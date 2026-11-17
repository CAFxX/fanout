package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32; F2 string }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint64 }

func main() {
    v1 := make(map[string]string)
    v2 := make(map[int]uint32)
    var v3 float64 = (0.0 + (float64(((1.5 * 0.0) - 1e10)) * (3.14159 / float64(float32(float64((-2.25 / 1.0)))))))
    var v4 int64 = -1
    v5 := []int{0, 127, -2147483648, -1}
    fmt.Println(v4, v3)
    v4 = -881
    v3 -= v3
    for fi6 := 0; fi6 < 10; fi6++ {
        fmt.Println(v4, fi6)
        _ = fi6
    }
    v7 := GenS1{F0: -311, F1: -1, F2: ""}
    var v8 bool = (((v3 == float64((1e10 + 0.0))) || true) || (2 == 0))
    if v8 {
        v3 -= v3
        for fi9 := 0; fi9 < 9; fi9++ {
            v10 := GenS2{F0: 4294967295}
            var v11 string = "hello"
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            _ = fi9
            _ = v10
            _ = v11
        }
        var v12 int16 = -32768
        _ = v12
    } else {
        switch 2147483647 {
        case 32767:
            var v13 float32 = 1.5
            _ = v13
        default:
            for fi14 := 0; fi14 < 18; fi14++ {
                fmt.Println(fi14, v8)
                _ = fi14
            }
        }
        v3 = v3
    }
    v4 = (v4 << ((973) & 63))
    for fi15 := 0; fi15 < 6; fi15++ {
        fmt.Println(fi15)
        _ = fi15
    }
    v3 = v3
    var v16 int = 0
    if (!v8) {
        v17 := GenS3{F0: 18446744073709551615}
        _ = v17
    } else {
        var v18 uint = 486
        v8 = (v8 || ((((v3 * 1.5) - (v3 - v3)) - (v3 / v3)) != v3))
        var v19 float64 = 0.1
        _ = v18
        _ = v19
    }
    switch 768 {
    case 42:
        fmt.Println(v8)
    }
    for fi20 := 0; fi20 < 14; fi20++ {
        v5 = append(v5, -9223372036854775808)
        for fi21 := 0; fi21 < 11; fi21++ {
            if false {
                v22 := GenS3{F0: 4294967295}
                var v23 uint64 = 1
                v3 -= ((-2.25 / 0.1) * (v3 / v3))
                _ = v22
                _ = v23
            }
            _ = fi21
        }
        _ = fi20
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
