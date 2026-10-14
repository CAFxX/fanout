package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 bool; F1 int; F2 uint8 }
type GenS4 struct { F0 int16; F1 uint32; F2 uint16 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 int16; F1 int16 }
type GenS8 struct { F0 int8; F1 uint8 }
func genFn0(p0 int) int16 {
    v5 := GenS1{F0: p0, F1: ("foo bar" + "hello")}
    var v6 uint32 = 65535
    if (18446744073709551615 != 140) {
        v7 := GenS1{F0: p0, F1: "\n\t"}
        switch "\n\t" {
        case "\n\t":
            for fi8 := 0; fi8 < 12; fi8++ {
                fmt.Println(p0, p0)
                _ = fi8
            }
        default:
            fmt.Println(p0, v6)
        }
        _ = v7
    } else {
        p0 = (p0 >> ((218) & 63))
    }
    fmt.Println(v6)
    v9 := []int{}
    _ = p0
    _ = v5
    _ = v6
    _ = v9
    return -219
}

func main() {
    var v1 int16 = -424
    v2 := []string{("foo bar" + "")}
    var v3 bool = ((((1e10 - 1e10) * float64(1.5)) + float64(3.14159)) > 0.1)
    v4 := make(map[int]uint32)
    switch "hello" {
    case "":
        v3 = v3
    case "a":
        switch 0 {
        case 0:
            v10 := GenS2{F0: 2}
            _ = v10
        default:
            _ = len(v2)
        }
    case "hello":
        switch "foo bar" {
        case "foo bar":
            var v11 bool = (0 != 2)
            _ = v11
        case "hello":
            v12 := GenS2{F0: 42}
            _ = v12
        case "":
            v13 := GenS1{F0: -2147483648, F1: "hello"}
            _ = v13
        default:
            v3 = (3.14159 != (3.14159 * -2.25))
        }
    default:
        v3 = (v3 || ((1e10 * 1.5) <= 1.5))
    }
    var v14 uint = 2
    v3 = (0 <= 255)
    if (2147483647 >= 0) {
        var v15 int32 = -2147483648
        for fi16 := 0; fi16 < 14; fi16++ {
            for fi17 := 0; fi17 < 1; fi17++ {
                switch "a" {
                case "":
                    switch fi17 {
                    case 2147483647:
                        var v18 string = "a"
                        _ = v18
                    case 187:
                        var v19 string = (("\n\t" + ("a" + "foo bar")) + (((("hello" + "foo bar") + "\n\t") + ("" + ((("hello" + ("\n\t" + "hello")) + ("\n\t" + "\n\t")) + "\n\t"))) + ((("a" + "foo bar") + (("foo bar" + "") + "")) + "hello")))
                        _ = v19
                    case 4294967295:
                        if (676 > v14) {
                            var v20 uint64 = 724
                            _ = v20
                        } else {
                            var v21 int = 32767
                            v22 := GenS2{F0: 18446744073709551615}
                            fmt.Println(v15, v21, v3)
                            _ = v21
                            _ = v22
                        }
                    default:
                        var v23 uint32 = 65535
                        switch 42 {
                        case -763:
                            var v24 int8 = -1
                            _ = v24
                        case 128:
                            v25 := GenS3{F0: ((!(1.5 <= -2.25)) || (!((fi17 / ((fi16) | 1)) != fi17))), F1: fi16, F2: 52}
                            _ = v25
                        case -26:
                            var v26 uint32 = (v23 / ((1) | 1))
                            _ = v26
                        }
                        var v27 int8 = -128
                        _ = v23
                        _ = v27
                    }
                case "\n\t":
                    var v28 uint64 = 0
                    _ = v28
                case "hello":
                    var v29 float32 = (3.14159 * 1.5)
                    _ = v29
                default:
                    _ = len(v2)
                }
                var v30 int16 = int16(fi16)
                _ = fi17
                _ = v30
            }
            var v31 int32 = (v15 * v15)
            if true {
                var v32 uint32 = 1
                _ = v32
            } else {
                for fi33 := 0; fi33 < 3; fi33++ {
                    if v3 {
                        _ = len(v4)
                        if (v14 >= (^(v14))) {
                            if (genFn0((-(fi33))) <= v1) {
                                if (v14 > v14) {
                                    switch 9223372036854775807 {
                                    case 32767:
                                        v34 := v4[fi33]
                                        _ = v34
                                    }
                                    var v35 int16 = -32768
                                    if (!v3) {
                                        var v36 int64 = 1
                                        v35 = v1
                                        var v37 uint32 = 2
                                        _ = v36
                                        _ = v37
                                    }
                                    _ = v35
                                }
                                var v38 int64 = 9223372036854775807
                                _ = v38
                            } else {
                                v2 = append(v2, ("hello" + ""))
                                v4[(-(fi33))] = 0
                                if v3 {
                                    fmt.Println(v14)
                                    v39 := GenS4{F0: v1, F1: 230, F2: 2}
                                    _ = v39
                                } else {
                                    fmt.Println(v14, fi33)
                                    fmt.Println(fi16, fi16, v3)
                                }
                            }
                        }
                        if (1 > 42) {
                            var v40 int32 = (v15 / (((v15 | v31)) | 1))
                            _ = len(v4)
                            _ = v40
                        }
                    } else {
                        v2 = append(v2, "\n\t")
                        v3 = (v3 && v3)
                    }
                    _ = len(v4)
                    _ = fi33
                }
                var v41 bool = v3
                _ = v41
            }
            _ = fi16
            _ = v31
        }
        _ = v15
    }
    if (!v3) {
        v3 = (42 >= 0)
        for fi42 := 0; fi42 < 2; fi42++ {
            if (v3 && (!false)) {
                var v43 bool = (9223372036854775807 < 9223372036854775807)
                _ = v43
            } else {
                v44 := GenS2{F0: 4294967295}
                _ = len(v2)
                var v45 uint32 = 2147483648
                _ = v44
                _ = v45
            }
            _ = fi42
        }
        switch 32767 {
        case 32767:
            switch 32767 {
            case 42:
                switch ("hello" + ("foo bar" + (("" + (("a" + "") + (("a" + ((("a" + "\n\t") + "") + ("a" + ("hello" + (("" + "a") + "a"))))) + "foo bar"))) + ""))) {
                case "a":
                    if (v3 && v3) {
                        v46 := v4[-849]
                        _ = v46
                    } else {
                        v1 = (v1 % ((genFn0(32767)) | 1))
                        v14 += 0
                        v47 := GenS5{F0: 311}
                        _ = v47
                    }
                case "foo bar":
                    v3 = false
                case "":
                    for fi48 := 0; fi48 < 10; fi48++ {
                        switch fi48 {
                        case 373:
                            for fi49 := 0; fi49 < 13; fi49++ {
                                if len(v2) > 0 {
                                    v2[0] = ""
                                }
                                switch (fi49 | fi49) {
                                case 911:
                                    v3 = (false || (32768 >= 865))
                                case 9223372036854775807:
                                    var v50 float64 = float64(((1e10 - float32((float64(((3.14159 - 3.14159) / -2.25)) + 1e10))) - 3.14159))
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    _ = v50
                                default:
                                    switch fi49 {
                                    case -445:
                                        if v3 {
                                            v1 += (v1 / (((^(v1))) | 1))
                                            v51 := GenS6{F0: (693 == (v14 << (((v14 >> ((18446744073709551615) & 63))) & 63)))}
                                            _ = len(v4)
                                            if true {
                                                var v52 string = ("\n\t" + "")
                                                _ = v52
                                            } else {
                                                var v53 bool = true
                                                _ = v53
                                            }
                                            fmt.Println(fi49)
                                            _ = v51
                                        } else {
                                            if (uint32(v14) == 42) {
                                                if ("\n\t" == ("foo bar" + "foo bar")) {
                                                    var v54 int64 = 1
                                                    _ = len(v2)
                                                    var v55 uint32 = 83
                                                    _ = v54
                                                    _ = v55
                                                }
                                                switch "foo bar" {
                                                case "\n\t":
                                                    v14 += v14
                                                case "hello":
                                                    fmt.Println(fi48, v3, v3)
                                                default:
                                                    switch 4294967295 {
                                                    case 127:
                                                        v14 = (v14 + (^(v14)))
                                                    case 0:
                                                        if v3 {
                                                            if (!(((303 < 0) || false) && ((fi48 % ((-57) | 1)) != -192))) {
                                                                if v3 {
                                                                    if (399 < 1) {
                                                                        v1 -= v1
                                                                        switch (fi49 & (fi48 << (((v14 ^ v14)) & 63))) {
                                                                        case -9223372036854775808:
                                                                            switch -936 {
                                                                            case 0:
                                                                                var v56 uint64 = 2
                                                                                _ = v56
                                                                            case 4294967295:
                                                                                var v57 uint64 = 708
                                                                                _ = v57
                                                                            }
                                                                        case -782:
                                                                            switch (fi48 + 4294967295) {
                                                                            case -522:
                                                                                var v58 int = fi48
                                                                                _ = v58
                                                                            default:
                                                                                switch (("" + "a") + ("a" + (("a" + "") + "a"))) {
                                                                                case "":
                                                                                    v1 = (v1 * (v1 + (v1 % ((genFn0(-595)) | 1))))
                                                                                case "a":
                                                                                    switch (fi48 / (((fi48 | -409)) | 1)) {
                                                                                    case 4294967295:
                                                                                        fmt.Println(fi48, v1)
                                                                                    default:
                                                                                        var v59 bool = v3
                                                                                        _ = v59
                                                                                    }
                                                                                case "foo bar":
                                                                                    v1 = (v1 + v1)
                                                                                }
                                                                            }
                                                                        case 32767:
                                                                            v60 := GenS2{F0: 0}
                                                                            _ = v60
                                                                        default:
                                                                            v61 := GenS3{F0: (602 == 0), F1: fi49, F2: 98}
                                                                            _ = v61
                                                                        }
                                                                    }
                                                                    var v62 int64 = -173
                                                                    _ = v62
                                                                }
                                                                if false {
                                                                    v63 := GenS2{F0: 992}
                                                                    v64 := GenS7{F0: v1, F1: 482}
                                                                    _ = v63
                                                                    _ = v64
                                                                }
                                                                v3 = (!v3)
                                                                var v65 int8 = 1
                                                                _ = v65
                                                            }
                                                            switch "a" {
                                                            case "a":
                                                                v66 := GenS2{F0: uint64(v14)}
                                                                _ = v66
                                                            case "":
                                                                fmt.Println(v3, fi48, v14)
                                                            case "\n\t":
                                                                v67 := GenS1{F0: (fi48 * fi48), F1: ("hello" + "\n\t")}
                                                                _ = v67
                                                            }
                                                        }
                                                    default:
                                                        v68 := GenS2{F0: 508}
                                                        _ = v68
                                                    }
                                                }
                                                switch "hello" {
                                                case "\n\t":
                                                    if true {
                                                        var v69 bool = true
                                                        _ = v69
                                                    } else {
                                                        if (("a" <= "foo bar") || (v3 && (18446744073709551615 <= 285))) {
                                                            if (2147483647 != 65535) {
                                                                v70 := GenS3{F0: false, F1: fi49, F2: 2}
                                                                v4[fi49] = 337
                                                                _ = len(v4)
                                                                if (0 > -9223372036854775808) {
                                                                    v71 := GenS1{F0: -1, F1: ("\n\t" + ("hello" + ((("" + "foo bar") + ("a" + "foo bar")) + "\n\t")))}
                                                                    v72 := GenS8{F0: int8(v14), F1: 128}
                                                                    _ = v71
                                                                    _ = v72
                                                                } else {
                                                                    var v73 uint8 = 0
                                                                    fmt.Println(v3, v1)
                                                                    _ = v73
                                                                }
                                                                _ = v70
                                                            }
                                                            v4[32768] = 760
                                                        }
                                                    }
                                                }
                                            }
                                            if true {
                                                fmt.Println(fi48, v14, v3)
                                            }
                                            var v74 int16 = -1
                                            _ = v74
                                        }
                                    }
                                }
                                _ = fi49
                            }
                        }
                        if false {
                            v75 := GenS4{F0: v1, F1: 65535, F2: 65535}
                            var v76 string = "a"
                            v2 = append(v2, (v76 + v76))
                            _ = v75
                            _ = v76
                        }
                        v1 = v1
                        _ = fi48
                    }
                default:
                    fmt.Println(v3)
                }
            case 32767:
                var v77 bool = (42 > 2147483648)
                _ = v77
            case 655:
                if true {
                    var v78 float64 = (1e10 - ((0.0 * 1.5) * -2.25))
                    for fi79 := 0; fi79 < 12; fi79++ {
                        var v80 float32 = ((-2.25 / (float32(1e10) * (1.5 * (1e10 + -2.25)))) - (1.5 * 0.1))
                        _ = fi79
                        _ = v80
                    }
                    _ = v78
                } else {
                    var v81 float32 = 3.14159
                    var v82 uint = (v14 >> (((v14 ^ (v14 - v14))) & 63))
                    _ = v81
                    _ = v82
                }
            }
        case -2147483648:
            v3 = ("hello" < ("foo bar" + ""))
        case -9223372036854775808:
            switch 128 {
            case 9223372036854775807:
                v1 -= -80
            }
        }
    } else {
        if v3 {
            fmt.Println(v1, v14)
        } else {
            v83 := GenS6{F0: (false && (0 < 32767))}
            _ = len(v2)
            _ = v83
        }
        v1 -= (v1 ^ v1)
        for fi84 := 0; fi84 < 12; fi84++ {
            v85 := GenS2{F0: 64}
            _ = len(v4)
            _ = fi84
            _ = v85
        }
    }
    v86 := GenS1{F0: 26, F1: ("hello" + "\n\t")}
    v1 -= (v1 & v1)
    if ((9223372036854775807 > 128) || false) {
        var v87 int64 = 0
        _ = v87
    }
    var v88 int32 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v88:", v88)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v86
    _ = v88
}
