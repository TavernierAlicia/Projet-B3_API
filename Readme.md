# *** API Order'NDrink *** 

## App client 


### GET
#### Show all bars with or without filters -> "/app/show"
  ##### Params 
  *Can be Null*
  - type (Biere, Vin... if null set default = all)
  - distance (km)
  - popularity (fav or new, if null set default = all)
  - lat (user geolocalisation)
  - long (user geolocalisation)

  ##### Headers
  - Authorization = access token

  ##### Body
  *Null*

#### Show searchbar results -> "app/search"
  ##### Params 
  *Can be Null*
  - search (Paris, vin, artisanal, rue de..)
  
  ##### Headers
  - Authorization = access token

  ##### Body
  *Null*

#### Get user profile -> "app/profile"
  ##### Params 
  *Null*
  
  ##### Headers
  - Authorization = access token

  ##### Body
  *Null*

#### Show details and menu for 1 etab -> "app/show/:id"
  ##### Params //bar id in path
  *Null*
  
  ##### Headers
  - Authorization = access token

  ##### Body
  *Null*

#### Show orders -> "app/showOrders"
  ##### Params 
  *Null*
  
  ##### Headers
  - Authorization = access token

  ##### Body
  *Null*

#### Get details for 1 order -> "app/getOrder/:comandid"
  ##### Params //command id in path 
  *Null*
  
  ##### Headers
  - Authorization = access token

  ##### Body
  *Null*


### POST
#### Create account -> "app/createUser/"
  ##### Params 
  *Null*
  
  ##### Headers
  *Null*

  ##### Body raw JSON
  {
	"name": string,
	"surname": string,
	"mail": string,
	"pass": string,
	"birth": string YYYY-MM-DD,
	"phone": string LIMIT 10,
	"confirmPass": string
}

#### Connect to account -> "app/auth/"
  ##### Params 
  *Null*
  
  ##### Headers
  *Null*

  ##### Body raw JSON
  {
	  "mail": string,
	  "pass": string
  }


#### Add Fav -> "app/favs/add/:etabid"
  ##### Params //bar id in path
  *Null*
  
  ##### Headers
  - Authorization = access token

  ##### Body 
  *Null*


#### Take Order -> "app/takeOrder"
  ##### Params
  *Null*
  
  ##### Headers
  - Authorization = access token

  ##### Body raw JSON
  {
    "etab_id": int,
    "instructions": string,
    "waiting_time": string HH:MM,
    "payment": string,
    "tip": int,
    "items_id": [int, int, int]
  }

### PUT
#### Edit profile -> "app/profile/edit/"
  ##### Params
  *Null*
  
  ##### Headers
  - Authorization = access token

  ##### Body raw JSON !! Password required if newPassword not empty !!
  {
	"name": string,
	"surname": string,
	"birth": string YYYY-MM-DD,
	"mail": string,
	"phone": string LIMIT 10,
	"pass": string,
	"newPass": string,
	"pic": string
  }

### DELETE 
#### Delete Fav -> "app/favs/delete/:etabid"
  ##### Params //bar id in path 
  *Null*
  
  ##### Headers
  - Authorization = access token

  ##### Body
  *Null*




  ### ERRORS
  0 --> No error
  1 --> This route doesn't exists
  2 --> Login or password wrong
  3 --> Mail already exists
  4 --> Mismatch passwords
  5 --> An error occured
  6 --> Field(s) missing
  7 --> No password
  8 --> Same passwords
  9 --> Incorrect password
