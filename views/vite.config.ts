import { fileURLToPath, URL } from "node:url"

import { defineConfig } from "vite"
import vue from "@vitejs/plugin-vue"
import cssInjectedByJsPlugin from "vite-plugin-css-injected-by-js"
import { quasar, transformAssetUrls } from "@quasar/vite-plugin"

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    vue({ transformAssetUrls }),
    cssInjectedByJsPlugin(),
    quasar({ sassVariables: "src/quasar-variables.sass" }),
  ],
  build: {
    rollupOptions: {
      // @ts-ignore
      output: {
        manualChunks: false,
        inlineDynamicImports: true,
        entryFileNames: "[name].js",
        assetFileNames: "[name].[ext]",
      },
    },
  },
  resolve: {
    alias: {
      "@": fileURLToPath(new URL("./src", import.meta.url)),
    },
  },
  server: {
    proxy: {
      "/api": "http://localhost:9443",
      "/o/oauth/token": "http://localhost:9443",
      "/.well-known": "http://localhost:9443",
    },
  },
})
