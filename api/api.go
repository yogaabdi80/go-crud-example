package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/yogaabdi80/go-crud-example/model"
	"github.com/yogaabdi80/go-crud-example/repository"
)


type response struct{
    Status int `json:"status"`
    Message string `json:"message"`
    Result interface{}`json:"result"`
}

var resp response

func Router() *mux.Router {

	router := mux.NewRouter()

    // http.HandleFunc("/api/getAll", api.GetProducts)
	// http.HandleFunc("/api/product/", api.GetProduct)

	router.HandleFunc("/api/getAll", getProducts).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/product/{id}", getProduct).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/create", createProduct).Methods("POST", "OPTIONS")
	// router.HandleFunc("/api/buku/{id}", controller.UpdateBuku).Methods("PUT", "OPTIONS")
	// router.HandleFunc("/api/buku/{id}", controller.HapusBuku).Methods("DELETE", "OPTIONS")

	return router
}

func createProduct(w http.ResponseWriter, r *http.Request){
    p := new (model.Product)
    var err error
    decoder := json.NewDecoder(r.Body)
    if err := decoder.Decode(&p); err != nil {
        respondWithError(w, http.StatusBadRequest, "Invalid request payload")
        return
    }
    defer r.Body.Close()
    if p.ID == 0{
        p, err = repository.CreateProduct(p)
    } else{
        p, err = repository.UpdateProduct(p)
    }

    if err != nil {
        resp.Status=99
        resp.Message = err.Error()
        respondWithJSON(w, http.StatusBadRequest, &resp)
        return
    }
    resp.Status=200
    resp.Message = "Data Berhasil Disimpan!"
    resp.Result = p
    respondWithJSON(w, http.StatusCreated, &resp)
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	productList,err := repository.GetProducts()
	if err != nil {
        respondWithError(w, http.StatusInternalServerError, err.Error())
        return
    }
	respondWithJSON(w, http.StatusOK, productList)
}

func getProduct(w http.ResponseWriter, r *http.Request){
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        respondWithError(w, http.StatusBadRequest, "Invalid product ID")
        return
    }

	p := new (model.Product)
	p.ID = id
	if err != nil {
        respondWithError(w, http.StatusBadRequest, "Invalid product ID")
        return
    }
	p, err = repository.GetProduct(p)
	if err!=nil{
        respondWithError(w, http.StatusInternalServerError, err.Error())
        return
	}
	respondWithJSON(w, http.StatusOK, p)

}

func respondWithError(w http.ResponseWriter, code int, message string) {
    respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
    response, _ := json.Marshal(payload)

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(code)
    w.Write(response)
}