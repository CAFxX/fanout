package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint64; F1 uint16 }
func genFn0(p0 int32, p1 uint16, p2 uint32) int {
    fmt.Println(p2, p0, p1)
    p2 -= p2
    v4 := GenS1{F0: p2}
    switch 0 {
    case 9223372036854775807:
        v5 := []int{128}
        _ = v5
    case 485:
        v6 := []int{32768, -9223372036854775808}
        _ = v6
    case 32767:
        v7 := []int{-964, -313, -275, 684}
        _ = v7
    }
    p0 -= (p0 & (p0 - p0))
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 2147483647
}

func main() {
    v1 := []int{577, 42, 2147483647}
    var v2 uint64 = 1
    v3 := make(map[uint32]uint32)
    v2 = v2
    switch 32768 {
    case 32767:
        for fi8 := 0; fi8 < 2; fi8++ {
            v9 := GenS2{F0: (v2 * v2), F1: 65535}
            _ = fi8
            _ = v9
        }
    case -9223372036854775808:
        switch 9223372036854775807 {
        case 1:
            v3[2] = 0
        case 128:
            delete(v3, 42)
        case 127:
            if (true || true) {
                var v10 float64 = ((3.14159 / 1.0) + (0.0 / 3.14159))
                var v11 uint = 336
                for fi12 := 0; fi12 < 17; fi12++ {
                    if ((uint16(v11) >= 255) || true) {
                        switch -9223372036854775808 {
                        case 32768:
                            v13 := GenS1{F0: 4294967295}
                            _ = v13
                        case 2147483647:
                            v14 := GenS1{F0: 42}
                            _ = v14
                        case -2147483648:
                            v11 += 255
                        default:
                            v2 -= (v2 >> ((v11) & 63))
                        }
                        v2 = v2
                        var v15 uint = 4294967295
                        _ = v15
                    }
                    _ = fi12
                }
                _ = v10
                _ = v11
            } else {
                fmt.Println(v2)
            }
        default:
            v16 := v3[4294967295]
            _ = v16
        }
    default:
        if ((0 == -128) || true) {
            v2 -= (v2 - v2)
        }
    }
    v17 := GenS2{F0: (v2 % ((285) | 1)), F1: 2}
    fmt.Println(v2)
    fmt.Println(v2)
    fmt.Println(v2)
    switch "a" {
    case "hello":
        v2 -= v2
    case "foo bar":
        switch 32768 {
        case 1:
            var v18 int = 2147483647
            _ = v18
        case 42:
            var v19 float64 = float64(0.1)
            _ = v19
        }
    case "a":
        var v20 int32 = -2147483648
        _ = v20
    }
    switch "a" {
    case "foo bar":
        v2 += (v2 << ((18446744073709551615) & 63))
    }
    v2 = 4294967295
    var v21 int16 = 85
    for fi22 := 0; fi22 < 9; fi22++ {
        _ = len(v3)
        v2 -= (v2 | v2)
        _ = fi22
    }
    var v23 float32 = 1e10
    switch 0 {
    case -123:
        switch "\n\t" {
        case "a":
            var v24 int16 = (v21 % ((v21) | 1))
            _ = v24
        case "hello":
            _ = len(v1)
        default:
            fmt.Println(v21)
        }
    case -799:
        var v25 int32 = -1
        _ = v25
    }
    switch 1 {
    case 0:
        var v26 uint64 = v2
        _ = v26
    case 329:
        fmt.Println(v2, v2, v21)
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v23 -= 1.5
    for fi27 := 0; fi27 < 10; fi27++ {
        fmt.Println(v23)
        _ = fi27
    }
    if (false || true) {
        fmt.Println(v2, v23, v21)
        var v28 string = ""
        if false {
            v23 = v23
        } else {
            fmt.Println(v28, v28)
        }
        _ = v28
    } else {
        _ = len(v1)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v17
    _ = v2
    _ = v21
    _ = v23
    _ = v3
}
