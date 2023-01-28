package main

import (
	"fmt"

	"github.com/ulfox/dby/db"
)

func generateYAMLFile(albumJSON []album) {
	fmt.Println(albumJSON)

	state, err := db.NewStorageFactory(dbFile)
	if err != nil {
		logger.Fatalf(err.Error())
	}

	err = state.Upsert(
		"albums",
		albumJSON,
	)
	if err != nil {
		logger.Fatalf(err.Error())
	}
}

func getYAMLFile() {
	state, err := db.NewStorageFactory(dbFile)
	if err != nil {
		logger.Fatalf(err.Error())
	}
}
