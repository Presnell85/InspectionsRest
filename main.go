package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//inspections Struct

type Inspection struct {
	ID               string  `json: "id"`
	Vin              int     `json: "vin"`
	Stock_Number     int     `json: "stock_number"`
	Lot_Number       int     `json: "lot_number"`
	PO_Number        int     `json: "po_number"`
	Year             int     `json: "year"`
	Make             string  `json: "make"`
	Model            string  `json: "model"`
	Type             string  `json: "type"`
	OD_Miles         int     `json: "od_miles"`
	Engine_Make      string  `json: "engine_make"`
	Engine_Detail    string  `json: "engine_detail"`
	Engine_HP        int     `json: "engine_hp"`
	Engine_Brake     bool    `json: "engine_brake"`
	Fuel_Type        string  `json: "fuel_type"`
	Fuel_Tanks       int     `json: "fuel_tanks"`
	Fuel_Capac       int     `json: "fuel_capac"`
	Exhaust          string  `json: "exhaust"`
	Trans_make       string  `json: "trans_make"`
	Trans_model      string  `json: "trans_model"`
	Trans_Spd        int     `json: "trans_spd"`
	Trans_Type       string  `json: "trans_type"`
	Cruise           bool    `json: "cruise"`
	Wet_Line         bool    `json: "wet_line"`
	PWR_Steer        bool    `json: "pwr_steer"`
	Rear_Spec        string  `json:"rear_spec"`
	Rear_Ratio       int     `json: "rear_ratio"`
	Axle_Type        string  `json: "axle_type"`
	Axle_F           int     `json: "axle_f"`
	Axle_R           int     `json: "axle_r"`
	Suspension       string  `json: "suspension"`
	WheelBase        int     `json: "wheelbase"`
	Tire_Size_R      float32 `json: "tire_size_r"`
	Tire_Size_F      float32 `json: "tire_size_f"`
	Wheel_F          string  `json: "wheel_f"`
	Wheel_R          string  `json: "wheel_r"`
	Wheel_5th        string  `json: "whhel_5th"`
	Air_Cond         bool    `json: "air_cond"`
	Sleeper_Size     int     `json: "sleeper_size"`
	Dual_Bunks       bool    `json: "dual_bunks"`
	Cab_Style        string  `json: "cab_style"`
	Int_Color        string  `json: "int_color"`
	Radio            string  `json: "radio"`
	Radio_Operable   bool    `json: "radio_operable"`
	Power_Mirrors    string  `json: "power_mirrors"`
	Hight_Back_Seats string  `json: "high_back_seats"`
	Pass_Seat        string  `json: "pass_seat"`
	EXT_Color        string  `json: "ext_color"`
	Bumper           string  `json: "bumper"`
	Heated_Mirror    bool    `json: "heated_mirror"`
	Fairings         string  `json: "fairings"`
	Visor            bool    `json: "visor"`
}

//Init inspections var as a slice inspections struct

var inspections []Inspection

//Get Inspections List

func getInspectionsList(w http.ResponseWriter, router *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(inspections)
}

//Get Inspection By ID

func getInspectionBYID(w http.ResponseWriter, router *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(router) //get the oparams
	for _, item := range inspections {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Inspection{})
}

//Update Inspections

func updateInspections(w http.ResponseWriter, router *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(router)
	for index, item := range inspections {
		if item.ID == params["id"] {
			inspections = append(inspections[:index], inspections[index+1:]...)
			var Inspection Inspection
			_ = json.NewDecoder(router.Body).Decode(&Inspection)
			Inspection.ID = params["id"]
			inspections = append(inspections, Inspection)
			json.NewEncoder(w).Encode(&Inspection)
			return
		}
	}

}

//create Inspections List

func createInspectionsList(w http.ResponseWriter, router *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var inspection Inspection
	_ = json.NewDecoder(router.Body).Decode(&inspection)
	inspection.ID = strconv.Itoa(rand.Intn(1000000000)) //this is not a real ID
	inspections = append(inspections, inspection)
	json.NewEncoder(w).Encode(inspection)
}

//Delete Inspections

func deleteInspections(w http.ResponseWriter, router *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(router)
	for index, item := range inspections {
		if item.ID == params["id"] {
			inspections = append(inspections[:index], inspections[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(inspections)
}

func main() {
	//Init the router

	router := mux.NewRouter()

	//mock data - todo implement MongoDB Collection
	inspections = append(inspections, Inspection{Vin: 123456789, Year: 2003, Make: "Mack", Model: "BMF", ID: "1"})
	//Rout Handlers / Endpoints

	router.HandleFunc("/api/inspectionsList", getInspectionsList).Methods("GET")
	router.HandleFunc("/api/inspections{id}", updateInspections).Methods("PUT")
	router.HandleFunc("/api/inspectionsList", createInspectionsList).Methods("POST")
	router.HandleFunc("/api/inspections{id}", deleteInspections).Methods("DELETE")
	router.HandleFunc("/api/inspections{id}", getInspectionBYID).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))
}
