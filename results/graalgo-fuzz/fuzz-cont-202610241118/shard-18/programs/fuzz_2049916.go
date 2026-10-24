package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64 }
type GenS2 struct { F0 uint32; F1 uint16; F2 uint64 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint8; F1 bool; F2 bool }
type GenS5 struct { F0 uint8; F1 uint; F2 uint64 }
type GenS6 struct { F0 int8; F1 uint8 }
type GenS7 struct { F0 uint64; F1 int32 }
type GenS8 struct { F0 uint16; F1 uint16; F2 string }
type GenS9 struct { F0 int32; F1 uint16; F2 int8 }
type GenS10 struct { F0 uint; F1 int8 }
type GenS11 struct { F0 int16 }
type GenS12 struct { F0 uint32; F1 int8; F2 uint8 }

func main() {
    v1 := []uint32{1, 65535, 2147483648, 2147483647}
    v2 := []int{}
    var v3 int8 = 42
    if true {
        for fi4 := 0; fi4 < 14; fi4++ {
            switch (fi4 * (fi4 / ((1) | 1))) {
            case 42:
                v2 = append(v2, (fi4 + fi4))
            case -2147483648:
                v3 += v3
            default:
                if (0.0 >= 0.0) {
                    v5 := GenS1{F0: 255, F1: 0}
                    _ = v5
                }
            }
            switch "foo bar" {
            case "foo bar":
                for fi6 := 0; fi6 < 10; fi6++ {
                    switch ("foo bar" + (("\n\t" + ("\n\t" + "a")) + "\n\t")) {
                    case "hello":
                        v7 := GenS2{F0: 0, F1: 2, F2: 0}
                        _ = v7
                    case "a":
                        v3 = v3
                    }
                    var v8 int32 = 795
                    switch ("hello" + "hello") {
                    case "foo bar":
                        v9 := GenS2{F0: 65535, F1: 32767, F2: 18446744073709551615}
                        _ = v9
                    case "":
                        if true {
                            if len(v2) > 0 {
                                v2[0] = 9223372036854775807
                            }
                        } else {
                            fmt.Println(v8, fi6, v3)
                            if len(v2) > 0 {
                                v2[0] = fi4
                            }
                        }
                    case "hello":
                        var v10 uint = 18446744073709551615
                        _ = v10
                    default:
                        v3 -= v3
                    }
                    switch fi4 {
                    case -9223372036854775808:
                        var v11 int32 = (v8 ^ (v8 & (v8 & v8)))
                        _ = v11
                    default:
                        _ = len(v1)
                    }
                    if (65535 > 0) {
                        if len(v2) > 0 {
                            v2[0] = fi6
                        }
                        if (!(406 <= -1)) {
                            v12 := GenS1{F0: 255, F1: 2}
                            switch "foo bar" {
                            case "":
                                fmt.Println(v8, v8)
                            default:
                                v2 = append(v2, (fi6 << ((4294967295) & 63)))
                            }
                            _ = v12
                        } else {
                            switch "hello" {
                            case "a":
                                if (-1 == -32768) {
                                    if true {
                                        switch fi4 {
                                        case 876:
                                            v8 += v8
                                        case 9223372036854775807:
                                            var v13 bool = (("\n\t" + "\n\t") < "")
                                            _ = v13
                                        default:
                                            v2 = append(v2, fi4)
                                        }
                                    } else {
                                        fmt.Println(v3, fi4, v8)
                                        if false {
                                            var v14 bool = (2 == 18446744073709551615)
                                            _ = v14
                                        }
                                        v8 = 42
                                    }
                                } else {
                                    switch ("a" + "foo bar") {
                                    case "":
                                        var v15 uint = 557
                                        _ = v15
                                    case "hello":
                                        var v16 uint8 = 253
                                        _ = v16
                                    }
                                    v8 = v8
                                }
                                var v17 bool = ("" < "")
                                v8 += v8
                                v2 = append(v2, fi6)
                                var v18 float32 = 3.14159
                                _ = v17
                                _ = v18
                            default:
                                switch (("" + "") + "a") {
                                case "a":
                                    var v19 string = "\n\t"
                                    _ = v19
                                default:
                                    v20 := GenS2{F0: 908, F1: 2, F2: 18446744073709551615}
                                    _ = v20
                                }
                            }
                            v21 := GenS3{F0: -610}
                            v22 := GenS4{F0: 255, F1: false, F2: false}
                            _ = v21
                            _ = v22
                        }
                    } else {
                        v23 := GenS4{F0: 127, F1: (663 != 18446744073709551615), F2: true}
                        var v24 bool = (!true)
                        _ = v23
                        _ = v24
                    }
                    fmt.Println(v8)
                    if (false && false) {
                        v2 = append(v2, 32768)
                        var v25 uint8 = 128
                        switch (fi6 + (fi6 << ((uint(v25)) & 63))) {
                        case 9223372036854775807:
                            var v26 uint16 = 443
                            _ = v26
                        }
                        _ = v25
                    } else {
                        var v27 uint16 = 32767
                        var v28 float32 = (((1e10 * 1.5) / 1.0) * 3.14159)
                        if (!(0 >= 107)) {
                            v8 += v8
                            v29 := GenS3{F0: 0}
                            var v30 int16 = 32767
                            _ = v29
                            _ = v30
                        } else {
                            v3 += (-(v3))
                        }
                        v8 = -219
                        v31 := GenS2{F0: 252, F1: 308, F2: 4294967295}
                        v32 := GenS5{F0: 0, F1: uint(v27), F2: 819}
                        v33 := GenS6{F0: -128, F1: 42}
                        v34 := GenS6{F0: (v3 ^ (v3 - v3)), F1: 255}
                        _ = v27
                        _ = v28
                        _ = v31
                        _ = v32
                        _ = v33
                        _ = v34
                    }
                    _ = fi6
                    _ = v8
                }
            default:
                v3 = v3
            }
            v3 += v3
            _ = fi4
        }
        if false {
            for fi35 := 0; fi35 < 6; fi35++ {
                if (2 <= 18446744073709551615) {
                    v3 = (v3 + (v3 * (-(v3))))
                    for fi36 := 0; fi36 < 4; fi36++ {
                        var v37 uint8 = 1
                        if (-1 < 1) {
                            var v38 int32 = -886
                            switch 0 {
                            case -9223372036854775808:
                                v38 = (-(v38))
                            case 4294967295:
                                _ = len(v2)
                            case 2147483647:
                                var v39 bool = true
                                _ = v39
                            }
                            var v40 int16 = 1000
                            _ = v38
                            _ = v40
                        } else {
                            switch fi36 {
                            case 0:
                                v2 = append(v2, (fi35 << ((65535) & 63)))
                            }
                            if (-2.25 == 3.14159) {
                                var v41 bool = false
                                _ = v41
                            } else {
                                switch "\n\t" {
                                case "hello":
                                    if ((fi36 >= 0) && (true || true)) {
                                        v3 -= v3
                                        var v42 uint32 = 0
                                        _ = v42
                                    } else {
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        v3 = (v3 - v3)
                                        var v43 int = 1
                                        if (true && false) {
                                            var v44 int16 = -463
                                            _ = len(v2)
                                            var v45 float32 = (-2.25 * 0.0)
                                            _ = v44
                                            _ = v45
                                        }
                                        var v46 bool = true
                                        _ = v43
                                        _ = v46
                                    }
                                case "\n\t":
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                }
                                _ = len(v1)
                            }
                        }
                        var v47 int64 = 42
                        _ = fi36
                        _ = v37
                        _ = v47
                    }
                } else {
                    for fi48 := 0; fi48 < 4; fi48++ {
                        switch "" {
                        case "hello":
                            if true {
                                var v49 string = (("hello" + "") + ("\n\t" + ("" + "")))
                                _ = v49
                            }
                        }
                        var v50 string = "a"
                        v51 := GenS3{F0: 41}
                        if (2147483647 >= 42) {
                            fmt.Println(fi48, fi35, fi48)
                        }
                        v1 = append(v1, 1)
                        v52 := GenS5{F0: 1, F1: 42, F2: 0}
                        _ = fi48
                        _ = v50
                        _ = v51
                        _ = v52
                    }
                }
                switch "foo bar" {
                case "foo bar":
                    var v53 uint8 = 133
                    _ = v53
                case "hello":
                    switch "foo bar" {
                    case "foo bar":
                        for fi54 := 0; fi54 < 15; fi54++ {
                            if ((((((3.14159 / -2.25) / ((0.1 - 3.14159) / 1.0)) * 0.0) - (1.5 - (1.5 + -2.25))) + -2.25) < 0.0) {
                                var v55 int64 = 111
                                _ = v55
                            }
                            if true {
                                v56 := GenS2{F0: 65535, F1: 1, F2: 0}
                                if true {
                                    fmt.Println(fi54, v3, v3)
                                }
                                _ = v56
                            } else {
                                if (fi54 < fi54) {
                                    fmt.Println(v3, fi54)
                                }
                                if false {
                                    if ("\n\t" > "foo bar") {
                                        if false {
                                            var v57 uint64 = 18446744073709551615
                                            var v58 uint = 771
                                            _ = v57
                                            _ = v58
                                        } else {
                                            var v59 uint64 = 380
                                            _ = v59
                                        }
                                        v60 := GenS3{F0: -672}
                                        var v61 uint16 = 255
                                        _ = v60
                                        _ = v61
                                    }
                                    v3 = v3
                                } else {
                                    v62 := GenS2{F0: 4294967295, F1: 765, F2: 0}
                                    _ = len(v2)
                                    v3 = v3
                                    _ = v62
                                }
                            }
                            if false {
                                if (!(!false)) {
                                    v63 := GenS3{F0: 1000}
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    if (1 < -164) {
                                        if false {
                                            v64 := GenS7{F0: 131, F1: 0}
                                            v2 = append(v2, (fi54 * fi35))
                                            _ = v64
                                        } else {
                                            v2 = append(v2, (fi54 >> ((65535) & 63)))
                                            v2 = append(v2, (fi35 + fi54))
                                            v3 -= v3
                                        }
                                    }
                                    _ = v63
                                } else {
                                    var v65 float32 = -2.25
                                    _ = v65
                                }
                                var v66 int64 = 0
                                _ = v66
                            } else {
                                if (1 >= 128) {
                                    var v67 float32 = float32((3.14159 + (1.5 - (float64(1.5) - 3.14159))))
                                    v67 = v67
                                    _ = v67
                                }
                            }
                            _ = fi54
                        }
                    }
                case "":
                    v3 = v3
                }
                if (false || true) {
                    v2 = append(v2, fi35)
                    v3 = (-(v3))
                    v3 = 88
                }
                _ = fi35
            }
            v3 = v3
            v68 := GenS8{F0: 928, F1: 496, F2: (("foo bar" + "") + (("foo bar" + ("hello" + "foo bar")) + "hello"))}
            _ = v68
        } else {
            if true {
                if true {
                    var v69 int64 = -1
                    _ = v69
                }
                v70 := GenS9{F0: 0, F1: 65535, F2: (v3 ^ 64)}
                _ = v70
            } else {
                v3 -= v3
                v71 := GenS9{F0: -2147483648, F1: 1, F2: (v3 * (v3 | (v3 | (v3 ^ 1))))}
                if (1 == 310) {
                    var v72 uint64 = 1
                    v73 := GenS10{F0: 42, F1: v3}
                    _ = v72
                    _ = v73
                }
                _ = v71
            }
        }
    }
    var v74 int = 742
    var v75 string = "\n\t"
    var v76 int = v74
    if false {
        switch (-(v76)) {
        case 4294967295:
            switch v75 {
            case "hello":
                var v77 int = -2147483648
                _ = v77
            }
        case 9223372036854775807:
            v78 := GenS11{F0: 32767}
            _ = v78
        }
        v79 := GenS12{F0: 658, F1: v3, F2: 2}
        switch "" {
        case "":
            if (true || (!(230 == 2))) {
                var v80 int8 = (-(v3))
                switch (v76 * v74) {
                case 0:
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                case -1:
                    var v81 uint8 = 2
                    _ = v81
                case 127:
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                }
                _ = v80
            }
        case "hello":
            var v82 float64 = -2.25
            _ = v82
        case "a":
            v83 := GenS5{F0: 204, F1: 255, F2: 4294967295}
            _ = v83
        default:
            v84 := GenS9{F0: 42, F1: 769, F2: -1}
            _ = v84
        }
        _ = v79
    } else {
        v3 += (v3 << ((4294967295) & 63))
    }
    if false {
        fmt.Println(v76, v3, v74)
        var v85 int16 = 32767
        var v86 uint64 = 1
        _ = v85
        _ = v86
    } else {
        v87 := GenS10{F0: 65535, F1: 127}
        var v88 uint32 = 1
        _ = v87
        _ = v88
    }
    switch (("" + "\n\t") + v75) {
    case "foo bar":
        var v89 uint64 = 153
        _ = v89
    }
    var v90 bool = (1e10 == float64((1e10 / 1.0)))
    if (v74 <= (v76 * -200)) {
        fmt.Println(v90, v3)
        _ = len(v2)
    }
    fmt.Println(v76, v75, v74)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v74:", v74)
    fmt.Println("v75:", v75)
    fmt.Println("v76:", v76)
    fmt.Println("v90:", v90)
    _ = v1
    _ = v2
    _ = v3
    _ = v74
    _ = v75
    _ = v76
    _ = v90
}
