## How to run this project

---

1. import startup_db.sql
2. edit configure environment in config.env
3. run the project

# Build project

```
go build -o main
```

# Run app

```
go run main.go
```

# Api Spec

### GROUP: Authentication

- [1] - Register
- [POST] : {root.api}/api/{version}/users

```json
Request:
    {
        "name": "Adam",
        "occupation" : "Programmer",
        "email" : "email3@gmail.com",
        "password" : "password"
    }

Response:
    {
        "meta": {
            "message": "Account has ben registered",
            "code": 200,
            "status": "success"
        },
        "data": {
            "id": 17,
            "name": "Adam",
            "occupation": "Programmer",
            "email": "email3@gmail.com",
            "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxN30.09wYEC-GsUwGd_zj5DXDndcagccfyAuISEzG6swq9Ig",
            "image_url": ""
        }
    }
```

- [2] - Login
- [POST] : {root.api}/api/{version}/sessions

```json
Request:
    {
        "email" : "email3@gmail.com",
        "password" : "password"
    }

Response:
{
    "meta": {
        "message": "Login Successfully",
        "code": 200,
        "status": "success"
    },
    "data": {
        "id": 17,
        "name": "Adam",
        "occupation": "Programmer",
        "email": "email3@gmail.com",
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxN30.09wYEC-GsUwGd_zj5DXDndcagccfyAuISEzG6swq9Ig",
        "image_url": ""
    }
}
```

### GROUP: User

- [1] - Upload photo
- [POST] : {root.api}/api/{version}/users

```
Headers :

    - key : Authorization
    - Value : Bearer {token}

Request (from-data):

    - key : avatar{type file}  Value : file name
```

```json
Response:
{
    "meta": {
        "message": "Avatar successfuly uploaded",
        "code": 200,
        "status": "success"
    },
    "data": {
        "is_uploaded": true
    }
}
```

- [2] - Check Email Availability
- [POST] : {root.api}/api/{version}/email_checkers

```json
Request:
{
    "email" : "email@gmail.com"
}
Response:
{
    "meta": {
        "message": "Email is available",
        "code": 200,
        "status": "success"
    },
    "data": {
        "is_available": true
    }
}
```

Email registered

```json
Response:
{
    "meta": {
        "message": "Email has been registered",
        "code": 200,
        "status": "success"
    },
    "data": {
        "is_available": false
    }
}

```

- [3] - fetch user
- [GET] : {root.api}/api/{version}/email_checkers

```json
Headers :

    - key : Authorization
    - Value : Bearer {token}

Response:
{
    "meta": {
        "message": "Successfuly fetch user data",
        "code": 200,
        "status": "success"
    },
    "data": {
        "id": 17,
        "name": "Adam",
        "occupation": "Programmer",
        "email": "email3@gmail.com",
        "token": "",
        "image_url": "images/17-undraw_profile_pic_ic5t.png"
    }
}
```

### GROUP: Campaign

- [1] - List campaigns
- [GET] : {root.api}/api/{version}/campaigns

```json
params :(can be empty)
    - key : user_id
    - Value : int


Response:
{
    "meta": {
        "message": "List of campaigns",
        "code": 200,
        "status": "success"
    },
    "data": [
        {
            "id": 1,
            "user_id": 2,
            "name": "Campaign punya Adam nasrudin",
            "short_description": "Cras faucibus magna vel blandit ultrices. Integer vitae aliquam nisl.",
            "image_url": "images/campaign-1-undraw_profile_pic_ic5t.png",
            "goal_amount": 200000,
            "current_amount": 0,
            "slug": "campaign-punya-user-1"
        },
        {
            "id": 2,
            "user_id": 3,
            "name": "project name",
            "short_description": "project ini adalah",
            "image_url": "",
            "goal_amount": 500000,
            "current_amount": 350000,
            "slug": "project-by-fahri-16"
        }
    ]
}
```

- [2] - Create campaigns
- [POST] : {root.api}/api/{version}/campaigns

```json
Headers :

    - key : Authorization
    - Value : Bearer {token}

Request:
{
    "name": "campaign name sangat SEDERHANA",
    "short_description": "short_description",
    "description": "Penjelasan secara lebar",
    "goal_amount": 10000000,
    "perks": "Keuntungan  1, keuntungan 2 , perks 3, ini ke empat"
}

Response:
{
    "meta": {
        "message": "Success to create campaign",
        "code": 200,
        "status": "success"
    },
    "data": {
        "id": 8,
        "user_id": 1,
        "name": "campaign name sangat SEDERHANA",
        "short_description": "short_description",
        "image_url": "",
        "goal_amount": 10000000,
        "current_amount": 0,
        "slug": "campaign-name-sangat-sederhana-1"
    }
}
```

