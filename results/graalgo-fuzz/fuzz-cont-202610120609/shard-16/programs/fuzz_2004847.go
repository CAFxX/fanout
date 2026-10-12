package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int8; F1 uint32 }
func genFn0() uint {
    var v4 int = -1
    switch "" {
    case "a":
        var v5 string = "foo bar"
        _ = v5
    case "foo bar":
        v4 += 42
    case "\n\t":
        if ((!true) && (("foo bar" + ("a" + "")) >= "a")) {
            for fi6 := 0; fi6 < 4; fi6++ {
                v4 -= v4
                v4 = (fi6 / ((457) | 1))
                var v7 bool = ((-2.25 == 1.5) || (false || (42 <= 2147483648)))
                _ = fi6
                _ = v7
            }
        } else {
            v4 -= v4
            var v8 uint32 = 142
            _ = v8
        }
    }
    var v9 int16 = 1000
    _ = v4
    _ = v9
    return 109
}
func genFn1(p0 bool) uint {
    fmt.Println(p0)
    switch "hello" {
    case "a":
        var v10 int16 = -995
        _ = v10
    case "\n\t":
        v11 := GenS1{F0: "a"}
        _ = v11
    case "hello":
        for fi12 := 0; fi12 < 16; fi12++ {
            var v13 bool = (p0 && p0)
            fmt.Println(p0, v13, v13)
            v13 = p0
            _ = fi12
            _ = v13
        }
    default:
        p0 = ((0.1 * ((1.5 + -2.25) * ((((float64(1e10) * -2.25) / 1.0) - (1e10 / 1.0)) + 1e10))) == (3.14159 + 1e10))
    }
    for fi14 := 0; fi14 < 19; fi14++ {
        var v15 int8 = -1
        v15 = v15
        switch "\n\t" {
        case "foo bar":
            switch -730 {
            case -889:
                for fi16 := 0; fi16 < 1; fi16++ {
                    if (56 != (v15 + (v15 + 100))) {
                        v17 := GenS1{F0: "a"}
                        v18 := GenS1{F0: "foo bar"}
                        _ = v17
                        _ = v18
                    }
                    _ = fi16
                }
            default:
                var v19 uint32 = 2147483647
                _ = v19
            }
        case "a":
            v15 += (^(v15))
        case "":
            fmt.Println(fi14, v15, fi14)
        }
        _ = fi14
        _ = v15
    }
    for fi20 := 0; fi20 < 9; fi20++ {
        p0 = (-9223372036854775808 <= -9223372036854775808)
        for fi21 := 0; fi21 < 18; fi21++ {
            var v22 int = fi20
            _ = fi21
            _ = v22
        }
        _ = fi20
    }
    if p0 {
        v23 := GenS2{F0: 127, F1: 0}
        _ = v23
    } else {
        var v24 int16 = 42
        var v25 uint32 = 42
        _ = v24
        _ = v25
    }
    _ = p0
    return 97
}

func main() {
    var v1 uint32 = 882
    var v2 uint32 = v1
    v3 := []string{("hello" + ("a" + "foo bar")), ("hello" + "foo bar")}
    if len(v3) > 0 {
        v3[0] = ""
    }
    switch "" {
    case "a":
        v1 -= v1
    case "hello":
        if (((-2.25 + 3.14159) * (1e10 / float32(float64(0.1)))) > float32(3.14159)) {
            v1 = v2
            fmt.Println(v2, v2)
        } else {
            for fi26 := 0; fi26 < 17; fi26++ {
                _ = len(v3)
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                v27 := GenS1{F0: ("" + "hello")}
                _ = fi26
                _ = v27
            }
        }
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v1 = 0
    v2 -= v2
    v28 := GenS2{F0: -1, F1: v2}
    var v29 float64 = (0.0 / 1.5)
    if true {
        for fi30 := 0; fi30 < 4; fi30++ {
            fmt.Println(v1, fi30, fi30)
            _ = fi30
        }
        v29 += (v29 / (0.1 + (3.14159 - 0.1)))
        var v31 int32 = 1
        _ = v31
    } else {
        if len(v3) > 0 {
            v3[0] = "hello"
        }
    }
    var v32 string = ("foo bar" + (("a" + "a") + ""))
    v29 = float64((-2.25 * 0.1))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v32
}
