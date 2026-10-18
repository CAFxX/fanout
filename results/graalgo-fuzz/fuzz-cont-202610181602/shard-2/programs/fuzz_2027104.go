package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint }
type GenS2 struct { F0 int32; F1 uint32 }
func genFn0(p0 bool) uint8 {
    v7 := []int{32767, 127, -222}
    var v8 string = ""
    for fi9 := 0; fi9 < 2; fi9++ {
        var v10 float32 = 3.14159
        _ = len(v7)
        for fi11 := 0; fi11 < 3; fi11++ {
            var v12 string = (v8 + (("foo bar" + "foo bar") + "a"))
            var v13 float64 = 1.5
            v14 := GenS1{F0: 64, F1: 0}
            _ = fi11
            _ = v12
            _ = v13
            _ = v14
        }
        _ = fi9
        _ = v10
    }
    switch "foo bar" {
    case "\n\t":
        v8 = (v8 + (v8 + "hello"))
    case "a":
        if false {
            var v15 uint16 = 58
            _ = v15
        }
    case "hello":
        v8 = ("" + "")
    default:
        switch v8 {
        case "hello":
            var v16 uint8 = 197
            _ = v16
        case "a":
            var v17 uint64 = 2
            _ = v17
        }
    }
    _ = p0
    _ = v7
    _ = v8
    return 112
}
func genFn1() uint8 {
    switch "foo bar" {
    case "\n\t":
        var v18 uint8 = 124
        _ = v18
    case "":
        var v19 bool = (("\n\t" + "") >= "hello")
        _ = v19
    default:
        if (0 >= -185) {
            var v20 int8 = -16
            v21 := GenS1{F0: -1, F1: 42}
            _ = v20
            _ = v21
        }
    }
    v22 := []int{1}
    _ = len(v22)
    _ = v22
    return 128
}

func main() {
    var v1 int8 = 47
    var v2 float32 = 0.1
    var v3 int = 1
    var v4 int16 = -32768
    var v5 uint32 = 42
    var v6 int = -9223372036854775808
    v23 := []int{v6, (v3 % (((v6 * v3)) | 1)), (v3 | v6)}
    fmt.Println(v2)
    _ = len(v23)
    fmt.Println(v2, v5, v6)
    fmt.Println(v2)
    var v24 uint32 = v5
    var v25 int16 = 42
    fmt.Println(v25)
    if (166 == uint8(v5)) {
        var v26 float64 = 0.0
        for fi27 := 0; fi27 < 5; fi27++ {
            switch "" {
            case "\n\t":
                v4 -= -550
            case "foo bar":
                fmt.Println(v1, v25, v1)
            case "hello":
                switch fi27 {
                case -2147483648:
                    var v28 uint = 255
                    _ = v28
                default:
                    for fi29 := 0; fi29 < 13; fi29++ {
                        v6 -= -621
                        if len(v23) > 0 {
                            v23[0] = 0
                        }
                        fmt.Println(v26)
                        _ = fi29
                    }
                }
            }
            _ = fi27
        }
        _ = v26
    }
    switch 42 {
    case 292:
        v30 := GenS2{F0: -2147483648, F1: v5}
        _ = v30
    }
    for fi31 := 0; fi31 < 17; fi31++ {
        if len(v23) >= 2 {
            v23 = v23[1:len(v23)]
        }
        var v32 uint32 = 1
        _ = fi31
        _ = v32
    }
    v5 -= v24
    if (v2 >= (((((float32((0.0 / 3.14159)) * v2) * v2) + 0.0) + (v2 * 0.1)) - v2)) {
        for fi33 := 0; fi33 < 18; fi33++ {
            v34 := GenS2{F0: 42, F1: (v24 % (((v24 % ((1) | 1))) | 1))}
            if len(v23) >= 2 {
                v23 = v23[1:len(v23)]
            }
            switch -2147483648 {
            case -583:
                switch 182 {
                case -842:
                    v3 += v6
                }
            default:
                var v35 uint32 = (v5 ^ v24)
                _ = v35
            }
            _ = fi33
            _ = v34
        }
    } else {
        for fi36 := 0; fi36 < 15; fi36++ {
            for fi37 := 0; fi37 < 9; fi37++ {
                var v38 int8 = (v1 + (^(v1)))
                v39 := GenS2{F0: 42, F1: (v24 * v5)}
                _ = fi37
                _ = v38
                _ = v39
            }
            _ = fi36
        }
        if ((("" + "foo bar") + "hello") < "hello") {
            if ((v6 <= v6) || false) {
                fmt.Println(v3)
                v4 += v25
                v2 = v2
            }
        }
    }
    v1 -= v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
