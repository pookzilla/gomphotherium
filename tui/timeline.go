package tui

import (
	"fmt"

	"github.com/mrusme/gomphotherium/mast"
)

func RenderTimeline(
	timeline *mast.Timeline,
	width int,
	showImages bool,
	justifyText bool) (string, error) {
	var output string = ""
	var err error = nil

	var tootOutput string = ""
	newRenderedIndex := len(timeline.Toots)
	for i := (timeline.LastRenderedIndex + 1); i < newRenderedIndex; i++ {
		toot := &timeline.Toots[i]
		isFocus := timeline.Focus != nil && toot.Status.ID == timeline.Focus.Status.ID
		tootWidth := width
		if isFocus {
			tootWidth -= 4
		}
		tootOutput, err = RenderToot(toot, tootWidth, showImages, justifyText, isFocus)
		output = fmt.Sprintf("%s%s\n", output, tootOutput)
	}

	timeline.LastRenderedIndex = (newRenderedIndex - 1)
	return output, err
}
