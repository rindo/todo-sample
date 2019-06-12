<template lang="pug">
  #app
    b-table(:items="todos" :fields="fields" striped)
      //- template(slot="name" slot-scope="row")
      //-   p(v-if="" @cliick="")
      //-   b-form-input(v-model="text" placeholder="Enter your name")

      template(slot="delete" slot-scope="row")
        b-button(size="sm" @click="deleteTodo(row.item)") Delete

    b-form(inline)
      b-input#field(placeholder="What's yuo want to do?" v-model="input")
      b-btn(variant="primary" @click="addTodo(input)" :disabled="!canAdd") Add
</template>

<script>
export default {
  name: 'app',
  data () {
    return {
      fields: ['id', 'name', 'created_at', 'updated_at', 'delete'],
      input: "",
      editing: ""
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
    addTodo: function (name) {
      this.$store.dispatch('addTodo', name)
    },
    deleteTodo: function (item) {
      this.$store.dispatch('deleteTodo', item.Id)
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
