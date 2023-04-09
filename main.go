package main

// Thank you for trying to develop NeuralStack Plugin!
// All those unused functions and variables should not be removed
// because the Plugin Loader requires them to load your plugin.

// This boilerplate used go embed fs to package frontend assets
// Make sure you run `yarn install` and `yarn run build` in views folder before compiling

// You should compile this project in Linux / macOS / FreeBSD,
// because Go Plugin System only supports those platforms.
// If you are a Windows user, please try WSL(Windows Subsystem Linux)

// The right command to compile this project is
// `go build -buildmode=plugin -o example.plug.so .`
// After compile, you will get a .plug.so file.
// Move this file into your NeuralStack instance /plugins folder
// After restart, this plugin will be loaded

// Get more info by read README.md

import (
	_ "embed"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	api "repo.smartsheep.studio/smartsheep/neuralstack-api"
	"smartsheep.studio/neuralstack/plugins/example/routes"
)

var (
	//go:embed views/dist/index.js
	script string
)

var Assets = api.PluginAssets{
	Apps: nil,
}

// Manifest for plugin loader, must name be Manifest
var Manifest = api.Plugin{
	Name:        "Example",
	PackageID:   "repo.smartsheep.studio/example/example",
	Description: "NeuralStack Plugin Boilerplate",
	Tags:        "Miscellaneous",
	Version:     "0.0.1",
	Author:      "You",

	Assets: &Assets,

	Setup:   Setup,
	Migrate: Migrate,
}

func Setup(p *api.Plugin, router gin.IRouter) {
	routes.Init(*p, router)

	p.Assets.Apps = []api.WebApp{
		{
			RootElement:        "plugins-example",
			DisplayOnLaunchpad: true,
			ID:                 "example",
			Icon:               "mdi-ab-testing",
			Name:               "Example",
			Descriptions:       "NeuralStack Example Plugin",
			WindowOptions: api.WebWindowOptions{
				Width:     800,
				Height:    400,
				Resizable: true,
				Closeable: true,
				Meta:      map[string]any{},
			},
			Script:    script,
			PackageID: p.PackageID,
		},
	}
}

func Migrate(datasource *gorm.DB) {

}
