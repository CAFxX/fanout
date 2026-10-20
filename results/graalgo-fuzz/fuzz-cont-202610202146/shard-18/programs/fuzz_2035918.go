package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8; F2 int16 }

func main() {
    v1 := make(map[int]string)
    var v2 string = "foo bar"
    var v3 float64 = 1e10
    v4 := make(map[int]bool)
    v5 := make(map[string]bool)
    delete(v4, 0)
    switch v2 {
    case "hello":
        v6 := []int{}
        _ = v6
    case "foo bar":
        v7 := GenS1{F0: true, F1: 48, F2: -32768}
        _ = v7
    case "":
        delete(v1, 32768)
    default:
        v8 := GenS1{F0: ((1 <= 42) && (true && false)), F1: 176, F2: -32768}
        _ = v8
    }
    v9 := []int{}
    var v10 float64 = float64(3.14159)
    v10 = v10
    if len(v9) >= 2 {
        v9 = v9[1:len(v9)]
    }
    v2 = "hello"
    if true {
        switch "foo bar" {
        case "a":
            var v11 uint16 = 246
            _ = v11
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
