package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int32; F1 uint32; F2 bool }
type GenS3 struct { F0 int16; F1 int16; F2 bool }
type GenS4 struct { F0 uint; F1 int32; F2 int32 }
type GenS5 struct { F0 uint32; F1 uint; F2 uint8 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 int16; F1 uint16; F2 uint16 }
type GenS8 struct { F0 int64; F1 int64 }
type GenS9 struct { F0 bool }

func main() {
    var v1 string = ("\n\t" + ("foo bar" + "hello"))
    var v2 int16 = -823
    v3 := []uint32{65535}
    var v4 uint16 = 806
    v5 := make(map[int]bool)
    var v6 float32 = -2.25
    v7 := GenS1{F0: true}
    if (2147483647 == -1) {
        switch v1 {
        case "foo bar":
            if false {
                var v8 bool = ("\n\t" <= ("" + ((v1 + "foo bar") + v1)))
                v9 := GenS2{F0: -1, F1: 4294967295, F2: (((v6 >= float32(0.1)) && true) && (!(v4 > v4)))}
                _ = v8
                _ = v9
            }
        case "hello":
            v10 := GenS2{F0: -576, F1: 826, F2: (true || (1e10 > (0.1 + (3.14159 / float64((3.14159 / 3.14159))))))}
            _ = v10
        case "a":
            v1 = ("\n\t" + "hello")
        }
    } else {
        v11 := GenS2{F0: -1, F1: 2147483647, F2: (0 >= 2)}
        switch ((v1 + v1) + v1) {
        case "":
            var v12 int16 = v2
            _ = v12
        case "foo bar":
            for fi13 := 0; fi13 < 13; fi13++ {
                v3 = append(v3, 2147483648)
                for fi14 := 0; fi14 < 12; fi14++ {
                    switch v1 {
                    case "a":
                        v3 = append(v3, 216)
                    case "foo bar":
                        if (100 >= -101) {
                            fmt.Println(v2)
                            v15 := GenS3{F0: (-(v2)), F1: 42, F2: false}
                            _ = v15
                        } else {
                            delete(v5, (-(fi14)))
                        }
                    }
                    _ = fi14
                }
                for fi16 := 0; fi16 < 5; fi16++ {
                    v1 = "a"
                    _ = fi16
                }
                _ = fi13
            }
        case "hello":
            var v17 float64 = 3.14159
            _ = v17
        default:
            v18 := GenS1{F0: true}
            _ = v18
        }
        for fi19 := 0; fi19 < 18; fi19++ {
            switch (fi19 - (fi19 / ((fi19) | 1))) {
            case -2147483648:
                fmt.Println(v2, v2, v1)
            case -1:
                switch ((v1 + v1) + v1) {
                case "foo bar":
                    if false {
                        var v20 uint = 354
                        _ = v20
                    } else {
                        v21 := GenS3{F0: v2, F1: -145, F2: false}
                        switch (fi19 % ((fi19) | 1)) {
                        case 9223372036854775807:
                            for fi22 := 0; fi22 < 19; fi22++ {
                                if false {
                                    var v23 int = (^(fi19))
                                    var v24 uint16 = (v4 & (v4 % ((1) | 1)))
                                    _ = v23
                                    _ = v24
                                }
                                _ = fi22
                            }
                        default:
                            v6 += (((v6 / v6) * ((((0.1 / float32(3.14159)) - v6) - (v6 - ((v6 + 0.0) + v6))) * ((0.0 * (0.0 / v6)) - 0.1))) + float32((((((1e10 - 3.14159) - 0.1) - 0.1) - 0.0) / 3.14159)))
                        }
                        if true {
                            fmt.Println(v6, v2)
                            var v25 int8 = 0
                            fmt.Println(v1, v2)
                            _ = v25
                        } else {
                            var v26 int32 = 748
                            v5[fi19] = true
                            _ = v26
                        }
                        _ = v21
                    }
                }
            case 0:
                for fi27 := 0; fi27 < 4; fi27++ {
                    if ((((v1 + ("hello" + v1)) + v1) < ((v1 + "hello") + "hello")) && ((v2 << ((1) & 63)) <= v2)) {
                        var v28 int = fi27
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        _ = v28
                    } else {
                        switch (v1 + (v1 + v1)) {
                        case "":
                            var v29 uint16 = (^(v4))
                            _ = v29
                        case "hello":
                            v5[(^(fi27))] = (9223372036854775807 == -724)
                            switch fi27 {
                            case 0:
                                if (false && false) {
                                    var v30 float64 = (float64(v6) + 0.0)
                                    v4 = (v4 * v4)
                                    if ("a" == v1) {
                                        switch "\n\t" {
                                        case "a":
                                            v6 = 1.5
                                        case "hello":
                                            fmt.Println(v2)
                                        default:
                                            v30 = v30
                                        }
                                    }
                                    v1 = ((v1 + "foo bar") + v1)
                                    _ = v30
                                } else {
                                    fmt.Println(v1, v4, v4)
                                    var v31 uint16 = v4
                                    _ = v31
                                }
                            case -1:
                                var v32 uint = 765
                                _ = v32
                            case 127:
                                var v33 int64 = 406
                                _ = v33
                            }
                            if (false && false) {
                                v34 := GenS2{F0: 229, F1: 1, F2: true}
                                _ = v34
                            } else {
                                fmt.Println(fi27, v6)
                            }
                        }
                        var v35 uint64 = 1
                        fmt.Println(v2)
                        if true {
                            fmt.Println(v2, fi19)
                        } else {
                            switch fi19 {
                            case 1:
                                v36 := GenS4{F0: 0, F1: 2147483647, F2: -473}
                                _ = v36
                            }
                            v35 -= 42
                        }
                        var v37 int16 = v2
                        _ = v35
                        _ = v37
                    }
                    switch v1 {
                    case "hello":
                        if true {
                            var v38 uint8 = 255
                            var v39 int8 = 42
                            v40 := GenS4{F0: 232, F1: 429, F2: 42}
                            _ = v38
                            _ = v39
                            _ = v40
                        } else {
                            switch (fi19 >> ((18446744073709551615) & 63)) {
                            case 0:
                                var v41 uint32 = 2
                                _ = v41
                            default:
                                v42 := GenS2{F0: -463, F1: 529, F2: false}
                                _ = v42
                            }
                            if false {
                                v43 := GenS2{F0: -2147483648, F1: 523, F2: (18446744073709551615 < 42)}
                                v6 = float32(1.5)
                                v5[fi19] = true
                                _ = v43
                            }
                        }
                    case "\n\t":
                        v4 = v4
                        v44 := GenS5{F0: 184, F1: 157, F2: 127}
                        _ = v44
                    case "a":
                        fmt.Println(v1, v4)
                    default:
                        _ = len(v5)
                    }
                    _ = fi27
                }
            }
            if ((1e10 + (3.14159 * v6)) <= v6) {
                if (true && (526 == 65535)) {
                    v6 = 1e10
                    v1 = v1
                    v45 := GenS6{F0: 879}
                    _ = v45
                }
                var v46 float64 = 0.1
                _ = len(v5)
                _ = v46
            }
            _ = fi19
        }
        _ = v11
    }
    v4 = v4
    _ = len(v3)
    v47 := GenS7{F0: -1, F1: v4, F2: (v4 % (((v4 / ((2) | 1))) | 1))}
    delete(v5, -2147483648)
    v2 = (v2 % (((-(v2))) | 1))
    _ = len(v5)
    switch 4294967295 {
    case -440:
        switch v1 {
        case "hello":
            for fi48 := 0; fi48 < 15; fi48++ {
                v2 += (v2 % ((v2) | 1))
                _ = fi48
            }
        default:
            if (-128 >= 114) {
                v3 = append(v3, 42)
                if (!(((-2.25 + (1.5 * (3.14159 + 3.14159))) / (3.14159 + -2.25)) > 0.0)) {
                    switch 127 {
                    case 689:
                        var v49 int64 = -1
                        _ = v49
                    case 42:
                        _ = len(v3)
                    case 32768:
                        switch (("a" + v1) + "\n\t") {
                        case "a":
                            v4 += v4
                        case "foo bar":
                            switch 0 {
                            case 4294967295:
                                if true {
                                    v1 = ("" + v1)
                                    for fi50 := 0; fi50 < 15; fi50++ {
                                        if (false && ("a" >= (v1 + v1))) {
                                            var v51 int = (-(fi50))
                                            v6 -= v6
                                            _ = v51
                                        } else {
                                            v5[-1] = true
                                            var v52 float64 = (3.14159 - 3.14159)
                                            var v53 uint8 = 0
                                            _ = v52
                                            _ = v53
                                        }
                                        v6 = (v6 - v6)
                                        v6 = (((float32(1e10) / v6) - 1.5) * (float32(1.5) + v6))
                                        _ = fi50
                                    }
                                }
                            case 2147483647:
                                var v54 uint8 = 2
                                _ = v54
                            default:
                                var v55 uint8 = 1
                                _ = v55
                            }
                        case "\n\t":
                            var v56 int8 = 0
                            _ = v56
                        }
                    default:
                        var v57 uint16 = (v4 + v4)
                        _ = v57
                    }
                } else {
                    if false {
                        for fi58 := 0; fi58 < 5; fi58++ {
                            if false {
                                v6 = v6
                                v4 += 479
                                v1 = v1
                            } else {
                                var v59 string = "\n\t"
                                for fi60 := 0; fi60 < 3; fi60++ {
                                    if true {
                                        var v61 uint32 = 703
                                        _ = v61
                                    } else {
                                        var v62 uint64 = 18446744073709551615
                                        _ = v62
                                    }
                                    _ = fi60
                                }
                                _ = v59
                            }
                            for fi63 := 0; fi63 < 19; fi63++ {
                                if (1.5 == 3.14159) {
                                    if ((v2 & (v2 - v2)) != int16(fi63)) {
                                        v3 = append(v3, 2147483648)
                                        v64 := GenS8{F0: -1, F1: 1}
                                        _ = v64
                                    } else {
                                        var v65 uint32 = 400
                                        _ = v65
                                    }
                                    v66 := v5[(fi58 - 1)]
                                    if (!(2 == 1)) {
                                        var v67 uint32 = 2147483648
                                        if v66 {
                                            _ = len(v3)
                                            switch (fi63 << ((42) & 63)) {
                                            case 127:
                                                var v68 uint = 42
                                                _ = v68
                                            default:
                                                v2 -= (v2 / ((v2) | 1))
                                            }
                                            var v69 int32 = -2147483648
                                            _ = v69
                                        } else {
                                            if (!(("foo bar" + (v1 + "hello")) <= v1)) {
                                                fmt.Println(fi63, v1)
                                            }
                                            if (fi58 != (fi58 >> ((1) & 63))) {
                                                _ = len(v3)
                                            }
                                        }
                                        _ = v67
                                    } else {
                                        _ = len(v5)
                                    }
                                    _ = v66
                                } else {
                                    var v70 bool = true
                                    var v71 int8 = 127
                                    _ = v70
                                    _ = v71
                                }
                                if len(v3) > 0 {
                                    v3[0] = 745
                                }
                                fmt.Println(v4)
                                _ = fi63
                            }
                            for fi72 := 0; fi72 < 4; fi72++ {
                                var v73 int32 = 0
                                _ = fi72
                                _ = v73
                            }
                            _ = fi58
                        }
                        if (true || false) {
                            for fi74 := 0; fi74 < 1; fi74++ {
                                for fi75 := 0; fi75 < 20; fi75++ {
                                    var v76 float64 = 0.1
                                    v77 := GenS7{F0: 1, F1: v4, F2: (v4 << ((4294967295) & 63))}
                                    _ = fi75
                                    _ = v76
                                    _ = v77
                                }
                                var v78 int32 = 231
                                var v79 int8 = -1
                                _ = fi74
                                _ = v78
                                _ = v79
                            }
                        }
                    } else {
                        v1 = (v1 + "a")
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    }
                    fmt.Println(v6, v1, v4)
                }
                var v80 float32 = v6
                _ = v80
            }
        }
    }
    fmt.Println(v6)
    v4 = 42
    v81 := GenS9{F0: (1 < -128)}
    if (false || (!false)) {
        switch "a" {
        case "a":
            v1 = "foo bar"
        case "\n\t":
            v1 = (v1 + v1)
        case "foo bar":
            if false {
                fmt.Println(v4, v2, v6)
                if true {
                    var v82 int64 = 941
                    if true {
                        switch (v1 + ((v1 + "") + (((((("" + "a") + v1) + "a") + ("a" + "\n\t")) + ("hello" + (v1 + "a"))) + v1))) {
                        case "foo bar":
                            if true {
                                v4 -= v4
                            } else {
                                v4 = (^(v4))
                                switch 42 {
                                case 127:
                                    v4 += (v4 / ((v4) | 1))
                                default:
                                    v82 = 1
                                }
                                var v83 int16 = v2
                                _ = v83
                            }
                        case "hello":
                            v2 -= (v2 >> ((643) & 63))
                        case "a":
                            var v84 int = 128
                            _ = v84
                        }
                        var v85 int32 = 2147483647
                        _ = v85
                    } else {
                        var v86 float64 = (float64(v6) + 0.1)
                        _ = v86
                    }
                    _ = v82
                }
                var v87 int = -623
                _ = v87
            } else {
                v2 -= -368
                if false {
                    v3 = append(v3, 803)
                    v5[-9223372036854775808] = ((-1 <= 1) || (452 != 42))
                }
            }
        }
    } else {
        var v88 int16 = v2
        _ = v88
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v5
    _ = v6
    _ = v7
    _ = v81
}
