import Vue from 'vue'
import Vuex from 'vuex'
import api from './utils/api.js'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    todos: []
  },
  mutations: {
    setTodos (state, todos) {
      state.todos = todos
    }
  },
  actions: {
    getTodos (context) {
      api.todos().then(res => {
        this.commit('setTodos', res.data.todos)
      })
    },
    addTodo (context, name) {
      // TODO: add to state
      api.addTodo(name).then(res => {
        this.dispatch("getTodos")
      })
    },
    updateTodo (context, id, name, done) {
      api.updateTodo(id, name, done).then(res => {
        // TODO: update from state
        this.dispatch("getTodos")
      })
    },
    deleteTodo (context, id) {
      api.deleteTodo(id).then(res => {
        // TODO: delete from state
        this.dispatch("getTodos")
      })
    }
  }
})
