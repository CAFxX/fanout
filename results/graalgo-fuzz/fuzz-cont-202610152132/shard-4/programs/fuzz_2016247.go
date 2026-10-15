package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint }
type GenS2 struct { F0 int16; F1 uint16 }
func genFn0(p0 int8, p1 int64) uint {
    fmt.Println(p0, p1)
    switch "a" {
    case "":
        if (932 < 127) {
            for fi7 := 0; fi7 < 13; fi7++ {
                fmt.Println(p1, fi7)
                v8 := []int{fi7, (fi7 & fi7)}
                if len(v8) >= 2 {
                    v8 = v8[1:len(v8)]
                }
                _ = fi7
                _ = v8
            }
            switch "\n\t" {
            case "":
                for fi9 := 0; fi9 < 10; fi9++ {
                    var v10 int32 = -2147483648
                    _ = fi9
                    _ = v10
                }
            case "\n\t":
                p0 -= -1
            case "foo bar":
                v11 := GenS1{F0: 312, F1: 255}
                _ = v11
            }
        }
    case "foo bar":
        v12 := GenS1{F0: 2, F1: 801}
        _ = v12
    case "hello":
        p1 = (p1 >> ((18446744073709551615) & 63))
    }
    _ = p0
    _ = p1
    return 109
}

func main() {
    v1 := []int{-1, 42}
    var v2 uint32 = 65535
    v3 := []int{0}
    v4 := []int{251, 371, 32768, 0}
    v5 := make(map[int]uint32)
    v6 := make(map[string]string)
    if false {
        for fi13 := 0; fi13 < 15; fi13++ {
            switch fi13 {
            case -1:
                _ = len(v6)
            case -2147483648:
                for fi14 := 0; fi14 < 12; fi14++ {
                    v15 := GenS2{F0: -747, F1: 255}
                    _ = fi14
                    _ = v15
                }
            default:
                v2 += (v2 << ((157) & 63))
            }
            for fi16 := 0; fi16 < 15; fi16++ {
                var v17 uint64 = 0
                _ = fi16
                _ = v17
            }
            if (2147483648 <= v2) {
                v18 := GenS2{F0: 0, F1: 65535}
                var v19 bool = (1.5 > (-2.25 + (float32(float64(0.1)) - 1e10)))
                _ = v18
                _ = v19
            }
            _ = fi13
        }
        v2 += (v2 >> ((42) & 63))
    } else {
        var v20 float64 = float64(float32((1e10 / -2.25)))
        if len(v1) > 0 {
            v1[0] = int(v2)
        }
        fmt.Println(v20, v2)
        _ = v20
    }
    v21 := GenS2{F0: -846, F1: 1}
    v2 += v2
    var v22 uint8 = 1
    var v23 int8 = 95
    v24 := GenS1{F0: 42, F1: uint(v22)}
    var v25 int = -2147483648
    var v26 float32 = (float32(0.0) / (1.5 / 1.5))
    if len(v1) > 0 {
        v1[0] = 4294967295
    }
    fmt.Println(v26, v22)
    for fi27 := 0; fi27 < 5; fi27++ {
        v22 = v22
        _ = fi27
    }
    fmt.Println(v23)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
