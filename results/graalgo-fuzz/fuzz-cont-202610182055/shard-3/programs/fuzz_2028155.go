package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int; F1 uint16 }
type GenS3 struct { F0 uint16; F1 uint64 }
type GenS4 struct { F0 uint64; F1 uint16 }
func genFn0(p0 uint16, p1 int, p2 string) int8 {
    var v7 uint64 = 18446744073709551615
    for fi8 := 0; fi8 < 1; fi8++ {
        v9 := []int{}
        _ = fi8
        _ = v9
    }
    var v10 int16 = 390
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v7
    return -128
}

func main() {
    v1 := []uint32{1, 2, 1}
    v2 := make(map[uint32]int)
    var v3 uint32 = 915
    var v4 int = 967
    var v5 uint64 = 663
    v6 := []int{v4}
    v5 += v5
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v11 int = v4
    var v12 string = "hello"
    var v13 uint64 = 34
    v14 := GenS1{F0: "hello"}
    var v15 int8 = 44
    if ((float64((((0.1 * 1.5) + 1e10) * 3.14159)) - (-2.25 - 0.1)) > (3.14159 + 0.0)) {
        v16 := GenS2{F0: (^(v11)), F1: 685}
        for fi17 := 0; fi17 < 19; fi17++ {
            fmt.Println(v13)
            v18 := GenS2{F0: v4, F1: 65535}
            v2[v3] = fi17
            _ = fi17
            _ = v18
        }
        v4 += 0
        _ = v16
    }
    if true {
        v19 := GenS3{F0: uint16(v3), F1: v5}
        _ = v19
    } else {
        v4 += (v11 & (v11 % ((v11) | 1)))
    }
    _ = len(v6)
    v20 := GenS4{F0: 18446744073709551615, F1: 939}
    _ = len(v2)
    switch (v11 + (v11 >> ((4294967295) & 63))) {
    case -2147483648:
        v11 -= (v11 | v11)
    }
    v21 := GenS2{F0: (v11 + v4), F1: 0}
    v4 += 0
    v22 := GenS1{F0: "foo bar"}
    for fi23 := 0; fi23 < 2; fi23++ {
        var v24 uint = 0
        if (true || true) {
            switch fi23 {
            case 42:
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
            case -9223372036854775808:
                delete(v2, v3)
            case 4294967295:
                var v25 float64 = -2.25
                _ = v25
            default:
                fmt.Println(v24, v12)
            }
            fmt.Println(v5)
        }
        _ = fi23
        _ = v24
    }
    fmt.Println(v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
