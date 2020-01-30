package main

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/dgraph-io/dgo/v2"
	"github.com/dgraph-io/dgo/v2/protos/api"
	"google.golang.org/grpc"
	"log"
)

const databaseUrl = "my-release-dgraph-alpha.acubed:9080"

func newClient() *dgo.Dgraph {
	// Dial a gRPC connection. The address to dial to can be configured when
	// setting up the dgraph cluster.
	d, err := grpc.Dial(databaseUrl, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	return dgo.NewDgraphClient(
		api.NewDgraphClient(d),
	)
}

func UpdateProfileForUuid(ctx context.Context, uuid, firstName, lastName string) error {
	c := newClient()

	variables := map[string]string{"$uuid": uuid}
	q := `
		query x($uuid: string){
			account(func: eq(uuid, $uuid)) {
				uid
				uuid
				profile {
					uid
					firstName
					lastName
				}
			}
		}
	`

	txn := c.NewTxn()
	resp, err := txn.QueryWithVars(ctx, q, variables)
	if err != nil {
		return err
	}

	var decode struct {
		All []struct {
			Uid string `json:"uid"`
			Uuid string `json:"uuid"`
			Profile struct {
				// TODO: not an array? needs uid?
				Uid string `json:"uid"`
				FirstName string `json:"firstName"`
				LastName string `json:"lastName"`
			} `json:"profile"`
		} `json:"uuid"`
	}
	log.Println("JSON: " + string(resp.GetJson()))
	if err := json.Unmarshal(resp.GetJson(), &decode); err != nil {
		return err
	}

	if len(decode.All) == 0 {
		return errors.New("couldn't find account by uid")
	}

	// TODO: what if matching failed because we don't have a profile yet? probably want to remove profile from query
	decode.All[0].Profile.FirstName = firstName
	decode.All[0].Profile.LastName = lastName

	// insert again
	out, err := json.Marshal(decode.All[0].Profile)
	if err != nil {
		return err
	}

	_, err = txn.Mutate(context.Background(), &api.Mutation{SetJson: out, CommitNow: true})
	if err != nil {
		return err
	}

	return nil
}