**Overview**
----
This project is completed as a test task for golang developer.  
It consists of two services connecting through grpc.  
Spawn-service has one method that generate random string contains  
digits '0-9' and latin character 'A-Z' and 'a-z'.  
Users-service has two http endpoints. First gets user's email and  
password, then creates user in mongo database.  Second endpoint  
return found user in database by given email.  

**Content**
----  

- [Status codes](#status-codes)
- [Create user](#create-user)
- [Get user](#get-user)
- [Usage](#usage)
  
**Status codes**
----

The following table gives an overview of how the API functions generally behave.


The following table shows the possible return codes for API requests.

| Return values | Description |
| ------------- | ----------- |
| `200 OK` | request was successful. |
| `201 Created` | The `POST` request was successful. |
| `400 Bad Request` | Wrond data in request body. |
| `404 Not Found` | User with such email does not exist |
| `405 Method Not Allowed` | The request is not supported. |
| `409 Conflict` | User with given email already exists |
| `500 Server Error` | While handling the request something went wrong on server-side. |  

**Create user**
----
  Accept json with two field 'email' and 'password'. Both should not be empty.  
  Email should be valid.  

* **URL**

  /create-user

* **Method:**

  `POST`
  
*  **URL Params**

    None

* **Data Params**

```json
{
  "email": "some email",
  "password": "some password",
}
```

* **Success Response:**

  * **Code:** 201 <br />
    **Content:** 

```json
{}
```
 
* **Error Response:**

  * *Request has wrong json fields*
    **Code:** 400 BAD REQUEST <br />
    **Content:** 
```json
{
    "error": "json format is not correct"
}
```

  OR with details

```json
{
    "error": "json body has empty fields",
    "detail": "email has wrong format"
}
```  

  * *There is already user with same email*
    **Code:** 409 STATUS CONFLICT <br />
    **Content:** 
```json
{
    "error": "user with such email already exists"
}
```

**Get user**
----
  Return JSON with user's information.

* **URL**

  /get-user/{email}

* **Method:**

  `GET`
  
*  **URL Params**

   None

* **Data Params**

  None

* **Success Response:**

  * **Code:** 200 OK <br />
    **Content:** 

```json
{
    {
      "id": "63eb882f4388086f4a1c0e0f",
      "email": "some@email.com",
      "salt": "Q3O58J32dHTx",
      "password": "19c8a5497ebb1cf76916d9cefb1caeef"
    }
}
```

* **Error Response:**

  * *There is no user with that email*
    **Code:** 404 NOT FOUND <br />
    **Content:** 
```json
{
    "error": "user does not exist"
}
```


**Usage**
----
Run app
```
make up
```  
Open swagger
```
http://localhost:8081/swagger/index.html#
```  
Stop container
```
make down
```  
Begin tests
```
cd ./users-service
make test
```  
Begin tests with cover
```
cd ./users-service
make cover
```  