package main

var locationRoutes = Routes{
	Route{
		"LocationCreate",
		"POST",
		"/locations",
		LocationCreate,
	},
    Route{
         "LocationUpdate",
         "POST",
         "/locations/{locationId}",
         LocationUpdate,
    },
	Route{
		"LocationDelete",
		"DELETE",
		"/locations/{locationId}",
		LocationDelete,
	},
    Route{
        "LocationList",
        "GET",
        "/locations",
        LocationList,
    },
    Route{
        "LocationGet",
        "GET",
        "/locations/{locationId}",
        LocationGet,
    },

}
