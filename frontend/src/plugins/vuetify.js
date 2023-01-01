// Styles
import "@mdi/font/css/materialdesignicons.css"
import "vuetify/styles"

// Vuetify
import {createVuetify} from "vuetify"

// Components
import {VBtn} from "vuetify/components"

export default createVuetify({
    aliases: {
        VBtnAlt: VBtn,
    },
    // https://next.vuetifyjs.com/features/global-configuration/
    defaults: {
        global: {
            rounded: "sm",
        },
        VAppBar: {
            flat: true,
        },
        VBtn: {
            color: "primary",
            height: 44,
        },
        VBtnAlt: {
            color: "primary",
            height: 48,
            variant: "text",
        },
        VSheet: {
            color: "#212121",
        },
    },
    // https://next.vuetifyjs.com/features/theme/
    theme: {
        defaultTheme: "light",
        themes: {
            dark: {
                dark: true,
                colors: {
                    primary: "#1697f6",
                },
            },
            light: {
                dark: false,
                colors: {
                    background: "#FFFFFF",
                    surface: "#F0F0F6",
                    primary: "#2f90b9",
                    "primary-darken-1": "#3700B3",
                    secondary: "#03DAC6",
                    "secondary-darken-1": "#018786",
                    error: "#B00020",
                    info: "#2196F3",
                    success: "#4CAF50",
                    warning: "#FB8C00",
                },
            },
        },
    },
})
