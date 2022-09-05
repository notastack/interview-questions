Feature: printallnumbersdivisiblebyz
  Scenario: prints all the numbers divisible by z
    Given a lower limit
      And an upper limit
      And a Z int variable
    When the function is called
    Then goes through all the numbers between the two limits
       And prints out all the numbers divisible by Z
