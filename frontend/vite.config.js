import {fileURLToPath, URL} from "node:url"

import {defineConfig} from "vite"
import vue from "@vitejs/plugin-vue"
import vuetify from "vite-plugin-vuetify"

// https://vitejs.dev/config/
export default defineConfig({
    plugins: [vue(), vuetify()],
    resolve: {
        alias: {
            "@": fileURLToPath(new URL("./src", import.meta.url)),
        },
    },
    server: {
        fs: {
            allow: [
                // search up for workspace root
                // searchForWorkspaceRoot(process.cwd()),
                // your custom rules
                "..",
            ],
        },
        proxy: {
            "/api": {
                target: "http://localhost:8000",
                changeOrigin: true,
                rewrite: (path) => path.replace(/^\/api/, ""),
            },
        },
    },
})
