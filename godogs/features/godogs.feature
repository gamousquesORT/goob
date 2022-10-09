Feature: Create Genres
  In order to add films
  As a Admin user
  I need to be able add gneres

  Scenario: Add valid genre
    Given I am logged as Admin
    When I create new genre
    Then I should see "Genre was successfully created."
