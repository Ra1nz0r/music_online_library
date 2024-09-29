// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
            "name": "Artem Rylskii",
            "url": "https://t.me/Rainz0r",
            "email": "n52rus@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/library/add": {
            "post": {
                "description": "Добавляет песню в базу данных и делает запрос во внешнее API для получения дополнительных сведений. Если внешнее API недоступно, песня добавляется без дополнительных данных.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "text/plain",
                    "application/json"
                ],
                "tags": [
                    "library"
                ],
                "summary": "Добавляет песню в онлайн библиотеку.",
                "parameters": [
                    {
                        "description": "Данные из запроса для добавления песни.",
                        "name": "db.AddParams",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/db.AddParams"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Успешное добавление песни без дополнительных данных.",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "201": {
                        "description": "Успешное добавление песни с полными данными.",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "integer"
                            }
                        }
                    },
                    "400": {
                        "description": "Некорректный запрос.",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Ошибка сервера при добавлении или обновлении песни.",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/library/delete": {
            "delete": {
                "description": "Обрабатывает DELETE запрос и удаляет песню из библиотеки по указанному ID.",
                "consumes": [
                    "text/plain"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "library"
                ],
                "summary": "Удаляет песню из онлайн библиотеки.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Необходимый ID для удаления песни.",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Успешное удаление песни.",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Некорректный запрос.",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Ошибка сервера при удалении песни.",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/library/update": {
            "put": {
                "description": "По указанному ID обновляет releaseDate, text, link у песни.",
                "consumes": [
                    "text/plain"
                ],
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "library"
                ],
                "summary": "Обновляет параметры песни.",
                "parameters": [
                    {
                        "description": "Данные для обновления.",
                        "name": "models.SongDetail",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.SongDetail"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Успешный запрос и обновление параметров.",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Некорректный запрос.",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Ошибка сервера при обновлении параметров.",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/list": {
            "get": {
                "description": "Получает данные из базы и выводит весь список песен из библиотеки в соответствии с фильтрами.",
                "consumes": [
                    "text/plain"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "library"
                ],
                "summary": "Выводит весь список песен из библиотеки в соответствии с фильтрами.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Имя группы для фильтрации.",
                        "name": "group",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Название композиции для фильтрации.",
                        "name": "song",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Дата релиза для фильтрации.",
                        "name": "releaseDate",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Слова в тексте песни для фильтрации",
                        "name": "text",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Лимит для создания пагинации, по-умолчанию 10.",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Смещение для создания пагинации, по-умолчанию 0.",
                        "name": "offset",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Успешный запрос с учётом фильтрации.",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/db.Library"
                            }
                        }
                    },
                    "400": {
                        "description": "Некорректный запрос.",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Ошибка сервера при создании фильтрации.",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/song/couplet": {
            "get": {
                "description": "Выводит текст по указанному ID, разбитый на куплеты по страницам.",
                "consumes": [
                    "text/plain"
                ],
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "library"
                ],
                "summary": "Текст песни по куплетам.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID группы для поиска композиции.",
                        "name": "id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Номер страницы для пагинации.",
                        "name": "page",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Успешный запрос и разбивка на куплеты.",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Некорректный запрос.",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "db.AddParams": {
            "type": "object",
            "properties": {
                "group": {
                    "type": "string"
                },
                "song": {
                    "type": "string"
                }
            }
        },
        "db.Library": {
            "type": "object",
            "properties": {
                "group": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "link": {
                    "type": "string"
                },
                "releaseDate": {
                    "type": "string"
                },
                "song": {
                    "type": "string"
                },
                "text": {
                    "type": "string"
                }
            }
        },
        "models.SongDetail": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "link": {
                    "type": "string"
                },
                "releaseDate": {
                    "type": "string"
                },
                "text": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:7654",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Music Library",
	Description:      "Implementation of an online song library.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}