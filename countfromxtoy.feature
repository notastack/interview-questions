Feature: countfromxtoy
  Scenario: count up untill reaching the limit
    Given two limits
    When the function is called
    Then from the lower limit, count up to the upper limit
       And prints out all the numbers in between
