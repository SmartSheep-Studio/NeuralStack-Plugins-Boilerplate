# NeuralStack Plugin Boilerplate

Before using this boilerplate to create your own NeuralStack Plugin,
make sure you have installed Golang the latest version,
NodeJS the latest version and any WSL distribution(for Windows users)

For testing your plugin, you need to install a NeuralStack instance on your develop machine.
Visit [README.md](https://repo.smartsheep.studio/SmartSheep/NeuralStack/src/branch/master/README.md) for some details.

This boilerplate is created with Vue.js 3, TypeScript, Pinia and Quasar, you can change any frontend library you like.
But a Single Page Application framework is recommended.

## Getting Start

Before compile go package. You need to compile the frontend.

Switch work directory into `views` and run:

```shell
$ npm install
$ npm run build # or npm run build-only for skip type checking
```

Then, switch work directory to project root, run those commands to compile plugin:

```shell
$ go build -buildmode=plugin -o example.plug.so .
```

After that, you can find a `example.plug.so` file in your project root. Move this file into your NeuralStack instance /plugins folder and restart it.

Boom! The Example App is displayed on your launchpad!

Explore more API details in our knowledge base in the wiki tab of our repository.

## Start Developing

To start developing plugin. You need to read those tips.

1. When you are trying to update something, make sure you have to read the comments in the code.
2. Look for help on forum before you create an issue.
3. Explore other plugins source code to learn how to take your plugin to the next level