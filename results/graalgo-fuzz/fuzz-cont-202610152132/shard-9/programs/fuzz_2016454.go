package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64 }
type GenS2 struct { F0 uint32; F1 int8; F2 int16 }
type GenS3 struct { F0 bool; F1 int16; F2 int8 }
type GenS4 struct { F0 uint64; F1 uint8 }
type GenS5 struct { F0 uint; F1 int16 }
type GenS6 struct { F0 int16; F1 uint8; F2 uint8 }
type GenS7 struct { F0 int32; F1 uint8; F2 uint8 }
type GenS8 struct { F0 int8; F1 string; F2 uint64 }
type GenS9 struct { F0 int64; F1 uint16 }
type GenS10 struct { F0 int8; F1 int32; F2 uint8 }
type GenS11 struct { F0 uint64; F1 int8 }
type GenS12 struct { F0 uint; F1 uint32 }
type GenS13 struct { F0 uint16 }
type GenS14 struct { F0 int; F1 uint32 }
type GenS15 struct { F0 uint64 }
type GenS16 struct { F0 int16 }
type GenS17 struct { F0 int }
type GenS18 struct { F0 int32 }
type GenS19 struct { F0 string; F1 int; F2 int16 }
type GenS20 struct { F0 int64; F1 int64 }
type GenS21 struct { F0 uint64 }
func genFn0() uint16 {
    v4 := GenS1{F0: 244, F1: 1}
    var v5 uint16 = 32767
    v5 = v5
    var v6 int8 = -128
    _ = v4
    _ = v5
    _ = v6
    return 65535
}

