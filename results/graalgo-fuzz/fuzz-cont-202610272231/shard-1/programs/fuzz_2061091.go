package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 bool; F1 int }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int32; F1 uint8 }
type GenS6 struct { F0 uint64; F1 uint16 }
type GenS7 struct { F0 uint32; F1 int32 }
type GenS8 struct { F0 int; F1 int }
type GenS9 struct { F0 int32; F1 uint64; F2 uint16 }
type GenS10 struct { F0 uint; F1 uint64 }
type GenS11 struct { F0 int32; F1 uint64 }
type GenS12 struct { F0 bool; F1 bool; F2 uint16 }
type GenS13 struct { F0 uint8; F1 uint }
type GenS14 struct { F0 int64; F1 bool }
type GenS15 struct { F0 bool; F1 uint; F2 int }
func genFn0(p0 uint32, p1 int64, p2 int) int {
    v6 := []int{(p2 & (p2 % (((p2 * p2)) | 1))), (p2 << ((789) & 63))}
    fmt.Println(p1, p0)
    v6 = append(v6, (^(p2)))
    v7 := GenS1{F0: 127}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return (p2 + (p2 / ((42) | 1)))
}

func main() {
    v1 := make(map[int]string)
    var v2 float32 = 3.14159
    var v3 uint32 = 42
    v4 := []string{"a", ("hello" + "\n\t"), ("\n\t" + ""), "\n\t"}
    v5 := make(map[string]string)
    v1[127] = ""
    delete(v5, ("\n\t" + ""))
    v3 = v3
    v8 := GenS1{F0: 100}
    v2 -= v2
    v2 += float32(1.5)
    for fi9 := 0; fi9 < 14; fi9++ {
        fmt.Println(v2)
        switch "foo bar" {
        case "\n\t":
            var v10 string = "hello"
            _ = v10
        }
        for fi11 := 0; fi11 < 4; fi11++ {
            if false {
                v2 += v2
                fmt.Println(fi9)
                var v12 bool = (false && true)
                _ = v12
            } else {
                if (true && false) {
                    var v13 int8 = 127
                    switch ((("hello" + "hello") + ("\n\t" + "\n\t")) + (("a" + "foo bar") + "foo bar")) {
                    case "hello":
                        if true {
                            v13 = v13
                            switch -852 {
                            case -2147483648:
                                var v14 uint16 = 42
                                fmt.Println(v2)
                                v13 = v13
                                _ = v14
                            case 533:
                                v3 = (^(v3))
                            case -464:
                                v3 += v3
                            default:
                                fmt.Println(v13, v3, v13)
                            }
                            if true {
                                var v15 float32 = v2
                                _ = v15
                            } else {
                                v5[""] = "foo bar"
                            }
                        } else {
                            var v16 int = fi11
                            fmt.Println(fi11, fi11)
                            var v17 float64 = -2.25
                            var v18 int8 = 46
                            _ = v16
                            _ = v17
                            _ = v18
                        }
                    default:
                        fmt.Println(v13)
                    }
                    switch (fi9 & (fi11 >> ((174) & 63))) {
                    case 128:
                        var v19 string = "hello"
                        _ = v19
                    case -664:
                        var v20 string = ("" + ("" + "a"))
                        _ = v20
                    case 314:
                        if false {
                            if (false || false) {
                                if false {
                                    switch fi9 {
                                    case 231:
                                        if (fi9 != fi9) {
                                            var v21 uint64 = 0
                                            switch fi11 {
                                            case 32767:
                                                fmt.Println(v3, v21, v2)
                                            case -9223372036854775808:
                                                var v22 int = (fi11 << ((18446744073709551615) & 63))
                                                _ = v22
                                            case 1:
                                                switch "foo bar" {
                                                case "a":
                                                    v23 := GenS1{F0: v13}
                                                    _ = v23
                                                case "":
                                                    if (!false) {
                                                        if true {
                                                            if ("foo bar" == ("\n\t" + ("hello" + ""))) {
                                                                v13 -= -1
                                                                fmt.Println(v13, v3, v2)
                                                            } else {
                                                                switch 9223372036854775807 {
                                                                case 648:
                                                                    var v24 uint = 378
                                                                    _ = v24
                                                                case 4294967295:
                                                                    fmt.Println(fi11)
                                                                case -1:
                                                                    var v25 int16 = -1
                                                                    _ = v25
                                                                default:
                                                                    if (true || true) {
                                                                        var v26 int32 = 42
                                                                        var v27 uint64 = v21
                                                                        fmt.Println(fi9, fi9, v21)
                                                                        _ = v26
                                                                        _ = v27
                                                                    } else {
                                                                        var v28 int64 = -1
                                                                        if (693 <= 626) {
                                                                            var v29 int32 = -1
                                                                            _ = v29
                                                                        } else {
                                                                            var v30 float32 = v2
                                                                            switch "hello" {
                                                                            case "foo bar":
                                                                                var v31 float64 = 3.14159
                                                                                _ = v31
                                                                            case "hello":
                                                                                var v32 int = (fi11 | (fi9 & (fi9 % ((fi9) | 1))))
                                                                                _ = v32
                                                                            default:
                                                                                if false {
                                                                                    var v33 uint8 = 37
                                                                                    v34 := GenS1{F0: v13}
                                                                                    v4 = append(v4, ("hello" + ("hello" + ("a" + "foo bar"))))
                                                                                    _ = v33
                                                                                    _ = v34
                                                                                }
                                                                            }
                                                                            if (v3 == (v3 & (^(v3)))) {
                                                                                v5["hello"] = "a"
                                                                                v13 -= 127
                                                                            } else {
                                                                                switch (fi9 | genFn0((v3 * (v3 / ((v3) | 1))), (v28 | (v28 << ((175) & 63))), fi11)) {
                                                                                case 0:
                                                                                    var v35 uint = uint(v21)
                                                                                    _ = v35
                                                                                case 1:
                                                                                    if len(v4) > 0 {
                                                                                        v4[0] = ""
                                                                                    }
                                                                                case -92:
                                                                                    delete(v1, 4294967295)
                                                                                }
                                                                                if true {
                                                                                    if false {
                                                                                        switch "" {
                                                                                        case "\n\t":
                                                                                            var v36 int32 = -2147483648
                                                                                            _ = v36
                                                                                        case "a":
                                                                                            if true {
                                                                                                v30 = (v2 + (float32(0.0) + v2))
                                                                                            } else {
                                                                                                v13 = (v13 - v13)
                                                                                                v37 := GenS1{F0: 1}
                                                                                                var v38 bool = (("\n\t" + (((("" + "foo bar") + "foo bar") + ("\n\t" + "hello")) + "foo bar")) <= "")
                                                                                                if v38 {
                                                                                                    switch ("\n\t" + "") {
                                                                                                    case "foo bar":
                                                                                                        var v39 string = "a"
                                                                                                        _ = v39
                                                                                                    default:
                                                                                                        if v38 {
                                                                                                            var v40 uint8 = 128
                                                                                                            _ = v40
                                                                                                        }
                                                                                                    }
                                                                                                    fmt.Println(v28)
                                                                                                    v2 += -2.25
                                                                                                } else {
                                                                                                    switch "" {
                                                                                                    case "a":
                                                                                                        _ = len(v4)
                                                                                                    }
                                                                                                    fmt.Println(fi9)
                                                                                                    v28 -= (v28 & (v28 + (v28 | v28)))
                                                                                                }
                                                                                                v13 -= v13
                                                                                                fmt.Println(v2)
                                                                                                _ = v37
                                                                                                _ = v38
                                                                                            }
                                                                                        default:
                                                                                            v41 := GenS1{F0: -1}
                                                                                            _ = v41
                                                                                        }
                                                                                        v4 = append(v4, (("" + (("a" + "hello") + ("\n\t" + "hello"))) + "\n\t"))
                                                                                    } else {
                                                                                        fmt.Println(v21, v13, v28)
                                                                                        if false {
                                                                                            var v42 int8 = v13
                                                                                            _ = v42
                                                                                        } else {
                                                                                            var v43 int8 = (^(v13))
                                                                                            var v44 int8 = 0
                                                                                            if len(v4) >= 2 {
                                                                                                v4 = v4[1:len(v4)]
                                                                                            }
                                                                                            _ = v43
                                                                                            _ = v44
                                                                                        }
                                                                                        v45 := GenS2{F0: v28}
                                                                                        fmt.Println(v2, v30, fi11)
                                                                                        _ = v45
                                                                                    }
                                                                                }
                                                                            }
                                                                            _ = v30
                                                                        }
                                                                        _ = v28
                                                                    }
                                                                    var v46 int8 = -74
                                                                    _ = v46
                                                                }
                                                                if true {
                                                                    v3 -= 701
                                                                    v47 := GenS1{F0: (v13 - v13)}
                                                                    _ = v47
                                                                }
                                                            }
                                                            var v48 uint = 1
                                                            _ = v48
                                                        }
                                                        var v49 uint8 = 42
                                                        var v50 uint32 = (v3 << ((325) & 63))
                                                        v3 -= (v3 & uint32(v49))
                                                        _ = v49
                                                        _ = v50
                                                    } else {
                                                        v4 = append(v4, "hello")
                                                    }
                                                case "foo bar":
                                                    var v51 int8 = (v13 - v13)
                                                    _ = v51
                                                default:
                                                    if false {
                                                        v52 := GenS1{F0: (v13 & -128)}
                                                        _ = v52
                                                    } else {
                                                        v53 := GenS2{F0: 57}
                                                        _ = v53
                                                    }
                                                    var v54 int8 = 117
                                                    switch "foo bar" {
                                                    case "hello":
                                                        v55 := GenS1{F0: 127}
                                                        _ = v55
                                                    case "foo bar":
                                                        switch genFn0(v3, -9223372036854775808, (fi11 * fi11)) {
                                                        case 957:
                                                            _ = len(v4)
                                                        case 32768:
                                                            v56 := GenS3{F0: (2147483647 == 42), F1: (fi9 | (^(fi11)))}
                                                            _ = v56
                                                        default:
                                                            v3 = v3
                                                        }
                                                    case "":
                                                        switch -9223372036854775808 {
                                                        case 1:
                                                            v57 := GenS2{F0: -52}
                                                            _ = v57
                                                        case 4294967295:
                                                            if (((((-2.25 + 0.0) + (1.5 - (0.1 - (0.1 / 0.1)))) - 1e10) * (1e10 - 0.0)) > ((((-2.25 / (3.14159 / ((1e10 * float64(float32(1.5))) - 0.0))) + float64(v2)) * (((-2.25 - 3.14159) / 1.0) + ((3.14159 - (1.5 - 3.14159)) * 1e10))) / 1.0)) {
                                                                _ = len(v1)
                                                                v2 = v2
                                                            }
                                                        case 0:
                                                            switch ("" + (("a" + "a") + "")) {
                                                            case "foo bar":
                                                                switch fi11 {
                                                                case -9223372036854775808:
                                                                    if len(v4) >= 2 {
                                                                        v4 = v4[1:len(v4)]
                                                                    }
                                                                case 128:
                                                                    v58 := GenS4{F0: 32767}
                                                                    v13 = (v13 - -1)
                                                                    var v59 uint16 = 1
                                                                    _ = v58
                                                                    _ = v59
                                                                default:
                                                                    switch ("\n\t" + "") {
                                                                    case "\n\t":
                                                                        switch ("hello" + "foo bar") {
                                                                        case "":
                                                                            v60 := GenS5{F0: 0, F1: 2}
                                                                            _ = v60
                                                                        case "\n\t":
                                                                            var v61 string = ("foo bar" + "")
                                                                            switch 913 {
                                                                            case 0:
                                                                                if true {
                                                                                    v62 := GenS3{F0: (true || (1 < 0)), F1: fi11}
                                                                                    _ = v62
                                                                                } else {
                                                                                    switch 127 {
                                                                                    case 32768:
                                                                                        var v63 string = ""
                                                                                        switch v61 {
                                                                                        case "\n\t":
                                                                                            switch fi9 {
                                                                                            case 0:
                                                                                                var v64 bool = true
                                                                                                _ = v64
                                                                                            case 42:
                                                                                                if len(v4) > 0 {
                                                                                                    v4[0] = v63
                                                                                                }
                                                                                                v61 = ""
                                                                                                _ = len(v1)
                                                                                            case -1:
                                                                                                fmt.Println(v21, v54)
                                                                                                fmt.Println(v54)
                                                                                                var v65 uint = 96
                                                                                                v66 := GenS4{F0: 42}
                                                                                                switch fi11 {
                                                                                                case 4294967295:
                                                                                                    var v67 uint32 = v3
                                                                                                    _ = v67
                                                                                                default:
                                                                                                    v68 := GenS5{F0: -2147483648, F1: 2}
                                                                                                    _ = v68
                                                                                                }
                                                                                                _ = v65
                                                                                                _ = v66
                                                                                            default:
                                                                                                switch "\n\t" {
                                                                                                case "a":
                                                                                                    delete(v1, fi9)
                                                                                                }
                                                                                            }
                                                                                            if ((-1 < 1000) || true) {
                                                                                                fmt.Println(fi9)
                                                                                                if len(v4) >= 2 {
                                                                                                    v4 = v4[1:len(v4)]
                                                                                                }
                                                                                                fmt.Println(v61, v2)
                                                                                            }
                                                                                        case "foo bar":
                                                                                            if len(v4) > 0 {
                                                                                                v4[0] = (v61 + ((((("a" + v63) + "hello") + v61) + ("" + ("foo bar" + "a"))) + "a"))
                                                                                            }
                                                                                        default:
                                                                                            _ = len(v5)
                                                                                        }
                                                                                        v3 -= v3
                                                                                        _ = v63
                                                                                    default:
                                                                                        v3 -= 0
                                                                                    }
                                                                                    _ = len(v5)
                                                                                }
                                                                            case -940:
                                                                                v69 := GenS6{F0: (v21 / ((v21) | 1)), F1: 42}
                                                                                var v70 uint = 729
                                                                                var v71 bool = true
                                                                                _ = v69
                                                                                _ = v70
                                                                                _ = v71
                                                                            case 459:
                                                                                if (42 >= 579) {
                                                                                    if true {
                                                                                        v61 = ("\n\t" + v61)
                                                                                    }
                                                                                }
                                                                            }
                                                                            fmt.Println(fi9, fi9)
                                                                            _ = v61
                                                                        case "hello":
                                                                            _ = len(v4)
                                                                        }
                                                                    case "":
                                                                        v72 := GenS2{F0: 42}
                                                                        _ = v72
                                                                    }
                                                                }
                                                            }
                                                        }
                                                    default:
                                                        if (!false) {
                                                            v13 -= (v54 >> ((515) & 63))
                                                            v3 = (v3 / (((v3 - v3)) | 1))
                                                        }
                                                    }
                                                    _ = v54
                                                }
                                            }
                                            _ = v21
                                        }
                                    default:
                                        var v73 int64 = 9223372036854775807
                                        _ = v73
                                    }
                                    v74 := GenS5{F0: -2147483648, F1: 1}
                                    if true {
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                        v4 = append(v4, ("" + "foo bar"))
                                        if ("" > "hello") {
                                            if (v3 <= v3) {
                                                if (32767 > (fi9 + (^(fi9)))) {
                                                    var v75 uint = 7
                                                    v76 := GenS7{F0: v3, F1: 42}
                                                    v77 := GenS5{F0: 2147483647, F1: uint8(v3)}
                                                    _ = v75
                                                    _ = v76
                                                    _ = v77
                                                } else {
                                                    var v78 int64 = -9223372036854775808
                                                    var v79 float32 = v2
                                                    if (int64(fi11) <= (-(v78))) {
                                                        _ = len(v1)
                                                        v79 = v2
                                                        fmt.Println(v3, fi9, v2)
                                                        _ = len(v1)
                                                        if len(v4) > 0 {
                                                            v4[0] = "\n\t"
                                                        }
                                                    }
                                                    _ = v78
                                                    _ = v79
                                                }
                                            }
                                        }
                                        v2 = 0.0
                                    } else {
                                        v80 := GenS8{F0: (-(fi11)), F1: 931}
                                        delete(v1, (fi9 >> ((42) & 63)))
                                        _ = v80
                                    }
                                    _ = v74
                                } else {
                                    v13 = (^(v13))
                                    var v81 bool = (!false)
                                    _ = v81
                                }
                                if true {
                                    fmt.Println(v2)
                                    if true {
                                        switch 128 {
                                        case 4294967295:
                                            switch "" {
                                            case "hello":
                                                v82 := GenS5{F0: 154, F1: 255}
                                                _ = v82
                                            case "\n\t":
                                                if (!(3.14159 >= (1.5 - 0.0))) {
                                                    var v83 int8 = (v13 + 42)
                                                    _ = v83
                                                }
                                            }
                                        }
                                        var v84 uint = 347
                                        _ = v84
                                    }
                                } else {
                                    var v85 int16 = 32767
                                    fmt.Println(fi11, v85, v2)
                                    var v86 int32 = 153
                                    _ = v85
                                    _ = v86
                                }
                                if ((1e10 + float64(1.5)) >= ((3.14159 + (1.5 / 3.14159)) / 1.5)) {
                                    _ = len(v1)
                                }
                                switch 127 {
                                case 4294967295:
                                    switch (fi9 << ((4294967295) & 63)) {
                                    case 127:
                                        v87 := GenS4{F0: 0}
                                        _ = v87
                                    case 128:
                                        if ((!false) || true) {
                                            v2 = v2
                                            v13 += v13
                                            if true {
                                                _ = len(v4)
                                                v2 += 0.1
                                                if false {
                                                    if false {
                                                        v3 += v3
                                                        var v88 uint = 694
                                                        _ = v88
                                                    } else {
                                                        var v89 float64 = float64(v2)
                                                        v13 = 42
                                                        v89 += float64(v2)
                                                        _ = v89
                                                    }
                                                    var v90 bool = true
                                                    v5[("a" + ("" + ("foo bar" + "a")))] = ""
                                                    _ = v90
                                                } else {
                                                    v1[genFn0(v3, -1, fi9)] = "foo bar"
                                                }
                                                v91 := GenS8{F0: 9223372036854775807, F1: genFn0(65535, -1, fi9)}
                                                var v92 uint8 = 128
                                                _ = v91
                                                _ = v92
                                            } else {
                                                if (-32768 < -32768) {
                                                    var v93 int32 = 2147483647
                                                    _ = v93
                                                }
                                                var v94 int16 = 1000
                                                v95 := GenS9{F0: 1, F1: 4294967295, F2: 983}
                                                _ = v94
                                                _ = v95
                                            }
                                            _ = len(v4)
                                        } else {
                                            fmt.Println(v13, fi9)
                                            fmt.Println(v2, fi11)
                                            fmt.Println(fi11)
                                        }
                                    case 1:
                                        v96 := GenS10{F0: 42, F1: 720}
                                        _ = v96
                                    }
                                }
                                switch "hello" {
                                case "":
                                    fmt.Println(v2)
                                case "\n\t":
                                    v97 := GenS4{F0: 42}
                                    _ = v97
                                }
                            }
                            var v98 uint32 = v3
                            _ = v98
                        } else {
                            if (-857 != 9223372036854775807) {
                                v99 := GenS11{F0: 42, F1: 4294967295}
                                v5[(("\n\t" + "") + ("foo bar" + "\n\t"))] = ""
                                _ = v99
                            } else {
                                v100 := GenS12{F0: false, F1: ("\n\t" == (("foo bar" + "a") + "\n\t")), F2: 0}
                                if (88 > 128) {
                                    if true {
                                        v101 := GenS5{F0: 42, F1: 2}
                                        var v102 string = "foo bar"
                                        fmt.Println(v2, fi11)
                                        _ = v101
                                        _ = v102
                                    } else {
                                        fmt.Println(v3, v13)
                                        var v103 float64 = 3.14159
                                        _ = v103
                                    }
                                    var v104 int = fi11
                                    _ = v104
                                } else {
                                    if (32767 >= 626) {
                                        if (!(128 != 127)) {
                                            if true {
                                                v2 = v2
                                                var v105 int = fi11
                                                v4 = append(v4, "foo bar")
                                                _ = v105
                                            } else {
                                                fmt.Println(v13, v2)
                                                fmt.Println(fi9)
                                            }
                                            v2 = (v2 / (v2 - v2))
                                        }
                                    } else {
                                        fmt.Println(v3, v3)
                                        v2 += v2
                                    }
                                    var v106 float32 = 0.0
                                    switch ("foo bar" + "foo bar") {
                                    case "foo bar":
                                        _ = len(v4)
                                        var v107 int32 = -2147483648
                                        _ = v107
                                    case "":
                                        v108 := GenS1{F0: (v13 - (v13 >> ((531) & 63)))}
                                        _ = v108
                                    case "a":
                                        var v109 int = fi11
                                        _ = v109
                                    }
                                    _ = v106
                                }
                                v13 += 1
                                if false {
                                    switch ("a" + ("hello" + ("foo bar" + ("foo bar" + "hello")))) {
                                    case "hello":
                                        switch "\n\t" {
                                        case "":
                                            switch genFn0((v3 + (^(v3))), 0, 42) {
                                            case 0:
                                                var v110 int8 = v13
                                                _ = v110
                                            }
                                        }
                                    case "":
                                        if len(v4) > 0 {
                                            v4[0] = ""
                                        }
                                    }
                                    var v111 uint8 = 136
                                    _ = v111
                                }
                                if (688 >= 42) {
                                    switch "hello" {
                                    case "foo bar":
                                        if ((59 == 128) || (true && true)) {
                                            if (113 == 1) {
                                                v4 = append(v4, ((("foo bar" + ("a" + "foo bar")) + "foo bar") + "hello"))
                                                v112 := GenS13{F0: 0, F1: 0}
                                                _ = v112
                                            }
                                            v3 -= v3
                                        }
                                    }
                                    v2 -= v2
                                }
                                _ = v100
                            }
                            if ((true || true) || true) {
                                v113 := GenS14{F0: 42, F1: false}
                                v2 -= v2
                                var v114 float32 = 0.1
                                _ = v113
                                _ = v114
                            } else {
                                v115 := GenS14{F0: 42, F1: (0 > -2147483648)}
                                v2 += v2
                                var v116 int32 = -2147483648
                                _ = v115
                                _ = v116
                            }
                        }
                    default:
                        fmt.Println(fi9)
                    }
                    _ = v13
                } else {
                    var v117 uint = 4294967295
                    v2 = (1.5 / (v2 * v2))
                    if false {
                        var v118 string = "a"
                        if true {
                            v117 += v117
                        }
                        var v119 uint64 = 2
                        _ = v118
                        _ = v119
                    }
                    if false {
                        v2 += v2
                        if (0 != 42) {
                            v117 += 184
                        }
                    }
                    _ = v117
                }
            }
            var v120 uint32 = 789
            var v121 int32 = -1
            switch (fi11 << ((255) & 63)) {
            case 1:
                v122 := GenS7{F0: v120, F1: 0}
                _ = v122
            }
            _ = fi11
            _ = v120
            _ = v121
        }
        _ = fi9
    }
    v3 -= v3
    fmt.Println(v2)
    var v123 uint8 = uint8(v3)
    switch ("foo bar" + "") {
    case "hello":
        v1[42] = "foo bar"
    case "":
        var v124 int8 = 127
        _ = v124
    }
    v4 = append(v4, "a")
    fmt.Println(v123, v123)
    var v125 uint32 = (^(v3))
    _ = len(v1)
    if true {
        v126 := GenS6{F0: 42, F1: 454}
        _ = v126
    }
    var v127 uint32 = 4294967295
    if false {
        v128 := GenS12{F0: true, F1: (1e10 < 0.1), F2: uint16(v123)}
        for fi129 := 0; fi129 < 14; fi129++ {
            if (((549 > 255) && false) && true) {
                v123 = (v123 >> ((42) & 63))
                v125 = (^(v127))
                v130 := GenS15{F0: ((fi129 % ((fi129) | 1)) != (fi129 ^ genFn0((v125 / (((v127 | (^(v125)))) | 1)), -9223372036854775808, genFn0(v3, 247, fi129)))), F1: 0, F2: (fi129 + 198)}
                _ = v130
            }
            var v131 uint32 = (^(v127))
            fmt.Println(v2, fi129)
            _ = fi129
            _ = v131
        }
        _ = v128
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v123:", v123)
    fmt.Println("v125:", v125)
    fmt.Println("v127:", v127)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v123
    _ = v125
    _ = v127
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
