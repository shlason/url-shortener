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
        "contact": {
            "name": "sidesideeffect.io",
            "url": "https://github.com/shlason/url-shortener",
            "email": "nocvi111@gmail.com"
        },
        "license": {
            "name": "MIT",
            "url": "https://github.com/shlason/url-shortener/blob/main/LICENSE"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/short": {
            "post": {
                "description": "藉由 timestamp 轉 base62 的方式產生 unique ID 來作為短網址的 ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "short"
                ],
                "summary": "建立短網址",
                "parameters": [
                    {
                        "description": "Original URL",
                        "name": "url",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.createShortURLResponsePayload"
                        }
                    }
                }
            }
        },
        "/{shortID}": {
            "get": {
                "description": "Use 301 redirect by short ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "short"
                ],
                "summary": "Redirect by short ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "shortID redirect use",
                        "name": "shortID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        }
    },
    "definitions": {
        "controllers.createShortURLResponsePayload": {
            "type": "object",
            "properties": {
                "shortId": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "short.sidesideeffect.io",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "URL-Shortener Example API",
	Description:      "This is a sample server celler server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
