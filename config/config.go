package config

import (
	// ex_rep "eco_points/internal/features/exchanges/repository"
	// f_rep "eco_points/internal/features/feedbacks/repository"
	// l_rep "eco_points/internal/features/locations/repository"
	// r_rep "eco_points/internal/features/rewards/repository"
	// t_rep "eco_points/internal/features/trashes/repository"
	// u_rep "eco_points/internal/features/users/repository"
	// d_rep "eco_points/internal/features/waste_deposits/repository"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type setting struct {
	User        string
	Host        string
	Password    string
	Port        string
	DBName      string
	JWTSecret   string
	CldKey      string
	MidTransKey string
	Schema      string
}

type mailSetting struct {
	Host     string
	Port     int
	Name     string
	Email    string
	Password string
}

func ImportGomailSetting() mailSetting {
	var result mailSetting

	if _, err := os.Stat(".env"); !os.IsNotExist(err) {
		err := godotenv.Load(".env")
		if err != nil {
			return mailSetting{}
		}
	} else {
		log.Println("file not exist")
	}
	result.Host = os.Getenv("GOMAIL_HOST")
	result.Port, _ = strconv.Atoi(os.Getenv("GOMAIL_PORT"))
	result.Name = os.Getenv("GOMAIL_NAME")
	result.Email = os.Getenv("GOMAIL_EMAIL")
	result.Password = os.Getenv("GOMAIL_PASSWORD")
	return result

}

func ImportSetting() setting {
	var result setting

	if _, err := os.Stat(".env"); !os.IsNotExist(err) {
		err := godotenv.Load(".env")
		if err != nil {
			return setting{}
		}
	} else {
		log.Println("file not exist")
	}
	result.User = os.Getenv("DB_USER")
	result.Host = os.Getenv("DB_HOST")
	result.Port = os.Getenv("DB_PORT")
	result.DBName = os.Getenv("DB_NAME")
	result.Password = os.Getenv("DB_PASSWORD")
	result.JWTSecret = os.Getenv("JWT_SECRET")
	result.CldKey = os.Getenv("CLOUDINARY_KEY")
	result.MidTransKey = os.Getenv("MIDTRANS_KEY")
	result.Schema = os.Getenv("SCHEMA")
	return result
}

func ConnectDB() (*gorm.DB, error) {
	s := ImportSetting()
	var connStr = fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s", s.Host, s.User, s.Password, s.Port, s.DBName)
	schem := ImportSetting().Schema
	schem = schem + "."
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: schem,
		},
	})
	if err != nil {
		return nil, err
	}
	// err = db.AutoMigrate(&u_rep.User{}, &t_rep.Trash{}, &d_rep.WasteDeposit{}, &l_rep.Location{}, &r_rep.Reward{}, &ex_rep.Exchange{}, &f_rep.Feedback{})
	// if err != nil {
	// 	return nil, err
	// }
	return db, nil
}
