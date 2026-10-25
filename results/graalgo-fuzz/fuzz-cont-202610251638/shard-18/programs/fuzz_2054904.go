package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int; F1 uint64; F2 uint64 }
type GenS5 struct { F0 string; F1 uint }
type GenS6 struct { F0 string }
type GenS7 struct { F0 string; F1 bool }
type GenS8 struct { F0 uint; F1 int; F2 int32 }
type GenS9 struct { F0 int; F1 uint; F2 uint64 }
type GenS10 struct { F0 uint }
type GenS11 struct { F0 uint }
type GenS12 struct { F0 int; F1 int8 }
type GenS13 struct { F0 int64; F1 bool }
type GenS14 struct { F0 int32; F1 uint64; F2 string }
func genFn0(p0 string) string {
    switch 42 {
    case 127:
        p0 = p0
    default:
        if true {
            switch p0 {
            case "hello":
                p0 = ("a" + p0)
            case "\n\t":
                v7 := []int{1, -102, 999}
                _ = v7
            case "":
                var v8 uint = 42
                _ = v8
            default:
                fmt.Println(p0)
            }
            v9 := GenS1{F0: false}
            _ = v9
        } else {
            var v10 bool = false
            v10 = (p0 > "")
            p0 = ((p0 + p0) + p0)
            _ = v10
        }
    }
    v11 := []int{42, 0}
    if false {
        if len(v11) >= 2 {
            v11 = v11[1:len(v11)]
        }
        p0 = (p0 + (p0 + p0))
        for fi12 := 0; fi12 < 3; fi12++ {
            for fi13 := 0; fi13 < 20; fi13++ {
                v14 := GenS1{F0: ((!((32767 <= 32767) || (2147483647 > int32(fi12)))) || false)}
                _ = fi13
                _ = v14
            }
            _ = fi12
        }
    }
    p0 = ("\n\t" + "a")
    _ = p0
    _ = v11
    return (p0 + "a")
}
func genFn1(p0 int16, p1 int) uint {
    if true {
        if false {
            p0 = 42
            for fi15 := 0; fi15 < 6; fi15++ {
                fmt.Println(p0, p1, fi15)
                fmt.Println(p1)
                _ = fi15
            }
        } else {
            var v16 bool = ((42 <= 127) || false)
            p1 -= (p1 >> ((0) & 63))
            v17 := []int{(^(p1))}
            _ = v16
            _ = v17
        }
        if ((-728 <= -2147483648) || (!false)) {
            p1 += (p1 | p1)
            v18 := GenS2{F0: ((127 <= 0) || (4294967295 <= 4294967295))}
            _ = v18
        }
        fmt.Println(p1, p0)
    } else {
        v19 := GenS1{F0: (1 < 795)}
        if false {
            switch "" {
            case "hello":
                p1 = p1
            case "a":
                fmt.Println(p0)
            case "\n\t":
                p1 += (p1 ^ (p1 >> ((260) & 63)))
            default:
                var v20 int8 = 1
                _ = v20
            }
        }
        _ = v19
    }
    if ((!(775 > 857)) || (true && false)) {
        v21 := []int{4294967295, p1, p1, 724}
        if len(v21) >= 2 {
            v21 = v21[1:len(v21)]
        }
        for fi22 := 0; fi22 < 20; fi22++ {
            if len(v21) >= 2 {
                v21 = v21[1:len(v21)]
            }
            var v23 float32 = 0.1
            v23 -= (3.14159 + v23)
            _ = fi22
            _ = v23
        }
        _ = v21
    }
    v24 := []int{(p1 - (-(p1))), p1, (p1 + (^(p1))), 0}
    _ = p0
    _ = p1
    _ = v24
    return 0
}

