package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint; F2 int16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint; F1 int8; F2 int8 }
type GenS5 struct { F0 uint8; F1 string; F2 bool }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 int8; F1 uint64; F2 int32 }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 int; F1 string; F2 uint32 }
type GenS10 struct { F0 int32; F1 int8 }
type GenS11 struct { F0 uint16; F1 int64 }
type GenS12 struct { F0 bool }
type GenS13 struct { F0 int32; F1 int16; F2 string }
type GenS14 struct { F0 int }
type GenS15 struct { F0 uint8; F1 uint; F2 bool }
type GenS16 struct { F0 int32; F1 uint64 }
type GenS17 struct { F0 int64; F1 int }
type GenS18 struct { F0 uint8; F1 int32 }
func genFn0(p0 uint32, p1 uint, p2 uint) uint {
    switch -2147483648 {
    case 128:
        p2 += (p2 % (((p2 % ((p1) | 1))) | 1))
    default:
        v5 := GenS1{F0: (p0 * p0), F1: (p1 & p1), F2: 32767}
        _ = v5
    }
    p1 = p1
    v6 := GenS1{F0: p0, F1: 904, F2: 0}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return p2
}
func genFn1(p0 string) int32 {
    var v7 string = "hello"
    p0 = (v7 + "foo bar")
    switch -2147483648 {
    case 4294967295:
        for fi8 := 0; fi8 < 2; fi8++ {
            v7 = "foo bar"
            var v9 float32 = -2.25
            if ((!false) || (1 != 0)) {
                var v10 uint32 = 42
                v10 -= 2147483648
                for fi11 := 0; fi11 < 11; fi11++ {
                    v7 = ("" + v7)
                    v12 := []int{}
                    _ = fi11
                    _ = v12
                }
                _ = v10
            } else {
                switch fi8 {
                case 9223372036854775807:
                    p0 = (("" + "\n\t") + (v7 + "a"))
                case -9223372036854775808:
                    fmt.Println(p0, v9)
                default:
                    p0 = "hello"
                }
                var v13 float32 = 0.0
                _ = v13
            }
            _ = fi8
            _ = v9
        }
    }
    p0 = (p0 + p0)
    _ = p0
    _ = v7
    return -200
}

