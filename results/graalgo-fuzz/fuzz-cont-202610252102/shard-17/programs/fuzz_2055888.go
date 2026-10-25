package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16 }
type GenS2 struct { F0 int32; F1 int32; F2 uint }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint; F1 uint32; F2 uint8 }
type GenS6 struct { F0 int8; F1 uint32 }

func main() {
    var v1 uint = 2
    var v2 uint64 = 42
    var v3 uint = 65535
    var v4 uint = (v3 + (v1 * (v1 + (v1 + v1))))
    fmt.Println(v3)
    var v5 float32 = 3.14159
    v3 += 150
    var v6 float32 = (v5 * v5)
    v7 := GenS1{F0: 0, F1: 1}
    v8 := []int{-2147483648, 32767, 4294967295, -230}
    if len(v8) >= 2 {
        v8 = v8[1:len(v8)]
    }
    if (false && true) {
        v2 -= v2
        fmt.Println(v1)
    } else {
        fmt.Println(v3, v6)
        v3 -= v1
        var v9 uint = 725
        _ = v9
    }
    fmt.Println(v4, v5, v1)
    if (v5 == (v6 * float32(0.1))) {
        if (164 != -9223372036854775808) {
            v2 -= (v2 | (v2 + 18446744073709551615))
        } else {
            var v10 uint = (v3 + v4)
            switch 0 {
            case -355:
                var v11 int16 = 0
                _ = v11
            }
            _ = v10
        }
    }
    v4 -= v4
    if false {
        v2 += v2
        if (!true) {
            if ((((-1 > 100) && ((false && false) || ((((v5 / (v5 / v5)) / (float32(((0.1 * 1e10) * (0.0 + 1.5))) + v5)) + (1.5 / (float32((1e10 / 3.14159)) + ((v6 + ((v5 - ((v6 * (-2.25 + -2.25)) / float32(float64(3.14159)))) * ((v5 - (v5 / 3.14159)) * (((v6 - 1.5) * float32(((float64(v5) / 1e10) + (0.0 / 1e10)))) - (float32(1.5) / 1.5))))) - v5)))) > v6))) || (0 != 74)) && true) {
                if true {
                    if (true && (v2 <= 1)) {
                        v8 = append(v8, 32767)
                        v2 += v2
                    } else {
                        fmt.Println(v2, v6)
                        v2 -= 4294967295
                    }
                    var v12 float64 = 0.1
                    var v13 bool = false
                    _ = v12
                    _ = v13
                } else {
                    var v14 bool = true
                    _ = v14
                }
                if false {
                    fmt.Println(v5, v3)
                    switch 2147483647 {
                    case -879:
                        if true {
                            switch "\n\t" {
                            case "a":
                                switch "hello" {
                                case "a":
                                    var v15 uint8 = 1
                                    _ = v15
                                case "\n\t":
                                    var v16 uint32 = 689
                                    _ = v16
                                case "hello":
                                    for fi17 := 0; fi17 < 14; fi17++ {
                                        v18 := GenS2{F0: 2147483647, F1: -2147483648, F2: 65535}
                                        _ = fi17
                                        _ = v18
                                    }
                                default:
                                    var v19 int8 = 0
                                    _ = v19
                                }
                            case "foo bar":
                                var v20 int32 = -100
                                _ = v20
                            case "":
                                if (1 >= 255) {
                                    for fi21 := 0; fi21 < 7; fi21++ {
                                        for fi22 := 0; fi22 < 9; fi22++ {
                                            v5 -= (v5 * v5)
                                            v23 := GenS1{F0: 144, F1: 404}
                                            var v24 uint = 4294967295
                                            _ = len(v8)
                                            var v25 int8 = 83
                                            var v26 int8 = v25
                                            _ = fi22
                                            _ = v23
                                            _ = v24
                                            _ = v25
                                            _ = v26
                                        }
                                        _ = fi21
                                    }
                                    v27 := GenS1{F0: 76, F1: 1}
                                    _ = v27
                                } else {
                                    v5 -= float32(1e10)
                                    _ = len(v8)
                                    fmt.Println(v2, v6, v5)
                                }
                            default:
                                v5 = 0.0
                            }
                        } else {
                            for fi28 := 0; fi28 < 15; fi28++ {
                                fmt.Println(v4, v2, v1)
                                v29 := GenS1{F0: 44, F1: 32767}
                                _ = fi28
                                _ = v29
                            }
                            for fi30 := 0; fi30 < 11; fi30++ {
                                _ = len(v8)
                                _ = fi30
                            }
                        }
                    case -106:
                        for fi31 := 0; fi31 < 3; fi31++ {
                            if (42 >= -542) {
                                var v32 int8 = 5
                                _ = v32
                            }
                            v1 = v3
                            _ = fi31
                        }
                    case 0:
                        v3 += (v4 - 42)
                    default:
                        for fi33 := 0; fi33 < 5; fi33++ {
                            switch (fi33 * fi33) {
                            case -1:
                                v8 = append(v8, fi33)
                            case -9223372036854775808:
                                v6 -= (1e10 + v5)
                            default:
                                v34 := GenS2{F0: -921, F1: -857, F2: v4}
                                _ = v34
                            }
                            switch ("a" + "a") {
                            case "foo bar":
                                switch "foo bar" {
                                case "\n\t":
                                    switch fi33 {
                                    case -2147483648:
                                        switch "" {
                                        case "":
                                            var v35 uint = (^(v4))
                                            _ = v35
                                        case "a":
                                            if len(v8) > 0 {
                                                v8[0] = (fi33 & fi33)
                                            }
                                        default:
                                            switch ("\n\t" + ("" + "hello")) {
                                            case "hello":
                                                v36 := GenS2{F0: -2147483648, F1: -2147483648, F2: uint(v2)}
                                                _ = v36
                                            default:
                                                v2 = (v2 * (v2 / ((v2) | 1)))
                                            }
                                        }
                                    case 42:
                                        if ((0 == 606) && (1000 > -679)) {
                                            if true {
                                                v8 = append(v8, (fi33 ^ fi33))
                                                switch (("hello" + "\n\t") + "hello") {
                                                case "hello":
                                                    var v37 float32 = (3.14159 - v6)
                                                    _ = v37
                                                case "a":
                                                    fmt.Println(fi33, v4, v3)
                                                }
                                                if len(v8) >= 2 {
                                                    v8 = v8[1:len(v8)]
                                                }
                                            } else {
                                                v1 = v1
                                            }
                                        }
                                    case 2147483647:
                                        switch "foo bar" {
                                        case "":
                                            v1 = (v3 >> (((v1 * (v3 % ((v3) | 1)))) & 63))
                                        case "a":
                                            v4 = v4
                                        case "hello":
                                            fmt.Println(v4, v2)
                                        }
                                    }
                                case "hello":
                                    v1 += v1
                                case "a":
                                    v38 := GenS3{F0: -1}
                                    _ = v38
                                default:
                                    if len(v8) >= 2 {
                                        v8 = v8[1:len(v8)]
                                    }
                                }
                            default:
                                _ = len(v8)
                            }
                            fmt.Println(v2, v6, v1)
                            _ = fi33
                        }
                    }
                    v39 := GenS1{F0: 42, F1: 255}
                    _ = v39
                } else {
                    var v40 uint64 = (v2 | (^(v2)))
                    _ = v40
                }
                v4 += 2
            }
            v2 -= (^(v2))
        } else {
            var v41 float32 = (float32(1.5) - float32(3.14159))
            if false {
                switch 32768 {
                case -31:
                    if true {
                        v4 += (v3 * v1)
                        _ = len(v8)
                    }
                case 42:
                    _ = len(v8)
                case -789:
                    switch ("hello" + "") {
                    case "":
                        if (32768 != -2147483648) {
                            var v42 float64 = 1.5
                            if (2 >= 65535) {
                                v5 -= -2.25
                                _ = len(v8)
                            } else {
                                v8 = append(v8, 2147483647)
                            }
                            _ = v42
                        }
                    case "foo bar":
                        v5 -= v5
                    default:
                        if false {
                            switch ("hello" + ("\n\t" + "hello")) {
                            case "":
                                if true {
                                    v1 += (v1 - v3)
                                }
                            case "\n\t":
                                v2 = (v2 | v2)
                            case "foo bar":
                                if len(v8) > 0 {
                                    v8[0] = 32768
                                }
                            }
                            var v43 int64 = 42
                            _ = v43
                        }
                    }
                default:
                    _ = len(v8)
                }
                if (-32768 == 1) {
                    for fi44 := 0; fi44 < 4; fi44++ {
                        var v45 int8 = -107
                        _ = fi44
                        _ = v45
                    }
                } else {
                    var v46 int32 = -2147483648
                    v47 := GenS4{F0: 42}
                    var v48 string = ("hello" + "a")
                    _ = v46
                    _ = v47
                    _ = v48
                }
                v3 += v1
            } else {
                v5 = (v5 * 0.1)
                switch 32768 {
                case 1:
                    if len(v8) >= 2 {
                        v8 = v8[1:len(v8)]
                    }
                case -521:
                    var v49 bool = (false && true)
                    _ = v49
                }
                fmt.Println(v41, v5)
            }
            _ = v41
        }
        _ = len(v8)
    }
    v50 := GenS5{F0: (v1 * 255), F1: 42, F2: 127}
    v51 := GenS6{F0: 100, F1: 4294967295}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v51
    _ = v6
    _ = v7
    _ = v8
}
