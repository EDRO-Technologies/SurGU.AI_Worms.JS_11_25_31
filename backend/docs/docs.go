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
        "/api/legendaryPrompt": {
            "post": {
                "description": "Post prompt",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Prompts"
                ],
                "summary": "PingPong",
                "parameters": [
                    {
                        "description": "prompt",
                        "name": "prompt",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.Prompt"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/api/prompt": {
            "post": {
                "description": "Post prompt",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Prompts"
                ],
                "summary": "PostPrompt",
                "parameters": [
                    {
                        "description": "prompt",
                        "name": "prompt",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.Prompt"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "types.Message": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "string"
                },
                "is_ai": {
                    "type": "boolean"
                }
            }
        },
        "types.Prompt": {
            "type": "object",
            "properties": {
                "act_id": {
                    "type": "integer"
                },
                "msg": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/types.Message"
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
