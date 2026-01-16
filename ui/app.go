package ui

import "github.com/rivo/tview"

type App struct {
	App   *tview.Application
	Pages *tview.Pages
}

func NewApp() *App {
	app := tview.NewApplication()
	pages := tview.NewPages()

	// Logo view
	logo := tview.NewTextView().
		SetTextAlign(tview.AlignCenter).
		SetText(Logo()).
		SetDynamicColors(true)

	pages.AddPage("logo", logo, true, true)

	return &App{
		App:   app,
		Pages: pages,
	}
}

func (a *App) Run() error {
	a.App.SetRoot(a.Pages, true)
	a.App.SetFocus(a.Pages)
	return a.App.Run()
}
