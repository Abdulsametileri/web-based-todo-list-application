import {shallowMount} from "@vue/test-utils";
import flushPromises from "flush-promises";
import App from "@/App";
import API from '@/api'

jest.mock("@/api")

describe('App.vue', () => {
    it('should add a todo when user enters input and clicks the button', async () => {
        const inputTodo = {
            id: 1,
            description: "dummy todo"
        }

        API.getTodoList.mockResolvedValue({data: []})
        API.addTodo.mockResolvedValue({ data: inputTodo })

        const wrapper = shallowMount(App)
        await flushPromises();

        const input = wrapper.find('input')
        await input.setValue(inputTodo.description)

        const button = wrapper.find('#addTodo')
        await button.trigger('click')
        await flushPromises();

        expect(wrapper.vm.todo).toBe('')
        expect(wrapper.vm.todoList[0].id).toEqual(inputTodo.id)
        expect(wrapper.vm.todoList[0].description).toBe(inputTodo.description)
        expect(wrapper.find('p').text()).toBe('1. ' + inputTodo.description)
    })
    it('delete all button exist', async () => {
        const wrapper = shallowMount(App)

        const deleteAllButton = wrapper.find('#deleteAll')

        expect(deleteAllButton.text()).toContain("Delete All")
    })
    it('delete all click propogates', async () => {
        const deleteAllTodos = jest.fn()
        const wrapper = shallowMount(App, {
            methods: {
                deleteAllTodos
            }
        })

        const deleteAllButton = wrapper.find('#deleteAll')

        await deleteAllButton.trigger('click')
        expect(deleteAllTodos).toBeCalled()
    })
})