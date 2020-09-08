<template>
  <div class="dx-table">
    <vue-good-table
      :columns="getColumns()"
      :rows="diagnostics"
      :fixed-header="true"
      styleClass="vgt-table bn"
      @on-row-click="select"
    >
      <template slot="table-row" slot-scope="props">
        <span v-if="props.column.field == 'name' && props.row.testUrl">
          <a :href="props.row.testUrl">{{ props.row.name }}</a>
        </span>
        <span v-else-if="props.column.field == 'diagnosticType'">
          <div>{{ props.row.diagnosticType.name }}</div>
        </span>
        <span v-else-if="props.column.field == 'sampleTypes'">
          <div v-for="st in props.row.sampleTypes" :key="st.id">{{ st.name }}</div>
        </span>
        <span v-else-if="props.column.field == 'pcrPlatforms'">
          <div v-for="p in props.row.pcrPlatforms" :key="p.id">{{ p.name }}</div>
        </span>
        <span v-else-if="props.column.field == 'regulatoryApprovals'">
          <div v-for="ra in props.row.regulatoryApprovals" :key="ra.id">
            <div v-if="props.row.sourceOfPerfData">
              <a :href="props.row.sourceOfPerfData">IFU/EUA</a>
            </div>
            <div v-else>
              IFU/EUA
            </div>
          </div>
        </span>
        <span v-else-if="props.column.field == 'performanceUrl'">
          <a :href="props.row.performanceUrl">Performance</a>
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
  // import the styles
  import 'vue-good-table/dist/vue-good-table.css'

  export default {
    methods: {
      select(params) {
        const dx = params.row
        this.$emit('select:dx', dx);
      },
      convertBoolToYN(value) {
        return value ? "Y" : "N"
      },
      filterYN(data, filterString) {
        return (data && filterString === 'Y') || (!data && filterString === 'N')
      },
      getColumns(){
        const stl = this.sampleTypeList.length > 0 ? this.sampleTypeList.map((i) => {
          return i.name
        }).sort() : ['Loading...'];
        const pcr = this.pcrPlatformList.length > 0 ? this.pcrPlatformList.map((i) => {
          return i.name
        }).sort() : ['Loading...'];
        const dt = this.diagnosticTypeList.length > 0 ? this.diagnosticTypeList.map((i) => {
          return i.name
        }).sort() : ['Loading...'];
        return [
          { 
            'label': 'Company',
            'field':'company.name',
            'sortable': true,
            'thClass':'bn',
            'filterOptions':{
              'enabled': true
            }
          },
          { 
            'label': 'Test Name',
            'field':'name',
            'sortable': true,
            'filterOptions':{
              'enabled': true
            }
          },
          { 
            'label': 'Type',
            'field': 'diagnosticType',
            'sortable': true,
            'filterOptions':{
              'enabled': true,
              'placeholder': 'All',
              'filterDropdownItems': dt,
              'filterFn': (data, filterString) => {
                return data.name === filterString;
              }
            }
          },
          { 
            'label': 'Instrument/Platform',
            'field':'pcrPlatforms',
            'filterOptions':{
              'enabled': true,
              'placeholder': 'All',
              'filterDropdownItems': pcr,
              'filterFn': (data, filterString) => {
                return data.filter(p => p.name === filterString).length > 0
              }
            }
          },
          { 
            'label': 'IFU/EUA',
            'field': 'regulatoryApprovals',
            'sortable': true,
            'filterOptions':{
              'enabled': false,
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
                return data.filter(st => st.name === filterString).length > 0
              }
            }
          },
          { 
            'label': 'Point-of-Care',
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
          { 
            'label': 'Performance Data',
            'field':'performanceUrl',
            'sortable': false,
            'filterOptions':{
              'enabled': false,
            },
            'width': '4em',
          },
        ]
      }
    },
    name: 'DiagnosticTable',
    props: {
      diagnostics: Array,
      selectedDx: Object,
      diagnosticTypeList: Array,
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
<style lang="css" scoped>
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
  color: #8bc53f;
}
.vgt-left-align .bn .sortable{
  background-color: #213964 !important;
  background: none !important;
}
</style>