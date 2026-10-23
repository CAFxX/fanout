package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16; F2 uint16 }
type GenS2 struct { F0 uint8; F1 uint64 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int; F1 int32 }

func main() {
    var v1 uint8 = 255
    var v2 int8 = 100
    v3 := make(map[int]string)
    var v4 uint64 = 18446744073709551615
    var v5 int32 = 42
    for fi6 := 0; fi6 < 20; fi6++ {
        var v7 string = ("a" + "\n\t")
        _ = fi6
        _ = v7
    }
    v8 := []int{}
    v9 := GenS1{F0: 323, F1: 103, F2: 81}
    if len(v8) > 0 {
        v8[0] = -158
    }
    switch 42 {
    case 1:
        var v10 int8 = v2
        _ = v10
    case 42:
        for fi11 := 0; fi11 < 6; fi11++ {
            var v12 float32 = 0.1
            var v13 int8 = v2
            _ = fi11
            _ = v12
            _ = v13
        }
    case -758:
        v14 := GenS2{F0: v1, F1: 42}
        _ = v14
    }
    if false {
        var v15 int32 = (v5 * v5)
        v2 += (v2 * 42)
        var v16 uint64 = 42
        _ = v15
        _ = v16
    }
    v1 += 10
    var v17 string = ("hello" + "")
    var v18 uint = 1
    var v19 uint16 = 32767
    for fi20 := 0; fi20 < 13; fi20++ {
        var v21 int8 = v2
        v22 := GenS1{F0: (v4 << (((v18 + v18)) & 63)), F1: 1, F2: v19}
        _ = fi20
        _ = v21
        _ = v22
    }
    for fi23 := 0; fi23 < 15; fi23++ {
        var v24 float64 = (0.1 * 3.14159)
        v25 := GenS3{F0: v17}
        for fi26 := 0; fi26 < 16; fi26++ {
            if false {
                v3[(fi26 + -224)] = (v17 + v17)
            } else {
                var v27 bool = false
                _ = v27
            }
            v3[(^(fi26))] = v17
            var v28 uint = (v18 % (((v18 - (^(v18)))) | 1))
            if ((v19 < 189) || true) {
                if (!false) {
                    v29 := GenS4{F0: (fi23 / ((fi26) | 1)), F1: v5}
                    _ = v29
                }
            }
            _ = fi26
            _ = v28
        }
        _ = fi23
        _ = v24
        _ = v25
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
