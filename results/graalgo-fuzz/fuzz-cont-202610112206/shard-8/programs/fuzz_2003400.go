package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32 }
type GenS2 struct { F0 int8; F1 int }
type GenS3 struct { F0 string; F1 uint64; F2 uint }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 string; F1 int32; F2 int32 }
type GenS6 struct { F0 int64; F1 uint32; F2 string }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 int8; F1 uint32; F2 uint8 }
type GenS9 struct { F0 string; F1 int; F2 uint }
type GenS10 struct { F0 int16; F1 uint64; F2 int64 }
type GenS11 struct { F0 string; F1 string }
type GenS12 struct { F0 int8; F1 int8 }
type GenS13 struct { F0 int16 }
type GenS14 struct { F0 uint8; F1 int64 }
type GenS15 struct { F0 int8 }
type GenS16 struct { F0 string }
type GenS17 struct { F0 int16; F1 uint; F2 int }
type GenS18 struct { F0 uint8; F1 string }
func genFn0() int8 {
    var v5 int64 = -1
    if (1 < 42) {
        v5 += 42
    }
    var v6 int64 = -9223372036854775808
    _ = v5
    _ = v6
    return 2
}
func genFn1() uint32 {
    v7 := []int{9223372036854775807, 127, 32767, 32767}
    var v8 float64 = 0.0
    switch "a" {
    case "hello":
        for fi9 := 0; fi9 < 8; fi9++ {
            if (57 == 0) {
                fmt.Println(v8)
            } else {
                for fi10 := 0; fi10 < 14; fi10++ {
                    if len(v7) >= 2 {
                        v7 = v7[1:len(v7)]
                    }
                    v11 := GenS1{F0: 42, F1: 26}
                    if false {
                        var v12 uint8 = 42
                        _ = v12
                    }
                    _ = fi10
                    _ = v11
                }
            }
            if len(v7) > 0 {
                v7[0] = (-(fi9))
            }
            _ = fi9
        }
    default:
        for fi13 := 0; fi13 < 17; fi13++ {
            v14 := GenS1{F0: genFn0(), F1: 952}
            v15 := GenS2{F0: genFn0(), F1: 1}
            _ = fi13
            _ = v14
            _ = v15
        }
    }
    if (-637 == 9223372036854775807) {
        v16 := GenS2{F0: genFn0(), F1: 4294967295}
        _ = v16
    } else {
        var v17 uint16 = 255
        _ = v17
    }
    _ = v7
    _ = v8
    return 711
}

