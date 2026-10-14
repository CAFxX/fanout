package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string }
type GenS2 struct { F0 bool; F1 uint16 }
type GenS3 struct { F0 uint64; F1 uint8 }

func main() {
    var v1 int16 = -448
    v2 := make(map[int]bool)
    var v3 uint64 = 4294967295
    var v4 float32 = 1.5
    var v5 int = 32767
    v6 := make(map[string]string)
    v3 = v3
    v5 = v5
    v1 -= (-(v1))
    for fi7 := 0; fi7 < 4; fi7++ {
        for fi8 := 0; fi8 < 13; fi8++ {
            var v9 int8 = -119
            _ = fi8
            _ = v9
        }
        v4 += (v4 + (v4 / (0.1 - v4)))
        v10 := []int{}
        _ = fi7
        _ = v10
    }
    var v11 int32 = 2147483647
    v12 := GenS1{F0: 0, F1: ""}
    for fi13 := 0; fi13 < 12; fi13++ {
        v14 := GenS2{F0: true, F1: 42}
        v1 -= v1
        _ = fi13
        _ = v14
    }
    for fi15 := 0; fi15 < 15; fi15++ {
        for fi16 := 0; fi16 < 11; fi16++ {
            var v17 uint32 = 42
            _ = fi16
            _ = v17
        }
        _ = fi15
    }
    if (!false) {
        var v18 float32 = (v4 / ((v4 + float32(float64(0.0))) / v4))
        v19 := []int{}
        _ = v18
        _ = v19
    } else {
        if (0 != v11) {
            v20 := GenS3{F0: v3, F1: 1}
            var v21 int8 = -1
            _ = v20
            _ = v21
        }
    }
    if false {
        v22 := v2[v5]
        _ = v22
    }
    if false {
        v11 -= (-(v11))
        v6[""] = ("foo bar" + "a")
    }
    v5 = v5
    if true {
        var v23 float64 = 3.14159
        _ = v23
    } else {
        var v24 float64 = (0.1 - 0.0)
        var v25 uint = 42
        fmt.Println(v4)
        _ = v24
        _ = v25
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
