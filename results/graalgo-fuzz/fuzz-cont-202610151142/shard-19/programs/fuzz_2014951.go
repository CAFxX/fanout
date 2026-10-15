package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64 }
type GenS2 struct { F0 int; F1 int16 }

func main() {
    v1 := []string{"a", "\n\t", "\n\t"}
    var v2 uint32 = 42
    var v3 uint16 = 763
    if false {
        var v4 int32 = -2147483648
        _ = v4
    } else {
        for fi5 := 0; fi5 < 2; fi5++ {
            if (0 != 42) {
                for fi6 := 0; fi6 < 15; fi6++ {
                    if (127 <= 157) {
                        v7 := GenS1{F0: (!false), F1: -939}
                        if (162 > 460) {
                            if false {
                                if (true && true) {
                                    var v8 int64 = -9223372036854775808
                                    var v9 float64 = ((0.1 + 1.5) / 1.5)
                                    v10 := GenS1{F0: (1 != 2), F1: 9223372036854775807}
                                    _ = v10
                                    _ = v8
                                    _ = v9
                                }
                            }
                        }
                        _ = len(v1)
                        _ = v7
                    }
                    if (-12 != -45) {
                        v1 = append(v1, "\n\t")
                    }
                    _ = fi6
                }
            }
            v2 += v2
            _ = fi5
        }
        if false {
            _ = len(v1)
            v2 -= 2147483648
        }
        v11 := GenS2{F0: -321, F1: 42}
        _ = v11
    }
    v3 = v3
    fmt.Println(v3, v3)
    var v12 int8 = -128
    fmt.Println(v12, v3, v3)
    if (-989 <= -9223372036854775808) {
        fmt.Println(v12)
    } else {
        fmt.Println(v12)
    }
    if false {
        if (42 <= 18446744073709551615) {
            v13 := GenS2{F0: 1, F1: -105}
            v14 := GenS1{F0: false, F1: 521}
            var v15 uint = 4294967295
            _ = v13
            _ = v14
            _ = v15
        } else {
            v16 := GenS1{F0: false, F1: 42}
            fmt.Println(v3, v12)
            _ = v16
        }
    }
    for fi17 := 0; fi17 < 1; fi17++ {
        fmt.Println(fi17, fi17)
        _ = fi17
    }
    v2 -= (v2 - v2)
    v12 -= v12
    var v18 int8 = v12
    switch ("\n\t" + "\n\t") {
    case "":
        _ = len(v1)
    case "hello":
        switch -222 {
        case -1:
            for fi19 := 0; fi19 < 18; fi19++ {
                var v20 int8 = v18
                var v21 bool = true
                _ = fi19
                _ = v20
                _ = v21
            }
        default:
            fmt.Println(v18, v2, v2)
        }
    case "a":
        var v22 uint = 749
        _ = v22
    }
    if (v12 == v12) {
        var v23 float64 = ((float64(1.5) * (((0.0 / ((0.0 + -2.25) / 3.14159)) + ((-2.25 * 1.5) - (1.5 / 1.5))) / (((float64(1e10) - ((((-2.25 + (0.1 + 0.0)) - float64(0.1)) * 1.5) + 1e10)) + -2.25) - -2.25))) + ((0.1 - 1e10) * 3.14159))
        _ = v23
    }
    var v24 bool = (225 <= (v2 * (v2 * v2)))
    var v25 uint16 = (v3 % (((v3 << ((255) & 63))) | 1))
    for fi26 := 0; fi26 < 10; fi26++ {
        v1 = append(v1, ("foo bar" + "a"))
        _ = fi26
    }
    var v27 uint8 = 232
    var v28 uint16 = v3
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v18
    _ = v2
    _ = v24
    _ = v25
    _ = v27
    _ = v28
    _ = v3
}