func main() {
    var v1 float32 = 0.1
    var v2 string = ""
    var v3 uint32 = 0
    var v4 int8 = 96
    fmt.Println(v1, v2, v3)
    var v18 float32 = (0.0 + (v1 / (float32(1.5) * v1)))
    v19 := GenS2{F0: genFn0(), F1: 0}
    if false {
        var v20 uint64 = 18446744073709551615
        if true {
            if (!(uint64(v3) >= uint64(v3))) {
                var v21 uint32 = v3
                _ = v21
            } else {
                switch v2 {
                case "":
                    v3 = (v3 & uint32(v20))
                case "hello":
                    var v22 uint64 = v20
                    _ = v22
                default:
                    v23 := GenS1{F0: genFn0(), F1: genFn1()}
                    _ = v23
                }
                for fi24 := 0; fi24 < 20; fi24++ {
                    if true {
                        v4 = 127
                        if ((((1.5 / (3.14159 + 0.1)) - (-2.25 * ((1.5 - (1.5 - 0.1)) * (((0.0 / 0.1) + 0.0) - 0.1)))) - -2.25) < ((-2.25 - 1.5) + 1e10)) {
                            v1 = (-2.25 - v18)
                            if (((genFn1() < (^(v3))) || (0 == 255)) && (!(!false))) {
                                var v25 string = v2
                                _ = v25
                            }
                        } else {
                            var v26 bool = false
                            var v27 bool = true
                            var v28 uint16 = 0
                            _ = v26
                            _ = v27
                            _ = v28
                        }
                    }
                    v29 := []int{}
                    switch fi24 {
                    case 0:
                        v2 = v2
                    case -1:
                        v1 = v1
                    }
                    _ = fi24
                    _ = v29
                }
            }
            switch -795 {
            case -724:
                var v30 int8 = (v4 + v4)
                _ = v30
            case 611:
                for fi31 := 0; fi31 < 3; fi31++ {
                    for fi32 := 0; fi32 < 16; fi32++ {
                        if false {
                            switch 9223372036854775807 {
                            case -1:
                                if (!(fi31 <= (fi32 * 118))) {
                                    v33 := GenS2{F0: v4, F1: 949}
                                    v3 -= v3
                                    v34 := []int{(fi31 | (fi31 & (fi32 ^ fi31))), (fi31 << ((448) & 63)), (fi32 + (fi31 % ((fi31) | 1)))}
                                    switch v2 {
                                    case "\n\t":
                                        var v35 int = 0
                                        _ = v35
                                    case "":
                                        switch fi32 {
                                        case -16:
                                            v18 += 0.1
                                        default:
                                            v34 = append(v34, (fi32 - fi32))
                                        }
                                    default:
                                        if false {
                                            if false {
                                                v36 := GenS3{F0: ((v2 + (v2 + (v2 + "\n\t"))) + v2), F1: (v20 / (((v20 & uint64(v3))) | 1)), F2: 0}
                                                v34 = append(v34, 32768)
                                                var v37 string = ((((v2 + v2) + v2) + v2) + "hello")
                                                fmt.Println(fi31)
                                                if false {
                                                    fmt.Println(fi32, fi32)
                                                    fmt.Println(v4, fi32, fi31)
                                                    switch fi31 {
                                                    case 32767:
                                                        var v38 string = "foo bar"
                                                        _ = v38
                                                    case 220:
                                                        var v39 int32 = 2147483647
                                                        _ = v39
                                                    case 127:
                                                        _ = len(v34)
                                                    }
                                                }
                                                _ = v36
                                                _ = v37
                                            }
                                            switch (fi31 + (fi32 | (fi32 - (^(fi32))))) {
                                            case 128:
                                                if len(v34) > 0 {
                                                    v34[0] = fi31
                                                }
                                            }
                                        } else {
                                            if false {
                                                if len(v34) > 0 {
                                                    v34[0] = fi31
                                                }
                                                switch 4294967295 {
                                                case 9223372036854775807:
                                                    fmt.Println(v20, v3)
                                                case -9223372036854775808:
                                                    switch (fi31 & fi32) {
                                                    case -9223372036854775808:
                                                        if len(v34) >= 2 {
                                                            v34 = v34[1:len(v34)]
                                                        }
                                                        if (!false) {
                                                            fmt.Println(v20, v2)
                                                            v40 := GenS4{F0: 42}
                                                            var v41 uint64 = (v20 * 720)
                                                            _ = v40
                                                            _ = v41
                                                        } else {
                                                            var v42 int64 = 9223372036854775807
                                                            switch fi31 {
                                                            case 2147483647:
                                                                var v43 float32 = (float32(0.1) - 3.14159)
                                                                _ = v43
                                                            case -795:
                                                                if (v4 > genFn0()) {
                                                                    v20 = 643
                                                                } else {
                                                                    var v44 uint32 = 58
                                                                    var v45 int64 = (v42 - 42)
                                                                    _ = v44
                                                                    _ = v45
                                                                }
                                                            case 128:
                                                                switch 0 {
                                                                case -296:
                                                                    if ((!true) || (3.14159 != (0.0 - 3.14159))) {
                                                                        if ((v3 != v3) || false) {
                                                                            v3 = (v3 * genFn1())
                                                                            v18 = float32(3.14159)
                                                                            v46 := GenS5{F0: "a", F1: -449, F2: -1}
                                                                            _ = v46
                                                                        } else {
                                                                            var v47 float64 = 0.0
                                                                            v3 = (v3 % (((v3 - genFn1())) | 1))
                                                                            _ = v47
                                                                        }
                                                                        switch v2 {
                                                                        case "a":
                                                                            v34 = append(v34, (^(fi32)))
                                                                        case "\n\t":
                                                                            fmt.Println(v3)
                                                                        default:
                                                                            var v48 int64 = v42
                                                                            _ = v48
                                                                        }
                                                                    } else {
                                                                        v49 := GenS6{F0: v42, F1: (v3 >> ((2) & 63)), F2: v2}
                                                                        if false {
                                                                            var v50 int = fi32
                                                                            var v51 float32 = v1
                                                                            _ = v50
                                                                            _ = v51
                                                                        }
                                                                        _ = v49
                                                                    }
                                                                    switch v2 {
                                                                    case "foo bar":
                                                                        v52 := GenS1{F0: (v4 | -1), F1: v3}
                                                                        _ = v52
                                                                    case "a":
                                                                        fmt.Println(v20, fi31)
                                                                    default:
                                                                        switch (v2 + "hello") {
                                                                        case "\n\t":
                                                                            v18 += v18
                                                                        case "a":
                                                                            var v53 string = ("\n\t" + ("foo bar" + (("\n\t" + "hello") + v2)))
                                                                            _ = v53
                                                                        }
                                                                    }
                                                                    fmt.Println(fi31, fi32, fi31)
                                                                case -850:
                                                                    fmt.Println(v42, v18)
                                                                case -553:
                                                                    v54 := GenS7{F0: 1}
                                                                    _ = v54
                                                                default:
                                                                    fmt.Println(v42, fi32, v18)
                                                                    switch v2 {
                                                                    case "a":
                                                                        v2 = "hello"
                                                                    case "\n\t":
                                                                        v34 = append(v34, (fi32 - (-(fi31))))
                                                                    case "hello":
                                                                        if (!true) {
                                                                            v20 = (v20 ^ (v20 >> ((65535) & 63)))
                                                                        } else {
                                                                            v55 := GenS8{F0: 127, F1: (^(v3)), F2: 128}
                                                                            _ = v55
                                                                        }
                                                                    default:
                                                                        v3 += genFn1()
                                                                    }
                                                                }
                                                            }
                                                            switch fi31 {
                                                            case 32768:
                                                                switch fi32 {
                                                                case 564:
                                                                    v2 = "a"
                                                                case -9223372036854775808:
                                                                    var v56 uint = 534
                                                                    _ = v56
                                                                case 0:
                                                                    v34 = append(v34, (fi32 >> ((65535) & 63)))
                                                                }
                                                            case 0:
                                                                fmt.Println(v20, v20, v20)
                                                            case 670:
                                                                v20 -= v20
                                                            default:
                                                                switch (fi31 & (fi32 % (((fi31 * -938)) | 1))) {
                                                                case 444:
                                                                    v34 = append(v34, (fi32 + (fi32 + fi31)))
                                                                default:
                                                                    if (0 == 4294967295) {
                                                                        var v57 uint8 = 128
                                                                        v58 := GenS9{F0: "hello", F1: fi31, F2: 65535}
                                                                        _ = v57
                                                                        _ = v58
                                                                    }
                                                                }
                                                            }
                                                            _ = v42
                                                        }
                                                    }
                                                case 42:
                                                    fmt.Println(fi31)
                                                }
                                            }
                                        }
                                    }
                                    _ = v33
                                    _ = v34
                                } else {
                                    var v59 uint = 823
                                    var v60 int8 = 100
                                    if (!false) {
                                        var v61 uint16 = 65535
                                        switch v2 {
                                        case "a":
                                            fmt.Println(v60, v60, v4)
                                        case "\n\t":
                                            if false {
                                                var v62 float64 = ((0.0 / -2.25) + ((1e10 * float64((-2.25 + 3.14159))) + -2.25))
                                                v20 = v20
                                                switch "a" {
                                                case "\n\t":
                                                    fmt.Println(fi32, v62)
                                                case "foo bar":
                                                    var v63 uint16 = v61
                                                    _ = v63
                                                }
                                                if (false && false) {
                                                    if (((v3 & v3) != v3) || false) {
                                                        var v64 int8 = v4
                                                        v2 = (v2 + v2)
                                                        var v65 int8 = -79
                                                        _ = v64
                                                        _ = v65
                                                    } else {
                                                        switch fi31 {
                                                        case 128:
                                                            switch "hello" {
                                                            case "foo bar":
                                                                var v66 string = (v2 + "\n\t")
                                                                _ = v66
                                                            }
                                                        case -210:
                                                            v67 := GenS10{F0: 32767, F1: uint64(v59), F2: -1}
                                                            _ = v67
                                                        case -9223372036854775808:
                                                            switch (v2 + (v2 + (v2 + ""))) {
                                                            case "foo bar":
                                                                var v68 uint = (v59 - 4294967295)
                                                                _ = v68
                                                            case "":
                                                                v69 := GenS11{F0: ((("\n\t" + ("foo bar" + "\n\t")) + ("foo bar" + (("foo bar" + ("foo bar" + "a")) + "hello"))) + ((v2 + v2) + "hello")), F1: v2}
                                                                _ = v69
                                                            case "a":
                                                                var v70 bool = false
                                                                _ = v70
                                                            }
                                                        }
                                                        var v71 uint16 = v61
                                                        switch (fi32 & fi31) {
                                                        case 1:
                                                            var v72 uint64 = 671
                                                            _ = v72
                                                        case 463:
                                                            var v73 string = (v2 + (v2 + ("foo bar" + ("\n\t" + "foo bar"))))
                                                            switch fi32 {
                                                            case -102:
                                                                var v74 int = (fi31 / (((fi32 ^ (fi32 * -9223372036854775808))) | 1))
                                                                if true {
                                                                    var v75 int64 = 1
                                                                    v60 = v4
                                                                    _ = v75
                                                                }
                                                                v76 := []int{}
                                                                _ = v74
                                                                _ = v76
                                                            case 4294967295:
                                                                v77 := []int{32768, fi31, 1, fi31}
                                                                _ = v77
                                                            case 0:
                                                                fmt.Println(v1)
                                                            default:
                                                                v73 = v73
                                                            }
                                                            switch "" {
                                                            case "foo bar":
                                                                if false {
                                                                    v60 -= genFn0()
                                                                    v2 = v73
                                                                } else {
                                                                    v78 := []int{fi32}
                                                                    _ = v78
                                                                }
                                                            default:
                                                                var v79 uint = v59
                                                                _ = v79
                                                            }
                                                            v80 := []int{(fi32 % ((-762) | 1)), fi31, -2147483648}
                                                            v61 += uint16(v3)
                                                            _ = v73
                                                            _ = v80
                                                        default:
                                                            v81 := GenS4{F0: (v59 | (^(v59)))}
                                                            _ = v81
                                                        }
                                                        _ = v71
                                                    }
                                                    v82 := []int{(-(fi31))}
                                                    _ = v82
                                                }
                                                _ = v62
                                            }
                                        }
                                        if false {
                                            v83 := GenS4{F0: (^(v59))}
                                            _ = v83
                                        } else {
                                            v84 := GenS12{F0: -128, F1: (v60 << ((v59) & 63))}
                                            if true {
                                                v4 = (v4 << ((v59) & 63))
                                                if false {
                                                    v60 = (v4 ^ (^(v60)))
                                                }
                                                var v85 uint64 = v20
                                                _ = v85
                                            }
                                            if (true || (!false)) {
                                                if (v18 <= float32(3.14159)) {
                                                    switch ("" + "a") {
                                                    case "hello":
                                                        switch v2 {
                                                        case "hello":
                                                            fmt.Println(v18, v1, v20)
                                                        case "foo bar":
                                                            if (42 == -9223372036854775808) {
                                                                v4 = (v4 | v4)
                                                                fmt.Println(v1, v60, v18)
                                                                v86 := []int{}
                                                                _ = v86
                                                            } else {
                                                                v3 = (^(v3))
                                                                if false {
                                                                    var v87 uint16 = 127
                                                                    var v88 uint = v59
                                                                    if true {
                                                                        v87 -= (v87 ^ (^(v61)))
                                                                    }
                                                                    _ = v87
                                                                    _ = v88
                                                                }
                                                            }
                                                        default:
                                                            var v89 int = fi32
                                                            _ = v89
                                                        }
                                                    case "\n\t":
                                                        v90 := GenS13{F0: 0}
                                                        _ = v90
                                                    default:
                                                        v91 := GenS12{F0: genFn0(), F1: genFn0()}
                                                        _ = v91
                                                    }
                                                    switch v2 {
                                                    case "":
                                                        var v92 string = v2
                                                        _ = v92
                                                    case "hello":
                                                        fmt.Println(v60, fi31)
                                                        v4 = 42
                                                    case "a":
                                                        switch 164 {
                                                        case 67:
                                                            v93 := GenS12{F0: v4, F1: (-(v60))}
                                                            _ = v93
                                                        case 9223372036854775807:
                                                            v4 = v4
                                                        case -2147483648:
                                                            v4 -= (v60 / ((0) | 1))
                                                            var v94 uint64 = v20
                                                            fmt.Println(v4, v2, v2)
                                                            var v95 int32 = -1
                                                            _ = v94
                                                            _ = v95
                                                        default:
                                                            v3 += v3
                                                        }
                                                    }
                                                }
                                                v96 := []int{(fi32 * -893)}
                                                _ = v96
                                            } else {
                                                switch (v2 + "") {
                                                case "hello":
                                                    v97 := GenS10{F0: 42, F1: (v20 / ((uint64(v3)) | 1)), F2: -265}
                                                    _ = v97
                                                case "a":
                                                    if (v61 >= (v61 | (^(v61)))) {
                                                        v98 := []int{(fi31 % ((int(v3)) | 1)), (fi32 - (fi31 % ((fi32) | 1))), (^(fi31)), 42}
                                                        _ = v98
                                                    }
                                                case "":
                                                    var v99 uint64 = v20
                                                    _ = v99
                                                default:
                                                    v4 = v60
                                                }
                                                var v100 float64 = 0.0
                                                _ = v100
                                            }
                                            _ = v84
                                        }
                                        _ = v61
                                    } else {
                                        switch (v2 + v2) {
                                        case "":
                                            var v101 int64 = 42
                                            _ = v101
                                        case "hello":
                                            v1 = (v18 - v18)
                                        }
                                        var v102 uint = (^(v59))
                                        switch "a" {
                                        case "hello":
                                            fmt.Println(fi31, v102, v18)
                                        }
                                        _ = v102
                                    }
                                    var v103 bool = (!((!(((v20 << (((v59 - (v59 - v59))) & 63)) <= (^(v20))) || true)) && (32767 != 1)))
                                    v104 := GenS14{F0: 128, F1: 0}
                                    switch ((("hello" + (((v2 + v2) + (v2 + v2)) + "hello")) + "\n\t") + (v2 + ("foo bar" + "\n\t"))) {
                                    case "hello":
                                        v105 := GenS1{F0: v4, F1: (^(v3))}
                                        _ = v105
                                    case "a":
                                        v2 = v2
                                    case "foo bar":
                                        if true {
                                            fmt.Println(v4, v103)
                                            v106 := []int{}
                                            var v107 float32 = v18
                                            v108 := GenS11{F0: v2, F1: ((v2 + v2) + (v2 + (v2 + v2)))}
                                            var v109 float64 = 1e10
                                            _ = v106
                                            _ = v107
                                            _ = v108
                                            _ = v109
                                        } else {
                                            fmt.Println(fi31)
                                        }
                                    default:
                                        var v110 bool = v103
                                        _ = v110
                                    }
                                    _ = v103
                                    _ = v104
                                    _ = v59
                                    _ = v60
                                }
                            case -7:
                                v111 := GenS13{F0: -503}
                                var v112 bool = false
                                _ = v111
                                _ = v112
                            case 636:
                                if false {
                                    var v113 int8 = -116
                                    _ = v113
                                } else {
                                    switch (fi32 + 933) {
                                    case 299:
                                        if (255 > 0) {
                                            if (!true) {
                                                v3 = genFn1()
                                            } else {
                                                var v114 uint8 = 127
                                                _ = v114
                                            }
                                            var v115 uint16 = 0
                                            _ = v115
                                        } else {
                                            v116 := []int{fi32, -151, (fi32 % ((fi31) | 1))}
                                            fmt.Println(v2, fi31)
                                            switch fi32 {
                                            case 9223372036854775807:
                                                v117 := GenS1{F0: v4, F1: genFn1()}
                                                v18 = v1
                                                _ = v117
                                            case 687:
                                                v18 = 3.14159
                                            case 32768:
                                                if len(v116) > 0 {
                                                    v116[0] = (^(fi31))
                                                }
                                                if len(v116) >= 2 {
                                                    v116 = v116[1:len(v116)]
                                                }
                                            }
                                            var v118 int32 = -2147483648
                                            _ = v116
                                            _ = v118
                                        }
                                    case -1:
                                        v18 += (v1 + float32((1.5 / 1.5)))
                                    case -156:
                                        switch v2 {
                                        case "":
                                            switch fi31 {
                                            case -9223372036854775808:
                                                v18 += (1e10 - 1e10)
                                            case -2147483648:
                                                if false {
                                                    switch "\n\t" {
                                                    case "\n\t":
                                                        if false {
                                                            fmt.Println(v18)
                                                            var v119 int16 = -1
                                                            _ = v119
                                                        } else {
                                                            v120 := GenS13{F0: -66}
                                                            v121 := GenS8{F0: genFn0(), F1: v3, F2: 78}
                                                            var v122 uint8 = 127
                                                            _ = v120
                                                            _ = v121
                                                            _ = v122
                                                        }
                                                    }
                                                } else {
                                                    var v123 float32 = 0.1
                                                    var v124 string = v2
                                                    var v125 int64 = -9223372036854775808
                                                    _ = v123
                                                    _ = v124
                                                    _ = v125
                                                }
                                                fmt.Println(v2)
                                            }
                                        case "\n\t":
                                            switch int(v20) {
                                            case 9223372036854775807:
                                                var v126 uint16 = 42
                                                _ = v126
                                            case 445:
                                                v1 = v18
                                            }
                                        }
                                    default:
                                        var v127 uint32 = (^(v3))
                                        _ = v127
                                    }
                                }
                            }
                            v4 = 0
                        }
                        v128 := GenS15{F0: v4}
                        switch fi32 {
                        case -988:
                            fmt.Println(fi32)
                        }
                        _ = fi32
                        _ = v128
                    }
                    _ = fi31
                }
            case 857:
                switch 503 {
                case -2147483648:
                    switch "hello" {
                    case "a":
                        var v129 int64 = int64(v3)
                        _ = v129
                    case "foo bar":
                        v3 = 957
                    }
                case 105:
                    switch ("a" + (v2 + v2)) {
                    case "foo bar":
                        fmt.Println(v1)
                    case "hello":
                        v130 := []int{-956, 123}
                        _ = v130
                    case "a":
                        for fi131 := 0; fi131 < 10; fi131++ {
                            v4 = -128
                            v132 := GenS15{F0: (v4 % ((v4) | 1))}
                            var v133 uint32 = v3
                            _ = fi131
                            _ = v132
                            _ = v133
                        }
                    }
                }
            }
            if false {
                var v134 int32 = -487
                _ = v134
            } else {
                var v135 float64 = 1e10
                for fi136 := 0; fi136 < 11; fi136++ {
                    if (3.14159 <= float32(v135)) {
                        v137 := GenS16{F0: ("foo bar" + "a")}
                        v4 = 92
                        var v138 uint32 = v3
                        _ = v137
                        _ = v138
                    }
                    var v139 uint = 18446744073709551615
                    for fi140 := 0; fi140 < 15; fi140++ {
                        v141 := GenS8{F0: v4, F1: (^(v3)), F2: 0}
                        if false {
                            var v142 int64 = 42
                            var v143 int32 = 42
                            _ = v142
                            _ = v143
                        } else {
                            v3 = v3
                            if (!false) {
                                if true {
                                    fmt.Println(v1, v3)
                                } else {
                                    v18 -= float32(v135)
                                    v144 := GenS17{F0: -1, F1: 255, F2: fi136}
                                    _ = v144
                                }
                                v145 := []int{fi140, (fi136 >> (((v139 | v139)) & 63))}
                                _ = v145
                            }
                        }
                        _ = fi140
                        _ = v141
                    }
                    _ = fi136
                    _ = v139
                }
                _ = v135
            }
        }
        v2 = v2
        _ = v20
    }
    var v146 uint32 = v3
    v147 := []int{}
    v146 += genFn1()
    var v148 float64 = 0.1
    fmt.Println(v4, v18, v148)
    var v149 uint32 = (v3 - (v3 << ((42) & 63)))
    if len(v147) > 0 {
        v147[0] = -154
    }
    var v150 uint32 = v146
    v4 += 42
    if len(v147) > 0 {
        v147[0] = 4294967295
    }
    var v151 uint64 = 136
    if true {
        if ((1 == 1) && true) {
            v147 = append(v147, -9223372036854775808)
            v151 -= 910
        }
        switch ("a" + (v2 + ((v2 + "foo bar") + v2))) {
        case "hello":
            if (0 <= -1) {
                if ((true && true) && (true && false)) {
                    v151 -= (^(v151))
                    v152 := GenS18{F0: 133, F1: v2}
                    var v153 int8 = 4
                    _ = v152
                    _ = v153
                }
            }
        case "\n\t":
            v1 = v18
        default:
            var v154 uint32 = (v146 & 2)
            _ = v154
        }
    } else {
        var v155 string = ((("\n\t" + v2) + "foo bar") + v2)
        v148 = v148
        _ = v155
    }
    fmt.Println("v1:", v1)
    fmt.Println("v146:", v146)
    fmt.Println("slicelen v147:", len(v147), v147)
    fmt.Println("v148:", v148)
    fmt.Println("v149:", v149)
    fmt.Println("v150:", v150)
    fmt.Println("v151:", v151)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v146
    _ = v147
    _ = v148
    _ = v149
    _ = v150
    _ = v151
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
