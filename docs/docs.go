// Code generated by swaggo/swag. DO NOT EDIT
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
        "/api/v1/address/search": {
            "get": {
                "description": "Get addresses by string.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Addresses"
                ],
                "summary": "Get address suggestions by string",
                "operationId": "get-addresses-by-string",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Address ID",
                        "name": "q",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/domain.AutocompletePrediction"
                            }
                        }
                    },
                    "404": {
                        "description": "Please input a valid string",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/address/validate": {
            "get": {
                "description": "Get an address by its ID.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Addresses"
                ],
                "summary": "Get address by ID",
                "operationId": "get-address-by-id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Address ID",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.Address"
                        }
                    },
                    "404": {
                        "description": "address not found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "domain.Address": {
            "type": "object",
            "properties": {
                "addressComponents": {
                    "description": "The components of the address, such as street number, street name, city, state, etc."
                },
                "formattedAddress": {
                    "description": "The formatted address string.",
                    "type": "string"
                },
                "type": {
                    "description": "The type of the address, such as \"google\" or \"geoscape\".",
                    "type": "string"
                }
            }
        },
        "domain.AutocompletePrediction": {
            "type": "object",
            "properties": {
                "description": {
                    "description": "Formatted Address.",
                    "type": "string"
                },
                "place_id": {
                    "description": "address id.",
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1",
	Host:             "localhost:3000",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Location Service",
	Description:      "Location Service",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
