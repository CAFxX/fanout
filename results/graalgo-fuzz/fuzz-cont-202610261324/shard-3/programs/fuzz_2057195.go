package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16 }
type GenS2 struct { F0 uint32; F1 int64; F2 int16 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int64; F1 uint8; F2 int8 }
type GenS7 struct { F0 uint; F1 uint8 }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 string; F1 bool }
type GenS10 struct { F0 string }
type GenS11 struct { F0 uint32; F1 int64; F2 uint32 }
type GenS12 struct { F0 uint16; F1 uint64 }
func genFn0() uint8 {
    var v4 uint = 0
    if false {
        v5 := GenS1{F0: false, F1: 42}
        _ = v5
    }
    _ = v4
    return 127
}
func genFn1(p0 int64, p1 int8, p2 string) int32 {
    p1 -= (p1 << ((978) & 63))
    for fi6 := 0; fi6 < 18; fi6++ {
        fmt.Println(p0)
        v7 := GenS1{F0: (-48 != p1), F1: 0}
        _ = fi6
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    return 368
}

func main() {
    var v1 int16 = 247
    var v2 int = 0
    var v3 float32 = -2.25
    for fi8 := 0; fi8 < 11; fi8++ {
        v9 := GenS1{F0: (!(false && (!true))), F1: 65535}
        if true {
            v10 := []int{(fi8 << ((0) & 63)), (v2 % ((fi8) | 1))}
            if ((!(false || (v3 <= (v3 - float32((-2.25 * 0.0)))))) && false) {
                var v11 uint16 = 32767
                fmt.Println(v1)
                _ = v11
            }
            _ = v10
        } else {
            v1 -= v1
            v1 = 844
            if true {
                var v12 int8 = 45
                var v13 bool = false
                _ = v12
                _ = v13
            }
        }
        _ = fi8
        _ = v9
    }
    v1 -= v1
    if (((0.1 * -2.25) + float64(float32(1.5))) > ((0.1 * 1e10) + -2.25)) {
        fmt.Println(v2, v3)
        v14 := GenS1{F0: false, F1: 32767}
        v3 -= (float32(3.14159) - 1.5)
        _ = v14
    } else {
        switch (("a" + "") + ("hello" + "a")) {
        case "foo bar":
            if (v1 <= (v1 % ((v1) | 1))) {
                var v15 int32 = 726
                switch ("hello" + ((("foo bar" + "") + "\n\t") + "foo bar")) {
                case "\n\t":
                    for fi16 := 0; fi16 < 17; fi16++ {
                        switch (v2 ^ fi16) {
                        case 4294967295:
                            if (248 >= -472) {
                                v17 := GenS2{F0: 2, F1: 280, F2: (v1 - v1)}
                                v18 := GenS2{F0: 42, F1: 1, F2: 236}
                                v19 := GenS2{F0: 2147483647, F1: -9223372036854775808, F2: (-(v1))}
                                _ = v17
                                _ = v18
                                _ = v19
                            } else {
                                for fi20 := 0; fi20 < 19; fi20++ {
                                    var v21 int = v2
                                    if (!true) {
                                        var v22 int64 = 9223372036854775807
                                        _ = v22
                                    } else {
                                        if true {
                                            if false {
                                                switch "hello" {
                                                case "a":
                                                    var v23 int = (-(fi20))
                                                    _ = v23
                                                default:
                                                    var v24 int32 = (v15 / ((-2147483648) | 1))
                                                    _ = v24
                                                }
                                                v25 := []int{v21, fi16, -2147483648, fi16}
                                                v26 := GenS1{F0: false, F1: 255}
                                                _ = v25
                                                _ = v26
                                            }
                                        } else {
                                            if (false || (!false)) {
                                                var v27 string = ("a" + ("hello" + "a"))
                                                switch (fi16 | (fi16 + -843)) {
                                                case 1:
                                                    v28 := GenS3{F0: 9223372036854775807}
                                                    _ = v28
                                                }
                                                fmt.Println(fi16)
                                                var v29 float32 = v3
                                                var v30 uint32 = 1
                                                _ = v27
                                                _ = v29
                                                _ = v30
                                            }
                                            if false {
                                                if (65535 <= 0) {
                                                    switch ("" + ("foo bar" + "\n\t")) {
                                                    case "foo bar":
                                                        v31 := GenS4{F0: 195}
                                                        v32 := GenS4{F0: 0}
                                                        _ = v31
                                                        _ = v32
                                                    }
                                                }
                                                fmt.Println(fi20)
                                            } else {
                                                if true {
                                                    fmt.Println(fi16, v1)
                                                } else {
                                                    if true {
                                                        v21 = 32768
                                                        var v33 bool = true
                                                        v34 := GenS3{F0: -654}
                                                        _ = v33
                                                        _ = v34
                                                    } else {
                                                        if (false || (!(false && (float64(1.5) > 1.5)))) {
                                                            v3 = v3
                                                        } else {
                                                            v2 = v21
                                                            v35 := []int{fi20}
                                                            _ = v35
                                                        }
                                                        if (175 != 1) {
                                                            fmt.Println(v1)
                                                            v36 := []int{v2}
                                                            if true {
                                                                if (1 > genFn0()) {
                                                                    var v37 string = (("foo bar" + (("\n\t" + (("" + "") + ("a" + ""))) + "")) + ("a" + "\n\t"))
                                                                    v21 = v21
                                                                    fmt.Println(fi16, fi20)
                                                                    _ = v37
                                                                }
                                                            }
                                                            v38 := GenS5{F0: 0}
                                                            var v39 float64 = 0.0
                                                            switch (("\n\t" + (("\n\t" + "foo bar") + "foo bar")) + ("foo bar" + ("hello" + ("\n\t" + ("hello" + "foo bar"))))) {
                                                            case "":
                                                                if false {
                                                                    var v40 string = "hello"
                                                                    _ = v40
                                                                }
                                                            case "foo bar":
                                                                if len(v36) > 0 {
                                                                    v36[0] = (^(fi16))
                                                                }
                                                            case "a":
                                                                v3 += v3
                                                                v3 = ((0.0 / (((float32((0.0 + ((v39 / float64((v3 + float32(float64(((float32(v39) * (0.0 - v3)) / 1.0)))))) * v39))) * v3) * v3) * v3)) + (0.1 * float32(v39)))
                                                                var v41 bool = true
                                                                _ = v41
                                                            default:
                                                                var v42 uint = 255
                                                                _ = v42
                                                            }
                                                            _ = v36
                                                            _ = v38
                                                            _ = v39
                                                        }
                                                    }
                                                    if false {
                                                        var v43 uint16 = 424
                                                        var v44 int16 = v1
                                                        var v45 int8 = 1
                                                        var v46 uint16 = 1
                                                        _ = v43
                                                        _ = v44
                                                        _ = v45
                                                        _ = v46
                                                    } else {
                                                        if ((("hello" + (("hello" + (("\n\t" + (("foo bar" + ("\n\t" + "\n\t")) + "\n\t")) + "hello")) + "a")) + "") < "foo bar") {
                                                            switch (v2 | -2147483648) {
                                                            case 326:
                                                                fmt.Println(v1, fi16, v1)
                                                            case 0:
                                                                v47 := GenS1{F0: (-89 <= 0), F1: 255}
                                                                _ = v47
                                                            case -9223372036854775808:
                                                                fmt.Println(fi16, fi16, v2)
                                                                var v48 string = "\n\t"
                                                                if false {
                                                                    fmt.Println(v3, fi20, v2)
                                                                    switch ((("a" + (v48 + ("a" + v48))) + ((v48 + v48) + ("" + ((v48 + v48) + v48)))) + (v48 + v48)) {
                                                                    case "a":
                                                                        v49 := GenS1{F0: (!(!true)), F1: 255}
                                                                        _ = v49
                                                                    case "\n\t":
                                                                        if true {
                                                                            var v50 int32 = -884
                                                                            v51 := GenS5{F0: 1}
                                                                            v52 := []int{}
                                                                            _ = v50
                                                                            _ = v51
                                                                            _ = v52
                                                                        }
                                                                    default:
                                                                        if true {
                                                                            fmt.Println(v21, v2, v21)
                                                                        } else {
                                                                            fmt.Println(v3)
                                                                            if true {
                                                                                v48 = ""
                                                                                var v53 int32 = (v15 + (-(v15)))
                                                                                v54 := GenS3{F0: -9223372036854775808}
                                                                                v55 := GenS1{F0: false, F1: 1}
                                                                                v21 = 32768
                                                                                v56 := GenS3{F0: -1}
                                                                                var v57 uint64 = 4294967295
                                                                                _ = v53
                                                                                _ = v54
                                                                                _ = v55
                                                                                _ = v56
                                                                                _ = v57
                                                                            } else {
                                                                                fmt.Println(fi20)
                                                                                var v58 string = "\n\t"
                                                                                var v59 int64 = -1
                                                                                _ = v58
                                                                                _ = v59
                                                                            }
                                                                        }
                                                                    }
                                                                    if false {
                                                                        v60 := []int{(-(fi16)), (fi20 - (fi16 | (v21 + fi16)))}
                                                                        var v61 uint8 = 99
                                                                        var v62 bool = false
                                                                        _ = v60
                                                                        _ = v61
                                                                        _ = v62
                                                                    }
                                                                }
                                                                var v63 int64 = 1
                                                                v15 -= (v15 + genFn1((^(v63)), 118, "hello"))
                                                                _ = v48
                                                                _ = v63
                                                            }
                                                            var v64 uint = 255
                                                            if true {
                                                                switch ("\n\t" + "a") {
                                                                case "hello":
                                                                    switch (-(fi16)) {
                                                                    case -543:
                                                                        if false {
                                                                            v65 := []int{(fi20 >> ((v64) & 63)), fi16}
                                                                            if ((v1 != -1) && true) {
                                                                                var v66 float64 = float64((float32(0.0) * v3))
                                                                                var v67 uint8 = 0
                                                                                var v68 uint32 = 2147483647
                                                                                _ = v66
                                                                                _ = v67
                                                                                _ = v68
                                                                            } else {
                                                                                fmt.Println(fi16)
                                                                            }
                                                                            _ = v65
                                                                        } else {
                                                                            v2 += v21
                                                                            fmt.Println(v21, v64)
                                                                            if ((v64 * 645) == (v64 | v64)) {
                                                                                var v69 int = (v2 / ((int(v1)) | 1))
                                                                                v70 := GenS1{F0: false, F1: uint16(v64)}
                                                                                var v71 uint16 = 329
                                                                                _ = v69
                                                                                _ = v70
                                                                                _ = v71
                                                                            }
                                                                        }
                                                                    case -965:
                                                                        v72 := []int{v2}
                                                                        _ = v72
                                                                    case -631:
                                                                        var v73 float32 = v3
                                                                        _ = v73
                                                                    default:
                                                                        var v74 uint64 = 4294967295
                                                                        _ = v74
                                                                    }
                                                                case "\n\t":
                                                                    if (148 < 227) {
                                                                        v75 := []int{(-(v2)), 42, 42}
                                                                        v76 := GenS1{F0: false, F1: 32767}
                                                                        v15 -= (v15 & v15)
                                                                        v77 := GenS3{F0: 0}
                                                                        _ = v75
                                                                        _ = v76
                                                                        _ = v77
                                                                    }
                                                                case "":
                                                                    fmt.Println(v64, fi20)
                                                                default:
                                                                    v1 -= v1
                                                                }
                                                            }
                                                            _ = v64
                                                        } else {
                                                            if (-1 <= 1) {
                                                                if true {
                                                                    fmt.Println(v3)
                                                                }
                                                                if false {
                                                                    var v78 int16 = v1
                                                                    v79 := GenS4{F0: 127}
                                                                    var v80 uint16 = 589
                                                                    _ = v78
                                                                    _ = v79
                                                                    _ = v80
                                                                } else {
                                                                    if true {
                                                                        fmt.Println(v3, v21)
                                                                        var v81 uint32 = 4294967295
                                                                        v82 := GenS1{F0: (!false), F1: 0}
                                                                        _ = v81
                                                                        _ = v82
                                                                    }
                                                                    if ("\n\t" < "hello") {
                                                                        v1 += (v1 ^ (v1 ^ v1))
                                                                        var v83 bool = ((float32(0.0) * 0.1) <= float32((1.5 * -2.25)))
                                                                        fmt.Println(v83, v2, v21)
                                                                        _ = v83
                                                                    }
                                                                    v15 = (v15 / ((2147483647) | 1))
                                                                }
                                                                v15 = (^(v15))
                                                            } else {
                                                                v84 := []int{(-(v21))}
                                                                var v85 int64 = 1
                                                                _ = v84
                                                                _ = v85
                                                            }
                                                            fmt.Println(v3, fi16)
                                                            v86 := GenS6{F0: -9223372036854775808, F1: 128, F2: 100}
                                                            _ = v86
                                                        }
                                                        v87 := GenS3{F0: 9223372036854775807}
                                                        switch "foo bar" {
                                                        case "foo bar":
                                                            switch 0 {
                                                            case 4294967295:
                                                                v2 = fi20
                                                            default:
                                                                var v88 int = v21
                                                                _ = v88
                                                            }
                                                        case "\n\t":
                                                            var v89 uint16 = 2
                                                            _ = v89
                                                        default:
                                                            v90 := []int{v21, -439, 601, (^(v2))}
                                                            _ = v90
                                                        }
                                                        _ = v87
                                                    }
                                                    if false {
                                                        var v91 int8 = 42
                                                        v92 := GenS7{F0: 906, F1: 0}
                                                        fmt.Println(v21, fi16, v1)
                                                        _ = v91
                                                        _ = v92
                                                    } else {
                                                        fmt.Println(v21, fi16, v1)
                                                        var v93 uint64 = 42
                                                        _ = v93
                                                    }
                                                    fmt.Println(v3, fi20, v3)
                                                    v1 = 1
                                                    v94 := []int{}
                                                    _ = v94
                                                }
                                            }
                                        }
                                        v2 -= fi20
                                        var v95 uint32 = 65535
                                        switch "hello" {
                                        case "foo bar":
                                            if (0 != 230) {
                                                switch -2147483648 {
                                                case -354:
                                                    v96 := GenS6{F0: 9223372036854775807, F1: 2, F2: 0}
                                                    _ = v96
                                                case 9223372036854775807:
                                                    v3 += v3
                                                default:
                                                    v97 := GenS8{F0: int64(v95)}
                                                    switch (fi20 << ((4294967295) & 63)) {
                                                    case -354:
                                                        var v98 int32 = (v15 ^ 0)
                                                        _ = v98
                                                    case 0:
                                                        if false {
                                                            v99 := GenS5{F0: -128}
                                                            _ = v99
                                                        }
                                                    default:
                                                        var v100 float64 = 1.5
                                                        _ = v100
                                                    }
                                                    _ = v97
                                                }
                                                v101 := GenS8{F0: 0}
                                                var v102 int8 = 7
                                                v103 := GenS2{F0: 2147483648, F1: 1, F2: (^(v1))}
                                                _ = v101
                                                _ = v102
                                                _ = v103
                                            } else {
                                                if false {
                                                    if false {
                                                        switch 472 {
                                                        case 720:
                                                            var v104 string = ((("\n\t" + ("" + "hello")) + "foo bar") + "foo bar")
                                                            _ = v104
                                                        }
                                                        v95 += (^(v95))
                                                        switch "foo bar" {
                                                        case "\n\t":
                                                            v105 := GenS9{F0: "hello", F1: (v15 != v15)}
                                                            _ = v105
                                                        case "a":
                                                            v106 := []int{}
                                                            _ = v106
                                                        case "foo bar":
                                                            fmt.Println(v2)
                                                        }
                                                    } else {
                                                        v15 = genFn1(986, -1, "a")
                                                    }
                                                    var v107 int8 = 100
                                                    var v108 int = fi16
                                                    _ = v107
                                                    _ = v108
                                                } else {
                                                    switch fi16 {
                                                    case -427:
                                                        v15 = (v15 % ((v15) | 1))
                                                    case -380:
                                                        v109 := GenS7{F0: 175, F1: 127}
                                                        _ = v109
                                                    case 4294967295:
                                                        v1 -= (v1 - v1)
                                                    }
                                                    switch (^(v21)) {
                                                    case 547:
                                                        var v110 uint8 = 127
                                                        _ = v110
                                                    case -72:
                                                        if ((v15 | genFn1(1, 42, "a")) < v15) {
                                                            var v111 string = "a"
                                                            switch "hello" {
                                                            case "":
                                                                switch v111 {
                                                                case "\n\t":
                                                                    var v112 int64 = -984
                                                                    _ = v112
                                                                }
                                                            case "\n\t":
                                                                v21 -= (v2 >> ((18446744073709551615) & 63))
                                                            }
                                                            _ = v111
                                                        } else {
                                                            switch "" {
                                                            case "":
                                                                v15 -= (v15 * -1)
                                                            default:
                                                                fmt.Println(v15)
                                                            }
                                                        }
                                                    }
                                                    v21 = (fi20 & 112)
                                                }
                                            }
                                        default:
                                            var v113 bool = (!(true || (!(1 > 65535))))
                                            _ = v113
                                        }
                                        _ = v95
                                    }
                                    v114 := []int{fi20, fi20}
                                    _ = fi20
                                    _ = v114
                                    _ = v21
                                }
                            }
                        case 127:
                            if (42 == 100) {
                                v3 = v3
                                var v115 int = int(v15)
                                var v116 int64 = -9223372036854775808
                                _ = v115
                                _ = v116
                            } else {
                                v117 := GenS10{F0: ""}
                                if false {
                                    fmt.Println(v1, v3, fi16)
                                } else {
                                    var v118 int = fi16
                                    v1 = (^(v1))
                                    var v119 int8 = 42
                                    _ = v118
                                    _ = v119
                                }
                                _ = v117
                            }
                        }
                        _ = fi16
                    }
                case "a":
                    fmt.Println(v1, v15, v1)
                case "foo bar":
                    v120 := []int{v2, (v2 * (v2 << ((255) & 63))), 2147483647}
                    _ = v120
                default:
                    var v121 int = v2
                    _ = v121
                }
                _ = v15
            } else {
                var v122 int = (-(v2))
                fmt.Println(v3)
                _ = v122
            }
        case "hello":
            var v123 int8 = -108
            _ = v123
        case "\n\t":
            v2 = (v2 >> ((1) & 63))
        default:
            for fi124 := 0; fi124 < 8; fi124++ {
                switch 128 {
                case 127:
                    fmt.Println(v3, fi124)
                case 1:
                    var v125 int = (v2 >> ((122) & 63))
                    _ = v125
                }
                v3 -= v3
                _ = fi124
            }
        }
        v3 = (-2.25 * ((v3 + 1e10) * v3))
        var v126 uint32 = 65535
        _ = v126
    }
    var v127 float32 = 1e10
    if true {
        var v128 int64 = -251
        for fi129 := 0; fi129 < 6; fi129++ {
            var v130 int32 = 0
            v3 += -2.25
            _ = fi129
            _ = v130
        }
        _ = v128
    }
    v3 += ((0.0 / 1.5) - (0.0 - float32(float64((1e10 / 1.5)))))
    for fi131 := 0; fi131 < 10; fi131++ {
        fmt.Println(v2, v1)
        v132 := []int{v2, fi131}
        _ = fi131
        _ = v132
    }
    var v133 int = v2
    var v134 string = ("" + ("" + "\n\t"))
    if false {
        var v135 float32 = ((v3 * 1.5) - 1e10)
        fmt.Println(v134, v1)
        _ = v135
    } else {
        v133 = v2
        for fi136 := 0; fi136 < 15; fi136++ {
            v137 := GenS6{F0: 42, F1: genFn0(), F2: -1}
            _ = fi136
            _ = v137
        }
        var v138 string = ((((v134 + ((((("\n\t" + (v134 + v134)) + v134) + (((v134 + "a") + (v134 + "foo bar")) + "")) + v134) + (("" + v134) + "foo bar"))) + (v134 + v134)) + v134) + (v134 + v134))
        _ = v138
    }
    var v139 string = "a"
    for fi140 := 0; fi140 < 5; fi140++ {
        v2 = v2
        switch v2 {
        case 0:
            if true {
                v141 := GenS10{F0: "foo bar"}
                fmt.Println(v133)
                _ = v141
            }
        case 9223372036854775807:
            if (!(true || false)) {
                for fi142 := 0; fi142 < 4; fi142++ {
                    if false {
                        var v143 bool = (223 > 4294967295)
                        switch fi142 {
                        case 32768:
                            if (!(9223372036854775807 > 9223372036854775807)) {
                                v144 := []int{609, (^(fi142)), (v133 - int(v1))}
                                _ = v144
                            }
                        case 0:
                            switch (fi142 & (v133 | v133)) {
                            case 293:
                                var v145 uint32 = 293
                                _ = v145
                            case 42:
                                var v146 uint64 = 2
                                _ = v146
                            case -732:
                                v139 = "\n\t"
                            }
                        case 2147483647:
                            fmt.Println(v3, fi142, v1)
                        }
                        _ = v143
                    }
                    _ = fi142
                }
                if (float64(1e10) < ((-2.25 * -2.25) - 0.1)) {
                    switch v134 {
                    case "foo bar":
                        v133 += v2
                    case "hello":
                        v147 := []int{(fi140 & -2147483648), (^(fi140)), -730, (^(v133))}
                        _ = v147
                    default:
                        if true {
                            v148 := []int{32768}
                            for fi149 := 0; fi149 < 11; fi149++ {
                                var v150 float32 = (v127 + 3.14159)
                                _ = fi149
                                _ = v150
                            }
                            _ = v148
                        }
                    }
                    var v151 int64 = -9223372036854775808
                    _ = v151
                } else {
                    if ((float64(1e10) + 3.14159) >= -2.25) {
                        v152 := GenS11{F0: 2147483647, F1: 1, F2: 532}
                        switch v139 {
                        case "\n\t":
                            for fi153 := 0; fi153 < 10; fi153++ {
                                v127 = v3
                                switch v134 {
                                case "a":
                                    var v154 int32 = 2147483647
                                    _ = v154
                                }
                                _ = fi153
                            }
                        case "a":
                            var v155 int64 = 1
                            _ = v155
                        }
                        _ = v152
                    }
                }
            } else {
                fmt.Println(fi140, v2, v134)
            }
        case 2147483647:
            v156 := GenS9{F0: (v134 + (v134 + v139)), F1: false}
            _ = v156
        }
        _ = fi140
    }
    if (true && false) {
        v157 := GenS12{F0: 429, F1: 282}
        v158 := []int{v133, 128, v133}
        var v159 uint = 255
        _ = v157
        _ = v158
        _ = v159
    }
    fmt.Println("v1:", v1)
    fmt.Println("v127:", v127)
    fmt.Println("v133:", v133)
    fmt.Println("v134:", v134)
    fmt.Println("v139:", v139)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v127
    _ = v133
    _ = v134
    _ = v139
    _ = v2
    _ = v3
}
