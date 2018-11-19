<template>
  <v-container fluid grid-list-xs>
    <v-layout row wrap>
      <v-flex xs6 xl6>
        <psql-card/>
      </v-flex>
      <v-flex xs6 xl6>
        <mysql-card />
      </v-flex>
      <v-flex xs6 xl6 class="mt-2">
        <mongo-card/>
      </v-flex>
      <v-flex xs6 xl6 class="mt-2">
        <redis-card />
      </v-flex>
    </v-layout>
  </v-container>
</template>

<script>
  import PsqlCard from '../components/PsqlCard'
  import MysqlCard from '../components/MysqlCard'
  import MongoCard from '../components/MongoCard'
  import RedisCard from '../components/RedisCard'

  export default {
    components: {
      PsqlCard,
      MysqlCard,
      MongoCard,
      RedisCard
    }
  }
</script>
