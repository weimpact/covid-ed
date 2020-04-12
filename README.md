# Covid Ed

[![Run in Postman](https://run.pstmn.io/button.svg)](https://app.getpostman.com/run-collection/9fa9cf5d968cb5850c51)

Project which aims to educate users with information about corona and it's seriousness in local language.


`GET /countries/{country}/cases ? top=10`
* top countries with high total cases

```
{
  "timezone": "IST", [ {
    "date": "dateformat",
    "count": 33,
  },
  {
    "date": "dateformat",
    "count": 30,
  }
  ]
}
```

`GET /countries/{country}/cases/aggreggated ? countries=IN,SG & interval=daily`
* to get countries cases aggregated over an interval

`GET /facts_myths`
* Get list of facts and myths

```
[
    {
        "fact": {
            "id": 1,
            "title": "Mask only prevents spreading from you",
            "description": "If you wear mask and sneeze you will not spread to others"
        },
        "myth": {
            "id": 0,
            "title": "Face masks protect against coronavirus",
            "description": "if you wear face mask then you wont get corona"
        }
    }
] 
```
`pkg/client` is built on `https://api.covid19api.com/`
