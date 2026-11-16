package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint8; F1 int8; F2 int16 }
type GenS3 struct { F0 int32; F1 uint32; F2 uint }
type GenS4 struct { F0 uint8; F1 int8; F2 bool }
type GenS5 struct { F0 int16; F1 int32; F2 uint16 }
func genFn0() uint64 {
    var v4 uint64 = 1
    var v5 int32 = -2147483648
    _ = v4
    _ = v5
    return v4
}

func main() {
    v1 := []int{2147483647, 4294967295, 127}
    var v2 uint32 = 1
    var v3 int32 = -1
    for fi6 := 0; fi6 < 4; fi6++ {
        var v7 uint64 = 18446744073709551615
        v8 := GenS1{F0: v3}
        fmt.Println(fi6, v2, v7)
        _ = fi6
        _ = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 7; fi9++ {
        v3 -= int32(fi9)
        switch "hello" {
        case "a":
            _ = len(v1)
        case "foo bar":
            if len(v1) > 0 {
                v1[0] = (^(fi9))
            }
        case "hello":
            if (-1 != 554) {
                if (1.5 != (1e10 * 0.0)) {
                    if (!(((3.14159 >= float32(0.0)) || (((("" + "\n\t") + "\n\t") == "hello") || false)) || true)) {
                        v10 := GenS2{F0: 0, F1: 42, F2: 42}
                        fmt.Println(v3, v3, v3)
                        _ = v10
                    }
                }
                if true {
                    v11 := GenS1{F0: 0}
                    for fi12 := 0; fi12 < 9; fi12++ {
                        v2 = (v2 * (v2 / (((v2 % (((v2 / (((^(v2))) | 1))) | 1))) | 1)))
                        v2 += (v2 * 1)
                        _ = len(v1)
                        var v13 int8 = -128
                        _ = fi12
                        _ = v13
                    }
                    switch 837 {
                    case 127:
                        var v14 float64 = 0.0
                        _ = v14
                    }
                    _ = v11
                }
            } else {
                var v15 float32 = (-2.25 * float32(1e10))
                _ = v15
            }
        }
        var v16 int32 = v3
        _ = fi9
        _ = v16
    }
    fmt.Println(v3, v2)
    _ = len(v1)
    var v17 uint = 0
    switch -137 {
    case 4294967295:
        fmt.Println(v17)
    case -1:
        var v18 int8 = 127
        _ = v18
    case 128:
        v19 := GenS3{F0: (v3 / ((v3) | 1)), F1: 4294967295, F2: (v17 / (((^(v17))) | 1))}
        _ = v19
    }
    v2 = (v2 >> ((2) & 63))
    v20 := GenS4{F0: 2, F1: 100, F2: (2147483647 != 741)}
    v2 = (^(v2))
    v2 += (^(v2))
    if false {
        var v21 float32 = float32(3.14159)
        _ = v21
    }
    v22 := GenS5{F0: 1, F1: v3, F2: 32767}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v17
    _ = v2
    _ = v20
    _ = v22
    _ = v3
}
