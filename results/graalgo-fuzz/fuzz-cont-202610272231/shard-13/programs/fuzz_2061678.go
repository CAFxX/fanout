package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool; F2 uint32 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int16; F1 int64; F2 uint64 }
type GenS4 struct { F0 bool; F1 uint64 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 string; F1 int16 }
type GenS7 struct { F0 int; F1 uint32 }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 uint16; F1 string; F2 string }
type GenS10 struct { F0 uint32; F1 uint64 }
type GenS11 struct { F0 int8; F1 uint16; F2 uint32 }
type GenS12 struct { F0 uint64 }
type GenS13 struct { F0 uint }
type GenS14 struct { F0 int }
type GenS15 struct { F0 uint64 }
type GenS16 struct { F0 int32 }
type GenS17 struct { F0 bool; F1 int16; F2 bool }
type GenS18 struct { F0 bool }
type GenS19 struct { F0 bool; F1 bool }
type GenS20 struct { F0 uint8; F1 int8; F2 uint8 }
type GenS21 struct { F0 uint8 }
type GenS22 struct { F0 bool }
type GenS23 struct { F0 uint16 }
type GenS24 struct { F0 uint32 }
type GenS25 struct { F0 uint64; F1 int; F2 int8 }
type GenS26 struct { F0 int8 }
type GenS27 struct { F0 int; F1 uint8; F2 int16 }
type GenS28 struct { F0 uint32 }
type GenS29 struct { F0 int8 }
type GenS30 struct { F0 uint16; F1 int16 }
type GenS31 struct { F0 uint8 }
type GenS32 struct { F0 uint8; F1 uint }
type GenS33 struct { F0 bool; F1 int }
type GenS34 struct { F0 uint8 }
type GenS35 struct { F0 uint32 }
type GenS36 struct { F0 int; F1 int32; F2 string }
type GenS37 struct { F0 int16; F1 int8 }
type GenS38 struct { F0 int }
type GenS39 struct { F0 uint; F1 bool; F2 int8 }
type GenS40 struct { F0 uint8; F1 uint16; F2 uint32 }
type GenS41 struct { F0 uint64; F1 int8 }
type GenS42 struct { F0 string; F1 uint; F2 int }
type GenS43 struct { F0 uint32; F1 bool; F2 uint32 }
type GenS44 struct { F0 string }
type GenS45 struct { F0 uint16; F1 bool; F2 bool }
type GenS46 struct { F0 int }
type GenS47 struct { F0 uint8; F1 int16; F2 string }
type GenS48 struct { F0 int64 }
type GenS49 struct { F0 string }
type GenS50 struct { F0 uint64; F1 bool; F2 int8 }
type GenS51 struct { F0 int32; F1 uint; F2 int64 }
type GenS52 struct { F0 uint64; F1 int32; F2 int }
type GenS53 struct { F0 int32; F1 int; F2 int8 }
func genFn0(p0 uint, p1 int16, p2 uint64) uint32 {
    p2 = p2
    p0 = (^(p0))
    var v7 uint8 = 249
    fmt.Println(p2)
    var v8 uint8 = v7
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 42
}
func genFn1(p0 uint8, p1 int8, p2 bool) uint32 {
    fmt.Println(p0, p0)
    fmt.Println(p0)
    v9 := GenS1{F0: "hello", F1: (!(209 > 0)), F2: genFn0(42, -1, 1)}
    var v10 uint16 = 307
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return 4294967295
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[int]uint32)
    v3 := []uint32{2}
    v4 := make(map[string]uint32)
    v5 := []string{"foo bar", ("\n\t" + "a"), (("foo bar" + "foo bar") + ("hello" + ("" + ("a" + "\n\t")))), ("\n\t" + (("\n\t" + "") + "hello"))}
    var v6 uint32 = 937
    for fi11 := 0; fi11 < 2; fi11++ {
        switch "" {
        case "hello":
            if false {
                var v12 int = fi11
                for fi13 := 0; fi13 < 2; fi13++ {
                    v14 := GenS2{F0: 445}
                    _ = fi13
                    _ = v14
                }
                switch "\n\t" {
                case "hello":
                    v1[v12] = "hello"
                case "a":
                    if true {
                        switch (("" + "a") + ((("\n\t" + ("\n\t" + (((("a" + ("\n\t" + "foo bar")) + ("hello" + ((("a" + "foo bar") + "") + "foo bar"))) + "hello") + "foo bar"))) + "foo bar") + "a")) {
                        case "":
                            if true {
                                for fi15 := 0; fi15 < 8; fi15++ {
                                    fmt.Println(fi11, fi15)
                                    _ = fi15
                                }
                                v16 := GenS3{F0: 1000, F1: 42, F2: 460}
                                _ = v16
                            }
                        case "a":
                            for fi17 := 0; fi17 < 11; fi17++ {
                                if (!false) {
                                    fmt.Println(fi11)
                                    _ = len(v2)
                                    v6 = (^(v6))
                                } else {
                                    v6 = v6
                                }
                                var v18 bool = false
                                _ = fi17
                                _ = v18
                            }
                        case "hello":
                            for fi19 := 0; fi19 < 17; fi19++ {
                                v12 += v12
                                _ = fi19
                            }
                        default:
                            var v20 int = (v12 << ((18446744073709551615) & 63))
                            _ = v20
                        }
                        if false {
                            var v21 uint8 = 92
                            var v22 uint16 = 500
                            _ = v21
                            _ = v22
                        } else {
                            v12 -= -2147483648
                            if false {
                                var v23 float32 = 1e10
                                _ = v23
                            } else {
                                v6 = (v6 & 499)
                                v24 := GenS4{F0: false, F1: 834}
                                if true {
                                    if true {
                                        if len(v5) > 0 {
                                            v5[0] = "hello"
                                        }
                                    } else {
                                        v25 := GenS1{F0: "hello", F1: (!(73 > 42)), F2: 4294967295}
                                        _ = v25
                                    }
                                    v26 := GenS5{F0: 42}
                                    if (v12 <= (-(v12))) {
                                        v27 := GenS6{F0: "foo bar", F1: -493}
                                        for fi28 := 0; fi28 < 10; fi28++ {
                                            fmt.Println(fi28, fi11)
                                            if ((0 <= 127) || (1 <= -128)) {
                                                v29 := v4["hello"]
                                                var v30 int8 = 1
                                                switch "\n\t" {
                                                case "a":
                                                    if false {
                                                        fmt.Println(fi11, v12, fi11)
                                                        switch ("foo bar" + "\n\t") {
                                                        case "":
                                                            if false {
                                                                v6 += v6
                                                                v6 = (^(v29))
                                                            } else {
                                                                var v31 uint = 704
                                                                var v32 int8 = v30
                                                                switch "\n\t" {
                                                                case "":
                                                                    if ((float64(3.14159) < 1e10) && (v31 != (^(v31)))) {
                                                                        switch 732 {
                                                                        case 4294967295:
                                                                            v30 -= v30
                                                                        }
                                                                        fmt.Println(v29, v12)
                                                                    } else {
                                                                        var v33 string = "foo bar"
                                                                        if false {
                                                                            v34 := GenS7{F0: (fi28 | fi28), F1: v29}
                                                                            _ = v34
                                                                        }
                                                                        _ = v33
                                                                    }
                                                                case "foo bar":
                                                                    var v35 int64 = 42
                                                                    _ = v35
                                                                case "\n\t":
                                                                    switch "a" {
                                                                    case "foo bar":
                                                                        v2[fi11] = (v29 / (((^(v6))) | 1))
                                                                    }
                                                                default:
                                                                    v12 -= 0
                                                                }
                                                                v36 := GenS2{F0: v31}
                                                                if len(v5) >= 2 {
                                                                    v5 = v5[1:len(v5)]
                                                                }
                                                                _ = v31
                                                                _ = v32
                                                                _ = v36
                                                            }
                                                        default:
                                                            fmt.Println(v29, v12, v29)
                                                            var v37 int16 = 42
                                                            _ = v37
                                                        }
                                                    } else {
                                                        v38 := GenS8{F0: -2147483648}
                                                        if (0 != 778) {
                                                            v6 += genFn0(1, 0, 4294967295)
                                                            _ = len(v5)
                                                        } else {
                                                            if false {
                                                                if false {
                                                                    if false {
                                                                        _ = len(v5)
                                                                    } else {
                                                                        if false {
                                                                            v39 := GenS2{F0: 2}
                                                                            _ = v39
                                                                        } else {
                                                                            v40 := GenS7{F0: 4294967295, F1: v29}
                                                                            _ = v40
                                                                        }
                                                                        v12 = 920
                                                                    }
                                                                    var v41 string = ""
                                                                    v42 := GenS2{F0: 485}
                                                                    _ = v41
                                                                    _ = v42
                                                                } else {
                                                                    if (!(-32768 > -1)) {
                                                                        v12 -= fi11
                                                                    } else {
                                                                        v43 := GenS3{F0: 1, F1: 9223372036854775807, F2: 851}
                                                                        var v44 uint32 = 22
                                                                        v44 += (v44 + v29)
                                                                        _ = v43
                                                                        _ = v44
                                                                    }
                                                                }
                                                            } else {
                                                                var v45 float32 = 3.14159
                                                                _ = v45
                                                            }
                                                            if (2 != 2) {
                                                                delete(v4, "")
                                                                switch "a" {
                                                                case "a":
                                                                    if (65535 <= 65535) {
                                                                        v12 = fi11
                                                                        if false {
                                                                            var v46 int64 = 42
                                                                            v47 := GenS5{F0: 18446744073709551615}
                                                                            _ = v46
                                                                            _ = v47
                                                                        }
                                                                        if false {
                                                                            if false {
                                                                                v48 := v4["\n\t"]
                                                                                var v49 uint64 = 2
                                                                                _ = v48
                                                                                _ = v49
                                                                            } else {
                                                                                var v50 float64 = ((3.14159 + 1e10) / 1.0)
                                                                                v51 := GenS8{F0: 413}
                                                                                _ = v50
                                                                                _ = v51
                                                                            }
                                                                        } else {
                                                                            var v52 int32 = -65
                                                                            v6 = genFn1(255, -128, true)
                                                                            _ = v52
                                                                        }
                                                                    } else {
                                                                        v12 = -169
                                                                        switch (v12 / ((9223372036854775807) | 1)) {
                                                                        case 1:
                                                                            if len(v3) > 0 {
                                                                                v3[0] = 2147483647
                                                                            }
                                                                        default:
                                                                            v6 = 2147483648
                                                                        }
                                                                    }
                                                                case "":
                                                                    switch "\n\t" {
                                                                    case "a":
                                                                        var v53 float64 = float64(0.0)
                                                                        _ = len(v5)
                                                                        _ = v53
                                                                    case "":
                                                                        v54 := GenS9{F0: 0, F1: "", F2: ("foo bar" + ("foo bar" + "\n\t"))}
                                                                        _ = v54
                                                                    case "foo bar":
                                                                        if len(v3) > 0 {
                                                                            v3[0] = genFn1(255, (v30 ^ (-(v30))), (-128 == 42))
                                                                        }
                                                                    }
                                                                case "foo bar":
                                                                    var v55 float32 = (0.0 * 3.14159)
                                                                    _ = v55
                                                                default:
                                                                    v29 += (v6 << ((65535) & 63))
                                                                    var v56 int16 = -32768
                                                                    var v57 uint64 = 4294967295
                                                                    _ = v56
                                                                    _ = v57
                                                                }
                                                            }
                                                        }
                                                        _ = v38
                                                    }
                                                }
                                                _ = v29
                                                _ = v30
                                            } else {
                                                var v58 float64 = 1e10
                                                v59 := GenS10{F0: 42, F1: 0}
                                                delete(v2, (-(fi28)))
                                                _ = v58
                                                _ = v59
                                            }
                                            if false {
                                                v6 -= v6
                                                var v60 int = -9223372036854775808
                                                v61 := GenS5{F0: 843}
                                                _ = v60
                                                _ = v61
                                            }
                                            _ = fi28
                                        }
                                        _ = v27
                                    } else {
                                        for fi62 := 0; fi62 < 17; fi62++ {
                                            v6 += (v6 ^ v6)
                                            v12 += v12
                                            _ = fi62
                                        }
                                    }
                                    _ = v26
                                } else {
                                    fmt.Println(v12)
                                }
                                _ = v24
                            }
                        }
                    } else {
                        var v63 bool = (true || true)
                        v64 := GenS10{F0: (v6 - v6), F1: 1}
                        if len(v3) > 0 {
                            v3[0] = (v6 - v6)
                        }
                        _ = v63
                        _ = v64
                    }
                case "foo bar":
                    v65 := GenS9{F0: 0, F1: "foo bar", F2: "hello"}
                    _ = v65
                }
                _ = v12
            }
        default:
            switch "foo bar" {
            case "foo bar":
                v6 += (^(v6))
            case "":
                for fi66 := 0; fi66 < 2; fi66++ {
                    v6 = (v6 % ((v6) | 1))
                    v6 += v6
                    _ = fi66
                }
            case "a":
                v67 := GenS2{F0: 1}
                _ = v67
            }
        }
        for fi68 := 0; fi68 < 20; fi68++ {
            fmt.Println(fi68, fi68, v6)
            var v69 uint32 = (v6 % ((4294967295) | 1))
            _ = fi68
            _ = v69
        }
        _ = len(v2)
        _ = fi11
    }
    if (153 < 0) {
        v6 += (v6 + 902)
    }
    v70 := GenS7{F0: -2147483648, F1: (v6 + genFn1(128, 127, false))}
    delete(v4, "hello")
    if true {
        v5 = append(v5, "a")
        fmt.Println(v6)
        switch 916 {
        case -2147483648:
            var v71 uint64 = 0
            _ = v71
        case 32767:
            var v72 int64 = 0
            _ = v72
        case 127:
            if ((("foo bar" + "\n\t") + ("a" + "\n\t")) <= "") {
                for fi73 := 0; fi73 < 10; fi73++ {
                    if false {
                        var v74 uint64 = 42
                        if (true && (0 < 0)) {
                            switch ("foo bar" + "") {
                            case "":
                                for fi75 := 0; fi75 < 19; fi75++ {
                                    if (!false) {
                                        fmt.Println(fi73)
                                        v2[fi75] = v6
                                    }
                                    v74 -= 42
                                    var v76 bool = true
                                    _ = fi75
                                    _ = v76
                                }
                            case "foo bar":
                                if false {
                                    for fi77 := 0; fi77 < 16; fi77++ {
                                        v74 += (v74 + (v74 % (((v74 ^ (^(v74)))) | 1)))
                                        _ = fi77
                                    }
                                    var v78 uint16 = 0
                                    _ = v78
                                } else {
                                    v6 += genFn1(127, 100, false)
                                }
                            default:
                                if (!false) {
                                    if ((-1 > 1) || ((false && false) && true)) {
                                        if true {
                                            for fi79 := 0; fi79 < 17; fi79++ {
                                                if (191 >= 5) {
                                                    v6 += (v6 ^ 914)
                                                }
                                                var v80 uint16 = 32767
                                                _ = fi79
                                                _ = v80
                                            }
                                            switch "\n\t" {
                                            case "foo bar":
                                                v74 += v74
                                            case "a":
                                                switch fi73 {
                                                case -9223372036854775808:
                                                    fmt.Println(v74, v6, v6)
                                                case -1:
                                                    fmt.Println(v6, v6)
                                                case 4294967295:
                                                    v6 = genFn0(2, 154, v74)
                                                default:
                                                    var v81 float32 = (1.5 - 0.1)
                                                    _ = v81
                                                }
                                            }
                                        }
                                    }
                                    delete(v4, "\n\t")
                                    var v82 float64 = (float64(3.14159) - ((3.14159 - (((1e10 + 3.14159) * (float64(0.1) / 1.0)) + 1e10)) * 1.5))
                                    _ = v82
                                } else {
                                    v3 = append(v3, 2147483648)
                                    fmt.Println(v74, v74)
                                    var v83 int32 = -1
                                    _ = v83
                                }
                            }
                            for fi84 := 0; fi84 < 13; fi84++ {
                                v74 = (^(v74))
                                if (!(!false)) {
                                    switch ("" + ("" + (("foo bar" + ("foo bar" + (("" + (("hello" + (("\n\t" + "hello") + "hello")) + "foo bar")) + (("" + "") + "hello")))) + ("a" + ("hello" + ""))))) {
                                    case "hello":
                                        if (((0.0 / (-2.25 + 0.1)) * 0.1) == 0.0) {
                                            var v85 bool = false
                                            var v86 string = ""
                                            _ = v85
                                            _ = v86
                                        }
                                    case "foo bar":
                                        switch -135 {
                                        case 0:
                                            delete(v1, fi84)
                                        case 9223372036854775807:
                                            var v87 string = "foo bar"
                                            _ = v87
                                        }
                                    case "a":
                                        switch "hello" {
                                        case "foo bar":
                                            switch fi84 {
                                            case 127:
                                                var v88 bool = true
                                                fmt.Println(fi84, v74)
                                                var v89 string = "hello"
                                                _ = v88
                                                _ = v89
                                            case 2147483647:
                                                v6 += 4294967295
                                            case 4294967295:
                                                v90 := GenS5{F0: 1}
                                                _ = v90
                                            default:
                                                if false {
                                                    var v91 string = "foo bar"
                                                    _ = v91
                                                } else {
                                                    v74 += v74
                                                    var v92 uint8 = 2
                                                    _ = v92
                                                }
                                            }
                                        case "\n\t":
                                            v74 += 502
                                        case "":
                                            if false {
                                                fmt.Println(fi84, v74, fi73)
                                            } else {
                                                if ("hello" <= "") {
                                                    v6 = (v6 % ((42) | 1))
                                                    var v93 bool = (!true)
                                                    fmt.Println(fi84, v6)
                                                    if v93 {
                                                        if v93 {
                                                            v93 = ((1.5 / -2.25) == 1e10)
                                                        }
                                                        fmt.Println(v93)
                                                        switch -9223372036854775808 {
                                                        case 127:
                                                            var v94 int16 = -770
                                                            _ = v94
                                                        case 4294967295:
                                                            v93 = (!(1 != 0))
                                                            if len(v5) >= 2 {
                                                                v5 = v5[1:len(v5)]
                                                            }
                                                            if false {
                                                                v3 = append(v3, (v6 | uint32(v74)))
                                                                var v95 uint8 = 78
                                                                if false {
                                                                    switch ("hello" + "foo bar") {
                                                                    case "\n\t":
                                                                        v96 := GenS11{F0: 104, F1: 255, F2: 734}
                                                                        _ = v96
                                                                    case "hello":
                                                                        switch (fi73 << ((42) & 63)) {
                                                                        case 32768:
                                                                            v93 = v93
                                                                        case -323:
                                                                            v5 = append(v5, ("\n\t" + ("hello" + "hello")))
                                                                        case 128:
                                                                            v6 += 1
                                                                        }
                                                                    default:
                                                                        switch "" {
                                                                        case "":
                                                                            var v97 uint16 = 255
                                                                            _ = v97
                                                                        case "foo bar":
                                                                            var v98 bool = (((1.5 * 3.14159) + 0.1) > ((((3.14159 + 3.14159) - (((-2.25 * 1.5) / (0.0 - 1e10)) / 0.1)) * 1e10) + -2.25))
                                                                            _ = v98
                                                                        case "a":
                                                                            fmt.Println(fi73, v74)
                                                                            var v99 float32 = ((1.5 / (0.0 + ((3.14159 + 3.14159) / ((-2.25 * 1e10) * (-2.25 / (3.14159 - (0.1 - 1e10))))))) / ((3.14159 + -2.25) / float32((3.14159 / -2.25))))
                                                                            var v100 float64 = -2.25
                                                                            _ = v100
                                                                            _ = v99
                                                                        default:
                                                                            if (255 == 2) {
                                                                                v95 = v95
                                                                                v95 = 0
                                                                                v5 = append(v5, (((("hello" + "hello") + ("\n\t" + "\n\t")) + "\n\t") + "foo bar"))
                                                                            } else {
                                                                                v101 := GenS7{F0: (fi73 ^ (fi73 / (((fi84 + 596)) | 1))), F1: genFn0(1, -1, (v74 ^ v74))}
                                                                                var v102 uint = 4294967295
                                                                                _ = v101
                                                                                _ = v102
                                                                            }
                                                                        }
                                                                    }
                                                                    var v103 int64 = -9223372036854775808
                                                                    if len(v3) > 0 {
                                                                        v3[0] = (v6 | v6)
                                                                    }
                                                                    var v104 int32 = 42
                                                                    v104 -= 42
                                                                    _ = v103
                                                                    _ = v104
                                                                } else {
                                                                    if true {
                                                                        v93 = true
                                                                    } else {
                                                                        var v105 uint16 = 744
                                                                        v105 = (v105 % (((v105 >> ((255) & 63))) | 1))
                                                                        v95 = v95
                                                                        if (false || (255 != 1)) {
                                                                            var v106 int64 = -9223372036854775808
                                                                            _ = len(v3)
                                                                            _ = v106
                                                                        }
                                                                        var v107 float32 = float32(-2.25)
                                                                        _ = v105
                                                                        _ = v107
                                                                    }
                                                                }
                                                                _ = v95
                                                            } else {
                                                                v74 -= 1
                                                                v74 += 42
                                                                v108 := GenS8{F0: int32(v74)}
                                                                _ = v108
                                                            }
                                                        default:
                                                            v5 = append(v5, "hello")
                                                        }
                                                    } else {
                                                        if (!v93) {
                                                            var v109 int8 = 127
                                                            switch fi73 {
                                                            case 127:
                                                                if (1000 < -146) {
                                                                    v110 := GenS2{F0: 255}
                                                                    switch "" {
                                                                    case "\n\t":
                                                                        v109 += 1
                                                                        fmt.Println(v109, fi73, v93)
                                                                    default:
                                                                        v111 := GenS12{F0: (v74 + v74)}
                                                                        _ = v111
                                                                    }
                                                                    v112 := v4[""]
                                                                    v109 = (v109 * (v109 ^ 42))
                                                                    delete(v1, fi73)
                                                                    _ = v110
                                                                    _ = v112
                                                                }
                                                                if len(v3) > 0 {
                                                                    v3[0] = (v6 | 65535)
                                                                }
                                                                v109 = (v109 % (((v109 / (((v109 & (v109 * (v109 - v109)))) | 1))) | 1))
                                                            case -466:
                                                                fmt.Println(v74, fi73, v74)
                                                            case 224:
                                                                var v113 float64 = 1e10
                                                                _ = v113
                                                            default:
                                                                switch fi84 {
                                                                case 32767:
                                                                    var v114 uint16 = 164
                                                                    _ = v114
                                                                case -141:
                                                                    if (!(1.5 < -2.25)) {
                                                                        if len(v3) > 0 {
                                                                            v3[0] = genFn0(uint(v74), -381, v74)
                                                                        }
                                                                    }
                                                                    switch ("a" + "\n\t") {
                                                                    case "foo bar":
                                                                        var v115 uint32 = genFn0(133, 42, (v74 << ((1) & 63)))
                                                                        _ = v115
                                                                    default:
                                                                        if (-1 <= -589) {
                                                                            v2[fi84] = v6
                                                                        }
                                                                    }
                                                                    if len(v3) >= 2 {
                                                                        v3 = v3[1:len(v3)]
                                                                    }
                                                                case 484:
                                                                    v74 = (v74 & (v74 + (^(v74))))
                                                                }
                                                            }
                                                            v74 += (^(v74))
                                                            _ = v109
                                                        } else {
                                                            v116 := GenS5{F0: 357}
                                                            switch "\n\t" {
                                                            case "\n\t":
                                                                var v117 int = (^(fi73))
                                                                _ = v117
                                                            case "foo bar":
                                                                fmt.Println(v74)
                                                            case "":
                                                                if v93 {
                                                                    v93 = (635 < fi84)
                                                                    fmt.Println(v93)
                                                                }
                                                            }
                                                            _ = v116
                                                        }
                                                        if len(v5) >= 2 {
                                                            v5 = v5[1:len(v5)]
                                                        }
                                                        fmt.Println(fi73)
                                                        var v118 int = (fi84 - -9223372036854775808)
                                                        _ = v118
                                                    }
                                                    _ = v93
                                                } else {
                                                    v74 -= uint64(v6)
                                                    if (((0 < 42) || true) && (!(false || (false && (int8(v6) >= -1))))) {
                                                        v74 = (^(v74))
                                                    }
                                                }
                                            }
                                        default:
                                            v6 = (v6 & (^(v6)))
                                            switch fi84 {
                                            case -9223372036854775808:
                                                if ((!(22 < 127)) || ("\n\t" >= ("" + "a"))) {
                                                    fmt.Println(fi73, v74, v6)
                                                    v6 += v6
                                                    var v119 string = ""
                                                    var v120 string = ((v119 + ("foo bar" + v119)) + v119)
                                                    v120 = (v120 + v120)
                                                    if true {
                                                        if true {
                                                            v121 := GenS7{F0: (fi84 ^ (^(fi84))), F1: v6}
                                                            v6 -= (v6 >> ((540) & 63))
                                                            var v122 bool = (!(65535 < 0))
                                                            _ = v121
                                                            _ = v122
                                                        } else {
                                                            var v123 uint8 = 1
                                                            var v124 uint = 255
                                                            if ((!true) && (!((-655 != 1) && false))) {
                                                                fmt.Println(v120, v123)
                                                                var v125 int8 = 127
                                                                _ = v125
                                                            } else {
                                                                v6 = 42
                                                                fmt.Println(v119, v123, v124)
                                                                if false {
                                                                    v120 = v119
                                                                } else {
                                                                    if len(v3) > 0 {
                                                                        v3[0] = 799
                                                                    }
                                                                    v126 := GenS13{F0: v124}
                                                                    if (!false) {
                                                                        switch 2147483647 {
                                                                        case -9223372036854775808:
                                                                            if false {
                                                                                switch fi73 {
                                                                                case -751:
                                                                                    v123 -= 42
                                                                                case 2147483647:
                                                                                    if false {
                                                                                        if true {
                                                                                            if (("\n\t" + "hello") == v119) {
                                                                                                v127 := GenS1{F0: v120, F1: (-625 > 1), F2: (v6 * (v6 & (^(v6))))}
                                                                                                _ = v127
                                                                                            } else {
                                                                                                var v128 float64 = ((float64(0.1) + float64(1e10)) * 0.0)
                                                                                                if (733 == (v74 / (((v74 ^ v74)) | 1))) {
                                                                                                    if (!(!false)) {
                                                                                                        var v129 uint16 = 0
                                                                                                        if (v123 != v123) {
                                                                                                            v120 = v119
                                                                                                            _ = len(v5)
                                                                                                            if false {
                                                                                                                v120 = "hello"
                                                                                                                _ = len(v5)
                                                                                                                var v130 int = (fi84 << (((v124 + 0)) & 63))
                                                                                                                _ = v130
                                                                                                            } else {
                                                                                                                var v131 int8 = 101
                                                                                                                if (v129 < uint16(v6)) {
                                                                                                                    switch (fi84 - -203) {
                                                                                                                    case 0:
                                                                                                                        v132 := GenS9{F0: (v129 | (v129 / ((v129) | 1))), F1: v120, F2: "\n\t"}
                                                                                                                        _ = v132
                                                                                                                    case -9223372036854775808:
                                                                                                                        if len(v3) >= 2 {
                                                                                                                            v3 = v3[1:len(v3)]
                                                                                                                        }
                                                                                                                        v74 -= v74
                                                                                                                    case 9223372036854775807:
                                                                                                                        var v133 uint16 = v129
                                                                                                                        _ = v133
                                                                                                                    default:
                                                                                                                        var v134 uint16 = v129
                                                                                                                        _ = v134
                                                                                                                    }
                                                                                                                    switch "" {
                                                                                                                    case "\n\t":
                                                                                                                        var v135 uint32 = 150
                                                                                                                        _ = v135
                                                                                                                    case "":
                                                                                                                        v136 := GenS14{F0: -509}
                                                                                                                        _ = v136
                                                                                                                    case "foo bar":
                                                                                                                        fmt.Println(v6)
                                                                                                                    }
                                                                                                                    v137 := v4[v120]
                                                                                                                    _ = v137
                                                                                                                }
                                                                                                                _ = v131
                                                                                                            }
                                                                                                        }
                                                                                                        if ((v124 >= 42) && (true && (uint16(v124) >= v129))) {
                                                                                                            var v138 uint32 = 2
                                                                                                            if ((v123 << ((0) & 63)) != v123) {
                                                                                                                v120 = ""
                                                                                                                if true {
                                                                                                                    if false {
                                                                                                                        v119 = v119
                                                                                                                        v139 := GenS14{F0: fi84}
                                                                                                                        v140 := GenS15{F0: v74}
                                                                                                                        v129 -= v129
                                                                                                                        _ = v139
                                                                                                                        _ = v140
                                                                                                                    } else {
                                                                                                                        v129 = 65535
                                                                                                                        switch fi84 {
                                                                                                                        case 4294967295:
                                                                                                                            var v141 int64 = int64(v74)
                                                                                                                            _ = v141
                                                                                                                        case 522:
                                                                                                                            fmt.Println(v138, fi73)
                                                                                                                        case -292:
                                                                                                                            v119 = v119
                                                                                                                        }
                                                                                                                        switch fi84 {
                                                                                                                        case 42:
                                                                                                                            switch fi84 {
                                                                                                                            case 0:
                                                                                                                                var v142 float64 = 3.14159
                                                                                                                                _ = v142
                                                                                                                            default:
                                                                                                                                v143 := GenS16{F0: 42}
                                                                                                                                _ = v143
                                                                                                                            }
                                                                                                                        default:
                                                                                                                            v144 := GenS12{F0: v74}
                                                                                                                            _ = v144
                                                                                                                        }
                                                                                                                    }
                                                                                                                    var v145 int64 = -1
                                                                                                                    var v146 string = ("foo bar" + v120)
                                                                                                                    v147 := GenS17{F0: (false || true), F1: 1, F2: false}
                                                                                                                    fmt.Println(v124)
                                                                                                                    _ = v145
                                                                                                                    _ = v146
                                                                                                                    _ = v147
                                                                                                                } else {
                                                                                                                    v148 := GenS18{F0: false}
                                                                                                                    var v149 int16 = 1000
                                                                                                                    var v150 uint = v124
                                                                                                                    _ = v148
                                                                                                                    _ = v149
                                                                                                                    _ = v150
                                                                                                                }
                                                                                                            }
                                                                                                            var v151 int = (fi73 * fi84)
                                                                                                            _ = v138
                                                                                                            _ = v151
                                                                                                        }
                                                                                                        v120 = (v119 + v120)
                                                                                                        v152 := GenS19{F0: false, F1: ((-9223372036854775808 < -1) && (-128 <= 1))}
                                                                                                        _ = v129
                                                                                                        _ = v152
                                                                                                    } else {
                                                                                                        if true {
                                                                                                            v153 := GenS7{F0: (^(fi73)), F1: (v6 + (^(v6)))}
                                                                                                            _ = v153
                                                                                                        } else {
                                                                                                            if true {
                                                                                                                switch (fi84 - (fi73 - fi73)) {
                                                                                                                case 42:
                                                                                                                    delete(v4, (v119 + v120))
                                                                                                                case 127:
                                                                                                                    v128 -= ((v128 + v128) - v128)
                                                                                                                case 869:
                                                                                                                    v5 = append(v5, v119)
                                                                                                                }
                                                                                                                if true {
                                                                                                                    fmt.Println(v119, fi73)
                                                                                                                }
                                                                                                                var v154 int8 = 1
                                                                                                                _ = v154
                                                                                                            } else {
                                                                                                                v155 := GenS11{F0: -1, F1: 42, F2: v6}
                                                                                                                _ = v155
                                                                                                            }
                                                                                                        }
                                                                                                        fmt.Println(v119)
                                                                                                        var v156 float32 = 3.14159
                                                                                                        v157 := GenS11{F0: 107, F1: 32767, F2: genFn0((v124 + v124), 32767, 2)}
                                                                                                        _ = len(v5)
                                                                                                        _ = v156
                                                                                                        _ = v157
                                                                                                    }
                                                                                                    v158 := GenS20{F0: (v123 * v123), F1: -1, F2: (v123 * 127)}
                                                                                                    var v159 bool = false
                                                                                                    if (false && v159) {
                                                                                                        switch ((v119 + v120) + v120) {
                                                                                                        case "\n\t":
                                                                                                            v160 := v2[fi73]
                                                                                                            var v161 uint64 = v74
                                                                                                            if true {
                                                                                                                _ = len(v5)
                                                                                                                fmt.Println(v160, v128, v123)
                                                                                                                var v162 bool = v159
                                                                                                                _ = v162
                                                                                                            } else {
                                                                                                                switch fi84 {
                                                                                                                case 127:
                                                                                                                    if (1 == -9223372036854775808) {
                                                                                                                        fmt.Println(v123)
                                                                                                                    } else {
                                                                                                                        v6 = 697
                                                                                                                        var v163 int16 = 42
                                                                                                                        _ = v163
                                                                                                                    }
                                                                                                                case 49:
                                                                                                                    fmt.Println(v160)
                                                                                                                case 0:
                                                                                                                    if len(v3) >= 2 {
                                                                                                                        v3 = v3[1:len(v3)]
                                                                                                                    }
                                                                                                                    switch v120 {
                                                                                                                    case "foo bar":
                                                                                                                        fmt.Println(v161)
                                                                                                                    }
                                                                                                                default:
                                                                                                                    fmt.Println(v74, v120, v161)
                                                                                                                }
                                                                                                                v4[""] = v160
                                                                                                            }
                                                                                                            _ = v160
                                                                                                            _ = v161
                                                                                                        case "":
                                                                                                            fmt.Println(v6, v159)
                                                                                                        }
                                                                                                        var v164 uint64 = (v74 - v74)
                                                                                                        _ = len(v3)
                                                                                                        _ = v164
                                                                                                    }
                                                                                                    if (-938 >= -32768) {
                                                                                                        fmt.Println(v123, fi73)
                                                                                                        if len(v3) >= 2 {
                                                                                                            v3 = v3[1:len(v3)]
                                                                                                        }
                                                                                                        switch fi73 {
                                                                                                        case 32768:
                                                                                                            v124 -= v124
                                                                                                        case -1:
                                                                                                            v165 := v1[fi73]
                                                                                                            _ = v165
                                                                                                        }
                                                                                                    }
                                                                                                    var v166 uint64 = 922
                                                                                                    _ = v158
                                                                                                    _ = v159
                                                                                                    _ = v166
                                                                                                } else {
                                                                                                    var v167 bool = (!true)
                                                                                                    v168 := v2[(fi73 % (((fi84 / (((fi73 * 398)) | 1))) | 1))]
                                                                                                    _ = v167
                                                                                                    _ = v168
                                                                                                }
                                                                                                _ = v128
                                                                                            }
                                                                                            var v169 uint32 = (v6 * genFn0((v124 + v124), 32767, (^(v74))))
                                                                                            v170 := GenS6{F0: v120, F1: 0}
                                                                                            switch 32767 {
                                                                                            case 400:
                                                                                                fmt.Println(v169, fi73, fi84)
                                                                                            case 324:
                                                                                                if len(v3) > 0 {
                                                                                                    v3[0] = v169
                                                                                                }
                                                                                            case -9223372036854775808:
                                                                                                if true {
                                                                                                    v123 -= 28
                                                                                                    v171 := GenS5{F0: 1}
                                                                                                    var v172 uint = (v124 / (((v124 / (((v124 >> ((v124) & 63))) | 1))) | 1))
                                                                                                    v173 := v2[fi84]
                                                                                                    fmt.Println(v119, v169)
                                                                                                    _ = v171
                                                                                                    _ = v172
                                                                                                    _ = v173
                                                                                                }
                                                                                            default:
                                                                                                var v174 int8 = 112
                                                                                                _ = v174
                                                                                            }
                                                                                            _ = v169
                                                                                            _ = v170
                                                                                        } else {
                                                                                            if true {
                                                                                                v120 = (v119 + v120)
                                                                                                var v175 float32 = 3.14159
                                                                                                switch (fi84 << ((v124) & 63)) {
                                                                                                case 42:
                                                                                                    v175 -= v175
                                                                                                case 985:
                                                                                                    if false {
                                                                                                        v6 = genFn1((v123 | (v123 & (v123 & (v123 << ((uint(v6)) & 63))))), -62, (!true))
                                                                                                    } else {
                                                                                                        v123 = v123
                                                                                                        if false {
                                                                                                            fmt.Println(v119, v123, v124)
                                                                                                        } else {
                                                                                                            v176 := GenS21{F0: (^(v123))}
                                                                                                            v177 := GenS7{F0: 4294967295, F1: v6}
                                                                                                            if (true || (!true)) {
                                                                                                                var v178 uint32 = (v6 * (v6 / (((v6 * (v6 * 65535))) | 1)))
                                                                                                                _ = v178
                                                                                                            } else {
                                                                                                                var v179 string = (v120 + ("foo bar" + "a"))
                                                                                                                if true {
                                                                                                                    if true {
                                                                                                                        v180 := GenS5{F0: (v124 | (^(v124)))}
                                                                                                                        _ = v180
                                                                                                                    }
                                                                                                                    v181 := v1[(fi84 | fi84)]
                                                                                                                    _ = v181
                                                                                                                }
                                                                                                                _ = v179
                                                                                                            }
                                                                                                            _ = v176
                                                                                                            _ = v177
                                                                                                        }
                                                                                                        v123 = v123
                                                                                                    }
                                                                                                }
                                                                                                _ = v175
                                                                                            }
                                                                                            var v182 uint8 = v123
                                                                                            _ = v182
                                                                                        }
                                                                                        if (float32(3.14159) == (0.0 * 1e10)) {
                                                                                            var v183 uint = (v124 & (^(v124)))
                                                                                            if false {
                                                                                                v184 := GenS7{F0: fi73, F1: v6}
                                                                                                delete(v4, v120)
                                                                                                _ = v184
                                                                                            } else {
                                                                                                if false {
                                                                                                    v185 := GenS22{F0: true}
                                                                                                    delete(v2, int(v6))
                                                                                                    v123 += (v123 + 42)
                                                                                                    v123 += 127
                                                                                                    var v186 uint32 = (v6 % ((101) | 1))
                                                                                                    _ = v185
                                                                                                    _ = v186
                                                                                                } else {
                                                                                                    v187 := GenS21{F0: 1}
                                                                                                    if true {
                                                                                                        if false {
                                                                                                            if len(v3) >= 2 {
                                                                                                                v3 = v3[1:len(v3)]
                                                                                                            }
                                                                                                        } else {
                                                                                                            var v188 int = (fi84 >> (((v124 - v124)) & 63))
                                                                                                            _ = v188
                                                                                                        }
                                                                                                        if (1000 != -449) {
                                                                                                            if len(v3) > 0 {
                                                                                                                v3[0] = (^(v6))
                                                                                                            }
                                                                                                            fmt.Println(v119)
                                                                                                            v1[(fi84 - -651)] = v120
                                                                                                        } else {
                                                                                                            if (42 != 1000) {
                                                                                                                v189 := GenS18{F0: (-1 <= 1)}
                                                                                                                fmt.Println(v183, v74)
                                                                                                                _ = v189
                                                                                                            } else {
                                                                                                                var v190 string = v120
                                                                                                                _ = v190
                                                                                                            }
                                                                                                            v123 += uint8(v6)
                                                                                                            if true {
                                                                                                                var v191 uint16 = 706
                                                                                                                switch ("a" + "\n\t") {
                                                                                                                case "\n\t":
                                                                                                                    switch fi84 {
                                                                                                                    case 32767:
                                                                                                                        v192 := GenS23{F0: v191}
                                                                                                                        _ = v192
                                                                                                                    case -992:
                                                                                                                        var v193 int64 = -295
                                                                                                                        _ = v193
                                                                                                                    default:
                                                                                                                        v191 = 65535
                                                                                                                    }
                                                                                                                case "":
                                                                                                                    fmt.Println(v123, v119)
                                                                                                                    v119 = ("a" + ((v119 + (v119 + v120)) + ""))
                                                                                                                case "hello":
                                                                                                                    switch v119 {
                                                                                                                    case "hello":
                                                                                                                        if (!(((v74 << (((v124 & v183)) & 63)) != v74) && (((v6 ^ (v6 ^ (^(v6)))) != 260) || (true && (!(true && ((!false) || true))))))) {
                                                                                                                            v119 = ((v119 + "") + v120)
                                                                                                                            var v194 string = v119
                                                                                                                            _ = v194
                                                                                                                        }
                                                                                                                    default:
                                                                                                                        var v195 int64 = -112
                                                                                                                        _ = v195
                                                                                                                    }
                                                                                                                default:
                                                                                                                    fmt.Println(v191)
                                                                                                                }
                                                                                                                _ = v191
                                                                                                            } else {
                                                                                                                v196 := GenS9{F0: 1, F1: v119, F2: "a"}
                                                                                                                v120 = (v119 + "a")
                                                                                                                var v197 int32 = -2147483648
                                                                                                                _ = v196
                                                                                                                _ = v197
                                                                                                            }
                                                                                                        }
                                                                                                        if false {
                                                                                                            v198 := GenS23{F0: 42}
                                                                                                            v120 = "hello"
                                                                                                            _ = v198
                                                                                                        }
                                                                                                    }
                                                                                                    _ = v187
                                                                                                }
                                                                                            }
                                                                                            _ = v183
                                                                                        }
                                                                                    }
                                                                                case -9223372036854775808:
                                                                                    fmt.Println(v120, fi73, v120)
                                                                                }
                                                                            }
                                                                        case 9223372036854775807:
                                                                            var v199 uint32 = (v6 - v6)
                                                                            _ = v199
                                                                        default:
                                                                            _ = len(v5)
                                                                        }
                                                                    }
                                                                    _ = v126
                                                                }
                                                            }
                                                            fmt.Println(v124, v119)
                                                            _ = v123
                                                            _ = v124
                                                        }
                                                        v200 := GenS10{F0: v6, F1: v74}
                                                        v119 = v119
                                                        if len(v3) > 0 {
                                                            v3[0] = (v6 ^ (v6 & genFn1(0, 0, false)))
                                                        }
                                                        if true {
                                                            if false {
                                                                fmt.Println(v120, v6, v74)
                                                                v1[fi84] = "foo bar"
                                                                _ = len(v2)
                                                            }
                                                        } else {
                                                            var v201 float32 = 0.1
                                                            if true {
                                                                v120 = v120
                                                                fmt.Println(v6)
                                                                v202 := GenS9{F0: 0, F1: "foo bar", F2: "\n\t"}
                                                                _ = v202
                                                            }
                                                            switch v119 {
                                                            case "":
                                                                if true {
                                                                    v6 -= v6
                                                                }
                                                            case "\n\t":
                                                                v6 = genFn1(1, 0, (3.14159 != 3.14159))
                                                            case "foo bar":
                                                                switch (v120 + v120) {
                                                                case "":
                                                                    v6 -= 0
                                                                case "a":
                                                                    v3 = append(v3, 65535)
                                                                case "foo bar":
                                                                    var v203 int = fi84
                                                                    _ = v203
                                                                default:
                                                                    if (300 >= 1000) {
                                                                        v3 = append(v3, (^(v6)))
                                                                        switch (^(fi73)) {
                                                                        case -559:
                                                                            switch v120 {
                                                                            case "":
                                                                                v119 = "\n\t"
                                                                            case "\n\t":
                                                                                v204 := GenS24{F0: v6}
                                                                                _ = v204
                                                                            default:
                                                                                switch v119 {
                                                                                case "foo bar":
                                                                                    v205 := v1[-310]
                                                                                    _ = v205
                                                                                }
                                                                            }
                                                                        case -191:
                                                                            v201 += float32(-2.25)
                                                                        case 4294967295:
                                                                            v206 := v4[(v119 + ((v120 + v120) + "foo bar"))]
                                                                            _ = v206
                                                                        default:
                                                                            v207 := GenS20{F0: uint8(v6), F1: -1, F2: 2}
                                                                            _ = v207
                                                                        }
                                                                        _ = len(v2)
                                                                    }
                                                                    if ((v74 + v74) != v74) {
                                                                        v120 = ((v119 + "") + v119)
                                                                    } else {
                                                                        var v208 int64 = 0
                                                                        v209 := GenS24{F0: v6}
                                                                        switch (fi73 - fi84) {
                                                                        case -2147483648:
                                                                            v120 = (v119 + (("foo bar" + "foo bar") + "hello"))
                                                                        case 4294967295:
                                                                            v2[42] = 2147483647
                                                                        case -16:
                                                                            switch fi73 {
                                                                            case 459:
                                                                                if len(v5) >= 2 {
                                                                                    v5 = v5[1:len(v5)]
                                                                                }
                                                                            case 803:
                                                                                v208 -= v208
                                                                            }
                                                                        }
                                                                        _ = v208
                                                                        _ = v209
                                                                    }
                                                                }
                                                            default:
                                                                switch (v119 + "") {
                                                                case "a":
                                                                    fmt.Println(fi73)
                                                                }
                                                            }
                                                            _ = v201
                                                        }
                                                        switch v119 {
                                                        case "a":
                                                            if len(v3) > 0 {
                                                                v3[0] = (v6 << ((803) & 63))
                                                            }
                                                            var v210 string = (v120 + ("a" + (v119 + "")))
                                                            var v211 int32 = 42
                                                            _ = v210
                                                            _ = v211
                                                        }
                                                        var v212 float32 = (1.5 + ((0.0 * (3.14159 - (1.5 / float32(0.1)))) + -2.25))
                                                        _ = v200
                                                        _ = v212
                                                    } else {
                                                        _ = len(v5)
                                                        if false {
                                                            fmt.Println(fi84, fi73, v120)
                                                            if ((v6 * v6) != genFn1(178, 0, (true || true))) {
                                                                fmt.Println(fi84, v6, v119)
                                                                if ((128 == 0) && (true && true)) {
                                                                    v213 := GenS6{F0: ((v120 + (("" + "\n\t") + (v120 + ((v120 + (v120 + v119)) + ("hello" + "foo bar"))))) + ((((v120 + v119) + v119) + (v119 + (v120 + v120))) + (((("" + "foo bar") + "foo bar") + ((v120 + "hello") + v120)) + v119))), F1: 281}
                                                                    v214 := GenS24{F0: v6}
                                                                    v119 = (v119 + v120)
                                                                    _ = v213
                                                                    _ = v214
                                                                } else {
                                                                    fmt.Println(v6, fi84)
                                                                    if ((^(v74)) < v74) {
                                                                        var v215 int16 = 1
                                                                        v119 = "a"
                                                                        _ = v215
                                                                    }
                                                                    var v216 uint16 = 573
                                                                    _ = v216
                                                                }
                                                            }
                                                            _ = len(v5)
                                                            _ = len(v3)
                                                            var v217 uint8 = 42
                                                            if (42 != 119) {
                                                                var v218 uint32 = genFn1(v217, 16, (false && false))
                                                                var v219 uint64 = v74
                                                                _ = v218
                                                                _ = v219
                                                            }
                                                            if false {
                                                                _ = len(v3)
                                                                var v220 uint32 = (v6 / (((^(v6))) | 1))
                                                                v1[(fi84 * int(v217))] = v120
                                                                _ = v220
                                                            } else {
                                                                v221 := GenS25{F0: v74, F1: (fi73 ^ fi73), F2: 1}
                                                                fmt.Println(v119, v217, v74)
                                                                _ = v221
                                                            }
                                                            _ = v217
                                                        } else {
                                                            v120 = ""
                                                            fmt.Println(fi84)
                                                            if (2 != uint16(v6)) {
                                                                v222 := GenS3{F0: 42, F1: 9223372036854775807, F2: v74}
                                                                switch (^(fi73)) {
                                                                case -27:
                                                                    v6 -= (v6 + v6)
                                                                case -1:
                                                                    delete(v4, "foo bar")
                                                                case 0:
                                                                    if false {
                                                                        if ((0.0 + 3.14159) > 1e10) {
                                                                            if (!true) {
                                                                                if true {
                                                                                    var v223 uint64 = (v74 | v74)
                                                                                    var v224 int32 = 1
                                                                                    var v225 int64 = 9223372036854775807
                                                                                    var v226 float32 = 3.14159
                                                                                    _ = v223
                                                                                    _ = v224
                                                                                    _ = v225
                                                                                    _ = v226
                                                                                } else {
                                                                                    fmt.Println(fi84)
                                                                                }
                                                                                if (float32(((-2.25 * 0.0) * (0.1 - 0.0))) >= 1.5) {
                                                                                    if (!(42 <= 65535)) {
                                                                                        var v227 uint32 = v6
                                                                                        _ = v227
                                                                                    }
                                                                                    var v228 bool = (v6 == v6)
                                                                                    switch fi84 {
                                                                                    case 42:
                                                                                        var v229 uint64 = v74
                                                                                        _ = v229
                                                                                    case 2147483647:
                                                                                        v6 += 880
                                                                                    case 128:
                                                                                        v228 = (v228 && true)
                                                                                    }
                                                                                    switch v120 {
                                                                                    case "":
                                                                                        v6 += v6
                                                                                    case "hello":
                                                                                        v230 := GenS26{F0: 127}
                                                                                        _ = v230
                                                                                    }
                                                                                    var v231 int64 = -9223372036854775808
                                                                                    _ = v228
                                                                                    _ = v231
                                                                                }
                                                                                var v232 int8 = 42
                                                                                _ = v232
                                                                            } else {
                                                                                var v233 float32 = 3.14159
                                                                                var v234 float64 = 3.14159
                                                                                var v235 string = v120
                                                                                _ = v233
                                                                                _ = v234
                                                                                _ = v235
                                                                            }
                                                                            var v236 float32 = -2.25
                                                                            fmt.Println(v74, v236, v6)
                                                                            _ = v236
                                                                        }
                                                                        _ = len(v1)
                                                                        var v237 int = fi84
                                                                        _ = v237
                                                                    } else {
                                                                        if (1.5 == 1e10) {
                                                                            fmt.Println(v119, v120)
                                                                        }
                                                                        switch fi84 {
                                                                        case 4294967295:
                                                                            var v238 float64 = 0.0
                                                                            _ = v238
                                                                        default:
                                                                            var v239 uint16 = 813
                                                                            _ = v239
                                                                        }
                                                                        fmt.Println(fi73, fi84, fi73)
                                                                    }
                                                                default:
                                                                    v120 = ((v120 + ((("hello" + "") + v119) + (v119 + ""))) + "\n\t")
                                                                }
                                                                _ = v222
                                                            }
                                                            var v240 int64 = 9223372036854775807
                                                            switch "" {
                                                            case "hello":
                                                                fmt.Println(v74, fi73)
                                                            case "a":
                                                                v241 := GenS2{F0: 1}
                                                                _ = v241
                                                            case "foo bar":
                                                                switch 32768 {
                                                                case 127:
                                                                    var v242 float64 = 1.5
                                                                    _ = v242
                                                                case 0:
                                                                    v6 += v6
                                                                case 1:
                                                                    if (!(!true)) {
                                                                        v243 := GenS14{F0: fi84}
                                                                        switch "a" {
                                                                        case "\n\t":
                                                                            v6 += (^(v6))
                                                                        case "foo bar":
                                                                            switch (fi84 * (-(fi73))) {
                                                                            case -597:
                                                                                switch (v120 + v120) {
                                                                                case "a":
                                                                                    var v244 float64 = 1e10
                                                                                    _ = v244
                                                                                case "":
                                                                                    if (false || false) {
                                                                                        switch v120 {
                                                                                        case "foo bar":
                                                                                            switch -333 {
                                                                                            case 0:
                                                                                                if true {
                                                                                                    v240 = v240
                                                                                                    fmt.Println(v74, fi84)
                                                                                                }
                                                                                                var v245 uint64 = (v74 & 1)
                                                                                                v246 := GenS27{F0: -621, F1: 111, F2: 175}
                                                                                                _ = v245
                                                                                                _ = v246
                                                                                            case 128:
                                                                                                if true {
                                                                                                    v247 := GenS28{F0: (^(v6))}
                                                                                                    _ = v247
                                                                                                }
                                                                                            default:
                                                                                                var v248 float64 = 0.0
                                                                                                _ = v248
                                                                                            }
                                                                                            if true {
                                                                                                v74 = v74
                                                                                            } else {
                                                                                                if (v240 == (^(v240))) {
                                                                                                    if true {
                                                                                                        if (((-2.25 * 0.1) + -2.25) <= -2.25) {
                                                                                                            if false {
                                                                                                                _ = len(v2)
                                                                                                                var v249 int = (fi73 * (fi84 - fi84))
                                                                                                                v119 = (("foo bar" + v120) + (("foo bar" + (("\n\t" + v120) + "hello")) + v120))
                                                                                                                var v250 float64 = 3.14159
                                                                                                                _ = v249
                                                                                                                _ = v250
                                                                                                            }
                                                                                                        }
                                                                                                        v119 = v120
                                                                                                        var v251 uint8 = uint8(v74)
                                                                                                        _ = v251
                                                                                                    }
                                                                                                } else {
                                                                                                    if true {
                                                                                                        delete(v4, v119)
                                                                                                        v74 = (^(v74))
                                                                                                    } else {
                                                                                                        if true {
                                                                                                            v119 = v120
                                                                                                            if (("foo bar" + v120) < (v120 + ((("hello" + ("foo bar" + v120)) + "") + ("" + "hello")))) {
                                                                                                                switch v120 {
                                                                                                                case "hello":
                                                                                                                    v4[(v120 + v120)] = (v6 % (((v6 & (v6 & v6))) | 1))
                                                                                                                case "foo bar":
                                                                                                                    var v252 int64 = (v240 + v240)
                                                                                                                    _ = v252
                                                                                                                }
                                                                                                                delete(v4, "hello")
                                                                                                            }
                                                                                                            var v253 uint64 = 18446744073709551615
                                                                                                            _ = v253
                                                                                                        }
                                                                                                        var v254 uint64 = v74
                                                                                                        v74 = (v74 ^ v254)
                                                                                                        _ = v254
                                                                                                    }
                                                                                                    if ((65535 <= 42) || ((true && true) && false)) {
                                                                                                        switch 127 {
                                                                                                        case 0:
                                                                                                            v255 := v2[(fi73 * fi84)]
                                                                                                            _ = v255
                                                                                                        case -681:
                                                                                                            switch 4294967295 {
                                                                                                            case 2147483647:
                                                                                                                v120 = v119
                                                                                                                v256 := GenS17{F0: (false || true), F1: -646, F2: (false || true)}
                                                                                                                v119 = ((v120 + ((v119 + v119) + v120)) + "foo bar")
                                                                                                                _ = v256
                                                                                                            }
                                                                                                        case -9223372036854775808:
                                                                                                            fmt.Println(v120, v120)
                                                                                                        }
                                                                                                        var v257 string = "\n\t"
                                                                                                        if (!(float32((0.0 / 3.14159)) != (0.1 + 1.5))) {
                                                                                                            if len(v3) >= 2 {
                                                                                                                v3 = v3[1:len(v3)]
                                                                                                            }
                                                                                                            switch fi84 {
                                                                                                            case 1:
                                                                                                                if true {
                                                                                                                    switch fi73 {
                                                                                                                    case -145:
                                                                                                                        if true {
                                                                                                                            var v258 string = "a"
                                                                                                                            var v259 uint16 = 32767
                                                                                                                            var v260 int16 = -1
                                                                                                                            v261 := v4["hello"]
                                                                                                                            _ = v258
                                                                                                                            _ = v259
                                                                                                                            _ = v260
                                                                                                                            _ = v261
                                                                                                                        }
                                                                                                                        v5 = append(v5, v257)
                                                                                                                        if false {
                                                                                                                            if false {
                                                                                                                                fmt.Println(v257)
                                                                                                                                var v262 string = "a"
                                                                                                                                v5 = append(v5, ("hello" + "foo bar"))
                                                                                                                                _ = v262
                                                                                                                            } else {
                                                                                                                                fmt.Println(fi84, v257, v119)
                                                                                                                                switch ("hello" + "hello") {
                                                                                                                                case "foo bar":
                                                                                                                                    if (-2147483648 != 188) {
                                                                                                                                        v74 += v74
                                                                                                                                        v263 := GenS29{F0: 1}
                                                                                                                                        _ = v263
                                                                                                                                    } else {
                                                                                                                                        var v264 uint8 = 38
                                                                                                                                        _ = v264
                                                                                                                                    }
                                                                                                                                case "a":
                                                                                                                                    var v265 int8 = -112
                                                                                                                                    _ = v265
                                                                                                                                default:
                                                                                                                                    switch fi84 {
                                                                                                                                    case 292:
                                                                                                                                        fmt.Println(v120, fi73)
                                                                                                                                    case -430:
                                                                                                                                        var v266 int64 = v240
                                                                                                                                        _ = v266
                                                                                                                                    default:
                                                                                                                                        v6 += (v6 << ((0) & 63))
                                                                                                                                    }
                                                                                                                                }
                                                                                                                            }
                                                                                                                        } else {
                                                                                                                            var v267 uint32 = v6
                                                                                                                            v267 -= genFn1(91, 42, (95 < 202))
                                                                                                                            v257 = (("a" + (v257 + "\n\t")) + "")
                                                                                                                            _ = v267
                                                                                                                        }
                                                                                                                        if (true || (v74 > (^(v74)))) {
                                                                                                                            var v268 int8 = -1
                                                                                                                            var v269 uint = 0
                                                                                                                            if len(v5) > 0 {
                                                                                                                                v5[0] = ((v119 + "") + "foo bar")
                                                                                                                            }
                                                                                                                            _ = v268
                                                                                                                            _ = v269
                                                                                                                        } else {
                                                                                                                            v270 := GenS30{F0: 654, F1: -46}
                                                                                                                            switch (v119 + v120) {
                                                                                                                            case "\n\t":
                                                                                                                                if (v6 != (^(v6))) {
                                                                                                                                    fmt.Println(v74, v240, v240)
                                                                                                                                    var v271 uint = 65535
                                                                                                                                    var v272 int32 = 42
                                                                                                                                    _ = v271
                                                                                                                                    _ = v272
                                                                                                                                } else {
                                                                                                                                    fmt.Println(v119, v6)
                                                                                                                                }
                                                                                                                            case "":
                                                                                                                                if (((-(fi73)) <= fi73) || (false || (v6 == v6))) {
                                                                                                                                    var v273 uint8 = uint8(v6)
                                                                                                                                    _ = v273
                                                                                                                                } else {
                                                                                                                                    fmt.Println(v257, v74, v119)
                                                                                                                                    delete(v4, v119)
                                                                                                                                }
                                                                                                                            default:
                                                                                                                                var v274 uint32 = (v6 & (v6 + v6))
                                                                                                                                v275 := GenS23{F0: 42}
                                                                                                                                delete(v2, (fi84 * fi73))
                                                                                                                                _ = v274
                                                                                                                                _ = v275
                                                                                                                            }
                                                                                                                            v6 = (v6 ^ 1)
                                                                                                                            _ = v270
                                                                                                                        }
                                                                                                                        var v276 uint64 = 2
                                                                                                                        if true {
                                                                                                                            switch (-(fi84)) {
                                                                                                                            case -326:
                                                                                                                                v276 += v74
                                                                                                                            }
                                                                                                                        } else {
                                                                                                                            switch fi84 {
                                                                                                                            case 32768:
                                                                                                                                var v277 int = 551
                                                                                                                                _ = v277
                                                                                                                            case -9223372036854775808:
                                                                                                                                var v278 string = "\n\t"
                                                                                                                                _ = v278
                                                                                                                            }
                                                                                                                            _ = len(v5)
                                                                                                                            if false {
                                                                                                                                v257 = v120
                                                                                                                                v257 = ((v120 + (v119 + v120)) + (v257 + ""))
                                                                                                                            } else {
                                                                                                                                v6 += (v6 ^ v6)
                                                                                                                                v279 := GenS31{F0: 1}
                                                                                                                                var v280 float64 = float64(0.0)
                                                                                                                                _ = v279
                                                                                                                                _ = v280
                                                                                                                            }
                                                                                                                            var v281 int = fi73
                                                                                                                            if true {
                                                                                                                                _ = len(v4)
                                                                                                                                if false {
                                                                                                                                    var v282 uint8 = 255
                                                                                                                                    v276 = v276
                                                                                                                                    _ = v282
                                                                                                                                }
                                                                                                                                switch fi84 {
                                                                                                                                case 32768:
                                                                                                                                    fmt.Println(v119, v281, v6)
                                                                                                                                case 705:
                                                                                                                                    _ = len(v3)
                                                                                                                                default:
                                                                                                                                    v283 := GenS32{F0: 70, F1: 255}
                                                                                                                                    _ = v283
                                                                                                                                }
                                                                                                                            } else {
                                                                                                                                var v284 int16 = -1
                                                                                                                                _ = v284
                                                                                                                            }
                                                                                                                            _ = v281
                                                                                                                        }
                                                                                                                        _ = v276
                                                                                                                    }
                                                                                                                    v285 := GenS19{F0: (-2147483648 == 2147483647), F1: (!false)}
                                                                                                                    var v286 uint64 = (v74 - v74)
                                                                                                                    _ = v285
                                                                                                                    _ = v286
                                                                                                                }
                                                                                                            }
                                                                                                            v287 := GenS33{F0: (!((1 == 128) || (!true))), F1: -1}
                                                                                                            switch (fi73 & 32768) {
                                                                                                            case 257:
                                                                                                                switch fi73 {
                                                                                                                case 32768:
                                                                                                                    if false {
                                                                                                                        v288 := GenS28{F0: 4294967295}
                                                                                                                        _ = v288
                                                                                                                    }
                                                                                                                default:
                                                                                                                    var v289 float64 = 1e10
                                                                                                                    _ = v289
                                                                                                                }
                                                                                                            case -1:
                                                                                                                v290 := v1[fi73]
                                                                                                                _ = v290
                                                                                                            case -430:
                                                                                                                var v291 int64 = (v240 + v240)
                                                                                                                _ = v291
                                                                                                            }
                                                                                                            _ = v287
                                                                                                        }
                                                                                                        var v292 uint8 = 2
                                                                                                        _ = v257
                                                                                                        _ = v292
                                                                                                    } else {
                                                                                                        v293 := GenS34{F0: 127}
                                                                                                        _ = v293
                                                                                                    }
                                                                                                    if true {
                                                                                                        if false {
                                                                                                            var v294 string = "\n\t"
                                                                                                            var v295 int32 = int32(v240)
                                                                                                            v5 = append(v5, v294)
                                                                                                            _ = v294
                                                                                                            _ = v295
                                                                                                        } else {
                                                                                                            var v296 int = fi73
                                                                                                            switch (fi73 << ((255) & 63)) {
                                                                                                            case 128:
                                                                                                                if len(v5) > 0 {
                                                                                                                    v5[0] = ("foo bar" + v119)
                                                                                                                }
                                                                                                            case 32768:
                                                                                                                if (fi73 == v296) {
                                                                                                                    delete(v1, v296)
                                                                                                                    if false {
                                                                                                                        if false {
                                                                                                                            if len(v5) >= 2 {
                                                                                                                                v5 = v5[1:len(v5)]
                                                                                                                            }
                                                                                                                            delete(v1, (fi84 & v296))
                                                                                                                        }
                                                                                                                    } else {
                                                                                                                        fmt.Println(v74, fi73, fi73)
                                                                                                                    }
                                                                                                                    if len(v3) >= 2 {
                                                                                                                        v3 = v3[1:len(v3)]
                                                                                                                    }
                                                                                                                } else {
                                                                                                                    v297 := GenS35{F0: v6}
                                                                                                                    _ = v297
                                                                                                                }
                                                                                                            case 2147483647:
                                                                                                                v296 = (fi73 % ((fi84) | 1))
                                                                                                                v298 := GenS16{F0: -1}
                                                                                                                _ = v298
                                                                                                            }
                                                                                                            v296 -= fi73
                                                                                                            switch v120 {
                                                                                                            case "foo bar":
                                                                                                                if (((!((true || (!false)) && true)) && false) || (false || (true && false))) {
                                                                                                                    v299 := GenS17{F0: (-62 <= 32767), F1: 898, F2: (127 == -1)}
                                                                                                                    fmt.Println(v119, v74, v74)
                                                                                                                    _ = v299
                                                                                                                }
                                                                                                            case "":
                                                                                                                if true {
                                                                                                                    var v300 uint8 = 255
                                                                                                                    var v301 int8 = int8(v300)
                                                                                                                    switch -476 {
                                                                                                                    case -1:
                                                                                                                        v302 := GenS36{F0: (fi73 / (((fi84 >> ((265) & 63))) | 1)), F1: int32(v6), F2: v119}
                                                                                                                        _ = v302
                                                                                                                    case 9223372036854775807:
                                                                                                                        var v303 bool = (!true)
                                                                                                                        _ = v303
                                                                                                                    }
                                                                                                                    _ = v300
                                                                                                                    _ = v301
                                                                                                                }
                                                                                                            default:
                                                                                                                var v304 int32 = -570
                                                                                                                _ = v304
                                                                                                            }
                                                                                                            switch v119 {
                                                                                                            case "a":
                                                                                                                switch (^(v296)) {
                                                                                                                case -1:
                                                                                                                    if true {
                                                                                                                        switch v119 {
                                                                                                                        case "hello":
                                                                                                                            if (-1 == 1000) {
                                                                                                                                v305 := GenS37{F0: 1000, F1: -4}
                                                                                                                                _ = len(v1)
                                                                                                                                v74 = (v74 & 305)
                                                                                                                                _ = v305
                                                                                                                            }
                                                                                                                        default:
                                                                                                                            var v306 int32 = -1
                                                                                                                            _ = v306
                                                                                                                        }
                                                                                                                    }
                                                                                                                case 127:
                                                                                                                    var v307 float64 = 3.14159
                                                                                                                    _ = v307
                                                                                                                }
                                                                                                            case "hello":
                                                                                                                fmt.Println(fi84, v120)
                                                                                                            case "\n\t":
                                                                                                                switch v119 {
                                                                                                                case "hello":
                                                                                                                    var v308 uint64 = (v74 / ((4294967295) | 1))
                                                                                                                    v296 = (-(fi84))
                                                                                                                    _ = len(v2)
                                                                                                                    _ = v308
                                                                                                                case "a":
                                                                                                                    v309 := GenS38{F0: 1}
                                                                                                                    _ = v309
                                                                                                                case "foo bar":
                                                                                                                    v310 := GenS39{F0: 867, F1: false, F2: -92}
                                                                                                                    _ = v310
                                                                                                                default:
                                                                                                                    fmt.Println(fi84, v6)
                                                                                                                }
                                                                                                            }
                                                                                                            if true {
                                                                                                                var v311 bool = (690 >= 0)
                                                                                                                v3 = append(v3, genFn0(2, 42, (v74 << ((255) & 63))))
                                                                                                                fmt.Println(fi73, v119, v296)
                                                                                                                _ = v311
                                                                                                            } else {
                                                                                                                if false {
                                                                                                                    var v312 string = ("\n\t" + v120)
                                                                                                                    _ = v312
                                                                                                                }
                                                                                                            }
                                                                                                            if (65535 >= uint16(v74)) {
                                                                                                                fmt.Println(v119)
                                                                                                                v74 = (v74 | v74)
                                                                                                            } else {
                                                                                                                if len(v5) >= 2 {
                                                                                                                    v5 = v5[1:len(v5)]
                                                                                                                }
                                                                                                                var v313 bool = false
                                                                                                                _ = v313
                                                                                                            }
                                                                                                            _ = v296
                                                                                                        }
                                                                                                        if true {
                                                                                                            switch ("\n\t" + ("a" + v119)) {
                                                                                                            case "foo bar":
                                                                                                                if false {
                                                                                                                    var v314 uint8 = 255
                                                                                                                    v315 := GenS17{F0: true, F1: 1, F2: (!(v240 == (v240 / ((v240) | 1))))}
                                                                                                                    _ = v314
                                                                                                                    _ = v315
                                                                                                                } else {
                                                                                                                    v316 := v1[2147483647]
                                                                                                                    v317 := GenS40{F0: 27, F1: 0, F2: (v6 * (v6 & v6))}
                                                                                                                    _ = v316
                                                                                                                    _ = v317
                                                                                                                }
                                                                                                            case "\n\t":
                                                                                                                v2[fi84] = v6
                                                                                                            case "":
                                                                                                                if (-10 == -35) {
                                                                                                                    var v318 uint64 = 2
                                                                                                                    var v319 uint32 = genFn0(18446744073709551615, 1000, (v74 >> ((65535) & 63)))
                                                                                                                    _ = v318
                                                                                                                    _ = v319
                                                                                                                } else {
                                                                                                                    var v320 int32 = 42
                                                                                                                    if (1 != (^(fi73))) {
                                                                                                                        if (!((v320 == (v320 << ((18446744073709551615) & 63))) && (1 > 0))) {
                                                                                                                            v120 = (v120 + "foo bar")
                                                                                                                            var v321 int32 = v320
                                                                                                                            v322 := GenS41{F0: 4294967295, F1: 1}
                                                                                                                            _ = v321
                                                                                                                            _ = v322
                                                                                                                        } else {
                                                                                                                            v323 := GenS3{F0: -109, F1: (v240 | v240), F2: v74}
                                                                                                                            v324 := GenS42{F0: "a", F1: 18446744073709551615, F2: (fi73 ^ fi84)}
                                                                                                                            fmt.Println(v119)
                                                                                                                            _ = v323
                                                                                                                            _ = v324
                                                                                                                        }
                                                                                                                        var v325 uint16 = 806
                                                                                                                        v320 = 1
                                                                                                                        _ = v325
                                                                                                                    } else {
                                                                                                                        var v326 uint32 = v6
                                                                                                                        v327 := GenS19{F0: true, F1: true}
                                                                                                                        _ = v326
                                                                                                                        _ = v327
                                                                                                                    }
                                                                                                                    v328 := GenS36{F0: fi73, F1: (v320 + v320), F2: "\n\t"}
                                                                                                                    _ = v320
                                                                                                                    _ = v328
                                                                                                                }
                                                                                                            }
                                                                                                        } else {
                                                                                                            switch fi73 {
                                                                                                            case 572:
                                                                                                                if false {
                                                                                                                    v119 = "\n\t"
                                                                                                                } else {
                                                                                                                    if (-981 != -2147483648) {
                                                                                                                        v74 = v74
                                                                                                                        v329 := GenS24{F0: (v6 / ((42) | 1))}
                                                                                                                        fmt.Println(fi73, v240, v74)
                                                                                                                        if (42 <= 42) {
                                                                                                                            var v330 uint8 = 2
                                                                                                                            v331 := GenS43{F0: 4294967295, F1: false, F2: genFn0(4294967295, 981, v74)}
                                                                                                                            _ = v330
                                                                                                                            _ = v331
                                                                                                                        } else {
                                                                                                                            if len(v5) > 0 {
                                                                                                                                v5[0] = v119
                                                                                                                            }
                                                                                                                            if len(v3) > 0 {
                                                                                                                                v3[0] = (v6 / ((v6) | 1))
                                                                                                                            }
                                                                                                                            fmt.Println(v6, v240)
                                                                                                                            v332 := GenS44{F0: v119}
                                                                                                                            _ = v332
                                                                                                                        }
                                                                                                                        v74 -= v74
                                                                                                                        _ = v329
                                                                                                                    } else {
                                                                                                                        v6 -= (^(v6))
                                                                                                                        var v333 int8 = 40
                                                                                                                        _ = v333
                                                                                                                    }
                                                                                                                    var v334 uint64 = v74
                                                                                                                    _ = v334
                                                                                                                }
                                                                                                            }
                                                                                                        }
                                                                                                    } else {
                                                                                                        var v335 string = v120
                                                                                                        v336 := GenS45{F0: 555, F1: ((("\n\t" + "\n\t") + v119) > v335), F2: false}
                                                                                                        _ = v335
                                                                                                        _ = v336
                                                                                                    }
                                                                                                }
                                                                                                var v337 float32 = -2.25
                                                                                                v74 -= 2
                                                                                                var v338 uint16 = 374
                                                                                                fmt.Println(v119, v337, fi84)
                                                                                                var v339 int16 = 181
                                                                                                _ = v337
                                                                                                _ = v338
                                                                                                _ = v339
                                                                                            }
                                                                                        case "":
                                                                                            v240 += v240
                                                                                        case "hello":
                                                                                            var v340 int8 = 100
                                                                                            _ = v340
                                                                                        }
                                                                                        switch (fi73 * (fi73 * 2147483647)) {
                                                                                        case -856:
                                                                                            var v341 int16 = 716
                                                                                            _ = v341
                                                                                        case 2147483647:
                                                                                            fmt.Println(v119)
                                                                                        case 128:
                                                                                            v342 := GenS46{F0: 128}
                                                                                            _ = v342
                                                                                        default:
                                                                                            var v343 int = fi84
                                                                                            _ = v343
                                                                                        }
                                                                                    }
                                                                                }
                                                                            default:
                                                                                switch (v119 + "\n\t") {
                                                                                case "hello":
                                                                                    v344 := GenS47{F0: 1, F1: -921, F2: v120}
                                                                                    _ = v344
                                                                                case "":
                                                                                    v119 = (v119 + v119)
                                                                                    if true {
                                                                                        v345 := GenS21{F0: 249}
                                                                                        fmt.Println(v74)
                                                                                        _ = v345
                                                                                    }
                                                                                    switch "" {
                                                                                    case "":
                                                                                        fmt.Println(v119, v119)
                                                                                    case "a":
                                                                                        var v346 bool = (false || (32767 != 302))
                                                                                        v347 := GenS34{F0: 70}
                                                                                        _ = v346
                                                                                        _ = v347
                                                                                    case "foo bar":
                                                                                        v1[-82] = (v119 + v119)
                                                                                    default:
                                                                                        var v348 uint64 = 2
                                                                                        _ = v348
                                                                                    }
                                                                                    if (((("" + v120) + ("hello" + v119)) + "") < v119) {
                                                                                        var v349 uint = 4294967295
                                                                                        _ = v349
                                                                                    }
                                                                                    switch fi73 {
                                                                                    case -299:
                                                                                        fmt.Println(v240, fi84, v120)
                                                                                    }
                                                                                default:
                                                                                    if (!(1 < 42)) {
                                                                                        switch fi84 {
                                                                                        case 32768:
                                                                                            v350 := GenS1{F0: (v119 + v120), F1: false, F2: genFn1(44, 1, (false && ((v74 << ((18446744073709551615) & 63)) >= (v74 >> ((18446744073709551615) & 63)))))}
                                                                                            _ = v350
                                                                                        }
                                                                                        v6 = 576
                                                                                        if ((v6 >= genFn1(0, 100, ((v74 - v74) > v74))) && (false && false)) {
                                                                                            v74 += v74
                                                                                            fmt.Println(v6, v119)
                                                                                        } else {
                                                                                            v120 = v120
                                                                                            var v351 string = v120
                                                                                            _ = v351
                                                                                        }
                                                                                    }
                                                                                }
                                                                            }
                                                                        }
                                                                        _ = v243
                                                                    } else {
                                                                        v74 -= (v74 - v74)
                                                                    }
                                                                default:
                                                                    var v352 int = -420
                                                                    _ = v352
                                                                }
                                                            default:
                                                                v240 -= v240
                                                            }
                                                            var v353 int64 = -263
                                                            v354 := GenS47{F0: 128, F1: 32767, F2: "foo bar"}
                                                            _ = v240
                                                            _ = v353
                                                            _ = v354
                                                        }
                                                    }
                                                    v74 -= (v74 & 2)
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                    _ = len(v3)
                                                    _ = v119
                                                    _ = v120
                                                } else {
                                                    switch "hello" {
                                                    case "foo bar":
                                                        var v355 uint32 = 211
                                                        _ = v355
                                                    case "\n\t":
                                                        v356 := GenS48{F0: 42}
                                                        _ = v356
                                                    default:
                                                        if false {
                                                            v357 := GenS49{F0: ("" + "foo bar")}
                                                            v2[fi73] = v6
                                                            _ = v357
                                                        }
                                                    }
                                                    v358 := GenS50{F0: (^(v74)), F1: true, F2: 1}
                                                    _ = v358
                                                }
                                            case 2147483647:
                                                switch (fi84 % ((fi73) | 1)) {
                                                case 127:
                                                    v359 := GenS51{F0: 1, F1: 255, F2: 42}
                                                    _ = v359
                                                }
                                            }
                                            v6 += v6
                                        }
                                    }
                                }
                                if (false && true) {
                                    var v360 int32 = 523
                                    fmt.Println(fi73)
                                    v74 = (^(v74))
                                    _ = v360
                                }
                                _ = fi84
                            }
                        }
                        v74 += v74
                        _ = v74
                    } else {
                        var v361 float32 = (0.0 / -2.25)
                        var v362 string = "foo bar"
                        _ = v361
                        _ = v362
                    }
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    _ = fi73
                }
                for fi363 := 0; fi363 < 8; fi363++ {
                    v364 := GenS1{F0: ("\n\t" + "\n\t"), F1: (!((fi363 - fi363) >= fi363)), F2: v6}
                    v6 = genFn0(4294967295, 1000, 2)
                    _ = fi363
                    _ = v364
                }
            }
        default:
            var v365 int16 = 770
            _ = v365
        }
    }
    switch "hello" {
    case "a":
        for fi366 := 0; fi366 < 6; fi366++ {
            v367 := GenS33{F0: (false || true), F1: fi366}
            for fi368 := 0; fi368 < 20; fi368++ {
                v6 -= v6
                _ = fi368
            }
            _ = fi366
            _ = v367
        }
    case "\n\t":
        for fi369 := 0; fi369 < 5; fi369++ {
            v370 := GenS12{F0: 881}
            _ = fi369
            _ = v370
        }
    default:
        if ("" > "hello") {
            if true {
                var v371 int16 = 873
                v371 = v371
                fmt.Println(v6, v371)
                _ = v371
            } else {
                v6 += (^(v6))
                var v372 float64 = (1e10 - 3.14159)
                v6 += (v6 & v6)
                _ = v372
            }
            var v373 uint32 = v6
            _ = v373
        } else {
            v374 := v1[128]
            _ = v374
        }
    }
    for fi375 := 0; fi375 < 14; fi375++ {
        if false {
            v376 := GenS12{F0: 2}
            _ = v376
        } else {
            for fi377 := 0; fi377 < 6; fi377++ {
                if (!true) {
                    if false {
                        v6 += (^(v6))
                    } else {
                        v378 := GenS26{F0: 127}
                        v379 := GenS52{F0: 2, F1: 42, F2: fi375}
                        var v380 float64 = ((float64(1.5) - 1e10) / (((3.14159 * -2.25) - 1e10) - ((1e10 / 1e10) + 3.14159)))
                        _ = v378
                        _ = v379
                        _ = v380
                    }
                    var v381 uint16 = 255
                    v381 -= (v381 << ((4294967295) & 63))
                    if (!(false && false)) {
                        if ((("\n\t" + "hello") + ("foo bar" + "\n\t")) >= "a") {
                            v382 := GenS53{F0: 42, F1: 9223372036854775807, F2: 1}
                            _ = len(v1)
                            _ = v382
                        } else {
                            var v383 uint16 = 255
                            var v384 uint32 = 0
                            v385 := GenS21{F0: 60}
                            v2[876] = v384
                            _ = v383
                            _ = v384
                            _ = v385
                        }
                    }
                    var v386 int64 = -9223372036854775808
                    if (((0.0 - (1e10 * -2.25)) + -2.25) > (0.0 / 3.14159)) {
                        v386 += v386
                        v386 = 0
                    }
                    _ = v381
                    _ = v386
                }
                _ = fi377
            }
        }
        switch "a" {
        case "a":
            v6 -= v6
        default:
            var v387 int = (fi375 << ((65535) & 63))
            _ = v387
        }
        _ = fi375
    }
    var v388 float64 = 3.14159
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v388:", v388)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v388
    _ = v4
    _ = v5
    _ = v6
    _ = v70
}
