package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64; F2 int16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int32; F1 int32 }
type GenS4 struct { F0 uint16; F1 uint32 }
type GenS5 struct { F0 uint32; F1 int32; F2 int }
type GenS6 struct { F0 int64; F1 string; F2 int }
type GenS7 struct { F0 uint; F1 bool; F2 uint32 }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 int8 }
type GenS10 struct { F0 int8; F1 uint; F2 int16 }
type GenS11 struct { F0 int; F1 uint32 }
type GenS12 struct { F0 uint32 }
type GenS13 struct { F0 uint64; F1 int32 }
type GenS14 struct { F0 uint64; F1 uint8; F2 uint16 }
type GenS15 struct { F0 uint; F1 uint16; F2 int16 }
type GenS16 struct { F0 uint; F1 int64; F2 bool }
type GenS17 struct { F0 uint8; F1 int64; F2 int32 }
type GenS18 struct { F0 int64; F1 int; F2 uint }
type GenS19 struct { F0 uint }
type GenS20 struct { F0 bool }
type GenS21 struct { F0 uint64; F1 uint }
type GenS22 struct { F0 int; F1 int16; F2 string }
type GenS23 struct { F0 int }
type GenS24 struct { F0 uint16 }
type GenS25 struct { F0 uint32; F1 uint32 }

