package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64 }
type GenS2 struct { F0 bool; F1 uint64 }
type GenS3 struct { F0 int64; F1 int }
func genFn0(p0 uint64, p1 int8, p2 int16) uint8 {
    switch "a" {
    case "\n\t":
        v6 := []int{-31, 42, 9223372036854775807, 4294967295}
        _ = v6
    case "":
        v7 := GenS1{F0: 42, F1: -1}
        _ = v7
    default:
        p1 -= (p1 ^ (-(p1)))
    }
    switch (("foo bar" + "a") + "foo bar") {
    case "a":
        var v8 float64 = -2.25
        _ = v8
    case "\n\t":
        v9 := []int{}
        _ = v9
    case "hello":
        v10 := []int{1, 47, 2147483647, -2147483648}
        _ = v10
    }
    switch 346 {
    case -2147483648:
        v11 := []int{-9223372036854775808, -9223372036854775808, -9223372036854775808, -895}
        _ = v11
    }
    v12 := GenS2{F0: true, F1: (p0 - p0)}
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return 127
}

func main() {
    v1 := []string{(("hello" + "a") + ("\n\t" + (("\n\t" + "foo bar") + "foo bar"))), "", ("\n\t" + ""), ("" + "\n\t")}
    v2 := make(map[int]uint32)
    v3 := []int{128, 4294967295, 401, 127}
    var v4 int8 = -1
    var v5 uint16 = 65535
    var v13 uint64 = 575
    v14 := GenS1{F0: v13, F1: 945}
    for fi15 := 0; fi15 < 6; fi15++ {
        var v16 uint8 = 2
        _ = fi15
        _ = v16
    }
    if true {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        var v17 uint8 = uint8(v13)
        _ = v17
    }
    if false {
        var v18 string = "a"
        _ = v18
    } else {
        if (!true) {
            var v19 uint64 = v13
            switch "hello" {
            case "a":
                if (42 <= 255) {
                    switch 424 {
                    case 32767:
                        _ = len(v2)
                    }
                }
            case "hello":
                fmt.Println(v19, v19, v13)
            default:
                var v20 uint32 = 109
                _ = v20
            }
            _ = v19
        }
        if (-2.25 != 1.5) {
            v5 += v5
        }
    }
    for fi21 := 0; fi21 < 5; fi21++ {
        v22 := GenS3{F0: 292, F1: 1}
        v2[(fi21 + 4294967295)] = 754
        _ = fi21
        _ = v22
    }
    var v23 float64 = 1.5
    if len(v3) > 0 {
        v3[0] = -871
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
