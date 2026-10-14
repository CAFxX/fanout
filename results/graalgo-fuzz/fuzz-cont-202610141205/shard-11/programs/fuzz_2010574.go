package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64; F2 uint32 }
type GenS2 struct { F0 uint16; F1 uint }
func genFn0(p0 uint16, p1 uint16, p2 string) uint32 {
    fmt.Println(p0, p1, p2)
    p1 -= (p1 >> ((2) & 63))
    var v5 bool = true
    fmt.Println(p0, p2, p1)
    fmt.Println(p1, p1, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 42
}
func genFn1(p0 uint) uint32 {
    v6 := GenS1{F0: -555, F1: 844, F2: 482}
    fmt.Println(p0)
    var v7 int64 = 0
    for fi8 := 0; fi8 < 10; fi8++ {
        var v9 uint16 = 42
        if false {
            var v10 uint32 = 42
            v9 = v9
            var v11 uint32 = v10
            _ = v10
            _ = v11
        } else {
            if (1e10 < (0.0 + (1e10 + -2.25))) {
                fmt.Println(v9, v7)
            }
            for fi12 := 0; fi12 < 5; fi12++ {
                v13 := GenS2{F0: (v9 >> (((^(p0))) & 63)), F1: p0}
                _ = fi12
                _ = v13
            }
            p0 = p0
        }
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = v6
    _ = v7
    return 2147483648
}

func main() {
    var v1 uint = 4294967295
    v2 := []int{-439, -189}
    var v3 int64 = 0
    var v4 int = 357
    v1 = v1
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v1 = 106
    v3 = v3
    var v14 uint8 = 14
    var v15 uint32 = 42
    for fi16 := 0; fi16 < 13; fi16++ {
        v17 := GenS2{F0: 65535, F1: 0}
        _ = fi16
        _ = v17
    }
    fmt.Println(v4, v15)
    v18 := GenS2{F0: 32767, F1: (^(v1))}
    switch (("a" + "foo bar") + "a") {
    case "hello":
        fmt.Println(v4)
    case "":
        var v19 uint8 = v14
        _ = v19
    }
    if len(v2) > 0 {
        v2[0] = -134
    }
    switch v4 {
    case 32767:
        var v20 string = "foo bar"
        _ = v20
    case 42:
        v15 += (^(v15))
    case -52:
        v3 += 0
    }
    if true {
        fmt.Println(v1, v4)
    } else {
        v1 += v1
        if (-1 > -1) {
            v14 = (v14 & v14)
            v2 = append(v2, (v4 >> (((v1 + (v1 * (v1 | (v1 % (((v1 * v1)) | 1)))))) & 63)))
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
        var v21 int8 = -1
        _ = v21
    }
    for fi22 := 0; fi22 < 16; fi22++ {
        if true {
            v14 = v14
            v14 -= v14
            v15 += (v15 | v15)
        }
        for fi23 := 0; fi23 < 16; fi23++ {
            var v24 float64 = 0.0
            _ = fi23
            _ = v24
        }
        v25 := GenS2{F0: 1, F1: (v1 / ((65535) | 1))}
        _ = fi22
        _ = v25
    }
    var v26 uint64 = 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v18
    _ = v2
    _ = v26
    _ = v3
    _ = v4
}
