# order-api

This repository contains an incomplete implementation of a RESTful API for Orders.  Complete the tasks in `Tasks to be completed by you` section at the bottom.

* Ask as many questions about the requirements as you like
* Feel free to modify the code in any way 
* Use any resource you normally would during development
* You are encouraged to critique the code
* You are encouraged to write tests

## API design


### Create an order.

Request: 
```
POST http://host:port/orders
{
    location: "location_id",
    item: "custom item"
}
```

Response(s):
```
201 CREATED
{
    "order": {
        "id": "new order id",
        "location": "location_id",
        "item": "custom item"
    },
    "error": ""
}
```

### Get an order by ID

Request: 
```
GET http://host:port/orders/:id
```

Response(s): 
```
200 OK
{
    "order": {
        "id": "desired order id",
        "location": "location_id",
        "item": "custom item"
    },
    "error": ""
}

---
404 NOT FOUND
```

### Query for orders by location and item.

Request: 
```
GET http://host:port/orders?location=90210&item=cheese
```

Response(s): 
```
200 OK
{
    orders: [{
        "id": "new order id",
        "location": "location_id",
        "item": "desired item"
    }]
    error: "",
}

```


## Tasks to be completed by you

* Add some kind of mechanism for storing orders.  For now it is fine to store orders in memory, but in the future we may want to store them in postgres or elsewhere.  Implement with this in mind.
* Add support for retrieving an order by ID in server and client.
* Add support for querying orders by location and item.
