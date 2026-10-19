package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int64; F1 int }
type GenS3 struct { F0 uint8 }
func genFn0(p0 string, p1 uint64) uint16 {
    v6 := []int{32767, 899, 2147483647, -374}
    if true {
        for fi7 := 0; fi7 < 6; fi7++ {
            var v8 float32 = 1.5
            fmt.Println(p0)
            for fi9 := 0; fi9 < 7; fi9++ {
                fmt.Println(v8, fi9, p0)
                var v10 int32 = -2147483648
                v10 += (v10 << ((4294967295) & 63))
                p0 = p0
                _ = fi9
                _ = v10
            }
            _ = fi7
            _ = v8
        }
        var v11 uint8 = 42
        v6 = append(v6, 4294967295)
        _ = v11
    } else {
        p0 = p0
    }
    if false {
        _ = len(v6)
        fmt.Println(p1, p0)
        var v12 float32 = ((((0.1 / 1.5) + 1.5) - 3.14159) / 1.0)
        _ = v12
    } else {
        if len(v6) > 0 {
            v6[0] = -9223372036854775808
        }
        switch 32768 {
        case -197:
            fmt.Println(p1)
        default:
            var v13 int16 = 32767
            _ = v13
        }
    }
    _ = p0
    _ = p1
    _ = v6
    return 32767
}
func genFn1(p0 int64, p1 string, p2 int64) uint {
    fmt.Println(p0)
    for fi14 := 0; fi14 < 1; fi14++ {
        if false {
            for fi15 := 0; fi15 < 14; fi15++ {
                v16 := GenS1{F0: (fi15 / (((-(fi15))) | 1))}
                var v17 float64 = 0.0
                _ = fi15
                _ = v16
                _ = v17
            }
        }
        p1 = p1
        v18 := GenS2{F0: p0, F1: 127}
        _ = fi14
        _ = v18
    }
    v19 := GenS2{F0: (p2 & (p0 / ((p0) | 1))), F1: -722}
    _ = p0
    _ = p1
    _ = p2
    _ = v19
    return 4294967295
}

func main() {
    v1 := []int{}
    v2 := make(map[uint32]string)
    var v3 int16 = 32767
    var v4 bool = false
    var v5 int32 = 56
    fmt.Println(v4)
    var v20 string = "a"
    v21 := GenS3{F0: 127}
    _ = len(v2)
    for fi22 := 0; fi22 < 16; fi22++ {
        v20 = v20
        delete(v2, 0)
        if false {
            v23 := GenS1{F0: 127}
            for fi24 := 0; fi24 < 12; fi24++ {
                v25 := v2[4294967295]
                var v26 uint64 = 0
                _ = fi24
                _ = v25
                _ = v26
            }
            _ = v23
        }
        _ = fi22
    }
    v3 -= 1000
    v27 := GenS2{F0: 9223372036854775807, F1: 1}
    v28 := GenS2{F0: 0, F1: -2147483648}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
