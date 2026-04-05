package main

import (
	"image"
	"image/color"
	_ "image/png"
	"log"
	"os"
	"os/exec"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/paint"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

func main() {
	go func() {
		window := new(app.Window)
		err := run(window)
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

func run(window *app.Window) error {
	theme := material.NewTheme()
	var ops op.Ops

	imgFile, err := os.Open("thumb-1920-1145714.png")
	if err != nil {
		return err
	}
	defer imgFile.Close()

	maImage, _, err := image.Decode(imgFile)
	if err != nil {
		return err
	}
	imgWidget := widget.Image{Src: paint.NewImageOp(maImage), Fit: widget.Cover, Position: layout.Center}
	var startButton widget.Clickable

	for {
		switch e := window.Event().(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:

			gtx := app.NewContext(&ops, e)

			if startButton.Clicked(gtx) {
				exec.Command("explorer", "sword.mp4").Start()
			}

			title := material.H1(theme, "Peakland!")
			title.Font.Typeface = "Impact"

			red := color.NRGBA{R: 217, G: 43, B: 15, A: 255}
			title.Color = red

			title.Alignment = text.Middle

			layout.Stack{}.Layout(gtx,
				layout.Expanded(func(gtx layout.Context) layout.Dimensions {
					return imgWidget.Layout(gtx)
				}),
				layout.Stacked(func(gtx layout.Context) layout.Dimensions {
					return layout.Inset{Top: unit.Dp(16)}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						return title.Layout(gtx)
					})
				}),
				layout.Expanded(func(gtx layout.Context) layout.Dimensions {
					return layout.Flex{Axis: layout.Vertical, Alignment: layout.Middle}.Layout(gtx,
						layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
							return layout.Dimensions{Size: gtx.Constraints.Max}
						}),
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
								btn := material.Button(theme, &startButton, "Enter")
								btn.Background = color.NRGBA{R: 217, G: 43, B: 15, A: 255}
								return btn.Layout(gtx)
							})
						}),
						layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
							return layout.Dimensions{Size: gtx.Constraints.Max}
						}),
					)
				}),
			)

			e.Frame(gtx.Ops)
		}
	}
}
