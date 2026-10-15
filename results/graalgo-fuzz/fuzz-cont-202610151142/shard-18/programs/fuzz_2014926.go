package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool; F2 string }
type GenS2 struct { F0 uint64; F1 int32; F2 uint64 }
type GenS3 struct { F0 uint; F1 uint; F2 uint64 }
type GenS4 struct { F0 uint8; F1 bool }
type GenS5 struct { F0 uint; F1 int8; F2 int8 }
type GenS6 struct { F0 int8 }
func genFn0(p0 bool, p1 uint, p2 uint32) uint64 {
    for fi6 := 0; fi6 < 20; fi6++ {
        if (!((0.1 / -2.25) != 0.0)) {
            v7 := GenS1{F0: p0, F1: (!p0), F2: ("" + "foo bar")}
            _ = v7
        } else {
            p1 += (p1 - 1)
            fmt.Println(p0, fi6)
        }
        var v8 float32 = ((float32(float64((((1e10 / -2.25) * (1.5 * 0.1)) + 1.5))) * 0.1) * 0.0)
        v9 := []int{-405, fi6}
        _ = fi6
        _ = v8
        _ = v9
    }
    v10 := []int{2147483647, 32767, 42}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return 1
}
func genFn1(p0 uint, p1 bool) int16 {
    p1 = p1
    var v11 uint64 = genFn0(p1, (p0 ^ p0), 1)
    for fi12 := 0; fi12 < 1; fi12++ {
        v11 -= 0
        for fi13 := 0; fi13 < 15; fi13++ {
            v11 = 4294967295
            var v14 int = (fi13 << ((65535) & 63))
            _ = fi13
            _ = v14
        }
        _ = fi12
    }
    if (p1 || (v11 < (^(v11)))) {
        if (0 < -2147483648) {
            p0 -= (p0 << (((p0 % ((255) | 1))) & 63))
            p1 = false
        }
        v15 := []int{904}
        v11 -= (v11 << ((1) & 63))
        _ = v15
    }
    p1 = (2 < 65535)
    _ = p0
    _ = p1
    _ = v11
    return 1
}

