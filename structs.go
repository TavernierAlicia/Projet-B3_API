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
	StreetNum   int64   `db:"street_num"`
	StreetName  string  `db:"street_name"`
	City        string  `db:"city"`
	Zip         string  `db:"zip"`

	Happy    string `db:"happy"`
	HappyEnd string `db:"happy_end"`
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
	IsFav       bool   `db:"is_fav"`
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
	Type        string  `db:"type"`
}

//get bars
type BarsInFavs struct {
	Id          int64  `db:"id"`
	Name        string `db:"name"`
	Type        string `db:"type"`
	Subtype     string `db:"subtype"`
	Street_num  string `db:"street_num"`
	Street_name string `db:"street_name"`
	Pic         string `db:"main_pic"`
	DateAdded   string `db:"date"`
	NbFavs      int64  `db:"nbFavs"`
}

type TakeOrder struct {
	Etab_id      int64   `json:"etab_id"`
	Instructions string  `json:"instructions"`
	Waiting_time string  `json:"waiting_time"`
	Payment      string  `json:"payment"`
	Tip          int64   `json:"tip"`
	Items        []int64 `json:"items_id"`
}

type OrderItems struct {
	Etab_id  int64   `db:"etab_id"`
	Id       int64   `db:"id"`
	Name     string  `db:"name"`
	Price    float64 `db:"price"`
	Sale     float64 `db:"sale"`
	NewPrice float64 `db:"newprice"`
}

type Command struct {
	Id        int64   `db:"id"`
	Etab_id   int64   `db:"etab_id"`
	Etab_name string  `db:"etab_name"`
	Pic       string  `db:"main_pic"`
	Date      string  `db:"cmd_date"`
	Price     float64 `db:"totalprice"`
	Status    string  `db:"status"`
}

type CommandItems struct {
	CommandId int64   `db:"command_id"`
	Item_id   int64   `db:"item_id"`
	Quantity  int64   `db:"quantity"`
	Name      string  `db:"name"`
	Price     float64 `db:"price"`
}

type Commands struct {
	Cmd      Command
	CmdItems []*CommandItems
}

type Auth struct {
	Mail string `json:"mail"`
	Pass string `json:"pass"`
}

type UserEdit struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Birth   string `json:"birth"`
	Phone   string `json:"phone"`
	Mail    string `json:"mail"`
	Pass    string `json:"pass"`
	NewPass string `json:"newPass"`
	Pic     string `json:"pic"`
}

type UserCreate struct {
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	Pass        string `json:"pass"`
	Mail        string `json:"mail"`
	Birth       string `json:"birth"`
	Phone       string `json:"phone"`
	ConfirmPass string `json:"confirmPass"`
}
