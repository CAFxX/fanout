package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16; F2 bool }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int32; F1 uint16 }
type GenS4 struct { F0 uint64; F1 string }
func genFn0(p0 int64, p1 uint) uint64 {
    if (!(42 != -32768)) {
        v5 := GenS1{F0: 176, F1: 1000, F2: false}
        p1 = 0
        for fi6 := 0; fi6 < 18; fi6++ {
            v7 := []int{fi6, (fi6 | (^(fi6))), -882, (fi6 ^ 9223372036854775807)}
            v8 := GenS2{F0: 255}
            if false {
                v7 = append(v7, -35)
                var v9 uint = (p1 * p1)
                for fi10 := 0; fi10 < 14; fi10++ {
                    v11 := GenS1{F0: 42, F1: -854, F2: true}
                    p1 = 0
                    var v12 int16 = -721
                    _ = fi10
                    _ = v11
                    _ = v12
                }
                _ = v9
            } else {
                v13 := GenS2{F0: 127}
                if (!true) {
                    fmt.Println(p1, fi6)
                }
                var v14 float64 = 0.1
                _ = v13
                _ = v14
            }
            _ = fi6
            _ = v7
            _ = v8
        }
        _ = v5
    }
    if false {
        v15 := GenS3{F0: 0, F1: 639}
        var v16 int32 = 0
        _ = v15
        _ = v16
    }
    _ = p0
    _ = p1
    return 935
}
func genFn1() int {
    var v17 int = 42
    v18 := GenS4{F0: 0, F1: "\n\t"}
    _ = v17
    _ = v18
    return (v17 + (^(v17)))
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := []uint32{4294967295, 0, 1, 0}
    v3 := []uint32{2, 42, 2}
    var v4 int32 = -2147483648
    if (-2147483648 <= genFn1()) {
        v4 -= v4
        var v19 int8 = 1
        _ = v19
    }
    v4 = -38
    for fi20 := 0; fi20 < 12; fi20++ {
        for fi21 := 0; fi21 < 9; fi21++ {
            if (2 != 328) {
                var v22 uint8 = 42
                _ = v22
            } else {
                var v23 int16 = -1
                _ = v23
            }
            _ = fi21
        }
        _ = fi20
    }
    switch 42 {
    case 4294967295:
        fmt.Println(v4)
    }
    fmt.Println(v4)
    v4 += (v4 & v4)
    var v24 uint16 = 603
    v25 := GenS4{F0: 1, F1: "hello"}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
}
