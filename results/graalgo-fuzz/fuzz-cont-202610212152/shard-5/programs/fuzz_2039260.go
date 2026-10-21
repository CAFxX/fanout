package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string; F2 int8 }
func genFn0() uint32 {
    v4 := []int{127, 1, 0, -938}
    var v5 int8 = 1
    v4 = append(v4, 1)
    for fi6 := 0; fi6 < 3; fi6++ {
        for fi7 := 0; fi7 < 11; fi7++ {
            v8 := GenS1{F0: 0, F1: "\n\t", F2: v5}
            _ = fi7
            _ = v8
        }
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        var v9 int32 = 598
        _ = fi6
        _ = v9
    }
    _ = v4
    _ = v5
    return 65535
}
func genFn1(p0 uint64) int8 {
    var v10 int64 = -1
    if ((-108 == 127) && false) {
        v11 := GenS1{F0: p0, F1: "\n\t", F2: -1}
        switch "hello" {
        case "hello":
            if true {
                v12 := []int{9223372036854775807, -580, 42, 9223372036854775807}
                _ = v12
            }
        case "\n\t":
            v13 := GenS1{F0: p0, F1: ("\n\t" + ""), F2: -1}
            _ = v13
        default:
            for fi14 := 0; fi14 < 3; fi14++ {
                var v15 uint = 0
                _ = fi14
                _ = v15
            }
        }
        _ = v11
    }
    switch 32768 {
    case 127:
        switch 1 {
        case 233:
            if false {
                p0 -= 42
            } else {
                fmt.Println(v10, v10)
                var v16 string = ("" + ("foo bar" + "a"))
                fmt.Println(v16, v16, v10)
                _ = v16
            }
        default:
            v10 -= 9223372036854775807
        }
    case 519:
        fmt.Println(v10, p0)
    case 777:
        for fi17 := 0; fi17 < 12; fi17++ {
            for fi18 := 0; fi18 < 1; fi18++ {
                v19 := GenS1{F0: 4294967295, F1: (("hello" + "") + ("" + (("\n\t" + "") + "hello"))), F2: 42}
                _ = fi18
                _ = v19
            }
            _ = fi17
        }
    default:
        var v20 string = "a"
        _ = v20
    }
    p0 -= p0
    var v21 uint = 255
    _ = p0
    _ = v10
    _ = v21
    return 42
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[string]uint32)
    v3 := []string{"a", "foo bar"}
    v22 := GenS1{F0: 1, F1: (("foo bar" + "foo bar") + "a"), F2: 0}
    var v23 float64 = ((-2.25 / 1e10) / 1.5)
    v23 -= v23
    v23 = v23
    fmt.Println(v23)
    for fi24 := 0; fi24 < 11; fi24++ {
        var v25 uint8 = 42
        v26 := GenS1{F0: 847, F1: "a", F2: 42}
        var v27 bool = true
        _ = fi24
        _ = v25
        _ = v26
        _ = v27
    }
    fmt.Println(v23)
    v28 := GenS1{F0: 437, F1: "\n\t", F2: -76}
    if (!(127 != -128)) {
        var v29 int64 = -709
        v23 = v23
        v23 = ((0.1 / ((1e10 / float64(float32((v23 + ((-2.25 * float64((1.5 * -2.25))) * 0.1))))) + float64((1.5 + 0.1)))) / v23)
        _ = v29
    }
    v23 -= (v23 / 1e10)
    v23 -= ((-2.25 - v23) + (v23 + 1e10))
    for fi30 := 0; fi30 < 7; fi30++ {
        v23 = 1.5
        v23 = 0.0
        v23 -= 3.14159
        _ = fi30
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v28
    _ = v3
}
