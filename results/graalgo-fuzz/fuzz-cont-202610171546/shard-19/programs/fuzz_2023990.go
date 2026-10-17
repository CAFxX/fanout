package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 string; F1 int16 }

func main() {
    var v1 int32 = -2147483648
    var v2 int64 = 581
    var v3 uint = 18446744073709551615
    var v4 uint16 = 985
    var v5 float32 = float32((-2.25 + (1e10 + 3.14159)))
    var v6 uint = (v3 ^ 4294967295)
    if (0 >= (v3 * 7)) {
        v7 := GenS1{F0: 0, F1: 18446744073709551615}
        if ((v4 * (v4 % ((v4) | 1))) == v4) {
            switch 32767 {
            case 32767:
                if (-1 >= 32767) {
                    var v8 int = 853
                    _ = v8
                } else {
                    v6 += (v3 & (v6 ^ (v6 >> (((^(v6))) & 63))))
                }
            case -983:
                v4 = (v4 >> (((v3 / ((42) | 1))) & 63))
            case 127:
                var v9 string = "a"
                _ = v9
            }
            fmt.Println(v3)
            if (!(!(false || false))) {
                v10 := GenS1{F0: 497, F1: 0}
                switch "hello" {
                case "":
                    switch "hello" {
                    case "\n\t":
                        switch 413 {
                        case 32767:
                            v11 := []int{-486, -106, 767}
                            _ = v11
                        default:
                            v2 = (v2 >> ((v3) & 63))
                        }
                    case "a":
                        v6 = (v3 & 1)
                    case "hello":
                        if false {
                            if false {
                                fmt.Println(v6, v5)
                            }
                        }
                    }
                }
                v2 = (v2 & v2)
                _ = v10
            } else {
                switch ("a" + "foo bar") {
                case "hello":
                    v12 := GenS1{F0: 0, F1: 2}
                    _ = v12
                case "foo bar":
                    v13 := []int{}
                    _ = v13
                case "a":
                    for fi14 := 0; fi14 < 3; fi14++ {
                        v15 := GenS2{F0: (("\n\t" + "\n\t") + "hello")}
                        v16 := GenS1{F0: 1000, F1: 2}
                        fmt.Println(v1, v5, v6)
                        _ = fi14
                        _ = v15
                        _ = v16
                    }
                default:
                    switch "" {
                    case "\n\t":
                        v3 = (v6 | 119)
                    case "":
                        v6 -= uint(v4)
                    case "a":
                        var v17 uint = v6
                        _ = v17
                    }
                }
                for fi18 := 0; fi18 < 2; fi18++ {
                    var v19 uint16 = v4
                    _ = fi18
                    _ = v19
                }
            }
        }
        _ = v7
    } else {
        var v20 uint64 = uint64(v4)
        _ = v20
    }
    if ((false && false) || ((18446744073709551615 == 0) && ((true || (!true)) || true))) {
        v6 -= v6
        var v21 uint32 = 2
        v5 += v5
        _ = v21
    } else {
        v22 := []int{1, 588, 4294967295}
        var v23 uint64 = 42
        _ = v22
        _ = v23
    }
    v24 := GenS3{F0: ("a" + ("hello" + "foo bar")), F1: 0}
    v4 = v4
    v25 := GenS1{F0: 0, F1: 4294967295}
    var v26 uint64 = 4294967295
    if (v4 <= (^(v4))) {
        var v27 int16 = -1
        if (!false) {
            v28 := []int{127, 1}
            v29 := GenS2{F0: "\n\t"}
            _ = v28
            _ = v29
        }
        v3 += (v3 | uint(v4))
        _ = v27
    }
    switch 42 {
    case 4294967295:
        fmt.Println(v6, v3, v1)
    case 9223372036854775807:
        var v30 uint8 = uint8(v26)
        _ = v30
    default:
        var v31 float64 = ((0.1 + (1e10 + (-2.25 * 3.14159))) - 3.14159)
        _ = v31
    }
    switch 319 {
    case 42:
        var v32 int = 128
        _ = v32
    case 127:
        fmt.Println(v6, v26, v5)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
