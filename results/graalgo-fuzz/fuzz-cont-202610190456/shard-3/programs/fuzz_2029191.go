package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint8; F1 int16 }
type GenS3 struct { F0 int16; F1 int64 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 bool; F1 int32 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 uint64; F1 string }
type GenS10 struct { F0 bool; F1 uint16; F2 string }
type GenS11 struct { F0 int8; F1 bool; F2 uint8 }
type GenS12 struct { F0 uint64; F1 int16; F2 string }
type GenS13 struct { F0 int; F1 int32; F2 uint32 }
type GenS14 struct { F0 uint32; F1 bool; F2 int32 }
func genFn0() uint32 {
    var v5 float32 = 1.5
    var v6 uint32 = 890
    for fi7 := 0; fi7 < 17; fi7++ {
        if (v5 > (v5 + (v5 + v5))) {
            if false {
                switch ("a" + ("" + ((("foo bar" + "") + "\n\t") + "a"))) {
                case "hello":
                    if (4294967295 >= 0) {
                        if (!true) {
                            v8 := []int{fi7, fi7}
                            var v9 uint32 = 899
                            if false {
                                v5 -= v5
                                var v10 int16 = -1
                                _ = v10
                            } else {
                                for fi11 := 0; fi11 < 5; fi11++ {
                                    v12 := GenS1{F0: 379}
                                    if (true || false) {
                                        v5 = v5
                                        if len(v8) > 0 {
                                            v8[0] = (fi11 ^ fi7)
                                        }
                                    }
                                    v13 := GenS1{F0: 710}
                                    _ = fi11
                                    _ = v12
                                    _ = v13
                                }
                                switch (("" + (("foo bar" + "") + "\n\t")) + "hello") {
                                case "hello":
                                    var v14 uint = 255
                                    _ = v14
                                }
                            }
                            _ = v8
                            _ = v9
                        } else {
                            if false {
                                var v15 uint32 = v6
                                if true {
                                    if true {
                                        if (42 != 9223372036854775807) {
                                            var v16 int32 = 0
                                            _ = v16
                                        } else {
                                            v17 := GenS2{F0: 2, F1: 42}
                                            v15 += (^(v15))
                                            _ = v17
                                        }
                                    }
                                    v18 := GenS3{F0: -164, F1: -9223372036854775808}
                                    v6 -= (v15 & (v6 << ((2) & 63)))
                                    _ = v18
                                } else {
                                    switch "hello" {
                                    case "foo bar":
                                        var v19 int8 = -128
                                        _ = v19
                                    case "hello":
                                        v20 := GenS2{F0: 127, F1: 590}
                                        _ = v20
                                    }
                                    for fi21 := 0; fi21 < 4; fi21++ {
                                        var v22 float64 = 1e10
                                        fmt.Println(v6, fi7, v5)
                                        v15 += 2
                                        _ = fi21
                                        _ = v22
                                    }
                                }
                                _ = v15
                            } else {
                                fmt.Println(v5, v5)
                                for fi23 := 0; fi23 < 11; fi23++ {
                                    v24 := GenS2{F0: 1, F1: 66}
                                    var v25 float64 = ((-2.25 - 1e10) / (1.5 * -2.25))
                                    var v26 uint32 = v6
                                    _ = fi23
                                    _ = v24
                                    _ = v25
                                    _ = v26
                                }
                            }
                        }
                        for fi27 := 0; fi27 < 15; fi27++ {
                            v28 := []int{fi27}
                            fmt.Println(fi7)
                            v6 = 65535
                            if (375 >= 660) {
                                if (972 <= 1) {
                                    v29 := GenS4{F0: 255}
                                    fmt.Println(v6, v6, v5)
                                    v5 = v5
                                    v5 -= 1e10
                                    _ = v29
                                }
                                _ = len(v28)
                                if ((255 >= 47) || true) {
                                    var v30 int32 = 0
                                    v31 := GenS1{F0: 1}
                                    if ((((1 != 801) || true) && false) && false) {
                                        if true {
                                            var v32 uint64 = 18446744073709551615
                                            _ = v32
                                        }
                                        var v33 float32 = 0.1
                                        var v34 int16 = 1000
                                        if len(v28) > 0 {
                                            v28[0] = -174
                                        }
                                        if ((3.14159 / -2.25) <= 1e10) {
                                            if true {
                                                switch fi27 {
                                                case 4294967295:
                                                    if (!(fi27 < fi27)) {
                                                        if true {
                                                            fmt.Println(v30)
                                                            var v35 string = "foo bar"
                                                            v36 := GenS1{F0: 0}
                                                            _ = len(v28)
                                                            v35 = v35
                                                            var v37 int16 = v34
                                                            _ = v35
                                                            _ = v36
                                                            _ = v37
                                                        } else {
                                                            if (!(!("hello" < ("foo bar" + "a")))) {
                                                                fmt.Println(fi7, fi27, v5)
                                                                v6 = (v6 / (((v6 + v6)) | 1))
                                                                var v38 int32 = v30
                                                                _ = v38
                                                            }
                                                        }
                                                    } else {
                                                        v6 = (v6 % ((v6) | 1))
                                                    }
                                                    if (((v5 + (v5 * v33)) / 1.0) != (v5 * ((v33 - float32(float64(v33))) + v33))) {
                                                        _ = len(v28)
                                                        var v39 float32 = v33
                                                        _ = v39
                                                    }
                                                    var v40 uint64 = 18446744073709551615
                                                    _ = v40
                                                case -2147483648:
                                                    switch (fi7 << ((4294967295) & 63)) {
                                                    case 128:
                                                        var v41 int16 = (v34 + v34)
                                                        fmt.Println(fi27)
                                                        _ = v41
                                                    case -604:
                                                        if len(v28) >= 2 {
                                                            v28 = v28[1:len(v28)]
                                                        }
                                                    case 2147483647:
                                                        v5 = float32(float64(v5))
                                                    default:
                                                        fmt.Println(v33, v5, fi27)
                                                    }
                                                }
                                                if len(v28) >= 2 {
                                                    v28 = v28[1:len(v28)]
                                                }
                                                if (false && (("foo bar" + "foo bar") <= "")) {
                                                    v42 := GenS2{F0: 204, F1: (v34 & v34)}
                                                    switch "\n\t" {
                                                    case "foo bar":
                                                        switch "a" {
                                                        case "":
                                                            _ = len(v28)
                                                        }
                                                    case "\n\t":
                                                        v30 -= (v30 | int32(v6))
                                                    default:
                                                        v43 := GenS5{F0: (-128 == -102), F1: 639}
                                                        _ = v43
                                                    }
                                                    var v44 bool = true
                                                    switch "foo bar" {
                                                    case "hello":
                                                        v45 := GenS4{F0: 170}
                                                        var v46 int64 = 0
                                                        v34 = (v34 & (v34 >> ((1) & 63)))
                                                        if true {
                                                            var v47 uint = 118
                                                            if v44 {
                                                                var v48 float32 = v5
                                                                var v49 int32 = (v30 % ((42) | 1))
                                                                if v44 {
                                                                    v44 = v44
                                                                } else {
                                                                    if (v6 < 290) {
                                                                        v5 += v33
                                                                        var v50 int16 = v34
                                                                        if ((!(1 > 11)) && (!((v6 < 42) && v44))) {
                                                                            var v51 int = 4294967295
                                                                            fmt.Println(v48, v34, v51)
                                                                            _ = v51
                                                                        } else {
                                                                            v30 += (v49 | (v30 & (v30 >> ((222) & 63))))
                                                                            v34 -= -32768
                                                                        }
                                                                        _ = v50
                                                                    }
                                                                    var v52 uint8 = 1
                                                                    fmt.Println(v5)
                                                                    _ = v52
                                                                }
                                                                var v53 int64 = v46
                                                                if len(v28) > 0 {
                                                                    v28[0] = -2147483648
                                                                }
                                                                _ = v48
                                                                _ = v49
                                                                _ = v53
                                                            } else {
                                                                if ((v46 * (-(v46))) < v46) {
                                                                    v34 += (v34 % ((v34) | 1))
                                                                    switch "hello" {
                                                                    case "":
                                                                        if len(v28) >= 2 {
                                                                            v28 = v28[1:len(v28)]
                                                                        }
                                                                    default:
                                                                        fmt.Println(v33)
                                                                    }
                                                                } else {
                                                                    var v54 bool = (fi7 <= (fi7 & fi7))
                                                                    var v55 uint = v47
                                                                    _ = v54
                                                                    _ = v55
                                                                }
                                                                v56 := GenS5{F0: (("" + "foo bar") >= ("" + "foo bar")), F1: -1}
                                                                if ((4294967295 == 42) || v44) {
                                                                    var v57 uint32 = (v6 | v6)
                                                                    var v58 uint32 = (v6 << (((v47 | 18446744073709551615)) & 63))
                                                                    var v59 int16 = (v34 ^ v34)
                                                                    _ = v57
                                                                    _ = v58
                                                                    _ = v59
                                                                } else {
                                                                    switch "\n\t" {
                                                                    case "":
                                                                        var v60 int32 = (v30 << (((v47 ^ (v47 << ((2) & 63)))) & 63))
                                                                        _ = v60
                                                                    default:
                                                                        fmt.Println(v30)
                                                                    }
                                                                    v28 = append(v28, 32767)
                                                                }
                                                                _ = v56
                                                            }
                                                            v30 = v30
                                                            _ = v47
                                                        }
                                                        v34 += (v34 * (v34 | (v34 & v34)))
                                                        if v44 {
                                                            v33 -= ((1.5 - v33) / v5)
                                                        }
                                                        if len(v28) > 0 {
                                                            v28[0] = (fi27 << ((18446744073709551615) & 63))
                                                        }
                                                        if (4294967295 != v6) {
                                                            v61 := GenS6{F0: "foo bar"}
                                                            _ = v61
                                                        }
                                                        v30 -= -1
                                                        _ = v45
                                                        _ = v46
                                                    case "foo bar":
                                                        var v62 float64 = 0.1
                                                        _ = v62
                                                    default:
                                                        v6 -= v6
                                                    }
                                                    var v63 int32 = v30
                                                    _ = v42
                                                    _ = v44
                                                    _ = v63
                                                }
                                            } else {
                                                var v64 float32 = (v5 + (0.1 * v33))
                                                v64 = ((v64 / v5) - v5)
                                                v6 += (^(v6))
                                                if len(v28) > 0 {
                                                    v28[0] = (-(fi27))
                                                }
                                                _ = v64
                                            }
                                            if ((1 >= v6) || (127 < 42)) {
                                                switch (("a" + "foo bar") + ("\n\t" + "\n\t")) {
                                                case "hello":
                                                    v33 = float32((1e10 * 0.0))
                                                case "foo bar":
                                                    if (true && (!false)) {
                                                        v65 := GenS3{F0: (v34 / ((v34) | 1)), F1: -1}
                                                        switch fi27 {
                                                        case -1:
                                                            var v66 uint64 = 42
                                                            _ = v66
                                                        default:
                                                            switch "" {
                                                            case "foo bar":
                                                                if len(v28) > 0 {
                                                                    v28[0] = 128
                                                                }
                                                            default:
                                                                v67 := GenS7{F0: false}
                                                                _ = v67
                                                            }
                                                        }
                                                        _ = v65
                                                    } else {
                                                        if true {
                                                            if true {
                                                                v68 := GenS8{F0: 341}
                                                                v30 -= v30
                                                                v28 = append(v28, fi7)
                                                                v30 = v30
                                                                v33 -= (1.5 / 1.5)
                                                                var v69 uint8 = 128
                                                                _ = v68
                                                                _ = v69
                                                            }
                                                            v5 -= ((1e10 / 1e10) * 3.14159)
                                                        } else {
                                                            var v70 int8 = -122
                                                            v28 = append(v28, fi7)
                                                            _ = v70
                                                        }
                                                        var v71 uint8 = 1
                                                        v30 -= (v30 << ((42) & 63))
                                                        _ = v71
                                                    }
                                                default:
                                                    if ((fi27 >> ((148) & 63)) <= fi27) {
                                                        v72 := GenS9{F0: 0, F1: ("a" + ("" + "foo bar"))}
                                                        v34 = v34
                                                        var v73 bool = true
                                                        _ = v72
                                                        _ = v73
                                                    } else {
                                                        if (1 < 0) {
                                                            switch "foo bar" {
                                                            case "\n\t":
                                                                if true {
                                                                    v74 := GenS5{F0: (fi7 < fi7), F1: v30}
                                                                    _ = len(v28)
                                                                    var v75 bool = false
                                                                    _ = v74
                                                                    _ = v75
                                                                }
                                                            default:
                                                                var v76 uint64 = 2
                                                                _ = v76
                                                            }
                                                            switch (fi7 << ((42) & 63)) {
                                                            case 260:
                                                                v34 -= v34
                                                            case 9223372036854775807:
                                                                fmt.Println(v33, v5, v30)
                                                            }
                                                            if len(v28) > 0 {
                                                                v28[0] = fi7
                                                            }
                                                        } else {
                                                            v77 := GenS8{F0: 32767}
                                                            fmt.Println(fi27)
                                                            switch int(v6) {
                                                            case -1:
                                                                v6 = (v6 - v6)
                                                            case -2:
                                                                fmt.Println(v34, v30)
                                                            }
                                                            _ = v77
                                                        }
                                                        switch fi27 {
                                                        case 32768:
                                                            var v78 int32 = (v30 / ((v30) | 1))
                                                            _ = v78
                                                        case 128:
                                                            fmt.Println(v6, v30)
                                                        }
                                                    }
                                                }
                                                var v79 int8 = 1
                                                var v80 uint32 = (v6 * v6)
                                                v30 = v30
                                                _ = v79
                                                _ = v80
                                            } else {
                                                v81 := GenS3{F0: v34, F1: 102}
                                                v34 -= v34
                                                _ = v81
                                            }
                                            v34 = (v34 / (((v34 << ((310) & 63))) | 1))
                                        }
                                        _ = v33
                                        _ = v34
                                    }
                                    var v82 uint16 = 1
                                    _ = v30
                                    _ = v31
                                    _ = v82
                                } else {
                                    v83 := GenS4{F0: 815}
                                    if ((0 != 42) || (32767 > 32767)) {
                                        var v84 bool = false
                                        if (!(65535 == 888)) {
                                            if len(v28) >= 2 {
                                                v28 = v28[1:len(v28)]
                                            }
                                        }
                                        _ = v84
                                    } else {
                                        fmt.Println(v5, fi7)
                                        var v85 float64 = 1e10
                                        _ = v85
                                    }
                                    var v86 int32 = 0
                                    _ = v83
                                    _ = v86
                                }
                                fmt.Println(v6, fi7)
                            }
                            v87 := GenS10{F0: (42 == 0), F1: 42, F2: "\n\t"}
                            _ = fi27
                            _ = v28
                            _ = v87
                        }
                    } else {
                        if (float32(3.14159) <= (v5 * v5)) {
                            v6 = (v6 >> ((4294967295) & 63))
                            if (!(!false)) {
                                switch fi7 {
                                case 0:
                                    if false {
                                        if false {
                                            v88 := []int{}
                                            _ = len(v88)
                                            v6 -= 2147483648
                                            _ = v88
                                        } else {
                                            fmt.Println(v5)
                                            if ((false || false) && (false && false)) {
                                                fmt.Println(fi7, fi7)
                                                v89 := GenS9{F0: 2, F1: ("a" + "a")}
                                                var v90 uint8 = 2
                                                _ = v89
                                                _ = v90
                                            } else {
                                                switch "a" {
                                                case "hello":
                                                    v91 := GenS11{F0: -1, F1: (!("\n\t" < ((("hello" + "foo bar") + ("" + "\n\t")) + ""))), F2: 233}
                                                    _ = v91
                                                case "a":
                                                    if (!((false || false) || (float32(3.14159) == (v5 * v5)))) {
                                                        v92 := []int{(fi7 >> ((1) & 63)), (fi7 / ((32767) | 1)), 41, fi7}
                                                        _ = v92
                                                    } else {
                                                        v6 = (^(v6))
                                                    }
                                                case "\n\t":
                                                    v6 += v6
                                                default:
                                                    v6 -= v6
                                                }
                                            }
                                        }
                                        if false {
                                            fmt.Println(v6, v5)
                                        } else {
                                            var v93 int32 = 0
                                            v5 = (v5 * (1e10 / 1e10))
                                            fmt.Println(v93)
                                            _ = v93
                                        }
                                        for fi94 := 0; fi94 < 6; fi94++ {
                                            v6 = (v6 % ((v6) | 1))
                                            if (!true) {
                                                v95 := GenS6{F0: "hello"}
                                                v6 -= 1
                                                _ = v95
                                            }
                                            v6 -= (v6 & 4294967295)
                                            _ = fi94
                                        }
                                    } else {
                                        for fi96 := 0; fi96 < 1; fi96++ {
                                            var v97 bool = true
                                            v97 = v97
                                            v98 := []int{fi7, fi96, (fi96 | fi7), (-(fi96))}
                                            _ = fi96
                                            _ = v97
                                            _ = v98
                                        }
                                        var v99 bool = false
                                        _ = v99
                                    }
                                case -2147483648:
                                    for fi100 := 0; fi100 < 16; fi100++ {
                                        v5 -= v5
                                        v5 = -2.25
                                        _ = fi100
                                    }
                                default:
                                    switch 32767 {
                                    case 0:
                                        v101 := []int{fi7}
                                        _ = v101
                                    case -2147483648:
                                        v102 := []int{-628, 42, (fi7 << ((18446744073709551615) & 63)), 986}
                                        _ = v102
                                    case 32768:
                                        for fi103 := 0; fi103 < 1; fi103++ {
                                            var v104 int = 747
                                            _ = fi103
                                            _ = v104
                                        }
                                    default:
                                        var v105 uint8 = 42
                                        _ = v105
                                    }
                                }
                                v106 := []int{fi7, fi7, (-(fi7)), fi7}
                                var v107 uint64 = 556
                                _ = v106
                                _ = v107
                            }
                        } else {
                            v5 += v5
                            var v108 float64 = (-2.25 - 1.5)
                            _ = v108
                        }
                    }
                default:
                    var v109 uint32 = (v6 / ((v6) | 1))
                    _ = v109
                }
                var v110 uint16 = 930
                _ = v110
            } else {
                for fi111 := 0; fi111 < 9; fi111++ {
                    var v112 int16 = 32767
                    v113 := []int{fi7, fi111, (fi7 >> ((42) & 63))}
                    fmt.Println(v5)
                    v5 = float32(0.0)
                    _ = fi111
                    _ = v112
                    _ = v113
                }
            }
            v5 -= 0.1
            var v114 int16 = -1
            _ = v114
        }
        _ = fi7
    }
    var v115 int32 = -1
    _ = v115
    _ = v5
    _ = v6
    return (v6 * v6)
}
func genFn1() uint32 {
    var v116 int16 = 42
    var v117 float64 = (1.5 - 0.0)
    var v118 float32 = (0.1 + 1e10)
    _ = v116
    _ = v117
    _ = v118
    return genFn0()
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[string]int)
    var v3 uint32 = 219
    var v4 uint16 = 1
    if (0 <= -1) {
        if (-716 < 0) {
            if true {
                for fi119 := 0; fi119 < 9; fi119++ {
                    if false {
                        v120 := GenS12{F0: 2, F1: 0, F2: ""}
                        _ = v120
                    }
                    fmt.Println(v3, v3, v3)
                    _ = fi119
                }
                for fi121 := 0; fi121 < 12; fi121++ {
                    var v122 uint = 4294967295
                    _ = fi121
                    _ = v122
                }
                fmt.Println(v4, v3)
            } else {
                v123 := []int{-1}
                v124 := GenS13{F0: -2147483648, F1: -128, F2: v3}
                _ = v123
                _ = v124
            }
        }
        var v125 int = -853
        switch "" {
        case "\n\t":
            v4 = (v4 * (v4 * 255))
        default:
            for fi126 := 0; fi126 < 7; fi126++ {
                v127 := GenS7{F0: (!true)}
                delete(v2, "a")
                _ = fi126
                _ = v127
            }
        }
        _ = v125
    } else {
        v128 := GenS14{F0: genFn1(), F1: (441 == 32768), F2: 2147483647}
        for fi129 := 0; fi129 < 7; fi129++ {
            var v130 bool = ((-32768 == 0) && false)
            v131 := []int{(-(fi129)), fi129, (fi129 ^ (fi129 - fi129))}
            _ = fi129
            _ = v130
            _ = v131
        }
        _ = v128
    }
    v132 := GenS5{F0: false, F1: 684}
    for fi133 := 0; fi133 < 7; fi133++ {
        v134 := []int{128}
        if (1e10 > (-2.25 * -2.25)) {
            var v135 uint = uint(v3)
            var v136 float32 = 1.5
            fmt.Println(v135, v3)
            _ = v135
            _ = v136
        } else {
            v3 += (v3 & genFn1())
            if (-1 != -71) {
                var v137 int16 = 443
                var v138 int64 = 289
                _ = v137
                _ = v138
            } else {
                var v139 bool = (!true)
                for fi140 := 0; fi140 < 4; fi140++ {
                    v141 := GenS8{F0: v4}
                    delete(v2, "\n\t")
                    fmt.Println(v4)
                    _ = fi140
                    _ = v141
                }
                v3 = (^(v3))
                _ = v139
            }
        }
        _ = fi133
        _ = v134
    }
    fmt.Println(v3, v4)
    for fi142 := 0; fi142 < 10; fi142++ {
        switch "hello" {
        case "hello":
            var v143 int = 4294967295
            _ = v143
        case "foo bar":
            v144 := []int{2147483647, fi142, fi142, 41}
            _ = v144
        default:
            var v145 int = fi142
            _ = v145
        }
        var v146 int32 = 332
        var v147 int = fi142
        _ = fi142
        _ = v146
        _ = v147
    }
    if false {
        var v148 int = -308
        _ = v148
    }
    v149 := GenS4{F0: 255}
    delete(v1, -9223372036854775808)
    var v150 uint16 = v4
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v150:", v150)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v132
    _ = v149
    _ = v150
    _ = v2
    _ = v3
    _ = v4
}
