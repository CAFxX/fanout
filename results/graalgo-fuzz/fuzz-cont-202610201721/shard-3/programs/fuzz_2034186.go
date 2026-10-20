package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint8; F1 uint8 }
type GenS3 struct { F0 int64; F1 string; F2 uint16 }
type GenS4 struct { F0 string; F1 uint; F2 uint16 }
type GenS5 struct { F0 uint16; F1 uint32 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 uint64; F1 uint64 }
type GenS8 struct { F0 uint8; F1 int16; F2 uint8 }
type GenS9 struct { F0 int8; F1 uint8; F2 bool }
type GenS10 struct { F0 int; F1 string }
type GenS11 struct { F0 int32 }
type GenS12 struct { F0 string; F1 int16; F2 int64 }
type GenS13 struct { F0 uint }
type GenS14 struct { F0 int32; F1 uint32 }
type GenS15 struct { F0 uint16; F1 int }

func main() {
    var v1 int64 = 42
    var v2 float32 = 1.5
    var v3 int = 32767
    var v4 int8 = 42
    switch "" {
    case "a":
        v4 += v4
    }
    var v5 float32 = 3.14159
    v1 -= (v1 - -1)
    v1 += (-(v1))
    if (!(2147483648 <= 2)) {
        if (!true) {
            v6 := GenS1{F0: 255}
            _ = v6
        } else {
            for fi7 := 0; fi7 < 5; fi7++ {
                if (42 > 0) {
                    switch (fi7 | fi7) {
                    case 0:
                        v2 -= float32(0.1)
                    case -352:
                        var v8 float32 = v5
                        _ = v8
                    }
                } else {
                    if (-9223372036854775808 > (fi7 % ((fi7) | 1))) {
                        if (660 != 0) {
                            v3 = v3
                            v3 = (v3 | 127)
                            if (42 != 65535) {
                                for fi9 := 0; fi9 < 14; fi9++ {
                                    v10 := []int{(fi9 - 128), (fi9 / (((fi7 / ((v3) | 1))) | 1)), (^(fi7))}
                                    switch (v3 / ((v3) | 1)) {
                                    case 32768:
                                        var v11 string = "\n\t"
                                        _ = v11
                                    default:
                                        if len(v10) >= 2 {
                                            v10 = v10[1:len(v10)]
                                        }
                                    }
                                    if (v1 != v1) {
                                        v3 = (fi9 % (((fi9 >> ((18446744073709551615) & 63))) | 1))
                                        switch "foo bar" {
                                        case "":
                                            _ = len(v10)
                                        case "foo bar":
                                            v10 = append(v10, 0)
                                        default:
                                            if true {
                                                var v12 string = ""
                                                v2 += (v5 / float32((1.5 / 1.5)))
                                                v13 := GenS1{F0: 255}
                                                _ = v12
                                                _ = v13
                                            }
                                        }
                                        v1 = v1
                                    } else {
                                        _ = len(v10)
                                        fmt.Println(v4)
                                        v14 := GenS2{F0: 198, F1: 2}
                                        _ = v14
                                    }
                                    _ = fi9
                                    _ = v10
                                }
                            }
                        } else {
                            var v15 int = (^(v3))
                            v16 := GenS2{F0: 2, F1: 68}
                            if (0 == 41) {
                                for fi17 := 0; fi17 < 13; fi17++ {
                                    var v18 float64 = 3.14159
                                    v19 := GenS1{F0: 1}
                                    if ((!(1000 < 0)) && false) {
                                        var v20 int8 = v4
                                        var v21 uint64 = 2
                                        var v22 string = ""
                                        if false {
                                            var v23 uint32 = 42
                                            v24 := GenS1{F0: 32767}
                                            _ = v23
                                            _ = v24
                                        } else {
                                            var v25 int16 = 0
                                            if false {
                                                var v26 int64 = v1
                                                v27 := GenS3{F0: (v26 ^ (v26 << ((255) & 63))), F1: ("\n\t" + v22), F2: 255}
                                                _ = v26
                                                _ = v27
                                            } else {
                                                v28 := []int{fi17, (^(v3)), 253}
                                                _ = v28
                                            }
                                            if true {
                                                v29 := []int{(fi17 ^ (v15 ^ v3)), 127, v3}
                                                fmt.Println(v22, v25, v1)
                                                _ = v29
                                            } else {
                                                v30 := GenS4{F0: v22, F1: 306, F2: 42}
                                                if (((!true) || true) || (2147483647 == 2)) {
                                                    var v31 uint8 = 2
                                                    v18 += 0.0
                                                    v32 := []int{v15, 9223372036854775807, (fi17 >> ((2) & 63)), fi7}
                                                    _ = v31
                                                    _ = v32
                                                } else {
                                                    var v33 uint64 = v21
                                                    _ = v33
                                                }
                                                _ = v30
                                            }
                                            v34 := GenS5{F0: 1, F1: 2147483648}
                                            _ = v25
                                            _ = v34
                                        }
                                        v15 -= (fi7 ^ fi17)
                                        if ((v21 / ((v21) | 1)) >= (v21 | (v21 >> ((42) & 63)))) {
                                            v20 = v4
                                            var v35 uint = 793
                                            switch (^(fi7)) {
                                            case 128:
                                                v35 += 184
                                            default:
                                                fmt.Println(v20)
                                            }
                                            _ = v35
                                        } else {
                                            var v36 int8 = v20
                                            v1 = v1
                                            var v37 uint = 0
                                            _ = v36
                                            _ = v37
                                        }
                                        fmt.Println(v3)
                                        _ = v20
                                        _ = v21
                                        _ = v22
                                    } else {
                                        switch "a" {
                                        case "a":
                                            var v38 uint = 2
                                            _ = v38
                                        case "\n\t":
                                            v39 := []int{}
                                            _ = v39
                                        case "":
                                            v40 := GenS6{F0: v4}
                                            _ = v40
                                        default:
                                            switch -660 {
                                            case 32768:
                                                switch fi7 {
                                                case -2147483648:
                                                    var v41 uint32 = 4294967295
                                                    _ = v41
                                                case 313:
                                                    switch ("hello" + "a") {
                                                    case "\n\t":
                                                        v1 += 9223372036854775807
                                                    case "":
                                                        fmt.Println(v3, v2)
                                                    }
                                                case 555:
                                                    switch "hello" {
                                                    case "hello":
                                                        v5 = (1.5 - v5)
                                                    case "foo bar":
                                                        if true {
                                                            var v42 float64 = (3.14159 + 1e10)
                                                            var v43 float32 = 1.5
                                                            v5 += (v2 * ((v5 + (0.1 + (v5 - ((float32(0.0) * 1.5) * (((v5 - 1.5) + (3.14159 / v2)) + v43))))) - v2))
                                                            _ = v42
                                                            _ = v43
                                                        } else {
                                                            fmt.Println(v3, v2, v3)
                                                        }
                                                    }
                                                default:
                                                    if false {
                                                        if true {
                                                            var v44 uint32 = 42
                                                            if false {
                                                                var v45 float32 = (float32(((v18 * (v18 - (float64((((v5 + 0.1) - v5) / float32(1.5))) - (-2.25 + 1.5)))) * ((v18 + float64(-2.25)) + v18))) / (v2 / v2))
                                                                if true {
                                                                    if ((v4 | (v4 | (v4 - -14))) <= 0) {
                                                                        if (fi7 != (v3 - 32767)) {
                                                                            switch "\n\t" {
                                                                            case "a":
                                                                                v46 := []int{(v15 - 42), (fi17 - -1), fi7, fi7}
                                                                                _ = v46
                                                                            case "\n\t":
                                                                                if true {
                                                                                    switch v15 {
                                                                                    case 4294967295:
                                                                                        var v47 uint32 = 2
                                                                                        _ = v47
                                                                                    }
                                                                                    v48 := []int{2147483647, v3, (fi17 * int(v4))}
                                                                                    fmt.Println(v45, fi17)
                                                                                    _ = v48
                                                                                }
                                                                            default:
                                                                                var v49 int8 = v4
                                                                                _ = v49
                                                                            }
                                                                            v50 := GenS7{F0: 1, F1: 0}
                                                                            _ = v50
                                                                        } else {
                                                                            v51 := GenS8{F0: 132, F1: 0, F2: 2}
                                                                            v1 = v1
                                                                            switch 588 {
                                                                            case -9223372036854775808:
                                                                                v52 := GenS7{F0: 42, F1: 480}
                                                                                _ = v52
                                                                            case 42:
                                                                                fmt.Println(v4, v15)
                                                                            case -2147483648:
                                                                                v53 := []int{}
                                                                                _ = v53
                                                                            }
                                                                            _ = v51
                                                                        }
                                                                    } else {
                                                                        var v54 string = ""
                                                                        v55 := GenS9{F0: (v4 & v4), F1: 245, F2: (!false)}
                                                                        v56 := GenS10{F0: 32768, F1: ((((("foo bar" + v54) + (v54 + v54)) + ((v54 + (((v54 + "hello") + "\n\t") + v54)) + v54)) + v54) + "a")}
                                                                        _ = v54
                                                                        _ = v55
                                                                        _ = v56
                                                                    }
                                                                    var v57 float32 = 1.5
                                                                    _ = v57
                                                                } else {
                                                                    fmt.Println(v45)
                                                                }
                                                                var v58 bool = false
                                                                _ = v45
                                                                _ = v58
                                                            }
                                                            v59 := []int{}
                                                            _ = v44
                                                            _ = v59
                                                        }
                                                        var v60 int = (^(fi7))
                                                        _ = v60
                                                    }
                                                }
                                            case 964:
                                                v61 := GenS11{F0: -281}
                                                _ = v61
                                            case -325:
                                                v62 := []int{v15, v15, fi7, 32768}
                                                _ = v62
                                            }
                                        }
                                        var v63 int = v15
                                        v2 += 0.1
                                        _ = v63
                                    }
                                    _ = fi17
                                    _ = v18
                                    _ = v19
                                }
                                fmt.Println(v1, v2, fi7)
                                var v64 uint = 212
                                _ = v64
                            } else {
                                fmt.Println(v5, fi7, v4)
                                if (557 <= 790) {
                                    switch "hello" {
                                    case "":
                                        v65 := GenS11{F0: 1}
                                        _ = v65
                                    case "a":
                                        v66 := []int{v3, v15}
                                        _ = v66
                                    default:
                                        v67 := GenS4{F0: "foo bar", F1: 4294967295, F2: 1}
                                        _ = v67
                                    }
                                } else {
                                    v1 += 352
                                    v4 += -15
                                }
                                for fi68 := 0; fi68 < 6; fi68++ {
                                    v69 := []int{-1, v3}
                                    _ = fi68
                                    _ = v69
                                }
                            }
                            _ = v15
                            _ = v16
                        }
                        if false {
                            if (100 < 128) {
                                for fi70 := 0; fi70 < 12; fi70++ {
                                    switch "\n\t" {
                                    case "":
                                        if (-1 == v1) {
                                            var v71 float64 = 3.14159
                                            var v72 float32 = 1e10
                                            _ = v71
                                            _ = v72
                                        } else {
                                            v73 := []int{fi70, (-(v3)), 0, (fi7 << ((255) & 63))}
                                            _ = v73
                                        }
                                    case "foo bar":
                                        switch (fi70 - 9223372036854775807) {
                                        case 42:
                                            v74 := GenS12{F0: ("\n\t" + "hello"), F1: -32768, F2: v1}
                                            _ = v74
                                        case 0:
                                            v75 := []int{}
                                            _ = v75
                                        case 9223372036854775807:
                                            var v76 int64 = v1
                                            _ = v76
                                        default:
                                            switch 128 {
                                            case -204:
                                                v77 := GenS13{F0: 553}
                                                _ = v77
                                            case 128:
                                                var v78 int8 = (-(v4))
                                                _ = v78
                                            default:
                                                switch (fi7 + (v3 ^ fi70)) {
                                                case -9223372036854775808:
                                                    v79 := []int{(^(fi70)), 101, v3, fi7}
                                                    _ = v79
                                                case 2147483647:
                                                    var v80 uint32 = 717
                                                    _ = v80
                                                case -2147483648:
                                                    if true {
                                                        fmt.Println(v4, v5)
                                                    }
                                                default:
                                                    v4 += v4
                                                }
                                            }
                                        }
                                    }
                                    var v81 uint16 = 42
                                    var v82 uint32 = 340
                                    _ = fi70
                                    _ = v81
                                    _ = v82
                                }
                                switch -483 {
                                case -9223372036854775808:
                                    v4 = v4
                                case 718:
                                    v83 := []int{v3, (fi7 - v3), 9223372036854775807}
                                    _ = v83
                                case 489:
                                    for fi84 := 0; fi84 < 14; fi84++ {
                                        var v85 int = v3
                                        if (42 >= 407) {
                                            if ((!true) || false) {
                                                var v86 uint8 = 1
                                                v86 = v86
                                                _ = v86
                                            }
                                            fmt.Println(v4)
                                        } else {
                                            var v87 bool = (42 < 84)
                                            var v88 int64 = (v1 ^ 505)
                                            v89 := []int{int(v4), 134, v85, -2147483648}
                                            _ = v87
                                            _ = v88
                                            _ = v89
                                        }
                                        _ = fi84
                                        _ = v85
                                    }
                                default:
                                    v1 += v1
                                }
                            }
                            v90 := []int{fi7, (v3 + fi7)}
                            _ = v90
                        } else {
                            v91 := []int{}
                            _ = v91
                        }
                    } else {
                        v4 += v4
                    }
                    v3 -= (fi7 * (^(v3)))
                    v92 := []int{fi7}
                    _ = v92
                }
                _ = fi7
            }
            for fi93 := 0; fi93 < 5; fi93++ {
                v3 = fi93
                _ = fi93
            }
            switch (-(v3)) {
            case 0:
                fmt.Println(v1)
            default:
                var v94 int8 = v4
                _ = v94
            }
        }
        v95 := GenS14{F0: 0, F1: 123}
        switch (v3 | 389) {
        case 0:
            fmt.Println(v2, v1, v5)
        case -637:
            v3 = v3
        }
        _ = v95
    }
    var v96 uint8 = 127
    v97 := []int{}
    var v98 int32 = 1
    v99 := GenS8{F0: (v96 & (v96 & v96)), F1: -32768, F2: (v96 ^ v96)}
    switch (("foo bar" + ((("\n\t" + "foo bar") + "\n\t") + "\n\t")) + "") {
    case "\n\t":
        var v100 int8 = v4
        _ = v100
    case "":
        switch v3 {
        case -713:
            v2 = (0.1 + 1e10)
        case -867:
            v101 := GenS15{F0: 15, F1: v3}
            _ = v101
        default:
            _ = len(v97)
        }
    case "a":
        var v102 int16 = 1000
        _ = v102
    }
    var v103 uint = 802
    var v104 string = ("" + (("a" + "hello") + (("hello" + "a") + "foo bar")))
    var v105 uint64 = 2
    var v106 uint64 = uint64(v96)
    v107 := GenS7{F0: (v105 / (((v106 | v105)) | 1)), F1: 18446744073709551615}
    fmt.Println("v1:", v1)
    fmt.Println("v103:", v103)
    fmt.Println("v104:", v104)
    fmt.Println("v105:", v105)
    fmt.Println("v106:", v106)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v96:", v96)
    fmt.Println("slicelen v97:", len(v97), v97)
    fmt.Println("v98:", v98)
    _ = v1
    _ = v103
    _ = v104
    _ = v105
    _ = v106
    _ = v107
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v96
    _ = v97
    _ = v98
    _ = v99
}
