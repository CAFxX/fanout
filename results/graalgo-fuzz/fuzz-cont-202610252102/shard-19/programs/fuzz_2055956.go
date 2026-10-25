package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64 }
type GenS2 struct { F0 int16; F1 uint64; F2 int32 }
type GenS3 struct { F0 bool; F1 uint16; F2 string }
type GenS4 struct { F0 uint32; F1 int16 }
type GenS5 struct { F0 uint64; F1 int8 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 int; F1 string }
type GenS9 struct { F0 uint; F1 int64 }
type GenS10 struct { F0 uint64 }
type GenS11 struct { F0 uint32; F1 uint64; F2 bool }
type GenS12 struct { F0 uint32; F1 uint }
type GenS13 struct { F0 int8; F1 int64; F2 uint32 }
type GenS14 struct { F0 int16; F1 uint8; F2 bool }
type GenS15 struct { F0 string }
type GenS16 struct { F0 uint }
type GenS17 struct { F0 uint64; F1 uint; F2 uint }
type GenS18 struct { F0 int; F1 uint16; F2 int8 }
type GenS19 struct { F0 int; F1 uint }
type GenS20 struct { F0 bool; F1 uint }
type GenS21 struct { F0 int; F1 int8; F2 int }
type GenS22 struct { F0 uint16 }
func genFn0(p0 string, p1 bool) int {
    for fi5 := 0; fi5 < 2; fi5++ {
        v6 := []int{fi5, fi5, 0}
        _ = fi5
        _ = v6
    }
    if p1 {
        p0 = p0
        v7 := []int{}
        for fi8 := 0; fi8 < 1; fi8++ {
            if ("\n\t" <= p0) {
                if p1 {
                    for fi9 := 0; fi9 < 11; fi9++ {
                        var v10 uint32 = 0
                        _ = fi9
                        _ = v10
                    }
                    for fi11 := 0; fi11 < 17; fi11++ {
                        v12 := GenS1{F0: "\n\t", F1: 473}
                        _ = fi11
                        _ = v12
                    }
                    switch fi8 {
                    case 128:
                        for fi13 := 0; fi13 < 19; fi13++ {
                            p1 = (0 == 643)
                            switch ("" + p0) {
                            case "foo bar":
                                var v14 float32 = 1.5
                                _ = v14
                            }
                            switch "foo bar" {
                            case "":
                                if (0.0 > (1.5 - 3.14159)) {
                                    if len(v7) > 0 {
                                        v7[0] = fi8
                                    }
                                    if (1 <= 57) {
                                        p1 = p1
                                        v15 := GenS2{F0: 990, F1: 42, F2: -1}
                                        p0 = p0
                                        if p1 {
                                            var v16 string = ""
                                            if p1 {
                                                var v17 uint = 361
                                                _ = v17
                                            }
                                            _ = v16
                                        }
                                        var v18 uint = 2
                                        _ = v15
                                        _ = v18
                                    } else {
                                        var v19 float64 = -2.25
                                        v20 := GenS1{F0: p0, F1: 9223372036854775807}
                                        _ = v19
                                        _ = v20
                                    }
                                    p0 = (("a" + p0) + (p0 + p0))
                                }
                            case "hello":
                                switch (p0 + "\n\t") {
                                case "\n\t":
                                    if (p1 || true) {
                                        p1 = true
                                    } else {
                                        if (0 >= -3) {
                                            var v21 uint8 = 127
                                            _ = v21
                                        }
                                        var v22 string = "hello"
                                        p1 = (734 > 42)
                                        if (4294967295 < 42) {
                                            switch v22 {
                                            case "\n\t":
                                                var v23 int32 = 42
                                                _ = v23
                                            case "hello":
                                                if (p1 || true) {
                                                    var v24 int16 = -889
                                                    _ = v24
                                                } else {
                                                    var v25 float64 = (3.14159 + -2.25)
                                                    _ = v25
                                                }
                                            }
                                        } else {
                                            switch ("a" + "foo bar") {
                                            case "foo bar":
                                                v22 = ""
                                            case "\n\t":
                                                p1 = (1e10 == ((0.1 + 1e10) + float64((3.14159 * (1.5 - 1e10)))))
                                            }
                                        }
                                        if p1 {
                                            switch (v22 + (("" + "\n\t") + (v22 + v22))) {
                                            case "":
                                                if (!((1e10 * (float64(3.14159) - (1.5 * (3.14159 + 1.5)))) == 0.0)) {
                                                    if (-1 <= 1) {
                                                        v26 := GenS1{F0: "foo bar", F1: 42}
                                                        _ = v26
                                                    }
                                                    v27 := GenS1{F0: p0, F1: -9223372036854775808}
                                                    if (1 <= 9223372036854775807) {
                                                        switch -2147483648 {
                                                        case 4294967295:
                                                            if len(v7) > 0 {
                                                                v7[0] = fi13
                                                            }
                                                        case 9223372036854775807:
                                                            v7 = append(v7, (fi13 % (((-(fi13))) | 1)))
                                                        case 32767:
                                                            v28 := GenS1{F0: p0, F1: 0}
                                                            _ = v28
                                                        default:
                                                            if (255 != 973) {
                                                                var v29 uint8 = 128
                                                                fmt.Println(v29)
                                                                _ = v29
                                                            } else {
                                                                var v30 int64 = 9223372036854775807
                                                                var v31 string = v22
                                                                switch v31 {
                                                                case "a":
                                                                    var v32 int8 = 42
                                                                    _ = v32
                                                                case "foo bar":
                                                                    var v33 int64 = (v30 >> ((612) & 63))
                                                                    _ = v33
                                                                default:
                                                                    p1 = p1
                                                                }
                                                                _ = v30
                                                                _ = v31
                                                            }
                                                        }
                                                        if false {
                                                            if (-752 > -32768) {
                                                                var v34 int16 = 1
                                                                fmt.Println(fi8, v34)
                                                                v34 = v34
                                                                _ = v34
                                                            }
                                                            var v35 int8 = -1
                                                            fmt.Println(p0)
                                                            v36 := GenS3{F0: p1, F1: 1, F2: p0}
                                                            _ = v35
                                                            _ = v36
                                                        } else {
                                                            var v37 int32 = -2147483648
                                                            _ = v37
                                                        }
                                                        p0 = "foo bar"
                                                    } else {
                                                        v38 := GenS4{F0: 871, F1: 42}
                                                        var v39 int = fi13
                                                        if len(v7) > 0 {
                                                            v7[0] = (fi13 << ((42) & 63))
                                                        }
                                                        switch fi8 {
                                                        case -1:
                                                            p1 = (42 < 134)
                                                        case -554:
                                                            p0 = v22
                                                            if len(v7) >= 2 {
                                                                v7 = v7[1:len(v7)]
                                                            }
                                                            if p1 {
                                                                fmt.Println(p0, fi13, fi8)
                                                            } else {
                                                                v7 = append(v7, (-(v39)))
                                                            }
                                                        case 127:
                                                            fmt.Println(fi13)
                                                        }
                                                        p0 = "\n\t"
                                                        _ = v38
                                                        _ = v39
                                                    }
                                                    if false {
                                                        if (!(-2147483648 >= fi8)) {
                                                            p1 = p1
                                                            p0 = ("\n\t" + "foo bar")
                                                        }
                                                        if len(v7) > 0 {
                                                            v7[0] = (fi13 + fi8)
                                                        }
                                                        var v40 uint64 = 663
                                                        fmt.Println(v22, fi13, v40)
                                                        if p1 {
                                                            v7 = append(v7, 42)
                                                            if p1 {
                                                                v41 := GenS1{F0: p0, F1: -9223372036854775808}
                                                                _ = v41
                                                            }
                                                            switch ("foo bar" + (((v22 + ((p0 + "a") + "foo bar")) + ("\n\t" + ("a" + ("a" + (((("foo bar" + p0) + (p0 + (("a" + (p0 + v22)) + v22))) + v22) + "hello"))))) + "")) {
                                                            case "":
                                                                fmt.Println(v40, fi13)
                                                            }
                                                            _ = len(v7)
                                                            var v42 uint = 18446744073709551615
                                                            v40 = 634
                                                            _ = v42
                                                        } else {
                                                            if (!(p1 && true)) {
                                                                var v43 uint64 = 382
                                                                if len(v7) > 0 {
                                                                    v7[0] = (fi8 % (((^(fi13))) | 1))
                                                                }
                                                                fmt.Println(fi8, v22, p1)
                                                                var v44 string = v22
                                                                _ = v43
                                                                _ = v44
                                                            } else {
                                                                switch "\n\t" {
                                                                case "\n\t":
                                                                    var v45 string = ("\n\t" + "\n\t")
                                                                    _ = v45
                                                                case "foo bar":
                                                                    var v46 int64 = -1
                                                                    _ = v46
                                                                default:
                                                                    var v47 uint64 = 4294967295
                                                                    _ = v47
                                                                }
                                                                if len(v7) > 0 {
                                                                    v7[0] = 127
                                                                }
                                                                fmt.Println(v22, v22, v22)
                                                            }
                                                            if false {
                                                                if len(v7) > 0 {
                                                                    v7[0] = (fi13 >> ((0) & 63))
                                                                }
                                                                v48 := GenS1{F0: p0, F1: 9223372036854775807}
                                                                v49 := GenS5{F0: (v40 ^ 4294967295), F1: 85}
                                                                _ = v48
                                                                _ = v49
                                                            } else {
                                                                if len(v7) >= 2 {
                                                                    v7 = v7[1:len(v7)]
                                                                }
                                                                switch p0 {
                                                                case "foo bar":
                                                                    var v50 uint32 = 141
                                                                    _ = v50
                                                                }
                                                            }
                                                            _ = len(v7)
                                                            var v51 int64 = 9223372036854775807
                                                            var v52 int = 4294967295
                                                            _ = v51
                                                            _ = v52
                                                        }
                                                        var v53 bool = (v22 >= ("a" + ((("" + p0) + "foo bar") + "")))
                                                        if len(v7) >= 2 {
                                                            v7 = v7[1:len(v7)]
                                                        }
                                                        fmt.Println(v22, fi8, v22)
                                                        var v54 bool = (0 < 2)
                                                        switch "a" {
                                                        case "foo bar":
                                                            var v55 uint = 962
                                                            _ = v55
                                                        }
                                                        _ = v40
                                                        _ = v53
                                                        _ = v54
                                                    } else {
                                                        if ((("\n\t" + p0) != ((v22 + (v22 + v22)) + p0)) && (!true)) {
                                                            if true {
                                                                if p1 {
                                                                    var v56 int64 = 596
                                                                    _ = v56
                                                                }
                                                            } else {
                                                                if len(v7) > 0 {
                                                                    v7[0] = 32768
                                                                }
                                                                fmt.Println(fi13, p0, fi8)
                                                            }
                                                            switch v22 {
                                                            case "foo bar":
                                                                if (!true) {
                                                                    if len(v7) > 0 {
                                                                        v7[0] = fi8
                                                                    }
                                                                    if p1 {
                                                                        fmt.Println(v22, v22)
                                                                    } else {
                                                                        fmt.Println(fi8, fi13)
                                                                        if p1 {
                                                                            v7 = append(v7, (-(fi8)))
                                                                            _ = len(v7)
                                                                            if len(v7) >= 2 {
                                                                                v7 = v7[1:len(v7)]
                                                                            }
                                                                        } else {
                                                                            if p1 {
                                                                                var v57 float32 = 0.0
                                                                                var v58 int = fi13
                                                                                if (0 <= 1) {
                                                                                    v59 := GenS2{F0: 1000, F1: 4294967295, F2: 42}
                                                                                    var v60 float64 = 0.0
                                                                                    v22 = (p0 + p0)
                                                                                    _ = v59
                                                                                    _ = v60
                                                                                }
                                                                                v7 = append(v7, fi8)
                                                                                fmt.Println(p0, v58, fi13)
                                                                                v57 -= v57
                                                                                if (-2.25 != float64(float32(float64((v57 + v57))))) {
                                                                                    if (!p1) {
                                                                                        var v61 uint16 = 255
                                                                                        _ = v61
                                                                                    }
                                                                                    if len(v7) > 0 {
                                                                                        v7[0] = fi13
                                                                                    }
                                                                                }
                                                                                v22 = p0
                                                                                _ = v57
                                                                                _ = v58
                                                                            } else {
                                                                                switch v22 {
                                                                                case "foo bar":
                                                                                    var v62 uint32 = 1
                                                                                    _ = v62
                                                                                default:
                                                                                    if len(v7) >= 2 {
                                                                                        v7 = v7[1:len(v7)]
                                                                                    }
                                                                                }
                                                                            }
                                                                            if false {
                                                                                if len(v7) > 0 {
                                                                                    v7[0] = fi8
                                                                                }
                                                                            } else {
                                                                                v63 := GenS1{F0: "", F1: -1}
                                                                                _ = v63
                                                                            }
                                                                        }
                                                                    }
                                                                    p0 = (p0 + p0)
                                                                } else {
                                                                    switch v22 {
                                                                    case "foo bar":
                                                                        v64 := GenS3{F0: (0 != 128), F1: 958, F2: v22}
                                                                        _ = v64
                                                                    }
                                                                    fmt.Println(p0, p0, fi13)
                                                                }
                                                            default:
                                                                if (true || false) {
                                                                    var v65 uint16 = 583
                                                                    _ = v65
                                                                }
                                                            }
                                                        }
                                                    }
                                                    _ = v27
                                                }
                                            case "\n\t":
                                                var v66 uint16 = 1
                                                _ = v66
                                            case "a":
                                                switch (v22 + (v22 + "")) {
                                                case "hello":
                                                    var v67 int64 = 42
                                                    _ = v67
                                                case "":
                                                    v68 := GenS3{F0: (!(true || p1)), F1: 42, F2: v22}
                                                    switch 4294967295 {
                                                    case 0:
                                                        var v69 float64 = 1.5
                                                        _ = v69
                                                    case -2147483648:
                                                        var v70 uint64 = 952
                                                        _ = v70
                                                    case -1:
                                                        v71 := GenS2{F0: 0, F1: 4294967295, F2: 2147483647}
                                                        _ = v71
                                                    }
                                                    var v72 uint64 = 42
                                                    _ = v68
                                                    _ = v72
                                                default:
                                                    if len(v7) >= 2 {
                                                        v7 = v7[1:len(v7)]
                                                    }
                                                }
                                            }
                                            fmt.Println(p0)
                                            var v73 int64 = 0
                                            switch -1 {
                                            case 935:
                                                v73 = (v73 + v73)
                                            default:
                                                if (true || false) {
                                                    if ((false && (42 > 163)) || true) {
                                                        var v74 string = ("\n\t" + (v22 + "a"))
                                                        p1 = (false && (1.5 > 3.14159))
                                                        _ = v74
                                                    }
                                                    fmt.Println(v73)
                                                    var v75 int64 = 42
                                                    fmt.Println(v73, fi13)
                                                    fmt.Println(v22, fi13, v22)
                                                    var v76 float32 = 3.14159
                                                    fmt.Println(v75, p1)
                                                    var v77 uint = 1
                                                    _ = v75
                                                    _ = v76
                                                    _ = v77
                                                }
                                                var v78 int32 = 2147483647
                                                _ = v78
                                            }
                                            switch fi8 {
                                            case 32768:
                                                var v79 uint32 = 0
                                                _ = v79
                                            case -9223372036854775808:
                                                var v80 uint32 = 0
                                                _ = v80
                                            case 2147483647:
                                                var v81 uint16 = 2
                                                _ = v81
                                            default:
                                                var v82 bool = ((9223372036854775807 >= -1) || true)
                                                if len(v7) >= 2 {
                                                    v7 = v7[1:len(v7)]
                                                }
                                                p1 = true
                                                _ = v82
                                            }
                                            _ = v73
                                        }
                                        _ = v22
                                    }
                                case "hello":
                                    var v83 uint8 = 1
                                    _ = v83
                                case "a":
                                    var v84 int = 1
                                    _ = v84
                                default:
                                    if len(v7) > 0 {
                                        v7[0] = 2147483647
                                    }
                                }
                            case "foo bar":
                                if len(v7) >= 2 {
                                    v7 = v7[1:len(v7)]
                                }
                            default:
                                p0 = p0
                            }
                            _ = fi13
                        }
                    case 277:
                        p1 = (true || (!(42 != 62)))
                    default:
                        p0 = ("foo bar" + (p0 + p0))
                    }
                }
                for fi85 := 0; fi85 < 11; fi85++ {
                    p1 = true
                    if (!p1) {
                        fmt.Println(fi85, fi85, fi8)
                        var v86 bool = false
                        v86 = ("\n\t" == "hello")
                        _ = v86
                    } else {
                        if (!true) {
                            fmt.Println(fi85, fi8, p0)
                        } else {
                            p1 = p1
                            if len(v7) > 0 {
                                v7[0] = fi8
                            }
                            if (1 > 18446744073709551615) {
                                p1 = (!true)
                                if (-2147483648 == 814) {
                                    switch fi8 {
                                    case 0:
                                        p0 = "hello"
                                    case -961:
                                        p0 = ""
                                    default:
                                        if len(v7) >= 2 {
                                            v7 = v7[1:len(v7)]
                                        }
                                    }
                                    if len(v7) >= 2 {
                                        v7 = v7[1:len(v7)]
                                    }
                                } else {
                                    fmt.Println(fi85)
                                    var v87 uint32 = 4294967295
                                    _ = v87
                                }
                                if len(v7) > 0 {
                                    v7[0] = fi85
                                }
                                p0 = (p0 + "foo bar")
                            } else {
                                var v88 uint8 = 255
                                if ((true && p1) || (65535 <= 32767)) {
                                    var v89 uint32 = 0
                                    _ = len(v7)
                                    _ = v89
                                }
                                p1 = p1
                                _ = v88
                            }
                            var v90 int64 = 9223372036854775807
                            _ = v90
                        }
                    }
                    _ = len(v7)
                    _ = fi85
                }
                for fi91 := 0; fi91 < 16; fi91++ {
                    if (-32768 != -297) {
                        p1 = true
                        _ = len(v7)
                        if p1 {
                            if p1 {
                                p1 = (375 < 42)
                                if p1 {
                                    switch p0 {
                                    case "\n\t":
                                        v92 := GenS6{F0: 92}
                                        var v93 string = "\n\t"
                                        p0 = (p0 + (p0 + p0))
                                        _ = v92
                                        _ = v93
                                    case "a":
                                        if p1 {
                                            var v94 float32 = (0.0 + 1.5)
                                            _ = v94
                                        }
                                    }
                                    if len(v7) >= 2 {
                                        v7 = v7[1:len(v7)]
                                    }
                                    if (4294967295 != 65535) {
                                        if (42 <= 42) {
                                            v7 = append(v7, fi8)
                                            v95 := GenS4{F0: 349, F1: 1}
                                            _ = v95
                                        } else {
                                            fmt.Println(fi91, p1)
                                        }
                                    }
                                }
                            } else {
                                v96 := GenS7{F0: 0}
                                var v97 int64 = 42
                                _ = v96
                                _ = v97
                            }
                        }
                        if len(v7) >= 2 {
                            v7 = v7[1:len(v7)]
                        }
                    }
                    v98 := GenS8{F0: 312, F1: "a"}
                    if (4294967295 >= 65535) {
                        var v99 float32 = 3.14159
                        var v100 uint = 42
                        _ = v100
                        _ = v99
                    } else {
                        v101 := GenS9{F0: 570, F1: 945}
                        v102 := GenS8{F0: fi91, F1: p0}
                        p0 = ""
                        v103 := GenS10{F0: 0}
                        _ = v101
                        _ = v102
                        _ = v103
                    }
                    _ = fi91
                    _ = v98
                }
            }
            _ = fi8
        }
        _ = v7
    } else {
        v104 := []int{432, 290, -702}
        if p1 {
            for fi105 := 0; fi105 < 11; fi105++ {
                p0 = "hello"
                var v106 uint32 = 1
                v107 := GenS5{F0: 2, F1: 127}
                _ = fi105
                _ = v106
                _ = v107
            }
            v108 := GenS2{F0: 1000, F1: 42, F2: 182}
            _ = v108
        } else {
            p1 = p1
            if len(v104) >= 2 {
                v104 = v104[1:len(v104)]
            }
            for fi109 := 0; fi109 < 19; fi109++ {
                var v110 uint64 = 1
                if p1 {
                    p1 = p1
                    if false {
                        for fi111 := 0; fi111 < 18; fi111++ {
                            var v112 float64 = 0.1
                            var v113 int16 = 0
                            _ = fi111
                            _ = v112
                            _ = v113
                        }
                    } else {
                        v104 = append(v104, 1)
                        if (1 > 0) {
                            if (!true) {
                                if (fi109 < fi109) {
                                    fmt.Println(p1, fi109)
                                    if len(v104) >= 2 {
                                        v104 = v104[1:len(v104)]
                                    }
                                    p0 = p0
                                } else {
                                    if p1 {
                                        var v114 string = "\n\t"
                                        var v115 uint16 = 1
                                        if false {
                                            v115 += (v115 / ((uint16(v110)) | 1))
                                        } else {
                                            var v116 string = v114
                                            if len(v104) >= 2 {
                                                v104 = v104[1:len(v104)]
                                            }
                                            _ = v116
                                        }
                                        _ = v114
                                        _ = v115
                                    }
                                    if (4294967295 == 65535) {
                                        var v117 uint8 = 43
                                        _ = v117
                                    } else {
                                        v118 := GenS2{F0: -503, F1: (v110 ^ (^(v110))), F2: 42}
                                        switch p0 {
                                        case "hello":
                                            if false {
                                                fmt.Println(p0)
                                            } else {
                                                switch fi109 {
                                                case -2147483648:
                                                    for fi119 := 0; fi119 < 16; fi119++ {
                                                        if p1 {
                                                            var v120 int8 = -128
                                                            var v121 int8 = (v120 - v120)
                                                            p1 = false
                                                            _ = v120
                                                            _ = v121
                                                        }
                                                        fmt.Println(p0, fi119)
                                                        if true {
                                                            v110 += v110
                                                            var v122 uint32 = 65535
                                                            v122 += 42
                                                            _ = v122
                                                        } else {
                                                            _ = len(v104)
                                                        }
                                                        _ = fi119
                                                    }
                                                case 2147483647:
                                                    switch (p0 + p0) {
                                                    case "hello":
                                                        if false {
                                                            var v123 uint32 = 486
                                                            if p1 {
                                                                p0 = ""
                                                                if len(v104) > 0 {
                                                                    v104[0] = fi109
                                                                }
                                                                if (!(((true || p1) || p1) || (p1 && p1))) {
                                                                    if len(v104) >= 2 {
                                                                        v104 = v104[1:len(v104)]
                                                                    }
                                                                    switch -296 {
                                                                    case 2147483647:
                                                                        fmt.Println(p0)
                                                                    case 127:
                                                                        var v124 uint = 255
                                                                        _ = v124
                                                                    }
                                                                    var v125 float32 = 1e10
                                                                    _ = v125
                                                                } else {
                                                                    fmt.Println(p0, p0)
                                                                    if len(v104) > 0 {
                                                                        v104[0] = 899
                                                                    }
                                                                }
                                                            } else {
                                                                v104 = append(v104, fi109)
                                                                var v126 uint16 = 328
                                                                _ = v126
                                                            }
                                                            _ = v123
                                                        } else {
                                                            v110 -= 482
                                                        }
                                                    default:
                                                        v110 -= v110
                                                    }
                                                case 9223372036854775807:
                                                    v127 := GenS11{F0: 42, F1: v110, F2: false}
                                                    _ = v127
                                                default:
                                                    if p1 {
                                                        p1 = p1
                                                        v128 := GenS7{F0: 1}
                                                        v104 = append(v104, 128)
                                                        _ = v128
                                                    } else {
                                                        var v129 int32 = 623
                                                        fmt.Println(p1, p0, v129)
                                                        switch p0 {
                                                        case "foo bar":
                                                            if len(v104) > 0 {
                                                                v104[0] = fi109
                                                            }
                                                        case "hello":
                                                            p1 = p1
                                                        case "":
                                                            v130 := GenS12{F0: 48, F1: 2}
                                                            _ = v130
                                                        default:
                                                            if (127 > 42) {
                                                                fmt.Println(v129, v110)
                                                                switch "\n\t" {
                                                                case "hello":
                                                                    if p1 {
                                                                        if (0.0 <= -2.25) {
                                                                            var v131 float64 = 1.5
                                                                            _ = len(v104)
                                                                            switch fi109 {
                                                                            case 128:
                                                                                var v132 uint8 = 192
                                                                                _ = v132
                                                                            case -2147483648:
                                                                                if (p0 != "\n\t") {
                                                                                    fmt.Println(v131, fi109, fi109)
                                                                                } else {
                                                                                    fmt.Println(v110, p0)
                                                                                    v131 += (float64(1e10) + v131)
                                                                                    for fi133 := 0; fi133 < 3; fi133++ {
                                                                                        if len(v104) > 0 {
                                                                                            v104[0] = (fi133 * (fi109 * (-(fi109))))
                                                                                        }
                                                                                        _ = fi133
                                                                                    }
                                                                                }
                                                                            }
                                                                            _ = v131
                                                                        } else {
                                                                            v134 := GenS7{F0: 42}
                                                                            for fi135 := 0; fi135 < 5; fi135++ {
                                                                                if (-791 != 770) {
                                                                                    var v136 bool = false
                                                                                    if p1 {
                                                                                        v104 = append(v104, 32767)
                                                                                        fmt.Println(v136, p0)
                                                                                    }
                                                                                    switch (p0 + p0) {
                                                                                    case "":
                                                                                        v137 := GenS4{F0: 2, F1: -528}
                                                                                        _ = v137
                                                                                    case "foo bar":
                                                                                        var v138 int64 = 9223372036854775807
                                                                                        _ = v138
                                                                                    }
                                                                                    _ = v136
                                                                                } else {
                                                                                    v139 := GenS13{F0: 127, F1: -712, F2: 792}
                                                                                    var v140 int8 = 1
                                                                                    switch (fi109 / ((fi109) | 1)) {
                                                                                    case -99:
                                                                                        v140 += 108
                                                                                        if len(v104) > 0 {
                                                                                            v104[0] = 0
                                                                                        }
                                                                                        if true {
                                                                                            if true {
                                                                                                fmt.Println(fi109, v140, fi109)
                                                                                            }
                                                                                            var v141 uint32 = 2147483648
                                                                                            _ = v141
                                                                                        } else {
                                                                                            var v142 int32 = (v129 + int32(v110))
                                                                                            var v143 int = fi135
                                                                                            _ = v142
                                                                                            _ = v143
                                                                                        }
                                                                                        var v144 uint8 = 42
                                                                                        _ = v144
                                                                                    case 2147483647:
                                                                                        v145 := GenS8{F0: 0, F1: (p0 + "")}
                                                                                        _ = v145
                                                                                    case 107:
                                                                                        v146 := GenS14{F0: -1, F1: 1, F2: p1}
                                                                                        _ = v146
                                                                                    default:
                                                                                        var v147 uint = 18446744073709551615
                                                                                        _ = v147
                                                                                    }
                                                                                    _ = v139
                                                                                    _ = v140
                                                                                }
                                                                                if ((v110 ^ 726) < (v110 * v110)) {
                                                                                    v148 := GenS3{F0: p1, F1: 32767, F2: p0}
                                                                                    _ = v148
                                                                                }
                                                                                p0 = (((p0 + ("\n\t" + "\n\t")) + p0) + "a")
                                                                                _ = fi135
                                                                            }
                                                                            _ = v134
                                                                        }
                                                                        _ = len(v104)
                                                                    }
                                                                case "":
                                                                    _ = len(v104)
                                                                default:
                                                                    _ = len(v104)
                                                                }
                                                            }
                                                        }
                                                        _ = v129
                                                    }
                                                }
                                                if p1 {
                                                    p1 = (-38 < -1)
                                                }
                                            }
                                        case "foo bar":
                                            for fi149 := 0; fi149 < 11; fi149++ {
                                                v150 := GenS4{F0: 2, F1: int16(fi109)}
                                                _ = fi149
                                                _ = v150
                                            }
                                        case "\n\t":
                                            v151 := GenS15{F0: (p0 + (p0 + "a"))}
                                            _ = v151
                                        }
                                        _ = v118
                                    }
                                    var v152 int = (fi109 >> ((779) & 63))
                                    _ = v152
                                }
                                var v153 float32 = 1e10
                                _ = v153
                            }
                            var v154 uint32 = 42
                            _ = v154
                        }
                    }
                    var v155 bool = true
                    _ = v155
                }
                v156 := GenS5{F0: v110, F1: 42}
                _ = fi109
                _ = v110
                _ = v156
            }
        }
        switch -285 {
        case 32767:
            if (!false) {
                if p1 {
                    var v157 int16 = 1
                    var v158 int16 = 0
                    if p1 {
                        v158 = v157
                        v104 = append(v104, -472)
                        v159 := GenS4{F0: 0, F1: v158}
                        _ = v159
                    } else {
                        if p1 {
                            var v160 uint64 = 42
                            switch p0 {
                            case "":
                                if false {
                                    p0 = "hello"
                                    v161 := GenS16{F0: 65535}
                                    _ = v161
                                } else {
                                    v157 = -1
                                }
                            case "\n\t":
                                for fi162 := 0; fi162 < 8; fi162++ {
                                    v163 := GenS17{F0: (v160 % (((v160 % ((v160) | 1))) | 1)), F1: 65535, F2: 241}
                                    _ = fi162
                                    _ = v163
                                }
                            case "hello":
                                v157 -= v158
                            }
                            if ((v160 - v160) < (v160 % (((v160 - (v160 >> ((231) & 63)))) | 1))) {
                                if len(v104) > 0 {
                                    v104[0] = 309
                                }
                                fmt.Println(v158)
                            }
                            _ = v160
                        } else {
                            fmt.Println(v158, p1)
                            p1 = ((("a" >= p0) || (126 == 1)) && p1)
                        }
                        switch p0 {
                        case "hello":
                            var v164 uint8 = 42
                            _ = v164
                        case "":
                            if len(v104) > 0 {
                                v104[0] = 9223372036854775807
                            }
                        }
                        var v165 uint64 = 18446744073709551615
                        _ = v165
                    }
                    _ = v157
                    _ = v158
                } else {
                    _ = len(v104)
                }
                var v166 uint = 65535
                fmt.Println(p0)
                _ = v166
            }
        case 127:
            var v167 float32 = (((1.5 - (0.1 / (1.5 * -2.25))) + (1.5 - 1.5)) + (1e10 / 1.0))
            _ = v167
        case 1:
            var v168 uint32 = 695
            _ = v168
        default:
            var v169 int16 = 1000
            _ = v169
        }
        _ = v104
    }
    for fi170 := 0; fi170 < 18; fi170++ {
        v171 := []int{fi170, 32767}
        if (!(((p0 + p0) < (p0 + (("\n\t" + ((((p0 + "a") + (p0 + "\n\t")) + (p0 + (p0 + p0))) + p0)) + p0))) || true)) {
            p0 = (p0 + p0)
            for fi172 := 0; fi172 < 5; fi172++ {
                switch p0 {
                case "foo bar":
                    if (!p1) {
                        p0 = (("foo bar" + "foo bar") + ("a" + p0))
                        p1 = true
                        var v173 bool = p1
                        _ = v173
                    } else {
                        switch ("\n\t" + "foo bar") {
                        case "\n\t":
                            switch (("" + (("" + ("a" + (((p0 + p0) + (p0 + p0)) + "foo bar"))) + "\n\t")) + "a") {
                            case "":
                                var v174 int = fi170
                                _ = v174
                            case "hello":
                                var v175 uint16 = 2
                                _ = v175
                            }
                        default:
                            switch fi172 {
                            case -771:
                                v171 = append(v171, fi170)
                            default:
                                var v176 int8 = -1
                                _ = v176
                            }
                        }
                        p1 = (1000 > -453)
                        var v177 uint = 18446744073709551615
                        _ = v177
                    }
                case "a":
                    if (p1 || p1) {
                        var v178 int32 = -365
                        switch 127 {
                        case -1:
                            var v179 string = (p0 + p0)
                            if (4294967295 < 2) {
                                var v180 int = (fi170 - fi170)
                                switch (("a" + p0) + (v179 + v179)) {
                                case "a":
                                    v181 := GenS12{F0: 4294967295, F1: 4294967295}
                                    _ = v181
                                case "":
                                    switch "hello" {
                                    case "\n\t":
                                        var v182 int8 = 120
                                        _ = v182
                                    case "foo bar":
                                        switch (v179 + v179) {
                                        case "\n\t":
                                            if len(v171) >= 2 {
                                                v171 = v171[1:len(v171)]
                                            }
                                        case "a":
                                            var v183 int16 = -26
                                            _ = v183
                                        }
                                    case "hello":
                                        if p1 {
                                            v171 = append(v171, -9223372036854775808)
                                            var v184 uint = 429
                                            if len(v171) >= 2 {
                                                v171 = v171[1:len(v171)]
                                            }
                                            _ = v184
                                        }
                                    }
                                }
                                if p1 {
                                    v180 = (^(fi172))
                                } else {
                                    fmt.Println(fi170)
                                    v171 = append(v171, (fi170 / (((fi170 / (((fi172 * fi170)) | 1))) | 1)))
                                    var v185 bool = p1
                                    var v186 int16 = 0
                                    v179 = v179
                                    _ = v185
                                    _ = v186
                                }
                                _ = v180
                            }
                            if (191 < 4294967295) {
                                v187 := GenS5{F0: 459, F1: -115}
                                if true {
                                    v188 := GenS4{F0: 600, F1: 32767}
                                    var v189 string = "foo bar"
                                    _ = v188
                                    _ = v189
                                } else {
                                    switch (fi170 << ((463) & 63)) {
                                    case 843:
                                        switch ("\n\t" + (v179 + (v179 + "foo bar"))) {
                                        case "hello":
                                            if true {
                                                var v190 float32 = 0.1
                                                var v191 uint = 42
                                                _ = v190
                                                _ = v191
                                            }
                                        case "":
                                            if len(v171) >= 2 {
                                                v171 = v171[1:len(v171)]
                                            }
                                        case "a":
                                            var v192 uint8 = 62
                                            _ = v192
                                        default:
                                            var v193 uint64 = 251
                                            _ = v193
                                        }
                                    case 57:
                                        v194 := GenS4{F0: 710, F1: -32768}
                                        _ = v194
                                    case 32767:
                                        fmt.Println(v178, p1, p1)
                                    }
                                    v195 := GenS3{F0: true, F1: 0, F2: (v179 + p0)}
                                    var v196 uint16 = 602
                                    _ = v195
                                    _ = v196
                                }
                                if p1 {
                                    v178 += (v178 ^ v178)
                                }
                                _ = v187
                            } else {
                                var v197 int8 = -32
                                if ((801 <= 697) || p1) {
                                    fmt.Println(v178)
                                }
                                if (1 != 0) {
                                    v198 := GenS2{F0: 0, F1: 2, F2: v178}
                                    var v199 float64 = 1.5
                                    _ = v198
                                    _ = v199
                                } else {
                                    if len(v171) >= 2 {
                                        v171 = v171[1:len(v171)]
                                    }
                                    var v200 int16 = 42
                                    p1 = (2147483647 < 698)
                                    _ = v200
                                }
                                v201 := GenS18{F0: fi172, F1: 255, F2: (v197 * v197)}
                                _ = v197
                                _ = v201
                            }
                            _ = v179
                        default:
                            if false {
                                v202 := GenS9{F0: 0, F1: 9223372036854775807}
                                switch (p0 + p0) {
                                case "a":
                                    if p1 {
                                        if p1 {
                                            var v203 string = (((p0 + p0) + p0) + p0)
                                            if (v203 <= "a") {
                                                switch ((v203 + "a") + (((((("foo bar" + (("hello" + "\n\t") + v203)) + v203) + p0) + (p0 + v203)) + v203) + v203)) {
                                                case "":
                                                    if false {
                                                        switch 127 {
                                                        case 128:
                                                            v203 = (("a" + v203) + "a")
                                                        case 42:
                                                            p0 = v203
                                                        default:
                                                            if len(v171) > 0 {
                                                                v171[0] = 32768
                                                            }
                                                        }
                                                    } else {
                                                        switch fi172 {
                                                        case 4294967295:
                                                            if (true && p1) {
                                                                if (((0.1 / (1.5 - (1e10 + float32(-2.25)))) - 0.0) <= 1.5) {
                                                                    if p1 {
                                                                        p0 = p0
                                                                        if ((fi170 * (fi170 % ((fi172) | 1))) < 9223372036854775807) {
                                                                            var v204 bool = ((19 == 128) && (4294967295 < 2))
                                                                            fmt.Println(p0, p0, v203)
                                                                            var v205 int32 = (v178 - (v178 << ((293) & 63)))
                                                                            _ = v204
                                                                            _ = v205
                                                                        } else {
                                                                            fmt.Println(v178, fi170, v203)
                                                                        }
                                                                        if (p1 && p1) {
                                                                            v206 := GenS19{F0: 0, F1: 2}
                                                                            var v207 int16 = -94
                                                                            p0 = v203
                                                                            var v208 bool = (0.1 < 1e10)
                                                                            v207 += v207
                                                                            if len(v171) > 0 {
                                                                                v171[0] = fi170
                                                                            }
                                                                            _ = v206
                                                                            _ = v207
                                                                            _ = v208
                                                                        }
                                                                        _ = len(v171)
                                                                        v203 = v203
                                                                    } else {
                                                                        if true {
                                                                            fmt.Println(p1, v203)
                                                                        }
                                                                        var v209 int32 = (v178 + v178)
                                                                        _ = v209
                                                                    }
                                                                }
                                                                p0 = v203
                                                            }
                                                            v178 = (v178 * v178)
                                                            fmt.Println(p1, p0)
                                                        case 127:
                                                            if (128 <= 2) {
                                                                switch -2147483648 {
                                                                case -9223372036854775808:
                                                                    if len(v171) >= 2 {
                                                                        v171 = v171[1:len(v171)]
                                                                    }
                                                                default:
                                                                    if p1 {
                                                                        if ((v178 / ((v178) | 1)) < (^(v178))) {
                                                                            v171 = append(v171, (fi170 ^ fi172))
                                                                            fmt.Println(v203)
                                                                            fmt.Println(v178)
                                                                        }
                                                                        v210 := GenS7{F0: 1}
                                                                        _ = v210
                                                                    } else {
                                                                        v211 := GenS20{F0: (100 <= 88), F1: 982}
                                                                        _ = v211
                                                                    }
                                                                }
                                                            } else {
                                                                var v212 int8 = 1
                                                                v213 := GenS16{F0: 217}
                                                                v178 += v178
                                                                var v214 uint32 = 2
                                                                fmt.Println(p0, p1, p0)
                                                                v178 = (^(v178))
                                                                _ = v212
                                                                _ = v213
                                                                _ = v214
                                                            }
                                                        }
                                                        var v215 float32 = (1.5 - 0.0)
                                                        if p1 {
                                                            if p1 {
                                                                p1 = p1
                                                                fmt.Println(v178)
                                                            }
                                                        }
                                                        var v216 int32 = v178
                                                        _ = v215
                                                        _ = v216
                                                    }
                                                default:
                                                    var v217 int32 = 42
                                                    _ = v217
                                                }
                                                if len(v171) > 0 {
                                                    v171[0] = (fi172 - 42)
                                                }
                                            }
                                            _ = v203
                                        }
                                    }
                                case "\n\t":
                                    if len(v171) >= 2 {
                                        v171 = v171[1:len(v171)]
                                    }
                                case "":
                                    v178 = (v178 * (v178 & (v178 & -846)))
                                }
                                fmt.Println(fi172)
                                _ = v202
                            }
                        }
                        fmt.Println(p1)
                        _ = v178
                    } else {
                        p1 = p1
                        p1 = (false || p1)
                        var v218 float32 = 3.14159
                        _ = v218
                    }
                case "hello":
                    p1 = ((1000 < 32767) && ((1e10 + 1e10) <= (3.14159 / float64((0.1 + 1.5)))))
                }
                if p1 {
                    var v219 uint64 = 42
                    v220 := GenS11{F0: 1, F1: 175, F2: (false || p1)}
                    p0 = (p0 + p0)
                    _ = v219
                    _ = v220
                }
                _ = fi172
            }
            var v221 uint64 = 1
            _ = v221
        } else {
            p0 = "foo bar"
        }
        fmt.Println(p0, p1, p1)
        _ = fi170
        _ = v171
    }
    for fi222 := 0; fi222 < 9; fi222++ {
        for fi223 := 0; fi223 < 1; fi223++ {
            if p1 {
                var v224 bool = p1
                v225 := []int{}
                _ = v224
                _ = v225
            }
            p1 = p1
            if p1 {
                if ((p0 + p0) >= (("hello" + (p0 + ("foo bar" + ""))) + "a")) {
                    v226 := GenS21{F0: (fi222 % ((-929) | 1)), F1: -1, F2: (fi223 * fi223)}
                    v227 := GenS10{F0: 727}
                    p0 = "hello"
                    _ = v226
                    _ = v227
                }
            } else {
                var v228 uint32 = 2
                var v229 string = "foo bar"
                _ = v228
                _ = v229
            }
            var v230 uint64 = 311
            if p1 {
                v231 := GenS10{F0: (^(v230))}
                v230 += (v230 ^ v230)
                _ = v231
            } else {
                v232 := []int{(fi222 | (-(fi223))), 451, int(v230)}
                _ = v232
            }
            _ = fi223
            _ = v230
        }
        _ = fi222
    }
    _ = p0
    _ = p1
    return -9223372036854775808
}
func genFn1(p0 string, p1 uint) int16 {
    if (!true) {
        for fi233 := 0; fi233 < 14; fi233++ {
            v234 := []int{}
            _ = fi233
            _ = v234
        }
        if true {
            var v235 int = 0
            p1 += p1
            _ = v235
        }
    }
    if (false || true) {
        var v236 int32 = -2147483648
        _ = v236
    }
    fmt.Println(p0, p1)
    v237 := GenS6{F0: 1}
    for fi238 := 0; fi238 < 18; fi238++ {
        for fi239 := 0; fi239 < 9; fi239++ {
            fmt.Println(p1)
            _ = fi239
        }
        _ = fi238
    }
    _ = p0
    _ = p1
    _ = v237
    return -141
}

