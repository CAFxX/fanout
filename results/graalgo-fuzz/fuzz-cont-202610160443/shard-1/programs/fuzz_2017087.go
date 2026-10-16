package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64; F2 bool }
type GenS2 struct { F0 int32; F1 int32 }

func main() {
    var v1 uint64 = 42
    var v2 string = "a"
    var v3 int = 32768
    v4 := make(map[uint32]string)
    v5 := []int{v3, v3, v3}
    var v6 string = v2
    if (!(!((v3 << ((65535) & 63)) < 0))) {
        var v7 int16 = 32767
        _ = v7
    }
    v8 := GenS1{F0: 2, F1: 4294967295, F2: ((1.5 * 3.14159) <= ((0.1 / (0.0 + (0.0 + 0.1))) + 1e10))}
    var v9 int16 = -32768
    for fi10 := 0; fi10 < 1; fi10++ {
        for fi11 := 0; fi11 < 20; fi11++ {
            var v12 float64 = 0.0
            _ = len(v5)
            v13 := v4[28]
            fmt.Println(v3)
            _ = fi11
            _ = v12
            _ = v13
        }
        _ = fi10
    }
    v14 := GenS1{F0: 32767, F1: 1, F2: false}
    v1 -= v1
    switch 731 {
    case 9223372036854775807:
        fmt.Println(v9)
    case -911:
        if len(v5) > 0 {
            v5[0] = 32767
        }
    default:
        switch "foo bar" {
        case "foo bar":
            var v15 int8 = -1
            _ = v15
        case "hello":
            switch 32768 {
            case 466:
                if (0.1 > float32(((0.0 / 1.0) - 1e10))) {
                    if len(v5) > 0 {
                        v5[0] = (-(v3))
                    }
                }
            case -9223372036854775808:
                var v16 uint32 = 2147483647
                _ = v16
            case 9223372036854775807:
                v5 = append(v5, -367)
            default:
                fmt.Println(v2)
            }
        default:
            if true {
                v2 = "foo bar"
                if true {
                    v17 := GenS1{F0: 32767, F1: 4294967295, F2: (1 > 18446744073709551615)}
                    _ = v17
                } else {
                    v18 := GenS2{F0: 0, F1: 2147483647}
                    if len(v5) > 0 {
                        v5[0] = v3
                    }
                    _ = v18
                }
                for fi19 := 0; fi19 < 6; fi19++ {
                    for fi20 := 0; fi20 < 7; fi20++ {
                        var v21 int16 = v9
                        _ = fi20
                        _ = v21
                    }
                    v22 := GenS1{F0: 965, F1: v1, F2: (!false)}
                    _ = fi19
                    _ = v22
                }
            } else {
                if false {
                    var v23 float64 = (-2.25 - (-2.25 * 1e10))
                    _ = v23
                }
                v4[2147483648] = ""
            }
        }
    }
    var v24 bool = (false || false)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v14
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
