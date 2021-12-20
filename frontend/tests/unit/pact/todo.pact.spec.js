const {API} = require("@/api");
const {pactWith} = require("jest-pact");
const interactions = require('./todo.pact.interactions')

pactWith(
    {
        consumer: "TodoFrontend",
        provider: "TodoBackend",
        pactfileWriteMode: 'merge',
    }, provider => {
        let api;

        describe("todo", () => {
            beforeEach(() => {
                api = new API(provider.mockService.baseUrl)
            })
            describe("get todo list", () => {
                it('get todo list successfully', async () => {
                    await provider.addInteraction(interactions.couldGetTodoList())
                    const res = await api.getTodoList()
                    console.log(res)
                })
            })
        })
    })