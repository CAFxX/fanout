package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8 }
type GenS2 struct { F0 int; F1 int }
type GenS3 struct { F0 uint16; F1 uint8 }
func genFn0(p0 uint, p1 int) uint64 {
    var v4 float64 = float64((float32(-2.25) - (0.1 + 3.14159)))
    v5 := GenS1{F0: true, F1: 0}
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 2
}
func genFn1() uint8 {
    var v6 uint32 = 65535
    var v7 int = 746
    v7 = (v7 & v7)
    _ = v6
    _ = v7
    return 30
}

func main() {
    var v1 uint16 = 562
    v2 := make(map[uint32]int)
    var v3 uint16 = 2
    var v8 float32 = -2.25
    v9 := GenS1{F0: (608 != 163), F1: 0}
    var v10 uint32 = 0
    v11 := GenS1{F0: false, F1: 92}
    v1 = 58
    var v12 uint64 = 445
    if true {
        for fi13 := 0; fi13 < 8; fi13++ {
            var v14 int8 = 0
            _ = fi13
            _ = v14
        }
        for fi15 := 0; fi15 < 7; fi15++ {
            if true {
                switch "\n\t" {
                case "":
                    var v16 uint8 = uint8(v1)
                    _ = v16
                case "\n\t":
                    fmt.Println(v8, v12, v8)
                default:
                    v12 = (v12 >> ((65535) & 63))
                }
                v8 += v8
            } else {
                for fi17 := 0; fi17 < 4; fi17++ {
                    v18 := v2[v10]
                    if (true && true) {
                        var v19 uint = 65535
                        v12 -= (v12 >> (((^(v19))) & 63))
                        _ = v19
                    }
                    var v20 float32 = (v8 * 3.14159)
                    _ = fi17
                    _ = v18
                    _ = v20
                }
                for fi21 := 0; fi21 < 6; fi21++ {
                    v12 = (v12 * v12)
                    _ = fi21
                }
            }
            v22 := []int{128, (^(fi15))}
            if false {
                v3 = 925
            }
            _ = fi15
            _ = v22
        }
        var v23 int16 = -336
        _ = v23
    } else {
        var v24 int64 = -1
        _ = v24
    }
    v25 := v2[4294967295]
    for fi26 := 0; fi26 < 14; fi26++ {
        v27 := GenS2{F0: (fi26 | (fi26 | v25)), F1: v25}
        var v28 int64 = 739
        if true {
            fmt.Println(v8, v25, v25)
            if true {
                fmt.Println(v3)
                v25 = fi26
            }
            var v29 int16 = 42
            _ = v29
        }
        _ = fi26
        _ = v27
        _ = v28
    }
    v30 := GenS2{F0: -2147483648, F1: v25}
    v31 := GenS3{F0: 149, F1: 2}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v25
    _ = v3
    _ = v30
    _ = v31
    _ = v8
    _ = v9
}
