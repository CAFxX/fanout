package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint64; F1 bool; F2 uint8 }
type GenS5 struct { F0 uint8; F1 uint32; F2 int64 }
type GenS6 struct { F0 uint64; F1 uint32; F2 int8 }
type GenS7 struct { F0 bool; F1 uint8; F2 string }
type GenS8 struct { F0 uint64; F1 uint64 }
type GenS9 struct { F0 uint16 }
type GenS10 struct { F0 int; F1 int8 }
type GenS11 struct { F0 int8; F1 int; F2 int32 }
type GenS12 struct { F0 uint64; F1 string }
type GenS13 struct { F0 bool; F1 uint32; F2 uint8 }
func genFn0() uint {
    var v5 int64 = 0
    v5 -= v5
    v5 = v5
    _ = v5
    return 1
}
func genFn1(p0 string, p1 int64) uint64 {
    if true {
        p1 -= (-(p1))
    } else {
        v6 := []int{128, 32767, 128}
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        var v7 uint16 = 1
        _ = v6
        _ = v7
    }
    var v8 uint32 = 1
    v9 := GenS1{F0: 42, F1: p1}
    for fi10 := 0; fi10 < 20; fi10++ {
        if ((1.5 / 1e10) <= 3.14159) {
            var v11 float32 = 1e10
            _ = v11
        } else {
            if (!false) {
                for fi12 := 0; fi12 < 17; fi12++ {
                    var v13 float32 = float32(3.14159)
                    _ = fi12
                    _ = v13
                }
                var v14 bool = (false && false)
                var v15 int64 = p1
                _ = v14
                _ = v15
            }
        }
        for fi16 := 0; fi16 < 16; fi16++ {
            var v17 uint8 = 240
            if (225 > -32768) {
                v18 := []int{fi10, fi16, -1}
                var v19 bool = (v17 >= v17)
                _ = v18
                _ = v19
            } else {
                p0 = (p0 + "a")
            }
            switch (fi10 * fi10) {
            case 127:
                var v20 uint64 = 18446744073709551615
                _ = v20
            case 308:
                v21 := GenS1{F0: 127, F1: -15}
                _ = v21
            case 746:
                fmt.Println(v17, p1)
            default:
                var v22 uint = 18446744073709551615
                _ = v22
            }
            var v23 uint8 = v17
            switch 4294967295 {
            case 742:
                if ((3.14159 / 1e10) != (0.1 * ((1.5 - -2.25) - 0.0))) {
                    switch fi10 {
                    case 32768:
                        v23 = (v23 & (^(v23)))
                    case 42:
                        var v24 float64 = (-2.25 * (0.1 / (float64((3.14159 - -2.25)) + 1.5)))
                        if ((v23 << ((1) & 63)) != v23) {
                            var v25 uint16 = 614
                            var v26 string = ""
                            v26 = v26
                            _ = v25
                            _ = v26
                        }
                        v27 := GenS1{F0: 127, F1: (p1 / ((42) | 1))}
                        var v28 uint32 = (v8 - v8)
                        _ = v24
                        _ = v27
                        _ = v28
                    default:
                        switch fi16 {
                        case 127:
                            v29 := GenS2{F0: 65535}
                            _ = v29
                        case -320:
                            fmt.Println(fi16, p0)
                        case -590:
                            var v30 uint16 = 252
                            _ = v30
                        }
                    }
                } else {
                    v31 := GenS1{F0: -1, F1: p1}
                    switch 544 {
                    case 473:
                        v32 := GenS3{F0: ("hello" + "\n\t")}
                        p0 = "a"
                        _ = v32
                    case -620:
                        if (1e10 <= (0.0 - 1.5)) {
                            if (65535 <= 2) {
                                switch (^(fi10)) {
                                case 0:
                                    v33 := GenS2{F0: 0}
                                    var v34 uint = uint(v8)
                                    v8 = (^(v8))
                                    switch p0 {
                                    case "a":
                                        switch fi10 {
                                        case 127:
                                            var v35 int32 = -2147483648
                                            _ = v35
                                        default:
                                            switch (^(fi10)) {
                                            case 42:
                                                if false {
                                                    if true {
                                                        switch fi16 {
                                                        case -9223372036854775808:
                                                            p1 = (p1 / ((p1) | 1))
                                                        case 9223372036854775807:
                                                            p1 += (p1 % ((0) | 1))
                                                        case 128:
                                                            var v36 int = fi16
                                                            _ = v36
                                                        }
                                                    }
                                                    p0 = p0
                                                } else {
                                                    if false {
                                                        if (32767 != 1000) {
                                                            v37 := []int{(fi10 & (fi16 ^ (fi10 << ((genFn0()) & 63))))}
                                                            _ = v37
                                                        }
                                                    } else {
                                                        fmt.Println(p0, p0)
                                                        switch fi16 {
                                                        case 2147483647:
                                                            switch (fi16 + (fi10 + fi16)) {
                                                            case 4294967295:
                                                                v8 -= 42
                                                            case 2147483647:
                                                                var v38 float32 = 1.5
                                                                _ = v38
                                                            default:
                                                                if (((0 >= 82) && (!(!(1 > 1)))) || true) {
                                                                    v39 := []int{fi16, (fi16 % ((fi10) | 1)), fi10}
                                                                    fmt.Println(v34, p1, v34)
                                                                    _ = v39
                                                                } else {
                                                                    p0 = p0
                                                                    if true {
                                                                        if false {
                                                                            v40 := GenS1{F0: 42, F1: -1}
                                                                            var v41 float64 = (-2.25 - 1.5)
                                                                            _ = v40
                                                                            _ = v41
                                                                        }
                                                                        if true {
                                                                            v42 := GenS2{F0: 468}
                                                                            var v43 string = (p0 + "hello")
                                                                            if (3.14159 > 1.5) {
                                                                                var v44 uint8 = (v23 * v17)
                                                                                _ = v44
                                                                            } else {
                                                                                if true {
                                                                                    fmt.Println(v17)
                                                                                    fmt.Println(p0, p0)
                                                                                    switch p0 {
                                                                                    case "hello":
                                                                                        if (false && true) {
                                                                                            var v45 float64 = 0.0
                                                                                            _ = v45
                                                                                        }
                                                                                    case "a":
                                                                                        v34 = (v34 - v34)
                                                                                    case "foo bar":
                                                                                        if ((42 == v17) && (!false)) {
                                                                                            var v46 uint8 = (^(v23))
                                                                                            _ = v46
                                                                                        }
                                                                                    }
                                                                                }
                                                                                var v47 int64 = (p1 << ((v34) & 63))
                                                                                _ = v47
                                                                            }
                                                                            v48 := []int{(fi10 << ((genFn0()) & 63)), fi10, (fi16 - fi10), (fi10 << ((2) & 63))}
                                                                            var v49 int8 = 95
                                                                            _ = v42
                                                                            _ = v43
                                                                            _ = v48
                                                                            _ = v49
                                                                        }
                                                                        v50 := GenS4{F0: 31, F1: (!((!(845 < 65535)) || false)), F2: (v17 << ((genFn0()) & 63))}
                                                                        _ = v50
                                                                    }
                                                                    v51 := []int{(^(fi10)), fi10}
                                                                    _ = v51
                                                                }
                                                            }
                                                        case 557:
                                                            var v52 float32 = 3.14159
                                                            var v53 uint = v34
                                                            _ = v52
                                                            _ = v53
                                                        }
                                                        var v54 uint32 = 0
                                                        v55 := []int{(fi16 >> (((v34 - (v34 | genFn0()))) & 63)), 128, fi10}
                                                        v8 = v54
                                                        _ = v54
                                                        _ = v55
                                                    }
                                                    var v56 int16 = -1
                                                    var v57 bool = (1.5 >= 3.14159)
                                                    _ = v56
                                                    _ = v57
                                                }
                                            }
                                        }
                                    }
                                    fmt.Println(v34, v17)
                                    _ = v33
                                    _ = v34
                                default:
                                    v58 := GenS4{F0: 42, F1: true, F2: (v17 % (((v23 & v17)) | 1))}
                                    _ = v58
                                }
                            } else {
                                v59 := GenS5{F0: (v17 + (v17 * (v23 ^ (^(v17))))), F1: (v8 + v8), F2: (p1 & (p1 | 42))}
                                v60 := []int{}
                                if false {
                                    v61 := GenS6{F0: 1, F1: v8, F2: 1}
                                    if (((1 >= 2) || true) || false) {
                                        if false {
                                            fmt.Println(v17)
                                            switch 9223372036854775807 {
                                            case 0:
                                                v62 := GenS6{F0: 0, F1: 479, F2: -91}
                                                _ = v62
                                            case 32768:
                                                var v63 uint = 861
                                                _ = v63
                                            case 42:
                                                if false {
                                                    switch fi16 {
                                                    case 824:
                                                        v64 := GenS7{F0: true, F1: 127, F2: ("hello" + "hello")}
                                                        _ = v64
                                                    case 128:
                                                        var v65 uint16 = 635
                                                        _ = v65
                                                    default:
                                                        if true {
                                                            switch p0 {
                                                            case "foo bar":
                                                                p1 = (p1 << ((genFn0()) & 63))
                                                            case "hello":
                                                                var v66 int64 = -1
                                                                _ = v66
                                                            }
                                                            var v67 uint32 = uint32(v23)
                                                            _ = v67
                                                        } else {
                                                            var v68 int64 = p1
                                                            v69 := GenS8{F0: 1, F1: 773}
                                                            switch (fi16 ^ 127) {
                                                            case 128:
                                                                var v70 uint8 = (v17 << ((42) & 63))
                                                                _ = v70
                                                            case -272:
                                                                fmt.Println(v8)
                                                            case 32768:
                                                                v71 := GenS9{F0: 42}
                                                                _ = v71
                                                            default:
                                                                v8 += (^(v8))
                                                            }
                                                            _ = v68
                                                            _ = v69
                                                        }
                                                    }
                                                    p1 = (-(p1))
                                                }
                                                fmt.Println(p1)
                                                v72 := GenS1{F0: 127, F1: (p1 % ((p1) | 1))}
                                                _ = v72
                                            default:
                                                fmt.Println(fi16)
                                            }
                                        }
                                    }
                                    _ = v61
                                } else {
                                    if ((p1 != p1) && ((v17 & v17) <= v17)) {
                                        v73 := GenS10{F0: fi10, F1: -42}
                                        switch p0 {
                                        case "foo bar":
                                            v23 = v17
                                            switch fi10 {
                                            case 4294967295:
                                                switch (p0 + (p0 + "foo bar")) {
                                                case "":
                                                    v74 := GenS10{F0: (fi16 ^ (fi16 * -96)), F1: 117}
                                                    _ = v74
                                                case "hello":
                                                    v17 -= 2
                                                case "a":
                                                    v17 -= v23
                                                }
                                            }
                                            v17 += v23
                                            if (!(("" + (p0 + (p0 + p0))) == p0)) {
                                                fmt.Println(v8, fi16, fi16)
                                                var v75 uint64 = 2
                                                _ = v75
                                            } else {
                                                v76 := GenS10{F0: fi16, F1: 1}
                                                var v77 uint8 = (v17 ^ 1)
                                                _ = v76
                                                _ = v77
                                            }
                                        default:
                                            var v78 uint = genFn0()
                                            _ = v78
                                        }
                                        switch fi16 {
                                        case 4294967295:
                                            p0 = "foo bar"
                                        }
                                        _ = v73
                                    }
                                }
                                _ = v59
                                _ = v60
                            }
                            p1 = (p1 & (p1 / (((p1 / (((p1 & 42)) | 1))) | 1)))
                            var v79 float64 = float64(0.1)
                            _ = v79
                        }
                    default:
                        var v80 int8 = -116
                        _ = v80
                    }
                    _ = v31
                }
            case 2147483647:
                v81 := []int{fi16}
                _ = v81
            case -558:
                var v82 int8 = 100
                _ = v82
            default:
                v23 -= (v23 >> ((42) & 63))
            }
            _ = fi16
            _ = v17
            _ = v23
        }
        var v83 int = (fi10 << ((781) & 63))
        _ = fi10
        _ = v83
    }
    var v84 uint64 = 0
    _ = p0
    _ = p1
    _ = v8
    _ = v84
    _ = v9
    return (v84 >> ((218) & 63))
}

