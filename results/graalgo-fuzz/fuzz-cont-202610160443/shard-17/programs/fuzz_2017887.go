package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 bool; F1 uint32; F2 uint8 }
type GenS3 struct { F0 string; F1 uint8 }
type GenS4 struct { F0 int16; F1 uint32 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int64; F1 string; F2 uint8 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 bool; F1 int }
type GenS9 struct { F0 int8; F1 uint }
type GenS10 struct { F0 int64; F1 int64 }
type GenS11 struct { F0 string }
type GenS12 struct { F0 int16 }
type GenS13 struct { F0 int64; F1 uint }
type GenS14 struct { F0 uint64; F1 uint; F2 uint16 }
type GenS15 struct { F0 uint; F1 uint8 }
type GenS16 struct { F0 uint32; F1 int; F2 int8 }
type GenS17 struct { F0 uint8; F1 uint }
type GenS18 struct { F0 int16; F1 int32 }
type GenS19 struct { F0 int; F1 int8 }
type GenS20 struct { F0 int32; F1 uint8 }
func genFn0(p0 int, p1 int16, p2 int64) int8 {
    p1 -= p1
    p1 = (p1 / (((p1 - p1)) | 1))
    _ = p0
    _ = p1
    _ = p2
    return 100
}
func genFn1(p0 int, p1 uint32, p2 int16) int {
    p2 = (p2 >> ((42) & 63))
    p1 = 673
    _ = p0
    _ = p1
    _ = p2
    return 42
}

