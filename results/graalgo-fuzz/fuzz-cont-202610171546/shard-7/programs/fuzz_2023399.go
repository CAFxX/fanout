package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int }
type GenS2 struct { F0 uint16 }

func main() {
    var v1 uint64 = 417
    var v2 float32 = 3.14159
    var v3 uint8 = 246
    var v4 bool = false
    var v5 string = "a"
    v6 := []int{}
    for fi7 := 0; fi7 < 16; fi7++ {
        _ = len(v6)
        v8 := GenS1{F0: 1, F1: -721}
        _ = fi7
        _ = v8
    }
    v1 -= v1
    if v4 {
        v2 += ((float32(1.5) / (v2 + float32((-2.25 + 1e10)))) + (3.14159 * (1.5 * 1e10)))
    }
    switch (v5 + ((v5 + (v5 + "hello")) + (v5 + (v5 + "a")))) {
    case "a":
        v5 = "a"
    case "hello":
        for fi9 := 0; fi9 < 6; fi9++ {
            var v10 int32 = 234
            var v11 float64 = 0.1
            _ = fi9
            _ = v10
            _ = v11
        }
    }
    var v12 float64 = (((((1e10 / float64(v2)) + ((1e10 - 1e10) - (3.14159 - 0.0))) - (1e10 / -2.25)) * float64(v2)) / 0.1)
    if (v2 > (v2 / (v2 + 1.5))) {
        var v13 int = 96
        var v14 int16 = 32767
        _ = v13
        _ = v14
    }
    v5 = v5
    fmt.Println(v3, v4)
    for fi15 := 0; fi15 < 8; fi15++ {
        var v16 uint32 = 1
        var v17 uint = 1
        v18 := GenS2{F0: 32767}
        _ = fi15
        _ = v16
        _ = v17
        _ = v18
    }
    if false {
        if len(v6) > 0 {
            v6[0] = -903
        }
        var v19 bool = (false || v4)
        _ = v19
    } else {
        v12 = v12
    }
    if (!v4) {
        _ = len(v6)
        fmt.Println(v3)
        if (1 < -841) {
            var v20 int32 = 0
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            var v21 uint = 18446744073709551615
            _ = v20
            _ = v21
        }
    } else {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
