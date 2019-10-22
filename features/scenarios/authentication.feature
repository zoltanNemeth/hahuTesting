Feature: authentication

  Scenario Outline: Login with invalid credentials
    Given user is on the site
    And the user accepts the privacy information
    When the user chooses the Login menu
    Then ensure to provide a login form with username and password field
    When the user submits the form with invalid credentials ("<username>", "<password>")
    Then an error message appears.

    Examples:
      | username |     password     |
      | username |   wrongPassword  |