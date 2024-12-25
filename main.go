package main

import (
	"creality-print-cli/components"
	"creality-print-cli/config"
	"creality-print-cli/data"
	"creality-print-cli/styling"
	"fmt"
	"os"
	"strconv"
	"time"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/buger/goterm"
)

type model struct {
}

func initialModel() model {
	return model{}
}

func (m model) Init() tea.Cmd {
	return tea.Batch(tickCmd(), tea.SetWindowTitle("Creality Print CLI"))
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}

	case tickMsg:
		return m, tickCmd()
	}

	return m, nil
}

type tickMsg time.Time

func tickCmd() tea.Cmd {
	return tea.Tick(time.Millisecond*time.Duration(config.Config.UpdateUIEveryXMilliseconds), func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func (m model) View() string {
	s := styling.Indent(components.VersionNumber()+"\n\n", config.Config.UIPaddingIndentAmount)
	s += styling.Indent(fmt.Sprintf("%s\n\n", data.CurrentData.CurPosition), config.Config.UIPaddingIndentAmount)
	nozzleTemp, _ := strconv.ParseFloat(data.CurrentData.NozzleTemp, 64)
	bedTemp, _ := strconv.ParseFloat(data.CurrentData.BedTemp0, 64)
	speed, _ := strconv.ParseFloat(data.CurrentData.RealTimeSpeed, 64)

	s += styling.Indent(fmt.Sprintf("Speed: %s mm/s\n\n", styling.ColorFg(fmt.Sprintf("%0.0f", speed), styling.HighlightedColor)), config.Config.UIPaddingIndentAmount)

	s += styling.Indent(fmt.Sprintf("Nozzle: %s/%d.00 °C\n", styling.ColorFg(fmt.Sprintf("%0.2f", nozzleTemp), styling.HighlightedColor), data.CurrentData.TargetNozzleTemp), config.Config.UIPaddingIndentAmount)
	s += styling.Indent(fmt.Sprintf("Bed: %s/%d.00 °C\n", styling.ColorFg(fmt.Sprintf("%0.2f", bedTemp), styling.HighlightedColor), data.CurrentData.TargetBedTemp0), config.Config.UIPaddingIndentAmount)
	s += styling.Indent(fmt.Sprintf("Enclosure: %s/%d.00 °C\n\n", styling.ColorFg(fmt.Sprintf("%d.00", data.CurrentData.BoxTemp), styling.HighlightedColor), data.CurrentData.BoxTemp), config.Config.UIPaddingIndentAmount)

	s += styling.Indent(fmt.Sprintf("Model Fan: %s %%\n", styling.ColorFg(fmt.Sprintf("%d", data.CurrentData.ModelFanPct), styling.HighlightedColor)), config.Config.UIPaddingIndentAmount)
	s += styling.Indent(fmt.Sprintf("Aux Fan: %s %%\n", styling.ColorFg(fmt.Sprintf("%d", data.CurrentData.AuxiliaryFanPct), styling.HighlightedColor)), config.Config.UIPaddingIndentAmount)
	s += styling.Indent(fmt.Sprintf("Enclosure Fan: %s %%\n\n", styling.ColorFg(fmt.Sprintf("%d", data.CurrentData.CaseFanPct), styling.HighlightedColor)), config.Config.UIPaddingIndentAmount)

	s += styling.Indent(fmt.Sprintf("Layer %s/%d\n\n", styling.ColorFg(fmt.Sprintf("%d", data.CurrentData.Layer), styling.HighlightedColor), data.CurrentData.TotalLayer), config.Config.UIPaddingIndentAmount)

	s += styling.Indent(fmt.Sprintf("Used %s m of filament\n\n", styling.ColorFg(fmt.Sprintf("%0.2f", float64(data.CurrentData.UsedMaterialLength)/1000), styling.HighlightedColor)), config.Config.UIPaddingIndentAmount)

	s += styling.Indent(fmt.Sprintf("Has been printing for %s\n", styling.ColorFg((time.Duration(data.CurrentData.PrintJobTime)*time.Second).String(), styling.HighlightedColor)), config.Config.UIPaddingIndentAmount)
	s += styling.Indent(fmt.Sprintf("Print finishes in %s\n\n", styling.ColorFg((time.Duration(data.CurrentData.PrintLeftTime)*time.Second).String(), styling.HighlightedColor)), config.Config.UIPaddingIndentAmount)

	width := goterm.Width()
	s += styling.Indent(components.Progressbar(width-config.Config.UIPaddingIndentAmount*2, float64(data.CurrentData.PrintProgress)/100)+"\n\n", config.Config.UIPaddingIndentAmount)
	s += components.KeybindsHints([]string{"q: quit"})

	return s
}

func main() {
	config.Init()
	go data.Init()
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}

}
