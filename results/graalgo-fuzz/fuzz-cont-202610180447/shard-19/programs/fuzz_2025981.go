package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16; F2 uint8 }
func genFn0(p0 string, p1 int32) string {
    var v7 int64 = 1
    fmt.Println(v7, p1)
    v8 := []int{1, 127, -1}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return p0
}

func main() {
    var v1 uint8 = 0
    v2 := []uint32{}
    var v3 int32 = 0
    v4 := []int{-748, 127, -2147483648, 4294967295}
    var v5 float32 = ((float32(0.1) + (0.0 + 0.1)) - 0.1)
    var v6 int16 = 32767
    switch 986 {
    case 128:
        _ = len(v2)
    case 127:
        v1 -= v1
    case 548:
        for fi9 := 0; fi9 < 7; fi9++ {
            switch -1 {
            case -800:
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            case -9223372036854775808:
                for fi10 := 0; fi10 < 9; fi10++ {
                    var v11 uint8 = (v1 & v1)
                    _ = fi10
                    _ = v11
                }
            default:
                for fi12 := 0; fi12 < 15; fi12++ {
                    fmt.Println(v3, v3, v6)
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = fi12
                }
            }
            v6 += (v6 - v6)
            if len(v4) > 0 {
                v4[0] = fi9
            }
            _ = fi9
        }
    default:
        v3 = -632
    }
    v1 = 2
    var v13 float64 = 1e10
    var v14 int32 = v3
    switch "hello" {
    case "\n\t":
        v15 := GenS1{F0: true, F1: 0, F2: 224}
        _ = v15
    }
    var v16 int = -212
    var v17 int64 = 0
    v13 = v13
    var v18 bool = ((-62 >= -1) || true)
    v16 += (v16 * v16)
    v19 := GenS1{F0: (v5 >= (1.5 * v5)), F1: 42, F2: (v1 & 1)}
    if v18 {
        var v20 int8 = -74
        var v21 uint32 = 65535
        _ = v20
        _ = v21
    } else {
        var v22 uint64 = 4294967295
        _ = v22
    }
    v1 = (^(v1))
    v23 := GenS1{F0: v18, F1: 32767, F2: (^(v1))}
    v16 = (v16 & -1)
    var v24 float32 = 0.1
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
