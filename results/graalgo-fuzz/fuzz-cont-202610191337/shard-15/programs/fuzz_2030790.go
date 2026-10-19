package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint; F2 uint8 }
type GenS2 struct { F0 bool; F1 uint64 }
type GenS3 struct { F0 uint64; F1 uint16 }
type GenS4 struct { F0 uint16; F1 bool; F2 uint16 }
func genFn0(p0 uint64, p1 uint16) uint64 {
    v4 := GenS1{F0: 0, F1: 2, F2: 128}
    v5 := GenS1{F0: 32767, F1: 598, F2: 42}
    fmt.Println(p1)
    for fi6 := 0; fi6 < 15; fi6++ {
        for fi7 := 0; fi7 < 16; fi7++ {
            fmt.Println(p0)
            _ = fi7
        }
        _ = fi6
    }
    p0 = 0
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return (p0 | uint64(p1))
}
func genFn1(p0 uint8, p1 string) int {
    var v8 int16 = 211
    p1 = ("a" + p1)
    _ = p0
    _ = p1
    _ = v8
    return -2147483648
}

func main() {
    v1 := []int{307, 2147483647, -268, -2147483648}
    var v2 int32 = -46
    v3 := []uint32{}
    var v9 bool = false
    switch "" {
    case "foo bar":
        if (234 != 4294967295) {
            fmt.Println(v9, v2)
            if len(v3) > 0 {
                v3[0] = 962
            }
        }
    case "a":
        fmt.Println(v2, v9)
    }
    v9 = (1 >= 128)
    if (((0.0 / -2.25) / (1e10 - -2.25)) < 1.5) {
        fmt.Println(v9, v2)
        var v10 int8 = -128
        _ = v10
    } else {
        switch 536 {
        case -408:
            if (166 >= 2) {
                if (((v2 / ((1) | 1)) != v2) && true) {
                    for fi11 := 0; fi11 < 5; fi11++ {
                        var v12 uint16 = 255
                        v13 := GenS1{F0: genFn1(4, ("" + "hello")), F1: 42, F2: 2}
                        v14 := GenS1{F0: 32767, F1: 538, F2: 127}
                        _ = fi11
                        _ = v12
                        _ = v13
                        _ = v14
                    }
                }
                for fi15 := 0; fi15 < 16; fi15++ {
                    fmt.Println(v9)
                    _ = fi15
                }
                fmt.Println(v9)
            }
        case 128:
            for fi16 := 0; fi16 < 3; fi16++ {
                _ = len(v1)
                _ = fi16
            }
        default:
            if ("" <= "hello") {
                v17 := GenS1{F0: 32767, F1: 361, F2: 176}
                _ = v17
            } else {
                v9 = (2 > 2147483648)
            }
        }
    }
    for fi18 := 0; fi18 < 10; fi18++ {
        var v19 int = 0
        var v20 int = 4294967295
        _ = fi18
        _ = v19
        _ = v20
    }
    v2 = v2
    v21 := GenS1{F0: genFn1(1, "\n\t"), F1: 194, F2: 0}
    var v22 int32 = v2
    if (95 >= 227) {
        if (!v9) {
            var v23 uint8 = 152
            switch genFn1((v23 >> ((830) & 63)), "") {
            case -159:
                fmt.Println(v2, v9)
            case 2147483647:
                if v9 {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                }
            case 0:
                var v24 int8 = 100
                _ = v24
            default:
                fmt.Println(v2)
            }
            _ = v23
        }
        for fi25 := 0; fi25 < 8; fi25++ {
            var v26 uint = 255
            _ = fi25
            _ = v26
        }
    } else {
        v27 := GenS2{F0: true, F1: 2}
        var v28 uint8 = 2
        if (2 > 2) {
            v3 = append(v3, 510)
            switch "hello" {
            case "\n\t":
                v29 := GenS3{F0: 2, F1: 283}
                _ = v29
            case "hello":
                _ = len(v3)
            case "":
                v2 -= (v2 << ((18446744073709551615) & 63))
            default:
                if v9 {
                    fmt.Println(v22, v22, v9)
                } else {
                    fmt.Println(v28)
                    switch ("foo bar" + "hello") {
                    case "a":
                        if len(v3) > 0 {
                            v3[0] = 2147483647
                        }
                    case "":
                        fmt.Println(v22, v9, v28)
                    case "\n\t":
                        if true {
                            v22 = (v22 / (((v22 * (v2 ^ 0))) | 1))
                            if v9 {
                                switch 2147483647 {
                                case -7:
                                    fmt.Println(v9, v28, v2)
                                default:
                                    if v9 {
                                        var v30 uint32 = 1
                                        _ = v30
                                    } else {
                                        switch 127 {
                                        case 42:
                                            var v31 bool = v9
                                            _ = v31
                                        default:
                                            for fi32 := 0; fi32 < 7; fi32++ {
                                                fmt.Println(fi32, v9)
                                                _ = fi32
                                            }
                                        }
                                    }
                                }
                                v33 := GenS1{F0: -776, F1: 971, F2: 128}
                                var v34 float64 = (0.0 + 0.0)
                                _ = v33
                                _ = v34
                            } else {
                                for fi35 := 0; fi35 < 8; fi35++ {
                                    var v36 int8 = 42
                                    _ = len(v3)
                                    v22 = 2147483647
                                    _ = fi35
                                    _ = v36
                                }
                                v3 = append(v3, 840)
                                v28 += v28
                            }
                            var v37 int8 = -128
                            _ = v37
                        }
                    default:
                        for fi38 := 0; fi38 < 17; fi38++ {
                            v2 = (v2 * 1)
                            var v39 string = ("hello" + ("" + "hello"))
                            _ = fi38
                            _ = v39
                        }
                    }
                    var v40 int16 = 42
                    _ = v40
                }
            }
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
        _ = v27
        _ = v28
    }
    var v41 int = genFn1(128, (("hello" + "\n\t") + ""))
    switch v41 {
    case -1:
        for fi42 := 0; fi42 < 5; fi42++ {
            if v9 {
                if len(v1) > 0 {
                    v1[0] = 1
                }
                var v43 uint16 = 42
                v43 = (v43 & (v43 / ((v43) | 1)))
                _ = v43
            } else {
                switch "a" {
                case "hello":
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                case "a":
                    var v44 float64 = (3.14159 + 1.5)
                    _ = v44
                default:
                    v3 = append(v3, 42)
                }
                v9 = (-2.25 == float64(((((1e10 * (3.14159 + (-2.25 / float32(1.5)))) + (3.14159 - 3.14159)) * (0.0 + float32(-2.25))) - (-2.25 / float32((1.5 / 1.0))))))
            }
            v45 := GenS4{F0: 65535, F1: (!true), F2: 65535}
            _ = fi42
            _ = v45
        }
    case 545:
        if (v9 || v9) {
            switch ("" + ((("" + ("a" + "foo bar")) + "\n\t") + "hello")) {
            case "":
                for fi46 := 0; fi46 < 20; fi46++ {
                    var v47 uint8 = 2
                    v3 = append(v3, 2147483648)
                    var v48 int8 = 100
                    _ = fi46
                    _ = v47
                    _ = v48
                }
            case "hello":
                var v49 int64 = 9223372036854775807
                _ = v49
            }
        } else {
            v41 += (v41 - (v41 + v41))
            if len(v1) > 0 {
                v1[0] = (v41 % (((v41 + genFn1(128, ("" + "a")))) | 1))
            }
            var v50 string = "a"
            _ = v50
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v41:", v41)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v41
    _ = v9
}
