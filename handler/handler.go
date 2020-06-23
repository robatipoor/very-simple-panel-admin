package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/robatipoor/very-simple-panel-admin/model"
	"github.com/robatipoor/very-simple-panel-admin/model"
)

// Home handler
func Home(resp http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	respondJSON(resp, http.StatusOK, "Wellcome !!!")
}

// AddProduct handler
func AddProduct(resp http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	product.InsertProduct(getProduct(req))
	respondJSON(resp, http.StatusOK, "Success")
}

func DeleteProduct(resp http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	product.Delete(getId(req))
	respondJSON(resp, http.StatusOK, "Success")
}

func GetProduct(resp http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	p := product.FindById(getId(req))
	respondJSON(resp, http.StatusOK, p)
}

func GetAllProduct(resp http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	p := product.FindAll()
	respondJSON(resp, http.StatusOK, p)
}

func UpdateProduct(resp http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	product.Update(getId(req), getProduct(req))
	respondJSON(resp, http.StatusOK, "Success")
}

// respondJSON makes the response with payload as json format
func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(response))
}

// respondError makes the error response with payload as json format
func respondError(w http.ResponseWriter, code int, message string) {
	respondJSON(w, code, map[string]string{"error": message})
}

func getProduct(req *http.Request) *product.Product {
	decoder := json.NewDecoder(req.Body)
	p := &product.Product{}
	err := decoder.Decode(p)
	if err != nil {
		log.Fatal(err)
	}
	return p
}

func getId(req *http.Request) uint{
	u64, err := strconv.ParseUint(mux.Vars(req)["id"], 10, 32)
    if err != nil {
        log.Fatal(err)
    }
   return uint(u64)
}