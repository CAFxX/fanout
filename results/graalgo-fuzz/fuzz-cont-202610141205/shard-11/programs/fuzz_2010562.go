package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint; F2 int32 }
type GenS2 struct { F0 int64; F1 string }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint64; F1 bool; F2 uint8 }
type GenS5 struct { F0 int64; F1 uint32 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 uint16; F1 int64 }
type GenS8 struct { F0 uint; F1 string }
type GenS9 struct { F0 uint; F1 string }
type GenS10 struct { F0 bool; F1 int16 }
type GenS11 struct { F0 int64; F1 int }
type GenS12 struct { F0 int16 }
type GenS13 struct { F0 int16; F1 string }
type GenS14 struct { F0 uint; F1 int8; F2 uint }
type GenS15 struct { F0 uint64; F1 int64; F2 uint64 }
type GenS16 struct { F0 int8 }
type GenS17 struct { F0 int64; F1 uint64; F2 bool }
type GenS18 struct { F0 uint16; F1 uint8 }
type GenS19 struct { F0 int64 }
type GenS20 struct { F0 string }
type GenS21 struct { F0 uint32 }
func genFn0() bool {
    v5 := []int{1, 2147483647, 2147483647, 0}
    v6 := GenS1{F0: "hello", F1: 1, F2: 1}
    _ = v5
    _ = v6
    return (669 <= 18446744073709551615)
}
func genFn1() uint8 {
    v7 := GenS1{F0: "hello", F1: 376, F2: 2147483647}
    for fi8 := 0; fi8 < 4; fi8++ {
        var v9 float64 = 1e10
        switch "a" {
        case "foo bar":
            v10 := GenS1{F0: ("a" + "\n\t"), F1: 477, F2: 2147483647}
            _ = v10
        case "\n\t":
            var v11 uint = 4294967295
            _ = v11
        case "":
            v12 := []int{(fi8 | (fi8 >> ((1) & 63)))}
            _ = v12
        default:
            var v13 string = "foo bar"
            _ = v13
        }
        _ = fi8
        _ = v9
    }
    var v14 int = -9223372036854775808
    if ("foo bar" == "") {
        var v15 int = (^(v14))
        if (true && (!(-1 != 0))) {
            for fi16 := 0; fi16 < 11; fi16++ {
                v17 := GenS1{F0: "hello", F1: 718, F2: -26}
                var v18 int32 = -2147483648
                if ((880 <= 2147483647) || true) {
                    if false {
                        var v19 float32 = -2.25
                        v20 := []int{}
                        _ = v19
                        _ = v20
                    }
                    v21 := GenS1{F0: (("\n\t" + "hello") + "\n\t"), F1: 42, F2: v18}
                    _ = v21
                } else {
                    if true {
                        v22 := GenS2{F0: 42, F1: "a"}
                        var v23 uint8 = 20
                        _ = v22
                        _ = v23
                    } else {
                        for fi24 := 0; fi24 < 11; fi24++ {
                            switch fi24 {
                            case 127:
                                v14 = v15
                            case 368:
                                var v25 string = (("" + "\n\t") + "\n\t")
                                v15 -= (fi16 - (fi24 << ((502) & 63)))
                                if false {
                                    var v26 int16 = 715
                                    var v27 int16 = (v26 % ((v26) | 1))
                                    switch "\n\t" {
                                    case "\n\t":
                                        v28 := []int{fi24, (v15 / ((fi16) | 1))}
                                        switch "hello" {
                                        case "hello":
                                            _ = len(v28)
                                        case "\n\t":
                                            if true {
                                                switch v25 {
                                                case "a":
                                                    _ = len(v28)
                                                }
                                                if len(v28) > 0 {
                                                    v28[0] = v14
                                                }
                                                if (1 > 42) {
                                                    v15 = int(v18)
                                                } else {
                                                    if false {
                                                        switch v25 {
                                                        case "foo bar":
                                                            v29 := GenS2{F0: 1, F1: v25}
                                                            v28 = append(v28, fi24)
                                                            _ = v29
                                                        case "a":
                                                            if false {
                                                                var v30 uint64 = 4294967295
                                                                var v31 uint16 = 1
                                                                v14 += fi24
                                                                _ = v30
                                                                _ = v31
                                                            }
                                                            v32 := GenS2{F0: -612, F1: v25}
                                                            if (3.14159 <= (0.1 * float32(3.14159))) {
                                                                if true {
                                                                    var v33 int8 = 0
                                                                    v18 = (v18 | v18)
                                                                    var v34 uint64 = 2
                                                                    v35 := GenS2{F0: -563, F1: "a"}
                                                                    _ = v33
                                                                    _ = v34
                                                                    _ = v35
                                                                }
                                                                _ = len(v28)
                                                                switch v25 {
                                                                case "foo bar":
                                                                    v36 := GenS3{F0: 0}
                                                                    _ = v36
                                                                case "\n\t":
                                                                    fmt.Println(v18, v25)
                                                                }
                                                            }
                                                            _ = v32
                                                        default:
                                                            v37 := GenS3{F0: 9223372036854775807}
                                                            _ = v37
                                                        }
                                                        switch ((v25 + v25) + ("" + v25)) {
                                                        case "foo bar":
                                                            var v38 uint32 = 1
                                                            _ = v38
                                                        default:
                                                            var v39 float32 = (((0.1 / 1e10) * (-2.25 / (1e10 - (float32(3.14159) - (0.0 / (((1.5 / 1.5) - ((((0.1 + 0.1) + 0.1) / 3.14159) / 1.5)) - (-2.25 / ((-2.25 / (((((0.0 / -2.25) * 0.1) * 1.5) - (3.14159 / 0.1)) * 1.5)) + 1.5)))))))) * ((-2.25 + 1e10) + ((float32(0.0) - 1.5) * -2.25)))
                                                            _ = v39
                                                        }
                                                        fmt.Println(fi24)
                                                    }
                                                    var v40 int64 = -9223372036854775808
                                                    _ = v40
                                                }
                                            } else {
                                                v26 += v26
                                                v15 -= 0
                                                v26 = v27
                                                v18 = int32(v26)
                                                if true {
                                                    v41 := GenS4{F0: 27, F1: (4294967295 == 2), F2: 127}
                                                    switch ("hello" + v25) {
                                                    case "hello":
                                                        v14 = v14
                                                    case "":
                                                        v42 := GenS2{F0: -609, F1: v25}
                                                        _ = v42
                                                    case "a":
                                                        v43 := GenS1{F0: v25, F1: 4294967295, F2: v18}
                                                        _ = v43
                                                    default:
                                                        switch (fi16 << ((172) & 63)) {
                                                        case -1:
                                                            if len(v28) > 0 {
                                                                v28[0] = 9223372036854775807
                                                            }
                                                        case 175:
                                                            if false {
                                                                v18 = v18
                                                            } else {
                                                                if (42 < 2147483648) {
                                                                    var v44 int16 = v27
                                                                    v28 = append(v28, (fi24 + (fi16 >> ((42) & 63))))
                                                                    if len(v28) >= 2 {
                                                                        v28 = v28[1:len(v28)]
                                                                    }
                                                                    if (-524 > 9223372036854775807) {
                                                                        var v45 int = 32768
                                                                        var v46 int32 = (v18 % ((v18) | 1))
                                                                        switch (^(v45)) {
                                                                        case 128:
                                                                            v47 := GenS1{F0: v25, F1: 42, F2: 42}
                                                                            _ = v47
                                                                        default:
                                                                            var v48 uint64 = 891
                                                                            _ = v48
                                                                        }
                                                                        _ = v45
                                                                        _ = v46
                                                                    } else {
                                                                        fmt.Println(v26, fi24)
                                                                        _ = len(v28)
                                                                    }
                                                                    if (485 >= 437) {
                                                                        var v49 string = "a"
                                                                        _ = v49
                                                                    }
                                                                    switch "" {
                                                                    case "a":
                                                                        switch (v14 / ((4294967295) | 1)) {
                                                                        case 32767:
                                                                            fmt.Println(v27, v14, v18)
                                                                        case 127:
                                                                            v50 := GenS5{F0: -9223372036854775808, F1: 0}
                                                                            _ = v50
                                                                        default:
                                                                            v18 = v18
                                                                        }
                                                                        switch v25 {
                                                                        case "hello":
                                                                            if (v25 >= v25) {
                                                                                fmt.Println(v26, fi24, v44)
                                                                            } else {
                                                                                var v51 uint8 = 1
                                                                                v52 := GenS6{F0: -785}
                                                                                switch 127 {
                                                                                case 128:
                                                                                    var v53 float64 = float64(0.0)
                                                                                    _ = v53
                                                                                case 824:
                                                                                    var v54 uint64 = 599
                                                                                    _ = v54
                                                                                }
                                                                                _ = v51
                                                                                _ = v52
                                                                            }
                                                                        case "foo bar":
                                                                            var v55 float64 = ((1e10 + (3.14159 + 0.0)) / 3.14159)
                                                                            _ = v55
                                                                        }
                                                                        switch fi24 {
                                                                        case -9223372036854775808:
                                                                            switch (v25 + (("\n\t" + (v25 + ("" + v25))) + (v25 + v25))) {
                                                                            case "a":
                                                                                if len(v28) > 0 {
                                                                                    v28[0] = 4294967295
                                                                                }
                                                                            case "foo bar":
                                                                                fmt.Println(v27)
                                                                            default:
                                                                                v28 = append(v28, 1)
                                                                            }
                                                                        default:
                                                                            var v56 uint16 = 255
                                                                            v56 -= v56
                                                                            v25 = v25
                                                                            v28 = append(v28, 128)
                                                                            fmt.Println(v44, v18, fi16)
                                                                            _ = v56
                                                                        }
                                                                    case "\n\t":
                                                                        switch (fi16 + 32767) {
                                                                        case 577:
                                                                            fmt.Println(v44, v18)
                                                                        case -9223372036854775808:
                                                                            fmt.Println(v15, v14)
                                                                        case 4294967295:
                                                                            if (!(((("\n\t" + v25) + v25) + v25) == v25)) {
                                                                                if true {
                                                                                    v57 := GenS7{F0: 1, F1: 42}
                                                                                    if len(v28) >= 2 {
                                                                                        v28 = v28[1:len(v28)]
                                                                                    }
                                                                                    var v58 float64 = 1.5
                                                                                    _ = v57
                                                                                    _ = v58
                                                                                }
                                                                                _ = len(v28)
                                                                            }
                                                                        default:
                                                                            v59 := GenS7{F0: 1, F1: -70}
                                                                            _ = v59
                                                                        }
                                                                    case "foo bar":
                                                                        switch (v25 + v25) {
                                                                        case "a":
                                                                            var v60 uint32 = 65535
                                                                            _ = v60
                                                                        case "\n\t":
                                                                            v18 -= 2147483647
                                                                        }
                                                                    default:
                                                                        v61 := GenS4{F0: 4294967295, F1: (0 != 9223372036854775807), F2: 123}
                                                                        _ = v61
                                                                    }
                                                                    _ = v44
                                                                }
                                                                v18 = v18
                                                                var v62 uint16 = 152
                                                                v27 += v26
                                                                if (!(((fi24 | 563) != v14) || (true && true))) {
                                                                    fmt.Println(v25, v25, v26)
                                                                }
                                                                v26 -= (-(v26))
                                                                _ = v62
                                                            }
                                                        default:
                                                            v28 = append(v28, (v14 ^ 9223372036854775807))
                                                        }
                                                    }
                                                    _ = v41
                                                } else {
                                                    var v63 uint64 = 384
                                                    var v64 int64 = 0
                                                    var v65 uint32 = 4294967295
                                                    _ = v63
                                                    _ = v64
                                                    _ = v65
                                                }
                                            }
                                        case "foo bar":
                                            if false {
                                                v15 += v15
                                                fmt.Println(fi24)
                                                if false {
                                                    v66 := GenS4{F0: 4294967295, F1: false, F2: 109}
                                                    fmt.Println(v27, v26)
                                                    _ = v66
                                                }
                                            } else {
                                                if false {
                                                    v14 = fi24
                                                    if ((!true) || true) {
                                                        v14 = fi24
                                                        v27 -= v26
                                                        v18 = v18
                                                        if len(v28) > 0 {
                                                            v28[0] = (-(v14))
                                                        }
                                                    } else {
                                                        v15 = (v14 & v14)
                                                        switch ("foo bar" + "") {
                                                        case "foo bar":
                                                            var v67 int32 = v18
                                                            _ = v67
                                                        case "a":
                                                            v15 += 127
                                                            if false {
                                                                fmt.Println(v27, v27)
                                                            } else {
                                                                if (2 <= 2) {
                                                                    _ = len(v28)
                                                                    fmt.Println(v27, v18, fi16)
                                                                    fmt.Println(v26, v15)
                                                                }
                                                            }
                                                            if (!((fi24 | 128) == (v15 + fi16))) {
                                                                var v68 float32 = 0.0
                                                                fmt.Println(fi24, fi16, fi16)
                                                                _ = v68
                                                            } else {
                                                                fmt.Println(v14)
                                                            }
                                                        case "hello":
                                                            var v69 string = v25
                                                            _ = v69
                                                        }
                                                        v18 = v18
                                                    }
                                                }
                                                v70 := GenS7{F0: 2, F1: -1}
                                                var v71 uint8 = 255
                                                _ = v70
                                                _ = v71
                                            }
                                        default:
                                            switch "hello" {
                                            case "\n\t":
                                                fmt.Println(v26)
                                            case "":
                                                var v72 int8 = -13
                                                _ = v72
                                            }
                                        }
                                        v25 = "\n\t"
                                        _ = v28
                                    }
                                    _ = v26
                                    _ = v27
                                }
                                _ = v25
                            case -9223372036854775808:
                                v14 -= -659
                            }
                            if true {
                                var v73 int8 = -7
                                v74 := []int{(v15 / (((fi24 - (fi24 & v14))) | 1)), (fi16 + v15)}
                                fmt.Println(v73, v15)
                                _ = v73
                                _ = v74
                            } else {
                                v75 := GenS7{F0: 607, F1: -9223372036854775808}
                                if (v18 != (v18 | (v18 >> ((1) & 63)))) {
                                    var v76 int32 = v18
                                    v77 := []int{v14, (v14 >> ((0) & 63)), (fi16 >> ((2) & 63))}
                                    _ = v76
                                    _ = v77
                                } else {
                                    switch v14 {
                                    case 9223372036854775807:
                                        var v78 string = "foo bar"
                                        _ = v78
                                    case 2147483647:
                                        var v79 uint32 = 2
                                        fmt.Println(v14)
                                        _ = v79
                                    case 42:
                                        if ((fi24 ^ (v14 >> ((595) & 63))) >= (fi24 >> ((255) & 63))) {
                                            var v80 uint = 0
                                            _ = v80
                                        } else {
                                            v81 := GenS1{F0: "foo bar", F1: 0, F2: 1}
                                            v82 := GenS2{F0: 0, F1: "foo bar"}
                                            if (!false) {
                                                var v83 uint8 = 2
                                                v84 := []int{(fi16 >> ((65535) & 63)), (v15 % (((^(fi16))) | 1))}
                                                if (v83 >= (v83 ^ v83)) {
                                                    if len(v84) > 0 {
                                                        v84[0] = (v14 - fi16)
                                                    }
                                                    if (2 >= 2) {
                                                        fmt.Println(v14, v18)
                                                        var v85 float64 = -2.25
                                                        if false {
                                                            v18 -= 2147483647
                                                            var v86 uint8 = (^(v83))
                                                            fmt.Println(fi16, v85, fi16)
                                                            _ = v86
                                                        }
                                                        _ = v85
                                                    } else {
                                                        var v87 int16 = -73
                                                        if ((v15 - 2147483647) <= (fi16 + 0)) {
                                                            var v88 float64 = 0.1
                                                            switch v15 {
                                                            case 9223372036854775807:
                                                                fmt.Println(v14, fi16, v83)
                                                            case -9223372036854775808:
                                                                v89 := GenS1{F0: "\n\t", F1: 946, F2: v18}
                                                                v83 -= 0
                                                                v90 := GenS3{F0: 1}
                                                                _ = v89
                                                                _ = v90
                                                            case 530:
                                                                v14 = (v14 / ((fi16) | 1))
                                                            default:
                                                                v84 = append(v84, 9223372036854775807)
                                                            }
                                                            _ = v88
                                                        }
                                                        v91 := GenS8{F0: 360, F1: "foo bar"}
                                                        _ = v87
                                                        _ = v91
                                                    }
                                                    v84 = append(v84, (v15 + (v15 + (fi16 / ((32768) | 1)))))
                                                } else {
                                                    if true {
                                                        if (v83 == (v83 & (v83 - (v83 * (v83 & v83))))) {
                                                            switch -427 {
                                                            case 820:
                                                                var v92 int16 = -88
                                                                _ = v92
                                                            case 878:
                                                                switch 0 {
                                                                case 867:
                                                                    fmt.Println(v15, fi24, fi16)
                                                                case 32768:
                                                                    _ = len(v84)
                                                                default:
                                                                    var v93 float32 = 1.5
                                                                    _ = v93
                                                                }
                                                                switch (^(fi24)) {
                                                                case -1:
                                                                    var v94 int64 = -491
                                                                    _ = v94
                                                                }
                                                                fmt.Println(fi16, v15)
                                                            case 64:
                                                                fmt.Println(v18)
                                                            }
                                                            v95 := GenS9{F0: 4294967295, F1: "a"}
                                                            v15 = 127
                                                            _ = v95
                                                        }
                                                        v83 -= v83
                                                    } else {
                                                        if len(v84) >= 2 {
                                                            v84 = v84[1:len(v84)]
                                                        }
                                                        if (false && true) {
                                                            var v96 uint32 = 397
                                                            v97 := GenS10{F0: (uint16(v83) != 32767), F1: 0}
                                                            _ = v96
                                                            _ = v97
                                                        } else {
                                                            v15 = v15
                                                            if len(v84) >= 2 {
                                                                v84 = v84[1:len(v84)]
                                                            }
                                                        }
                                                        if true {
                                                            if len(v84) >= 2 {
                                                                v84 = v84[1:len(v84)]
                                                            }
                                                            var v98 bool = (("foo bar" >= "a") || (672 >= 187))
                                                            _ = v98
                                                        } else {
                                                            v83 = 42
                                                        }
                                                    }
                                                    fmt.Println(v15, v15)
                                                }
                                                _ = v83
                                                _ = v84
                                            }
                                            _ = v81
                                            _ = v82
                                        }
                                    default:
                                        switch (((("foo bar" + "foo bar") + "a") + "a") + "foo bar") {
                                        case "\n\t":
                                            fmt.Println(v18, fi16, v18)
                                        case "hello":
                                            switch (v15 % ((fi16) | 1)) {
                                            case -898:
                                                switch 127 {
                                                case 786:
                                                    v99 := GenS3{F0: 1}
                                                    _ = v99
                                                }
                                            }
                                        case "":
                                            fmt.Println(v15)
                                        }
                                    }
                                    var v100 uint = 4294967295
                                    _ = v100
                                }
                                var v101 uint16 = 65535
                                _ = v101
                                _ = v75
                            }
                            _ = fi24
                        }
                        switch (("" + (("foo bar" + "\n\t") + "a")) + ("\n\t" + "foo bar")) {
                        case "foo bar":
                            var v102 int = (v14 << ((426) & 63))
                            _ = v102
                        }
                    }
                    switch "foo bar" {
                    case "foo bar":
                        v15 = -775
                    }
                }
                _ = fi16
                _ = v17
                _ = v18
            }
            fmt.Println(v14, v15)
        }
        _ = v15
    }
    if (-35 < 115) {
        v14 -= v14
        fmt.Println(v14)
    }
    _ = v14
    _ = v7
    return 255
}

