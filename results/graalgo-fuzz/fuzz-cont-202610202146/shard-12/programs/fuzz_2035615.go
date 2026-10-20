package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint; F1 int; F2 int8 }
type GenS3 struct { F0 uint8; F1 string }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 int16; F1 uint }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 uint; F1 uint16; F2 uint64 }
type GenS10 struct { F0 uint32 }
type GenS11 struct { F0 uint8; F1 uint; F2 bool }
type GenS12 struct { F0 int64 }
type GenS13 struct { F0 uint32 }
type GenS14 struct { F0 uint32; F1 int }
type GenS15 struct { F0 int16 }
type GenS16 struct { F0 int32; F1 int32 }
type GenS17 struct { F0 int8; F1 uint64 }
type GenS18 struct { F0 uint16 }
type GenS19 struct { F0 int64; F1 uint64 }
type GenS20 struct { F0 string; F1 int32; F2 uint }
type GenS21 struct { F0 uint32; F1 int64; F2 uint16 }
type GenS22 struct { F0 uint64 }
type GenS23 struct { F0 uint64 }
type GenS24 struct { F0 int16; F1 uint16; F2 int32 }
type GenS25 struct { F0 uint16 }
type GenS26 struct { F0 uint8; F1 int8; F2 int64 }
type GenS27 struct { F0 uint32; F1 uint32 }
type GenS28 struct { F0 uint8 }
func genFn0(p0 int16, p1 int8, p2 int8) uint8 {
    v4 := []int{4294967295, 1, 32767}
    for fi5 := 0; fi5 < 9; fi5++ {
        v4 = append(v4, 128)
        _ = fi5
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 128
}

func main() {
    v1 := []uint32{0, 759, 1}
    v2 := []string{"\n\t", "foo bar", "", "a"}
    var v3 uint32 = 2
    for fi6 := 0; fi6 < 14; fi6++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        switch fi6 {
        case -2147483648:
            v3 = 2147483647
        case -716:
            if (((!(true && false)) && true) && false) {
                var v7 uint64 = 790
                _ = v7
            } else {
                v3 = v3
                var v8 uint32 = v3
                _ = v8
            }
        case 127:
            var v9 uint16 = 42
            _ = v9
        }
        _ = fi6
    }
    var v10 int8 = 100
    v3 -= v3
    v3 = (^(v3))
    if true {
        if (3.14159 <= 3.14159) {
            if (!false) {
                v3 = (^(v3))
            }
            v3 -= v3
            v11 := GenS1{F0: 1000}
            _ = v11
        }
        if len(v2) > 0 {
            v2[0] = ("" + "hello")
        }
    } else {
        for fi12 := 0; fi12 < 10; fi12++ {
            v10 -= (v10 | v10)
            if true {
                var v13 string = ("a" + "a")
                _ = v13
            } else {
                for fi14 := 0; fi14 < 14; fi14++ {
                    switch "\n\t" {
                    case "":
                        v3 = 0
                        v15 := GenS1{F0: 32767}
                        _ = v15
                    case "\n\t":
                        var v16 uint64 = 2
                        _ = v16
                    default:
                        var v17 uint8 = 105
                        _ = v17
                    }
                    v1 = append(v1, v3)
                    _ = fi14
                }
            }
            var v18 uint8 = 182
            _ = fi12
            _ = v18
        }
    }
    fmt.Println(v10, v10)
    for fi19 := 0; fi19 < 20; fi19++ {
        fmt.Println(v10, fi19, v10)
        _ = fi19
    }
    v10 -= v10
    if (-128 == -1) {
        v10 = 0
    }
    v1 = append(v1, v3)
    for fi20 := 0; fi20 < 19; fi20++ {
        v21 := GenS1{F0: 0}
        if true {
            var v22 float64 = 0.1
            fmt.Println(v22)
            if true {
                v1 = append(v1, v3)
                var v23 uint8 = genFn0(-562, (v10 % ((127) | 1)), -108)
                for fi24 := 0; fi24 < 8; fi24++ {
                    if false {
                        if false {
                            var v25 uint64 = 4294967295
                            _ = v25
                        } else {
                            var v26 int64 = 42
                            switch ((("foo bar" + "") + "hello") + (("foo bar" + "\n\t") + "hello")) {
                            case "\n\t":
                                var v27 float32 = (1e10 + (-2.25 + 1.5))
                                _ = v27
                            }
                            var v28 float32 = 1e10
                            v29 := GenS1{F0: 535}
                            _ = v26
                            _ = v28
                            _ = v29
                        }
                        v30 := GenS1{F0: 42}
                        v3 += (v3 - (v3 >> ((1) & 63)))
                        var v31 uint = 65535
                        var v32 uint64 = 277
                        _ = v30
                        _ = v31
                        _ = v32
                    }
                    v3 -= (v3 | (^(v3)))
                    _ = fi24
                }
                _ = v23
            }
            _ = v22
        }
        _ = fi20
        _ = v21
    }
    v33 := GenS1{F0: 1}
    for fi34 := 0; fi34 < 19; fi34++ {
        v10 += (-(v10))
        if false {
            v35 := GenS2{F0: 88, F1: 127, F2: v10}
            switch fi34 {
            case 979:
                switch "a" {
                case "a":
                    switch fi34 {
                    case 1:
                        v10 = v10
                    default:
                        v36 := GenS2{F0: 1, F1: fi34, F2: (v10 << ((18446744073709551615) & 63))}
                        _ = v36
                    }
                case "hello":
                    for fi37 := 0; fi37 < 16; fi37++ {
                        switch ((("hello" + "a") + "\n\t") + ("foo bar" + ("\n\t" + "foo bar"))) {
                        case "foo bar":
                            fmt.Println(v3, fi34)
                        case "hello":
                            fmt.Println(fi37)
                        case "\n\t":
                            if (v3 > (v3 & 0)) {
                                var v38 float64 = -2.25
                                _ = v38
                            } else {
                                v10 = (v10 << ((0) & 63))
                                if len(v2) > 0 {
                                    v2[0] = (("" + ("" + "")) + "a")
                                }
                                v3 = 86
                                v3 += v3
                            }
                        }
                        switch "foo bar" {
                        case "foo bar":
                            var v39 uint32 = v3
                            _ = v39
                        case "hello":
                            v40 := GenS2{F0: 18446744073709551615, F1: (fi34 + 986), F2: v10}
                            if false {
                                var v41 float64 = (float64(1e10) / 1e10)
                                _ = v41
                            } else {
                                if ((false && (!false)) || true) {
                                    switch "a" {
                                    case "foo bar":
                                        v10 = 127
                                    case "":
                                        v42 := GenS3{F0: 0, F1: ""}
                                        _ = v42
                                    case "hello":
                                        v3 -= (v3 ^ v3)
                                    default:
                                        fmt.Println(fi37, fi34)
                                    }
                                    if true {
                                        switch ("\n\t" + "foo bar") {
                                        case "":
                                            switch ("hello" + "a") {
                                            case "hello":
                                                _ = len(v2)
                                            case "foo bar":
                                                v43 := GenS2{F0: 255, F1: 196, F2: v10}
                                                _ = v43
                                            case "\n\t":
                                                v44 := GenS4{F0: 0}
                                                _ = v44
                                            }
                                        case "a":
                                            if (fi34 != -1) {
                                                fmt.Println(v3, fi34, v3)
                                                switch ("a" + "a") {
                                                case "\n\t":
                                                    if (((!true) && true) && (!true)) {
                                                        v45 := GenS5{F0: v3}
                                                        fmt.Println(fi34, fi37)
                                                        _ = v45
                                                    } else {
                                                        var v46 uint64 = 4294967295
                                                        _ = v46
                                                    }
                                                case "foo bar":
                                                    var v47 bool = (!((fi34 + fi37) == fi37))
                                                    _ = v47
                                                case "a":
                                                    if (!((true && true) && (("a" + ("hello" + "\n\t")) <= ("hello" + "a")))) {
                                                        var v48 uint8 = 127
                                                        var v49 string = "foo bar"
                                                        _ = v48
                                                        _ = v49
                                                    } else {
                                                        var v50 uint32 = 4294967295
                                                        if ((-9223372036854775808 == 697) || false) {
                                                            var v51 int = fi37
                                                            v51 -= fi34
                                                            _ = v51
                                                        } else {
                                                            fmt.Println(fi34, v10, v10)
                                                            fmt.Println(v10, fi34, v3)
                                                        }
                                                        _ = v50
                                                    }
                                                default:
                                                    if (0 != -148) {
                                                        v3 -= (v3 << ((42) & 63))
                                                        var v52 uint32 = v3
                                                        v52 += v52
                                                        _ = v52
                                                    }
                                                }
                                                fmt.Println(fi37, v10, v10)
                                            } else {
                                                v53 := GenS6{F0: 997}
                                                v3 -= 65535
                                                v10 += 42
                                                _ = len(v2)
                                                _ = v53
                                            }
                                        case "\n\t":
                                            if (196 == 42) {
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                            }
                                        default:
                                            v10 = (v10 | v10)
                                        }
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        v54 := GenS7{F0: 467, F1: 1}
                                        if true {
                                            v10 -= v10
                                            var v55 uint8 = 1
                                            _ = v55
                                        }
                                        switch fi37 {
                                        case 0:
                                            var v56 uint64 = 967
                                            _ = v56
                                        default:
                                            v57 := GenS8{F0: (-714 >= -1)}
                                            _ = v57
                                        }
                                        switch "hello" {
                                        case "\n\t":
                                            v1 = append(v1, 240)
                                            if false {
                                                v3 = v3
                                                v58 := GenS2{F0: 42, F1: fi37, F2: 0}
                                                _ = v58
                                            } else {
                                                var v59 float32 = 0.0
                                                switch "a" {
                                                case "\n\t":
                                                    v60 := GenS8{F0: false}
                                                    _ = v60
                                                case "a":
                                                    if false {
                                                        var v61 bool = true
                                                        v62 := GenS3{F0: 204, F1: ("\n\t" + ("foo bar" + "hello"))}
                                                        _ = v61
                                                        _ = v62
                                                    }
                                                case "":
                                                    var v63 bool = true
                                                    _ = v63
                                                default:
                                                    if (("hello" + "a") >= "a") {
                                                        if (32767 < 0) {
                                                            var v64 int8 = v10
                                                            fmt.Println(v3, fi37)
                                                            var v65 int = (fi37 % ((9223372036854775807) | 1))
                                                            v66 := GenS5{F0: 148}
                                                            if false {
                                                                if len(v1) >= 2 {
                                                                    v1 = v1[1:len(v1)]
                                                                }
                                                            } else {
                                                                var v67 int = -685
                                                                if false {
                                                                    switch "foo bar" {
                                                                    case "a":
                                                                        switch (fi37 - fi34) {
                                                                        case -262:
                                                                            var v68 float32 = v59
                                                                            _ = v68
                                                                        case -204:
                                                                            v69 := GenS4{F0: 2}
                                                                            _ = v69
                                                                        case -796:
                                                                            var v70 int32 = 690
                                                                            _ = v70
                                                                        default:
                                                                            var v71 uint8 = 33
                                                                            switch (((("foo bar" + (("foo bar" + "foo bar") + "a")) + (("hello" + "") + "foo bar")) + (("" + "foo bar") + "hello")) + "") {
                                                                            case "":
                                                                                fmt.Println(fi34, v65, fi34)
                                                                            }
                                                                            v59 = (v59 * -2.25)
                                                                            var v72 uint64 = uint64(v3)
                                                                            _ = v71
                                                                            _ = v72
                                                                        }
                                                                        var v73 int32 = 0
                                                                        v59 -= (v59 * (0.1 * v59))
                                                                        _ = v73
                                                                    case "hello":
                                                                        if (229 <= 0) {
                                                                            var v74 uint = 18446744073709551615
                                                                            if false {
                                                                                v65 += (^(v67))
                                                                                if true {
                                                                                    var v75 uint16 = uint16(v3)
                                                                                    fmt.Println(fi37, v67, v65)
                                                                                    _ = v75
                                                                                } else {
                                                                                    v76 := GenS9{F0: 0, F1: 313, F2: 32}
                                                                                    v2 = append(v2, "\n\t")
                                                                                    var v77 int = -558
                                                                                    switch ("" + "hello") {
                                                                                    case "a":
                                                                                        v10 += 127
                                                                                    }
                                                                                    if (v74 >= v74) {
                                                                                        if (true && false) {
                                                                                            v1 = append(v1, v3)
                                                                                            v59 += (v59 * -2.25)
                                                                                            v78 := GenS9{F0: (v74 >> ((11) & 63)), F1: uint16(v3), F2: 42}
                                                                                            _ = v78
                                                                                        } else {
                                                                                            if (int(v74) <= 885) {
                                                                                                var v79 string = "\n\t"
                                                                                                _ = v79
                                                                                            } else {
                                                                                                switch "\n\t" {
                                                                                                case "\n\t":
                                                                                                    if ((v77 % ((2147483647) | 1)) >= fi37) {
                                                                                                        switch "hello" {
                                                                                                        case "a":
                                                                                                            _ = len(v1)
                                                                                                        case "foo bar":
                                                                                                            fmt.Println(fi37, v3, v64)
                                                                                                        }
                                                                                                        switch fi34 {
                                                                                                        case -2147483648:
                                                                                                            var v80 int = fi37
                                                                                                            _ = v80
                                                                                                        default:
                                                                                                            var v81 uint8 = genFn0(42, v10, (v10 + (v64 << (((v74 | v74)) & 63))))
                                                                                                            _ = v81
                                                                                                        }
                                                                                                    }
                                                                                                }
                                                                                                if true {
                                                                                                    var v82 int8 = (v64 << ((v74) & 63))
                                                                                                    v10 = v64
                                                                                                    v67 = v65
                                                                                                    _ = v82
                                                                                                } else {
                                                                                                    if true {
                                                                                                        v67 = v77
                                                                                                    } else {
                                                                                                        v83 := GenS6{F0: 42}
                                                                                                        if len(v2) > 0 {
                                                                                                            v2[0] = "\n\t"
                                                                                                        }
                                                                                                        _ = v83
                                                                                                    }
                                                                                                }
                                                                                            }
                                                                                            if len(v2) >= 2 {
                                                                                                v2 = v2[1:len(v2)]
                                                                                            }
                                                                                            v77 = v65
                                                                                        }
                                                                                        if len(v1) >= 2 {
                                                                                            v1 = v1[1:len(v1)]
                                                                                        }
                                                                                        var v84 uint32 = (v3 * v3)
                                                                                        if (255 == 127) {
                                                                                            v77 -= 218
                                                                                            switch "foo bar" {
                                                                                            case "":
                                                                                                v2 = append(v2, ("\n\t" + "hello"))
                                                                                            }
                                                                                        }
                                                                                        _ = v84
                                                                                    } else {
                                                                                        var v85 int16 = -32768
                                                                                        _ = v85
                                                                                    }
                                                                                    _ = v76
                                                                                    _ = v77
                                                                                }
                                                                                _ = len(v1)
                                                                            }
                                                                            switch ("hello" + "a") {
                                                                            case "foo bar":
                                                                                if (127 != 255) {
                                                                                    fmt.Println(v64, v3, fi34)
                                                                                    switch ("\n\t" + ("" + ("foo bar" + "foo bar"))) {
                                                                                    case "hello":
                                                                                        var v86 int16 = 1
                                                                                        _ = v86
                                                                                    case "":
                                                                                        v59 += 0.1
                                                                                    case "foo bar":
                                                                                        var v87 string = ("\n\t" + "")
                                                                                        _ = v87
                                                                                    }
                                                                                } else {
                                                                                    v59 = ((0.1 * v59) * float32(0.0))
                                                                                    v10 -= (v10 & -1)
                                                                                }
                                                                            case "hello":
                                                                                fmt.Println(v67, v67)
                                                                            case "\n\t":
                                                                                v74 += v74
                                                                            default:
                                                                                if true {
                                                                                    v88 := GenS10{F0: v3}
                                                                                    _ = v88
                                                                                } else {
                                                                                    fmt.Println(v3, v10)
                                                                                    switch (("a" + "hello") + "hello") {
                                                                                    case "a":
                                                                                        fmt.Println(v10, v3)
                                                                                    case "\n\t":
                                                                                        if (v3 >= (v3 - uint32(v74))) {
                                                                                            v59 += float32(float64(0.1))
                                                                                        } else {
                                                                                            if ((v65 << (((v74 + 251)) & 63)) > (v65 << ((v74) & 63))) {
                                                                                                fmt.Println(v64, v10)
                                                                                            }
                                                                                            var v89 int = -90
                                                                                            _ = v89
                                                                                        }
                                                                                    case "":
                                                                                        var v90 int = fi34
                                                                                        _ = v90
                                                                                    }
                                                                                    v91 := GenS7{F0: 1, F1: (v74 % ((v74) | 1))}
                                                                                    _ = v91
                                                                                }
                                                                            }
                                                                            _ = v74
                                                                        } else {
                                                                            switch (fi37 << ((4294967295) & 63)) {
                                                                            case 4294967295:
                                                                                v1 = append(v1, 42)
                                                                            }
                                                                            v10 += (v64 | 1)
                                                                            v92 := GenS11{F0: 86, F1: 0, F2: (995 < 18446744073709551615)}
                                                                            v93 := GenS9{F0: 255, F1: 2, F2: 2}
                                                                            fmt.Println(v65)
                                                                            _ = v92
                                                                            _ = v93
                                                                        }
                                                                    default:
                                                                        v94 := GenS1{F0: 303}
                                                                        _ = v94
                                                                    }
                                                                    switch (v67 * (v67 | fi37)) {
                                                                    case -70:
                                                                        v64 = v64
                                                                    }
                                                                }
                                                                _ = v67
                                                            }
                                                            _ = v64
                                                            _ = v65
                                                            _ = v66
                                                        }
                                                        if (0 != 9223372036854775807) {
                                                            var v95 int32 = -2147483648
                                                            var v96 uint8 = uint8(v3)
                                                            _ = v95
                                                            _ = v96
                                                        } else {
                                                            v59 = (v59 * -2.25)
                                                        }
                                                        var v97 bool = true
                                                        _ = v97
                                                    } else {
                                                        fmt.Println(fi34, fi34, v10)
                                                    }
                                                }
                                                _ = v59
                                            }
                                            fmt.Println(fi37, fi37, fi37)
                                            fmt.Println(fi34, fi34)
                                        case "":
                                            if (!true) {
                                                v1 = append(v1, v3)
                                            } else {
                                                var v98 int = 4294967295
                                                if (true || false) {
                                                    if false {
                                                        var v99 uint8 = 0
                                                        v100 := GenS11{F0: genFn0(-32768, (v10 / (((-(v10))) | 1)), v10), F1: 2, F2: false}
                                                        _ = v100
                                                        _ = v99
                                                    } else {
                                                        v3 += (v3 + v3)
                                                        if len(v1) >= 2 {
                                                            v1 = v1[1:len(v1)]
                                                        }
                                                    }
                                                    fmt.Println(fi34, fi34, v98)
                                                    if len(v1) > 0 {
                                                        v1[0] = (v3 % (((v3 - (v3 << ((4294967295) & 63)))) | 1))
                                                    }
                                                }
                                                var v101 int64 = 0
                                                _ = v101
                                                _ = v98
                                            }
                                        default:
                                            v10 = v10
                                        }
                                        _ = v54
                                    }
                                } else {
                                    if false {
                                        if true {
                                            switch (fi37 << ((42) & 63)) {
                                            case 32767:
                                                v3 = (v3 << ((1) & 63))
                                            case 2147483647:
                                                var v102 uint = 375
                                                _ = v102
                                            case 116:
                                                v103 := GenS12{F0: -1}
                                                _ = v103
                                            default:
                                                var v104 int32 = 362
                                                _ = v104
                                            }
                                            var v105 uint8 = 248
                                            v105 += v105
                                            var v106 uint = uint(v105)
                                            _ = v105
                                            _ = v106
                                        } else {
                                            switch 9223372036854775807 {
                                            case 4294967295:
                                                switch fi37 {
                                                case -624:
                                                    fmt.Println(fi34, fi37, fi37)
                                                case 0:
                                                    var v107 string = ("a" + ("hello" + ("a" + ("\n\t" + "foo bar"))))
                                                    _ = v107
                                                case 347:
                                                    switch "hello" {
                                                    case "hello":
                                                        var v108 int8 = (v10 % (((v10 / ((v10) | 1))) | 1))
                                                        _ = v108
                                                    case "\n\t":
                                                        v2 = append(v2, "hello")
                                                    case "":
                                                        _ = len(v1)
                                                        v109 := GenS13{F0: 4294967295}
                                                        _ = v109
                                                    default:
                                                        v110 := GenS11{F0: 30, F1: 451, F2: true}
                                                        _ = v110
                                                    }
                                                }
                                            case 2147483647:
                                                var v111 int16 = 32767
                                                _ = v111
                                            }
                                        }
                                        var v112 uint8 = 42
                                        _ = v112
                                    } else {
                                        var v113 uint32 = 2
                                        fmt.Println(fi34, v10)
                                        _ = v113
                                    }
                                    if (65535 != 0) {
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        v10 = (v10 << ((38) & 63))
                                    } else {
                                        if len(v1) > 0 {
                                            v1[0] = (v3 << ((18446744073709551615) & 63))
                                        }
                                        switch "\n\t" {
                                        case "foo bar":
                                            if false {
                                                v114 := GenS11{F0: 5, F1: 2, F2: (!false)}
                                                fmt.Println(fi37, v10, fi37)
                                                if (false && (false || (false && (genFn0(1, v10, 1) < 128)))) {
                                                    if false {
                                                        v3 += (v3 / ((v3) | 1))
                                                        v3 = v3
                                                    } else {
                                                        if len(v2) > 0 {
                                                            v2[0] = ("hello" + ("" + "a"))
                                                        }
                                                        v115 := GenS14{F0: v3, F1: fi34}
                                                        var v116 uint16 = uint16(v3)
                                                        switch ("a" + "") {
                                                        case "a":
                                                            var v117 uint32 = 341
                                                            _ = v117
                                                        default:
                                                            switch fi34 {
                                                            case -9223372036854775808:
                                                                if (0 >= 32767) {
                                                                    if false {
                                                                        _ = len(v1)
                                                                        v3 -= v3
                                                                        fmt.Println(v3)
                                                                    } else {
                                                                        var v118 uint64 = 2
                                                                        if len(v1) >= 2 {
                                                                            v1 = v1[1:len(v1)]
                                                                        }
                                                                        if (-962 == -1) {
                                                                            v3 = v3
                                                                            v119 := GenS3{F0: genFn0(1000, 42, 42), F1: ("foo bar" + "\n\t")}
                                                                            var v120 int = 617
                                                                            fmt.Println(v120)
                                                                            var v121 uint8 = 71
                                                                            _ = v119
                                                                            _ = v120
                                                                            _ = v121
                                                                        } else {
                                                                            fmt.Println(v118, v116, v10)
                                                                            if false {
                                                                                var v122 string = "\n\t"
                                                                                var v123 int32 = 1
                                                                                _ = v122
                                                                                _ = v123
                                                                            } else {
                                                                                if false {
                                                                                    v2 = append(v2, ("\n\t" + ("foo bar" + ((("" + ("foo bar" + "foo bar")) + (("" + (("" + ("hello" + "hello")) + "")) + (("a" + (("a" + ("\n\t" + ("foo bar" + "\n\t"))) + "hello")) + "hello"))) + (("a" + ("" + "a")) + "foo bar")))))
                                                                                    v2 = append(v2, ((("foo bar" + "\n\t") + "foo bar") + ""))
                                                                                }
                                                                                switch "" {
                                                                                case "a":
                                                                                    v2 = append(v2, ("hello" + ("a" + "\n\t")))
                                                                                }
                                                                                if (0.0 >= ((-2.25 - 0.0) + 0.1)) {
                                                                                    if (genFn0(1, int8(v118), (v10 ^ v10)) != 42) {
                                                                                        v1 = append(v1, (v3 + v3))
                                                                                    } else {
                                                                                        v118 = v118
                                                                                    }
                                                                                    if true {
                                                                                        v118 = v118
                                                                                    } else {
                                                                                        v2 = append(v2, ("a" + ("hello" + (("hello" + "hello") + "foo bar"))))
                                                                                        switch (fi37 - fi37) {
                                                                                        case -9223372036854775808:
                                                                                            var v124 int16 = 32767
                                                                                            _ = v124
                                                                                        }
                                                                                        if len(v1) >= 2 {
                                                                                            v1 = v1[1:len(v1)]
                                                                                        }
                                                                                    }
                                                                                    if false {
                                                                                        v3 = (v3 | v3)
                                                                                    }
                                                                                }
                                                                                switch (fi34 & fi37) {
                                                                                case 1:
                                                                                    _ = len(v1)
                                                                                    v125 := GenS8{F0: true}
                                                                                    _ = v125
                                                                                case 8:
                                                                                    var v126 uint32 = 42
                                                                                    _ = v126
                                                                                case -9223372036854775808:
                                                                                    switch ("\n\t" + "foo bar") {
                                                                                    case "foo bar":
                                                                                        v127 := GenS6{F0: int16(v118)}
                                                                                        _ = v127
                                                                                    }
                                                                                default:
                                                                                    fmt.Println(v10, v10)
                                                                                }
                                                                                v128 := GenS15{F0: 32767}
                                                                                if ("" == "a") {
                                                                                    v118 += v118
                                                                                }
                                                                                _ = v128
                                                                            }
                                                                            fmt.Println(v118, v3)
                                                                            v3 = (v3 % (((^(v3))) | 1))
                                                                        }
                                                                        _ = v118
                                                                    }
                                                                    fmt.Println(v116, fi34)
                                                                    v129 := GenS9{F0: 2, F1: (v116 >> ((1) & 63)), F2: 0}
                                                                    _ = v129
                                                                } else {
                                                                    if (true && false) {
                                                                        switch "a" {
                                                                        case "foo bar":
                                                                            v116 = 163
                                                                        case "":
                                                                            var v130 uint64 = 198
                                                                            _ = v130
                                                                        case "\n\t":
                                                                            var v131 bool = false
                                                                            _ = v131
                                                                        }
                                                                        var v132 float32 = -2.25
                                                                        _ = v132
                                                                    } else {
                                                                        v116 -= (v116 << ((838) & 63))
                                                                    }
                                                                    v10 -= (v10 % ((v10) | 1))
                                                                    v133 := GenS7{F0: 971, F1: 42}
                                                                    var v134 uint32 = (v3 ^ 2147483648)
                                                                    if true {
                                                                        var v135 int8 = (^(v10))
                                                                        _ = v135
                                                                    }
                                                                    _ = v133
                                                                    _ = v134
                                                                }
                                                            case 9223372036854775807:
                                                                v136 := GenS10{F0: v3}
                                                                _ = v136
                                                            default:
                                                                v116 -= (v116 * (v116 / ((v116) | 1)))
                                                            }
                                                        }
                                                        _ = v115
                                                        _ = v116
                                                    }
                                                    switch "foo bar" {
                                                    case "\n\t":
                                                        v10 -= v10
                                                    case "":
                                                        var v137 uint8 = genFn0(32767, -1, (v10 << ((144) & 63)))
                                                        _ = v137
                                                    }
                                                }
                                                if (int32(v3) < 186) {
                                                    v3 += v3
                                                    var v138 int16 = -1
                                                    v10 += v10
                                                    _ = v138
                                                }
                                                switch -9223372036854775808 {
                                                case 4294967295:
                                                    var v139 float64 = 1e10
                                                    _ = v139
                                                case -171:
                                                    if (false || (!false)) {
                                                        fmt.Println(fi37)
                                                        fmt.Println(v3, v3)
                                                        var v140 int8 = 0
                                                        var v141 int32 = 1
                                                        _ = v140
                                                        _ = v141
                                                    }
                                                case 408:
                                                    v2 = append(v2, "hello")
                                                    if ((^(v3)) <= (v3 / ((569) | 1))) {
                                                        v142 := GenS12{F0: -468}
                                                        switch (fi37 & -9223372036854775808) {
                                                        case 127:
                                                            if false {
                                                                var v143 uint16 = 2
                                                                _ = v143
                                                            } else {
                                                                if (1.5 >= 3.14159) {
                                                                    fmt.Println(v10, fi37)
                                                                    var v144 uint16 = 2
                                                                    _ = v144
                                                                }
                                                            }
                                                        case 9223372036854775807:
                                                            fmt.Println(fi34)
                                                            if (1 >= 1) {
                                                                var v145 int16 = 1
                                                                fmt.Println(v10, v145, fi34)
                                                                v146 := GenS16{F0: 62, F1: -2147483648}
                                                                _ = v145
                                                                _ = v146
                                                            }
                                                            var v147 uint32 = (v3 << ((1) & 63))
                                                            v147 = (^(v147))
                                                            _ = v147
                                                        case 128:
                                                            v148 := GenS17{F0: 100, F1: 18446744073709551615}
                                                            var v149 bool = false
                                                            switch 42 {
                                                            case 42:
                                                                v149 = (("" + "\n\t") <= "\n\t")
                                                            case 0:
                                                                switch "foo bar" {
                                                                case "\n\t":
                                                                    v10 -= (v10 ^ v10)
                                                                    var v150 int = fi37
                                                                    _ = v150
                                                                case "foo bar":
                                                                    v3 = 267
                                                                }
                                                            case -128:
                                                                var v151 uint16 = 255
                                                                _ = v151
                                                            default:
                                                                if true {
                                                                    if (42 >= 0) {
                                                                        if len(v1) > 0 {
                                                                            v1[0] = 1
                                                                        }
                                                                        if len(v2) > 0 {
                                                                            v2[0] = ("\n\t" + "")
                                                                        }
                                                                        if (!true) {
                                                                            switch "foo bar" {
                                                                            case "hello":
                                                                                fmt.Println(v149, fi37, v10)
                                                                            default:
                                                                                v10 = (v10 & (v10 >> ((65535) & 63)))
                                                                            }
                                                                            fmt.Println(fi34)
                                                                            fmt.Println(fi34, fi34, v3)
                                                                            v149 = v149
                                                                        } else {
                                                                            v149 = (4294967295 <= 0)
                                                                            v3 += (v3 & v3)
                                                                            var v152 float64 = 3.14159
                                                                            _ = v152
                                                                        }
                                                                        v153 := GenS12{F0: -9223372036854775808}
                                                                        v10 = (v10 % (((v10 % ((v10) | 1))) | 1))
                                                                        _ = v153
                                                                    }
                                                                }
                                                            }
                                                            _ = v148
                                                            _ = v149
                                                        default:
                                                            v1 = append(v1, (v3 & (v3 << ((1) & 63))))
                                                        }
                                                        _ = v142
                                                    } else {
                                                        var v154 uint32 = v3
                                                        _ = v154
                                                    }
                                                    v10 = v10
                                                default:
                                                    if false {
                                                        if (!true) {
                                                            fmt.Println(fi37, v10)
                                                            v155 := GenS14{F0: 4294967295, F1: (fi37 * (fi34 ^ fi34))}
                                                            _ = v155
                                                        }
                                                    }
                                                }
                                                _ = v114
                                            } else {
                                                v3 = 2
                                                var v156 int8 = 66
                                                var v157 uint64 = 824
                                                _ = v156
                                                _ = v157
                                            }
                                        }
                                        if false {
                                            if (!true) {
                                                if len(v1) > 0 {
                                                    v1[0] = (v3 << ((65535) & 63))
                                                }
                                                switch "" {
                                                case "":
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                default:
                                                    v3 = 2147483648
                                                }
                                            } else {
                                                var v158 string = "foo bar"
                                                v3 -= (v3 * v3)
                                                _ = v158
                                            }
                                        } else {
                                            v159 := GenS9{F0: 65535, F1: 1, F2: 18446744073709551615}
                                            v3 += (v3 << ((255) & 63))
                                            _ = v159
                                        }
                                    }
                                    var v160 int8 = 100
                                    _ = v160
                                }
                                fmt.Println(fi37, fi37, v3)
                                if false {
                                    switch (fi37 >> ((42) & 63)) {
                                    case 2147483647:
                                        v10 = (v10 % ((127) | 1))
                                    case 627:
                                        if (1 < 42) {
                                            if true {
                                                v161 := GenS18{F0: 0}
                                                switch ("foo bar" + "\n\t") {
                                                case "foo bar":
                                                    v3 = v3
                                                case "\n\t":
                                                    v162 := GenS19{F0: -1, F1: 2}
                                                    _ = v162
                                                case "hello":
                                                    var v163 uint32 = 0
                                                    _ = v163
                                                }
                                                v3 -= v3
                                                if ((!(0 == 65535)) || true) {
                                                    if false {
                                                        var v164 int = fi34
                                                        _ = v164
                                                    }
                                                } else {
                                                    if ("foo bar" > "\n\t") {
                                                        switch (fi34 & 42) {
                                                        case -9223372036854775808:
                                                            if true {
                                                                var v165 int = 738
                                                                _ = v165
                                                            } else {
                                                                var v166 string = "\n\t"
                                                                _ = v166
                                                            }
                                                        default:
                                                            v3 = (v3 << ((294) & 63))
                                                        }
                                                    }
                                                }
                                                var v167 float64 = 1e10
                                                v167 += 1e10
                                                _ = v161
                                                _ = v167
                                            }
                                        } else {
                                            if false {
                                                var v168 int32 = 103
                                                _ = v168
                                            } else {
                                                v169 := GenS12{F0: -1}
                                                var v170 float32 = 3.14159
                                                _ = len(v1)
                                                _ = v169
                                                _ = v170
                                            }
                                            if (!(922 >= 847)) {
                                                v2 = append(v2, "a")
                                                fmt.Println(fi37, fi34, fi34)
                                            } else {
                                                v171 := GenS19{F0: -1, F1: 624}
                                                if (!(-2147483648 >= 0)) {
                                                    v3 -= (^(v3))
                                                } else {
                                                    var v172 int64 = -296
                                                    v172 += -1
                                                    _ = v172
                                                }
                                                fmt.Println(v10, fi37)
                                                _ = v171
                                            }
                                            v10 += (-(v10))
                                        }
                                    }
                                    if true {
                                        var v173 int = fi34
                                        if (1 < 32767) {
                                            switch (((("\n\t" + "foo bar") + ("foo bar" + "a")) + ("" + "")) + "foo bar") {
                                            case "":
                                                var v174 uint32 = (v3 * v3)
                                                _ = v174
                                            case "a":
                                                switch -287 {
                                                case 127:
                                                    _ = len(v1)
                                                case -503:
                                                    v173 -= fi37
                                                case -9223372036854775808:
                                                    v175 := GenS6{F0: 32767}
                                                    _ = v175
                                                default:
                                                    fmt.Println(v10)
                                                }
                                            }
                                            if len(v1) > 0 {
                                                v1[0] = (v3 + (v3 | v3))
                                            }
                                        } else {
                                            v173 = 4294967295
                                            if len(v2) > 0 {
                                                v2[0] = ("a" + ("hello" + ("a" + "a")))
                                            }
                                        }
                                        var v176 int32 = 379
                                        _ = v173
                                        _ = v176
                                    } else {
                                        if (!false) {
                                            v3 -= 1
                                            if (((-459 <= 1000) || false) && (65535 == (v3 + v3))) {
                                                var v177 uint8 = 251
                                                _ = v177
                                            } else {
                                                var v178 float32 = 0.1
                                                _ = v178
                                            }
                                            var v179 uint = uint(v3)
                                            v3 -= (^(v3))
                                            _ = v179
                                        } else {
                                            switch "a" {
                                            case "a":
                                                var v180 float32 = (0.0 - 1.5)
                                                _ = v180
                                            case "foo bar":
                                                if ((2 > 0) && true) {
                                                    v3 = 577
                                                    if ((fi37 | fi34) != 4294967295) {
                                                        var v181 uint = 0
                                                        _ = v181
                                                    } else {
                                                        v10 += 100
                                                        v3 -= (^(v3))
                                                        if false {
                                                            v3 += (v3 + v3)
                                                            v3 -= 547
                                                        } else {
                                                            if true {
                                                                v182 := GenS12{F0: 1}
                                                                _ = v182
                                                            }
                                                            v183 := GenS14{F0: (v3 ^ (v3 | v3)), F1: (fi37 << ((961) & 63))}
                                                            switch (fi37 << ((4294967295) & 63)) {
                                                            case 127:
                                                                v184 := GenS2{F0: 821, F1: (fi34 >> ((uint(v3)) & 63)), F2: -1}
                                                                _ = v184
                                                            case 0:
                                                                if false {
                                                                    if true {
                                                                        fmt.Println(fi37, fi37, v3)
                                                                    } else {
                                                                        v10 -= (v10 ^ v10)
                                                                    }
                                                                    var v185 string = "hello"
                                                                    _ = v185
                                                                } else {
                                                                    fmt.Println(v10, fi37, fi37)
                                                                    v2 = append(v2, ("foo bar" + "foo bar"))
                                                                }
                                                            case -2147483648:
                                                                fmt.Println(v3, v3)
                                                            default:
                                                                var v186 int32 = -648
                                                                _ = v186
                                                            }
                                                            if false {
                                                                switch 9223372036854775807 {
                                                                case -407:
                                                                    var v187 string = ""
                                                                    _ = v187
                                                                case 82:
                                                                    var v188 int64 = 42
                                                                    _ = v188
                                                                case 127:
                                                                    v10 = v10
                                                                }
                                                                switch (^(fi34)) {
                                                                case 49:
                                                                    var v189 float64 = 0.0
                                                                    _ = v189
                                                                case 0:
                                                                    v2 = append(v2, "hello")
                                                                    v190 := GenS18{F0: uint16(v3)}
                                                                    _ = v190
                                                                case 32768:
                                                                    var v191 string = ("" + "a")
                                                                    _ = v191
                                                                }
                                                            } else {
                                                                switch "a" {
                                                                case "foo bar":
                                                                    var v192 uint32 = (v3 - v3)
                                                                    _ = v192
                                                                case "\n\t":
                                                                    v193 := GenS16{F0: 0, F1: 766}
                                                                    if true {
                                                                        v194 := GenS4{F0: 243}
                                                                        switch fi34 {
                                                                        case 0:
                                                                            var v195 uint = 4294967295
                                                                            _ = v195
                                                                        case 215:
                                                                            if (255 >= 2) {
                                                                                switch -9223372036854775808 {
                                                                                case 776:
                                                                                    v196 := GenS20{F0: "", F1: 605, F2: 2}
                                                                                    _ = v196
                                                                                case 4294967295:
                                                                                    fmt.Println(v3, v10, fi37)
                                                                                default:
                                                                                    if true {
                                                                                        var v197 int32 = 379
                                                                                        switch (fi37 << ((708) & 63)) {
                                                                                        case -1:
                                                                                            if ((1e10 - float64(1e10)) >= -2.25) {
                                                                                                var v198 uint32 = (v3 << ((uint(v3)) & 63))
                                                                                                var v199 float64 = (1.5 / 3.14159)
                                                                                                if len(v2) > 0 {
                                                                                                    v2[0] = ("hello" + "\n\t")
                                                                                                }
                                                                                                _ = v198
                                                                                                _ = v199
                                                                                            } else {
                                                                                                v197 = 42
                                                                                                _ = len(v2)
                                                                                                var v200 uint16 = 32767
                                                                                                var v201 uint = 65535
                                                                                                v202 := GenS21{F0: v3, F1: 42, F2: 65535}
                                                                                                if false {
                                                                                                    var v203 float32 = 1e10
                                                                                                    if false {
                                                                                                        var v204 string = ""
                                                                                                        _ = v204
                                                                                                    }
                                                                                                    _ = v203
                                                                                                } else {
                                                                                                    fmt.Println(fi37, fi34)
                                                                                                }
                                                                                                _ = v200
                                                                                                _ = v201
                                                                                                _ = v202
                                                                                            }
                                                                                        case 4294967295:
                                                                                            if len(v1) > 0 {
                                                                                                v1[0] = (v3 >> ((0) & 63))
                                                                                            }
                                                                                        case -2147483648:
                                                                                            var v205 uint8 = 9
                                                                                            _ = v205
                                                                                        default:
                                                                                            v206 := GenS10{F0: v3}
                                                                                            _ = v206
                                                                                        }
                                                                                        _ = v197
                                                                                    }
                                                                                }
                                                                            }
                                                                        case -126:
                                                                            var v207 int8 = v10
                                                                            _ = v207
                                                                        }
                                                                        v10 = v10
                                                                        _ = v194
                                                                    } else {
                                                                        v10 = 1
                                                                        if true {
                                                                            fmt.Println(v10, v3, v3)
                                                                            fmt.Println(fi34, fi34, fi34)
                                                                            switch fi37 {
                                                                            case 32767:
                                                                                v208 := GenS15{F0: -1}
                                                                                _ = v208
                                                                            case -2147483648:
                                                                                if ((3.14159 - ((float64(0.0) / 0.1) / 3.14159)) > 1e10) {
                                                                                    v209 := GenS22{F0: 0}
                                                                                    if false {
                                                                                        var v210 float32 = (0.0 + 0.0)
                                                                                        if true {
                                                                                            var v211 uint64 = 482
                                                                                            var v212 float32 = (v210 / v210)
                                                                                            if (!((!true) || true)) {
                                                                                                v2 = append(v2, "a")
                                                                                                var v213 bool = (749 >= 1)
                                                                                                _ = v213
                                                                                            } else {
                                                                                                _ = len(v1)
                                                                                                v211 = (v211 ^ v211)
                                                                                                v214 := GenS11{F0: 1, F1: 0, F2: true}
                                                                                                _ = v214
                                                                                            }
                                                                                            _ = v211
                                                                                            _ = v212
                                                                                        } else {
                                                                                            fmt.Println(v10, v3, v3)
                                                                                            v10 += v10
                                                                                            v10 = 30
                                                                                        }
                                                                                        fmt.Println(fi34)
                                                                                        _ = v210
                                                                                    } else {
                                                                                        v3 += v3
                                                                                        fmt.Println(fi34, fi37)
                                                                                        v215 := GenS7{F0: -1, F1: 42}
                                                                                        if true {
                                                                                            v3 = (v3 - (v3 % (((v3 % ((v3) | 1))) | 1)))
                                                                                        } else {
                                                                                            if len(v1) > 0 {
                                                                                                v1[0] = (v3 ^ v3)
                                                                                            }
                                                                                        }
                                                                                        _ = v215
                                                                                    }
                                                                                    var v216 int64 = -677
                                                                                    _ = v209
                                                                                    _ = v216
                                                                                } else {
                                                                                    var v217 uint8 = 1
                                                                                    _ = v217
                                                                                }
                                                                            }
                                                                            if false {
                                                                                v3 = v3
                                                                                v3 += v3
                                                                            } else {
                                                                                if ((-964 >= 42) || false) {
                                                                                    var v218 int32 = -1
                                                                                    _ = v218
                                                                                }
                                                                            }
                                                                        } else {
                                                                            v219 := GenS23{F0: 42}
                                                                            var v220 string = ("foo bar" + (("" + "\n\t") + "foo bar"))
                                                                            _ = v219
                                                                            _ = v220
                                                                        }
                                                                        switch (("a" + ("a" + "foo bar")) + (("foo bar" + "") + "\n\t")) {
                                                                        case "\n\t":
                                                                            var v221 int8 = v10
                                                                            _ = v221
                                                                        case "foo bar":
                                                                            v2 = append(v2, ("\n\t" + "hello"))
                                                                        }
                                                                        var v222 int32 = 836
                                                                        v3 -= v3
                                                                        _ = v222
                                                                    }
                                                                    v3 += v3
                                                                    _ = v193
                                                                }
                                                                v10 = 42
                                                                fmt.Println(fi37)
                                                            }
                                                            var v223 uint8 = 236
                                                            _ = v183
                                                            _ = v223
                                                        }
                                                    }
                                                    v10 = v10
                                                } else {
                                                    switch (-(fi37)) {
                                                    case 4294967295:
                                                        var v224 string = "\n\t"
                                                        _ = v224
                                                    }
                                                    if true {
                                                        if false {
                                                            switch "hello" {
                                                            case "\n\t":
                                                                var v225 int16 = 28
                                                                var v226 string = "hello"
                                                                switch (^(fi34)) {
                                                                case 0:
                                                                    v225 += (v225 % (((v225 * v225)) | 1))
                                                                case 127:
                                                                    var v227 uint32 = v3
                                                                    _ = v227
                                                                }
                                                                _ = v225
                                                                _ = v226
                                                            case "hello":
                                                                _ = len(v2)
                                                            case "foo bar":
                                                                v10 = v10
                                                            }
                                                            var v228 uint64 = 18446744073709551615
                                                            if len(v1) > 0 {
                                                                v1[0] = v3
                                                            }
                                                            _ = v228
                                                        }
                                                        fmt.Println(fi37, v3)
                                                    } else {
                                                        fmt.Println(v10, fi34)
                                                        v2 = append(v2, "a")
                                                    }
                                                    var v229 bool = false
                                                    _ = v229
                                                }
                                            case "hello":
                                                v230 := GenS22{F0: 1}
                                                if false {
                                                    v1 = append(v1, (^(v3)))
                                                    v10 += (-(v10))
                                                }
                                                var v231 int64 = -9223372036854775808
                                                _ = v230
                                                _ = v231
                                            }
                                            if true {
                                                var v232 uint = 65535
                                                _ = v232
                                            }
                                            v3 = (^(v3))
                                        }
                                        v1 = append(v1, 42)
                                        var v233 int64 = -1
                                        _ = v233
                                    }
                                    switch "foo bar" {
                                    case "":
                                        switch "" {
                                        case "\n\t":
                                            if len(v1) > 0 {
                                                v1[0] = v3
                                            }
                                        case "hello":
                                            v3 = (v3 * v3)
                                        case "a":
                                            fmt.Println(v10)
                                        }
                                    default:
                                        fmt.Println(v10, fi37, fi34)
                                    }
                                } else {
                                    if (false && (42 > 0)) {
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        var v234 int = (fi37 | fi34)
                                        fmt.Println(fi37, v234)
                                        _ = v234
                                    } else {
                                        switch ("a" + "a") {
                                        case "":
                                            if (!(v3 > 829)) {
                                                fmt.Println(v10, fi37)
                                            } else {
                                                switch (("foo bar" + "hello") + ("foo bar" + (((("foo bar" + "\n\t") + "\n\t") + "\n\t") + ("hello" + ("\n\t" + ""))))) {
                                                case "a":
                                                    v3 = (v3 | (v3 * v3))
                                                case "":
                                                    var v235 float32 = 3.14159
                                                    _ = v235
                                                default:
                                                    switch fi37 {
                                                    case -964:
                                                        switch "hello" {
                                                        case "foo bar":
                                                            var v236 int64 = -1
                                                            _ = v236
                                                        default:
                                                            var v237 int16 = 804
                                                            _ = v237
                                                        }
                                                    case 52:
                                                        v1 = append(v1, v3)
                                                    }
                                                }
                                                v238 := GenS3{F0: genFn0(0, (-(v10)), (v10 % (((v10 / ((v10) | 1))) | 1))), F1: "\n\t"}
                                                switch "\n\t" {
                                                case "hello":
                                                    var v239 int32 = -772
                                                    _ = v239
                                                }
                                                _ = v238
                                            }
                                        case "foo bar":
                                            fmt.Println(v3)
                                        }
                                        if (!true) {
                                            if false {
                                                v10 = (v10 / (((v10 << ((312) & 63))) | 1))
                                                if len(v2) > 0 {
                                                    v2[0] = "\n\t"
                                                }
                                            }
                                            var v240 uint32 = 707
                                            v241 := GenS24{F0: 32767, F1: 501, F2: 606}
                                            _ = v240
                                            _ = v241
                                        }
                                    }
                                }
                            }
                            if false {
                                fmt.Println(v10)
                                var v242 uint8 = 127
                                switch (fi37 >> ((1) & 63)) {
                                case 411:
                                    var v243 uint16 = uint16(v242)
                                    _ = v243
                                case -1:
                                    v244 := GenS17{F0: v10, F1: 1}
                                    _ = v244
                                case -9223372036854775808:
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                }
                                _ = v242
                            }
                            if (!false) {
                                v245 := GenS25{F0: 255}
                                _ = v245
                            } else {
                                switch fi37 {
                                case 32767:
                                    if (false && (v3 > v3)) {
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        var v246 int32 = 0
                                        var v247 int = (fi34 - (fi34 ^ fi37))
                                        _ = v246
                                        _ = v247
                                    } else {
                                        var v248 uint64 = 268
                                        var v249 float32 = (1e10 + float32(3.14159))
                                        if true {
                                            fmt.Println(v249, fi37)
                                            v248 -= (^(v248))
                                            if len(v2) > 0 {
                                                v2[0] = "\n\t"
                                            }
                                        } else {
                                            v250 := GenS21{F0: (v3 | v3), F1: -9223372036854775808, F2: 1}
                                            _ = v250
                                        }
                                        var v251 int8 = (v10 << ((65535) & 63))
                                        _ = v248
                                        _ = v249
                                        _ = v251
                                    }
                                }
                                fmt.Println(fi34)
                                v252 := GenS21{F0: 1, F1: -1, F2: 65535}
                                _ = v252
                            }
                            _ = v40
                        }
                        _ = fi37
                    }
                }
            }
            _ = v35
        } else {
            var v253 string = ("a" + ("a" + "a"))
            for fi254 := 0; fi254 < 18; fi254++ {
                v255 := GenS8{F0: (v3 <= v3)}
                if false {
                    v10 -= (v10 >> ((830) & 63))
                    if ((v3 + (v3 - v3)) >= (v3 << ((192) & 63))) {
                        if (!(v3 != v3)) {
                            v253 = v253
                            var v256 uint8 = 246
                            v10 -= v10
                            v257 := GenS26{F0: 0, F1: v10, F2: int64(v10)}
                            if ((127 <= v10) || false) {
                                var v258 float32 = 3.14159
                                _ = len(v2)
                                _ = v258
                            }
                            _ = v256
                            _ = v257
                        }
                        fmt.Println(fi34, v253)
                        v3 -= (v3 & (^(v3)))
                    }
                    var v259 uint8 = 127
                    _ = v259
                } else {
                    v260 := GenS27{F0: v3, F1: v3}
                    _ = v260
                }
                _ = fi254
                _ = v255
            }
            v3 = v3
            _ = v253
        }
        _ = fi34
    }
    v261 := GenS28{F0: genFn0(42, 0, -128)}
    v262 := GenS28{F0: 42}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v263 int32 = 2147483647
    var v264 float64 = 1.5
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v263:", v263)
    fmt.Println("v264:", v264)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v2
    _ = v261
    _ = v262
    _ = v263
    _ = v264
    _ = v3
    _ = v33
}
