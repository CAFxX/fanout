package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint }
type GenS2 struct { F0 int16; F1 int }
type GenS3 struct { F0 uint64; F1 uint16; F2 bool }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int8; F1 int16; F2 int8 }
type GenS6 struct { F0 int8; F1 int16 }
type GenS7 struct { F0 string; F1 string; F2 uint64 }
type GenS8 struct { F0 int }
type GenS9 struct { F0 int64 }
type GenS10 struct { F0 bool; F1 int16 }
type GenS11 struct { F0 uint8; F1 uint8 }
type GenS12 struct { F0 int; F1 uint16 }
type GenS13 struct { F0 uint }
type GenS14 struct { F0 int8; F1 uint }
type GenS15 struct { F0 uint32; F1 int64 }
type GenS16 struct { F0 string; F1 bool }
type GenS17 struct { F0 uint32; F1 int64 }
type GenS18 struct { F0 int8; F1 int64; F2 string }
type GenS19 struct { F0 uint64 }
type GenS20 struct { F0 uint32; F1 bool; F2 uint8 }
func genFn0(p0 uint8, p1 uint) uint32 {
    var v5 int64 = 42
    var v6 string = "\n\t"
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 289
}
func genFn1() bool {
    v7 := GenS1{F0: 9223372036854775807, F1: 128}
    for fi8 := 0; fi8 < 4; fi8++ {
        var v9 string = (("a" + "") + "a")
        _ = fi8
        _ = v9
    }
    v10 := GenS2{F0: 32767, F1: 4294967295}
    var v11 uint8 = 76
    _ = v10
    _ = v11
    _ = v7
    return false
}

