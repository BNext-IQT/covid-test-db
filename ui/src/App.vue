<template>
  <div id="app">
    <div class="logos">
      <img src="./assets/b.next logo.png" height="72px" width="280px">
      <img src="./assets/IQT_Logo.png" height="72px" width="200px">
    </div>
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

h1 {
  color: #213964;
}

.logos {
  width:100%;
  display: flex;
  flex-flow: row;
  padding: 0.5em;
  justify-content: flex-end;
  align-content: flex-end;
  align-self: flex-end;
}
</style>
