package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64; F2 int32 }
type GenS2 struct { F0 uint; F1 uint; F2 uint32 }
func genFn0(p0 int, p1 int, p2 int32) int64 {
    fmt.Println(p1)
    for fi4 := 0; fi4 < 6; fi4++ {
        fmt.Println(p0, p1)
        _ = fi4
    }
    v5 := []int{}
    if (!false) {
        v5 = append(v5, (p1 >> ((4294967295) & 63)))
        var v6 float64 = 0.1
        fmt.Println(v6, p1)
        _ = v6
    } else {
        var v7 float64 = ((0.0 - ((((0.0 - (0.0 * 3.14159)) * -2.25) / (1.5 + -2.25)) * (-2.25 + 0.1))) * float64(0.0))
        v7 = v7
        _ = v7
    }
    for fi8 := 0; fi8 < 1; fi8++ {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        var v9 string = ("a" + "foo bar")
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return -472
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int8 = 41
    var v3 int32 = 0
    v10 := []int{-1, 33, 127, -2147483648}
    for fi11 := 0; fi11 < 16; fi11++ {
        v1[(^(fi11))] = 2147483647
        _ = fi11
    }
    var v12 uint8 = 91
    v13 := GenS1{F0: 1000, F1: 2, F2: -2147483648}
    v2 = (v2 | int8(v12))
    v14 := GenS1{F0: 1, F1: 1, F2: int32(v12)}
    v15 := GenS2{F0: 1, F1: 65535, F2: 1}
    delete(v1, -277)
    var v16 uint16 = 2
    v3 = v3
    switch 677 {
    case 128:
        fmt.Println(v12, v2, v12)
    }
    v10 = append(v10, -1)
    switch 32768 {
    case 32767:
        delete(v1, 128)
    case -537:
        if (!true) {
            fmt.Println(v12)
            var v17 int64 = 0
            v18 := GenS1{F0: 1000, F1: 844, F2: int32(v17)}
            _ = v17
            _ = v18
        }
    case -9223372036854775808:
        for fi19 := 0; fi19 < 19; fi19++ {
            if len(v10) > 0 {
                v10[0] = (fi19 * fi19)
            }
            var v20 int64 = -148
            _ = fi19
            _ = v20
        }
    default:
        v2 += (-(v2))
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
}
