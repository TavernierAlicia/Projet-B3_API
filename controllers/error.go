package main
//handle custom error page
func returnError(c *gin.Context) {
	var cause string

	cause = "Nous sommes desoles, une erreur s'est produite"

	//display error page
	data := Error{
		ErrorCause: cause,
	}
	c.JSON(404, data)
}