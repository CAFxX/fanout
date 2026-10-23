package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32 }
type GenS2 struct { F0 bool; F1 int8; F2 bool }
type GenS3 struct { F0 string; F1 uint64 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 bool; F1 uint }
type GenS6 struct { F0 string; F1 int; F2 bool }
type GenS7 struct { F0 uint; F1 uint32; F2 int64 }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 int; F1 bool }
func genFn0(p0 bool) int64 {
    var v7 float64 = 0.0
    for fi8 := 0; fi8 < 19; fi8++ {
        v7 -= v7
        _ = fi8
    }
    v9 := GenS1{F0: -3, F1: 1}
    var v10 int8 = 42
    v11 := []int{4294967295, 128}
    _ = p0
    _ = v10
    _ = v11
    _ = v7
    _ = v9
    return -787
}
func genFn1() int32 {
    if (0 < 1) {
        var v12 bool = (1.5 >= 3.14159)
        v13 := GenS1{F0: 107, F1: 42}
        var v14 float32 = 3.14159
        _ = v12
        _ = v13
        _ = v14
    }
    if false {
        v15 := []int{-2147483648, 710, -1}
        _ = v15
    } else {
        if false {
            var v16 bool = (763 > 1)
            switch 541 {
            case 2147483647:
                switch (("foo bar" + ("\n\t" + "\n\t")) + "foo bar") {
                case "hello":
                    v17 := GenS1{F0: 0, F1: 2147483648}
                    _ = v17
                case "a":
                    for fi18 := 0; fi18 < 13; fi18++ {
                        switch ("" + "") {
                        case "":
                            switch "foo bar" {
                            case "":
                                var v19 float64 = -2.25
                                _ = v19
                            case "hello":
                                v20 := GenS2{F0: v16, F1: 1, F2: (202 >= 42)}
                                _ = v20
                            case "\n\t":
                                var v21 int32 = 0
                                _ = v21
                            default:
                                if true {
                                    fmt.Println(v16)
                                    if false {
                                        switch (fi18 | -1) {
                                        case 4294967295:
                                            for fi22 := 0; fi22 < 11; fi22++ {
                                                var v23 int16 = -32768
                                                _ = fi22
                                                _ = v23
                                            }
                                        case -9223372036854775808:
                                            var v24 uint16 = 42
                                            _ = v24
                                        default:
                                            for fi25 := 0; fi25 < 11; fi25++ {
                                                var v26 float32 = 1e10
                                                _ = fi25
                                                _ = v26
                                            }
                                        }
                                        for fi27 := 0; fi27 < 3; fi27++ {
                                            var v28 bool = (1 != 65535)
                                            switch "hello" {
                                            case "hello":
                                                v29 := GenS2{F0: v28, F1: -128, F2: ("a" > "\n\t")}
                                                _ = v29
                                            }
                                            v30 := GenS3{F0: "a", F1: 282}
                                            _ = fi27
                                            _ = v28
                                            _ = v30
                                        }
                                    }
                                } else {
                                    switch (fi18 | fi18) {
                                    case 1:
                                        if true {
                                            for fi31 := 0; fi31 < 5; fi31++ {
                                                v32 := GenS4{F0: "a"}
                                                v16 = (((("a" + "hello") + ("\n\t" + "\n\t")) == ("foo bar" + "")) || (v16 && v16))
                                                v33 := GenS4{F0: "a"}
                                                _ = fi31
                                                _ = v32
                                                _ = v33
                                            }
                                        } else {
                                            v16 = (!v16)
                                            if v16 {
                                                if (genFn0(("" == ("a" + ("\n\t" + ("a" + ""))))) < 674) {
                                                    var v34 uint64 = 4294967295
                                                    _ = v34
                                                }
                                                var v35 uint8 = 128
                                                _ = v35
                                            } else {
                                                if v16 {
                                                    for fi36 := 0; fi36 < 16; fi36++ {
                                                        var v37 int32 = -1
                                                        v37 -= v37
                                                        _ = fi36
                                                        _ = v37
                                                    }
                                                    for fi38 := 0; fi38 < 10; fi38++ {
                                                        v39 := []int{(fi18 | (^(fi18)))}
                                                        if len(v39) >= 2 {
                                                            v39 = v39[1:len(v39)]
                                                        }
                                                        fmt.Println(fi18)
                                                        _ = fi38
                                                        _ = v39
                                                    }
                                                    v16 = (fi18 < (fi18 - -861))
                                                } else {
                                                    v16 = ((0.1 - 1.5) <= 1.5)
                                                    if (("" + "hello") != ((("a" + "hello") + "foo bar") + "\n\t")) {
                                                        v40 := []int{9223372036854775807, (fi18 & fi18), -291}
                                                        v41 := GenS2{F0: v16, F1: 100, F2: (float32((1.5 + (0.0 - 3.14159))) > ((((1.5 - float32((0.1 - (1e10 * -2.25)))) / 1.0) * ((1e10 / 1.5) - 0.1)) * (1e10 * 1e10)))}
                                                        var v42 uint16 = 851
                                                        _ = v40
                                                        _ = v41
                                                        _ = v42
                                                    }
                                                }
                                            }
                                            var v43 bool = v16
                                            _ = v43
                                        }
                                    case 39:
                                        v16 = false
                                    default:
                                        for fi44 := 0; fi44 < 12; fi44++ {
                                            v16 = (128 > 95)
                                            var v45 uint = 0
                                            v46 := GenS4{F0: "a"}
                                            _ = fi44
                                            _ = v45
                                            _ = v46
                                        }
                                    }
                                }
                            }
                        case "hello":
                            for fi47 := 0; fi47 < 19; fi47++ {
                                v16 = (32767 == (fi18 - fi18))
                                v48 := GenS5{F0: (fi18 < 9223372036854775807), F1: 334}
                                if v16 {
                                    v49 := GenS6{F0: "a", F1: 128, F2: (18446744073709551615 != 880)}
                                    v16 = (!((fi18 == fi18) || v16))
                                    v16 = v16
                                    _ = v49
                                }
                                _ = fi47
                                _ = v48
                            }
                        case "a":
                            if v16 {
                                if true {
                                    switch "foo bar" {
                                    case "foo bar":
                                        v16 = ("foo bar" < ("a" + "\n\t"))
                                    case "":
                                        switch ("hello" + "a") {
                                        case "hello":
                                            switch "\n\t" {
                                            case "foo bar":
                                                v50 := []int{32767, -851, -9223372036854775808, 0}
                                                _ = v50
                                            default:
                                                v51 := []int{(^(fi18)), (^(fi18))}
                                                _ = v51
                                            }
                                        case "":
                                            var v52 string = ""
                                            _ = v52
                                        case "foo bar":
                                            var v53 int16 = 302
                                            _ = v53
                                        default:
                                            var v54 uint32 = 92
                                            _ = v54
                                        }
                                    default:
                                        v55 := GenS1{F0: 1, F1: 429}
                                        _ = v55
                                    }
                                    for fi56 := 0; fi56 < 13; fi56++ {
                                        switch (-(fi56)) {
                                        case -447:
                                            var v57 uint64 = 1
                                            _ = v57
                                        default:
                                            if true {
                                                v58 := []int{fi18, (fi56 % (((^(fi18))) | 1))}
                                                v16 = ((-9223372036854775808 != 42) || (0 > 255))
                                                v16 = (v16 && v16)
                                                _ = v58
                                            }
                                        }
                                        var v59 uint8 = 42
                                        v60 := []int{128, fi18}
                                        _ = fi56
                                        _ = v59
                                        _ = v60
                                    }
                                } else {
                                    for fi61 := 0; fi61 < 8; fi61++ {
                                        fmt.Println(v16)
                                        _ = fi61
                                    }
                                    for fi62 := 0; fi62 < 18; fi62++ {
                                        v16 = true
                                        switch ("foo bar" + "\n\t") {
                                        case "":
                                            v16 = (v16 || false)
                                        default:
                                            v63 := GenS5{F0: (-39 != 0), F1: 4294967295}
                                            var v64 uint = 42
                                            _ = v63
                                            _ = v64
                                        }
                                        _ = fi62
                                    }
                                }
                                fmt.Println(fi18, fi18)
                                fmt.Println(v16, fi18)
                            } else {
                                var v65 float64 = 0.0
                                _ = v65
                            }
                        default:
                            var v66 uint16 = 1
                            _ = v66
                        }
                        _ = fi18
                    }
                }
            case 127:
                if v16 {
                    var v67 bool = (v16 && (!v16))
                    var v68 bool = v16
                    _ = v67
                    _ = v68
                } else {
                    for fi69 := 0; fi69 < 14; fi69++ {
                        switch "a" {
                        case "a":
                            var v70 int64 = 0
                            _ = v70
                        case "hello":
                            v71 := GenS7{F0: 2, F1: 1, F2: -484}
                            _ = v71
                        case "foo bar":
                            var v72 int32 = 42
                            _ = v72
                        default:
                            var v73 uint8 = 0
                            _ = v73
                        }
                        _ = fi69
                    }
                }
            case 9223372036854775807:
                v74 := []int{-9223372036854775808}
                _ = v74
            }
            _ = v16
        }
    }
    return 1
}

