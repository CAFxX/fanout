package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string }
type GenS2 struct { F0 int64; F1 uint8; F2 int64 }
type GenS3 struct { F0 uint32; F1 int32 }
type GenS4 struct { F0 int16; F1 int64 }
type GenS5 struct { F0 int64; F1 uint8; F2 uint }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 int16; F1 bool }
type GenS8 struct { F0 int16; F1 uint64; F2 uint8 }
type GenS9 struct { F0 uint; F1 int32; F2 int }
type GenS10 struct { F0 string; F1 int16 }
type GenS11 struct { F0 int64; F1 uint }
type GenS12 struct { F0 bool; F1 int64; F2 uint16 }
type GenS13 struct { F0 int32; F1 int32; F2 bool }
type GenS14 struct { F0 uint64 }
type GenS15 struct { F0 string }
type GenS16 struct { F0 bool }
type GenS17 struct { F0 uint8; F1 uint16; F2 uint64 }
type GenS18 struct { F0 uint8; F1 uint64 }
type GenS19 struct { F0 string }
type GenS20 struct { F0 int8; F1 uint; F2 int }
type GenS21 struct { F0 uint32; F1 int8 }
type GenS22 struct { F0 string; F1 uint8 }
type GenS23 struct { F0 uint8 }
type GenS24 struct { F0 uint8; F1 int32 }
type GenS25 struct { F0 uint64; F1 int; F2 int64 }
type GenS26 struct { F0 int }
type GenS27 struct { F0 int8; F1 uint32 }
type GenS28 struct { F0 uint64; F1 int8 }
type GenS29 struct { F0 uint16; F1 uint16 }
type GenS30 struct { F0 uint64 }
type GenS31 struct { F0 bool; F1 uint }
type GenS32 struct { F0 uint32 }
type GenS33 struct { F0 uint64; F1 bool; F2 int32 }
type GenS34 struct { F0 uint64 }
type GenS35 struct { F0 int16; F1 uint64 }
type GenS36 struct { F0 string; F1 int; F2 int }
type GenS37 struct { F0 bool }
type GenS38 struct { F0 string; F1 uint }
type GenS39 struct { F0 int32; F1 int; F2 string }
type GenS40 struct { F0 int16; F1 bool; F2 bool }
type GenS41 struct { F0 uint }
type GenS42 struct { F0 int32; F1 uint32 }
type GenS43 struct { F0 uint32; F1 int16 }
type GenS44 struct { F0 uint; F1 bool; F2 string }

