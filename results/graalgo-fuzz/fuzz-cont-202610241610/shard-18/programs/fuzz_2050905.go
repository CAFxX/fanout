package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64; F2 uint }
type GenS2 struct { F0 int16 }
func genFn0(p0 int16, p1 int) uint8 {
    if true {
        switch (p1 & (p1 + (p1 & p1))) {
        case 0:
            if (11 >= 18446744073709551615) {
                p1 -= p1
            }
        case 9223372036854775807:
            v5 := []int{p1, p1}
            _ = v5
        default:
            v6 := GenS1{F0: 487, F1: 0, F2: 42}
            _ = v6
        }
        p0 -= (p0 ^ (p0 % (((p0 ^ p0)) | 1)))
    } else {
        var v7 int8 = -19
        _ = v7
    }
    switch p1 {
    case -512:
        var v8 int64 = 42
        _ = v8
    case 32767:
        v9 := GenS2{F0: p0}
        _ = v9
    case 0:
        for fi10 := 0; fi10 < 8; fi10++ {
            for fi11 := 0; fi11 < 20; fi11++ {
                if (-2.25 >= (((0.1 / 3.14159) * (-2.25 * -2.25)) - ((3.14159 * 1.5) * (3.14159 * 0.0)))) {
                    var v12 int8 = 42
                    _ = v12
                } else {
                    p0 = 32767
                    var v13 float64 = float64(1e10)
                    fmt.Println(fi10)
                    _ = v13
                }
                var v14 int16 = -1
                _ = fi11
                _ = v14
            }
            if ("" != "hello") {
                v15 := []int{fi10, fi10, (fi10 >> ((4294967295) & 63))}
                _ = v15
            }
            v16 := []int{}
            _ = fi10
            _ = v16
        }
    }
    var v17 uint32 = 42
    var v18 uint = 0
    _ = p0
    _ = p1
    _ = v17
    _ = v18
    return 151
}
func genFn1() uint16 {
    var v19 uint = 2
    v20 := GenS2{F0: 42}
    _ = v19
    _ = v20
    return 979
}

func main() {
    v1 := make(map[int]string)
    v2 := []string{("\n\t" + "hello")}
    var v3 uint8 = 0
    v4 := []string{}
    v3 += (v3 & (v3 & (v3 << ((2) & 63))))
    for fi21 := 0; fi21 < 1; fi21++ {
        var v22 int64 = 1
        _ = fi21
        _ = v22
    }
    v23 := GenS2{F0: 1000}
    if true {
        for fi24 := 0; fi24 < 1; fi24++ {
            var v25 int8 = -128
            _ = fi24
            _ = v25
        }
        v26 := GenS1{F0: 950, F1: 11, F2: 0}
        _ = len(v4)
        _ = v26
    }
    v3 = genFn0(42, 754)
    if false {
        var v27 int8 = 0
        _ = v27
    }
    for fi28 := 0; fi28 < 4; fi28++ {
        fmt.Println(fi28, fi28)
        v3 += v3
        fmt.Println(v3, v3)
        _ = fi28
    }
    var v29 bool = true
    fmt.Println(v3, v3)
    v3 += (v3 - (^(v3)))
    v3 = genFn0(-802, 9223372036854775807)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v23
    _ = v29
    _ = v3
    _ = v4
}
