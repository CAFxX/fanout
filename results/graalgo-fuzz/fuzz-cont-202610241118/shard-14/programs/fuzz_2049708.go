package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool; F2 uint32 }
type GenS2 struct { F0 int16; F1 string; F2 uint64 }
type GenS3 struct { F0 uint }
func genFn0() int {
    var v7 int8 = -84
    fmt.Println(v7)
    fmt.Println(v7)
    if true {
        fmt.Println(v7)
        for fi8 := 0; fi8 < 10; fi8++ {
            if false {
                var v9 string = ("\n\t" + (("" + "hello") + ""))
                _ = v9
            } else {
                v10 := []int{9223372036854775807, fi8, fi8}
                v11 := GenS1{F0: 1, F1: false, F2: 1}
                _ = v10
                _ = v11
            }
            switch "a" {
            case "\n\t":
                var v12 int16 = -1
                _ = v12
            }
            fmt.Println(v7)
            _ = fi8
        }
    } else {
        v7 = v7
    }
    _ = v7
    return 986
}
func genFn1(p0 int8, p1 int8, p2 string) int16 {
    var v13 bool = false
    var v14 float64 = 1.5
    var v15 string = "hello"
    var v16 int64 = 489
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    return -32768
}

func main() {
    var v1 uint = 42
    v2 := make(map[int]string)
    var v3 uint = v1
    var v4 int16 = -71
    v5 := []string{("\n\t" + ("a" + "foo bar")), "hello", (("" + "foo bar") + "foo bar")}
    var v6 int64 = 9223372036854775807
    var v17 uint8 = 128
    v17 = (v17 - v17)
    v1 += 1
    fmt.Println(v6, v6)
    var v18 string = ("\n\t" + ("a" + "foo bar"))
    switch "foo bar" {
    case "\n\t":
        var v19 uint64 = 0
        _ = v19
    case "a":
        for fi20 := 0; fi20 < 9; fi20++ {
            v21 := GenS1{F0: v6, F1: (!false), F2: 2147483647}
            v1 -= (v3 % ((2) | 1))
            _ = fi20
            _ = v21
        }
    default:
        for fi22 := 0; fi22 < 3; fi22++ {
            v23 := GenS1{F0: 9223372036854775807, F1: true, F2: 2}
            switch v18 {
            case "foo bar":
                v17 = (^(v17))
            case "\n\t":
                v4 += genFn1(int8(v1), -1, (v18 + v18))
            case "a":
                var v24 int32 = -2147483648
                _ = v24
            }
            v25 := GenS2{F0: v4, F1: "a", F2: 42}
            _ = fi22
            _ = v23
            _ = v25
        }
    }
    v26 := GenS1{F0: v6, F1: (-1 <= 60), F2: 254}
    v4 += v4
    if (float32((0.1 * float64(1e10))) < (-2.25 / 1e10)) {
        v27 := GenS1{F0: v6, F1: (127 >= -1), F2: 755}
        _ = v27
    } else {
        v28 := GenS3{F0: (v3 << ((42) & 63))}
        _ = v28
    }
    var v29 int = 617
    if false {
        v6 = (v6 >> (((v3 ^ v1)) & 63))
    } else {
        var v30 uint64 = 4294967295
        _ = v30
    }
    var v31 int64 = (-(v6))
    v32 := GenS3{F0: v3}
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v26
    _ = v29
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
