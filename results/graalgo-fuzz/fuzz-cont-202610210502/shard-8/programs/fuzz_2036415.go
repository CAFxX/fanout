package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 bool; F1 int8 }
type GenS5 struct { F0 bool; F1 uint16 }
type GenS6 struct { F0 int; F1 string }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 uint16; F1 int8; F2 uint }
type GenS10 struct { F0 bool }
type GenS11 struct { F0 int64 }
type GenS12 struct { F0 bool }
type GenS13 struct { F0 uint32; F1 bool }
type GenS14 struct { F0 int; F1 int64; F2 int8 }
type GenS15 struct { F0 string; F1 int64; F2 int64 }

func main() {
    var v1 uint32 = 597
    v2 := []int{}
    var v3 int8 = 1
    var v4 float64 = float64(-2.25)
    v5 := make(map[uint32]uint32)
    v6 := []string{("foo bar" + ((("" + ("\n\t" + "foo bar")) + "hello") + "hello"))}
    v4 = float64(1.5)
    var v7 float32 = 3.14159
    v7 = float32(v4)
    v3 += v3
    v7 += float32((v4 - v4))
    v8 := GenS1{F0: 0, F1: 128}
    v1 = 4294967295
    for fi9 := 0; fi9 < 16; fi9++ {
        v7 -= 1e10
        for fi10 := 0; fi10 < 19; fi10++ {
            switch fi10 {
            case 535:
                var v11 int8 = -70
                _ = v11
            case 1:
                switch (("" + "") + "\n\t") {
                case "a":
                    fmt.Println(fi10, fi10, fi9)
                case "foo bar":
                    v1 += (^(v1))
                default:
                    _ = len(v2)
                }
            case 4294967295:
                fmt.Println(v7, v1, fi10)
            default:
                v7 = (((float32(0.0) * v7) / 3.14159) - 0.0)
            }
            fmt.Println(fi9, fi9)
            _ = fi10
        }
        _ = fi9
    }
    if false {
        if false {
            var v12 int32 = -160
            _ = v12
        } else {
            v13 := GenS1{F0: 42, F1: 1}
            _ = v13
        }
    } else {
        var v14 float64 = (v4 - v4)
        _ = v14
    }
    switch ("\n\t" + ("hello" + ("foo bar" + "a"))) {
    case "hello":
        if ((^(v1)) != v1) {
            delete(v5, 0)
            for fi15 := 0; fi15 < 13; fi15++ {
                switch int(v3) {
                case 32767:
                    if (true || false) {
                        switch 521 {
                        case -956:
                            var v16 int = fi15
                            _ = v16
                        case 128:
                            for fi17 := 0; fi17 < 19; fi17++ {
                                v7 = v7
                                _ = fi17
                            }
                        }
                        v18 := GenS1{F0: 42, F1: 227}
                        _ = v18
                    } else {
                        switch "a" {
                        case "hello":
                            for fi19 := 0; fi19 < 1; fi19++ {
                                switch (^(fi15)) {
                                case 128:
                                    v5[(v1 / ((65535) | 1))] = 1
                                case 383:
                                    var v20 uint8 = 240
                                    _ = v20
                                case 634:
                                    if (!(float64((v7 / v7)) == 1e10)) {
                                        var v21 uint8 = 1
                                        _ = v21
                                    }
                                }
                                switch "foo bar" {
                                case "a":
                                    v7 = v7
                                }
                                fmt.Println(fi19, fi15, v7)
                                if (127 != (v1 << ((49) & 63))) {
                                    var v22 int8 = (v3 * v3)
                                    _ = v22
                                }
                                switch -1 {
                                case 32767:
                                    if false {
                                        if false {
                                            var v23 uint32 = v1
                                            var v24 bool = true
                                            var v25 bool = ((v3 << ((uint(v23)) & 63)) >= 1)
                                            _ = v23
                                            _ = v24
                                            _ = v25
                                        }
                                        switch "a" {
                                        case "":
                                            if true {
                                                if (v4 <= v4) {
                                                    if false {
                                                        v2 = append(v2, fi19)
                                                    } else {
                                                        v26 := GenS1{F0: 4294967295, F1: 2}
                                                        _ = v26
                                                    }
                                                    if (!true) {
                                                        var v27 float64 = 3.14159
                                                        switch fi19 {
                                                        case -2147483648:
                                                            v4 = (1.5 + v27)
                                                        }
                                                        if ((^(fi15)) >= fi15) {
                                                            v28 := GenS2{F0: 43}
                                                            _ = v28
                                                        }
                                                        v4 += v27
                                                        v1 = 2147483648
                                                        _ = v27
                                                    }
                                                }
                                            }
                                        default:
                                            if false {
                                                v29 := GenS1{F0: 4294967295, F1: 0}
                                                var v30 float32 = v7
                                                switch fi15 {
                                                case 127:
                                                    if true {
                                                        v7 -= (1.5 - v30)
                                                    }
                                                case 861:
                                                    fmt.Println(fi15, fi15)
                                                default:
                                                    var v31 int8 = 35
                                                    _ = v31
                                                }
                                                _ = len(v2)
                                                if true {
                                                    fmt.Println(fi15)
                                                } else {
                                                    var v32 uint8 = 255
                                                    if len(v6) > 0 {
                                                        v6[0] = "a"
                                                    }
                                                    var v33 int = (fi19 >> ((1) & 63))
                                                    if (v4 <= 0.0) {
                                                        var v34 string = "\n\t"
                                                        v35 := GenS3{F0: v32}
                                                        v36 := GenS4{F0: true, F1: (^(v3))}
                                                        _ = v34
                                                        _ = v35
                                                        _ = v36
                                                    } else {
                                                        if ((v3 ^ -128) != -128) {
                                                            v37 := GenS1{F0: 2, F1: (v32 << ((42) & 63))}
                                                            switch (v33 >> ((561) & 63)) {
                                                            case 1:
                                                                v38 := GenS1{F0: 1, F1: (^(v32))}
                                                                _ = v38
                                                            case 457:
                                                                switch 288 {
                                                                case 128:
                                                                    _ = len(v5)
                                                                case 134:
                                                                    switch "" {
                                                                    case "":
                                                                        var v39 float64 = v4
                                                                        _ = v39
                                                                    case "foo bar":
                                                                        v40 := GenS4{F0: (("a" == "\n\t") && (!((0.0 / v4) == v4))), F1: 0}
                                                                        _ = v40
                                                                    default:
                                                                        if true {
                                                                            var v41 int8 = v3
                                                                            _ = v41
                                                                        }
                                                                    }
                                                                    v42 := GenS5{F0: true, F1: 925}
                                                                    _ = v42
                                                                case 508:
                                                                    v43 := GenS5{F0: false, F1: 32767}
                                                                    _ = v43
                                                                default:
                                                                    if (442 > 18446744073709551615) {
                                                                        if (("a" + ("hello" + "a")) > ((("a" + (("hello" + "foo bar") + "a")) + ("" + "a")) + ((("\n\t" + ("" + (("\n\t" + "\n\t") + "foo bar"))) + "") + (((("foo bar" + ("\n\t" + "hello")) + "a") + (("foo bar" + "\n\t") + "foo bar")) + "hello")))) {
                                                                            switch fi19 {
                                                                            case -1:
                                                                                var v44 int32 = 2147483647
                                                                                _ = v44
                                                                            }
                                                                        }
                                                                    }
                                                                }
                                                            }
                                                            var v45 int64 = 0
                                                            _ = v37
                                                            _ = v45
                                                        } else {
                                                            var v46 string = "\n\t"
                                                            fmt.Println(v4)
                                                            _ = v46
                                                        }
                                                        _ = len(v2)
                                                    }
                                                    if true {
                                                        switch "\n\t" {
                                                        case "a":
                                                            fmt.Println(v30, v4, fi19)
                                                        }
                                                    } else {
                                                        v30 -= v30
                                                        switch 32768 {
                                                        case -668:
                                                            v47 := GenS4{F0: true, F1: -128}
                                                            _ = v47
                                                        case 0:
                                                            var v48 uint32 = (v1 | (v1 - v1))
                                                            var v49 int16 = int16(v3)
                                                            var v50 string = "a"
                                                            _ = v48
                                                            _ = v49
                                                            _ = v50
                                                        }
                                                        if true {
                                                            if len(v6) > 0 {
                                                                v6[0] = "\n\t"
                                                            }
                                                            if len(v2) > 0 {
                                                                v2[0] = 2147483647
                                                            }
                                                            _ = len(v5)
                                                        } else {
                                                            if (-2.25 == v30) {
                                                                v30 = 0.0
                                                            } else {
                                                                if false {
                                                                    v4 = v4
                                                                    v3 = -107
                                                                    v32 -= (^(v32))
                                                                } else {
                                                                    if len(v6) > 0 {
                                                                        v6[0] = (("a" + "\n\t") + ("foo bar" + ("foo bar" + ("a" + "hello"))))
                                                                    }
                                                                    if ("a" >= "foo bar") {
                                                                        var v51 int8 = int8(v1)
                                                                        _ = len(v5)
                                                                        v52 := GenS6{F0: (fi19 / ((0) | 1)), F1: ("hello" + "\n\t")}
                                                                        _ = v51
                                                                        _ = v52
                                                                    } else {
                                                                        if false {
                                                                            var v53 float32 = v7
                                                                            v53 -= -2.25
                                                                            _ = v53
                                                                        }
                                                                    }
                                                                    if false {
                                                                        var v54 int16 = 125
                                                                        _ = v54
                                                                    }
                                                                    if (v30 < ((1.5 + 1e10) * (1.5 + v7))) {
                                                                        v30 += 3.14159
                                                                        var v55 uint8 = v32
                                                                        v56 := GenS7{F0: (false && ((!true) || (523 == 754)))}
                                                                        _ = v55
                                                                        _ = v56
                                                                    }
                                                                }
                                                                fmt.Println(fi19, v1)
                                                            }
                                                        }
                                                    }
                                                    _ = v32
                                                    _ = v33
                                                }
                                                _ = v29
                                                _ = v30
                                            }
                                        }
                                    }
                                case 42:
                                    if len(v2) > 0 {
                                        v2[0] = (fi15 | fi15)
                                    }
                                default:
                                    if ("" == ("hello" + "hello")) {
                                        switch "" {
                                        case "a":
                                            if len(v2) > 0 {
                                                v2[0] = -648
                                            }
                                        }
                                        v3 = -128
                                        v3 = (v3 >> ((18446744073709551615) & 63))
                                    } else {
                                        v57 := GenS3{F0: 0}
                                        _ = v57
                                    }
                                }
                                var v58 int64 = 42
                                var v59 int8 = (v3 & v3)
                                if true {
                                    var v60 uint64 = 616
                                    if true {
                                        v7 = v7
                                        v61 := GenS8{F0: 417}
                                        _ = v61
                                    } else {
                                        v7 = (v7 + 3.14159)
                                        v5[v1] = 2
                                    }
                                    _ = len(v5)
                                    v62 := GenS1{F0: 0, F1: 1}
                                    if ("foo bar" > "") {
                                        fmt.Println(v59, v59, v3)
                                        _ = len(v6)
                                        var v63 int8 = v3
                                        _ = v63
                                    }
                                    _ = v60
                                    _ = v62
                                } else {
                                    v64 := GenS9{F0: 2, F1: v3, F2: 975}
                                    _ = v64
                                }
                                _ = fi19
                                _ = v58
                                _ = v59
                            }
                        case "foo bar":
                            if false {
                                _ = len(v5)
                                if true {
                                    if (18446744073709551615 > 0) {
                                        v3 += 89
                                        switch (fi15 + fi15) {
                                        case 128:
                                            var v65 uint16 = 32767
                                            _ = v65
                                        default:
                                            _ = len(v5)
                                        }
                                        v7 -= 1e10
                                    }
                                } else {
                                    var v66 float32 = ((float32(-2.25) / (1.5 - v7)) - 3.14159)
                                    var v67 int = (fi15 / (((fi15 & fi15)) | 1))
                                    for fi68 := 0; fi68 < 16; fi68++ {
                                        v66 += (0.0 * ((float32((((float64(v66) + v4) * v4) * v4)) * v7) * float32(1e10)))
                                        switch "hello" {
                                        case "hello":
                                            v69 := GenS5{F0: (v4 < v4), F1: 55}
                                            _ = v69
                                        case "foo bar":
                                            v70 := GenS10{F0: (true || (255 >= 127))}
                                            _ = v70
                                        case "":
                                            v4 += 0.1
                                        default:
                                            var v71 int32 = -63
                                            _ = v71
                                        }
                                        _ = fi68
                                    }
                                    _ = v66
                                    _ = v67
                                }
                            } else {
                                v4 = -2.25
                                fmt.Println(fi15, v1, v1)
                                for fi72 := 0; fi72 < 1; fi72++ {
                                    var v73 float32 = float32(v4)
                                    _ = len(v6)
                                    var v74 float32 = 1.5
                                    _ = fi72
                                    _ = v73
                                    _ = v74
                                }
                            }
                        }
                    }
                case 177:
                    v75 := GenS10{F0: (!(1 > 1))}
                    _ = v75
                }
                _ = fi15
            }
            var v76 float32 = v7
            _ = v76
        } else {
            for fi77 := 0; fi77 < 12; fi77++ {
                for fi78 := 0; fi78 < 2; fi78++ {
                    v79 := GenS5{F0: true, F1: 1}
                    switch "foo bar" {
                    case "":
                        var v80 uint64 = 0
                        v6 = append(v6, "foo bar")
                        _ = v80
                    case "\n\t":
                        v4 = 0.0
                    case "foo bar":
                        var v81 int = -9223372036854775808
                        _ = v81
                    default:
                        fmt.Println(fi77, fi77, fi77)
                    }
                    var v82 string = "\n\t"
                    _ = fi78
                    _ = v79
                    _ = v82
                }
                v83 := GenS11{F0: -879}
                v7 -= float32((v4 / v4))
                _ = fi77
                _ = v83
            }
        }
    case "foo bar":
        if true {
            if (42 == 224) {
                fmt.Println(v1, v7, v1)
            } else {
                if false {
                    var v84 uint32 = v1
                    if false {
                        for fi85 := 0; fi85 < 2; fi85++ {
                            var v86 int64 = -1
                            for fi87 := 0; fi87 < 3; fi87++ {
                                if false {
                                    v1 = v84
                                    switch -9223372036854775808 {
                                    case 769:
                                        if (2 != 128) {
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                            if (v7 != v7) {
                                                var v88 bool = false
                                                if v88 {
                                                    var v89 uint8 = 2
                                                    v90 := GenS12{F0: false}
                                                    var v91 uint8 = (^(v89))
                                                    _ = v89
                                                    _ = v90
                                                    _ = v91
                                                } else {
                                                    switch (^(fi87)) {
                                                    case 1:
                                                        v92 := v5[(v1 * v1)]
                                                        _ = v92
                                                    default:
                                                        v4 -= 0.0
                                                    }
                                                    var v93 bool = (607 > 882)
                                                    _ = v93
                                                }
                                                var v94 uint32 = (v84 & 42)
                                                _ = v88
                                                _ = v94
                                            }
                                        } else {
                                            delete(v5, 0)
                                        }
                                    default:
                                        v4 = (v4 / (0.0 + v4))
                                    }
                                    fmt.Println(v84)
                                }
                                _ = fi87
                            }
                            if (-642 != (fi85 << ((614) & 63))) {
                                for fi95 := 0; fi95 < 17; fi95++ {
                                    var v96 int32 = 0
                                    v86 = (v86 / ((int64(v1)) | 1))
                                    if ((v4 + -2.25) > (0.0 + 0.0)) {
                                        v97 := GenS3{F0: 29}
                                        _ = v97
                                    }
                                    fmt.Println(v3, v7)
                                    _ = fi95
                                    _ = v96
                                }
                                v98 := GenS1{F0: 42, F1: uint8(v84)}
                                _ = v98
                            } else {
                                v99 := GenS13{F0: 0, F1: false}
                                _ = v99
                            }
                            _ = fi85
                            _ = v86
                        }
                        var v100 uint16 = 255
                        _ = v100
                    } else {
                        v101 := GenS14{F0: 1, F1: 961, F2: (^(v3))}
                        _ = v101
                    }
                    _ = v84
                } else {
                    var v102 int16 = -699
                    v6 = append(v6, ("hello" + ""))
                    _ = v102
                }
            }
            _ = len(v2)
        }
    case "a":
        v103 := GenS15{F0: ("\n\t" + (("a" + ("" + "")) + "")), F1: -1, F2: 1}
        _ = v103
    default:
        if true {
            v4 = (v4 + v4)
        }
    }
    v5[(v1 & 2)] = 42
    v104 := GenS9{F0: 42, F1: (-(v3)), F2: 42}
    v105 := GenS9{F0: 2, F1: v3, F2: 65535}
    var v106 bool = true
    for fi107 := 0; fi107 < 7; fi107++ {
        fmt.Println(v106)
        var v108 uint16 = 0
        _ = fi107
        _ = v108
    }
    var v109 uint16 = uint16(v1)
    fmt.Println("v1:", v1)
    fmt.Println("v106:", v106)
    fmt.Println("v109:", v109)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v104
    _ = v105
    _ = v106
    _ = v109
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
