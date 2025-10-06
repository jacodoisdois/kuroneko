package ui

import (
	"bytes"
	"image"
	_ "image/jpeg"
	"kuroneko/internal/handler"
	errorUtils "kuroneko/internal/ui/utils"
	"os/exec"

	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func NewCircularImage(res fyne.Resource) *canvas.Raster {
	reader := bytes.NewReader(res.Content())
	img, _, err := image.Decode(reader)
	if err != nil {
		fyne.LogError("Failed to decode image", err)
		size := 80
		dest := image.NewRGBA(image.Rect(0, 0, size, size))
		radius := float64(size) / 2
		centerX := radius
		centerY := radius
		for y := 0; y < size; y++ {
			for x := 0; x < size; x++ {
				dx := float64(x) - centerX
				dy := float64(y) - centerY
				if dx*dx+dy*dy <= radius*radius {
					dest.Set(x, y, color.NRGBA{R: 240, G: 237, B: 230, A: 255})
				}
			}
		}
		return canvas.NewRasterFromImage(dest)
	}
	size := 80
	dest := image.NewRGBA(image.Rect(0, 0, size, size))
	radius := float64(size) / 2
	centerX := radius
	centerY := radius
	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			dx := float64(x) - centerX
			dy := float64(y) - centerY
			if dx*dx+dy*dy <= radius*radius {
				srcX := int((float64(x)/float64(size))*float64(img.Bounds().Dx()) + float64(img.Bounds().Min.X))
				srcY := int((float64(y)/float64(size))*float64(img.Bounds().Dy()) + float64(img.Bounds().Min.Y))
				if srcX >= img.Bounds().Min.X && srcX < img.Bounds().Max.X && srcY >= img.Bounds().Min.Y && srcY < img.Bounds().Max.Y {
					dest.Set(x, y, img.At(srcX, srcY))
				}
			}
		}
	}
	return canvas.NewRasterFromImage(dest)
}

func CreateTitleContainer() fyne.CanvasObject {
	title := widget.NewLabel("Kuroneko 🐾")
	title.Alignment = fyne.TextAlignCenter
	title.TextStyle = fyne.TextStyle{Bold: true}

	image := NewCircularImage(resourceLogoJpg)
	image.SetMinSize(fyne.NewSize(80, 80))
	image.Resize(fyne.NewSize(80, 80))

	circle := canvas.NewCircle(color.Transparent)
	circle.Resize(fyne.NewSize(86, 86))
	circle.Move(fyne.NewPos(-3, -3))
	circle.StrokeColor = color.Black
	circle.StrokeWidth = 3

	logoContainer := container.NewStack(circle, container.NewCenter(image))

	subtitle := widget.NewLabel("Manga CBZ Creator")
	subtitle.Alignment = fyne.TextAlignCenter
	subtitle.TextStyle = fyne.TextStyle{Italic: true}

	titleContainer := container.NewVBox(
		container.NewCenter(logoContainer),
		container.NewCenter(title),
		container.NewCenter(subtitle),
	)

	return titleContainer
}

func CreateInputContainer(w fyne.Window, progressBar *widget.ProgressBar, statusLabel *widget.Label) fyne.CanvasObject {
	folderEntry := widget.NewEntry()
	folderEntry.SetPlaceHolder("Select a folder containing manga chapters...")
	folderEntry.Disable()

	prefs := fyne.CurrentApp().Preferences()
	lastPath := prefs.String("lastFolderPath")
	if lastPath != "" {
		folderEntry.SetText(lastPath)
		statusLabel.SetText("Last folder loaded. Ready to create CBZ files.")
	}

	folderButton := widget.NewButtonWithIcon("Browse", theme.FolderOpenIcon(), func() {
		d := dialog.NewFolderOpen(func(uri fyne.ListableURI, err error) {
			if uri != nil {
				folderPath := uri.Path()
				folderEntry.SetText(folderPath)
				prefs.SetString("lastFolderPath", folderPath)
				statusLabel.SetText("Folder selected. Click 'Create CBZ' to process.")
			}
		}, w)
		d.Show()
	})

	folderContainer := container.NewBorder(
		nil, nil, nil, folderButton,
		container.NewPadded(folderEntry),
	)

	progressBar.Hide()

	runButton := widget.NewButtonWithIcon("Create CBZ", resourcePawIcon2Png, func() {
		folderPath := folderEntry.Text
		if folderPath == "" {
			errorUtils.ThrowError(w, "Please select a folder first.")
			return
		}
		statusLabel.SetText("Processing... Please wait.")
		progressBar.Show()
		progressBar.SetValue(0.1)

		go func() {
			handler.GenerateCbzFiles(folderPath, w, func(progress float64) {
				progressBar.SetValue(0.3 + progress*0.7)
			})
			progressBar.Hide()
		}()
	})

	openButton := widget.NewButtonWithIcon("Open in Finder", theme.ComputerIcon(), func() {
		folderPath := folderEntry.Text
		if folderPath == "" {
			errorUtils.ThrowError(w, "Please select a folder first.")
			return
		}
		exec.Command("open", folderPath).Start()
	})

	buttonsContainer := container.NewHBox(runButton, openButton)

	inputContainer := container.NewVBox(
		statusLabel,
		widget.NewSeparator(),
		container.NewPadded(widget.NewLabel("Select Manga Folder:")),
		folderContainer,
		container.NewPadded(progressBar),
		container.NewPadded(buttonsContainer),
	)

	return container.NewPadded(inputContainer)
}

func NewMainWindow(a fyne.App) fyne.Window {
	window := a.NewWindow("Kuroneko - Manga CBZ Creator")
	window.Resize(fyne.NewSize(700, 500))
	window.SetIcon(resourcePawIcon2Png)

	titleContainer := CreateTitleContainer()

	statusLabel := widget.NewLabel("Ready to create CBZ files")
	statusLabel.Alignment = fyne.TextAlignCenter
	statusLabel.TextStyle = fyne.TextStyle{Italic: true}

	progressBar := widget.NewProgressBar()

	inputContainer := CreateInputContainer(window, progressBar, statusLabel)

	footer := widget.NewLabel("🐾 Support Kuroneko on GitHub • Star the project if you like it!")
	footer.Alignment = fyne.TextAlignCenter
	footer.TextStyle = fyne.TextStyle{Italic: true}

	content := container.NewVBox(
		container.NewPadded(titleContainer),
		widget.NewSeparator(),
		inputContainer,
		widget.NewSeparator(),
		container.NewCenter(footer),
	)

	window.SetContent(container.NewPadded(content))
	return window
}
