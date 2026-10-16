package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64 }
type GenS2 struct { F0 uint32; F1 int; F2 uint8 }
type GenS3 struct { F0 int32; F1 uint32 }
type GenS4 struct { F0 int64; F1 int32; F2 int64 }
type GenS5 struct { F0 int8; F1 bool }
type GenS6 struct { F0 int32; F1 int16; F2 int16 }
type GenS7 struct { F0 int32; F1 bool; F2 int16 }
type GenS8 struct { F0 uint32; F1 uint64 }
type GenS9 struct { F0 uint; F1 int; F2 uint64 }
type GenS10 struct { F0 int64 }
type GenS11 struct { F0 uint64; F1 uint64; F2 uint8 }
type GenS12 struct { F0 int64; F1 uint64 }

func main() {
    v1 := []string{"\n\t", ("\n\t" + "")}
    v2 := make(map[string]bool)
    v3 := []uint32{}
    v4 := []int{}
    v5 := make(map[uint32]string)
    for fi6 := 0; fi6 < 10; fi6++ {
        switch "\n\t" {
        case "":
            var v7 string = "hello"
            _ = v7
        case "a":
            var v8 int = (fi6 | fi6)
            _ = v8
        case "hello":
            switch "" {
            case "hello":
                switch fi6 {
                case 32768:
                    if false {
                        var v9 uint = 255
                        v9 = (v9 * (v9 / ((v9) | 1)))
                        _ = v9
                    }
                }
            case "foo bar":
                if false {
                    var v10 int = fi6
                    for fi11 := 0; fi11 < 1; fi11++ {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        switch "" {
                        case "a":
                            v10 += (fi11 % ((fi6) | 1))
                        default:
                            v10 = fi11
                        }
                        v10 += fi11
                        _ = fi11
                    }
                    _ = v10
                } else {
                    if false {
                        var v12 uint8 = 127
                        _ = v12
                    }
                    var v13 int64 = 9223372036854775807
                    v14 := GenS1{F0: 462, F1: (v13 - 42)}
                    _ = v13
                    _ = v14
                }
            case "\n\t":
                v15 := GenS1{F0: 1000, F1: 1}
                _ = v15
            default:
                if true {
                    if len(v3) > 0 {
                        v3[0] = 42
                    }
                    switch "foo bar" {
                    case "a":
                        for fi16 := 0; fi16 < 15; fi16++ {
                            if (!(301 <= 255)) {
                                if false {
                                    v17 := v2[("\n\t" + "foo bar")]
                                    fmt.Println(fi16, v17)
                                    var v18 uint64 = 157
                                    _ = v17
                                    _ = v18
                                } else {
                                    switch ("a" + "") {
                                    case "":
                                        if (1e10 == float32(float64(0.1))) {
                                            var v19 int16 = -32768
                                            fmt.Println(v19, v19)
                                            var v20 uint32 = 2147483648
                                            _ = v19
                                            _ = v20
                                        } else {
                                            fmt.Println(fi6)
                                            _ = len(v3)
                                        }
                                    case "hello":
                                        _ = len(v5)
                                    }
                                }
                                v21 := GenS1{F0: -103, F1: -1}
                                _ = v21
                            } else {
                                if (("" == "\n\t") && ((475 == 76) || false)) {
                                    var v22 float32 = 1e10
                                    var v23 int = fi6
                                    _ = v22
                                    _ = v23
                                } else {
                                    v24 := GenS1{F0: 42, F1: 9223372036854775807}
                                    if ((1.5 + -2.25) < -2.25) {
                                        var v25 float64 = 3.14159
                                        fmt.Println(fi16, v25)
                                        _ = v25
                                    }
                                    _ = v24
                                }
                            }
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            _ = fi16
                        }
                    case "":
                        for fi26 := 0; fi26 < 3; fi26++ {
                            var v27 uint8 = 0
                            if true {
                                switch (fi6 / ((int(v27)) | 1)) {
                                case 32768:
                                    var v28 float64 = 1.5
                                    fmt.Println(fi6, fi26)
                                    switch "\n\t" {
                                    case "":
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                    case "\n\t":
                                        var v29 int64 = 87
                                        _ = v29
                                    case "foo bar":
                                        if (1 == -627) {
                                            v27 = v27
                                            switch "hello" {
                                            case "a":
                                                fmt.Println(fi26)
                                            default:
                                                fmt.Println(fi26, fi6, fi26)
                                            }
                                            var v30 int64 = 42
                                            fmt.Println(fi26, fi6, v30)
                                            v30 = 42
                                            _ = v30
                                        } else {
                                            if false {
                                                if (9223372036854775807 > -1) {
                                                    v27 = (v27 * v27)
                                                } else {
                                                    var v31 int = (fi6 >> ((2) & 63))
                                                    switch (fi26 << ((4294967295) & 63)) {
                                                    case 417:
                                                        if true {
                                                            var v32 int8 = 42
                                                            _ = v32
                                                        }
                                                    }
                                                    delete(v2, "a")
                                                    if ((!(!(((v28 / v28) * 1.5) < (3.14159 / v28)))) || (0.1 != (1.5 - (0.0 * (float32(((((float64(1.5) * ((v28 + v28) - v28)) - 0.1) / 1e10) + (v28 / v28))) - 0.0))))) {
                                                        var v33 uint64 = 0
                                                        var v34 uint = 2
                                                        if (18446744073709551615 <= v34) {
                                                            v31 = (v31 * fi6)
                                                            v27 = (v27 % ((127) | 1))
                                                            v35 := GenS1{F0: 42, F1: 0}
                                                            _ = v35
                                                        } else {
                                                            fmt.Println(fi26, fi26)
                                                            var v36 string = ((("a" + "a") + "hello") + "")
                                                            var v37 float64 = v28
                                                            _ = v36
                                                            _ = v37
                                                        }
                                                        _ = v33
                                                        _ = v34
                                                    } else {
                                                        if true {
                                                            var v38 int64 = 1
                                                            v39 := GenS2{F0: 660, F1: -277, F2: v27}
                                                            var v40 int8 = 100
                                                            _ = v38
                                                            _ = v39
                                                            _ = v40
                                                        } else {
                                                            v28 = 0.0
                                                        }
                                                        if len(v3) > 0 {
                                                            v3[0] = 2147483647
                                                        }
                                                        v41 := GenS3{F0: 776, F1: 861}
                                                        _ = v41
                                                    }
                                                    var v42 float64 = float64(float32((v28 * 0.0)))
                                                    v42 += 3.14159
                                                    _ = v31
                                                    _ = v42
                                                }
                                                if (true && true) {
                                                    v43 := GenS1{F0: -1, F1: -9223372036854775808}
                                                    v27 += (v27 >> ((255) & 63))
                                                    var v44 bool = (958 == 255)
                                                    _ = v43
                                                    _ = v44
                                                }
                                                if false {
                                                    var v45 string = "foo bar"
                                                    switch fi26 {
                                                    case 2147483647:
                                                        v4 = append(v4, 42)
                                                    case 67:
                                                        v45 = ""
                                                    }
                                                    _ = v45
                                                } else {
                                                    switch (fi26 >> ((44) & 63)) {
                                                    case -618:
                                                        fmt.Println(v28, fi26)
                                                    case -9223372036854775808:
                                                        var v46 int = -9223372036854775808
                                                        _ = v46
                                                    case 734:
                                                        var v47 uint32 = 42
                                                        _ = v47
                                                    default:
                                                        if false {
                                                            v48 := GenS4{F0: 0, F1: -1, F2: -9223372036854775808}
                                                            v49 := GenS5{F0: 127, F1: true}
                                                            var v50 uint16 = 855
                                                            var v51 uint32 = 2
                                                            fmt.Println(v28)
                                                            _ = v48
                                                            _ = v49
                                                            _ = v50
                                                            _ = v51
                                                        }
                                                    }
                                                }
                                            }
                                            switch "hello" {
                                            case "\n\t":
                                                v52 := v2[("" + ((("a" + (("hello" + "\n\t") + "a")) + "hello") + "foo bar"))]
                                                _ = v52
                                            case "":
                                                var v53 uint64 = 2
                                                _ = v53
                                            default:
                                                fmt.Println(v27, v28)
                                            }
                                            v54 := GenS3{F0: 42, F1: 4294967295}
                                            if true {
                                                v27 = v27
                                                v55 := GenS6{F0: 42, F1: 42, F2: -274}
                                                var v56 uint16 = 2
                                                _ = v55
                                                _ = v56
                                            } else {
                                                if true {
                                                    delete(v2, "\n\t")
                                                } else {
                                                    if false {
                                                        switch (fi26 - fi26) {
                                                        case -9223372036854775808:
                                                            switch ((("foo bar" + "foo bar") + ("\n\t" + ("" + ("\n\t" + (("\n\t" + "\n\t") + "foo bar"))))) + "\n\t") {
                                                            case "foo bar":
                                                                var v57 int16 = 1
                                                                _ = v57
                                                            case "\n\t":
                                                                var v58 string = ("foo bar" + "")
                                                                v28 = v28
                                                                var v59 string = ""
                                                                _ = v58
                                                                _ = v59
                                                            case "a":
                                                                var v60 uint8 = (v27 - (v27 / (((v27 >> ((2) & 63))) | 1)))
                                                                _ = v60
                                                            }
                                                        case 418:
                                                            if (3.14159 <= 1.5) {
                                                                v61 := GenS7{F0: 42, F1: false, F2: 148}
                                                                if ((v28 + ((float64(1e10) / 1.0) * v28)) == 0.1) {
                                                                    v62 := GenS4{F0: 42, F1: -292, F2: -1}
                                                                    _ = v62
                                                                }
                                                                _ = v61
                                                            }
                                                        default:
                                                            _ = len(v3)
                                                            var v63 uint8 = (v27 >> ((106) & 63))
                                                            fmt.Println(fi6)
                                                            _ = v63
                                                        }
                                                        if len(v4) > 0 {
                                                            v4[0] = fi26
                                                        }
                                                    } else {
                                                        if true {
                                                            v28 -= (v28 - 0.0)
                                                            switch ((("foo bar" + "a") + ("hello" + ((("\n\t" + "foo bar") + "\n\t") + "a"))) + "") {
                                                            case "a":
                                                                fmt.Println(v28, v27)
                                                            case "":
                                                                v64 := GenS8{F0: 761, F1: 0}
                                                                _ = v64
                                                            }
                                                        } else {
                                                            if (42 != 1000) {
                                                                var v65 int32 = -2147483648
                                                                _ = v65
                                                            }
                                                            var v66 float64 = v28
                                                            fmt.Println(v66, fi6)
                                                            _ = v66
                                                        }
                                                    }
                                                    delete(v2, "")
                                                    v67 := GenS9{F0: 494, F1: fi26, F2: 1}
                                                    _ = v67
                                                }
                                                v28 = v28
                                                var v68 uint32 = 492
                                                if (446 >= 0) {
                                                    if (2 >= 255) {
                                                        v69 := GenS10{F0: -1}
                                                        v68 += (v68 * (^(v68)))
                                                        var v70 uint32 = (v68 << ((198) & 63))
                                                        _ = v69
                                                        _ = v70
                                                    } else {
                                                        v71 := GenS5{F0: -1, F1: ((-(fi26)) != fi6)}
                                                        if false {
                                                            var v72 int64 = 9223372036854775807
                                                            v1 = append(v1, "")
                                                            var v73 bool = (((!true) || false) && (false || ((v27 % ((128) | 1)) != (v27 & (^(v27))))))
                                                            _ = v72
                                                            _ = v73
                                                        }
                                                        if ((false && ((0 < 1) && true)) && true) {
                                                            if (2147483648 != (v68 & (v68 % ((v68) | 1)))) {
                                                                if len(v4) > 0 {
                                                                    v4[0] = 509
                                                                }
                                                                if len(v4) > 0 {
                                                                    v4[0] = (fi6 << ((349) & 63))
                                                                }
                                                            }
                                                        }
                                                        _ = v71
                                                    }
                                                    if ((18446744073709551615 == 18) || (42 >= 82)) {
                                                        var v74 uint = 65535
                                                        if (2 != 4294967295) {
                                                            v5[v68] = "foo bar"
                                                        }
                                                        v68 = 2147483647
                                                        var v75 bool = (!(((true && (2 <= 184)) || (2 > 65535)) || ((("" <= "foo bar") || (v27 == (v27 | (v27 + (v27 - v27))))) && (!((v28 * ((((((v28 + v28) - ((v28 / v28) - float64(1.5))) - float64(3.14159)) / float64((3.14159 * 0.1))) / 1.0) / 1.5)) != 0.0)))))
                                                        _ = v74
                                                        _ = v75
                                                    } else {
                                                        var v76 bool = (-128 < 127)
                                                        _ = v76
                                                    }
                                                }
                                                _ = v68
                                            }
                                            _ = v54
                                        }
                                    default:
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                    }
                                    _ = v28
                                case -810:
                                    delete(v5, 65535)
                                case -9223372036854775808:
                                    var v77 int = fi26
                                    _ = v77
                                default:
                                    fmt.Println(fi6, fi6)
                                }
                            } else {
                                v3 = append(v3, 571)
                            }
                            _ = fi26
                            _ = v27
                        }
                    default:
                        var v78 int = -9223372036854775808
                        _ = v78
                    }
                }
            }
        }
        _ = fi6
    }
    var v79 string = ""
    switch (v79 + "hello") {
    case "hello":
        var v80 bool = true
        _ = v80
    default:
        v79 = v79
    }
    var v81 float64 = (1e10 + 1e10)
    _ = len(v2)
    for fi82 := 0; fi82 < 8; fi82++ {
        v79 = (v79 + (v79 + "\n\t"))
        var v83 int8 = 0
        _ = fi82
        _ = v83
    }
    v81 = v81
    v79 = ("hello" + ((v79 + (v79 + v79)) + v79))
    v81 += (0.1 + (v81 - 1.5))
    _ = len(v1)
    if (4294967295 == 860) {
        switch v79 {
        case "\n\t":
            _ = len(v5)
        }
        v84 := GenS3{F0: 42, F1: 2}
        if len(v3) > 0 {
            v3[0] = 1
        }
        _ = v84
    } else {
        v81 += 1.5
        v79 = (((((v79 + "") + v79) + "foo bar") + v79) + v79)
    }
    if (float64((1e10 / 3.14159)) == v81) {
        if (4294967295 != 727) {
            v79 = (v79 + v79)
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            for fi85 := 0; fi85 < 13; fi85++ {
                if (!false) {
                    switch -889 {
                    case -2147483648:
                        v5[4294967295] = v79
                    default:
                        var v86 uint32 = 65535
                        _ = v86
                    }
                    for fi87 := 0; fi87 < 12; fi87++ {
                        if false {
                            v81 -= (1e10 * v81)
                            if false {
                                v88 := GenS3{F0: -2147483648, F1: 1}
                                _ = v88
                            } else {
                                v79 = v79
                            }
                            v79 = "foo bar"
                            if (65535 >= 1) {
                                v89 := GenS11{F0: 2, F1: 309, F2: 20}
                                switch v79 {
                                case "foo bar":
                                    var v90 int16 = 1
                                    switch ("foo bar" + "") {
                                    case "hello":
                                        v91 := GenS8{F0: 0, F1: 858}
                                        _ = v91
                                    case "a":
                                        if true {
                                            v90 = 260
                                            v92 := v2[v79]
                                            v93 := GenS7{F0: 206, F1: (!(!false)), F2: v90}
                                            _ = v92
                                            _ = v93
                                        }
                                    case "foo bar":
                                        _ = len(v2)
                                    }
                                    _ = v90
                                }
                                _ = v89
                            } else {
                                var v94 uint32 = 2147483648
                                var v95 int64 = 0
                                _ = v94
                                _ = v95
                            }
                            if false {
                                v96 := GenS12{F0: 605, F1: 2}
                                _ = v96
                            }
                            if (1 > 213) {
                                fmt.Println(fi87)
                                v2[""] = ((124 > 903) || true)
                            }
                            v97 := GenS8{F0: 0, F1: 2}
                            _ = v97
                        } else {
                            v81 -= float64(1.5)
                        }
                        var v98 int32 = -59
                        var v99 uint64 = 1
                        _ = fi87
                        _ = v98
                        _ = v99
                    }
                    v100 := GenS11{F0: 42, F1: 0, F2: 128}
                    _ = v100
                }
                _ = fi85
            }
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v79:", v79)
    fmt.Println("v81:", v81)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v79
    _ = v81
}
