<template>
    <v-card>
        <v-card-title>
            <span class="title">PostgreSQL</span>
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
import axios from 'axios'
export default {
    data(){
        return{
            monitoringData: {
                version: "",
                active_client: 0,
                health: {
                    usage: []
                }
            }
        }
    },
    created(){
        this.getData()
    },
    methods: {
        getData(){
            this.$http.get("http://localhost:8180/getpsql")
            .then(resp => {
                console.log(resp)
            })
        }
    }
}
</script>
