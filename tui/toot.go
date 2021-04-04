package tui

import (
  "fmt"
  // "time"
  // "context"

  "github.com/mattn/go-runewidth"
  "github.com/grokify/html-strip-tags-go"
  "html"

  "image/color"
  "github.com/eliukblau/pixterm/pkg/ansimage"

  // "github.com/mattn/go-mastodon"
  "github.com/mrusme/gomphotherium/mast"
)

func RenderToot(toot *mast.Toot, width int, showImages bool) (string, error) {
  var output string = ""
  var err error = nil

  status := &toot.Status

  createdAt := status.CreatedAt

  account := status.Account.Acct
  if account == "" {
    account = status.Account.Username
  }

  inReplyTo := ""
  if status.InReplyToID != nil {
    inReplyTo = " \xe2\x87\x9f"
  }

  idPadding :=
    width -
    len(string(toot.ID)) -
    runewidth.StringWidth(status.Account.DisplayName) -
    len(account) -
    // https://github.com/mattn/go-runewidth/issues/36
    runewidth.StringWidth(inReplyTo)

  output = fmt.Sprintf("%s[blue]%s[-] [grey]%s[-][magenta]%s[-][grey]%*d[-]\n",
    output,
    status.Account.DisplayName,
    account,
    inReplyTo,
    idPadding,
    toot.ID,
  )
  output = fmt.Sprintf("%s%s\n",
    output,
    html.UnescapeString(strip.StripTags(status.Content)),
  )

  if showImages == true {
    for _, attachment := range status.MediaAttachments {
      pix, err := ansimage.NewScaledFromURL(
        attachment.PreviewURL,
        int((float64(width) * 0.75)),
        width,
        color.Transparent,
        ansimage.ScaleModeResize,
        ansimage.NoDithering,
      )
      if err == nil {
        output = fmt.Sprintf("%s\n%s\n", output, pix.RenderExt(false, false))
      }
    }
  }

  output = fmt.Sprintf("%s[magenta]\xe2\x86\xab %d[-] ",
    output,
    status.RepliesCount,
  )
  output = fmt.Sprintf("%s[green]\xe2\x86\xbb %d[-] ",
    output,
    status.ReblogsCount,
  )
  output = fmt.Sprintf("%s[yellow]\xe2\x98\x85 %d[-] ",
    output,
    status.FavouritesCount,
  )
  output = fmt.Sprintf("%s[grey]on %s at %s[-]\n",
    output,
    createdAt.Format("Jan 2"),
    createdAt.Format("15:04"),
  )

  output = fmt.Sprintf("%s\n", output)
  return output, err
}
