package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint32; F1 int16 }
type GenS4 struct { F0 uint64; F1 uint }
type GenS5 struct { F0 int; F1 uint8; F2 uint64 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 int; F1 bool }
type GenS10 struct { F0 int64 }
type GenS11 struct { F0 int64 }
type GenS12 struct { F0 string; F1 int8; F2 uint32 }
type GenS13 struct { F0 int32; F1 int32; F2 string }
func genFn0(p0 uint8, p1 uint, p2 int16) bool {
    var v7 int64 = 42
    var v8 string = "hello"
    switch -801 {
    case -9223372036854775808:
        if true {
            v9 := []int{}
            var v10 uint16 = 32767
            if (!(p1 >= 2)) {
                fmt.Println(v8, v10)
                fmt.Println(p1)
                var v11 int16 = (p2 | (p2 >> ((18446744073709551615) & 63)))
                _ = v11
            }
            _ = v10
            _ = v9
        }
    case 9223372036854775807:
        switch "foo bar" {
        case "a":
            var v12 int = 786
            _ = v12
        case "foo bar":
            for fi13 := 0; fi13 < 14; fi13++ {
                var v14 uint8 = (p0 | p0)
                _ = fi13
                _ = v14
            }
        case "\n\t":
            p2 += (p2 % ((p2) | 1))
        default:
            for fi15 := 0; fi15 < 10; fi15++ {
                switch v8 {
                case "foo bar":
                    var v16 int = (fi15 % ((127) | 1))
                    _ = v16
                }
                _ = fi15
            }
        }
    case 32768:
        var v17 int32 = 42
        _ = v17
    }
    var v18 int64 = (v7 ^ v7)
    _ = p0
    _ = p1
    _ = p2
    _ = v18
    _ = v7
    _ = v8
    return (int(v18) > -797)
}
func genFn1(p0 uint64) uint {
    var v19 uint8 = 255
    v20 := GenS1{F0: 65535, F1: 100}
    if (!false) {
        switch "a" {
        case "\n\t":
            p0 += (^(p0))
        case "foo bar":
            v19 = (v19 << ((42) & 63))
        default:
            if true {
                for fi21 := 0; fi21 < 19; fi21++ {
                    fmt.Println(p0, fi21)
                    _ = fi21
                }
                fmt.Println(p0, p0)
                var v22 int16 = -582
                _ = v22
            }
        }
        for fi23 := 0; fi23 < 7; fi23++ {
            for fi24 := 0; fi24 < 1; fi24++ {
                switch (fi24 >> ((482) & 63)) {
                case 128:
                    v19 += (^(v19))
                }
                _ = fi24
            }
            v19 += 42
            v19 += (^(v19))
            _ = fi23
        }
        v25 := []int{4294967295, 9223372036854775807}
        _ = v25
    } else {
        var v26 uint8 = (v19 + 1)
        _ = v26
    }
    p0 -= p0
    _ = p0
    _ = v19
    _ = v20
    return 42
}