func main() {
    v1 := make(map[uint32]int)
    var v2 uint8 = 1
    var v3 string = "a"
    var v4 bool = true
    fmt.Println(v3)
    if (!true) {
        switch (v3 + v3) {
        case "":
            if v4 {
                v3 = ("\n\t" + v3)
            }
        case "foo bar":
            switch v3 {
            case "foo bar":
                var v5 uint16 = 255
                _ = v5
            }
        case "hello":
            if (v4 || v4) {
                v4 = ((255 != 65535) && v4)
                v6 := []int{-1, 4294967295}
                if false {
                    if v4 {
                        v7 := GenS1{F0: (9223372036854775807 < 0), F1: v3}
                        _ = v7
                    } else {
                        for fi8 := 0; fi8 < 19; fi8++ {
                            if (!(42 < 42)) {
                                var v9 float64 = float64(1e10)
                                v1[2147483648] = 0
                                _ = v9
                            } else {
                                for fi10 := 0; fi10 < 7; fi10++ {
                                    if v4 {
                                        fmt.Println(v3)
                                        if (false || v4) {
                                            v4 = (!v4)
                                            fmt.Println(fi10, v2, fi10)
                                        }
                                        var v11 uint16 = 2
                                        _ = v11
                                    } else {
                                        if (100 >= 0) {
                                            v12 := GenS2{F0: -61, F1: 129, F2: -9223372036854775808}
                                            var v13 uint16 = 42
                                            _ = v12
                                            _ = v13
                                        } else {
                                            var v14 uint16 = 1
                                            switch "foo bar" {
                                            case "hello":
                                                fmt.Println(v3, fi10, v2)
                                            case "foo bar":
                                                v15 := v1[441]
                                                _ = v15
                                            default:
                                                v2 = v2
                                            }
                                            _ = v14
                                        }
                                    }
                                    if v4 {
                                        if (!(0 <= 2)) {
                                            var v16 uint64 = 0
                                            switch (fi10 ^ (-(fi8))) {
                                            case -2147483648:
                                                if v4 {
                                                    switch int(v16) {
                                                    case 0:
                                                        fmt.Println(v16, v4)
                                                        _ = len(v1)
                                                        if v4 {
                                                            _ = len(v6)
                                                        } else {
                                                            fmt.Println(v4)
                                                            _ = len(v1)
                                                            v16 -= (v16 & (^(v16)))
                                                            fmt.Println(fi8)
                                                        }
                                                    default:
                                                        delete(v1, 65535)
                                                    }
                                                }
                                            case -373:
                                                fmt.Println(v2)
                                                var v17 bool = (((1.5 - 3.14159) + 0.0) > (1e10 - ((1.5 * (float64(float32(((((1.5 + (1.5 - (1.5 / 1.5))) / float64(-2.25)) + 0.1) * 3.14159))) / float64((0.0 + 1.5)))) / 1.5)))
                                                _ = v17
                                            case -1:
                                                fmt.Println(v16)
                                                v18 := GenS3{F0: 4294967295, F1: -2147483648}
                                                delete(v1, 2)
                                                _ = v18
                                            default:
                                                switch (^(fi8)) {
                                                case 0:
                                                    var v19 uint64 = 32
                                                    _ = v19
                                                case 128:
                                                    switch 32767 {
                                                    case -9223372036854775808:
                                                        var v20 string = "a"
                                                        _ = v20
                                                    case -2147483648:
                                                        delete(v1, 1)
                                                    }
                                                }
                                                switch -288 {
                                                case 32768:
                                                    v21 := GenS1{F0: v4, F1: v3}
                                                    _ = v21
                                                case -462:
                                                    var v22 float64 = (1e10 * 0.1)
                                                    _ = v22
                                                }
                                                var v23 int16 = 0
                                                v4 = ((!v4) && ((-2.25 - -2.25) < float32(3.14159)))
                                                _ = v23
                                            }
                                            _ = v16
                                        } else {
                                            if ((0.0 / -2.25) == 0.1) {
                                                v24 := GenS4{F0: 1000, F1: 9223372036854775807}
                                                if ((1.5 * (1e10 / -2.25)) == 3.14159) {
                                                    v4 = false
                                                    switch fi8 {
                                                    case 1:
                                                        fmt.Println(v4, fi10)
                                                    case -9223372036854775808:
                                                        v25 := GenS2{F0: -143, F1: v2, F2: 1}
                                                        _ = v25
                                                    }
                                                    var v26 float32 = 1.5
                                                    v27 := GenS2{F0: -1, F1: (v2 >> ((1) & 63)), F2: 0}
                                                    _ = v26
                                                    _ = v27
                                                }
                                                _ = v24
                                            }
                                        }
                                    } else {
                                        switch "\n\t" {
                                        case "":
                                            switch (fi8 ^ fi8) {
                                            case -9223372036854775808:
                                                var v28 float32 = -2.25
                                                fmt.Println(fi10, v4)
                                                if true {
                                                    v29 := GenS2{F0: 9223372036854775807, F1: v2, F2: -9223372036854775808}
                                                    v6 = append(v6, fi10)
                                                    _ = v29
                                                }
                                                fmt.Println(fi10, fi8)
                                                _ = v28
                                            }
                                        case "foo bar":
                                            v30 := GenS5{F0: 1, F1: v2, F2: 502}
                                            switch "\n\t" {
                                            case "foo bar":
                                                if len(v6) > 0 {
                                                    v6[0] = fi8
                                                }
                                            }
                                            if true {
                                                if len(v6) > 0 {
                                                    v6[0] = fi10
                                                }
                                                switch ((v3 + (v3 + "hello")) + (v3 + v3)) {
                                                case "":
                                                    v31 := GenS2{F0: 1, F1: (v2 & 42), F2: 42}
                                                    _ = v31
                                                case "\n\t":
                                                    switch "" {
                                                    case "foo bar":
                                                        if (42 >= (v2 & v2)) {
                                                            v3 = v3
                                                            v4 = v4
                                                            var v32 uint64 = 1
                                                            fmt.Println(v2)
                                                            v33 := GenS6{F0: 2147483647}
                                                            _ = v32
                                                            _ = v33
                                                        } else {
                                                            var v34 int32 = 42
                                                            fmt.Println(v4, v34)
                                                            switch fi8 {
                                                            case 969:
                                                                v3 = v3
                                                            default:
                                                                if (-2.25 > 3.14159) {
                                                                    if false {
                                                                        if len(v6) > 0 {
                                                                            v6[0] = fi10
                                                                        }
                                                                        if (v4 && true) {
                                                                            v34 = (v34 % (((v34 * v34)) | 1))
                                                                            var v35 bool = ((1.5 > 1.5) && v4)
                                                                            _ = v35
                                                                        } else {
                                                                            v34 -= (v34 | v34)
                                                                            var v36 uint8 = (v2 + v2)
                                                                            _ = v36
                                                                        }
                                                                        switch ("hello" + (v3 + v3)) {
                                                                        case "foo bar":
                                                                            v37 := GenS7{F0: -32768, F1: false}
                                                                            _ = v37
                                                                        }
                                                                    } else {
                                                                        switch fi10 {
                                                                        case 32767:
                                                                            switch (fi10 + int(v34)) {
                                                                            case 560:
                                                                                v38 := GenS3{F0: 42, F1: v34}
                                                                                _ = v38
                                                                            case -246:
                                                                                var v39 uint = 2
                                                                                _ = v39
                                                                            case 0:
                                                                                var v40 uint = 233
                                                                                _ = v40
                                                                            }
                                                                        case 127:
                                                                            fmt.Println(v3, v4, v2)
                                                                        case 32768:
                                                                            switch v3 {
                                                                            case "":
                                                                                var v41 uint = 65535
                                                                                _ = v41
                                                                            case "foo bar":
                                                                                var v42 int = fi8
                                                                                _ = v42
                                                                            case "a":
                                                                                v4 = ("\n\t" >= (v3 + v3))
                                                                            default:
                                                                                switch v3 {
                                                                                case "hello":
                                                                                    fmt.Println(v2, fi10, v3)
                                                                                case "\n\t":
                                                                                    var v43 uint16 = 255
                                                                                    _ = v43
                                                                                case "":
                                                                                    if v4 {
                                                                                        switch (-(fi8)) {
                                                                                        case 19:
                                                                                            if (255 >= 52) {
                                                                                                switch (fi8 / (((fi10 << ((2) & 63))) | 1)) {
                                                                                                case 9223372036854775807:
                                                                                                    var v44 bool = (1e10 == 3.14159)
                                                                                                    _ = v44
                                                                                                case -2147483648:
                                                                                                    if (!(2147483648 != 2)) {
                                                                                                        if (!((32767 == -32768) && (v3 <= "a"))) {
                                                                                                            v3 = v3
                                                                                                            var v45 int8 = 127
                                                                                                            _ = v45
                                                                                                        } else {
                                                                                                            var v46 int8 = 36
                                                                                                            v47 := GenS8{F0: int16(fi8), F1: 2, F2: (v2 | (^(v2)))}
                                                                                                            fmt.Println(v34, v3, v3)
                                                                                                            _ = v46
                                                                                                            _ = v47
                                                                                                        }
                                                                                                        v48 := GenS8{F0: 678, F1: 18446744073709551615, F2: v2}
                                                                                                        _ = v48
                                                                                                    } else {
                                                                                                        fmt.Println(v4, v3, v3)
                                                                                                        switch v3 {
                                                                                                        case "\n\t":
                                                                                                            v49 := GenS8{F0: -32768, F1: 18446744073709551615, F2: 1}
                                                                                                            _ = v49
                                                                                                        case "hello":
                                                                                                            var v50 bool = ((((42 >= (^(v2))) && v4) && v4) && v4)
                                                                                                            _ = v50
                                                                                                        default:
                                                                                                            var v51 float64 = 1e10
                                                                                                            _ = v51
                                                                                                        }
                                                                                                        switch "" {
                                                                                                        case "\n\t":
                                                                                                            var v52 uint32 = 378
                                                                                                            _ = v52
                                                                                                        default:
                                                                                                            if v4 {
                                                                                                                switch "" {
                                                                                                                case "hello":
                                                                                                                    switch 4294967295 {
                                                                                                                    case 2147483647:
                                                                                                                        if false {
                                                                                                                            _ = len(v1)
                                                                                                                            var v53 int = fi10
                                                                                                                            _ = v53
                                                                                                                        } else {
                                                                                                                            fmt.Println(v4, fi8)
                                                                                                                            var v54 int16 = 32767
                                                                                                                            _ = v54
                                                                                                                        }
                                                                                                                    default:
                                                                                                                        var v55 float64 = -2.25
                                                                                                                        v4 = v4
                                                                                                                        var v56 uint8 = v2
                                                                                                                        _ = v55
                                                                                                                        _ = v56
                                                                                                                    }
                                                                                                                case "\n\t":
                                                                                                                    var v57 float64 = 0.0
                                                                                                                    _ = v57
                                                                                                                case "foo bar":
                                                                                                                    v58 := GenS6{F0: (v34 ^ 0)}
                                                                                                                    _ = v58
                                                                                                                }
                                                                                                            } else {
                                                                                                                v3 = v3
                                                                                                            }
                                                                                                            switch ((v3 + (v3 + ((v3 + v3) + v3))) + ((("" + (((("\n\t" + ("\n\t" + v3)) + (v3 + ((v3 + v3) + ((("foo bar" + v3) + (((v3 + v3) + ("foo bar" + "foo bar")) + v3)) + v3)))) + "hello") + "\n\t")) + ((v3 + v3) + v3)) + ((v3 + (v3 + v3)) + "hello"))) {
                                                                                                            case "":
                                                                                                                var v59 int64 = -674
                                                                                                                _ = v59
                                                                                                            default:
                                                                                                                var v60 uint32 = 918
                                                                                                                _ = v60
                                                                                                            }
                                                                                                            var v61 uint8 = (v2 ^ v2)
                                                                                                            _ = v61
                                                                                                        }
                                                                                                        v62 := GenS7{F0: -32768, F1: (0.0 <= (((0.1 * 1.5) + 0.1) + (-2.25 - 0.1)))}
                                                                                                        _ = v62
                                                                                                    }
                                                                                                default:
                                                                                                    var v63 float64 = 1.5
                                                                                                    _ = v63
                                                                                                }
                                                                                            } else {
                                                                                                var v64 int64 = 42
                                                                                                if (!v4) {
                                                                                                    v65 := GenS1{F0: (2 > 111), F1: v3}
                                                                                                    _ = len(v1)
                                                                                                    _ = v65
                                                                                                }
                                                                                                if (127 >= (v2 % (((v2 << ((4294967295) & 63))) | 1))) {
                                                                                                    switch fi8 {
                                                                                                    case -356:
                                                                                                        v66 := GenS9{F0: 42, F1: v34, F2: 4294967295}
                                                                                                        _ = v66
                                                                                                    case -198:
                                                                                                        v34 = (v34 << ((18446744073709551615) & 63))
                                                                                                    case -135:
                                                                                                        if v4 {
                                                                                                            v67 := GenS10{F0: ("a" + (("foo bar" + (v3 + (v3 + v3))) + (v3 + (((v3 + v3) + "a") + (("foo bar" + ((v3 + ("hello" + (v3 + "\n\t"))) + v3)) + ("\n\t" + "")))))), F1: 884}
                                                                                                            var v68 uint32 = 2147483648
                                                                                                            _ = v67
                                                                                                            _ = v68
                                                                                                        } else {
                                                                                                            v34 -= v34
                                                                                                        }
                                                                                                    default:
                                                                                                        v64 = (v64 % ((v64) | 1))
                                                                                                        var v69 bool = true
                                                                                                        if (1e10 < float64(((-2.25 + 3.14159) - 0.1))) {
                                                                                                            var v70 uint16 = 964
                                                                                                            switch v3 {
                                                                                                            case "a":
                                                                                                                v71 := GenS2{F0: v64, F1: v2, F2: v64}
                                                                                                                fmt.Println(v3, v34, v69)
                                                                                                                _ = v71
                                                                                                            case "":
                                                                                                                var v72 string = "hello"
                                                                                                                _ = v72
                                                                                                            case "foo bar":
                                                                                                                var v73 uint = 349
                                                                                                                _ = v73
                                                                                                            }
                                                                                                            _ = v70
                                                                                                        }
                                                                                                        fmt.Println(v34, fi10, v3)
                                                                                                        v74 := GenS11{F0: v64, F1: 2}
                                                                                                        _ = v69
                                                                                                        _ = v74
                                                                                                    }
                                                                                                    v64 += (v64 * (v64 ^ v64))
                                                                                                } else {
                                                                                                    if (false && false) {
                                                                                                        if (0 < 42) {
                                                                                                            if false {
                                                                                                                v75 := GenS8{F0: -960, F1: 0, F2: v2}
                                                                                                                _ = v75
                                                                                                            }
                                                                                                            _ = len(v6)
                                                                                                            fmt.Println(v2, v34, v4)
                                                                                                        }
                                                                                                        v76 := GenS4{F0: 42, F1: (v64 << ((4294967295) & 63))}
                                                                                                        _ = v76
                                                                                                    } else {
                                                                                                        v3 = (v3 + v3)
                                                                                                        switch v3 {
                                                                                                        case "":
                                                                                                            if v4 {
                                                                                                                v77 := GenS5{F0: (v64 / ((v64) | 1)), F1: (^(v2)), F2: 18446744073709551615}
                                                                                                                if v4 {
                                                                                                                    v34 = (v34 >> ((42) & 63))
                                                                                                                    v2 = v2
                                                                                                                    _ = len(v1)
                                                                                                                    v78 := GenS12{F0: v4, F1: (v64 | v64), F2: 555}
                                                                                                                    _ = v78
                                                                                                                }
                                                                                                                _ = v77
                                                                                                            } else {
                                                                                                                v6 = append(v6, (fi10 >> ((134) & 63)))
                                                                                                                fmt.Println(v34, v64, v3)
                                                                                                                var v79 int64 = v64
                                                                                                                v80 := GenS11{F0: -9223372036854775808, F1: 4294967295}
                                                                                                                _ = v79
                                                                                                                _ = v80
                                                                                                            }
                                                                                                        case "a":
                                                                                                            v3 = (v3 + (v3 + (v3 + v3)))
                                                                                                        }
                                                                                                    }
                                                                                                }
                                                                                                fmt.Println(fi10)
                                                                                                switch fi8 {
                                                                                                case 32767:
                                                                                                    var v81 int = -2147483648
                                                                                                    _ = v81
                                                                                                }
                                                                                                _ = v64
                                                                                            }
                                                                                        default:
                                                                                            if v4 {
                                                                                                delete(v1, 2147483648)
                                                                                            }
                                                                                        }
                                                                                        var v82 uint8 = v2
                                                                                        if true {
                                                                                            v3 = v3
                                                                                            var v83 float64 = (-2.25 * 0.1)
                                                                                            _ = v83
                                                                                        } else {
                                                                                            fmt.Println(v3)
                                                                                        }
                                                                                        _ = v82
                                                                                    } else {
                                                                                        var v84 float64 = 0.1
                                                                                        v2 += (^(v2))
                                                                                        if len(v6) >= 2 {
                                                                                            v6 = v6[1:len(v6)]
                                                                                        }
                                                                                        _ = v84
                                                                                    }
                                                                                    v34 -= int32(fi10)
                                                                                    v85 := GenS9{F0: 949, F1: (v34 + (v34 & (v34 - -491))), F2: fi8}
                                                                                    _ = v85
                                                                                default:
                                                                                    if v4 {
                                                                                        v86 := GenS5{F0: 9223372036854775807, F1: (v2 % ((35) | 1)), F2: 18446744073709551615}
                                                                                        fmt.Println(v3, fi8, v4)
                                                                                        var v87 uint16 = 1
                                                                                        _ = v86
                                                                                        _ = v87
                                                                                    }
                                                                                }
                                                                            }
                                                                        default:
                                                                            v2 -= 127
                                                                        }
                                                                    }
                                                                    fmt.Println(v3)
                                                                } else {
                                                                    if (true && v4) {
                                                                        if true {
                                                                            var v88 int16 = -428
                                                                            _ = v88
                                                                        }
                                                                        var v89 uint32 = 2
                                                                        var v90 int64 = -2
                                                                        if v4 {
                                                                            var v91 string = v3
                                                                            _ = v91
                                                                        } else {
                                                                            v90 += v90
                                                                            if (v90 >= 9223372036854775807) {
                                                                                if (2 >= 0) {
                                                                                    var v92 uint32 = (v89 - (v89 & 4294967295))
                                                                                    _ = v92
                                                                                } else {
                                                                                    v89 = 42
                                                                                    var v93 uint8 = uint8(v89)
                                                                                    _ = v93
                                                                                }
                                                                                v94 := GenS2{F0: v90, F1: v2, F2: -577}
                                                                                if false {
                                                                                    _ = len(v1)
                                                                                    fmt.Println(v3)
                                                                                } else {
                                                                                    v3 = v3
                                                                                }
                                                                                _ = v94
                                                                            } else {
                                                                                var v95 uint64 = 599
                                                                                switch "" {
                                                                                case "a":
                                                                                    if v4 {
                                                                                        v96 := GenS2{F0: (v90 ^ v90), F1: v2, F2: v90}
                                                                                        v3 = v3
                                                                                        _ = v96
                                                                                    } else {
                                                                                        v97 := GenS6{F0: (v34 | v34)}
                                                                                        _ = v97
                                                                                    }
                                                                                    var v98 uint64 = 1
                                                                                    var v99 uint32 = v89
                                                                                    _ = v98
                                                                                    _ = v99
                                                                                default:
                                                                                    var v100 uint32 = (v89 / (((v89 / (((v89 % ((v89) | 1))) | 1))) | 1))
                                                                                    _ = v100
                                                                                }
                                                                                _ = v95
                                                                            }
                                                                        }
                                                                        _ = v89
                                                                        _ = v90
                                                                    } else {
                                                                        v4 = true
                                                                        if (!(-1 > -9223372036854775808)) {
                                                                            if false {
                                                                                if len(v6) > 0 {
                                                                                    v6[0] = 127
                                                                                }
                                                                                v34 = 0
                                                                                var v101 bool = (("a" >= v3) || false)
                                                                                var v102 int32 = v34
                                                                                v3 = v3
                                                                                _ = len(v1)
                                                                                if (!true) {
                                                                                    v3 = (((v3 + ((("a" + "") + v3) + "foo bar")) + "") + "")
                                                                                    v103 := GenS5{F0: 42, F1: v2, F2: 469}
                                                                                    _ = v103
                                                                                }
                                                                                _ = len(v1)
                                                                                v104 := GenS10{F0: v3, F1: 42}
                                                                                fmt.Println(v101)
                                                                                v4 = v101
                                                                                v3 = v3
                                                                                _ = v101
                                                                                _ = v102
                                                                                _ = v104
                                                                            }
                                                                        } else {
                                                                            var v105 int16 = 1
                                                                            var v106 int32 = (v34 >> ((919) & 63))
                                                                            if (!v4) {
                                                                                v6 = append(v6, (-(fi10)))
                                                                            } else {
                                                                                v6 = append(v6, fi8)
                                                                                var v107 uint16 = 307
                                                                                _ = v107
                                                                            }
                                                                            _ = v105
                                                                            _ = v106
                                                                        }
                                                                        fmt.Println(v34, v4, v34)
                                                                        v4 = true
                                                                    }
                                                                }
                                                            }
                                                            _ = v34
                                                        }
                                                    case "\n\t":
                                                        v3 = v3
                                                    }
                                                    v108 := GenS11{F0: 0, F1: 521}
                                                    _ = v108
                                                case "foo bar":
                                                    if ((^(v2)) <= v2) {
                                                        _ = len(v1)
                                                        var v109 uint8 = (v2 - (^(v2)))
                                                        _ = v109
                                                    } else {
                                                        if v4 {
                                                            if (-9223372036854775808 < -9223372036854775808) {
                                                                if v4 {
                                                                    switch (fi8 ^ fi8) {
                                                                    case -985:
                                                                        var v110 int64 = 42
                                                                        _ = v110
                                                                    case 593:
                                                                        v2 += (v2 * (v2 ^ v2))
                                                                    }
                                                                } else {
                                                                    v4 = v4
                                                                }
                                                                switch 32767 {
                                                                case 1:
                                                                    v2 -= (v2 * (v2 | (^(v2))))
                                                                case -1:
                                                                    var v111 uint64 = 0
                                                                    _ = v111
                                                                }
                                                                var v112 int8 = -4
                                                                _ = v112
                                                            } else {
                                                                v4 = v4
                                                                v113 := GenS12{F0: v4, F1: -1, F2: 1}
                                                                v2 -= 2
                                                                _ = v113
                                                            }
                                                            fmt.Println(v4)
                                                        }
                                                        if (42 >= 42) {
                                                            var v114 int16 = -129
                                                            _ = v114
                                                        } else {
                                                            var v115 uint32 = 345
                                                            var v116 uint = uint(v115)
                                                            switch v3 {
                                                            case "a":
                                                                var v117 int32 = -841
                                                                _ = v117
                                                            default:
                                                                if v4 {
                                                                    if v4 {
                                                                        if true {
                                                                            if (true && true) {
                                                                                _ = len(v6)
                                                                            } else {
                                                                                fmt.Println(fi10, v4)
                                                                                v2 -= 193
                                                                                switch (-(fi10)) {
                                                                                case 586:
                                                                                    fmt.Println(fi8)
                                                                                case -2147483648:
                                                                                    delete(v1, v115)
                                                                                case 1:
                                                                                    fmt.Println(v115, v116)
                                                                                }
                                                                                v116 = v116
                                                                            }
                                                                        } else {
                                                                            switch fi10 {
                                                                            case 2147483647:
                                                                                v118 := GenS13{F0: -1, F1: 1, F2: false}
                                                                                _ = v118
                                                                            }
                                                                            if ((v116 % (((v116 | v116)) | 1)) != v116) {
                                                                                if (445 < 0) {
                                                                                    var v119 uint = (v116 | (v116 % (((v116 % ((v116) | 1))) | 1)))
                                                                                    _ = v119
                                                                                }
                                                                                v120 := GenS14{F0: 1}
                                                                                _ = v120
                                                                            }
                                                                            fmt.Println(v116)
                                                                        }
                                                                    } else {
                                                                        v4 = ((32767 == -648) && (v4 && (!false)))
                                                                    }
                                                                    var v121 int8 = 42
                                                                    _ = v121
                                                                } else {
                                                                    if false {
                                                                        var v122 uint32 = (v115 * 0)
                                                                        delete(v1, (v122 / ((v115) | 1)))
                                                                        _ = v122
                                                                    }
                                                                    var v123 int64 = -9223372036854775808
                                                                    _ = v123
                                                                }
                                                            }
                                                            fmt.Println(v116)
                                                            _ = v115
                                                            _ = v116
                                                        }
                                                        var v124 uint64 = 2
                                                        var v125 int64 = -9223372036854775808
                                                        v126 := v1[947]
                                                        var v127 float64 = 0.1
                                                        fmt.Println(fi8, v3)
                                                        if v4 {
                                                            var v128 int32 = -2147483648
                                                            v126 = fi8
                                                            v129 := GenS15{F0: (v3 + "foo bar")}
                                                            v130 := GenS15{F0: "hello"}
                                                            v131 := GenS16{F0: (v4 && (0 != 0))}
                                                            if true {
                                                                delete(v1, 2)
                                                            } else {
                                                                var v132 string = "hello"
                                                                _ = v132
                                                            }
                                                            fmt.Println(v124)
                                                            fmt.Println(v2, fi8)
                                                            _ = v128
                                                            _ = v129
                                                            _ = v130
                                                            _ = v131
                                                        } else {
                                                            if v4 {
                                                                if (!false) {
                                                                    v133 := GenS17{F0: v2, F1: 467, F2: 42}
                                                                    _ = v133
                                                                } else {
                                                                    if (!("hello" == (v3 + v3))) {
                                                                        v134 := GenS18{F0: v2, F1: (v124 * (v124 | v124))}
                                                                        _ = v134
                                                                    }
                                                                }
                                                                v3 = ("\n\t" + v3)
                                                                switch fi8 {
                                                                case -2147483648:
                                                                    if false {
                                                                        v4 = false
                                                                    }
                                                                    if (v2 != v2) {
                                                                        v1[65535] = (v126 * (^(fi10)))
                                                                        if (-1 != (^(v125))) {
                                                                            v1[2147483647] = v126
                                                                        }
                                                                        v124 = (v124 << ((18446744073709551615) & 63))
                                                                    } else {
                                                                        var v135 int32 = -462
                                                                        var v136 int16 = 1000
                                                                        if true {
                                                                            v137 := GenS19{F0: v3}
                                                                            var v138 uint8 = v2
                                                                            v126 += -393
                                                                            _ = v137
                                                                            _ = v138
                                                                        } else {
                                                                            var v139 uint64 = (v124 / ((742) | 1))
                                                                            _ = v139
                                                                        }
                                                                        delete(v1, 2147483648)
                                                                        _ = v135
                                                                        _ = v136
                                                                    }
                                                                case 128:
                                                                    var v140 float64 = (v127 - 3.14159)
                                                                    _ = v140
                                                                default:
                                                                    v141 := GenS20{F0: 100, F1: 255, F2: fi8}
                                                                    _ = v141
                                                                }
                                                            }
                                                            fmt.Println(v126, v4)
                                                        }
                                                        _ = v124
                                                        _ = v125
                                                        _ = v126
                                                        _ = v127
                                                    }
                                                }
                                                v4 = (0 < 2147483647)
                                                fmt.Println(fi10, v2)
                                            } else {
                                                if true {
                                                    var v142 int = fi10
                                                    v4 = v4
                                                    v142 -= fi10
                                                    _ = v142
                                                } else {
                                                    switch (^(fi8)) {
                                                    case 32768:
                                                        if len(v6) >= 2 {
                                                            v6 = v6[1:len(v6)]
                                                        }
                                                    case 612:
                                                        var v143 uint = 4294967295
                                                        _ = v143
                                                    case -9223372036854775808:
                                                        v3 = v3
                                                    }
                                                }
                                                v4 = v4
                                            }
                                            _ = v30
                                        case "hello":
                                            if true {
                                                var v144 int32 = -2147483648
                                                switch v3 {
                                                case "a":
                                                    fmt.Println(v3)
                                                case "hello":
                                                    switch (fi10 % ((fi10) | 1)) {
                                                    case -1:
                                                        v145 := v1[558]
                                                        _ = v145
                                                    case 525:
                                                        switch v3 {
                                                        case "":
                                                            v2 = 0
                                                        default:
                                                            if len(v6) >= 2 {
                                                                v6 = v6[1:len(v6)]
                                                            }
                                                        }
                                                    }
                                                case "":
                                                    switch (v3 + ("foo bar" + (v3 + (v3 + v3)))) {
                                                    case "a":
                                                        if true {
                                                            fmt.Println(fi8, v4)
                                                            v4 = (255 != 689)
                                                            switch ("a" + "a") {
                                                            case "foo bar":
                                                                var v146 string = (v3 + (v3 + v3))
                                                                _ = v146
                                                            case "\n\t":
                                                                var v147 int64 = 9223372036854775807
                                                                _ = v147
                                                            default:
                                                                if v4 {
                                                                    var v148 float32 = (float32(3.14159) + (3.14159 / 1e10))
                                                                    _ = len(v6)
                                                                    if v4 {
                                                                        v148 += v148
                                                                        _ = len(v6)
                                                                        var v149 string = (("hello" + ("\n\t" + "foo bar")) + ("hello" + ""))
                                                                        _ = v149
                                                                    }
                                                                    var v150 string = ("foo bar" + (v3 + ("a" + (((v3 + "foo bar") + "a") + (((v3 + v3) + v3) + (v3 + ("foo bar" + "")))))))
                                                                    v148 = 0.0
                                                                    switch "foo bar" {
                                                                    case "\n\t":
                                                                        v144 += v144
                                                                    case "a":
                                                                        v151 := GenS8{F0: -1, F1: 31, F2: (^(v2))}
                                                                        _ = v151
                                                                    case "hello":
                                                                        if true {
                                                                            v148 += 0.0
                                                                        }
                                                                    }
                                                                    var v152 float32 = 1e10
                                                                    _ = v148
                                                                    _ = v150
                                                                    _ = v152
                                                                } else {
                                                                    var v153 int64 = -93
                                                                    v154 := GenS21{F0: uint32(v2), F1: 1}
                                                                    _ = v153
                                                                    _ = v154
                                                                }
                                                            }
                                                        } else {
                                                            var v155 float64 = float64(0.0)
                                                            _ = v155
                                                        }
                                                    default:
                                                        if (9223372036854775807 < 874) {
                                                            var v156 uint32 = 2147483647
                                                            _ = v156
                                                        }
                                                    }
                                                }
                                                v3 = v3
                                                _ = v144
                                            } else {
                                                fmt.Println(v3, fi8)
                                                if (v2 >= v2) {
                                                    switch "a" {
                                                    case "a":
                                                        var v157 uint8 = v2
                                                        _ = v157
                                                    }
                                                    v158 := GenS22{F0: ("a" + ""), F1: (v2 + 255)}
                                                    v159 := GenS23{F0: (v2 >> ((42) & 63))}
                                                    if v4 {
                                                        if v4 {
                                                            var v160 int64 = -319
                                                            v160 += 536
                                                            fmt.Println(fi10, fi10)
                                                            _ = v160
                                                        }
                                                        v161 := GenS6{F0: -534}
                                                        var v162 int8 = 127
                                                        if (!v4) {
                                                            var v163 int = fi8
                                                            if true {
                                                                delete(v1, 772)
                                                            }
                                                            _ = v163
                                                        } else {
                                                            v6 = append(v6, (fi8 & -417))
                                                            v164 := GenS3{F0: 151, F1: 221}
                                                            _ = v164
                                                        }
                                                        _ = v161
                                                        _ = v162
                                                    }
                                                    _ = v158
                                                    _ = v159
                                                }
                                                _ = len(v1)
                                            }
                                            fmt.Println(fi10, v2)
                                            if false {
                                                v165 := GenS17{F0: v2, F1: 255, F2: 189}
                                                _ = v165
                                            } else {
                                                v6 = append(v6, fi8)
                                                v6 = append(v6, (fi10 * fi10))
                                                var v166 float32 = 0.0
                                                _ = v166
                                            }
                                        }
                                    }
                                    _ = fi10
                                }
                                if (!false) {
                                    for fi167 := 0; fi167 < 6; fi167++ {
                                        var v168 int32 = -1
                                        fmt.Println(fi8)
                                        var v169 int16 = -511
                                        v2 -= v2
                                        _ = fi167
                                        _ = v168
                                        _ = v169
                                    }
                                } else {
                                    if (-1 != 1) {
                                        for fi170 := 0; fi170 < 12; fi170++ {
                                            v171 := GenS14{F0: 4294967295}
                                            v4 = (!(0.0 <= ((3.14159 - (3.14159 + float64((1e10 + 1.5)))) + (1.5 / 1.0))))
                                            _ = fi170
                                            _ = v171
                                        }
                                        v4 = false
                                        for fi172 := 0; fi172 < 15; fi172++ {
                                            v2 = 255
                                            switch (fi172 % (((fi172 - fi172)) | 1)) {
                                            case 414:
                                                fmt.Println(v2)
                                            case 128:
                                                if (float32((1.5 / 1.0)) != 0.1) {
                                                    if ((4294967295 < 508) || (!((-1 >= -93) && (false || true)))) {
                                                        var v173 float64 = -2.25
                                                        var v174 int32 = 0
                                                        var v175 int32 = -70
                                                        v176 := GenS7{F0: -968, F1: v4}
                                                        if v4 {
                                                            var v177 int32 = (v174 << ((12) & 63))
                                                            _ = v177
                                                        } else {
                                                            if v4 {
                                                                var v178 uint16 = 65535
                                                                _ = v178
                                                            }
                                                            if (4294967295 < 65535) {
                                                                v179 := GenS20{F0: 23, F1: 0, F2: (fi8 - (fi8 - (fi8 | int(v2))))}
                                                                v180 := GenS24{F0: 128, F1: (^(v174))}
                                                                v175 = (v175 + v174)
                                                                _ = v179
                                                                _ = v180
                                                            } else {
                                                                if v4 {
                                                                    v181 := GenS25{F0: 353, F1: (-(fi8)), F2: -1}
                                                                    var v182 uint32 = 54
                                                                    v174 = (v174 - v174)
                                                                    v175 -= v175
                                                                    _ = v181
                                                                    _ = v182
                                                                } else {
                                                                    switch -1 {
                                                                    case 38:
                                                                        if (!v4) {
                                                                            var v183 uint16 = 2
                                                                            var v184 int64 = -961
                                                                            switch v3 {
                                                                            case "a":
                                                                                if len(v6) > 0 {
                                                                                    v6[0] = (fi172 << ((1) & 63))
                                                                                }
                                                                            }
                                                                            _ = v183
                                                                            _ = v184
                                                                        } else {
                                                                            if len(v6) >= 2 {
                                                                                v6 = v6[1:len(v6)]
                                                                            }
                                                                            if v4 {
                                                                                var v185 float32 = 3.14159
                                                                                _ = v185
                                                                            }
                                                                            v175 = v174
                                                                            v174 -= 763
                                                                            switch "a" {
                                                                            case "":
                                                                                switch fi172 {
                                                                                case -2147483648:
                                                                                    if (65535 != 937) {
                                                                                        var v186 int = fi8
                                                                                        v4 = ((-(fi172)) != fi8)
                                                                                        _ = v186
                                                                                    }
                                                                                case 9223372036854775807:
                                                                                    v173 -= v173
                                                                                case 128:
                                                                                    v187 := GenS9{F0: 18446744073709551615, F1: v174, F2: (fi8 >> ((255) & 63))}
                                                                                    _ = v187
                                                                                }
                                                                            case "a":
                                                                                v174 = 1
                                                                            case "foo bar":
                                                                                delete(v1, uint32(v2))
                                                                            default:
                                                                                v4 = (255 != 894)
                                                                            }
                                                                            v3 = ("" + "hello")
                                                                            v188 := GenS26{F0: fi8}
                                                                            var v189 int32 = 207
                                                                            v4 = v4
                                                                            if (!v4) {
                                                                                if (1000 < 42) {
                                                                                    v190 := GenS27{F0: -1, F1: 2147483648}
                                                                                    fmt.Println(v173, v189, v174)
                                                                                    _ = v190
                                                                                } else {
                                                                                    v191 := GenS28{F0: 813, F1: 42}
                                                                                    v174 -= -866
                                                                                    var v192 uint64 = 4294967295
                                                                                    _ = v191
                                                                                    _ = v192
                                                                                }
                                                                                fmt.Println(v174, v174)
                                                                            }
                                                                            v193 := GenS29{F0: 113, F1: 943}
                                                                            var v194 int64 = 888
                                                                            _ = v188
                                                                            _ = v189
                                                                            _ = v193
                                                                            _ = v194
                                                                        }
                                                                    case 42:
                                                                        var v195 uint8 = 127
                                                                        _ = v195
                                                                    case 128:
                                                                        var v196 bool = (1 < 2)
                                                                        switch ("hello" + ((v3 + (v3 + v3)) + (v3 + ((v3 + (v3 + "\n\t")) + v3)))) {
                                                                        case "hello":
                                                                            var v197 uint32 = 4294967295
                                                                            _ = v197
                                                                        default:
                                                                            v198 := GenS28{F0: 0, F1: -1}
                                                                            _ = v198
                                                                        }
                                                                        var v199 int32 = int32(v2)
                                                                        _ = v196
                                                                        _ = v199
                                                                    default:
                                                                        switch fi172 {
                                                                        case 32768:
                                                                            v200 := GenS3{F0: 1, F1: v175}
                                                                            _ = v200
                                                                        case 766:
                                                                            if (255 != 0) {
                                                                                if v4 {
                                                                                    var v201 int8 = 42
                                                                                    v174 = v174
                                                                                    if v4 {
                                                                                        fmt.Println(v175)
                                                                                        switch "" {
                                                                                        case "\n\t":
                                                                                            v6 = append(v6, (-(fi8)))
                                                                                        case "hello":
                                                                                            if (255 != 1) {
                                                                                                if (32767 >= uint16(v2)) {
                                                                                                    var v202 int32 = (-(v174))
                                                                                                    switch (fi172 >> ((255) & 63)) {
                                                                                                    case -1:
                                                                                                        fmt.Println(v4, v173)
                                                                                                    case -75:
                                                                                                        v203 := GenS23{F0: (v2 >> ((597) & 63))}
                                                                                                        _ = v203
                                                                                                    case -2147483648:
                                                                                                        var v204 float64 = float64(float32((0.0 * ((((((((-2.25 - 0.1) + 3.14159) * (float64(float32((float64((0.1 + float32(v173))) - v173))) * (1.5 / 1.5))) / v173) / 0.1) * v173) * v173) + v173))))
                                                                                                        _ = v204
                                                                                                    }
                                                                                                    _ = v202
                                                                                                } else {
                                                                                                    v4 = v4
                                                                                                    v205 := GenS30{F0: 585}
                                                                                                    _ = v205
                                                                                                }
                                                                                            }
                                                                                        case "foo bar":
                                                                                            fmt.Println(v4, v174)
                                                                                        default:
                                                                                            if (v174 == 1) {
                                                                                                v206 := GenS31{F0: v4, F1: 849}
                                                                                                _ = v206
                                                                                            } else {
                                                                                                v2 += (v2 + v2)
                                                                                                if (v4 || true) {
                                                                                                    if (uint16(v2) < 820) {
                                                                                                        var v207 int = (fi172 << ((42) & 63))
                                                                                                        _ = v207
                                                                                                    }
                                                                                                } else {
                                                                                                    v6 = append(v6, (-(fi8)))
                                                                                                    fmt.Println(v173, v4, v3)
                                                                                                }
                                                                                                fmt.Println(v2)
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                    _ = v201
                                                                                }
                                                                                if false {
                                                                                    v174 -= (v174 & v174)
                                                                                }
                                                                                var v208 uint = 1
                                                                                _ = v208
                                                                            }
                                                                        }
                                                                    }
                                                                    if (v174 < v174) {
                                                                        var v209 uint = 4294967295
                                                                        _ = v209
                                                                    }
                                                                    var v210 uint64 = 0
                                                                    _ = v210
                                                                }
                                                                _ = len(v6)
                                                            }
                                                        }
                                                        v2 = (^(v2))
                                                        _ = v173
                                                        _ = v174
                                                        _ = v175
                                                        _ = v176
                                                    } else {
                                                        if false {
                                                            var v211 string = ("foo bar" + "hello")
                                                            _ = v211
                                                        }
                                                        delete(v1, 65535)
                                                        var v212 uint16 = 852
                                                        _ = v212
                                                    }
                                                    if (v4 && ((v2 - (v2 << ((4294967295) & 63))) == v2)) {
                                                        var v213 int16 = 32767
                                                        v2 = 83
                                                        _ = v213
                                                    }
                                                    v214 := GenS16{F0: (((!v4) || (v4 && v4)) && (100 >= 1))}
                                                    v4 = false
                                                    _ = v214
                                                } else {
                                                    var v215 uint16 = 42
                                                    var v216 float32 = 1e10
                                                    v217 := GenS19{F0: (v3 + "\n\t")}
                                                    _ = v215
                                                    _ = v216
                                                    _ = v217
                                                }
                                            case 9223372036854775807:
                                                if v4 {
                                                    v218 := GenS2{F0: -802, F1: 128, F2: 819}
                                                    if (v4 || (0.0 >= 0.1)) {
                                                        v4 = (0 != 32767)
                                                        v6 = append(v6, (fi172 | (fi172 ^ (fi172 | 9223372036854775807))))
                                                        v4 = v4
                                                        fmt.Println(v3, v3, v2)
                                                        if (0.0 >= 1.5) {
                                                            if (!false) {
                                                                v4 = false
                                                                v219 := GenS6{F0: -1}
                                                                v4 = v4
                                                                _ = v219
                                                            }
                                                            var v220 int64 = 0
                                                            _ = v220
                                                        } else {
                                                            if v4 {
                                                                var v221 float32 = 3.14159
                                                                v221 += v221
                                                                _ = v221
                                                            } else {
                                                                var v222 int16 = -32768
                                                                _ = v222
                                                            }
                                                        }
                                                        v223 := GenS9{F0: 1, F1: -2147483648, F2: fi8}
                                                        _ = v223
                                                    } else {
                                                        v1[4294967295] = fi172
                                                    }
                                                    _ = v218
                                                }
                                            default:
                                                if v4 {
                                                    var v224 uint8 = v2
                                                    fmt.Println(v3, v3)
                                                    _ = v224
                                                }
                                            }
                                            v2 = v2
                                            _ = fi172
                                        }
                                    }
                                    for fi225 := 0; fi225 < 7; fi225++ {
                                        v226 := GenS3{F0: 2147483648, F1: 1}
                                        v227 := GenS32{F0: 760}
                                        _ = fi225
                                        _ = v226
                                        _ = v227
                                    }
                                    for fi228 := 0; fi228 < 14; fi228++ {
                                        v229 := v1[510]
                                        if len(v6) >= 2 {
                                            v6 = v6[1:len(v6)]
                                        }
                                        _ = len(v6)
                                        _ = fi228
                                        _ = v229
                                    }
                                }
                                var v230 float64 = 0.1
                                _ = v230
                            }
                            _ = fi8
                        }
                        var v231 uint64 = 4294967295
                        v2 += 42
                        _ = v231
                    }
                    var v232 bool = (802 != 671)
                    if v232 {
                        if ((126 < 0) || v4) {
                            switch "a" {
                            case "hello":
                                _ = len(v1)
                            default:
                                fmt.Println(v2, v232)
                            }
                        }
                        v4 = v4
                    } else {
                        var v233 bool = (!(1 > -128))
                        if v4 {
                            if (!(!v233)) {
                                switch -979 {
                                case 1:
                                    var v234 int = -2147483648
                                    _ = v234
                                case 4294967295:
                                    if true {
                                        delete(v1, 2)
                                        if v4 {
                                            switch 0 {
                                            case -707:
                                                v4 = true
                                            }
                                        }
                                    } else {
                                        if len(v6) >= 2 {
                                            v6 = v6[1:len(v6)]
                                        }
                                        var v235 int16 = -37
                                        _ = v235
                                    }
                                case 426:
                                    if (float64((((3.14159 / -2.25) - 1.5) + 1e10)) != 3.14159) {
                                        var v236 float64 = (-2.25 / (((1e10 * float64(0.1)) * (0.0 + ((1e10 / 0.1) * 3.14159))) * float64(1.5)))
                                        v6 = append(v6, 717)
                                        _ = v236
                                    } else {
                                        fmt.Println(v3, v232, v232)
                                    }
                                }
                                v232 = (-9223372036854775808 < 1)
                                fmt.Println(v3, v232)
                            } else {
                                switch v3 {
                                case "a":
                                    v2 -= (v2 << ((571) & 63))
                                case "":
                                    v237 := GenS18{F0: (v2 - v2), F1: 42}
                                    _ = v237
                                case "\n\t":
                                    for fi238 := 0; fi238 < 5; fi238++ {
                                        if (-2.25 <= (0.1 + float32(1e10))) {
                                            v232 = false
                                            if len(v6) > 0 {
                                                v6[0] = (fi238 + (fi238 - (fi238 ^ (fi238 - (fi238 + fi238)))))
                                            }
                                        } else {
                                            for fi239 := 0; fi239 < 10; fi239++ {
                                                if (!v232) {
                                                    fmt.Println(fi239, v4, v4)
                                                    fmt.Println(fi239, v4)
                                                }
                                                var v240 int64 = -1
                                                v4 = v232
                                                v232 = (515 <= 1)
                                                v241 := GenS5{F0: v240, F1: 233, F2: 0}
                                                _ = len(v1)
                                                _ = fi239
                                                _ = v240
                                                _ = v241
                                            }
                                            var v242 int32 = -1
                                            if ((fi238 & -9223372036854775808) > fi238) {
                                                switch (v3 + "a") {
                                                case "foo bar":
                                                    var v243 int = 42
                                                    _ = v243
                                                default:
                                                    switch (-(fi238)) {
                                                    case -423:
                                                        for fi244 := 0; fi244 < 19; fi244++ {
                                                            switch "\n\t" {
                                                            case "a":
                                                                var v245 bool = (2 > 0)
                                                                _ = v245
                                                            }
                                                            var v246 uint16 = 2
                                                            _ = len(v1)
                                                            if ((("foo bar" + (v3 + (v3 + (v3 + "foo bar")))) + v3) < (("foo bar" + ("foo bar" + v3)) + "\n\t")) {
                                                                if v232 {
                                                                    v3 = v3
                                                                }
                                                                v246 += (v246 / ((v246) | 1))
                                                                v2 = v2
                                                            }
                                                            var v247 uint32 = 68
                                                            _ = fi244
                                                            _ = v246
                                                            _ = v247
                                                        }
                                                    }
                                                }
                                                v232 = (!v4)
                                                fmt.Println(v233, v233, v232)
                                            } else {
                                                v2 -= (v2 / ((v2) | 1))
                                            }
                                            _ = v242
                                        }
                                        if ((!((v3 + ((v3 + (v3 + "a")) + v3)) == v3)) && v232) {
                                            switch (fi238 << ((480) & 63)) {
                                            case 127:
                                                var v248 uint8 = (v2 >> ((2) & 63))
                                                _ = v248
                                            default:
                                                var v249 int8 = 0
                                                _ = v249
                                            }
                                            for fi250 := 0; fi250 < 2; fi250++ {
                                                fmt.Println(v3, v2, fi250)
                                                var v251 int32 = -1
                                                _ = fi250
                                                _ = v251
                                            }
                                        } else {
                                            if len(v6) >= 2 {
                                                v6 = v6[1:len(v6)]
                                            }
                                            if false {
                                                var v252 int64 = 99
                                                if (-2147483648 <= 0) {
                                                    v232 = v4
                                                    v253 := GenS33{F0: 1, F1: (v233 && v4), F2: 0}
                                                    v233 = (2147483647 <= 1)
                                                    _ = v253
                                                } else {
                                                    _ = len(v1)
                                                    var v254 int16 = 1
                                                    _ = v254
                                                }
                                                var v255 float32 = (((0.1 * 0.0) + (float32((0.0 + 1.5)) * 1e10)) * 3.14159)
                                                _ = v252
                                                _ = v255
                                            } else {
                                                for fi256 := 0; fi256 < 14; fi256++ {
                                                    v257 := GenS34{F0: 1}
                                                    _ = fi256
                                                    _ = v257
                                                }
                                            }
                                            v233 = v232
                                        }
                                        _ = fi238
                                    }
                                }
                                v232 = v232
                                v4 = v4
                            }
                            v3 = v3
                            for fi258 := 0; fi258 < 17; fi258++ {
                                var v259 bool = (396 < 65535)
                                if ((32767 == 42) && (!false)) {
                                    for fi260 := 0; fi260 < 17; fi260++ {
                                        switch fi260 {
                                        case 1:
                                            if v233 {
                                                fmt.Println(v3)
                                                switch -2147483648 {
                                                case 32768:
                                                    var v261 string = (v3 + "\n\t")
                                                    _ = v261
                                                }
                                                if len(v6) >= 2 {
                                                    v6 = v6[1:len(v6)]
                                                }
                                            } else {
                                                var v262 uint16 = 2
                                                switch (fi258 << ((2) & 63)) {
                                                case 1:
                                                    _ = len(v6)
                                                    v2 = 42
                                                default:
                                                    v259 = (v259 || v233)
                                                }
                                                _ = v262
                                            }
                                        case 32768:
                                            if true {
                                                if len(v6) >= 2 {
                                                    v6 = v6[1:len(v6)]
                                                }
                                                var v263 uint = 1
                                                _ = v263
                                            }
                                        default:
                                            v6 = append(v6, fi260)
                                        }
                                        v6 = append(v6, (fi258 >> ((177) & 63)))
                                        _ = fi260
                                    }
                                    for fi264 := 0; fi264 < 15; fi264++ {
                                        if ((!(18446744073709551615 > 2)) || v4) {
                                            v259 = v259
                                            var v265 int16 = 1000
                                            v265 = (v265 / (((v265 * v265)) | 1))
                                            _ = v265
                                        } else {
                                            v266 := GenS10{F0: v3, F1: 42}
                                            v232 = (!(2 > 1))
                                            _ = v266
                                        }
                                        delete(v1, 96)
                                        v267 := GenS24{F0: v2, F1: -1}
                                        v233 = v233
                                        _ = fi264
                                        _ = v267
                                    }
                                }
                                if false {
                                    if v259 {
                                        fmt.Println(v233, v233)
                                        _ = len(v6)
                                    } else {
                                        v6 = append(v6, 508)
                                        var v268 uint16 = 32767
                                        _ = v268
                                    }
                                    if len(v6) > 0 {
                                        v6[0] = (fi258 ^ fi258)
                                    }
                                    fmt.Println(fi258, v232)
                                } else {
                                    if true {
                                        for fi269 := 0; fi269 < 13; fi269++ {
                                            delete(v1, 2147483647)
                                            if (!false) {
                                                var v270 uint32 = 65535
                                                fmt.Println(v232, v232)
                                                _ = v270
                                            } else {
                                                switch (^(fi269)) {
                                                case 128:
                                                    var v271 int = 32767
                                                    _ = v271
                                                default:
                                                    switch fi269 {
                                                    case -2147483648:
                                                        var v272 uint16 = 65535
                                                        _ = v272
                                                    }
                                                }
                                                v1[42] = fi258
                                                v4 = v232
                                            }
                                            fmt.Println(v3)
                                            _ = fi269
                                        }
                                    }
                                }
                                _ = fi258
                                _ = v259
                            }
                        }
                        _ = v233
                    }
                    _ = v232
                }
                _ = v6
            }
        }
    }
    v273 := []int{1, 9223372036854775807, -453}
    if (216 == 789) {
        for fi274 := 0; fi274 < 17; fi274++ {
            var v275 int32 = -1
            _ = fi274
            _ = v275
        }
    } else {
        if (86 < 100) {
            if (!(!(1e10 != (1e10 + 1e10)))) {
                var v276 int32 = 1
                for fi277 := 0; fi277 < 16; fi277++ {
                    v276 -= (v276 - v276)
                    for fi278 := 0; fi278 < 11; fi278++ {
                        var v279 uint = 1
                        _ = fi278
                        _ = v279
                    }
                    _ = fi277
                }
                v280 := GenS11{F0: 144, F1: 42}
                _ = v276
                _ = v280
            }
            _ = len(v273)
        }
    }
    if true {
        fmt.Println(v3, v2)
        if false {
            v4 = (float32(1.5) >= (((0.1 / 3.14159) + 1.5) + 1.5))
            switch -2147483648 {
            case -1:
                v4 = v4
            case 0:
                for fi281 := 0; fi281 < 19; fi281++ {
                    _ = len(v273)
                    var v282 bool = false
                    for fi283 := 0; fi283 < 13; fi283++ {
                        v284 := GenS33{F0: 199, F1: v282, F2: 0}
                        var v285 uint16 = 42
                        var v286 int32 = 2147483647
                        v273 = append(v273, fi281)
                        var v287 uint64 = 0
                        _ = fi283
                        _ = v284
                        _ = v285
                        _ = v286
                        _ = v287
                    }
                    _ = fi281
                    _ = v282
                }
            default:
                var v288 int8 = -1
                _ = v288
            }
            v4 = v4
        } else {
            var v289 uint8 = v2
            _ = v289
        }
        var v290 int32 = -1
        _ = v290
    }
    v3 = v3
    for fi291 := 0; fi291 < 6; fi291++ {
        switch ("\n\t" + v3) {
        case "\n\t":
            _ = len(v273)
        case "hello":
            var v292 int32 = 42
            _ = v292
        }
        v4 = false
        v293 := GenS27{F0: 42, F1: 896}
        _ = fi291
        _ = v293
    }
    var v294 uint32 = 2147483647
    var v295 uint32 = 42
    v4 = (!(1.5 > -2.25))
    for fi296 := 0; fi296 < 15; fi296++ {
        if false {
            for fi297 := 0; fi297 < 15; fi297++ {
                var v298 uint8 = v2
                v298 = v2
                v298 -= 128
                _ = fi297
                _ = v298
            }
            var v299 int32 = -1
            var v300 string = "\n\t"
            _ = v299
            _ = v300
        }
        for fi301 := 0; fi301 < 8; fi301++ {
            if v4 {
                v3 = "\n\t"
            }
            v295 -= v294
            switch v3 {
            case "a":
                fmt.Println(v295)
            case "foo bar":
                var v302 uint64 = 273
                _ = v302
            default:
                if (v4 || (!false)) {
                    v3 = ("hello" + ("a" + v3))
                    v303 := GenS35{F0: 1, F1: 42}
                    _ = v303
                } else {
                    if (32767 != -512) {
                        v304 := GenS18{F0: (v2 * v2), F1: 1}
                        if v4 {
                            fmt.Println(v3, v3)
                            var v305 float32 = float32(1e10)
                            _ = v305
                        }
                        if (616 != 42) {
                            switch "a" {
                            case "hello":
                                delete(v1, v294)
                                if v4 {
                                    if (32767 == -103) {
                                        switch int(v2) {
                                        case 127:
                                            if v4 {
                                                var v306 uint = 18446744073709551615
                                                _ = v306
                                            } else {
                                                switch (v3 + v3) {
                                                case "\n\t":
                                                    v273 = append(v273, (fi301 << ((340) & 63)))
                                                case "a":
                                                    var v307 int32 = 2147483647
                                                    _ = v307
                                                default:
                                                    v308 := GenS36{F0: v3, F1: fi301, F2: (^(fi301))}
                                                    _ = v308
                                                }
                                            }
                                        default:
                                            var v309 int16 = 0
                                            _ = v309
                                        }
                                    } else {
                                        v310 := GenS24{F0: (^(v2)), F1: int32(fi296)}
                                        fmt.Println(fi296, fi301, v4)
                                        v311 := GenS22{F0: v3, F1: (v2 / ((1) | 1))}
                                        _ = v310
                                        _ = v311
                                    }
                                    v3 = v3
                                    var v312 uint16 = 212
                                    _ = v312
                                }
                            default:
                                v294 -= v295
                            }
                        }
                        _ = v304
                    }
                    if len(v273) > 0 {
                        v273[0] = fi301
                    }
                    var v313 int8 = -1
                    _ = v313
                }
            }
            _ = fi301
        }
        var v314 int16 = -862
        _ = fi296
        _ = v314
    }
    v315 := GenS37{F0: v4}
    if (4294967295 > 2) {
        if v4 {
            if true {
                if v4 {
                    if len(v273) >= 2 {
                        v273 = v273[1:len(v273)]
                    }
                }
                switch -2147483648 {
                case -1:
                    switch v3 {
                    case "foo bar":
                        if v4 {
                            var v316 uint16 = 65535
                            if v4 {
                                var v317 uint16 = 1
                                v317 += (v316 >> ((4294967295) & 63))
                                for fi318 := 0; fi318 < 15; fi318++ {
                                    if (fi318 <= (fi318 * fi318)) {
                                        var v319 int = (fi318 | fi318)
                                        for fi320 := 0; fi320 < 12; fi320++ {
                                            v295 = (v295 * v295)
                                            if v4 {
                                                var v321 float64 = (1e10 * float64((float32(float64(float32(1.5))) * (0.1 * ((0.0 * 3.14159) * 0.0)))))
                                                v322 := GenS15{F0: v3}
                                                _ = v321
                                                _ = v322
                                            }
                                            v316 = (v316 * (v317 * v316))
                                            _ = fi320
                                        }
                                        _ = v319
                                    } else {
                                        v323 := GenS27{F0: 127, F1: (v295 ^ v295)}
                                        v324 := GenS37{F0: false}
                                        _ = v323
                                        _ = v324
                                    }
                                    v294 += 2147483648
                                    fmt.Println(v4, v3, v317)
                                    _ = fi318
                                }
                                _ = v317
                            } else {
                                v316 += v316
                            }
                            if v4 {
                                v2 += uint8(v316)
                                v316 -= (v316 << ((42) & 63))
                            } else {
                                v1[(v295 % ((v294) | 1))] = 42
                                v295 = (^(v294))
                                var v325 uint = 908
                                _ = v325
                            }
                            _ = v316
                        } else {
                            fmt.Println(v294, v294)
                            switch 701 {
                            case 0:
                                if (!((!v4) && true)) {
                                    var v326 uint32 = (v295 << ((662) & 63))
                                    _ = len(v1)
                                    _ = v326
                                }
                            case 468:
                                v4 = v4
                            case 127:
                                v327 := v1[276]
                                _ = v327
                            default:
                                var v328 uint64 = 0
                                _ = v328
                            }
                            var v329 uint8 = 2
                            _ = v329
                        }
                    case "":
                        for fi330 := 0; fi330 < 13; fi330++ {
                            v294 += (v295 >> ((305) & 63))
                            for fi331 := 0; fi331 < 13; fi331++ {
                                if v4 {
                                    switch (fi331 - 128) {
                                    case 127:
                                        if false {
                                            if (0 > 166) {
                                                var v332 float64 = 1e10
                                                if v4 {
                                                    switch (fi330 ^ fi331) {
                                                    case 4294967295:
                                                        v2 -= 42
                                                    case 0:
                                                        var v333 int = 1
                                                        _ = v333
                                                    }
                                                    v295 -= v294
                                                }
                                                _ = v332
                                            } else {
                                                var v334 uint64 = 1
                                                if len(v273) >= 2 {
                                                    v273 = v273[1:len(v273)]
                                                }
                                                _ = v334
                                            }
                                        } else {
                                            _ = len(v1)
                                            var v335 uint16 = 0
                                            _ = v335
                                        }
                                    default:
                                        v336 := GenS38{F0: "\n\t", F1: 655}
                                        var v337 uint16 = 65535
                                        if true {
                                            v2 = 127
                                            if v4 {
                                                var v338 int = 763
                                                var v339 bool = (!false)
                                                _ = v338
                                                _ = v339
                                            } else {
                                                var v340 int8 = -128
                                                fmt.Println(v2, v4)
                                                _ = v340
                                            }
                                        } else {
                                            _ = len(v1)
                                            var v341 uint64 = 378
                                            _ = v341
                                        }
                                        _ = v336
                                        _ = v337
                                    }
                                    switch -1 {
                                    case 127:
                                        if (318 >= (fi331 >> ((298) & 63))) {
                                            var v342 int32 = int32(v2)
                                            _ = v342
                                        }
                                    case -9223372036854775808:
                                        _ = len(v273)
                                    case -1:
                                        fmt.Println(v2, v294, fi330)
                                        v3 = "hello"
                                        var v343 bool = false
                                        _ = v343
                                    default:
                                        var v344 uint = 18446744073709551615
                                        _ = v344
                                    }
                                    switch 1 {
                                    case 128:
                                        v294 += v295
                                    case 0:
                                        v345 := GenS39{F0: -765, F1: (fi330 + (fi330 - -594)), F2: v3}
                                        _ = v345
                                    case 978:
                                        var v346 int16 = 111
                                        _ = v346
                                    }
                                }
                                v4 = ((1 <= 369) || (2 >= 596))
                                _ = fi331
                            }
                            _ = fi330
                        }
                    }
                case -2147483648:
                    var v347 uint16 = 32767
                    _ = v347
                case 32767:
                    v3 = (v3 + ("a" + v3))
                }
                switch -174 {
                case 4294967295:
                    if (0 == 128) {
                        for fi348 := 0; fi348 < 11; fi348++ {
                            for fi349 := 0; fi349 < 18; fi349++ {
                                v294 += v295
                                _ = fi349
                            }
                            fmt.Println(v4, fi348, v4)
                            v3 = v3
                            _ = fi348
                        }
                    } else {
                        for fi350 := 0; fi350 < 3; fi350++ {
                            if false {
                                switch v3 {
                                case "foo bar":
                                    for fi351 := 0; fi351 < 3; fi351++ {
                                        delete(v1, v295)
                                        _ = fi351
                                    }
                                default:
                                    v3 = v3
                                }
                                fmt.Println(v295)
                                _ = len(v1)
                            } else {
                                v4 = (v4 || ((0.0 - -2.25) >= 0.1))
                                for fi352 := 0; fi352 < 12; fi352++ {
                                    switch v3 {
                                    case "hello":
                                        v273 = append(v273, (fi350 + (fi350 ^ 259)))
                                        var v353 int16 = 731
                                        var v354 uint32 = 2
                                        _ = v353
                                        _ = v354
                                    case "":
                                        fmt.Println(v4, fi350)
                                    }
                                    _ = fi352
                                }
                                v355 := GenS40{F0: -811, F1: (42 <= 295), F2: (("hello" == v3) && true)}
                                _ = v355
                            }
                            var v356 uint = 1
                            v4 = (208 <= 0)
                            _ = fi350
                            _ = v356
                        }
                        if (579 <= 32767) {
                            for fi357 := 0; fi357 < 17; fi357++ {
                                var v358 uint64 = 2
                                _ = fi357
                                _ = v358
                            }
                            v359 := GenS41{F0: 42}
                            for fi360 := 0; fi360 < 11; fi360++ {
                                if v4 {
                                    if (1 <= 1) {
                                        switch (((("" + v3) + "\n\t") + ("" + v3)) + v3) {
                                        case "hello":
                                            var v361 uint16 = 373
                                            _ = v361
                                        default:
                                            v362 := GenS31{F0: v4, F1: 1}
                                            _ = v362
                                        }
                                    } else {
                                        var v363 uint = 42
                                        _ = v363
                                    }
                                }
                                v364 := GenS42{F0: -1, F1: (v295 >> ((1) & 63))}
                                _ = fi360
                                _ = v364
                            }
                            _ = v359
                        }
                        v295 = v294
                    }
                case -539:
                    v4 = true
                case -527:
                    switch 42 {
                    case -695:
                        for fi365 := 0; fi365 < 7; fi365++ {
                            var v366 int = fi365
                            var v367 float64 = 0.1
                            if (!false) {
                                v368 := GenS18{F0: (v2 + v2), F1: 2}
                                for fi369 := 0; fi369 < 18; fi369++ {
                                    var v370 uint8 = 50
                                    v371 := GenS22{F0: "\n\t", F1: 2}
                                    _ = fi369
                                    _ = v370
                                    _ = v371
                                }
                                var v372 int64 = 9223372036854775807
                                _ = v368
                                _ = v372
                            } else {
                                if len(v273) > 0 {
                                    v273[0] = (-(v366))
                                }
                            }
                            _ = fi365
                            _ = v366
                            _ = v367
                        }
                    default:
                        for fi373 := 0; fi373 < 4; fi373++ {
                            var v374 uint = 2
                            _ = len(v273)
                            _ = fi373
                            _ = v374
                        }
                    }
                default:
                    for fi375 := 0; fi375 < 3; fi375++ {
                        for fi376 := 0; fi376 < 19; fi376++ {
                            v377 := GenS43{F0: v295, F1: -32768}
                            _ = fi376
                            _ = v377
                        }
                        v294 = (^(v295))
                        var v378 bool = v4
                        _ = fi375
                        _ = v378
                    }
                }
            } else {
                v294 = (^(v294))
            }
        }
        switch -449 {
        case 823:
            var v379 uint32 = v294
            _ = v379
        case -9223372036854775808:
            for fi380 := 0; fi380 < 12; fi380++ {
                if v4 {
                    v381 := GenS13{F0: -1, F1: 1, F2: ((0 > 124) || false)}
                    v2 -= (^(v2))
                    if (-550 < 42) {
                        if ((v2 & (v2 % ((v2) | 1))) < v2) {
                            for fi382 := 0; fi382 < 15; fi382++ {
                                var v383 uint64 = 18446744073709551615
                                v2 = (^(v2))
                                switch fi380 {
                                case -87:
                                    v294 += v294
                                    v295 = v295
                                case -1:
                                    v1[(^(v294))] = fi382
                                default:
                                    v295 -= v294
                                }
                                if v4 {
                                    var v384 int32 = 2147483647
                                    switch v3 {
                                    case "":
                                        v295 += (v294 + 2147483647)
                                    case "a":
                                        switch (v3 + v3) {
                                        case "hello":
                                            var v385 int64 = 778
                                            _ = v385
                                        case "a":
                                            v386 := GenS44{F0: 255, F1: (!(!v4)), F2: "foo bar"}
                                            _ = v386
                                        }
                                    default:
                                        var v387 uint16 = 65535
                                        _ = v387
                                    }
                                    var v388 bool = (1.5 < (1e10 - -2.25))
                                    _ = v384
                                    _ = v388
                                }
                                _ = fi382
                                _ = v383
                            }
                        } else {
                            v3 = "a"
                        }
                        v4 = true
                    } else {
                        if v4 {
                            for fi389 := 0; fi389 < 10; fi389++ {
                                v2 = 0
                                v2 -= uint8(v294)
                                var v390 uint16 = 255
                                _ = fi389
                                _ = v390
                            }
                            delete(v1, (v294 + v294))
                            if len(v273) >= 2 {
                                v273 = v273[1:len(v273)]
                            }
                        }
                    }
                    _ = v381
                }
                var v391 float64 = (3.14159 + -2.25)
                _ = fi380
                _ = v391
            }
        case 4294967295:
            fmt.Println(v4, v3)
        }
    }
    v392 := GenS33{F0: 1, F1: false, F2: 42}
    v295 += v294
    v295 = (^(v295))
    v393 := GenS39{F0: 2147483647, F1: 209, F2: "hello"}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v273:", len(v273), v273)
    fmt.Println("v294:", v294)
    fmt.Println("v295:", v295)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v273
    _ = v294
    _ = v295
    _ = v3
    _ = v315
    _ = v392
    _ = v393
    _ = v4
}
