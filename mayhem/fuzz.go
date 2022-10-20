package fuzz

import "ariga.io/entimport/internal/mux"

func mayhemit(bytes []byte) int {

    content := string(bytes)
    var test mux.Mux
    test.OpenImport(content)
    return 0
}

func Fuzz(data []byte) int {
    _ = mayhemit(data)
    return 0
}