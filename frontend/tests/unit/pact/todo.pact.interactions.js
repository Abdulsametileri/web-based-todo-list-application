const {Matchers} = require("@pact-foundation/pact");
const {eachLike, integer, like} = Matchers

const getAllColorsRequest = () => ({
    method: 'GET',
    path: '/getTodoList'
})

const expectedGetTodoListResponse = {
    status: 200,
    headers: {
        'Content-Type': 'application/json; charset=UTF-8'
    },
    body: {
        code: integer(),
        data: eachLike({
            id: integer(),
            description: like("")
        }),
        message: like("")
    }
}

module.exports = {
    couldGetTodoList: () => ({
        state: 'i could get todo list',
        uponReceiving: "a request for getting all todo list",
        withRequest: getAllColorsRequest(),
        willRespondWith: expectedGetTodoListResponse
    })
}