package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16; F2 uint8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint32; F1 int16 }
type GenS4 struct { F0 int32; F1 int8; F2 string }
type GenS5 struct { F0 uint16; F1 uint64 }
type GenS6 struct { F0 uint32; F1 uint8 }
type GenS7 struct { F0 uint32; F1 int64 }
type GenS8 struct { F0 int8; F1 int }
type GenS9 struct { F0 uint8 }
type GenS10 struct { F0 int16; F1 uint8 }
type GenS11 struct { F0 int32 }
type GenS12 struct { F0 uint8 }
type GenS13 struct { F0 int32; F1 uint8 }
type GenS14 struct { F0 uint16 }
type GenS15 struct { F0 uint32 }
type GenS16 struct { F0 uint; F1 int64; F2 uint64 }
type GenS17 struct { F0 int64; F1 string }
type GenS18 struct { F0 uint64; F1 int32; F2 bool }
type GenS19 struct { F0 string }
type GenS20 struct { F0 bool; F1 uint64 }
type GenS21 struct { F0 uint }
type GenS22 struct { F0 uint8; F1 int64; F2 uint64 }
type GenS23 struct { F0 string; F1 uint64; F2 bool }
type GenS24 struct { F0 bool; F1 int16 }
type GenS25 struct { F0 uint16; F1 uint8; F2 uint32 }
type GenS26 struct { F0 uint16; F1 uint; F2 uint32 }
type GenS27 struct { F0 string; F1 uint32; F2 uint }
type GenS28 struct { F0 string; F1 uint8 }
type GenS29 struct { F0 int32 }
type GenS30 struct { F0 uint32 }
func genFn0(p0 int16, p1 int8, p2 uint32) bool {
    var v5 int8 = (p1 >> ((65535) & 63))
    switch (("foo bar" + "foo bar") + "foo bar") {
    case "hello":
        p1 = 74
    }
    if (false && true) {
        var v6 int = 32767
        p1 = 127
        _ = v6
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return (uint8(p2) < 1)
}
func genFn1(p0 uint, p1 uint8, p2 uint64) uint8 {
    v7 := []int{}
    for fi8 := 0; fi8 < 16; fi8++ {
        p1 = 0
        _ = fi8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return (^(p1))
}

func main() {
    v1 := make(map[string]bool)
    var v2 float32 = 3.14159
    var v3 int64 = 9223372036854775807
    var v4 uint32 = 33
    v9 := []int{}
    v2 = v2
    v10 := GenS1{F0: 0, F1: 2, F2: 2}
    v11 := GenS1{F0: -968, F1: 255, F2: 128}
    v12 := GenS1{F0: 70, F1: 42, F2: 0}
    if (v2 != (v2 + ((v2 * v2) + 0.0))) {
        var v13 uint64 = 2
        v14 := GenS1{F0: 338, F1: 963, F2: 226}
        _ = v13
        _ = v14
    }
    for fi15 := 0; fi15 < 13; fi15++ {
        if ((fi15 - fi15) == (fi15 << ((18446744073709551615) & 63))) {
            v4 += (v4 + 589)
        } else {
            for fi16 := 0; fi16 < 13; fi16++ {
                fmt.Println(fi15, fi15)
                if (!true) {
                    var v17 float32 = (0.1 * v2)
                    _ = v17
                }
                _ = fi16
            }
            for fi18 := 0; fi18 < 6; fi18++ {
                switch "\n\t" {
                case "foo bar":
                    v1["foo bar"] = true
                case "":
                    if true {
                        var v19 uint64 = 170
                        switch "a" {
                        case "foo bar":
                            var v20 uint32 = (v4 & 0)
                            _ = v20
                        case "\n\t":
                            if ((float32(0.0) >= (v2 * (float32(3.14159) - (1.5 * v2)))) || true) {
                                v4 = v4
                                v2 = (v2 / (float32(3.14159) / (v2 + float32((0.1 - float64(((v2 - v2) + float32(0.1))))))))
                                var v21 float32 = ((v2 + v2) + (v2 + v2))
                                v21 -= v21
                                var v22 int8 = 1
                                v23 := GenS1{F0: 4294967295, F1: 2, F2: genFn1(18446744073709551615, 2, v19)}
                                _ = v21
                                _ = v22
                                _ = v23
                            } else {
                                v19 = (v19 / ((v19) | 1))
                                if (625 > v4) {
                                    if (!(!false)) {
                                        if true {
                                            v24 := GenS1{F0: fi15, F1: 109, F2: 11}
                                            _ = v24
                                        }
                                        v25 := GenS1{F0: 0, F1: 431, F2: 70}
                                        _ = v25
                                    }
                                    if ((-2.25 > 0.1) && true) {
                                        delete(v1, "foo bar")
                                        var v26 uint8 = 127
                                        _ = v26
                                    } else {
                                        if true {
                                            v3 = -1
                                        } else {
                                            switch ((((("" + "a") + "\n\t") + "hello") + (("\n\t" + "hello") + "hello")) + "") {
                                            case "foo bar":
                                                switch (fi15 | (fi15 % (((fi18 % (((^(fi15))) | 1))) | 1))) {
                                                case 4294967295:
                                                    fmt.Println(fi15)
                                                }
                                                v19 -= 18446744073709551615
                                                v2 = 3.14159
                                            case "":
                                                v2 = 1.5
                                            case "\n\t":
                                                v3 = (v3 + 1)
                                            default:
                                                if len(v9) >= 2 {
                                                    v9 = v9[1:len(v9)]
                                                }
                                            }
                                            v3 += 9223372036854775807
                                        }
                                        var v27 int8 = 111
                                        var v28 uint16 = 65535
                                        _ = v27
                                        _ = v28
                                    }
                                    if true {
                                        fmt.Println(v19, v3)
                                        var v29 int16 = -1
                                        _ = v29
                                    }
                                } else {
                                    v30 := GenS1{F0: fi18, F1: 0, F2: 42}
                                    switch (fi18 - (-(fi18))) {
                                    case 4294967295:
                                        v19 = (v19 | 0)
                                    case 32767:
                                        if ((v4 * v4) > 2) {
                                            _ = len(v1)
                                            fmt.Println(v19)
                                            if (!false) {
                                                var v31 string = "a"
                                                _ = v31
                                            }
                                        } else {
                                            switch "hello" {
                                            case "":
                                                fmt.Println(v3, v19, v19)
                                            }
                                        }
                                    }
                                    _ = v30
                                }
                                var v32 bool = false
                                _ = v32
                            }
                        case "":
                            if false {
                                v33 := GenS1{F0: fi18, F1: 32767, F2: 1}
                                switch (fi15 * (fi18 % (((fi15 - 128)) | 1))) {
                                case 2147483647:
                                    v34 := v1[("a" + "\n\t")]
                                    v35 := GenS2{F0: ("a" + ("hello" + ""))}
                                    fmt.Println(v2, fi18, v34)
                                    v4 = v4
                                    _ = v34
                                    _ = v35
                                case 616:
                                    switch "foo bar" {
                                    case "":
                                        if true {
                                            if (-562 <= -32768) {
                                                switch 3 {
                                                case 1:
                                                    switch (fi15 / ((fi18) | 1)) {
                                                    case 933:
                                                        var v36 float64 = 1e10
                                                        _ = v36
                                                    case 0:
                                                        var v37 uint16 = 42
                                                        _ = v37
                                                    default:
                                                        if (1000 > -32768) {
                                                            var v38 float64 = ((1.5 / (-2.25 * ((float64(((float32(-2.25) - v2) - v2)) / 1e10) + (0.0 - 1.5)))) - 3.14159)
                                                            _ = v38
                                                        }
                                                    }
                                                case 4294967295:
                                                    v4 = 130
                                                    fmt.Println(v19, v3)
                                                    v39 := GenS2{F0: "\n\t"}
                                                    _ = v39
                                                case 2147483647:
                                                    var v40 int = (-(fi15))
                                                    _ = v40
                                                default:
                                                    _ = len(v9)
                                                }
                                                if (v19 > v19) {
                                                    var v41 int64 = v3
                                                    _ = v41
                                                } else {
                                                    var v42 int32 = int32(v4)
                                                    _ = v42
                                                }
                                                switch fi15 {
                                                case 32767:
                                                    v43 := GenS3{F0: 571, F1: 32767}
                                                    _ = v43
                                                case 4294967295:
                                                    fmt.Println(v3)
                                                    if false {
                                                        var v44 int32 = 1
                                                        if true {
                                                            v44 = 907
                                                            if false {
                                                                var v45 string = "a"
                                                                if (!false) {
                                                                    switch v45 {
                                                                    case "hello":
                                                                        v19 -= (v19 >> ((500) & 63))
                                                                    case "foo bar":
                                                                        if false {
                                                                            v46 := GenS4{F0: 0, F1: 119, F2: v45}
                                                                            switch fi15 {
                                                                            case 0:
                                                                                _ = len(v1)
                                                                            case -2147483648:
                                                                                if true {
                                                                                    if (65535 > 18446744073709551615) {
                                                                                        v44 = v44
                                                                                    } else {
                                                                                        var v47 int32 = (v44 / (((v44 | (-(v44)))) | 1))
                                                                                        _ = v47
                                                                                    }
                                                                                    fmt.Println(v4)
                                                                                } else {
                                                                                    fmt.Println(fi15, fi18)
                                                                                    var v48 int = 9223372036854775807
                                                                                    _ = v48
                                                                                }
                                                                            case 128:
                                                                                switch v45 {
                                                                                case "foo bar":
                                                                                    if (0 > uint16(v19)) {
                                                                                        v49 := GenS5{F0: 32767, F1: v19}
                                                                                        var v50 string = v45
                                                                                        fmt.Println(v4)
                                                                                        if false {
                                                                                            v51 := GenS6{F0: (^(v4)), F1: 112}
                                                                                            fmt.Println(v2, fi15, v19)
                                                                                            _ = len(v1)
                                                                                            _ = v51
                                                                                        }
                                                                                        _ = v49
                                                                                        _ = v50
                                                                                    } else {
                                                                                        switch 4294967295 {
                                                                                        case 127:
                                                                                            fmt.Println(v44, fi18)
                                                                                            v52 := GenS2{F0: "a"}
                                                                                            fmt.Println(v19, v3, fi18)
                                                                                            _ = v52
                                                                                        case 2147483647:
                                                                                            v19 = (v19 | 1)
                                                                                        case -809:
                                                                                            v19 -= (^(v19))
                                                                                        }
                                                                                    }
                                                                                case "hello":
                                                                                    v53 := v1["a"]
                                                                                    _ = v53
                                                                                }
                                                                            default:
                                                                                switch (fi18 >> ((685) & 63)) {
                                                                                case 128:
                                                                                    v3 -= v3
                                                                                case -715:
                                                                                    var v54 uint64 = v19
                                                                                    _ = v54
                                                                                case -9223372036854775808:
                                                                                    var v55 int = (fi18 ^ fi15)
                                                                                    _ = v55
                                                                                }
                                                                            }
                                                                            if false {
                                                                                v19 += (^(v19))
                                                                            } else {
                                                                                v56 := GenS7{F0: v4, F1: (v3 << ((2) & 63))}
                                                                                delete(v1, "foo bar")
                                                                                v4 = v4
                                                                                _ = v56
                                                                            }
                                                                            if (false && true) {
                                                                                var v57 int = fi18
                                                                                _ = v57
                                                                            } else {
                                                                                if ((fi15 << ((470) & 63)) <= (fi18 + (fi15 | (fi18 >> ((720) & 63))))) {
                                                                                    if false {
                                                                                        if (-576 == (v44 ^ (-(v44)))) {
                                                                                            v2 -= v2
                                                                                            var v58 float64 = ((3.14159 * 0.1) / -2.25)
                                                                                            _ = v58
                                                                                        }
                                                                                        v3 -= -1
                                                                                    }
                                                                                    v59 := GenS6{F0: 717, F1: 120}
                                                                                    var v60 uint8 = 131
                                                                                    _ = v59
                                                                                    _ = v60
                                                                                } else {
                                                                                    if (!false) {
                                                                                        v45 = v45
                                                                                    } else {
                                                                                        v2 = 1.5
                                                                                    }
                                                                                }
                                                                                if true {
                                                                                    if (0 >= 814) {
                                                                                        _ = len(v1)
                                                                                    }
                                                                                    if ((^(v4)) != (v4 * v4)) {
                                                                                        var v61 uint64 = (v19 | (v19 >> ((347) & 63)))
                                                                                        v62 := v1[v45]
                                                                                        fmt.Println(fi15, v2)
                                                                                        _ = v61
                                                                                        _ = v62
                                                                                    } else {
                                                                                        if (!(!true)) {
                                                                                            v44 += (v44 >> ((2) & 63))
                                                                                            if true {
                                                                                                v1[((v45 + "hello") + "")] = true
                                                                                            }
                                                                                            if len(v9) > 0 {
                                                                                                v9[0] = fi15
                                                                                            }
                                                                                        } else {
                                                                                            var v63 uint = 2
                                                                                            switch v45 {
                                                                                            case "":
                                                                                                switch (v45 + ("hello" + ("hello" + "foo bar"))) {
                                                                                                case "":
                                                                                                    if (1000 > 42) {
                                                                                                        fmt.Println(fi18, v44)
                                                                                                        var v64 uint32 = v4
                                                                                                        var v65 float64 = -2.25
                                                                                                        var v66 uint16 = 1
                                                                                                        if (v63 >= (v63 + (v63 % (((v63 | uint(v19))) | 1)))) {
                                                                                                            if len(v9) >= 2 {
                                                                                                                v9 = v9[1:len(v9)]
                                                                                                            }
                                                                                                            v67 := v1[v45]
                                                                                                            _ = v67
                                                                                                        }
                                                                                                        _ = v64
                                                                                                        _ = v65
                                                                                                        _ = v66
                                                                                                    } else {
                                                                                                        if len(v9) >= 2 {
                                                                                                            v9 = v9[1:len(v9)]
                                                                                                        }
                                                                                                        v19 = (^(v19))
                                                                                                    }
                                                                                                case "hello":
                                                                                                    v68 := v1[v45]
                                                                                                    _ = v68
                                                                                                case "\n\t":
                                                                                                    if len(v9) >= 2 {
                                                                                                        v9 = v9[1:len(v9)]
                                                                                                    }
                                                                                                }
                                                                                            case "a":
                                                                                                v1["a"] = (-1 < -1)
                                                                                            }
                                                                                            _ = v63
                                                                                        }
                                                                                    }
                                                                                    v3 -= v3
                                                                                }
                                                                            }
                                                                            fmt.Println(fi18)
                                                                            fmt.Println(v3, v45, v44)
                                                                            var v69 int = 686
                                                                            _ = v46
                                                                            _ = v69
                                                                        } else {
                                                                            v45 = "a"
                                                                        }
                                                                    }
                                                                    fmt.Println(fi15)
                                                                    if false {
                                                                        v2 = (float32((1e10 * (3.14159 - 0.1))) * float32((1.5 + 0.0)))
                                                                        var v70 float32 = (v2 + ((-2.25 - float32(0.1)) / v2))
                                                                        _ = v70
                                                                    } else {
                                                                        _ = len(v9)
                                                                        v71 := GenS8{F0: 0, F1: 32768}
                                                                        _ = len(v9)
                                                                        _ = v71
                                                                    }
                                                                    var v72 bool = true
                                                                    _ = v72
                                                                } else {
                                                                    switch v45 {
                                                                    case "":
                                                                        v9 = append(v9, 4294967295)
                                                                    case "a":
                                                                        v45 = (v45 + "foo bar")
                                                                    }
                                                                }
                                                                v73 := GenS3{F0: v4, F1: 1000}
                                                                _ = v45
                                                                _ = v73
                                                            } else {
                                                                v74 := GenS3{F0: (v4 * (^(v4))), F1: 136}
                                                                _ = v74
                                                            }
                                                            switch "" {
                                                            case "":
                                                                v44 = v44
                                                            default:
                                                                var v75 int32 = v44
                                                                _ = v75
                                                            }
                                                        }
                                                        v76 := GenS4{F0: (v44 / ((v44) | 1)), F1: -128, F2: ""}
                                                        _ = v44
                                                        _ = v76
                                                    }
                                                default:
                                                    v4 = v4
                                                }
                                            }
                                            if (-32768 >= 0) {
                                                v2 += (v2 + 0.0)
                                            }
                                            switch "" {
                                            case "a":
                                                v19 = 4294967295
                                            case "foo bar":
                                                delete(v1, "")
                                            case "":
                                                var v77 uint8 = 128
                                                _ = v77
                                            default:
                                                if true {
                                                    v1[("foo bar" + ("a" + "foo bar"))] = (!true)
                                                    fmt.Println(fi18)
                                                }
                                            }
                                            fmt.Println(fi15, fi18, v2)
                                        } else {
                                            if true {
                                                v3 = v3
                                                v3 += (^(v3))
                                            }
                                        }
                                    case "foo bar":
                                        var v78 int = (-(fi15))
                                        _ = v78
                                    case "hello":
                                        if false {
                                            v3 = 42
                                            var v79 int8 = -1
                                            switch fi18 {
                                            case 1:
                                                if (!false) {
                                                    var v80 int64 = -9223372036854775808
                                                    _ = v80
                                                } else {
                                                    v81 := GenS6{F0: uint32(v19), F1: 128}
                                                    _ = v81
                                                }
                                            case 127:
                                                var v82 uint32 = v4
                                                _ = v82
                                            }
                                            _ = v79
                                        }
                                    }
                                case 490:
                                    fmt.Println(v2, v3)
                                }
                                if (914 <= -980) {
                                    fmt.Println(v4, v2, fi15)
                                    fmt.Println(v2, fi18)
                                    v19 += (^(v19))
                                } else {
                                    v19 += (v19 >> ((uint(v19)) & 63))
                                    var v83 int32 = -2147483648
                                    _ = v83
                                }
                                _ = v33
                            }
                        default:
                            if len(v9) > 0 {
                                v9[0] = 4294967295
                            }
                            fmt.Println(v4, v2)
                        }
                        if false {
                            delete(v1, "hello")
                            if false {
                                switch 127 {
                                case -680:
                                    switch (("hello" + ("a" + "a")) + "foo bar") {
                                    case "a":
                                        var v84 float64 = 0.0
                                        _ = v84
                                    case "\n\t":
                                        if ((fi15 % ((fi18) | 1)) <= fi18) {
                                            var v85 float64 = (((0.1 * 0.1) + 1.5) * -2.25)
                                            v86 := GenS1{F0: fi15, F1: 32767, F2: 2}
                                            _ = v85
                                            _ = v86
                                        }
                                    case "":
                                        if true {
                                            if len(v9) > 0 {
                                                v9[0] = -770
                                            }
                                        }
                                    default:
                                        var v87 string = "a"
                                        _ = v87
                                    }
                                }
                                var v88 string = ((("hello" + ("a" + ("" + ("\n\t" + "\n\t")))) + "foo bar") + "")
                                if len(v9) >= 2 {
                                    v9 = v9[1:len(v9)]
                                }
                                _ = v88
                            }
                            var v89 float64 = 0.0
                            fmt.Println(v3, fi15)
                            switch (fi18 % (((fi18 | (^(fi18)))) | 1)) {
                            case 9223372036854775807:
                                switch "" {
                                case "\n\t":
                                    fmt.Println(fi18, v4, v2)
                                case "hello":
                                    v1["a"] = true
                                }
                            case 32767:
                                v4 -= v4
                            case 127:
                                if true {
                                    if false {
                                        v90 := v1[""]
                                        _ = len(v9)
                                        _ = v90
                                    } else {
                                        switch fi18 {
                                        case -2147483648:
                                            if len(v9) > 0 {
                                                v9[0] = (fi18 / (((-(fi18))) | 1))
                                            }
                                            if len(v9) >= 2 {
                                                v9 = v9[1:len(v9)]
                                            }
                                            fmt.Println(fi15, fi15)
                                        case 1:
                                            var v91 uint32 = (^(v4))
                                            _ = v91
                                        default:
                                            v92 := GenS6{F0: v4, F1: 0}
                                            _ = v92
                                        }
                                        v1[((((("foo bar" + "\n\t") + "") + "foo bar") + ("" + "\n\t")) + "foo bar")] = false
                                    }
                                    var v93 int32 = 700
                                    _ = v93
                                } else {
                                    var v94 int8 = -1
                                    _ = v94
                                }
                            default:
                                switch fi15 {
                                case 897:
                                    if (2147483647 < -337) {
                                        if false {
                                            fmt.Println(v89, fi18)
                                            fmt.Println(fi18)
                                        }
                                        if false {
                                            switch "foo bar" {
                                            case "foo bar":
                                                v4 = v4
                                            case "":
                                                v89 = ((v89 - v89) - v89)
                                            default:
                                                var v95 uint = 0
                                                _ = v95
                                            }
                                            var v96 bool = false
                                            switch (fi15 % (((fi18 - (-(fi15)))) | 1)) {
                                            case 59:
                                                if v96 {
                                                    var v97 uint8 = genFn1(4294967295, 2, v19)
                                                    var v98 uint8 = (v97 | 0)
                                                    _ = v97
                                                    _ = v98
                                                } else {
                                                    var v99 int8 = 0
                                                    v99 = -1
                                                    fmt.Println(v19, v99)
                                                    _ = len(v1)
                                                    _ = v99
                                                }
                                            case -553:
                                                v89 = v89
                                            case 1:
                                                switch "\n\t" {
                                                case "a":
                                                    v9 = append(v9, fi15)
                                                case "":
                                                    v100 := GenS9{F0: 0}
                                                    _ = v100
                                                default:
                                                    v3 += v3
                                                }
                                            default:
                                                v101 := GenS10{F0: 1000, F1: 127}
                                                _ = v101
                                            }
                                            _ = v96
                                        } else {
                                            v89 = v89
                                            if (false && true) {
                                                v19 -= 18446744073709551615
                                                switch "hello" {
                                                case "":
                                                    v102 := GenS11{F0: -1}
                                                    _ = v102
                                                }
                                            } else {
                                                v89 -= float64(1.5)
                                                v103 := GenS12{F0: 86}
                                                fmt.Println(v19)
                                                v104 := GenS7{F0: (v4 << ((18446744073709551615) & 63)), F1: v3}
                                                v19 -= (v19 / (((^(v19))) | 1))
                                                _ = v103
                                                _ = v104
                                            }
                                            switch "\n\t" {
                                            case "hello":
                                                if false {
                                                    v105 := GenS13{F0: 369, F1: 1}
                                                    var v106 string = "foo bar"
                                                    var v107 int = fi18
                                                    _ = v105
                                                    _ = v106
                                                    _ = v107
                                                } else {
                                                    if false {
                                                        v108 := v1["hello"]
                                                        switch "hello" {
                                                        case "hello":
                                                            if (("foo bar" + ("foo bar" + ("foo bar" + ""))) <= "\n\t") {
                                                                if true {
                                                                    v109 := GenS14{F0: 255}
                                                                    if len(v9) >= 2 {
                                                                        v9 = v9[1:len(v9)]
                                                                    }
                                                                    _ = v109
                                                                } else {
                                                                    if (127 >= 122) {
                                                                        v9 = append(v9, (fi18 - (fi18 % (((^(fi15))) | 1))))
                                                                        v110 := GenS7{F0: (v4 * v4), F1: 0}
                                                                        var v111 uint16 = 453
                                                                        v3 = int64(fi18)
                                                                        fmt.Println(v2)
                                                                        _ = v110
                                                                        _ = v111
                                                                    } else {
                                                                        var v112 int16 = -32768
                                                                        _ = v112
                                                                    }
                                                                    v19 = (v19 & v19)
                                                                    if false {
                                                                        var v113 uint = 92
                                                                        v3 = (v3 | 9223372036854775807)
                                                                        if v108 {
                                                                            var v114 uint64 = (^(v19))
                                                                            delete(v1, "a")
                                                                            if true {
                                                                                if (88 > 0) {
                                                                                    switch 9223372036854775807 {
                                                                                    case 32767:
                                                                                        v115 := GenS9{F0: 42}
                                                                                        _ = v115
                                                                                    default:
                                                                                        switch fi15 {
                                                                                        case 2147483647:
                                                                                            v116 := v1[("foo bar" + "\n\t")]
                                                                                            _ = v116
                                                                                        case -80:
                                                                                            v108 = v108
                                                                                        default:
                                                                                            fmt.Println(fi18, v113)
                                                                                            v117 := GenS15{F0: v4}
                                                                                            _ = v117
                                                                                        }
                                                                                    }
                                                                                    v118 := GenS16{F0: (v113 - v113), F1: int64(v114), F2: uint64(v4)}
                                                                                    v119 := GenS17{F0: (v3 >> ((v113) & 63)), F1: "a"}
                                                                                    _ = v118
                                                                                    _ = v119
                                                                                } else {
                                                                                    switch (fi18 ^ -4) {
                                                                                    case 42:
                                                                                        fmt.Println(v2, v108)
                                                                                    case 2147483647:
                                                                                        var v120 uint32 = (v4 - (v4 & v4))
                                                                                        _ = v120
                                                                                    case -2147483648:
                                                                                        switch ("\n\t" + "\n\t") {
                                                                                        case "\n\t":
                                                                                            v3 -= 1
                                                                                        case "foo bar":
                                                                                            v3 -= v3
                                                                                            var v121 int64 = 150
                                                                                            _ = v121
                                                                                        }
                                                                                    default:
                                                                                        fmt.Println(v89, v3, v19)
                                                                                    }
                                                                                    var v122 uint32 = (^(v4))
                                                                                    if (127 <= 0) {
                                                                                        v114 -= v19
                                                                                        if (!v108) {
                                                                                            var v123 float64 = 0.0
                                                                                            var v124 int8 = 0
                                                                                            var v125 int16 = 559
                                                                                            _ = v123
                                                                                            _ = v124
                                                                                            _ = v125
                                                                                        } else {
                                                                                            v126 := v1["hello"]
                                                                                            v127 := GenS17{F0: v3, F1: "hello"}
                                                                                            _ = v126
                                                                                            _ = v127
                                                                                        }
                                                                                    }
                                                                                    _ = v122
                                                                                }
                                                                                v113 = v113
                                                                                var v128 int16 = -525
                                                                                _ = v128
                                                                            }
                                                                            _ = v114
                                                                        }
                                                                        _ = v113
                                                                    } else {
                                                                        _ = len(v1)
                                                                    }
                                                                    var v129 int8 = 42
                                                                    _ = v129
                                                                }
                                                            } else {
                                                                var v130 uint64 = uint64(v4)
                                                                _ = v130
                                                            }
                                                        case "a":
                                                            fmt.Println(v2, v2, v19)
                                                        case "foo bar":
                                                            if len(v9) >= 2 {
                                                                v9 = v9[1:len(v9)]
                                                            }
                                                        default:
                                                            var v131 uint8 = 41
                                                            _ = v131
                                                        }
                                                        v108 = (!v108)
                                                        _ = v108
                                                    } else {
                                                        v4 += v4
                                                        var v132 uint16 = 664
                                                        _ = v132
                                                    }
                                                }
                                            case "":
                                                v89 = ((v89 * ((v89 * v89) / (v89 / -2.25))) / -2.25)
                                            default:
                                                fmt.Println(v19)
                                            }
                                            switch "a" {
                                            case "hello":
                                                var v133 int8 = 1
                                                _ = v133
                                            case "foo bar":
                                                v4 -= v4
                                            case "a":
                                                v134 := GenS5{F0: uint16(v4), F1: (v19 + v19)}
                                                _ = v134
                                            default:
                                                if true {
                                                    var v135 float64 = (-2.25 - 0.1)
                                                    var v136 int64 = v3
                                                    _ = v135
                                                    _ = v136
                                                }
                                            }
                                            v137 := GenS6{F0: (^(v4)), F1: 2}
                                            _ = v137
                                        }
                                    } else {
                                        v19 = v19
                                        v138 := GenS5{F0: 42, F1: (v19 >> ((0) & 63))}
                                        v139 := GenS16{F0: 951, F1: (-(v3)), F2: uint64(v4)}
                                        _ = v138
                                        _ = v139
                                    }
                                    if (((0 > v4) && false) && (1 != uint8(v4))) {
                                        fmt.Println(v19, fi15)
                                        if len(v9) > 0 {
                                            v9[0] = -1
                                        }
                                    } else {
                                        v3 -= v3
                                        if false {
                                            var v140 float64 = v89
                                            var v141 bool = false
                                            _ = v140
                                            _ = v141
                                        } else {
                                            if true {
                                                var v142 uint64 = v19
                                                var v143 string = "\n\t"
                                                _ = v142
                                                _ = v143
                                            } else {
                                                var v144 uint = 18446744073709551615
                                                if (32767 == -61) {
                                                    if (644 > 32767) {
                                                        var v145 int64 = (-(v3))
                                                        v144 += (v144 - v144)
                                                        if (249 < -487) {
                                                            v4 -= (v4 << ((0) & 63))
                                                        }
                                                        if (!false) {
                                                            if false {
                                                                if ((v2 - ((float32((v89 * v89)) + v2) * 0.1)) > (v2 + (0.1 / v2))) {
                                                                    if false {
                                                                        if len(v9) > 0 {
                                                                            v9[0] = fi18
                                                                        }
                                                                        var v146 uint8 = 2
                                                                        switch "\n\t" {
                                                                        case "":
                                                                            v145 = v3
                                                                        case "a":
                                                                            fmt.Println(v144, v144)
                                                                        case "foo bar":
                                                                            if false {
                                                                                var v147 uint8 = v146
                                                                                var v148 string = (("" + "hello") + "")
                                                                                var v149 int64 = (v145 & v145)
                                                                                v150 := GenS13{F0: 75, F1: genFn1(v144, (v146 | genFn1((v144 >> ((975) & 63)), 2, (v19 % ((985) | 1)))), (v19 + 2))}
                                                                                fmt.Println(v4, v146, fi15)
                                                                                _ = v147
                                                                                _ = v148
                                                                                _ = v149
                                                                                _ = v150
                                                                            }
                                                                        }
                                                                        _ = v146
                                                                    } else {
                                                                        var v151 string = "\n\t"
                                                                        if false {
                                                                            switch fi18 {
                                                                            case 32767:
                                                                                var v152 int16 = 1000
                                                                                _ = v152
                                                                            default:
                                                                                v153 := GenS18{F0: 0, F1: 0, F2: (false || (0 == 100))}
                                                                                _ = v153
                                                                            }
                                                                            var v154 uint8 = 42
                                                                            if (0 != 0) {
                                                                                var v155 uint64 = 2
                                                                                v156 := GenS8{F0: 42, F1: (fi18 / ((128) | 1))}
                                                                                _ = v155
                                                                                _ = v156
                                                                            } else {
                                                                                v89 = 3.14159
                                                                                v151 = v151
                                                                                if true {
                                                                                    if false {
                                                                                        v157 := GenS19{F0: (v151 + v151)}
                                                                                        if ((v145 % ((v3) | 1)) < (v145 / ((v145) | 1))) {
                                                                                            fmt.Println(v151, v89)
                                                                                            v151 = ("a" + "hello")
                                                                                            if (true || true) {
                                                                                                v19 -= (v19 * v19)
                                                                                            }
                                                                                        } else {
                                                                                            if (false && true) {
                                                                                                var v158 uint8 = v154
                                                                                                var v159 int8 = 127
                                                                                                _ = v158
                                                                                                _ = v159
                                                                                            } else {
                                                                                                var v160 bool = ((v2 > ((0.0 + (float32(v89) - 0.0)) - v2)) || ((fi18 & fi18) > -827))
                                                                                                _ = v160
                                                                                            }
                                                                                        }
                                                                                        fmt.Println(fi15)
                                                                                        _ = v157
                                                                                    } else {
                                                                                        v9 = append(v9, (fi18 >> ((v144) & 63)))
                                                                                        var v161 string = "foo bar"
                                                                                        v4 = (v4 ^ (^(v4)))
                                                                                        _ = v161
                                                                                    }
                                                                                    v4 += v4
                                                                                }
                                                                            }
                                                                            switch (fi15 >> ((1) & 63)) {
                                                                            case 0:
                                                                                if false {
                                                                                    v9 = append(v9, int(v3))
                                                                                    var v162 float32 = 0.1
                                                                                    var v163 int16 = int16(v154)
                                                                                    _ = v162
                                                                                    _ = v163
                                                                                }
                                                                            }
                                                                            fmt.Println(v154)
                                                                            if (!(!(!((-(fi15)) <= (-(fi18)))))) {
                                                                                v164 := GenS14{F0: 0}
                                                                                var v165 uint16 = 2
                                                                                _ = v164
                                                                                _ = v165
                                                                            } else {
                                                                                var v166 uint16 = 255
                                                                                _ = v166
                                                                            }
                                                                            _ = v154
                                                                        } else {
                                                                            if len(v9) >= 2 {
                                                                                v9 = v9[1:len(v9)]
                                                                            }
                                                                        }
                                                                        _ = v151
                                                                    }
                                                                } else {
                                                                    var v167 int16 = 42
                                                                    _ = v167
                                                                }
                                                                var v168 uint8 = uint8(v19)
                                                                v169 := GenS20{F0: (true || (v4 >= v4)), F1: 810}
                                                                _ = v168
                                                                _ = v169
                                                            } else {
                                                                v170 := GenS21{F0: (v144 & (v144 | v144))}
                                                                switch fi18 {
                                                                case 0:
                                                                    v89 -= float64(0.0)
                                                                }
                                                                _ = v170
                                                            }
                                                            fmt.Println(v2, v3)
                                                            v171 := GenS22{F0: genFn1(662, 255, (v19 << ((v144) & 63))), F1: 42, F2: v19}
                                                            _ = v171
                                                        }
                                                        v89 = v89
                                                        if (true || true) {
                                                            v4 = 2147483648
                                                            switch "a" {
                                                            case "foo bar":
                                                                v172 := GenS23{F0: "hello", F1: (v19 & (v19 & v19)), F2: false}
                                                                var v173 float64 = 3.14159
                                                                _ = v172
                                                                _ = v173
                                                            case "":
                                                                v19 -= v19
                                                            case "\n\t":
                                                                var v174 int32 = -1
                                                                _ = v174
                                                            default:
                                                                v175 := GenS24{F0: (true && true), F1: 32767}
                                                                _ = v175
                                                            }
                                                            if true {
                                                                if true {
                                                                    if true {
                                                                        v9 = append(v9, (fi18 / ((-175) | 1)))
                                                                        var v176 int8 = 42
                                                                        if (fi15 != (fi18 + (fi15 + 1))) {
                                                                            fmt.Println(fi15, v144, v89)
                                                                        }
                                                                        _ = v176
                                                                    }
                                                                } else {
                                                                    v177 := v1[("foo bar" + "hello")]
                                                                    var v178 string = "\n\t"
                                                                    v179 := GenS20{F0: v177, F1: v19}
                                                                    v144 += 65535
                                                                    if v177 {
                                                                        v180 := GenS1{F0: 1, F1: 2, F2: 152}
                                                                        _ = v180
                                                                    } else {
                                                                        delete(v1, v178)
                                                                        if v177 {
                                                                            v177 = (v178 >= "foo bar")
                                                                            var v181 uint64 = v19
                                                                            _ = v181
                                                                        }
                                                                    }
                                                                    _ = v177
                                                                    _ = v178
                                                                    _ = v179
                                                                }
                                                            }
                                                            if (42 >= v3) {
                                                                var v182 uint32 = 412
                                                                if true {
                                                                    var v183 bool = true
                                                                    switch (-(fi15)) {
                                                                    case -917:
                                                                        var v184 float32 = (3.14159 - (v2 / ((0.1 / (v2 / v2)) - 1.5)))
                                                                        _ = v184
                                                                    }
                                                                    _ = v183
                                                                } else {
                                                                    v144 = v144
                                                                }
                                                                v185 := GenS25{F0: 472, F1: 0, F2: (v4 - (v182 & (v182 - (^(v4)))))}
                                                                _ = v182
                                                                _ = v185
                                                            } else {
                                                                fmt.Println(fi15, fi15)
                                                                if ((v144 % (((^(v144))) | 1)) >= 1) {
                                                                    switch fi18 {
                                                                    case -340:
                                                                        fmt.Println(v89, v144, fi15)
                                                                    case 32767:
                                                                        v3 -= v3
                                                                    case 0:
                                                                        v3 -= (^(v145))
                                                                    default:
                                                                        var v186 uint64 = 18446744073709551615
                                                                        _ = v186
                                                                    }
                                                                    switch ("" + ("foo bar" + "\n\t")) {
                                                                    case "foo bar":
                                                                        v187 := GenS19{F0: "foo bar"}
                                                                        _ = v187
                                                                    case "":
                                                                        var v188 uint8 = 127
                                                                        var v189 uint8 = (v188 | v188)
                                                                        _ = v188
                                                                        _ = v189
                                                                    default:
                                                                        v190 := v1["foo bar"]
                                                                        _ = v190
                                                                    }
                                                                    switch (fi15 << ((v144) & 63)) {
                                                                    case 977:
                                                                        v2 += (3.14159 - v2)
                                                                    case -2147483648:
                                                                        v191 := GenS26{F0: 42, F1: (v144 - uint(v19)), F2: 391}
                                                                        _ = v191
                                                                    case 127:
                                                                        _ = len(v9)
                                                                    default:
                                                                        var v192 int8 = -1
                                                                        _ = v192
                                                                    }
                                                                }
                                                                v1[""] = (((-2.25 == v2) || (true && ((((0.0 - 1e10) - -2.25) * (v89 + (float64(v2) - v89))) <= 1.5))) || (-206 == 1))
                                                            }
                                                            v9 = append(v9, fi18)
                                                        }
                                                        var v193 float64 = float64((v2 / 0.1))
                                                        _ = v145
                                                        _ = v193
                                                    } else {
                                                        var v194 uint32 = uint32(v19)
                                                        v195 := v1["\n\t"]
                                                        _ = v194
                                                        _ = v195
                                                    }
                                                    if len(v9) > 0 {
                                                        v9[0] = 32768
                                                    }
                                                    var v196 int32 = -2147483648
                                                    _ = v196
                                                } else {
                                                    var v197 uint8 = 2
                                                    if false {
                                                        var v198 uint64 = (v19 ^ v19)
                                                        var v199 uint = v144
                                                        _ = v198
                                                        _ = v199
                                                    }
                                                    switch fi15 {
                                                    case 578:
                                                        if len(v9) >= 2 {
                                                            v9 = v9[1:len(v9)]
                                                        }
                                                    case -9223372036854775808:
                                                        var v200 float32 = -2.25
                                                        if ((false && false) || (false || false)) {
                                                            v197 -= 128
                                                            v200 -= v200
                                                            v89 -= float64((1.5 - 0.1))
                                                        } else {
                                                            fmt.Println(v4, v2)
                                                            v201 := GenS6{F0: (v4 ^ v4), F1: 42}
                                                            v19 = 1
                                                            _ = v201
                                                        }
                                                        _ = v200
                                                    default:
                                                        v202 := GenS13{F0: -776, F1: (v197 / (((v197 * uint8(v19))) | 1))}
                                                        _ = v202
                                                    }
                                                    v1["\n\t"] = false
                                                    _ = v197
                                                }
                                                _ = len(v9)
                                                _ = v144
                                            }
                                        }
                                    }
                                case 9223372036854775807:
                                    if true {
                                        var v203 uint8 = 18
                                        _ = v203
                                    } else {
                                        v204 := GenS15{F0: v4}
                                        fmt.Println(fi15)
                                        fmt.Println(v3, v3)
                                        _ = len(v1)
                                        _ = v204
                                    }
                                case 6:
                                    v205 := GenS27{F0: "", F1: v4, F2: 255}
                                    _ = v205
                                default:
                                    v4 = (^(v4))
                                }
                            }
                            _ = v89
                        }
                        _ = v19
                    } else {
                        v3 -= (v3 & v3)
                        if (false && (42 > 1)) {
                            if ((((float64((v2 / v2)) != float64(float32((0.0 / (1e10 - 3.14159))))) && (1 != -1)) || (2 != 179)) || (839 == fi18)) {
                                v206 := GenS26{F0: 0, F1: 0, F2: 2147483647}
                                var v207 uint32 = (^(v4))
                                if (42 >= 1) {
                                    var v208 float32 = v2
                                    if (false && false) {
                                        var v209 uint64 = 114
                                        v210 := GenS28{F0: "\n\t", F1: 2}
                                        _ = v209
                                        _ = v210
                                    } else {
                                        if false {
                                            switch "foo bar" {
                                            case "a":
                                                delete(v1, "foo bar")
                                            }
                                            switch "" {
                                            case "a":
                                                var v211 float32 = ((((v2 / v208) - (v2 + float32(0.0))) * 0.0) / v2)
                                                _ = v211
                                            case "\n\t":
                                                switch ("" + "") {
                                                case "hello":
                                                    v212 := GenS18{F0: 0, F1: -890, F2: false}
                                                    _ = v212
                                                case "":
                                                    switch (fi15 | fi18) {
                                                    case 32768:
                                                        v2 -= 1.5
                                                    case 127:
                                                        if ("\n\t" < ("foo bar" + "\n\t")) {
                                                            fmt.Println(fi18, v207, v208)
                                                            switch -9223372036854775808 {
                                                            case 2147483647:
                                                                v4 = 42
                                                            case -2147483648:
                                                                if false {
                                                                    delete(v1, (("hello" + "") + ""))
                                                                }
                                                            }
                                                            if len(v9) > 0 {
                                                                v9[0] = (fi15 << ((65535) & 63))
                                                            }
                                                        } else {
                                                            if true {
                                                                v3 -= (v3 + (v3 * v3))
                                                                var v213 int = -466
                                                                var v214 float32 = ((1e10 + (((1.5 / -2.25) * (0.0 * 0.1)) / 0.1)) + float32((3.14159 + 1.5)))
                                                                _ = v213
                                                                _ = v214
                                                            }
                                                            var v215 int16 = 1000
                                                            v208 -= v208
                                                            _ = v215
                                                        }
                                                    }
                                                }
                                            }
                                        } else {
                                            if ((0.1 * (1e10 * (0.0 / (1e10 - 3.14159)))) != (1.5 - 3.14159)) {
                                                _ = len(v1)
                                                v216 := GenS14{F0: 42}
                                                v3 = v3
                                                v2 = 1.5
                                                _ = v216
                                            } else {
                                                v2 -= (3.14159 / float32(1.5))
                                            }
                                            fmt.Println(v2, v207)
                                            var v217 uint16 = 255
                                            _ = v217
                                        }
                                        var v218 uint = 1
                                        _ = v218
                                    }
                                    _ = v208
                                }
                                _ = v206
                                _ = v207
                            }
                        }
                        switch ((((("a" + "\n\t") + "a") + "") + ("a" + ("a" + (("\n\t" + ("foo bar" + "")) + "\n\t")))) + "foo bar") {
                        case "\n\t":
                            if (696 == 4294967295) {
                                fmt.Println(v3)
                                v4 = (v4 - (v4 / ((v4) | 1)))
                            } else {
                                var v219 float32 = 3.14159
                                if false {
                                    v220 := GenS25{F0: 65535, F1: 157, F2: v4}
                                    if (true || true) {
                                        if (42 > 128) {
                                            delete(v1, "\n\t")
                                        } else {
                                            v3 += (v3 | (v3 * (v3 | v3)))
                                        }
                                    } else {
                                        v4 = (^(v4))
                                        switch (("a" + "foo bar") + (("\n\t" + "hello") + "hello")) {
                                        case "":
                                            v221 := GenS29{F0: -1}
                                            _ = v221
                                        case "\n\t":
                                            switch fi15 {
                                            case 0:
                                                delete(v1, "")
                                            case -284:
                                                var v222 bool = (42 != (v4 ^ v4))
                                                _ = v222
                                            }
                                            v3 -= v3
                                            delete(v1, "")
                                        default:
                                            v3 -= v3
                                        }
                                        v223 := GenS21{F0: 0}
                                        _ = v223
                                    }
                                    v3 = v3
                                    _ = v220
                                } else {
                                    v3 -= v3
                                }
                                _ = v219
                            }
                        case "":
                            v3 = 9223372036854775807
                        case "foo bar":
                            v1["\n\t"] = false
                        }
                    }
                default:
                    v224 := GenS29{F0: -2147483648}
                    _ = v224
                }
                var v225 uint16 = 42
                v2 += 1e10
                _ = fi18
                _ = v225
            }
        }
        for fi226 := 0; fi226 < 1; fi226++ {
            v2 += (v2 * 3.14159)
            var v227 uint8 = genFn1(42, 107, 18446744073709551615)
            v2 = float32(1e10)
            _ = fi226
            _ = v227
        }
        v2 = float32(float64(v2))
        _ = fi15
    }
    var v228 int8 = 127
    if false {
        var v229 int16 = 0
        for fi230 := 0; fi230 < 5; fi230++ {
            var v231 string = "\n\t"
            fmt.Println(fi230)
            v232 := GenS30{F0: 631}
            _ = fi230
            _ = v231
            _ = v232
        }
        _ = v229
    } else {
        if ("" >= "foo bar") {
            var v233 uint8 = 172
            _ = v233
        } else {
            var v234 int = 548
            if false {
                v2 += v2
                if true {
                    for fi235 := 0; fi235 < 18; fi235++ {
                        if true {
                            if true {
                                v236 := GenS29{F0: 2147483647}
                                _ = v236
                            }
                            var v237 uint16 = 255
                            fmt.Println(v237, v3)
                            _ = v237
                        } else {
                            fmt.Println(v228, v2, v234)
                            var v238 uint = 196
                            _ = v238
                        }
                        v239 := v1["a"]
                        var v240 int16 = 1
                        _ = fi235
                        _ = v239
                        _ = v240
                    }
                } else {
                    var v241 bool = false
                    var v242 uint = 255
                    _ = v241
                    _ = v242
                }
            } else {
                var v243 int = 127
                var v244 float32 = 1.5
                _ = v243
                _ = v244
            }
            var v245 bool = false
            _ = v234
            _ = v245
        }
        v2 -= 3.14159
    }
    for fi246 := 0; fi246 < 10; fi246++ {
        v247 := v1["foo bar"]
        switch 127 {
        case -2147483648:
            v4 -= v4
        default:
            v4 -= v4
        }
        var v248 int8 = v228
        _ = fi246
        _ = v247
        _ = v248
    }
    for fi249 := 0; fi249 < 19; fi249++ {
        for fi250 := 0; fi250 < 2; fi250++ {
            v4 = v4
            var v251 float32 = 0.1
            v252 := GenS22{F0: uint8(v4), F1: (v3 << ((1) & 63)), F2: 578}
            _ = fi250
            _ = v251
            _ = v252
        }
        var v253 int = (fi249 << ((209) & 63))
        _ = fi249
        _ = v253
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v228:", v228)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v228
    _ = v3
    _ = v4
    _ = v9
}
