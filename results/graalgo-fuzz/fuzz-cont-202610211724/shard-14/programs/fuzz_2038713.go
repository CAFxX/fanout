package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint16; F1 uint32 }
type GenS4 struct { F0 uint32; F1 bool }
type GenS5 struct { F0 uint64; F1 int8; F2 uint16 }
type GenS6 struct { F0 uint64; F1 string; F2 uint64 }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 uint32; F1 int }
type GenS9 struct { F0 int8; F1 int32 }
type GenS10 struct { F0 bool; F1 uint64; F2 uint32 }
type GenS11 struct { F0 string }
func genFn0(p0 int16, p1 uint32) string {
    for fi4 := 0; fi4 < 6; fi4++ {
        p0 -= p0
        v5 := GenS1{F0: "hello", F1: int16(fi4)}
        fmt.Println(p0, fi4)
        _ = fi4
        _ = v5
    }
    for fi6 := 0; fi6 < 10; fi6++ {
        p0 = p0
        _ = fi6
    }
    if (32767 == 9223372036854775807) {
        switch (("hello" + ("a" + "foo bar")) + "hello") {
        case "\n\t":
            p0 -= p0
        case "a":
            v7 := []int{-605}
            _ = v7
        }
        if false {
            p1 = (p1 << ((19) & 63))
        } else {
            v8 := []int{}
            fmt.Println(p0)
            _ = v8
        }
    } else {
        switch 4294967295 {
        case 128:
            if true {
                switch 9223372036854775807 {
                case -1:
                    p1 -= 42
                }
            } else {
                var v9 uint8 = 129
                switch -1 {
                case 1:
                    v10 := []int{128}
                    _ = v10
                }
                p1 -= (p1 >> ((65535) & 63))
                _ = v9
            }
        }
    }
    switch 728 {
    case 127:
        p0 += 32767
    case -880:
        for fi11 := 0; fi11 < 6; fi11++ {
            v12 := []int{}
            if ((p1 >> ((0) & 63)) >= (^(p1))) {
                for fi13 := 0; fi13 < 7; fi13++ {
                    v14 := GenS1{F0: ((("" + "hello") + "hello") + "a"), F1: p0}
                    var v15 uint8 = 199
                    switch fi13 {
                    case 42:
                        if false {
                            v16 := GenS1{F0: "hello", F1: int16(p1)}
                            var v17 bool = false
                            var v18 uint8 = v15
                            _ = v16
                            _ = v17
                            _ = v18
                        }
                    case 32768:
                        p1 = p1
                    case 9223372036854775807:
                        v15 = v15
                    }
                    _ = fi13
                    _ = v14
                    _ = v15
                }
            }
            var v19 int32 = -2147483648
            _ = fi11
            _ = v12
            _ = v19
        }
    case -657:
        v20 := GenS1{F0: "hello", F1: p0}
        _ = v20
    default:
        p1 -= p1
    }
    v21 := []int{}
    _ = p0
    _ = p1
    _ = v21
    return ((("\n\t" + "\n\t") + "hello") + ("a" + "a"))
}
func genFn1() uint16 {
    if (242 > 1) {
        var v22 uint64 = 451
        _ = v22
    } else {
        var v23 bool = false
        if (!(!v23)) {
            if (-32768 == 42) {
                v23 = false
                if (890 != 1) {
                    v24 := []int{}
                    if len(v24) > 0 {
                        v24[0] = 128
                    }
                    _ = v24
                }
                v25 := []int{}
                _ = v25
            }
        }
        v23 = (18446744073709551615 <= 2)
        _ = v23
    }
    v26 := []int{0}
    if (true || true) {
        if (1000 != -882) {
            _ = len(v26)
        }
        switch "foo bar" {
        case "hello":
            var v27 uint64 = 42
            _ = v27
        case "a":
            for fi28 := 0; fi28 < 7; fi28++ {
                var v29 string = "a"
                _ = fi28
                _ = v29
            }
        default:
            if (0.1 == -2.25) {
                for fi30 := 0; fi30 < 7; fi30++ {
                    fmt.Println(fi30)
                    if (!(int32(fi30) <= 2147483647)) {
                        v31 := GenS1{F0: ("" + ("a" + "\n\t")), F1: -32768}
                        _ = v31
                    }
                    _ = fi30
                }
                var v32 bool = (!(!false))
                var v33 int = 32767
                _ = v32
                _ = v33
            } else {
                v34 := GenS1{F0: ("foo bar" + "\n\t"), F1: 548}
                _ = v34
            }
        }
    }
    if (0 >= 42) {
        if false {
            v35 := GenS2{F0: 177}
            _ = v35
        } else {
            var v36 uint16 = 255
            if len(v26) >= 2 {
                v26 = v26[1:len(v26)]
            }
            switch "a" {
            case "":
                switch 1 {
                case 128:
                    var v37 uint64 = 18446744073709551615
                    _ = v37
                case 302:
                    switch "foo bar" {
                    case "a":
                        fmt.Println(v36)
                    case "hello":
                        if len(v26) >= 2 {
                            v26 = v26[1:len(v26)]
                        }
                    }
                case -582:
                    if ((127 > 128) || true) {
                        v36 = v36
                        v38 := GenS1{F0: "foo bar", F1: 1000}
                        _ = v38
                    }
                }
            case "a":
                if (1 != -1) {
                    var v39 int32 = -2147483648
                    switch (("hello" + ("a" + (("hello" + "hello") + "a"))) + "a") {
                    case "hello":
                        if false {
                            if false {
                                v40 := GenS3{F0: (v36 % ((1) | 1)), F1: 4294967295}
                                _ = v40
                            } else {
                                if len(v26) >= 2 {
                                    v26 = v26[1:len(v26)]
                                }
                                var v41 string = "\n\t"
                                var v42 uint32 = 65535
                                _ = v41
                                _ = v42
                            }
                            if (!((!true) && (234 != 0))) {
                                v43 := GenS3{F0: v36, F1: 1}
                                if (!false) {
                                    fmt.Println(v39, v39)
                                    v44 := GenS4{F0: 2147483648, F1: true}
                                    _ = v44
                                }
                                _ = v43
                            }
                            var v45 int64 = 42
                            _ = v45
                        } else {
                            var v46 int = -247
                            switch (v46 | (-(v46))) {
                            case 32768:
                                if len(v26) > 0 {
                                    v26[0] = 983
                                }
                            case 871:
                                if ((!(("foo bar" <= "\n\t") || ((!false) || (!(true && false))))) && (-54 == -63)) {
                                    var v47 int16 = -347
                                    _ = v47
                                } else {
                                    v36 = v36
                                    if false {
                                        var v48 uint16 = (v36 & (v36 >> ((18446744073709551615) & 63)))
                                        if ("" >= "") {
                                            v48 -= v48
                                            fmt.Println(v48)
                                            var v49 int64 = -300
                                            _ = v49
                                        } else {
                                            v26 = append(v26, 157)
                                            fmt.Println(v36)
                                            for fi50 := 0; fi50 < 19; fi50++ {
                                                var v51 uint8 = 1
                                                v46 = 32767
                                                v48 = 2
                                                _ = fi50
                                                _ = v51
                                            }
                                        }
                                        for fi52 := 0; fi52 < 2; fi52++ {
                                            v36 += 255
                                            v46 = 2147483647
                                            _ = fi52
                                        }
                                        _ = v48
                                    } else {
                                        v53 := GenS3{F0: (^(v36)), F1: 0}
                                        var v54 uint = 255
                                        var v55 int16 = 0
                                        _ = v53
                                        _ = v54
                                        _ = v55
                                    }
                                }
                            case 1:
                                for fi56 := 0; fi56 < 20; fi56++ {
                                    var v57 uint32 = 65535
                                    if (-9223372036854775808 != -1) {
                                        var v58 int32 = -1
                                        _ = v58
                                    } else {
                                        var v59 int8 = -9
                                        _ = v59
                                    }
                                    _ = fi56
                                    _ = v57
                                }
                            }
                            _ = v46
                        }
                    case "a":
                        v26 = append(v26, 128)
                    case "foo bar":
                        var v60 float64 = 1.5
                        _ = v60
                    default:
                        if true {
                            fmt.Println(v36, v36)
                            var v61 uint = 2
                            switch (("hello" + "foo bar") + ("hello" + "hello")) {
                            case "hello":
                                if len(v26) >= 2 {
                                    v26 = v26[1:len(v26)]
                                }
                            case "\n\t":
                                switch "a" {
                                case "hello":
                                    v62 := GenS3{F0: 1, F1: 512}
                                    _ = v62
                                default:
                                    for fi63 := 0; fi63 < 8; fi63++ {
                                        v64 := GenS5{F0: 4294967295, F1: 42, F2: v36}
                                        v36 = v36
                                        v61 = (^(v61))
                                        _ = fi63
                                        _ = v64
                                    }
                                }
                            case "":
                                for fi65 := 0; fi65 < 15; fi65++ {
                                    switch ("hello" + ((("a" + ("" + "foo bar")) + "a") + "foo bar")) {
                                    case "foo bar":
                                        switch ("\n\t" + "\n\t") {
                                        case "\n\t":
                                            if ((-2.25 + 1e10) >= 1.5) {
                                                switch 34 {
                                                case 128:
                                                    if (1 >= 2) {
                                                        if len(v26) > 0 {
                                                            v26[0] = (fi65 + (fi65 | -9223372036854775808))
                                                        }
                                                        _ = len(v26)
                                                    } else {
                                                        var v66 float64 = 1e10
                                                        if len(v26) > 0 {
                                                            v26[0] = fi65
                                                        }
                                                        _ = v66
                                                    }
                                                case 42:
                                                    v36 -= (v36 / (((v36 * 744)) | 1))
                                                default:
                                                    var v67 uint8 = 91
                                                    _ = v67
                                                }
                                            }
                                        case "":
                                            var v68 uint32 = 0
                                            _ = v68
                                        default:
                                            fmt.Println(v36, v36, v39)
                                        }
                                    case "hello":
                                        var v69 int8 = 7
                                        _ = v69
                                    default:
                                        v61 -= 2
                                    }
                                    v39 -= -2147483648
                                    v61 = 470
                                    _ = fi65
                                }
                            }
                            _ = v61
                        }
                    }
                    _ = v39
                }
            case "hello":
                for fi70 := 0; fi70 < 13; fi70++ {
                    for fi71 := 0; fi71 < 20; fi71++ {
                        var v72 bool = false
                        _ = fi71
                        _ = v72
                    }
                    _ = fi70
                }
            }
            _ = v36
        }
    } else {
        var v73 uint64 = 0
        var v74 bool = true
        _ = v73
        _ = v74
    }
    var v75 uint16 = 2
    _ = v26
    _ = v75
    return 0
}

