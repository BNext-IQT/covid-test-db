<template>
  <div class="pocs">
      <PocTable :pocs="pocs" @select:poc="setSelectedPoc" />
      <hr>
      <PocForm :poc="selectedPoc" @update:poc="updatePoc" @clear:poc="clearSelection" />
  </div>
</template>

<script>
  import axios from "axios";

  import PocTable from '@/components/pocs/PocTable.vue';
  import PocForm from '@/components/pocs/PocForm.vue';

  export default {
    data (){
      return {
        pocs: null,
        selectedPoc: { 'id':null, 'name': '', 'phone':'', 'email':''}
      }
    },
    methods: {
      getData(){
        axios
        .get("http://localhost:5000/pocs")
        .then((res) => {
          this.pocs = res.data;
        })
        .catch((err) => {
          console.log("error: %o", err);
        })
      },
      updatePoc(poc) {
        if(poc.name && (poc.phone || poc.email)){
          const body = poc
          if(poc.id == null){
            axios
            .post("http://localhost:5000/pocs", body)
            .then((res) => {
              this.selectedPoc = res.data;
              this.getData();
            })
            .catch((err) => {
              console.log("error: %o", err);
            })
          }else if(poc.id != null){
            axios
            .put("http://localhost:5000/pocs/" + poc.id, body)
            .then((res) => {
              this.selectedPoc = res.data;
              this.getData();
            })
            .catch((err) => {
              console.log("error: %o", err);
            })
          }
        }
      },
      setSelectedPoc(poc) {
        this.selectedPoc = poc;
      },
      clearSelection() {
        this.selectedPoc = { 'id':null, 'name': '', 'phone':'', 'email':''};
      }
    },
    mounted (){
      this.getData();
    },
    name: 'Pocs',
    components: {
      PocTable,
      PocForm,
    }
    // props: {
    //   msg: String
    // }
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