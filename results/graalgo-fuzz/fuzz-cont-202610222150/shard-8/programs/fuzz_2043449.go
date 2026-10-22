package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int64; F1 int16; F2 string }
type GenS3 struct { F0 int32; F1 uint16 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int64; F1 int64 }
type GenS6 struct { F0 int64; F1 uint }
type GenS7 struct { F0 int16; F1 int64; F2 string }

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 uint8 = 85
    v3 := make(map[string]bool)
    var v4 float32 = float32(float64(0.0))
    v5 := make(map[uint32]int)
    for fi6 := 0; fi6 < 10; fi6++ {
        v5[42] = (fi6 & fi6)
        if (v2 != v2) {
            if false {
                _ = len(v3)
                _ = len(v3)
            } else {
                fmt.Println(v1, v2)
                v7 := GenS1{F0: "foo bar"}
                _ = v7
            }
            fmt.Println(v4, v1, v1)
            v8 := []int{(fi6 >> ((0) & 63)), fi6, (fi6 & 4294967295)}
            _ = v8
        }
        _ = fi6
    }
    v1 -= v1
    v2 += (^(v2))
    var v9 uint8 = (v2 + 1)
    fmt.Println(v9, v4)
    switch "a" {
    case "\n\t":
        var v10 uint64 = 0
        _ = v10
    case "a":
        v11 := GenS2{F0: 61, F1: 32767, F2: ""}
        _ = v11
    default:
        v2 -= v9
    }
    switch ("\n\t" + "a") {
    case "\n\t":
        var v12 string = ("\n\t" + "foo bar")
        _ = v12
    case "hello":
        var v13 uint8 = 128
        _ = v13
    default:
        switch "" {
        case "":
            if (-89 <= -9223372036854775808) {
                for fi14 := 0; fi14 < 9; fi14++ {
                    v2 += v2
                    v15 := GenS3{F0: 42, F1: uint16(v1)}
                    for fi16 := 0; fi16 < 6; fi16++ {
                        var v17 bool = false
                        v2 = (v2 % ((255) | 1))
                        _ = fi16
                        _ = v17
                    }
                    _ = fi14
                    _ = v15
                }
                v4 = (3.14159 / v4)
                v18 := []int{117}
                _ = v18
            } else {
                if (!true) {
                    for fi19 := 0; fi19 < 9; fi19++ {
                        v20 := []int{(fi19 << ((18446744073709551615) & 63))}
                        for fi21 := 0; fi21 < 10; fi21++ {
                            v22 := GenS1{F0: (((((("hello" + "a") + ("" + "foo bar")) + ((("foo bar" + "") + "hello") + "hello")) + "foo bar") + "foo bar") + "hello")}
                            v9 = (v9 / (((v9 >> ((132) & 63))) | 1))
                            fmt.Println(fi21, fi21, fi21)
                            _ = fi21
                            _ = v22
                        }
                        _ = fi19
                        _ = v20
                    }
                }
            }
        }
    }
    var v23 int16 = 1000
    if false {
        if true {
            if ((1 == 516) || (32767 != 42)) {
                _ = len(v5)
                switch -9223372036854775808 {
                case 0:
                    v24 := GenS4{F0: (false || true)}
                    _ = v24
                case 9223372036854775807:
                    var v25 int16 = v23
                    _ = v25
                }
                var v26 int32 = -1
                _ = v26
            } else {
                v23 = v23
                switch -248 {
                case 127:
                    var v27 int32 = -2147483648
                    _ = v27
                case -2147483648:
                    switch ("a" + ("hello" + "\n\t")) {
                    case "hello":
                        v28 := GenS5{F0: 1, F1: 788}
                        _ = v28
                    case "a":
                        v1 = v1
                    case "foo bar":
                        delete(v5, 0)
                    default:
                        if (v4 > v4) {
                            v2 = 35
                            for fi29 := 0; fi29 < 8; fi29++ {
                                v30 := GenS5{F0: 42, F1: -9223372036854775808}
                                for fi31 := 0; fi31 < 9; fi31++ {
                                    var v32 uint = 1
                                    fmt.Println(v32, v32)
                                    _ = fi31
                                    _ = v32
                                }
                                delete(v5, 596)
                                _ = fi29
                                _ = v30
                            }
                        }
                    }
                case 1:
                    v33 := GenS6{F0: -1, F1: 702}
                    _ = v33
                default:
                    switch "a" {
                    case "a":
                        delete(v3, (("hello" + "foo bar") + "\n\t"))
                    }
                }
                var v34 int16 = -32768
                _ = v34
            }
            fmt.Println(v4)
        }
    }
    var v35 int32 = 2147483647
    v36 := GenS7{F0: int16(v9), F1: 654, F2: ""}
    v37 := []int{}
    if ((((3.14159 != -2.25) || ((3.14159 * (float64((v4 * (float32(-2.25) + v4))) * 0.0)) <= (0.0 + ((((3.14159 / float64(((float32(1e10) - float32((-2.25 - (1e10 - ((3.14159 - 0.0) - (-2.25 / (1e10 / 1e10))))))) - ((1.5 / (v4 - 0.0)) / v4)))) - 1e10) - -2.25) * (1e10 * (-2.25 - (0.1 - 1e10))))))) || (466 >= 255)) || false) {
        v38 := GenS5{F0: 9223372036854775807, F1: 629}
        _ = v38
    }
    delete(v3, "")
    v39 := GenS2{F0: -506, F1: (v23 % ((v23) | 1)), F2: "a"}
    v23 = (v23 & (v23 + (-(v23))))
    fmt.Println(v23)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v37:", len(v37), v37)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v39
    _ = v4
    _ = v5
    _ = v9
}
