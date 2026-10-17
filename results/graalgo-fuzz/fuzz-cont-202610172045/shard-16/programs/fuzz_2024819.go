package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8 }
type GenS2 struct { F0 string; F1 string }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int16; F1 int32 }
type GenS7 struct { F0 uint8; F1 int16 }
type GenS8 struct { F0 uint8; F1 int8 }
type GenS9 struct { F0 uint16; F1 uint; F2 int16 }
type GenS10 struct { F0 string; F1 int64 }
type GenS11 struct { F0 uint16 }
type GenS12 struct { F0 uint32; F1 int16 }
type GenS13 struct { F0 int32 }
type GenS14 struct { F0 string }
type GenS15 struct { F0 uint8; F1 int64; F2 int8 }
type GenS16 struct { F0 bool }
type GenS17 struct { F0 int8 }
type GenS18 struct { F0 uint32; F1 uint32; F2 int32 }
type GenS19 struct { F0 uint16 }
type GenS20 struct { F0 uint8; F1 uint16; F2 int32 }
type GenS21 struct { F0 uint32; F1 int8; F2 uint32 }
type GenS22 struct { F0 uint; F1 uint; F2 bool }
type GenS23 struct { F0 string; F1 int16 }
type GenS24 struct { F0 int32; F1 int16 }

