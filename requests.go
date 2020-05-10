package main

//create account (date format 'yyyy-mm-dd')
const (
	createAccount = `INSERT INTO clients (name, surname, mail, password, birth_date, phone_number, token) VALUES (?, ?, ?, ?, ?, ?, ?)`
)

//authentification
const (
	authReq = `SELECT token FROM clients WHERE mail = ? AND password = ?`
)

//get user id
const (
	getUID = `SELECT id FROM clients WHERE token = ?`
)

//get all etabs
const (
	getAllEtabs = `SELECT id, name, description, type, latitude, longitude, main_pic, date, subtype FROM etabs ORDER BY date DESC`
)

//get etabs by text
const (
	searchResult = `SELECT id, name, description, type, latitude, longitude, main_pic, date, subtype FROM etabs WHERE subtype LIKE ? OR name LIKE ? OR city LIKE ? OR description LIKE ? ORDER BY date DESC`
)

//get favs
const (
	getFavs = `SELECT etabs.id, etabs.name, etabs.description, etabs.type, etabs.latitude, etabs.longitude, etabs.main_pic, favoris.date AS date, etabs.subtype FROM favoris JOIN etabs ON etab_id = etabs.id WHERE user_id = ? ORDER BY date DESC`
)

//get user details
const (
	getUser = `SELECT id, name, surname, mail, profile_pic, birth_date, password FROM clients WHERE id = ?`
)

//change user data
const (
	editUserCm = `UPDATE clients SET name = ?, surname = ?, birth_date = ?, mail = ? WHERE id = ?`
)

//change user pwd
const (
	editUserPwd = `UPDATE clients SET password = ?, token = ? WHERE id = ?`
)

//show bar
const (
	showBarDetails = `SELECT id, name, description, street_num, street_name, address_complement, city, zip, type, subtype, main_pic, date, happy, happy_end, tempFav.favNum AS fav FROM etabs LEFT JOIN (SELECT COUNT(user_id) AS favNum, etab_id FROM favoris GROUP BY etab_id) AS tempFav ON tempFav.etab_id = id WHERE id = ?`
)

const (
	showBarPictures = `SELECT id, etab_id, path FROM etab_pictures WHERE etab_id = ?`
)

const (
	showBarItems = `SELECT i.id, i.name, i.description, i.price, IFNULL(i.sale, 1) AS sale , (i.price * IFNULL(i.sale, 1)) AS newprice, i.type, IFNULL(ci.quantity, 0) AS quantity FROM items AS i 
										LEFT JOIN 
											(SELECT count(cart_items.id) AS quantity, cart_items.item_id FROM carts 
										LEFT JOIN cart_items ON cart_items.cart_id = carts.id 
											WHERE carts.etab_id = ? AND carts.client_id = ? AND carts.status = "current" 
											GROUP BY cart_items.item_id) AS ci ON ci.item_id = i.id 
									WHERE i.etab_id = ?`
)


//----------------- PROS -----------------
//GET
//POST
//PUT
//DELETE

*/
