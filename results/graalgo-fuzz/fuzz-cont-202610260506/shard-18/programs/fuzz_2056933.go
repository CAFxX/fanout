package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int64; F1 int16 }
type GenS3 struct { F0 int16; F1 uint16; F2 uint }
func genFn0(p0 uint16) int64 {
    p0 -= p0
    v7 := []int{-9223372036854775808, 2147483647}
    var v8 uint16 = (p0 / ((p0) | 1))
    _ = p0
    _ = v7
    _ = v8
    return 1
}

func main() {
    v1 := []uint32{2, 2147483648}
    v2 := make(map[int]int)
    var v3 string = ""
    v4 := []int{-767, 613}
    v5 := make(map[uint32]bool)
    var v6 uint16 = 2
    switch 1 {
    case 127:
        fmt.Println(v6, v6)
    }
    switch "" {
    case "foo bar":
        v9 := GenS1{F0: v6}
        _ = v9
    default:
        var v10 bool = (-1 >= -2147483648)
        _ = v10
    }
    switch 9223372036854775807 {
    case 259:
        v11 := GenS2{F0: 0, F1: 0}
        _ = v11
    default:
        for fi12 := 0; fi12 < 4; fi12++ {
            v13 := GenS1{F0: 1}
            switch 4294967295 {
            case -425:
                if (((float64(0.0) - 0.0) * (0.1 / 1e10)) > 0.1) {
                    v6 -= (v6 + v6)
                    fmt.Println(v3, v3)
                    var v14 int8 = 76
                    _ = v14
                } else {
                    var v15 float32 = (0.1 / 1.0)
                    _ = v15
                }
            }
            switch (v3 + "a") {
            case "\n\t":
                for fi16 := 0; fi16 < 14; fi16++ {
                    var v17 int8 = -21
                    var v18 uint64 = 873
                    v19 := GenS1{F0: v6}
                    v20 := GenS3{F0: 1000, F1: v6, F2: 159}
                    _ = fi16
                    _ = v17
                    _ = v18
                    _ = v19
                    _ = v20
                }
            default:
                for fi21 := 0; fi21 < 12; fi21++ {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    _ = fi21
                }
            }
            _ = fi12
            _ = v13
        }
    }
    switch "a" {
    case "hello":
        v3 = "a"
    }
    var v22 float32 = ((((-2.25 + 3.14159) * 3.14159) + -2.25) + 3.14159)
    switch -9223372036854775808 {
    case -759:
        v3 = v3
    }
    v4 = append(v4, 32767)
    for fi23 := 0; fi23 < 2; fi23++ {
        var v24 uint16 = v6
        _ = fi23
        _ = v24
    }
    fmt.Println(v6, v3, v22)
    v3 = v3
    var v25 float64 = 1e10
    var v26 float32 = v22
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
