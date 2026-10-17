package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int; F2 int32 }
type GenS2 struct { F0 uint32 }
func genFn0(p0 int, p1 uint32) int8 {
    v4 := GenS1{F0: -1, F1: 32767, F2: 1}
    v5 := GenS1{F0: int32(p0), F1: 9223372036854775807, F2: -426}
    if true {
        var v6 uint8 = 47
        if (2 != 0) {
            v7 := []int{p0, (p0 >> ((42) & 63))}
            _ = v7
        }
        var v8 int32 = -1
        _ = v6
        _ = v8
    }
    switch (p0 * p0) {
    case 9223372036854775807:
        var v9 uint = 18446744073709551615
        _ = v9
    case 1:
        v10 := []int{2147483647}
        _ = v10
    case 32767:
        if true {
            switch p0 {
            case 4294967295:
                for fi11 := 0; fi11 < 3; fi11++ {
                    p0 -= fi11
                    var v12 uint32 = 42
                    v13 := GenS2{F0: (v12 % ((v12) | 1))}
                    _ = fi11
                    _ = v12
                    _ = v13
                }
            case 32767:
                fmt.Println(p1, p0)
            }
            for fi14 := 0; fi14 < 9; fi14++ {
                p1 = (p1 * (^(p1)))
                _ = fi14
            }
        } else {
            p0 = -1
        }
    default:
        if false {
            var v15 int32 = -367
            var v16 int32 = v15
            _ = v15
            _ = v16
        } else {
            p1 -= p1
            var v17 uint32 = (p1 / ((p1) | 1))
            fmt.Println(p0)
            _ = v17
        }
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 1
}
func genFn1(p0 uint, p1 uint8) int16 {
    var v18 string = "a"
    switch 32768 {
    case 128:
        v19 := []int{}
        _ = v19
    case 1:
        p0 = (p0 % (((p0 & (p0 - uint(p1)))) | 1))
    }
    _ = p0
    _ = p1
    _ = v18
    return -139
}

func main() {
    v1 := make(map[uint32]int)
    var v2 uint16 = 822
    var v3 uint = 53
    var v20 int32 = 42
    if false {
        v2 = v2
        var v21 uint64 = 109
        _ = v21
    }
    var v22 uint64 = uint64(v3)
    var v23 float64 = float64(1.5)
    for fi24 := 0; fi24 < 19; fi24++ {
        delete(v1, 0)
        _ = fi24
    }
    v22 -= v22
    var v25 uint16 = 703
    var v26 float64 = v23
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v20
    _ = v22
    _ = v23
    _ = v25
    _ = v26
    _ = v3
}
