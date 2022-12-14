# healthy_web
## Starting Server

1. Starting this web application by the following command.
    ```bash
    go run main.go
    ```
1. When startup is complete, the console shows the following message:
    ```
    http server started on [::]:8080
    ```
1. Access [http://localhost:8080](http://localhost:8080) in your browser.
1. Login with the following username and password.
    - username : ``test``
    - password : ``test``
## Build executable file
Build this source code by the following command.
```bash
go build main.go
```

## Project Map
The follwing figure is the map of this sample project.

```
  + config                  … Define configurations of this system.
  + logger                  … Provide loggers.
  + middleware              … Define custom middleware.
  + migration               … Provide database migration service for development.
  + router                  … Define routing.
  + controller              … Define controllers.
  + model                   … Define models.
  + repository              … Provide a service of database access.
  + service                 … Provide a service of book management.
  + session                 … Provide session management.
  - main.go                 … Entry Point.
```

## Services
This sample provides 3 services: meals, users, and food management.

### Meals Management
There are the following services in the Meals management.

|Service Name|HTTP Method|URL|Parameter|Summary|
|:---|:---:|:---|:---|:---|
|Get Service|GET|``/api/meal/[MEAL_ID]``|Meal ID|Get a meal data.|
|List/Search Service|GET|``/api/meal?query=[KEYWORD]&page=[PAGE_NUMBER]&size=[PAGE_SIZE]``|Page, Keyword(Optional)|Get a list of meal.|
|Regist Service|POST|``/api/meals``|Meal|Regist a meal data.|

### Users Management
There are the following services in the Users management.

|Service Name|HTTP Method|URL|Parameter|Summary|
|:---|:---:|:---|:---|:---|
|Login Service|POST|``/api/auth/login``|Session ID, User Name, Password|Session authentication with username and password.|
|Logout Service|POST|``/api/auth/logout``|Session ID|Logout a user.|
|Login Status Check Service|GET|``/api/auth/loginStatus``|Session ID|Check if the user is logged in.|
|Login Username Service|GET|``/api/auth/loginAccount``|Session ID|Get the login user's username.|

### Food Management
There are the following services in the Food management.

|Service Name|HTTP Method|URL|Parameter|Summary|
|:---|:---:|:---|:---|:---|
|Food List Service|GET|``/api/foods``|Nothing|Get a list of foods.|
