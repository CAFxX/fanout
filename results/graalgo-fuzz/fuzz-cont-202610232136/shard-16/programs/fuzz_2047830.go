package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string }

func main() {
    var v1 string = "hello"
    var v2 uint = 1
    var v3 int16 = -1
    v2 -= (v2 >> (((^(v2))) & 63))
    var v4 int32 = 1
    v5 := GenS1{F0: 0, F1: ""}
    var v6 float64 = 1e10
    if false {
        v7 := []int{-607}
        _ = v7
    }
    v8 := []int{590}
    var v9 string = ""
    v8 = append(v8, 42)
    var v10 float64 = -2.25
    v11 := GenS1{F0: 100, F1: v1}
    if len(v8) > 0 {
        v8[0] = 458
    }
    fmt.Println(v4, v6)
    fmt.Println(v6, v2, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
