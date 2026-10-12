package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64; F2 uint16 }
type GenS2 struct { F0 int8; F1 int; F2 uint64 }
func genFn0(p0 uint, p1 int8, p2 int64) uint {
    fmt.Println(p1, p0)
    var v4 int64 = (p2 - (p2 * p2))
    var v5 uint32 = 1
    var v6 string = "a"
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    return (p0 >> ((2) & 63))
}
func genFn1(p0 int32, p1 bool, p2 int32) uint64 {
    if p1 {
        if true {
            var v7 int16 = -32768
            v8 := []int{32767, 32768}
            for fi9 := 0; fi9 < 18; fi9++ {
                v10 := GenS1{F0: 42, F1: -503, F2: 2}
                p1 = p1
                _ = fi9
                _ = v10
            }
            _ = v7
            _ = v8
        }
        fmt.Println(p2)
    }
    v11 := []int{}
    fmt.Println(p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return 698
}

func main() {
    var v1 float64 = (0.1 / 1e10)
    var v2 int = -2147483648
    var v3 uint = 0
    v1 = 0.0
    v3 = 18446744073709551615
    v1 = 1e10
    v12 := GenS1{F0: 1, F1: 0, F2: 517}
    if false {
        v3 += (v3 << (((^(v3))) & 63))
        switch v2 {
        case 659:
            var v13 int = 415
            _ = v13
        case -2147483648:
            v14 := []int{v2}
            _ = v14
        case 1:
            for fi15 := 0; fi15 < 16; fi15++ {
                v2 = fi15
                v1 = v1
                var v16 uint8 = 160
                _ = fi15
                _ = v16
            }
        }
    } else {
        fmt.Println(v2)
    }
    v17 := GenS1{F0: 42, F1: 638, F2: uint16(v3)}
    v1 = ((v1 * (((0.1 - -2.25) / (1e10 + v1)) + (v1 + ((0.0 / v1) - v1)))) * (v1 - v1))
    var v18 bool = false
    var v19 string = "\n\t"
    v18 = (!v18)
    v20 := GenS2{F0: -37, F1: v2, F2: 849}
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
}
