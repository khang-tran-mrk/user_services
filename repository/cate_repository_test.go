package repository

import (
	"ChoTot/config"
	"ChoTot/entity"
	"log"
	"testing"
)

func Test_cateRepository_Create(t *testing.T) {
	info := &entity.SubCategories{
		Type_name: "test cate",
		Cat_id:    "1",
	}

	db := config.ConnectDatabase()
	cr := NewCateRepository(db)
	if err := cr.Create(info); err != nil {
		log.Fatalln(err)
	}
}
