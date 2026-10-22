package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint; F2 int32 }
type GenS2 struct { F0 int8; F1 uint8 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint; F1 int16; F2 uint16 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 int32; F1 uint64 }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 uint32; F1 int32 }
type GenS10 struct { F0 uint16; F1 int64; F2 int8 }
type GenS11 struct { F0 int16; F1 bool }
type GenS12 struct { F0 uint16; F1 uint64; F2 int8 }
type GenS13 struct { F0 int; F1 bool }
type GenS14 struct { F0 int64; F1 int64 }
type GenS15 struct { F0 uint8; F1 int64; F2 uint64 }
type GenS16 struct { F0 int32 }
type GenS17 struct { F0 uint8; F1 int64 }
type GenS18 struct { F0 string; F1 uint64; F2 int32 }
type GenS19 struct { F0 int; F1 uint64 }
type GenS20 struct { F0 uint16 }
type GenS21 struct { F0 int32 }
type GenS22 struct { F0 string }
type GenS23 struct { F0 int16; F1 bool; F2 string }
type GenS24 struct { F0 bool; F1 int8 }
type GenS25 struct { F0 uint8 }
type GenS26 struct { F0 int64; F1 string; F2 string }
type GenS27 struct { F0 int8; F1 uint8; F2 string }
func genFn0() uint16 {
    v7 := []int{128, -2147483648}
    switch "" {
    case "foo bar":
        v7 = append(v7, 9223372036854775807)
    case "":
        var v8 string = (("" + "a") + (("hello" + "hello") + "\n\t"))
        _ = v8
    }
    _ = v7
    return 930
}
func genFn1(p0 uint8) bool {
    switch -305 {
    case -2147483648:
        fmt.Println(p0)
    }
    v9 := GenS1{F0: 2, F1: 42, F2: 932}
    var v10 uint16 = genFn0()
    var v11 int8 = 72
    var v12 uint = uint(v10)
    _ = p0
    _ = v10
    _ = v11
    _ = v12
    _ = v9
    return (0.0 <= 3.14159)
}

