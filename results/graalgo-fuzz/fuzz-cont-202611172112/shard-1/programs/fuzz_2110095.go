package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int16; F1 int }
type GenS5 struct { F0 string }

func main() {
    v1 := make(map[int]bool)
    var v2 float64 = 3.14159
    var v3 uint64 = 1
    var v4 int64 = 204
    v5 := []int{42, -2147483648}
    v6 := GenS1{F0: -730}
    switch 128 {
    case 128:
        v3 -= (^(v3))
    }
    v3 += 580
    var v7 uint = 65535
    for fi8 := 0; fi8 < 17; fi8++ {
        fmt.Println(v4, v3)
        if false {
            if (0 == 32767) {
                var v9 int8 = 75
                for fi10 := 0; fi10 < 6; fi10++ {
                    _ = len(v1)
                    switch ("" + "a") {
                    case "hello":
                        v11 := GenS2{F0: (v3 << ((v7) & 63))}
                        _ = v11
                    case "foo bar":
                        v7 += uint(v3)
                    case "\n\t":
                        var v12 float32 = (-2.25 + (3.14159 * (0.1 - 3.14159)))
                        _ = v12
                    default:
                        delete(v1, int(v7))
                        v13 := GenS3{F0: (v4 | v4)}
                        if (v4 > v4) {
                            v3 -= v3
                        }
                        _ = v13
                    }
                    _ = fi10
                }
                _ = v9
            }
            for fi14 := 0; fi14 < 14; fi14++ {
                v3 = 0
                fmt.Println(fi8, fi8, fi14)
                _ = fi14
            }
            fmt.Println(v7)
        } else {
            var v15 bool = true
            _ = v15
        }
        _ = fi8
    }
    if false {
        fmt.Println(v4, v7, v7)
        v7 = v7
    }
    v16 := GenS4{F0: -32768, F1: -9223372036854775808}
    v17 := GenS2{F0: 2}
    if true {
        var v18 int16 = -238
        switch ((("hello" + "foo bar") + (("a" + "") + ("foo bar" + ("" + "\n\t")))) + "foo bar") {
        case "":
            v18 = v18
        case "a":
            var v19 uint16 = 991
            _ = v19
        }
        _ = v18
    }
    var v20 uint8 = 53
    for fi21 := 0; fi21 < 16; fi21++ {
        _ = len(v1)
        if (v2 > 0.0) {
            v4 += v4
            v20 -= 97
            for fi22 := 0; fi22 < 6; fi22++ {
                var v23 uint = (v7 % ((1) | 1))
                v23 += (v7 & v23)
                _ = fi22
                _ = v23
            }
        }
        if false {
            v24 := GenS3{F0: v4}
            _ = v24
        }
        _ = fi21
    }
    fmt.Println(v3, v4)
    v25 := GenS5{F0: "\n\t"}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
