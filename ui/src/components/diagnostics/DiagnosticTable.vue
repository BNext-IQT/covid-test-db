<template>
  <div class="dx-table">
    <table class="ui celled table">
      <thead>
        <tr>
          <th>Company</th>
          <th>Name</th>
          <th>Type</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="dx in diagnostics" :key="dx.id" @click="select(dx)">
          <td>{{ dx.company.name }}</td>
          <td>{{ dx.name }}</td>
          <td>{{ dx.diagnosticType.name }}</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
  export default {
    methods: {
      select(dx) {
        console.log("selected: %s", dx.id);
        this.$emit('select:poc', dx);
      },
    },
    name: 'DiagnosticTable',
    props: {
      diagnostics: Array
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