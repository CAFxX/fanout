package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64; F2 uint }
type GenS2 struct { F0 string; F1 uint32 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint16; F1 uint32 }
func genFn0(p0 int8, p1 int32) int16 {
    var v4 uint64 = 2
    for fi5 := 0; fi5 < 7; fi5++ {
        p0 = p0
        _ = fi5
    }
    var v6 string = "foo bar"
    _ = p0
    _ = p1
    _ = v4
    _ = v6
    return -464
}
func genFn1(p0 uint16, p1 int8) uint16 {
    var v7 bool = (1.5 >= (float32(1e10) + -2.25))
    if (((^(p0)) != p0) && (genFn0((p1 / ((1) | 1)), 936) != genFn0((p1 % (((p1 | p1)) | 1)), -1))) {
        var v8 uint64 = 4294967295
        _ = v8
    }
    fmt.Println(p0)
    switch "\n\t" {
    case "foo bar":
        if (((-745 <= -553) && (1e10 != -2.25)) && ((true && v7) || true)) {
            switch -133 {
            case 1:
                v9 := []int{-1, -162}
                _ = v9
            case -848:
                v10 := []int{2147483647, 127}
                _ = v10
            case -377:
                v11 := []int{-9223372036854775808, 42, 128, -515}
                _ = v11
            }
        }
    case "\n\t":
        fmt.Println(p0, p1, v7)
    default:
        v12 := GenS1{F0: (v7 && true), F1: 1, F2: 255}
        _ = v12
    }
    var v13 uint = 978
    _ = p0
    _ = p1
    _ = v13
    _ = v7
    return (^(p0))
}

func main() {
    var v1 uint64 = 18446744073709551615
    v2 := []string{}
    v3 := make(map[uint32]int)
    v1 -= (v1 & (v1 ^ (v1 | v1)))
    if (true || (true || false)) {
        if true {
            for fi14 := 0; fi14 < 1; fi14++ {
                var v15 bool = (0 > -9223372036854775808)
                v1 -= (v1 << ((670) & 63))
                var v16 string = "a"
                _ = fi14
                _ = v15
                _ = v16
            }
            fmt.Println(v1)
            for fi17 := 0; fi17 < 7; fi17++ {
                v1 = 129
                v1 -= v1
                fmt.Println(fi17)
                _ = fi17
            }
        }
        _ = len(v3)
        var v18 int16 = 1000
        _ = v18
    }
    v1 -= (v1 >> ((65535) & 63))
    v2 = append(v2, ("\n\t" + (("a" + ("foo bar" + "hello")) + ((((("hello" + ("a" + ("foo bar" + "hello"))) + "hello") + "foo bar") + "\n\t") + ""))))
    v3[248] = 0
    v1 += 42
    var v19 float32 = (0.0 + 0.1)
    v20 := GenS2{F0: "\n\t", F1: uint32(v1)}
    for fi21 := 0; fi21 < 16; fi21++ {
        if (42 == 127) {
            switch (fi21 >> ((18446744073709551615) & 63)) {
            case -9223372036854775808:
                v22 := GenS3{F0: -1000}
                _ = v22
            case 32767:
                switch (fi21 + (fi21 >> ((0) & 63))) {
                case -278:
                    v23 := GenS4{F0: ("hello" + "a")}
                    _ = v23
                case 113:
                    if (!true) {
                        var v24 int64 = -9223372036854775808
                        fmt.Println(fi21, v24)
                        _ = v24
                    } else {
                        v1 = (v1 - 436)
                        v19 = (v19 - (v19 * ((float32(0.1) + (1.5 * v19)) * 3.14159)))
                        fmt.Println(v1, v19, v19)
                    }
                case 4294967295:
                    for fi25 := 0; fi25 < 1; fi25++ {
                        var v26 uint32 = 1
                        if false {
                            fmt.Println(v26, v19)
                        }
                        v1 -= v1
                        _ = fi25
                        _ = v26
                    }
                default:
                    for fi27 := 0; fi27 < 4; fi27++ {
                        v28 := GenS5{F0: 679, F1: 297}
                        if len(v2) > 0 {
                            v2[0] = ""
                        }
                        v19 = ((v19 - 0.1) * 3.14159)
                        _ = fi27
                        _ = v28
                    }
                }
            }
            for fi29 := 0; fi29 < 9; fi29++ {
                fmt.Println(fi29)
                var v30 uint32 = 398
                v30 -= 711
                _ = fi29
                _ = v30
            }
        } else {
            v19 += -2.25
            var v31 uint32 = 2147483647
            v32 := GenS5{F0: 2, F1: 1}
            _ = v31
            _ = v32
        }
        _ = fi21
    }
    var v33 uint8 = 42
    v33 -= (v33 + v33)
    v34 := v3[1]
    v33 = 127
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v33
    _ = v34
}
