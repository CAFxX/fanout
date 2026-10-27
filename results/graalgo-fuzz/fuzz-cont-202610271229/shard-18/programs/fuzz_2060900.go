package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint16; F1 int64; F2 int }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int8; F1 int8; F2 bool }
type GenS6 struct { F0 uint32; F1 uint }
type GenS7 struct { F0 string; F1 string; F2 uint }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 uint32; F1 uint16 }
type GenS10 struct { F0 int }
type GenS11 struct { F0 uint64; F1 int32; F2 uint8 }
type GenS12 struct { F0 uint16; F1 uint8 }
type GenS13 struct { F0 int }
type GenS14 struct { F0 int }
type GenS15 struct { F0 uint64; F1 uint }
type GenS16 struct { F0 bool; F1 string; F2 uint32 }
type GenS17 struct { F0 int32; F1 uint16 }
type GenS18 struct { F0 string; F1 int64 }
type GenS19 struct { F0 uint8; F1 int16; F2 uint64 }
type GenS20 struct { F0 string }
type GenS21 struct { F0 int32; F1 bool; F2 uint32 }
func genFn0(p0 bool) int64 {
    if p0 {
        var v4 int32 = 1
        v4 += v4
        v4 = 42
        _ = v4
    }
    p0 = (!true)
    _ = p0
    return 9223372036854775807
}
func genFn1(p0 int64, p1 string) uint32 {
    for fi5 := 0; fi5 < 7; fi5++ {
        if false {
            v6 := []int{}
            var v7 float64 = 1e10
            _ = v6
            _ = v7
        }
        v8 := GenS1{F0: 207, F1: 127}
        p0 += p0
        _ = fi5
        _ = v8
    }
    if (!(!false)) {
        var v9 int = 2147483647
        v10 := GenS1{F0: 255, F1: 100}
        _ = v10
        _ = v9
    } else {
        v11 := []int{663, 127}
        _ = v11
    }
    fmt.Println(p0, p1)
    p1 = p1
    _ = p0
    _ = p1
    return 2
}

