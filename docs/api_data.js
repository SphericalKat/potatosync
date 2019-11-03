define({ "api": [
  {
    "type": "post",
    "url": "/api/users/new",
    "title": "",
    "name": "Create_a_new_user_account",
    "group": "Authorization",
    "permission": [
      {
        "name": "anyone"
      }
    ],
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "email",
            "description": "<p>User's email address</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "username",
            "description": "<p>Username</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "size": "8..60",
            "optional": false,
            "field": "password",
            "description": "<p>Password to use for this account</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "request-example",
          "content": "\n{\n\t\"email\": \"john.doe@example.com\",\n\t\"username\": \"JohnDoe\",\n\t\"password\": \"password123\"\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "controllers/auth_controller.go",
    "groupTitle": "Authorization"
  },
  {
    "type": "post",
    "url": "/api/users/delete",
    "title": "",
    "name": "Delete_a_user_account",
    "group": "Authorization",
    "permission": [
      {
        "name": "Logged-in users"
      }
    ],
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "string",
            "optional": false,
            "field": "Authorization",
            "description": "<p>JWT token associated with user account</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Header-Example:",
          "content": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjV9.FTIhjfCLND1L-hvhgH9_TC_P7MbGQnjnNnFOjJL8Q1k",
          "type": "string"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "controllers/auth_controller.go",
    "groupTitle": "Authorization"
  },
  {
    "type": "post",
    "url": "/api/users/login",
    "title": "",
    "name": "Log_into_user_account",
    "group": "Authorization",
    "permission": [
      {
        "name": "anyone"
      }
    ],
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "string",
            "optional": true,
            "field": "email",
            "description": "<p>User's email address</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": true,
            "field": "username",
            "description": "<p>Username</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "size": "8..60",
            "optional": false,
            "field": "password",
            "description": "<p>Password to use for this account</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "request-example",
          "content": "\n{\n\t\"email\": \"john.doe@example.com\",\n\t\"username\": \"JohnDoe\",\n\t\"password\": \"password123\"\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "controllers/auth_controller.go",
    "groupTitle": "Authorization"
  },
  {
    "type": "post",
    "url": "/api/users/manage/password",
    "title": "",
    "name": "Modify_account_s_password",
    "group": "Authorization",
    "permission": [
      {
        "name": "Logged-in users"
      }
    ],
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "password",
            "description": "<p>The new password</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "request-example",
          "content": "\n{\n\t\"password\": \"password123\",\n}",
          "type": "json"
        }
      ]
    },
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "string",
            "optional": false,
            "field": "Authorization",
            "description": "<p>JWT token associated with user account</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Header-Example:",
          "content": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjV9.FTIhjfCLND1L-hvhgH9_TC_P7MbGQnjnNnFOjJL8Q1k",
          "type": "string"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "controllers/mgmt_controller.go",
    "groupTitle": "Authorization"
  },
  {
    "type": "post",
    "url": "/api/users/manage/image",
    "title": "",
    "name": "Modify_account_s_password",
    "group": "Authorization",
    "permission": [
      {
        "name": "Logged-in users"
      }
    ],
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "image_url",
            "description": "<p>The url to the user's profile image</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "request-example",
          "content": "\n{\n\t\"image_url\": \"https://example.com/1234567\",\n}",
          "type": "json"
        }
      ]
    },
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "string",
            "optional": false,
            "field": "Authorization",
            "description": "<p>JWT token associated with user account</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Header-Example:",
          "content": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjV9.FTIhjfCLND1L-hvhgH9_TC_P7MbGQnjnNnFOjJL8Q1k",
          "type": "string"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "controllers/mgmt_controller.go",
    "groupTitle": "Authorization"
  },
  {
    "type": "post",
    "url": "/api/users/manage/username",
    "title": "",
    "name": "Modify_account_s_username",
    "group": "Authorization",
    "permission": [
      {
        "name": "Logged-in users"
      }
    ],
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "email",
            "description": "<p>The user's email</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "username",
            "description": "<p>The new username</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "request-example",
          "content": "\n{\n\t\"email\": \"john.doe@example.com\",\n\t\"username\": \"JohnDoe\",\n}",
          "type": "json"
        }
      ]
    },
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "string",
            "optional": false,
            "field": "Authorization",
            "description": "<p>JWT token associated with user account</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Header-Example:",
          "content": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjV9.FTIhjfCLND1L-hvhgH9_TC_P7MbGQnjnNnFOjJL8Q1k",
          "type": "string"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "controllers/mgmt_controller.go",
    "groupTitle": "Authorization"
  },
  {
    "type": "get",
    "url": "/api/users/info",
    "title": "",
    "name": "View_info_about_user_account",
    "group": "Authorization",
    "permission": [
      {
        "name": "Logged-in users"
      }
    ],
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "string",
            "optional": false,
            "field": "Authorization",
            "description": "<p>JWT token associated with user account</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Header-Example:",
          "content": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjV9.FTIhjfCLND1L-hvhgH9_TC_P7MbGQnjnNnFOjJL8Q1k",
          "type": "string"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "controllers/auth_controller.go",
    "groupTitle": "Authorization"
  },
  {
    "type": "post",
    "url": "/api/notes/save",
    "title": "",
    "name": "Create_or_update_a_note",
    "group": "Notes",
    "permission": [
      {
        "name": "Logged-in users"
      }
    ],
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "number",
            "optional": false,
            "field": "note_id",
            "description": "<p>The note's ID</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "title",
            "description": "<p>The note's title</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "content",
            "description": "<p>The note's content</p>"
          },
          {
            "group": "Parameter",
            "type": "boolean",
            "optional": false,
            "field": "is_starred",
            "description": "<p>Whether the note is starred or not</p>"
          },
          {
            "group": "Parameter",
            "type": "number",
            "optional": false,
            "field": "date",
            "description": "<p>Date of note creation</p>"
          },
          {
            "group": "Parameter",
            "type": "number",
            "optional": false,
            "field": "color",
            "description": "<p>The color of the note</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "image_path",
            "description": "<p>Image path</p>"
          },
          {
            "group": "Parameter",
            "type": "boolean",
            "optional": false,
            "field": "is_list",
            "description": "<p>Whether the note contains a list or not</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "list_parse_string",
            "description": "<p>List parse string</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "reminders",
            "description": "<p>Reminders associated with the note</p>"
          },
          {
            "group": "Parameter",
            "type": "boolean",
            "optional": false,
            "field": "hide_content",
            "description": "<p>Whether to hide the note or not</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "pin",
            "description": "<p>Pin</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "password",
            "description": "<p>Password to lock the note with</p>"
          },
          {
            "group": "Parameter",
            "type": "boolean",
            "optional": false,
            "field": "is_deleted",
            "description": "<p>Is deleted or not</p>"
          },
          {
            "group": "Parameter",
            "type": "boolean",
            "optional": false,
            "field": "is_archived",
            "description": "<p>Is archived or not</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "request-example",
          "content": "\n{\n\t\"note_id\": 5,\n\t\"title\": \"Sample note 2\",\n\t\"content\": \"This is a sample note 2\",\n\t\"is_starred\": true,\n\t\"date\": 20191211,\n\t\"color\": 4,\n\t\"image_path\": \"abcd\",\n\t\"is_list\": false,\n\t\"list_parse_string\": \"abcd\",\n\t\"reminders\": \"abcd\",\n\t\"hide_content\": false,\n\t\"pin\": \"smth\",\n\t\"password\": \"password123\",\n\t\"is_deleted\": false,\n\t\"is_archived\": false\n}",
          "type": "json"
        }
      ]
    },
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "string",
            "optional": false,
            "field": "Authorization",
            "description": "<p>JWT token associated with user account</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Header-Example:",
          "content": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjV9.FTIhjfCLND1L-hvhgH9_TC_P7MbGQnjnNnFOjJL8Q1k",
          "type": "string"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "controllers/notes_controller.go",
    "groupTitle": "Notes"
  },
  {
    "type": "post",
    "url": "/api/notes/delete",
    "title": "",
    "name": "Delete_a_saved_note",
    "group": "Notes",
    "permission": [
      {
        "name": "Logged-in users"
      }
    ],
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "number",
            "optional": false,
            "field": "note_id",
            "description": "<p>The note's ID</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "request-example",
          "content": "\n{\n\t\"note_id\": 5,\n}",
          "type": "json"
        }
      ]
    },
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "string",
            "optional": false,
            "field": "Authorization",
            "description": "<p>JWT token associated with user account</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Header-Example:",
          "content": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjV9.FTIhjfCLND1L-hvhgH9_TC_P7MbGQnjnNnFOjJL8Q1k",
          "type": "string"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "controllers/notes_controller.go",
    "groupTitle": "Notes"
  },
  {
    "type": "post",
    "url": "/api/notes/deleteall",
    "title": "",
    "name": "Delete_all_saved_notes_for_a_user",
    "group": "Notes",
    "permission": [
      {
        "name": "Logged-in users"
      }
    ],
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "string",
            "optional": false,
            "field": "Authorization",
            "description": "<p>JWT token associated with user account</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Header-Example:",
          "content": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjV9.FTIhjfCLND1L-hvhgH9_TC_P7MbGQnjnNnFOjJL8Q1k",
          "type": "string"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "controllers/notes_controller.go",
    "groupTitle": "Notes"
  },
  {
    "type": "get",
    "url": "/api/notes/list",
    "title": "",
    "name": "List_notes_associated_with_a_user",
    "group": "Notes",
    "permission": [
      {
        "name": "Logged-in users"
      }
    ],
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "string",
            "optional": false,
            "field": "Authorization",
            "description": "<p>JWT token associated with user account</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Header-Example:",
          "content": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjV9.FTIhjfCLND1L-hvhgH9_TC_P7MbGQnjnNnFOjJL8Q1k",
          "type": "string"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "controllers/notes_controller.go",
    "groupTitle": "Notes"
  }
] });
