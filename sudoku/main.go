package main

import (
    "encoding/json"
    "net/http"

    "github.com/gorilla/mux"
)

func generateSudokuHandler(w http.ResponseWriter, r *http.Request) {
    newSudoku := sudoku.GenerateSudoku()
    json.NewEncoder(w).Encode(newSudoku)
}

func solveSudokuHandler(w http.ResponseWriter, r *http.Request) {
    var inputSudoku [][]int
    // 解析请求中的数独盘面
    err := json.NewDecoder(r.Body).Decode(&inputSudoku)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    solution := su.SolveSudoku(inputSudoku)
    json.NewEncoder(w).Encode(solution)
}

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/generate", generateSudokuHandler).Methods("GET")
    r.HandleFunc("/solve", solveSudokuHandler).Methods("POST")

    http.Handle("/", r)
    http.ListenAndServe(":8080", nil)
}