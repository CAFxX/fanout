package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int }
type GenS2 struct { F0 uint8; F1 bool }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 string; F1 uint32 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint64; F1 uint32; F2 uint8 }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 int; F1 uint; F2 int16 }
type GenS10 struct { F0 uint16; F1 uint64; F2 string }

func main() {
    var v1 float32 = (0.1 + 0.0)
    v2 := []string{"\n\t", "a", "\n\t"}
    v3 := []string{""}
    var v4 uint32 = 2147483648
    v4 -= (v4 * v4)
    for fi5 := 0; fi5 < 16; fi5++ {
        v6 := GenS1{F0: (v4 - (v4 >> ((2) & 63))), F1: 128}
        _ = fi5
        _ = v6
    }
    v7 := GenS1{F0: (v4 & 42), F1: 4294967295}
    v4 += v4
    var v8 int32 = 0
    if (128 >= 1) {
        if (-1 == -987) {
            for fi9 := 0; fi9 < 8; fi9++ {
                if true {
                    var v10 int = (fi9 % ((9223372036854775807) | 1))
                    if (255 > 199) {
                        if ((18446744073709551615 > 2) || (!(923 == 0))) {
                            if ((0 == 245) && false) {
                                for fi11 := 0; fi11 < 3; fi11++ {
                                    v4 = (v4 * 0)
                                    if len(v3) > 0 {
                                        v3[0] = "a"
                                    }
                                    v8 += v8
                                    _ = fi11
                                }
                                if false {
                                    v12 := GenS1{F0: v4, F1: (fi9 * v10)}
                                    _ = v12
                                } else {
                                    for fi13 := 0; fi13 < 20; fi13++ {
                                        var v14 uint8 = 255
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        _ = fi13
                                        _ = v14
                                    }
                                    _ = len(v2)
                                    v8 = (v8 * 2147483647)
                                }
                            }
                            var v15 int32 = 42
                            v8 -= (v8 % ((v8) | 1))
                            _ = v15
                        } else {
                            v4 -= (v4 / ((v4) | 1))
                        }
                    } else {
                        v4 += 128
                    }
                    _ = v10
                } else {
                    v8 = -1
                    switch fi9 {
                    case -1:
                        var v16 int32 = 42
                        _ = v16
                    case 286:
                        var v17 int32 = (^(v8))
                        _ = v17
                    }
                }
                var v18 uint8 = 0
                var v19 uint8 = 149
                _ = fi9
                _ = v18
                _ = v19
            }
        } else {
            var v20 float32 = -2.25
            if true {
                switch "foo bar" {
                case "hello":
                    fmt.Println(v1, v4)
                default:
                    v20 += v20
                }
                for fi21 := 0; fi21 < 18; fi21++ {
                    v22 := GenS2{F0: 2, F1: false}
                    if (v1 >= v20) {
                        var v23 int32 = v8
                        if true {
                            v4 = (v4 & v4)
                            if false {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            }
                            var v24 int64 = 0
                            _ = v24
                        } else {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            v25 := GenS1{F0: v4, F1: (fi21 << ((4294967295) & 63))}
                            _ = v25
                        }
                        _ = v23
                    } else {
                        var v26 int8 = 100
                        var v27 bool = (v4 <= v4)
                        if false {
                            v4 -= v4
                        } else {
                            if (false || (!v27)) {
                                for fi28 := 0; fi28 < 19; fi28++ {
                                    if v27 {
                                        if v27 {
                                            var v29 float64 = 1.5
                                            _ = v29
                                        }
                                        v30 := GenS2{F0: 164, F1: (float64(v20) > -2.25)}
                                        if (!((^(fi21)) != (fi21 >> ((4294967295) & 63)))) {
                                            var v31 int32 = (v8 % ((v8) | 1))
                                            if false {
                                                var v32 int = -832
                                                _ = v32
                                            } else {
                                                if (!false) {
                                                    var v33 int8 = (^(v26))
                                                    var v34 float64 = 1e10
                                                    _ = v33
                                                    _ = v34
                                                } else {
                                                    _ = len(v2)
                                                    v35 := GenS3{F0: 4294967295}
                                                    if len(v3) > 0 {
                                                        v3[0] = "foo bar"
                                                    }
                                                    if (v27 || (false && false)) {
                                                        v36 := GenS3{F0: 4294967295}
                                                        if v27 {
                                                            if v27 {
                                                                if false {
                                                                    var v37 int32 = 42
                                                                    v38 := GenS2{F0: 127, F1: v27}
                                                                    _ = v37
                                                                    _ = v38
                                                                } else {
                                                                    v39 := GenS4{F0: "foo bar", F1: (v4 / ((v4) | 1))}
                                                                    _ = v39
                                                                }
                                                                switch fi28 {
                                                                case 32768:
                                                                    var v40 float32 = v1
                                                                    _ = v40
                                                                }
                                                                switch -9223372036854775808 {
                                                                case 32767:
                                                                    var v41 string = "a"
                                                                    _ = v41
                                                                }
                                                            } else {
                                                                if false {
                                                                    v42 := GenS5{F0: -794}
                                                                    v31 += v8
                                                                    v20 = (-2.25 / v20)
                                                                    _ = v42
                                                                } else {
                                                                    switch "" {
                                                                    case "\n\t":
                                                                        v20 -= (float32((((1e10 / 1.0) / (-2.25 - 1e10)) + 3.14159)) / v20)
                                                                    default:
                                                                        switch (fi21 >> ((42) & 63)) {
                                                                        case 128:
                                                                            var v43 bool = true
                                                                            _ = v43
                                                                        case 1:
                                                                            if v27 {
                                                                                v8 = v8
                                                                            } else {
                                                                                switch "\n\t" {
                                                                                case "foo bar":
                                                                                    if true {
                                                                                        var v44 int = -862
                                                                                        _ = v44
                                                                                    } else {
                                                                                        if len(v2) > 0 {
                                                                                            v2[0] = "\n\t"
                                                                                        }
                                                                                    }
                                                                                case "":
                                                                                    if false {
                                                                                        v45 := GenS6{F0: 1, F1: 2147483647, F2: 2}
                                                                                        var v46 int64 = 9223372036854775807
                                                                                        var v47 uint = 1
                                                                                        _ = v45
                                                                                        _ = v46
                                                                                        _ = v47
                                                                                    }
                                                                                default:
                                                                                    v26 += 42
                                                                                }
                                                                            }
                                                                        case 42:
                                                                            if true {
                                                                                if (("" + "") > ("hello" + "foo bar")) {
                                                                                    if (!(255 > 255)) {
                                                                                        v4 = v4
                                                                                    }
                                                                                    if v27 {
                                                                                        fmt.Println(v26, fi21, v20)
                                                                                        v1 -= ((((0.1 * v1) + 1e10) * float32(-2.25)) + float32(((0.1 / 3.14159) + 1.5)))
                                                                                        v8 = v31
                                                                                    } else {
                                                                                        if (false && v27) {
                                                                                            v2 = append(v2, ("hello" + "\n\t"))
                                                                                        }
                                                                                        var v48 uint64 = 239
                                                                                        _ = v48
                                                                                    }
                                                                                    var v49 uint32 = (v4 | (v4 % ((v4) | 1)))
                                                                                    var v50 uint64 = uint64(v49)
                                                                                    v1 += 3.14159
                                                                                    switch ("\n\t" + ("foo bar" + "\n\t")) {
                                                                                    case "a":
                                                                                        var v51 int16 = 42
                                                                                        _ = v51
                                                                                    case "hello":
                                                                                        fmt.Println(v31, fi28, v4)
                                                                                    case "foo bar":
                                                                                        if (!false) {
                                                                                            v49 = (v49 - (v49 | (^(v4))))
                                                                                            if ((^(v50)) == (v50 ^ (v50 * v50))) {
                                                                                                v52 := GenS5{F0: 1}
                                                                                                v53 := GenS1{F0: (^(v49)), F1: (fi21 | 9223372036854775807)}
                                                                                                _ = v52
                                                                                                _ = v53
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                    _ = v49
                                                                                    _ = v50
                                                                                }
                                                                                if len(v2) > 0 {
                                                                                    v2[0] = "\n\t"
                                                                                }
                                                                            } else {
                                                                                v8 += v8
                                                                                v20 += -2.25
                                                                                var v54 uint = 4294967295
                                                                                var v55 uint16 = 255
                                                                                _ = v54
                                                                                _ = v55
                                                                            }
                                                                        }
                                                                    }
                                                                    var v56 int32 = int32(v4)
                                                                    _ = v56
                                                                }
                                                                v20 -= 0.1
                                                                v57 := GenS2{F0: 1, F1: v27}
                                                                switch "a" {
                                                                case "hello":
                                                                    var v58 string = (("" + "hello") + "foo bar")
                                                                    _ = v58
                                                                default:
                                                                    if true {
                                                                        var v59 uint = 0
                                                                        var v60 string = ""
                                                                        _ = v59
                                                                        _ = v60
                                                                    }
                                                                }
                                                                fmt.Println(fi21)
                                                                v61 := GenS7{F0: 255}
                                                                if (2 > 1) {
                                                                    _ = len(v3)
                                                                    var v62 int = (fi28 & -1)
                                                                    _ = v62
                                                                }
                                                                _ = v57
                                                                _ = v61
                                                            }
                                                        }
                                                        _ = v36
                                                    }
                                                    _ = v35
                                                }
                                                v20 += 1e10
                                                v20 = v1
                                                var v63 float64 = -2.25
                                                var v64 int8 = (^(v26))
                                                _ = v63
                                                _ = v64
                                            }
                                            _ = v31
                                        }
                                        _ = v30
                                    } else {
                                        _ = len(v3)
                                    }
                                    var v65 uint32 = (v4 >> ((42) & 63))
                                    _ = fi28
                                    _ = v65
                                }
                            } else {
                                var v66 float64 = (0.0 - 0.0)
                                _ = v66
                            }
                        }
                        _ = v26
                        _ = v27
                    }
                    _ = fi21
                    _ = v22
                }
                if false {
                    if false {
                        switch 4294967295 {
                        case -321:
                            v4 -= (v4 * v4)
                        }
                        v4 += 65535
                        for fi67 := 0; fi67 < 8; fi67++ {
                            v4 = v4
                            v68 := GenS8{F0: (!false)}
                            v69 := GenS9{F0: -910, F1: 526, F2: 1000}
                            _ = fi67
                            _ = v68
                            _ = v69
                        }
                    } else {
                        fmt.Println(v8)
                        if (float64(v1) > 0.1) {
                            switch 9223372036854775807 {
                            case 313:
                                for fi70 := 0; fi70 < 13; fi70++ {
                                    v1 += v1
                                    _ = fi70
                                }
                            case 564:
                                v4 -= v4
                            case 128:
                                v20 = v1
                            default:
                                fmt.Println(v20, v8)
                            }
                            var v71 uint64 = 0
                            var v72 int64 = 0
                            _ = v71
                            _ = v72
                        }
                    }
                    var v73 int64 = 9223372036854775807
                    _ = v73
                } else {
                    v74 := GenS1{F0: 846, F1: 713}
                    v1 += 0.0
                    _ = v74
                }
            } else {
                if (1 == 348) {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    var v75 uint64 = 2
                    _ = v75
                } else {
                    v3 = append(v3, (("foo bar" + "") + "a"))
                    var v76 int16 = -1
                    var v77 int = 2147483647
                    _ = v76
                    _ = v77
                }
                if true {
                    _ = len(v3)
                    var v78 float32 = v1
                    _ = v78
                } else {
                    v20 = (0.0 - ((v1 / v1) + 0.1))
                    switch "hello" {
                    case "foo bar":
                        var v79 int16 = 1
                        _ = v79
                    case "\n\t":
                        var v80 bool = false
                        _ = v80
                    default:
                        fmt.Println(v8)
                    }
                }
            }
            _ = v20
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    var v81 uint64 = 0
    v82 := GenS7{F0: 255}
    for fi83 := 0; fi83 < 2; fi83++ {
        switch ("foo bar" + "") {
        case "foo bar":
            var v84 string = "hello"
            _ = v84
        default:
            for fi85 := 0; fi85 < 18; fi85++ {
                var v86 int64 = -473
                v87 := GenS10{F0: uint16(v4), F1: uint64(v4), F2: ("hello" + "")}
                fmt.Println(v4, v81, fi85)
                switch "hello" {
                case "":
                    v88 := GenS9{F0: (fi83 - 4294967295), F1: 42, F2: -1}
                    _ = v88
                default:
                    v81 = (v81 ^ (v81 & v81))
                }
                var v89 uint = 304
                _ = fi85
                _ = v86
                _ = v87
                _ = v89
            }
        }
        var v90 bool = ((float64(v1) - -2.25) != 0.1)
        _ = fi83
        _ = v90
    }
    var v91 int16 = 1000
    var v92 uint16 = 1
    v4 = v4
    if true {
        for fi93 := 0; fi93 < 6; fi93++ {
            var v94 bool = (!(((v92 * v92) >= (v92 % (((v92 >> ((4294967295) & 63))) | 1))) || false))
            if v94 {
                var v95 uint = 2
                var v96 bool = (!v94)
                for fi97 := 0; fi97 < 11; fi97++ {
                    v1 += v1
                    _ = fi97
                }
                _ = v95
                _ = v96
            }
            _ = fi93
            _ = v94
        }
        switch "foo bar" {
        case "foo bar":
            switch "hello" {
            case "hello":
                fmt.Println(v8, v91, v4)
            case "":
                if (42 >= 139) {
                    var v98 bool = (0 < 119)
                    var v99 uint16 = (v92 >> ((42) & 63))
                    for fi100 := 0; fi100 < 5; fi100++ {
                        var v101 float32 = ((3.14159 / 1.5) - 0.0)
                        if ((v1 > (v1 + float32(float64(1.5)))) && true) {
                            fmt.Println(v98, v101)
                            if (v98 && ((v81 - 2) >= 1)) {
                                v91 += v91
                            } else {
                                for fi102 := 0; fi102 < 10; fi102++ {
                                    var v103 int64 = int64(v8)
                                    fmt.Println(v101, v99)
                                    var v104 uint8 = 128
                                    _ = fi102
                                    _ = v103
                                    _ = v104
                                }
                            }
                        } else {
                            var v105 bool = (2 > v4)
                            _ = v105
                        }
                        for fi106 := 0; fi106 < 10; fi106++ {
                            v92 += (v99 / (((v92 | (^(v92)))) | 1))
                            if (398 == 1) {
                                v2 = append(v2, "\n\t")
                                v4 -= (v4 | (v4 >> ((65535) & 63)))
                                switch "" {
                                case "":
                                    v99 -= (^(v92))
                                default:
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                }
                            }
                            _ = fi106
                        }
                        _ = fi100
                        _ = v101
                    }
                    _ = v98
                    _ = v99
                } else {
                    if len(v3) > 0 {
                        v3[0] = "hello"
                    }
                    v107 := GenS9{F0: 827, F1: 42, F2: 32767}
                    for fi108 := 0; fi108 < 14; fi108++ {
                        if true {
                            var v109 int32 = 1
                            var v110 uint = 255
                            var v111 bool = true
                            _ = v109
                            _ = v110
                            _ = v111
                        }
                        _ = fi108
                    }
                    _ = v107
                }
            }
        case "\n\t":
            v112 := GenS9{F0: 42, F1: 947, F2: (v91 >> ((0) & 63))}
            _ = v112
        }
        if (32768 != 4294967295) {
            var v113 int8 = 127
            v92 -= (v92 + uint16(v81))
            fmt.Println(v92)
            _ = v113
        }
    } else {
        v114 := GenS3{F0: v81}
        _ = v114
    }
    v81 += (v81 << ((18446744073709551615) & 63))
    for fi115 := 0; fi115 < 15; fi115++ {
        var v116 string = ("hello" + "")
        _ = fi115
        _ = v116
    }
    fmt.Println(v91)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("v81:", v81)
    fmt.Println("v91:", v91)
    fmt.Println("v92:", v92)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v81
    _ = v82
    _ = v91
    _ = v92
}
