*** API Order'NDrink ***

*App "Client"*

** GET **
# Show all bars with or without filters -> "/app/show"
  ## Params 
  *Can be Null*
  - type (Biere, Vin... if null set default = all)
  - distance (km)
  - popularity (fav or new, if null set default = all)
  - lat (user geolocalisation)
  - long (user geolocalisation)

  ## Headers
  - Authorization = access token

  ## Body
  *Null*

# Show searchbar results -> "app/search"
  ## Params 
  *Can be Null*
  - search (Paris, vin, artisanal, rue de..)
  
  ## Headers
  - Authorization = access token

  ## Body
  *Null*

# Get user profile -> "app/profile"
  ## Params 
  *Null*
  
  ## Headers
  - Authorization = access token

  ## Body
  *Null*

# Show details and menu for 1 etab -> "app/show/:id"
  ## Params //bar id in path
  *Null*
  
  ## Headers
  - Authorization = access token

  ## Body
  *Null*

# Show orders -> "app/showOrders"
  ## Params 
  *Null*
  
  ## Headers
  - Authorization = access token

  ## Body
  *Null*

# Get details for 1 order -> "app/getOrder/:comandid"
  ## Params //command id in path 
  *Null*
  
  ## Headers
  - Authorization = access token

  ## Body
  *Null*


** POST **
# Create account -> "app/createUser/"
  ## Params 
  *Null*
  
  ## Headers
  *Null*

  ## Body x-www-form-urlencoded
  - name
  - surname
  - mail
  - password
  - confirmPassword
  - birth (yyyy-mm-dd)
  - phone (0102030405)

# Connect to account -> "app/auth/"
  ## Params 
  *Null*
  
  ## Headers
  *Null*

  ## Body x-www-form-urlencoded
  - mail
  - password


# Add Fav -> "app/favs/add/:etabid"
  ## Params //bar id in path
  *Null*
  
  ## Headers
  - Authorization = access token

  ## Body 
  *Null*


# Take Order -> "app/takeOrder"
  ## Params
  *Null*
  
  ## Headers
  - Authorization = access token

  ## Body raw JSON
  {
    "etab_id": int,
    "instructions": string,
    "waiting_time": string HH:MM,
    "payment": string,
    "tip": int,
    "items_id": [int, int, int]
  }

** PUT **
# Edit profile -> "app/profile/edit/"
  ## Params
  *Null*
  
  ## Headers
  - Authorization = access token

  ## Body x-www-form-urlencoded !! Password required if newPassword not empty !!
  - name
  - surname
  - pic
  - mail
  - password
  - newPassword
  - birth (yyyy-mm-dd)
  - phone (0102030405)

** DELETE **
# Delete Fav -> "app/favs/delete/:etabid"
  ## Params //bar id in path 
  *Null*
  
  ## Headers
  - Authorization = access token

  ## Body
  *Null*