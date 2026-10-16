package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint16; F1 uint; F2 int64 }
type GenS5 struct { F0 uint16; F1 uint64 }
type GenS6 struct { F0 int16; F1 uint; F2 bool }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 int32; F1 uint32 }
type GenS10 struct { F0 uint32; F1 uint64 }
type GenS11 struct { F0 uint8; F1 bool; F2 int16 }
type GenS12 struct { F0 uint32; F1 uint64 }
type GenS13 struct { F0 int8; F1 int8; F2 uint8 }
type GenS14 struct { F0 bool; F1 uint64; F2 uint16 }
type GenS15 struct { F0 string }
type GenS16 struct { F0 uint8; F1 uint; F2 int64 }
type GenS17 struct { F0 uint64; F1 uint; F2 bool }
type GenS18 struct { F0 int8; F1 int64 }
type GenS19 struct { F0 int8 }
type GenS20 struct { F0 uint8; F1 uint16 }
type GenS21 struct { F0 int }
type GenS22 struct { F0 bool; F1 int; F2 uint8 }
type GenS23 struct { F0 int32; F1 uint64 }
type GenS24 struct { F0 string; F1 bool; F2 uint32 }

func main() {
    var v1 int64 = 1
    v2 := []int{}
    var v3 uint32 = 2147483647
    v4 := []uint32{(v3 + (v3 - v3)), 925, (v3 | v3)}
    v1 += 9223372036854775807
    var v5 uint64 = 0
    fmt.Println(v3, v1)
    if (((v5 >> ((1) & 63)) <= v5) || (65535 < 2)) {
        for fi6 := 0; fi6 < 6; fi6++ {
            v7 := GenS1{F0: 86, F1: false}
            v3 += v3
            v3 = 2147483648
            _ = fi6
            _ = v7
        }
    }
    v8 := GenS1{F0: 2, F1: (2 > 564)}
    v9 := GenS2{F0: v5}
    fmt.Println(v5, v1)
    fmt.Println(v1)
    switch 42 {
    case -766:
        for fi10 := 0; fi10 < 5; fi10++ {
            for fi11 := 0; fi11 < 14; fi11++ {
                var v12 int16 = 42
                if len(v4) > 0 {
                    v4[0] = 0
                }
                _ = fi11
                _ = v12
            }
            var v13 int16 = -408
            _ = fi10
            _ = v13
        }
    case 0:
        switch -794 {
        case -355:
            v14 := GenS3{F0: 65535}
            _ = v14
        case -2147483648:
            switch "hello" {
            case "a":
                if ((207 != 42) && (v3 != (v3 & uint32(v5)))) {
                    _ = len(v2)
                } else {
                    switch 2147483647 {
                    case -642:
                        v1 = (-(v1))
                    case 308:
                        v1 = (-(v1))
                    case 586:
                        for fi15 := 0; fi15 < 10; fi15++ {
                            for fi16 := 0; fi16 < 3; fi16++ {
                                if (v1 != v1) {
                                    var v17 uint64 = v5
                                    fmt.Println(v1, fi15)
                                    _ = len(v2)
                                    fmt.Println(v5, fi16, v17)
                                    _ = v17
                                }
                                if false {
                                    if (-2.25 < float32((0.1 / (1e10 - 0.1)))) {
                                        v1 = v1
                                        v1 = (v1 + v1)
                                    } else {
                                        switch "foo bar" {
                                        case "hello":
                                            if ("a" < "\n\t") {
                                                v5 = v5
                                                v3 = (^(v3))
                                                v3 = v3
                                                if true {
                                                    switch (fi16 | (fi16 + 4294967295)) {
                                                    case 9223372036854775807:
                                                        if false {
                                                            fmt.Println(v5)
                                                            if true {
                                                                var v18 uint8 = 128
                                                                v19 := GenS1{F0: (^(v18)), F1: (42 >= int8(v1))}
                                                                switch fi16 {
                                                                case 32767:
                                                                    var v20 int16 = 32767
                                                                    _ = v20
                                                                case 9223372036854775807:
                                                                    v21 := GenS1{F0: v18, F1: true}
                                                                    _ = v21
                                                                default:
                                                                    v18 += (v18 % ((v18) | 1))
                                                                }
                                                                _ = v18
                                                                _ = v19
                                                            }
                                                            _ = len(v2)
                                                            var v22 int16 = 740
                                                            _ = v22
                                                        } else {
                                                            var v23 float64 = 1e10
                                                            v5 += 228
                                                            if len(v4) > 0 {
                                                                v4[0] = (v3 - v3)
                                                            }
                                                            if false {
                                                                fmt.Println(v5, fi15, fi15)
                                                            }
                                                            var v24 uint8 = 255
                                                            switch "\n\t" {
                                                            case "\n\t":
                                                                fmt.Println(v3, v3)
                                                                var v25 uint8 = (^(v24))
                                                                _ = v25
                                                            case "hello":
                                                                if (v1 != (v1 % (((v1 & (v1 - 343))) | 1))) {
                                                                    if len(v2) > 0 {
                                                                        v2[0] = fi16
                                                                    }
                                                                } else {
                                                                    var v26 uint64 = (v5 % (((v5 / ((463) | 1))) | 1))
                                                                    v23 = 3.14159
                                                                    if len(v4) >= 2 {
                                                                        v4 = v4[1:len(v4)]
                                                                    }
                                                                    _ = v26
                                                                }
                                                            case "a":
                                                                if len(v2) > 0 {
                                                                    v2[0] = -1
                                                                }
                                                            }
                                                            v3 = v3
                                                            v3 -= v3
                                                            _ = v23
                                                            _ = v24
                                                        }
                                                    case -49:
                                                        var v27 int8 = 4
                                                        fmt.Println(v27, v27)
                                                        switch "a" {
                                                        case "\n\t":
                                                            switch (fi15 ^ (fi16 & 32768)) {
                                                            case 2147483647:
                                                                var v28 int16 = 1
                                                                _ = v28
                                                            default:
                                                                var v29 int = fi15
                                                                _ = v29
                                                            }
                                                        }
                                                        var v30 int16 = 32767
                                                        v30 += 777
                                                        if (!false) {
                                                            v31 := GenS3{F0: v3}
                                                            switch 127 {
                                                            case -891:
                                                                var v32 uint16 = 32767
                                                                _ = v32
                                                            case 2147483647:
                                                                v33 := GenS4{F0: 1, F1: 2, F2: v1}
                                                                _ = v33
                                                            case 1:
                                                                fmt.Println(fi16, v3)
                                                            default:
                                                                if false {
                                                                    if true {
                                                                        v34 := GenS5{F0: 854, F1: (^(v5))}
                                                                        v27 = (v27 - (v27 >> ((0) & 63)))
                                                                        fmt.Println(v30)
                                                                        _ = v34
                                                                    } else {
                                                                        switch ("foo bar" + "hello") {
                                                                        case "\n\t":
                                                                            if false {
                                                                                v30 += (v30 % ((32767) | 1))
                                                                            } else {
                                                                                if false {
                                                                                    v35 := GenS5{F0: 267, F1: v5}
                                                                                    _ = v35
                                                                                }
                                                                            }
                                                                        default:
                                                                            if len(v4) > 0 {
                                                                                v4[0] = 119
                                                                            }
                                                                        }
                                                                        if false {
                                                                            v4 = append(v4, v3)
                                                                            var v36 uint16 = 0
                                                                            _ = v36
                                                                        } else {
                                                                            v37 := GenS6{F0: (v30 / (((v30 % ((v30) | 1))) | 1)), F1: 4294967295, F2: true}
                                                                            var v38 int64 = v1
                                                                            _ = v37
                                                                            _ = v38
                                                                        }
                                                                        var v39 string = ("\n\t" + "")
                                                                        _ = v39
                                                                    }
                                                                }
                                                            }
                                                            _ = len(v2)
                                                            _ = v31
                                                        } else {
                                                            _ = len(v2)
                                                            v27 = 67
                                                        }
                                                        v3 -= (^(v3))
                                                        _ = v27
                                                        _ = v30
                                                    case 0:
                                                        if len(v4) >= 2 {
                                                            v4 = v4[1:len(v4)]
                                                        }
                                                    }
                                                    v3 = v3
                                                    if (v1 >= -1) {
                                                        v5 += v5
                                                        v40 := GenS4{F0: 51, F1: 702, F2: (v1 * int64(v3))}
                                                        fmt.Println(v1)
                                                        _ = v40
                                                    } else {
                                                        var v41 int16 = 1
                                                        _ = v41
                                                    }
                                                    if false {
                                                        switch ("\n\t" + (("a" + ((("a" + "\n\t") + "foo bar") + (("a" + "hello") + ("a" + "hello")))) + "\n\t")) {
                                                        case "\n\t":
                                                            var v42 bool = false
                                                            _ = v42
                                                        default:
                                                            v3 = v3
                                                        }
                                                        v4 = append(v4, 2147483647)
                                                        var v43 float64 = -2.25
                                                        var v44 float32 = 0.0
                                                        v45 := GenS7{F0: 847}
                                                        _ = v43
                                                        _ = v44
                                                        _ = v45
                                                    } else {
                                                        switch (fi15 - 1) {
                                                        case 826:
                                                            v1 -= (v1 / ((9223372036854775807) | 1))
                                                        case -9223372036854775808:
                                                            var v46 int16 = 1000
                                                            _ = v46
                                                        default:
                                                            fmt.Println(fi15)
                                                        }
                                                        var v47 int = (^(fi15))
                                                        _ = len(v2)
                                                        v48 := GenS2{F0: (v5 * v5)}
                                                        _ = v47
                                                        _ = v48
                                                    }
                                                } else {
                                                    var v49 float64 = 3.14159
                                                    _ = v49
                                                }
                                            }
                                        default:
                                            v50 := GenS6{F0: -32768, F1: 255, F2: false}
                                            _ = v50
                                        }
                                        if (true || ((255 >= 42) && (true && true))) {
                                            v3 -= 2
                                        } else {
                                            switch "foo bar" {
                                            case "":
                                                v51 := GenS4{F0: 366, F1: 4294967295, F2: (v1 << ((42) & 63))}
                                                _ = v51
                                            case "a":
                                                v52 := GenS8{F0: 1}
                                                _ = v52
                                            case "hello":
                                                switch (("\n\t" + "foo bar") + "hello") {
                                                case "":
                                                    v53 := GenS6{F0: -1, F1: 595, F2: true}
                                                    _ = v53
                                                case "hello":
                                                    switch (fi15 ^ (-(fi15))) {
                                                    case 127:
                                                        switch "foo bar" {
                                                        case "a":
                                                            v3 = 1
                                                        case "hello":
                                                            _ = len(v4)
                                                            if len(v2) >= 2 {
                                                                v2 = v2[1:len(v2)]
                                                            }
                                                            switch fi15 {
                                                            case 9223372036854775807:
                                                                switch "a" {
                                                                case "foo bar":
                                                                    var v54 uint8 = 22
                                                                    _ = v54
                                                                case "hello":
                                                                    if ("hello" == "a") {
                                                                        if (true || false) {
                                                                            var v55 int = (^(fi16))
                                                                            switch v55 {
                                                                            case 9223372036854775807:
                                                                                v5 -= 4294967295
                                                                            case -2147483648:
                                                                                v55 -= fi16
                                                                            }
                                                                            _ = v55
                                                                        }
                                                                        var v56 uint = 796
                                                                        fmt.Println(fi16, fi16, v3)
                                                                        _ = v56
                                                                    } else {
                                                                        if len(v2) > 0 {
                                                                            v2[0] = 32768
                                                                        }
                                                                        if false {
                                                                            fmt.Println(fi16, v3)
                                                                            if len(v2) > 0 {
                                                                                v2[0] = (fi15 - fi16)
                                                                            }
                                                                            v5 += 521
                                                                        }
                                                                    }
                                                                }
                                                            case 1:
                                                                v2 = append(v2, fi15)
                                                            case 0:
                                                                if false {
                                                                    v5 = v5
                                                                    var v57 int16 = 1
                                                                    v58 := GenS9{F0: -2147483648, F1: v3}
                                                                    _ = v57
                                                                    _ = v58
                                                                } else {
                                                                    v59 := GenS7{F0: 65535}
                                                                    _ = v59
                                                                }
                                                            default:
                                                                var v60 int64 = (-(v1))
                                                                _ = v60
                                                            }
                                                        default:
                                                            fmt.Println(fi16, v1, v3)
                                                        }
                                                    case -1:
                                                        v1 += v1
                                                    case 32768:
                                                        if true {
                                                            var v61 uint16 = 42
                                                            if (v3 >= 59) {
                                                                switch ("foo bar" + ("hello" + "a")) {
                                                                case "hello":
                                                                    v1 = (v1 & v1)
                                                                case "\n\t":
                                                                    var v62 bool = (1 >= -128)
                                                                    _ = v62
                                                                }
                                                                v63 := GenS10{F0: v3, F1: 1}
                                                                if (2147483647 != 0) {
                                                                    v64 := GenS11{F0: 137, F1: true, F2: 0}
                                                                    _ = v64
                                                                } else {
                                                                    switch (fi16 - (fi15 - fi16)) {
                                                                    case 0:
                                                                        switch "a" {
                                                                        case "foo bar":
                                                                            v61 += (v61 ^ v61)
                                                                        default:
                                                                            if (!(102 <= 0)) {
                                                                                v61 += 1
                                                                                var v65 string = ""
                                                                                _ = v65
                                                                            } else {
                                                                                v61 -= v61
                                                                                fmt.Println(v61, fi15, fi15)
                                                                                var v66 int8 = 17
                                                                                _ = v66
                                                                            }
                                                                        }
                                                                    case 9223372036854775807:
                                                                        var v67 float64 = 1e10
                                                                        _ = v67
                                                                    case 1:
                                                                        switch "foo bar" {
                                                                        case "foo bar":
                                                                            v1 -= v1
                                                                        case "\n\t":
                                                                            v5 += (v5 | v5)
                                                                        case "":
                                                                            if len(v4) >= 2 {
                                                                                v4 = v4[1:len(v4)]
                                                                            }
                                                                        default:
                                                                            if true {
                                                                                v61 += (v61 | v61)
                                                                            } else {
                                                                                fmt.Println(v1, v61)
                                                                                fmt.Println(fi15, fi15)
                                                                            }
                                                                        }
                                                                    }
                                                                    var v68 uint64 = v5
                                                                    v1 = (v1 << ((764) & 63))
                                                                    _ = v68
                                                                }
                                                                v69 := GenS5{F0: (v61 ^ 0), F1: (v5 + (v5 >> ((821) & 63)))}
                                                                fmt.Println(v1, fi15)
                                                                _ = v63
                                                                _ = v69
                                                            } else {
                                                                var v70 float64 = 1.5
                                                                fmt.Println(v3, v70, v70)
                                                                if false {
                                                                    v71 := GenS12{F0: v3, F1: 139}
                                                                    _ = v71
                                                                } else {
                                                                    v61 = v61
                                                                }
                                                                v4 = append(v4, (^(v3)))
                                                                switch 9223372036854775807 {
                                                                case 127:
                                                                    switch (fi15 - (fi15 % (((fi16 - fi15)) | 1))) {
                                                                    case 635:
                                                                        switch (fi16 % (((^(fi16))) | 1)) {
                                                                        case 32768:
                                                                            var v72 int = (fi15 ^ fi15)
                                                                            _ = v72
                                                                        case -2147483648:
                                                                            if false {
                                                                                if len(v2) > 0 {
                                                                                    v2[0] = -1
                                                                                }
                                                                            } else {
                                                                                fmt.Println(v70, fi15)
                                                                                _ = len(v2)
                                                                                if (true && (((v61 - 976) > v61) && false)) {
                                                                                    v73 := GenS3{F0: v3}
                                                                                    if len(v2) > 0 {
                                                                                        v2[0] = (-(fi16))
                                                                                    }
                                                                                    switch 1 {
                                                                                    case -2147483648:
                                                                                        v4 = append(v4, 2147483647)
                                                                                    case 1:
                                                                                        switch (fi16 << ((240) & 63)) {
                                                                                        case 2147483647:
                                                                                            var v74 uint16 = (v61 ^ v61)
                                                                                            _ = v74
                                                                                        case 867:
                                                                                            var v75 int8 = -128
                                                                                            _ = v75
                                                                                        case 32768:
                                                                                            v1 -= (v1 & (v1 >> ((uint(v5)) & 63)))
                                                                                        }
                                                                                    default:
                                                                                        v3 += (v3 + v3)
                                                                                    }
                                                                                    _ = v73
                                                                                }
                                                                            }
                                                                        case 127:
                                                                            switch (fi15 + 0) {
                                                                            case -132:
                                                                                if (-128 < -128) {
                                                                                    if false {
                                                                                        fmt.Println(v3, v61, v5)
                                                                                        switch -153 {
                                                                                        case 32767:
                                                                                            v76 := GenS6{F0: 42, F1: 260, F2: (774 != 65535)}
                                                                                            _ = v76
                                                                                        }
                                                                                    } else {
                                                                                        v77 := GenS8{F0: v61}
                                                                                        if len(v4) >= 2 {
                                                                                            v4 = v4[1:len(v4)]
                                                                                        }
                                                                                        if false {
                                                                                            fmt.Println(v61, v61)
                                                                                            var v78 bool = (!true)
                                                                                            if len(v4) > 0 {
                                                                                                v4[0] = v3
                                                                                            }
                                                                                            if true {
                                                                                                fmt.Println(fi16)
                                                                                                v79 := GenS13{F0: -1, F1: -43, F2: 84}
                                                                                                _ = len(v2)
                                                                                                _ = v79
                                                                                            } else {
                                                                                                var v80 string = "foo bar"
                                                                                                var v81 float64 = v70
                                                                                                if (v61 < 65535) {
                                                                                                    fmt.Println(v81, v81)
                                                                                                    v3 += 42
                                                                                                    v2 = append(v2, (-(fi15)))
                                                                                                    v82 := GenS14{F0: false, F1: 0, F2: (v61 >> ((42) & 63))}
                                                                                                    _ = v82
                                                                                                } else {
                                                                                                    var v83 uint32 = (v3 << ((42) & 63))
                                                                                                    switch fi16 {
                                                                                                    case -113:
                                                                                                        v70 = v70
                                                                                                    }
                                                                                                    v80 = ""
                                                                                                    _ = v83
                                                                                                }
                                                                                                _ = v80
                                                                                                _ = v81
                                                                                            }
                                                                                            _ = v78
                                                                                        }
                                                                                        _ = v77
                                                                                    }
                                                                                    var v84 uint32 = v3
                                                                                    _ = v84
                                                                                }
                                                                            case -9223372036854775808:
                                                                                if false {
                                                                                    if len(v4) >= 2 {
                                                                                        v4 = v4[1:len(v4)]
                                                                                    }
                                                                                    if ("a" >= "a") {
                                                                                        if (!false) {
                                                                                            _ = len(v4)
                                                                                        } else {
                                                                                            switch int(v5) {
                                                                                            case 348:
                                                                                                v61 -= 32767
                                                                                            }
                                                                                            v5 += (v5 / (((^(v5))) | 1))
                                                                                            if false {
                                                                                                fmt.Println(v70, v3)
                                                                                                switch "hello" {
                                                                                                case "":
                                                                                                    if true {
                                                                                                        var v85 float64 = (((float64((1.5 + 1.5)) + v70) / 3.14159) / -2.25)
                                                                                                        v61 = (v61 & (v61 >> ((4294967295) & 63)))
                                                                                                        if ((float64(1.5) - v70) <= (1.5 * 1e10)) {
                                                                                                            fmt.Println(v5, v3, v5)
                                                                                                            fmt.Println(v5, fi15)
                                                                                                            if len(v2) >= 2 {
                                                                                                                v2 = v2[1:len(v2)]
                                                                                                            }
                                                                                                            v85 = -2.25
                                                                                                        } else {
                                                                                                            if false {
                                                                                                                if len(v4) > 0 {
                                                                                                                    v4[0] = (v3 + 987)
                                                                                                                }
                                                                                                            } else {
                                                                                                                var v86 int32 = 2147483647
                                                                                                                fmt.Println(v5, v61, v1)
                                                                                                                v87 := GenS15{F0: ("hello" + "a")}
                                                                                                                _ = v86
                                                                                                                _ = v87
                                                                                                            }
                                                                                                        }
                                                                                                        _ = v85
                                                                                                    } else {
                                                                                                        fmt.Println(fi15)
                                                                                                        fmt.Println(v1, v1, fi15)
                                                                                                    }
                                                                                                case "hello":
                                                                                                    var v88 int64 = (v1 - 1)
                                                                                                    _ = v88
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                        v61 = (v61 ^ v61)
                                                                                        v1 = 42
                                                                                    }
                                                                                    var v89 string = "hello"
                                                                                    _ = v89
                                                                                }
                                                                            default:
                                                                                v90 := GenS13{F0: -119, F1: 100, F2: 1}
                                                                                _ = len(v2)
                                                                                if (644 <= 1) {
                                                                                    v70 -= (((v70 - 0.0) * 3.14159) + ((v70 + (v70 / float64(3.14159))) - 3.14159))
                                                                                }
                                                                                fmt.Println(v70)
                                                                                if true {
                                                                                    if len(v4) > 0 {
                                                                                        v4[0] = v3
                                                                                    }
                                                                                } else {
                                                                                    if ((((fi15 % ((93) | 1)) < int(v5)) || false) || true) {
                                                                                        v3 = v3
                                                                                        var v91 bool = (0.1 > v70)
                                                                                        _ = v91
                                                                                    } else {
                                                                                        v92 := GenS2{F0: (v5 | uint64(v61))}
                                                                                        _ = v92
                                                                                    }
                                                                                }
                                                                                v70 = v70
                                                                                _ = v90
                                                                            }
                                                                        }
                                                                    case 1:
                                                                        var v93 string = (("hello" + "a") + ((("foo bar" + "\n\t") + "") + (("hello" + "\n\t") + "")))
                                                                        _ = v93
                                                                    }
                                                                case -494:
                                                                    if false {
                                                                        var v94 int16 = -948
                                                                        _ = v94
                                                                    }
                                                                }
                                                                _ = v70
                                                            }
                                                            fmt.Println(v1, v5)
                                                            _ = v61
                                                        } else {
                                                            var v95 bool = true
                                                            _ = len(v2)
                                                            v3 = 65535
                                                            _ = v95
                                                        }
                                                    }
                                                case "\n\t":
                                                    if (!false) {
                                                        if false {
                                                            v96 := GenS12{F0: (v3 % ((v3) | 1)), F1: v5}
                                                            _ = v96
                                                        } else {
                                                            _ = len(v2)
                                                            var v97 float64 = -2.25
                                                            _ = v97
                                                        }
                                                    }
                                                }
                                            default:
                                                var v98 int = (fi15 / (((fi15 ^ -755)) | 1))
                                                _ = v98
                                            }
                                            v5 = v5
                                        }
                                    }
                                    var v99 int32 = -1
                                    _ = v99
                                } else {
                                    v1 = v1
                                    if false {
                                        switch fi15 {
                                        case 32767:
                                            v5 = v5
                                        }
                                    } else {
                                        fmt.Println(v1, fi16, fi16)
                                        var v100 int32 = -2147483648
                                        switch "foo bar" {
                                        case "hello":
                                            var v101 int64 = 0
                                            _ = v101
                                        case "":
                                            var v102 uint8 = uint8(v5)
                                            _ = v102
                                        case "\n\t":
                                            fmt.Println(v1)
                                        default:
                                            switch fi15 {
                                            case 42:
                                                if false {
                                                    if true {
                                                        var v103 float64 = 0.1
                                                        _ = v103
                                                    } else {
                                                        if len(v4) > 0 {
                                                            v4[0] = (^(v3))
                                                        }
                                                    }
                                                    var v104 float32 = 0.1
                                                    _ = v104
                                                }
                                            case 128:
                                                var v105 int8 = 42
                                                _ = v105
                                            case 1:
                                                v106 := GenS13{F0: -76, F1: 42, F2: 42}
                                                _ = v106
                                            }
                                        }
                                        if ((0 < 32767) && ((1 < 255) && (!(v100 >= v100)))) {
                                            if false {
                                                var v107 uint16 = 466
                                                var v108 float32 = 1.5
                                                _ = v107
                                                _ = v108
                                            } else {
                                                var v109 uint16 = 462
                                                var v110 uint8 = 127
                                                _ = v109
                                                _ = v110
                                            }
                                        } else {
                                            fmt.Println(fi16)
                                        }
                                        var v111 uint8 = 170
                                        _ = v100
                                        _ = v111
                                    }
                                    v112 := GenS16{F0: 149, F1: 0, F2: v1}
                                    _ = v112
                                }
                                fmt.Println(v1)
                                if len(v4) > 0 {
                                    v4[0] = v3
                                }
                                _ = fi16
                            }
                            fmt.Println(fi15, v3, fi15)
                            if (2 != 42) {
                                switch ("hello" + "a") {
                                case "\n\t":
                                    v113 := GenS12{F0: (v3 % (((v3 - 2)) | 1)), F1: (v5 - v5)}
                                    _ = v113
                                case "foo bar":
                                    for fi114 := 0; fi114 < 1; fi114++ {
                                        fmt.Println(v1, v5, v1)
                                        _ = fi114
                                    }
                                case "hello":
                                    switch 1 {
                                    case 0:
                                        var v115 uint32 = 2147483648
                                        _ = v115
                                    case 32768:
                                        var v116 int16 = -1
                                        _ = v116
                                    default:
                                        var v117 uint64 = 4294967295
                                        _ = v117
                                    }
                                default:
                                    var v118 int64 = (v1 << ((18446744073709551615) & 63))
                                    _ = v118
                                }
                                v119 := GenS15{F0: ("hello" + "foo bar")}
                                v120 := GenS2{F0: (v5 ^ (^(v5)))}
                                _ = v119
                                _ = v120
                            }
                            _ = fi15
                        }
                    default:
                        if ((v1 + 0) != -9223372036854775808) {
                            for fi121 := 0; fi121 < 14; fi121++ {
                                var v122 uint32 = v3
                                _ = len(v2)
                                if false {
                                    if true {
                                        if len(v4) > 0 {
                                            v4[0] = 553
                                        }
                                        v123 := GenS17{F0: (v5 << ((57) & 63)), F1: 4294967295, F2: (100 <= 34)}
                                        if true {
                                            if true {
                                                if false {
                                                    _ = len(v4)
                                                } else {
                                                    if true {
                                                        v4 = append(v4, v122)
                                                    } else {
                                                        switch "" {
                                                        case "a":
                                                            var v124 uint16 = 32767
                                                            _ = v124
                                                        case "hello":
                                                            for fi125 := 0; fi125 < 19; fi125++ {
                                                                var v126 float32 = 1e10
                                                                if (v5 <= (v5 - (v5 * v5))) {
                                                                    v127 := GenS18{F0: 67, F1: (^(v1))}
                                                                    _ = v127
                                                                } else {
                                                                    v5 = 18446744073709551615
                                                                    if len(v2) >= 2 {
                                                                        v2 = v2[1:len(v2)]
                                                                    }
                                                                    if false {
                                                                        v3 += v122
                                                                    }
                                                                    v128 := GenS19{F0: 42}
                                                                    _ = v128
                                                                }
                                                                _ = fi125
                                                                _ = v126
                                                            }
                                                        case "":
                                                            var v129 string = "\n\t"
                                                            _ = v129
                                                        }
                                                        var v130 uint64 = 4294967295
                                                        fmt.Println(v5, v130, v3)
                                                        _ = v130
                                                    }
                                                }
                                                var v131 uint32 = (v3 - (v122 >> ((255) & 63)))
                                                _ = v131
                                            } else {
                                                var v132 int16 = int16(v1)
                                                var v133 bool = (-2147483648 == 1)
                                                fmt.Println(v133, v1, v132)
                                                _ = v132
                                                _ = v133
                                            }
                                            if (18446744073709551615 > 18446744073709551615) {
                                                var v134 string = "\n\t"
                                                if (127 < -1) {
                                                    switch v134 {
                                                    case "hello":
                                                        v5 = v5
                                                    case "\n\t":
                                                        fmt.Println(fi121)
                                                    case "a":
                                                        var v135 uint16 = 0
                                                        _ = v135
                                                    }
                                                    fmt.Println(v5)
                                                }
                                                for fi136 := 0; fi136 < 3; fi136++ {
                                                    var v137 float32 = (0.1 * float32(0.1))
                                                    var v138 int32 = 0
                                                    _ = fi136
                                                    _ = v137
                                                    _ = v138
                                                }
                                                _ = v134
                                            }
                                            v139 := GenS15{F0: "\n\t"}
                                            _ = v139
                                        }
                                        _ = v123
                                    } else {
                                        v1 = v1
                                        if len(v2) > 0 {
                                            v2[0] = (fi121 >> ((399) & 63))
                                        }
                                    }
                                    for fi140 := 0; fi140 < 2; fi140++ {
                                        var v141 uint64 = v5
                                        v3 = v122
                                        _ = fi140
                                        _ = v141
                                    }
                                    v122 -= 42
                                } else {
                                    for fi142 := 0; fi142 < 16; fi142++ {
                                        if len(v4) > 0 {
                                            v4[0] = (v3 ^ v122)
                                        }
                                        v143 := GenS20{F0: 1, F1: 65535}
                                        v144 := GenS1{F0: 164, F1: (v1 != 378)}
                                        _ = fi142
                                        _ = v143
                                        _ = v144
                                    }
                                }
                                _ = fi121
                                _ = v122
                            }
                        } else {
                            if false {
                                if (("foo bar" == "foo bar") && (0 < uint16(v5))) {
                                    v3 = uint32(v5)
                                    for fi145 := 0; fi145 < 15; fi145++ {
                                        v146 := GenS15{F0: "a"}
                                        v3 = v3
                                        switch ("" + "\n\t") {
                                        case "a":
                                            if true {
                                                if len(v4) > 0 {
                                                    v4[0] = (^(v3))
                                                }
                                                v5 = (v5 >> ((65535) & 63))
                                            }
                                        case "hello":
                                            var v147 int8 = int8(fi145)
                                            _ = v147
                                        case "foo bar":
                                            v5 = 2
                                        default:
                                            for fi148 := 0; fi148 < 7; fi148++ {
                                                _ = len(v2)
                                                fmt.Println(v1, v3)
                                                var v149 int8 = 1
                                                v1 -= (^(v1))
                                                if true {
                                                    v149 = v149
                                                    var v150 int8 = (v149 | int8(fi148))
                                                    _ = v150
                                                } else {
                                                    v1 -= v1
                                                    if (v3 >= v3) {
                                                        if true {
                                                            var v151 uint16 = 773
                                                            if false {
                                                                v152 := GenS12{F0: v3, F1: (v5 % ((v5) | 1))}
                                                                _ = v152
                                                            }
                                                            _ = v151
                                                        }
                                                        v153 := GenS12{F0: (^(v3)), F1: v5}
                                                        fmt.Println(v1, v1)
                                                        _ = v153
                                                    }
                                                }
                                                if ((-861 != 495) && false) {
                                                    v154 := GenS13{F0: v149, F1: v149, F2: 40}
                                                    var v155 float32 = 1.5
                                                    if false {
                                                        v1 -= (v1 / ((v1) | 1))
                                                        v2 = append(v2, 9223372036854775807)
                                                    } else {
                                                        fmt.Println(v5, v149, v155)
                                                        var v156 float32 = v155
                                                        _ = v156
                                                    }
                                                    var v157 uint32 = 4294967295
                                                    v158 := GenS21{F0: (fi145 / ((974) | 1))}
                                                    _ = v154
                                                    _ = v155
                                                    _ = v157
                                                    _ = v158
                                                } else {
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                    var v159 int16 = -505
                                                    _ = v159
                                                }
                                                if ((("a" + "foo bar") + ("foo bar" + "\n\t")) == ("hello" + "\n\t")) {
                                                    if ((!false) && true) {
                                                        v160 := GenS22{F0: (!true), F1: (fi148 % ((fi148) | 1)), F2: 229}
                                                        _ = v160
                                                    } else {
                                                        v1 = (v1 >> ((2) & 63))
                                                        v2 = append(v2, (fi148 % ((fi148) | 1)))
                                                    }
                                                    if len(v2) > 0 {
                                                        v2[0] = (fi148 * (fi145 * 770))
                                                    }
                                                    fmt.Println(v5, v3, v3)
                                                } else {
                                                    switch "\n\t" {
                                                    case "":
                                                        v1 = (v1 ^ -568)
                                                    case "a":
                                                        if true {
                                                            var v161 int64 = 9223372036854775807
                                                            v162 := GenS17{F0: 0, F1: 42, F2: true}
                                                            _ = v161
                                                            _ = v162
                                                        } else {
                                                            var v163 uint = 132
                                                            var v164 int64 = int64(v149)
                                                            switch "" {
                                                            case "":
                                                                var v165 bool = ((v5 * uint64(v3)) < v5)
                                                                _ = v165
                                                            case "a":
                                                                v2 = append(v2, (fi145 << ((v163) & 63)))
                                                            case "hello":
                                                                v163 = v163
                                                                var v166 int16 = 42
                                                                _ = v166
                                                            }
                                                            _ = v163
                                                            _ = v164
                                                        }
                                                    case "\n\t":
                                                        var v167 uint64 = (v5 >> ((uint(v5)) & 63))
                                                        v3 = (v3 * uint32(v5))
                                                        _ = v167
                                                    default:
                                                        var v168 uint = 2
                                                        _ = v168
                                                    }
                                                }
                                                _ = fi148
                                                _ = v149
                                            }
                                        }
                                        _ = fi145
                                        _ = v146
                                    }
                                }
                                for fi169 := 0; fi169 < 15; fi169++ {
                                    var v170 uint32 = (^(v3))
                                    _ = fi169
                                    _ = v170
                                }
                                _ = len(v4)
                            }
                            var v171 uint = 578
                            for fi172 := 0; fi172 < 3; fi172++ {
                                v3 = v3
                                v171 = v171
                                _ = fi172
                            }
                            _ = v171
                        }
                    }
                    fmt.Println(v5)
                    if len(v2) > 0 {
                        v2[0] = 42
                    }
                }
            default:
                fmt.Println(v3)
            }
        case 23:
            if true {
                switch -9223372036854775808 {
                case 127:
                    if false {
                        var v173 string = "foo bar"
                        for fi174 := 0; fi174 < 3; fi174++ {
                            fmt.Println(v5, v1)
                            for fi175 := 0; fi175 < 8; fi175++ {
                                if false {
                                    v173 = "hello"
                                }
                                v1 = v1
                                switch "\n\t" {
                                case "\n\t":
                                    v173 = ((v173 + ("a" + v173)) + ("hello" + ("hello" + "foo bar")))
                                case "hello":
                                    var v176 uint = 617
                                    _ = v176
                                case "foo bar":
                                    v177 := GenS12{F0: (v3 | 2), F1: (v5 * (v5 ^ v5))}
                                    _ = v177
                                }
                                switch 2147483647 {
                                case 2147483647:
                                    v173 = "foo bar"
                                case -9223372036854775808:
                                    fmt.Println(v5, fi174, v173)
                                }
                                if (v1 <= v1) {
                                    v178 := GenS23{F0: 0, F1: (^(v5))}
                                    fmt.Println(v173, v173)
                                    switch (fi175 / (((fi175 + fi175)) | 1)) {
                                    case 722:
                                        switch (fi175 % ((fi174) | 1)) {
                                        case 538:
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                        }
                                    }
                                    _ = v178
                                }
                                _ = fi175
                            }
                            _ = fi174
                        }
                        _ = len(v2)
                        _ = v173
                    } else {
                        v179 := GenS24{F0: "a", F1: true, F2: v3}
                        var v180 float64 = 1.5
                        _ = v179
                        _ = v180
                    }
                case 9223372036854775807:
                    if (631 != 255) {
                        if (v5 > v5) {
                            switch "\n\t" {
                            case "hello":
                                v181 := GenS6{F0: -1, F1: 1, F2: true}
                                _ = v181
                            case "":
                                v5 = v5
                            }
                            v5 += (v5 / (((v5 << ((616) & 63))) | 1))
                        } else {
                            v2 = append(v2, 66)
                            switch "hello" {
                            case "\n\t":
                                v182 := GenS23{F0: -326, F1: v5}
                                _ = v182
                            case "foo bar":
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            }
                        }
                    } else {
                        v1 = v1
                    }
                }
                var v183 float64 = 0.1
                v3 += v3
                _ = v183
            } else {
                v184 := GenS22{F0: false, F1: 9223372036854775807, F2: 116}
                _ = v184
            }
        default:
            v185 := GenS15{F0: ("\n\t" + "a")}
            _ = v185
        }
    }
    fmt.Println(v1)
    for fi186 := 0; fi186 < 12; fi186++ {
        var v187 uint16 = 42
        var v188 int64 = v1
        fmt.Println(v187, fi186, v187)
        _ = fi186
        _ = v187
        _ = v188
    }
    v3 = (^(v3))
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v189 uint = uint(v5)
    switch "" {
    case "a":
        v189 -= v189
    }
    fmt.Println("v1:", v1)
    fmt.Println("v189:", v189)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v189
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
