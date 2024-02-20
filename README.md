# Vending Machine Service

## Table of Contents
- [Getting Started](#getting-started)
  - [Installation](#installation)
    - [Web Application Setup](#web-application-setup)
     
    
    - [With CLI (Command Line Interface)](#cli)

  
- [Using API](#vending-machine-api-documentation)


## Getting Started

### Installation


1. Clone the repository:

   ```bash
   git clone https://github.com/Hamedblue1381/vending-machine.git
   cd vending-machine

#### Web Application Setup

  ```bash
        make build
  ```

#### Cli

  ```bash
        make cli
  ``` 



# Vending Machine API Documentation

Welcome to our Vending Machine API! This repository contains a simple RESTful API designed to manage vending machines and their inventory.

## Base URL

All API endpoints are accessible at `http://localhost:8080/`.

## Endpoints

### GET /item/{item}

Dispense an item by its unique identifier.

#### Request

- Method: GET
- Path: `/item/{item}`
- Parameters:
  - `item` (required): The ID of the desired item.

#### Response
- Status code: 200 OK
- Body:

- Status code: 200 OK
- Body:
  ```json
  {
    "message": "Item dispensed successfully!"
  }

### GET /insert

Retrieve a list of items currently available for purchase.

#### Request

- Method: GET
- Path: `/insert`

#### Response

- Status code: 200 OK
- Body:
  ```json
  {
    [
        {
          "ID": 1,
          "Name": "Soda",
          "Count": 20
          },
          {
          "ID": 2,
          "Name": "Coffee",
          "Count": 19
          }
    ]
  }

### GET /machines

Obtain information about existing vending machines and their current stock levels.

#### Request

- Method: GET
- Path: `/machines`

#### Response

- Status code: 200 OK
- Body:
 ```json
{
    [
        {
        "id": 0,
        "stock": 40,
        "coins": 0,
        "items": [
            {
            "ID": 1,
            "Name": "Soda",
            "Count": 20
            },
            {
            "ID": 2,
            "Name": "Coffee",
            "Count": 20
            }
                ]
        }
    ]
}
```
### POST /add-machine

Add a new vending machine to the system.

#### Request

- Method: POST
- Path: `/add-machine`
- Body:
```json
{
"count": 20
}
```

