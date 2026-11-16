package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int; F2 int }

func main() {
    var v1 string = ""
    v2 := []string{v1, v1, "\n\t"}
    var v3 uint32 = 42
    v4 := []uint32{(v3 * v3)}
    var v5 int16 = -529
    v6 := GenS1{F0: (v1 + (v1 + "foo bar")), F1: 772, F2: 4294967295}
    v5 = v5
    v4 = append(v4, (^(v3)))
    v7 := GenS1{F0: v1, F1: 363, F2: 72}
    v1 = v1
    v5 = v5
    var v8 int8 = 61
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
