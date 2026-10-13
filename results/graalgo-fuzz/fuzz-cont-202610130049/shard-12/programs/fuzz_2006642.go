package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8; F2 uint }
type GenS2 struct { F0 bool; F1 int; F2 string }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 bool }

func main() {
    var v1 bool = false
    v2 := []int{114, 128, 128}
    v3 := make(map[string]int)
    switch "\n\t" {
    case "\n\t":
        for fi4 := 0; fi4 < 13; fi4++ {
            v1 = true
            for fi5 := 0; fi5 < 13; fi5++ {
                v6 := GenS1{F0: 33, F1: 94, F2: 42}
                v7 := GenS2{F0: (-2.25 <= 0.1), F1: (-(fi4)), F2: "foo bar"}
                v8 := GenS3{F0: (v1 || v1)}
                _ = fi5
                _ = v6
                _ = v7
                _ = v8
            }
            _ = fi4
        }
    }
    var v9 uint64 = 740
    var v10 float64 = -2.25
    for fi11 := 0; fi11 < 7; fi11++ {
        delete(v3, "\n\t")
        var v12 int = fi11
        _ = fi11
        _ = v12
    }
    var v13 uint16 = 63
    v14 := GenS4{F0: true}
    var v15 float64 = v10
    v16 := GenS1{F0: 2, F1: 255, F2: 828}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v9
}
