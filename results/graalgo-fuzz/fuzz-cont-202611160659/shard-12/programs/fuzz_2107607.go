package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 string; F1 int8; F2 int8 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 int; F1 uint8; F2 uint16 }
type GenS10 struct { F0 int; F1 int }
type GenS11 struct { F0 string; F1 string }
type GenS12 struct { F0 uint64 }
type GenS13 struct { F0 int64; F1 uint8 }
type GenS14 struct { F0 int16; F1 uint64; F2 uint }
type GenS15 struct { F0 uint32 }
type GenS16 struct { F0 uint16; F1 uint32; F2 uint64 }
type GenS17 struct { F0 int64; F1 bool }
type GenS18 struct { F0 uint16 }
type GenS19 struct { F0 uint16; F1 int16 }
type GenS20 struct { F0 string }
func genFn0() int32 {
    for fi7 := 0; fi7 < 9; fi7++ {
        var v8 uint = 65535
        v8 += (v8 - v8)
        if false {
            fmt.Println(v8, v8)
        }
        _ = fi7
        _ = v8
    }
    var v9 uint64 = 4294967295
    v9 = 62
    if (true || false) {
        var v10 int = 127
        v10 = v10
        _ = v10
    }
    _ = v9
    return -1
}

func main() {
    var v1 uint8 = 89
    var v2 int32 = 1
    v3 := []string{("" + "")}
    var v4 string = "\n\t"
    var v5 uint = 273
    var v6 float32 = (3.14159 + float32(3.14159))
    switch v4 {
    case "a":
        v11 := GenS1{F0: 4294967295, F1: 676}
        _ = v11
    case "hello":
        for fi12 := 0; fi12 < 1; fi12++ {
            v3 = append(v3, (v4 + (v4 + v4)))
            for fi13 := 0; fi13 < 12; fi13++ {
                var v14 float32 = v6
                _ = fi13
                _ = v14
            }
            var v15 string = v4
            _ = fi12
            _ = v15
        }
    default:
        v1 = (v1 | (^(v1)))
    }
    var v16 uint8 = v1
    for fi17 := 0; fi17 < 7; fi17++ {
        for fi18 := 0; fi18 < 20; fi18++ {
            if len(v3) > 0 {
                v3[0] = v4
            }
            var v19 int16 = 1
            if (1 >= 0) {
                if (fi17 >= (fi17 / (((fi17 & fi17)) | 1))) {
                    fmt.Println(v19, v16, fi17)
                    v20 := GenS2{F0: 42}
                    switch fi18 {
                    case 2147483647:
                        var v21 int64 = 9223372036854775807
                        _ = v21
                    case 42:
                        switch (^(fi17)) {
                        case 0:
                            v22 := GenS2{F0: int64(v19)}
                            _ = v22
                        default:
                            var v23 int8 = 110
                            _ = v23
                        }
                    }
                    _ = v20
                } else {
                    fmt.Println(v19)
                    v1 = (v1 | v1)
                    switch fi18 {
                    case 9223372036854775807:
                        var v24 float64 = (((1e10 + -2.25) * 1.5) / (1.5 / 1e10))
                        _ = v24
                    case 0:
                        if (fi18 < fi17) {
                            v25 := GenS1{F0: 328, F1: 0}
                            v6 -= ((0.0 + -2.25) + -2.25)
                            _ = v25
                        } else {
                            var v26 uint64 = 39
                            var v27 int64 = 1
                            var v28 uint16 = 937
                            _ = v26
                            _ = v27
                            _ = v28
                        }
                    }
                    v1 -= (^(v16))
                    switch (v4 + v4) {
                    case "foo bar":
                        var v29 bool = (false || false)
                        v30 := GenS3{F0: 853}
                        var v31 bool = (1 == 0)
                        _ = v29
                        _ = v30
                        _ = v31
                    }
                }
            }
            _ = fi18
            _ = v19
        }
        _ = fi17
    }
    v32 := GenS4{F0: -9223372036854775808}
    if (v16 < v16) {
        for fi33 := 0; fi33 < 14; fi33++ {
            v34 := GenS4{F0: -1}
            v2 = v2
            if true {
                var v35 int16 = 925
                _ = v35
            }
            _ = fi33
            _ = v34
        }
    }
    if (1 != -546) {
        for fi36 := 0; fi36 < 14; fi36++ {
            v6 += v6
            for fi37 := 0; fi37 < 6; fi37++ {
                fmt.Println(v6, v6)
                switch v4 {
                case "hello":
                    v6 = v6
                default:
                    v38 := GenS5{F0: (("\n\t" + v4) + v4), F1: 127, F2: 56}
                    _ = v38
                }
                switch v4 {
                case "foo bar":
                    fmt.Println(fi37, v6)
                case "\n\t":
                    v39 := GenS1{F0: 657, F1: 888}
                    _ = v39
                }
                _ = fi37
            }
            _ = fi36
        }
    } else {
        fmt.Println(v16)
        v40 := GenS5{F0: (v4 + (v4 + v4)), F1: 100, F2: -128}
        switch "hello" {
        case "hello":
            var v41 uint16 = 0
            _ = v41
        case "":
            if (-1 <= 0) {
                var v42 int64 = 42
                _ = v42
            } else {
                v43 := GenS5{F0: ("foo bar" + v4), F1: -42, F2: -1}
                if (0 <= 32767) {
                    v44 := GenS5{F0: v4, F1: 3, F2: 13}
                    var v45 uint32 = 42
                    _ = v44
                    _ = v45
                } else {
                    switch -1 {
                    case 32767:
                        v46 := GenS5{F0: v4, F1: 76, F2: 1}
                        _ = v46
                    default:
                        if (116 < -70) {
                            for fi47 := 0; fi47 < 8; fi47++ {
                                if true {
                                    for fi48 := 0; fi48 < 12; fi48++ {
                                        var v49 float64 = 0.1
                                        _ = fi48
                                        _ = v49
                                    }
                                    _ = len(v3)
                                    for fi50 := 0; fi50 < 2; fi50++ {
                                        var v51 float64 = float64((3.14159 / (v6 * v6)))
                                        if len(v3) > 0 {
                                            v3[0] = v4
                                        }
                                        if ((v4 + "\n\t") == "a") {
                                            if (!true) {
                                                switch "foo bar" {
                                                case "\n\t":
                                                    if true {
                                                        if (!false) {
                                                            fmt.Println(v6)
                                                        } else {
                                                            v51 = v51
                                                        }
                                                        if (!true) {
                                                            var v52 int16 = 42
                                                            _ = v52
                                                        } else {
                                                            fmt.Println(v2, v16)
                                                            switch (fi50 >> ((v5) & 63)) {
                                                            case 0:
                                                                if true {
                                                                    fmt.Println(v4)
                                                                }
                                                            case -1:
                                                                switch v4 {
                                                                case "":
                                                                    if (!false) {
                                                                        switch fi50 {
                                                                        case -85:
                                                                            fmt.Println(v2, v2, v5)
                                                                        case 892:
                                                                            if (459 <= 1) {
                                                                                v3 = append(v3, "")
                                                                            } else {
                                                                                v53 := GenS5{F0: v4, F1: 127, F2: 1}
                                                                                _ = v53
                                                                            }
                                                                        case 128:
                                                                            fmt.Println(v6, v6, fi47)
                                                                        default:
                                                                            if true {
                                                                                if false {
                                                                                    v1 = (v16 << (((v5 - v5)) & 63))
                                                                                } else {
                                                                                    var v54 string = "foo bar"
                                                                                    fmt.Println(v54, v54)
                                                                                    switch (v4 + (("a" + (v54 + v4)) + v4)) {
                                                                                    case "":
                                                                                        var v55 int = (fi50 * (fi50 ^ fi47))
                                                                                        _ = v55
                                                                                    case "hello":
                                                                                        if (fi50 == fi47) {
                                                                                            if (-1 > -85) {
                                                                                                v4 = v4
                                                                                                v5 = (v5 * (v5 ^ 0))
                                                                                            } else {
                                                                                                if (true || (-32768 > 526)) {
                                                                                                    if len(v3) >= 2 {
                                                                                                        v3 = v3[1:len(v3)]
                                                                                                    }
                                                                                                    if false {
                                                                                                        switch "hello" {
                                                                                                        case "hello":
                                                                                                            fmt.Println(v54, v51)
                                                                                                        case "a":
                                                                                                            if len(v3) >= 2 {
                                                                                                                v3 = v3[1:len(v3)]
                                                                                                            }
                                                                                                        case "\n\t":
                                                                                                            v1 += v1
                                                                                                        }
                                                                                                        fmt.Println(v6)
                                                                                                        if (1 == 674) {
                                                                                                            if true {
                                                                                                                if (-171 < 0) {
                                                                                                                    if len(v3) > 0 {
                                                                                                                        v3[0] = (("foo bar" + "") + "")
                                                                                                                    }
                                                                                                                    _ = len(v3)
                                                                                                                }
                                                                                                                if (true && false) {
                                                                                                                    var v56 float64 = v51
                                                                                                                    switch fi50 {
                                                                                                                    case -9223372036854775808:
                                                                                                                        var v57 int16 = -32768
                                                                                                                        _ = v57
                                                                                                                    case 9223372036854775807:
                                                                                                                        if true {
                                                                                                                            switch 2147483647 {
                                                                                                                            case -796:
                                                                                                                                v2 -= (^(v2))
                                                                                                                            case 4294967295:
                                                                                                                                if (!true) {
                                                                                                                                    v58 := GenS3{F0: 65535}
                                                                                                                                    if (!(!(v51 > v56))) {
                                                                                                                                        v56 = 1e10
                                                                                                                                        if false {
                                                                                                                                            v59 := GenS1{F0: 1, F1: 281}
                                                                                                                                            var v60 int16 = -916
                                                                                                                                            var v61 int64 = 0
                                                                                                                                            _ = v59
                                                                                                                                            _ = v60
                                                                                                                                            _ = v61
                                                                                                                                        } else {
                                                                                                                                            switch "\n\t" {
                                                                                                                                            case "\n\t":
                                                                                                                                                var v62 uint64 = 4294967295
                                                                                                                                                _ = v62
                                                                                                                                            case "a":
                                                                                                                                                var v63 uint8 = v1
                                                                                                                                                _ = v63
                                                                                                                                            case "foo bar":
                                                                                                                                                switch v4 {
                                                                                                                                                case "\n\t":
                                                                                                                                                    var v64 int8 = 42
                                                                                                                                                    _ = v64
                                                                                                                                                }
                                                                                                                                            }
                                                                                                                                            v2 -= (-(v2))
                                                                                                                                        }
                                                                                                                                    }
                                                                                                                                    var v65 int8 = -1
                                                                                                                                    var v66 int32 = v2
                                                                                                                                    _ = v58
                                                                                                                                    _ = v65
                                                                                                                                    _ = v66
                                                                                                                                } else {
                                                                                                                                    var v67 uint16 = 42
                                                                                                                                    var v68 int = (fi50 ^ fi47)
                                                                                                                                    v3 = append(v3, "foo bar")
                                                                                                                                    _ = v67
                                                                                                                                    _ = v68
                                                                                                                                }
                                                                                                                            case -1:
                                                                                                                                if len(v3) > 0 {
                                                                                                                                    v3[0] = v4
                                                                                                                                }
                                                                                                                                v5 = 2
                                                                                                                            default:
                                                                                                                                switch "hello" {
                                                                                                                                case "\n\t":
                                                                                                                                    v69 := GenS6{F0: 255}
                                                                                                                                    _ = v69
                                                                                                                                case "":
                                                                                                                                    if len(v3) > 0 {
                                                                                                                                        v3[0] = "foo bar"
                                                                                                                                    }
                                                                                                                                case "foo bar":
                                                                                                                                    v2 += (v2 - (v2 ^ -1))
                                                                                                                                default:
                                                                                                                                    switch (v54 + v4) {
                                                                                                                                    case "":
                                                                                                                                        v54 = (("" + v54) + "")
                                                                                                                                    case "hello":
                                                                                                                                        var v70 int8 = 1
                                                                                                                                        _ = v70
                                                                                                                                    }
                                                                                                                                }
                                                                                                                            }
                                                                                                                            switch 866 {
                                                                                                                            case 32767:
                                                                                                                                if len(v3) > 0 {
                                                                                                                                    v3[0] = "hello"
                                                                                                                                }
                                                                                                                            case 127:
                                                                                                                                switch "hello" {
                                                                                                                                case "foo bar":
                                                                                                                                    var v71 bool = false
                                                                                                                                    _ = v71
                                                                                                                                }
                                                                                                                            default:
                                                                                                                                v54 = "a"
                                                                                                                            }
                                                                                                                        } else {
                                                                                                                            _ = len(v3)
                                                                                                                            if len(v3) > 0 {
                                                                                                                                v3[0] = ((v54 + ("foo bar" + v54)) + (v54 + v54))
                                                                                                                            }
                                                                                                                        }
                                                                                                                    }
                                                                                                                    v2 += v2
                                                                                                                    _ = v56
                                                                                                                } else {
                                                                                                                    v5 -= (v5 ^ (v5 >> (((^(v5))) & 63)))
                                                                                                                    v5 = 1
                                                                                                                    v2 += (v2 / (((-(v2))) | 1))
                                                                                                                }
                                                                                                                v2 = v2
                                                                                                            } else {
                                                                                                                var v72 uint32 = 1
                                                                                                                _ = v72
                                                                                                            }
                                                                                                            var v73 float32 = v6
                                                                                                            v73 += v6
                                                                                                            v74 := GenS1{F0: 397, F1: 1}
                                                                                                            v75 := GenS7{F0: 0}
                                                                                                            switch (fi47 % ((42) | 1)) {
                                                                                                            case -9223372036854775808:
                                                                                                                var v76 float32 = v73
                                                                                                                _ = v76
                                                                                                            default:
                                                                                                                v77 := GenS8{F0: 100}
                                                                                                                _ = v77
                                                                                                            }
                                                                                                            var v78 int16 = 42
                                                                                                            _ = v73
                                                                                                            _ = v74
                                                                                                            _ = v75
                                                                                                            _ = v78
                                                                                                        } else {
                                                                                                            if (47 > 15) {
                                                                                                                var v79 bool = (false || (!(!false)))
                                                                                                                _ = v79
                                                                                                            } else {
                                                                                                                v2 += (v2 + int32(v1))
                                                                                                                v4 = ""
                                                                                                                _ = len(v3)
                                                                                                            }
                                                                                                            v80 := GenS9{F0: 0, F1: (^(v1)), F2: 0}
                                                                                                            v81 := GenS10{F0: fi47, F1: (fi47 << (((v5 + 65535)) & 63))}
                                                                                                            if len(v3) > 0 {
                                                                                                                v3[0] = ""
                                                                                                            }
                                                                                                            _ = v80
                                                                                                            _ = v81
                                                                                                        }
                                                                                                        var v82 int32 = genFn0()
                                                                                                        switch (v4 + ("\n\t" + v4)) {
                                                                                                        case "foo bar":
                                                                                                            if true {
                                                                                                                v82 += genFn0()
                                                                                                                _ = len(v3)
                                                                                                            } else {
                                                                                                                if false {
                                                                                                                    v3 = append(v3, "")
                                                                                                                    v3 = append(v3, v54)
                                                                                                                }
                                                                                                            }
                                                                                                        }
                                                                                                        _ = v82
                                                                                                    } else {
                                                                                                        switch ((v54 + v4) + "hello") {
                                                                                                        case "":
                                                                                                            if (((((!false) && true) || true) || (2 == 971)) || false) {
                                                                                                                var v83 float32 = v6
                                                                                                                _ = v83
                                                                                                            } else {
                                                                                                                if len(v3) > 0 {
                                                                                                                    v3[0] = v54
                                                                                                                }
                                                                                                            }
                                                                                                        case "hello":
                                                                                                            var v84 uint16 = uint16(v16)
                                                                                                            _ = v84
                                                                                                        default:
                                                                                                            if (true || false) {
                                                                                                                if true {
                                                                                                                    v54 = "foo bar"
                                                                                                                    v85 := GenS7{F0: 64}
                                                                                                                    var v86 int16 = 1000
                                                                                                                    v87 := GenS9{F0: (fi50 | fi47), F1: v1, F2: 42}
                                                                                                                    v1 -= 255
                                                                                                                    _ = v85
                                                                                                                    _ = v86
                                                                                                                    _ = v87
                                                                                                                } else {
                                                                                                                    v51 -= ((1.5 * v51) * 1e10)
                                                                                                                }
                                                                                                                fmt.Println(v2)
                                                                                                                switch fi47 {
                                                                                                                case -241:
                                                                                                                    var v88 int16 = 310
                                                                                                                    _ = v88
                                                                                                                }
                                                                                                                if false {
                                                                                                                    v89 := GenS1{F0: 2, F1: 0}
                                                                                                                    switch v54 {
                                                                                                                    case "a":
                                                                                                                        v90 := GenS6{F0: 2}
                                                                                                                        v3 = append(v3, (v54 + ("\n\t" + "a")))
                                                                                                                        v6 = 1.5
                                                                                                                        v16 = 255
                                                                                                                        _ = v90
                                                                                                                    default:
                                                                                                                        var v91 bool = (100 != -128)
                                                                                                                        _ = v91
                                                                                                                    }
                                                                                                                    var v92 uint16 = 32767
                                                                                                                    _ = v89
                                                                                                                    _ = v92
                                                                                                                } else {
                                                                                                                    v93 := GenS11{F0: v54, F1: "\n\t"}
                                                                                                                    _ = v93
                                                                                                                }
                                                                                                            }
                                                                                                        }
                                                                                                        if true {
                                                                                                            v54 = v54
                                                                                                            var v94 uint16 = 301
                                                                                                            fmt.Println(v5, v4, v5)
                                                                                                            _ = v94
                                                                                                        }
                                                                                                    }
                                                                                                    var v95 uint = 0
                                                                                                    _ = v95
                                                                                                } else {
                                                                                                    var v96 float64 = v51
                                                                                                    v16 = 161
                                                                                                    _ = v96
                                                                                                }
                                                                                                if len(v3) >= 2 {
                                                                                                    v3 = v3[1:len(v3)]
                                                                                                }
                                                                                            }
                                                                                            fmt.Println(v16, v54, v1)
                                                                                            if (!(!(v51 <= (1e10 - float64(v6))))) {
                                                                                                _ = len(v3)
                                                                                                v16 += 255
                                                                                            } else {
                                                                                                v97 := GenS2{F0: 9223372036854775807}
                                                                                                fmt.Println(v6, v4)
                                                                                                _ = v97
                                                                                            }
                                                                                        } else {
                                                                                            var v98 int8 = 42
                                                                                            _ = v98
                                                                                        }
                                                                                    default:
                                                                                        if len(v3) >= 2 {
                                                                                            v3 = v3[1:len(v3)]
                                                                                        }
                                                                                    }
                                                                                    v51 += (0.0 + (float64(v6) * v51))
                                                                                    _ = v54
                                                                                }
                                                                                v99 := GenS8{F0: 42}
                                                                                _ = v99
                                                                            } else {
                                                                                v2 = v2
                                                                            }
                                                                        }
                                                                        if true {
                                                                            switch v4 {
                                                                            case "\n\t":
                                                                                if len(v3) >= 2 {
                                                                                    v3 = v3[1:len(v3)]
                                                                                }
                                                                            default:
                                                                                v100 := GenS12{F0: 840}
                                                                                _ = v100
                                                                            }
                                                                            v5 += (v5 % ((951) | 1))
                                                                            var v101 int = (^(fi47))
                                                                            switch ("\n\t" + v4) {
                                                                            case "":
                                                                                var v102 uint = 65535
                                                                                _ = v102
                                                                            case "\n\t":
                                                                                _ = len(v3)
                                                                                v103 := GenS10{F0: 999, F1: (fi50 << ((v5) & 63))}
                                                                                _ = v103
                                                                            case "hello":
                                                                                switch -1 {
                                                                                case 909:
                                                                                    switch v4 {
                                                                                    case "\n\t":
                                                                                        var v104 string = v4
                                                                                        fmt.Println(fi50, v2, fi50)
                                                                                        v16 = v16
                                                                                        var v105 int = -9223372036854775808
                                                                                        switch "foo bar" {
                                                                                        case "\n\t":
                                                                                            _ = len(v3)
                                                                                        default:
                                                                                            if false {
                                                                                                var v106 bool = false
                                                                                                v6 = (v6 * float32(1.5))
                                                                                                var v107 int16 = 1000
                                                                                                _ = v106
                                                                                                _ = v107
                                                                                            }
                                                                                        }
                                                                                        _ = v104
                                                                                        _ = v105
                                                                                    case "":
                                                                                        if (true && (!(genFn0() != v2))) {
                                                                                            var v108 int64 = 1
                                                                                            if (true && (42 > 36)) {
                                                                                                _ = len(v3)
                                                                                            }
                                                                                            if ((v5 > v5) || (!false)) {
                                                                                                var v109 uint8 = 127
                                                                                                _ = v109
                                                                                            } else {
                                                                                                var v110 int64 = 890
                                                                                                v111 := GenS11{F0: v4, F1: ((v4 + v4) + v4)}
                                                                                                switch int(v110) {
                                                                                                case 42:
                                                                                                    if len(v3) > 0 {
                                                                                                        v3[0] = "a"
                                                                                                    }
                                                                                                case -644:
                                                                                                    v6 -= 3.14159
                                                                                                default:
                                                                                                    if len(v3) > 0 {
                                                                                                        v3[0] = (v4 + v4)
                                                                                                    }
                                                                                                }
                                                                                                _ = v110
                                                                                                _ = v111
                                                                                            }
                                                                                            _ = v108
                                                                                        }
                                                                                    default:
                                                                                        var v112 uint8 = (v16 | (v16 - (v1 % ((v1) | 1))))
                                                                                        _ = v112
                                                                                    }
                                                                                    if (!(2147483647 != 42)) {
                                                                                        fmt.Println(v6, v1, v1)
                                                                                        v113 := GenS12{F0: 18446744073709551615}
                                                                                        _ = v113
                                                                                    } else {
                                                                                        v6 = (float32(v51) / -2.25)
                                                                                        v114 := GenS11{F0: "foo bar", F1: v4}
                                                                                        var v115 int32 = (v2 / (((v2 ^ (v2 | (^(v2))))) | 1))
                                                                                        if (fi47 != (-(fi50))) {
                                                                                            var v116 int = 246
                                                                                            _ = v116
                                                                                        }
                                                                                        v51 = (v51 + 0.0)
                                                                                        _ = v114
                                                                                        _ = v115
                                                                                    }
                                                                                case 4294967295:
                                                                                    v16 = 254
                                                                                case -930:
                                                                                    v101 = 9223372036854775807
                                                                                    var v117 int16 = 1
                                                                                    if (2147483648 > 1) {
                                                                                        fmt.Println(fi50)
                                                                                    } else {
                                                                                        if false {
                                                                                            fmt.Println(v117, v117, fi50)
                                                                                        }
                                                                                        switch (fi50 % (((-(fi47))) | 1)) {
                                                                                        case 208:
                                                                                            if ((fi50 + 2147483647) > v101) {
                                                                                                if len(v3) >= 2 {
                                                                                                    v3 = v3[1:len(v3)]
                                                                                                }
                                                                                                fmt.Println(v51, v1)
                                                                                            }
                                                                                        case 128:
                                                                                            var v118 uint8 = 1
                                                                                            _ = v118
                                                                                        default:
                                                                                            if len(v3) >= 2 {
                                                                                                v3 = v3[1:len(v3)]
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                    _ = v117
                                                                                default:
                                                                                    v119 := GenS13{F0: 9223372036854775807, F1: v1}
                                                                                    fmt.Println(v6, v51)
                                                                                    if len(v3) >= 2 {
                                                                                        v3 = v3[1:len(v3)]
                                                                                    }
                                                                                    _ = v119
                                                                                }
                                                                            }
                                                                            v6 += v6
                                                                            _ = v101
                                                                        }
                                                                        fmt.Println(v5, v5)
                                                                    } else {
                                                                        var v120 int16 = 1000
                                                                        v121 := GenS11{F0: v4, F1: "a"}
                                                                        _ = v120
                                                                        _ = v121
                                                                    }
                                                                    var v122 uint = 261
                                                                    v123 := GenS1{F0: 1, F1: 0}
                                                                    if len(v3) > 0 {
                                                                        v3[0] = (v4 + ((("\n\t" + "\n\t") + "hello") + "hello"))
                                                                    }
                                                                    var v124 uint = (^(v122))
                                                                    _ = v122
                                                                    _ = v123
                                                                    _ = v124
                                                                case "a":
                                                                    v5 = 2
                                                                case "\n\t":
                                                                    var v125 uint32 = 65535
                                                                    _ = v125
                                                                default:
                                                                    v3 = append(v3, (v4 + (v4 + v4)))
                                                                }
                                                            case 101:
                                                                v126 := GenS14{F0: 847, F1: 18446744073709551615, F2: v5}
                                                                _ = v126
                                                            }
                                                        }
                                                        v4 = v4
                                                    } else {
                                                        if len(v3) >= 2 {
                                                            v3 = v3[1:len(v3)]
                                                        }
                                                    }
                                                }
                                                var v127 float32 = ((((float32((v51 * ((v51 + 1.5) / v51))) * float32(v51)) - 0.1) + v6) - (v6 + 3.14159))
                                                _ = v127
                                            }
                                            v128 := GenS15{F0: 391}
                                            _ = v128
                                        } else {
                                            var v129 uint = v5
                                            _ = v129
                                        }
                                        var v130 int = (fi50 * fi47)
                                        _ = fi50
                                        _ = v130
                                        _ = v51
                                    }
                                } else {
                                    for fi131 := 0; fi131 < 19; fi131++ {
                                        v132 := GenS16{F0: 65535, F1: 4294967295, F2: 1}
                                        _ = fi131
                                        _ = v132
                                    }
                                }
                                switch v4 {
                                case "hello":
                                    _ = len(v3)
                                default:
                                    var v133 int = (fi47 ^ 128)
                                    _ = v133
                                }
                                var v134 float32 = 1.5
                                _ = fi47
                                _ = v134
                            }
                            var v135 string = ((v4 + v4) + "a")
                            if false {
                                for fi136 := 0; fi136 < 10; fi136++ {
                                    v6 = v6
                                    for fi137 := 0; fi137 < 10; fi137++ {
                                        fmt.Println(fi137, v2)
                                        fmt.Println(v5, v1, fi136)
                                        if len(v3) > 0 {
                                            v3[0] = v4
                                        }
                                        _ = fi137
                                    }
                                    _ = fi136
                                }
                            } else {
                                v138 := GenS4{F0: -711}
                                fmt.Println(v6, v16, v2)
                                if (true && (2 < 973)) {
                                    for fi139 := 0; fi139 < 1; fi139++ {
                                        if (!true) {
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                            var v140 int32 = -46
                                            for fi141 := 0; fi141 < 8; fi141++ {
                                                var v142 int64 = 459
                                                v3 = append(v3, "hello")
                                                _ = fi141
                                                _ = v142
                                            }
                                            _ = v140
                                        } else {
                                            _ = len(v3)
                                        }
                                        switch v4 {
                                        case "":
                                            _ = len(v3)
                                        case "foo bar":
                                            v4 = (("foo bar" + ((("foo bar" + "\n\t") + v135) + v4)) + (((v135 + "a") + v4) + "foo bar"))
                                        default:
                                            v143 := GenS1{F0: 0, F1: uint32(v1)}
                                            _ = v143
                                        }
                                        var v144 float32 = v6
                                        _ = fi139
                                        _ = v144
                                    }
                                }
                                _ = v138
                            }
                            _ = v135
                        } else {
                            var v145 int32 = (^(v2))
                            if true {
                                fmt.Println(v6, v145, v4)
                            } else {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            }
                            if true {
                                for fi146 := 0; fi146 < 5; fi146++ {
                                    v145 -= (v2 % ((genFn0()) | 1))
                                    _ = fi146
                                }
                            } else {
                                for fi147 := 0; fi147 < 15; fi147++ {
                                    if ((1 == -847) || false) {
                                        if true {
                                            v148 := GenS17{F0: 1, F1: (!false)}
                                            for fi149 := 0; fi149 < 10; fi149++ {
                                                v145 = v145
                                                fmt.Println(v1)
                                                v4 = v4
                                                _ = fi149
                                            }
                                            _ = len(v3)
                                            _ = v148
                                        } else {
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                            if ((v5 | 2) > 1) {
                                                v145 = v145
                                                v6 = (v6 + 0.1)
                                                v2 = v145
                                            }
                                            switch fi147 {
                                            case 0:
                                                if (1 != -32768) {
                                                    var v150 uint = v5
                                                    for fi151 := 0; fi151 < 3; fi151++ {
                                                        fmt.Println(v2, v6)
                                                        fmt.Println(fi151, v1, v1)
                                                        _ = len(v3)
                                                        _ = fi151
                                                    }
                                                    _ = v150
                                                } else {
                                                    switch (fi147 | fi147) {
                                                    case 128:
                                                        var v152 int32 = genFn0()
                                                        _ = v152
                                                    }
                                                    fmt.Println(v16, v2, v2)
                                                }
                                            default:
                                                var v153 int64 = -9223372036854775808
                                                _ = v153
                                            }
                                        }
                                        if false {
                                            switch v4 {
                                            case "foo bar":
                                                fmt.Println(v2)
                                            case "hello":
                                                for fi154 := 0; fi154 < 8; fi154++ {
                                                    v4 = v4
                                                    if (!false) {
                                                        fmt.Println(v145, v6, v4)
                                                        if true {
                                                            var v155 int = fi154
                                                            var v156 uint16 = 255
                                                            fmt.Println(fi147, fi154, v16)
                                                            _ = v155
                                                            _ = v156
                                                        }
                                                        if len(v3) >= 2 {
                                                            v3 = v3[1:len(v3)]
                                                        }
                                                        if (65535 < 32767) {
                                                            switch ("a" + v4) {
                                                            case "foo bar":
                                                                v145 += (v145 / ((v2) | 1))
                                                            }
                                                            v1 -= (v1 / ((v1) | 1))
                                                        }
                                                        v6 = 1e10
                                                    }
                                                    _ = fi154
                                                }
                                            case "a":
                                                _ = len(v3)
                                            }
                                        }
                                    }
                                    var v157 uint16 = 980
                                    _ = fi147
                                    _ = v157
                                }
                                switch v4 {
                                case "\n\t":
                                    if true {
                                        var v158 uint32 = 4294967295
                                        switch -943 {
                                        case 128:
                                            var v159 int32 = v2
                                            _ = v159
                                        case -1:
                                            var v160 int16 = int16(v16)
                                            _ = v160
                                        case 42:
                                            v161 := GenS10{F0: 42, F1: 1}
                                            _ = v161
                                        }
                                        _ = v158
                                    }
                                }
                                v6 += v6
                            }
                            _ = v145
                        }
                    }
                    if false {
                        var v162 float64 = (1e10 / 3.14159)
                        var v163 string = v4
                        for fi164 := 0; fi164 < 9; fi164++ {
                            var v165 bool = true
                            v162 = (1e10 * v162)
                            _ = fi164
                            _ = v165
                        }
                        _ = v162
                        _ = v163
                    }
                    if len(v3) > 0 {
                        v3[0] = "a"
                    }
                }
                _ = v43
            }
        case "a":
            v166 := GenS18{F0: 42}
            _ = v166
        default:
            fmt.Println(v2)
        }
        _ = v40
    }
    _ = len(v3)
    var v167 float64 = 1.5
    if false {
        v3 = append(v3, ("" + "a"))
        if ((-156 <= 0) || (!((true && (1 != 1)) && false))) {
            var v168 float32 = 1.5
            v5 = (^(v5))
            var v169 uint64 = 18446744073709551615
            _ = v168
            _ = v169
        } else {
            for fi170 := 0; fi170 < 13; fi170++ {
                for fi171 := 0; fi171 < 9; fi171++ {
                    if (false || false) {
                        if true {
                            switch fi171 {
                            case 9223372036854775807:
                                v4 = v4
                            case 32768:
                                v167 -= 3.14159
                            }
                            switch (fi171 + 127) {
                            case 981:
                                var v172 uint64 = 362
                                _ = v172
                            default:
                                if ((347 >= 821) && false) {
                                    v3 = append(v3, (v4 + v4))
                                    v16 -= v16
                                } else {
                                    var v173 uint8 = 2
                                    _ = v173
                                }
                            }
                            v1 = v16
                            if ((fi171 - (fi171 & fi170)) == fi170) {
                                var v174 int8 = -79
                                v2 = v2
                                _ = v174
                            } else {
                                var v175 float32 = 0.1
                                fmt.Println(v16, v4)
                                if len(v3) > 0 {
                                    v3[0] = v4
                                }
                                var v176 int32 = (v2 - 0)
                                _ = v175
                                _ = v176
                            }
                        } else {
                            if ((v16 / (((^(v1))) | 1)) == v16) {
                                v6 += v6
                                switch fi171 {
                                case -9223372036854775808:
                                    v2 -= (v2 * (v2 ^ v2))
                                case 2147483647:
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                case 128:
                                    v6 = 0.1
                                    if ((811 >= 2147483647) && (v167 > 3.14159)) {
                                        _ = len(v3)
                                        var v177 float64 = ((1e10 - (v167 - v167)) + float64((1.5 / float32(v167))))
                                        v177 -= v177
                                        _ = v177
                                    }
                                    v5 += v5
                                    if ((1 != 0) && ((v5 + v5) < (v5 ^ 18446744073709551615))) {
                                        v3 = append(v3, ("foo bar" + "foo bar"))
                                        v6 -= -2.25
                                        v178 := GenS5{F0: v4, F1: -128, F2: -42}
                                        _ = v178
                                    } else {
                                        var v179 float32 = 1.5
                                        fmt.Println(v179, v16, v1)
                                        _ = v179
                                    }
                                default:
                                    var v180 float64 = v167
                                    _ = v180
                                }
                            } else {
                                v3 = append(v3, v4)
                                var v181 string = v4
                                _ = v181
                            }
                        }
                        fmt.Println(v1, v6)
                        v167 += float64((((v6 * ((float32(v167) + v6) / v6)) + -2.25) * 3.14159))
                        var v182 int64 = 9223372036854775807
                        var v183 int32 = v2
                        if false {
                            if false {
                                if (!false) {
                                    switch v4 {
                                    case "\n\t":
                                        switch 205 {
                                        case 1:
                                            v5 += (v5 % (((^(v5))) | 1))
                                        case 9223372036854775807:
                                            var v184 string = "foo bar"
                                            switch (fi171 % (((fi171 + (fi170 ^ (fi171 ^ 4294967295)))) | 1)) {
                                            case 2147483647:
                                                if (false || (false || (255 != 959))) {
                                                    v184 = (("foo bar" + (("a" + v184) + v184)) + "\n\t")
                                                } else {
                                                    var v185 uint = v5
                                                    var v186 uint16 = 0
                                                    _ = v185
                                                    _ = v186
                                                }
                                            case -2147483648:
                                                v5 -= v5
                                            default:
                                                v183 = -2147483648
                                            }
                                            _ = len(v3)
                                            _ = v184
                                        }
                                    default:
                                        v6 -= v6
                                    }
                                    if ((!(false || true)) && (!false)) {
                                        if true {
                                            v4 = ""
                                            v3 = append(v3, ("foo bar" + (v4 + ("foo bar" + (((("\n\t" + (v4 + v4)) + "foo bar") + v4) + v4)))))
                                        } else {
                                            v187 := GenS19{F0: 42, F1: 1000}
                                            _ = v187
                                        }
                                    } else {
                                        _ = len(v3)
                                    }
                                    v1 -= 1
                                    var v188 uint16 = 949
                                    _ = v188
                                }
                            } else {
                                v5 = v5
                                var v189 uint8 = 162
                                _ = v189
                            }
                            var v190 uint64 = 1
                            switch -189 {
                            case -547:
                                if true {
                                    var v191 uint = (v5 - (^(v5)))
                                    _ = v191
                                }
                            case 128:
                                v192 := GenS1{F0: (v190 | v190), F1: 2147483648}
                                _ = v192
                            case 32767:
                                var v193 uint16 = 0
                                _ = v193
                            default:
                                v167 += v167
                            }
                            v194 := GenS11{F0: "", F1: v4}
                            _ = v190
                            _ = v194
                        } else {
                            var v195 int32 = (v183 >> ((v5) & 63))
                            _ = v195
                        }
                        v196 := GenS17{F0: (v182 + (v182 >> (((v5 & v5)) & 63))), F1: false}
                        _ = v182
                        _ = v183
                        _ = v196
                    }
                    _ = fi171
                }
                _ = fi170
            }
        }
    } else {
        if false {
            switch ("" + v4) {
            case "hello":
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            default:
                v4 = (v4 + ("\n\t" + ""))
            }
            v197 := GenS20{F0: v4}
            v4 = v4
            _ = v197
        }
        var v198 int8 = 100
        _ = v198
    }
    fmt.Println(v2)
    var v199 int32 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v167:", v167)
    fmt.Println("v199:", v199)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v167
    _ = v199
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
