package main

import (
	"errors"
	pb "github.com/acubed-tm/profile-service/protofiles"
	"github.com/neo4j/neo4j-go-driver/neo4j"
	"log"
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

func CreateProfileForUuid(uuid, profileUuid, firstName, lastName, description string) error {
	query := "MATCH (a:Account{uuid:{accountUuid}}) CREATE (a)-[:HAS_PROFILE]->(p:Profile{uuid: {profileUuid}, firstName: {firstName}, name: {name}, description: {description}})"
	variables := map[string]interface{}{"accountUuid": uuid, "profileUuid": profileUuid, "firstName": firstName, "name": lastName, "description": description}
	return Write(query, variables)
}

func GetProfileByUuid(uuid string) (firstName, lastName, description string, err error) {
	query := "MATCH (a:Account{uuid:{uuid}})-[:HAS_PROFILE]->(p:Profile) RETURN p.firstName, p.name, p.description"
	variables := map[string]interface{}{"uuid": uuid}

	type profile struct {
		firstName, lastName, description string
	}
	obj, err := Fetch(query, variables, func(res neo4j.Result) (interface{}, error) {
		if res.Next() {
			return profile{
				firstName:   res.Record().GetByIndex(0).(string),
				lastName:    res.Record().GetByIndex(1).(string),
				description: res.Record().GetByIndex(2).(string),
			}, nil
		}
		return nil, res.Err()
	})
	if err != nil {
		return "", "", "", err
	}

	if obj == nil {
		return "", "", "", errors.New("couldn't find profile")
	}

	p := obj.(profile)
	return p.firstName, p.lastName, p.description, nil
}

func UpdateProfileForUuid(uuid, firstName, lastName, description string) error {
	query := "MATCH (:Account{uuid:{uuid}})-[:HAS_PROFILE]->(p:Profile) SET p.firstName={firstName}, p.name={name}, p.description={description}"
	variables := map[string]interface{}{"uuid": uuid, "firstName": firstName, "name": lastName, "description": description}
	return Write(query, variables)
}

func CreateOrganizationProfileForUuid(uuid, profileUuid, displayName, description string) error {
	query := "MATCH (a:Organisation{uuid:{uuid}}) CREATE (a)-[:HAS_PROFILE]->(p:Profile{uuid: {profileUuid}, displayName: {name}, description: {description}})"
	variables := map[string]interface{}{"uuid": uuid, "profileUuid": profileUuid, "name": displayName, "description": description}
	return Write(query, variables)
}

func GetOrganizationProfileByUuid(uuid string) (displayName, description string, err error) {
	query := "MATCH (a:Organisation{uuid:{uuid}})-[:HAS_PROFILE]->(p:Profile) RETURN p.displayName, p.description"
	variables := map[string]interface{}{"uuid": uuid}

	type profile struct {
		displayName, description string
	}
	obj, err := Fetch(query, variables, func(res neo4j.Result) (interface{}, error) {
		if res.Next() {
			return profile{
				displayName: res.Record().GetByIndex(0).(string),
				description: res.Record().GetByIndex(1).(string),
			}, nil
		}
		return nil, res.Err()
	})
	if err != nil {
		return "", "", err
	}

	if obj == nil {
		return "", "", errors.New("couldn't find profile")
	}

	p := obj.(profile)
	return p.displayName, p.description, nil
}

func UpdateOrganizationProfileForUuid(uuid, displayName, description string) error {
	query := "MATCH (:Organisation{uuid:{uuid}})-[:HAS_PROFILE]->(p:Profile) SET p.displayName={name}, p.description={description}"
	variables := map[string]interface{}{"uuid": uuid, "name": displayName, "description": description}
	return Write(query, variables)
}

func GetEmailsByUuid(uuid string) ([]*pb.UserEmail, error) {
	query := "MATCH (e:Email)<-[:HAS_EMAIL]-(:Account{uuid:{uuid}}) RETURN e.emailAddress, e.isPrimary"
	variables := map[string]interface{}{"uuid": uuid}

	obj, err := Fetch(query, variables, func(res neo4j.Result) (interface{}, error) {
		var ret []*pb.UserEmail
		for res.Next() {
			rec := &pb.UserEmail{
				Email:     res.Record().GetByIndex(0).(string),
				IsPrimary: res.Record().GetByIndex(1).(bool),
			}
			log.Printf("aaa: %v", rec)
			ret = append(ret, rec)
		}

		log.Printf("final len: %v", len(ret))
		log.Printf("ret int: %v", ret)
		return ret, nil
	})

	if err != nil {
		return nil, err
	}

	if obj == nil {
		return nil, errors.New("couldn't find account by uuid")
	}

	log.Printf("returning: %v", obj)

	return obj.([]*pb.UserEmail), nil
}
