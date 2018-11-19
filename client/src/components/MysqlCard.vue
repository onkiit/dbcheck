<template>
    <v-card>
        <v-card-title>
            <span class="title">MySQL</span>
        </v-card-title>
        <v-card-text>
            <v-layout row wrap>
                <v-flex xs12>
                    <v-flex xs6>
                        <span class="body-2">Version</span>
                    </v-flex>
                    <v-flex xs6>
                        <span class="caption">{{ monitoringData.version }}</span>
                    </v-flex>
                </v-flex>
                <v-flex xs12>
                    <v-flex xs6>
                        <span class="body-2">Active Clients</span>
                    </v-flex>
                    <v-flex xs6>
                        <span class="caption">{{ monitoringData.active_client }}</span>
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
            monitoringData: {
                version: "",
                active_client: 0
            }
        }
    },
    created(){
        this.monitoringData.version = "PostgreSQL version 3"
    },
    methods: {
        
    }
}
</script>
