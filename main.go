package main

import (
	"embed"
	"log"
	"order/api"
	"order/api/pkg/logging"
	"order/api/service"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

// go:embed all:frontend/dist
var assets embed.FS

func init() {
	log.SetOutput(logging.FilePtr("system.log"))
}

func main() {
	app := api.NewApp()

	err := wails.Run(&options.App{
		Title:  "进销存管理系统",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.Startup,
		Bind: []interface{}{
			app, service.NewService(),
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
