package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool }

func main() {
    v1 := []int{32768, 32767}
    v2 := make(map[string]uint32)
    var v3 uint8 = 42
    var v4 uint64 = 299
    var v5 string = "\n\t"
    v6 := GenS1{F0: (^(v4)), F1: true}
    v4 += 599
    if (((v5 + ("" + v5)) + v5) > v5) {
        v7 := GenS1{F0: (v4 >> ((0) & 63)), F1: false}
        _ = v7
    }
    var v8 uint64 = 4294967295
    var v9 float64 = (((1.5 + -2.25) * 0.0) / 0.1)
    for fi10 := 0; fi10 < 3; fi10++ {
        var v11 uint64 = v8
        for fi12 := 0; fi12 < 18; fi12++ {
            if len(v1) > 0 {
                v1[0] = 42
            }
            _ = fi12
        }
        _ = fi10
        _ = v11
    }
    _ = len(v2)
    var v13 int32 = -732
    var v14 uint = 255
    fmt.Println(v3, v14)
    v15 := GenS1{F0: v4, F1: false}
    switch 32767 {
    case -447:
        for fi16 := 0; fi16 < 1; fi16++ {
            var v17 float32 = -2.25
            fmt.Println(v3, v8)
            if false {
                fmt.Println(v9)
            }
            _ = fi16
            _ = v17
        }
    case 127:
        v8 += v8
    default:
        var v18 uint8 = v3
        _ = v18
    }
    if len(v1) > 0 {
        v1[0] = 128
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
