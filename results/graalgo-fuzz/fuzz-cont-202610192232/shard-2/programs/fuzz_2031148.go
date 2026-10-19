package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int; F1 string }
type GenS3 struct { F0 uint8; F1 int64; F2 bool }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 int32; F1 int8; F2 int16 }
type GenS7 struct { F0 int; F1 int32; F2 int32 }
type GenS8 struct { F0 string }
type GenS9 struct { F0 uint32 }
type GenS10 struct { F0 uint; F1 int16 }
type GenS11 struct { F0 uint32; F1 uint16 }
type GenS12 struct { F0 int8; F1 uint16 }
type GenS13 struct { F0 uint64; F1 int8 }
type GenS14 struct { F0 int64; F1 int32; F2 int8 }
type GenS15 struct { F0 uint16; F1 int16 }
type GenS16 struct { F0 int32; F1 int }
type GenS17 struct { F0 uint8; F1 uint64; F2 uint16 }
type GenS18 struct { F0 int8 }
type GenS19 struct { F0 int64; F1 string; F2 int64 }
type GenS20 struct { F0 int64 }
type GenS21 struct { F0 int64 }
type GenS22 struct { F0 uint32; F1 int8 }
type GenS23 struct { F0 int32; F1 int16 }
type GenS24 struct { F0 bool }
type GenS25 struct { F0 uint; F1 uint8; F2 int16 }
type GenS26 struct { F0 int8; F1 int32 }

