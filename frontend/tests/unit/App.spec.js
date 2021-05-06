import {shallowMount} from "@vue/test-utils";
import App from "@/App";

describe('App.vue', () => {
    it('add a todo when clicking the button', async () => {
        const wrapper = shallowMount(App)

        const inputToAdd = "dummy todo"

        const input = wrapper.find('input')
        await input.setValue(inputToAdd)

        const button = wrapper.find('button')
        await button.trigger('click')

        expect(wrapper.vm.todoList[0]).toBe(inputToAdd)
        expect(wrapper.vm.todo).toBe('')
        expect(wrapper.find('p').text()).toBe('1. ' + inputToAdd)
    })
})