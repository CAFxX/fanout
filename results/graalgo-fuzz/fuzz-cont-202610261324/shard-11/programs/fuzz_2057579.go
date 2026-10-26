package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int; F1 int64; F2 string }
type GenS4 struct { F0 uint32; F1 bool; F2 string }
type GenS5 struct { F0 uint16; F1 int8 }
func genFn0(p0 int64, p1 uint32, p2 int64) int {
    if (((("foo bar" + "foo bar") + "") + (((("" + "foo bar") + "a") + (("" + "foo bar") + ("a" + "a"))) + "")) > "\n\t") {
        v4 := GenS1{F0: 42}
        _ = v4
    } else {
        p2 += (p0 - 976)
    }
    p1 = (p1 & 42)
    var v5 int32 = -1
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 383
}

func main() {
    var v1 float64 = (3.14159 - -2.25)
    v2 := make(map[string]bool)
    var v3 uint32 = 65535
    v6 := GenS1{F0: 255}
    var v7 int = -816
    for fi8 := 0; fi8 < 3; fi8++ {
        v9 := []int{(fi8 - genFn0(9223372036854775807, (v3 << ((0) & 63)), -9223372036854775808))}
        _ = fi8
        _ = v9
    }
    v1 = (3.14159 - 1.5)
    var v10 int32 = -204
    var v11 uint32 = 42
    switch "\n\t" {
    case "\n\t":
        var v12 uint = 255
        _ = v12
    default:
        var v13 int32 = v10
        _ = v13
    }
    if (false && false) {
        var v14 uint = 18446744073709551615
        if true {
            if (true && (406 >= 1)) {
                for fi15 := 0; fi15 < 2; fi15++ {
                    v7 = (v7 * 127)
                    switch "hello" {
                    case "foo bar":
                        v10 -= (v10 | v10)
                    case "a":
                        v16 := []int{-1}
                        _ = v16
                    case "hello":
                        fmt.Println(v7, v1)
                    default:
                        fmt.Println(v3, v3, v11)
                    }
                    _ = fi15
                }
            } else {
                var v17 uint = v14
                for fi18 := 0; fi18 < 8; fi18++ {
                    v19 := GenS2{F0: ((42 > 32767) && (uint16(v17) >= 0))}
                    _ = fi18
                    _ = v19
                }
                _ = v17
            }
            v2[(("\n\t" + "a") + ("a" + ""))] = false
        }
        var v20 bool = (42 > 240)
        _ = v14
        _ = v20
    } else {
        fmt.Println(v11)
        v21 := GenS3{F0: v7, F1: 1, F2: "hello"}
        _ = v21
    }
    for fi22 := 0; fi22 < 4; fi22++ {
        for fi23 := 0; fi23 < 9; fi23++ {
            if true {
                _ = len(v2)
                var v24 int8 = 127
                _ = v24
            }
            _ = fi23
        }
        v25 := GenS4{F0: (v3 - v11), F1: (v10 != v10), F2: "foo bar"}
        v26 := GenS3{F0: int(v10), F1: -1, F2: (("hello" + ("\n\t" + ("foo bar" + "a"))) + "hello")}
        _ = fi22
        _ = v25
        _ = v26
    }
    switch "hello" {
    case "hello":
        v11 = (v11 - (^(v3)))
    case "a":
        var v27 float64 = ((1e10 + 0.0) + (v1 + v1))
        _ = v27
    default:
        if false {
            v28 := GenS5{F0: 192, F1: 0}
            v29 := GenS5{F0: 160, F1: 42}
            _ = v28
            _ = v29
        }
    }
    v11 -= v11
    v30 := GenS2{F0: true}
    var v31 string = ""
    v11 -= (v11 | v3)
    v1 += 0.0
    if (!(1000 <= 32767)) {
        var v32 int32 = (v10 - (v10 & (-(v10))))
        if (v3 > (v11 + (v3 - 2))) {
            var v33 int64 = 185
            v10 -= v32
            v34 := []int{32768, v7}
            _ = v33
            _ = v34
        }
        _ = v32
    }
    for fi35 := 0; fi35 < 4; fi35++ {
        v36 := GenS2{F0: ((1.5 - -2.25) <= float32(v1))}
        _ = fi35
        _ = v36
    }
    var v37 float64 = (v1 - (0.0 * v1))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v37:", v37)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v37
    _ = v6
    _ = v7
}
