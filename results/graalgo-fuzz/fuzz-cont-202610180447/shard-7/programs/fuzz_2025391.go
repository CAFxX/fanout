package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint16; F1 int16; F2 uint }
type GenS3 struct { F0 string; F1 int32 }
type GenS4 struct { F0 int8; F1 uint64; F2 uint }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 bool; F1 int; F2 bool }

func main() {
    var v1 int = -2147483648
    var v2 int8 = -128
    var v3 uint = 2
    v4 := []int{(v1 | v1), 656, (v1 ^ v1)}
    for fi5 := 0; fi5 < 4; fi5++ {
        if false {
            v6 := GenS1{F0: 32767}
            var v7 string = ("" + "a")
            v8 := GenS1{F0: 1000}
            _ = v6
            _ = v7
            _ = v8
        } else {
            switch (("foo bar" + "") + "foo bar") {
            case "\n\t":
                v9 := GenS2{F0: 1, F1: -831, F2: 18446744073709551615}
                _ = v9
            case "foo bar":
                v10 := GenS2{F0: 42, F1: 42, F2: (v3 - 4294967295)}
                _ = v10
            case "a":
                switch v1 {
                case 32768:
                    fmt.Println(v1, v3, v1)
                default:
                    for fi11 := 0; fi11 < 8; fi11++ {
                        if (142 < 1) {
                            v2 = 100
                        }
                        v12 := GenS1{F0: 1000}
                        v3 = (^(v3))
                        _ = fi11
                        _ = v12
                    }
                }
            }
        }
        switch (v1 * (^(v1))) {
        case 1:
            v4 = append(v4, 1)
        case 128:
            var v13 string = "a"
            _ = v13
        }
        for fi14 := 0; fi14 < 20; fi14++ {
            v15 := GenS1{F0: 444}
            v1 = fi5
            _ = fi14
            _ = v15
        }
        _ = fi5
    }
    switch ((("foo bar" + "a") + "a") + "a") {
    case "foo bar":
        v2 = v2
    default:
        for fi16 := 0; fi16 < 2; fi16++ {
            for fi17 := 0; fi17 < 2; fi17++ {
                switch "\n\t" {
                case "\n\t":
                    switch "\n\t" {
                    case "hello":
                        v2 = 99
                    }
                    v3 = (v3 ^ (v3 + (v3 | v3)))
                case "hello":
                    _ = len(v4)
                default:
                    fmt.Println(fi17)
                }
                v3 = (v3 + (v3 | (v3 << (((v3 & (^(v3)))) & 63))))
                v18 := GenS3{F0: "\n\t", F1: 2147483647}
                fmt.Println(v2, fi17)
                _ = fi17
                _ = v18
            }
            for fi19 := 0; fi19 < 8; fi19++ {
                v20 := GenS4{F0: -1, F1: 883, F2: (v3 >> ((255) & 63))}
                v3 = v3
                if true {
                    switch (fi16 * fi16) {
                    case 9223372036854775807:
                        v3 = (v3 % ((v3) | 1))
                    default:
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    }
                    var v21 string = "\n\t"
                    _ = v21
                } else {
                    switch ("foo bar" + ("hello" + "")) {
                    case "a":
                        var v22 uint32 = 0
                        _ = v22
                    default:
                        var v23 int32 = -580
                        var v24 int16 = -32768
                        _ = v23
                        _ = v24
                    }
                    switch ("foo bar" + ("foo bar" + ("a" + ""))) {
                    case "a":
                        switch ("a" + "") {
                        case "a":
                            v2 -= v2
                            var v25 float64 = float64(((3.14159 - ((1.5 - 3.14159) / (3.14159 - float32((float64((0.1 * 0.0)) / ((1e10 / 1.5) + float64(1.5))))))) * 1e10))
                            _ = v25
                        default:
                            switch (-(fi19)) {
                            case -2147483648:
                                var v26 string = "\n\t"
                                _ = v26
                            case 873:
                                var v27 uint64 = 0
                                _ = v27
                            default:
                                v28 := GenS2{F0: 65535, F1: -698, F2: 57}
                                _ = v28
                            }
                        }
                    case "foo bar":
                        var v29 uint8 = 42
                        _ = v29
                    case "\n\t":
                        _ = len(v4)
                    }
                    if (0 >= 195) {
                        fmt.Println(v3, v2, v1)
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        v30 := GenS5{F0: uint64(v3)}
                        _ = v30
                    }
                }
                _ = fi19
                _ = v20
            }
            fmt.Println(v1, v1)
            _ = fi16
        }
    }
    for fi31 := 0; fi31 < 15; fi31++ {
        if len(v4) > 0 {
            v4[0] = -395
        }
        var v32 int16 = -977
        _ = fi31
        _ = v32
    }
    v1 = 9223372036854775807
    v3 = 4294967295
    var v33 uint = v3
    var v34 bool = true
    v35 := GenS6{F0: v34, F1: v1, F2: (((-32768 != 42) && ((2 > 4294967295) || (-2.25 == 0.0))) && (1 == 2147483647))}
    var v36 int8 = 86
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v4
}
