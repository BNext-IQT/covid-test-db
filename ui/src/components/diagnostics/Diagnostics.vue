<template>
  <div class="diagnostics">
      <DiagnosticTable :diagnostics="diagnostics" :diagnosticTypeList="diagnosticTypeList" :sampleTypeList="sampleTypeList" :pcrPlatformList="pcrPlatformList" :selectedDx="selectedDx" @select:dx="setSelectedDx" />
     <DiagnosticDetail :diagnostic="selectedDx"/>
  </div>
</template>

<script>
  import axios from "axios";

  import DiagnosticTable from '@/components/diagnostics/DiagnosticTable.vue';
  import DiagnosticDetail from '@/components/diagnostics/DiagnosticDetail.vue';

  export default {
    data (){
      return {
        diagnostics: [],
        selectedDx: { },
        sampleTypeList: [],
        pcrPlatformList: [],
        diagnosticTypeList: [],
      }
    },
    methods: {
      getData(){
        axios
        .get("/api/diagnostics")
        .then((res) => {
          this.diagnostics = res.data;

          for(const dx of this.diagnostics){
            const url = new URL('performance/page_results', window.location);
            url.searchParams.append('company', dx.company.name);
            url.searchParams.append('test', dx.name);
            dx['performanceUrl'] = url
          }
        })
        .catch((err) => {
          console.log("error: %o", err);
        })

        axios
        .get("/api/diagnostictypes")
        .then((res) => {
          this.diagnosticTypeList = JSON.parse(JSON.stringify(res.data));
        })
        .catch((err) => {
          console.log("error: %o", err);
        })

        axios
        .get("/api/sampletypes")
        .then((res) => {
          this.sampleTypeList = JSON.parse(JSON.stringify(res.data));
        })
        .catch((err) => {
          console.log("error: %o", err);
        })

        axios
        .get("/api/pcrplatforms")
        .then((res) => {
          this.pcrPlatformList = JSON.parse(JSON.stringify(res.data));
        })
        .catch((err) => {
          console.log("error: %o", err);
        })
      },
      setSelectedDx(dx) {
        this.selectedDx = dx;
        console.log("set selected called with dx = %o", dx);
        this.$modal.show('diagnostic-detail', {'diagnostic': dx})
      },
      clearSelection() {
        this.selectedDx = {};
      }
    },
    mounted (){
      this.getData();
    },
    name: 'Diagnostics',
    components: {
      DiagnosticTable,
      DiagnosticDetail
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