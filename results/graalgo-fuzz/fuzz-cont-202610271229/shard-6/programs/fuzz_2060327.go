package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string; F2 string }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint16; F1 uint16; F2 int }
type GenS5 struct { F0 int; F1 int16 }
type GenS6 struct { F0 uint32; F1 uint16; F2 uint32 }
type GenS7 struct { F0 uint8; F1 int32 }
type GenS8 struct { F0 int16; F1 uint64 }
type GenS9 struct { F0 int64; F1 uint32 }
type GenS10 struct { F0 uint }
type GenS11 struct { F0 int; F1 uint16; F2 bool }
type GenS12 struct { F0 int64; F1 uint16; F2 int64 }
type GenS13 struct { F0 uint64; F1 string }
type GenS14 struct { F0 int32 }
type GenS15 struct { F0 int; F1 bool; F2 int8 }
type GenS16 struct { F0 int64; F1 int16 }
type GenS17 struct { F0 int64 }
type GenS18 struct { F0 bool }
type GenS19 struct { F0 int64; F1 int; F2 uint16 }
type GenS20 struct { F0 bool }
type GenS21 struct { F0 int; F1 uint64; F2 int8 }
type GenS22 struct { F0 uint32; F1 string; F2 bool }
type GenS23 struct { F0 string }
type GenS24 struct { F0 uint16; F1 int16; F2 string }
func genFn0(p0 int16, p1 int64) uint16 {
    for fi5 := 0; fi5 < 11; fi5++ {
        v6 := GenS1{F0: 598, F1: ("foo bar" + ("" + ((("" + "hello") + "\n\t") + "hello"))), F2: ("\n\t" + "")}
        v7 := []int{(fi5 | 9223372036854775807), (fi5 << ((2) & 63)), fi5, 2147483647}
        fmt.Println(p1, p1)
        _ = fi5
        _ = v6
        _ = v7
    }
    v8 := GenS2{F0: "foo bar"}
    if false {
        p1 = (p1 ^ (p1 / (((p1 - p1)) | 1)))
        v9 := GenS3{F0: (-(p0))}
        _ = v9
    }
    var v10 int64 = 42
    v10 += -570
    _ = p0
    _ = p1
    _ = v10
    _ = v8
    return 920
}

