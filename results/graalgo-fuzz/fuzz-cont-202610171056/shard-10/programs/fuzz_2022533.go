package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 string; F1 int64 }

func main() {
    var v1 uint16 = 1
    var v2 int = 893
    var v3 float32 = 0.1
    fmt.Println(v2, v1, v2)
    fmt.Println(v1)
    v4 := GenS1{F0: ("\n\t" + "a")}
    v3 += v3
    for fi5 := 0; fi5 < 16; fi5++ {
        v3 = (v3 - (((v3 * float32(1.5)) * (float32(-2.25) - float32(float64((v3 / v3))))) - v3))
        fmt.Println(v2, fi5)
        v6 := []int{fi5, fi5}
        _ = fi5
        _ = v6
    }
    var v7 int64 = 9223372036854775807
    v8 := []int{-9}
    v9 := GenS2{F0: "", F1: (-(v7))}
    v10 := GenS2{F0: "a", F1: v7}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
