package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64; F2 int }
type GenS2 struct { F0 uint64; F1 uint8; F2 int8 }
type GenS3 struct { F0 uint32; F1 int8 }
type GenS4 struct { F0 string; F1 int; F2 int64 }
type GenS5 struct { F0 int16 }

func main() {
    var v1 uint64 = 2
    v2 := make(map[string]bool)
    v3 := make(map[string]int)
    var v4 float64 = ((float64(((1.5 - 0.0) / -2.25)) - 0.0) - -2.25)
    var v5 int8 = -128
    v6 := GenS1{F0: -1, F1: v1, F2: 9223372036854775807}
    if (v5 > (v5 * v5)) {
        if false {
            v7 := v3["a"]
            fmt.Println(v4, v7, v7)
            _ = v7
        }
        v8 := GenS1{F0: -1, F1: v1, F2: 961}
        v4 += v4
        _ = v8
    }
    for fi9 := 0; fi9 < 17; fi9++ {
        if ("" >= "hello") {
            v10 := GenS2{F0: 18446744073709551615, F1: 127, F2: (^(v5))}
            fmt.Println(v5, v4)
            var v11 uint8 = 2
            _ = v10
            _ = v11
        } else {
            switch "" {
            case "hello":
                if false {
                    for fi12 := 0; fi12 < 7; fi12++ {
                        fmt.Println(v5, v1)
                        var v13 uint = 0
                        v14 := GenS3{F0: 2147483648, F1: (^(v5))}
                        _ = fi12
                        _ = v13
                        _ = v14
                    }
                } else {
                    v5 += (^(v5))
                    v15 := GenS4{F0: "", F1: (fi9 ^ fi9), F2: 965}
                    _ = v15
                }
            default:
                for fi16 := 0; fi16 < 13; fi16++ {
                    switch "a" {
                    case "\n\t":
                        if (0 <= v1) {
                            var v17 int64 = 0
                            _ = v17
                        } else {
                            var v18 uint32 = 4294967295
                            _ = v18
                        }
                    case "hello":
                        var v19 int8 = (v5 >> ((42) & 63))
                        _ = v19
                    }
                    _ = fi16
                }
            }
            v20 := GenS5{F0: 1}
            var v21 int64 = 42
            _ = v20
            _ = v21
        }
        if (-2147483648 >= int32(v5)) {
            fmt.Println(v1, v5)
        } else {
            if true {
                v22 := GenS5{F0: -369}
                v1 -= 367
                v1 += 759
                _ = v22
            }
        }
        var v23 float32 = ((0.0 - 1e10) - 0.1)
        _ = fi9
        _ = v23
    }
    for fi24 := 0; fi24 < 8; fi24++ {
        v25 := []int{fi24, fi24}
        var v26 float64 = v4
        _ = fi24
        _ = v25
        _ = v26
    }
    for fi27 := 0; fi27 < 1; fi27++ {
        fmt.Println(v5, v4, v4)
        _ = fi27
    }
    _ = len(v2)
    v28 := []int{-2147483648, -2147483648, 9223372036854775807, 796}
    v4 = v4
    v5 -= (v5 + 100)
    var v29 int32 = 0
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
