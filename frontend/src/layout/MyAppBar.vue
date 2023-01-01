<template>
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
        <v-btn icon @click="toggleTheme">
            <v-icon>mdi-theme-light-dark</v-icon>
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
</template>

<script>
import {useTheme} from "vuetify"
export default {
    name: "MyAppBar",
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
    setup() {
        const theme = useTheme()
        const toggleTheme = () => {
            theme.global.name.value = theme.global.current.value.dark
                ? "light"
                : "dark"
        }
        return {toggleTheme}
    },
}
</script>