func main() {
    var v1 int32 = -2147483648
    v2 := []uint32{1, 894}
    v3 := make(map[string]string)
    var v4 float32 = ((1e10 - (((0.1 - (((0.0 - 0.0) / (float32(1.5) * 3.14159)) * 3.14159)) / 0.1) + 1e10)) / float32(((3.14159 * (0.0 * ((0.1 / 3.14159) * 0.1))) - (1e10 - float64(3.14159)))))
    v4 += (v4 - 3.14159)
    for fi14 := 0; fi14 < 20; fi14++ {
        for fi15 := 0; fi15 < 13; fi15++ {
            if false {
                var v16 int64 = -9223372036854775808
                _ = v16
            } else {
                var v17 int64 = 619
                var v18 uint8 = 128
                _ = v17
                _ = v18
            }
            v4 = 0.1
            switch -1 {
            case 2147483647:
                if (837 != 1) {
                    v19 := GenS1{F0: 2147483648, F1: 1, F2: 32767}
                    _ = v19
                } else {
                    v1 = (v1 >> ((2) & 63))
                    var v20 string = ("foo bar" + "hello")
                    _ = v20
                }
            }
            _ = fi15
        }
        if false {
            var v21 uint16 = 866
            var v22 uint = 255
            var v23 int64 = -1
            _ = v21
            _ = v22
            _ = v23
        } else {
            v1 += genFn1("a")
            for fi24 := 0; fi24 < 17; fi24++ {
                var v25 int8 = 0
                v26 := GenS1{F0: 1, F1: 4294967295, F2: 0}
                switch ("foo bar" + "a") {
                case "\n\t":
                    v27 := v3["hello"]
                    _ = v27
                case "hello":
                    switch 9223372036854775807 {
                    case 42:
                        if len(v2) > 0 {
                            v2[0] = 1
                        }
                    }
                default:
                    delete(v3, "a")
                }
                _ = fi24
                _ = v25
                _ = v26
            }
            var v28 int = (-(fi14))
            _ = v28
        }
        _ = fi14
    }
    fmt.Println(v4, v1)
    if (-168 == 500) {
        if (true && ("" == "foo bar")) {
            for fi29 := 0; fi29 < 1; fi29++ {
                var v30 uint8 = 42
                _ = fi29
                _ = v30
            }
            v1 = (-(v1))
            v31 := GenS1{F0: 290, F1: 148, F2: -32768}
            _ = v31
        } else {
            if (-1 >= 0) {
                if false {
                    if true {
                        for fi32 := 0; fi32 < 19; fi32++ {
                            v33 := GenS1{F0: 154, F1: 0, F2: int16(fi32)}
                            _ = fi32
                            _ = v33
                        }
                        for fi34 := 0; fi34 < 12; fi34++ {
                            if (!false) {
                                if true {
                                    v35 := GenS2{F0: 154}
                                    v4 = (v4 - -2.25)
                                    _ = v35
                                }
                            }
                            for fi36 := 0; fi36 < 16; fi36++ {
                                if true {
                                    switch fi36 {
                                    case -2147483648:
                                        var v37 bool = (4294967295 == 953)
                                        _ = v37
                                    }
                                }
                                if false {
                                    var v38 int32 = -66
                                    _ = v38
                                }
                                switch (fi34 % (((fi34 & fi34)) | 1)) {
                                case -2147483648:
                                    switch 32767 {
                                    case 4294967295:
                                        v1 -= genFn1("a")
                                        if true {
                                            v4 -= v4
                                        } else {
                                            fmt.Println(fi36)
                                            if true {
                                                switch "foo bar" {
                                                case "foo bar":
                                                    if false {
                                                        var v39 int32 = v1
                                                        fmt.Println(v39)
                                                        if true {
                                                            var v40 int16 = -1
                                                            _ = v40
                                                        }
                                                        _ = v39
                                                    }
                                                }
                                                var v41 string = "a"
                                                switch (fi34 ^ fi36) {
                                                case 4294967295:
                                                    v41 = ((v41 + ("" + (v41 + v41))) + v41)
                                                case 127:
                                                    v41 = (v41 + v41)
                                                case 9223372036854775807:
                                                    var v42 string = "\n\t"
                                                    _ = v42
                                                default:
                                                    fmt.Println(fi36, fi34)
                                                }
                                                _ = v41
                                            } else {
                                                v43 := GenS2{F0: 3}
                                                if true {
                                                    v44 := GenS1{F0: 42, F1: 625, F2: 987}
                                                    var v45 uint = 4294967295
                                                    _ = v44
                                                    _ = v45
                                                }
                                                _ = v43
                                            }
                                            var v46 bool = false
                                            _ = v46
                                        }
                                        var v47 float64 = (((0.0 - 1.5) / 1.0) - 0.0)
                                        _ = v47
                                    case 1:
                                        v1 += v1
                                    }
                                    switch 32768 {
                                    case -976:
                                        if ("" != "") {
                                            switch fi36 {
                                            case 1:
                                                switch (^(fi36)) {
                                                case -2147483648:
                                                    fmt.Println(fi36, v4, v4)
                                                case 9223372036854775807:
                                                    if false {
                                                        if (!((-2.25 / 1e10) != 0.0)) {
                                                            if (42 == 100) {
                                                                v48 := GenS3{F0: "a"}
                                                                if true {
                                                                    v4 += v4
                                                                } else {
                                                                    if ((((v4 / -2.25) * v4) == v4) && true) {
                                                                        v1 = (v1 | -2147483648)
                                                                    }
                                                                    v1 -= -2147483648
                                                                    v4 -= (-2.25 * 1e10)
                                                                    var v49 uint64 = 4294967295
                                                                    v1 = v1
                                                                    _ = v49
                                                                }
                                                                _ = v48
                                                            }
                                                        }
                                                        if ((v1 & (v1 * v1)) < 42) {
                                                            if len(v2) > 0 {
                                                                v2[0] = 252
                                                            }
                                                        }
                                                    }
                                                default:
                                                    v50 := GenS2{F0: 1}
                                                    _ = v50
                                                }
                                            case 606:
                                                v51 := v3[((("\n\t" + "foo bar") + "foo bar") + "\n\t")]
                                                _ = v51
                                            case 283:
                                                v52 := GenS2{F0: 176}
                                                _ = v52
                                            }
                                            var v53 bool = (242 < 675)
                                            v53 = (42 >= -609)
                                            _ = v53
                                        } else {
                                            var v54 bool = true
                                            switch (fi34 * 128) {
                                            case 373:
                                                switch "a" {
                                                case "hello":
                                                    switch fi36 {
                                                    case 996:
                                                        switch fi34 {
                                                        case 127:
                                                            fmt.Println(v1, v1, v4)
                                                        case 0:
                                                            var v55 uint32 = 2
                                                            _ = v55
                                                        case 128:
                                                            _ = len(v2)
                                                        }
                                                    case 32768:
                                                        v56 := v3[""]
                                                        _ = v56
                                                    case 42:
                                                        if len(v2) >= 2 {
                                                            v2 = v2[1:len(v2)]
                                                        }
                                                    default:
                                                        if true {
                                                            if len(v2) > 0 {
                                                                v2[0] = 42
                                                            }
                                                            var v57 uint64 = 448
                                                            _ = v57
                                                        }
                                                    }
                                                }
                                                switch fi34 {
                                                case 1:
                                                    v54 = false
                                                case 32767:
                                                    v58 := GenS2{F0: 42}
                                                    _ = v58
                                                case 128:
                                                    delete(v3, "")
                                                default:
                                                    v1 = v1
                                                }
                                            }
                                            v59 := v3[""]
                                            _ = v54
                                            _ = v59
                                        }
                                    }
                                default:
                                    _ = len(v3)
                                }
                                _ = fi36
                            }
                            switch "\n\t" {
                            case "foo bar":
                                if true {
                                    switch ("hello" + "a") {
                                    case "foo bar":
                                        v1 += (v1 - (v1 - -1))
                                    case "hello":
                                        v1 += v1
                                    case "":
                                        fmt.Println(v1, v1)
                                    default:
                                        var v60 uint = 42
                                        _ = v60
                                    }
                                    var v61 float64 = -2.25
                                    _ = v61
                                }
                            }
                            _ = fi34
                        }
                    }
                    var v62 string = ""
                    v62 = (v62 + v62)
                    _ = v62
                }
                v63 := GenS2{F0: 128}
                _ = v63
            } else {
                var v64 string = "a"
                if false {
                    var v65 int32 = (v1 + (v1 % ((v1) | 1)))
                    v65 = v1
                    v65 -= v65
                    _ = v65
                } else {
                    var v66 uint64 = 4294967295
                    if true {
                        if ((!false) || true) {
                            var v67 string = "foo bar"
                            for fi68 := 0; fi68 < 5; fi68++ {
                                _ = len(v2)
                                if (100 >= -1) {
                                    for fi69 := 0; fi69 < 20; fi69++ {
                                        fmt.Println(fi68, v64, v1)
                                        _ = fi69
                                    }
                                    v70 := GenS4{F0: 990, F1: -128, F2: 0}
                                    var v71 int64 = 0
                                    _ = v70
                                    _ = v71
                                } else {
                                    var v72 int32 = genFn1((("hello" + v64) + ((v67 + v67) + (("a" + v64) + v64))))
                                    var v73 int = fi68
                                    v1 += v72
                                    _ = v72
                                    _ = v73
                                }
                                _ = fi68
                            }
                            _ = v67
                        } else {
                            v64 = ""
                            for fi74 := 0; fi74 < 16; fi74++ {
                                var v75 uint = 65535
                                _ = fi74
                                _ = v75
                            }
                            if ((128 > 112) && (42 < 9223372036854775807)) {
                                v76 := GenS1{F0: 2, F1: genFn0(2, 65535, 18446744073709551615), F2: 480}
                                var v77 uint32 = 42
                                for fi78 := 0; fi78 < 8; fi78++ {
                                    switch ("foo bar" + v64) {
                                    case "":
                                        fmt.Println(fi78)
                                    case "\n\t":
                                        switch v64 {
                                        case "":
                                            v64 = "\n\t"
                                        case "hello":
                                            for fi79 := 0; fi79 < 5; fi79++ {
                                                var v80 string = (v64 + v64)
                                                if true {
                                                    v77 += (v77 % ((v77) | 1))
                                                    if false {
                                                        var v81 int8 = 0
                                                        _ = len(v2)
                                                        var v82 int = (-(fi79))
                                                        var v83 string = v64
                                                        var v84 bool = false
                                                        _ = v81
                                                        _ = v82
                                                        _ = v83
                                                        _ = v84
                                                    } else {
                                                        if (42 > 1) {
                                                            var v85 string = "a"
                                                            fmt.Println(v66)
                                                            if (v4 <= v4) {
                                                                if true {
                                                                    if false {
                                                                        v64 = "foo bar"
                                                                        fmt.Println(fi78, v77, v85)
                                                                    } else {
                                                                        switch ("\n\t" + ((v80 + ((("foo bar" + v85) + v80) + "hello")) + (((("" + "a") + (v80 + "\n\t")) + v64) + ((((v80 + v85) + "") + "hello") + v80)))) {
                                                                        case "hello":
                                                                            v66 = (v66 ^ (^(v66)))
                                                                        default:
                                                                            fmt.Println(v80)
                                                                        }
                                                                        if (!(false && (false || false))) {
                                                                            if true {
                                                                                switch (fi79 % (((fi79 * fi79)) | 1)) {
                                                                                case -830:
                                                                                    v4 -= (0.0 / 1.5)
                                                                                case 128:
                                                                                    v77 += v77
                                                                                default:
                                                                                    var v86 float32 = v4
                                                                                    _ = v86
                                                                                }
                                                                                fmt.Println(fi78, fi78)
                                                                                if len(v2) > 0 {
                                                                                    v2[0] = (v77 % ((65535) | 1))
                                                                                }
                                                                                v87 := GenS5{F0: 127, F1: ("a" + "a"), F2: false}
                                                                                _ = v87
                                                                            } else {
                                                                                v88 := GenS4{F0: 4294967295, F1: 42, F2: -128}
                                                                                v89 := GenS2{F0: 2}
                                                                                _ = v88
                                                                                _ = v89
                                                                            }
                                                                            if len(v2) >= 2 {
                                                                                v2 = v2[1:len(v2)]
                                                                            }
                                                                            if ((0.0 * (1.5 + (3.14159 / (0.0 - 3.14159)))) > (0.1 + -2.25)) {
                                                                                if true {
                                                                                    v80 = v64
                                                                                    fmt.Println(v80)
                                                                                    v64 = "\n\t"
                                                                                    if ((v77 + v77) <= (v77 - 288)) {
                                                                                        fmt.Println(v85, v85, fi79)
                                                                                        var v90 int8 = -1
                                                                                        var v91 bool = false
                                                                                        _ = v90
                                                                                        _ = v91
                                                                                    }
                                                                                }
                                                                                v85 = "hello"
                                                                                var v92 uint8 = 58
                                                                                _ = len(v2)
                                                                                _ = v92
                                                                            } else {
                                                                                switch fi79 {
                                                                                case 0:
                                                                                    switch v85 {
                                                                                    case "foo bar":
                                                                                        v93 := GenS6{F0: 0}
                                                                                        _ = v93
                                                                                    case "hello":
                                                                                        _ = len(v3)
                                                                                    }
                                                                                default:
                                                                                    fmt.Println(v1)
                                                                                }
                                                                                v94 := GenS7{F0: 95, F1: 993, F2: v1}
                                                                                var v95 int32 = genFn1("foo bar")
                                                                                _ = v94
                                                                                _ = v95
                                                                            }
                                                                        } else {
                                                                            var v96 uint8 = 42
                                                                            if ((v96 == (v96 | 255)) && false) {
                                                                                v77 = 42
                                                                                switch fi78 {
                                                                                case 127:
                                                                                    fmt.Println(fi79)
                                                                                    if true {
                                                                                        v77 = 65535
                                                                                        fmt.Println(v66)
                                                                                        var v97 int = -2147483648
                                                                                        _ = v97
                                                                                    }
                                                                                default:
                                                                                    var v98 int32 = v1
                                                                                    _ = v98
                                                                                }
                                                                            }
                                                                            v4 = (3.14159 - v4)
                                                                            _ = v96
                                                                        }
                                                                        if true {
                                                                            v4 += v4
                                                                            var v99 string = (v80 + v85)
                                                                            if ("hello" == (("hello" + v99) + (("" + "foo bar") + (("a" + v80) + v80)))) {
                                                                                v66 -= v66
                                                                                var v100 int8 = 1
                                                                                fmt.Println(fi79)
                                                                                _ = v100
                                                                            } else {
                                                                                var v101 uint64 = v66
                                                                                var v102 int32 = v1
                                                                                _ = v101
                                                                                _ = v102
                                                                            }
                                                                            _ = v99
                                                                        }
                                                                    }
                                                                    v4 += 3.14159
                                                                }
                                                                if true {
                                                                    v103 := GenS8{F0: -32768}
                                                                    var v104 int64 = -641
                                                                    if (592 != uint16(v66)) {
                                                                        fmt.Println(v104)
                                                                        fmt.Println(v1)
                                                                        delete(v3, (v64 + v64))
                                                                    }
                                                                    _ = v103
                                                                    _ = v104
                                                                } else {
                                                                    v3["a"] = (v64 + v64)
                                                                }
                                                            }
                                                            _ = v85
                                                        }
                                                        v4 -= 1.5
                                                        v77 = (v77 ^ v77)
                                                    }
                                                } else {
                                                    if (genFn0((^(v77)), 2, 42) < genFn0((v77 / ((328) | 1)), 1, 2)) {
                                                        v64 = v64
                                                        v105 := GenS9{F0: (fi78 & fi79), F1: v80, F2: v77}
                                                        _ = len(v3)
                                                        _ = v105
                                                    } else {
                                                        var v106 float32 = 1.5
                                                        _ = v106
                                                    }
                                                    if true {
                                                        v107 := GenS5{F0: 128, F1: "a", F2: false}
                                                        if true {
                                                            if false {
                                                                v108 := GenS10{F0: (v1 + (v1 << ((uint(v66)) & 63))), F1: 1}
                                                                var v109 int8 = -93
                                                                switch (v64 + v64) {
                                                                case "\n\t":
                                                                    if (uint(v77) != 1) {
                                                                        v77 = uint32(v66)
                                                                        fmt.Println(v1)
                                                                        v80 = v80
                                                                    }
                                                                case "hello":
                                                                    var v110 uint = 18446744073709551615
                                                                    _ = v110
                                                                }
                                                                _ = v108
                                                                _ = v109
                                                            }
                                                            fmt.Println(v64)
                                                            if false {
                                                                switch v64 {
                                                                case "foo bar":
                                                                    _ = len(v2)
                                                                case "":
                                                                    switch (^(fi78)) {
                                                                    case -9223372036854775808:
                                                                        _ = len(v2)
                                                                    case 4294967295:
                                                                        v77 += v77
                                                                    }
                                                                    _ = len(v3)
                                                                    if false {
                                                                        v111 := GenS4{F0: 65535, F1: 1, F2: -1}
                                                                        _ = v111
                                                                    } else {
                                                                        var v112 bool = ((v64 + (((v64 + (v64 + v64)) + "\n\t") + ("\n\t" + "hello"))) > (v64 + "hello"))
                                                                        v4 -= (v4 + (v4 - (3.14159 * 3.14159)))
                                                                        switch (v80 + v64) {
                                                                        case "\n\t":
                                                                            v64 = v80
                                                                        case "foo bar":
                                                                            var v113 float64 = (1.5 + float64(v4))
                                                                            _ = v113
                                                                        default:
                                                                            var v114 float32 = v4
                                                                            _ = v114
                                                                        }
                                                                        if (fi79 == fi78) {
                                                                            if (true && v112) {
                                                                                v115 := GenS11{F0: 1, F1: -729}
                                                                                fmt.Println(v77, v64, v112)
                                                                                _ = v115
                                                                            }
                                                                            v116 := v3[v80]
                                                                            _ = v116
                                                                        } else {
                                                                            _ = len(v2)
                                                                            v66 -= (v66 << ((genFn0((v77 | 65535), genFn0(v77, 255, 4294967295), 403)) & 63))
                                                                            var v117 float32 = v4
                                                                            v80 = v64
                                                                            if (101 < 22) {
                                                                                if len(v2) > 0 {
                                                                                    v2[0] = 65535
                                                                                }
                                                                                if ((fi79 / ((fi79) | 1)) < (fi78 ^ fi79)) {
                                                                                    if (42 <= 0) {
                                                                                        if len(v2) >= 2 {
                                                                                            v2 = v2[1:len(v2)]
                                                                                        }
                                                                                        switch -2147483648 {
                                                                                        case 2147483647:
                                                                                            v118 := GenS1{F0: v77, F1: genFn0(v77, 42, uint(v77)), F2: 1}
                                                                                            _ = v118
                                                                                        case -31:
                                                                                            v119 := GenS3{F0: (v64 + v80)}
                                                                                            v80 = v80
                                                                                            var v120 bool = (1 >= -712)
                                                                                            _ = v119
                                                                                            _ = v120
                                                                                        }
                                                                                        v1 = (v1 | genFn1(("hello" + (("a" + v80) + v80))))
                                                                                    } else {
                                                                                        var v121 uint64 = v66
                                                                                        fmt.Println(v117, fi79)
                                                                                        _ = v121
                                                                                    }
                                                                                    v122 := v3[v80]
                                                                                    _ = v122
                                                                                } else {
                                                                                    var v123 int = (fi78 ^ (fi78 * (fi79 % (((^(fi79))) | 1))))
                                                                                    v64 = "\n\t"
                                                                                    _ = v123
                                                                                }
                                                                                v4 = v117
                                                                            }
                                                                            _ = v117
                                                                        }
                                                                        _ = v112
                                                                    }
                                                                case "\n\t":
                                                                    var v124 float64 = -2.25
                                                                    _ = v124
                                                                default:
                                                                    if false {
                                                                        var v125 float64 = 0.0
                                                                        v64 = v80
                                                                        _ = v125
                                                                    }
                                                                }
                                                                if false {
                                                                    switch v80 {
                                                                    case "a":
                                                                        v126 := GenS5{F0: 127, F1: "foo bar", F2: false}
                                                                        v2 = append(v2, v77)
                                                                        _ = len(v2)
                                                                        _ = v126
                                                                    default:
                                                                        var v127 uint32 = 1
                                                                        _ = v127
                                                                    }
                                                                    switch 4294967295 {
                                                                    case 1:
                                                                        _ = len(v2)
                                                                    default:
                                                                        v66 -= v66
                                                                    }
                                                                } else {
                                                                    switch (fi78 << ((genFn0((v77 | 202), 255, 42)) & 63)) {
                                                                    case 9223372036854775807:
                                                                        v80 = "\n\t"
                                                                        if true {
                                                                            if len(v2) >= 2 {
                                                                                v2 = v2[1:len(v2)]
                                                                            }
                                                                        } else {
                                                                            v1 -= v1
                                                                        }
                                                                        var v128 bool = true
                                                                        _ = v128
                                                                    case 127:
                                                                        var v129 uint = 2
                                                                        _ = v129
                                                                    }
                                                                    var v130 int = fi78
                                                                    if false {
                                                                        if false {
                                                                            if false {
                                                                                var v131 bool = true
                                                                                if v131 {
                                                                                    v66 += 689
                                                                                    fmt.Println(v77, v77)
                                                                                } else {
                                                                                    v132 := GenS6{F0: 0}
                                                                                    v131 = (!v131)
                                                                                    v133 := GenS7{F0: 0, F1: (v66 + (v66 << ((genFn0(4294967295, 2, 42)) & 63))), F2: v1}
                                                                                    v66 += v66
                                                                                    _ = len(v2)
                                                                                    _ = v132
                                                                                    _ = v133
                                                                                }
                                                                                v3[(("a" + ("foo bar" + ((v64 + "\n\t") + v64))) + ("" + ((v80 + ("a" + v64)) + v80)))] = ""
                                                                                _ = v131
                                                                            } else {
                                                                                v134 := GenS3{F0: v80}
                                                                                v130 += fi79
                                                                                if len(v2) >= 2 {
                                                                                    v2 = v2[1:len(v2)]
                                                                                }
                                                                                _ = v134
                                                                            }
                                                                            fmt.Println(v80, fi78, fi78)
                                                                            if (-1 >= 32767) {
                                                                                v4 += float32(0.0)
                                                                                v135 := GenS12{F0: ((false && (v66 == v66)) || false)}
                                                                                _ = v135
                                                                            }
                                                                            if (v77 >= v77) {
                                                                                v136 := GenS13{F0: v1, F1: int16(v77), F2: v64}
                                                                                _ = v136
                                                                            }
                                                                            v137 := GenS1{F0: (^(v77)), F1: 731, F2: -32768}
                                                                            _ = v137
                                                                        } else {
                                                                            var v138 string = v80
                                                                            _ = v138
                                                                        }
                                                                        fmt.Println(v4, v1, v130)
                                                                    }
                                                                    _ = v130
                                                                }
                                                            } else {
                                                                var v139 int64 = 0
                                                                v140 := GenS3{F0: (((v64 + v80) + v64) + v64)}
                                                                _ = v139
                                                                _ = v140
                                                            }
                                                        } else {
                                                            v141 := GenS7{F0: -37, F1: v66, F2: v1}
                                                            switch "foo bar" {
                                                            case "hello":
                                                                switch fi79 {
                                                                case 427:
                                                                    switch v64 {
                                                                    case "foo bar":
                                                                        v142 := GenS2{F0: uint8(v77)}
                                                                        _ = v142
                                                                    case "hello":
                                                                        switch (fi79 % ((fi78) | 1)) {
                                                                        case 4294967295:
                                                                            v64 = ("hello" + "a")
                                                                        case 1:
                                                                            v143 := GenS14{F0: -2147483648}
                                                                            _ = v143
                                                                        case 2147483647:
                                                                            var v144 float64 = ((-2.25 + (float64(v4) / -2.25)) - (3.14159 * -2.25))
                                                                            _ = v144
                                                                        default:
                                                                            if (42 <= 565) {
                                                                                fmt.Println(fi78)
                                                                                var v145 int = (fi78 % ((162) | 1))
                                                                                v146 := GenS15{F0: 128, F1: 65535, F2: true}
                                                                                _ = v145
                                                                                _ = v146
                                                                            }
                                                                        }
                                                                    }
                                                                default:
                                                                    v77 = (v77 % ((v77) | 1))
                                                                }
                                                            default:
                                                                var v147 float32 = v4
                                                                _ = v147
                                                            }
                                                            _ = v141
                                                        }
                                                        switch (fi78 % ((0) | 1)) {
                                                        case 42:
                                                            v148 := GenS16{F0: (v1 - (v1 * (v1 ^ v1))), F1: 386}
                                                            _ = v148
                                                        default:
                                                            v2 = append(v2, v77)
                                                        }
                                                        v1 = v1
                                                        var v149 int16 = 1
                                                        v77 += v77
                                                        if (!(false && true)) {
                                                            if false {
                                                                if true {
                                                                    v150 := GenS3{F0: ("" + "a")}
                                                                    var v151 int32 = (v1 << ((genFn0((v77 - 735), 0, genFn0(1, 255, 65535))) & 63))
                                                                    _ = v150
                                                                    _ = v151
                                                                }
                                                                var v152 uint32 = v77
                                                                _ = v152
                                                            }
                                                        } else {
                                                            if (false || (-1 > (^(v149)))) {
                                                                v153 := GenS17{F0: -1, F1: fi78}
                                                                v77 += v77
                                                                v64 = v64
                                                                switch v80 {
                                                                case "\n\t":
                                                                    if false {
                                                                        v3[((v64 + ((v64 + ("a" + v80)) + "")) + ("a" + "hello"))] = ((v64 + v64) + ("" + ""))
                                                                        v80 = "hello"
                                                                        var v154 uint64 = 249
                                                                        var v155 int32 = -711
                                                                        v156 := GenS4{F0: 954, F1: 42, F2: 0}
                                                                        _ = v154
                                                                        _ = v155
                                                                        _ = v156
                                                                    } else {
                                                                        var v157 uint32 = (v77 | v77)
                                                                        _ = v157
                                                                    }
                                                                }
                                                                _ = v153
                                                            }
                                                            v158 := GenS6{F0: int16(v77)}
                                                            fmt.Println(v64, v80, v66)
                                                            _ = v158
                                                        }
                                                        v149 -= v149
                                                        v64 = v64
                                                        var v159 uint8 = 42
                                                        switch "foo bar" {
                                                        case "foo bar":
                                                            _ = len(v3)
                                                        }
                                                        _ = v107
                                                        _ = v149
                                                        _ = v159
                                                    }
                                                }
                                                var v160 string = ("\n\t" + "\n\t")
                                                _ = fi79
                                                _ = v160
                                                _ = v80
                                            }
                                        case "foo bar":
                                            for fi161 := 0; fi161 < 13; fi161++ {
                                                v64 = v64
                                                var v162 uint8 = 255
                                                var v163 int64 = 1
                                                _ = fi161
                                                _ = v162
                                                _ = v163
                                            }
                                        }
                                    case "foo bar":
                                        v164 := GenS14{F0: fi78}
                                        _ = v164
                                    }
                                    if true {
                                        v165 := GenS2{F0: 127}
                                        _ = v165
                                    }
                                    _ = fi78
                                }
                                _ = v76
                                _ = v77
                            } else {
                                v66 -= 1
                                if true {
                                    v166 := GenS11{F0: 2, F1: 1}
                                    var v167 int32 = (-(v1))
                                    var v168 int16 = -562
                                    _ = v166
                                    _ = v167
                                    _ = v168
                                }
                                for fi169 := 0; fi169 < 6; fi169++ {
                                    for fi170 := 0; fi170 < 19; fi170++ {
                                        if true {
                                            var v171 int8 = 1
                                            v64 = ""
                                            var v172 int32 = (v1 & (v1 + v1))
                                            _ = v171
                                            _ = v172
                                        } else {
                                            v4 = 0.0
                                            if true {
                                                var v173 string = (v64 + v64)
                                                if ((v66 - v66) != v66) {
                                                    _ = len(v3)
                                                    v66 -= (v66 ^ 0)
                                                }
                                                _ = v173
                                            } else {
                                                _ = len(v3)
                                            }
                                            v174 := GenS18{F0: 1, F1: (v1 * 1)}
                                            _ = v174
                                        }
                                        _ = fi170
                                    }
                                    _ = fi169
                                }
                            }
                        }
                        v1 = v1
                    }
                    _ = v66
                }
                _ = v64
            }
        }
        v4 += v4
    } else {
        var v175 float32 = 1.5
        v176 := GenS15{F0: 2, F1: 196, F2: false}
        if false {
            v177 := GenS4{F0: 65535, F1: -128, F2: 0}
            switch (("" + ("a" + "a")) + ("a" + "\n\t")) {
            case "a":
                fmt.Println(v1, v1)
            case "\n\t":
                var v178 int32 = v1
                _ = v178
            }
            v1 += v1
            _ = v177
        }
        _ = v175
        _ = v176
    }
    v1 = v1
    fmt.Println(v1)
    var v179 uint16 = 42
    var v180 uint = 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("v179:", v179)
    fmt.Println("v180:", v180)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v179
    _ = v180
    _ = v2
    _ = v3
    _ = v4
}
