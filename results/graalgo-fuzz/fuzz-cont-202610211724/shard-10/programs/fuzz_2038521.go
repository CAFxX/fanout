package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32 }
type GenS2 struct { F0 int16; F1 uint8 }

func main() {
    var v1 string = "a"
    v2 := []uint32{2147483647, 876}
    var v3 int32 = 425
    fmt.Println(v1, v3)
    fmt.Println(v1, v1)
    var v4 uint32 = 2147483647
    for fi5 := 0; fi5 < 20; fi5++ {
        for fi6 := 0; fi6 < 8; fi6++ {
            v7 := GenS1{F0: 24, F1: v3}
            _ = fi6
            _ = v7
        }
        _ = fi5
    }
    for fi8 := 0; fi8 < 16; fi8++ {
        fmt.Println(v3, v1, v4)
        _ = fi8
    }
    v3 = (v3 / (((v3 | -760)) | 1))
    if ((int(v3) < 32768) || (v1 != v1)) {
        var v9 float64 = 0.0
        v4 = (^(v4))
        var v10 int = -304
        _ = v10
        _ = v9
    } else {
        v4 -= 514
        for fi11 := 0; fi11 < 8; fi11++ {
            v1 = v1
            _ = fi11
        }
    }
    fmt.Println(v1)
    v4 += 372
    var v12 float64 = 0.1
    var v13 uint32 = (v4 - v4)
    if true {
        var v14 uint = 386
        fmt.Println(v1)
        _ = v14
    } else {
        v15 := GenS2{F0: -333, F1: 1}
        v12 = -2.25
        _ = v15
    }
    if (1000 == 1) {
        var v16 uint64 = 0
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = v16
    } else {
        v12 = (v12 / ((0.0 * float64((float32((0.0 * float64((1.5 - 0.0)))) - 0.1))) + 1e10))
    }
    var v17 int8 = -64
    var v18 float32 = float32(v12)
    var v19 uint32 = 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
