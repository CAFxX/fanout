package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool; F2 uint16 }
type GenS2 struct { F0 int32; F1 string }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 int; F1 uint8 }
type GenS8 struct { F0 string }
type GenS9 struct { F0 string; F1 int16 }
type GenS10 struct { F0 bool }

func main() {
    v1 := []string{}
    v2 := make(map[int]bool)
    v3 := make(map[int]uint32)
    v4 := make(map[string]uint32)
    var v5 uint16 = 42
    var v6 uint16 = 2
    fmt.Println(v6, v6)
    v5 = 98
    fmt.Println(v5, v6)
    if false {
        var v7 int8 = 0
        _ = v7
    } else {
        var v8 int8 = 42
        v9 := GenS1{F0: 719, F1: (!(3.14159 == (-2.25 * (1e10 - (3.14159 * 1.5))))), F2: (v5 % ((v6) | 1))}
        _ = v8
        _ = v9
    }
    if (19 != 42) {
        var v10 uint64 = 0
        if true {
            v10 = v10
            if true {
                for fi11 := 0; fi11 < 4; fi11++ {
                    if (0 < 1) {
                        var v12 uint8 = 42
                        _ = v12
                    }
                    v10 = 18446744073709551615
                    _ = fi11
                }
                switch 32767 {
                case 42:
                    v10 += 18446744073709551615
                case 0:
                    v10 = v10
                default:
                    var v13 uint16 = 32767
                    _ = v13
                }
            } else {
                for fi14 := 0; fi14 < 4; fi14++ {
                    v4["\n\t"] = 983
                    var v15 int16 = -32768
                    var v16 string = "hello"
                    _ = fi14
                    _ = v15
                    _ = v16
                }
                v17 := GenS1{F0: (v10 & 546), F1: true, F2: (v5 * (v6 >> ((1) & 63)))}
                _ = v17
            }
            switch 4294967295 {
            case -1:
                delete(v2, 42)
            }
        } else {
            var v18 int16 = -32768
            _ = v18
        }
        _ = v10
    }
    if false {
        var v19 int64 = 0
        var v20 uint32 = 1
        _ = v19
        _ = v20
    }
    switch ("\n\t" + "\n\t") {
    case "":
        v21 := GenS2{F0: -2147483648, F1: ("hello" + ("hello" + "\n\t"))}
        _ = v21
    case "\n\t":
        var v22 uint32 = 4294967295
        _ = v22
    }
    for fi23 := 0; fi23 < 4; fi23++ {
        if (992 == 2147483647) {
            switch "\n\t" {
            case "":
                fmt.Println(v6, fi23, v5)
            case "foo bar":
                switch (((("\n\t" + "foo bar") + "hello") + "hello") + "\n\t") {
                case "":
                    var v24 string = "foo bar"
                    _ = v24
                case "hello":
                    v25 := GenS1{F0: 73, F1: (1 >= 248), F2: (v6 ^ 702)}
                    _ = v25
                case "foo bar":
                    if false {
                        fmt.Println(fi23)
                        _ = len(v1)
                    } else {
                        v6 = 1
                    }
                default:
                    var v26 int8 = 0
                    _ = v26
                }
            case "hello":
                if ((42 >= 18446744073709551615) || (0 > 37)) {
                    if true {
                        v27 := GenS1{F0: 4294967295, F1: false, F2: v5}
                        _ = v27
                    }
                } else {
                    for fi28 := 0; fi28 < 11; fi28++ {
                        v6 -= (v6 % ((65535) | 1))
                        if false {
                            if false {
                                var v29 int64 = -145
                                _ = v29
                            } else {
                                if true {
                                    var v30 bool = false
                                    v1 = append(v1, "")
                                    _ = v30
                                }
                                if true {
                                    if true {
                                        if false {
                                            var v31 uint = 1
                                            if true {
                                                if (42 > 114) {
                                                    switch fi23 {
                                                    case -489:
                                                        fmt.Println(v5, v31, v31)
                                                        v3[fi23] = 2147483647
                                                    case 32767:
                                                        if (571 == 0) {
                                                            var v32 float64 = float64(((-2.25 - 1e10) / -2.25))
                                                            v33 := GenS2{F0: 380, F1: "foo bar"}
                                                            _ = v32
                                                            _ = v33
                                                        }
                                                    }
                                                } else {
                                                    var v34 bool = true
                                                    _ = v34
                                                }
                                                var v35 uint = (^(v31))
                                                _ = v35
                                            } else {
                                                v31 = v31
                                                v5 = v5
                                                v36 := GenS2{F0: -1, F1: ("foo bar" + "a")}
                                                _ = v36
                                            }
                                            var v37 float64 = (float64((((((1.5 + float32(((3.14159 / 3.14159) * -2.25))) * 0.0) + 1.5) / ((3.14159 * 0.0) - (0.1 * 1.5))) - 1e10)) * (0.0 + 1e10))
                                            fmt.Println(v37)
                                            var v38 uint8 = uint8(v6)
                                            _ = v31
                                            _ = v37
                                            _ = v38
                                        } else {
                                            _ = len(v1)
                                            v39 := GenS1{F0: 4294967295, F1: (!((32767 > -51) && (1 > 65535))), F2: (^(v5))}
                                            var v40 int64 = 134
                                            v6 = 255
                                            _ = v39
                                            _ = v40
                                        }
                                    } else {
                                        var v41 float64 = ((0.0 + float64(0.0)) - 1e10)
                                        v6 = (^(v6))
                                        _ = v41
                                    }
                                } else {
                                    if false {
                                        if true {
                                            if true {
                                                var v42 int64 = -9223372036854775808
                                                v5 = 42
                                                if (false || false) {
                                                    v43 := GenS3{F0: false}
                                                    _ = v43
                                                } else {
                                                    var v44 uint64 = 2
                                                    v45 := GenS3{F0: ("\n\t" >= "\n\t")}
                                                    if true {
                                                        switch -9223372036854775808 {
                                                        case 2147483647:
                                                            v44 = v44
                                                        case 4294967295:
                                                            _ = len(v2)
                                                        default:
                                                            v44 = v44
                                                        }
                                                        var v46 int16 = 127
                                                        _ = v46
                                                    } else {
                                                        if true {
                                                            if true {
                                                                var v47 int32 = 1
                                                                _ = v47
                                                            }
                                                            v4["foo bar"] = 770
                                                            v6 = (^(v6))
                                                        } else {
                                                            v3[(fi28 * -413)] = 4294967295
                                                            if (0 > 406) {
                                                                fmt.Println(v44)
                                                            }
                                                            var v48 uint16 = (v5 >> ((950) & 63))
                                                            _ = v48
                                                        }
                                                    }
                                                    _ = v44
                                                    _ = v45
                                                }
                                                var v49 float64 = -2.25
                                                var v50 float32 = 0.1
                                                if (("\n\t" + "a") < ("\n\t" + "\n\t")) {
                                                    var v51 int16 = -32768
                                                    if true {
                                                        v51 = (-(v51))
                                                    }
                                                    var v52 uint64 = 42
                                                    _ = v51
                                                    _ = v52
                                                } else {
                                                    var v53 int = fi23
                                                    var v54 bool = (v50 >= -2.25)
                                                    _ = v53
                                                    _ = v54
                                                }
                                                var v55 int64 = v42
                                                _ = v42
                                                _ = v49
                                                _ = v50
                                                _ = v55
                                            }
                                        }
                                    } else {
                                        v56 := v4["a"]
                                        _ = v56
                                    }
                                    if false {
                                        if (317 != 2147483647) {
                                            fmt.Println(fi28, v5)
                                            var v57 string = "\n\t"
                                            _ = v57
                                        } else {
                                            fmt.Println(v6, fi28, v6)
                                        }
                                        switch fi28 {
                                        case 1:
                                            v58 := GenS4{F0: fi23}
                                            _ = v58
                                        case 604:
                                            var v59 int32 = 1
                                            _ = v59
                                        case -9223372036854775808:
                                            if (!(0 >= 1)) {
                                                if true {
                                                    v60 := v3[fi28]
                                                    var v61 uint16 = v6
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                    _ = v60
                                                    _ = v61
                                                } else {
                                                    var v62 uint = 681
                                                    _ = v62
                                                }
                                                fmt.Println(fi23)
                                                v63 := GenS5{F0: -9223372036854775808}
                                                _ = v63
                                            }
                                        }
                                    }
                                    var v64 int8 = 1
                                    _ = v64
                                }
                                var v65 uint16 = v6
                                var v66 int16 = 1000
                                v66 = (v66 + 1)
                                switch 626 {
                                case -788:
                                    var v67 float64 = float64(1.5)
                                    _ = v67
                                case -503:
                                    if false {
                                        v68 := GenS1{F0: 18446744073709551615, F1: true, F2: 375}
                                        switch "foo bar" {
                                        case "a":
                                            switch "a" {
                                            case "foo bar":
                                                var v69 uint = 18446744073709551615
                                                _ = v69
                                            case "":
                                                fmt.Println(v5, v6, fi23)
                                            case "a":
                                                var v70 int16 = 105
                                                _ = v70
                                            }
                                        }
                                        if (((127 >= 42) || (2147483647 >= 4294967295)) && false) {
                                            switch "foo bar" {
                                            case "foo bar":
                                                var v71 string = ("hello" + "")
                                                _ = v71
                                            }
                                        } else {
                                            v66 -= (v66 & v66)
                                            v6 = (v6 ^ (^(v65)))
                                            if (!(255 > 65535)) {
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                                if false {
                                                    var v72 int32 = 0
                                                    _ = v72
                                                }
                                            } else {
                                                var v73 string = ("a" + ("" + "\n\t"))
                                                fmt.Println(v5)
                                                var v74 uint = 372
                                                if true {
                                                    if false {
                                                        if (!true) {
                                                            var v75 uint32 = uint32(v74)
                                                            v76 := GenS4{F0: 517}
                                                            v77 := GenS2{F0: -411, F1: (v73 + "\n\t")}
                                                            _ = v75
                                                            _ = v76
                                                            _ = v77
                                                        }
                                                        if true {
                                                            fmt.Println(v66)
                                                            if (127 < 1) {
                                                                v78 := GenS6{F0: fi23}
                                                                v79 := v4["hello"]
                                                                v80 := v3[fi23]
                                                                _ = v78
                                                                _ = v79
                                                                _ = v80
                                                            } else {
                                                                var v81 int32 = 1
                                                                _ = len(v1)
                                                                fmt.Println(fi28, v5, fi23)
                                                                _ = v81
                                                            }
                                                        }
                                                        fmt.Println(fi23)
                                                        v5 = (^(v65))
                                                        var v82 uint16 = v6
                                                        _ = v82
                                                    } else {
                                                        v65 -= (v5 & v5)
                                                        var v83 uint8 = 128
                                                        if len(v1) > 0 {
                                                            v1[0] = v73
                                                        }
                                                        if ((false && (4294967295 < 18446744073709551615)) || true) {
                                                            fmt.Println(v83)
                                                            if true {
                                                                switch v73 {
                                                                case "hello":
                                                                    if true {
                                                                        var v84 uint64 = 42
                                                                        _ = v84
                                                                    } else {
                                                                        v5 += v65
                                                                        var v85 uint32 = 875
                                                                        var v86 int8 = int8(v85)
                                                                        _ = v85
                                                                        _ = v86
                                                                    }
                                                                case "a":
                                                                    if (!true) {
                                                                        fmt.Println(v74, v73)
                                                                        v6 -= (v65 | 32767)
                                                                    }
                                                                }
                                                                var v87 uint = v74
                                                                v88 := GenS7{F0: fi23, F1: v83}
                                                                _ = v87
                                                                _ = v88
                                                            }
                                                        } else {
                                                            var v89 float32 = (1.5 / 1.5)
                                                            fmt.Println(v65)
                                                            v90 := GenS5{F0: 1}
                                                            if false {
                                                                v65 = v5
                                                            } else {
                                                                if (!true) {
                                                                    v66 -= v66
                                                                }
                                                                v74 += (^(v74))
                                                                v91 := GenS8{F0: (v73 + v73)}
                                                                var v92 int64 = 115
                                                                v93 := GenS8{F0: "a"}
                                                                _ = v91
                                                                _ = v92
                                                                _ = v93
                                                            }
                                                            _ = v89
                                                            _ = v90
                                                        }
                                                        var v94 uint16 = uint16(v74)
                                                        _ = v83
                                                        _ = v94
                                                    }
                                                    v65 -= (v5 % ((1) | 1))
                                                    v66 += int16(v74)
                                                    switch v73 {
                                                    case "":
                                                        var v95 uint = v74
                                                        _ = v95
                                                    case "hello":
                                                        v96 := GenS5{F0: -1}
                                                        _ = v96
                                                    case "foo bar":
                                                        fmt.Println(fi28)
                                                    }
                                                    var v97 int64 = 0
                                                    _ = v97
                                                }
                                                _ = v73
                                                _ = v74
                                            }
                                        }
                                        _ = v68
                                    } else {
                                        if (34 == 0) {
                                            switch "hello" {
                                            case "hello":
                                                if len(v1) > 0 {
                                                    v1[0] = ""
                                                }
                                            default:
                                                _ = len(v1)
                                                var v98 int32 = -966
                                                v99 := GenS6{F0: -9223372036854775808}
                                                v6 += (^(v5))
                                                _ = v98
                                                _ = v99
                                            }
                                        }
                                        fmt.Println(v66, v5, v65)
                                        v100 := GenS9{F0: (((("" + ("foo bar" + "\n\t")) + "hello") + "\n\t") + (("" + "") + "hello")), F1: v66}
                                        _ = v100
                                    }
                                    switch 42 {
                                    case 128:
                                        if (true || true) {
                                            v6 += (^(v6))
                                        } else {
                                            v5 += (v5 >> ((4294967295) & 63))
                                            fmt.Println(v65, fi23)
                                        }
                                    case -299:
                                        var v101 int64 = 612
                                        _ = v101
                                    }
                                    var v102 int16 = int16(v65)
                                    _ = v102
                                }
                                fmt.Println(v5, v6, v5)
                                switch "hello" {
                                case "\n\t":
                                    var v103 int32 = 1
                                    _ = len(v4)
                                    _ = v103
                                case "hello":
                                    fmt.Println(v6, fi28)
                                default:
                                    var v104 int32 = 42
                                    _ = v104
                                }
                                _ = v65
                                _ = v66
                            }
                            v5 += v6
                            delete(v2, (fi28 << ((2) & 63)))
                        } else {
                            v6 = v6
                        }
                        var v105 uint32 = uint32(v6)
                        _ = fi28
                        _ = v105
                    }
                }
            default:
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
            v5 = (v5 >> ((243) & 63))
        }
        _ = fi23
    }
    if true {
        if len(v1) > 0 {
            v1[0] = "\n\t"
        }
        switch "a" {
        case "\n\t":
            for fi106 := 0; fi106 < 19; fi106++ {
                v6 -= 40
                var v107 uint64 = 42
                _ = fi106
                _ = v107
            }
        case "foo bar":
            v6 += v6
        }
        var v108 string = "foo bar"
        _ = v108
    } else {
        for fi109 := 0; fi109 < 18; fi109++ {
            var v110 int16 = 42
            v111 := GenS7{F0: (fi109 - fi109), F1: 127}
            for fi112 := 0; fi112 < 5; fi112++ {
                v110 = v110
                v113 := GenS10{F0: true}
                _ = fi112
                _ = v113
            }
            _ = fi109
            _ = v110
            _ = v111
        }
    }
    var v114 int16 = 32767
    switch 135 {
    case -9223372036854775808:
        var v115 uint8 = 127
        _ = v115
    case -186:
        switch "foo bar" {
        case "hello":
            v6 += (v5 + v5)
        case "foo bar":
            fmt.Println(v114)
        }
    }
    if false {
        var v116 float64 = 1.5
        if true {
            v117 := GenS4{F0: 280}
            _ = v117
        }
        if true {
            v6 += 65535
            fmt.Println(v5)
        }
        _ = v116
    } else {
        v118 := GenS1{F0: 18446744073709551615, F1: true, F2: (v5 ^ v6)}
        if true {
            fmt.Println(v114)
        }
        if (true && (v5 > 356)) {
            v6 = 255
            for fi119 := 0; fi119 < 10; fi119++ {
                fmt.Println(v5)
                _ = fi119
            }
        }
        _ = v118
    }
    var v120 uint64 = 0
    fmt.Println(v5, v114, v114)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v114:", v114)
    fmt.Println("v120:", v120)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v114
    _ = v120
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
