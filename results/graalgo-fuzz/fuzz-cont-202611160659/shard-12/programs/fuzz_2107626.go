package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint32; F1 uint16; F2 uint32 }
type GenS4 struct { F0 int16; F1 int8 }
type GenS5 struct { F0 int16; F1 uint32; F2 uint8 }
type GenS6 struct { F0 uint8; F1 string; F2 uint8 }
type GenS7 struct { F0 string; F1 uint }
type GenS8 struct { F0 string }
type GenS9 struct { F0 uint8 }
type GenS10 struct { F0 int8 }
type GenS11 struct { F0 uint8; F1 uint8 }
type GenS12 struct { F0 uint8; F1 uint16 }
type GenS13 struct { F0 uint8 }
type GenS14 struct { F0 string; F1 int16; F2 int64 }
type GenS15 struct { F0 bool; F1 int }
type GenS16 struct { F0 uint8; F1 int32 }
type GenS17 struct { F0 uint16; F1 int64; F2 uint8 }
func genFn0(p0 uint8) int8 {
    var v6 int32 = -1
    if (0 > 4294967295) {
        var v7 uint8 = p0
        _ = v7
    } else {
        var v8 float32 = 0.0
        if false {
            for fi9 := 0; fi9 < 18; fi9++ {
                switch (("" + ("hello" + "foo bar")) + (("\n\t" + "a") + "")) {
                case "hello":
                    var v10 int8 = 0
                    _ = v10
                case "\n\t":
                    var v11 float32 = ((float32((1.5 * 0.1)) * (0.1 + 3.14159)) / v8)
                    _ = v11
                default:
                    var v12 float32 = 3.14159
                    _ = v12
                }
                for fi13 := 0; fi13 < 20; fi13++ {
                    v6 = 0
                    v14 := GenS1{F0: -9223372036854775808, F1: 255}
                    _ = fi13
                    _ = v14
                }
                v15 := []int{(-(fi9))}
                _ = fi9
                _ = v15
            }
            v8 = v8
            switch 1 {
            case -1:
                fmt.Println(v6)
            case 672:
                var v16 uint32 = 2
                _ = v16
            case 898:
                v17 := GenS2{F0: 1}
                _ = v17
            default:
                v18 := GenS3{F0: 4294967295, F1: 665, F2: 1}
                _ = v18
            }
        }
        var v19 int = 42
        _ = v19
        _ = v8
    }
    switch 128 {
    case 32768:
        for fi20 := 0; fi20 < 11; fi20++ {
            v6 -= (v6 * v6)
            fmt.Println(p0, p0, p0)
            _ = fi20
        }
    case 127:
        v6 -= -1
    default:
        fmt.Println(v6, v6)
    }
    v21 := GenS1{F0: 9223372036854775807, F1: 255}
    _ = p0
    _ = v21
    _ = v6
    return int8(v6)
}
func genFn1() uint64 {
    for fi22 := 0; fi22 < 11; fi22++ {
        if ((float64((1.5 + 0.0)) * 3.14159) <= 3.14159) {
            var v23 int8 = 127
            var v24 int32 = -1
            _ = v23
            _ = v24
        } else {
            switch "" {
            case "foo bar":
                var v25 uint16 = 2
                _ = v25
            default:
                v26 := GenS4{F0: 0, F1: 127}
                _ = v26
            }
            for fi27 := 0; fi27 < 7; fi27++ {
                var v28 uint64 = 1
                switch "foo bar" {
                case "a":
                    v29 := GenS5{F0: -260, F1: 285, F2: 42}
                    _ = v29
                }
                switch (fi27 ^ fi27) {
                case 32768:
                    v30 := GenS6{F0: 171, F1: ("foo bar" + "a"), F2: 2}
                    _ = v30
                case 0:
                    var v31 bool = ((-1 == -636) || (false && (genFn0(1) < 100)))
                    _ = v31
                case 936:
                    fmt.Println(fi22, fi22)
                }
                _ = fi27
                _ = v28
            }
            v32 := GenS6{F0: 0, F1: "a", F2: 62}
            _ = v32
        }
        _ = fi22
    }
    v33 := GenS6{F0: 83, F1: "\n\t", F2: 128}
    v34 := []int{2147483647, -9223372036854775808, -989}
    _ = v33
    _ = v34
    return 81
}

