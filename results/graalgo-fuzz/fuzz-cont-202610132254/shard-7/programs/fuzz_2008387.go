package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint64; F1 string }
func genFn0(p0 int32, p1 int8) uint16 {
    var v4 bool = false
    v5 := []int{0}
    v4 = (873 != 18446744073709551615)
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 0
}
func genFn1(p0 int64, p1 uint, p2 uint32) uint {
    var v6 uint8 = 255
    p1 = p1
    var v7 float32 = -2.25
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return (p1 - (p1 << (((p1 - 255)) & 63)))
}

func main() {
    var v1 uint32 = 0
    v2 := []uint32{}
    v3 := make(map[uint32]int)
    var v8 uint8 = 255
    if false {
        v9 := GenS1{F0: 42}
        switch 32768 {
        case 859:
            var v10 uint8 = 128
            _ = v10
        case 475:
            v11 := GenS1{F0: int16(v8)}
            _ = v11
        default:
            v12 := GenS2{F0: 0, F1: "\n\t"}
            _ = v12
        }
        _ = v9
    }
    if (!false) {
        v8 = (v8 - v8)
        var v13 uint = 530
        fmt.Println(v8)
        _ = v13
    }
    v14 := GenS1{F0: 999}
    var v15 uint8 = v8
    v8 = v15
    var v16 uint8 = (v15 | (v8 & 36))
    v17 := v3[(v1 % ((v1) | 1))]
    var v18 float64 = (1.5 * -2.25)
    if (!(18446744073709551615 == 255)) {
        fmt.Println(v17)
        for fi19 := 0; fi19 < 2; fi19++ {
            var v20 int8 = -79
            _ = fi19
            _ = v20
        }
    }
    var v21 int32 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v3
    _ = v8
}
