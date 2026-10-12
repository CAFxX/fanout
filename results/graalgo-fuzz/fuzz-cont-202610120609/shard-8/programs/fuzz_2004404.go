package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint16; F1 int8; F2 uint64 }
type GenS3 struct { F0 int16; F1 int64; F2 uint32 }
type GenS4 struct { F0 int; F1 string; F2 int32 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 int8; F1 int; F2 uint8 }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 string }
type GenS10 struct { F0 uint16; F1 uint16; F2 int32 }
type GenS11 struct { F0 int16; F1 uint64 }
type GenS12 struct { F0 string; F1 int32; F2 uint64 }
type GenS13 struct { F0 uint16 }
type GenS14 struct { F0 int32; F1 int16; F2 bool }
type GenS15 struct { F0 uint16; F1 int }

func main() {
    v1 := []string{"", "\n\t", ("foo bar" + (("" + ("foo bar" + ("a" + ("a" + "hello")))) + "a"))}
    v2 := []string{"", "hello", ""}
    var v3 int32 = -841
    var v4 uint16 = 65535
    var v5 float32 = (1e10 * (1e10 / 1e10))
    v6 := GenS1{F0: (v3 << ((255) & 63))}
    fmt.Println(v5, v3, v4)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if (65535 < 42) {
        v4 = (v4 << ((42) & 63))
    }
    if true {
        var v7 uint = 42
        for fi8 := 0; fi8 < 3; fi8++ {
            if (((0.0 * (0.0 - float64(v5))) != float64((-2.25 / v5))) && true) {
                if true {
                    for fi9 := 0; fi9 < 9; fi9++ {
                        switch 127 {
                        case 32768:
                            _ = len(v2)
                        default:
                            v7 -= 255
                        }
                        var v10 int32 = (v3 | (^(v3)))
                        switch 0 {
                        case 2147483647:
                            v11 := GenS2{F0: (v4 + v4), F1: -1, F2: 1}
                            _ = v11
                        case -715:
                            if (2 != 356) {
                                v12 := GenS2{F0: (v4 | (v4 & v4)), F1: 0, F2: 1}
                                v13 := GenS3{F0: -373, F1: -511, F2: 0}
                                _ = v12
                                _ = v13
                            } else {
                                v14 := GenS4{F0: (fi9 % (((fi8 + 42)) | 1)), F1: "", F2: -924}
                                v15 := GenS2{F0: 858, F1: -96, F2: 0}
                                var v16 int16 = 32767
                                _ = v14
                                _ = v15
                                _ = v16
                            }
                        default:
                            if true {
                                var v17 int64 = 42
                                if ("a" > "foo bar") {
                                    switch (fi9 | fi8) {
                                    case 173:
                                        if false {
                                            var v18 uint64 = 0
                                            v19 := GenS3{F0: -1, F1: v17, F2: 584}
                                            var v20 int32 = (v3 / ((int32(v7)) | 1))
                                            _ = v18
                                            _ = v19
                                            _ = v20
                                        } else {
                                            if false {
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                                var v21 int32 = (v3 / ((v10) | 1))
                                                var v22 int8 = 42
                                                _ = v21
                                                _ = v22
                                            } else {
                                                var v23 uint = v7
                                                _ = v23
                                            }
                                            switch fi9 {
                                            case 128:
                                                if (v4 <= 104) {
                                                    v24 := GenS5{F0: -201}
                                                    v25 := GenS6{F0: 463}
                                                    v3 = 829
                                                    var v26 float64 = ((((0.0 - 0.1) - (3.14159 * ((0.0 / float64(float32(3.14159))) - 1.5))) + 1.5) - 3.14159)
                                                    if ((!true) && false) {
                                                        v4 = (^(v4))
                                                        if (v7 <= (v7 % (((v7 ^ uint(v4))) | 1))) {
                                                            if (uint16(v7) <= (v4 << ((v7) & 63))) {
                                                                if true {
                                                                    _ = len(v2)
                                                                    fmt.Println(v7, fi8, v26)
                                                                    v2 = append(v2, "hello")
                                                                }
                                                                v17 = 42
                                                                if ((3.14159 / v5) < v5) {
                                                                    switch "\n\t" {
                                                                    case "hello":
                                                                        var v27 int32 = (v3 | (v10 - v3))
                                                                        _ = v27
                                                                    case "foo bar":
                                                                        if (((-(v17)) >= (-(v17))) && (true || ((v7 + (v7 & (v7 ^ (v7 - (v7 - v7))))) < 0))) {
                                                                            fmt.Println(v4, v17, v17)
                                                                        } else {
                                                                            v28 := GenS5{F0: (v10 & v10)}
                                                                            var v29 int8 = -117
                                                                            switch ("a" + "") {
                                                                            case "":
                                                                                _ = len(v2)
                                                                            case "foo bar":
                                                                                if true {
                                                                                    v30 := GenS7{F0: (^(v29)), F1: 1, F2: 2}
                                                                                    if true {
                                                                                        if ((v7 + (v7 * v7)) <= (v7 & 2)) {
                                                                                            var v31 int32 = (v10 + v10)
                                                                                            switch fi9 {
                                                                                            case -327:
                                                                                                if ((fi8 - fi9) >= (-(fi8))) {
                                                                                                    var v32 float32 = v5
                                                                                                    v29 = 123
                                                                                                    _ = v32
                                                                                                } else {
                                                                                                    switch ((("\n\t" + "a") + "hello") + "\n\t") {
                                                                                                    case "":
                                                                                                        v7 = v7
                                                                                                    case "foo bar":
                                                                                                        v17 += v17
                                                                                                    default:
                                                                                                        if (2147483647 > (fi8 * fi9)) {
                                                                                                            var v33 float32 = (float32((1.5 + v26)) / (((v5 + 0.1) - float32((0.1 + v26))) + ((1e10 / (float32((3.14159 - float64(v5))) / (v5 / v5))) * (0.1 * (-2.25 * -2.25)))))
                                                                                                            var v34 int64 = v17
                                                                                                            switch 337 {
                                                                                                            case 32768:
                                                                                                                fmt.Println(v7)
                                                                                                                switch (fi8 | (fi8 * (fi9 >> ((uint(v4)) & 63)))) {
                                                                                                                case -9223372036854775808:
                                                                                                                    v35 := GenS1{F0: v3}
                                                                                                                    _ = v35
                                                                                                                case -1:
                                                                                                                    _ = len(v2)
                                                                                                                case 32767:
                                                                                                                    v36 := GenS3{F0: 32767, F1: 1, F2: 1}
                                                                                                                    _ = v36
                                                                                                                }
                                                                                                            case 729:
                                                                                                                if ((v29 << ((v7) & 63)) > 83) {
                                                                                                                    if len(v1) > 0 {
                                                                                                                        v1[0] = ""
                                                                                                                    }
                                                                                                                    var v37 float32 = (1.5 * ((v33 - ((0.0 * (v33 + v33)) + v33)) - (1.5 - v33)))
                                                                                                                    _ = v37
                                                                                                                } else {
                                                                                                                    v38 := GenS8{F0: (v4 | v4)}
                                                                                                                    v7 = 0
                                                                                                                    _ = v38
                                                                                                                }
                                                                                                            case 127:
                                                                                                                v39 := GenS1{F0: v10}
                                                                                                                var v40 int16 = -838
                                                                                                                _ = v39
                                                                                                                _ = v40
                                                                                                            default:
                                                                                                                if false {
                                                                                                                    var v41 string = "foo bar"
                                                                                                                    fmt.Println(v7, v17, v4)
                                                                                                                    _ = v41
                                                                                                                }
                                                                                                            }
                                                                                                            fmt.Println(fi9)
                                                                                                            var v42 bool = (!(fi8 >= int(v7)))
                                                                                                            v42 = (v42 && v42)
                                                                                                            _ = v33
                                                                                                            _ = v34
                                                                                                            _ = v42
                                                                                                        }
                                                                                                    }
                                                                                                    if len(v1) > 0 {
                                                                                                        v1[0] = (((("\n\t" + "a") + "\n\t") + "a") + "a")
                                                                                                    }
                                                                                                }
                                                                                            case 42:
                                                                                                if true {
                                                                                                    v43 := GenS3{F0: 1000, F1: (v17 * 0), F2: 2}
                                                                                                    var v44 bool = (float64(v5) <= (float64((0.1 / (v5 * (-2.25 - 3.14159)))) / v26))
                                                                                                    _ = v43
                                                                                                    _ = v44
                                                                                                }
                                                                                            }
                                                                                            fmt.Println(v31, fi9, v26)
                                                                                            _ = v31
                                                                                        } else {
                                                                                            fmt.Println(v3)
                                                                                        }
                                                                                    }
                                                                                    var v45 uint64 = 510
                                                                                    _ = v30
                                                                                    _ = v45
                                                                                } else {
                                                                                    v46 := GenS9{F0: "a"}
                                                                                    var v47 uint32 = 217
                                                                                    v48 := GenS7{F0: v29, F1: fi8, F2: 2}
                                                                                    _ = v46
                                                                                    _ = v47
                                                                                    _ = v48
                                                                                }
                                                                            case "hello":
                                                                                v2 = append(v2, "hello")
                                                                            }
                                                                            _ = v28
                                                                            _ = v29
                                                                        }
                                                                        v49 := GenS3{F0: -770, F1: v17, F2: 700}
                                                                        if (true || (0 > 0)) {
                                                                            fmt.Println(v26, v10)
                                                                            v50 := GenS8{F0: v4}
                                                                            v7 = (v7 / (((v7 | 4294967295)) | 1))
                                                                            _ = v50
                                                                        }
                                                                        _ = v49
                                                                    default:
                                                                        switch fi9 {
                                                                        case 128:
                                                                            if (((-2.25 == 0.1) && false) || (3.14159 <= v5)) {
                                                                                if len(v1) > 0 {
                                                                                    v1[0] = "a"
                                                                                }
                                                                                if (true && true) {
                                                                                    v51 := GenS7{F0: int8(v10), F1: (fi9 / (((fi8 ^ (fi9 - 82))) | 1)), F2: 1}
                                                                                    var v52 int8 = int8(v7)
                                                                                    _ = v51
                                                                                    _ = v52
                                                                                }
                                                                                v53 := GenS10{F0: 348, F1: uint16(v7), F2: (v10 - v3)}
                                                                                _ = v53
                                                                            }
                                                                        case 42:
                                                                            _ = len(v2)
                                                                        case 50:
                                                                            if len(v2) > 0 {
                                                                                v2[0] = ""
                                                                            }
                                                                            var v54 int16 = -32768
                                                                            _ = v54
                                                                        default:
                                                                            v4 -= v4
                                                                        }
                                                                    }
                                                                } else {
                                                                    v4 = v4
                                                                    v55 := GenS9{F0: "hello"}
                                                                    _ = v55
                                                                }
                                                            }
                                                        }
                                                        switch "a" {
                                                        case "a":
                                                            switch ("" + "") {
                                                            case "a":
                                                                v17 = -541
                                                            case "":
                                                                switch "a" {
                                                                case "a":
                                                                    v56 := GenS2{F0: 182, F1: 42, F2: 614}
                                                                    _ = v56
                                                                }
                                                            default:
                                                                if false {
                                                                    switch -284 {
                                                                    case -861:
                                                                        _ = len(v1)
                                                                    case 32767:
                                                                        var v57 int = (^(fi9))
                                                                        _ = v57
                                                                    default:
                                                                        _ = len(v1)
                                                                    }
                                                                } else {
                                                                    v58 := GenS11{F0: -653, F1: 946}
                                                                    v59 := GenS4{F0: int(v4), F1: "hello", F2: (^(v10))}
                                                                    _ = v58
                                                                    _ = v59
                                                                }
                                                            }
                                                        case "foo bar":
                                                            var v60 uint8 = 127
                                                            _ = v60
                                                        case "\n\t":
                                                            fmt.Println(v3, v10)
                                                        }
                                                    }
                                                    switch (fi9 ^ (^(fi9))) {
                                                    case -414:
                                                        v7 += 859
                                                    }
                                                    _ = v24
                                                    _ = v25
                                                    _ = v26
                                                }
                                            case 0:
                                                v2 = append(v2, "foo bar")
                                            case 32767:
                                                var v61 int64 = (-(v17))
                                                _ = v61
                                            }
                                            if true {
                                                if true {
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                    if false {
                                                        v62 := GenS2{F0: 404, F1: int8(v17), F2: 323}
                                                        var v63 int64 = (v17 * -946)
                                                        _ = v62
                                                        _ = v63
                                                    }
                                                }
                                                var v64 string = ("\n\t" + ("hello" + (((("hello" + "foo bar") + ("\n\t" + "\n\t")) + "") + "foo bar")))
                                                _ = v64
                                            }
                                        }
                                    case 2147483647:
                                        if (!true) {
                                            if len(v1) > 0 {
                                                v1[0] = "hello"
                                            }
                                            v10 += (v3 >> ((v7) & 63))
                                            var v65 uint32 = 65535
                                            if true {
                                                var v66 int8 = 78
                                                _ = v66
                                            } else {
                                                v1 = append(v1, (("a" + (("" + ("hello" + "\n\t")) + "")) + ("\n\t" + "a")))
                                                if true {
                                                    var v67 uint8 = 127
                                                    if false {
                                                        if len(v2) >= 2 {
                                                            v2 = v2[1:len(v2)]
                                                        }
                                                    }
                                                    v67 = v67
                                                    _ = v67
                                                } else {
                                                    var v68 string = "foo bar"
                                                    var v69 float32 = v5
                                                    var v70 int64 = v17
                                                    _ = v68
                                                    _ = v69
                                                    _ = v70
                                                }
                                                var v71 uint = (v7 / ((42) | 1))
                                                _ = v71
                                            }
                                            var v72 int8 = 86
                                            if ((1 != 42) || true) {
                                                var v73 float32 = (v5 + (3.14159 / 1e10))
                                                v73 = v73
                                                fmt.Println(fi9, fi8, v5)
                                                _ = v73
                                            } else {
                                                v1 = append(v1, "\n\t")
                                                if (v7 < v7) {
                                                    if true {
                                                        v7 = v7
                                                        v74 := GenS8{F0: v4}
                                                        _ = v74
                                                    } else {
                                                        v75 := GenS7{F0: (v72 % (((v72 & 44)) | 1)), F1: int(v7), F2: 162}
                                                        var v76 string = "foo bar"
                                                        switch ("" + "\n\t") {
                                                        case "\n\t":
                                                            v4 -= 2
                                                        case "foo bar":
                                                            var v77 int = (fi8 ^ fi9)
                                                            _ = v77
                                                        default:
                                                            var v78 uint64 = 2
                                                            _ = v78
                                                        }
                                                        _ = v75
                                                        _ = v76
                                                    }
                                                } else {
                                                    if ("" > "hello") {
                                                        switch "hello" {
                                                        case "\n\t":
                                                            v79 := GenS10{F0: uint16(v7), F1: 0, F2: (^(v10))}
                                                            _ = v79
                                                        default:
                                                            v3 = (v10 >> (((v7 ^ v7)) & 63))
                                                        }
                                                    }
                                                    var v80 int16 = 42
                                                    switch -810 {
                                                    case 1:
                                                        if len(v1) >= 2 {
                                                            v1 = v1[1:len(v1)]
                                                        }
                                                    case 32768:
                                                        if len(v1) > 0 {
                                                            v1[0] = "hello"
                                                        }
                                                        v10 -= (^(v3))
                                                    }
                                                    _ = v80
                                                }
                                            }
                                            _ = v65
                                            _ = v72
                                        } else {
                                            v4 = (v4 | v4)
                                        }
                                        if false {
                                            v81 := GenS12{F0: (("hello" + "\n\t") + "hello"), F1: (v3 >> ((2) & 63)), F2: 589}
                                            _ = v81
                                        }
                                        _ = len(v2)
                                    default:
                                        v82 := GenS8{F0: 559}
                                        _ = v82
                                    }
                                } else {
                                    if ((32767 == -460) && ("\n\t" == "")) {
                                        var v83 uint16 = (v4 & v4)
                                        _ = v83
                                    } else {
                                        v84 := GenS5{F0: 542}
                                        _ = v84
                                    }
                                    v85 := GenS13{F0: (^(v4))}
                                    switch "" {
                                    case "hello":
                                        fmt.Println(v17)
                                    case "a":
                                        if false {
                                            var v86 int32 = (v10 % ((v3) | 1))
                                            _ = v86
                                        }
                                    default:
                                        switch "\n\t" {
                                        case "foo bar":
                                            v17 = v17
                                        case "a":
                                            fmt.Println(v4, v10, v3)
                                        default:
                                            var v87 float64 = 0.1
                                            _ = v87
                                        }
                                    }
                                    _ = v85
                                }
                                fmt.Println(fi9, fi9, v4)
                                _ = v17
                            }
                        }
                        _ = fi9
                        _ = v10
                    }
                    fmt.Println(fi8)
                } else {
                    var v88 int64 = -1
                    v3 = v3
                    var v89 uint8 = 141
                    _ = v88
                    _ = v89
                }
                var v90 int64 = 979
                for fi91 := 0; fi91 < 1; fi91++ {
                    if ("" > (("hello" + ("" + "")) + "\n\t")) {
                        var v92 uint64 = 144
                        _ = v92
                    }
                    _ = fi91
                }
                _ = v90
            }
            switch ("\n\t" + (("foo bar" + "hello") + "a")) {
            case "a":
                if ((v4 | 255) <= v4) {
                    var v93 uint64 = 42
                    _ = v93
                } else {
                    for fi94 := 0; fi94 < 1; fi94++ {
                        if true {
                            var v95 int = fi94
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            v4 -= (v4 / ((885) | 1))
                            _ = v95
                        } else {
                            v96 := GenS3{F0: 965, F1: 951, F2: 25}
                            v4 += v4
                            if len(v1) > 0 {
                                v1[0] = ((("foo bar" + "a") + (("foo bar" + ("a" + (("hello" + "\n\t") + ""))) + "a")) + ((("\n\t" + "foo bar") + "foo bar") + ("\n\t" + "foo bar")))
                            }
                            v97 := GenS8{F0: v4}
                            _ = v96
                            _ = v97
                        }
                        if false {
                            v2 = append(v2, "foo bar")
                        }
                        _ = fi94
                    }
                }
            case "foo bar":
                for fi98 := 0; fi98 < 16; fi98++ {
                    v1 = append(v1, ("" + ("\n\t" + "")))
                    if len(v1) > 0 {
                        v1[0] = ((("" + "foo bar") + "\n\t") + ("a" + ("hello" + "hello")))
                    }
                    _ = fi98
                }
            case "":
                for fi99 := 0; fi99 < 15; fi99++ {
                    v4 += (^(v4))
                    _ = fi99
                }
            default:
                v3 -= v3
            }
            _ = fi8
        }
        switch "hello" {
        case "foo bar":
            if true {
                for fi100 := 0; fi100 < 12; fi100++ {
                    fmt.Println(v3, fi100, v4)
                    v2 = append(v2, "\n\t")
                    _ = fi100
                }
                fmt.Println(v4, v7, v7)
                var v101 float64 = (0.1 / (0.1 * 1e10))
                _ = v101
            }
        }
        _ = v7
    } else {
        v3 = 42
        var v102 float32 = v5
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = v102
    }
    fmt.Println(v4, v3, v4)
    fmt.Println(v3)
    v4 = v4
    for fi103 := 0; fi103 < 1; fi103++ {
        fmt.Println(v5, v4, v5)
        if (false || true) {
            var v104 bool = true
            _ = v104
        }
        switch (fi103 * (fi103 + fi103)) {
        case 38:
            var v105 bool = ((true && (float32(0.0) <= v5)) || false)
            _ = v105
        case 32768:
            v4 = (v4 / ((2) | 1))
        default:
            if true {
                _ = len(v2)
            } else {
                var v106 uint32 = 978
                var v107 uint = 103
                _ = v106
                _ = v107
            }
        }
        _ = fi103
    }
    for fi108 := 0; fi108 < 16; fi108++ {
        if (int16(v3) >= 42) {
            var v109 uint64 = 4294967295
            for fi110 := 0; fi110 < 10; fi110++ {
                v4 -= 2
                _ = fi110
            }
            switch (fi108 - fi108) {
            case -177:
                var v111 uint8 = 0
                _ = v111
            case 918:
                v5 = v5
            default:
                fmt.Println(v4)
            }
            _ = v109
        } else {
            v112 := GenS3{F0: 42, F1: 1, F2: 65535}
            var v113 uint64 = 726
            switch 0 {
            case 2147483647:
                switch (fi108 * (fi108 | fi108)) {
                case -960:
                    switch "\n\t" {
                    case "\n\t":
                        switch ("\n\t" + ("foo bar" + "foo bar")) {
                        case "a":
                            for fi114 := 0; fi114 < 8; fi114++ {
                                var v115 uint32 = 2
                                if (!true) {
                                    v2 = append(v2, "a")
                                    var v116 int = (fi114 & fi108)
                                    _ = v116
                                }
                                var v117 int32 = (^(v3))
                                _ = fi114
                                _ = v115
                                _ = v117
                            }
                        case "hello":
                            v4 -= 65535
                        default:
                            v118 := GenS14{F0: -552, F1: 1000, F2: (!(-660 <= -1))}
                            _ = v118
                        }
                    default:
                        var v119 uint = 132
                        _ = v119
                    }
                case 663:
                    var v120 uint8 = 42
                    _ = v120
                case 32767:
                    switch "" {
                    case "foo bar":
                        v3 -= v3
                    }
                }
            case 127:
                var v121 string = ("" + "")
                _ = v121
            default:
                for fi122 := 0; fi122 < 11; fi122++ {
                    var v123 uint = 42
                    v124 := GenS15{F0: 42, F1: (-(fi122))}
                    var v125 float32 = float32(0.0)
                    _ = fi122
                    _ = v123
                    _ = v124
                    _ = v125
                }
            }
            _ = v112
            _ = v113
        }
        v3 -= (v3 / (((-(v3))) | 1))
        _ = fi108
    }
    v4 += v4
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
