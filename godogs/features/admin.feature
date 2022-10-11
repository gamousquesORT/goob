Feature: Create Genres
  In order to add films
  As an Admin user
  I need to be able add gneres

  Scenario: Add valid genre
    Given I am logged as Admin
    When I create new genre "Terror", "Genero para tener miedo"
    Then I should be able to retrieve it getting "Terror", "Genero para tener miedo"

Scenario: Add several Genres
    Given I am logged as Admin
    When I add the following Genres
    | Comedia | para reirse |
    | Suspenso | para pasar un buen rato |
    | Documental | para aprender |
    Then I should have 4 Genres in the app


