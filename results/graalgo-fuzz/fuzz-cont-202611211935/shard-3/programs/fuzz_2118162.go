package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int16; F1 int64; F2 uint }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint16; F1 uint32; F2 int }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 int8; F1 uint8 }
type GenS7 struct { F0 string; F1 int64; F2 int8 }
type GenS8 struct { F0 bool; F1 int }
type GenS9 struct { F0 int; F1 int }
type GenS10 struct { F0 int8; F1 uint8; F2 uint8 }
type GenS11 struct { F0 int8; F1 uint32; F2 uint }
type GenS12 struct { F0 int16; F1 uint; F2 uint16 }
type GenS13 struct { F0 int16; F1 int32; F2 uint32 }
type GenS14 struct { F0 uint }
type GenS15 struct { F0 int16; F1 string; F2 uint64 }
type GenS16 struct { F0 uint64; F1 uint64; F2 uint }
type GenS17 struct { F0 uint; F1 uint64 }
type GenS18 struct { F0 uint16; F1 int16; F2 uint32 }
type GenS19 struct { F0 uint8; F1 string }
type GenS20 struct { F0 string; F1 int32; F2 int }
type GenS21 struct { F0 uint8; F1 int; F2 int }
type GenS22 struct { F0 uint16; F1 uint }
type GenS23 struct { F0 uint8 }
type GenS24 struct { F0 int64; F1 int }
type GenS25 struct { F0 int64; F1 int }
func genFn0(p0 int, p1 int16) string {
    if (851 >= 1) {
        if true {
            var v6 bool = false
            p1 -= (p1 % ((p1) | 1))
            for fi7 := 0; fi7 < 7; fi7++ {
                for fi8 := 0; fi8 < 5; fi8++ {
                    var v9 int16 = p1
                    v10 := []int{(p0 % ((32768) | 1)), fi8, (p0 / ((1) | 1))}
                    v6 = (fi7 <= (p0 % ((186) | 1)))
                    var v11 uint = 0
                    _ = fi8
                    _ = v10
                    _ = v11
                    _ = v9
                }
                for fi12 := 0; fi12 < 18; fi12++ {
                    v13 := []int{(p0 ^ -1), p0, (p0 - (fi12 >> ((165) & 63))), (fi12 >> ((18446744073709551615) & 63))}
                    switch ((("foo bar" + (("foo bar" + "a") + "")) + "\n\t") + "foo bar") {
                    case "foo bar":
                        if (!(!(((0 > 65535) && v6) && (11 == 127)))) {
                            if len(v13) > 0 {
                                v13[0] = -226
                            }
                            fmt.Println(fi12, p0, p0)
                            v14 := GenS1{F0: -195}
                            p0 = -316
                            v15 := GenS1{F0: -1}
                            _ = v14
                            _ = v15
                        }
                    }
                    v16 := GenS2{F0: p1, F1: 42, F2: 1}
                    _ = fi12
                    _ = v13
                    _ = v16
                }
                for fi17 := 0; fi17 < 4; fi17++ {
                    v18 := GenS3{F0: 752}
                    var v19 int64 = -1
                    switch fi7 {
                    case 42:
                        if (!true) {
                            if false {
                                p0 -= fi17
                                var v20 int = -9223372036854775808
                                if (-2147483648 <= 1) {
                                    v21 := GenS4{F0: 457, F1: 1, F2: (p0 ^ (p0 * p0))}
                                    fmt.Println(v20, p0, v19)
                                    switch fi17 {
                                    case 0:
                                        v22 := GenS2{F0: p1, F1: (v19 >> ((89) & 63)), F2: 2}
                                        _ = v22
                                    case 4294967295:
                                        if v6 {
                                            var v23 int16 = (p1 & p1)
                                            if ((454 <= 1) || v6) {
                                                p1 = 32767
                                                switch "\n\t" {
                                                case "a":
                                                    if false {
                                                        if v6 {
                                                            v23 = (v23 * v23)
                                                            var v24 string = ("a" + "a")
                                                            _ = v24
                                                        } else {
                                                            switch "" {
                                                            case "":
                                                                v23 = -1
                                                            case "a":
                                                                v20 -= 32768
                                                            default:
                                                                if false {
                                                                    switch (^(fi17)) {
                                                                    case 0:
                                                                        v25 := []int{}
                                                                        _ = v25
                                                                    }
                                                                    var v26 float32 = 1.5
                                                                    _ = v26
                                                                } else {
                                                                    var v27 int8 = -109
                                                                    v20 = fi17
                                                                    v28 := GenS5{F0: 18446744073709551615}
                                                                    _ = v27
                                                                    _ = v28
                                                                }
                                                            }
                                                            switch (("foo bar" + "\n\t") + "foo bar") {
                                                            case "":
                                                                var v29 int = 128
                                                                _ = v29
                                                            case "foo bar":
                                                                switch (p0 / ((fi17) | 1)) {
                                                                case 408:
                                                                    var v30 float64 = 1e10
                                                                    _ = v30
                                                                case 1:
                                                                    fmt.Println(p0, v6, v19)
                                                                case 9223372036854775807:
                                                                    switch "a" {
                                                                    case "a":
                                                                        v31 := GenS6{F0: -47, F1: 0}
                                                                        _ = v31
                                                                    }
                                                                default:
                                                                    fmt.Println(v23, v19)
                                                                }
                                                            case "a":
                                                                v23 = v23
                                                            }
                                                            var v32 int16 = v23
                                                            switch (fi17 ^ 2147483647) {
                                                            case 42:
                                                                p0 = v20
                                                            case -203:
                                                                switch (fi7 >> ((446) & 63)) {
                                                                case 2147483647:
                                                                    v33 := []int{-457}
                                                                    _ = v33
                                                                }
                                                            default:
                                                                switch "" {
                                                                case "hello":
                                                                    var v34 uint8 = 13
                                                                    _ = v34
                                                                case "foo bar":
                                                                    v35 := []int{}
                                                                    _ = v35
                                                                default:
                                                                    v36 := GenS7{F0: "\n\t", F1: 42, F2: -128}
                                                                    _ = v36
                                                                }
                                                            }
                                                            _ = v32
                                                        }
                                                        v37 := GenS1{F0: v19}
                                                        v38 := GenS3{F0: -9223372036854775808}
                                                        _ = v37
                                                        _ = v38
                                                    } else {
                                                        v19 -= (v19 / ((v19) | 1))
                                                        v39 := GenS7{F0: "foo bar", F1: (v19 * 1), F2: 0}
                                                        var v40 string = "\n\t"
                                                        if (1 > 18446744073709551615) {
                                                            v41 := GenS8{F0: (1.5 != 1.5), F1: (fi7 << ((42) & 63))}
                                                            _ = v41
                                                        }
                                                        v42 := GenS1{F0: v19}
                                                        v20 = (p0 >> ((42) & 63))
                                                        _ = v39
                                                        _ = v40
                                                        _ = v42
                                                    }
                                                case "":
                                                    if v6 {
                                                        var v43 int = 42
                                                        v44 := []int{p0, fi17}
                                                        v45 := GenS9{F0: v20, F1: v20}
                                                        _ = v43
                                                        _ = v44
                                                        _ = v45
                                                    }
                                                default:
                                                    fmt.Println(v19, v6)
                                                }
                                            }
                                            _ = v23
                                        } else {
                                            var v46 uint32 = 42
                                            p0 = (fi7 * (v20 << ((18446744073709551615) & 63)))
                                            v19 += 1
                                            _ = v46
                                        }
                                    case 625:
                                        v19 += (v19 ^ v19)
                                    default:
                                        var v47 uint16 = 0
                                        _ = v47
                                    }
                                    v48 := GenS9{F0: fi17, F1: -2147483648}
                                    var v49 uint8 = 255
                                    switch "\n\t" {
                                    case "foo bar":
                                        switch "foo bar" {
                                        case "hello":
                                            v50 := []int{-276, (-(v20))}
                                            _ = v50
                                        case "":
                                            v51 := []int{}
                                            _ = v51
                                        }
                                    case "hello":
                                        var v52 float64 = -2.25
                                        _ = v52
                                    case "\n\t":
                                        if false {
                                            fmt.Println(p0, v49, v20)
                                            v53 := GenS10{F0: -1, F1: (v49 >> ((1) & 63)), F2: (v49 | v49)}
                                            if (!false) {
                                                fmt.Println(p0, v19, p0)
                                                v49 -= (v49 % ((1) | 1))
                                                p1 = p1
                                            } else {
                                                v20 = int(p1)
                                                v54 := GenS1{F0: -1}
                                                _ = v54
                                            }
                                            var v55 bool = (v6 && true)
                                            _ = v53
                                            _ = v55
                                        } else {
                                            switch (fi17 * (p0 ^ v20)) {
                                            case -495:
                                                v56 := GenS11{F0: 123, F1: 42, F2: 1}
                                                _ = v56
                                            case -863:
                                                var v57 uint16 = 32767
                                                _ = v57
                                            default:
                                                v58 := GenS5{F0: 4294967295}
                                                _ = v58
                                            }
                                            v59 := []int{fi17, 127, int(v49), (v20 / (((-(fi17))) | 1))}
                                            if len(v59) >= 2 {
                                                v59 = v59[1:len(v59)]
                                            }
                                            p0 = (^(p0))
                                            _ = v59
                                        }
                                    default:
                                        if ("" >= "a") {
                                            var v60 int64 = v19
                                            var v61 bool = v6
                                            _ = v60
                                            _ = v61
                                        }
                                    }
                                    _ = v21
                                    _ = v48
                                    _ = v49
                                } else {
                                    fmt.Println(v19, p0, fi7)
                                    switch (p0 | v20) {
                                    case -9223372036854775808:
                                        if true {
                                            v62 := []int{fi7, (fi7 >> ((1) & 63)), (-(p0))}
                                            p0 += (fi7 / ((p0) | 1))
                                            _ = v62
                                        }
                                    }
                                }
                                _ = v20
                            }
                            switch "\n\t" {
                            case "\n\t":
                                v63 := GenS12{F0: p1, F1: 1, F2: 42}
                                _ = v63
                            default:
                                v64 := []int{fi7}
                                _ = v64
                            }
                            v65 := []int{fi7, (fi17 * fi17)}
                            _ = v65
                        }
                    case 625:
                        v66 := []int{-2147483648}
                        _ = v66
                    default:
                        v67 := GenS10{F0: 46, F1: 128, F2: 168}
                        _ = v67
                    }
                    v68 := []int{p0, p0, (fi7 >> ((4294967295) & 63)), fi7}
                    v69 := GenS13{F0: p1, F1: 159, F2: 814}
                    v70 := GenS14{F0: 255}
                    p0 -= fi7
                    _ = fi17
                    _ = v18
                    _ = v19
                    _ = v68
                    _ = v69
                    _ = v70
                }
                _ = fi7
            }
            _ = v6
        }
        var v71 uint16 = 255
        if true {
            v72 := []int{}
            if false {
                var v73 string = ("\n\t" + "hello")
                _ = v73
            } else {
                v74 := GenS2{F0: p1, F1: 9223372036854775807, F2: 255}
                for fi75 := 0; fi75 < 16; fi75++ {
                    if false {
                        for fi76 := 0; fi76 < 6; fi76++ {
                            var v77 uint16 = v71
                            switch "" {
                            case "hello":
                                p0 += fi75
                            case "":
                                if len(v72) > 0 {
                                    v72[0] = (fi76 ^ -829)
                                }
                            }
                            _ = fi76
                            _ = v77
                        }
                    } else {
                        var v78 string = ""
                        switch fi75 {
                        case -1:
                            v79 := GenS15{F0: -648, F1: v78, F2: 309}
                            _ = v79
                        case 9223372036854775807:
                            _ = len(v72)
                        case 991:
                            for fi80 := 0; fi80 < 7; fi80++ {
                                fmt.Println(fi80)
                                if (!false) {
                                    var v81 int16 = 0
                                    if false {
                                        v72 = append(v72, 169)
                                        if len(v72) > 0 {
                                            v72[0] = (-(fi80))
                                        }
                                    }
                                    if (((3.14159 * (1.5 / 0.1)) != float64(0.1)) && false) {
                                        v82 := GenS6{F0: -128, F1: 127}
                                        _ = v82
                                    } else {
                                        var v83 uint32 = 2
                                        _ = v83
                                    }
                                    _ = v81
                                }
                                _ = fi80
                            }
                        }
                        _ = v78
                    }
                    if len(v72) >= 2 {
                        v72 = v72[1:len(v72)]
                    }
                    fmt.Println(v71, v71)
                    _ = fi75
                }
                v84 := GenS16{F0: 406, F1: 18446744073709551615, F2: uint(v71)}
                _ = v74
                _ = v84
            }
            _ = v72
        } else {
            v85 := GenS11{F0: 1, F1: 865, F2: 855}
            _ = v85
        }
        _ = v71
    } else {
        for fi86 := 0; fi86 < 4; fi86++ {
            var v87 int16 = p1
            var v88 uint8 = 128
            v87 += p1
            _ = fi86
            _ = v87
            _ = v88
        }
        p1 -= 1000
        var v89 int = int(p1)
        _ = v89
    }
    fmt.Println(p1, p1)
    for fi90 := 0; fi90 < 8; fi90++ {
        var v91 uint = 2
        v92 := GenS3{F0: p0}
        for fi93 := 0; fi93 < 20; fi93++ {
            if true {
                v94 := []int{fi90}
                fmt.Println(fi90, fi90, p1)
                if false {
                    var v95 int8 = 0
                    var v96 bool = (!true)
                    v97 := GenS1{F0: 0}
                    _ = v95
                    _ = v96
                    _ = v97
                } else {
                    switch fi93 {
                    case 922:
                        if len(v94) > 0 {
                            v94[0] = 2147483647
                        }
                    case 914:
                        if ((!(42 == 2)) || true) {
                            var v98 int16 = (p1 | (p1 % ((p1) | 1)))
                            v99 := GenS17{F0: (v91 ^ v91), F1: 42}
                            _ = v98
                            _ = v99
                        }
                    }
                    v100 := GenS18{F0: 2, F1: (p1 >> ((v91) & 63)), F2: 1}
                    v91 += v91
                    _ = v100
                }
                _ = v94
            } else {
                p0 = 1
                if (fi90 <= (fi93 - 712)) {
                    v101 := GenS17{F0: (v91 >> ((v91) & 63)), F1: 2}
                    if (2147483647 > 65535) {
                        p0 -= fi90
                    } else {
                        v102 := []int{fi93, fi93, fi90, (fi93 << ((v91) & 63))}
                        if len(v102) >= 2 {
                            v102 = v102[1:len(v102)]
                        }
                        v103 := GenS6{F0: 100, F1: 255}
                        _ = v102
                        _ = v103
                    }
                    _ = v101
                }
            }
            if (2 < 127) {
                p0 = (fi90 / ((127) | 1))
            } else {
                var v104 float32 = 1.5
                if true {
                    var v105 uint8 = 67
                    var v106 uint8 = (v105 / ((v105) | 1))
                    p0 -= (-(fi90))
                    _ = v105
                    _ = v106
                }
                _ = v104
            }
            var v107 int64 = 0
            v108 := GenS19{F0: 127, F1: "hello"}
            _ = fi93
            _ = v107
            _ = v108
        }
        _ = fi90
        _ = v91
        _ = v92
    }
    switch (p0 + p0) {
    case 127:
        for fi109 := 0; fi109 < 20; fi109++ {
            v110 := GenS20{F0: ("\n\t" + "a"), F1: 42, F2: p0}
            p0 -= -1
            for fi111 := 0; fi111 < 5; fi111++ {
                var v112 int8 = 42
                v112 -= -14
                fmt.Println(p1)
                v113 := GenS21{F0: 128, F1: (^(fi109)), F2: (fi111 + (-(p0)))}
                v114 := []int{(-(fi111)), 42, (fi109 * fi109), p0}
                switch "\n\t" {
                case "":
                    switch -999 {
                    case 42:
                        if false {
                            v112 = (v112 / (((-(v112))) | 1))
                            switch 9223372036854775807 {
                            case 42:
                                var v115 bool = true
                                _ = v115
                            default:
                                fmt.Println(fi111, fi111, p1)
                            }
                        } else {
                            p1 -= (p1 / ((p1) | 1))
                        }
                    }
                }
                _ = fi111
                _ = v112
                _ = v113
                _ = v114
            }
            _ = fi109
            _ = v110
        }
    default:
        v116 := GenS22{F0: 22, F1: 2}
        _ = v116
    }
    _ = p0
    _ = p1
    return ("" + "\n\t")
}
func genFn1(p0 string, p1 int8, p2 uint64) bool {
    v117 := GenS5{F0: 255}
    v118 := []int{42}
    for fi119 := 0; fi119 < 14; fi119++ {
        switch "\n\t" {
        case "a":
            fmt.Println(p1)
        case "\n\t":
            p2 = 2
        }
        v120 := GenS2{F0: -32768, F1: -1, F2: 42}
        for fi121 := 0; fi121 < 3; fi121++ {
            var v122 float32 = (3.14159 * 1.5)
            p0 = ""
            _ = fi121
            _ = v122
        }
        _ = fi119
        _ = v120
    }
    v123 := GenS23{F0: 0}
    _ = p0
    _ = p1
    _ = p2
    _ = v117
    _ = v118
    _ = v123
    return (true || true)
}

func main() {
    var v1 uint8 = 96
    var v2 int16 = 1000
    var v3 uint = 255
    v4 := []int{854, -1, -2147483648, 1}
    v5 := []string{"a", "hello"}
    v1 -= 2
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v124 int16 = (v2 << ((268) & 63))
    var v125 float32 = 1.5
    v126 := GenS24{F0: 890, F1: 1}
    v127 := GenS25{F0: 259, F1: 128}
    var v128 int = -9223372036854775808
    var v129 uint64 = 18446744073709551615
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v130 int32 = -2147483648
    fmt.Println("v1:", v1)
    fmt.Println("v124:", v124)
    fmt.Println("v125:", v125)
    fmt.Println("v128:", v128)
    fmt.Println("v129:", v129)
    fmt.Println("v130:", v130)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v124
    _ = v125
    _ = v126
    _ = v127
    _ = v128
    _ = v129
    _ = v130
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
