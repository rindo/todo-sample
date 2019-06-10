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
        commit('setTodos', res.data)
      })
    },
    addTodo (context, name) {
      api.addTodo(name).then(res => {
        getTodos()
      })
    },
    deleteTodo () {
      
    }
  }
})
