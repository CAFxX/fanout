package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64 }
type GenS2 struct { F0 int; F1 uint8 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int64; F1 int; F2 int32 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint64; F1 bool }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 int8; F1 int32 }
type GenS9 struct { F0 int64 }
type GenS10 struct { F0 uint; F1 int32 }
type GenS11 struct { F0 string; F1 int32 }
type GenS12 struct { F0 int; F1 uint32 }
func genFn0() uint {
    for fi5 := 0; fi5 < 10; fi5++ {
        for fi6 := 0; fi6 < 20; fi6++ {
            v7 := GenS1{F0: "\n\t", F1: 0}
            v8 := GenS1{F0: (("foo bar" + "foo bar") + ""), F1: 9223372036854775807}
            _ = fi6
            _ = v7
            _ = v8
        }
        v9 := GenS2{F0: -309, F1: 105}
        _ = fi5
        _ = v9
    }
    var v10 int32 = 2147483647
    if (-1 == -302) {
        for fi11 := 0; fi11 < 7; fi11++ {
            var v12 uint = 508
            _ = fi11
            _ = v12
        }
    } else {
        v13 := GenS1{F0: "\n\t", F1: int64(v10)}
        var v14 int = 32767
        v14 += (v14 * (v14 << ((2) & 63)))
        _ = v13
        _ = v14
    }
    var v15 int8 = 1
    _ = v10
    _ = v15
    return 65535
}
func genFn1(p0 int64, p1 bool, p2 uint) uint32 {
    if p1 {
        fmt.Println(p1)
        p2 = p2
        if p1 {
            v16 := []int{32767, -727, 32767, -707}
            _ = v16
        }
    } else {
        fmt.Println(p1)
    }
    v17 := GenS3{F0: 100}
    if p1 {
        if true {
            if ((4294967295 == 4294967295) && (!true)) {
                for fi18 := 0; fi18 < 9; fi18++ {
                    var v19 int = (fi18 % ((fi18) | 1))
                    switch 127 {
                    case 32768:
                        var v20 uint16 = 255
                        _ = v20
                    case -783:
                        p2 = p2
                    }
                    switch fi18 {
                    case 121:
                        p1 = (true && (v19 == v19))
                    default:
                        v21 := []int{(fi18 ^ (fi18 % (((-(fi18))) | 1))), 4294967295, (fi18 >> (((p2 + p2)) & 63))}
                        _ = v21
                    }
                    _ = fi18
                    _ = v19
                }
            }
            p0 -= p0
        }
        p2 = (p2 / (((p2 % (((p2 >> ((42) & 63))) | 1))) | 1))
    } else {
        if p1 {
            fmt.Println(p2, p0, p0)
            if (p1 || (0.1 > (0.0 - 3.14159))) {
                v22 := GenS4{F0: (p0 + 0), F1: int(p2), F2: 2147483647}
                _ = v22
            }
        } else {
            v23 := GenS5{F0: 1}
            p1 = p1
            _ = v23
        }
        p2 += (p2 - (^(p2)))
    }
    v24 := GenS2{F0: 127, F1: 255}
    var v25 uint16 = 65535
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    _ = v24
    _ = v25
    return 0
}

