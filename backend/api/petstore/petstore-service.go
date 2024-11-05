//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=types.cfg.yaml ../../../spec/openapi-spec-v1.yaml
//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=server.cfg.yaml ../../../spec/openapi-spec-v1.yaml

package petstore

import (
	gen "cdd/backend/api/petstore/gen"
	"fmt"
	"net/http"
	"sort"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type PetStore struct {
	Pets map[string]gen.Pet
	Lock sync.Mutex
}

func NewPetStore() *PetStore {
	store := &PetStore{
		Pets: make(map[string]gen.Pet),
	}

	// Add fixtures
	doggoId := uuid.New().String()
	doggoTag := "Doggo"
	store.Pets[doggoId] = gen.Pet{
		Id:   doggoId,
		Name: "Snoops",
		Tag:  &doggoTag,
	}

	birbId := uuid.New().String()
	birbTag := "Birb"
	store.Pets[birbId] = gen.Pet{
		Id:   birbId,
		Name: "Birb",
		Tag:  &birbTag,
	}

	return store
}

// sendPetStoreError wraps sending of an error in the Error format, and
// handling the failure to marshal that.
func sendPetStoreError(c *gin.Context, code int, message string) {
	petErr := gen.Error{
		Code:    int32(code),
		Message: message,
	}
	c.JSON(code, petErr)
}

func getSortedPetList(p *PetStore) []gen.Pet {
	petSlice := make([]gen.Pet, 0, len(p.Pets))
	for _, pet := range p.Pets {
		petSlice = append(petSlice, pet)
	}

	// Sort the slice
	sort.Slice(petSlice, func(i, j int) bool {
		return petSlice[i].Id < petSlice[j].Id
	})

	return petSlice
}

// FindPets implements all the handlers in the ServerInterface
func (p *PetStore) ListPets(c *gin.Context, params gen.ListPetsParams) {
	p.Lock.Lock()
	defer p.Lock.Unlock()

	petList := getSortedPetList(p)
	var result []gen.Pet

	for _, pet := range petList {
		result = append(result, pet)

		if params.Limit != nil {
			l := int(*params.Limit)
			if len(result) >= l {
				// We're at the limit
				break
			}
		}
	}

	c.JSON(http.StatusOK, result)
}

func (p *PetStore) CreatePets(c *gin.Context) {
	// We expect a NewPet object in the request body.
	var newPet gen.NewPet
	err := c.Bind(&newPet)
	if err != nil {
		sendPetStoreError(c, http.StatusBadRequest, "Invalid format for NewPet")
		return
	}
	// We now have a pet, let's add it to our "database".

	// We're always asynchronous, so lock unsafe operations below
	p.Lock.Lock()
	defer p.Lock.Unlock()

	// Insert into map
	id := uuid.New().String()
	pet := gen.Pet{
		Id:   id,
		Name: newPet.Name,
		Tag:  newPet.Tag,
	}
	p.Pets[id] = pet

	c.Status(http.StatusCreated)
}

func (p *PetStore) ShowPetById(c *gin.Context, petId string) {
	p.Lock.Lock()
	defer p.Lock.Unlock()

	pet, found := p.Pets[petId]
	if !found {
		sendPetStoreError(c, http.StatusNotFound, fmt.Sprintf("Could not find pet with ID %s", petId))
		return
	}
	c.JSON(http.StatusOK, pet)
}
