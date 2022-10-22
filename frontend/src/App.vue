<!-- App.vue -->


<!-- <script>
export default {
    name: "App",
    data() {
        return {
            drawer: false,
            group: null,
            items: [
                {
                    title: "Home",
                    to: "/home",
                },
                {
                    title: "About",
                    to: "/about",
                },
            ],
        }
    },
    watch: {
        group() {
            this.drawer = false
        },
    },
}
</script> -->

<template>
    <v-app>
        <MyAppBar />

        <v-main app>
            <v-container>
                <router-view />
            </v-container>
        </v-main>

        <v-footer app>
            <v-container
                class="text-overline d-flex align-center justify-space-between"
            >
                <div>Copyright &copy; 2022 HomeApp</div>
                <v-icon icon="mdi-vuetify" size="x-large"></v-icon>
            </v-container>
        </v-footer>
    </v-app>
</template>

<script>
import MyAppBar from "./layout/MyAppBar.vue"
export default {
    name: "App",
    components: {
        MyAppBar,
    },
}
</script>