func main() {
    v1 := []string{"a", "hello", "\n\t"}
    v2 := make(map[uint32]uint32)
    var v3 uint8 = 2
    v4 := []string{"", "\n\t", "", ("a" + "\n\t")}
    var v5 uint16 = 82
    var v6 float64 = float64((float32(float64(0.1)) / -2.25))
    v3 -= 255
    v3 = v3
    var v13 string = ("a" + "hello")
    fmt.Println(v5, v13, v6)
    _ = len(v1)
    if true {
        if true {
            for fi14 := 0; fi14 < 10; fi14++ {
                switch "foo bar" {
                case "":
                    if (false && (127 <= 1)) {
                        for fi15 := 0; fi15 < 10; fi15++ {
                            if false {
                                v16 := GenS2{F0: -128, F1: (v3 & v3)}
                                v17 := v2[2147483647]
                                v18 := GenS1{F0: 27, F1: 65535, F2: 1}
                                _ = v16
                                _ = v17
                                _ = v18
                            } else {
                                var v19 int32 = 92
                                fmt.Println(v5, v5, v6)
                                var v20 string = "foo bar"
                                _ = v19
                                _ = v20
                            }
                            fmt.Println(v6, fi15, fi14)
                            switch ((v13 + ("foo bar" + "a")) + "a") {
                            case "foo bar":
                                v21 := GenS1{F0: 2, F1: 1, F2: 1}
                                _ = v21
                            case "hello":
                                if ((0.1 / 3.14159) > v6) {
                                    v13 = v13
                                    fmt.Println(v6, v3)
                                    fmt.Println(v5, v6)
                                    _ = len(v2)
                                    var v22 int64 = -9223372036854775808
                                    _ = v22
                                } else {
                                    switch fi14 {
                                    case 32768:
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                    }
                                    if (false && false) {
                                        switch (^(fi14)) {
                                        case -700:
                                            v3 -= (v3 ^ (v3 << ((806) & 63)))
                                        case 128:
                                            if ((-2.25 + -2.25) != ((0.0 / ((1.5 * (1.5 + 0.0)) + 3.14159)) - (3.14159 / 1.0))) {
                                                fmt.Println(fi15, v6)
                                            } else {
                                                if (0.0 > 3.14159) {
                                                    var v23 float64 = v6
                                                    _ = v23
                                                } else {
                                                    switch fi15 {
                                                    case 32767:
                                                        v13 = ((("a" + "a") + (v13 + "foo bar")) + v13)
                                                    default:
                                                        if false {
                                                            switch ((v13 + v13) + (("foo bar" + v13) + v13)) {
                                                            case "a":
                                                                var v24 int32 = 1
                                                                _ = v24
                                                            case "":
                                                                if true {
                                                                    v25 := v2[0]
                                                                    v13 = v13
                                                                    _ = v25
                                                                } else {
                                                                    if true {
                                                                        if ((fi14 / (((fi14 ^ (fi14 >> ((1) & 63)))) | 1)) == fi15) {
                                                                            if true {
                                                                                v13 = "a"
                                                                                fmt.Println(v3, v3)
                                                                                if len(v4) >= 2 {
                                                                                    v4 = v4[1:len(v4)]
                                                                                }
                                                                            }
                                                                            var v26 int16 = int16(fi15)
                                                                            if false {
                                                                                var v27 int64 = -9223372036854775808
                                                                                _ = v27
                                                                            }
                                                                            _ = v26
                                                                        }
                                                                        v13 = ""
                                                                        if false {
                                                                            v13 = (v13 + "foo bar")
                                                                            var v28 int = (fi14 + fi14)
                                                                            switch -2147483648 {
                                                                            case -31:
                                                                                if (0 >= 42) {
                                                                                    v29 := v2[2]
                                                                                    switch (fi15 & fi14) {
                                                                                    case 403:
                                                                                        switch fi14 {
                                                                                        case 727:
                                                                                            v30 := GenS2{F0: 127, F1: 255}
                                                                                            _ = v30
                                                                                        default:
                                                                                            v31 := GenS2{F0: 64, F1: 42}
                                                                                            v3 += v3
                                                                                            if (42 <= -682) {
                                                                                                if false {
                                                                                                    v29 -= v29
                                                                                                }
                                                                                                if ((v29 & 629) <= 58) {
                                                                                                    v32 := v2[1]
                                                                                                    _ = v32
                                                                                                }
                                                                                                var v33 int16 = 0
                                                                                                _ = v33
                                                                                            } else {
                                                                                                v29 = (v29 & uint32(v3))
                                                                                                switch 4294967295 {
                                                                                                case 128:
                                                                                                    delete(v2, 4294967295)
                                                                                                case -1:
                                                                                                    v6 = v6
                                                                                                default:
                                                                                                    v34 := GenS2{F0: 0, F1: (v3 * v3)}
                                                                                                    _ = v34
                                                                                                }
                                                                                                switch (fi15 | v28) {
                                                                                                case -1:
                                                                                                    var v35 int64 = -633
                                                                                                    _ = v35
                                                                                                case -9223372036854775808:
                                                                                                    switch 32767 {
                                                                                                    case 343:
                                                                                                        var v36 uint = 2
                                                                                                        if len(v1) >= 2 {
                                                                                                            v1 = v1[1:len(v1)]
                                                                                                        }
                                                                                                        if false {
                                                                                                            v28 = fi15
                                                                                                            v5 = (v5 * (v5 ^ (^(v5))))
                                                                                                            v13 = (v13 + v13)
                                                                                                            v3 = 128
                                                                                                            fmt.Println(v13, v29, v3)
                                                                                                            fmt.Println(v3, v28, v5)
                                                                                                        } else {
                                                                                                            v37 := GenS3{F0: v3}
                                                                                                            var v38 uint32 = (v29 % ((v29) | 1))
                                                                                                            if false {
                                                                                                                v28 = fi14
                                                                                                            } else {
                                                                                                                switch v13 {
                                                                                                                case "hello":
                                                                                                                    if true {
                                                                                                                        var v39 uint8 = v3
                                                                                                                        var v40 int64 = -807
                                                                                                                        _ = v39
                                                                                                                        _ = v40
                                                                                                                    } else {
                                                                                                                        switch "foo bar" {
                                                                                                                        case "foo bar":
                                                                                                                            var v41 float32 = 3.14159
                                                                                                                            _ = v41
                                                                                                                        case "hello":
                                                                                                                            var v42 bool = true
                                                                                                                            _ = v42
                                                                                                                        case "":
                                                                                                                            v43 := v2[4]
                                                                                                                            _ = v43
                                                                                                                        default:
                                                                                                                            var v44 float32 = ((0.0 - ((0.0 - 0.1) - (3.14159 * -2.25))) / 3.14159)
                                                                                                                            _ = v44
                                                                                                                        }
                                                                                                                    }
                                                                                                                case "foo bar":
                                                                                                                    v45 := GenS1{F0: 0, F1: (^(v36)), F2: -347}
                                                                                                                    _ = v45
                                                                                                                default:
                                                                                                                    if (true || true) {
                                                                                                                        if ((!false) && false) {
                                                                                                                            var v46 uint32 = v29
                                                                                                                            _ = v46
                                                                                                                        }
                                                                                                                        var v47 uint16 = v5
                                                                                                                        _ = v47
                                                                                                                    } else {
                                                                                                                        if false {
                                                                                                                            v29 = 65535
                                                                                                                            if (v3 <= v3) {
                                                                                                                                var v48 bool = ((42 >= 519) && false)
                                                                                                                                v49 := GenS2{F0: -1, F1: 41}
                                                                                                                                _ = v48
                                                                                                                                _ = v49
                                                                                                                            } else {
                                                                                                                                v50 := GenS4{F0: uint(v29), F1: 239, F2: v5}
                                                                                                                                var v51 int = fi15
                                                                                                                                v5 = (^(v5))
                                                                                                                                switch (fi14 - fi14) {
                                                                                                                                case -93:
                                                                                                                                    var v52 int8 = 0
                                                                                                                                    _ = v52
                                                                                                                                }
                                                                                                                                v53 := v2[(v38 >> ((v36) & 63))]
                                                                                                                                if len(v1) > 0 {
                                                                                                                                    v1[0] = (v13 + ((v13 + v13) + v13))
                                                                                                                                }
                                                                                                                                v54 := GenS5{F0: 791}
                                                                                                                                if false {
                                                                                                                                    var v55 float64 = v6
                                                                                                                                    if false {
                                                                                                                                        delete(v2, 0)
                                                                                                                                    } else {
                                                                                                                                        v13 = ("hello" + v13)
                                                                                                                                        var v56 float32 = 1.5
                                                                                                                                        var v57 int8 = -128
                                                                                                                                        var v58 int = fi15
                                                                                                                                        v59 := GenS2{F0: (v57 * v57), F1: (^(v3))}
                                                                                                                                        var v60 uint32 = 2147483648
                                                                                                                                        var v61 uint32 = v38
                                                                                                                                        _ = v56
                                                                                                                                        _ = v57
                                                                                                                                        _ = v58
                                                                                                                                        _ = v59
                                                                                                                                        _ = v60
                                                                                                                                        _ = v61
                                                                                                                                    }
                                                                                                                                    var v62 uint8 = 42
                                                                                                                                    _ = v55
                                                                                                                                    _ = v62
                                                                                                                                } else {
                                                                                                                                    fmt.Println(v38)
                                                                                                                                    if true {
                                                                                                                                        switch (v28 % ((int(v5)) | 1)) {
                                                                                                                                        case 636:
                                                                                                                                            if true {
                                                                                                                                                v51 = -198
                                                                                                                                                var v63 uint8 = 42
                                                                                                                                                v64 := GenS5{F0: 42}
                                                                                                                                                _ = v63
                                                                                                                                                _ = v64
                                                                                                                                            } else {
                                                                                                                                                var v65 string = "hello"
                                                                                                                                                v38 = 2147483648
                                                                                                                                                switch fi14 {
                                                                                                                                                case -9223372036854775808:
                                                                                                                                                    fmt.Println(v29, fi15, v65)
                                                                                                                                                    v3 -= uint8(v5)
                                                                                                                                                    v13 = ((v13 + v65) + ((v13 + "a") + "hello"))
                                                                                                                                                case 902:
                                                                                                                                                    switch v28 {
                                                                                                                                                    case 17:
                                                                                                                                                        var v66 int64 = -371
                                                                                                                                                        _ = v66
                                                                                                                                                    case -9223372036854775808:
                                                                                                                                                        v67 := GenS2{F0: 127, F1: (v3 | v3)}
                                                                                                                                                        _ = v67
                                                                                                                                                    default:
                                                                                                                                                        var v68 uint8 = (v3 + (v3 / ((uint8(v36)) | 1)))
                                                                                                                                                        _ = v68
                                                                                                                                                    }
                                                                                                                                                }
                                                                                                                                                _ = v65
                                                                                                                                            }
                                                                                                                                        case 128:
                                                                                                                                            switch v13 {
                                                                                                                                            case "foo bar":
                                                                                                                                                if false {
                                                                                                                                                    v69 := GenS2{F0: -128, F1: v3}
                                                                                                                                                    v5 -= v5
                                                                                                                                                    _ = v69
                                                                                                                                                }
                                                                                                                                            default:
                                                                                                                                                var v70 string = (v13 + "foo bar")
                                                                                                                                                _ = v70
                                                                                                                                            }
                                                                                                                                        }
                                                                                                                                    } else {
                                                                                                                                        v51 = 128
                                                                                                                                        switch "hello" {
                                                                                                                                        case "a":
                                                                                                                                            fmt.Println(v13, v29, v13)
                                                                                                                                        case "foo bar":
                                                                                                                                            v2[748] = uint32(v36)
                                                                                                                                        case "":
                                                                                                                                            var v71 uint8 = 248
                                                                                                                                            _ = v71
                                                                                                                                        default:
                                                                                                                                            fmt.Println(v29)
                                                                                                                                        }
                                                                                                                                        v72 := GenS5{F0: 0}
                                                                                                                                        v73 := GenS6{F0: 1}
                                                                                                                                        var v74 uint64 = 18446744073709551615
                                                                                                                                        var v75 int16 = 42
                                                                                                                                        _ = v72
                                                                                                                                        _ = v73
                                                                                                                                        _ = v74
                                                                                                                                        _ = v75
                                                                                                                                    }
                                                                                                                                }
                                                                                                                                _ = v50
                                                                                                                                _ = v51
                                                                                                                                _ = v53
                                                                                                                                _ = v54
                                                                                                                            }
                                                                                                                        } else {
                                                                                                                            switch fi14 {
                                                                                                                            case 4294967295:
                                                                                                                                fmt.Println(v5, v5, v6)
                                                                                                                            case -806:
                                                                                                                                v76 := GenS7{F0: -661, F1: 124}
                                                                                                                                _ = v76
                                                                                                                            default:
                                                                                                                                var v77 float32 = (float32(0.0) + 1.5)
                                                                                                                                _ = v77
                                                                                                                            }
                                                                                                                            if true {
                                                                                                                                v2[1] = v38
                                                                                                                                var v78 float64 = (-2.25 - v6)
                                                                                                                                v28 = (fi15 | fi14)
                                                                                                                                _ = v78
                                                                                                                            } else {
                                                                                                                                var v79 int32 = -2147483648
                                                                                                                                _ = v79
                                                                                                                            }
                                                                                                                        }
                                                                                                                        v80 := GenS8{F0: 1}
                                                                                                                        _ = v80
                                                                                                                    }
                                                                                                                }
                                                                                                                v81 := GenS1{F0: 4294967295, F1: v36, F2: 809}
                                                                                                                _ = v81
                                                                                                            }
                                                                                                            _ = v37
                                                                                                            _ = v38
                                                                                                        }
                                                                                                        _ = v36
                                                                                                    case 127:
                                                                                                        v82 := GenS9{F0: (v29 % (((v29 & uint32(v3))) | 1)), F1: -685}
                                                                                                        _ = v82
                                                                                                    default:
                                                                                                        fmt.Println(v3, v29, v28)
                                                                                                    }
                                                                                                case -2147483648:
                                                                                                    var v83 bool = (96 != 0)
                                                                                                    _ = v83
                                                                                                default:
                                                                                                    v5 = v5
                                                                                                    var v84 int64 = 886
                                                                                                    _ = v84
                                                                                                }
                                                                                            }
                                                                                            _ = v31
                                                                                        }
                                                                                    }
                                                                                    _ = v29
                                                                                }
                                                                            default:
                                                                                v85 := GenS2{F0: 42, F1: 1}
                                                                                _ = v85
                                                                            }
                                                                            if (false && (!((-2.25 - ((-2.25 - 3.14159) * 0.1)) <= 0.0))) {
                                                                                if (42 > -953) {
                                                                                    v86 := GenS10{F0: 1, F1: 42, F2: 0}
                                                                                    _ = v86
                                                                                }
                                                                                if (1e10 == (0.1 * float32((0.0 * v6)))) {
                                                                                    var v87 int = (v28 | fi15)
                                                                                    fmt.Println(v13, v6, v6)
                                                                                    _ = v87
                                                                                } else {
                                                                                    var v88 int16 = 376
                                                                                    v6 = float64(-2.25)
                                                                                    _ = v88
                                                                                }
                                                                                var v89 uint8 = v3
                                                                                _ = v89
                                                                            } else {
                                                                                v28 = (v28 ^ v28)
                                                                                v6 += (3.14159 + (v6 + (float64(1e10) - 0.1)))
                                                                                if (917 > 766) {
                                                                                    if (2 < 836) {
                                                                                        v2[125] = uint32(v3)
                                                                                        if (42 != v5) {
                                                                                            v90 := GenS11{F0: -450, F1: (genFn0() <= (^(v5)))}
                                                                                            var v91 uint64 = 4294967295
                                                                                            fmt.Println(v28, v3, v28)
                                                                                            _ = v90
                                                                                            _ = v91
                                                                                        }
                                                                                        switch "foo bar" {
                                                                                        case "foo bar":
                                                                                            fmt.Println(v6, fi15)
                                                                                        case "\n\t":
                                                                                            v92 := GenS5{F0: -2147483648}
                                                                                            _ = v92
                                                                                        case "a":
                                                                                            v28 = fi14
                                                                                        }
                                                                                    }
                                                                                    if false {
                                                                                        var v93 int = 784
                                                                                        var v94 uint64 = 810
                                                                                        _ = v93
                                                                                        _ = v94
                                                                                    }
                                                                                    v95 := GenS12{F0: genFn0(), F1: 439, F2: 109}
                                                                                    _ = v95
                                                                                } else {
                                                                                    v5 = (^(v5))
                                                                                    fmt.Println(v28, v3)
                                                                                }
                                                                            }
                                                                            _ = v28
                                                                        }
                                                                    } else {
                                                                        if false {
                                                                            if (v3 > v3) {
                                                                                v96 := GenS13{F0: fi14, F1: ((1.5 + (-2.25 * ((0.0 - -2.25) - -2.25))) < ((float32(v6) * -2.25) - (((3.14159 * -2.25) * ((float32((v6 + (v6 + (float64(-2.25) - ((((float64(0.1) / v6) * v6) * (3.14159 / float64(((1.5 + 1e10) + 3.14159)))) / float64(3.14159)))))) - 1.5) * 0.1)) * float32(v6))))}
                                                                                v97 := GenS10{F0: 564, F1: 9223372036854775807, F2: -1}
                                                                                if false {
                                                                                    _ = len(v4)
                                                                                    v5 = 0
                                                                                    v6 -= v6
                                                                                }
                                                                                _ = v96
                                                                                _ = v97
                                                                            }
                                                                            var v98 int = (fi14 % (((fi15 * (fi15 + (fi14 >> ((1) & 63))))) | 1))
                                                                            _ = v98
                                                                        }
                                                                        if (!false) {
                                                                            fmt.Println(fi14)
                                                                        } else {
                                                                            if true {
                                                                                v99 := GenS7{F0: 221, F1: 133}
                                                                                _ = v99
                                                                            }
                                                                            delete(v2, 2147483647)
                                                                            _ = len(v4)
                                                                        }
                                                                        v5 = v5
                                                                    }
                                                                    var v100 uint32 = 2147483648
                                                                    var v101 uint64 = 42
                                                                    switch "a" {
                                                                    case "":
                                                                        var v102 uint32 = 889
                                                                        _ = v102
                                                                    default:
                                                                        if true {
                                                                            var v103 int8 = 100
                                                                            if false {
                                                                                v13 = v13
                                                                                if len(v4) >= 2 {
                                                                                    v4 = v4[1:len(v4)]
                                                                                }
                                                                            }
                                                                            _ = v103
                                                                        }
                                                                    }
                                                                    _ = v100
                                                                    _ = v101
                                                                }
                                                            default:
                                                                delete(v2, 27)
                                                                var v104 int16 = 42
                                                                var v105 bool = (-2147483648 < (fi14 + fi14))
                                                                _ = v104
                                                                _ = v105
                                                            }
                                                            v5 += 500
                                                            v5 = v5
                                                        } else {
                                                            if false {
                                                                v106 := v2[559]
                                                                _ = v106
                                                            } else {
                                                                v13 = v13
                                                                var v107 uint64 = uint64(v5)
                                                                v5 = (v5 >> ((377) & 63))
                                                                _ = v107
                                                            }
                                                        }
                                                    }
                                                    var v108 int64 = -1
                                                    v13 = ("" + (v13 + v13))
                                                    _ = v108
                                                }
                                            }
                                        default:
                                            fmt.Println(v6)
                                        }
                                    } else {
                                        if true {
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            v3 += v3
                                            v109 := GenS9{F0: 0, F1: -711}
                                            _ = v109
                                        }
                                    }
                                }
                            }
                            fmt.Println(v13, v6, fi14)
                            _ = fi15
                        }
                        switch ("" + "foo bar") {
                        case "foo bar":
                            v6 += (-2.25 + (0.1 + 3.14159))
                        case "hello":
                            v110 := GenS14{F0: 0, F1: -9223372036854775808}
                            _ = v110
                        case "a":
                            v111 := GenS15{F0: 128, F1: -1, F2: 284}
                            _ = v111
                        }
                    }
                default:
                    switch (fi14 << ((4294967295) & 63)) {
                    case -9223372036854775808:
                        v112 := GenS16{F0: 42}
                        _ = v112
                    }
                }
                _ = fi14
            }
            _ = len(v2)
            switch 9223372036854775807 {
            case -994:
                v5 = genFn0()
            case -302:
                var v113 uint64 = 1
                _ = v113
            }
        }
        v114 := GenS17{F0: v3, F1: 1}
        _ = v114
    } else {
        var v115 int64 = -1
        var v116 int16 = 67
        _ = v115
        _ = v116
    }
    fmt.Println(v5, v6)
    v117 := GenS14{F0: -866, F1: 0}
    if len(v1) > 0 {
        v1[0] = "foo bar"
    }
    switch "\n\t" {
    case "hello":
        v118 := GenS16{F0: -2147483648}
        _ = v118
    case "\n\t":
        v119 := GenS4{F0: 255, F1: 864, F2: (v5 >> ((42) & 63))}
        _ = v119
    case "":
        var v120 int8 = 76
        _ = v120
    default:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    for fi121 := 0; fi121 < 4; fi121++ {
        v122 := GenS1{F0: 124, F1: 42, F2: 628}
        if (!((698 != 255) || true)) {
            for fi123 := 0; fi123 < 15; fi123++ {
                v6 += ((v6 - (v6 / v6)) + v6)
                v5 += v5
                if (3.14159 < 0.1) {
                    var v124 float64 = 1.5
                    fmt.Println(v13, v5)
                    _ = v124
                } else {
                    if (false || false) {
                        fmt.Println(v6, fi121)
                        if false {
                            v4 = append(v4, "hello")
                            switch -2147483648 {
                            case 4294967295:
                                switch ((v13 + ("a" + ((v13 + v13) + (((v13 + (v13 + "foo bar")) + v13) + v13)))) + (v13 + v13)) {
                                case "":
                                    if (!((!true) && (false && (-1 >= 100)))) {
                                        v125 := GenS1{F0: 18446744073709551615, F1: 1, F2: 42}
                                        _ = v125
                                    } else {
                                        if true {
                                            v4 = append(v4, "a")
                                        }
                                        var v126 bool = (false || true)
                                        _ = v126
                                    }
                                case "a":
                                    switch (fi123 / ((4294967295) | 1)) {
                                    case 775:
                                        v127 := GenS18{F0: v13, F1: 742, F2: 1}
                                        _ = v127
                                    case 128:
                                        v1 = append(v1, "hello")
                                    case 303:
                                        var v128 int8 = 100
                                        _ = v128
                                    }
                                case "\n\t":
                                    v129 := GenS19{F0: (-(fi123)), F1: 263}
                                    _ = v129
                                default:
                                    if (42 >= 67) {
                                        v130 := v2[65535]
                                        v4 = append(v4, v13)
                                        switch v13 {
                                        case "\n\t":
                                            var v131 uint = 42
                                            _ = v131
                                        case "foo bar":
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                        case "":
                                            var v132 int8 = -94
                                            _ = v132
                                        }
                                        _ = v130
                                    }
                                    switch 2147483647 {
                                    case 613:
                                        v1 = append(v1, (v13 + ("a" + v13)))
                                    case 9223372036854775807:
                                        v133 := GenS7{F0: 42, F1: 340}
                                        var v134 int8 = -1
                                        _ = v133
                                        _ = v134
                                    default:
                                        var v135 int16 = -32768
                                        _ = v135
                                    }
                                    if ((^(fi121)) > (fi123 / ((int(v3)) | 1))) {
                                        v136 := GenS16{F0: 544}
                                        _ = v136
                                    } else {
                                        _ = len(v2)
                                    }
                                }
                            case 1:
                                v6 = v6
                            case 128:
                                v6 = v6
                            default:
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            }
                        }
                    } else {
                        fmt.Println(v5, fi123, v5)
                        v137 := GenS20{F0: (v5 & (v5 * (^(v5))))}
                        _ = v137
                    }
                    switch ("\n\t" + "a") {
                    case "hello":
                        if false {
                            if (("\n\t" + "\n\t") < (v13 + ((((v13 + (v13 + v13)) + (v13 + v13)) + v13) + v13))) {
                                v3 = 42
                                v5 -= uint16(v3)
                                var v138 uint64 = 2
                                var v139 float32 = 1e10
                                _ = v138
                                _ = v139
                            }
                            var v140 int = 4294967295
                            var v141 bool = (!false)
                            v141 = (18446744073709551615 < 1)
                            v13 = ("a" + v13)
                            if (190 < 0) {
                                var v142 float64 = (0.1 / (v6 - float64(((0.0 + 0.1) + 0.0))))
                                _ = v142
                            } else {
                                var v143 uint8 = v3
                                _ = v143
                            }
                            _ = len(v4)
                            _ = v140
                            _ = v141
                        }
                    case "a":
                        v13 = v13
                    case "foo bar":
                        v13 = "a"
                    }
                    fmt.Println(v3, fi123, v3)
                }
                if (0 >= -505) {
                    if false {
                        v5 += 255
                        v144 := v2[2147483648]
                        _ = v144
                    } else {
                        _ = len(v1)
                        switch fi123 {
                        case -27:
                            fmt.Println(fi123, v5)
                        }
                        v145 := GenS20{F0: genFn0()}
                        _ = v145
                    }
                    v6 = (v6 * v6)
                }
                _ = len(v2)
                _ = fi123
            }
            if true {
                v146 := GenS7{F0: 475, F1: 2}
                v147 := GenS7{F0: -2147483648, F1: 1}
                if (!false) {
                    switch "\n\t" {
                    case "foo bar":
                        v148 := GenS21{F0: 2147483647}
                        _ = v148
                    default:
                        v149 := GenS16{F0: -2147483648}
                        _ = v149
                    }
                    if (4294967295 > 1) {
                        for fi150 := 0; fi150 < 14; fi150++ {
                            v151 := GenS22{F0: v13}
                            fmt.Println(fi150)
                            _ = fi150
                            _ = v151
                        }
                        v3 -= v3
                    }
                }
                _ = v146
                _ = v147
            } else {
                switch v13 {
                case "a":
                    if true {
                        if true {
                            switch v13 {
                            case "foo bar":
                                var v152 float32 = 0.1
                                _ = v152
                            case "":
                                var v153 uint8 = (v3 >> ((1) & 63))
                                _ = v153
                            case "hello":
                                if true {
                                    v154 := GenS16{F0: 1}
                                    for fi155 := 0; fi155 < 3; fi155++ {
                                        var v156 int8 = 0
                                        _ = fi155
                                        _ = v156
                                    }
                                    v157 := v2[42]
                                    _ = v154
                                    _ = v157
                                }
                            }
                            for fi158 := 0; fi158 < 7; fi158++ {
                                v6 += v6
                                v13 = "foo bar"
                                switch ("foo bar" + "") {
                                case "foo bar":
                                    v6 = (v6 + v6)
                                case "":
                                    if (!(65535 > 447)) {
                                        v159 := GenS9{F0: 0, F1: 0}
                                        v3 -= 42
                                        switch fi121 {
                                        case 856:
                                            var v160 int16 = -1
                                            var v161 int64 = 1
                                            v3 = 127
                                            var v162 float32 = (float32(v6) - (0.0 - 0.1))
                                            fmt.Println(v6, v161)
                                            if true {
                                                if (!(801 <= v161)) {
                                                    v6 += v6
                                                } else {
                                                    v163 := GenS23{F0: int16(fi121), F1: (float64((v162 * 1.5)) != ((v6 + v6) - v6)), F2: ("foo bar" + v13)}
                                                    switch v13 {
                                                    case "foo bar":
                                                        fmt.Println(v162, v160)
                                                    }
                                                    if false {
                                                        var v164 uint = 4294967295
                                                        switch "" {
                                                        case "hello":
                                                            if false {
                                                                var v165 int64 = (-(v161))
                                                                v166 := GenS24{F0: (!(2147483647 == -1)), F1: 19}
                                                                _ = v165
                                                                _ = v166
                                                            } else {
                                                                v1 = append(v1, v13)
                                                                _ = len(v2)
                                                            }
                                                        case "\n\t":
                                                            switch v13 {
                                                            case "foo bar":
                                                                var v167 int64 = (^(v161))
                                                                _ = v167
                                                            case "hello":
                                                                var v168 uint64 = 18446744073709551615
                                                                _ = v168
                                                            case "":
                                                                v160 -= (-(v160))
                                                            default:
                                                                if len(v1) > 0 {
                                                                    v1[0] = (v13 + v13)
                                                                }
                                                            }
                                                        default:
                                                            var v169 uint = (v164 << (((v164 + v164)) & 63))
                                                            _ = v169
                                                        }
                                                        fmt.Println(v162)
                                                        _ = v164
                                                    } else {
                                                        var v170 uint = uint(v3)
                                                        var v171 bool = ((true || false) || ((v161 <= v161) && false))
                                                        fmt.Println(v3)
                                                        var v172 float32 = ((((v162 + v162) * 1.5) * -2.25) * v162)
                                                        _ = v170
                                                        _ = v171
                                                        _ = v172
                                                    }
                                                    switch ((v13 + (((v13 + v13) + (("foo bar" + ("" + ("" + v13))) + v13)) + (v13 + ((v13 + v13) + v13)))) + v13) {
                                                    case "hello":
                                                        fmt.Println(v13, v13, v3)
                                                    }
                                                    if (65535 <= v5) {
                                                        fmt.Println(fi121)
                                                        var v173 uint16 = (v5 + 48)
                                                        _ = v173
                                                    } else {
                                                        if (uint(v5) <= 1) {
                                                            var v174 float32 = 0.1
                                                            switch v13 {
                                                            case "foo bar":
                                                                fmt.Println(v13, v160, v161)
                                                            case "":
                                                                var v175 float64 = (v6 / v6)
                                                                _ = v175
                                                            default:
                                                                if false {
                                                                    if (!(((1e10 / (v6 / v6)) * v6) >= v6)) {
                                                                        v176 := GenS1{F0: 964, F1: 446, F2: -2147483648}
                                                                        var v177 int = fi121
                                                                        if true {
                                                                            v178 := GenS4{F0: 0, F1: v160, F2: uint16(v3)}
                                                                            switch fi158 {
                                                                            case 1:
                                                                                if false {
                                                                                    var v179 uint16 = v5
                                                                                    if ((v5 ^ v5) == genFn0()) {
                                                                                        switch v13 {
                                                                                        case "hello":
                                                                                            var v180 uint8 = v3
                                                                                            _ = v180
                                                                                        default:
                                                                                            fmt.Println(v179)
                                                                                        }
                                                                                        v162 -= 1.5
                                                                                        v161 += -1
                                                                                    }
                                                                                    _ = v179
                                                                                }
                                                                            case 32767:
                                                                                var v181 float32 = (v162 + v174)
                                                                                _ = v181
                                                                            }
                                                                            v160 = (v160 ^ (v160 ^ -1))
                                                                            var v182 uint = 255
                                                                            _ = v178
                                                                            _ = v182
                                                                        } else {
                                                                            _ = len(v4)
                                                                        }
                                                                        _ = v176
                                                                        _ = v177
                                                                    }
                                                                }
                                                            }
                                                            _ = v174
                                                        }
                                                        v183 := GenS3{F0: v3}
                                                        _ = v183
                                                    }
                                                    v5 = 42
                                                    v4 = append(v4, (((v13 + "a") + ((("a" + ("\n\t" + v13)) + ("foo bar" + v13)) + "foo bar")) + v13))
                                                    if true {
                                                        var v184 float32 = ((v162 + v162) / v162)
                                                        _ = len(v1)
                                                        _ = v184
                                                    }
                                                    var v185 uint32 = 593
                                                    v5 = (v5 & v5)
                                                    switch (fi121 << ((2) & 63)) {
                                                    case -92:
                                                        var v186 uint32 = uint32(v3)
                                                        _ = v186
                                                    case 97:
                                                        v13 = ((v13 + "") + v13)
                                                    case 42:
                                                        if (((v5 % (((v5 / (((v5 << ((0) & 63))) | 1))) | 1)) < genFn0()) && (v185 > 1)) {
                                                            switch v13 {
                                                            case "foo bar":
                                                                fmt.Println(v13, fi121, v5)
                                                            case "a":
                                                                v5 += 261
                                                            case "\n\t":
                                                                v160 = v160
                                                            default:
                                                                v187 := GenS3{F0: 1}
                                                                _ = v187
                                                            }
                                                            var v188 uint8 = 0
                                                            _ = len(v1)
                                                            _ = v188
                                                        } else {
                                                            v189 := GenS5{F0: -8}
                                                            var v190 uint64 = 2
                                                            delete(v2, v185)
                                                            _ = v189
                                                            _ = v190
                                                        }
                                                        v161 = v161
                                                        v4 = append(v4, v13)
                                                    default:
                                                        v191 := GenS13{F0: fi158, F1: ((v13 + "hello") != v13)}
                                                        _ = v191
                                                    }
                                                    _ = v163
                                                    _ = v185
                                                }
                                                var v192 float64 = 0.0
                                                v13 = ""
                                                _ = v192
                                            } else {
                                                v3 = v3
                                                _ = len(v4)
                                                var v193 int = fi158
                                                _ = v193
                                            }
                                            _ = v160
                                            _ = v161
                                            _ = v162
                                        default:
                                            v2[0] = 0
                                        }
                                        var v194 uint = 0
                                        _ = v159
                                        _ = v194
                                    }
                                case "a":
                                    fmt.Println(fi121, fi158, fi121)
                                default:
                                    var v195 float64 = 1.5
                                    _ = v195
                                }
                                v196 := GenS9{F0: 301, F1: 1}
                                fmt.Println(v5)
                                _ = fi158
                                _ = v196
                            }
                        } else {
                            var v197 uint32 = 0
                            v198 := GenS23{F0: 86, F1: (v6 <= (((3.14159 - v6) + 3.14159) * (float64(((1.5 - float32(float64(0.0))) - (-2.25 + (1e10 * 1e10)))) * float64(1.5)))), F2: ((v13 + ("hello" + "a")) + ("hello" + (v13 + (v13 + v13))))}
                            _ = v197
                            _ = v198
                        }
                        var v199 int = fi121
                        _ = v199
                    } else {
                        v3 = (v3 << ((672) & 63))
                    }
                case "hello":
                    v200 := GenS25{F0: v3}
                    _ = v200
                case "foo bar":
                    switch -757 {
                    case 4294967295:
                        if false {
                            for fi201 := 0; fi201 < 2; fi201++ {
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                if false {
                                    if false {
                                        fmt.Println(v6, v13, v13)
                                        fmt.Println(v6)
                                        switch (("hello" + "a") + (((v13 + ("hello" + v13)) + "\n\t") + ("a" + (((v13 + (v13 + v13)) + "hello") + v13)))) {
                                        case "foo bar":
                                            fmt.Println(fi121, v6, fi201)
                                        default:
                                            _ = len(v2)
                                        }
                                    } else {
                                        _ = len(v1)
                                        if true {
                                            v202 := GenS26{F0: -1, F1: v13, F2: v13}
                                            if false {
                                                v6 += v6
                                            }
                                            var v203 int = (fi121 << ((1) & 63))
                                            switch (fi201 - fi201) {
                                            case 128:
                                                if true {
                                                    if false {
                                                        var v204 uint64 = 562
                                                        _ = v204
                                                    }
                                                }
                                            case -2147483648:
                                                if false {
                                                    v203 = 42
                                                    var v205 bool = ((!true) || (!(132 <= 0)))
                                                    v2[905] = 42
                                                    _ = v205
                                                }
                                            case -640:
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                            }
                                            v6 += v6
                                            _ = v202
                                            _ = v203
                                        }
                                    }
                                    fmt.Println(fi201, v3)
                                }
                                var v206 uint = 1
                                _ = fi201
                                _ = v206
                            }
                        }
                    case 2147483647:
                        var v207 int64 = -9223372036854775808
                        _ = v207
                    }
                }
                var v208 int32 = 42
                var v209 float32 = ((1.5 + 0.0) - ((3.14159 * 0.0) / 1.0))
                _ = v208
                _ = v209
            }
        } else {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            if false {
                for fi210 := 0; fi210 < 18; fi210++ {
                    if true {
                        var v211 int = (fi210 + (fi121 & (fi210 + 478)))
                        if true {
                            var v212 int8 = -56
                            _ = len(v2)
                            v213 := GenS9{F0: 42, F1: 220}
                            _ = v212
                            _ = v213
                        } else {
                            if (1 > 65535) {
                                switch (fi121 - -2147483648) {
                                case 915:
                                    var v214 int8 = 39
                                    _ = v214
                                case 4294967295:
                                    if false {
                                        var v215 bool = (972 < 4294967295)
                                        v1 = append(v1, v13)
                                        _ = v215
                                    } else {
                                        var v216 float64 = (((v6 - float64((0.0 - (0.0 / (float32((v6 - v6)) / (3.14159 + 0.1)))))) - v6) * (float64((1.5 / 1.0)) / v6))
                                        v217 := v2[293]
                                        _ = v216
                                        _ = v217
                                    }
                                default:
                                    var v218 float64 = v6
                                    _ = v218
                                }
                                var v219 uint64 = 4294967295
                                v13 = ("foo bar" + ((v13 + "hello") + (v13 + v13)))
                                _ = v219
                            } else {
                                v5 -= 673
                                var v220 bool = true
                                fmt.Println(fi121, v6)
                                v221 := GenS7{F0: -1, F1: 979}
                                _ = v220
                                _ = v221
                            }
                            var v222 int32 = -2147483648
                            var v223 string = v13
                            _ = v222
                            _ = v223
                        }
                        _ = v211
                    } else {
                        var v224 int8 = -44
                        if (4294967295 <= 1) {
                            v225 := GenS10{F0: (^(v5)), F1: 9223372036854775807, F2: 42}
                            v5 = genFn0()
                            var v226 uint8 = 255
                            _ = v225
                            _ = v226
                        }
                        delete(v2, 91)
                        if (("" + v13) == v13) {
                            var v227 uint8 = 203
                            v228 := GenS4{F0: 269, F1: 1, F2: 904}
                            v229 := v2[2]
                            _ = v227
                            _ = v228
                            _ = v229
                        } else {
                            if true {
                                v2[915] = 2147483648
                                fmt.Println(v3, v224)
                                fmt.Println(fi210)
                            }
                        }
                        var v230 bool = (!false)
                        v13 = "hello"
                        _ = v224
                        _ = v230
                    }
                    if (32767 <= v5) {
                        fmt.Println(v3, fi210, fi210)
                        if true {
                            fmt.Println(v3, fi121, v3)
                            v13 = (("\n\t" + "hello") + v13)
                        } else {
                            fmt.Println(v6)
                        }
                        if (1 <= 1) {
                            var v231 float32 = (1.5 / ((3.14159 / float32((v6 - v6))) / (1e10 + 1.5)))
                            var v232 uint8 = v3
                            var v233 int16 = 366
                            _ = v231
                            _ = v232
                            _ = v233
                        }
                    } else {
                        v13 = v13
                        v234 := GenS25{F0: (v3 | v3)}
                        _ = v234
                    }
                    v5 = (^(v5))
                    _ = fi210
                }
                fmt.Println(v13, v5, v3)
            }
        }
        v235 := v2[1]
        _ = fi121
        _ = v122
        _ = v235
    }
    var v236 float32 = -2.25
    v237 := GenS27{F0: 0, F1: v3, F2: (v13 + v13)}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v236:", v236)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v117
    _ = v13
    _ = v2
    _ = v236
    _ = v237
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
