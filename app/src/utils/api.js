'use strict';

import axios from 'axios'

const client = axios.create({
  baseURL: 'api:8080'
})

export default {
  addTodo: (name) => {
    client.post('/todos', { name })
  },

  todos: () => {
    client.get('/todos')
  },

  updateTodo: (id, name, done) => {
    client.put(`/todos/${id}`)
  },

  deleteToodo: (id) => {
    client.delete(`/todos/${id}`)
  }

}
