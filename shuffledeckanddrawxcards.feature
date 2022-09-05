Feature: draw from shuffled deck
  Scenario: shuffle a deck and draw x cards
    Given a deck has been created
    When the function is called
    Then it shuffles the deck
    And draw X different cards at random
