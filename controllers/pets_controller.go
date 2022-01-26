package controllers

import (
	"api_integration/client"
	"api_integration/models"
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// PetsController ...
type PetsController struct {
	client client.ClientInterface
}

type PetsControllerInterface interface {
	GetPets(c *gin.Context)
	PostPets(c *gin.Context)
	GetPetById(c *gin.Context)
}

func NewPetsController(client client.ClientInterface) PetsControllerInterface {
	return &PetsController{
		client: client,
	}
}

func (p *PetsController) GetPets(c *gin.Context) {

	body := p.client.NewRequest(http.MethodGet, "http://petstore-demo-endpoint.execute-api.com/petstore/pets", nil)

	petsResp := new(models.GetPetsResponse)

	err := json.Unmarshal(body, petsResp)
	if err != nil {
		log.Fatalln("Errored while JSON Unmarshal")
	}

	c.JSON(http.StatusOK, gin.H{"getPets": petsResp})
}

func (p *PetsController) PostPets(c *gin.Context) {

	petsReq := new(models.PostPetsRequest)

	if err := c.BindJSON(petsReq); err != nil {
		log.Fatal("Error while BindJSON")
	}

	parameter, _ := json.Marshal(petsReq)

	body := p.client.NewRequest(http.MethodPost, "http://petstore-demo-endpoint.execute-api.com/petstore/pets", bytes.NewBuffer(parameter))

	postPets := new(models.PostPetsResponse)

	err := json.Unmarshal(body, postPets)
	if err != nil {
		log.Fatalln("Errored while JSON Unmarshal")
	}

	c.JSON(http.StatusOK, gin.H{"postPets": postPets})
}

func (p *PetsController) GetPetById(c *gin.Context) {

	id, ok := c.Params.Get("id")

	if !ok {
		log.Fatal("Error while fetching params id")
	}

	body := p.client.NewRequest(http.MethodGet, "http://petstore-demo-endpoint.execute-api.com/petstore/pets/"+id, nil)

	getPetById := new(models.PostPetsRequest)

	err := json.Unmarshal(body, getPetById)
	if err != nil {
		log.Fatalln("Errored while JSON Unmarshal")
	}

	c.JSON(http.StatusOK, gin.H{"getPetById": getPetById})
}
