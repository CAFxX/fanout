package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8 }
type GenS2 struct { F0 int16; F1 uint8; F2 int32 }
type GenS3 struct { F0 uint32; F1 int8; F2 uint16 }
type GenS4 struct { F0 int32; F1 int }
type GenS5 struct { F0 string; F1 bool }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 int64; F1 uint32; F2 bool }
type GenS8 struct { F0 int8; F1 uint64; F2 int }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 uint64; F1 int }
type GenS11 struct { F0 int16 }
type GenS12 struct { F0 uint16; F1 int8; F2 bool }
type GenS13 struct { F0 uint32; F1 uint32; F2 uint }
type GenS14 struct { F0 int64; F1 int16 }
type GenS15 struct { F0 uint64 }
type GenS16 struct { F0 int }
type GenS17 struct { F0 uint8; F1 int8 }
type GenS18 struct { F0 uint64 }
type GenS19 struct { F0 int16; F1 uint16 }

func main() {
    v1 := []string{("" + "hello"), ("\n\t" + ("\n\t" + "foo bar")), "\n\t"}
    v2 := []uint32{2147483647}
    var v3 float32 = (0.0 - float32(1.5))
    var v4 int = -9223372036854775808
    var v5 int = v4
    fmt.Println(v4, v5, v3)
    v6 := GenS1{F0: 408, F1: 100}
    var v7 uint16 = 32767
    if ((255 != 44) || (true && true)) {
        if (!true) {
            v4 = v5
            _ = len(v2)
        }
        fmt.Println(v7, v3, v7)
    }
    switch ("a" + "\n\t") {
    case "\n\t":
        switch (("a" + "a") + "") {
        case "hello":
            var v8 uint16 = v7
            _ = v8
        case "a":
            v5 = v5
        case "\n\t":
            v9 := GenS1{F0: (v7 ^ (v7 & 255)), F1: 127}
            _ = v9
        default:
            if (!(226 >= 1)) {
                var v10 bool = false
                if (-8 < int8(v7)) {
                    v3 = v3
                    if (0 > -517) {
                        switch ("foo bar" + "hello") {
                        case "":
                            for fi11 := 0; fi11 < 20; fi11++ {
                                fmt.Println(v5, fi11)
                                var v12 bool = v10
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                _ = fi11
                                _ = v12
                            }
                        case "foo bar":
                            switch "foo bar" {
                            case "":
                                var v13 int = (v5 >> ((20) & 63))
                                _ = v13
                            case "hello":
                                _ = len(v2)
                            case "a":
                                v14 := GenS2{F0: 1, F1: 42, F2: -1}
                                _ = v14
                            }
                        default:
                            if v10 {
                                var v15 float32 = (v3 / float32((1.5 + (float64((v3 - v3)) + 0.1))))
                                v7 = v7
                                fmt.Println(v5, v7)
                                _ = v15
                            } else {
                                switch (("hello" + "foo bar") + "") {
                                case "hello":
                                    if (!v10) {
                                        var v16 int = 128
                                        _ = v16
                                    } else {
                                        fmt.Println(v3, v3, v7)
                                    }
                                default:
                                    switch "a" {
                                    case "\n\t":
                                        if (42 < -1) {
                                            switch "" {
                                            case "a":
                                                if v10 {
                                                    if (1 <= 128) {
                                                        if (!((-32768 > -32768) && v10)) {
                                                            v4 = v4
                                                        }
                                                        v17 := GenS3{F0: 2147483647, F1: 74, F2: v7}
                                                        for fi18 := 0; fi18 < 19; fi18++ {
                                                            if len(v2) > 0 {
                                                                v2[0] = 753
                                                            }
                                                            var v19 string = (("foo bar" + ("hello" + ("hello" + ("hello" + "a")))) + ("hello" + ""))
                                                            _ = fi18
                                                            _ = v19
                                                        }
                                                        _ = v17
                                                    }
                                                    for fi20 := 0; fi20 < 14; fi20++ {
                                                        v21 := GenS3{F0: 4294967295, F1: 1, F2: v7}
                                                        if (!(-2147483648 <= -101)) {
                                                            var v22 int16 = 0
                                                            switch (^(v5)) {
                                                            case 128:
                                                                fmt.Println(v4, v3, v4)
                                                            default:
                                                                if (!(782 < 0)) {
                                                                    v23 := GenS1{F0: v7, F1: -36}
                                                                    if v10 {
                                                                        v22 += v22
                                                                        if false {
                                                                            v5 -= fi20
                                                                        }
                                                                        v24 := GenS4{F0: 0, F1: (fi20 >> ((56) & 63))}
                                                                        _ = v24
                                                                    } else {
                                                                        if len(v1) > 0 {
                                                                            v1[0] = ("\n\t" + "hello")
                                                                        }
                                                                    }
                                                                    _ = v23
                                                                } else {
                                                                    if (1 == 2147483647) {
                                                                        for fi25 := 0; fi25 < 9; fi25++ {
                                                                            _ = len(v1)
                                                                            v26 := GenS3{F0: 65535, F1: 100, F2: 7}
                                                                            v7 -= 594
                                                                            _ = fi25
                                                                            _ = v26
                                                                        }
                                                                        v7 = v7
                                                                    } else {
                                                                        switch "foo bar" {
                                                                        case "\n\t":
                                                                            v2 = append(v2, 2147483648)
                                                                        case "foo bar":
                                                                            if (42 < 712) {
                                                                                var v27 float64 = (1e10 / -2.25)
                                                                                if v10 {
                                                                                    var v28 string = ("" + "hello")
                                                                                    _ = v28
                                                                                } else {
                                                                                    for fi29 := 0; fi29 < 20; fi29++ {
                                                                                        if len(v1) >= 2 {
                                                                                            v1 = v1[1:len(v1)]
                                                                                        }
                                                                                        fmt.Println(v10, v3)
                                                                                        _ = fi29
                                                                                    }
                                                                                    var v30 uint64 = 42
                                                                                    _ = v30
                                                                                }
                                                                                switch ("" + "\n\t") {
                                                                                case "":
                                                                                    v31 := GenS5{F0: "\n\t", F1: (!v10)}
                                                                                    _ = v31
                                                                                case "\n\t":
                                                                                    for fi32 := 0; fi32 < 9; fi32++ {
                                                                                        if (v10 || ((255 >= 0) || true)) {
                                                                                            fmt.Println(v7, fi32, fi20)
                                                                                            fmt.Println(v7, v4, v27)
                                                                                            fmt.Println(v4, fi20, v7)
                                                                                        } else {
                                                                                            var v33 float32 = (1.5 + 1e10)
                                                                                            var v34 int16 = 1000
                                                                                            _ = v33
                                                                                            _ = v34
                                                                                        }
                                                                                        if (1 != 515) {
                                                                                            if false {
                                                                                                if len(v2) >= 2 {
                                                                                                    v2 = v2[1:len(v2)]
                                                                                                }
                                                                                                v7 = v7
                                                                                                var v35 float32 = v3
                                                                                                _ = v35
                                                                                            } else {
                                                                                                var v36 uint32 = 2
                                                                                                _ = v36
                                                                                            }
                                                                                        }
                                                                                        var v37 bool = (65535 > 255)
                                                                                        var v38 uint64 = 18446744073709551615
                                                                                        _ = fi32
                                                                                        _ = v37
                                                                                        _ = v38
                                                                                    }
                                                                                case "a":
                                                                                    for fi39 := 0; fi39 < 5; fi39++ {
                                                                                        switch "foo bar" {
                                                                                        case "foo bar":
                                                                                            v7 = (v7 * 65535)
                                                                                        case "\n\t":
                                                                                            _ = len(v2)
                                                                                        default:
                                                                                            if (!(v10 || (843 != 1))) {
                                                                                                fmt.Println(v27)
                                                                                                fmt.Println(fi20, v27)
                                                                                                v7 = v7
                                                                                                var v40 int32 = -1
                                                                                                if (!(v40 >= (v40 % ((-359) | 1)))) {
                                                                                                    v7 -= v7
                                                                                                    switch "hello" {
                                                                                                    case "foo bar":
                                                                                                        v22 = v22
                                                                                                    }
                                                                                                }
                                                                                                fmt.Println(v10, v5)
                                                                                                _ = v40
                                                                                            }
                                                                                        }
                                                                                        if v10 {
                                                                                            _ = len(v2)
                                                                                            if true {
                                                                                                if (false || (true && v10)) {
                                                                                                    fmt.Println(v5, fi20)
                                                                                                    if len(v1) > 0 {
                                                                                                        v1[0] = ""
                                                                                                    }
                                                                                                    if v10 {
                                                                                                        if (!v10) {
                                                                                                            switch "\n\t" {
                                                                                                            case "a":
                                                                                                                var v41 float32 = (((((v3 / (0.1 * -2.25)) * 3.14159) - ((v3 - 0.1) - (float32(v27) * (v3 - v3)))) * 1e10) - float32(1e10))
                                                                                                                v42 := GenS6{F0: 74}
                                                                                                                _ = v41
                                                                                                                _ = v42
                                                                                                            case "":
                                                                                                                var v43 int = v5
                                                                                                                _ = v43
                                                                                                            case "hello":
                                                                                                                v2 = append(v2, 947)
                                                                                                            default:
                                                                                                                v7 += v7
                                                                                                            }
                                                                                                            var v44 int64 = 9223372036854775807
                                                                                                            fmt.Println(v22, v27)
                                                                                                            _ = v44
                                                                                                        } else {
                                                                                                            if v10 {
                                                                                                                if (0 != 926) {
                                                                                                                    fmt.Println(v7, v22, v5)
                                                                                                                    v45 := GenS2{F0: (-(v22)), F1: 127, F2: -940}
                                                                                                                    fmt.Println(v4)
                                                                                                                    _ = v45
                                                                                                                }
                                                                                                            } else {
                                                                                                                switch fi20 {
                                                                                                                case 512:
                                                                                                                    v10 = v10
                                                                                                                case 32767:
                                                                                                                    if len(v1) >= 2 {
                                                                                                                        v1 = v1[1:len(v1)]
                                                                                                                    }
                                                                                                                }
                                                                                                            }
                                                                                                            v4 += 32767
                                                                                                            if (true || (("hello" + "\n\t") < ("a" + "foo bar"))) {
                                                                                                                switch "hello" {
                                                                                                                case "foo bar":
                                                                                                                    if (!v10) {
                                                                                                                        v4 -= v4
                                                                                                                        if len(v2) >= 2 {
                                                                                                                            v2 = v2[1:len(v2)]
                                                                                                                        }
                                                                                                                    } else {
                                                                                                                        var v46 int16 = (v22 << ((129) & 63))
                                                                                                                        v47 := GenS6{F0: 128}
                                                                                                                        _ = v46
                                                                                                                        _ = v47
                                                                                                                    }
                                                                                                                case "\n\t":
                                                                                                                    if (-128 == 42) {
                                                                                                                        if (!(v27 != (v27 + v27))) {
                                                                                                                            switch (v5 % ((fi39) | 1)) {
                                                                                                                            case -713:
                                                                                                                                if len(v2) >= 2 {
                                                                                                                                    v2 = v2[1:len(v2)]
                                                                                                                                }
                                                                                                                            case 790:
                                                                                                                                var v48 uint16 = v7
                                                                                                                                var v49 int = (v4 - 32767)
                                                                                                                                if (v3 <= v3) {
                                                                                                                                    v50 := GenS7{F0: 82, F1: 42, F2: (float32(0.0) >= (v3 * (v3 + v3)))}
                                                                                                                                    var v51 int = (-(fi39))
                                                                                                                                    if (((!(false || v10)) && (-91 >= 42)) || (-385 != 0)) {
                                                                                                                                        if v10 {
                                                                                                                                            var v52 int64 = 1
                                                                                                                                            v10 = (42 <= 18446744073709551615)
                                                                                                                                            _ = v52
                                                                                                                                        }
                                                                                                                                    }
                                                                                                                                    _ = v50
                                                                                                                                    _ = v51
                                                                                                                                }
                                                                                                                                v4 += fi20
                                                                                                                                _ = v48
                                                                                                                                _ = v49
                                                                                                                            case 127:
                                                                                                                                var v53 int8 = -128
                                                                                                                                _ = v53
                                                                                                                            }
                                                                                                                            v22 = int16(v7)
                                                                                                                        } else {
                                                                                                                            fmt.Println(v22, v4)
                                                                                                                            v22 -= 0
                                                                                                                            if true {
                                                                                                                                switch ("hello" + ("a" + "\n\t")) {
                                                                                                                                case "a":
                                                                                                                                    v54 := GenS8{F0: 42, F1: 4294967295, F2: fi39}
                                                                                                                                    _ = v54
                                                                                                                                case "\n\t":
                                                                                                                                    v55 := GenS9{F0: 744}
                                                                                                                                    _ = v55
                                                                                                                                case "foo bar":
                                                                                                                                    v7 -= (v7 % (((v7 >> ((255) & 63))) | 1))
                                                                                                                                }
                                                                                                                                v22 -= 1000
                                                                                                                                v10 = false
                                                                                                                            }
                                                                                                                            switch (v5 >> ((42) & 63)) {
                                                                                                                            case -765:
                                                                                                                                if len(v2) > 0 {
                                                                                                                                    v2[0] = 482
                                                                                                                                }
                                                                                                                            case -1:
                                                                                                                                v56 := GenS10{F0: 4294967295, F1: (fi39 - (fi39 % (((v4 - -422)) | 1)))}
                                                                                                                                _ = v56
                                                                                                                            case 2147483647:
                                                                                                                                if v10 {
                                                                                                                                    v57 := GenS2{F0: (v22 & v22), F1: 128, F2: 0}
                                                                                                                                    _ = v57
                                                                                                                                } else {
                                                                                                                                    switch (fi39 | v5) {
                                                                                                                                    case 42:
                                                                                                                                        var v58 uint16 = v7
                                                                                                                                        _ = v58
                                                                                                                                    }
                                                                                                                                    v27 = ((v27 * 0.0) + 1e10)
                                                                                                                                    var v59 uint64 = 4294967295
                                                                                                                                    _ = v59
                                                                                                                                }
                                                                                                                            default:
                                                                                                                                _ = len(v1)
                                                                                                                            }
                                                                                                                            if (v22 == v22) {
                                                                                                                                v60 := GenS11{F0: v22}
                                                                                                                                fmt.Println(fi20, v22)
                                                                                                                                _ = v60
                                                                                                                            }
                                                                                                                        }
                                                                                                                        var v61 int64 = 1
                                                                                                                        v10 = true
                                                                                                                        _ = v61
                                                                                                                    } else {
                                                                                                                        v4 -= 450
                                                                                                                        v62 := GenS12{F0: (v7 / ((1) | 1)), F1: 42, F2: true}
                                                                                                                        v3 += 3.14159
                                                                                                                        _ = v62
                                                                                                                    }
                                                                                                                default:
                                                                                                                    switch ("a" + "hello") {
                                                                                                                    case "a":
                                                                                                                        if v10 {
                                                                                                                            v10 = (!((!true) || ((v7 ^ (v7 * 42)) == (v7 << ((255) & 63)))))
                                                                                                                            v22 = v22
                                                                                                                            var v63 float32 = 0.1
                                                                                                                            _ = v63
                                                                                                                        } else {
                                                                                                                            v22 += v22
                                                                                                                            if ((^(fi39)) >= (fi20 * v5)) {
                                                                                                                                if len(v2) >= 2 {
                                                                                                                                    v2 = v2[1:len(v2)]
                                                                                                                                }
                                                                                                                                v2 = append(v2, 2)
                                                                                                                                if true {
                                                                                                                                    if v10 {
                                                                                                                                        v64 := GenS13{F0: 2, F1: 723, F2: 1}
                                                                                                                                        _ = v64
                                                                                                                                    }
                                                                                                                                    if len(v2) >= 2 {
                                                                                                                                        v2 = v2[1:len(v2)]
                                                                                                                                    }
                                                                                                                                }
                                                                                                                            } else {
                                                                                                                                _ = len(v1)
                                                                                                                            }
                                                                                                                        }
                                                                                                                        v65 := GenS11{F0: -32768}
                                                                                                                        _ = v65
                                                                                                                    }
                                                                                                                }
                                                                                                            } else {
                                                                                                                var v66 int = fi39
                                                                                                                switch (("hello" + "a") + "a") {
                                                                                                                case "":
                                                                                                                    v67 := GenS10{F0: 373, F1: (fi20 & (^(v5)))}
                                                                                                                    _ = v67
                                                                                                                default:
                                                                                                                    v68 := GenS13{F0: 42, F1: 664, F2: 65535}
                                                                                                                    fmt.Println(fi20)
                                                                                                                    _ = v68
                                                                                                                }
                                                                                                                _ = v66
                                                                                                            }
                                                                                                        }
                                                                                                    } else {
                                                                                                        switch fi20 {
                                                                                                        case 4294967295:
                                                                                                            switch v5 {
                                                                                                            case 0:
                                                                                                                if false {
                                                                                                                    if true {
                                                                                                                        if v10 {
                                                                                                                            v69 := GenS8{F0: -125, F1: 335, F2: (fi20 * 32767)}
                                                                                                                            _ = len(v2)
                                                                                                                            _ = v69
                                                                                                                        } else {
                                                                                                                            v4 += int(v7)
                                                                                                                        }
                                                                                                                    }
                                                                                                                    var v70 uint8 = 0
                                                                                                                    _ = v70
                                                                                                                } else {
                                                                                                                    v71 := GenS4{F0: 0, F1: 815}
                                                                                                                    if len(v1) >= 2 {
                                                                                                                        v1 = v1[1:len(v1)]
                                                                                                                    }
                                                                                                                    if (100 <= -1) {
                                                                                                                        v27 -= v27
                                                                                                                        var v72 int = (^(fi39))
                                                                                                                        _ = v72
                                                                                                                    }
                                                                                                                    _ = v71
                                                                                                                }
                                                                                                            case -263:
                                                                                                                v73 := GenS2{F0: (v22 % ((v22) | 1)), F1: 255, F2: 785}
                                                                                                                _ = v73
                                                                                                            }
                                                                                                        case 127:
                                                                                                            if len(v2) >= 2 {
                                                                                                                v2 = v2[1:len(v2)]
                                                                                                            }
                                                                                                        case -318:
                                                                                                            fmt.Println(v5)
                                                                                                        }
                                                                                                    }
                                                                                                } else {
                                                                                                    var v74 uint32 = 4294967295
                                                                                                    if v10 {
                                                                                                        v4 = (v4 + 42)
                                                                                                        v75 := GenS14{F0: 1, F1: 42}
                                                                                                        v76 := GenS9{F0: 4294967295}
                                                                                                        v77 := GenS14{F0: 9223372036854775807, F1: (^(v22))}
                                                                                                        _ = v75
                                                                                                        _ = v76
                                                                                                        _ = v77
                                                                                                    } else {
                                                                                                        if (!true) {
                                                                                                            var v78 float32 = v3
                                                                                                            var v79 uint32 = 0
                                                                                                            if len(v1) >= 2 {
                                                                                                                v1 = v1[1:len(v1)]
                                                                                                            }
                                                                                                            _ = v78
                                                                                                            _ = v79
                                                                                                        } else {
                                                                                                            if len(v2) > 0 {
                                                                                                                v2[0] = v74
                                                                                                            }
                                                                                                        }
                                                                                                    }
                                                                                                    _ = v74
                                                                                                }
                                                                                                var v80 int8 = 46
                                                                                                switch "a" {
                                                                                                case "hello":
                                                                                                    v3 += ((v3 + ((float32(v27) / (float32(v27) * v3)) - (v3 * v3))) / (v3 * v3))
                                                                                                case "a":
                                                                                                    v4 += -518
                                                                                                default:
                                                                                                    if len(v1) > 0 {
                                                                                                        v1[0] = ("hello" + "")
                                                                                                    }
                                                                                                }
                                                                                                _ = v80
                                                                                            } else {
                                                                                                switch 4294967295 {
                                                                                                case -50:
                                                                                                    var v81 uint16 = (v7 * (v7 ^ (v7 - 42)))
                                                                                                    _ = v81
                                                                                                case -2147483648:
                                                                                                    fmt.Println(fi39)
                                                                                                    var v82 float64 = 0.0
                                                                                                    _ = v82
                                                                                                case -815:
                                                                                                    var v83 int8 = 127
                                                                                                    _ = v83
                                                                                                }
                                                                                            }
                                                                                        } else {
                                                                                            v84 := GenS3{F0: 0, F1: -128, F2: (v7 % (((v7 * v7)) | 1))}
                                                                                            v85 := GenS14{F0: 9223372036854775807, F1: v22}
                                                                                            v7 = v7
                                                                                            _ = v84
                                                                                            _ = v85
                                                                                        }
                                                                                        _ = fi39
                                                                                    }
                                                                                }
                                                                                _ = v27
                                                                            } else {
                                                                                v86 := GenS5{F0: "hello", F1: ((-(v4)) >= (v5 * (v4 * 1)))}
                                                                                v5 = v5
                                                                                _ = v86
                                                                            }
                                                                        }
                                                                        if len(v2) >= 2 {
                                                                            v2 = v2[1:len(v2)]
                                                                        }
                                                                        switch (fi20 / ((v4) | 1)) {
                                                                        case 378:
                                                                            var v87 uint16 = 478
                                                                            _ = v87
                                                                        default:
                                                                            v7 = (v7 * (v7 | (^(v7))))
                                                                        }
                                                                    }
                                                                }
                                                            }
                                                            v3 -= float32(-2.25)
                                                            _ = v22
                                                        }
                                                        var v88 float64 = 0.1
                                                        _ = fi20
                                                        _ = v21
                                                        _ = v88
                                                    }
                                                } else {
                                                    v2 = append(v2, 2147483648)
                                                    if v10 {
                                                        v2 = append(v2, 86)
                                                        var v89 int64 = -1
                                                        _ = v89
                                                    }
                                                }
                                            default:
                                                v7 = 65535
                                            }
                                            for fi90 := 0; fi90 < 7; fi90++ {
                                                if true {
                                                    var v91 float32 = v3
                                                    _ = v91
                                                }
                                                var v92 bool = (42 < -849)
                                                for fi93 := 0; fi93 < 4; fi93++ {
                                                    if ("a" > "hello") {
                                                        switch "hello" {
                                                        case "\n\t":
                                                            v10 = (true && v92)
                                                            var v94 int64 = 0
                                                            var v95 bool = v92
                                                            v92 = (v92 && v92)
                                                            if (false || v95) {
                                                                fmt.Println(v94, v4, v10)
                                                                var v96 uint = 65535
                                                                _ = v96
                                                            } else {
                                                                if len(v2) >= 2 {
                                                                    v2 = v2[1:len(v2)]
                                                                }
                                                                fmt.Println(v3, v95)
                                                                if true {
                                                                    v97 := GenS3{F0: 2147483648, F1: 105, F2: v7}
                                                                    _ = v97
                                                                }
                                                            }
                                                            v98 := GenS9{F0: 59}
                                                            if (v10 && ("" > ("a" + (("hello" + "a") + "")))) {
                                                                var v99 int32 = -931
                                                                if (574 != 65535) {
                                                                    var v100 int32 = (^(v99))
                                                                    var v101 uint16 = (v7 << ((42) & 63))
                                                                    _ = v100
                                                                    _ = v101
                                                                }
                                                                var v102 int32 = (v99 >> ((580) & 63))
                                                                _ = len(v1)
                                                                _ = v102
                                                                _ = v99
                                                            }
                                                            _ = v94
                                                            _ = v95
                                                            _ = v98
                                                        case "foo bar":
                                                            v103 := GenS10{F0: 4294967295, F1: -110}
                                                            _ = v103
                                                        }
                                                        if (42 < -32768) {
                                                            v3 = (v3 / (1e10 / 1.5))
                                                            v104 := GenS15{F0: 898}
                                                            switch (("hello" + (("foo bar" + ("hello" + (("foo bar" + ("" + "hello")) + ""))) + "a")) + "\n\t") {
                                                            case "foo bar":
                                                                var v105 int8 = 42
                                                                _ = v105
                                                            case "hello":
                                                                v92 = v10
                                                            case "":
                                                                if v10 {
                                                                    if v92 {
                                                                        switch fi93 {
                                                                        case -797:
                                                                            v106 := GenS6{F0: 2}
                                                                            _ = v106
                                                                        }
                                                                        var v107 float64 = 1.5
                                                                        if (-1 == -32768) {
                                                                            v4 += (fi93 | (v4 | v4))
                                                                            var v108 int8 = 67
                                                                            v109 := GenS2{F0: 681, F1: 42, F2: -2147483648}
                                                                            _ = v108
                                                                            _ = v109
                                                                        } else {
                                                                            var v110 float32 = v3
                                                                            v111 := GenS16{F0: v5}
                                                                            _ = v110
                                                                            _ = v111
                                                                        }
                                                                        _ = v107
                                                                    } else {
                                                                        v112 := GenS2{F0: 42, F1: 127, F2: 0}
                                                                        v113 := GenS17{F0: 42, F1: -128}
                                                                        fmt.Println(fi93)
                                                                        _ = v112
                                                                        _ = v113
                                                                    }
                                                                } else {
                                                                    if (false && v10) {
                                                                        v114 := GenS5{F0: (("foo bar" + ("" + ("\n\t" + "foo bar"))) + "hello"), F1: v92}
                                                                        var v115 bool = (!true)
                                                                        v115 = v10
                                                                        _ = v114
                                                                        _ = v115
                                                                    }
                                                                    var v116 int8 = 127
                                                                    v117 := GenS18{F0: 1}
                                                                    _ = v116
                                                                    _ = v117
                                                                }
                                                            default:
                                                                fmt.Println(v7, v92, v5)
                                                            }
                                                            _ = v104
                                                        }
                                                    } else {
                                                        var v118 int8 = -92
                                                        _ = v118
                                                    }
                                                    _ = fi93
                                                }
                                                _ = fi90
                                                _ = v92
                                            }
                                            fmt.Println(v10, v10, v5)
                                        } else {
                                            _ = len(v1)
                                        }
                                    case "a":
                                        var v119 int64 = 1
                                        _ = v119
                                    case "":
                                        v120 := GenS6{F0: 0}
                                        _ = v120
                                    }
                                }
                            }
                        }
                    } else {
                        var v121 uint16 = v7
                        var v122 uint = 18446744073709551615
                        _ = v121
                        _ = v122
                    }
                    for fi123 := 0; fi123 < 6; fi123++ {
                        switch "hello" {
                        case "foo bar":
                            var v124 int8 = 100
                            _ = v124
                        case "hello":
                            switch "foo bar" {
                            case "hello":
                                if v10 {
                                    for fi125 := 0; fi125 < 4; fi125++ {
                                        var v126 int16 = -1
                                        if true {
                                            v3 = (v3 + 1.5)
                                            if len(v2) > 0 {
                                                v2[0] = 1
                                            }
                                            v126 = v126
                                        } else {
                                            var v127 uint8 = 1
                                            var v128 int16 = v126
                                            if len(v2) > 0 {
                                                v2[0] = 472
                                            }
                                            _ = v127
                                            _ = v128
                                        }
                                        v7 -= (v7 + (v7 ^ v7))
                                        _ = fi125
                                        _ = v126
                                    }
                                    v4 = v5
                                }
                            default:
                                v4 += -2147483648
                            }
                        case "a":
                            v3 = v3
                        default:
                            for fi129 := 0; fi129 < 14; fi129++ {
                                v10 = v10
                                fmt.Println(v4)
                                var v130 int32 = -912
                                v10 = v10
                                var v131 int16 = 42
                                _ = fi129
                                _ = v130
                                _ = v131
                            }
                        }
                        var v132 uint16 = 451
                        var v133 int8 = 100
                        _ = fi123
                        _ = v132
                        _ = v133
                    }
                }
                v3 -= ((v3 - v3) - v3)
                _ = v10
            }
        }
    case "a":
        var v134 float64 = (float64(v3) + (float64(v3) + -2.25))
        _ = v134
    }
    var v135 int = (v5 >> ((2) & 63))
    if len(v1) > 0 {
        v1[0] = "a"
    }
    var v136 uint16 = v7
    var v137 uint32 = 2
    if ("\n\t" < "\n\t") {
        for fi138 := 0; fi138 < 7; fi138++ {
            var v139 uint16 = v7
            _ = fi138
            _ = v139
        }
    } else {
        if len(v1) > 0 {
            v1[0] = (("foo bar" + "") + "foo bar")
        }
        switch (("a" + ("a" + "")) + "") {
        case "a":
            for fi140 := 0; fi140 < 6; fi140++ {
                var v141 int64 = 9223372036854775807
                _ = fi140
                _ = v141
            }
        default:
            var v142 uint16 = v136
            _ = v142
        }
    }
    v5 = v135
    for fi143 := 0; fi143 < 12; fi143++ {
        v144 := GenS10{F0: 931, F1: (v5 / (((v4 + (-(v4)))) | 1))}
        var v145 int16 = -1
        if ((false && true) && (119 != uint8(v136))) {
            v4 = fi143
        }
        _ = fi143
        _ = v144
        _ = v145
    }
    var v146 float64 = 3.14159
    if (!((v135 < (v5 >> ((1) & 63))) && (false || true))) {
        _ = len(v1)
        switch "foo bar" {
        case "":
            v3 -= v3
        case "hello":
            fmt.Println(v136, v146)
        default:
            var v147 string = ""
            _ = v147
        }
        _ = len(v2)
    } else {
        for fi148 := 0; fi148 < 4; fi148++ {
            if len(v2) > 0 {
                v2[0] = 544
            }
            v149 := GenS11{F0: 325}
            var v150 uint32 = v137
            _ = fi148
            _ = v149
            _ = v150
        }
        var v151 uint32 = v137
        switch "foo bar" {
        case "hello":
            for fi152 := 0; fi152 < 10; fi152++ {
                v153 := GenS19{F0: -593, F1: 255}
                v151 += v151
                fmt.Println(v137)
                _ = fi152
                _ = v153
            }
        case "\n\t":
            fmt.Println(v7)
        default:
            var v154 uint32 = (v137 >> ((255) & 63))
            _ = v154
        }
        _ = v151
    }
    if ((-(v135)) > v135) {
        var v155 int64 = 1
        _ = v155
    } else {
        for fi156 := 0; fi156 < 13; fi156++ {
            v136 -= v7
            _ = fi156
        }
        var v157 uint8 = 42
        _ = v157
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v135:", v135)
    fmt.Println("v136:", v136)
    fmt.Println("v137:", v137)
    fmt.Println("v146:", v146)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v135
    _ = v136
    _ = v137
    _ = v146
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
