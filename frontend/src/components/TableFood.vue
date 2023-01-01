<template>
    <v-hover v-slot="{isHovering, props}">
        <v-card
            class="rounded-xl"
            v-bind="props"
            :elevation="isHovering ? 24 : 6"
            color="accent"
        >
            <v-card-title class="text-center" color="#AAA">
                Available foods
            </v-card-title>
            <v-table fixed-header height="300px" class="rounded-0">
                <thead>
                    <tr>
                        <th class="text-left" v-for="head in heads" :key="head">
                            {{ head }}
                        </th>
                    </tr>
                </thead>
                <tbody>
                    <tr v-for="item in items" :key="item.name">
                        <td v-for="text in item" :key="text">{{ text }}</td>
                    </tr>
                </tbody>
            </v-table>
        </v-card>
    </v-hover>
</template>

<script>
export default {
    props: ["items", "heads"],
    created() {
        // console.log(this.heads)
        // console.log(this.items)
    },
}
</script>
