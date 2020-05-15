<template>
  <div class="dx-table">
    <table class="ui celled table">
      <thead>
        <tr>
          <th>Company</th>
          <th>Name</th>
          <th>PoC</th>
          <th>Type</th>
          <th>Regulatory Status</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="dx in diagnostics" :key="dx.id" @click="select(dx)">
          <td>
            <div v-if="dx.id !== selectedDx.id">
                <strong>{{ dx.company.name }}</strong><br/>
                {{ dx.company.state }} {{dx.company.state !== "" ? ", " : ""}}{{ dx.company.country }}
            </div>
            <div v-else>
                <strong>{{ dx.company.name }}</strong><br/>
                <div v-if="dx.company.streetAddress !== ''">{{dx.company.streetAddress}}</div>
                {{ dx.company.city }} {{dx.company.city !== "" ? ", " : ""}}{{ dx.company.state }}&nbsp;{{ dx.company.postalCode }}<br v-if="dx.company.city || dx.company.state || dx.company.postalCode" />
                <div v-if="dx.company.country !== ''">{{dx.company.country}}</div>
            </div>
          </td>
          <td>
            <div v-if="dx.id !== selectedDx.id">
                <strong>{{ dx.name }}</strong>
            </div>
            <div v-else>
                <strong>{{ dx.name }}</strong><br/>
                <table>
                  <tr>
                    <td><strong>Verified LOD</strong></td><td>{{ dx.verifiedLod }}</td>
                    <td><strong>Average Ct</strong></td><td>{{ dx.avgCt != 0 ? dx.avgCt : "-" }}</td>
                  </tr>
                  <tr>
                    <td colspan="2"><strong>Sample Types</strong></td>
                    <td colspan="2"><strong>Gene Targets</strong></td>
                  </tr>
                  <tr>
                    <td colspan="2">
                      <div v-if="!dx.sampleTypes || dx.sampleTypes.length === 0"> -No Data-</div>
                      <div v-for="st in dx.sampleTypes" :key="st.id">{{ st.name }}</div>
                    </td>
                    <td colspan="2">
                      <div v-if="!dx.diagnosticTargets || dx.diagnosticTargets.length === 0"> -No Data-</div>
                      <div v-for="dt in dx.diagnosticTargets" :key="dt.id">{{ dt.name }}</div>
                    </td>
                  </tr>
                </table>
            </div>
            
          </td>
          <td>
            <strong>{{ dx.poc.name }}</strong><br/>
            {{ dx.poc.phone }} <br v-if="dx.poc.phone !== ''"/>
            <a v-if="dx.poc.email.includes('@')" :href="'mailto:' + dx.poc.email">{{ dx.poc.email }}</a>
            <a v-else-if="dx.poc.email.includes('http:')" :href="dx.poc.email">{{ dx.poc.email }}</a>
            <span v-else> {{ dx.poc.email }} </span>
            
          </td>
          <td>{{ dx.diagnosticType.name }}</td>
          <td>
              <div v-for="ra in dx.regulatoryApprovals" :key="ra.id">
                {{ ra.name }}
              </div>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
  export default {
    methods: {
      select(dx) {
        this.$emit('select:poc', dx);
      },
    },
    name: 'DiagnosticTable',
    props: {
      diagnostics: Array,
      selectedDx: Object,
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