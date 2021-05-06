<template>
  <div id="app">
    <h1>Todo App</h1>

    <input id="todo-input" v-model="todo"/>
    <button @click="addTodo">Add</button>

    <section class="todo-list">
      <p v-for="(todo, index) in todoList" :key="index">{{ index + 1 }}. {{ todo }}</p>
    </section>
  </div>
</template>

<script>
export default {
  data() {
    return {
      todo: '',
      todoList: [],
    }
  },
  methods: {
    addTodo() {
      if (this.todo === '')
        return

      this.todoList.push(this.todo)

      this.todo = ''
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