func main() {
    v1 := []uint32{2147483647}
    var v2 string = (("" + (("" + "\n\t") + "")) + ("" + "a"))
    var v3 bool = false
    var v4 float32 = (0.1 - (-2.25 - -2.25))
    var v11 int32 = -1
    if (0.1 >= 0.1) {
        for fi12 := 0; fi12 < 13; fi12++ {
            if false {
                if v3 {
                    for fi13 := 0; fi13 < 3; fi13++ {
                        switch "a" {
                        case "\n\t":
                            v2 = v2
                        case "":
                            switch v2 {
                            case "":
                                v14 := GenS4{F0: 32767, F1: 255, F2: (fi13 / (((^(fi12))) | 1))}
                                _ = v14
                            default:
                                v3 = v3
                            }
                        }
                        v3 = ((false && false) && (110 <= 326))
                        v15 := GenS4{F0: 65535, F1: 916, F2: fi12}
                        v1 = append(v1, 65535)
                        _ = fi13
                        _ = v15
                    }
                }
                v16 := GenS2{F0: "hello"}
                _ = v16
            } else {
                var v17 string = v2
                for fi18 := 0; fi18 < 7; fi18++ {
                    if v3 {
                        if len(v1) > 0 {
                            v1[0] = 240
                        }
                        switch "" {
                        case "a":
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        case "hello":
                            _ = len(v1)
                        case "":
                            var v19 float64 = 1.5
                            if v3 {
                                v1 = append(v1, 2)
                            }
                            _ = len(v1)
                            if (!v3) {
                                fmt.Println(v2, v4)
                                v20 := GenS4{F0: 32767, F1: 1, F2: fi18}
                                fmt.Println(fi18)
                                _ = v20
                            }
                            _ = v19
                        default:
                            switch (fi12 + (fi18 & 1)) {
                            case 9223372036854775807:
                                v2 = v17
                            default:
                                v2 = v2
                            }
                        }
                    } else {
                        if true {
                            fmt.Println(fi12, v11)
                            v3 = (9223372036854775807 < -9223372036854775808)
                        } else {
                            var v21 int16 = -915
                            switch (fi18 - fi12) {
                            case 4294967295:
                                switch (v17 + "") {
                                case "":
                                    v3 = v3
                                case "hello":
                                    if (v3 || v3) {
                                        fmt.Println(v17, v11, v21)
                                        v11 = int32(v21)
                                        fmt.Println(v4)
                                        var v22 float32 = 0.1
                                        v11 += (v11 % (((v11 * v11)) | 1))
                                        switch fi12 {
                                        case -2147483648:
                                            if ((v21 ^ v21) >= -1) {
                                                var v23 int16 = (^(v21))
                                                _ = v23
                                            } else {
                                                v24 := GenS5{F0: (fi18 << ((255) & 63)), F1: 1}
                                                _ = v24
                                            }
                                        default:
                                            fmt.Println(v2, fi18, v2)
                                            switch 32768 {
                                            case -2147483648:
                                                fmt.Println(v11)
                                            }
                                        }
                                        _ = v22
                                    } else {
                                        v17 = v2
                                    }
                                }
                            }
                            v3 = ((!v3) || (fi12 == fi12))
                            _ = v21
                        }
                        v25 := GenS6{F0: 65535, F1: 65535, F2: 1}
                        v3 = v3
                        _ = v25
                    }
                    switch (fi12 * fi12) {
                    case 32767:
                        v17 = "foo bar"
                    default:
                        if ((fi18 + fi12) == 0) {
                            var v26 uint32 = 2147483647
                            _ = v26
                        } else {
                            var v27 float64 = (0.0 - (((float64(v4) + (0.0 / float64(-2.25))) + ((1.5 / 1.0) + 0.0)) - ((3.14159 + 1e10) * ((0.1 + (((0.0 * float64(v4)) - 0.1) - float64((v4 * 0.0)))) * 0.1))))
                            var v28 bool = (fi12 <= 621)
                            if false {
                                if (255 >= 167) {
                                    var v29 int = fi12
                                    v30 := GenS6{F0: 534, F1: 834, F2: 1}
                                    if len(v1) > 0 {
                                        v1[0] = 65535
                                    }
                                    _ = v29
                                    _ = v30
                                }
                            } else {
                                v1 = append(v1, 209)
                                var v31 uint64 = 265
                                var v32 int8 = 48
                                _ = v31
                                _ = v32
                            }
                            _ = v27
                            _ = v28
                        }
                    }
                    _ = fi18
                }
                v2 = ((v2 + v2) + v2)
                _ = v17
            }
            if (!(v3 && (v3 || v3))) {
                v4 = (v4 / (3.14159 - ((v4 / -2.25) + v4)))
                if v3 {
                    if v3 {
                        _ = len(v1)
                        fmt.Println(v4, v11, fi12)
                        v33 := GenS4{F0: 255, F1: 1, F2: fi12}
                        _ = v33
                    } else {
                        v1 = append(v1, 65535)
                        fmt.Println(v4, v2)
                        v4 = v4
                    }
                    switch (fi12 / (((fi12 / (((fi12 + fi12)) | 1))) | 1)) {
                    case 4294967295:
                        var v34 uint = 2
                        _ = v34
                    }
                } else {
                    var v35 float32 = ((1e10 * (v4 / (3.14159 - (v4 / (v4 * (v4 * 0.0)))))) / v4)
                    if ((0 != 2) || true) {
                        if v3 {
                            v36 := GenS4{F0: 165, F1: 32767, F2: (fi12 * fi12)}
                            v2 = v2
                            _ = v36
                        }
                        v37 := GenS2{F0: "\n\t"}
                        if (v3 || (4294967295 != 603)) {
                            for fi38 := 0; fi38 < 19; fi38++ {
                                if len(v1) > 0 {
                                    v1[0] = 42
                                }
                                var v39 float32 = v35
                                v4 = (float32(0.1) + -2.25)
                                switch (v2 + ("hello" + "a")) {
                                case "foo bar":
                                    var v40 int64 = 9223372036854775807
                                    _ = v40
                                case "\n\t":
                                    if (111 > -28) {
                                        fmt.Println(v3)
                                        var v41 int = fi12
                                        var v42 uint16 = 0
                                        _ = v41
                                        _ = v42
                                    } else {
                                        switch fi38 {
                                        case 0:
                                            if (65535 != 2) {
                                                switch fi38 {
                                                case 127:
                                                    v2 = v2
                                                case 32767:
                                                    var v43 uint = 438
                                                    _ = v43
                                                case 9223372036854775807:
                                                    switch "a" {
                                                    case "hello":
                                                        switch (-(fi38)) {
                                                        case -2147483648:
                                                            var v44 int8 = 42
                                                            _ = v44
                                                        case 2147483647:
                                                            var v45 bool = false
                                                            _ = v45
                                                        case 0:
                                                            _ = len(v1)
                                                        default:
                                                            var v46 uint8 = 127
                                                            _ = v46
                                                        }
                                                    case "a":
                                                        if false {
                                                            var v47 uint64 = 42
                                                            if true {
                                                                if false {
                                                                    v35 = 1.5
                                                                    if (uint8(v47) != 1) {
                                                                        _ = len(v1)
                                                                        fmt.Println(v11)
                                                                    } else {
                                                                        switch "foo bar" {
                                                                        case "":
                                                                            v3 = true
                                                                        default:
                                                                            var v48 float32 = (float32((float64((1e10 / v35)) + ((3.14159 + (-2.25 * 1.5)) - 0.0))) * 0.1)
                                                                            _ = v48
                                                                        }
                                                                        switch 814 {
                                                                        case 127:
                                                                            var v49 string = "hello"
                                                                            _ = v49
                                                                        case 0:
                                                                            switch (fi12 - 32768) {
                                                                            case -1:
                                                                                _ = len(v1)
                                                                            default:
                                                                                if (!(!(255 <= 2))) {
                                                                                    if (9223372036854775807 != -9223372036854775808) {
                                                                                        if (!(false && v3)) {
                                                                                            switch (-(fi38)) {
                                                                                            case 42:
                                                                                                var v50 uint8 = 255
                                                                                                _ = v50
                                                                                            case 127:
                                                                                                var v51 int16 = 42
                                                                                                _ = v51
                                                                                            case 251:
                                                                                                if v3 {
                                                                                                    v39 = v35
                                                                                                    if len(v1) >= 2 {
                                                                                                        v1 = v1[1:len(v1)]
                                                                                                    }
                                                                                                } else {
                                                                                                    var v52 float64 = ((1e10 + 0.1) - ((0.0 / 0.1) / (-2.25 - (((float64(v35) * 0.0) + (1.5 - 1e10)) - 1e10))))
                                                                                                    var v53 int32 = v11
                                                                                                    _ = v52
                                                                                                    _ = v53
                                                                                                }
                                                                                            }
                                                                                            if (0.1 <= ((3.14159 + 0.0) / 3.14159)) {
                                                                                                v2 = "\n\t"
                                                                                            }
                                                                                        }
                                                                                        v39 -= (0.1 - v39)
                                                                                        v11 -= (v11 >> ((65535) & 63))
                                                                                    } else {
                                                                                        if true {
                                                                                            var v54 int16 = 42
                                                                                            switch "foo bar" {
                                                                                            case "\n\t":
                                                                                                v39 = v4
                                                                                            case "hello":
                                                                                                if len(v1) > 0 {
                                                                                                    v1[0] = 349
                                                                                                }
                                                                                            case "a":
                                                                                                fmt.Println(v4, v54)
                                                                                            }
                                                                                            _ = v54
                                                                                        } else {
                                                                                            v47 = (v47 >> ((278) & 63))
                                                                                            v55 := GenS6{F0: 317, F1: genFn0(-32768, 42), F2: 284}
                                                                                            var v56 string = ("a" + (("foo bar" + ((v2 + "") + "")) + (v2 + (v2 + ((v2 + v2) + "foo bar")))))
                                                                                            _ = v55
                                                                                            _ = v56
                                                                                        }
                                                                                        if (!false) {
                                                                                            var v57 float64 = ((3.14159 - -2.25) + 0.1)
                                                                                            v58 := GenS7{F0: 1, F1: v11}
                                                                                            _ = v57
                                                                                            _ = v58
                                                                                        }
                                                                                    }
                                                                                    v39 += 0.0
                                                                                } else {
                                                                                    var v59 int8 = 0
                                                                                    if len(v1) > 0 {
                                                                                        v1[0] = 818
                                                                                    }
                                                                                    _ = v59
                                                                                }
                                                                            }
                                                                        case 92:
                                                                            if len(v1) > 0 {
                                                                                v1[0] = 227
                                                                            }
                                                                        default:
                                                                            v60 := GenS8{F0: 1, F1: (v47 >> ((958) & 63))}
                                                                            _ = v60
                                                                        }
                                                                        v11 = v11
                                                                        v61 := GenS4{F0: 1, F1: 32767, F2: fi38}
                                                                        if ((v3 && (!v3)) || true) {
                                                                            if (491 >= 766) {
                                                                                if len(v1) > 0 {
                                                                                    v1[0] = 534
                                                                                }
                                                                                v62 := GenS1{F0: genFn0(1, 1), F1: (v2 + ((v2 + "hello") + "hello")), F2: (v2 + (v2 + v2))}
                                                                                fmt.Println(fi38, v4, v4)
                                                                                _ = v62
                                                                            }
                                                                        }
                                                                        v47 = (v47 | 42)
                                                                        _ = v61
                                                                    }
                                                                    switch (fi38 >> ((2) & 63)) {
                                                                    case 127:
                                                                        switch (fi12 / (((fi12 | fi38)) | 1)) {
                                                                        case 42:
                                                                            if (int32(v47) == -1) {
                                                                                v47 += (v47 | (v47 / (((v47 * v47)) | 1)))
                                                                                if (v2 == "a") {
                                                                                    fmt.Println(v39)
                                                                                    var v63 int64 = 0
                                                                                    if len(v1) > 0 {
                                                                                        v1[0] = 42
                                                                                    }
                                                                                    _ = v63
                                                                                } else {
                                                                                    v3 = false
                                                                                    if v3 {
                                                                                        if (false && (!(-1 == 127))) {
                                                                                            v64 := GenS4{F0: 264, F1: 65535, F2: -2147483648}
                                                                                            fmt.Println(v4)
                                                                                            var v65 bool = v3
                                                                                            _ = v64
                                                                                            _ = v65
                                                                                        } else {
                                                                                            switch "foo bar" {
                                                                                            case "\n\t":
                                                                                                var v66 uint64 = 230
                                                                                                _ = v66
                                                                                            }
                                                                                            if v3 {
                                                                                                if true {
                                                                                                    v4 = v39
                                                                                                }
                                                                                            }
                                                                                            if len(v1) >= 2 {
                                                                                                v1 = v1[1:len(v1)]
                                                                                            }
                                                                                        }
                                                                                        v39 -= (v35 * float32(float64(0.1)))
                                                                                        if (1 < 652) {
                                                                                            if false {
                                                                                                v1 = append(v1, 2147483648)
                                                                                                v47 = (^(v47))
                                                                                                v67 := GenS1{F0: 32767, F1: "", F2: v2}
                                                                                                _ = v67
                                                                                            }
                                                                                            var v68 uint64 = v47
                                                                                            _ = v68
                                                                                        } else {
                                                                                            fmt.Println(v3, fi38)
                                                                                            var v69 uint = 255
                                                                                            v3 = (!(!((1 < -1) || v3)))
                                                                                            _ = v69
                                                                                        }
                                                                                    }
                                                                                }
                                                                                v39 -= 0.0
                                                                            } else {
                                                                                fmt.Println(fi38)
                                                                                v3 = v3
                                                                            }
                                                                        default:
                                                                            switch v2 {
                                                                            case "foo bar":
                                                                                if (-404 <= 42) {
                                                                                    v70 := GenS7{F0: 134, F1: (v11 << ((0) & 63))}
                                                                                    v71 := GenS3{F0: 42}
                                                                                    if (18446744073709551615 == 395) {
                                                                                        if v3 {
                                                                                            var v72 float32 = (v4 - v39)
                                                                                            v73 := GenS9{F0: 0, F1: 2}
                                                                                            switch "hello" {
                                                                                            case "\n\t":
                                                                                                v39 += 0.0
                                                                                            }
                                                                                            _ = v72
                                                                                            _ = v73
                                                                                        }
                                                                                    }
                                                                                    var v74 int64 = 165
                                                                                    _ = v70
                                                                                    _ = v71
                                                                                    _ = v74
                                                                                }
                                                                            }
                                                                        }
                                                                    }
                                                                } else {
                                                                    v11 = (-(v11))
                                                                    fmt.Println(v4)
                                                                }
                                                            } else {
                                                                fmt.Println(v39, fi12, v11)
                                                                if v3 {
                                                                    v2 = "hello"
                                                                    var v75 uint16 = 255
                                                                    _ = v75
                                                                }
                                                                var v76 int16 = 0
                                                                _ = v76
                                                            }
                                                            _ = v47
                                                        } else {
                                                            switch fi38 {
                                                            case -9223372036854775808:
                                                                v11 -= v11
                                                            case 386:
                                                                var v77 uint64 = 42
                                                                _ = v77
                                                            case 949:
                                                                if false {
                                                                    fmt.Println(fi38, v2, v39)
                                                                } else {
                                                                    if v3 {
                                                                        switch (^(fi38)) {
                                                                        case 32768:
                                                                            var v78 int64 = 527
                                                                            _ = v78
                                                                        case 0:
                                                                            v79 := GenS10{F0: 1}
                                                                            _ = v79
                                                                        }
                                                                        fmt.Println(fi38, v4)
                                                                        fmt.Println(fi12)
                                                                        var v80 float32 = 0.1
                                                                        _ = v80
                                                                    }
                                                                    v3 = true
                                                                }
                                                            }
                                                            _ = len(v1)
                                                        }
                                                    default:
                                                        v81 := GenS8{F0: -554, F1: 18446744073709551615}
                                                        _ = v81
                                                    }
                                                }
                                            } else {
                                                _ = len(v1)
                                            }
                                        case 128:
                                            if (10 <= 4294967295) {
                                                fmt.Println(fi12, v3, v11)
                                                if (31 < 120) {
                                                    if len(v1) > 0 {
                                                        v1[0] = 236
                                                    }
                                                }
                                                switch v2 {
                                                case "\n\t":
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                    switch (fi12 + 2147483647) {
                                                    case 127:
                                                        switch ("\n\t" + v2) {
                                                        case "":
                                                            switch "foo bar" {
                                                            case "a":
                                                                v3 = true
                                                            case "hello":
                                                                v11 -= (v11 ^ 2147483647)
                                                            case "\n\t":
                                                                var v82 int64 = 1
                                                                _ = v82
                                                            default:
                                                                v83 := GenS11{F0: fi12, F1: 707, F2: (!(0.0 == 0.1))}
                                                                _ = v83
                                                            }
                                                        case "foo bar":
                                                            if v3 {
                                                                v1 = append(v1, 1)
                                                                fmt.Println(v3)
                                                                switch (fi38 & fi38) {
                                                                case 889:
                                                                    if len(v1) >= 2 {
                                                                        v1 = v1[1:len(v1)]
                                                                    }
                                                                case 4294967295:
                                                                    var v84 uint32 = 2147483648
                                                                    _ = v84
                                                                }
                                                            } else {
                                                                v3 = false
                                                            }
                                                        case "\n\t":
                                                            v4 -= v4
                                                            var v85 string = (v2 + v2)
                                                            _ = v85
                                                        }
                                                    default:
                                                        var v86 uint8 = 235
                                                        _ = v86
                                                    }
                                                    var v87 string = v2
                                                    _ = v87
                                                case "hello":
                                                    v3 = v3
                                                case "a":
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                }
                                                if v3 {
                                                    switch "hello" {
                                                    case "\n\t":
                                                        if (genFn0(-32768, -548) > 1) {
                                                            if ((865 == 18446744073709551615) && (fi12 == (fi38 & 732))) {
                                                                if (v11 <= (v11 ^ (v11 % (((v11 - v11)) | 1)))) {
                                                                    if false {
                                                                        var v88 int8 = -128
                                                                        if (1 == 19) {
                                                                            fmt.Println(v88)
                                                                            if (4294967295 >= 384) {
                                                                                fmt.Println(v88)
                                                                                if v3 {
                                                                                    fmt.Println(v88, v35)
                                                                                } else {
                                                                                    v89 := GenS2{F0: v2}
                                                                                    v39 -= (v4 + 1e10)
                                                                                    if (!(v3 || true)) {
                                                                                        var v90 uint64 = 1
                                                                                        v91 := GenS12{F0: 9223372036854775807, F1: 255, F2: 0}
                                                                                        _ = v90
                                                                                        _ = v91
                                                                                    }
                                                                                    _ = v89
                                                                                }
                                                                                if (v3 || false) {
                                                                                    v11 = -979
                                                                                    var v92 uint64 = 4294967295
                                                                                    v35 = float32(3.14159)
                                                                                    _ = v92
                                                                                } else {
                                                                                    var v93 int16 = -1
                                                                                    var v94 string = ((v2 + v2) + "a")
                                                                                    _ = v93
                                                                                    _ = v94
                                                                                }
                                                                                if len(v1) > 0 {
                                                                                    v1[0] = 0
                                                                                }
                                                                            }
                                                                        }
                                                                        _ = v88
                                                                    }
                                                                }
                                                                fmt.Println(fi12, v2, v39)
                                                            }
                                                            if len(v1) > 0 {
                                                                v1[0] = 1
                                                            }
                                                            switch -9223372036854775808 {
                                                            case 0:
                                                                fmt.Println(v3, v35)
                                                            case -9223372036854775808:
                                                                if (0.1 != ((float64(float32(1.5)) + 1.5) - 3.14159)) {
                                                                    var v95 float64 = (float64(-2.25) - ((0.0 * (3.14159 / 1.0)) + (float64(v35) + (3.14159 / 1.0))))
                                                                    v96 := GenS2{F0: v2}
                                                                    _ = v95
                                                                    _ = v96
                                                                }
                                                            }
                                                            _ = len(v1)
                                                        } else {
                                                            if len(v1) >= 2 {
                                                                v1 = v1[1:len(v1)]
                                                            }
                                                            if false {
                                                                v4 = (0.0 + v35)
                                                                var v97 int64 = -9223372036854775808
                                                                _ = v97
                                                            } else {
                                                                var v98 string = v2
                                                                switch (fi12 + 30) {
                                                                case 4294967295:
                                                                    var v99 bool = v3
                                                                    _ = v99
                                                                case 560:
                                                                    if v3 {
                                                                        if len(v1) > 0 {
                                                                            v1[0] = 65535
                                                                        }
                                                                        v100 := GenS3{F0: 1}
                                                                        _ = v100
                                                                    }
                                                                case 128:
                                                                    var v101 int64 = -1
                                                                    _ = v101
                                                                }
                                                                var v102 int64 = 0
                                                                fmt.Println(fi38, v39, v11)
                                                                if v3 {
                                                                    if ((^(fi12)) == 2147483647) {
                                                                        if (!v3) {
                                                                            if v3 {
                                                                                switch (fi12 + fi38) {
                                                                                case 1:
                                                                                    v2 = (v2 + v98)
                                                                                case -2147483648:
                                                                                    var v103 uint64 = 178
                                                                                    _ = v103
                                                                                default:
                                                                                    fmt.Println(v4, v11, fi12)
                                                                                }
                                                                            } else {
                                                                                if (true && (42 == 1)) {
                                                                                    if (genFn0(-836, int64(fi38)) == 868) {
                                                                                        v1 = append(v1, 0)
                                                                                        switch v2 {
                                                                                        case "\n\t":
                                                                                            v3 = (42 <= 32767)
                                                                                        }
                                                                                        fmt.Println(v11, v2)
                                                                                    } else {
                                                                                        fmt.Println(v11, v2, v102)
                                                                                    }
                                                                                    if len(v1) > 0 {
                                                                                        v1[0] = 2
                                                                                    }
                                                                                    v102 -= (v102 % (((v102 ^ v102)) | 1))
                                                                                    switch ((v98 + v98) + v2) {
                                                                                    case "foo bar":
                                                                                        v35 += v39
                                                                                    }
                                                                                    v104 := GenS6{F0: 65535, F1: 0, F2: 4294967295}
                                                                                    _ = v104
                                                                                }
                                                                            }
                                                                            var v105 uint32 = 636
                                                                            if v3 {
                                                                                v106 := GenS10{F0: 65535}
                                                                                var v107 uint32 = (v105 / ((v105) | 1))
                                                                                _ = v106
                                                                                _ = v107
                                                                            } else {
                                                                                switch 9223372036854775807 {
                                                                                case 32767:
                                                                                    v11 -= v11
                                                                                }
                                                                                switch (((v2 + ((v98 + v98) + ("" + "a"))) + (v2 + v2)) + (((v2 + v2) + "hello") + "hello")) {
                                                                                case "hello":
                                                                                    v3 = (v105 == v105)
                                                                                case "a":
                                                                                    v11 = int32(fi12)
                                                                                default:
                                                                                    var v108 uint32 = (^(v105))
                                                                                    _ = v108
                                                                                }
                                                                                var v109 float32 = v39
                                                                                _ = v109
                                                                            }
                                                                            _ = v105
                                                                        } else {
                                                                            switch (v98 + "foo bar") {
                                                                            case "\n\t":
                                                                                v1 = append(v1, 4294967295)
                                                                            case "foo bar":
                                                                                var v110 int8 = 127
                                                                                _ = v110
                                                                            }
                                                                            v1 = append(v1, 739)
                                                                        }
                                                                    } else {
                                                                        if len(v1) >= 2 {
                                                                            v1 = v1[1:len(v1)]
                                                                        }
                                                                        var v111 int8 = -128
                                                                        var v112 uint8 = 2
                                                                        v113 := GenS11{F0: (fi12 / (((fi12 << ((18446744073709551615) & 63))) | 1)), F1: 475, F2: (v102 <= (^(v102)))}
                                                                        var v114 int16 = 442
                                                                        _ = v111
                                                                        _ = v112
                                                                        _ = v113
                                                                        _ = v114
                                                                    }
                                                                    if len(v1) > 0 {
                                                                        v1[0] = 65535
                                                                    }
                                                                    if (!true) {
                                                                        v11 = (v11 << ((42) & 63))
                                                                        v2 = v2
                                                                    } else {
                                                                        fmt.Println(v102)
                                                                    }
                                                                } else {
                                                                    if v3 {
                                                                        var v115 uint8 = 1
                                                                        v3 = v3
                                                                        var v116 int16 = -32768
                                                                        var v117 int64 = -486
                                                                        _ = v115
                                                                        _ = v116
                                                                        _ = v117
                                                                    }
                                                                    var v118 float64 = 0.0
                                                                    v119 := GenS13{F0: 4294967295, F1: ("foo bar" + "a")}
                                                                    _ = v118
                                                                    _ = v119
                                                                }
                                                                _ = v102
                                                                _ = v98
                                                            }
                                                            fmt.Println(fi12, v2)
                                                            v3 = (int8(v11) < -128)
                                                            fmt.Println(fi12, v39, v35)
                                                            if len(v1) >= 2 {
                                                                v1 = v1[1:len(v1)]
                                                            }
                                                        }
                                                    }
                                                    switch (fi38 << ((9) & 63)) {
                                                    case 42:
                                                        if (!true) {
                                                            if false {
                                                                var v120 float32 = (((v4 - v4) * 0.0) * ((1e10 - float32(1.5)) + v35))
                                                                _ = v120
                                                            } else {
                                                                switch fi12 {
                                                                case 0:
                                                                    v4 -= 3.14159
                                                                case 1:
                                                                    switch "hello" {
                                                                    case "foo bar":
                                                                        if (65535 >= 1) {
                                                                            var v121 float32 = (float32(1.5) + ((float32(3.14159) / v4) / (1.5 * -2.25)))
                                                                            v39 += (v121 + 3.14159)
                                                                            _ = v121
                                                                        } else {
                                                                            if (!v3) {
                                                                                var v122 int = (-(fi12))
                                                                                if (v3 && v3) {
                                                                                    v123 := GenS14{F0: (v11 << ((0) & 63))}
                                                                                    var v124 uint64 = 980
                                                                                    v3 = v3
                                                                                    if (536 >= 2147483647) {
                                                                                        var v125 uint = 1
                                                                                        switch v122 {
                                                                                        case -827:
                                                                                            v124 = v124
                                                                                        case 128:
                                                                                            v11 += v11
                                                                                            v126 := GenS4{F0: 65535, F1: 394, F2: (fi12 - (v122 ^ (fi38 & fi12)))}
                                                                                            _ = v126
                                                                                        case -841:
                                                                                            var v127 uint64 = v124
                                                                                            _ = v127
                                                                                        default:
                                                                                            fmt.Println(fi38, v4, v11)
                                                                                        }
                                                                                        v2 = ""
                                                                                        fmt.Println(v125, v4, v11)
                                                                                        v122 = (^(v122))
                                                                                        var v128 uint16 = 890
                                                                                        if (-9223372036854775808 > int64(fi12)) {
                                                                                            v129 := GenS15{F0: fi38, F1: v3, F2: 46}
                                                                                            switch fi38 {
                                                                                            case 1:
                                                                                                v4 += v35
                                                                                            case 42:
                                                                                                var v130 int64 = 1
                                                                                                _ = v130
                                                                                            }
                                                                                            _ = v129
                                                                                        }
                                                                                        _ = v125
                                                                                        _ = v128
                                                                                    }
                                                                                    v122 += (fi12 - fi38)
                                                                                    _ = v123
                                                                                    _ = v124
                                                                                }
                                                                                switch ("hello" + v2) {
                                                                                case "foo bar":
                                                                                    fmt.Println(fi38)
                                                                                case "hello":
                                                                                    if (v2 >= "\n\t") {
                                                                                        v11 += (v11 & v11)
                                                                                        v122 += v122
                                                                                    }
                                                                                }
                                                                                _ = v122
                                                                            }
                                                                            v131 := GenS5{F0: -9223372036854775808, F1: -1}
                                                                            fmt.Println(v3)
                                                                            v132 := GenS10{F0: 620}
                                                                            if (3.14159 < 0.0) {
                                                                                var v133 uint64 = 0
                                                                                var v134 float64 = -2.25
                                                                                _ = v133
                                                                                _ = v134
                                                                            }
                                                                            _ = v131
                                                                            _ = v132
                                                                        }
                                                                    case "a":
                                                                        if len(v1) > 0 {
                                                                            v1[0] = 106
                                                                        }
                                                                    }
                                                                default:
                                                                    if (!(false && (v3 || (true && v3)))) {
                                                                        var v135 uint = 4294967295
                                                                        fmt.Println(v11)
                                                                        var v136 uint32 = 1
                                                                        _ = v135
                                                                        _ = v136
                                                                    }
                                                                }
                                                                if (!v3) {
                                                                    var v137 int16 = 819
                                                                    v3 = v3
                                                                    _ = v137
                                                                } else {
                                                                    var v138 uint32 = 2147483648
                                                                    _ = v138
                                                                }
                                                            }
                                                        } else {
                                                            switch fi12 {
                                                            case -9223372036854775808:
                                                                var v139 int64 = 1
                                                                _ = v139
                                                            case -472:
                                                                var v140 int64 = 663
                                                                _ = v140
                                                            }
                                                            if v3 {
                                                                if (!(786 > 127)) {
                                                                    fmt.Println(fi12)
                                                                    fmt.Println(v11, v39)
                                                                }
                                                                v2 = ((v2 + ("hello" + (v2 + v2))) + v2)
                                                            }
                                                        }
                                                    case 9223372036854775807:
                                                        if v3 {
                                                            switch fi12 {
                                                            case 32767:
                                                                if (42 < 0) {
                                                                    v11 = (v11 >> ((449) & 63))
                                                                    var v141 uint64 = 18446744073709551615
                                                                    var v142 uint32 = 1
                                                                    _ = v141
                                                                    _ = v142
                                                                }
                                                            }
                                                        }
                                                    }
                                                    v143 := GenS5{F0: 498, F1: -32768}
                                                    _ = v143
                                                } else {
                                                    if v3 {
                                                        v144 := GenS16{F0: 718, F1: 42}
                                                        v1 = append(v1, 2147483647)
                                                        _ = v144
                                                    }
                                                }
                                                v39 += v39
                                            } else {
                                                v2 = v2
                                                if (v3 || v3) {
                                                    fmt.Println(v39)
                                                    fmt.Println(v3)
                                                    var v145 uint16 = 1
                                                    if len(v1) > 0 {
                                                        v1[0] = 2
                                                    }
                                                    _ = v145
                                                }
                                                var v146 bool = (-116 < -105)
                                                v11 = v11
                                                v35 += (v4 * (0.1 + float32(0.1)))
                                                _ = v146
                                            }
                                        }
                                        if (!true) {
                                            fmt.Println(v3, v11, fi38)
                                        } else {
                                            switch fi12 {
                                            case 4294967295:
                                                if (63 > 0) {
                                                    var v147 float32 = 3.14159
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                    _ = v147
                                                } else {
                                                    switch "foo bar" {
                                                    case "":
                                                        v1 = append(v1, 386)
                                                    case "a":
                                                        if v3 {
                                                            if (100 > -51) {
                                                                v1 = append(v1, 0)
                                                                v148 := GenS3{F0: 1}
                                                                var v149 int8 = 100
                                                                _ = v148
                                                                _ = v149
                                                            } else {
                                                                var v150 int = 0
                                                                _ = v150
                                                            }
                                                        }
                                                    case "hello":
                                                        switch fi12 {
                                                        case 32768:
                                                            var v151 uint32 = 2
                                                            _ = v151
                                                        }
                                                        var v152 uint16 = 255
                                                        switch 2147483647 {
                                                        case 1:
                                                            fmt.Println(v2, v2, v4)
                                                        default:
                                                            v153 := GenS17{F0: 1}
                                                            _ = v153
                                                        }
                                                        var v154 int8 = 100
                                                        v154 -= (v154 / (((^(v154))) | 1))
                                                        _ = v152
                                                        _ = v154
                                                    default:
                                                        var v155 float32 = v35
                                                        _ = v155
                                                    }
                                                    fmt.Println(v4, v4, v4)
                                                }
                                            case 193:
                                                var v156 uint64 = 42
                                                v157 := GenS18{F0: (!(0.1 <= 1e10))}
                                                if v3 {
                                                    switch (fi12 & -9223372036854775808) {
                                                    case 0:
                                                        switch fi38 {
                                                        case 842:
                                                            if (v2 != (v2 + "foo bar")) {
                                                                _ = len(v1)
                                                            } else {
                                                                if v3 {
                                                                    fmt.Println(v4, v4)
                                                                    v1 = append(v1, 0)
                                                                } else {
                                                                    var v158 int64 = -767
                                                                    v159 := GenS6{F0: 148, F1: 42, F2: 2147483648}
                                                                    _ = v158
                                                                    _ = v159
                                                                }
                                                            }
                                                        default:
                                                            v39 += (0.0 + v35)
                                                        }
                                                    default:
                                                        var v160 uint8 = 127
                                                        _ = v160
                                                    }
                                                    var v161 int8 = 0
                                                    if len(v1) > 0 {
                                                        v1[0] = 1
                                                    }
                                                    _ = v161
                                                } else {
                                                    fmt.Println(v39, v2, v35)
                                                }
                                                _ = v156
                                                _ = v157
                                            case 127:
                                                var v162 int8 = -16
                                                _ = v162
                                            default:
                                                v35 = float32(((0.0 + 3.14159) / (3.14159 * (float64((v39 * v39)) - 0.1))))
                                            }
                                            var v163 int64 = 42
                                            _ = v163
                                        }
                                        v35 -= v39
                                        fmt.Println(fi38, v2, v11)
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                    }
                                }
                                v3 = (((((0 != -1) && (v4 > -2.25)) && v3) && (-32768 < -818)) || (0 > (fi38 << ((330) & 63))))
                                _ = fi38
                                _ = v39
                            }
                            var v164 uint8 = 2
                            switch "" {
                            case "":
                                if (true || true) {
                                    fmt.Println(v4, v4)
                                }
                            case "a":
                                for fi165 := 0; fi165 < 9; fi165++ {
                                    v164 += (v164 * v164)
                                    _ = fi165
                                }
                            default:
                                _ = len(v1)
                            }
                            _ = v164
                        }
                        _ = v37
                    } else {
                        var v166 uint64 = 839
                        v3 = (885 >= 172)
                        if (42 == 127) {
                            v167 := GenS19{F0: 0, F1: (fi12 - fi12), F2: 255}
                            v11 -= (^(v11))
                            fmt.Println(v2, v11, fi12)
                            _ = v167
                        } else {
                            var v168 int32 = v11
                            switch (v2 + "hello") {
                            case "hello":
                                fmt.Println(v11, v168)
                            }
                            _ = v168
                        }
                        _ = v166
                    }
                    _ = v35
                }
            }
            for fi169 := 0; fi169 < 19; fi169++ {
                _ = len(v1)
                _ = fi169
            }
            _ = fi12
        }
        _ = len(v1)
        if v3 {
            if (0 <= 0) {
                if v3 {
                    fmt.Println(v4, v4, v4)
                }
            } else {
                v11 += v11
                v3 = (1 == 1)
                var v170 int16 = 1000
                _ = v170
            }
            if (42 > -87) {
                var v171 uint64 = 2
                _ = v171
            } else {
                if (1 > 42) {
                    for fi172 := 0; fi172 < 4; fi172++ {
                        fmt.Println(v3, v3)
                        if v3 {
                            var v173 uint64 = 2
                            var v174 int32 = (v11 & (v11 | (v11 & v11)))
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            _ = v173
                            _ = v174
                        } else {
                            if true {
                                for fi175 := 0; fi175 < 16; fi175++ {
                                    v3 = v3
                                    var v176 int64 = 42
                                    _ = fi175
                                    _ = v176
                                }
                            } else {
                                if true {
                                    v177 := GenS17{F0: -1}
                                    var v178 string = (v2 + "foo bar")
                                    for fi179 := 0; fi179 < 20; fi179++ {
                                        var v180 int = 32767
                                        _ = fi179
                                        _ = v180
                                    }
                                    _ = v177
                                    _ = v178
                                }
                                switch (v2 + v2) {
                                case "":
                                    switch 537 {
                                    case 127:
                                        v181 := GenS12{F0: 0, F1: 406, F2: 1}
                                        _ = v181
                                    case 332:
                                        v4 -= v4
                                    case -537:
                                        switch (fi172 * fi172) {
                                        case 9223372036854775807:
                                            switch (fi172 & (fi172 % (((fi172 % (((-(fi172))) | 1))) | 1))) {
                                            case -296:
                                                var v182 float64 = (-2.25 * 1.5)
                                                _ = v182
                                            }
                                        }
                                    }
                                case "a":
                                    switch v2 {
                                    case "":
                                        var v183 uint64 = 163
                                        _ = v183
                                    }
                                }
                            }
                            for fi184 := 0; fi184 < 17; fi184++ {
                                var v185 uint = 47
                                fmt.Println(v4, v185, v11)
                                _ = fi184
                                _ = v185
                            }
                            if (1 < 100) {
                                switch ((v2 + "\n\t") + "a") {
                                case "a":
                                    fmt.Println(v2, v11)
                                case "hello":
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                default:
                                    v2 = v2
                                }
                                switch v2 {
                                case "foo bar":
                                    v3 = v3
                                case "hello":
                                    for fi186 := 0; fi186 < 10; fi186++ {
                                        v187 := GenS7{F0: 128, F1: (^(v11))}
                                        _ = len(v1)
                                        _ = fi186
                                        _ = v187
                                    }
                                }
                                v188 := GenS20{F0: v3}
                                _ = v188
                            } else {
                                if (1 > (v11 | v11)) {
                                    var v189 uint = 470
                                    v190 := GenS15{F0: (fi172 * (fi172 >> ((780) & 63))), F1: v3, F2: -1}
                                    v3 = (!false)
                                    _ = v189
                                    _ = v190
                                } else {
                                    for fi191 := 0; fi191 < 8; fi191++ {
                                        var v192 float64 = (0.1 / 1e10)
                                        var v193 float64 = ((1e10 / 1e10) + -2.25)
                                        if (2147483647 > 2147483648) {
                                            var v194 bool = v3
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            v192 = -2.25
                                            if len(v1) > 0 {
                                                v1[0] = 4294967295
                                            }
                                            _ = v194
                                        }
                                        _ = fi191
                                        _ = v192
                                        _ = v193
                                    }
                                    for fi195 := 0; fi195 < 3; fi195++ {
                                        var v196 uint8 = 128
                                        fmt.Println(v11, fi195, fi195)
                                        switch v2 {
                                        case "\n\t":
                                            v2 = ("\n\t" + "hello")
                                        default:
                                            v196 -= (^(v196))
                                        }
                                        _ = fi195
                                        _ = v196
                                    }
                                }
                                _ = len(v1)
                                if true {
                                    fmt.Println(v11)
                                    var v197 bool = (-1 >= 0)
                                    _ = v197
                                }
                            }
                        }
                        _ = fi172
                    }
                    switch 127 {
                    case 128:
                        fmt.Println(v11, v2)
                    case 4294967295:
                        v198 := GenS21{F0: 607, F1: 0, F2: -128}
                        _ = v198
                    }
                } else {
                    switch ((((v2 + v2) + "\n\t") + (v2 + ((v2 + ("hello" + "")) + ""))) + v2) {
                    case "a":
                        if v3 {
                            if (!(125 == 100)) {
                                var v199 int64 = 0
                                _ = v199
                            }
                            fmt.Println(v3, v2)
                        } else {
                            var v200 bool = (446 < 815)
                            v4 -= v4
                            _ = v200
                        }
                    case "hello":
                        var v201 int16 = 42
                        _ = v201
                    case "foo bar":
                        if v3 {
                            var v202 float64 = 0.0
                            if len(v1) > 0 {
                                v1[0] = 2147483648
                            }
                            _ = v202
                        } else {
                            switch "a" {
                            case "a":
                                v203 := GenS3{F0: 52}
                                _ = v203
                            case "\n\t":
                                v3 = ((((!(32767 == 1000)) && v3) && (126 == 332)) && true)
                            case "foo bar":
                                v204 := GenS6{F0: 42, F1: 1, F2: 486}
                                _ = v204
                            }
                        }
                    }
                }
                if true {
                    v205 := GenS22{F0: 4294967295, F1: v2, F2: (255 > 1)}
                    for fi206 := 0; fi206 < 17; fi206++ {
                        var v207 int8 = 42
                        _ = fi206
                        _ = v207
                    }
                    _ = v205
                } else {
                    v208 := GenS18{F0: v3}
                    fmt.Println(v3, v3, v3)
                    _ = v208
                }
                v209 := GenS21{F0: 266, F1: 785, F2: 127}
                _ = v209
            }
            _ = len(v1)
        } else {
            _ = len(v1)
            switch 720 {
            case 127:
                for fi210 := 0; fi210 < 13; fi210++ {
                    fmt.Println(fi210, v4, v11)
                    v4 = v4
                    if false {
                        switch fi210 {
                        case 32767:
                            v3 = v3
                        default:
                            v211 := GenS13{F0: 467, F1: "a"}
                            _ = v211
                        }
                        if (v3 || (false && false)) {
                            switch fi210 {
                            case -841:
                                fmt.Println(v11, fi210)
                            }
                        } else {
                            var v212 uint64 = 42
                            var v213 uint32 = 4294967295
                            if (v212 <= v212) {
                                for fi214 := 0; fi214 < 9; fi214++ {
                                    v215 := GenS23{F0: v2}
                                    _ = fi214
                                    _ = v215
                                }
                                var v216 int16 = 42
                                _ = v216
                            } else {
                                _ = len(v1)
                            }
                            _ = v212
                            _ = v213
                        }
                        if true {
                            for fi217 := 0; fi217 < 16; fi217++ {
                                v4 = ((v4 - -2.25) - 1.5)
                                v3 = v3
                                _ = fi217
                            }
                            var v218 uint8 = 2
                            _ = v218
                        }
                    }
                    _ = fi210
                }
            default:
                v219 := GenS24{F0: 474, F1: -1, F2: v2}
                _ = v219
            }
        }
    }
    if v3 {
        var v220 bool = (((v4 <= v4) || v3) || ((-1 != -1) && v3))
        for fi221 := 0; fi221 < 5; fi221++ {
            var v222 int16 = 1
            if (v220 || (42 == 0)) {
                if (42 > 0) {
                    var v223 string = v2
                    var v224 float64 = (1.5 * float64(3.14159))
                    var v225 uint8 = 0
                    _ = v223
                    _ = v224
                    _ = v225
                }
                if v3 {
                    v222 = v222
                    fmt.Println(v2, v220)
                }
                v222 = (v222 << ((255) & 63))
            }
            v220 = v3
            _ = fi221
            _ = v222
        }
        _ = v220
    }
    for fi226 := 0; fi226 < 14; fi226++ {
        var v227 uint = 0
        if v3 {
            var v228 uint8 = 121
            _ = v228
        }
        var v229 uint16 = 255
        _ = fi226
        _ = v227
        _ = v229
    }
    var v230 int8 = -1
    for fi231 := 0; fi231 < 17; fi231++ {
        var v232 uint64 = 1
        _ = fi231
        _ = v232
    }
    v1 = append(v1, 976)
    v233 := GenS19{F0: 0, F1: 32767, F2: 864}
    for fi234 := 0; fi234 < 7; fi234++ {
        fmt.Println(v4)
        _ = fi234
    }
    for fi235 := 0; fi235 < 12; fi235++ {
        for fi236 := 0; fi236 < 12; fi236++ {
            v230 = (v230 ^ (^(v230)))
            _ = fi236
        }
        _ = fi235
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v230:", v230)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v2
    _ = v230
    _ = v233
    _ = v3
    _ = v4
}
