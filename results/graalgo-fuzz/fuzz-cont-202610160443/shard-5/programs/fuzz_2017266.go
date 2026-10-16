package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64 }
type GenS2 struct { F0 uint32; F1 uint8 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int32; F1 string; F2 uint8 }
type GenS5 struct { F0 string; F1 int16 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 int }
type GenS8 struct { F0 uint16; F1 uint8; F2 int }
type GenS9 struct { F0 uint8; F1 int }
type GenS10 struct { F0 int16; F1 int16 }
type GenS11 struct { F0 string; F1 int; F2 int64 }
type GenS12 struct { F0 uint8; F1 string; F2 string }
type GenS13 struct { F0 uint8; F1 uint8 }
type GenS14 struct { F0 uint16; F1 int64 }
type GenS15 struct { F0 uint64; F1 int16; F2 int64 }
type GenS16 struct { F0 int32; F1 string; F2 uint8 }
type GenS17 struct { F0 int8 }
type GenS18 struct { F0 uint32; F1 int16 }

func main() {
    var v1 uint64 = 657
    v2 := []int{0, 128, -9223372036854775808}
    v3 := make(map[string]string)
    var v4 uint32 = 1
    var v5 int = int(v1)
    var v6 int64 = -483
    if true {
        for fi7 := 0; fi7 < 16; fi7++ {
            switch ("a" + "a") {
            case "a":
                if ((!true) || (3.14159 >= 0.1)) {
                    if (!false) {
                        v8 := GenS1{F0: (v1 * 836), F1: (v1 * (v1 << ((1) & 63)))}
                        switch fi7 {
                        case -2147483648:
                            var v9 float32 = (1e10 - (float32(1e10) * (0.0 / (float32(float64(3.14159)) / 0.1))))
                            _ = v9
                        case -187:
                            if false {
                                v6 += (v6 ^ v6)
                                var v10 int64 = -53
                                var v11 uint = 523
                                _ = v10
                                _ = v11
                            } else {
                                var v12 uint8 = 128
                                for fi13 := 0; fi13 < 18; fi13++ {
                                    var v14 uint = 271
                                    if len(v2) > 0 {
                                        v2[0] = 9223372036854775807
                                    }
                                    _ = fi13
                                    _ = v14
                                }
                                fmt.Println(fi7, v6)
                                _ = v12
                            }
                        case 9223372036854775807:
                            fmt.Println(v1, v4, fi7)
                        }
                        var v15 int32 = -783
                        _ = v15
                        _ = v8
                    } else {
                        if ((((v6 / (((-(v6))) | 1)) <= (v6 & (v6 - (v6 + v6)))) || (-2.25 < (-2.25 + 1e10))) || true) {
                            v16 := GenS1{F0: (v1 | uint64(v4)), F1: (v1 + v1)}
                            _ = v16
                        } else {
                            switch "a" {
                            case "foo bar":
                                if (!true) {
                                    if false {
                                        var v17 bool = (127 != 0)
                                        _ = v17
                                    }
                                    _ = len(v2)
                                    for fi18 := 0; fi18 < 1; fi18++ {
                                        fmt.Println(fi18)
                                        switch "hello" {
                                        case "a":
                                            v3[("" + "")] = "\n\t"
                                        default:
                                            if true {
                                                v1 = v1
                                                v4 = 0
                                                var v19 string = "a"
                                                var v20 int8 = 1
                                                var v21 bool = ((-(v20)) > int8(v6))
                                                v21 = true
                                                if (126 == 126) {
                                                    var v22 int64 = v6
                                                    switch ("\n\t" + v19) {
                                                    case "":
                                                        v21 = (-2147483648 <= int32(v1))
                                                    case "\n\t":
                                                        v23 := GenS2{F0: (v4 >> ((211) & 63)), F1: uint8(v4)}
                                                        _ = v23
                                                    case "hello":
                                                        var v24 uint16 = 65535
                                                        _ = v24
                                                    }
                                                    if v21 {
                                                        if v21 {
                                                            if len(v2) > 0 {
                                                                v2[0] = fi18
                                                            }
                                                        } else {
                                                            var v25 float32 = 1.5
                                                            switch v19 {
                                                            case "foo bar":
                                                                v26 := v3[v19]
                                                                _ = v26
                                                            case "\n\t":
                                                                var v27 uint16 = 65535
                                                                _ = v27
                                                            case "hello":
                                                                v28 := GenS1{F0: v1, F1: (v1 >> ((65535) & 63))}
                                                                _ = v28
                                                            default:
                                                                fmt.Println(v20, v4, v25)
                                                            }
                                                            _ = v25
                                                        }
                                                        v19 = v19
                                                        if (!true) {
                                                            switch fi18 {
                                                            case 4294967295:
                                                                fmt.Println(fi18, v6)
                                                            case -666:
                                                                if v21 {
                                                                    switch (v19 + "hello") {
                                                                    case "a":
                                                                        fmt.Println(v5)
                                                                        v29 := v3["a"]
                                                                        v19 = v29
                                                                        var v30 int8 = (v20 | 71)
                                                                        var v31 uint64 = v1
                                                                        _ = v29
                                                                        _ = v30
                                                                        _ = v31
                                                                    default:
                                                                        if v21 {
                                                                            fmt.Println(fi18, v1, v19)
                                                                            var v32 int8 = v20
                                                                            v4 += (v4 + 701)
                                                                            _ = v32
                                                                        }
                                                                    }
                                                                    v1 = v1
                                                                } else {
                                                                    if ((!v21) && v21) {
                                                                        v22 -= v6
                                                                        switch (fi18 + 4294967295) {
                                                                        case 1:
                                                                            switch v19 {
                                                                            case "hello":
                                                                                v22 -= 269
                                                                            case "\n\t":
                                                                                v33 := GenS2{F0: (v4 | 2147483648), F1: 128}
                                                                                _ = v33
                                                                            }
                                                                        case 128:
                                                                            if false {
                                                                                switch 4294967295 {
                                                                                case 42:
                                                                                    v34 := GenS2{F0: (^(v4)), F1: 0}
                                                                                    _ = v34
                                                                                case 1:
                                                                                    v3[v19] = v19
                                                                                }
                                                                            } else {
                                                                                if v21 {
                                                                                    v35 := GenS3{F0: (v20 & -96)}
                                                                                    _ = v35
                                                                                }
                                                                            }
                                                                        }
                                                                        var v36 uint = 890
                                                                        _ = v36
                                                                    } else {
                                                                        var v37 float32 = 0.1
                                                                        _ = v37
                                                                    }
                                                                }
                                                            case -2147483648:
                                                                v38 := GenS1{F0: (^(v1)), F1: 0}
                                                                _ = v38
                                                            }
                                                            var v39 float32 = 0.0
                                                            var v40 uint16 = 424
                                                            switch v19 {
                                                            case "hello":
                                                                fmt.Println(v19, v20)
                                                            }
                                                            v41 := GenS3{F0: (v20 + (v20 / ((int8(v1)) | 1)))}
                                                            _ = v39
                                                            _ = v40
                                                            _ = v41
                                                        }
                                                    }
                                                    _ = v22
                                                }
                                                _ = v19
                                                _ = v20
                                                _ = v21
                                            }
                                        }
                                        _ = fi18
                                    }
                                }
                            case "\n\t":
                                v42 := GenS4{F0: -932, F1: "a", F2: 236}
                                _ = v42
                            }
                        }
                        fmt.Println(fi7, v1, fi7)
                    }
                    for fi43 := 0; fi43 < 15; fi43++ {
                        _ = len(v2)
                        _ = fi43
                    }
                    switch ("a" + "\n\t") {
                    case "\n\t":
                        v5 = v5
                    case "hello":
                        if true {
                            for fi44 := 0; fi44 < 8; fi44++ {
                                v45 := GenS3{F0: -1}
                                v1 = v1
                                _ = fi44
                                _ = v45
                            }
                        }
                    case "":
                        switch "a" {
                        case "":
                            for fi46 := 0; fi46 < 14; fi46++ {
                                v4 = 4294967295
                                if (false && (true || false)) {
                                    switch "a" {
                                    case "a":
                                        var v47 uint = 1
                                        _ = v47
                                    case "\n\t":
                                        if (true || (!(true || true))) {
                                            var v48 uint = 2
                                            if (!false) {
                                                var v49 int = v5
                                                if (!false) {
                                                    v48 += (v48 - 550)
                                                    v50 := GenS5{F0: "hello", F1: -1}
                                                    _ = v50
                                                }
                                                switch "" {
                                                case "hello":
                                                    v48 = 70
                                                case "foo bar":
                                                    v51 := GenS1{F0: 421, F1: v1}
                                                    _ = v51
                                                case "":
                                                    var v52 string = ("a" + ("foo bar" + ""))
                                                    _ = v52
                                                }
                                                _ = v49
                                            } else {
                                                fmt.Println(v5)
                                                v5 -= 32768
                                                switch "" {
                                                case "\n\t":
                                                    switch (fi7 >> ((uint(v4)) & 63)) {
                                                    case -536:
                                                        v4 = v4
                                                    case 4294967295:
                                                        fmt.Println(v5, fi7)
                                                    }
                                                default:
                                                    fmt.Println(v1, fi7)
                                                }
                                            }
                                            v53 := GenS6{F0: (v48 << ((v48) & 63))}
                                            _ = v48
                                            _ = v53
                                        }
                                    case "foo bar":
                                        v4 += (v4 | v4)
                                    default:
                                        v5 -= (fi7 << ((2) & 63))
                                    }
                                }
                                var v54 uint16 = 2
                                _ = fi46
                                _ = v54
                            }
                        }
                    default:
                        for fi55 := 0; fi55 < 20; fi55++ {
                            switch (fi55 % ((fi55) | 1)) {
                            case 42:
                                delete(v3, "")
                            case -497:
                                v5 += v5
                                var v56 uint32 = v4
                                switch fi7 {
                                case -1:
                                    var v57 uint64 = (v1 ^ v1)
                                    _ = v57
                                case -9223372036854775808:
                                    if false {
                                        v58 := GenS3{F0: -128}
                                        switch "hello" {
                                        case "foo bar":
                                            v5 += (fi7 ^ fi7)
                                        case "":
                                            v1 = v1
                                        case "a":
                                            if true {
                                                _ = len(v3)
                                                v59 := GenS2{F0: (v4 & 4294967295), F1: 192}
                                                _ = len(v3)
                                                _ = v59
                                            }
                                        default:
                                            var v60 int16 = 1000
                                            _ = v60
                                        }
                                        v6 = 0
                                        _ = v58
                                    }
                                case -826:
                                    switch (v5 * (fi7 / ((v5) | 1))) {
                                    case -802:
                                        v61 := GenS7{F0: v5}
                                        _ = v61
                                    case 1:
                                        if true {
                                            var v62 int32 = 42
                                            _ = v62
                                        } else {
                                            switch (fi7 % ((32768) | 1)) {
                                            case 42:
                                                v63 := GenS4{F0: -477, F1: ("hello" + "foo bar"), F2: 234}
                                                _ = v63
                                            case 9223372036854775807:
                                                v56 += (^(v4))
                                            case 89:
                                                if ((v56 >> ((18446744073709551615) & 63)) >= v4) {
                                                    switch (v5 & fi55) {
                                                    case 2147483647:
                                                        v64 := GenS1{F0: v1, F1: (v1 | v1)}
                                                        _ = v64
                                                    case 32768:
                                                        if true {
                                                            v65 := GenS8{F0: 32767, F1: 69, F2: (fi7 << ((42) & 63))}
                                                            if false {
                                                                if ((1.5 * (((1e10 / float64(3.14159)) * 1.5) + -2.25)) > (0.1 * (3.14159 - 1e10))) {
                                                                    _ = len(v3)
                                                                    v4 = (v4 % ((v4) | 1))
                                                                    fmt.Println(v56)
                                                                }
                                                            }
                                                            _ = v65
                                                        }
                                                    default:
                                                        switch (v5 << ((0) & 63)) {
                                                        case 127:
                                                            if (!false) {
                                                                if (32767 == 534) {
                                                                    v1 -= v1
                                                                    var v66 string = (("hello" + ("hello" + "a")) + "hello")
                                                                    _ = v66
                                                                }
                                                                var v67 float64 = float64(1.5)
                                                                _ = v67
                                                            } else {
                                                                var v68 uint64 = (v1 << ((42) & 63))
                                                                switch "hello" {
                                                                case "":
                                                                    v69 := GenS5{F0: "foo bar", F1: -32768}
                                                                    if true {
                                                                        switch (("foo bar" + ("\n\t" + ("\n\t" + "\n\t"))) + "foo bar") {
                                                                        case "\n\t":
                                                                            var v70 float64 = 3.14159
                                                                            v68 -= 399
                                                                            _ = v70
                                                                        case "foo bar":
                                                                            var v71 string = "\n\t"
                                                                            _ = v71
                                                                        case "hello":
                                                                            v72 := GenS5{F0: ("\n\t" + "foo bar"), F1: -616}
                                                                            _ = v72
                                                                        }
                                                                    }
                                                                    var v73 uint32 = v56
                                                                    var v74 uint64 = 805
                                                                    v4 -= 42
                                                                    _ = v69
                                                                    _ = v73
                                                                    _ = v74
                                                                default:
                                                                    switch "\n\t" {
                                                                    case "foo bar":
                                                                        fmt.Println(v1, v5, v6)
                                                                    case "\n\t":
                                                                        v6 -= v6
                                                                    case "a":
                                                                        fmt.Println(v5, v1, fi55)
                                                                    default:
                                                                        v6 = (v6 + (v6 >> ((4294967295) & 63)))
                                                                    }
                                                                }
                                                                switch ("a" + ("a" + ("foo bar" + "hello"))) {
                                                                case "hello":
                                                                    var v75 bool = true
                                                                    _ = v75
                                                                case "":
                                                                    fmt.Println(fi55, v4, v6)
                                                                default:
                                                                    v76 := GenS4{F0: 727, F1: ("a" + ("" + ("\n\t" + "a"))), F2: 227}
                                                                    _ = v76
                                                                }
                                                                _ = v68
                                                            }
                                                        case 42:
                                                            v77 := GenS6{F0: 0}
                                                            _ = v77
                                                        }
                                                    }
                                                    var v78 int64 = (^(v6))
                                                    delete(v3, "a")
                                                    _ = v78
                                                } else {
                                                    if (false || false) {
                                                        var v79 uint = 255
                                                        fmt.Println(v4)
                                                        var v80 int32 = 2147483647
                                                        _ = v79
                                                        _ = v80
                                                    } else {
                                                        _ = len(v3)
                                                        var v81 bool = (uint8(v4) == 87)
                                                        var v82 int = 4294967295
                                                        _ = v81
                                                        _ = v82
                                                    }
                                                    switch (fi55 + 128) {
                                                    case 32767:
                                                        v5 += fi7
                                                    case 754:
                                                        var v83 string = ("\n\t" + ("hello" + "hello"))
                                                        _ = v83
                                                    case -1:
                                                        switch "a" {
                                                        case "a":
                                                            if false {
                                                                var v84 int8 = 42
                                                                _ = v84
                                                            } else {
                                                                v6 += v6
                                                                switch "" {
                                                                case "a":
                                                                    var v85 int32 = -181
                                                                    _ = v85
                                                                case "":
                                                                    var v86 uint32 = v4
                                                                    _ = v86
                                                                default:
                                                                    fmt.Println(v6, fi55, v5)
                                                                }
                                                            }
                                                        }
                                                    default:
                                                        fmt.Println(fi7)
                                                    }
                                                }
                                            default:
                                                if len(v2) > 0 {
                                                    v2[0] = fi7
                                                }
                                            }
                                        }
                                    case -964:
                                        if false {
                                            v1 += v1
                                            v1 = (v1 - (^(v1)))
                                            fmt.Println(v4)
                                        }
                                    default:
                                        if true {
                                            v2 = append(v2, (^(fi7)))
                                        }
                                    }
                                    v56 = 512
                                    switch v5 {
                                    case 309:
                                        if ((!true) && (42 > -737)) {
                                            v4 += (^(v4))
                                            switch 42 {
                                            case 0:
                                                switch (fi7 - (^(fi7))) {
                                                case 4294967295:
                                                    var v87 uint8 = 210
                                                    _ = v87
                                                case 32767:
                                                    v88 := GenS1{F0: v1, F1: (v1 >> ((4294967295) & 63))}
                                                    _ = v88
                                                case -780:
                                                    if true {
                                                        if (!true) {
                                                            if (false && (1 == 32767)) {
                                                                v89 := GenS9{F0: 255, F1: v5}
                                                                v3["\n\t"] = "foo bar"
                                                                v56 = v4
                                                                _ = v89
                                                            } else {
                                                                v5 += (fi55 % ((fi55) | 1))
                                                                v90 := GenS8{F0: 405, F1: 2, F2: (-(fi7))}
                                                                if (((false && (!false)) || true) || (!true)) {
                                                                    if (-2.25 < 0.0) {
                                                                        v4 -= (v56 & (v56 * v56))
                                                                        v6 += (v6 & v6)
                                                                        _ = len(v3)
                                                                    } else {
                                                                        v1 -= v1
                                                                        _ = len(v3)
                                                                    }
                                                                    var v91 uint8 = 1
                                                                    if false {
                                                                        if (uint8(v1) == v91) {
                                                                            var v92 float64 = ((3.14159 * (-2.25 + 0.1)) + 0.1)
                                                                            _ = v92
                                                                        }
                                                                        v4 -= 42
                                                                        v93 := GenS9{F0: uint8(v56), F1: (fi55 - (fi55 / (((fi55 % ((1) | 1))) | 1)))}
                                                                        _ = v93
                                                                    }
                                                                    if len(v2) >= 2 {
                                                                        v2 = v2[1:len(v2)]
                                                                    }
                                                                    var v94 uint8 = 42
                                                                    _ = v91
                                                                    _ = v94
                                                                }
                                                                _ = v90
                                                            }
                                                            v6 -= 0
                                                        } else {
                                                            if true {
                                                                v6 += v6
                                                                v95 := v3["foo bar"]
                                                                fmt.Println(fi55, v95)
                                                                _ = v95
                                                            }
                                                            v4 -= v4
                                                        }
                                                    } else {
                                                        var v96 int16 = -1
                                                        if (false && true) {
                                                            var v97 int = -354
                                                            delete(v3, "hello")
                                                            _ = v97
                                                        }
                                                        _ = v96
                                                    }
                                                }
                                            case 127:
                                                if (1e10 > float64(0.1)) {
                                                    v1 = (v1 & v1)
                                                    v98 := v3[(("a" + "\n\t") + ("" + ""))]
                                                    var v99 string = ((("foo bar" + "\n\t") + (v98 + "a")) + v98)
                                                    _ = v98
                                                    _ = v99
                                                } else {
                                                    v1 -= (^(v1))
                                                }
                                            case 1:
                                                fmt.Println(v4, v4)
                                            default:
                                                if (0 != v1) {
                                                    v100 := GenS10{F0: 42, F1: -1}
                                                    v5 = fi7
                                                    var v101 uint = 793
                                                    _ = v100
                                                    _ = v101
                                                }
                                            }
                                            fmt.Println(v56, v1, v56)
                                        } else {
                                            switch "foo bar" {
                                            case "":
                                                v102 := GenS11{F0: "foo bar", F1: v5, F2: -953}
                                                _ = v102
                                            }
                                            var v103 uint16 = 523
                                            _ = v103
                                        }
                                    case -522:
                                        fmt.Println(v5)
                                    case -2147483648:
                                        fmt.Println(v6, v5)
                                    default:
                                        fmt.Println(fi55)
                                    }
                                    delete(v3, "hello")
                                    v104 := GenS11{F0: "hello", F1: 9223372036854775807, F2: (^(v6))}
                                    v105 := GenS12{F0: 255, F1: "hello", F2: "\n\t"}
                                    _ = v104
                                    _ = v105
                                }
                                fmt.Println(v56)
                                v56 -= (v56 / ((v4) | 1))
                                _ = v56
                            case 997:
                                if ((((1e10 - 0.1) + -2.25) * 1e10) >= ((-2.25 / 0.1) + 3.14159)) {
                                    v4 -= uint32(v1)
                                    var v106 uint8 = 128
                                    var v107 uint = 829
                                    switch fi55 {
                                    case -204:
                                        if false {
                                            v107 = v107
                                        } else {
                                            fmt.Println(v5, fi55)
                                            fmt.Println(fi55, v6)
                                            switch (v5 & (fi55 % ((int(v1)) | 1))) {
                                            case 32767:
                                                v108 := GenS13{F0: (v106 ^ uint8(v107)), F1: (^(v106))}
                                                _ = v108
                                            }
                                        }
                                    case 42:
                                        if true {
                                            if true {
                                                if ((((v4 >= v4) || false) && (!("foo bar" != "\n\t"))) && false) {
                                                    var v109 uint64 = 18446744073709551615
                                                    var v110 uint = (^(v107))
                                                    _ = v109
                                                    _ = v110
                                                } else {
                                                    var v111 uint16 = 32767
                                                    _ = v111
                                                }
                                                var v112 int8 = 100
                                                _ = v112
                                            } else {
                                                v1 += 336
                                                _ = len(v3)
                                                var v113 bool = true
                                                _ = v113
                                            }
                                            v114 := GenS7{F0: fi7}
                                            v4 += (v4 ^ (v4 & v4))
                                            _ = v114
                                        } else {
                                            fmt.Println(fi7, fi7, v6)
                                            if false {
                                                var v115 float32 = ((0.1 + 1e10) - 3.14159)
                                                switch "\n\t" {
                                                case "\n\t":
                                                    v5 = fi55
                                                case "hello":
                                                    v116 := GenS5{F0: "\n\t", F1: 42}
                                                    _ = v116
                                                case "foo bar":
                                                    fmt.Println(v107)
                                                default:
                                                    var v117 int16 = -1
                                                    _ = v117
                                                }
                                                fmt.Println(v6, v1, v6)
                                                _ = v115
                                            } else {
                                                var v118 int8 = -71
                                                v5 = 42
                                                _ = v118
                                            }
                                            switch -9223372036854775808 {
                                            case 127:
                                                v3[""] = ""
                                                var v119 bool = ((fi7 >> (((^(v107))) & 63)) > (v5 % (((^(fi55))) | 1)))
                                                _ = v119
                                            }
                                        }
                                    case 128:
                                        v6 = v6
                                    default:
                                        v3[("hello" + "")] = ""
                                    }
                                    _ = v106
                                    _ = v107
                                }
                            }
                            _ = fi55
                        }
                    }
                }
            case "":
                v120 := GenS14{F0: 105, F1: 353}
                _ = v120
            case "\n\t":
                fmt.Println(v1, fi7)
            default:
                fmt.Println(fi7, v1, v1)
            }
            v121 := GenS11{F0: "", F1: fi7, F2: (^(v6))}
            _ = fi7
            _ = v121
        }
        switch -733 {
        case 704:
            v122 := GenS4{F0: 2147483647, F1: "a", F2: 119}
            _ = v122
        case 1:
            var v123 int16 = -14
            _ = v123
        case -203:
            if (!(0.1 != float64(float32(3.14159)))) {
                v4 = (v4 / ((v4) | 1))
                v4 = 2147483648
                var v124 int8 = -17
                _ = v124
            } else {
                var v125 int64 = v6
                v126 := GenS14{F0: 0, F1: v6}
                for fi127 := 0; fi127 < 2; fi127++ {
                    var v128 uint64 = (^(v1))
                    v5 = 255
                    _ = fi127
                    _ = v128
                }
                _ = v125
                _ = v126
            }
        default:
            for fi129 := 0; fi129 < 16; fi129++ {
                v4 = v4
                _ = fi129
            }
        }
        v130 := GenS15{F0: (v1 << ((894) & 63)), F1: -104, F2: v6}
        _ = v130
    }
    var v131 uint8 = 42
    if (1e10 == ((3.14159 - 3.14159) * (((0.1 * 1e10) + (((float64((3.14159 * 0.1)) * ((1.5 * (-2.25 / (1e10 / float64(1.5)))) / 1.0)) - 3.14159) - 3.14159)) * 0.1))) {
        v132 := GenS16{F0: 42, F1: "foo bar", F2: 128}
        var v133 int = (v5 * (^(v5)))
        _ = v132
        _ = v133
    } else {
        fmt.Println(v4, v131)
    }
    v134 := GenS6{F0: 65535}
    v135 := GenS17{F0: 16}
    fmt.Println(v6)
    var v136 float32 = 3.14159
    if true {
        v4 -= (v4 | v4)
        v5 -= (v5 ^ 128)
        fmt.Println(v5)
    }
    v131 += v131
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi137 := 0; fi137 < 18; fi137++ {
        if ((^(v5)) != (v5 << ((954) & 63))) {
            delete(v3, "\n\t")
            v138 := GenS15{F0: v1, F1: 0, F2: v6}
            _ = v138
        } else {
            v139 := GenS5{F0: "\n\t", F1: -810}
            _ = v139
        }
        for fi140 := 0; fi140 < 6; fi140++ {
            if false {
                var v141 uint = 65535
                _ = v141
            } else {
                v142 := GenS9{F0: (^(v131)), F1: fi140}
                v143 := GenS3{F0: 0}
                fmt.Println(v4, fi140, v136)
                _ = v142
                _ = v143
            }
            v5 = (fi137 ^ (v5 * (-(v5))))
            _ = fi140
        }
        var v144 bool = (-9223372036854775808 == 0)
        _ = fi137
        _ = v144
    }
    fmt.Println(v5, v5)
    var v145 int = 128
    if (((true || true) || true) && (!(255 == 42))) {
        fmt.Println(v131, v145, v131)
    } else {
        switch "foo bar" {
        case "":
            v145 += v145
        case "hello":
            if (32767 < 42) {
                v1 = (v1 % (((v1 ^ (v1 & v1))) | 1))
                delete(v3, "a")
                var v146 int = v145
                _ = v146
            } else {
                v4 += 895
                if false {
                    v147 := GenS4{F0: 2147483647, F1: "foo bar", F2: (v131 % (((v131 - uint8(v1))) | 1))}
                    delete(v3, ("\n\t" + "foo bar"))
                    if false {
                        v131 += 0
                        v5 -= v5
                    } else {
                        v3[("foo bar" + "hello")] = "a"
                        var v148 float32 = (0.0 * 0.1)
                        _ = v148
                    }
                    _ = v147
                } else {
                    switch v145 {
                    case 9223372036854775807:
                        v149 := GenS1{F0: uint64(v131), F1: (v1 - v1)}
                        _ = v149
                    }
                }
            }
        }
        var v150 uint32 = 628
        switch v5 {
        case 1:
            if (v136 <= v136) {
                v136 = 1e10
            } else {
                v136 += v136
                fmt.Println(v136)
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            }
        case 4294967295:
            fmt.Println(v131, v131)
        case 623:
            if ((1e10 == (1e10 + (-2.25 * 3.14159))) || (!(!((^(v131)) <= v131)))) {
                for fi151 := 0; fi151 < 16; fi151++ {
                    var v152 uint16 = 152
                    for fi153 := 0; fi153 < 19; fi153++ {
                        if ((!false) && (-524 == (^(v6)))) {
                            if true {
                                switch int(v1) {
                                case 4294967295:
                                    v154 := GenS2{F0: 2147483648, F1: v131}
                                    _ = v154
                                case -1:
                                    v4 -= 2
                                }
                                if len(v2) > 0 {
                                    v2[0] = (fi153 & int(v131))
                                }
                                var v155 uint8 = (v131 + v131)
                                v136 = v136
                                v6 += v6
                                _ = v155
                            }
                            var v156 int8 = 127
                            var v157 int32 = -139
                            if true {
                                v158 := GenS13{F0: v131, F1: 128}
                                _ = v158
                            } else {
                                fmt.Println(v145, v152, v145)
                            }
                            _ = v156
                            _ = v157
                        } else {
                            var v159 uint16 = (v152 | v152)
                            _ = v159
                        }
                        switch "foo bar" {
                        case "":
                            v160 := GenS18{F0: v4, F1: -1}
                            _ = v160
                        case "hello":
                            v131 = 28
                        case "a":
                            var v161 uint8 = (v131 * v131)
                            v136 += (v136 * -2.25)
                            _ = v161
                        }
                        var v162 uint64 = (v1 + 477)
                        _ = fi153
                        _ = v162
                    }
                    _ = fi151
                    _ = v152
                }
                var v163 uint16 = 42
                _ = v163
            }
        default:
            var v164 int16 = 0
            _ = v164
        }
        _ = v150
    }
    v136 = (v136 / (((v136 + v136) + float32(1.5)) + (v136 * v136)))
    fmt.Println("v1:", v1)
    fmt.Println("v131:", v131)
    fmt.Println("v136:", v136)
    fmt.Println("v145:", v145)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v131
    _ = v134
    _ = v135
    _ = v136
    _ = v145
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
