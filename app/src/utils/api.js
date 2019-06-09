'use strict';

import axios from 'axios'

const client = axios.create({
  baseURL: 'http://localhost:8080'
})

const onSuccess = res => {
  console.log(`[API Response] ${JSON.stringify(res.data)}`)
}

const onError = res => {
  throw new Error('API Error')
}

export default {
  addTodo: (name) => {
    axios.post('/todos').then(onSuccess).catch(onError)
  },

  todos: () => {
    axios.get('/todos').then(onSuccess).catch(onError)  
  },

  updateTodo: (id, name, done) => {
    axios.put(`/todos/${id}`).then(onSuccess).catch(onError)
  },

  deleteToodo: (id) => {
    axios.delete(`/todos/${id}`).then(onSuccess).catch(onError)
  }

}
