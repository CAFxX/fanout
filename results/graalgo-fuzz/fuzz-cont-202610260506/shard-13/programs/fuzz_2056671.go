package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint16; F1 int8; F2 uint8 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 bool; F1 int16; F2 string }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 uint; F1 string }
type GenS7 struct { F0 int32; F1 uint64; F2 string }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 int32; F1 int32; F2 int }
type GenS10 struct { F0 uint8; F1 uint }
type GenS11 struct { F0 int8; F1 int16 }
type GenS12 struct { F0 uint64 }
type GenS13 struct { F0 uint64; F1 int64 }
type GenS14 struct { F0 uint; F1 int64; F2 uint8 }
type GenS15 struct { F0 int64; F1 string }
type GenS16 struct { F0 int; F1 int64; F2 int8 }
type GenS17 struct { F0 int32 }
type GenS18 struct { F0 int64; F1 uint64; F2 int8 }
type GenS19 struct { F0 uint; F1 uint; F2 int64 }
type GenS20 struct { F0 int64; F1 uint8; F2 uint16 }
type GenS21 struct { F0 uint16 }
type GenS22 struct { F0 uint8; F1 uint32; F2 uint64 }
type GenS23 struct { F0 int64; F1 bool }
type GenS24 struct { F0 uint8; F1 uint }
type GenS25 struct { F0 uint64 }
type GenS26 struct { F0 int32 }
type GenS27 struct { F0 int32; F1 uint32; F2 uint8 }
func genFn0() bool {
    switch 522 {
    case -2147483648:
        var v7 string = "foo bar"
        _ = v7
    }
    var v8 int8 = -128
    fmt.Println(v8)
    v8 += (-(v8))
    _ = v8
    return (("" + "a") == "foo bar")
}
func genFn1() uint64 {
    v9 := []int{1}
    var v10 string = "a"
    var v11 int = -253
    if (true || false) {
        v11 -= (v11 & v11)
    } else {
        v11 -= v11
    }
    switch (((v10 + v10) + "hello") + "a") {
    case "\n\t":
        switch ("foo bar" + v10) {
        case "hello":
            v11 = 32767
        case "\n\t":
            var v12 int64 = 0
            _ = v12
        case "":
            v10 = v10
        default:
            v13 := GenS1{F0: 585}
            _ = v13
        }
    case "":
        if (-34 < 80) {
            if len(v9) > 0 {
                v9[0] = 127
            }
        }
    }
    _ = v10
    _ = v11
    _ = v9
    return 1
}

