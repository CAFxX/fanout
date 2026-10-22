package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint64; F1 uint64 }
type GenS3 struct { F0 uint8 }

func main() {
    var v1 int16 = -32768
    var v2 uint8 = 1
    var v3 int64 = -293
    v4 := make(map[uint32]bool)
    v1 -= (v1 / ((131) | 1))
    v5 := GenS1{F0: 2}
    var v6 int32 = 2147483647
    var v7 uint8 = (v2 - 1)
    var v8 int = -1
    for fi9 := 0; fi9 < 8; fi9++ {
        switch (fi9 | (^(fi9))) {
        case 9223372036854775807:
            v7 += v2
        default:
            var v10 float64 = (1e10 + (((-2.25 * 1.5) + -2.25) * (-2.25 / 0.1)))
            _ = v10
        }
        _ = fi9
    }
    for fi11 := 0; fi11 < 14; fi11++ {
        if false {
            for fi12 := 0; fi12 < 14; fi12++ {
                if false {
                    var v13 int32 = v6
                    fmt.Println(v13, v13, fi11)
                    _ = v13
                } else {
                    v2 = 255
                    var v14 uint32 = 458
                    v4[0] = (((!((!true) && false)) && false) || true)
                    if true {
                        _ = len(v4)
                        v8 += (v8 | fi12)
                        v6 += 1
                        if true {
                            v3 = 42
                        } else {
                            var v15 int8 = 24
                            v16 := []int{831, fi11, (fi11 | fi12)}
                            _ = len(v16)
                            _ = v15
                            _ = v16
                        }
                    } else {
                        switch fi11 {
                        case -9223372036854775808:
                            var v17 int64 = 9223372036854775807
                            _ = v17
                        case 114:
                            v18 := []int{}
                            _ = v18
                        }
                        delete(v4, 699)
                        v19 := GenS1{F0: uint64(v7)}
                        _ = v19
                    }
                    _ = v14
                }
                switch "hello" {
                case "":
                    var v20 int16 = v1
                    _ = v20
                case "a":
                    _ = len(v4)
                default:
                    v1 += (v1 >> ((18446744073709551615) & 63))
                }
                switch "\n\t" {
                case "a":
                    var v21 int64 = (v3 ^ (^(v3)))
                    _ = v21
                default:
                    v22 := GenS1{F0: 2}
                    _ = v22
                }
                var v23 int16 = (v1 / ((-188) | 1))
                v7 -= v7
                _ = fi12
                _ = v23
            }
        } else {
            v24 := []int{-9223372036854775808}
            v25 := GenS2{F0: 18446744073709551615, F1: 2}
            v24 = append(v24, (v8 / ((fi11) | 1)))
            _ = v24
            _ = v25
        }
        v1 -= 1000
        _ = fi11
    }
    v26 := GenS2{F0: 17, F1: 2}
    v27 := GenS3{F0: v7}
    var v28 int32 = 2147483647
    fmt.Println(v3, v3)
    v2 -= v2
    v1 -= -1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