func main() {
    var v1 float32 = 1e10
    var v2 uint8 = 59
    var v3 int8 = 100
    var v4 float64 = 1e10
    var v5 bool = (true || (true || ((("hello" + "hello") + ((("hello" + "foo bar") + "\n\t") + ((("a" + "hello") + "foo bar") + "foo bar"))) == (("a" + "") + ""))))
    v6 := make(map[string]bool)
    for fi27 := 0; fi27 < 9; fi27++ {
        v4 = (((1e10 * v4) - 0.0) * 3.14159)
        v28 := GenS2{F0: 489}
        fmt.Println(v2, v3)
        _ = fi27
        _ = v28
    }
    switch "" {
    case "hello":
        v29 := GenS1{F0: 687, F1: v3}
        _ = v29
    case "\n\t":
        for fi30 := 0; fi30 < 20; fi30++ {
            switch 128 {
            case 127:
                v31 := []int{2147483647, 133, (fi30 << ((0) & 63))}
                _ = v31
            case 2147483647:
                v32 := []int{2147483647, fi30, fi30}
                _ = v32
            case -698:
                v3 = (-(v3))
            default:
                fmt.Println(v5)
            }
            switch (fi30 & fi30) {
            case -98:
                if false {
                    switch fi30 {
                    case 1:
                        for fi33 := 0; fi33 < 1; fi33++ {
                            var v34 float64 = v4
                            v6["a"] = v5
                            v35 := GenS2{F0: 156}
                            _ = fi33
                            _ = v34
                            _ = v35
                        }
                    case 127:
                        if ("hello" <= ("hello" + "\n\t")) {
                            var v36 float64 = v4
                            var v37 uint32 = 2147483648
                            _ = v36
                            _ = v37
                        } else {
                            if (((v2 + v2) < 206) && false) {
                                for fi38 := 0; fi38 < 20; fi38++ {
                                    v39 := v6["\n\t"]
                                    if true {
                                        var v40 int16 = 42
                                        if (v40 < v40) {
                                            if true {
                                                var v41 int32 = 2147483647
                                                var v42 float32 = (1e10 + v1)
                                                v43 := v6["\n\t"]
                                                v44 := GenS3{F0: 0, F1: (v40 << ((0) & 63))}
                                                _ = v41
                                                _ = v42
                                                _ = v43
                                                _ = v44
                                            } else {
                                                if v39 {
                                                    if v5 {
                                                        v4 -= v4
                                                        var v45 uint64 = 42
                                                        _ = v45
                                                    } else {
                                                        v46 := []int{}
                                                        v3 += v3
                                                        v47 := v6["a"]
                                                        fmt.Println(v39)
                                                        var v48 bool = (!(v5 && true))
                                                        _ = v46
                                                        _ = v47
                                                        _ = v48
                                                    }
                                                    var v49 float64 = v4
                                                    _ = v49
                                                }
                                                v40 += (v40 & (v40 & v40))
                                                var v50 int8 = (v3 ^ v3)
                                                _ = v50
                                            }
                                            switch 4294967295 {
                                            case 1:
                                                if (true || (0 != uint64(v2))) {
                                                    fmt.Println(fi30)
                                                    v2 -= 46
                                                    fmt.Println(fi38, fi38)
                                                } else {
                                                    switch "" {
                                                    case "a":
                                                        delete(v6, "hello")
                                                    case "foo bar":
                                                        if (!true) {
                                                            v51 := GenS4{F0: 1, F1: 0}
                                                            if (!(!true)) {
                                                                switch (fi38 >> ((1) & 63)) {
                                                                case -206:
                                                                    var v52 float32 = v1
                                                                    _ = v52
                                                                }
                                                                if true {
                                                                    switch "" {
                                                                    case "":
                                                                        v5 = true
                                                                    case "hello":
                                                                        v5 = ("hello" > "")
                                                                    default:
                                                                        var v53 uint = genFn1(532)
                                                                        _ = v53
                                                                    }
                                                                }
                                                                _ = len(v6)
                                                            } else {
                                                                v54 := GenS5{F0: (fi30 - (fi30 - (fi38 - 1))), F1: (^(v2)), F2: 146}
                                                                if (-386 != -1) {
                                                                    if v5 {
                                                                        if (!((genFn1(4294967295) <= genFn1(42)) && (42 == -66))) {
                                                                            if v5 {
                                                                                fmt.Println(v3)
                                                                                fmt.Println(v3, fi30, fi30)
                                                                            }
                                                                            v55 := v6["\n\t"]
                                                                            _ = v55
                                                                        }
                                                                        v40 = v40
                                                                    } else {
                                                                        var v56 uint32 = 232
                                                                        switch ("a" + "\n\t") {
                                                                        case "\n\t":
                                                                            if v5 {
                                                                                v6[(((("a" + ("a" + "hello")) + ("foo bar" + ("" + "\n\t"))) + "\n\t") + ("a" + ("\n\t" + "foo bar")))] = (1 < 0)
                                                                                v4 = 0.0
                                                                            } else {
                                                                                switch fi38 {
                                                                                case 9223372036854775807:
                                                                                    var v57 int32 = -2147483648
                                                                                    _ = v57
                                                                                case 1:
                                                                                    v40 = (v40 << ((255) & 63))
                                                                                }
                                                                            }
                                                                        case "":
                                                                            v39 = false
                                                                        default:
                                                                            var v58 uint64 = uint64(v56)
                                                                            _ = v58
                                                                        }
                                                                        _ = v56
                                                                    }
                                                                }
                                                                var v59 uint = 18446744073709551615
                                                                _ = v54
                                                                _ = v59
                                                            }
                                                            _ = v51
                                                        }
                                                    case "hello":
                                                        v60 := []int{fi38, (fi38 + (fi30 + (fi38 ^ 32767))), fi30, fi30}
                                                        _ = v60
                                                    default:
                                                        v61 := GenS6{F0: 1}
                                                        _ = v61
                                                    }
                                                }
                                                if v39 {
                                                    if v39 {
                                                        fmt.Println(v40)
                                                        v1 = (v1 - v1)
                                                    } else {
                                                        v3 += v3
                                                    }
                                                }
                                                fmt.Println(v1, v39)
                                            case 32767:
                                                v62 := GenS7{F0: 127}
                                                _ = v62
                                            default:
                                                if (-1 >= 354) {
                                                    v6["foo bar"] = (!(0 >= 1))
                                                    if v39 {
                                                        fmt.Println(fi38, v4)
                                                        delete(v6, ("a" + ""))
                                                        if v5 {
                                                            fmt.Println(v1, fi38)
                                                            var v63 uint16 = 32767
                                                            v64 := GenS8{F0: 128}
                                                            _ = v63
                                                            _ = v64
                                                        }
                                                        if false {
                                                            v5 = (true || (!v39))
                                                            var v65 string = "foo bar"
                                                            switch (v65 + "foo bar") {
                                                            case "foo bar":
                                                                v66 := GenS9{F0: fi30, F1: (65535 <= 42)}
                                                                _ = v66
                                                            case "":
                                                                var v67 uint16 = 0
                                                                _ = v67
                                                            case "hello":
                                                                delete(v6, (v65 + "foo bar"))
                                                            }
                                                            _ = v65
                                                        }
                                                    }
                                                } else {
                                                    var v68 int16 = v40
                                                    if false {
                                                        v69 := GenS7{F0: (v2 << ((126) & 63))}
                                                        _ = v69
                                                    } else {
                                                        v70 := []int{1}
                                                        _ = v70
                                                    }
                                                    v71 := v6["\n\t"]
                                                    fmt.Println(v40, v71, v39)
                                                    v68 = v68
                                                    v71 = ((fi30 & fi30) >= (-(fi30)))
                                                    switch ("\n\t" + "a") {
                                                    case "\n\t":
                                                        switch "foo bar" {
                                                        case "":
                                                            v72 := []int{int(v2)}
                                                            _ = v72
                                                        case "hello":
                                                            v1 = ((v1 * float32(v4)) * v1)
                                                        }
                                                    }
                                                    _ = v68
                                                    _ = v71
                                                }
                                            }
                                        } else {
                                            v73 := GenS7{F0: (v2 / (((^(v2))) | 1))}
                                            _ = v73
                                        }
                                        v74 := GenS10{F0: -9223372036854775808}
                                        _ = v40
                                        _ = v74
                                    }
                                    v2 = v2
                                    _ = fi38
                                    _ = v39
                                }
                            }
                            for fi75 := 0; fi75 < 11; fi75++ {
                                v4 -= 1.5
                                if false {
                                    if (32767 < 0) {
                                        v76 := GenS6{F0: 42}
                                        v77 := GenS7{F0: 245}
                                        _ = v76
                                        _ = v77
                                    }
                                    v5 = false
                                    var v78 uint = 2
                                    var v79 uint32 = 441
                                    var v80 float64 = float64(1.5)
                                    var v81 bool = (0 > 374)
                                    var v82 uint8 = v2
                                    if v5 {
                                        v80 = (v4 * (1.5 / v4))
                                    } else {
                                        v4 = v80
                                        switch -1 {
                                        case 128:
                                            switch "\n\t" {
                                            case "foo bar":
                                                var v83 float64 = (v80 + -2.25)
                                                _ = v83
                                            case "hello":
                                                if true {
                                                    switch "\n\t" {
                                                    case "a":
                                                        switch "foo bar" {
                                                        case "hello":
                                                            var v84 uint = v78
                                                            _ = v84
                                                        case "foo bar":
                                                            var v85 int = -148
                                                            _ = v85
                                                        case "a":
                                                            v86 := []int{(fi75 % (((^(fi75))) | 1)), (fi75 % (((-(fi75))) | 1))}
                                                            _ = v86
                                                        default:
                                                            v87 := GenS11{F0: 9223372036854775807}
                                                            _ = v87
                                                        }
                                                        if ((-(fi30)) == 128) {
                                                            v3 = 0
                                                        } else {
                                                            v80 -= -2.25
                                                            var v88 int8 = 106
                                                            _ = v88
                                                        }
                                                        var v89 uint32 = 65535
                                                        v90 := []int{(^(fi30)), fi75, (fi75 % (((^(fi30))) | 1))}
                                                        _ = v89
                                                        _ = v90
                                                    case "\n\t":
                                                        var v91 float64 = float64(1e10)
                                                        _ = v91
                                                    case "hello":
                                                        switch (fi30 << ((v78) & 63)) {
                                                        case -171:
                                                            v92 := []int{}
                                                            _ = v92
                                                        default:
                                                            v81 = false
                                                        }
                                                    }
                                                    var v93 uint16 = 65535
                                                    v6[("a" + "a")] = true
                                                    _ = v93
                                                }
                                            case "":
                                                v79 += v79
                                            default:
                                                fmt.Println(v80, v5)
                                            }
                                        case 222:
                                            var v94 string = "hello"
                                            if v5 {
                                                v5 = ((!v81) && v81)
                                                fmt.Println(v79)
                                                v78 = (v78 % ((v78) | 1))
                                            }
                                            if ((v94 == "") && false) {
                                                v2 = (v82 + 42)
                                                v5 = (32767 != 376)
                                            }
                                            _ = v94
                                        case 32768:
                                            var v95 int64 = 0
                                            _ = v95
                                        }
                                    }
                                    _ = v78
                                    _ = v79
                                    _ = v80
                                    _ = v81
                                    _ = v82
                                }
                                v96 := []int{}
                                _ = fi75
                                _ = v96
                            }
                            fmt.Println(fi30, v2, v5)
                        }
                    }
                    var v97 bool = (0 == 908)
                    v98 := []int{(fi30 ^ (-(fi30))), (-(fi30)), fi30}
                    _ = v97
                    _ = v98
                }
            case -9223372036854775808:
                var v99 int = fi30
                _ = v99
            }
            _ = fi30
        }
    case "foo bar":
        var v100 int32 = 761
        _ = v100
    default:
        if v5 {
            var v101 int16 = 32767
            _ = v101
        } else {
            if true {
                delete(v6, "hello")
            } else {
                fmt.Println(v3, v4)
                for fi102 := 0; fi102 < 2; fi102++ {
                    if false {
                        v103 := v6["foo bar"]
                        v104 := GenS2{F0: 151}
                        _ = v103
                        _ = v104
                    } else {
                        v3 -= -54
                        v105 := GenS12{F0: (((("\n\t" + "") + "") + (((("\n\t" + "hello") + "hello") + "\n\t") + "a")) + "\n\t"), F1: 1, F2: 1}
                        _ = v105
                    }
                    if ((!false) && (2147483647 > (fi102 % ((int(v2)) | 1)))) {
                        if false {
                            v5 = true
                            v106 := []int{(fi102 >> ((2) & 63)), (-(fi102))}
                            var v107 int16 = 660
                            _ = v106
                            _ = v107
                        }
                        v2 += 128
                        switch (^(fi102)) {
                        case -9223372036854775808:
                            v1 += ((v1 - 0.0) - (v1 / float32((v4 - float64(float32(float64(v1)))))))
                        case 0:
                            v2 -= v2
                        case 128:
                            for fi108 := 0; fi108 < 5; fi108++ {
                                if (!v5) {
                                    v109 := GenS7{F0: 1}
                                    _ = len(v6)
                                    var v110 int64 = -1
                                    _ = v109
                                    _ = v110
                                }
                                switch fi102 {
                                case 9223372036854775807:
                                    var v111 float64 = (float64(3.14159) * 3.14159)
                                    _ = v111
                                case -460:
                                    v112 := []int{2147483647, fi108}
                                    _ = v112
                                default:
                                    var v113 int16 = 42
                                    _ = v113
                                }
                                _ = fi108
                            }
                        default:
                            v114 := []int{fi102, -9223372036854775808}
                            _ = v114
                        }
                    }
                    v2 += v2
                    _ = fi102
                }
                v115 := GenS1{F0: 124, F1: v3}
                _ = v115
            }
            if true {
                switch 32767 {
                case 128:
                    fmt.Println(v4)
                }
            } else {
                var v116 string = ("a" + "\n\t")
                _ = v116
            }
            delete(v6, "hello")
        }
    }
    v4 = 1.5
    delete(v6, "a")
    v117 := []int{988, -2147483648}
    v1 = (0.1 * 3.14159)
    if len(v117) >= 2 {
        v117 = v117[1:len(v117)]
    }
    v4 -= 0.0
    var v118 float32 = v1
    var v119 int = 42
    for fi120 := 0; fi120 < 13; fi120++ {
        var v121 int32 = 42
        _ = fi120
        _ = v121
    }
    for fi122 := 0; fi122 < 4; fi122++ {
        v123 := GenS2{F0: 42}
        v124 := GenS13{F0: 2147483647, F1: 42, F2: "a"}
        _ = fi122
        _ = v123
        _ = v124
    }
    var v125 uint = genFn1(864)
    v6[("" + ("foo bar" + "\n\t"))] = (!true)
    var v126 uint32 = 65535
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v117:", len(v117), v117)
    fmt.Println("v118:", v118)
    fmt.Println("v119:", v119)
    fmt.Println("v125:", v125)
    fmt.Println("v126:", v126)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v117
    _ = v118
    _ = v119
    _ = v125
    _ = v126
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
