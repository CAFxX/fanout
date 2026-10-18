package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16 }
type GenS2 struct { F0 int; F1 bool; F2 uint16 }
type GenS3 struct { F0 uint32; F1 int32 }
func genFn0() bool {
    var v4 int32 = 564
    switch 2147483647 {
    case -2147483648:
        v4 = 0
    case 32768:
        v5 := GenS1{F0: 910, F1: 1000}
        _ = v5
    case 0:
        v6 := []int{-517, 128, 4294967295, -604}
        _ = v6
    default:
        v4 -= -1
    }
    _ = v4
    return (42 > 2)
}
func genFn1(p0 uint, p1 string, p2 int32) uint16 {
    v7 := []int{-1, 991}
    for fi8 := 0; fi8 < 13; fi8++ {
        var v9 float64 = 1.5
        fmt.Println(p0, p2)
        if (3.14159 < ((float32((0.0 / v9)) + 1e10) / -2.25)) {
            for fi10 := 0; fi10 < 8; fi10++ {
                p2 = p2
                var v11 int64 = 0
                _ = fi10
                _ = v11
            }
            var v12 float64 = 1.5
            if true {
                v12 += (0.0 * float64(float32((-2.25 * v9))))
            } else {
                for fi13 := 0; fi13 < 19; fi13++ {
                    p2 = p2
                    fmt.Println(v9, v12, p0)
                    v14 := GenS1{F0: 2147483648, F1: -1}
                    _ = fi13
                    _ = v14
                }
                if (18446744073709551615 > 18446744073709551615) {
                    v15 := GenS1{F0: 0, F1: 0}
                    v12 -= 1.5
                    _ = v15
                }
            }
            _ = v12
        } else {
            v7 = append(v7, (fi8 * fi8))
            var v16 bool = (874 == 42)
            _ = v16
        }
        _ = fi8
        _ = v9
    }
    fmt.Println(p2, p1, p2)
    var v17 uint32 = 166
    var v18 bool = (!((^(p2)) == p2))
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    _ = v18
    _ = v7
    return 255
}

func main() {
    v1 := make(map[int]int)
    v2 := []string{"\n\t"}
    v3 := []int{-1, 225, 2147483647, 42}
    var v19 int64 = 0
    for fi20 := 0; fi20 < 8; fi20++ {
        v2 = append(v2, "hello")
        v1[(fi20 >> ((255) & 63))] = (fi20 / ((fi20) | 1))
        _ = fi20
    }
    if (42 == 65535) {
        v21 := GenS2{F0: -207, F1: ((0.0 / (float32(float64(-2.25)) * (-2.25 + 0.0))) == float32(1e10)), F2: 2}
        v19 = v19
        if true {
            var v22 float32 = (0.1 / 0.1)
            _ = v22
        }
        _ = v21
    }
    if false {
        v23 := GenS3{F0: 196, F1: -1}
        v19 = (v19 % ((v19) | 1))
        v24 := GenS1{F0: 4294967295, F1: 1000}
        _ = v23
        _ = v24
    } else {
        _ = len(v3)
        var v25 uint64 = 42
        _ = v25
    }
    for fi26 := 0; fi26 < 1; fi26++ {
        v19 = v19
        v19 = 423
        var v27 int = (fi26 - 22)
        _ = fi26
        _ = v27
    }
    var v28 int16 = 1000
    fmt.Println(v19, v19)
    if (545 > 42) {
        fmt.Println(v28, v19)
    }
    v19 -= (v19 - (v19 + v19))
    var v29 uint = 1
    for fi30 := 0; fi30 < 13; fi30++ {
        var v31 int = (-(fi30))
        _ = fi30
        _ = v31
    }
    delete(v1, -2147483648)
    switch 1 {
    case -935:
        if ((v29 ^ v29) <= (v29 / ((335) | 1))) {
            for fi32 := 0; fi32 < 10; fi32++ {
                if true {
                    var v33 float32 = 0.0
                    delete(v1, fi32)
                    _ = v33
                }
                fmt.Println(v29, v19)
                var v34 float64 = (float64(0.0) / 1e10)
                _ = fi32
                _ = v34
            }
            fmt.Println(v29)
            v35 := GenS3{F0: 818, F1: 603}
            _ = v35
        }
    case 32767:
        v28 -= v28
    default:
        for fi36 := 0; fi36 < 3; fi36++ {
            var v37 float32 = (-2.25 * 0.1)
            _ = fi36
            _ = v37
        }
    }
    switch 2147483647 {
    case 0:
        v28 -= v28
    }
    v2 = append(v2, "a")
    v38 := GenS1{F0: 2, F1: v28}
    var v39 int8 = -128
    var v40 int16 = -1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v39:", v39)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v19
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v38
    _ = v39
    _ = v40
}
