package database

import (
	"HttpServer/pkg/consts"
	"fmt"
	"sync"

	"github.com/sirupsen/logrus"
)

var (
	lock *sync.RWMutex
)

type Database struct {
	table map[string]string
}

func (d *Database) Init() {
	d.table = make(map[string]string)
}

func (d *Database) Get(key string) (string, error) {
	logrus.Infof("Get date from db key: %s", key)
	lock.RLock()
	defer lock.Unlock()

	if _, exist := d.table[key]; !exist {
		return consts.EmptyString, fmt.Errorf(consts.NotFoundError)
	}

	return d.table[key], nil
}

func (d *Database) Insert(key string, value string) error {
	logrus.Infof("Insert data to db, key: %s, value: %s", key, value)
	lock.Lock()
	defer lock.Unlock()

	if _, exist := d.table[key]; exist {
		return fmt.Errorf(consts.DataExistError)
	}

	d.table[key] = value
	return nil
}

func (d *Database) BatchInsert(data map[string]string) error {
	logrus.Infof("BatchInsert data to db, %s", data)
	lock.Lock()
	defer lock.Unlock()

	for k := range data {
		if _, exist := d.table[k]; exist {
			logrus.Errorf("key exist: %s", k)
			return fmt.Errorf(consts.DataExistError)
		}
	}

	for k, v := range data {
		d.table[k] = v
	}

	return nil
}

func (d *Database) Update(key string, value string) error {
	logrus.Infof("Insert data to db, key: %s, value: %s", key, value)
	lock.Lock()
	defer lock.Unlock()

	if _, exist := d.table[key]; !exist {
		return fmt.Errorf(consts.NotFoundError)
	}

	d.table[key] = value
	return nil
}
