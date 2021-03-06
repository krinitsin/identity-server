// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "#### API for identity management\"\n",
    "title": "Identity API",
    "contact": {
      "email": "krinitsinv@gmail.com"
    },
    "version": "1.0.0"
  },
  "host": "localhost:8080",
  "basePath": "/api/v1",
  "paths": {
    "/private/identity": {
      "get": {
        "security": [
          {
            "basicAuth": []
          }
        ],
        "description": "View identity after it was seted",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "private"
        ],
        "operationId": "getPrivateIdentity",
        "responses": {
          "200": {
            "description": "View private Identity",
            "schema": {
              "$ref": "#/definitions/IdentityResponse"
            }
          },
          "401": {
            "$ref": "#/responses/UnauthorizedError"
          },
          "412": {
            "$ref": "#/responses/IdentityIsNotSetError"
          },
          "default": {
            "$ref": "#/responses/InternalServerErrorResponse"
          }
        }
      },
      "post": {
        "security": [
          {
            "basicAuth": []
          }
        ],
        "description": "Set identity after registration",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "private"
        ],
        "operationId": "setIdentity",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/SetIdentityRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Identity assigned"
          },
          "401": {
            "$ref": "#/responses/UnauthorizedError"
          },
          "409": {
            "description": "Ethereum address is already assigned to another identity",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "$ref": "#/responses/InternalServerErrorResponse"
          }
        }
      }
    },
    "/public/country/{address}": {
      "get": {
        "description": "Get country assigned to Etherium address",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "public"
        ],
        "operationId": "getPublicCountry",
        "parameters": [
          {
            "minLength": 42,
            "type": "string",
            "name": "address",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Country for Etherium address",
            "schema": {
              "$ref": "#/definitions/CountryResponse"
            }
          },
          "400": {
            "description": "Etherium address is invalid",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "$ref": "#/responses/InternalServerErrorResponse"
          }
        }
      }
    },
    "/public/registration": {
      "post": {
        "description": "Register new account",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "public"
        ],
        "operationId": "registration",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/RegistrationRequest"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "User created"
          },
          "409": {
            "description": "Username already exists",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "$ref": "#/responses/InternalServerErrorResponse"
          }
        }
      }
    }
  },
  "definitions": {
    "CountryResponse": {
      "type": "object",
      "properties": {
        "Country": {
          "type": "string",
          "title": "Country"
        }
      }
    },
    "IdentityResponse": {
      "type": "object",
      "properties": {
        "Country": {
          "type": "string",
          "title": "Country"
        },
        "Username": {
          "type": "string",
          "title": "Username"
        },
        "eth_address": {
          "type": "string",
          "title": "Etherium address"
        }
      }
    },
    "RegistrationRequest": {
      "type": "object",
      "required": [
        "username",
        "password"
      ],
      "properties": {
        "password": {
          "type": "string",
          "title": "Password",
          "minLength": 8
        },
        "username": {
          "type": "string",
          "title": "Username",
          "minLength": 8
        }
      }
    },
    "SetIdentityRequest": {
      "type": "object",
      "required": [
        "eth_address",
        "country"
      ],
      "properties": {
        "country": {
          "type": "string",
          "title": "Country",
          "minLength": 2
        },
        "eth_address": {
          "type": "string",
          "title": "Etherium address",
          "minLength": 8
        }
      }
    },
    "error": {
      "type": "object",
      "properties": {
        "code": {
          "type": "integer",
          "format": "int64"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "principal": {
      "type": "object",
      "properties": {
        "password": {
          "type": "string"
        },
        "username": {
          "type": "string"
        }
      }
    }
  },
  "responses": {
    "IdentityIsNotSetError": {
      "description": "Identity is not set",
      "schema": {
        "$ref": "#/definitions/error"
      }
    },
    "InternalServerErrorResponse": {
      "description": "Internal server error",
      "schema": {
        "$ref": "#/definitions/error"
      }
    },
    "UnauthorizedError": {
      "description": "Authentication information is missing or invalid",
      "headers": {
        "WWW_Authenticate": {
          "type": "string"
        }
      }
    }
  },
  "securityDefinitions": {
    "basicAuth": {
      "type": "basic"
    }
  },
  "tags": [
    {
      "description": "Publicly available path",
      "name": "public"
    },
    {
      "description": "Private path with authorization layer",
      "name": "private"
    }
  ]
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "#### API for identity management\"\n",
    "title": "Identity API",
    "contact": {
      "email": "krinitsinv@gmail.com"
    },
    "version": "1.0.0"
  },
  "host": "localhost:8080",
  "basePath": "/api/v1",
  "paths": {
    "/private/identity": {
      "get": {
        "security": [
          {
            "basicAuth": []
          }
        ],
        "description": "View identity after it was seted",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "private"
        ],
        "operationId": "getPrivateIdentity",
        "responses": {
          "200": {
            "description": "View private Identity",
            "schema": {
              "$ref": "#/definitions/IdentityResponse"
            }
          },
          "401": {
            "description": "Authentication information is missing or invalid",
            "headers": {
              "WWW_Authenticate": {
                "type": "string"
              }
            }
          },
          "412": {
            "description": "Identity is not set",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "description": "Internal server error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "post": {
        "security": [
          {
            "basicAuth": []
          }
        ],
        "description": "Set identity after registration",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "private"
        ],
        "operationId": "setIdentity",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/SetIdentityRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Identity assigned"
          },
          "401": {
            "description": "Authentication information is missing or invalid",
            "headers": {
              "WWW_Authenticate": {
                "type": "string"
              }
            }
          },
          "409": {
            "description": "Ethereum address is already assigned to another identity",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "description": "Internal server error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/public/country/{address}": {
      "get": {
        "description": "Get country assigned to Etherium address",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "public"
        ],
        "operationId": "getPublicCountry",
        "parameters": [
          {
            "minLength": 42,
            "type": "string",
            "name": "address",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Country for Etherium address",
            "schema": {
              "$ref": "#/definitions/CountryResponse"
            }
          },
          "400": {
            "description": "Etherium address is invalid",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "description": "Internal server error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/public/registration": {
      "post": {
        "description": "Register new account",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "public"
        ],
        "operationId": "registration",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/RegistrationRequest"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "User created"
          },
          "409": {
            "description": "Username already exists",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "description": "Internal server error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "CountryResponse": {
      "type": "object",
      "properties": {
        "Country": {
          "type": "string",
          "title": "Country"
        }
      }
    },
    "IdentityResponse": {
      "type": "object",
      "properties": {
        "Country": {
          "type": "string",
          "title": "Country"
        },
        "Username": {
          "type": "string",
          "title": "Username"
        },
        "eth_address": {
          "type": "string",
          "title": "Etherium address"
        }
      }
    },
    "RegistrationRequest": {
      "type": "object",
      "required": [
        "username",
        "password"
      ],
      "properties": {
        "password": {
          "type": "string",
          "title": "Password",
          "minLength": 8
        },
        "username": {
          "type": "string",
          "title": "Username",
          "minLength": 8
        }
      }
    },
    "SetIdentityRequest": {
      "type": "object",
      "required": [
        "eth_address",
        "country"
      ],
      "properties": {
        "country": {
          "type": "string",
          "title": "Country",
          "minLength": 2
        },
        "eth_address": {
          "type": "string",
          "title": "Etherium address",
          "minLength": 8
        }
      }
    },
    "error": {
      "type": "object",
      "properties": {
        "code": {
          "type": "integer",
          "format": "int64"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "principal": {
      "type": "object",
      "properties": {
        "password": {
          "type": "string"
        },
        "username": {
          "type": "string"
        }
      }
    }
  },
  "responses": {
    "IdentityIsNotSetError": {
      "description": "Identity is not set",
      "schema": {
        "$ref": "#/definitions/error"
      }
    },
    "InternalServerErrorResponse": {
      "description": "Internal server error",
      "schema": {
        "$ref": "#/definitions/error"
      }
    },
    "UnauthorizedError": {
      "description": "Authentication information is missing or invalid",
      "headers": {
        "WWW_Authenticate": {
          "type": "string"
        }
      }
    }
  },
  "securityDefinitions": {
    "basicAuth": {
      "type": "basic"
    }
  },
  "tags": [
    {
      "description": "Publicly available path",
      "name": "public"
    },
    {
      "description": "Private path with authorization layer",
      "name": "private"
    }
  ]
}`))
}
