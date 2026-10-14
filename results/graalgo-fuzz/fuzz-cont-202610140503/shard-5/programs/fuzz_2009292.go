package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint64; F1 int8 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 bool; F1 string }
type GenS5 struct { F0 int16; F1 int64; F2 string }
type GenS6 struct { F0 int8; F1 uint8 }
type GenS7 struct { F0 int8; F1 bool }
type GenS8 struct { F0 uint; F1 int32 }
type GenS9 struct { F0 uint32; F1 int16 }
type GenS10 struct { F0 uint16 }
type GenS11 struct { F0 int64 }
type GenS12 struct { F0 int8; F1 int; F2 bool }
type GenS13 struct { F0 int8 }
type GenS14 struct { F0 int8; F1 int32 }

func main() {
    var v1 int = 745
    var v2 uint8 = 193
    var v3 int8 = -128
    var v4 string = (("hello" + "\n\t") + "a")
    v5 := make(map[uint32]bool)
    var v6 int8 = (^(v3))
    switch v4 {
    case "hello":
        fmt.Println(v1)
    case "":
        for fi7 := 0; fi7 < 14; fi7++ {
            var v8 int = (-(fi7))
            _ = fi7
            _ = v8
        }
    case "foo bar":
        var v9 int16 = 1000
        _ = v9
    }
    v10 := GenS1{F0: (v3 << ((2) & 63))}
    switch v4 {
    case "\n\t":
        var v11 float64 = 0.1
        _ = v11
    }
    switch (("hello" + v4) + "\n\t") {
    case "foo bar":
        if (32767 == 1000) {
            var v12 float32 = 0.0
            v3 += -4
            _ = v12
        }
    case "hello":
        v6 = v3
    case "":
        for fi13 := 0; fi13 < 8; fi13++ {
            v4 = "foo bar"
            _ = fi13
        }
    default:
        if (!(1 < 42)) {
            var v14 int = 2147483647
            switch v14 {
            case -2147483648:
                if true {
                    v15 := []int{(v14 % (((v14 % ((-110) | 1))) | 1)), -184}
                    _ = v15
                } else {
                    v16 := []int{(v14 * (v14 - v14)), v1}
                    for fi17 := 0; fi17 < 1; fi17++ {
                        v18 := GenS1{F0: v6}
                        switch ((v4 + (v4 + v4)) + "a") {
                        case "foo bar":
                            fmt.Println(v3)
                        case "":
                            switch v1 {
                            case 47:
                                if len(v16) > 0 {
                                    v16[0] = (fi17 - fi17)
                                }
                            case -2147483648:
                                if (false || true) {
                                    var v19 int32 = -1
                                    _ = v19
                                }
                            }
                        case "hello":
                            var v20 uint = 1
                            _ = v20
                        default:
                            switch v4 {
                            case "foo bar":
                                if len(v16) >= 2 {
                                    v16 = v16[1:len(v16)]
                                }
                            default:
                                var v21 uint16 = 0
                                _ = v21
                            }
                        }
                        var v22 uint32 = 421
                        _ = fi17
                        _ = v18
                        _ = v22
                    }
                    _ = v16
                }
            default:
                v23 := GenS2{F0: 18446744073709551615, F1: (v6 % ((v6) | 1))}
                _ = v23
            }
            fmt.Println(v14, v4)
            _ = v14
        } else {
            v24 := GenS2{F0: 2, F1: (v3 & 67)}
            var v25 int = (v1 * (v1 / (((v1 & (v1 / ((v1) | 1)))) | 1)))
            var v26 string = v4
            _ = v24
            _ = v25
            _ = v26
        }
    }
    v27 := GenS3{F0: -81}
    v3 += 0
    var v28 int8 = -128
    if (false || false) {
        fmt.Println(v3)
    }
    v29 := GenS4{F0: (true || (!true)), F1: v4}
    fmt.Println(v6, v6)
    var v30 uint8 = v2
    v31 := []int{}
    v32 := GenS5{F0: -242, F1: -575, F2: "foo bar"}
    switch (^(v1)) {
    case -350:
        if false {
            if (42 >= 65535) {
                for fi33 := 0; fi33 < 1; fi33++ {
                    v1 += 4294967295
                    for fi34 := 0; fi34 < 17; fi34++ {
                        if false {
                            if false {
                                v35 := GenS3{F0: 42}
                                var v36 float64 = ((1e10 / 0.1) * float64(1.5))
                                var v37 uint16 = 42
                                _ = v35
                                _ = v36
                                _ = v37
                            }
                        } else {
                            var v38 string = "hello"
                            var v39 uint32 = 1
                            _ = v38
                            _ = v39
                        }
                        if ((1000 > -668) || (!false)) {
                            switch v4 {
                            case "a":
                                if false {
                                    v2 = 1
                                    var v40 bool = false
                                    v2 -= (^(v2))
                                    _ = v40
                                }
                                v6 = (^(v3))
                                var v41 int8 = -104
                                _ = v41
                            case "foo bar":
                                v1 += (fi33 & (fi34 ^ v1))
                            }
                            var v42 bool = false
                            if v42 {
                                if v42 {
                                    var v43 int = (fi33 * fi34)
                                    var v44 uint16 = 2
                                    _ = v43
                                    _ = v44
                                }
                                if ((float64(0.1) + 3.14159) == 0.0) {
                                    v45 := GenS6{F0: (v3 + v6), F1: (v2 << ((255) & 63))}
                                    switch (-(fi34)) {
                                    case 128:
                                        var v46 float32 = 1e10
                                        _ = v46
                                    default:
                                        v2 -= 127
                                    }
                                    v42 = (false && (4294967295 < 529))
                                    _ = v45
                                } else {
                                    var v47 int16 = 1000
                                    _ = v47
                                }
                                if v42 {
                                    var v48 int64 = 1
                                    var v49 int32 = -1
                                    _ = v48
                                    _ = v49
                                }
                            }
                            _ = v42
                        } else {
                            v30 = v2
                            var v50 int16 = -1
                            _ = len(v31)
                            if ((0.1 * 0.0) <= (1e10 - 3.14159)) {
                                v5[1] = false
                                fmt.Println(v30, v4)
                                fmt.Println(v4, v4)
                            }
                            _ = v50
                        }
                        var v51 uint64 = 56
                        _ = fi34
                        _ = v51
                    }
                    v2 += 1
                    _ = fi33
                }
                var v52 int32 = 1
                _ = v52
            } else {
                var v53 uint16 = 32767
                _ = v53
            }
        } else {
            for fi54 := 0; fi54 < 16; fi54++ {
                fmt.Println(v3)
                _ = fi54
            }
            switch (v1 + v1) {
            case 887:
                var v55 uint64 = 946
                _ = v55
            case -2147483648:
                switch v4 {
                case "hello":
                    v56 := GenS3{F0: -2147483648}
                    _ = v56
                default:
                    for fi57 := 0; fi57 < 18; fi57++ {
                        v30 += (v30 ^ (v30 + v2))
                        _ = fi57
                    }
                }
            case -219:
                for fi58 := 0; fi58 < 20; fi58++ {
                    if false {
                        var v59 float64 = 0.1
                        var v60 int16 = -657
                        _ = v59
                        _ = v60
                    } else {
                        v3 -= (v3 ^ 42)
                        if false {
                            v1 = fi58
                            v61 := GenS1{F0: v6}
                            var v62 bool = true
                            _ = v61
                            _ = v62
                        } else {
                            if (true && false) {
                                v4 = "\n\t"
                            } else {
                                var v63 uint8 = (v30 % (((^(v30))) | 1))
                                v30 -= 128
                                if (v2 <= (v30 + (v30 ^ (v2 << ((2) & 63))))) {
                                    fmt.Println(v6, v1)
                                }
                                _ = v63
                            }
                        }
                        v1 -= (fi58 << ((4294967295) & 63))
                    }
                    switch fi58 {
                    case 873:
                        v64 := GenS5{F0: 32767, F1: -1, F2: "foo bar"}
                        _ = v64
                    case 4294967295:
                        for fi65 := 0; fi65 < 6; fi65++ {
                            var v66 uint = 280
                            v30 = 255
                            _ = fi65
                            _ = v66
                        }
                    default:
                        v67 := GenS7{F0: v6, F1: (false || true)}
                        _ = v67
                    }
                    _ = fi58
                }
            }
        }
    default:
        for fi68 := 0; fi68 < 8; fi68++ {
            delete(v5, 2)
            fmt.Println(v3, v28)
            if false {
                if len(v31) > 0 {
                    v31[0] = fi68
                }
                var v69 uint8 = (v2 & 52)
                _ = v69
            }
            _ = fi68
        }
    }
    for fi70 := 0; fi70 < 13; fi70++ {
        var v71 bool = false
        switch v4 {
        case "":
            var v72 float32 = 1e10
            _ = v72
        case "hello":
            switch 128 {
            case 32768:
                if v71 {
                    if false {
                        for fi73 := 0; fi73 < 14; fi73++ {
                            v74 := GenS8{F0: 42, F1: 0}
                            var v75 uint64 = 2
                            _ = fi73
                            _ = v74
                            _ = v75
                        }
                        switch 1 {
                        case 78:
                            v3 = (v3 * (v3 % ((v6) | 1)))
                        case 0:
                            if (!(1 >= 242)) {
                                v5[703] = v71
                                fmt.Println(v28, v4, v3)
                                if (128 <= v1) {
                                    v28 = v6
                                    for fi76 := 0; fi76 < 6; fi76++ {
                                        if (1.5 >= 1e10) {
                                            var v77 int8 = 111
                                            _ = len(v5)
                                            _ = v77
                                        } else {
                                            if false {
                                                v28 -= 1
                                                v1 = 32767
                                                switch (fi76 % ((32767) | 1)) {
                                                case -1:
                                                    if (((-500 >= 1000) && ((!v71) || v71)) || (!(((65535 != uint32(v30)) && v71) && (42 >= 914)))) {
                                                        var v78 bool = (722 < 32767)
                                                        _ = v78
                                                    } else {
                                                        v79 := GenS6{F0: v3, F1: v2}
                                                        _ = v79
                                                    }
                                                default:
                                                    switch ((v4 + v4) + (v4 + v4)) {
                                                    case "":
                                                        var v80 uint16 = 42
                                                        _ = v80
                                                    case "\n\t":
                                                        _ = len(v5)
                                                    }
                                                }
                                            }
                                            if ((v2 + (v30 >> ((65535) & 63))) < (v2 >> ((4294967295) & 63))) {
                                                if (((("" + (v4 + v4)) + (v4 + "hello")) > v4) && v71) {
                                                    v6 += 127
                                                    switch ("foo bar" + v4) {
                                                    case "foo bar":
                                                        v81 := GenS3{F0: -2147483648}
                                                        _ = v81
                                                    default:
                                                        if ((-1 == 818) || (v71 && v71)) {
                                                            v6 = (v6 - (v28 >> ((18446744073709551615) & 63)))
                                                            if (1 == 0) {
                                                                switch "\n\t" {
                                                                case "\n\t":
                                                                    v1 = (fi70 % ((853) | 1))
                                                                case "a":
                                                                    v28 = (^(v3))
                                                                case "":
                                                                    var v82 uint16 = 32767
                                                                    _ = v82
                                                                }
                                                                if len(v31) > 0 {
                                                                    v31[0] = 128
                                                                }
                                                            } else {
                                                                fmt.Println(v2)
                                                                var v83 float64 = 1e10
                                                                fmt.Println(v71)
                                                                _ = v83
                                                            }
                                                            _ = len(v5)
                                                        } else {
                                                            v6 += v6
                                                            v3 = v3
                                                            switch ("\n\t" + ((v4 + "foo bar") + (v4 + v4))) {
                                                            case "foo bar":
                                                                var v84 int64 = 700
                                                                _ = v84
                                                            case "hello":
                                                                var v85 uint = 65535
                                                                _ = v85
                                                            case "a":
                                                                v1 -= (^(v1))
                                                            }
                                                            v86 := GenS1{F0: int8(v2)}
                                                            if false {
                                                                if len(v31) >= 2 {
                                                                    v31 = v31[1:len(v31)]
                                                                }
                                                                fmt.Println(v30)
                                                                if (v71 || (748 != 0)) {
                                                                    if v71 {
                                                                        fmt.Println(v28)
                                                                    }
                                                                    var v87 int32 = 2147483647
                                                                    var v88 uint16 = 65535
                                                                    var v89 float64 = ((((-2.25 + 1e10) - 3.14159) / ((float64(3.14159) * (float64(1e10) / 1e10)) * ((1.5 + 3.14159) * 1e10))) * ((((3.14159 * 1e10) + ((1.5 * float64(0.1)) * float64(0.0))) / 3.14159) / (0.0 + (3.14159 - 1e10))))
                                                                    _ = v87
                                                                    _ = v88
                                                                    _ = v89
                                                                }
                                                                _ = len(v31)
                                                            } else {
                                                                v3 += (^(v6))
                                                                v90 := GenS9{F0: 323, F1: -32768}
                                                                _ = v90
                                                            }
                                                            _ = v86
                                                        }
                                                    }
                                                    switch ("\n\t" + (v4 + (v4 + v4))) {
                                                    case "\n\t":
                                                        switch ("a" + v4) {
                                                        case "foo bar":
                                                            v2 += (^(v30))
                                                        case "":
                                                            if (v71 || v71) {
                                                                var v91 string = v4
                                                                var v92 string = (v91 + v91)
                                                                _ = v91
                                                                _ = v92
                                                            } else {
                                                                if (42 == 42) {
                                                                    if len(v31) > 0 {
                                                                        v31[0] = (^(v1))
                                                                    }
                                                                } else {
                                                                    var v93 int16 = 42
                                                                    if v71 {
                                                                        var v94 string = ("\n\t" + "\n\t")
                                                                        _ = v94
                                                                    }
                                                                    _ = v93
                                                                }
                                                            }
                                                            if len(v31) >= 2 {
                                                                v31 = v31[1:len(v31)]
                                                            }
                                                            var v95 int16 = -32768
                                                            _ = v95
                                                        case "\n\t":
                                                            var v96 string = "\n\t"
                                                            _ = v96
                                                        default:
                                                            v97 := GenS10{F0: 32767}
                                                            _ = v97
                                                        }
                                                    case "":
                                                        if (v71 || true) {
                                                            v98 := GenS1{F0: (^(v3))}
                                                            _ = v98
                                                        } else {
                                                            fmt.Println(fi70)
                                                        }
                                                    }
                                                }
                                                if v71 {
                                                    v2 = v2
                                                    var v99 int = fi70
                                                    _ = v99
                                                }
                                                if v71 {
                                                    fmt.Println(v4, v3)
                                                }
                                                var v100 string = "foo bar"
                                                _ = v100
                                            } else {
                                                if (0 == 0) {
                                                    if len(v31) >= 2 {
                                                        v31 = v31[1:len(v31)]
                                                    }
                                                }
                                                v2 = (^(v2))
                                                v101 := GenS11{F0: 1}
                                                var v102 int = (fi76 & v1)
                                                _ = v101
                                                _ = v102
                                            }
                                            if v71 {
                                                switch ((("foo bar" + v4) + v4) + v4) {
                                                case "":
                                                    v2 -= 2
                                                case "foo bar":
                                                    delete(v5, 892)
                                                }
                                                v4 = v4
                                                fmt.Println(v6, v4, v3)
                                            }
                                        }
                                        _ = fi76
                                    }
                                    var v103 int = fi70
                                    _ = v103
                                }
                            }
                        default:
                            if true {
                                v6 = -57
                                var v104 int64 = 9223372036854775807
                                var v105 uint8 = v2
                                _ = v104
                                _ = v105
                            } else {
                                v106 := GenS2{F0: 388, F1: (v28 * 0)}
                                for fi107 := 0; fi107 < 13; fi107++ {
                                    var v108 string = (v4 + ("" + (v4 + v4)))
                                    switch v4 {
                                    case "":
                                        if (985 == 774) {
                                            if false {
                                                v109 := GenS12{F0: -32, F1: 824, F2: (18446744073709551615 > 4294967295)}
                                                if (589 == -1) {
                                                    fmt.Println(v28, v6, fi107)
                                                }
                                                _ = v109
                                            }
                                            v108 = ("\n\t" + "")
                                        }
                                    case "a":
                                        fmt.Println(v71)
                                    default:
                                        v28 -= v6
                                    }
                                    if (!v71) {
                                        _ = len(v31)
                                        if v71 {
                                            v3 += v6
                                        }
                                    } else {
                                        v3 = (^(v3))
                                    }
                                    v6 = 100
                                    _ = fi107
                                    _ = v108
                                }
                                if false {
                                    if len(v31) > 0 {
                                        v31[0] = (v1 / (((fi70 + v1)) | 1))
                                    }
                                    if (-2.25 < (float32((float64((0.0 - -2.25)) / (1e10 / 1.0))) - 0.1)) {
                                        for fi110 := 0; fi110 < 10; fi110++ {
                                            v71 = (int(v2) == int(v3))
                                            switch v1 {
                                            case 2147483647:
                                                var v111 string = v4
                                                _ = v111
                                            case -254:
                                                if (!((((255 >= 824) || (v71 || v71)) && (!((v28 ^ (v3 & v3)) >= (^(v3))))) || (v1 > (fi70 + (v1 * (v1 ^ fi70)))))) {
                                                    if v71 {
                                                        v5[uint32(v30)] = ((0 < 830) && ((32767 != 52) || (325 == 18446744073709551615)))
                                                        if (false && v71) {
                                                            var v112 bool = true
                                                            if (1 <= 255) {
                                                                if ((fi70 % ((v1) | 1)) < (fi110 * 897)) {
                                                                    v2 = v2
                                                                }
                                                                if len(v31) >= 2 {
                                                                    v31 = v31[1:len(v31)]
                                                                }
                                                                var v113 uint16 = 42
                                                                _ = v113
                                                            }
                                                            _ = v112
                                                        }
                                                    }
                                                    var v114 int16 = -1
                                                    _ = v114
                                                } else {
                                                    switch 1 {
                                                    case -9223372036854775808:
                                                        var v115 uint32 = 300
                                                        _ = v115
                                                    case -2147483648:
                                                        fmt.Println(fi110, fi110, v1)
                                                    case 1:
                                                        if false {
                                                            v1 += -1
                                                            v3 -= -128
                                                            var v116 string = "foo bar"
                                                            v30 -= v2
                                                            v117 := GenS10{F0: 1}
                                                            v6 = -1
                                                            v71 = (0 == 0)
                                                            var v118 uint16 = 65535
                                                            _ = v116
                                                            _ = v117
                                                            _ = v118
                                                        }
                                                    default:
                                                        switch v4 {
                                                        case "hello":
                                                            if (-541 != (fi70 | v1)) {
                                                                if len(v31) >= 2 {
                                                                    v31 = v31[1:len(v31)]
                                                                }
                                                            }
                                                        case "\n\t":
                                                            if (2147483647 != 2) {
                                                                switch (v1 >> ((797) & 63)) {
                                                                case 42:
                                                                    fmt.Println(v1)
                                                                case 1:
                                                                    if v71 {
                                                                        if false {
                                                                            if ((1 >= v6) || ((65535 <= 0) && v71)) {
                                                                                if (-272 != 9223372036854775807) {
                                                                                    switch (fi110 << ((18446744073709551615) & 63)) {
                                                                                    case 4294967295:
                                                                                        v119 := GenS10{F0: 65535}
                                                                                        _ = v119
                                                                                    case 249:
                                                                                        var v120 string = "hello"
                                                                                        if false {
                                                                                            v3 = (v3 / ((v3) | 1))
                                                                                        }
                                                                                        _ = v120
                                                                                    case 2147483647:
                                                                                        switch (fi110 >> ((42) & 63)) {
                                                                                        case -696:
                                                                                            switch v4 {
                                                                                            case "":
                                                                                                v4 = (v4 + v4)
                                                                                            }
                                                                                        }
                                                                                        v30 += v2
                                                                                        _ = len(v5)
                                                                                    }
                                                                                }
                                                                                if false {
                                                                                    v121 := GenS12{F0: v3, F1: fi110, F2: false}
                                                                                    switch 127 {
                                                                                    case -186:
                                                                                        switch (v4 + v4) {
                                                                                        case "":
                                                                                            v1 = 4294967295
                                                                                        }
                                                                                    case 42:
                                                                                        switch v1 {
                                                                                        case 707:
                                                                                            v122 := v5[993]
                                                                                            _ = v122
                                                                                        }
                                                                                        if v71 {
                                                                                            v1 += fi110
                                                                                            v30 = (v2 | v30)
                                                                                            if (42 > 0) {
                                                                                                var v123 uint64 = 2
                                                                                                _ = v123
                                                                                            } else {
                                                                                                v2 = v30
                                                                                            }
                                                                                        } else {
                                                                                            if (1 == 0) {
                                                                                                v4 = ((((v4 + v4) + ("a" + v4)) + v4) + ((v4 + v4) + v4))
                                                                                                if (false && v71) {
                                                                                                    v2 += (v2 - (^(v30)))
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                    default:
                                                                                        var v124 uint64 = 2
                                                                                        _ = v124
                                                                                    }
                                                                                    v6 = v28
                                                                                    _ = v121
                                                                                } else {
                                                                                    v4 = v4
                                                                                }
                                                                            }
                                                                            v3 += v28
                                                                            switch (-(fi70)) {
                                                                            case -1:
                                                                                v125 := GenS7{F0: v28, F1: (!v71)}
                                                                                _ = v125
                                                                            default:
                                                                                if (v71 || (1 >= 42)) {
                                                                                    var v126 float32 = (3.14159 / (1.5 / (-2.25 - 0.1)))
                                                                                    if (uint64(v30) <= 187) {
                                                                                        v127 := GenS2{F0: 1, F1: v28}
                                                                                        var v128 bool = (v71 || false)
                                                                                        switch v1 {
                                                                                        case -2147483648:
                                                                                            v4 = v4
                                                                                        case 937:
                                                                                            switch fi110 {
                                                                                            case 86:
                                                                                                switch (^(v1)) {
                                                                                                case -902:
                                                                                                    var v129 uint32 = 2147483647
                                                                                                    _ = v129
                                                                                                case -154:
                                                                                                    var v130 int64 = -147
                                                                                                    _ = v130
                                                                                                case -9223372036854775808:
                                                                                                    var v131 int16 = -1
                                                                                                    _ = v131
                                                                                                default:
                                                                                                    var v132 int = (fi110 & fi70)
                                                                                                    _ = v132
                                                                                                }
                                                                                                v71 = true
                                                                                            case 42:
                                                                                                if len(v31) >= 2 {
                                                                                                    v31 = v31[1:len(v31)]
                                                                                                }
                                                                                            default:
                                                                                                v30 = v30
                                                                                            }
                                                                                        case 18:
                                                                                            var v133 string = v4
                                                                                            _ = v133
                                                                                        default:
                                                                                            v3 += v3
                                                                                        }
                                                                                        _ = v127
                                                                                        _ = v128
                                                                                    } else {
                                                                                        v134 := v5[0]
                                                                                        var v135 int = v1
                                                                                        _ = v134
                                                                                        _ = v135
                                                                                    }
                                                                                    v71 = (65535 != 1)
                                                                                    _ = v126
                                                                                }
                                                                            }
                                                                            v136 := GenS10{F0: 721}
                                                                            _ = v136
                                                                        } else {
                                                                            v137 := GenS8{F0: 702, F1: -1}
                                                                            switch 0 {
                                                                            case 106:
                                                                                v30 = (v30 + 2)
                                                                            }
                                                                            switch fi70 {
                                                                            case -640:
                                                                                if v71 {
                                                                                    _ = len(v5)
                                                                                    _ = len(v31)
                                                                                    v5[4294967295] = v71
                                                                                } else {
                                                                                    v138 := GenS13{F0: v6}
                                                                                    _ = v138
                                                                                }
                                                                            default:
                                                                                var v139 uint64 = 4294967295
                                                                                _ = v139
                                                                            }
                                                                            _ = v137
                                                                        }
                                                                        v3 = 1
                                                                        var v140 int8 = (v3 + int8(v1))
                                                                        _ = v140
                                                                    } else {
                                                                        delete(v5, 2147483648)
                                                                        var v141 bool = v71
                                                                        var v142 int32 = -2147483648
                                                                        _ = v141
                                                                        _ = v142
                                                                    }
                                                                    v143 := GenS8{F0: 107, F1: 1}
                                                                    _ = v143
                                                                }
                                                                v30 = v30
                                                                var v144 uint = 0
                                                                v3 = 100
                                                                if (!(v71 || v71)) {
                                                                    fmt.Println(v2, v71, fi70)
                                                                    v145 := GenS11{F0: 42}
                                                                    if (-32768 == 457) {
                                                                        var v146 uint64 = 4294967295
                                                                        if v71 {
                                                                            fmt.Println(v3, v6, v71)
                                                                        } else {
                                                                            var v147 int8 = (v28 * 112)
                                                                            v148 := GenS4{F0: v71, F1: v4}
                                                                            v149 := GenS4{F0: false, F1: v4}
                                                                            v1 = -1
                                                                            fmt.Println(fi110, v144)
                                                                            _ = v147
                                                                            _ = v148
                                                                            _ = v149
                                                                        }
                                                                        _ = v146
                                                                    }
                                                                    if (-932 > 2147483647) {
                                                                        fmt.Println(v1)
                                                                        var v150 int64 = -100
                                                                        v6 -= int8(v150)
                                                                        _ = v150
                                                                    } else {
                                                                        v151 := GenS11{F0: 594}
                                                                        fmt.Println(v1)
                                                                        _ = v151
                                                                    }
                                                                    switch (v1 / ((fi70) | 1)) {
                                                                    case 32767:
                                                                        v5[42] = ((!(593 <= 1)) && ((v2 << (((v144 * (^(v144)))) & 63)) != (^(v2))))
                                                                    case -1:
                                                                        if ((-2.25 + (((0.0 + 1e10) - (-2.25 * (0.0 - (3.14159 - 0.1)))) * ((-2.25 - 0.1) * -2.25))) > (0.0 - 0.1)) {
                                                                            var v152 int32 = -2147483648
                                                                            _ = v152
                                                                        } else {
                                                                            if v71 {
                                                                                if v71 {
                                                                                    fmt.Println(fi110)
                                                                                    var v153 int32 = -863
                                                                                    _ = v153
                                                                                }
                                                                            } else {
                                                                                fmt.Println(fi110, v4)
                                                                                var v154 int32 = 1
                                                                                if v71 {
                                                                                    v30 = (v2 + v2)
                                                                                    v5[775] = v71
                                                                                }
                                                                                if true {
                                                                                    var v155 int16 = 1000
                                                                                    _ = len(v31)
                                                                                    fmt.Println(v155)
                                                                                    _ = v155
                                                                                } else {
                                                                                    switch v4 {
                                                                                    case "hello":
                                                                                        if false {
                                                                                            v4 = v4
                                                                                            var v156 float64 = 0.1
                                                                                            _ = v156
                                                                                        }
                                                                                    }
                                                                                    v157 := GenS7{F0: (v28 / ((v28) | 1)), F1: (1e10 != (0.1 * -2.25))}
                                                                                    v158 := GenS14{F0: v3, F1: (v154 * v154)}
                                                                                    _ = v157
                                                                                    _ = v158
                                                                                }
                                                                                _ = v154
                                                                            }
                                                                            v6 = (-(v3))
                                                                            fmt.Println(v2)
                                                                        }
                                                                    case 997:
                                                                        var v159 bool = v71
                                                                        _ = v159
                                                                    }
                                                                    _ = v145
                                                                } else {
                                                                    switch -728 {
                                                                    case 790:
                                                                        fmt.Println(v1)
                                                                    case 0:
                                                                        _ = len(v5)
                                                                    case -68:
                                                                        _ = len(v5)
                                                                    }
                                                                    v6 -= 48
                                                                }
                                                                _ = v144
                                                            } else {
                                                                v30 = (v30 & v2)
                                                                _ = len(v5)
                                                            }
                                                        case "foo bar":
                                                            fmt.Println(v30, v30)
                                                        }
                                                        fmt.Println(fi110, v2)
                                                    }
                                                }
                                            }
                                            _ = fi110
                                        }
                                        if (!(v28 != (^(v3)))) {
                                            var v160 uint = 0
                                            v3 -= v28
                                            v161 := GenS4{F0: ((((float32(float64(0.0)) + 1e10) - 3.14159) != 0.1) && (!true)), F1: "foo bar"}
                                            _ = v160
                                            _ = v161
                                        } else {
                                            var v162 int8 = v3
                                            var v163 int16 = -400
                                            var v164 int8 = (v162 << ((18446744073709551615) & 63))
                                            _ = v162
                                            _ = v163
                                            _ = v164
                                        }
                                        var v165 uint32 = 279
                                        _ = v165
                                    }
                                } else {
                                    v71 = v71
                                    fmt.Println(v71, v3)
                                }
                                _ = v106
                            }
                        }
                        var v166 int64 = 9223372036854775807
                        _ = v166
                    }
                    var v167 int16 = 0
                    _ = v167
                }
            }
        default:
            _ = len(v5)
        }
        _ = fi70
        _ = v71
    }
    if true {
        var v168 int64 = -1
        for fi169 := 0; fi169 < 7; fi169++ {
            for fi170 := 0; fi170 < 8; fi170++ {
                if false {
                    v28 = (v3 % (((^(v3))) | 1))
                } else {
                    if len(v31) >= 2 {
                        v31 = v31[1:len(v31)]
                    }
                    var v171 string = ("foo bar" + ((v4 + "foo bar") + ""))
                    fmt.Println(v30, v30, v30)
                    var v172 int = fi170
                    var v173 uint8 = 42
                    if len(v31) >= 2 {
                        v31 = v31[1:len(v31)]
                    }
                    if false {
                        v174 := v5[0]
                        var v175 uint16 = 2
                        _ = v174
                        _ = v175
                    }
                    fmt.Println(v6)
                    _ = v171
                    _ = v172
                    _ = v173
                }
                var v176 uint8 = 2
                _ = fi170
                _ = v176
            }
            var v177 int = fi169
            _ = fi169
            _ = v177
        }
        _ = v168
    }
    v178 := GenS4{F0: true, F1: (v4 + v4)}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v178
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
