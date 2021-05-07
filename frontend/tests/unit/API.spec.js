import API from '@/api'
import nock from 'nock'

describe("API", () => {
    it("get all todos", async () => {
        const todos = [
            {id: 0, description: 'Todo 1'},
            {id: 1, description: 'Todo 2'},
            {id: 2, description: 'Todo 3'},
        ]
        nock(API.url)
            .get('/getTodoList')
            .reply(200,
                todos)
        const respTodos = await API.getTodoList()
        expect(respTodos).toEqual(todos)
    })

    it("add todo", async () => {
        const addedTodo = { id: 0, description: 'Dummy todo' }

        nock(API.url)
            .post('/addTodo', { task_description: 'Dummy todo' })
            .reply(201, addedTodo)

        const respTodo = await API.addTodo(addedTodo.description)
        expect(respTodo).toEqual(addedTodo)
    })
})