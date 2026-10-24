package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16 }
type GenS2 struct { F0 int32; F1 uint8; F2 uint8 }
func genFn0(p0 int64, p1 int16) uint8 {
    v4 := GenS1{F0: 65535, F1: 2}
    var v5 int64 = p0
    for fi6 := 0; fi6 < 1; fi6++ {
        v7 := GenS2{F0: 42, F1: 98, F2: 2}
        _ = fi6
        _ = v7
    }
    var v8 uint = 255
    switch 32767 {
    case 127:
        v9 := GenS1{F0: 2, F1: 42}
        _ = v9
    case -899:
        fmt.Println(v5)
    default:
        p0 = v5
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v8
    return 2
}
func genFn1(p0 int, p1 bool) bool {
    var v10 uint16 = 1
    if true {
        var v11 float32 = 0.0
        _ = v11
    }
    _ = p0
    _ = p1
    _ = v10
    return (!(p1 || false))
}

func main() {
    var v1 int = 758
    v2 := make(map[uint32]bool)
    var v3 float32 = -2.25
    v12 := v2[688]
    var v13 float64 = (0.1 * 3.14159)
    switch ("" + "a") {
    case "\n\t":
        v13 = ((3.14159 * v13) + v13)
    }
    for fi14 := 0; fi14 < 19; fi14++ {
        fmt.Println(v3, v13, v1)
        var v15 int = (v1 & (fi14 % ((fi14) | 1)))
        _ = fi14
        _ = v15
    }
    for fi16 := 0; fi16 < 12; fi16++ {
        if (736 < 964) {
            switch "\n\t" {
            case "hello":
                for fi17 := 0; fi17 < 18; fi17++ {
                    v18 := GenS2{F0: 1, F1: 128, F2: 255}
                    v12 = false
                    var v19 float32 = (0.0 + ((v3 + (float32(v13) - v3)) - (0.0 / v3)))
                    var v20 int64 = 377
                    _ = fi17
                    _ = v18
                    _ = v19
                    _ = v20
                }
            default:
                if (v3 > float32((v13 * float64(v3)))) {
                    v21 := GenS2{F0: -382, F1: 2, F2: genFn0(42, 1000)}
                    var v22 float32 = 0.0
                    _ = v21
                    _ = v22
                } else {
                    var v23 uint = 65535
                    v24 := GenS2{F0: 42, F1: 127, F2: 127}
                    _ = v23
                    _ = v24
                }
            }
            v3 = 1.5
        } else {
            for fi25 := 0; fi25 < 18; fi25++ {
                if false {
                    if ((127 == 42) || (1 == 172)) {
                        var v26 int16 = -1
                        fmt.Println(v26, v3)
                        _ = v26
                    } else {
                        var v27 uint64 = 115
                        v3 = (-2.25 - (1.5 + (v3 / (v3 + v3))))
                        _ = v27
                    }
                }
                v12 = v12
                _ = fi25
            }
            v28 := []int{4294967295, 127}
            _ = v28
        }
        _ = fi16
    }
    v29 := []int{}
    var v30 int64 = 9223372036854775807
    if v12 {
        var v31 uint = 131
        _ = v31
    }
    fmt.Println(v1, v30)
    _ = len(v29)
    v1 += (v1 * (v1 - 963))
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v29
    _ = v3
    _ = v30
}
