package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16 }
type GenS2 struct { F0 uint16; F1 int }
type GenS3 struct { F0 bool; F1 int16; F2 uint }
func genFn0(p0 uint16, p1 uint32) uint {
    p0 -= p0
    v4 := []int{791, -9223372036854775808, -9223372036854775808, 127}
    switch -9223372036854775808 {
    case -1:
        fmt.Println(p0, p0)
    default:
        v5 := GenS1{F0: 901, F1: -32768}
        _ = v5
    }
    _ = p0
    _ = p1
    _ = v4
    return 18446744073709551615
}
func genFn1(p0 int8, p1 uint32, p2 int16) uint16 {
    v6 := []int{-2147483648}
    switch 32767 {
    case -825:
        p0 += 127
    case 573:
        p1 -= 411
    default:
        for fi7 := 0; fi7 < 4; fi7++ {
            fmt.Println(p0, p1, p1)
            _ = fi7
        }
    }
    if (!true) {
        var v8 uint16 = 32767
        var v9 uint8 = 138
        p0 += (p0 ^ int8(p1))
        _ = v8
        _ = v9
    }
    var v10 uint = 33
    v6 = append(v6, 931)
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v6
    return 2
}

func main() {
    var v1 uint8 = 2
    v2 := make(map[string]uint32)
    v3 := make(map[int]string)
    if false {
        v1 += 159
        v11 := GenS2{F0: 32767, F1: 403}
        _ = v11
    } else {
        var v12 bool = (703 == 42)
        _ = v12
    }
    if (1.5 < (float64(1e10) - 1e10)) {
        v13 := []int{136, -280, 384, -469}
        _ = v13
    }
    _ = len(v3)
    v1 = v1
    for fi14 := 0; fi14 < 14; fi14++ {
        v1 += 255
        _ = fi14
    }
    v1 += v1
    if ((!false) || ((true && true) && ((!(376 > 903)) || ((612 == 32767) || true)))) {
        v1 += v1
        v15 := GenS3{F0: false, F1: 650, F2: 383}
        _ = v15
    } else {
        delete(v3, 32767)
        var v16 int32 = -2147483648
        _ = v16
    }
    v1 = 42
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v3
}
