package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16 }
type GenS2 struct { F0 uint64; F1 bool }
type GenS3 struct { F0 bool; F1 int8 }
type GenS4 struct { F0 uint8; F1 int16 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 bool; F1 uint64; F2 bool }
type GenS8 struct { F0 string }
type GenS9 struct { F0 uint64 }
func genFn0(p0 uint8, p1 uint16) int16 {
    for fi6 := 0; fi6 < 2; fi6++ {
        fmt.Println(p0, fi6, p1)
        v7 := []int{fi6}
        v8 := GenS1{F0: p0, F1: 32767}
        _ = fi6
        _ = v7
        _ = v8
    }
    fmt.Println(p0, p0)
    var v9 string = ("hello" + ((("hello" + "foo bar") + "foo bar") + (("hello" + "hello") + (("foo bar" + "a") + "a"))))
    _ = p0
    _ = p1
    _ = v9
    return 1
}
func genFn1() uint64 {
    var v10 uint = 4294967295
    switch "hello" {
    case "a":
        var v11 int32 = 2147483647
        _ = v11
    }
    var v12 uint = 255
    for fi13 := 0; fi13 < 20; fi13++ {
        if false {
            v10 = v10
            if true {
                fmt.Println(fi13, v12)
            }
            v12 += 812
        } else {
            var v14 int64 = -9223372036854775808
            _ = v14
        }
        _ = fi13
    }
    _ = v10
    _ = v12
    return 4
}

