package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int; F2 int8 }
type GenS2 struct { F0 int; F1 uint64 }
func genFn0(p0 uint, p1 uint16) int32 {
    p0 -= p0
    var v7 int8 = -1
    v7 = v7
    var v8 uint32 = 595
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 2147483647
}
func genFn1(p0 int64, p1 string, p2 int) int8 {
    p1 = p1
    var v9 uint8 = 0
    if true {
        v10 := []int{p2, -1}
        if len(v10) >= 2 {
            v10 = v10[1:len(v10)]
        }
        var v11 uint8 = (v9 << ((65535) & 63))
        _ = v10
        _ = v11
    } else {
        fmt.Println(v9, p2)
        p2 = (^(p2))
        if (151 > 947) {
            v12 := GenS1{F0: (p2 ^ (p2 / ((1) | 1))), F1: (p2 >> ((255) & 63)), F2: -85}
            p1 = (p1 + p1)
            _ = v12
        } else {
            v13 := []int{p2}
            _ = v13
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 100
}

func main() {
    var v1 int = 42
    var v2 string = "a"
    v3 := []uint32{104, 221}
    var v4 uint32 = 171
    var v5 float64 = (((0.1 * 0.0) + ((0.1 / 1e10) + 1.5)) * ((1e10 + 3.14159) / (((float64(1e10) + ((3.14159 - 1.5) - 1.5)) + -2.25) - -2.25)))
    v6 := []int{v1, -801, -630}
    if ((1 >= -1) && (genFn0(491, 274) == -1)) {
        v14 := GenS1{F0: (v1 * (-(v1))), F1: (v1 & v1), F2: -1}
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        switch v2 {
        case "":
            for fi15 := 0; fi15 < 8; fi15++ {
                if (-1 > -1) {
                    var v16 uint = 18446744073709551615
                    _ = v16
                } else {
                    fmt.Println(fi15, v1)
                    var v17 uint32 = v4
                    var v18 uint32 = (v17 / ((v4) | 1))
                    _ = v17
                    _ = v18
                }
                _ = fi15
            }
        }
        _ = v14
    } else {
        v5 -= v5
    }
    if (2147483647 <= -1) {
        fmt.Println(v2, v4, v4)
    }
    v5 += v5
    var v19 float32 = -2.25
    var v20 float32 = v19
    v19 -= -2.25
    v4 = (v4 | 2147483647)
    fmt.Println(v1, v20, v4)
    var v21 bool = false
    if ((v4 / (((v4 - v4)) | 1)) >= (v4 % ((4294967295) | 1))) {
        for fi22 := 0; fi22 < 11; fi22++ {
            var v23 bool = (0 == 1)
            _ = fi22
            _ = v23
        }
        v4 = 4294967295
    } else {
        v24 := GenS2{F0: int(v4), F1: 2}
        switch v2 {
        case "\n\t":
            var v25 uint8 = 42
            _ = v25
        }
        _ = v24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
