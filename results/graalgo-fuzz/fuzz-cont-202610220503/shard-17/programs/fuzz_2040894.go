package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int64; F1 uint16 }
type GenS3 struct { F0 int32; F1 int; F2 uint }

func main() {
    var v1 int16 = 0
    v2 := []int{4294967295, 128, -2147483648, 32768}
    v3 := []int{-885}
    var v4 bool = (2 > 2147483647)
    v4 = v4
    if (((false || (v1 != v1)) || false) && v4) {
        switch 4294967295 {
        case -663:
            _ = len(v2)
        case 42:
            fmt.Println(v1, v4)
        case 2147483647:
            for fi5 := 0; fi5 < 6; fi5++ {
                v4 = true
                _ = fi5
            }
        default:
            var v6 uint = 831
            _ = v6
        }
        var v7 uint64 = 4294967295
        _ = v7
    }
    var v8 uint16 = 338
    if (425 != 1) {
        if v4 {
            var v9 int8 = -114
            v1 = (v1 << ((4294967295) & 63))
            _ = v9
        } else {
            for fi10 := 0; fi10 < 19; fi10++ {
                v1 = (v1 ^ v1)
                _ = fi10
            }
        }
        fmt.Println(v4, v1, v1)
    }
    for fi11 := 0; fi11 < 14; fi11++ {
        if true {
            if v4 {
                v4 = ((1 <= 0) && false)
                if len(v3) > 0 {
                    v3[0] = (fi11 | fi11)
                }
            }
            fmt.Println(fi11)
            for fi12 := 0; fi12 < 2; fi12++ {
                v1 += (v1 >> ((509) & 63))
                if v4 {
                    v8 += 710
                    fmt.Println(v1, v8, fi11)
                    v13 := GenS1{F0: 2147483647}
                    switch fi12 {
                    case -9223372036854775808:
                        v8 += v8
                    case 42:
                        var v14 int = 545
                        _ = v14
                    case 127:
                        v15 := GenS2{F0: 41, F1: v8}
                        _ = v15
                    }
                    _ = v13
                }
                _ = fi12
            }
        }
        v1 -= (v1 * v1)
        if v4 {
            v16 := GenS3{F0: 42, F1: fi11, F2: 4294967295}
            switch "" {
            case "a":
                _ = len(v2)
            default:
                v2 = append(v2, fi11)
            }
            fmt.Println(fi11, v1, v4)
            _ = v16
        }
        _ = fi11
    }
    var v17 uint = 42
    var v18 int32 = 2147483647
    v18 -= v18
    switch "foo bar" {
    case "hello":
        if ((1e10 - -2.25) == -2.25) {
            if len(v3) > 0 {
                v3[0] = -771
            }
        }
    }
    var v19 uint16 = v8
    v4 = (true || true)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
