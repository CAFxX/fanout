package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int }

func main() {
    var v1 int8 = 100
    var v2 float64 = 0.1
    var v3 uint = 466
    for fi4 := 0; fi4 < 8; fi4++ {
        v5 := []int{}
        v3 -= 687
        var v6 uint64 = 4294967295
        _ = fi4
        _ = v5
        _ = v6
    }
    fmt.Println(v2, v3)
    for fi7 := 0; fi7 < 6; fi7++ {
        fmt.Println(v1)
        v8 := GenS1{F0: false, F1: (fi7 & 128)}
        _ = fi7
        _ = v8
    }
    v9 := []int{}
    fmt.Println(v3, v3, v2)
    v3 = 855
    var v10 uint32 = 65535
    fmt.Println(v3)
    if ((float64(float32(0.0)) - 1e10) >= (v2 + v2)) {
        _ = len(v9)
    }
    if true {
        var v11 string = (("\n\t" + "hello") + "foo bar")
        _ = v11
    }
    v2 = v2
    var v12 uint = (^(v3))
    var v13 int8 = (v1 % (((v1 * (v1 - v1))) | 1))
    var v14 uint64 = 42
    v14 += v14
    v12 = (v3 / ((v12) | 1))
    fmt.Println(v14, v10, v10)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v9
}
