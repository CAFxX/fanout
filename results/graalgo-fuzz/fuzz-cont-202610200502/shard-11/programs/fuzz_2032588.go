package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64 }
type GenS2 struct { F0 uint64; F1 uint16; F2 uint16 }
type GenS3 struct { F0 uint16; F1 uint8; F2 uint64 }
func genFn0(p0 int16, p1 uint32) uint32 {
    switch -919 {
    case -803:
        switch 32768 {
        case -613:
            p0 += -915
        default:
            for fi6 := 0; fi6 < 1; fi6++ {
                var v7 uint32 = 65535
                v8 := []int{}
                if false {
                    v7 = (p1 * 65535)
                    v8 = append(v8, 1)
                    if false {
                        fmt.Println(p0)
                    }
                } else {
                    switch "" {
                    case "\n\t":
                        v9 := GenS1{F0: 100, F1: -9223372036854775808}
                        _ = v9
                    }
                    var v10 int16 = p0
                    _ = v10
                }
                _ = fi6
                _ = v7
                _ = v8
            }
        }
    default:
        v11 := GenS1{F0: 127, F1: -1}
        _ = v11
    }
    v12 := []int{}
    v13 := GenS1{F0: 83, F1: 42}
    v14 := GenS2{F0: 42, F1: 65535, F2: 2}
    v15 := GenS2{F0: 945, F1: 2, F2: 42}
    _ = p0
    _ = p1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    return 208
}
func genFn1(p0 int32, p1 string) bool {
    v16 := GenS2{F0: 1, F1: 65535, F2: 576}
    p0 -= p0
    _ = p0
    _ = p1
    _ = v16
    return (0 > -1)
}

func main() {
    var v1 int64 = 9223372036854775807
    v2 := make(map[string]bool)
    var v3 float32 = -2.25
    v4 := make(map[string]uint32)
    var v5 int8 = -128
    fmt.Println(v1, v3, v3)
    v3 += (3.14159 - v3)
    for fi17 := 0; fi17 < 11; fi17++ {
        v4["a"] = 2
        _ = fi17
    }
    for fi18 := 0; fi18 < 7; fi18++ {
        v19 := []int{}
        v1 += (v1 + (v1 ^ v1))
        _ = fi18
        _ = v19
    }
    var v20 float32 = v3
    v21 := v4["foo bar"]
    v5 = (v5 << ((65535) & 63))
    v1 = v1
    v22 := GenS2{F0: 2, F1: 826, F2: 0}
    v23 := GenS3{F0: 210, F1: 255, F2: 783}
    v1 = (-(v1))
    v24 := GenS3{F0: 739, F1: 0, F2: 236}
    var v25 uint32 = v21
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
