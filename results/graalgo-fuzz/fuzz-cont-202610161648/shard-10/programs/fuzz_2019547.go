package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16; F2 int32 }
type GenS2 struct { F0 int8; F1 int8 }
type GenS3 struct { F0 uint8 }

func main() {
    v1 := []uint32{1}
    v2 := []uint32{0, 65535}
    var v3 uint16 = 510
    v4 := make(map[int]string)
    var v5 uint8 = 235
    fmt.Println(v5)
    var v6 float64 = -2.25
    fmt.Println(v6, v3)
    fmt.Println(v5)
    for fi7 := 0; fi7 < 6; fi7++ {
        v8 := GenS1{F0: fi7, F1: 255, F2: -1}
        _ = fi7
        _ = v8
    }
    fmt.Println(v5)
    if false {
        v9 := GenS2{F0: -1, F1: 100}
        if false {
            for fi10 := 0; fi10 < 17; fi10++ {
                v6 -= (float64(0.0) * float64(3.14159))
                if false {
                    if (61 >= -128) {
                        var v11 int8 = 0
                        v4[-341] = "\n\t"
                        _ = v11
                    }
                    v3 = (v3 / ((v3) | 1))
                } else {
                    _ = len(v1)
                }
                var v12 int16 = -136
                _ = fi10
                _ = v12
            }
        } else {
            if (true && (42 < 127)) {
                var v13 int8 = -1
                _ = v13
            }
        }
        _ = v9
    } else {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        var v14 string = "\n\t"
        v1 = append(v1, 411)
        _ = v14
    }
    switch 127 {
    case -391:
        v15 := GenS3{F0: 233}
        _ = v15
    default:
        var v16 float64 = v6
        _ = v16
    }
    v5 -= 2
    var v17 int32 = -1
    if true {
        var v18 uint8 = (^(v5))
        var v19 uint8 = (v5 * (v18 & 128))
        var v20 int8 = 100
        _ = v18
        _ = v19
        _ = v20
    }
    v21 := GenS1{F0: -9223372036854775808, F1: v3, F2: 42}
    _ = len(v2)
    var v22 uint64 = 555
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
