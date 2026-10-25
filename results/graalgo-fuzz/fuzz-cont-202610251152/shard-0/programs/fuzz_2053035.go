package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint32; F1 int64; F2 uint16 }
type GenS3 struct { F0 int64; F1 uint8; F2 int16 }
type GenS4 struct { F0 int16; F1 int16; F2 string }
type GenS5 struct { F0 int8; F1 uint; F2 int32 }
type GenS6 struct { F0 int8; F1 uint8 }
type GenS7 struct { F0 int64; F1 string; F2 uint32 }

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[int]bool)
    var v3 int32 = 2147483647
    var v4 float32 = 0.0
    v5 := make(map[int]string)
    switch -9223372036854775808 {
    case 127:
        v6 := []int{}
        _ = v6
    default:
        if (42 >= 4294967295) {
            if false {
                v7 := v1[527]
                v3 += v3
                _ = v7
            }
            var v8 uint32 = 0
            var v9 float64 = 1.5
            _ = v8
            _ = v9
        } else {
            v3 = v3
            switch 128 {
            case -2147483648:
                v3 -= (v3 ^ (v3 >> ((18446744073709551615) & 63)))
            case 4294967295:
                var v10 string = "\n\t"
                _ = v10
            case 2147483647:
                v11 := []int{1, 42}
                _ = v11
            default:
                var v12 bool = true
                _ = v12
            }
            delete(v2, 0)
        }
    }
    var v13 float64 = float64(v4)
    var v14 uint64 = 287
    var v15 uint16 = 139
    if true {
        v16 := GenS1{F0: -9223372036854775808}
        for fi17 := 0; fi17 < 8; fi17++ {
            for fi18 := 0; fi18 < 13; fi18++ {
                switch "\n\t" {
                case "hello":
                    v19 := v5[fi17]
                    _ = v19
                }
                var v20 bool = (uint32(v14) > 2147483648)
                fmt.Println(v15, v14)
                _ = fi18
                _ = v20
            }
            if (!((v15 - v15) >= v15)) {
                if (false && (237 != (v15 & (v15 / (((^(v15))) | 1))))) {
                    for fi21 := 0; fi21 < 1; fi21++ {
                        if (v4 != v4) {
                            v4 = (float32(v13) / 3.14159)
                            var v22 string = "\n\t"
                            _ = v22
                        }
                        v3 += 42
                        v4 = (v4 - ((v4 - 3.14159) + float32(0.0)))
                        _ = fi21
                    }
                }
                v4 += 1e10
                v23 := GenS1{F0: 422}
                _ = v23
            } else {
                var v24 string = "foo bar"
                switch v24 {
                case "":
                    fmt.Println(v3, v14, v24)
                case "foo bar":
                    if (1 != -32768) {
                        var v25 string = v24
                        for fi26 := 0; fi26 < 12; fi26++ {
                            var v27 uint8 = 159
                            if ((v3 / ((v3) | 1)) < v3) {
                                if true {
                                    fmt.Println(v4, v4, v13)
                                } else {
                                    v15 = (v15 ^ v15)
                                    var v28 int16 = 1000
                                    _ = v28
                                }
                            } else {
                                var v29 float32 = 3.14159
                                var v30 uint64 = (v14 % ((v14) | 1))
                                _ = v29
                                _ = v30
                            }
                            _ = fi26
                            _ = v27
                        }
                        switch v25 {
                        case "":
                            if false {
                                var v31 uint = 710
                                _ = v31
                            } else {
                                var v32 int16 = -502
                                if false {
                                    for fi33 := 0; fi33 < 5; fi33++ {
                                        if (127 > 127) {
                                            var v34 string = v24
                                            if ((-430 >= v3) || (true && true)) {
                                                switch 127 {
                                                case 127:
                                                    if ((v14 - 0) < (v14 - v14)) {
                                                        if ((1 >= 39) && false) {
                                                            v35 := GenS2{F0: 2147483647, F1: 513, F2: v15}
                                                            v36 := GenS2{F0: 1, F1: -9, F2: v15}
                                                            _ = v35
                                                            _ = v36
                                                        } else {
                                                            v32 += v32
                                                        }
                                                    } else {
                                                        v24 = (v24 + v24)
                                                        switch (fi33 + (fi17 ^ (fi33 + int(v32)))) {
                                                        case 128:
                                                            v37 := GenS3{F0: 866, F1: 255, F2: (-(v32))}
                                                            _ = v37
                                                        default:
                                                            v25 = "a"
                                                        }
                                                    }
                                                }
                                            } else {
                                                if (v14 > 4294967295) {
                                                    v13 -= 0.0
                                                    v15 = v15
                                                }
                                                fmt.Println(fi33, v4)
                                            }
                                            delete(v5, -136)
                                            _ = v34
                                        } else {
                                            switch (((("hello" + v24) + "") + "hello") + "a") {
                                            case "hello":
                                                v4 = (3.14159 + ((1e10 / v4) * 0.1))
                                            case "foo bar":
                                                v38 := v2[fi17]
                                                _ = v38
                                            }
                                        }
                                        v39 := GenS4{F0: 406, F1: v32, F2: (v25 + ((v25 + ("foo bar" + v25)) + v24))}
                                        if (fi33 != fi33) {
                                            var v40 float32 = v4
                                            var v41 int16 = v32
                                            v3 -= 1
                                            _ = v40
                                            _ = v41
                                        }
                                        _ = fi33
                                        _ = v39
                                    }
                                    switch (fi17 % (((fi17 >> ((42) & 63))) | 1)) {
                                    case 0:
                                        v24 = (v25 + v25)
                                    }
                                }
                                v4 -= v4
                                _ = v32
                            }
                        case "foo bar":
                            fmt.Println(v15, v3, v25)
                        case "\n\t":
                            var v42 int8 = 127
                            _ = v42
                        }
                        _ = v25
                    } else {
                        v3 -= (v3 & v3)
                    }
                }
                v14 += (^(v14))
                _ = v24
            }
            if true {
                var v43 float32 = 1e10
                _ = v43
            } else {
                for fi44 := 0; fi44 < 16; fi44++ {
                    switch "" {
                    case "":
                        if false {
                            v45 := v2[fi17]
                            var v46 uint8 = 102
                            v47 := []int{(fi44 | (fi17 >> ((42) & 63))), fi17, (^(fi44)), fi44}
                            _ = v45
                            _ = v46
                            _ = v47
                        }
                    }
                    _ = fi44
                }
                var v48 int64 = 9223372036854775807
                switch "" {
                case "":
                    v13 -= v13
                case "hello":
                    var v49 uint32 = 65535
                    _ = v49
                }
                _ = v48
            }
            _ = fi17
        }
        v15 += v15
        _ = v16
    }
    fmt.Println(v3, v4)
    if (true && false) {
        fmt.Println(v13, v14, v14)
    }
    var v50 int32 = v3
    v3 = v3
    var v51 uint32 = 234
    fmt.Println(v50, v13, v51)
    var v52 int32 = (v3 ^ -6)
    if true {
        for fi53 := 0; fi53 < 11; fi53++ {
            v54 := GenS1{F0: -173}
            for fi55 := 0; fi55 < 2; fi55++ {
                v50 -= -543
                var v56 int8 = -1
                _ = fi55
                _ = v56
            }
            var v57 float32 = v4
            _ = fi53
            _ = v54
            _ = v57
        }
    } else {
        delete(v1, v51)
        if true {
            if (2 >= 0) {
                v14 = (v14 << ((255) & 63))
                var v58 uint16 = (v15 << ((42) & 63))
                _ = v58
            } else {
                v13 = 3.14159
                var v59 uint32 = v51
                _ = v59
            }
        }
        for fi60 := 0; fi60 < 18; fi60++ {
            fmt.Println(v3)
            switch ("\n\t" + ("foo bar" + "foo bar")) {
            case "\n\t":
                if (fi60 >= (^(fi60))) {
                    v14 -= 536
                    if true {
                        var v61 float64 = v13
                        _ = v61
                    } else {
                        v62 := []int{(fi60 >> ((2) & 63)), (fi60 & fi60), (^(fi60))}
                        for fi63 := 0; fi63 < 2; fi63++ {
                            var v64 int64 = -1
                            _ = len(v1)
                            var v65 uint16 = v15
                            v13 = 1.5
                            switch (fi63 >> ((992) & 63)) {
                            case 140:
                                if (163 == 1) {
                                    switch fi60 {
                                    case 127:
                                        fmt.Println(v4, v4)
                                    default:
                                        v51 -= v51
                                    }
                                }
                            case 32767:
                                var v66 uint16 = v65
                                v67 := GenS5{F0: int8(v50), F1: 4294967295, F2: 0}
                                _ = v66
                                _ = v67
                            case 0:
                                _ = len(v5)
                            }
                            var v68 uint16 = (^(v15))
                            if (!(0 >= 255)) {
                                v69 := GenS2{F0: v51, F1: v64, F2: v68}
                                fmt.Println(v50, v65, v52)
                                var v70 float64 = 0.0
                                _ = v69
                                _ = v70
                            }
                            _ = fi63
                            _ = v64
                            _ = v65
                            _ = v68
                        }
                        _ = v62
                    }
                    var v71 string = "foo bar"
                    _ = v71
                } else {
                    var v72 int64 = -657
                    if (0.0 <= v4) {
                        for fi73 := 0; fi73 < 17; fi73++ {
                            v74 := GenS5{F0: int8(v15), F1: 784, F2: v52}
                            if ((v52 << ((42) & 63)) == (v52 - (v50 << ((1) & 63)))) {
                                v51 = (v51 / (((v51 >> ((1) & 63))) | 1))
                                var v75 int32 = v52
                                switch "hello" {
                                case "foo bar":
                                    switch ("hello" + "a") {
                                    case "hello":
                                        if (!true) {
                                            if false {
                                                v76 := GenS3{F0: v72, F1: 144, F2: 1}
                                                v50 = v52
                                                if false {
                                                    switch "\n\t" {
                                                    case "a":
                                                        var v77 uint64 = v14
                                                        _ = v77
                                                    case "":
                                                        var v78 uint32 = (^(v51))
                                                        _ = v78
                                                    case "foo bar":
                                                        v79 := GenS3{F0: v72, F1: 42, F2: -683}
                                                        _ = v79
                                                    }
                                                    if true {
                                                        switch (("foo bar" + ("\n\t" + (("" + "hello") + ""))) + ("hello" + ("" + (("\n\t" + "") + "")))) {
                                                        case "\n\t":
                                                            v80 := []int{(fi60 * int(v15)), fi60, (fi60 % ((fi60) | 1))}
                                                            var v81 bool = true
                                                            _ = v80
                                                            _ = v81
                                                        case "":
                                                            v75 -= v52
                                                        }
                                                        v75 = (v52 - int32(v14))
                                                        v15 = v15
                                                        var v82 bool = (1 != 4294967295)
                                                        _ = v82
                                                    }
                                                    v50 += 0
                                                } else {
                                                    v83 := GenS5{F0: 71, F1: 65535, F2: v52}
                                                    var v84 int16 = -46
                                                    _ = v83
                                                    _ = v84
                                                }
                                                _ = v76
                                            } else {
                                                var v85 uint64 = v14
                                                _ = v85
                                            }
                                            var v86 int16 = -1
                                            _ = v86
                                        }
                                    default:
                                        v87 := GenS1{F0: v72}
                                        _ = v87
                                    }
                                default:
                                    var v88 uint8 = 42
                                    _ = v88
                                }
                                _ = v75
                            }
                            v72 = (v72 >> ((2) & 63))
                            _ = fi73
                            _ = v74
                        }
                        delete(v2, -728)
                    } else {
                        switch fi60 {
                        case -1:
                            switch "a" {
                            case "\n\t":
                                var v89 uint = 18446744073709551615
                                _ = v89
                            case "foo bar":
                                if (true && true) {
                                    _ = len(v1)
                                    v90 := GenS2{F0: v51, F1: v72, F2: (v15 * v15)}
                                    _ = v90
                                }
                            case "a":
                                fmt.Println(v4, v15, v51)
                            default:
                                switch (fi60 / ((fi60) | 1)) {
                                case 4294967295:
                                    switch "hello" {
                                    case "a":
                                        v91 := GenS6{F0: 36, F1: 23}
                                        _ = v91
                                    case "hello":
                                        fmt.Println(v3, v14)
                                    case "foo bar":
                                        var v92 string = "hello"
                                        _ = v92
                                    }
                                case 0:
                                    if false {
                                        if (1e10 == -2.25) {
                                            v93 := GenS4{F0: -1, F1: -32768, F2: ""}
                                            _ = v93
                                        }
                                        v52 += v50
                                        v15 += 65535
                                    } else {
                                        v94 := GenS1{F0: v72}
                                        v72 += (v72 >> ((1) & 63))
                                        switch fi60 {
                                        case -9223372036854775808:
                                            v95 := GenS2{F0: v51, F1: v72, F2: 65535}
                                            _ = v95
                                        case -528:
                                            for fi96 := 0; fi96 < 17; fi96++ {
                                                if true {
                                                    if (0 <= (v72 << ((65535) & 63))) {
                                                        v4 -= v4
                                                    }
                                                    var v97 uint64 = (^(v14))
                                                    fmt.Println(v51, v13, v4)
                                                    _ = v97
                                                }
                                                var v98 int32 = 0
                                                _ = fi96
                                                _ = v98
                                            }
                                        default:
                                            switch ("a" + ("\n\t" + "\n\t")) {
                                            case "a":
                                                var v99 uint32 = v51
                                                _ = v99
                                            case "\n\t":
                                                for fi100 := 0; fi100 < 20; fi100++ {
                                                    v101 := GenS7{F0: v72, F1: "foo bar", F2: (v51 - v51)}
                                                    _ = fi100
                                                    _ = v101
                                                }
                                            case "hello":
                                                switch "" {
                                                case "a":
                                                    delete(v1, v51)
                                                case "foo bar":
                                                    delete(v5, fi60)
                                                case "":
                                                    v13 = (float64((1.5 * v4)) + v13)
                                                default:
                                                    fmt.Println(v15, v4)
                                                }
                                            }
                                        }
                                        _ = v94
                                    }
                                }
                            }
                        case 32768:
                            delete(v1, v51)
                        }
                        var v102 int32 = v3
                        _ = v102
                    }
                    if (42 <= 2) {
                        var v103 uint32 = v51
                        delete(v5, (-(fi60)))
                        var v104 uint8 = 0
                        _ = v103
                        _ = v104
                    }
                    _ = v72
                }
            default:
                var v105 uint = 993
                _ = v105
            }
            v106 := []int{(fi60 | (fi60 + -516)), 42, -942}
            _ = fi60
            _ = v106
        }
    }
    var v107 uint = 2
    var v108 uint64 = v14
    var v109 uint8 = 36
    switch ("a" + ("a" + "hello")) {
    case "\n\t":
        if false {
            _ = len(v1)
            if (-1 < 0) {
                v110 := v1[0]
                _ = v110
            } else {
                switch (("foo bar" + "") + "") {
                case "hello":
                    fmt.Println(v15)
                }
                var v111 int = 698
                for fi112 := 0; fi112 < 11; fi112++ {
                    v111 -= 42
                    _ = fi112
                }
                _ = v111
            }
            var v113 bool = true
            _ = v113
        }
    case "hello":
        for fi114 := 0; fi114 < 19; fi114++ {
            switch (fi114 & (fi114 % ((fi114) | 1))) {
            case 4294967295:
                v108 -= 42
            case -794:
                for fi115 := 0; fi115 < 11; fi115++ {
                    v14 = v108
                    v116 := []int{fi115, fi114, (fi115 << ((0) & 63)), int(v107)}
                    var v117 bool = false
                    _ = fi115
                    _ = v116
                    _ = v117
                }
            }
            _ = fi114
        }
    case "foo bar":
        v51 = v51
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v107:", v107)
    fmt.Println("v108:", v108)
    fmt.Println("v109:", v109)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v107
    _ = v108
    _ = v109
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v51
    _ = v52
}
