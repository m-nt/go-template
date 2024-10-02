package internal

import (
	"os"
	"reflect"

	"github.com/joho/godotenv"
)

type Settings struct {
	DATABASE_URL    string `default:"postgres://root:12345@localhost:5432/root"`
	SERVICE_ADDRESS string `default:":3000"`
}

func load_Env(settings interface{}) {
	// load data from env
	godotenv.Load()
	v := reflect.ValueOf(settings).Elem()

	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)
		value := os.Getenv(field.Name)
		if value == "" {
			value = field.Tag.Get("default")
		}
		if v.Field(i).CanSet() && field.Type.Kind() == reflect.String {
			v.Field(i).SetString(value)
		}
	}
}

func Get_Settings() Settings {
	var settings Settings
	load_Env(&settings)
	return settings
}
