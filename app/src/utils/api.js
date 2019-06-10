'use strict';

import axios from 'axios'
axios.defaults.headers.post["Content-Type"] = "application/json";

const client = axios.create({
  baseURL: 'http://localhost:8080'
})

export default {
  addTodo: (name) => {
    return client.post('/todos', { name: name })
  },

  todos: () => {
    return client.get('/todos')
  },

  updateTodo: (id, name, done) => {
    return client.put(`/todos/${id}`, { name: name, done: done })
  },

  deleteToodo: (id) => {
    return client.delete(`/todos/${id}`)
  }
}
