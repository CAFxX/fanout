package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 bool; F1 uint64; F2 int64 }
type GenS3 struct { F0 bool; F1 bool; F2 bool }
type GenS4 struct { F0 uint }
func genFn0(p0 uint, p1 bool, p2 int16) int64 {
    switch "a" {
    case "\n\t":
        var v5 int64 = -1
        _ = v5
    }
    var v6 float64 = 1e10
    var v7 int8 = -1
    switch "" {
    case "foo bar":
        p0 = p0
    case "":
        v8 := GenS1{F0: ("" + "foo bar")}
        _ = v8
    case "\n\t":
        var v9 int32 = 42
        _ = v9
    default:
        var v10 uint64 = 524
        _ = v10
    }
    var v11 int32 = 954
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v6
    _ = v7
    return -9223372036854775808
}

func main() {
    v1 := []int{-150, 42, 9223372036854775807}
    var v2 int16 = 42
    var v3 uint32 = 2
    var v4 int64 = 1
    var v12 uint8 = 42
    fmt.Println(v2, v2)
    v3 -= (v3 / (((v3 / ((812) | 1))) | 1))
    for fi13 := 0; fi13 < 19; fi13++ {
        v12 += (v12 % (((v12 - 128)) | 1))
        _ = fi13
    }
    v12 = 42
    fmt.Println(v2, v4, v3)
    switch -9223372036854775808 {
    case 42:
        v14 := GenS1{F0: ("\n\t" + "a")}
        _ = v14
    case 1:
        var v15 uint8 = (v12 / (((v12 - 128)) | 1))
        _ = v15
    }
    fmt.Println(v3, v3)
    v16 := GenS2{F0: true, F1: 18446744073709551615, F2: v4}
    v17 := GenS3{F0: false, F1: (4294967295 > -1), F2: (!(2 < 127))}
    if (!(true || false)) {
        if ((uint(v12) == 2) && true) {
            if false {
                fmt.Println(v12)
                var v18 uint64 = 18446744073709551615
                v2 = v2
                _ = v18
            }
        } else {
            v19 := GenS4{F0: 65535}
            var v20 int64 = 0
            _ = v19
            _ = v20
        }
        var v21 int8 = 127
        for fi22 := 0; fi22 < 14; fi22++ {
            v4 += genFn0(712, (false || true), (v2 ^ v2))
            fmt.Println(v12, fi22, v3)
            _ = fi22
        }
        _ = v21
    }
    for fi23 := 0; fi23 < 16; fi23++ {
        if true {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            var v24 int8 = -128
            _ = v24
        }
        v25 := GenS1{F0: "a"}
        _ = fi23
        _ = v25
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v1 = append(v1, -2147483648)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
