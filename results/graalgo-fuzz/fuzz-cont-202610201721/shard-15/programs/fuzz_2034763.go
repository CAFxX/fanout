package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string; F2 int8 }
type GenS2 struct { F0 int8; F1 int16 }
type GenS3 struct { F0 int32; F1 uint16; F2 int64 }
type GenS4 struct { F0 int32; F1 int32 }
type GenS5 struct { F0 uint; F1 int64; F2 int32 }
type GenS6 struct { F0 int64; F1 int64; F2 int8 }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 uint64; F1 uint16 }
type GenS9 struct { F0 string; F1 uint32 }
type GenS10 struct { F0 bool }
type GenS11 struct { F0 uint16; F1 int16; F2 bool }
type GenS12 struct { F0 uint8; F1 int; F2 uint }

func main() {
    var v1 uint32 = 232
    var v2 float64 = 0.0
    v3 := []string{"a", "\n\t", (((("foo bar" + "a") + "\n\t") + "a") + "\n\t"), "hello"}
    var v4 int8 = 0
    var v5 int16 = 1
    if ((-123 < 32768) && (("foo bar" + "\n\t") != "\n\t")) {
        var v6 int64 = 0
        for fi7 := 0; fi7 < 13; fi7++ {
            var v8 int8 = v4
            for fi9 := 0; fi9 < 19; fi9++ {
                var v10 uint = 2
                _ = fi9
                _ = v10
            }
            if (v2 < float64((((1.5 + 3.14159) * float32(v2)) + ((0.1 * 0.0) * (float32((float64(0.0) / v2)) + 1.5))))) {
                switch (-(fi7)) {
                case 2147483647:
                    if false {
                        if len(v3) > 0 {
                            v3[0] = "\n\t"
                        }
                        if true {
                            v11 := GenS1{F0: -32768, F1: "\n\t", F2: v4}
                            v3 = append(v3, ("foo bar" + "hello"))
                            var v12 bool = ("foo bar" <= "foo bar")
                            _ = v11
                            _ = v12
                        } else {
                            v6 = (v6 | 42)
                            fmt.Println(v2)
                            v13 := GenS2{F0: v8, F1: 351}
                            _ = v13
                        }
                    } else {
                        v6 = 989
                        v5 = v5
                    }
                case 32767:
                    v1 = v1
                default:
                    if true {
                        fmt.Println(v6, fi7, v2)
                    } else {
                        for fi14 := 0; fi14 < 4; fi14++ {
                            _ = len(v3)
                            v4 -= 1
                            v15 := GenS2{F0: (v8 >> ((18446744073709551615) & 63)), F1: -41}
                            _ = fi14
                            _ = v15
                        }
                    }
                }
            } else {
                v16 := GenS3{F0: 2147483647, F1: 255, F2: 985}
                _ = v16
            }
            _ = fi7
            _ = v8
        }
        v4 += v4
        _ = v6
    } else {
        switch -9223372036854775808 {
        case -755:
            v17 := GenS1{F0: 1000, F1: "\n\t", F2: v4}
            _ = v17
        case -327:
            var v18 bool = (((v2 - v2) + ((((v2 + v2) / (3.14159 + -2.25)) + (v2 + ((3.14159 + v2) - ((float64(-2.25) * v2) - v2)))) - 0.1)) < v2)
            _ = v18
        case -1:
            var v19 int64 = int64(v5)
            _ = v19
        default:
            v5 = -32768
        }
        if false {
            if false {
                switch 0 {
                case 757:
                    v2 = v2
                case 762:
                    if false {
                        if (false && (true || false)) {
                            switch "\n\t" {
                            case "a":
                                for fi20 := 0; fi20 < 16; fi20++ {
                                    v21 := GenS1{F0: -32768, F1: "\n\t", F2: (v4 >> ((4294967295) & 63))}
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    _ = fi20
                                    _ = v21
                                }
                            case "hello":
                                v4 = (v4 << ((255) & 63))
                            case "\n\t":
                                v4 -= v4
                            }
                            if (0 > 255) {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                if ((true && (false && false)) && true) {
                                    if len(v3) > 0 {
                                        v3[0] = ((("" + ("hello" + "a")) + "") + ("a" + "\n\t"))
                                    }
                                } else {
                                    var v22 int16 = v5
                                    _ = v22
                                }
                                var v23 float64 = 1e10
                                _ = v23
                            }
                            if ((!(!(v1 > 684))) || false) {
                                v5 = -899
                                var v24 uint = 505
                                var v25 float64 = ((v2 / 1e10) / (((v2 + ((float64(-2.25) / 3.14159) + (v2 / v2))) - (0.0 + v2)) - (v2 + 0.0)))
                                _ = v24
                                _ = v25
                            }
                        } else {
                            fmt.Println(v4, v5)
                        }
                        var v26 uint16 = 548
                        fmt.Println(v2)
                        _ = v26
                    }
                }
            }
        } else {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
        var v27 float32 = ((3.14159 / 0.1) / 1.0)
        _ = v27
    }
    switch 75 {
    case -2147483648:
        if (true && false) {
            var v28 uint64 = 4294967295
            _ = v28
        } else {
            for fi29 := 0; fi29 < 12; fi29++ {
                for fi30 := 0; fi30 < 20; fi30++ {
                    v31 := GenS1{F0: (v5 << ((1) & 63)), F1: "foo bar", F2: v4}
                    var v32 int = (fi30 >> ((2) & 63))
                    _ = fi30
                    _ = v31
                    _ = v32
                }
                v2 += float64((float32(0.1) - -2.25))
                v33 := GenS4{F0: -2147483648, F1: 1}
                _ = fi29
                _ = v33
            }
            var v34 float32 = 1e10
            _ = v34
        }
    case -426:
        if true {
            if ((65535 < 32767) || false) {
                var v35 string = "hello"
                fmt.Println(v1)
                v36 := GenS5{F0: 255, F1: 423, F2: 0}
                _ = v35
                _ = v36
            } else {
                switch 0 {
                case 690:
                    if (!(true || true)) {
                        var v37 uint = 18446744073709551615
                        if len(v3) > 0 {
                            v3[0] = ""
                        }
                        v2 -= ((-2.25 - (v2 * 0.0)) + float64(1e10))
                        _ = v37
                    }
                case 32767:
                    v3 = append(v3, "\n\t")
                case 127:
                    for fi38 := 0; fi38 < 11; fi38++ {
                        var v39 int32 = -2147483648
                        _ = fi38
                        _ = v39
                    }
                }
            }
        } else {
            if (!true) {
                var v40 int16 = 1
                _ = v40
            }
            if (4294967295 < 4294967295) {
                var v41 int16 = 378
                for fi42 := 0; fi42 < 13; fi42++ {
                    fmt.Println(v41, v41)
                    v3 = append(v3, (("\n\t" + ("" + "")) + "a"))
                    for fi43 := 0; fi43 < 11; fi43++ {
                        fmt.Println(v5)
                        var v44 float32 = float32(1e10)
                        _ = fi43
                        _ = v44
                    }
                    _ = fi42
                }
                _ = v41
            }
            var v45 uint = 18446744073709551615
            _ = v45
        }
    }
    switch -865 {
    case 0:
        v3 = append(v3, ("\n\t" + ""))
    default:
        v3 = append(v3, "hello")
    }
    fmt.Println(v1, v2, v5)
    v46 := GenS6{F0: -9223372036854775808, F1: -1, F2: v4}
    v1 += (v1 << ((434) & 63))
    v47 := GenS2{F0: (^(v4)), F1: (v5 - v5)}
    switch -2147483648 {
    case 32767:
        for fi48 := 0; fi48 < 5; fi48++ {
            switch 1 {
            case 32767:
                if false {
                    v49 := GenS2{F0: v4, F1: v5}
                    for fi50 := 0; fi50 < 16; fi50++ {
                        v51 := GenS6{F0: 0, F1: 9223372036854775807, F2: int8(fi48)}
                        v52 := GenS4{F0: 1, F1: 42}
                        _ = fi50
                        _ = v51
                        _ = v52
                    }
                    _ = v49
                } else {
                    var v53 int8 = v4
                    _ = v53
                }
            case 1:
                _ = len(v3)
            }
            v2 -= 1e10
            _ = fi48
        }
    default:
        v54 := GenS7{F0: v4}
        _ = v54
    }
    for fi55 := 0; fi55 < 13; fi55++ {
        fmt.Println(fi55)
        _ = fi55
    }
    v56 := GenS8{F0: 42, F1: uint16(v1)}
    for fi57 := 0; fi57 < 19; fi57++ {
        var v58 bool = true
        _ = fi57
        _ = v58
    }
    for fi59 := 0; fi59 < 1; fi59++ {
        var v60 float32 = (1e10 - (0.1 + (-2.25 + ((-2.25 * float32(((v2 * 0.1) * v2))) * 0.0))))
        fmt.Println(v2, v4, v5)
        if (false || ((v1 ^ v1) != (v1 * 0))) {
            for fi61 := 0; fi61 < 11; fi61++ {
                fmt.Println(fi59)
                v62 := GenS9{F0: (("a" + "") + ("\n\t" + ("hello" + "a"))), F1: 544}
                var v63 uint = 57
                _ = fi61
                _ = v62
                _ = v63
            }
            if (!false) {
                v1 -= v1
                v4 = (v4 / ((int8(v5)) | 1))
            } else {
                if (!true) {
                    v4 = v4
                } else {
                    switch "\n\t" {
                    case "":
                        for fi64 := 0; fi64 < 4; fi64++ {
                            v1 += v1
                            var v65 uint32 = 450
                            _ = fi64
                            _ = v65
                        }
                    case "hello":
                        if (!true) {
                            if (9223372036854775807 > -902) {
                                v1 += (v1 * (v1 % ((288) | 1)))
                                for fi66 := 0; fi66 < 15; fi66++ {
                                    switch "a" {
                                    case "":
                                        v67 := GenS10{F0: (((((v2 + v2) - float64(v60)) / float64(((v60 - v60) + (1e10 + -2.25)))) + v2) < (v2 * v2))}
                                        _ = v67
                                    }
                                    if false {
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        v68 := GenS9{F0: "a", F1: v1}
                                        _ = v68
                                    }
                                    v69 := GenS11{F0: 790, F1: v5, F2: ((((float32(3.14159) / v60) + v60) / v60) != 3.14159)}
                                    _ = fi66
                                    _ = v69
                                }
                                _ = len(v3)
                            }
                            switch -867 {
                            case 9223372036854775807:
                                v1 = (^(v1))
                            case 2147483647:
                                v70 := GenS12{F0: 255, F1: fi59, F2: 18446744073709551615}
                                _ = v70
                            }
                            v5 = (v5 * 42)
                        }
                    case "a":
                        _ = len(v3)
                    default:
                        v1 += 390
                    }
                    switch (fi59 & fi59) {
                    case -156:
                        var v71 int8 = -1
                        _ = v71
                    case 2147483647:
                        v3 = append(v3, "")
                    case 127:
                        var v72 int32 = 2147483647
                        _ = v72
                    default:
                        for fi73 := 0; fi73 < 4; fi73++ {
                            var v74 uint8 = 2
                            v60 -= 0.1
                            _ = fi73
                            _ = v74
                        }
                    }
                }
                var v75 int = (fi59 % ((fi59) | 1))
                _ = v75
            }
        }
        _ = fi59
        _ = v60
    }
    var v76 uint16 = 65535
    v4 -= (v4 % (((v4 << ((18446744073709551615) & 63))) | 1))
    if ((true || (((v5 - v5) != (v5 & -371)) && (true && true))) || ("a" > "hello")) {
        for fi77 := 0; fi77 < 12; fi77++ {
            for fi78 := 0; fi78 < 7; fi78++ {
                var v79 float64 = v2
                var v80 int = fi78
                _ = fi78
                _ = v79
                _ = v80
            }
            v76 += v76
            if false {
                var v81 uint = 1
                if true {
                    fmt.Println(v81)
                    v2 -= (1.5 / v2)
                    v5 = v5
                } else {
                    var v82 string = (("hello" + "\n\t") + "")
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    _ = v82
                }
                var v83 bool = true
                _ = v81
                _ = v83
            }
            _ = fi77
        }
        fmt.Println(v4)
    }
    v84 := GenS9{F0: "\n\t", F1: (^(v1))}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v76:", v76)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v47
    _ = v5
    _ = v56
    _ = v76
    _ = v84
}
