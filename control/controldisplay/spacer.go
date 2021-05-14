package controldisplay

import (
	"fmt"
	"strings"
)

type SpacerRenderer struct {
	width int
}

func NewSpacerRenderer(width int) *SpacerRenderer {
	return &SpacerRenderer{width}
}

// Render returns a divider string os format: "......."
func (d SpacerRenderer) Render() (string, int) {
	// we always have a leading and trailing space
	if d.width < 2 {
		return strings.Repeat(" ", d.width), d.width
	}

	// allow for spaces
	numberOfDots := d.width - 2
	return fmt.Sprintf(" %s ", colorSpacer(strings.Repeat(".", numberOfDots))), d.width
}
