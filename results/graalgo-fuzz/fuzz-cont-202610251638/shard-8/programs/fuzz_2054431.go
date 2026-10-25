package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32 }
type GenS2 struct { F0 uint32; F1 uint16; F2 int8 }
type GenS3 struct { F0 string; F1 int8 }
type GenS4 struct { F0 uint; F1 int32 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint32; F1 uint; F2 uint8 }
type GenS7 struct { F0 string; F1 uint32; F2 string }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 int64; F1 int16 }
type GenS10 struct { F0 string; F1 int16; F2 bool }
type GenS11 struct { F0 int32 }
type GenS12 struct { F0 uint8; F1 uint8 }
type GenS13 struct { F0 bool; F1 uint8; F2 uint32 }
type GenS14 struct { F0 int8; F1 int8 }
type GenS15 struct { F0 uint64; F1 int64 }
type GenS16 struct { F0 uint64 }
type GenS17 struct { F0 uint32 }
type GenS18 struct { F0 string; F1 uint; F2 int64 }
type GenS19 struct { F0 int }
type GenS20 struct { F0 uint; F1 int64; F2 int16 }
type GenS21 struct { F0 uint8; F1 bool }
func genFn0() uint16 {
    var v4 int = -372
    if (198 <= 373) {
        if true {
            v5 := GenS1{F0: (v4 >> ((65535) & 63)), F1: 0}
            _ = v5
        } else {
            v6 := []int{(v4 * (v4 << ((2) & 63))), v4}
            _ = len(v6)
            v7 := GenS1{F0: 32768, F1: 639}
            _ = v6
            _ = v7
        }
    } else {
        if ((^(v4)) > v4) {
            v4 -= -1
            for fi8 := 0; fi8 < 19; fi8++ {
                var v9 int = (v4 << ((2) & 63))
                v4 -= (v4 & (fi8 - v9))
                v4 -= v9
                _ = fi8
                _ = v9
            }
        }
        v4 = (v4 / ((v4) | 1))
    }
    var v10 uint64 = 1
    var v11 uint32 = 65535
    _ = v10
    _ = v11
    _ = v4
    return 101
}
func genFn1() uint16 {
    var v12 uint = 255
    var v13 uint16 = 42
    v14 := []int{0}
    if len(v14) > 0 {
        v14[0] = 2147483647
    }
    _ = v12
    _ = v13
    _ = v14
    return (v13 << (((v12 & (v12 << ((v12) & 63)))) & 63))
}

