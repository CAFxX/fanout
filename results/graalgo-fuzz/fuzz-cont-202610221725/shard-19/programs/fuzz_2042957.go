package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32; F2 int16 }
type GenS2 struct { F0 uint8; F1 uint64 }
type GenS3 struct { F0 string; F1 uint32 }
type GenS4 struct { F0 string; F1 string; F2 uint64 }
type GenS5 struct { F0 bool; F1 string }
type GenS6 struct { F0 string; F1 uint8 }
type GenS7 struct { F0 int32; F1 uint32 }
type GenS8 struct { F0 int8 }
func genFn0(p0 uint, p1 bool, p2 bool) int8 {
    v5 := []int{}
    fmt.Println(p0)
    _ = len(v5)
    switch -2147483648 {
    case -9223372036854775808:
        v6 := GenS1{F0: 1, F1: 0, F2: 42}
        _ = v6
    case -583:
        var v7 float64 = -2.25
        _ = v7
    case 127:
        p1 = ((p1 || p1) && (255 != 891))
    default:
        v8 := GenS1{F0: 1, F1: 42, F2: 1}
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 100
}
func genFn1() int {
    var v9 int = -1
    var v10 bool = true
    v11 := []int{}
    for fi12 := 0; fi12 < 10; fi12++ {
        v9 = fi12
        _ = fi12
    }
    _ = v10
    _ = v11
    _ = v9
    return (-(v9))
}

func main() {
    var v1 uint8 = 42
    v2 := []uint32{}
    var v3 uint = 255
    var v4 int32 = 2147483647
    var v13 int8 = -128
    v3 += 4294967295
    v3 -= 2
    switch "\n\t" {
    case "hello":
        if false {
            v13 = (-(v13))
        } else {
            v13 += v13
            v1 -= v1
            v4 = 997
        }
    case "\n\t":
        if len(v2) > 0 {
            v2[0] = 2147483647
        }
    }
    v4 -= 1
    if ((127 != 127) && true) {
        switch ("a" + "a") {
        case "hello":
            var v14 float32 = 0.1
            _ = v14
        case "foo bar":
            for fi15 := 0; fi15 < 5; fi15++ {
                v16 := GenS1{F0: v1, F1: v4, F2: -918}
                _ = fi15
                _ = v16
            }
        default:
            if false {
                if false {
                    for fi17 := 0; fi17 < 8; fi17++ {
                        switch -297 {
                        case -9223372036854775808:
                            for fi18 := 0; fi18 < 16; fi18++ {
                                v13 = (v13 & v13)
                                _ = fi18
                            }
                        case 232:
                            v2 = append(v2, 918)
                        default:
                            var v19 bool = true
                            _ = v19
                        }
                        v2 = append(v2, 2147483647)
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        _ = fi17
                    }
                    v20 := GenS1{F0: v1, F1: v4, F2: 719}
                    fmt.Println(v1)
                    _ = v20
                } else {
                    for fi21 := 0; fi21 < 2; fi21++ {
                        fmt.Println(v1, v1, v1)
                        if true {
                            fmt.Println(fi21, v3)
                        } else {
                            var v22 bool = false
                            v23 := GenS1{F0: 127, F1: v4, F2: -1}
                            if v22 {
                                v3 += (^(v3))
                            } else {
                                for fi24 := 0; fi24 < 2; fi24++ {
                                    if ((v3 | (^(v3))) == (^(v3))) {
                                        v13 += int8(v3)
                                        v1 = (v1 << ((65535) & 63))
                                        fmt.Println(fi24)
                                    } else {
                                        if len(v2) > 0 {
                                            v2[0] = 65535
                                        }
                                    }
                                    _ = fi24
                                }
                                v4 -= v4
                            }
                            _ = v22
                            _ = v23
                        }
                        _ = fi21
                    }
                }
                var v25 int8 = 37
                v3 += (^(v3))
                _ = v25
            } else {
                for fi26 := 0; fi26 < 20; fi26++ {
                    v4 += v4
                    _ = fi26
                }
            }
        }
        switch -9223372036854775808 {
        case 1:
            v3 = v3
        case 142:
            if (false && (!(!true))) {
                v4 = (v4 & -311)
                switch 1 {
                case -1:
                    v3 -= (v3 - (v3 | 18))
                }
            } else {
                _ = len(v2)
                switch ("" + ("" + "\n\t")) {
                case "":
                    switch 9223372036854775807 {
                    case 32767:
                        for fi27 := 0; fi27 < 4; fi27++ {
                            switch 2147483647 {
                            case 127:
                                v3 -= (v3 - (^(v3)))
                            case 42:
                                var v28 int8 = (v13 & (v13 - v13))
                                _ = v28
                            case -1:
                                var v29 uint32 = 2147483648
                                _ = v29
                            }
                            if ((("foo bar" <= ("" + (("a" + "") + "a"))) || (genFn0(v3, true, true) != 127)) || (571 != 4294967295)) {
                                _ = len(v2)
                                var v30 int8 = 86
                                _ = v30
                            } else {
                                var v31 int = fi27
                                if false {
                                    var v32 string = "hello"
                                    if len(v2) > 0 {
                                        v2[0] = 1
                                    }
                                    v3 = (v3 % ((v3) | 1))
                                    _ = v32
                                }
                                _ = v31
                            }
                            _ = fi27
                        }
                    }
                case "foo bar":
                    var v33 int = -2147483648
                    _ = v33
                default:
                    v34 := GenS1{F0: (v1 << (((v3 | v3)) & 63)), F1: v4, F2: 0}
                    _ = v34
                }
                for fi35 := 0; fi35 < 8; fi35++ {
                    v13 = v13
                    for fi36 := 0; fi36 < 18; fi36++ {
                        var v37 int32 = (v4 >> ((v3) & 63))
                        _ = fi36
                        _ = v37
                    }
                    _ = fi35
                }
            }
        case 2147483647:
            var v38 int = -1
            _ = v38
        }
        if (!true) {
            v39 := GenS2{F0: (v1 % ((v1) | 1)), F1: 711}
            v3 = 441
            _ = v39
        }
    } else {
        v4 += (v4 ^ -865)
        v40 := GenS1{F0: v1, F1: (v4 ^ v4), F2: int16(v3)}
        _ = v40
    }
    var v41 uint32 = 2
    for fi42 := 0; fi42 < 15; fi42++ {
        v41 += 4294967295
        if (!(18446744073709551615 != 42)) {
            var v43 uint16 = 0
            var v44 int32 = v4
            v45 := GenS3{F0: ("\n\t" + ("hello" + "")), F1: uint32(v1)}
            _ = v43
            _ = v44
            _ = v45
        }
        if len(v2) > 0 {
            v2[0] = 756
        }
        _ = fi42
    }
    v13 += v13
    switch "" {
    case "hello":
        var v46 uint32 = (v41 << ((v3) & 63))
        _ = v46
    case "\n\t":
        v13 = (v13 << (((v3 | 501)) & 63))
    case "foo bar":
        if false {
            fmt.Println(v3, v1)
            var v47 int = genFn1()
            fmt.Println(v1, v4)
            _ = v47
        } else {
            if (!(v3 < (v3 << ((v3) & 63)))) {
                fmt.Println(v13)
            } else {
                if true {
                    for fi48 := 0; fi48 < 16; fi48++ {
                        fmt.Println(v41, v4)
                        v2 = append(v2, v41)
                        for fi49 := 0; fi49 < 11; fi49++ {
                            if false {
                                if len(v2) > 0 {
                                    v2[0] = v41
                                }
                                v4 = 796
                                v50 := GenS1{F0: v1, F1: (^(v4)), F2: 0}
                                _ = v50
                            } else {
                                switch "\n\t" {
                                case "":
                                    switch (fi49 + fi48) {
                                    case 42:
                                        var v51 int8 = 42
                                        _ = v51
                                    case -1:
                                        var v52 int64 = -759
                                        _ = len(v2)
                                        v13 = (v13 ^ (v13 - (v13 / ((v13) | 1))))
                                        fmt.Println(v4, v52)
                                        _ = v52
                                    default:
                                        v2 = append(v2, (v41 - v41))
                                    }
                                case "foo bar":
                                    v13 -= (v13 & (v13 - -72))
                                }
                                if len(v2) > 0 {
                                    v2[0] = (v41 / (((v41 - 42)) | 1))
                                }
                            }
                            _ = fi49
                        }
                        _ = fi48
                    }
                    v41 += v41
                }
                if (127 > 1) {
                    v4 -= (v4 % ((int32(v41)) | 1))
                }
            }
            if (false && (0.0 > -2.25)) {
                switch genFn1() {
                case 762:
                    if true {
                        if (0.1 != 1.5) {
                            var v53 float64 = 1e10
                            v54 := GenS4{F0: ((("a" + "\n\t") + "\n\t") + ("" + "hello")), F1: (("hello" + ("" + "hello")) + "a"), F2: 82}
                            _ = v53
                            _ = v54
                        }
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        v55 := GenS4{F0: "hello", F1: ("a" + ""), F2: 2}
                        _ = v55
                    }
                case -9223372036854775808:
                    v13 += 111
                }
            } else {
                if true {
                    fmt.Println(v1, v4, v3)
                } else {
                    switch "" {
                    case "foo bar":
                        for fi56 := 0; fi56 < 18; fi56++ {
                            v2 = append(v2, v41)
                            switch ("hello" + "a") {
                            case "foo bar":
                                fmt.Println(v1)
                            case "\n\t":
                                v41 -= 4294967295
                            default:
                                if len(v2) > 0 {
                                    v2[0] = (v41 - 946)
                                }
                            }
                            switch ("hello" + (("" + "\n\t") + "\n\t")) {
                            case "\n\t":
                                v13 += v13
                            case "":
                                switch "foo bar" {
                                case "":
                                    switch (fi56 - (fi56 - -652)) {
                                    case 9223372036854775807:
                                        if (533 >= 65535) {
                                            if (false || (!(!false))) {
                                                v57 := GenS5{F0: true, F1: "hello"}
                                                switch (fi56 | (fi56 / (((fi56 + fi56)) | 1))) {
                                                case 32768:
                                                    var v58 int32 = (^(v4))
                                                    _ = v58
                                                }
                                                var v59 uint = 1
                                                _ = v57
                                                _ = v59
                                            }
                                            fmt.Println(v13)
                                            var v60 int = fi56
                                            _ = v60
                                        }
                                    case 32768:
                                        fmt.Println(fi56, v4, v13)
                                    default:
                                        v4 = (v4 / ((v4) | 1))
                                    }
                                case "\n\t":
                                    v61 := GenS6{F0: "hello", F1: (v1 - (v1 ^ (v1 & v1)))}
                                    _ = v61
                                }
                            case "foo bar":
                                v4 -= 2147483647
                            }
                            _ = fi56
                        }
                    }
                    v4 -= -1
                }
                for fi62 := 0; fi62 < 11; fi62++ {
                    v13 -= (v13 & 42)
                    for fi63 := 0; fi63 < 11; fi63++ {
                        switch "a" {
                        case "":
                            var v64 string = ("a" + "hello")
                            v1 += 149
                            v13 -= v13
                            _ = v64
                        default:
                            v4 += (v4 - v4)
                            v65 := GenS5{F0: (((!(!false)) && false) || (-9223372036854775808 < -1)), F1: ""}
                            v3 = v3
                            _ = v65
                        }
                        if true {
                            if (-728 != 1) {
                                if (!(42 > 915)) {
                                    var v66 int32 = (v4 ^ (-(v4)))
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    var v67 int16 = -1
                                    _ = v66
                                    _ = v67
                                }
                                fmt.Println(v1, v13)
                                fmt.Println(v4, fi62)
                                v68 := GenS2{F0: (v1 + (^(v1))), F1: 42}
                                _ = v68
                            } else {
                                v3 = v3
                                if true {
                                    v3 = (v3 << ((v3) & 63))
                                }
                            }
                            switch (^(fi63)) {
                            case 32768:
                                v1 = (v1 | (^(v1)))
                            case 0:
                                v69 := GenS2{F0: v1, F1: 1}
                                _ = v69
                            }
                        }
                        var v70 int8 = (^(v13))
                        _ = fi63
                        _ = v70
                    }
                    v41 = 807
                    _ = fi62
                }
            }
            v71 := GenS5{F0: (-794 >= -9223372036854775808), F1: "hello"}
            _ = v71
        }
    }
    fmt.Println(v3, v13)
    if (!(((v4 % (((v4 % ((-411) | 1))) | 1)) >= v4) && ((v41 | (^(v41))) >= 298))) {
        if (v13 < (v13 + v13)) {
            for fi72 := 0; fi72 < 4; fi72++ {
                fmt.Println(v13, v41, v1)
                v73 := GenS7{F0: v4, F1: (v41 >> ((484) & 63))}
                _ = fi72
                _ = v73
            }
            v2 = append(v2, 2147483648)
        }
        v1 = 2
    } else {
        fmt.Println(v41, v3, v41)
    }
    v74 := GenS8{F0: (v13 << (((v3 + uint(v1))) & 63))}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v74
}