func main() {
    var v1 uint = 18446744073709551615
    v2 := make(map[int]int)
    var v3 bool = false
    v4 := make(map[string]string)
    v3 = v3
    for fi103 := 0; fi103 < 19; fi103++ {
        for fi104 := 0; fi104 < 5; fi104++ {
            fmt.Println(fi103, fi103, fi104)
            if v3 {
                v105 := []int{}
                if len(v105) > 0 {
                    v105[0] = (fi104 & 188)
                }
                if v3 {
                    var v106 uint16 = 404
                    _ = len(v4)
                    _ = v106
                } else {
                    v1 -= v1
                }
                if (244 <= 7) {
                    fmt.Println(v1)
                    fmt.Println(v1, fi104)
                } else {
                    var v107 int16 = -538
                    _ = v107
                }
                switch fi104 {
                case -825:
                    v3 = (0.0 <= (1.5 - (-2.25 - (1e10 - 1e10))))
                default:
                    switch ("" + "a") {
                    case "a":
                        var v108 int64 = 558
                        _ = v108
                    default:
                        var v109 uint64 = 1
                        v3 = v3
                        v110 := GenS3{F0: 0}
                        var v111 uint64 = v109
                        fmt.Println(v3, fi103)
                        var v112 bool = (-277 > -108)
                        if false {
                            var v113 uint32 = 986
                            var v114 int64 = -9223372036854775808
                            var v115 int16 = int16(v114)
                            _ = v113
                            _ = v114
                            _ = v115
                        }
                        v116 := GenS11{F0: 9223372036854775807, F1: fi104}
                        var v117 uint64 = (v109 >> (((v1 / ((v1) | 1))) & 63))
                        var v118 int64 = 0
                        v105 = append(v105, fi104)
                        var v119 uint32 = 0
                        _ = v109
                        _ = v110
                        _ = v111
                        _ = v112
                        _ = v116
                        _ = v117
                        _ = v118
                        _ = v119
                    }
                }
                if true {
                    v105 = append(v105, (fi104 - fi104))
                    var v120 int = fi104
                    _ = v120
                } else {
                    var v121 float64 = -2.25
                    v1 -= (^(v1))
                    _ = v121
                }
                v122 := GenS8{F0: (v1 + v1), F1: "a"}
                fmt.Println(v3, fi104)
                v1 += (v1 * v1)
                _ = v105
                _ = v122
            }
            var v123 int64 = 1
            _ = fi104
            _ = v123
        }
        _ = fi103
    }
    if (false && v3) {
        if (!false) {
            v124 := []int{}
            for fi125 := 0; fi125 < 18; fi125++ {
                v124 = append(v124, 0)
                for fi126 := 0; fi126 < 3; fi126++ {
                    if v3 {
                        if len(v124) >= 2 {
                            v124 = v124[1:len(v124)]
                        }
                        var v127 uint32 = 707
                        var v128 int32 = 0
                        _ = v127
                        _ = v128
                    } else {
                        switch ("" + ("a" + ((("hello" + "a") + ("a" + ("" + "hello"))) + ("a" + "foo bar")))) {
                        case "\n\t":
                            v3 = ((fi126 + fi125) != (fi125 << (((v1 >> ((65535) & 63))) & 63)))
                        case "":
                            v1 += (^(v1))
                        }
                        v129 := GenS12{F0: 1000}
                        if true {
                            v130 := GenS13{F0: -32768, F1: "foo bar"}
                            _ = v130
                        } else {
                            fmt.Println(v1, v1)
                            if true {
                                if true {
                                    var v131 int = (fi125 - fi126)
                                    fmt.Println(v3)
                                    v3 = v3
                                    _ = v131
                                } else {
                                    if (-2.25 == -2.25) {
                                        var v132 int = fi125
                                        _ = v132
                                    } else {
                                        var v133 int16 = 1000
                                        v133 -= 0
                                        _ = v133
                                    }
                                    if (!true) {
                                        if len(v124) > 0 {
                                            v124[0] = fi126
                                        }
                                    }
                                    v134 := GenS14{F0: 4294967295, F1: 79, F2: v1}
                                    _ = v134
                                }
                            }
                        }
                        var v135 int = fi126
                        switch (v135 << ((v1) & 63)) {
                        case -1:
                            if false {
                                v136 := GenS15{F0: 1, F1: 137, F2: 0}
                                v135 = (-(fi125))
                                var v137 int16 = 368
                                _ = v136
                                _ = v137
                            }
                        case 127:
                            if v3 {
                                v1 += (v1 + v1)
                            } else {
                                v3 = (!true)
                                var v138 uint8 = 2
                                _ = v138
                            }
                        case 1:
                            v1 = v1
                        }
                        _ = len(v2)
                        if false {
                            var v139 int = (fi125 % (((fi126 ^ (fi126 << (((v1 - (v1 * (^(v1))))) & 63)))) | 1))
                            if (2147483648 < 275) {
                                v3 = (v1 == (v1 | (v1 ^ (v1 & 2))))
                                fmt.Println(v1, fi125, v139)
                                fmt.Println(v135, v1)
                            }
                            _ = v139
                        }
                        v3 = (9223372036854775807 >= (v135 << ((18446744073709551615) & 63)))
                        _ = v129
                        _ = v135
                    }
                    v140 := GenS16{F0: 111}
                    _ = fi126
                    _ = v140
                }
                var v141 bool = ("foo bar" == ("\n\t" + ("\n\t" + ("a" + "foo bar"))))
                _ = fi125
                _ = v141
            }
            var v142 int32 = -571
            _ = v124
            _ = v142
        } else {
            switch -576 {
            case -9223372036854775808:
                for fi143 := 0; fi143 < 16; fi143++ {
                    var v144 uint = v1
                    _ = fi143
                    _ = v144
                }
            case -218:
                var v145 int32 = 838
                _ = v145
            case -741:
                for fi146 := 0; fi146 < 4; fi146++ {
                    v147 := GenS1{F0: "a", F1: v1, F2: 958}
                    _ = fi146
                    _ = v147
                }
            default:
                switch 32768 {
                case -342:
                    if v3 {
                        v3 = v3
                    }
                case -418:
                    if (((1.5 / ((1.5 * 3.14159) + -2.25)) * -2.25) >= float32((0.0 * 1e10))) {
                        if v3 {
                            if true {
                                for fi148 := 0; fi148 < 13; fi148++ {
                                    var v149 int16 = -808
                                    v4["\n\t"] = ("" + "a")
                                    var v150 bool = (-2.25 < (1e10 * 1e10))
                                    _ = fi148
                                    _ = v149
                                    _ = v150
                                }
                                if (v3 && true) {
                                    v1 += v1
                                } else {
                                    v1 += (v1 | (v1 - (v1 >> ((v1) & 63))))
                                }
                            } else {
                                var v151 float64 = 1.5
                                var v152 int8 = -128
                                _ = v151
                                _ = v152
                            }
                        } else {
                            v153 := []int{127}
                            if v3 {
                                switch 9223372036854775807 {
                                case 32767:
                                    if ((v3 || false) || false) {
                                        fmt.Println(v1)
                                        v154 := GenS17{F0: 868, F1: 988, F2: v3}
                                        if v3 {
                                            fmt.Println(v3, v3)
                                            if v3 {
                                                v155 := GenS18{F0: 70, F1: 242}
                                                _ = v155
                                            } else {
                                                v1 = (v1 + 42)
                                                for fi156 := 0; fi156 < 6; fi156++ {
                                                    v3 = (v3 && false)
                                                    fmt.Println(v3)
                                                    fmt.Println(v3)
                                                    _ = fi156
                                                }
                                            }
                                            for fi157 := 0; fi157 < 8; fi157++ {
                                                var v158 uint8 = 2
                                                _ = fi157
                                                _ = v158
                                            }
                                        } else {
                                            v159 := GenS14{F0: (v1 + v1), F1: 1, F2: 355}
                                            _ = v159
                                        }
                                        _ = v154
                                    } else {
                                        var v160 uint8 = 136
                                        fmt.Println(v3)
                                        _ = v160
                                    }
                                default:
                                    var v161 uint32 = 2147483647
                                    _ = v161
                                }
                                var v162 uint16 = 1
                                _ = v162
                            } else {
                                if len(v153) > 0 {
                                    v153[0] = 1
                                }
                                var v163 int16 = 42
                                _ = v163
                            }
                            var v164 int64 = 9223372036854775807
                            _ = v153
                            _ = v164
                        }
                    } else {
                        var v165 int8 = 1
                        _ = v165
                    }
                default:
                    for fi166 := 0; fi166 < 10; fi166++ {
                        switch 128 {
                        case 42:
                            var v167 string = "a"
                            _ = v167
                        case 127:
                            for fi168 := 0; fi168 < 14; fi168++ {
                                v1 = 65535
                                _ = fi168
                            }
                        case -2147483648:
                            if (int16(fi166) != -32768) {
                                for fi169 := 0; fi169 < 20; fi169++ {
                                    var v170 uint8 = 2
                                    _ = fi169
                                    _ = v170
                                }
                                v171 := v4["\n\t"]
                                _ = v171
                            }
                        default:
                            switch fi166 {
                            case 32767:
                                v172 := []int{}
                                _ = v172
                            }
                        }
                        _ = fi166
                    }
                }
            }
            if ((0 >= 1000) || (0 > 1)) {
                if true {
                    var v173 int64 = 1
                    var v174 uint64 = 2
                    for fi175 := 0; fi175 < 8; fi175++ {
                        v3 = (100 < 6)
                        _ = fi175
                    }
                    _ = v173
                    _ = v174
                }
                var v176 uint = (v1 & 18446744073709551615)
                _ = v176
            } else {
                fmt.Println(v1)
            }
        }
        v1 = v1
    } else {
        switch 4294967295 {
        case 128:
            v1 = 255
        case 689:
            v3 = v3
        default:
            for fi177 := 0; fi177 < 10; fi177++ {
                if (!(-128 != 6)) {
                    var v178 uint16 = 613
                    _ = v178
                }
                for fi179 := 0; fi179 < 11; fi179++ {
                    switch "a" {
                    case "a":
                        var v180 uint32 = 990
                        _ = v180
                    case "":
                        var v181 uint64 = 501
                        _ = v181
                    }
                    _ = fi179
                }
                fmt.Println(v1)
                _ = fi177
            }
        }
        var v182 int16 = 32767
        switch (("" + "") + "a") {
        case "hello":
            if v3 {
                var v183 string = (("hello" + ("a" + ("" + "foo bar"))) + (("" + "\n\t") + "hello"))
                _ = v183
            }
        case "\n\t":
            var v184 int64 = 1
            _ = v184
        case "foo bar":
            switch ("foo bar" + "hello") {
            case "\n\t":
                if (126 > 2) {
                    v3 = (!v3)
                }
            case "a":
                if ("\n\t" <= "\n\t") {
                    v185 := []int{-551, 521, -868, -84}
                    _ = v185
                }
            case "hello":
                v186 := GenS19{F0: 0}
                _ = v186
            }
        default:
            var v187 uint16 = 719
            _ = v187
        }
        _ = v182
    }
    switch 1 {
    case 967:
        if (-944 >= 745) {
            for fi188 := 0; fi188 < 10; fi188++ {
                for fi189 := 0; fi189 < 4; fi189++ {
                    switch int(v1) {
                    case 895:
                        v1 = v1
                    case 9223372036854775807:
                        fmt.Println(fi188)
                    case 0:
                        var v190 string = ("a" + "a")
                        _ = v190
                    }
                    switch "\n\t" {
                    case "foo bar":
                        v3 = ((0 != 65535) || (1 > 18446744073709551615))
                    case "\n\t":
                        v4["foo bar"] = ("a" + ("" + "\n\t"))
                    }
                    if v3 {
                        switch "\n\t" {
                        case "hello":
                            if (9223372036854775807 >= -9223372036854775808) {
                                v1 += v1
                                if (((65535 <= 2) || ((v3 && (!v3)) && false)) || (!(float32((0.1 / 1e10)) < 0.1))) {
                                    var v191 int64 = 9223372036854775807
                                    v192 := GenS3{F0: (v191 / ((-598) | 1))}
                                    if true {
                                        var v193 float64 = 0.1
                                        var v194 float64 = 0.0
                                        _ = v193
                                        _ = v194
                                    }
                                    _ = v191
                                    _ = v192
                                } else {
                                    v3 = v3
                                }
                            } else {
                                _ = len(v4)
                                var v195 int8 = -70
                                v195 = (v195 ^ (v195 % ((v195) | 1)))
                                _ = v195
                            }
                        }
                    } else {
                        v3 = true
                        var v196 float32 = 1.5
                        var v197 uint32 = 65535
                        _ = v196
                        _ = v197
                    }
                    _ = fi189
                }
                var v198 uint64 = 42
                _ = fi188
                _ = v198
            }
            if v3 {
                var v199 uint = v1
                fmt.Println(v199, v199, v199)
                var v200 int = -1
                _ = v199
                _ = v200
            }
            if (0 > 65535) {
                v201 := GenS18{F0: 715, F1: 1}
                switch 42 {
                case -1:
                    fmt.Println(v3, v3)
                default:
                    for fi202 := 0; fi202 < 6; fi202++ {
                        v203 := GenS20{F0: "a"}
                        for fi204 := 0; fi204 < 4; fi204++ {
                            delete(v4, "foo bar")
                            v3 = v3
                            var v205 int64 = 208
                            var v206 int32 = 2147483647
                            fmt.Println(v1, v1, fi204)
                            var v207 string = ""
                            var v208 int = 9223372036854775807
                            _ = fi204
                            _ = v205
                            _ = v206
                            _ = v207
                            _ = v208
                        }
                        _ = fi202
                        _ = v203
                    }
                }
                var v209 int16 = -32768
                _ = v201
                _ = v209
            }
        } else {
            if v3 {
                var v210 float32 = (3.14159 - 1.5)
                _ = v210
            }
            switch -1 {
            case 1:
                for fi211 := 0; fi211 < 17; fi211++ {
                    for fi212 := 0; fi212 < 4; fi212++ {
                        var v213 uint16 = 0
                        if v3 {
                            v214 := GenS20{F0: ""}
                            if (-373 >= -573) {
                                if false {
                                    v213 = v213
                                    v4["foo bar"] = (("" + "\n\t") + "")
                                    var v215 int = 316
                                    _ = v215
                                } else {
                                    var v216 float32 = (float32((((-2.25 + 0.0) + float64((((1e10 - float32((0.0 - (1.5 + float64((-2.25 / 0.1)))))) * 0.1) - ((1e10 / 1e10) / 0.1)))) - ((1e10 * (0.0 / 1e10)) * 0.0))) / 1.5)
                                    _ = v216
                                }
                            }
                            _ = v214
                        } else {
                            var v217 uint8 = 42
                            var v218 float64 = 1.5
                            v218 = 1.5
                            _ = v217
                            _ = v218
                        }
                        _ = fi212
                        _ = v213
                    }
                    _ = fi211
                }
            case 127:
                var v219 uint8 = 0
                _ = v219
            case -9223372036854775808:
                var v220 uint = (v1 | 255)
                _ = v220
            default:
                var v221 uint32 = 808
                _ = v221
            }
        }
    case -505:
        v222 := GenS13{F0: 42, F1: "hello"}
        _ = v222
    case 1:
        switch 4294967295 {
        case 0:
            if (!(728 != 1)) {
                fmt.Println(v1, v3)
                v223 := []int{-59, 168}
                _ = v223
            }
        case -2147483648:
            switch ("" + "hello") {
            case "\n\t":
                for fi224 := 0; fi224 < 2; fi224++ {
                    v3 = ((v1 & 2) > (v1 - 443))
                    v1 = (^(v1))
                    _ = fi224
                }
            case "":
                v225 := []int{-9223372036854775808, 24, 127}
                _ = v225
            case "hello":
                switch -544 {
                case 127:
                    var v226 uint32 = 2147483648
                    _ = v226
                case 32768:
                    var v227 float32 = float32((((0.1 - 1.5) - float64(float32(-2.25))) / 1.5))
                    _ = v227
                case -2147483648:
                    v228 := v2[-829]
                    _ = v228
                default:
                    v3 = (!(!v3))
                }
            }
        case 32767:
            v1 = v1
        default:
            var v229 int = -74
            _ = v229
        }
    default:
        var v230 int8 = -113
        _ = v230
    }
    v231 := GenS21{F0: 325}
    fmt.Println(v3, v3)
    var v232 int32 = 0
    var v233 uint16 = 202
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v232:", v232)
    fmt.Println("v233:", v233)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v231
    _ = v232
    _ = v233
    _ = v3
    _ = v4
}
