package postgres

import (
	"fmt"
	"os"
	"strconv"

	"github.com/beclab/terminus-app-demo/pkg/model"
	"github.com/gofiber/fiber/v2/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PGOperator struct {
	db *gorm.DB
}

var (
	db_host     = ""
	db_port     = 0
	db_username = ""
	db_password = ""
	db_name     = ""
)

func init() {
	var err error

	db_host = os.Getenv("DB_HOST")
	db_port, err = strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		panic(err)
	}
	db_username = os.Getenv("DB_USER")
	db_password = os.Getenv("DB_PWD")
	db_name = os.Getenv("DB_NAME")
}

func New() *PGOperator {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
		db_host, db_username, db_password, db_name, db_port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&model.Note{})
	if err != nil {
		panic(err)
	}

	return &PGOperator{db: db}
}

func (o *PGOperator) Create(data *model.Note) (*model.Note, error) {
	err := o.db.Create(&data).Error
	if err != nil {
		log.Error("exec sql error, ", err)
		return nil, err
	}

	return data, nil
}

func (o *PGOperator) Get(id int) (*model.Note, error) {
	var note *model.Note
	err := o.db.Where("id = ?", id).First(&note).Error
	if err != nil {
		log.Error("exec sql error, ", err)
		return nil, err
	}

	return note, nil
}

func (o *PGOperator) List() ([]*model.Note, error) {
	list := make([]*model.Note, 0)
	err := o.db.Order("id desc").Find(&list).Error
	if err != nil {
		log.Error("exec sql error, ", err)
		return nil, err
	}

	return list, nil
}

func (o *PGOperator) Delete(id int) error {
	err := o.db.Where("id = ?", id).Delete(&model.Note{}).Error
	if err != nil {
		log.Error("exec sql error, ", err)
		return err
	}

	return nil
}

func (o *PGOperator) Update(data *model.Note) (*model.Note, error) {
	err := o.db.Save(data).Error
	if err != nil {
		log.Error("exec sql error, ", err)
		return nil, err
	}

	return data, nil
}

func (o *PGOperator) Close() {

}
