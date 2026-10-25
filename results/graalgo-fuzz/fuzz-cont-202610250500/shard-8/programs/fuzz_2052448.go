package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32 }
type GenS2 struct { F0 uint; F1 string; F2 int }
type GenS3 struct { F0 int64 }
func genFn0(p0 int8, p1 int32, p2 int8) uint16 {
    v4 := []int{-946}
    v5 := GenS1{F0: 1, F1: 807}
    var v6 uint16 = 255
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    return 447
}
func genFn1(p0 int32, p1 int64, p2 int64) int64 {
    fmt.Println(p1, p2)
    switch "hello" {
    case "\n\t":
        var v7 uint64 = 1
        _ = v7
    case "a":
        var v8 int = 32767
        _ = v8
    case "":
        v9 := []int{162, -9223372036854775808, 1}
        _ = v9
    }
    var v10 uint64 = 1
    p2 = -9223372036854775808
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return (-(p2))
}

func main() {
    var v1 uint = 13
    v2 := []uint32{}
    var v3 int64 = 9223372036854775807
    v11 := GenS2{F0: 2, F1: ("foo bar" + ("" + "")), F2: 9223372036854775807}
    v12 := GenS2{F0: (v1 >> ((v1) & 63)), F1: ("\n\t" + ""), F2: 924}
    fmt.Println(v1, v3)
    if true {
        fmt.Println(v3, v3)
        fmt.Println(v1, v1)
    }
    v3 += 0
    v13 := GenS3{F0: (v3 / (((v3 << ((4294967295) & 63))) | 1))}
    v3 += (v3 | (v3 & (v3 + v3)))
    var v14 uint8 = 127
    if true {
        v3 = v3
        v14 += (v14 | v14)
        fmt.Println(v1, v3)
    } else {
        v2 = append(v2, 2)
    }
    v3 -= (v3 >> (((v1 % (((v1 / (((v1 & (v1 ^ v1))) | 1))) | 1))) & 63))
    if (-1 <= -119) {
        for fi15 := 0; fi15 < 1; fi15++ {
            var v16 int = (^(fi15))
            if (true || false) {
                if (0.1 != ((1.5 + (0.1 - 1e10)) + 1e10)) {
                    var v17 string = ""
                    v3 = genFn1(350, v3, genFn1(-1, (v3 | -113), v3))
                    _ = v17
                }
                v1 = v1
            }
            var v18 int16 = -1
            _ = fi15
            _ = v16
            _ = v18
        }
        v14 += v14
    }
    if true {
        var v19 uint8 = v14
        var v20 int16 = -1
        _ = len(v2)
        _ = v19
        _ = v20
    }
    if false {
        if len(v2) > 0 {
            v2[0] = 4294967295
        }
        switch ("" + "a") {
        case "\n\t":
            v21 := GenS2{F0: v1, F1: ("a" + "\n\t"), F2: 1}
            _ = v21
        case "hello":
            fmt.Println(v14, v3, v1)
        default:
            var v22 uint = 4294967295
            _ = v22
        }
    }
    var v23 int16 = 42
    var v24 float64 = 1.5
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v23
    _ = v24
    _ = v3
}
