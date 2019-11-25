Feature: searching

  Scenario Outline: Limiting the result to "<numberOfResults>"
    Given user is on the site
    When the visitor limits the potential results to "<numberOfResults>"
    And the visitor clicks on the search button
    Then "<numberOfResults>" results should appear on the page

    Examples:
      | numberOfResults |
      |       10        |
      |       20        |
      |       50        |

  Scenario: Check sum of results selecting a specific brand
    Given user is on the site
    When the user select a specific brand
    Then the sum of results in the search button should equal to the number after the name of the brand