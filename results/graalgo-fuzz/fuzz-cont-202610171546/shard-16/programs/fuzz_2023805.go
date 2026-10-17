package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8; F2 string }
type GenS2 struct { F0 int; F1 int32; F2 string }
type GenS3 struct { F0 int16; F1 int16; F2 bool }
type GenS4 struct { F0 uint8; F1 uint8 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int64; F1 int; F2 uint16 }
type GenS7 struct { F0 int32; F1 uint32 }
type GenS8 struct { F0 int64; F1 string; F2 bool }
type GenS9 struct { F0 int8; F1 uint32; F2 uint32 }
type GenS10 struct { F0 uint32; F1 int32; F2 int16 }
type GenS11 struct { F0 uint16; F1 uint8 }
type GenS12 struct { F0 uint8; F1 int32 }
type GenS13 struct { F0 bool }
type GenS14 struct { F0 uint16 }
type GenS15 struct { F0 uint16; F1 int32; F2 int16 }
type GenS16 struct { F0 bool; F1 bool; F2 string }
func genFn0(p0 uint8) uint8 {
    var v5 uint64 = 0
    v6 := []int{-774, -363, 1}
    v7 := GenS1{F0: 1, F1: 8, F2: "hello"}
    if ((false || true) || (false || ((-62 >= 42) && false))) {
        v8 := GenS1{F0: v5, F1: -1, F2: ("a" + "a")}
        _ = v8
    } else {
        if (!(2 <= 303)) {
            fmt.Println(p0, v5)
        }
        var v9 int32 = 0
        p0 += (p0 ^ (p0 % (((p0 % ((p0) | 1))) | 1)))
        _ = v9
    }
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return p0
}

