<template>
    <v-card>
        <v-card-title>
            <span class="title">Mongo</span>
        </v-card-title>
        <v-card-text>
            <v-layout row wrap>
                <v-flex xs12>
                    <v-flex xs6>
                        Version
                    </v-flex>
                    <v-flex xs6>
                        {{ monitoringData.version }}
                    </v-flex>
                </v-flex>
            </v-layout>
        </v-card-text>
    </v-card>
</template>

<script>
export default {
    data(){
        return{
            monitoringData: {}
        }
    },
    created(){
        this.monitoringData.version = "PostgreSQL version 3"
    },
    methods: {
        
    }
}
</script>
