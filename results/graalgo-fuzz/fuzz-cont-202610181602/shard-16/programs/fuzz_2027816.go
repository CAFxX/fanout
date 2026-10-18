package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32 }
type GenS2 struct { F0 uint16; F1 int16 }
type GenS3 struct { F0 int16; F1 uint64; F2 int32 }

func main() {
    v1 := make(map[int]string)
    v2 := []int{42}
    var v3 uint8 = 0
    v4 := make(map[string]bool)
    var v5 int64 = -9223372036854775808
    var v6 int16 = 1
    fmt.Println(v3, v5, v3)
    var v7 int8 = -108
    for fi8 := 0; fi8 < 9; fi8++ {
        v9 := GenS1{F0: (false || (v3 < (v3 + 42))), F1: 55}
        var v10 int32 = -922
        _ = fi8
        _ = v10
        _ = v9
    }
    var v11 int = -1
    for fi12 := 0; fi12 < 16; fi12++ {
        var v13 uint = 0
        for fi14 := 0; fi14 < 7; fi14++ {
            fmt.Println(v7, v11)
            if ((0.0 * 0.1) < (float32(1e10) + (0.0 / float32(-2.25)))) {
                v3 = v3
                var v15 uint64 = 0
                _ = v15
            } else {
                v7 = (^(v7))
                if true {
                    v16 := GenS1{F0: true, F1: -1}
                    v11 = (fi14 << (((v13 + 738)) & 63))
                    v17 := GenS2{F0: 32767, F1: v6}
                    switch "hello" {
                    case "a":
                        delete(v4, "a")
                    case "\n\t":
                        delete(v1, v11)
                    default:
                        v18 := GenS2{F0: 810, F1: (v6 & v6)}
                        _ = v18
                    }
                    v3 += (v3 - (v3 * (v3 >> (((v13 / ((v13) | 1))) & 63))))
                    switch fi14 {
                    case -1:
                        var v19 bool = true
                        _ = v19
                    case 4294967295:
                        switch fi12 {
                        case 4294967295:
                            v7 = v7
                        case 2147483647:
                            v5 += (v5 + 9223372036854775807)
                        case 128:
                            v13 = v13
                        default:
                            v7 += (v7 >> (((v13 | 4294967295)) & 63))
                        }
                        v11 = 127
                    default:
                        var v20 uint64 = 359
                        _ = v20
                    }
                    _ = v16
                    _ = v17
                } else {
                    _ = len(v2)
                    fmt.Println(v5, fi12)
                    v2 = append(v2, (fi14 % ((fi14) | 1)))
                }
                switch "" {
                case "hello":
                    var v21 float64 = float64(3.14159)
                    _ = v21
                case "foo bar":
                    fmt.Println(v11, v3)
                case "":
                    delete(v1, (fi12 % ((fi14) | 1)))
                }
            }
            v3 = v3
            _ = fi14
        }
        _ = fi12
        _ = v13
    }
    var v22 int = 32767
    v23 := GenS3{F0: v6, F1: 42, F2: 42}
    for fi24 := 0; fi24 < 7; fi24++ {
        v6 -= (v6 ^ int16(v7))
        _ = fi24
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
