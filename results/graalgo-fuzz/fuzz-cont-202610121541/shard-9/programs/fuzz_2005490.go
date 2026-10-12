package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint32; F1 uint64; F2 int8 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int; F1 uint64 }
type GenS7 struct { F0 bool; F1 uint8 }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 bool; F1 int32; F2 uint16 }
type GenS10 struct { F0 uint64; F1 uint16 }
type GenS11 struct { F0 uint64 }
type GenS12 struct { F0 int64 }
func genFn0() uint32 {
    switch 4294967295 {
    case 536:
        if (1 >= 917) {
            v5 := GenS1{F0: 1}
            _ = v5
        } else {
            if true {
                var v6 int = 9223372036854775807
                v7 := []int{}
                v8 := GenS1{F0: 1}
                _ = v6
                _ = v7
                _ = v8
            }
            v9 := GenS1{F0: 65535}
            for fi10 := 0; fi10 < 15; fi10++ {
                var v11 int8 = 127
                fmt.Println(fi10, fi10)
                if (!false) {
                    var v12 uint64 = 2
                    _ = v12
                } else {
                    switch -373 {
                    case 9223372036854775807:
                        var v13 int = 128
                        _ = v13
                    }
                    fmt.Println(fi10, v11)
                    switch fi10 {
                    case 32768:
                        if true {
                            if true {
                                switch 4294967295 {
                                case -122:
                                    if ("hello" > "foo bar") {
                                        v11 -= 90
                                    }
                                case 4294967295:
                                    var v14 uint32 = 433
                                    _ = v14
                                case 883:
                                    v11 = v11
                                }
                                v15 := GenS1{F0: 18446744073709551615}
                                v11 = -75
                                _ = v15
                            } else {
                                v16 := []int{-400, (fi10 * 1), fi10, (fi10 / ((fi10) | 1))}
                                v11 -= (v11 * (-(v11)))
                                _ = v16
                            }
                            v11 -= v11
                        }
                    case 9223372036854775807:
                        for fi17 := 0; fi17 < 1; fi17++ {
                            v18 := GenS2{F0: 2147483648}
                            _ = fi17
                            _ = v18
                        }
                    case -1:
                        v11 = (^(v11))
                    }
                }
                _ = fi10
                _ = v11
            }
            _ = v9
        }
    case 297:
        v19 := GenS3{F0: 2147483648, F1: 18446744073709551615, F2: 0}
        _ = v19
    default:
        v20 := []int{-924, -9223372036854775808, -9223372036854775808}
        _ = v20
    }
    if true {
        if (true && true) {
            v21 := []int{42, -924, 42, 42}
            _ = v21
        }
        if (!false) {
            v22 := []int{0, 127}
            if true {
                for fi23 := 0; fi23 < 13; fi23++ {
                    var v24 uint = 733
                    if (("foo bar" == "\n\t") || false) {
                        if true {
                            v25 := GenS4{F0: 2}
                            _ = v25
                        } else {
                            for fi26 := 0; fi26 < 13; fi26++ {
                                v24 = 18446744073709551615
                                _ = fi26
                            }
                            var v27 uint = v24
                            _ = v27
                        }
                    }
                    var v28 string = (("foo bar" + ("foo bar" + "\n\t")) + ("foo bar" + ""))
                    _ = fi23
                    _ = v24
                    _ = v28
                }
            } else {
                var v29 uint = 881
                _ = v29
            }
            v22 = append(v22, 9223372036854775807)
            _ = v22
        }
        if (false || true) {
            v30 := GenS1{F0: 18446744073709551615}
            _ = v30
        } else {
            v31 := []int{-466, 127, -902, 4294967295}
            _ = v31
        }
    }
    for fi32 := 0; fi32 < 10; fi32++ {
        v33 := []int{32767}
        var v34 uint64 = 0
        _ = fi32
        _ = v33
        _ = v34
    }
    v35 := GenS1{F0: 65535}
    var v36 float32 = -2.25
    _ = v35
    _ = v36
    return 70
}

