package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 bool; F1 uint }

func main() {
    var v1 uint8 = 42
    var v2 float32 = (0.1 - (-2.25 + 1.5))
    var v3 int32 = 42
    var v4 int8 = -1
    for fi5 := 0; fi5 < 12; fi5++ {
        v6 := GenS1{F0: 567}
        var v7 int16 = 42
        if true {
            v4 += v4
        }
        _ = fi5
        _ = v6
        _ = v7
    }
    var v8 int32 = v3
    switch (((("foo bar" + "hello") + "foo bar") + "foo bar") + ("\n\t" + "a")) {
    case "\n\t":
        var v9 float32 = v2
        _ = v9
    case "":
        fmt.Println(v3)
    case "hello":
        v10 := []int{}
        _ = v10
    default:
        if true {
            switch "hello" {
            case "":
                if (true && (357 == 18446744073709551615)) {
                    switch 127 {
                    case 1:
                        var v11 uint16 = 42
                        _ = v11
                    case 2147483647:
                        switch 683 {
                        case 2147483647:
                            v12 := GenS1{F0: -309}
                            _ = v12
                        default:
                            if (!(v2 != float32((((0.0 + 3.14159) + (((float64((v2 * 1e10)) - 0.0) / (float64(v2) / 3.14159)) - float64(((-2.25 / (v2 / float32((1.5 * 1.5)))) + v2)))) * 1.5)))) {
                                switch "foo bar" {
                                case "\n\t":
                                    for fi13 := 0; fi13 < 12; fi13++ {
                                        for fi14 := 0; fi14 < 11; fi14++ {
                                            var v15 uint8 = 135
                                            _ = fi14
                                            _ = v15
                                        }
                                        v4 += v4
                                        var v16 uint64 = 344
                                        _ = fi13
                                        _ = v16
                                    }
                                case "a":
                                    for fi17 := 0; fi17 < 15; fi17++ {
                                        v8 += (v3 >> ((314) & 63))
                                        v1 -= 201
                                        _ = fi17
                                    }
                                default:
                                    v18 := GenS1{F0: 42}
                                    _ = v18
                                }
                                var v19 uint = 42
                                _ = v19
                            } else {
                                var v20 int16 = 32767
                                _ = v20
                            }
                        }
                    case -2147483648:
                        v21 := []int{-1, 2147483647, -812, 32768}
                        _ = v21
                    }
                }
            default:
                var v22 int = 32768
                _ = v22
            }
            var v23 uint8 = v1
            fmt.Println(v2, v1, v3)
            _ = v23
        }
    }
    v24 := []int{127, -509, 9223372036854775807, 2147483647}
    fmt.Println(v3, v8, v3)
    switch 199 {
    case 128:
        var v25 int = -753
        _ = v25
    }
    fmt.Println(v3, v2, v3)
    switch "" {
    case "hello":
        switch "a" {
        case "foo bar":
            v1 = 0
        case "":
            var v26 uint64 = 42
            _ = v26
        }
    default:
        v27 := GenS2{F0: (((!(!false)) && (-438 != 9223372036854775807)) || true), F1: 1}
        _ = v27
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v8
}
