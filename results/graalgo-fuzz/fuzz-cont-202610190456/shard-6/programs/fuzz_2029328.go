package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int }
type GenS2 struct { F0 int8; F1 uint32 }
func genFn0(p0 uint8, p1 uint16, p2 int) uint8 {
    v6 := []int{42, 1, p2, 1}
    if true {
        fmt.Println(p2)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 0
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 float64 = -2.25
    v3 := make(map[uint32]uint32)
    var v4 float64 = (v2 - -2.25)
    var v5 uint32 = 65535
    v5 = (^(v5))
    v7 := []int{-1, -1}
    _ = len(v7)
    v8 := GenS1{F0: 249, F1: 183}
    v1[v5] = (v5 % ((v5) | 1))
    fmt.Println(v5)
    for fi9 := 0; fi9 < 12; fi9++ {
        v4 -= (v4 * (float64(0.0) * float64(1.5)))
        if (1 <= 0) {
            for fi10 := 0; fi10 < 20; fi10++ {
                v2 += -2.25
                var v11 int16 = -901
                _ = fi10
                _ = v11
            }
        }
        for fi12 := 0; fi12 < 20; fi12++ {
            v7 = append(v7, fi12)
            fmt.Println(fi9)
            var v13 int16 = 32767
            _ = fi12
            _ = v13
        }
        _ = fi9
    }
    v2 += (1e10 - (v4 * v2))
    if false {
        var v14 uint8 = 127
        if ((!false) || (1e10 <= float32(v2))) {
            if ((-37 < 127) && true) {
                var v15 int = 32767
                _ = v15
            } else {
                for fi16 := 0; fi16 < 10; fi16++ {
                    switch fi16 {
                    case 273:
                        v2 -= (v2 / ((v2 * (v2 * float64((((1.5 * 1e10) * (((1.5 + 1e10) * float32((float64((((1.5 - float32((3.14159 / v2))) / float32((1.5 + v2))) - -2.25)) / v4))) * (((0.1 + (float32(float64(1e10)) / (float32((1e10 * float64(-2.25))) * 1.5))) / (((1e10 + float32((v2 + (1.5 - v4)))) * ((((float32(v2) * 3.14159) * 1.5) - 3.14159) / 3.14159)) / -2.25)) - 0.0))) / 1.0)))) - 1e10))
                    default:
                        v17 := GenS2{F0: 42, F1: v5}
                        _ = v17
                    }
                    _ = fi16
                }
                v5 = (v5 << ((255) & 63))
                var v18 int8 = 100
                _ = v18
            }
            fmt.Println(v5)
        } else {
            v14 = (v14 ^ 0)
            var v19 int = 32767
            _ = v19
        }
        _ = v14
    } else {
        fmt.Println(v4)
        if len(v7) > 0 {
            v7[0] = 2147483647
        }
        switch 1 {
        case -20:
            v5 -= (v5 | 42)
        case 127:
            var v20 int16 = 42
            _ = v20
        case 128:
            var v21 int32 = -193
            _ = v21
        }
    }
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    for fi22 := 0; fi22 < 3; fi22++ {
        fmt.Println(v5)
        v4 -= v2
        _ = fi22
    }
    v23 := v1[195]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
