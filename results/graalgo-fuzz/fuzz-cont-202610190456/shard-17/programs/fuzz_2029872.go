package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 int16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint8; F1 int32 }
type GenS4 struct { F0 int16; F1 string; F2 int8 }
type GenS5 struct { F0 uint8 }

func main() {
    var v1 int64 = 42
    v2 := make(map[string]string)
    var v3 int8 = 22
    var v4 uint64 = 380
    switch 183 {
    case -69:
        switch ("\n\t" + "") {
        case "\n\t":
            v3 += v3
        case "":
            v5 := []int{2147483647}
            _ = v5
        case "foo bar":
            var v6 float32 = float32((1e10 + 3.14159))
            _ = v6
        }
    case 1:
        v7 := GenS1{F0: 2, F1: 133, F2: 1000}
        _ = v7
    case 32767:
        if true {
            if false {
                for fi8 := 0; fi8 < 1; fi8++ {
                    switch 4294967295 {
                    case 32767:
                        v9 := []int{}
                        _ = v9
                    case -791:
                        v2["hello"] = ("hello" + "hello")
                    case 128:
                        switch fi8 {
                        case -9223372036854775808:
                            v10 := []int{fi8, 128, 1}
                            _ = v10
                        case 42:
                            v1 -= v1
                        case 71:
                            switch "\n\t" {
                            case "":
                                var v11 uint8 = 0
                                _ = v11
                            case "hello":
                                v3 -= (v3 & (v3 / (((^(v3))) | 1)))
                            case "foo bar":
                                if (!true) {
                                    for fi12 := 0; fi12 < 20; fi12++ {
                                        v1 += v1
                                        switch "hello" {
                                        case "foo bar":
                                            v13 := GenS2{F0: 143}
                                            _ = v13
                                        case "a":
                                            fmt.Println(v1, v4)
                                        }
                                        v4 += v4
                                        _ = fi12
                                    }
                                    v1 -= v1
                                }
                            default:
                                var v14 bool = true
                                _ = v14
                            }
                        default:
                            v4 += (v4 ^ (^(v4)))
                        }
                    default:
                        if (((-2.25 - float64((0.1 / float32(1.5)))) - (1.5 + (((-2.25 - (1.5 - 1e10)) + 0.1) / -2.25))) != 0.0) {
                            var v15 int16 = 42
                            v16 := GenS1{F0: 381, F1: 107, F2: (v15 * v15)}
                            _ = v15
                            _ = v16
                        } else {
                            var v17 string = ((("hello" + "a") + (("\n\t" + ("a" + "a")) + ("a" + "\n\t"))) + ("\n\t" + "hello"))
                            _ = v17
                        }
                    }
                    _ = fi8
                }
            } else {
                v1 += v1
            }
        } else {
            switch "\n\t" {
            case "":
                switch "hello" {
                case "hello":
                    if (!false) {
                        var v18 uint16 = 0
                        var v19 uint16 = (v18 & (v18 ^ v18))
                        v20 := GenS1{F0: 767, F1: 0, F2: 0}
                        _ = v18
                        _ = v19
                        _ = v20
                    }
                case "\n\t":
                    var v21 uint8 = 147
                    _ = v21
                default:
                    for fi22 := 0; fi22 < 13; fi22++ {
                        var v23 bool = ((false || true) || (931 == -2147483648))
                        fmt.Println(fi22, v4)
                        v24 := GenS3{F0: 2, F1: 275}
                        _ = fi22
                        _ = v23
                        _ = v24
                    }
                }
            case "foo bar":
                v25 := GenS4{F0: 1, F1: ("foo bar" + "hello"), F2: 1}
                _ = v25
            }
        }
    default:
        for fi26 := 0; fi26 < 15; fi26++ {
            fmt.Println(v4, v3)
            _ = fi26
        }
    }
    v27 := GenS5{F0: 127}
    v28 := GenS2{F0: 65535}
    var v29 uint16 = 782
    if true {
        v3 = int8(v29)
        fmt.Println(v1)
    }
    v1 = (v1 >> ((271) & 63))
    fmt.Println(v4, v3, v4)
    v30 := GenS3{F0: 255, F1: -609}
    var v31 bool = ((^(v29)) == (v29 * (v29 / ((v29) | 1))))
    var v32 uint = 18446744073709551615
    var v33 uint32 = 42
    v1 += (v1 & v1)
    v3 += (v3 / ((-105) | 1))
    if false {
        if ((0 == -128) && (("" + "\n\t") != "a")) {
            if v31 {
                var v34 uint = (v32 % ((uint(v33)) | 1))
                _ = v34
            }
        }
        switch "a" {
        case "foo bar":
            var v35 uint32 = v33
            _ = v35
        case "\n\t":
            if ((v3 | (v3 >> ((v32) & 63))) > int8(v29)) {
                v36 := GenS2{F0: (v32 - (v32 ^ 4294967295))}
                fmt.Println(v1, v1, v3)
                v4 -= 297
                _ = v36
            } else {
                v37 := v2["\n\t"]
                _ = v37
            }
        }
        fmt.Println(v29)
    }
    var v38 uint64 = v4
    v32 += (v32 + 42)
    for fi39 := 0; fi39 < 4; fi39++ {
        v29 -= (v29 % ((713) | 1))
        var v40 float64 = float64(1e10)
        _ = fi39
        _ = v40
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v38
    _ = v4
}
