package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint; F1 int64; F2 string }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint64; F1 int16 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint32; F1 int64; F2 string }
type GenS7 struct { F0 int32; F1 bool }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 int16; F1 int16 }
type GenS10 struct { F0 uint32; F1 uint64; F2 int32 }
type GenS11 struct { F0 uint64 }
type GenS12 struct { F0 bool; F1 uint32 }
type GenS13 struct { F0 int32; F1 uint32; F2 uint }
type GenS14 struct { F0 int32; F1 uint32 }
type GenS15 struct { F0 int8 }
type GenS16 struct { F0 uint; F1 int64 }
func genFn0() bool {
    v4 := []int{-2147483648, -74}
    var v5 int8 = 100
    v5 += v5
    _ = v4
    _ = v5
    return false
}
func genFn1(p0 bool, p1 bool) uint16 {
    v6 := []int{-9223372036854775808, 4294967295, 32767}
    var v7 uint16 = 921
    var v8 uint32 = 65535
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    for fi9 := 0; fi9 < 15; fi9++ {
        if p0 {
            p0 = p1
        } else {
            var v10 uint16 = v7
            if false {
                v11 := GenS1{F0: v8}
                for fi12 := 0; fi12 < 16; fi12++ {
                    if ((1e10 + 0.1) == 1e10) {
                        v13 := GenS2{F0: 537, F1: 673, F2: ("a" + "")}
                        _ = v13
                    } else {
                        switch ("a" + "\n\t") {
                        case "a":
                            v10 = v7
                        case "foo bar":
                            if p1 {
                                if false {
                                    var v14 uint8 = 142
                                    var v15 string = (("hello" + (("\n\t" + "hello") + "foo bar")) + ("\n\t" + ""))
                                    switch v15 {
                                    case "foo bar":
                                        v16 := GenS1{F0: 65535}
                                        _ = v16
                                    }
                                    _ = v14
                                    _ = v15
                                }
                                if len(v6) >= 2 {
                                    v6 = v6[1:len(v6)]
                                }
                            }
                        case "\n\t":
                            v7 -= 255
                            fmt.Println(p0)
                            fmt.Println(fi9, fi12, p1)
                        default:
                            switch (("a" + (("\n\t" + "\n\t") + "hello")) + "") {
                            case "a":
                                v8 = v8
                            case "\n\t":
                                var v17 uint16 = (v7 | 0)
                                _ = v17
                            default:
                                p0 = p1
                            }
                            if true {
                                v18 := GenS3{F0: p1}
                                _ = v18
                            }
                        }
                        if len(v6) > 0 {
                            v6[0] = (fi9 * fi9)
                        }
                    }
                    if true {
                        fmt.Println(v8)
                    } else {
                        switch ("foo bar" + ("\n\t" + "a")) {
                        case "a":
                            v7 = (v7 & (^(v7)))
                        default:
                            v19 := GenS1{F0: 42}
                            _ = v19
                        }
                        fmt.Println(v10, v8)
                    }
                    v20 := GenS4{F0: uint64(v8), F1: 317}
                    _ = fi12
                    _ = v20
                }
                switch "foo bar" {
                case "":
                    var v21 int = fi9
                    _ = v21
                case "\n\t":
                    p1 = p0
                case "a":
                    if (723 <= 511) {
                        v22 := GenS2{F0: 42, F1: 0, F2: "foo bar"}
                        var v23 uint8 = 2
                        _ = v22
                        _ = v23
                    } else {
                        var v24 uint64 = 0
                        _ = v24
                    }
                default:
                    fmt.Println(v10, v7, v8)
                }
                _ = v11
            } else {
                switch ("a" + "") {
                case "":
                    if false {
                        v10 = v10
                        var v25 int32 = -1
                        _ = v25
                    } else {
                        var v26 uint32 = v8
                        fmt.Println(fi9, v7, fi9)
                        _ = v26
                    }
                case "hello":
                    var v27 string = ("" + ("a" + (("a" + "\n\t") + ("hello" + ("" + "\n\t")))))
                    _ = v27
                }
                var v28 uint8 = 2
                v29 := GenS3{F0: ("foo bar" >= "hello")}
                _ = v28
                _ = v29
            }
            _ = v10
        }
        fmt.Println(v8, v7, v7)
        _ = fi9
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return (v7 - v7)
}

func main() {
    var v1 string = "foo bar"
    var v2 uint64 = 42
    var v3 int16 = -278
    var v30 int32 = 627
    switch "a" {
    case "foo bar":
        var v31 int16 = 1
        _ = v31
    case "\n\t":
        if ((1e10 + 1e10) > ((0.0 / (1.5 / 1.0)) * (((0.0 - float64(((float32(-2.25) * (float32(0.0) - ((0.0 / -2.25) * ((0.1 + 0.1) / 1e10)))) / -2.25))) / 0.1) - (3.14159 / float64(0.1))))) {
            v32 := GenS3{F0: (false || (!false))}
            if false {
                for fi33 := 0; fi33 < 1; fi33++ {
                    for fi34 := 0; fi34 < 1; fi34++ {
                        v35 := GenS5{F0: v1}
                        var v36 string = (v1 + "a")
                        if (817 != 616) {
                            v1 = v36
                            if (100 <= 127) {
                                v1 = "\n\t"
                                v37 := GenS6{F0: 2147483648, F1: 0, F2: ("hello" + "foo bar")}
                                fmt.Println(fi33, v3, fi33)
                                _ = v37
                            }
                        } else {
                            v38 := []int{fi33, (fi34 / (((fi34 & fi34)) | 1)), 32767, (fi34 + -1)}
                            var v39 string = "a"
                            fmt.Println(fi34, fi34)
                            _ = v38
                            _ = v39
                        }
                        switch fi33 {
                        case -588:
                            var v40 int8 = -80
                            _ = v40
                        case 128:
                            if (255 <= 1) {
                                var v41 uint64 = 4294967295
                                _ = v41
                            }
                        case 4294967295:
                            v42 := []int{fi33, (fi33 + (fi34 + fi33)), 127, (fi34 - fi33)}
                            _ = v42
                        default:
                            if (!false) {
                                if (2 >= 2) {
                                    var v43 uint8 = 125
                                    var v44 bool = ((!true) && (150 >= 73))
                                    _ = v43
                                    _ = v44
                                } else {
                                    var v45 int16 = 32767
                                    var v46 uint8 = 128
                                    v47 := GenS6{F0: 42, F1: 0, F2: v1}
                                    _ = v45
                                    _ = v46
                                    _ = v47
                                }
                                if false {
                                    var v48 string = ("foo bar" + v36)
                                    v49 := []int{}
                                    fmt.Println(v2)
                                    _ = v48
                                    _ = v49
                                } else {
                                    switch fi34 {
                                    case 4294967295:
                                        var v50 uint8 = 2
                                        _ = v50
                                    case -912:
                                        v1 = ((((v36 + ("hello" + "a")) + v1) + ((v1 + ("\n\t" + "")) + (v36 + ((v36 + v1) + v1)))) + v36)
                                    }
                                }
                            }
                        }
                        _ = fi34
                        _ = v35
                        _ = v36
                    }
                    _ = fi33
                }
            } else {
                v2 += 4294967295
                for fi51 := 0; fi51 < 3; fi51++ {
                    v3 = 1
                    if false {
                        v30 = 42
                    } else {
                        if (true && false) {
                            switch "\n\t" {
                            case "foo bar":
                                for fi52 := 0; fi52 < 6; fi52++ {
                                    if (genFn1(true, (uint(v2) < 18446744073709551615)) <= 32767) {
                                        if true {
                                            var v53 uint = 65535
                                            _ = v53
                                        }
                                        var v54 string = "a"
                                        _ = v54
                                    }
                                    _ = fi52
                                }
                            }
                            switch (v1 + (v1 + v1)) {
                            case "\n\t":
                                switch "\n\t" {
                                case "a":
                                    v55 := GenS3{F0: false}
                                    _ = v55
                                case "\n\t":
                                    v56 := []int{(-(fi51)), (^(fi51))}
                                    _ = v56
                                default:
                                    var v57 uint16 = 255
                                    _ = v57
                                }
                            case "":
                                if true {
                                    switch (-(fi51)) {
                                    case -783:
                                        fmt.Println(v30, fi51)
                                    case 524:
                                        if (2 > 4294967295) {
                                            var v58 bool = (65535 <= 42)
                                            _ = v58
                                        } else {
                                            var v59 int = (fi51 | 4294967295)
                                            _ = v59
                                        }
                                    default:
                                        var v60 int32 = 1
                                        _ = v60
                                    }
                                    fmt.Println(v2)
                                } else {
                                    v61 := []int{746, (fi51 - int(v30)), 867}
                                    _ = v61
                                }
                            case "foo bar":
                                v62 := GenS7{F0: (v30 % ((v30) | 1)), F1: (false || (true && (0.0 == (-2.25 / (0.1 + (0.0 - 3.14159))))))}
                                _ = v62
                            }
                            v3 = -919
                        }
                    }
                    if false {
                        var v63 uint64 = (v2 << ((2) & 63))
                        switch (v1 + (v1 + (v1 + ("foo bar" + (v1 + v1))))) {
                        case "\n\t":
                            v64 := []int{fi51, (fi51 * -252), fi51}
                            _ = v64
                        case "a":
                            switch v1 {
                            case "":
                                if true {
                                    fmt.Println(fi51, v63, fi51)
                                    var v65 int32 = v30
                                    var v66 uint32 = 4294967295
                                    _ = v65
                                    _ = v66
                                } else {
                                    v67 := GenS8{F0: 196}
                                    if ((!(fi51 < fi51)) || false) {
                                        for fi68 := 0; fi68 < 5; fi68++ {
                                            v69 := []int{fi68, -1, -1, (-(fi51))}
                                            _ = fi68
                                            _ = v69
                                        }
                                        var v70 int = (fi51 + fi51)
                                        v71 := GenS9{F0: (v3 ^ v3), F1: v3}
                                        _ = v70
                                        _ = v71
                                    } else {
                                        var v72 float64 = -2.25
                                        _ = v72
                                    }
                                    _ = v67
                                }
                            case "a":
                                v73 := []int{fi51, 32767, fi51, (fi51 | (fi51 % ((int(v2)) | 1)))}
                                _ = v73
                            default:
                                if true {
                                    switch fi51 {
                                    case 270:
                                        for fi74 := 0; fi74 < 16; fi74++ {
                                            if true {
                                                v75 := GenS10{F0: 704, F1: (^(v2)), F2: v30}
                                                _ = v75
                                            } else {
                                                v76 := []int{(fi74 << ((454) & 63)), -1, (fi74 / ((-9223372036854775808) | 1)), -9223372036854775808}
                                                _ = v76
                                            }
                                            switch fi51 {
                                            case 128:
                                                switch ((v1 + ("\n\t" + "\n\t")) + (v1 + (v1 + (((v1 + v1) + v1) + ("\n\t" + "hello"))))) {
                                                case "":
                                                    v30 -= v30
                                                    var v77 int = -443
                                                    _ = v77
                                                default:
                                                    fmt.Println(fi51)
                                                }
                                            case -9223372036854775808:
                                                fmt.Println(v30, fi51, fi51)
                                            case 42:
                                                v3 = 42
                                            }
                                            _ = fi74
                                        }
                                    case 0:
                                        v78 := []int{fi51, (fi51 << ((4294967295) & 63)), (fi51 >> ((42) & 63)), (fi51 * fi51)}
                                        _ = v78
                                    case 431:
                                        switch ("foo bar" + ("hello" + v1)) {
                                        case "":
                                            v2 -= (v2 ^ 2)
                                        case "foo bar":
                                            v30 = -2147483648
                                        case "\n\t":
                                            var v79 bool = true
                                            _ = v79
                                        default:
                                            switch (fi51 << ((18446744073709551615) & 63)) {
                                            case 128:
                                                if (!false) {
                                                    if (-686 >= (-(fi51))) {
                                                        v30 = v30
                                                        for fi80 := 0; fi80 < 5; fi80++ {
                                                            v3 -= (-(v3))
                                                            _ = fi80
                                                        }
                                                        if true {
                                                            for fi81 := 0; fi81 < 12; fi81++ {
                                                                v63 = (^(v63))
                                                                var v82 int = (fi81 | (^(fi81)))
                                                                var v83 float32 = 0.0
                                                                _ = fi81
                                                                _ = v82
                                                                _ = v83
                                                            }
                                                        } else {
                                                            for fi84 := 0; fi84 < 1; fi84++ {
                                                                v85 := GenS11{F0: (v63 | 0)}
                                                                _ = fi84
                                                                _ = v85
                                                            }
                                                            var v86 uint32 = 2147483647
                                                            _ = v86
                                                        }
                                                    } else {
                                                        v63 += (v63 + (^(v2)))
                                                        switch fi51 {
                                                        case -905:
                                                            v87 := []int{(fi51 / ((fi51) | 1)), fi51, int(v30)}
                                                            _ = v87
                                                        case 32767:
                                                            if (!true) {
                                                                fmt.Println(fi51, v30, v63)
                                                                for fi88 := 0; fi88 < 1; fi88++ {
                                                                    v2 -= 206
                                                                    _ = fi88
                                                                }
                                                                v3 += (v3 - -1)
                                                            }
                                                        }
                                                        var v89 uint64 = (v63 - (v2 * 4294967295))
                                                        _ = v89
                                                    }
                                                    switch 32768 {
                                                    case 9223372036854775807:
                                                        var v90 int8 = -128
                                                        _ = v90
                                                    default:
                                                        var v91 int16 = (-(v3))
                                                        _ = v91
                                                    }
                                                } else {
                                                    fmt.Println(v1, fi51)
                                                    v92 := []int{(-(fi51))}
                                                    if true {
                                                        if len(v92) > 0 {
                                                            v92[0] = fi51
                                                        }
                                                        for fi93 := 0; fi93 < 14; fi93++ {
                                                            v94 := GenS12{F0: false, F1: 1}
                                                            if len(v92) >= 2 {
                                                                v92 = v92[1:len(v92)]
                                                            }
                                                            var v95 uint = 0
                                                            v96 := GenS13{F0: 255, F1: 2147483648, F2: (v95 * (v95 % (((v95 - v95)) | 1)))}
                                                            v92 = append(v92, -654)
                                                            _ = fi93
                                                            _ = v94
                                                            _ = v95
                                                            _ = v96
                                                        }
                                                    }
                                                    _ = v92
                                                }
                                            case 32767:
                                                for fi97 := 0; fi97 < 9; fi97++ {
                                                    v98 := GenS14{F0: (^(v30)), F1: 1}
                                                    if (false || false) {
                                                        fmt.Println(v30, v30)
                                                    } else {
                                                        v99 := []int{fi97, 426, fi97, fi97}
                                                        fmt.Println(v1, v3)
                                                        if len(v99) >= 2 {
                                                            v99 = v99[1:len(v99)]
                                                        }
                                                        _ = v99
                                                    }
                                                    _ = fi97
                                                    _ = v98
                                                }
                                            default:
                                                var v100 int16 = (v3 - (-(v3)))
                                                _ = v100
                                            }
                                        }
                                    }
                                }
                            }
                        default:
                            switch v1 {
                            case "hello":
                                v101 := []int{fi51}
                                _ = v101
                            case "foo bar":
                                v102 := GenS3{F0: ((!false) || false)}
                                _ = v102
                            default:
                                var v103 int8 = 20
                                _ = v103
                            }
                        }
                        var v104 float64 = ((0.1 + 1.5) * (0.0 / float64((0.0 - 1e10))))
                        _ = v104
                        _ = v63
                    } else {
                        var v105 float32 = -2.25
                        for fi106 := 0; fi106 < 13; fi106++ {
                            if false {
                                if (!(v3 != 42)) {
                                    var v107 int8 = 1
                                    _ = v107
                                }
                                v108 := []int{(fi106 << ((255) & 63)), 9223372036854775807, (-(fi51)), fi106}
                                _ = v108
                            } else {
                                if false {
                                    v105 = (v105 * v105)
                                }
                                fmt.Println(v30)
                                v109 := GenS9{F0: 1, F1: v3}
                                _ = v109
                            }
                            _ = fi106
                        }
                        for fi110 := 0; fi110 < 8; fi110++ {
                            v2 += (v2 & v2)
                            v111 := []int{fi51}
                            _ = fi110
                            _ = v111
                        }
                        _ = v105
                    }
                    _ = fi51
                }
                if (false && (4294967295 < 1)) {
                    v2 -= 0
                }
            }
            switch -2147483648 {
            case 488:
                if (v1 != "hello") {
                    var v112 int32 = v30
                    v113 := []int{-337}
                    v112 = v30
                    _ = v112
                    _ = v113
                }
            case 2147483647:
                if false {
                    var v114 bool = false
                    v115 := GenS15{F0: -1}
                    v1 = ""
                    _ = v114
                    _ = v115
                }
            default:
                var v116 uint8 = 85
                _ = v116
            }
            _ = v32
        }
    case "a":
        v2 += 0
    }
    var v117 uint = 0
    v3 -= v3
    fmt.Println(v30, v1)
    v117 = (v117 & v117)
    if ((!true) && true) {
        var v118 uint32 = 0
        v119 := GenS15{F0: 42}
        _ = v118
        _ = v119
    } else {
        v3 += (v3 >> (((v117 | (v117 >> ((v117) & 63)))) & 63))
    }
    v2 += 42
    for fi120 := 0; fi120 < 12; fi120++ {
        var v121 int32 = v30
        v3 -= -32768
        _ = fi120
        _ = v121
    }
    v2 += (v2 + (v2 >> (((v117 * (^(v117)))) & 63)))
    var v122 string = ""
    v123 := GenS16{F0: v117, F1: 641}
    switch -866 {
    case -924:
        var v124 uint8 = 1
        _ = v124
    default:
        var v125 float32 = -2.25
        _ = v125
    }
    v126 := GenS8{F0: 163}
    fmt.Println("v1:", v1)
    fmt.Println("v117:", v117)
    fmt.Println("v122:", v122)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    _ = v1
    _ = v117
    _ = v122
    _ = v123
    _ = v126
    _ = v2
    _ = v3
    _ = v30
}
