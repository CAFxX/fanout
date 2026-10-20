package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64 }
type GenS2 struct { F0 uint32; F1 int16; F2 uint }
type GenS3 struct { F0 int32 }
func genFn0(p0 uint16, p1 int16) int16 {
    fmt.Println(p1)
    if (2147483647 != -216) {
        p0 += (p0 >> ((42) & 63))
        for fi6 := 0; fi6 < 9; fi6++ {
            var v7 int = fi6
            _ = fi6
            _ = v7
        }
        p0 = (^(p0))
    }
    v8 := GenS1{F0: 128, F1: 2}
    p0 -= (p0 - (p0 | p0))
    p1 = p1
    _ = p0
    _ = p1
    _ = v8
    return p1
}
func genFn1(p0 int64, p1 int32, p2 int32) uint16 {
    var v9 int32 = -1
    v10 := []int{-1, 0, int(p0)}
    v9 = p2
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return 42
}

func main() {
    v1 := []uint32{}
    var v2 int32 = 1
    var v3 uint32 = 42
    var v4 int32 = v2
    var v5 string = ("\n\t" + "\n\t")
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v4 += -2147483648
    v2 -= 42
    v4 = (v4 | v4)
    if len(v1) > 0 {
        v1[0] = v3
    }
    for fi11 := 0; fi11 < 9; fi11++ {
        if true {
            var v12 string = v5
            for fi13 := 0; fi13 < 5; fi13++ {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = fi13
            }
            _ = v12
        }
        _ = fi11
    }
    v1 = append(v1, (v3 >> ((805) & 63)))
    switch ((v5 + v5) + ("\n\t" + "hello")) {
    case "":
        v14 := GenS2{F0: v3, F1: 0, F2: 2}
        _ = v14
    case "hello":
        var v15 int8 = -26
        _ = v15
    case "\n\t":
        fmt.Println(v5, v5)
    default:
        if (100 <= 127) {
            for fi16 := 0; fi16 < 11; fi16++ {
                v17 := GenS3{F0: (v4 % (((v2 / ((v2) | 1))) | 1))}
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = fi16
                _ = v17
            }
        } else {
            v2 = v4
            _ = len(v1)
            var v18 int8 = 100
            _ = v18
        }
    }
    var v19 float64 = (((float64(float32((0.1 * (1.5 / (float64(0.0) - 1e10))))) - 1e10) - (0.0 * 0.0)) * float64(1.5))
    var v20 uint64 = 1
    var v21 string = (v5 + v5)
    if (!(1 >= 42)) {
        v2 -= v2
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
