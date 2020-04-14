<template>
  <div class="poc-form">
    <sui-form @submit.prevent="update">
      <sui-form-field inline>
        <input v-model="poc.name" type="text" placeholder="Name" />
        <!-- <sui-label basic color="red" pointing>Name cannot be blank.</sui-label> -->
      </sui-form-field>
      <sui-form-field inline>
        <input v-model="poc.phone" type="text" placeholder="Phone" />
      </sui-form-field>
      <sui-form-field inline>
        <input v-model="poc.email" type="text" placeholder="Email" />
      </sui-form-field>
      <button class="ui button">{{buttonText}}</button>
      <button class="ui button" @click="clear">Clear</button>
    </sui-form>
  </div>
</template>

<script>
  export default {
    methods: {
      update() {
        this.$emit('update:poc', this.poc)
      },
      clear(){
        this.$emit('clear:poc', this.poc)
      },
    },
    computed: {
      buttonText() {
        return this.poc.id != null ? "Update PoC" : "Add PoC"
      },
    },
    name: 'PocTable',
    props: {
      poc: Object
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