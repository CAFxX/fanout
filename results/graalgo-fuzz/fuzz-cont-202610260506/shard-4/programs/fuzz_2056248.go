package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64 }
type GenS2 struct { F0 uint64; F1 string; F2 uint16 }
type GenS3 struct { F0 int64; F1 int; F2 bool }
type GenS4 struct { F0 int16; F1 uint16 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 string; F1 uint; F2 uint32 }
type GenS8 struct { F0 int; F1 int; F2 uint }
type GenS9 struct { F0 int }
type GenS10 struct { F0 uint8; F1 bool }
type GenS11 struct { F0 int16; F1 bool; F2 bool }
type GenS12 struct { F0 int64; F1 bool }
type GenS13 struct { F0 bool; F1 int16 }
type GenS14 struct { F0 uint64; F1 uint16 }
type GenS15 struct { F0 string; F1 uint64 }
type GenS16 struct { F0 uint32 }

func main() {
    v1 := make(map[string]bool)
    v2 := []int{1, 602}
    var v3 uint64 = 42
    var v4 int32 = -2147483648
    for fi5 := 0; fi5 < 7; fi5++ {
        var v6 uint64 = 18446744073709551615
        v3 = v6
        if (!true) {
            v3 = 2
            var v7 uint16 = 78
            switch -749 {
            case 831:
                _ = len(v2)
            case -673:
                if (!true) {
                    v8 := GenS1{F0: 214, F1: (^(v6))}
                    _ = v8
                } else {
                    if false {
                        for fi9 := 0; fi9 < 18; fi9++ {
                            v3 += v6
                            _ = fi9
                        }
                    } else {
                        var v10 uint8 = 127
                        _ = v10
                    }
                    for fi11 := 0; fi11 < 7; fi11++ {
                        var v12 string = (("hello" + ("\n\t" + ("foo bar" + ((("\n\t" + "a") + "a") + "foo bar")))) + "")
                        v12 = "foo bar"
                        _ = fi11
                        _ = v12
                    }
                }
            }
            _ = v7
        }
        _ = fi5
        _ = v6
    }
    switch -9223372036854775808 {
    case 127:
        switch 9223372036854775807 {
        case 4294967295:
            if false {
                if ((2147483647 != 688) || true) {
                    switch ("foo bar" + "hello") {
                    case "hello":
                        v4 = -1
                    case "foo bar":
                        var v13 uint64 = (^(v3))
                        _ = v13
                    case "\n\t":
                        var v14 uint16 = 712
                        _ = v14
                    }
                    switch -762 {
                    case -69:
                        var v15 uint = 18446744073709551615
                        _ = v15
                    case 127:
                        v4 = v4
                    case 2147483647:
                        fmt.Println(v4, v4)
                    default:
                        if (false && ((float32((((3.14159 * (((-2.25 / 1.0) / float64((0.1 + (3.14159 / (0.1 * (0.1 - (1e10 + -2.25))))))) - ((0.1 * -2.25) / (0.1 + 3.14159)))) * (0.0 - (0.1 + (0.0 * (3.14159 * 0.1))))) / 1e10)) / (0.1 + (((0.1 - 3.14159) * 0.1) + (0.1 / 1e10)))) <= float32(3.14159))) {
                            if true {
                                if (v3 < v3) {
                                    var v16 int64 = 497
                                    switch (("a" + "\n\t") + "\n\t") {
                                    case "hello":
                                        v2 = append(v2, 2147483647)
                                    }
                                    _ = v16
                                } else {
                                    var v17 int = 0
                                    var v18 int = 264
                                    v3 += v3
                                    _ = v17
                                    _ = v18
                                }
                                v3 = (v3 & v3)
                            } else {
                                fmt.Println(v4, v4)
                                for fi19 := 0; fi19 < 5; fi19++ {
                                    for fi20 := 0; fi20 < 4; fi20++ {
                                        fmt.Println(fi19, v3, fi19)
                                        v3 = (v3 << ((42) & 63))
                                        _ = fi20
                                    }
                                    if false {
                                        v3 -= (v3 ^ 90)
                                    } else {
                                        v3 = v3
                                    }
                                    _ = fi19
                                }
                            }
                            fmt.Println(v3, v3)
                        } else {
                            v3 = v3
                            switch -860 {
                            case 127:
                                v4 = (^(v4))
                            case -930:
                                var v21 uint64 = v3
                                _ = v21
                            }
                        }
                    }
                } else {
                    fmt.Println(v4, v4)
                }
                if (-128 <= 68) {
                    if len(v2) > 0 {
                        v2[0] = -236
                    }
                    v2 = append(v2, -1)
                }
                v3 += v3
            } else {
                switch "a" {
                case "\n\t":
                    var v22 float32 = (float32(float64((-2.25 / 3.14159))) / float32(1.5))
                    _ = v22
                case "foo bar":
                    v4 += v4
                default:
                    var v23 float64 = 3.14159
                    _ = v23
                }
                delete(v1, (("" + "hello") + ("foo bar" + ("foo bar" + ((("\n\t" + "\n\t") + "a") + "")))))
            }
        default:
            if (((v4 >> ((255) & 63)) > v4) || (((((("\n\t" + "\n\t") + "\n\t") + "a") + ("hello" + "a")) <= ("\n\t" + ("\n\t" + ""))) || false)) {
                var v24 uint8 = 1
                v3 = v3
                _ = v24
            } else {
                v25 := GenS1{F0: 372, F1: 1}
                v2 = append(v2, 9223372036854775807)
                _ = v25
            }
        }
    case 0:
        v3 = 830
    }
    var v26 uint64 = (v3 / ((1) | 1))
    var v27 int64 = -1
    if (1 == v3) {
        switch ((("\n\t" + (("hello" + "foo bar") + "\n\t")) + "foo bar") + "hello") {
        case "hello":
            fmt.Println(v27, v26, v3)
        case "foo bar":
            v28 := GenS1{F0: 2147483648, F1: (v3 % (((v3 + 0)) | 1))}
            _ = v28
        case "a":
            var v29 int64 = -487
            _ = v29
        default:
            for fi30 := 0; fi30 < 2; fi30++ {
                for fi31 := 0; fi31 < 1; fi31++ {
                    v32 := GenS1{F0: 549, F1: 106}
                    v4 += v4
                    if (!true) {
                        v33 := GenS2{F0: (v3 - v26), F1: "\n\t", F2: 42}
                        fmt.Println(fi31, fi30)
                        var v34 uint8 = 42
                        _ = v33
                        _ = v34
                    } else {
                        v3 -= v26
                        var v35 uint32 = 2147483648
                        _ = v35
                    }
                    v26 = v3
                    _ = fi31
                    _ = v32
                }
                for fi36 := 0; fi36 < 5; fi36++ {
                    switch (fi30 | fi36) {
                    case -9223372036854775808:
                        if false {
                            v37 := GenS3{F0: (v27 >> ((2) & 63)), F1: (fi36 & 0), F2: false}
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            var v38 int8 = -1
                            _ = v37
                            _ = v38
                        }
                    case 32767:
                        v39 := GenS1{F0: 65535, F1: (v3 - 4294967295)}
                        _ = v39
                    default:
                        switch "foo bar" {
                        case "\n\t":
                            v3 -= 850
                        case "hello":
                            var v40 int64 = (v27 / (((v27 | v27)) | 1))
                            switch "a" {
                            case "a":
                                v27 = 84
                            case "hello":
                                fmt.Println(v4, fi30)
                            case "\n\t":
                                var v41 uint16 = 0
                                v40 = (-(v27))
                                if true {
                                    v27 = 1
                                    v42 := GenS4{F0: int16(v26), F1: v41}
                                    _ = v42
                                } else {
                                    if (1 <= 42) {
                                        v43 := GenS1{F0: 50, F1: (^(v26))}
                                        var v44 uint8 = 255
                                        v40 -= (v40 * v27)
                                        _ = v43
                                        _ = v44
                                    }
                                    v45 := GenS3{F0: 42, F1: (fi30 - 0), F2: true}
                                    delete(v1, ("a" + "a"))
                                    _ = v45
                                }
                                v40 += (v40 % ((v40) | 1))
                                if true {
                                    fmt.Println(v4, v26, v40)
                                    v4 = (v4 ^ -1)
                                    var v46 string = "\n\t"
                                    _ = v46
                                }
                                v47 := GenS5{F0: 42}
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                _ = v41
                                _ = v47
                            default:
                                var v48 float32 = -2.25
                                _ = v48
                            }
                            _ = v40
                        default:
                            var v49 int = fi30
                            _ = v49
                        }
                    }
                    var v50 int16 = 42
                    var v51 string = "a"
                    delete(v1, v51)
                    _ = fi36
                    _ = v50
                    _ = v51
                }
                _ = fi30
            }
        }
        fmt.Println(v27, v3)
        v3 = (^(v3))
    } else {
        fmt.Println(v27)
        switch "foo bar" {
        case "hello":
            v4 = v4
        case "\n\t":
            if false {
                v52 := GenS6{F0: ("\n\t" + "hello")}
                v26 += (v3 | v3)
                _ = v52
            } else {
                for fi53 := 0; fi53 < 15; fi53++ {
                    v54 := GenS4{F0: 0, F1: 2}
                    var v55 bool = (195 >= 1)
                    if (true || (((v27 <= int64(v26)) && (215 >= 0)) || (!true))) {
                        for fi56 := 0; fi56 < 16; fi56++ {
                            v57 := GenS5{F0: 573}
                            if ((v4 << ((18446744073709551615) & 63)) > v4) {
                                v1["foo bar"] = v55
                                var v58 int64 = (-(v27))
                                switch "" {
                                case "\n\t":
                                    v4 += v4
                                    _ = len(v2)
                                    fmt.Println(fi53, v4, fi53)
                                case "a":
                                    fmt.Println(fi56, v26, v58)
                                default:
                                    v55 = v55
                                }
                                _ = v58
                            } else {
                                if v55 {
                                    v59 := GenS2{F0: 1, F1: ((("foo bar" + "hello") + "") + ("foo bar" + "foo bar")), F2: 1}
                                    var v60 float64 = -2.25
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    v3 -= (v26 - v3)
                                    fmt.Println(v26)
                                    v61 := v1[("a" + "")]
                                    _ = v59
                                    _ = v60
                                    _ = v61
                                } else {
                                    _ = len(v2)
                                    var v62 int64 = 0
                                    var v63 int8 = -120
                                    _ = v62
                                    _ = v63
                                }
                                v64 := GenS7{F0: ("a" + ("a" + "hello")), F1: 1, F2: 0}
                                _ = v64
                            }
                            var v65 int8 = -26
                            var v66 int64 = v27
                            _ = fi56
                            _ = v57
                            _ = v65
                            _ = v66
                        }
                        var v67 bool = v55
                        _ = v67
                    }
                    _ = fi53
                    _ = v54
                    _ = v55
                }
                var v68 int32 = -2147483648
                _ = v68
            }
        default:
            v69 := GenS8{F0: -1, F1: 4294967295, F2: 65535}
            _ = v69
        }
    }
    switch 1 {
    case 32768:
        v70 := GenS9{F0: 32767}
        _ = v70
    case -748:
        if len(v2) > 0 {
            v2[0] = 9223372036854775807
        }
    default:
        for fi71 := 0; fi71 < 9; fi71++ {
            switch ("hello" + ("a" + "")) {
            case "\n\t":
                for fi72 := 0; fi72 < 6; fi72++ {
                    v27 += (v27 - v27)
                    v26 += v26
                    _ = fi72
                }
            default:
                var v73 int16 = -32768
                _ = v73
            }
            if false {
                if (fi71 >= fi71) {
                    v74 := GenS10{F0: 1, F1: (!true)}
                    fmt.Println(fi71)
                    _ = v74
                } else {
                    if (!(false && true)) {
                        v26 -= (v26 & v3)
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    } else {
                        for fi75 := 0; fi75 < 12; fi75++ {
                            _ = len(v1)
                            _ = fi75
                        }
                    }
                }
                switch -490 {
                case 2147483647:
                    if true {
                        fmt.Println(fi71, v26)
                    }
                case -120:
                    var v76 uint32 = 2
                    _ = v76
                case 9223372036854775807:
                    v3 += v26
                }
            }
            var v77 int64 = v27
            _ = fi71
            _ = v77
        }
    }
    switch 32768 {
    case -9223372036854775808:
        for fi78 := 0; fi78 < 10; fi78++ {
            for fi79 := 0; fi79 < 8; fi79++ {
                if ((v27 != 974) && true) {
                    v3 = (v3 / (((^(v3))) | 1))
                    v4 += v4
                    var v80 string = "hello"
                    _ = v80
                } else {
                    if true {
                        var v81 uint8 = 241
                        _ = v81
                    } else {
                        fmt.Println(v27, v4, v3)
                        switch "foo bar" {
                        case "hello":
                            v3 = v3
                        case "\n\t":
                            var v82 bool = true
                            _ = v82
                        case "":
                            v26 -= (v3 % ((v26) | 1))
                        }
                        if true {
                            v26 = v3
                            v4 += v4
                            v4 = (v4 * v4)
                        }
                    }
                    if len(v2) > 0 {
                        v2[0] = 454
                    }
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    fmt.Println(v4)
                }
                _ = fi79
            }
            fmt.Println(fi78)
            _ = fi78
        }
    }
    if false {
        if false {
            v3 = (v26 | v3)
            if len(v2) > 0 {
                v2[0] = 2147483647
            }
        } else {
            for fi83 := 0; fi83 < 19; fi83++ {
                for fi84 := 0; fi84 < 12; fi84++ {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    if (false && true) {
                        v4 -= 2147483647
                        var v85 uint8 = 127
                        v27 += -817
                        _ = v85
                    }
                    v86 := GenS11{F0: 32767, F1: true, F2: ((fi83 >> ((1) & 63)) <= (^(fi83)))}
                    _ = fi84
                    _ = v86
                }
                _ = fi83
            }
            var v87 float32 = (3.14159 * 1e10)
            _ = v87
        }
        v4 -= (^(v4))
        v88 := GenS6{F0: "a"}
        _ = v88
    }
    v26 = (v26 << ((4294967295) & 63))
    v89 := v1["\n\t"]
    for fi90 := 0; fi90 < 9; fi90++ {
        switch "\n\t" {
        case "foo bar":
            switch fi90 {
            case 32768:
                switch int(v27) {
                case -896:
                    switch (fi90 + fi90) {
                    case 1:
                        v2 = append(v2, (fi90 & 138))
                    }
                case 0:
                    var v91 uint = 534
                    _ = v91
                default:
                    for fi92 := 0; fi92 < 15; fi92++ {
                        var v93 float64 = float64(1e10)
                        if v89 {
                            v94 := GenS4{F0: 1, F1: 2}
                            if false {
                                switch ("a" + (("foo bar" + ("foo bar" + "foo bar")) + "hello")) {
                                case "\n\t":
                                    if ((!v89) && false) {
                                        switch (fi90 / (((-(fi92))) | 1)) {
                                        case 4294967295:
                                            var v95 float32 = 0.1
                                            _ = v95
                                        }
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        var v96 int16 = 1
                                        _ = v96
                                    } else {
                                        var v97 uint16 = 0
                                        if true {
                                            var v98 uint = 255
                                            _ = v98
                                        } else {
                                            v26 = 4294967295
                                        }
                                        _ = v97
                                    }
                                case "hello":
                                    v89 = false
                                default:
                                    switch (fi90 | (fi92 & fi92)) {
                                    case 32767:
                                        v2 = append(v2, fi92)
                                    }
                                }
                                v3 = 253
                            }
                            fmt.Println(v27, v93)
                            _ = v94
                        }
                        var v99 bool = (v93 < v93)
                        _ = fi92
                        _ = v93
                        _ = v99
                    }
                }
            case 127:
                if (!v89) {
                    v100 := GenS12{F0: v27, F1: v89}
                    for fi101 := 0; fi101 < 19; fi101++ {
                        var v102 uint64 = (v26 | v26)
                        fmt.Println(v3, v4, v4)
                        _ = fi101
                        _ = v102
                    }
                    v27 += v27
                    _ = v100
                }
            }
        case "hello":
            v89 = ((v26 << ((255) & 63)) >= v3)
        case "":
            v103 := GenS9{F0: fi90}
            _ = v103
        default:
            fmt.Println(v27, fi90)
        }
        fmt.Println(v89, v26, v3)
        _ = fi90
    }
    for fi104 := 0; fi104 < 9; fi104++ {
        v26 -= v3
        fmt.Println(v27, v89)
        var v105 uint = 630
        _ = fi104
        _ = v105
    }
    for fi106 := 0; fi106 < 13; fi106++ {
        if v89 {
            for fi107 := 0; fi107 < 12; fi107++ {
                fmt.Println(v3)
                var v108 uint32 = 435
                fmt.Println(fi106)
                _ = fi107
                _ = v108
            }
            if (false && ("foo bar" != (("foo bar" + ((("" + "") + ("\n\t" + "foo bar")) + ("foo bar" + "foo bar"))) + "hello"))) {
                switch ((((("a" + "hello") + "a") + "") + "hello") + "hello") {
                case "a":
                    if (!v89) {
                        if (v89 && v89) {
                            v89 = v89
                        }
                        if true {
                            if ((!v89) && v89) {
                                fmt.Println(fi106, v27, fi106)
                            } else {
                                fmt.Println(v27, v4, v89)
                            }
                            fmt.Println(v27, v4, fi106)
                        } else {
                            for fi109 := 0; fi109 < 14; fi109++ {
                                v110 := GenS13{F0: (v27 >= v27), F1: -1}
                                if ((!(!true)) && (255 <= 0)) {
                                    var v111 uint16 = 42
                                    var v112 bool = v89
                                    _ = v111
                                    _ = v112
                                }
                                switch (fi109 << ((1) & 63)) {
                                case 32767:
                                    v26 = 4294967295
                                case -9223372036854775808:
                                    v89 = v89
                                case -256:
                                    _ = len(v2)
                                default:
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                }
                                if (!v89) {
                                    if v89 {
                                        var v113 int16 = 1000
                                        v113 = int16(v4)
                                        v114 := GenS14{F0: v3, F1: 42}
                                        _ = v113
                                        _ = v114
                                    }
                                    v4 = 2147483647
                                }
                                if false {
                                    v115 := GenS12{F0: 9223372036854775807, F1: v89}
                                    v4 = (v4 << ((255) & 63))
                                    _ = len(v2)
                                    _ = v115
                                }
                                _ = fi109
                                _ = v110
                            }
                        }
                        for fi116 := 0; fi116 < 5; fi116++ {
                            if (!true) {
                                if ((("\n\t" + "foo bar") + "a") >= ("hello" + "")) {
                                    v26 -= 2
                                    if len(v2) > 0 {
                                        v2[0] = (fi106 & (fi116 & (fi116 ^ fi116)))
                                    }
                                    var v117 int = 30
                                    var v118 int8 = -73
                                    fmt.Println(fi106, v89, v26)
                                    switch fi116 {
                                    case 32768:
                                        switch fi106 {
                                        case 1:
                                            var v119 int = (v117 + (fi106 & (v117 * fi106)))
                                            _ = v119
                                        case 51:
                                            fmt.Println(v89, v27, v27)
                                        case 32768:
                                            v2 = append(v2, fi106)
                                        default:
                                            v4 = v4
                                            if false {
                                                v4 = v4
                                            }
                                            if len(v2) > 0 {
                                                v2[0] = fi106
                                            }
                                        }
                                    }
                                    _ = v117
                                    _ = v118
                                }
                            } else {
                                if len(v2) > 0 {
                                    v2[0] = fi106
                                }
                                switch (fi116 >> ((360) & 63)) {
                                case 42:
                                    v89 = ((v3 + 501) > 4294967295)
                                    v120 := GenS2{F0: 4294967295, F1: (("hello" + ("\n\t" + (((("hello" + "a") + "foo bar") + ("\n\t" + "hello")) + ("\n\t" + (("" + (("foo bar" + "a") + ((("a" + "foo bar") + (("foo bar" + "foo bar") + ("a" + "foo bar"))) + "foo bar"))) + (((("foo bar" + ("hello" + "hello")) + ("" + "foo bar")) + ((((("\n\t" + "foo bar") + "a") + ("hello" + ((("hello" + "\n\t") + "") + "hello"))) + "") + "\n\t")) + "a")))))) + "foo bar"), F2: 65535}
                                    switch fi106 {
                                    case 496:
                                        v3 -= (^(v3))
                                        v26 += v26
                                    default:
                                        v121 := GenS15{F0: ("" + ""), F1: v26}
                                        _ = v121
                                    }
                                    _ = v120
                                case 9223372036854775807:
                                    v4 -= 2147483647
                                case -2147483648:
                                    v26 = (v26 & (v3 + 42))
                                }
                            }
                            _ = fi116
                        }
                    }
                default:
                    var v122 uint64 = (v26 + 355)
                    _ = v122
                }
            } else {
                switch int(v27) {
                case 4294967295:
                    v123 := GenS15{F0: ("a" + ("\n\t" + "hello")), F1: 4294967295}
                    _ = v123
                }
            }
            var v124 uint = 488
            _ = v124
        }
        var v125 float64 = 1.5
        v26 = 2
        _ = fi106
        _ = v125
    }
    if (!(897 == v4)) {
        for fi126 := 0; fi126 < 6; fi126++ {
            var v127 bool = (!(65535 >= 173))
            var v128 int64 = v27
            _ = fi126
            _ = v127
            _ = v128
        }
        if v89 {
            v3 -= (^(v3))
            if (!v89) {
                var v129 string = ""
                _ = v129
            }
        } else {
            switch "\n\t" {
            case "a":
                var v130 int32 = (v4 & -309)
                _ = v130
            case "hello":
                if v89 {
                    for fi131 := 0; fi131 < 5; fi131++ {
                        v132 := GenS5{F0: 800}
                        switch (fi131 % ((fi131) | 1)) {
                        case 4294967295:
                            delete(v1, "hello")
                        }
                        v133 := GenS16{F0: 2}
                        _ = fi131
                        _ = v132
                        _ = v133
                    }
                    var v134 int64 = v27
                    _ = v134
                }
            case "foo bar":
                var v135 int32 = v4
                _ = v135
            }
            v136 := v1["hello"]
            _ = v136
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v89:", v89)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v89
}
