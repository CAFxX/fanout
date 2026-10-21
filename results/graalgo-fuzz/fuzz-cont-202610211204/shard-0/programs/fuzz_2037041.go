package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64; F2 int8 }
type GenS2 struct { F0 uint8; F1 uint }
type GenS3 struct { F0 int32; F1 int8; F2 uint8 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int8; F1 uint16 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 int8; F1 string; F2 uint32 }
type GenS8 struct { F0 uint8; F1 int64 }
type GenS9 struct { F0 uint16; F1 int16; F2 int }
type GenS10 struct { F0 int8; F1 int64; F2 int16 }
type GenS11 struct { F0 uint32; F1 int64 }
type GenS12 struct { F0 uint8; F1 uint8 }
type GenS13 struct { F0 string }
func genFn0(p0 int16, p1 string) uint32 {
    for fi7 := 0; fi7 < 4; fi7++ {
        p0 = -32768
        v8 := GenS1{F0: false, F1: -1, F2: 127}
        _ = fi7
        _ = v8
    }
    var v9 uint = 668
    var v10 uint64 = 1
    for fi11 := 0; fi11 < 5; fi11++ {
        for fi12 := 0; fi12 < 11; fi12++ {
            switch "\n\t" {
            case "hello":
                fmt.Println(fi11, v9, v9)
            case "":
                if (0 != -4) {
                    v13 := []int{(fi12 & fi11)}
                    if (((false || ((p1 != (p1 + p1)) && (true || true))) && (127 == 100)) || (!false)) {
                        p0 = (p0 + p0)
                        var v14 int8 = 42
                        v15 := GenS2{F0: 42, F1: 0}
                        var v16 uint8 = 155
                        _ = v14
                        _ = v15
                        _ = v16
                    } else {
                        v10 -= (v10 | v10)
                        var v17 string = "hello"
                        _ = v17
                    }
                    _ = v13
                }
            case "\n\t":
                v18 := GenS2{F0: 2, F1: (v9 >> (((v9 << (((v9 << (((v9 >> (((v9 | v9)) & 63))) & 63))) & 63))) & 63))}
                _ = v18
            default:
                p0 = 0
            }
            if true {
                var v19 bool = true
                v20 := GenS1{F0: (128 < 127), F1: -293, F2: -1}
                _ = v19
                _ = v20
            }
            v9 -= uint(v10)
            _ = fi12
        }
        if ((!(1 != 127)) || (255 != 265)) {
            fmt.Println(v10, v9)
            var v21 uint32 = 1
            switch p1 {
            case "a":
                var v22 int16 = p0
                _ = v22
            case "\n\t":
                if (-2.25 > 1e10) {
                    switch p1 {
                    case "foo bar":
                        var v23 int = (-(fi11))
                        _ = v23
                    }
                    var v24 int8 = int8(fi11)
                    _ = v24
                }
            case "hello":
                var v25 uint64 = 216
                _ = v25
            }
            _ = v21
        } else {
            for fi26 := 0; fi26 < 2; fi26++ {
                switch (("hello" + (p1 + p1)) + p1) {
                case "":
                    var v27 uint = (v9 - 4294967295)
                    _ = v27
                case "foo bar":
                    var v28 uint64 = 747
                    _ = v28
                case "a":
                    fmt.Println(v9)
                }
                if ("a" >= "a") {
                    switch (p1 + p1) {
                    case "hello":
                        switch (fi26 + 133) {
                        case -1:
                            v29 := GenS1{F0: (!((fi26 + 127) >= fi26)), F1: 0, F2: -128}
                            _ = v29
                        }
                    case "\n\t":
                        if false {
                            v30 := GenS1{F0: (-1 != 42), F1: 812, F2: 0}
                            var v31 int16 = p0
                            _ = v30
                            _ = v31
                        } else {
                            v32 := []int{int(v9), (fi26 + fi11)}
                            _ = v32
                        }
                    }
                    var v33 uint64 = (v10 - 4294967295)
                    if false {
                        v34 := []int{734}
                        _ = v34
                    }
                    _ = v33
                }
                _ = fi26
            }
            for fi35 := 0; fi35 < 20; fi35++ {
                var v36 uint64 = v10
                var v37 int8 = 96
                switch fi35 {
                case -1:
                    fmt.Println(p1)
                default:
                    var v38 uint64 = 2
                    _ = v38
                }
                if (true || true) {
                    var v39 int = (fi11 & 42)
                    switch 42 {
                    case 42:
                        switch ("" + (("foo bar" + (p1 + "\n\t")) + (((p1 + (p1 + (("" + p1) + (p1 + "")))) + (("" + p1) + p1)) + ""))) {
                        case "\n\t":
                            var v40 float32 = (0.0 / float32(-2.25))
                            _ = v40
                        case "hello":
                            var v41 uint8 = 255
                            _ = v41
                        case "":
                            switch 32768 {
                            case 4294967295:
                                var v42 uint16 = 560
                                _ = v42
                            default:
                                var v43 uint = (v9 >> ((65535) & 63))
                                _ = v43
                            }
                        default:
                            var v44 int32 = 1
                            _ = v44
                        }
                    case -570:
                        v45 := []int{fi35, (fi35 - v39), int(v9)}
                        _ = v45
                    case -674:
                        var v46 uint16 = 1
                        _ = v46
                    }
                    var v47 float64 = -2.25
                    _ = v39
                    _ = v47
                } else {
                    fmt.Println(fi35, v36)
                    p1 = "a"
                    v48 := GenS3{F0: -1, F1: v37, F2: 127}
                    switch (fi35 >> ((v9) & 63)) {
                    case 682:
                        v49 := GenS4{F0: v36}
                        _ = v49
                    case 474:
                        v50 := []int{fi35, (-(fi11)), (fi11 << ((255) & 63))}
                        _ = v50
                    }
                    _ = v48
                }
                fmt.Println(fi11)
                _ = fi35
                _ = v36
                _ = v37
            }
            switch "hello" {
            case "foo bar":
                if true {
                    fmt.Println(fi11)
                    if true {
                        switch p1 {
                        case "a":
                            v51 := GenS4{F0: 2}
                            _ = v51
                        case "":
                            switch ("a" + "hello") {
                            case "foo bar":
                                for fi52 := 0; fi52 < 7; fi52++ {
                                    fmt.Println(v10, fi11)
                                    v53 := GenS4{F0: v10}
                                    v54 := GenS5{F0: 122, F1: 0}
                                    _ = fi52
                                    _ = v53
                                    _ = v54
                                }
                            default:
                                v55 := []int{fi11}
                                _ = v55
                            }
                        }
                        v56 := GenS2{F0: 1, F1: v9}
                        var v57 uint32 = 2147483647
                        _ = v56
                        _ = v57
                    } else {
                        var v58 bool = false
                        if v58 {
                            var v59 uint8 = 191
                            var v60 int16 = (p0 >> ((v9) & 63))
                            _ = v59
                            _ = v60
                        } else {
                            var v61 bool = v58
                            _ = v61
                        }
                        var v62 uint16 = 42
                        _ = v58
                        _ = v62
                    }
                    v63 := GenS4{F0: (v10 * (v10 - v10))}
                    _ = v63
                } else {
                    var v64 uint64 = v10
                    for fi65 := 0; fi65 < 19; fi65++ {
                        v66 := []int{2147483647, fi65}
                        v67 := GenS5{F0: 1, F1: 96}
                        _ = fi65
                        _ = v66
                        _ = v67
                    }
                    _ = v64
                }
            case "a":
                for fi68 := 0; fi68 < 15; fi68++ {
                    v69 := []int{int(p0)}
                    switch 666 {
                    case 431:
                        v70 := GenS6{F0: 469}
                        _ = v70
                    default:
                        if true {
                            switch 418 {
                            case -1:
                                switch p1 {
                                case "hello":
                                    fmt.Println(v9, p0, v9)
                                default:
                                    fmt.Println(fi68, fi68, fi11)
                                }
                            case 743:
                                v71 := GenS7{F0: -6, F1: (p1 + p1), F2: 2147483648}
                                _ = v71
                            }
                            var v72 int8 = 42
                            var v73 float64 = 0.0
                            if len(v69) >= 2 {
                                v69 = v69[1:len(v69)]
                            }
                            _ = v72
                            _ = v73
                        } else {
                            v74 := GenS7{F0: -10, F1: ("a" + (("" + p1) + p1)), F2: 1}
                            v75 := GenS8{F0: 128, F1: -1}
                            _ = v74
                            _ = v75
                        }
                    }
                    v76 := GenS1{F0: (2147483647 != 42), F1: 0, F2: 42}
                    _ = fi68
                    _ = v69
                    _ = v76
                }
            case "":
                fmt.Println(fi11, p0, fi11)
            }
        }
        _ = fi11
    }
    p0 = -161
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return 850
}
func genFn1(p0 bool, p1 uint64) int32 {
    var v77 bool = ((255 < 995) || (!p0))
    switch "hello" {
    case "\n\t":
        fmt.Println(p0, v77, v77)
    case "foo bar":
        fmt.Println(p0, p1)
    }
    if p0 {
        if p0 {
            p1 = (p1 | 574)
            var v78 int8 = 100
            _ = v78
        } else {
            if p0 {
                for fi79 := 0; fi79 < 6; fi79++ {
                    v80 := GenS1{F0: (!p0), F1: 0, F2: 0}
                    _ = fi79
                    _ = v80
                }
                if ((!p0) || v77) {
                    p1 = p1
                    p0 = true
                } else {
                    v77 = v77
                }
                if (!false) {
                    var v81 uint32 = 2147483648
                    var v82 float64 = -2.25
                    _ = v81
                    _ = v82
                }
            } else {
                fmt.Println(v77, v77)
                v77 = p0
                if p0 {
                    var v83 float64 = 1.5
                    _ = v83
                } else {
                    if (v77 && p0) {
                        var v84 uint32 = 912
                        p1 -= 18446744073709551615
                        _ = v84
                    }
                    for fi85 := 0; fi85 < 17; fi85++ {
                        if (!(42 <= 2147483647)) {
                            for fi86 := 0; fi86 < 3; fi86++ {
                                var v87 int = fi86
                                var v88 uint8 = 255
                                _ = fi86
                                _ = v87
                                _ = v88
                            }
                        }
                        _ = fi85
                    }
                    if (p1 > p1) {
                        p0 = p0
                        var v89 string = "hello"
                        v90 := []int{}
                        _ = v89
                        _ = v90
                    }
                }
            }
        }
        var v91 int16 = 32767
        v92 := GenS3{F0: 42, F1: 0, F2: 50}
        _ = v91
        _ = v92
    } else {
        v93 := GenS9{F0: 42, F1: -617, F2: 4294967295}
        _ = v93
    }
    v94 := GenS9{F0: 42, F1: 987, F2: 4294967295}
    for fi95 := 0; fi95 < 16; fi95++ {
        if (-128 > -93) {
            var v96 int16 = 458
            for fi97 := 0; fi97 < 3; fi97++ {
                var v98 string = "hello"
                _ = fi97
                _ = v98
            }
            _ = v96
        } else {
            if p0 {
                v99 := GenS10{F0: 0, F1: 877, F2: 32767}
                var v100 uint64 = p1
                var v101 int32 = -2147483648
                _ = v100
                _ = v101
                _ = v99
            }
            p0 = true
        }
        v102 := GenS11{F0: genFn0(1000, ""), F1: 758}
        p1 -= (p1 + (^(p1)))
        _ = fi95
        _ = v102
    }
    _ = p0
    _ = p1
    _ = v77
    _ = v94
    return 2147483647
}

