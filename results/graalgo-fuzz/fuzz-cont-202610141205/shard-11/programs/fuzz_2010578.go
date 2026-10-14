package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string; F2 bool }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint32; F1 uint8 }
type GenS5 struct { F0 int64; F1 int32; F2 uint32 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 int; F1 uint32; F2 int16 }
type GenS8 struct { F0 bool; F1 uint64; F2 uint64 }
type GenS9 struct { F0 uint64; F1 int64; F2 bool }
type GenS10 struct { F0 string; F1 int }
type GenS11 struct { F0 uint }
type GenS12 struct { F0 uint; F1 int8 }
type GenS13 struct { F0 bool; F1 int64 }
type GenS14 struct { F0 string }
type GenS15 struct { F0 int32; F1 uint; F2 uint8 }
type GenS16 struct { F0 int; F1 uint8 }
type GenS17 struct { F0 uint16 }
type GenS18 struct { F0 string; F1 uint32; F2 bool }
type GenS19 struct { F0 int8; F1 uint16; F2 string }
type GenS20 struct { F0 bool }
type GenS21 struct { F0 int8; F1 uint }
type GenS22 struct { F0 uint16; F1 int; F2 uint32 }
type GenS23 struct { F0 int16; F1 uint64 }
type GenS24 struct { F0 int32 }
type GenS25 struct { F0 int64; F1 uint32 }
func genFn0(p0 string, p1 uint64) uint {
    var v4 int64 = 1
    v5 := GenS1{F0: -9223372036854775808, F1: "foo bar", F2: (32767 < -1)}
    v6 := []int{-1}
    var v7 uint8 = 255
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return 42
}
func genFn1(p0 bool, p1 string) uint8 {
    var v8 int64 = -9223372036854775808
    v9 := GenS1{F0: 266, F1: p1, F2: p0}
    v8 = v8
    p0 = p0
    v8 = v8
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return 1
}

