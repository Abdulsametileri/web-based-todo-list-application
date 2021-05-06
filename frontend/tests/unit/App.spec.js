import {shallowMount} from "@vue/test-utils";
import flushPromises from "flush-promises";
import App from "@/App";

jest.mock("axios", () => ({
    get: () => Promise.resolve({ data: { data: [] } }),
    post: () => Promise.resolve({ data: { data: { id: 0, description: 'dummy todo' }} })
}));

describe('App.vue', () => {
    it('should add a todo when user enters input and clicks the button', async () => {
        const inputToAdd = "dummy todo"

        const wrapper = shallowMount(App)
        await flushPromises();

        const input = wrapper.find('input')
        await input.setValue(inputToAdd)

        const button = wrapper.find('button')
        await button.trigger('click')
        await flushPromises();

        expect(wrapper.vm.todo).toBe('')
        expect(wrapper.vm.todoList[0].id).toEqual(0)
        expect(wrapper.vm.todoList[0].description).toBe(inputToAdd)
        expect(wrapper.find('p').text()).toBe('1. ' + inputToAdd)
    })
})