package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/CarApp/internal/model"
	"github.com/CarApp/internal/services"
	"github.com/CarApp/internal/utils"
)

type svc struct {
	ddd *services.CarRepo
}

func HandlerRequests(repo *services.CarRepo) {
	exp := svc{repo}
	http.HandleFunc("/", exp.listCarsDetails)
	http.HandleFunc("/all", exp.listAllCars)
	http.HandleFunc("/add", exp.addCarsDetails)
	http.HandleFunc("/update", exp.updateCarDetails)
}

func (s *svc) listCarsDetails(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		// Add the response return message
		HandlerMessage := []byte(`{
			"success": false,
			"message": "Check your HTTP method: Invalid HTTP method executed",
		   }`)

		utils.ReturnJsonResponse(w, http.StatusMethodNotAllowed, HandlerMessage)
		return
	}

	if _, ok := req.URL.Query()["id"]; !ok {
		// Add the response return message
		HandlerMessage := []byte(`{
		 "success": false,
		 "message": "This method requires unique id",
		}`)

		utils.ReturnJsonResponse(w, http.StatusInternalServerError, HandlerMessage)
		return
	}

	id := req.URL.Query()["id"][0]
	response := s.ddd.ListCar(id)
	if response == nil {
		HandlerMessage := []byte(`{
			"success": false,
			"message": "No car records found against this ID",
		   }`)

		utils.ReturnJsonResponse(w, http.StatusNotFound, HandlerMessage)
		return
	}
	byteData, err := json.Marshal(response)
	if err != nil {
		fmt.Printf("%s", err)
		return
	}

	utils.ReturnJsonResponse(w, http.StatusOK, byteData)
}

func (s *svc) listAllCars(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		// Add the response return message
		HandlerMessage := []byte(`{
			"success": false,
			"message": "Check your HTTP method: Invalid HTTP method executed",
		   }`)

		utils.ReturnJsonResponse(w, http.StatusMethodNotAllowed, HandlerMessage)
		return
	}

	response := s.ddd.ListAll()
	byteData, err := json.Marshal(response)
	if err != nil {
		fmt.Printf("%s", err)
		return
	}
	utils.ReturnJsonResponse(w, http.StatusOK, byteData)

}

func (s *svc) addCarsDetails(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		// Add the response return message
		HandlerMessage := []byte(`{
			"success": false,
			"message": "Check your HTTP method: Invalid HTTP method executed",
		   }`)

		utils.ReturnJsonResponse(w, http.StatusMethodNotAllowed, HandlerMessage)
		return
	}

	var carsData []model.Car
	defer req.Body.Close()

	if err := json.NewDecoder(req.Body).Decode(&carsData); err != nil {
		// Add the response return message
		HandlerMessage := []byte(`{
			"success": false,
			"message": "Error parsing the movie data",
   		}`)

		utils.ReturnJsonResponse(w, http.StatusInternalServerError, HandlerMessage)
		return
	}

	// Store it in DB
	byteData, _ := json.Marshal(carsData)
	s.ddd.Dataa = byteData
	_ = s.ddd.AddCarDetails()

	successMessage := []byte(fmt.Sprintf(`{
		"success": true,
		"message": "Data stored successfully",
		"count": %d,
	   }`, len(carsData)))

	utils.ReturnJsonResponse(w, http.StatusOK, successMessage)
}

func (s *svc) updateCarDetails(w http.ResponseWriter, req *http.Request) {
	if req.Method != "PUT" {
		// Add the response return message
		HandlerMessage := []byte(`{
			"success": false,
			"message": "Check your HTTP method: Invalid HTTP method executed",
		   }`)

		utils.ReturnJsonResponse(w, http.StatusMethodNotAllowed, HandlerMessage)
		return
	}

	var carData model.Car
	defer req.Body.Close()

	if err := json.NewDecoder(req.Body).Decode(&carData); err != nil {
		// Add the response return message
		HandlerMessage := []byte(`{
			"success": false,
			"message": "Error parsing the car data",
   		}`)

		utils.ReturnJsonResponse(w, http.StatusInternalServerError, HandlerMessage)
		return
	}

	if err := s.ddd.UpdateCarDetails(carData); err != nil {
		// Add the response return message
		HandlerMessage := []byte(fmt.Sprintf(`{
			"success": false,
			"message": "%s",
   		}`, err))

		utils.ReturnJsonResponse(w, http.StatusNotFound, HandlerMessage)
		return
	}

	utils.ReturnJsonResponse(w, http.StatusOK, []byte("Data updated successfully"))
}
