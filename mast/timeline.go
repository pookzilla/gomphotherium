package mast

import (
	"context"
	"errors"

	"github.com/mattn/go-mastodon"
)

type TimelineType int

const (
	TimelineHome          TimelineType = 0
	TimelineLocal                      = 1
	TimelinePublic                     = 2
	TimelineNotifications              = 3
	TimelineHashtag                    = 4
	TimelineUser                       = 5
	TimelineThread                     = 6
)

type TimelineOptions struct {
	IsLocal    bool
	Hashtag    string
	User       mastodon.Account
	ThreadToot mastodon.Status
}

type Timeline struct {
	client          *mastodon.Client
	timelineType    TimelineType
	timelineOptions TimelineOptions

	LastRenderedIndex int

	Account                   mastodon.Account
	Toots                     []Toot
	TootIndexStatusIDMappings map[string]int
	KnownUsers                map[string]string
	Focus                     *Toot
}

func NewTimeline(mastodonClient *mastodon.Client) Timeline {
	timeline := Timeline{
		client:       mastodonClient,
		timelineType: TimelineHome,

		LastRenderedIndex:         -1,
		TootIndexStatusIDMappings: make(map[string]int),
		KnownUsers:                make(map[string]string),
	}

	return timeline
}

func (timeline *Timeline) Switch(timelineType TimelineType, options *TimelineOptions) {
	if timeline.timelineType != timelineType ||
		timelineType == TimelineHashtag ||
		timelineType == TimelineUser ||
		timelineType == TimelineThread {
		timeline.timelineType = timelineType
		if options != nil {
			timeline.timelineOptions = *options
		}
		timeline.Toots = []Toot{}
		timeline.TootIndexStatusIDMappings = make(map[string]int)
		timeline.LastRenderedIndex = -1
	}
}

func (timeline *Timeline) GetCurrentType() TimelineType {
	return timeline.timelineType
}

func (timeline *Timeline) GetCurrentOptions() TimelineOptions {
	return timeline.timelineOptions
}

func (timeline *Timeline) Load() error {
	var statuses []*mastodon.Status
	var notifications []*mastodon.Notification
	var err error
	var focusStatus *mastodon.Status

	timeline.Focus = nil

	account, err := timeline.client.GetAccountCurrentUser(context.Background())
	if err != nil {
		return err
	}

	timeline.Account = *account

	switch timeline.timelineType {
	case TimelineHome:
		statuses, err =
			timeline.client.GetTimelineHome(context.Background(), nil)
	case TimelineLocal:
		statuses, err =
			timeline.client.GetTimelinePublic(context.Background(), true, nil)
	case TimelinePublic:
		statuses, err =
			timeline.client.GetTimelinePublic(context.Background(), false, nil)
	case TimelineNotifications:
		notifications, err =
			timeline.client.GetNotifications(context.Background(), nil)
		if err != nil {
			return err
		}

		for _, notification := range notifications {
			statuses = append(statuses, notification.Status)
		}
	case TimelineHashtag:
		statuses, err =
			timeline.client.GetTimelineHashtag(
				context.Background(),
				timeline.timelineOptions.Hashtag,
				timeline.timelineOptions.IsLocal,
				nil,
			)
	case TimelineUser:
		statuses, err =
			timeline.client.GetAccountStatuses(
				context.Background(),
				timeline.timelineOptions.User.ID,
				nil,
			)
	case TimelineThread:
		context, err := timeline.client.GetStatusContext(
			context.Background(),
			timeline.timelineOptions.ThreadToot.ID,
		)

		focusStatus = &timeline.timelineOptions.ThreadToot
		if err != nil {
			statuses = make([]*mastodon.Status, 1)
			statuses = append(statuses, &timeline.timelineOptions.ThreadToot)
		} else {
			statuses = make([]*mastodon.Status, len(context.Ancestors)+len(context.Descendants)+1)
			statuses = context.Ancestors
			statuses = append(statuses, &timeline.timelineOptions.ThreadToot)
			statuses = append(statuses, context.Descendants...)
		}

		// reverse
		for i, j := 0, len(statuses)-1; i < j; i, j = i+1, j-1 {
			statuses[i], statuses[j] = statuses[j], statuses[i]
		}
	}

	if err != nil {
		return err
	}

	oldestStatusIndex := len(statuses) - 1
	oldestNotificationIndex := len(notifications) - 1
	for i := oldestStatusIndex; i >= 0; i-- {
		status := statuses[i]
		if status == nil {
			continue
		}
		var notification *mastodon.Notification = nil

		if oldestNotificationIndex >= i {
			notification = notifications[i]
		}

		statusID := string(status.ID)
		tootIndex, exists := timeline.TootIndexStatusIDMappings[statusID]
		if exists == false {
			tootIndex := len(timeline.Toots)
			toot := NewToot(timeline.client, status, notification, tootIndex)
			if focusStatus != nil && focusStatus.ID == toot.Status.ID {
				timeline.Focus = &toot
			}
			timeline.Toots =
				append(timeline.Toots, toot)

			timeline.TootIndexStatusIDMappings[statusID] = tootIndex
			timeline.KnownUsers[string(status.Account.ID)] = status.Account.Acct
		} else if focusStatus != nil {
			toot := timeline.Toots[tootIndex]
			if focusStatus.ID == toot.Status.ID {
				timeline.Focus = &toot
			}
		}
	}

	return nil
}

func (timeline *Timeline) Toot(
	status *string,
	inReplyTo int,
	filesToUpload []string,
	visibility *string,
	sensitive bool,
	spoiler *string) (*mastodon.Status, error) {
	newToot := mastodon.Toot{
		Status:      *status,
		Visibility:  *visibility,
		Sensitive:   sensitive,
		SpoilerText: *spoiler,
	}

	if inReplyTo > -1 {
		if inReplyTo >= len(timeline.Toots) {
			return nil, errors.New("Replied-to toot does not exist")
		}
		newToot.InReplyToID = timeline.Toots[inReplyTo].Status.ID
	}

	if len(filesToUpload) > 0 {
		var mediaIDs []mastodon.ID

		for _, fileToUpload := range filesToUpload {
			attachment, err :=
				timeline.client.UploadMedia(context.Background(), fileToUpload)
			if err != nil {
				return nil, err
			}

			mediaIDs = append(mediaIDs, attachment.ID)
		}

		newToot.MediaIDs = mediaIDs
	}

	return timeline.client.PostStatus(context.Background(), &newToot)
}

func (timeline *Timeline) Boost(
	tootID int,
	shallBe bool) (*mastodon.Status, error) {
	id := timeline.Toots[tootID].Status.ID

	if shallBe == true {
		return timeline.client.Reblog(context.Background(), id)
	}
	return timeline.client.Unreblog(context.Background(), id)
}

func (timeline *Timeline) Fav(
	tootID int,
	shallBe bool) (*mastodon.Status, error) {
	id := timeline.Toots[tootID].Status.ID

	if shallBe == true {
		return timeline.client.Favourite(context.Background(), id)
	}
	return timeline.client.Unfavourite(context.Background(), id)
}

func (timeline *Timeline) User(
	ID string) (*mastodon.Account, error) {
	return timeline.client.GetAccount(context.Background(), mastodon.ID(ID))
}

func (timeline *Timeline) SearchUser(
	query string,
	limit int64) ([]*mastodon.Account, error) {
	return timeline.client.AccountsSearch(context.Background(), query, limit)
}
