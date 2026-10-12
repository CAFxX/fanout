package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string }
type GenS2 struct { F0 int64; F1 uint16 }
type GenS3 struct { F0 int32 }
func genFn0(p0 uint16, p1 int32) uint {
    fmt.Println(p0)
    v4 := []int{}
    _ = p0
    _ = p1
    _ = v4
    return 0
}
func genFn1(p0 uint64, p1 int16, p2 uint) int32 {
    p1 = p1
    if ((^(p2)) >= (^(p2))) {
        if ((!true) || (2147483647 == -1)) {
            p1 = p1
        }
    }
    switch -444 {
    case 744:
        p0 += (p0 ^ (p0 | 210))
    case 33:
        var v5 float64 = (float64(-2.25) * 1e10)
        _ = v5
    default:
        switch 42 {
        case 4294967295:
            v6 := []int{-465, 0, -2147483648, -2147483648}
            _ = v6
        }
    }
    for fi7 := 0; fi7 < 12; fi7++ {
        v8 := GenS1{F0: 42, F1: "\n\t"}
        _ = fi7
        _ = v8
    }
    fmt.Println(p1, p1, p0)
    _ = p0
    _ = p1
    _ = p2
    return -940
}

func main() {
    var v1 uint = 517
    v2 := make(map[int]bool)
    var v3 float32 = 0.0
    if (0 < 4294967295) {
        if ((0 > 18446744073709551615) && true) {
            v1 = (^(v1))
            v1 -= (^(v1))
        }
        var v9 float32 = v3
        v10 := v2[9223372036854775807]
        _ = v10
        _ = v9
    }
    v11 := GenS1{F0: 1, F1: ("\n\t" + "")}
    v12 := []int{-2147483648, 32767}
    if (37 > 42) {
        var v13 uint64 = 32
        for fi14 := 0; fi14 < 16; fi14++ {
            var v15 int64 = -471
            _ = fi14
            _ = v15
        }
        _ = v13
    }
    var v16 uint32 = 71
    for fi17 := 0; fi17 < 11; fi17++ {
        var v18 uint = v1
        _ = fi17
        _ = v18
    }
    if (!true) {
        v19 := GenS1{F0: 127, F1: "foo bar"}
        switch "a" {
        case "hello":
            v3 += v3
        default:
            v20 := GenS1{F0: 2, F1: (("" + "\n\t") + "")}
            _ = v20
        }
        v3 -= v3
        _ = v19
    } else {
        var v21 int8 = 0
        _ = v21
    }
    var v22 int = 32767
    switch "a" {
    case "\n\t":
        if true {
            v23 := GenS2{F0: 1, F1: 2}
            v3 = -2.25
            _ = v23
        }
    case "":
        v24 := GenS3{F0: 172}
        _ = v24
    case "a":
        var v25 float64 = 0.0
        _ = v25
    default:
        for fi26 := 0; fi26 < 19; fi26++ {
            fmt.Println(v22)
            for fi27 := 0; fi27 < 15; fi27++ {
                switch (("\n\t" + "\n\t") + ("\n\t" + ("\n\t" + ""))) {
                case "a":
                    v16 = (v16 - v16)
                    if true {
                        if (false || (0 >= -32768)) {
                            if false {
                                v28 := GenS2{F0: -334, F1: 65535}
                                _ = v28
                            } else {
                                var v29 int8 = -128
                                _ = v29
                            }
                            switch ("hello" + ("hello" + (((("" + "hello") + "hello") + "") + ("\n\t" + ("hello" + "a"))))) {
                            case "foo bar":
                                var v30 int = (-(fi27))
                                _ = v30
                            default:
                                v16 += v16
                            }
                            v22 = v22
                        } else {
                            if false {
                                v31 := GenS1{F0: 93, F1: "foo bar"}
                                var v32 uint16 = 342
                                _ = v31
                                _ = v32
                            } else {
                                fmt.Println(v16, v3, v3)
                            }
                            switch "\n\t" {
                            case "\n\t":
                                v3 -= -2.25
                            case "foo bar":
                                v3 -= (v3 / float32(0.1))
                            case "hello":
                                v16 += v16
                            default:
                                v33 := GenS1{F0: uint8(v16), F1: (("" + (("a" + ("hello" + (("\n\t" + "\n\t") + (("" + (("\n\t" + ((("a" + "hello") + (("foo bar" + ("" + "hello")) + "")) + (("hello" + "") + "foo bar"))) + "a")) + ("" + ("\n\t" + ("a" + "a"))))))) + "")) + "hello")}
                                _ = v33
                            }
                            var v34 int16 = 1000
                            _ = v34
                        }
                    }
                }
                _ = fi27
            }
            _ = fi26
        }
    }
    for fi35 := 0; fi35 < 12; fi35++ {
        delete(v2, fi35)
        v36 := GenS2{F0: 156, F1: 65535}
        _ = fi35
        _ = v36
    }
    var v37 uint8 = 255
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    _ = v1
    _ = v11
    _ = v12
    _ = v16
    _ = v2
    _ = v22
    _ = v3
    _ = v37
}
