package configs

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"gopkg.in/go-ini/ini.v1"
)

// ConfigList : config struct
type ConfigList struct {
	User            string
	Password        string
	Host            string
	Port            string
	Database        string
	ServerAccount   string
	ServerAccountPW string
	Signature       string
	BucketPath      string
}

// Config : config
var Config ConfigList

// InitConfig : config init
// DEV
func InitConfig() error {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		return errors.Wrap(err, "InitConfig()")
	}

	Config = ConfigList{
		User:            cfg.Section("DB").Key("User").String(),
		Password:        cfg.Section("DB").Key("Password").String(),
		Host:            cfg.Section("DB").Key("Host").String(),
		Port:            cfg.Section("DB").Key("Port").String(),
		Database:        cfg.Section("DB").Key("Database").String(),
		ServerAccount:   cfg.Section("Email").Key("ServerAccount").String(),
		ServerAccountPW: cfg.Section("Email").Key("ServerAccountPW").String(),
		Signature:       cfg.Section("JWT").Key("Signature").String(),
		BucketPath:      cfg.Section("STORAGE").Key("BucketPath").String(),
	}

	return nil
}
