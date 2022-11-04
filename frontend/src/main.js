import {createApp} from "vue"
import App from "./App.vue"
import router from "./router"

import axios from "axios"
import VueAxios from "vue-axios"
import registerPlugins from "@/plugins"
import vuetify from "@/plugins/vuetify"

const app = createApp(App)
registerPlugins(app)

app.use(router)
app.use(vuetify)
app.use(VueAxios, axios)

app.mount("#app")
