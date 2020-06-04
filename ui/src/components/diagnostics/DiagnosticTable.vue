<template>
  <div class="dx-table">
    <vue-good-table
      :columns="getColumns()"
      :rows="diagnostics"
      :search-options="{
        enabled: true,
        skipDiacritics: true,
      }"
    >
      <template slot="table-row" slot-scope="props">
        <span v-if="props.column.field == 'name' && props.row.testUrl">
          <a :href="props.row.testUrl">{{ props.row.name }}</a>
        </span>
        <span v-else-if="props.column.field == 'sampleTypes'">
          <div v-for="st in props.row.sampleTypes" :key="st.id">{{ st.name }}</div>
        </span>
        <span v-else-if="props.column.field == 'pcrPlatforms'">
          <div v-for="p in props.row.pcrPlatforms" :key="p.id">{{ p.name }}</div>
        </span>
        <span v-else-if="props.column.field == 'regulatoryApprovals'">
          <div v-for="ra in props.row.regulatoryApprovals" :key="ra.id">{{ ra.name }}</div>
        </span>
        <span v-else>
          {{props.formattedRow[props.column.field]}}
        </span>
      </template>
    </vue-good-table>
  </div>
</template>

<script>
  import { VueGoodTable } from 'vue-good-table';

  export default {
    methods: {
      select(dx) {
        this.$emit('select:poc', dx);
      },
      convertBoolToYN(value) {
        return value ? "Y" : "N"
      },
      filterYN(data, filterString) {
        console.log("fileterString: %s", filterString);
        console.log(data);
        return (data && filterString === 'Y') || (!data && filterString === 'N')
      },
      getColumns(){
        const stl = this.sampleTypeList.length > 0 ? this.sampleTypeList.map((i) => {
          return {'value': i.id, 'text':i.name}
        }) : ['Loading...'];
        const pcr = this.pcrPlatformList.length > 0 ? this.pcrPlatformList.map((i) => {
          return {'value': i.id, 'text':i.name}
        }) : ['Loading...'];
        return [
          { 
            'label': 'Company',
            'field':'company.name',
            'sortable': true,
            'filterOptions':{
              'enabled': true
            }
          },
          { 
            'label': 'Name',
            'field':'name',
            'sortable': true,
            'filterOptions':{
              'enabled': true
            }
          },
          { 
            'label': 'PCR Platform',
            'field':'pcrPlatforms',
            'filterOptions':{
              'enabled': true,
              'placeholder': 'All',
              'filterDropdownItems': pcr,
              'filterFn': (data, filterString) => {
                return data.filter(p => p.id === filterString).length > 0
              }
            }
          },
          { 
            'label': 'Sensitivity',
            'field':'sensitivity',
            'width': '50px',
            'sortable': true,
            'filterOptions':{
              'enabled': true
            }
          },
           { 
            'label': 'Specificity',
            'field':'specificity',
            'width': '50px',
            'sortable': true,
            'filterOptions':{
              'enabled': true
            }
          },
          { 
            'label': 'Regulatory Status',
            'field': 'regulatoryApprovals',
            'sortable': true,
            'filterOptions':{
              'enabled': true,
              'filterFn': (data, filterString) => {
                return data.filter(ra => ra.name.includes(filterString)).length > 0
              }
            }
          },
          { 
            'label': 'Sample Types',
            'field':'sampleTypes',
            'filterOptions':{
              'enabled': true,
              'placeholder': 'All',
              'filterDropdownItems': stl,
              'filterFn': (data, filterString) => {
                return data.filter(st => st.id === filterString).length > 0
              }
            }
          },
          { 
            'label': 'Point of Care',
            'field':'pointOfCare',
            'sortable': true,
            'filterOptions':{
              'enabled': true,
              'placeholder': ' ',
              'filterDropdownItems': ['Y', 'N'],
              'filterFn':  this.filterYN
            },
            'formatFn': this.convertBoolToYN,
            'width': '50px',

          },
          { 
            'label': 'Integrated Sample Prep',
            'field':'prepIntegrated',
            'sortable': true,
            'filterOptions':{
              'enabled': true,
              'placeholder': ' ',
              'filterDropdownItems': ['Y', 'N'],
              'filterFn':  this.filterYN
            },
            'formatFn': this.convertBoolToYN,
            'width': '4em',
          },
        ]
      }
    },
    name: 'DiagnosticTable',
    props: {
      diagnostics: Array,
      selectedDx: Object,
      sampleTypeList: Array,
      pcrPlatformList: Array,
    },
    components: {
      VueGoodTable,
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