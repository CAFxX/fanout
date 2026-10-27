package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16 }
type GenS2 struct { F0 int8; F1 string }
type GenS3 struct { F0 uint16; F1 int32; F2 int8 }
type GenS4 struct { F0 string; F1 int32 }
type GenS5 struct { F0 int64; F1 int8; F2 uint }
type GenS6 struct { F0 bool; F1 uint64; F2 uint16 }
type GenS7 struct { F0 int8; F1 uint; F2 int32 }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 uint; F1 int32 }
type GenS10 struct { F0 int }
type GenS11 struct { F0 uint }
type GenS12 struct { F0 uint32 }
type GenS13 struct { F0 uint16; F1 int64 }
func genFn0(p0 int64, p1 int8) int16 {
    var v7 uint16 = 2
    p1 = (-(p1))
    p0 += p0
    var v8 uint = 0
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 32767
}
func genFn1(p0 int16, p1 int, p2 uint8) uint64 {
    var v9 uint16 = 688
    var v10 int8 = 42
    v11 := GenS1{F0: uint16(p2), F1: 2}
    for fi12 := 0; fi12 < 13; fi12++ {
        v13 := []int{p1, fi12}
        if true {
            for fi14 := 0; fi14 < 15; fi14++ {
                var v15 int16 = (p0 % ((p0) | 1))
                fmt.Println(v10, p2)
                v16 := GenS1{F0: (^(v9)), F1: (v9 / (((v9 / ((730) | 1))) | 1))}
                if false {
                    p2 -= p2
                } else {
                    fmt.Println(p2, fi12, v9)
                    v13 = append(v13, fi12)
                    var v17 float64 = ((3.14159 + 3.14159) + float64((3.14159 * 0.0)))
                    _ = v17
                }
                _ = fi14
                _ = v15
                _ = v16
            }
        }
        v9 += 0
        _ = fi12
        _ = v13
    }
    if (-423 != 1) {
        for fi18 := 0; fi18 < 15; fi18++ {
            v19 := []int{}
            _ = fi18
            _ = v19
        }
        var v20 bool = (!true)
        if ((false || false) || v20) {
            p0 = (p0 / ((genFn0(-9223372036854775808, int8(p2))) | 1))
        } else {
            v10 = -1
            for fi21 := 0; fi21 < 8; fi21++ {
                for fi22 := 0; fi22 < 13; fi22++ {
                    v20 = v20
                    _ = fi22
                }
                switch (-(fi21)) {
                case -143:
                    v10 += v10
                case -789:
                    if (v20 && (!false)) {
                        switch (fi21 * -103) {
                        case 32767:
                            v23 := GenS1{F0: v9, F1: 42}
                            _ = v23
                        case 42:
                            v24 := []int{}
                            _ = v24
                        }
                    } else {
                        var v25 uint64 = 1
                        _ = v25
                    }
                }
                var v26 uint8 = (p2 / ((p2) | 1))
                _ = fi21
                _ = v26
            }
        }
        _ = v20
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v9
    return 114
}

func main() {
    var v1 bool = true
    var v2 int16 = -657
    var v3 uint64 = 4294967295
    v4 := make(map[uint32]uint32)
    var v5 bool = (2 < (v3 / (((v3 * (v3 + (^(v3))))) | 1)))
    v6 := []int{-2147483648, 9223372036854775807, 842, 0}
    switch "hello" {
    case "a":
        if (!false) {
            var v27 int8 = 92
            _ = v27
        } else {
            for fi28 := 0; fi28 < 11; fi28++ {
                v6 = append(v6, fi28)
                var v29 string = "a"
                v30 := GenS2{F0: 127, F1: "foo bar"}
                _ = fi28
                _ = v29
                _ = v30
            }
            var v31 int = 445
            _ = v31
        }
    }
    v32 := GenS2{F0: 42, F1: ("hello" + "foo bar")}
    v5 = ((929 < (v3 / ((2) | 1))) || (566 == 268))
    if (("" + ("" + "foo bar")) <= "\n\t") {
        switch (("" + "") + "foo bar") {
        case "a":
            var v33 uint16 = 0
            _ = v33
        default:
            if false {
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                for fi34 := 0; fi34 < 13; fi34++ {
                    var v35 string = ""
                    v36 := GenS3{F0: 988, F1: -1, F2: -1}
                    _ = fi34
                    _ = v35
                    _ = v36
                }
                if v5 {
                    var v37 int64 = 0
                    var v38 int32 = int32(v37)
                    _ = v37
                    _ = v38
                }
            } else {
                switch ("hello" + "hello") {
                case "foo bar":
                    for fi39 := 0; fi39 < 10; fi39++ {
                        fmt.Println(fi39, v5)
                        _ = fi39
                    }
                case "hello":
                    var v40 string = "a"
                    _ = v40
                default:
                    for fi41 := 0; fi41 < 5; fi41++ {
                        v2 = (^(v2))
                        if v1 {
                            _ = len(v4)
                        } else {
                            for fi42 := 0; fi42 < 5; fi42++ {
                                var v43 int = 127
                                v6 = append(v6, 42)
                                fmt.Println(v3, fi42, v43)
                                _ = fi42
                                _ = v43
                            }
                            v3 += 1
                            for fi44 := 0; fi44 < 6; fi44++ {
                                var v45 int64 = -356
                                var v46 uint = 0
                                if (0 > -1) {
                                    var v47 float64 = 0.0
                                    v3 = 427
                                    _ = v47
                                } else {
                                    var v48 bool = v1
                                    v48 = false
                                    var v49 bool = (274 < 0)
                                    var v50 int8 = -107
                                    if ("" < ("foo bar" + ((("\n\t" + ("foo bar" + (("a" + (((("a" + (("foo bar" + ("\n\t" + "foo bar")) + "a")) + ("a" + "hello")) + ("" + "hello")) + "")) + "a"))) + "hello") + "\n\t"))) {
                                        if len(v6) > 0 {
                                            v6[0] = -9223372036854775808
                                        }
                                    }
                                    _ = v48
                                    _ = v49
                                    _ = v50
                                }
                                _ = fi44
                                _ = v45
                                _ = v46
                            }
                        }
                        _ = fi41
                    }
                }
            }
        }
        v1 = (!v1)
        v51 := GenS3{F0: 1, F1: -730, F2: 42}
        _ = v51
    }
    var v52 uint = 42
    v1 = v5
    v53 := GenS1{F0: 65535, F1: 255}
    switch ("a" + "hello") {
    case "hello":
        v54 := GenS4{F0: "a", F1: -583}
        _ = v54
    case "foo bar":
        v55 := GenS5{F0: 0, F1: -110, F2: 65535}
        _ = v55
    case "":
        v52 -= 780
    default:
        _ = len(v4)
    }
    fmt.Println(v1)
    v56 := GenS6{F0: v1, F1: (v3 ^ 42), F2: 42}
    v52 -= (v52 & v52)
    if v5 {
        for fi57 := 0; fi57 < 1; fi57++ {
            switch fi57 {
            case 4294967295:
                if ((v2 | (v2 / (((v2 >> ((v52) & 63))) | 1))) < (v2 & v2)) {
                    switch fi57 {
                    case 603:
                        switch (-(fi57)) {
                        case 42:
                            var v58 int64 = -9223372036854775808
                            _ = v58
                        case 1:
                            fmt.Println(v2)
                        default:
                            fmt.Println(v52, v1, v52)
                        }
                    }
                    v2 -= (v2 | (v2 << (((^(v52))) & 63)))
                    if ("" >= "") {
                        var v59 int64 = -991
                        _ = v59
                    } else {
                        fmt.Println(v5)
                    }
                } else {
                    if (!((((!(false || v1)) || (-1 != 821)) || true) && (1 == 2147483647))) {
                        var v60 uint = (v52 ^ 42)
                        for fi61 := 0; fi61 < 14; fi61++ {
                            delete(v4, 998)
                            if (v3 == 821) {
                                switch (fi57 & fi57) {
                                case 0:
                                    switch "foo bar" {
                                    case "foo bar":
                                        if v1 {
                                            if true {
                                                if true {
                                                    v1 = (true || v5)
                                                    switch (fi57 | 127) {
                                                    case 127:
                                                        if v5 {
                                                            if ((fi61 < -478) || true) {
                                                                if false {
                                                                    v5 = v1
                                                                    fmt.Println(fi61, v5, fi57)
                                                                    fmt.Println(v52)
                                                                }
                                                                if ((65535 < 186) && v1) {
                                                                    switch (fi57 % (((^(fi57))) | 1)) {
                                                                    case 42:
                                                                        v3 = v3
                                                                    case 4294967295:
                                                                        v6 = append(v6, 127)
                                                                        if v5 {
                                                                            v62 := GenS7{F0: 42, F1: 42, F2: -1}
                                                                            _ = v62
                                                                        } else {
                                                                            var v63 int = fi61
                                                                            var v64 bool = (!v5)
                                                                            _ = v63
                                                                            _ = v64
                                                                        }
                                                                        v60 += v52
                                                                    case 1:
                                                                        var v65 string = "\n\t"
                                                                        _ = v65
                                                                    default:
                                                                        if (false || (42 >= 4294967295)) {
                                                                            switch 127 {
                                                                            case -1:
                                                                                fmt.Println(v2)
                                                                            case 42:
                                                                                v2 = v2
                                                                            }
                                                                            var v66 int = fi57
                                                                            var v67 float32 = ((-2.25 / 0.1) - (((float32(0.1) - ((3.14159 + (float32(0.0) - 3.14159)) * 0.0)) * 1e10) - float32(1e10)))
                                                                            _ = v66
                                                                            _ = v67
                                                                        }
                                                                    }
                                                                } else {
                                                                    if true {
                                                                        if (fi57 >= fi61) {
                                                                            fmt.Println(fi61)
                                                                        } else {
                                                                            var v68 int = fi57
                                                                            var v69 int32 = 10
                                                                            _ = v68
                                                                            _ = v69
                                                                        }
                                                                    } else {
                                                                        if v1 {
                                                                            var v70 uint64 = (^(v3))
                                                                            var v71 float64 = 3.14159
                                                                            if v5 {
                                                                                v4[2147483647] = 948
                                                                                if ((!false) && true) {
                                                                                    var v72 uint64 = v70
                                                                                    if v1 {
                                                                                        var v73 int64 = -1
                                                                                        v74 := GenS8{F0: 255}
                                                                                        if len(v6) >= 2 {
                                                                                            v6 = v6[1:len(v6)]
                                                                                        }
                                                                                        _ = v73
                                                                                        _ = v74
                                                                                    } else {
                                                                                        var v75 uint8 = 15
                                                                                        var v76 float32 = 1e10
                                                                                        _ = v75
                                                                                        _ = v76
                                                                                    }
                                                                                    _ = v72
                                                                                }
                                                                            } else {
                                                                                if (("\n\t" + "hello") == "a") {
                                                                                    v3 = (v3 & genFn1(v2, fi61, 1))
                                                                                    var v77 uint8 = 123
                                                                                    _ = v77
                                                                                }
                                                                                if v5 {
                                                                                    v78 := GenS5{F0: 9223372036854775807, F1: 100, F2: 605}
                                                                                    v70 += 18446744073709551615
                                                                                    v6 = append(v6, (fi57 / (((fi61 >> (((v60 - 0)) & 63))) | 1)))
                                                                                    _ = v78
                                                                                } else {
                                                                                    v1 = (65535 >= 32767)
                                                                                    var v79 uint16 = 950
                                                                                    v80 := v4[0]
                                                                                    _ = v79
                                                                                    _ = v80
                                                                                }
                                                                            }
                                                                            _ = v70
                                                                            _ = v71
                                                                        } else {
                                                                            v1 = (!(v1 || (!v1)))
                                                                            v81 := GenS7{F0: 0, F1: v52, F2: 49}
                                                                            _ = v81
                                                                        }
                                                                        var v82 uint16 = 1
                                                                        fmt.Println(v1, v82)
                                                                        _ = v82
                                                                    }
                                                                }
                                                                v6 = append(v6, 910)
                                                            } else {
                                                                var v83 int64 = -843
                                                                v60 = v52
                                                                _ = v83
                                                            }
                                                        }
                                                        var v84 uint = (v52 + (v52 % (((v60 << (((v52 << (((^(v52))) & 63))) & 63))) | 1)))
                                                        var v85 float64 = -2.25
                                                        v5 = (135 != -1)
                                                        _ = v84
                                                        _ = v85
                                                    case 42:
                                                        v60 -= v60
                                                    }
                                                    v4[972] = uint32(v3)
                                                } else {
                                                    fmt.Println(fi57, v52)
                                                    if (0 >= 127) {
                                                        v86 := GenS9{F0: v52, F1: -1}
                                                        if (888 < 2147483647) {
                                                            if (v1 && false) {
                                                                switch "a" {
                                                                case "\n\t":
                                                                    v6 = append(v6, 0)
                                                                default:
                                                                    v6 = append(v6, 32767)
                                                                }
                                                            } else {
                                                                v52 -= 0
                                                                v52 = (v60 & (v52 | (^(v52))))
                                                                v5 = (0 <= -9223372036854775808)
                                                            }
                                                        } else {
                                                            v87 := GenS1{F0: 32767, F1: 880}
                                                            v52 = v52
                                                            v88 := GenS7{F0: 127, F1: 154, F2: 924}
                                                            _ = v87
                                                            _ = v88
                                                        }
                                                        _ = v86
                                                    } else {
                                                        v60 -= v60
                                                        var v89 float32 = 1e10
                                                        _ = v89
                                                    }
                                                    v5 = ((v60 % ((v52) | 1)) > (v60 >> (((v60 - v60)) & 63)))
                                                }
                                                delete(v4, 0)
                                                v90 := GenS5{F0: 0, F1: 100, F2: v60}
                                                _ = v90
                                            } else {
                                                var v91 int8 = 50
                                                _ = v91
                                            }
                                            v92 := GenS10{F0: fi57}
                                            fmt.Println(fi61, v3, v60)
                                            _ = v92
                                        }
                                    case "hello":
                                        if len(v6) >= 2 {
                                            v6 = v6[1:len(v6)]
                                        }
                                        v2 = (v2 / (((v2 + (v2 | (v2 / (((v2 ^ v2)) | 1))))) | 1))
                                        var v93 uint32 = 685
                                        v1 = false
                                        _ = v93
                                    case "":
                                        var v94 uint8 = 2
                                        _ = v94
                                    default:
                                        if v1 {
                                            fmt.Println(v52)
                                        }
                                    }
                                default:
                                    fmt.Println(fi61, v2)
                                }
                                v95 := GenS1{F0: 47, F1: 643}
                                _ = v95
                            }
                            if len(v6) > 0 {
                                v6[0] = (fi57 / (((-(fi57))) | 1))
                            }
                            _ = fi61
                        }
                        fmt.Println(v3, v3, v3)
                        _ = v60
                    }
                    var v96 int16 = (v2 & (^(v2)))
                    _ = v96
                }
            }
            _ = fi57
        }
    } else {
        for fi97 := 0; fi97 < 6; fi97++ {
            var v98 uint8 = 237
            _ = len(v6)
            v99 := GenS9{F0: v52, F1: -847}
            _ = fi97
            _ = v98
            _ = v99
        }
        if (!(v5 && ((true && v1) || false))) {
            switch 423 {
            case 1:
                v100 := GenS4{F0: "\n\t", F1: 0}
                _ = v100
            }
        }
        if true {
            switch -2147483648 {
            case -355:
                _ = len(v4)
            }
        } else {
            v101 := GenS7{F0: 0, F1: 290, F2: 42}
            v1 = v1
            if (42 > 490) {
                fmt.Println(v1, v52)
            }
            _ = v101
        }
    }
    switch ("hello" + "foo bar") {
    case "":
        v5 = v5
    }
    v102 := GenS11{F0: (v52 - (v52 >> ((v52) & 63)))}
    var v103 bool = ((v3 >> ((v52) & 63)) <= (^(v3)))
    v104 := GenS12{F0: uint32(v3)}
    var v105 uint8 = 128
    if v5 {
        v106 := GenS13{F0: uint16(v105), F1: -9223372036854775808}
        if (!v103) {
            v107 := GenS3{F0: uint16(v52), F1: -1, F2: 42}
            _ = v107
        }
        switch "a" {
        case "a":
            if ("" >= "foo bar") {
                v52 = 18446744073709551615
            }
        case "":
            fmt.Println(v105)
        default:
            v6 = append(v6, -9223372036854775808)
        }
        _ = v106
    }
    fmt.Println("v1:", v1)
    fmt.Println("v103:", v103)
    fmt.Println("v105:", v105)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v52:", v52)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v102
    _ = v103
    _ = v104
    _ = v105
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v52
    _ = v53
    _ = v56
    _ = v6
}
