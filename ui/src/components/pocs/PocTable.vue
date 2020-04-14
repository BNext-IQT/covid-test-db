<template>
  <div class="poc-table">
    <table class="ui celled table">
      <thead>
        <tr>
          <th>Name</th>
          <th>Phone</th>
          <th>Email</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="poc in pocs" :key="poc.id" @click="select(poc)">
          <td>{{ poc.name }}</td>
          <td>{{ poc.phone }}</td>
          <td>{{ poc.email }}</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
  export default {
    methods: {
      select(poc) {
        console.log("selected: %s", poc.id);
        this.$emit('select:poc', poc);
      },
    },
    name: 'PocTable',
    props: {
      pocs: Array
    }
  }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>