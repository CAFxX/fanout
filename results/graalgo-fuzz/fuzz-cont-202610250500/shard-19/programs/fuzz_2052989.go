package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16; F2 uint8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint; F1 int16 }
type GenS6 struct { F0 uint; F1 int32; F2 uint16 }
type GenS7 struct { F0 uint32; F1 int64; F2 string }
type GenS8 struct { F0 int16; F1 bool }
type GenS9 struct { F0 int32; F1 string }
type GenS10 struct { F0 int32; F1 int16; F2 int32 }
type GenS11 struct { F0 int16; F1 int64; F2 int8 }
type GenS12 struct { F0 string; F1 int32; F2 int }
type GenS13 struct { F0 string; F1 int16; F2 uint64 }
type GenS14 struct { F0 string }
type GenS15 struct { F0 uint32; F1 uint }
type GenS16 struct { F0 int8; F1 string; F2 uint32 }
type GenS17 struct { F0 int32 }
type GenS18 struct { F0 int }
type GenS19 struct { F0 int64; F1 int8 }
func genFn0(p0 uint32, p1 bool) uint8 {
    p1 = (12 != 1)
    v5 := GenS1{F0: 42, F1: 32767, F2: 2}
    _ = p0
    _ = p1
    _ = v5
    return 2
}
func genFn1() int {
    var v6 uint64 = 42
    v6 += (v6 + 824)
    _ = v6
    return 32768
}