func main() {
    var v1 float64 = float64(-2.25)
    v2 := []int{32768, 632, 889, 1}
    var v3 bool = true
    var v4 uint16 = 2
    fmt.Println(v1, v3)
    switch "a" {
    case "\n\t":
        var v26 uint16 = (v4 * 2)
        _ = v26
    case "a":
        v3 = (0 > -881)
    case "":
        for fi27 := 0; fi27 < 17; fi27++ {
            for fi28 := 0; fi28 < 3; fi28++ {
                v4 -= (v4 >> ((714) & 63))
                _ = fi28
            }
            if len(v2) > 0 {
                v2[0] = (fi27 % ((42) | 1))
            }
            v4 = (^(v4))
            _ = fi27
        }
    }
    var v29 int16 = 1000
    var v30 int64 = 9223372036854775807
    for fi31 := 0; fi31 < 5; fi31++ {
        if (!(96 >= -2147483648)) {
            var v32 uint8 = 42
            _ = v32
        }
        v33 := GenS6{F0: 1, F1: (((!v3) && v3) || true)}
        if v3 {
            var v34 float64 = (0.0 / v1)
            _ = v34
        } else {
            switch "hello" {
            case "foo bar":
                var v35 uint32 = 242
                _ = v35
            case "":
                var v36 uint64 = 18446744073709551615
                _ = v36
            default:
                fmt.Println(fi31, v1, v4)
            }
            for fi37 := 0; fi37 < 14; fi37++ {
                v38 := GenS7{F0: 4294967295}
                if (v4 < v4) {
                    v30 = -12
                }
                _ = fi37
                _ = v38
            }
            switch (fi31 % (((fi31 * (fi31 - fi31))) | 1)) {
            case 42:
                var v39 int64 = v30
                _ = v39
            case 884:
                for fi40 := 0; fi40 < 10; fi40++ {
                    if v3 {
                        var v41 int16 = (v29 & (v29 * -1))
                        _ = v41
                    }
                    _ = fi40
                }
            default:
                fmt.Println(v4, v30, v29)
            }
        }
        _ = fi31
        _ = v33
    }
    var v42 float64 = -2.25
    v4 += v4
    for fi43 := 0; fi43 < 8; fi43++ {
        if (v29 == v29) {
            if (float32(-2.25) <= (0.0 * (((0.0 / ((0.0 / -2.25) - -2.25)) + 0.0) + (1e10 * 1.5)))) {
                for fi44 := 0; fi44 < 11; fi44++ {
                    v45 := GenS8{F0: 1, F1: -2147483648}
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    v46 := GenS6{F0: 42, F1: v3}
                    _ = fi44
                    _ = v45
                    _ = v46
                }
                for fi47 := 0; fi47 < 18; fi47++ {
                    if (fi43 <= (fi43 + (fi43 | fi47))) {
                        switch "a" {
                        case "hello":
                            v4 = v4
                        case "a":
                            var v48 uint64 = 4294967295
                            if ((((("" + "a") + ("foo bar" + "\n\t")) + ("hello" + "\n\t")) != ("hello" + "hello")) || (!((^(v29)) > (v29 % (((v29 & (-(v29)))) | 1))))) {
                                v1 += 0.0
                                if v3 {
                                    v49 := GenS2{F0: 42, F1: 255}
                                    if (42 < 1) {
                                        v42 -= v1
                                        switch "hello" {
                                        case "a":
                                            v29 += v29
                                        case "foo bar":
                                            if v3 {
                                                fmt.Println(v48, v4)
                                                v48 -= (v48 - 4294967295)
                                                v4 += (v4 >> ((42) & 63))
                                            } else {
                                                fmt.Println(v29, v42, v29)
                                                v29 -= v29
                                                v48 -= v48
                                            }
                                        case "\n\t":
                                            switch 42 {
                                            case 32768:
                                                v50 := GenS9{F0: v30}
                                                _ = v50
                                            default:
                                                var v51 float32 = 1e10
                                                _ = v51
                                            }
                                        default:
                                            v48 = (v48 + v48)
                                        }
                                    } else {
                                        if len(v2) > 0 {
                                            v2[0] = 925
                                        }
                                    }
                                    _ = v49
                                }
                                var v52 int32 = 0
                                _ = v52
                            } else {
                                v2 = append(v2, (fi43 & fi47))
                                v53 := GenS1{F0: ("foo bar" + "hello"), F1: (^(v30))}
                                _ = v53
                            }
                            _ = v48
                        case "":
                            fmt.Println(v4)
                        }
                        _ = len(v2)
                    } else {
                        var v54 int = fi43
                        v55 := GenS8{F0: -128, F1: -811}
                        v29 -= (-(v29))
                        _ = v54
                        _ = v55
                    }
                    if (v3 || true) {
                        switch (fi47 + (fi43 * -2147483648)) {
                        case 4294967295:
                            if (!(v29 < v29)) {
                                v56 := GenS4{F0: (v30 & 885), F1: (fi47 ^ -878), F2: 1}
                                if v3 {
                                    v30 = (v30 % ((19) | 1))
                                    var v57 uint16 = 65535
                                    _ = v57
                                } else {
                                    var v58 uint16 = 65535
                                    switch (fi47 * int(v58)) {
                                    case 0:
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                    }
                                    if len(v2) > 0 {
                                        v2[0] = -9223372036854775808
                                    }
                                    v1 = (v1 / v42)
                                    _ = v58
                                }
                                if (1.5 < 1e10) {
                                    v3 = (2147483647 < fi43)
                                    if ((v29 * (v29 / (((v29 + 0)) | 1))) <= 1000) {
                                        v42 += v42
                                    }
                                } else {
                                    var v59 uint16 = v4
                                    v60 := GenS8{F0: -1, F1: -1}
                                    v29 += (v29 + 1000)
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    v59 -= (v4 | v59)
                                    if v3 {
                                        var v61 float64 = ((((1e10 - (-2.25 * v42)) / ((float64(float32(v1)) + v1) + ((1.5 + ((-2.25 + (v42 + (v1 + -2.25))) / v1)) - v42))) - (float64(-2.25) - v42)) / v1)
                                        var v62 uint32 = 108
                                        var v63 uint16 = 32767
                                        _ = v61
                                        _ = v62
                                        _ = v63
                                    }
                                    _ = v59
                                    _ = v60
                                }
                                if v3 {
                                    if len(v2) > 0 {
                                        v2[0] = -2147483648
                                    }
                                    if v3 {
                                        switch "foo bar" {
                                        case "hello":
                                            v1 = v42
                                        default:
                                            var v64 bool = ((fi47 + (^(fi43))) < 127)
                                            _ = v64
                                        }
                                        fmt.Println(fi47, v1, v1)
                                        if false {
                                            if ((((0.0 - float64(0.0)) - -2.25) < v42) && (v3 || v3)) {
                                                var v65 uint = 18446744073709551615
                                                var v66 int32 = 42
                                                _ = v65
                                                _ = v66
                                            } else {
                                                var v67 string = ""
                                                v29 -= 1
                                                v3 = (float32((v1 + v42)) >= 1.5)
                                                var v68 uint8 = 128
                                                _ = v67
                                                _ = v68
                                            }
                                            v42 = -2.25
                                            v30 = v30
                                        }
                                    }
                                } else {
                                    if v3 {
                                        if (65535 > 2147483647) {
                                            var v69 float32 = 3.14159
                                            v70 := GenS10{F0: 42, F1: 493}
                                            _ = v69
                                            _ = v70
                                        }
                                        v3 = false
                                        v30 -= (^(v30))
                                    }
                                    v71 := GenS2{F0: fi43, F1: 127}
                                    _ = v71
                                }
                                _ = v56
                            }
                        default:
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        }
                    }
                    _ = fi47
                }
            } else {
                if (1e10 < 0.0) {
                    var v72 float64 = v42
                    for fi73 := 0; fi73 < 7; fi73++ {
                        v2 = append(v2, fi73)
                        v74 := GenS11{F0: "", F1: 335}
                        _ = fi73
                        _ = v74
                    }
                    _ = v72
                } else {
                    switch "a" {
                    case "hello":
                        _ = len(v2)
                    case "":
                        var v75 int = 500
                        _ = v75
                    case "\n\t":
                        v29 += v29
                    default:
                        var v76 int32 = -607
                        _ = v76
                    }
                    var v77 int = 2147483647
                    _ = v77
                }
            }
        } else {
            var v78 uint32 = 2147483648
            var v79 float64 = v1
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = v78
            _ = v79
        }
        switch (fi43 << ((498) & 63)) {
        case 9223372036854775807:
            for fi80 := 0; fi80 < 19; fi80++ {
                if v3 {
                    var v81 string = ("hello" + "\n\t")
                    v82 := GenS12{F0: fi80, F1: 704}
                    var v83 float32 = 1e10
                    v3 = ((!(true && (!(!(v3 || v3))))) && v3)
                    _ = v81
                    _ = v82
                    _ = v83
                }
                var v84 uint = 508
                if false {
                    v1 = (v42 + (v1 - (0.1 * 0.1)))
                }
                var v85 uint32 = 0
                var v86 uint64 = 4294967295
                _ = fi80
                _ = v84
                _ = v85
                _ = v86
            }
        default:
            if v3 {
                switch fi43 {
                case -533:
                    v30 = (-(v30))
                default:
                    for fi87 := 0; fi87 < 6; fi87++ {
                        var v88 uint32 = 202
                        var v89 uint64 = 1
                        if (127 <= -30) {
                            fmt.Println(v30, v29, fi43)
                        }
                        _ = fi87
                        _ = v88
                        _ = v89
                    }
                }
            } else {
                if v3 {
                    var v90 uint = 42
                    _ = v90
                } else {
                    v4 -= 32767
                    v30 = 42
                    for fi91 := 0; fi91 < 20; fi91++ {
                        var v92 uint64 = 922
                        _ = fi91
                        _ = v92
                    }
                }
                v42 = v1
            }
        }
        for fi93 := 0; fi93 < 14; fi93++ {
            fmt.Println(v1)
            v94 := GenS5{F0: 176}
            _ = fi93
            _ = v94
        }
        _ = fi43
    }
    fmt.Println(v30, v1, v30)
    v95 := GenS6{F0: 42, F1: (!(false && (42 <= 1)))}
    if ((v29 ^ v29) != 1000) {
        if v3 {
            v4 = 65535
            if (723 < 1) {
                v4 -= v4
                if v3 {
                    v96 := GenS2{F0: 32768, F1: 42}
                    var v97 float32 = (0.0 - ((float32(((v1 * v1) * v1)) + (1e10 / ((-2.25 / 1.5) / -2.25))) / 1e10))
                    _ = v96
                    _ = v97
                }
                var v98 uint64 = 1
                _ = v98
            }
            for fi99 := 0; fi99 < 15; fi99++ {
                fmt.Println(fi99)
                v2 = append(v2, fi99)
                for fi100 := 0; fi100 < 4; fi100++ {
                    if len(v2) > 0 {
                        v2[0] = (fi100 * fi99)
                    }
                    if ("" == "\n\t") {
                        fmt.Println(fi99, v29, v1)
                    }
                    _ = fi100
                }
                _ = fi99
            }
        }
        v101 := GenS2{F0: 9223372036854775807, F1: 2}
        _ = v101
    }
    v1 = v42
    v29 -= (v29 + v29)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v42
    _ = v95
}
