package main

var itemRoutes = Routes{
	Route{
		"ItemCreate",
		"POST",
		"/items",
		ItemCreate,
	},
    Route{
         "ItemUpdate",
         "POST",
         "/items/{itemId}",
         ItemUpdate,
    },
	Route{
		"ItemDelete",
		"DELETE",
		"/items/{itemId}",
		ItemDelete,
	},
    Route{
        "ItemList",
        "GET",
        "/items",
        ItemList,
    },
    Route{
        "ItemGet",
        "GET",
        "/items/{itemId}",
        ItemGet,
    },

}
