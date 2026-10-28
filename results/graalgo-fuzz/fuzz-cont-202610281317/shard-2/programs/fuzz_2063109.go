package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int32; F1 uint16; F2 uint64 }
type GenS5 struct { F0 uint64; F1 string }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 string; F1 string; F2 int8 }
type GenS9 struct { F0 int32 }
type GenS10 struct { F0 bool }
type GenS11 struct { F0 uint; F1 int16; F2 int16 }
type GenS12 struct { F0 int8; F1 int }
type GenS13 struct { F0 uint8; F1 int16; F2 int8 }
type GenS14 struct { F0 uint16 }
type GenS15 struct { F0 int32 }
type GenS16 struct { F0 int }
type GenS17 struct { F0 int; F1 int8; F2 int }
type GenS18 struct { F0 int32; F1 uint8; F2 uint8 }
type GenS19 struct { F0 uint; F1 int8 }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[int]int)
    var v3 uint = 91
    var v4 float64 = 1e10
    v5 := []int{127, -808, 421}
    v6 := GenS1{F0: 556, F1: 9223372036854775807}
    switch 4294967295 {
    case -245:
        var v7 uint8 = 128
        _ = v7
    default:
        var v8 bool = (!("foo bar" > "\n\t"))
        _ = v8
    }
    switch -1 {
    case 42:
        for fi9 := 0; fi9 < 20; fi9++ {
            if false {
                for fi10 := 0; fi10 < 12; fi10++ {
                    if (!true) {
                        var v11 int8 = 1
                        v11 += -128
                        var v12 string = ("foo bar" + "hello")
                        _ = v11
                        _ = v12
                    }
                    fmt.Println(fi9, v3, fi9)
                    switch "\n\t" {
                    case "":
                        if (!(!((0 < -32768) || false))) {
                            v3 = (^(v3))
                            switch "a" {
                            case "hello":
                                if true {
                                    v3 = v3
                                    v4 -= v4
                                    if true {
                                        v3 = (v3 - (v3 << ((18446744073709551615) & 63)))
                                        if false {
                                            if false {
                                                if false {
                                                    fmt.Println(fi9)
                                                    v13 := GenS1{F0: 42, F1: (fi10 / ((127) | 1))}
                                                    _ = v13
                                                }
                                                var v14 int32 = -917
                                                v3 = v3
                                                _ = v14
                                            } else {
                                                fmt.Println(v3)
                                                if (!true) {
                                                    var v15 int64 = 0
                                                    fmt.Println(fi10)
                                                    switch (("foo bar" + "foo bar") + "foo bar") {
                                                    case "foo bar":
                                                        v16 := GenS2{F0: 4294967295}
                                                        _ = v16
                                                    case "a":
                                                        if ((!(false && (!true))) && (275 < 2)) {
                                                            var v17 float64 = (0.0 + (v4 / -2.25))
                                                            v1[866] = 4294967295
                                                            _ = v17
                                                        } else {
                                                            switch fi9 {
                                                            case 697:
                                                                fmt.Println(fi10, v4)
                                                            case 128:
                                                                v4 = float64(3.14159)
                                                            }
                                                            if len(v5) > 0 {
                                                                v5[0] = 4294967295
                                                            }
                                                            v3 += (v3 | (v3 & v3))
                                                        }
                                                        switch (("a" + ("foo bar" + "foo bar")) + ((("foo bar" + "foo bar") + (("" + "\n\t") + "a")) + "hello")) {
                                                        case "":
                                                            var v18 string = "foo bar"
                                                            _ = v18
                                                        case "hello":
                                                            var v19 uint8 = 127
                                                            _ = v19
                                                        case "foo bar":
                                                            var v20 uint8 = 127
                                                            _ = v20
                                                        }
                                                        v15 += v15
                                                        v21 := v1[2147483647]
                                                        _ = v21
                                                    case "hello":
                                                        fmt.Println(fi10, fi10, v15)
                                                    }
                                                    var v22 bool = true
                                                    if len(v5) >= 2 {
                                                        v5 = v5[1:len(v5)]
                                                    }
                                                    _ = v15
                                                    _ = v22
                                                }
                                            }
                                        }
                                        switch "" {
                                        case "hello":
                                            if (-43 >= 0) {
                                                var v23 float64 = float64(-2.25)
                                                _ = v23
                                            }
                                        case "foo bar":
                                            if true {
                                                v4 -= v4
                                                switch ("a" + "\n\t") {
                                                case "a":
                                                    if false {
                                                        if (466 > 4294967295) {
                                                            var v24 uint8 = 0
                                                            if (true || (709 != 0)) {
                                                                if true {
                                                                    if false {
                                                                        if true {
                                                                            var v25 uint16 = 32767
                                                                            var v26 float32 = 1e10
                                                                            _ = v25
                                                                            _ = v26
                                                                        } else {
                                                                            var v27 string = ""
                                                                            _ = v27
                                                                        }
                                                                        var v28 uint16 = 42
                                                                        if true {
                                                                            v1[317] = 65535
                                                                            var v29 uint64 = 609
                                                                            _ = v29
                                                                        } else {
                                                                            switch "foo bar" {
                                                                            case "a":
                                                                                v24 = (v24 / (((v24 - (v24 - (^(v24))))) | 1))
                                                                                if (true || (!(v4 == 0.0))) {
                                                                                    v28 = 830
                                                                                    var v30 bool = (((v28 != (v28 & 65535)) && (("hello" + "\n\t") <= "hello")) && (!(!(-9223372036854775808 < int64(v28)))))
                                                                                    _ = v30
                                                                                } else {
                                                                                    v24 = 119
                                                                                }
                                                                                _ = len(v5)
                                                                            case "hello":
                                                                                if true {
                                                                                    var v31 uint = v3
                                                                                    switch "" {
                                                                                    case "foo bar":
                                                                                        var v32 float64 = ((v4 / float64(3.14159)) - 0.1)
                                                                                        fmt.Println(fi9, v24, v3)
                                                                                        if len(v5) >= 2 {
                                                                                            v5 = v5[1:len(v5)]
                                                                                        }
                                                                                        v3 = 65535
                                                                                        _ = v32
                                                                                    case "\n\t":
                                                                                        if ((42 == -2147483648) && (9223372036854775807 != -9223372036854775808)) {
                                                                                            if len(v5) > 0 {
                                                                                                v5[0] = -904
                                                                                            }
                                                                                            if (-444 < (fi10 * fi10)) {
                                                                                                v24 += 42
                                                                                                v33 := GenS3{F0: 0}
                                                                                                _ = v33
                                                                                            }
                                                                                        } else {
                                                                                            switch "a" {
                                                                                            case "\n\t":
                                                                                                v3 = 2
                                                                                            case "foo bar":
                                                                                                fmt.Println(v28, v3, v31)
                                                                                            case "":
                                                                                                v34 := GenS3{F0: 1}
                                                                                                _ = v34
                                                                                            }
                                                                                        }
                                                                                    case "hello":
                                                                                        switch 612 {
                                                                                        case 9223372036854775807:
                                                                                            if false {
                                                                                                v35 := GenS3{F0: 17}
                                                                                                var v36 bool = (fi10 <= (fi9 >> ((v31) & 63)))
                                                                                                if v36 {
                                                                                                    var v37 uint32 = 2
                                                                                                    var v38 uint16 = 2
                                                                                                    if (!(!(((-2.25 + float32(v4)) - -2.25) < ((float32(v4) - (1e10 * ((0.1 / 3.14159) / -2.25))) / 1.0)))) {
                                                                                                        if v36 {
                                                                                                            if (!v36) {
                                                                                                                v24 = v24
                                                                                                                v28 = (v28 | v38)
                                                                                                                var v39 float64 = v4
                                                                                                                _ = v39
                                                                                                            }
                                                                                                        } else {
                                                                                                            v31 = v31
                                                                                                            v38 = uint16(v24)
                                                                                                            var v40 int = -2147483648
                                                                                                            _ = v40
                                                                                                        }
                                                                                                        v37 = (v37 << ((v31) & 63))
                                                                                                    } else {
                                                                                                        v41 := GenS4{F0: -1, F1: (^(v28)), F2: uint64(v3)}
                                                                                                        v42 := GenS3{F0: uint64(v37)}
                                                                                                        _ = v41
                                                                                                        _ = v42
                                                                                                    }
                                                                                                    if (!v36) {
                                                                                                        if v36 {
                                                                                                            switch "hello" {
                                                                                                            case "\n\t":
                                                                                                                v31 = 4294967295
                                                                                                            case "a":
                                                                                                                if len(v5) > 0 {
                                                                                                                    v5[0] = 9223372036854775807
                                                                                                                }
                                                                                                            default:
                                                                                                                if false {
                                                                                                                    v43 := GenS5{F0: 2, F1: "\n\t"}
                                                                                                                    if v36 {
                                                                                                                        fmt.Println(v36, v3, v28)
                                                                                                                        var v44 int64 = 0
                                                                                                                        v37 = (v37 | 42)
                                                                                                                        _ = v44
                                                                                                                    } else {
                                                                                                                        var v45 float32 = (3.14159 * (1.5 - 0.1))
                                                                                                                        _ = v45
                                                                                                                    }
                                                                                                                    v36 = v36
                                                                                                                    _ = v43
                                                                                                                }
                                                                                                            }
                                                                                                        } else {
                                                                                                            v46 := GenS4{F0: int32(v37), F1: (v38 & 964), F2: 442}
                                                                                                            _ = v46
                                                                                                        }
                                                                                                        if (18446744073709551615 == 4294967295) {
                                                                                                            v3 = (v3 & (v31 % ((65535) | 1)))
                                                                                                            fmt.Println(v3)
                                                                                                            v36 = true
                                                                                                        } else {
                                                                                                            var v47 bool = ((!false) && false)
                                                                                                            _ = v47
                                                                                                        }
                                                                                                        v28 += v28
                                                                                                        if (127 < 127) {
                                                                                                            var v48 int64 = -284
                                                                                                            _ = v48
                                                                                                        }
                                                                                                    }
                                                                                                    var v49 uint16 = (^(v38))
                                                                                                    if ((^(v31)) > 307) {
                                                                                                        var v50 uint64 = 2
                                                                                                        switch (fi9 - fi9) {
                                                                                                        case 175:
                                                                                                            if ((false || (1.5 < (v4 / 1e10))) || ((v24 % ((196) | 1)) <= 1)) {
                                                                                                                delete(v1, v37)
                                                                                                                var v51 int = 42
                                                                                                                v52 := GenS6{F0: v37}
                                                                                                                _ = len(v2)
                                                                                                                _ = v51
                                                                                                                _ = v52
                                                                                                            }
                                                                                                        }
                                                                                                        switch (fi10 + fi10) {
                                                                                                        case 1:
                                                                                                            var v53 string = ("foo bar" + (("a" + "") + "a"))
                                                                                                            _ = v53
                                                                                                        }
                                                                                                        _ = v50
                                                                                                    }
                                                                                                    fmt.Println(v3, fi10)
                                                                                                    _ = v37
                                                                                                    _ = v38
                                                                                                    _ = v49
                                                                                                } else {
                                                                                                    var v54 int8 = 53
                                                                                                    v28 += v28
                                                                                                    v55 := GenS7{F0: v24}
                                                                                                    _ = v54
                                                                                                    _ = v55
                                                                                                }
                                                                                                _ = v35
                                                                                                _ = v36
                                                                                            }
                                                                                        case -1:
                                                                                            var v56 int64 = 42
                                                                                            _ = v56
                                                                                        case 128:
                                                                                            v4 += -2.25
                                                                                            switch "" {
                                                                                            case "foo bar":
                                                                                                v57 := v2[fi9]
                                                                                                _ = v57
                                                                                            case "":
                                                                                                v3 -= v31
                                                                                                var v58 uint32 = 4294967295
                                                                                                var v59 int8 = 127
                                                                                                _ = v58
                                                                                                _ = v59
                                                                                            default:
                                                                                                v4 -= float64(1.5)
                                                                                            }
                                                                                            switch fi9 {
                                                                                            case 943:
                                                                                                v60 := GenS1{F0: (v28 - v28), F1: 0}
                                                                                                _ = v60
                                                                                            }
                                                                                        default:
                                                                                            var v61 uint16 = v28
                                                                                            _ = v61
                                                                                        }
                                                                                    default:
                                                                                        if (true && (true && true)) {
                                                                                            if ((1 > 947) && true) {
                                                                                                v31 = 255
                                                                                                var v62 float64 = 1e10
                                                                                                if len(v5) >= 2 {
                                                                                                    v5 = v5[1:len(v5)]
                                                                                                }
                                                                                                _ = v62
                                                                                            }
                                                                                            if true {
                                                                                                switch (fi9 | fi9) {
                                                                                                case 9223372036854775807:
                                                                                                    v1[412] = 4294967295
                                                                                                default:
                                                                                                    delete(v1, 4294967295)
                                                                                                }
                                                                                            } else {
                                                                                                switch -2147483648 {
                                                                                                case -851:
                                                                                                    v24 += (v24 | (v24 % ((v24) | 1)))
                                                                                                default:
                                                                                                    v63 := GenS7{F0: v24}
                                                                                                    _ = v63
                                                                                                }
                                                                                                var v64 float64 = (v4 + float64(0.1))
                                                                                                if false {
                                                                                                    var v65 float64 = (v4 - -2.25)
                                                                                                    _ = v65
                                                                                                } else {
                                                                                                    v64 += (0.0 + float64(float32(v4)))
                                                                                                    v66 := GenS1{F0: 2, F1: 418}
                                                                                                    if (!true) {
                                                                                                        v28 = (v28 / ((500) | 1))
                                                                                                        if (131 <= v24) {
                                                                                                            v24 += 255
                                                                                                        }
                                                                                                    } else {
                                                                                                        var v67 float64 = v4
                                                                                                        v68 := GenS2{F0: 0}
                                                                                                        _ = v67
                                                                                                        _ = v68
                                                                                                    }
                                                                                                    _ = v66
                                                                                                }
                                                                                                var v69 int16 = 1000
                                                                                                v4 += v64
                                                                                                v64 += ((1e10 * v4) / (v64 * v64))
                                                                                                fmt.Println(v31, v24, v31)
                                                                                                var v70 uint64 = 2
                                                                                                var v71 uint64 = v70
                                                                                                _ = v64
                                                                                                _ = v69
                                                                                                _ = v70
                                                                                                _ = v71
                                                                                            }
                                                                                            if false {
                                                                                                v72 := GenS6{F0: 10}
                                                                                                var v73 string = ""
                                                                                                v74 := GenS5{F0: 655, F1: v73}
                                                                                                _ = v72
                                                                                                _ = v73
                                                                                                _ = v74
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                    fmt.Println(fi9, v31)
                                                                                    _ = v31
                                                                                } else {
                                                                                    var v75 string = (((("foo bar" + "hello") + "foo bar") + ("hello" + "")) + ("" + "\n\t"))
                                                                                    var v76 uint32 = 42
                                                                                    _ = v75
                                                                                    _ = v76
                                                                                }
                                                                            case "foo bar":
                                                                                v77 := v2[-311]
                                                                                _ = v77
                                                                            default:
                                                                                switch "a" {
                                                                                case "a":
                                                                                    v78 := GenS7{F0: (v24 | v24)}
                                                                                    _ = v78
                                                                                case "foo bar":
                                                                                    if false {
                                                                                        v79 := v2[(fi9 << ((v3) & 63))]
                                                                                        _ = v79
                                                                                    }
                                                                                case "hello":
                                                                                    if true {
                                                                                        fmt.Println(fi10)
                                                                                        var v80 int8 = 0
                                                                                        v81 := GenS8{F0: "foo bar", F1: "\n\t", F2: v80}
                                                                                        _ = v80
                                                                                        _ = v81
                                                                                    }
                                                                                default:
                                                                                    if false {
                                                                                        if false {
                                                                                            v24 -= v24
                                                                                            if len(v5) > 0 {
                                                                                                v5[0] = (fi10 - fi9)
                                                                                            }
                                                                                            v28 = 762
                                                                                        } else {
                                                                                            v28 = (^(v28))
                                                                                            v5 = append(v5, fi9)
                                                                                            var v82 string = "foo bar"
                                                                                            _ = v82
                                                                                        }
                                                                                        var v83 int8 = 42
                                                                                        v28 = v28
                                                                                        _ = v83
                                                                                    } else {
                                                                                        v84 := GenS5{F0: 2, F1: "a"}
                                                                                        if len(v5) > 0 {
                                                                                            v5[0] = (fi9 << ((v3) & 63))
                                                                                        }
                                                                                        v85 := GenS5{F0: 4294967295, F1: ("a" + "\n\t")}
                                                                                        _ = v84
                                                                                        _ = v85
                                                                                    }
                                                                                }
                                                                            }
                                                                        }
                                                                        _ = v28
                                                                    }
                                                                    var v86 float64 = v4
                                                                    _ = v86
                                                                } else {
                                                                    if false {
                                                                        if (v4 >= (3.14159 + 1e10)) {
                                                                            v4 -= (v4 + v4)
                                                                            v87 := GenS4{F0: -102, F1: 65535, F2: 1}
                                                                            var v88 float32 = 1.5
                                                                            _ = v87
                                                                            _ = v88
                                                                        } else {
                                                                            v4 -= (v4 - (v4 + v4))
                                                                        }
                                                                        v4 = v4
                                                                        if (v24 > (v24 - v24)) {
                                                                            v4 += v4
                                                                            if true {
                                                                                var v89 bool = true
                                                                                var v90 uint8 = (v24 / ((128) | 1))
                                                                                switch (-(fi10)) {
                                                                                case -907:
                                                                                    v90 = (v24 % ((v90) | 1))
                                                                                case 32767:
                                                                                    v90 = 127
                                                                                case 726:
                                                                                    v24 += v90
                                                                                default:
                                                                                    switch ("foo bar" + (("hello" + "\n\t") + "foo bar")) {
                                                                                    case "\n\t":
                                                                                        switch fi9 {
                                                                                        case 32768:
                                                                                            if v89 {
                                                                                                v91 := GenS6{F0: 65535}
                                                                                                var v92 float64 = float64(-2.25)
                                                                                                _ = v91
                                                                                                _ = v92
                                                                                            } else {
                                                                                                var v93 int16 = 1
                                                                                                v90 -= v90
                                                                                                var v94 int32 = 158
                                                                                                _ = v93
                                                                                                _ = v94
                                                                                            }
                                                                                        }
                                                                                    case "foo bar":
                                                                                        v24 = (v24 >> ((236) & 63))
                                                                                        v3 = v3
                                                                                    default:
                                                                                        var v95 int = fi9
                                                                                        _ = v95
                                                                                    }
                                                                                }
                                                                                _ = v89
                                                                                _ = v90
                                                                            } else {
                                                                                v96 := GenS2{F0: 2}
                                                                                _ = v96
                                                                            }
                                                                            v4 = float64(0.1)
                                                                        } else {
                                                                            switch "hello" {
                                                                            case "":
                                                                                v24 = v24
                                                                            case "a":
                                                                                v97 := GenS9{F0: -2147483648}
                                                                                _ = v97
                                                                            }
                                                                            switch 2147483647 {
                                                                            case -2147483648:
                                                                                v5 = append(v5, fi10)
                                                                            case -857:
                                                                                if false {
                                                                                    switch "hello" {
                                                                                    case "a":
                                                                                        if (false || ((^(v24)) > (v24 % ((1) | 1)))) {
                                                                                            var v98 uint8 = 255
                                                                                            if (((-(fi10)) <= (fi9 ^ (fi10 >> (((^(v3))) & 63)))) && ("hello" <= ("foo bar" + "hello"))) {
                                                                                                if (32767 == 0) {
                                                                                                    if (("" + "hello") <= "foo bar") {
                                                                                                        if true {
                                                                                                            fmt.Println(fi10)
                                                                                                        } else {
                                                                                                            v98 = v98
                                                                                                            switch int(v98) {
                                                                                                            case -591:
                                                                                                                fmt.Println(v3, fi9)
                                                                                                            default:
                                                                                                                var v99 uint32 = 4294967295
                                                                                                                _ = v99
                                                                                                            }
                                                                                                        }
                                                                                                        fmt.Println(v3, v3)
                                                                                                        var v100 float64 = ((v4 + (1.5 + 0.1)) / (0.0 * v4))
                                                                                                        var v101 int64 = 42
                                                                                                        _ = v100
                                                                                                        _ = v101
                                                                                                    }
                                                                                                }
                                                                                                delete(v2, (^(fi10)))
                                                                                                var v102 int32 = 0
                                                                                                _ = v102
                                                                                            } else {
                                                                                                v103 := v1[490]
                                                                                                if (1 <= (v3 << ((1) & 63))) {
                                                                                                    var v104 uint32 = 0
                                                                                                    fmt.Println(fi10, v3)
                                                                                                    switch ((("\n\t" + "") + "\n\t") + ("\n\t" + "\n\t")) {
                                                                                                    case "":
                                                                                                        v24 = (v24 << ((v3) & 63))
                                                                                                    case "a":
                                                                                                        v98 -= v98
                                                                                                    }
                                                                                                    _ = v104
                                                                                                } else {
                                                                                                    v4 = v4
                                                                                                    delete(v1, (^(v103)))
                                                                                                }
                                                                                                var v105 int16 = -1
                                                                                                var v106 uint16 = 32767
                                                                                                var v107 int8 = 100
                                                                                                _ = v103
                                                                                                _ = v105
                                                                                                _ = v106
                                                                                                _ = v107
                                                                                            }
                                                                                            _ = v98
                                                                                        }
                                                                                    case "\n\t":
                                                                                        var v108 uint16 = uint16(v24)
                                                                                        _ = v108
                                                                                    default:
                                                                                        v24 += v24
                                                                                    }
                                                                                    fmt.Println(v4)
                                                                                } else {
                                                                                    var v109 int16 = 32767
                                                                                    _ = v109
                                                                                }
                                                                            }
                                                                            if false {
                                                                                fmt.Println(v4, v4, fi10)
                                                                                switch (fi9 >> ((v3) & 63)) {
                                                                                case 32768:
                                                                                    v24 = (v24 / ((v24) | 1))
                                                                                default:
                                                                                    if true {
                                                                                        v110 := GenS10{F0: false}
                                                                                        _ = v110
                                                                                    }
                                                                                }
                                                                                v24 += 255
                                                                            } else {
                                                                                if false {
                                                                                    var v111 uint = v3
                                                                                    _ = v111
                                                                                }
                                                                                v3 = (v3 * v3)
                                                                                v24 -= v24
                                                                            }
                                                                        }
                                                                    }
                                                                }
                                                            }
                                                            var v112 string = ("\n\t" + "a")
                                                            _ = v112
                                                            _ = v24
                                                        } else {
                                                            v113 := GenS10{F0: false}
                                                            _ = v113
                                                        }
                                                        v4 = v4
                                                        v3 = v3
                                                    } else {
                                                        v3 = v3
                                                        switch (("hello" + "a") + "hello") {
                                                        case "a":
                                                            var v114 float32 = (0.0 * (((0.1 * (((-2.25 * (3.14159 + (1e10 * 1e10))) + 0.0) * (0.0 * (((0.0 * 0.1) - 0.0) + -2.25)))) / (1.5 / -2.25)) * (((1.5 - ((1.5 * (3.14159 - (0.1 * 1e10))) * -2.25)) + -2.25) / 1.5)))
                                                            _ = v114
                                                        case "\n\t":
                                                            var v115 uint8 = 61
                                                            _ = v115
                                                        default:
                                                            var v116 int64 = 0
                                                            _ = v116
                                                        }
                                                    }
                                                case "\n\t":
                                                    v3 = 4294967295
                                                case "":
                                                    v117 := GenS11{F0: 42, F1: 275, F2: -1}
                                                    _ = v117
                                                }
                                                _ = len(v1)
                                                if (v3 != (v3 % ((0) | 1))) {
                                                    if (42 <= 2) {
                                                        fmt.Println(fi9, fi9, v3)
                                                        switch "" {
                                                        case "foo bar":
                                                            if (1.5 < 0.1) {
                                                                switch fi9 {
                                                                case 32768:
                                                                    var v118 bool = true
                                                                    _ = v118
                                                                }
                                                                switch "\n\t" {
                                                                case "\n\t":
                                                                    var v119 uint8 = 128
                                                                    _ = v119
                                                                case "hello":
                                                                    v120 := GenS4{F0: -461, F1: 2, F2: 2}
                                                                    _ = v120
                                                                default:
                                                                    v3 = 570
                                                                }
                                                            }
                                                        default:
                                                            switch fi10 {
                                                            case 32767:
                                                                var v121 bool = false
                                                                _ = v121
                                                            }
                                                        }
                                                        v3 -= v3
                                                    } else {
                                                        fmt.Println(v4)
                                                        v3 = 42
                                                        v3 = (^(v3))
                                                    }
                                                    v4 = v4
                                                    if false {
                                                        var v122 uint8 = 134
                                                        if true {
                                                            var v123 int8 = -1
                                                            v4 -= (3.14159 + ((3.14159 - ((((((v4 / v4) / ((float64(1e10) - (v4 + (0.0 - (((v4 * (v4 + v4)) / 1.5) - 0.1)))) / (0.1 + v4))) / float64(1.5)) * 0.1) + (v4 - v4)) + v4)) + ((float64(-2.25) + float64((0.1 + 0.0))) - v4)))
                                                            _ = len(v5)
                                                            _ = v123
                                                        } else {
                                                            if (!(v122 >= v122)) {
                                                                v5 = append(v5, fi9)
                                                                var v124 float64 = float64(3.14159)
                                                                _ = v124
                                                            } else {
                                                                if true {
                                                                    var v125 uint64 = 18446744073709551615
                                                                    v126 := GenS5{F0: (v125 << (((v3 << ((v3) & 63))) & 63)), F1: "a"}
                                                                    v127 := v2[127]
                                                                    _ = v125
                                                                    _ = v126
                                                                    _ = v127
                                                                } else {
                                                                    v4 -= (((((v4 - 1e10) / (float64(3.14159) + v4)) - v4) + (((1.5 / 1.5) + v4) / v4)) + 1e10)
                                                                    v128 := GenS6{F0: 42}
                                                                    v129 := v1[999]
                                                                    v129 += 218
                                                                    var v130 int64 = 165
                                                                    _ = v128
                                                                    _ = v129
                                                                    _ = v130
                                                                }
                                                            }
                                                            var v131 int32 = 1
                                                            v132 := GenS12{F0: -1, F1: fi10}
                                                            var v133 int = (fi9 & 32768)
                                                            _ = len(v5)
                                                            _ = v131
                                                            _ = v132
                                                            _ = v133
                                                        }
                                                        _ = v122
                                                    } else {
                                                        var v134 float64 = (((v4 / v4) - -2.25) * float64(((3.14159 * ((-2.25 / (((float32(float64(float32((1.5 * v4)))) * 3.14159) * (0.0 / 1.0)) / (0.0 - 1e10))) + float32(float64(1.5)))) * 1.5)))
                                                        v4 = ((v134 * -2.25) * v4)
                                                        _ = v134
                                                    }
                                                } else {
                                                    v135 := GenS13{F0: 2, F1: -720, F2: 37}
                                                    var v136 int32 = -2147483648
                                                    _ = v135
                                                    _ = v136
                                                }
                                                var v137 uint16 = 42
                                                _ = v137
                                            } else {
                                                var v138 uint = (^(v3))
                                                if true {
                                                    v4 = v4
                                                    var v139 int64 = 1
                                                    _ = v139
                                                }
                                                switch "\n\t" {
                                                case "a":
                                                    switch "\n\t" {
                                                    case "":
                                                        var v140 float64 = (v4 * v4)
                                                        _ = v140
                                                    case "a":
                                                        fmt.Println(v3, fi9, v3)
                                                        v138 = v138
                                                    case "hello":
                                                        var v141 uint32 = 0
                                                        _ = v141
                                                    default:
                                                        if ((v138 << ((399) & 63)) == 128) {
                                                            if true {
                                                                switch 128 {
                                                                case -1:
                                                                    v142 := GenS3{F0: 4294967295}
                                                                    _ = v142
                                                                default:
                                                                    if len(v5) >= 2 {
                                                                        v5 = v5[1:len(v5)]
                                                                    }
                                                                }
                                                                var v143 int16 = 1
                                                                _ = v143
                                                            } else {
                                                                if len(v5) > 0 {
                                                                    v5[0] = (fi9 - 1)
                                                                }
                                                                if false {
                                                                    v144 := GenS14{F0: 42}
                                                                    _ = v144
                                                                } else {
                                                                    switch (-(fi9)) {
                                                                    case 32767:
                                                                        switch (fi9 >> ((v3) & 63)) {
                                                                        case -123:
                                                                            if false {
                                                                                var v145 uint8 = 255
                                                                                delete(v2, fi9)
                                                                                var v146 uint8 = v145
                                                                                _ = v145
                                                                                _ = v146
                                                                            } else {
                                                                                var v147 int = fi9
                                                                                v147 += v147
                                                                                v138 = v138
                                                                                var v148 uint16 = 32767
                                                                                _ = v147
                                                                                _ = v148
                                                                            }
                                                                        case 128:
                                                                            var v149 uint32 = 2147483648
                                                                            v5 = append(v5, 32767)
                                                                            if (!(!true)) {
                                                                                v138 -= (v3 % ((v3) | 1))
                                                                            } else {
                                                                                fmt.Println(fi9)
                                                                                v150 := GenS7{F0: 42}
                                                                                switch "" {
                                                                                case "hello":
                                                                                    if len(v5) >= 2 {
                                                                                        v5 = v5[1:len(v5)]
                                                                                    }
                                                                                default:
                                                                                    var v151 uint32 = (v149 & (v149 + v149))
                                                                                    _ = v151
                                                                                }
                                                                                _ = v150
                                                                            }
                                                                            _ = v149
                                                                        default:
                                                                            if (true && (!false)) {
                                                                                v152 := GenS1{F0: 65535, F1: fi10}
                                                                                _ = v152
                                                                            } else {
                                                                                v138 = (v138 >> ((535) & 63))
                                                                            }
                                                                        }
                                                                    case 9223372036854775807:
                                                                        var v153 uint32 = 2147483648
                                                                        _ = v153
                                                                    case 1:
                                                                        v154 := v1[51]
                                                                        _ = v154
                                                                    default:
                                                                        v155 := GenS12{F0: -128, F1: (fi10 << ((v3) & 63))}
                                                                        _ = v155
                                                                    }
                                                                    var v156 int64 = 42
                                                                    _ = v156
                                                                }
                                                            }
                                                        } else {
                                                            fmt.Println(v4, v3, v3)
                                                        }
                                                    }
                                                case "hello":
                                                    v2[(-(fi9))] = (fi10 * fi9)
                                                case "\n\t":
                                                    if false {
                                                        v157 := GenS15{F0: 2147483647}
                                                        switch "" {
                                                        case "\n\t":
                                                            var v158 float32 = 3.14159
                                                            _ = v158
                                                        case "foo bar":
                                                            switch ("a" + ((("a" + ((("\n\t" + ("hello" + "")) + ("foo bar" + "hello")) + ("\n\t" + ""))) + ("foo bar" + ((("a" + "foo bar") + "") + "hello"))) + ("foo bar" + ""))) {
                                                            case "hello":
                                                                var v159 uint32 = 949
                                                                _ = v159
                                                            default:
                                                                v160 := GenS12{F0: 0, F1: 1}
                                                                _ = v160
                                                            }
                                                        default:
                                                            var v161 uint8 = 69
                                                            _ = v161
                                                        }
                                                        _ = v157
                                                    } else {
                                                        v4 += float64(-2.25)
                                                        switch fi10 {
                                                        case 193:
                                                            v162 := GenS16{F0: fi9}
                                                            _ = v162
                                                        case 9223372036854775807:
                                                            v138 = (v3 << ((v3) & 63))
                                                        case 694:
                                                            if false {
                                                                var v163 int64 = int64(fi10)
                                                                v163 = v163
                                                                _ = v163
                                                            }
                                                        default:
                                                            v164 := v1[1]
                                                            var v165 uint32 = v164
                                                            var v166 string = ""
                                                            _ = v164
                                                            _ = v165
                                                            _ = v166
                                                        }
                                                    }
                                                }
                                                _ = v138
                                            }
                                        case "a":
                                            v5 = append(v5, (fi10 >> ((v3) & 63)))
                                        }
                                    }
                                } else {
                                    v3 = v3
                                    var v167 uint = 845
                                    switch "\n\t" {
                                    case "":
                                        var v168 uint16 = 212
                                        _ = v168
                                    case "foo bar":
                                        fmt.Println(v167, v4, v4)
                                    }
                                    fmt.Println(v167)
                                    fmt.Println(v4)
                                    var v169 uint32 = 1
                                    _ = v167
                                    _ = v169
                                }
                            case "\n\t":
                                v2[(fi9 - (fi9 + (fi10 % (((fi9 & fi9)) | 1))))] = 32768
                            case "":
                                var v170 uint16 = 186
                                _ = v170
                            }
                            v171 := GenS15{F0: -942}
                            _ = v171
                        } else {
                            var v172 int8 = int8(fi10)
                            _ = v172
                        }
                    case "foo bar":
                        v173 := GenS12{F0: -1, F1: fi10}
                        _ = v173
                    }
                    _ = fi10
                }
                v4 = float64(0.1)
                for fi174 := 0; fi174 < 18; fi174++ {
                    v4 = v4
                    _ = fi174
                }
            } else {
                if (18446744073709551615 > 0) {
                    v175 := GenS17{F0: fi9, F1: 0, F2: 127}
                    for fi176 := 0; fi176 < 12; fi176++ {
                        fmt.Println(v3)
                        v4 = v4
                        _ = fi176
                    }
                    _ = v175
                }
                switch fi9 {
                case 1:
                    var v177 uint = (v3 - v3)
                    _ = v177
                }
            }
            if len(v5) > 0 {
                v5[0] = -677
            }
            _ = fi9
        }
    case -730:
        if true {
            v4 = (v4 / 3.14159)
            var v178 uint = (v3 & 147)
            fmt.Println(v4, v4, v3)
            _ = v178
        }
    case 2147483647:
        var v179 uint = (v3 - v3)
        _ = v179
    }
    v3 = (v3 % ((v3) | 1))
    for fi180 := 0; fi180 < 14; fi180++ {
        for fi181 := 0; fi181 < 16; fi181++ {
            switch fi181 {
            case -204:
                switch -1 {
                case -2147483648:
                    switch "" {
                    case "a":
                        fmt.Println(v4, fi181)
                    case "\n\t":
                        v3 = 0
                    }
                case -1:
                    v182 := GenS14{F0: 2}
                    _ = v182
                default:
                    v3 = (v3 / ((v3) | 1))
                }
            case -2147483648:
                switch fi180 {
                case 9223372036854775807:
                    v183 := GenS18{F0: 2147483647, F1: 127, F2: 166}
                    _ = v183
                case -2147483648:
                    v3 -= 4294967295
                }
            }
            _ = fi181
        }
        var v184 int32 = -2147483648
        fmt.Println(v184, v3, v4)
        _ = fi180
        _ = v184
    }
    fmt.Println(v4)
    delete(v1, 1)
    var v185 bool = ((0.1 * (float32(3.14159) * (0.0 * (float32(1e10) + 0.1)))) == 3.14159)
    v186 := GenS16{F0: -9223372036854775808}
    fmt.Println(v4)
    var v187 uint32 = 2147483647
    if v185 {
        v187 -= v187
        if v185 {
            var v188 uint64 = 0
            fmt.Println(v185, v188)
            _ = v188
        }
        v189 := GenS1{F0: 1, F1: 777}
        _ = v189
    } else {
        var v190 float64 = (float64(0.1) * v4)
        switch ("hello" + "foo bar") {
        case "a":
            v187 -= v187
        case "\n\t":
            v191 := GenS19{F0: (v3 << (((^(v3))) & 63)), F1: -50}
            _ = v191
        case "hello":
            fmt.Println(v3, v187)
        default:
            v3 = (v3 % ((170) | 1))
        }
        _ = v190
    }
    var v192 int32 = -760
    if (!true) {
        v193 := v1[v187]
        if v185 {
            fmt.Println(v4)
        } else {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        }
        v194 := v1[v187]
        _ = v193
        _ = v194
    }
    fmt.Println(v4)
    v187 += (^(v187))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v185:", v185)
    fmt.Println("v187:", v187)
    fmt.Println("v192:", v192)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v185
    _ = v186
    _ = v187
    _ = v192
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
