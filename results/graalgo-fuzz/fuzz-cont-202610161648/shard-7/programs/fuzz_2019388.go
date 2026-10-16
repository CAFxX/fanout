package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string; F2 int16 }

func main() {
    var v1 float64 = 0.0
    var v2 int32 = 578
    var v3 int64 = -1
    var v4 uint64 = 2
    for fi5 := 0; fi5 < 17; fi5++ {
        v6 := GenS1{F0: (^(fi5)), F1: "hello", F2: 1000}
        v4 += (v4 << ((962) & 63))
        fmt.Println(fi5, v1, v4)
        _ = fi5
        _ = v6
    }
    fmt.Println(v1)
    v2 -= v2
    for fi7 := 0; fi7 < 19; fi7++ {
        var v8 int16 = 1000
        var v9 uint64 = v4
        var v10 float32 = ((0.1 - (0.1 + ((0.0 * 0.0) - 3.14159))) + -2.25)
        _ = fi7
        _ = v10
        _ = v8
        _ = v9
    }
    fmt.Println(v1, v1, v4)
    for fi11 := 0; fi11 < 16; fi11++ {
        if true {
            var v12 uint16 = 32767
            v13 := GenS1{F0: fi11, F1: "hello", F2: 32767}
            _ = v12
            _ = v13
        }
        v4 = (v4 * (^(v4)))
        v1 = v1
        _ = fi11
    }
    if (0.1 <= 0.1) {
        switch "" {
        case "a":
            fmt.Println(v1, v4)
        case "":
            if ((127 == 2) && true) {
                var v14 int32 = v2
                _ = v14
            }
        default:
            v15 := []int{}
            _ = v15
        }
        var v16 uint64 = (v4 % (((v4 * v4)) | 1))
        v3 = v3
        _ = v16
    }
    fmt.Println(v1)
    v17 := []int{-1}
    var v18 bool = true
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