func main() {
    var v1 int8 = -101
    v2 := make(map[string]bool)
    var v3 int = 42
    var v4 bool = false
    var v37 uint8 = 128
    if (-481 != -1) {
        v38 := GenS5{F0: 381}
        if (!v4) {
            switch v3 {
            case -208:
                v37 -= 250
            case 1:
                var v39 int8 = 127
                _ = v39
            default:
                var v40 int8 = v1
                _ = v40
            }
            v1 = v1
            switch int(v37) {
            case 4294967295:
                v3 = (v3 >> ((18446744073709551615) & 63))
            case 9223372036854775807:
                var v41 uint = 255
                _ = v41
            default:
                if v4 {
                    v42 := []int{v3, (^(v3))}
                    var v43 uint = 65535
                    var v44 uint = v43
                    _ = v42
                    _ = v43
                    _ = v44
                }
            }
        } else {
            v45 := GenS3{F0: genFn0(), F1: 2, F2: (-(v1))}
            if v4 {
                v3 = (v3 & 0)
            } else {
                v46 := GenS6{F0: v3, F1: 1}
                var v47 uint = 2
                _ = v46
                _ = v47
            }
            _ = v45
        }
        _ = v38
    }
    fmt.Println(v3)
    if ((42 == 0) || (v4 || v4)) {
        if (!(-116 >= v1)) {
            var v48 int8 = (^(v1))
            _ = v48
        }
        v37 = 127
    } else {
        for fi49 := 0; fi49 < 1; fi49++ {
            if v4 {
                var v50 string = "hello"
                for fi51 := 0; fi51 < 19; fi51++ {
                    var v52 float64 = 1e10
                    _ = fi51
                    _ = v52
                }
                v53 := GenS6{F0: 4294967295, F1: 18446744073709551615}
                _ = v50
                _ = v53
            } else {
                for fi54 := 0; fi54 < 13; fi54++ {
                    v55 := GenS7{F0: (-46 == v1), F1: v37}
                    if true {
                        v56 := []int{fi54, fi54}
                        var v57 uint16 = 377
                        _ = v56
                        _ = v57
                    }
                    v3 += fi49
                    _ = fi54
                    _ = v55
                }
                if false {
                    var v58 uint = 18446744073709551615
                    if true {
                        v59 := GenS8{F0: (v58 | v58)}
                        _ = v59
                    } else {
                        fmt.Println(v4)
                        var v60 int16 = 962
                        fmt.Println(v58, v37)
                        _ = v60
                    }
                    _ = v58
                }
            }
            v37 -= (v37 ^ (v37 ^ v37))
            var v61 bool = true
            _ = fi49
            _ = v61
        }
    }
    v62 := v2["a"]
    v63 := GenS6{F0: v3, F1: 1}
    for fi64 := 0; fi64 < 12; fi64++ {
        var v65 int64 = 1
        if v62 {
            if v62 {
                v66 := GenS9{F0: (1 >= 2), F1: 2147483647, F2: 255}
                var v67 uint8 = v37
                switch ("hello" + "hello") {
                case "\n\t":
                    v37 = 0
                case "hello":
                    var v68 int32 = 0
                    _ = v68
                case "":
                    fmt.Println(v3, v65)
                }
                _ = v66
                _ = v67
            } else {
                if true {
                    fmt.Println(v3, fi64)
                    if true {
                        v1 = v1
                        for fi69 := 0; fi69 < 16; fi69++ {
                            v2[("\n\t" + ((("foo bar" + ("\n\t" + ("a" + "foo bar"))) + "a") + "foo bar"))] = (genFn0() != 2)
                            v70 := GenS9{F0: false, F1: -1, F2: 0}
                            if v62 {
                                _ = len(v2)
                                v2[("foo bar" + (("" + "\n\t") + ""))] = (fi69 == fi69)
                                if (882 >= 262) {
                                    v71 := GenS9{F0: (9223372036854775807 <= v65), F1: -2147483648, F2: 42}
                                    var v72 string = ("a" + ("\n\t" + "hello"))
                                    _ = v71
                                    _ = v72
                                } else {
                                    var v73 float32 = ((float32(0.1) + (0.0 + float32((1.5 * float64((float32(-2.25) - (-2.25 + 0.1))))))) + float32(1.5))
                                    if true {
                                        var v74 int = (^(v3))
                                        if (v62 || (-2.25 < -2.25)) {
                                            if (!("\n\t" < "hello")) {
                                                fmt.Println(v65)
                                                var v75 int = (v74 | (fi64 / ((fi69) | 1)))
                                                _ = v75
                                            } else {
                                                if v62 {
                                                    switch "" {
                                                    case "\n\t":
                                                        v76 := []int{461, 42}
                                                        _ = v76
                                                    default:
                                                        var v77 float32 = v73
                                                        _ = v77
                                                    }
                                                    fmt.Println(fi64, v74, v1)
                                                } else {
                                                    var v78 bool = false
                                                    var v79 uint = 640
                                                    fmt.Println(v37, v3, v62)
                                                    _ = v78
                                                    _ = v79
                                                }
                                            }
                                            switch v74 {
                                            case -600:
                                                fmt.Println(v65, v1, v1)
                                                if v4 {
                                                    if (((v3 << ((18446744073709551615) & 63)) <= v3) || (255 <= 4294967295)) {
                                                        v80 := GenS5{F0: -691}
                                                        v81 := GenS6{F0: 127, F1: 0}
                                                        if ((fi64 << ((0) & 63)) >= (v3 % (((v74 / ((-669) | 1))) | 1))) {
                                                            var v82 uint32 = 258
                                                            _ = v82
                                                        } else {
                                                            fmt.Println(v73)
                                                            var v83 int32 = -2147483648
                                                            v65 += (-(v65))
                                                            _ = v83
                                                        }
                                                        _ = v80
                                                        _ = v81
                                                    }
                                                    v2["foo bar"] = ((1 == 42) && v4)
                                                    var v84 uint = 65535
                                                    _ = v84
                                                } else {
                                                    v73 += float32((-2.25 + 0.0))
                                                }
                                                fmt.Println(v73, v37, fi64)
                                                switch v3 {
                                                case 2147483647:
                                                    var v85 int16 = 449
                                                    _ = v85
                                                case -9223372036854775808:
                                                    v65 = -184
                                                case 429:
                                                    var v86 uint64 = 270
                                                    v87 := GenS10{F0: (v86 & 42), F1: 255}
                                                    v88 := GenS1{F0: 42}
                                                    var v89 bool = v4
                                                    if (v73 > ((float32(-2.25) - (((1e10 * ((v73 * v73) + v73)) * float32(((3.14159 + 3.14159) - 3.14159))) + v73)) - (v73 + v73))) {
                                                        v90 := GenS7{F0: true, F1: 42}
                                                        delete(v2, "")
                                                        v91 := v2[("" + ((("" + "a") + "foo bar") + "hello"))]
                                                        _ = v90
                                                        _ = v91
                                                    }
                                                    _ = v86
                                                    _ = v87
                                                    _ = v88
                                                    _ = v89
                                                default:
                                                    var v92 uint64 = 18446744073709551615
                                                    _ = v92
                                                }
                                            case 216:
                                                v93 := GenS2{F0: genFn0()}
                                                _ = v93
                                            default:
                                                var v94 bool = v4
                                                _ = v94
                                            }
                                        }
                                        _ = v74
                                    } else {
                                        v95 := []int{}
                                        _ = v95
                                    }
                                    _ = v73
                                }
                            } else {
                                fmt.Println(v4, v62, v1)
                            }
                            var v96 float32 = 0.1
                            _ = fi69
                            _ = v70
                            _ = v96
                        }
                        if false {
                            var v97 float32 = ((1e10 - (0.0 + 1.5)) + (((1e10 + float32(0.1)) - (1.5 / 1e10)) + 1.5))
                            _ = v97
                        }
                    }
                }
            }
            v98 := v2["a"]
            _ = v98
        }
        _ = fi64
        _ = v65
    }
    fmt.Println(v37, v1, v62)
    var v99 int32 = -589
    v100 := []int{v3, 4294967295}
    v101 := GenS11{F0: uint64(v37)}
    delete(v2, "foo bar")
    v99 = v99
    for fi102 := 0; fi102 < 19; fi102++ {
        for fi103 := 0; fi103 < 10; fi103++ {
            v104 := GenS9{F0: true, F1: (v99 >> ((2) & 63)), F2: 2}
            fmt.Println(v3, v3, v37)
            _ = fi103
            _ = v104
        }
        for fi105 := 0; fi105 < 15; fi105++ {
            var v106 string = "foo bar"
            var v107 int64 = 9223372036854775807
            switch v106 {
            case "\n\t":
                if (4294967295 == 0) {
                    v99 = 1
                    _ = len(v100)
                } else {
                    switch v106 {
                    case "\n\t":
                        delete(v2, "hello")
                    case "":
                        var v108 float32 = (((3.14159 + -2.25) - 0.0) / ((0.0 - (1e10 - 0.1)) * ((1.5 - 3.14159) * 1e10)))
                        _ = v108
                    case "hello":
                        var v109 uint8 = v37
                        _ = v109
                    }
                    fmt.Println(fi105, v106)
                }
            }
            _ = fi105
            _ = v106
            _ = v107
        }
        for fi110 := 0; fi110 < 10; fi110++ {
            if (!((810 == 924) && (3.14159 > 0.1))) {
                v3 += (v3 << ((255) & 63))
                v1 = (^(v1))
            } else {
                if v62 {
                    var v111 uint16 = 255
                    switch v3 {
                    case 42:
                        var v112 float64 = (3.14159 * float64(3.14159))
                        _ = v112
                    case 4294967295:
                        switch "" {
                        case "foo bar":
                            if len(v100) >= 2 {
                                v100 = v100[1:len(v100)]
                            }
                        case "":
                            _ = len(v100)
                        case "a":
                            var v113 uint16 = (v111 << ((65535) & 63))
                            _ = v113
                        }
                    case 0:
                        if (true || true) {
                            v114 := GenS12{F0: 9223372036854775807}
                            _ = len(v100)
                            v115 := GenS2{F0: genFn0()}
                            _ = v114
                            _ = v115
                        }
                    default:
                        v37 -= (v37 * v37)
                    }
                    _ = v111
                }
                var v116 int16 = 0
                _ = v116
            }
            var v117 int64 = 1
            var v118 float32 = (1e10 / ((0.1 * ((0.1 - 3.14159) * -2.25)) + 1.5))
            _ = fi110
            _ = v117
            _ = v118
        }
        _ = fi102
    }
    var v119 int = -381
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v100:", len(v100), v100)
    fmt.Println("v119:", v119)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v62:", v62)
    fmt.Println("v99:", v99)
    _ = v1
    _ = v100
    _ = v101
    _ = v119
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v62
    _ = v63
    _ = v99
}
