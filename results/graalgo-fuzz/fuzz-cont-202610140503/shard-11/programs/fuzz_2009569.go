package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool; F2 int16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int8; F1 uint32; F2 int }
type GenS4 struct { F0 uint16; F1 int64 }
type GenS5 struct { F0 int32; F1 uint8 }

func main() {
    var v1 float64 = 1e10
    var v2 float32 = float32(0.1)
    var v3 int32 = -2147483648
    v4 := make(map[int]string)
    var v5 int64 = 1
    if (true || false) {
        if (718 != 2) {
            switch ("foo bar" + "a") {
            case "a":
                switch "\n\t" {
                case "\n\t":
                    var v6 int = -128
                    _ = v6
                case "foo bar":
                    fmt.Println(v1)
                }
            }
            v7 := GenS1{F0: 127, F1: ((28 == 0) && true), F2: int16(v3)}
            v5 -= v5
            _ = v7
        } else {
            switch 898 {
            case 42:
                v8 := GenS2{F0: -1}
                _ = v8
            case 127:
                switch 9223372036854775807 {
                case 2147483647:
                    v1 -= (v1 / v1)
                }
            }
            v1 -= (float64(-2.25) / v1)
            if (v2 == v2) {
                switch "" {
                case "hello":
                    var v9 int8 = 1
                    _ = v9
                }
                _ = len(v4)
            }
        }
    }
    var v10 int32 = (v3 * v3)
    switch ("\n\t" + "hello") {
    case "a":
        for fi11 := 0; fi11 < 3; fi11++ {
            switch 1 {
            case 32767:
                var v12 int64 = v5
                _ = v12
            case 88:
                v13 := GenS2{F0: 32767}
                _ = v13
            default:
                var v14 int = fi11
                _ = v14
            }
            _ = fi11
        }
    default:
        v15 := GenS3{F0: -128, F1: 431, F2: -9223372036854775808}
        _ = v15
    }
    if (255 < 185) {
        var v16 uint8 = 128
        switch ("hello" + "foo bar") {
        case "hello":
            var v17 int = 9223372036854775807
            _ = v17
        case "\n\t":
            for fi18 := 0; fi18 < 1; fi18++ {
                if (-125 == -1) {
                    var v19 uint32 = 0
                    for fi20 := 0; fi20 < 4; fi20++ {
                        var v21 int16 = 1000
                        var v22 float32 = (float32(0.0) * float32(v1))
                        var v23 uint = 594
                        var v24 uint8 = 127
                        _ = fi20
                        _ = v21
                        _ = v22
                        _ = v23
                        _ = v24
                    }
                    v2 = (v2 + ((float32(v1) - (float32(3.14159) + (v2 - (v2 + (1.5 * (v2 * ((v2 - v2) - v2))))))) - float32(1.5)))
                    _ = v19
                } else {
                    var v25 uint8 = v16
                    _ = v25
                }
                _ = fi18
            }
        case "a":
            var v26 int8 = -75
            _ = v26
        default:
            v16 += 42
        }
        _ = v16
    } else {
        v27 := []int{}
        _ = v27
    }
    v4[-1] = "\n\t"
    var v28 uint16 = 15
    fmt.Println(v1)
    var v29 uint = 0
    var v30 uint16 = v28
    for fi31 := 0; fi31 < 18; fi31++ {
        v32 := GenS4{F0: v30, F1: 0}
        var v33 string = "a"
        _ = fi31
        _ = v32
        _ = v33
    }
    v29 += 81
    for fi34 := 0; fi34 < 11; fi34++ {
        fmt.Println(v28, v5, fi34)
        v35 := GenS5{F0: -2147483648, F1: 255}
        var v36 uint = (v29 % ((v29) | 1))
        _ = fi34
        _ = v35
        _ = v36
    }
    switch ("" + "a") {
    case "a":
        v37 := []int{-59, 2147483647, 128}
        _ = v37
    case "\n\t":
        v38 := []int{446, 128, -1}
        _ = v38
    default:
        v39 := GenS4{F0: 255, F1: v5}
        _ = v39
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
