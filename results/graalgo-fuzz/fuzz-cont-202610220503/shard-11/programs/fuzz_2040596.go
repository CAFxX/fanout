package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64 }
type GenS2 struct { F0 int8; F1 int64 }
type GenS3 struct { F0 int; F1 uint }
type GenS4 struct { F0 int8; F1 uint16 }
type GenS5 struct { F0 uint32; F1 uint16; F2 bool }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 int64; F1 uint8 }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 int64; F1 uint32; F2 uint }
type GenS11 struct { F0 uint64 }
type GenS12 struct { F0 string }
type GenS13 struct { F0 int64; F1 int16 }
type GenS14 struct { F0 uint32; F1 int16 }
func genFn0(p0 uint, p1 uint8) uint8 {
    for fi7 := 0; fi7 < 18; fi7++ {
        fmt.Println(fi7)
        if (p1 < p1) {
            v8 := GenS1{F0: -541, F1: -530}
            fmt.Println(p0)
            v9 := GenS2{F0: 127, F1: 9223372036854775807}
            _ = v8
            _ = v9
        }
        v10 := GenS2{F0: -1, F1: 1}
        _ = fi7
        _ = v10
    }
    v11 := []int{128, 128, 128, 42}
    fmt.Println(p1, p1)
    _ = p0
    _ = p1
    _ = v11
    return 144
}
func genFn1(p0 uint, p1 uint, p2 uint32) string {
    for fi12 := 0; fi12 < 11; fi12++ {
        var v13 bool = false
        var v14 int64 = 9223372036854775807
        for fi15 := 0; fi15 < 20; fi15++ {
            v16 := []int{0, fi15}
            if true {
                if v13 {
                    v17 := GenS3{F0: fi15, F1: p0}
                    _ = v17
                }
            } else {
                if (v13 && ("hello" == "\n\t")) {
                    if (genFn0(p0, 255) < 0) {
                        v13 = v13
                        var v18 int64 = v14
                        if true {
                            fmt.Println(fi12)
                            p0 = p0
                            v19 := GenS4{F0: 100, F1: 457}
                            _ = v19
                        } else {
                            v20 := GenS4{F0: -1, F1: 42}
                            _ = v20
                        }
                        _ = v18
                    } else {
                        fmt.Println(fi12, p1)
                        switch "a" {
                        case "hello":
                            var v21 float64 = (float64((1e10 + 0.0)) / 3.14159)
                            _ = v21
                        case "a":
                            if (32767 >= 270) {
                                var v22 int = fi12
                                _ = v22
                            } else {
                                p2 = 2147483647
                            }
                        case "\n\t":
                            if ((-(fi12)) == -2147483648) {
                                var v23 int = -1
                                _ = v23
                            } else {
                                if len(v16) > 0 {
                                    v16[0] = (fi15 | fi12)
                                }
                                v16 = append(v16, fi12)
                            }
                        }
                    }
                    v13 = v13
                } else {
                    v24 := GenS5{F0: p2, F1: 32767, F2: v13}
                    _ = v24
                }
            }
            v25 := GenS5{F0: p2, F1: 32767, F2: v13}
            _ = fi15
            _ = v16
            _ = v25
        }
        _ = fi12
        _ = v13
        _ = v14
    }
    var v26 float32 = (0.1 / -2.25)
    v27 := []int{-9223372036854775808}
    _ = p0
    _ = p1
    _ = p2
    _ = v26
    _ = v27
    return "a"
}

