package cmd

// import (
// 	"encoding/json"
// 	"flag"
// 	"fmt"
// 	"matrix" // Import the matrix package.
// 	"net/http"
// )

// // MatrixPayload defines the structure for incoming matrix data.

// func healthCheck(w http.ResponseWriter, r *http.Request) {
// 	w.WriteHeader(http.StatusOK)
// }

// func handlerFunction(w http.ResponseWriter, r *http.Request) {
// 	operation := r.URL.Query().Get("operation")

// 	var payload matrix.MatrixPayload
// 	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
// 		fmt.Println("Decode error:", err)
// 		http.Error(w, "Invalid request body", http.StatusBadRequest)
// 		return
// 	}

// 	result, err := matrix.DistributeTask(payload, operation)
// 	if err != nil {
// 		fmt.Println("DistributeTask error:", err)
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	resp, err := json.Marshal(result.Data)
// 	if err != nil {
// 		fmt.Println("Marshal error:", err)
// 		http.Error(w, "Error marshalling response", http.StatusInternalServerError)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.Write(resp)
// }

// func main() {
// 	port := flag.String("port", "8080", "Port to run the server on")
// 	flag.Parse()

// 	http.HandleFunc("/health", healthCheck)
// 	http.HandleFunc("/multiply", multiplicationHandler)
// 	http.HandleFunc("/", handlerFunction)

// 	fmt.Println("Server is starting on port", *port)

// 	http.ListenAndServe(":"+*port, nil)
// }

// func multiplicationHandler(w http.ResponseWriter, r *http.Request) {
// 	var payload matrix.MatrixPayload
// 	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
// 		http.Error(w, "Invalid request body", http.StatusBadRequest)
// 		return
// 	}

// 	matrixA := matrix.NewMatrix(payload.MatrixA)
// 	matrixB := matrix.NewMatrix(payload.MatrixB)

// 	result, err := matrix.Multiply(matrixA, matrixB)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	resp, err := json.Marshal(result.Data)
// 	if err != nil {
// 		http.Error(w, "Error marshalling response", http.StatusInternalServerError)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.Write(resp)
// }
