package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8 }
type GenS2 struct { F0 uint16; F1 int16 }
type GenS3 struct { F0 uint16; F1 int; F2 int8 }
type GenS4 struct { F0 uint }
func genFn0(p0 string, p1 uint) uint64 {
    var v7 float64 = float64(0.1)
    var v8 uint32 = 408
    if true {
        v9 := []int{2147483647, 0}
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 2
}

func main() {
    var v1 int64 = -1
    v2 := make(map[uint32]uint32)
    v3 := make(map[string]string)
    v4 := make(map[int]bool)
    var v5 uint64 = 18446744073709551615
    var v6 int16 = -328
    fmt.Println(v6, v6, v5)
    if (("" <= "a") && (330 <= 910)) {
        if ((!((114 != 0) || (589 == 1))) || false) {
            v10 := GenS1{F0: 1, F1: 0}
            var v11 float64 = 1e10
            for fi12 := 0; fi12 < 10; fi12++ {
                var v13 int64 = 1
                v5 = (v5 - v5)
                _ = fi12
                _ = v13
            }
            _ = v10
            _ = v11
        } else {
            v14 := GenS2{F0: 42, F1: v6}
            var v15 float64 = 0.0
            _ = v14
            _ = v15
        }
        v5 -= (v5 ^ genFn0("\n\t", 2))
    } else {
        fmt.Println(v1)
    }
    v5 = (v5 - genFn0("\n\t", 358))
    v5 = 857
    v5 = v5
    if true {
        switch "\n\t" {
        case "a":
            v5 = v5
        case "hello":
            v6 = 0
        }
        v16 := GenS3{F0: 65535, F1: 9223372036854775807, F2: 100}
        _ = v16
    } else {
        v1 += v1
        v17 := []int{9223372036854775807}
        v18 := GenS4{F0: 18446744073709551615}
        _ = v17
        _ = v18
    }
    v2[65535] = 0
    v19 := []int{}
    for fi20 := 0; fi20 < 6; fi20++ {
        if true {
            v5 += (v5 ^ (v5 * (v5 << ((255) & 63))))
        }
        for fi21 := 0; fi21 < 6; fi21++ {
            v1 = v1
            _ = fi21
        }
        if len(v19) >= 2 {
            v19 = v19[1:len(v19)]
        }
        _ = fi20
    }
    var v22 bool = true
    var v23 string = "hello"
    fmt.Println(v23, v23)
    var v24 int32 = 2147483647
    switch "foo bar" {
    case "a":
        if len(v19) >= 2 {
            v19 = v19[1:len(v19)]
        }
    default:
        var v25 int64 = (v1 >> ((615) & 63))
        _ = v25
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