func main() {
    var v1 string = "foo bar"
    var v2 uint8 = 1
    v3 := []int{42}
    var v4 int64 = -9223372036854775808
    var v5 uint64 = 580
    v6 := []uint32{353}
    if true {
        var v28 int64 = -176
        var v29 string = v1
        _ = v28
        _ = v29
    } else {
        if true {
            fmt.Println(v4, v2, v1)
            var v30 int64 = v4
            v1 = (v1 + v1)
            _ = v30
        } else {
            if (0 >= 127) {
                v4 += (v4 << ((200) & 63))
                var v31 int64 = -9223372036854775808
                for fi32 := 0; fi32 < 5; fi32++ {
                    var v33 int16 = 951
                    for fi34 := 0; fi34 < 16; fi34++ {
                        if true {
                            var v35 int = (fi32 / ((fi32) | 1))
                            v3 = append(v3, v35)
                            var v36 uint8 = genFn0(255, v2)
                            _ = v35
                            _ = v36
                        }
                        var v37 uint = 327
                        v4 = v4
                        fmt.Println(v4, fi34, v4)
                        var v38 int32 = -1
                        v39 := GenS6{F0: v37}
                        switch -840 {
                        case 9223372036854775807:
                            switch fi34 {
                            case -530:
                                switch fi34 {
                                case 4294967295:
                                    switch (("" + (v1 + (v1 + "foo bar"))) + v1) {
                                    case "\n\t":
                                        fmt.Println(v33, v1, v37)
                                    case "a":
                                        if (!(v38 <= (v38 & int32(v37)))) {
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                            var v40 uint32 = 65535
                                            var v41 uint64 = 1
                                            _ = v40
                                            _ = v41
                                        } else {
                                            v42 := GenS7{F0: (v33 & v33)}
                                            if len(v3) > 0 {
                                                v3[0] = -462
                                            }
                                            _ = v42
                                        }
                                        switch (v1 + (v1 + ("\n\t" + v1))) {
                                        case "a":
                                            fmt.Println(v37)
                                        case "\n\t":
                                            v43 := GenS1{F0: 2147483647, F1: v31}
                                            _ = v43
                                        }
                                    case "hello":
                                        var v44 uint8 = 0
                                        _ = v44
                                    default:
                                        var v45 uint16 = 65535
                                        _ = v45
                                    }
                                }
                            }
                        case 42:
                            v46 := GenS8{F0: -9223372036854775808, F1: (v2 << (((v37 + 935)) & 63))}
                            _ = v46
                        default:
                            fmt.Println(v2, fi34)
                        }
                        _ = fi34
                        _ = v37
                        _ = v38
                        _ = v39
                    }
                    _ = fi32
                    _ = v33
                }
                _ = v31
            }
            _ = len(v6)
        }
        if false {
            var v47 int = 4294967295
            for fi48 := 0; fi48 < 13; fi48++ {
                switch fi48 {
                case 128:
                    v5 -= v5
                case 1:
                    switch "\n\t" {
                    case "hello":
                        switch v1 {
                        case "":
                            fmt.Println(v1, v2)
                        case "a":
                            v49 := GenS9{F0: 42}
                            _ = v49
                        case "foo bar":
                            if ((v2 << ((261) & 63)) != (v2 | genFn0(42, uint8(v5)))) {
                                v47 -= 32767
                            }
                        default:
                            for fi50 := 0; fi50 < 12; fi50++ {
                                switch v47 {
                                case -632:
                                    var v51 float32 = 3.14159
                                    v5 += v5
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    _ = v51
                                case 127:
                                    v1 = (v1 + ((("foo bar" + v1) + (("\n\t" + v1) + ("" + (((v1 + ("a" + ("foo bar" + "\n\t"))) + "\n\t") + v1)))) + ("hello" + v1)))
                                case 4294967295:
                                    var v52 uint8 = (v2 & (v2 * (v2 % ((74) | 1))))
                                    _ = v52
                                }
                                v53 := GenS3{F0: (-(fi48)), F1: 0}
                                var v54 float32 = 1.5
                                v55 := GenS10{F0: int64(v2), F1: 2147483647, F2: 42}
                                switch fi48 {
                                case 914:
                                    if (0.0 >= 1.5) {
                                        if ((v5 % (((v5 % (((v5 & 856)) | 1))) | 1)) > 2) {
                                            if (!false) {
                                                fmt.Println(v5)
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                                fmt.Println(v54, v4)
                                            } else {
                                                if false {
                                                    v54 += ((v54 + 0.1) + v54)
                                                }
                                                var v56 uint = 65535
                                                if false {
                                                    var v57 string = (v1 + v1)
                                                    var v58 uint = v56
                                                    var v59 float32 = v54
                                                    _ = v57
                                                    _ = v58
                                                    _ = v59
                                                }
                                                var v60 float32 = v54
                                                _ = v56
                                                _ = v60
                                            }
                                            var v61 uint64 = 729
                                            if (float32(0.1) == 3.14159) {
                                                fmt.Println(fi48, v54, v61)
                                                v3 = append(v3, 503)
                                            } else {
                                                if len(v3) > 0 {
                                                    v3[0] = 443
                                                }
                                                fmt.Println(v61, v1, v2)
                                                if true {
                                                    if true {
                                                        if (!false) {
                                                            _ = len(v3)
                                                        }
                                                        v2 = v2
                                                        switch (v1 + (("\n\t" + ((v1 + v1) + (v1 + "foo bar"))) + v1)) {
                                                        case "hello":
                                                            switch v1 {
                                                            case "":
                                                                v62 := GenS11{F0: v5}
                                                                _ = v62
                                                            }
                                                        default:
                                                            if false {
                                                                v54 += (v54 * (v54 + float32(3.14159)))
                                                            } else {
                                                                var v63 uint64 = 270
                                                                var v64 int64 = v4
                                                                v6 = append(v6, 2)
                                                                _ = v63
                                                                _ = v64
                                                            }
                                                        }
                                                        var v65 float64 = 1.5
                                                        _ = v65
                                                    }
                                                }
                                            }
                                            _ = v61
                                        }
                                        v66 := GenS8{F0: 1, F1: v2}
                                        v67 := GenS9{F0: 1}
                                        _ = v66
                                        _ = v67
                                    }
                                }
                                var v68 int64 = v4
                                _ = fi50
                                _ = v53
                                _ = v54
                                _ = v55
                                _ = v68
                            }
                        }
                    }
                }
                _ = fi48
            }
            _ = v47
        }
    }
    if (-1 >= 0) {
        for fi69 := 0; fi69 < 4; fi69++ {
            if false {
                for fi70 := 0; fi70 < 8; fi70++ {
                    switch "foo bar" {
                    case "hello":
                        _ = len(v6)
                    }
                    _ = fi70
                }
            }
            var v71 uint64 = v5
            v72 := GenS12{F0: "foo bar"}
            _ = fi69
            _ = v71
            _ = v72
        }
        if (1e10 != 0.0) {
            v73 := GenS2{F0: -87, F1: v4}
            _ = v73
        } else {
            for fi74 := 0; fi74 < 7; fi74++ {
                var v75 uint64 = (v5 >> ((uint(v5)) & 63))
                _ = fi74
                _ = v75
            }
            v76 := GenS9{F0: 719}
            _ = v76
        }
        switch 1 {
        case 1:
            if true {
                v1 = ("a" + v1)
            }
        }
    }
    for fi77 := 0; fi77 < 1; fi77++ {
        switch (v1 + v1) {
        case "\n\t":
            for fi78 := 0; fi78 < 11; fi78++ {
                v2 = v2
                v5 += v5
                _ = fi78
            }
        default:
            if false {
                v2 -= (v2 ^ 255)
            } else {
                v79 := GenS13{F0: 361, F1: 32767}
                _ = v79
            }
        }
        _ = fi77
    }
    for fi80 := 0; fi80 < 12; fi80++ {
        var v81 uint32 = 42
        switch (v1 + v1) {
        case "\n\t":
            v6 = append(v6, 901)
        case "a":
            var v82 float32 = 3.14159
            _ = v82
        default:
            fmt.Println(v4, fi80, fi80)
        }
        _ = fi80
        _ = v81
    }
    var v83 int = 561
    fmt.Println(v5, v4, v2)
    if false {
        switch v1 {
        case "a":
            if ((18446744073709551615 == (v5 * (^(v5)))) || (2147483647 != 65535)) {
                var v84 int = v83
                fmt.Println(v83)
                _ = v84
            } else {
                var v85 int16 = 718
                _ = v85
            }
        }
        var v86 uint = 65535
        v4 -= v4
        _ = v86
    } else {
        switch v1 {
        case "\n\t":
            v83 += -66
        case "hello":
            if (!(v4 == v4)) {
                var v87 float32 = 1.5
                _ = v87
            }
        default:
            if (!true) {
                switch (^(v83)) {
                case 218:
                    for fi88 := 0; fi88 < 7; fi88++ {
                        v5 += 0
                        _ = fi88
                    }
                case 0:
                    var v89 float32 = (0.0 / 3.14159)
                    _ = v89
                }
                var v90 float32 = 0.0
                _ = v90
            } else {
                fmt.Println(v2, v5, v2)
                v91 := GenS5{F0: 2147483648, F1: 0, F2: false}
                if len(v3) > 0 {
                    v3[0] = (v83 + v83)
                }
                _ = v91
            }
        }
        var v92 uint64 = v5
        var v93 int64 = v4
        _ = v92
        _ = v93
    }
    for fi94 := 0; fi94 < 11; fi94++ {
        v95 := GenS14{F0: 4294967295, F1: 42}
        v1 = ""
        if (uint(v5) <= 1) {
            var v96 float64 = ((0.1 + 0.1) * 1e10)
            _ = v96
        } else {
            if ((v83 & (v83 & (fi94 | v83))) > (fi94 % (((^(v83))) | 1))) {
                if ((fi94 | int(v4)) != (-(fi94))) {
                    if (!false) {
                        fmt.Println(fi94, v2)
                    }
                    v1 = v1
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                }
                var v97 string = "hello"
                if true {
                    var v98 int32 = 0
                    var v99 int = (v83 ^ (fi94 >> ((255) & 63)))
                    _ = v98
                    _ = v99
                }
                _ = v97
            }
            var v100 int64 = int64(v5)
            _ = v100
        }
        _ = fi94
        _ = v95
    }
    v3 = append(v3, v83)
    if (!(true || false)) {
        var v101 int16 = 42
        _ = v101
    } else {
        switch 32768 {
        case 80:
            v5 = uint64(v2)
        case -9223372036854775808:
            if false {
                var v102 float32 = 1e10
                fmt.Println(v1, v2)
                v5 -= v5
                _ = v102
            } else {
                switch ("foo bar" + ("\n\t" + "")) {
                case "foo bar":
                    var v103 float64 = ((1e10 + float64((((0.1 + 0.1) * (1.5 / (-2.25 * float32(float64(((1e10 * float32((float64(1.5) - 0.1))) * 0.1)))))) / 1e10))) * 1.5)
                    _ = v103
                default:
                    var v104 uint32 = 2147483648
                    _ = v104
                }
            }
        case -135:
            v2 += v2
        default:
            for fi105 := 0; fi105 < 6; fi105++ {
                var v106 uint = 18446744073709551615
                v2 = 2
                for fi107 := 0; fi107 < 20; fi107++ {
                    switch (v1 + v1) {
                    case "\n\t":
                        var v108 int8 = 127
                        _ = v108
                    }
                    _ = fi107
                }
                _ = fi105
                _ = v106
            }
        }
        v5 = (v5 % ((v5) | 1))
    }
    v83 -= v83
    for fi109 := 0; fi109 < 8; fi109++ {
        fmt.Println(v4, v83)
        _ = fi109
    }
    for fi110 := 0; fi110 < 3; fi110++ {
        var v111 uint16 = 1
        v112 := GenS8{F0: (v4 % (((-(v4))) | 1)), F1: (^(v2))}
        _ = fi110
        _ = v111
        _ = v112
    }
    if false {
        fmt.Println(v4, v83)
        v113 := GenS12{F0: (v1 + v1)}
        _ = v113
    }
    v1 = (v1 + v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v83:", v83)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v83
}
