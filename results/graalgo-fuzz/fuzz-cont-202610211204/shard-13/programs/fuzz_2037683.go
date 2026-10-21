package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int64; F1 uint; F2 int16 }
type GenS4 struct { F0 int64; F1 uint64; F2 uint8 }
type GenS5 struct { F0 int16; F1 uint8; F2 int8 }
type GenS6 struct { F0 string; F1 uint64 }
type GenS7 struct { F0 string; F1 int32 }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 uint8; F1 int32 }
type GenS11 struct { F0 uint64; F1 int8; F2 int16 }
type GenS12 struct { F0 int8; F1 uint64 }
type GenS13 struct { F0 int64 }

func main() {
    var v1 string = (("" + ("hello" + "a")) + ("foo bar" + ""))
    v2 := []uint32{840}
    v3 := []int{32768, 871, 127, -520}
    var v4 int = 384
    v5 := make(map[int]int)
    for fi6 := 0; fi6 < 10; fi6++ {
        for fi7 := 0; fi7 < 16; fi7++ {
            var v8 int32 = -2147483648
            _ = fi7
            _ = v8
        }
        var v9 uint32 = 2
        v10 := GenS1{F0: "foo bar", F1: 229}
        _ = fi6
        _ = v10
        _ = v9
    }
    v4 -= v4
    v1 = v1
    if ((v1 + v1) < v1) {
        var v11 uint64 = 1
        var v12 uint8 = 128
        _ = v11
        _ = v12
    }
    fmt.Println(v4, v4)
    for fi13 := 0; fi13 < 4; fi13++ {
        var v14 uint16 = 32767
        delete(v5, 192)
        _ = fi13
        _ = v14
    }
    _ = len(v2)
    var v15 float64 = ((1.5 - 0.0) + float64(0.1))
    v16 := GenS1{F0: v1, F1: 128}
    var v17 float32 = 0.1
    v18 := GenS2{F0: 508}
    var v19 string = (v1 + v1)
    if (("foo bar" < v19) || (v17 <= (1.5 / v17))) {
        if true {
            switch (v1 + v1) {
            case "":
                if (true || (363 <= 65535)) {
                    var v20 bool = false
                    v4 -= (^(v4))
                    switch (v4 + 507) {
                    case -2147483648:
                        v19 = v1
                    case -486:
                        v21 := GenS3{F0: -1, F1: 42, F2: 1000}
                        _ = v21
                    }
                    _ = v20
                } else {
                    if false {
                        var v22 uint8 = 185
                        v23 := GenS1{F0: (v1 + "hello"), F1: 255}
                        _ = v22
                        _ = v23
                    } else {
                        var v24 int16 = 1000
                        fmt.Println(v24, v19, v17)
                        var v25 string = "hello"
                        _ = v24
                        _ = v25
                    }
                }
            }
        } else {
            var v26 string = (v19 + "")
            _ = v26
        }
    }
    switch v4 {
    case -9223372036854775808:
        switch ("" + "a") {
        case "hello":
            if true {
                var v27 uint16 = 42
                v28 := GenS3{F0: 864, F1: 1, F2: -231}
                _ = v27
                _ = v28
            }
        case "foo bar":
            if false {
                if true {
                    v4 = v4
                    if true {
                        v29 := GenS4{F0: -9223372036854775808, F1: 0, F2: 152}
                        fmt.Println(v17, v15, v15)
                        _ = v29
                    }
                    fmt.Println(v19, v19, v17)
                } else {
                    v30 := GenS5{F0: -1, F1: 2, F2: -128}
                    if true {
                        if true {
                            fmt.Println(v4, v1)
                            switch v19 {
                            case "hello":
                                var v31 bool = true
                                _ = v31
                            case "foo bar":
                                switch 1 {
                                case 32768:
                                    for fi32 := 0; fi32 < 13; fi32++ {
                                        fmt.Println(v17)
                                        var v33 int64 = -9223372036854775808
                                        var v34 uint32 = 4294967295
                                        _ = fi32
                                        _ = v33
                                        _ = v34
                                    }
                                case 128:
                                    if (v15 >= float64(float32(v15))) {
                                        var v35 int64 = 341
                                        if true {
                                            v36 := GenS1{F0: v1, F1: 46}
                                            _ = v36
                                        } else {
                                            if (42 > -1) {
                                                var v37 int64 = v35
                                                for fi38 := 0; fi38 < 5; fi38++ {
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                    var v39 int32 = -177
                                                    _ = fi38
                                                    _ = v39
                                                }
                                                _ = v37
                                            }
                                            v3 = append(v3, v4)
                                        }
                                        _ = v35
                                    } else {
                                        var v40 float64 = v15
                                        if false {
                                            v1 = v19
                                            switch "hello" {
                                            case "hello":
                                                switch ("foo bar" + (v19 + ((((v1 + v19) + ("foo bar" + v19)) + ((v19 + (v19 + "\n\t")) + v1)) + (("\n\t" + "hello") + "hello")))) {
                                                case "a":
                                                    var v41 uint64 = 0
                                                    _ = v41
                                                case "":
                                                    var v42 int16 = 68
                                                    _ = v42
                                                }
                                            case "a":
                                                if (2147483647 <= 1) {
                                                    if (255 >= 474) {
                                                        v43 := GenS6{F0: "foo bar", F1: 196}
                                                        for fi44 := 0; fi44 < 4; fi44++ {
                                                            var v45 string = ((v1 + v19) + (((v1 + "\n\t") + (v1 + v1)) + (((v1 + (("hello" + ("a" + "a")) + (v1 + (("foo bar" + "foo bar") + v19)))) + v1) + (v19 + "foo bar"))))
                                                            var v46 int = (v4 - (fi44 - (fi44 ^ 1)))
                                                            switch v1 {
                                                            case "foo bar":
                                                                var v47 int8 = 1
                                                                _ = v47
                                                            case "":
                                                                switch v45 {
                                                                case "a":
                                                                    v15 += v15
                                                                }
                                                            }
                                                            _ = fi44
                                                            _ = v45
                                                            _ = v46
                                                        }
                                                        for fi48 := 0; fi48 < 7; fi48++ {
                                                            var v49 uint = 1
                                                            var v50 uint32 = 2147483648
                                                            _ = fi48
                                                            _ = v49
                                                            _ = v50
                                                        }
                                                        _ = v43
                                                    } else {
                                                        var v51 uint64 = 2
                                                        fmt.Println(v15, v1, v15)
                                                        if (919 != 42) {
                                                            v52 := GenS7{F0: v1, F1: 775}
                                                            if false {
                                                                for fi53 := 0; fi53 < 13; fi53++ {
                                                                    for fi54 := 0; fi54 < 18; fi54++ {
                                                                        v5[629] = (fi53 - 0)
                                                                        var v55 uint8 = 2
                                                                        var v56 float32 = ((0.0 + 1.5) / (v17 - v17))
                                                                        _ = fi54
                                                                        _ = v55
                                                                        _ = v56
                                                                    }
                                                                    _ = fi53
                                                                }
                                                                if true {
                                                                    v57 := GenS8{F0: ((v40 <= (float64((1e10 * v17)) + v40)) || true)}
                                                                    fmt.Println(v1, v19)
                                                                    for fi58 := 0; fi58 < 18; fi58++ {
                                                                        for fi59 := 0; fi59 < 16; fi59++ {
                                                                            if (!(42 < -1)) {
                                                                                var v60 float64 = 1e10
                                                                                _ = v60
                                                                            }
                                                                            _ = fi59
                                                                        }
                                                                        _ = fi58
                                                                    }
                                                                    _ = v57
                                                                }
                                                            } else {
                                                                fmt.Println(v15)
                                                                if len(v3) > 0 {
                                                                    v3[0] = -913
                                                                }
                                                                if false {
                                                                    v61 := GenS9{F0: v51}
                                                                    for fi62 := 0; fi62 < 7; fi62++ {
                                                                        if true {
                                                                            v63 := GenS10{F0: 128, F1: 2147483647}
                                                                            if (false && (1 != -495)) {
                                                                                v64 := GenS11{F0: 715, F1: 1, F2: 42}
                                                                                if (-65 > 2147483647) {
                                                                                    var v65 int16 = 799
                                                                                    _ = v65
                                                                                } else {
                                                                                    if false {
                                                                                        switch 0 {
                                                                                        case 2147483647:
                                                                                            _ = len(v5)
                                                                                        case -123:
                                                                                            var v66 float64 = float64(v17)
                                                                                            _ = v66
                                                                                        case -2147483648:
                                                                                            var v67 uint16 = 65535
                                                                                            _ = v67
                                                                                        default:
                                                                                            var v68 int = (fi62 + -862)
                                                                                            _ = v68
                                                                                        }
                                                                                        var v69 string = ((v19 + v19) + "")
                                                                                        _ = v69
                                                                                    } else {
                                                                                        var v70 float32 = 0.0
                                                                                        v5[(fi62 / ((9223372036854775807) | 1))] = 617
                                                                                        _ = v70
                                                                                    }
                                                                                    v40 -= float64((float32(v15) - v17))
                                                                                }
                                                                                delete(v5, v4)
                                                                                _ = v64
                                                                            } else {
                                                                                for fi71 := 0; fi71 < 11; fi71++ {
                                                                                    if (0 < 128) {
                                                                                        var v72 uint16 = 65535
                                                                                        if len(v2) >= 2 {
                                                                                            v2 = v2[1:len(v2)]
                                                                                        }
                                                                                        v73 := v5[(fi62 ^ -13)]
                                                                                        _ = v72
                                                                                        _ = v73
                                                                                    } else {
                                                                                        if false {
                                                                                            var v74 int16 = 298
                                                                                            if (42 > 896) {
                                                                                                v1 = ""
                                                                                                switch (((v19 + "foo bar") + "\n\t") + v1) {
                                                                                                case "":
                                                                                                    v75 := GenS12{F0: 127, F1: 2}
                                                                                                    _ = v75
                                                                                                case "\n\t":
                                                                                                    _ = len(v5)
                                                                                                }
                                                                                            }
                                                                                            _ = v74
                                                                                        } else {
                                                                                            switch (fi71 + 4294967295) {
                                                                                            case 32767:
                                                                                                var v76 int8 = -10
                                                                                                _ = v76
                                                                                            case 127:
                                                                                                v77 := GenS8{F0: false}
                                                                                                _ = v77
                                                                                            default:
                                                                                                var v78 bool = false
                                                                                                _ = v78
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                    v3 = append(v3, (-(v4)))
                                                                                    switch "foo bar" {
                                                                                    case "":
                                                                                        v79 := GenS13{F0: 0}
                                                                                        _ = v79
                                                                                    case "foo bar":
                                                                                        var v80 int = (fi62 * (fi62 | -2147483648))
                                                                                        if false {
                                                                                            v40 -= 1.5
                                                                                            v19 = ((v1 + v1) + (v1 + v19))
                                                                                        }
                                                                                        var v81 float64 = (v40 * 3.14159)
                                                                                        _ = v80
                                                                                        _ = v81
                                                                                    case "hello":
                                                                                        var v82 string = v1
                                                                                        _ = v82
                                                                                    }
                                                                                    _ = fi71
                                                                                }
                                                                                var v83 int = (-(v4))
                                                                                _ = v83
                                                                            }
                                                                            _ = len(v5)
                                                                            _ = v63
                                                                        }
                                                                        _ = fi62
                                                                    }
                                                                    v15 = v15
                                                                    _ = v61
                                                                } else {
                                                                    v84 := GenS3{F0: 1, F1: 4294967295, F2: -777}
                                                                    _ = v84
                                                                }
                                                            }
                                                            _ = v52
                                                        } else {
                                                            var v85 int32 = 0
                                                            _ = v85
                                                        }
                                                        _ = v51
                                                    }
                                                    var v86 uint8 = 127
                                                    _ = v86
                                                }
                                            }
                                        }
                                        v4 -= v4
                                        _ = v40
                                    }
                                case 42:
                                    var v87 uint64 = 581
                                    _ = v87
                                default:
                                    fmt.Println(v15)
                                }
                            case "a":
                                switch v1 {
                                case "hello":
                                    if false {
                                        fmt.Println(v1, v4)
                                        v3 = append(v3, v4)
                                    } else {
                                        v1 = "a"
                                    }
                                case "":
                                    switch (v4 >> ((631) & 63)) {
                                    case -30:
                                        var v88 int32 = 0
                                        _ = v88
                                    case 32768:
                                        var v89 int8 = -1
                                        _ = v89
                                    }
                                }
                            default:
                                var v90 float64 = 0.1
                                _ = v90
                            }
                        }
                        for fi91 := 0; fi91 < 9; fi91++ {
                            var v92 bool = (1 != 224)
                            fmt.Println(v19)
                            _ = len(v5)
                            _ = fi91
                            _ = v92
                        }
                        for fi93 := 0; fi93 < 3; fi93++ {
                            for fi94 := 0; fi94 < 2; fi94++ {
                                v95 := v5[(fi94 << ((4294967295) & 63))]
                                _ = fi94
                                _ = v95
                            }
                            _ = fi93
                        }
                    } else {
                        var v96 uint32 = 4294967295
                        _ = v96
                    }
                    for fi97 := 0; fi97 < 6; fi97++ {
                        v98 := GenS7{F0: (v19 + v19), F1: -893}
                        var v99 float32 = -2.25
                        _ = fi97
                        _ = v98
                        _ = v99
                    }
                    _ = v30
                }
                for fi100 := 0; fi100 < 11; fi100++ {
                    var v101 uint32 = 2147483647
                    var v102 bool = false
                    _ = fi100
                    _ = v101
                    _ = v102
                }
            } else {
                for fi103 := 0; fi103 < 6; fi103++ {
                    if (!(("hello" != v1) || false)) {
                        var v104 int8 = 12
                        var v105 int = v4
                        var v106 int32 = -666
                        _ = v104
                        _ = v105
                        _ = v106
                    }
                    _ = fi103
                }
                v3 = append(v3, 4294967295)
            }
        }
    case 1:
        v15 = v15
    }
    for fi107 := 0; fi107 < 13; fi107++ {
        if true {
            switch (((v19 + (v19 + v19)) + ("" + v1)) + ((v1 + v1) + v1)) {
            case "\n\t":
                var v108 float64 = 1e10
                _ = v108
            case "hello":
                var v109 int64 = -262
                _ = v109
            case "foo bar":
                var v110 float64 = 3.14159
                _ = v110
            }
            var v111 uint16 = 0
            var v112 int16 = -135
            _ = v111
            _ = v112
        }
        var v113 int16 = 1
        fmt.Println(v4)
        _ = fi107
        _ = v113
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
