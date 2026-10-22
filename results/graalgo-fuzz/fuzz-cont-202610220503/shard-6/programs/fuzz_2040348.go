package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64; F2 int8 }
type GenS2 struct { F0 int32; F1 uint64 }
type GenS3 struct { F0 uint64; F1 uint16; F2 uint32 }
type GenS4 struct { F0 int; F1 uint }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 uint16; F1 int16; F2 int64 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 bool; F1 int32; F2 int }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 int64 }
type GenS11 struct { F0 int8; F1 bool }
type GenS12 struct { F0 int64 }
type GenS13 struct { F0 uint32; F1 int16; F2 int32 }
type GenS14 struct { F0 int32; F1 uint64; F2 uint8 }
type GenS15 struct { F0 bool }
type GenS16 struct { F0 int8; F1 int16 }
type GenS17 struct { F0 uint32 }
type GenS18 struct { F0 int32; F1 int32; F2 uint16 }
type GenS19 struct { F0 uint }

func main() {
    var v1 uint = 52
    v2 := make(map[uint32]string)
    v3 := []string{"a", ("hello" + "hello"), (("foo bar" + "foo bar") + "\n\t")}
    var v4 float64 = 0.0
    if (!(-1 < 676)) {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    } else {
        var v5 float64 = -2.25
        v6 := GenS1{F0: 42, F1: 42, F2: 0}
        _ = v5
        _ = v6
    }
    v7 := GenS1{F0: 4294967295, F1: -9223372036854775808, F2: 0}
    if (42 >= 100) {
        if false {
            if true {
                switch ("\n\t" + "hello") {
                case "foo bar":
                    switch -9223372036854775808 {
                    case -9223372036854775808:
                        for fi8 := 0; fi8 < 6; fi8++ {
                            v4 -= float64(3.14159)
                            switch "a" {
                            case "\n\t":
                                fmt.Println(fi8, v4)
                            case "hello":
                                switch fi8 {
                                case 32768:
                                    switch (-(fi8)) {
                                    case -9223372036854775808:
                                        var v9 int16 = 1
                                        _ = v9
                                    }
                                case -9223372036854775808:
                                    _ = len(v3)
                                }
                            case "foo bar":
                                if (true || false) {
                                    v3 = append(v3, "a")
                                    v4 = 3.14159
                                    if false {
                                        for fi10 := 0; fi10 < 4; fi10++ {
                                            if ("a" >= "hello") {
                                                v11 := GenS2{F0: -1, F1: 2}
                                                var v12 uint32 = 2147483648
                                                if false {
                                                    var v13 bool = false
                                                    var v14 float32 = 1.5
                                                    _ = v13
                                                    _ = v14
                                                }
                                                _ = v11
                                                _ = v12
                                            } else {
                                                delete(v2, uint32(v1))
                                                var v15 int64 = -9223372036854775808
                                                _ = v15
                                            }
                                            var v16 int8 = 100
                                            var v17 string = "hello"
                                            _ = fi10
                                            _ = v16
                                            _ = v17
                                        }
                                        if false {
                                            v18 := GenS3{F0: 2, F1: uint16(v1), F2: 0}
                                            _ = v18
                                        }
                                        v1 += (^(v1))
                                    } else {
                                        v2[684] = "hello"
                                    }
                                }
                            }
                            if ((3.14159 >= 1e10) && (((1e10 / float64(((1.5 / 0.1) + 0.0))) / 1.5) <= -2.25)) {
                                var v19 float32 = ((1e10 - 0.1) - (0.0 - 1e10))
                                v4 += v4
                                for fi20 := 0; fi20 < 18; fi20++ {
                                    v1 = v1
                                    v1 = v1
                                    _ = fi20
                                }
                                _ = v19
                            }
                            _ = fi8
                        }
                    default:
                        if false {
                            if ((v1 - v1) != 4294967295) {
                                switch "a" {
                                case "foo bar":
                                    switch -359 {
                                    case 32767:
                                        for fi21 := 0; fi21 < 8; fi21++ {
                                            fmt.Println(v4, v1, v1)
                                            _ = fi21
                                        }
                                    }
                                default:
                                    delete(v2, 4294967295)
                                }
                                v4 += (1e10 * (v4 - (float64(1.5) / 0.1)))
                            } else {
                                var v22 string = "\n\t"
                                var v23 float64 = (v4 * ((3.14159 + (float64(((0.1 - 0.0) + (3.14159 / float32(v4)))) - v4)) + v4))
                                fmt.Println(v22, v4)
                                _ = v22
                                _ = v23
                            }
                            v4 += ((3.14159 * (v4 - (v4 - v4))) + v4)
                            switch "" {
                            case "hello":
                                fmt.Println(v4, v1)
                            case "":
                                for fi24 := 0; fi24 < 9; fi24++ {
                                    v4 += (v4 - 1.5)
                                    switch ("a" + ("a" + "\n\t")) {
                                    case "\n\t":
                                        fmt.Println(v4, v4)
                                    case "":
                                        v4 = (v4 - v4)
                                    case "hello":
                                        var v25 int8 = 20
                                        _ = v25
                                    default:
                                        v26 := GenS4{F0: (fi24 ^ fi24), F1: (v1 & v1)}
                                        _ = v26
                                    }
                                    if (true || (-874 < -525)) {
                                        for fi27 := 0; fi27 < 9; fi27++ {
                                            if false {
                                                var v28 bool = false
                                                switch (fi27 % ((fi27) | 1)) {
                                                case 42:
                                                    fmt.Println(v4, fi27)
                                                default:
                                                    fmt.Println(v4, v1, fi27)
                                                }
                                                v29 := GenS3{F0: 18446744073709551615, F1: 65535, F2: 4294967295}
                                                v4 -= 1e10
                                                var v30 int = -2147483648
                                                _ = v28
                                                _ = v29
                                                _ = v30
                                            }
                                            delete(v2, 1)
                                            _ = fi27
                                        }
                                        var v31 uint32 = 598
                                        switch "foo bar" {
                                        case "a":
                                            var v32 uint64 = 4294967295
                                            _ = v32
                                        case "hello":
                                            if true {
                                                switch fi24 {
                                                case -1:
                                                    for fi33 := 0; fi33 < 9; fi33++ {
                                                        switch -19 {
                                                        case 2147483647:
                                                            var v34 uint16 = 725
                                                            _ = v34
                                                        case 32768:
                                                            switch (fi24 % (((fi24 ^ (fi24 - fi33))) | 1)) {
                                                            case 9223372036854775807:
                                                                if len(v3) > 0 {
                                                                    v3[0] = "hello"
                                                                }
                                                            case 42:
                                                                switch "a" {
                                                                case "a":
                                                                    v35 := GenS3{F0: 106, F1: 73, F2: v31}
                                                                    _ = v35
                                                                case "\n\t":
                                                                    v1 = (v1 + (v1 - 673))
                                                                default:
                                                                    var v36 int16 = 1000
                                                                    _ = v36
                                                                }
                                                            }
                                                        case 128:
                                                            v37 := GenS5{F0: 0}
                                                            _ = v37
                                                        }
                                                        _ = fi33
                                                    }
                                                case -2147483648:
                                                    switch (^(fi24)) {
                                                    case -126:
                                                        switch fi24 {
                                                        case -9223372036854775808:
                                                            if len(v3) > 0 {
                                                                v3[0] = "foo bar"
                                                            }
                                                        }
                                                    case 2147483647:
                                                        v38 := GenS5{F0: 26}
                                                        _ = v38
                                                    case -530:
                                                        for fi39 := 0; fi39 < 17; fi39++ {
                                                            v31 += v31
                                                            var v40 uint64 = 2
                                                            _ = fi39
                                                            _ = v40
                                                        }
                                                    }
                                                }
                                                if false {
                                                    v4 += v4
                                                    v41 := GenS1{F0: 0, F1: -351, F2: 38}
                                                    if (-32768 >= -809) {
                                                        v31 = 2147483648
                                                        v31 -= v31
                                                        v1 -= v1
                                                    } else {
                                                        v2[42] = ("\n\t" + "\n\t")
                                                    }
                                                    _ = v41
                                                } else {
                                                    var v42 int8 = 97
                                                    if false {
                                                        var v43 uint16 = 32767
                                                        switch fi24 {
                                                        case -9223372036854775808:
                                                            v42 += (v42 / ((v42) | 1))
                                                        }
                                                        _ = v43
                                                    } else {
                                                        if false {
                                                            var v44 uint16 = 42
                                                            _ = v44
                                                        } else {
                                                            v45 := GenS6{F0: uint16(v1), F1: 0, F2: 42}
                                                            v46 := GenS2{F0: 113, F1: 42}
                                                            _ = v45
                                                            _ = v46
                                                        }
                                                    }
                                                    switch ("" + "foo bar") {
                                                    case "\n\t":
                                                        if (!false) {
                                                            v1 = 255
                                                            v4 = (v4 * ((v4 * float64((0.1 / 1e10))) - (v4 * v4)))
                                                        } else {
                                                            fmt.Println(v31)
                                                        }
                                                    }
                                                    _ = v42
                                                }
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                            } else {
                                                var v47 string = ("" + "\n\t")
                                                switch v47 {
                                                case "hello":
                                                    v48 := GenS2{F0: -2147483648, F1: 42}
                                                    _ = v48
                                                case "a":
                                                    for fi49 := 0; fi49 < 11; fi49++ {
                                                        _ = len(v3)
                                                        if false {
                                                            var v50 uint16 = 571
                                                            var v51 int32 = 930
                                                            v51 -= v51
                                                            _ = v50
                                                            _ = v51
                                                        } else {
                                                            var v52 int64 = 1
                                                            _ = len(v2)
                                                            v1 += v1
                                                            _ = v52
                                                        }
                                                        _ = fi49
                                                    }
                                                case "foo bar":
                                                    v47 = (v47 + "foo bar")
                                                }
                                                _ = v47
                                            }
                                        }
                                        _ = v31
                                    } else {
                                        if len(v3) > 0 {
                                            v3[0] = "foo bar"
                                        }
                                        if (false && true) {
                                            switch "a" {
                                            case "\n\t":
                                                if false {
                                                    _ = len(v2)
                                                } else {
                                                    for fi53 := 0; fi53 < 18; fi53++ {
                                                        var v54 int = fi24
                                                        fmt.Println(v4, fi53, v4)
                                                        v3 = append(v3, "foo bar")
                                                        _ = fi53
                                                        _ = v54
                                                    }
                                                    var v55 float32 = -2.25
                                                    _ = v55
                                                }
                                            case "hello":
                                                for fi56 := 0; fi56 < 1; fi56++ {
                                                    if (2147483647 <= 227) {
                                                        if true {
                                                            if true {
                                                                var v57 float64 = v4
                                                                _ = v57
                                                            } else {
                                                                switch ("\n\t" + ("" + "")) {
                                                                case "foo bar":
                                                                    var v58 float32 = (0.0 / 0.1)
                                                                    var v59 float64 = (3.14159 * 0.1)
                                                                    _ = v58
                                                                    _ = v59
                                                                case "a":
                                                                    switch -363 {
                                                                    case 128:
                                                                        v60 := GenS3{F0: 2, F1: 1, F2: 42}
                                                                        _ = v60
                                                                    case -9223372036854775808:
                                                                        switch (fi24 >> ((v1) & 63)) {
                                                                        case 805:
                                                                            var v61 uint64 = 18446744073709551615
                                                                            _ = v61
                                                                        case 9223372036854775807:
                                                                            var v62 uint = v1
                                                                            _ = v62
                                                                        }
                                                                    }
                                                                case "":
                                                                    switch (-(fi56)) {
                                                                    case -966:
                                                                        fmt.Println(v1, fi56)
                                                                        v63 := GenS1{F0: 2, F1: 9223372036854775807, F2: -1}
                                                                        if false {
                                                                            v64 := GenS7{F0: (float32(3.14159) < 3.14159)}
                                                                            v1 += (v1 * (v1 & (v1 - v1)))
                                                                            _ = v64
                                                                        } else {
                                                                            var v65 int64 = -1
                                                                            _ = v65
                                                                        }
                                                                        _ = v63
                                                                    case 32768:
                                                                        v1 = (^(v1))
                                                                    default:
                                                                        v1 -= (^(v1))
                                                                    }
                                                                }
                                                                var v66 int32 = 2147483647
                                                                v4 = (((v4 + (float64(3.14159) + v4)) - float64(1e10)) * (v4 - v4))
                                                                var v67 uint32 = 2147483647
                                                                _ = len(v3)
                                                                _ = v66
                                                                _ = v67
                                                            }
                                                        } else {
                                                            if false {
                                                                v4 = (v4 - 3.14159)
                                                                v4 -= 1.5
                                                                v4 -= (v4 * 1.5)
                                                                var v68 int8 = 1
                                                                var v69 int16 = -92
                                                                var v70 uint8 = 0
                                                                switch -696 {
                                                                case -9223372036854775808:
                                                                    if len(v3) >= 2 {
                                                                        v3 = v3[1:len(v3)]
                                                                    }
                                                                default:
                                                                    fmt.Println(v70, v70, fi24)
                                                                }
                                                                _ = v68
                                                                _ = v69
                                                                _ = v70
                                                            }
                                                            if (1e10 == 1e10) {
                                                                v1 += v1
                                                            }
                                                            if (!true) {
                                                                switch "" {
                                                                case "\n\t":
                                                                    if true {
                                                                        v1 = v1
                                                                        v4 = 0.0
                                                                        switch ("\n\t" + ((("a" + "\n\t") + "hello") + ("a" + ("\n\t" + ((("hello" + (("foo bar" + "foo bar") + "hello")) + "a") + "hello"))))) {
                                                                        case "":
                                                                            fmt.Println(fi24)
                                                                        }
                                                                    }
                                                                    v71 := v2[42]
                                                                    var v72 int16 = -39
                                                                    _ = v71
                                                                    _ = v72
                                                                default:
                                                                    _ = len(v3)
                                                                }
                                                                switch "hello" {
                                                                case "hello":
                                                                    if len(v3) > 0 {
                                                                        v3[0] = ("a" + ("hello" + "hello"))
                                                                    }
                                                                case "":
                                                                    if len(v3) > 0 {
                                                                        v3[0] = ("" + "")
                                                                    }
                                                                case "\n\t":
                                                                    var v73 bool = true
                                                                    _ = v73
                                                                default:
                                                                    if (!(true || (!(1000 != 959)))) {
                                                                        if true {
                                                                            fmt.Println(fi56, fi24, v1)
                                                                            var v74 uint16 = 2
                                                                            v4 = v4
                                                                            if ((((0.1 - 0.0) - 3.14159) < (3.14159 / -2.25)) && true) {
                                                                                v4 = float64((-2.25 - (float32(v4) + ((0.1 + 0.0) * (1.5 * 0.0)))))
                                                                                v74 += v74
                                                                            }
                                                                            var v75 uint64 = 1
                                                                            var v76 int64 = -1
                                                                            var v77 uint8 = 5
                                                                            _ = v74
                                                                            _ = v75
                                                                            _ = v76
                                                                            _ = v77
                                                                        }
                                                                        fmt.Println(fi56, fi56)
                                                                        if (2 <= 42) {
                                                                            if false {
                                                                                var v78 uint16 = 42
                                                                                _ = v78
                                                                            } else {
                                                                                var v79 int64 = 9223372036854775807
                                                                                switch "\n\t" {
                                                                                case "foo bar":
                                                                                    var v80 bool = (1 != 32767)
                                                                                    _ = v80
                                                                                case "\n\t":
                                                                                    switch "" {
                                                                                    case "\n\t":
                                                                                        if (0 != 1) {
                                                                                            var v81 string = "\n\t"
                                                                                            v4 = float64(1e10)
                                                                                            v4 = (((float64(1.5) / -2.25) * (v4 * 1e10)) * 0.0)
                                                                                            _ = v81
                                                                                        }
                                                                                    case "a":
                                                                                        switch "\n\t" {
                                                                                        case "a":
                                                                                            if len(v3) >= 2 {
                                                                                                v3 = v3[1:len(v3)]
                                                                                            }
                                                                                        case "hello":
                                                                                            v82 := GenS8{F0: (v1 == 18446744073709551615), F1: 42, F2: fi56}
                                                                                            _ = v82
                                                                                        case "":
                                                                                            if (0 != 42) {
                                                                                                _ = len(v2)
                                                                                            } else {
                                                                                                v2[2147483647] = "foo bar"
                                                                                                if (-32768 >= 0) {
                                                                                                    if ((!(((-1 < -128) || false) || false)) || false) {
                                                                                                        switch "hello" {
                                                                                                        case "\n\t":
                                                                                                            var v83 uint64 = 2
                                                                                                            _ = v83
                                                                                                        case "hello":
                                                                                                            if true {
                                                                                                                v84 := GenS9{F0: false}
                                                                                                                v2[921] = (("foo bar" + (("" + "foo bar") + "foo bar")) + ("hello" + "a"))
                                                                                                                _ = len(v3)
                                                                                                                var v85 float32 = 0.0
                                                                                                                _ = v84
                                                                                                                _ = v85
                                                                                                            } else {
                                                                                                                v1 = v1
                                                                                                            }
                                                                                                        case "a":
                                                                                                            fmt.Println(fi56, v1, v4)
                                                                                                        }
                                                                                                    }
                                                                                                    switch (-(fi56)) {
                                                                                                    case 32768:
                                                                                                        v1 += v1
                                                                                                    default:
                                                                                                        v86 := GenS4{F0: 4294967295, F1: 18446744073709551615}
                                                                                                        _ = v86
                                                                                                    }
                                                                                                    if (2147483647 <= -230) {
                                                                                                        if ((!(false || (false && ((((true || (v1 >= 255)) && (float64(0.0) <= 0.1)) && true) && false)))) || false) {
                                                                                                            var v87 int8 = 100
                                                                                                            var v88 uint16 = 255
                                                                                                            var v89 int64 = v79
                                                                                                            _ = v87
                                                                                                            _ = v88
                                                                                                            _ = v89
                                                                                                        } else {
                                                                                                            v1 = v1
                                                                                                            if (1 <= 42) {
                                                                                                                var v90 float32 = (float32(3.14159) + -2.25)
                                                                                                                if (42 <= 206) {
                                                                                                                    var v91 string = ("" + "")
                                                                                                                    v1 = v1
                                                                                                                    _ = v91
                                                                                                                } else {
                                                                                                                    v1 = (v1 % ((v1) | 1))
                                                                                                                    fmt.Println(fi24, fi56, fi24)
                                                                                                                }
                                                                                                                v92 := v2[0]
                                                                                                                _ = v90
                                                                                                                _ = v92
                                                                                                            }
                                                                                                        }
                                                                                                        fmt.Println(v1, v1)
                                                                                                        v4 += v4
                                                                                                    }
                                                                                                    v93 := GenS10{F0: -1}
                                                                                                    _ = v93
                                                                                                }
                                                                                                var v94 float64 = v4
                                                                                                _ = v94
                                                                                            }
                                                                                        }
                                                                                    case "":
                                                                                        if false {
                                                                                            fmt.Println(v4, v1)
                                                                                        } else {
                                                                                            var v95 string = ""
                                                                                            var v96 int16 = -86
                                                                                            v97 := GenS10{F0: (-(v79))}
                                                                                            var v98 int8 = 127
                                                                                            _ = v95
                                                                                            _ = v96
                                                                                            _ = v97
                                                                                            _ = v98
                                                                                        }
                                                                                    }
                                                                                case "hello":
                                                                                    switch "\n\t" {
                                                                                    case "hello":
                                                                                        v99 := GenS11{F0: -1, F1: true}
                                                                                        _ = v99
                                                                                    default:
                                                                                        switch (fi24 / (((fi24 ^ fi24)) | 1)) {
                                                                                        case 701:
                                                                                            v100 := GenS10{F0: (v79 >> ((65535) & 63))}
                                                                                            _ = v100
                                                                                        case -1:
                                                                                            if true {
                                                                                                v1 += (v1 / ((160) | 1))
                                                                                                if (!(100 > 1)) {
                                                                                                    v101 := GenS12{F0: 1}
                                                                                                    _ = v101
                                                                                                }
                                                                                                var v102 uint8 = 147
                                                                                                _ = v102
                                                                                            } else {
                                                                                                var v103 uint8 = 58
                                                                                                if true {
                                                                                                    fmt.Println(fi24, fi56, v4)
                                                                                                    fmt.Println(v103)
                                                                                                    fmt.Println(v4, fi24)
                                                                                                }
                                                                                                v104 := GenS13{F0: 2147483647, F1: -1, F2: 1}
                                                                                                _ = v103
                                                                                                _ = v104
                                                                                            }
                                                                                        case 32767:
                                                                                            v105 := GenS6{F0: 42, F1: 42, F2: -9223372036854775808}
                                                                                            _ = v105
                                                                                        }
                                                                                    }
                                                                                }
                                                                                _ = v79
                                                                            }
                                                                            var v106 int8 = -128
                                                                            _ = v106
                                                                        } else {
                                                                            v1 -= v1
                                                                        }
                                                                    }
                                                                }
                                                                v2[1] = "foo bar"
                                                                v4 -= float64(-2.25)
                                                            } else {
                                                                var v107 string = "hello"
                                                                fmt.Println(v4, fi24)
                                                                _ = v107
                                                            }
                                                        }
                                                        v4 += (((v4 / -2.25) / v4) / 3.14159)
                                                        fmt.Println(v4)
                                                    }
                                                    switch fi56 {
                                                    case 841:
                                                        var v108 int = (^(fi24))
                                                        _ = v108
                                                    default:
                                                        if false {
                                                            var v109 uint32 = 42
                                                            if true {
                                                                if (1 > 42) {
                                                                    var v110 uint32 = 42
                                                                    var v111 bool = (true || (!true))
                                                                    _ = v110
                                                                    _ = v111
                                                                }
                                                            } else {
                                                                var v112 int8 = 127
                                                                _ = v112
                                                            }
                                                            _ = v109
                                                        }
                                                    }
                                                    _ = fi56
                                                }
                                            case "foo bar":
                                                v113 := GenS13{F0: 2147483647, F1: 1, F2: -1}
                                                _ = v113
                                            }
                                        } else {
                                            if (false && true) {
                                                switch "a" {
                                                case "":
                                                    var v114 uint32 = 815
                                                    _ = v114
                                                case "a":
                                                    if false {
                                                        fmt.Println(v4, v1)
                                                        v1 = 826
                                                        var v115 uint8 = 0
                                                        _ = v115
                                                    }
                                                }
                                            } else {
                                                v1 -= (v1 + 42)
                                                switch ("" + "hello") {
                                                case "a":
                                                    fmt.Println(v1, v1, fi24)
                                                case "hello":
                                                    for fi116 := 0; fi116 < 10; fi116++ {
                                                        v1 = 18446744073709551615
                                                        v4 = v4
                                                        _ = fi116
                                                    }
                                                case "":
                                                    v117 := GenS14{F0: 2147483647, F1: 1, F2: uint8(v1)}
                                                    _ = v117
                                                }
                                                switch "\n\t" {
                                                case "foo bar":
                                                    var v118 uint64 = 4294967295
                                                    _ = v118
                                                case "\n\t":
                                                    if (-1 == -751) {
                                                        var v119 int16 = 42
                                                        v120 := GenS6{F0: 2, F1: v119, F2: 42}
                                                        _ = v119
                                                        _ = v120
                                                    } else {
                                                        v4 -= (v4 - 1.5)
                                                    }
                                                case "a":
                                                    v1 = 2
                                                default:
                                                    if true {
                                                        for fi121 := 0; fi121 < 20; fi121++ {
                                                            if true {
                                                                var v122 float32 = (0.0 / float32(3.14159))
                                                                switch "hello" {
                                                                case "hello":
                                                                    var v123 uint8 = 255
                                                                    _ = v123
                                                                case "\n\t":
                                                                    if true {
                                                                        var v124 string = "a"
                                                                        _ = v124
                                                                    } else {
                                                                        switch "foo bar" {
                                                                        case "\n\t":
                                                                            v1 = v1
                                                                            v4 += float64(v122)
                                                                        case "hello":
                                                                            v125 := GenS6{F0: 771, F1: 0, F2: 572}
                                                                            _ = v125
                                                                        }
                                                                        v126 := GenS12{F0: 9223372036854775807}
                                                                        _ = v126
                                                                    }
                                                                default:
                                                                    var v127 int16 = 1000
                                                                    _ = v127
                                                                }
                                                                _ = v122
                                                            } else {
                                                                fmt.Println(fi24, v1, fi24)
                                                                if (v4 != (v4 + -2.25)) {
                                                                    switch "hello" {
                                                                    case "a":
                                                                        var v128 float32 = (1e10 / float32(v4))
                                                                        _ = v128
                                                                    case "":
                                                                        var v129 int64 = -504
                                                                        _ = v129
                                                                    default:
                                                                        delete(v2, 36)
                                                                    }
                                                                    _ = len(v2)
                                                                    v130 := GenS12{F0: 0}
                                                                    _ = v130
                                                                } else {
                                                                    v131 := v2[663]
                                                                    v132 := GenS10{F0: -9223372036854775808}
                                                                    var v133 uint8 = 42
                                                                    _ = v131
                                                                    _ = v132
                                                                    _ = v133
                                                                }
                                                                v134 := GenS4{F0: -9223372036854775808, F1: 42}
                                                                _ = v134
                                                            }
                                                            v1 += (^(v1))
                                                            switch -1 {
                                                            case 364:
                                                                v4 -= (v4 / 1.5)
                                                            default:
                                                                fmt.Println(v4, fi121)
                                                            }
                                                            var v135 uint32 = 65535
                                                            _ = fi121
                                                            _ = v135
                                                        }
                                                        if len(v3) > 0 {
                                                            v3[0] = "\n\t"
                                                        }
                                                        v136 := GenS11{F0: -122, F1: false}
                                                        _ = v136
                                                    } else {
                                                        for fi137 := 0; fi137 < 4; fi137++ {
                                                            v3 = append(v3, "a")
                                                            v4 = 1e10
                                                            var v138 int64 = -1
                                                            _ = fi137
                                                            _ = v138
                                                        }
                                                        delete(v2, 573)
                                                    }
                                                }
                                            }
                                            if false {
                                                v139 := GenS8{F0: (1.5 > ((3.14159 + 0.0) * -2.25)), F1: -22, F2: fi24}
                                                v1 = 42
                                                var v140 int = 32768
                                                _ = v139
                                                _ = v140
                                            }
                                        }
                                    }
                                    _ = fi24
                                }
                            case "foo bar":
                                for fi141 := 0; fi141 < 12; fi141++ {
                                    v1 += v1
                                    _ = fi141
                                }
                            default:
                                var v142 int64 = 102
                                _ = v142
                            }
                        } else {
                            for fi143 := 0; fi143 < 10; fi143++ {
                                if true {
                                    for fi144 := 0; fi144 < 18; fi144++ {
                                        v4 = (v4 + float64(3.14159))
                                        _ = fi144
                                    }
                                }
                                var v145 bool = (true || false)
                                _ = fi143
                                _ = v145
                            }
                            if false {
                                var v146 uint8 = 1
                                if (-116 > 1) {
                                    _ = len(v3)
                                    for fi147 := 0; fi147 < 10; fi147++ {
                                        _ = len(v3)
                                        var v148 float32 = 3.14159
                                        _ = fi147
                                        _ = v148
                                    }
                                } else {
                                    var v149 uint64 = 18446744073709551615
                                    _ = v149
                                }
                                if (!false) {
                                    _ = len(v3)
                                    for fi150 := 0; fi150 < 11; fi150++ {
                                        var v151 bool = true
                                        if (0 != 42) {
                                            var v152 uint8 = (^(v146))
                                            for fi153 := 0; fi153 < 15; fi153++ {
                                                v152 -= (^(v146))
                                                if (((("\n\t" + "\n\t") + "hello") >= "foo bar") || v151) {
                                                    v154 := GenS4{F0: fi153, F1: (v1 / (((^(v1))) | 1))}
                                                    _ = v154
                                                } else {
                                                    var v155 int64 = 42
                                                    var v156 bool = v151
                                                    switch "hello" {
                                                    case "":
                                                        v146 -= 2
                                                    case "\n\t":
                                                        v151 = v156
                                                    case "hello":
                                                        switch 0 {
                                                        case 127:
                                                            var v157 uint32 = 2147483647
                                                            _ = v157
                                                        case 4294967295:
                                                            v158 := GenS15{F0: v156}
                                                            _ = v158
                                                        case 247:
                                                            fmt.Println(v155, v146)
                                                        }
                                                        var v159 float32 = (0.1 / 0.1)
                                                        if v151 {
                                                            fmt.Println(v151, v159)
                                                            v151 = v151
                                                        }
                                                        _ = v159
                                                    }
                                                    _ = v155
                                                    _ = v156
                                                }
                                                _ = fi153
                                            }
                                            _ = v152
                                        }
                                        _ = fi150
                                        _ = v151
                                    }
                                    var v160 int8 = 0
                                    _ = v160
                                } else {
                                    v1 = v1
                                    var v161 int = 4294967295
                                    for fi162 := 0; fi162 < 14; fi162++ {
                                        for fi163 := 0; fi163 < 8; fi163++ {
                                            v164 := GenS13{F0: 761, F1: -1, F2: -2147483648}
                                            _ = fi163
                                            _ = v164
                                        }
                                        _ = fi162
                                    }
                                    _ = v161
                                }
                                _ = v146
                            } else {
                                fmt.Println(v1, v1)
                            }
                            var v165 bool = (false && false)
                            _ = v165
                        }
                    }
                case "":
                    v166 := v2[65535]
                    _ = v166
                default:
                    v167 := GenS14{F0: 567, F1: 42, F2: 128}
                    _ = v167
                }
            } else {
                _ = len(v3)
                for fi168 := 0; fi168 < 3; fi168++ {
                    fmt.Println(v4, v4, v1)
                    _ = fi168
                }
            }
        } else {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            if (0 != 0) {
                var v169 int16 = 1
                _ = v169
            }
        }
        for fi170 := 0; fi170 < 19; fi170++ {
            _ = len(v2)
            v171 := GenS16{F0: -61, F1: int16(fi170)}
            if (!true) {
                v4 -= v4
                switch ("a" + "a") {
                case "a":
                    _ = len(v2)
                case "foo bar":
                    v4 = v4
                case "":
                    for fi172 := 0; fi172 < 14; fi172++ {
                        v173 := GenS17{F0: uint32(v1)}
                        _ = fi172
                        _ = v173
                    }
                }
                switch (("hello" + "hello") + "foo bar") {
                case "":
                    fmt.Println(v4, v1, v4)
                case "a":
                    v1 = (v1 + (v1 & (v1 / ((v1) | 1))))
                default:
                    v4 += v4
                }
            }
            _ = fi170
            _ = v171
        }
    }
    switch "" {
    case "":
        v4 += 1.5
    }
    v3 = append(v3, (("\n\t" + ("a" + ("\n\t" + "hello"))) + (("" + ("\n\t" + "a")) + "a")))
    if (42 <= 42) {
        if true {
            var v174 float64 = 1.5
            _ = v174
        }
        fmt.Println(v4)
    }
    if len(v3) > 0 {
        v3[0] = ""
    }
    v175 := GenS18{F0: 42, F1: -793, F2: 159}
    if ((v1 ^ v1) == v1) {
        _ = len(v3)
    }
    var v176 int16 = 0
    fmt.Println(v176)
    v1 += v1
    v177 := GenS19{F0: 27}
    var v178 float32 = 0.0
    fmt.Println("v1:", v1)
    fmt.Println("v176:", v176)
    fmt.Println("v178:", v178)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v175
    _ = v176
    _ = v177
    _ = v178
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
