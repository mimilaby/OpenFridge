<!-- App.vue -->

<template>
    <v-app>
        <v-app-bar>
            <v-app-bar-nav-icon
                variant="text"
                @click.stop="drawer = !drawer"
            ></v-app-bar-nav-icon>
            <v-app-bar-title class="pl-0">
                <div class="d-flex align-center">
                    <v-avatar
                        rounded="0"
                        class="mr-3"
                        image="https://cdn.vuetifyjs.com/docs/images/logos/v.png"
                    />
                    HomeApp
                </div>
            </v-app-bar-title>

            <v-spacer></v-spacer>

            <v-btn icon>
                <v-icon>mdi-magnify</v-icon>
            </v-btn>
            <v-btn icon>
                <v-icon>mdi-dots-vertical</v-icon>
            </v-btn>
        </v-app-bar>
        <v-navigation-drawer v-model="drawer" bottom permanent>
            <v-list>
                <v-list-item
                    v-for="(item, i) in items"
                    :key="i"
                    :title="item.title"
                    :to="item.to"
                    exact
                    active-class="primary"
                ></v-list-item>
            </v-list>
        </v-navigation-drawer>

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
</script>
