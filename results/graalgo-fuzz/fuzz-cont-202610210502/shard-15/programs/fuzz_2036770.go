package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint; F1 uint64; F2 int16 }
func genFn0(p0 int, p1 int64, p2 uint32) int8 {
    p0 = (-(p0))
    if false {
        var v7 uint32 = (p2 / ((p2) | 1))
        var v8 uint16 = 255
        _ = v7
        _ = v8
    } else {
        p2 = p2
    }
    var v9 int8 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return v9
}
func genFn1(p0 bool, p1 uint32) int64 {
    var v10 float32 = ((-2.25 * 1.5) * 1.5)
    var v11 int8 = genFn0(-9223372036854775808, -268, 65535)
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    return -9223372036854775808
}

func main() {
    var v1 bool = (258 <= -1)
    var v2 int32 = -832
    v3 := []uint32{}
    var v4 int8 = -128
    var v5 int32 = (v2 << ((4294967295) & 63))
    v6 := make(map[int]bool)
    v12 := GenS1{F0: 65535}
    _ = len(v3)
    fmt.Println(v5)
    v13 := GenS2{F0: 787, F1: 0, F2: -130}
    switch -1 {
    case 2147483647:
        for fi14 := 0; fi14 < 11; fi14++ {
            switch "a" {
            case "foo bar":
                fmt.Println(v5, v4, fi14)
            case "hello":
                var v15 bool = (0 >= 890)
                _ = v15
            case "\n\t":
                v1 = v1
            default:
                var v16 string = ("a" + "foo bar")
                _ = v16
            }
            fmt.Println(v5, v2)
            v2 = (v2 % (((-(v2))) | 1))
            _ = fi14
        }
    }
    v5 = 42
    switch ("hello" + (("hello" + (("hello" + "a") + "hello")) + ("\n\t" + "foo bar"))) {
    case "foo bar":
        v17 := GenS1{F0: 65535}
        _ = v17
    case "\n\t":
        var v18 int = 665
        _ = v18
    }
    v1 = true
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v1 = v1
    if (1 != 1) {
        v3 = append(v3, 2)
        fmt.Println(v2, v5, v2)
    }
    if false {
        var v19 bool = false
        v2 += v2
        switch 128 {
        case -911:
            var v20 int64 = 112
            _ = v20
        case 9223372036854775807:
            fmt.Println(v2, v1, v5)
        }
        _ = v19
    } else {
        v6[326] = (!(9223372036854775807 == genFn1(v1, 2147483648)))
        v21 := GenS1{F0: 723}
        switch -2147483648 {
        case 525:
            v6[127] = true
        case 4294967295:
            for fi22 := 0; fi22 < 16; fi22++ {
                var v23 string = "\n\t"
                _ = fi22
                _ = v23
            }
        default:
            var v24 uint32 = 2147483647
            _ = v24
        }
        _ = v21
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
