Feature: create auth-app token
  As a user
  I want to create auth-app Tokens
  So that I can use 3rd party apps

  Background:
    Given user "Alice" has been created with default attributes


  Scenario: user creates auth-app token
    When user "Alice" creates auth-app token with expiration time "72h" using the auth-app API
    Then the HTTP status code should be "200"
    And the JSON data of the response should match
      """
      {
        "type": "object",
        "required": [
          "token",
          "expiration_date",
          "created_date",
          "label"
        ],
        "properties": {
          "token": {
            "pattern": "^[a-zA-Z0-9]{16}$"
          },
          "label": {
            "const": "Generated via API"
          }
        }
      }
      """


  Scenario: user creates auth-app token with custom label
    When user "Alice" creates auth-app token with expiration time "72h" and label "Custom label" using the auth-app API
    Then the HTTP status code should be "200"
    And the JSON data of the response should match
      """
      {
        "type": "object",
        "required": [
          "token",
          "expiration_date",
          "created_date",
          "label"
        ],
        "properties": {
          "token": {
            "pattern": "^[a-zA-Z0-9]{16}$"
          },
          "label": {
            "const": "Custom label"
          }
        }
      }
      """

  @env-config
  Scenario: user lists auth-app tokens generated by different auth-app api
    Given the config "AUTH_APP_ENABLE_IMPERSONATION" has been set to "true"
    And user "Alice" has created auth-app token with expiration time "72h" using the auth-app API
    And user "Admin" has created auth-app token for user "Alice" with expiration time "72h" using the auth-app API
    When user "Alice" lists all created tokens using the auth-app API
    Then the HTTP status code should be "200"
    And the JSON data of the response should match
      """
      {
        "type": "array",
        "minItems": 2,
        "maxItems": 2,
        "uniqueItems": true,
        "items": {
          "oneOf": [
            {
              "type": "object",
              "required": [
                "token",
                "expiration_date",
                "created_date",
                "label"
              ],
              "properties": {
                "token": {
                  "pattern": "^\\$2a\\$11\\$[A-Za-z0-9./]{53}$"
                },
                "label": {
                  "const": "Generated via API"
                }
              }
            },
            {
              "type": "object",
              "required": [
                "token",
                "expiration_date",
                "created_date",
                "label"
              ],
              "properties": {
                "token": {
                  "pattern": "^\\$2a\\$11\\$[A-Za-z0-9./]{53}$"
                },
                "label": {
                  "const": "Generated via Impersonation API"
                }
              }
            }
          ]
        }
      }
      """

  @env-config
  Scenario: admin creates auth-app token for other user
    Given the config "AUTH_APP_ENABLE_IMPERSONATION" has been set to "true"
    When user "Admin" creates auth-app token for user "Alice" with expiration time "72h" using the auth-app API
    Then the HTTP status code should be "200"
    And the JSON data of the response should match
      """
      {
        "type": "object",
        "required": [
          "token",
          "expiration_date",
          "created_date",
          "label"
        ],
        "properties": {
          "token": {
            "pattern": "^[a-zA-Z0-9]{16}$"
          },
          "label": {
            "const": "Generated via Impersonation API"
          }
        }
      }
      """

  @env-config
  Scenario: user deletes the created auth-app token
    Given the config "AUTH_APP_ENABLE_IMPERSONATION" has been set to "true"
    And user "Alice" has created auth-app token with expiration time "72h" using the auth-app API
    And user "Admin" has created auth-app token for user "Alice" with expiration time "72h" using the auth-app API
    When user "Alice" deletes all the created auth-app tokens using the auth-app API
    Then the HTTP status code of responses on all endpoints should be "200"
    And user "Alice" should have "0" auth-app tokens
    When user "Alice" lists all created tokens using the auth-app API
    Then the HTTP status code should be "200"
    And the JSON data of the response should match
      """
      {
        "type": "array",
        "minItems": 0,
        "maxItems": 0
      }
      """


  Scenario: admin tries to create auth-app token for other user without impersonation enabled
    When user "Admin" tries to create auth-app token for user "Alice" with expiration time "72h" using the auth-app API
    Then the HTTP status code should be "403"
    And the content in the response should include the following content:
      """
      impersonation is not allowed
      """


  Scenario: non-admin user tries to create auth-app token without expiry
    When user "Alice" tries to create auth-app token with expiration time "" using the auth-app API
    Then the HTTP status code should be "400"
    And the content in the response should include the following content:
      """
      error parsing expiry. Use e.g. 30m or 72h
      """

  @env-config
  Scenario: admin tries to create auth-app token for other users without expiry
    Given the config "AUTH_APP_ENABLE_IMPERSONATION" has been set to "true"
    When user "Admin" tries to create auth-app token for user "Alice" with expiration time "" using the auth-app API
    Then the HTTP status code should be "400"
    And the content in the response should include the following content:
      """
      error parsing expiry. Use e.g. 30m or 72h
      """

  @env-config
  Scenario: non-admin user tries to create an auth-app token for another user
    Given user "Brian" has been created with default attributes
    And the config "AUTH_APP_ENABLE_IMPERSONATION" has been set to "true"
    When user "Alice" tries to create auth-app token for user "Brian" with expiration time "72h" using the auth-app API
    Then the HTTP status code should be "403"

  @env-config @issue-10815
  Scenario: admin tries to create auth-app token for non-existing user
    Given the config "AUTH_APP_ENABLE_IMPERSONATION" has been set to "true"
    When user "Admin" creates auth-app token for user "Brian" with expiration time "72h" using the auth-app API
    Then the HTTP status code should be "403"

  @env-config @issue-10815
  Scenario: admin user tries to delete auth-app token of another user with impersonation enabled
    Given the config "AUTH_APP_ENABLE_IMPERSONATION" has been set to "true"
    And user "Admin" has created auth-app token for user "Alice" with expiration time "72h" using the auth-app API
    When user "Admin" tries to delete the last created auth-app token using the auth-app API
    Then the HTTP status code should be "403"

  @issue-10921
  Scenario: admin user tries to delete auth-app token of a normal user
    Given user "Alice" has created auth-app token with expiration time "72h" using the auth-app API
    When user "Admin" tries to delete the last created auth-app token using the auth-app API
    Then the HTTP status code should be "403"

  @issue-10921
  Scenario: non-admin user tries to delete auth-app token of another user
    Given user "Brian" has been created with default attributes
    And user "Brian" has created auth-app token with expiration time "72h" using the auth-app API
    When user "Admin" tries to delete the last created auth-app token using the auth-app API
    Then the HTTP status code should be "403"