func main() {
    var v1 int8 = -1
    v2 := make(map[int]string)
    var v3 int16 = -1
    v4 := make(map[uint32]string)
    v7 := GenS2{F0: 1}
    if true {
        if (2 <= 127) {
            var v8 uint32 = 2
            var v9 int32 = 42
            v9 = (v9 >> ((331) & 63))
            _ = v8
            _ = v9
        } else {
            for fi10 := 0; fi10 < 14; fi10++ {
                v11 := []int{genFn1()}
                _ = fi10
                _ = v11
            }
            v4[1] = "a"
            fmt.Println(v1)
        }
        for fi12 := 0; fi12 < 10; fi12++ {
            for fi13 := 0; fi13 < 5; fi13++ {
                switch "\n\t" {
                case "":
                    var v14 uint32 = 4294967295
                    _ = v14
                }
                _ = fi13
            }
            v3 += -476
            v1 += v1
            _ = fi12
        }
    } else {
        for fi15 := 0; fi15 < 5; fi15++ {
            fmt.Println(fi15, v3, v3)
            _ = fi15
        }
    }
    v3 += 49
    v1 = int8(v3)
    for fi16 := 0; fi16 < 18; fi16++ {
        _ = len(v4)
        _ = fi16
    }
    fmt.Println(v1, v3)
    for fi17 := 0; fi17 < 4; fi17++ {
        if (!true) {
            if false {
                if true {
                    v1 = (v1 & v1)
                    fmt.Println(v1, v3, fi17)
                    var v18 string = "hello"
                    _ = v18
                } else {
                    var v19 uint = 4294967295
                    _ = v19
                }
                if true {
                    if (!(-2.25 < 0.1)) {
                        _ = len(v4)
                    } else {
                        if (!(42 != 32767)) {
                            v20 := []int{fi17, 1, 32768, fi17}
                            _ = v20
                        } else {
                            if (!(!(0 == 1))) {
                                var v21 uint8 = 255
                                var v22 int64 = 0
                                _ = v21
                                _ = v22
                            }
                            for fi23 := 0; fi23 < 6; fi23++ {
                                fmt.Println(v1, v1, fi17)
                                _ = fi23
                            }
                        }
                    }
                }
            }
            fmt.Println(fi17, fi17, fi17)
        } else {
            if false {
                switch ("\n\t" + "hello") {
                case "a":
                    fmt.Println(v1)
                case "foo bar":
                    switch fi17 {
                    case 4294967295:
                        for fi24 := 0; fi24 < 2; fi24++ {
                            v25 := GenS1{F0: -690, F1: v3, F2: genFn0(2147483647, true)}
                            v1 = v1
                            v1 = int8(v3)
                            _ = fi24
                            _ = v25
                        }
                    case 32768:
                        if (genFn0(23, false) != 128) {
                            switch genFn1() {
                            case 2147483647:
                                v26 := []int{-563}
                                _ = v26
                            case -30:
                                if true {
                                    v3 = -647
                                } else {
                                    fmt.Println(fi17)
                                }
                            default:
                                var v27 bool = ((^(v1)) < (^(v1)))
                                _ = v27
                            }
                        } else {
                            var v28 int16 = v3
                            for fi29 := 0; fi29 < 6; fi29++ {
                                switch "a" {
                                case "":
                                    fmt.Println(fi17, fi29, v28)
                                default:
                                    v28 -= v3
                                }
                                var v30 int32 = -1
                                var v31 uint16 = 65535
                                _ = fi29
                                _ = v30
                                _ = v31
                            }
                            _ = v28
                        }
                    }
                }
                v3 += -32768
            }
        }
        v1 += v1
        if true {
            switch "foo bar" {
            case "a":
                fmt.Println(v3)
            case "foo bar":
                v1 = v1
            }
            for fi32 := 0; fi32 < 6; fi32++ {
                v33 := []int{(-(fi32)), fi17}
                v2[(fi17 | (fi32 & (fi32 << ((2) & 63))))] = "a"
                v1 += (v1 | (-(v1)))
                v34 := GenS3{F0: 0}
                _ = fi32
                _ = v33
                _ = v34
            }
        } else {
            if true {
                v1 -= (v1 * -1)
                for fi35 := 0; fi35 < 10; fi35++ {
                    if (1e10 < ((0.1 * float32((((0.1 - 0.1) * (3.14159 + 1.5)) - (0.1 + 1e10)))) / 1e10)) {
                        if (0 == 0) {
                            var v36 uint32 = 770
                            v1 -= v1
                            switch (fi17 & fi17) {
                            case -647:
                                v37 := []int{genFn1(), genFn1(), fi35}
                                _ = v37
                            case 243:
                                switch (("a" + "hello") + ("hello" + "hello")) {
                                case "\n\t":
                                    v38 := v2[(fi35 % ((genFn1()) | 1))]
                                    _ = v38
                                }
                            }
                            _ = v36
                        } else {
                            v39 := GenS3{F0: 32767}
                            v40 := GenS4{F0: true}
                            var v41 int8 = (v1 * v1)
                            _ = v39
                            _ = v40
                            _ = v41
                        }
                        if (!false) {
                            if false {
                                if ((!true) || ((fi35 << ((1) & 63)) > 32767)) {
                                    v3 += (v3 - -63)
                                }
                            }
                            var v42 int = 484
                            v1 = v1
                            if false {
                                if (false || false) {
                                    v43 := []int{fi17, fi17}
                                    v44 := GenS3{F0: 2}
                                    v1 -= v1
                                    _ = v43
                                    _ = v44
                                } else {
                                    v45 := GenS5{F0: 4294967295, F1: (v3 - (v3 % (((v3 % (((v3 % (((v3 / (((v3 - (v3 << ((0) & 63)))) | 1))) | 1))) | 1))) | 1)))}
                                    fmt.Println(v1, v1)
                                    switch ("\n\t" + "hello") {
                                    case "foo bar":
                                        v46 := GenS3{F0: 51}
                                        _ = v46
                                    default:
                                        v1 = 1
                                    }
                                    var v47 bool = (!((v3 + (v3 & (v3 - (v3 & v3)))) >= v3))
                                    fmt.Println(v1, fi35, v1)
                                    _ = v45
                                    _ = v47
                                }
                                var v48 int8 = 0
                                fmt.Println(v1)
                                switch "foo bar" {
                                case "\n\t":
                                    v42 += -64
                                case "foo bar":
                                    var v49 uint32 = 2147483647
                                    _ = v49
                                default:
                                    if false {
                                        v42 = (fi17 % ((9223372036854775807) | 1))
                                        switch "foo bar" {
                                        case "hello":
                                            var v50 float32 = -2.25
                                            _ = v50
                                        default:
                                            if false {
                                                switch "a" {
                                                case "a":
                                                    if (-9223372036854775808 >= 128) {
                                                        var v51 uint32 = 1
                                                        if (true && true) {
                                                            v3 += v3
                                                        } else {
                                                            fmt.Println(v51, v51, fi35)
                                                        }
                                                        v52 := GenS2{F0: 0}
                                                        _ = v51
                                                        _ = v52
                                                    } else {
                                                        var v53 uint32 = 429
                                                        var v54 bool = true
                                                        var v55 int32 = -2147483648
                                                        _ = v53
                                                        _ = v54
                                                        _ = v55
                                                    }
                                                case "":
                                                    v56 := GenS6{F0: 2, F1: -434, F2: 478}
                                                    _ = v56
                                                case "\n\t":
                                                    var v57 uint8 = 196
                                                    v58 := GenS7{F0: 4294967295, F1: -9223372036854775808, F2: "\n\t"}
                                                    _ = v57
                                                    _ = v58
                                                default:
                                                    v59 := []int{}
                                                    _ = v59
                                                }
                                                fmt.Println(fi35, v3, v42)
                                                v60 := GenS1{F0: 98, F1: v3, F2: 0}
                                                _ = v60
                                            } else {
                                                v61 := []int{fi35, v42}
                                                _ = v61
                                            }
                                        }
                                        var v62 int32 = 0
                                        _ = v62
                                    } else {
                                        v63 := GenS1{F0: 9223372036854775807, F1: v3, F2: 42}
                                        _ = v63
                                    }
                                }
                                _ = v48
                            } else {
                                switch (("" + "hello") + "") {
                                case "a":
                                    var v64 int32 = 0
                                    _ = v64
                                case "foo bar":
                                    v42 = 4294967295
                                case "hello":
                                    var v65 int16 = v3
                                    var v66 int8 = (v1 ^ v1)
                                    _ = v65
                                    _ = v66
                                }
                            }
                            v67 := GenS8{F0: -916, F1: true}
                            var v68 uint32 = 2147483647
                            v42 = 4294967295
                            _ = v42
                            _ = v67
                            _ = v68
                        } else {
                            v69 := GenS9{F0: -700, F1: ""}
                            if ((-2.25 < 3.14159) && false) {
                                var v70 uint64 = 0
                                _ = v70
                            } else {
                                v3 = 1000
                                v3 = v3
                                var v71 bool = (true && true)
                                _ = v71
                            }
                            _ = v69
                        }
                    } else {
                        var v72 int = (fi17 % (((fi17 | -1)) | 1))
                        _ = len(v4)
                        v1 -= (^(v1))
                        var v73 uint8 = 128
                        var v74 int64 = 26
                        v3 -= -793
                        _ = v72
                        _ = v73
                        _ = v74
                    }
                    switch fi17 {
                    case -487:
                        switch ("foo bar" + ("\n\t" + "a")) {
                        case "hello":
                            var v75 int16 = 1000
                            _ = v75
                        }
                        v3 += v3
                        v76 := GenS10{F0: -2147483648, F1: 32767, F2: -1}
                        var v77 int32 = -1
                        var v78 int = fi35
                        var v79 int8 = v1
                        _ = v76
                        _ = v77
                        _ = v78
                        _ = v79
                    case 874:
                        delete(v4, 42)
                    case -2147483648:
                        v80 := []int{}
                        _ = v80
                    }
                    v81 := v2[genFn1()]
                    v1 -= (^(v1))
                    v82 := []int{42, fi35}
                    fmt.Println(v81, fi35)
                    v83 := GenS3{F0: 380}
                    _ = fi35
                    _ = v81
                    _ = v82
                    _ = v83
                }
                var v84 bool = (!(0 >= 32767))
                _ = v84
            } else {
                for fi85 := 0; fi85 < 10; fi85++ {
                    if ((^(v3)) < (^(v3))) {
                        var v86 int32 = -2147483648
                        v87 := GenS8{F0: v3, F1: true}
                        _ = v86
                        _ = v87
                    } else {
                        v1 = v1
                    }
                    v88 := GenS6{F0: 255, F1: 1, F2: 65535}
                    v3 = v3
                    _ = fi85
                    _ = v88
                }
                delete(v2, fi17)
                switch fi17 {
                case 172:
                    switch "" {
                    case "foo bar":
                        var v89 int = fi17
                        _ = v89
                    default:
                        v90 := []int{(fi17 - -318)}
                        _ = v90
                    }
                default:
                    switch "" {
                    case "a":
                        delete(v2, genFn1())
                    }
                }
            }
            var v91 uint16 = 543
            for fi92 := 0; fi92 < 19; fi92++ {
                v91 = v91
                v91 -= (v91 & 42)
                var v93 string = "\n\t"
                v93 = v93
                v3 += (v3 >> ((4294967295) & 63))
                _ = fi92
                _ = v93
            }
            _ = v91
        }
        _ = fi17
    }
    if (!false) {
        switch 2147483647 {
        case 281:
            if true {
                if true {
                    if false {
                        if (128 <= 4294967295) {
                            var v94 bool = (127 <= genFn0(728, true))
                            _ = v94
                        } else {
                            switch (("hello" + ("\n\t" + ("foo bar" + (("foo bar" + "\n\t") + "a")))) + ("hello" + "")) {
                            case "a":
                                fmt.Println(v3, v1)
                            case "hello":
                                v95 := GenS11{F0: (-(v3)), F1: 42, F2: (v1 >> ((18446744073709551615) & 63))}
                                _ = v95
                            default:
                                var v96 uint64 = 2
                                _ = v96
                            }
                            delete(v4, 2)
                        }
                        for fi97 := 0; fi97 < 19; fi97++ {
                            var v98 uint64 = 2
                            _ = fi97
                            _ = v98
                        }
                        var v99 uint8 = 179
                        _ = v99
                    } else {
                        v100 := GenS12{F0: (("a" + "") + ("hello" + ("foo bar" + "hello"))), F1: 2147483647, F2: 32768}
                        _ = v100
                    }
                    fmt.Println(v3, v1)
                    if true {
                        var v101 float64 = 0.1
                        _ = v101
                    }
                } else {
                    if ((2 >= genFn0(1, (false && false))) || false) {
                        if (false || true) {
                            if true {
                                for fi102 := 0; fi102 < 6; fi102++ {
                                    v103 := GenS13{F0: "\n\t", F1: v3, F2: 2}
                                    for fi104 := 0; fi104 < 6; fi104++ {
                                        var v105 uint64 = 777
                                        var v106 int64 = -878
                                        if (fi104 <= (-(fi102))) {
                                            v107 := []int{}
                                            v108 := GenS13{F0: ((("hello" + "hello") + "hello") + "a"), F1: (v3 ^ -32768), F2: 4294967295}
                                            _ = v107
                                            _ = v108
                                        }
                                        var v109 uint = 1
                                        _ = fi104
                                        _ = v105
                                        _ = v106
                                        _ = v109
                                    }
                                    for fi110 := 0; fi110 < 3; fi110++ {
                                        var v111 int8 = v1
                                        v112 := []int{(fi110 << ((373) & 63)), (fi102 << ((1) & 63)), -1}
                                        _ = fi110
                                        _ = v111
                                        _ = v112
                                    }
                                    _ = fi102
                                    _ = v103
                                }
                            }
                            switch -735 {
                            case -2147483648:
                                v1 -= 42
                            case 42:
                                v113 := GenS10{F0: 2147483647, F1: v3, F2: 54}
                                _ = v113
                            case -630:
                                fmt.Println(v1)
                            }
                            v114 := GenS10{F0: -1, F1: v3, F2: 2147483647}
                            _ = v114
                        }
                        if (42 > 830) {
                            var v115 uint16 = 255
                            v116 := GenS14{F0: ("a" + "foo bar")}
                            _ = v115
                            _ = v116
                        } else {
                            var v117 int8 = (v1 & v1)
                            var v118 string = ""
                            _ = v117
                            _ = v118
                        }
                        var v119 float64 = -2.25
                        _ = v119
                    }
                    v1 -= v1
                }
                fmt.Println(v1, v1)
                if false {
                    v120 := GenS15{F0: 115, F1: 1}
                    v121 := []int{genFn1(), 2147483647, -8, 1}
                    _ = v120
                    _ = v121
                }
            } else {
                if (-278 <= v3) {
                    if ("foo bar" != ("foo bar" + "\n\t")) {
                        v122 := GenS16{F0: 1, F1: ("foo bar" + ""), F2: 4294967295}
                        if ((0.1 * (-2.25 - -2.25)) > ((float64(0.0) - 0.1) / 0.1)) {
                            v123 := GenS12{F0: "\n\t", F1: 42, F2: 128}
                            if (v1 == -128) {
                                if false {
                                    v1 += v1
                                } else {
                                    v3 -= (v3 + v3)
                                    if (3.14159 > (float64((((float32((((3.14159 - (1.5 / float64((-2.25 - 1e10)))) / (3.14159 + 1e10)) / 1.0)) / 0.1) + (3.14159 + 1.5)) + float32(float64(0.1)))) + 3.14159)) {
                                        v124 := GenS17{F0: 42}
                                        if false {
                                            if false {
                                                for fi125 := 0; fi125 < 2; fi125++ {
                                                    v126 := v2[genFn1()]
                                                    if true {
                                                        v3 = (v3 % (((v3 << ((0) & 63))) | 1))
                                                    }
                                                    v127 := GenS13{F0: v126, F1: (^(v3)), F2: 42}
                                                    _ = fi125
                                                    _ = v126
                                                    _ = v127
                                                }
                                                switch 42 {
                                                case 2147483647:
                                                    v128 := GenS18{F0: 128}
                                                    _ = v128
                                                }
                                                if false {
                                                    for fi129 := 0; fi129 < 8; fi129++ {
                                                        for fi130 := 0; fi130 < 15; fi130++ {
                                                            v3 = v3
                                                            var v131 string = ((("\n\t" + "a") + ("" + ("" + "foo bar"))) + "a")
                                                            _ = fi130
                                                            _ = v131
                                                        }
                                                        var v132 uint32 = 2
                                                        v3 = (v3 ^ 32767)
                                                        _ = fi129
                                                        _ = v132
                                                    }
                                                    v3 -= (v3 + v3)
                                                }
                                            } else {
                                                v2[32767] = "foo bar"
                                                for fi133 := 0; fi133 < 5; fi133++ {
                                                    var v134 uint8 = 107
                                                    _ = fi133
                                                    _ = v134
                                                }
                                            }
                                            fmt.Println(v1)
                                            v1 = (v1 * (v1 << ((65535) & 63)))
                                        }
                                        _ = v124
                                    }
                                    v1 = 127
                                }
                                fmt.Println(v1)
                            }
                            _ = len(v2)
                            _ = v123
                        } else {
                            v135 := []int{9223372036854775807}
                            v136 := GenS19{F0: -797, F1: (v1 | (v1 - (v1 | (v1 ^ 0))))}
                            var v137 int16 = int16(v1)
                            _ = v135
                            _ = v136
                            _ = v137
                        }
                        _ = v122
                    } else {
                        var v138 string = "hello"
                        if (0 != -480) {
                            var v139 int32 = 2147483647
                            var v140 string = ((("\n\t" + ("hello" + "a")) + "\n\t") + v138)
                            _ = v139
                            _ = v140
                        } else {
                            v3 = v3
                        }
                        _ = v138
                    }
                    switch -940 {
                    case 2147483647:
                        v141 := GenS16{F0: v1, F1: "", F2: 950}
                        _ = v141
                    case -2147483648:
                        var v142 int32 = 0
                        _ = v142
                    }
                    v143 := GenS14{F0: "\n\t"}
                    _ = v143
                } else {
                    var v144 float64 = (1.5 + -2.25)
                    _ = v144
                }
            }
        case -1:
            var v145 int8 = (v1 << ((18446744073709551615) & 63))
            _ = v145
        case 128:
            if false {
                delete(v4, 1)
                v3 += (v3 | (v3 - (v3 * (v3 / (((v3 + (^(v3)))) | 1)))))
                v3 = v3
            }
        }
        v146 := []int{4294967295}
        var v147 string = ((("\n\t" + (("a" + "hello") + "foo bar")) + "\n\t") + "\n\t")
        _ = v146
        _ = v147
    } else {
        if false {
            var v148 float32 = ((3.14159 + (3.14159 * -2.25)) / 1.0)
            _ = v148
        }
        var v149 uint16 = 255
        _ = v149
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
