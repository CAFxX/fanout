package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint16; F1 uint32 }
type GenS3 struct { F0 uint8; F1 int64; F2 string }
type GenS4 struct { F0 int8; F1 string; F2 uint8 }
type GenS5 struct { F0 uint32; F1 int8; F2 uint16 }
type GenS6 struct { F0 uint; F1 string; F2 string }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 bool; F1 string }
type GenS9 struct { F0 int32; F1 uint; F2 uint }
type GenS10 struct { F0 uint32; F1 int8; F2 bool }
type GenS11 struct { F0 int; F1 uint8; F2 uint64 }
func genFn0(p0 int, p1 int) uint64 {
    var v7 float64 = 0.0
    p0 = p1
    for fi8 := 0; fi8 < 3; fi8++ {
        v9 := []int{p1, p1, -221}
        _ = fi8
        _ = v9
    }
    v10 := []int{p1, (p0 << ((255) & 63)), p0, p0}
    if (84 >= 128) {
        p0 += p1
        fmt.Println(v7, v7)
        for fi11 := 0; fi11 < 20; fi11++ {
            v12 := GenS1{F0: (p0 - (p0 / (((-(p1))) | 1)))}
            var v13 int = (fi11 + -1)
            v13 = v13
            _ = fi11
            _ = v12
            _ = v13
        }
    }
    _ = p0
    _ = p1
    _ = v10
    _ = v7
    return 236
}

