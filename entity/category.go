package entity

type SubCategories struct {
	Id        string `json:"id" gorm:"primary_key;auto_increment"`
	Type_name string `json:"type_name`
	Cat_id    string `json:"cat_id"`
}