func main() {
    var v1 int64 = 9223372036854775807
    var v2 uint8 = 128
    var v3 bool = (!false)
    v4 := []uint32{309}
    var v5 bool = (!(-509 > -1))
    v35 := GenS6{F0: v2, F1: ("\n\t" + ""), F2: v2}
    var v36 float64 = 1.5
    if v3 {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        if true {
            var v37 bool = (false && (1e10 > 0.0))
            var v38 bool = (65535 != 65535)
            fmt.Println(v38, v1)
            _ = v37
            _ = v38
        }
    }
    if (v5 && (255 >= (v2 | 0))) {
        switch "\n\t" {
        case "a":
            switch -543 {
            case 127:
                for fi39 := 0; fi39 < 20; fi39++ {
                    switch "\n\t" {
                    case "\n\t":
                        v3 = (v1 != 568)
                    case "hello":
                        v40 := GenS7{F0: ("" + "a"), F1: 255}
                        _ = v40
                    case "a":
                        for fi41 := 0; fi41 < 1; fi41++ {
                            var v42 uint64 = 1
                            fmt.Println(v2, v36, fi41)
                            _ = fi41
                            _ = v42
                        }
                    }
                    _ = fi39
                }
            }
        }
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = len(v4)
    }
    for fi43 := 0; fi43 < 18; fi43++ {
        fmt.Println(fi43, v2, fi43)
        if (42 < 18446744073709551615) {
            for fi44 := 0; fi44 < 15; fi44++ {
                v45 := GenS1{F0: v1, F1: 1}
                if (v3 || (("foo bar" + "") <= ("foo bar" + ""))) {
                    if len(v4) > 0 {
                        v4[0] = 2147483648
                    }
                    fmt.Println(fi43, v2)
                    if ((^(v2)) <= 218) {
                        fmt.Println(fi44, v3, fi43)
                    } else {
                        fmt.Println(fi44, fi44)
                    }
                    fmt.Println(v5, v36, fi43)
                    fmt.Println(fi43, v5)
                } else {
                    var v46 int = fi43
                    _ = v46
                }
                _ = fi44
                _ = v45
            }
            var v47 uint32 = uint32(v2)
            _ = v47
        } else {
            if (true && (true || v5)) {
                if (v5 && false) {
                    if false {
                        v48 := GenS8{F0: "\n\t"}
                        if len(v4) > 0 {
                            v4[0] = 385
                        }
                        for fi49 := 0; fi49 < 12; fi49++ {
                            v36 -= float64(1e10)
                            _ = fi49
                        }
                        _ = v48
                    }
                    if (!(764 > 0)) {
                        var v50 uint8 = (v2 & v2)
                        switch 9223372036854775807 {
                        case -1:
                            switch fi43 {
                            case 9223372036854775807:
                                v50 += 127
                            default:
                                for fi51 := 0; fi51 < 16; fi51++ {
                                    if v5 {
                                        if v5 {
                                            if (v3 || v5) {
                                                fmt.Println(fi51, v3, v36)
                                                var v52 int8 = 0
                                                _ = v52
                                            }
                                            if v3 {
                                                v1 = (-(v1))
                                            } else {
                                                v2 = v2
                                                v4 = append(v4, 0)
                                                if v3 {
                                                    switch "hello" {
                                                    case "foo bar":
                                                        if v5 {
                                                            _ = len(v4)
                                                            v53 := GenS8{F0: "\n\t"}
                                                            switch "\n\t" {
                                                            case "\n\t":
                                                                if len(v4) >= 2 {
                                                                    v4 = v4[1:len(v4)]
                                                                }
                                                            default:
                                                                v54 := GenS1{F0: (v1 & 1), F1: 560}
                                                                _ = v54
                                                            }
                                                            _ = v53
                                                        }
                                                        v4 = append(v4, 2147483648)
                                                        v55 := GenS9{F0: 105}
                                                        _ = v55
                                                    case "hello":
                                                        v1 = -9223372036854775808
                                                    }
                                                    v56 := GenS10{F0: 1}
                                                    v3 = v5
                                                    _ = v56
                                                }
                                            }
                                            v57 := GenS11{F0: 127, F1: v50}
                                            _ = v57
                                        } else {
                                            v36 = v36
                                        }
                                        v50 -= 127
                                        switch (fi43 & (fi43 & fi43)) {
                                        case 0:
                                            if v3 {
                                                fmt.Println(v50, v50, v5)
                                                v58 := GenS6{F0: v2, F1: ((("hello" + "hello") + "hello") + ""), F2: (v50 | v2)}
                                                _ = v58
                                            } else {
                                                switch fi43 {
                                                case -464:
                                                    switch int(v2) {
                                                    case 32767:
                                                        var v59 uint8 = v2
                                                        _ = v59
                                                    }
                                                    var v60 uint8 = 253
                                                    switch "foo bar" {
                                                    case "hello":
                                                        var v61 uint32 = 42
                                                        _ = v61
                                                    default:
                                                        if len(v4) >= 2 {
                                                            v4 = v4[1:len(v4)]
                                                        }
                                                    }
                                                    _ = v60
                                                case 32767:
                                                    if false {
                                                        var v62 int64 = (v1 / ((v1) | 1))
                                                        _ = v62
                                                    } else {
                                                        if (0.0 != -2.25) {
                                                            var v63 int16 = -844
                                                            var v64 int32 = -209
                                                            _ = v63
                                                            _ = v64
                                                        } else {
                                                            var v65 int8 = 100
                                                            fmt.Println(fi51, v2, fi51)
                                                            v66 := GenS11{F0: 159, F1: 42}
                                                            var v67 uint8 = v2
                                                            fmt.Println(v2, fi51)
                                                            v2 += v50
                                                            v3 = (32767 < 42)
                                                            _ = v65
                                                            _ = v66
                                                            _ = v67
                                                        }
                                                        _ = len(v4)
                                                    }
                                                    if len(v4) > 0 {
                                                        v4[0] = 2
                                                    }
                                                default:
                                                    v68 := GenS12{F0: (v2 + (v2 * (v50 | 42))), F1: 65535}
                                                    _ = v68
                                                }
                                                v69 := GenS13{F0: v2}
                                                _ = v69
                                            }
                                        case 4294967295:
                                            var v70 int8 = genFn0((v50 / ((v2) | 1)))
                                            _ = v70
                                        }
                                        if (1 == int32(fi43)) {
                                            v71 := GenS5{F0: 1, F1: 1, F2: 179}
                                            if len(v4) > 0 {
                                                v4[0] = 0
                                            }
                                            _ = v71
                                        } else {
                                            v36 += v36
                                        }
                                    } else {
                                        v72 := GenS14{F0: "foo bar", F1: 1, F2: v1}
                                        var v73 uint = 255
                                        v1 = (-(v1))
                                        _ = v72
                                        _ = v73
                                    }
                                    v50 -= v2
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                    _ = fi51
                                }
                            }
                        case 2147483647:
                            v3 = false
                        }
                        switch (-(fi43)) {
                        case 9223372036854775807:
                            if v3 {
                                v50 = v50
                                fmt.Println(fi43, v3, fi43)
                            }
                        case -308:
                            if v3 {
                                v50 -= 186
                            } else {
                                fmt.Println(v36)
                            }
                        }
                        _ = v50
                    } else {
                        var v74 float32 = 1e10
                        _ = v74
                    }
                    for fi75 := 0; fi75 < 10; fi75++ {
                        switch fi75 {
                        case 0:
                            v3 = false
                        }
                        v76 := GenS15{F0: (v3 && v3), F1: (fi43 + fi43)}
                        v5 = (126 > -1)
                        v77 := GenS16{F0: 42, F1: 786}
                        v78 := GenS1{F0: (v1 ^ (v1 << ((2) & 63))), F1: 42}
                        _ = fi75
                        _ = v76
                        _ = v77
                        _ = v78
                    }
                }
                if len(v4) > 0 {
                    v4[0] = 449
                }
            }
            switch 0 {
            case 0:
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            case -9223372036854775808:
                if len(v4) > 0 {
                    v4[0] = 4294967295
                }
            case 4294967295:
                for fi79 := 0; fi79 < 1; fi79++ {
                    var v80 string = ("\n\t" + ("\n\t" + (("\n\t" + "\n\t") + "a")))
                    if ((v36 / float64(((-2.25 / 1e10) * 0.1))) >= v36) {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        v2 -= 128
                        v80 = v80
                    }
                    _ = fi79
                    _ = v80
                }
            }
            v81 := GenS1{F0: v1, F1: 921}
            _ = v81
        }
        v82 := GenS4{F0: 0, F1: -106}
        _ = fi43
        _ = v82
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v83 string = ("hello" + "hello")
    fmt.Println(v2, v83)
    if v5 {
        if (!(((true && true) && v3) && false)) {
            v4 = append(v4, 42)
            _ = len(v4)
            var v84 uint8 = (v2 - v2)
            _ = v84
        }
        v85 := GenS13{F0: v2}
        _ = v85
    } else {
        v1 += -996
        v86 := GenS17{F0: 42, F1: v1, F2: v2}
        for fi87 := 0; fi87 < 4; fi87++ {
            v1 += (v1 / ((810) | 1))
            var v88 string = v83
            _ = fi87
            _ = v88
        }
        _ = v86
    }
    v3 = (2 >= 4294967295)
    v36 = ((v36 + (v36 * (v36 * v36))) * float64(0.1))
    var v89 float64 = float64(-2.25)
    switch v83 {
    case "":
        v89 += (-2.25 * (1.5 + v36))
    case "hello":
        v36 = ((v89 / -2.25) * ((v89 - (v89 - (v89 / (v36 / float64(float32(v36)))))) / (v89 + v36)))
    default:
        var v90 uint64 = 1
        _ = v90
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v83:", v83)
    fmt.Println("v89:", v89)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v83
    _ = v89
}
