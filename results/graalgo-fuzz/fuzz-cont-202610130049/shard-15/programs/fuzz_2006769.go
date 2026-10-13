package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16 }
type GenS2 struct { F0 int64; F1 string; F2 uint }
func genFn0(p0 int) uint32 {
    p0 = p0
    var v4 int32 = -193
    p0 = (p0 >> ((975) & 63))
    _ = p0
    _ = v4
    return 0
}
func genFn1() int16 {
    v5 := []int{1, -821, 42, -221}
    v5 = append(v5, 31)
    switch "a" {
    case "foo bar":
        var v6 float32 = float32(1e10)
        _ = v6
    }
    switch (("foo bar" + "\n\t") + "a") {
    case "":
        for fi7 := 0; fi7 < 3; fi7++ {
            var v8 int32 = 0
            fmt.Println(fi7, fi7)
            _ = fi7
            _ = v8
        }
    case "\n\t":
        var v9 float64 = 1.5
        _ = v9
    case "a":
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    }
    var v10 uint32 = 1
    _ = v10
    _ = v5
    return 49
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[int]string)
    var v3 float32 = 3.14159
    v3 = 0.1
    fmt.Println(v3)
    for fi11 := 0; fi11 < 2; fi11++ {
        var v12 uint = 65535
        _ = fi11
        _ = v12
    }
    if (874 == 2) {
        fmt.Println(v3)
    } else {
        v13 := GenS1{F0: 0, F1: 540}
        var v14 bool = false
        _ = v13
        _ = v14
    }
    var v15 int64 = -9223372036854775808
    var v16 uint64 = 42
    v15 -= v15
    v17 := []int{2147483647, 42}
    var v18 uint16 = 0
    switch 42 {
    case -9223372036854775808:
        var v19 uint = 42
        _ = v19
    default:
        for fi20 := 0; fi20 < 5; fi20++ {
            v15 -= (v15 % ((v15) | 1))
            switch "a" {
            case "":
                delete(v2, (fi20 - 510))
            case "\n\t":
                if false {
                    if ((v18 * (v18 | v18)) < 0) {
                        v21 := GenS2{F0: v15, F1: "hello", F2: 0}
                        var v22 int16 = -743
                        _ = v21
                        _ = v22
                    } else {
                        switch "a" {
                        case "\n\t":
                            v18 += 42
                        }
                        v23 := v1[fi20]
                        fmt.Println(v23, v16)
                        _ = v23
                    }
                }
            case "hello":
                v15 = -9223372036854775808
            default:
                v24 := GenS1{F0: -2147483648, F1: 1000}
                _ = v24
            }
            _ = fi20
        }
    }
    var v25 uint64 = v16
    for fi26 := 0; fi26 < 20; fi26++ {
        var v27 int16 = 32767
        var v28 int64 = (-(v15))
        v27 = v27
        _ = fi26
        _ = v27
        _ = v28
    }
    v29 := GenS1{F0: 1, F1: 42}
    switch "\n\t" {
    case "hello":
        switch ("a" + "a") {
        case "a":
            delete(v1, 9223372036854775807)
        }
    case "a":
        v30 := GenS1{F0: -485, F1: 318}
        _ = v30
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v25
    _ = v29
    _ = v3
}