func main() {
    var v1 bool = false
    var v2 int16 = 1000
    v3 := []uint32{65535}
    v4 := make(map[int]int)
    if v1 {
        v5 := GenS1{F0: 42, F1: 0}
        _ = v5
    }
    var v6 string = "hello"
    if v1 {
        var v7 uint = 42
        if (v7 > (v7 | v7)) {
            var v8 uint8 = 128
            var v9 int32 = 2147483647
            _ = v8
            _ = v9
        }
        v7 = 18446744073709551615
        _ = v7
    } else {
        switch 42 {
        case -566:
            for fi10 := 0; fi10 < 20; fi10++ {
                v11 := GenS1{F0: fi10, F1: 128}
                if (!(!false)) {
                    v12 := GenS1{F0: (fi10 - 32768), F1: 42}
                    switch (fi10 - 875) {
                    case 32768:
                        v6 = ""
                    case 32767:
                        if v1 {
                            for fi13 := 0; fi13 < 9; fi13++ {
                                v14 := GenS2{F0: "hello", F1: v6}
                                fmt.Println(v6, v6)
                                _ = fi13
                                _ = v14
                            }
                        } else {
                            v4[(fi10 >> ((149) & 63))] = fi10
                            v15 := GenS2{F0: (("\n\t" + "") + ""), F1: "foo bar"}
                            _ = v15
                        }
                    }
                    for fi16 := 0; fi16 < 13; fi16++ {
                        v2 -= v2
                        _ = fi16
                    }
                    _ = v12
                }
                v2 += v2
                _ = fi10
                _ = v11
            }
        case 42:
            v6 = ""
        }
        for fi17 := 0; fi17 < 4; fi17++ {
            v2 += (v2 | v2)
            _ = fi17
        }
    }
    var v18 int64 = 9223372036854775807
    for fi19 := 0; fi19 < 5; fi19++ {
        switch -1 {
        case -181:
            if true {
                if true {
                    v20 := GenS2{F0: (v6 + v6), F1: v6}
                    for fi21 := 0; fi21 < 18; fi21++ {
                        switch (fi21 - 1) {
                        case 0:
                            v18 += v18
                        case -751:
                            _ = len(v4)
                        default:
                            if v1 {
                                fmt.Println(fi21, v18, fi21)
                                var v22 int64 = (^(v18))
                                if v1 {
                                    switch v6 {
                                    case "foo bar":
                                        var v23 int16 = (v2 * (v2 + (v2 << ((988) & 63))))
                                        v22 = (v22 * (^(v18)))
                                        fmt.Println(v2)
                                        v18 = (v22 % (((-(v22))) | 1))
                                        if false {
                                            if (0 != 2) {
                                                var v24 float64 = ((3.14159 - 1.5) - 1.5)
                                                v25 := GenS3{F0: 289}
                                                _ = v24
                                                _ = v25
                                            }
                                            if ("" >= "\n\t") {
                                                fmt.Println(v6, v18)
                                                fmt.Println(fi19, v22)
                                                var v26 uint64 = 42
                                                var v27 int64 = (v22 >> ((407) & 63))
                                                _ = v26
                                                _ = v27
                                            }
                                        } else {
                                            var v28 bool = (!false)
                                            _ = v28
                                        }
                                        _ = v23
                                    default:
                                        if (734 < 4294967295) {
                                            v18 += -518
                                        }
                                    }
                                } else {
                                    v29 := GenS4{F0: 42}
                                    _ = v29
                                }
                                v3 = append(v3, 65535)
                                var v30 float32 = 1.5
                                _ = v22
                                _ = v30
                            } else {
                                if ("\n\t" < (v6 + ((("\n\t" + (v6 + ("" + "\n\t"))) + v6) + v6))) {
                                    switch v6 {
                                    case "a":
                                        fmt.Println(v6)
                                    default:
                                        v1 = (255 >= 42)
                                    }
                                    var v31 uint32 = 0
                                    _ = v31
                                } else {
                                    var v32 uint64 = 42
                                    _ = len(v4)
                                    var v33 uint16 = 52
                                    if v1 {
                                        v1 = true
                                    } else {
                                        if v1 {
                                            var v34 bool = v1
                                            _ = v34
                                        } else {
                                            v35 := GenS1{F0: (fi19 | (fi19 + fi21)), F1: 0}
                                            fmt.Println(v32, v33)
                                            fmt.Println(v32, v6)
                                            _ = v35
                                        }
                                        switch -9223372036854775808 {
                                        case 127:
                                            if (!(2 <= 255)) {
                                                switch v6 {
                                                case "hello":
                                                    var v36 string = "a"
                                                    _ = v36
                                                default:
                                                    v32 += (v32 / ((v32) | 1))
                                                }
                                                v37 := GenS5{F0: (v33 | v33)}
                                                _ = v37
                                            } else {
                                                var v38 uint8 = 1
                                                fmt.Println(fi19, v38, fi21)
                                                var v39 int64 = (^(v18))
                                                _ = v38
                                                _ = v39
                                            }
                                        case 4294967295:
                                            if v1 {
                                                if (((fi19 | fi19) <= 128) && v1) {
                                                    v18 += v18
                                                    if (-76 <= -128) {
                                                        v40 := GenS6{F0: (v2 >> ((0) & 63)), F1: 1}
                                                        _ = v40
                                                    }
                                                    fmt.Println(v6)
                                                } else {
                                                    _ = len(v4)
                                                    fmt.Println(fi21, fi19, v18)
                                                }
                                                if true {
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                    v18 -= -9223372036854775808
                                                    v6 = ("\n\t" + v6)
                                                    fmt.Println(fi21)
                                                }
                                                var v41 string = v6
                                                v6 = v6
                                                if v1 {
                                                    v42 := GenS7{F0: 127, F1: v2}
                                                    _ = v42
                                                } else {
                                                    if (255 < 2) {
                                                        switch 32767 {
                                                        case 713:
                                                            if v1 {
                                                                v43 := GenS5{F0: (^(v33))}
                                                                _ = v43
                                                            }
                                                        case 9223372036854775807:
                                                            delete(v4, fi19)
                                                        case 585:
                                                            v44 := GenS8{F0: 248, F1: 86}
                                                            _ = v44
                                                        default:
                                                            if true {
                                                                v3 = append(v3, 4294967295)
                                                                if v1 {
                                                                    v45 := GenS1{F0: 816, F1: 104}
                                                                    if len(v3) > 0 {
                                                                        v3[0] = 42
                                                                    }
                                                                    switch v6 {
                                                                    case "":
                                                                        switch fi19 {
                                                                        case -1:
                                                                            v33 = v33
                                                                        case 42:
                                                                            switch (v41 + "a") {
                                                                            case "hello":
                                                                                fmt.Println(v18, v18)
                                                                            case "\n\t":
                                                                                switch v6 {
                                                                                case "\n\t":
                                                                                    _ = len(v3)
                                                                                case "foo bar":
                                                                                    v41 = ""
                                                                                case "a":
                                                                                    if v1 {
                                                                                        if v1 {
                                                                                            var v46 int32 = 0
                                                                                            _ = v46
                                                                                        }
                                                                                        switch 1 {
                                                                                        case 0:
                                                                                            if true {
                                                                                                if ((1.5 - ((float64(0.0) + -2.25) + (0.0 * 0.1))) == 0.0) {
                                                                                                    v41 = v41
                                                                                                } else {
                                                                                                    v41 = (v41 + v6)
                                                                                                    if (false || false) {
                                                                                                        if (4294967295 > 451) {
                                                                                                            v32 -= (v32 >> ((0) & 63))
                                                                                                            if v1 {
                                                                                                                var v47 int = (fi19 ^ -9223372036854775808)
                                                                                                                fmt.Println(fi19)
                                                                                                                _ = v47
                                                                                                            } else {
                                                                                                                if v1 {
                                                                                                                    v18 -= (v18 ^ v18)
                                                                                                                }
                                                                                                                if len(v3) > 0 {
                                                                                                                    v3[0] = 2147483648
                                                                                                                }
                                                                                                                v6 = v41
                                                                                                            }
                                                                                                        }
                                                                                                        var v48 uint8 = 127
                                                                                                        var v49 string = v41
                                                                                                        _ = v48
                                                                                                        _ = v49
                                                                                                    }
                                                                                                    v1 = (9223372036854775807 <= 42)
                                                                                                }
                                                                                                if v1 {
                                                                                                    if v1 {
                                                                                                        v50 := v4[fi21]
                                                                                                        var v51 int64 = int64(v2)
                                                                                                        _ = v50
                                                                                                        _ = v51
                                                                                                    } else {
                                                                                                        if ((-811 < (v18 & 0)) || v1) {
                                                                                                            var v52 int32 = 0
                                                                                                            if (((1.5 / (1e10 * -2.25)) * 1.5) < 0.0) {
                                                                                                                var v53 float64 = (0.1 * float64(1e10))
                                                                                                                _ = v53
                                                                                                            } else {
                                                                                                                v54 := GenS6{F0: (v2 / ((v2) | 1)), F1: 0}
                                                                                                                v55 := GenS4{F0: v33}
                                                                                                                v2 = (v2 >> ((665) & 63))
                                                                                                                v32 -= v32
                                                                                                                v52 = int32(v2)
                                                                                                                _ = v54
                                                                                                                _ = v55
                                                                                                            }
                                                                                                            _ = v52
                                                                                                        } else {
                                                                                                            switch (fi19 - fi21) {
                                                                                                            case 2147483647:
                                                                                                                v18 = 9223372036854775807
                                                                                                            case -575:
                                                                                                                v56 := GenS2{F0: v41, F1: "foo bar"}
                                                                                                                _ = v56
                                                                                                            case 32767:
                                                                                                                _ = len(v3)
                                                                                                            default:
                                                                                                                v57 := GenS2{F0: "a", F1: "hello"}
                                                                                                                _ = v57
                                                                                                            }
                                                                                                            if v1 {
                                                                                                                var v58 int = 420
                                                                                                                v1 = (!(2147483647 > 918))
                                                                                                                v59 := GenS1{F0: (fi21 - (fi19 | (fi21 - fi21))), F1: 1}
                                                                                                                _ = v58
                                                                                                                _ = v59
                                                                                                            }
                                                                                                        }
                                                                                                        if (0 > 128) {
                                                                                                            v3 = append(v3, 2)
                                                                                                            v6 = "\n\t"
                                                                                                            fmt.Println(v41, fi21)
                                                                                                        } else {
                                                                                                            delete(v4, 128)
                                                                                                        }
                                                                                                        delete(v4, -387)
                                                                                                    }
                                                                                                    switch v41 {
                                                                                                    case "hello":
                                                                                                        var v60 string = v6
                                                                                                        _ = v60
                                                                                                    case "\n\t":
                                                                                                        var v61 int16 = v2
                                                                                                        _ = v61
                                                                                                    case "foo bar":
                                                                                                        var v62 bool = v1
                                                                                                        _ = v62
                                                                                                    default:
                                                                                                        v4[(fi21 >> ((65535) & 63))] = (fi21 >> ((18446744073709551615) & 63))
                                                                                                    }
                                                                                                }
                                                                                            } else {
                                                                                                var v63 uint64 = uint64(v33)
                                                                                                _ = v63
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                    switch fi19 {
                                                                                    case 4294967295:
                                                                                        if v1 {
                                                                                            delete(v4, fi21)
                                                                                        } else {
                                                                                            switch v41 {
                                                                                            case "":
                                                                                                switch v41 {
                                                                                                case "\n\t":
                                                                                                    v64 := GenS1{F0: -1, F1: 42}
                                                                                                    _ = v64
                                                                                                case "":
                                                                                                    var v65 uint8 = 127
                                                                                                    var v66 uint16 = v33
                                                                                                    _ = v65
                                                                                                    _ = v66
                                                                                                }
                                                                                            case "a":
                                                                                                switch ((v41 + ("\n\t" + v6)) + v41) {
                                                                                                case "\n\t":
                                                                                                    if (v1 && (1 >= 1)) {
                                                                                                        if (65535 < uint32(v33)) {
                                                                                                            var v67 int16 = 0
                                                                                                            switch ((v6 + (v41 + (v41 + ""))) + ((v6 + "") + v6)) {
                                                                                                            case "":
                                                                                                                v41 = v6
                                                                                                            case "a":
                                                                                                                fmt.Println(v67, v67, fi19)
                                                                                                            }
                                                                                                            _ = v67
                                                                                                        } else {
                                                                                                            var v68 float32 = 0.1
                                                                                                            var v69 float64 = 1.5
                                                                                                            _ = v68
                                                                                                            _ = v69
                                                                                                        }
                                                                                                        if (1 >= 0) {
                                                                                                            var v70 uint = 80
                                                                                                            v3 = append(v3, 2147483647)
                                                                                                            _ = v70
                                                                                                        } else {
                                                                                                            v71 := GenS1{F0: -9223372036854775808, F1: 127}
                                                                                                            _ = v71
                                                                                                        }
                                                                                                    }
                                                                                                case "hello":
                                                                                                    if len(v3) > 0 {
                                                                                                        v3[0] = 65535
                                                                                                    }
                                                                                                    v2 += v2
                                                                                                    var v72 int = 127
                                                                                                    _ = v72
                                                                                                case "":
                                                                                                    switch -1 {
                                                                                                    case 0:
                                                                                                        v3 = append(v3, 2)
                                                                                                    case 128:
                                                                                                        var v73 uint32 = 70
                                                                                                        _ = v73
                                                                                                    default:
                                                                                                        v74 := GenS9{F0: 873, F1: 137, F2: 84}
                                                                                                        _ = v74
                                                                                                    }
                                                                                                default:
                                                                                                    v3 = append(v3, 42)
                                                                                                }
                                                                                            case "foo bar":
                                                                                                v6 = ""
                                                                                            }
                                                                                            v75 := GenS10{F0: v41, F1: 598}
                                                                                            if (!(false || false)) {
                                                                                                var v76 int8 = -128
                                                                                                var v77 int = fi21
                                                                                                _ = len(v4)
                                                                                                _ = v76
                                                                                                _ = v77
                                                                                            }
                                                                                            _ = v75
                                                                                        }
                                                                                    case -117:
                                                                                        v1 = v1
                                                                                        if v1 {
                                                                                            var v78 string = ((("foo bar" + v41) + v41) + "a")
                                                                                            var v79 uint8 = 2
                                                                                            v80 := GenS11{F0: v33}
                                                                                            _ = v78
                                                                                            _ = v79
                                                                                            _ = v80
                                                                                        } else {
                                                                                            v1 = false
                                                                                        }
                                                                                    case 1:
                                                                                        fmt.Println(fi21, v2)
                                                                                    default:
                                                                                        if v1 {
                                                                                            _ = len(v3)
                                                                                            if len(v3) > 0 {
                                                                                                v3[0] = 1
                                                                                            }
                                                                                            if ((fi21 & (^(fi21))) >= (-(fi21))) {
                                                                                                fmt.Println(v33)
                                                                                                if false {
                                                                                                    if (43 < 127) {
                                                                                                        if v1 {
                                                                                                            if v1 {
                                                                                                                var v81 int8 = 42
                                                                                                                var v82 int = (fi19 % ((fi19) | 1))
                                                                                                                _ = v81
                                                                                                                _ = v82
                                                                                                            }
                                                                                                            v2 = (v2 & v2)
                                                                                                        }
                                                                                                        if (987 == 42) {
                                                                                                            v18 = int64(v2)
                                                                                                        }
                                                                                                    }
                                                                                                    v3 = append(v3, uint32(v33))
                                                                                                    if (!(v2 >= v2)) {
                                                                                                        _ = len(v3)
                                                                                                        if ((v18 << ((4294967295) & 63)) == 0) {
                                                                                                            if false {
                                                                                                                fmt.Println(fi19)
                                                                                                            }
                                                                                                            if (-114 < int8(v2)) {
                                                                                                                switch ("hello" + ("hello" + v6)) {
                                                                                                                case "\n\t":
                                                                                                                    v41 = v6
                                                                                                                case "foo bar":
                                                                                                                    var v83 int64 = v18
                                                                                                                    _ = v83
                                                                                                                }
                                                                                                                v84 := GenS12{F0: 2147483648, F1: (v2 >> ((18446744073709551615) & 63))}
                                                                                                                fmt.Println(fi19)
                                                                                                                if ((v41 + v6) == "\n\t") {
                                                                                                                    var v85 int8 = -1
                                                                                                                    if (!false) {
                                                                                                                        v41 = ("\n\t" + "")
                                                                                                                    }
                                                                                                                    _ = v85
                                                                                                                }
                                                                                                                fmt.Println(v32, v32, fi19)
                                                                                                                v86 := GenS3{F0: (v33 * v33)}
                                                                                                                _ = v84
                                                                                                                _ = v86
                                                                                                            } else {
                                                                                                                switch (fi19 - (fi21 * fi19)) {
                                                                                                                case 127:
                                                                                                                    v87 := GenS6{F0: v2, F1: -1}
                                                                                                                    _ = v87
                                                                                                                case 32767:
                                                                                                                    var v88 int = (fi21 >> ((1) & 63))
                                                                                                                    _ = v88
                                                                                                                }
                                                                                                                if ((v2 != 0) && false) {
                                                                                                                    var v89 float64 = 3.14159
                                                                                                                    v3 = append(v3, 1)
                                                                                                                    var v90 uint8 = 1
                                                                                                                    var v91 int8 = 42
                                                                                                                    _ = v89
                                                                                                                    _ = v90
                                                                                                                    _ = v91
                                                                                                                } else {
                                                                                                                    v33 += 649
                                                                                                                }
                                                                                                            }
                                                                                                            fmt.Println(v2, fi19, v33)
                                                                                                        }
                                                                                                        fmt.Println(v6, v6)
                                                                                                    } else {
                                                                                                        var v92 bool = (v6 >= v41)
                                                                                                        v41 = ("foo bar" + (v6 + ("a" + (v6 + v6))))
                                                                                                        _ = v92
                                                                                                    }
                                                                                                } else {
                                                                                                    v18 -= (v18 / (((v18 / ((v18) | 1))) | 1))
                                                                                                    v4[fi21] = 128
                                                                                                    var v93 int16 = v2
                                                                                                    _ = v93
                                                                                                }
                                                                                                switch "foo bar" {
                                                                                                case "a":
                                                                                                    var v94 uint16 = v33
                                                                                                    _ = v94
                                                                                                case "foo bar":
                                                                                                    var v95 float32 = 3.14159
                                                                                                    _ = v95
                                                                                                default:
                                                                                                    fmt.Println(v41)
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                    delete(v4, fi21)
                                                                                }
                                                                            }
                                                                        case -9223372036854775808:
                                                                            if v1 {
                                                                                v18 += v18
                                                                            } else {
                                                                                var v96 int16 = (-(v2))
                                                                                if true {
                                                                                    _ = len(v3)
                                                                                } else {
                                                                                    fmt.Println(v1, fi21)
                                                                                    var v97 uint8 = 161
                                                                                    _ = v97
                                                                                }
                                                                                delete(v4, (fi21 / ((fi21) | 1)))
                                                                                _ = v96
                                                                            }
                                                                        default:
                                                                            if true {
                                                                                if (65535 > 2) {
                                                                                    if (2 < 1) {
                                                                                        if v1 {
                                                                                            if (!(32767 > (fi21 * fi21))) {
                                                                                                var v98 uint16 = 631
                                                                                                v41 = "\n\t"
                                                                                                v18 = (v18 - (-(v18)))
                                                                                                _ = v98
                                                                                            }
                                                                                            switch (fi21 % ((32768) | 1)) {
                                                                                            case -158:
                                                                                                var v99 int32 = 232
                                                                                                _ = v99
                                                                                            case 1:
                                                                                                v32 -= uint64(v33)
                                                                                            default:
                                                                                                v41 = ("a" + v6)
                                                                                            }
                                                                                            v2 -= (v2 & (v2 | (v2 / ((v2) | 1))))
                                                                                        } else {
                                                                                            v1 = (2 >= 18446744073709551615)
                                                                                            var v100 uint = 2
                                                                                            _ = v100
                                                                                        }
                                                                                        if len(v3) > 0 {
                                                                                            v3[0] = 1
                                                                                        }
                                                                                        if v1 {
                                                                                            v1 = (-128 >= -1)
                                                                                        } else {
                                                                                            switch fi21 {
                                                                                            case 402:
                                                                                                if (!true) {
                                                                                                    v2 = -32768
                                                                                                    if (false && v1) {
                                                                                                        var v101 uint32 = 0
                                                                                                        fmt.Println(v6, v32, v2)
                                                                                                        v32 += v32
                                                                                                        if v1 {
                                                                                                            if v1 {
                                                                                                                if len(v3) >= 2 {
                                                                                                                    v3 = v3[1:len(v3)]
                                                                                                                }
                                                                                                            } else {
                                                                                                                v41 = "foo bar"
                                                                                                            }
                                                                                                            v101 += v101
                                                                                                            if (((v6 + "foo bar") >= (v6 + v6)) && false) {
                                                                                                                v41 = "hello"
                                                                                                                v102 := GenS13{F0: 2147483647}
                                                                                                                switch int(v33) {
                                                                                                                case 818:
                                                                                                                    switch int(v101) {
                                                                                                                    case -1:
                                                                                                                        if (!false) {
                                                                                                                            v32 = v32
                                                                                                                            switch v6 {
                                                                                                                            case "hello":
                                                                                                                                var v103 int16 = v2
                                                                                                                                _ = v103
                                                                                                                            default:
                                                                                                                                v18 = v18
                                                                                                                                v104 := GenS12{F0: 2, F1: v2}
                                                                                                                                if true {
                                                                                                                                    v105 := v4[(fi21 ^ (fi21 | -2147483648))]
                                                                                                                                    _ = v105
                                                                                                                                }
                                                                                                                                _ = v104
                                                                                                                            }
                                                                                                                        }
                                                                                                                    case 1:
                                                                                                                        v18 = v18
                                                                                                                    }
                                                                                                                case 32768:
                                                                                                                    v33 = v33
                                                                                                                case 127:
                                                                                                                    v2 += v2
                                                                                                                }
                                                                                                                if false {
                                                                                                                    v106 := GenS9{F0: v33, F1: 0, F2: (v2 & v2)}
                                                                                                                    var v107 int8 = -78
                                                                                                                    v3 = append(v3, v101)
                                                                                                                    v18 += 977
                                                                                                                    v108 := GenS3{F0: 2}
                                                                                                                    _ = v106
                                                                                                                    _ = v107
                                                                                                                    _ = v108
                                                                                                                }
                                                                                                                var v109 bool = (!(22 == 0))
                                                                                                                _ = v102
                                                                                                                _ = v109
                                                                                                            }
                                                                                                        }
                                                                                                        var v110 int8 = 0
                                                                                                        v18 -= (-(v18))
                                                                                                        _ = v101
                                                                                                        _ = v110
                                                                                                    }
                                                                                                    var v111 int16 = v2
                                                                                                    _ = v111
                                                                                                } else {
                                                                                                    switch v6 {
                                                                                                    case "a":
                                                                                                        fmt.Println(v2, v6)
                                                                                                    }
                                                                                                    if v1 {
                                                                                                        v1 = v1
                                                                                                        var v112 string = ("hello" + v41)
                                                                                                        _ = v112
                                                                                                    }
                                                                                                }
                                                                                            case 94:
                                                                                                var v113 bool = v1
                                                                                                _ = v113
                                                                                            case 468:
                                                                                                if v1 {
                                                                                                    v33 += 42
                                                                                                    var v114 int = -891
                                                                                                    _ = v114
                                                                                                } else {
                                                                                                    v4[int(v2)] = fi21
                                                                                                    if v1 {
                                                                                                        fmt.Println(v2, v33)
                                                                                                        var v115 int64 = v18
                                                                                                        var v116 bool = false
                                                                                                        _ = v115
                                                                                                        _ = v116
                                                                                                    } else {
                                                                                                        if ((fi19 < 32768) && (uint16(v32) == (v33 * (v33 ^ v33)))) {
                                                                                                            var v117 uint8 = 127
                                                                                                            v118 := GenS4{F0: 614}
                                                                                                            _ = v117
                                                                                                            _ = v118
                                                                                                        } else {
                                                                                                            v41 = v6
                                                                                                        }
                                                                                                    }
                                                                                                    fmt.Println(v2, v41, fi21)
                                                                                                }
                                                                                            default:
                                                                                                var v119 uint8 = 1
                                                                                                _ = v119
                                                                                            }
                                                                                            switch fi21 {
                                                                                            case 1:
                                                                                                switch (fi19 - (^(fi21))) {
                                                                                                case 0:
                                                                                                    switch 32767 {
                                                                                                    case -1:
                                                                                                        var v120 uint32 = 2
                                                                                                        _ = v120
                                                                                                    case 32767:
                                                                                                        if false {
                                                                                                            v121 := GenS2{F0: "hello", F1: (v41 + "a")}
                                                                                                            v122 := GenS14{F0: v41}
                                                                                                            var v123 bool = true
                                                                                                            _ = v121
                                                                                                            _ = v122
                                                                                                            _ = v123
                                                                                                        }
                                                                                                    }
                                                                                                    var v124 float64 = 1.5
                                                                                                    _ = v124
                                                                                                case 127:
                                                                                                    v4[fi19] = (fi19 * (fi19 / ((fi21) | 1)))
                                                                                                }
                                                                                            default:
                                                                                                if (!(v1 && (v41 != ""))) {
                                                                                                    switch (fi21 >> ((42) & 63)) {
                                                                                                    case 4294967295:
                                                                                                        switch ((v6 + v41) + v6) {
                                                                                                        case "a":
                                                                                                            if ((631 != 65535) || (432 >= 628)) {
                                                                                                                _ = len(v4)
                                                                                                                fmt.Println(v2, v32)
                                                                                                                v125 := GenS15{F0: 41, F1: -389, F2: 42}
                                                                                                                _ = v125
                                                                                                            } else {
                                                                                                                fmt.Println(v2, v18)
                                                                                                                v126 := GenS3{F0: (v33 / (((v33 | (v33 % ((v33) | 1)))) | 1))}
                                                                                                                var v127 int32 = 500
                                                                                                                _ = v126
                                                                                                                _ = v127
                                                                                                            }
                                                                                                        case "":
                                                                                                            v33 = (v33 >> ((827) & 63))
                                                                                                        default:
                                                                                                            var v128 uint8 = 42
                                                                                                            _ = v128
                                                                                                        }
                                                                                                    }
                                                                                                    fmt.Println(v32)
                                                                                                }
                                                                                                var v129 uint8 = 127
                                                                                                if (((^(v2)) != -318) && v1) {
                                                                                                    v130 := GenS4{F0: v33}
                                                                                                    _ = v130
                                                                                                }
                                                                                                _ = v129
                                                                                            }
                                                                                            if true {
                                                                                                fmt.Println(v6)
                                                                                                fmt.Println(v41, v32, v41)
                                                                                                var v131 uint32 = 121
                                                                                                _ = len(v3)
                                                                                                var v132 uint16 = v33
                                                                                                _ = v131
                                                                                                _ = v132
                                                                                            } else {
                                                                                                var v133 float64 = float64(((3.14159 / float32((1.5 - float64(0.0)))) - 1.5))
                                                                                                var v134 int = (fi21 % (((fi19 | fi19)) | 1))
                                                                                                var v135 uint64 = uint64(v33)
                                                                                                if v1 {
                                                                                                    var v136 bool = false
                                                                                                    switch (v41 + ("foo bar" + v41)) {
                                                                                                    case "hello":
                                                                                                        v134 = 127
                                                                                                    case "\n\t":
                                                                                                        if (42 < v18) {
                                                                                                            v136 = (-2147483648 < -2147483648)
                                                                                                            v1 = (0.1 < float64((-2.25 / 1.5)))
                                                                                                            var v137 uint16 = (v33 + (v33 - v33))
                                                                                                            _ = v137
                                                                                                        } else {
                                                                                                            if ((v133 * (v133 + 3.14159)) <= v133) {
                                                                                                                v41 = v6
                                                                                                                var v138 uint8 = 128
                                                                                                                if true {
                                                                                                                    switch (fi19 | -915) {
                                                                                                                    case 0:
                                                                                                                        v4[127] = -1
                                                                                                                    }
                                                                                                                    _ = len(v4)
                                                                                                                }
                                                                                                                _ = v138
                                                                                                            }
                                                                                                        }
                                                                                                    default:
                                                                                                        v18 = 42
                                                                                                    }
                                                                                                    v33 = 0
                                                                                                    var v139 uint = 2
                                                                                                    if len(v3) > 0 {
                                                                                                        v3[0] = 65535
                                                                                                    }
                                                                                                    v140 := GenS15{F0: 0, F1: v18, F2: 18}
                                                                                                    _ = v136
                                                                                                    _ = v139
                                                                                                    _ = v140
                                                                                                }
                                                                                                if len(v3) > 0 {
                                                                                                    v3[0] = 65535
                                                                                                }
                                                                                                _ = v133
                                                                                                _ = v134
                                                                                                _ = v135
                                                                                            }
                                                                                            var v141 int64 = v18
                                                                                            _ = v141
                                                                                        }
                                                                                    } else {
                                                                                        switch (v6 + "hello") {
                                                                                        case "":
                                                                                            if ((v2 < v2) && (v1 || (v1 && false))) {
                                                                                                var v142 float64 = 1e10
                                                                                                _ = v142
                                                                                            } else {
                                                                                                v33 = v33
                                                                                                v32 += 0
                                                                                                fmt.Println(v41, v33, fi19)
                                                                                            }
                                                                                        case "\n\t":
                                                                                            fmt.Println(v33, v33)
                                                                                        case "foo bar":
                                                                                            _ = len(v4)
                                                                                        }
                                                                                    }
                                                                                } else {
                                                                                    var v143 uint32 = 248
                                                                                    _ = v143
                                                                                }
                                                                                _ = len(v3)
                                                                                switch (^(fi19)) {
                                                                                case 1:
                                                                                    v144 := GenS1{F0: (fi21 << ((18446744073709551615) & 63)), F1: 42}
                                                                                    _ = v144
                                                                                case -9223372036854775808:
                                                                                    if (804 <= 2) {
                                                                                        v6 = v6
                                                                                        fmt.Println(v2)
                                                                                    } else {
                                                                                        if (!(0 == -133)) {
                                                                                            v18 = int64(v33)
                                                                                            fmt.Println(v32, v33)
                                                                                        }
                                                                                        var v145 uint32 = 4294967295
                                                                                        _ = v145
                                                                                    }
                                                                                case -1:
                                                                                    if v1 {
                                                                                        var v146 int8 = 0
                                                                                        var v147 uint16 = (v33 / ((v33) | 1))
                                                                                        if len(v3) > 0 {
                                                                                            v3[0] = 829
                                                                                        }
                                                                                        _ = v146
                                                                                        _ = v147
                                                                                    } else {
                                                                                        if ((!(853 < 65535)) && false) {
                                                                                            var v148 float32 = 1e10
                                                                                            _ = v148
                                                                                        }
                                                                                    }
                                                                                default:
                                                                                    v149 := GenS11{F0: 32767}
                                                                                    _ = v149
                                                                                }
                                                                            }
                                                                        }
                                                                    case "a":
                                                                        _ = len(v4)
                                                                        var v150 uint32 = 2147483648
                                                                        v41 = "hello"
                                                                        if len(v3) >= 2 {
                                                                            v3 = v3[1:len(v3)]
                                                                        }
                                                                        switch (fi21 / ((-97) | 1)) {
                                                                        case 665:
                                                                            _ = len(v3)
                                                                        case 4294967295:
                                                                            var v151 bool = (0 <= 201)
                                                                            _ = v151
                                                                        case 9223372036854775807:
                                                                            var v152 int = 128
                                                                            _ = v152
                                                                        default:
                                                                            switch ("" + v41) {
                                                                            case "hello":
                                                                                v41 = "foo bar"
                                                                            case "\n\t":
                                                                                v1 = (v18 != (v18 ^ (^(v18))))
                                                                            case "a":
                                                                                v153 := GenS4{F0: (^(v33))}
                                                                                _ = v153
                                                                            default:
                                                                                switch ((v6 + (v6 + "a")) + "\n\t") {
                                                                                case "":
                                                                                    v32 -= v32
                                                                                case "a":
                                                                                    fmt.Println(v18)
                                                                                case "foo bar":
                                                                                    if false {
                                                                                        var v154 float32 = 0.0
                                                                                        v6 = ("foo bar" + v6)
                                                                                        _ = v154
                                                                                    }
                                                                                }
                                                                            }
                                                                        }
                                                                        switch (fi21 >> ((136) & 63)) {
                                                                        case -494:
                                                                            switch v41 {
                                                                            case "\n\t":
                                                                                if ((fi19 > (fi21 * fi21)) && false) {
                                                                                    v155 := GenS9{F0: (^(v33)), F1: uint(v32), F2: (v2 | (v2 | int16(v18)))}
                                                                                    v1 = v1
                                                                                    v2 = 7
                                                                                    v156 := GenS16{F0: (255 >= 0)}
                                                                                    _ = v155
                                                                                    _ = v156
                                                                                } else {
                                                                                    v157 := GenS14{F0: "\n\t"}
                                                                                    if v1 {
                                                                                        switch (((v41 + v41) + v41) + (v6 + "")) {
                                                                                        case "":
                                                                                            switch v6 {
                                                                                            case "\n\t":
                                                                                                fmt.Println(v6)
                                                                                            }
                                                                                        }
                                                                                        var v158 float64 = 0.0
                                                                                        _ = v158
                                                                                    }
                                                                                    _ = v157
                                                                                }
                                                                            case "hello":
                                                                                if (65535 > 18446744073709551615) {
                                                                                    var v159 int32 = 42
                                                                                    v160 := GenS17{F0: 100}
                                                                                    _ = v159
                                                                                    _ = v160
                                                                                }
                                                                            default:
                                                                                v6 = (v41 + v6)
                                                                            }
                                                                        }
                                                                        if (v1 && (!(42 != 36))) {
                                                                            fmt.Println(fi21)
                                                                        }
                                                                        _ = v150
                                                                    case "hello":
                                                                        if (2147483647 <= 0) {
                                                                            var v161 int = (fi19 * -9223372036854775808)
                                                                            if v1 {
                                                                                if (1 < 18446744073709551615) {
                                                                                    fmt.Println(v2)
                                                                                    v18 -= (v18 << ((508) & 63))
                                                                                    if v1 {
                                                                                        v2 -= -601
                                                                                    } else {
                                                                                        var v162 string = ((v41 + v6) + "\n\t")
                                                                                        _ = v162
                                                                                    }
                                                                                    v2 = v2
                                                                                } else {
                                                                                    if v1 {
                                                                                        if true {
                                                                                            if (false || (((v33 % ((255) | 1)) == (v33 - v33)) && v1)) {
                                                                                                v4[(-(fi19))] = (-(fi21))
                                                                                            } else {
                                                                                                fmt.Println(v41)
                                                                                            }
                                                                                            v4[fi19] = fi19
                                                                                            v33 -= 200
                                                                                        }
                                                                                        v3 = append(v3, 825)
                                                                                    } else {
                                                                                        var v163 uint64 = (v32 >> ((4294967295) & 63))
                                                                                        if (v1 && v1) {
                                                                                            if (-68 < -128) {
                                                                                                if false {
                                                                                                    if len(v3) > 0 {
                                                                                                        v3[0] = 590
                                                                                                    }
                                                                                                }
                                                                                            } else {
                                                                                                var v164 uint8 = 255
                                                                                                _ = v164
                                                                                            }
                                                                                            v165 := GenS9{F0: 553, F1: 421, F2: v2}
                                                                                            var v166 uint = 42
                                                                                            _ = v165
                                                                                            _ = v166
                                                                                        }
                                                                                        fmt.Println(v33, v41, v1)
                                                                                        _ = v163
                                                                                    }
                                                                                }
                                                                            }
                                                                            v167 := GenS10{F0: v6, F1: v18}
                                                                            _ = v161
                                                                            _ = v167
                                                                        } else {
                                                                            var v168 uint = 255
                                                                            v168 -= 65535
                                                                            _ = v168
                                                                        }
                                                                    }
                                                                    _ = v45
                                                                } else {
                                                                    if len(v3) > 0 {
                                                                        v3[0] = 1
                                                                    }
                                                                }
                                                                fmt.Println(v6)
                                                            } else {
                                                                v33 = v33
                                                            }
                                                        }
                                                        v3 = append(v3, 4294967295)
                                                    }
                                                    switch 127 {
                                                    case 2147483647:
                                                        fmt.Println(v33, v32)
                                                    default:
                                                        if false {
                                                            switch "a" {
                                                            case "\n\t":
                                                                v3 = append(v3, 244)
                                                            case "":
                                                                if len(v3) >= 2 {
                                                                    v3 = v3[1:len(v3)]
                                                                }
                                                                var v169 uint64 = v32
                                                                _ = v169
                                                            default:
                                                                var v170 uint64 = (^(v32))
                                                                _ = v170
                                                            }
                                                        } else {
                                                            var v171 uint8 = 105
                                                            v172 := GenS9{F0: (v33 | v33), F1: 2, F2: (v2 | (v2 / ((v2) | 1)))}
                                                            var v173 uint = 1
                                                            var v174 int16 = 42
                                                            var v175 int8 = 100
                                                            _ = v171
                                                            _ = v172
                                                            _ = v173
                                                            _ = v174
                                                            _ = v175
                                                        }
                                                    }
                                                    v176 := GenS18{F0: 645, F1: 42, F2: 24}
                                                    v177 := GenS1{F0: (fi19 / (((fi21 * (-(fi19)))) | 1)), F1: 42}
                                                    v6 = "foo bar"
                                                    fmt.Println(v33)
                                                    _ = v176
                                                    _ = v177
                                                }
                                                _ = v41
                                            }
                                        }
                                        if v1 {
                                            fmt.Println(fi21)
                                            if ((^(v33)) != 2) {
                                                v178 := GenS19{F0: (v33 % ((v33) | 1))}
                                                var v179 uint = 4294967295
                                                _ = v178
                                                _ = v179
                                            }
                                            v180 := v4[(fi19 | fi21)]
                                            _ = v180
                                        }
                                    }
                                    _ = v32
                                    _ = v33
                                }
                                if v1 {
                                    var v181 int64 = (v18 + v18)
                                    v4[(fi21 * 128)] = fi21
                                    _ = v181
                                }
                                v1 = false
                            }
                        }
                        _ = fi21
                    }
                    if v1 {
                        if false {
                            if (2147483647 < 228) {
                                v18 = (v18 << ((0) & 63))
                                var v182 float64 = 1.5
                                v183 := GenS7{F0: 1, F1: int16(fi19)}
                                _ = v182
                                _ = v183
                            }
                            if false {
                                if (4294967295 == 678) {
                                    v184 := GenS9{F0: 0, F1: 42, F2: v2}
                                    if v1 {
                                        if v1 {
                                            var v185 uint16 = 32767
                                            if (v1 && (0 > 100)) {
                                                switch ((((("\n\t" + v6) + ((v6 + v6) + (v6 + (v6 + (v6 + v6))))) + v6) + (v6 + ("\n\t" + v6))) + "hello") {
                                                case "":
                                                    var v186 uint8 = 127
                                                    _ = v186
                                                }
                                                for fi187 := 0; fi187 < 7; fi187++ {
                                                    v188 := GenS20{F0: 136, F1: v185, F2: 931}
                                                    _ = fi187
                                                    _ = v188
                                                }
                                                var v189 string = v6
                                                _ = v189
                                            } else {
                                                var v190 uint16 = v185
                                                v185 -= (v190 | (v190 / ((42) | 1)))
                                                switch "hello" {
                                                case "hello":
                                                    v191 := GenS21{F0: 42, F1: -1, F2: 2147483648}
                                                    _ = v191
                                                }
                                                _ = v190
                                            }
                                            _ = v185
                                        } else {
                                            if ((true && (!true)) && v1) {
                                                v192 := GenS16{F0: v1}
                                                v193 := GenS6{F0: (v2 % ((v2) | 1)), F1: 37}
                                                _ = v192
                                                _ = v193
                                            } else {
                                                v2 = v2
                                                switch ("" + ("hello" + v6)) {
                                                case "\n\t":
                                                    if false {
                                                        v3 = append(v3, 65535)
                                                        v6 = ((v6 + "") + v6)
                                                    } else {
                                                        var v194 int8 = 100
                                                        if (!v1) {
                                                            var v195 uint32 = 2147483647
                                                            var v196 uint8 = 127
                                                            var v197 float32 = 0.1
                                                            _ = v195
                                                            _ = v196
                                                            _ = v197
                                                        } else {
                                                            for fi198 := 0; fi198 < 15; fi198++ {
                                                                if false {
                                                                    if (v2 >= v2) {
                                                                        _ = len(v3)
                                                                        if (fi198 < (fi19 << ((255) & 63))) {
                                                                            var v199 int = (fi198 << ((204) & 63))
                                                                            v194 = (v194 - (-(v194)))
                                                                            _ = v199
                                                                        } else {
                                                                            var v200 uint16 = 2
                                                                            _ = v200
                                                                        }
                                                                        var v201 int = fi19
                                                                        v2 = (v2 / ((v2) | 1))
                                                                        v202 := GenS22{F0: 66, F1: 698, F2: true}
                                                                        if (186 <= 42) {
                                                                            fmt.Println(fi198)
                                                                        } else {
                                                                            var v203 int = (-(fi198))
                                                                            var v204 bool = ((3.14159 * (0.1 - 0.0)) != 3.14159)
                                                                            fmt.Println(v2, v18, v201)
                                                                            var v205 uint16 = 42
                                                                            _ = v203
                                                                            _ = v204
                                                                            _ = v205
                                                                        }
                                                                        _ = v201
                                                                        _ = v202
                                                                    } else {
                                                                        if v1 {
                                                                            var v206 float32 = 1e10
                                                                            switch fi198 {
                                                                            case 71:
                                                                                var v207 int64 = (v18 / (((^(v18))) | 1))
                                                                                _ = v207
                                                                            case 32767:
                                                                                var v208 int8 = (v194 | -1)
                                                                                _ = v208
                                                                            case 42:
                                                                                switch fi19 {
                                                                                case -1:
                                                                                    var v209 uint32 = 530
                                                                                    _ = v209
                                                                                default:
                                                                                    var v210 int8 = v194
                                                                                    _ = v210
                                                                                }
                                                                            default:
                                                                                if v1 {
                                                                                    var v211 uint8 = 116
                                                                                    _ = v211
                                                                                }
                                                                            }
                                                                            fmt.Println(v18, v6)
                                                                            var v212 bool = (!v1)
                                                                            var v213 float64 = ((-2.25 / -2.25) / 1e10)
                                                                            _ = v206
                                                                            _ = v212
                                                                            _ = v213
                                                                        } else {
                                                                            if false {
                                                                                _ = len(v4)
                                                                            }
                                                                            var v214 bool = false
                                                                            switch (fi19 % ((fi198) | 1)) {
                                                                            case 95:
                                                                                var v215 int32 = 1
                                                                                _ = v215
                                                                            default:
                                                                                var v216 bool = v1
                                                                                _ = v216
                                                                            }
                                                                            _ = v214
                                                                        }
                                                                    }
                                                                }
                                                                _ = fi198
                                                            }
                                                            v217 := GenS2{F0: v6, F1: ((v6 + ("a" + ((("foo bar" + "hello") + "a") + v6))) + "hello")}
                                                            if ((int(v18) == (fi19 + fi19)) && (18446744073709551615 <= 255)) {
                                                                var v218 uint32 = 1
                                                                _ = v218
                                                            }
                                                            _ = v217
                                                        }
                                                        _ = v194
                                                    }
                                                case "a":
                                                    for fi219 := 0; fi219 < 11; fi219++ {
                                                        if len(v3) > 0 {
                                                            v3[0] = 2
                                                        }
                                                        _ = fi219
                                                    }
                                                }
                                                var v220 int8 = 127
                                                _ = v220
                                            }
                                            v4[fi19] = 740
                                        }
                                        v221 := GenS11{F0: 881}
                                        var v222 int64 = (v18 ^ v18)
                                        _ = v221
                                        _ = v222
                                    } else {
                                        for fi223 := 0; fi223 < 3; fi223++ {
                                            fmt.Println(v2, v6)
                                            switch v6 {
                                            case "foo bar":
                                                var v224 uint = 772
                                                _ = v224
                                            case "":
                                                fmt.Println(v18, v6)
                                            case "\n\t":
                                                v1 = (-2147483648 >= int32(v2))
                                            }
                                            v1 = true
                                            var v225 float32 = (1.5 - -2.25)
                                            _ = fi223
                                            _ = v225
                                        }
                                    }
                                    v1 = true
                                    _ = v184
                                } else {
                                    for fi226 := 0; fi226 < 7; fi226++ {
                                        v4[fi226] = (fi226 & (fi226 & (fi226 / ((fi226) | 1))))
                                        var v227 int = fi19
                                        _ = fi226
                                        _ = v227
                                    }
                                }
                                fmt.Println(fi19, v1)
                                v18 = (v18 | v18)
                            }
                            var v228 int32 = 1
                            _ = v228
                        }
                    } else {
                        var v229 uint32 = 2
                        if v1 {
                            var v230 uint64 = 306
                            var v231 string = v6
                            _ = v230
                            _ = v231
                        } else {
                            v2 -= 32767
                        }
                        for fi232 := 0; fi232 < 15; fi232++ {
                            v1 = v1
                            if v1 {
                                var v233 bool = false
                                v233 = true
                                v1 = (3.14159 >= (1e10 / (float32((-2.25 - 3.14159)) / 1.5)))
                                v234 := GenS6{F0: v2, F1: 0}
                                _ = v233
                                _ = v234
                            }
                            v1 = v1
                            fmt.Println(v1, v229, v229)
                            var v235 int = 9223372036854775807
                            _ = fi232
                            _ = v235
                        }
                        _ = v229
                    }
                    _ = v20
                }
            }
        case 42:
            delete(v4, (fi19 >> ((896) & 63)))
        default:
            v236 := GenS8{F0: 1, F1: -128}
            _ = v236
        }
        _ = fi19
    }
    var v237 int16 = (v2 << ((607) & 63))
    fmt.Println(v2, v2)
    v18 = (v18 & (v18 / (((v18 + v18)) | 1)))
    v238 := GenS23{F0: v6, F1: v2}
    var v239 int = 662
    for fi240 := 0; fi240 < 17; fi240++ {
        if true {
            if false {
                for fi241 := 0; fi241 < 11; fi241++ {
                    fmt.Println(fi240, v239)
                    _ = fi241
                }
                fmt.Println(fi240)
            }
            for fi242 := 0; fi242 < 7; fi242++ {
                v4[v239] = (-(fi240))
                switch "hello" {
                case "foo bar":
                    if len(v3) > 0 {
                        v3[0] = 473
                    }
                }
                if len(v3) > 0 {
                    v3[0] = 4294967295
                }
                _ = fi242
            }
            var v243 float64 = 1.5
            _ = v243
        } else {
            for fi244 := 0; fi244 < 16; fi244++ {
                v237 = v2
                fmt.Println(v6)
                _ = fi244
            }
            var v245 int16 = (v237 * 1000)
            _ = v245
        }
        var v246 int64 = 42
        _ = fi240
        _ = v246
    }
    v247 := GenS24{F0: 316, F1: v237}
    v3 = append(v3, 0)
    var v248 int64 = (-(v18))
    fmt.Println(v18, v239)
    switch (v239 / ((-9223372036854775808) | 1)) {
    case 127:
        v248 -= (v248 ^ v18)
    case -1:
        var v249 bool = (1 >= 325)
        _ = v249
    case -396:
        var v250 uint64 = 739
        _ = v250
    }
    v6 = v6
    fmt.Println(v6)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v237:", v237)
    fmt.Println("v239:", v239)
    fmt.Println("v248:", v248)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v237
    _ = v238
    _ = v239
    _ = v247
    _ = v248
    _ = v3
    _ = v4
    _ = v6
}
