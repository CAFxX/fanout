package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32 }
type GenS2 struct { F0 int16; F1 int8; F2 uint64 }
func genFn0(p0 int8, p1 int16) int {
    if true {
        for fi7 := 0; fi7 < 2; fi7++ {
            for fi8 := 0; fi8 < 5; fi8++ {
                v9 := GenS1{F0: fi7, F1: 42}
                var v10 bool = (!(0 < 2147483647))
                _ = fi8
                _ = v10
                _ = v9
            }
            fmt.Println(p0, fi7, fi7)
            _ = fi7
        }
        if (2 == 0) {
            if ((!true) && ((!(-9223372036854775808 != 281)) || (65535 < 2147483648))) {
                var v11 string = "hello"
                _ = v11
            }
            switch -1 {
            case 42:
                v12 := GenS1{F0: -1, F1: -2147483648}
                _ = v12
            case -2147483648:
                var v13 uint64 = 621
                _ = v13
            case -486:
                p0 += (p0 / ((127) | 1))
            default:
                var v14 uint = 255
                _ = v14
            }
        } else {
            v15 := []int{-1, -266, -1}
            fmt.Println(p1)
            p0 = (p0 << ((654) & 63))
            _ = v15
        }
    } else {
        p1 += p1
        for fi16 := 0; fi16 < 2; fi16++ {
            var v17 bool = true
            var v18 uint64 = 574
            _ = fi16
            _ = v17
            _ = v18
        }
        v19 := []int{128, 1, 32767}
        _ = v19
    }
    p0 -= (p0 - p0)
    var v20 int32 = -418
    switch "a" {
    case "hello":
        v21 := []int{2147483647}
        _ = v21
    case "foo bar":
        v22 := []int{42, int(v20), int(p1), -9223372036854775808}
        _ = v22
    default:
        var v23 int8 = -128
        _ = v23
    }
    _ = p0
    _ = p1
    _ = v20
    return 84
}
func genFn1(p0 int) uint32 {
    var v24 int16 = 1000
    switch p0 {
    case 42:
        for fi25 := 0; fi25 < 1; fi25++ {
            p0 -= 369
            for fi26 := 0; fi26 < 16; fi26++ {
                p0 = (fi25 >> ((87) & 63))
                _ = fi26
            }
            _ = fi25
        }
    case -9223372036854775808:
        p0 -= (p0 >> ((61) & 63))
    }
    _ = p0
    _ = v24
    return 374
}

func main() {
    v1 := []uint32{2147483648}
    var v2 float64 = 0.0
    v3 := make(map[uint32]bool)
    var v4 float32 = (float32(((v2 + ((3.14159 - float64(((-2.25 - 0.1) / (3.14159 - (((3.14159 + (1.5 * -2.25)) * 3.14159) - 0.1))))) + (float64(float32((v2 - v2))) - v2))) + (v2 - ((float64(0.0) / v2) - 0.0)))) - (1e10 / 1.0))
    var v5 int16 = 32767
    v6 := make(map[string]int)
    v4 += 0.0
    if true {
        fmt.Println(v5)
    }
    v5 = v5
    if false {
        switch ("\n\t" + ("\n\t" + "hello")) {
        case "":
            fmt.Println(v5)
        case "foo bar":
            v2 -= (v2 + float64(0.1))
        default:
            fmt.Println(v4, v5, v4)
        }
    }
    v27 := GenS2{F0: v5, F1: -91, F2: 421}
    fmt.Println(v5)
    var v28 float64 = (((0.0 / v2) - v2) + v2)
    v4 -= ((((v4 + ((((v4 * (0.1 - (v4 / v4))) - 3.14159) + float32(0.0)) - (float32(v2) * v4))) / v4) + -2.25) - -2.25)
    if len(v1) > 0 {
        v1[0] = 42
    }
    v5 = v5
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
