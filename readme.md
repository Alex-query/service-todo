## Create Todo

Request for creating todo

```http
POST /api/todo
```

| Parameter | Type | Description |
| :--- | :--- | :--- |
| `title` | `string` | **Required**. Title todo |

Example

```javascript
{
    "title":string,
}
```

### Responses

Endpoints return the JSON representation of the resources created or edited. However, if an invalid request is submitted, or some other error occurs, API returns a JSON response in the following format:

```javascript
{
    "data": {
        "id": string,
        "title": string
    },
    "meta": {
        "status": "success"
    }
}
```

## Get Todo

Request for creating todo

```http
GET /api/todo
```

| Parameter | Type | Description |
| :--- | :--- | :--- |
| `id` | `string` | **Required**. Title id |

Example

```javascript
{
    "id":string,
}
```

### Responses

Endpoints return the JSON representation of the resources created or edited. However, if an invalid request is submitted, or some other error occurs, API returns a JSON response in the following format:

```javascript
{
    "data": {
        "id": string,
        "title": string
    },
    "meta": {
        "status": "success"
    }
}
```

```javascript
{
    "data": "not found",
    "meta": {
        "status": "error"
    }
}
```