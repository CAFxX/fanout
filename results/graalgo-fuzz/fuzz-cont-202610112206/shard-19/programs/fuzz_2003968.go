package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint8; F1 uint8; F2 uint64 }
type GenS3 struct { F0 uint; F1 int64 }
type GenS4 struct { F0 uint16; F1 uint; F2 uint32 }
type GenS5 struct { F0 bool; F1 int8 }
type GenS6 struct { F0 uint32; F1 string; F2 int32 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 bool; F1 string }
type GenS10 struct { F0 uint64; F1 bool }
type GenS11 struct { F0 int32; F1 bool }
type GenS12 struct { F0 int; F1 uint64; F2 int32 }
type GenS13 struct { F0 int8; F1 uint8; F2 bool }
func genFn0(p0 uint16, p1 uint32) uint64 {
    var v6 int = 277
    v7 := GenS1{F0: 100}
    v8 := []int{4294967295}
    v6 = (v6 % (((v6 * v6)) | 1))
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return 131
}
func genFn1() int {
    var v9 int64 = 585
    if (true || (!(float32((((-2.25 + float64(0.1)) / (1.5 * 0.1)) * 0.0)) != 0.1))) {
        switch "" {
        case "hello":
            fmt.Println(v9)
        default:
            switch 1 {
            case 2147483647:
                fmt.Println(v9)
            }
        }
        fmt.Println(v9)
    }
    _ = v9
    return 373
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int8 = -1
    var v3 float64 = 3.14159
    var v4 uint32 = 2147483648
    var v5 uint16 = 1
    switch 9223372036854775807 {
    case 1:
        if false {
            fmt.Println(v3, v3)
            var v10 uint = 450
            _ = v10
        } else {
            switch "foo bar" {
            case "foo bar":
                v11 := GenS1{F0: 42}
                _ = v11
            case "":
                v1[v4] = ""
            case "hello":
                var v12 string = "a"
                _ = v12
            }
            for fi13 := 0; fi13 < 19; fi13++ {
                var v14 int64 = -9223372036854775808
                fmt.Println(v4, v14)
                _ = fi13
                _ = v14
            }
            switch int(v2) {
            case 128:
                for fi15 := 0; fi15 < 2; fi15++ {
                    v1[v4] = ("\n\t" + ("foo bar" + "\n\t"))
                    var v16 int = fi15
                    _ = fi15
                    _ = v16
                }
            case -1:
                switch int(v2) {
                case 4294967295:
                    if (!false) {
                        v17 := GenS1{F0: v2}
                        var v18 int16 = 1
                        v19 := []int{617}
                        _ = v17
                        _ = v18
                        _ = v19
                    } else {
                        var v20 uint8 = 218
                        _ = v20
                    }
                case 2147483647:
                    v4 -= 65535
                }
            case 127:
                v21 := []int{4294967295}
                _ = v21
            }
        }
    case 32768:
        var v22 uint = 18446744073709551615
        _ = v22
    }
    if (!false) {
        if false {
            v23 := []int{0, 506}
            v4 += v4
            _ = v23
        } else {
            for fi24 := 0; fi24 < 12; fi24++ {
                for fi25 := 0; fi25 < 9; fi25++ {
                    fmt.Println(fi24)
                    _ = fi25
                }
                var v26 int16 = 0
                _ = fi24
                _ = v26
            }
            var v27 uint16 = v5
            v28 := GenS2{F0: 255, F1: 231, F2: 4294967295}
            _ = v27
            _ = v28
        }
        if (670 <= 1) {
            v1[v4] = "foo bar"
        } else {
            for fi29 := 0; fi29 < 19; fi29++ {
                var v30 int8 = v2
                _ = fi29
                _ = v30
            }
            v31 := GenS3{F0: 436, F1: -1}
            _ = v31
        }
        for fi32 := 0; fi32 < 12; fi32++ {
            if false {
                switch "\n\t" {
                case "":
                    if (0 <= 1) {
                        if false {
                            v33 := GenS3{F0: 42, F1: -1}
                            v2 = v2
                            _ = v33
                        }
                        if (!(!(genFn1() < fi32))) {
                            v34 := GenS3{F0: 65535, F1: -815}
                            for fi35 := 0; fi35 < 12; fi35++ {
                                v5 = uint16(v4)
                                fmt.Println(v2)
                                _ = fi35
                            }
                            _ = v34
                        }
                    } else {
                        if (!false) {
                            v36 := GenS1{F0: v2}
                            fmt.Println(v2, v3)
                            _ = v36
                        } else {
                            for fi37 := 0; fi37 < 5; fi37++ {
                                v5 -= v5
                                var v38 uint64 = 42
                                if ((v38 << ((255) & 63)) < 18446744073709551615) {
                                    fmt.Println(v38, v3, v3)
                                    fmt.Println(fi37)
                                    v39 := []int{genFn1()}
                                    _ = v39
                                }
                                _ = fi37
                                _ = v38
                            }
                            for fi40 := 0; fi40 < 2; fi40++ {
                                var v41 float32 = -2.25
                                if ((725 > (fi32 << ((919) & 63))) || false) {
                                    if ((!(genFn0((v5 + v5), (^(v4))) < 431)) && (0 > -1)) {
                                        if true {
                                            if false {
                                                fmt.Println(fi40)
                                            }
                                        } else {
                                            var v42 int8 = (v2 >> ((255) & 63))
                                            var v43 uint16 = (v5 | v5)
                                            v44 := GenS3{F0: 2, F1: 23}
                                            _ = v42
                                            _ = v43
                                            _ = v44
                                        }
                                        if true {
                                            var v45 int32 = -1
                                            if false {
                                                fmt.Println(v41)
                                            } else {
                                                _ = len(v1)
                                                _ = len(v1)
                                                v46 := []int{fi32, (fi32 & 4294967295), (^(fi40)), fi40}
                                                _ = v46
                                            }
                                            _ = v45
                                        } else {
                                            v4 = (v4 & v4)
                                            if (int16(v5) < 0) {
                                                if true {
                                                    var v47 uint64 = 1
                                                    fmt.Println(v41, fi40, v5)
                                                    if (!(v5 >= 2)) {
                                                        v48 := GenS3{F0: 65535, F1: int64(v47)}
                                                        var v49 int16 = 32767
                                                        v50 := GenS1{F0: (v2 | v2)}
                                                        _ = v48
                                                        _ = v49
                                                        _ = v50
                                                    }
                                                    v4 = v4
                                                    v51 := GenS1{F0: (-(v2))}
                                                    _ = v47
                                                    _ = v51
                                                }
                                                v41 = (v41 * float32(((v3 * v3) + v3)))
                                                var v52 bool = true
                                                _ = v52
                                            } else {
                                                v3 -= v3
                                                if false {
                                                    v53 := GenS1{F0: 0}
                                                    _ = v53
                                                }
                                            }
                                        }
                                    } else {
                                        v4 -= 0
                                        v54 := GenS4{F0: (v5 * (^(v5))), F1: uint(v4), F2: (v4 & 42)}
                                        if true {
                                            var v55 uint8 = 0
                                            _ = v55
                                        } else {
                                            v56 := []int{fi40, 127}
                                            if true {
                                                v2 = v2
                                            }
                                            _ = v56
                                        }
                                        _ = v54
                                    }
                                } else {
                                    if false {
                                        v2 = -37
                                        v5 = (v5 & (v5 >> ((1) & 63)))
                                        if true {
                                            v41 -= float32(0.1)
                                        }
                                    }
                                    v3 = v3
                                }
                                v57 := []int{}
                                _ = fi40
                                _ = v41
                                _ = v57
                            }
                        }
                        v3 = v3
                    }
                }
                _ = len(v1)
            } else {
                v4 = 2
            }
            for fi58 := 0; fi58 < 10; fi58++ {
                if (-687 < -693) {
                    fmt.Println(v2, v2, v3)
                    var v59 int = (fi32 - fi32)
                    _ = v59
                } else {
                    v2 = v2
                    v60 := []int{}
                    _ = v60
                }
                _ = fi58
            }
            for fi61 := 0; fi61 < 20; fi61++ {
                fmt.Println(v2)
                if (1 <= 0) {
                    _ = len(v1)
                } else {
                    fmt.Println(v2, v3, v2)
                    v3 = float64(0.0)
                }
                var v62 int16 = 1000
                v63 := []int{(fi32 - (fi32 & (fi32 >> ((615) & 63))))}
                fmt.Println(v2, v2, v62)
                _ = fi61
                _ = v62
                _ = v63
            }
            _ = fi32
        }
    } else {
        if (42 < 42) {
            for fi64 := 0; fi64 < 12; fi64++ {
                v1[(v4 | 1)] = (("" + "a") + "a")
                _ = fi64
            }
            v65 := []int{genFn1()}
            for fi66 := 0; fi66 < 4; fi66++ {
                if (true && (330 <= 2)) {
                    switch "a" {
                    case "\n\t":
                        if ((fi66 | -336) > fi66) {
                            if true {
                                v67 := GenS5{F0: (0 >= -32768), F1: v2}
                                var v68 uint64 = 1
                                fmt.Println(v4, v68, v4)
                                _ = v67
                                _ = v68
                            } else {
                                v2 -= v2
                                for fi69 := 0; fi69 < 16; fi69++ {
                                    _ = len(v65)
                                    switch ("\n\t" + ("a" + (((("foo bar" + "hello") + "hello") + "foo bar") + ("" + "a")))) {
                                    case "\n\t":
                                        switch genFn1() {
                                        case -2147483648:
                                            if true {
                                                v3 += 0.1
                                                var v70 int = 128
                                                _ = v70
                                            } else {
                                                switch fi66 {
                                                case 0:
                                                    if (-9223372036854775808 == -9223372036854775808) {
                                                        v3 += (v3 + 1e10)
                                                    } else {
                                                        var v71 int8 = (v2 % (((v2 / ((-128) | 1))) | 1))
                                                        var v72 int32 = -1
                                                        v73 := GenS6{F0: v4, F1: "a", F2: int32(v4)}
                                                        _ = v71
                                                        _ = v72
                                                        _ = v73
                                                    }
                                                default:
                                                    var v74 uint8 = 182
                                                    _ = v74
                                                }
                                                _ = len(v1)
                                            }
                                        case 658:
                                            v2 -= -128
                                        case 2147483647:
                                            var v75 uint32 = v4
                                            _ = v75
                                        default:
                                            if false {
                                                delete(v1, v4)
                                                fmt.Println(v5)
                                                if (!((0.0 - -2.25) == (float64(float32(v3)) / 3.14159))) {
                                                    switch -870 {
                                                    case 2147483647:
                                                        switch "foo bar" {
                                                        case "":
                                                            delete(v1, (v4 << ((2) & 63)))
                                                        case "a":
                                                            v4 -= 42
                                                        case "hello":
                                                            if true {
                                                                v65 = append(v65, fi66)
                                                                fmt.Println(v3)
                                                            } else {
                                                                var v76 uint8 = 1
                                                                _ = v76
                                                            }
                                                        default:
                                                            var v77 int16 = -138
                                                            _ = v77
                                                        }
                                                    }
                                                    switch (("hello" + ("hello" + ("" + "foo bar"))) + "") {
                                                    case "foo bar":
                                                        switch "\n\t" {
                                                        case "":
                                                            delete(v1, (v4 / ((1) | 1)))
                                                        case "a":
                                                            if (!true) {
                                                                if (166 < 1) {
                                                                    v78 := GenS1{F0: (v2 >> ((977) & 63))}
                                                                    var v79 uint = 4294967295
                                                                    fmt.Println(v3)
                                                                    if (!(!("a" != ((("foo bar" + "\n\t") + "a") + ("hello" + ("\n\t" + "a")))))) {
                                                                        v79 += 688
                                                                        var v80 uint32 = (^(v4))
                                                                        var v81 float32 = 1e10
                                                                        _ = v80
                                                                        _ = v81
                                                                    } else {
                                                                        v82 := GenS5{F0: true, F1: (v2 % ((v2) | 1))}
                                                                        if (((fi66 >> ((35) & 63)) > genFn1()) && false) {
                                                                            if len(v65) >= 2 {
                                                                                v65 = v65[1:len(v65)]
                                                                            }
                                                                            v4 -= 376
                                                                            _ = len(v1)
                                                                        } else {
                                                                            var v83 int = (fi69 % ((int(v79)) | 1))
                                                                            _ = v83
                                                                        }
                                                                        fmt.Println(v79, v2, v4)
                                                                        if false {
                                                                            v3 -= v3
                                                                            var v84 bool = (1 != int32(v79))
                                                                            _ = v84
                                                                        } else {
                                                                            var v85 int64 = -737
                                                                            _ = v85
                                                                        }
                                                                        _ = v82
                                                                    }
                                                                    _ = v78
                                                                    _ = v79
                                                                }
                                                                v3 -= 0.1
                                                                v3 = 0.1
                                                            } else {
                                                                if len(v65) >= 2 {
                                                                    v65 = v65[1:len(v65)]
                                                                }
                                                            }
                                                        case "hello":
                                                            v86 := GenS2{F0: 110, F1: 172, F2: 18446744073709551615}
                                                            v5 = v5
                                                            if (("a" + "hello") <= (("hello" + "a") + ("hello" + ("a" + "")))) {
                                                                var v87 int64 = 843
                                                                _ = v87
                                                            }
                                                            _ = v86
                                                        default:
                                                            v5 -= (v5 / ((2) | 1))
                                                        }
                                                    case "hello":
                                                        if (!true) {
                                                            if false {
                                                                fmt.Println(v3, fi66)
                                                                fmt.Println(v4, v3)
                                                                if (0 == (v4 + 2147483647)) {
                                                                    var v88 int = genFn1()
                                                                    if (true && false) {
                                                                        v5 = (^(v5))
                                                                    }
                                                                    var v89 float64 = v3
                                                                    _ = v88
                                                                    _ = v89
                                                                }
                                                            }
                                                            var v90 uint64 = uint64(v5)
                                                            if true {
                                                                v65 = append(v65, genFn1())
                                                                v91 := GenS2{F0: 255, F1: 247, F2: genFn0(v5, v4)}
                                                                _ = v91
                                                            } else {
                                                                v5 = v5
                                                            }
                                                            v90 += (v90 | (v90 << ((18446744073709551615) & 63)))
                                                            _ = v90
                                                        } else {
                                                            if (127 < 58) {
                                                                v92 := GenS7{F0: (1 < 0)}
                                                                _ = v92
                                                            } else {
                                                                _ = len(v65)
                                                            }
                                                            if false {
                                                                v3 = float64(3.14159)
                                                                v5 -= v5
                                                            }
                                                            v93 := GenS8{F0: -63}
                                                            var v94 uint32 = v4
                                                            v95 := GenS9{F0: (("" + ("a" + "")) >= "hello"), F1: "foo bar"}
                                                            _ = v93
                                                            _ = v94
                                                            _ = v95
                                                        }
                                                    case "a":
                                                        switch ("foo bar" + "foo bar") {
                                                        case "\n\t":
                                                            var v96 uint = 1
                                                            _ = v96
                                                        }
                                                        v97 := GenS9{F0: (-9223372036854775808 <= -894), F1: "a"}
                                                        var v98 uint = 1
                                                        v3 = v3
                                                        _ = v97
                                                        _ = v98
                                                    }
                                                } else {
                                                    v99 := GenS9{F0: (!true), F1: "hello"}
                                                    _ = len(v1)
                                                    v2 += v2
                                                    switch 32768 {
                                                    case 32767:
                                                        v4 = (v4 + v4)
                                                    default:
                                                        switch (fi66 >> ((255) & 63)) {
                                                        case -9223372036854775808:
                                                            var v100 float32 = float32(((v3 * (1.5 - -2.25)) + v3))
                                                            _ = v100
                                                        case -848:
                                                            if len(v65) >= 2 {
                                                                v65 = v65[1:len(v65)]
                                                            }
                                                        default:
                                                            if (42 >= v5) {
                                                                v2 -= (v2 >> ((4294967295) & 63))
                                                                v5 = (^(v5))
                                                                if true {
                                                                    v1[486] = ((("\n\t" + ("a" + "")) + "hello") + "foo bar")
                                                                    v101 := GenS9{F0: true, F1: (("foo bar" + ((("foo bar" + ("hello" + (("\n\t" + "a") + "foo bar"))) + "foo bar") + (("hello" + "a") + ""))) + (("foo bar" + (("" + "foo bar") + "\n\t")) + "foo bar"))}
                                                                    var v102 uint32 = 42
                                                                    v103 := GenS4{F0: (v5 | (v5 - 65535)), F1: 37, F2: (^(v4))}
                                                                    _ = v101
                                                                    _ = v102
                                                                    _ = v103
                                                                } else {
                                                                    var v104 uint32 = (^(v4))
                                                                    _ = v104
                                                                }
                                                            }
                                                        }
                                                    }
                                                    _ = v99
                                                }
                                            } else {
                                                v105 := GenS1{F0: 42}
                                                var v106 int = (fi66 % ((genFn1()) | 1))
                                                var v107 float32 = 1.5
                                                var v108 int8 = (v2 & (v2 | v2))
                                                v109 := GenS4{F0: v5, F1: 2, F2: (v4 & 559)}
                                                _ = v105
                                                _ = v106
                                                _ = v107
                                                _ = v108
                                                _ = v109
                                            }
                                        }
                                    }
                                    _ = fi69
                                }
                            }
                            for fi110 := 0; fi110 < 19; fi110++ {
                                v111 := GenS4{F0: v5, F1: 979, F2: v4}
                                var v112 int8 = int8(v5)
                                if len(v65) > 0 {
                                    v65[0] = (fi110 & 42)
                                }
                                switch genFn1() {
                                case 2147483647:
                                    if false {
                                        v4 = v4
                                        var v113 int32 = 0
                                        fmt.Println(v2, v113)
                                        _ = v113
                                    } else {
                                        fmt.Println(v2, v4, v2)
                                        v114 := v1[(v4 >> ((42) & 63))]
                                        var v115 float32 = 1e10
                                        v116 := GenS7{F0: ((42 >= 4294967295) || false)}
                                        _ = v114
                                        _ = v115
                                        _ = v116
                                    }
                                case -1:
                                    fmt.Println(v3, fi110)
                                case 4294967295:
                                    var v117 bool = false
                                    _ = v117
                                default:
                                    v3 -= v3
                                }
                                v3 += v3
                                _ = fi110
                                _ = v111
                                _ = v112
                            }
                            v5 = 42
                        }
                    default:
                        var v118 uint = 0
                        _ = v118
                    }
                } else {
                    fmt.Println(v5, v5, v2)
                    v3 -= -2.25
                    v4 = v4
                }
                for fi119 := 0; fi119 < 18; fi119++ {
                    switch -493 {
                    case 127:
                        v5 = 536
                    case 32767:
                        var v120 uint8 = 31
                        _ = v120
                    case 42:
                        v121 := GenS9{F0: false, F1: "a"}
                        _ = v121
                    }
                    var v122 uint = 4294967295
                    _ = fi119
                    _ = v122
                }
                switch fi66 {
                case 2147483647:
                    for fi123 := 0; fi123 < 20; fi123++ {
                        if (true || (false || true)) {
                            fmt.Println(fi123, v2)
                            var v124 string = "a"
                            _ = v124
                        }
                        var v125 bool = false
                        _ = fi123
                        _ = v125
                    }
                case -905:
                    for fi126 := 0; fi126 < 9; fi126++ {
                        var v127 int64 = 42
                        var v128 string = "\n\t"
                        v129 := GenS10{F0: 1, F1: (false && true)}
                        v130 := GenS11{F0: 0, F1: ((v2 < v2) && false)}
                        var v131 int8 = int8(v4)
                        _ = fi126
                        _ = v127
                        _ = v128
                        _ = v129
                        _ = v130
                        _ = v131
                    }
                }
                _ = fi66
            }
            _ = v65
        }
        var v132 int8 = (-(v2))
        _ = v132
    }
    switch "foo bar" {
    case "":
        if (42 != 2) {
            v133 := GenS3{F0: uint(v5), F1: -9223372036854775808}
            _ = v133
        } else {
            fmt.Println(v2, v4)
            _ = len(v1)
            for fi134 := 0; fi134 < 1; fi134++ {
                v5 = (v5 / ((1) | 1))
                _ = fi134
            }
        }
    }
    for fi135 := 0; fi135 < 5; fi135++ {
        switch ("\n\t" + ("hello" + "foo bar")) {
        case "a":
            var v136 uint8 = 255
            _ = v136
        default:
            if true {
                for fi137 := 0; fi137 < 1; fi137++ {
                    v3 += 1e10
                    v138 := GenS1{F0: (v2 * v2)}
                    if false {
                        if ((-(v2)) < int8(v5)) {
                            v139 := []int{(fi137 / ((fi135) | 1))}
                            _ = v139
                        } else {
                            v2 -= (v2 & (v2 << ((65535) & 63)))
                            fmt.Println(fi137, v5, v4)
                            if true {
                                v140 := GenS4{F0: v5, F1: 581, F2: 4294967295}
                                _ = v140
                            } else {
                                v5 -= v5
                            }
                        }
                        v4 -= v4
                        if true {
                            v141 := v1[2]
                            v142 := GenS9{F0: ((fi135 % ((-1) | 1)) != 1), F1: (v141 + (v141 + "\n\t"))}
                            v143 := GenS11{F0: -393, F1: true}
                            _ = v141
                            _ = v142
                            _ = v143
                        }
                        fmt.Println(v5, v3, v2)
                    }
                    _ = fi137
                    _ = v138
                }
            } else {
                for fi144 := 0; fi144 < 11; fi144++ {
                    v145 := GenS12{F0: fi144, F1: 893, F2: -2147483648}
                    _ = fi144
                    _ = v145
                }
                switch (("a" + "foo bar") + "") {
                case "hello":
                    var v146 uint = 18446744073709551615
                    _ = v146
                default:
                    switch "foo bar" {
                    case "a":
                        for fi147 := 0; fi147 < 5; fi147++ {
                            v148 := GenS9{F0: (!true), F1: "a"}
                            v1[42] = "\n\t"
                            _ = fi147
                            _ = v148
                        }
                    default:
                        v4 += v4
                    }
                }
                _ = len(v1)
            }
        }
        if false {
            var v149 uint8 = 29
            v149 -= 2
            _ = v149
        } else {
            switch (fi135 * -305) {
            case 128:
                for fi150 := 0; fi150 < 16; fi150++ {
                    switch fi150 {
                    case -2147483648:
                        var v151 float64 = ((1e10 / (v3 / (v3 * 3.14159))) - (1.5 + v3))
                        _ = v151
                    case 4294967295:
                        var v152 int8 = (v2 & v2)
                        _ = v152
                    }
                    v153 := GenS7{F0: false}
                    fmt.Println(fi135)
                    _ = fi150
                    _ = v153
                }
            }
            fmt.Println(fi135)
        }
        _ = fi135
    }
    for fi154 := 0; fi154 < 7; fi154++ {
        var v155 int8 = -3
        var v156 float64 = ((float64(((1e10 / (-2.25 + 3.14159)) * (0.1 + (float32((1e10 * -2.25)) * -2.25)))) + ((((v3 * (1e10 + (v3 * v3))) + 1.5) + (1.5 / v3)) + -2.25)) - v3)
        for fi157 := 0; fi157 < 4; fi157++ {
            _ = len(v1)
            var v158 uint64 = 65
            _ = fi157
            _ = v158
        }
        _ = fi154
        _ = v155
        _ = v156
    }
    v5 -= v5
    var v159 int32 = 102
    v4 += 557
    if (false || false) {
        v160 := GenS13{F0: (v2 / ((v2) | 1)), F1: 128, F2: true}
        _ = v160
    } else {
        v5 = v5
        v161 := GenS13{F0: v2, F1: 255, F2: (v159 >= -1)}
        for fi162 := 0; fi162 < 3; fi162++ {
            for fi163 := 0; fi163 < 9; fi163++ {
                v164 := GenS10{F0: 25, F1: true}
                v165 := GenS4{F0: 839, F1: 2, F2: 1}
                _ = fi163
                _ = v164
                _ = v165
            }
            v166 := GenS12{F0: (fi162 - (-(fi162))), F1: 4294967295, F2: v159}
            v167 := GenS6{F0: 4294967295, F1: ("" + ("a" + "\n\t")), F2: (^(v159))}
            _ = fi162
            _ = v166
            _ = v167
        }
        _ = v161
    }
    switch "a" {
    case "foo bar":
        switch -2147483648 {
        case 694:
            switch -477 {
            case 32768:
                v2 += -47
            case -9223372036854775808:
                switch genFn1() {
                case -2147483648:
                    v168 := []int{1, 32767, 4294967295, 4294967295}
                    _ = v168
                case -837:
                    v5 += v5
                case 227:
                    v159 -= (v159 & 2147483647)
                default:
                    if (54 == 206) {
                        v4 -= (^(v4))
                    }
                }
            case 127:
                switch (((((("" + "foo bar") + "\n\t") + "") + "hello") + "a") + "foo bar") {
                case "":
                    fmt.Println(v3, v159, v4)
                }
            default:
                var v169 uint16 = (v5 % (((v5 / ((680) | 1))) | 1))
                _ = v169
            }
        default:
            var v170 int32 = int32(v5)
            _ = v170
        }
    case "\n\t":
        v3 -= v3
    default:
        for fi171 := 0; fi171 < 18; fi171++ {
            v172 := GenS12{F0: (fi171 - 9223372036854775807), F1: 525, F2: v159}
            fmt.Println(v3)
            _ = fi171
            _ = v172
        }
    }
    v2 = v2
    var v173 int64 = -9223372036854775808
    v1[(v4 - 65535)] = "a"
    if (((-2.25 - v3) * v3) <= v3) {
        fmt.Println(v3, v159)
        v3 -= 0.1
    } else {
        for fi174 := 0; fi174 < 6; fi174++ {
            if false {
                if true {
                    v5 -= v5
                    v3 += -2.25
                    v3 += (1e10 + ((float64((0.1 + (float32((float64(float32(float64((-2.25 - 0.1)))) - v3)) / 1.0))) * float64(float32((1e10 * v3)))) + v3))
                }
                for fi175 := 0; fi175 < 18; fi175++ {
                    delete(v1, (v4 + (^(v4))))
                    var v176 int8 = v2
                    _ = fi175
                    _ = v176
                }
            }
            v3 -= (v3 - 1.5)
            _ = fi174
        }
    }
    var v177 uint8 = 42
    var v178 uint16 = 1
    v5 += (v5 << ((0) & 63))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v159:", v159)
    fmt.Println("v173:", v173)
    fmt.Println("v177:", v177)
    fmt.Println("v178:", v178)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v159
    _ = v173
    _ = v177
    _ = v178
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
