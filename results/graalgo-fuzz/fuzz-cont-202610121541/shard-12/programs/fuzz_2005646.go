package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8; F2 string }
type GenS2 struct { F0 uint; F1 uint8 }
type GenS3 struct { F0 int16; F1 int32; F2 uint16 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 string; F1 uint; F2 int8 }
type GenS6 struct { F0 uint32; F1 int64 }
type GenS7 struct { F0 uint64; F1 uint64; F2 uint8 }
type GenS8 struct { F0 int8; F1 uint; F2 uint16 }
type GenS9 struct { F0 int64 }
type GenS10 struct { F0 bool }
type GenS11 struct { F0 int32; F1 int8 }
type GenS12 struct { F0 int16 }
type GenS13 struct { F0 int64 }
type GenS14 struct { F0 int8; F1 uint }
func genFn0(p0 uint64) uint64 {
    v5 := []int{-1, 127}
    _ = len(v5)
    var v6 float32 = 1.5
    var v7 int16 = 1
    for fi8 := 0; fi8 < 2; fi8++ {
        for fi9 := 0; fi9 < 7; fi9++ {
            if false {
                var v10 bool = ("hello" == ("a" + "a"))
                v11 := GenS1{F0: 42, F1: 127, F2: ("" + "")}
                _ = v10
                _ = v11
            } else {
                v6 = -2.25
            }
            v5 = append(v5, fi8)
            _ = fi9
        }
        switch 0 {
        case 128:
            if len(v5) > 0 {
                v5[0] = fi8
            }
        }
        _ = fi8
    }
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return (p0 / ((4294967295) | 1))
}
func genFn1(p0 int8, p1 uint32) uint16 {
    var v12 int8 = (p0 & (^(p0)))
    var v13 bool = (-128 >= p0)
    v14 := []int{int(p1), -1, -1, -167}
    _ = p0
    _ = p1
    _ = v12
    _ = v13
    _ = v14
    return 755
}

