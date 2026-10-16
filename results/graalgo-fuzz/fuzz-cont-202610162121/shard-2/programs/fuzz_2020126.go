package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int }
type GenS2 struct { F0 string; F1 int64; F2 uint8 }
type GenS3 struct { F0 int32 }
func genFn0(p0 uint32) uint8 {
    var v4 uint16 = 42
    v4 = 255
    v4 += v4
    v5 := GenS1{F0: 0, F1: -747}
    _ = p0
    _ = v4
    _ = v5
    return 2
}

func main() {
    var v1 int16 = 498
    v2 := make(map[string]uint32)
    var v3 int8 = 1
    _ = len(v2)
    var v6 int8 = v3
    v7 := GenS1{F0: (v1 + 42), F1: -1}
    v3 += -45
    for fi8 := 0; fi8 < 5; fi8++ {
        for fi9 := 0; fi9 < 6; fi9++ {
            fmt.Println(v3)
            _ = fi9
        }
        v10 := []int{(fi8 << ((0) & 63))}
        v11 := GenS2{F0: ("a" + "\n\t"), F1: 42, F2: 141}
        _ = fi8
        _ = v10
        _ = v11
    }
    v12 := GenS2{F0: "\n\t", F1: 9223372036854775807, F2: 255}
    var v13 float32 = (0.0 - 3.14159)
    v13 = v13
    v3 -= v6
    for fi14 := 0; fi14 < 15; fi14++ {
        if false {
            fmt.Println(v3, fi14)
            var v15 uint16 = 1
            switch fi14 {
            case 127:
                v15 = v15
            case -2147483648:
                if (42 < 820) {
                    fmt.Println(v1)
                    fmt.Println(v6)
                    var v16 uint32 = 93
                    _ = v16
                }
            default:
                fmt.Println(fi14, v15)
            }
            _ = v15
        } else {
            var v17 uint64 = 2
            var v18 int16 = (v1 >> ((1) & 63))
            v19 := GenS3{F0: 42}
            _ = v17
            _ = v18
            _ = v19
        }
        v20 := []int{}
        _ = fi14
        _ = v20
    }
    var v21 uint8 = 1
    v22 := []int{128}
    _ = len(v2)
    v23 := GenS3{F0: -729}
    var v24 int64 = 0
    if false {
        _ = len(v22)
        for fi25 := 0; fi25 < 16; fi25++ {
            if false {
                if len(v22) > 0 {
                    v22[0] = (-(fi25))
                }
            }
            _ = fi25
        }
    } else {
        for fi26 := 0; fi26 < 8; fi26++ {
            fmt.Println(v24, v3)
            var v27 uint32 = 0
            _ = fi26
            _ = v27
        }
    }
    fmt.Println(v1, v13, v13)
    var v28 int = 128
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v24:", v24)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v28
    _ = v3
    _ = v6
    _ = v7
}
