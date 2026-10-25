package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string }
type GenS2 struct { F0 int64; F1 int }
type GenS3 struct { F0 string; F1 int8; F2 int16 }
type GenS4 struct { F0 int16; F1 int16 }
type GenS5 struct { F0 bool; F1 int8 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 bool; F1 uint16 }
type GenS10 struct { F0 int64 }
type GenS11 struct { F0 uint64; F1 uint16 }
func genFn0(p0 bool, p1 int32, p2 int16) string {
    for fi6 := 0; fi6 < 18; fi6++ {
        for fi7 := 0; fi7 < 7; fi7++ {
            v8 := []int{(fi6 >> ((655) & 63)), 0}
            p0 = p0
            var v9 int16 = (-(p2))
            v10 := GenS1{F0: 42, F1: (("foo bar" + "") + "\n\t")}
            var v11 uint8 = 55
            v12 := GenS1{F0: 470, F1: ("" + "foo bar")}
            _ = fi7
            _ = v10
            _ = v11
            _ = v12
            _ = v8
            _ = v9
        }
        if true {
            if (!(!p0)) {
                switch (fi6 % ((fi6) | 1)) {
                case -222:
                    for fi13 := 0; fi13 < 18; fi13++ {
                        v14 := []int{fi6, fi13, int(p2), fi6}
                        _ = len(v14)
                        p2 += 32767
                        _ = fi13
                        _ = v14
                    }
                case -897:
                    if (!(42 > 42)) {
                        for fi15 := 0; fi15 < 9; fi15++ {
                            if ((!p0) && p0) {
                                fmt.Println(fi15, p2, fi6)
                                var v16 uint32 = 15
                                v17 := GenS2{F0: 504, F1: int(v16)}
                                _ = v16
                                _ = v17
                            } else {
                                switch (fi6 - (-(fi15))) {
                                case 1:
                                    if ("foo bar" > "\n\t") {
                                        p2 = 1
                                        p1 -= p1
                                    } else {
                                        var v18 int16 = (p2 / ((-716) | 1))
                                        _ = v18
                                    }
                                default:
                                    v19 := GenS3{F0: "a", F1: 100, F2: (p2 - p2)}
                                    _ = v19
                                }
                                v20 := []int{}
                                var v21 float32 = 1e10
                                var v22 int8 = -1
                                fmt.Println(p1, v22)
                                _ = v20
                                _ = v21
                                _ = v22
                            }
                            _ = fi15
                        }
                    } else {
                        var v23 int = (fi6 & fi6)
                        _ = v23
                    }
                }
                for fi24 := 0; fi24 < 8; fi24++ {
                    fmt.Println(fi6, p2)
                    _ = fi24
                }
                v25 := []int{}
                _ = v25
            }
        } else {
            for fi26 := 0; fi26 < 1; fi26++ {
                var v27 uint16 = 1
                _ = fi26
                _ = v27
            }
        }
        if (891 > 255) {
            for fi28 := 0; fi28 < 5; fi28++ {
                if p0 {
                    p2 = (p2 / ((-32768) | 1))
                    var v29 int16 = p2
                    var v30 uint32 = 761
                    if (100 > 127) {
                        var v31 uint = 465
                        p1 = p1
                        _ = v31
                    } else {
                        var v32 uint32 = v30
                        var v33 bool = p0
                        v34 := []int{(fi28 / ((fi6) | 1)), fi28, 1, 0}
                        _ = v32
                        _ = v33
                        _ = v34
                    }
                    var v35 float32 = (1.5 * ((0.1 * 0.0) / (-2.25 * 3.14159)))
                    _ = v29
                    _ = v30
                    _ = v35
                } else {
                    var v36 uint16 = 42
                    var v37 int32 = 1
                    _ = v36
                    _ = v37
                }
                if p0 {
                    if ((p0 || p0) || false) {
                        var v38 string = (("a" + "\n\t") + "hello")
                        switch v38 {
                        case "\n\t":
                            var v39 int32 = -2147483648
                            _ = v39
                        }
                        v40 := []int{9223372036854775807, fi28, fi28, 32767}
                        v41 := GenS4{F0: (p2 | 59), F1: (p2 + 32767)}
                        _ = v38
                        _ = v40
                        _ = v41
                    }
                    p1 = p1
                    var v42 uint8 = 255
                    _ = v42
                }
                p0 = (112 >= 42)
                switch "\n\t" {
                case "":
                    p2 += (p2 ^ p2)
                case "foo bar":
                    v43 := GenS1{F0: 2, F1: ("" + "foo bar")}
                    _ = v43
                case "hello":
                    if p0 {
                        p1 = -1
                        v44 := GenS1{F0: 18446744073709551615, F1: "a"}
                        _ = v44
                    } else {
                        v45 := GenS3{F0: ("\n\t" + (("a" + "\n\t") + "a")), F1: -95, F2: (p2 & p2)}
                        if ((^(p1)) == -1) {
                            fmt.Println(fi28)
                            v46 := []int{(fi6 << ((65535) & 63)), -792, (fi6 % ((fi28) | 1))}
                            _ = v46
                        } else {
                            switch (("foo bar" + "hello") + ("foo bar" + "\n\t")) {
                            case "\n\t":
                                if p0 {
                                    v47 := []int{1, 2147483647, -2147483648, -365}
                                    v47 = append(v47, (fi6 / ((fi28) | 1)))
                                    var v48 uint8 = 42
                                    var v49 uint64 = 1
                                    fmt.Println(fi6, v48)
                                    if len(v47) >= 2 {
                                        v47 = v47[1:len(v47)]
                                    }
                                    _ = v47
                                    _ = v48
                                    _ = v49
                                } else {
                                    v50 := GenS5{F0: ((42 > 2) && (255 >= 255)), F1: 0}
                                    fmt.Println(p0, p1, p1)
                                    _ = v50
                                }
                            case "foo bar":
                                v51 := []int{fi6, -9223372036854775808}
                                _ = v51
                            default:
                                p2 = p2
                            }
                            v52 := []int{851, (fi6 << ((4294967295) & 63))}
                            _ = v52
                        }
                        v53 := []int{}
                        _ = v45
                        _ = v53
                    }
                default:
                    p0 = (3.14159 == 3.14159)
                }
                if false {
                    var v54 uint8 = 2
                    switch ("foo bar" + "") {
                    case "hello":
                        v55 := []int{}
                        _ = v55
                    case "\n\t":
                        if true {
                            v56 := []int{fi6, 127, int(v54), -2147483648}
                            if p0 {
                                v57 := GenS6{F0: 1}
                                var v58 float32 = (1.5 / 1.5)
                                _ = v57
                                _ = v58
                            }
                            v59 := GenS5{F0: p0, F1: -128}
                            _ = v56
                            _ = v59
                        } else {
                            fmt.Println(p1, p0)
                            p2 = 42
                            p2 += (p2 / ((32767) | 1))
                        }
                    case "":
                        fmt.Println(p2, v54, v54)
                    default:
                        var v60 uint64 = 0
                        _ = v60
                    }
                    var v61 uint32 = 2147483647
                    var v62 int = -855
                    if p0 {
                        fmt.Println(fi28)
                        if false {
                            v63 := []int{}
                            p1 += (p1 << ((2) & 63))
                            switch int(p2) {
                            case 42:
                                switch 128 {
                                case 174:
                                    p1 -= p1
                                case 42:
                                    p2 = p2
                                    var v64 int = int(v54)
                                    if p0 {
                                        if (!(p0 && p0)) {
                                            if (556 <= 2) {
                                                fmt.Println(fi28, fi6, p0)
                                            } else {
                                                v64 -= (-(fi28))
                                                if p0 {
                                                    v61 -= uint32(v54)
                                                    v65 := GenS7{F0: 255}
                                                    _ = v65
                                                }
                                            }
                                            v66 := GenS8{F0: p1}
                                            switch fi28 {
                                            case 1:
                                                if false {
                                                    var v67 string = ""
                                                    _ = v67
                                                } else {
                                                    if len(v63) >= 2 {
                                                        v63 = v63[1:len(v63)]
                                                    }
                                                    switch fi6 {
                                                    case 32768:
                                                        v63 = append(v63, v64)
                                                    default:
                                                        fmt.Println(fi6)
                                                    }
                                                }
                                            case -997:
                                                var v68 int8 = -1
                                                switch (v64 - 813) {
                                                case -550:
                                                    p2 = (p2 + 1)
                                                case -967:
                                                    var v69 string = "\n\t"
                                                    _ = v69
                                                case 1:
                                                    switch 9223372036854775807 {
                                                    case 127:
                                                        var v70 uint64 = 46
                                                        _ = v70
                                                    }
                                                default:
                                                    fmt.Println(fi28, p1)
                                                }
                                                _ = v68
                                            case 127:
                                                p1 += p1
                                            }
                                            _ = v66
                                        }
                                        if len(v63) > 0 {
                                            v63[0] = fi28
                                        }
                                        var v71 int64 = int64(p1)
                                        _ = v71
                                    } else {
                                        _ = len(v63)
                                        switch "a" {
                                        case "a":
                                            v64 = fi6
                                        }
                                    }
                                    _ = v64
                                }
                            }
                            _ = v63
                        }
                    } else {
                        if p0 {
                            var v72 int = v62
                            _ = v72
                        }
                        if (2 != 65535) {
                            var v73 uint16 = 255
                            if (p0 || true) {
                                p0 = (!false)
                                fmt.Println(fi6)
                                v74 := []int{fi6, fi28, 2147483647}
                                _ = v74
                            } else {
                                if false {
                                    var v75 int64 = 663
                                    var v76 int64 = (v75 - 371)
                                    _ = v75
                                    _ = v76
                                }
                                var v77 int32 = int32(p2)
                                var v78 uint8 = (v54 >> ((18446744073709551615) & 63))
                                if (!false) {
                                    v54 += 127
                                    if p0 {
                                        fmt.Println(v61, fi6, fi28)
                                        v62 += v62
                                        var v79 float64 = (-2.25 - -2.25)
                                        _ = v79
                                    }
                                }
                                v80 := []int{}
                                var v81 int32 = (v77 | p1)
                                _ = v77
                                _ = v78
                                _ = v80
                                _ = v81
                            }
                            _ = v73
                        }
                    }
                    _ = v54
                    _ = v61
                    _ = v62
                } else {
                    var v82 uint16 = 151
                    _ = v82
                }
                p2 = int16(p1)
                if (((!((false || (3.14159 <= -2.25)) && (!p0))) && (!(1.5 == float32((0.1 * 1.5))))) || p0) {
                    v83 := GenS9{F0: (0 <= 0), F1: 65535}
                    p2 = p2
                    if (255 <= 4294967295) {
                        var v84 int32 = (^(p1))
                        _ = v84
                    }
                    if (42 < 1) {
                        switch fi6 {
                        case 32768:
                            if (!p0) {
                                fmt.Println(fi28, fi6)
                                var v85 uint16 = 0
                                var v86 uint64 = 42
                                _ = v85
                                _ = v86
                            } else {
                                v87 := GenS10{F0: -1}
                                _ = v87
                            }
                        }
                    }
                    v88 := []int{}
                    _ = v83
                    _ = v88
                }
                if p0 {
                    if ((true || (-908 >= -1)) && p0) {
                        var v89 uint16 = 776
                        if true {
                            v90 := GenS5{F0: p0, F1: 100}
                            _ = v90
                        } else {
                            switch (fi28 + (fi28 << ((255) & 63))) {
                            case -932:
                                if true {
                                    var v91 int32 = -2147483648
                                    v89 = v89
                                    if (127 > 118) {
                                        if (fi6 <= (fi6 ^ (^(fi28)))) {
                                            var v92 int8 = 100
                                            var v93 uint16 = 1
                                            v94 := []int{}
                                            _ = v92
                                            _ = v93
                                            _ = v94
                                        } else {
                                            v89 = 378
                                            if true {
                                                fmt.Println(p1, v91, p1)
                                                fmt.Println(v89, p1)
                                            }
                                            v91 -= (p1 / ((p1) | 1))
                                            var v95 uint32 = 888
                                            _ = v95
                                        }
                                    }
                                    _ = v91
                                }
                            case -380:
                                v96 := GenS11{F0: 2, F1: 1}
                                _ = v96
                            case 32768:
                                if p0 {
                                    fmt.Println(fi28)
                                } else {
                                    var v97 bool = (!(p0 || (false || (!(!(p1 == p1))))))
                                    var v98 int64 = 42
                                    _ = v97
                                    _ = v98
                                }
                            default:
                                v99 := GenS7{F0: 65535}
                                _ = v99
                            }
                        }
                        _ = v89
                    }
                }
                if p0 {
                    p2 += p2
                    var v100 uint8 = 127
                    _ = v100
                }
                _ = fi28
            }
            var v101 bool = (515 <= 259)
            _ = v101
        }
        _ = fi6
    }
    for fi102 := 0; fi102 < 15; fi102++ {
        switch "" {
        case "foo bar":
            p0 = p0
        case "hello":
            if p0 {
                if (p0 && (42 == 128)) {
                    v103 := []int{fi102, -9223372036854775808, (fi102 / (((fi102 & (fi102 << ((1) & 63)))) | 1)), 128}
                    for fi104 := 0; fi104 < 11; fi104++ {
                        fmt.Println(p0)
                        p1 = p1
                        _ = fi104
                    }
                    fmt.Println(p0, fi102)
                    _ = v103
                } else {
                    fmt.Println(p2, p1)
                }
                var v105 float64 = 1e10
                _ = v105
            }
        }
        p1 -= (-(p1))
        p2 -= (p2 >> ((4294967295) & 63))
        _ = fi102
    }
    v106 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v106
    return ""
}

