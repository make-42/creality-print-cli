package components

import (
	"creality-print-cli/config"
	"creality-print-cli/constants"
	"creality-print-cli/styling"
	"fmt"
	"math"
	"strings"

	"github.com/muesli/termenv"
)

func VersionNumber() string {
	return "creality-print-cli " + styling.ColorFg(constants.AppVersion, styling.HighlightedColor)
}

func KeybindsHints(keybinds []string) string {
	s := ""
	for index, keybind := range keybinds {
		if index != 0 {
			s += styling.Dot
		}
		s += styling.Subtle(keybind)
	}
	return styling.Indent(s, config.Config.UIPaddingIndentAmount)
}

func Checkbox(label string, checked bool, selected bool) string {
	s := fmt.Sprintf("[ ] %s", label)
	if checked {
		s = "[x] " + label
	}
	if selected {
		return styling.ColorFg(s, styling.HighlightedColor)
	}
	return s
}

func Progressbar(width int, percent float64) string {
	w := float64(styling.ProgressBarWidth)

	fullSize := int(math.Round(w * percent))
	var fullCells string
	for i := 0; i < fullSize; i++ {
		fullCells += termenv.String(styling.ProgressFullChar).Foreground(styling.Term.Color(styling.Ramp[i])).String()
	}

	emptySize := int(w) - fullSize
	emptyCells := strings.Repeat(styling.ProgressEmpty, emptySize)

	return fmt.Sprintf("%s%s %3.0f", fullCells, emptyCells, math.Round(percent*100))
}

func Header(header string) string {
	return styling.Bold(styling.Indent(header, config.Config.UIPaddingIndentAmount))
}
