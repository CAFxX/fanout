package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 bool; F1 uint64 }
func genFn0(p0 int, p1 uint, p2 int64) uint32 {
    v6 := GenS1{F0: (^(p1))}
    v7 := GenS2{F0: p2}
    v8 := []int{(-(p0)), (p0 >> ((p1) & 63)), p0, (p0 & (p0 * p0))}
    var v9 uint32 = 965
    p0 += p0
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 65535
}
func genFn1(p0 uint32, p1 string, p2 uint32) int64 {
    if false {
        for fi10 := 0; fi10 < 8; fi10++ {
            p0 = 65535
            fmt.Println(p2)
            _ = fi10
        }
    } else {
        p1 = p1
        fmt.Println(p2, p2, p1)
    }
    for fi11 := 0; fi11 < 16; fi11++ {
        var v12 string = ("\n\t" + p1)
        fmt.Println(p1, p1, p1)
        _ = fi11
        _ = v12
    }
    var v13 float64 = 1e10
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return 9223372036854775807
}

func main() {
    v1 := []int{}
    var v2 int = 128
    v3 := make(map[string]int)
    v4 := make(map[int]string)
    v5 := make(map[uint32]bool)
    for fi14 := 0; fi14 < 20; fi14++ {
        var v15 float64 = 0.1
        v15 += 0.1
        _ = fi14
        _ = v15
    }
    v2 -= 42
    if true {
        if (false || (-1 <= -31)) {
            v16 := GenS3{F0: true, F1: 42}
            var v17 float64 = (0.0 * ((1.5 + (((0.0 - ((1e10 * float64(((3.14159 / 0.1) * 1.5))) - 1.5)) * ((1e10 + 0.0) - (-2.25 * 0.1))) - 1e10)) * 1e10))
            _ = v16
            _ = v17
        } else {
            if false {
                v2 = 127
                _ = len(v4)
                var v18 int16 = -32768
                _ = v18
            } else {
                v2 -= (v2 >> ((18446744073709551615) & 63))
                fmt.Println(v2)
            }
        }
        var v19 int32 = -1
        var v20 string = ""
        _ = v19
        _ = v20
    }
    for fi21 := 0; fi21 < 19; fi21++ {
        var v22 int = (v2 & -2147483648)
        v23 := GenS3{F0: true, F1: 0}
        _ = fi21
        _ = v22
        _ = v23
    }
    var v24 uint = 491
    v25 := GenS1{F0: (v24 + 4294967295)}
    fmt.Println(v24, v24)
    v2 = v2
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
