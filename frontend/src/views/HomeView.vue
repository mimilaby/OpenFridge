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
            head_general: [],
        }
    },
    methods: {
        async getFoods() {
            const response = await fetch("api/food/get_available")
            const data = await response.json()
            return data
        },
        selectGeneral() {
            let food = {}
            this.head_general = ["Name", "Amount", "Calories", "Days left"]
            let food_general =
                '{"name": "","amount": "","calories": "","description": ""}'
            food_general = JSON.parse(food_general)

            for (const indx in this.foods_all) {
                food = this.foods_all[indx]
                for (const key in food_general) {
                    food_general[key] = food[key]
                }
                console.log(food_general)
                this.foods_general.append(food_general)
                console.log(this.foods_general)
            }
        },
    },
    async mounted() {
        const foods = await this.getFoods()
        this.foods_all = foods.data.foods
        console.log(this.foods_all)
        this.selectGeneral()
        // console.log(this.foods_general)
    },
}
</script>
