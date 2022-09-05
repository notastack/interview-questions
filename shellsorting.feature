Feature: shell sorting
  Scenario: sorts an array
    Given a random int array has been created
    When the function is called
    Then it sorts the array from lower to higher
    And prints the sorted array