func main() {
    var v1 uint64 = 57
    var v2 int64 = 9223372036854775807
    v3 := make(map[int]string)
    var v4 uint8 = 129
    var v5 uint = 18446744073709551615
    v6 := make(map[int]int)
    switch "foo bar" {
    case "\n\t":
        v5 += (v5 << (((v5 / ((0) | 1))) & 63))
    }
    var v14 int8 = -1
    v2 = 0
    var v15 int16 = 0
    v4 = v4
    fmt.Println(v2)
    v16 := GenS2{F0: 1, F1: v14, F2: v4}
    v14 = (v14 * -128)
    if (175 == 32768) {
        var v17 int = 525
        v15 -= v15
        if (64 <= 65535) {
            v14 -= v14
            v18 := GenS2{F0: 647, F1: v14, F2: (v4 | (v4 | (^(v4))))}
            _ = v18
        } else {
            v2 += (v2 * v2)
            fmt.Println(v15)
        }
        _ = v17
    } else {
        var v19 string = ""
        v20 := GenS2{F0: 1, F1: (^(v14)), F2: 128}
        v6[32767] = -752
        _ = v19
        _ = v20
    }
    var v21 uint8 = (^(v4))
    for fi22 := 0; fi22 < 18; fi22++ {
        v23 := []int{(fi22 | (fi22 % ((-9223372036854775808) | 1))), (^(fi22)), fi22, (fi22 / ((int(v4)) | 1))}
        for fi24 := 0; fi24 < 3; fi24++ {
            var v25 int8 = v14
            var v26 int64 = (v2 % ((0) | 1))
            v4 = (v4 + v21)
            switch ("" + (("hello" + "foo bar") + "hello")) {
            case "foo bar":
                switch "" {
                case "a":
                    var v27 float64 = ((-2.25 + ((1e10 / float64(-2.25)) * 1.5)) * (0.0 * 0.1))
                    _ = v27
                case "foo bar":
                    var v28 bool = false
                    var v29 uint16 = 255
                    _ = v28
                    _ = v29
                case "\n\t":
                    if false {
                        switch 0 {
                        case -2147483648:
                            delete(v3, fi24)
                        case -860:
                            if (v1 != genFn1()) {
                                if (!(140 != 0)) {
                                    var v30 uint64 = 42
                                    _ = v30
                                }
                            } else {
                                v31 := v3[fi24]
                                var v32 int32 = 1
                                _ = v31
                                _ = v32
                            }
                        case 127:
                            if false {
                                fmt.Println(v2)
                                v33 := GenS2{F0: 32767, F1: (v25 & 98), F2: 127}
                                if len(v23) >= 2 {
                                    v23 = v23[1:len(v23)]
                                }
                                _ = v33
                            } else {
                                v26 = 9223372036854775807
                                switch fi22 {
                                case -9223372036854775808:
                                    v26 = v26
                                case -1:
                                    if false {
                                        var v34 int = -9223372036854775808
                                        if ((v15 / ((v15) | 1)) >= (v15 * v15)) {
                                            v21 += v4
                                            fmt.Println(v25, v25, fi22)
                                            switch v34 {
                                            case 336:
                                                v35 := GenS1{F0: (v26 << (((v5 % (((^(v5))) | 1))) & 63))}
                                                _ = v35
                                            case 2147483647:
                                                fmt.Println(v5, v1, v21)
                                                if (868 == 1) {
                                                    v5 = (v5 >> ((0) & 63))
                                                    if false {
                                                        var v36 float32 = 0.1
                                                        v37 := GenS3{F0: (v15 ^ (v15 + 32767))}
                                                        if true {
                                                            var v38 int = (fi24 % ((fi24) | 1))
                                                            var v39 int = 127
                                                            if (!(!true)) {
                                                                if false {
                                                                    if (false || false) {
                                                                        if true {
                                                                            v36 -= 1.5
                                                                            if (!(((v14 | (v25 << ((v5) & 63))) <= (v25 << ((0) & 63))) && ("a" >= ""))) {
                                                                                switch "foo bar" {
                                                                                case "\n\t":
                                                                                    v1 += (v1 | (v1 / (((v1 & (v1 - v1))) | 1)))
                                                                                    var v40 bool = true
                                                                                    var v41 uint = v5
                                                                                    _ = v40
                                                                                    _ = v41
                                                                                case "a":
                                                                                    if true {
                                                                                        v1 = 578
                                                                                        v5 -= (v5 + (v5 / ((uint(v4)) | 1)))
                                                                                    }
                                                                                case "hello":
                                                                                    v42 := GenS3{F0: (v15 % ((v15) | 1))}
                                                                                    _ = v42
                                                                                default:
                                                                                    v15 -= v15
                                                                                }
                                                                                if (v26 < v26) {
                                                                                    var v43 uint32 = 65535
                                                                                    var v44 int32 = 42
                                                                                    _ = v43
                                                                                    _ = v44
                                                                                }
                                                                                var v45 bool = false
                                                                                _ = v45
                                                                            } else {
                                                                                if (!(uint64(v5) < v1)) {
                                                                                    fmt.Println(v36)
                                                                                }
                                                                            }
                                                                        } else {
                                                                            fmt.Println(fi22, v34, v26)
                                                                            v25 = 100
                                                                        }
                                                                        v46 := GenS3{F0: v15}
                                                                        if (-640 < 432) {
                                                                            v3[(v39 >> (((v5 - 1)) & 63))] = ("" + ("a" + (("\n\t" + "\n\t") + "foo bar")))
                                                                            if (-2.25 != 1.5) {
                                                                                v1 = (^(v1))
                                                                            } else {
                                                                                v47 := GenS1{F0: (v2 & -9223372036854775808)}
                                                                                v48 := GenS1{F0: v26}
                                                                                _ = v47
                                                                                _ = v48
                                                                            }
                                                                            fmt.Println(v34, v21)
                                                                        } else {
                                                                            v25 += v25
                                                                            v38 = v39
                                                                            fmt.Println(v34, v34)
                                                                        }
                                                                        _ = v46
                                                                    } else {
                                                                        v4 -= v21
                                                                        if len(v23) >= 2 {
                                                                            v23 = v23[1:len(v23)]
                                                                        }
                                                                        v26 = (v26 ^ v2)
                                                                        _ = len(v3)
                                                                    }
                                                                    v49 := GenS2{F0: 2, F1: (v14 / (((v25 >> ((42) & 63))) | 1)), F2: (v4 >> (((^(v5))) & 63))}
                                                                    _ = v49
                                                                } else {
                                                                    if (!(96 == (fi22 * (fi24 | (v39 * fi24))))) {
                                                                        var v50 int32 = 2147483647
                                                                        _ = v50
                                                                    } else {
                                                                        switch ("" + "a") {
                                                                        case "\n\t":
                                                                            if false {
                                                                                fmt.Println(v2, v39)
                                                                            } else {
                                                                                if true {
                                                                                    v23 = append(v23, v39)
                                                                                    var v51 uint8 = v4
                                                                                    var v52 bool = false
                                                                                    _ = v51
                                                                                    _ = v52
                                                                                } else {
                                                                                    var v53 bool = true
                                                                                    v54 := GenS3{F0: v15}
                                                                                    switch fi24 {
                                                                                    case -483:
                                                                                        switch ("\n\t" + "") {
                                                                                        case "a":
                                                                                            _ = len(v23)
                                                                                        }
                                                                                    default:
                                                                                        fmt.Println(v1, v36)
                                                                                        _ = len(v3)
                                                                                        var v55 uint = v5
                                                                                        _ = v55
                                                                                    }
                                                                                    _ = v53
                                                                                    _ = v54
                                                                                }
                                                                                fmt.Println(v25, v5)
                                                                            }
                                                                            if false {
                                                                                var v56 uint16 = 255
                                                                                _ = v56
                                                                            }
                                                                            v57 := GenS4{F0: (!(!(true && false))), F1: -801, F2: (("" + ("a" + "hello")) + ("foo bar" + ((("\n\t" + "a") + "hello") + "hello")))}
                                                                            _ = v57
                                                                        case "hello":
                                                                            if (v1 >= 0) {
                                                                                fmt.Println(fi22, v4)
                                                                                if (false || ((v26 / ((v2) | 1)) <= v2)) {
                                                                                    if (v4 < (^(v4))) {
                                                                                        var v58 int32 = 42
                                                                                        _ = v58
                                                                                    } else {
                                                                                        var v59 uint64 = 4294967295
                                                                                        v21 += v21
                                                                                        switch 32768 {
                                                                                        case 366:
                                                                                            if (-1 != -1) {
                                                                                                switch (("hello" + "foo bar") + ("" + "foo bar")) {
                                                                                                case "\n\t":
                                                                                                    v60 := GenS5{F0: (v5 + 65535)}
                                                                                                    _ = v60
                                                                                                default:
                                                                                                    _ = len(v3)
                                                                                                }
                                                                                            }
                                                                                        default:
                                                                                            if ((((0.1 / v36) * v36) == ((v36 * v36) / float32(-2.25))) && (false || true)) {
                                                                                                if true {
                                                                                                    fmt.Println(v34, v21)
                                                                                                }
                                                                                                switch "hello" {
                                                                                                case "a":
                                                                                                    fmt.Println(v2, v2, v39)
                                                                                                }
                                                                                                switch ("a" + (((((("a" + "") + "hello") + (("a" + ("foo bar" + "foo bar")) + "\n\t")) + "\n\t") + (("a" + "hello") + "foo bar")) + "")) {
                                                                                                case "hello":
                                                                                                    var v61 int32 = -2147483648
                                                                                                    _ = v61
                                                                                                case "foo bar":
                                                                                                    v26 = (v26 & v2)
                                                                                                case "":
                                                                                                    fmt.Println(v59, v36, v2)
                                                                                                }
                                                                                            } else {
                                                                                                if false {
                                                                                                    fmt.Println(v2, v26)
                                                                                                    if true {
                                                                                                        var v62 int64 = v2
                                                                                                        v63 := GenS4{F0: true, F1: (v15 | v15), F2: "a"}
                                                                                                        v64 := GenS6{F0: 18446744073709551615, F1: ""}
                                                                                                        _ = v62
                                                                                                        _ = v63
                                                                                                        _ = v64
                                                                                                    }
                                                                                                    v36 = v36
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                        _ = v59
                                                                                    }
                                                                                    v65 := GenS7{F0: 409, F1: 42, F2: "hello"}
                                                                                    _ = v65
                                                                                }
                                                                            }
                                                                        }
                                                                        var v66 uint16 = 42
                                                                        var v67 uint8 = (v21 & (v4 >> ((592) & 63)))
                                                                        _ = v66
                                                                        _ = v67
                                                                    }
                                                                    v21 -= 255
                                                                    v25 = v25
                                                                }
                                                                v68 := v6[v34]
                                                                _ = v68
                                                            } else {
                                                                switch (("foo bar" + ("a" + "")) + "\n\t") {
                                                                case "hello":
                                                                    if len(v23) >= 2 {
                                                                        v23 = v23[1:len(v23)]
                                                                    }
                                                                default:
                                                                    v38 = v39
                                                                }
                                                                var v69 int64 = v26
                                                                if len(v23) >= 2 {
                                                                    v23 = v23[1:len(v23)]
                                                                }
                                                                _ = v69
                                                            }
                                                            v4 = v21
                                                            switch "foo bar" {
                                                            case "hello":
                                                                var v70 string = ("a" + ("a" + "a"))
                                                                _ = v70
                                                            }
                                                            switch ("foo bar" + "\n\t") {
                                                            case "foo bar":
                                                                var v71 uint8 = 46
                                                                _ = v71
                                                            default:
                                                                var v72 uint16 = uint16(v21)
                                                                _ = v72
                                                            }
                                                            if false {
                                                                if len(v23) >= 2 {
                                                                    v23 = v23[1:len(v23)]
                                                                }
                                                                v23 = append(v23, (v38 << ((948) & 63)))
                                                                var v73 bool = (!true)
                                                                _ = v73
                                                            }
                                                            _ = v38
                                                            _ = v39
                                                        } else {
                                                            if len(v23) >= 2 {
                                                                v23 = v23[1:len(v23)]
                                                            }
                                                            if false {
                                                                v74 := GenS7{F0: 0, F1: v1, F2: "hello"}
                                                                if false {
                                                                    fmt.Println(v26, v21, v4)
                                                                    v34 = 42
                                                                    var v75 int32 = -55
                                                                    var v76 float64 = -2.25
                                                                    _ = v75
                                                                    _ = v76
                                                                } else {
                                                                    var v77 int16 = (v15 << ((2) & 63))
                                                                    _ = v77
                                                                }
                                                                _ = v74
                                                            }
                                                            v21 += (v4 - 127)
                                                            fmt.Println(v34, v26)
                                                        }
                                                        v14 = int8(v2)
                                                        _ = v36
                                                        _ = v37
                                                    } else {
                                                        fmt.Println(v2, fi24, v15)
                                                    }
                                                    v21 += (v4 >> (((v5 & uint(v1))) & 63))
                                                } else {
                                                    v78 := GenS7{F0: -1, F1: genFn1(), F2: "foo bar"}
                                                    v79 := v6[(v34 + (fi22 / ((2147483647) | 1)))]
                                                    fmt.Println(v21)
                                                    v14 -= (v25 << ((18446744073709551615) & 63))
                                                    _ = v78
                                                    _ = v79
                                                }
                                                var v80 int = 32768
                                                if (v1 != (v1 * (v1 + uint64(v21)))) {
                                                    var v81 uint8 = v4
                                                    var v82 uint8 = 42
                                                    var v83 float32 = float32((3.14159 - 3.14159))
                                                    _ = v81
                                                    _ = v82
                                                    _ = v83
                                                } else {
                                                    v84 := GenS1{F0: v2}
                                                    switch v34 {
                                                    case 1:
                                                        v26 = v2
                                                        var v85 uint16 = 0
                                                        _ = v85
                                                    case 589:
                                                        v86 := GenS5{F0: 690}
                                                        fmt.Println(v25, v21)
                                                        fmt.Println(fi24, v80)
                                                        switch "" {
                                                        case "\n\t":
                                                            if true {
                                                                if false {
                                                                    if (1 == 42) {
                                                                        var v87 int = v80
                                                                        fmt.Println(v2, v87, v15)
                                                                        _ = v87
                                                                    }
                                                                    v21 = (v4 + 110)
                                                                    v88 := GenS4{F0: (0 != 708), F1: v15, F2: (("hello" + "a") + "\n\t")}
                                                                    _ = v88
                                                                }
                                                                fmt.Println(v25, v2)
                                                                var v89 bool = (false || false)
                                                                _ = v89
                                                            } else {
                                                                var v90 uint = (v5 % ((uint(v1)) | 1))
                                                                switch v80 {
                                                                case -1:
                                                                    if ((v21 + v21) >= 15) {
                                                                        switch fi24 {
                                                                        case -1:
                                                                            v91 := GenS8{F0: 1}
                                                                            _ = v91
                                                                        }
                                                                        var v92 uint32 = 4294967295
                                                                        if (127 == fi22) {
                                                                            if (!true) {
                                                                                if true {
                                                                                    if (!(!(true && (int(v4) <= v80)))) {
                                                                                        var v93 uint = (v5 >> ((v5) & 63))
                                                                                        _ = v93
                                                                                    }
                                                                                    switch "hello" {
                                                                                    case "foo bar":
                                                                                        v21 = (v4 << (((v90 / ((uint(v1)) | 1))) & 63))
                                                                                    case "hello":
                                                                                        if false {
                                                                                            v94 := GenS9{F0: 0, F1: -934, F2: v80}
                                                                                            v26 -= -49
                                                                                            v95 := GenS9{F0: 2147483647, F1: -1, F2: (^(v34))}
                                                                                            fmt.Println(v34, v2)
                                                                                            fmt.Println(v26)
                                                                                            _ = v94
                                                                                            _ = v95
                                                                                        } else {
                                                                                            switch 128 {
                                                                                            case 1:
                                                                                                var v96 uint16 = 1
                                                                                                v14 += (v14 * (v25 ^ v25))
                                                                                                _ = v96
                                                                                            case -2147483648:
                                                                                                var v97 float64 = 0.1
                                                                                                _ = v97
                                                                                            case 2147483647:
                                                                                                var v98 int16 = v15
                                                                                                _ = v98
                                                                                            }
                                                                                        }
                                                                                    default:
                                                                                        if false {
                                                                                            v14 += v25
                                                                                            var v99 uint32 = 83
                                                                                            fmt.Println(v5, v90, v15)
                                                                                            fmt.Println(v80, v5)
                                                                                            var v100 string = ("\n\t" + "")
                                                                                            _ = v100
                                                                                            _ = v99
                                                                                        } else {
                                                                                            var v101 int = (fi24 | (^(v80)))
                                                                                            _ = v101
                                                                                        }
                                                                                    }
                                                                                } else {
                                                                                    switch "hello" {
                                                                                    case "foo bar":
                                                                                        v15 += (^(v15))
                                                                                        v25 = int8(v15)
                                                                                    default:
                                                                                        fmt.Println(v4, fi22, v34)
                                                                                    }
                                                                                }
                                                                            } else {
                                                                                if ((1 != 42) && false) {
                                                                                    if true {
                                                                                        fmt.Println(v92, fi24, fi22)
                                                                                        var v102 uint16 = 42
                                                                                        switch "foo bar" {
                                                                                        case "hello":
                                                                                            delete(v3, fi22)
                                                                                        case "":
                                                                                            var v103 int64 = -1
                                                                                            _ = v103
                                                                                        default:
                                                                                            if true {
                                                                                                v104 := GenS3{F0: v15}
                                                                                                v105 := GenS7{F0: -2147483648, F1: genFn1(), F2: ("a" + "")}
                                                                                                _ = v104
                                                                                                _ = v105
                                                                                            } else {
                                                                                                fmt.Println(v21, v26)
                                                                                                v106 := GenS10{F0: v21, F1: 255}
                                                                                                switch "\n\t" {
                                                                                                case "a":
                                                                                                    switch -9223372036854775808 {
                                                                                                    case 1:
                                                                                                        delete(v6, int(v25))
                                                                                                    case -2147483648:
                                                                                                        v107 := GenS11{F0: v14, F1: (v15 << (((v90 + 255)) & 63))}
                                                                                                        _ = v107
                                                                                                    default:
                                                                                                        v34 -= fi24
                                                                                                    }
                                                                                                default:
                                                                                                    v4 -= v4
                                                                                                }
                                                                                                _ = v106
                                                                                            }
                                                                                        }
                                                                                        _ = v102
                                                                                    }
                                                                                    var v108 uint = uint(v92)
                                                                                    var v109 int32 = 53
                                                                                    _ = v108
                                                                                    _ = v109
                                                                                } else {
                                                                                    switch 763 {
                                                                                    case -22:
                                                                                        v110 := GenS6{F0: (v90 ^ (v5 % ((v90) | 1))), F1: ("" + "foo bar")}
                                                                                        _ = v110
                                                                                    }
                                                                                }
                                                                                var v111 string = "hello"
                                                                                _ = len(v3)
                                                                                v23 = append(v23, -9223372036854775808)
                                                                                _ = v111
                                                                            }
                                                                            v1 = (v1 - (v1 - (^(v1))))
                                                                            if (2 < 255) {
                                                                                var v112 float32 = 1e10
                                                                                _ = v112
                                                                            }
                                                                        } else {
                                                                            v21 -= v21
                                                                        }
                                                                        _ = v92
                                                                    } else {
                                                                        v5 = v5
                                                                    }
                                                                case 128:
                                                                    var v113 uint64 = 782
                                                                    _ = v113
                                                                case 0:
                                                                    var v114 uint8 = 139
                                                                    _ = v114
                                                                }
                                                                v115 := GenS12{F0: (^(v1))}
                                                                _ = v115
                                                                _ = v90
                                                            }
                                                        case "hello":
                                                            if len(v23) > 0 {
                                                                v23[0] = (v34 >> ((v5) & 63))
                                                            }
                                                        default:
                                                            var v116 int8 = (v25 & (v14 << (((v5 | (v5 | v5))) & 63)))
                                                            _ = v116
                                                        }
                                                        var v117 uint = v5
                                                        v34 += v80
                                                        switch (("" + "") + (("hello" + "a") + "\n\t")) {
                                                        case "foo bar":
                                                            v118 := GenS13{F0: (v1 >> (((v5 << ((v117) & 63))) & 63)), F1: v26}
                                                            _ = v118
                                                        default:
                                                            if (!((v25 << (((^(v5))) & 63)) >= v14)) {
                                                                switch 0 {
                                                                case 9223372036854775807:
                                                                    v2 = (v2 << ((v5) & 63))
                                                                case 2147483647:
                                                                    switch "a" {
                                                                    case "a":
                                                                        var v119 int64 = v2
                                                                        v14 = int8(v119)
                                                                        var v120 uint64 = (v1 % ((genFn1()) | 1))
                                                                        _ = v119
                                                                        _ = v120
                                                                    }
                                                                case -787:
                                                                    switch "hello" {
                                                                    case "foo bar":
                                                                        var v121 int64 = (v26 & (v26 << ((18446744073709551615) & 63)))
                                                                        _ = v121
                                                                    case "a":
                                                                        _ = len(v23)
                                                                    case "":
                                                                        _ = len(v23)
                                                                    }
                                                                default:
                                                                    v117 -= (^(v5))
                                                                }
                                                                if false {
                                                                    fmt.Println(v15, v34, v25)
                                                                    var v122 string = "hello"
                                                                    if false {
                                                                        var v123 uint64 = (v1 ^ v1)
                                                                        if ((v5 * v5) <= v5) {
                                                                            fmt.Println(v80, v123, v2)
                                                                            fmt.Println(v14, v34, v80)
                                                                            switch v122 {
                                                                            case "foo bar":
                                                                                v124 := GenS6{F0: v117, F1: (v122 + (v122 + v122))}
                                                                                _ = v124
                                                                            case "\n\t":
                                                                                fmt.Println(v122, v80)
                                                                                v125 := GenS5{F0: (v117 / ((65535) | 1))}
                                                                                fmt.Println(v122, v34)
                                                                                _ = v125
                                                                            default:
                                                                                if false {
                                                                                    if (((v122 + (v122 + (v122 + v122))) + (v122 + v122)) == v122) {
                                                                                        if true {
                                                                                            var v126 int64 = 0
                                                                                            _ = v126
                                                                                        }
                                                                                    }
                                                                                    if false {
                                                                                        _ = len(v23)
                                                                                        fmt.Println(v21)
                                                                                    } else {
                                                                                        v2 += v2
                                                                                        v127 := v3[4294967295]
                                                                                        v15 = v15
                                                                                        _ = v127
                                                                                    }
                                                                                } else {
                                                                                    fmt.Println(v25, v5)
                                                                                    if len(v23) >= 2 {
                                                                                        v23 = v23[1:len(v23)]
                                                                                    }
                                                                                }
                                                                            }
                                                                        }
                                                                        var v128 int = 32767
                                                                        _ = v123
                                                                        _ = v128
                                                                    } else {
                                                                        if true {
                                                                            v23 = append(v23, 4294967295)
                                                                            _ = len(v3)
                                                                        }
                                                                    }
                                                                    _ = v122
                                                                } else {
                                                                    var v129 int64 = v26
                                                                    _ = v129
                                                                }
                                                                var v130 string = "a"
                                                                _ = v130
                                                            } else {
                                                                v4 = v4
                                                                v2 = (v2 * v26)
                                                            }
                                                        }
                                                        _ = v117
                                                        _ = v86
                                                    case -1:
                                                        fmt.Println(v34, v4)
                                                    default:
                                                        if (false && true) {
                                                            var v131 uint64 = (v1 + uint64(v4))
                                                            _ = v131
                                                        }
                                                    }
                                                    if false {
                                                        v80 -= (v80 << ((uint(v1)) & 63))
                                                        var v132 string = "\n\t"
                                                        v23 = append(v23, (fi22 >> ((945) & 63)))
                                                        _ = v132
                                                    } else {
                                                        fmt.Println(v2, v2, v25)
                                                        switch ("a" + "") {
                                                        case "hello":
                                                            v133 := GenS14{F0: (v5 / ((1) | 1)), F1: 0, F2: 128}
                                                            _ = v133
                                                        case "":
                                                            var v134 float32 = 0.0
                                                            _ = v134
                                                        default:
                                                            v23 = append(v23, fi24)
                                                        }
                                                        fmt.Println(v26, v1, fi24)
                                                    }
                                                    _ = v84
                                                }
                                                _ = v80
                                            }
                                        } else {
                                            if true {
                                                v135 := GenS15{F0: -9223372036854775808, F1: "\n\t"}
                                                var v136 int32 = 126
                                                var v137 float64 = (1.5 - 0.0)
                                                _ = v135
                                                _ = v136
                                                _ = v137
                                            } else {
                                                if true {
                                                    switch -883 {
                                                    case -2147483648:
                                                        var v138 int16 = 1000
                                                        _ = v138
                                                    default:
                                                        v14 += 0
                                                    }
                                                    fmt.Println(v4, v34, v4)
                                                    var v139 int32 = 0
                                                    _ = v139
                                                } else {
                                                    if ((v14 <= (-(v25))) && ((!true) || false)) {
                                                        v140 := GenS16{F0: fi24, F1: (v2 | -9223372036854775808), F2: int8(v1)}
                                                        _ = v140
                                                    }
                                                }
                                            }
                                            v25 += v25
                                            if len(v23) >= 2 {
                                                v23 = v23[1:len(v23)]
                                            }
                                        }
                                        v141 := v3[fi24]
                                        _ = v141
                                        _ = v34
                                    } else {
                                        v142 := GenS7{F0: -616, F1: (v1 & (v1 * (v1 - genFn1()))), F2: "foo bar"}
                                        if (true && false) {
                                            v26 -= (v26 << ((uint(v21)) & 63))
                                        } else {
                                            fmt.Println(v1)
                                            if true {
                                                if len(v23) > 0 {
                                                    v23[0] = (fi22 | (fi22 | (fi24 + (fi22 ^ fi24))))
                                                }
                                                var v143 int64 = (v2 | v2)
                                                var v144 int32 = 0
                                                if (int32(v14) != v144) {
                                                    v1 = 491
                                                } else {
                                                    var v145 uint16 = uint16(v5)
                                                    var v146 uint32 = 2
                                                    _ = v145
                                                    _ = v146
                                                }
                                                _ = v143
                                                _ = v144
                                            }
                                            v147 := GenS12{F0: v1}
                                            _ = v147
                                        }
                                        if false {
                                            var v148 uint32 = 785
                                            v1 += (^(v1))
                                            _ = v148
                                        }
                                        _ = v142
                                    }
                                case 42:
                                    v23 = append(v23, (fi24 << ((630) & 63)))
                                default:
                                    v149 := GenS17{F0: 1}
                                    _ = v149
                                }
                                var v150 string = "foo bar"
                                _ = v150
                            }
                        }
                        var v151 int32 = -1
                        _ = v151
                    } else {
                        var v152 uint16 = 42
                        _ = v152
                    }
                }
            }
            _ = fi24
            _ = v25
            _ = v26
        }
        _ = fi22
        _ = v23
    }
    fmt.Println(v4)
    for fi153 := 0; fi153 < 20; fi153++ {
        if (!(false || true)) {
            fmt.Println(v5)
        }
        _ = fi153
    }
    for fi154 := 0; fi154 < 16; fi154++ {
        for fi155 := 0; fi155 < 3; fi155++ {
            var v156 int = 127
            v1 = v1
            var v157 uint32 = 0
            _ = fi155
            _ = v156
            _ = v157
        }
        for fi158 := 0; fi158 < 1; fi158++ {
            if true {
                var v159 int16 = (^(v15))
                if false {
                    v160 := GenS18{F0: v2, F1: 70, F2: v14}
                    var v161 int32 = -2147483648
                    v1 -= 18446744073709551615
                    var v162 string = "foo bar"
                    v163 := GenS4{F0: ((1e10 == (-2.25 - 3.14159)) || false), F1: (v15 ^ (v15 * int16(v1))), F2: (v162 + (("foo bar" + v162) + ((v162 + ((((("foo bar" + (v162 + v162)) + (v162 + (v162 + ("" + "\n\t")))) + v162) + "foo bar") + "hello")) + "")))}
                    var v164 int64 = (v2 + 0)
                    if true {
                        v165 := GenS4{F0: false, F1: (v15 | v159), F2: (v162 + v162)}
                        _ = v165
                    } else {
                        v166 := v6[167]
                        v164 += v2
                        v167 := []int{(v166 / ((fi158) | 1)), (fi158 | 128), fi154}
                        v168 := GenS14{F0: (v5 & v5), F1: (v2 & -549), F2: v4}
                        var v169 uint8 = 12
                        var v170 int64 = v164
                        if true {
                            v171 := GenS19{F0: v5, F1: v5, F2: (v164 | v164)}
                            if (0.0 < (float32(-2.25) + ((3.14159 + ((-2.25 / 0.1) / 1.0)) + 3.14159))) {
                                var v172 float64 = ((1.5 / float64(1.5)) * -2.25)
                                if true {
                                    var v173 uint16 = 486
                                    var v174 int64 = v164
                                    if (100 >= (v14 | (v14 - (v14 + (v14 / ((int8(v161)) | 1)))))) {
                                        v173 = v173
                                        var v175 int64 = 737
                                        var v176 uint16 = v173
                                        _ = v175
                                        _ = v176
                                    }
                                    _ = v173
                                    _ = v174
                                } else {
                                    delete(v3, fi158)
                                    v15 = v15
                                    if true {
                                        switch "hello" {
                                        case "foo bar":
                                            _ = len(v167)
                                        case "\n\t":
                                            var v177 float32 = -2.25
                                            _ = v177
                                        case "":
                                            delete(v6, 2147483647)
                                        default:
                                            if len(v167) > 0 {
                                                v167[0] = fi158
                                            }
                                            v4 = (v21 << ((838) & 63))
                                            var v178 int = (v166 / ((fi158) | 1))
                                            _ = v178
                                        }
                                    } else {
                                        var v179 int = 4
                                        _ = v179
                                    }
                                    switch fi154 {
                                    case 0:
                                        var v180 int8 = v14
                                        _ = v180
                                    }
                                    var v181 string = v162
                                    v14 = (-(v14))
                                    _ = v181
                                }
                                v2 -= (v164 % ((v170) | 1))
                                _ = v172
                            } else {
                                v182 := GenS20{F0: v170, F1: v169, F2: 322}
                                _ = v182
                            }
                            var v183 uint = v5
                            _ = v171
                            _ = v183
                        }
                        _ = v166
                        _ = v167
                        _ = v168
                        _ = v169
                        _ = v170
                    }
                    v184 := GenS21{F0: 42}
                    fmt.Println(v5, v159, v15)
                    _ = v160
                    _ = v161
                    _ = v162
                    _ = v163
                    _ = v164
                    _ = v184
                } else {
                    _ = len(v6)
                    switch "" {
                    case "foo bar":
                        switch (fi154 * (-(fi154))) {
                        case -280:
                            var v185 uint16 = 876
                            _ = v185
                        case 0:
                            v186 := GenS22{F0: (v4 >> (((^(v5))) & 63)), F1: 63, F2: v1}
                            _ = v186
                        default:
                            v187 := GenS13{F0: (^(v1)), F1: int64(v14)}
                            _ = v187
                        }
                    }
                    fmt.Println(v1, v2, fi158)
                }
                v188 := GenS7{F0: -2147483648, F1: genFn1(), F2: "foo bar"}
                _ = v159
                _ = v188
            } else {
                v4 += 255
                switch "hello" {
                case "hello":
                    v189 := []int{fi158}
                    _ = v189
                case "":
                    switch fi158 {
                    case 537:
                        fmt.Println(fi158, v4)
                        fmt.Println(v15)
                        v190 := []int{(^(fi158)), fi154}
                        switch "a" {
                        case "foo bar":
                            var v191 float32 = 1e10
                            _ = v191
                        case "hello":
                            if ((0.1 / ((0.0 + 1.5) * 1e10)) >= ((-2.25 + (1.5 * ((-2.25 + 0.1) - ((0.0 / float64(((float32(0.1) + float32(0.0)) * 1e10))) / (float64(1.5) + float64((float32(1.5) * -2.25))))))) + 0.1)) {
                                switch "" {
                                case "":
                                    v4 = (^(v4))
                                case "\n\t":
                                    v192 := GenS18{F0: v2, F1: 42, F2: -53}
                                    _ = v192
                                }
                            }
                        }
                        if (!false) {
                            if false {
                                v1 = genFn1()
                            } else {
                                var v193 uint16 = 255
                                v194 := GenS9{F0: 0, F1: 0, F2: fi158}
                                if ((!(true || true)) && false) {
                                    fmt.Println(v4, v2)
                                    if ((1e10 < 1e10) && true) {
                                        if (((("hello" + "foo bar") + ("" + "a")) + "foo bar") >= "a") {
                                            _ = len(v190)
                                            switch fi158 {
                                            case 32767:
                                                v195 := GenS23{F0: (v2 * int64(fi154)), F1: true}
                                                _ = v195
                                            default:
                                                v14 += -128
                                            }
                                        } else {
                                            switch "hello" {
                                            case "a":
                                                var v196 uint = v5
                                                _ = v196
                                            default:
                                                var v197 float64 = 0.1
                                                _ = v197
                                            }
                                        }
                                        if true {
                                            var v198 uint64 = v1
                                            var v199 int16 = (v15 << ((42) & 63))
                                            _ = v198
                                            _ = v199
                                        }
                                        var v200 int16 = v15
                                        _ = v200
                                    }
                                } else {
                                    if (!true) {
                                        switch "foo bar" {
                                        case "foo bar":
                                            switch "foo bar" {
                                            case "":
                                                v6[fi154] = fi154
                                            case "a":
                                                v1 -= 0
                                            default:
                                                if true {
                                                    v14 = 5
                                                }
                                            }
                                        case "":
                                            if (v1 <= (v1 * (v1 | (v1 * v1)))) {
                                                if (true && (!false)) {
                                                    if (true && (!(v15 <= int16(v14)))) {
                                                        var v201 uint64 = (v1 >> ((v5) & 63))
                                                        switch (fi158 | (fi158 & (fi158 + (fi158 - (fi154 | (-(fi158))))))) {
                                                        case 9223372036854775807:
                                                            var v202 int8 = -1
                                                            _ = v202
                                                        case 4294967295:
                                                            _ = len(v190)
                                                        case 2147483647:
                                                            delete(v3, 2147483647)
                                                        }
                                                        _ = v201
                                                    } else {
                                                        var v203 bool = (("\n\t" + ("a" + "hello")) >= "hello")
                                                        _ = v203
                                                    }
                                                    if false {
                                                        v14 = v14
                                                        var v204 bool = true
                                                        delete(v6, fi158)
                                                        if len(v190) >= 2 {
                                                            v190 = v190[1:len(v190)]
                                                        }
                                                        v205 := GenS14{F0: v5, F1: (-(v2)), F2: 42}
                                                        if true {
                                                            if ((v14 << (((v5 | uint(v193))) & 63)) >= int8(v5)) {
                                                                v5 -= 42
                                                            } else {
                                                                v206 := GenS24{F0: v21, F1: v5}
                                                                _ = v206
                                                            }
                                                        } else {
                                                            if len(v190) > 0 {
                                                                v190[0] = -204
                                                            }
                                                        }
                                                        v193 += 255
                                                        v204 = false
                                                        v190 = append(v190, -2147483648)
                                                        v21 = (v4 / (((v21 ^ v21)) | 1))
                                                        _ = v204
                                                        _ = v205
                                                    } else {
                                                        if false {
                                                            v15 += 590
                                                        }
                                                        _ = len(v3)
                                                        _ = len(v6)
                                                        v207 := v3[451]
                                                        _ = v207
                                                    }
                                                }
                                                var v208 float64 = 1.5
                                                _ = v208
                                            }
                                        }
                                        v14 = -101
                                        if len(v190) >= 2 {
                                            v190 = v190[1:len(v190)]
                                        }
                                        v209 := GenS25{F0: (v1 >> ((92) & 63))}
                                        _ = v209
                                    }
                                    v210 := GenS26{F0: 2147483647}
                                    v211 := GenS5{F0: v5}
                                    if true {
                                        if true {
                                            var v212 uint16 = 32767
                                            _ = v212
                                        }
                                        if false {
                                            v21 += (v21 << ((509) & 63))
                                            switch (fi154 % ((fi158) | 1)) {
                                            case 4294967295:
                                                v213 := GenS27{F0: -726, F1: 2147483648, F2: (v21 & (v21 % (((v4 << ((0) & 63))) | 1)))}
                                                _ = v213
                                            case -2147483648:
                                                v6[9223372036854775807] = fi158
                                            default:
                                                v14 += v14
                                            }
                                        }
                                    } else {
                                        var v214 int16 = v15
                                        v4 += 127
                                        _ = v214
                                    }
                                    _ = v210
                                    _ = v211
                                }
                                _ = v193
                                _ = v194
                            }
                        }
                        fmt.Println(v14)
                        _ = v190
                    case 628:
                        var v215 uint16 = 255
                        _ = v215
                    case 128:
                        v216 := GenS1{F0: v2}
                        _ = v216
                    }
                case "a":
                    fmt.Println(v14, v5)
                default:
                    var v217 uint64 = 4294967295
                    _ = v217
                }
                v218 := []int{fi154, (fi158 + int(v14)), 962}
                _ = v218
            }
            fmt.Println(v21, v5)
            v219 := v3[(fi158 / ((fi154) | 1))]
            v14 = 35
            _ = fi158
            _ = v219
        }
        _ = fi154
    }
    delete(v3, -2147483648)
    v220 := GenS24{F0: (^(v4)), F1: v5}
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v21
    _ = v220
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
