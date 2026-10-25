package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 bool; F1 uint16 }
type GenS3 struct { F0 string }
func genFn0(p0 int8, p1 uint16) int64 {
    p0 = p0
    if (!false) {
        switch "a" {
        case "a":
            fmt.Println(p0, p0)
        case "":
            p0 += p0
        case "hello":
            p1 = p1
        default:
            fmt.Println(p1, p1)
        }
    }
    p0 = p0
    _ = p0
    _ = p1
    return -1
}
func genFn1(p0 bool) int16 {
    fmt.Println(p0)
    p0 = (32767 <= 372)
    for fi5 := 0; fi5 < 9; fi5++ {
        v6 := GenS1{F0: -128}
        _ = fi5
        _ = v6
    }
    _ = p0
    return 362
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []uint32{}
    v3 := []uint32{42}
    var v4 int8 = 1
    v7 := v1[155]
    var v8 int64 = 1
    _ = len(v2)
    var v9 bool = (921 != 0)
    if (42 <= 2) {
        var v10 uint8 = 9
        var v11 uint8 = (v10 ^ (^(v10)))
        _ = v10
        _ = v11
    } else {
        var v12 int8 = (v4 ^ (v4 >> ((255) & 63)))
        for fi13 := 0; fi13 < 6; fi13++ {
            var v14 uint = 65535
            switch (fi13 % ((fi13) | 1)) {
            case 128:
                v8 = -9223372036854775808
            case 2147483647:
                if (v14 <= (v14 & (v14 & 691))) {
                    var v15 int64 = (v8 | v8)
                    _ = v15
                }
            default:
                v16 := GenS1{F0: v12}
                _ = v16
            }
            _ = fi13
            _ = v14
        }
        _ = v12
    }
    if ((!(3.14159 <= 0.0)) && (v9 && v9)) {
        v9 = (!(0 <= 0))
        v17 := GenS2{F0: v9, F1: 600}
        _ = v17
    } else {
        switch "foo bar" {
        case "":
            if len(v3) > 0 {
                v3[0] = v7
            }
        case "a":
            if (v8 >= (v8 << ((uint(v7)) & 63))) {
                v4 = -1
            }
        default:
            var v18 int32 = 42
            _ = v18
        }
    }
    if (4294967295 > 18446744073709551615) {
        delete(v1, (v7 | (^(v7))))
        var v19 float64 = 1e10
        _ = v19
    } else {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        v20 := GenS3{F0: "a"}
        _ = v20
    }
    fmt.Println(v7, v4)
    v8 = (v8 + (v8 >> ((65535) & 63)))
    var v21 uint8 = 153
    if true {
        v9 = (v8 <= v8)
        fmt.Println(v7)
        for fi22 := 0; fi22 < 11; fi22++ {
            var v23 int8 = v4
            fmt.Println(v9, v4)
            _ = fi22
            _ = v23
        }
    }
    _ = len(v3)
    fmt.Println(v7, v9)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
