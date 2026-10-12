package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16; F2 int16 }
type GenS2 struct { F0 uint8; F1 uint16; F2 bool }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 bool; F1 uint32 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint8; F1 uint16; F2 int }
type GenS7 struct { F0 uint8; F1 string }
type GenS8 struct { F0 uint16; F1 uint64 }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 uint16; F1 bool; F2 string }
type GenS11 struct { F0 string; F1 bool }
func genFn0() uint {
    var v5 uint32 = 65535
    v5 = v5
    v6 := GenS1{F0: 32768, F1: 0, F2: 1000}
    var v7 float32 = -2.25
    _ = v5
    _ = v6
    _ = v7
    return 255
}

func main() {
    var v1 int32 = 0
    var v2 int64 = -797
    var v3 uint16 = 0
    v4 := []string{"\n\t", "foo bar", (("a" + ("\n\t" + "a")) + "a"), ""}
    v8 := GenS1{F0: 9223372036854775807, F1: (v3 << ((631) & 63)), F2: 797}
    if false {
        if (!(!((v2 << ((2) & 63)) != (v2 >> ((65535) & 63))))) {
            switch ("\n\t" + (("hello" + "a") + "\n\t")) {
            case "\n\t":
                if true {
                    for fi9 := 0; fi9 < 12; fi9++ {
                        if ((float64(((0.1 + 1.5) - 1e10)) / ((-2.25 * 3.14159) / 0.1)) <= (float64(float32(3.14159)) / (((-2.25 + float64(1.5)) / float64((1.5 + (1.5 + 0.0)))) * 1e10))) {
                            for fi10 := 0; fi10 < 14; fi10++ {
                                var v11 uint = 502
                                var v12 bool = false
                                fmt.Println(v2)
                                _ = fi10
                                _ = v11
                                _ = v12
                            }
                        } else {
                            _ = len(v4)
                            var v13 uint32 = 590
                            _ = v13
                        }
                        switch fi9 {
                        case 127:
                            if (false || (v1 != 2147483647)) {
                                switch ("" + "a") {
                                case "":
                                    var v14 int16 = 1
                                    _ = v14
                                case "hello":
                                    fmt.Println(v3, v2)
                                case "\n\t":
                                    v1 = v1
                                default:
                                    switch "hello" {
                                    case "\n\t":
                                        var v15 int16 = -970
                                        _ = v15
                                    case "":
                                        fmt.Println(v2, v1, v1)
                                    }
                                }
                                v16 := GenS2{F0: 123, F1: 2, F2: true}
                                _ = v16
                            } else {
                                for fi17 := 0; fi17 < 18; fi17++ {
                                    fmt.Println(v1, v3)
                                    _ = fi17
                                }
                                if (-1 != 0) {
                                    switch 127 {
                                    case -864:
                                        fmt.Println(v1, v1, v2)
                                    case -567:
                                        var v18 int16 = 776
                                        _ = v18
                                    default:
                                        var v19 int16 = 710
                                        _ = v19
                                    }
                                    fmt.Println(v3)
                                    v20 := GenS1{F0: (fi9 << ((4294967295) & 63)), F1: (v3 % ((492) | 1)), F2: 1}
                                    _ = v20
                                } else {
                                    v21 := GenS1{F0: fi9, F1: v3, F2: 816}
                                    if (902 == (v2 ^ v2)) {
                                        switch "\n\t" {
                                        case "foo bar":
                                            for fi22 := 0; fi22 < 1; fi22++ {
                                                if len(v4) >= 2 {
                                                    v4 = v4[1:len(v4)]
                                                }
                                                v1 = 42
                                                if false {
                                                    if ((1e10 / 1.0) <= float32(1.5)) {
                                                        v3 = (v3 ^ 42)
                                                        v2 = 42
                                                        v23 := GenS2{F0: 2, F1: (v3 | (^(v3))), F2: true}
                                                        _ = v23
                                                    }
                                                    if true {
                                                        var v24 string = "\n\t"
                                                        v24 = v24
                                                        _ = v24
                                                    } else {
                                                        if ("\n\t" != "\n\t") {
                                                            var v25 float64 = (1.5 / 0.1)
                                                            _ = v25
                                                        }
                                                    }
                                                }
                                                v26 := GenS2{F0: 245, F1: (^(v3)), F2: (v3 < (v3 << ((108) & 63)))}
                                                v1 += v1
                                                fmt.Println(v3, v2, v3)
                                                _ = fi22
                                                _ = v26
                                            }
                                        case "\n\t":
                                            switch (fi9 * (fi9 | fi9)) {
                                            case 32767:
                                                if len(v4) > 0 {
                                                    v4[0] = "\n\t"
                                                }
                                            default:
                                                if (false || (!false)) {
                                                    if true {
                                                        v1 += (v1 << ((0) & 63))
                                                    } else {
                                                        v27 := GenS3{F0: 1}
                                                        v2 -= (v2 << ((538) & 63))
                                                        v28 := GenS4{F0: (int(v1) != (fi9 - 358)), F1: 0}
                                                        _ = v27
                                                        _ = v28
                                                    }
                                                    v29 := GenS1{F0: fi9, F1: (v3 + (v3 + v3)), F2: -266}
                                                    _ = v29
                                                } else {
                                                    v3 = v3
                                                    v2 += 42
                                                    if ((true && (127 != -58)) || true) {
                                                        if ((!false) && true) {
                                                            for fi30 := 0; fi30 < 5; fi30++ {
                                                                switch fi30 {
                                                                case -287:
                                                                    if true {
                                                                        switch ("" + "hello") {
                                                                        case "hello":
                                                                            var v31 uint64 = 4294967295
                                                                            _ = v31
                                                                        case "\n\t":
                                                                            var v32 bool = false
                                                                            _ = v32
                                                                        }
                                                                        if len(v4) >= 2 {
                                                                            v4 = v4[1:len(v4)]
                                                                        }
                                                                        if (v3 == v3) {
                                                                            var v33 int64 = 42
                                                                            fmt.Println(v2, fi9)
                                                                            _ = v33
                                                                        } else {
                                                                            var v34 bool = true
                                                                            _ = v34
                                                                        }
                                                                    }
                                                                case 0:
                                                                    v35 := GenS4{F0: false, F1: uint32(v3)}
                                                                    _ = v35
                                                                case 2147483647:
                                                                    switch "a" {
                                                                    case "":
                                                                        fmt.Println(v3, v2, fi30)
                                                                        v1 -= 2147483647
                                                                        v1 = (v1 - (v1 / ((2147483647) | 1)))
                                                                        var v36 int64 = v2
                                                                        if false {
                                                                            var v37 float64 = float64((3.14159 - -2.25))
                                                                            if (3.14159 > -2.25) {
                                                                                if (!true) {
                                                                                    if true {
                                                                                        if ((true || (true || (56 != 299))) || (1 > 4294967295)) {
                                                                                            v38 := GenS5{F0: 1}
                                                                                            fmt.Println(v1, v36)
                                                                                            v37 = ((0.1 + 1.5) - 1e10)
                                                                                            fmt.Println(fi30, fi9, v36)
                                                                                            var v39 int = -2147483648
                                                                                            _ = v38
                                                                                            _ = v39
                                                                                        }
                                                                                    }
                                                                                    if ((!false) || (42 >= 2)) {
                                                                                        _ = len(v4)
                                                                                        v4 = append(v4, "\n\t")
                                                                                    } else {
                                                                                        var v40 uint16 = (^(v3))
                                                                                        _ = v40
                                                                                    }
                                                                                    var v41 string = "a"
                                                                                    _ = v41
                                                                                }
                                                                            } else {
                                                                                if (!(-1 < 42)) {
                                                                                    fmt.Println(fi9, v37, v37)
                                                                                }
                                                                            }
                                                                            var v42 bool = false
                                                                            _ = v37
                                                                            _ = v42
                                                                        } else {
                                                                            if false {
                                                                                v4 = append(v4, "hello")
                                                                                var v43 float32 = 3.14159
                                                                                v3 += (^(v3))
                                                                                var v44 int32 = (v1 >> ((18446744073709551615) & 63))
                                                                                v4 = append(v4, "")
                                                                                if (true || true) {
                                                                                    fmt.Println(v36)
                                                                                    if false {
                                                                                        fmt.Println(v43, fi9)
                                                                                        _ = len(v4)
                                                                                    } else {
                                                                                        if (((!((v1 ^ (v44 % (((v1 % (((v1 | 1)) | 1))) | 1))) > v44)) && false) || (598 == 255)) {
                                                                                            fmt.Println(v43)
                                                                                        }
                                                                                        var v45 int8 = 1
                                                                                        _ = v45
                                                                                    }
                                                                                } else {
                                                                                    fmt.Println(v36, fi30, v3)
                                                                                    if len(v4) > 0 {
                                                                                        v4[0] = "foo bar"
                                                                                    }
                                                                                    if (2 == 0) {
                                                                                        fmt.Println(v1)
                                                                                        if false {
                                                                                            v46 := GenS2{F0: 2, F1: (v3 >> ((0) & 63)), F2: false}
                                                                                            v36 -= -1
                                                                                            var v47 int16 = -1
                                                                                            _ = v46
                                                                                            _ = v47
                                                                                        } else {
                                                                                            v3 = v3
                                                                                        }
                                                                                        v36 += (v2 - 999)
                                                                                        fmt.Println(v36, v3, v3)
                                                                                    }
                                                                                    fmt.Println(v43)
                                                                                }
                                                                                _ = v43
                                                                                _ = v44
                                                                            } else {
                                                                                switch (("a" + "") + ("\n\t" + "hello")) {
                                                                                case "\n\t":
                                                                                    fmt.Println(fi9)
                                                                                case "":
                                                                                    var v48 uint = 610
                                                                                    _ = v48
                                                                                case "a":
                                                                                    v49 := GenS6{F0: 128, F1: v3, F2: (fi30 << ((uint(v3)) & 63))}
                                                                                    _ = v49
                                                                                }
                                                                            }
                                                                            var v50 uint8 = 127
                                                                            _ = v50
                                                                        }
                                                                        v2 -= (v36 ^ v36)
                                                                        _ = v36
                                                                    case "hello":
                                                                        if ((0 == 42) && true) {
                                                                            if (!true) {
                                                                                var v51 uint64 = 0
                                                                                v51 = 1
                                                                                _ = v51
                                                                            }
                                                                            if true {
                                                                                if true {
                                                                                    switch -2147483648 {
                                                                                    case 128:
                                                                                        _ = len(v4)
                                                                                    case 1:
                                                                                        v2 += 42
                                                                                    }
                                                                                    var v52 int64 = (^(v2))
                                                                                    _ = v52
                                                                                } else {
                                                                                    if ((707 != 816) && ((("hello" + (("a" + "a") + "")) + "\n\t") > "\n\t")) {
                                                                                        switch "foo bar" {
                                                                                        case "foo bar":
                                                                                            v53 := GenS3{F0: 253}
                                                                                            _ = v53
                                                                                        default:
                                                                                            var v54 uint = 42
                                                                                            _ = v54
                                                                                        }
                                                                                        var v55 int8 = 0
                                                                                        var v56 float32 = 0.0
                                                                                        v55 += (v55 * v55)
                                                                                        _ = v55
                                                                                        _ = v56
                                                                                    } else {
                                                                                        switch ("\n\t" + "foo bar") {
                                                                                        case "foo bar":
                                                                                            switch fi30 {
                                                                                            case 0:
                                                                                                fmt.Println(v2, fi30, fi30)
                                                                                            case 32768:
                                                                                                var v57 int8 = 95
                                                                                                _ = v57
                                                                                            case 2147483647:
                                                                                                var v58 int = int(v2)
                                                                                                _ = v58
                                                                                            }
                                                                                        case "hello":
                                                                                            v59 := GenS4{F0: false, F1: 65535}
                                                                                            _ = v59
                                                                                        case "":
                                                                                            v3 = v3
                                                                                            v2 -= (v2 - -528)
                                                                                        default:
                                                                                            var v60 uint64 = 4294967295
                                                                                            _ = v60
                                                                                        }
                                                                                    }
                                                                                    fmt.Println(fi9)
                                                                                    var v61 uint8 = 255
                                                                                    _ = v61
                                                                                }
                                                                            }
                                                                        }
                                                                    case "a":
                                                                        fmt.Println(v2, v1, fi30)
                                                                    }
                                                                default:
                                                                    v3 = v3
                                                                }
                                                                var v62 int8 = -128
                                                                fmt.Println(v3, fi30, v62)
                                                                _ = fi30
                                                                _ = v62
                                                            }
                                                            fmt.Println(v2, fi9, fi9)
                                                        } else {
                                                            for fi63 := 0; fi63 < 7; fi63++ {
                                                                v64 := GenS3{F0: 127}
                                                                var v65 int32 = v1
                                                                _ = fi63
                                                                _ = v64
                                                                _ = v65
                                                            }
                                                            if len(v4) > 0 {
                                                                v4[0] = ("\n\t" + "a")
                                                            }
                                                            v66 := GenS2{F0: 0, F1: (^(v3)), F2: (false && (2147483647 <= v1))}
                                                            _ = v66
                                                        }
                                                    } else {
                                                        if len(v4) >= 2 {
                                                            v4 = v4[1:len(v4)]
                                                        }
                                                        fmt.Println(fi9, fi9, v1)
                                                        for fi67 := 0; fi67 < 4; fi67++ {
                                                            var v68 int64 = v2
                                                            switch (("" + "foo bar") + "\n\t") {
                                                            case "":
                                                                if (168 != 51) {
                                                                    var v69 int64 = (v2 & v2)
                                                                    var v70 string = "hello"
                                                                    _ = v69
                                                                    _ = v70
                                                                }
                                                            default:
                                                                v68 -= (v2 - v68)
                                                            }
                                                            var v71 float32 = 1.5
                                                            _ = fi67
                                                            _ = v68
                                                            _ = v71
                                                        }
                                                    }
                                                }
                                            }
                                        }
                                        v3 = 588
                                        if false {
                                            v3 = (v3 ^ (^(v3)))
                                            var v72 uint64 = 921
                                            _ = v72
                                        } else {
                                            if false {
                                                var v73 uint8 = 127
                                                _ = v73
                                            } else {
                                                var v74 float64 = 0.1
                                                _ = v74
                                            }
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                        }
                                    } else {
                                        v1 -= (v1 << ((875) & 63))
                                    }
                                    for fi75 := 0; fi75 < 5; fi75++ {
                                        if false {
                                            if (false && true) {
                                                if true {
                                                    var v76 uint = 652
                                                    _ = v76
                                                } else {
                                                    v1 -= (v1 | (v1 / (((v1 | (v1 >> ((255) & 63)))) | 1)))
                                                    var v77 uint16 = (v3 | 32767)
                                                    v3 = 264
                                                    _ = v77
                                                }
                                                if len(v4) > 0 {
                                                    v4[0] = "a"
                                                }
                                                var v78 float32 = 0.1
                                                v79 := GenS5{F0: 1}
                                                _ = v78
                                                _ = v79
                                            }
                                            if true {
                                                v1 = -1
                                                v80 := GenS6{F0: 42, F1: (v3 | v3), F2: (-(fi9))}
                                                v4 = append(v4, "")
                                                _ = v80
                                            } else {
                                                v81 := GenS5{F0: 42}
                                                _ = len(v4)
                                                v82 := GenS7{F0: 0, F1: "foo bar"}
                                                _ = v81
                                                _ = v82
                                            }
                                            v83 := GenS8{F0: (v3 & (v3 << ((975) & 63))), F1: 406}
                                            _ = v83
                                        }
                                        if false {
                                            if true {
                                                fmt.Println(v2, fi75)
                                                v84 := GenS9{F0: false}
                                                _ = v84
                                            } else {
                                                v1 -= v1
                                            }
                                            var v85 int16 = 208
                                            v1 += (v1 * 42)
                                            _ = v85
                                        } else {
                                            switch ((((("foo bar" + "") + "a") + "a") + "foo bar") + (("\n\t" + "\n\t") + "foo bar")) {
                                            case "hello":
                                                var v86 int32 = v1
                                                _ = v86
                                            case "foo bar":
                                                switch -625 {
                                                case -533:
                                                    if ((v2 & v2) < 99) {
                                                        if len(v4) > 0 {
                                                            v4[0] = (((("" + "\n\t") + "a") + "\n\t") + (("foo bar" + ("" + ((("" + "hello") + "hello") + ("hello" + ("hello" + ("foo bar" + "\n\t")))))) + ""))
                                                        }
                                                        var v87 string = (("hello" + "hello") + ("foo bar" + "foo bar"))
                                                        v88 := GenS5{F0: 2}
                                                        _ = v87
                                                        _ = v88
                                                    }
                                                case 127:
                                                    v1 -= (v1 + 0)
                                                default:
                                                    v2 = (v2 + v2)
                                                }
                                            default:
                                                v1 = v1
                                            }
                                        }
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                        _ = fi75
                                    }
                                    _ = v21
                                }
                            }
                        case 32767:
                            switch 4294967295 {
                            case 1:
                                v2 = (v2 % (((v2 + (v2 >> ((4294967295) & 63)))) | 1))
                            case 42:
                                if len(v4) > 0 {
                                    v4[0] = "\n\t"
                                }
                            case 4294967295:
                                for fi89 := 0; fi89 < 13; fi89++ {
                                    if len(v4) > 0 {
                                        v4[0] = "foo bar"
                                    }
                                    v4 = append(v4, "\n\t")
                                    var v90 uint32 = 221
                                    _ = fi89
                                    _ = v90
                                }
                            }
                        }
                        _ = fi9
                    }
                    v4 = append(v4, "hello")
                    for fi91 := 0; fi91 < 1; fi91++ {
                        fmt.Println(fi91, v3)
                        _ = fi91
                    }
                } else {
                    v2 += v2
                    for fi92 := 0; fi92 < 15; fi92++ {
                        if (((true || false) && (!true)) && false) {
                            var v93 uint = 2
                            switch ((("a" + ("hello" + "foo bar")) + ("hello" + "")) + ("\n\t" + "")) {
                            case "foo bar":
                                switch (fi92 ^ 128) {
                                case 4294967295:
                                    for fi94 := 0; fi94 < 10; fi94++ {
                                        switch 4294967295 {
                                        case -1:
                                            var v95 float64 = 1e10
                                            _ = v95
                                        case -978:
                                            switch (fi92 ^ fi92) {
                                            case -804:
                                                if len(v4) >= 2 {
                                                    v4 = v4[1:len(v4)]
                                                }
                                            }
                                        case -2147483648:
                                            var v96 uint32 = 65535
                                            v93 -= v93
                                            if (true || (!true)) {
                                                switch (fi92 / (((fi92 / ((fi92) | 1))) | 1)) {
                                                case 9223372036854775807:
                                                    switch "" {
                                                    case "a":
                                                        v93 = uint(v96)
                                                    case "\n\t":
                                                        var v97 string = ""
                                                        _ = v97
                                                    case "foo bar":
                                                        switch fi94 {
                                                        case 4294967295:
                                                            var v98 uint = uint(v96)
                                                            _ = v98
                                                        default:
                                                            switch "foo bar" {
                                                            case "\n\t":
                                                                if ((v2 / ((int64(v93)) | 1)) > (v2 - 1)) {
                                                                    var v99 uint16 = 32767
                                                                    v100 := GenS10{F0: v99, F1: true, F2: "hello"}
                                                                    var v101 float32 = 1e10
                                                                    var v102 int16 = 633
                                                                    v93 -= v93
                                                                    _ = v100
                                                                    _ = v101
                                                                    _ = v102
                                                                    _ = v99
                                                                } else {
                                                                    if len(v4) >= 2 {
                                                                        v4 = v4[1:len(v4)]
                                                                    }
                                                                }
                                                            case "":
                                                                var v103 float32 = 3.14159
                                                                _ = v103
                                                            }
                                                        }
                                                    }
                                                case -2147483648:
                                                    v104 := GenS11{F0: (("foo bar" + (("\n\t" + ("hello" + "hello")) + "hello")) + "\n\t"), F1: true}
                                                    _ = v104
                                                case -39:
                                                    fmt.Println(v2)
                                                }
                                                v2 -= 700
                                            }
                                            _ = v96
                                        }
                                        v1 = (-(v1))
                                        _ = fi94
                                    }
                                default:
                                    if len(v4) > 0 {
                                        v4[0] = ""
                                    }
                                }
                            case "\n\t":
                                var v105 string = ("" + "\n\t")
                                _ = v105
                            default:
                                if ((80 < 127) && ((fi92 ^ (fi92 + (fi92 & (fi92 % (((fi92 ^ -868)) | 1))))) >= fi92)) {
                                    v106 := GenS8{F0: v3, F1: 630}
                                    _ = v106
                                } else {
                                    v1 = v1
                                    v93 = v93
                                }
                            }
                            v1 -= (v1 / ((1) | 1))
                            _ = v93
                        }
                        var v107 float64 = (((1.5 / 0.1) - 3.14159) + float64(1e10))
                        _ = fi92
                        _ = v107
                    }
                    var v108 uint8 = 0
                    _ = v108
                }
            case "":
                switch "a" {
                case "foo bar":
                    if (350 < 65535) {
                        if true {
                            var v109 int8 = 1
                            _ = v109
                        }
                    }
                }
            default:
                fmt.Println(v1)
            }
            v3 = (v3 / ((v3) | 1))
        }
        v110 := GenS9{F0: (!true)}
        if true {
            for fi111 := 0; fi111 < 14; fi111++ {
                for fi112 := 0; fi112 < 7; fi112++ {
                    v2 += (-(v2))
                    var v113 uint16 = 32767
                    if ((!true) || (23 <= 255)) {
                        v114 := GenS6{F0: 128, F1: v3, F2: (fi111 & fi112)}
                        _ = v114
                    }
                    _ = fi112
                    _ = v113
                }
                v3 += v3
                for fi115 := 0; fi115 < 5; fi115++ {
                    var v116 int = fi111
                    _ = fi115
                    _ = v116
                }
                _ = fi111
            }
        }
        _ = v110
    }
    v117 := GenS10{F0: v3, F1: true, F2: "foo bar"}
    fmt.Println(v1, v1)
    _ = len(v4)
    fmt.Println(v2)
    v118 := GenS2{F0: 0, F1: 216, F2: (!(-99 >= 127))}
    for fi119 := 0; fi119 < 18; fi119++ {
        for fi120 := 0; fi120 < 16; fi120++ {
            var v121 int8 = -128
            var v122 int16 = 202
            _ = fi120
            _ = v121
            _ = v122
        }
        v3 = (v3 ^ (^(v3)))
        if true {
            v2 = v2
            v4 = append(v4, ((("foo bar" + "hello") + "hello") + "\n\t"))
        }
        _ = fi119
    }
    v4 = append(v4, ("a" + ("" + "foo bar")))
    v1 = (v1 & v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v117
    _ = v118
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