func main() {
    var v1 int32 = 0
    v2 := []string{}
    v3 := []int{}
    var v12 int8 = -128
    v13 := GenS2{F0: 83}
    var v14 string = ("hello" + (("\n\t" + "a") + "foo bar"))
    if false {
        for fi15 := 0; fi15 < 8; fi15++ {
            v1 -= (^(v1))
            var v16 string = "a"
            _ = fi15
            _ = v16
        }
        if len(v2) > 0 {
            v2[0] = "hello"
        }
        if (!true) {
            if len(v3) > 0 {
                v3[0] = 128
            }
            switch 32768 {
            case 921:
                var v17 uint = 2
                _ = v17
            }
            fmt.Println(v1)
        } else {
            v14 = v14
            for fi18 := 0; fi18 < 16; fi18++ {
                if true {
                    v12 = (v12 * (v12 - (v12 % ((-1) | 1))))
                }
                v1 = (v1 % ((v1) | 1))
                _ = fi18
            }
            v12 -= (^(v12))
        }
    }
    for fi19 := 0; fi19 < 7; fi19++ {
        for fi20 := 0; fi20 < 11; fi20++ {
            var v21 uint8 = 0
            v12 = v12
            v14 = v14
            _ = fi20
            _ = v21
        }
        v1 = (-(v1))
        switch "foo bar" {
        case "hello":
            var v22 uint64 = 42
            _ = v22
        case "\n\t":
            for fi23 := 0; fi23 < 9; fi23++ {
                v12 += v12
                var v24 int32 = v1
                if (true || false) {
                    var v25 int32 = (v24 << ((398) & 63))
                    fmt.Println(v14, fi19)
                    var v26 uint32 = 2147483647
                    _ = v25
                    _ = v26
                }
                v1 = (v24 & 42)
                var v27 uint64 = 4294967295
                _ = fi23
                _ = v24
                _ = v27
            }
        case "a":
            if true {
                if ((1000 <= 0) || (!(1 > 2))) {
                    v12 = -88
                    for fi28 := 0; fi28 < 18; fi28++ {
                        var v29 int64 = 92
                        var v30 uint = 4294967295
                        v29 = (v29 ^ (v29 + genFn0(false)))
                        _ = fi28
                        _ = v29
                        _ = v30
                    }
                }
            } else {
                var v31 int32 = v1
                _ = v31
            }
        default:
            v14 = (("hello" + v14) + (v14 + v14))
        }
        _ = fi19
    }
    var v32 int64 = 1
    var v33 int16 = 42
    for fi34 := 0; fi34 < 17; fi34++ {
        for fi35 := 0; fi35 < 7; fi35++ {
            if false {
                switch (fi35 & (^(fi34))) {
                case 1:
                    if len(v2) > 0 {
                        v2[0] = v14
                    }
                case 2147483647:
                    v36 := GenS1{F0: 300, F1: v12}
                    _ = v36
                case 951:
                    v1 = (v1 & -2147483648)
                default:
                    v12 = (v12 | 100)
                }
            } else {
                if true {
                    v37 := GenS2{F0: 255}
                    var v38 float32 = 3.14159
                    switch (("" + "foo bar") + v14) {
                    case "a":
                        v32 = (^(v32))
                    }
                    _ = v37
                    _ = v38
                } else {
                    if true {
                        v39 := GenS1{F0: 0, F1: v12}
                        _ = v39
                    } else {
                        var v40 uint = 65535
                        _ = v40
                    }
                    v33 -= (v33 | (v33 << ((65535) & 63)))
                    if false {
                        switch v14 {
                        case "a":
                            var v41 int16 = v33
                            _ = v41
                        case "foo bar":
                            fmt.Println(v14, fi34)
                        case "\n\t":
                            v42 := GenS3{F0: 81, F1: v32, F2: fi35}
                            var v43 int8 = -1
                            v43 += 0
                            _ = v42
                            _ = v43
                        default:
                            v44 := GenS1{F0: 0, F1: 127}
                            _ = v44
                        }
                    } else {
                        v33 += (^(v33))
                        var v45 int = fi35
                        _ = v45
                    }
                    v46 := GenS4{F0: v32}
                    if len(v3) > 0 {
                        v3[0] = (^(fi35))
                    }
                    _ = v46
                }
                if (("hello" < "\n\t") || ((true || ((fi35 / ((int(v32)) | 1)) >= (^(fi34)))) && (42 <= 4294967295))) {
                    v47 := GenS5{F0: 100, F1: v12, F2: true}
                    v12 = v12
                    v14 = v14
                    _ = v47
                } else {
                    v48 := GenS6{F0: 148, F1: 0}
                    var v49 uint64 = 42
                    var v50 int64 = int64(v1)
                    _ = v48
                    _ = v49
                    _ = v50
                }
                v51 := GenS7{F0: "a", F1: v14, F2: 18446744073709551615}
                _ = v51
            }
            _ = fi35
        }
        v2 = append(v2, "a")
        _ = fi34
    }
    var v52 uint = 66
    var v53 int64 = v32
    if (2 > 130) {
        v54 := GenS3{F0: 2, F1: (-(v53)), F2: -9223372036854775808}
        if false {
            if (true || true) {
                var v55 int32 = v1
                v56 := GenS8{F0: 125}
                _ = v55
                _ = v56
            }
            var v57 uint8 = 255
            _ = v57
        }
        _ = v54
    }
    var v58 string = v14
    for fi59 := 0; fi59 < 6; fi59++ {
        var v60 int32 = (v1 ^ -137)
        switch (-(fi59)) {
        case 127:
            v60 += (v60 ^ -198)
        }
        _ = fi59
        _ = v60
    }
    v61 := GenS9{F0: 490, F1: 32767}
    _ = len(v2)
    fmt.Println(v33)
    v62 := GenS10{F0: 436}
    for fi63 := 0; fi63 < 3; fi63++ {
        if false {
            if (false || false) {
                if (int16(v52) > v33) {
                    var v64 int8 = -1
                    _ = v64
                } else {
                    switch ((v14 + "foo bar") + v14) {
                    case "\n\t":
                        v12 -= 44
                    case "foo bar":
                        v65 := GenS11{F0: 2, F1: (v1 % ((v1) | 1)), F2: 42}
                        _ = v65
                    case "hello":
                        var v66 uint32 = 78
                        _ = v66
                    }
                    v33 -= -1
                    for fi67 := 0; fi67 < 3; fi67++ {
                        var v68 uint64 = 2
                        _ = fi67
                        _ = v68
                    }
                }
                fmt.Println(fi63, v33)
                for fi69 := 0; fi69 < 15; fi69++ {
                    switch ("foo bar" + v14) {
                    case "\n\t":
                        if true {
                            switch (fi69 & (^(fi63))) {
                            case 239:
                                switch ("hello" + (v14 + ("a" + v14))) {
                                case "a":
                                    if false {
                                        v70 := GenS12{F0: 0, F1: 25}
                                        v71 := GenS13{F0: (fi69 & (^(fi69)))}
                                        _ = v70
                                        _ = v71
                                    } else {
                                        var v72 int16 = -32768
                                        var v73 string = ("\n\t" + v58)
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        _ = v72
                                        _ = v73
                                    }
                                }
                            default:
                                if (2147483647 < 2) {
                                    var v74 string = v14
                                    if (2147483647 == 1) {
                                        var v75 int64 = (v32 ^ (-(v53)))
                                        _ = v75
                                    }
                                    switch v58 {
                                    case "\n\t":
                                        v76 := GenS9{F0: 599, F1: 42}
                                        _ = v76
                                    case "hello":
                                        var v77 int8 = 0
                                        _ = v77
                                    default:
                                        var v78 uint64 = 1
                                        _ = v78
                                    }
                                    _ = v74
                                }
                            }
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            v33 -= (-(v33))
                        }
                    }
                    switch v14 {
                    case "a":
                        v3 = append(v3, 689)
                    case "hello":
                        var v79 uint8 = 1
                        _ = v79
                    case "":
                        if true {
                            var v80 int8 = -128
                            switch 127 {
                            case 541:
                                if (!true) {
                                    if len(v2) > 0 {
                                        v2[0] = ((v14 + v14) + v58)
                                    }
                                    var v81 bool = ((658 <= 18446744073709551615) && false)
                                    fmt.Println(fi63)
                                    _ = v81
                                } else {
                                    v82 := GenS14{F0: fi69}
                                    v83 := GenS3{F0: 694, F1: v53, F2: fi69}
                                    _ = v82
                                    _ = v83
                                }
                            default:
                                v3 = append(v3, (fi69 / (((fi63 ^ 9223372036854775807)) | 1)))
                            }
                            _ = v80
                        } else {
                            v14 = v14
                            switch fi69 {
                            case 9223372036854775807:
                                switch (fi69 >> (((v52 - (^(v52)))) & 63)) {
                                case -9223372036854775808:
                                    v84 := GenS15{F0: 42, F1: 67}
                                    _ = v84
                                }
                            case -45:
                                var v85 bool = true
                                _ = v85
                            default:
                                switch ("" + v58) {
                                case "hello":
                                    if ((((("hello" + "hello") + (("\n\t" + "hello") + ("foo bar" + v58))) + v14) == "\n\t") && (!(1 != 38))) {
                                        if true {
                                            if true {
                                                if false {
                                                    if (!(!((v32 / ((21) | 1)) < (v53 >> ((42) & 63))))) {
                                                        v14 = v14
                                                        v12 = (v12 << ((v52) & 63))
                                                        v86 := GenS10{F0: 128}
                                                        _ = v86
                                                    } else {
                                                        if false {
                                                            var v87 uint16 = 892
                                                            fmt.Println(v87, v52)
                                                            _ = v87
                                                        } else {
                                                            v3 = append(v3, 32768)
                                                            v2 = append(v2, v58)
                                                            _ = len(v2)
                                                            v88 := GenS16{F0: true, F1: v58, F2: 0}
                                                            _ = v88
                                                        }
                                                    }
                                                    switch ("" + (v14 + v58)) {
                                                    case "":
                                                        v52 += (v52 ^ (v52 >> ((v52) & 63)))
                                                    case "foo bar":
                                                        var v89 uint8 = 253
                                                        _ = v89
                                                    default:
                                                        var v90 string = ((v14 + v58) + "hello")
                                                        _ = v90
                                                    }
                                                }
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                            }
                                            if (3.14159 == (float32((1e10 - (1e10 + (-2.25 * ((1.5 * (0.1 * 0.1)) - -2.25))))) + 1e10)) {
                                                v91 := GenS1{F0: 42, F1: v12}
                                                _ = v91
                                            }
                                            fmt.Println(v58, v32, v12)
                                        }
                                        var v92 string = (v58 + "a")
                                        if false {
                                            if (true || (507 > 4294967295)) {
                                                if ((float32(1.5) + 0.1) > -2.25) {
                                                    v93 := GenS16{F0: (false || (fi69 > fi69)), F1: v14, F2: 0}
                                                    v92 = "hello"
                                                    if len(v2) > 0 {
                                                        v2[0] = v58
                                                    }
                                                    _ = v93
                                                } else {
                                                    v14 = v14
                                                }
                                            }
                                            fmt.Println(fi69, v53, v58)
                                            if false {
                                                var v94 int32 = 1
                                                _ = v94
                                            } else {
                                                v33 -= (-(v33))
                                                if true {
                                                    if ((v1 & v1) >= (v1 % (((^(v1))) | 1))) {
                                                        v95 := GenS17{F0: (v1 - (v1 + (v1 % (((v1 / ((v1) | 1))) | 1)))), F1: 0}
                                                        fmt.Println(v58, v58, v1)
                                                        _ = v95
                                                    }
                                                    v96 := GenS16{F0: (-2.25 >= (3.14159 - 0.1)), F1: ("a" + ((v14 + v92) + ("a" + "a"))), F2: genFn1(-9223372036854775808, "")}
                                                    if ((fi63 <= fi69) && true) {
                                                        v3 = append(v3, fi69)
                                                        fmt.Println(v52, v52, v33)
                                                        fmt.Println(v32, v32, v14)
                                                    } else {
                                                        v58 = v92
                                                    }
                                                    _ = v96
                                                } else {
                                                    v32 -= v32
                                                    v97 := GenS8{F0: 172}
                                                    _ = v97
                                                }
                                                var v98 uint8 = 127
                                                _ = v98
                                            }
                                        } else {
                                            v52 -= (v52 & v52)
                                        }
                                        var v99 bool = false
                                        _ = v92
                                        _ = v99
                                    } else {
                                        var v100 uint8 = 42
                                        _ = v100
                                    }
                                case "":
                                    v32 = v53
                                default:
                                    v14 = (v58 + "\n\t")
                                }
                            }
                            v52 = v52
                        }
                    default:
                        if (127 < int(v33)) {
                            var v101 string = v58
                            var v102 bool = (!((!(!true)) || (!false)))
                            var v103 int32 = 42
                            _ = v101
                            _ = v102
                            _ = v103
                        }
                    }
                    v104 := GenS11{F0: 42, F1: (^(v1)), F2: 2}
                    _ = fi69
                    _ = v104
                }
            }
            var v105 int8 = (v12 / ((40) | 1))
            _ = v105
        }
        var v106 int = fi63
        for fi107 := 0; fi107 < 2; fi107++ {
            switch 9223372036854775807 {
            case 4294967295:
                if true {
                    if true {
                        v32 += genFn0(true)
                        if ((v33 >> (((v52 - v52)) & 63)) < v33) {
                            v2 = append(v2, v58)
                            v1 = v1
                            fmt.Println(v106, v14)
                        }
                        switch (("hello" + v58) + (v58 + (v58 + ""))) {
                        case "hello":
                            v108 := GenS5{F0: (v12 * -122), F1: v12, F2: (("a" + v58) == v14)}
                            _ = v108
                        default:
                            v106 -= 1
                        }
                        var v109 int = v106
                        switch (v109 / (((v109 ^ 4294967295)) | 1)) {
                        case 1:
                            var v110 string = v58
                            _ = v110
                        case 0:
                            var v111 uint32 = 210
                            _ = v111
                        case 9223372036854775807:
                            v106 = -9223372036854775808
                        }
                        _ = v109
                    }
                } else {
                    switch 127 {
                    case 731:
                        var v112 uint16 = 0
                        _ = v112
                    case 32767:
                        var v113 float64 = ((0.0 / (1.5 - 3.14159)) / 1.0)
                        _ = v113
                    case 4294967295:
                        if false {
                            var v114 uint16 = 0
                            _ = len(v3)
                            _ = v114
                        } else {
                            if true {
                                var v115 int16 = v33
                                _ = v115
                            } else {
                                v116 := GenS18{F0: ((((v58 + v58) + v14) + v58) + v58), F1: v32}
                                _ = v116
                            }
                        }
                    }
                    v117 := GenS19{F0: 127, F1: 1000, F2: 143}
                    switch v106 {
                    case -850:
                        switch (fi107 - fi107) {
                        case 1:
                            if false {
                                switch v58 {
                                case "foo bar":
                                    v33 = (v33 / (((v33 / ((42) | 1))) | 1))
                                default:
                                    if (!true) {
                                        v118 := GenS20{F0: "foo bar"}
                                        _ = v118
                                    }
                                }
                                var v119 string = v58
                                _ = v119
                            }
                        default:
                            var v120 int32 = (v1 ^ (v1 / ((v1) | 1)))
                            _ = v120
                        }
                    case -9223372036854775808:
                        v121 := GenS21{F0: v1, F1: false, F2: 4294967295}
                        _ = v121
                    }
                    _ = v117
                }
            default:
                if (-454 > (v1 | (v1 | -1))) {
                    var v122 int32 = v1
                    var v123 int32 = v1
                    v32 = -136
                    _ = v122
                    _ = v123
                }
            }
            _ = fi107
        }
        _ = fi63
        _ = v106
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    fmt.Println("v58:", v58)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v52
    _ = v53
    _ = v58
    _ = v61
    _ = v62
}