func main() {
    var v1 int16 = 0
    v2 := make(map[int]int)
    var v3 float32 = 1e10
    var v4 int16 = (-(v1))
    var v5 int = 9223372036854775807
    v6 := make(map[int]string)
    v1 += -1
    fmt.Println(v3, v5)
    for fi7 := 0; fi7 < 11; fi7++ {
        for fi8 := 0; fi8 < 9; fi8++ {
            v3 = 0.0
            if false {
                v9 := GenS1{F0: ("hello" + "\n\t")}
                _ = v9
            } else {
                v10 := []int{fi7, (v5 * v5)}
                switch "foo bar" {
                case "foo bar":
                    if false {
                        if true {
                            v4 -= v4
                            var v11 uint64 = 42
                            _ = v11
                        }
                        v12 := GenS2{F0: (^(fi8)), F1: "foo bar"}
                        _ = v12
                    }
                default:
                    switch fi8 {
                    case 4294967295:
                        if (0 == 128) {
                            if false {
                                var v13 int8 = 0
                                _ = v13
                            } else {
                                v5 = (-(v5))
                                _ = len(v10)
                                if (!true) {
                                    var v14 int32 = 42
                                    var v15 bool = (54 < 2)
                                    v15 = v15
                                    var v16 uint64 = 42
                                    _ = v14
                                    _ = v15
                                    _ = v16
                                } else {
                                    if (true && ((!true) && false)) {
                                        var v17 int8 = 98
                                        fmt.Println(v5, fi8)
                                        _ = v17
                                    } else {
                                        if len(v10) >= 2 {
                                            v10 = v10[1:len(v10)]
                                        }
                                        if len(v10) >= 2 {
                                            v10 = v10[1:len(v10)]
                                        }
                                    }
                                    var v18 float32 = 3.14159
                                    _ = v18
                                }
                                if false {
                                    var v19 uint8 = 2
                                    _ = v19
                                }
                            }
                            v20 := v6[(fi8 - (fi8 - fi8))]
                            _ = v20
                        }
                    case 26:
                        switch "hello" {
                        case "a":
                            v5 += (fi8 >> ((42) & 63))
                        case "\n\t":
                            if (true && (92 == 4294967295)) {
                                if true {
                                    v10 = append(v10, v5)
                                    _ = len(v10)
                                    v1 -= v4
                                    v4 = (v1 % (((v4 - -1)) | 1))
                                } else {
                                    if true {
                                        if (!("hello" < "hello")) {
                                            var v21 uint16 = 65535
                                            _ = v21
                                        }
                                        var v22 int64 = 772
                                        switch (fi8 ^ (v5 * fi8)) {
                                        case 42:
                                            if true {
                                                if (1 != 42) {
                                                    if false {
                                                        if (0.0 > 0.1) {
                                                            var v23 float32 = (float32(float64(((1.5 + (float32(1e10) + ((v3 * (v3 - v3)) * 3.14159))) * (v3 - 1e10)))) * 0.1)
                                                            if true {
                                                                if len(v10) >= 2 {
                                                                    v10 = v10[1:len(v10)]
                                                                }
                                                            } else {
                                                                fmt.Println(v22, fi8, v22)
                                                                fmt.Println(v3)
                                                                v24 := GenS2{F0: -760, F1: "hello"}
                                                                var v25 float32 = ((float32(float64((v3 * float32(float64(v23))))) * float32(3.14159)) + ((float32(0.0) + v23) * 1e10))
                                                                _ = v24
                                                                _ = v25
                                                            }
                                                            if (26 <= 2) {
                                                                fmt.Println(fi8, fi8)
                                                            }
                                                            v23 += (v23 * (v23 - v23))
                                                            if true {
                                                                var v26 bool = true
                                                                switch ("\n\t" + (("hello" + "") + ("" + "\n\t"))) {
                                                                case "foo bar":
                                                                    v23 += (float32((((1.5 + 0.1) * -2.25) - 3.14159)) + v3)
                                                                case "hello":
                                                                    switch (-(fi7)) {
                                                                    case -903:
                                                                        var v27 int16 = v1
                                                                        _ = v27
                                                                    default:
                                                                        v28 := GenS2{F0: -9223372036854775808, F1: (("a" + "") + ("\n\t" + ("foo bar" + ("" + ("hello" + "a")))))}
                                                                        _ = v28
                                                                    }
                                                                default:
                                                                    fmt.Println(v26, v1, v22)
                                                                }
                                                                _ = v26
                                                            } else {
                                                                v23 = float32((1e10 / 1.0))
                                                                v29 := GenS1{F0: "a"}
                                                                _ = v29
                                                            }
                                                            switch fi8 {
                                                            case 1:
                                                                if (0 >= 255) {
                                                                    fmt.Println(v22)
                                                                } else {
                                                                    var v30 float64 = (float64((v3 + (v23 * v3))) / (1e10 * (0.1 * 1.5)))
                                                                    fmt.Println(fi7, fi8)
                                                                    _ = v30
                                                                }
                                                            case 32768:
                                                                _ = len(v10)
                                                            }
                                                            v23 = float32((1.5 * 0.0))
                                                            if false {
                                                                v31 := GenS1{F0: ("\n\t" + "foo bar")}
                                                                if len(v10) > 0 {
                                                                    v10[0] = fi7
                                                                }
                                                                var v32 uint16 = 655
                                                                _ = v31
                                                                _ = v32
                                                            } else {
                                                                var v33 uint32 = 42
                                                                _ = v33
                                                            }
                                                            var v34 int8 = -104
                                                            if (65535 != 42) {
                                                                fmt.Println(v4, fi8, v23)
                                                                v5 += (v5 >> ((65535) & 63))
                                                            }
                                                            var v35 uint = 1
                                                            var v36 int16 = (v4 % ((v4) | 1))
                                                            _ = v23
                                                            _ = v34
                                                            _ = v35
                                                            _ = v36
                                                        } else {
                                                            v37 := GenS1{F0: "a"}
                                                            if len(v10) >= 2 {
                                                                v10 = v10[1:len(v10)]
                                                            }
                                                            v3 += 0.0
                                                            _ = v37
                                                        }
                                                    } else {
                                                        switch fi8 {
                                                        case -361:
                                                            if false {
                                                                fmt.Println(v22)
                                                                var v38 float32 = (v3 * (v3 + 3.14159))
                                                                v3 = v38
                                                                _ = v38
                                                            }
                                                        case 9223372036854775807:
                                                            switch v5 {
                                                            case 1:
                                                                fmt.Println(v22, v4, fi7)
                                                                fmt.Println(v4)
                                                            case 260:
                                                                v6[(-(fi7))] = ("hello" + "\n\t")
                                                            default:
                                                                fmt.Println(fi7)
                                                            }
                                                        default:
                                                            var v39 int8 = 42
                                                            _ = v39
                                                        }
                                                        v1 -= (v1 ^ 32767)
                                                    }
                                                    v4 -= -1
                                                    _ = len(v2)
                                                    switch "" {
                                                    case "hello":
                                                        var v40 string = ("" + "a")
                                                        _ = v40
                                                    case "\n\t":
                                                        if false {
                                                            delete(v6, fi8)
                                                            fmt.Println(v5, v4)
                                                            _ = len(v10)
                                                        }
                                                    }
                                                    var v41 int32 = -381
                                                    _ = v41
                                                } else {
                                                    v42 := GenS3{F0: 124, F1: (^(v22)), F2: (true || (!true))}
                                                    v1 -= (-(v1))
                                                    _ = v42
                                                }
                                            }
                                        case -2147483648:
                                            if (false && (2 != 2147483647)) {
                                                switch (fi7 | (fi7 & v5)) {
                                                case -9223372036854775808:
                                                    v22 = (-(v22))
                                                case -455:
                                                    if false {
                                                        _ = len(v6)
                                                        fmt.Println(v1)
                                                        v3 = 0.1
                                                    }
                                                default:
                                                    if (!(!(((-23 != -1) || false) && false))) {
                                                        if (("a" == "") || true) {
                                                            switch -2147483648 {
                                                            case -752:
                                                                var v43 bool = false
                                                                if v43 {
                                                                    switch "hello" {
                                                                    case "":
                                                                        if ((2 < 2) && v43) {
                                                                            v4 = v1
                                                                            if v43 {
                                                                                fmt.Println(v5, fi7, v1)
                                                                                switch "hello" {
                                                                                case "foo bar":
                                                                                    v5 -= (fi8 << ((895) & 63))
                                                                                case "a":
                                                                                    if false {
                                                                                        var v44 string = "hello"
                                                                                        fmt.Println(fi8)
                                                                                        _ = v44
                                                                                    }
                                                                                }
                                                                                v45 := GenS3{F0: 127, F1: (v22 * v22), F2: (255 <= 2)}
                                                                                _ = v45
                                                                            }
                                                                            if len(v10) > 0 {
                                                                                v10[0] = (-(fi7))
                                                                            }
                                                                        } else {
                                                                            switch fi7 {
                                                                            case 587:
                                                                                if (!((v22 == -9223372036854775808) && (1 != (v22 & -1)))) {
                                                                                    v46 := GenS4{F0: fi7}
                                                                                    if len(v10) > 0 {
                                                                                        v10[0] = fi7
                                                                                    }
                                                                                    _ = v46
                                                                                }
                                                                            case 42:
                                                                                v47 := GenS2{F0: (fi8 / (((v5 & (^(fi8)))) | 1)), F1: ""}
                                                                                _ = v47
                                                                            case 127:
                                                                                if v43 {
                                                                                    var v48 uint = 2
                                                                                    _ = v48
                                                                                }
                                                                            }
                                                                        }
                                                                    case "hello":
                                                                        switch (fi8 >> ((255) & 63)) {
                                                                        case -919:
                                                                            var v49 int32 = -2147483648
                                                                            _ = v49
                                                                        default:
                                                                            var v50 int64 = (v22 / (((v22 << ((1) & 63))) | 1))
                                                                            _ = v50
                                                                        }
                                                                    case "a":
                                                                        v5 += (v5 & 0)
                                                                        v10 = append(v10, fi8)
                                                                    }
                                                                } else {
                                                                    fmt.Println(v22, v1, v22)
                                                                }
                                                                _ = v43
                                                            case 863:
                                                                var v51 int8 = int8(v22)
                                                                _ = v51
                                                            default:
                                                                if (!(((v3 + (float32((float64(3.14159) - 0.1)) / (v3 + v3))) + float32(-2.25)) <= (v3 + float32(-2.25)))) {
                                                                    var v52 uint64 = 4294967295
                                                                    v53 := v2[v5]
                                                                    if (false || true) {
                                                                        v54 := GenS5{F0: v52}
                                                                        v55 := GenS2{F0: -475, F1: ""}
                                                                        _ = v54
                                                                        _ = v55
                                                                    } else {
                                                                        v53 -= 32767
                                                                        v5 = (-(fi7))
                                                                    }
                                                                    switch 9223372036854775807 {
                                                                    case -126:
                                                                        fmt.Println(v53)
                                                                    case -1:
                                                                        var v56 uint32 = 18
                                                                        _ = v56
                                                                    case 437:
                                                                        var v57 int32 = 0
                                                                        _ = v57
                                                                    }
                                                                    if false {
                                                                        var v58 int16 = v4
                                                                        v59 := GenS4{F0: fi8}
                                                                        fmt.Println(v5)
                                                                        v60 := GenS6{F0: 2147483647, F1: 0, F2: v4}
                                                                        fmt.Println(v4, v4)
                                                                        fmt.Println(v4)
                                                                        _ = v58
                                                                        _ = v59
                                                                        _ = v60
                                                                    } else {
                                                                        var v61 int16 = 0
                                                                        var v62 int = (-(v53))
                                                                        _ = v61
                                                                        _ = v62
                                                                    }
                                                                    _ = v52
                                                                    _ = v53
                                                                }
                                                            }
                                                        } else {
                                                            v63 := GenS7{F0: (fi7 - 42), F1: -768, F2: 1}
                                                            var v64 int64 = (^(v22))
                                                            var v65 bool = (!(!(338 != (v64 + (-(v22))))))
                                                            if (v4 > (v4 & v1)) {
                                                                v66 := GenS2{F0: int(v64), F1: "\n\t"}
                                                                _ = v66
                                                            } else {
                                                                if (2147483647 >= int32(fi8)) {
                                                                    if false {
                                                                        var v67 int32 = -2147483648
                                                                        _ = v67
                                                                    }
                                                                    if v65 {
                                                                        v65 = true
                                                                        var v68 uint64 = 4294967295
                                                                        _ = v68
                                                                    }
                                                                    switch v5 {
                                                                    case 42:
                                                                        v69 := GenS8{F0: ("\n\t" + "")}
                                                                        _ = v69
                                                                    }
                                                                } else {
                                                                    if (!false) {
                                                                        v1 = 1000
                                                                        var v70 uint8 = 1
                                                                        v10 = append(v10, 4294967295)
                                                                        _ = v70
                                                                    } else {
                                                                        if v65 {
                                                                            var v71 uint8 = 128
                                                                            _ = v71
                                                                        } else {
                                                                            switch (((("" + "foo bar") + "hello") + (("\n\t" + "\n\t") + "foo bar")) + "hello") {
                                                                            case "":
                                                                                v65 = (((v65 && v65) && v65) && (127 <= 255))
                                                                            case "foo bar":
                                                                                if (42 < 2) {
                                                                                    v64 = (v64 / ((42) | 1))
                                                                                    var v72 int16 = (v4 - (v4 + v1))
                                                                                    _ = v72
                                                                                }
                                                                            case "hello":
                                                                                if true {
                                                                                    _ = len(v6)
                                                                                    fmt.Println(v64)
                                                                                } else {
                                                                                    if v65 {
                                                                                        switch "" {
                                                                                        case "":
                                                                                            v73 := GenS9{F0: 0}
                                                                                            _ = v73
                                                                                        default:
                                                                                            var v74 string = ("foo bar" + "foo bar")
                                                                                            _ = v74
                                                                                        }
                                                                                    }
                                                                                    v64 = 9223372036854775807
                                                                                }
                                                                            default:
                                                                                switch "\n\t" {
                                                                                case "a":
                                                                                    if true {
                                                                                        v22 -= (v64 & v64)
                                                                                        var v75 string = "foo bar"
                                                                                        _ = v75
                                                                                    } else {
                                                                                        v22 += (v64 | v64)
                                                                                        switch "" {
                                                                                        case "":
                                                                                            var v76 uint64 = 314
                                                                                            v77 := GenS10{F0: 42, F1: v4}
                                                                                            _ = v76
                                                                                            _ = v77
                                                                                        }
                                                                                    }
                                                                                }
                                                                            }
                                                                        }
                                                                        v3 -= ((v3 - (1e10 + v3)) + v3)
                                                                    }
                                                                }
                                                                switch "" {
                                                                case "a":
                                                                    var v78 uint8 = 127
                                                                    _ = v78
                                                                case "":
                                                                    v79 := GenS11{F0: 2147483647, F1: 65535}
                                                                    _ = v79
                                                                }
                                                                v22 += (v64 % ((v64) | 1))
                                                                var v80 int32 = 200
                                                                var v81 float32 = v3
                                                                _ = v80
                                                                _ = v81
                                                            }
                                                            _ = v63
                                                            _ = v64
                                                            _ = v65
                                                        }
                                                        if false {
                                                            var v82 float32 = v3
                                                            v5 = 32768
                                                            v83 := GenS12{F0: 127, F1: 0}
                                                            _ = v82
                                                            _ = v83
                                                        } else {
                                                            var v84 int = (v5 ^ fi7)
                                                            _ = v84
                                                        }
                                                        v10 = append(v10, (-(fi7)))
                                                        v4 = (^(v1))
                                                    }
                                                }
                                                if (v3 != v3) {
                                                    v85 := GenS13{F0: 4294967295, F1: 70}
                                                    if (65535 <= 255) {
                                                        var v86 uint8 = 2
                                                        _ = v86
                                                    } else {
                                                        v4 += (v1 | (v1 & v4))
                                                        switch (("" + "") + "\n\t") {
                                                        case "a":
                                                            _ = len(v2)
                                                        case "\n\t":
                                                            var v87 uint8 = 1
                                                            _ = v87
                                                        case "foo bar":
                                                            var v88 bool = (!(false || (42 <= -128)))
                                                            _ = v88
                                                        default:
                                                            switch (fi7 & (v5 | fi8)) {
                                                            case -1:
                                                                switch "" {
                                                                case "":
                                                                    _ = len(v6)
                                                                }
                                                            default:
                                                                if false {
                                                                    v89 := GenS14{F0: (v22 | (^(v22))), F1: -306, F2: 8}
                                                                    var v90 float64 = -2.25
                                                                    var v91 string = ""
                                                                    delete(v6, v5)
                                                                    _ = v89
                                                                    _ = v90
                                                                    _ = v91
                                                                }
                                                            }
                                                        }
                                                    }
                                                    if false {
                                                        v22 -= (^(v22))
                                                    }
                                                    _ = v85
                                                } else {
                                                    switch "foo bar" {
                                                    case "hello":
                                                        var v92 float64 = -2.25
                                                        _ = v92
                                                    case "foo bar":
                                                        v93 := GenS15{F0: 255, F1: v1}
                                                        _ = v93
                                                    }
                                                }
                                            }
                                        case 32768:
                                            if ((("\n\t" <= "") || (64 < 4294967295)) && ((42 <= 2) && true)) {
                                                fmt.Println(fi7)
                                                v94 := v2[(^(fi7))]
                                                _ = v94
                                            }
                                        default:
                                            var v95 uint = 1
                                            _ = v95
                                        }
                                        _ = v22
                                    }
                                    var v96 int64 = -9223372036854775808
                                    v4 += v4
                                    _ = v96
                                }
                            } else {
                                fmt.Println(v3, v4, v1)
                                if false {
                                    if ((4294967295 != 0) && (1 > 42)) {
                                        var v97 int32 = -359
                                        v98 := GenS16{F0: v97, F1: fi8}
                                        _ = v97
                                        _ = v98
                                    } else {
                                        fmt.Println(v5)
                                        if (!false) {
                                            if ((((v3 / v3) * float32((3.14159 / float64(-2.25)))) * v3) > v3) {
                                                v99 := GenS10{F0: 4294967295, F1: -668}
                                                v3 = ((v3 * (((((v3 + 1.5) - -2.25) - (v3 - v3)) + (v3 * v3)) - 1.5)) - v3)
                                                if len(v10) > 0 {
                                                    v10[0] = 128
                                                }
                                                v10 = append(v10, fi8)
                                                _ = v99
                                            } else {
                                                v1 = v4
                                                switch v5 {
                                                case 32767:
                                                    v100 := GenS17{F0: 127, F1: 4294967295, F2: 42}
                                                    _ = v100
                                                case 811:
                                                    if len(v10) >= 2 {
                                                        v10 = v10[1:len(v10)]
                                                    }
                                                }
                                                var v101 int = (fi8 / (((fi7 | v5)) | 1))
                                                switch 1 {
                                                case 44:
                                                    v102 := GenS18{F0: 100}
                                                    _ = v102
                                                default:
                                                    v103 := v6[v101]
                                                    _ = v103
                                                }
                                                _ = v101
                                            }
                                            fmt.Println(fi7)
                                        }
                                        if (("foo bar" > (("" + "") + (("" + "hello") + "\n\t"))) && ((65535 != 4294967295) && true)) {
                                            fmt.Println(fi8)
                                            v104 := GenS19{F0: 395, F1: "foo bar", F2: 42}
                                            delete(v2, v5)
                                            _ = v104
                                        } else {
                                            switch v5 {
                                            case 0:
                                                v5 -= v5
                                            case -9223372036854775808:
                                                var v105 int32 = -1
                                                _ = v105
                                            default:
                                                v5 = (v5 * fi7)
                                            }
                                            v4 = v4
                                        }
                                    }
                                    var v106 int = 127
                                    _ = v106
                                }
                            }
                        case "foo bar":
                            v1 = (v4 % (((v4 * v1)) | 1))
                        }
                    }
                }
                _ = v10
            }
            _ = fi8
        }
        v2[v5] = (v5 ^ 42)
        v5 = fi7
        _ = fi7
    }
    switch "foo bar" {
    case "foo bar":
        for fi107 := 0; fi107 < 13; fi107++ {
            switch -347 {
            case 0:
                for fi108 := 0; fi108 < 5; fi108++ {
                    v4 -= (v1 % ((-32768) | 1))
                    switch v5 {
                    case 976:
                        var v109 int16 = v4
                        _ = v109
                    }
                    switch (fi108 & fi108) {
                    case 616:
                        v6[(fi107 - -24)] = "a"
                    case 1:
                        if true {
                            v1 = (v4 ^ (-(v4)))
                            fmt.Println(v5, v5)
                            v3 = float32(1.5)
                            var v110 uint16 = 2
                            _ = v110
                        } else {
                            var v111 uint32 = 686
                            _ = v111
                        }
                    default:
                        var v112 int8 = -16
                        _ = v112
                    }
                    _ = fi108
                }
            case 1:
                switch "foo bar" {
                case "\n\t":
                    var v113 string = ("foo bar" + "")
                    _ = v113
                case "":
                    v4 = 42
                }
            }
            var v114 uint16 = 902
            _ = fi107
            _ = v114
        }
    case "":
        switch (^(v5)) {
        case 42:
            v115 := GenS16{F0: -1, F1: 32768}
            _ = v115
        default:
            v4 += v1
        }
    case "hello":
        v3 -= 1.5
    default:
        v116 := []int{}
        _ = v116
    }
    for fi117 := 0; fi117 < 7; fi117++ {
        var v118 float64 = 0.0
        _ = fi117
        _ = v118
    }
    v119 := []int{(v5 % (((v5 - v5)) | 1)), 127, v5}
    var v120 uint64 = 981
    v1 += v4
    v121 := GenS20{F0: 42}
    if (!(!(9223372036854775807 == 9223372036854775807))) {
        for fi122 := 0; fi122 < 6; fi122++ {
            if (557 >= 1) {
                v123 := GenS14{F0: 833, F1: 2147483647, F2: 45}
                for fi124 := 0; fi124 < 4; fi124++ {
                    if (!false) {
                        v125 := GenS21{F0: 0}
                        _ = len(v2)
                        _ = v125
                    } else {
                        v126 := GenS15{F0: 1, F1: (v4 & v1)}
                        fmt.Println(v5)
                        _ = v126
                    }
                    var v127 uint16 = 893
                    if true {
                        v128 := GenS22{F0: 4294967295, F1: 42}
                        v129 := GenS23{F0: -2147483648, F1: 0}
                        v3 = v3
                        _ = v128
                        _ = v129
                    } else {
                        var v130 float64 = ((3.14159 - (float64((v3 * v3)) / float64(v3))) + 1.5)
                        v1 += v1
                        var v131 int = fi122
                        _ = v130
                        _ = v131
                    }
                    _ = fi124
                    _ = v127
                }
                v6[(v5 * (fi122 * (fi122 / (((fi122 / (((fi122 ^ (-(fi122)))) | 1))) | 1))))] = "foo bar"
                _ = v123
            } else {
                v120 -= 18446744073709551615
                if true {
                    v1 = v4
                    if true {
                        fmt.Println(fi122, v5)
                        var v132 int = (v5 ^ v5)
                        switch fi122 {
                        case -9223372036854775808:
                            v132 = 307
                        case 127:
                            if false {
                                v6[127] = ("" + ((("\n\t" + "a") + "") + "hello"))
                                var v133 int32 = 42
                                _ = v133
                            }
                        case 86:
                            var v134 int8 = 0
                            _ = v134
                        default:
                            var v135 int64 = -1
                            _ = v135
                        }
                        _ = v132
                    }
                } else {
                    var v136 string = ("foo bar" + (("\n\t" + (("" + "\n\t") + "a")) + ""))
                    _ = v136
                }
            }
            for fi137 := 0; fi137 < 2; fi137++ {
                if ((false && (-128 > -1)) && false) {
                    _ = len(v2)
                } else {
                    v138 := GenS12{F0: 127, F1: 65535}
                    if (1 >= 0) {
                        if false {
                            if true {
                                var v139 float64 = (0.1 / 0.1)
                                if (4294967295 == 2147483648) {
                                    var v140 int64 = -25
                                    _ = v140
                                } else {
                                    fmt.Println(fi137, v3)
                                    v139 -= float64(v3)
                                }
                                var v141 int16 = 32767
                                var v142 int = v5
                                _ = v139
                                _ = v141
                                _ = v142
                            } else {
                                var v143 uint64 = (v120 >> ((18446744073709551615) & 63))
                                _ = v143
                            }
                            if true {
                                fmt.Println(v1)
                                v1 -= (-(v1))
                                switch "foo bar" {
                                case "a":
                                    if true {
                                        v5 += fi122
                                        if true {
                                            switch (fi137 ^ v5) {
                                            case -924:
                                                v144 := GenS24{F0: false}
                                                _ = v144
                                            case 844:
                                                if true {
                                                    v1 = (v4 / (((v4 * v1)) | 1))
                                                }
                                            case 9223372036854775807:
                                                v145 := GenS8{F0: "hello"}
                                                _ = v145
                                            }
                                            v5 -= (fi137 ^ 128)
                                            v146 := GenS8{F0: "foo bar"}
                                            _ = v146
                                        }
                                    } else {
                                        var v147 int8 = 60
                                        _ = v147
                                    }
                                case "":
                                    v1 += 42
                                case "foo bar":
                                    switch 464 {
                                    case 23:
                                        if true {
                                            switch 1 {
                                            case -9223372036854775808:
                                                if (("hello" + "") != "a") {
                                                    fmt.Println(v1, v5, v5)
                                                    v3 = v3
                                                    v3 += v3
                                                    switch "a" {
                                                    case "foo bar":
                                                        fmt.Println(fi122, v5)
                                                    case "hello":
                                                        v119 = append(v119, fi137)
                                                        v120 -= v120
                                                    }
                                                } else {
                                                    var v148 uint32 = 0
                                                    fmt.Println(v120, v1)
                                                    _ = v148
                                                }
                                            case 1:
                                                var v149 uint16 = 310
                                                _ = v149
                                            default:
                                                var v150 float64 = 1.5
                                                _ = v150
                                            }
                                            v2[2147483647] = (-(fi122))
                                            var v151 int32 = 0
                                            _ = v151
                                        }
                                    }
                                default:
                                    var v152 uint8 = 42
                                    _ = v152
                                }
                                var v153 int8 = 1
                                _ = v153
                            }
                        } else {
                            _ = len(v2)
                        }
                        v119 = append(v119, v5)
                        v1 -= (v1 ^ (v4 ^ v4))
                        if true {
                            fmt.Println(v4, v120, fi122)
                            switch "\n\t" {
                            case "\n\t":
                                v1 = v4
                            }
                            var v154 int16 = v1
                            _ = v154
                        }
                        fmt.Println(fi122, fi122, v4)
                        if true {
                            v4 += v1
                        }
                    }
                    switch "hello" {
                    case "foo bar":
                        v119 = append(v119, fi137)
                    case "hello":
                        v155 := GenS25{F0: 42, F1: 234, F2: v4}
                        _ = v155
                    }
                    _ = v138
                }
                var v156 int8 = 42
                var v157 uint64 = 774
                var v158 int32 = -2147483648
                _ = fi137
                _ = v156
                _ = v157
                _ = v158
            }
            _ = fi122
        }
        _ = len(v119)
    } else {
        if (42 >= 0) {
            var v159 uint16 = 0
            for fi160 := 0; fi160 < 18; fi160++ {
                var v161 uint16 = v159
                var v162 bool = true
                var v163 float32 = (1e10 / -2.25)
                _ = fi160
                _ = v161
                _ = v162
                _ = v163
            }
            _ = v159
        } else {
            var v164 int32 = -1
            v120 -= 90
            _ = v164
        }
        var v165 uint16 = 42
        switch ("hello" + (("" + "") + "hello")) {
        case "\n\t":
            v1 = (v4 << ((0) & 63))
        case "foo bar":
            var v166 int = (v5 ^ (-(v5)))
            _ = v166
        case "a":
            if true {
                if true {
                    for fi167 := 0; fi167 < 10; fi167++ {
                        if false {
                            v120 = v120
                            fmt.Println(v5)
                        }
                        _ = fi167
                    }
                    v1 -= (-(v4))
                } else {
                    if false {
                        v5 = 1
                    } else {
                        if (2147483648 < uint32(v165)) {
                            if false {
                                v168 := GenS26{F0: 42, F1: 1}
                                _ = v168
                            }
                        }
                        _ = len(v119)
                    }
                }
            }
        default:
            v4 = 1000
        }
        _ = v165
    }
    v169 := GenS4{F0: (v5 + 42)}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v119:", len(v119), v119)
    fmt.Println("v120:", v120)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v119
    _ = v120
    _ = v121
    _ = v169
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
