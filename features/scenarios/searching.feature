Feature: searching

  Scenario: Limiting the results to a specific number
    Given user is on the site
    When the visitor limits the potential results to "50"
    And the visitor clicks on the search button
    Then "50" results should appear on the page

  Scenario: Check sum of results selecting a specific brand
    Given user is on the site
    When the user select a specific brand
    Then the sum of results in the search button should equal to the number after the name of the brand