package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint16; F1 int32; F2 string }
type GenS4 struct { F0 uint; F1 int16; F2 bool }
func genFn0(p0 uint, p1 bool, p2 int16) uint16 {
    fmt.Println(p2, p0, p1)
    if ((p0 - p0) < p0) {
        p2 += (p2 ^ p2)
        var v6 int64 = 754
        _ = v6
    } else {
        v7 := GenS1{F0: 42, F1: 0}
        for fi8 := 0; fi8 < 13; fi8++ {
            for fi9 := 0; fi9 < 19; fi9++ {
                var v10 float64 = (0.0 / 1.0)
                v11 := []int{fi9}
                if (945 == uint64(p0)) {
                    p1 = true
                    p0 = (^(p0))
                }
                fmt.Println(p0, fi8)
                v12 := GenS1{F0: 9223372036854775807, F1: 759}
                fmt.Println(v10)
                _ = fi9
                _ = v10
                _ = v11
                _ = v12
            }
            fmt.Println(p2, fi8)
            v13 := []int{}
            _ = fi8
            _ = v13
        }
        v14 := []int{}
        _ = v14
        _ = v7
    }
    p1 = (false && p1)
    _ = p0
    _ = p1
    _ = p2
    return 255
}

func main() {
    var v1 uint16 = 2
    var v2 uint16 = (v1 & (v1 + 42))
    v3 := make(map[int]bool)
    var v4 uint32 = 0
    var v5 uint64 = 2
    v4 = v4
    v15 := GenS2{F0: false}
    switch 127 {
    case -911:
        var v16 bool = (!(2 < 1))
        _ = v16
    default:
        fmt.Println(v4, v5, v1)
    }
    for fi17 := 0; fi17 < 8; fi17++ {
        var v18 float32 = float32(1.5)
        _ = fi17
        _ = v18
    }
    var v19 int16 = 32767
    if (60 <= -1) {
        var v20 uint = 18446744073709551615
        _ = v20
    }
    v1 += v1
    if ((((^(v4)) > v4) || (419 >= 2147483647)) && true) {
        if (20 < 189) {
            v2 = 255
            v21 := []int{-59, 32768, 32767}
            var v22 uint16 = uint16(v5)
            _ = v21
            _ = v22
        } else {
            fmt.Println(v4, v2, v19)
        }
    }
    delete(v3, -9223372036854775808)
    if ((v4 << ((215) & 63)) > 0) {
        fmt.Println(v4, v19)
        v23 := GenS2{F0: true}
        var v24 float32 = (0.1 - 0.1)
        _ = v23
        _ = v24
    }
    v25 := GenS1{F0: -270, F1: v5}
    if true {
        v4 -= v4
    }
    fmt.Println(v19, v1)
    var v26 int16 = (v19 ^ (v19 + v19))
    v27 := v3[32767]
    v1 += (v1 | v2)
    v28 := GenS3{F0: genFn0(uint(v2), v27, 32767), F1: 1, F2: (("\n\t" + "\n\t") + "foo bar")}
    switch 127 {
    case 600:
        switch ("a" + "foo bar") {
        case "a":
            v29 := []int{-2147483648}
            _ = v29
        case "\n\t":
            v26 += (v26 >> ((18446744073709551615) & 63))
        }
    case 32767:
        v30 := GenS3{F0: uint16(v4), F1: -981, F2: "\n\t"}
        _ = v30
    case 1:
        v31 := GenS4{F0: 65535, F1: v26, F2: ((1.5 / 3.14159) > 1e10)}
        _ = v31
    default:
        v4 = v4
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v19
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
