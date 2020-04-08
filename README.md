# Covid Ed

Project which aims to educate users with information about corona and it's seriousness in local language.


Req: `GET /countries/{country}/cases ? top=10

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

similarly for `GET /countries/{country}/cases/aggreggated ? countries=IN,SG & interval=daily
```

```



`pkg/client` is built on `https://api.covid19api.com/`
