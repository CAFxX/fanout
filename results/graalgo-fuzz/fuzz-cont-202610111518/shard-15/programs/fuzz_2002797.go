package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8 }
type GenS2 struct { F0 int; F1 uint32; F2 int16 }
type GenS3 struct { F0 uint64; F1 int }
func genFn0(p0 int16, p1 uint32, p2 uint8) int64 {
    for fi7 := 0; fi7 < 12; fi7++ {
        if true {
            p0 += 1
        }
        v8 := []int{(-(fi7)), 42, fi7, (fi7 >> ((1) & 63))}
        p0 = (p0 - (p0 & 0))
        _ = fi7
        _ = v8
    }
    v9 := []int{4294967295, 712, -963, -254}
    v9 = append(v9, -1)
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 1
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int8 = -1
    var v3 bool = (50 == 359)
    var v4 int32 = 415
    v5 := make(map[string]int)
    var v6 float32 = ((1e10 / 1e10) + 1.5)
    v10 := GenS1{F0: int32(v2), F1: (v2 + v2)}
    var v11 uint16 = 65535
    v3 = v3
    v6 = v6
    v6 += v6
    for fi12 := 0; fi12 < 18; fi12++ {
        v13 := []int{(fi12 * fi12), fi12, 128}
        _ = fi12
        _ = v13
    }
    if true {
        v14 := GenS2{F0: -315, F1: 2147483647, F2: 8}
        v11 = (^(v11))
        _ = v14
    } else {
        _ = len(v1)
    }
    switch "" {
    case "foo bar":
        for fi15 := 0; fi15 < 13; fi15++ {
            var v16 int64 = -1
            if true {
                v17 := GenS1{F0: v4, F1: 127}
                switch -1 {
                case 32767:
                    v18 := []int{}
                    _ = v18
                case -992:
                    v19 := GenS3{F0: 18446744073709551615, F1: 2147483647}
                    _ = v19
                default:
                    fmt.Println(v6)
                }
                var v20 uint8 = 127
                _ = v17
                _ = v20
            }
            _ = fi15
            _ = v16
        }
    }
    switch -213 {
    case -564:
        var v21 uint32 = 65535
        _ = v21
    case -9223372036854775808:
        v22 := []int{-179, 521, 32767}
        _ = v22
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
