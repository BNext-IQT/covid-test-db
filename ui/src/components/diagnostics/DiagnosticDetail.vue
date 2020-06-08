<template>
  <modal 
        name="diagnostic-detail"
        :draggable="true"
        :adaptive="true"
        :scrollable="true"
        :reset="true"
        width="60%"
        height="auto">
    <div class="content">
      <div class="left-col">
        <div>
            <strong>Name: </strong> {{this.diagnostic.name}}
        </div>
        <div>
            <strong>Product No: </strong> {{this.diagnostic.catalogNo}}
        </div>
        <div>
            <strong>Sensitivity: </strong> {{this.diagnostic.sensitivity}}
        </div>
        <div>
            <strong>Specificity: </strong> {{this.diagnostic.specificity}}
        </div>
        <div>
            <strong>Verified LoD: </strong> {{this.diagnostic.verifiedLod}}
        </div>
        <div>
            <strong>Average CT: </strong> {{this.diagnostic.averageCt}}
        </div>
        <div>
            <strong>Tests per Run: </strong> {{this.diagnostic.testsPerRun}}
        </div>
        <div>
            <strong>Tests per Kit: </strong> {{this.diagnostic.testsPerKit}}
        </div>
        <div>
            <strong>Point of Care: </strong> {{this.convertBoolToYN(this.diagnostic.pointOfCare)}}
        </div>
        <div>
            <strong>Integrated Sample Prep: </strong> {{this.convertBoolToYN(this.diagnostic.prepIntegrated)}}
        </div>
        <div>
            <strong>Cost per kit: </strong> {{this.diagnostic.costPerKit}}
        </div>
        <div>
            <strong>In Stock: </strong> {{this.convertBoolToYN(this.diagnostic.inStock)}}
        </div>
        <div>
            <strong>Lead Time: </strong> {{this.diagnostic.costPerKit}}
        </div>
        <div v-if="this.diagnostic.regulatoryApprovals !== null" class="regulatory-data">
            <strong>Regulatory Status </strong>
            <div v-for="ra in this.diagnostic.regulatoryApprovals" :key="ra.id">{{ ra.name }}</div>
        </div>
        <div v-if="this.diagnostic.pcrPlatforms !== null" class="platform-data">
            <strong>PCR Platforms </strong>
            <div v-for="pp in this.diagnostic.pcrPlatforms" :key="pp.id">{{ pp.name }}</div>
        </div>
      </div>
      <div class="right-col">
        <div v-if="this.diagnostic.company" class="company-data">
            <strong>Company Data </strong>
            <div>
                {{this.diagnostic.company.name}}
            </div>
            <div>
                {{this.diagnostic.company.streetAddress}}
            </div>
            <div>
                {{this.diagnostic.company.city}}, {{this.diagnostic.company.state}} {{this.diagnostic.company.country}}
            </div>
        </div>
        <div v-if="this.diagnostic.poc" class="contact-data">
            <strong>Contact </strong>
            <div>
                {{this.diagnostic.poc.name}}
            </div>
            <div>
                {{this.diagnostic.poc.phone}}
            </div>
            <div>
                {{this.diagnostic.poc.email}}
            </div>
        </div>
        <div v-if="this.diagnostic.sampleTypes !== null" class="sample-data">
            <strong>Sample Types </strong>
            <div v-for="st in this.diagnostic.sampleTypes" :key="st.id">{{ st.name }}</div>
        </div>
      </div>
    </div>
  </modal>
</template>

<script>
  export default {
    methods: {
      convertBoolToYN(value) {
        return value ? "Y" : "N"
      },
    },
    name: 'DiagnosticDetail',
    props: {
      diagnostic: Object
    }
  }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.content {
    display: flex;
    flex-flow: row;
    padding: 0.5em;
}
.left-col {
    flex: 1;
    border: 2px solid black;
    flex-flow: column;
    padding: 0.5em;
    margin: 0.5em;
    justify-content: flex-start;
    align-content: flex-start;
    align-self: flex-start;
    text-align: left;
}
.right-col {
    flex: 1;
    border: 2px solid black;
    flex-flow: column;
    padding: 0.5em;
    margin: 0.5em;
    justify-content: flex-start;
    align-content: flex-start;
    align-self: flex-start;
    text-align: left;
}
.company-data {
    flex: 1;
    border: 2px solid navy;
    flex-flow: column;
    padding: 0.5em;
    justify-content: flex-start;
    align-content: flex-start;
    align-self: flex-start;

}
.poc-data {
    flex: 1;
    border: 2px solid navy;
    flex-flow: column;
    padding: 0.5em;
    justify-content: flex-start;
    align-content: flex-start;
    align-self: flex-start;
}
.regulatory-data {
    flex: 1;
    border: 2px solid navy;
    flex-flow: column;
    padding: 0.5em;
    justify-content: flex-start;
    align-content: flex-start;
    align-self: flex-start;

}
.sample-data {
    flex: 1;
    border: 2px solid navy;
    flex-flow: column;
    padding: 0.5em;
    justify-content: flex-start;
    align-content: flex-start;
    align-self: flex-start;

}
.platform-data {
    flex: 1;
    border: 2px solid navy;
    flex-flow: column;
    padding: 0.5em;
    justify-content: flex-start;
    align-content: flex-start;
    align-self: flex-start;

}
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