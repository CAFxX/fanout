package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32 }
type GenS2 struct { F0 uint8; F1 uint64 }
type GenS3 struct { F0 uint16; F1 uint32 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint64; F1 int32; F2 uint16 }
type GenS6 struct { F0 int32; F1 int16; F2 string }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 bool; F1 string }
type GenS9 struct { F0 int8 }
type GenS10 struct { F0 int }
type GenS11 struct { F0 int8; F1 uint8; F2 bool }

func main() {
    v1 := make(map[uint32]int)
    var v2 bool = true
    v3 := make(map[string]int)
    v4 := []int{589}
    var v5 string = ("hello" + "a")
    v6 := GenS1{F0: v2, F1: 813}
    for fi7 := 0; fi7 < 2; fi7++ {
        if v2 {
            var v8 uint64 = 42
            fmt.Println(v8)
            v5 = v5
            _ = v8
        }
        v9 := GenS1{F0: true, F1: 2147483647}
        _ = fi7
        _ = v9
    }
    for fi10 := 0; fi10 < 11; fi10++ {
        var v11 uint = 4294967295
        switch fi10 {
        case 718:
            v12 := v3[("a" + "foo bar")]
            _ = v12
        default:
            var v13 bool = v2
            _ = v13
        }
        v11 += (v11 | (v11 ^ v11))
        _ = fi10
        _ = v11
    }
    v14 := GenS1{F0: (128 <= 115), F1: 1}
    if v2 {
        _ = len(v3)
    } else {
        if (v2 || ((true && v2) && v2)) {
            var v15 bool = v2
            _ = v15
        } else {
            var v16 int8 = 1
            var v17 int32 = 1
            _ = v16
            _ = v17
        }
    }
    if v2 {
        v2 = true
        fmt.Println(v2, v2)
    } else {
        for fi18 := 0; fi18 < 17; fi18++ {
            for fi19 := 0; fi19 < 19; fi19++ {
                v3["a"] = fi18
                switch fi18 {
                case -589:
                    var v20 int8 = -1
                    _ = v20
                case 2147483647:
                    v2 = (!false)
                case 42:
                    v5 = v5
                default:
                    v5 = "a"
                }
                switch (fi18 - (fi18 << ((42) & 63))) {
                case 0:
                    var v21 float32 = (float32(float64(0.1)) / (1.5 / 1e10))
                    _ = v21
                case -344:
                    var v22 float32 = (1.5 * 3.14159)
                    _ = v22
                }
                _ = fi19
            }
            _ = fi18
        }
        var v23 bool = (!(!(1e10 == float64(1e10))))
        var v24 uint64 = 0
        _ = v23
        _ = v24
    }
    delete(v1, 789)
    if (433 != 0) {
        switch -573 {
        case 9223372036854775807:
            if (((true || (-1 == 1)) || false) || (53 <= 1)) {
                if (v2 && true) {
                    if (!(367 != 2)) {
                        if v2 {
                            if (424 != 2) {
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                v2 = (v2 && (978 > -32768))
                                v25 := GenS1{F0: v2, F1: 2147483648}
                                _ = v25
                            }
                        }
                        for fi26 := 0; fi26 < 1; fi26++ {
                            var v27 uint32 = 0
                            if (v2 || v2) {
                                var v28 uint32 = v27
                                _ = v28
                            }
                            _ = fi26
                            _ = v27
                        }
                        fmt.Println(v2, v5)
                    } else {
                        for fi29 := 0; fi29 < 14; fi29++ {
                            if (927 != 42) {
                                v2 = v2
                                switch fi29 {
                                case 2147483647:
                                    var v30 uint = 244
                                    _ = v30
                                case 530:
                                    switch (fi29 * (fi29 / (((fi29 ^ fi29)) | 1))) {
                                    case 2147483647:
                                        delete(v1, 4294967295)
                                    case 208:
                                        var v31 float32 = (float32((1e10 + -2.25)) + 3.14159)
                                        _ = v31
                                    case -762:
                                        var v32 bool = (4294967295 < 370)
                                        _ = v32
                                    default:
                                        switch (^(fi29)) {
                                        case -2147483648:
                                            if false {
                                                v2 = (905 == 764)
                                                for fi33 := 0; fi33 < 8; fi33++ {
                                                    if false {
                                                        v34 := GenS2{F0: 127, F1: 648}
                                                        var v35 uint = 42
                                                        _ = v34
                                                        _ = v35
                                                    } else {
                                                        v36 := GenS1{F0: (664 == 668), F1: 2147483647}
                                                        v5 = (v5 + (v5 + "hello"))
                                                        _ = v36
                                                    }
                                                    _ = fi33
                                                }
                                            } else {
                                                delete(v1, 42)
                                            }
                                        case -152:
                                            v2 = (84 != 2)
                                        case 172:
                                            switch v5 {
                                            case "\n\t":
                                                for fi37 := 0; fi37 < 16; fi37++ {
                                                    var v38 uint32 = 463
                                                    switch (v5 + ("a" + "\n\t")) {
                                                    case "foo bar":
                                                        switch "\n\t" {
                                                        case "":
                                                            var v39 int8 = 0
                                                            _ = v39
                                                        case "\n\t":
                                                            if true {
                                                                v2 = false
                                                                v5 = v5
                                                                if false {
                                                                    v2 = (false && false)
                                                                    if true {
                                                                        if len(v4) >= 2 {
                                                                            v4 = v4[1:len(v4)]
                                                                        }
                                                                        v40 := v1[v38]
                                                                        _ = v40
                                                                    } else {
                                                                        var v41 int = (fi29 << ((18446744073709551615) & 63))
                                                                        v42 := GenS1{F0: false, F1: v38}
                                                                        _ = v41
                                                                        _ = v42
                                                                    }
                                                                } else {
                                                                    var v43 int = -1
                                                                    fmt.Println(fi37, v5, v2)
                                                                    delete(v1, 2)
                                                                    _ = v43
                                                                }
                                                            } else {
                                                                var v44 uint8 = 127
                                                                var v45 int64 = -1
                                                                _ = v44
                                                                _ = v45
                                                            }
                                                        default:
                                                            v2 = (!(39 != 127))
                                                        }
                                                    case "a":
                                                        v38 -= 61
                                                    }
                                                    v46 := GenS3{F0: 65535, F1: 4294967295}
                                                    _ = fi37
                                                    _ = v38
                                                    _ = v46
                                                }
                                            case "foo bar":
                                                for fi47 := 0; fi47 < 20; fi47++ {
                                                    var v48 uint16 = 32767
                                                    _ = fi47
                                                    _ = v48
                                                }
                                            }
                                        }
                                    }
                                default:
                                    if v2 {
                                        var v49 float64 = (-2.25 - 0.1)
                                        fmt.Println(v5, v5, v5)
                                        if (v2 && (-2.25 >= (float32(v49) * 1e10))) {
                                            v50 := GenS2{F0: 1, F1: 18446744073709551615}
                                            _ = v50
                                        } else {
                                            for fi51 := 0; fi51 < 1; fi51++ {
                                                fmt.Println(v49, v5, fi51)
                                                v5 = "foo bar"
                                                _ = fi51
                                            }
                                        }
                                        _ = v49
                                    } else {
                                        switch (fi29 & (^(fi29))) {
                                        case -914:
                                            if (218 >= 0) {
                                                switch v5 {
                                                case "foo bar":
                                                    delete(v1, 0)
                                                default:
                                                    var v52 uint64 = 2
                                                    _ = v52
                                                }
                                            } else {
                                                v5 = v5
                                                v5 = "a"
                                            }
                                        case 127:
                                            v53 := GenS4{F0: 314}
                                            _ = v53
                                        case 0:
                                            switch "foo bar" {
                                            case "":
                                                for fi54 := 0; fi54 < 7; fi54++ {
                                                    if (-1 <= 1) {
                                                        switch v5 {
                                                        case "":
                                                            v5 = "foo bar"
                                                        case "hello":
                                                            var v55 int = fi54
                                                            _ = v55
                                                        }
                                                        v2 = false
                                                        switch v5 {
                                                        case "a":
                                                            v5 = ("foo bar" + v5)
                                                        default:
                                                            v56 := GenS4{F0: 2}
                                                            _ = v56
                                                        }
                                                    }
                                                    v57 := GenS3{F0: 2, F1: 2147483647}
                                                    _ = fi54
                                                    _ = v57
                                                }
                                            default:
                                                switch "\n\t" {
                                                case "hello":
                                                    if (0 > 22) {
                                                        if len(v4) > 0 {
                                                            v4[0] = fi29
                                                        }
                                                        switch "\n\t" {
                                                        case "hello":
                                                            if (101 == -2147483648) {
                                                                v2 = (65535 >= 65535)
                                                                fmt.Println(v2)
                                                            } else {
                                                                v2 = (!true)
                                                                if (!(!(!(!((("foo bar" + v5) < v5) && true))))) {
                                                                    if (v2 || (127 == 42)) {
                                                                        v2 = v2
                                                                    } else {
                                                                        v2 = v2
                                                                        var v58 int = -716
                                                                        if true {
                                                                            var v59 uint16 = 2
                                                                            _ = v59
                                                                        } else {
                                                                            if v2 {
                                                                                delete(v3, ("hello" + v5))
                                                                                if ("foo bar" > v5) {
                                                                                    var v60 int64 = 1
                                                                                    if (-2.25 == (0.0 / (-2.25 / (3.14159 / 1.0)))) {
                                                                                        fmt.Println(v5)
                                                                                        switch v58 {
                                                                                        case 9223372036854775807:
                                                                                            if (v60 < v60) {
                                                                                                var v61 int64 = v60
                                                                                                var v62 uint32 = 1
                                                                                                v61 = (v60 * (v60 | (v60 ^ v61)))
                                                                                                _ = v61
                                                                                                _ = v62
                                                                                            }
                                                                                        default:
                                                                                            v5 = v5
                                                                                        }
                                                                                        switch (v58 >> ((64) & 63)) {
                                                                                        case 0:
                                                                                            fmt.Println(fi29, v58, fi29)
                                                                                        case 488:
                                                                                            v63 := GenS4{F0: 4294967295}
                                                                                            _ = v63
                                                                                        default:
                                                                                            if ("" > ("hello" + ((("a" + (v5 + v5)) + "hello") + "hello"))) {
                                                                                                var v64 uint64 = 999
                                                                                                _ = v64
                                                                                            } else {
                                                                                                v60 = 1
                                                                                                v65 := GenS2{F0: 42, F1: 18446744073709551615}
                                                                                                _ = v65
                                                                                            }
                                                                                        }
                                                                                    } else {
                                                                                        v5 = "\n\t"
                                                                                    }
                                                                                    v66 := GenS5{F0: 1, F1: 1, F2: 2}
                                                                                    _ = v60
                                                                                    _ = v66
                                                                                }
                                                                            }
                                                                            delete(v3, "a")
                                                                            v67 := GenS4{F0: 2}
                                                                            _ = v67
                                                                        }
                                                                        _ = v58
                                                                    }
                                                                    for fi68 := 0; fi68 < 18; fi68++ {
                                                                        v5 = (v5 + (("\n\t" + v5) + ("a" + (v5 + v5))))
                                                                        v2 = true
                                                                        if (!true) {
                                                                            if (true || false) {
                                                                                v69 := GenS6{F0: int32(fi68), F1: -786, F2: "a"}
                                                                                v2 = v2
                                                                                var v70 uint16 = 0
                                                                                _ = v69
                                                                                _ = v70
                                                                            }
                                                                            v2 = v2
                                                                        }
                                                                        _ = fi68
                                                                    }
                                                                    _ = len(v3)
                                                                }
                                                                for fi71 := 0; fi71 < 7; fi71++ {
                                                                    v5 = v5
                                                                    if ((((v5 + v5) + "\n\t") + v5) < (v5 + v5)) {
                                                                        var v72 uint8 = 42
                                                                        if v2 {
                                                                            v73 := GenS7{F0: 127}
                                                                            fmt.Println(v5, v72, fi29)
                                                                            _ = v73
                                                                        }
                                                                        v74 := GenS8{F0: true, F1: ("\n\t" + "")}
                                                                        v72 = (^(v72))
                                                                        _ = v72
                                                                        _ = v74
                                                                    }
                                                                    delete(v3, v5)
                                                                    _ = fi71
                                                                }
                                                            }
                                                        }
                                                    }
                                                case "":
                                                    var v75 bool = (!(!true))
                                                    _ = v75
                                                default:
                                                    var v76 int64 = 565
                                                    _ = v76
                                                }
                                            }
                                        }
                                        for fi77 := 0; fi77 < 4; fi77++ {
                                            v2 = v2
                                            if (v2 || (((v5 + v5) + ("a" + v5)) > ((v5 + v5) + ""))) {
                                                fmt.Println(fi29)
                                                if len(v4) >= 2 {
                                                    v4 = v4[1:len(v4)]
                                                }
                                                switch v5 {
                                                case "foo bar":
                                                    var v78 bool = v2
                                                    _ = v78
                                                case "a":
                                                    if len(v4) > 0 {
                                                        v4[0] = (fi77 | (fi77 & (-(fi77))))
                                                    }
                                                default:
                                                    if (818 == -1) {
                                                        delete(v1, 1)
                                                    } else {
                                                        if ((1.5 < ((0.0 * 3.14159) * (3.14159 * ((1.5 + float32(-2.25)) / 1e10)))) || true) {
                                                            v79 := GenS7{F0: 0}
                                                            _ = v79
                                                        }
                                                        var v80 int32 = 2147483647
                                                        if ((1000 < -32768) && ((v5 + "foo bar") == ("hello" + (v5 + v5)))) {
                                                            v2 = v2
                                                        }
                                                        _ = v80
                                                    }
                                                }
                                            }
                                            if true {
                                                switch v5 {
                                                case "":
                                                    v81 := GenS2{F0: 1, F1: 2}
                                                    v82 := v1[992]
                                                    v5 = (("a" + v5) + v5)
                                                    if (974 <= int16(fi77)) {
                                                        v83 := GenS3{F0: 42, F1: 284}
                                                        var v84 string = (("" + v5) + "")
                                                        switch -180 {
                                                        case 127:
                                                            var v85 string = (v5 + v5)
                                                            _ = v85
                                                        }
                                                        _ = v83
                                                        _ = v84
                                                    }
                                                    _ = v81
                                                    _ = v82
                                                case "foo bar":
                                                    fmt.Println(v5)
                                                case "\n\t":
                                                    switch (fi29 << ((2) & 63)) {
                                                    case 301:
                                                        if (127 > -1) {
                                                            switch ("foo bar" + "foo bar") {
                                                            case "foo bar":
                                                                v86 := GenS9{F0: -128}
                                                                _ = v86
                                                            case "hello":
                                                                v5 = v5
                                                            default:
                                                                _ = len(v4)
                                                            }
                                                        } else {
                                                            fmt.Println(fi77, fi77)
                                                            var v87 int = 960
                                                            v88 := GenS6{F0: 0, F1: -32768, F2: "\n\t"}
                                                            _ = v87
                                                            _ = v88
                                                        }
                                                    case 128:
                                                        fmt.Println(v2)
                                                    case 201:
                                                        switch (-(fi29)) {
                                                        case -9223372036854775808:
                                                            var v89 int32 = -2147483648
                                                            _ = v89
                                                        }
                                                    }
                                                default:
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                }
                                                var v90 uint8 = 228
                                                _ = v90
                                            }
                                            _ = fi77
                                        }
                                    }
                                }
                            }
                            switch 4294967295 {
                            case 2147483647:
                                v91 := GenS1{F0: v2, F1: 0}
                                _ = v91
                            default:
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            }
                            _ = fi29
                        }
                    }
                    switch 127 {
                    case 4294967295:
                        switch (((v5 + "") + v5) + (v5 + ("a" + (v5 + ("hello" + v5))))) {
                        case "foo bar":
                            v5 = ("" + "")
                        }
                    default:
                        v92 := GenS8{F0: v2, F1: (v5 + "foo bar")}
                        _ = v92
                    }
                    var v93 uint64 = 2
                    _ = v93
                } else {
                    for fi94 := 0; fi94 < 11; fi94++ {
                        for fi95 := 0; fi95 < 9; fi95++ {
                            if v2 {
                                v96 := GenS3{F0: 773, F1: 2147483648}
                                v5 = (v5 + "hello")
                                _ = len(v4)
                                var v97 bool = v2
                                _ = v96
                                _ = v97
                            } else {
                                if (65535 > 114) {
                                    var v98 bool = false
                                    _ = v98
                                } else {
                                    if true {
                                        v5 = (v5 + ("" + "foo bar"))
                                    }
                                    fmt.Println(fi95, fi94, v5)
                                    var v99 uint8 = 42
                                    v2 = v2
                                    var v100 int8 = -128
                                    _ = v100
                                    _ = v99
                                }
                                switch ((v5 + v5) + v5) {
                                case "":
                                    v2 = v2
                                case "a":
                                    v101 := GenS2{F0: 255, F1: 457}
                                    _ = v101
                                case "\n\t":
                                    fmt.Println(fi94, fi95)
                                default:
                                    var v102 uint64 = 412
                                    _ = v102
                                }
                                var v103 int64 = 42
                                _ = v103
                            }
                            if true {
                                if v2 {
                                    fmt.Println(fi94)
                                    _ = len(v4)
                                    if (0 < 0) {
                                        v5 = v5
                                        v104 := GenS7{F0: 1}
                                        var v105 int8 = 1
                                        _ = v104
                                        _ = v105
                                    } else {
                                        switch (v5 + "foo bar") {
                                        case "foo bar":
                                            switch fi95 {
                                            case -493:
                                                switch (fi94 | fi95) {
                                                case -2147483648:
                                                    var v106 uint8 = 228
                                                    _ = v106
                                                default:
                                                    v5 = v5
                                                }
                                            case -1:
                                                v107 := GenS1{F0: (!((!(-2147483648 == 794)) && (((1 != -264) && v2) && v2))), F1: 2147483647}
                                                _ = v107
                                            case 127:
                                                if false {
                                                    switch v5 {
                                                    case "":
                                                        v108 := GenS4{F0: 65535}
                                                        _ = v108
                                                    default:
                                                        var v109 int32 = -921
                                                        _ = v109
                                                    }
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                    var v110 int8 = -122
                                                    v2 = v2
                                                    var v111 uint = 42
                                                    _ = v110
                                                    _ = v111
                                                }
                                            }
                                        default:
                                            fmt.Println(v2, fi95)
                                        }
                                        v5 = (v5 + ((v5 + "") + v5))
                                        v2 = v2
                                    }
                                }
                            }
                            _ = fi95
                        }
                        _ = fi94
                    }
                }
            }
        case -566:
            switch 486 {
            case 32767:
                var v112 int32 = 2147483647
                _ = v112
            case 32768:
                v113 := GenS10{F0: 4294967295}
                _ = v113
            case 128:
                _ = len(v1)
            }
        case -85:
            _ = len(v1)
        default:
            v114 := GenS11{F0: 1, F1: 36, F2: v2}
            _ = v114
        }
    }
    var v115 bool = (254 == 1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v115:", v115)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v115
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
