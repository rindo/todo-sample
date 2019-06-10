'use strict';

import axios from 'axios'

axios.defaults.headers.post["Content-Type"] = "application/x-www-form-urlencoded";

const client = axios.create({
  baseURL: 'api:8080'
})

export default {
  addTodo: (name) => {
    return client.post('/todos', { name })
  },

  todos: () => {
    return client.get('/todos')
  },

  updateTodo: (id, name, done) => {
    return client.put(`/todos/${id}`)
  },

  deleteToodo: (id) => {
    return client.delete(`/todos/${id}`)
  }

}