func main() {
    v1 := []uint32{968, 42, 4294967295, 1}
    var v2 uint16 = 962
    v3 := make(map[string]int)
    v2 -= 0
    switch "foo bar" {
    case "":
        v2 -= 1
    case "hello":
        v10 := GenS2{F0: -32768}
        _ = v10
    case "\n\t":
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    switch ("hello" + ("\n\t" + "")) {
    case "hello":
        for fi11 := 0; fi11 < 7; fi11++ {
            switch (fi11 / ((fi11) | 1)) {
            case 2147483647:
                for fi12 := 0; fi12 < 6; fi12++ {
                    _ = len(v3)
                    switch "hello" {
                    case "foo bar":
                        if (-2.25 >= 0.0) {
                            if false {
                                var v13 string = "\n\t"
                                v14 := GenS3{F0: 2}
                                _ = v13
                                _ = v14
                            }
                            if true {
                                switch "foo bar" {
                                case "hello":
                                    var v15 float32 = 0.1
                                    _ = v15
                                case "":
                                    var v16 uint32 = 0
                                    v16 += (v16 << ((255) & 63))
                                    if (int8(fi12) != 127) {
                                        fmt.Println(v2)
                                    }
                                    v16 = 739
                                    _ = v16
                                case "\n\t":
                                    if false {
                                        if true {
                                            switch (fi12 << ((0) & 63)) {
                                            case -699:
                                                var v17 uint16 = v2
                                                _ = v17
                                            case -689:
                                                v18 := GenS3{F0: 2}
                                                _ = v18
                                            case 51:
                                                switch ("hello" + (("" + "") + "\n\t")) {
                                                case "hello":
                                                    var v19 string = ""
                                                    _ = v19
                                                case "":
                                                    if false {
                                                        if true {
                                                            fmt.Println(fi11)
                                                        }
                                                        var v20 int64 = -9223372036854775808
                                                        _ = v20
                                                    } else {
                                                        if len(v1) > 0 {
                                                            v1[0] = 4294967295
                                                        }
                                                    }
                                                case "foo bar":
                                                    var v21 int16 = 0
                                                    _ = v21
                                                default:
                                                    v2 = (v2 * v2)
                                                }
                                            default:
                                                if false {
                                                    var v22 float32 = 0.1
                                                    _ = v22
                                                } else {
                                                    switch fi12 {
                                                    case 4294967295:
                                                        var v23 bool = (-1 <= 0)
                                                        _ = v23
                                                    default:
                                                        v2 -= 65535
                                                    }
                                                    fmt.Println(fi12, fi11)
                                                    var v24 uint64 = 4294967295
                                                    if (("foo bar" >= "\n\t") || false) {
                                                        v24 -= v24
                                                        if (4294967295 > 631) {
                                                            v24 = v24
                                                        } else {
                                                            fmt.Println(fi11, fi11)
                                                            if (!(false && false)) {
                                                                v24 += (v24 - (^(v24)))
                                                            }
                                                            v25 := GenS4{F0: 261, F1: 127}
                                                            if (!(3.14159 <= -2.25)) {
                                                                switch "\n\t" {
                                                                case "\n\t":
                                                                    var v26 uint = 0
                                                                    _ = v26
                                                                default:
                                                                    v2 = (v2 * 1)
                                                                }
                                                            }
                                                            if (v24 == (v24 - 632)) {
                                                                v24 = 0
                                                            }
                                                            _ = v25
                                                        }
                                                        v2 -= v2
                                                    } else {
                                                        if false {
                                                            var v27 uint = uint(v24)
                                                            _ = v27
                                                        }
                                                        fmt.Println(v24)
                                                        v28 := GenS5{F0: 9223372036854775807, F1: 1, F2: 312}
                                                        _ = v28
                                                    }
                                                    v29 := GenS6{F0: fi11}
                                                    v30 := GenS1{F0: fi12, F1: "hello", F2: (false && true)}
                                                    if (true && (((!false) && true) && true)) {
                                                        v31 := GenS5{F0: 95, F1: 210, F2: 369}
                                                        var v32 uint64 = (v24 << ((18446744073709551615) & 63))
                                                        _ = v31
                                                        _ = v32
                                                    } else {
                                                        _ = len(v3)
                                                        if true {
                                                            v3["hello"] = 9223372036854775807
                                                            var v33 uint = 0
                                                            _ = v33
                                                        }
                                                        if (127 == -22) {
                                                            v34 := GenS7{F0: (fi12 + (fi12 >> ((822) & 63))), F1: 196, F2: -1}
                                                            _ = v34
                                                        }
                                                    }
                                                    var v35 float32 = float32(0.0)
                                                    fmt.Println(v2, v24)
                                                    fmt.Println(v35, v35)
                                                    fmt.Println(fi11)
                                                    _ = v24
                                                    _ = v29
                                                    _ = v30
                                                    _ = v35
                                                }
                                            }
                                            v36 := GenS6{F0: fi12}
                                            _ = v36
                                        } else {
                                            if (((!false) || true) && false) {
                                                delete(v3, "\n\t")
                                            }
                                        }
                                        switch ("a" + "foo bar") {
                                        case "foo bar":
                                            v2 = 42
                                        }
                                        var v37 int32 = 0
                                        v38 := GenS8{F0: (0 < 1), F1: 2, F2: 18446744073709551615}
                                        var v39 float64 = 3.14159
                                        _ = v37
                                        _ = v38
                                        _ = v39
                                    }
                                default:
                                    v40 := GenS9{F0: 811, F1: 67, F2: (!true)}
                                    _ = v40
                                }
                                var v41 bool = false
                                _ = v41
                            } else {
                                if (false || (4294967295 != genFn0("", 1))) {
                                    v42 := GenS9{F0: 418, F1: -585, F2: (-1 <= -61)}
                                    _ = v42
                                } else {
                                    switch fi12 {
                                    case 9223372036854775807:
                                        v2 = 32767
                                    default:
                                        fmt.Println(fi11)
                                    }
                                    v2 = (v2 & 1)
                                    var v43 float64 = 1.5
                                    _ = v43
                                }
                                switch fi11 {
                                case 432:
                                    if true {
                                        switch "hello" {
                                        case "\n\t":
                                            var v44 bool = true
                                            _ = v44
                                        case "a":
                                            fmt.Println(fi11, v2, fi12)
                                        case "hello":
                                            v45 := GenS1{F0: 969, F1: "foo bar", F2: ((175 >= 162) || true)}
                                            _ = v45
                                        }
                                    } else {
                                        switch -152 {
                                        case 168:
                                            v2 += 407
                                        case 128:
                                            if (false || (("" + "") <= "")) {
                                                v2 = (v2 % (((v2 ^ 32767)) | 1))
                                                v2 += v2
                                                var v46 uint8 = 161
                                                if len(v1) > 0 {
                                                    v1[0] = 2147483647
                                                }
                                                var v47 int64 = 0
                                                var v48 uint64 = 828
                                                var v49 float64 = -2.25
                                                v48 = uint64(v46)
                                                _ = v46
                                                _ = v47
                                                _ = v48
                                                _ = v49
                                            } else {
                                                var v50 uint16 = v2
                                                v50 -= (v50 & 65535)
                                                var v51 uint8 = 255
                                                if ((v2 - uint16(v51)) < v2) {
                                                    if (7 < 898) {
                                                        var v52 int32 = -2147483648
                                                        delete(v3, "")
                                                        if true {
                                                            var v53 int = 128
                                                            fmt.Println(fi11)
                                                            fmt.Println(v2, v53, v53)
                                                            v50 += (v2 / (((v50 & v50)) | 1))
                                                            switch 9223372036854775807 {
                                                            case 0:
                                                                v53 += v53
                                                            case 9223372036854775807:
                                                                switch (v53 * fi11) {
                                                                case 128:
                                                                    fmt.Println(v2, v52)
                                                                case 32767:
                                                                    if len(v1) > 0 {
                                                                        v1[0] = 2
                                                                    }
                                                                case 2147483647:
                                                                    v53 = (v53 / ((int(v2)) | 1))
                                                                default:
                                                                    _ = len(v3)
                                                                }
                                                            case 212:
                                                                switch "\n\t" {
                                                                case "\n\t":
                                                                    v54 := GenS5{F0: 0, F1: (-(v52)), F2: 0}
                                                                    if len(v1) > 0 {
                                                                        v1[0] = uint32(v2)
                                                                    }
                                                                    _ = v54
                                                                default:
                                                                    v2 = v2
                                                                }
                                                            default:
                                                                fmt.Println(v50, v52, fi12)
                                                            }
                                                            _ = v53
                                                        } else {
                                                            if ((1 >= 9223372036854775807) && true) {
                                                                if len(v1) >= 2 {
                                                                    v1 = v1[1:len(v1)]
                                                                }
                                                            } else {
                                                                v50 -= v2
                                                                var v55 float64 = 1.5
                                                                _ = v55
                                                            }
                                                            if true {
                                                                var v56 float32 = float32((-2.25 * float64(1.5)))
                                                                _ = v56
                                                            }
                                                            if (2 != genFn0("", 4294967295)) {
                                                                var v57 int16 = 32767
                                                                _ = v57
                                                            }
                                                            v58 := GenS4{F0: 1, F1: v51}
                                                            v51 = (v51 >> ((4294967295) & 63))
                                                            var v59 string = "foo bar"
                                                            _ = v58
                                                            _ = v59
                                                        }
                                                        _ = v52
                                                    }
                                                    v1 = append(v1, 0)
                                                } else {
                                                    if (313 > 18446744073709551615) {
                                                        v50 = (^(v2))
                                                        var v60 float32 = 1.5
                                                        if (108 <= 1) {
                                                            var v61 int16 = 1
                                                            v2 += v50
                                                            _ = v61
                                                        }
                                                        if (true && false) {
                                                            fmt.Println(fi11, v2, v51)
                                                        }
                                                        _ = v60
                                                    } else {
                                                        fmt.Println(fi12)
                                                        v2 += (v50 | (v2 | (^(v2))))
                                                    }
                                                    v2 = (v2 ^ v50)
                                                    fmt.Println(fi12)
                                                }
                                                _ = v50
                                                _ = v51
                                            }
                                        default:
                                            fmt.Println(fi11, fi11)
                                        }
                                        var v62 string = "a"
                                        _ = v62
                                    }
                                case 127:
                                    switch "\n\t" {
                                    case "foo bar":
                                        v2 = v2
                                    case "":
                                        var v63 uint16 = (v2 & (^(v2)))
                                        _ = v63
                                    case "hello":
                                        var v64 uint8 = 127
                                        var v65 uint32 = 0
                                        _ = v64
                                        _ = v65
                                    default:
                                        if false {
                                            var v66 float32 = -2.25
                                            _ = v66
                                        } else {
                                            fmt.Println(fi12, fi11)
                                            fmt.Println(fi12, v2)
                                        }
                                    }
                                default:
                                    _ = len(v3)
                                    var v67 uint16 = 65535
                                    _ = v67
                                }
                                v68 := GenS10{F0: ("" + "hello"), F1: fi12}
                                _ = v68
                            }
                        }
                    }
                    _ = fi12
                }
            default:
                switch -2147483648 {
                case 128:
                    v2 -= 1
                }
            }
            _ = fi11
        }
    case "":
        for fi69 := 0; fi69 < 8; fi69++ {
            var v70 int32 = 89
            v70 = v70
            fmt.Println(v70)
            _ = fi69
            _ = v70
        }
    case "a":
        var v71 uint64 = 0
        _ = v71
    default:
        fmt.Println(v2)
    }
    switch "\n\t" {
    case "\n\t":
        for fi72 := 0; fi72 < 6; fi72++ {
            switch (fi72 + fi72) {
            case 9223372036854775807:
                var v73 int64 = -640
                _ = v73
            default:
                switch "hello" {
                case "foo bar":
                    fmt.Println(v2, v2)
                case "a":
                    if ((fi72 - 4294967295) == fi72) {
                        _ = len(v1)
                        v1 = append(v1, 2147483647)
                        if false {
                            v74 := GenS4{F0: 317, F1: 250}
                            var v75 int8 = 42
                            _ = v74
                            _ = v75
                        } else {
                            fmt.Println(fi72, v2)
                            v76 := v3["hello"]
                            _ = v76
                        }
                    } else {
                        if (987 >= 2) {
                            _ = len(v3)
                        } else {
                            for fi77 := 0; fi77 < 18; fi77++ {
                                v2 += (v2 << ((1) & 63))
                                switch "hello" {
                                case "":
                                    fmt.Println(fi77, fi77)
                                case "\n\t":
                                    var v78 int16 = int16(v2)
                                    _ = v78
                                }
                                _ = fi77
                            }
                            fmt.Println(v2)
                            for fi79 := 0; fi79 < 10; fi79++ {
                                var v80 uint16 = (^(v2))
                                _ = fi79
                                _ = v80
                            }
                        }
                    }
                default:
                    v2 -= 57
                }
            }
            var v81 int64 = 42
            _ = fi72
            _ = v81
        }
    case "foo bar":
        if (!true) {
            if false {
                var v82 int32 = 0
                _ = v82
            }
        }
    case "hello":
        if false {
            v83 := GenS11{F0: 0}
            var v84 uint16 = 255
            for fi85 := 0; fi85 < 10; fi85++ {
                v2 = (v2 | v84)
                v2 = (^(v2))
                for fi86 := 0; fi86 < 16; fi86++ {
                    v2 = v2
                    if len(v1) > 0 {
                        v1[0] = 209
                    }
                    _ = fi86
                }
                _ = fi85
            }
            _ = v83
            _ = v84
        } else {
            fmt.Println(v2)
            if (!false) {
                v2 += 545
                for fi87 := 0; fi87 < 17; fi87++ {
                    if (true && (true || true)) {
                        var v88 float32 = ((0.0 - 0.1) / -2.25)
                        v89 := GenS7{F0: fi87, F1: 2147483647, F2: 1}
                        _ = v88
                        _ = v89
                    }
                    if false {
                        fmt.Println(v2, v2)
                        for fi90 := 0; fi90 < 20; fi90++ {
                            var v91 float32 = (-2.25 - -2.25)
                            _ = fi90
                            _ = v91
                        }
                        var v92 uint32 = 65535
                        _ = v92
                    } else {
                        if false {
                            if (-212 > fi87) {
                                v2 -= 829
                                v93 := GenS12{F0: 527, F1: 127}
                                switch (fi87 & fi87) {
                                case 9223372036854775807:
                                    var v94 float32 = (float32(float64(1.5)) + (0.1 / (1.5 + 1.5)))
                                    _ = v94
                                case -694:
                                    for fi95 := 0; fi95 < 19; fi95++ {
                                        var v96 int64 = 74
                                        _ = fi95
                                        _ = v96
                                    }
                                case 2147483647:
                                    if false {
                                        for fi97 := 0; fi97 < 16; fi97++ {
                                            var v98 bool = false
                                            v98 = false
                                            if len(v1) > 0 {
                                                v1[0] = 4294967295
                                            }
                                            _ = fi97
                                            _ = v98
                                        }
                                    } else {
                                        v2 -= (v2 >> ((4294967295) & 63))
                                    }
                                }
                                _ = v93
                            } else {
                                delete(v3, "foo bar")
                            }
                            var v99 int32 = 42
                            _ = v99
                        } else {
                            switch "foo bar" {
                            case "foo bar":
                                switch (fi87 >> ((221) & 63)) {
                                case -175:
                                    for fi100 := 0; fi100 < 2; fi100++ {
                                        v101 := GenS13{F0: ((1e10 * (3.14159 + 1.5)) == 1e10), F1: 0}
                                        _ = fi100
                                        _ = v101
                                    }
                                }
                            case "hello":
                                v102 := GenS5{F0: 478, F1: 42, F2: 42}
                                _ = v102
                            case "":
                                if false {
                                    v2 -= (v2 ^ (v2 % ((v2) | 1)))
                                }
                            default:
                                fmt.Println(v2, fi87)
                            }
                            v103 := GenS3{F0: 2}
                            _ = v103
                        }
                        fmt.Println(v2, v2)
                    }
                    var v104 int = fi87
                    _ = fi87
                    _ = v104
                }
            } else {
                v105 := GenS3{F0: 65535}
                _ = v105
            }
        }
    }
    fmt.Println(v2)
    var v106 uint32 = 42
    v3["foo bar"] = 127
    var v107 uint = 65535
    for fi108 := 0; fi108 < 5; fi108++ {
        for fi109 := 0; fi109 < 20; fi109++ {
            fmt.Println(fi108, fi109)
            if ((v2 % ((v2) | 1)) != (v2 - v2)) {
                switch (fi109 % (((fi109 << ((18446744073709551615) & 63))) | 1)) {
                case 0:
                    if true {
                        var v110 uint = 18446744073709551615
                        var v111 string = "foo bar"
                        _ = v110
                        _ = v111
                    }
                case 127:
                    if ((v2 >> ((1) & 63)) == uint16(v107)) {
                        var v112 uint8 = 216
                        _ = v112
                    } else {
                        var v113 int64 = 9223372036854775807
                        if true {
                            v114 := GenS14{F0: "a"}
                            _ = v114
                        }
                        delete(v3, "hello")
                        _ = v113
                    }
                default:
                    var v115 int16 = -32768
                    _ = v115
                }
            } else {
                if true {
                    v116 := GenS8{F0: (((1e10 - (1.5 + ((-2.25 - 0.1) - 1.5))) / 1.5) <= 3.14159), F1: 175, F2: 2}
                    switch ("hello" + ("foo bar" + "foo bar")) {
                    case "foo bar":
                        if (255 < (v2 & v2)) {
                            var v117 uint8 = 1
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            if false {
                                var v118 int16 = 558
                                _ = v118
                            } else {
                                var v119 string = "hello"
                                if false {
                                    v119 = ("\n\t" + ("" + v119))
                                    switch v119 {
                                    case "a":
                                        var v120 uint64 = 843
                                        _ = v120
                                    case "hello":
                                        v106 -= 24
                                    case "foo bar":
                                        fmt.Println(v2, fi108)
                                        switch ((("hello" + ("\n\t" + ((v119 + (v119 + v119)) + ("" + "a")))) + v119) + (v119 + "")) {
                                        case "\n\t":
                                            v119 = ("hello" + v119)
                                        default:
                                            var v121 int16 = -967
                                            _ = v121
                                        }
                                    default:
                                        var v122 float64 = (((1.5 + (0.1 + 1e10)) + 3.14159) + 3.14159)
                                        _ = v122
                                    }
                                } else {
                                    v1 = append(v1, (v106 & (^(v106))))
                                    switch 709 {
                                    case 6:
                                        v119 = (v119 + "a")
                                    case 32767:
                                        fmt.Println(v106, v117, v119)
                                    case 127:
                                        v123 := GenS15{F0: 61, F1: 413, F2: v117}
                                        _ = v123
                                    }
                                    if (v119 != v119) {
                                        fmt.Println(v117)
                                        switch v119 {
                                        case "foo bar":
                                            v124 := GenS16{F0: (fi108 - (fi108 >> ((v107) & 63))), F1: (v117 | (v117 << ((588) & 63)))}
                                            _ = v124
                                        case "":
                                            var v125 string = ((v119 + v119) + "")
                                            _ = v125
                                        case "\n\t":
                                            v126 := GenS17{F0: (v2 * 32767)}
                                            _ = v126
                                        default:
                                            v2 -= v2
                                        }
                                    } else {
                                        fmt.Println(v106)
                                        var v127 float32 = 3.14159
                                        _ = v127
                                    }
                                }
                                if (!(699 < 2147483647)) {
                                    v1 = append(v1, v106)
                                } else {
                                    var v128 int16 = 1
                                    if true {
                                        switch fi108 {
                                        case -2147483648:
                                            switch "" {
                                            case "hello":
                                                var v129 int8 = -128
                                                _ = v129
                                            case "a":
                                                switch (fi108 & (-(fi108))) {
                                                case 4294967295:
                                                    v128 = (v128 / ((32767) | 1))
                                                default:
                                                    v2 = 2
                                                }
                                            case "":
                                                v130 := GenS8{F0: ((v119 + v119) < v119), F1: 18446744073709551615, F2: 393}
                                                _ = v130
                                            default:
                                                v119 = ""
                                                var v131 uint = v107
                                                _ = v131
                                            }
                                        case -969:
                                            fmt.Println(fi109, v107, v107)
                                        }
                                    }
                                    v1 = append(v1, (v106 * (v106 - v106)))
                                    _ = v128
                                }
                                v117 += (v117 << ((713) & 63))
                                switch 32767 {
                                case 128:
                                    v132 := GenS5{F0: -171, F1: 0, F2: 2}
                                    _ = v132
                                }
                                v133 := GenS18{F0: ((v119 + "\n\t") + "\n\t"), F1: 961, F2: (float32(0.1) <= 1e10)}
                                if true {
                                    v134 := GenS19{F0: 51, F1: v2, F2: v119}
                                    switch "\n\t" {
                                    case "a":
                                        switch ("foo bar" + v119) {
                                        case "hello":
                                            v119 = (v119 + v119)
                                        case "a":
                                            v135 := GenS14{F0: ("\n\t" + "")}
                                            _ = v135
                                        case "\n\t":
                                            switch v119 {
                                            case "\n\t":
                                                v107 = v107
                                            case "hello":
                                                switch (("" + "a") + v119) {
                                                case "":
                                                    v2 = (v2 | (v2 * v2))
                                                }
                                            case "a":
                                                if (42 <= 42) {
                                                    if false {
                                                        switch ((v119 + (v119 + v119)) + "\n\t") {
                                                        case "hello":
                                                            switch fi108 {
                                                            case 1:
                                                                v117 = 2
                                                            default:
                                                                var v136 string = v119
                                                                v137 := GenS8{F0: (1 == 1), F1: 931, F2: 0}
                                                                switch v136 {
                                                                case "hello":
                                                                    v119 = ((v136 + (v136 + v136)) + ("" + "hello"))
                                                                default:
                                                                    var v138 uint32 = v106
                                                                    _ = v138
                                                                }
                                                                _ = v136
                                                                _ = v137
                                                            }
                                                        case "a":
                                                            switch v119 {
                                                            case "\n\t":
                                                                if (1 != -931) {
                                                                    if len(v1) > 0 {
                                                                        v1[0] = v106
                                                                    }
                                                                }
                                                            case "hello":
                                                                v117 = (v117 | (^(v117)))
                                                            }
                                                        case "\n\t":
                                                            if (v2 < (v2 - 32767)) {
                                                                if (true || (((((0.1 / 3.14159) - 3.14159) * (1.5 * 0.1)) + (3.14159 / 1e10)) < 0.0)) {
                                                                    v106 -= uint32(v117)
                                                                    v139 := v3[v119]
                                                                    if len(v1) >= 2 {
                                                                        v1 = v1[1:len(v1)]
                                                                    }
                                                                    _ = v139
                                                                } else {
                                                                    var v140 int16 = 32767
                                                                    _ = v140
                                                                }
                                                                var v141 bool = false
                                                                fmt.Println(v141)
                                                                _ = v141
                                                            }
                                                        default:
                                                            if ((v119 + v119) != v119) {
                                                                var v142 int32 = 1
                                                                switch v119 {
                                                                case "a":
                                                                    v1 = append(v1, 4294967295)
                                                                default:
                                                                    v143 := GenS12{F0: genFn0((v119 + v119), uint64(v2)), F1: 42}
                                                                    _ = v143
                                                                }
                                                                _ = v142
                                                            }
                                                        }
                                                        if (0 == 42) {
                                                            v106 += v106
                                                            if len(v1) > 0 {
                                                                v1[0] = (^(v106))
                                                            }
                                                        } else {
                                                            var v144 uint = (v107 / ((v107) | 1))
                                                            if (false || true) {
                                                                v145 := GenS4{F0: (v106 ^ (v106 / (((v106 >> ((genFn0(v119, 4294967295)) & 63))) | 1))), F1: v117}
                                                                v117 = (v117 << ((42) & 63))
                                                                v2 = v2
                                                                _ = v145
                                                            }
                                                            _ = v144
                                                        }
                                                        v146 := GenS20{F0: ((v119 > ((("a" + v119) + "foo bar") + v119)) || (1 >= (v117 * v117)))}
                                                        var v147 uint64 = 42
                                                        v148 := GenS8{F0: true, F1: uint64(v107), F2: (^(v147))}
                                                        v2 -= v2
                                                        fmt.Println(fi109, v119)
                                                        _ = v146
                                                        _ = v147
                                                        _ = v148
                                                    } else {
                                                        v119 = v119
                                                        fmt.Println(v117, v106)
                                                        switch v119 {
                                                        case "":
                                                            fmt.Println(v119)
                                                        case "a":
                                                            var v149 int32 = 0
                                                            _ = v149
                                                        case "foo bar":
                                                            v117 = (v117 - v117)
                                                        default:
                                                            var v150 float32 = -2.25
                                                            _ = v150
                                                        }
                                                    }
                                                    v106 = 42
                                                    v151 := GenS4{F0: 0, F1: (v117 + 1)}
                                                    _ = v151
                                                } else {
                                                    switch v119 {
                                                    case "foo bar":
                                                        v152 := GenS12{F0: 42, F1: 100}
                                                        var v153 string = v119
                                                        v117 += 23
                                                        _ = v152
                                                        _ = v153
                                                    }
                                                    var v154 uint16 = (^(v2))
                                                    if (-1 > 1) {
                                                        var v155 uint = v107
                                                        if (!(!(float32(float64(3.14159)) > (1.5 * 1e10)))) {
                                                            v117 = genFn1((365 < 4294967295), v119)
                                                            var v156 uint = (v155 % (((v107 >> (((v107 << ((v155) & 63))) & 63))) | 1))
                                                            switch v119 {
                                                            case "\n\t":
                                                                if true {
                                                                    var v157 uint8 = uint8(v2)
                                                                    _ = len(v3)
                                                                    v158 := GenS17{F0: 42}
                                                                    _ = v157
                                                                    _ = v158
                                                                }
                                                            case "":
                                                                var v159 float64 = float64(1e10)
                                                                _ = v159
                                                            case "hello":
                                                                _ = len(v1)
                                                            default:
                                                                switch v119 {
                                                                case "hello":
                                                                    var v160 bool = (32767 < 703)
                                                                    if (562 >= 449) {
                                                                        var v161 int8 = 1
                                                                        var v162 int16 = 1
                                                                        _ = v161
                                                                        _ = v162
                                                                    } else {
                                                                        if true {
                                                                            if false {
                                                                                delete(v3, (v119 + ""))
                                                                                fmt.Println(v106)
                                                                                if ((!(!((!false) && v160))) || (uint16(v117) < 255)) {
                                                                                    var v163 int16 = 818
                                                                                    switch fi109 {
                                                                                    case 4294967295:
                                                                                        var v164 uint = (^(v107))
                                                                                        _ = v164
                                                                                    case 117:
                                                                                        v119 = "hello"
                                                                                        if v160 {
                                                                                            if true {
                                                                                                fmt.Println(v2, v106)
                                                                                            }
                                                                                            v154 = (^(v154))
                                                                                            v165 := GenS21{F0: -128, F1: (^(v107))}
                                                                                            _ = v165
                                                                                        } else {
                                                                                            v107 = v156
                                                                                        }
                                                                                        fmt.Println(v107, v106, v154)
                                                                                        fmt.Println(fi108)
                                                                                        switch (v119 + v119) {
                                                                                        case "\n\t":
                                                                                            v156 = (v156 - (v156 * v107))
                                                                                        case "a":
                                                                                            switch v119 {
                                                                                            case "hello":
                                                                                                _ = len(v3)
                                                                                                v155 = v155
                                                                                                v166 := v3[((((v119 + v119) + (v119 + (v119 + (v119 + ("hello" + ""))))) + "\n\t") + "\n\t")]
                                                                                                var v167 int8 = 42
                                                                                                var v168 uint = v156
                                                                                                _ = v166
                                                                                                _ = v167
                                                                                                _ = v168
                                                                                            case "foo bar":
                                                                                                v154 += uint16(v107)
                                                                                            default:
                                                                                                v169 := GenS22{F0: (v2 % (((v154 - 65535)) | 1)), F1: fi109, F2: v106}
                                                                                                switch "a" {
                                                                                                case "foo bar":
                                                                                                    v154 = (^(v154))
                                                                                                default:
                                                                                                    v170 := GenS5{F0: 42, F1: 413, F2: (v106 | 230)}
                                                                                                    _ = v170
                                                                                                }
                                                                                                _ = v169
                                                                                            }
                                                                                        }
                                                                                        if (v119 <= v119) {
                                                                                            v156 += (v155 ^ (v155 >> ((uint(v117)) & 63)))
                                                                                            var v171 uint = (v107 / ((uint(v154)) | 1))
                                                                                            v172 := GenS4{F0: 1, F1: (^(v117))}
                                                                                            _ = v171
                                                                                            _ = v172
                                                                                        }
                                                                                    case -987:
                                                                                        var v173 bool = (genFn0(v119, 42) == v107)
                                                                                        _ = v173
                                                                                    }
                                                                                    _ = v163
                                                                                } else {
                                                                                    v174 := GenS17{F0: (^(v2))}
                                                                                    var v175 uint8 = genFn1((v160 && v160), v119)
                                                                                    if v160 {
                                                                                        v175 = (v175 & v117)
                                                                                        var v176 int16 = -32768
                                                                                        v154 -= 32767
                                                                                        _ = v176
                                                                                    } else {
                                                                                        var v177 int64 = 1
                                                                                        _ = v177
                                                                                    }
                                                                                    var v178 string = ("a" + "foo bar")
                                                                                    fmt.Println(v106)
                                                                                    _ = v174
                                                                                    _ = v175
                                                                                    _ = v178
                                                                                }
                                                                            } else {
                                                                                fmt.Println(v154, v160)
                                                                            }
                                                                            if ((-32768 < 1) || ((!((v155 << (((v155 ^ (v156 % (((v155 >> (((v107 * (v107 << ((934) & 63)))) & 63))) | 1)))) & 63)) == (^(v155)))) || (1e10 >= float64(((3.14159 * (1.5 * 1.5)) / 1.5))))) {
                                                                                var v179 int = fi108
                                                                                var v180 uint64 = 42
                                                                                _ = v179
                                                                                _ = v180
                                                                            }
                                                                            if true {
                                                                                if false {
                                                                                    v156 = genFn0(v119, uint64(v154))
                                                                                    fmt.Println(v117)
                                                                                    switch (v119 + (v119 + "\n\t")) {
                                                                                    case "hello":
                                                                                        v1 = append(v1, (v106 << (((v107 % ((v155) | 1))) & 63)))
                                                                                    case "a":
                                                                                        v107 = (v107 - (^(v107)))
                                                                                    case "":
                                                                                        _ = len(v1)
                                                                                    }
                                                                                } else {
                                                                                    if true {
                                                                                        var v181 int32 = 1
                                                                                        v2 = (v154 & (v154 & (v2 - (v2 - 42))))
                                                                                        v182 := GenS8{F0: (42 > -32768), F1: 0, F2: 4294967295}
                                                                                        _ = v181
                                                                                        _ = v182
                                                                                    } else {
                                                                                        switch "hello" {
                                                                                        case "a":
                                                                                            if v160 {
                                                                                                v183 := GenS2{F0: 0}
                                                                                                _ = v183
                                                                                            } else {
                                                                                                var v184 float32 = 0.0
                                                                                                var v185 int = 42
                                                                                                _ = v184
                                                                                                _ = v185
                                                                                            }
                                                                                        case "":
                                                                                            if false {
                                                                                                v2 = (v2 % ((212) | 1))
                                                                                            }
                                                                                        case "hello":
                                                                                            var v186 float64 = (((3.14159 - ((0.1 * 3.14159) + (0.0 / 0.1))) + 1.5) - (1e10 / 1e10))
                                                                                            _ = v186
                                                                                        default:
                                                                                            var v187 int8 = -128
                                                                                            _ = v187
                                                                                        }
                                                                                        if v160 {
                                                                                            v106 = v106
                                                                                            var v188 uint32 = v106
                                                                                            _ = v188
                                                                                        }
                                                                                    }
                                                                                }
                                                                            } else {
                                                                                v107 = v107
                                                                            }
                                                                        }
                                                                    }
                                                                    if (false || (false || (1e10 == -2.25))) {
                                                                        if true {
                                                                            fmt.Println(v155, v160)
                                                                            v119 = ""
                                                                            fmt.Println(v156, v155)
                                                                            var v189 uint = v155
                                                                            if len(v1) > 0 {
                                                                                v1[0] = v106
                                                                            }
                                                                            _ = v189
                                                                        }
                                                                    } else {
                                                                        _ = len(v3)
                                                                        var v190 uint = genFn0(v119, 4294967295)
                                                                        v119 = "a"
                                                                        switch fi108 {
                                                                        case 1:
                                                                            v190 = v107
                                                                        }
                                                                        _ = v190
                                                                    }
                                                                    _ = v160
                                                                case "foo bar":
                                                                    v191 := v3[(("foo bar" + v119) + (v119 + "a"))]
                                                                    _ = v191
                                                                default:
                                                                    var v192 uint16 = v154
                                                                    _ = v192
                                                                }
                                                            }
                                                            _ = v156
                                                        }
                                                        v2 += (v2 >> (((v155 << (((v107 / ((v107) | 1))) & 63))) & 63))
                                                        _ = v155
                                                    } else {
                                                        var v193 uint8 = (v117 * genFn1(false, v119))
                                                        _ = v193
                                                    }
                                                    _ = v154
                                                }
                                            }
                                        default:
                                            var v194 int64 = 1
                                            _ = v194
                                        }
                                    default:
                                        if ((v117 ^ (v117 + 2)) <= genFn1((0 > 0), v119)) {
                                            v117 += v117
                                            if ((71 >= (^(v117))) && ((v106 * v106) > uint32(v117))) {
                                                v195 := GenS23{F0: -32768, F1: 1}
                                                _ = v195
                                            } else {
                                                var v196 uint8 = (v117 - (v117 % (((^(v117))) | 1)))
                                                _ = v196
                                            }
                                        }
                                    }
                                    _ = v134
                                } else {
                                    v107 += 806
                                }
                                switch 4294967295 {
                                case 4294967295:
                                    if true {
                                        if (590 >= 919) {
                                            if true {
                                                v107 = (v107 & uint(v106))
                                                fmt.Println(v119)
                                            } else {
                                                v197 := GenS23{F0: -13, F1: 96}
                                                _ = len(v1)
                                                v117 = (v117 >> ((v107) & 63))
                                                _ = v197
                                            }
                                            v119 = v119
                                            v198 := GenS24{F0: 0}
                                            _ = v198
                                        }
                                        v199 := GenS25{F0: int64(v2), F1: (v106 & v106)}
                                        _ = v199
                                    } else {
                                        v119 = (((v119 + v119) + "") + v119)
                                    }
                                case 0:
                                    v3[((v119 + v119) + v119)] = 259
                                default:
                                    switch v119 {
                                    case "foo bar":
                                        v200 := GenS8{F0: false, F1: 42, F2: 1}
                                        _ = v200
                                    default:
                                        switch v119 {
                                        case "\n\t":
                                            v201 := GenS14{F0: (v119 + v119)}
                                            _ = v201
                                        case "hello":
                                            switch (v119 + (v119 + v119)) {
                                            case "\n\t":
                                                _ = len(v1)
                                            default:
                                                v107 -= v107
                                            }
                                        default:
                                            var v202 int8 = -37
                                            _ = v202
                                        }
                                    }
                                }
                                _ = len(v3)
                                var v203 int = 4294967295
                                _ = v119
                                _ = v133
                                _ = v203
                            }
                            _ = v117
                        } else {
                            var v204 int16 = 32767
                            fmt.Println(fi108)
                            _ = v204
                        }
                    case "a":
                        v2 = (v2 + v2)
                    }
                    _ = v116
                }
                switch ((("" + "a") + "foo bar") + "\n\t") {
                case "foo bar":
                    v205 := GenS12{F0: (v107 >> (((^(v107))) & 63)), F1: -54}
                    _ = v205
                }
                if true {
                    switch ("\n\t" + "foo bar") {
                    case "foo bar":
                        switch (fi108 ^ -9223372036854775808) {
                        case 32767:
                            v2 = v2
                        default:
                            var v206 int64 = 126
                            _ = v206
                        }
                    case "a":
                        v106 -= v106
                    }
                    var v207 int = 1
                    fmt.Println(v106)
                    _ = v207
                }
            }
            _ = fi109
        }
        _ = len(v3)
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi108
    }
    var v208 int = 0
    fmt.Println(v208, v107)
    v2 = (^(v2))
    fmt.Println(v106, v208)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v106:", v106)
    fmt.Println("v107:", v107)
    fmt.Println("v2:", v2)
    fmt.Println("v208:", v208)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v106
    _ = v107
    _ = v2
    _ = v208
    _ = v3
}
