<template>
  <div id="app">
    <h1>Todo App</h1>

    <input id="todo-input" v-model="todo"/>
    <button id="addTodo" @click="addTodo">Add</button>
    <button id="deleteAll" @click="deleteAllTodos">Delete All</button>

    <section class="todo-list">
      <p v-for="(todo, index) in todoList" :key="index">{{ index + 1 }}. {{ todo.description }}</p>
    </section>

    <p>{{ error }}</p>
  </div>
</template>

<script>
import API from '@/api'

export default {
  data() {
    return {
      todo: '',
      todoList: [],
      error: ''
    }
  },
  methods: {
    async addTodo() {
      if (this.todo === '')
        return

      try {
        const {data} = await API.addTodo(this.todo)
        this.todoList.push(data)
      } catch (e) {
        console.error(e)
        this.error = e
      } finally {
        this.todo = ''
      }
    },
    async deleteAllTodos() {
      try {
        await API.deleteAllTodos()
        this.todoList = []
      } catch (e) {
        console.error(e)
        this.error = e
      }
    }
  },
  async created() {
    try {
      const {data} = await API.getTodoList()
      this.todoList = data
    } catch (e) {
      console.error(e)
      this.error = e
    }
  }
}
</script>

<style lang="scss">
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
  text-align: center;
}
</style>
