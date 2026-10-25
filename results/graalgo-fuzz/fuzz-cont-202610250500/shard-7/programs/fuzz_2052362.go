package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string; F2 string }
type GenS2 struct { F0 int64; F1 int; F2 string }
type GenS3 struct { F0 int32; F1 uint16; F2 int64 }
type GenS4 struct { F0 uint; F1 bool; F2 bool }
type GenS5 struct { F0 int8 }

func main() {
    var v1 string = "hello"
    v2 := make(map[int]int)
    v3 := []uint32{978, 2147483647, 1, 0}
    v4 := GenS1{F0: 1, F1: (v1 + ((((v1 + "hello") + v1) + v1) + v1)), F2: ""}
    if len(v3) > 0 {
        v3[0] = 42
    }
    if false {
        v1 = (v1 + (v1 + "hello"))
    } else {
        v5 := GenS2{F0: -9223372036854775808, F1: 2147483647, F2: "a"}
        _ = v5
    }
    v1 = v1
    v6 := GenS3{F0: 2147483647, F1: 42, F2: 9223372036854775807}
    switch "" {
    case "\n\t":
        var v7 uint32 = 299
        _ = v7
    case "a":
        var v8 int8 = 100
        _ = v8
    case "hello":
        v3 = append(v3, 0)
    }
    v9 := GenS4{F0: 428, F1: false, F2: (100 >= 100)}
    var v10 bool = false
    fmt.Println(v10, v1)
    for fi11 := 0; fi11 < 18; fi11++ {
        delete(v2, fi11)
        v12 := GenS5{F0: -50}
        fmt.Println(fi11, fi11)
        _ = fi11
        _ = v12
    }
    switch "\n\t" {
    case "hello":
        var v13 string = "hello"
        _ = v13
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v9
}
