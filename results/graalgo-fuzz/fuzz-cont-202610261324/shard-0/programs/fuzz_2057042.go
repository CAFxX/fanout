package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 string; F1 string; F2 string }
type GenS4 struct { F0 int16 }

func main() {
    var v1 int64 = 42
    var v2 float64 = 1e10
    var v3 uint16 = 2
    var v4 bool = (1 >= int32(v3))
    var v5 int16 = 843
    var v6 int16 = (v5 / ((v5) | 1))
    switch "" {
    case "hello":
        var v7 uint32 = 4294967295
        _ = v7
    case "\n\t":
        v4 = (3.14159 < v2)
    }
    v8 := GenS1{F0: uint64(v3)}
    v5 = -291
    if v4 {
        v9 := GenS1{F0: 42}
        var v10 uint64 = 166
        for fi11 := 0; fi11 < 7; fi11++ {
            v1 += (v1 * v1)
            var v12 float32 = (0.1 - float32(3.14159))
            v13 := GenS2{F0: (0 <= 698)}
            _ = fi11
            _ = v12
            _ = v13
        }
        _ = v10
        _ = v9
    } else {
        var v14 uint64 = 973
        for fi15 := 0; fi15 < 20; fi15++ {
            if (42 <= 0) {
                fmt.Println(v5, v2)
            }
            v16 := GenS2{F0: true}
            _ = fi15
            _ = v16
        }
        switch ("foo bar" + ("" + ("hello" + "hello"))) {
        case "hello":
            v17 := GenS1{F0: 673}
            _ = v17
        case "a":
            for fi18 := 0; fi18 < 9; fi18++ {
                v5 -= v6
                _ = fi18
            }
        case "\n\t":
            var v19 float64 = ((float64(float32(((v2 - v2) / ((v2 + 3.14159) / v2)))) * v2) * float64(3.14159))
            _ = v19
        default:
            switch -1 {
            case 355:
                switch "foo bar" {
                case "":
                    for fi20 := 0; fi20 < 16; fi20++ {
                        v3 = 32767
                        v6 = 42
                        _ = fi20
                    }
                case "hello":
                    v14 -= (v14 % ((v14) | 1))
                default:
                    if (!(!(((!v4) && (42 != 0)) && ((v1 * (v1 - (v1 * (v1 >> ((65535) & 63))))) != v1)))) {
                        if v4 {
                            var v21 int16 = v6
                            _ = v21
                        }
                    } else {
                        var v22 int = 128
                        v23 := []int{v22, 32768, (v22 | v22)}
                        _ = v22
                        _ = v23
                    }
                }
            case 42:
                if v4 {
                    var v24 int = 127
                    var v25 int = (^(v24))
                    _ = v24
                    _ = v25
                }
            default:
                v2 += 0.1
            }
        }
        _ = v14
    }
    for fi26 := 0; fi26 < 8; fi26++ {
        v1 -= (-(v1))
        fmt.Println(v2)
        _ = fi26
    }
    if v4 {
        v3 -= (v3 * (v3 / ((v3) | 1)))
        v2 -= v2
    }
    v6 = 42
    fmt.Println(v5)
    var v27 float64 = ((v2 + ((v2 + ((-2.25 * (v2 / 1.0)) + v2)) * (-2.25 - (v2 * -2.25)))) / 1.0)
    v27 = ((-2.25 * 0.1) * v2)
    if (215 < (v3 * (^(v3)))) {
        var v28 uint = 4294967295
        if ((v6 & v5) != int16(v1)) {
            for fi29 := 0; fi29 < 6; fi29++ {
                if v4 {
                    var v30 int = 42
                    _ = v30
                }
                _ = fi29
            }
            var v31 bool = v4
            _ = v31
        }
        var v32 uint8 = 127
        _ = v28
        _ = v32
    }
    fmt.Println(v27, v4)
    v6 -= v6
    if ((v3 - 11) > v3) {
        v5 = v6
    }
    switch 693 {
    case 9223372036854775807:
        v33 := GenS3{F0: "hello", F1: ("" + "foo bar"), F2: "\n\t"}
        _ = v33
    default:
        v6 -= (^(v6))
    }
    v34 := GenS2{F0: (128 != 149)}
    if (v6 > 1000) {
        for fi35 := 0; fi35 < 2; fi35++ {
            for fi36 := 0; fi36 < 6; fi36++ {
                v3 = v3
                v37 := []int{}
                _ = fi36
                _ = v37
            }
            v38 := GenS4{F0: -32768}
            _ = fi35
            _ = v38
        }
        v4 = (v4 && false)
        switch "foo bar" {
        case "foo bar":
            v39 := []int{127, 4294967295, 32768, 2147483647}
            _ = v39
        case "hello":
            fmt.Println(v3)
        default:
            v6 = -338
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
