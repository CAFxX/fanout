package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint; F2 uint32 }
type GenS2 struct { F0 uint32; F1 uint16; F2 string }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint8; F1 uint64 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 int }
type GenS9 struct { F0 int32; F1 uint32 }
type GenS10 struct { F0 int32 }
type GenS11 struct { F0 int32; F1 uint16 }
type GenS12 struct { F0 uint16; F1 int32 }
func genFn0(p0 uint64, p1 int16, p2 uint32) int64 {
    var v6 int64 = 9223372036854775807
    p2 += (^(p2))
    var v7 uint8 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return (v6 << ((0) & 63))
}

func main() {
    v1 := make(map[int]uint32)
    v2 := []int{-539, 127, 561, 0}
    v3 := []string{"foo bar"}
    v4 := make(map[string]string)
    var v5 uint16 = 0
    switch "foo bar" {
    case "foo bar":
        var v8 uint = 255
        _ = v8
    case "":
        v9 := GenS1{F0: -916, F1: 0, F2: 42}
        _ = v9
    case "a":
        v5 = 32767
    }
    v5 += (v5 >> ((547) & 63))
    switch "hello" {
    case "":
        var v10 uint8 = 0
        _ = v10
    case "hello":
        if true {
            for fi11 := 0; fi11 < 15; fi11++ {
                for fi12 := 0; fi12 < 19; fi12++ {
                    if (2 >= 4294967295) {
                        fmt.Println(v5, fi11, fi11)
                        fmt.Println(fi11)
                        fmt.Println(fi12)
                    } else {
                        v13 := GenS2{F0: 1, F1: (v5 / (((v5 << ((4294967295) & 63))) | 1)), F2: ("\n\t" + ("foo bar" + "foo bar"))}
                        _ = v13
                    }
                    v5 = (v5 / (((v5 + v5)) | 1))
                    _ = fi12
                }
                for fi14 := 0; fi14 < 9; fi14++ {
                    if (42 >= 2147483648) {
                        if true {
                            v5 = 54
                        } else {
                            var v15 uint32 = 770
                            var v16 int32 = 1
                            _ = v15
                            _ = v16
                        }
                        var v17 float64 = 3.14159
                        if false {
                            v5 = 262
                            if true {
                                var v18 uint = 457
                                fmt.Println(fi14, fi11, fi11)
                                _ = v18
                            }
                            switch fi14 {
                            case -1:
                                switch (fi14 - (fi14 / (((fi14 >> ((42) & 63))) | 1))) {
                                case 127:
                                    var v19 int16 = 0
                                    _ = v19
                                case -85:
                                    v17 = v17
                                    var v20 uint16 = v5
                                    switch (fi14 << ((208) & 63)) {
                                    case -173:
                                        v21 := GenS3{F0: 1}
                                        _ = v21
                                    case -2147483648:
                                        fmt.Println(fi14, v17, v5)
                                    default:
                                        v1[fi14] = 2147483647
                                        fmt.Println(v17, v20)
                                    }
                                    _ = v20
                                case 15:
                                    var v22 int8 = -124
                                    _ = v22
                                }
                                v5 = (^(v5))
                            case 350:
                                if true {
                                    var v23 float32 = 0.1
                                    var v24 string = ("\n\t" + "\n\t")
                                    if true {
                                        if true {
                                            v25 := GenS1{F0: (fi11 - (-(fi11))), F1: 1, F2: 2147483648}
                                            switch 1 {
                                            case -806:
                                                if (1 < 4294967295) {
                                                    v24 = (("\n\t" + ("foo bar" + v24)) + "foo bar")
                                                    var v26 uint64 = 2
                                                    _ = v26
                                                } else {
                                                    var v27 int16 = 985
                                                    v28 := GenS4{F0: 128}
                                                    _ = v27
                                                    _ = v28
                                                }
                                            case 127:
                                                switch v24 {
                                                case "foo bar":
                                                    var v29 float32 = 0.0
                                                    _ = v29
                                                default:
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                }
                                            default:
                                                switch 395 {
                                                case 1:
                                                    if true {
                                                        switch ((v24 + ("\n\t" + v24)) + "hello") {
                                                        case "":
                                                            switch v24 {
                                                            case "":
                                                                v23 = (v23 * (v23 + (v23 * float32(v17))))
                                                            default:
                                                                var v30 int = (fi14 & fi11)
                                                                v31 := GenS5{F0: 0, F1: 0}
                                                                _ = v30
                                                                _ = v31
                                                            }
                                                        case "hello":
                                                            _ = len(v3)
                                                        case "foo bar":
                                                            var v32 uint8 = 1
                                                            _ = v32
                                                        }
                                                        if len(v3) >= 2 {
                                                            v3 = v3[1:len(v3)]
                                                        }
                                                        if true {
                                                            v33 := GenS4{F0: 255}
                                                            var v34 int8 = 0
                                                            switch (fi14 & (fi14 - (fi11 >> ((204) & 63)))) {
                                                            case -2147483648:
                                                                v35 := GenS6{F0: true}
                                                                _ = v35
                                                            default:
                                                                var v36 uint = 18446744073709551615
                                                                v24 = (v24 + ("a" + v24))
                                                                var v37 float32 = 1.5
                                                                _ = v36
                                                                _ = v37
                                                            }
                                                            _ = v33
                                                            _ = v34
                                                        }
                                                    }
                                                case -2147483648:
                                                    switch 0 {
                                                    case 0:
                                                        v5 -= 42
                                                    case -1:
                                                        var v38 float32 = v23
                                                        _ = v38
                                                    case -9223372036854775808:
                                                        v23 += v23
                                                    }
                                                case 127:
                                                    if false {
                                                        if (true || (0.1 >= (v17 - v17))) {
                                                            fmt.Println(v23, v5)
                                                        }
                                                        if false {
                                                            var v39 int64 = 223
                                                            v40 := GenS2{F0: 1, F1: v5, F2: "a"}
                                                            v41 := GenS7{F0: 229}
                                                            _ = v39
                                                            _ = v40
                                                            _ = v41
                                                        }
                                                        var v42 int32 = 42
                                                        fmt.Println(v5, fi11)
                                                        var v43 string = v24
                                                        var v44 string = (v24 + ("\n\t" + "foo bar"))
                                                        v42 -= v42
                                                        _ = v42
                                                        _ = v43
                                                        _ = v44
                                                    }
                                                default:
                                                    v45 := GenS4{F0: 42}
                                                    _ = v45
                                                }
                                            }
                                            if false {
                                                var v46 uint8 = 104
                                                if (!true) {
                                                    var v47 uint32 = 2
                                                    _ = v47
                                                }
                                                _ = v46
                                            } else {
                                                if (410 <= 415) {
                                                    v48 := v4[(v24 + ((v24 + v24) + "hello"))]
                                                    v23 = v23
                                                    _ = v48
                                                } else {
                                                    v2 = append(v2, (^(fi11)))
                                                    if false {
                                                        var v49 uint = 663
                                                        v50 := GenS8{F0: (fi11 ^ (fi11 << ((v49) & 63)))}
                                                        v51 := GenS7{F0: 42}
                                                        _ = v49
                                                        _ = v50
                                                        _ = v51
                                                    }
                                                    v4["a"] = ((v24 + v24) + v24)
                                                    fmt.Println(fi11)
                                                    if true {
                                                        v3 = append(v3, v24)
                                                        fmt.Println(v24)
                                                        fmt.Println(v5, v5)
                                                        v5 -= v5
                                                        _ = len(v3)
                                                    }
                                                    if (1 <= 255) {
                                                        if true {
                                                            fmt.Println(fi14)
                                                            v23 = -2.25
                                                            if ((((((v24 + v24) + v24) + ("\n\t" + "")) + (v24 + ("" + ((("foo bar" + "\n\t") + v24) + "foo bar")))) + "hello") < "hello") {
                                                                if false {
                                                                    var v52 bool = (!(v17 <= v17))
                                                                    _ = v52
                                                                }
                                                            }
                                                        }
                                                    } else {
                                                        _ = len(v2)
                                                        var v53 uint = 231
                                                        v3 = append(v3, v24)
                                                        _ = v53
                                                    }
                                                }
                                                v54 := GenS2{F0: 42, F1: (v5 ^ 32767), F2: ("a" + "")}
                                                if len(v2) > 0 {
                                                    v2[0] = (fi14 ^ fi14)
                                                }
                                                _ = v54
                                            }
                                            _ = v25
                                        }
                                        switch (fi11 ^ (fi14 << ((870) & 63))) {
                                        case 2147483647:
                                            delete(v4, v24)
                                        case 4294967295:
                                            var v55 int32 = -2147483648
                                            _ = v55
                                        }
                                    }
                                    _ = v23
                                    _ = v24
                                } else {
                                    var v56 float32 = ((1.5 + 1.5) - (1e10 + 0.1))
                                    if ((v17 / ((float64(0.0) / (((v17 - v17) * 3.14159) / -2.25)) - v17)) > float64(v56)) {
                                        fmt.Println(v17, fi11)
                                        fmt.Println(v56)
                                        v57 := GenS5{F0: 42, F1: 1}
                                        _ = v57
                                    } else {
                                        if (32767 != (v5 - v5)) {
                                            fmt.Println(v56, v56)
                                            v5 = v5
                                        } else {
                                            switch fi11 {
                                            case 542:
                                                fmt.Println(fi14, v56, fi11)
                                            case -849:
                                                if (((((!false) && (992 <= 18446744073709551615)) || true) || false) && false) {
                                                    switch "\n\t" {
                                                    case "foo bar":
                                                        var v58 string = ("\n\t" + "")
                                                        switch "a" {
                                                        case "":
                                                            v59 := GenS6{F0: true}
                                                            _ = v59
                                                        case "foo bar":
                                                            if false {
                                                                v56 = v56
                                                            }
                                                        case "\n\t":
                                                            v2 = append(v2, 805)
                                                            var v60 string = ("\n\t" + v58)
                                                            v61 := GenS2{F0: uint32(v5), F1: v5, F2: "hello"}
                                                            _ = v60
                                                            _ = v61
                                                        default:
                                                            if (fi11 > fi11) {
                                                                v62 := GenS7{F0: 42}
                                                                var v63 uint32 = 861
                                                                _ = v62
                                                                _ = v63
                                                            } else {
                                                                v2 = append(v2, 32767)
                                                            }
                                                        }
                                                        if (v17 > -2.25) {
                                                            v5 = (v5 | v5)
                                                        }
                                                        _ = v58
                                                    case "\n\t":
                                                        v3 = append(v3, "foo bar")
                                                    case "a":
                                                        if ((!(!((1.5 + (1e10 / v56)) <= v56))) || false) {
                                                            var v64 float64 = 3.14159
                                                            if (true && false) {
                                                                v65 := GenS9{F0: 328, F1: 319}
                                                                if true {
                                                                    if true {
                                                                        v66 := GenS10{F0: -2147483648}
                                                                        var v67 int64 = -9223372036854775808
                                                                        var v68 int32 = int32(v67)
                                                                        _ = v66
                                                                        _ = v67
                                                                        _ = v68
                                                                    } else {
                                                                        v5 += (v5 + v5)
                                                                        v69 := v1[(fi11 >> ((2) & 63))]
                                                                        _ = v69
                                                                    }
                                                                    v64 = v64
                                                                    var v70 uint32 = uint32(v5)
                                                                    _ = v70
                                                                } else {
                                                                    v56 += v56
                                                                    v17 = (1.5 + ((3.14159 / v17) - (float64(v56) / v64)))
                                                                }
                                                                _ = v65
                                                            } else {
                                                                var v71 uint16 = v5
                                                                switch (fi11 >> ((658) & 63)) {
                                                                case 42:
                                                                    fmt.Println(v64, v5)
                                                                default:
                                                                    switch 668 {
                                                                    case 42:
                                                                        if (!false) {
                                                                            v72 := GenS10{F0: -1}
                                                                            _ = v72
                                                                        }
                                                                        switch "\n\t" {
                                                                        case "foo bar":
                                                                            v5 = v5
                                                                        default:
                                                                            switch fi11 {
                                                                            case 392:
                                                                                switch "a" {
                                                                                case "":
                                                                                    var v73 float32 = (1e10 * 0.0)
                                                                                    _ = v73
                                                                                default:
                                                                                    fmt.Println(fi14, v56, v5)
                                                                                }
                                                                            case 1:
                                                                                delete(v1, fi14)
                                                                                fmt.Println(v17)
                                                                            default:
                                                                                var v74 uint8 = 26
                                                                                _ = v74
                                                                            }
                                                                        }
                                                                    case 128:
                                                                        var v75 int = 127
                                                                        _ = v75
                                                                    case -337:
                                                                        v17 -= float64(v56)
                                                                    default:
                                                                        v5 = 32767
                                                                    }
                                                                }
                                                                v71 -= v5
                                                                var v76 string = "hello"
                                                                _ = v71
                                                                _ = v76
                                                            }
                                                            if true {
                                                                v17 -= v64
                                                                var v77 uint = 505
                                                                if false {
                                                                    if true {
                                                                        var v78 int8 = -94
                                                                        _ = v78
                                                                    } else {
                                                                        if (false && false) {
                                                                            fmt.Println(v17, v64, fi11)
                                                                            v64 += (((v64 * (float64((float32(-2.25) - float32(((v17 - 1e10) + v64)))) + (v17 - 3.14159))) * v64) / 1.0)
                                                                            v56 = v56
                                                                            var v79 uint16 = 937
                                                                            fmt.Println(v56)
                                                                            _ = v79
                                                                        }
                                                                        v56 = ((v56 / ((-2.25 - v56) / v56)) + 0.0)
                                                                    }
                                                                } else {
                                                                    v80 := GenS6{F0: (!(!false))}
                                                                    v17 += 1.5
                                                                    if ((42 >= 1) || true) {
                                                                        var v81 float64 = v64
                                                                        var v82 uint8 = 23
                                                                        _ = v81
                                                                        _ = v82
                                                                    }
                                                                    _ = v80
                                                                }
                                                                _ = v77
                                                            } else {
                                                                if (375 != 2147483648) {
                                                                    v17 += (v17 * float64(1e10))
                                                                    _ = len(v1)
                                                                    var v83 float32 = 0.0
                                                                    _ = v83
                                                                }
                                                                v84 := GenS1{F0: (fi11 * (fi14 ^ fi14)), F1: 239, F2: 1}
                                                                v85 := v4["\n\t"]
                                                                v86 := GenS2{F0: 65535, F1: (v5 / (((v5 << ((0) & 63))) | 1)), F2: v85}
                                                                _ = v84
                                                                _ = v85
                                                                _ = v86
                                                            }
                                                            _ = v64
                                                        }
                                                    }
                                                    v5 = (^(v5))
                                                } else {
                                                    if false {
                                                        v1[fi14] = 65535
                                                        var v87 int32 = 42
                                                        _ = v87
                                                    }
                                                    v56 -= 1e10
                                                    switch fi11 {
                                                    case 56:
                                                        if false {
                                                            if ("" == "foo bar") {
                                                                var v88 float32 = v56
                                                                var v89 int32 = 2147483647
                                                                v90 := GenS7{F0: 0}
                                                                var v91 bool = (-429 > 1)
                                                                v4["hello"] = "foo bar"
                                                                _ = v88
                                                                _ = v89
                                                                _ = v90
                                                                _ = v91
                                                            } else {
                                                                v5 = (v5 - (v5 * (^(v5))))
                                                                var v92 uint32 = 49
                                                                if len(v2) >= 2 {
                                                                    v2 = v2[1:len(v2)]
                                                                }
                                                                _ = v92
                                                            }
                                                        } else {
                                                            fmt.Println(v56)
                                                            v93 := GenS11{F0: 569, F1: v5}
                                                            _ = v93
                                                        }
                                                    default:
                                                        var v94 uint8 = 1
                                                        _ = v94
                                                    }
                                                }
                                            case 0:
                                                v17 = (((0.0 - 1e10) + v17) - 0.1)
                                            default:
                                                v95 := GenS12{F0: (v5 | 518), F1: 566}
                                                _ = v95
                                            }
                                            v17 -= v17
                                            v1[fi14] = 635
                                        }
                                    }
                                    switch ("" + "foo bar") {
                                    case "hello":
                                        v17 -= v17
                                    default:
                                        if len(v2) > 0 {
                                            v2[0] = fi11
                                        }
                                        var v96 int8 = 83
                                        v17 += v17
                                        var v97 int32 = -1
                                        fmt.Println(v17, v5)
                                        _ = v96
                                        _ = v97
                                    }
                                    _ = v56
                                }
                            case 2147483647:
                                var v98 string = ("foo bar" + "")
                                var v99 uint16 = 42
                                _ = v98
                                _ = v99
                            default:
                                v100 := GenS6{F0: ((("hello" + "foo bar") > "") && false)}
                                if true {
                                    switch 9223372036854775807 {
                                    case -929:
                                        fmt.Println(fi14)
                                    default:
                                        v17 = v17
                                    }
                                    var v101 uint16 = v5
                                    _ = v101
                                }
                                var v102 uint32 = 65535
                                _ = len(v3)
                                var v103 uint = 0
                                _ = v100
                                _ = v102
                                _ = v103
                            }
                            v5 = v5
                        }
                        _ = v17
                    }
                    var v104 bool = ((float32(3.14159) - (0.1 * (3.14159 / 1e10))) <= 0.0)
                    _ = fi14
                    _ = v104
                }
                v5 = v5
                _ = fi11
            }
        }
    }
    if false {
        if len(v2) > 0 {
            v2[0] = 2147483647
        }
    }
    var v105 uint64 = 2
    _ = len(v1)
    v5 = (v5 * (^(v5)))
    v106 := GenS11{F0: -2147483648, F1: uint16(v105)}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v105:", v105)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v105
    _ = v106
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
