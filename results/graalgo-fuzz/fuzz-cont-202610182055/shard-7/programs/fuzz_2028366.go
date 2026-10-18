package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool; F2 int }
type GenS2 struct { F0 uint8; F1 int32; F2 int16 }
type GenS3 struct { F0 uint8; F1 int16; F2 int }
type GenS4 struct { F0 int16; F1 string }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 int64; F1 bool; F2 string }
type GenS9 struct { F0 int32 }
type GenS10 struct { F0 uint32 }
type GenS11 struct { F0 uint64; F1 bool }
type GenS12 struct { F0 bool; F1 string }
type GenS13 struct { F0 int64 }
type GenS14 struct { F0 bool; F1 string }
type GenS15 struct { F0 int64 }
type GenS16 struct { F0 int16; F1 int8; F2 bool }
type GenS17 struct { F0 uint }
type GenS18 struct { F0 uint; F1 int64; F2 uint64 }
type GenS19 struct { F0 int64; F1 int8 }
type GenS20 struct { F0 uint8; F1 uint }
type GenS21 struct { F0 uint }
type GenS22 struct { F0 int16; F1 bool; F2 uint }
type GenS23 struct { F0 int32 }
type GenS24 struct { F0 uint64 }
type GenS25 struct { F0 bool; F1 uint32 }
func genFn0(p0 bool) int {
    v5 := GenS1{F0: "\n\t", F1: (p0 || p0), F2: -650}
    if false {
        for fi6 := 0; fi6 < 12; fi6++ {
            for fi7 := 0; fi7 < 1; fi7++ {
                if (!(-91 <= 100)) {
                    p0 = (128 >= 75)
                    if ((float32((float64(0.0) - -2.25)) > 1e10) || (true || p0)) {
                        v8 := []int{fi7, fi7}
                        p0 = (255 == 1)
                        _ = v8
                    }
                } else {
                    switch "foo bar" {
                    case "foo bar":
                        v9 := GenS1{F0: ("foo bar" + ("\n\t" + ("foo bar" + "\n\t"))), F1: (!p0), F2: (fi6 >> ((262) & 63))}
                        _ = v9
                    case "\n\t":
                        var v10 float32 = 0.1
                        _ = v10
                    }
                }
                p0 = (!(42 != 159))
                _ = fi7
            }
            if p0 {
                v11 := GenS2{F0: 128, F1: -2147483648, F2: -1}
                if (1 != 0) {
                    p0 = p0
                    p0 = false
                } else {
                    fmt.Println(fi6, fi6)
                }
                for fi12 := 0; fi12 < 9; fi12++ {
                    v13 := []int{fi12, (fi6 + (fi12 << ((468) & 63)))}
                    p0 = p0
                    _ = fi12
                    _ = v13
                }
                _ = v11
            } else {
                v14 := GenS3{F0: 91, F1: 1, F2: fi6}
                _ = v14
            }
            v15 := []int{}
            _ = fi6
            _ = v15
        }
    }
    for fi16 := 0; fi16 < 1; fi16++ {
        v17 := []int{fi16, fi16, fi16}
        _ = fi16
        _ = v17
    }
    _ = p0
    _ = v5
    return 939
}
func genFn1(p0 uint64) int64 {
    p0 -= p0
    fmt.Println(p0)
    v18 := []int{-9223372036854775808, -2147483648}
    _ = p0
    _ = v18
    return -1
}

