package main

import (
	"calculator_service/internal/calculator"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type Request struct {
	Expression string `json:"expression"`
}

type Response struct {
	Result string `json:"result,omitempty"`
	Error  string `json:"error,omitempty"`
}

func calculateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var request Request
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&request); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if strings.Contains(request.Expression, "$") {
		http.Error(w, `{"error":"Internal server error"}`, http.StatusInternalServerError)
		return
	}

	res, err := calculator.Calc(request.Expression)
	var resp Response
	//fmt.Printf("%s\n", err)

	if err != nil {
		resp.Error = err.Error()
		w.WriteHeader(http.StatusUnprocessableEntity)
	} else {
		resp.Result = fmt.Sprintf("%f", res)
		w.WriteHeader(http.StatusOK)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func main() {
	fmt.Printf("server start\n")
	http.HandleFunc("/api/v1/calculate", calculateHandler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", "localhost", "8080"), nil))
}
