package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32 }
type GenS2 struct { F0 int16; F1 string }

func main() {
    v1 := []uint32{2, 1, 0}
    var v2 float64 = 1.5
    var v3 bool = ((v2 - (v2 * 1.5)) > v2)
    var v4 float32 = (-2.25 - 3.14159)
    for fi5 := 0; fi5 < 1; fi5++ {
        for fi6 := 0; fi6 < 12; fi6++ {
            var v7 int8 = 89
            _ = fi6
            _ = v7
        }
        _ = fi5
    }
    for fi8 := 0; fi8 < 9; fi8++ {
        v9 := GenS1{F0: 901, F1: 291}
        var v10 float64 = v2
        _ = fi8
        _ = v10
        _ = v9
    }
    var v11 int8 = 42
    for fi12 := 0; fi12 < 15; fi12++ {
        v2 = v2
        v13 := GenS2{F0: -1, F1: "hello"}
        _ = len(v1)
        _ = fi12
        _ = v13
    }
    if (250 < 42) {
        var v14 int64 = -753
        _ = v14
    }
    v15 := GenS2{F0: 261, F1: "foo bar"}
    for fi16 := 0; fi16 < 11; fi16++ {
        v11 -= (-(v11))
        _ = fi16
    }
    v3 = false
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
