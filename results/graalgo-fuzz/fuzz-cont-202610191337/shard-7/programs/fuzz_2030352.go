package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64 }
type GenS2 struct { F0 int16; F1 int16; F2 uint32 }
type GenS3 struct { F0 uint16; F1 uint16; F2 uint32 }
type GenS4 struct { F0 uint8; F1 int; F2 uint32 }
type GenS5 struct { F0 uint16; F1 bool }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 uint16; F1 uint32; F2 uint8 }
type GenS8 struct { F0 int16; F1 uint16 }
type GenS9 struct { F0 uint8 }
type GenS10 struct { F0 uint8; F1 uint32 }
type GenS11 struct { F0 uint16; F1 uint8; F2 int16 }
type GenS12 struct { F0 int32; F1 int64; F2 uint16 }
type GenS13 struct { F0 string; F1 uint32 }
type GenS14 struct { F0 string; F1 uint; F2 uint }
type GenS15 struct { F0 int }
type GenS16 struct { F0 int }
type GenS17 struct { F0 uint64; F1 int16; F2 string }
type GenS18 struct { F0 int16; F1 uint64 }
type GenS19 struct { F0 uint16; F1 int16 }
type GenS20 struct { F0 bool; F1 int8 }

func main() {
    var v1 uint = 52
    v2 := make(map[string]int)
    v3 := make(map[string]bool)
    var v4 int64 = 412
    var v5 uint = v1
    v1 = (v5 >> (((^(v5))) & 63))
    v6 := GenS1{F0: v4, F1: 0}
    for fi7 := 0; fi7 < 18; fi7++ {
        v4 = v4
        v8 := GenS1{F0: (v4 + v4), F1: 5}
        _ = fi7
        _ = v8
    }
    v9 := GenS2{F0: 42, F1: -1, F2: 0}
    v10 := GenS3{F0: 0, F1: 1, F2: 2147483648}
    _ = len(v3)
    if (255 != 168) {
        if true {
            fmt.Println(v1, v4)
            v11 := GenS4{F0: 42, F1: 387, F2: 468}
            _ = v11
        } else {
            for fi12 := 0; fi12 < 16; fi12++ {
                v1 = 296
                var v13 int = fi12
                _ = fi12
                _ = v13
            }
            var v14 string = (("hello" + ("hello" + "\n\t")) + "foo bar")
            _ = v14
        }
    } else {
        switch "" {
        case "foo bar":
            var v15 uint32 = 2
            _ = v15
        }
    }
    v5 = v1
    if (2147483647 == -2147483648) {
        for fi16 := 0; fi16 < 2; fi16++ {
            for fi17 := 0; fi17 < 3; fi17++ {
                if (1 != 2) {
                    if true {
                        v5 += v5
                        v18 := GenS5{F0: 255, F1: ((!(18446744073709551615 != 2)) || false)}
                        v19 := GenS6{F0: 65535}
                        v1 = (v5 | v5)
                        _ = v18
                        _ = v19
                    } else {
                        switch "\n\t" {
                        case "":
                            v20 := GenS7{F0: 1, F1: 2, F2: 0}
                            _ = v20
                        default:
                            var v21 uint16 = 42
                            _ = v21
                        }
                        if (!true) {
                            v22 := GenS2{F0: -32768, F1: 350, F2: 42}
                            switch (fi17 >> (((v1 - (v1 & v5))) & 63)) {
                            case 32768:
                                v23 := []int{(fi16 - (fi17 - (fi17 | fi17))), (fi16 | (fi17 >> ((65535) & 63))), (fi17 % ((-2147483648) | 1))}
                                _ = v23
                            default:
                                switch fi17 {
                                case 4294967295:
                                    var v24 uint8 = 42
                                    _ = v24
                                case -288:
                                    if (255 <= 2) {
                                        v4 = (v4 & (v4 >> (((v1 * 18446744073709551615)) & 63)))
                                    }
                                default:
                                    if false {
                                        if false {
                                            v25 := GenS2{F0: -1, F1: -1, F2: 947}
                                            v26 := v2[""]
                                            _ = v25
                                            _ = v26
                                        }
                                        switch (^(fi17)) {
                                        case 684:
                                            v4 += v4
                                        case -280:
                                            if ((1.5 * -2.25) != (1.5 - ((0.1 + 0.0) * float64((((((3.14159 - -2.25) + 0.0) - ((1.5 - 3.14159) + 1e10)) + 0.1) + -2.25))))) {
                                                var v27 uint16 = 0
                                                fmt.Println(v1, fi17)
                                                v4 = v4
                                                switch fi16 {
                                                case -9223372036854775808:
                                                    if ((-2.25 / (0.1 + 1.5)) != -2.25) {
                                                        switch ("a" + "a") {
                                                        case "a":
                                                            v1 = (v5 % (((v1 & v5)) | 1))
                                                        case "\n\t":
                                                            if false {
                                                                _ = len(v3)
                                                                switch fi16 {
                                                                case 2147483647:
                                                                    switch "a" {
                                                                    case "hello":
                                                                        var v28 int = fi16
                                                                        _ = v28
                                                                    default:
                                                                        var v29 string = "foo bar"
                                                                        _ = v29
                                                                    }
                                                                default:
                                                                    switch fi17 {
                                                                    case -33:
                                                                        switch (("hello" + "hello") + "a") {
                                                                        case "a":
                                                                            var v30 string = ("\n\t" + "foo bar")
                                                                            _ = v30
                                                                        case "foo bar":
                                                                            var v31 int16 = 32767
                                                                            _ = v31
                                                                        }
                                                                    case 2147483647:
                                                                        var v32 bool = (42 >= 0)
                                                                        _ = v32
                                                                    case 1:
                                                                        fmt.Println(v27, v1, fi17)
                                                                    default:
                                                                        v33 := []int{}
                                                                        _ = v33
                                                                    }
                                                                }
                                                            }
                                                        case "hello":
                                                            var v34 float64 = 3.14159
                                                            _ = v34
                                                        }
                                                        switch -1 {
                                                        case 940:
                                                            v35 := GenS8{F0: -32768, F1: (v27 ^ 2)}
                                                            if false {
                                                                if (425 < -22) {
                                                                    v36 := GenS3{F0: v27, F1: 2, F2: 176}
                                                                    _ = v36
                                                                } else {
                                                                    switch "\n\t" {
                                                                    case "\n\t":
                                                                        fmt.Println(fi17, v4, v27)
                                                                    default:
                                                                        v37 := []int{206, fi17}
                                                                        _ = v37
                                                                    }
                                                                }
                                                            }
                                                            _ = v35
                                                        default:
                                                            if (!(true || (!(208 == 2)))) {
                                                                fmt.Println(v1, fi16)
                                                                if (false && (1e10 == -2.25)) {
                                                                    var v38 bool = ((67 == 4294967295) && true)
                                                                    _ = v38
                                                                } else {
                                                                    v1 = 1
                                                                    if false {
                                                                        fmt.Println(fi17, v1)
                                                                        switch ("a" + "a") {
                                                                        case "":
                                                                            var v39 uint64 = uint64(v27)
                                                                            _ = v39
                                                                        default:
                                                                            var v40 int64 = v4
                                                                            v4 -= v40
                                                                            var v41 uint32 = 2
                                                                            v41 = (v41 - v41)
                                                                            if (!true) {
                                                                                if (!(255 >= 255)) {
                                                                                    v42 := []int{(-(fi16))}
                                                                                    _ = v42
                                                                                }
                                                                                switch 32768 {
                                                                                case 42:
                                                                                    if (0.0 != (1.5 * ((((1.5 / ((1e10 + -2.25) - 0.1)) * (3.14159 + 1e10)) * -2.25) / 1.0))) {
                                                                                        fmt.Println(fi17)
                                                                                        v5 = (v1 | 42)
                                                                                        switch "foo bar" {
                                                                                        case "\n\t":
                                                                                            fmt.Println(fi16, fi16)
                                                                                        case "":
                                                                                            if (3.14159 >= (0.1 / 1.0)) {
                                                                                                fmt.Println(v1)
                                                                                                v43 := []int{(^(fi17))}
                                                                                                v1 += 65535
                                                                                                v44 := GenS6{F0: 162}
                                                                                                _ = v43
                                                                                                _ = v44
                                                                                            }
                                                                                        case "a":
                                                                                            v2["\n\t"] = (fi17 | (fi16 << ((v5) & 63)))
                                                                                        }
                                                                                    }
                                                                                case -207:
                                                                                    if (v27 != (v27 * 986)) {
                                                                                        if (((v40 * 0) <= v40) && false) {
                                                                                            fmt.Println(v27, fi16, v40)
                                                                                            var v45 uint64 = 4294967295
                                                                                            _ = v45
                                                                                        }
                                                                                        var v46 string = (("a" + ("hello" + ("a" + "hello"))) + "\n\t")
                                                                                        v46 = (v46 + ("a" + (v46 + "foo bar")))
                                                                                        if false {
                                                                                            v3[((((v46 + v46) + v46) + ((v46 + v46) + (v46 + v46))) + "foo bar")] = (-128 > -39)
                                                                                            v47 := GenS9{F0: uint8(v5)}
                                                                                            v48 := []int{}
                                                                                            var v49 int8 = -121
                                                                                            _ = v47
                                                                                            _ = v48
                                                                                            _ = v49
                                                                                        }
                                                                                        if (v5 < v5) {
                                                                                            v1 = v5
                                                                                            fmt.Println(v41, v46, v4)
                                                                                        }
                                                                                        _ = v46
                                                                                    }
                                                                                case 32768:
                                                                                    v27 -= v27
                                                                                default:
                                                                                    v41 -= 65535
                                                                                    v40 += -9223372036854775808
                                                                                    fmt.Println(v27)
                                                                                    if (0.0 <= (3.14159 / 3.14159)) {
                                                                                        var v50 float32 = 0.0
                                                                                        if (-32768 == 410) {
                                                                                            v1 -= uint(v27)
                                                                                        }
                                                                                        v5 -= 1
                                                                                        var v51 uint = (v1 ^ v5)
                                                                                        delete(v3, ("a" + ((("\n\t" + ("foo bar" + "a")) + ("a" + "a")) + ("hello" + ""))))
                                                                                        _ = v50
                                                                                        _ = v51
                                                                                    }
                                                                                    var v52 string = ("foo bar" + "foo bar")
                                                                                    if (2 != 18446744073709551615) {
                                                                                        if (!true) {
                                                                                            v53 := GenS5{F0: (^(v27)), F1: ((1e10 > (float64((float32(3.14159) * 1e10)) * (1.5 / (1.5 + 3.14159)))) || (65535 < v27))}
                                                                                            v54 := GenS10{F0: 127, F1: v41}
                                                                                            switch v52 {
                                                                                            case "\n\t":
                                                                                                fmt.Println(v52)
                                                                                            case "":
                                                                                                switch (fi16 << ((18446744073709551615) & 63)) {
                                                                                                case 2147483647:
                                                                                                    v55 := GenS11{F0: 65535, F1: 255, F2: -1}
                                                                                                    _ = v55
                                                                                                case 32767:
                                                                                                    switch (fi17 + fi16) {
                                                                                                    case 551:
                                                                                                        if true {
                                                                                                            if true {
                                                                                                                if (((0 <= 2) || ((32767 >= 65) && false)) && (-2.25 < 0.1)) {
                                                                                                                    var v56 int = fi17
                                                                                                                    _ = v56
                                                                                                                }
                                                                                                                v57 := GenS12{F0: 0, F1: (v40 + 0), F2: v27}
                                                                                                                if (v1 != v1) {
                                                                                                                    v58 := GenS4{F0: 0, F1: fi16, F2: 320}
                                                                                                                    _ = v58
                                                                                                                }
                                                                                                                var v59 uint32 = (^(v41))
                                                                                                                _ = v57
                                                                                                                _ = v59
                                                                                                            }
                                                                                                            switch ((v52 + (("foo bar" + "") + ("a" + "\n\t"))) + (v52 + v52)) {
                                                                                                            case "foo bar":
                                                                                                                delete(v2, (("" + "\n\t") + ""))
                                                                                                            }
                                                                                                            v60 := GenS13{F0: "hello", F1: v41}
                                                                                                            _ = v60
                                                                                                        } else {
                                                                                                            v61 := v3[("a" + v52)]
                                                                                                            if (!((v27 - uint16(v5)) == v27)) {
                                                                                                                v62 := GenS13{F0: (((((v52 + v52) + "\n\t") + v52) + v52) + ((((("\n\t" + "") + ((v52 + v52) + (v52 + v52))) + v52) + "\n\t") + "\n\t")), F1: (v41 << ((v1) & 63))}
                                                                                                                var v63 float32 = (float32(0.0) * -2.25)
                                                                                                                _ = v62
                                                                                                                _ = v63
                                                                                                            }
                                                                                                            v64 := GenS14{F0: v52, F1: 65535, F2: v5}
                                                                                                            _ = v61
                                                                                                            _ = v64
                                                                                                        }
                                                                                                    case 0:
                                                                                                        delete(v2, v52)
                                                                                                    case -1:
                                                                                                        switch int(v40) {
                                                                                                        case 32767:
                                                                                                            v65 := GenS12{F0: 2147483647, F1: (v40 % (((v40 << (((v1 << ((v1) & 63))) & 63))) | 1)), F2: (v27 | (v27 / ((0) | 1)))}
                                                                                                            _ = v65
                                                                                                        case -9223372036854775808:
                                                                                                            v66 := GenS15{F0: (fi16 / ((-9223372036854775808) | 1))}
                                                                                                            _ = v66
                                                                                                        case -1:
                                                                                                            v27 = v27
                                                                                                        }
                                                                                                    default:
                                                                                                        v67 := []int{(fi16 % ((fi16) | 1)), (fi17 & 128)}
                                                                                                        _ = v67
                                                                                                    }
                                                                                                case -9223372036854775808:
                                                                                                    switch v52 {
                                                                                                    case "":
                                                                                                        var v68 int64 = v4
                                                                                                        _ = v68
                                                                                                    case "a":
                                                                                                        fmt.Println(v4)
                                                                                                    case "\n\t":
                                                                                                        fmt.Println(v4)
                                                                                                        var v69 float32 = (0.0 + float32(1e10))
                                                                                                        v70 := []int{fi16, (fi16 % (((fi17 | (fi16 * fi16))) | 1))}
                                                                                                        switch (^(fi17)) {
                                                                                                        case -9223372036854775808:
                                                                                                            v71 := GenS4{F0: 157, F1: fi17, F2: 0}
                                                                                                            _ = v71
                                                                                                        case 2147483647:
                                                                                                            switch v52 {
                                                                                                            case "hello":
                                                                                                                v72 := GenS13{F0: v52, F1: 0}
                                                                                                                _ = v72
                                                                                                            default:
                                                                                                                var v73 int64 = (v40 | v40)
                                                                                                                _ = v73
                                                                                                            }
                                                                                                        case 316:
                                                                                                            fmt.Println(v69, v4, v40)
                                                                                                        }
                                                                                                        switch fi16 {
                                                                                                        case 4294967295:
                                                                                                            v74 := GenS16{F0: -2147483648}
                                                                                                            _ = v74
                                                                                                        case -9223372036854775808:
                                                                                                            switch (v52 + v52) {
                                                                                                            case "\n\t":
                                                                                                                if ((true && false) && true) {
                                                                                                                    if len(v70) >= 2 {
                                                                                                                        v70 = v70[1:len(v70)]
                                                                                                                    }
                                                                                                                }
                                                                                                            case "foo bar":
                                                                                                                var v75 uint32 = v41
                                                                                                                var v76 uint32 = 582
                                                                                                                v76 -= v75
                                                                                                                _ = v75
                                                                                                                _ = v76
                                                                                                            case "hello":
                                                                                                                var v77 uint64 = 1
                                                                                                                _ = v77
                                                                                                            }
                                                                                                            var v78 float32 = (1e10 * (v69 + v69))
                                                                                                            if (false || true) {
                                                                                                                fmt.Println(v40, v52)
                                                                                                                var v79 int16 = -699
                                                                                                                _ = v79
                                                                                                            }
                                                                                                            _ = v78
                                                                                                        default:
                                                                                                            var v80 int16 = 99
                                                                                                            _ = v80
                                                                                                        }
                                                                                                        _ = v69
                                                                                                        _ = v70
                                                                                                    default:
                                                                                                        v41 -= v41
                                                                                                        var v81 uint = uint(v41)
                                                                                                        fmt.Println(v40, v5, v52)
                                                                                                        var v82 uint8 = 197
                                                                                                        _ = v81
                                                                                                        _ = v82
                                                                                                    }
                                                                                                default:
                                                                                                    v40 += (v4 | v40)
                                                                                                }
                                                                                            case "hello":
                                                                                                delete(v2, (v52 + v52))
                                                                                            }
                                                                                            v52 = v52
                                                                                            delete(v3, (("a" + (((v52 + v52) + v52) + v52)) + v52))
                                                                                            switch (^(fi17)) {
                                                                                            case 2147483647:
                                                                                                v83 := GenS6{F0: 255}
                                                                                                var v84 uint = v1
                                                                                                switch (fi17 & fi17) {
                                                                                                case 32767:
                                                                                                    if true {
                                                                                                        v1 += v1
                                                                                                        v1 += 65535
                                                                                                        fmt.Println(fi17, v27, fi17)
                                                                                                    }
                                                                                                case 233:
                                                                                                    if true {
                                                                                                        var v85 uint32 = (v41 << ((v1) & 63))
                                                                                                        if ((v84 | (^(v1))) != (v1 & v1)) {
                                                                                                            var v86 uint16 = v27
                                                                                                            v85 += (v41 % ((v85) | 1))
                                                                                                            _ = v86
                                                                                                        }
                                                                                                        var v87 uint16 = v27
                                                                                                        v84 = v1
                                                                                                        _ = v85
                                                                                                        _ = v87
                                                                                                    } else {
                                                                                                        var v88 uint8 = 1
                                                                                                        _ = v88
                                                                                                    }
                                                                                                case 2147483647:
                                                                                                    var v89 int64 = (v4 | (v40 ^ v4))
                                                                                                    _ = v89
                                                                                                default:
                                                                                                    v1 = v5
                                                                                                }
                                                                                                _ = v83
                                                                                                _ = v84
                                                                                            case 140:
                                                                                                var v90 int8 = 0
                                                                                                _ = v90
                                                                                            default:
                                                                                                v91 := GenS1{F0: v4, F1: 4294967295}
                                                                                                var v92 bool = true
                                                                                                var v93 float32 = 1.5
                                                                                                _ = v91
                                                                                                _ = v92
                                                                                                _ = v93
                                                                                            }
                                                                                            v94 := GenS3{F0: (v27 | v27), F1: v27, F2: v41}
                                                                                            v4 -= (v40 << ((v5) & 63))
                                                                                            _ = v53
                                                                                            _ = v54
                                                                                            _ = v94
                                                                                        } else {
                                                                                            fmt.Println(v1, v5)
                                                                                        }
                                                                                    } else {
                                                                                        var v95 uint = 65535
                                                                                        _ = v95
                                                                                    }
                                                                                    var v96 uint16 = (v27 | 0)
                                                                                    _ = v52
                                                                                    _ = v96
                                                                                }
                                                                            } else {
                                                                                switch (fi16 << (((v5 << ((255) & 63))) & 63)) {
                                                                                case 1:
                                                                                    var v97 bool = false
                                                                                    _ = v97
                                                                                case 9223372036854775807:
                                                                                    var v98 bool = true
                                                                                    _ = v98
                                                                                default:
                                                                                    v40 -= (v4 + int64(v5))
                                                                                }
                                                                                var v99 bool = (3.14159 == ((1.5 + -2.25) * (3.14159 + 1e10)))
                                                                                _ = v99
                                                                            }
                                                                            _ = v40
                                                                            _ = v41
                                                                        }
                                                                        v5 = v1
                                                                        var v100 int = (-(fi16))
                                                                        _ = v100
                                                                    }
                                                                    var v101 int16 = 32767
                                                                    _ = v101
                                                                }
                                                                v102 := []int{-569, (fi16 ^ (fi17 / ((fi17) | 1)))}
                                                                _ = v102
                                                            }
                                                        }
                                                    }
                                                case 263:
                                                    var v103 uint = (^(v5))
                                                    _ = v103
                                                }
                                                if (0 == 32767) {
                                                    v27 -= v27
                                                    if false {
                                                        v5 -= 4294967295
                                                    } else {
                                                        v104 := []int{int(v4), (fi17 % ((fi16) | 1)), (fi17 << ((30) & 63))}
                                                        if ((v5 & 175) >= 2) {
                                                            if (v27 <= (v27 << (((v1 % ((18446744073709551615) | 1))) & 63))) {
                                                                if (true || true) {
                                                                    if true {
                                                                        var v105 string = "foo bar"
                                                                        v105 = v105
                                                                        _ = v105
                                                                    } else {
                                                                        _ = len(v104)
                                                                        var v106 int8 = 100
                                                                        _ = v106
                                                                    }
                                                                    v27 = (v27 / ((uint16(v5)) | 1))
                                                                    fmt.Println(v27, fi17)
                                                                } else {
                                                                    fmt.Println(fi16, v27, fi17)
                                                                    v5 -= v5
                                                                    if len(v104) >= 2 {
                                                                        v104 = v104[1:len(v104)]
                                                                    }
                                                                }
                                                                switch (fi17 * fi17) {
                                                                case 0:
                                                                    v1 = (v1 % ((1) | 1))
                                                                case 2147483647:
                                                                    var v107 uint = (^(v5))
                                                                    _ = v107
                                                                case 636:
                                                                    v108 := GenS13{F0: ("\n\t" + ("" + "a")), F1: 2147483647}
                                                                    _ = v108
                                                                }
                                                                if (0.0 != 0.1) {
                                                                    fmt.Println(v4)
                                                                }
                                                            }
                                                        }
                                                        _ = v104
                                                    }
                                                } else {
                                                    var v109 uint32 = 1
                                                    v4 -= (v4 * v4)
                                                    var v110 float64 = 1.5
                                                    var v111 int8 = 42
                                                    var v112 uint = v5
                                                    _ = v109
                                                    _ = v110
                                                    _ = v111
                                                    _ = v112
                                                }
                                                fmt.Println(v27, v5, fi16)
                                                _ = v27
                                            }
                                        case 32767:
                                            if true {
                                                fmt.Println(fi16)
                                            }
                                        }
                                        v4 += v4
                                    }
                                }
                            }
                            if true {
                                var v113 uint16 = 2
                                v1 += v5
                                fmt.Println(fi17, v4, v113)
                                _ = v113
                            }
                            _ = v22
                        } else {
                            v114 := GenS17{F0: 4294967295, F1: 32767, F2: ("hello" + "foo bar")}
                            var v115 uint32 = 683
                            if false {
                                if (-317 <= -1) {
                                    v116 := GenS7{F0: 0, F1: (v115 << ((v1) & 63)), F2: 183}
                                    v117 := []int{fi16, fi16, int(v1), (-(fi16))}
                                    _ = v116
                                    _ = v117
                                } else {
                                    var v118 int16 = 42
                                    v119 := GenS3{F0: 0, F1: 42, F2: (v115 ^ 69)}
                                    _ = v118
                                    _ = v119
                                }
                                switch (fi17 ^ 128) {
                                case 9223372036854775807:
                                    var v120 string = "a"
                                    _ = v120
                                }
                            }
                            v121 := v2[""]
                            _ = v114
                            _ = v115
                            _ = v121
                        }
                        fmt.Println(fi16, fi17, v4)
                        v2["hello"] = fi16
                    }
                    v122 := GenS18{F0: 832, F1: 42}
                    v4 += (v4 * 9223372036854775807)
                    _ = v122
                }
                fmt.Println(v1, fi17)
                _ = fi17
            }
            _ = fi16
        }
        fmt.Println(v5, v5, v4)
    } else {
        delete(v3, "hello")
    }
    var v123 string = ("\n\t" + "\n\t")
    var v124 uint16 = 65535
    v125 := GenS19{F0: (v124 / (((^(v124))) | 1)), F1: 1000}
    var v126 uint16 = v124
    var v127 float64 = float64(-2.25)
    if true {
        var v128 int = -9223372036854775808
        if (65535 >= 255) {
            v129 := GenS20{F0: (18446744073709551615 < 78), F1: 80}
            v128 = v128
            _ = v129
        } else {
            v130 := GenS11{F0: v124, F1: 42, F2: 1}
            var v131 uint8 = 42
            _ = v130
            _ = v131
        }
        _ = v128
    }
    for fi132 := 0; fi132 < 20; fi132++ {
        delete(v3, "\n\t")
        var v133 uint8 = 255
        for fi134 := 0; fi134 < 2; fi134++ {
            var v135 uint32 = 689
            v133 += v133
            if (576 >= (fi134 + fi134)) {
                if ((v133 >> ((v5) & 63)) > (^(v133))) {
                    if false {
                        fmt.Println(v124)
                    }
                } else {
                    v4 -= 0
                }
            }
            _ = fi134
            _ = v135
        }
        _ = fi132
        _ = v133
    }
    v136 := GenS2{F0: -32768, F1: -783, F2: 2147483648}
    fmt.Println("v1:", v1)
    fmt.Println("v123:", v123)
    fmt.Println("v124:", v124)
    fmt.Println("v126:", v126)
    fmt.Println("v127:", v127)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v123
    _ = v124
    _ = v125
    _ = v126
    _ = v127
    _ = v136
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
