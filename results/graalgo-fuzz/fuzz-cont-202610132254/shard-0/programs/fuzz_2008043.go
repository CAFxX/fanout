package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int8; F1 uint16 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 bool; F1 int; F2 int }
type GenS8 struct { F0 int32; F1 int; F2 int64 }
type GenS9 struct { F0 uint32 }
type GenS10 struct { F0 bool; F1 string }
type GenS11 struct { F0 uint64; F1 bool; F2 int8 }
type GenS12 struct { F0 uint64; F1 string; F2 int64 }
type GenS13 struct { F0 string }
type GenS14 struct { F0 int16 }
type GenS15 struct { F0 int16; F1 bool; F2 int }
func genFn0(p0 uint) int32 {
    p0 = p0
    v4 := GenS1{F0: -583}
    for fi5 := 0; fi5 < 7; fi5++ {
        var v6 int = (fi5 >> ((954) & 63))
        p0 -= p0
        var v7 bool = true
        _ = fi5
        _ = v6
        _ = v7
    }
    for fi8 := 0; fi8 < 9; fi8++ {
        fmt.Println(fi8)
        if (!(4294967295 >= 0)) {
            for fi9 := 0; fi9 < 14; fi9++ {
                p0 = p0
                if (432 == 451) {
                    if false {
                        switch "foo bar" {
                        case "":
                            if (false || false) {
                                fmt.Println(fi9)
                            } else {
                                if (!(42 > 578)) {
                                    var v10 uint8 = 255
                                    _ = v10
                                } else {
                                    v11 := GenS1{F0: -9223372036854775808}
                                    p0 = 18446744073709551615
                                    v12 := []int{fi9, (fi9 + (fi9 / (((fi8 / (((fi8 - 0)) | 1))) | 1)))}
                                    if len(v12) >= 2 {
                                        v12 = v12[1:len(v12)]
                                    }
                                    if len(v12) >= 2 {
                                        v12 = v12[1:len(v12)]
                                    }
                                    _ = v11
                                    _ = v12
                                }
                                p0 = (p0 ^ (p0 | 18446744073709551615))
                            }
                        }
                    }
                    if false {
                        if true {
                            v13 := []int{fi8, (fi9 | (fi8 / (((^(fi8))) | 1))), fi9}
                            v14 := GenS1{F0: -1}
                            var v15 float64 = (((1.5 * float64(0.0)) * (((1e10 - (1e10 * 0.1)) / (0.1 + (0.0 + 1e10))) - (0.1 + (0.0 + (-2.25 * 1e10))))) * -2.25)
                            _ = v13
                            _ = v14
                            _ = v15
                        } else {
                            fmt.Println(p0)
                            v16 := GenS1{F0: -1}
                            if (((0 <= p0) || false) && (1 >= -1)) {
                                v17 := []int{4294967295, (fi9 / (((fi8 + fi8)) | 1)), 2147483647}
                                if ((-783 <= -1) || (false && (!true))) {
                                    p0 += 2
                                    p0 += p0
                                    p0 = (^(p0))
                                } else {
                                    var v18 float32 = (1e10 - (1e10 - 3.14159))
                                    _ = v18
                                }
                                if len(v17) > 0 {
                                    v17[0] = (^(fi8))
                                }
                                _ = v17
                            } else {
                                v19 := GenS2{F0: 1}
                                _ = v19
                            }
                            _ = v16
                        }
                    }
                    if (!(true && false)) {
                        var v20 bool = false
                        if v20 {
                            var v21 int16 = 523
                            v22 := GenS1{F0: 657}
                            _ = v21
                            _ = v22
                        } else {
                            if v20 {
                                switch 161 {
                                case 4294967295:
                                    switch (fi9 | fi9) {
                                    case 42:
                                        var v23 int16 = 42
                                        _ = v23
                                    case 4294967295:
                                        v24 := []int{}
                                        _ = v24
                                    default:
                                        var v25 float32 = 0.0
                                        if (-123 < 42) {
                                            v25 -= (1e10 - ((0.0 - -2.25) * float32(float64(v25))))
                                            if (!((((fi9 & -9223372036854775808) >= (fi9 << ((571) & 63))) && (3.14159 == 3.14159)) && true)) {
                                                fmt.Println(p0, v20, fi8)
                                                var v26 int = (fi9 & (fi9 >> (((^(p0))) & 63)))
                                                switch -113 {
                                                case 32767:
                                                    if (207 == -1) {
                                                        p0 = p0
                                                        p0 = 2
                                                        switch (fi8 & (fi9 << (((p0 << (((^(p0))) & 63))) & 63))) {
                                                        case -734:
                                                            if (444 >= 0) {
                                                                v27 := []int{fi8, 42, (fi8 - (v26 << ((138) & 63)))}
                                                                switch 2147483647 {
                                                                case 32767:
                                                                    v20 = (-2147483648 == 31)
                                                                }
                                                                switch ("foo bar" + ("a" + "a")) {
                                                                case "a":
                                                                    if len(v27) >= 2 {
                                                                        v27 = v27[1:len(v27)]
                                                                    }
                                                                case "foo bar":
                                                                    fmt.Println(v20, v25)
                                                                case "":
                                                                    if (!v20) {
                                                                        if false {
                                                                            var v28 uint = p0
                                                                            if false {
                                                                                switch fi8 {
                                                                                case 128:
                                                                                    switch ("\n\t" + ((("hello" + "hello") + ((("foo bar" + "hello") + "\n\t") + "a")) + "hello")) {
                                                                                    case "hello":
                                                                                        switch ("foo bar" + "\n\t") {
                                                                                        case "":
                                                                                            v29 := GenS1{F0: 42}
                                                                                            _ = v29
                                                                                        case "a":
                                                                                            p0 += 18446744073709551615
                                                                                            v30 := GenS3{F0: -12, F1: 1}
                                                                                            switch "\n\t" {
                                                                                            case "":
                                                                                                var v31 string = ((("\n\t" + ("foo bar" + ("" + ("\n\t" + "foo bar")))) + "\n\t") + "")
                                                                                                _ = v31
                                                                                            case "\n\t":
                                                                                                if len(v27) > 0 {
                                                                                                    v27[0] = v26
                                                                                                }
                                                                                            }
                                                                                            var v32 uint32 = uint32(p0)
                                                                                            _ = v30
                                                                                            _ = v32
                                                                                        case "foo bar":
                                                                                            if (32767 <= 42) {
                                                                                                switch "a" {
                                                                                                case "":
                                                                                                    var v33 float32 = 3.14159
                                                                                                    _ = v33
                                                                                                case "hello":
                                                                                                    var v34 int64 = -1
                                                                                                    _ = v34
                                                                                                default:
                                                                                                    switch (("" + "") + "") {
                                                                                                    case "\n\t":
                                                                                                        if (!false) {
                                                                                                            switch (fi9 % ((128) | 1)) {
                                                                                                            case 0:
                                                                                                                fmt.Println(v26, v25)
                                                                                                            case 42:
                                                                                                                if v20 {
                                                                                                                    switch "a" {
                                                                                                                    case "\n\t":
                                                                                                                        if len(v27) >= 2 {
                                                                                                                            v27 = v27[1:len(v27)]
                                                                                                                        }
                                                                                                                    case "hello":
                                                                                                                        var v35 float32 = v25
                                                                                                                        _ = v35
                                                                                                                    case "foo bar":
                                                                                                                        v36 := GenS4{F0: fi8}
                                                                                                                        _ = v36
                                                                                                                    default:
                                                                                                                        if false {
                                                                                                                            var v37 string = "hello"
                                                                                                                            if (fi8 != (fi8 + fi9)) {
                                                                                                                                if len(v27) > 0 {
                                                                                                                                    v27[0] = fi8
                                                                                                                                }
                                                                                                                                var v38 int16 = -32768
                                                                                                                                switch "" {
                                                                                                                                case "\n\t":
                                                                                                                                    v38 = (v38 * (v38 | (v38 << ((v28) & 63))))
                                                                                                                                    if ((v37 + "hello") <= ("foo bar" + v37)) {
                                                                                                                                        if (v38 >= v38) {
                                                                                                                                            fmt.Println(fi9)
                                                                                                                                            fmt.Println(p0)
                                                                                                                                        }
                                                                                                                                    }
                                                                                                                                case "foo bar":
                                                                                                                                    v25 = v25
                                                                                                                                }
                                                                                                                                _ = v38
                                                                                                                            } else {
                                                                                                                                var v39 float32 = v25
                                                                                                                                _ = v39
                                                                                                                            }
                                                                                                                            _ = v37
                                                                                                                        }
                                                                                                                    }
                                                                                                                } else {
                                                                                                                    v27 = append(v27, fi8)
                                                                                                                    switch fi9 {
                                                                                                                    case 4294967295:
                                                                                                                        var v40 uint16 = 515
                                                                                                                        _ = v40
                                                                                                                    }
                                                                                                                    v41 := GenS5{F0: 0}
                                                                                                                    _ = v41
                                                                                                                }
                                                                                                            default:
                                                                                                                v42 := GenS1{F0: 234}
                                                                                                                _ = v42
                                                                                                            }
                                                                                                            v28 = (^(v28))
                                                                                                        } else {
                                                                                                            v27 = append(v27, (fi9 - v26))
                                                                                                        }
                                                                                                    }
                                                                                                }
                                                                                                v25 += ((float32(1.5) / 0.1) + (((v25 / v25) - (float32(float64(float32(0.1))) + 3.14159)) - v25))
                                                                                            } else {
                                                                                                fmt.Println(p0, fi8, v28)
                                                                                                fmt.Println(fi8, v28, v26)
                                                                                            }
                                                                                        default:
                                                                                            v43 := GenS2{F0: 255}
                                                                                            switch (v26 << ((0) & 63)) {
                                                                                            case 12:
                                                                                                var v44 float64 = ((float64(v25) - 0.0) * float64(float32(0.1)))
                                                                                                _ = v44
                                                                                            case 2147483647:
                                                                                                if true {
                                                                                                    fmt.Println(fi8, v26)
                                                                                                }
                                                                                            }
                                                                                            var v45 int16 = -32768
                                                                                            _ = v43
                                                                                            _ = v45
                                                                                        }
                                                                                    case "":
                                                                                        var v46 uint32 = 2147483647
                                                                                        _ = v46
                                                                                    case "\n\t":
                                                                                        var v47 uint8 = 61
                                                                                        _ = v47
                                                                                    default:
                                                                                        v25 = (0.1 * float32(0.1))
                                                                                    }
                                                                                case 2147483647:
                                                                                    v25 = v25
                                                                                    if v20 {
                                                                                        var v48 float64 = 0.0
                                                                                        v28 += (p0 | 1)
                                                                                        p0 = 2
                                                                                        v48 = (v48 + v48)
                                                                                        if v20 {
                                                                                            p0 -= (p0 % (((p0 / (((v28 & p0)) | 1))) | 1))
                                                                                        }
                                                                                        _ = v48
                                                                                    }
                                                                                    var v49 float32 = (v25 + 0.0)
                                                                                    _ = v49
                                                                                }
                                                                                v26 = (fi9 * (fi8 / ((fi8) | 1)))
                                                                                if len(v27) > 0 {
                                                                                    v27[0] = (fi8 << (((^(p0))) & 63))
                                                                                }
                                                                            }
                                                                            v50 := GenS4{F0: -9223372036854775808}
                                                                            _ = v28
                                                                            _ = v50
                                                                        }
                                                                    } else {
                                                                        v51 := GenS1{F0: 0}
                                                                        var v52 uint64 = 1
                                                                        var v53 uint16 = 2
                                                                        _ = v51
                                                                        _ = v52
                                                                        _ = v53
                                                                    }
                                                                }
                                                                switch v26 {
                                                                case 2147483647:
                                                                    switch "foo bar" {
                                                                    case "\n\t":
                                                                        switch "" {
                                                                        case "\n\t":
                                                                            if len(v27) >= 2 {
                                                                                v27 = v27[1:len(v27)]
                                                                            }
                                                                        case "a":
                                                                            if false {
                                                                                switch v26 {
                                                                                case -1:
                                                                                    if len(v27) >= 2 {
                                                                                        v27 = v27[1:len(v27)]
                                                                                    }
                                                                                default:
                                                                                    if (100 == 127) {
                                                                                        v54 := GenS5{F0: -32768}
                                                                                        v27 = append(v27, -603)
                                                                                        _ = v54
                                                                                    } else {
                                                                                        var v55 int64 = -9223372036854775808
                                                                                        switch ("foo bar" + ("hello" + ("" + ("\n\t" + "a")))) {
                                                                                        case "\n\t":
                                                                                            p0 = (^(p0))
                                                                                        case "hello":
                                                                                            if v20 {
                                                                                                var v56 int8 = 0
                                                                                                _ = v56
                                                                                            }
                                                                                            _ = len(v27)
                                                                                        case "a":
                                                                                            v57 := GenS6{F0: (4294967295 > 1)}
                                                                                            _ = v57
                                                                                        default:
                                                                                            v58 := GenS7{F0: (-376 >= -1), F1: 2147483647, F2: fi8}
                                                                                            _ = v58
                                                                                        }
                                                                                        _ = v55
                                                                                    }
                                                                                }
                                                                                var v59 int8 = -84
                                                                                var v60 int = (-(fi9))
                                                                                _ = v59
                                                                                _ = v60
                                                                            } else {
                                                                                if len(v27) >= 2 {
                                                                                    v27 = v27[1:len(v27)]
                                                                                }
                                                                                if ((255 > 1) || (1 != 255)) {
                                                                                    v61 := GenS8{F0: 42, F1: 128, F2: -9223372036854775808}
                                                                                    var v62 bool = v20
                                                                                    v26 = fi8
                                                                                    _ = v61
                                                                                    _ = v62
                                                                                }
                                                                                v20 = v20
                                                                                var v63 float64 = 3.14159
                                                                                _ = v63
                                                                            }
                                                                        case "":
                                                                            v64 := GenS9{F0: 2147483647}
                                                                            _ = v64
                                                                        default:
                                                                            if true {
                                                                                switch (v26 & (fi9 + (^(fi9)))) {
                                                                                case -80:
                                                                                    var v65 float32 = ((1.5 - 1.5) + -2.25)
                                                                                    _ = v65
                                                                                default:
                                                                                    v66 := GenS9{F0: 1}
                                                                                    _ = v66
                                                                                }
                                                                                p0 = (p0 % (((p0 * 42)) | 1))
                                                                            }
                                                                        }
                                                                    }
                                                                case 127:
                                                                    if len(v27) > 0 {
                                                                        v27[0] = -687
                                                                    }
                                                                default:
                                                                    var v67 uint32 = 4294967295
                                                                    _ = v67
                                                                }
                                                                v26 += v26
                                                                _ = v27
                                                            } else {
                                                                var v68 int16 = 32767
                                                                _ = v68
                                                            }
                                                        case -780:
                                                            var v69 uint = (p0 | (p0 % ((84) | 1)))
                                                            _ = v69
                                                        default:
                                                            v70 := GenS6{F0: v20}
                                                            _ = v70
                                                        }
                                                    } else {
                                                        v71 := []int{(v26 % ((fi9) | 1)), (-(v26)), (v26 + -2147483648), fi8}
                                                        _ = v71
                                                    }
                                                default:
                                                    v72 := GenS10{F0: (2 > 968), F1: "\n\t"}
                                                    _ = v72
                                                }
                                                _ = v26
                                            } else {
                                                v20 = true
                                            }
                                            if (!v20) {
                                                var v73 bool = false
                                                if (!true) {
                                                    v74 := []int{9223372036854775807, fi8, fi8, (fi9 * (fi8 | fi8))}
                                                    _ = v74
                                                }
                                                _ = v73
                                            }
                                        } else {
                                            if v20 {
                                                v75 := GenS5{F0: 37}
                                                _ = v75
                                            } else {
                                                fmt.Println(v20, v20)
                                                p0 = p0
                                            }
                                            v76 := []int{(fi8 / ((fi8) | 1)), fi8, 0}
                                            v77 := GenS11{F0: 877, F1: (2 <= 2), F2: -51}
                                            _ = v76
                                            _ = v77
                                        }
                                        _ = v25
                                    }
                                case 128:
                                    if v20 {
                                        if ((v20 || v20) && true) {
                                            v20 = false
                                        } else {
                                            p0 = (p0 << ((0) & 63))
                                        }
                                        p0 = 953
                                    }
                                    var v78 int16 = 852
                                    v20 = true
                                    _ = v78
                                default:
                                    if v20 {
                                        var v79 bool = v20
                                        fmt.Println(v79, p0, p0)
                                        v80 := GenS12{F0: 0, F1: "a", F2: 1}
                                        _ = v79
                                        _ = v80
                                    } else {
                                        p0 = 960
                                        v20 = v20
                                    }
                                }
                            } else {
                                v20 = v20
                            }
                        }
                        _ = v20
                    }
                } else {
                    if false {
                        if false {
                            p0 = p0
                            v81 := GenS12{F0: 1, F1: "", F2: 0}
                            switch fi8 {
                            case 32767:
                                var v82 uint16 = 65535
                                _ = v82
                            default:
                                v83 := GenS13{F0: ""}
                                _ = v83
                            }
                            _ = v81
                        } else {
                            p0 -= (p0 * 2)
                            switch (("foo bar" + (("\n\t" + "\n\t") + "hello")) + "a") {
                            case "a":
                                switch (((("a" + "a") + "\n\t") + "foo bar") + "foo bar") {
                                case "\n\t":
                                    if true {
                                        switch -508 {
                                        case 32767:
                                            if false {
                                                v84 := []int{fi8, fi8}
                                                if len(v84) > 0 {
                                                    v84[0] = 461
                                                }
                                                fmt.Println(fi8, fi8)
                                                _ = len(v84)
                                                fmt.Println(p0, p0)
                                                _ = v84
                                            } else {
                                                v85 := []int{42, fi9, 101}
                                                _ = v85
                                            }
                                        case -2147483648:
                                            v86 := GenS12{F0: 42, F1: "", F2: 1}
                                            v87 := []int{32768, fi8}
                                            if (930 != 46) {
                                                switch 128 {
                                                case -322:
                                                    fmt.Println(fi8)
                                                }
                                                p0 = p0
                                            } else {
                                                if false {
                                                    if (false || ((float64(-2.25) - (1e10 / 1.0)) == 3.14159)) {
                                                        if (!false) {
                                                            var v88 int = (fi9 << ((65535) & 63))
                                                            fmt.Println(p0)
                                                            p0 += (p0 + p0)
                                                            if true {
                                                                p0 -= p0
                                                                p0 -= 65535
                                                                var v89 int32 = 2147483647
                                                                _ = v89
                                                            } else {
                                                                v90 := GenS7{F0: false, F1: (fi9 | 1), F2: 9223372036854775807}
                                                                _ = v90
                                                            }
                                                            if len(v87) > 0 {
                                                                v87[0] = fi8
                                                            }
                                                            v88 += 403
                                                            _ = v88
                                                        }
                                                    } else {
                                                        if (p0 < (p0 * 42)) {
                                                            var v91 int64 = 660
                                                            fmt.Println(p0, p0, v91)
                                                            if false {
                                                                var v92 float32 = 0.0
                                                                p0 += 18446744073709551615
                                                                _ = v92
                                                            }
                                                            if (814 == 548) {
                                                                p0 = p0
                                                            } else {
                                                                if (true && false) {
                                                                    switch "foo bar" {
                                                                    case "a":
                                                                        p0 += 65535
                                                                    case "\n\t":
                                                                        if len(v87) > 0 {
                                                                            v87[0] = 1
                                                                        }
                                                                    }
                                                                    if true {
                                                                        var v93 uint64 = 163
                                                                        _ = v93
                                                                    } else {
                                                                        fmt.Println(fi8, fi8)
                                                                    }
                                                                } else {
                                                                    var v94 int = fi9
                                                                    _ = v94
                                                                }
                                                                fmt.Println(p0)
                                                            }
                                                            _ = v91
                                                        }
                                                        var v95 uint16 = 0
                                                        _ = v95
                                                    }
                                                    if false {
                                                        var v96 uint32 = 2147483648
                                                        p0 -= (p0 << (((p0 / (((p0 * (p0 | 42))) | 1))) & 63))
                                                        _ = v96
                                                    }
                                                    p0 += 18446744073709551615
                                                }
                                                v97 := GenS4{F0: fi9}
                                                _ = v97
                                            }
                                            _ = v86
                                            _ = v87
                                        }
                                        var v98 string = "foo bar"
                                        _ = v98
                                    }
                                case "":
                                    if (!(fi8 != 32767)) {
                                        v99 := []int{}
                                        _ = v99
                                    } else {
                                        switch ("a" + "a") {
                                        case "":
                                            fmt.Println(fi8, fi9)
                                        case "foo bar":
                                            v100 := []int{9223372036854775807, (fi9 % ((fi9) | 1)), fi9, 9223372036854775807}
                                            _ = v100
                                        case "\n\t":
                                            v101 := []int{}
                                            _ = v101
                                        default:
                                            p0 = (p0 - 18446744073709551615)
                                        }
                                    }
                                default:
                                    p0 = (^(p0))
                                }
                                var v102 string = "\n\t"
                                v103 := GenS14{F0: -49}
                                _ = v102
                                _ = v103
                            default:
                                p0 = p0
                            }
                        }
                    } else {
                        p0 += p0
                    }
                    var v104 float64 = (float64(0.1) / -2.25)
                    if ((("a" + (("hello" + "foo bar") + ("" + "\n\t"))) < "") && (!true)) {
                        v105 := []int{}
                        _ = v105
                    }
                    var v106 bool = ((42 <= 4294967295) || false)
                    _ = v104
                    _ = v106
                }
                _ = fi9
            }
            p0 -= (p0 ^ (p0 ^ 65535))
            p0 = (p0 / ((p0) | 1))
        } else {
            v107 := []int{}
            var v108 uint32 = 618
            _ = v107
            _ = v108
        }
        _ = fi8
    }
    _ = p0
    _ = v4
    return 0
}

