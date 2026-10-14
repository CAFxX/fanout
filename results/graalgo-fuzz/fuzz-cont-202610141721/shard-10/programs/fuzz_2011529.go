package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64 }
type GenS2 struct { F0 int8 }
func genFn0(p0 uint, p1 uint32, p2 uint8) int16 {
    var v7 int = 832
    if false {
        v7 = v7
    } else {
        v7 = (v7 | -9223372036854775808)
        v8 := GenS1{F0: p2, F1: 18446744073709551615}
        v7 -= (v7 - v7)
        _ = v8
    }
    var v9 float32 = 1e10
    if true {
        v9 = (1e10 + float32((1e10 * ((-2.25 - ((1e10 / 0.1) / 1.0)) * 1.5))))
        for fi10 := 0; fi10 < 20; fi10++ {
            fmt.Println(v9, v9)
            for fi11 := 0; fi11 < 20; fi11++ {
                v7 = (fi10 ^ fi11)
                var v12 uint = p0
                _ = fi11
                _ = v12
            }
            var v13 uint32 = p1
            _ = fi10
            _ = v13
        }
        p2 = p2
    }
    var v14 float32 = v9
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    _ = v7
    _ = v9
    return 42
}
func genFn1(p0 int16, p1 uint64, p2 bool) int64 {
    v15 := []int{290, -1, 126}
    p1 -= (p1 % ((862) | 1))
    fmt.Println(p2, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    return -1
}

func main() {
    v1 := make(map[uint32]string)
    v2 := []uint32{646, 2147483647}
    v3 := []int{9223372036854775807}
    var v4 uint8 = 124
    var v5 uint16 = 0
    var v6 float32 = (1e10 + 1e10)
    v4 -= (v4 % ((v4) | 1))
    v16 := GenS2{F0: -128}
    fmt.Println(v4)
    v17 := v1[0]
    v18 := GenS2{F0: 46}
    for fi19 := 0; fi19 < 17; fi19++ {
        for fi20 := 0; fi20 < 1; fi20++ {
            v21 := GenS2{F0: 1}
            v6 += -2.25
            _ = fi20
            _ = v21
        }
        v17 = (v17 + v17)
        v2 = append(v2, 566)
        _ = fi19
    }
    fmt.Println(v6, v6)
    _ = len(v1)
    var v22 uint = 744
    v23 := GenS1{F0: 128, F1: 42}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
