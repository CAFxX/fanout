package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint64; F1 int64 }
type GenS3 struct { F0 uint32; F1 int8; F2 int64 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int32; F1 string }
type GenS7 struct { F0 int32; F1 int32 }
type GenS8 struct { F0 uint32; F1 int8; F2 uint8 }
type GenS9 struct { F0 uint32; F1 int8; F2 uint }
type GenS10 struct { F0 int16; F1 int }
type GenS11 struct { F0 int32 }
type GenS12 struct { F0 uint64; F1 int32 }
type GenS13 struct { F0 uint32 }
func genFn0(p0 int16, p1 int32) int16 {
    var v5 uint32 = 2
    v6 := []int{-816, -1, -815, 32768}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 42
}
func genFn1(p0 int32, p1 uint8) uint {
    p1 = (p1 >> ((4294967295) & 63))
    var v7 int32 = (p0 ^ -1)
    _ = p0
    _ = p1
    _ = v7
    return 2
}

func main() {
    v1 := []uint32{2147483647, 2147483648, 65535}
    v2 := []int{-9223372036854775808, 4294967295, 127, -383}
    v3 := make(map[string]string)
    var v4 uint64 = 42
    if (true && ((255 == 0) && (42 < 261))) {
        v4 -= v4
        fmt.Println(v4)
        if false {
            var v8 string = "hello"
            for fi9 := 0; fi9 < 9; fi9++ {
                v10 := GenS1{F0: 18}
                var v11 int32 = -2147483648
                _ = len(v3)
                _ = fi9
                _ = v10
                _ = v11
            }
            v12 := GenS2{F0: (v4 - 42), F1: -9223372036854775808}
            _ = v12
            _ = v8
        }
    } else {
        var v13 uint32 = 0
        if ((("hello" + "\n\t") + "hello") != ("\n\t" + "\n\t")) {
            fmt.Println(v13, v4)
        } else {
            for fi14 := 0; fi14 < 19; fi14++ {
                v4 -= 4294967295
                v15 := GenS2{F0: 2, F1: -478}
                if (!false) {
                    v4 = (v4 * v4)
                } else {
                    if false {
                        v16 := GenS2{F0: (^(v4)), F1: 1}
                        for fi17 := 0; fi17 < 20; fi17++ {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            _ = fi17
                        }
                        fmt.Println(fi14, fi14, v4)
                        _ = v16
                    }
                    v18 := GenS2{F0: v4, F1: -1}
                    if (128 <= 128) {
                        fmt.Println(v13, v13, fi14)
                    } else {
                        if true {
                            var v19 int32 = 0
                            _ = v19
                        } else {
                            if true {
                                switch fi14 {
                                case 4294967295:
                                    if false {
                                        v13 += 4294967295
                                        fmt.Println(fi14, v4)
                                        for fi20 := 0; fi20 < 19; fi20++ {
                                            v21 := GenS3{F0: 42, F1: 0, F2: 1}
                                            if (2147483647 != 662) {
                                                v3["\n\t"] = "foo bar"
                                                v4 += v4
                                                switch (^(fi14)) {
                                                case 127:
                                                    fmt.Println(v4, v13, v13)
                                                case 0:
                                                    v3[("foo bar" + "hello")] = ("foo bar" + "hello")
                                                default:
                                                    var v22 uint64 = v4
                                                    _ = v22
                                                }
                                            }
                                            _ = fi20
                                            _ = v21
                                        }
                                    } else {
                                        if (!true) {
                                            v23 := GenS4{F0: 32767}
                                            _ = len(v2)
                                            _ = v23
                                        }
                                        if (((^(v13)) == (^(v13))) || false) {
                                            if ((genFn0(-32768, 2147483647) == 32767) || true) {
                                                v24 := GenS2{F0: v4, F1: 9223372036854775807}
                                                if true {
                                                    var v25 uint32 = (v13 + v13)
                                                    switch ("a" + "\n\t") {
                                                    case "":
                                                        v26 := GenS3{F0: (v13 << ((255) & 63)), F1: 42, F2: 264}
                                                        _ = v26
                                                    case "foo bar":
                                                        var v27 uint64 = (v4 + (^(v4)))
                                                        _ = v27
                                                    case "a":
                                                        var v28 uint64 = 4294967295
                                                        _ = v28
                                                    }
                                                    switch ("\n\t" + "hello") {
                                                    case "hello":
                                                        for fi29 := 0; fi29 < 12; fi29++ {
                                                            v2 = append(v2, fi29)
                                                            v25 = (v25 + (^(v13)))
                                                            switch "a" {
                                                            case "a":
                                                                var v30 uint8 = 128
                                                                _ = v30
                                                            case "hello":
                                                                v4 = 18446744073709551615
                                                            case "foo bar":
                                                                v13 -= (v25 / ((224) | 1))
                                                            }
                                                            _ = fi29
                                                        }
                                                    case "":
                                                        v31 := GenS3{F0: 1, F1: 0, F2: -533}
                                                        _ = v31
                                                    case "\n\t":
                                                        if false {
                                                            var v32 int8 = 41
                                                            for fi33 := 0; fi33 < 7; fi33++ {
                                                                _ = len(v3)
                                                                var v34 float64 = 0.0
                                                                _ = fi33
                                                                _ = v34
                                                            }
                                                            fmt.Println(v25, v4)
                                                            _ = v32
                                                        } else {
                                                            v35 := GenS5{F0: 0}
                                                            fmt.Println(v13, v25, v4)
                                                            _ = v35
                                                        }
                                                    }
                                                    _ = v25
                                                }
                                                _ = v24
                                            } else {
                                                v36 := GenS6{F0: -169, F1: ("hello" + "hello")}
                                                if (9223372036854775807 >= -9223372036854775808) {
                                                    fmt.Println(v13, fi14, fi14)
                                                    fmt.Println(v4)
                                                } else {
                                                    switch ("" + ("a" + "\n\t")) {
                                                    case "a":
                                                        if len(v1) >= 2 {
                                                            v1 = v1[1:len(v1)]
                                                        }
                                                    case "hello":
                                                        v37 := GenS4{F0: fi14}
                                                        _ = v37
                                                    case "":
                                                        v13 = 2
                                                    }
                                                    for fi38 := 0; fi38 < 14; fi38++ {
                                                        delete(v3, ("foo bar" + "\n\t"))
                                                        if len(v2) >= 2 {
                                                            v2 = v2[1:len(v2)]
                                                        }
                                                        _ = fi38
                                                    }
                                                }
                                                _ = v36
                                            }
                                            fmt.Println(v4, v4)
                                            var v39 uint16 = 42
                                            _ = v39
                                        } else {
                                            v13 = v13
                                            switch fi14 {
                                            case 32768:
                                                var v40 int16 = 99
                                                _ = v40
                                            default:
                                                var v41 int32 = -116
                                                _ = v41
                                            }
                                        }
                                    }
                                case 9223372036854775807:
                                    for fi42 := 0; fi42 < 18; fi42++ {
                                        var v43 int32 = 775
                                        _ = fi42
                                        _ = v43
                                    }
                                default:
                                    fmt.Println(v4, fi14)
                                }
                                v13 += v13
                            }
                        }
                    }
                    _ = v18
                }
                _ = fi14
                _ = v15
            }
            switch ("hello" + "\n\t") {
            case "":
                v3["\n\t"] = "a"
            case "a":
                fmt.Println(v13)
            case "foo bar":
                if true {
                    v44 := GenS4{F0: 32768}
                    _ = v44
                }
            }
        }
        if len(v1) > 0 {
            v1[0] = 65535
        }
        _ = v13
    }
    if len(v2) > 0 {
        v2[0] = 32767
    }
    var v45 int64 = 531
    switch -2147483648 {
    case 127:
        var v46 uint = 214
        _ = v46
    case -611:
        var v47 uint64 = v4
        _ = v47
    case 32768:
        if (v4 != (^(v4))) {
            v48 := GenS4{F0: 4294967295}
            if false {
                if false {
                    switch "" {
                    case "a":
                        var v49 float32 = 3.14159
                        _ = v49
                    case "foo bar":
                        for fi50 := 0; fi50 < 2; fi50++ {
                            v51 := GenS7{F0: 920, F1: -803}
                            v3["hello"] = ("foo bar" + "")
                            _ = fi50
                            _ = v51
                        }
                    }
                    v52 := GenS4{F0: 442}
                    _ = v52
                }
                for fi53 := 0; fi53 < 12; fi53++ {
                    switch ((("hello" + "foo bar") + "a") + ("a" + ((((("foo bar" + "foo bar") + "") + ("foo bar" + "a")) + "") + ("hello" + "foo bar")))) {
                    case "a":
                        fmt.Println(v45, v4)
                    default:
                        if true {
                            if ((v45 ^ v45) <= v45) {
                                v45 = v45
                            } else {
                                var v54 uint16 = 1
                                fmt.Println(v45, v45, v4)
                                var v55 uint64 = (v4 / (((^(v4))) | 1))
                                _ = v54
                                _ = v55
                            }
                            delete(v3, "\n\t")
                            if ((0.1 * (3.14159 * 0.1)) >= (3.14159 * 1e10)) {
                                v56 := GenS8{F0: uint32(v4), F1: -37, F2: 128}
                                fmt.Println(v45)
                                _ = v56
                            }
                        } else {
                            var v57 uint16 = 255
                            fmt.Println(fi53, v4)
                            _ = v57
                        }
                    }
                    for fi58 := 0; fi58 < 11; fi58++ {
                        v59 := GenS9{F0: 620, F1: 42, F2: 255}
                        v60 := GenS9{F0: 4294967295, F1: -128, F2: genFn1(1, 90)}
                        v45 = v45
                        if false {
                            var v61 int32 = -1
                            _ = v61
                        }
                        v62 := GenS10{F0: 32767, F1: (^(fi58))}
                        fmt.Println(fi53)
                        switch (("hello" + (("" + (("a" + ((("\n\t" + "\n\t") + "hello") + ("\n\t" + "a"))) + (("foo bar" + "hello") + ""))) + "hello")) + "foo bar") {
                        case "":
                            if true {
                                fmt.Println(v45, v4)
                                fmt.Println(fi58)
                                if len(v2) > 0 {
                                    v2[0] = 42
                                }
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            }
                        case "hello":
                            switch (^(fi58)) {
                            case -9223372036854775808:
                                v63 := GenS8{F0: 491, F1: 100, F2: 165}
                                _ = v63
                            default:
                                if ((fi53 * (fi53 % (((fi53 / ((fi58) | 1))) | 1))) == 1) {
                                    if (((v4 % ((v4) | 1)) == (v4 ^ v4)) || (2 == 810)) {
                                        var v64 uint8 = 2
                                        fmt.Println(v4, v4)
                                        var v65 int = fi53
                                        fmt.Println(v65, v45)
                                        v64 = v64
                                        v66 := GenS1{F0: 80}
                                        var v67 int8 = int8(v65)
                                        v3[("\n\t" + "foo bar")] = "a"
                                        _ = v64
                                        _ = v65
                                        _ = v66
                                        _ = v67
                                    }
                                    var v68 float64 = (((((0.0 / -2.25) - 3.14159) / -2.25) - (0.0 * 1.5)) + 0.0)
                                    _ = v68
                                }
                            }
                        }
                        if (v4 <= (v4 / (((v4 | (^(v4)))) | 1))) {
                            if (0 > 2147483648) {
                                var v69 int32 = -2147483648
                                if true {
                                    v70 := GenS8{F0: 42, F1: -1, F2: 255}
                                    v71 := GenS7{F0: (^(v69)), F1: -2147483648}
                                    switch fi58 {
                                    case -971:
                                        v45 += (v45 * (v45 >> ((615) & 63)))
                                    case -526:
                                        v45 += v45
                                        switch (fi58 ^ fi58) {
                                        case -1:
                                            v69 = -807
                                        }
                                        v3[("" + ("foo bar" + "hello"))] = "\n\t"
                                    case 128:
                                        var v72 string = ("\n\t" + (("foo bar" + "a") + "a"))
                                        _ = v72
                                    }
                                    _ = v70
                                    _ = v71
                                }
                                _ = len(v3)
                                _ = v69
                            }
                        } else {
                            v73 := GenS8{F0: 654, F1: 12, F2: 127}
                            v74 := v3[(("foo bar" + "foo bar") + ("\n\t" + ""))]
                            v75 := GenS11{F0: -2147483648}
                            _ = v73
                            _ = v74
                            _ = v75
                        }
                        var v76 string = ("" + "a")
                        v45 += (v45 + 42)
                        _ = fi58
                        _ = v59
                        _ = v60
                        _ = v62
                        _ = v76
                    }
                    fmt.Println(fi53, v45, v4)
                    _ = fi53
                }
            } else {
                var v77 string = ("" + "foo bar")
                _ = v77
            }
            var v78 int64 = v45
            _ = v48
            _ = v78
        }
    default:
        v45 -= (v45 >> ((845) & 63))
    }
    v79 := v3["a"]
    v80 := v3[v79]
    switch 42 {
    case -2147483648:
        fmt.Println(v4, v79)
    case -559:
        fmt.Println(v4, v79)
    default:
        if ((-2.25 - 1.5) >= float32(0.0)) {
            var v81 uint = 624
            fmt.Println(v45)
            fmt.Println(v80)
            _ = v81
        }
    }
    v82 := GenS4{F0: -91}
    v3["foo bar"] = v79
    if (65535 != 2) {
        v83 := GenS9{F0: 65535, F1: -37, F2: 42}
        _ = v83
    } else {
        if len(v1) > 0 {
            v1[0] = 65535
        }
        switch 2147483647 {
        case 827:
            if true {
                switch 2147483647 {
                case -590:
                    var v84 int16 = 1
                    _ = v84
                default:
                    var v85 string = "\n\t"
                    _ = v85
                }
                for fi86 := 0; fi86 < 17; fi86++ {
                    v87 := GenS12{F0: v4, F1: -1}
                    var v88 uint8 = 2
                    switch (fi86 << ((162) & 63)) {
                    case 914:
                        if (0 <= 65535) {
                            if (false && true) {
                                if true {
                                    v4 -= v4
                                } else {
                                    v89 := GenS11{F0: 1}
                                    _ = v89
                                }
                                switch (("\n\t" + ((v80 + v80) + ("hello" + v79))) + ("foo bar" + v80)) {
                                case "hello":
                                    var v90 int16 = 715
                                    _ = v90
                                }
                                v3[((v79 + (v79 + "hello")) + v80)] = "\n\t"
                            }
                            for fi91 := 0; fi91 < 15; fi91++ {
                                _ = len(v1)
                                var v92 uint64 = v4
                                if (!(v79 > ("" + "hello"))) {
                                    v93 := GenS8{F0: 0, F1: 100, F2: (v88 % (((v88 % ((153) | 1))) | 1))}
                                    fmt.Println(v88, fi91, fi91)
                                    fmt.Println(fi91, v88)
                                    _ = v93
                                }
                                fmt.Println(v88, v45)
                                _ = fi91
                                _ = v92
                            }
                        } else {
                            var v94 float32 = float32(float64(0.0))
                            _ = v94
                        }
                    case 485:
                        if true {
                            v88 += (v88 * (v88 >> ((952) & 63)))
                        }
                    }
                    _ = fi86
                    _ = v87
                    _ = v88
                }
            } else {
                v1 = append(v1, 1)
            }
        case 0:
            v95 := GenS13{F0: 2}
            _ = v95
        case -903:
            var v96 int16 = -66
            _ = v96
        }
        for fi97 := 0; fi97 < 6; fi97++ {
            var v98 bool = false
            _ = fi97
            _ = v98
        }
    }
    for fi99 := 0; fi99 < 15; fi99++ {
        var v100 int8 = -1
        var v101 float64 = 1e10
        _ = fi99
        _ = v100
        _ = v101
    }
    v79 = "a"
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v79:", v79)
    fmt.Println("v80:", v80)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v79
    _ = v80
    _ = v82
}
