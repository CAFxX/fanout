package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8 }
type GenS2 struct { F0 int16; F1 uint; F2 bool }

func main() {
    var v1 bool = ((42 < 87) && (true || true))
    var v2 float64 = (0.0 / 3.14159)
    var v3 float32 = (-2.25 / -2.25)
    v4 := make(map[int]bool)
    v5 := make(map[int]uint32)
    var v6 string = ("\n\t" + "a")
    fmt.Println(v3, v1)
    if (!false) {
        v3 += v3
        var v7 int8 = -1
        _ = v7
    } else {
        v6 = "foo bar"
        if (128 != 42) {
            for fi8 := 0; fi8 < 8; fi8++ {
                if (2147483648 >= 1) {
                    v3 -= (((float32(v2) + float32(v2)) - v3) * v3)
                }
                _ = fi8
            }
            for fi9 := 0; fi9 < 19; fi9++ {
                v10 := GenS1{F0: 17, F1: 13}
                _ = fi9
                _ = v10
            }
            v3 += v3
        }
        var v11 uint64 = 4294967295
        _ = v11
    }
    v6 = "foo bar"
    if (42 < -9223372036854775808) {
        var v12 float64 = v2
        _ = v12
    }
    v3 -= v3
    switch -655 {
    case -780:
        v13 := []int{-320, -621, 9223372036854775807, 32767}
        _ = v13
    case -2147483648:
        v2 -= (1.5 * v2)
    case -246:
        v14 := GenS1{F0: 0, F1: -1}
        _ = v14
    }
    switch 32768 {
    case 2147483647:
        switch 0 {
        case -9223372036854775808:
            var v15 uint32 = 65535
            _ = v15
        case 32767:
            var v16 int64 = -9223372036854775808
            _ = v16
        case 289:
            v17 := GenS1{F0: 42, F1: 7}
            _ = v17
        }
    case -444:
        v18 := []int{-258, 1, -2147483648, -737}
        _ = v18
    default:
        for fi19 := 0; fi19 < 17; fi19++ {
            if (42 <= 2) {
                var v20 int16 = 42
                _ = v20
            } else {
                var v21 uint32 = 1
                var v22 int32 = -1
                var v23 uint16 = 42
                _ = v21
                _ = v22
                _ = v23
            }
            for fi24 := 0; fi24 < 19; fi24++ {
                v25 := []int{fi19, (^(fi19)), (fi24 + (fi24 % ((fi24) | 1))), (^(fi19))}
                fmt.Println(v6, v6, fi19)
                _ = fi24
                _ = v25
            }
            _ = fi19
        }
    }
    v26 := GenS2{F0: 0, F1: 783, F2: (-126 == -1)}
    var v27 bool = ((42 < 4294967295) || false)
    v28 := GenS2{F0: 0, F1: 386, F2: (-1 <= -1)}
    v6 = (v6 + ("\n\t" + ("" + v6)))
    v29 := []int{-1, -117, 2147483647}
    v6 = v6
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
