<template>
    <v-table fixed-header height="300px">
        <thead>
            <tr>
                <th class="text-left">Name</th>
                <th class="text-left">Calories</th>
            </tr>
        </thead>
        <tbody>
            <tr v-for="item in foods" :key="item.name">
                <td>{{ item.name }}</td>
                <td>{{ item.price }}</td>
            </tr>
        </tbody>
    </v-table>
</template>

<script>
export default {
    data() {
        return {
            foods: [],
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
        this.foods = foods.data.foods
        console.log(this.foods)
    },
}
</script>
