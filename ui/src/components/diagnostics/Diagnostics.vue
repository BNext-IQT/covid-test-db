<template>
  <div class="diagnostics">
      <DiagnosticTable :diagnostics="diagnostics" @select:poc="setSelectedDx" />
     <!--  <hr>
      <PocForm :poc="selectedPoc" @update:poc="updatePoc" @clear:poc="clearSelection" /> -->
  </div>
</template>

<script>
  import axios from "axios";

  import DiagnosticTable from '@/components/diagnostics/DiagnosticTable.vue';
  //import PocForm from '@/components/pocs/PocForm.vue';

  export default {
    data (){
      return {
        diagnostics: null,
        selectedDx: { }
      }
    },
    methods: {
      getData(){
        axios
        .get("http://localhost:5000/diagnostics", {crossdomain: true})
        .then((res) => {
          this.diagnostics = res.data;
        })
        .catch((err) => {
          console.log("error: %o", err);
        })
      },
      // updatePoc(poc) {
      //   if(poc.name && (poc.phone || poc.email)){
      //     const body = poc
      //     if(poc.id == null){
      //       axios
      //       .post("http://localhost:5000/pocs", body)
      //       .then((res) => {
      //         this.selectedPoc = res.data;
      //         this.getData();
      //       })
      //       .catch((err) => {
      //         console.log("error: %o", err);
      //       })
      //     }else if(poc.id != null){
      //       axios
      //       .put("http://localhost:5000/pocs/" + poc.id, body)
      //       .then((res) => {
      //         this.selectedPoc = res.data;
      //         this.getData();
      //       })
      //       .catch((err) => {
      //         console.log("error: %o", err);
      //       })
      //     }
      //   }
      // },
      setSelectedDx(dx) {
        this.selectedDx = dx;
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
      //PocForm,
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