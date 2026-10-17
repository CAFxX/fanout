package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool }

func main() {
    var v1 int8 = -91
    v2 := make(map[string]int)
    v3 := []uint32{1}
    var v4 uint64 = 18446744073709551615
    v5 := GenS1{F0: 0, F1: true}
    v4 = v4
    v6 := GenS1{F0: 170, F1: (false || (!false))}
    for fi7 := 0; fi7 < 2; fi7++ {
        v1 = (v1 * (v1 * (v1 << ((731) & 63))))
        v8 := GenS1{F0: 1, F1: false}
        _ = fi7
        _ = v8
    }
    v9 := GenS1{F0: 221, F1: false}
    var v10 uint32 = 65535
    var v11 string = "a"
    _ = len(v3)
    if false {
        _ = len(v2)
        var v12 uint32 = (v10 / ((900) | 1))
        _ = v12
    }
    v1 = v1
    fmt.Println(v10)
    if true {
        if (!(v10 == v10)) {
            v3 = append(v3, (^(v10)))
            var v13 uint64 = (v4 - (v4 | v4))
            _ = v13
        }
        if true {
            v10 = (v10 - v10)
            v1 += 127
        } else {
            v1 = v1
            var v14 float64 = 0.1
            var v15 int16 = 42
            _ = v14
            _ = v15
        }
    } else {
        fmt.Println(v1, v4)
    }
    var v16 int16 = 151
    _ = len(v2)
    for fi17 := 0; fi17 < 17; fi17++ {
        v1 += (v1 ^ (-(v1)))
        _ = fi17
    }
    v16 = -32768
    v4 = 196
    v1 = 127
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
