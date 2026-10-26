package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8 }
type GenS2 struct { F0 string; F1 uint8; F2 bool }
type GenS3 struct { F0 bool; F1 string }
type GenS4 struct { F0 int16; F1 uint64; F2 uint }
func genFn0(p0 uint32) string {
    var v7 int64 = 562
    p0 = (p0 * (p0 - (p0 / ((4294967295) | 1))))
    v8 := GenS1{F0: 42, F1: -1}
    _ = p0
    _ = v7
    _ = v8
    return ("\n\t" + "a")
}

func main() {
    v1 := []uint32{}
    v2 := make(map[uint32]bool)
    v3 := []string{}
    var v4 uint = 608
    var v5 int32 = 503
    var v6 uint8 = 128
    switch "\n\t" {
    case "hello":
        if true {
            var v9 uint16 = 0
            var v10 int16 = -1
            if true {
                v11 := GenS2{F0: ("" + (("hello" + "hello") + "")), F1: (v6 << ((v4) & 63)), F2: true}
                _ = v11
            } else {
                v12 := GenS2{F0: "a", F1: v6, F2: false}
                _ = v12
            }
            _ = v10
            _ = v9
        } else {
            var v13 int16 = 32767
            _ = v13
        }
    case "\n\t":
        v5 += v5
    default:
        if (255 == 32767) {
            for fi14 := 0; fi14 < 18; fi14++ {
                var v15 int32 = v5
                if false {
                    v16 := GenS3{F0: (0.1 == (((0.0 / 3.14159) - 0.0) + -2.25)), F1: ("foo bar" + (("a" + "foo bar") + "hello"))}
                    for fi17 := 0; fi17 < 16; fi17++ {
                        fmt.Println(v4, v4, fi14)
                        if (65535 >= 42) {
                            var v18 int16 = 866
                            v18 -= v18
                            var v19 uint64 = 26
                            _ = v18
                            _ = v19
                        }
                        var v20 int8 = 0
                        _ = fi17
                        _ = v20
                    }
                    var v21 uint8 = (v6 | 1)
                    _ = v16
                    _ = v21
                } else {
                    for fi22 := 0; fi22 < 12; fi22++ {
                        if len(v3) > 0 {
                            v3[0] = ("hello" + "")
                        }
                        _ = fi22
                    }
                }
                var v23 bool = true
                _ = fi14
                _ = v15
                _ = v23
            }
            v3 = append(v3, "a")
        }
    }
    switch ("foo bar" + "hello") {
    case "hello":
        var v24 int = -233
        _ = v24
    case "foo bar":
        v25 := GenS4{F0: 42, F1: 699, F2: v4}
        _ = v25
    case "a":
        fmt.Println(v4)
    }
    v26 := GenS1{F0: 65535, F1: -128}
    v4 = (v4 << (((v4 % ((v4) | 1))) & 63))
    var v27 uint64 = 4294967295
    fmt.Println(v27)
    var v28 float64 = -2.25
    if false {
        v29 := GenS4{F0: 1000, F1: v27, F2: 1}
        fmt.Println(v5)
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = v29
    }
    v1 = append(v1, 2147483647)
    v30 := v2[1]
    for fi31 := 0; fi31 < 13; fi31++ {
        v2[2] = v30
        _ = fi31
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
