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
		tootWidth := width
		if i == timeline.Focus {
			tootWidth -= 4
		}
		tootOutput, err = RenderToot(&timeline.Toots[i], tootWidth, showImages, justifyText, i == timeline.Focus)
		output = fmt.Sprintf("%s%s\n", output, tootOutput)
	}

	timeline.LastRenderedIndex = (newRenderedIndex - 1)
	return output, err
}
