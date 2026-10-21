package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8 }
type GenS2 struct { F0 int32; F1 uint; F2 int16 }
type GenS3 struct { F0 uint32; F1 int64; F2 bool }
func genFn0(p0 int8, p1 int) uint8 {
    var v7 int64 = -741
    if false {
        var v8 int64 = v7
        _ = v8
    } else {
        v9 := []int{(p1 ^ p1)}
        if len(v9) > 0 {
            v9[0] = (p1 / ((279) | 1))
        }
        for fi10 := 0; fi10 < 1; fi10++ {
            var v11 uint64 = 4294967295
            for fi12 := 0; fi12 < 4; fi12++ {
                var v13 float64 = float64(0.0)
                v14 := GenS1{F0: (p0 + p0), F1: 128}
                var v15 uint = 18446744073709551615
                if len(v9) > 0 {
                    v9[0] = 4294967295
                }
                _ = fi12
                _ = v13
                _ = v14
                _ = v15
            }
            v16 := GenS2{F0: -2147483648, F1: 18446744073709551615, F2: 166}
            _ = fi10
            _ = v11
            _ = v16
        }
        _ = v9
    }
    var v17 float64 = 0.0
    v18 := []int{(p1 | (p1 << ((18446744073709551615) & 63))), p1, (^(p1)), (^(p1))}
    _ = p0
    _ = p1
    _ = v17
    _ = v18
    _ = v7
    return 128
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint32 = 77
    var v3 bool = (42 > 1)
    v4 := make(map[int]uint32)
    v5 := make(map[string]uint32)
    var v6 uint8 = 1
    v19 := GenS1{F0: 42, F1: 2}
    if (!true) {
        switch "hello" {
        case "foo bar":
            var v20 int64 = -1
            _ = v20
        default:
            if true {
                fmt.Println(v2)
                switch -1 {
                case -600:
                    v3 = (507 == 0)
                }
                switch 128 {
                case -2147483648:
                    var v21 uint16 = 236
                    _ = v21
                case 679:
                    fmt.Println(v3, v2, v3)
                default:
                    v22 := []int{}
                    _ = v22
                }
            } else {
                v23 := []int{9223372036854775807, 32767, 128}
                for fi24 := 0; fi24 < 16; fi24++ {
                    v6 = 42
                    _ = fi24
                }
                switch 4294967295 {
                case 2147483647:
                    var v25 uint16 = 767
                    _ = v25
                case -836:
                    var v26 uint = uint(v2)
                    _ = v26
                default:
                    v23 = append(v23, -9223372036854775808)
                }
                _ = v23
            }
        }
    }
    switch "\n\t" {
    case "hello":
        switch 4294967295 {
        case 32767:
            v3 = (553 < 1000)
        case 4294967295:
            v6 = 42
        case 9223372036854775807:
            var v27 int32 = -886
            _ = v27
        default:
            var v28 int16 = 635
            _ = v28
        }
    case "":
        v3 = (100 == 1)
    case "a":
        var v29 float32 = -2.25
        _ = v29
    default:
        fmt.Println(v2, v6)
    }
    v30 := []int{860, 1, 1}
    v6 += 188
    var v31 int16 = 32767
    fmt.Println(v31, v3)
    var v32 bool = true
    v33 := GenS3{F0: (v2 * v2), F1: -611, F2: false}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
