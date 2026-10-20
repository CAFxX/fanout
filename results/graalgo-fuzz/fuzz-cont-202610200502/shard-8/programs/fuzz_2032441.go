package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string; F2 bool }
type GenS2 struct { F0 uint64; F1 int16; F2 uint }
type GenS3 struct { F0 int; F1 int8; F2 bool }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int8; F1 uint16; F2 uint16 }
type GenS6 struct { F0 uint; F1 bool; F2 uint32 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 uint16; F1 uint16; F2 uint32 }
type GenS9 struct { F0 int32; F1 int16; F2 uint64 }
type GenS10 struct { F0 int8; F1 string; F2 uint64 }
type GenS11 struct { F0 int8 }
func genFn0(p0 int8, p1 int16) uint64 {
    for fi4 := 0; fi4 < 15; fi4++ {
        var v5 int = fi4
        v6 := []int{(^(v5)), (-(v5)), fi4, 32768}
        _ = fi4
        _ = v5
        _ = v6
    }
    v7 := GenS1{F0: 255, F1: "\n\t", F2: false}
    if (18446744073709551615 != 4294967295) {
        fmt.Println(p0, p0)
        var v8 float64 = float64(((0.0 - 0.0) - 0.1))
        if true {
            for fi9 := 0; fi9 < 20; fi9++ {
                fmt.Println(fi9, fi9)
                v10 := GenS1{F0: 56, F1: "foo bar", F2: (true && true)}
                _ = fi9
                _ = v10
            }
            for fi11 := 0; fi11 < 1; fi11++ {
                v12 := GenS1{F0: 71, F1: "", F2: false}
                if (0 != 1) {
                    fmt.Println(v8, p1, p0)
                } else {
                    if false {
                        v13 := GenS2{F0: 42, F1: p1, F2: 255}
                        var v14 uint8 = 88
                        var v15 uint64 = 42
                        _ = v13
                        _ = v14
                        _ = v15
                    }
                    v16 := GenS3{F0: fi11, F1: p0, F2: true}
                    _ = v16
                }
                p1 = 1000
                _ = fi11
                _ = v12
            }
        } else {
            var v17 uint32 = 42
            for fi18 := 0; fi18 < 8; fi18++ {
                if false {
                    var v19 int32 = -191
                    for fi20 := 0; fi20 < 12; fi20++ {
                        switch -2147483648 {
                        case 127:
                            fmt.Println(p1, v8, fi18)
                        case 4294967295:
                            v17 += v17
                        }
                        if (true && (-1 == (fi18 * (fi18 ^ fi20)))) {
                            switch (fi20 & -2147483648) {
                            case 32768:
                                p0 += p0
                            case 1:
                                fmt.Println(fi20, fi20, fi20)
                                var v21 uint8 = 204
                                _ = v21
                            case 4294967295:
                                fmt.Println(fi20, fi20, fi18)
                            default:
                                var v22 int64 = 42
                                _ = v22
                            }
                            v23 := []int{(fi20 % ((fi20) | 1)), (fi18 ^ fi20)}
                            _ = v23
                        } else {
                            v24 := GenS2{F0: 88, F1: p1, F2: 42}
                            _ = v24
                        }
                        _ = fi20
                    }
                    _ = v19
                }
                _ = fi18
            }
            v25 := []int{}
            _ = v17
            _ = v25
        }
        _ = v8
    } else {
        var v26 int32 = 319
        for fi27 := 0; fi27 < 7; fi27++ {
            p0 += (p0 * p0)
            _ = fi27
        }
        if true {
            v28 := GenS4{F0: 2}
            _ = v28
        }
        _ = v26
    }
    v29 := []int{1, -9223372036854775808}
    if (!false) {
        switch 1 {
        case 9223372036854775807:
            switch -862 {
            case 32768:
                if (false && ((false || false) && ("hello" <= ((("" + (("a" + "") + (("" + ("hello" + "foo bar")) + ("hello" + ("a" + ("\n\t" + "")))))) + "\n\t") + "a")))) {
                    for fi30 := 0; fi30 < 12; fi30++ {
                        if (0.0 == 0.1) {
                            var v31 float64 = (-2.25 * 0.1)
                            if (((v31 + v31) + 1.5) <= (((3.14159 / ((v31 - (v31 / v31)) / float64(3.14159))) - v31) + 1.5)) {
                                p0 = 127
                                switch "\n\t" {
                                case "":
                                    v32 := GenS5{F0: p0, F1: 1, F2: 42}
                                    _ = v32
                                case "hello":
                                    v33 := GenS6{F0: 0, F1: (!(true || false)), F2: 2}
                                    _ = v33
                                case "a":
                                    var v34 int16 = (p1 * p1)
                                    _ = v34
                                default:
                                    var v35 uint32 = 2
                                    _ = v35
                                }
                                if true {
                                    var v36 float32 = 1.5
                                    for fi37 := 0; fi37 < 5; fi37++ {
                                        p1 = 32767
                                        fmt.Println(fi37, fi37)
                                        v38 := GenS5{F0: (p0 + p0), F1: 32767, F2: 255}
                                        _ = fi37
                                        _ = v38
                                    }
                                    _ = v36
                                }
                            }
                            for fi39 := 0; fi39 < 19; fi39++ {
                                if (752 < 65535) {
                                    var v40 int16 = 366
                                    _ = v40
                                } else {
                                    v29 = append(v29, 127)
                                }
                                _ = fi39
                            }
                            _ = v31
                        }
                        _ = fi30
                    }
                    for fi41 := 0; fi41 < 1; fi41++ {
                        p1 -= (p1 & p1)
                        _ = fi41
                    }
                }
            case 128:
                v29 = append(v29, -9223372036854775808)
            case 32767:
                v42 := GenS5{F0: p0, F1: 2, F2: 317}
                _ = v42
            default:
                switch ("\n\t" + ("foo bar" + "hello")) {
                case "hello":
                    if len(v29) > 0 {
                        v29[0] = 127
                    }
                default:
                    switch "foo bar" {
                    case "a":
                        v43 := GenS4{F0: 567}
                        _ = v43
                    case "foo bar":
                        if (2147483647 >= 2) {
                            for fi44 := 0; fi44 < 3; fi44++ {
                                var v45 uint32 = 814
                                switch "a" {
                                case "foo bar":
                                    fmt.Println(p0)
                                case "hello":
                                    fmt.Println(v45, fi44)
                                case "\n\t":
                                    v45 -= (v45 + 112)
                                default:
                                    fmt.Println(p1, p1, v45)
                                }
                                _ = fi44
                                _ = v45
                            }
                            v29 = append(v29, -1)
                            if false {
                                var v46 uint16 = 424
                                fmt.Println(v46, v46)
                                _ = len(v29)
                                _ = v46
                            }
                        } else {
                            var v47 int16 = -246
                            _ = v47
                        }
                    }
                }
            }
        default:
            for fi48 := 0; fi48 < 5; fi48++ {
                for fi49 := 0; fi49 < 2; fi49++ {
                    fmt.Println(fi48, fi48)
                    switch ("a" + "foo bar") {
                    case "hello":
                        v50 := GenS5{F0: p0, F1: 1, F2: 32767}
                        _ = v50
                    case "\n\t":
                        p0 -= (p0 + (p0 & (p0 % (((p0 << ((4294967295) & 63))) | 1))))
                    case "a":
                        v51 := GenS6{F0: 255, F1: false, F2: 65535}
                        _ = v51
                    }
                    p1 = 939
                    _ = fi49
                }
                p0 += (p0 - (-(p0)))
                var v52 float64 = 1.5
                _ = fi48
                _ = v52
            }
        }
        if len(v29) > 0 {
            v29[0] = 128
        }
    }
    _ = p0
    _ = p1
    _ = v29
    _ = v7
    return 4294967295
}
func genFn1(p0 int) uint {
    v53 := []int{(^(p0)), (p0 >> ((2) & 63)), (p0 * p0)}
    if (-128 <= -128) {
        if false {
            v54 := GenS7{F0: (0 <= 65535)}
            _ = v54
        }
        p0 += 4294967295
    }
    _ = p0
    _ = v53
    return 335
}

