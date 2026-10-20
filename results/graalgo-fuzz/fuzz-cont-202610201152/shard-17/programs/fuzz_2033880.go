package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string; F2 int16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int16; F1 string }
type GenS4 struct { F0 int32; F1 uint32 }
type GenS5 struct { F0 uint64; F1 uint16 }
type GenS6 struct { F0 int8; F1 int32; F2 int16 }
type GenS7 struct { F0 uint8; F1 int32; F2 string }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 uint8; F1 string }
type GenS11 struct { F0 uint8 }
type GenS12 struct { F0 bool; F1 string; F2 uint8 }
type GenS13 struct { F0 uint32 }
type GenS14 struct { F0 int8 }
type GenS15 struct { F0 uint32 }
type GenS16 struct { F0 int }

func main() {
    v1 := make(map[string]bool)
    v2 := []string{}
    v3 := make(map[int]int)
    var v4 uint8 = 1
    var v5 float64 = (3.14159 * (float64(0.1) * 1e10))
    var v6 int16 = 1000
    v5 -= (v5 / 1.0)
    var v7 float64 = (v5 + v5)
    for fi8 := 0; fi8 < 5; fi8++ {
        fmt.Println(v5, v7, v5)
        switch (fi8 + fi8) {
        case -9223372036854775808:
            switch fi8 {
            case 164:
                for fi9 := 0; fi9 < 13; fi9++ {
                    if false {
                        v10 := GenS1{F0: 65535, F1: (("a" + ("" + "a")) + ("foo bar" + "")), F2: v6}
                        v11 := v1["a"]
                        _ = v10
                        _ = v11
                    }
                    _ = fi9
                }
            case 127:
                switch (("" + "") + "\n\t") {
                case "\n\t":
                    v7 = 1.5
                case "hello":
                    var v12 int16 = (v6 | (v6 & (v6 & 32767)))
                    _ = v12
                default:
                    var v13 uint = 1
                    _ = v13
                }
            case 0:
                for fi14 := 0; fi14 < 2; fi14++ {
                    if ((3.14159 < -2.25) || true) {
                        v4 = v4
                        if false {
                            if false {
                                v6 = v6
                                var v15 bool = (0 > -90)
                                v16 := GenS2{F0: 4294967295}
                                _ = v15
                                _ = v16
                            } else {
                                var v17 int8 = -71
                                _ = len(v1)
                                if true {
                                    v6 = -32768
                                }
                                fmt.Println(fi8, v6, v5)
                                _ = v17
                            }
                            switch "a" {
                            case "a":
                                fmt.Println(fi8, fi8)
                            default:
                                var v18 int = -9223372036854775808
                                _ = v18
                            }
                            if (32768 != (fi14 + 4294967295)) {
                                delete(v1, "")
                                v19 := GenS3{F0: v6, F1: "foo bar"}
                                switch (("foo bar" + "\n\t") + ("hello" + "\n\t")) {
                                case "\n\t":
                                    v20 := GenS2{F0: 1}
                                    _ = v20
                                case "a":
                                    var v21 uint64 = 1
                                    _ = v21
                                default:
                                    fmt.Println(v4, fi14, fi14)
                                }
                                _ = v19
                            } else {
                                v22 := v3[0]
                                v6 = (v6 | v6)
                                switch ("\n\t" + ("a" + "")) {
                                case "":
                                    var v23 int = fi14
                                    _ = v23
                                case "foo bar":
                                    v5 = float64(-2.25)
                                case "a":
                                    if (true && true) {
                                        if (uint(v4) < 42) {
                                            v22 -= fi8
                                            if len(v2) > 0 {
                                                v2[0] = "foo bar"
                                            }
                                        } else {
                                            _ = len(v2)
                                        }
                                        switch ("\n\t" + ("foo bar" + (("hello" + ("a" + "")) + "a"))) {
                                        case "":
                                            switch fi14 {
                                            case -9223372036854775808:
                                                var v24 int16 = v6
                                                _ = v24
                                            }
                                        case "foo bar":
                                            var v25 int16 = int16(v4)
                                            _ = v25
                                        }
                                        if (((true || (false && false)) || (true || false)) || (42 > 32767)) {
                                            v26 := GenS4{F0: 42, F1: 2}
                                            var v27 uint32 = 42
                                            var v28 uint32 = v27
                                            _ = v26
                                            _ = v27
                                            _ = v28
                                        }
                                    }
                                default:
                                    var v29 int32 = 969
                                    _ = v29
                                }
                                _ = v22
                            }
                        }
                    }
                    _ = fi14
                }
            }
        case -695:
            var v30 float32 = 0.0
            _ = v30
        }
        var v31 float64 = -2.25
        _ = fi8
        _ = v31
    }
    v4 -= v4
    v2 = append(v2, (("hello" + ("\n\t" + "hello")) + "foo bar"))
    var v32 uint64 = 0
    v33 := GenS4{F0: 42, F1: 882}
    v34 := GenS5{F0: v32, F1: 667}
    switch 9223372036854775807 {
    case 493:
        v6 += v6
    case 4294967295:
        for fi35 := 0; fi35 < 7; fi35++ {
            if true {
                v32 = (v32 + 1)
            }
            if (((fi35 + fi35) >= fi35) || true) {
                if (true && (42 != 65535)) {
                    var v36 int64 = 42
                    v5 = 1e10
                    if (true && true) {
                        v4 += (v4 - 42)
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    } else {
                        if (true && false) {
                            for fi37 := 0; fi37 < 5; fi37++ {
                                var v38 bool = false
                                switch (fi35 << ((255) & 63)) {
                                case 32767:
                                    if v38 {
                                        v36 = (v36 ^ (v36 << ((4294967295) & 63)))
                                        if true {
                                            v39 := GenS4{F0: 1, F1: 0}
                                            var v40 int = -9223372036854775808
                                            var v41 int8 = int8(v40)
                                            _ = v39
                                            _ = v40
                                            _ = v41
                                        } else {
                                            if (!(!(!v38))) {
                                                v6 = v6
                                            } else {
                                                fmt.Println(fi37, v5)
                                                if false {
                                                    fmt.Println(v4, v32, v36)
                                                    switch (^(fi37)) {
                                                    case 42:
                                                        var v42 uint8 = v4
                                                        _ = v42
                                                    }
                                                } else {
                                                    v5 += (v7 + 0.1)
                                                }
                                                delete(v1, "hello")
                                                switch (fi35 / (((fi37 - (fi37 - fi35))) | 1)) {
                                                case 127:
                                                    if (-2147483648 < 705) {
                                                        switch fi35 {
                                                        case -1:
                                                            switch "foo bar" {
                                                            case "\n\t":
                                                                var v43 float32 = -2.25
                                                                _ = v43
                                                            default:
                                                                if (!v38) {
                                                                    v38 = false
                                                                }
                                                            }
                                                        case 128:
                                                            var v44 float32 = 3.14159
                                                            _ = v44
                                                        case 32768:
                                                            v45 := GenS1{F0: 42, F1: "hello", F2: v6}
                                                            _ = v45
                                                        }
                                                    } else {
                                                        switch (fi35 * (fi35 | fi35)) {
                                                        case 42:
                                                            v5 += 0.0
                                                        }
                                                        v46 := GenS6{F0: 1, F1: 0, F2: 0}
                                                        _ = v46
                                                    }
                                                default:
                                                    if v38 {
                                                        var v47 int8 = -1
                                                        var v48 bool = false
                                                        var v49 uint64 = v32
                                                        _ = v47
                                                        _ = v48
                                                        _ = v49
                                                    } else {
                                                        if len(v2) >= 2 {
                                                            v2 = v2[1:len(v2)]
                                                        }
                                                    }
                                                }
                                                var v50 uint64 = v32
                                                var v51 int64 = v36
                                                _ = v50
                                                _ = v51
                                            }
                                            v52 := GenS4{F0: 0, F1: 65535}
                                            v36 = v36
                                            if (int16(v36) >= -32768) {
                                                if len(v2) > 0 {
                                                    v2[0] = ""
                                                }
                                                fmt.Println(fi35, v36, v38)
                                                v5 = (((float64(3.14159) * (v5 / 3.14159)) * 1e10) + ((v5 + (0.0 - (float64(((0.1 + 3.14159) * 0.0)) * ((1.5 + 0.0) + v5)))) / v5))
                                            }
                                            if v38 {
                                                if v38 {
                                                    v53 := GenS3{F0: (v6 ^ v6), F1: "hello"}
                                                    _ = len(v3)
                                                    var v54 uint16 = 0
                                                    _ = v53
                                                    _ = v54
                                                } else {
                                                    v55 := GenS5{F0: 18446744073709551615, F1: 454}
                                                    switch fi35 {
                                                    case -2147483648:
                                                        if (0 != 1) {
                                                            v4 = v4
                                                            v56 := GenS5{F0: 668, F1: 0}
                                                            if (v38 && (835 != v32)) {
                                                                var v57 float64 = (v5 / (0.1 / float64(3.14159)))
                                                                var v58 float64 = v7
                                                                var v59 int = fi35
                                                                var v60 uint = 255
                                                                fmt.Println(v5)
                                                                _ = v57
                                                                _ = v58
                                                                _ = v59
                                                                _ = v60
                                                            } else {
                                                                fmt.Println(v6, v32)
                                                                var v61 uint8 = (^(v4))
                                                                _ = v61
                                                            }
                                                            _ = v56
                                                        } else {
                                                            switch "" {
                                                            case "foo bar":
                                                                if (v38 || v38) {
                                                                    var v62 int64 = (^(v36))
                                                                    _ = v62
                                                                }
                                                            default:
                                                                _ = len(v2)
                                                            }
                                                            _ = len(v2)
                                                        }
                                                        v6 = v6
                                                    default:
                                                        _ = len(v3)
                                                    }
                                                    _ = v55
                                                }
                                                var v63 uint16 = 0
                                                _ = v63
                                            }
                                            _ = v52
                                        }
                                    }
                                case 127:
                                    var v64 int32 = 42
                                    v4 = (^(v4))
                                    if v38 {
                                        v65 := GenS6{F0: 1, F1: v64, F2: 1000}
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        _ = v65
                                    }
                                    _ = v64
                                }
                                var v66 uint = 255
                                _ = len(v2)
                                var v67 float64 = 1e10
                                _ = fi37
                                _ = v38
                                _ = v66
                                _ = v67
                            }
                            for fi68 := 0; fi68 < 2; fi68++ {
                                var v69 bool = (!true)
                                switch "a" {
                                case "":
                                    v36 = (v36 >> ((42) & 63))
                                case "hello":
                                    v3[fi35] = (fi68 * fi35)
                                default:
                                    v70 := GenS7{F0: v4, F1: -955, F2: ((("foo bar" + ("foo bar" + ("a" + ("hello" + "a")))) + "") + "hello")}
                                    _ = v70
                                }
                                _ = fi68
                                _ = v69
                            }
                            var v71 string = "foo bar"
                            _ = v71
                        }
                        var v72 uint8 = 128
                        fmt.Println(v6)
                        _ = v72
                    }
                    _ = v36
                }
                var v73 uint32 = 807
                if (0 <= 0) {
                    if (42 == 32767) {
                        v74 := GenS5{F0: (v32 >> ((0) & 63)), F1: 223}
                        _ = v74
                    }
                    switch (-(fi35)) {
                    case 128:
                        if (-988 == -158) {
                            v32 += v32
                            v75 := GenS8{F0: (v73 / (((v73 | (v73 * 2147483648))) | 1))}
                            _ = v75
                        }
                    case 1:
                        for fi76 := 0; fi76 < 5; fi76++ {
                            fmt.Println(v32, fi35)
                            var v77 int16 = -32768
                            v6 = v77
                            _ = fi76
                            _ = v77
                        }
                    default:
                        v2 = append(v2, "hello")
                    }
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                } else {
                    if (-1 < -2147483648) {
                        switch (^(fi35)) {
                        case 366:
                            var v78 int8 = 42
                            _ = v78
                        case 4294967295:
                            delete(v3, fi35)
                        case -433:
                            v4 = v4
                        }
                        var v79 uint32 = uint32(v4)
                        _ = v79
                    }
                    v5 += (((float64((3.14159 * ((0.0 / (0.1 / (1.5 * 0.1))) / (-2.25 + (float32(v5) - (3.14159 / -2.25)))))) * (v5 - (((((v7 + 3.14159) * v5) - v7) + (1e10 * ((1.5 - (0.0 + v7)) / (3.14159 - v5)))) / (((1.5 + 1e10) - (((0.1 / (float64(3.14159) * (0.0 - ((0.0 - 3.14159) + ((((-2.25 + float64(((0.0 / 3.14159) / -2.25))) / (v7 + (v5 * v5))) * v5) / (v7 - (((v5 + (((v5 / (1.5 - v7)) + v5) * 0.0)) - (float64((1.5 - (-2.25 - (-2.25 * (-2.25 / 0.1))))) / ((((v7 / ((0.1 * 0.1) + v5)) / v7) + v5) * (v5 + (float64(1.5) * (v5 * v5)))))) + (float64(0.0) / (v7 - v7))))))))) + v7) + v7)) * (float64(float32((-2.25 - v5))) / v5))))) - v5) + v7)
                }
                _ = v73
            } else {
                v80 := GenS7{F0: (^(v4)), F1: -442, F2: ("a" + "")}
                for fi81 := 0; fi81 < 11; fi81++ {
                    if true {
                        fmt.Println(fi35)
                        if false {
                            if (v4 == (v4 >> ((734) & 63))) {
                                v3[fi35] = -9223372036854775808
                                v6 = (v6 ^ (v6 + v6))
                            }
                            if (-627 >= (v6 - 32767)) {
                                v82 := GenS9{F0: 486}
                                var v83 uint8 = 255
                                _ = v82
                                _ = v83
                            } else {
                                var v84 float64 = v5
                                _ = v84
                            }
                            v1["a"] = (!(!true))
                        } else {
                            v5 = (-2.25 / 1.5)
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            v4 = (v4 / ((uint8(v32)) | 1))
                            switch ((("a" + "") + ("a" + "")) + "foo bar") {
                            case "a":
                                var v85 int = fi35
                                _ = v85
                            case "\n\t":
                                v7 -= (float64((1.5 - (3.14159 - ((1.5 + float32(float64(1.5))) * ((float32((v5 + v5)) - (1e10 - (0.0 / 3.14159))) * 0.0))))) + float64(((1.5 - float32((0.1 / ((((-2.25 * -2.25) + 3.14159) - ((1e10 * (v5 / v7)) + v5)) / ((3.14159 * v7) / (v7 / (v7 + v7))))))) * 1e10)))
                            case "hello":
                                var v86 uint = 437
                                _ = v86
                            default:
                                fmt.Println(v6, fi81)
                            }
                            _ = len(v2)
                        }
                    }
                    fmt.Println(v5)
                    _ = fi81
                }
                _ = v80
            }
            _ = fi35
        }
    default:
        fmt.Println(v6, v32, v32)
    }
    switch "a" {
    case "foo bar":
        var v87 bool = true
        _ = v87
    case "\n\t":
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    v6 += 990
    if true {
        var v88 bool = (true && ((-128 < 127) || true))
        _ = v88
    }
    switch 817 {
    case 4294967295:
        var v89 int16 = (v6 >> ((739) & 63))
        _ = v89
    case 1:
        for fi90 := 0; fi90 < 19; fi90++ {
            var v91 int8 = -128
            if false {
                if (false && true) {
                    var v92 int16 = v6
                    v93 := GenS4{F0: 42, F1: 2}
                    if (-372 <= 0) {
                        _ = len(v1)
                        if true {
                            if (!true) {
                                v32 -= v32
                            }
                            v7 = (v5 * float64(0.0))
                            if false {
                                v94 := GenS6{F0: 34, F1: 42, F2: (v6 / ((v92) | 1))}
                                v6 = v6
                                var v95 int8 = (v91 ^ (^(v91)))
                                _ = v94
                                _ = v95
                            } else {
                                v2 = append(v2, "a")
                                fmt.Println(fi90)
                            }
                        } else {
                            v96 := GenS10{F0: 42, F1: "\n\t"}
                            switch "" {
                            case "foo bar":
                                v97 := GenS4{F0: 82, F1: 42}
                                _ = v97
                            case "\n\t":
                                fmt.Println(v32, v5)
                            default:
                                fmt.Println(v4, v92, v32)
                            }
                            switch "" {
                            case "foo bar":
                                v98 := GenS3{F0: 0, F1: "foo bar"}
                                _ = v98
                            default:
                                v6 = -32768
                            }
                            _ = v96
                        }
                    }
                    _ = v92
                    _ = v93
                } else {
                    var v99 int16 = v6
                    _ = v99
                }
            } else {
                var v100 bool = (v4 <= v4)
                _ = v100
            }
            _ = fi90
            _ = v91
        }
    case -1:
        v101 := GenS10{F0: v4, F1: (("a" + ("foo bar" + ("" + "a"))) + "a")}
        _ = v101
    }
    if ((!false) || false) {
        v32 = (v32 - (v32 / ((v32) | 1)))
        if (!true) {
            v102 := GenS9{F0: 65535}
            v32 += 629
            for fi103 := 0; fi103 < 6; fi103++ {
                for fi104 := 0; fi104 < 4; fi104++ {
                    if ((("a" + "\n\t") > "\n\t") || (2 != 199)) {
                        var v105 int16 = (v6 >> ((65535) & 63))
                        if (v4 < 1) {
                            delete(v1, "hello")
                        } else {
                            v6 = (-(v6))
                            switch ("\n\t" + "hello") {
                            case "\n\t":
                                v5 = v5
                            case "hello":
                                if false {
                                    if (!true) {
                                        fmt.Println(fi104, v6)
                                        switch ("foo bar" + ("a" + "\n\t")) {
                                        case "foo bar":
                                            if (false || (-97 == 42)) {
                                                fmt.Println(v7, fi104)
                                                var v106 bool = (65535 != uint32(v32))
                                                var v107 int16 = (v6 - v105)
                                                v106 = (uint16(v32) == 32767)
                                                fmt.Println(v105, v32)
                                                _ = v106
                                                _ = v107
                                            } else {
                                                v4 += (v4 / ((v4) | 1))
                                                switch ("\n\t" + "foo bar") {
                                                case "":
                                                    var v108 int32 = 2147483647
                                                    _ = v108
                                                }
                                            }
                                        }
                                        var v109 int32 = 782
                                        _ = v109
                                    }
                                    v7 -= -2.25
                                    _ = len(v2)
                                    if len(v2) > 0 {
                                        v2[0] = "hello"
                                    }
                                    v105 -= (v6 >> ((531) & 63))
                                    switch fi104 {
                                    case -175:
                                        var v110 int64 = 1
                                        _ = v110
                                    case -1:
                                        if len(v2) > 0 {
                                            v2[0] = ""
                                        }
                                    default:
                                        v4 = v4
                                    }
                                }
                            case "a":
                                switch fi104 {
                                case -931:
                                    v6 = v105
                                case 671:
                                    fmt.Println(v5)
                                case 32767:
                                    var v111 int32 = -944
                                    _ = v111
                                default:
                                    var v112 int32 = 2147483647
                                    _ = v112
                                }
                            default:
                                v5 -= -2.25
                            }
                            switch (fi103 % ((-2147483648) | 1)) {
                            case 0:
                                v4 = v4
                            case 4294967295:
                                var v113 string = "hello"
                                _ = v113
                            }
                        }
                        var v114 int16 = (^(v6))
                        _ = v105
                        _ = v114
                    } else {
                        var v115 uint8 = v4
                        if len(v2) > 0 {
                            v2[0] = ("hello" + "\n\t")
                        }
                        if false {
                            switch -231 {
                            case 4294967295:
                                fmt.Println(fi104, fi104)
                            case -552:
                                if false {
                                    switch 32767 {
                                    case 2147483647:
                                        var v116 uint8 = (v4 << ((4294967295) & 63))
                                        _ = v116
                                    case 1:
                                        v117 := GenS5{F0: (v32 & (^(v32))), F1: 0}
                                        _ = v117
                                    case 42:
                                        fmt.Println(fi103)
                                    default:
                                        v7 += v7
                                    }
                                    var v118 int32 = 2147483647
                                    _ = v118
                                }
                            default:
                                if true {
                                    fmt.Println(fi104)
                                    v5 -= float64(float32(float64(-2.25)))
                                    v119 := GenS1{F0: 42, F1: "foo bar", F2: v6}
                                    _ = v119
                                } else {
                                    if false {
                                        v120 := GenS5{F0: 42, F1: 65535}
                                        _ = v120
                                    }
                                    switch ("" + "") {
                                    case "hello":
                                        v7 = 0.0
                                    case "a":
                                        var v121 uint64 = (v32 - v32)
                                        _ = v121
                                    default:
                                        v122 := GenS4{F0: -1, F1: 0}
                                        _ = v122
                                    }
                                    v5 = v7
                                }
                            }
                            switch 42 {
                            case -241:
                                v115 = (v4 % ((1) | 1))
                            case 128:
                                var v123 int64 = -875
                                _ = v123
                            case 9223372036854775807:
                                if true {
                                    v5 = v5
                                    var v124 int8 = 42
                                    if ((0.1 * -2.25) > (0.0 / 1e10)) {
                                        v115 = (v4 ^ (^(v4)))
                                        v125 := GenS2{F0: 4294967295}
                                        _ = v125
                                    } else {
                                        if ((^(v4)) != v4) {
                                            fmt.Println(v124)
                                            v126 := GenS3{F0: v6, F1: "a"}
                                            switch "hello" {
                                            case "":
                                                _ = len(v2)
                                            default:
                                                var v127 uint32 = 42
                                                _ = v127
                                            }
                                            _ = v126
                                        }
                                        fmt.Println(v115)
                                        if (((-(v6)) >= 969) || true) {
                                            switch (("foo bar" + ("foo bar" + ("hello" + "hello"))) + "") {
                                            case "a":
                                                switch fi103 {
                                                case -2147483648:
                                                    fmt.Println(v4, fi103, v32)
                                                case 127:
                                                    v115 = (v115 - v4)
                                                case -9223372036854775808:
                                                    v4 += 38
                                                default:
                                                    var v128 uint32 = 65535
                                                    _ = v128
                                                }
                                            case "\n\t":
                                                var v129 float32 = 3.14159
                                                _ = v129
                                            default:
                                                v3[(fi103 % ((fi103) | 1))] = 32767
                                            }
                                            v32 -= v32
                                        }
                                        if false {
                                            var v130 uint32 = 16
                                            if true {
                                                v6 -= int16(v4)
                                            }
                                            if (v115 >= v4) {
                                                var v131 uint64 = v32
                                                switch fi103 {
                                                case 2147483647:
                                                    v132 := GenS7{F0: v4, F1: 0, F2: ("a" + "foo bar")}
                                                    _ = v132
                                                case 0:
                                                    var v133 uint16 = 613
                                                    _ = v133
                                                }
                                                _ = v131
                                            } else {
                                                v4 += 42
                                            }
                                            _ = v130
                                        } else {
                                            if len(v2) > 0 {
                                                v2[0] = "hello"
                                            }
                                            var v134 string = ("a" + ("foo bar" + "hello"))
                                            _ = v134
                                        }
                                    }
                                    _ = v124
                                }
                            default:
                                var v135 int64 = 1
                                _ = v135
                            }
                            var v136 string = "\n\t"
                            _ = v136
                        } else {
                            if false {
                                v137 := GenS9{F0: 4294967295}
                                if len(v2) > 0 {
                                    v2[0] = "hello"
                                }
                                var v138 uint32 = 4294967295
                                _ = v137
                                _ = v138
                            }
                        }
                        _ = v115
                    }
                    _ = fi104
                }
                var v139 string = "a"
                if (v6 > (-(v6))) {
                    switch fi103 {
                    case -859:
                        v1[v139] = false
                    case 70:
                        var v140 uint8 = v4
                        _ = v140
                    case 42:
                        fmt.Println(v5, v139)
                    }
                    v141 := v1[("a" + (((v139 + ((((((v139 + v139) + v139) + ("foo bar" + ("foo bar" + "a"))) + ((v139 + v139) + "hello")) + (v139 + ((v139 + "a") + (v139 + ("" + v139))))) + (v139 + ("hello" + (v139 + v139))))) + v139) + ""))]
                    _ = v141
                }
                _ = fi103
                _ = v139
            }
            _ = v102
        } else {
            var v142 int64 = 42
            var v143 int64 = v142
            _ = v142
            _ = v143
        }
        var v144 uint = 255
        _ = v144
    } else {
        for fi145 := 0; fi145 < 17; fi145++ {
            if false {
                switch "hello" {
                case "a":
                    switch ("" + "") {
                    case "":
                        var v146 int = (fi145 / ((fi145) | 1))
                        _ = v146
                    case "a":
                        for fi147 := 0; fi147 < 4; fi147++ {
                            v7 = v5
                            v7 = v7
                            _ = fi147
                        }
                    case "hello":
                        switch (("\n\t" + "hello") + ("" + ("hello" + "foo bar"))) {
                        case "hello":
                            v32 = (^(v32))
                        case "":
                            for fi148 := 0; fi148 < 17; fi148++ {
                                if false {
                                    var v149 uint32 = 1
                                    fmt.Println(fi145)
                                    fmt.Println(fi148, v6, v149)
                                    _ = v149
                                } else {
                                    v1["hello"] = (42 <= 1)
                                    if false {
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        var v150 uint = 18446744073709551615
                                        v151 := GenS4{F0: -1, F1: 2147483647}
                                        _ = v150
                                        _ = v151
                                    }
                                    if ("hello" <= "a") {
                                        var v152 float64 = 1e10
                                        _ = v152
                                    } else {
                                        if ((0 <= 2147483647) || true) {
                                            if ((v32 * 42) != (v32 - (v32 + v32))) {
                                                switch "a" {
                                                case "a":
                                                    v32 = (v32 % (((v32 % (((v32 << ((65535) & 63))) | 1))) | 1))
                                                case "foo bar":
                                                    if (42 > v32) {
                                                        if true {
                                                            v4 = v4
                                                            if (!true) {
                                                                if len(v2) >= 2 {
                                                                    v2 = v2[1:len(v2)]
                                                                }
                                                                var v153 float32 = (0.0 / 3.14159)
                                                                _ = v153
                                                            } else {
                                                                switch 211 {
                                                                case -1:
                                                                    fmt.Println(v7, v6)
                                                                case 0:
                                                                    var v154 uint8 = 127
                                                                    _ = v154
                                                                case 173:
                                                                    switch int(v4) {
                                                                    case -9223372036854775808:
                                                                        var v155 uint16 = 2
                                                                        _ = v155
                                                                    case 32767:
                                                                        fmt.Println(v6, v32, v32)
                                                                    case 0:
                                                                        var v156 string = "foo bar"
                                                                        _ = v156
                                                                    default:
                                                                        var v157 uint32 = 1
                                                                        _ = v157
                                                                    }
                                                                default:
                                                                    v7 -= (v7 * float64(1e10))
                                                                }
                                                                _ = len(v1)
                                                            }
                                                        }
                                                        if ((0.0 * (3.14159 - 1e10)) != float32(v5)) {
                                                            var v158 int16 = v6
                                                            v7 += float64((0.1 - 0.0))
                                                            v159 := GenS11{F0: v4}
                                                            switch "hello" {
                                                            case "a":
                                                                v160 := GenS12{F0: false, F1: ("\n\t" + "\n\t"), F2: (v4 << ((18446744073709551615) & 63))}
                                                                _ = v160
                                                            case "":
                                                                var v161 float64 = v7
                                                                _ = v161
                                                            }
                                                            if false {
                                                                v162 := GenS3{F0: (v6 ^ (v6 | (-(v158)))), F1: ("\n\t" + "a")}
                                                                _ = v162
                                                            } else {
                                                                switch fi145 {
                                                                case -1:
                                                                    fmt.Println(v158)
                                                                default:
                                                                    v158 = v158
                                                                }
                                                                if (v158 > (v158 - 1)) {
                                                                    v163 := GenS13{F0: 65535}
                                                                    _ = v163
                                                                }
                                                                var v164 uint8 = v4
                                                                _ = v164
                                                            }
                                                            fmt.Println(v5)
                                                            if ((-1 != 0) && (v7 == (v5 - (1e10 * float64(1e10))))) {
                                                                v165 := GenS6{F0: 7, F1: 0, F2: (v158 + v158)}
                                                                var v166 uint16 = 2
                                                                _ = v165
                                                                _ = v166
                                                            } else {
                                                                var v167 float32 = (1.5 + (((1.5 * (0.1 / 0.1)) / 1e10) / 1e10))
                                                                _ = v167
                                                            }
                                                            var v168 int16 = (v158 ^ (-(v158)))
                                                            _ = v158
                                                            _ = v159
                                                            _ = v168
                                                        }
                                                    }
                                                    switch fi148 {
                                                    case 32768:
                                                        v169 := GenS14{F0: 1}
                                                        _ = v169
                                                    case -5:
                                                        v5 = v5
                                                    case 128:
                                                        var v170 uint32 = 856
                                                        _ = v170
                                                    }
                                                    if true {
                                                        v7 = 3.14159
                                                        var v171 string = "foo bar"
                                                        v172 := GenS1{F0: 255, F1: v171, F2: (v6 + v6)}
                                                        _ = v171
                                                        _ = v172
                                                    } else {
                                                        var v173 int = 0
                                                        var v174 bool = true
                                                        _ = v173
                                                        _ = v174
                                                    }
                                                    if (-888 != (^(fi145))) {
                                                        delete(v3, (fi148 - fi148))
                                                        v7 -= (v5 + v7)
                                                    } else {
                                                        v175 := GenS15{F0: 585}
                                                        var v176 int64 = -188
                                                        _ = v175
                                                        _ = v176
                                                    }
                                                    _ = len(v3)
                                                    v32 += v32
                                                    switch 128 {
                                                    case 886:
                                                        if false {
                                                            v4 = (v4 >> ((838) & 63))
                                                            var v177 int = int(v32)
                                                            var v178 float32 = (0.0 + 1.5)
                                                            var v179 int8 = 127
                                                            v6 -= (-(v6))
                                                            _ = v177
                                                            _ = v178
                                                            _ = v179
                                                        }
                                                    case -1:
                                                        var v180 int32 = 1
                                                        _ = v180
                                                    }
                                                    if (v6 < 42) {
                                                        var v181 uint32 = 2147483647
                                                        v182 := GenS16{F0: fi145}
                                                        var v183 uint64 = (v32 | v32)
                                                        _ = v181
                                                        _ = v182
                                                        _ = v183
                                                    }
                                                    v184 := GenS2{F0: 4294967295}
                                                    _ = len(v2)
                                                    _ = v184
                                                case "hello":
                                                    var v185 int32 = 0
                                                    _ = v185
                                                default:
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                }
                                                v5 = 1e10
                                            } else {
                                                var v186 uint8 = (v4 >> ((18446744073709551615) & 63))
                                                var v187 uint8 = (v186 << ((2) & 63))
                                                var v188 bool = true
                                                _ = v186
                                                _ = v187
                                                _ = v188
                                            }
                                            if (!(4294967295 < 825)) {
                                                var v189 int = fi145
                                                var v190 int16 = v6
                                                fmt.Println(v7)
                                                _ = v189
                                                _ = v190
                                            }
                                        } else {
                                            v6 = (v6 / ((-447) | 1))
                                        }
                                    }
                                    v191 := v1[""]
                                    var v192 uint16 = 42
                                    _ = v191
                                    _ = v192
                                }
                                v4 = v4
                                _ = fi148
                            }
                        }
                    }
                case "hello":
                    for fi193 := 0; fi193 < 4; fi193++ {
                        if false {
                            v2 = append(v2, "a")
                            if len(v2) > 0 {
                                v2[0] = "\n\t"
                            }
                            fmt.Println(fi193, v4)
                        } else {
                            var v194 uint32 = 2147483648
                            if ((fi145 / ((fi193) | 1)) != fi145) {
                                v5 = 0.0
                                var v195 int8 = -128
                                var v196 string = ("\n\t" + "hello")
                                _ = v195
                                _ = v196
                            } else {
                                v194 = v194
                                v197 := GenS8{F0: (v194 + 2147483648)}
                                v194 += v194
                                _ = v197
                            }
                            _ = v194
                        }
                        var v198 uint64 = 4294967295
                        var v199 float32 = 3.14159
                        _ = fi193
                        _ = v198
                        _ = v199
                    }
                case "\n\t":
                    _ = len(v3)
                default:
                    var v200 uint16 = 65535
                    _ = v200
                }
            }
            var v201 uint32 = 4294967295
            var v202 int32 = 1
            _ = fi145
            _ = v201
            _ = v202
        }
    }
    var v203 int8 = 127
    fmt.Println(v4, v7)
    for fi204 := 0; fi204 < 19; fi204++ {
        var v205 int16 = 0
        var v206 int64 = 9223372036854775807
        _ = fi204
        _ = v205
        _ = v206
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v203:", v203)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v203
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
