# Covid Ed

Project which aims to educate users with information about corona and it's seriousness in local language.


Req: `GET /country/{country}/cases ? count=daily|total`

Response: 
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

similarly for `GET /country/{country}/state/{state}/cases ? count=daily|total`


Req: `Get /countries/cases ? top=N`