func main() {
    v1 := make(map[int]int)
    v2 := []uint32{2147483647, 0, 1, 646}
    v3 := make(map[string]uint32)
    var v76 string = ""
    v77 := GenS6{F0: 0, F1: v76, F2: 2}
    var v78 uint8 = 2
    if false {
        var v79 uint16 = 1
        v78 = 128
        _ = v79
    } else {
        v80 := GenS6{F0: 657, F1: v76, F2: 1}
        _ = v80
    }
    if false {
        v78 = 180
    }
    if false {
        for fi81 := 0; fi81 < 15; fi81++ {
            for fi82 := 0; fi82 < 18; fi82++ {
                var v83 uint = 18446744073709551615
                v84 := v1[(fi81 * fi82)]
                v85 := GenS7{F0: (v78 >> ((4294967295) & 63))}
                var v86 float64 = 0.1
                _ = fi82
                _ = v83
                _ = v84
                _ = v85
                _ = v86
            }
            for fi87 := 0; fi87 < 20; fi87++ {
                var v88 float64 = 0.0
                v89 := GenS8{F0: 866, F1: fi87}
                fmt.Println(fi87)
                v90 := GenS4{F0: 2147483648, F1: true}
                _ = fi87
                _ = v88
                _ = v89
                _ = v90
            }
            switch v76 {
            case "foo bar":
                v76 = "foo bar"
            default:
                for fi91 := 0; fi91 < 19; fi91++ {
                    _ = len(v2)
                    _ = fi91
                }
            }
            _ = fi81
        }
        v1[42] = 886
        v92 := GenS1{F0: v76, F1: 1000}
        _ = v92
    }
    for fi93 := 0; fi93 < 13; fi93++ {
        v78 = v78
        if true {
            v94 := GenS5{F0: 298, F1: 112, F2: 1}
            _ = v94
        }
        _ = fi93
    }
    switch (("\n\t" + v76) + v76) {
    case "":
        for fi95 := 0; fi95 < 8; fi95++ {
            switch -9223372036854775808 {
            case 102:
                v96 := v1[(^(fi95))]
                _ = v96
            default:
                v76 = (("\n\t" + "foo bar") + v76)
            }
            for fi97 := 0; fi97 < 15; fi97++ {
                var v98 int8 = 127
                v76 = v76
                var v99 uint16 = 2
                _ = fi97
                _ = v98
                _ = v99
            }
            _ = fi95
        }
    case "hello":
        for fi100 := 0; fi100 < 15; fi100++ {
            if false {
                for fi101 := 0; fi101 < 5; fi101++ {
                    _ = len(v1)
                    _ = fi101
                }
                v102 := GenS9{F0: 90, F1: 941}
                _ = v102
            } else {
                v78 -= (v78 ^ (v78 | v78))
                v103 := GenS3{F0: 255, F1: 0}
                _ = v103
            }
            _ = fi100
        }
    case "a":
        for fi104 := 0; fi104 < 15; fi104++ {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            v78 = v78
            _ = fi104
        }
    }
    var v105 float64 = 0.0
    if (!(false && false)) {
        var v106 uint32 = 0
        var v107 float64 = 1e10
        var v108 uint16 = 783
        _ = v106
        _ = v107
        _ = v108
    }
    for fi109 := 0; fi109 < 9; fi109++ {
        for fi110 := 0; fi110 < 2; fi110++ {
            fmt.Println(fi109)
            v78 -= 1
            if false {
                _ = len(v2)
                v111 := GenS1{F0: v76, F1: 42}
                if true {
                    _ = len(v2)
                    v112 := GenS3{F0: 32767, F1: 1}
                    var v113 int = (fi110 / (((fi109 - 128)) | 1))
                    switch (v76 + "hello") {
                    case "\n\t":
                        v105 = (((v105 - -2.25) + v105) + 0.0)
                    default:
                        v113 -= fi110
                    }
                    _ = v112
                    _ = v113
                } else {
                    switch "\n\t" {
                    case "a":
                        v114 := GenS10{F0: true, F1: 876, F2: 65535}
                        _ = v114
                    case "hello":
                        v76 = v76
                    case "\n\t":
                        var v115 float64 = v105
                        _ = v115
                    default:
                        v105 = (v105 + (v105 + v105))
                    }
                    v116 := GenS10{F0: false, F1: 859, F2: 43}
                    _ = v116
                }
                _ = v111
            } else {
                if true {
                    v117 := GenS11{F0: "a"}
                    _ = v117
                } else {
                    v2 = append(v2, 735)
                    var v118 bool = (-9223372036854775808 >= -1)
                    _ = v118
                }
            }
            _ = fi110
        }
        var v119 string = v76
        _ = fi109
        _ = v119
    }
    fmt.Println(v78, v76)
    var v120 string = v76
    v121 := GenS2{F0: 0}
    for fi122 := 0; fi122 < 17; fi122++ {
        v123 := GenS11{F0: (v120 + (v76 + v76))}
        _ = fi122
        _ = v123
    }
    v105 += v105
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v105:", v105)
    fmt.Println("v120:", v120)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v76:", v76)
    fmt.Println("v78:", v78)
    _ = v1
    _ = v105
    _ = v120
    _ = v121
    _ = v2
    _ = v3
    _ = v76
    _ = v77
    _ = v78
}
