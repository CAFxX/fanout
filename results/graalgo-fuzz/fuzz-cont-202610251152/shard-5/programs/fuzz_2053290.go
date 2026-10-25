package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16 }
type GenS2 struct { F0 uint64; F1 uint64 }
func genFn0(p0 uint8, p1 int16, p2 uint32) int32 {
    v5 := []int{32768}
    for fi6 := 0; fi6 < 2; fi6++ {
        if len(v5) > 0 {
            v5[0] = fi6
        }
        v7 := GenS1{F0: true, F1: 1}
        _ = fi6
        _ = v7
    }
    if (((p0 * 255) < p0) || false) {
        var v8 float32 = 3.14159
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return -2147483648
}

func main() {
    v1 := make(map[int]uint32)
    var v2 uint32 = 567
    var v3 string = ("a" + ("" + "a"))
    var v4 int = -693
    for fi9 := 0; fi9 < 4; fi9++ {
        for fi10 := 0; fi10 < 15; fi10++ {
            v11 := []int{(-(v4))}
            switch "a" {
            case "\n\t":
                if ((false || false) || false) {
                    if false {
                        var v12 uint = uint(v2)
                        _ = v12
                    } else {
                        v4 -= (^(fi9))
                    }
                }
            case "hello":
                var v13 float32 = (1.5 + (1.5 * float32((0.0 - (((1e10 * ((float64(0.1) / 1.0) - 0.0)) + 3.14159) / 1e10)))))
                _ = v13
            case "foo bar":
                var v14 int32 = int32(v2)
                _ = v14
            default:
                var v15 bool = (825 == 255)
                _ = v15
            }
            _ = len(v1)
            _ = fi10
            _ = v11
        }
        _ = fi9
    }
    v16 := GenS2{F0: 543, F1: 0}
    if true {
        switch v4 {
        case -342:
            v3 = v3
        default:
            v4 = (v4 >> ((4294967295) & 63))
        }
    } else {
        var v17 int = (v4 >> ((255) & 63))
        _ = len(v1)
        _ = v17
    }
    v18 := []int{}
    fmt.Println(v3, v4, v4)
    v4 += 2147483647
    var v19 int8 = 127
    v3 = v3
    _ = len(v1)
    if true {
        if (973 > 1) {
            v2 += (v2 % ((v2) | 1))
        } else {
            for fi20 := 0; fi20 < 18; fi20++ {
                if false {
                    var v21 int = 1
                    var v22 int16 = -641
                    _ = v21
                    _ = v22
                }
                _ = fi20
            }
            var v23 int32 = genFn0(128, 32767, (v2 % ((1) | 1)))
            var v24 uint8 = 255
            _ = v23
            _ = v24
        }
        v2 += v2
    }
    if true {
        fmt.Println(v2, v4)
    }
    if false {
        v4 -= v4
    } else {
        var v25 uint16 = 216
        _ = v25
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
