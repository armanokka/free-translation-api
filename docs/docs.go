// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/detect": {
            "post": {
                "description": "Detect language the text is written in",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Detect language the text is written in",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Text to translate",
                        "name": "text",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.ResponseDetection"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/Response"
                        }
                    }
                }
            }
        },
        "/translate": {
            "post": {
                "description": "Translate text and detect its language",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Translate text",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Text to translate",
                        "name": "text",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Source language",
                        "name": "from",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "Target language",
                        "name": "to",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/ResponseTranslation"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "Response": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                },
                "ok": {
                    "type": "boolean"
                }
            }
        },
        "ResponseTranslation": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                },
                "ok": {
                    "type": "boolean"
                },
                "text_lang": {
                    "type": "string"
                },
                "translated_text": {
                    "description": "string if only one text is translated, array if more",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "v1.ResponseDetection": {
            "type": "object",
            "properties": {
                "detected_languages": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "error": {
                    "type": "string"
                },
                "ok": {
                    "type": "boolean"
                }
            }
        }
    },
    "securityDefinitions": {
        "BasicAuth": {
            "type": "basic"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:80",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Swagger Example API",
	Description:      "The free Google translation API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