func main() {
    v1 := []string{"\n\t", "foo bar", "hello", "a"}
    var v2 uint16 = 65535
    var v3 string = "hello"
    var v15 uint = 255
    for fi16 := 0; fi16 < 18; fi16++ {
        switch fi16 {
        case 9223372036854775807:
            switch (v3 + "\n\t") {
            case "\n\t":
                for fi17 := 0; fi17 < 18; fi17++ {
                    switch (v3 + v3) {
                    case "hello":
                        _ = len(v1)
                    case "":
                        if false {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            fmt.Println(v2, v15, fi16)
                            v18 := GenS1{F0: fi17, F1: 231}
                            var v19 int32 = 42
                            var v20 float64 = 1e10
                            _ = v18
                            _ = v19
                            _ = v20
                        }
                    default:
                        var v21 uint16 = (v2 >> ((v15) & 63))
                        _ = v21
                    }
                    fmt.Println(v3, fi17)
                    var v22 int8 = -13
                    _ = fi17
                    _ = v22
                }
            case "hello":
                v1 = append(v1, "a")
            case "foo bar":
                for fi23 := 0; fi23 < 6; fi23++ {
                    if (fi23 >= (^(fi23))) {
                        var v24 uint64 = 2
                        v25 := GenS1{F0: 0, F1: -133}
                        _ = v24
                        _ = v25
                    } else {
                        v15 += (v15 % ((v15) | 1))
                        v3 = (v3 + "foo bar")
                        v26 := GenS2{F0: 2, F1: genFn1(), F2: 47}
                        fmt.Println(fi23)
                        if true {
                            v1 = append(v1, "foo bar")
                        } else {
                            v15 = 255
                            fmt.Println(v3)
                            if (true || false) {
                                fmt.Println(v3, v2)
                                var v27 uint32 = 0
                                _ = v27
                            }
                            var v28 uint = (v15 + 4294967295)
                            v2 = 297
                            _ = v28
                        }
                        switch fi23 {
                        case 180:
                            if false {
                                v1 = append(v1, "")
                                fmt.Println(v3)
                                v29 := GenS2{F0: 2, F1: (v2 + (v2 << (((v15 ^ v15)) & 63))), F2: int8(v2)}
                                v1 = append(v1, ("foo bar" + v3))
                                var v30 uint64 = 0
                                _ = v29
                                _ = v30
                            } else {
                                fmt.Println(v3, fi16, v2)
                            }
                        default:
                            var v31 string = "a"
                            var v32 uint8 = 1
                            var v33 int16 = -32768
                            if false {
                                var v34 string = (v31 + (v31 + v31))
                                if len(v1) > 0 {
                                    v1[0] = v31
                                }
                                v32 -= (^(v32))
                                _ = v34
                            } else {
                                if false {
                                    if false {
                                        fmt.Println(v32)
                                        var v35 uint64 = uint64(v2)
                                        _ = v35
                                    } else {
                                        var v36 uint16 = (v2 / ((v2) | 1))
                                        v1 = append(v1, v3)
                                        if true {
                                            var v37 uint8 = (v32 - (^(v32)))
                                            var v38 int = 42
                                            var v39 uint8 = v37
                                            _ = v37
                                            _ = v38
                                            _ = v39
                                        }
                                        v31 = v3
                                        var v40 uint32 = 777
                                        _ = v36
                                        _ = v40
                                    }
                                    switch (^(fi16)) {
                                    case -1:
                                        if true {
                                            v41 := GenS1{F0: -2147483648, F1: -1}
                                            switch v31 {
                                            case "a":
                                                var v42 int16 = (-(v33))
                                                _ = v42
                                            case "":
                                                v32 -= (v32 & (^(v32)))
                                            case "\n\t":
                                                var v43 uint = v15
                                                _ = v43
                                            }
                                            if (0 == 837) {
                                                v3 = ("hello" + ("hello" + v3))
                                            } else {
                                                var v44 uint = (v15 << ((v15) & 63))
                                                var v45 string = v3
                                                _ = v44
                                                _ = v45
                                            }
                                            _ = v41
                                        } else {
                                            switch "hello" {
                                            case "foo bar":
                                                fmt.Println(v15, fi16)
                                            case "":
                                                v46 := GenS2{F0: 2, F1: v2, F2: 127}
                                                _ = v46
                                            case "\n\t":
                                                var v47 int16 = (v33 & (v33 / ((v33) | 1)))
                                                _ = v47
                                            }
                                            if true {
                                                if false {
                                                    if false {
                                                        v15 = v15
                                                        v48 := GenS3{F0: v31, F1: 0}
                                                        var v49 uint = (v15 % (((^(v15))) | 1))
                                                        var v50 int = -570
                                                        _ = v48
                                                        _ = v49
                                                        _ = v50
                                                    } else {
                                                        if (v32 >= uint8(v15)) {
                                                            switch "" {
                                                            case "hello":
                                                                v32 = 0
                                                            case "":
                                                                var v51 string = (v31 + v3)
                                                                _ = v51
                                                            case "foo bar":
                                                                if true {
                                                                    switch (v31 + (v3 + v31)) {
                                                                    case "hello":
                                                                        fmt.Println(v15, v33, v3)
                                                                        v2 -= (^(v2))
                                                                        var v52 int32 = 1
                                                                        _ = v52
                                                                    case "a":
                                                                        if true {
                                                                            v2 = (v2 % ((v2) | 1))
                                                                        } else {
                                                                            var v53 uint16 = v2
                                                                            var v54 int8 = 0
                                                                            var v55 uint = 4294967295
                                                                            _ = v53
                                                                            _ = v54
                                                                            _ = v55
                                                                        }
                                                                    default:
                                                                        if false {
                                                                            switch v3 {
                                                                            case "\n\t":
                                                                                v2 = v2
                                                                            case "a":
                                                                                v3 = ((((v3 + "") + ((v31 + (v31 + v31)) + v31)) + (((v3 + ("hello" + "hello")) + v3) + "foo bar")) + "\n\t")
                                                                            }
                                                                            if len(v1) > 0 {
                                                                                v1[0] = "a"
                                                                            }
                                                                            v31 = "a"
                                                                            if false {
                                                                                v3 = (v31 + v31)
                                                                                v15 = (v15 + 292)
                                                                            } else {
                                                                                var v56 float64 = float64(0.0)
                                                                                _ = v56
                                                                            }
                                                                            v57 := GenS4{F0: uint(v2), F1: 2147483647}
                                                                            _ = v57
                                                                        }
                                                                    }
                                                                } else {
                                                                    var v58 int = fi16
                                                                    v59 := GenS5{F0: -1}
                                                                    v32 = (^(v32))
                                                                    _ = v58
                                                                    _ = v59
                                                                }
                                                            }
                                                            fmt.Println(v3, v32, v2)
                                                            switch (fi23 + (fi23 >> (((v15 % ((v15) | 1))) & 63))) {
                                                            case 619:
                                                                fmt.Println(v32)
                                                            }
                                                        } else {
                                                            v60 := GenS6{F0: 2147483647, F1: uint(v32), F2: 2}
                                                            fmt.Println(v2)
                                                            _ = v60
                                                        }
                                                        var v61 uint8 = (^(v32))
                                                        if len(v1) >= 2 {
                                                            v1 = v1[1:len(v1)]
                                                        }
                                                        v31 = v31
                                                        _ = v61
                                                    }
                                                } else {
                                                    v15 += 590
                                                }
                                                var v62 uint64 = 18446744073709551615
                                                _ = v62
                                            }
                                        }
                                    case 0:
                                        switch v31 {
                                        case "\n\t":
                                            v63 := GenS7{F0: v3, F1: 2, F2: v3}
                                            v64 := GenS8{F0: -1}
                                            switch v31 {
                                            case "hello":
                                                var v65 int = 42
                                                _ = v65
                                            case "":
                                                v33 -= 42
                                                switch fi16 {
                                                case 0:
                                                    v31 = (("a" + ("foo bar" + v3)) + (v31 + "foo bar"))
                                                default:
                                                    v66 := GenS3{F0: "hello", F1: 127}
                                                    _ = v66
                                                }
                                                var v67 int64 = 877
                                                switch v31 {
                                                case "hello":
                                                    var v68 uint32 = 994
                                                    _ = v68
                                                }
                                                _ = v67
                                            }
                                            _ = v63
                                            _ = v64
                                        case "hello":
                                            switch v3 {
                                            case "hello":
                                                if (1e10 == 0.0) {
                                                    _ = len(v1)
                                                    switch 628 {
                                                    case 127:
                                                        var v69 bool = false
                                                        _ = v69
                                                    case 4294967295:
                                                        switch (v31 + "hello") {
                                                        case "":
                                                            v15 += 2
                                                        default:
                                                            v70 := GenS9{F0: -899, F1: -32768}
                                                            _ = v70
                                                        }
                                                    case 72:
                                                        _ = len(v1)
                                                    default:
                                                        if ((fi23 % ((32767) | 1)) != fi23) {
                                                            var v71 int8 = 12
                                                            _ = v71
                                                        } else {
                                                            fmt.Println(v33, v15)
                                                        }
                                                    }
                                                    if (!true) {
                                                        var v72 float32 = 1e10
                                                        _ = len(v1)
                                                        _ = v72
                                                    }
                                                    var v73 float64 = (0.1 - (0.1 * 0.0))
                                                    _ = v73
                                                } else {
                                                    v1 = append(v1, "a")
                                                    if len(v1) > 0 {
                                                        v1[0] = ""
                                                    }
                                                    v32 = (^(v32))
                                                }
                                            case "":
                                                var v74 string = v31
                                                _ = v74
                                            case "foo bar":
                                                v32 = (^(v32))
                                            }
                                        case "":
                                            var v75 int64 = -9223372036854775808
                                            _ = v75
                                        default:
                                            v1 = append(v1, v31)
                                        }
                                    case -2147483648:
                                        if len(v1) > 0 {
                                            v1[0] = v31
                                        }
                                    default:
                                        v76 := GenS8{F0: -285}
                                        _ = v76
                                    }
                                }
                            }
                            _ = v31
                            _ = v32
                            _ = v33
                        }
                        v15 -= (^(v15))
                        _ = v26
                    }
                    v15 += (v15 ^ (v15 << (((v15 << ((65535) & 63))) & 63)))
                    _ = fi23
                }
            default:
                var v77 uint8 = 197
                _ = v77
            }
        default:
            switch v3 {
            case "":
                fmt.Println(v15)
            default:
                if (-863 != 1000) {
                    switch v3 {
                    case "a":
                        for fi78 := 0; fi78 < 9; fi78++ {
                            v79 := GenS4{F0: (v15 & 255), F1: 0}
                            v3 = "hello"
                            v80 := GenS2{F0: 0, F1: 0, F2: 36}
                            _ = fi78
                            _ = v79
                            _ = v80
                        }
                    case "foo bar":
                        for fi81 := 0; fi81 < 19; fi81++ {
                            var v82 uint = (v15 - 1)
                            var v83 int8 = -128
                            if true {
                                switch 38 {
                                case 2147483647:
                                    switch ((v3 + v3) + v3) {
                                    case "hello":
                                        v84 := GenS6{F0: 401, F1: 676, F2: 0}
                                        _ = v84
                                    case "":
                                        if true {
                                            v82 += (v15 * uint(v2))
                                        }
                                    case "foo bar":
                                        if len(v1) > 0 {
                                            v1[0] = ""
                                        }
                                    default:
                                        if false {
                                            v85 := GenS8{F0: -1}
                                            if false {
                                                v3 = (v3 + v3)
                                            }
                                            switch (v3 + "hello") {
                                            case "a":
                                                if true {
                                                    v82 += v82
                                                    v86 := GenS10{F0: ("foo bar" + v3), F1: 994, F2: (false || (((v83 / (((v83 | v83)) | 1)) >= (v83 << (((v82 * (v15 << (((v82 + (v15 >> ((v82) & 63)))) & 63)))) & 63))) && false))}
                                                    if (true || (!false)) {
                                                        switch (^(fi81)) {
                                                        case 1:
                                                            if (!("" >= v3)) {
                                                                v87 := GenS11{F0: 1}
                                                                var v88 uint = v82
                                                                _ = v87
                                                                _ = v88
                                                            } else {
                                                                switch v3 {
                                                                case "\n\t":
                                                                    var v89 float64 = (0.1 * (1.5 / 1.0))
                                                                    _ = v89
                                                                case "":
                                                                    v3 = (v3 + v3)
                                                                }
                                                            }
                                                        case -111:
                                                            fmt.Println(v82, v3, v82)
                                                        }
                                                        var v90 int64 = 0
                                                        _ = v90
                                                    }
                                                    switch v3 {
                                                    case "a":
                                                        var v91 int = 9223372036854775807
                                                        _ = v91
                                                    case "hello":
                                                        fmt.Println(v15, v2)
                                                    case "\n\t":
                                                        if (((0.1 / 1.5) / 1.0) != 0.1) {
                                                            var v92 float64 = 0.1
                                                            v3 = "hello"
                                                            _ = v92
                                                        } else {
                                                            switch (v3 + "\n\t") {
                                                            case "":
                                                                var v93 int = fi81
                                                                _ = v93
                                                            }
                                                            fmt.Println(v3, v2)
                                                            v15 = (^(v15))
                                                            _ = len(v1)
                                                            _ = len(v1)
                                                            _ = len(v1)
                                                            v94 := GenS12{F0: 127, F1: 255}
                                                            v95 := GenS5{F0: 0}
                                                            _ = v94
                                                            _ = v95
                                                        }
                                                    default:
                                                        if false {
                                                            v96 := GenS3{F0: "hello", F1: -74}
                                                            v2 -= v2
                                                            _ = v96
                                                        }
                                                        fmt.Println(v3, v2)
                                                        var v97 uint = 803
                                                        _ = v97
                                                    }
                                                    switch fi81 {
                                                    case 4294967295:
                                                        var v98 int64 = -9223372036854775808
                                                        if (!false) {
                                                            v99 := GenS1{F0: 128, F1: 887}
                                                            _ = v99
                                                        }
                                                        _ = v98
                                                    default:
                                                        if len(v1) >= 2 {
                                                            v1 = v1[1:len(v1)]
                                                        }
                                                    }
                                                    v82 = (v15 % (((v82 >> (((v15 * 719)) & 63))) | 1))
                                                    _ = v86
                                                }
                                            default:
                                                v2 = v2
                                            }
                                            _ = v85
                                        }
                                    }
                                    if true {
                                        switch fi16 {
                                        case 128:
                                            v15 = v82
                                            var v100 uint16 = (v2 * (^(v2)))
                                            if len(v1) > 0 {
                                                v1[0] = ("" + v3)
                                            }
                                            fmt.Println(v15)
                                            if len(v1) > 0 {
                                                v1[0] = "a"
                                            }
                                            _ = v100
                                        case -781:
                                            fmt.Println(fi81, v15, v82)
                                        }
                                        v15 = v15
                                        v101 := GenS2{F0: 342, F1: uint16(v82), F2: (v83 >> (((v15 << ((v82) & 63))) & 63))}
                                        _ = v101
                                    } else {
                                        if (!(0 > 487)) {
                                            if ((!(-2.25 >= 0.1)) && false) {
                                                if ((^(v15)) <= v82) {
                                                    var v102 int8 = -128
                                                    _ = v102
                                                } else {
                                                    v103 := GenS10{F0: v3, F1: 0, F2: false}
                                                    _ = v103
                                                }
                                                if ((false && (9223372036854775807 == 9223372036854775807)) && true) {
                                                    if false {
                                                        switch "\n\t" {
                                                        case "a":
                                                            switch -9223372036854775808 {
                                                            case 32768:
                                                                _ = len(v1)
                                                            case -119:
                                                                v2 = v2
                                                                v15 -= (v82 / ((v15) | 1))
                                                                var v104 uint = v82
                                                                _ = v104
                                                            default:
                                                                v105 := GenS5{F0: 42}
                                                                _ = v105
                                                            }
                                                        case "":
                                                            v1 = append(v1, "a")
                                                        }
                                                        if true {
                                                            if true {
                                                                switch -631 {
                                                                case 42:
                                                                    fmt.Println(v3, v2, v3)
                                                                }
                                                                var v106 uint64 = 42
                                                                var v107 int16 = 1
                                                                _ = v106
                                                                _ = v107
                                                            }
                                                            fmt.Println(v3, fi81)
                                                        }
                                                        if false {
                                                            switch (fi81 | -8) {
                                                            case 1:
                                                                switch fi16 {
                                                                case 128:
                                                                    var v108 string = v3
                                                                    _ = v108
                                                                case -2147483648:
                                                                    var v109 uint64 = 1
                                                                    _ = v109
                                                                default:
                                                                    if (true || true) {
                                                                        if (!true) {
                                                                            var v110 float32 = 3.14159
                                                                            v111 := GenS13{F0: true, F1: 42, F2: 42}
                                                                            var v112 float32 = -2.25
                                                                            _ = v110
                                                                            _ = v111
                                                                            _ = v112
                                                                        } else {
                                                                            v113 := GenS14{F0: v83, F1: 95}
                                                                            var v114 int16 = 1000
                                                                            var v115 string = "hello"
                                                                            var v116 int = -58
                                                                            fmt.Println(v115, v115)
                                                                            _ = v113
                                                                            _ = v114
                                                                            _ = v115
                                                                            _ = v116
                                                                        }
                                                                    }
                                                                }
                                                            case 2147483647:
                                                                switch ((v3 + v3) + "") {
                                                                case "hello":
                                                                    var v117 bool = false
                                                                    _ = v117
                                                                case "a":
                                                                    v3 = ""
                                                                    var v118 bool = (true || true)
                                                                    _ = v118
                                                                case "":
                                                                    switch "hello" {
                                                                    case "\n\t":
                                                                        v119 := GenS15{F0: 2, F1: -9223372036854775808}
                                                                        if ((fi81 ^ 128) != 9223372036854775807) {
                                                                            v120 := GenS3{F0: "hello", F1: (v83 * v83)}
                                                                            v2 += genFn1()
                                                                            _ = v120
                                                                        }
                                                                        _ = v119
                                                                    }
                                                                }
                                                                if false {
                                                                    if (!false) {
                                                                        if (128 == 2) {
                                                                            v82 = (v15 / ((0) | 1))
                                                                        } else {
                                                                            fmt.Println(v82, fi81)
                                                                            if true {
                                                                                if false {
                                                                                    if len(v1) > 0 {
                                                                                        v1[0] = "a"
                                                                                    }
                                                                                    fmt.Println(v2, v15)
                                                                                }
                                                                            } else {
                                                                                v2 -= (v2 % (((v2 + (v2 | v2))) | 1))
                                                                                var v121 uint64 = 42
                                                                                v121 = 0
                                                                                _ = v121
                                                                            }
                                                                            if (0.1 < -2.25) {
                                                                                v15 += 1
                                                                                var v122 uint16 = (v2 | genFn0())
                                                                                _ = v122
                                                                            }
                                                                        }
                                                                        if false {
                                                                            if false {
                                                                                fmt.Println(v15)
                                                                                if (-291 <= -881) {
                                                                                    fmt.Println(v3, v15)
                                                                                }
                                                                            } else {
                                                                                v123 := GenS15{F0: 931, F1: -1}
                                                                                switch 42 {
                                                                                case 1:
                                                                                    switch fi16 {
                                                                                    case 0:
                                                                                        var v124 int16 = 32767
                                                                                        _ = v124
                                                                                    case 1:
                                                                                        var v125 uint64 = 0
                                                                                        _ = v125
                                                                                    default:
                                                                                        var v126 string = v3
                                                                                        _ = v126
                                                                                    }
                                                                                }
                                                                                var v127 int = fi81
                                                                                _ = v123
                                                                                _ = v127
                                                                            }
                                                                            v128 := GenS1{F0: (fi81 ^ 502), F1: int32(v15)}
                                                                            v129 := GenS2{F0: uint32(v2), F1: (^(v2)), F2: v83}
                                                                            _ = v128
                                                                            _ = v129
                                                                        } else {
                                                                            var v130 int = 32768
                                                                            fmt.Println(v83, fi81, v3)
                                                                            var v131 int64 = -833
                                                                            _ = v130
                                                                            _ = v131
                                                                        }
                                                                    } else {
                                                                        switch v3 {
                                                                        case "foo bar":
                                                                            var v132 int32 = -297
                                                                            _ = v132
                                                                        case "\n\t":
                                                                            v133 := GenS7{F0: "foo bar", F1: 693, F2: v3}
                                                                            v134 := GenS1{F0: fi81, F1: 1}
                                                                            var v135 int8 = (v83 >> ((v15) & 63))
                                                                            v136 := GenS16{F0: 1}
                                                                            _ = v133
                                                                            _ = v134
                                                                            _ = v135
                                                                            _ = v136
                                                                        case "a":
                                                                            if false {
                                                                                var v137 uint32 = 2147483647
                                                                                _ = v137
                                                                            } else {
                                                                                v15 = v82
                                                                                switch "hello" {
                                                                                case "hello":
                                                                                    var v138 uint8 = 1
                                                                                    _ = v138
                                                                                case "foo bar":
                                                                                    v82 = v82
                                                                                case "":
                                                                                    fmt.Println(v83)
                                                                                }
                                                                            }
                                                                        default:
                                                                            if (!(657 != 18446744073709551615)) {
                                                                                v139 := GenS17{F0: 42}
                                                                                if true {
                                                                                    var v140 uint8 = 135
                                                                                    v141 := GenS17{F0: 42}
                                                                                    _ = v140
                                                                                    _ = v141
                                                                                } else {
                                                                                    v15 = v82
                                                                                }
                                                                                var v142 int = fi81
                                                                                _ = v139
                                                                                _ = v142
                                                                            } else {
                                                                                v82 -= v15
                                                                                switch fi81 {
                                                                                case 0:
                                                                                    v82 = v15
                                                                                }
                                                                            }
                                                                            v3 = "\n\t"
                                                                        }
                                                                        if ((fi16 >> ((4294967295) & 63)) < (fi16 ^ (fi81 - 448))) {
                                                                            switch v3 {
                                                                            case "hello":
                                                                                v82 -= 272
                                                                            }
                                                                            _ = len(v1)
                                                                        }
                                                                        var v143 bool = true
                                                                        _ = v143
                                                                    }
                                                                }
                                                                if false {
                                                                    v1 = append(v1, "foo bar")
                                                                    if len(v1) > 0 {
                                                                        v1[0] = "hello"
                                                                    }
                                                                    if true {
                                                                        v144 := GenS15{F0: 2, F1: 9223372036854775807}
                                                                        _ = v144
                                                                    }
                                                                }
                                                            case 127:
                                                                v15 += v82
                                                            default:
                                                                v83 = v83
                                                            }
                                                        }
                                                    } else {
                                                        if false {
                                                            v145 := GenS6{F0: 42, F1: v82, F2: 65}
                                                            _ = v145
                                                        } else {
                                                            var v146 string = (((v3 + ("" + v3)) + v3) + (v3 + "\n\t"))
                                                            _ = len(v1)
                                                            var v147 uint16 = (^(v2))
                                                            v3 = "a"
                                                            _ = v146
                                                            _ = v147
                                                        }
                                                        if true {
                                                            v148 := GenS6{F0: 2147483647, F1: (v15 & (^(v15))), F2: 127}
                                                            var v149 bool = false
                                                            v83 += (v83 + -107)
                                                            _ = v148
                                                            _ = v149
                                                        } else {
                                                            var v150 uint64 = 18446744073709551615
                                                            v15 = 1
                                                            fmt.Println(v3, v2, fi16)
                                                            _ = v150
                                                        }
                                                        switch (fi16 % (((-(fi16))) | 1)) {
                                                        case -1:
                                                            var v151 uint = 1
                                                            _ = v151
                                                        case 127:
                                                            v152 := GenS2{F0: 42, F1: 453, F2: (v83 * (v83 + -59))}
                                                            _ = v152
                                                        case 540:
                                                            fmt.Println(v15, v2, v82)
                                                        default:
                                                            if false {
                                                                v15 -= v82
                                                            } else {
                                                                fmt.Println(v3)
                                                                if false {
                                                                    var v153 int64 = 9223372036854775807
                                                                    _ = v153
                                                                } else {
                                                                    _ = len(v1)
                                                                    v154 := GenS18{F0: (v3 + v3), F1: 215, F2: 1}
                                                                    _ = v154
                                                                }
                                                            }
                                                        }
                                                    }
                                                    fmt.Println(fi81, v15)
                                                    fmt.Println(fi16)
                                                    var v155 uint64 = 0
                                                    _ = v155
                                                }
                                                var v156 string = ("hello" + v3)
                                                _ = v156
                                            }
                                            if true {
                                                if (true || false) {
                                                    fmt.Println(v15)
                                                    if (1e10 <= (3.14159 * 3.14159)) {
                                                        if (false || (uint16(v15) >= v2)) {
                                                            var v157 uint = v82
                                                            v158 := GenS8{F0: -2147483648}
                                                            if true {
                                                                fmt.Println(fi81, v2)
                                                            } else {
                                                                v2 = genFn1()
                                                                var v159 string = v3
                                                                var v160 int = -2147483648
                                                                _ = v159
                                                                _ = v160
                                                            }
                                                            _ = v157
                                                            _ = v158
                                                        } else {
                                                            var v161 float32 = float32(0.0)
                                                            v161 = ((0.0 / -2.25) - (0.1 - -2.25))
                                                            var v162 int64 = 0
                                                            _ = v161
                                                            _ = v162
                                                        }
                                                        _ = len(v1)
                                                    }
                                                    switch 127 {
                                                    case -242:
                                                        v163 := GenS1{F0: fi16, F1: -650}
                                                        _ = v163
                                                    }
                                                } else {
                                                    v164 := GenS9{F0: 42, F1: -32768}
                                                    v15 -= v15
                                                    v83 -= v83
                                                    _ = v164
                                                }
                                                var v165 uint64 = 0
                                                var v166 uint32 = 65535
                                                _ = len(v1)
                                                if len(v1) > 0 {
                                                    v1[0] = "\n\t"
                                                }
                                                var v167 int8 = v83
                                                var v168 int32 = 42
                                                _ = v165
                                                _ = v166
                                                _ = v167
                                                _ = v168
                                            } else {
                                                var v169 uint8 = 57
                                                _ = v169
                                            }
                                            if false {
                                                fmt.Println(v3, v2, v15)
                                                fmt.Println(v83, v82)
                                                fmt.Println(v83, v15)
                                            }
                                        }
                                        var v170 uint32 = 110
                                        _ = v170
                                    }
                                    fmt.Println(v15, v2)
                                case 1:
                                    fmt.Println(v82)
                                default:
                                    v171 := GenS19{F0: (fi81 << ((v82) & 63))}
                                    _ = v171
                                }
                                fmt.Println(v3)
                            } else {
                                v83 = (v83 + v83)
                                fmt.Println(v15, v83)
                                var v172 float64 = 1e10
                                _ = v172
                            }
                            _ = fi81
                            _ = v82
                            _ = v83
                        }
                    case "":
                        if len(v1) > 0 {
                            v1[0] = v3
                        }
                    }
                    var v173 uint = v15
                    _ = v173
                }
            }
        }
        _ = fi16
    }
    switch 117 {
    case -386:
        var v174 uint16 = 2
        _ = v174
    default:
        if (1 < 128) {
            var v175 int = -236
            _ = v175
        } else {
            var v176 uint = (v15 << (((^(v15))) & 63))
            _ = len(v1)
            for fi177 := 0; fi177 < 11; fi177++ {
                if ("hello" < v3) {
                    v176 = v15
                    v178 := GenS2{F0: 1, F1: v2, F2: 100}
                    v15 = (v176 % (((^(v15))) | 1))
                    _ = v178
                }
                _ = fi177
            }
            _ = v176
        }
    }
    var v179 uint32 = 2
    v3 = v3
    v180 := GenS20{F0: (^(v15)), F1: 42, F2: -675}
    fmt.Println(v3, v179, v3)
    var v181 string = "\n\t"
    fmt.Println(v2, v179, v15)
    var v182 float64 = ((1.5 / 1e10) / -2.25)
    for fi183 := 0; fi183 < 8; fi183++ {
        var v184 int16 = -1
        v185 := GenS18{F0: "hello", F1: 908, F2: 1}
        if (!true) {
            v182 += v182
            v181 = ""
            v186 := GenS21{F0: 203, F1: (((1e10 / 1.5) / 1e10) == 0.0)}
            _ = v186
        }
        _ = fi183
        _ = v184
        _ = v185
    }
    switch ((v3 + ("\n\t" + v3)) + v3) {
    case "a":
        v2 -= v2
    case "foo bar":
        v187 := GenS1{F0: -9223372036854775808, F1: 885}
        _ = v187
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v179:", v179)
    fmt.Println("v181:", v181)
    fmt.Println("v182:", v182)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v15
    _ = v179
    _ = v180
    _ = v181
    _ = v182
    _ = v2
    _ = v3
}
