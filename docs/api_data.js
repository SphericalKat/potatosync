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
            "optional": false,
            "field": "password",
            "description": "<p>Password to use for this account</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "request-example",
          "content": "\n{\n\t\"email\": \"amolele@gmail.com\",\n\t\"username\": \"SphericalKat\",\n\t\"password\": \"forthencho\"\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "./controllers/auth_controller.go",
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
    "filename": "./controllers/auth_controller.go",
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
            "optional": false,
            "field": "password",
            "description": "<p>Password to use for this account</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "request-example",
          "content": "\n{\n\t\"email\": \"amolele@gmail.com\",\n\t\"username\": \"SphericalKat\",\n\t\"password\": \"forthencho\"\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "./controllers/auth_controller.go",
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
    "filename": "./controllers/auth_controller.go",
    "groupTitle": "Authorization"
  }
] });