func main() {
    var v1 uint32 = 65535
    var v2 uint = 65535
    v3 := make(map[string]uint32)
    v4 := make(map[uint32]string)
    fmt.Println(v1, v2)
    if (-9 > 1) {
        for fi19 := 0; fi19 < 18; fi19++ {
            if (!false) {
                v1 += (^(v1))
                v20 := GenS4{F0: -32768, F1: ""}
                _ = v20
            }
            v2 = 255
            _ = fi19
        }
    }
    var v21 uint64 = 248
    fmt.Println(v1, v1)
    if (42 <= 127) {
        var v22 uint64 = uint64(v1)
        var v23 int = -344
        for fi24 := 0; fi24 < 15; fi24++ {
            switch fi24 {
            case 831:
                var v25 float32 = float32(3.14159)
                _ = v25
            case 394:
                if false {
                    for fi26 := 0; fi26 < 7; fi26++ {
                        v27 := GenS3{F0: 127, F1: 42, F2: (fi24 * fi26)}
                        v21 += v22
                        if (uint8(v21) != 0) {
                            var v28 uint16 = 2
                            v29 := GenS5{F0: (v1 % (((v1 / ((v1) | 1))) | 1))}
                            _ = v28
                            _ = v29
                        } else {
                            v1 = uint32(v2)
                            v23 = -744
                            if ((v2 - (v2 | v2)) <= v2) {
                                v22 = (v22 + (v21 | 4294967295))
                                v30 := GenS6{F0: fi24}
                                v31 := GenS2{F0: 42, F1: 1, F2: -32768}
                                if false {
                                    v32 := []int{(v23 << (((v2 ^ (v2 & (v2 & v2)))) & 63)), (fi24 + (v23 & v23)), v23}
                                    if (1 >= 2147483647) {
                                        var v33 int64 = 599
                                        _ = v33
                                    }
                                    v34 := GenS5{F0: (v1 & v1)}
                                    _ = v32
                                    _ = v34
                                } else {
                                    if true {
                                        if true {
                                            var v35 float64 = 0.1
                                            _ = v35
                                        } else {
                                            var v36 bool = ((!((65535 == (^(v2))) || true)) && (-128 != -1))
                                            _ = v36
                                        }
                                        v2 -= v2
                                        fmt.Println(v2, fi26, v1)
                                        delete(v4, (v1 + 65535))
                                        v37 := GenS7{F0: 255}
                                        var v38 int8 = -128
                                        switch 360 {
                                        case 4294967295:
                                            if (uint8(v22) != 42) {
                                                if true {
                                                    v39 := GenS5{F0: 1}
                                                    v40 := GenS8{F0: 0, F1: ((^(v1)) != (v1 >> ((uint(v1)) & 63))), F2: ("foo bar" + "hello")}
                                                    _ = v39
                                                    _ = v40
                                                } else {
                                                    var v41 uint32 = 4294967295
                                                    var v42 string = "hello"
                                                    var v43 uint32 = (v41 + (v1 + v1))
                                                    v23 = (fi26 / (((v23 & (^(fi26)))) | 1))
                                                    _ = v41
                                                    _ = v42
                                                    _ = v43
                                                }
                                                switch (fi24 % ((23) | 1)) {
                                                case 867:
                                                    v38 += -75
                                                case 673:
                                                    v21 = v21
                                                case 936:
                                                    var v44 int = 42
                                                    _ = v44
                                                default:
                                                    if false {
                                                        v45 := []int{genFn0(true), (fi26 / ((v23) | 1)), (v23 | v23)}
                                                        switch "foo bar" {
                                                        case "":
                                                            if ((^(v21)) < 0) {
                                                                var v46 uint64 = (v21 + (v21 | 510))
                                                                _ = v46
                                                            } else {
                                                                if true {
                                                                    if true {
                                                                        if (!((v21 - (v22 / (((v22 | v21)) | 1))) != 227)) {
                                                                            v2 += v2
                                                                        }
                                                                    } else {
                                                                        if false {
                                                                            var v47 float32 = (3.14159 * (1e10 / (0.1 - (1e10 * 1e10))))
                                                                            v21 -= (v22 + v22)
                                                                            var v48 uint16 = 1
                                                                            var v49 uint32 = uint32(v2)
                                                                            _ = v47
                                                                            _ = v48
                                                                            _ = v49
                                                                        }
                                                                    }
                                                                } else {
                                                                    v22 = (v21 >> (((v2 - (^(v2)))) & 63))
                                                                }
                                                            }
                                                        }
                                                        v45 = append(v45, (^(fi24)))
                                                        if true {
                                                            if ((("hello" + ("a" + "")) + ((((("\n\t" + "foo bar") + "") + "") + "foo bar") + "hello")) == "\n\t") {
                                                                if false {
                                                                    var v50 float64 = 0.0
                                                                    v45 = append(v45, (fi24 ^ int(v2)))
                                                                    _ = v50
                                                                }
                                                                switch "" {
                                                                case "a":
                                                                    v22 -= 2
                                                                default:
                                                                    var v51 uint8 = 42
                                                                    v52 := GenS9{F0: -941}
                                                                    _ = v51
                                                                    _ = v52
                                                                }
                                                                var v53 int32 = 2147483647
                                                                _ = v53
                                                            } else {
                                                                if false {
                                                                    v45 = append(v45, fi24)
                                                                    if false {
                                                                        if len(v45) >= 2 {
                                                                            v45 = v45[1:len(v45)]
                                                                        }
                                                                        v2 -= v2
                                                                        v54 := GenS10{F0: v1}
                                                                        _ = v54
                                                                    } else {
                                                                        if (!false) {
                                                                            v55 := GenS3{F0: 42, F1: 42, F2: 128}
                                                                            v22 = v22
                                                                            _ = v55
                                                                        }
                                                                        fmt.Println(v38, v2)
                                                                    }
                                                                    v2 += (^(v2))
                                                                }
                                                                _ = len(v45)
                                                                v2 = 65535
                                                            }
                                                            var v56 int32 = -879
                                                            if false {
                                                                if false {
                                                                    v22 -= v21
                                                                    if true {
                                                                        if (false || ((!(v22 > 4294967295)) && (((v56 | (v56 - (v56 & 0))) != (v56 & v56)) || true))) {
                                                                            var v57 float32 = (-2.25 * float32(((float64(0.1) + 1e10) + 1.5)))
                                                                            switch genFn0((v21 >= (v21 | (v21 * (v22 * v21))))) {
                                                                            case 4294967295:
                                                                                v3["foo bar"] = v1
                                                                            case -2147483648:
                                                                                var v58 int8 = v38
                                                                                _ = v58
                                                                            case 0:
                                                                                v23 -= (fi24 << ((205) & 63))
                                                                            default:
                                                                                var v59 uint = (v2 / (((v2 ^ 1)) | 1))
                                                                                _ = v59
                                                                            }
                                                                            _ = v57
                                                                        } else {
                                                                            var v60 string = ""
                                                                            v60 = v60
                                                                            _ = v60
                                                                        }
                                                                    }
                                                                    v23 = fi26
                                                                }
                                                                if ((v56 & v56) > 1) {
                                                                    var v61 uint8 = 255
                                                                    _ = v61
                                                                }
                                                                v62 := GenS11{F0: (v22 * 992), F1: ("a" != (("hello" + "foo bar") + "\n\t"))}
                                                                _ = v62
                                                            } else {
                                                                if (!true) {
                                                                    var v63 int8 = 127
                                                                    fmt.Println(v1, v22)
                                                                    if (!true) {
                                                                        v64 := GenS12{F0: (-679 < 1), F1: "foo bar"}
                                                                        v63 -= v38
                                                                        v65 := GenS13{F0: 9223372036854775807}
                                                                        _ = v64
                                                                        _ = v65
                                                                    } else {
                                                                        var v66 uint32 = (v1 ^ (v1 / ((763) | 1)))
                                                                        _ = v66
                                                                    }
                                                                    _ = v63
                                                                } else {
                                                                    v67 := GenS1{F0: "hello", F1: (true || (1000 < 0)), F2: (fi26 << ((v2) & 63))}
                                                                    var v68 uint8 = 2
                                                                    fmt.Println(v21, v38, fi24)
                                                                    _ = v67
                                                                    _ = v68
                                                                }
                                                            }
                                                            _ = v56
                                                        } else {
                                                            if false {
                                                                v69 := GenS14{F0: true, F1: "\n\t"}
                                                                if false {
                                                                    var v70 int16 = -32768
                                                                    var v71 uint8 = 128
                                                                    v1 -= v1
                                                                    _ = v70
                                                                    _ = v71
                                                                } else {
                                                                    v72 := GenS15{F0: 0}
                                                                    _ = v72
                                                                }
                                                                _ = v69
                                                            } else {
                                                                var v73 uint64 = (^(v22))
                                                                _ = v73
                                                            }
                                                            _ = len(v45)
                                                        }
                                                        var v74 uint8 = 0
                                                        if ("" <= "foo bar") {
                                                            if false {
                                                                delete(v4, v1)
                                                            } else {
                                                                v75 := GenS16{F0: 1, F1: (v38 << ((v2) & 63)), F2: true}
                                                                switch (fi24 ^ genFn0(true)) {
                                                                case 1:
                                                                    v76 := v3["a"]
                                                                    _ = v76
                                                                default:
                                                                    if (-1 == 42) {
                                                                        fmt.Println(v22, v38)
                                                                        v38 -= (v38 + v38)
                                                                        var v77 uint64 = (v22 & v22)
                                                                        v78 := GenS7{F0: 108}
                                                                        _ = v77
                                                                        _ = v78
                                                                    }
                                                                }
                                                                v74 -= (v74 >> ((v2) & 63))
                                                                _ = v75
                                                            }
                                                            v79 := v4[2147483648]
                                                            _ = v79
                                                        } else {
                                                            v1 = v1
                                                            var v80 uint32 = v1
                                                            if false {
                                                                var v81 uint16 = 2
                                                                var v82 string = ("" + "a")
                                                                switch 128 {
                                                                case 32768:
                                                                    v83 := GenS9{F0: 0}
                                                                    _ = v83
                                                                case 0:
                                                                    var v84 float32 = 3.14159
                                                                    _ = v84
                                                                }
                                                                _ = v81
                                                                _ = v82
                                                            }
                                                            _ = v80
                                                        }
                                                        v85 := GenS9{F0: 635}
                                                        _ = v45
                                                        _ = v74
                                                        _ = v85
                                                    } else {
                                                        var v86 uint32 = (v1 - v1)
                                                        _ = len(v4)
                                                        v87 := []int{fi24, fi24, genFn0(true)}
                                                        if true {
                                                            v88 := GenS17{F0: v2}
                                                            _ = v88
                                                        }
                                                        v89 := v3["\n\t"]
                                                        _ = v86
                                                        _ = v87
                                                        _ = v89
                                                    }
                                                }
                                            } else {
                                                v38 = 127
                                                var v90 uint8 = 247
                                                v90 -= v90
                                                _ = v90
                                            }
                                        default:
                                            if (!true) {
                                                v2 = v2
                                            }
                                            var v91 uint8 = 42
                                            _ = v91
                                        }
                                        _ = v37
                                        _ = v38
                                    }
                                    switch v23 {
                                    case 2147483647:
                                        var v92 string = (("a" + ("" + "a")) + (("hello" + "\n\t") + "\n\t"))
                                        _ = v92
                                    }
                                }
                                v93 := []int{}
                                v1 -= (v1 % ((v1) | 1))
                                fmt.Println(v23, v1, v1)
                                _ = v30
                                _ = v31
                                _ = v93
                            }
                            var v94 float64 = 1e10
                            _ = v94
                        }
                        _ = fi26
                        _ = v27
                    }
                    var v95 uint16 = uint16(v1)
                    _ = v95
                } else {
                    switch (v23 / ((fi24) | 1)) {
                    case 0:
                        if true {
                            v2 = (v2 / ((v2) | 1))
                            if (59 == 127) {
                                switch v23 {
                                case 32767:
                                    v96 := GenS18{F0: v2, F1: genFn1(133), F2: v22}
                                    _ = v96
                                default:
                                    v97 := GenS15{F0: -499}
                                    _ = v97
                                }
                                switch genFn0(((false && ("\n\t" == "hello")) || true)) {
                                case -736:
                                    switch (^(v23)) {
                                    case 543:
                                        v2 += v2
                                    case -2147483648:
                                        for fi98 := 0; fi98 < 19; fi98++ {
                                            v99 := GenS19{F0: genFn1((v21 << ((v2) & 63))), F1: int8(v2)}
                                            _ = fi98
                                            _ = v99
                                        }
                                    default:
                                        fmt.Println(v2, fi24, v23)
                                    }
                                case -552:
                                    v23 -= int(v1)
                                case 530:
                                    v1 += (v1 + v1)
                                default:
                                    v100 := []int{(v23 >> (((v2 / ((uint(v21)) | 1))) & 63)), v23}
                                    _ = v100
                                }
                            }
                        } else {
                            switch ("\n\t" + ("hello" + "foo bar")) {
                            case "":
                                v21 = v21
                            }
                        }
                    case -1:
                        if false {
                            var v101 uint8 = 0
                            _ = len(v4)
                            v21 = 2
                            _ = v101
                        } else {
                            switch (("\n\t" + (("\n\t" + "foo bar") + "")) + "foo bar") {
                            case "\n\t":
                                var v102 uint8 = 58
                                _ = v102
                            case "foo bar":
                                if true {
                                    switch "a" {
                                    case "\n\t":
                                        if (!((!(false || false)) && false)) {
                                            var v103 int32 = -2147483648
                                            var v104 uint16 = 2
                                            v105 := GenS17{F0: uint(v22)}
                                            _ = v103
                                            _ = v104
                                            _ = v105
                                        }
                                    case "":
                                        v106 := GenS20{F0: 255, F1: (^(v2))}
                                        _ = v106
                                    case "hello":
                                        switch (fi24 + (-(fi24))) {
                                        case -2147483648:
                                            var v107 int64 = 0
                                            _ = v107
                                        case 9223372036854775807:
                                            switch "\n\t" {
                                            case "":
                                                v2 += v2
                                            }
                                        }
                                    default:
                                        var v108 uint8 = 2
                                        _ = v108
                                    }
                                    var v109 int32 = 0
                                    v1 += (v1 + (v1 + 96))
                                    _ = v109
                                } else {
                                    if true {
                                        v2 = 385
                                        v22 = (v22 | (^(v22)))
                                    } else {
                                        if (!true) {
                                            if true {
                                                var v110 string = ""
                                                fmt.Println(v2, v22, fi24)
                                                _ = v110
                                            }
                                        }
                                        var v111 int = fi24
                                        if ("foo bar" != "a") {
                                            for fi112 := 0; fi112 < 14; fi112++ {
                                                var v113 float64 = -2.25
                                                if (false && true) {
                                                    if false {
                                                        v114 := GenS21{F0: 255}
                                                        v21 -= (v21 ^ v22)
                                                        var v115 uint16 = 2
                                                        if (-2147483648 <= 795) {
                                                            if true {
                                                                v4[(v1 * uint32(v115))] = "hello"
                                                                v2 = uint(v22)
                                                                v115 = v115
                                                            }
                                                            v116 := GenS22{F0: 0, F1: false, F2: v2}
                                                            v117 := []int{v23, genFn0((true && (!true))), (fi24 | v111)}
                                                            _ = v116
                                                            _ = v117
                                                        }
                                                        if false {
                                                            v111 += genFn0((v113 != ((v113 - v113) * v113)))
                                                            var v118 int16 = -32768
                                                            if (fi112 >= (fi24 + fi112)) {
                                                                var v119 int64 = -1
                                                                v120 := GenS23{F0: int32(v119)}
                                                                var v121 uint32 = (v1 << ((v2) & 63))
                                                                _ = v119
                                                                _ = v120
                                                                _ = v121
                                                            } else {
                                                                fmt.Println(v1, v23)
                                                            }
                                                            _ = v118
                                                        } else {
                                                            var v122 int32 = -1
                                                            _ = v122
                                                        }
                                                        _ = v114
                                                        _ = v115
                                                    } else {
                                                        v111 += (fi24 + genFn0(true))
                                                        var v123 uint16 = 79
                                                        _ = v123
                                                    }
                                                    v111 += 368
                                                    var v124 uint16 = 1
                                                    _ = v124
                                                }
                                                var v125 uint8 = 14
                                                _ = fi112
                                                _ = v113
                                                _ = v125
                                            }
                                            var v126 uint = v2
                                            v127 := GenS24{F0: (v22 - v21)}
                                            _ = v126
                                            _ = v127
                                        } else {
                                            v128 := []int{}
                                            if false {
                                                if (140 != 65535) {
                                                    if len(v128) > 0 {
                                                        v128[0] = -441
                                                    }
                                                    for fi129 := 0; fi129 < 7; fi129++ {
                                                        var v130 int64 = 0
                                                        _ = fi129
                                                        _ = v130
                                                    }
                                                    var v131 int32 = -851
                                                    _ = v131
                                                }
                                                v132 := GenS10{F0: v1}
                                                v111 = (fi24 >> ((0) & 63))
                                                _ = v132
                                            }
                                            _ = v128
                                        }
                                        _ = v111
                                    }
                                    switch "hello" {
                                    case "a":
                                        if (0 < 127) {
                                            v21 -= v21
                                            v21 += 0
                                        }
                                    case "hello":
                                        for fi133 := 0; fi133 < 17; fi133++ {
                                            if (((float32((0.0 + 3.14159)) / (1e10 / float32(0.1))) >= float32(((-2.25 * ((1e10 * 0.1) + 1e10)) - 1e10))) || false) {
                                                var v134 uint16 = 0
                                                _ = v134
                                            }
                                            _ = fi133
                                        }
                                    default:
                                        if false {
                                            v135 := []int{-280, fi24}
                                            if (178 != 299) {
                                                var v136 float64 = (((-2.25 / 1.0) * 1e10) - 0.0)
                                                v137 := GenS10{F0: 2147483648}
                                                _ = v136
                                                _ = v137
                                            } else {
                                                fmt.Println(v23, v1)
                                                for fi138 := 0; fi138 < 11; fi138++ {
                                                    var v139 int8 = 0
                                                    _ = fi138
                                                    _ = v139
                                                }
                                            }
                                            _ = v135
                                        } else {
                                            if (!("a" <= ((("foo bar" + (("\n\t" + ("hello" + "foo bar")) + ("\n\t" + ("\n\t" + ((("" + "") + ("" + (("" + "\n\t") + ("" + (("foo bar" + ("\n\t" + ("foo bar" + "a"))) + "a"))))) + "a"))))) + "\n\t") + ("hello" + ("hello" + (("hello" + (("\n\t" + "\n\t") + "a")) + ("hello" + ("a" + "foo bar")))))))) {
                                                v140 := []int{genFn0((!true)), fi24, (fi24 + (fi24 ^ (v23 | 0))), -881}
                                                _ = v140
                                            }
                                        }
                                    }
                                    var v141 int32 = -2147483648
                                    _ = v141
                                }
                            case "a":
                                v142 := []int{}
                                _ = v142
                            }
                            switch (("hello" + (("foo bar" + ("foo bar" + ("hello" + ("" + ("hello" + "foo bar"))))) + ("a" + "hello"))) + (("\n\t" + "a") + "hello")) {
                            case "hello":
                                var v143 uint = v2
                                _ = v143
                            }
                            var v144 int8 = 127
                            _ = v144
                        }
                    case 9223372036854775807:
                        var v145 uint = (v2 | v2)
                        _ = v145
                    default:
                        fmt.Println(v21, v2)
                    }
                    v146 := []int{fi24, genFn0(true), (v23 / ((32768) | 1))}
                    fmt.Println(v1)
                    _ = v146
                }
            case 127:
                var v147 uint64 = v22
                _ = v147
            default:
                v148 := []int{v23, v23, 127}
                _ = v148
            }
            if (4294967295 >= v1) {
                var v149 int = v23
                _ = v149
            } else {
                v22 = v21
                for fi150 := 0; fi150 < 6; fi150++ {
                    var v151 uint32 = 2147483647
                    fmt.Println(fi150)
                    _ = fi150
                    _ = v151
                }
            }
            var v152 uint8 = 0
            _ = fi24
            _ = v152
        }
        _ = v22
        _ = v23
    }
    var v153 uint64 = v21
    var v154 uint8 = 27
    if (!false) {
        var v155 uint64 = 0
        _ = v155
    } else {
        fmt.Println(v2, v153, v2)
        v156 := GenS8{F0: int64(v1), F1: true, F2: "a"}
        _ = v156
    }
    for fi157 := 0; fi157 < 8; fi157++ {
        v158 := []int{}
        v159 := GenS8{F0: -916, F1: (false && false), F2: "\n\t"}
        v154 -= v154
        _ = fi157
        _ = v158
        _ = v159
    }
    var v160 int8 = 42
    v4[962] = "foo bar"
    v2 -= v2
    var v161 string = ("a" + "foo bar")
    v1 = v1
    for fi162 := 0; fi162 < 17; fi162++ {
        for fi163 := 0; fi163 < 18; fi163++ {
            var v164 bool = true
            v165 := []int{genFn0(v164), genFn0(v164), fi162, (fi162 >> ((v2) & 63))}
            _ = fi163
            _ = v164
            _ = v165
        }
        v166 := GenS25{F0: ((!true) || true), F1: v1}
        _ = fi162
        _ = v166
    }
    v153 = 4294967295
    _ = len(v3)
    var v167 int = -378
    fmt.Println("v1:", v1)
    fmt.Println("v153:", v153)
    fmt.Println("v154:", v154)
    fmt.Println("v160:", v160)
    fmt.Println("v161:", v161)
    fmt.Println("v167:", v167)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v153
    _ = v154
    _ = v160
    _ = v161
    _ = v167
    _ = v2
    _ = v21
    _ = v3
    _ = v4
}
