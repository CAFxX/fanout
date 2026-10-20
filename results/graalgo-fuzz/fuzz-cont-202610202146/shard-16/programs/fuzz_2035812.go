package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16 }
type GenS2 struct { F0 bool; F1 uint8; F2 uint64 }
type GenS3 struct { F0 string; F1 int16 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 string; F1 int8 }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 int32; F1 uint16 }
type GenS10 struct { F0 bool; F1 bool }
type GenS11 struct { F0 int64; F1 uint8 }
func genFn0(p0 uint64, p1 int8, p2 bool) uint64 {
    p1 -= p1
    switch 4294967295 {
    case 142:
        var v5 int16 = 42
        _ = v5
    case 32768:
        fmt.Println(p0, p0, p0)
    }
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = p2
    return (^(p0))
}

func main() {
    v1 := make(map[string]int)
    v2 := []int{42, -9223372036854775808}
    v3 := make(map[uint32]int)
    var v4 uint8 = 1
    v2 = append(v2, -2147483648)
    v4 -= 2
    var v6 int8 = -1
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v7 int8 = 100
    var v8 bool = true
    var v9 float32 = 0.1
    if v8 {
        switch ((("\n\t" + ("foo bar" + "")) + "hello") + "\n\t") {
        case "a":
            v10 := GenS1{F0: (309 >= 255), F1: 32767}
            _ = v10
        }
        fmt.Println(v4)
    } else {
        v7 = (-(v6))
        for fi11 := 0; fi11 < 12; fi11++ {
            var v12 int = 1
            for fi13 := 0; fi13 < 10; fi13++ {
                if ("" > ("" + (("a" + ("a" + ("" + (("\n\t" + "a") + "foo bar")))) + (("" + "") + "a")))) {
                    switch fi11 {
                    case -2147483648:
                        if v8 {
                            if ((-(v12)) > -127) {
                                v14 := GenS1{F0: false, F1: 255}
                                _ = v14
                            }
                        }
                    }
                    if (-9223372036854775808 == 9223372036854775807) {
                        v15 := GenS1{F0: (v8 || (!(v4 != (^(v4))))), F1: 1}
                        var v16 string = "a"
                        v17 := GenS1{F0: false, F1: 65535}
                        _ = v15
                        _ = v16
                        _ = v17
                    } else {
                        _ = len(v1)
                        v18 := GenS2{F0: (!false), F1: 42, F2: 985}
                        var v19 int32 = 1
                        _ = v18
                        _ = v19
                    }
                    v20 := GenS2{F0: v8, F1: (v4 << ((255) & 63)), F2: 4294967295}
                    if (!false) {
                        v21 := GenS3{F0: "\n\t", F1: -32768}
                        if v8 {
                            v22 := GenS3{F0: ("" + "\n\t"), F1: 330}
                            if true {
                                v23 := GenS1{F0: v8, F1: 42}
                                v24 := v1["\n\t"]
                                switch "" {
                                case "a":
                                    v8 = ("" > ("a" + ((("" + "foo bar") + "") + "")))
                                case "":
                                    v25 := GenS4{F0: 1000}
                                    _ = v25
                                }
                                _ = v23
                                _ = v24
                            } else {
                                if false {
                                    v7 = (v6 * (v7 & 100))
                                }
                                fmt.Println(v8, fi13, v9)
                                var v26 int16 = 1000
                                _ = v26
                            }
                            fmt.Println(v6, v9, fi11)
                            var v27 uint16 = 2
                            _ = v22
                            _ = v27
                        } else {
                            v28 := GenS1{F0: ((!v8) || false), F1: 255}
                            _ = v28
                        }
                        var v29 uint16 = 758
                        _ = v21
                        _ = v29
                    } else {
                        v30 := GenS1{F0: ((0 >= 32767) && (1 != 748)), F1: 1}
                        switch (fi11 << ((42) & 63)) {
                        case -2147483648:
                            if (1.5 < -2.25) {
                                var v31 uint8 = (v4 | (v4 ^ 124))
                                var v32 int8 = v7
                                _ = v31
                                _ = v32
                            } else {
                                if false {
                                    if ((v9 - float32(1e10)) == (-2.25 - float32((((-2.25 + 0.0) / 1.5) - ((0.1 * (-2.25 * float64(v9))) * 3.14159))))) {
                                        var v33 uint32 = 2
                                        v33 -= (v33 + v33)
                                        if true {
                                            v34 := GenS1{F0: v8, F1: 935}
                                            _ = v34
                                        }
                                        _ = v33
                                    }
                                } else {
                                    switch 2147483647 {
                                    case 107:
                                        v1[("hello" + "hello")] = 32768
                                    default:
                                        var v35 uint32 = 1
                                        _ = v35
                                    }
                                    if ((v6 | (v7 & (v7 - (v7 / ((1) | 1))))) == 1) {
                                        fmt.Println(v12, v8)
                                        switch v12 {
                                        case 0:
                                            if len(v2) > 0 {
                                                v2[0] = (fi13 - fi11)
                                            }
                                        case 9223372036854775807:
                                            fmt.Println(v12, v7, v7)
                                        }
                                        v36 := GenS5{F0: 32767}
                                        fmt.Println(v9)
                                        v37 := GenS1{F0: true, F1: 65535}
                                        _ = v36
                                        _ = v37
                                    } else {
                                        v38 := GenS1{F0: (128 < v4), F1: 255}
                                        _ = v38
                                    }
                                    v39 := GenS6{F0: ("" + ("\n\t" + ("foo bar" + (("a" + ("foo bar" + "")) + ("foo bar" + "\n\t"))))), F1: v6}
                                    v40 := GenS2{F0: (1e10 != v9), F1: (v4 % ((1) | 1)), F2: 18446744073709551615}
                                    v9 = v9
                                    _ = v39
                                    _ = v40
                                }
                                var v41 uint = 777
                                _ = v41
                            }
                        case -947:
                            var v42 int64 = 0
                            v2 = append(v2, fi11)
                            v43 := GenS2{F0: (v4 < 125), F1: v4, F2: 1}
                            _ = v42
                            _ = v43
                        case 9223372036854775807:
                            v2 = append(v2, (v12 / ((-391) | 1)))
                        default:
                            fmt.Println(v4)
                        }
                        if (!(!v8)) {
                            var v44 int16 = 32767
                            _ = v44
                        }
                        v6 = (v7 ^ 0)
                        _ = v30
                    }
                    switch ("\n\t" + "a") {
                    case "hello":
                        v4 -= v4
                    case "\n\t":
                        var v45 bool = v8
                        _ = v45
                    }
                    v46 := GenS7{F0: -9223372036854775808}
                    _ = v20
                    _ = v46
                } else {
                    var v47 uint = 4294967295
                    if (!true) {
                        var v48 float64 = 1.5
                        _ = v48
                    } else {
                        var v49 string = "hello"
                        if (!v8) {
                            v50 := GenS2{F0: v8, F1: v4, F2: 42}
                            if false {
                                var v51 float32 = v9
                                v9 += 0.0
                                _ = v51
                            }
                            switch v49 {
                            case "hello":
                                var v52 int64 = 9223372036854775807
                                _ = v52
                            default:
                                switch ("\n\t" + "hello") {
                                case "a":
                                    v53 := GenS3{F0: (v49 + v49), F1: int16(v4)}
                                    _ = v53
                                }
                            }
                            _ = v50
                        }
                        _ = v49
                    }
                    _ = v47
                }
                var v54 int32 = 815
                v55 := GenS1{F0: (256 > 32767), F1: 232}
                if v8 {
                    var v56 int = (-(fi13))
                    delete(v3, 1)
                    v54 = (v54 + 695)
                    _ = v56
                }
                switch 0 {
                case 2147483647:
                    if v8 {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        fmt.Println(v12, v12)
                        switch "hello" {
                        case "":
                            if (v6 >= -128) {
                                var v57 bool = (773 > 1)
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                _ = v57
                            } else {
                                delete(v3, 2147483648)
                            }
                        case "foo bar":
                            v4 = (v4 << ((2) & 63))
                        default:
                            if v8 {
                                var v58 uint8 = v4
                                var v59 int64 = 244
                                switch (("a" + ("a" + ("\n\t" + ("" + ("foo bar" + ("a" + "")))))) + "\n\t") {
                                case "\n\t":
                                    if (v8 || v8) {
                                        switch v12 {
                                        case 0:
                                            v4 = (v58 & (v4 + v4))
                                        case 1:
                                            fmt.Println(v4, fi13)
                                        case 859:
                                            if true {
                                                v54 += 0
                                                fmt.Println(v59)
                                                fmt.Println(v59)
                                            }
                                        default:
                                            v54 += (v54 << ((18446744073709551615) & 63))
                                            v59 = v59
                                        }
                                        var v60 uint32 = 0
                                        _ = v60
                                    }
                                case "hello":
                                    switch ("\n\t" + "\n\t") {
                                    case "":
                                        v9 -= v9
                                    case "\n\t":
                                        v59 = 859
                                    case "hello":
                                        v59 += (v59 + -9223372036854775808)
                                    default:
                                        var v61 float64 = -2.25
                                        _ = v61
                                    }
                                case "a":
                                    v62 := GenS7{F0: -9223372036854775808}
                                    _ = v62
                                default:
                                    v6 -= 100
                                }
                                _ = v58
                                _ = v59
                            } else {
                                v4 -= v4
                                delete(v1, "")
                                v4 += 129
                            }
                        }
                        if v8 {
                            v4 -= (v4 & v4)
                            delete(v3, 1)
                            if (v7 < v7) {
                                v8 = v8
                                if v8 {
                                    var v63 float64 = 1e10
                                    var v64 float64 = v63
                                    _ = v63
                                    _ = v64
                                } else {
                                    var v65 float64 = 1.5
                                    _ = v65
                                }
                                fmt.Println(v4)
                                if (0 > 0) {
                                    delete(v3, 42)
                                }
                                v9 += (v9 * (v9 * v9))
                                v7 = (v7 >> ((42) & 63))
                                v66 := GenS3{F0: ("hello" + ""), F1: 42}
                                delete(v1, "\n\t")
                                _ = v66
                            }
                            if (1000 < 42) {
                                v6 += v7
                                var v67 uint16 = 2
                                _ = v67
                            }
                            v68 := GenS7{F0: 9223372036854775807}
                            _ = v68
                        } else {
                            var v69 string = "a"
                            if v8 {
                                var v70 uint32 = 0
                                switch "a" {
                                case "":
                                    if (985 > 42) {
                                        var v71 int64 = 42
                                        _ = v71
                                    } else {
                                        v72 := GenS5{F0: 1}
                                        if (0 == 1) {
                                            var v73 float32 = float32(1.5)
                                            var v74 int = fi11
                                            v8 = (100 <= v6)
                                            _ = v73
                                            _ = v74
                                        }
                                        v54 += (v54 | 432)
                                        var v75 uint = 710
                                        _ = v72
                                        _ = v75
                                    }
                                default:
                                    fmt.Println(fi11, fi13)
                                }
                                var v76 uint8 = (v4 - 2)
                                _ = v70
                                _ = v76
                            }
                            v12 = 9223372036854775807
                            _ = v69
                        }
                        fmt.Println(v12, v6)
                    } else {
                        if true {
                            if v8 {
                                switch 909 {
                                case 32767:
                                    v4 = (^(v4))
                                case 9223372036854775807:
                                    _ = len(v2)
                                }
                            } else {
                                fmt.Println(v8)
                                var v77 bool = v8
                                fmt.Println(v8, v6)
                                _ = v77
                            }
                        }
                        v12 = v12
                        if (false && v8) {
                            v8 = true
                            fmt.Println(v6, fi13)
                        } else {
                            v9 += 0.0
                        }
                    }
                }
                v6 = (v7 * v7)
                var v78 uint8 = 2
                _ = fi13
                _ = v54
                _ = v55
                _ = v78
            }
            _ = fi11
            _ = v12
        }
    }
    for fi79 := 0; fi79 < 7; fi79++ {
        v6 -= v6
        if ((!v8) || v8) {
            v8 = true
            if (65535 != 908) {
                v4 = (v4 % ((v4) | 1))
            }
        } else {
            for fi80 := 0; fi80 < 20; fi80++ {
                var v81 int16 = 0
                if false {
                    switch 159 {
                    case 39:
                        if ((-2147483648 < 2147483647) && (-1 < (fi80 + (fi80 / ((-54) | 1))))) {
                            delete(v1, ("a" + ("a" + "")))
                        } else {
                            v4 = 255
                            fmt.Println(v81, v8, v7)
                            v82 := GenS2{F0: (v81 <= (^(v81))), F1: 80, F2: 2}
                            v83 := GenS8{F0: 2}
                            switch -372 {
                            case 9223372036854775807:
                                if (false && v8) {
                                    switch (fi80 + (fi80 & 2147483647)) {
                                    case 4294967295:
                                        var v84 float64 = 1e10
                                        _ = v84
                                    case -2147483648:
                                        var v85 uint8 = v4
                                        _ = v85
                                    case -447:
                                        v86 := GenS9{F0: 993, F1: 32767}
                                        v87 := GenS7{F0: 0}
                                        _ = v86
                                        _ = v87
                                    }
                                    var v88 string = "a"
                                    var v89 float32 = v9
                                    _ = v88
                                    _ = v89
                                } else {
                                    v90 := GenS1{F0: false, F1: 42}
                                    fmt.Println(v4)
                                    _ = v90
                                }
                            default:
                                if true {
                                    v6 = v7
                                    if true {
                                        if false {
                                            var v91 float64 = 0.1
                                            var v92 int = fi80
                                            _ = v91
                                            _ = v92
                                        }
                                        if (602 == 9223372036854775807) {
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                            if len(v2) > 0 {
                                                v2[0] = fi80
                                            }
                                            var v93 float64 = (((3.14159 - (3.14159 * (float64(3.14159) * ((0.0 + 1.5) - 1.5)))) + float64((((-2.25 + v9) - -2.25) * v9))) / 3.14159)
                                            _ = v93
                                        }
                                        if v8 {
                                            switch "a" {
                                            case "\n\t":
                                                var v94 uint32 = 2
                                                _ = v94
                                            case "":
                                                v7 += (v7 | v7)
                                            case "foo bar":
                                                fmt.Println(v81, v4, v81)
                                            default:
                                                v7 -= 0
                                            }
                                            switch "" {
                                            case "":
                                                v9 += v9
                                                v7 += 127
                                                v4 = (v4 / ((v4) | 1))
                                            case "foo bar":
                                                var v95 int16 = 0
                                                _ = v95
                                            }
                                        } else {
                                            v6 -= v6
                                            var v96 int16 = v81
                                            _ = v96
                                        }
                                        v97 := GenS10{F0: true, F1: v8}
                                        _ = v97
                                    } else {
                                        v98 := GenS5{F0: 1}
                                        var v99 uint64 = 18446744073709551615
                                        switch fi80 {
                                        case 42:
                                            if v8 {
                                                fmt.Println(fi79, v7)
                                                switch 545 {
                                                case -202:
                                                    v1[("a" + ("foo bar" + "\n\t"))] = 2147483647
                                                }
                                            }
                                        case 0:
                                            if v8 {
                                                var v100 bool = v8
                                                var v101 uint8 = (^(v4))
                                                var v102 int64 = 528
                                                _ = v100
                                                _ = v101
                                                _ = v102
                                            } else {
                                                fmt.Println(v8)
                                            }
                                        case 32768:
                                            switch 2147483647 {
                                            case 120:
                                                v4 = 1
                                                fmt.Println(v4)
                                            case 127:
                                                var v103 uint8 = (v4 - 42)
                                                v104 := GenS7{F0: -918}
                                                switch fi79 {
                                                case -9223372036854775808:
                                                    v99 -= (v99 / ((v99) | 1))
                                                }
                                                _ = v103
                                                _ = v104
                                            case -1:
                                                if (!((v7 - (v7 >> ((42) & 63))) <= v6)) {
                                                    v6 = v7
                                                } else {
                                                    _ = len(v2)
                                                }
                                            default:
                                                fmt.Println(v99, v81)
                                            }
                                        }
                                        _ = v98
                                        _ = v99
                                    }
                                }
                                if v8 {
                                    v4 += 247
                                }
                            }
                            v9 = 1.5
                            v4 = (^(v4))
                            var v105 int16 = int16(v4)
                            _ = v105
                            _ = v82
                            _ = v83
                        }
                    }
                    fmt.Println(v81)
                    switch (fi79 / (((fi80 - (fi79 | -506))) | 1)) {
                    case 833:
                        switch ("a" + (("\n\t" + "a") + (("a" + "foo bar") + "foo bar"))) {
                        case "":
                            var v106 uint64 = genFn0(825, -23, false)
                            _ = v106
                        case "a":
                            v81 -= v81
                        case "foo bar":
                            var v107 uint64 = 18446744073709551615
                            _ = v107
                        }
                    case 42:
                        if v8 {
                            delete(v1, "a")
                        } else {
                            var v108 bool = v8
                            _ = v108
                        }
                    }
                }
                switch "\n\t" {
                case "foo bar":
                    if len(v2) > 0 {
                        v2[0] = 127
                    }
                case "hello":
                    v8 = (v8 || v8)
                }
                v8 = ((!(42 > 42)) || ((2 > 793) && v8))
                _ = fi80
                _ = v81
            }
        }
        var v109 int = fi79
        _ = fi79
        _ = v109
    }
    switch -864 {
    case -21:
        v110 := GenS2{F0: v8, F1: v4, F2: 0}
        _ = v110
    case 128:
        v6 -= v6
    default:
        switch 166 {
        case 32767:
            if len(v2) > 0 {
                v2[0] = -417
            }
        case -795:
            var v111 string = ""
            _ = v111
        case 9223372036854775807:
            v8 = (530 != 0)
        }
    }
    v8 = v8
    v8 = v8
    v8 = (32767 == 65535)
    if (!false) {
        v9 = 0.0
        v6 = v7
        for fi112 := 0; fi112 < 18; fi112++ {
            var v113 uint32 = 0
            _ = fi112
            _ = v113
        }
    } else {
        v114 := GenS11{F0: 42, F1: v4}
        var v115 float32 = 0.0
        _ = len(v2)
        _ = v114
        _ = v115
    }
    fmt.Println(v4, v8)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
