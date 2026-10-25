package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64; F2 int16 }

func main() {
    var v1 float64 = float64(1.5)
    v2 := []uint32{}
    var v3 int64 = 800
    var v4 int32 = -1
    fmt.Println(v1)
    if true {
        for fi5 := 0; fi5 < 1; fi5++ {
            fmt.Println(v1)
            v2 = append(v2, 86)
            _ = fi5
        }
    }
    fmt.Println(v3, v1, v1)
    var v6 string = "foo bar"
    var v7 float32 = ((1e10 * (0.0 - (-2.25 - 0.0))) - float32(float64(float32(v1))))
    v8 := GenS1{F0: "a", F1: 18446744073709551615, F2: 1000}
    if (v1 < (v1 - v1)) {
        v2 = append(v2, 4294967295)
    } else {
        v6 = "a"
        if ((!(1 >= -15)) && true) {
            v9 := GenS1{F0: v6, F1: 247, F2: 1}
            fmt.Println(v1, v1)
            var v10 bool = false
            _ = v10
            _ = v9
        } else {
            v11 := GenS1{F0: v6, F1: 42, F2: 587}
            var v12 int = 377
            _ = v11
            _ = v12
        }
    }
    fmt.Println(v1, v7)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
}
