package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64; F2 int }
type GenS2 struct { F0 int8; F1 uint8; F2 uint }
func genFn0(p0 int64) uint8 {
    for fi4 := 0; fi4 < 3; fi4++ {
        v5 := GenS1{F0: 42, F1: p0, F2: fi4}
        fmt.Println(p0, p0)
        _ = fi4
        _ = v5
    }
    fmt.Println(p0)
    v6 := []int{-1, int(p0), 9223372036854775807, 127}
    _ = p0
    _ = v6
    return 127
}

func main() {
    var v1 uint64 = 42
    var v2 int64 = -406
    v3 := []string{}
    fmt.Println(v2)
    if false {
        var v7 float64 = -2.25
        var v8 uint32 = 2147483647
        _ = v7
        _ = v8
    } else {
        _ = len(v3)
        for fi9 := 0; fi9 < 20; fi9++ {
            for fi10 := 0; fi10 < 5; fi10++ {
                v1 += (v1 | (v1 | (v1 % ((v1) | 1))))
                v11 := GenS1{F0: 967, F1: v2, F2: fi9}
                var v12 float32 = 0.0
                var v13 float64 = 1.5
                _ = fi10
                _ = v11
                _ = v12
                _ = v13
            }
            for fi14 := 0; fi14 < 10; fi14++ {
                _ = len(v3)
                _ = fi14
            }
            _ = fi9
        }
        var v15 uint64 = v1
        _ = v15
    }
    for fi16 := 0; fi16 < 17; fi16++ {
        fmt.Println(fi16, v1, v1)
        _ = fi16
    }
    v17 := GenS2{F0: 127, F1: 42, F2: 159}
    v18 := GenS1{F0: 829, F1: v2, F2: 9223372036854775807}
    for fi19 := 0; fi19 < 20; fi19++ {
        var v20 uint = 714
        _ = fi19
        _ = v20
    }
    var v21 uint16 = 1
    var v22 uint64 = (v1 - v1)
    v23 := GenS2{F0: -34, F1: 128, F2: 65535}
    v1 += (v1 * v1)
    var v24 uint = 589
    if false {
        var v25 bool = ((!false) || true)
        v22 = (v1 / ((42) | 1))
        v1 += (v1 << ((0) & 63))
        _ = v25
    }
    if true {
        var v26 uint8 = 127
        for fi27 := 0; fi27 < 9; fi27++ {
            v28 := GenS2{F0: -128, F1: v26, F2: (v24 & 42)}
            _ = fi27
            _ = v28
        }
        _ = v26
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v24 -= v24
    v22 = v1
    var v29 int64 = 9223372036854775807
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v29
    _ = v3
}
