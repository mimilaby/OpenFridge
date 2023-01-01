import {fileURLToPath, URL} from "node:url"

import {defineConfig, loadEnv} from "vite"
import vue from "@vitejs/plugin-vue"
import vuetify from "vite-plugin-vuetify"
import process from "node:process"

// https://vitejs.dev/config/
export default defineConfig(({command, mode}) => {
    const env = loadEnv(mode, process.cwd())
    if (command === "build") {
        return {
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
                        target: `${env.VITE_BACKEND_URL}/api`,
                        // changeOrigin: true,
                        rewrite: (path) => path.replace(/^\/api/, ""),
                    },
                },
            },
        }
    } else if (command === "serve") {
        return {
            plugins: [vue(), vuetify()],
            resolve: {
                alias: {
                    "@": fileURLToPath(new URL("./src", import.meta.url)),
                },
            },
            server: {
                proxy: {
                    "/api": {
                        target: `${env.VITE_BACKEND_URL}/api`,
                        changeOrigin: true,
                        rewrite: (path) => path.replace(/^\/api/, ""),
                    },
                },
            },
        }
    } 
})