func main() {
    var v1 uint = 255
    v2 := []int{4294967295, 32768, -9223372036854775808, 9223372036854775807}
    var v3 uint32 = 1
    v4 := make(map[string]uint32)
    v5 := make(map[string]uint32)
    var v107 float32 = -2.25
    switch "\n\t" {
    case "":
        v107 += (v107 / v107)
    }
    _ = len(v2)
    for fi108 := 0; fi108 < 9; fi108++ {
        var v109 int16 = 1000
        for fi110 := 0; fi110 < 5; fi110++ {
            v107 = float32(1.5)
            var v111 uint8 = 1
            _ = fi110
            _ = v111
        }
        _ = fi108
        _ = v109
    }
    for fi112 := 0; fi112 < 16; fi112++ {
        var v113 float32 = (v107 - 1e10)
        v114 := GenS1{F0: (v1 | (v1 ^ (v1 * v1))), F1: "a"}
        for fi115 := 0; fi115 < 19; fi115++ {
            if ((!true) || (!false)) {
                fmt.Println(v1)
            }
            v1 -= v1
            _ = fi115
        }
        _ = fi112
        _ = v113
        _ = v114
    }
    delete(v5, ("foo bar" + ("hello" + "a")))
    var v116 string = ("hello" + "")
    var v117 float32 = 0.1
    fmt.Println("v1:", v1)
    fmt.Println("v107:", v107)
    fmt.Println("v116:", v116)
    fmt.Println("v117:", v117)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v107
    _ = v116
    _ = v117
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
