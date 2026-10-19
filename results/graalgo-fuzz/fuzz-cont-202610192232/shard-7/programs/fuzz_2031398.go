package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8; F2 uint16 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int64; F1 uint16 }
func genFn0(p0 int16, p1 uint8, p2 int16) uint32 {
    v6 := []int{579, 4294967295, -342, -205}
    if (true || (!false)) {
        p1 += 0
    }
    v7 := GenS1{F0: -177, F1: 1, F2: 176}
    p1 += p1
    if false {
        p0 = (p2 | (p0 - 32767))
        switch "\n\t" {
        case "hello":
            p1 = p1
        }
        switch ("foo bar" + ((("\n\t" + "") + "\n\t") + "hello")) {
        case "hello":
            var v8 float32 = -2.25
            _ = v8
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 65535
}
func genFn1() uint32 {
    v9 := []int{151}
    var v10 int16 = -1
    v10 = 825
    var v11 float32 = float32(1e10)
    var v12 int64 = -9223372036854775808
    _ = v10
    _ = v11
    _ = v12
    _ = v9
    return 718
}

func main() {
    v1 := make(map[int]string)
    v2 := []int{438, 2147483647}
    var v3 int64 = 1
    v4 := make(map[uint32]int)
    var v5 string = "\n\t"
    var v13 int32 = 0
    var v14 int32 = (-(v13))
    for fi15 := 0; fi15 < 12; fi15++ {
        switch v5 {
        case "foo bar":
            v16 := GenS2{F0: (v3 >> ((929) & 63))}
            _ = v16
        default:
            if (false && ((255 == 0) || true)) {
                var v17 int8 = 7
                _ = v17
            } else {
                v18 := GenS3{F0: v3, F1: 368}
                v14 = v13
                _ = v18
            }
        }
        v5 = v5
        _ = fi15
    }
    v19 := GenS3{F0: (v3 | (v3 - v3)), F1: 255}
    v20 := GenS1{F0: 9223372036854775807, F1: 0, F2: 0}
    switch (v5 + v5) {
    case "hello":
        var v21 uint16 = 42
        _ = v21
    case "\n\t":
        fmt.Println(v13, v14, v13)
    default:
        fmt.Println(v13, v3)
    }
    if false {
        if (0.0 >= -2.25) {
            v3 = (v3 & v3)
        } else {
            if len(v2) > 0 {
                v2[0] = 4294967295
            }
            for fi22 := 0; fi22 < 19; fi22++ {
                fmt.Println(fi22, v13)
                v5 = v5
                v5 = v5
                _ = fi22
            }
            if (-235 < -1) {
                v13 -= (v13 ^ -459)
            } else {
                var v23 float32 = 0.0
                _ = v23
            }
        }
    }
    switch -173 {
    case 4294967295:
        fmt.Println(v3)
    default:
        var v24 int8 = 2
        _ = v24
    }
    fmt.Println(v3)
    v5 = v5
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