func main() {
    var v1 bool = false
    var v2 int8 = 56
    var v3 float64 = float64(0.1)
    v7 := []int{0, 431}
    switch "a" {
    case "hello":
        if (!(float32((float64(1e10) / v3)) >= -2.25)) {
            for fi8 := 0; fi8 < 11; fi8++ {
                if v1 {
                    fmt.Println(v3, v3)
                } else {
                    var v9 float32 = float32(0.1)
                    for fi10 := 0; fi10 < 15; fi10++ {
                        var v11 uint64 = 2
                        switch (((("foo bar" + "hello") + (("hello" + "\n\t") + ("hello" + ("" + "")))) + "hello") + "") {
                        case "\n\t":
                            var v12 uint = 18446744073709551615
                            _ = v12
                        case "a":
                            switch "\n\t" {
                            case "\n\t":
                                var v13 float32 = v9
                                fmt.Println(v13, v11)
                                _ = v13
                            case "hello":
                                v3 = v3
                            default:
                                switch (-(fi10)) {
                                case 9223372036854775807:
                                    v9 += 0.1
                                case -1:
                                    fmt.Println(v9, v9)
                                case 1:
                                    switch ("foo bar" + (("\n\t" + ("foo bar" + ((("\n\t" + "\n\t") + ("" + "foo bar")) + "a"))) + "foo bar")) {
                                    case "a":
                                        v9 -= (0.1 * ((0.1 / v9) * (v9 + v9)))
                                    case "":
                                        v14 := GenS1{F0: 0, F1: 4294967295}
                                        _ = v14
                                    }
                                }
                            }
                        default:
                            v15 := GenS2{F0: 4294967295, F1: (v2 << ((1) & 63)), F2: -1}
                            _ = v15
                        }
                        _ = fi10
                        _ = v11
                    }
                    fmt.Println(v1, v1, v1)
                    _ = v9
                }
                if (-1 < 1) {
                    for fi16 := 0; fi16 < 8; fi16++ {
                        if (509 <= -853) {
                            v17 := GenS1{F0: 921, F1: 654}
                            v2 -= v2
                            v1 = (0 != 18446744073709551615)
                            var v18 uint64 = 42
                            v18 += 0
                            _ = v17
                            _ = v18
                        } else {
                            switch fi16 {
                            case 127:
                                var v19 int16 = 1
                                _ = v19
                            case -2147483648:
                                if ((2147483647 < 42) || (v1 && v1)) {
                                    if (32767 < -878) {
                                        v2 += (v2 ^ (v2 >> ((18446744073709551615) & 63)))
                                    }
                                    v20 := GenS3{F0: v1, F1: 435, F2: (v2 >> ((65535) & 63))}
                                    _ = v20
                                } else {
                                    switch 128 {
                                    case 128:
                                        v21 := GenS1{F0: 18446744073709551615, F1: 0}
                                        v1 = v1
                                        v2 = 100
                                        _ = v21
                                    case 751:
                                        v2 += v2
                                    case 32767:
                                        v3 = (v3 / v3)
                                    }
                                    if (((!(2 >= 389)) && v1) || (2 > 42)) {
                                        v2 += v2
                                        if (1 >= 127) {
                                            if len(v7) >= 2 {
                                                v7 = v7[1:len(v7)]
                                            }
                                        } else {
                                            if len(v7) >= 2 {
                                                v7 = v7[1:len(v7)]
                                            }
                                        }
                                        if ((fi16 * fi16) != (fi16 << ((2) & 63))) {
                                            v2 -= (-(v2))
                                        }
                                        if v1 {
                                            var v22 int8 = -1
                                            _ = v22
                                        } else {
                                            switch "\n\t" {
                                            case "":
                                                fmt.Println(fi16, v2, fi8)
                                            case "\n\t":
                                                fmt.Println(fi16, fi8)
                                            }
                                        }
                                    } else {
                                        v1 = (0 <= 42)
                                        v2 = 42
                                    }
                                }
                            case 32767:
                                v2 = (v2 | v2)
                            default:
                                var v23 string = ("" + "foo bar")
                                _ = v23
                            }
                            v24 := GenS4{F0: 870, F1: 87}
                            _ = v24
                        }
                        if (v1 || (v1 || ((v2 * v2) >= (v2 * v2)))) {
                            var v25 uint64 = 845
                            _ = v25
                        }
                        _ = fi16
                    }
                    var v26 int64 = 990
                    var v27 int32 = 1
                    _ = v26
                    _ = v27
                } else {
                    fmt.Println(fi8)
                    switch fi8 {
                    case 4294967295:
                        for fi28 := 0; fi28 < 19; fi28++ {
                            fmt.Println(v1)
                            _ = fi28
                        }
                    default:
                        if ("foo bar" > "") {
                            if (v3 != ((((v3 + (((v3 / (v3 + (v3 + (float64(((((-2.25 + 1e10) - float32((0.0 + -2.25))) * 1e10) + (float32(((v3 - 1e10) * -2.25)) + (3.14159 * ((1.5 * (0.1 + 3.14159)) + float32(float64(3.14159))))))) + ((v3 + (v3 * 3.14159)) * v3))))) / v3) + (1.5 - v3))) - -2.25) - (float64(3.14159) * v3)) + float64(float32(v3)))) {
                                var v29 uint16 = 65535
                                var v30 string = ("a" + "hello")
                                fmt.Println(v1, v30, v30)
                                _ = v29
                                _ = v30
                            } else {
                                if (0 < 2147483648) {
                                    v7 = append(v7, fi8)
                                    if v1 {
                                        var v31 uint = 0
                                        _ = v31
                                    } else {
                                        fmt.Println(v2)
                                        if (173 <= 42) {
                                            if (166 == 42) {
                                                if len(v7) > 0 {
                                                    v7[0] = fi8
                                                }
                                                fmt.Println(v3)
                                            } else {
                                                if len(v7) >= 2 {
                                                    v7 = v7[1:len(v7)]
                                                }
                                            }
                                        }
                                    }
                                    v1 = false
                                }
                            }
                            switch ("foo bar" + "a") {
                            case "":
                                v32 := GenS5{F0: 0, F1: -32768}
                                _ = v32
                            case "hello":
                                v1 = ((112 < 42) || (166 != 255))
                            case "a":
                                var v33 float64 = v3
                                _ = v33
                            }
                            if v1 {
                                switch fi8 {
                                case -15:
                                    if (!v1) {
                                        var v34 uint16 = 32767
                                        _ = v34
                                    } else {
                                        for fi35 := 0; fi35 < 9; fi35++ {
                                            if (fi8 > -239) {
                                                if len(v7) > 0 {
                                                    v7[0] = -2147483648
                                                }
                                                if v1 {
                                                    v36 := GenS6{F0: 521, F1: 238, F2: 0}
                                                    switch 128 {
                                                    case 4294967295:
                                                        switch 42 {
                                                        case -1:
                                                            v2 = (v2 / (((v2 + -128)) | 1))
                                                        case 9223372036854775807:
                                                            var v37 int = (fi8 >> ((255) & 63))
                                                            _ = v37
                                                        case 128:
                                                            v38 := GenS3{F0: v1, F1: -332, F2: v2}
                                                            _ = v38
                                                        }
                                                    case -967:
                                                        fmt.Println(v2, fi8)
                                                    case 128:
                                                        var v39 int64 = 42
                                                        if ((fi8 <= 0) || (v1 && v1)) {
                                                            if (!(!false)) {
                                                                var v40 int32 = 1
                                                                _ = v40
                                                            } else {
                                                                if ((^(v39)) < (v39 ^ v39)) {
                                                                    var v41 float64 = v3
                                                                    _ = v41
                                                                }
                                                                if len(v7) > 0 {
                                                                    v7[0] = (fi8 >> ((255) & 63))
                                                                }
                                                                if v1 {
                                                                    v42 := GenS6{F0: -146, F1: 42, F2: 127}
                                                                    if (416 != 1) {
                                                                        v39 += (v39 ^ -926)
                                                                        var v43 int64 = v39
                                                                        _ = v43
                                                                    } else {
                                                                        fmt.Println(fi8)
                                                                    }
                                                                    v44 := GenS1{F0: 42, F1: 1}
                                                                    _ = v42
                                                                    _ = v44
                                                                } else {
                                                                    var v45 bool = (fi8 > (fi8 / (((fi8 >> ((91) & 63))) | 1)))
                                                                    var v46 int8 = v2
                                                                    v46 -= v2
                                                                    _ = v45
                                                                    _ = v46
                                                                }
                                                                fmt.Println(v2, fi35)
                                                                var v47 uint16 = 301
                                                                var v48 uint16 = v47
                                                                v49 := GenS7{F0: 2147483647, F1: 0, F2: 128}
                                                                _ = v47
                                                                _ = v48
                                                                _ = v49
                                                            }
                                                        } else {
                                                            fmt.Println(v3, v3)
                                                            if v1 {
                                                                if v1 {
                                                                    v50 := GenS4{F0: 2, F1: 0}
                                                                    if v1 {
                                                                        v7 = append(v7, (fi8 | (fi8 << ((255) & 63))))
                                                                        v51 := GenS8{F0: v2, F1: ("hello" + "\n\t"), F2: 512}
                                                                        _ = v51
                                                                    } else {
                                                                        _ = len(v7)
                                                                        v52 := GenS1{F0: 42, F1: 0}
                                                                        _ = v52
                                                                    }
                                                                    fmt.Println(fi8, v3, v1)
                                                                    _ = v50
                                                                }
                                                                switch 9223372036854775807 {
                                                                case 32767:
                                                                    v53 := GenS9{F0: (-(v39)), F1: 541}
                                                                    _ = v53
                                                                default:
                                                                    if len(v7) > 0 {
                                                                        v7[0] = fi35
                                                                    }
                                                                }
                                                            } else {
                                                                var v54 string = "foo bar"
                                                                v2 += v2
                                                                if v1 {
                                                                    var v55 uint = 2
                                                                    switch "\n\t" {
                                                                    case "":
                                                                        v2 = (v2 - v2)
                                                                    case "hello":
                                                                        if (v2 != (v2 * v2)) {
                                                                            v54 = v54
                                                                        }
                                                                    case "foo bar":
                                                                        var v56 uint32 = 2147483648
                                                                        _ = v56
                                                                    }
                                                                    _ = v55
                                                                }
                                                                _ = v54
                                                            }
                                                        }
                                                        _ = v39
                                                    default:
                                                        v1 = (false && (!(v1 || true)))
                                                    }
                                                    var v57 int = 42
                                                    _ = v36
                                                    _ = v57
                                                }
                                                switch ((((("foo bar" + "a") + "a") + ("hello" + (("" + "hello") + "\n\t"))) + ("a" + ("hello" + "hello"))) + "hello") {
                                                case "a":
                                                    if ("a" > ((("a" + "a") + "a") + "\n\t")) {
                                                        switch (fi8 - 2147483647) {
                                                        case 127:
                                                            _ = len(v7)
                                                        }
                                                    } else {
                                                        var v58 string = "a"
                                                        _ = v58
                                                    }
                                                case "\n\t":
                                                    v1 = false
                                                case "foo bar":
                                                    v59 := GenS3{F0: (v1 || (fi8 <= (fi8 - 9223372036854775807))), F1: 1, F2: 41}
                                                    _ = v59
                                                }
                                            } else {
                                                fmt.Println(fi8)
                                            }
                                            v7 = append(v7, (fi8 << ((676) & 63)))
                                            _ = fi35
                                        }
                                        fmt.Println(v1, fi8, v3)
                                        _ = len(v7)
                                    }
                                case 622:
                                    var v60 int8 = v2
                                    _ = v60
                                case -9223372036854775808:
                                    _ = len(v7)
                                default:
                                    var v61 uint = 0
                                    _ = v61
                                }
                            } else {
                                var v62 float32 = 1.5
                                if (!v1) {
                                    v63 := GenS10{F0: v2, F1: 2147483647, F2: 1}
                                    switch -9223372036854775808 {
                                    case -2147483648:
                                        for fi64 := 0; fi64 < 10; fi64++ {
                                            if (v1 && v1) {
                                                if (-1 != 1) {
                                                    v65 := GenS2{F0: 2, F1: 127, F2: 56}
                                                    var v66 int8 = v2
                                                    if ((348 >= 4294967295) && false) {
                                                        v7 = append(v7, fi8)
                                                    } else {
                                                        v1 = (1 <= -845)
                                                        var v67 uint16 = 463
                                                        var v68 uint8 = 1
                                                        v1 = (((("a" + "\n\t") >= ("a" + "\n\t")) && v1) || (42 <= -32768))
                                                        v3 = -2.25
                                                        _ = v67
                                                        _ = v68
                                                    }
                                                    _ = v65
                                                    _ = v66
                                                } else {
                                                    var v69 float64 = (-2.25 * float64(v62))
                                                    if (v2 > 0) {
                                                        if (v1 && v1) {
                                                            var v70 bool = v1
                                                            if (!v70) {
                                                                fmt.Println(v69, v69)
                                                            }
                                                            if (v1 && (433 <= 42)) {
                                                                v7 = append(v7, -482)
                                                                v2 += (-(v2))
                                                            }
                                                            var v71 string = ""
                                                            _ = v70
                                                            _ = v71
                                                        }
                                                        if len(v7) >= 2 {
                                                            v7 = v7[1:len(v7)]
                                                        }
                                                        switch "" {
                                                        case "":
                                                            v1 = (!(0 < 4294967295))
                                                        case "hello":
                                                            if (-720 < 42) {
                                                                if true {
                                                                    switch "hello" {
                                                                    case "hello":
                                                                        var v72 uint16 = genFn0()
                                                                        _ = v72
                                                                    case "a":
                                                                        if len(v7) > 0 {
                                                                            v7[0] = (fi8 << ((65535) & 63))
                                                                        }
                                                                    default:
                                                                        if false {
                                                                            fmt.Println(v1, v3, fi64)
                                                                            if false {
                                                                                if len(v7) > 0 {
                                                                                    v7[0] = -855
                                                                                }
                                                                            }
                                                                        }
                                                                    }
                                                                    var v73 string = "a"
                                                                    _ = v73
                                                                } else {
                                                                    v7 = append(v7, (^(fi8)))
                                                                    fmt.Println(v1, v1, v69)
                                                                }
                                                                if ((v1 || v1) && (255 < 0)) {
                                                                    switch "foo bar" {
                                                                    case "\n\t":
                                                                        v3 += 0.1
                                                                    case "foo bar":
                                                                        var v74 uint8 = 251
                                                                        _ = v74
                                                                    case "a":
                                                                        var v75 uint16 = 0
                                                                        _ = v75
                                                                    default:
                                                                        var v76 bool = (!(391 <= 366))
                                                                        _ = v76
                                                                    }
                                                                    fmt.Println(v1)
                                                                }
                                                                var v77 uint8 = 127
                                                                _ = v77
                                                            } else {
                                                                v69 += (v3 * -2.25)
                                                            }
                                                            v7 = append(v7, -644)
                                                        }
                                                    }
                                                    _ = v69
                                                }
                                                v1 = (255 < 42)
                                                v78 := GenS11{F0: 2, F1: -1}
                                                if len(v7) > 0 {
                                                    v7[0] = fi8
                                                }
                                                var v79 uint64 = 42
                                                v62 -= (v62 / v62)
                                                fmt.Println(v1)
                                                if len(v7) > 0 {
                                                    v7[0] = fi8
                                                }
                                                v7 = append(v7, fi64)
                                                _ = v78
                                                _ = v79
                                            } else {
                                                switch fi8 {
                                                case 32768:
                                                    var v80 uint = 601
                                                    _ = v80
                                                }
                                                v1 = ("a" > ("a" + "hello"))
                                                var v81 uint64 = 42
                                                _ = v81
                                            }
                                            _ = fi64
                                        }
                                    case 653:
                                        if len(v7) > 0 {
                                            v7[0] = 42
                                        }
                                    }
                                    _ = v63
                                }
                                _ = v62
                            }
                        } else {
                            if len(v7) > 0 {
                                v7[0] = (^(fi8))
                            }
                            for fi82 := 0; fi82 < 2; fi82++ {
                                switch "a" {
                                case "hello":
                                    if len(v7) > 0 {
                                        v7[0] = (fi8 + fi82)
                                    }
                                case "foo bar":
                                    var v83 float64 = ((v3 / v3) - v3)
                                    _ = v83
                                case "a":
                                    v1 = (v1 && false)
                                default:
                                    fmt.Println(v2)
                                }
                                if len(v7) >= 2 {
                                    v7 = v7[1:len(v7)]
                                }
                                v7 = append(v7, 128)
                                _ = fi82
                            }
                        }
                    }
                    v84 := GenS12{F0: 284, F1: 0}
                    _ = v84
                }
                _ = fi8
            }
            var v85 uint = 845
            _ = v85
        } else {
            fmt.Println(v3)
            switch ((("foo bar" + "a") + ("" + "foo bar")) + ("hello" + ((("a" + "a") + "a") + (("foo bar" + "") + "")))) {
            case "\n\t":
                if (42 > 118) {
                    v86 := GenS3{F0: (0 <= -1), F1: -823, F2: 1}
                    v1 = false
                    _ = v86
                } else {
                    var v87 uint32 = 65535
                    _ = v87
                }
            case "":
                for fi88 := 0; fi88 < 4; fi88++ {
                    v7 = append(v7, 187)
                    _ = fi88
                }
            case "a":
                v89 := GenS2{F0: 2147483648, F1: (v2 / ((v2) | 1)), F2: 950}
                _ = v89
            }
            var v90 int8 = (v2 * v2)
            _ = v90
        }
    case "foo bar":
        v3 = v3
    case "":
        switch -9223372036854775808 {
        case 153:
            fmt.Println(v1)
        default:
            for fi91 := 0; fi91 < 14; fi91++ {
                if (2 < 0) {
                    for fi92 := 0; fi92 < 1; fi92++ {
                        v2 = (v2 << ((18446744073709551615) & 63))
                        switch ("" + "hello") {
                        case "a":
                            if (genFn0() < 42) {
                                v7 = append(v7, int(v2))
                                v93 := GenS13{F0: 2}
                                if true {
                                    v3 += 0.1
                                    switch "a" {
                                    case "":
                                        fmt.Println(fi92, v1, v2)
                                    case "\n\t":
                                        if true {
                                            fmt.Println(v3)
                                            v7 = append(v7, 717)
                                            fmt.Println(fi92)
                                        }
                                    default:
                                        if len(v7) > 0 {
                                            v7[0] = fi92
                                        }
                                    }
                                    var v94 uint64 = 4294967295
                                    _ = v94
                                } else {
                                    v1 = v1
                                    var v95 uint32 = 2
                                    v2 += v2
                                    _ = v95
                                }
                                _ = v93
                            }
                        case "":
                            v3 += v3
                        case "\n\t":
                            v96 := GenS12{F0: 352, F1: 0}
                            _ = v96
                        }
                        v97 := GenS12{F0: 0, F1: 540}
                        _ = fi92
                        _ = v97
                    }
                    if v1 {
                        switch "foo bar" {
                        case "\n\t":
                            switch "hello" {
                            case "\n\t":
                                switch fi91 {
                                case -2147483648:
                                    v7 = append(v7, -9223372036854775808)
                                }
                            case "foo bar":
                                var v98 uint32 = 857
                                _ = v98
                            case "":
                                fmt.Println(v3, v1, fi91)
                            }
                        default:
                            var v99 int8 = v2
                            _ = v99
                        }
                        var v100 int8 = (v2 << ((2) & 63))
                        _ = v100
                    } else {
                        v101 := GenS14{F0: 0, F1: 2}
                        var v102 int = fi91
                        _ = v101
                        _ = v102
                    }
                } else {
                    v103 := GenS14{F0: (fi91 ^ fi91), F1: 4294967295}
                    if true {
                        var v104 float64 = v3
                        if v1 {
                            switch int(v2) {
                            case -289:
                                for fi105 := 0; fi105 < 12; fi105++ {
                                    v7 = append(v7, 987)
                                    _ = fi105
                                }
                            case -9223372036854775808:
                                var v106 int64 = 42
                                _ = v106
                            default:
                                for fi107 := 0; fi107 < 3; fi107++ {
                                    var v108 int8 = v2
                                    switch fi91 {
                                    case 715:
                                        switch fi107 {
                                        case 32767:
                                            var v109 int8 = (v108 ^ v2)
                                            _ = v109
                                        case -9223372036854775808:
                                            v1 = (2 != 49)
                                        }
                                    case 606:
                                        v104 += v3
                                    }
                                    v1 = v1
                                    var v110 uint8 = 1
                                    _ = fi107
                                    _ = v108
                                    _ = v110
                                }
                            }
                        } else {
                            switch ("\n\t" + ((("a" + "\n\t") + "foo bar") + "")) {
                            case "":
                                if len(v7) >= 2 {
                                    v7 = v7[1:len(v7)]
                                }
                            case "foo bar":
                                fmt.Println(v104, v104)
                            case "a":
                                fmt.Println(v3, v104)
                            }
                            v3 -= (-2.25 - v3)
                            if len(v7) >= 2 {
                                v7 = v7[1:len(v7)]
                            }
                        }
                        for fi111 := 0; fi111 < 16; fi111++ {
                            fmt.Println(fi91, v3)
                            if v1 {
                                if len(v7) >= 2 {
                                    v7 = v7[1:len(v7)]
                                }
                                if ("\n\t" > "a") {
                                    var v112 int16 = 42
                                    _ = v112
                                } else {
                                    var v113 int16 = 28
                                    var v114 int64 = 0
                                    v114 += v114
                                    _ = v113
                                    _ = v114
                                }
                                v115 := GenS15{F0: 123}
                                _ = v115
                            }
                            fmt.Println(v1)
                            _ = fi111
                        }
                        _ = v104
                    } else {
                        for fi116 := 0; fi116 < 8; fi116++ {
                            fmt.Println(fi91)
                            v1 = v1
                            _ = fi116
                        }
                        var v117 int16 = 1000
                        _ = v117
                    }
                    _ = v103
                }
                v118 := GenS2{F0: 159, F1: -41, F2: -32768}
                v3 = (v3 + v3)
                _ = fi91
                _ = v118
            }
        }
    }
    fmt.Println(v1, v3)
    v7 = append(v7, 42)
    if true {
        v2 += v2
    } else {
        var v119 float64 = v3
        _ = v119
    }
    if true {
        switch "foo bar" {
        case "hello":
            v3 = 1e10
        case "":
            fmt.Println(v3, v3, v1)
        case "foo bar":
            var v120 uint64 = 459
            _ = v120
        default:
            for fi121 := 0; fi121 < 17; fi121++ {
                fmt.Println(v1, fi121)
                var v122 int16 = -374
                _ = fi121
                _ = v122
            }
        }
        for fi123 := 0; fi123 < 12; fi123++ {
            for fi124 := 0; fi124 < 1; fi124++ {
                if false {
                    var v125 int = -367
                    var v126 uint16 = 65535
                    if (421 <= 4294967295) {
                        var v127 int8 = v2
                        v128 := GenS15{F0: 42}
                        v2 -= v127
                        if ((v2 % (((^(v2))) | 1)) != v2) {
                            if v1 {
                                if false {
                                    v125 -= int(v2)
                                    var v129 string = "foo bar"
                                    var v130 float32 = 1.5
                                    _ = v129
                                    _ = v130
                                } else {
                                    if (true || v1) {
                                        var v131 int32 = 1
                                        fmt.Println(v126)
                                        var v132 int16 = -594
                                        switch "foo bar" {
                                        case "hello":
                                            var v133 int32 = 767
                                            _ = v133
                                        case "\n\t":
                                            _ = len(v7)
                                        case "foo bar":
                                            fmt.Println(v132, v127)
                                            v7 = append(v7, int(v132))
                                        default:
                                            if (0 != v132) {
                                                v134 := GenS16{F0: v132}
                                                if (v1 || (((v126 - 65535) <= v126) && v1)) {
                                                    if false {
                                                        if ((!(!v1)) || ((-309 != -9223372036854775808) && v1)) {
                                                            v1 = (int8(v132) != (v127 & v2))
                                                            v1 = ((1 > 432) && (int8(v132) < (-(v127))))
                                                            v132 = 973
                                                        } else {
                                                            if (false && (18446744073709551615 >= 4294967295)) {
                                                                v135 := GenS17{F0: fi124}
                                                                _ = v135
                                                            } else {
                                                                _ = len(v7)
                                                                if len(v7) >= 2 {
                                                                    v7 = v7[1:len(v7)]
                                                                }
                                                                var v136 int32 = v131
                                                                v131 = v131
                                                                if len(v7) > 0 {
                                                                    v7[0] = fi123
                                                                }
                                                                _ = v136
                                                            }
                                                        }
                                                        v126 = (v126 ^ genFn0())
                                                        v137 := GenS18{F0: v131}
                                                        _ = v137
                                                    }
                                                    if len(v7) >= 2 {
                                                        v7 = v7[1:len(v7)]
                                                    }
                                                    switch (((("\n\t" + ("" + "")) + ("foo bar" + ("hello" + ""))) + "hello") + "hello") {
                                                    case "\n\t":
                                                        v138 := GenS7{F0: (v131 + v131), F1: 1, F2: 204}
                                                        _ = v138
                                                    case "foo bar":
                                                        var v139 uint8 = 42
                                                        _ = v139
                                                    default:
                                                        v126 -= (^(v126))
                                                    }
                                                    var v140 uint = 989
                                                    v141 := GenS6{F0: (v132 | v132), F1: 1, F2: 1}
                                                    var v142 string = ("hello" + "a")
                                                    v3 = v3
                                                    _ = v140
                                                    _ = v141
                                                    _ = v142
                                                }
                                                v127 -= v127
                                                var v143 uint16 = (v126 & (v126 + 852))
                                                _ = v134
                                                _ = v143
                                            } else {
                                                if (v1 || false) {
                                                    var v144 uint64 = 976
                                                    _ = v144
                                                } else {
                                                    if v1 {
                                                        v2 += (v127 % (((v2 | (v2 * (v2 >> ((253) & 63))))) | 1))
                                                        if false {
                                                            v145 := GenS13{F0: genFn0()}
                                                            if (!true) {
                                                                var v146 uint16 = 2
                                                                var v147 int16 = (^(v132))
                                                                v126 -= v126
                                                                _ = v146
                                                                _ = v147
                                                            }
                                                            v131 -= (^(v131))
                                                            _ = v145
                                                        }
                                                        if (v125 != -225) {
                                                            if (0.1 != 3.14159) {
                                                                var v148 uint64 = 2
                                                                _ = v148
                                                            }
                                                            if v1 {
                                                                v149 := GenS6{F0: (v132 + (-(v132))), F1: 127, F2: 0}
                                                                _ = v149
                                                            } else {
                                                                var v150 int32 = (v131 * v131)
                                                                _ = v150
                                                            }
                                                            if (!(v1 && (1.5 <= 1e10))) {
                                                                v125 = fi123
                                                            }
                                                        }
                                                    }
                                                    switch fi124 {
                                                    case 287:
                                                        if v1 {
                                                            v151 := GenS18{F0: v131}
                                                            v152 := GenS1{F0: 37, F1: 0}
                                                            switch -218 {
                                                            case 1:
                                                                fmt.Println(fi124)
                                                            case -1:
                                                                fmt.Println(v132, v132, v126)
                                                                var v153 uint16 = 42
                                                                v1 = false
                                                                _ = v153
                                                            default:
                                                                if false {
                                                                    var v154 int64 = -9223372036854775808
                                                                    fmt.Println(v154)
                                                                    if ((-2.25 / ((0.0 / (3.14159 / 1.0)) - (0.1 + ((((0.0 + ((0.0 - float32(0.1)) / (1e10 * float32((0.0 + v3))))) * 0.0) * float32(v3)) + -2.25)))) > (0.1 * (3.14159 * 1e10))) {
                                                                        var v155 int8 = v127
                                                                        fmt.Println(v131)
                                                                        _ = v155
                                                                    } else {
                                                                        switch fi123 {
                                                                        case -813:
                                                                            var v156 uint16 = (v126 | v126)
                                                                            _ = v156
                                                                        case -59:
                                                                            var v157 uint8 = 121
                                                                            _ = v157
                                                                        default:
                                                                            v158 := GenS3{F0: v1, F1: -32768, F2: (v127 >> ((649) & 63))}
                                                                            _ = v158
                                                                        }
                                                                        v125 += (fi124 * fi123)
                                                                        if v1 {
                                                                            var v159 string = "\n\t"
                                                                            _ = v159
                                                                        } else {
                                                                            if v1 {
                                                                                switch fi124 {
                                                                                case 42:
                                                                                    var v160 float32 = 0.0
                                                                                    _ = v160
                                                                                default:
                                                                                    if (4294967295 >= 42) {
                                                                                        var v161 string = "foo bar"
                                                                                        v162 := GenS14{F0: (fi124 & -704), F1: 2}
                                                                                        var v163 int64 = (v154 & int64(v131))
                                                                                        v7 = append(v7, fi124)
                                                                                        switch fi124 {
                                                                                        case -739:
                                                                                            v3 += (0.1 - v3)
                                                                                        case -2147483648:
                                                                                            var v164 float32 = float32(v3)
                                                                                            _ = v164
                                                                                        default:
                                                                                            if v1 {
                                                                                                v165 := GenS7{F0: v131, F1: 196, F2: 0}
                                                                                                _ = v165
                                                                                            }
                                                                                        }
                                                                                        var v166 int16 = 1000
                                                                                        fmt.Println(v3, v127, v161)
                                                                                        _ = v161
                                                                                        _ = v162
                                                                                        _ = v163
                                                                                        _ = v166
                                                                                    } else {
                                                                                        switch v125 {
                                                                                        case 1:
                                                                                            fmt.Println(v132, v131, fi124)
                                                                                        case 127:
                                                                                            var v167 float32 = (0.0 * 3.14159)
                                                                                            _ = v167
                                                                                        }
                                                                                        v7 = append(v7, (-(fi124)))
                                                                                        if v1 {
                                                                                            if (!((v3 / 0.1) != v3)) {
                                                                                                v7 = append(v7, (fi124 / ((fi123) | 1)))
                                                                                                if ((^(v154)) <= (v154 / ((v154) | 1))) {
                                                                                                    v168 := GenS19{F0: "", F1: (^(v125)), F2: 1}
                                                                                                    _ = v168
                                                                                                } else {
                                                                                                    var v169 uint = 42
                                                                                                    _ = v169
                                                                                                }
                                                                                                v170 := GenS4{F0: 513, F1: 127}
                                                                                                _ = v170
                                                                                            }
                                                                                            var v171 float64 = v3
                                                                                            fmt.Println(v131)
                                                                                            var v172 float64 = float64((1e10 - (float32(v3) / float32((float64(float32(0.1)) + v3)))))
                                                                                            _ = v171
                                                                                            _ = v172
                                                                                        }
                                                                                        if len(v7) > 0 {
                                                                                            v7[0] = fi124
                                                                                        }
                                                                                    }
                                                                                }
                                                                            }
                                                                            var v173 uint64 = 18446744073709551615
                                                                            _ = v173
                                                                        }
                                                                        var v174 uint = 537
                                                                        _ = v174
                                                                    }
                                                                    fmt.Println(v154, v154, v132)
                                                                    _ = v154
                                                                }
                                                            }
                                                            _ = v151
                                                            _ = v152
                                                        } else {
                                                            if (379 == int64(v126)) {
                                                                var v175 uint64 = 4294967295
                                                                _ = v175
                                                            }
                                                            v176 := GenS8{F0: (v127 % (((v2 * -1)) | 1)), F1: "hello", F2: 420}
                                                            v3 = 0.1
                                                            _ = v176
                                                        }
                                                    default:
                                                        v7 = append(v7, 127)
                                                    }
                                                    if (!(genFn0() == (v126 & (v126 & v126)))) {
                                                        fmt.Println(v3, v126)
                                                        v2 = 82
                                                        if len(v7) >= 2 {
                                                            v7 = v7[1:len(v7)]
                                                        }
                                                    }
                                                }
                                            }
                                        }
                                        var v177 float64 = 1.5
                                        _ = v131
                                        _ = v132
                                        _ = v177
                                    }
                                    _ = len(v7)
                                    if v1 {
                                        if v1 {
                                            fmt.Println(v127, v1, v126)
                                            var v178 string = ("foo bar" + "")
                                            _ = v178
                                        }
                                    } else {
                                        var v179 uint8 = 0
                                        if (1 < -650) {
                                            v7 = append(v7, fi124)
                                            if (((fi123 ^ v125) != (v125 >> ((1) & 63))) || (v179 == (^(v179)))) {
                                                fmt.Println(v1, v125, v125)
                                                v1 = true
                                            } else {
                                                var v180 int32 = -893
                                                _ = v180
                                            }
                                        } else {
                                            var v181 uint8 = v179
                                            _ = v181
                                        }
                                        _ = v179
                                    }
                                }
                                v182 := GenS20{F0: -420, F1: 0}
                                v126 = v126
                                _ = v182
                            } else {
                                v3 = v3
                                var v183 float32 = 0.0
                                switch (("" + "") + ((("\n\t" + "a") + "\n\t") + "\n\t")) {
                                case "a":
                                    var v184 string = (("\n\t" + "a") + "hello")
                                    _ = v184
                                case "":
                                    v185 := GenS2{F0: 2147483648, F1: -128, F2: 1000}
                                    _ = v185
                                default:
                                    var v186 int64 = 9223372036854775807
                                    fmt.Println(v1, v1)
                                    _ = v186
                                }
                                _ = v183
                            }
                            var v187 int32 = 42
                            _ = v187
                        } else {
                            v127 -= (v2 / ((v2) | 1))
                            switch (fi124 | (fi124 * (v125 / (((v125 + (^(fi123)))) | 1)))) {
                            case -2147483648:
                                v1 = (((float32(3.14159) + 1.5) <= 0.1) && false)
                            case -1:
                                switch int(v126) {
                                case 9223372036854775807:
                                    if false {
                                        v7 = append(v7, fi123)
                                        var v188 float32 = 1.5
                                        if ((32768 > fi123) && false) {
                                            var v189 uint8 = 128
                                            v126 += v126
                                            _ = v189
                                        }
                                        _ = v188
                                    }
                                case -1:
                                    v2 = (v127 / ((35) | 1))
                                }
                            case -9223372036854775808:
                                var v190 int32 = 2147483647
                                _ = v190
                            }
                        }
                        _ = v127
                        _ = v128
                    } else {
                        switch "hello" {
                        case "a":
                            var v191 uint8 = 42
                            _ = v191
                        }
                        _ = len(v7)
                    }
                    v192 := GenS3{F0: true, F1: -330, F2: 127}
                    v193 := GenS21{F0: 666}
                    _ = v125
                    _ = v126
                    _ = v192
                    _ = v193
                } else {
                    v7 = append(v7, fi123)
                    v1 = v1
                }
                _ = fi124
            }
            _ = fi123
        }
        var v194 uint32 = 1
        _ = v194
    }
    switch -9223372036854775808 {
    case 296:
        v195 := GenS18{F0: 0}
        _ = v195
    case 0:
        v196 := GenS12{F0: 0, F1: 2}
        _ = v196
    case 2147483647:
        fmt.Println(v3)
    default:
        v197 := GenS5{F0: 0, F1: 1000}
        _ = v197
    }
    v1 = false
    var v198 string = ("foo bar" + ("hello" + "\n\t"))
    v2 -= v2
    v2 = v2
    v199 := GenS11{F0: 0, F1: 39}
    fmt.Println(v2, v3)
    switch 32767 {
    case 9223372036854775807:
        for fi200 := 0; fi200 < 11; fi200++ {
            _ = len(v7)
            _ = len(v7)
            _ = fi200
        }
    case 4294967295:
        var v201 int8 = v2
        _ = v201
    }
    var v202 float32 = 3.14159
    var v203 uint16 = 255
    v203 -= v203
    fmt.Println("v1:", v1)
    fmt.Println("v198:", v198)
    fmt.Println("v2:", v2)
    fmt.Println("v202:", v202)
    fmt.Println("v203:", v203)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v198
    _ = v199
    _ = v2
    _ = v202
    _ = v203
    _ = v3
    _ = v7
}
