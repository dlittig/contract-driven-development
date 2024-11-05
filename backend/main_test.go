package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	api "cdd/backend/api/petstore"
	gen "cdd/backend/api/petstore/gen"

	"github.com/oapi-codegen/testutil"
	"github.com/stretchr/testify/assert"
)

func doGet(t *testing.T, handler http.Handler, url string) *httptest.ResponseRecorder {
	response := testutil.NewRequest().Get(url).WithAcceptJson().GoWithHTTPHandler(t, handler)
	return response.Recorder
}

func TestPetStore(t *testing.T) {
	var err error
	store := api.NewPetStore()
	ginPetServer := NewGinPetServer(store, "8080")
	r := ginPetServer.Handler

	t.Run("Add pet", func(t *testing.T) {
		tag := "TagOfSpot"
		newPet := gen.NewPet{
			Name: "Spot",
			Tag:  &tag,
		}

		rr := testutil.NewRequest().Post("/pets").WithJsonBody(newPet).GoWithHTTPHandler(t, r).Recorder
		assert.Equal(t, http.StatusCreated, rr.Code)
		assert.Empty(t, rr.Body.String())
		assert.Equal(t, 3, len(store.Pets)) // By default 2 pets are already in the store
	})

	t.Run("Find pet by ID", func(t *testing.T) {
		pet := gen.Pet{
			Id: "4d56e405-443e-4ba6-bad9-49b41b7d38ec",
		}

		store.Pets[pet.Id] = pet
		rr := doGet(t, r, fmt.Sprintf("/pets/%s", pet.Id))

		var resultPet gen.Pet
		err = json.NewDecoder(rr.Body).Decode(&resultPet)
		assert.NoError(t, err, "error getting pet")
		assert.Equal(t, pet, resultPet)
	})

	t.Run("Pet not found", func(t *testing.T) {
		rr := doGet(t, r, "/pets/27179095781")
		assert.Equal(t, http.StatusNotFound, rr.Code)

		var petError gen.Error
		err = json.NewDecoder(rr.Body).Decode(&petError)
		assert.NoError(t, err, "error getting response", err)
		assert.Equal(t, int32(http.StatusNotFound), petError.Code)
	})

	t.Run("List all pets", func(t *testing.T) {
		store.Pets = map[string]gen.Pet{"4d56e405-443e-4ba6-bad9-49b41b7d38ec": {}, "3d0c09a8-0999-4e7e-bf35-945513cc9611": {}}

		// Now, list all pets, we should have two
		rr := doGet(t, r, "/pets")
		assert.Equal(t, http.StatusOK, rr.Code)

		var petList []gen.Pet
		err = json.NewDecoder(rr.Body).Decode(&petList)
		assert.NoError(t, err, "error getting response", err)
		assert.Equal(t, 2, len(petList))
	})
}
