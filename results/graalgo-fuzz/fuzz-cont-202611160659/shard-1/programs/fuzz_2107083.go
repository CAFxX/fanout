package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string; F2 uint16 }
type GenS2 struct { F0 uint32 }

func main() {
    var v1 string = ("" + "foo bar")
    var v2 uint8 = 73
    var v3 uint8 = (v2 ^ v2)
    v4 := []uint32{}
    v1 = ("foo bar" + (v1 + (v1 + v1)))
    var v5 uint16 = 32767
    for fi6 := 0; fi6 < 18; fi6++ {
        var v7 uint16 = v5
        _ = fi6
        _ = v7
    }
    v8 := GenS1{F0: -1, F1: v1, F2: 2}
    v9 := GenS1{F0: -9223372036854775808, F1: "", F2: v5}
    var v10 uint64 = 42
    var v11 int64 = 833
    if (0 >= 4) {
        v12 := GenS1{F0: (v11 - v11), F1: "a", F2: (v5 * v5)}
        var v13 float64 = (0.0 / float64(1.5))
        _ = v12
        _ = v13
    }
    v14 := GenS2{F0: 662}
    v15 := GenS2{F0: 690}
    if len(v4) > 0 {
        v4[0] = 668
    }
    var v16 uint16 = v5
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
