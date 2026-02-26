package spritesheets

type SpriteSlice struct {
	Image string
	X int
	Y int
	Width int
	Height int
}

type SpriteSheetResponse struct {
	Images map[string]string
	ImageDict map[string]SpriteSlice
}