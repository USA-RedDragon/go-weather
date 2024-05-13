package app

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	xwidget "fyne.io/x/fyne/widget"
)

type App struct {
	app    fyne.App
	window fyne.Window
	appBar fyne.CanvasObject
}

func New() *App {
	app := app.NewWithID("dev.mcswain.radar")

	ret := &App{
		app: app,
	}

	menuButton := widget.NewButtonWithIcon("", theme.MenuIcon(), ret.toggleAppBar)
	appBar := container.NewVBox(menuButton)
	ret.appBar = appBar

	mapView := container.NewStack()
	mapView.Add(xwidget.NewMap())
	mapView.Add(widget.NewLabel("Component name"))

	content := container.NewBorder(nil, nil, appBar, nil, mapView)
	ret.window = app.NewWindow("Radar")
	ret.window.SetContent(content)
	return ret
}

func (a *App) Run() {
	if !fyne.CurrentDevice().IsMobile() {
		a.window.Resize(fyne.NewSize(1280, 720))
	}
	a.window.Show()
	a.app.Run()
}

func (a *App) toggleAppBar() {
	fmt.Println("Menu clicked")
}
