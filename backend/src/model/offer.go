package model

import "time"

type Offer struct {
	Id       string    `json:"id"`
	Vehicle  Vehicle   `json:"vehicle"`
	Price    int       `json:"price"`
	Date     time.Time `json:"date"` // publish date
	Location string    `json:"location"`
	Rates    []Rate    `json:"rates"`
	Comments []Comment `json:"comments"`
}

/*
[
    {
        "id": "ad2dbbb5-e13e-4fab-8c56-0ee3a1a0f62d",
        "vehicle": {
            "id": "1",
            "make": "BMW",
            "model": "320",
            "date": "2020-05-05T00:00:00Z",
            "hp": 150,
            "cubic": 2000
        },
        "price": 20000,
        "date": "2020-05-05T00:00:00Z",
        "location": "Novi Sad",
        "rates": [
            {
                "id": "551414fe-c8a6-4b3a-8387-a3f3c7f3cc42",
                "offerId": "ad2dbbb5-e13e-4fab-8c56-0ee3a1a0f62d",
                "mark": 3
            }
        ],
        "comments": [
            {
                "id": "420b9b22-83e1-4716-b5ff-ac5d2ea6828a",
                "offerId": "ad2dbbb5-e13e-4fab-8c56-0ee3a1a0f62d",
                "string": "com 1"
            }
        ]
    }
]
*/