- [3] - upload campaign images
- [POST] : {root.api}/api/{version}/campaign-images

```json
Headers :

    - key : Authorization
    - Value : Bearer {token}

Request (from-data):

    - key : avatar{type file}  Value : file name
    - key : is_primary{type text}  Value : true/false
    - key : campaign_id{type text}  Value : int
Response:
{
    "meta": {
        "message": "Campaign image successfuly uploaded",
        "code": 200,
        "status": "success"
    },
    "data": {
        "is_uploaded": true
    }
}
```

- [4] - Update campaigns
- [PUT] : {root.api}/api/{version}/campaigns/:id

```json
Headers :

    - key : Authorization
    - Value : Bearer {token}

Request:
{
    "name": "Update campaign  ",
    "short_description": "short_ description ",
    "description": "Penjelasan secara lebar ",
    "goal_amount": 15000000,
    "perks": "Keuntungan  1 Update, keuntungan 2 , perks 3 Update,  ini ke empat Update"
}

Response:
{
    "meta": {
        "message": "Success to updated campaign",
        "code": 200,
        "status": "success"
    },
    "data": {
        "id": 8,
        "user_id": 1,
        "name": "Update campaign  ",
        "short_description": "short_ description ",
        "image_url": "images/campaign-1-undraw_profile_pic_ic5t.png",
        "goal_amount": 15000000,
        "current_amount": 0,
        "slug": "campaign-name-sangat-sederhana-1"
    }
}
```

- [5] - Get campaigns Detail
- [GET] : {root.api}/api/{version}/campaigns/:id

```json
Response:
{
    "meta": {
        "message": "Campaign detail",
        "code": 200,
        "status": "success"
    },
    "data": {
        "id": 8,
        "name": "Update campaign  ",
        "short_description": "short_ description ",
        "description": "Penjelasan secara lebar ",
        "image_url": "images/campaign-1-undraw_profile_pic_ic5t.png",
        "goal_amount": 15000000,
        "current_amount": 0,
        "backer_count": 0,
        "user_id": 1,
        "slug": "campaign-name-sangat-sederhana-1",
        "perks": [
            "Keuntungan  1 Update",
            "keuntungan 2",
            "perks 3 Update",
            "ini ke empat Update"
        ],
        "user": {
            "name": "Adam",
            "image_url": "images/1-undraw_profile_pic_ic5t.png"
        },
        "images": [
            {
                "image_url": "images/campaign-1-undraw_profile_pic_ic5t.png",
                "is_primary": true
            }
        ]
    }
}
```

### GROUP: Transaction

- [1] - Create Transaction
- [POST] : {root.api}/api/{version}/transactions

```json
Headers :

    - key : Authorization
    - Value : Bearer {token}

Request:
{
    "campaign_id": 8,
    "amount": 350000
}

Response:
{
    "meta": {
        "message": "Create transaction successfully",
        "code": 200,
        "status": "success"
    },
    "data": {
        "id": 23,
        "campaign_id": 8,
        "user_id": 17,
        "amount": 350000,
        "status": "pending",
        "code": "",
        "payment_url": "https://app.sandbox.midtrans.com/snap/v2/vtweb/c63ac421-3ae7-4489-befa-597dec644f84"
    }
}
```

- [2] - List Transaction by user
- [GET] : {root.api}/api/{version}/transactions

```json
Headers :

    - key : Authorization
    - Value : Bearer {token}

Response:
{
    "meta": {
        "message": "List of users transaction",
        "code": 200,
        "status": "success"
    },
    "data": [
        {
            "id": 23,
            "amount": 350000,
            "status": "pending",
            "created_at": "2021-03-04T14:43:01+07:00",
            "campaign": {
                "name": "Update campaign  ",
                "image_url": "images/campaign-1-undraw_profile_pic_ic5t.png"
            }
        },
        {
            "id": 22,
            "amount": 350000,
            "status": "pending",
            "created_at": "2021-03-04T14:41:15+07:00",
            "campaign": {
                "name": "Update campaign  ",
                "image_url": "images/campaign-1-undraw_profile_pic_ic5t.png"
            }
        }
    ]
}
```

- [3] - List Transaction by campaigns id
- [GET] : {root.api}/api/{version}/campaigns/:id/transactions

```json
Headers :

    - key : Authorization
    - Value : Bearer {token}

Response:
{
    "meta": {
        "message": "List of Campaigns transaction",
        "code": 200,
        "status": "success"
    },
    "data": [
        {
            "id": 15,
            "name": "Adam Nasrudin",
            "amount": 12345678,
            "created_at": "2021-03-02T13:56:19+07:00"
        }
    ]
}
```

## Web CMS {Dashboard admin}

```
url : {root}/login
```