func main() {
    v1 := make(map[string]int)
    var v2 int16 = -348
    v3 := []uint32{0}
    v4 := make(map[int]bool)
    v5 := make(map[uint32]int)
    for fi15 := 0; fi15 < 12; fi15++ {
        for fi16 := 0; fi16 < 4; fi16++ {
            v2 = genFn0(0, 0)
            _ = fi16
        }
        switch "\n\t" {
        case "hello":
            v17 := GenS2{F0: 18446744073709551615, F1: (true && (586 >= 1))}
            _ = v17
        }
        _ = fi15
    }
    switch 471 {
    case -1:
        var v18 uint64 = 845
        _ = v18
    case 357:
        var v19 uint16 = 1
        _ = v19
    case 32767:
        if false {
            v3 = append(v3, 2)
        } else {
            for fi20 := 0; fi20 < 2; fi20++ {
                var v21 string = "\n\t"
                v22 := GenS1{F0: 1, F1: -1}
                if (!(((fi20 != fi20) || true) && false)) {
                    var v23 int64 = 42
                    _ = v23
                } else {
                    var v24 int8 = -128
                    _ = len(v3)
                    if true {
                        v2 = 42
                        if ((true && true) && (v2 != (v2 - (v2 / ((0) | 1))))) {
                            for fi25 := 0; fi25 < 3; fi25++ {
                                if true {
                                    fmt.Println(fi25)
                                } else {
                                    v24 = -128
                                    switch (fi25 | fi20) {
                                    case -279:
                                        var v26 uint64 = 658
                                        _ = v26
                                    case -2147483648:
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                    case -546:
                                        v27 := GenS3{F0: (493 > 2147483648), F1: v24}
                                        _ = v27
                                    default:
                                        var v28 int8 = 42
                                        _ = v28
                                    }
                                    v2 -= v2
                                }
                                var v29 float64 = (0.0 - 0.1)
                                var v30 int64 = 977
                                var v31 int = (^(fi25))
                                if (genFn0(1, 1) != (v2 | -28)) {
                                    v32 := GenS4{F0: 255, F1: v2}
                                    v24 += (v24 + -1)
                                    _ = v32
                                }
                                switch ((v21 + v21) + v21) {
                                case "hello":
                                    var v33 uint32 = 2147483648
                                    _ = v33
                                case "foo bar":
                                    fmt.Println(fi20, v21)
                                }
                                v31 = fi25
                                _ = fi25
                                _ = v29
                                _ = v30
                                _ = v31
                            }
                            if (true && false) {
                                var v34 uint16 = 261
                                _ = v34
                            }
                        } else {
                            var v35 uint32 = 478
                            var v36 int = 32768
                            for fi37 := 0; fi37 < 4; fi37++ {
                                if (!true) {
                                    v35 = (v35 << ((4294967295) & 63))
                                    var v38 float32 = 3.14159
                                    var v39 int = fi37
                                    v40 := GenS5{F0: (-(v36))}
                                    v35 = v35
                                    _ = v38
                                    _ = v39
                                    _ = v40
                                } else {
                                    var v41 int64 = -1
                                    _ = v41
                                }
                                var v42 string = ("a" + (v21 + ((v21 + "foo bar") + (v21 + v21))))
                                _ = fi37
                                _ = v42
                            }
                            _ = v35
                            _ = v36
                        }
                    } else {
                        for fi43 := 0; fi43 < 2; fi43++ {
                            if (true || true) {
                                v2 = (v2 << ((2) & 63))
                            } else {
                                var v44 int = (fi20 - (fi20 & fi43))
                                v24 -= (-(v24))
                                if ((v2 >> ((171) & 63)) > int16(v24)) {
                                    fmt.Println(fi20)
                                    v45 := GenS6{F0: 65535}
                                    if (!false) {
                                        v46 := GenS1{F0: 80, F1: v2}
                                        _ = v46
                                    } else {
                                        if false {
                                            v44 -= 128
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                            var v47 int64 = 0
                                            v47 -= -751
                                            if (!(0 <= 606)) {
                                                var v48 uint64 = 2
                                                v3 = append(v3, 2)
                                                var v49 int8 = v24
                                                fmt.Println(v21, v21)
                                                switch (fi20 | (-(fi43))) {
                                                case -9223372036854775808:
                                                    v21 = v21
                                                    if (v48 == v48) {
                                                        var v50 int = (-(fi20))
                                                        _ = v50
                                                    } else {
                                                        v48 = v48
                                                        if false {
                                                            if true {
                                                                v51 := GenS6{F0: 381}
                                                                v24 = (v24 << ((255) & 63))
                                                                _ = v51
                                                            } else {
                                                                if (42 == 2147483648) {
                                                                    fmt.Println(v2, v24)
                                                                    var v52 uint64 = (v48 ^ 1)
                                                                    fmt.Println(v44)
                                                                    _ = v52
                                                                }
                                                            }
                                                            v53 := GenS3{F0: true, F1: v24}
                                                            _ = v53
                                                        } else {
                                                            _ = len(v3)
                                                        }
                                                        var v54 int32 = 0
                                                        _ = v54
                                                    }
                                                    v55 := GenS5{F0: v44}
                                                    _ = v55
                                                case 9223372036854775807:
                                                    var v56 string = (v21 + v21)
                                                    _ = v56
                                                case 127:
                                                    var v57 uint16 = 180
                                                    _ = v57
                                                default:
                                                    v21 = v21
                                                }
                                                v24 -= v24
                                                _ = v48
                                                _ = v49
                                            } else {
                                                switch v44 {
                                                case 1:
                                                    v44 += -914
                                                    delete(v5, 995)
                                                    var v58 uint64 = 2
                                                    switch fi43 {
                                                    case 493:
                                                        if len(v3) >= 2 {
                                                            v3 = v3[1:len(v3)]
                                                        }
                                                    case 42:
                                                        if true {
                                                            v59 := GenS6{F0: 24}
                                                            _ = v59
                                                        }
                                                    case 1:
                                                        _ = len(v5)
                                                    }
                                                    if false {
                                                        _ = len(v5)
                                                        v60 := v1[(v21 + ("foo bar" + v21))]
                                                        var v61 float32 = 0.1
                                                        _ = v60
                                                        _ = v61
                                                    } else {
                                                        if (53 <= 2) {
                                                            v2 = (v2 << ((669) & 63))
                                                        } else {
                                                            v2 = v2
                                                            var v62 uint32 = 364
                                                            v24 += 28
                                                            fmt.Println(fi43)
                                                            v58 -= v58
                                                            _ = v62
                                                        }
                                                        var v63 int = -1
                                                        var v64 float64 = (((-2.25 / -2.25) + 1e10) * (1.5 * 1e10))
                                                        _ = v63
                                                        _ = v64
                                                    }
                                                    var v65 bool = false
                                                    _ = v58
                                                    _ = v65
                                                case 79:
                                                    if false {
                                                        v44 = (fi20 / ((4294967295) | 1))
                                                        var v66 int = fi20
                                                        fmt.Println(v2, v21, v24)
                                                        _ = v66
                                                    }
                                                default:
                                                    if false {
                                                        var v67 uint8 = 42
                                                        if (!("\n\t" <= "foo bar")) {
                                                            v67 += v67
                                                            v68 := GenS7{F0: true, F1: 2, F2: (true && true)}
                                                            switch (("foo bar" + v21) + (v21 + "\n\t")) {
                                                            case "":
                                                                v69 := GenS2{F0: 597, F1: false}
                                                                _ = v69
                                                            }
                                                            _ = v68
                                                        } else {
                                                            if true {
                                                                v1["foo bar"] = (v44 | (fi20 / (((v44 & 127)) | 1)))
                                                            }
                                                            v70 := GenS8{F0: v21}
                                                            var v71 uint16 = 42
                                                            _ = v70
                                                            _ = v71
                                                        }
                                                        if true {
                                                            v47 = (v47 | 9223372036854775807)
                                                        }
                                                        v21 = v21
                                                        v21 = "foo bar"
                                                        _ = v67
                                                    } else {
                                                        var v72 int = (fi20 >> ((4294967295) & 63))
                                                        var v73 int32 = 668
                                                        _ = v72
                                                        _ = v73
                                                    }
                                                }
                                                fmt.Println(v21, v21)
                                            }
                                            _ = v47
                                        }
                                    }
                                    _ = v45
                                }
                                switch v21 {
                                case "\n\t":
                                    if true {
                                        switch (v21 + "\n\t") {
                                        case "foo bar":
                                            var v74 uint8 = 1
                                            _ = v74
                                        case "hello":
                                            switch 9223372036854775807 {
                                            case 32767:
                                                if false {
                                                    var v75 uint16 = 0
                                                    v76 := GenS1{F0: 137, F1: 1}
                                                    _ = v75
                                                    _ = v76
                                                }
                                            case 42:
                                                fmt.Println(v21, v24)
                                            }
                                        }
                                    }
                                    var v77 int64 = 764
                                    _ = v77
                                case "":
                                    var v78 bool = (709 > 42)
                                    _ = v78
                                case "a":
                                    if (false || (true && ((170 == 0) && false))) {
                                        fmt.Println(v2, v44)
                                        if (-1 < -1) {
                                            switch fi43 {
                                            case -1:
                                                fmt.Println(fi43, v21)
                                            default:
                                                v24 = (v24 >> ((18446744073709551615) & 63))
                                            }
                                            var v79 string = ((v21 + "a") + v21)
                                            _ = v79
                                        } else {
                                            if len(v3) > 0 {
                                                v3[0] = 42
                                            }
                                        }
                                        v80 := GenS3{F0: (2147483648 <= 2147483648), F1: v24}
                                        _ = v80
                                    } else {
                                        _ = len(v3)
                                        v81 := GenS6{F0: 32767}
                                        _ = v81
                                    }
                                default:
                                    v24 += 0
                                    if len(v3) > 0 {
                                        v3[0] = 2147483647
                                    }
                                    var v82 int64 = 854
                                    v21 = ("" + ((("foo bar" + v21) + v21) + v21))
                                    _ = v82
                                }
                                _ = v44
                            }
                            if true {
                                v4[fi43] = (4294967295 <= genFn1())
                                var v83 uint = 42
                                _ = v83
                            } else {
                                var v84 int = (fi43 / ((fi43) | 1))
                                v24 += (v24 | v24)
                                if len(v3) > 0 {
                                    v3[0] = 1
                                }
                                _ = v84
                            }
                            v21 = v21
                            _ = fi43
                        }
                        var v85 uint8 = 1
                        var v86 int = -25
                        _ = v85
                        _ = v86
                    }
                    _ = v24
                }
                _ = fi20
                _ = v21
                _ = v22
            }
            fmt.Println(v2)
            for fi87 := 0; fi87 < 14; fi87++ {
                if (0 > fi87) {
                    v3 = append(v3, 65535)
                }
                var v88 string = ("foo bar" + ("hello" + "hello"))
                _ = fi87
                _ = v88
            }
        }
    }
    fmt.Println(v2)
    var v89 bool = false
    v90 := v4[127]
    for fi91 := 0; fi91 < 16; fi91++ {
        fmt.Println(v2, v90, v90)
        _ = fi91
    }
    switch 4294967295 {
    case -182:
        v2 -= 1
    case 9223372036854775807:
        v90 = v90
    default:
        for fi92 := 0; fi92 < 16; fi92++ {
            var v93 string = ("a" + "a")
            _ = fi92
            _ = v93
        }
    }
    v94 := GenS9{F0: 637}
    switch ("" + (("a" + "") + "\n\t")) {
    case "hello":
        for fi95 := 0; fi95 < 16; fi95++ {
            var v96 int8 = 0
            _ = fi95
            _ = v96
        }
    default:
        v90 = v89
    }
    var v97 int16 = genFn0(176, 2)
    v89 = (!v89)
    var v98 int8 = 57
    fmt.Println(v89)
    v89 = false
    var v99 float64 = 3.14159
    v90 = ((v98 >> ((135) & 63)) >= -101)
    fmt.Println(v90)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v89:", v89)
    fmt.Println("v90:", v90)
    fmt.Println("v97:", v97)
    fmt.Println("v98:", v98)
    fmt.Println("v99:", v99)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v89
    _ = v90
    _ = v94
    _ = v97
    _ = v98
    _ = v99
}