func main() {
    var v1 uint64 = 4294967295
    v2 := make(map[string]string)
    var v3 int64 = 42
    v4 := make(map[uint32]uint32)
    for fi12 := 0; fi12 < 6; fi12++ {
        var v13 int8 = -1
        if (true && true) {
            v1 = v1
            var v14 float64 = 0.1
            _ = v14
        } else {
            v15 := GenS3{F0: 42, F1: 65535, F2: true}
            if true {
                if false {
                    switch "" {
                    case "\n\t":
                        fmt.Println(v1, fi12)
                    case "a":
                        if (217 == uint8(v1)) {
                            if false {
                                var v16 uint64 = v1
                                _ = v16
                            } else {
                                v17 := GenS2{F0: -32768, F1: (fi12 | 128)}
                                _ = v17
                            }
                            fmt.Println(v1, v1, v3)
                            v18 := GenS4{F0: 2147483647}
                            _ = v18
                        } else {
                            var v19 int = 590
                            switch -2147483648 {
                            case -491:
                                if (false && true) {
                                    v1 = v1
                                    var v20 int8 = (-(v13))
                                    _ = v20
                                } else {
                                    if true {
                                        var v21 uint64 = (v1 | (v1 >> ((255) & 63)))
                                        _ = v21
                                    }
                                    v3 -= (v3 + (v3 ^ -9223372036854775808))
                                }
                            case 909:
                                fmt.Println(v19, v19, v13)
                            case 32768:
                                fmt.Println(v1)
                            }
                            _ = v19
                        }
                    case "":
                        for fi22 := 0; fi22 < 17; fi22++ {
                            if true {
                                v23 := GenS4{F0: 4294967295}
                                _ = v23
                            } else {
                                if true {
                                    v4[genFn0(uint8(v1), 41)] = 842
                                    if (!(!(!(false && false)))) {
                                        switch ("\n\t" + ((("" + "hello") + ("a" + "foo bar")) + "foo bar")) {
                                        case "a":
                                            _ = len(v4)
                                        default:
                                            if (!true) {
                                                var v24 uint = 65535
                                                _ = v24
                                            }
                                        }
                                        v25 := []int{}
                                        var v26 float32 = -2.25
                                        _ = v25
                                        _ = v26
                                    } else {
                                        switch (fi22 << ((255) & 63)) {
                                        case -2147483648:
                                            if (true || true) {
                                                v27 := GenS1{F0: fi12, F1: 18446744073709551615}
                                                if (65535 > 42) {
                                                    v28 := GenS5{F0: (v13 / (((v13 << ((57) & 63))) | 1)), F1: 32767, F2: v13}
                                                    var v29 uint32 = 664
                                                    _ = v28
                                                    _ = v29
                                                } else {
                                                    switch ("" + "") {
                                                    case "\n\t":
                                                        switch (("" + "") + "foo bar") {
                                                        case "a":
                                                            var v30 uint = 0
                                                            _ = v30
                                                        default:
                                                            v31 := []int{(fi12 ^ (fi12 << ((4294967295) & 63)))}
                                                            var v32 float64 = -2.25
                                                            v33 := GenS3{F0: v1, F1: 32767, F2: false}
                                                            _ = v31
                                                            _ = v32
                                                            _ = v33
                                                        }
                                                    }
                                                }
                                                var v34 int = 19
                                                _ = v27
                                                _ = v34
                                            } else {
                                                v35 := GenS6{F0: v13, F1: 32767}
                                                if true {
                                                    v1 = (v1 & 18446744073709551615)
                                                    v1 -= v1
                                                    v36 := []int{-417, (^(fi12)), fi22, (fi22 + 127)}
                                                    _ = v36
                                                }
                                                _ = v35
                                            }
                                        case 32767:
                                            var v37 int = (fi12 * fi22)
                                            _ = v37
                                        default:
                                            var v38 float64 = (float64((1e10 * (0.0 / 1.5))) - ((1e10 / 1e10) + 1.5))
                                            _ = v38
                                        }
                                        v39 := []int{}
                                        fmt.Println(v1, fi12, fi22)
                                        _ = v39
                                    }
                                    if false {
                                        var v40 int32 = 1
                                        if true {
                                            switch fi22 {
                                            case -199:
                                                v13 -= (v13 << ((349) & 63))
                                            case 173:
                                                v4[65535] = 4294967295
                                            case 422:
                                                v41 := GenS3{F0: v1, F1: 2, F2: true}
                                                _ = v41
                                            default:
                                                var v42 int64 = (v3 ^ 182)
                                                v40 = 1
                                                if true {
                                                    v1 = (v1 + (v1 >> ((18446744073709551615) & 63)))
                                                    _ = len(v4)
                                                }
                                                _ = v42
                                            }
                                            switch (fi12 + fi12) {
                                            case -1:
                                                v43 := GenS3{F0: (v1 << ((255) & 63)), F1: 255, F2: true}
                                                _ = v43
                                            }
                                            v40 += int32(v13)
                                        }
                                        switch (fi22 % ((9223372036854775807) | 1)) {
                                        case -2147483648:
                                            v44 := GenS7{F0: "hello", F1: ("foo bar" + "foo bar"), F2: (v1 >> ((540) & 63))}
                                            _ = v44
                                        case 215:
                                            v45 := GenS4{F0: 916}
                                            _ = v45
                                        case 1:
                                            var v46 bool = (-4 != 347)
                                            _ = v46
                                        default:
                                            v1 += (v1 ^ 18446744073709551615)
                                            v47 := GenS8{F0: (fi12 >> ((255) & 63))}
                                            _ = v47
                                        }
                                        _ = v40
                                    } else {
                                        var v48 uint64 = (v1 / (((v1 + (v1 + (^(v1))))) | 1))
                                        v49 := GenS9{F0: (^(v3))}
                                        _ = v48
                                        _ = v49
                                    }
                                    v50 := GenS10{F0: true, F1: 0}
                                    _ = v50
                                }
                                var v51 float32 = 0.0
                                v51 += ((((v51 - 0.1) + 1e10) / 1.0) - float32(1.5))
                                _ = v51
                            }
                            _ = fi22
                        }
                    default:
                        if true {
                            var v52 bool = (!true)
                            for fi53 := 0; fi53 < 10; fi53++ {
                                v54 := GenS11{F0: 1, F1: 127}
                                switch fi53 {
                                case 32767:
                                    if false {
                                        v1 -= (v1 + (v1 >> ((2) & 63)))
                                        v1 -= 494
                                        _ = len(v4)
                                    }
                                case 1:
                                    v55 := v2["foo bar"]
                                    _ = v55
                                case -2147483648:
                                    var v56 uint8 = 90
                                    _ = v56
                                default:
                                    v13 = v13
                                }
                                _ = fi53
                                _ = v54
                            }
                            var v57 float32 = (((((-2.25 / (-2.25 + float32((0.1 * (-2.25 / 1.5))))) + 1e10) - 0.0) / 0.1) / 1e10)
                            _ = v52
                            _ = v57
                        } else {
                            v1 = (^(v1))
                            var v58 string = "\n\t"
                            if (true && (false && (false && true))) {
                                for fi59 := 0; fi59 < 10; fi59++ {
                                    if true {
                                        fmt.Println(v13)
                                    }
                                    _ = fi59
                                }
                            } else {
                                v60 := GenS12{F0: (fi12 ^ (fi12 - 1)), F1: 42}
                                fmt.Println(v3, v58, v58)
                                if false {
                                    for fi61 := 0; fi61 < 8; fi61++ {
                                        var v62 int8 = v13
                                        v63 := []int{fi12}
                                        _ = fi61
                                        _ = v62
                                        _ = v63
                                    }
                                    for fi64 := 0; fi64 < 8; fi64++ {
                                        switch ((v58 + (v58 + v58)) + v58) {
                                        case "\n\t":
                                            v13 = v13
                                        case "":
                                            v3 = (v3 | v3)
                                        default:
                                            var v65 string = "hello"
                                            _ = v65
                                        }
                                        _ = fi64
                                    }
                                } else {
                                    fmt.Println(fi12, v13, v3)
                                    switch fi12 {
                                    case -1:
                                        v66 := GenS13{F0: 65535}
                                        _ = v66
                                    case -2147483648:
                                        if true {
                                            var v67 uint = 0
                                            _ = v67
                                        }
                                    case -704:
                                        var v68 float64 = (((1e10 / 0.1) * (float64((1e10 + 1.5)) + 1.5)) - 0.0)
                                        _ = v68
                                    default:
                                        var v69 float64 = float64(((0.1 - (0.1 / (float32((1.5 - float64((float32((3.14159 / 1e10)) * (1.5 / -2.25))))) + 1e10))) / (0.0 - ((0.0 - -2.25) + 0.0))))
                                        _ = v69
                                    }
                                }
                                _ = v60
                            }
                            _ = v58
                        }
                    }
                    var v70 uint16 = 752
                    _ = v70
                } else {
                    v13 = v13
                }
                if true {
                    for fi71 := 0; fi71 < 15; fi71++ {
                        v3 -= v3
                        if ((!false) || ((fi71 ^ fi71) >= 0)) {
                            v3 -= (v3 ^ (v3 >> ((2) & 63)))
                            switch "" {
                            case "\n\t":
                                if (fi12 >= (fi12 * (fi71 >> ((65535) & 63)))) {
                                    _ = len(v2)
                                    v72 := GenS7{F0: "hello", F1: "\n\t", F2: v1}
                                    _ = v72
                                }
                            case "a":
                                if true {
                                    v73 := []int{fi12, fi12, (fi12 >> ((4294967295) & 63))}
                                    v3 += (^(v3))
                                    v13 = int8(v1)
                                    _ = v73
                                }
                            }
                        } else {
                            if false {
                                v13 -= (v13 >> ((1) & 63))
                                switch "a" {
                                case "a":
                                    var v74 int64 = (v3 + v3)
                                    _ = v74
                                case "":
                                    if (false && false) {
                                        switch -2147483648 {
                                        case 876:
                                            v75 := GenS14{F0: v13, F1: 2}
                                            var v76 int32 = -2147483648
                                            _ = v75
                                            _ = v76
                                        case 772:
                                            v77 := GenS9{F0: (-(v3))}
                                            _ = v77
                                        case 128:
                                            var v78 bool = (false || true)
                                            _ = v78
                                        }
                                        v13 = int8(v1)
                                        var v79 uint32 = 1
                                        if false {
                                            var v80 string = "hello"
                                            if ((v80 + v80) <= v80) {
                                                switch v80 {
                                                case "hello":
                                                    v81 := GenS7{F0: v80, F1: v80, F2: v1}
                                                    _ = v81
                                                case "a":
                                                    if (42 < 1) {
                                                        v82 := GenS2{F0: 906, F1: (fi12 + (fi12 / ((fi71) | 1)))}
                                                        _ = v82
                                                    }
                                                    var v83 uint16 = 841
                                                    var v84 uint8 = 255
                                                    _ = v83
                                                    _ = v84
                                                case "":
                                                    var v85 string = ((v80 + "foo bar") + "")
                                                    _ = v85
                                                default:
                                                    v3 -= v3
                                                    fmt.Println(v79, v80, fi71)
                                                    v86 := GenS5{F0: -1, F1: 210, F2: (v13 - v13)}
                                                    _ = v86
                                                }
                                                if (1e10 >= -2.25) {
                                                    v1 = (v1 % (((v1 / ((v1) | 1))) | 1))
                                                    fmt.Println(v1)
                                                    v87 := GenS15{F0: genFn0(255, uint(v79)), F1: (v3 << ((1) & 63))}
                                                    _ = v87
                                                }
                                                v80 = (v80 + "")
                                            } else {
                                                var v88 float32 = float32(0.0)
                                                v89 := []int{(fi12 | fi71), fi71}
                                                _ = v88
                                                _ = v89
                                            }
                                            v2[v80] = ("a" + "\n\t")
                                            _ = v80
                                        }
                                        _ = v79
                                    }
                                default:
                                    if false {
                                        if false {
                                            var v90 float64 = (1.5 - float64(float32(-2.25)))
                                            v1 = (v1 << ((416) & 63))
                                            _ = v90
                                        } else {
                                            v3 = v3
                                        }
                                        v3 = v3
                                        v13 = (-(v13))
                                    } else {
                                        switch (fi12 % (((fi12 * fi12)) | 1)) {
                                        case 42:
                                            if ("\n\t" > "a") {
                                                v3 = (v3 / ((v3) | 1))
                                                var v91 uint = 18446744073709551615
                                                _ = v91
                                            }
                                        }
                                        var v92 uint = 754
                                        v92 += 2
                                        _ = v92
                                    }
                                }
                            }
                            fmt.Println(fi12, v1)
                        }
                        var v93 float32 = (((((3.14159 / 3.14159) + -2.25) - -2.25) + (((1.5 * float32(float64(0.0))) * (1e10 + -2.25)) * 0.0)) / -2.25)
                        var v94 int = 338
                        v2["\n\t"] = "a"
                        _ = fi71
                        _ = v93
                        _ = v94
                    }
                    for fi95 := 0; fi95 < 10; fi95++ {
                        v13 -= v13
                        switch ("" + "") {
                        case "":
                            var v96 int64 = (v3 ^ (v3 % (((^(v3))) | 1)))
                            _ = v96
                        case "\n\t":
                            if ((2147483647 != 713) && true) {
                                var v97 int = 0
                                v1 = 610
                                v13 = (v13 - v13)
                                if (1e10 < (1e10 + 0.0)) {
                                    fmt.Println(fi12)
                                    if (127 >= 127) {
                                        var v98 uint64 = v1
                                        _ = v98
                                    } else {
                                        var v99 int64 = v3
                                        v100 := []int{(fi12 | v97), fi95, fi95, 1}
                                        var v101 uint64 = v1
                                        _ = v100
                                        _ = v101
                                        _ = v99
                                    }
                                }
                                var v102 float32 = float32((((1.5 * ((1e10 + 0.0) + 1.5)) * -2.25) + 3.14159))
                                var v103 string = ("a" + "a")
                                _ = v102
                                _ = v103
                                _ = v97
                            } else {
                                if false {
                                    if ("hello" != "a") {
                                        v104 := GenS4{F0: 4294967295}
                                        v105 := []int{(fi12 + (fi12 & (^(fi12)))), (fi95 / ((9223372036854775807) | 1)), -2147483648}
                                        _ = v104
                                        _ = v105
                                    }
                                }
                                fmt.Println(v13, v1, v1)
                                v1 -= (v1 | (v1 * (v1 | (v1 >> ((1) & 63)))))
                            }
                        default:
                            var v106 float64 = (3.14159 * 3.14159)
                            _ = v106
                        }
                        var v107 bool = false
                        _ = fi95
                        _ = v107
                    }
                }
            }
            if false {
                for fi108 := 0; fi108 < 17; fi108++ {
                    fmt.Println(v1, fi12, v3)
                    _ = fi108
                }
                fmt.Println(v3)
            } else {
                fmt.Println(v1, v1, v13)
                v109 := []int{fi12, (fi12 + 1)}
                switch "foo bar" {
                case "":
                    var v110 int64 = (v3 - 935)
                    _ = v110
                case "a":
                    if (42 == 0) {
                        v1 -= (v1 >> ((65535) & 63))
                        var v111 int16 = 42
                        v1 = v1
                        _ = v111
                    } else {
                        for fi112 := 0; fi112 < 14; fi112++ {
                            if (1e10 <= float64(0.1)) {
                                v3 = (v3 / ((v3) | 1))
                                v13 -= 1
                                fmt.Println(v13, fi112, v1)
                            } else {
                                v113 := GenS16{F0: ("foo bar" + "hello"), F1: true}
                                v114 := GenS12{F0: 1000, F1: 948}
                                switch "" {
                                case "foo bar":
                                    if (128 > 2) {
                                        if (1 <= 4) {
                                            if true {
                                                delete(v2, "\n\t")
                                                v4[genFn0(uint8(v1), 911)] = 65535
                                                v115 := GenS13{F0: 18446744073709551615}
                                                if false {
                                                    fmt.Println(fi112, v3)
                                                }
                                                fmt.Println(v1)
                                                if false {
                                                    switch ("hello" + "\n\t") {
                                                    case "hello":
                                                        var v116 uint = 525
                                                        _ = v116
                                                    case "foo bar":
                                                        switch "a" {
                                                        case "\n\t":
                                                            v1 -= (v1 - v1)
                                                        case "foo bar":
                                                            v13 = 54
                                                        default:
                                                            if (0.0 != (float64(0.0) / (((float64(float32(((0.1 + ((3.14159 - 0.0) * ((0.1 + 1e10) + (float64(3.14159) - (float64(0.0) + ((3.14159 * (3.14159 / 1.5)) * (3.14159 / ((-2.25 + 1.5) + 1e10)))))))) + float64(((-2.25 / 1.0) / 1e10))))) * 3.14159) - (0.0 - 3.14159)) - (float64(0.0) - 1.5)))) {
                                                                fmt.Println(fi112, v13)
                                                                _ = len(v109)
                                                                v117 := GenS17{F0: 1, F1: 42}
                                                                _ = v117
                                                            } else {
                                                                if false {
                                                                    v1 = v1
                                                                    v118 := v4[0]
                                                                    var v119 bool = (v13 > (v13 + v13))
                                                                    v109 = append(v109, fi12)
                                                                    _ = v118
                                                                    _ = v119
                                                                } else {
                                                                    fmt.Println(v13)
                                                                    _ = len(v109)
                                                                    if false {
                                                                        switch 128 {
                                                                        case 2147483647:
                                                                            v13 -= v13
                                                                        case 128:
                                                                            if true {
                                                                                var v120 uint64 = 1
                                                                                v121 := GenS14{F0: v13, F1: 710}
                                                                                v120 += v120
                                                                                _ = v120
                                                                                _ = v121
                                                                            } else {
                                                                                v122 := GenS18{F0: 42, F1: (-(v3)), F2: "\n\t"}
                                                                                _ = v122
                                                                            }
                                                                        case 42:
                                                                            v2["\n\t"] = "a"
                                                                        default:
                                                                            _ = len(v109)
                                                                        }
                                                                        var v123 uint64 = (v1 & v1)
                                                                        v123 = (v123 & v1)
                                                                        fmt.Println(v3)
                                                                        v109 = append(v109, fi112)
                                                                        _ = v123
                                                                    } else {
                                                                        v1 = 18446744073709551615
                                                                        v1 += (v1 + 669)
                                                                    }
                                                                }
                                                            }
                                                        }
                                                    }
                                                } else {
                                                    v1 -= v1
                                                    v124 := GenS5{F0: v13, F1: 1, F2: -31}
                                                    _ = v124
                                                }
                                                v125 := GenS19{F0: (v1 | v1)}
                                                var v126 bool = true
                                                _ = v115
                                                _ = v125
                                                _ = v126
                                            }
                                            switch (fi12 | (-(fi12))) {
                                            case 42:
                                                var v127 uint64 = 527
                                                _ = v127
                                            case 852:
                                                v128 := v2["hello"]
                                                _ = v128
                                            }
                                            v1 = 150
                                        }
                                        var v129 float32 = 3.14159
                                        _ = v129
                                    }
                                case "\n\t":
                                    var v130 float64 = (1.5 / 1.0)
                                    _ = v130
                                default:
                                    v1 = v1
                                }
                                _ = v113
                                _ = v114
                            }
                            _ = fi112
                        }
                        for fi131 := 0; fi131 < 11; fi131++ {
                            var v132 int = (fi12 | fi131)
                            v132 = 1
                            _ = fi131
                            _ = v132
                        }
                    }
                case "\n\t":
                    if (0 == 267) {
                        v133 := v2[""]
                        var v134 uint64 = v1
                        for fi135 := 0; fi135 < 12; fi135++ {
                            if len(v109) > 0 {
                                v109[0] = (fi12 / ((fi135) | 1))
                            }
                            var v136 int = (fi135 & fi135)
                            _ = len(v109)
                            _ = fi135
                            _ = v136
                        }
                        _ = v133
                        _ = v134
                    } else {
                        v137 := GenS20{F0: 2, F1: true, F2: 121}
                        if false {
                            if false {
                                var v138 uint16 = 135
                                var v139 uint16 = (^(v138))
                                var v140 string = "foo bar"
                                _ = v138
                                _ = v139
                                _ = v140
                            } else {
                                if len(v109) > 0 {
                                    v109[0] = (fi12 | (-(fi12)))
                                }
                                for fi141 := 0; fi141 < 20; fi141++ {
                                    v1 -= 4294967295
                                    _ = fi141
                                }
                            }
                        } else {
                            fmt.Println(v1, fi12, v13)
                            var v142 bool = true
                            _ = v142
                        }
                        _ = v137
                    }
                default:
                    if true {
                        v143 := GenS15{F0: 2147483647, F1: v3}
                        _ = v143
                    }
                }
                _ = v109
            }
            _ = v15
        }
        var v144 uint16 = uint16(v1)
        _ = fi12
        _ = v13
        _ = v144
    }
    fmt.Println(v3, v1)
    v145 := GenS6{F0: -74, F1: -32768}
    _ = len(v4)
    v1 = 0
    if (((!false) && true) && (89 == 591)) {
        var v146 uint = 0
        delete(v2, "")
        _ = v146
    }
    v147 := []int{2147483647, int(v1)}
    var v148 int32 = 156
    v1 += (v1 * v1)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v147:", len(v147), v147)
    fmt.Println("v148:", v148)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v145
    _ = v147
    _ = v148
    _ = v2
    _ = v3
    _ = v4
}
