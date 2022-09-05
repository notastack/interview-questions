Feature: countdown
  Scenario: count down by decrements of Z
    Given two limits and a Z decrement
    When the function is called
    Then from the upper limit, count down to the lower limit by decrements of Z
