package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int16; F1 bool; F2 uint64 }
type GenS3 struct { F0 bool; F1 int64; F2 uint64 }
type GenS4 struct { F0 uint64; F1 uint16 }
func genFn0(p0 int16, p1 uint32) uint32 {
    v5 := []int{128, 0, 1}
    var v6 uint32 = p1
    v7 := GenS1{F0: int64(v6)}
    for fi8 := 0; fi8 < 12; fi8++ {
        switch ("hello" + "foo bar") {
        case "a":
            switch "foo bar" {
            case "\n\t":
                if false {
                    if ((((((0.0 - 0.0) + float64(0.0)) * 1e10) / 1e10) <= float64((0.1 - (float32(((((((0.1 * (1.5 / 1.0)) - ((3.14159 * -2.25) + float64(float32(0.1)))) + 1.5) / 0.1) - 3.14159) + (0.1 + 0.0))) - -2.25)))) || false) {
                        v9 := GenS1{F0: -1}
                        p0 = p0
                        var v10 int16 = 32767
                        _ = v10
                        _ = v9
                    }
                }
            case "":
                switch ("\n\t" + (("" + "\n\t") + (("hello" + ("hello" + ("a" + "a"))) + ("foo bar" + "a")))) {
                case "foo bar":
                    v6 = (v6 % ((v6) | 1))
                case "a":
                    for fi11 := 0; fi11 < 1; fi11++ {
                        p1 = p1
                        var v12 float64 = -2.25
                        if true {
                            switch "foo bar" {
                            case "a":
                                v13 := GenS2{F0: 1000, F1: ((^(fi8)) <= 128), F2: 42}
                                _ = v13
                            case "\n\t":
                                if true {
                                    v12 = v12
                                }
                            case "hello":
                                v12 += v12
                            }
                        }
                        _ = fi11
                        _ = v12
                    }
                default:
                    v6 -= 65535
                }
            default:
                var v14 float64 = (((float64(((float32(1e10) - ((0.0 - 0.0) - 0.1)) + 1.5)) + ((float64((1e10 + (0.1 / 3.14159))) / ((((3.14159 / float64(1e10)) - 0.1) + 0.0) / 1.0)) + -2.25)) + 3.14159) * (1.5 + (1.5 / 1e10)))
                _ = v14
            }
        case "foo bar":
            v6 = (p1 * v6)
        }
        _ = fi8
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return 169
}

func main() {
    var v1 int32 = -1
    v2 := []int{-9223372036854775808, 4294967295, 768, 127}
    v3 := make(map[string]int)
    var v4 int64 = -86
    v3[("hello" + ("\n\t" + "a"))] = -671
    var v15 uint64 = 18446744073709551615
    if (uint32(v15) >= 42) {
        fmt.Println(v4)
        for fi16 := 0; fi16 < 19; fi16++ {
            var v17 uint32 = 0
            v1 = v1
            _ = fi16
            _ = v17
        }
        for fi18 := 0; fi18 < 1; fi18++ {
            var v19 uint8 = 127
            _ = fi18
            _ = v19
        }
    }
    v15 = (^(v15))
    var v20 int8 = 127
    for fi21 := 0; fi21 < 9; fi21++ {
        var v22 float64 = 0.1
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = fi21
        _ = v22
    }
    var v23 int = 32768
    var v24 float64 = 3.14159
    if (!(((0.0 / ((1.5 * 0.0) - 3.14159)) / 1e10) >= 1e10)) {
        _ = len(v2)
        v20 -= v20
        v3[("foo bar" + ("\n\t" + ""))] = 1
    } else {
        v25 := GenS3{F0: (18446744073709551615 != 1), F1: -569, F2: v15}
        _ = v25
    }
    switch "foo bar" {
    case "":
        var v26 bool = (1 >= 32767)
        _ = v26
    }
    if (false && ((v1 ^ (v1 | (v1 & (v1 & (v1 * (v1 % (((v1 / ((v1) | 1))) | 1))))))) <= 42)) {
        fmt.Println(v15)
        for fi27 := 0; fi27 < 12; fi27++ {
            if true {
                var v28 int8 = (v20 ^ (v20 - 0))
                v29 := GenS3{F0: false, F1: (v4 % ((v4) | 1)), F2: v15}
                _ = v28
                _ = v29
            }
            _ = len(v3)
            var v30 int8 = (-(v20))
            _ = fi27
            _ = v30
        }
    } else {
        _ = len(v2)
        var v31 uint32 = 65535
        v32 := GenS4{F0: 106, F1: 1}
        _ = v31
        _ = v32
    }
    v1 = 1
    v33 := GenS4{F0: 4294967295, F1: 255}
    v4 += (v4 / (((v4 << ((18446744073709551615) & 63))) | 1))
    if ((("a" + (("" + ("" + "hello")) + ("\n\t" + "foo bar"))) < "") || false) {
        fmt.Println(v24)
        var v34 int16 = 1
        _ = v34
    } else {
        if (v1 < (v1 / ((v1) | 1))) {
            v23 = (v23 / ((v23) | 1))
        }
    }
    v4 = (v4 << ((1) & 63))
    v35 := GenS2{F0: 600, F1: (136 <= 42), F2: (v15 ^ 1)}
    switch v23 {
    case -2147483648:
        for fi36 := 0; fi36 < 12; fi36++ {
            fmt.Println(v1, v23)
            if (("\n\t" >= "\n\t") && ((true && false) && (!((((!(-883 >= -692)) || (float32(float64(0.0)) <= float32(v24))) || false) || (0 <= 976))))) {
                var v37 uint8 = 163
                for fi38 := 0; fi38 < 9; fi38++ {
                    fmt.Println(v37)
                    _ = fi38
                }
                _ = v37
            } else {
                switch "a" {
                case "":
                    var v39 uint16 = 1
                    _ = v39
                case "foo bar":
                    v1 -= -1
                }
            }
            _ = fi36
        }
    case 127:
        v15 += 687
    case 734:
        if (genFn0(1, 42) == 0) {
            v1 += v1
            v1 -= v1
        } else {
            v23 = v23
            for fi40 := 0; fi40 < 7; fi40++ {
                if (false && true) {
                    var v41 uint64 = v15
                    v20 += (v20 + 100)
                    _ = v41
                }
                switch ("a" + "a") {
                case "hello":
                    if true {
                        _ = len(v3)
                    }
                }
                _ = fi40
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v2
    _ = v20
    _ = v23
    _ = v24
    _ = v3
    _ = v33
    _ = v35
    _ = v4
}
