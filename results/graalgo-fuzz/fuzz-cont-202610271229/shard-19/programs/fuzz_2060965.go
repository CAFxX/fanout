package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int8; F1 uint8; F2 uint32 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 string; F1 uint; F2 uint64 }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 int; F1 uint32; F2 int32 }
type GenS10 struct { F0 uint }
type GenS11 struct { F0 bool; F1 int32; F2 bool }
type GenS12 struct { F0 int64; F1 uint32; F2 int8 }
type GenS13 struct { F0 int16 }
type GenS14 struct { F0 int32; F1 int8 }
type GenS15 struct { F0 uint32 }
type GenS16 struct { F0 string; F1 int }
type GenS17 struct { F0 uint8; F1 int16; F2 uint8 }
type GenS18 struct { F0 int; F1 bool }
type GenS19 struct { F0 uint }
type GenS20 struct { F0 uint64; F1 int }
type GenS21 struct { F0 int8; F1 uint32; F2 int8 }
type GenS22 struct { F0 uint16 }
type GenS23 struct { F0 int32; F1 uint }

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[string]int)
    v3 := make(map[uint32]int)
    v4 := []string{"", (("" + "a") + "a"), ("hello" + ("" + "a"))}
    for fi5 := 0; fi5 < 4; fi5++ {
        var v6 float32 = 3.14159
        for fi7 := 0; fi7 < 13; fi7++ {
            switch "foo bar" {
            case "\n\t":
                if false {
                    v8 := GenS1{F0: -1}
                    fmt.Println(fi7, v6, fi7)
                    _ = v8
                } else {
                    var v9 uint32 = 2147483648
                    _ = v9
                }
                var v10 uint16 = 255
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                _ = v10
            case "a":
                switch (fi7 & (fi5 >> ((85) & 63))) {
                case -1:
                    if false {
                        if true {
                            if (1 > 127) {
                                var v11 bool = false
                                _ = len(v3)
                                if v11 {
                                    v11 = (!(0 <= 0))
                                    v12 := GenS1{F0: 0}
                                    var v13 float64 = 1e10
                                    _ = v12
                                    _ = v13
                                } else {
                                    switch -2147483648 {
                                    case -343:
                                        delete(v3, 2)
                                    default:
                                        if (false && true) {
                                            v11 = (((("foo bar" + "\n\t") + "\n\t") > "foo bar") || true)
                                        } else {
                                            fmt.Println(fi5, fi7, v11)
                                        }
                                    }
                                }
                                switch "" {
                                case "hello":
                                    var v14 uint64 = 18446744073709551615
                                    _ = v14
                                case "foo bar":
                                    v15 := GenS1{F0: 431}
                                    v11 = v11
                                    _ = v15
                                }
                                _ = v11
                            }
                        } else {
                            v16 := GenS1{F0: -350}
                            if true {
                                var v17 int32 = -2147483648
                                fmt.Println(fi5)
                                _ = v17
                            } else {
                                var v18 float32 = 1.5
                                if (true && false) {
                                    switch fi7 {
                                    case -646:
                                        v3[2147483647] = (fi5 - fi5)
                                    case 9223372036854775807:
                                        switch (fi5 & (fi5 ^ fi7)) {
                                        case -2147483648:
                                            v3[2147483647] = fi5
                                        case -974:
                                            if (true || (!(((0.0 * 1e10) + (0.0 * 3.14159)) == 3.14159))) {
                                                fmt.Println(fi5, fi7)
                                                fmt.Println(v18, fi7)
                                                if (true && false) {
                                                    v6 = -2.25
                                                    var v19 int32 = int32(fi5)
                                                    _ = v19
                                                } else {
                                                    v6 = v18
                                                    var v20 int32 = int32(fi5)
                                                    _ = v20
                                                }
                                            }
                                        default:
                                            var v21 uint8 = 246
                                            _ = v21
                                        }
                                    case 804:
                                        var v22 float64 = 0.0
                                        _ = v22
                                    default:
                                        switch fi5 {
                                        case -2147483648:
                                            _ = len(v2)
                                        case -821:
                                            switch "foo bar" {
                                            case "foo bar":
                                                switch 4294967295 {
                                                case 128:
                                                    v1[528] = "\n\t"
                                                case -9223372036854775808:
                                                    v18 = v18
                                                case 452:
                                                    v18 = -2.25
                                                default:
                                                    fmt.Println(fi7, v6, v6)
                                                }
                                            case "hello":
                                                v18 = v6
                                            }
                                            var v23 float32 = v6
                                            _ = v23
                                        default:
                                            if len(v4) > 0 {
                                                v4[0] = "a"
                                            }
                                        }
                                    }
                                    v6 = (((((1.5 - 3.14159) + v18) / ((1e10 / v6) / (-2.25 / (0.0 / v6)))) / 1.5) + v18)
                                    if len(v4) > 0 {
                                        v4[0] = "\n\t"
                                    }
                                    var v24 int64 = -9223372036854775808
                                    _ = v24
                                } else {
                                    v18 = 0.0
                                    var v25 string = "a"
                                    delete(v3, 952)
                                    var v26 uint16 = 797
                                    var v27 int = (^(fi5))
                                    v26 += v26
                                    v27 = fi7
                                    if false {
                                        v28 := GenS1{F0: 42}
                                        var v29 uint16 = v26
                                        _ = v28
                                        _ = v29
                                    }
                                    _ = v25
                                    _ = v26
                                    _ = v27
                                }
                                if (849 == 299) {
                                    var v30 int16 = -127
                                    var v31 uint8 = 127
                                    fmt.Println(fi5)
                                    v18 = (float32(float64(v18)) / 1.0)
                                    _ = len(v4)
                                    _ = v30
                                    _ = v31
                                } else {
                                    var v32 uint16 = 832
                                    v32 = (^(v32))
                                    _ = v32
                                }
                                _ = v18
                            }
                            v6 -= v6
                            var v33 uint8 = 128
                            var v34 int32 = 2147483647
                            fmt.Println(fi7, v33)
                            _ = v16
                            _ = v33
                            _ = v34
                        }
                        v35 := v1[4294967295]
                        v36 := GenS2{F0: 0, F1: 127, F2: 2147483647}
                        v6 -= (v6 / ((v6 * v6) - (1e10 + (v6 - 0.0))))
                        if false {
                            v4 = append(v4, v35)
                            var v37 int8 = -37
                            fmt.Println(fi7)
                            if true {
                                fmt.Println(fi7, v6)
                                v35 = (v35 + v35)
                                fmt.Println(v6)
                            }
                            _ = v37
                        } else {
                            v6 += 1.5
                            if true {
                                v35 = v35
                                v2[("hello" + "hello")] = 984
                            } else {
                                if (!false) {
                                    fmt.Println(v6)
                                    var v38 uint32 = 925
                                    _ = v38
                                } else {
                                    if true {
                                        var v39 uint32 = 42
                                        var v40 uint32 = (v39 / ((0) | 1))
                                        var v41 int8 = 1
                                        _ = v39
                                        _ = v40
                                        _ = v41
                                    } else {
                                        if true {
                                            if ((fi7 - 1) > fi7) {
                                                var v42 int32 = -2147483648
                                                _ = v42
                                            } else {
                                                fmt.Println(fi5, fi7)
                                                v1[564] = "\n\t"
                                            }
                                            switch "" {
                                            case "\n\t":
                                                var v43 int64 = 42
                                                _ = v43
                                            case "a":
                                                if (78 < 255) {
                                                    if false {
                                                        var v44 string = v35
                                                        v45 := GenS3{F0: 0}
                                                        _ = v44
                                                        _ = v45
                                                    } else {
                                                        v6 = v6
                                                        if (true || ((!true) || (647 == -9223372036854775808))) {
                                                            v35 = "\n\t"
                                                            fmt.Println(v6)
                                                            v46 := GenS3{F0: 0}
                                                            _ = v46
                                                        }
                                                        if true {
                                                            var v47 int = -2147483648
                                                            if (42 <= 65535) {
                                                                switch fi7 {
                                                                case 32768:
                                                                    var v48 bool = true
                                                                    _ = v48
                                                                case 0:
                                                                    if (42 <= 0) {
                                                                        if len(v4) > 0 {
                                                                            v4[0] = ""
                                                                        }
                                                                        v6 = v6
                                                                        switch ("a" + (("\n\t" + v35) + (v35 + (v35 + ("hello" + ""))))) {
                                                                        case "\n\t":
                                                                            if len(v4) > 0 {
                                                                                v4[0] = ("a" + ("a" + v35))
                                                                            }
                                                                        case "hello":
                                                                            v49 := GenS1{F0: -2147483648}
                                                                            _ = v49
                                                                        case "a":
                                                                            if false {
                                                                                if true {
                                                                                    _ = len(v1)
                                                                                    v50 := GenS4{F0: 42}
                                                                                    if false {
                                                                                        var v51 float64 = 0.0
                                                                                        v52 := GenS1{F0: 2147483647}
                                                                                        _ = v51
                                                                                        _ = v52
                                                                                    } else {
                                                                                        var v53 int16 = 1
                                                                                        var v54 int32 = 459
                                                                                        v6 += v6
                                                                                        var v55 float32 = (-2.25 + 1e10)
                                                                                        _ = v53
                                                                                        _ = v54
                                                                                        _ = v55
                                                                                    }
                                                                                    _ = v50
                                                                                }
                                                                                if (42 > 2) {
                                                                                    if false {
                                                                                        v6 += 1e10
                                                                                        var v56 uint = 2
                                                                                        if ((v56 & (v56 % (((v56 + 65535)) | 1))) <= 42) {
                                                                                            v57 := GenS4{F0: 442}
                                                                                            v1[951] = v35
                                                                                            _ = v57
                                                                                        } else {
                                                                                            v58 := GenS5{F0: 959}
                                                                                            switch ("\n\t" + "") {
                                                                                            case "":
                                                                                                var v59 int32 = 0
                                                                                                _ = v59
                                                                                            case "a":
                                                                                                var v60 int64 = 1
                                                                                                _ = v60
                                                                                            }
                                                                                            _ = v58
                                                                                        }
                                                                                        _ = v56
                                                                                    } else {
                                                                                        var v61 string = ""
                                                                                        _ = v61
                                                                                    }
                                                                                }
                                                                            } else {
                                                                                v47 += v47
                                                                            }
                                                                        }
                                                                    } else {
                                                                        if (!false) {
                                                                            switch (v47 - fi7) {
                                                                            case -9223372036854775808:
                                                                                _ = len(v4)
                                                                            }
                                                                            v6 += v6
                                                                            var v62 int32 = 0
                                                                            _ = v62
                                                                        } else {
                                                                            v6 -= float32((3.14159 - 0.1))
                                                                            if (false || true) {
                                                                                if true {
                                                                                    fmt.Println(v47, v6, fi5)
                                                                                    v63 := GenS1{F0: -1}
                                                                                    v6 = v6
                                                                                    _ = v63
                                                                                }
                                                                                v35 = "hello"
                                                                                v4 = append(v4, (v35 + v35))
                                                                                var v64 uint = 2
                                                                                v6 = v6
                                                                                _ = v64
                                                                            } else {
                                                                                if ((0.0 / 0.1) < 3.14159) {
                                                                                    var v65 int8 = 0
                                                                                    v66 := GenS4{F0: 32767}
                                                                                    v67 := GenS6{F0: (v35 + "a"), F1: 2, F2: 683}
                                                                                    _ = v65
                                                                                    _ = v66
                                                                                    _ = v67
                                                                                } else {
                                                                                    fmt.Println(v35, fi5)
                                                                                    v68 := GenS1{F0: 0}
                                                                                    _ = v68
                                                                                }
                                                                                v47 -= (v47 ^ fi5)
                                                                            }
                                                                        }
                                                                    }
                                                                }
                                                                v69 := GenS7{F0: -51}
                                                                v70 := GenS4{F0: 0}
                                                                _ = v69
                                                                _ = v70
                                                            }
                                                            if false {
                                                                v71 := v1[2147483648]
                                                                v72 := GenS8{F0: -54}
                                                                v73 := GenS8{F0: 100}
                                                                switch v35 {
                                                                case "foo bar":
                                                                    v74 := GenS9{F0: (fi5 >> ((255) & 63)), F1: 65535, F2: -86}
                                                                    _ = v74
                                                                case "":
                                                                    if (4294967295 == 2147483648) {
                                                                        fmt.Println(v35, v35, v71)
                                                                        v75 := GenS1{F0: 42}
                                                                        _ = v75
                                                                    }
                                                                }
                                                                _ = v71
                                                                _ = v72
                                                                _ = v73
                                                            } else {
                                                                switch 42 {
                                                                case 32768:
                                                                    v6 = float32(0.0)
                                                                case 32767:
                                                                    v6 = 0.1
                                                                case -9223372036854775808:
                                                                    if ((true || false) && (1 >= 715)) {
                                                                        v76 := v3[740]
                                                                        var v77 int16 = 877
                                                                        fmt.Println(v76)
                                                                        _ = v76
                                                                        _ = v77
                                                                    } else {
                                                                        switch 4294967295 {
                                                                        case -9223372036854775808:
                                                                            fmt.Println(v6)
                                                                        }
                                                                        var v78 uint32 = 2
                                                                        if (!(0.0 != ((-2.25 * -2.25) / ((((3.14159 / 0.1) - 1.5) / float64(v6)) / 1.0)))) {
                                                                            var v79 uint32 = (^(v78))
                                                                            v80 := GenS6{F0: v35, F1: 1, F2: 993}
                                                                            v79 = (v79 - 42)
                                                                            v81 := GenS3{F0: 831}
                                                                            v82 := GenS5{F0: 42}
                                                                            _ = v79
                                                                            _ = v80
                                                                            _ = v81
                                                                            _ = v82
                                                                        } else {
                                                                            if (3.14159 < 3.14159) {
                                                                                fmt.Println(v6)
                                                                                var v83 bool = false
                                                                                fmt.Println(fi5, v78, v78)
                                                                                v35 = ("foo bar" + v35)
                                                                                var v84 uint = 4294967295
                                                                                _ = v83
                                                                                _ = v84
                                                                            }
                                                                            if true {
                                                                                if false {
                                                                                    var v85 uint8 = 2
                                                                                    var v86 bool = (!(true || true))
                                                                                    switch v35 {
                                                                                    case "hello":
                                                                                        switch v35 {
                                                                                        case "":
                                                                                            v85 -= (^(v85))
                                                                                            switch ("hello" + (v35 + ((v35 + v35) + v35))) {
                                                                                            case "foo bar":
                                                                                                v47 -= (^(v47))
                                                                                            }
                                                                                        case "foo bar":
                                                                                            fmt.Println(fi7)
                                                                                        }
                                                                                    case "":
                                                                                        fmt.Println(fi5)
                                                                                    default:
                                                                                        fmt.Println(v78)
                                                                                    }
                                                                                    _ = v85
                                                                                    _ = v86
                                                                                }
                                                                                delete(v2, v35)
                                                                            }
                                                                        }
                                                                        _ = v78
                                                                    }
                                                                }
                                                            }
                                                            _ = v47
                                                        }
                                                    }
                                                } else {
                                                    if true {
                                                        v35 = ("hello" + "a")
                                                        var v87 int16 = 32767
                                                        _ = v87
                                                    } else {
                                                        switch (fi7 - fi7) {
                                                        case -1:
                                                            var v88 int16 = -1
                                                            _ = v88
                                                        case 42:
                                                            v89 := GenS10{F0: 1}
                                                            _ = v89
                                                        }
                                                        var v90 int16 = -1
                                                        if len(v4) > 0 {
                                                            v4[0] = (((v35 + "") + v35) + v35)
                                                        }
                                                        switch ((v35 + "foo bar") + v35) {
                                                        case "hello":
                                                            v90 = (v90 * (^(v90)))
                                                        case "foo bar":
                                                            fmt.Println(fi5, v90)
                                                        case "":
                                                            v4 = append(v4, (v35 + v35))
                                                        }
                                                        if true {
                                                            if (true || (!(!(-2147483648 == -976)))) {
                                                                var v91 bool = false
                                                                _ = v91
                                                            } else {
                                                                v92 := GenS6{F0: v35, F1: 48, F2: 4294967295}
                                                                var v93 int32 = -33
                                                                fmt.Println(v35, v93)
                                                                _ = v92
                                                                _ = v93
                                                            }
                                                            var v94 uint64 = 2
                                                            v35 = "hello"
                                                            _ = v94
                                                        }
                                                        _ = v90
                                                    }
                                                    var v95 int = 4294967295
                                                    if false {
                                                        var v96 int16 = 0
                                                        fmt.Println(fi7, fi7)
                                                        var v97 uint16 = 255
                                                        _ = v96
                                                        _ = v97
                                                    } else {
                                                        v6 += 0.0
                                                    }
                                                    _ = v95
                                                }
                                            }
                                        } else {
                                            v98 := GenS10{F0: 255}
                                            _ = v98
                                        }
                                        switch ("a" + "") {
                                        case "foo bar":
                                            if (573 >= 0) {
                                                if (true && false) {
                                                    v99 := GenS3{F0: -1}
                                                    v6 -= (float32(-2.25) - v6)
                                                    var v100 uint16 = 32767
                                                    _ = v100
                                                    _ = v99
                                                } else {
                                                    v101 := GenS11{F0: (true || (1 == -9223372036854775808)), F1: 1, F2: (!false)}
                                                    _ = v101
                                                }
                                                switch -9223372036854775808 {
                                                case 1:
                                                    v102 := GenS3{F0: -255}
                                                    _ = v102
                                                case -2147483648:
                                                    v6 = ((v6 - v6) + float32((1.5 * 0.0)))
                                                case -562:
                                                    v103 := GenS1{F0: -2147483648}
                                                    _ = v103
                                                default:
                                                    fmt.Println(v35)
                                                }
                                                v104 := GenS12{F0: 0, F1: 65535, F2: 72}
                                                v105 := GenS13{F0: -1}
                                                v35 = "\n\t"
                                                _ = v104
                                                _ = v105
                                            } else {
                                                var v106 uint64 = 406
                                                var v107 int8 = 0
                                                _ = v106
                                                _ = v107
                                            }
                                        case "a":
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                        case "":
                                            var v108 string = (v35 + v35)
                                            _ = v108
                                        }
                                    }
                                    if (v35 > "\n\t") {
                                        fmt.Println(fi5)
                                        v35 = (v35 + ((("\n\t" + v35) + "") + (("\n\t" + "\n\t") + (((v35 + v35) + (v35 + "\n\t")) + "foo bar"))))
                                        v109 := GenS14{F0: 1, F1: -15}
                                        _ = v109
                                    } else {
                                        v6 += ((v6 - v6) * v6)
                                        if (!(860 <= 2147483648)) {
                                            var v110 float64 = 0.0
                                            _ = v110
                                        }
                                    }
                                    fmt.Println(fi7, v35, fi5)
                                    if (0 == 32768) {
                                        fmt.Println(fi5)
                                        v35 = v35
                                    } else {
                                        v111 := GenS10{F0: 2}
                                        v6 = v6
                                        _ = v111
                                    }
                                }
                                if (100 != -128) {
                                    var v112 uint16 = 255
                                    if false {
                                        v112 = v112
                                        if len(v4) > 0 {
                                            v4[0] = "hello"
                                        }
                                    }
                                    switch "" {
                                    case "foo bar":
                                        delete(v1, 1)
                                    default:
                                        switch fi7 {
                                        case -1:
                                            v35 = ((v35 + v35) + ((((("foo bar" + v35) + "a") + (((v35 + v35) + v35) + (v35 + ((v35 + ((v35 + "hello") + "\n\t")) + ("hello" + v35))))) + (v35 + "foo bar")) + "a"))
                                        case 9223372036854775807:
                                            v35 = "hello"
                                        default:
                                            if false {
                                                v6 = float32(0.0)
                                            }
                                        }
                                    }
                                    _ = v112
                                } else {
                                    if (255 < 42) {
                                        var v113 uint = 65535
                                        v35 = ("hello" + (v35 + (v35 + v35)))
                                        _ = v113
                                    } else {
                                        v35 = v35
                                        fmt.Println(fi7, v6, fi5)
                                    }
                                    v35 = (v35 + v35)
                                }
                                v35 = ("\n\t" + v35)
                            }
                            v114 := GenS10{F0: 2}
                            _ = v114
                        }
                        _ = v35
                        _ = v36
                    }
                default:
                    switch ("a" + "") {
                    case "foo bar":
                        v1[65535] = ("hello" + ("\n\t" + "a"))
                    default:
                        if len(v4) > 0 {
                            v4[0] = (("a" + "a") + "")
                        }
                    }
                    if len(v4) > 0 {
                        v4[0] = "\n\t"
                    }
                }
            case "hello":
                switch fi5 {
                case 269:
                    fmt.Println(fi7)
                case -9223372036854775808:
                    v4 = append(v4, (("hello" + ("" + "hello")) + ("foo bar" + (("a" + ((("hello" + ("hello" + "a")) + "a") + "hello")) + ("a" + ("foo bar" + "foo bar"))))))
                case 0:
                    delete(v1, 2147483648)
                }
            }
            _ = len(v4)
            v6 = 3.14159
            _ = fi7
        }
        _ = fi5
        _ = v6
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if (true || true) {
        if (!true) {
            var v115 int64 = -189
            _ = v115
        } else {
            v116 := GenS4{F0: 255}
            var v117 int8 = 39
            _ = v116
            _ = v117
        }
        v118 := v2["a"]
        _ = v118
    } else {
        for fi119 := 0; fi119 < 17; fi119++ {
            if true {
                var v120 bool = (2 < 255)
                v121 := GenS15{F0: 2147483647}
                var v122 float32 = float32((((-2.25 + (0.0 + 0.0)) / 1e10) / 1.5))
                _ = v120
                _ = v121
                _ = v122
            } else {
                var v123 int64 = -557
                _ = v123
            }
            var v124 int8 = 50
            v125 := GenS8{F0: -122}
            _ = fi119
            _ = v124
            _ = v125
        }
        var v126 int32 = 1
        _ = v126
    }
    var v127 string = ""
    if (33 != 42) {
        if false {
            for fi128 := 0; fi128 < 15; fi128++ {
                v129 := GenS16{F0: v127, F1: (fi128 >> ((2) & 63))}
                switch 1 {
                case 42:
                    if false {
                        switch v127 {
                        case "":
                            if len(v4) > 0 {
                                v4[0] = v127
                            }
                        case "a":
                            var v130 bool = ((!false) || true)
                            _ = v130
                        case "foo bar":
                            v127 = ((v127 + "") + "hello")
                        default:
                            var v131 int16 = 1
                            _ = v131
                        }
                    }
                case 4294967295:
                    v132 := GenS17{F0: 128, F1: 433, F2: 255}
                    _ = v132
                case 211:
                    fmt.Println(fi128, fi128)
                default:
                    var v133 int8 = int8(fi128)
                    _ = v133
                }
                for fi134 := 0; fi134 < 17; fi134++ {
                    switch fi134 {
                    case -531:
                        switch (v127 + ((v127 + "") + v127)) {
                        case "a":
                            var v135 bool = ((42 >= 42) && false)
                            _ = v135
                        }
                    case 32767:
                        v136 := GenS6{F0: "", F1: 690, F2: 404}
                        _ = v136
                    }
                    v127 = (("hello" + v127) + v127)
                    _ = fi134
                }
                _ = fi128
                _ = v129
            }
            v137 := GenS2{F0: -69, F1: 110, F2: 2147483647}
            _ = v137
        } else {
            if len(v4) > 0 {
                v4[0] = v127
            }
            v4 = append(v4, v127)
            v127 = v127
        }
    } else {
        switch (v127 + "\n\t") {
        case "a":
            fmt.Println(v127)
        case "":
            if (360 != 42) {
                v127 = "hello"
                if true {
                    v138 := GenS10{F0: 4294967295}
                    for fi139 := 0; fi139 < 13; fi139++ {
                        if false {
                            if false {
                                var v140 int8 = 42
                                v1[271] = v127
                                v127 = ("foo bar" + (v127 + (v127 + (("\n\t" + "foo bar") + "foo bar"))))
                                _ = v140
                            } else {
                                var v141 bool = true
                                v141 = v141
                                for fi142 := 0; fi142 < 15; fi142++ {
                                    v141 = (-509 != 0)
                                    switch v127 {
                                    case "":
                                        v143 := GenS18{F0: (fi139 ^ (^(fi139))), F1: (v141 && v141)}
                                        _ = v143
                                    case "a":
                                        fmt.Println(fi139)
                                    case "\n\t":
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                    }
                                    v144 := GenS13{F0: 1}
                                    var v145 bool = (-128 < 49)
                                    if ((!v141) && false) {
                                        v146 := GenS10{F0: 255}
                                        v141 = v145
                                        var v147 uint8 = 2
                                        _ = v146
                                        _ = v147
                                    }
                                    v141 = v145
                                    v148 := v2["foo bar"]
                                    _ = fi142
                                    _ = v144
                                    _ = v145
                                    _ = v148
                                }
                                _ = v141
                            }
                            var v149 int16 = 758
                            _ = v149
                        } else {
                            var v150 bool = true
                            if (!false) {
                                switch (v127 + (v127 + v127)) {
                                case "foo bar":
                                    switch fi139 {
                                    case 2147483647:
                                        for fi151 := 0; fi151 < 18; fi151++ {
                                            var v152 uint64 = 1
                                            if v150 {
                                                v153 := v3[42]
                                                delete(v2, ("foo bar" + v127))
                                                v152 = v152
                                                _ = v153
                                            } else {
                                                if false {
                                                    _ = len(v4)
                                                    if true {
                                                        v150 = v150
                                                        _ = len(v1)
                                                    }
                                                    _ = len(v4)
                                                }
                                                if true {
                                                    v152 = 2
                                                    var v154 uint16 = 32767
                                                    _ = v154
                                                } else {
                                                    var v155 int8 = 127
                                                    _ = v155
                                                }
                                                v156 := GenS19{F0: 255}
                                                _ = v156
                                            }
                                            fmt.Println(fi151, fi139)
                                            _ = fi151
                                            _ = v152
                                        }
                                    case 1:
                                        if (855 < 32767) {
                                            fmt.Println(v127, fi139, v150)
                                            if v150 {
                                                if v150 {
                                                    switch fi139 {
                                                    case -1:
                                                        var v157 bool = (!v150)
                                                        _ = v157
                                                    }
                                                } else {
                                                    var v158 bool = (486 <= 32767)
                                                    var v159 bool = (-1 >= 0)
                                                    for fi160 := 0; fi160 < 9; fi160++ {
                                                        v158 = v158
                                                        _ = fi160
                                                    }
                                                    _ = v158
                                                    _ = v159
                                                }
                                                for fi161 := 0; fi161 < 9; fi161++ {
                                                    v162 := GenS20{F0: 18446744073709551615, F1: (fi139 >> ((255) & 63))}
                                                    _ = len(v4)
                                                    _ = fi161
                                                    _ = v162
                                                }
                                                if (v150 || (2 != 924)) {
                                                    var v163 int16 = -32768
                                                    var v164 uint64 = 0
                                                    _ = v163
                                                    _ = v164
                                                }
                                            }
                                        } else {
                                            if v150 {
                                                v127 = (v127 + "a")
                                                v4 = append(v4, v127)
                                                for fi165 := 0; fi165 < 7; fi165++ {
                                                    fmt.Println(v127)
                                                    var v166 int16 = 42
                                                    var v167 string = ("" + "")
                                                    _ = fi165
                                                    _ = v166
                                                    _ = v167
                                                }
                                            } else {
                                                v127 = "foo bar"
                                            }
                                            var v168 string = "a"
                                            v150 = true
                                            _ = v168
                                        }
                                    case 128:
                                        v169 := GenS21{F0: 127, F1: 2147483647, F2: 100}
                                        _ = v169
                                    }
                                case "a":
                                    v127 = (v127 + (v127 + ("hello" + "")))
                                case "hello":
                                    if (527 >= 552) {
                                        for fi170 := 0; fi170 < 16; fi170++ {
                                            switch "foo bar" {
                                            case "hello":
                                                switch "foo bar" {
                                                case "a":
                                                    v127 = (v127 + (v127 + v127))
                                                case "hello":
                                                    delete(v2, ("foo bar" + v127))
                                                }
                                                var v171 uint = 495
                                                v171 = (v171 << ((v171) & 63))
                                                var v172 int16 = 265
                                                var v173 uint64 = 1
                                                _ = v171
                                                _ = v172
                                                _ = v173
                                            }
                                            var v174 uint = 4294967295
                                            fmt.Println(fi170)
                                            _ = fi170
                                            _ = v174
                                        }
                                    }
                                }
                            } else {
                                if v150 {
                                    if (826 <= 536) {
                                        switch v127 {
                                        case "\n\t":
                                            fmt.Println(v127)
                                        case "":
                                            v127 = "\n\t"
                                        }
                                        if v150 {
                                            v127 = ""
                                        } else {
                                            switch "hello" {
                                            case "foo bar":
                                                v175 := GenS22{F0: 2}
                                                _ = v175
                                            case "\n\t":
                                                _ = len(v2)
                                            }
                                        }
                                        if (!v150) {
                                            var v176 int8 = -128
                                            for fi177 := 0; fi177 < 18; fi177++ {
                                                var v178 int = fi177
                                                _ = fi177
                                                _ = v178
                                            }
                                            switch ("\n\t" + (v127 + v127)) {
                                            case "foo bar":
                                                var v179 string = v127
                                                _ = v179
                                            case "\n\t":
                                                v150 = true
                                            }
                                            _ = v176
                                        }
                                    } else {
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                    }
                                } else {
                                    fmt.Println(v127)
                                    v180 := GenS13{F0: -1}
                                    var v181 uint64 = 0
                                    _ = v180
                                    _ = v181
                                }
                                if v150 {
                                    fmt.Println(v127, v150)
                                    var v182 uint8 = 1
                                    if (0.0 < 0.0) {
                                        var v183 bool = ((float64(-2.25) * 3.14159) < (0.0 / 0.1))
                                        for fi184 := 0; fi184 < 11; fi184++ {
                                            if (!(1 >= 964)) {
                                                v4 = append(v4, "a")
                                            } else {
                                                var v185 float32 = 0.1
                                                var v186 float32 = (v185 + float32(0.0))
                                                _ = v185
                                                _ = v186
                                            }
                                            v182 += (v182 ^ (v182 & v182))
                                            var v187 uint16 = 1
                                            _ = fi184
                                            _ = v187
                                        }
                                        v1[2147483647] = "a"
                                        _ = v183
                                    }
                                    _ = v182
                                }
                            }
                            for fi188 := 0; fi188 < 19; fi188++ {
                                var v189 int32 = -834
                                _ = fi188
                                _ = v189
                            }
                            _ = v150
                        }
                        for fi190 := 0; fi190 < 17; fi190++ {
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            _ = len(v2)
                            _ = fi190
                        }
                        _ = fi139
                    }
                    _ = v138
                } else {
                    var v191 uint16 = 959
                    _ = v191
                }
            } else {
                var v192 bool = true
                _ = v192
            }
        }
        var v193 int16 = 1000
        v194 := GenS23{F0: -796, F1: 255}
        _ = v193
        _ = v194
    }
    switch 0 {
    case 174:
        v127 = v127
    case -503:
        v127 = v127
    case -1:
        fmt.Println(v127)
    }
    for fi195 := 0; fi195 < 19; fi195++ {
        _ = len(v4)
        _ = fi195
    }
    delete(v2, v127)
    for fi196 := 0; fi196 < 17; fi196++ {
        _ = len(v4)
        fmt.Println(fi196, fi196)
        for fi197 := 0; fi197 < 2; fi197++ {
            var v198 string = ("foo bar" + v127)
            _ = fi197
            _ = v198
        }
        _ = fi196
    }
    v127 = v127
    v127 = "a"
    v127 = v127
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v127:", v127)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v127
    _ = v2
    _ = v3
    _ = v4
}
