package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32 }
type GenS2 struct { F0 uint8; F1 uint8; F2 int8 }
type GenS3 struct { F0 uint }
func genFn0(p0 bool) int16 {
    v5 := GenS1{F0: -2147483648, F1: 0}
    v6 := []int{2147483647, 32767, 128}
    v7 := GenS2{F0: 36, F1: 0, F2: 1}
    var v8 uint32 = 980
    _ = len(v6)
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return 32767
}
func genFn1(p0 uint, p1 int8, p2 int) int32 {
    var v9 int = p2
    var v10 float64 = 3.14159
    v11 := GenS2{F0: 48, F1: uint8(p0), F2: (p1 >> (((p0 * 255)) & 63))}
    var v12 string = "a"
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v12
    _ = v9
    return 1
}

func main() {
    var v1 uint64 = 2
    v2 := make(map[string]bool)
    v3 := []uint32{269, 4294967295, 372, 42}
    v4 := make(map[int]uint32)
    v1 -= v1
    switch 961 {
    case -1:
        var v13 float64 = -2.25
        _ = v13
    case 42:
        if false {
            v1 += 1
            v1 = (v1 - 622)
            v1 += v1
        } else {
            _ = len(v3)
            var v14 uint32 = 65535
            _ = v14
        }
    }
    for fi15 := 0; fi15 < 6; fi15++ {
        var v16 int16 = 42
        var v17 uint16 = 255
        if false {
            switch -626 {
            case 32768:
                v16 = v16
            case 128:
                v16 = genFn0(true)
            case -2147483648:
                v16 += v16
            }
            for fi18 := 0; fi18 < 18; fi18++ {
                var v19 int32 = -803
                v20 := GenS3{F0: 2}
                _ = fi18
                _ = v19
                _ = v20
            }
            v3 = append(v3, 42)
        }
        _ = fi15
        _ = v16
        _ = v17
    }
    var v21 int16 = 1000
    v1 += 4294967295
    for fi22 := 0; fi22 < 17; fi22++ {
        v23 := GenS3{F0: 1}
        _ = fi22
        _ = v23
    }
    v1 = (v1 + v1)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v1 = v1
    fmt.Println(v1, v1)
    v1 -= 18446744073709551615
    fmt.Println(v1, v21)
    _ = len(v3)
    var v24 uint8 = 238
    var v25 int8 = -128
    for fi26 := 0; fi26 < 1; fi26++ {
        if (1e10 < 0.1) {
            var v27 uint64 = (v1 ^ v1)
            fmt.Println(v25)
            _ = v27
        }
        v21 = (v21 >> ((255) & 63))
        _ = fi26
    }
    var v28 float64 = -2.25
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v21
    _ = v24
    _ = v25
    _ = v28
    _ = v3
    _ = v4
}
