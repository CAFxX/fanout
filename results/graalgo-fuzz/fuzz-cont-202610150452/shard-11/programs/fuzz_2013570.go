package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 int64 }
type GenS2 struct { F0 int64; F1 bool }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int; F1 bool; F2 uint }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 int32; F1 int8; F2 uint8 }
type GenS8 struct { F0 uint8; F1 string; F2 uint8 }
type GenS9 struct { F0 bool; F1 uint64; F2 uint }
type GenS10 struct { F0 bool; F1 uint16; F2 string }
type GenS11 struct { F0 uint32 }
type GenS12 struct { F0 string; F1 int32; F2 int }
type GenS13 struct { F0 int8; F1 uint32 }
func genFn0(p0 string) uint {
    v5 := []int{0}
    if (0.0 > 0.1) {
        if (!false) {
            p0 = (((p0 + p0) + p0) + ("hello" + p0))
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        }
        switch ((p0 + (p0 + ("hello" + p0))) + p0) {
        case "\n\t":
            v6 := GenS1{F0: 255, F1: 255, F2: 42}
            _ = v6
        default:
            var v7 int16 = -1
            _ = v7
        }
        var v8 int32 = -1
        _ = v8
    }
    _ = p0
    _ = v5
    return 862
}
func genFn1(p0 uint64, p1 int) uint32 {
    if false {
        p1 = (p1 & (^(p1)))
    }
    for fi9 := 0; fi9 < 4; fi9++ {
        v10 := GenS2{F0: -513, F1: (!true)}
        _ = fi9
        _ = v10
    }
    v11 := []int{p1, (p1 << ((2) & 63))}
    if (127 >= 127) {
        if false {
            var v12 float64 = float64((1.5 - 0.1))
            var v13 int = p1
            _ = v12
            _ = v13
        }
    }
    _ = p0
    _ = p1
    _ = v11
    return 483
}

