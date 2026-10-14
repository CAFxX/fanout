package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8 }
type GenS2 struct { F0 bool; F1 uint64 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int; F1 bool }
type GenS5 struct { F0 uint8; F1 uint8 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 uint32; F1 string }
type GenS8 struct { F0 uint16; F1 int16; F2 int64 }
type GenS9 struct { F0 uint16; F1 bool; F2 uint64 }
type GenS10 struct { F0 uint16; F1 int }
type GenS11 struct { F0 uint16 }
func genFn0(p0 int32, p1 uint32) uint16 {
    switch "a" {
    case "":
        v6 := GenS1{F0: 134, F1: 98}
        _ = v6
    case "hello":
        fmt.Println(p0, p1)
    case "\n\t":
        var v7 int8 = 127
        _ = v7
    }
    v8 := []int{4294967295, 1}
    _ = p0
    _ = p1
    _ = v8
    return 255
}
func genFn1(p0 bool, p1 int8) uint16 {
    v9 := []int{}
    _ = len(v9)
    _ = p0
    _ = p1
    _ = v9
    return 319
}

func main() {
    v1 := []string{}
    var v2 bool = (-32768 < 1)
    v3 := []int{}
    var v4 uint = 1
    var v5 uint16 = 255
    if ((-798 > 867) || (2147483648 < 2147483647)) {
        var v10 int16 = 32767
        v11 := GenS2{F0: v2, F1: 1}
        _ = v10
        _ = v11
    } else {
        switch 4294967295 {
        case 32768:
            fmt.Println(v5, v2)
        case -9223372036854775808:
            v4 += v4
        case -1:
            var v12 uint64 = 18446744073709551615
            _ = v12
        default:
            if v2 {
                if ((((("" + "hello") + "a") != (("" + (("\n\t" + ("" + "a")) + (("hello" + "\n\t") + "hello"))) + ("foo bar" + ("hello" + ("" + "a"))))) || ((!(((!((v4 == v4) || (79 < 100))) && true) && (127 >= 100))) && (!(!v2)))) || (3.14159 != (-2.25 + float64(float32(float64(0.0)))))) {
                    v4 += v4
                    var v13 int8 = -16
                    _ = v13
                } else {
                    for fi14 := 0; fi14 < 2; fi14++ {
                        for fi15 := 0; fi15 < 20; fi15++ {
                            fmt.Println(fi14, v2, fi14)
                            v4 += (^(v4))
                            fmt.Println(v2)
                            _ = fi15
                        }
                        fmt.Println(v5, v5)
                        for fi16 := 0; fi16 < 12; fi16++ {
                            switch 624 {
                            case -9223372036854775808:
                                var v17 bool = true
                                _ = v17
                            case 139:
                                var v18 string = "foo bar"
                                _ = v18
                            case -269:
                                v19 := GenS2{F0: v2, F1: 18446744073709551615}
                                _ = v19
                            }
                            if v2 {
                                switch "\n\t" {
                                case "foo bar":
                                    var v20 int = fi16
                                    _ = v20
                                case "a":
                                    if v2 {
                                        fmt.Println(v4, fi16)
                                        var v21 float32 = ((float32(float64(1.5)) / -2.25) - (0.0 + 0.1))
                                        if v2 {
                                            v4 += (^(v4))
                                            switch (("a" + ("" + "hello")) + ("hello" + "")) {
                                            case "":
                                                v21 -= 0.1
                                            case "a":
                                                switch fi14 {
                                                case 128:
                                                    v22 := GenS1{F0: 127, F1: 2}
                                                    _ = v22
                                                case 82:
                                                    var v23 uint32 = 718
                                                    _ = v23
                                                case 2147483647:
                                                    switch 127 {
                                                    case 738:
                                                        fmt.Println(v21, v2, fi14)
                                                    case 4294967295:
                                                        v24 := GenS1{F0: uint8(v4), F1: 42}
                                                        _ = v24
                                                    }
                                                default:
                                                    var v25 int8 = -1
                                                    _ = v25
                                                }
                                            case "hello":
                                                v26 := GenS2{F0: (2 < 18446744073709551615), F1: 1}
                                                _ = v26
                                            default:
                                                var v27 uint16 = v5
                                                _ = v27
                                            }
                                        }
                                        _ = v21
                                    }
                                    switch fi14 {
                                    case -1:
                                        var v28 uint16 = genFn1(false, 62)
                                        _ = v28
                                    }
                                    if (!v2) {
                                        v3 = append(v3, 375)
                                        v2 = false
                                    } else {
                                        v4 = 18446744073709551615
                                    }
                                    v5 = (^(v5))
                                }
                                if v2 {
                                    v4 = (v4 - 18446744073709551615)
                                    var v29 float32 = ((1.5 - (1e10 / -2.25)) * 3.14159)
                                    switch (^(fi14)) {
                                    case -2147483648:
                                        if (!true) {
                                            switch "hello" {
                                            case "\n\t":
                                                switch "a" {
                                                case "a":
                                                    var v30 uint64 = 1
                                                    _ = v30
                                                case "hello":
                                                    var v31 string = "hello"
                                                    _ = v31
                                                case "":
                                                    v32 := GenS3{F0: v2}
                                                    _ = v32
                                                }
                                            case "foo bar":
                                                v2 = v2
                                                v4 -= v4
                                                switch fi16 {
                                                case -957:
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                case 127:
                                                    switch -1 {
                                                    case -881:
                                                        v33 := GenS3{F0: true}
                                                        _ = v33
                                                    case 577:
                                                        var v34 uint32 = 2147483647
                                                        _ = v34
                                                    }
                                                case 4294967295:
                                                    var v35 int16 = 744
                                                    _ = v35
                                                }
                                            default:
                                                v5 -= (v5 | v5)
                                            }
                                            var v36 uint32 = 42
                                            var v37 uint16 = v5
                                            switch "foo bar" {
                                            case "foo bar":
                                                v37 += v5
                                                var v38 int64 = 330
                                                var v39 float32 = ((((v29 * v29) / v29) / v29) - v29)
                                                _ = v38
                                                _ = v39
                                            case "":
                                                var v40 float64 = 0.0
                                                _ = v40
                                            case "\n\t":
                                                switch (("" + "hello") + ("a" + ("a" + (("a" + "hello") + "")))) {
                                                case "\n\t":
                                                    v36 = 852
                                                }
                                            default:
                                                if (!true) {
                                                    var v41 int16 = 32767
                                                    _ = v41
                                                }
                                            }
                                            if v2 {
                                                if (653 != 0) {
                                                    if v2 {
                                                        if false {
                                                            if len(v3) >= 2 {
                                                                v3 = v3[1:len(v3)]
                                                            }
                                                            switch fi14 {
                                                            case 32768:
                                                                v42 := GenS4{F0: (fi14 * 9223372036854775807), F1: (!true)}
                                                                _ = v42
                                                            default:
                                                                v43 := GenS4{F0: 4294967295, F1: false}
                                                                _ = v43
                                                            }
                                                            fmt.Println(v4, v5)
                                                        }
                                                        if v2 {
                                                            fmt.Println(v2, fi16)
                                                        } else {
                                                            var v44 int = fi14
                                                            _ = len(v3)
                                                            _ = v44
                                                        }
                                                        var v45 int64 = -9223372036854775808
                                                        _ = v45
                                                    }
                                                    v5 -= v5
                                                }
                                                switch "a" {
                                                case "hello":
                                                    if v2 {
                                                        if v2 {
                                                            v37 += 32767
                                                            v3 = append(v3, (fi16 >> ((v4) & 63)))
                                                        }
                                                    } else {
                                                        if (-1 <= -128) {
                                                            v46 := GenS2{F0: (2147483647 >= (-(fi14))), F1: 18446744073709551615}
                                                            _ = v46
                                                        } else {
                                                            var v47 float32 = float32(1.5)
                                                            var v48 uint16 = 989
                                                            _ = v47
                                                            _ = v48
                                                        }
                                                        if len(v1) > 0 {
                                                            v1[0] = ("hello" + "\n\t")
                                                        }
                                                    }
                                                case "a":
                                                    var v49 int64 = 42
                                                    _ = v49
                                                default:
                                                    v50 := GenS1{F0: 128, F1: 0}
                                                    _ = v50
                                                }
                                            } else {
                                                var v51 int32 = 302
                                                v52 := GenS2{F0: ((42 > uint64(v36)) && v2), F1: 709}
                                                _ = v51
                                                _ = v52
                                            }
                                            switch fi16 {
                                            case 235:
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                            case 1:
                                                fmt.Println(v4)
                                            default:
                                                _ = len(v3)
                                                switch "" {
                                                case "\n\t":
                                                    if (4294967295 >= 522) {
                                                        switch (fi14 % ((fi16) | 1)) {
                                                        case 0:
                                                            switch "hello" {
                                                            case "a":
                                                                switch "" {
                                                                case "\n\t":
                                                                    fmt.Println(v37)
                                                                default:
                                                                    fmt.Println(fi14, v2, fi14)
                                                                }
                                                            case "foo bar":
                                                                var v53 string = ("\n\t" + ("a" + "foo bar"))
                                                                _ = v53
                                                            case "":
                                                                fmt.Println(v5)
                                                            }
                                                        case -1:
                                                            fmt.Println(v29, v29)
                                                        }
                                                        if len(v3) >= 2 {
                                                            v3 = v3[1:len(v3)]
                                                        }
                                                    } else {
                                                        if len(v1) >= 2 {
                                                            v1 = v1[1:len(v1)]
                                                        }
                                                    }
                                                case "foo bar":
                                                    v54 := GenS5{F0: 199, F1: 77}
                                                    _ = v54
                                                case "":
                                                    var v55 int = -128
                                                    _ = v55
                                                default:
                                                    if len(v1) > 0 {
                                                        v1[0] = ("\n\t" + "hello")
                                                    }
                                                }
                                            }
                                            var v56 bool = ((100 == 0) || false)
                                            var v57 int64 = -566
                                            if v2 {
                                                switch (("" + "") + "\n\t") {
                                                case "hello":
                                                    fmt.Println(v2, fi14, v4)
                                                case "foo bar":
                                                    var v58 bool = true
                                                    if (v29 >= (v29 - ((0.0 - (0.1 + (v29 / ((1e10 + v29) * v29)))) + v29))) {
                                                        v59 := GenS6{F0: (fi14 | (fi16 % ((176) | 1)))}
                                                        _ = len(v3)
                                                        _ = v59
                                                    }
                                                    _ = v58
                                                case "a":
                                                    var v60 float32 = (v29 + 1.5)
                                                    _ = v60
                                                default:
                                                    switch (fi14 * (fi14 ^ fi14)) {
                                                    case 42:
                                                        v3 = append(v3, 948)
                                                    case -2147483648:
                                                        if (v57 > 9223372036854775807) {
                                                            v61 := GenS6{F0: (fi14 + fi14)}
                                                            _ = v61
                                                        }
                                                    case 127:
                                                        if (v36 < (v36 * (^(v36)))) {
                                                            v2 = v2
                                                        }
                                                    }
                                                    switch fi14 {
                                                    case 353:
                                                        if len(v1) >= 2 {
                                                            v1 = v1[1:len(v1)]
                                                        }
                                                    case -9223372036854775808:
                                                        if (0 >= 2) {
                                                            v62 := GenS2{F0: false, F1: 1}
                                                            if ((v4 >> ((1) & 63)) <= (v4 << ((v4) & 63))) {
                                                                switch ("hello" + ("hello" + (("" + (("" + ("a" + "a")) + "\n\t")) + ("\n\t" + "hello")))) {
                                                                case "":
                                                                    v5 += (^(v37))
                                                                default:
                                                                    v56 = ((0.1 / 1.0) < 1e10)
                                                                }
                                                                if true {
                                                                    if true {
                                                                        var v63 uint32 = v36
                                                                        _ = v63
                                                                    } else {
                                                                        v64 := GenS5{F0: 2, F1: 24}
                                                                        if (1 != 128) {
                                                                            var v65 float32 = (v29 / (0.1 * (v29 * float32(float64(float32(1.5))))))
                                                                            _ = v65
                                                                        }
                                                                        _ = v64
                                                                    }
                                                                    switch (("hello" + ("" + "")) + "a") {
                                                                    case "a":
                                                                        var v66 float64 = 0.0
                                                                        _ = v66
                                                                    case "":
                                                                        if (!false) {
                                                                            v67 := GenS1{F0: 255, F1: uint8(v5)}
                                                                            var v68 int8 = 48
                                                                            var v69 uint64 = 4294967295
                                                                            _ = v67
                                                                            _ = v68
                                                                            _ = v69
                                                                        }
                                                                    }
                                                                    v70 := GenS7{F0: v36, F1: ""}
                                                                    var v71 int = 32768
                                                                    var v72 uint64 = 42
                                                                    _ = v70
                                                                    _ = v71
                                                                    _ = v72
                                                                }
                                                                if v56 {
                                                                    if len(v1) > 0 {
                                                                        v1[0] = "\n\t"
                                                                    }
                                                                    switch "foo bar" {
                                                                    case "\n\t":
                                                                        fmt.Println(v5, v57)
                                                                    case "":
                                                                        var v73 string = "foo bar"
                                                                        _ = v73
                                                                    case "foo bar":
                                                                        var v74 int64 = v57
                                                                        _ = v74
                                                                    default:
                                                                        v29 = v29
                                                                        v75 := GenS7{F0: v36, F1: "\n\t"}
                                                                        _ = len(v1)
                                                                        _ = v75
                                                                    }
                                                                }
                                                            }
                                                            v4 -= (v4 ^ v4)
                                                            _ = v62
                                                        }
                                                    case 309:
                                                        var v76 uint64 = 677
                                                        _ = v76
                                                    default:
                                                        var v77 int64 = v57
                                                        _ = v77
                                                    }
                                                    var v78 int = (fi16 / (((-(fi16))) | 1))
                                                    _ = v78
                                                }
                                                v1 = append(v1, ("\n\t" + "hello"))
                                            }
                                            var v79 uint32 = v36
                                            _ = v36
                                            _ = v37
                                            _ = v56
                                            _ = v57
                                            _ = v79
                                        } else {
                                            if (-128 != 42) {
                                                v5 = (v5 % ((genFn0(-170, 611)) | 1))
                                                var v80 string = (("hello" + ((("\n\t" + ("a" + "foo bar")) + "\n\t") + "foo bar")) + "\n\t")
                                                var v81 int = fi16
                                                v82 := GenS8{F0: genFn1(false, -128), F1: 32767, F2: 1}
                                                _ = v80
                                                _ = v81
                                                _ = v82
                                            } else {
                                                var v83 int64 = -9223372036854775808
                                                if v2 {
                                                    fmt.Println(v2, v4)
                                                    fmt.Println(v4, fi14)
                                                    var v84 float64 = 0.1
                                                    _ = v84
                                                } else {
                                                    if len(v3) > 0 {
                                                        v3[0] = int(v5)
                                                    }
                                                }
                                                _ = v83
                                            }
                                            v85 := GenS9{F0: genFn0(0, 2147483647), F1: (341 == -1), F2: 42}
                                            _ = v85
                                        }
                                    default:
                                        fmt.Println(fi14, v4, v29)
                                    }
                                    _ = v29
                                } else {
                                    switch fi16 {
                                    case 9223372036854775807:
                                        v2 = (!v2)
                                    case -2147483648:
                                        v5 -= genFn0(-2147483648, 4294967295)
                                    case -475:
                                        v4 += v4
                                    }
                                }
                            } else {
                                v86 := GenS10{F0: (v5 | v5), F1: fi14}
                                _ = v86
                            }
                            _ = fi16
                        }
                        _ = fi14
                    }
                    v87 := GenS1{F0: 127, F1: 0}
                    _ = v87
                }
            }
        }
    }
    var v88 uint = (v4 & v4)
    v89 := GenS11{F0: 18}
    switch ("a" + ("" + "hello")) {
    case "":
        v4 = (v88 & 2)
    case "hello":
        v90 := GenS2{F0: v2, F1: 940}
        _ = v90
    case "a":
        if len(v3) > 0 {
            v3[0] = -2147483648
        }
    default:
        var v91 uint16 = v5
        _ = v91
    }
    v2 = (1 > 0)
    fmt.Println(v4)
    if (((false || v2) || (((((1e10 - (3.14159 * 1.5)) / (((float64(((0.1 / (-2.25 - float32((float64((3.14159 / -2.25)) - 1.5)))) - -2.25)) - 0.1) + 0.1) + -2.25)) * 0.0) - (3.14159 + float64(((0.0 + 1.5) - 1.5)))) > 1e10)) && true) {
        fmt.Println(v88, v4)
    }
    var v92 int64 = 42
    var v93 string = ""
    var v94 uint16 = uint16(v4)
    v95 := GenS9{F0: (v94 + 65535), F1: false, F2: 790}
    v2 = (0.1 != (1e10 - float64(float32(float64(0.0)))))
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v1 = append(v1, v93)
    if false {
        v96 := GenS11{F0: genFn0(42, 2147483648)}
        _ = v96
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v88:", v88)
    fmt.Println("v92:", v92)
    fmt.Println("v93:", v93)
    fmt.Println("v94:", v94)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v88
    _ = v89
    _ = v92
    _ = v93
    _ = v94
    _ = v95
}
