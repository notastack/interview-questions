Feature: rps
  Scenario: play a game of rock paper scissors
    Given the user choose his play
    When the function is called
    Then generate the computer's play
    And prints out the results of the game
