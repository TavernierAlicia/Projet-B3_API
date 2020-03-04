package main

//// CONTROLLERS EXAMPLE ////
/*
//handle index page
func indexPage(c *gin.Context) {
	data := Preview{}
	c.HTML(200, "index.html", data)
}


//handle form pro
func formProPage(c *gin.Context) {
	data := Preview{}
	c.HTML(200, "form-pro.html", data)
}

func receptForm(c *gin.Context) {

	data := Preview{}

	fmt.Println(data)
	c.Request.ParseForm()

	mail := strings.Join(c.Request.PostForm["from"], " ")
	name := strings.Join(c.Request.PostForm["name"], " ")
	surname := strings.Join(c.Request.PostForm["surname"], " ")
	subjectNum := strings.Join(c.Request.PostForm["subject"], " ")
	cmdNumber := strings.Join(c.Request.PostForm["cmdNumber"], " ")
	message := strings.Join(c.Request.PostForm["message"], " ")
	pro := false

	path := c.FullPath()

	//Define is the client is a professionnal or not
	if path == "/contact/form-pro" || path == "/professionnal/form-pro" || path == "/form-pro" {
		pro = true
	} else {
		pro = false
	}

	//choose subject and send mail
	subject := SelectSubj(pro, subjectNum)
	response := SendMail(mail, name, surname, subject, cmdNumber, message, pro)

	if response == true {
		c.Redirect(http.StatusMovedPermanently, "/success")
	} else {
		c.Redirect(http.StatusMovedPermanently, "/error")
	}
}
*/