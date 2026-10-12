package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int }
type GenS2 struct { F0 uint16; F1 int16; F2 int }
type GenS3 struct { F0 uint64; F1 string }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int16; F1 uint8; F2 uint32 }
type GenS6 struct { F0 uint32; F1 uint64 }
type GenS7 struct { F0 uint }
func genFn0(p0 bool, p1 uint, p2 int32) int8 {
    var v5 bool = (p2 == (p2 % ((p2) | 1)))
    switch 603 {
    case 115:
        p1 -= (p1 - 1)
    case 1:
        p1 = p1
    default:
        if (2 == 0) {
            v6 := []int{}
            if (!p0) {
                if len(v6) > 0 {
                    v6[0] = -192
                }
                var v7 string = "a"
                _ = v7
            } else {
                _ = len(v6)
            }
            v8 := GenS1{F0: p2, F1: -865}
            _ = v6
            _ = v8
        }
    }
    p2 = (p2 ^ (p2 | -361))
    for fi9 := 0; fi9 < 16; fi9++ {
        for fi10 := 0; fi10 < 4; fi10++ {
            var v11 float64 = (1.5 * 1e10)
            _ = fi10
            _ = v11
        }
        _ = fi9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 0
}
func genFn1(p0 uint, p1 uint8) uint32 {
    if false {
        for fi12 := 0; fi12 < 17; fi12++ {
            var v13 uint8 = p1
            _ = fi12
            _ = v13
        }
    } else {
        v14 := []int{-1, 2147483647, -2147483648, -50}
        _ = v14
    }
    for fi15 := 0; fi15 < 17; fi15++ {
        p0 = 42
        var v16 uint64 = 1
        _ = fi15
        _ = v16
    }
    var v17 bool = false
    var v18 int8 = 1
    _ = p0
    _ = p1
    _ = v17
    _ = v18
    return uint32(p1)
}

func main() {
    v1 := []int{}
    var v2 int32 = 337
    var v3 bool = true
    v4 := make(map[int]string)
    v3 = (1 == v2)
    var v19 uint16 = 255
    _ = len(v4)
    v20 := GenS2{F0: (v19 + 625), F1: 42, F2: 127}
    fmt.Println(v3, v2, v19)
    var v21 float64 = (((3.14159 + (1e10 * (0.0 * (float64((3.14159 + ((0.0 * (1.5 / (1.5 / -2.25))) - 1.5))) + 3.14159)))) - float64(3.14159)) / 1.0)
    for fi22 := 0; fi22 < 10; fi22++ {
        var v23 uint = 65535
        _ = fi22
        _ = v23
    }
    v2 += (v2 / ((v2) | 1))
    if (((v3 && (true && true)) || (!((-117 > int8(v19)) || false))) && true) {
        switch -2147483648 {
        case 42:
            switch "\n\t" {
            case "":
                var v24 int = 1
                _ = v24
            case "a":
                v25 := GenS2{F0: v19, F1: -612, F2: 0}
                _ = v25
            case "\n\t":
                var v26 int16 = 1
                _ = v26
            }
        case 2147483647:
            for fi27 := 0; fi27 < 14; fi27++ {
                var v28 uint16 = 42
                v28 = (v28 & 32767)
                v29 := GenS2{F0: (v19 - (v19 << ((2) & 63))), F1: -1, F2: (fi27 & (fi27 << ((65535) & 63)))}
                _ = fi27
                _ = v28
                _ = v29
            }
        }
        v3 = false
    } else {
        if true {
            v30 := GenS3{F0: 2, F1: "foo bar"}
            _ = v30
        } else {
            v2 = v2
            v19 -= (^(v19))
        }
        for fi31 := 0; fi31 < 2; fi31++ {
            switch "" {
            case "hello":
                for fi32 := 0; fi32 < 9; fi32++ {
                    if false {
                        var v33 int64 = 592
                        if v3 {
                            v21 -= 3.14159
                        } else {
                            if v3 {
                                v34 := GenS2{F0: (v19 - (v19 * (v19 / (((v19 / ((v19) | 1))) | 1)))), F1: 32767, F2: 9223372036854775807}
                                _ = v34
                            } else {
                                if (false || (!(v33 != (v33 << ((425) & 63))))) {
                                    v21 = ((((-2.25 - (1.5 - float64((-2.25 - 0.0)))) - (v21 * ((-2.25 / 1.5) - float64(3.14159)))) + 1e10) + (v21 / v21))
                                    var v35 uint16 = v19
                                    v21 = (3.14159 - (1e10 - 1e10))
                                    switch fi32 {
                                    case 0:
                                        var v36 uint32 = 2147483648
                                        _ = v36
                                    case -1:
                                        _ = len(v1)
                                    case 32768:
                                        v37 := GenS3{F0: 42, F1: ("" + ((("foo bar" + "a") + (("" + "a") + ((("foo bar" + "a") + ("" + "")) + "foo bar"))) + "\n\t"))}
                                        _ = v37
                                    default:
                                        _ = len(v4)
                                    }
                                    fmt.Println(v3)
                                    if (v3 && true) {
                                        delete(v4, (-(fi32)))
                                        switch "foo bar" {
                                        case "a":
                                            delete(v4, (fi31 & (fi32 - fi31)))
                                        }
                                    } else {
                                        var v38 uint8 = 255
                                        v19 += (v35 ^ v19)
                                        v39 := GenS4{F0: 4294967295}
                                        _ = v38
                                        _ = v39
                                    }
                                    var v40 int32 = (v2 & (^(v2)))
                                    _ = v35
                                    _ = v40
                                } else {
                                    v41 := GenS2{F0: (v19 & (v19 - v19)), F1: 32767, F2: -122}
                                    fmt.Println(fi31, v3)
                                    v19 = v19
                                    _ = v41
                                }
                            }
                            if (false && ((252 == 1) || v3)) {
                                switch "" {
                                case "foo bar":
                                    if len(v1) > 0 {
                                        v1[0] = fi31
                                    }
                                default:
                                    switch ("foo bar" + "a") {
                                    case "foo bar":
                                        var v42 float32 = (float32(1e10) / 1.0)
                                        _ = v42
                                    case "hello":
                                        fmt.Println(v19, fi32, fi31)
                                    }
                                }
                                v43 := GenS4{F0: 678}
                                fmt.Println(v2)
                                _ = v43
                            }
                            if (-842 > v33) {
                                var v44 uint32 = 65535
                                if (v3 && v3) {
                                    var v45 bool = v3
                                    v2 += v2
                                    var v46 int = (fi31 >> ((64) & 63))
                                    _ = v45
                                    _ = v46
                                } else {
                                    var v47 string = ("\n\t" + "")
                                    if false {
                                        switch fi32 {
                                        case 1:
                                            v48 := GenS5{F0: 0, F1: 255, F2: (v44 + (v44 % ((v44) | 1)))}
                                            _ = v48
                                        case 32767:
                                            v49 := GenS4{F0: 42}
                                            _ = v49
                                        default:
                                            var v50 uint8 = 2
                                            _ = v50
                                        }
                                    }
                                    var v51 int8 = 127
                                    _ = v47
                                    _ = v51
                                }
                                v33 += 811
                                var v52 float64 = 1.5
                                _ = v44
                                _ = v52
                            }
                        }
                        if v3 {
                            switch (fi31 | 3) {
                            case -9223372036854775808:
                                var v53 uint = 18446744073709551615
                                _ = v53
                            default:
                                var v54 bool = (!(!(v3 && ((v19 * v19) < (v19 * v19)))))
                                _ = v54
                            }
                            if len(v1) > 0 {
                                v1[0] = 32768
                            }
                            _ = len(v4)
                            v3 = v3
                            switch (fi32 & fi31) {
                            case 326:
                                var v55 uint64 = 2
                                _ = v55
                            }
                        } else {
                            v21 -= 0.0
                            if v3 {
                                var v56 int = 701
                                v57 := GenS4{F0: 18446744073709551615}
                                v58 := GenS4{F0: 4294967295}
                                v3 = (((9223372036854775807 < 4294967295) && ((v3 || v3) && v3)) || ((!(("hello" + "hello") != "a")) && v3))
                                if (0.1 == -2.25) {
                                    v59 := GenS6{F0: 2147483647, F1: 0}
                                    _ = v59
                                }
                                _ = v56
                                _ = v57
                                _ = v58
                            }
                        }
                        _ = v33
                    }
                    fmt.Println(v19, v2, v2)
                    switch (fi32 << ((4294967295) & 63)) {
                    case -2147483648:
                        if len(v1) > 0 {
                            v1[0] = -239
                        }
                    }
                    _ = fi32
                }
            default:
                if (42 >= genFn0(false, 42, 2147483647)) {
                    v3 = (genFn0(((1e10 + (3.14159 * 0.0)) <= (1.5 - 1e10)), 18446744073709551615, (v2 ^ -730)) < genFn0(v3, 42, v2))
                    fmt.Println(fi31, v2, fi31)
                    fmt.Println(v3)
                }
            }
            if true {
                v19 -= (v19 - (v19 + (v19 * v19)))
                v19 = (^(v19))
                fmt.Println(v3)
            } else {
                if true {
                    var v60 int32 = v2
                    v61 := GenS1{F0: (v2 >> ((65535) & 63)), F1: fi31}
                    var v62 int64 = 1
                    _ = v60
                    _ = v61
                    _ = v62
                } else {
                    v2 = 2147483647
                    v19 = v19
                }
                if ((!v3) && ((v21 != 0.0) && (fi31 > fi31))) {
                    fmt.Println(v19, v3, v21)
                }
                var v63 int16 = 42
                _ = v63
            }
            var v64 int16 = 1000
            _ = fi31
            _ = v64
        }
        for fi65 := 0; fi65 < 20; fi65++ {
            fmt.Println(v19)
            if false {
                for fi66 := 0; fi66 < 5; fi66++ {
                    if (0 < 65535) {
                        _ = len(v1)
                        v21 -= 0.0
                        v3 = (65535 != 65535)
                        switch "a" {
                        case "":
                            switch "a" {
                            case "a":
                                var v67 uint64 = 1
                                _ = v67
                            case "foo bar":
                                switch 171 {
                                case 2147483647:
                                    v3 = false
                                case -466:
                                    v68 := v4[9223372036854775807]
                                    _ = v68
                                }
                            }
                        case "hello":
                            var v69 int16 = 0
                            _ = v69
                        case "a":
                            if true {
                                var v70 int32 = (^(v2))
                                var v71 uint64 = 1
                                v72 := GenS3{F0: 2, F1: "\n\t"}
                                _ = v70
                                _ = v71
                                _ = v72
                            }
                        }
                        v3 = v3
                    }
                    _ = fi66
                }
            } else {
                v2 += v2
                for fi73 := 0; fi73 < 2; fi73++ {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    fmt.Println(fi65, v2, fi73)
                    _ = fi73
                }
            }
            if (18446744073709551615 <= 1) {
                switch fi65 {
                case 32768:
                    switch "hello" {
                    case "\n\t":
                        for fi74 := 0; fi74 < 1; fi74++ {
                            if false {
                                if len(v1) > 0 {
                                    v1[0] = (fi74 + fi74)
                                }
                                v1 = append(v1, (fi65 % ((-942) | 1)))
                                v2 = 0
                            } else {
                                v2 = v2
                                v2 = (v2 << ((537) & 63))
                                v2 += (^(v2))
                                v4[fi74] = "hello"
                            }
                            _ = fi74
                        }
                    case "hello":
                        for fi75 := 0; fi75 < 11; fi75++ {
                            var v76 int32 = v2
                            v77 := GenS1{F0: -1, F1: 670}
                            fmt.Println(v21, v2)
                            _ = fi75
                            _ = v76
                            _ = v77
                        }
                    case "":
                        var v78 float32 = 0.1
                        _ = v78
                    default:
                        var v79 float32 = 3.14159
                        _ = v79
                    }
                case 4294967295:
                    _ = len(v4)
                case -996:
                    if v3 {
                        v2 += (v2 ^ v2)
                        for fi80 := 0; fi80 < 14; fi80++ {
                            var v81 float32 = (0.1 / 0.1)
                            _ = fi80
                            _ = v81
                        }
                        switch ("foo bar" + "\n\t") {
                        case "hello":
                            v3 = (!(255 <= 586))
                        }
                    } else {
                        var v82 uint = 65535
                        var v83 int = (fi65 + fi65)
                        _ = v82
                        _ = v83
                    }
                default:
                    if (int(v19) == (^(fi65))) {
                        var v84 uint = 466
                        _ = v84
                    } else {
                        var v85 float32 = 1e10
                        fmt.Println(v2, v2)
                        v86 := GenS7{F0: 469}
                        _ = v85
                        _ = v86
                    }
                }
            } else {
                var v87 uint16 = 10
                if len(v1) > 0 {
                    v1[0] = fi65
                }
                _ = v87
            }
            _ = fi65
        }
    }
    switch 0 {
    case -964:
        var v88 int32 = v2
        _ = v88
    }
    v89 := GenS6{F0: uint32(v19), F1: 4294967295}
    if (798 < 0) {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        var v90 uint64 = 2
        _ = v90
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if v3 {
        v21 += v21
    } else {
        v21 = v21
        if (!v3) {
            v2 = 1
            var v91 uint32 = 2147483648
            _ = v91
        } else {
            var v92 uint8 = 0
            for fi93 := 0; fi93 < 1; fi93++ {
                for fi94 := 0; fi94 < 5; fi94++ {
                    var v95 int32 = -2147483648
                    _ = fi94
                    _ = v95
                }
                var v96 uint = 6
                _ = fi93
                _ = v96
            }
            _ = v92
        }
    }
    var v97 int16 = 32767
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v97:", v97)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v89
    _ = v97
}
