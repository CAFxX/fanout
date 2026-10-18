package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 bool; F1 uint }

func main() {
    var v1 int64 = 1
    v2 := []string{("" + ""), "\n\t", ("a" + "\n\t"), "foo bar"}
    var v3 uint16 = 763
    v4 := []int{}
    if ((true && (0 >= 9223372036854775807)) || ((v1 | (-(v1))) <= v1)) {
        fmt.Println(v3, v3)
        v1 = v1
    }
    var v5 float64 = 0.1
    _ = len(v4)
    v3 += v3
    v5 = 0.0
    v3 = v3
    if (((((-2.25 - 1.5) + 0.0) - ((float64((0.0 * ((0.1 * (0.1 + float32(1.5))) * 0.0))) + -2.25) + float64(3.14159))) * (v5 - v5)) <= 0.1) {
        for fi6 := 0; fi6 < 14; fi6++ {
            v7 := GenS1{F0: 255}
            if false {
                v5 -= 3.14159
            }
            _ = fi6
            _ = v7
        }
        switch -285 {
        case -1:
            for fi8 := 0; fi8 < 17; fi8++ {
                if true {
                    var v9 int8 = 1
                    fmt.Println(v5, fi8)
                    for fi10 := 0; fi10 < 2; fi10++ {
                        v1 = (v1 - v1)
                        fmt.Println(v3, fi10, v9)
                        v2 = append(v2, ("\n\t" + "foo bar"))
                        v2 = append(v2, ("\n\t" + ("\n\t" + ("\n\t" + "foo bar"))))
                        _ = fi10
                    }
                    _ = v9
                }
                var v11 int32 = -1
                v3 = v3
                _ = fi8
                _ = v11
            }
        case 32768:
            v5 += (v5 / 1.0)
        case 127:
            var v12 float32 = 1.5
            _ = v12
        }
    } else {
        fmt.Println(v1, v1, v3)
        v3 = (^(v3))
    }
    var v13 float32 = float32(v5)
    v2 = append(v2, "")
    if false {
        for fi14 := 0; fi14 < 17; fi14++ {
            if false {
                var v15 uint64 = 1
                v15 -= 372
                _ = v15
            }
            v5 = 1e10
            _ = fi14
        }
    } else {
        v16 := GenS1{F0: 956}
        _ = v16
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    switch -802 {
    case 0:
        switch "\n\t" {
        case "foo bar":
            v17 := GenS1{F0: 306}
            _ = v17
        case "hello":
            if len(v4) > 0 {
                v4[0] = -1
            }
        default:
            fmt.Println(v5, v5)
        }
    default:
        for fi18 := 0; fi18 < 12; fi18++ {
            fmt.Println(v1)
            v19 := GenS1{F0: 4294967295}
            _ = fi18
            _ = v19
        }
    }
    for fi20 := 0; fi20 < 18; fi20++ {
        switch (-(fi20)) {
        case -9223372036854775808:
            for fi21 := 0; fi21 < 3; fi21++ {
                fmt.Println(v1)
                var v22 uint64 = 0
                switch "a" {
                case "hello":
                    if (255 >= 216) {
                        v2 = append(v2, ("\n\t" + "a"))
                    }
                    v23 := GenS2{F0: true, F1: 985}
                    v1 += (-(v1))
                    v3 -= v3
                    _ = v23
                case "foo bar":
                    var v24 uint64 = (v22 * (v22 << ((18446744073709551615) & 63)))
                    _ = v24
                }
                _ = fi21
                _ = v22
            }
        default:
            if (871 <= 2147483648) {
                if len(v2) > 0 {
                    v2[0] = "foo bar"
                }
                v13 = v13
                v25 := GenS1{F0: 708}
                _ = v25
            }
        }
        _ = fi20
    }
    var v26 string = ("foo bar" + "foo bar")
    if len(v2) > 0 {
        v2[0] = (("\n\t" + v26) + v26)
    }
    var v27 string = (v26 + (v26 + ("foo bar" + (v26 + v26))))
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
