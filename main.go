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
	"embed"
	"encoding/json"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	api "repo.smartsheep.studio/atom/electron"
	"repo.smartsheep.studio/atom/electron/utils"
	"smartsheep.studio/neuralstack/plugins/example/routes"
)

var (
	//go:embed all:views/dist
	views embed.FS

	//go:embed manifest.json
	rawManifest []byte
)

var Assets = api.PluginAssets{
	Apps: nil,
}

// Manifest for plugin loader, must name be Manifest
var Manifest = api.Plugin{
	Assets: &Assets,

	Setup:   Setup,
	Migrate: Migrate,
	Init: func(p *api.Plugin) {
		var manifest api.PluginManifest
		json.Unmarshal(rawManifest, &manifest)
		p.Manifest = manifest
	},
}

// Setup function will call when loading plugin.
// Usually do mapping routes and read configurations.
func Setup(p *api.Plugin, router gin.IRouter) {
	routes.Init(*p, router)

	p.Assets.Apps = []api.WebApp{
		{
			RootFile:    "index.html",                 // Root file when no path, generally no need to change
			ID:          "example",                    // App ID, if you changed this, make sure you updated the base path in views/vite.config.ts
			Icon:        "mdi-ab-testing",             // App Logo, display on title and launchpad, support mdi v7 icons(https://materialdesignicons.com)
			Name:        "Example",                    // App Title and Label display in launchpad
			Description: "NeuralStack Example Plugin", // App Description
			Permissions: &api.WebAppPermissionsConfig{},
			Assets:      utils.EmbedFolder(views, "views/dist"),
			PackageID:   p.Manifest.Package,
		},
	}
}

// Migrate function will call when datasource is ready.
// Usually do migrate things like auto creating table.
// Recommended use utils.HandleFatalError to handle error.
func Migrate(datasource *gorm.DB) {

}