func main() {
    v1 := make(map[int]uint32)
    var v2 float64 = 0.1
    v3 := make(map[string]string)
    var v4 int8 = -36
    for fi85 := 0; fi85 < 18; fi85++ {
        fmt.Println(v4, fi85)
        v4 += v4
        for fi86 := 0; fi86 < 14; fi86++ {
            v4 += v4
            var v87 float32 = 1e10
            _ = fi86
            _ = v87
        }
        _ = fi85
    }
    fmt.Println(v4)
    v4 += (v4 % (((v4 - (v4 ^ v4))) | 1))
    v4 += (v4 << ((255) & 63))
    switch "\n\t" {
    case "hello":
        for fi88 := 0; fi88 < 16; fi88++ {
            if (v4 == v4) {
                v89 := GenS10{F0: (-(fi88)), F1: -17}
                for fi90 := 0; fi90 < 12; fi90++ {
                    v4 += (-(v4))
                    _ = fi90
                }
                var v91 uint64 = 2
                _ = v89
                _ = v91
            } else {
                v2 += float64(0.1)
                v92 := GenS8{F0: 2, F1: 0}
                _ = v92
            }
            _ = fi88
        }
    case "":
        v2 = (-2.25 - ((-2.25 * float64(-2.25)) * ((v2 * float64(0.0)) - float64(-2.25))))
    default:
        switch ("\n\t" + "") {
        case "\n\t":
            if (-357 > 32767) {
                delete(v3, ("hello" + "foo bar"))
                for fi93 := 0; fi93 < 2; fi93++ {
                    var v94 string = "\n\t"
                    switch "a" {
                    case "foo bar":
                        v2 = 1.5
                    case "hello":
                        v95 := GenS1{F0: (v4 >> ((genFn0()) & 63)), F1: 0}
                        _ = v95
                    default:
                        var v96 float32 = float32((v2 - v2))
                        _ = v96
                    }
                    _ = fi93
                    _ = v94
                }
            } else {
                v4 = v4
                switch 127 {
                case 5:
                    for fi97 := 0; fi97 < 10; fi97++ {
                        switch (("foo bar" + "foo bar") + "") {
                        case "\n\t":
                            if (-677 > 0) {
                                if false {
                                    var v98 uint = 1
                                    var v99 bool = (!(-1 != 0))
                                    if (0.1 < 1e10) {
                                        fmt.Println(v2, fi97, v99)
                                        v100 := []int{-9223372036854775808, (fi97 & -1), fi97, 9223372036854775807}
                                        _ = v100
                                    }
                                    _ = v98
                                    _ = v99
                                }
                                v2 = v2
                            }
                        case "hello":
                            v4 -= v4
                        case "a":
                            var v101 int64 = -197
                            _ = v101
                        }
                        v102 := GenS9{F0: 42}
                        _ = fi97
                        _ = v102
                    }
                case 341:
                    if false {
                        for fi103 := 0; fi103 < 13; fi103++ {
                            v4 -= v4
                            v104 := GenS11{F0: 127, F1: (fi103 % ((32768) | 1)), F2: 0}
                            _ = fi103
                            _ = v104
                        }
                    } else {
                        fmt.Println(v4, v4)
                        for fi105 := 0; fi105 < 18; fi105++ {
                            if (127 == 255) {
                                switch "\n\t" {
                                case "\n\t":
                                    _ = len(v1)
                                case "foo bar":
                                    for fi106 := 0; fi106 < 3; fi106++ {
                                        v2 += ((v2 + (float64((1.5 - (-2.25 - (1.5 * (-2.25 + 3.14159))))) - v2)) + -2.25)
                                        v2 = (v2 / v2)
                                        _ = fi106
                                    }
                                default:
                                    fmt.Println(fi105)
                                }
                                _ = len(v3)
                            }
                            _ = fi105
                        }
                    }
                }
            }
        case "":
            switch 32768 {
            case 0:
                v2 -= float64(1e10)
            case 127:
                var v107 uint8 = 128
                _ = v107
            }
        default:
            var v108 uint16 = 0
            _ = v108
        }
    }
    var v109 uint32 = 42
    for fi110 := 0; fi110 < 7; fi110++ {
        var v111 int16 = -1
        var v112 uint64 = genFn1("foo bar", 0)
        _ = fi110
        _ = v111
        _ = v112
    }
    v2 = v2
    v113 := []int{-2147483648, -705, -523, 128}
    v109 -= v109
    if (4294967295 == 214) {
        var v114 int16 = 1000
        _ = v114
    } else {
        v115 := GenS12{F0: genFn1("a", -830), F1: "hello"}
        fmt.Println(v109, v109)
        var v116 uint16 = 0
        _ = v115
        _ = v116
    }
    var v117 int16 = 42
    var v118 uint = 0
    v119 := GenS13{F0: (uint16(v109) < 807), F1: v109, F2: 85}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v109:", v109)
    fmt.Println("slicelen v113:", len(v113), v113)
    fmt.Println("v117:", v117)
    fmt.Println("v118:", v118)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v109
    _ = v113
    _ = v117
    _ = v118
    _ = v119
    _ = v2
    _ = v3
    _ = v4
}
