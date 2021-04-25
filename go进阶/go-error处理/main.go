package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/go-xorm/xorm"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

//Engine 配置项
type Orm struct {
	Engine        *xorm.Engine
	Database      string `yaml:"database"`
	Host          string `yaml:"host"`
	Port          int    `yaml:"port"`
	User          string `yaml:"user"`
	Password      string `yaml:"password"`
	Socket        string `yaml:"socket"`
	Debug         bool   `yaml:"debug"`
	Charset       string `yaml:"charset"`
	MaxIdleConn   int    `yaml:"max_idle_conn"`
	DbMaxOpenConn int    `yaml:"max_open_conn"`
	MaxLifeTime   int    `yaml:"max_life_time"`
}

// 根据结构体初始化
func (it *Orm) init() (err error) {
	var dbDef string

	if it.Host == "localhost" {
		dbDef = fmt.Sprintf("%s:%s@unix(%s)/%s?charset=%s", it.Host, it.User, it.Socket, it.Database, it.Charset)
	} else {
		dbDef = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", it.User, it.Password, it.Host, it.Port, it.Database, it.Charset)
	}

	it.Engine, err = xorm.NewEngine("mysql", dbDef)
	if err != nil {
		return errors.Wrap(err, "new mysql engine err")
	}

	if it.Debug {
		it.Engine.ShowSQL(true)
	}
	if it.MaxIdleConn > 0 {
		it.Engine.SetMaxIdleConns(it.MaxIdleConn)
	}
	if it.DbMaxOpenConn > 0 {
		it.Engine.SetMaxOpenConns(it.DbMaxOpenConn)
	}
	if it.MaxLifeTime > 0 {
		it.Engine.SetConnMaxLifetime(time.Duration(it.MaxLifeTime) * time.Second)
	}
	err = it.Engine.Ping()
	if err != nil {
		return errors.Wrap(err, "engine ping failed")
	}
	return
}

// 通过结构体解析单个orm
func UnmarshalOrmByConf(engine *Orm) (res *xorm.Engine, err error) {
	err = engine.init()
	res = engine.Engine
	return
}

var (
	ormConf   *Orm
	ormEngine *xorm.Engine
)

func InitDao(confPath string) error {
	dataBytes, err := ioutil.ReadFile(confPath)
	if err != nil {
		return errors.Wrap(err, "read config file err")
	}

	err = yaml.Unmarshal(dataBytes, ormConf)
	if err != nil {
		return errors.Wrap(err, "unmarshal orm config bytes err")
	}

	ormEngine, err = UnmarshalOrmByConf(ormConf)
	if err != nil {
		return errors.Wrap(err, "unmarshal orm config err")
	}

	return nil
}

type Test struct {
	Id         int64
	Name       string
	Age        int64
	InsertTime int64
	UpdateTime int64
}

func GetAgeByTest(name string) (bool, error) {
	var test Test

	//ErrNoRows 类型的错误不应该返回给上层,而是通过has字段告诉上层是否存在该行数据
	has, err := ormEngine.Where("name=?", name).Get(&test)
	if err != nil {
		return true, errors.Wrap(err, "get test name err")
	}
	return has, nil
}

func main() {
	err := InitDao("./orm.yaml")
	if err != nil {
		log.Fatal(err)
	}

	has, err := GetAgeByTest("bai")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(has)
}
