Feature: moneycounter
  Scenario: divide the total amount of money by bills
    Given the list of possible bills
      And the total amount of money
    When the function is called
    Then prints out the money divided by as few bills as possible
       And prints out all the numbers in between
