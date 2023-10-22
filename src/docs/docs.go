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
        "/enclosure": {
            "get": {
                "description": "Returns the whole enclosure",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "enclosure"
                ],
                "summary": "Get enclosure",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Enclosure"
                        }
                    }
                }
            }
        },
        "/{boxId}/fans/{fanId}": {
            "get": {
                "description": "Retrieve a fan object by its ID and the ID of the box it belongs to",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "fan"
                ],
                "summary": "Get fan",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Box ID",
                        "name": "boxId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Fan ID",
                        "name": "fanId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Fan"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "Update a fan object by its ID and the ID of the box it belongs to",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "fan"
                ],
                "summary": "Update fan",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Box ID",
                        "name": "boxId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Fan ID",
                        "name": "fanId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Fan object",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Fan"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Fan"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/{boxId}/lights/{lightId}": {
            "get": {
                "description": "Get information about a specific light in a specific box",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "light"
                ],
                "summary": "Get light",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Box ID",
                        "name": "boxId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Light ID",
                        "name": "lightId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Light"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "Update the level of a specific light in a specific box",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "light"
                ],
                "summary": "Update light",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Box ID",
                        "name": "boxId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Light ID",
                        "name": "lightId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Light object",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Light"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Light"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/{boxId}/sensors/{sensorId}": {
            "get": {
                "description": "Get a sensor by its box and sensor ID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "sensor"
                ],
                "summary": "Get Sensor",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Box ID",
                        "name": "boxId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Sensor ID",
                        "name": "sensorId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Sensor"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/{boxId}/sensors/{sensorId}/data": {
            "get": {
                "description": "Get time-series data for a sensor by its box and sensor ID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "sensor"
                ],
                "summary": "Get Sensor Data",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Box ID",
                        "name": "boxId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Sensor ID",
                        "name": "sensorId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.TimeSeries"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.Box": {
            "type": "object",
            "properties": {
                "fans": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Fan"
                    }
                },
                "id": {
                    "type": "string"
                },
                "lights": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Light"
                    }
                },
                "location": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "sensors": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Sensor"
                    }
                }
            }
        },
        "model.Enclosure": {
            "type": "object",
            "properties": {
                "boxes": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Box"
                    }
                },
                "id": {
                    "type": "string"
                },
                "location": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "model.Fan": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "level": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "model.Light": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "level": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "state": {
                    "type": "boolean"
                },
                "type": {
                    "$ref": "#/definitions/model.LightType"
                }
            }
        },
        "model.LightType": {
            "type": "string",
            "enum": [
                "MONO"
            ],
            "x-enum-varnames": [
                "MONO"
            ]
        },
        "model.Sensor": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "pin": {
                    "type": "integer"
                },
                "target": {
                    "type": "integer"
                },
                "type": {
                    "$ref": "#/definitions/model.SensorType"
                },
                "unit": {
                    "type": "string"
                }
            }
        },
        "model.SensorType": {
            "type": "string",
            "enum": [
                "TEMP"
            ],
            "x-enum-varnames": [
                "TEMP"
            ]
        },
        "model.TimeSeries": {
            "type": "object",
            "properties": {
                "times": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "values": {
                    "type": "array",
                    "items": {
                        "type": "number"
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "pBox2 API-Docs",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
