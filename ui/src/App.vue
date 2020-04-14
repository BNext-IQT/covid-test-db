<template>
  <div id="app">
    <h1>COVID-19 Testing Database</h1>
    <Pocs />
  </div>
</template>

<script>
import Pocs from '@/components/pocs/Pocs.vue'

export default {
  name: 'App',
  components: {
    Pocs
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
