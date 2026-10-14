package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool }
type GenS2 struct { F0 int32; F1 uint }
type GenS3 struct { F0 int32; F1 int }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 int; F1 int8; F2 int64 }
type GenS8 struct { F0 uint32; F1 uint64 }
func genFn0() int64 {
    for fi7 := 0; fi7 < 17; fi7++ {
        var v8 float64 = 0.1
        v9 := []int{fi7, fi7}
        _ = fi7
        _ = v8
        _ = v9
    }
    var v10 int64 = 42
    _ = v10
    return 1
}
func genFn1(p0 int32, p1 int16, p2 int64) uint64 {
    if false {
        v11 := GenS1{F0: 527, F1: false}
        _ = v11
    } else {
        fmt.Println(p1)
        var v12 uint = 2
        _ = v12
    }
    v13 := GenS1{F0: 1, F1: (32767 >= 2)}
    p1 -= 42
    var v14 string = "foo bar"
    p0 = -575
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v14
    return 1
}

func main() {
    var v1 int16 = 1
    var v2 int = 32768
    v3 := []uint32{2, 2147483647}
    var v4 string = "foo bar"
    var v5 uint8 = 1
    var v6 bool = ((false && (1 != 1)) && false)
    for fi15 := 0; fi15 < 18; fi15++ {
        for fi16 := 0; fi16 < 2; fi16++ {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            switch (fi16 | -670) {
            case 32767:
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                fmt.Println(fi15, v6)
                if v6 {
                    v17 := GenS1{F0: 2, F1: (242 > 4294967295)}
                    v1 -= (v1 - int16(fi16))
                    v5 += v5
                    _ = v17
                } else {
                    if v6 {
                        if (true || true) {
                            switch v4 {
                            case "foo bar":
                                switch v4 {
                                case "a":
                                    fmt.Println(v1, v5)
                                default:
                                    v5 = (v5 >> ((18446744073709551615) & 63))
                                }
                            case "":
                                if false {
                                    v2 -= (fi15 << ((65535) & 63))
                                    fmt.Println(v6)
                                }
                            case "a":
                                if len(v3) > 0 {
                                    v3[0] = 0
                                }
                            }
                            fmt.Println(fi16, v2)
                            v5 -= v5
                            v3 = append(v3, 652)
                        } else {
                            v6 = v6
                            v2 -= 2147483647
                            var v18 uint64 = 0
                            _ = v18
                        }
                        var v19 uint32 = 441
                        v20 := GenS2{F0: 42, F1: 1}
                        var v21 float64 = (1e10 / 1.0)
                        if false {
                            v1 += 32767
                            if (("" + (v4 + (((("" + "") + v4) + "a") + v4))) >= "\n\t") {
                                fmt.Println(v21)
                                fmt.Println(v1, v2, v21)
                                v21 = float64(((0.1 - 0.0) * float32(float64(1e10))))
                            } else {
                                if v6 {
                                    if (genFn0() < -9223372036854775808) {
                                        if true {
                                            var v22 uint16 = 1000
                                            if v6 {
                                                v23 := GenS2{F0: -229, F1: 4294967295}
                                                if false {
                                                    fmt.Println(v6, v5, v19)
                                                    if (!(42 > 308)) {
                                                        fmt.Println(fi16)
                                                    }
                                                    switch (^(v2)) {
                                                    case 1:
                                                        v1 = (-(v1))
                                                    case -9223372036854775808:
                                                        var v24 uint64 = 0
                                                        _ = v24
                                                    default:
                                                        var v25 uint8 = (v5 + v5)
                                                        _ = v25
                                                    }
                                                }
                                                _ = v23
                                            } else {
                                                v26 := GenS1{F0: 349, F1: (981 > 155)}
                                                v21 = float64(1.5)
                                                _ = v26
                                            }
                                            _ = v22
                                        } else {
                                            v5 += (v5 * v5)
                                            v6 = (v21 > float64((1e10 - (1.5 + 0.0))))
                                            if v6 {
                                                v4 = ""
                                                v19 = (^(v19))
                                                if len(v3) > 0 {
                                                    v3[0] = (v19 % (((v19 << ((1) & 63))) | 1))
                                                }
                                                _ = len(v3)
                                                var v27 int64 = genFn0()
                                                _ = v27
                                            } else {
                                                if v6 {
                                                    if len(v3) > 0 {
                                                        v3[0] = v19
                                                    }
                                                }
                                            }
                                        }
                                        _ = len(v3)
                                    } else {
                                        if (false && (v6 || v6)) {
                                            switch v4 {
                                            case "hello":
                                                if (v6 && v6) {
                                                    var v28 uint16 = 459
                                                    switch (("" + v4) + v4) {
                                                    case "foo bar":
                                                        v28 -= (v28 << ((18446744073709551615) & 63))
                                                    case "a":
                                                        v29 := GenS1{F0: 181, F1: (!v6)}
                                                        _ = v29
                                                    case "":
                                                        v3 = append(v3, (v19 ^ (^(v19))))
                                                    }
                                                    v1 = v1
                                                    v3 = append(v3, 2147483647)
                                                    _ = v28
                                                }
                                            }
                                            if (645 > 897) {
                                                var v30 int32 = -2147483648
                                                v3 = append(v3, v19)
                                                _ = v30
                                            } else {
                                                var v31 float64 = 0.1
                                                fmt.Println(v19, v21)
                                                var v32 int = (-(fi16))
                                                _ = v31
                                                _ = v32
                                            }
                                            v4 = (((v4 + ("a" + (v4 + v4))) + (("" + v4) + ((("hello" + "") + v4) + (v4 + v4)))) + v4)
                                        } else {
                                            if (v4 == "") {
                                                var v33 bool = ((v1 - (-(v1))) < v1)
                                                v4 = "a"
                                                _ = len(v3)
                                                _ = v33
                                            }
                                            v4 = "hello"
                                        }
                                        var v34 int = (fi16 ^ (fi15 | fi15))
                                        if (-2147483648 != -1) {
                                            var v35 uint32 = v19
                                            _ = v35
                                        } else {
                                            fmt.Println(v1)
                                            var v36 int = (v2 + int(v5))
                                            v37 := GenS2{F0: -42, F1: 42}
                                            _ = v36
                                            _ = v37
                                        }
                                        _ = v34
                                    }
                                    switch ("" + (("hello" + v4) + ("\n\t" + ("\n\t" + ("hello" + "hello"))))) {
                                    case "foo bar":
                                        if ((10 >= 100) && (!v6)) {
                                            v3 = append(v3, v19)
                                        } else {
                                            fmt.Println(v21, fi16)
                                        }
                                    case "\n\t":
                                        var v38 float32 = (-2.25 * 0.0)
                                        _ = v38
                                    case "":
                                        var v39 int16 = -32768
                                        _ = v39
                                    }
                                    v4 = "hello"
                                    var v40 float64 = (v21 / -2.25)
                                    if (!(!(((v5 << ((uint(v5)) & 63)) > 254) && false))) {
                                        if v6 {
                                            fmt.Println(fi15, v21, fi16)
                                        }
                                        var v41 float32 = ((-2.25 - 0.1) - 1.5)
                                        _ = v41
                                    } else {
                                        fmt.Println(v6)
                                        if true {
                                            v42 := GenS1{F0: genFn1(-733, v1, 155), F1: false}
                                            _ = v42
                                        }
                                    }
                                    _ = v40
                                } else {
                                    switch v4 {
                                    case "foo bar":
                                        if (v6 || (v6 || ((!v6) || (!true)))) {
                                            v6 = (!(!v6))
                                            if (1 != v1) {
                                                var v43 int64 = 0
                                                _ = v43
                                            }
                                            v4 = v4
                                        }
                                    case "":
                                        fmt.Println(v1)
                                    default:
                                        fmt.Println(v19, v1, v2)
                                    }
                                    var v44 int16 = v1
                                    fmt.Println(v2, v4, v2)
                                    switch (fi15 % (((fi15 * (fi16 + fi15))) | 1)) {
                                    case -2147483648:
                                        if len(v3) > 0 {
                                            v3[0] = v19
                                        }
                                    default:
                                        fmt.Println(fi15, v4, fi16)
                                    }
                                    _ = v44
                                }
                                v45 := GenS1{F0: 4294967295, F1: true}
                                var v46 uint16 = 176
                                _ = v45
                                _ = v46
                            }
                            v47 := GenS1{F0: 352, F1: v6}
                            _ = v47
                        }
                        v48 := GenS3{F0: 2147483647, F1: fi16}
                        v19 += (v19 / (((v19 - v19)) | 1))
                        _ = v19
                        _ = v20
                        _ = v21
                        _ = v48
                    } else {
                        if len(v3) > 0 {
                            v3[0] = 2147483648
                        }
                        v49 := GenS1{F0: 475, F1: (-2.25 > 0.1)}
                        v1 = -491
                        _ = v49
                    }
                    switch (fi16 | (fi15 - v2)) {
                    case 42:
                        if v6 {
                            v50 := GenS3{F0: -594, F1: -624}
                            switch (-(v2)) {
                            case 330:
                                v3 = append(v3, uint32(v5))
                                var v51 int32 = int32(v1)
                                if false {
                                    v52 := GenS3{F0: 2147483647, F1: (^(v2))}
                                    _ = v52
                                }
                                _ = v51
                            }
                            if false {
                                v4 = ""
                                var v53 int = fi15
                                var v54 uint16 = uint16(v5)
                                _ = v53
                                _ = v54
                            }
                            _ = v50
                        }
                    case 4294967295:
                        fmt.Println(fi16, v5)
                    default:
                        var v55 int8 = -1
                        _ = v55
                    }
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                }
                if len(v3) > 0 {
                    v3[0] = 0
                }
                if false {
                    _ = len(v3)
                    switch 4294967295 {
                    case 1:
                        v5 -= (v5 % (((^(v5))) | 1))
                        var v56 int16 = v1
                        _ = v56
                    case 32768:
                        if v6 {
                            var v57 uint16 = 2
                            v4 = ((((((v4 + (v4 + "")) + v4) + (v4 + ("\n\t" + ((v4 + "foo bar") + "a")))) + v4) + v4) + ("foo bar" + v4))
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            _ = v57
                        } else {
                            switch (fi16 * fi15) {
                            case 918:
                                v58 := GenS4{F0: 1}
                                _ = v58
                            case 4294967295:
                                v5 -= 217
                                if (-2.25 <= 1e10) {
                                    v2 = (fi15 & 1)
                                    v1 = (-(v1))
                                    if true {
                                        var v59 uint = 65535
                                        switch -9223372036854775808 {
                                        case 32768:
                                            v2 = (fi15 / ((fi15) | 1))
                                        }
                                        _ = v59
                                    }
                                } else {
                                    fmt.Println(v4)
                                }
                            default:
                                if false {
                                    fmt.Println(v2, v2, v2)
                                    var v60 uint32 = 4294967295
                                    v61 := GenS5{F0: 42}
                                    _ = len(v3)
                                    var v62 int8 = -6
                                    _ = v60
                                    _ = v61
                                    _ = v62
                                } else {
                                    v4 = v4
                                    switch v4 {
                                    case "a":
                                        var v63 string = ""
                                        _ = v63
                                    case "hello":
                                        v4 = "foo bar"
                                    default:
                                        v4 = "\n\t"
                                    }
                                }
                            }
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            switch v4 {
                            case "a":
                                var v64 uint64 = 0
                                _ = v64
                            case "":
                                v65 := GenS6{F0: (v4 >= v4)}
                                _ = v65
                            }
                        }
                    case -2147483648:
                        v66 := GenS4{F0: 406}
                        _ = v66
                    default:
                        fmt.Println(fi16, v2, v4)
                    }
                    if false {
                        var v67 uint = 771
                        var v68 uint32 = 2
                        v5 = (v5 ^ 1)
                        _ = v67
                        _ = v68
                    }
                }
            case 42:
                switch v4 {
                case "a":
                    v69 := GenS7{F0: (fi16 | (v2 ^ v2)), F1: 0, F2: -575}
                    _ = v69
                case "":
                    switch fi15 {
                    case 128:
                        var v70 uint32 = 1
                        _ = v70
                    case -9223372036854775808:
                        var v71 string = "\n\t"
                        _ = v71
                    default:
                        fmt.Println(fi16)
                    }
                }
            case -9223372036854775808:
                if ((2 <= 222) && false) {
                    switch -2147483648 {
                    case 285:
                        if (315 < 144) {
                            v72 := GenS7{F0: (-(v2)), F1: 100, F2: 9223372036854775807}
                            _ = v72
                        }
                    case -1:
                        var v73 int = (fi16 & (^(fi16)))
                        _ = v73
                    case 9223372036854775807:
                        switch (^(fi16)) {
                        case 32767:
                            var v74 bool = ((867 <= 2147483647) && v6)
                            _ = v74
                        case -9223372036854775808:
                            switch v4 {
                            case "foo bar":
                                var v75 uint = 428
                                _ = v75
                            case "\n\t":
                                v4 = v4
                            default:
                                if (v2 > (v2 & (v2 << ((428) & 63)))) {
                                    v5 += v5
                                    _ = len(v3)
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                } else {
                                    fmt.Println(v6)
                                    v5 = 131
                                    v5 = v5
                                }
                            }
                        default:
                            if v6 {
                                var v76 uint64 = 477
                                if (!v6) {
                                    v6 = ((float64((3.14159 - 1e10)) - 0.1) >= (1.5 / 1.5))
                                }
                                v77 := GenS2{F0: -2147483648, F1: uint(v5)}
                                v4 = v4
                                fmt.Println(v76)
                                _ = v76
                                _ = v77
                            } else {
                                v78 := GenS1{F0: 1, F1: (int8(fi15) != 127)}
                                if ((v6 || false) || false) {
                                    if len(v3) > 0 {
                                        v3[0] = 2147483648
                                    }
                                } else {
                                    var v79 uint32 = 2147483647
                                    var v80 string = (((v4 + (v4 + v4)) + v4) + (((v4 + v4) + "a") + "a"))
                                    if v6 {
                                        var v81 float32 = 0.0
                                        _ = v81
                                    }
                                    _ = v79
                                    _ = v80
                                }
                                switch fi15 {
                                case 32768:
                                    if v6 {
                                        var v82 uint16 = 0
                                        v83 := GenS3{F0: 1, F1: (^(fi16))}
                                        _ = v82
                                        _ = v83
                                    }
                                }
                                _ = v78
                            }
                        }
                    default:
                        v5 = v5
                    }
                }
                fmt.Println(v2, fi15)
            default:
                if v6 {
                    var v84 int = fi15
                    v5 -= (v5 / ((255) | 1))
                    v85 := GenS8{F0: 2147483648, F1: 558}
                    _ = v84
                    _ = v85
                } else {
                    var v86 uint16 = 2
                    v2 -= 4294967295
                    var v87 float32 = 1.5
                    _ = v86
                    _ = v87
                }
            }
            _ = fi16
        }
        fmt.Println(fi15, v4, fi15)
        v2 -= (v2 * 42)
        _ = fi15
    }
    switch v4 {
    case "hello":
        for fi88 := 0; fi88 < 17; fi88++ {
            v6 = true
            for fi89 := 0; fi89 < 5; fi89++ {
                v1 = v1
                var v90 float32 = -2.25
                _ = fi89
                _ = v90
            }
            _ = fi88
        }
    case "":
        var v91 int = v2
        _ = v91
    default:
        for fi92 := 0; fi92 < 9; fi92++ {
            v93 := GenS6{F0: (v1 < v1)}
            _ = fi92
            _ = v93
        }
    }
    var v94 int16 = v1
    if len(v3) > 0 {
        v3[0] = 2147483647
    }
    v94 -= v94
    switch v4 {
    case "hello":
        var v95 float32 = 3.14159
        _ = v95
    default:
        v5 += v5
    }
    v5 = (v5 >> ((0) & 63))
    if (885 <= 4294967295) {
        var v96 int = (v2 * (v2 / (((v2 ^ v2)) | 1)))
        _ = v96
    }
    v6 = v6
    var v97 float32 = 1.5
    var v98 int32 = 42
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v99 uint16 = 65535
    switch (-(v2)) {
    case 128:
        _ = len(v3)
    case -585:
        var v100 float32 = v97
        _ = v100
    default:
        var v101 int16 = (v1 >> ((1) & 63))
        _ = v101
    }
    for fi102 := 0; fi102 < 19; fi102++ {
        fmt.Println(v98, v97)
        if (!(false || (v94 < v1))) {
            var v103 uint8 = 1
            if v6 {
                v104 := GenS5{F0: 255}
                v94 = (^(v94))
                _ = v104
            } else {
                fmt.Println(v98, v99)
            }
            v103 = 203
            _ = v103
        } else {
            if ((v94 >> ((1) & 63)) > 1000) {
                if (42 != 618) {
                    if v6 {
                        var v105 uint64 = 2
                        _ = v105
                    } else {
                        if (((!(18446744073709551615 < 0)) || v6) && (v6 || v6)) {
                            v5 -= 42
                            var v106 string = ((v4 + ("\n\t" + v4)) + v4)
                            for fi107 := 0; fi107 < 19; fi107++ {
                                var v108 uint = 2
                                _ = len(v3)
                                _ = fi107
                                _ = v108
                            }
                            _ = v106
                        }
                        fmt.Println(v98, v97, v98)
                    }
                } else {
                    fmt.Println(v94)
                    _ = len(v3)
                }
                var v109 int8 = 100
                _ = v109
            } else {
                for fi110 := 0; fi110 < 8; fi110++ {
                    if v6 {
                        if (602 < 7) {
                            switch 32767 {
                            case 32768:
                                var v111 uint64 = genFn1((-(v98)), 1000, genFn0())
                                _ = v111
                            }
                            var v112 int8 = 127
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            _ = v112
                        } else {
                            _ = len(v3)
                        }
                        v1 -= (v1 & v1)
                    }
                    _ = fi110
                }
                fmt.Println(v6, v98)
            }
        }
        _ = fi102
    }
    v113 := GenS3{F0: int32(v94), F1: v2}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v94:", v94)
    fmt.Println("v97:", v97)
    fmt.Println("v98:", v98)
    fmt.Println("v99:", v99)
    _ = v1
    _ = v113
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v94
    _ = v97
    _ = v98
    _ = v99
}
