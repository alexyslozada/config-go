package config_test

import (
	"testing"

	"github.com/alexyslozada/config-go"
)

func TestLoadFile(t *testing.T) {
	_, err := config.LoadFile("config.json.example")
	if err != nil {
		t.Errorf("no se pudo leer el archivo: %v", err)
	}
}

func TestLoadBytes(t *testing.T) {
	j := []byte(`{
	  "db_server": "127.0.0.1",
	  "db_port": 5432,
	  "db_name": "mibasededatos",
	"db_user": "usuario",
	"db_password": "clave"
	}`)
	c := config.Configuration{}
	err := config.LoadBytes(j, &c)
	if err != nil {
		t.Errorf("no se pudo leer los bytes: %v", err)
	}
}

func TestNew(t *testing.T) {
	_, err := config.New("config.json.example")
	if err != nil {
		t.Errorf("no se pudo cargar la configuracion: %v", err)
	}
}

func TestConfiguration_Validate(t *testing.T) {
	c, err := config.New("config.json.example")
	if err != nil {
		t.Errorf("no se pudo cargar la configuracion: %v", err)
	}

	err = c.Validate("db_server", "db_port", "db_user", "db_password", "is_secure")
	if err != nil {
		t.Error(err)
	}
}

func TestConfiguration_Get(t *testing.T) {
	c, err := config.New("config.json.example")
	if err != nil {
		t.Errorf("no se pudo cargar la configuracion: %v", err)
	}

	_, err = c.Get("db_server")
	if err != nil {
		t.Error(err)
	}
}

func TestConfiguration_GetInt(t *testing.T) {
	c, err := config.New("config.json.example")
	if err != nil {
		t.Errorf("no se pudo cargar la configuracion: %v", err)
	}

	_, err = c.GetInt("db_port")
	if err != nil {
		t.Error(err)
	}
}

func TestConfiguration_GetFloat(t *testing.T) {
	c, err := config.New("config.json.example")
	if err != nil {
		t.Errorf("no se pudo cargar la configuracion: %v", err)
	}

	_, err = c.GetFloat("db_port")
	if err != nil {
		t.Error(err)
	}
}

func TestConfiguration_GetBool(t *testing.T) {
	c, err := config.New("config.json.example")
	if err != nil {
		t.Errorf("no se pudo cargar la configuracion: %v", err)
	}

	_, err = c.GetBool("is_secure")
	if err != nil {
		t.Error(err)
	}
}
