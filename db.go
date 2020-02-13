package main

import (
	"github.com/neo4j/neo4j-go-driver/neo4j"
)

const databaseUrl = "bolt://neo4j-public.default:7687"
const username = "neo4j"
const password = ""

var driver neo4j.Driver = nil

func newSession(accessMode neo4j.AccessMode) neo4j.Session {
	var (
		err     error
		session neo4j.Session
	)

	if driver == nil {
		driver, err = neo4j.NewDriver(databaseUrl, neo4j.BasicAuth(username, password, ""))
		if err != nil {
			panic(err.Error())
		}
	}

	session, err = driver.Session(accessMode)
	if err != nil {
		panic(err.Error())
	}

	return session
}

func UpdateProfileForUuid(uuid, firstName, lastName, description string) error {
	query := "MATCH (:Account{uuid:{uuid}})-[:HAS_PROFILE]->(p:Profile) SET p.firstName={firstName}, p.name={name}, p.description={description}"
	variables := map[string]interface{}{"uuid": uuid, "firstName": firstName, "name": lastName, "description": description}
	return Write(query, variables)
}
