package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int64; F1 int16; F2 uint }
type GenS3 struct { F0 uint8; F1 uint }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint32; F1 uint64; F2 uint64 }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 uint64; F1 bool; F2 uint }
type GenS9 struct { F0 int8 }
type GenS10 struct { F0 uint16 }
func genFn0() uint8 {
    for fi6 := 0; fi6 < 9; fi6++ {
        v7 := []int{32767, fi6, (-(fi6)), -2147483648}
        var v8 uint8 = 130
        var v9 uint8 = v8
        _ = fi6
        _ = v7
        _ = v8
        _ = v9
    }
    var v10 float64 = 1.5
    _ = v10
    return 2
}

func main() {
    var v1 int16 = 32767
    var v2 float64 = 1.5
    var v3 string = (("" + "\n\t") + (("a" + "a") + "\n\t"))
    v4 := make(map[int]uint32)
    var v5 int8 = 42
    v2 = (((3.14159 / v2) / v2) * 0.0)
    v11 := []int{0, 4294967295, -9223372036854775808, -393}
    fmt.Println(v1, v3, v1)
    v1 = (v1 * (v1 - (-(v1))))
    if true {
        v12 := v4[128]
        v12 += (^(v12))
        var v13 int8 = 100
        _ = v12
        _ = v13
    } else {
        switch 4294967295 {
        case -9223372036854775808:
            var v14 uint16 = 739
            _ = v14
        case 154:
            var v15 float32 = ((-2.25 + 0.1) + (((((-2.25 * (((1.5 / -2.25) + 1e10) / ((1.5 - (0.0 * (1e10 * (-2.25 * 0.0)))) / 1.0))) + (0.0 + (((float32(((v2 + (((1.5 * 0.0) * float64(3.14159)) + v2)) - ((-2.25 / v2) + v2))) * -2.25) + 0.0) + -2.25))) - ((0.1 * (3.14159 * (1.5 + (float32(((1.5 + v2) * (v2 * (v2 * v2)))) + 3.14159)))) / (0.1 / (-2.25 - 3.14159)))) + 3.14159) / (float32(0.1) + (1.5 - 1e10))))
            _ = v15
        default:
            var v16 int = 2147483647
            _ = v16
        }
    }
    var v17 uint8 = 42
    v18 := GenS1{F0: 7}
    fmt.Println(v3)
    if (false && ((429 != 136) || ((v1 / ((-303) | 1)) != 32767))) {
        for fi19 := 0; fi19 < 11; fi19++ {
            switch ("" + v3) {
            case "hello":
                delete(v4, fi19)
            case "a":
                if (2 < 323) {
                    v20 := GenS2{F0: 9223372036854775807, F1: v1, F2: 2}
                    if false {
                        var v21 string = v3
                        v17 += 1
                        _ = v21
                    }
                    if true {
                        v22 := GenS2{F0: -196, F1: (v1 / ((788) | 1)), F2: 1}
                        if false {
                            if (!(!((fi19 + (-(fi19))) > 32767))) {
                                var v23 int16 = v1
                                if (!((fi19 ^ 128) >= int(v17))) {
                                    v2 -= (-2.25 * (0.1 + -2.25))
                                }
                                fmt.Println(v23, fi19, v3)
                                _ = v23
                            } else {
                                if true {
                                    v11 = append(v11, (fi19 >> ((2) & 63)))
                                    var v24 uint = 4294967295
                                    v17 += v17
                                    _ = v24
                                } else {
                                    for fi25 := 0; fi25 < 15; fi25++ {
                                        v5 = -12
                                        var v26 string = v3
                                        v2 -= (v2 / ((v2 - (v2 + (v2 * (float64(-2.25) - (0.1 / v2))))) / v2))
                                        _ = fi25
                                        _ = v26
                                    }
                                    var v27 int8 = (v5 % ((v5) | 1))
                                    _ = v27
                                }
                            }
                            v28 := GenS1{F0: 42}
                            switch (v3 + ("foo bar" + v3)) {
                            case "foo bar":
                                for fi29 := 0; fi29 < 19; fi29++ {
                                    v30 := GenS3{F0: 42, F1: 551}
                                    if len(v11) > 0 {
                                        v11[0] = (^(fi29))
                                    }
                                    _ = fi29
                                    _ = v30
                                }
                            case "":
                                v3 = "\n\t"
                            case "\n\t":
                                v31 := GenS3{F0: (v17 & genFn0()), F1: uint(v17)}
                                _ = v31
                            }
                            _ = v28
                        }
                        var v32 int = fi19
                        _ = v22
                        _ = v32
                    }
                    _ = v20
                } else {
                    var v33 uint8 = v17
                    v5 -= v5
                    _ = v33
                }
            case "foo bar":
                for fi34 := 0; fi34 < 7; fi34++ {
                    v5 = 2
                    switch fi19 {
                    case -938:
                        var v35 int16 = (v1 >> ((65535) & 63))
                        _ = v35
                    default:
                        switch 42 {
                        case 128:
                            v36 := GenS1{F0: 4294967295}
                            _ = v36
                        }
                    }
                    v37 := GenS1{F0: 259}
                    _ = fi34
                    _ = v37
                }
            default:
                switch (fi19 | fi19) {
                case -323:
                    var v38 uint8 = v17
                    _ = v38
                case -758:
                    if true {
                        var v39 uint8 = 128
                        var v40 uint32 = 2147483647
                        v41 := GenS2{F0: -1, F1: v1, F2: 4294967295}
                        _ = v39
                        _ = v40
                        _ = v41
                    } else {
                        switch ("\n\t" + ("hello" + v3)) {
                        case "\n\t":
                            var v42 uint = 231
                            _ = v42
                        case "hello":
                            fmt.Println(v2, v17, v2)
                        default:
                            fmt.Println(v3, v17)
                        }
                    }
                case 533:
                    fmt.Println(v1, v1)
                default:
                    v3 = "foo bar"
                }
            }
            switch ("" + v3) {
            case "a":
                var v43 uint32 = 0
                _ = v43
            default:
                if (3.14159 < -2.25) {
                    v3 = v3
                    v17 -= (v17 / ((v17) | 1))
                    var v44 uint16 = 65535
                    _ = v44
                }
            }
            _ = fi19
        }
        v17 = v17
    } else {
        var v45 bool = true
        v2 -= v2
        if ((836 <= 943) && (4294967295 < 2)) {
            var v46 int8 = v5
            v2 = v2
            for fi47 := 0; fi47 < 17; fi47++ {
                var v48 uint32 = 650
                switch v3 {
                case "hello":
                    for fi49 := 0; fi49 < 5; fi49++ {
                        v50 := GenS4{F0: true}
                        if (int16(v48) != (v1 % ((v1) | 1))) {
                            switch 42 {
                            case -9223372036854775808:
                                fmt.Println(v2, fi49, v45)
                            case 128:
                                if v45 {
                                    fmt.Println(v1, v1, fi49)
                                }
                            default:
                                v1 -= v1
                            }
                            if len(v11) > 0 {
                                v11[0] = -2147483648
                            }
                            v2 = float64(((1e10 * 3.14159) + 1.5))
                            v51 := GenS4{F0: v45}
                            if (!true) {
                                v5 = (v5 & -1)
                                v2 = v2
                                fmt.Println(v1, v17, fi49)
                            }
                            v1 += v1
                            switch fi47 {
                            case 128:
                                switch v3 {
                                case "foo bar":
                                    var v52 uint16 = 32767
                                    var v53 int8 = (v46 - v5)
                                    var v54 int = (fi49 & -263)
                                    _ = v52
                                    _ = v53
                                    _ = v54
                                case "hello":
                                    var v55 int64 = -1
                                    _ = v55
                                case "":
                                    delete(v4, 127)
                                }
                            case 589:
                                switch v3 {
                                case "a":
                                    if (v48 < 445) {
                                        if (!false) {
                                            v56 := v4[fi49]
                                            _ = v56
                                        } else {
                                            v46 += v5
                                            fmt.Println(v3, v45)
                                        }
                                    }
                                case "\n\t":
                                    v48 += 2147483648
                                }
                            case 928:
                                if v45 {
                                    v5 = 65
                                    var v57 uint64 = 787
                                    _ = v57
                                } else {
                                    fmt.Println(v2)
                                    switch "" {
                                    case "hello":
                                        var v58 int32 = 1
                                        _ = v58
                                    }
                                }
                                var v59 uint32 = 752
                                switch "a" {
                                case "hello":
                                    v2 -= (0.0 + (v2 / v2))
                                case "foo bar":
                                    var v60 float64 = v2
                                    _ = v60
                                case "a":
                                    v61 := GenS5{F0: v45}
                                    _ = v61
                                default:
                                    var v62 uint16 = 0
                                    _ = v62
                                }
                                _ = v59
                            }
                            if (true && false) {
                                var v63 int = (fi49 ^ (fi47 << ((4294967295) & 63)))
                                var v64 int64 = 824
                                _ = v63
                                _ = v64
                            }
                            if (1 > (v5 - (v46 / ((int8(fi47)) | 1)))) {
                                if true {
                                    switch fi49 {
                                    case 248:
                                        v3 = (v3 + v3)
                                    case 32768:
                                        v46 += v46
                                        var v65 int32 = -854
                                        if (!true) {
                                            if true {
                                                var v66 uint64 = 18446744073709551615
                                                var v67 float32 = float32(v2)
                                                switch fi49 {
                                                case -595:
                                                    fmt.Println(fi49, v5, v5)
                                                case 0:
                                                    var v68 uint = 0
                                                    _ = v68
                                                }
                                                var v69 bool = ((v65 * v65) != v65)
                                                if (int(v66) < (-(fi49))) {
                                                    _ = len(v11)
                                                    var v70 string = (v3 + "")
                                                    _ = v70
                                                } else {
                                                    if (!(-14 <= (fi47 + fi49))) {
                                                        var v71 uint16 = 1
                                                        v72 := v4[fi49]
                                                        v73 := GenS4{F0: false}
                                                        _ = v71
                                                        _ = v72
                                                        _ = v73
                                                    } else {
                                                        v74 := GenS6{F0: v48, F1: (v66 | v66), F2: v66}
                                                        _ = v74
                                                    }
                                                }
                                                _ = v66
                                                _ = v67
                                                _ = v69
                                            }
                                            var v75 uint = 206
                                            _ = v75
                                        }
                                        _ = v65
                                    case 236:
                                        v2 = v2
                                    default:
                                        v76 := GenS2{F0: -388, F1: -1, F2: 0}
                                        _ = v76
                                    }
                                }
                                switch 32768 {
                                case 42:
                                    var v77 uint8 = (v17 ^ genFn0())
                                    _ = v77
                                default:
                                    fmt.Println(v3, v48, v45)
                                    fmt.Println(v1, v2, v17)
                                    var v78 int32 = 1
                                    _ = v78
                                }
                                var v79 uint32 = v48
                                _ = v79
                            } else {
                                switch (fi49 & (fi47 + fi49)) {
                                case 32767:
                                    v17 -= v17
                                case 0:
                                    v45 = (true && true)
                                default:
                                    if v45 {
                                        if ((fi47 >> ((18446744073709551615) & 63)) > 4294967295) {
                                            switch "a" {
                                            case "\n\t":
                                                v80 := v4[(fi49 | 32767)]
                                                _ = v80
                                            }
                                        } else {
                                            v11 = append(v11, 180)
                                            var v81 uint8 = 1
                                            v82 := GenS7{F0: 42}
                                            _ = v81
                                            _ = v82
                                        }
                                    } else {
                                        var v83 int = (^(fi47))
                                        v84 := GenS8{F0: uint64(v48), F1: ((0 >= 0) || (!true)), F2: 65535}
                                        var v85 int16 = int16(v46)
                                        _ = v83
                                        _ = v84
                                        _ = v85
                                    }
                                }
                                v46 = v46
                                if v45 {
                                    var v86 int = (fi49 ^ fi49)
                                    if v45 {
                                        if v45 {
                                            v17 = 127
                                            var v87 uint16 = 1
                                            _ = v87
                                        } else {
                                            var v88 uint16 = 32767
                                            var v89 string = "foo bar"
                                            _ = v88
                                            _ = v89
                                        }
                                        fmt.Println(v46, v2, v48)
                                    } else {
                                        v17 = v17
                                        v90 := GenS9{F0: (v46 / ((int8(v17)) | 1))}
                                        if (42 > 1) {
                                            var v91 float64 = 0.1
                                            v1 = (v1 | -985)
                                            var v92 uint = 65535
                                            _ = v91
                                            _ = v92
                                        } else {
                                            switch 1 {
                                            case 9223372036854775807:
                                                v93 := GenS4{F0: v45}
                                                _ = v93
                                            case -2147483648:
                                                v45 = (true && v45)
                                            }
                                            v5 -= v46
                                            if v45 {
                                                v45 = v45
                                            } else {
                                                switch -408 {
                                                case 42:
                                                    switch (fi49 >> ((541) & 63)) {
                                                    case 958:
                                                        fmt.Println(v3, fi49, v1)
                                                    }
                                                }
                                            }
                                        }
                                        _ = v90
                                    }
                                    if (65535 <= v48) {
                                        v4[v86] = uint32(v17)
                                        fmt.Println(v46, v1, v2)
                                        var v94 uint8 = v17
                                        _ = v94
                                    } else {
                                        v48 -= (v48 / (((^(v48))) | 1))
                                        v95 := GenS5{F0: v45}
                                        fmt.Println(v48, v48)
                                        _ = v95
                                    }
                                    _ = v86
                                } else {
                                    v96 := GenS10{F0: 42}
                                    switch (v3 + ("foo bar" + v3)) {
                                    case "hello":
                                        _ = len(v4)
                                    case "\n\t":
                                        var v97 string = v3
                                        _ = v97
                                    }
                                    _ = v96
                                }
                            }
                            _ = v51
                        }
                        var v98 uint = 953
                        _ = fi49
                        _ = v50
                        _ = v98
                    }
                case "a":
                    if true {
                        _ = len(v11)
                    }
                }
                _ = fi47
                _ = v48
            }
            _ = v46
        }
        _ = v45
    }
    if false {
        v3 = v3
    } else {
        switch 1 {
        case 42:
            var v99 uint = 18446744073709551615
            _ = v99
        case 513:
            v2 = v2
        }
    }
    for fi100 := 0; fi100 < 13; fi100++ {
        if (!((v17 * v17) <= genFn0())) {
            if (!false) {
                fmt.Println(v2, v17)
            }
            v101 := v4[(fi100 / ((fi100) | 1))]
            fmt.Println(v2)
            _ = v101
        }
        fmt.Println(fi100)
        _ = fi100
    }
    v1 += (v1 >> ((255) & 63))
    var v102 string = ("\n\t" + (("" + v3) + v3))
    v102 = ("\n\t" + v102)
    if (true || false) {
        for fi103 := 0; fi103 < 9; fi103++ {
            var v104 int8 = 46
            var v105 int32 = -294
            var v106 uint = 18446744073709551615
            _ = fi103
            _ = v104
            _ = v105
            _ = v106
        }
    } else {
        v107 := v4[0]
        _ = v107
    }
    fmt.Println("v1:", v1)
    fmt.Println("v102:", v102)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v102
    _ = v11
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
