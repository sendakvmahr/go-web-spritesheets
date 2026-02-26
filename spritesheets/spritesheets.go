package spritesheets

type SpriteSlice struct {
	Image string
	X int
	Y int
	Width int
	Height int
}

type SpriteSheetResponse struct {
	Images []string
	ImageDict map[string]SpriteSlice
}