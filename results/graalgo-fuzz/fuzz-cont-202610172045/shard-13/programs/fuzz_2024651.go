package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint32; F1 int8; F2 uint32 }
type GenS3 struct { F0 int32; F1 uint32 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint16; F1 int }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 int64; F1 bool; F2 int16 }
type GenS8 struct { F0 string; F1 bool }
type GenS9 struct { F0 uint8; F1 uint64; F2 int16 }
type GenS10 struct { F0 bool; F1 uint; F2 int }
type GenS11 struct { F0 int32; F1 uint16; F2 uint64 }
type GenS12 struct { F0 int32; F1 uint; F2 int64 }
type GenS13 struct { F0 int64 }
type GenS14 struct { F0 uint8; F1 uint; F2 uint8 }
type GenS15 struct { F0 uint }
type GenS16 struct { F0 string }
type GenS17 struct { F0 uint16; F1 int8 }
type GenS18 struct { F0 int32; F1 uint64; F2 uint64 }

func main() {
    v1 := []uint32{2, 1, 65535, 4294967295}
    var v2 int8 = -128
    v3 := []uint32{1, 186, 164, 2}
    var v4 float64 = (0.0 / 3.14159)
    v5 := GenS1{F0: 868}
    v2 -= (v2 * 100)
    var v6 uint16 = 2
    var v7 uint8 = 2
    fmt.Println(v6, v7)
    v8 := GenS1{F0: (v6 % ((v6) | 1))}
    var v9 float64 = (v4 / v4)
    if false {
        v4 = v4
        v10 := GenS1{F0: v6}
        if (!true) {
            v11 := GenS2{F0: 2, F1: -1, F2: 42}
            v12 := GenS3{F0: 42, F1: 0}
            _ = v11
            _ = v12
        } else {
            v7 += v7
            for fi13 := 0; fi13 < 5; fi13++ {
                v6 = 0
                var v14 uint = 18446744073709551615
                _ = fi13
                _ = v14
            }
            v2 = (v2 | 69)
        }
        _ = v10
    }
    for fi15 := 0; fi15 < 17; fi15++ {
        if false {
            var v16 float64 = v4
            var v17 uint16 = v6
            var v18 uint8 = v7
            _ = v16
            _ = v17
            _ = v18
        }
        if ((2147483647 <= 65535) && (((float32(0.1) * 0.1) * (1.5 + (0.0 / 1.5))) < 0.0)) {
            if ((!false) && ((!(((3.14159 * (0.0 * 0.1)) + 0.1) <= ((0.0 - 3.14159) + 1e10))) && (true && (false || false)))) {
                _ = len(v1)
            }
            var v19 int64 = 42
            _ = v19
        } else {
            v4 -= v9
        }
        for fi20 := 0; fi20 < 1; fi20++ {
            fmt.Println(v2, v2)
            fmt.Println(v6)
            _ = fi20
        }
        _ = fi15
    }
    if true {
        if false {
            if ((92 <= 1) || (false || (true && true))) {
                fmt.Println(v9, v9)
            } else {
                if len(v3) > 0 {
                    v3[0] = 2147483648
                }
                var v21 bool = true
                var v22 int16 = 1
                _ = v21
                _ = v22
            }
            for fi23 := 0; fi23 < 9; fi23++ {
                switch "a" {
                case "a":
                    if true {
                        if (-1 > 42) {
                            _ = len(v1)
                            var v24 int = 9223372036854775807
                            switch (v24 | (v24 + v24)) {
                            case 127:
                                var v25 int = (v24 | v24)
                                _ = v25
                            }
                            _ = v24
                        }
                    } else {
                        for fi26 := 0; fi26 < 17; fi26++ {
                            if (v7 >= v7) {
                                var v27 int32 = 2147483647
                                if true {
                                    if true {
                                        if false {
                                            if true {
                                                if true {
                                                    v27 -= (^(v27))
                                                    switch fi23 {
                                                    case 44:
                                                        v4 = (v4 / (float64(float32((v4 / v9))) - (v9 * v4)))
                                                    case 4294967295:
                                                        var v28 string = "hello"
                                                        _ = v28
                                                    default:
                                                        _ = len(v3)
                                                    }
                                                }
                                                switch "a" {
                                                case "a":
                                                    _ = len(v3)
                                                case "hello":
                                                    switch (fi26 ^ -491) {
                                                    case -2147483648:
                                                        v7 -= (^(v7))
                                                    case 127:
                                                        fmt.Println(v6, v4, fi26)
                                                    case 0:
                                                        var v29 uint16 = (v6 + 42)
                                                        _ = v29
                                                    }
                                                default:
                                                    switch fi23 {
                                                    case -902:
                                                        _ = len(v1)
                                                    case -2147483648:
                                                        var v30 int64 = -657
                                                        _ = v30
                                                    case 128:
                                                        var v31 uint16 = v6
                                                        _ = v31
                                                    default:
                                                        switch ("\n\t" + ("a" + ("foo bar" + ("hello" + "")))) {
                                                        case "hello":
                                                            if (("foo bar" + "hello") == "a") {
                                                                v6 += (^(v6))
                                                                v32 := GenS4{F0: 263}
                                                                _ = v32
                                                            } else {
                                                                if true {
                                                                    fmt.Println(v27)
                                                                    var v33 uint8 = 1
                                                                    _ = v33
                                                                } else {
                                                                    switch fi26 {
                                                                    case 2147483647:
                                                                        switch 475 {
                                                                        case 42:
                                                                            if true {
                                                                                var v34 int32 = (-(v27))
                                                                                v2 -= 37
                                                                                _ = v34
                                                                            }
                                                                        default:
                                                                            switch ("hello" + "a") {
                                                                            case "a":
                                                                                v9 = v9
                                                                            case "foo bar":
                                                                                v6 = 74
                                                                            case "":
                                                                                if (!(false || false)) {
                                                                                    var v35 float32 = 1.5
                                                                                    _ = v35
                                                                                } else {
                                                                                    var v36 int = -2147483648
                                                                                    _ = v36
                                                                                }
                                                                            }
                                                                            v37 := GenS5{F0: v6, F1: -625}
                                                                            if len(v3) >= 2 {
                                                                                v3 = v3[1:len(v3)]
                                                                            }
                                                                            _ = v37
                                                                        }
                                                                    case 42:
                                                                        var v38 string = "foo bar"
                                                                        _ = v38
                                                                    case -1:
                                                                        var v39 uint16 = (^(v6))
                                                                        _ = v39
                                                                    default:
                                                                        v7 = (^(v7))
                                                                    }
                                                                    var v40 int64 = 42
                                                                    _ = v40
                                                                }
                                                            }
                                                        case "a":
                                                            v2 = v2
                                                        case "":
                                                            var v41 uint8 = (v7 << ((42) & 63))
                                                            _ = v41
                                                        default:
                                                            v7 = 1
                                                        }
                                                    }
                                                }
                                            } else {
                                                if (4294967295 != 2) {
                                                    switch 2147483647 {
                                                    case 4294967295:
                                                        if len(v1) > 0 {
                                                            v1[0] = 4294967295
                                                        }
                                                    case -132:
                                                        if len(v3) > 0 {
                                                            v3[0] = 2147483647
                                                        }
                                                    default:
                                                        var v42 int8 = v2
                                                        switch (fi26 * 127) {
                                                        case -9223372036854775808:
                                                            if true {
                                                                if (42 != 7) {
                                                                    var v43 int = (fi26 ^ int(v7))
                                                                    if len(v1) >= 2 {
                                                                        v1 = v1[1:len(v1)]
                                                                    }
                                                                    _ = v43
                                                                }
                                                            }
                                                        case 42:
                                                            var v44 int32 = -2147483648
                                                            _ = v44
                                                        case 127:
                                                            v2 += v2
                                                        }
                                                        var v45 uint64 = 4294967295
                                                        var v46 uint64 = 53
                                                        _ = v42
                                                        _ = v45
                                                        _ = v46
                                                    }
                                                    fmt.Println(v2, v7)
                                                    if (true && false) {
                                                        v6 += (v6 ^ v6)
                                                        v27 -= (v27 + v27)
                                                        _ = len(v1)
                                                        fmt.Println(v9, v6)
                                                        if (false && (!(("foo bar" + "") <= ((("" + "a") + "a") + "foo bar")))) {
                                                            var v47 int32 = (v27 | (v27 | (v27 >> ((18446744073709551615) & 63))))
                                                            var v48 uint32 = 65535
                                                            v2 -= 127
                                                            v7 = 2
                                                            _ = v47
                                                            _ = v48
                                                        } else {
                                                            if false {
                                                                var v49 string = "foo bar"
                                                                var v50 int64 = 844
                                                                _ = v49
                                                                _ = v50
                                                            } else {
                                                                v51 := GenS1{F0: v6}
                                                                if true {
                                                                    if false {
                                                                        v3 = append(v3, 42)
                                                                        if false {
                                                                            if ("\n\t" != "\n\t") {
                                                                                switch 38 {
                                                                                case 586:
                                                                                    if ((!false) || ((v27 + 0) > v27)) {
                                                                                        var v52 float64 = 3.14159
                                                                                        _ = len(v3)
                                                                                        _ = v52
                                                                                    }
                                                                                }
                                                                            }
                                                                            var v53 int = 32767
                                                                            v27 -= (v27 | 214)
                                                                            _ = v53
                                                                        } else {
                                                                            fmt.Println(v27, v9)
                                                                            switch (fi26 - 0) {
                                                                            case 4294967295:
                                                                                if true {
                                                                                    v54 := GenS6{F0: 42}
                                                                                    if (350 != 56) {
                                                                                        v55 := GenS6{F0: -32768}
                                                                                        var v56 uint64 = 1
                                                                                        if len(v1) > 0 {
                                                                                            v1[0] = 42
                                                                                        }
                                                                                        v7 += (^(v7))
                                                                                        _ = v55
                                                                                        _ = v56
                                                                                    }
                                                                                    if false {
                                                                                        v6 = v6
                                                                                        v57 := GenS2{F0: 0, F1: v2, F2: 555}
                                                                                        _ = v57
                                                                                    } else {
                                                                                        var v58 string = "foo bar"
                                                                                        _ = v58
                                                                                    }
                                                                                    _ = v54
                                                                                } else {
                                                                                    switch "a" {
                                                                                    case "foo bar":
                                                                                        if len(v3) >= 2 {
                                                                                            v3 = v3[1:len(v3)]
                                                                                        }
                                                                                    case "":
                                                                                        fmt.Println(fi26)
                                                                                    }
                                                                                    v59 := GenS6{F0: 32767}
                                                                                    var v60 uint16 = 274
                                                                                    _ = v59
                                                                                    _ = v60
                                                                                }
                                                                            case -2147483648:
                                                                                if true {
                                                                                    fmt.Println(v9, v7)
                                                                                    v61 := GenS5{F0: (v6 % ((1) | 1)), F1: fi23}
                                                                                    _ = v61
                                                                                }
                                                                            default:
                                                                                switch fi23 {
                                                                                case 32767:
                                                                                    v4 = ((v4 + float64((float32((((3.14159 + (v9 * 3.14159)) + (v9 * 1e10)) + v9)) - 0.1))) / float64(1e10))
                                                                                case 4294967295:
                                                                                    var v62 string = "hello"
                                                                                    _ = v62
                                                                                case 42:
                                                                                    switch ((("a" + "\n\t") + "a") + "") {
                                                                                    case "hello":
                                                                                        v63 := GenS7{F0: 42, F1: false, F2: -1}
                                                                                        _ = v63
                                                                                    case "\n\t":
                                                                                        switch (fi26 & (fi23 >> ((636) & 63))) {
                                                                                        case 128:
                                                                                            var v64 int8 = int8(v27)
                                                                                            _ = v64
                                                                                        case 328:
                                                                                            var v65 uint = 294
                                                                                            _ = v65
                                                                                        case 883:
                                                                                            v4 -= v9
                                                                                        default:
                                                                                            v7 -= (v7 % ((2) | 1))
                                                                                        }
                                                                                    case "":
                                                                                        v66 := GenS8{F0: "", F1: (0.0 <= ((1e10 + 3.14159) * 1.5))}
                                                                                        _ = v66
                                                                                    }
                                                                                default:
                                                                                    var v67 int64 = 1
                                                                                    _ = v67
                                                                                }
                                                                            }
                                                                            fmt.Println(fi26, v6)
                                                                        }
                                                                        v6 = (v6 << ((0) & 63))
                                                                    } else {
                                                                        if false {
                                                                            var v68 uint64 = 974
                                                                            _ = v68
                                                                        } else {
                                                                            v69 := GenS8{F0: "", F1: (!false)}
                                                                            var v70 bool = (18446744073709551615 == 926)
                                                                            _ = len(v3)
                                                                            _ = v69
                                                                            _ = v70
                                                                        }
                                                                    }
                                                                } else {
                                                                    switch "\n\t" {
                                                                    case "a":
                                                                        var v71 uint16 = (^(v6))
                                                                        _ = v71
                                                                    case "hello":
                                                                        v2 -= v2
                                                                    case "\n\t":
                                                                        if false {
                                                                            switch -791 {
                                                                            case 0:
                                                                                if len(v3) > 0 {
                                                                                    v3[0] = 4294967295
                                                                                }
                                                                            case -1:
                                                                                v7 -= v7
                                                                            }
                                                                        }
                                                                    default:
                                                                        if true {
                                                                            if false {
                                                                                var v72 string = "foo bar"
                                                                                var v73 uint = 18446744073709551615
                                                                                _ = v72
                                                                                _ = v73
                                                                            } else {
                                                                                fmt.Println(v4, v2)
                                                                                _ = len(v1)
                                                                            }
                                                                            if ("foo bar" == "hello") {
                                                                                if (true && true) {
                                                                                    if len(v1) >= 2 {
                                                                                        v1 = v1[1:len(v1)]
                                                                                    }
                                                                                    v74 := GenS9{F0: (v7 ^ 79), F1: 0, F2: 32767}
                                                                                    if false {
                                                                                        var v75 int8 = (v2 << ((42) & 63))
                                                                                        _ = v75
                                                                                    }
                                                                                    _ = v74
                                                                                } else {
                                                                                    v4 = ((v9 / v9) * ((float64(float32(v4)) * float64((3.14159 + -2.25))) * ((3.14159 - v9) - (v9 + (3.14159 * 1e10)))))
                                                                                }
                                                                            }
                                                                            var v76 int64 = -1
                                                                            _ = v76
                                                                        }
                                                                    }
                                                                    v2 -= (v2 << ((4294967295) & 63))
                                                                    var v77 float32 = (float32(v9) + -2.25)
                                                                    _ = v77
                                                                }
                                                                var v78 float32 = 1e10
                                                                if len(v3) >= 2 {
                                                                    v3 = v3[1:len(v3)]
                                                                }
                                                                _ = v51
                                                                _ = v78
                                                            }
                                                        }
                                                    }
                                                    v27 = v27
                                                    v7 += v7
                                                    fmt.Println(v27, fi26, v9)
                                                } else {
                                                    v79 := GenS2{F0: 483, F1: v2, F2: 4294967295}
                                                    var v80 int16 = 666
                                                    v80 += v80
                                                    _ = v79
                                                    _ = v80
                                                }
                                            }
                                            v81 := GenS9{F0: v7, F1: 4294967295, F2: 1}
                                            v82 := GenS8{F0: "a", F1: true}
                                            _ = v81
                                            _ = v82
                                        }
                                    }
                                    fmt.Println(fi26)
                                    v83 := GenS5{F0: v6, F1: (^(fi26))}
                                    _ = v83
                                } else {
                                    var v84 int16 = 1000
                                    fmt.Println(v6, fi26, v27)
                                    v6 = 255
                                    _ = v84
                                }
                                var v85 uint32 = 106
                                _ = v27
                                _ = v85
                            } else {
                                if false {
                                    var v86 uint64 = 0
                                    switch ("a" + "\n\t") {
                                    case "hello":
                                        v87 := GenS3{F0: 646, F1: 2147483648}
                                        _ = v87
                                    case "foo bar":
                                        var v88 int64 = 9223372036854775807
                                        _ = v88
                                    default:
                                        v86 = v86
                                    }
                                    var v89 string = "a"
                                    _ = v86
                                    _ = v89
                                } else {
                                    v7 = (v7 << ((351) & 63))
                                    var v90 string = "\n\t"
                                    _ = v90
                                }
                                if true {
                                    if (!(float32(v4) >= 0.1)) {
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        var v91 float64 = v4
                                        _ = v91
                                    }
                                } else {
                                    switch fi23 {
                                    case -162:
                                        v6 -= (v6 << ((1) & 63))
                                    default:
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                    }
                                    var v92 uint = 398
                                    if false {
                                        if ((v6 + (v6 % ((v6) | 1))) == v6) {
                                            _ = len(v3)
                                            var v93 float64 = v4
                                            _ = v93
                                        } else {
                                            if false {
                                                v7 = v7
                                                v94 := GenS9{F0: v7, F1: 18446744073709551615, F2: 42}
                                                _ = v94
                                            } else {
                                                v6 += v6
                                            }
                                        }
                                    }
                                    _ = v92
                                }
                                fmt.Println(fi26, v7)
                            }
                            fmt.Println(v6, v7)
                            if true {
                                switch "foo bar" {
                                case "foo bar":
                                    var v95 string = ("" + "hello")
                                    _ = v95
                                default:
                                    switch "foo bar" {
                                    case "hello":
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                    case "foo bar":
                                        if false {
                                            if true {
                                                var v96 uint64 = 42
                                                v96 = v96
                                                v97 := GenS5{F0: 65535, F1: fi23}
                                                _ = v96
                                                _ = v97
                                            } else {
                                                v6 = (v6 * (v6 * 0))
                                                switch "\n\t" {
                                                case "foo bar":
                                                    if (1e10 != 3.14159) {
                                                        if (false || (32767 >= (fi26 % (((fi26 * fi23)) | 1)))) {
                                                            if true {
                                                                if false {
                                                                    v6 -= v6
                                                                }
                                                                v7 -= v7
                                                                var v98 uint8 = (v7 << ((409) & 63))
                                                                _ = v98
                                                            } else {
                                                                if false {
                                                                    switch (fi26 ^ fi23) {
                                                                    case 9223372036854775807:
                                                                        switch "\n\t" {
                                                                        case "a":
                                                                            if (!false) {
                                                                                if (v7 < (v7 & 202)) {
                                                                                    if len(v3) > 0 {
                                                                                        v3[0] = 0
                                                                                    }
                                                                                    if true {
                                                                                        if true {
                                                                                            v1 = append(v1, 77)
                                                                                        } else {
                                                                                            fmt.Println(v6)
                                                                                            v3 = append(v3, 2)
                                                                                        }
                                                                                    }
                                                                                    if (v4 >= (float64(float32(1.5)) * v9)) {
                                                                                        if (-2147483648 >= -923) {
                                                                                            var v99 int64 = -9223372036854775808
                                                                                            var v100 string = ("foo bar" + "a")
                                                                                            if true {
                                                                                                v2 -= v2
                                                                                                if (925 > 65535) {
                                                                                                    switch "\n\t" {
                                                                                                    case "foo bar":
                                                                                                        if true {
                                                                                                            v101 := GenS10{F0: true, F1: 42, F2: -2147483648}
                                                                                                            _ = v101
                                                                                                        }
                                                                                                    case "":
                                                                                                        v102 := GenS11{F0: 0, F1: (v6 * v6), F2: 42}
                                                                                                        switch (fi26 - (fi23 + (fi26 << ((871) & 63)))) {
                                                                                                        case 32767:
                                                                                                            if true {
                                                                                                                var v103 string = (v100 + v100)
                                                                                                                fmt.Println(v9, fi26, fi23)
                                                                                                                switch 4294967295 {
                                                                                                                case 9223372036854775807:
                                                                                                                    switch (fi23 + fi23) {
                                                                                                                    case -377:
                                                                                                                        v3 = append(v3, 2147483648)
                                                                                                                    case 819:
                                                                                                                        var v104 uint = 254
                                                                                                                        _ = v104
                                                                                                                    case 142:
                                                                                                                        switch (fi26 << ((2) & 63)) {
                                                                                                                        case 42:
                                                                                                                            if false {
                                                                                                                                v7 = (v7 >> ((42) & 63))
                                                                                                                                v7 -= (v7 % ((42) | 1))
                                                                                                                            } else {
                                                                                                                                if false {
                                                                                                                                    if true {
                                                                                                                                        var v105 uint64 = uint64(v6)
                                                                                                                                        _ = v105
                                                                                                                                    }
                                                                                                                                }
                                                                                                                            }
                                                                                                                        default:
                                                                                                                            v2 += (^(v2))
                                                                                                                            if false {
                                                                                                                                v2 = v2
                                                                                                                                v6 = 32767
                                                                                                                                var v106 uint32 = 637
                                                                                                                                _ = v106
                                                                                                                            }
                                                                                                                        }
                                                                                                                    default:
                                                                                                                        v107 := GenS12{F0: -2147483648, F1: 0, F2: -1}
                                                                                                                        _ = v107
                                                                                                                    }
                                                                                                                }
                                                                                                                _ = v103
                                                                                                            } else {
                                                                                                                if true {
                                                                                                                    var v108 uint = 1
                                                                                                                    _ = v108
                                                                                                                }
                                                                                                                v109 := GenS12{F0: -1, F1: 255, F2: v99}
                                                                                                                var v110 uint = 0
                                                                                                                switch -2147483648 {
                                                                                                                case 2147483647:
                                                                                                                    v99 = (v99 ^ v99)
                                                                                                                case -9223372036854775808:
                                                                                                                    var v111 int64 = v99
                                                                                                                    var v112 uint8 = (v7 / ((v7) | 1))
                                                                                                                    switch fi26 {
                                                                                                                    case -1:
                                                                                                                        var v113 int = (fi23 * (fi23 / ((fi23) | 1)))
                                                                                                                        _ = v113
                                                                                                                    }
                                                                                                                    _ = v111
                                                                                                                    _ = v112
                                                                                                                case -924:
                                                                                                                    fmt.Println(v7, v100, v99)
                                                                                                                default:
                                                                                                                    if (!(!(v99 < 1))) {
                                                                                                                        switch ("hello" + ("foo bar" + v100)) {
                                                                                                                        case "":
                                                                                                                            var v114 uint8 = 0
                                                                                                                            _ = v114
                                                                                                                        case "a":
                                                                                                                            var v115 int16 = -816
                                                                                                                            _ = v115
                                                                                                                        case "\n\t":
                                                                                                                            var v116 int = (fi23 - (fi26 ^ 32767))
                                                                                                                            _ = v116
                                                                                                                        }
                                                                                                                    }
                                                                                                                    if len(v3) > 0 {
                                                                                                                        v3[0] = 4294967295
                                                                                                                    }
                                                                                                                }
                                                                                                                _ = v109
                                                                                                                _ = v110
                                                                                                            }
                                                                                                        case 42:
                                                                                                            if len(v3) > 0 {
                                                                                                                v3[0] = 1
                                                                                                            }
                                                                                                        default:
                                                                                                            v6 = (v6 % (((v6 + (^(v6)))) | 1))
                                                                                                        }
                                                                                                        v117 := GenS13{F0: v99}
                                                                                                        _ = v102
                                                                                                        _ = v117
                                                                                                    default:
                                                                                                        _ = len(v1)
                                                                                                    }
                                                                                                    var v118 float64 = 3.14159
                                                                                                    v2 = v2
                                                                                                    _ = v118
                                                                                                } else {
                                                                                                    v4 = v4
                                                                                                }
                                                                                            }
                                                                                            _ = v100
                                                                                            _ = v99
                                                                                        }
                                                                                        v119 := GenS10{F0: ((v7 > v7) && false), F1: 4294967295, F2: fi23}
                                                                                        if (-1 == -1) {
                                                                                            v6 = v6
                                                                                        }
                                                                                        _ = v119
                                                                                    } else {
                                                                                        if true {
                                                                                            v120 := GenS14{F0: v7, F1: 4294967295, F2: (v7 + v7)}
                                                                                            v6 += v6
                                                                                            fmt.Println(fi23, v6, fi26)
                                                                                            if true {
                                                                                                fmt.Println(v4, v7, v2)
                                                                                                v2 += v2
                                                                                                v7 = 128
                                                                                                v6 += (v6 * v6)
                                                                                                if (!(((^(v6)) < v6) || (42 < -2147483648))) {
                                                                                                    fmt.Println(v2)
                                                                                                    v121 := GenS5{F0: v6, F1: -2147483648}
                                                                                                    v9 = v4
                                                                                                    _ = v121
                                                                                                }
                                                                                            }
                                                                                            _ = v120
                                                                                        }
                                                                                        if (true || false) {
                                                                                            v7 += (^(v7))
                                                                                            v6 = (v6 ^ 1)
                                                                                        } else {
                                                                                            v4 = (v4 - 0.0)
                                                                                        }
                                                                                    }
                                                                                }
                                                                            }
                                                                        case "":
                                                                            var v122 int16 = 42
                                                                            _ = v122
                                                                        case "hello":
                                                                            v123 := GenS15{F0: 255}
                                                                            _ = v123
                                                                        default:
                                                                            v7 -= v7
                                                                        }
                                                                    case 705:
                                                                        v4 += v9
                                                                    case 4294967295:
                                                                        v124 := GenS10{F0: (42 == 578), F1: 916, F2: (fi23 << ((255) & 63))}
                                                                        _ = v124
                                                                    default:
                                                                        var v125 int32 = -2147483648
                                                                        _ = v125
                                                                    }
                                                                    var v126 uint32 = 2147483648
                                                                    var v127 int32 = 0
                                                                    _ = v126
                                                                    _ = v127
                                                                }
                                                                if (2 != 2) {
                                                                    var v128 int8 = 100
                                                                    v129 := GenS6{F0: 603}
                                                                    _ = v128
                                                                    _ = v129
                                                                } else {
                                                                    v6 += 0
                                                                }
                                                                v130 := GenS16{F0: "foo bar"}
                                                                _ = v130
                                                            }
                                                        }
                                                    }
                                                default:
                                                    v131 := GenS10{F0: true, F1: 4294967295, F2: fi23}
                                                    v7 = (v7 / ((v7) | 1))
                                                    if true {
                                                        if (float64(3.14159) >= v9) {
                                                            v9 += v4
                                                        } else {
                                                            v1 = append(v1, 1)
                                                        }
                                                    } else {
                                                        v6 -= (v6 & v6)
                                                    }
                                                    _ = v131
                                                }
                                                var v132 int8 = v2
                                                _ = v132
                                            }
                                        } else {
                                            var v133 int64 = 1
                                            _ = v133
                                        }
                                    default:
                                        var v134 int64 = 53
                                        _ = v134
                                    }
                                }
                                v135 := GenS10{F0: false, F1: uint(v6), F2: fi23}
                                var v136 int16 = 263
                                _ = v135
                                _ = v136
                            }
                            _ = fi26
                        }
                        v137 := GenS17{F0: v6, F1: (^(v2))}
                        _ = v137
                    }
                }
                _ = fi23
            }
        } else {
            switch -9223372036854775808 {
            case 4294967295:
                v138 := GenS17{F0: (^(v6)), F1: v2}
                _ = v138
            case 127:
                if false {
                    v139 := GenS18{F0: -676, F1: 498, F2: 42}
                    var v140 uint32 = 2147483648
                    _ = v139
                    _ = v140
                }
            case -2147483648:
                var v141 int16 = 1
                _ = v141
            }
            v2 += (v2 % (((^(v2))) | 1))
        }
    }
    v142 := GenS14{F0: (^(v7)), F1: 42, F2: 0}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v142
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
