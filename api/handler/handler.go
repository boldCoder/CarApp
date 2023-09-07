package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/CarApp/internal/model"
	"github.com/CarApp/internal/services"
	"github.com/CarApp/internal/utils"
)

type service struct {
	svc *services.CarRepo
}

// Register requests
func HandlerRequests(repo *services.CarRepo) {
	exp := service{repo}
	http.HandleFunc("/get", exp.listCarDetails)
	http.HandleFunc("/all", exp.listAllCars)
	http.HandleFunc("/add", exp.addCarsDetails)
	http.HandleFunc("/update", exp.updateCarDetails)
}

func (s *service) listCarDetails(w http.ResponseWriter, req *http.Request) {
	// Check if request method is not GET
	if req.Method != "GET" {
		// Add the response return message
		HandlerMessage := []byte(`{
			"success": false,
			"message": "Check your HTTP method: Invalid HTTP method executed",
		   }`)

		// Return response in JSON with http-status and message
		utils.ReturnJsonResponse(w, http.StatusMethodNotAllowed, HandlerMessage)
		return
	}

	if _, ok := req.URL.Query()["id"]; !ok {
		// Add the response return message
		HandlerMessage := []byte(`{
		 "success": false,
		 "message": "This method requires unique id",
		}`)

		// Return response in JSON with http-status and message
		utils.ReturnJsonResponse(w, http.StatusInternalServerError, HandlerMessage)
		return
	}

	id := req.URL.Query()["id"][0]
	response := s.svc.ListCar(id)
	if response == nil {
		HandlerMessage := []byte(`{
			"success": false,
			"message": "No car records found against this ID",
		   }`)

		// Return response in JSON with http-status and message
		utils.ReturnJsonResponse(w, http.StatusNotFound, HandlerMessage)
		return
	}

	byteData, err := json.Marshal(response)
	if err != nil {
		HandlerMessage := []byte(fmt.Sprintf(`{
			"success": false,
			"message": "%s",
		   }`, err.Error()))

		// Return response in JSON with http-status and message
		utils.ReturnJsonResponse(w, http.StatusInternalServerError, HandlerMessage)
		return
	}
	// Return response in JSON with http-status and message
	utils.ReturnJsonResponse(w, http.StatusOK, byteData)
}

func (s *service) listAllCars(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		// Add the response return message
		HandlerMessage := []byte(`{
			"success": false,
			"message": "Check your HTTP method: Invalid HTTP method executed",
		   }`)

		// Return response in JSON with http-status and message
		utils.ReturnJsonResponse(w, http.StatusMethodNotAllowed, HandlerMessage)
		return
	}

	response := s.svc.ListAll()
	byteData, err := json.Marshal(response)
	if err != nil {
		HandlerMessage := []byte(fmt.Sprintf(`{
			"success": false,
			"message": "%s",
		   }`, err.Error()))

		// Return response in JSON with http-status and message
		utils.ReturnJsonResponse(w, http.StatusInternalServerError, HandlerMessage)
		return
	}

	// Return response in JSON with http-status and message
	utils.ReturnJsonResponse(w, http.StatusOK, byteData)

}

func (s *service) addCarsDetails(w http.ResponseWriter, req *http.Request) {
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

	byteData, err := json.Marshal(carsData)
	if err != nil {
		HandlerMessage := []byte(fmt.Sprintf(`{
			"success": false,
			"message": "%s",
		   }`, err.Error()))

		// Return response in JSON with http-status and message
		utils.ReturnJsonResponse(w, http.StatusInternalServerError, HandlerMessage)
		return
	}
	// Store it in DB
	s.svc.Details = byteData
	if err = s.svc.AddCarDetails(); err != nil {
		HandlerMessage := []byte(fmt.Sprintf(`{
			"success": false,
			"message": "%s",
		   }`, err.Error()))

		// Return response in JSON with http-status and message
		utils.ReturnJsonResponse(w, http.StatusInternalServerError, HandlerMessage)
		return
	}

	successMessage := []byte(fmt.Sprintf(`{
		"success": true,
		"message": "Data stored successfully",
		"count": %d,
	   }`, len(carsData)))

	// Return response in JSON with http-status and message
	utils.ReturnJsonResponse(w, http.StatusOK, successMessage)
}

func (s *service) updateCarDetails(w http.ResponseWriter, req *http.Request) {
	if req.Method != "PUT" {
		// Add the response return message
		HandlerMessage := []byte(`{
			"success": false,
			"message": "Check your HTTP method: Invalid HTTP method executed",
		   }`)

		// Return response in JSON with http-status and message
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

		// Return response in JSON with http-status and message
		utils.ReturnJsonResponse(w, http.StatusInternalServerError, HandlerMessage)
		return
	}

	if err := s.svc.UpdateCarDetails(carData); err != nil {
		// Add the response return message
		HandlerMessage := []byte(fmt.Sprintf(`{
			"success": false,
			"message": "%s",
   		}`, err))

		// Return response in JSON with http-status and message
		utils.ReturnJsonResponse(w, http.StatusNotFound, HandlerMessage)
		return
	}

	// Return response in JSON with http-status and message
	utils.ReturnJsonResponse(w, http.StatusOK, []byte("Data updated successfully"))
}
