package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 int8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint; F1 int64; F2 uint32 }
type GenS5 struct { F0 uint16; F1 int8 }
type GenS6 struct { F0 int16; F1 int32; F2 int64 }
type GenS7 struct { F0 uint64; F1 int }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 uint8 }
type GenS10 struct { F0 uint }
type GenS11 struct { F0 uint16; F1 uint32; F2 uint32 }
type GenS12 struct { F0 int64 }
type GenS13 struct { F0 uint16; F1 int64; F2 int32 }
type GenS14 struct { F0 int64 }
type GenS15 struct { F0 uint8; F1 int64; F2 int64 }
type GenS16 struct { F0 uint32; F1 int; F2 uint32 }
func genFn0() int {
    v7 := GenS1{F0: 858, F1: 68, F2: -128}
    for fi8 := 0; fi8 < 9; fi8++ {
        fmt.Println(fi8)
        fmt.Println(fi8)
        _ = fi8
    }
    var v9 uint32 = 1
    _ = v7
    _ = v9
    return 0
}
func genFn1(p0 int16, p1 uint, p2 uint) int64 {
    v10 := []int{42}
    p0 -= (p0 * p0)
    for fi11 := 0; fi11 < 20; fi11++ {
        v12 := GenS2{F0: 42}
        for fi13 := 0; fi13 < 3; fi13++ {
            v14 := GenS1{F0: 255, F1: 127, F2: 127}
            if true {
                v15 := GenS3{F0: 2}
                switch "a" {
                case "a":
                    p2 = (p2 >> ((p1) & 63))
                case "foo bar":
                    var v16 string = (("hello" + ("" + "\n\t")) + "hello")
                    _ = v16
                case "":
                    var v17 int32 = 42
                    _ = v17
                }
                v18 := GenS3{F0: 2}
                _ = v15
                _ = v18
            }
            _ = fi13
            _ = v14
        }
        var v19 int8 = 0
        _ = fi11
        _ = v12
        _ = v19
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return 0
}

func main() {
    v1 := []uint32{0, 2147483648}
    v2 := make(map[int]string)
    var v3 bool = ((1.5 / (1e10 / (float64((float32(((0.0 + 1e10) + 0.0)) * -2.25)) / 1e10))) > (-2.25 + 1.5))
    var v4 uint16 = 352
    v5 := make(map[int]bool)
    var v6 int8 = 0
    if true {
        var v20 uint16 = v4
        _ = v20
    } else {
        var v21 int16 = -1
        v6 = (v6 & 0)
        _ = v21
    }
    fmt.Println(v4, v3)
    if v3 {
        for fi22 := 0; fi22 < 8; fi22++ {
            for fi23 := 0; fi23 < 3; fi23++ {
                fmt.Println(fi22, v4, fi22)
                v6 += (v6 % (((v6 % ((-128) | 1))) | 1))
                v24 := GenS2{F0: -128}
                _ = fi23
                _ = v24
            }
            for fi25 := 0; fi25 < 1; fi25++ {
                v26 := GenS2{F0: 127}
                fmt.Println(fi25, fi22, v6)
                _ = fi25
                _ = v26
            }
            v3 = false
            _ = fi22
        }
        v27 := GenS3{F0: 269}
        _ = v27
    }
    if v3 {
        if (genFn1(-32768, 1, 25) >= 984) {
            var v28 float32 = (1e10 - 3.14159)
            var v29 uint16 = v4
            _ = v28
            _ = v29
        } else {
            v6 = v6
        }
        if v3 {
            switch 369 {
            case 0:
                v30 := GenS2{F0: v6}
                _ = v30
            case -2147483648:
                for fi31 := 0; fi31 < 1; fi31++ {
                    v32 := GenS2{F0: (^(v6))}
                    switch ("" + ((("hello" + "foo bar") + "hello") + "a")) {
                    case "a":
                        switch ((("a" + ("" + "")) + ("hello" + (("hello" + "hello") + ("" + (((("\n\t" + "") + "\n\t") + "foo bar") + "hello"))))) + "") {
                        case "":
                            fmt.Println(v3)
                        case "a":
                            var v33 int32 = 1
                            _ = v33
                        }
                    case "\n\t":
                        v6 += (v6 + v6)
                    case "":
                        var v34 uint8 = 128
                        _ = v34
                    }
                    if ((1e10 - float64(1e10)) >= 1e10) {
                        v6 += v6
                        var v35 uint8 = 1
                        v35 = uint8(v4)
                        _ = v35
                    }
                    _ = fi31
                    _ = v32
                }
            case 127:
                for fi36 := 0; fi36 < 15; fi36++ {
                    var v37 float64 = 1e10
                    switch "foo bar" {
                    case "hello":
                        var v38 uint32 = 633
                        _ = v38
                    case "foo bar":
                        switch 0 {
                        case 128:
                            if v3 {
                                switch "" {
                                case "hello":
                                    fmt.Println(v3, v4)
                                default:
                                    for fi39 := 0; fi39 < 16; fi39++ {
                                        switch (fi39 + (fi36 ^ fi36)) {
                                        case 42:
                                            if (!v3) {
                                                var v40 int64 = genFn1(-32768, 18446744073709551615, 1)
                                                v41 := GenS1{F0: v4, F1: 154, F2: v6}
                                                if ("a" >= "") {
                                                    if v3 {
                                                        v4 += 0
                                                    } else {
                                                        var v42 float64 = float64(((-2.25 / 1.0) * 0.1))
                                                        _ = v42
                                                    }
                                                } else {
                                                    switch ("\n\t" + ("foo bar" + (("\n\t" + "") + (("hello" + (("" + (("a" + "a") + ("foo bar" + "a"))) + ("hello" + ("hello" + "")))) + "a")))) {
                                                    case "hello":
                                                        v43 := GenS4{F0: 255, F1: v40, F2: 1}
                                                        _ = v43
                                                    case "\n\t":
                                                        v37 -= v37
                                                    default:
                                                        switch fi36 {
                                                        case 128:
                                                            switch "" {
                                                            case "hello":
                                                                if (false || ((!false) && false)) {
                                                                    if len(v1) > 0 {
                                                                        v1[0] = 891
                                                                    }
                                                                    fmt.Println(v3)
                                                                } else {
                                                                    var v44 bool = true
                                                                    var v45 int16 = -1
                                                                    _ = v44
                                                                    _ = v45
                                                                }
                                                            case "\n\t":
                                                                v37 = (0.1 / v37)
                                                            }
                                                        case 9223372036854775807:
                                                            if (v37 < v37) {
                                                                var v46 int64 = (^(v40))
                                                                _ = v46
                                                            } else {
                                                                if false {
                                                                    v40 = genFn1(311, 255, 2)
                                                                }
                                                                if v3 {
                                                                    var v47 uint8 = 206
                                                                    fmt.Println(v3, fi36, fi39)
                                                                    _ = v47
                                                                }
                                                                switch int(v40) {
                                                                case 751:
                                                                    var v48 uint16 = v4
                                                                    _ = v48
                                                                }
                                                            }
                                                            switch (fi39 << ((2) & 63)) {
                                                            case 42:
                                                                v49 := GenS4{F0: 2, F1: (-(v40)), F2: 2}
                                                                _ = v49
                                                            default:
                                                                if v3 {
                                                                    if v3 {
                                                                        var v50 int32 = -682
                                                                        _ = v50
                                                                    } else {
                                                                        switch "a" {
                                                                        case "a":
                                                                            v51 := GenS5{F0: 700, F1: 127}
                                                                            _ = v51
                                                                        default:
                                                                            if (65535 != 0) {
                                                                                var v52 uint8 = 1
                                                                                _ = v52
                                                                            } else {
                                                                                v40 = (v40 >> ((42) & 63))
                                                                                v53 := GenS6{F0: 1000, F1: -1, F2: 542}
                                                                                v54 := GenS3{F0: 18446744073709551615}
                                                                                _ = v53
                                                                                _ = v54
                                                                            }
                                                                        }
                                                                        v55 := GenS7{F0: 10, F1: 490}
                                                                        _ = v55
                                                                    }
                                                                    fmt.Println(fi39)
                                                                } else {
                                                                    v6 = v6
                                                                    v56 := GenS8{F0: 22}
                                                                    _ = v56
                                                                }
                                                            }
                                                            var v57 uint32 = 985
                                                            var v58 int32 = int32(v57)
                                                            var v59 float32 = ((1e10 + -2.25) / float32((v37 + (-2.25 * float64(1.5)))))
                                                            _ = v57
                                                            _ = v58
                                                            _ = v59
                                                        case -9223372036854775808:
                                                            v60 := GenS9{F0: 1}
                                                            _ = v60
                                                        default:
                                                            v61 := GenS10{F0: 90}
                                                            _ = v61
                                                        }
                                                    }
                                                    switch (fi36 * genFn0()) {
                                                    case 127:
                                                        v62 := GenS11{F0: (v4 + 1), F1: 42, F2: 0}
                                                        _ = v62
                                                    case 9223372036854775807:
                                                        v6 = 0
                                                    }
                                                }
                                                _ = len(v5)
                                                v63 := GenS12{F0: v40}
                                                _ = v40
                                                _ = v41
                                                _ = v63
                                            }
                                        case -2147483648:
                                            if (fi36 <= 1) {
                                                if (true && (0 <= 42)) {
                                                    fmt.Println(fi36, fi39)
                                                    v4 -= 32767
                                                    fmt.Println(v37, v37, v6)
                                                }
                                                v64 := GenS13{F0: (v4 >> ((2) & 63)), F1: 0, F2: 1}
                                                v6 = -1
                                                _ = v64
                                            } else {
                                                var v65 uint64 = 280
                                                var v66 bool = v3
                                                switch "\n\t" {
                                                case "\n\t":
                                                    _ = len(v1)
                                                default:
                                                    v67 := GenS11{F0: 751, F1: 430, F2: 2147483648}
                                                    _ = v67
                                                }
                                                v68 := GenS14{F0: 504}
                                                if v66 {
                                                    v69 := GenS9{F0: 2}
                                                    if (963 <= 0) {
                                                        var v70 int64 = int64(v6)
                                                        var v71 uint32 = 2147483648
                                                        v72 := GenS2{F0: -19}
                                                        _ = v70
                                                        _ = v71
                                                        _ = v72
                                                    }
                                                    fmt.Println(v6, v3)
                                                    var v73 float64 = 0.1
                                                    _ = len(v2)
                                                    _ = v69
                                                    _ = v73
                                                }
                                                var v74 float32 = -2.25
                                                _ = v65
                                                _ = v66
                                                _ = v68
                                                _ = v74
                                            }
                                        default:
                                            v75 := GenS5{F0: v4, F1: (v6 / ((v6) | 1))}
                                            _ = v75
                                        }
                                        if (-1 == int32(fi39)) {
                                            if v3 {
                                                fmt.Println(v6, fi39)
                                                switch int(v6) {
                                                case 32767:
                                                    if false {
                                                        v37 -= (v37 / float64(float32((float64(0.0) + v37))))
                                                    } else {
                                                        fmt.Println(fi39)
                                                    }
                                                case -2147483648:
                                                    var v76 float32 = (1e10 - (float32(((float64(1e10) + v37) + v37)) + 1e10))
                                                    _ = v76
                                                case 127:
                                                    fmt.Println(fi36, v37, fi39)
                                                default:
                                                    var v77 float64 = (3.14159 * float64(((0.1 * -2.25) / float32(1.5))))
                                                    _ = v77
                                                }
                                            } else {
                                                v78 := GenS5{F0: v4, F1: v6}
                                                v4 = v4
                                                _ = v78
                                            }
                                            var v79 int64 = -9223372036854775808
                                            _ = v79
                                        }
                                        _ = fi39
                                    }
                                }
                            }
                        case -323:
                            switch fi36 {
                            case 645:
                                fmt.Println(v37, fi36, v4)
                            case 32767:
                                v80 := GenS9{F0: 188}
                                _ = v80
                            default:
                                fmt.Println(v6, v6, v4)
                            }
                        default:
                            var v81 int32 = 42
                            _ = v81
                        }
                    }
                    _ = fi36
                    _ = v37
                }
            default:
                var v82 int32 = -2147483648
                _ = v82
            }
            fmt.Println(v3)
        } else {
            v3 = true
            var v83 int64 = 9223372036854775807
            if (224 != uint8(v4)) {
                v84 := GenS8{F0: (v6 & (^(v6)))}
                fmt.Println(v83)
                _ = v84
            }
            _ = v83
        }
        v4 -= 1
    } else {
        for fi85 := 0; fi85 < 17; fi85++ {
            switch -2147483648 {
            case 701:
                v5[-2147483648] = v3
            case 1:
                if v3 {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    fmt.Println(v6, v3)
                    for fi86 := 0; fi86 < 7; fi86++ {
                        v6 -= 100
                        v87 := GenS15{F0: 1, F1: 0, F2: 822}
                        _ = fi86
                        _ = v87
                    }
                }
            case 32767:
                for fi88 := 0; fi88 < 13; fi88++ {
                    v2[127] = "\n\t"
                    switch "foo bar" {
                    case "a":
                        switch (fi85 | (fi88 & fi85)) {
                        case 32768:
                            fmt.Println(v6, v3)
                        default:
                            fmt.Println(v3)
                        }
                    default:
                        var v89 uint = 262
                        _ = v89
                    }
                    fmt.Println(v4, v4)
                    if true {
                        var v90 uint16 = (v4 | v4)
                        if len(v1) > 0 {
                            v1[0] = 42
                        }
                        v3 = false
                        _ = v90
                    }
                    _ = fi88
                }
            }
            _ = fi85
        }
        if (1e10 == (0.1 * 3.14159)) {
            v3 = v3
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
        switch "hello" {
        case "a":
            switch "hello" {
            case "foo bar":
                var v91 uint = 124
                _ = v91
            }
        default:
            for fi92 := 0; fi92 < 11; fi92++ {
                var v93 uint16 = v4
                for fi94 := 0; fi94 < 20; fi94++ {
                    v6 = v6
                    _ = fi94
                }
                switch 387 {
                case 32767:
                    for fi95 := 0; fi95 < 3; fi95++ {
                        v3 = true
                        _ = fi95
                    }
                }
                _ = fi92
                _ = v93
            }
        }
    }
    _ = len(v5)
    for fi96 := 0; fi96 < 18; fi96++ {
        var v97 int = (fi96 | (fi96 + fi96))
        for fi98 := 0; fi98 < 2; fi98++ {
            v99 := GenS11{F0: v4, F1: 4294967295, F2: 34}
            v4 += (v4 << ((562) & 63))
            _ = fi98
            _ = v99
        }
        _ = fi96
        _ = v97
    }
    if (-850 == 1) {
        if false {
            switch ("" + "hello") {
            case "hello":
                for fi100 := 0; fi100 < 6; fi100++ {
                    var v101 bool = v3
                    v102 := GenS16{F0: 2147483648, F1: (fi100 * 4294967295), F2: 4294967295}
                    var v103 int32 = 2147483647
                    _ = fi100
                    _ = v101
                    _ = v102
                    _ = v103
                }
            }
            fmt.Println(v3, v4, v4)
            for fi104 := 0; fi104 < 1; fi104++ {
                switch -167 {
                case 127:
                    switch (^(fi104)) {
                    case -1:
                        var v105 uint8 = 1
                        _ = v105
                    case 64:
                        for fi106 := 0; fi106 < 18; fi106++ {
                            var v107 int64 = -1
                            var v108 uint8 = 128
                            _ = fi106
                            _ = v107
                            _ = v108
                        }
                    case 9223372036854775807:
                        if len(v1) > 0 {
                            v1[0] = 2147483648
                        }
                    default:
                        for fi109 := 0; fi109 < 17; fi109++ {
                            fmt.Println(v6, fi104)
                            if v3 {
                                fmt.Println(v6, fi104)
                                var v110 int64 = 0
                                v3 = true
                                _ = v110
                            }
                            v6 = v6
                            _ = fi109
                        }
                    }
                case -9223372036854775808:
                    switch "foo bar" {
                    case "foo bar":
                        v5[fi104] = true
                    }
                case 4294967295:
                    _ = len(v1)
                default:
                    var v111 uint = 42
                    _ = v111
                }
                _ = fi104
            }
        }
    } else {
        v112 := GenS16{F0: 65535, F1: 42, F2: 555}
        for fi113 := 0; fi113 < 9; fi113++ {
            var v114 float64 = 0.0
            if v3 {
                _ = len(v1)
            } else {
                v6 += 100
            }
            var v115 int16 = 853
            _ = fi113
            _ = v114
            _ = v115
        }
        fmt.Println(v3, v3, v4)
        _ = v112
    }
    var v116 uint8 = 250
    v6 -= v6
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v116:", v116)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v116
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
