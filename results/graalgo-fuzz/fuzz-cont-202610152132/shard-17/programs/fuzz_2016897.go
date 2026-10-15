package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32 }
type GenS2 struct { F0 int16; F1 uint16; F2 string }
type GenS3 struct { F0 bool; F1 uint16; F2 bool }
type GenS4 struct { F0 uint32; F1 int16 }
type GenS5 struct { F0 int16; F1 uint64; F2 uint16 }
type GenS6 struct { F0 uint }
func genFn0(p0 uint64, p1 int, p2 string) string {
    p0 = (p0 >> ((42) & 63))
    v4 := GenS1{F0: p1, F1: 2147483647}
    fmt.Println(p1, p1)
    v5 := GenS1{F0: p1, F1: 42}
    if true {
        switch p2 {
        case "foo bar":
            p0 -= (p0 & 0)
        case "a":
            fmt.Println(p0, p2, p2)
        default:
            if true {
                if (-91 == (p1 << ((255) & 63))) {
                    v6 := []int{}
                    if len(v6) > 0 {
                        v6[0] = p1
                    }
                    for fi7 := 0; fi7 < 8; fi7++ {
                        if (0 >= 1) {
                            v8 := GenS2{F0: -32768, F1: 32767, F2: p2}
                            _ = v8
                        } else {
                            for fi9 := 0; fi9 < 12; fi9++ {
                                p2 = (((p2 + (("a" + p2) + p2)) + p2) + p2)
                                _ = fi9
                            }
                        }
                        _ = fi7
                    }
                    _ = v6
                }
                v10 := []int{42}
                _ = v10
            }
        }
        if (p0 >= 0) {
            for fi11 := 0; fi11 < 12; fi11++ {
                p2 = ("a" + "hello")
                if (!(71 > -1)) {
                    var v12 float64 = -2.25
                    if (591 != 2147483648) {
                        v13 := GenS1{F0: fi11, F1: -2147483648}
                        var v14 uint16 = 255
                        fmt.Println(p2)
                        _ = v13
                        _ = v14
                    } else {
                        for fi15 := 0; fi15 < 12; fi15++ {
                            var v16 string = p2
                            var v17 string = ""
                            var v18 int64 = 968
                            _ = fi15
                            _ = v16
                            _ = v17
                            _ = v18
                        }
                        v19 := GenS1{F0: p1, F1: -1}
                        _ = v19
                    }
                    for fi20 := 0; fi20 < 13; fi20++ {
                        v12 = v12
                        var v21 int8 = -126
                        switch "hello" {
                        case "":
                            v12 = (0.1 + (0.1 / v12))
                        case "foo bar":
                            if true {
                                v22 := []int{489, fi20, (p1 * (fi11 / (((fi20 | 42)) | 1))), p1}
                                var v23 uint64 = p0
                                if true {
                                    _ = len(v22)
                                    var v24 float32 = 1.5
                                    var v25 string = "\n\t"
                                    var v26 int = (^(p1))
                                    _ = v24
                                    _ = v25
                                    _ = v26
                                }
                                _ = v22
                                _ = v23
                            }
                        case "a":
                            var v27 uint = 18446744073709551615
                            _ = v27
                        default:
                            switch 347 {
                            case 78:
                                if false {
                                    p0 -= (p0 >> ((284) & 63))
                                    v28 := GenS3{F0: true, F1: 547, F2: false}
                                    v29 := []int{13, 0, 127, (p1 | (^(fi11)))}
                                    var v30 int64 = 290
                                    v29 = append(v29, fi20)
                                    _ = v28
                                    _ = v29
                                    _ = v30
                                } else {
                                    v31 := GenS4{F0: 954, F1: 0}
                                    var v32 string = p2
                                    _ = v31
                                    _ = v32
                                }
                            case -145:
                                v12 += v12
                            }
                        }
                        _ = fi20
                        _ = v21
                    }
                    _ = v12
                }
                _ = fi11
            }
            v33 := []int{2147483647, p1, (p1 | (p1 & p1)), (p1 * (^(p1)))}
            p0 -= p0
            _ = v33
        } else {
            v34 := GenS5{F0: 1000, F1: (p0 + (p0 + p0)), F2: 180}
            _ = v34
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return "hello"
}
func genFn1(p0 int32, p1 int8) uint32 {
    v35 := GenS1{F0: -1, F1: p0}
    if (2 == 32767) {
        v36 := GenS5{F0: 32767, F1: 383, F2: 0}
        if true {
            p0 = 1
            fmt.Println(p0, p0)
        } else {
            var v37 float64 = -2.25
            _ = v37
        }
        p0 -= 42
        _ = v36
    }
    p1 = (p1 % ((4) | 1))
    if (true || (2 > 0)) {
        v38 := []int{0, 127, 42, 165}
        _ = v38
    }
    _ = p0
    _ = p1
    _ = v35
    return 0
}

func main() {
    var v1 uint16 = 2
    var v2 uint64 = 4294967295
    v3 := []int{42}
    v1 += (v1 * (v1 << ((2) & 63)))
    if len(v3) > 0 {
        v3[0] = 32768
    }
    switch 2147483647 {
    case 4294967295:
        for fi39 := 0; fi39 < 14; fi39++ {
            if false {
                if (false || (!((501 > v1) && (!(false && true))))) {
                    for fi40 := 0; fi40 < 20; fi40++ {
                        switch ("hello" + "\n\t") {
                        case "a":
                            fmt.Println(fi40, fi39)
                        case "foo bar":
                            v2 += (v2 | (v2 - (v2 | (v2 - (^(v2))))))
                        case "":
                            var v41 uint64 = 42
                            _ = v41
                        default:
                            switch ("hello" + "hello") {
                            case "foo bar":
                                if (!((true || (("\n\t" + ("foo bar" + "")) > ("hello" + "\n\t"))) || true)) {
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    v2 += 231
                                }
                            case "\n\t":
                                var v42 uint16 = (v1 / ((v1) | 1))
                                _ = v42
                            case "a":
                                v43 := GenS4{F0: 0, F1: 1000}
                                _ = v43
                            default:
                                v44 := GenS3{F0: true, F1: v1, F2: (!(1 >= -2147483648))}
                                _ = v44
                            }
                        }
                        var v45 int32 = 0
                        v2 += v2
                        _ = fi40
                        _ = v45
                    }
                    var v46 int = fi39
                    _ = v46
                } else {
                    var v47 uint8 = 220
                    v1 += 42
                    if false {
                        var v48 int16 = 42
                        v47 -= (v47 << ((18446744073709551615) & 63))
                        _ = v48
                    } else {
                        var v49 int = fi39
                        fmt.Println(v1)
                        _ = v49
                    }
                    _ = v47
                }
            }
            if (!true) {
                if true {
                    switch fi39 {
                    case 4294967295:
                        var v50 uint16 = v1
                        _ = v50
                    case 128:
                        if false {
                            switch "hello" {
                            case "\n\t":
                                var v51 int32 = -214
                                _ = v51
                            case "a":
                                if (false && (-663 <= 0)) {
                                    switch (fi39 ^ fi39) {
                                    case -913:
                                        if true {
                                            v1 = 0
                                        } else {
                                            v2 = 1
                                            var v52 int8 = -77
                                            switch -1 {
                                            case 4294967295:
                                                fmt.Println(v52)
                                            case -738:
                                                if (((3.14159 - 0.1) / float32(((3.14159 - 0.0) / (3.14159 + (0.1 * float64(float32(float64(((-2.25 + (-2.25 - 1.5)) + 0.1))))))))) >= 1e10) {
                                                    for fi53 := 0; fi53 < 11; fi53++ {
                                                        switch fi53 {
                                                        case 9223372036854775807:
                                                            var v54 float64 = (0.0 + 1e10)
                                                            _ = v54
                                                        case 128:
                                                            var v55 uint16 = 2
                                                            _ = v55
                                                        case -2147483648:
                                                            var v56 int32 = -1
                                                            _ = v56
                                                        }
                                                        v52 += v52
                                                        _ = fi53
                                                    }
                                                } else {
                                                    v2 += v2
                                                    for fi57 := 0; fi57 < 15; fi57++ {
                                                        if true {
                                                            var v58 uint = 65535
                                                            _ = v58
                                                        } else {
                                                            var v59 uint16 = 255
                                                            v60 := GenS2{F0: 843, F1: v59, F2: "foo bar"}
                                                            _ = v59
                                                            _ = v60
                                                        }
                                                        _ = fi57
                                                    }
                                                }
                                            case 0:
                                                fmt.Println(v1)
                                            default:
                                                var v61 int = 9223372036854775807
                                                _ = v61
                                            }
                                            _ = v52
                                        }
                                    case 32768:
                                        v62 := GenS4{F0: 4294967295, F1: 32767}
                                        _ = v62
                                    }
                                }
                            case "hello":
                                var v63 uint64 = 0
                                _ = v63
                            }
                        } else {
                            var v64 int = -1
                            if (("\n\t" + "a") >= "foo bar") {
                                var v65 int32 = -2147483648
                                _ = v65
                            } else {
                                if (-128 != 0) {
                                    fmt.Println(v1, v2, v64)
                                }
                                var v66 bool = true
                                _ = v66
                            }
                            _ = v64
                        }
                    default:
                        var v67 uint = 1
                        _ = v67
                    }
                    for fi68 := 0; fi68 < 1; fi68++ {
                        if true {
                            switch fi39 {
                            case 42:
                                if len(v3) > 0 {
                                    v3[0] = (fi68 << ((255) & 63))
                                }
                                v2 += v2
                                v1 = v1
                            case 128:
                                switch fi68 {
                                case 128:
                                    var v69 string = ("foo bar" + ("\n\t" + "hello"))
                                    _ = v69
                                case -135:
                                    switch ("foo bar" + ("foo bar" + "")) {
                                    case "":
                                        var v70 float64 = (1e10 * ((3.14159 * (((0.0 - -2.25) + (0.1 * 1e10)) + 1.5)) / 1.0))
                                        _ = v70
                                    }
                                default:
                                    switch "" {
                                    case "":
                                        var v71 uint16 = (v1 & (v1 >> ((0) & 63)))
                                        var v72 int16 = 0
                                        fmt.Println(v1, v71)
                                        var v73 bool = false
                                        _ = v71
                                        _ = v72
                                        _ = v73
                                    case "a":
                                        v1 -= (v1 >> ((255) & 63))
                                    case "\n\t":
                                        v2 += v2
                                    }
                                }
                            default:
                                v1 -= (v1 ^ 2)
                            }
                        } else {
                            fmt.Println(fi39, v1, fi39)
                            v2 = v2
                            switch fi68 {
                            case 102:
                                var v74 float64 = 3.14159
                                _ = v74
                            case -394:
                                v1 = v1
                            case -1:
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            }
                        }
                        if (!true) {
                            var v75 bool = false
                            switch "\n\t" {
                            case "hello":
                                fmt.Println(v2)
                            case "":
                                var v76 int64 = -1
                                _ = v76
                            case "a":
                                var v77 int = 127
                                _ = v77
                            }
                            var v78 uint32 = 2147483648
                            v3 = append(v3, (fi68 ^ (fi39 - fi68)))
                            fmt.Println(v1, fi68)
                            _ = v75
                            _ = v78
                        }
                        v3 = append(v3, fi68)
                        _ = fi68
                    }
                } else {
                    var v79 string = ""
                    _ = v79
                }
                switch (^(fi39)) {
                case 0:
                    var v80 bool = (v1 <= (v1 + (v1 - 1)))
                    _ = v80
                case 127:
                    v2 = 192
                case 845:
                    for fi81 := 0; fi81 < 4; fi81++ {
                        var v82 int16 = -963
                        var v83 bool = ((v1 >= v1) || false)
                        fmt.Println(fi81, v1)
                        _ = fi81
                        _ = v82
                        _ = v83
                    }
                }
            } else {
                switch 9223372036854775807 {
                case 2147483647:
                    if false {
                        if (!(1 >= 255)) {
                            v1 += 680
                            var v84 uint8 = 2
                            var v85 int8 = -25
                            _ = v84
                            _ = v85
                        } else {
                            if ((415 != 1) || true) {
                                v86 := GenS3{F0: ((^(v1)) == 42), F1: 0, F2: (false || false)}
                                var v87 int = (fi39 / (((fi39 - fi39)) | 1))
                                v1 += (^(v1))
                                _ = v86
                                _ = v87
                            } else {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                _ = len(v3)
                                var v88 int16 = -1
                                _ = v88
                            }
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        }
                        fmt.Println(fi39)
                        var v89 bool = false
                        _ = v89
                    } else {
                        for fi90 := 0; fi90 < 17; fi90++ {
                            var v91 uint16 = (v1 >> ((255) & 63))
                            _ = fi90
                            _ = v91
                        }
                        v1 = 0
                        var v92 int16 = -32768
                        _ = v92
                    }
                case -9223372036854775808:
                    _ = len(v3)
                }
            }
            _ = fi39
        }
    case 65:
        var v93 int64 = int64(v2)
        _ = v93
    case 853:
        var v94 float32 = 3.14159
        _ = v94
    default:
        if ((111 != 30) || true) {
            for fi95 := 0; fi95 < 15; fi95++ {
                v1 -= 558
                if true {
                    v1 = (v1 | (v1 * 0))
                    switch 4294967295 {
                    case 0:
                        v96 := GenS6{F0: 42}
                        _ = v96
                    case 9223372036854775807:
                        v2 += v2
                    case -895:
                        for fi97 := 0; fi97 < 18; fi97++ {
                            var v98 string = (((("a" + ("a" + ("foo bar" + ("" + "")))) + (("\n\t" + "a") + ("hello" + "hello"))) + "a") + "foo bar")
                            v99 := GenS1{F0: 128, F1: -673}
                            fmt.Println(v1)
                            _ = fi97
                            _ = v98
                            _ = v99
                        }
                    }
                }
                _ = fi95
            }
            var v100 string = ("hello" + "")
            _ = v100
        }
    }
    var v101 uint8 = 1
    var v102 int32 = 1
    var v103 int8 = -60
    if (!(!false)) {
        for fi104 := 0; fi104 < 10; fi104++ {
            for fi105 := 0; fi105 < 9; fi105++ {
                switch "\n\t" {
                case "foo bar":
                    v106 := GenS1{F0: (fi104 - 2147483647), F1: v102}
                    _ = v106
                case "a":
                    var v107 uint = 4294967295
                    v107 -= v107
                    var v108 uint8 = (v101 >> ((v107) & 63))
                    _ = v107
                    _ = v108
                case "":
                    fmt.Println(v101)
                default:
                    fmt.Println(fi105, fi104)
                }
                fmt.Println(fi105, v1, v2)
                v101 = v101
                _ = fi105
            }
            _ = fi104
        }
    }
    v103 = (v103 + (v103 & (v103 >> ((321) & 63))))
    var v109 uint64 = v2
    v102 = v102
    for fi110 := 0; fi110 < 19; fi110++ {
        v102 += v102
        _ = len(v3)
        _ = fi110
    }
    if len(v3) > 0 {
        v3[0] = 32768
    }
    var v111 uint32 = 0
    for fi112 := 0; fi112 < 1; fi112++ {
        var v113 int64 = 9223372036854775807
        var v114 uint64 = (v2 % ((v2) | 1))
        v103 += 100
        _ = fi112
        _ = v113
        _ = v114
    }
    v115 := GenS1{F0: 128, F1: 1}
    v102 += (v102 * -1)
    v116 := GenS3{F0: (-1 <= -1), F1: (^(v1)), F2: false}
    fmt.Println("v1:", v1)
    fmt.Println("v101:", v101)
    fmt.Println("v102:", v102)
    fmt.Println("v103:", v103)
    fmt.Println("v109:", v109)
    fmt.Println("v111:", v111)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v101
    _ = v102
    _ = v103
    _ = v109
    _ = v111
    _ = v115
    _ = v116
    _ = v2
    _ = v3
}