func main() {
    var v1 int32 = 1
    var v2 uint16 = 255
    v3 := make(map[string]string)
    var v4 int32 = v1
    v5 := make(map[int]uint32)
    v6 := make(map[string]string)
    v4 -= v1
    switch 9223372036854775807 {
    case 32767:
        if true {
            v7 := v3["hello"]
            _ = len(v3)
            fmt.Println(v4, v7)
            _ = v7
        } else {
            v2 -= (^(v2))
            if (!(true || true)) {
                v8 := GenS1{F0: -96}
                v4 -= v4
                v9 := GenS2{F0: true, F1: 0, F2: 42}
                _ = v8
                _ = v9
            }
            if false {
                var v10 uint32 = 2147483647
                var v11 int32 = 2147483647
                _ = v10
                _ = v11
            } else {
                var v12 int64 = 0
                v13 := GenS1{F0: -59}
                _ = v12
                _ = v13
            }
        }
    case 128:
        v1 -= (v1 * (v4 >> ((176) & 63)))
    }
    v14 := GenS1{F0: genFn0(0, -161, 673)}
    for fi15 := 0; fi15 < 9; fi15++ {
        if (386 == 264) {
            for fi16 := 0; fi16 < 12; fi16++ {
                fmt.Println(fi15)
                delete(v6, "\n\t")
                _ = fi16
            }
        } else {
            v17 := []int{(fi15 - fi15), fi15}
            fmt.Println(v2, v2)
            _ = v17
        }
        switch "\n\t" {
        case "\n\t":
            if ((false && true) && true) {
                v18 := GenS1{F0: 100}
                v19 := []int{fi15, fi15, fi15, (fi15 % (((^(fi15))) | 1))}
                v6[(("hello" + ("\n\t" + "foo bar")) + ("a" + ("\n\t" + "hello")))] = (("hello" + ("" + ("\n\t" + "a"))) + "hello")
                _ = v18
                _ = v19
            }
        case "foo bar":
            _ = len(v5)
        case "":
            switch "a" {
            case "foo bar":
                for fi20 := 0; fi20 < 17; fi20++ {
                    v1 = v1
                    var v21 uint8 = 255
                    _ = fi20
                    _ = v21
                }
            case "hello":
                v22 := []int{fi15, -425}
                _ = v22
            case "":
                v23 := GenS2{F0: (1.5 <= 1.5), F1: 65535, F2: 0}
                _ = v23
            }
        default:
            var v24 uint64 = 18446744073709551615
            _ = v24
        }
        v2 += 938
        _ = fi15
    }
    var v25 uint8 = 167
    switch "foo bar" {
    case "\n\t":
        var v26 int32 = 951
        _ = v26
    }
    if false {
        var v27 bool = false
        if (!true) {
            v1 -= v4
            if true {
                v28 := GenS3{F0: "foo bar", F1: 127}
                var v29 uint = 2
                _ = v28
                _ = v29
            } else {
                _ = len(v5)
                if (150 != 42) {
                    for fi30 := 0; fi30 < 8; fi30++ {
                        switch genFn1(fi30, 65535, 521) {
                        case 32767:
                            v4 -= v4
                        default:
                            for fi31 := 0; fi31 < 12; fi31++ {
                                var v32 string = ""
                                fmt.Println(v4, v1, v25)
                                if ((v1 * (v4 / ((v4) | 1))) >= v4) {
                                    v4 = (v1 | (v1 >> ((252) & 63)))
                                    v4 = v1
                                    if (v27 || v27) {
                                        v4 += v4
                                        v2 -= v2
                                        if true {
                                            if v27 {
                                                if v27 {
                                                    var v33 uint32 = 42
                                                    if (1 == 1) {
                                                        v34 := GenS1{F0: 75}
                                                        v2 -= v2
                                                        _ = v34
                                                    } else {
                                                        fmt.Println(fi31, v1, v4)
                                                    }
                                                    _ = v33
                                                }
                                                v35 := []int{fi31, (fi31 & 0), (fi30 % ((1) | 1)), fi31}
                                                v27 = true
                                                _ = v35
                                            } else {
                                                v36 := GenS1{F0: 100}
                                                _ = v36
                                            }
                                            v4 = -1
                                            switch (fi31 | (fi31 % ((fi31) | 1))) {
                                            case 128:
                                                if v27 {
                                                    var v37 int = (fi30 % (((fi30 & genFn1(fi30, 1, 32767))) | 1))
                                                    _ = v37
                                                }
                                                v1 = v1
                                                var v38 int64 = 42
                                                if v27 {
                                                    v39 := GenS1{F0: -1}
                                                    _ = v39
                                                }
                                                _ = v38
                                            case 9223372036854775807:
                                                if false {
                                                    switch fi30 {
                                                    case 2147483647:
                                                        var v40 int16 = -1
                                                        _ = v40
                                                    default:
                                                        var v41 uint16 = 0
                                                        _ = v41
                                                    }
                                                    var v42 int16 = -692
                                                    if v27 {
                                                        v6[(("foo bar" + ((v32 + ("foo bar" + "a")) + v32)) + v32)] = (("foo bar" + v32) + "a")
                                                        var v43 uint8 = 0
                                                        v42 = v42
                                                        _ = v43
                                                    } else {
                                                        v2 = (v2 << ((0) & 63))
                                                        switch v32 {
                                                        case "":
                                                            v44 := GenS2{F0: (true || (v4 != 2147483647)), F1: 0, F2: (v25 % ((255) | 1))}
                                                            _ = v44
                                                        }
                                                        if (128 == 128) {
                                                            v27 = (835 == 742)
                                                            switch ("" + "foo bar") {
                                                            case "foo bar":
                                                                var v45 uint8 = (^(v25))
                                                                _ = v45
                                                            case "hello":
                                                                delete(v6, v32)
                                                            }
                                                        }
                                                        var v46 float64 = 0.0
                                                        _ = v46
                                                    }
                                                    _ = v42
                                                } else {
                                                    var v47 uint = 4294967295
                                                    v47 -= (v47 >> (((v47 - v47)) & 63))
                                                    switch "foo bar" {
                                                    case "hello":
                                                        _ = len(v6)
                                                    case "":
                                                        v2 = 0
                                                    default:
                                                        var v48 uint8 = (v25 / ((v25) | 1))
                                                        switch "foo bar" {
                                                        case "":
                                                            if ((v27 && v27) && v27) {
                                                                v6[v32] = v32
                                                            }
                                                        case "hello":
                                                            v49 := GenS4{F0: 501, F1: 2}
                                                            _ = v49
                                                        }
                                                        _ = v48
                                                    }
                                                    switch v32 {
                                                    case "":
                                                        var v50 float32 = (0.0 + 1e10)
                                                        _ = v50
                                                    case "\n\t":
                                                        var v51 uint32 = 4294967295
                                                        _ = v51
                                                    case "a":
                                                        v47 += (v47 >> (((v47 & 4294967295)) & 63))
                                                    default:
                                                        if true {
                                                            v52 := GenS3{F0: v32, F1: (^(v25))}
                                                            var v53 int16 = 761
                                                            switch (fi30 >> (((^(v47))) & 63)) {
                                                            case 0:
                                                                switch (fi31 / ((genFn1((-(fi31)), 771, v53)) | 1)) {
                                                                case 1:
                                                                    var v54 int16 = 1
                                                                    _ = v54
                                                                case 127:
                                                                    var v55 string = (v32 + "")
                                                                    _ = v55
                                                                case -101:
                                                                    var v56 string = ((v32 + (v32 + ((("hello" + (("\n\t" + v32) + "a")) + "hello") + "foo bar"))) + ("\n\t" + ("foo bar" + ("\n\t" + "\n\t"))))
                                                                    _ = v56
                                                                }
                                                            case 2147483647:
                                                                v57 := GenS4{F0: v53, F1: 65535}
                                                                _ = v57
                                                            case 32768:
                                                                v58 := []int{(-(fi30)), 0, fi30, (fi30 << ((v47) & 63))}
                                                                _ = v58
                                                            }
                                                            _ = v52
                                                            _ = v53
                                                        }
                                                    }
                                                    v59 := []int{fi31, genFn1(32767, 49, -1), -2147483648}
                                                    _ = v47
                                                    _ = v59
                                                }
                                            }
                                            if true {
                                                v4 += v1
                                                v2 -= (v2 * v2)
                                                if true {
                                                    v60 := v5[32767]
                                                    fmt.Println(fi31, v25, v4)
                                                    _ = v60
                                                } else {
                                                    v4 += v4
                                                }
                                            }
                                            v3[v32] = ("" + v32)
                                            v2 = v2
                                        }
                                    }
                                    if v27 {
                                        var v61 uint8 = v25
                                        v62 := GenS3{F0: "", F1: 255}
                                        v25 += v61
                                        _ = v61
                                        _ = v62
                                    }
                                    delete(v3, v32)
                                } else {
                                    if true {
                                        var v63 int8 = 60
                                        _ = v63
                                    } else {
                                        switch "" {
                                        case "\n\t":
                                            var v64 bool = v27
                                            _ = v64
                                        case "foo bar":
                                            v65 := GenS2{F0: false, F1: 2147483648, F2: 2}
                                            _ = v65
                                        }
                                        v66 := GenS2{F0: (!(1 < 0)), F1: 193, F2: 115}
                                        v1 += (v1 | (-(v4)))
                                        _ = v66
                                    }
                                }
                                fmt.Println(fi31, fi30, v2)
                                _ = fi31
                                _ = v32
                            }
                        }
                        _ = fi30
                    }
                    v67 := v3[(("hello" + "foo bar") + "")]
                    for fi68 := 0; fi68 < 10; fi68++ {
                        var v69 int64 = -1
                        _ = fi68
                        _ = v69
                    }
                    _ = v67
                }
                v25 = (v25 ^ (v25 * v25))
            }
            if (!(v2 >= (v2 % ((v2) | 1)))) {
                for fi70 := 0; fi70 < 7; fi70++ {
                    for fi71 := 0; fi71 < 2; fi71++ {
                        v2 = 973
                        v4 = (v4 >> ((598) & 63))
                        _ = fi71
                    }
                    var v72 int8 = 100
                    switch fi70 {
                    case 4294967295:
                        if (!(v27 || true)) {
                            v73 := GenS5{F0: v25}
                            v74 := []int{}
                            for fi75 := 0; fi75 < 12; fi75++ {
                                var v76 uint8 = v25
                                switch genFn1((fi70 & fi70), 2, 42) {
                                case 1:
                                    if (float64((float32(0.0) + (0.1 / 1.0))) >= -2.25) {
                                        fmt.Println(v72, v1)
                                        v77 := GenS6{F0: int64(v1), F1: "foo bar", F2: (v76 - (v76 / (((v25 % ((v76) | 1))) | 1)))}
                                        _ = v77
                                    } else {
                                        switch (fi75 % ((fi75) | 1)) {
                                        case 42:
                                            v74 = append(v74, (^(fi70)))
                                            var v78 uint8 = v76
                                            v79 := GenS3{F0: "\n\t", F1: v25}
                                            if false {
                                                v80 := GenS7{F0: 65535}
                                                if (v1 >= (v4 & (v4 >> ((4294967295) & 63)))) {
                                                    switch ("\n\t" + ("hello" + "\n\t")) {
                                                    case "foo bar":
                                                        var v81 int = genFn1((fi75 >> ((641) & 63)), 546, -1)
                                                        _ = v81
                                                    case "":
                                                        switch (("" + "\n\t") + "\n\t") {
                                                        case "hello":
                                                            if ((2 != 1) && (!(("hello" <= "\n\t") || (0 > 0)))) {
                                                                v74 = append(v74, genFn1(fi70, 0, -32768))
                                                                var v82 int64 = 0
                                                                v83 := GenS8{F0: (v27 && v27), F1: fi70}
                                                                _ = v82
                                                                _ = v83
                                                            }
                                                        case "foo bar":
                                                            var v84 uint8 = 255
                                                            _ = v84
                                                        }
                                                    }
                                                    v85 := GenS4{F0: -32768, F1: 842}
                                                    _ = v85
                                                }
                                                fmt.Println(v76, v78, v76)
                                                _ = v80
                                            }
                                            if v27 {
                                                v78 += (v78 * 0)
                                                if v27 {
                                                    _ = len(v3)
                                                    if (-2.25 != 3.14159) {
                                                        var v86 float64 = ((3.14159 - (((3.14159 * 0.0) * 1.5) - 0.0)) / float64(1.5))
                                                        if false {
                                                            v1 = 42
                                                            if len(v74) > 0 {
                                                                v74[0] = fi75
                                                            }
                                                        }
                                                        _ = v86
                                                    } else {
                                                        v2 -= v2
                                                        if (!v27) {
                                                            var v87 uint = 474
                                                            if (-1 < (-(v72))) {
                                                                switch "a" {
                                                                case "\n\t":
                                                                    var v88 string = ""
                                                                    _ = v88
                                                                case "a":
                                                                    var v89 int32 = 1
                                                                    _ = v89
                                                                case "foo bar":
                                                                    if false {
                                                                        fmt.Println(v78, v87, v76)
                                                                        if v27 {
                                                                            v74 = append(v74, fi70)
                                                                            if v27 {
                                                                                v1 -= v1
                                                                            }
                                                                            var v90 bool = v27
                                                                            var v91 uint8 = (v76 | v25)
                                                                            _ = v90
                                                                            _ = v91
                                                                        } else {
                                                                            switch ("\n\t" + "a") {
                                                                            case "":
                                                                                switch ("" + "hello") {
                                                                                case "\n\t":
                                                                                    if (v76 < (v78 >> ((703) & 63))) {
                                                                                        var v92 float64 = 1e10
                                                                                        _ = v92
                                                                                    }
                                                                                case "foo bar":
                                                                                    if (fi70 <= genFn1((fi70 / (((^(fi70))) | 1)), 2147483648, 1000)) {
                                                                                        switch "foo bar" {
                                                                                        case "a":
                                                                                            fmt.Println(v76, v2, fi70)
                                                                                        case "foo bar":
                                                                                            v93 := GenS9{F0: (v72 - v72), F1: v87}
                                                                                            _ = v93
                                                                                        }
                                                                                        var v94 string = ("foo bar" + ("a" + ""))
                                                                                        v1 = (v1 % ((1) | 1))
                                                                                        v72 = (v72 & v72)
                                                                                        if ((v87 << ((v87) & 63)) >= (v87 | uint(v2))) {
                                                                                            v95 := GenS10{F0: -798, F1: -1}
                                                                                            _ = v95
                                                                                        }
                                                                                        v78 = (v78 + uint8(v87))
                                                                                        var v96 int32 = 0
                                                                                        _ = v94
                                                                                        _ = v96
                                                                                    } else {
                                                                                        var v97 string = "foo bar"
                                                                                        var v98 float64 = 3.14159
                                                                                        _ = v97
                                                                                        _ = v98
                                                                                    }
                                                                                }
                                                                            case "\n\t":
                                                                                var v99 uint16 = 2
                                                                                _ = v99
                                                                            case "a":
                                                                                fmt.Println(v76)
                                                                            default:
                                                                                fmt.Println(v27, v2, fi75)
                                                                            }
                                                                            var v100 int32 = (^(v1))
                                                                            var v101 uint = v87
                                                                            var v102 uint64 = 332
                                                                            _ = len(v5)
                                                                            var v103 uint16 = (v2 << ((v101) & 63))
                                                                            _ = v100
                                                                            _ = v101
                                                                            _ = v102
                                                                            _ = v103
                                                                        }
                                                                    }
                                                                }
                                                            }
                                                            _ = v87
                                                        } else {
                                                            v104 := GenS6{F0: -9223372036854775808, F1: "", F2: 127}
                                                            var v105 uint = 255
                                                            v25 = v25
                                                            _ = v104
                                                            _ = v105
                                                        }
                                                    }
                                                }
                                            } else {
                                                delete(v5, int(v2))
                                                var v106 int64 = -62
                                                var v107 uint8 = (v25 << ((42) & 63))
                                                _ = v106
                                                _ = v107
                                            }
                                            _ = v78
                                            _ = v79
                                        case 2147483647:
                                            if false {
                                                switch "hello" {
                                                case "":
                                                    if len(v74) > 0 {
                                                        v74[0] = 32768
                                                    }
                                                }
                                                var v108 uint = 4294967295
                                                v109 := v5[fi70]
                                                switch "" {
                                                case "":
                                                    switch 532 {
                                                    case 2147483647:
                                                        if false {
                                                            fmt.Println(v109, v4, fi70)
                                                        } else {
                                                            var v110 string = "a"
                                                            v1 = v4
                                                            _ = v110
                                                        }
                                                    case 0:
                                                        if v27 {
                                                            var v111 int = -2147483648
                                                            _ = len(v5)
                                                            _ = v111
                                                        }
                                                    default:
                                                        if true {
                                                            if len(v74) >= 2 {
                                                                v74 = v74[1:len(v74)]
                                                            }
                                                        }
                                                    }
                                                case "hello":
                                                    v109 -= v109
                                                case "foo bar":
                                                    fmt.Println(v109, v109, v4)
                                                default:
                                                    var v112 int32 = (v4 - v4)
                                                    _ = v112
                                                }
                                                _ = v108
                                                _ = v109
                                            } else {
                                                var v113 float32 = (-2.25 / (1e10 + 0.1))
                                                switch genFn1(genFn1((-(fi75)), 1, -1), 943, 0) {
                                                case -1:
                                                    _ = len(v3)
                                                }
                                                _ = v113
                                            }
                                        case -1:
                                            fmt.Println(fi70, v4, v25)
                                        default:
                                            if ((v72 / (((v72 << ((65535) & 63))) | 1)) >= v72) {
                                                if (0 != uint32(v2)) {
                                                    if len(v74) > 0 {
                                                        v74[0] = (^(fi70))
                                                    }
                                                    var v114 uint64 = 560
                                                    _ = v114
                                                }
                                            }
                                        }
                                        v115 := GenS1{F0: genFn0((fi70 - (fi75 + -9223372036854775808)), 1, -704)}
                                        _ = v115
                                    }
                                case 32767:
                                    v4 += (v1 % ((v4) | 1))
                                }
                                v2 = (v2 & v2)
                                v76 -= (v25 >> ((1) & 63))
                                _ = fi75
                                _ = v76
                            }
                            _ = v73
                            _ = v74
                        }
                    }
                    _ = fi70
                    _ = v72
                }
            } else {
                fmt.Println(v25, v25)
            }
        }
        _ = v27
    } else {
        v1 += (v4 >> ((395) & 63))
    }
    var v116 uint64 = 18446744073709551615
    if false {
        var v117 int16 = -334
        _ = v117
    } else {
        for fi118 := 0; fi118 < 12; fi118++ {
            if (true || ("foo bar" >= "")) {
                for fi119 := 0; fi119 < 11; fi119++ {
                    v2 -= v2
                    var v120 int8 = 74
                    _ = fi119
                    _ = v120
                }
                var v121 int16 = 32767
                switch ("hello" + ("" + "a")) {
                case "foo bar":
                    v116 = 467
                }
                _ = v121
            } else {
                delete(v6, "")
                switch "foo bar" {
                case "hello":
                    if (false || false) {
                        for fi122 := 0; fi122 < 8; fi122++ {
                            var v123 int = (fi122 % (((fi118 % (((fi122 + 420)) | 1))) | 1))
                            _ = fi122
                            _ = v123
                        }
                        for fi124 := 0; fi124 < 19; fi124++ {
                            var v125 bool = false
                            var v126 float32 = -2.25
                            _ = fi124
                            _ = v125
                            _ = v126
                        }
                    }
                case "a":
                    if (true && true) {
                        if true {
                            v4 += v1
                            var v127 float32 = 0.1
                            _ = v127
                        }
                        v116 = (v116 + v116)
                    } else {
                        var v128 int8 = 127
                        _ = v128
                    }
                case "foo bar":
                    var v129 int = (fi118 * (fi118 % ((fi118) | 1)))
                    _ = v129
                }
                v130 := []int{genFn1(-768, uint32(v25), -1)}
                _ = v130
            }
            _ = fi118
        }
        delete(v5, -614)
        fmt.Println(v116)
    }
    v131 := GenS4{F0: -32768, F1: 4294967295}
    v132 := v3[("\n\t" + "hello")]
    fmt.Println(v132, v2)
    switch ((v132 + v132) + v132) {
    case "a":
        if (!(((((!true) || false) && false) && (1e10 > (1.5 / -2.25))) && false)) {
            if (false && false) {
                for fi133 := 0; fi133 < 8; fi133++ {
                    for fi134 := 0; fi134 < 2; fi134++ {
                        var v135 bool = false
                        v136 := GenS11{F0: ("hello" + v132)}
                        v137 := GenS2{F0: v135, F1: 42, F2: (v25 % ((v25) | 1))}
                        _ = fi134
                        _ = v135
                        _ = v136
                        _ = v137
                    }
                    switch 264 {
                    case -9223372036854775808:
                        switch v132 {
                        case "\n\t":
                            v138 := v6["\n\t"]
                            _ = v138
                        default:
                            fmt.Println(v2)
                        }
                    case -38:
                        fmt.Println(v132, v116, v25)
                    }
                    _ = fi133
                }
                v139 := []int{genFn1(128, 0, 1), 32767, 923}
                _ = v139
            } else {
                v140 := GenS7{F0: 212}
                switch v132 {
                case "":
                    v141 := GenS1{F0: 1}
                    _ = v141
                case "\n\t":
                    fmt.Println(v1)
                default:
                    fmt.Println(v116, v116, v132)
                }
                _ = v140
            }
        }
    case "\n\t":
        if ((true || (!(-9223372036854775808 == -735))) || false) {
            for fi142 := 0; fi142 < 12; fi142++ {
                v143 := []int{fi142}
                _ = fi142
                _ = v143
            }
            var v144 int32 = v1
            switch (v132 + v132) {
            case "\n\t":
                v116 = v116
            case "a":
                if (3.14159 != ((1e10 + float64(3.14159)) * (-2.25 * -2.25))) {
                    if false {
                        for fi145 := 0; fi145 < 12; fi145++ {
                            var v146 int = fi145
                            v147 := GenS1{F0: 0}
                            var v148 float64 = float64(float32(((-2.25 / 1.5) / 0.1)))
                            _ = fi145
                            _ = v146
                            _ = v147
                            _ = v148
                        }
                        if (0 > -197) {
                            if true {
                                for fi149 := 0; fi149 < 14; fi149++ {
                                    v150 := GenS12{F0: 888}
                                    fmt.Println(fi149)
                                    _ = fi149
                                    _ = v150
                                }
                                if false {
                                    v116 -= (^(v116))
                                } else {
                                    fmt.Println(v25, v1, v1)
                                }
                                v116 -= 1
                            } else {
                                v151 := GenS13{F0: 0, F1: 255}
                                v2 = v2
                                _ = v151
                            }
                            var v152 int = -1
                            _ = v152
                        } else {
                            var v153 uint32 = 65535
                            v153 = 65535
                            _ = v153
                        }
                        if (false && false) {
                            delete(v6, (v132 + (v132 + "")))
                        }
                    }
                    if true {
                        fmt.Println(v4)
                    }
                }
            case "hello":
                for fi154 := 0; fi154 < 1; fi154++ {
                    var v155 bool = true
                    v132 = v132
                    v156 := []int{fi154}
                    _ = fi154
                    _ = v155
                    _ = v156
                }
            default:
                for fi157 := 0; fi157 < 16; fi157++ {
                    switch (v132 + (v132 + v132)) {
                    case "hello":
                        switch 641 {
                        case 128:
                            v25 = (v25 & v25)
                        case 32768:
                            for fi158 := 0; fi158 < 6; fi158++ {
                                v25 = v25
                                if true {
                                    var v159 bool = true
                                    v116 = (v116 << ((482) & 63))
                                    v4 += (v4 - 578)
                                    var v160 uint16 = (v2 / ((32767) | 1))
                                    var v161 int = int(v25)
                                    _ = v159
                                    _ = v160
                                    _ = v161
                                }
                                fmt.Println(v25, v132, v4)
                                if true {
                                    switch fi157 {
                                    case -9223372036854775808:
                                        switch fi158 {
                                        case 2147483647:
                                            var v162 uint16 = uint16(v116)
                                            v144 += 7
                                            _ = v162
                                        case 9223372036854775807:
                                            v4 += v4
                                            switch v132 {
                                            case "":
                                                switch "" {
                                                case "foo bar":
                                                    var v163 string = "\n\t"
                                                    fmt.Println(v116, fi158)
                                                    _ = v163
                                                }
                                            case "foo bar":
                                                var v164 float32 = float32(float64((1e10 + 0.1)))
                                                _ = v164
                                            }
                                        default:
                                            var v165 int = fi158
                                            _ = v165
                                        }
                                    case -2147483648:
                                        delete(v5, fi157)
                                    }
                                    fmt.Println(v116)
                                }
                                v166 := GenS6{F0: -588, F1: "foo bar", F2: (v25 - (v25 >> ((18446744073709551615) & 63)))}
                                _ = fi158
                                _ = v166
                            }
                        default:
                            for fi167 := 0; fi167 < 10; fi167++ {
                                fmt.Println(fi167)
                                if ((v116 >> ((1) & 63)) != 4294967295) {
                                    v4 = int32(v116)
                                    v168 := GenS8{F0: false, F1: 15}
                                    var v169 uint64 = v116
                                    switch fi157 {
                                    case 2147483647:
                                        var v170 bool = (genFn0(-9223372036854775808, -416, 9223372036854775807) == 32)
                                        _ = v170
                                    case 1:
                                        fmt.Println(v116)
                                    default:
                                        v171 := GenS6{F0: 42, F1: "hello", F2: 128}
                                        _ = v171
                                    }
                                    delete(v3, v132)
                                    _ = v168
                                    _ = v169
                                }
                                var v172 int64 = -1
                                var v173 string = v132
                                v174 := []int{fi167, fi157}
                                _ = fi167
                                _ = v172
                                _ = v173
                                _ = v174
                            }
                        }
                    case "\n\t":
                        v4 = (^(v144))
                    }
                    switch int(v116) {
                    case 42:
                        if (-621 <= -713) {
                            for fi175 := 0; fi175 < 5; fi175++ {
                                var v176 int32 = -576
                                var v177 float32 = 3.14159
                                _ = fi175
                                _ = v176
                                _ = v177
                            }
                        } else {
                            v178 := GenS14{F0: (v116 >> ((65535) & 63)), F1: 139, F2: v2}
                            var v179 uint32 = 2
                            _ = v178
                            _ = v179
                        }
                    }
                    switch genFn1(fi157, 2, 699) {
                    case 1:
                        if true {
                            if ((32767 > -9223372036854775808) || ((v132 + v132) > (v132 + v132))) {
                                v3["foo bar"] = (("\n\t" + "a") + "")
                                fmt.Println(v2)
                            }
                        }
                    case 127:
                        v180 := GenS8{F0: false, F1: (-(fi157))}
                        _ = v180
                    case -200:
                        fmt.Println(v144)
                    default:
                        if true {
                            v181 := GenS5{F0: (v25 * v25)}
                            _ = v181
                        }
                    }
                    _ = fi157
                }
            }
            _ = v144
        } else {
            for fi182 := 0; fi182 < 17; fi182++ {
                if true {
                    v183 := GenS15{F0: 255, F1: uint8(v2)}
                    _ = v183
                } else {
                    _ = len(v3)
                    if false {
                        if true {
                            v184 := GenS16{F0: 2147483648, F1: int(v2), F2: 100}
                            _ = v184
                        }
                    }
                    if false {
                        v25 = 0
                    } else {
                        var v185 float64 = (0.1 - 3.14159)
                        _ = v185
                    }
                }
                _ = fi182
            }
            for fi186 := 0; fi186 < 12; fi186++ {
                v187 := GenS2{F0: false, F1: 68, F2: 42}
                switch "a" {
                case "\n\t":
                    for fi188 := 0; fi188 < 8; fi188++ {
                        var v189 uint64 = v116
                        switch int(v4) {
                        case -636:
                            if true {
                                v190 := []int{(^(fi186)), -9223372036854775808, (fi186 << ((2) & 63))}
                                _ = len(v6)
                                v191 := GenS1{F0: -121}
                                if (!((!false) && false)) {
                                    v192 := GenS13{F0: 1, F1: 758}
                                    fmt.Println(v116, v1, v116)
                                    _ = len(v190)
                                    _ = v192
                                }
                                _ = v190
                                _ = v191
                            } else {
                                v193 := GenS17{F0: v25, F1: 42}
                                fmt.Println(v25, fi186, v1)
                                _ = v193
                            }
                        case 32768:
                            switch v132 {
                            case "hello":
                                v132 = "a"
                            default:
                                var v194 int32 = 42
                                _ = v194
                            }
                        case 4294967295:
                            fmt.Println(v116, fi188, v132)
                        default:
                            var v195 uint = 42
                            _ = v195
                        }
                        _ = fi188
                        _ = v189
                    }
                case "foo bar":
                    var v196 int = (fi186 & fi186)
                    _ = v196
                case "hello":
                    v197 := GenS7{F0: 2147483647}
                    _ = v197
                }
                _ = fi186
                _ = v187
            }
        }
    }
    fmt.Println(v132, v25, v4)
    if false {
        var v198 int8 = genFn0(573, -638, 1)
        v199 := GenS2{F0: (false || false), F1: 92, F2: v25}
        _ = v198
        _ = v199
    } else {
        v200 := GenS18{F0: 0, F1: (^(v4))}
        _ = v200
    }
    var v201 uint8 = (^(v25))
    for fi202 := 0; fi202 < 9; fi202++ {
        if false {
            for fi203 := 0; fi203 < 3; fi203++ {
                v204 := []int{fi203, fi203, fi203}
                _ = fi203
                _ = v204
            }
            if ((1e10 - -2.25) >= (1e10 + ((0.1 + 3.14159) / (1e10 / 0.1)))) {
                var v205 int16 = 32767
                switch "\n\t" {
                case "foo bar":
                    v206 := GenS15{F0: 0, F1: 11}
                    _ = v206
                case "a":
                    v1 = v4
                case "hello":
                    for fi207 := 0; fi207 < 6; fi207++ {
                        if (1 >= 1) {
                            v208 := v6[v132]
                            _ = v208
                        } else {
                            v209 := []int{(fi202 - fi202), -9223372036854775808, 42, fi202}
                            var v210 int = fi202
                            _ = v209
                            _ = v210
                        }
                        v211 := GenS19{F0: (-(fi207)), F1: 104}
                        _ = fi207
                        _ = v211
                    }
                }
                _ = v205
            } else {
                v4 += int32(fi202)
                switch (fi202 / (((fi202 - (fi202 << ((1) & 63)))) | 1)) {
                case 127:
                    v212 := []int{fi202, int(v116), (fi202 - 0), (fi202 + fi202)}
                    _ = v212
                default:
                    v1 = v1
                }
            }
            var v213 float64 = (-2.25 * -2.25)
            _ = v213
        }
        var v214 int32 = (v1 & (v1 - int32(fi202)))
        _ = fi202
        _ = v214
    }
    for fi215 := 0; fi215 < 7; fi215++ {
        for fi216 := 0; fi216 < 7; fi216++ {
            v217 := GenS20{F0: 42, F1: v25}
            var v218 string = ""
            _ = fi216
            _ = v217
            _ = v218
        }
        _ = fi215
    }
    fmt.Println("v1:", v1)
    fmt.Println("v116:", v116)
    fmt.Println("v132:", v132)
    fmt.Println("v2:", v2)
    fmt.Println("v201:", v201)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v116
    _ = v131
    _ = v132
    _ = v14
    _ = v2
    _ = v201
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
