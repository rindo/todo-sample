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
    },
    updateTodo(state, id, name, done) {
      var todo = state.todos.filter(it => it.id == id)
      todo.name = name
      todo.done = done
    },
    removeTodo (state, id) {
      state.todos = state.todos.filter(it => it.id != id)
    }
  },
  actions: {
    getTodos (context) {
      api.todos().then(res => {
        this.commit('setTodos', res.data)
      })
    },
    addTodo (context, name) {
      api.addTodo(name).then(res => {
        this.dispatch("getTodos")
      })
    },
    updateTodo (context, id, name, done) {
      api.updateTodo(id, name, done).then(res => {
        this.commit('updateTodo', id, name, done)
      })
    },
    deleteTodo (context, id) {
      api.deleteTodo(id).then(res => {
        this.commit('removeTodo', id)
      })
    }
  }
})
