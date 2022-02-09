# Jumia Challenger

Project developed to solve a test scenario applied by the company Jumia.

## How to run

---

For simplicity's sake, the basic commands for the correct operation of the application have been inserted into a Makefile.

It will be necessary to create an `.env` file to include the environment variables used by the project.

Here is an example of the `.env` file:

```env
PORT="8031"
WEB_PORT="8032"
```

Command to execute the code locally:

```bash
make build_react_app
```

and

```bash
make run
```

## How to test

---

Following the pattern of simplicity, a shortcut was created to run all tests in the project.

Command to execute all unit test locally:

```bash
make test
```

## Usage

---

List of endpoints released by the project.

## List Phones

Endpoint responsible to return a list of phones found on customer table.

**Method**

GET

**URL**

```
 http://localhost:8031/phones
```

**Response body**

```json
[
  {
    "country": "Moroco",
    "country_code": 212,
    "phone_number": "6007989253",
    "state": false
  },
  {
    "country": "Moroco",
    "country_code": 212,
    "phone_number": "698054317",
    "state": true
  }
]
```