func main() {
    v1 := []int{62, -1}
    v2 := []int{2147483647, 32767, 0, 1}
    v3 := []uint32{65535}
    v4 := make(map[string]uint32)
    var v10 uint32 = 1
    for fi11 := 0; fi11 < 6; fi11++ {
        for fi12 := 0; fi12 < 11; fi12++ {
            _ = len(v3)
            v13 := GenS1{F0: 1, F1: int8(v10), F2: ""}
            v14 := GenS1{F0: 4294967295, F1: 19, F2: "a"}
            _ = fi12
            _ = v13
            _ = v14
        }
        _ = fi11
    }
    var v15 float64 = (1e10 - 0.0)
    var v16 int16 = -32768
    fmt.Println(v16, v15, v15)
    var v17 string = ""
    switch -9223372036854775808 {
    case 32767:
        v18 := GenS2{F0: -56, F1: 606, F2: v17}
        _ = v18
    case 879:
        for fi19 := 0; fi19 < 19; fi19++ {
            fmt.Println(v17, v17, v15)
            var v20 float64 = float64(1.5)
            _ = fi19
            _ = v20
        }
    case 697:
        v21 := GenS3{F0: (v16 - v16), F1: (v16 - v16), F2: (469 != 65535)}
        _ = v21
    }
    for fi22 := 0; fi22 < 15; fi22++ {
        for fi23 := 0; fi23 < 15; fi23++ {
            if (true && true) {
                var v24 uint16 = 936
                _ = v24
            } else {
                if (float64((1e10 / -2.25)) != v15) {
                    v15 = v15
                }
                fmt.Println(v15)
            }
            switch 206 {
            case 360:
                v4[("foo bar" + "hello")] = v10
            case -1:
                fmt.Println(v15, v15)
            case -876:
                if ((!false) && false) {
                    var v25 int8 = 127
                    _ = v25
                }
            default:
                switch v17 {
                case "hello":
                    v15 += ((-2.25 + (1.5 - ((float64(0.1) + v15) + (float64((3.14159 / 1.0)) + (((((v15 / float64((1e10 * 1e10))) - (1.5 * v15)) - (float64(3.14159) - (((1e10 / 3.14159) - (float64(0.0) - 1.5)) * 3.14159))) - 3.14159) + v15))))) - float64(((1e10 - 0.1) / -2.25)))
                case "a":
                    if false {
                        v26 := GenS1{F0: 4294967295, F1: -128, F2: "foo bar"}
                        v10 += v10
                        v27 := GenS4{F0: 42, F1: 2}
                        _ = v26
                        _ = v27
                    } else {
                        v2 = append(v2, (fi23 >> ((384) & 63)))
                        if (0 > 171) {
                            switch (fi23 / ((2147483647) | 1)) {
                            case 2147483647:
                                v10 = v10
                            case 127:
                                var v28 uint32 = v10
                                _ = v28
                            case 4294967295:
                                switch ("foo bar" + v17) {
                                case "\n\t":
                                    var v29 int32 = 2147483647
                                    _ = v29
                                case "foo bar":
                                    v30 := GenS5{F0: (fi22 / (((^(fi22))) | 1))}
                                    _ = v30
                                case "hello":
                                    v16 = (-(v16))
                                default:
                                    v16 = v16
                                }
                            }
                        } else {
                            if (-9223372036854775808 < fi23) {
                                v16 += (v16 * (v16 & (^(v16))))
                                fmt.Println(v17, fi23)
                            }
                            fmt.Println(v10, v16)
                        }
                        switch fi22 {
                        case 127:
                            if (v10 < 2147483647) {
                                v31 := GenS3{F0: (v16 * (v16 * int16(v10))), F1: (v16 | (^(v16))), F2: true}
                                _ = v31
                            } else {
                                if (!true) {
                                    fmt.Println(v16, v17)
                                    if (genFn0(2) >= 132) {
                                        switch fi23 {
                                        case 875:
                                            if true {
                                                switch v17 {
                                                case "foo bar":
                                                    v16 += (v16 & v16)
                                                case "hello":
                                                    if (!true) {
                                                        v15 += float64((0.0 - 1e10))
                                                        var v32 float32 = 1.5
                                                        v33 := GenS5{F0: 127}
                                                        _ = v32
                                                        _ = v33
                                                    } else {
                                                        v17 = "a"
                                                        var v34 uint32 = 42
                                                        _ = v34
                                                    }
                                                default:
                                                    fmt.Println(v16, v10)
                                                }
                                            } else {
                                                switch fi22 {
                                                case 4294967295:
                                                    switch ("" + (v17 + v17)) {
                                                    case "\n\t":
                                                        v16 += (^(v16))
                                                    case "hello":
                                                        v17 = (v17 + ("a" + (v17 + v17)))
                                                    case "":
                                                        var v35 float64 = v15
                                                        _ = v35
                                                    }
                                                }
                                            }
                                        default:
                                            v15 -= ((1e10 / (float64(((3.14159 - (0.0 - 1e10)) / 1.5)) / v15)) - v15)
                                        }
                                        if len(v1) > 0 {
                                            v1[0] = (fi23 << ((255) & 63))
                                        }
                                        v15 = 0.0
                                    } else {
                                        fmt.Println(v10)
                                        switch fi22 {
                                        case -9223372036854775808:
                                            if len(v3) > 0 {
                                                v3[0] = (^(v10))
                                            }
                                        default:
                                            var v36 string = (v17 + v17)
                                            _ = v36
                                        }
                                        if (783 > 0) {
                                            if ((!true) || (1 == 5)) {
                                                switch v17 {
                                                case "":
                                                    if true {
                                                        fmt.Println(v15, fi23, v10)
                                                        v16 += v16
                                                        delete(v4, v17)
                                                    }
                                                case "\n\t":
                                                    v37 := GenS2{F0: (fi23 / (((fi22 - (fi23 & (fi22 % ((int(v10)) | 1))))) | 1)), F1: -490, F2: "foo bar"}
                                                    _ = v37
                                                }
                                                switch -234 {
                                                case 127:
                                                    v15 -= (1e10 + v15)
                                                case 32768:
                                                    var v38 uint = 2
                                                    _ = v38
                                                default:
                                                    var v39 float32 = 0.0
                                                    _ = v39
                                                }
                                            } else {
                                                fmt.Println(v15)
                                            }
                                            if true {
                                                switch fi23 {
                                                case 42:
                                                    v15 -= 0.1
                                                default:
                                                    switch 42 {
                                                    case 174:
                                                        var v40 int64 = 42
                                                        _ = v40
                                                    case 127:
                                                        var v41 int16 = (v16 | v16)
                                                        _ = v41
                                                    default:
                                                        if len(v2) >= 2 {
                                                            v2 = v2[1:len(v2)]
                                                        }
                                                    }
                                                }
                                            }
                                            v42 := GenS2{F0: 63, F1: 204, F2: v17}
                                            v10 = (v10 % ((524) | 1))
                                            _ = v42
                                        } else {
                                            var v43 int16 = (v16 >> ((65535) & 63))
                                            if true {
                                                v44 := GenS2{F0: -10, F1: -370, F2: v17}
                                                var v45 int = fi23
                                                _ = v44
                                                _ = v45
                                            }
                                            _ = v43
                                        }
                                        if (true || false) {
                                            fmt.Println(v17)
                                            if true {
                                                if true {
                                                    v10 -= (v10 * v10)
                                                    v46 := GenS6{F0: -1, F1: fi22, F2: 42}
                                                    _ = v46
                                                } else {
                                                    switch (fi23 * (fi23 * 4294967295)) {
                                                    case 4294967295:
                                                        v15 += v15
                                                    case 232:
                                                        switch (fi23 >> ((780) & 63)) {
                                                        case 0:
                                                            v47 := GenS7{F0: -2147483648, F1: 4294967295}
                                                            _ = v47
                                                        case 250:
                                                            var v48 int64 = -572
                                                            _ = v48
                                                        default:
                                                            if false {
                                                                v49 := GenS8{F0: 951, F1: "", F2: (!(false || (-80 == 42)))}
                                                                _ = v49
                                                            }
                                                            if true {
                                                                var v50 uint8 = 29
                                                                fmt.Println(v16, fi23, v50)
                                                                v51 := GenS3{F0: 42, F1: v16, F2: (!(91 <= 0))}
                                                                _ = v50
                                                                _ = v51
                                                            }
                                                            delete(v4, ("\n\t" + "a"))
                                                        }
                                                    case -930:
                                                        v52 := GenS9{F0: 42, F1: v10, F2: (v10 - v10)}
                                                        _ = v52
                                                    default:
                                                        fmt.Println(fi22, fi23)
                                                    }
                                                    if (uint16(v10) <= 42) {
                                                        var v53 bool = false
                                                        _ = v53
                                                    }
                                                    fmt.Println(v10)
                                                    v16 -= -708
                                                }
                                            } else {
                                                v15 += (v15 - (v15 / v15))
                                                fmt.Println(v10, fi23, fi23)
                                                if false {
                                                    v54 := GenS10{F0: v10, F1: 2147483647, F2: (v16 + (^(v16)))}
                                                    fmt.Println(v16, v10, fi23)
                                                    v10 += (v10 & 4294967295)
                                                    _ = v54
                                                } else {
                                                    if len(v3) > 0 {
                                                        v3[0] = (v10 + v10)
                                                    }
                                                    if (107 >= -53) {
                                                        v10 = (v10 << ((4294967295) & 63))
                                                        fmt.Println(v17)
                                                    } else {
                                                        switch (fi23 & fi22) {
                                                        case 32767:
                                                            v15 = float64(float32(v15))
                                                        case -2147483648:
                                                            var v55 int16 = v16
                                                            _ = v55
                                                        default:
                                                            v56 := GenS1{F0: 0, F1: 0, F2: (v17 + v17)}
                                                            _ = v56
                                                        }
                                                        fmt.Println(v17, v15)
                                                    }
                                                    var v57 float32 = ((0.1 / 1.0) / 1e10)
                                                    var v58 string = ("\n\t" + v17)
                                                    _ = v57
                                                    _ = v58
                                                }
                                            }
                                            v17 = (v17 + "hello")
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                            v59 := v4["\n\t"]
                                            v60 := GenS8{F0: -9223372036854775808, F1: (v17 + "foo bar"), F2: (-2.25 < (((-2.25 * v15) - (v15 - (-2.25 - (((((((1.5 - 3.14159) + 3.14159) + (((float64(1.5) + v15) * float64((-2.25 / (0.1 - 1.5)))) - v15)) + (v15 + (((v15 / v15) * float64(1e10)) + float64((3.14159 - (float32(((v15 / (v15 + (1.5 / (((float64(((-2.25 / 3.14159) - -2.25)) / (((0.1 + (v15 * (v15 - (v15 + float64(0.1))))) * (v15 + v15)) / v15)) - v15) / (((-2.25 * v15) + v15) + v15))))) + v15)) - 3.14159)))))) / float64((0.1 * float32((((v15 / v15) - v15) / (v15 / v15)))))) / v15) * (float64(1e10) / v15))))) / 0.1))}
                                            _ = v59
                                            _ = v60
                                        } else {
                                            v10 -= 373
                                        }
                                        var v61 string = v17
                                        _ = v61
                                    }
                                    v62 := GenS7{F0: 0, F1: (v10 - 2)}
                                    _ = v62
                                } else {
                                    v2 = append(v2, 127)
                                }
                                switch fi23 {
                                case 32768:
                                    v63 := v4[v17]
                                    _ = v63
                                default:
                                    v16 += (v16 + 32767)
                                }
                            }
                        }
                    }
                case "foo bar":
                    var v64 int8 = -128
                    _ = v64
                default:
                    delete(v4, v17)
                }
            }
            _ = fi23
        }
        for fi65 := 0; fi65 < 12; fi65++ {
            fmt.Println(v10)
            if false {
                v66 := GenS7{F0: 0, F1: 2}
                v17 = v17
                var v67 string = "hello"
                switch v67 {
                case "hello":
                    var v68 uint = 553
                    v10 = (v10 & (v10 | (v10 + 2)))
                    switch "foo bar" {
                    case "":
                        var v69 int32 = -2147483648
                        _ = v69
                    case "foo bar":
                        if (true && ((!(!(!false))) && ((int32(v10) >= 42) || false))) {
                            var v70 int64 = 42
                            v71 := GenS6{F0: -43, F1: (fi65 | fi65), F2: 42}
                            v68 = v68
                            _ = v70
                            _ = v71
                        }
                    case "a":
                        v72 := GenS2{F0: (fi65 & (fi65 * fi65)), F1: 42, F2: v67}
                        _ = v72
                    default:
                        var v73 float32 = 1.5
                        _ = v73
                    }
                    if true {
                        v68 -= (v68 - (v68 / ((v68) | 1)))
                        fmt.Println(v15)
                        switch v67 {
                        case "":
                            v74 := GenS10{F0: v10, F1: int32(v10), F2: (v16 + (v16 * v16))}
                            _ = v74
                        case "hello":
                            switch v67 {
                            case "\n\t":
                                var v75 string = v17
                                _ = v75
                            case "a":
                                if (41 > 0) {
                                    v67 = "\n\t"
                                    switch int(v10) {
                                    case 1:
                                        var v76 string = v17
                                        _ = v76
                                    case -1:
                                        switch fi65 {
                                        case 128:
                                            var v77 uint64 = 1
                                            _ = v77
                                        case 2147483647:
                                            if (!false) {
                                                switch (fi22 + 1) {
                                                case 4294967295:
                                                    var v78 float64 = v15
                                                    var v79 int8 = 17
                                                    v4[""] = (v10 * v10)
                                                    v80 := GenS9{F0: v79, F1: v10, F2: (v10 | (v10 - (^(v10))))}
                                                    _ = v78
                                                    _ = v79
                                                    _ = v80
                                                }
                                                v81 := GenS1{F0: 759, F1: -128, F2: ""}
                                                v67 = (v17 + v17)
                                                _ = v81
                                            } else {
                                                if ("foo bar" >= "") {
                                                    var v82 float32 = (((1.5 + (float32((0.1 * (0.1 * (v15 * float64(1.5))))) / (3.14159 / (1.5 / (0.0 - (0.1 * (-2.25 + (0.0 / -2.25)))))))) + -2.25) / 1.5)
                                                    fmt.Println(v10, fi65)
                                                    v83 := GenS11{F0: 255, F1: 156}
                                                    _ = v82
                                                    _ = v83
                                                } else {
                                                    v15 = (0.1 - v15)
                                                }
                                                fmt.Println(v17, v67, v10)
                                            }
                                        case -1:
                                            v16 = (^(v16))
                                        }
                                    case 32768:
                                        delete(v4, (v67 + v17))
                                    }
                                    fmt.Println(v16, fi22)
                                } else {
                                    v84 := GenS12{F0: 0, F1: 64}
                                    _ = v84
                                }
                            }
                        case "foo bar":
                            if (!(0 < 127)) {
                                v15 += ((float64(1.5) + v15) / float64(float32((1.5 / (0.0 / v15)))))
                                v85 := GenS12{F0: 0, F1: 42}
                                _ = v85
                            } else {
                                if false {
                                    v15 = float64(0.0)
                                    fmt.Println(v17)
                                }
                                if true {
                                    fmt.Println(v17, v68, v68)
                                } else {
                                    fmt.Println(v10)
                                }
                            }
                        default:
                            fmt.Println(v10, v16)
                        }
                    } else {
                        if ((!true) && false) {
                            if true {
                                if len(v3) > 0 {
                                    v3[0] = (v10 ^ v10)
                                }
                                var v86 uint8 = uint8(v10)
                                _ = v86
                            } else {
                                v16 -= (v16 * v16)
                                var v87 int32 = -2147483648
                                v16 = (^(v16))
                                _ = v87
                            }
                            v68 = v68
                        } else {
                            var v88 float64 = v15
                            v15 += (-2.25 / 3.14159)
                            _ = v88
                        }
                        v89 := GenS13{F0: (168 <= 255)}
                        if ((v16 << (((v68 / ((v68) | 1))) & 63)) < -648) {
                            switch v67 {
                            case "foo bar":
                                if true {
                                    v10 += 2147483647
                                }
                            case "":
                                v90 := GenS12{F0: 127, F1: -1}
                                _ = v90
                            case "\n\t":
                                v91 := GenS3{F0: v16, F1: (v16 << (((^(v68))) & 63)), F2: ("\n\t" > ((v67 + v17) + v17))}
                                _ = v91
                            default:
                                if false {
                                    fmt.Println(v15)
                                }
                            }
                            v16 = v16
                        } else {
                            fmt.Println(v15, fi65, v10)
                        }
                        _ = v89
                    }
                    v92 := GenS14{F0: 848}
                    v68 -= v68
                    _ = v68
                    _ = v92
                case "":
                    delete(v4, v67)
                case "\n\t":
                    fmt.Println(fi65)
                }
                _ = v66
                _ = v67
            }
            var v93 int16 = (v16 >> ((0) & 63))
            fmt.Println(fi65, v16, v16)
            v94 := GenS15{F0: 863, F1: -964, F2: v16}
            _ = fi65
            _ = v93
            _ = v94
        }
        _ = fi22
    }
    var v95 uint64 = 4294967295
    switch "foo bar" {
    case "a":
        if true {
            v95 += v95
        } else {
            var v96 int = 127
            v97 := GenS4{F0: genFn0(255), F1: 128}
            _ = v96
            _ = v97
        }
    case "hello":
        fmt.Println(v17)
    case "\n\t":
        var v98 int8 = -128
        _ = v98
    default:
        delete(v4, (v17 + "hello"))
    }
    v99 := GenS12{F0: genFn0(42), F1: 0}
    var v100 float64 = (1e10 / -2.25)
    switch v17 {
    case "":
        var v101 float64 = -2.25
        _ = v101
    case "hello":
        var v102 int = 32768
        _ = v102
    case "\n\t":
        v103 := v4[v17]
        _ = v103
    }
    _ = len(v1)
    var v104 int64 = 42
    v10 += (v10 * v10)
    if (v104 > -1) {
        switch 4294967295 {
        case -339:
            var v105 int8 = 127
            _ = v105
        case -9223372036854775808:
            v106 := v4[v17]
            _ = v106
        }
        switch "" {
        case "hello":
            var v107 uint32 = v10
            _ = v107
        default:
            v95 -= 0
        }
        var v108 uint = 989
        _ = v108
    }
    v109 := GenS16{F0: (true || (v104 != (-(v104)))), F1: ((uint8(v10) < 128) && ((v17 > "\n\t") || true)), F2: "a"}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v100:", v100)
    fmt.Println("v104:", v104)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v95:", v95)
    _ = v1
    _ = v10
    _ = v100
    _ = v104
    _ = v109
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v95
    _ = v99
}
