// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/auth/image/{image_id}": {
            "get": {
                "description": "get list image",
                "tags": [
                    "Images"
                ],
                "summary": "image",
                "parameters": [
                    {
                        "type": "string",
                        "description": "image id",
                        "name": "image_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "'Bearer ' + token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.ImageResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "update image",
                "tags": [
                    "Images"
                ],
                "summary": "image",
                "parameters": [
                    {
                        "type": "string",
                        "description": "image id",
                        "name": "image_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "'Bearer ' + token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "image content",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.ImageUpdate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/auth/images": {
            "get": {
                "description": "get list image",
                "tags": [
                    "Images"
                ],
                "summary": "image",
                "parameters": [
                    {
                        "type": "string",
                        "name": "block_id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "status",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "user_id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "'Bearer ' + token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "status",
                        "name": "status",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.ImageResponse"
                            }
                        }
                    }
                }
            },
            "delete": {
                "description": "get list image",
                "tags": [
                    "Images"
                ],
                "summary": "image",
                "parameters": [
                    {
                        "type": "string",
                        "name": "block_id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "status",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "user_id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "'Bearer ' + token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/auth/images/block-ids": {
            "get": {
                "description": "get list image",
                "tags": [
                    "Images"
                ],
                "summary": "image",
                "parameters": [
                    {
                        "type": "string",
                        "description": "'Bearer ' + token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/api/v1/auth/object": {
            "post": {
                "description": "upload object",
                "tags": [
                    "Object"
                ],
                "summary": "upload object",
                "parameters": [
                    {
                        "type": "string",
                        "description": "add block id",
                        "name": "block_id",
                        "in": "header"
                    },
                    {
                        "type": "string",
                        "description": "'Bearer ' + token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "add file multipart/form-data",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.ImageResponse"
                        }
                    },
                    "400": {
                        "description": "some info",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "not found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "...",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/auth/object/{id}": {
            "get": {
                "description": "download object",
                "tags": [
                    "Object"
                ],
                "summary": "download object",
                "parameters": [
                    {
                        "type": "string",
                        "description": "object id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "'Bearer ' + token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "integer"
                            }
                        }
                    },
                    "400": {
                        "description": "some info",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "not found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "...",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.ImageFilter": {
            "type": "object",
            "properties": {
                "block_id": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "model.ImageResponse": {
            "type": "object",
            "properties": {
                "block_id": {
                    "type": "string"
                },
                "create_at": {
                    "type": "string"
                },
                "data": {
                    "type": "string"
                },
                "error": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "model.ImageUpdate": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "string"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "OCR BACKEND API",
	Description: "",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
