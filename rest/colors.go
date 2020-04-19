package rest

type Color string

const (
	Red    Color = "#ff0000"
	Green  Color = "#0eff00"
	Gray   Color = "#808080"
	Yellow Color = "#ffff00"
	Blue   Color = "#0000ff"
	Purple Color = "#800080"
	White  Color = "#ffffff"
	Black  Color = "#000000"
)

func (c Color) String() string {
	return string(c)
}
