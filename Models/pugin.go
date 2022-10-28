package Models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/zxc10110/mvc_09/Config"
)

//create Picture ... Insert New Picture
func CreateStudent(student *Student) (err error) {
	if err = Config.DB.Create(student).Error; err != nil {
		return err
	}
	return
}

//list Student .. get all student
func GetAllStudent(student *[]Student) (err error) {
	if err = Config.DB.Find(student).Error; err != nil {
		return
	}
	return
}

//random sport
func RandomSport() (result Student, err error) {
	if err = Config.DB.Select("*").
		Table("Test.student").
		Order("RAND()").
		Limit("1").
		Find(&result).Error; err != nil {
		return
	}
	return
}

//get
func CreateSport(rand *Sport) (err error) {
	if err = Config.DB.Create(rand).Error; err != nil {
		return err
	}
	return
}

//count student
func CountStudent() (count Count, err error) {
	if err = Config.DB.Select("COUNT(student_id) as count").
		Table("Test.student").
		Find(&count).Error; err != nil {
		return
	}
	return
}

func GetSport(sport *[]Sport) (err error) {
	if err = Config.DB.Find(sport).Error; err != nil {
		return
	}
	return
}

func CountSport() (count Count, err error) {
	if err = Config.DB.Select("COUNT(sport_no) as count").
		Table("Test.sports").
		Find(&count).Error; err != nil {
		return
	}
	return
}
