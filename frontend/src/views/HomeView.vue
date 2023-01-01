<template>
    <v-container>
        <TableFood :items="foods_general" :heads="head_general" />
    </v-container>
</template>

<script>
import TableFood from "../components/TableFood.vue"

export default {
    name: "HomeView",
    components: {
        TableFood,
    },
    data() {
        return {
            foods_all: [],
            foods_general: [],
            head_general: ["Name", "Amount", "Calories", "Days left"],
            food_general_format: JSON.parse(
                JSON.stringify({
                    name: "",
                    amount: "",
                    calories: "",
                    description: "",
                })
            ),
        }
    },
    methods: {
        async getFoods() {
            const response = await fetch("api/food/get_available")
            const data = await response.json()
            return data
        },
    },
    async mounted() {
        const foods = await this.getFoods()
        this.foods_all = foods.data.foods
        console.log(this.foods_all)

        this.foods_all.map((item) => {
            let item_general = {}
            for (const key in this.food_general_format) {
                item_general[key] = item[key]
            }
            this.foods_general.push(item_general)
        })
    },
}
</script>
