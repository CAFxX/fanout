package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8; F2 uint16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint16; F1 string; F2 uint16 }
type GenS4 struct { F0 int64; F1 string }
type GenS5 struct { F0 int8; F1 uint16; F2 uint16 }
type GenS6 struct { F0 uint64; F1 int32 }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 uint16; F1 uint; F2 uint16 }
type GenS9 struct { F0 string; F1 uint32 }
type GenS10 struct { F0 uint8 }
type GenS11 struct { F0 int32; F1 uint32; F2 int32 }
type GenS12 struct { F0 bool; F1 uint32 }
type GenS13 struct { F0 uint8 }
type GenS14 struct { F0 int16; F1 uint64; F2 uint }
type GenS15 struct { F0 string }
type GenS16 struct { F0 uint64; F1 uint }
type GenS17 struct { F0 uint8 }
type GenS18 struct { F0 uint8; F1 string }
type GenS19 struct { F0 uint32; F1 uint }
type GenS20 struct { F0 uint32 }
type GenS21 struct { F0 uint8; F1 uint16 }
type GenS22 struct { F0 int64; F1 int16 }

func main() {
    var v1 int64 = 1
    v2 := make(map[uint32]bool)
    var v3 int16 = 32767
    v4 := make(map[string]int)
    v5 := make(map[string]int)
    v6 := []int{-9223372036854775808, -1, 32768, 32767}
    delete(v4, "\n\t")
    v1 += (v1 * 1)
    v1 -= (v1 - v1)
    v4["hello"] = 42
    switch 127 {
    case 32767:
        if false {
            v3 = v3
            v1 += v1
        }
    }
    var v7 string = "\n\t"
    v7 = v7
    _ = len(v6)
    for fi8 := 0; fi8 < 20; fi8++ {
        for fi9 := 0; fi9 < 19; fi9++ {
            _ = len(v6)
            v1 += (v1 & (v1 >> ((1) & 63)))
            _ = fi9
        }
        switch fi8 {
        case 9223372036854775807:
            switch v7 {
            case "hello":
                for fi10 := 0; fi10 < 1; fi10++ {
                    switch v7 {
                    case "foo bar":
                        switch fi8 {
                        case -1:
                            var v11 int64 = (v1 & v1)
                            _ = v11
                        case -820:
                            switch (fi8 + 2147483647) {
                            case 0:
                                _ = len(v2)
                            default:
                                v7 = (v7 + (v7 + (v7 + v7)))
                            }
                        }
                        var v12 int16 = -32768
                        _ = v12
                    case "a":
                        v13 := GenS1{F0: (fi8 / ((fi8) | 1)), F1: -70, F2: 1}
                        _ = v13
                    case "hello":
                        if (42 == 103) {
                            if ((v7 + "foo bar") >= v7) {
                                if (!true) {
                                    v14 := GenS1{F0: (fi8 / ((fi8) | 1)), F1: 107, F2: 65535}
                                    var v15 uint = 962
                                    _ = v14
                                    _ = v15
                                }
                                if true {
                                    if len(v6) > 0 {
                                        v6[0] = (fi8 * fi10)
                                    }
                                    if ((false || false) && true) {
                                        if false {
                                            fmt.Println(fi10, fi8, v7)
                                        }
                                        _ = len(v5)
                                        switch fi10 {
                                        case -1:
                                            v7 = v7
                                        case 128:
                                            if len(v6) > 0 {
                                                v6[0] = (fi10 & -886)
                                            }
                                        case 32768:
                                            switch fi8 {
                                            case -1:
                                                v16 := GenS1{F0: (fi8 << ((18446744073709551615) & 63)), F1: 42, F2: 0}
                                                _ = v16
                                            }
                                        default:
                                            v17 := GenS1{F0: (fi10 - (fi8 + 4294967295)), F1: 74, F2: 667}
                                            var v18 uint = 1
                                            _ = v17
                                            _ = v18
                                        }
                                    } else {
                                        v2[2147483648] = (v7 <= "a")
                                        if false {
                                            if (true && (127 >= 8)) {
                                                switch (^(fi8)) {
                                                case 42:
                                                    if true {
                                                        v7 = v7
                                                        v1 += (v1 + -490)
                                                        v19 := GenS2{F0: 677}
                                                        _ = v19
                                                    }
                                                case -384:
                                                    v3 += (v3 - (v3 + (v3 / ((v3) | 1))))
                                                default:
                                                    v20 := GenS2{F0: 42}
                                                    _ = v20
                                                }
                                                v3 += v3
                                            } else {
                                                if ((42 >= 399) || (127 < 42)) {
                                                    v21 := GenS2{F0: 4294967295}
                                                    _ = v21
                                                }
                                                v1 += (v1 & (v1 | v1))
                                                v22 := GenS3{F0: 2, F1: ("" + (("foo bar" + ("foo bar" + ("a" + (v7 + v7)))) + v7)), F2: 180}
                                                _ = v22
                                            }
                                            if true {
                                                if false {
                                                    if true {
                                                        v23 := GenS2{F0: 1}
                                                        _ = v23
                                                    }
                                                    v24 := GenS4{F0: (^(v1)), F1: ("a" + v7)}
                                                    _ = v24
                                                } else {
                                                    v1 -= -991
                                                }
                                                switch (v7 + "") {
                                                case "a":
                                                    _ = len(v5)
                                                case "hello":
                                                    v3 = 1000
                                                default:
                                                    v25 := GenS5{F0: 127, F1: 65535, F2: 32767}
                                                    _ = v25
                                                }
                                                var v26 float64 = 0.1
                                                if true {
                                                    fmt.Println(v7, v7)
                                                    var v27 uint8 = 127
                                                    _ = v27
                                                }
                                                _ = v26
                                            } else {
                                                var v28 int8 = 42
                                                delete(v5, v7)
                                                _ = v28
                                            }
                                        }
                                    }
                                    if true {
                                        if true {
                                            switch 2147483647 {
                                            case -1:
                                                v7 = (v7 + v7)
                                            case 314:
                                                switch 128 {
                                                case 1:
                                                    switch (((v7 + v7) + v7) + v7) {
                                                    case "":
                                                        v29 := GenS6{F0: 1, F1: 994}
                                                        _ = v29
                                                    case "hello":
                                                        var v30 int8 = 100
                                                        _ = v30
                                                    default:
                                                        if false {
                                                            var v31 uint8 = 128
                                                            var v32 int = 410
                                                            v31 = v31
                                                            _ = v31
                                                            _ = v32
                                                        }
                                                    }
                                                default:
                                                    v7 = (v7 + "a")
                                                }
                                            }
                                        } else {
                                            v1 = (v1 % (((v1 / ((v1) | 1))) | 1))
                                            v7 = (("hello" + v7) + ((v7 + ((v7 + "\n\t") + v7)) + (v7 + v7)))
                                        }
                                        fmt.Println(v7)
                                    }
                                } else {
                                    fmt.Println(fi8, v7)
                                    if (0 == 65535) {
                                        fmt.Println(v1, v3, v7)
                                        v33 := GenS7{F0: 529}
                                        if len(v6) >= 2 {
                                            v6 = v6[1:len(v6)]
                                        }
                                        _ = v33
                                    }
                                    v1 += v1
                                }
                                var v34 uint64 = 0
                                var v35 uint8 = 127
                                _ = v34
                                _ = v35
                            } else {
                                v1 = (v1 | (v1 / (((v1 - (v1 - 42))) | 1)))
                                delete(v2, 65535)
                            }
                            fmt.Println(v3)
                            var v36 int32 = 574
                            if (!false) {
                                v36 = v36
                                _ = len(v6)
                                var v37 uint64 = 4294967295
                                _ = v37
                            }
                            if (false && (true && (42 < 18446744073709551615))) {
                                v3 = v3
                                var v38 uint32 = 42
                                switch "hello" {
                                case "foo bar":
                                    delete(v5, ((v7 + v7) + (((("" + v7) + "foo bar") + v7) + v7)))
                                default:
                                    switch (("a" + ((v7 + "") + ("\n\t" + (v7 + v7)))) + ((v7 + v7) + v7)) {
                                    case "hello":
                                        var v39 int32 = 0
                                        _ = v39
                                    default:
                                        if (-8 < 65) {
                                            if true {
                                                v7 = "foo bar"
                                                var v40 int8 = -99
                                                if (v3 == (^(v3))) {
                                                    switch (fi10 << ((65535) & 63)) {
                                                    case 32768:
                                                        if false {
                                                            switch fi10 {
                                                            case 452:
                                                                if len(v6) > 0 {
                                                                    v6[0] = (fi8 / ((128) | 1))
                                                                }
                                                            case -9223372036854775808:
                                                                if true {
                                                                    fmt.Println(v3, fi10)
                                                                    if ((22 <= 2) || false) {
                                                                        fmt.Println(v38)
                                                                    } else {
                                                                        switch "foo bar" {
                                                                        case "a":
                                                                            v40 = 42
                                                                        case "":
                                                                            v36 = (^(v36))
                                                                        case "hello":
                                                                            v3 -= -830
                                                                        }
                                                                        var v41 uint32 = 2147483647
                                                                        delete(v4, v7)
                                                                        _ = v41
                                                                    }
                                                                }
                                                            default:
                                                                switch (fi8 + (fi8 << ((1) & 63))) {
                                                                case -1:
                                                                    switch (fi8 + (-(fi10))) {
                                                                    case 32768:
                                                                        if (false && ((v7 + v7) > (((v7 + "foo bar") + (v7 + ((v7 + v7) + ("hello" + "\n\t")))) + (v7 + v7)))) {
                                                                            v40 = (v40 << ((18446744073709551615) & 63))
                                                                        } else {
                                                                            var v42 int8 = 1
                                                                            _ = v42
                                                                        }
                                                                    default:
                                                                        switch fi8 {
                                                                        case 212:
                                                                            var v43 uint32 = (v38 | 2147483648)
                                                                            _ = v43
                                                                        case 574:
                                                                            var v44 uint = 42
                                                                            _ = v44
                                                                        }
                                                                    }
                                                                case 127:
                                                                    v40 -= v40
                                                                default:
                                                                    switch fi8 {
                                                                    case 9223372036854775807:
                                                                        delete(v5, v7)
                                                                        v45 := GenS8{F0: 255, F1: 33, F2: 2}
                                                                        var v46 int32 = -2147483648
                                                                        _ = v45
                                                                        _ = v46
                                                                    case 127:
                                                                        v47 := GenS9{F0: v7, F1: v38}
                                                                        _ = v47
                                                                    case 0:
                                                                        v3 -= (v3 >> ((736) & 63))
                                                                    }
                                                                }
                                                            }
                                                            if (255 == 318) {
                                                                if ((-(v1)) == -1) {
                                                                    v36 = (v36 ^ 42)
                                                                    v48 := GenS10{F0: 255}
                                                                    v40 += (v40 << ((0) & 63))
                                                                    switch 462 {
                                                                    case 32767:
                                                                        var v49 uint64 = 1
                                                                        _ = v49
                                                                    case 9223372036854775807:
                                                                        v36 = -2147483648
                                                                    case 127:
                                                                        switch v7 {
                                                                        case "":
                                                                            if true {
                                                                                if (!true) {
                                                                                    if ((!(2147483647 != v36)) && true) {
                                                                                        fmt.Println(v38, v40)
                                                                                        if len(v6) > 0 {
                                                                                            v6[0] = fi10
                                                                                        }
                                                                                        v50 := v2[v38]
                                                                                        _ = v50
                                                                                    } else {
                                                                                        var v51 uint8 = 42
                                                                                        v52 := GenS11{F0: -1, F1: v38, F2: int32(v51)}
                                                                                        var v53 float32 = (0.0 - 1e10)
                                                                                        var v54 uint16 = 949
                                                                                        v55 := GenS12{F0: (!(741 <= 0)), F1: v38}
                                                                                        _ = v51
                                                                                        _ = v52
                                                                                        _ = v53
                                                                                        _ = v54
                                                                                        _ = v55
                                                                                    }
                                                                                    v56 := GenS13{F0: 2}
                                                                                    _ = v56
                                                                                } else {
                                                                                    v36 += v36
                                                                                    if ((^(fi10)) >= 32768) {
                                                                                        v40 = v40
                                                                                    }
                                                                                }
                                                                                v57 := GenS14{F0: v3, F1: 2, F2: 255}
                                                                                switch (("" + "a") + (v7 + v7)) {
                                                                                case "hello":
                                                                                    v7 = (((("" + ("hello" + (v7 + ("\n\t" + v7)))) + (v7 + v7)) + v7) + "hello")
                                                                                }
                                                                                _ = v57
                                                                            } else {
                                                                                var v58 string = (v7 + v7)
                                                                                _ = v58
                                                                            }
                                                                        default:
                                                                            if false {
                                                                                var v59 uint8 = 128
                                                                                if false {
                                                                                    if false {
                                                                                        fmt.Println(v38)
                                                                                        if false {
                                                                                            var v60 float32 = float32(-2.25)
                                                                                            var v61 uint16 = 255
                                                                                            if (!false) {
                                                                                                var v62 float64 = float64(0.1)
                                                                                                _ = v62
                                                                                            } else {
                                                                                                v63 := GenS1{F0: fi10, F1: (v40 - (v40 << ((564) & 63))), F2: 32767}
                                                                                                if (v7 >= (v7 + v7)) {
                                                                                                    var v64 string = (("foo bar" + v7) + "\n\t")
                                                                                                    _ = v64
                                                                                                }
                                                                                                _ = v63
                                                                                            }
                                                                                            var v65 uint64 = 0
                                                                                            v66 := GenS4{F0: (v1 & (v1 & 82)), F1: (v7 + v7)}
                                                                                            _ = v60
                                                                                            _ = v61
                                                                                            _ = v65
                                                                                            _ = v66
                                                                                        }
                                                                                    } else {
                                                                                        if false {
                                                                                            var v67 uint8 = (v59 >> ((18446744073709551615) & 63))
                                                                                            v3 = v3
                                                                                            v68 := GenS13{F0: v67}
                                                                                            var v69 uint8 = (^(v67))
                                                                                            if true {
                                                                                                if len(v6) > 0 {
                                                                                                    v6[0] = 32768
                                                                                                }
                                                                                                v69 = (v59 * (v69 - v69))
                                                                                                if true {
                                                                                                    if true {
                                                                                                        fmt.Println(v69, v7, v67)
                                                                                                    }
                                                                                                    v70 := GenS14{F0: (v3 - int16(v69)), F1: 350, F2: 18446744073709551615}
                                                                                                    var v71 float32 = 0.1
                                                                                                    _ = v70
                                                                                                    _ = v71
                                                                                                } else {
                                                                                                    v72 := GenS15{F0: "a"}
                                                                                                    v73 := GenS10{F0: (v69 - 127)}
                                                                                                    v3 = (v3 + v3)
                                                                                                    _ = v72
                                                                                                    _ = v73
                                                                                                }
                                                                                            } else {
                                                                                                switch fi10 {
                                                                                                case 32767:
                                                                                                    switch ("\n\t" + (v7 + "hello")) {
                                                                                                    case "":
                                                                                                        _ = len(v5)
                                                                                                    case "\n\t":
                                                                                                        if (!true) {
                                                                                                            if false {
                                                                                                                if true {
                                                                                                                    var v74 string = ""
                                                                                                                    _ = v74
                                                                                                                } else {
                                                                                                                    delete(v2, v38)
                                                                                                                    var v75 float32 = 3.14159
                                                                                                                    _ = v75
                                                                                                                }
                                                                                                                var v76 int = (-(fi10))
                                                                                                                if false {
                                                                                                                    if (false || true) {
                                                                                                                        switch (fi8 & -410) {
                                                                                                                        case 298:
                                                                                                                            v76 -= (v76 * fi8)
                                                                                                                        case -786:
                                                                                                                            fmt.Println(v76, v38, v36)
                                                                                                                        case 2147483647:
                                                                                                                            if ((0 == fi10) || false) {
                                                                                                                                var v77 int64 = v1
                                                                                                                                _ = v77
                                                                                                                            }
                                                                                                                        }
                                                                                                                        _ = len(v5)
                                                                                                                        if (int(v36) < v76) {
                                                                                                                            v40 = 127
                                                                                                                        } else {
                                                                                                                            switch 267 {
                                                                                                                            case 128:
                                                                                                                                var v78 string = v7
                                                                                                                                _ = v78
                                                                                                                            case 127:
                                                                                                                                if true {
                                                                                                                                    if (!true) {
                                                                                                                                        var v79 uint8 = (v69 | v59)
                                                                                                                                        v76 -= (fi10 + (fi8 | v76))
                                                                                                                                        var v80 uint32 = (v38 & (^(v38)))
                                                                                                                                        _ = v79
                                                                                                                                        _ = v80
                                                                                                                                    }
                                                                                                                                    v67 -= v67
                                                                                                                                } else {
                                                                                                                                    if false {
                                                                                                                                        if false {
                                                                                                                                            var v81 bool = false
                                                                                                                                            _ = v81
                                                                                                                                        } else {
                                                                                                                                            var v82 int32 = v36
                                                                                                                                            var v83 uint64 = 4294967295
                                                                                                                                            var v84 float64 = -2.25
                                                                                                                                            delete(v5, "foo bar")
                                                                                                                                            v84 = v84
                                                                                                                                            _ = v82
                                                                                                                                            _ = v83
                                                                                                                                            _ = v84
                                                                                                                                        }
                                                                                                                                        if true {
                                                                                                                                            if true {
                                                                                                                                                v7 = v7
                                                                                                                                                if (((v7 + v7) > "hello") && (!(true && false))) {
                                                                                                                                                    var v85 float64 = 0.0
                                                                                                                                                    v7 = "\n\t"
                                                                                                                                                    _ = v85
                                                                                                                                                }
                                                                                                                                            }
                                                                                                                                            fmt.Println(v3, v76)
                                                                                                                                            if true {
                                                                                                                                                if (0 > 0) {
                                                                                                                                                    v86 := GenS10{F0: (v67 / (((v69 >> ((65535) & 63))) | 1))}
                                                                                                                                                    _ = v86
                                                                                                                                                } else {
                                                                                                                                                    v1 = v1
                                                                                                                                                    delete(v4, v7)
                                                                                                                                                }
                                                                                                                                                v7 = ""
                                                                                                                                                var v87 int32 = int32(v38)
                                                                                                                                                _ = v87
                                                                                                                                            } else {
                                                                                                                                                v7 = "foo bar"
                                                                                                                                                switch (v76 % ((fi8) | 1)) {
                                                                                                                                                case 32768:
                                                                                                                                                    var v88 uint32 = v38
                                                                                                                                                    _ = v88
                                                                                                                                                }
                                                                                                                                            }
                                                                                                                                        } else {
                                                                                                                                            fmt.Println(fi8)
                                                                                                                                        }
                                                                                                                                        var v89 int8 = (v40 >> ((841) & 63))
                                                                                                                                        v38 -= v38
                                                                                                                                        v90 := GenS10{F0: 255}
                                                                                                                                        delete(v2, (v38 >> ((90) & 63)))
                                                                                                                                        v67 = v59
                                                                                                                                        v89 += v89
                                                                                                                                        _ = v89
                                                                                                                                        _ = v90
                                                                                                                                    } else {
                                                                                                                                        switch "foo bar" {
                                                                                                                                        case "":
                                                                                                                                            v67 -= (v69 * (v67 << ((819) & 63)))
                                                                                                                                        case "foo bar":
                                                                                                                                            fmt.Println(v40, v3)
                                                                                                                                        case "\n\t":
                                                                                                                                            var v91 string = v7
                                                                                                                                            switch "\n\t" {
                                                                                                                                            case "a":
                                                                                                                                                if (false || false) {
                                                                                                                                                    var v92 int64 = int64(v67)
                                                                                                                                                    fmt.Println(v36, v69, v3)
                                                                                                                                                    _ = v92
                                                                                                                                                } else {
                                                                                                                                                    fmt.Println(v69, v76)
                                                                                                                                                    v3 = (v3 << ((4294967295) & 63))
                                                                                                                                                }
                                                                                                                                            case "foo bar":
                                                                                                                                                v1 = 9223372036854775807
                                                                                                                                            case "\n\t":
                                                                                                                                                v67 -= v59
                                                                                                                                                var v93 int8 = (v40 << ((2) & 63))
                                                                                                                                                _ = v93
                                                                                                                                            }
                                                                                                                                            v94 := v5[((v7 + v7) + v7)]
                                                                                                                                            _ = v91
                                                                                                                                            _ = v94
                                                                                                                                        default:
                                                                                                                                            v95 := GenS7{F0: 255}
                                                                                                                                            _ = v95
                                                                                                                                        }
                                                                                                                                    }
                                                                                                                                    if (313 > 1) {
                                                                                                                                        switch ("a" + (((v7 + (("hello" + (v7 + v7)) + ((v7 + v7) + v7))) + (v7 + ("" + (("hello" + v7) + v7)))) + ((v7 + "a") + (("\n\t" + (((v7 + "hello") + "hello") + "a")) + (("foo bar" + (((v7 + v7) + ((v7 + (v7 + "\n\t")) + v7)) + v7)) + ("a" + (("foo bar" + ("\n\t" + "")) + ("hello" + (v7 + "a"))))))))) {
                                                                                                                                        case "\n\t":
                                                                                                                                            delete(v4, "\n\t")
                                                                                                                                        case "a":
                                                                                                                                            v6 = append(v6, (v76 << ((uint(v38)) & 63)))
                                                                                                                                        default:
                                                                                                                                            if false {
                                                                                                                                                var v96 string = "hello"
                                                                                                                                                delete(v5, v7)
                                                                                                                                                _ = v96
                                                                                                                                            } else {
                                                                                                                                                var v97 uint8 = v67
                                                                                                                                                if (false && false) {
                                                                                                                                                    fmt.Println(v7)
                                                                                                                                                    switch v76 {
                                                                                                                                                    case 32768:
                                                                                                                                                        if len(v6) >= 2 {
                                                                                                                                                            v6 = v6[1:len(v6)]
                                                                                                                                                        }
                                                                                                                                                        v98 := GenS16{F0: 0, F1: 18446744073709551615}
                                                                                                                                                        if true {
                                                                                                                                                            v99 := GenS3{F0: 32767, F1: ("\n\t" + v7), F2: 2}
                                                                                                                                                            _ = v99
                                                                                                                                                        }
                                                                                                                                                        _ = v98
                                                                                                                                                    case 2147483647:
                                                                                                                                                        if true {
                                                                                                                                                            fmt.Println(v7)
                                                                                                                                                        } else {
                                                                                                                                                            if (1 > (v3 ^ int16(v69))) {
                                                                                                                                                                switch ("\n\t" + v7) {
                                                                                                                                                                case "\n\t":
                                                                                                                                                                    if true {
                                                                                                                                                                        var v100 uint32 = 1
                                                                                                                                                                        _ = v100
                                                                                                                                                                    }
                                                                                                                                                                default:
                                                                                                                                                                    var v101 int64 = v1
                                                                                                                                                                    _ = v101
                                                                                                                                                                }
                                                                                                                                                            } else {
                                                                                                                                                                if true {
                                                                                                                                                                    v36 -= 42
                                                                                                                                                                    v36 += v36
                                                                                                                                                                }
                                                                                                                                                                if ((-(v3)) < v3) {
                                                                                                                                                                    v40 = (v40 ^ v40)
                                                                                                                                                                    v102 := GenS7{F0: 65535}
                                                                                                                                                                    _ = v102
                                                                                                                                                                }
                                                                                                                                                            }
                                                                                                                                                            var v103 float64 = (-2.25 + 3.14159)
                                                                                                                                                            _ = v103
                                                                                                                                                        }
                                                                                                                                                    case 246:
                                                                                                                                                        fmt.Println(fi8, v40)
                                                                                                                                                    }
                                                                                                                                                }
                                                                                                                                                _ = v97
                                                                                                                                            }
                                                                                                                                        }
                                                                                                                                        v104 := GenS4{F0: v1, F1: v7}
                                                                                                                                        _ = v104
                                                                                                                                    }
                                                                                                                                    var v105 uint32 = v38
                                                                                                                                    _ = v105
                                                                                                                                }
                                                                                                                            case 0:
                                                                                                                                var v106 float64 = 1e10
                                                                                                                                _ = v106
                                                                                                                            default:
                                                                                                                                if false {
                                                                                                                                    switch 32767 {
                                                                                                                                    case 747:
                                                                                                                                        var v107 int32 = (-(v36))
                                                                                                                                        _ = v107
                                                                                                                                    }
                                                                                                                                    var v108 int8 = 42
                                                                                                                                    _ = v108
                                                                                                                                } else {
                                                                                                                                    var v109 float32 = 0.0
                                                                                                                                    v40 -= v40
                                                                                                                                    var v110 int = -761
                                                                                                                                    switch ("a" + "foo bar") {
                                                                                                                                    case "foo bar":
                                                                                                                                        var v111 float32 = v109
                                                                                                                                        _ = v111
                                                                                                                                    }
                                                                                                                                    _ = v109
                                                                                                                                    _ = v110
                                                                                                                                }
                                                                                                                            }
                                                                                                                            fmt.Println(fi8)
                                                                                                                        }
                                                                                                                    } else {
                                                                                                                        if true {
                                                                                                                            fmt.Println(v40, v69)
                                                                                                                        }
                                                                                                                        var v112 bool = ((0 == v67) && ((v40 * 127) == 1))
                                                                                                                        var v113 float64 = 0.1
                                                                                                                        fmt.Println(v112, v69, v36)
                                                                                                                        var v114 int = v76
                                                                                                                        _ = v112
                                                                                                                        _ = v113
                                                                                                                        _ = v114
                                                                                                                    }
                                                                                                                    v115 := GenS17{F0: (v69 << ((535) & 63))}
                                                                                                                    _ = v115
                                                                                                                } else {
                                                                                                                    v7 = ""
                                                                                                                    var v116 float32 = 1e10
                                                                                                                    var v117 float64 = (1.5 + (1.5 * 1.5))
                                                                                                                    _ = v116
                                                                                                                    _ = v117
                                                                                                                }
                                                                                                                _ = v76
                                                                                                            }
                                                                                                        }
                                                                                                    default:
                                                                                                        if (v3 >= (v3 << ((638) & 63))) {
                                                                                                            if (false || (v36 == int32(v3))) {
                                                                                                                if false {
                                                                                                                    if true {
                                                                                                                        v118 := GenS18{F0: (v59 * (v69 & v67)), F1: v7}
                                                                                                                        v59 += (^(v67))
                                                                                                                        _ = v118
                                                                                                                    }
                                                                                                                    v3 += 867
                                                                                                                }
                                                                                                                var v119 string = (v7 + "")
                                                                                                                var v120 int8 = v40
                                                                                                                _ = v119
                                                                                                                _ = v120
                                                                                                            }
                                                                                                            if true {
                                                                                                                var v121 int = 907
                                                                                                                v40 = (v40 % ((v40) | 1))
                                                                                                                v6 = append(v6, (fi8 << ((0) & 63)))
                                                                                                                _ = v121
                                                                                                            }
                                                                                                            v4[v7] = -885
                                                                                                            v38 = 4294967295
                                                                                                            v59 -= (v67 ^ 168)
                                                                                                            var v122 uint32 = 685
                                                                                                            switch (fi8 | (fi10 & 653)) {
                                                                                                            case 128:
                                                                                                                var v123 int16 = (^(v3))
                                                                                                                _ = v123
                                                                                                            }
                                                                                                            if true {
                                                                                                                v4["hello"] = int(v40)
                                                                                                                v67 = (^(v69))
                                                                                                                v6 = append(v6, fi8)
                                                                                                            } else {
                                                                                                                v67 = (^(v67))
                                                                                                            }
                                                                                                            _ = v122
                                                                                                        }
                                                                                                    }
                                                                                                case -2147483648:
                                                                                                    var v124 uint16 = 32767
                                                                                                    _ = v124
                                                                                                }
                                                                                                v69 -= v59
                                                                                                fmt.Println(fi8)
                                                                                            }
                                                                                            _ = v67
                                                                                            _ = v68
                                                                                            _ = v69
                                                                                        }
                                                                                    }
                                                                                    var v125 uint32 = 2147483648
                                                                                    var v126 float32 = float32(0.1)
                                                                                    _ = v125
                                                                                    _ = v126
                                                                                }
                                                                                _ = v59
                                                                            } else {
                                                                                v127 := GenS18{F0: 3, F1: v7}
                                                                                _ = v127
                                                                            }
                                                                        }
                                                                        _ = len(v2)
                                                                        var v128 uint16 = 1
                                                                        _ = v128
                                                                    }
                                                                    _ = v48
                                                                }
                                                                v3 = 0
                                                            }
                                                            switch fi8 {
                                                            case 755:
                                                                v2[v38] = false
                                                            }
                                                        } else {
                                                            var v129 string = ("a" + ((((v7 + v7) + v7) + v7) + ((v7 + v7) + "\n\t")))
                                                            _ = v129
                                                        }
                                                    }
                                                    v130 := GenS15{F0: "\n\t"}
                                                    _ = len(v6)
                                                    _ = v130
                                                } else {
                                                    var v131 float64 = 1e10
                                                    fmt.Println(fi10, v131, fi10)
                                                    _ = v131
                                                }
                                                _ = v40
                                            }
                                            v3 -= (v3 + (v3 - -610))
                                            v1 -= v1
                                            if len(v6) > 0 {
                                                v6[0] = (fi8 << ((396) & 63))
                                            }
                                            if true {
                                                v4[""] = 4294967295
                                                v6 = append(v6, (fi10 << ((4294967295) & 63)))
                                            } else {
                                                v132 := GenS19{F0: (v38 | (v38 & v38)), F1: 597}
                                                _ = v132
                                            }
                                        } else {
                                            v133 := GenS8{F0: 255, F1: 42, F2: 2}
                                            var v134 bool = false
                                            v134 = v134
                                            _ = v133
                                            _ = v134
                                        }
                                        var v135 uint16 = 65535
                                        switch (fi10 - 133) {
                                        case 0:
                                            var v136 int64 = (v1 >> ((1) & 63))
                                            _ = v136
                                        case 831:
                                            switch (fi8 / ((fi10) | 1)) {
                                            case -940:
                                                v135 = v135
                                            case 0:
                                                fmt.Println(fi10)
                                            case -1:
                                                switch "\n\t" {
                                                case "\n\t":
                                                    _ = len(v6)
                                                }
                                            }
                                        default:
                                            v36 += int32(v3)
                                        }
                                        var v137 float64 = 0.1
                                        v138 := GenS17{F0: 2}
                                        v1 = v1
                                        var v139 int8 = 0
                                        _ = v135
                                        _ = v137
                                        _ = v138
                                        _ = v139
                                    }
                                    v140 := GenS14{F0: v3, F1: 0, F2: 723}
                                    if len(v6) >= 2 {
                                        v6 = v6[1:len(v6)]
                                    }
                                    _ = v140
                                }
                                _ = v38
                            } else {
                                if (0 <= 1) {
                                    v7 = ((v7 + ("" + ((v7 + ((v7 + "hello") + "")) + v7))) + (v7 + v7))
                                    switch (v7 + v7) {
                                    case "foo bar":
                                        v1 -= (-(v1))
                                    case "":
                                        v141 := GenS20{F0: 216}
                                        _ = v141
                                    default:
                                        v36 = int32(fi8)
                                    }
                                }
                            }
                            var v142 int = fi8
                            _ = v142
                            _ = v36
                        } else {
                            v1 += 1
                            var v143 int32 = 0
                            v144 := GenS4{F0: 581, F1: ""}
                            _ = v143
                            _ = v144
                        }
                    default:
                        if (128 != 127) {
                            if true {
                                switch (fi8 - fi10) {
                                case 94:
                                    var v145 uint = 2
                                    _ = v145
                                case 0:
                                    var v146 int64 = 0
                                    _ = v146
                                }
                                var v147 uint = 42
                                v148 := GenS7{F0: 0}
                                _ = v147
                                _ = v148
                            } else {
                                if (191 <= 2) {
                                    v1 += v1
                                    var v149 uint8 = 128
                                    _ = v149
                                }
                                v3 += (-(v3))
                                fmt.Println(v1, v7, fi8)
                                if (((fi8 < fi8) || false) || true) {
                                    v150 := GenS11{F0: 0, F1: 4294967295, F2: 1}
                                    _ = v150
                                }
                            }
                        }
                    }
                    var v151 bool = true
                    _ = fi10
                    _ = v151
                }
            case "a":
                delete(v2, 4294967295)
            }
        case 42:
            if (4294967295 > 4294967295) {
                if (1 == 0) {
                    fmt.Println(fi8)
                } else {
                    v3 -= v3
                    fmt.Println(fi8, v1)
                }
                var v152 string = ((v7 + v7) + "hello")
                var v153 bool = ((0.1 - 0.1) != 0.0)
                _ = v152
                _ = v153
            } else {
                var v154 uint32 = 2147483648
                v155 := GenS10{F0: 42}
                switch (fi8 & (fi8 / ((fi8) | 1))) {
                case 42:
                    v3 += (v3 % ((int16(fi8)) | 1))
                case -1:
                    v1 += (-(v1))
                case -891:
                    if false {
                        var v156 int32 = 224
                        for fi157 := 0; fi157 < 16; fi157++ {
                            v156 = -646
                            _ = fi157
                        }
                        _ = v156
                    }
                }
                _ = v154
                _ = v155
            }
        case -622:
            fmt.Println(v7, v1)
        }
        v1 -= (v1 ^ -1)
        _ = fi8
    }
    if false {
        if (true && (127 > 1)) {
            fmt.Println(v7, v7, v1)
        } else {
            v158 := GenS21{F0: 128, F1: 32767}
            fmt.Println(v7, v7)
            v159 := GenS18{F0: 199, F1: ""}
            _ = v158
            _ = v159
        }
        var v160 string = v7
        _ = v160
    } else {
        v7 = v7
        if (9223372036854775807 != 128) {
            for fi161 := 0; fi161 < 4; fi161++ {
                switch fi161 {
                case 127:
                    fmt.Println(v7)
                default:
                    var v162 int32 = 48
                    _ = v162
                }
                v163 := GenS22{F0: (v1 + v1), F1: (-(v3))}
                _ = fi161
                _ = v163
            }
        }
    }
    v164 := v4["foo bar"]
    _ = len(v6)
    v164 = v164
    fmt.Println("v1:", v1)
    fmt.Println("v164:", v164)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v164
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
