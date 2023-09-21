// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/books": {
            "get": {
                "description": "Get All Books",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get All Books",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/main.Books"
                            }
                        }
                    }
                }
            }
        },
        "/books/search": {
            "get": {
                "description": "Get Books by search criteria",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Search Books",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Some ID",
                        "name": "some_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/books/{some_id}": {
            "get": {
                "description": "get book by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get Book by ID",
                "operationId": "int",
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "description": "Update a book",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Update a book",
                "operationId": "int",
                "parameters": [
                    {
                        "description": "body data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.Books"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "$ref": "#/definitions/main.Books"
                        }
                    }
                }
            },
            "post": {
                "description": "Add a new book",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Add a new book",
                "parameters": [
                    {
                        "description": "body data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.Books"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/main.Books"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a book",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Delete a Book",
                "operationId": "int",
                "responses": {
                    "204": {
                        "description": "No Content"
                    }
                }
            },
            "patch": {
                "description": "Partially update a book",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Partially update a book",
                "operationId": "int",
                "parameters": [
                    {
                        "description": "same as post body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.Books"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "$ref": "#/definitions/main.Books"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "main.Books": {
            "type": "object",
            "properties": {
                "book_author": {
                    "type": "string"
                },
                "book_genre": {
                    "type": "string"
                },
                "book_publishers": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "isbn": {
                    "type": "string"
                },
                "langauge": {
                    "type": "string"
                },
                "prices": {
                    "type": "number"
                },
                "status": {
                    "type": "boolean"
                },
                "title": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "localhost:5000",
	BasePath:         "/",
	Schemes:          []string{"http"},
	Title:            "Books API",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}