func main() {
    var v1 uint16 = 42
    v2 := make(map[string]uint32)
    v3 := make(map[uint32]uint32)
    v4 := []int{}
    var v5 uint8 = 130
    v6 := make(map[int]string)
    delete(v6, 1)
    v103 := GenS6{F0: 438}
    var v104 int32 = 0
    switch ("hello" + "a") {
    case "\n\t":
        if (18446744073709551615 < 42) {
            for fi105 := 0; fi105 < 16; fi105++ {
                switch ("a" + "a") {
                case "":
                    v106 := GenS10{F0: -98, F1: -9223372036854775808, F2: 368}
                    _ = v106
                case "\n\t":
                    v1 = 952
                default:
                    v5 = (v5 ^ (^(v5)))
                }
                if len(v4) > 0 {
                    v4[0] = fi105
                }
                _ = fi105
            }
            var v107 int32 = (v104 << ((416) & 63))
            for fi108 := 0; fi108 < 13; fi108++ {
                switch "\n\t" {
                case "":
                    if ((v104 * v107) < v104) {
                        fmt.Println(fi108, v1, fi108)
                        v4 = append(v4, (fi108 - -996))
                    } else {
                        var v109 string = ""
                        _ = v109
                    }
                case "\n\t":
                    var v110 uint8 = (v5 & v5)
                    _ = v110
                }
                fmt.Println(v107, v107, fi108)
                _ = fi108
            }
            _ = v107
        } else {
            var v111 uint16 = v1
            _ = v111
        }
    case "a":
        for fi112 := 0; fi112 < 13; fi112++ {
            v113 := GenS12{F0: 0, F1: v5}
            v114 := GenS4{F0: 4294967295}
            _ = fi112
            _ = v113
            _ = v114
        }
    case "":
        var v115 uint64 = 899
        _ = v115
    default:
        switch 258 {
        case 128:
            fmt.Println(v5)
        case 127:
            switch 477 {
            case 4294967295:
                v116 := GenS13{F0: ""}
                _ = v116
            default:
                v5 += (^(v5))
            }
        default:
            var v117 uint32 = genFn0(1, "a")
            _ = v117
        }
    }
    if true {
        if len(v4) > 0 {
            v4[0] = -9223372036854775808
        }
    } else {
        var v118 int64 = 9223372036854775807
        var v119 string = (("foo bar" + ("hello" + "")) + "a")
        for fi120 := 0; fi120 < 4; fi120++ {
            fmt.Println(v119, v119, v118)
            _ = fi120
        }
        _ = v118
        _ = v119
    }
    for fi121 := 0; fi121 < 13; fi121++ {
        var v122 int64 = -188
        _ = fi121
        _ = v122
    }
    for fi123 := 0; fi123 < 14; fi123++ {
        v104 -= (v104 * genFn1(false, uint64(v5)))
        v2["hello"] = 4294967295
        var v124 int16 = 768
        _ = fi123
        _ = v124
    }
    switch "hello" {
    case "":
        v125 := v6[127]
        _ = v125
    }
    var v126 bool = ("" > "a")
    fmt.Println(v1, v1, v104)
    fmt.Println("v1:", v1)
    fmt.Println("v104:", v104)
    fmt.Println("v126:", v126)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v103
    _ = v104
    _ = v126
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
