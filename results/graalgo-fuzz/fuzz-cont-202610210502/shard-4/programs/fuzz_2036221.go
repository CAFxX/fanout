package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int8; F1 uint64; F2 uint32 }
type GenS4 struct { F0 uint64; F1 uint16; F2 bool }
type GenS5 struct { F0 int32; F1 uint32 }
type GenS6 struct { F0 string; F1 uint32 }
type GenS7 struct { F0 int32; F1 uint64 }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 int64; F1 bool }
type GenS10 struct { F0 uint8; F1 int64; F2 uint32 }
type GenS11 struct { F0 int32; F1 int8; F2 int64 }
type GenS12 struct { F0 uint }
type GenS13 struct { F0 int16 }
type GenS14 struct { F0 uint }

func main() {
    var v1 bool = true
    var v2 uint32 = 212
    v3 := []uint32{2, v2, 0, v2}
    for fi4 := 0; fi4 < 20; fi4++ {
        v5 := GenS1{F0: 1, F1: true}
        var v6 int64 = -141
        v6 += v6
        _ = fi4
        _ = v5
        _ = v6
    }
    var v7 int64 = -9223372036854775808
    for fi8 := 0; fi8 < 2; fi8++ {
        for fi9 := 0; fi9 < 17; fi9++ {
            fmt.Println(v7, v1, fi8)
            var v10 uint32 = 4294967295
            v1 = (1e10 != (-2.25 - 3.14159))
            _ = fi9
            _ = v10
        }
        v7 = v7
        var v11 float32 = 1e10
        _ = fi8
        _ = v11
    }
    switch (("a" + ("a" + "hello")) + ("a" + "")) {
    case "\n\t":
        v7 = v7
    case "":
        for fi12 := 0; fi12 < 2; fi12++ {
            v7 += -1
            var v13 int64 = 9223372036854775807
            _ = fi12
            _ = v13
        }
    case "foo bar":
        v14 := GenS1{F0: -2147483648, F1: (!((v1 && true) && (v1 && v1)))}
        _ = v14
    }
    v15 := GenS1{F0: -815, F1: (-900 == -1)}
    var v16 int16 = 547
    for fi17 := 0; fi17 < 14; fi17++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        fmt.Println(v2, fi17, v7)
        _ = fi17
    }
    if (81 <= 128) {
        v18 := GenS1{F0: int32(v2), F1: v1}
        if len(v3) > 0 {
            v3[0] = v2
        }
        v1 = false
        _ = v18
    } else {
        for fi19 := 0; fi19 < 8; fi19++ {
            for fi20 := 0; fi20 < 3; fi20++ {
                fmt.Println(v1)
                switch ("hello" + (("hello" + ((("\n\t" + "\n\t") + (("a" + ((("a" + ("a" + "a")) + "hello") + "\n\t")) + ("hello" + "a"))) + "")) + "foo bar")) {
                case "foo bar":
                    switch (fi19 << ((18446744073709551615) & 63)) {
                    case 992:
                        if len(v3) > 0 {
                            v3[0] = (v2 << ((18446744073709551615) & 63))
                        }
                    case 128:
                        v2 = v2
                    default:
                        if (!true) {
                            switch ("" + ((("foo bar" + (("hello" + "hello") + ("\n\t" + ""))) + ("foo bar" + "a")) + ("\n\t" + ""))) {
                            case "hello":
                                if (v7 > (v7 + (v7 >> ((18446744073709551615) & 63)))) {
                                    if ((2 == 0) && (1 <= 255)) {
                                        v1 = v1
                                        fmt.Println(v16, fi19, v7)
                                        switch (fi19 << ((2) & 63)) {
                                        case -166:
                                            var v21 bool = false
                                            _ = v21
                                        case 1:
                                            if false {
                                                if (1 == 32767) {
                                                    var v22 string = (("foo bar" + ((("" + ("\n\t" + "foo bar")) + "") + ("hello" + ""))) + ("a" + "a"))
                                                    switch 1 {
                                                    case 128:
                                                        var v23 uint64 = 584
                                                        _ = v23
                                                    case 32768:
                                                        v22 = (v22 + v22)
                                                    }
                                                    _ = v22
                                                }
                                            }
                                        default:
                                            switch "" {
                                            case "":
                                                fmt.Println(v7, fi19)
                                            }
                                        }
                                        fmt.Println(fi19, fi19)
                                    }
                                    var v24 int16 = v16
                                    v1 = (!v1)
                                    _ = v24
                                } else {
                                    var v25 uint64 = 835
                                    _ = v25
                                }
                            case "a":
                                v26 := GenS2{F0: 776}
                                var v27 int64 = (v7 | 1)
                                if (2 != 255) {
                                    switch (-(fi19)) {
                                    case 42:
                                        if (268 != 255) {
                                            var v28 bool = v1
                                            if (!true) {
                                                var v29 int16 = 32767
                                                var v30 uint8 = uint8(v2)
                                                fmt.Println(v29, v27)
                                                _ = v29
                                                _ = v30
                                            } else {
                                                v16 = 622
                                                v7 -= (^(v7))
                                            }
                                            switch "foo bar" {
                                            case "foo bar":
                                                var v31 int8 = -128
                                                _ = v31
                                            default:
                                                var v32 uint8 = 1
                                                _ = v32
                                            }
                                            _ = v28
                                        } else {
                                            v1 = (1 <= 0)
                                            _ = len(v3)
                                        }
                                    case 9223372036854775807:
                                        if ((v16 & v16) != v16) {
                                            v1 = true
                                            if ((!(v2 == (v2 << ((42) & 63)))) && true) {
                                                switch (fi19 / (((fi19 >> ((255) & 63))) | 1)) {
                                                case 1:
                                                    v33 := GenS3{F0: -8, F1: 42, F2: (v2 >> ((0) & 63))}
                                                    var v34 int64 = v7
                                                    if v1 {
                                                        var v35 uint64 = 42
                                                        v35 -= (v35 >> ((255) & 63))
                                                        if len(v3) > 0 {
                                                            v3[0] = (^(v2))
                                                        }
                                                        _ = v35
                                                    } else {
                                                        v36 := GenS4{F0: 201, F1: 606, F2: (v1 && v1)}
                                                        if (18446744073709551615 == 54) {
                                                            switch ((("" + ("a" + ("hello" + "foo bar"))) + "a") + "") {
                                                            case "\n\t":
                                                                fmt.Println(v7)
                                                            }
                                                            v16 += v16
                                                        } else {
                                                            if (1 == 2) {
                                                                fmt.Println(v2)
                                                                v34 = (v7 + 44)
                                                            }
                                                            switch ("hello" + "foo bar") {
                                                            case "a":
                                                                v27 = v7
                                                            case "hello":
                                                                var v37 float64 = ((((1.5 - (-2.25 / (1e10 - 1.5))) / -2.25) - 0.1) - 1e10)
                                                                _ = v37
                                                            default:
                                                                if true {
                                                                    v3 = append(v3, (^(v2)))
                                                                    var v38 uint32 = (v2 & (v2 & (v2 * v2)))
                                                                    _ = v38
                                                                } else {
                                                                    v2 = v2
                                                                }
                                                            }
                                                            var v39 int = fi19
                                                            _ = v39
                                                        }
                                                        v40 := GenS4{F0: 2, F1: 407, F2: v1}
                                                        _ = v36
                                                        _ = v40
                                                    }
                                                    var v41 uint64 = 838
                                                    _ = len(v3)
                                                    _ = v33
                                                    _ = v34
                                                    _ = v41
                                                case -980:
                                                    if v1 {
                                                        if (float64(-2.25) == 0.0) {
                                                            v42 := GenS4{F0: 618, F1: 1, F2: v1}
                                                            v27 = int64(v2)
                                                            _ = v42
                                                        } else {
                                                            _ = len(v3)
                                                        }
                                                        if (!(v7 == (v27 ^ v27))) {
                                                            if v1 {
                                                                fmt.Println(v1)
                                                                var v43 uint64 = 42
                                                                switch "foo bar" {
                                                                case "\n\t":
                                                                    v44 := GenS3{F0: 72, F1: uint64(v2), F2: v2}
                                                                    if v1 {
                                                                        var v45 int = (fi20 << ((818) & 63))
                                                                        _ = v45
                                                                    } else {
                                                                        fmt.Println(v1, v43)
                                                                        v46 := GenS4{F0: (v43 / ((v43) | 1)), F1: 65535, F2: ((fi19 >> ((843) & 63)) < fi19)}
                                                                        _ = v46
                                                                    }
                                                                    fmt.Println(v16, v16)
                                                                    v27 = 0
                                                                    _ = v44
                                                                default:
                                                                    fmt.Println(v16, v2, v43)
                                                                }
                                                                var v47 int8 = 97
                                                                var v48 float32 = (-2.25 + -2.25)
                                                                _ = v43
                                                                _ = v47
                                                                _ = v48
                                                            } else {
                                                                if false {
                                                                    if v1 {
                                                                        v49 := GenS4{F0: 1, F1: 255, F2: (v1 || false)}
                                                                        switch (^(fi20)) {
                                                                        case -2147483648:
                                                                            if len(v3) > 0 {
                                                                                v3[0] = 4294967295
                                                                            }
                                                                        case 146:
                                                                            var v50 uint16 = 32767
                                                                            _ = v50
                                                                        case -639:
                                                                            v51 := GenS5{F0: -866, F1: (v2 - 722)}
                                                                            _ = v51
                                                                        }
                                                                        v16 += v16
                                                                        v52 := GenS4{F0: 2, F1: 735, F2: true}
                                                                        _ = v49
                                                                        _ = v52
                                                                    }
                                                                    var v53 int8 = 22
                                                                    var v54 uint32 = v2
                                                                    _ = v53
                                                                    _ = v54
                                                                }
                                                                if (true && ((v1 || ((fi20 ^ fi19) == fi19)) || (18446744073709551615 <= 1))) {
                                                                    fmt.Println(v2)
                                                                } else {
                                                                    switch 32768 {
                                                                    case -869:
                                                                        if (!true) {
                                                                            if v1 {
                                                                                var v55 bool = (1 < 2)
                                                                                var v56 uint8 = 240
                                                                                _ = v55
                                                                                _ = v56
                                                                            } else {
                                                                                v1 = v1
                                                                                switch "a" {
                                                                                case "foo bar":
                                                                                    v27 += (v7 << ((575) & 63))
                                                                                    var v57 uint64 = 1
                                                                                    v57 = (v57 & (v57 / ((4294967295) | 1)))
                                                                                    _ = v57
                                                                                case "":
                                                                                    switch "" {
                                                                                    case "a":
                                                                                        _ = len(v3)
                                                                                    }
                                                                                case "hello":
                                                                                    var v58 int8 = -35
                                                                                    _ = v58
                                                                                }
                                                                                fmt.Println(v1)
                                                                            }
                                                                            v59 := GenS6{F0: "a", F1: v2}
                                                                            _ = v59
                                                                        }
                                                                    case -9223372036854775808:
                                                                        if len(v3) > 0 {
                                                                            v3[0] = v2
                                                                        }
                                                                    case 300:
                                                                        v60 := GenS2{F0: 223}
                                                                        _ = v60
                                                                    default:
                                                                        var v61 uint32 = (^(v2))
                                                                        _ = v61
                                                                    }
                                                                    switch ("foo bar" + ("foo bar" + "a")) {
                                                                    case "foo bar":
                                                                        if len(v3) > 0 {
                                                                            v3[0] = v2
                                                                        }
                                                                    case "hello":
                                                                        v7 -= (v27 >> ((4294967295) & 63))
                                                                    default:
                                                                        var v62 string = ("foo bar" + ("foo bar" + "\n\t"))
                                                                        _ = v62
                                                                    }
                                                                    switch "\n\t" {
                                                                    case "hello":
                                                                        if (!((255 >= 42) && v1)) {
                                                                            v16 += (v16 % ((32767) | 1))
                                                                        }
                                                                    case "":
                                                                        fmt.Println(v7, v2, v1)
                                                                    }
                                                                }
                                                            }
                                                            var v63 int64 = v7
                                                            if (18446744073709551615 >= 625) {
                                                                if len(v3) > 0 {
                                                                    v3[0] = v2
                                                                }
                                                            } else {
                                                                fmt.Println(fi19, v63, v63)
                                                            }
                                                            _ = v63
                                                        } else {
                                                            if v1 {
                                                                _ = len(v3)
                                                            } else {
                                                                _ = len(v3)
                                                                var v64 int8 = -1
                                                                v65 := GenS6{F0: "hello", F1: 0}
                                                                _ = v64
                                                                _ = v65
                                                            }
                                                            v66 := GenS7{F0: -688, F1: 960}
                                                            v67 := GenS8{F0: 42}
                                                            v16 += (-(v16))
                                                            _ = v66
                                                            _ = v67
                                                        }
                                                        v7 = (v27 >> ((2) & 63))
                                                    } else {
                                                        switch 127 {
                                                        case -9223372036854775808:
                                                            fmt.Println(v1)
                                                            switch -691 {
                                                            case 309:
                                                                switch "hello" {
                                                                case "a":
                                                                    if len(v3) >= 2 {
                                                                        v3 = v3[1:len(v3)]
                                                                    }
                                                                case "":
                                                                    var v68 float32 = -2.25
                                                                    _ = v68
                                                                default:
                                                                    if len(v3) >= 2 {
                                                                        v3 = v3[1:len(v3)]
                                                                    }
                                                                }
                                                            default:
                                                                if len(v3) > 0 {
                                                                    v3[0] = (^(v2))
                                                                }
                                                            }
                                                            if len(v3) >= 2 {
                                                                v3 = v3[1:len(v3)]
                                                            }
                                                            v27 = (v27 << ((255) & 63))
                                                            v16 -= v16
                                                        case 0:
                                                            v2 = 0
                                                        case 68:
                                                            v69 := GenS9{F0: (v7 + v27), F1: (432 <= 52)}
                                                            _ = v69
                                                        }
                                                        v3 = append(v3, 237)
                                                    }
                                                default:
                                                    v2 += 2147483648
                                                }
                                            } else {
                                                if len(v3) > 0 {
                                                    v3[0] = v2
                                                }
                                                if v1 {
                                                    if false {
                                                        switch ("\n\t" + (((("" + "foo bar") + "hello") + ("" + "hello")) + "foo bar")) {
                                                        case "\n\t":
                                                            v70 := GenS10{F0: 24, F1: (^(v7)), F2: (v2 * (v2 << ((18446744073709551615) & 63)))}
                                                            _ = v70
                                                        case "foo bar":
                                                            v71 := GenS5{F0: 1, F1: v2}
                                                            _ = v71
                                                        default:
                                                            if false {
                                                                switch ("" + "hello") {
                                                                case "hello":
                                                                    var v72 int = fi20
                                                                    _ = v72
                                                                case "\n\t":
                                                                    if len(v3) > 0 {
                                                                        v3[0] = v2
                                                                    }
                                                                default:
                                                                    if (true || ((v16 <= v16) && ((v2 + (^(v2))) <= 2147483648))) {
                                                                        v27 -= (v7 + (v7 * (v27 % (((v7 >> ((443) & 63))) | 1))))
                                                                    } else {
                                                                        switch 1 {
                                                                        case 1:
                                                                            v16 += (v16 % ((v16) | 1))
                                                                        case 253:
                                                                            var v73 int32 = -1
                                                                            _ = v73
                                                                        case 32768:
                                                                            var v74 uint8 = 127
                                                                            _ = v74
                                                                        default:
                                                                            if v1 {
                                                                                var v75 string = "a"
                                                                                switch v75 {
                                                                                case "":
                                                                                    var v76 int8 = 1
                                                                                    _ = v76
                                                                                case "a":
                                                                                    var v77 uint = 255
                                                                                    _ = v77
                                                                                }
                                                                                if len(v3) > 0 {
                                                                                    v3[0] = v2
                                                                                }
                                                                                _ = v75
                                                                            }
                                                                        }
                                                                        fmt.Println(fi19, v16, v27)
                                                                    }
                                                                }
                                                                var v78 uint16 = 999
                                                                v7 = v7
                                                                var v79 bool = ("" <= ("foo bar" + "\n\t"))
                                                                var v80 bool = (210 <= 233)
                                                                _ = v78
                                                                _ = v79
                                                                _ = v80
                                                            } else {
                                                                if len(v3) > 0 {
                                                                    v3[0] = (v2 - 42)
                                                                }
                                                                v7 += 767
                                                                var v81 uint64 = uint64(v2)
                                                                _ = v81
                                                            }
                                                        }
                                                        v82 := GenS6{F0: "a", F1: v2}
                                                        v1 = (v1 && (!false))
                                                        _ = v82
                                                    } else {
                                                        var v83 float64 = (1.5 + (0.0 / (float64(float32(1.5)) - ((0.0 * 1.5) * (1.5 * 1e10)))))
                                                        _ = v83
                                                    }
                                                } else {
                                                    fmt.Println(v7, v2)
                                                    _ = len(v3)
                                                    v84 := GenS10{F0: 2, F1: (v7 << ((1) & 63)), F2: v2}
                                                    _ = v84
                                                }
                                            }
                                        }
                                    case 0:
                                        fmt.Println(v1, v16, v2)
                                    }
                                    v27 -= (v7 << ((125) & 63))
                                    var v85 uint64 = 42
                                    switch "\n\t" {
                                    case "a":
                                        if ((v16 % ((v16) | 1)) <= v16) {
                                            v1 = v1
                                            switch fi19 {
                                            case 42:
                                                v86 := GenS1{F0: 112, F1: v1}
                                                switch fi20 {
                                                case 9223372036854775807:
                                                    v2 += (v2 + 0)
                                                case 32768:
                                                    switch (fi20 ^ -9223372036854775808) {
                                                    case -801:
                                                        fmt.Println(v2, v16)
                                                    case 127:
                                                        var v87 int = (fi19 * (-(fi19)))
                                                        _ = v87
                                                    default:
                                                        v2 = (v2 + v2)
                                                    }
                                                case -2147483648:
                                                    v2 -= v2
                                                }
                                                _ = v86
                                            case -2147483648:
                                                v88 := GenS1{F0: 0, F1: (2 > 2)}
                                                _ = v88
                                            default:
                                                fmt.Println(v27, v7)
                                            }
                                        }
                                        v2 += v2
                                        var v89 uint64 = (v85 + v85)
                                        _ = v89
                                    case "\n\t":
                                        fmt.Println(fi20, v1)
                                    }
                                    if (!v1) {
                                        v90 := GenS11{F0: -1, F1: 42, F2: (v27 >> ((255) & 63))}
                                        switch "a" {
                                        case "hello":
                                            var v91 string = ("" + "")
                                            _ = v91
                                        default:
                                            switch "hello" {
                                            case "":
                                                v2 += v2
                                            }
                                        }
                                        switch "\n\t" {
                                        case "\n\t":
                                            var v92 float32 = (((0.0 / -2.25) + float32(0.0)) - (3.14159 + (0.0 / (-2.25 - (-2.25 + (-2.25 / 3.14159))))))
                                            _ = v92
                                        case "hello":
                                            var v93 float32 = (float32(0.0) - 1.5)
                                            _ = v93
                                        case "":
                                            v94 := GenS11{F0: 42, F1: 113, F2: -1}
                                            _ = v94
                                        default:
                                            var v95 bool = ("foo bar" >= ("" + "foo bar"))
                                            _ = v95
                                        }
                                        _ = v90
                                    }
                                    _ = v85
                                }
                                _ = v26
                                _ = v27
                            default:
                                v96 := GenS1{F0: 42, F1: v1}
                                _ = v96
                            }
                            v16 = (v16 >> ((1) & 63))
                        }
                    }
                }
                v97 := GenS10{F0: 127, F1: v7, F2: (v2 / ((v2) | 1))}
                _ = fi20
                _ = v97
            }
            v16 = (v16 & v16)
            if (295 < 2) {
                switch fi19 {
                case 4294967295:
                    v16 -= v16
                default:
                    switch fi19 {
                    case -651:
                        v1 = true
                    case 127:
                        v98 := GenS12{F0: 255}
                        _ = v98
                    case 1:
                        v99 := GenS13{F0: (v16 / ((v16) | 1))}
                        _ = v99
                    default:
                        v100 := GenS11{F0: 841, F1: -73, F2: -520}
                        _ = v100
                    }
                }
                switch ("foo bar" + (("a" + "") + "\n\t")) {
                case "":
                    var v101 uint32 = v2
                    _ = v101
                }
                v7 += (v7 / ((int64(fi19)) | 1))
            } else {
                if v1 {
                    switch "hello" {
                    case "foo bar":
                        var v102 float64 = (0.1 + ((3.14159 / (((0.1 / 3.14159) * 3.14159) + 3.14159)) - 3.14159))
                        _ = v102
                    case "hello":
                        v7 += v7
                    case "\n\t":
                        var v103 uint8 = uint8(v2)
                        _ = v103
                    }
                    v2 += (^(v2))
                }
                v104 := GenS14{F0: 86}
                _ = v104
            }
            _ = fi19
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v7
}