func main() {
    var v1 uint64 = 4294967295
    v2 := []int{4294967295, -9223372036854775808}
    var v3 int64 = -234
    for fi55 := 0; fi55 < 3; fi55++ {
        if (true || ((!false) || ((-(fi55)) > (^(fi55))))) {
            if true {
                v56 := GenS8{F0: 42, F1: 997, F2: 2147483647}
                v57 := GenS4{F0: 4294967295}
                _ = v56
                _ = v57
            } else {
                var v58 uint8 = 42
                _ = v58
            }
            _ = len(v2)
        }
        switch fi55 {
        case -582:
            switch "\n\t" {
            case "\n\t":
                if true {
                    v1 = (v1 / ((v1) | 1))
                    for fi59 := 0; fi59 < 1; fi59++ {
                        _ = len(v2)
                        var v60 float64 = 1e10
                        _ = fi59
                        _ = v60
                    }
                } else {
                    v3 = (v3 % ((-9223372036854775808) | 1))
                    for fi61 := 0; fi61 < 10; fi61++ {
                        var v62 int = fi61
                        _ = fi61
                        _ = v62
                    }
                    if (genFn0(1, 0) >= (v1 | 18446744073709551615)) {
                        v3 = v3
                        var v63 int32 = -1
                        fmt.Println(v63, v63, v63)
                        _ = v63
                    } else {
                        switch "" {
                        case "foo bar":
                            _ = len(v2)
                        }
                        var v64 bool = (!true)
                        _ = v64
                    }
                }
            }
        default:
            v1 = 732
        }
        _ = fi55
    }
    if true {
        var v65 int = -2147483648
        _ = v65
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v66 uint16 = 0
    v66 += (v66 + (v66 << ((2) & 63)))
    v1 = (v1 / ((v1) | 1))
    v67 := GenS8{F0: v66, F1: v66, F2: uint32(v1)}
    v68 := GenS9{F0: 42, F1: 528, F2: v1}
    v3 = (v3 << ((798) & 63))
    for fi69 := 0; fi69 < 20; fi69++ {
        for fi70 := 0; fi70 < 9; fi70++ {
            var v71 int8 = 81
            fmt.Println(v1, fi69)
            _ = fi70
            _ = v71
        }
        _ = fi69
    }
    fmt.Println(v3, v66, v66)
    switch -938 {
    case -355:
        v3 = v3
    case 9223372036854775807:
        v1 += (v1 >> ((65535) & 63))
    case 128:
        if true {
            v72 := GenS6{F0: 897, F1: ((false || (false && false)) && (-238 != -20)), F2: 588}
            for fi73 := 0; fi73 < 11; fi73++ {
                v74 := GenS10{F0: 127, F1: (("\n\t" + (("foo bar" + "") + ("\n\t" + ("hello" + "\n\t")))) + "\n\t"), F2: 974}
                for fi75 := 0; fi75 < 11; fi75++ {
                    if (!("hello" > "hello")) {
                        v1 = (^(v1))
                    } else {
                        if true {
                            v76 := GenS11{F0: 121}
                            var v77 string = (("" + "\n\t") + "foo bar")
                            v78 := GenS4{F0: 65535}
                            _ = v76
                            _ = v77
                            _ = v78
                        }
                    }
                    v2 = append(v2, fi75)
                    _ = fi75
                }
                _ = fi73
                _ = v74
            }
            var v79 uint = uint(v1)
            _ = v72
            _ = v79
        } else {
            fmt.Println(v1)
        }
    }
    v80 := GenS5{F0: 100, F1: (v66 ^ uint16(v1)), F2: 42}
    switch 0 {
    case 900:
        for fi81 := 0; fi81 < 20; fi81++ {
            fmt.Println(v66, v66, v3)
            _ = fi81
        }
    case -236:
        for fi82 := 0; fi82 < 14; fi82++ {
            for fi83 := 0; fi83 < 6; fi83++ {
                v84 := GenS6{F0: 62, F1: true, F2: 2147483647}
                if ("\n\t" <= "hello") {
                    v3 = (v3 >> ((136) & 63))
                    if (true && true) {
                        _ = len(v2)
                    }
                    v3 -= (v3 / ((9223372036854775807) | 1))
                }
                _ = fi83
                _ = v84
            }
            _ = fi82
        }
    default:
        v85 := GenS1{F0: 0, F1: "a", F2: false}
        _ = v85
    }
    fmt.Println(v66, v66, v66)
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v66:", v66)
    _ = v1
    _ = v2
    _ = v3
    _ = v66
    _ = v67
    _ = v68
    _ = v80
}
