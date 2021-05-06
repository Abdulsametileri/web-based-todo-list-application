// https://docs.cypress.io/api/introduction/api.html

describe('My First Test', () => {
  it('Visits the app root url', () => {
    cy.visit('/')
    cy.contains('h1', 'Todo App')
  })
  it('Type todo and click to add new todo', () => {
    cy.get('input').type('Dummy todo')
    cy.get('button').click()
    cy.get('p').contains('1. Dummy todo')
    cy.get('input').empty
  })
})