func main() {
    var v1 int16 = -882
    v2 := []string{(("hello" + "foo bar") + "foo bar"), ("foo bar" + ("a" + "")), "hello", "hello"}
    v3 := make(map[string]string)
    v4 := []int{4294967295, 42, -9223372036854775808}
    if (0.1 > (-2.25 - float64(1e10))) {
        for fi14 := 0; fi14 < 6; fi14++ {
            var v15 int = fi14
            var v16 float64 = 1.5
            for fi17 := 0; fi17 < 1; fi17++ {
                if false {
                    var v18 uint32 = 4294967295
                    v1 = 1
                    var v19 int64 = -611
                    _ = v18
                    _ = v19
                } else {
                    if true {
                        switch (-(fi14)) {
                        case 42:
                            v15 -= 42
                        }
                        fmt.Println(v15)
                        var v20 int8 = -128
                        _ = v20
                    }
                    var v21 float32 = (-2.25 / 1.0)
                    v4 = append(v4, (fi14 - 152))
                    v22 := GenS2{F0: -383, F1: (false && (!true))}
                    var v23 uint32 = 505
                    switch (("" + "foo bar") + "\n\t") {
                    case "foo bar":
                        switch "hello" {
                        case "":
                            v24 := GenS3{F0: v1}
                            _ = v24
                        case "\n\t":
                            switch fi14 {
                            case 42:
                                switch "a" {
                                case "hello":
                                    fmt.Println(v23, v1)
                                }
                            case 235:
                                _ = len(v2)
                            case -844:
                                var v25 uint = 255
                                _ = v25
                            default:
                                var v26 float64 = -2.25
                                _ = v26
                            }
                        case "hello":
                            var v27 uint32 = 42
                            _ = v27
                        }
                    }
                    _ = v21
                    _ = v22
                    _ = v23
                }
                if true {
                    var v28 uint8 = 255
                    _ = v28
                } else {
                    if (v1 == (v1 + (v1 * (v1 << ((255) & 63))))) {
                        var v29 int8 = 127
                        v29 = 0
                        switch v15 {
                        case -711:
                            v30 := GenS2{F0: 9223372036854775807, F1: false}
                            _ = v30
                        }
                        _ = v29
                    } else {
                        if len(v4) > 0 {
                            v4[0] = 32768
                        }
                        if (761 < 1) {
                            switch ("\n\t" + "foo bar") {
                            case "\n\t":
                                if true {
                                    var v31 float32 = float32(((0.1 / 1.5) * v16))
                                    if true {
                                        if (52 == 1) {
                                            var v32 float32 = v31
                                            if (382 != 65535) {
                                                fmt.Println(fi14, v32)
                                            }
                                            _ = v32
                                        } else {
                                            v16 = (float64(v31) * v16)
                                            var v33 float32 = 3.14159
                                            _ = v33
                                        }
                                        v16 += (v16 - float64((((1e10 + v31) / v31) + ((v31 + 1.5) + (-2.25 + 3.14159)))))
                                    } else {
                                        var v34 int8 = 42
                                        _ = len(v3)
                                        switch "hello" {
                                        case "\n\t":
                                            var v35 int8 = -128
                                            _ = v35
                                        default:
                                            if false {
                                                _ = len(v4)
                                                fmt.Println(v34, fi17, v16)
                                                fmt.Println(v1, v31)
                                            } else {
                                                switch (fi14 - (fi17 - -841)) {
                                                case -9223372036854775808:
                                                    if true {
                                                        v16 = (((float64(v31) - v16) / 3.14159) * float64(v31))
                                                    }
                                                case 32767:
                                                    v36 := GenS4{F0: 108}
                                                    _ = v36
                                                case 127:
                                                    var v37 int64 = 1
                                                    _ = v37
                                                default:
                                                    v31 = v31
                                                }
                                            }
                                        }
                                        if (42 > 2) {
                                            if true {
                                                var v38 uint16 = 42
                                                _ = v38
                                            } else {
                                                v34 -= -74
                                                var v39 uint = 0
                                                v40 := v3["a"]
                                                v31 = ((v31 * ((v31 * (v31 / 1.0)) - 1e10)) + float32(float64((0.1 / (3.14159 / float32(1e10))))))
                                                _ = v39
                                                _ = v40
                                            }
                                            var v41 string = ("hello" + "hello")
                                            var v42 int64 = -1
                                            if (!false) {
                                                if (fi14 > fi14) {
                                                    var v43 float32 = (v31 + v31)
                                                    _ = v43
                                                }
                                                switch "a" {
                                                case "":
                                                    fmt.Println(v15, v15)
                                                case "\n\t":
                                                    var v44 int = 9223372036854775807
                                                    _ = v44
                                                case "hello":
                                                    var v45 int32 = -2147483648
                                                    _ = v45
                                                }
                                            } else {
                                                if true {
                                                    if true {
                                                        var v46 int8 = v34
                                                        var v47 uint16 = 32767
                                                        var v48 float64 = float64(-2.25)
                                                        _ = v46
                                                        _ = v47
                                                        _ = v48
                                                    } else {
                                                        if true {
                                                            v16 = (v16 + float64(float32(v16)))
                                                        }
                                                    }
                                                }
                                                var v49 int32 = -450
                                                _ = v49
                                            }
                                            v50 := GenS4{F0: 67}
                                            if (1 != 18446744073709551615) {
                                                if len(v4) > 0 {
                                                    v4[0] = 1
                                                }
                                                v51 := GenS5{F0: (fi17 << ((genFn0("")) & 63)), F1: (4294967295 > 65535), F2: 0}
                                                switch (v41 + v41) {
                                                case "foo bar":
                                                    if (42 < -946) {
                                                        v16 -= (v16 + 0.1)
                                                    } else {
                                                        if (!(true && (v42 >= v42))) {
                                                            var v52 string = v41
                                                            _ = v52
                                                        }
                                                    }
                                                case "\n\t":
                                                    if len(v4) > 0 {
                                                        v4[0] = v15
                                                    }
                                                }
                                                _ = v51
                                            }
                                            _ = v41
                                            _ = v42
                                            _ = v50
                                        }
                                        _ = v34
                                    }
                                    _ = v31
                                } else {
                                    v16 = (v16 - 1e10)
                                    var v53 bool = true
                                    v54 := GenS6{F0: 0}
                                    var v55 float32 = 3.14159
                                    _ = v53
                                    _ = v54
                                    _ = v55
                                }
                                fmt.Println(v16, v15, v15)
                                var v56 int = (fi17 ^ v15)
                                _ = v56
                            case "foo bar":
                                if (!((((float32((((v16 / v16) - v16) + 1e10)) * 1e10) - 1e10) == (float32((((v16 - v16) * ((v16 * v16) + (v16 * v16))) / ((v16 - -2.25) + (v16 / float64(1e10))))) * 1.5)) || true)) {
                                    if ((1 < 65535) || ((fi14 & fi17) < (fi17 - 230))) {
                                        fmt.Println(v16, v1)
                                        switch (fi14 | fi17) {
                                        case -93:
                                            if true {
                                                fmt.Println(v1, fi14, fi14)
                                                var v57 float32 = 3.14159
                                                fmt.Println(v15, v16)
                                                _ = v57
                                            } else {
                                                var v58 float64 = float64(-2.25)
                                                _ = v58
                                            }
                                        default:
                                            if false {
                                                if len(v2) > 0 {
                                                    v2[0] = "foo bar"
                                                }
                                            }
                                        }
                                        v15 -= (fi17 / (((fi14 | 734)) | 1))
                                    }
                                    v15 += fi17
                                } else {
                                    fmt.Println(v16, v16, fi14)
                                }
                            case "a":
                                _ = len(v3)
                                var v59 int = v15
                                var v60 int64 = 1
                                _ = v59
                                _ = v60
                            }
                            var v61 uint = 0
                            _ = v61
                        } else {
                            fmt.Println(fi14, v1)
                            v16 = v16
                        }
                        if true {
                            v62 := GenS3{F0: 42}
                            var v63 uint8 = 128
                            v1 = 1000
                            _ = v62
                            _ = v63
                        }
                        var v64 uint32 = 143
                        if (243 == 0) {
                            switch fi17 {
                            case 1:
                                if false {
                                    v65 := GenS7{F0: -1, F1: -128, F2: 1}
                                    v4 = append(v4, -35)
                                    v16 -= (v16 / float64(-2.25))
                                    _ = v65
                                } else {
                                    var v66 float64 = (3.14159 * -2.25)
                                    v66 = 0.0
                                    if false {
                                        _ = len(v3)
                                        v1 -= v1
                                        switch "foo bar" {
                                        case "a":
                                            v67 := GenS5{F0: (v15 | (fi14 & -2147483648)), F1: true, F2: genFn0("\n\t")}
                                            _ = v67
                                        case "hello":
                                            v64 += 107
                                        case "\n\t":
                                            v64 -= (v64 | v64)
                                        }
                                    } else {
                                        v68 := GenS5{F0: fi14, F1: (-36 != 100), F2: 494}
                                        _ = v68
                                    }
                                    _ = v66
                                }
                            case -526:
                                switch "" {
                                case "foo bar":
                                    var v69 float64 = 1.5
                                    _ = v69
                                default:
                                    if (v15 > (v15 + v15)) {
                                        fmt.Println(fi14)
                                        v70 := GenS8{F0: 127, F1: (("a" + ("hello" + "\n\t")) + "hello"), F2: 42}
                                        var v71 uint64 = 740
                                        var v72 string = "a"
                                        _ = v70
                                        _ = v71
                                        _ = v72
                                    }
                                }
                            case -1:
                                var v73 bool = (v64 >= 2147483648)
                                _ = v73
                            default:
                                var v74 uint = 718
                                v75 := v3["foo bar"]
                                _ = v74
                                _ = v75
                            }
                            v76 := GenS1{F0: 4294967295, F1: 0, F2: 9223372036854775807}
                            v77 := GenS9{F0: false, F1: 101, F2: 284}
                            _ = v76
                            _ = v77
                        }
                        _ = v64
                    }
                }
                v78 := GenS10{F0: (v1 > v1), F1: 468, F2: ""}
                v79 := GenS2{F0: -245, F1: true}
                _ = fi17
                _ = v78
                _ = v79
            }
            _ = fi14
            _ = v15
            _ = v16
        }
    }
    var v80 uint64 = 0
    if false {
        switch 127 {
        case -894:
            if (false && (71 < 0)) {
                var v81 uint16 = 328
                _ = v81
            } else {
                if false {
                    v80 -= (v80 | 845)
                } else {
                    var v82 uint32 = 1
                    v83 := v3[("foo bar" + "hello")]
                    v83 = (v83 + ("" + v83))
                    _ = v82
                    _ = v83
                }
                var v84 float32 = 0.1
                _ = v84
            }
        case -767:
            for fi85 := 0; fi85 < 11; fi85++ {
                v80 = (v80 & 735)
                var v86 int64 = 1
                _ = fi85
                _ = v86
            }
        case 32768:
            var v87 float64 = 3.14159
            _ = v87
        }
        v2 = append(v2, "\n\t")
        v80 = v80
    } else {
        v1 -= v1
        switch 32768 {
        case 228:
            v88 := GenS11{F0: genFn1(301, -5)}
            _ = v88
        case 1:
            v1 += (v1 | -653)
        }
    }
    fmt.Println(v1, v80)
    v1 = 757
    v89 := GenS12{F0: "foo bar", F1: 2147483647, F2: 2147483647}
    switch 1 {
    case 2147483647:
        v90 := GenS13{F0: 100, F1: 2147483648}
        _ = v90
    case 127:
        switch "hello" {
        case "hello":
            delete(v3, (("\n\t" + "\n\t") + (("\n\t" + ((("" + ("hello" + "\n\t")) + "\n\t") + (("foo bar" + "") + "\n\t"))) + "hello")))
        case "":
            _ = len(v3)
        case "a":
            delete(v3, ("" + ("a" + "a")))
        }
    default:
        if len(v2) > 0 {
            v2[0] = "\n\t"
        }
    }
    v80 -= 0
    v91 := GenS7{F0: -727, F1: 42, F2: 40}
    v80 -= (v80 / (((v80 + v80)) | 1))
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if true {
        v92 := GenS13{F0: -123, F1: 2}
        _ = v92
    }
    for fi93 := 0; fi93 < 5; fi93++ {
        var v94 float32 = float32(((0.1 - (1.5 / ((0.0 + 1.5) / 1.5))) / 1.5))
        for fi95 := 0; fi95 < 6; fi95++ {
            if ((!((fi95 & fi95) == (fi95 ^ -1))) && ((2 >= 1) || true)) {
                fmt.Println(v80)
            } else {
                var v96 uint8 = 128
                var v97 uint8 = (v96 >> ((761) & 63))
                var v98 int32 = -2147483648
                _ = v96
                _ = v97
                _ = v98
            }
            _ = fi95
        }
        _ = fi93
        _ = v94
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v80:", v80)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v80
    _ = v89
    _ = v91
}
