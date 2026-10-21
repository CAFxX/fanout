package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 bool; F1 uint32 }
type GenS3 struct { F0 uint8; F1 string }
type GenS4 struct { F0 uint8; F1 uint32; F2 string }
func genFn0() int64 {
    v4 := GenS1{F0: 2147483647}
    var v5 uint = 0
    v5 = (v5 / ((v5) | 1))
    _ = v4
    _ = v5
    return 42
}

func main() {
    var v1 float32 = ((((1.5 / (3.14159 * 0.1)) - 0.1) / 1e10) / ((1.5 * 0.1) / 1.0))
    var v2 uint64 = 243
    v3 := make(map[string]int)
    switch 439 {
    case 238:
        if (-2.25 >= 3.14159) {
            for fi6 := 0; fi6 < 9; fi6++ {
                var v7 uint8 = 255
                var v8 float64 = 1.5
                var v9 int32 = -600
                _ = fi6
                _ = v7
                _ = v8
                _ = v9
            }
            var v10 bool = ((true && (2147483647 != 9223372036854775807)) || (!((^(v2)) <= v2)))
            v10 = (!false)
            _ = v10
        }
    default:
        _ = len(v3)
    }
    var v11 int64 = 1
    var v12 float64 = (1e10 * ((1.5 - 1.5) * 0.1))
    if (783 <= 1) {
        v13 := []int{42, 127, -9223372036854775808, -1}
        _ = v13
    } else {
        v14 := GenS2{F0: (!(true || true)), F1: 2147483647}
        _ = v14
    }
    var v15 uint = 0
    switch -9223372036854775808 {
    case 9223372036854775807:
        fmt.Println(v12)
    case 885:
        fmt.Println(v2)
    case 2147483647:
        v11 -= v11
    default:
        v16 := GenS3{F0: 2, F1: ("foo bar" + (("a" + "") + (("\n\t" + "foo bar") + "a")))}
        _ = v16
    }
    switch ("a" + (("a" + "\n\t") + "a")) {
    case "hello":
        if ((v12 != v12) || (uint16(v15) > 1)) {
            if true {
                var v17 string = "foo bar"
                if ((42 <= 842) || false) {
                    var v18 uint16 = 65535
                    switch v17 {
                    case "":
                        v11 = v11
                    default:
                        v19 := GenS4{F0: 0, F1: 65535, F2: "hello"}
                        _ = v19
                    }
                    _ = v18
                } else {
                    for fi20 := 0; fi20 < 4; fi20++ {
                        var v21 int32 = -1
                        _ = fi20
                        _ = v21
                    }
                    v22 := []int{1, 810}
                    _ = v22
                }
                _ = v17
            } else {
                v3["a"] = 32768
            }
            v2 = v2
            var v23 uint = 455
            _ = v23
        }
    case "a":
        switch 4294967295 {
        case 0:
            for fi24 := 0; fi24 < 2; fi24++ {
                var v25 string = "hello"
                _ = fi24
                _ = v25
            }
        default:
            for fi26 := 0; fi26 < 16; fi26++ {
                var v27 int8 = -1
                _ = fi26
                _ = v27
            }
        }
    case "foo bar":
        if ((true || (((1 <= -2147483648) && false) && (false || (32767 <= 32767)))) || true) {
            if false {
                v28 := GenS1{F0: -459}
                var v29 uint8 = 17
                _ = v28
                _ = v29
            }
            v30 := GenS2{F0: ((float64(v1) - 0.0) != 0.1), F1: 2147483648}
            _ = v30
        } else {
            v1 = 1.5
        }
    }
    v12 = v12
    v31 := GenS1{F0: -1}
    var v32 float32 = 0.0
    if ((v15 % ((v15) | 1)) > v15) {
        for fi33 := 0; fi33 < 13; fi33++ {
            v3["\n\t"] = (fi33 << ((v15) & 63))
            var v34 float32 = float32(v12)
            _ = fi33
            _ = v34
        }
    }
    v35 := []int{-810, -2147483648, int(v11)}
    v15 = v15
    switch (("a" + "hello") + "foo bar") {
    case "a":
        v2 += v2
    case "\n\t":
        var v36 int16 = 42
        _ = v36
    case "":
        v3["foo bar"] = 9223372036854775807
    default:
        v12 = (v12 * ((v12 - v12) / ((((v12 * ((v12 * 0.1) * (v12 + (v12 + ((v12 + v12) / v12))))) + v12) / v12) * v12)))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v35:", len(v35), v35)
    _ = v1
    _ = v11
    _ = v12
    _ = v15
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v35
}
