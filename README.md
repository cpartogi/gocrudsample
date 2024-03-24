# Application Documentation

## Folder Structure
1. /domain/tutorial -> application folder
2. /domain/tutorial/handler ->  contains handler for API Endpoints
3. /domain/tutorial/model -> contains model
4. /domain/tutorial/repo -> contains function to get and update data
5. /domain/tutorial/usecase -> contains business logic function
6. /init -> contains initial function
7. /lib -> internal library
8. /migration -> data migration script
9. /schema ->  contains request and response parameters for API Endpoints

## How To Install
1. If you don't have posgresql and redis installed, you can use docker with command :  
```bash 
make docker-restart
```
2. run command : 
```bash 
go mod init gocrudsample   
```
3. run command : 
```bash 
go mod tidy
```
4. run migration with this command : 
```bash 
make migration-up
```  
5. run the application with this command : 
```bash 
make run
```

6. Using postman, open application with address : 
```bash 
localhost:7100/tutorials/
```


## How to run unit test
1. Generate mock  with : 
```bash 
make mock-gen
```
2. run command  : 
```bash 
make test
```
3. if you want to view unit test coverage in browser run comand : 
```bash 
make coverage-report 
```

## API Documentation
#### Get Detail Tutorial

```http
  GET /tutorials/${id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. Id of item to fetch in UUID Format |

#### Get Tutorial Types

```http
  GET /tutorials/types
```

#### Get Tutorial List

```http
  GET /tutorials?tutorialTypeId=${tutorial_type_id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `tutorial_type_id`      | `string` |  Id of tutorial type to fetch in UUID Format, if empty value API will result all tutorials data  |


#### Add Tutorial

```http
  POST /tutorials
```

#### Request Parameters in JSON Format
| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `tutorialTypeId`      | `string` | **Required**. Id of tutorial type id in UUID Format |
| `title`      | `string` | **Required**. tutorial title |
| `keywords`      | `string` | keywords for tutorial |
| `sequence`      | `int` | sequence for displaying tutorial in order |
| `description`      | `string` | description |

### Example
```http
{
    "tutorialTypeId":"93d899ef-b918-4a94-b7fb-c51df7c7e144",
    "title":"tutorial title",
    "sequence": 1,
    "keywords":"tutorial",
    "description" : "tutorial description"
}
```

#### Update Tutorial

```http
  UPDATE /tutorials/${id}
```

#### Request Parameters in JSON Format
| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `tutorialTypeId`      | `string` | **Required**. Id of tutorial type id in UUID Format |
| `title`      | `string` | **Required**. tutorial title |
| `keywords`      | `string` | keywords for tutorial |
| `sequence`      | `int` | sequence for displaying tutorial in order |
| `description`      | `string` | description |

### Delete Tutorial

```http
  DELETE /tutorials/${id}
```
#### Patch Tutorial

```http
  PATCH /tutorials/${id}
```

#### Request Parameters in JSON Format
| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `title`      | `string` | **Required**. tutorial title |

### Example
```http
  {
    "title":"tutorial title",
  }
```