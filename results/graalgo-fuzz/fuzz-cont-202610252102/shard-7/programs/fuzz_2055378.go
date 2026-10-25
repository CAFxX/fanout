package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint }
type GenS2 struct { F0 uint64; F1 int8; F2 uint64 }
type GenS3 struct { F0 uint8; F1 uint }

func main() {
    v1 := []int{-92, 42, -2147483648, -9223372036854775808}
    v2 := []int{}
    v3 := []uint32{2}
    v4 := []string{"foo bar", "foo bar", "\n\t"}
    v5 := GenS1{F0: -363, F1: 932}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v6 float64 = (0.1 * 0.0)
    v6 += -2.25
    _ = len(v1)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if (false && (145 > 164)) {
        switch 9223372036854775807 {
        case -863:
            fmt.Println(v6)
        case 346:
            var v7 int64 = -186
            _ = v7
        }
    } else {
        v6 = v6
        var v8 uint = 422
        _ = v8
    }
    var v9 int16 = 0
    v10 := GenS2{F0: 0, F1: 127, F2: 825}
    switch 0 {
    case 42:
        v6 = (v6 * v6)
    case -609:
        fmt.Println(v6, v9)
    case 32768:
        var v11 bool = (18446744073709551615 < 42)
        _ = v11
    }
    fmt.Println(v6, v6)
    if true {
        if ((127 < 2) || (65535 == 966)) {
            var v12 float32 = -2.25
            _ = v12
        }
        v6 += ((v6 - v6) + v6)
        var v13 uint64 = 4294967295
        _ = v13
    } else {
        fmt.Println(v6, v9)
        v14 := GenS3{F0: 111, F1: 193}
        _ = v14
    }
    for fi15 := 0; fi15 < 15; fi15++ {
        for fi16 := 0; fi16 < 18; fi16++ {
            v6 = v6
            v6 = (float64(0.0) / 1e10)
            _ = fi16
        }
        switch (fi15 + (-(fi15))) {
        case -83:
            v17 := GenS3{F0: 2, F1: 1}
            _ = v17
        default:
            if (2147483647 != 1) {
                if len(v1) > 0 {
                    v1[0] = (fi15 & (fi15 << ((4294967295) & 63)))
                }
                if true {
                    var v18 int64 = 1
                    v18 -= v18
                    _ = v18
                }
            } else {
                _ = len(v3)
                for fi19 := 0; fi19 < 13; fi19++ {
                    v20 := GenS1{F0: (fi15 + (fi15 >> ((42) & 63))), F1: 940}
                    switch 2147483647 {
                    case 127:
                        v9 += (v9 * (v9 ^ (v9 | v9)))
                    }
                    _ = fi19
                    _ = v20
                }
            }
        }
        _ = fi15
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
