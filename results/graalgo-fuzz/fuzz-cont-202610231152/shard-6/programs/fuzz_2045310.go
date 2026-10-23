package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 int8 }
type GenS2 struct { F0 int64; F1 uint32; F2 int }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint }
func genFn0(p0 uint64) uint {
    p0 = 18446744073709551615
    if (128 < 1) {
        v6 := GenS1{F0: 0, F1: 0, F2: 100}
        p0 = (p0 % (((p0 / ((p0) | 1))) | 1))
        _ = v6
    }
    v7 := GenS2{F0: -9223372036854775808, F1: 42, F2: 2147483647}
    v8 := GenS3{F0: 990}
    var v9 uint = 42
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return 665
}

func main() {
    var v1 float64 = 0.0
    var v2 uint8 = 136
    var v3 uint32 = 81
    var v4 int8 = -128
    v5 := make(map[uint32]string)
    var v10 uint16 = 2
    v2 += (v2 & (^(v2)))
    v2 += v2
    var v11 float64 = v1
    v12 := []int{517, -669}
    v13 := GenS4{F0: v3}
    var v14 bool = true
    var v15 int16 = int16(v3)
    v15 += v15
    v5[(v3 * (v3 % ((v3) | 1)))] = ("foo bar" + "a")
    v14 = true
    v16 := GenS3{F0: int16(v4)}
    for fi17 := 0; fi17 < 5; fi17++ {
        if ((-9223372036854775808 > -203) && true) {
            if len(v12) >= 2 {
                v12 = v12[1:len(v12)]
            }
        } else {
            v18 := GenS5{F0: uint(v10)}
            if true {
                var v19 string = "hello"
                for fi20 := 0; fi20 < 15; fi20++ {
                    v12 = append(v12, fi17)
                    v21 := v5[v3]
                    _ = fi20
                    _ = v21
                }
                _ = v19
            } else {
                var v22 uint16 = v10
                _ = len(v12)
                v23 := v5[(v3 | (v3 * v3))]
                _ = v22
                _ = v23
            }
            _ = v18
        }
        _ = fi17
    }
    for fi24 := 0; fi24 < 13; fi24++ {
        switch fi24 {
        case -43:
            var v25 uint64 = 602
            _ = v25
        case 9223372036854775807:
            if false {
                for fi26 := 0; fi26 < 10; fi26++ {
                    var v27 uint8 = 255
                    _ = fi26
                    _ = v27
                }
            }
        }
        _ = fi24
    }
    v10 = (v10 - v10)
    var v28 int = 128
    var v29 uint32 = (^(v3))
    v11 += v1
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
