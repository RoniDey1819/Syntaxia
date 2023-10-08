package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/PratikforCoding/Syntaxia/json"
	model "github.com/PratikforCoding/Syntaxia/models"
)

func (apiCfg *APIConfig)HandlerRegister(w http.ResponseWriter, r *http.Request) {
	var attendee model.Attendee
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&attendee)
	if err != nil {
		reply.RespondWtihError(w, http.StatusInternalServerError, "Error decoding parameter!")
		return
	}
	createdAttendee, err := apiCfg.register(attendee)
	fmt.Println(createdAttendee)
	if err != nil {
		reply.RespondWtihError(w, 201, "Error registering new attendee")
		return
	}
	retAttendee := model.Attendee{
		ID: createdAttendee.ID,
		SerialNo: createdAttendee.SerialNo,
	}
	reply.RespondWithJson(w, 200, retAttendee)
}

func (apiCfg *APIConfig)HandlerGetAttendeebyId(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	serialNo := r.URL.Query().Get("serialno")
	foundAttendee, err := apiCfg.getAttendee(serialNo)
	if err != nil {
		reply.RespondWtihError(w, http.StatusNotFound, "Bus not found")
		return 
	}
	retAttendee := model.Attendee{
		FirstName: foundAttendee.FirstName,
		LastName: foundAttendee.LastName,
	}
	reply.RespondWithJson(w, http.StatusFound, retAttendee)
}

func (apiCfg *APIConfig)HandlerGetAllAttendees(w http.ResponseWriter, r *http.Request) {
	allAttendee, err := apiCfg.getAllAttendees()
	if err != nil {
		reply.RespondWtihError(w, http.StatusInternalServerError, "Error getting attendees")
		return
	}
	reply.RespondWithJson(w, http.StatusFound, allAttendee)
}

func (apiCfg *APIConfig)HandlerTaken(w http.ResponseWriter, r *http.Request) {
	var attendee model.Attendee
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&attendee)
	if err != nil {
		reply.RespondWtihError(w, http.StatusInternalServerError, "Error decoding parameter!")
		return
	}
	serialno := attendee.SerialNo
	claimedAttendee, err := apiCfg.claimefood(serialno)
	if err != nil {
		reply.RespondWtihError(w, http.StatusInternalServerError, "Error upadtaing attendee")
		return
	}
	reply.RespondWithJson(w, http.StatusAccepted, claimedAttendee)
}

func (apiCfg *APIConfig)HandlerGetAttendeesbyYear(w http.ResponseWriter, r *http.Request) {
	var attendee model.Attendee
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&attendee)
	if err != nil {
		reply.RespondWtihError(w, http.StatusInternalServerError, "Error decoding parameter!")
		return
	}
	year := attendee.Year
	yearAttendees, err := apiCfg.getAttendeesByYear(year)
	if err != nil {
		reply.RespondWtihError(w, http.StatusInternalServerError, "Error upadtaing attendee")
		return
	}
	reply.RespondWithJson(w, http.StatusAccepted, yearAttendees)
}