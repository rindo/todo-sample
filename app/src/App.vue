<template lang="pug">
  #app
    b-table(striped :items="todos")

    b-form(inline)
      b-input#field(placeholder="What's yuo want to do?" v-model="input")
      b-btn(variant="primary" @click="addTodo(input)" :disabled="!canAdd") Add
</template>

<script>
export default {
  name: 'app',
  data () {
    return {
      input: ""
    }
  },
  created () {
    this.$store.dispatch('getTodos')
  },
  computed: {
    todos() {
      return this.$store.state.todos
    },
    canAdd() {
      return this.input != ""
    }
  },
  methods: {
    addTodo: function () {
      this.$store.dispatch('addTodo', this.input)
    }
  }
}
</script>

<style lang="scss">
#app {
  margin: 10pt;

  #field {
    width: 50%;
  }
}
</style>