func main() {
    v1 := []int{32768, 32767}
    var v2 float32 = (0.0 + 0.1)
    v3 := make(map[string]uint32)
    var v4 int8 = 100
    var v5 uint = 13
    var v6 uint16 = 2
    if (!(636 == 686)) {
        for fi14 := 0; fi14 < 11; fi14++ {
            var v15 int = fi14
            v15 = (fi14 * fi14)
            _ = fi14
            _ = v15
        }
    }
    if len(v1) > 0 {
        v1[0] = 128
    }
    for fi16 := 0; fi16 < 8; fi16++ {
        if (false || (!true)) {
            var v17 int = fi16
            switch (v17 % ((fi16) | 1)) {
            case 32767:
                v18 := GenS1{F0: v17}
                _ = v18
            default:
                if false {
                    v2 = v2
                    v4 += -73
                    var v19 bool = (!true)
                    _ = v19
                } else {
                    v3[(("foo bar" + ("hello" + ("hello" + ""))) + ("foo bar" + "\n\t"))] = 1
                    var v20 uint32 = 4294967295
                    _ = v20
                }
            }
            _ = v17
        } else {
            var v21 int64 = -983
            for fi22 := 0; fi22 < 2; fi22++ {
                switch "hello" {
                case "":
                    if true {
                        var v23 float64 = 1.5
                        _ = v23
                    } else {
                        v6 -= v6
                        var v24 int32 = -1
                        v5 -= v5
                        v24 += 42
                        _ = v24
                    }
                default:
                    v25 := GenS1{F0: (fi22 % ((fi22) | 1))}
                    _ = v25
                }
                v26 := GenS1{F0: (fi16 * -1)}
                _ = fi22
                _ = v26
            }
            if true {
                v27 := GenS2{F0: (^(v6)), F1: 419}
                _ = v27
            } else {
                v6 += (v6 * (v6 ^ v6))
                fmt.Println(v4, v6, v21)
                if true {
                    for fi28 := 0; fi28 < 3; fi28++ {
                        v29 := GenS1{F0: -758}
                        _ = fi28
                        _ = v29
                    }
                } else {
                    v30 := GenS3{F0: 235, F1: v21, F2: ("\n\t" + "a")}
                    for fi31 := 0; fi31 < 19; fi31++ {
                        v32 := GenS1{F0: -2147483648}
                        v5 += v5
                        v33 := GenS1{F0: (fi16 ^ fi16)}
                        switch ((("hello" + (("" + "foo bar") + "a")) + "foo bar") + (("foo bar" + "hello") + "foo bar")) {
                        case "a":
                            fmt.Println(v6, v4)
                        case "hello":
                            v34 := GenS2{F0: v6, F1: 163}
                            _ = v34
                        case "foo bar":
                            fmt.Println(fi16, v4)
                        }
                        v4 = (v4 << (((v5 * v5)) & 63))
                        _ = len(v1)
                        var v35 bool = (!true)
                        _ = fi31
                        _ = v32
                        _ = v33
                        _ = v35
                    }
                    _ = v30
                }
            }
            _ = v21
        }
        _ = fi16
    }
    for fi36 := 0; fi36 < 6; fi36++ {
        for fi37 := 0; fi37 < 20; fi37++ {
            v38 := GenS3{F0: 128, F1: 9223372036854775807, F2: ""}
            _ = fi37
            _ = v38
        }
        fmt.Println(v5, fi36)
        _ = fi36
    }
    v5 -= v5
    fmt.Println(v5)
    for fi39 := 0; fi39 < 5; fi39++ {
        v4 = v4
        v40 := GenS4{F0: (^(v4)), F1: "a", F2: 42}
        v6 = 32767
        _ = fi39
        _ = v40
    }
    for fi41 := 0; fi41 < 8; fi41++ {
        v42 := GenS5{F0: 1, F1: v4, F2: v6}
        v6 -= v6
        v2 += 3.14159
        _ = fi41
        _ = v42
    }
    if (127 != 128) {
        if true {
            if (32767 != -715) {
                fmt.Println(v2, v2, v5)
                fmt.Println(v5, v2, v2)
            }
            var v43 int = -688
            _ = v43
        } else {
            v44 := GenS6{F0: v5, F1: "hello", F2: "hello"}
            _ = v44
        }
    }
    var v45 uint64 = 4294967295
    switch 127 {
    case 127:
        var v46 int64 = 1
        _ = v46
    case 0:
        if len(v1) > 0 {
            v1[0] = 0
        }
    default:
        for fi47 := 0; fi47 < 4; fi47++ {
            var v48 float64 = (-2.25 / 3.14159)
            _ = fi47
            _ = v48
        }
    }
    v49 := v3["foo bar"]
    v50 := GenS7{F0: 0}
    for fi51 := 0; fi51 < 17; fi51++ {
        v4 -= v4
        _ = fi51
    }
    var v52 float32 = (float32(-2.25) + (v2 + (v2 * v2)))
    if ((!false) && ((v4 ^ v4) < v4)) {
        v45 -= (v45 * v45)
        switch "" {
        case "a":
            if false {
                v5 -= v5
                fmt.Println(v6, v2)
            }
        default:
            switch "\n\t" {
            case "hello":
                var v53 int64 = -7
                _ = v53
            case "\n\t":
                if (false && true) {
                    var v54 uint = 1
                    _ = v54
                }
            case "a":
                switch "\n\t" {
                case "\n\t":
                    if (!true) {
                        var v55 uint32 = 2147483648
                        _ = v55
                    } else {
                        if ((v4 / (((v4 & v4)) | 1)) >= (v4 ^ (v4 % ((v4) | 1)))) {
                            if len(v1) > 0 {
                                v1[0] = -9223372036854775808
                            }
                            var v56 uint64 = (v45 % (((v45 >> ((193) & 63))) | 1))
                            v57 := GenS5{F0: (v49 % (((v49 / (((v49 << ((v5) & 63))) | 1))) | 1)), F1: v4, F2: v6}
                            _ = v56
                            _ = v57
                        }
                    }
                }
            }
        }
    } else {
        if true {
            for fi58 := 0; fi58 < 7; fi58++ {
                if true {
                    v59 := GenS5{F0: 1, F1: (v4 + (v4 >> ((v5) & 63))), F2: 952}
                    switch (fi58 << ((980) & 63)) {
                    case -1:
                        v60 := GenS1{F0: fi58}
                        _ = v60
                    default:
                        for fi61 := 0; fi61 < 18; fi61++ {
                            var v62 float64 = (float64(v52) / (1.5 + (3.14159 + (1.5 - 1.5))))
                            if true {
                                var v63 float32 = (float32(v62) + v52)
                                _ = v63
                            }
                            _ = fi61
                            _ = v62
                        }
                    }
                    _ = v59
                } else {
                    var v64 uint16 = v6
                    v3["\n\t"] = v49
                    v5 = v5
                    _ = v64
                }
                _ = fi58
            }
            if len(v1) > 0 {
                v1[0] = 42
            }
            var v65 int8 = v4
            _ = v65
        }
        if (true && (100 != 127)) {
            if false {
                var v66 uint16 = (v6 + v6)
                _ = v66
            }
            if (0.1 != 0.1) {
                v49 = uint32(v45)
                v67 := GenS6{F0: v5, F1: "foo bar", F2: ""}
                v3[("foo bar" + "foo bar")] = (v49 >> ((v5) & 63))
                _ = v67
            } else {
                v1 = append(v1, 32768)
                if len(v1) > 0 {
                    v1[0] = -1
                }
                if (((1.5 / (((v2 * (v2 * v52)) * (v52 / v52)) * 1e10)) > float32((0.0 + 0.0))) || true) {
                    if true {
                        if false {
                            var v68 uint = v5
                            for fi69 := 0; fi69 < 13; fi69++ {
                                if true {
                                    v1 = append(v1, (fi69 & fi69))
                                    var v70 int8 = v4
                                    fmt.Println(v52, v4)
                                    _ = v70
                                } else {
                                    var v71 uint = (v68 & (v68 / (((^(v68))) | 1)))
                                    if false {
                                        delete(v3, ("" + ""))
                                        v49 += v49
                                    } else {
                                        var v72 int8 = v4
                                        fmt.Println(fi69, v49, fi69)
                                        _ = v72
                                    }
                                    _ = v71
                                }
                                _ = fi69
                            }
                            _ = v68
                        }
                    } else {
                        fmt.Println(v49)
                        for fi73 := 0; fi73 < 12; fi73++ {
                            if (!false) {
                                for fi74 := 0; fi74 < 6; fi74++ {
                                    if (false && (1000 > 1000)) {
                                        v5 -= v5
                                        var v75 float32 = ((v2 + (v2 + float32((0.1 * -2.25)))) - float32(((1.5 / (3.14159 * (3.14159 * (-2.25 - ((0.1 / 0.1) * (1.5 / 1.0)))))) - 0.1)))
                                        _ = v75
                                    } else {
                                        if ((fi74 & int(v45)) >= (fi74 | (^(fi73)))) {
                                            v2 += 3.14159
                                        } else {
                                            switch fi73 {
                                            case 32767:
                                                var v76 uint16 = (v6 % ((255) | 1))
                                                _ = v76
                                            }
                                            v77 := GenS8{F0: (v6 >= (v6 * v6)), F1: "foo bar"}
                                            _ = v77
                                        }
                                        switch (fi74 & fi74) {
                                        case -9223372036854775808:
                                            v49 = uint32(v45)
                                        }
                                        if (!true) {
                                            v2 += v2
                                            if ((true || false) && false) {
                                                switch "\n\t" {
                                                case "\n\t":
                                                    switch "hello" {
                                                    case "a":
                                                        if false {
                                                            var v78 uint8 = 42
                                                            v79 := GenS7{F0: -965}
                                                            v6 -= 255
                                                            _ = v78
                                                            _ = v79
                                                        } else {
                                                            var v80 uint = 65535
                                                            if (9223372036854775807 != int64(v45)) {
                                                                var v81 int16 = 1
                                                                _ = v81
                                                            }
                                                            _ = v80
                                                        }
                                                    }
                                                case "hello":
                                                    if false {
                                                        v49 = (v49 >> ((v5) & 63))
                                                        v2 = v2
                                                    } else {
                                                        if false {
                                                            var v82 float32 = v2
                                                            var v83 int8 = (v4 ^ (v4 << ((2) & 63)))
                                                            switch (^(fi74)) {
                                                            case 128:
                                                                switch (fi73 + -2147483648) {
                                                                case -9223372036854775808:
                                                                    var v84 int = (fi74 >> (((^(v5))) & 63))
                                                                    _ = v84
                                                                default:
                                                                    switch (fi74 + fi74) {
                                                                    case -223:
                                                                        if len(v1) > 0 {
                                                                            v1[0] = fi73
                                                                        }
                                                                    case -551:
                                                                        v45 = genFn0(fi73, int(v6))
                                                                    case 128:
                                                                        v6 = v6
                                                                    }
                                                                }
                                                            case -888:
                                                                var v85 uint = (v5 & 4294967295)
                                                                _ = v85
                                                            default:
                                                                fmt.Println(v2, v52)
                                                            }
                                                            _ = v82
                                                            _ = v83
                                                        }
                                                    }
                                                case "foo bar":
                                                    v86 := GenS9{F0: -1, F1: v5, F2: 255}
                                                    _ = v86
                                                default:
                                                    v3["a"] = (v49 & (^(v49)))
                                                }
                                            } else {
                                                v4 = int8(fi73)
                                                var v87 int = fi74
                                                _ = v87
                                            }
                                        } else {
                                            var v88 bool = true
                                            fmt.Println(v45)
                                            var v89 bool = v88
                                            v45 = (v45 * 18446744073709551615)
                                            _ = v88
                                            _ = v89
                                        }
                                    }
                                    _ = fi74
                                }
                                v90 := GenS2{F0: v6, F1: (v49 >> (((v5 % (((v5 % (((v5 + (v5 / ((uint(v6)) | 1)))) | 1))) | 1))) & 63))}
                                for fi91 := 0; fi91 < 6; fi91++ {
                                    switch ("foo bar" + "hello") {
                                    case "\n\t":
                                        v49 = 2147483648
                                    case "a":
                                        v6 = v6
                                    case "":
                                        if (false || (v5 > v5)) {
                                            if true {
                                                var v92 uint8 = 2
                                                _ = v92
                                            } else {
                                                if (!true) {
                                                    var v93 uint64 = genFn0(fi91, -9223372036854775808)
                                                    v49 += (v49 / (((^(v49))) | 1))
                                                    v93 -= genFn0(int(v5), (fi91 | -9223372036854775808))
                                                    if len(v1) > 0 {
                                                        v1[0] = 128
                                                    }
                                                    var v94 float32 = ((v2 * v2) * (float32((1.5 + 1.5)) / (((v52 / (v52 / (((1e10 + 0.0) + float32(float64(float32(0.0)))) * v2))) - v52) - float32((1.5 * 1.5)))))
                                                    _ = v93
                                                    _ = v94
                                                } else {
                                                    if ((3.14159 / 0.1) >= 0.0) {
                                                        v52 = v2
                                                        v52 = (((v52 * (float32(((((0.1 / -2.25) + 1.5) * (-2.25 + 1.5)) + 0.1)) - (v2 / (v2 + v2)))) * float32(3.14159)) * v52)
                                                    }
                                                    v49 += (v49 % ((1) | 1))
                                                    fmt.Println(fi91, v6, fi73)
                                                }
                                            }
                                            v95 := GenS10{F0: (v49 % ((707) | 1)), F1: v4, F2: (32767 < 1000)}
                                            var v96 int = -9223372036854775808
                                            v97 := GenS7{F0: 42}
                                            _ = v95
                                            _ = v96
                                            _ = v97
                                        } else {
                                            v98 := GenS9{F0: 0, F1: (v5 | v5), F2: (v5 >> (((v5 + uint(v49))) & 63))}
                                            _ = v98
                                        }
                                    default:
                                        fmt.Println(v5, v52, v2)
                                    }
                                    _ = fi91
                                }
                                _ = v90
                            } else {
                                v52 -= ((v52 + (v2 * 1e10)) / 1e10)
                                v6 = (v6 / ((v6) | 1))
                            }
                            v99 := GenS10{F0: (v49 - (v49 % ((2147483648) | 1))), F1: 42, F2: (true && true)}
                            fmt.Println(v4)
                            _ = fi73
                            _ = v99
                        }
                    }
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                }
            }
            for fi100 := 0; fi100 < 2; fi100++ {
                switch -557 {
                case 32767:
                    var v101 float32 = float32(0.1)
                    _ = v101
                case -1:
                    v6 = 533
                default:
                    fmt.Println(v2, v2, v5)
                }
                switch (fi100 | (fi100 / ((127) | 1))) {
                case 952:
                    for fi102 := 0; fi102 < 13; fi102++ {
                        v103 := GenS5{F0: (^(v49)), F1: v4, F2: v6}
                        fmt.Println(fi102, fi100, v49)
                        _ = fi102
                        _ = v103
                    }
                case 127:
                    if (!false) {
                        v5 = 1
                        v1 = append(v1, int(v5))
                    } else {
                        for fi104 := 0; fi104 < 17; fi104++ {
                            if ((v5 == 1) || false) {
                                var v105 int64 = -39
                                v2 += v52
                                if false {
                                    if ("hello" != "foo bar") {
                                        var v106 bool = false
                                        if true {
                                            delete(v3, "hello")
                                        } else {
                                            v6 += v6
                                            fmt.Println(v105, v45)
                                        }
                                        _ = v106
                                    } else {
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        _ = len(v3)
                                    }
                                }
                                v107 := v3[""]
                                _ = v105
                                _ = v107
                            } else {
                                switch ("\n\t" + "\n\t") {
                                case "":
                                    v108 := GenS2{F0: 501, F1: (v49 % (((^(v49))) | 1))}
                                    _ = v108
                                }
                                if (42 > 42) {
                                    switch (("hello" + "a") + "a") {
                                    case "foo bar":
                                        v49 = 763
                                    case "a":
                                        delete(v3, "a")
                                    }
                                    v49 += 4294967295
                                } else {
                                    fmt.Println(v49, v52, v6)
                                    v45 += (v45 << ((v5) & 63))
                                }
                            }
                            if (true && ((!(42 == 42)) && (1 > 0))) {
                                fmt.Println(v5, v2)
                                switch 9223372036854775807 {
                                case -9223372036854775808:
                                    v49 += (v49 / ((4294967295) | 1))
                                case 128:
                                    switch (fi100 - 32767) {
                                    case -658:
                                        if false {
                                            var v109 string = ("\n\t" + "\n\t")
                                            if (v5 == v5) {
                                                _ = len(v1)
                                                fmt.Println(v109, v49)
                                                var v110 int8 = v4
                                                _ = v110
                                            } else {
                                                v52 -= (float32(float64(float32((((float64((1.5 / 3.14159)) + (float64(v2) * (0.0 + 3.14159))) * 1.5) * 1e10)))) * (float32(-2.25) + -2.25))
                                            }
                                            v2 = (((3.14159 - -2.25) - v52) * 0.0)
                                            _ = v109
                                        }
                                    case 32768:
                                        v45 = (v45 | (^(v45)))
                                    case 32767:
                                        if true {
                                            var v111 uint8 = 255
                                            switch 88 {
                                            case -715:
                                                fmt.Println(fi100, v52, v45)
                                            case -1:
                                                v112 := GenS11{F0: 127, F1: (v111 / ((uint8(v49)) | 1)), F2: (v45 + v45)}
                                                _ = v112
                                            case 128:
                                                if true {
                                                    var v113 int = 32768
                                                    v114 := GenS6{F0: 975, F1: (("" + "hello") + ("hello" + "a")), F2: "\n\t"}
                                                    _ = v113
                                                    _ = v114
                                                }
                                            }
                                            _ = v111
                                        } else {
                                            switch fi100 {
                                            case -9223372036854775808:
                                                v115 := GenS4{F0: v4, F1: "a", F2: 171}
                                                _ = v115
                                            }
                                        }
                                    default:
                                        _ = len(v1)
                                    }
                                case 32767:
                                    v5 -= (v5 & (v5 * (v5 >> ((v5) & 63))))
                                default:
                                    v116 := GenS6{F0: 0, F1: "", F2: ""}
                                    _ = v116
                                }
                                var v117 uint8 = 127
                                _ = v117
                            } else {
                                v5 = (v5 % ((v5) | 1))
                            }
                            _ = fi104
                        }
                    }
                case 32767:
                    _ = len(v1)
                default:
                    v118 := GenS2{F0: (v6 >> (((v5 & v5)) & 63)), F1: v49}
                    _ = v118
                }
                _ = fi100
            }
        }
    }
    v119 := GenS9{F0: 111, F1: v5, F2: (v5 - v5)}
    for fi120 := 0; fi120 < 16; fi120++ {
        var v121 int64 = 885
        _ = fi120
        _ = v121
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v52:", v52)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v119
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v49
    _ = v5
    _ = v50
    _ = v52
    _ = v6
}