func main() {
    var v1 uint8 = 0
    v2 := []uint32{0, 790}
    var v3 string = ""
    var v4 int32 = 42
    v3 = v3
    v3 = v3
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v4 += (v4 / ((v4) | 1))
    switch 1 {
    case 128:
        v240 := GenS9{F0: 65535, F1: 851}
        _ = v240
    case 32768:
        if (127 >= 99) {
            if false {
                if false {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    var v241 int8 = -128
                    var v242 uint8 = v1
                    _ = v241
                    _ = v242
                }
            }
            v3 = ("\n\t" + (v3 + (v3 + "")))
        }
    }
    if false {
        v243 := GenS22{F0: 42}
        var v244 float32 = 3.14159
        _ = v243
        _ = v244
    } else {
        v3 = "\n\t"
        switch 32768 {
        case 2147483647:
            v245 := GenS3{F0: ((0.0 * 1e10) > 0.1), F1: uint16(v1), F2: (("a" + v3) + "foo bar")}
            _ = v245
        default:
            v4 -= 42
        }
    }
    if true {
        if true {
            var v246 uint16 = 42
            var v247 int16 = -32768
            _ = v246
            _ = v247
        }
        for fi248 := 0; fi248 < 14; fi248++ {
            var v249 uint16 = 2
            v1 -= v1
            _ = fi248
            _ = v249
        }
        _ = len(v2)
    } else {
        if (!(false && false)) {
            if false {
                var v250 uint32 = 2147483647
                fmt.Println(v1, v250, v1)
                _ = v250
            } else {
                var v251 float32 = (3.14159 - 1e10)
                _ = v251
            }
        } else {
            v252 := GenS20{F0: ((((((0.0 * (0.0 * float64(1.5))) * float64(-2.25)) * 1.5) - 3.14159) - (-2.25 * float64(0.0))) < 0.1), F1: 2}
            _ = v252
        }
        v1 = v1
        var v253 float64 = (-2.25 + (0.1 / -2.25))
        _ = v253
    }
    var v254 uint8 = 169
    for fi255 := 0; fi255 < 8; fi255++ {
        switch fi255 {
        case 2147483647:
            v2 = append(v2, 0)
        case 9223372036854775807:
            var v256 int8 = 0
            _ = v256
        default:
            if false {
                v1 += v254
            }
        }
        _ = fi255
    }
    v1 += v1
    if false {
        v254 = 47
    } else {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v254:", v254)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v254
    _ = v3
    _ = v4
}
