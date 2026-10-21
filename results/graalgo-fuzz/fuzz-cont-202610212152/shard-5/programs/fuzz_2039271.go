package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 bool; F1 int; F2 int8 }
type GenS3 struct { F0 uint }

func main() {
    v1 := make(map[string]bool)
    var v2 float32 = 1.5
    v3 := []string{(("\n\t" + "a") + "a"), (((((((("\n\t" + ("foo bar" + "")) + "hello") + "a") + ((("\n\t" + "hello") + "hello") + "")) + ("a" + "foo bar")) + ("\n\t" + "a")) + "hello") + "a"), ("hello" + ""), "foo bar"}
    var v4 uint32 = 4294967295
    var v5 int16 = -1
    v6 := make(map[int]int)
    v2 = (float32(-2.25) - v2)
    v7 := GenS1{F0: 827}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if true {
        if true {
            for fi8 := 0; fi8 < 1; fi8++ {
                v9 := GenS2{F0: (2147483647 > 2), F1: (fi8 * (fi8 % ((42) | 1))), F2: 100}
                _ = fi8
                _ = v9
            }
        }
    }
    switch 128 {
    case 9223372036854775807:
        var v10 bool = (1 > 593)
        _ = v10
    case 2147483647:
        fmt.Println(v4, v2)
    }
    fmt.Println(v5)
    switch -2147483648 {
    case 1:
        v2 = v2
    case 32768:
        v11 := GenS3{F0: 48}
        _ = v11
    case 127:
        if false {
            v2 = (3.14159 - ((1.5 + (float32(float64(float32((float64(1.5) * (1e10 / (-2.25 / -2.25)))))) / v2)) + 0.0))
        }
    }
    for fi12 := 0; fi12 < 6; fi12++ {
        var v13 bool = (v4 != v4)
        switch ("a" + "hello") {
        case "a":
            if true {
                switch fi12 {
                case 503:
                    switch -440 {
                    case -9223372036854775808:
                        v5 = (^(v5))
                    case 1:
                        fmt.Println(v2, v2, v2)
                    case -964:
                        var v14 float64 = 1.5
                        _ = v14
                    default:
                        fmt.Println(v13, v4)
                    }
                case -2147483648:
                    v6[(fi12 >> ((18446744073709551615) & 63))] = (fi12 << ((323) & 63))
                case 42:
                    for fi15 := 0; fi15 < 18; fi15++ {
                        switch "\n\t" {
                        case "\n\t":
                            var v16 uint32 = v4
                            _ = v16
                        case "foo bar":
                            fmt.Println(v13)
                        case "":
                            v3 = append(v3, ("hello" + ("" + ("" + "foo bar"))))
                        }
                        _ = fi15
                    }
                default:
                    var v17 float32 = 1.5
                    _ = v17
                }
                v5 = (^(v5))
            } else {
                v4 += 4294967295
                fmt.Println(fi12)
                v18 := v6[fi12]
                _ = v18
            }
        case "":
            var v19 uint8 = 1
            _ = v19
        default:
            v4 = v4
        }
        for fi20 := 0; fi20 < 7; fi20++ {
            v6[fi20] = (fi12 & -2147483648)
            v2 -= float32((float64(((float32(3.14159) * (0.1 - v2)) + 1e10)) * 0.0))
            _ = fi20
        }
        _ = fi12
        _ = v13
    }
    if (!false) {
        var v21 int32 = -455
        _ = v21
    } else {
        var v22 bool = (false || true)
        _ = v22
    }
    var v23 uint8 = 41
    if false {
        if false {
            fmt.Println(v23, v4)
            v4 = v4
            v23 = (v23 % ((v23) | 1))
        }
        v23 += 1
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
