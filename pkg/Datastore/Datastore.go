package Datastore

import (
	"Lab3/pkg/utils"
	"errors"
	"strconv"
)

var DataStore = make(map[string]string)

func Read(ID string) (string, error) {
	if value, ok := DataStore[ID]; ok {
		return value, nil
	}
	return "", errors.New("Value with ID " + ID + " doesn't exist.")
}

func Create(value string) (string, error) {
	var ID = strconv.Itoa(utils.Random())
	if _, ok := DataStore[ID]; ok {
		return "", errors.New("ID " + ID + " already exists.")
	}
	DataStore[ID] = value
	return value, nil
}

func Update(ID string, value string) (string, error) {
	if _, ok := DataStore[ID]; !ok {
		return "", errors.New("Value with ID " + ID + " doesn't exist.")
	}

	DataStore[ID] = value
	return value, nil
}

func Delete(ID string) (string, error) {
	if _, ok := DataStore[ID]; !ok {
		return "", errors.New("Value with ID " + ID + " doesn't exist.")
	}

	delete(DataStore, ID)
	return "Item successfully deleted.", nil
}
