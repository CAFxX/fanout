package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int64; F1 uint }
type GenS5 struct { F0 uint; F1 string; F2 bool }
type GenS6 struct { F0 int16; F1 int8; F2 uint32 }
type GenS7 struct { F0 uint; F1 uint64 }
type GenS8 struct { F0 bool; F1 int; F2 uint16 }
type GenS9 struct { F0 int16 }
type GenS10 struct { F0 int64; F1 int; F2 string }
type GenS11 struct { F0 int64 }
type GenS12 struct { F0 int16; F1 int16; F2 uint64 }
func genFn0() uint64 {
    v5 := []int{-2147483648, 128, 9223372036854775807}
    var v6 int = 2147483647
    if (true && (false || false)) {
        v7 := GenS1{F0: 0}
        _ = v7
    } else {
        v8 := GenS2{F0: (v6 ^ -2147483648)}
        v9 := GenS1{F0: 939}
        _ = v8
        _ = v9
    }
    v6 = (v6 % (((^(v6))) | 1))
    _ = v5
    _ = v6
    return 42
}
func genFn1(p0 int) int8 {
    v10 := []int{}
    v11 := GenS3{F0: 18446744073709551615}
    _ = p0
    _ = v10
    _ = v11
    return -77
}

func main() {
    v1 := make(map[string]string)
    v2 := []uint32{0}
    v3 := make(map[string]int)
    var v4 int = 2147483647
    v4 = 127
    if true {
        if (-32768 == -32768) {
            if true {
                if (!(v4 >= 42)) {
                    v12 := GenS1{F0: 0}
                    _ = v12
                }
                switch "foo bar" {
                case "foo bar":
                    v4 = v4
                case "a":
                    var v13 uint8 = 177
                    _ = v13
                case "":
                    v4 += 127
                default:
                    for fi14 := 0; fi14 < 10; fi14++ {
                        v4 -= fi14
                        for fi15 := 0; fi15 < 7; fi15++ {
                            fmt.Println(fi15)
                            fmt.Println(fi15)
                            var v16 uint32 = 1
                            fmt.Println(fi14, v4, fi15)
                            _ = fi15
                            _ = v16
                        }
                        _ = fi14
                    }
                }
                v17 := GenS4{F0: -1, F1: 255}
                _ = v17
            } else {
                for fi18 := 0; fi18 < 5; fi18++ {
                    var v19 int64 = 9223372036854775807
                    var v20 bool = (true && (!false))
                    var v21 uint64 = 4294967295
                    _ = fi18
                    _ = v19
                    _ = v20
                    _ = v21
                }
                var v22 int = (v4 - v4)
                _ = v22
            }
            var v23 float64 = (0.0 * (3.14159 / -2.25))
            var v24 int16 = 1
            _ = v23
            _ = v24
        } else {
            if false {
                if true {
                    fmt.Println(v4)
                } else {
                    if false {
                        if true {
                            v4 = v4
                            v4 = v4
                            _ = len(v1)
                        } else {
                            if true {
                                if (1 < 42) {
                                    if (1 < 0) {
                                        v25 := GenS3{F0: 2}
                                        var v26 int8 = 1
                                        switch (v4 << ((18446744073709551615) & 63)) {
                                        case 0:
                                            v1["a"] = "foo bar"
                                        case 4294967295:
                                            fmt.Println(v26, v4)
                                        default:
                                            var v27 int16 = 754
                                            _ = v27
                                        }
                                        _ = v25
                                        _ = v26
                                    }
                                    var v28 uint16 = 185
                                    v29 := v1["a"]
                                    _ = v28
                                    _ = v29
                                }
                                delete(v1, "")
                                for fi30 := 0; fi30 < 8; fi30++ {
                                    v31 := GenS3{F0: 42}
                                    _ = fi30
                                    _ = v31
                                }
                            }
                        }
                        v4 -= (v4 - -251)
                        var v32 string = ("\n\t" + ((("\n\t" + ("hello" + "foo bar")) + ("" + "foo bar")) + ""))
                        _ = v32
                    } else {
                        if len(v2) > 0 {
                            v2[0] = 0
                        }
                    }
                    for fi33 := 0; fi33 < 17; fi33++ {
                        v4 += fi33
                        fmt.Println(fi33)
                        if (!(false && ((false || true) && true))) {
                            for fi34 := 0; fi34 < 5; fi34++ {
                                var v35 int = v4
                                v36 := GenS5{F0: 742, F1: "", F2: true}
                                _ = len(v2)
                                v35 -= -889
                                var v37 int = -1
                                if (1.5 <= 3.14159) {
                                    var v38 float64 = (float64(1.5) - ((1e10 + float64((-2.25 * (1.5 / float32(((-2.25 / -2.25) * (1e10 + ((float64((float32(((1e10 + -2.25) / 1.5)) / 1.0)) / (1.5 + float64((float32((-2.25 / -2.25)) * 1e10)))) + 3.14159)))))))) + 1.5))
                                    switch ("foo bar" + "") {
                                    case "\n\t":
                                        v37 = (fi33 << ((64) & 63))
                                    case "":
                                        fmt.Println(v37)
                                    case "foo bar":
                                        v35 += fi33
                                        switch "hello" {
                                        case "foo bar":
                                            v39 := GenS1{F0: 0}
                                            _ = v39
                                        case "hello":
                                            v4 -= 208
                                        case "\n\t":
                                            _ = len(v2)
                                        }
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                    default:
                                        var v40 uint16 = 652
                                        _ = v40
                                    }
                                    var v41 int8 = -37
                                    _ = v38
                                    _ = v41
                                }
                                _ = fi34
                                _ = v35
                                _ = v36
                                _ = v37
                            }
                            for fi42 := 0; fi42 < 6; fi42++ {
                                v4 = (fi33 ^ fi42)
                                if (-2.25 < -2.25) {
                                    var v43 uint = 227
                                    if len(v2) > 0 {
                                        v2[0] = 65535
                                    }
                                    _ = v43
                                } else {
                                    if true {
                                        if len(v2) > 0 {
                                            v2[0] = 2
                                        }
                                    } else {
                                        switch (^(fi42)) {
                                        case 9223372036854775807:
                                            switch "a" {
                                            case "":
                                                var v44 uint16 = 381
                                                var v45 uint16 = (^(v44))
                                                var v46 int64 = 1
                                                if (592 < 2) {
                                                    v4 += fi33
                                                    var v47 uint8 = 127
                                                    if ((v46 == int64(v4)) || (true || true)) {
                                                        switch ((("\n\t" + "") + "a") + "foo bar") {
                                                        case "\n\t":
                                                            switch ("\n\t" + (("" + "a") + "")) {
                                                            case "a":
                                                                var v48 uint = 42
                                                                _ = v48
                                                            case "\n\t":
                                                                if (v47 != 42) {
                                                                    var v49 float32 = (1e10 * (1e10 / 1.5))
                                                                    _ = v49
                                                                } else {
                                                                    switch (fi42 | v4) {
                                                                    case 2147483647:
                                                                        if (!(!(fi33 == fi42))) {
                                                                            v45 -= (v45 | 1)
                                                                            var v50 uint32 = 0
                                                                            v50 = (v50 / ((42) | 1))
                                                                            _ = v50
                                                                        }
                                                                    default:
                                                                        var v51 int64 = 9223372036854775807
                                                                        _ = v51
                                                                    }
                                                                    if true {
                                                                        v1["foo bar"] = (("hello" + "") + "foo bar")
                                                                    } else {
                                                                        fmt.Println(v4)
                                                                    }
                                                                    fmt.Println(fi33, v44)
                                                                }
                                                            case "":
                                                                var v52 uint64 = 0
                                                                _ = v52
                                                            default:
                                                                if false {
                                                                    switch (-(v4)) {
                                                                    case -748:
                                                                        if (226 == 2147483647) {
                                                                            var v53 int8 = genFn1((v4 ^ (v4 & 4294967295)))
                                                                            switch fi42 {
                                                                            case 32767:
                                                                                if false {
                                                                                    if false {
                                                                                        if (!true) {
                                                                                            switch "hello" {
                                                                                            case "foo bar":
                                                                                                if ((v53 << ((422) & 63)) < (v53 >> ((121) & 63))) {
                                                                                                    v46 += v46
                                                                                                }
                                                                                            case "a":
                                                                                                var v54 float64 = (3.14159 + 0.1)
                                                                                                _ = v54
                                                                                            case "hello":
                                                                                                v45 += (v45 * v45)
                                                                                            default:
                                                                                                fmt.Println(v47)
                                                                                            }
                                                                                            v47 -= v47
                                                                                        } else {
                                                                                            _ = len(v2)
                                                                                            v55 := GenS6{F0: 0, F1: 0, F2: 38}
                                                                                            var v56 int64 = 0
                                                                                            _ = v55
                                                                                            _ = v56
                                                                                        }
                                                                                    }
                                                                                    var v57 uint8 = (v47 >> ((255) & 63))
                                                                                    var v58 uint8 = (v47 * 127)
                                                                                    v59 := v1[""]
                                                                                    _ = v57
                                                                                    _ = v58
                                                                                    _ = v59
                                                                                }
                                                                            case -10:
                                                                                if true {
                                                                                    var v60 int16 = 1
                                                                                    _ = len(v1)
                                                                                    var v61 int16 = (v60 % ((v60) | 1))
                                                                                    v62 := v1["a"]
                                                                                    _ = v60
                                                                                    _ = v61
                                                                                    _ = v62
                                                                                } else {
                                                                                    switch ("hello" + "hello") {
                                                                                    case "":
                                                                                        if true {
                                                                                            v4 = fi42
                                                                                            if (611 != 0) {
                                                                                                var v63 int64 = -9223372036854775808
                                                                                                _ = v63
                                                                                            } else {
                                                                                                v64 := GenS3{F0: 629}
                                                                                                var v65 int16 = -1
                                                                                                _ = v64
                                                                                                _ = v65
                                                                                            }
                                                                                        } else {
                                                                                            _ = len(v3)
                                                                                        }
                                                                                    case "foo bar":
                                                                                        var v66 uint32 = 2147483647
                                                                                        _ = v66
                                                                                    default:
                                                                                        fmt.Println(fi42)
                                                                                    }
                                                                                    v47 += (v47 * v47)
                                                                                    v67 := GenS7{F0: 1, F1: 56}
                                                                                    _ = v67
                                                                                }
                                                                                fmt.Println(v47, v53, v53)
                                                                                fmt.Println(v47)
                                                                            case 0:
                                                                                v68 := GenS8{F0: ((!true) || true), F1: (fi42 / (((^(fi42))) | 1)), F2: (v44 ^ v44)}
                                                                                _ = v68
                                                                            }
                                                                            _ = v53
                                                                        }
                                                                    case 864:
                                                                        var v69 int8 = -1
                                                                        _ = v69
                                                                    default:
                                                                        v47 += uint8(v45)
                                                                    }
                                                                    v47 = v47
                                                                    if (!(124 < 65535)) {
                                                                        v44 = (v45 / ((272) | 1))
                                                                    }
                                                                }
                                                            }
                                                        case "":
                                                            fmt.Println(fi33, fi33)
                                                        case "foo bar":
                                                            v2 = append(v2, 65535)
                                                        }
                                                        var v70 float32 = 0.0
                                                        var v71 uint8 = 156
                                                        _ = v70
                                                        _ = v71
                                                    } else {
                                                        var v72 uint8 = 16
                                                        _ = v72
                                                    }
                                                    v45 += uint16(v47)
                                                    v45 = 42
                                                    delete(v3, "")
                                                    _ = v47
                                                } else {
                                                    v73 := GenS9{F0: 0}
                                                    v74 := GenS9{F0: 42}
                                                    _ = v73
                                                    _ = v74
                                                }
                                                _ = v44
                                                _ = v45
                                                _ = v46
                                            default:
                                                if len(v2) > 0 {
                                                    v2[0] = 2147483647
                                                }
                                            }
                                        case 42:
                                            var v75 int = fi33
                                            _ = v75
                                        case 0:
                                            fmt.Println(fi33, fi42)
                                        }
                                        v76 := GenS1{F0: -460}
                                        _ = v76
                                    }
                                    if (genFn0() < 42) {
                                        if (false || (-99 < v4)) {
                                            var v77 float32 = 1.5
                                            if false {
                                                var v78 uint16 = 951
                                                _ = v78
                                            }
                                            _ = v77
                                        }
                                        v79 := GenS9{F0: 1000}
                                        v4 -= (fi42 - (v4 & 4294967295))
                                        var v80 string = ("a" + "a")
                                        _ = v79
                                        _ = v80
                                    } else {
                                        if true {
                                            v81 := GenS10{F0: -9223372036854775808, F1: 0, F2: ""}
                                            if false {
                                                fmt.Println(v4, v4)
                                                v82 := GenS3{F0: 2}
                                                _ = v82
                                            } else {
                                                var v83 int32 = 763
                                                v4 = (fi42 & fi33)
                                                _ = v83
                                            }
                                            v84 := GenS10{F0: 833, F1: fi33, F2: ""}
                                            _ = v81
                                            _ = v84
                                        } else {
                                            switch v4 {
                                            case 32767:
                                                var v85 float32 = -2.25
                                                _ = v85
                                            default:
                                                v4 = (fi42 << ((1) & 63))
                                            }
                                            var v86 float64 = 1e10
                                            var v87 int = (fi33 % ((127) | 1))
                                            _ = v86
                                            _ = v87
                                        }
                                        var v88 int8 = 42
                                        _ = v88
                                    }
                                }
                                if true {
                                    v4 -= 127
                                    var v89 string = ("a" + "foo bar")
                                    _ = v89
                                } else {
                                    var v90 float32 = (((-2.25 * 3.14159) * 3.14159) * 3.14159)
                                    var v91 uint8 = 128
                                    v90 = v90
                                    _ = v90
                                    _ = v91
                                }
                                _ = fi42
                            }
                            v92 := GenS4{F0: 0, F1: 65535}
                            _ = v92
                        } else {
                            switch "foo bar" {
                            case "a":
                                v4 = fi33
                            case "":
                                v4 = (fi33 - fi33)
                            case "foo bar":
                                if len(v2) > 0 {
                                    v2[0] = 565
                                }
                            default:
                                v4 = (fi33 % ((128) | 1))
                            }
                        }
                        _ = fi33
                    }
                    v93 := GenS9{F0: 981}
                    _ = v93
                }
            } else {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                switch (v4 << ((192) & 63)) {
                case -2147483648:
                    for fi94 := 0; fi94 < 11; fi94++ {
                        for fi95 := 0; fi95 < 16; fi95++ {
                            var v96 int64 = -1
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            var v97 float64 = 1e10
                            _ = fi95
                            _ = v96
                            _ = v97
                        }
                        fmt.Println(v4, fi94)
                        _ = fi94
                    }
                case 42:
                    v4 = v4
                }
                v4 += v4
            }
            switch v4 {
            case -818:
                fmt.Println(v4)
            case 42:
                if false {
                    for fi98 := 0; fi98 < 13; fi98++ {
                        v99 := GenS2{F0: (fi98 & 42)}
                        v4 = 9223372036854775807
                        _ = fi98
                        _ = v99
                    }
                    fmt.Println(v4)
                }
            case 9223372036854775807:
                for fi100 := 0; fi100 < 11; fi100++ {
                    if false {
                        v4 = v4
                        v101 := GenS11{F0: -1}
                        _ = v101
                    }
                    _ = fi100
                }
            default:
                v4 = 1
            }
        }
        v102 := v3["foo bar"]
        v102 = 128
        _ = v102
    }
    for fi103 := 0; fi103 < 8; fi103++ {
        fmt.Println(v4, fi103)
        _ = fi103
    }
    var v104 uint16 = 2
    for fi105 := 0; fi105 < 12; fi105++ {
        if (false || false) {
            if (false && false) {
                v104 = (v104 ^ (v104 * v104))
                if false {
                    v104 = v104
                } else {
                    switch "\n\t" {
                    case "":
                        fmt.Println(v104, fi105, v104)
                    case "a":
                        v4 = (v4 << ((446) & 63))
                    }
                    for fi106 := 0; fi106 < 13; fi106++ {
                        v107 := GenS10{F0: 9223372036854775807, F1: -2147483648, F2: ((("" + ("\n\t" + "")) + "\n\t") + "\n\t")}
                        _ = fi106
                        _ = v107
                    }
                }
            }
            v108 := GenS12{F0: 1, F1: -564, F2: 839}
            _ = v108
        }
        if len(v2) > 0 {
            v2[0] = 1
        }
        _ = fi105
    }
    v104 = (v104 >> ((18446744073709551615) & 63))
    v4 -= 9223372036854775807
    var v109 int8 = 42
    fmt.Println(v104)
    if (true || false) {
        v109 = (v109 - v109)
        if ((^(v104)) == 255) {
            v110 := GenS4{F0: 1, F1: 1}
            _ = v110
        }
    }
    var v111 int32 = 2147483647
    var v112 bool = (2147483647 > 0)
    v104 += v104
    fmt.Println(v104)
    var v113 uint = 243
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v104:", v104)
    fmt.Println("v109:", v109)
    fmt.Println("v111:", v111)
    fmt.Println("v112:", v112)
    fmt.Println("v113:", v113)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v104
    _ = v109
    _ = v111
    _ = v112
    _ = v113
    _ = v2
    _ = v3
    _ = v4
}
