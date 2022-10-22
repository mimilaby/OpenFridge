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
        defaultTheme: "dark",
        themes: {
            dark: {
                dark: true,
                colors: {
                    primary: "#1697f6",
                },
            },
        },
    },
})
