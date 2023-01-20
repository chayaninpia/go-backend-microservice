package util

import (
	"fmt"
	"os"
	"strings"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"xorm.io/xorm"
	"xorm.io/xorm/log"
)

type ServerInfo struct {
	Host      string
	HostLocal string
	Port      int
	User      string
	Password  string
	Dbname    string
}

func InitXorm() (*xorm.Engine, error) {

	s := ServerInfo{
		Host:      viper.GetString(`db.host`),
		HostLocal: viper.GetString(`db.hostLocal`),
		Port:      viper.GetInt(`db.port`),
		User:      viper.GetString(`db.user`),
		Password:  viper.GetString(`db.password`),
		Dbname:    viper.GetString(`db.dbname`),
	}
	psqlInfo := ``
	hostName, _ := os.Hostname()
	if strings.Contains(hostName, `local`) {
		psqlInfo = fmt.Sprintf("host=%s port=%d user=%s "+
			"password=%s dbname=%s sslmode=disable",
			s.HostLocal, s.Port, s.User, s.Password, s.Dbname)
	} else {
		psqlInfo = fmt.Sprintf("host=%s port=%d user=%s "+
			"password=%s dbname=%s sslmode=disable",
			s.Host, s.Port, s.User, s.Password, s.Dbname)
	}

	e, err := xorm.NewEngine("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	e.Logger().SetLevel(log.LOG_INFO)
	e.ShowSQL(true)

	return e, err
}
