<template lang="pug">
  #app
    b-table(striped :items="todos")

    b-form(inline)
      b-input#field(placeholder="What's yuo want to do?" v-model="input")
      b-btn(variant="primary" @click="addTodo(input)" :disabled="!canAdd") Add
      p {{canAdd}}
</template>

<script>
export default {
  name: 'app',
  data: function () {
    return {
      todos: [
        {id: 1, name: 'todo1', done: true, created_at: 1, updated_at: 1}
      ],
      input: ""
    }
  },
  computed: {
    canAdd: function () {
      return this.input != ""
    }
  },
  methods: {
    addTodo: function () {
      // this.todos.push({
      //   id: 1, name: 'todo1', done: true, created_at: 1, updated_at: 1
      // })
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
