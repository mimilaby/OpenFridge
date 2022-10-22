<template>
    <v-container>
        <TableFood />
    </v-container>
</template>

<script>
import TableFood from "../components/TableFood.vue"
export default {
    name: "HomeView",
    components: {
        TableFood,
    },
}
</script>
