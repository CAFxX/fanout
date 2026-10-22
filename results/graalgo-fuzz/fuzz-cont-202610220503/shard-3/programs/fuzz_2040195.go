package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint; F2 int16 }
type GenS2 struct { F0 int32; F1 string }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint32; F1 int64; F2 int8 }
func genFn0(p0 uint64, p1 uint16) int64 {
    fmt.Println(p0, p0)
    p1 -= (p1 * p1)
    switch 127 {
    case 9223372036854775807:
        var v5 int64 = 0
        _ = v5
    case -2147483648:
        switch "hello" {
        case "hello":
            switch -9223372036854775808 {
            case 127:
                for fi6 := 0; fi6 < 6; fi6++ {
                    var v7 int = (fi6 << ((4294967295) & 63))
                    v7 = v7
                    var v8 uint64 = p0
                    _ = fi6
                    _ = v7
                    _ = v8
                }
            case -1:
                var v9 float32 = float32(1e10)
                _ = v9
            default:
                for fi10 := 0; fi10 < 15; fi10++ {
                    var v11 uint64 = p0
                    v12 := []int{fi10}
                    _ = fi10
                    _ = v11
                    _ = v12
                }
            }
        case "\n\t":
            for fi13 := 0; fi13 < 8; fi13++ {
                var v14 uint16 = 65535
                _ = fi13
                _ = v14
            }
        }
    }
    var v15 uint32 = 2147483647
    _ = p0
    _ = p1
    _ = v15
    return 0
}
func genFn1(p0 int64) int8 {
    p0 = p0
    p0 -= genFn0(4294967295, 0)
    p0 -= p0
    if (!false) {
        switch "" {
        case "hello":
            p0 += p0
        default:
            if false {
                v16 := GenS1{F0: 2147483647, F1: 775, F2: 1000}
                _ = v16
            }
        }
    } else {
        v17 := GenS1{F0: 4294967295, F1: 255, F2: 42}
        var v18 float64 = 1.5
        var v19 string = "foo bar"
        _ = v17
        _ = v18
        _ = v19
    }
    var v20 uint = 705
    _ = p0
    _ = v20
    return -1
}

func main() {
    var v1 int = 4294967295
    v2 := []int{v1}
    v3 := make(map[int]string)
    var v4 uint32 = 2147483648
    fmt.Println(v1, v1)
    v1 = v1
    var v21 int = v1
    for fi22 := 0; fi22 < 19; fi22++ {
        v4 = (v4 % ((4294967295) | 1))
        _ = fi22
    }
    v23 := GenS1{F0: v4, F1: 352, F2: -1}
    switch ("\n\t" + ("foo bar" + "\n\t")) {
    case "hello":
        v1 = (-(v21))
    default:
        for fi24 := 0; fi24 < 14; fi24++ {
            if (0 < 18446744073709551615) {
                var v25 string = ("hello" + "a")
                _ = v25
            }
            _ = fi24
        }
    }
    for fi26 := 0; fi26 < 18; fi26++ {
        if ((0 > -1) || (1 > 32767)) {
            var v27 int = (-(v21))
            v28 := GenS2{F0: 42, F1: "foo bar"}
            _ = v27
            _ = v28
        }
        _ = fi26
    }
    if len(v2) > 0 {
        v2[0] = v1
    }
    v21 = int(v4)
    switch v1 {
    case -535:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    case 9223372036854775807:
        v4 -= v4
    case 128:
        if false {
            switch "foo bar" {
            case "hello":
                fmt.Println(v21, v4, v21)
            }
            v1 -= (-(v21))
        } else {
            var v29 int64 = 9223372036854775807
            _ = len(v2)
            _ = v29
        }
    }
    switch ("a" + "") {
    case "a":
        fmt.Println(v1)
    }
    v30 := GenS3{F0: 65535}
    v31 := GenS4{F0: 608, F1: 95, F2: 42}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v23
    _ = v3
    _ = v30
    _ = v31
    _ = v4
}
