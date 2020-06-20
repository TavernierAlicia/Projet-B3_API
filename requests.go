package main

//----------------- CLIENTS -----------------
//create account (date format 'yyyy-mm-dd')
const (
	verifyDouble = `SELECT mail FROM clients WHERE mail = ?`
)

const (
	createAccount = `INSERT INTO clients (name, surname, mail, password, birth_date, phone_number, token) 
										SELECT * FROM (SELECT ? AS name, ? AS surname, ? AS mail, ? AS password, ? AS birth_date, ? AS phone_number, ? AS token) as ifexists 
											WHERE NOT EXISTS (SELECT mail FROM clients WHERE mail = ? ) LIMIT 1`
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
	getAllEtabs = `SELECT id, name, description, type, latitude, longitude, main_pic, date, subtype, street_num, street_name, city, zip, happy, happy_end FROM etabs ORDER BY date DESC`
)

//get etabs by text
const (
	searchResult = `SELECT id, name, description, type, latitude, longitude, main_pic, date, subtype, street_num, street_name, city, zip, happy, happy_end FROM etabs WHERE subtype LIKE ? OR name LIKE ? OR city LIKE ? OR description LIKE ? ORDER BY date DESC`
)

//get favs
const (
	getFavs = `SELECT etabs.id, etabs.name, etabs.type, etabs.subtype, etabs.street_num, etabs.street_name, etabs.main_pic, favoris.date AS date, favs.nbFavs FROM favoris JOIN etabs ON etab_id = etabs.id LEFT JOIN (SELECT COUNT(*) AS nbFavs, etab_id FROM favoris GROUP BY etab_id) AS favs ON favs.etab_id = favoris.etab_id WHERE user_id = ? ORDER BY date DESC`
)

//add favs
const (
	addFavs = `INSERT INTO favoris(user_id, etab_id) SELECT * FROM (SELECT ? AS user_id, ? AS etab_id) AS ifexists WHERE NOT EXISTS (SELECT * FROM favoris WHERE user_id = ? AND etab_id = ?) LIMIT 1`
)

//delete favs
const (
	deleteFav = `DELETE FROM favoris WHERE user_id = ? AND etab_id = ?`
)

//get user details
const (
	getUser = `SELECT id, name, surname, mail, phone_number AS phone, profile_pic, birth_date FROM clients WHERE id = ?`
)

//change user data
const (
	editUserCm = `UPDATE clients SET name = ?, surname = ?, birth_date = ?, phone_number = ?, profile_pic = ? WHERE id = ?`
)

//change user pwd
const (
	editUserPwd = `UPDATE clients SET password = ?, token = ? WHERE id = ?`
)

//show bar
const (
	showBarDetails = `SELECT IF(favoris.id IS NULL, 0, 1) AS is_fav, etabs.id, name, description, street_num, street_name, address_complement, city, zip, type, subtype, main_pic, etabs.date, happy, happy_end, IFNULL(tempFav.favNum, 0) AS fav FROM etabs LEFT JOIN (SELECT IFNULL(COUNT(user_id), 0) AS favNum, etab_id FROM favoris GROUP BY etab_id) AS tempFav ON tempFav.etab_id = etabs.id LEFT JOIN favoris ON favoris.etab_id = etabs.id AND user_id = ? WHERE etabs.id = ?`
)

const (
	showBarPictures = `SELECT id, etab_id, path FROM etab_pictures WHERE etab_id = ?`
)

const (
	showBarItems = `SELECT i.id, i.name, i.description, i.price, IFNULL(i.sale, 1) AS sale , (i.price * IFNULL(i.sale, 1)) AS newprice, i.type FROM items AS i WHERE i.etab_id = ?`
)

//take order
const (
	addOrder = `INSERT INTO commands (client_id, etab_id, instructions, waiting_time, payment, tip) VALUES (?, ?, ?, ?, ?, ?)`
)

//insert order
const (
	addOrderItems = `INSERT INTO command_items (command_id, item_id, price) VALUES (?, ?, (SELECT price * IFNULL(sale, 1) AS price FROM items WHERE id = ?))
	`
)

//calc order price
const (
	calcPrice = `UPDATE commands SET price = (SELECT IFNULL(SUM(price),0) FROM command_items WHERE command_id = ?) WHERE id = ?`
)

//show orders
const (
	showOrders = `SELECT IFNULL(commands.price, 0) AS totalprice, commands.etab_id, status, cmd_date, commands.id, name AS etab_name, main_pic FROM commands JOIN etabs ON commands.etab_id = etabs.id WHERE client_id = ? ORDER BY cmd_date DESC LIMIT 30`
)

//show orders details
const (
	showOrdersDetails = `SELECT command_id, command_items.item_id, COUNT(item_id) AS quantity, items.name, command_items.price FROM command_items JOIN items ON command_items.item_id = items.id WHERE command_id = ? GROUP BY item_id`
)

//show one order detail
const (
	showTheOrder = `SELECT etabs.street_num, etabs.street_name, etabs.city, IFNULL(commands.price, 0) AS totalprice, status, cmd_date, commands.id, commands.tip, commands.payment, name AS etab_name, commands.etab_id, main_pic FROM commands JOIN etabs ON commands.etab_id = etabs.id WHERE commands.id = ?`
)

//----------------- PROS -----------------
//GET
//POST
//PUT
//DELETE