func main() {
    var v1 int8 = 1
    var v2 uint32 = 2147483648
    var v3 bool = (0 == 753)
    var v4 float64 = float64((0.0 + (3.14159 + 1.5)))
    v5 := []int{-275, 1}
    v3 = (v3 || (true && v3))
    v1 = 42
    for fi16 := 0; fi16 < 9; fi16++ {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        v17 := GenS1{F0: v3, F1: v3, F2: ("hello" + "hello")}
        v3 = (!v3)
        _ = fi16
        _ = v17
    }
    switch (("foo bar" + "") + ("\n\t" + (("hello" + "foo bar") + "hello"))) {
    case "hello":
        fmt.Println(v2)
    default:
        v18 := GenS1{F0: true, F1: false, F2: "foo bar"}
        _ = v18
    }
    if (int64(v1) != 42) {
        var v19 uint64 = 4294967295
        _ = v19
    } else {
        fmt.Println(v1, v3, v4)
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        if v3 {
            if v3 {
                fmt.Println(v3)
                v2 += v2
                fmt.Println(v3, v3, v1)
            } else {
                for fi20 := 0; fi20 < 2; fi20++ {
                    if (v1 == -1) {
                        switch fi20 {
                        case 1:
                            if v3 {
                                fmt.Println(fi20)
                                v21 := GenS2{F0: 42, F1: -2147483648, F2: 4294967295}
                                v22 := GenS2{F0: 512, F1: -1, F2: 2}
                                _ = v21
                                _ = v22
                            }
                        case -785:
                            fmt.Println(fi20, fi20, v3)
                        }
                    }
                    _ = fi20
                }
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                for fi23 := 0; fi23 < 5; fi23++ {
                    v24 := GenS3{F0: 4294967295, F1: 1, F2: uint64(v2)}
                    switch (fi23 + fi23) {
                    case -9223372036854775808:
                        for fi25 := 0; fi25 < 5; fi25++ {
                            if v3 {
                                v3 = true
                            }
                            var v26 uint64 = 0
                            if ((!((2 < 334) && ((!v3) || v3))) || v3) {
                                if v3 {
                                    _ = len(v5)
                                    switch fi23 {
                                    case -166:
                                        v27 := GenS1{F0: v3, F1: v3, F2: ""}
                                        _ = v27
                                    case 32768:
                                        v2 = (v2 & (v2 / (((^(v2))) | 1)))
                                    case 127:
                                        switch (fi25 * fi25) {
                                        case 9223372036854775807:
                                            v28 := GenS4{F0: 42, F1: (9223372036854775807 > 9223372036854775807)}
                                            _ = v28
                                        case 32767:
                                            var v29 uint32 = (^(v2))
                                            _ = v29
                                        default:
                                            v1 = v1
                                        }
                                        v30 := GenS2{F0: v26, F1: -80, F2: (v26 - v26)}
                                        var v31 int16 = -1
                                        _ = v30
                                        _ = v31
                                    }
                                    if v3 {
                                        v32 := GenS1{F0: true, F1: false, F2: ("foo bar" + "")}
                                        fmt.Println(v3, fi25)
                                        v33 := GenS5{F0: 0, F1: v1, F2: v1}
                                        if v3 {
                                            _ = len(v5)
                                        }
                                        v34 := GenS1{F0: (32767 >= 1000), F1: ((^(v2)) <= (v2 / ((v2) | 1))), F2: "\n\t"}
                                        _ = v32
                                        _ = v33
                                        _ = v34
                                    }
                                } else {
                                    v4 = 3.14159
                                    var v35 int8 = v1
                                    v36 := GenS6{F0: -10}
                                    var v37 uint8 = 4
                                    _ = v35
                                    _ = v36
                                    _ = v37
                                }
                            }
                            _ = fi25
                            _ = v26
                        }
                    case 42:
                        var v38 float64 = -2.25
                        _ = v38
                    case -59:
                        if false {
                            v4 -= v4
                        }
                    default:
                        switch ("" + "\n\t") {
                        case "\n\t":
                            switch -422 {
                            case 9223372036854775807:
                                fmt.Println(fi23, v1)
                            case 1:
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                            case 32767:
                                switch (fi23 >> ((4294967295) & 63)) {
                                case -9223372036854775808:
                                    v2 -= (v2 << ((0) & 63))
                                case 128:
                                    if true {
                                        for fi39 := 0; fi39 < 16; fi39++ {
                                            var v40 int32 = 42
                                            v3 = true
                                            var v41 uint8 = 255
                                            _ = fi39
                                            _ = v40
                                            _ = v41
                                        }
                                        var v42 string = "foo bar"
                                        if (v3 || v3) {
                                            var v43 bool = (42 >= v1)
                                            _ = v43
                                        } else {
                                            fmt.Println(v42, v42, v4)
                                        }
                                        _ = v42
                                    } else {
                                        if v3 {
                                            var v44 uint64 = 18446744073709551615
                                            if true {
                                                if (!((v1 + v1) != v1)) {
                                                    if ((-179 >= 9223372036854775807) && (false || v3)) {
                                                        if v3 {
                                                            if len(v5) > 0 {
                                                                v5[0] = fi23
                                                            }
                                                            var v45 uint16 = 892
                                                            var v46 uint = 482
                                                            _ = v45
                                                            _ = v46
                                                        } else {
                                                            for fi47 := 0; fi47 < 3; fi47++ {
                                                                var v48 uint32 = 2
                                                                v2 += (v2 + v2)
                                                                v4 = v4
                                                                _ = fi47
                                                                _ = v48
                                                            }
                                                            for fi49 := 0; fi49 < 8; fi49++ {
                                                                if v3 {
                                                                    if (42 >= 957) {
                                                                        if len(v5) >= 2 {
                                                                            v5 = v5[1:len(v5)]
                                                                        }
                                                                        v1 = 127
                                                                        v1 = int8(v44)
                                                                        _ = len(v5)
                                                                        v3 = (int16(v44) < -1)
                                                                    }
                                                                }
                                                                if v3 {
                                                                    switch (fi49 + fi23) {
                                                                    case -883:
                                                                        if (0 <= 65535) {
                                                                            if v3 {
                                                                                var v50 uint = 42
                                                                                v51 := GenS4{F0: 201, F1: v3}
                                                                                if false {
                                                                                    if v3 {
                                                                                        v2 = (v2 * (v2 << ((v50) & 63)))
                                                                                    }
                                                                                    switch fi49 {
                                                                                    case 32767:
                                                                                        var v52 int64 = -9223372036854775808
                                                                                        _ = v52
                                                                                    default:
                                                                                        v3 = (-9223372036854775808 >= -9223372036854775808)
                                                                                    }
                                                                                    if len(v5) > 0 {
                                                                                        v5[0] = 9223372036854775807
                                                                                    }
                                                                                }
                                                                                _ = v50
                                                                                _ = v51
                                                                            } else {
                                                                                var v53 uint = 1
                                                                                if v3 {
                                                                                    var v54 uint64 = v44
                                                                                    v54 = genFn0((!(139 < 1)), 893, 2147483648)
                                                                                    if (v3 || (v3 || false)) {
                                                                                        v55 := GenS3{F0: 65535, F1: (^(v53)), F2: v44}
                                                                                        if v3 {
                                                                                            var v56 uint32 = v2
                                                                                            fmt.Println(v54)
                                                                                            if false {
                                                                                                if v3 {
                                                                                                    v56 = (^(v2))
                                                                                                }
                                                                                                v57 := GenS1{F0: (v1 <= v1), F1: (!(!v3)), F2: "foo bar"}
                                                                                                v54 -= 4294967295
                                                                                                _ = v57
                                                                                            }
                                                                                            _ = v56
                                                                                        }
                                                                                        fmt.Println(fi23, v44)
                                                                                        _ = v55
                                                                                    }
                                                                                    _ = v54
                                                                                }
                                                                                _ = v53
                                                                            }
                                                                            var v58 int16 = 42
                                                                            _ = v58
                                                                        } else {
                                                                            if v3 {
                                                                                if true {
                                                                                    v2 -= (v2 - (v2 * v2))
                                                                                    v4 = v4
                                                                                    v5 = append(v5, fi23)
                                                                                }
                                                                                if v3 {
                                                                                    var v59 int16 = -32768
                                                                                    _ = v59
                                                                                }
                                                                                if v3 {
                                                                                    v60 := GenS4{F0: 225, F1: (v4 != v4)}
                                                                                    v3 = v3
                                                                                    _ = v60
                                                                                } else {
                                                                                    v44 += (v44 - uint64(v2))
                                                                                    var v61 uint32 = (v2 & 2147483648)
                                                                                    var v62 uint16 = 2
                                                                                    fmt.Println(v61)
                                                                                    switch "a" {
                                                                                    case "foo bar":
                                                                                        v5 = append(v5, 697)
                                                                                    }
                                                                                    _ = v61
                                                                                    _ = v62
                                                                                }
                                                                            }
                                                                            switch fi23 {
                                                                            case 329:
                                                                                var v63 float32 = 0.0
                                                                                _ = v63
                                                                            case 32768:
                                                                                var v64 float32 = (1.5 / 1e10)
                                                                                _ = v64
                                                                            }
                                                                        }
                                                                    case 32768:
                                                                        if len(v5) >= 2 {
                                                                            v5 = v5[1:len(v5)]
                                                                        }
                                                                    default:
                                                                        fmt.Println(v1, v4)
                                                                    }
                                                                } else {
                                                                    var v65 int32 = 710
                                                                    _ = v65
                                                                }
                                                                _ = fi49
                                                            }
                                                        }
                                                    }
                                                    var v66 float32 = float32((1.5 * v4))
                                                    _ = v66
                                                }
                                            }
                                            _ = v44
                                        }
                                        if len(v5) > 0 {
                                            v5[0] = fi23
                                        }
                                    }
                                default:
                                    v1 = 127
                                }
                            default:
                                v2 += (v2 << ((255) & 63))
                            }
                        case "hello":
                            _ = len(v5)
                        case "":
                            v67 := GenS2{F0: 4294967295, F1: 42, F2: 0}
                            _ = v67
                        }
                    }
                    var v68 int32 = 0
                    _ = fi23
                    _ = v24
                    _ = v68
                }
            }
        } else {
            if v3 {
                v69 := GenS2{F0: 2, F1: int32(v2), F2: genFn0((v1 != (v1 << ((467) & 63))), 0, 4294967295)}
                v4 += float64(0.1)
                _ = v69
            } else {
                for fi70 := 0; fi70 < 3; fi70++ {
                    switch (^(fi70)) {
                    case 2147483647:
                        for fi71 := 0; fi71 < 12; fi71++ {
                            if false {
                                fmt.Println(v1)
                                v5 = append(v5, fi70)
                            }
                            _ = fi71
                        }
                    }
                    fmt.Println(v3)
                    _ = fi70
                }
                if true {
                    var v72 uint64 = 42
                    _ = v72
                }
            }
            fmt.Println(v3)
        }
    }
    if v3 {
        v1 = v1
        if ((195 > 1) || v3) {
            for fi73 := 0; fi73 < 4; fi73++ {
                v3 = false
                if false {
                    if len(v5) > 0 {
                        v5[0] = (-(fi73))
                    }
                    var v74 int32 = 664
                    _ = v74
                }
                fmt.Println(v2)
                _ = fi73
            }
            v2 += (^(v2))
        }
        v3 = ((-(v1)) == v1)
    } else {
        var v75 bool = v3
        _ = v75
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    if (!false) {
        v3 = v3
        v1 = (^(v1))
    }
    fmt.Println(v4, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
