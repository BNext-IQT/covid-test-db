<template>
  <div id="app">
    <h1>COVID-19 Testing Database</h1>
    <Diagnostics />
  </div>
</template>

<script>
import Diagnostics from '@/components/diagnostics/Diagnostics.vue'

export default {
  name: 'App',
  components: {
    Diagnostics
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
