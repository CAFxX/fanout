package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32; F2 int32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int64; F1 int32 }
type GenS4 struct { F0 int8; F1 int64 }
type GenS5 struct { F0 int64; F1 uint16; F2 bool }
type GenS6 struct { F0 int; F1 bool; F2 int8 }
type GenS7 struct { F0 bool; F1 int64 }
type GenS8 struct { F0 uint64; F1 bool }
type GenS9 struct { F0 uint16; F1 int; F2 uint32 }
type GenS10 struct { F0 uint64 }
type GenS11 struct { F0 int16; F1 uint32; F2 uint16 }

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[int]bool)
    v3 := []uint32{42, 2147483648}
    v4 := make(map[uint32]bool)
    v5 := v2[32767]
    if v5 {
        var v6 int64 = 9223372036854775807
        for fi7 := 0; fi7 < 7; fi7++ {
            if v5 {
                if ((2 > 4294967295) || (127 > 56)) {
                    v8 := GenS1{F0: 564, F1: 2147483647, F2: -2147483648}
                    _ = v8
                } else {
                    v5 = (228 >= 255)
                    v9 := GenS1{F0: 1, F1: 65535, F2: 2147483647}
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    _ = v9
                }
                v10 := GenS1{F0: 893, F1: 2147483647, F2: -2147483648}
                _ = v10
            } else {
                var v11 uint16 = 65535
                _ = v11
            }
            if false {
                var v12 uint = 255
                _ = v12
            }
            _ = fi7
        }
        _ = v6
    }
    if ("\n\t" == "\n\t") {
        if v5 {
            fmt.Println(v5)
            var v13 uint = 971
            var v14 float64 = 3.14159
            _ = v13
            _ = v14
        }
        switch 673 {
        case 295:
            v5 = (284 == 1)
        case 553:
            var v15 int64 = 9223372036854775807
            _ = v15
        case -2147483648:
            v5 = v5
        default:
            if ((float32(3.14159) / 0.1) == 1e10) {
                for fi16 := 0; fi16 < 13; fi16++ {
                    if v5 {
                        for fi17 := 0; fi17 < 16; fi17++ {
                            v18 := GenS1{F0: 0, F1: 2147483648, F2: 0}
                            _ = fi17
                            _ = v18
                        }
                    }
                    v19 := GenS2{F0: 42}
                    for fi20 := 0; fi20 < 10; fi20++ {
                        v5 = false
                        _ = fi20
                    }
                    _ = fi16
                    _ = v19
                }
                if (824 > 2) {
                    v5 = true
                    switch "hello" {
                    case "a":
                        for fi21 := 0; fi21 < 7; fi21++ {
                            v22 := GenS2{F0: 0}
                            _ = fi21
                            _ = v22
                        }
                    case "\n\t":
                        _ = len(v3)
                    case "":
                        if false {
                            for fi23 := 0; fi23 < 14; fi23++ {
                                v3 = append(v3, 65535)
                                _ = fi23
                            }
                            if v5 {
                                if (127 <= 100) {
                                    fmt.Println(v5)
                                    var v24 uint16 = 42
                                    _ = v24
                                } else {
                                    v5 = false
                                }
                                v25 := GenS1{F0: 839, F1: 1, F2: 0}
                                _ = v25
                            } else {
                                if v5 {
                                    var v26 uint = 42
                                    v27 := GenS1{F0: 4294967295, F1: 2, F2: -1}
                                    _ = v26
                                    _ = v27
                                }
                                for fi28 := 0; fi28 < 17; fi28++ {
                                    for fi29 := 0; fi29 < 9; fi29++ {
                                        var v30 float64 = float64(float32(-2.25))
                                        _ = len(v3)
                                        _ = fi29
                                        _ = v30
                                    }
                                    _ = fi28
                                }
                                v31 := GenS3{F0: 240, F1: -487}
                                _ = v31
                            }
                            var v32 int64 = -16
                            _ = v32
                        } else {
                            fmt.Println(v5)
                        }
                    default:
                        if len(v3) > 0 {
                            v3[0] = 0
                        }
                    }
                }
                var v33 uint = 0
                _ = v33
            }
        }
        v34 := GenS1{F0: 0, F1: 582, F2: -291}
        _ = v34
    } else {
        v35 := GenS4{F0: 85, F1: 1}
        var v36 int = 32768
        if (255 >= 4294967295) {
            v37 := GenS1{F0: 348, F1: 725, F2: 0}
            var v38 int = 2147483647
            for fi39 := 0; fi39 < 12; fi39++ {
                v40 := GenS2{F0: 0}
                _ = fi39
                _ = v40
            }
            _ = v37
            _ = v38
        } else {
            v36 = v36
        }
        _ = v35
        _ = v36
    }
    v5 = v5
    v5 = (true && (419 <= -661))
    for fi41 := 0; fi41 < 2; fi41++ {
        delete(v4, 2147483648)
        if (-38 != 100) {
            v5 = true
            v42 := GenS3{F0: 9223372036854775807, F1: 244}
            for fi43 := 0; fi43 < 10; fi43++ {
                v44 := GenS5{F0: 1, F1: 2, F2: false}
                _ = fi43
                _ = v44
            }
            _ = v42
        } else {
            if false {
                fmt.Println(v5, fi41)
            } else {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                v5 = (65535 != 42)
            }
            if (419 <= 0) {
                v5 = true
                if (!(("a" + "a") >= ("\n\t" + (("hello" + "hello") + ("" + ("foo bar" + ("\n\t" + (("hello" + "a") + ("" + "foo bar"))))))))) {
                    if v5 {
                        v45 := GenS5{F0: 9223372036854775807, F1: 32767, F2: true}
                        v3 = append(v3, 42)
                        _ = v45
                    } else {
                        for fi46 := 0; fi46 < 17; fi46++ {
                            v5 = ((v5 || (fi46 <= (fi46 ^ 32767))) && v5)
                            if (fi41 == (fi41 + (fi41 % (((fi46 + fi41)) | 1)))) {
                                if v5 {
                                    var v47 uint64 = 1
                                    var v48 float64 = 1e10
                                    _ = len(v1)
                                    _ = v47
                                    _ = v48
                                } else {
                                    v5 = v5
                                    switch fi46 {
                                    case 128:
                                        var v49 int64 = -308
                                        _ = v49
                                    }
                                }
                            }
                            _ = fi46
                        }
                        v50 := GenS3{F0: 0, F1: 0}
                        _ = v50
                    }
                    if len(v3) > 0 {
                        v3[0] = 1
                    }
                    switch "" {
                    case "hello":
                        var v51 float64 = (float64(((0.1 + float32((3.14159 / ((0.0 / 1.5) - 1e10)))) * ((((3.14159 / (1.5 + 0.1)) - 3.14159) / 3.14159) - 3.14159))) - 1.5)
                        _ = v51
                    case "\n\t":
                        for fi52 := 0; fi52 < 3; fi52++ {
                            v5 = (42 >= 335)
                            _ = fi52
                        }
                    case "foo bar":
                        if (!true) {
                            if ((4294967295 <= 914) || ("hello" != "foo bar")) {
                                for fi53 := 0; fi53 < 19; fi53++ {
                                    var v54 uint = 65535
                                    var v55 string = ("hello" + "\n\t")
                                    v5 = v5
                                    if (42 < -565) {
                                        var v56 uint32 = 2147483648
                                        var v57 float32 = 1e10
                                        fmt.Println(fi53, v56)
                                        v56 -= v56
                                        _ = len(v3)
                                        _ = v56
                                        _ = v57
                                    } else {
                                        v58 := GenS6{F0: (fi41 / ((fi41) | 1)), F1: true, F2: 100}
                                        _ = v58
                                    }
                                    _ = fi53
                                    _ = v54
                                    _ = v55
                                }
                                if (0 < 0) {
                                    var v59 bool = (!v5)
                                    var v60 uint16 = 42
                                    if false {
                                        v1[57] = "hello"
                                        var v61 bool = v5
                                        v59 = v61
                                        _ = v61
                                    }
                                    _ = v59
                                    _ = v60
                                } else {
                                    fmt.Println(fi41)
                                    fmt.Println(fi41)
                                }
                                fmt.Println(v5, fi41)
                            } else {
                                delete(v2, (fi41 | 1))
                                v5 = true
                                var v62 int8 = 0
                                _ = v62
                            }
                        } else {
                            v5 = v5
                            if len(v3) > 0 {
                                v3[0] = 1
                            }
                        }
                    }
                } else {
                    for fi63 := 0; fi63 < 15; fi63++ {
                        switch fi63 {
                        case -2147483648:
                            v64 := GenS6{F0: fi41, F1: v5, F2: -128}
                            _ = v64
                        case 32768:
                            var v65 string = "\n\t"
                            _ = v65
                        case -276:
                            switch fi63 {
                            case 682:
                                v5 = v5
                            case -2147483648:
                                if (42 <= 1000) {
                                    v66 := GenS7{F0: true, F1: 1}
                                    if (!true) {
                                        var v67 uint8 = 1
                                        if v5 {
                                            v68 := GenS8{F0: 162, F1: (3.14159 >= ((1e10 / 1.0) / (0.0 + float32(-2.25))))}
                                            var v69 uint64 = 473
                                            switch "hello" {
                                            case "\n\t":
                                                v70 := GenS3{F0: 286, F1: 1}
                                                switch ("a" + ("a" + "\n\t")) {
                                                case "foo bar":
                                                    if false {
                                                        var v71 uint32 = 293
                                                        if (!v5) {
                                                            if false {
                                                                v72 := GenS5{F0: -9223372036854775808, F1: 255, F2: (("\n\t" + "a") > "foo bar")}
                                                                fmt.Println(fi41, v5)
                                                                if ("hello" < "foo bar") {
                                                                    v3 = append(v3, (v71 >> ((304) & 63)))
                                                                    var v73 string = ("" + "foo bar")
                                                                    v69 = v69
                                                                    _ = v73
                                                                }
                                                                _ = v72
                                                            } else {
                                                                v74 := GenS2{F0: v67}
                                                                var v75 uint = 96
                                                                _ = v74
                                                                _ = v75
                                                            }
                                                            var v76 int = 32767
                                                            var v77 float64 = 1e10
                                                            _ = v76
                                                            _ = v77
                                                        } else {
                                                            if (-454 != -2147483648) {
                                                                fmt.Println(v5, fi63, fi41)
                                                                if false {
                                                                    var v78 string = "\n\t"
                                                                    var v79 int8 = 127
                                                                    v79 -= (v79 ^ v79)
                                                                    v67 -= uint8(v69)
                                                                    var v80 bool = ((((v69 >> ((0) & 63)) < 2) && v5) && v5)
                                                                    _ = v78
                                                                    _ = v79
                                                                    _ = v80
                                                                } else {
                                                                    var v81 int8 = 0
                                                                    var v82 uint64 = (v69 | (v69 >> ((255) & 63)))
                                                                    _ = v81
                                                                    _ = v82
                                                                }
                                                            }
                                                            var v83 uint32 = v71
                                                            _ = v83
                                                        }
                                                        v5 = v5
                                                        _ = v71
                                                    }
                                                case "":
                                                    v5 = ((1 <= 141) || (32767 != 65535))
                                                }
                                                _ = v70
                                            case "hello":
                                                v67 = (v67 | v67)
                                            default:
                                                var v84 uint16 = 2
                                                _ = v84
                                            }
                                            _ = v68
                                            _ = v69
                                        } else {
                                            _ = len(v3)
                                            switch (^(fi63)) {
                                            case 585:
                                                v85 := GenS9{F0: 32767, F1: 4294967295, F2: 2}
                                                _ = v85
                                            case 32768:
                                                var v86 int64 = -9223372036854775808
                                                _ = v86
                                            default:
                                                var v87 int8 = 42
                                                _ = v87
                                            }
                                            var v88 bool = (!(!false))
                                            var v89 int8 = -128
                                            _ = v88
                                            _ = v89
                                        }
                                        v90 := GenS10{F0: 881}
                                        _ = v67
                                        _ = v90
                                    }
                                    fmt.Println(fi63)
                                    _ = v66
                                }
                            case -2:
                                v5 = (1 < 520)
                            default:
                                v91 := GenS8{F0: 42, F1: true}
                                _ = v91
                            }
                        default:
                            v5 = (v5 && false)
                        }
                        if (-1 <= 547) {
                            v5 = (v5 || (-32768 == -974))
                            fmt.Println(fi63, fi41, fi63)
                            switch (fi63 << ((255) & 63)) {
                            case -93:
                                v5 = true
                            default:
                                fmt.Println(v5, v5)
                            }
                        } else {
                            switch "\n\t" {
                            case "foo bar":
                                var v92 float64 = -2.25
                                _ = v92
                            }
                            var v93 uint64 = 4294967295
                            fmt.Println(fi41, v5, fi41)
                            v93 -= v93
                            _ = v93
                        }
                        delete(v1, 4294967295)
                        fmt.Println(v5, fi41, fi63)
                        v94 := GenS11{F0: 3, F1: 42, F2: 42}
                        v5 = true
                        var v95 int8 = 31
                        _ = fi63
                        _ = v94
                        _ = v95
                    }
                    fmt.Println(fi41)
                    for fi96 := 0; fi96 < 15; fi96++ {
                        var v97 int8 = -1
                        switch (fi41 - -705) {
                        case -158:
                            var v98 uint16 = 255
                            _ = v98
                        default:
                            v1[2] = "foo bar"
                        }
                        _ = fi96
                        _ = v97
                    }
                }
                v5 = ((239 > 0) || (v5 || v5))
            } else {
                fmt.Println(v5, fi41)
            }
            if v5 {
                if (-2.25 == (float64(0.0) * (1.5 - 1e10))) {
                    var v99 float32 = (((1e10 + float32((1.5 * 3.14159))) * 0.0) * (-2.25 + (0.0 - 0.1)))
                    if v5 {
                        var v100 uint64 = 1
                        var v101 int64 = 796
                        _ = v100
                        _ = v101
                    }
                    _ = v99
                }
                var v102 uint16 = 689
                var v103 uint = 534
                _ = v102
                _ = v103
            }
        }
        _ = fi41
    }
    switch -2147483648 {
    case -2147483648:
        if false {
            v5 = true
            var v104 int = 127
            v1[309] = (("hello" + "foo bar") + "foo bar")
            _ = v104
        }
    case -880:
        fmt.Println(v5)
    }
    switch 42 {
    case -412:
        var v105 int8 = -128
        _ = v105
    default:
        var v106 string = "hello"
        _ = v106
    }
    var v107 int16 = -1
    switch "a" {
    case "foo bar":
        fmt.Println(v107)
    case "\n\t":
        v5 = (1 > 42)
    case "a":
        fmt.Println(v5)
    default:
        var v108 int = -695
        _ = v108
    }
    v109 := v2[4294967295]
    v5 = v5
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v107:", v107)
    fmt.Println("v109:", v109)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v107
    _ = v109
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
