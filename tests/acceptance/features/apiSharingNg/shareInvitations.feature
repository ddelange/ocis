Feature: Send a sharing invitations
  As the owner of a resource
  I want to be able to send invitations to other users
  So that they can have access to it

  https://owncloud.dev/libre-graph-api/#/drives.permissions/Invite

  Background:
    Given these users have been created with default attributes and without skeleton files:
      | username |
      | Alice    |
      | Brian    |


  Scenario Outline: send share invitation to user with different roles
    Given user "Alice" has uploaded file with content "to share" to "/textfile1.txt"
    And user "Alice" has created folder "FolderToShare"
    When user "Alice" sends the following share invitation using the Graph API:
      | resourceType | <resource-type> |
      | resource     | <path>          |
      | space        | Personal        |
      | sharee       | Brian           |
      | shareType    | user            |
      | role         | <role>          |
    Then the HTTP status code should be "200"
    And the JSON data of the response should match
      """
      {
        "type": "object",
        "required": [
          "value"
        ],
        "properties": {
          "value": {
            "type": "array",
            "items": {
              "type": "object",
              "required": [
                "id",
                "roles",
                "grantedToV2"
              ],
              "properties": {
                "id": {
                  "type": "string",
                  "pattern": "^%share_id_pattern%$"
                },
                "roles": {
                  "type": "array",
                  "items": {
                    "type": "string",
                    "pattern": "^%role_id_pattern%$"
                  }
                },
                "grantedToV2": {
                  "type": "object",
                  "required": [
                    "user"
                  ],
                  "properties": {
                    "user": {
                      "type": "object",
                      "required": [
                        "id",
                        "displayName"
                      ],
                      "properties": {
                        "id": {
                          "type": "string",
                          "pattern": "^%user_id_pattern%$"
                        },
                        "displayName": {
                          "type": "string",
                          "enum": [
                            "Brian Murphy"
                          ]
                        }
                      }
                    }
                  }
                }
              }
            }
          }
        }
      }
      """
    Examples:
      | role        | resource-type | path           |
      | Viewer      | file          | /textfile1.txt |
      | File Editor | file          | /textfile1.txt |
      | Co Owner    | file          | /textfile1.txt |
      | Manager     | file          | /textfile1.txt |
      | Viewer      | folder        | FolderToShare  |
      | Editor      | folder        | FolderToShare  |
      | Co Owner    | folder        | FolderToShare  |
      | Uploader    | folder        | FolderToShare  |
      | Manager     | folder        | FolderToShare  |


  Scenario Outline: send share invitation to group with different roles
    Given user "Carol" has been created with default attributes and without skeleton files
    And group "grp1" has been created
    And the following users have been added to the following groups
      | username | groupname |
      | Brian    | grp1      |
      | Carol    | grp1      |
    And user "Alice" has uploaded file with content "to share" to "/textfile1.txt"
    And user "Alice" has created folder "FolderToShare"
    When user "Alice" sends the following share invitation using the Graph API:
      | resourceType | <resource-type> |
      | resource     | <path>          |
      | space        | Personal        |
      | sharee       | grp1            |
      | shareType    | group           |
      | role         | <role>          |
    Then the HTTP status code should be "200"
    And the JSON data of the response should match
      """
      {
        "type": "object",
        "required": [
          "value"
        ],
        "properties": {
          "value": {
            "type": "array",
            "items": {
              "type": "object",
              "required": [
                "id",
                "roles",
                "grantedToV2"
              ],
              "properties": {
                "id": {
                  "type": "string",
                  "pattern": "^%share_id_pattern%$"
                },
                "roles": {
                  "type": "array",
                  "items": {
                    "type": "string",
                    "pattern": "^%role_id_pattern%$"
                  }
                },
                "grantedToV2": {
                  "type": "object",
                  "required": [
                    "group"
                  ],
                  "properties": {
                    "group": {
                      "type": "object",
                      "required": [
                        "id",
                        "displayName"
                      ],
                      "properties": {
                        "id": {
                          "type": "string",
                          "pattern": "^%group_id_pattern%$"
                        },
                        "displayName": {
                          "type": "string",
                          "enum": [
                            "grp1"
                          ]
                        }
                      }
                    }
                  }
                }
              }
            }
          }
        }
      }
      """
    Examples:
      | role        | resource-type | path           |
      | Viewer      | file          | /textfile1.txt |
      | File Editor | file          | /textfile1.txt |
      | Co Owner    | file          | /textfile1.txt |
      | Manager     | file          | /textfile1.txt |
      | Viewer      | folder        | FolderToShare  |
      | Editor      | folder        | FolderToShare  |
      | Co Owner    | folder        | FolderToShare  |
      | Uploader    | folder        | FolderToShare  |
      | Manager     | folder        | FolderToShare  |