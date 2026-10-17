package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32 }
type GenS2 struct { F0 uint64; F1 bool; F2 uint32 }
type GenS3 struct { F0 string; F1 int; F2 string }

func main() {
    v1 := []int{139, 4294967295, 841}
    var v2 float64 = 0.1
    v3 := make(map[string]uint32)
    if (-128 < -29) {
        v1 = append(v1, 32767)
        fmt.Println(v2)
        v4 := GenS1{F0: true, F1: 2147483647}
        _ = v4
    } else {
        var v5 uint64 = 42
        var v6 uint32 = 311
        _ = v5
        _ = v6
    }
    var v7 int32 = -2147483648
    fmt.Println(v2, v7)
    v7 = (v7 % ((120) | 1))
    _ = len(v1)
    for fi8 := 0; fi8 < 19; fi8++ {
        var v9 int = -1
        v10 := GenS2{F0: 42, F1: (32767 < 42), F2: 2147483647}
        v9 -= fi8
        _ = fi8
        _ = v10
        _ = v9
    }
    for fi11 := 0; fi11 < 14; fi11++ {
        if false {
            delete(v3, "hello")
        } else {
            var v12 int = 853
            fmt.Println(fi11, v2)
            for fi13 := 0; fi13 < 16; fi13++ {
                var v14 uint32 = 779
                v15 := GenS3{F0: "\n\t", F1: v12, F2: "foo bar"}
                _ = fi13
                _ = v14
                _ = v15
            }
            _ = v12
        }
        _ = fi11
    }
    var v16 int32 = v7
    if ((242 != 9223372036854775807) || true) {
        v3[("a" + ("a" + ""))] = 1
    } else {
        v16 = 2147483647
    }
    switch "hello" {
    case "hello":
        var v17 uint16 = 804
        _ = v17
    case "a":
        v2 = v2
    default:
        v7 += 871
    }
    var v18 int32 = (v7 - (-(v16)))
    v2 = v2
    v2 -= 1.5
    var v19 float64 = ((float64(0.0) / (v2 + v2)) * (0.1 / v2))
    var v20 uint64 = 18446744073709551615
    for fi21 := 0; fi21 < 3; fi21++ {
        switch 9223372036854775807 {
        case 9223372036854775807:
            fmt.Println(fi21, v16)
        case 42:
            _ = len(v3)
        case 186:
            v7 += (v7 ^ v7)
        }
        v3[("a" + "foo bar")] = 2147483648
        _ = fi21
    }
    if (false || false) {
        switch (("foo bar" + ("foo bar" + "foo bar")) + "a") {
        case "a":
            for fi22 := 0; fi22 < 7; fi22++ {
                var v23 int16 = 206
                var v24 uint = 651
                _ = fi22
                _ = v23
                _ = v24
            }
        case "foo bar":
            if ((0.0 - 0.0) >= -2.25) {
                v20 -= (v20 + 4294967295)
                fmt.Println(v2, v20, v16)
            }
        default:
            var v25 int = 42
            _ = v25
        }
    } else {
        fmt.Println(v19)
        fmt.Println(v2, v20)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v16
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v7
}
