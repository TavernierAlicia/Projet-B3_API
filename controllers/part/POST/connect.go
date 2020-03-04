package main


func connectPart(c *gin.Context) {

	request_name := "connect_part"

	requestDB(client_id, request_name)
	isConnected := true

	data := PartConnect{
		user_id: id,
		name: name,
		surname: surname,
		connected: isConnected,
	}
	c.JSON(200, data)
}