func main() {
    var v1 int8 = 42
    var v2 uint = 789
    v3 := make(map[string]bool)
    var v4 bool = false
    fmt.Println(v2, v2)
    if v4 {
        v15 := []int{9223372036854775807, 9223372036854775807, 127}
        for fi16 := 0; fi16 < 13; fi16++ {
            switch ("hello" + ("" + "\n\t")) {
            case "\n\t":
                v17 := GenS2{F0: (v2 & v2), F1: 255}
                _ = v17
            }
            for fi18 := 0; fi18 < 11; fi18++ {
                if (0 < 42) {
                    v4 = (!(0.1 <= 0.1))
                    switch fi18 {
                    case 1:
                        v4 = false
                    }
                    if false {
                        v4 = (!v4)
                    } else {
                        var v19 uint = 4294967295
                        var v20 string = ""
                        v1 = (v1 | v1)
                        _ = v19
                        _ = v20
                    }
                } else {
                    var v21 int = fi16
                    delete(v3, "hello")
                    if true {
                        if ((18446744073709551615 > 552) && ("foo bar" > ("a" + "\n\t"))) {
                            if (!false) {
                                var v22 float32 = 1.5
                                if false {
                                    v1 = (-(v1))
                                    v4 = true
                                }
                                fmt.Println(fi18)
                                _ = v22
                            } else {
                                var v23 uint16 = genFn1(v1, 42)
                                switch (fi16 >> ((v2) & 63)) {
                                case -77:
                                    switch (fi18 - fi18) {
                                    case 42:
                                        if len(v15) > 0 {
                                            v15[0] = 254
                                        }
                                    case 721:
                                        switch ((("\n\t" + ("\n\t" + "\n\t")) + ("\n\t" + "hello")) + ("" + "foo bar")) {
                                        case "foo bar":
                                            switch (fi18 | 9223372036854775807) {
                                            case 32767:
                                                if (true && false) {
                                                    var v24 int16 = 925
                                                    v1 = (^(v1))
                                                    _ = v24
                                                } else {
                                                    v1 += v1
                                                    fmt.Println(v23, v1)
                                                    v1 += (v1 ^ (v1 << ((v2) & 63)))
                                                    v4 = v4
                                                    v25 := GenS3{F0: 42, F1: 198, F2: v23}
                                                    _ = v25
                                                }
                                            case 32768:
                                                v26 := GenS2{F0: 0, F1: 127}
                                                _ = v26
                                            case 898:
                                                _ = len(v15)
                                            }
                                        case "\n\t":
                                            fmt.Println(v4, v21)
                                        case "a":
                                            var v27 uint16 = 2
                                            _ = v27
                                        default:
                                            if v4 {
                                                if v4 {
                                                    v28 := GenS3{F0: int16(v2), F1: -2147483648, F2: (^(v23))}
                                                    var v29 uint8 = 127
                                                    _ = v28
                                                    _ = v29
                                                }
                                                if len(v15) > 0 {
                                                    v15[0] = fi16
                                                }
                                                var v30 int16 = 199
                                                _ = v30
                                            }
                                        }
                                    }
                                    switch (v21 - v21) {
                                    case 9223372036854775807:
                                        if (!(-2.25 <= 0.0)) {
                                            fmt.Println(v21, v21)
                                            fmt.Println(v4, v4)
                                            _ = len(v15)
                                        } else {
                                            var v31 float64 = 0.0
                                            var v32 int8 = int8(fi18)
                                            v4 = (!false)
                                            v32 = v1
                                            var v33 uint32 = 0
                                            _ = v31
                                            _ = v32
                                            _ = v33
                                        }
                                    case 173:
                                        v34 := GenS1{F0: 1, F1: 42, F2: ("\n\t" + "\n\t")}
                                        _ = v34
                                    default:
                                        var v35 uint = 575
                                        _ = v35
                                    }
                                default:
                                    v21 -= fi16
                                    var v36 uint64 = 536
                                    if v4 {
                                        switch (v21 >> (((v2 - 897)) & 63)) {
                                        case -1:
                                            fmt.Println(v36, v21)
                                        case 47:
                                            switch ((("foo bar" + ("\n\t" + "hello")) + "") + "foo bar") {
                                            case "a":
                                                var v37 string = "a"
                                                _ = v37
                                            }
                                        case -224:
                                            v38 := GenS4{F0: uint16(v2)}
                                            _ = v38
                                        default:
                                            switch "" {
                                            case "a":
                                                var v39 uint32 = 2147483648
                                                _ = v39
                                            case "":
                                                if (192 >= 32767) {
                                                    if len(v15) >= 2 {
                                                        v15 = v15[1:len(v15)]
                                                    }
                                                } else {
                                                    v36 = 643
                                                    if ((0.1 / -2.25) > 3.14159) {
                                                        v40 := GenS5{F0: ("" + "foo bar"), F1: 42, F2: 1}
                                                        v41 := GenS6{F0: 42, F1: 1}
                                                        v42 := v3[(("" + "\n\t") + "a")]
                                                        v1 += 100
                                                        fmt.Println(v36, v42, v4)
                                                        fmt.Println(v1)
                                                        var v43 int64 = 0
                                                        _ = v40
                                                        _ = v41
                                                        _ = v42
                                                        _ = v43
                                                    } else {
                                                        v1 = -51
                                                        _ = len(v15)
                                                        v44 := GenS7{F0: genFn0((^(v36))), F1: (v36 >> (((v2 + (v2 << ((uint(v23)) & 63)))) & 63)), F2: 1}
                                                        _ = v44
                                                    }
                                                    _ = len(v3)
                                                    var v45 uint = 4294967295
                                                    _ = v45
                                                }
                                            }
                                        }
                                        var v46 bool = v4
                                        fmt.Println(v21)
                                        v46 = (42 < 1)
                                        if (4294967295 < 704) {
                                            if v46 {
                                                if v46 {
                                                    var v47 int8 = (v1 * (v1 ^ (v1 + 4)))
                                                    var v48 int64 = int64(v23)
                                                    v1 = v1
                                                    _ = v47
                                                    _ = v48
                                                } else {
                                                    fmt.Println(v4, v23, v21)
                                                    var v49 uint16 = (v23 ^ genFn1((v1 + v1), 782))
                                                    var v50 int8 = 127
                                                    _ = v49
                                                    _ = v50
                                                }
                                                var v51 int = fi16
                                                if v46 {
                                                    fmt.Println(v1, v36)
                                                }
                                                _ = v51
                                            } else {
                                                var v52 int32 = -2147483648
                                                fmt.Println(v23, fi16)
                                                _ = v52
                                            }
                                            if v4 {
                                                var v53 uint32 = 1
                                                var v54 uint = v2
                                                var v55 float64 = 3.14159
                                                v1 = int8(v23)
                                                if (45 > 255) {
                                                    _ = len(v3)
                                                }
                                                v1 = v1
                                                v56 := GenS4{F0: genFn1((v1 - (v1 % (((v1 % ((v1) | 1))) | 1))), (v53 * (^(v53))))}
                                                if (((("a" + "") + "") + ("hello" + ("" + "foo bar"))) < "hello") {
                                                    fmt.Println(v1)
                                                    v36 = uint64(v53)
                                                    v57 := GenS7{F0: v36, F1: (v36 / ((1) | 1)), F2: 42}
                                                    _ = v57
                                                }
                                                if (18446744073709551615 != (v36 - (v36 ^ v36))) {
                                                    v53 = (v53 >> ((255) & 63))
                                                    v58 := GenS8{F0: (v1 + v1), F1: uint(v36), F2: 1}
                                                    v59 := GenS1{F0: 0, F1: 100, F2: ("foo bar" + "hello")}
                                                    _ = v58
                                                    _ = v59
                                                }
                                                _ = v53
                                                _ = v54
                                                _ = v55
                                                _ = v56
                                            } else {
                                                if (!v4) {
                                                    delete(v3, (("a" + "a") + "\n\t"))
                                                } else {
                                                    switch v21 {
                                                    case -178:
                                                        if len(v15) >= 2 {
                                                            v15 = v15[1:len(v15)]
                                                        }
                                                    case 128:
                                                        if v46 {
                                                            if (!v4) {
                                                                var v60 int32 = 42
                                                                _ = v60
                                                            } else {
                                                                if false {
                                                                    var v61 string = ("hello" + "")
                                                                    if v4 {
                                                                        v62 := GenS9{F0: 0}
                                                                        var v63 float32 = 1e10
                                                                        fmt.Println(v23, v2)
                                                                        v1 += v1
                                                                        var v64 string = ""
                                                                        _ = v62
                                                                        _ = v63
                                                                        _ = v64
                                                                    }
                                                                    _ = v61
                                                                } else {
                                                                    if false {
                                                                        var v65 float32 = float32(3.14159)
                                                                        _ = v65
                                                                    } else {
                                                                        if v4 {
                                                                            fmt.Println(v36, v4)
                                                                            v66 := GenS4{F0: v23}
                                                                            var v67 int8 = -128
                                                                            _ = v66
                                                                            _ = v67
                                                                        }
                                                                        v36 = v36
                                                                    }
                                                                }
                                                                if v46 {
                                                                    v21 += fi18
                                                                    v46 = (true && false)
                                                                    var v68 string = ("" + "hello")
                                                                    _ = v68
                                                                } else {
                                                                    v46 = ((!false) || (!true))
                                                                }
                                                            }
                                                            if false {
                                                                _ = len(v3)
                                                                v15 = append(v15, (v21 + 32767))
                                                            } else {
                                                                var v69 string = (("foo bar" + ((("a" + "\n\t") + "foo bar") + ("\n\t" + ""))) + "foo bar")
                                                                _ = v69
                                                            }
                                                            v70 := GenS10{F0: v4}
                                                            _ = v70
                                                        } else {
                                                            var v71 float64 = ((1.5 + 1.5) * 1.5)
                                                            var v72 string = ("hello" + "\n\t")
                                                            if v46 {
                                                                if true {
                                                                    if v46 {
                                                                        v4 = (2 >= 2)
                                                                        if len(v15) >= 2 {
                                                                            v15 = v15[1:len(v15)]
                                                                        }
                                                                    } else {
                                                                        v1 = (v1 >> ((v2) & 63))
                                                                        v73 := GenS7{F0: (v36 & 1), F1: (^(v36)), F2: 255}
                                                                        v3[v72] = (v71 != 0.1)
                                                                        _ = v73
                                                                    }
                                                                    v72 = "hello"
                                                                    switch v72 {
                                                                    case "hello":
                                                                        v4 = false
                                                                    case "\n\t":
                                                                        var v74 float64 = ((-2.25 - v71) + 1.5)
                                                                        _ = v74
                                                                    case "":
                                                                        fmt.Println(v71, v46, v46)
                                                                    default:
                                                                        if len(v15) > 0 {
                                                                            v15[0] = 0
                                                                        }
                                                                    }
                                                                    v75 := GenS4{F0: (^(v23))}
                                                                    v4 = ((!true) && v4)
                                                                    fmt.Println(v2, fi16)
                                                                    _ = v75
                                                                } else {
                                                                    switch (("a" + ((((v72 + (("\n\t" + (v72 + v72)) + ((("hello" + "a") + "") + v72))) + (("foo bar" + (("\n\t" + (v72 + "\n\t")) + "hello")) + "")) + "hello") + (v72 + v72))) + v72) {
                                                                    case "\n\t":
                                                                        v2 = (v2 ^ uint(v23))
                                                                    case "":
                                                                        v76 := v3["hello"]
                                                                        _ = v76
                                                                    case "hello":
                                                                        v46 = v4
                                                                    }
                                                                    v2 += v2
                                                                    v1 = 1
                                                                }
                                                                if (!(int(v36) <= (^(v21)))) {
                                                                    var v77 float32 = ((((0.1 - (float32(float64(0.1)) / 1.0)) * (-2.25 / 3.14159)) + 3.14159) / (1.5 - 0.1))
                                                                    var v78 int64 = 9223372036854775807
                                                                    if true {
                                                                        fmt.Println(fi18, v78)
                                                                        fmt.Println(v77, v71, v78)
                                                                    } else {
                                                                        if v46 {
                                                                            if v4 {
                                                                                var v79 int = (fi18 & (fi18 << ((v2) & 63)))
                                                                                var v80 float64 = v71
                                                                                v2 = 18446744073709551615
                                                                                _ = v79
                                                                                _ = v80
                                                                            }
                                                                            var v81 uint32 = 1
                                                                            switch v21 {
                                                                            case 42:
                                                                                if ((!(!(((v23 + (v23 % ((genFn1(42, (v81 ^ (^(v81))))) | 1))) == v23) && v4))) && (4294967295 >= uint32(v23))) {
                                                                                    var v82 int64 = (v78 & v78)
                                                                                    _ = v82
                                                                                } else {
                                                                                    v83 := GenS11{F0: 1, F1: (-(v1))}
                                                                                    _ = v83
                                                                                }
                                                                            }
                                                                            _ = v81
                                                                        }
                                                                        v72 = ((v72 + v72) + ("hello" + ((v72 + "hello") + (((v72 + "hello") + v72) + ((("" + "a") + v72) + ((v72 + "foo bar") + (v72 + ("hello" + v72))))))))
                                                                        if len(v15) > 0 {
                                                                            v15[0] = fi16
                                                                        }
                                                                        var v84 int = 9223372036854775807
                                                                        switch ((v72 + ("hello" + v72)) + v72) {
                                                                        case "\n\t":
                                                                            if false {
                                                                                if (v4 && v46) {
                                                                                    switch fi16 {
                                                                                    case 765:
                                                                                        switch fi16 {
                                                                                        case 1:
                                                                                            var v85 uint32 = 4294967295
                                                                                            _ = v85
                                                                                        case 127:
                                                                                            var v86 int16 = -430
                                                                                            _ = v86
                                                                                        }
                                                                                    default:
                                                                                        v3[((v72 + (v72 + (v72 + v72))) + "foo bar")] = v4
                                                                                        var v87 uint64 = v36
                                                                                        if ((42 < 110) || (!(!((v87 & 930) <= v36)))) {
                                                                                            if (65535 >= uint32(v87)) {
                                                                                                v15 = append(v15, 25)
                                                                                                v88 := GenS6{F0: 65535, F1: v78}
                                                                                                v89 := GenS8{F0: v1, F1: 1, F2: 394}
                                                                                                _ = v88
                                                                                                _ = v89
                                                                                            } else {
                                                                                                var v90 uint16 = (v23 >> (((v2 + v2)) & 63))
                                                                                                fmt.Println(v1, v77, v71)
                                                                                                v91 := GenS9{F0: (v78 ^ (v78 % ((v78) | 1)))}
                                                                                                _ = v90
                                                                                                _ = v91
                                                                                            }
                                                                                            v23 = (v23 >> ((v2) & 63))
                                                                                        }
                                                                                        _ = v87
                                                                                    }
                                                                                }
                                                                            } else {
                                                                                var v92 uint8 = 109
                                                                                _ = v92
                                                                            }
                                                                        case "hello":
                                                                            delete(v3, (((v72 + v72) + (v72 + (v72 + "\n\t"))) + (v72 + "a")))
                                                                        }
                                                                        _ = v84
                                                                    }
                                                                    _ = v77
                                                                    _ = v78
                                                                } else {
                                                                    v93 := v3[v72]
                                                                    fmt.Println(v2, v93, v46)
                                                                    _ = v93
                                                                }
                                                                v2 += v2
                                                            } else {
                                                                v94 := GenS2{F0: v2, F1: 127}
                                                                fmt.Println(v71, v2)
                                                                _ = v94
                                                            }
                                                            _ = v71
                                                            _ = v72
                                                        }
                                                        var v95 int = fi16
                                                        _ = v95
                                                    }
                                                    if ((!true) && (false && (!v4))) {
                                                        v1 -= v1
                                                        var v96 float32 = (float32(1.5) - 0.1)
                                                        _ = v96
                                                    } else {
                                                        switch (fi18 % (((fi16 & (v21 ^ fi16))) | 1)) {
                                                        case 419:
                                                            v4 = v4
                                                        default:
                                                            switch "" {
                                                            case "a":
                                                                switch "foo bar" {
                                                                case "a":
                                                                    v23 = v23
                                                                case "foo bar":
                                                                    var v97 int32 = -1
                                                                    _ = v97
                                                                case "\n\t":
                                                                    var v98 float32 = -2.25
                                                                    _ = v98
                                                                }
                                                            }
                                                        }
                                                        v99 := GenS7{F0: v36, F1: v36, F2: 172}
                                                        _ = v99
                                                    }
                                                    v36 -= v36
                                                }
                                                switch "a" {
                                                case "":
                                                    v1 += (v1 * (v1 & 127))
                                                }
                                                if (!v46) {
                                                    var v100 uint64 = 42
                                                    _ = v100
                                                } else {
                                                    v23 = (^(v23))
                                                    switch "" {
                                                    case "a":
                                                        v3[""] = v4
                                                        delete(v3, ("" + ""))
                                                        var v101 int64 = 277
                                                        _ = v101
                                                    case "hello":
                                                        switch "a" {
                                                        case "foo bar":
                                                            _ = len(v3)
                                                        case "":
                                                            if false {
                                                                if ((!(false || v4)) || ((0.1 - 0.0) <= 1e10)) {
                                                                    var v102 bool = (-1 <= -2147483648)
                                                                    v103 := GenS1{F0: 188, F1: 0, F2: "a"}
                                                                    _ = v102
                                                                    _ = v103
                                                                }
                                                                switch ("" + "hello") {
                                                                case "foo bar":
                                                                    fmt.Println(v23)
                                                                case "a":
                                                                    var v104 uint64 = v36
                                                                    _ = v104
                                                                default:
                                                                    v21 += fi16
                                                                    _ = len(v15)
                                                                    if true {
                                                                        fmt.Println(v4, v23, v46)
                                                                        v15 = append(v15, 42)
                                                                    }
                                                                }
                                                                var v105 uint8 = 2
                                                                v106 := GenS12{F0: 32767}
                                                                _ = v105
                                                                _ = v106
                                                            }
                                                        default:
                                                            var v107 bool = v4
                                                            _ = v107
                                                        }
                                                    case "foo bar":
                                                        var v108 int64 = 0
                                                        _ = v108
                                                    }
                                                }
                                            }
                                            var v109 int8 = (v1 % ((v1) | 1))
                                            if v4 {
                                                v109 = (v1 >> (((^(v2))) & 63))
                                                v110 := GenS7{F0: 1, F1: v36, F2: 59}
                                                _ = len(v3)
                                                _ = v110
                                            }
                                            var v111 int8 = v109
                                            v46 = (v46 || (!((!true) || (0 != 0))))
                                            switch (fi16 - fi18) {
                                            case 42:
                                                if (fi16 <= 1) {
                                                    if true {
                                                        v4 = (7 == 130)
                                                        v111 += (v1 / (((-(v1))) | 1))
                                                    }
                                                }
                                            }
                                            _ = v109
                                            _ = v111
                                        }
                                        _ = v46
                                    }
                                    _ = v36
                                }
                                _ = v23
                            }
                            var v112 string = "a"
                            _ = v112
                        } else {
                            if (1 != genFn0(2)) {
                                var v113 string = "\n\t"
                                var v114 uint64 = 4294967295
                                v4 = v4
                                if true {
                                    var v115 uint8 = 68
                                    if (!(false && ((!v4) || (255 == (v115 - 250))))) {
                                        fmt.Println(v4, v1, v21)
                                    } else {
                                        v116 := GenS13{F0: -9223372036854775808}
                                        v117 := GenS14{F0: (v1 + (^(v1))), F1: (v2 / (((v2 * (v2 * v2))) | 1))}
                                        if len(v15) > 0 {
                                            v15[0] = fi16
                                        }
                                        _ = v116
                                        _ = v117
                                    }
                                    if len(v15) > 0 {
                                        v15[0] = (-(fi16))
                                    }
                                    _ = v115
                                } else {
                                    fmt.Println(v113, v21, fi16)
                                    v118 := v3[v113]
                                    if v118 {
                                        var v119 uint8 = 255
                                        _ = v119
                                    }
                                    _ = v118
                                }
                                _ = v113
                                _ = v114
                            } else {
                                if true {
                                    v120 := GenS12{F0: -337}
                                    _ = len(v15)
                                    v3["hello"] = ((205 <= 729) && v4)
                                    var v121 uint64 = 4294967295
                                    _ = v120
                                    _ = v121
                                } else {
                                    var v122 int8 = 2
                                    fmt.Println(v122, v4, v122)
                                    _ = v122
                                }
                                v4 = v4
                                v2 -= (v2 - v2)
                                var v123 int32 = 1
                                v21 = (v21 << ((42) & 63))
                                _ = v123
                            }
                        }
                        v124 := v3["hello"]
                        v125 := GenS5{F0: ("" + "a"), F1: v2, F2: v1}
                        _ = v124
                        _ = v125
                    } else {
                        v2 = 0
                        if ((2147483647 == 1) || (32767 == 34)) {
                            v4 = true
                        } else {
                            var v126 int = (v21 * fi18)
                            _ = v126
                        }
                        v2 = 42
                    }
                    _ = v21
                }
                if len(v15) > 0 {
                    v15[0] = fi16
                }
                var v127 int8 = v1
                _ = fi18
                _ = v127
            }
            _ = fi16
        }
        _ = v15
    }
    for fi128 := 0; fi128 < 16; fi128++ {
        fmt.Println(v4, fi128, fi128)
        _ = fi128
    }
    fmt.Println(v1, v4)
    _ = len(v3)
    if (!((!(1 <= 837)) && v4)) {
        var v129 int = 2147483647
        v1 = v1
        v1 = 127
        _ = v129
    }
    var v130 int8 = v1
    v130 = (-(v130))
    for fi131 := 0; fi131 < 11; fi131++ {
        switch fi131 {
        case 32768:
            var v132 uint = v2
            _ = v132
        default:
            v4 = v4
        }
        _ = fi131
    }
    v2 = (^(v2))
    v4 = v4
    v4 = v4
    v133 := []int{1, 32767, 2147483647, 2147483647}
    v2 = v2
    v2 = 0
    for fi134 := 0; fi134 < 9; fi134++ {
        v1 = v1
        _ = fi134
    }
    if len(v133) > 0 {
        v133[0] = -226
    }
    fmt.Println("v1:", v1)
    fmt.Println("v130:", v130)
    fmt.Println("slicelen v133:", len(v133), v133)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v130
    _ = v133
    _ = v2
    _ = v3
    _ = v4
}
