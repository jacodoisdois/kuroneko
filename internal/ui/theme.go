package ui

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type OffWhiteTheme struct{}

var _ fyne.Theme = (*OffWhiteTheme)(nil)

// Off-white color palette
var (
	offWhiteBackground = color.NRGBA{R: 250, G: 249, B: 246, A: 255}
	offWhiteForeground = color.Black
	offWhitePrimary    = color.NRGBA{R: 240, G: 237, B: 230, A: 255}
	offWhiteSurface    = color.NRGBA{R: 248, G: 246, B: 240, A: 255}
	offWhiteShadow     = color.NRGBA{R: 235, G: 233, B: 228, A: 255}
	offWhiteSeparator  = color.NRGBA{R: 245, G: 243, B: 238, A: 255}
	offWhiteDisabled   = color.NRGBA{R: 200, G: 195, B: 190, A: 255}
)

func (t *OffWhiteTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	switch name {
	case theme.ColorNameBackground:
		return offWhiteBackground
	case theme.ColorNameForeground:
		return offWhiteForeground
	case theme.ColorNamePrimary:
		return offWhitePrimary
	case theme.ColorNameButton:
		return offWhiteSurface
	case theme.ColorNameDisabled:
		return offWhiteDisabled
	case theme.ColorNameDisabledButton:
		return offWhiteDisabled
	case theme.ColorNameError:
		return color.NRGBA{R: 180, G: 50, B: 50, A: 255}
	case theme.ColorNameFocus:
		return color.NRGBA{R: 200, G: 200, B: 200, A: 255}
	case theme.ColorNameHover:
		return color.NRGBA{R: 230, G: 230, B: 230, A: 255}
	case theme.ColorNameInputBackground:
		return offWhiteSurface
	case theme.ColorNameMenuBackground:
		return offWhiteBackground
	case theme.ColorNameOverlayBackground:
		return color.NRGBA{R: 250, G: 249, B: 246, A: 180}
	case theme.ColorNamePlaceHolder:
		return color.NRGBA{R: 150, G: 150, B: 150, A: 255}
	case theme.ColorNamePressed:
		return color.NRGBA{R: 200, G: 200, B: 200, A: 255}
	case theme.ColorNameScrollBar:
		return color.NRGBA{R: 180, G: 180, B: 180, A: 255}
	case theme.ColorNameSelection:
		return color.NRGBA{R: 200, G: 220, B: 240, A: 255}
	case theme.ColorNameSeparator:
		return offWhiteSeparator
	case theme.ColorNameShadow:
		return offWhiteShadow
	default:
		return theme.DefaultTheme().Color(name, variant)
	}
}

func (t *OffWhiteTheme) Font(style fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(style)
}

func (t *OffWhiteTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

func (t *OffWhiteTheme) Size(name fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(name)
}
