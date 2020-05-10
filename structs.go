package main

//Error struct
type Error struct {
	Text string
}

//get bars
type Bars struct {
	Id          int64   `db:"id"`
	Name        string  `db:"name"`
	Description string  `db:"description"`
	Type        string  `db:"type"`
	Subtype     string  `db:"subtype"`
	Lat         float64 `db:"latitude"`
	Long        float64 `db:"longitude"`
	Pic         string  `db:"main_pic"`
	Date        string  `db:"date"`
}

type User struct {
	Id       int64  `db:"id"`
	Name     string `db:"name"`
	Surname  string `db:"surname"`
	Mail     string `db:"mail"`
	Pic      string `db:"profile_pic"`
	Birth    string `db:"birth_date"`
	Password string `db:"password"`
}

type BarView struct {
	BarDetails []*barDetails
	Pictures   []*picture
	Items      []*item
}

type barDetails struct {
	Id          int64  `db:"id"`
	Name        string `db:"name"`
	Description string `db:"description"`
	Favs        int64  `db:"fav"`
	Street_num  int64  `db:"street_num"`
	Street_name string `db:"street_name"`
	Complement  string `db:"address_complement"`
	City        string `db:"city"`
	Zip         string `db:"zip"`
	Type        string `db:"type"`
	Subtype     string `db:"subtype"`
	Pic         string `db:"main_pic"`
	Date        string `db:"date"`
	Open        string "23:59"
	HH          string `db:"happy"`
	HHEnd       string `db:"happy_end"`
}

type picture struct {
	Id      int64  `db:"id"`
	Etab_id int64  `db:"etab_id"`
	Path    string `db:"path"`
}

type item struct {
	Id          int64   `db:"id"`
	Name        string  `db:"name"`
	Description string  `db:"description"`
	Price       float64 `db:"price"`
	Sale        float64 `db:"sale"`
	NewPrice    float64 `db:"newprice"`
	Quantity    int64   `db:"quantity"`
	Type        string  `db:"type"`
}
