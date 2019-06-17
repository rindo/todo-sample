<template lang="pug">
  #app
    b-table(:items="todos" :fields="fields" striped)
      template(slot="delete" slot-scope="row")
        b-button(size="sm" variant="danger" @click="deleteTodo(row.item)") Delete

    b-form(inline)
      b-input#field(placeholder="What's yuo want to do?" v-model="input")
      b-btn(variant="primary" @click="addTodo()" :disabled="!canAdd") Add
</template>

<script>
export default {
  name: 'app',
  data () {
    return {
      fields: ['id', 'name', 'created_at', 'updated_at', 'delete'],
      input: "",
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
      this.input = ""
    },
    deleteTodo: function (item) {
      console.log(item)
      this.$store.dispatch('deleteTodo', item.id)
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
