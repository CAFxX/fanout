package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int8; F1 uint32; F2 uint32 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint16; F1 uint; F2 uint64 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 int }
type GenS10 struct { F0 uint16; F1 int8; F2 string }
type GenS11 struct { F0 string; F1 uint }
type GenS12 struct { F0 bool; F1 uint32 }
type GenS13 struct { F0 int64; F1 uint32 }
type GenS14 struct { F0 string; F1 int64; F2 uint32 }
type GenS15 struct { F0 string; F1 int; F2 bool }
type GenS16 struct { F0 bool; F1 uint8 }
type GenS17 struct { F0 int8; F1 uint }
type GenS18 struct { F0 int8; F1 int8; F2 int16 }
func genFn0(p0 uint32) uint16 {
    v6 := GenS1{F0: -824}
    fmt.Println(p0)
    var v7 uint32 = p0
    _ = p0
    _ = v6
    _ = v7
    return 42
}

func main() {
    v1 := []string{("\n\t" + "a"), "hello", ((((("foo bar" + "") + (("a" + "a") + "foo bar")) + "\n\t") + "") + "\n\t"), (("" + ("hello" + ("foo bar" + "foo bar"))) + "\n\t")}
    var v2 float32 = 3.14159
    v3 := make(map[int]int)
    v4 := make(map[string]int)
    var v5 float32 = 0.1
    v1 = append(v1, "hello")
    if true {
        var v8 int32 = 0
        for fi9 := 0; fi9 < 5; fi9++ {
            for fi10 := 0; fi10 < 3; fi10++ {
                switch "foo bar" {
                case "a":
                    var v11 int64 = 1
                    _ = v11
                case "\n\t":
                    if true {
                        var v12 int8 = 124
                        var v13 int32 = (v8 ^ -173)
                        var v14 int32 = (v13 >> ((1) & 63))
                        v15 := GenS1{F0: -32768}
                        var v16 uint32 = 2147483647
                        if true {
                            var v17 uint16 = 812
                            var v18 float32 = v5
                            switch 4294967295 {
                            case -1:
                                if (!true) {
                                    if ((!(!(32767 != 32767))) || ("\n\t" == "hello")) {
                                        var v19 uint64 = 42
                                        if ((int64(v16) != 0) && true) {
                                            v20 := GenS2{F0: (-(v12)), F1: 7, F2: (^(v16))}
                                            v16 = 2147483647
                                            if (42 <= 877) {
                                                if false {
                                                    var v21 uint16 = 255
                                                    _ = v21
                                                } else {
                                                    var v22 string = "hello"
                                                    switch int(v19) {
                                                    case 128:
                                                        v3[-2147483648] = fi10
                                                        v19 += (^(v19))
                                                        if true {
                                                            v23 := GenS1{F0: 759}
                                                            var v24 uint32 = 42
                                                            if true {
                                                                switch 1 {
                                                                case -2147483648:
                                                                    switch (v22 + (v22 + (v22 + v22))) {
                                                                    case "foo bar":
                                                                        if true {
                                                                            if (128 < 127) {
                                                                                if false {
                                                                                    _ = len(v3)
                                                                                    var v25 float64 = 1.5
                                                                                    _ = v25
                                                                                } else {
                                                                                    if (!true) {
                                                                                        var v26 int32 = (v8 + (v14 << ((2) & 63)))
                                                                                        v16 += (v24 ^ (v16 - v24))
                                                                                        v12 = v12
                                                                                        _ = v26
                                                                                    }
                                                                                    v13 = 0
                                                                                }
                                                                                v27 := GenS1{F0: 32767}
                                                                                v14 = 1
                                                                                _ = v27
                                                                            }
                                                                            fmt.Println(v2)
                                                                            var v28 int16 = -576
                                                                            _ = v28
                                                                        }
                                                                    default:
                                                                        var v29 int16 = 94
                                                                        _ = v29
                                                                    }
                                                                default:
                                                                    v30 := GenS1{F0: 1000}
                                                                    _ = v30
                                                                }
                                                            } else {
                                                                var v31 uint16 = (^(v17))
                                                                v5 = (((float32((3.14159 - (float64((float32(1.5) * v2)) + (0.0 / 1.0)))) / float32(float64(-2.25))) / v5) * v5)
                                                                _ = v31
                                                            }
                                                            _ = v23
                                                            _ = v24
                                                        } else {
                                                            v22 = v22
                                                            v2 = (0.0 / 1.0)
                                                            if (v19 > (v19 << ((1) & 63))) {
                                                                switch fi10 {
                                                                case -911:
                                                                    var v32 int = (fi10 ^ fi9)
                                                                    _ = v32
                                                                case 9223372036854775807:
                                                                    var v33 int8 = v12
                                                                    _ = v33
                                                                case -462:
                                                                    var v34 int = (fi10 | fi10)
                                                                    _ = v34
                                                                }
                                                            } else {
                                                                v35 := GenS2{F0: -128, F1: (v16 + uint32(v17)), F2: v16}
                                                                v2 += 1e10
                                                                switch (fi10 | 32767) {
                                                                case 978:
                                                                    if (v13 != (v14 | (-(v8)))) {
                                                                        delete(v3, (fi10 + 1))
                                                                        v14 += -68
                                                                        if true {
                                                                            if len(v1) >= 2 {
                                                                                v1 = v1[1:len(v1)]
                                                                            }
                                                                        } else {
                                                                            var v36 bool = false
                                                                            v36 = (9223372036854775807 >= 1)
                                                                            if (false || ((3.14159 * 1.5) != ((0.1 - float64((float32(float64(float32(0.1))) * v5))) * 0.0))) {
                                                                                var v37 float32 = ((v2 - v2) - float32((1.5 * 1.5)))
                                                                                fmt.Println(v12, fi10)
                                                                                _ = v37
                                                                            } else {
                                                                                var v38 float64 = float64(0.1)
                                                                                _ = v38
                                                                            }
                                                                            _ = v36
                                                                        }
                                                                    } else {
                                                                        if false {
                                                                            v39 := GenS2{F0: v12, F1: v16, F2: (v16 << ((1) & 63))}
                                                                            _ = v39
                                                                        }
                                                                        v1 = append(v1, (("hello" + v22) + (v22 + ((v22 + ((("\n\t" + v22) + v22) + (v22 + ((v22 + (v22 + v22)) + "\n\t")))) + v22))))
                                                                    }
                                                                case 42:
                                                                    v22 = ((((v22 + "hello") + v22) + v22) + ("hello" + (v22 + ("" + ("" + "foo bar")))))
                                                                default:
                                                                    if false {
                                                                        fmt.Println(v12)
                                                                        if len(v1) >= 2 {
                                                                            v1 = v1[1:len(v1)]
                                                                        }
                                                                        var v40 uint64 = v19
                                                                        var v41 uint32 = v16
                                                                        v42 := GenS2{F0: 42, F1: (v16 - (v16 ^ v41)), F2: (v41 % ((v16) | 1))}
                                                                        _ = v40
                                                                        _ = v41
                                                                        _ = v42
                                                                    }
                                                                }
                                                                _ = v35
                                                            }
                                                        }
                                                    case -2147483648:
                                                        var v43 int16 = 32767
                                                        _ = v43
                                                    default:
                                                        if (false || (1 == 1)) {
                                                            v3[(^(fi10))] = fi9
                                                            if true {
                                                                if (127 >= 101) {
                                                                    var v44 uint32 = (^(v16))
                                                                    _ = v44
                                                                }
                                                                delete(v3, 32767)
                                                            }
                                                            var v45 uint16 = 255
                                                            _ = v45
                                                        } else {
                                                            fmt.Println(v18, fi10, v16)
                                                        }
                                                        if true {
                                                            var v46 bool = ("\n\t" < v22)
                                                            var v47 float32 = v18
                                                            v17 += 416
                                                            _ = v46
                                                            _ = v47
                                                        } else {
                                                            v3[128] = (fi10 - fi9)
                                                            v13 = 42
                                                        }
                                                        var v48 uint64 = v19
                                                        _ = v48
                                                    }
                                                    v5 -= ((float32(3.14159) * float32(3.14159)) * v18)
                                                    v2 = (float32(-2.25) / v18)
                                                    switch v22 {
                                                    case "hello":
                                                        v22 = v22
                                                    default:
                                                        fmt.Println(v14)
                                                    }
                                                    _ = v22
                                                }
                                            }
                                            _ = v20
                                        } else {
                                            switch ("foo bar" + ((("hello" + ("foo bar" + "")) + ("foo bar" + ("\n\t" + ("foo bar" + (((("hello" + "a") + "\n\t") + "") + ((("\n\t" + "") + "foo bar") + "foo bar")))))) + "hello")) {
                                            case "foo bar":
                                                var v49 float32 = (((v2 * (v2 - (v5 + v18))) - v18) / (v2 - v2))
                                                _ = v49
                                            default:
                                                var v50 uint64 = 727
                                                v3[fi9] = (fi10 + fi10)
                                                switch (fi10 & fi9) {
                                                case -9223372036854775808:
                                                    v51 := GenS3{F0: 32767}
                                                    _ = v51
                                                case -480:
                                                    if false {
                                                        var v52 int8 = -23
                                                        _ = len(v1)
                                                        _ = v52
                                                    } else {
                                                        var v53 int32 = v13
                                                        _ = v53
                                                    }
                                                }
                                                var v54 int64 = 9223372036854775807
                                                v12 -= v12
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                                _ = v50
                                                _ = v54
                                            }
                                            v55 := GenS2{F0: -128, F1: 0, F2: (^(v16))}
                                            _ = v55
                                        }
                                        v56 := GenS3{F0: genFn0((^(v16)))}
                                        _ = v19
                                        _ = v56
                                    } else {
                                        var v57 bool = false
                                        _ = v57
                                    }
                                    v17 += 65535
                                    v17 = 28
                                    fmt.Println(fi10)
                                } else {
                                    v5 = v5
                                    if ((v12 | (v12 + (v12 >> ((18446744073709551615) & 63)))) < (v12 & (^(v12)))) {
                                        var v58 bool = (false && (!((^(v14)) == (v8 % ((int32(fi9)) | 1)))))
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        _ = v58
                                    }
                                    v59 := GenS4{F0: genFn0((^(v16))), F1: 255, F2: 42}
                                    _ = v59
                                }
                            }
                            _ = v17
                            _ = v18
                        }
                        if false {
                            if false {
                                switch 551 {
                                case 42:
                                    v13 -= v8
                                    v60 := GenS5{F0: "a"}
                                    _ = v60
                                case 32768:
                                    fmt.Println(v16)
                                case 9223372036854775807:
                                    v61 := GenS6{F0: 921}
                                    _ = v61
                                }
                                if ((((3.14159 * v2) + v5) != v5) || true) {
                                    if true {
                                        var v62 uint = 471
                                        switch ("a" + "a") {
                                        case "hello":
                                            var v63 int64 = -109
                                            _ = v63
                                        case "\n\t":
                                            if (((float64(v2) / 1.0) >= (1.5 - 3.14159)) && true) {
                                                fmt.Println(v12, v14, fi9)
                                                var v64 uint16 = 312
                                                _ = v64
                                            } else {
                                                _ = len(v4)
                                                var v65 int64 = 0
                                                fmt.Println(fi10, v5)
                                                _ = v65
                                            }
                                            if false {
                                                v14 -= (v14 << ((v62) & 63))
                                                fmt.Println(v5, v12)
                                                var v66 string = "foo bar"
                                                _ = v66
                                            }
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                        }
                                        _ = v62
                                    }
                                    if true {
                                        var v67 uint32 = 4294967295
                                        if true {
                                            switch "" {
                                            case "a":
                                                fmt.Println(v8, v12, v16)
                                            case "\n\t":
                                                _ = len(v1)
                                            }
                                        }
                                        if len(v1) > 0 {
                                            v1[0] = "foo bar"
                                        }
                                        _ = v67
                                    } else {
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        if false {
                                            v5 = (-2.25 / ((v5 / 3.14159) + (((0.0 + 1.5) * v5) + -2.25)))
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            v68 := GenS4{F0: 221, F1: 18446744073709551615, F2: 42}
                                            _ = v68
                                        }
                                    }
                                    v69 := v4["a"]
                                    _ = v69
                                } else {
                                    var v70 uint = 2
                                    var v71 uint = 957
                                    _ = v70
                                    _ = v71
                                }
                                v14 = v13
                            }
                            v13 = (v14 - v8)
                            if (((float64(v2) - -2.25) / (1.5 * float64(-2.25))) > (1.5 / 1.5)) {
                                v72 := GenS7{F0: int32(v12)}
                                if (-128 == (v12 % (((v12 * (v12 * v12))) | 1))) {
                                    var v73 int16 = 1
                                    v74 := GenS4{F0: 1, F1: 42, F2: 4294967295}
                                    if true {
                                        v75 := v4["foo bar"]
                                        v76 := GenS3{F0: 255}
                                        _ = v75
                                        _ = v76
                                    }
                                    _ = v73
                                    _ = v74
                                }
                                _ = v72
                            }
                        } else {
                            v16 = (v16 >> ((537) & 63))
                        }
                        _ = v12
                        _ = v13
                        _ = v14
                        _ = v15
                        _ = v16
                    }
                }
                v77 := GenS8{F0: 668}
                v78 := GenS9{F0: fi9}
                _ = fi10
                _ = v77
                _ = v78
            }
            v79 := GenS8{F0: 1}
            v80 := GenS3{F0: 197}
            _ = fi9
            _ = v79
            _ = v80
        }
        v81 := GenS10{F0: 255, F1: -1, F2: ((("\n\t" + ("a" + "")) + "hello") + "")}
        _ = v8
        _ = v81
    }
    var v82 string = ("" + ("a" + ("a" + "\n\t")))
    for fi83 := 0; fi83 < 11; fi83++ {
        v82 = "\n\t"
        if (!true) {
            for fi84 := 0; fi84 < 17; fi84++ {
                v1 = append(v1, "a")
                v82 = ("hello" + v82)
                switch fi83 {
                case -636:
                    switch ("a" + v82) {
                    case "foo bar":
                        v82 = "hello"
                        switch v82 {
                        case "a":
                            v2 -= (((v5 * v5) + float32((0.0 + 0.1))) / float32(-2.25))
                        case "foo bar":
                            v85 := GenS11{F0: "\n\t", F1: 0}
                            _ = v85
                        case "\n\t":
                            v82 = (("a" + v82) + ("" + ((v82 + v82) + (("" + (v82 + (v82 + "hello"))) + v82))))
                        }
                    case "a":
                        v82 = (v82 + v82)
                    case "":
                        v86 := GenS12{F0: true, F1: 942}
                        _ = v86
                    default:
                        var v87 uint8 = 12
                        _ = v87
                    }
                case 128:
                    v82 = (("hello" + (v82 + (v82 + ("hello" + v82)))) + v82)
                case 146:
                    var v88 uint8 = 42
                    _ = v88
                default:
                    if ((1.5 - 1e10) >= -2.25) {
                        if false {
                            fmt.Println(fi83, fi83, v2)
                            switch -555 {
                            case 0:
                                var v89 uint32 = 2147483648
                                _ = v89
                            case 127:
                                if (227 == 42) {
                                    if false {
                                        v90 := GenS13{F0: -9223372036854775808, F1: 2147483647}
                                        _ = v90
                                    } else {
                                        switch -9223372036854775808 {
                                        case 32768:
                                            v82 = ("a" + "foo bar")
                                        case 9223372036854775807:
                                            var v91 int = fi84
                                            _ = v91
                                        case 127:
                                            v92 := GenS6{F0: 42}
                                            _ = v92
                                        }
                                    }
                                    switch v82 {
                                    case "foo bar":
                                        if (-32768 != -1) {
                                            var v93 float64 = ((1.5 / 3.14159) * (0.1 + 1e10))
                                            v5 = v5
                                            _ = v93
                                        } else {
                                            var v94 uint = 1
                                            if (1 <= 254) {
                                                fmt.Println(fi83, v5)
                                                if false {
                                                    var v95 uint = (^(v94))
                                                    v2 = ((v2 / 0.1) / ((float32(-2.25) + -2.25) / v5))
                                                    _ = v95
                                                }
                                                v82 = v82
                                            }
                                            var v96 float64 = 0.1
                                            v94 = (v94 + (v94 * 4294967295))
                                            _ = v94
                                            _ = v96
                                        }
                                    case "hello":
                                        if (!true) {
                                            switch (^(fi84)) {
                                            case -708:
                                                v5 += (v5 + v5)
                                            case 258:
                                                v2 = (-2.25 / v5)
                                            default:
                                                _ = len(v4)
                                            }
                                        }
                                    case "a":
                                        _ = len(v1)
                                    default:
                                        v97 := GenS14{F0: v82, F1: 462, F2: 4294967295}
                                        _ = v97
                                    }
                                    v4[(v82 + "")] = -2147483648
                                    fmt.Println(v2)
                                }
                            case 236:
                                var v98 uint16 = 575
                                switch v82 {
                                case "":
                                    fmt.Println(fi83, v82, fi84)
                                case "a":
                                    v99 := GenS15{F0: "\n\t", F1: fi83, F2: true}
                                    _ = v99
                                default:
                                    var v100 uint32 = 2147483648
                                    _ = v100
                                }
                                v82 = v82
                                v101 := GenS5{F0: v82}
                                if (true && false) {
                                    var v102 int64 = 923
                                    v3[32767] = (fi83 % (((fi83 << ((18446744073709551615) & 63))) | 1))
                                    v103 := GenS16{F0: false, F1: 255}
                                    _ = v102
                                    _ = v103
                                }
                                switch ((((v82 + "hello") + ("\n\t" + ((v82 + v82) + ""))) + v82) + v82) {
                                case "\n\t":
                                    if false {
                                        switch (-(fi83)) {
                                        case -1:
                                            v104 := GenS3{F0: (v98 * (v98 + (v98 * 528)))}
                                            _ = v104
                                        default:
                                            v98 = (^(v98))
                                        }
                                        if len(v1) > 0 {
                                            v1[0] = "a"
                                        }
                                        v2 = 0.1
                                    } else {
                                        switch (fi84 / (((fi84 - fi83)) | 1)) {
                                        case 315:
                                            v4[(v82 + ("a" + "a"))] = 32767
                                        case 127:
                                            var v105 float32 = 0.0
                                            _ = v105
                                        case 1:
                                            v106 := GenS14{F0: (v82 + ""), F1: -100, F2: 0}
                                            _ = v106
                                        }
                                    }
                                default:
                                    v107 := GenS6{F0: 1}
                                    _ = v107
                                }
                                _ = v101
                                _ = v98
                            default:
                                v4[("foo bar" + (v82 + ("hello" + v82)))] = (fi83 - 349)
                            }
                        } else {
                            if len(v1) > 0 {
                                v1[0] = v82
                            }
                            if (!true) {
                                fmt.Println(v5)
                                fmt.Println(v82, v5)
                                var v108 string = ("" + v82)
                                _ = v108
                            } else {
                                var v109 int = (fi84 / ((fi84) | 1))
                                _ = v109
                            }
                            v5 -= v2
                        }
                    }
                }
                v5 = -2.25
                _ = fi84
            }
            for fi110 := 0; fi110 < 2; fi110++ {
                if (!false) {
                    v82 = ("\n\t" + "foo bar")
                    var v111 uint32 = 1
                    fmt.Println(v111, fi110, v5)
                    _ = v111
                }
                if (float32(0.1) == (v2 - float32((((1e10 + 1.5) - ((0.1 * 1e10) * -2.25)) + float64(float32(0.0)))))) {
                    if false {
                        switch ("foo bar" + "\n\t") {
                        case "a":
                            if (!true) {
                                v82 = v82
                                v112 := GenS11{F0: ("" + ""), F1: 255}
                                var v113 int8 = -93
                                _ = v112
                                _ = v113
                            }
                        }
                        v5 += (3.14159 + v2)
                        if (1 > 127) {
                            v2 = float32(1e10)
                        } else {
                            fmt.Println(v82, fi110)
                            v5 += v5
                        }
                    } else {
                        if true {
                            switch fi110 {
                            case -130:
                                switch "foo bar" {
                                case "a":
                                    var v114 int8 = -128
                                    _ = v114
                                default:
                                    var v115 uint16 = 2
                                    _ = v115
                                }
                            case 32767:
                                switch "" {
                                case "foo bar":
                                    switch v82 {
                                    case "a":
                                        _ = len(v4)
                                    case "hello":
                                        _ = len(v3)
                                    case "\n\t":
                                        v2 -= (0.0 * (float32(1e10) * float32(-2.25)))
                                    default:
                                        var v116 uint32 = 1
                                        _ = v116
                                    }
                                case "":
                                    fmt.Println(v2, fi110, fi110)
                                default:
                                    var v117 uint8 = 1
                                    _ = v117
                                }
                            case 4294967295:
                                switch ("" + "foo bar") {
                                case "foo bar":
                                    v5 = (3.14159 + v2)
                                case "a":
                                    v118 := GenS16{F0: true, F1: 1}
                                    _ = v118
                                case "\n\t":
                                    var v119 int = -9223372036854775808
                                    _ = v119
                                default:
                                    v120 := GenS14{F0: v82, F1: 9223372036854775807, F2: 0}
                                    _ = v120
                                }
                            }
                        }
                        v5 += v2
                    }
                    switch "a" {
                    case "a":
                        var v121 uint64 = 42
                        _ = v121
                    case "":
                        fmt.Println(v82, v82, v2)
                    default:
                        var v122 float32 = (v5 / v5)
                        _ = v122
                    }
                    delete(v4, "\n\t")
                }
                _ = fi110
            }
        } else {
            if (!false) {
                if len(v1) > 0 {
                    v1[0] = ((v82 + v82) + "a")
                }
                var v123 int32 = -566
                var v124 uint16 = 2
                _ = v123
                _ = v124
            }
            fmt.Println(v82, v82, fi83)
            v125 := GenS5{F0: "foo bar"}
            _ = v125
        }
        v126 := GenS14{F0: v82, F1: 9223372036854775807, F2: 4294967295}
        _ = fi83
        _ = v126
    }
    v4[v82] = 32768
    if false {
        if false {
            v127 := GenS14{F0: v82, F1: 527, F2: 1}
            for fi128 := 0; fi128 < 7; fi128++ {
                v129 := GenS17{F0: 127, F1: 42}
                v2 = v2
                if len(v1) > 0 {
                    v1[0] = ""
                }
                _ = fi128
                _ = v129
            }
            _ = v127
        } else {
            v2 = 3.14159
        }
        var v130 bool = false
        for fi131 := 0; fi131 < 19; fi131++ {
            switch "foo bar" {
            case "":
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            case "foo bar":
                if v130 {
                    v1 = append(v1, "\n\t")
                    v130 = (v130 && false)
                } else {
                    var v132 uint = 627
                    for fi133 := 0; fi133 < 7; fi133++ {
                        if (431 <= 964) {
                            if (-1 < -843) {
                                fmt.Println(v130)
                                var v134 int32 = -363
                                _ = v134
                            }
                            v5 -= (0.0 - (1.5 + v2))
                        }
                        v132 = 0
                        _ = fi133
                    }
                    _ = v132
                }
            default:
                _ = len(v4)
            }
            switch ("" + "hello") {
            case "foo bar":
                if (2 <= 298) {
                    v135 := GenS18{F0: 0, F1: 42, F2: 959}
                    v2 -= float32(3.14159)
                    v82 = v82
                    _ = v135
                }
            case "\n\t":
                if (v130 || v130) {
                    var v136 int16 = -32768
                    _ = v136
                } else {
                    fmt.Println(v2, fi131)
                }
            default:
                v137 := GenS4{F0: genFn0(65535), F1: 948, F2: 18446744073709551615}
                _ = v137
            }
            _ = fi131
        }
        _ = v130
    }
    v5 += ((v5 * v2) / v2)
    fmt.Println(v2, v2, v82)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if (false && ((!(v2 > (0.0 - v2))) && true)) {
        var v138 uint = 0
        if len(v1) > 0 {
            v1[0] = "\n\t"
        }
        _ = v138
    } else {
        v5 = v2
        fmt.Println(v82, v5, v2)
    }
    if ((!false) || ((!(32767 > -6)) || (false && (!true)))) {
        v2 = 1.5
        v82 = ("foo bar" + v82)
        v139 := GenS16{F0: false, F1: 2}
        _ = v139
    }
    v1 = append(v1, (v82 + v82))
    var v140 int16 = 0
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v140:", v140)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v82:", v82)
    _ = v1
    _ = v140
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v82
}
