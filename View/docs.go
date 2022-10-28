package View

import "github.com/swaggo/swag"

const docTemplate = `{
  "openapi": "3.0.3",
  "info": {
    "title": "Golang MVC Structure",
    "version": "1.0.11"
  },
  "tags": [
    {
      "name": "student",
      "description": "Testing of MVC's structure"
    }
  ],
  "paths": {
    "/createStudent": {
      "post": {
        "tags": [
          "student"
        ],
        "summary": "create student",
        "requestBody": {
          "description": "create student",
          "required": true,
          "content": {
            "application/json": {
              "schema": {
                "$ref": "#/components/schemas/student"
              }
            },
            "application/xml": {
              "schema": {
                "$ref": "#/components/schemas/student"
              }
            }
          }
        },
        "responses": {
          "201": {
            "description": "Upload Successful",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/student"
                }
              }
            }
          }
        }
      }
    },
    "/getStudent": {
      "get": {
        "tags": [
          "student"
        ],
        "summary": "list all student",
        "responses": {
          "201": {
            "description": "list Successful",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/student"
                }
              }
            }
          }
        }
      }
    },
    "/randomSport": {
      "get": {
        "tags": [
          "student"
        ],
        "summary": "random all student",
        "responses": {
          "201": {
            "description": "random Successful",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/student"
                }
              }
            }
          }
        }
      }
    },
    "/getSport": {
      "get": {
        "tags": [
          "student"
        ],
        "summary": "random all student",
        "responses": {
          "201": {
            "description": "random Successful",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/student"
                }
              }
            }
          }
        }
      }
    }
  },
  "components": {
    "schemas": {
      "student": {
        "type": "object",
        "properties": {
          "student_id": {
            "type": "integer",
            "format": "int64"
          },
          "first_name": {
            "type": "string"
          },
          "last_name": {
            "type": "string"
          },
          "year": {
            "type": "integer",
            "format": "int64"
          }
        },
        "xml": {
          "name": "Picture"
        }
      }
    }
  }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Golang MVC",
	Description:      "This is a sample server todo server",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