func main() {
    var v1 int64 = -9223372036854775808
    var v2 bool = false
    v3 := make(map[int]bool)
    v4 := make(map[uint32]string)
    var v5 int32 = -2147483648
    v6 := []uint32{327}
    switch 4294967295 {
    case 32767:
        switch 32768 {
        case 904:
            var v25 uint32 = 2
            _ = v25
        case -962:
            _ = len(v3)
        default:
            for fi26 := 0; fi26 < 14; fi26++ {
                fmt.Println(v2, v1)
                _ = fi26
            }
        }
    case 32768:
        switch "a" {
        case "\n\t":
            for fi27 := 0; fi27 < 5; fi27++ {
                fmt.Println(v5)
                fmt.Println(v5)
                var v28 uint32 = 2147483647
                _ = fi27
                _ = v28
            }
        case "hello":
            var v29 int8 = 1
            _ = v29
        case "":
            for fi30 := 0; fi30 < 16; fi30++ {
                v1 = v1
                _ = fi30
            }
        }
    }
    var v31 float32 = ((3.14159 / 1.0) + float32(((-2.25 - -2.25) * (1.5 - ((-2.25 * ((1.5 * (((0.0 / ((0.1 * 1e10) - 1e10)) - 1.5) - 0.1)) + 3.14159)) / 1.0)))))
    v1 -= v1
    for fi32 := 0; fi32 < 4; fi32++ {
        v5 = v5
        _ = fi32
    }
    if true {
        var v33 string = "foo bar"
        fmt.Println(v2, v2)
        switch 786 {
        case -9223372036854775808:
            v34 := GenS2{F0: (true && (false && (65535 <= 2)))}
            _ = v34
        case -2147483648:
            switch v33 {
            case "a":
                for fi35 := 0; fi35 < 13; fi35++ {
                    if (!true) {
                        if (((float32(1e10) + (v31 / (float32(((1.5 / (0.0 + float64(v31))) - 1.5)) / v31))) <= float32(1.5)) || (255 < 65535)) {
                            var v36 uint16 = 797
                            for fi37 := 0; fi37 < 4; fi37++ {
                                switch 42 {
                                case 9223372036854775807:
                                    var v38 uint32 = 2
                                    _ = v38
                                case 42:
                                    var v39 uint32 = uint32(v36)
                                    _ = v39
                                case 32767:
                                    v5 += v5
                                default:
                                    var v40 uint16 = (^(v36))
                                    _ = v40
                                }
                                fmt.Println(v31)
                                if true {
                                    v41 := GenS2{F0: true}
                                    _ = v41
                                }
                                var v42 uint64 = 1
                                var v43 int8 = -99
                                var v44 int = 128
                                if false {
                                    v45 := GenS1{F0: v2}
                                    if false {
                                        fmt.Println(v1, fi35, v42)
                                    } else {
                                        switch v33 {
                                        case "hello":
                                            var v46 float64 = 1.5
                                            var v47 uint8 = 128
                                            _ = v46
                                            _ = v47
                                        case "":
                                            v48 := v3[(fi37 - -684)]
                                            _ = v48
                                        case "a":
                                            if v2 {
                                                var v49 float64 = (((0.0 * (3.14159 - ((3.14159 - 1e10) * 1e10))) - 1.5) / 0.1)
                                                switch -9223372036854775808 {
                                                case 2147483647:
                                                    if false {
                                                        var v50 int32 = v5
                                                        var v51 int16 = -32768
                                                        _ = v50
                                                        _ = v51
                                                    }
                                                case -9223372036854775808:
                                                    delete(v3, (v44 >> ((2) & 63)))
                                                default:
                                                    var v52 uint32 = 65535
                                                    _ = v52
                                                }
                                                _ = v49
                                            }
                                        default:
                                            var v53 int64 = (^(v1))
                                            _ = v53
                                        }
                                    }
                                    var v54 bool = true
                                    v55 := GenS2{F0: (v1 <= v1)}
                                    _ = v45
                                    _ = v54
                                    _ = v55
                                } else {
                                    fmt.Println(v33, v42, v1)
                                    fmt.Println(v43)
                                    if len(v6) >= 2 {
                                        v6 = v6[1:len(v6)]
                                    }
                                }
                                _ = fi37
                                _ = v42
                                _ = v43
                                _ = v44
                            }
                            _ = v36
                        } else {
                            fmt.Println(fi35)
                            for fi56 := 0; fi56 < 8; fi56++ {
                                if v2 {
                                    v33 = "a"
                                    if v2 {
                                        v57 := GenS3{F0: 351}
                                        var v58 uint = 255
                                        fmt.Println(v58, fi35, v5)
                                        switch (fi56 % ((fi35) | 1)) {
                                        case -2147483648:
                                            v33 = ((("\n\t" + ("foo bar" + ("\n\t" + v33))) + ("a" + "")) + ("foo bar" + ((v33 + ((((v33 + (v33 + v33)) + v33) + (((("foo bar" + v33) + v33) + v33) + (v33 + v33))) + v33)) + (v33 + v33))))
                                        default:
                                            v2 = (!false)
                                        }
                                        _ = v57
                                        _ = v58
                                    }
                                    switch fi35 {
                                    case 2147483647:
                                        var v59 float64 = ((-2.25 * (1.5 * 1e10)) * 1e10)
                                        _ = v59
                                    default:
                                        v2 = (false && (984 == 297))
                                    }
                                } else {
                                    v31 = v31
                                }
                                v60 := GenS4{F0: 0, F1: 42, F2: 1}
                                if (42 != 0) {
                                    var v61 uint32 = 512
                                    _ = v61
                                } else {
                                    switch v33 {
                                    case "":
                                        if (fi56 < fi56) {
                                            if (v2 || (!true)) {
                                                fmt.Println(fi35)
                                            } else {
                                                v5 -= (v5 << ((1) & 63))
                                                switch "a" {
                                                case "\n\t":
                                                    fmt.Println(v1, fi56)
                                                }
                                            }
                                            var v62 bool = (1 != 18446744073709551615)
                                            _ = v62
                                        }
                                    case "hello":
                                        fmt.Println(v5, v1, v31)
                                    case "foo bar":
                                        var v63 int16 = 494
                                        _ = v63
                                    default:
                                        v5 += v5
                                    }
                                    v64 := GenS1{F0: v2}
                                    if false {
                                        _ = len(v6)
                                        v33 = (v33 + v33)
                                        var v65 string = v33
                                        if (("\n\t" == ((v65 + "\n\t") + ((("foo bar" + "\n\t") + "hello") + ("\n\t" + "")))) || false) {
                                            var v66 int64 = (v1 + (^(v1)))
                                            v3[(fi56 | (fi56 | 9223372036854775807))] = (-213 >= 42)
                                            v33 = v65
                                            _ = v66
                                        } else {
                                            fmt.Println(fi56, v2, fi35)
                                            switch v33 {
                                            case "":
                                                if (128 == 255) {
                                                    v67 := GenS2{F0: true}
                                                    switch (v65 + (v33 + ((((v65 + "hello") + ("" + "hello")) + v33) + "a"))) {
                                                    case "foo bar":
                                                        _ = len(v6)
                                                    case "":
                                                        switch (fi56 % ((32768) | 1)) {
                                                        case 9223372036854775807:
                                                            v68 := GenS5{F0: "", F1: 71}
                                                            _ = v68
                                                        case -2147483648:
                                                            v33 = ""
                                                        case -899:
                                                            var v69 int32 = (v5 >> ((42) & 63))
                                                            _ = v69
                                                        default:
                                                            v33 = (v65 + v65)
                                                        }
                                                    case "hello":
                                                        v1 = v1
                                                    default:
                                                        var v70 string = "hello"
                                                        _ = v70
                                                    }
                                                    _ = v67
                                                } else {
                                                    switch 978 {
                                                    case 554:
                                                        var v71 uint8 = 0
                                                        _ = v71
                                                    }
                                                    v1 -= (v1 + (v1 << ((4294967295) & 63)))
                                                }
                                            case "foo bar":
                                                if v2 {
                                                    v6 = append(v6, 65535)
                                                    v31 -= 0.1
                                                    var v72 uint8 = 0
                                                    _ = v72
                                                } else {
                                                    v1 = v1
                                                    if (42 < 366) {
                                                        v65 = v33
                                                    } else {
                                                        delete(v3, -1)
                                                        var v73 int64 = int64(fi56)
                                                        if true {
                                                            fmt.Println(v73, v1)
                                                        }
                                                        _ = v73
                                                    }
                                                    switch (v65 + v65) {
                                                    case "hello":
                                                        fmt.Println(fi56, fi56, v31)
                                                    case "\n\t":
                                                        fmt.Println(v2, v1)
                                                    case "a":
                                                        var v74 int16 = -515
                                                        _ = v74
                                                    default:
                                                        var v75 uint64 = 608
                                                        _ = v75
                                                    }
                                                }
                                            }
                                            var v76 uint8 = 2
                                            _ = len(v3)
                                            _ = len(v6)
                                            var v77 int = -9223372036854775808
                                            _ = v76
                                            _ = v77
                                        }
                                        _ = v65
                                    } else {
                                        switch (^(fi56)) {
                                        case 793:
                                            v78 := GenS2{F0: v2}
                                            _ = v78
                                        case -1:
                                            switch int(v5) {
                                            case -1:
                                                switch (-(fi35)) {
                                                case 128:
                                                    if v2 {
                                                        if false {
                                                            v5 = v5
                                                        }
                                                        v79 := GenS6{F0: v33}
                                                        if v2 {
                                                            var v80 string = v33
                                                            var v81 uint = 18446744073709551615
                                                            if true {
                                                                if (v2 || (v2 || (42 >= (^(v1))))) {
                                                                    fmt.Println(v80, fi56, fi56)
                                                                    fmt.Println(v5, v33, fi56)
                                                                    if false {
                                                                        if len(v6) >= 2 {
                                                                            v6 = v6[1:len(v6)]
                                                                        }
                                                                    } else {
                                                                        v1 = int64(fi56)
                                                                        var v82 int16 = 32767
                                                                        _ = v82
                                                                    }
                                                                } else {
                                                                    v81 = v81
                                                                    if (-738 == (fi35 ^ -192)) {
                                                                        v31 += v31
                                                                    } else {
                                                                        v83 := GenS1{F0: true}
                                                                        var v84 float32 = 0.1
                                                                        fmt.Println(v1, v80, v5)
                                                                        _ = v83
                                                                        _ = v84
                                                                    }
                                                                    var v85 int64 = v1
                                                                    _ = v85
                                                                }
                                                                fmt.Println(v33, v2, v80)
                                                            }
                                                            _ = v80
                                                            _ = v81
                                                        } else {
                                                            _ = len(v3)
                                                            var v86 int32 = v5
                                                            _ = len(v4)
                                                            _ = v86
                                                        }
                                                        if ((false || true) && true) {
                                                            var v87 int16 = 402
                                                            _ = v87
                                                        }
                                                        if false {
                                                            _ = len(v3)
                                                        } else {
                                                            v88 := GenS7{F0: ((("" + "hello") + ((v33 + "foo bar") + "\n\t")) + (v33 + "\n\t")), F1: (((v1 - v1) == -9223372036854775808) || (18446744073709551615 >= 255))}
                                                            _ = len(v6)
                                                            v89 := GenS4{F0: fi56, F1: 127, F2: 42}
                                                            v90 := GenS4{F0: (fi56 ^ (fi35 ^ 127)), F1: 435, F2: 2}
                                                            _ = v88
                                                            _ = v89
                                                            _ = v90
                                                        }
                                                        _ = v79
                                                    } else {
                                                        var v91 string = "hello"
                                                        _ = v91
                                                    }
                                                }
                                            case 127:
                                                delete(v4, 65535)
                                            case 128:
                                                var v92 uint8 = 123
                                                v93 := GenS7{F0: "", F1: (v2 && true)}
                                                fmt.Println(v33, v1)
                                                _ = v92
                                                _ = v93
                                            }
                                        case 127:
                                            var v94 float64 = 1.5
                                            _ = v94
                                        }
                                    }
                                    _ = len(v6)
                                    v1 += v1
                                    _ = v64
                                }
                                if v2 {
                                    fmt.Println(fi35)
                                    if v2 {
                                        switch ("foo bar" + "hello") {
                                        case "\n\t":
                                            v31 = v31
                                        case "a":
                                            _ = len(v3)
                                        case "foo bar":
                                            if len(v6) >= 2 {
                                                v6 = v6[1:len(v6)]
                                            }
                                        default:
                                            v95 := GenS2{F0: (!v2)}
                                            _ = v95
                                        }
                                    }
                                }
                                v96 := GenS3{F0: 65535}
                                _ = fi56
                                _ = v60
                                _ = v96
                            }
                            _ = len(v6)
                        }
                        _ = len(v6)
                    }
                    v97 := GenS8{F0: 613, F1: 947, F2: (v5 / (((v5 & v5)) | 1))}
                    _ = fi35
                    _ = v97
                }
            default:
                v5 -= 2147483647
            }
        }
        _ = v33
    } else {
        if (!v2) {
            if len(v6) > 0 {
                v6[0] = 42
            }
            for fi98 := 0; fi98 < 9; fi98++ {
                v99 := GenS4{F0: fi98, F1: 229, F2: 568}
                _ = fi98
                _ = v99
            }
            fmt.Println(v31, v5)
        }
        var v100 uint = 1
        v101 := v4[1]
        _ = v100
        _ = v101
    }
    var v102 bool = v2
    switch -253 {
    case 2147483647:
        if false {
            v102 = ((true && (true || v102)) && false)
            if true {
                for fi103 := 0; fi103 < 4; fi103++ {
                    switch "foo bar" {
                    case "\n\t":
                        for fi104 := 0; fi104 < 4; fi104++ {
                            var v105 int8 = 65
                            v5 -= v5
                            delete(v4, 2147483648)
                            _ = fi104
                            _ = v105
                        }
                    case "":
                        delete(v4, 951)
                    case "hello":
                        if (1 < v5) {
                            for fi106 := 0; fi106 < 4; fi106++ {
                                var v107 int32 = (v5 | v5)
                                var v108 int64 = -1
                                if false {
                                    if true {
                                        v31 = 1.5
                                        v5 = (v107 << ((4294967295) & 63))
                                    } else {
                                        var v109 uint32 = 0
                                        switch ("hello" + "\n\t") {
                                        case "":
                                            v6 = append(v6, 47)
                                        }
                                        v31 -= 1.5
                                        _ = v109
                                    }
                                    fmt.Println(fi103)
                                    if (18446744073709551615 == 0) {
                                        var v110 int = fi106
                                        if (1.5 < 1.5) {
                                            switch 9223372036854775807 {
                                            case 0:
                                                v111 := GenS9{F0: 0, F1: 18446744073709551615, F2: 4294967295}
                                                v2 = v102
                                                v2 = ((100 < 42) || (!v102))
                                                _ = v111
                                            case -9223372036854775808:
                                                var v112 uint32 = 414
                                                _ = v112
                                            case 882:
                                                var v113 int16 = int16(v107)
                                                _ = v113
                                            default:
                                                v114 := GenS2{F0: v102}
                                                _ = v114
                                            }
                                            fmt.Println(fi106, v1)
                                            var v115 uint8 = 127
                                            delete(v3, (fi106 - (v110 << ((42) & 63))))
                                            _ = v115
                                        } else {
                                            var v116 float32 = (float32(((((((-2.25 + 3.14159) / (1.5 / 1.5)) * 1e10) - 1e10) * 3.14159) + (1.5 - 3.14159))) / v31)
                                            v1 = int64(v107)
                                            v117 := GenS1{F0: true}
                                            _ = v116
                                            _ = v117
                                        }
                                        _ = v110
                                    }
                                } else {
                                    switch ("" + "a") {
                                    case "a":
                                        fmt.Println(v108)
                                    case "":
                                        v31 += 1e10
                                    case "\n\t":
                                        fmt.Println(v108)
                                    }
                                }
                                _ = fi106
                                _ = v107
                                _ = v108
                            }
                            for fi118 := 0; fi118 < 12; fi118++ {
                                v119 := GenS6{F0: (("hello" + (("foo bar" + "foo bar") + "hello")) + "")}
                                _ = fi118
                                _ = v119
                            }
                        }
                    default:
                        for fi120 := 0; fi120 < 10; fi120++ {
                            v2 = false
                            if ("foo bar" == ("a" + "")) {
                                if v2 {
                                    var v121 int64 = (v1 * (v1 | v1))
                                    switch "foo bar" {
                                    case "\n\t":
                                        v121 += (v121 << ((242) & 63))
                                    }
                                    _ = v121
                                }
                                if (!v2) {
                                    var v122 float64 = 1.5
                                    if (497 <= 65535) {
                                        v31 = 1.5
                                        fmt.Println(v5)
                                    } else {
                                        var v123 bool = (2 >= 2)
                                        v124 := GenS10{F0: 681}
                                        v125 := GenS11{F0: 0}
                                        switch -36 {
                                        case -1:
                                            v2 = (v31 != 0.1)
                                        case 128:
                                            switch (((("foo bar" + "a") + (("hello" + "a") + "a")) + "foo bar") + (("foo bar" + "\n\t") + "foo bar")) {
                                            case "hello":
                                                var v126 int64 = v1
                                                _ = v126
                                            case "foo bar":
                                                switch (-(fi103)) {
                                                case -306:
                                                    v123 = true
                                                case 32768:
                                                    if (9223372036854775807 <= -1) {
                                                        if v102 {
                                                            switch ("foo bar" + (("\n\t" + "\n\t") + (("" + "\n\t") + ("" + "hello")))) {
                                                            case "":
                                                                if v102 {
                                                                    v1 -= v1
                                                                }
                                                            default:
                                                                if (v102 && false) {
                                                                    _ = len(v4)
                                                                } else {
                                                                    var v127 int32 = v5
                                                                    switch "hello" {
                                                                    case "foo bar":
                                                                        fmt.Println(v2, v122, v102)
                                                                    default:
                                                                        if len(v6) > 0 {
                                                                            v6[0] = 2147483648
                                                                        }
                                                                    }
                                                                    _ = v127
                                                                }
                                                            }
                                                            var v128 int64 = v1
                                                            var v129 int8 = 10
                                                            _ = v128
                                                            _ = v129
                                                        } else {
                                                            var v130 uint16 = 255
                                                            var v131 int64 = v1
                                                            fmt.Println(v102, v102)
                                                            _ = len(v3)
                                                            _ = v130
                                                            _ = v131
                                                        }
                                                    } else {
                                                        var v132 float64 = 1e10
                                                        var v133 uint16 = 65535
                                                        var v134 uint = uint(v133)
                                                        _ = v132
                                                        _ = v133
                                                        _ = v134
                                                    }
                                                default:
                                                    var v135 int = fi120
                                                    var v136 int8 = -1
                                                    if v123 {
                                                        fmt.Println(v123)
                                                        v137 := GenS12{F0: (fi103 + (fi103 | fi103)), F1: v136}
                                                        var v138 uint16 = 999
                                                        _ = v137
                                                        _ = v138
                                                    }
                                                    _ = v135
                                                    _ = v136
                                                }
                                            default:
                                                if false {
                                                    var v139 uint32 = 65535
                                                    var v140 float64 = float64((v31 - ((v31 / 1e10) * 1.5)))
                                                    if v102 {
                                                        v4[563] = "foo bar"
                                                    }
                                                    _ = v139
                                                    _ = v140
                                                } else {
                                                    v31 = v31
                                                    var v141 uint = 894
                                                    var v142 uint64 = 1
                                                    _ = v141
                                                    _ = v142
                                                }
                                            }
                                        case 508:
                                            if false {
                                                v143 := GenS13{F0: (v1 + (v1 + v1)), F1: ((-(v1)) < v1)}
                                                _ = v143
                                            }
                                        }
                                        _ = v123
                                        _ = v124
                                        _ = v125
                                    }
                                    _ = v122
                                } else {
                                    var v144 int = (fi103 + (fi103 + fi120))
                                    _ = v144
                                }
                                if v2 {
                                    if (74 == 0) {
                                        var v145 uint32 = 268
                                        _ = v145
                                    } else {
                                        var v146 int8 = 100
                                        _ = v146
                                    }
                                }
                            }
                            _ = fi120
                        }
                    }
                    fmt.Println(fi103, v2)
                    _ = fi103
                }
                if true {
                    v5 = -1
                } else {
                    var v147 float32 = (float32(3.14159) + v31)
                    for fi148 := 0; fi148 < 3; fi148++ {
                        if ((0.1 != (-2.25 * float64(v147))) || (0 < -76)) {
                            switch "" {
                            case "":
                                v149 := GenS4{F0: 128, F1: 982, F2: 42}
                                _ = v149
                            default:
                                v150 := GenS14{F0: v5, F1: 4294967295, F2: "\n\t"}
                                _ = v150
                            }
                        }
                        fmt.Println(v31)
                        var v151 int64 = (v1 ^ v1)
                        _ = fi148
                        _ = v151
                    }
                    _ = v147
                }
                for fi152 := 0; fi152 < 14; fi152++ {
                    v153 := GenS1{F0: (32767 >= 0)}
                    v154 := v3[0]
                    _ = fi152
                    _ = v153
                    _ = v154
                }
            }
            fmt.Println(v31, v31, v5)
        } else {
            var v155 int8 = -128
            fmt.Println(v155, v155)
            _ = v155
        }
    case 32767:
        if (54 >= 236) {
            if v102 {
                v156 := v3[-2147483648]
                _ = v156
            } else {
                fmt.Println(v102)
            }
        } else {
            switch int(v1) {
            case 1:
                v31 -= v31
            case 0:
                var v157 uint64 = 4294967295
                _ = v157
            }
        }
    case 4294967295:
        v158 := GenS3{F0: 2147483648}
        _ = v158
    default:
        _ = len(v3)
    }
    switch 42 {
    case 46:
        if v2 {
            for fi159 := 0; fi159 < 9; fi159++ {
                var v160 float32 = v31
                _ = fi159
                _ = v160
            }
        }
    case -9223372036854775808:
        if true {
            if (!true) {
                var v161 uint = 42
                _ = v161
            }
            if (!v2) {
                v162 := GenS8{F0: 4294967295, F1: 4294967295, F2: (-(v5))}
                _ = len(v3)
                _ = v162
            } else {
                for fi163 := 0; fi163 < 6; fi163++ {
                    var v164 int8 = -127
                    v165 := GenS1{F0: ((1 > 255) || v2)}
                    var v166 float32 = v31
                    _ = fi163
                    _ = v164
                    _ = v165
                    _ = v166
                }
            }
            v167 := GenS1{F0: (-2.25 >= (3.14159 + (((3.14159 + float64(v31)) + 0.0) / (3.14159 / (3.14159 * -2.25)))))}
            _ = v167
        }
    }
    v1 = 42
    var v168 uint16 = 832
    v168 -= v168
    fmt.Println("v1:", v1)
    fmt.Println("v102:", v102)
    fmt.Println("v168:", v168)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v102
    _ = v168
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
