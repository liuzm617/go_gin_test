package models

import "log"

type Tag struct {
	Model

	Name       string `gorm:"type:varchar(100);unique_index" json:"name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}



func GetTags() ([]Tag, error) {
	var (
		tags []Tag
		err  error
	)
	// err = db.Where("").Find(&tags).Error
	query := make(map[string]interface{})
	query["id"] = 1
	err = db.Where(query).Find(&tags).Error
	return tags, err
}

func GetTag() []Tag {
	var (
		tag Tag
		err error
	)

	err = db.Where("id>=?", 1).Find(&tag).Error
	if err != nil {
		log.Printf("err:%s\n", err)
	}
	return append(make([]Tag, 0), tag)
}


func AddTag(name string, createdBy string)(int, error){
	tag := Tag{
		Name:name,
		CreatedBy:createdBy,
		State:0,
	}
	if err := db.Create(&tag).Error; err!= nil{
		return 0, err
	}
	return tag.ID ,nil

}

func DeleteTag(id int) error{
	var tag Tag
	err := db.Where("id=?", id).Delete(&tag).Error
	return err
}