func main() {
    var v1 uint8 = 128
    var v2 uint32 = 339
    var v3 int8 = -128
    var v4 int64 = 1
    v5 := []int{0, 0, 9223372036854775807, 32767}
    var v6 uint16 = 255
    var v75 int16 = 1000
    var v76 float32 = (float32(float64(3.14159)) * (1e10 * ((1.5 * -2.25) + 3.14159)))
    switch -130 {
    case -1:
        v5 = append(v5, 32767)
    default:
        var v77 uint = 2
        _ = v77
    }
    var v78 int64 = (v4 % (((-(v4))) | 1))
    v5 = append(v5, -9223372036854775808)
    switch "\n\t" {
    case "hello":
        fmt.Println(v1)
    case "foo bar":
        for fi79 := 0; fi79 < 17; fi79++ {
            if (2 < uint(v2)) {
                v4 -= int64(fi79)
            } else {
                var v80 float32 = v76
                _ = v80
            }
            if false {
                v81 := GenS7{F0: 657, F1: (v2 >> ((4294967295) & 63)), F2: 9223372036854775807}
                fmt.Println(v3)
                v3 -= v3
                _ = v81
            } else {
                v78 += v4
                for fi82 := 0; fi82 < 11; fi82++ {
                    v3 -= (v3 + 1)
                    v6 = v6
                    _ = fi82
                }
            }
            _ = fi79
        }
    }
    v3 += -1
    var v83 string = ((("hello" + ("a" + "a")) + "foo bar") + ("hello" + "\n\t"))
    v6 -= 894
    var v84 int64 = 850
    v3 += v3
    if ((true || true) && false) {
        v3 = v3
        for fi85 := 0; fi85 < 1; fi85++ {
            _ = len(v5)
            v1 = v1
            switch v83 {
            case "hello":
                switch fi85 {
                case -1:
                    fmt.Println(v4)
                default:
                    var v86 uint32 = v2
                    _ = v86
                }
            case "\n\t":
                var v87 bool = ("foo bar" > (v83 + v83))
                _ = v87
            }
            _ = fi85
        }
        var v88 int8 = (v3 << ((0) & 63))
        _ = v88
    } else {
        switch 2147483647 {
        case 149:
            if false {
                v89 := GenS5{F0: false, F1: 65535}
                switch (v83 + ((((v83 + v83) + v83) + "a") + (v83 + v83))) {
                case "\n\t":
                    v90 := GenS8{F0: v1}
                    _ = v90
                case "a":
                    fmt.Println(v4)
                default:
                    v78 = v84
                }
                fmt.Println(v4, v6, v6)
                _ = v89
            }
        default:
            fmt.Println(v84, v84)
        }
    }
    switch 127 {
    case 128:
        v5 = append(v5, 405)
    default:
        if (0 < 42) {
            var v91 int16 = v75
            _ = v91
        }
    }
    var v92 int = 2147483647
    v76 += v76
    v93 := GenS9{F0: (-(v92)), F1: true}
    var v94 float32 = (v76 + 0.0)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v75:", v75)
    fmt.Println("v76:", v76)
    fmt.Println("v78:", v78)
    fmt.Println("v83:", v83)
    fmt.Println("v84:", v84)
    fmt.Println("v92:", v92)
    fmt.Println("v94:", v94)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v75
    _ = v76
    _ = v78
    _ = v83
    _ = v84
    _ = v92
    _ = v93
    _ = v94
}
