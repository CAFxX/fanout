package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int }

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[uint32]uint32)
    v3 := []string{("hello" + ("foo bar" + "\n\t"))}
    for fi4 := 0; fi4 < 2; fi4++ {
        if (!(1 > 7)) {
            var v5 int = (^(fi4))
            v6 := GenS1{F0: 255, F1: (v5 * fi4)}
            _ = v5
            _ = v6
        }
        _ = fi4
    }
    v7 := GenS1{F0: 65535, F1: 70}
    var v8 uint64 = 42
    var v9 float32 = -2.25
    if ((!true) && true) {
        var v10 string = ("foo bar" + "foo bar")
        for fi11 := 0; fi11 < 17; fi11++ {
            v10 = v10
            if (v9 < (0.0 / -2.25)) {
                if (false && false) {
                    v9 -= ((((1e10 + v9) + ((1.5 + (float32(1e10) * float32(1e10))) * float32(3.14159))) + ((v9 * (((0.1 - v9) + ((v9 + 3.14159) + v9)) * (float32(1.5) - 0.0))) - 1.5)) * v9)
                    fmt.Println(v10, v10, v8)
                } else {
                    v12 := v2[42]
                    _ = len(v2)
                    _ = v12
                }
            } else {
                for fi13 := 0; fi13 < 1; fi13++ {
                    if (!(((-9223372036854775808 > 1) && ((!(false && true)) || false)) && false)) {
                        v9 += v9
                    }
                    _ = fi13
                }
                var v14 float64 = 0.0
                _ = v14
            }
            if true {
                var v15 uint32 = 156
                v8 = v8
                _ = v15
            }
            _ = fi11
        }
        _ = v10
    }
    if len(v3) > 0 {
        v3[0] = "hello"
    }
    v9 = 1.5
    v9 = (float32(1.5) * v9)
    var v16 float64 = 1e10
    _ = len(v3)
    if (715 > 138) {
        var v17 float32 = (v9 * ((v9 - v9) - ((3.14159 + ((float32((v16 - 1e10)) * ((0.1 / float32(float64(v9))) * (v9 * (3.14159 / (v9 + (v9 / v9)))))) - v9)) + 0.1)))
        if (25 == -1) {
            var v18 uint16 = 32767
            var v19 uint8 = 42
            switch 9223372036854775807 {
            case 724:
                v20 := GenS1{F0: 18446744073709551615, F1: 2147483647}
                _ = v20
            }
            _ = v18
            _ = v19
        } else {
            if len(v3) > 0 {
                v3[0] = "hello"
            }
            switch -570 {
            case 552:
                v21 := GenS1{F0: 661, F1: -9223372036854775808}
                _ = v21
            case 9223372036854775807:
                var v22 int8 = 42
                _ = v22
            case 223:
                var v23 int8 = -128
                _ = v23
            }
            fmt.Println(v16, v9)
        }
        switch 9223372036854775807 {
        case 128:
            var v24 int64 = -567
            _ = v24
        }
        _ = v17
    } else {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        for fi25 := 0; fi25 < 4; fi25++ {
            var v26 float64 = v16
            _ = fi25
            _ = v26
        }
    }
    fmt.Println(v16, v8)
    for fi27 := 0; fi27 < 18; fi27++ {
        var v28 int8 = -128
        for fi29 := 0; fi29 < 12; fi29++ {
            switch "a" {
            case "a":
                fmt.Println(v8, fi29, fi27)
            case "hello":
                v1[(fi29 ^ (fi27 % ((fi29) | 1)))] = true
            case "\n\t":
                var v30 int8 = (v28 * v28)
                _ = v30
            }
            _ = fi29
        }
        _ = fi27
        _ = v28
    }
    for fi31 := 0; fi31 < 2; fi31++ {
        _ = len(v1)
        var v32 int8 = 74
        var v33 float64 = v16
        _ = fi31
        _ = v32
        _ = v33
    }
    fmt.Println(v9, v9, v9)
    v8 = v8
    switch "\n\t" {
    case "hello":
        v9 = (v9 * (v9 - (v9 - v9)))
    default:
        v9 = v9
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