func main() {
    v1 := make(map[int]bool)
    var v2 uint32 = 65535
    v3 := make(map[uint32]bool)
    v109 := GenS15{F0: -1, F1: true, F2: 127}
    fmt.Println(v2)
    var v110 int = 127
    for fi111 := 0; fi111 < 14; fi111++ {
        var v112 uint = 42
        if ((v112 | 18446744073709551615) <= (v112 >> ((1) & 63))) {
            v113 := GenS5{F0: 32767}
            _ = v113
        }
        _ = fi111
        _ = v112
    }
    var v114 uint16 = 350
    fmt.Println(v2)
    v115 := v3[(v2 / ((1) | 1))]
    v116 := GenS4{F0: v110}
    switch "a" {
    case "hello":
        var v117 string = (((("" + ("foo bar" + "a")) + ("foo bar" + "a")) + "\n\t") + "\n\t")
        _ = v117
    case "":
        if true {
            v118 := GenS9{F0: (^(v2))}
            fmt.Println(v114, v115, v2)
            v119 := GenS13{F0: "\n\t"}
            _ = v118
            _ = v119
        }
    case "foo bar":
        var v120 string = "a"
        _ = v120
    default:
        var v121 int16 = 42
        _ = v121
    }
    if false {
        for fi122 := 0; fi122 < 6; fi122++ {
            v2 = (v2 | v2)
            _ = fi122
        }
    } else {
        v123 := GenS1{F0: -1}
        fmt.Println(v114, v2, v2)
        switch "a" {
        case "\n\t":
            v2 = v2
        default:
            for fi124 := 0; fi124 < 3; fi124++ {
                var v125 uint = 332
                _ = fi124
                _ = v125
            }
        }
        _ = v123
    }
    v115 = (42 < 1)
    var v126 int32 = -2147483648
    v2 -= (v2 >> ((229) & 63))
    for fi127 := 0; fi127 < 20; fi127++ {
        for fi128 := 0; fi128 < 14; fi128++ {
            var v129 float64 = ((float64((float32((1.5 * ((1.5 + 1e10) / 1e10))) - 0.1)) - (0.0 / 1.0)) * 0.0)
            _ = fi128
            _ = v129
        }
        _ = fi127
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v110:", v110)
    fmt.Println("v114:", v114)
    fmt.Println("v115:", v115)
    fmt.Println("v126:", v126)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v109
    _ = v110
    _ = v114
    _ = v115
    _ = v116
    _ = v126
    _ = v2
    _ = v3
}
