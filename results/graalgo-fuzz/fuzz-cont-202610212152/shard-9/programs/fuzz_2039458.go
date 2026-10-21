package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32 }
type GenS2 struct { F0 uint; F1 int8 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint16; F1 int; F2 int16 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 string; F1 int64 }
type GenS7 struct { F0 uint32; F1 int; F2 bool }
type GenS8 struct { F0 string; F1 string; F2 uint }
type GenS9 struct { F0 uint16 }
type GenS10 struct { F0 bool; F1 string }
type GenS11 struct { F0 uint; F1 uint64; F2 int64 }
type GenS12 struct { F0 uint32 }
func genFn0() string {
    if false {
        var v4 int8 = -128
        v5 := []int{137, 9223372036854775807, 2147483647, 448}
        _ = v4
        _ = v5
    } else {
        v6 := []int{606, 1}
        _ = v6
    }
    for fi7 := 0; fi7 < 3; fi7++ {
        switch "" {
        case "\n\t":
            for fi8 := 0; fi8 < 16; fi8++ {
                v9 := []int{}
                var v10 uint32 = 42
                var v11 int32 = -1
                _ = fi8
                _ = v10
                _ = v11
                _ = v9
            }
        case "a":
            fmt.Println(fi7)
        }
        _ = fi7
    }
    var v12 string = "hello"
    _ = v12
    return "hello"
}
func genFn1(p0 int, p1 uint8) uint8 {
    v13 := []int{(p0 >> ((42) & 63)), 32768, p0, 9223372036854775807}
    var v14 int16 = 1000
    if ((1.5 + 3.14159) < (1.5 + 3.14159)) {
        v14 -= v14
        fmt.Println(v14)
        if true {
            v13 = append(v13, p0)
            v13 = append(v13, (-(p0)))
        } else {
            v15 := GenS1{F0: (("hello" + "\n\t") + (("\n\t" + "\n\t") + ("hello" + "\n\t"))), F1: 1}
            if (int8(p0) > 0) {
                v16 := GenS1{F0: "", F1: 948}
                if false {
                    var v17 uint = 65535
                    var v18 string = ("\n\t" + "hello")
                    _ = v17
                    _ = v18
                }
                v13 = append(v13, 127)
                _ = v16
            }
            _ = v15
        }
    }
    p1 = (p1 - 244)
    _ = p0
    _ = p1
    _ = v13
    _ = v14
    return 128
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[int]bool)
    v3 := make(map[string]string)
    v19 := GenS2{F0: 936, F1: 42}
    v20 := GenS3{F0: 9}
    var v21 int8 = -1
    fmt.Println(v21)
    v22 := GenS4{F0: 746, F1: 32767, F2: 1}
    var v23 bool = (2 >= 2)
    v24 := GenS2{F0: 255, F1: (-(v21))}
    switch 127 {
    case -9223372036854775808:
        v21 = (v21 >> ((1) & 63))
    case -1:
        v25 := GenS5{F0: 127}
        _ = v25
    case 9223372036854775807:
        for fi26 := 0; fi26 < 19; fi26++ {
            var v27 bool = v23
            v28 := []int{fi26, (fi26 / (((fi26 + (-(fi26)))) | 1))}
            _ = fi26
            _ = v27
            _ = v28
        }
    }
    v29 := []int{-195, -1}
    var v30 int = 128
    v31 := GenS2{F0: 42, F1: 1}
    v30 = v30
    if ((1000 == 1000) && v23) {
        var v32 float32 = float32((1.5 - (3.14159 * float64(-2.25))))
        _ = v32
    } else {
        if v23 {
            v33 := GenS1{F0: ("a" + ("\n\t" + "a")), F1: 2}
            _ = v33
        } else {
            for fi34 := 0; fi34 < 4; fi34++ {
                switch "" {
                case "\n\t":
                    for fi35 := 0; fi35 < 19; fi35++ {
                        v23 = (1 == -827)
                        v36 := GenS2{F0: 4294967295, F1: (v21 / (((v21 - (-(v21)))) | 1))}
                        if v23 {
                            if (!true) {
                                fmt.Println(v21)
                                var v37 bool = (false && (128 == 232))
                                _ = v37
                            } else {
                                if v23 {
                                    v30 += (fi35 << ((15) & 63))
                                } else {
                                    switch "foo bar" {
                                    case "hello":
                                        delete(v1, (v30 << ((18446744073709551615) & 63)))
                                    default:
                                        fmt.Println(v23)
                                    }
                                }
                                v38 := v2[(fi34 ^ v30)]
                                _ = v38
                            }
                        } else {
                            fmt.Println(v23, v30, v21)
                            if true {
                                var v39 int32 = 276
                                _ = v39
                            } else {
                                if (2 < 255) {
                                    switch ("\n\t" + ("\n\t" + "a")) {
                                    case "hello":
                                        switch (fi35 % ((fi35) | 1)) {
                                        case 127:
                                            if v23 {
                                                var v40 uint16 = 42
                                                v41 := GenS6{F0: "foo bar", F1: 1}
                                                var v42 int64 = -601
                                                _ = v40
                                                _ = v41
                                                _ = v42
                                            } else {
                                                var v43 float32 = (((0.1 - 1.5) * -2.25) + 0.0)
                                                _ = v43
                                            }
                                        case 1:
                                            switch 42 {
                                            case 1:
                                                var v44 int32 = 740
                                                _ = v44
                                            case 0:
                                                v2[fi35] = (!((-830 == -1) || (!(1e10 != 0.1))))
                                            case -272:
                                                v21 += v21
                                            default:
                                                fmt.Println(v21, v30, fi35)
                                            }
                                        case -83:
                                            var v45 float64 = -2.25
                                            _ = v45
                                        default:
                                            v23 = true
                                        }
                                    }
                                    switch -1 {
                                    case 32767:
                                        v46 := GenS3{F0: 42}
                                        _ = v46
                                    }
                                    switch (fi35 ^ fi34) {
                                    case 42:
                                        fmt.Println(fi35)
                                    }
                                }
                                switch (^(fi35)) {
                                case 127:
                                    switch (fi34 + v30) {
                                    case 2147483647:
                                        switch 463 {
                                        case 127:
                                            var v47 string = ((("\n\t" + "\n\t") + "") + "hello")
                                            _ = v47
                                        case 32767:
                                            var v48 uint64 = 4294967295
                                            _ = v48
                                        case 128:
                                            switch fi35 {
                                            case 0:
                                                delete(v3, "foo bar")
                                            case -2147483648:
                                                var v49 uint32 = 639
                                                _ = v49
                                            default:
                                                var v50 string = "hello"
                                                _ = v50
                                            }
                                        }
                                    }
                                    var v51 int32 = 872
                                    if (2 >= 2) {
                                        switch ("hello" + "a") {
                                        case "\n\t":
                                            v52 := GenS7{F0: 2147483647, F1: v30, F2: (!false)}
                                            _ = v52
                                        case "":
                                            var v53 int8 = (-(v21))
                                            _ = v53
                                        default:
                                            var v54 uint16 = 181
                                            _ = v54
                                        }
                                    } else {
                                        v55 := GenS8{F0: "", F1: ("a" + ("hello" + "")), F2: 0}
                                        var v56 float64 = ((3.14159 + 0.1) - ((-2.25 * 3.14159) * (3.14159 / 1.5)))
                                        v57 := v2[v30]
                                        var v58 uint32 = 42
                                        _ = v55
                                        _ = v56
                                        _ = v57
                                        _ = v58
                                    }
                                    v30 -= (fi34 << ((65535) & 63))
                                    _ = v51
                                case 32767:
                                    var v59 uint8 = 42
                                    _ = v59
                                case 1:
                                    if (-134 < 32768) {
                                        fmt.Println(v21)
                                    } else {
                                        var v60 int32 = 715
                                        v23 = v23
                                        _ = v60
                                    }
                                }
                            }
                        }
                        _ = fi35
                        _ = v36
                    }
                }
                for fi61 := 0; fi61 < 14; fi61++ {
                    if v23 {
                        if (!((1.5 + 3.14159) > (-2.25 * 0.0))) {
                            fmt.Println(v21, fi61)
                        }
                        if (187 > genFn1((fi61 - fi61), 128)) {
                            v23 = (1e10 < (-2.25 / 1.0))
                            v62 := GenS8{F0: "\n\t", F1: "foo bar", F2: 65535}
                            fmt.Println(v30, fi34)
                            v63 := GenS6{F0: "\n\t", F1: -9223372036854775808}
                            if len(v29) >= 2 {
                                v29 = v29[1:len(v29)]
                            }
                            _ = v62
                            _ = v63
                        } else {
                            v64 := v2[(v30 << ((1) & 63))]
                            v65 := GenS5{F0: v21}
                            fmt.Println(v64, v64)
                            _ = v64
                            _ = v65
                        }
                        fmt.Println(v21, fi61, v30)
                    }
                    var v66 uint16 = 255
                    _ = fi61
                    _ = v66
                }
                _ = fi34
            }
            switch v30 {
            case 32767:
                var v67 uint = 253
                _ = v67
            case 127:
                switch (v30 ^ (v30 / ((v30) | 1))) {
                case -595:
                    var v68 float32 = (0.1 - ((1.5 / float32(float64(1.5))) + float32(float64(float32((0.0 / (float64((((1e10 + 1e10) / 1e10) + ((((-2.25 + 1.5) + 1e10) + (3.14159 * 0.0)) / 1.0))) + ((float64(0.1) + (-2.25 + 0.1)) - 1e10))))))))
                    _ = v68
                case 42:
                    var v69 uint16 = 255
                    _ = v69
                }
            case 463:
                if (!((!v23) || (v23 && ((v23 && false) || (v21 == 42))))) {
                    for fi70 := 0; fi70 < 9; fi70++ {
                        for fi71 := 0; fi71 < 10; fi71++ {
                            delete(v1, fi70)
                            switch (fi71 / ((-9223372036854775808) | 1)) {
                            case -467:
                                v72 := v1[fi71]
                                _ = v72
                            case -822:
                                var v73 int8 = v21
                                _ = v73
                            default:
                                if false {
                                    v74 := GenS9{F0: 255}
                                    v21 += -1
                                    if v23 {
                                        var v75 bool = v23
                                        var v76 int8 = v21
                                        _ = v75
                                        _ = v76
                                    } else {
                                        v77 := GenS10{F0: false, F1: "hello"}
                                        _ = v77
                                    }
                                    _ = v74
                                } else {
                                    switch (((("foo bar" + "\n\t") + "foo bar") + "foo bar") + "") {
                                    case "foo bar":
                                        var v78 bool = true
                                        _ = v78
                                    case "a":
                                        switch ("" + "") {
                                        case "a":
                                            var v79 uint64 = 18446744073709551615
                                            _ = v79
                                        case "":
                                            var v80 float64 = -2.25
                                            _ = v80
                                        default:
                                            v23 = true
                                            var v81 float32 = 0.0
                                            if (v23 && (!true)) {
                                                switch "" {
                                                case "":
                                                    var v82 uint = 4294967295
                                                    _ = v82
                                                case "hello":
                                                    switch -2147483648 {
                                                    case 32768:
                                                        v30 -= fi70
                                                    }
                                                case "foo bar":
                                                    v83 := GenS10{F0: false, F1: ("hello" + "\n\t")}
                                                    _ = v83
                                                default:
                                                    v81 = v81
                                                }
                                                fmt.Println(fi71, v23, v21)
                                                var v84 uint64 = 1
                                                _ = v84
                                            }
                                            _ = v81
                                        }
                                    case "\n\t":
                                        if v23 {
                                            var v85 uint32 = 2147483647
                                            _ = v85
                                        } else {
                                            var v86 int = 114
                                            var v87 uint8 = 42
                                            _ = v86
                                            _ = v87
                                        }
                                    default:
                                        v30 -= -9223372036854775808
                                        var v88 int32 = 1
                                        v2[4294967295] = true
                                        _ = v88
                                    }
                                }
                                fmt.Println(v30)
                            }
                            _ = fi71
                        }
                        v21 += -128
                        v21 = (v21 / ((v21) | 1))
                        _ = fi70
                    }
                }
            default:
                v30 = (v30 * 0)
            }
            var v89 int32 = -1
            _ = v89
        }
        fmt.Println(v21, v30, v21)
        var v90 uint64 = 947
        _ = v90
    }
    _ = len(v29)
    for fi91 := 0; fi91 < 4; fi91++ {
        for fi92 := 0; fi92 < 3; fi92++ {
            v3["a"] = ("" + ("hello" + "hello"))
            var v93 float32 = 1.5
            var v94 uint = 42
            switch "a" {
            case "":
                var v95 int8 = 100
                switch fi92 {
                case 9223372036854775807:
                    v23 = v23
                default:
                    if len(v29) >= 2 {
                        v29 = v29[1:len(v29)]
                    }
                }
                _ = v95
            default:
                v96 := GenS9{F0: 32767}
                _ = v96
            }
            _ = fi92
            _ = v93
            _ = v94
        }
        var v97 int32 = -1
        if true {
            fmt.Println(v30, v30, v30)
        } else {
            for fi98 := 0; fi98 < 5; fi98++ {
                var v99 uint8 = 2
                _ = fi98
                _ = v99
            }
            switch "hello" {
            case "":
                for fi100 := 0; fi100 < 20; fi100++ {
                    if v23 {
                        switch (fi91 & (-(fi100))) {
                        case -866:
                            v23 = true
                        case 32768:
                            var v101 int64 = 802
                            _ = v101
                        case -2147483648:
                            v102 := GenS4{F0: 1, F1: fi100, F2: -32768}
                            _ = v102
                        }
                        if v23 {
                            if (0 <= 42) {
                                fmt.Println(fi100)
                                var v103 string = "foo bar"
                                var v104 int64 = 9223372036854775807
                                var v105 string = v103
                                _ = v103
                                _ = v104
                                _ = v105
                            } else {
                                v97 -= (v97 >> ((255) & 63))
                                var v106 uint64 = 694
                                _ = v106
                            }
                            _ = len(v3)
                            switch (v30 | v30) {
                            case 398:
                                var v107 uint32 = 4294967295
                                _ = v107
                            case -1:
                                fmt.Println(fi91, v97)
                            case 32768:
                                switch (v30 * (fi100 << ((18446744073709551615) & 63))) {
                                case 128:
                                    switch "" {
                                    case "\n\t":
                                        if len(v29) >= 2 {
                                            v29 = v29[1:len(v29)]
                                        }
                                    case "foo bar":
                                        var v108 uint64 = 270
                                        _ = v108
                                    default:
                                        fmt.Println(v23)
                                    }
                                    v30 = fi91
                                }
                            }
                        }
                    }
                    if (53 != (v21 - int8(v97))) {
                        var v109 uint32 = 65535
                        v97 -= 0
                        if (65535 > (^(v109))) {
                            v109 = v109
                            v110 := GenS11{F0: 751, F1: 18446744073709551615, F2: -9223372036854775808}
                            switch fi100 {
                            case -91:
                                v111 := GenS9{F0: 42}
                                _ = v111
                            case 672:
                                var v112 float64 = 3.14159
                                _ = v112
                            }
                            _ = v110
                        }
                        _ = v109
                    }
                    _ = fi100
                }
            case "foo bar":
                if v23 {
                    var v113 uint = 778
                    _ = v113
                } else {
                    var v114 int64 = 9223372036854775807
                    var v115 uint = 4294967295
                    v115 -= 87
                    _ = v114
                    _ = v115
                }
            default:
                for fi116 := 0; fi116 < 13; fi116++ {
                    v117 := GenS1{F0: "", F1: 2147483648}
                    v23 = (!v23)
                    _ = fi116
                    _ = v117
                }
            }
        }
        _ = fi91
        _ = v97
    }
    for fi118 := 0; fi118 < 13; fi118++ {
        v119 := GenS12{F0: 520}
        _ = fi118
        _ = v119
    }
    var v120 bool = v23
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v120:", v120)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    _ = v1
    _ = v120
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v29
    _ = v3
    _ = v30
    _ = v31
}