func main() {
    v1 := []uint32{2147483648, 65535, 4294967295}
    v2 := make(map[string]uint32)
    var v3 int8 = 42
    v4 := []string{("\n\t" + "\n\t"), ("hello" + "a")}
    var v5 int8 = (-(v3))
    v6 := []uint32{0}
    fmt.Println(v5, v5)
    for fi7 := 0; fi7 < 14; fi7++ {
        v8 := GenS1{F0: -364, F1: 18446744073709551615, F2: 1000}
        v5 -= 100
        if false {
            if (-1 < -1) {
                v9 := v2["hello"]
                if (!(65535 > 65535)) {
                    fmt.Println(fi7, v5, v3)
                    var v10 uint = 524
                    v3 += (v3 + v3)
                    _ = v10
                }
                for fi11 := 0; fi11 < 8; fi11++ {
                    if false {
                        if true {
                            var v12 uint32 = v9
                            if false {
                                switch "hello" {
                                case "\n\t":
                                    if true {
                                        switch "a" {
                                        case "":
                                            var v13 float32 = -2.25
                                            _ = v13
                                        case "hello":
                                            var v14 int64 = -9223372036854775808
                                            _ = v14
                                        case "\n\t":
                                            var v15 uint64 = 2
                                            _ = v15
                                        }
                                        switch "\n\t" {
                                        case "foo bar":
                                            var v16 uint64 = 2
                                            _ = v16
                                        case "\n\t":
                                            if (true || (true && ((1e10 + ((3.14159 - (1e10 * (((((float64((((-2.25 - 0.0) * 0.0) - (1.5 - 3.14159))) - 0.1) - 1e10) + 1e10) / 0.1) * 3.14159))) * 0.1)) >= float64(0.0)))) {
                                                v5 -= 100
                                                v3 = v3
                                                v12 += v12
                                            } else {
                                                v17 := GenS1{F0: 1, F1: 42, F2: 0}
                                                v3 = v5
                                                var v18 uint = 4294967295
                                                var v19 uint = (v18 ^ 0)
                                                _ = v17
                                                _ = v18
                                                _ = v19
                                            }
                                        default:
                                            v20 := GenS1{F0: -9223372036854775808, F1: 18446744073709551615, F2: 687}
                                            _ = v20
                                        }
                                        if true {
                                            v5 += 1
                                            v9 += (v9 >> ((uint(v12)) & 63))
                                            v12 = v12
                                        }
                                    } else {
                                        v3 = v5
                                        var v21 int16 = -1
                                        _ = v21
                                    }
                                    fmt.Println(fi11, fi7, fi7)
                                    if (!(((128 <= 127) || false) || (337 < -32768))) {
                                        delete(v2, "foo bar")
                                        v3 = (v3 - v3)
                                        v22 := GenS1{F0: 0, F1: 4294967295, F2: 1}
                                        _ = v22
                                    }
                                default:
                                    var v23 uint = 4294967295
                                    _ = v23
                                }
                                v24 := v2["hello"]
                                _ = v24
                            }
                            var v25 int8 = v5
                            _ = v12
                            _ = v25
                        } else {
                            if (0 > -506) {
                                switch "hello" {
                                case "hello":
                                    v26 := GenS1{F0: -1, F1: 18446744073709551615, F2: 404}
                                    _ = v26
                                default:
                                    var v27 int = fi11
                                    _ = v27
                                }
                                switch ("" + "a") {
                                case "\n\t":
                                    if false {
                                        v9 -= v9
                                        v5 += 42
                                    } else {
                                        v9 += (^(v9))
                                        fmt.Println(v5)
                                        var v28 int8 = (v5 / (((v5 + -5)) | 1))
                                        _ = v28
                                    }
                                    switch (fi7 % (((fi11 << ((132) & 63))) | 1)) {
                                    case 128:
                                        var v29 uint8 = 255
                                        _ = v29
                                    case -787:
                                        var v30 uint8 = 1
                                        _ = v30
                                    }
                                    if false {
                                        v31 := v2[(("hello" + "a") + "hello")]
                                        switch "a" {
                                        case "a":
                                            v32 := GenS2{F0: ""}
                                            _ = v32
                                        case "hello":
                                            if true {
                                                fmt.Println(v5, v3, fi11)
                                            } else {
                                                v5 -= v3
                                                v3 = v5
                                            }
                                        }
                                        _ = v31
                                    } else {
                                        _ = len(v2)
                                    }
                                case "foo bar":
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                    v5 = (v3 << ((192) & 63))
                                    v33 := GenS3{F0: 42, F1: -1}
                                    _ = v33
                                default:
                                    var v34 int8 = 42
                                    _ = v34
                                }
                                switch fi11 {
                                case -971:
                                    if ((true || (1 != (v9 | v9))) && true) {
                                        v35 := GenS1{F0: -76, F1: 244, F2: 869}
                                        if (47 != 314) {
                                            var v36 uint64 = 1
                                            _ = v36
                                        } else {
                                            v37 := GenS4{F0: 656, F1: 1}
                                            fmt.Println(fi7, fi7, v3)
                                            if false {
                                                switch "\n\t" {
                                                case "\n\t":
                                                    fmt.Println(fi11)
                                                case "":
                                                    fmt.Println(v9, v5, fi11)
                                                }
                                                switch ("" + "\n\t") {
                                                case "foo bar":
                                                    v5 -= (^(v3))
                                                case "\n\t":
                                                    var v38 int32 = 0
                                                    v3 -= 26
                                                    _ = v38
                                                default:
                                                    var v39 int32 = 42
                                                    _ = v39
                                                }
                                                v40 := v2[((("a" + "") + "a") + "foo bar")]
                                                _ = v40
                                            } else {
                                                v41 := GenS1{F0: 42, F1: 42, F2: -1}
                                                _ = v41
                                            }
                                            _ = v37
                                        }
                                        switch ((("a" + "foo bar") + "hello") + "foo bar") {
                                        case "foo bar":
                                            var v42 string = ("a" + "a")
                                            var v43 int = fi11
                                            _ = v42
                                            _ = v43
                                        case "\n\t":
                                            v44 := GenS5{F0: (v9 | v9), F1: 713, F2: fi7}
                                            _ = v44
                                        default:
                                            v45 := GenS6{F0: -9223372036854775808, F1: "a", F2: 2147483647}
                                            _ = v45
                                        }
                                        _ = v35
                                    } else {
                                        var v46 uint8 = 255
                                        switch "" {
                                        case "":
                                            switch (fi7 ^ fi11) {
                                            case -781:
                                                switch fi7 {
                                                case -727:
                                                    if len(v4) > 0 {
                                                        v4[0] = (("hello" + ("a" + "hello")) + (("foo bar" + ("" + "")) + (("foo bar" + "a") + "hello")))
                                                    }
                                                case 127:
                                                    switch "a" {
                                                    case "foo bar":
                                                        switch -193 {
                                                        case -2147483648:
                                                            var v47 float64 = 1e10
                                                            _ = v47
                                                        case 32768:
                                                            v3 = v5
                                                        case 127:
                                                            var v48 uint8 = (^(v46))
                                                            _ = v48
                                                        default:
                                                            if (true && (-2.25 > 1.5)) {
                                                                switch "" {
                                                                case "hello":
                                                                    switch fi7 {
                                                                    case -1:
                                                                        fmt.Println(v9, v5, fi11)
                                                                    case -2147483648:
                                                                        switch "\n\t" {
                                                                        case "a":
                                                                            if true {
                                                                                v5 = (^(v5))
                                                                                v5 -= (v3 % ((v5) | 1))
                                                                            }
                                                                        }
                                                                    default:
                                                                        if (2 > 2) {
                                                                            var v49 uint = 4294967295
                                                                            fmt.Println(v46, fi11, fi7)
                                                                            _ = v49
                                                                        } else {
                                                                            var v50 uint8 = (^(v46))
                                                                            var v51 uint32 = (v9 / (((v9 - (v9 | v9))) | 1))
                                                                            _ = v50
                                                                            _ = v51
                                                                        }
                                                                    }
                                                                case "foo bar":
                                                                    if (-770 >= 264) {
                                                                        v52 := GenS4{F0: 1, F1: 417}
                                                                        _ = v52
                                                                    } else {
                                                                        if true {
                                                                            v9 += v9
                                                                            v9 = v9
                                                                            var v53 uint64 = 1
                                                                            switch fi7 {
                                                                            case 2147483647:
                                                                                switch fi11 {
                                                                                case 443:
                                                                                    v9 -= (^(v9))
                                                                                default:
                                                                                    if len(v6) > 0 {
                                                                                        v6[0] = v9
                                                                                    }
                                                                                }
                                                                                v9 = v9
                                                                                v2["hello"] = (^(v9))
                                                                                v5 -= v3
                                                                                fmt.Println(fi7, v46, v9)
                                                                            case 32768:
                                                                                var v54 uint32 = (v9 - v9)
                                                                                _ = v54
                                                                            }
                                                                            delete(v2, ("hello" + ("foo bar" + "\n\t")))
                                                                            _ = v53
                                                                        } else {
                                                                            v9 = v9
                                                                            var v55 uint32 = 65535
                                                                            v56 := GenS4{F0: 807, F1: 0}
                                                                            switch ((("" + ("" + "a")) + "hello") + "a") {
                                                                            case "":
                                                                                var v57 float64 = 3.14159
                                                                                _ = v57
                                                                            case "hello":
                                                                                v55 += 1
                                                                            default:
                                                                                switch 0 {
                                                                                case 889:
                                                                                    var v58 int32 = -230
                                                                                    _ = v58
                                                                                }
                                                                            }
                                                                            if (!true) {
                                                                                if true {
                                                                                    var v59 uint64 = 18446744073709551615
                                                                                    delete(v2, "hello")
                                                                                    _ = v59
                                                                                } else {
                                                                                    var v60 uint64 = 42
                                                                                    var v61 uint64 = v60
                                                                                    _ = v60
                                                                                    _ = v61
                                                                                }
                                                                            } else {
                                                                                v3 = (v5 % (((v5 % ((v3) | 1))) | 1))
                                                                                var v62 uint64 = 18446744073709551615
                                                                                _ = len(v2)
                                                                                _ = v62
                                                                            }
                                                                            var v63 uint64 = uint64(v9)
                                                                            delete(v2, "foo bar")
                                                                            _ = v55
                                                                            _ = v56
                                                                            _ = v63
                                                                        }
                                                                        var v64 int32 = -204
                                                                        _ = v64
                                                                    }
                                                                    if true {
                                                                        var v65 uint8 = v46
                                                                        if ((v5 * 0) != v3) {
                                                                            fmt.Println(v3, v65, fi7)
                                                                            var v66 float64 = (1e10 - ((3.14159 - -2.25) / 1.5))
                                                                            _ = v66
                                                                        }
                                                                        v46 += 1
                                                                        if true {
                                                                            v65 = (^(v46))
                                                                            if false {
                                                                                if (fi7 == fi7) {
                                                                                    _ = len(v2)
                                                                                    v67 := GenS7{F0: 527, F1: false, F2: 42}
                                                                                    v3 += int8(v46)
                                                                                    _ = v67
                                                                                } else {
                                                                                    if true {
                                                                                        var v68 float64 = ((-2.25 + 3.14159) + (-2.25 - float64(0.0)))
                                                                                        switch 128 {
                                                                                        case -170:
                                                                                            v5 += (v5 & v3)
                                                                                        default:
                                                                                            v5 = v3
                                                                                        }
                                                                                        v69 := GenS2{F0: "foo bar"}
                                                                                        _ = v68
                                                                                        _ = v69
                                                                                    } else {
                                                                                        var v70 int32 = 619
                                                                                        if len(v4) >= 2 {
                                                                                            v4 = v4[1:len(v4)]
                                                                                        }
                                                                                        _ = v70
                                                                                    }
                                                                                }
                                                                                if (("\n\t" + "foo bar") > "hello") {
                                                                                    v5 += (^(v3))
                                                                                    v71 := GenS4{F0: 42, F1: 2147483648}
                                                                                    v46 = 2
                                                                                    _ = v71
                                                                                } else {
                                                                                    var v72 uint = 414
                                                                                    v73 := GenS3{F0: 1, F1: 541}
                                                                                    if true {
                                                                                        var v74 bool = (false && true)
                                                                                        _ = len(v2)
                                                                                        _ = v74
                                                                                    }
                                                                                    _ = v72
                                                                                    _ = v73
                                                                                }
                                                                            }
                                                                        } else {
                                                                            v75 := GenS6{F0: -26, F1: "", F2: fi11}
                                                                            var v76 uint8 = v46
                                                                            switch (fi7 % ((4294967295) | 1)) {
                                                                            case 1:
                                                                                switch "" {
                                                                                case "hello":
                                                                                    fmt.Println(fi11, v5, v76)
                                                                                case "foo bar":
                                                                                    var v77 uint8 = (v46 - v46)
                                                                                    _ = v77
                                                                                }
                                                                            case 184:
                                                                                var v78 int32 = 42
                                                                                _ = v78
                                                                            case 128:
                                                                                fmt.Println(fi11, v5, v76)
                                                                            }
                                                                            _ = v75
                                                                            _ = v76
                                                                        }
                                                                        _ = v65
                                                                    } else {
                                                                        fmt.Println(v9, fi7)
                                                                    }
                                                                    fmt.Println(v5, fi7)
                                                                    _ = len(v2)
                                                                case "":
                                                                    if false {
                                                                        v3 = v3
                                                                        var v79 string = ("hello" + "hello")
                                                                        _ = v79
                                                                    } else {
                                                                        if (18446744073709551615 > 255) {
                                                                            delete(v2, ("a" + "foo bar"))
                                                                        }
                                                                        v5 = (v3 & (v3 | v3))
                                                                        if true {
                                                                            if false {
                                                                                if (true || ((!(v5 >= v3)) || true)) {
                                                                                    v5 = (v5 * (v5 | (-(v3))))
                                                                                    v46 = (^(v46))
                                                                                    var v80 int8 = v5
                                                                                    v81 := GenS2{F0: "hello"}
                                                                                    var v82 int8 = -124
                                                                                    v5 = -105
                                                                                    _ = v80
                                                                                    _ = v81
                                                                                    _ = v82
                                                                                } else {
                                                                                    switch "\n\t" {
                                                                                    case "a":
                                                                                        switch -2147483648 {
                                                                                        case -1:
                                                                                            if ((float32(1e10) - -2.25) > (3.14159 + 0.1)) {
                                                                                                v46 += (v46 & (v46 * (v46 * 128)))
                                                                                                v46 += v46
                                                                                                switch (^(fi7)) {
                                                                                                case 0:
                                                                                                    var v83 float32 = float32(0.0)
                                                                                                    _ = v83
                                                                                                }
                                                                                            }
                                                                                        case 127:
                                                                                            if false {
                                                                                                v5 = -128
                                                                                                fmt.Println(v3, v9, v9)
                                                                                            }
                                                                                        default:
                                                                                            if len(v1) >= 2 {
                                                                                                v1 = v1[1:len(v1)]
                                                                                            }
                                                                                        }
                                                                                    default:
                                                                                        var v84 uint32 = v9
                                                                                        var v85 uint32 = 626
                                                                                        fmt.Println(v5, v9, v46)
                                                                                        _ = v84
                                                                                        _ = v85
                                                                                    }
                                                                                    if true {
                                                                                        var v86 int32 = 304
                                                                                        _ = v86
                                                                                    }
                                                                                    if len(v1) > 0 {
                                                                                        v1[0] = v9
                                                                                    }
                                                                                }
                                                                                var v87 int32 = 42
                                                                                _ = v87
                                                                            }
                                                                            if true {
                                                                                v2[(("foo bar" + (("foo bar" + ((("a" + "foo bar") + "a") + "")) + "hello")) + "foo bar")] = (^(v9))
                                                                                if ((v46 >> ((0) & 63)) < 2) {
                                                                                    switch fi7 {
                                                                                    case 32768:
                                                                                        if (true && true) {
                                                                                            fmt.Println(fi11)
                                                                                            switch "hello" {
                                                                                            case "\n\t":
                                                                                                var v88 int8 = 42
                                                                                                _ = v88
                                                                                            case "foo bar":
                                                                                                var v89 int64 = 0
                                                                                                _ = v89
                                                                                            case "":
                                                                                                if (-2147483648 < 2147483647) {
                                                                                                    v9 -= (v9 * (v9 & 0))
                                                                                                    v2["foo bar"] = v9
                                                                                                    fmt.Println(v9, fi7, fi7)
                                                                                                    if (737 != uint16(v9)) {
                                                                                                        if true {
                                                                                                            if len(v1) >= 2 {
                                                                                                                v1 = v1[1:len(v1)]
                                                                                                            }
                                                                                                            v90 := v2["hello"]
                                                                                                            _ = v90
                                                                                                        }
                                                                                                    }
                                                                                                }
                                                                                            }
                                                                                            switch fi11 {
                                                                                            case 393:
                                                                                                switch fi7 {
                                                                                                case 249:
                                                                                                    v91 := v2[("\n\t" + "\n\t")]
                                                                                                    _ = v91
                                                                                                case 127:
                                                                                                    v5 += v5
                                                                                                    if len(v6) >= 2 {
                                                                                                        v6 = v6[1:len(v6)]
                                                                                                    }
                                                                                                case 32768:
                                                                                                    if false {
                                                                                                        v2["a"] = (v9 + (v9 ^ v9))
                                                                                                        v92 := GenS7{F0: 255, F1: (1 >= 65535), F2: 4294967295}
                                                                                                        if true {
                                                                                                            v9 += v9
                                                                                                            var v93 uint = 42
                                                                                                            switch (fi7 | 0) {
                                                                                                            case -2147483648:
                                                                                                                switch (fi7 % (((-(fi11))) | 1)) {
                                                                                                                case -959:
                                                                                                                    var v94 int32 = -664
                                                                                                                    _ = v94
                                                                                                                case 128:
                                                                                                                    v3 += v5
                                                                                                                    if len(v1) > 0 {
                                                                                                                        v1[0] = v9
                                                                                                                    }
                                                                                                                default:
                                                                                                                    delete(v2, (("hello" + "foo bar") + ("foo bar" + ("hello" + (("" + ((((("a" + "") + (("" + "\n\t") + ("foo bar" + (("a" + "\n\t") + "\n\t")))) + "a") + ((("hello" + "a") + "hello") + "a")) + "a")) + ((("foo bar" + ("foo bar" + "a")) + "a") + (("a" + "a") + "foo bar")))))))
                                                                                                                }
                                                                                                            default:
                                                                                                                var v95 int8 = v3
                                                                                                                fmt.Println(v95, v5)
                                                                                                                delete(v2, "\n\t")
                                                                                                                if (0 > 817) {
                                                                                                                    if false {
                                                                                                                        v96 := GenS4{F0: 485, F1: (^(v9))}
                                                                                                                        _ = v96
                                                                                                                    }
                                                                                                                }
                                                                                                                _ = v95
                                                                                                            }
                                                                                                            _ = v93
                                                                                                        } else {
                                                                                                            v3 -= (v5 * v3)
                                                                                                            fmt.Println(v3)
                                                                                                            v46 -= (v46 / ((v46) | 1))
                                                                                                        }
                                                                                                        _ = v92
                                                                                                    }
                                                                                                default:
                                                                                                    v6 = append(v6, 2)
                                                                                                }
                                                                                            case 1:
                                                                                                if false {
                                                                                                    var v97 uint32 = (v9 & (^(v9)))
                                                                                                    _ = v97
                                                                                                } else {
                                                                                                    v98 := GenS2{F0: ("a" + "\n\t")}
                                                                                                    switch "foo bar" {
                                                                                                    case "hello":
                                                                                                        if true {
                                                                                                            if (!(65535 < 0)) {
                                                                                                                if true {
                                                                                                                    if false {
                                                                                                                        if true {
                                                                                                                            if (true || (!true)) {
                                                                                                                                v5 = v5
                                                                                                                            } else {
                                                                                                                                var v99 uint64 = 2
                                                                                                                                _ = v99
                                                                                                                            }
                                                                                                                        }
                                                                                                                        if len(v1) > 0 {
                                                                                                                            v1[0] = (^(v9))
                                                                                                                        }
                                                                                                                        v3 = v3
                                                                                                                    } else {
                                                                                                                        fmt.Println(v46)
                                                                                                                        switch fi7 {
                                                                                                                        case 127:
                                                                                                                            _ = len(v2)
                                                                                                                        case 4294967295:
                                                                                                                            switch 32768 {
                                                                                                                            case 0:
                                                                                                                                v6 = append(v6, v9)
                                                                                                                            default:
                                                                                                                                v3 = 42
                                                                                                                            }
                                                                                                                        case 128:
                                                                                                                            switch (("hello" + "hello") + "") {
                                                                                                                            case "a":
                                                                                                                                if (9223372036854775807 <= 1) {
                                                                                                                                    delete(v2, "a")
                                                                                                                                }
                                                                                                                            default:
                                                                                                                                if false {
                                                                                                                                    v2["a"] = (v9 | (v9 / (((^(v9))) | 1)))
                                                                                                                                    v100 := v2[""]
                                                                                                                                    var v101 bool = true
                                                                                                                                    _ = v100
                                                                                                                                    _ = v101
                                                                                                                                }
                                                                                                                            }
                                                                                                                        }
                                                                                                                        v2["\n\t"] = (v9 / ((2147483647) | 1))
                                                                                                                        v4 = append(v4, "hello")
                                                                                                                    }
                                                                                                                    var v102 bool = (((!((true || false) || false)) || (!(4294967295 == 330))) && (1000 > -1))
                                                                                                                    switch fi11 {
                                                                                                                    case -744:
                                                                                                                        if v102 {
                                                                                                                            fmt.Println(v5, v46)
                                                                                                                            fmt.Println(v3, v3)
                                                                                                                            delete(v2, "a")
                                                                                                                            v103 := GenS8{F0: true}
                                                                                                                            if v102 {
                                                                                                                                var v104 uint16 = 2
                                                                                                                                _ = v104
                                                                                                                            } else {
                                                                                                                                fmt.Println(fi7)
                                                                                                                            }
                                                                                                                            var v105 uint64 = 42
                                                                                                                            _ = v103
                                                                                                                            _ = v105
                                                                                                                        } else {
                                                                                                                            v46 = (v46 << ((18446744073709551615) & 63))
                                                                                                                        }
                                                                                                                    default:
                                                                                                                        switch -2147483648 {
                                                                                                                        case 32768:
                                                                                                                            fmt.Println(fi7)
                                                                                                                        case 896:
                                                                                                                            v106 := v2["a"]
                                                                                                                            _ = v106
                                                                                                                        default:
                                                                                                                            _ = len(v2)
                                                                                                                        }
                                                                                                                    }
                                                                                                                    _ = v102
                                                                                                                } else {
                                                                                                                    fmt.Println(v3, fi11, fi11)
                                                                                                                    if true {
                                                                                                                        var v107 uint32 = (v9 * uint32(v46))
                                                                                                                        _ = v107
                                                                                                                    }
                                                                                                                    fmt.Println(fi11, v9, fi7)
                                                                                                                    var v108 int64 = -1
                                                                                                                    _ = v108
                                                                                                                }
                                                                                                            } else {
                                                                                                                fmt.Println(v9)
                                                                                                                v46 -= 154
                                                                                                                if false {
                                                                                                                    fmt.Println(v9)
                                                                                                                    v46 = v46
                                                                                                                }
                                                                                                            }
                                                                                                            v5 -= (v5 >> ((4294967295) & 63))
                                                                                                            var v109 int64 = 1
                                                                                                            _ = v109
                                                                                                        } else {
                                                                                                            if (1 != -1) {
                                                                                                                var v110 int32 = -713
                                                                                                                switch fi11 {
                                                                                                                case 32768:
                                                                                                                    var v111 int32 = (v110 >> ((794) & 63))
                                                                                                                    _ = v111
                                                                                                                case 0:
                                                                                                                    var v112 uint8 = v46
                                                                                                                    _ = v112
                                                                                                                }
                                                                                                                _ = v110
                                                                                                            }
                                                                                                        }
                                                                                                    case "a":
                                                                                                        if false {
                                                                                                            v113 := v2["\n\t"]
                                                                                                            v3 = 8
                                                                                                            v5 = (v3 % ((v5) | 1))
                                                                                                            v114 := GenS2{F0: "hello"}
                                                                                                            delete(v2, "hello")
                                                                                                            _ = v113
                                                                                                            _ = v114
                                                                                                        } else {
                                                                                                            switch (fi11 >> ((42) & 63)) {
                                                                                                            case 562:
                                                                                                                var v115 int = 9223372036854775807
                                                                                                                _ = v115
                                                                                                            case 222:
                                                                                                                var v116 uint64 = 0
                                                                                                                _ = v116
                                                                                                            case 4294967295:
                                                                                                                switch "a" {
                                                                                                                case "a":
                                                                                                                    v117 := GenS3{F0: -1, F1: -331}
                                                                                                                    _ = v117
                                                                                                                case "\n\t":
                                                                                                                    if true {
                                                                                                                        if (((float64((0.1 / 0.1)) - 3.14159) < (((1.5 - (-2.25 + 0.1)) + 0.1) / float64((0.0 + (1.5 * ((float32(-2.25) - 0.0) / -2.25)))))) || false) {
                                                                                                                            v118 := GenS2{F0: ("a" + "\n\t")}
                                                                                                                            v119 := GenS3{F0: -340, F1: -45}
                                                                                                                            var v120 int64 = 1
                                                                                                                            _ = v118
                                                                                                                            _ = v119
                                                                                                                            _ = v120
                                                                                                                        } else {
                                                                                                                            v3 = v3
                                                                                                                            v121 := GenS5{F0: uint32(v46), F1: -2147483648, F2: 4294967295}
                                                                                                                            _ = v121
                                                                                                                        }
                                                                                                                    }
                                                                                                                case "foo bar":
                                                                                                                    v122 := GenS2{F0: ""}
                                                                                                                    _ = v122
                                                                                                                }
                                                                                                            default:
                                                                                                                var v123 bool = (!false)
                                                                                                                _ = v123
                                                                                                            }
                                                                                                        }
                                                                                                    case "foo bar":
                                                                                                        if ((false || (("" >= "\n\t") && false)) || (1 >= 0)) {
                                                                                                            v124 := v2["a"]
                                                                                                            if false {
                                                                                                                v124 = v9
                                                                                                            } else {
                                                                                                                if (!true) {
                                                                                                                    var v125 int32 = 2147483647
                                                                                                                    _ = v125
                                                                                                                } else {
                                                                                                                    var v126 float64 = float64(-2.25)
                                                                                                                    switch "" {
                                                                                                                    case "foo bar":
                                                                                                                        v5 += v3
                                                                                                                        var v127 uint8 = 1
                                                                                                                        var v128 uint = 255
                                                                                                                        _ = v127
                                                                                                                        _ = v128
                                                                                                                    case "\n\t":
                                                                                                                        if len(v4) > 0 {
                                                                                                                            v4[0] = ""
                                                                                                                        }
                                                                                                                        v129 := GenS1{F0: -9223372036854775808, F1: 4294967295, F2: -132}
                                                                                                                        _ = v129
                                                                                                                    case "":
                                                                                                                        var v130 uint64 = 18446744073709551615
                                                                                                                        _ = v130
                                                                                                                    default:
                                                                                                                        v131 := GenS8{F0: true}
                                                                                                                        _ = v131
                                                                                                                    }
                                                                                                                    _ = len(v2)
                                                                                                                    _ = v126
                                                                                                                }
                                                                                                                v5 = v5
                                                                                                                if true {
                                                                                                                    if len(v1) >= 2 {
                                                                                                                        v1 = v1[1:len(v1)]
                                                                                                                    }
                                                                                                                    fmt.Println(fi11, v3, v124)
                                                                                                                    var v132 int8 = (v3 ^ v5)
                                                                                                                    _ = v132
                                                                                                                }
                                                                                                                _ = len(v6)
                                                                                                                if len(v4) > 0 {
                                                                                                                    v4[0] = "foo bar"
                                                                                                                }
                                                                                                            }
                                                                                                            var v133 uint = 42
                                                                                                            _ = v124
                                                                                                            _ = v133
                                                                                                        }
                                                                                                    default:
                                                                                                        var v134 int16 = -1
                                                                                                        _ = v134
                                                                                                    }
                                                                                                    v3 += (v3 - (v5 ^ (v3 * int8(v9))))
                                                                                                    _ = v98
                                                                                                }
                                                                                            default:
                                                                                                v135 := GenS9{F0: v5}
                                                                                                _ = v135
                                                                                            }
                                                                                            fmt.Println(v46, v3, v3)
                                                                                            var v136 float64 = 0.0
                                                                                            _ = v136
                                                                                        } else {
                                                                                            var v137 int8 = v3
                                                                                            var v138 uint = 2
                                                                                            var v139 bool = true
                                                                                            _ = v137
                                                                                            _ = v138
                                                                                            _ = v139
                                                                                        }
                                                                                    default:
                                                                                        v140 := GenS10{F0: (-(v3)), F1: 321, F2: 42}
                                                                                        _ = v140
                                                                                    }
                                                                                    var v141 int64 = 9223372036854775807
                                                                                    _ = len(v6)
                                                                                    if (!(!(!(!false)))) {
                                                                                        v46 = 255
                                                                                    } else {
                                                                                        v141 = v141
                                                                                    }
                                                                                    _ = v141
                                                                                } else {
                                                                                    v2["\n\t"] = v9
                                                                                }
                                                                            } else {
                                                                                switch (fi7 + fi11) {
                                                                                case -706:
                                                                                    fmt.Println(fi11)
                                                                                default:
                                                                                    var v142 int16 = 32767
                                                                                    _ = v142
                                                                                }
                                                                                var v143 bool = false
                                                                                var v144 string = ("foo bar" + (("\n\t" + "") + ""))
                                                                                var v145 string = (v144 + ("\n\t" + "a"))
                                                                                v146 := GenS2{F0: v145}
                                                                                _ = v143
                                                                                _ = v144
                                                                                _ = v145
                                                                                _ = v146
                                                                            }
                                                                            v9 = (v9 % ((v9) | 1))
                                                                            if (127 == v46) {
                                                                                v147 := v2["hello"]
                                                                                fmt.Println(v9, fi7, v5)
                                                                                var v148 int8 = 84
                                                                                _ = v147
                                                                                _ = v148
                                                                            }
                                                                            switch 128 {
                                                                            case 127:
                                                                                if false {
                                                                                    if false {
                                                                                        fmt.Println(fi7, v46, v5)
                                                                                    } else {
                                                                                        if true {
                                                                                            v5 = v5
                                                                                            if ((v3 / ((100) | 1)) > (v3 - 127)) {
                                                                                                v1 = append(v1, v9)
                                                                                                v149 := GenS11{F0: 2147483647, F1: (^(v9))}
                                                                                                _ = v149
                                                                                            }
                                                                                        } else {
                                                                                            v5 += (v3 ^ v5)
                                                                                            v5 -= 127
                                                                                            v46 = v46
                                                                                        }
                                                                                    }
                                                                                    switch -1 {
                                                                                    case -61:
                                                                                        fmt.Println(v3)
                                                                                    default:
                                                                                        v150 := v2[""]
                                                                                        var v151 string = ""
                                                                                        _ = v150
                                                                                        _ = v151
                                                                                    }
                                                                                    switch ("\n\t" + "a") {
                                                                                    case "\n\t":
                                                                                        v2["hello"] = v9
                                                                                    case "hello":
                                                                                        var v152 uint8 = (v46 << ((255) & 63))
                                                                                        _ = v152
                                                                                    case "foo bar":
                                                                                        if (9223372036854775807 <= 47) {
                                                                                            _ = len(v2)
                                                                                            var v153 uint8 = 127
                                                                                            var v154 uint16 = 65535
                                                                                            v155 := GenS5{F0: 2147483647, F1: 2147483647, F2: fi7}
                                                                                            var v156 int32 = -2147483648
                                                                                            _ = v153
                                                                                            _ = v154
                                                                                            _ = v155
                                                                                            _ = v156
                                                                                        } else {
                                                                                            if false {
                                                                                                var v157 int = (fi11 / ((42) | 1))
                                                                                                _ = v157
                                                                                            }
                                                                                            v1 = append(v1, (^(v9)))
                                                                                        }
                                                                                        var v158 uint8 = v46
                                                                                        _ = v158
                                                                                    default:
                                                                                        fmt.Println(v9)
                                                                                    }
                                                                                }
                                                                            case 0:
                                                                                var v159 uint8 = 2
                                                                                _ = v159
                                                                            case 174:
                                                                                var v160 uint8 = uint8(v9)
                                                                                _ = v160
                                                                            }
                                                                            v5 -= (v5 - (v5 & 34))
                                                                        }
                                                                    }
                                                                }
                                                                if false {
                                                                    var v161 int32 = -1
                                                                    v162 := GenS12{F0: (v9 & (v9 - (v9 >> ((42) & 63))))}
                                                                    if (!true) {
                                                                        v46 = v46
                                                                        var v163 uint32 = 2147483647
                                                                        var v164 uint64 = 911
                                                                        _ = v163
                                                                        _ = v164
                                                                    } else {
                                                                        var v165 int16 = 1
                                                                        v166 := GenS8{F0: ("hello" < (("foo bar" + "a") + (("a" + "a") + "\n\t")))}
                                                                        var v167 uint8 = 127
                                                                        var v168 float32 = 1.5
                                                                        _ = v165
                                                                        _ = v166
                                                                        _ = v167
                                                                        _ = v168
                                                                    }
                                                                    _ = v161
                                                                    _ = v162
                                                                } else {
                                                                    var v169 int16 = -32768
                                                                    if false {
                                                                        fmt.Println(v169)
                                                                        fmt.Println(fi11)
                                                                        v9 = 65535
                                                                    } else {
                                                                        v46 -= (v46 * (v46 >> ((541) & 63)))
                                                                        var v170 uint64 = 42
                                                                        var v171 uint8 = (v46 << ((244) & 63))
                                                                        var v172 int64 = 0
                                                                        v2[""] = 342
                                                                        _ = v170
                                                                        _ = v171
                                                                        _ = v172
                                                                    }
                                                                    var v173 float64 = 1.5
                                                                    _ = v169
                                                                    _ = v173
                                                                }
                                                                var v174 uint16 = 1
                                                                fmt.Println(fi11)
                                                                fmt.Println(fi7, v174, v46)
                                                                v175 := GenS13{F0: 1, F1: -149}
                                                                _ = v174
                                                                _ = v175
                                                            } else {
                                                                fmt.Println(v3, v5)
                                                                switch fi7 {
                                                                case 919:
                                                                    v1 = append(v1, 2147483647)
                                                                }
                                                                fmt.Println(v5, fi7)
                                                            }
                                                        }
                                                    case "a":
                                                        fmt.Println(v5)
                                                    case "hello":
                                                        fmt.Println(v5, fi7, v5)
                                                    }
                                                case 32767:
                                                    v176 := GenS12{F0: v9}
                                                    _ = v176
                                                default:
                                                    fmt.Println(fi11)
                                                }
                                            }
                                        case "a":
                                            switch fi7 {
                                            case 2147483647:
                                                v5 = 1
                                            case -990:
                                                var v177 int32 = 1
                                                _ = v177
                                            case 9223372036854775807:
                                                var v178 bool = (4294967295 >= uint(v46))
                                                _ = v178
                                            }
                                            v179 := GenS14{F0: 18446744073709551615, F1: 0, F2: 32767}
                                            _ = v179
                                        case "hello":
                                            var v180 float32 = (0.0 - 0.1)
                                            _ = v180
                                        }
                                        var v181 bool = false
                                        _ = v181
                                        _ = v46
                                    }
                                case -470:
                                    v3 = 42
                                    _ = len(v2)
                                }
                                if (1e10 >= ((0.1 + 0.0) * float64(0.0))) {
                                    v182 := GenS13{F0: 240, F1: -803}
                                    _ = v182
                                } else {
                                    v3 -= v5
                                }
                                var v183 int = (fi7 - fi11)
                                _ = v183
                            }
                            v3 = (v5 << ((1) & 63))
                            switch "a" {
                            case "":
                                v184 := GenS10{F0: (v5 + v5), F1: 381, F2: -1}
                                _ = v184
                            default:
                                switch (fi11 | (fi7 + 2147483647)) {
                                case 42:
                                    fmt.Println(v9, fi11, v9)
                                case -2147483648:
                                    switch ("\n\t" + ("foo bar" + "foo bar")) {
                                    case "foo bar":
                                        if true {
                                            if (!true) {
                                                v185 := GenS4{F0: 750, F1: v9}
                                                if false {
                                                    v186 := GenS13{F0: 42, F1: 184}
                                                    if (2 != 0) {
                                                        var v187 uint32 = (v9 << ((255) & 63))
                                                        _ = v187
                                                    } else {
                                                        v188 := GenS10{F0: -84, F1: 2, F2: int16(v3)}
                                                        if true {
                                                            var v189 int = (fi11 ^ -9223372036854775808)
                                                            var v190 uint8 = 0
                                                            v191 := GenS9{F0: v5}
                                                            _ = v189
                                                            _ = v190
                                                            _ = v191
                                                        } else {
                                                            var v192 uint64 = 42
                                                            var v193 uint32 = (^(v9))
                                                            fmt.Println(v9, v5)
                                                            _ = v192
                                                            _ = v193
                                                        }
                                                        var v194 int = (fi7 * fi11)
                                                        v194 -= fi7
                                                        _ = v188
                                                        _ = v194
                                                    }
                                                    var v195 float64 = 1e10
                                                    _ = v186
                                                    _ = v195
                                                } else {
                                                    v196 := v2["\n\t"]
                                                    _ = len(v4)
                                                    if (42 > 42) {
                                                        delete(v2, ("a" + "a"))
                                                        switch "hello" {
                                                        case "hello":
                                                            v9 = (v9 & v196)
                                                        case "\n\t":
                                                            if len(v4) >= 2 {
                                                                v4 = v4[1:len(v4)]
                                                            }
                                                        }
                                                    }
                                                    var v197 int64 = -1
                                                    _ = v196
                                                    _ = v197
                                                }
                                                switch ((("hello" + "") + "a") + "\n\t") {
                                                case "hello":
                                                    switch "foo bar" {
                                                    case "":
                                                        v198 := GenS15{F0: 255, F1: 378, F2: -888}
                                                        _ = v198
                                                    default:
                                                        var v199 uint = 255
                                                        _ = v199
                                                    }
                                                case "\n\t":
                                                    v200 := GenS16{F0: 962, F1: 0, F2: (-583 != 363)}
                                                    _ = v200
                                                case "":
                                                    v3 = v3
                                                }
                                                _ = v185
                                            }
                                            v5 = (v3 - int8(v9))
                                        } else {
                                            if ((0 == -2147483648) || (-449 == 2147483647)) {
                                                v201 := GenS7{F0: 4294967295, F1: (!false), F2: (v9 >> ((993) & 63))}
                                                var v202 float64 = ((1e10 * 3.14159) * 1.5)
                                                _ = v201
                                                _ = v202
                                            }
                                            v5 += v5
                                            v203 := GenS17{F0: 60, F1: -1, F2: -2147483648}
                                            _ = v203
                                        }
                                    }
                                default:
                                    switch 32768 {
                                    case -9223372036854775808:
                                        if (((468 < 2) || false) && ((v3 & v5) != (-(v3)))) {
                                            v5 += (v3 + v5)
                                            switch "a" {
                                            case "foo bar":
                                                v3 = (v5 ^ (v3 << ((18446744073709551615) & 63)))
                                            case "":
                                                v3 = 42
                                            case "hello":
                                                switch "foo bar" {
                                                case "foo bar":
                                                    v204 := GenS18{F0: 42, F1: (fi7 + fi11), F2: 18446744073709551615}
                                                    _ = v204
                                                case "":
                                                    v205 := GenS19{F0: 200}
                                                    _ = v205
                                                case "hello":
                                                    if false {
                                                        v206 := GenS20{F0: (!(1 != -859))}
                                                        _ = v206
                                                    }
                                                }
                                            default:
                                                switch ("" + "hello") {
                                                case "foo bar":
                                                    if ((v3 != 61) && true) {
                                                        v3 += 0
                                                        var v207 int32 = -1
                                                        _ = v207
                                                    }
                                                case "":
                                                    v2["hello"] = (^(v9))
                                                    fmt.Println(v9, fi7, v5)
                                                    var v208 uint8 = 196
                                                    _ = v208
                                                default:
                                                    v2[("a" + "a")] = (^(v9))
                                                    fmt.Println(v9, v9, fi7)
                                                }
                                            }
                                            var v209 float64 = (1.5 + 0.0)
                                            _ = v209
                                        } else {
                                            switch "\n\t" {
                                            case "\n\t":
                                                v9 += 4294967295
                                            default:
                                                if (true && true) {
                                                    if len(v6) > 0 {
                                                        v6[0] = 42
                                                    }
                                                    var v210 float32 = -2.25
                                                    _ = v210
                                                }
                                                var v211 uint64 = 4294967295
                                                _ = v211
                                            }
                                            var v212 int8 = v5
                                            _ = v212
                                        }
                                    case 32767:
                                        if (!((false || false) || true)) {
                                            if len(v1) > 0 {
                                                v1[0] = (v9 << ((834) & 63))
                                            }
                                            v213 := GenS21{F0: 4294967295, F1: 18446744073709551615}
                                            switch "foo bar" {
                                            case "a":
                                                v4 = append(v4, "hello")
                                            case "hello":
                                                fmt.Println(fi11, v5, fi7)
                                            default:
                                                var v214 int = 221
                                                _ = v214
                                            }
                                            fmt.Println(v3, fi11)
                                            _ = v213
                                        }
                                    case 9223372036854775807:
                                        var v215 uint16 = 0
                                        _ = v215
                                    }
                                }
                            }
                        }
                        var v216 int16 = 1
                        _ = v216
                    }
                    v5 = -1
                    _ = fi11
                }
                _ = v9
            } else {
                var v217 uint32 = 315
                _ = len(v2)
                _ = v217
            }
            var v218 uint = 2
            if ((true || false) && (!(32767 >= 32767))) {
                var v219 int16 = 42
                _ = v219
            }
            _ = v218
        }
        _ = fi7
        _ = v8
    }
    for fi220 := 0; fi220 < 18; fi220++ {
        _ = len(v2)
        v3 += v3
        var v221 int32 = 0
        _ = fi220
        _ = v221
    }
    switch 9223372036854775807 {
    case 2147483647:
        var v222 string = ("hello" + ("hello" + "\n\t"))
        _ = v222
    case 317:
        if (!(1 >= 255)) {
            fmt.Println(v3, v5)
            v3 += 0
            switch 127 {
            case 9223372036854775807:
                v223 := GenS22{F0: 32768, F1: 1000, F2: ("hello" + "foo bar")}
                _ = v223
            case 775:
                v3 = (-(v5))
            }
        }
    }
    _ = len(v2)
    var v224 uint64 = 0
    if true {
        var v225 uint16 = 0
        fmt.Println(v225)
        _ = v225
    } else {
        fmt.Println(v5, v5, v5)
    }
    var v226 uint32 = 1
    var v227 int = -1
    var v228 uint = 221
    if (128 != 127) {
        var v229 float64 = (float64(float32((-2.25 / (float64(1.5) / 3.14159)))) / 1.0)
        v226 -= 42
        if true {
            var v230 bool = ((false && true) || true)
            if (v230 || (456 >= v228)) {
                v229 = v229
                if ((v230 && false) && (false || (v227 <= 4294967295))) {
                    switch (("\n\t" + "a") + "hello") {
                    case "a":
                        var v231 int = 1
                        _ = v231
                    case "hello":
                        v224 += 180
                    case "":
                        var v232 int8 = (-(v3))
                        _ = v232
                    }
                    switch 32768 {
                    case 9223372036854775807:
                        var v233 uint8 = 0
                        _ = v233
                    }
                    var v234 float32 = 1.5
                    _ = v234
                }
            } else {
                switch "" {
                case "\n\t":
                    v224 = (v224 ^ v224)
                case "a":
                    fmt.Println(v229)
                case "foo bar":
                    v235 := GenS1{F0: 1, F1: (^(v224)), F2: int16(v3)}
                    _ = v235
                default:
                    if v230 {
                        for fi236 := 0; fi236 < 15; fi236++ {
                            for fi237 := 0; fi237 < 14; fi237++ {
                                if (0 == -9223372036854775808) {
                                    switch (("foo bar" + "hello") + "\n\t") {
                                    case "foo bar":
                                        if len(v6) >= 2 {
                                            v6 = v6[1:len(v6)]
                                        }
                                    }
                                } else {
                                    fmt.Println(fi237, v5, v3)
                                    v238 := GenS23{F0: v227}
                                    v230 = (!true)
                                    _ = v238
                                }
                                v239 := GenS4{F0: 720, F1: uint32(v228)}
                                _ = fi237
                                _ = v239
                            }
                            for fi240 := 0; fi240 < 15; fi240++ {
                                v3 += v5
                                fmt.Println(fi240)
                                _ = fi240
                            }
                            _ = fi236
                        }
                        fmt.Println(v229, v226, v5)
                    }
                }
                if false {
                    fmt.Println(v228, v229)
                }
            }
            _ = v230
        } else {
            v228 = (v228 | (v228 + v228))
            var v241 uint = 0
            _ = v241
        }
        _ = v229
    } else {
        for fi242 := 0; fi242 < 9; fi242++ {
            if (3.14159 > 1.5) {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                var v243 uint32 = v226
                var v244 uint8 = 42
                _ = v243
                _ = v244
            } else {
                var v245 int64 = -9223372036854775808
                v2[(("a" + "\n\t") + ("\n\t" + ("a" + "")))] = v226
                _ = v245
            }
            _ = fi242
        }
        if true {
            if len(v1) > 0 {
                v1[0] = v226
            }
            for fi246 := 0; fi246 < 3; fi246++ {
                var v247 uint = v228
                var v248 int32 = 274
                switch ("foo bar" + ("foo bar" + "a")) {
                case "foo bar":
                    if (v224 == (v224 % ((v224) | 1))) {
                        for fi249 := 0; fi249 < 9; fi249++ {
                            v3 += -74
                            _ = fi249
                        }
                        delete(v2, (("" + ("\n\t" + "")) + "hello"))
                    }
                default:
                    if (!(520 != v226)) {
                        v250 := GenS12{F0: 54}
                        if true {
                            fmt.Println(fi246, v3)
                        }
                        _ = v250
                    }
                }
                _ = fi246
                _ = v247
                _ = v248
            }
            fmt.Println(v226, v224)
        }
        var v251 int = (v227 ^ 468)
        _ = v251
    }
    if false {
        v252 := GenS8{F0: (true || (false && false))}
        if true {
            if true {
                var v253 uint64 = (v224 % ((v224) | 1))
                _ = v253
            } else {
                v228 -= v228
                v254 := GenS22{F0: 128, F1: 448, F2: ("foo bar" + "a")}
                _ = v254
            }
            v255 := GenS24{F0: 212}
            _ = v255
        }
        fmt.Println(v227, v228)
        _ = v252
    } else {
        var v256 bool = false
        v257 := GenS25{F0: v226, F1: (^(v226))}
        v258 := GenS2{F0: "hello"}
        _ = v256
        _ = v257
        _ = v258
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v224:", v224)
    fmt.Println("v226:", v226)
    fmt.Println("v227:", v227)
    fmt.Println("v228:", v228)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v224
    _ = v226
    _ = v227
    _ = v228
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
