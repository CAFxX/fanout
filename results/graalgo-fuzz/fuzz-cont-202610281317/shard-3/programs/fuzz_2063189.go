package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string; F2 int }
type GenS2 struct { F0 uint16 }
func genFn0(p0 int64) uint64 {
    switch "\n\t" {
    case "foo bar":
        v5 := GenS1{F0: -2147483648, F1: ("foo bar" + "hello"), F2: 127}
        _ = v5
    }
    var v6 uint8 = 84
    v7 := GenS1{F0: 42, F1: (("\n\t" + "") + "\n\t"), F2: 127}
    _ = p0
    _ = v6
    _ = v7
    return 1
}
func genFn1(p0 uint16, p1 int) int16 {
    if (-2147483648 < -145) {
        for fi8 := 0; fi8 < 3; fi8++ {
            v9 := []int{(fi8 ^ p1), (p1 << ((4294967295) & 63))}
            _ = len(v9)
            _ = fi8
            _ = v9
        }
        var v10 float32 = (1.5 + (0.1 * ((float32(3.14159) / 3.14159) - 0.0)))
        _ = v10
    }
    for fi11 := 0; fi11 < 19; fi11++ {
        if true {
            p1 = (p1 % ((fi11) | 1))
            for fi12 := 0; fi12 < 4; fi12++ {
                if true {
                    fmt.Println(fi12, fi12)
                    v13 := GenS1{F0: 928, F1: "", F2: (^(p1))}
                    if (false || false) {
                        fmt.Println(p0, p0)
                    }
                    _ = v13
                } else {
                    if (!(1.5 > ((1.5 + 1.5) * 0.1))) {
                        var v14 uint64 = uint64(p0)
                        v15 := GenS1{F0: -1, F1: "a", F2: (-(p1))}
                        _ = v14
                        _ = v15
                    } else {
                        v16 := GenS1{F0: 0, F1: (("foo bar" + "hello") + "hello"), F2: fi11}
                        fmt.Println(p1, fi11, fi12)
                        var v17 int = (p1 * (fi12 ^ -284))
                        _ = v16
                        _ = v17
                    }
                    var v18 int = 32768
                    p0 = 1
                    _ = v18
                }
                v19 := []int{(-(fi11))}
                _ = fi12
                _ = v19
            }
        }
        _ = fi11
    }
    var v20 int64 = 1
    switch (p1 & 9223372036854775807) {
    case -9223372036854775808:
        v21 := GenS2{F0: p0}
        _ = v21
    default:
        var v22 int32 = 2147483647
        _ = v22
    }
    _ = p0
    _ = p1
    _ = v20
    return 1
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[uint32]string)
    var v3 int8 = 122
    var v4 uint = 65535
    fmt.Println(v4, v3)
    var v23 uint32 = 2
    v4 += v4
    fmt.Println(v3, v23)
    var v24 string = "\n\t"
    v4 = v4
    v23 = (v23 << ((2) & 63))
    v4 += 18446744073709551615
    v4 = (v4 << ((v4) & 63))
    v4 = v4
    if (138 < 9223372036854775807) {
        switch ("a" + v24) {
        case "":
            v25 := []int{-135, -552}
            _ = v25
        case "hello":
            v26 := GenS1{F0: 42, F1: "", F2: -2147483648}
            _ = v26
        default:
            var v27 float64 = ((float64(1.5) / (1e10 - (1.5 + float64(float32(1e10))))) / 0.1)
            _ = v27
        }
    }
    for fi28 := 0; fi28 < 11; fi28++ {
        v23 -= v23
        _ = fi28
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
}
