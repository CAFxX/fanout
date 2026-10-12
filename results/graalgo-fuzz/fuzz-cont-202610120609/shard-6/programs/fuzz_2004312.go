package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16 }
type GenS2 struct { F0 int; F1 int64; F2 int8 }
type GenS3 struct { F0 uint; F1 uint16; F2 int16 }
type GenS4 struct { F0 uint; F1 string }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 int16; F1 uint8 }
func genFn0(p0 uint, p1 uint8) string {
    var v5 int32 = 2147483647
    if (-1 <= -1) {
        var v6 uint64 = 419
        _ = v6
    }
    for fi7 := 0; fi7 < 7; fi7++ {
        p1 += 255
        if false {
            var v8 int64 = 42
            v5 = (^(v5))
            _ = v8
        }
        for fi9 := 0; fi9 < 2; fi9++ {
            p0 -= (p0 >> (((p0 >> ((p0) & 63))) & 63))
            _ = fi9
        }
        _ = fi7
    }
    p1 = 0
    _ = p0
    _ = p1
    _ = v5
    return "hello"
}
func genFn1(p0 int64) uint8 {
    v10 := []int{127, 128}
    switch 463 {
    case -352:
        p0 -= p0
    default:
        if false {
            switch "a" {
            case "a":
                if true {
                    if len(v10) >= 2 {
                        v10 = v10[1:len(v10)]
                    }
                    p0 += 1
                    if len(v10) > 0 {
                        v10[0] = 42
                    }
                } else {
                    if len(v10) >= 2 {
                        v10 = v10[1:len(v10)]
                    }
                    p0 = (p0 - p0)
                }
            case "hello":
                v11 := GenS1{F0: 1, F1: 1}
                _ = v11
            default:
                for fi12 := 0; fi12 < 15; fi12++ {
                    for fi13 := 0; fi13 < 11; fi13++ {
                        if (!true) {
                            v10 = append(v10, (^(fi12)))
                        } else {
                            p0 += (p0 + p0)
                            var v14 int32 = 0
                            fmt.Println(p0, fi12, p0)
                            _ = v14
                        }
                        var v15 bool = true
                        v15 = false
                        v16 := GenS1{F0: 705, F1: -301}
                        _ = fi13
                        _ = v15
                        _ = v16
                    }
                    if len(v10) > 0 {
                        v10[0] = (^(fi12))
                    }
                    for fi17 := 0; fi17 < 14; fi17++ {
                        if true {
                            if true {
                                if true {
                                    var v18 bool = true
                                    var v19 uint = 0
                                    _ = len(v10)
                                    if (127 <= 42) {
                                        v19 -= 505
                                        _ = len(v10)
                                        if (0.1 <= 1e10) {
                                            fmt.Println(v19, v19, fi12)
                                        } else {
                                            if v18 {
                                                var v20 string = ("hello" + "a")
                                                _ = v20
                                            }
                                            var v21 float32 = float32(float64(((0.1 / ((1e10 * (float32((3.14159 / (1e10 * (-2.25 * 1e10)))) - 0.1)) + float32(-2.25))) - 1.5)))
                                            v22 := GenS2{F0: fi17, F1: 1, F2: 127}
                                            _ = v21
                                            _ = v22
                                        }
                                    }
                                    if v18 {
                                        var v23 uint16 = 32767
                                        _ = v23
                                    }
                                    _ = v18
                                    _ = v19
                                }
                            }
                        }
                        var v24 bool = (true && true)
                        _ = fi17
                        _ = v24
                    }
                    _ = fi12
                }
            }
            var v25 uint = 983
            v26 := GenS3{F0: v25, F1: 65535, F2: 42}
            _ = v25
            _ = v26
        } else {
            fmt.Println(p0)
            switch (((("hello" + "") + "hello") + "\n\t") + "") {
            case "":
                var v27 float64 = (float64((float32(1.5) * 0.0)) / 0.1)
                _ = v27
            case "a":
                var v28 int64 = p0
                _ = v28
            case "hello":
                if (false || true) {
                    switch 4294967295 {
                    case 4294967295:
                        v10 = append(v10, 158)
                    case 144:
                        p0 -= 132
                    case 2147483647:
                        for fi29 := 0; fi29 < 13; fi29++ {
                            switch "hello" {
                            case "foo bar":
                                p0 += p0
                            case "a":
                                var v30 int32 = 321
                                _ = v30
                            case "hello":
                                var v31 uint16 = 255
                                _ = v31
                            }
                            var v32 uint16 = 2
                            if (42 != -432) {
                                v33 := GenS4{F0: 0, F1: "\n\t"}
                                fmt.Println(p0, fi29)
                                _ = v33
                            } else {
                                v34 := GenS3{F0: 255, F1: (v32 >> ((2) & 63)), F2: -186}
                                p0 = p0
                                _ = v34
                            }
                            _ = fi29
                            _ = v32
                        }
                    }
                    var v35 uint16 = 0
                    var v36 bool = false
                    _ = v35
                    _ = v36
                }
            default:
                if (("" + ("hello" + "a")) != "\n\t") {
                    v37 := GenS5{F0: (!true)}
                    if (18446744073709551615 < 33) {
                        switch "" {
                        case "a":
                            var v38 uint32 = 0
                            _ = v38
                        default:
                            var v39 uint = 255
                            _ = v39
                        }
                        var v40 int64 = p0
                        _ = v40
                    }
                    p0 = (p0 ^ (^(p0)))
                    _ = v37
                } else {
                    p0 = p0
                    v41 := GenS5{F0: true}
                    var v42 bool = (!true)
                    _ = v41
                    _ = v42
                }
            }
        }
    }
    v43 := GenS1{F0: 2, F1: 1000}
    switch ("foo bar" + ("a" + "foo bar")) {
    case "hello":
        p0 = 42
    case "a":
        v44 := GenS4{F0: 1, F1: ("foo bar" + "a")}
        _ = v44
    }
    _ = p0
    _ = v10
    _ = v43
    return 255
}

func main() {
    var v1 uint8 = 67
    var v2 int8 = 1
    var v3 int = 4294967295
    var v4 int16 = 80
    v45 := GenS1{F0: 0, F1: (v4 - (v4 << ((95) & 63)))}
    var v46 uint32 = 65535
    v47 := GenS5{F0: (int64(v3) == 9223372036854775807)}
    var v48 string = "\n\t"
    fmt.Println(v1, v48)
    fmt.Println(v4, v46, v48)
    if false {
        v49 := GenS1{F0: 810, F1: v4}
        v50 := GenS6{F0: v4, F1: (v1 - v1)}
        _ = v49
        _ = v50
    }
    v51 := GenS6{F0: -362, F1: v1}
    var v52 int16 = v4
    v53 := []int{v3, 606, v3}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v48:", v48)
    fmt.Println("v52:", v52)
    fmt.Println("slicelen v53:", len(v53), v53)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v47
    _ = v48
    _ = v51
    _ = v52
    _ = v53
}
