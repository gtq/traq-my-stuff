package main

var containerRoutes = Routes{
	Route{
		"ContainerCreate",
		"POST",
		"/containers",
		ContainerCreate,
	},
    Route{
         "ContainerUpdate",
         "POST",
         "/containers/{containerId}",
         ContainerUpdate,
    },
	Route{
		"ContainerDelete",
		"DELETE",
		"/containers/{containerId}",
		ContainerDelete,
	},
    Route{
        "ContainerList",
        "GET",
        "/containers",
        ContainerList,
    },
    Route{
        "ContainerGet",
        "GET",
        "/containers/{containerId}",
        ContainerGet,
    },

}
