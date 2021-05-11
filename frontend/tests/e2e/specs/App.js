// https://docs.cypress.io/api/introduction/api.html

describe('My First Test', () => {
  it('Visits the app root url', () => {
    cy.visit('/')
    cy.contains('h1', 'Todo App')
  })
  it('Type todo and click to add new todo', () => {
    cy.get('input').type('testtset')
    cy.get('#addTodo').click()
    cy.get('.todo-list').contains('testtset')
    cy.get('input').empty
  })
  it('Given I have multiple inputs, When I click delete all button, ' +
      'then I should I see empty list?', () => {
    cy.get('input').type('item1')
    cy.get('#addTodo').click()
    cy.get('input').type('item2')
    cy.get('#addTodo').click()
    cy.get('#deleteAll').click()
    cy.get('.todo-list').find('p').should('have.length', 0)
  })
})
