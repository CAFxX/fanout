package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16 }
type GenS2 struct { F0 uint64; F1 int64; F2 int16 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int16; F1 uint64; F2 int64 }
type GenS5 struct { F0 int32; F1 bool; F2 uint8 }
type GenS6 struct { F0 bool; F1 bool }
func genFn0(p0 uint8) int8 {
    var v4 uint32 = 2
    v5 := []int{-806, 1, 0}
    v6 := GenS1{F0: -400, F1: 32767}
    v5 = append(v5, -1)
    var v7 uint = 634
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return 1
}

func main() {
    var v1 bool = false
    var v2 int8 = 42
    v3 := make(map[int]bool)
    v1 = true
    var v8 uint8 = 128
    if ((1 == v8) || ((float64((1.5 / (1e10 + float32((3.14159 - 3.14159))))) - float64(-2.25)) != 1.5)) {
        if ((!(true && (!v1))) && false) {
            fmt.Println(v8)
            if (510 > 154) {
                if (!(335 <= 127)) {
                    v9 := GenS1{F0: 32767, F1: 2}
                    v1 = (v1 || (!v1))
                    delete(v3, 9223372036854775807)
                    _ = v9
                }
            } else {
                if v1 {
                    v10 := []int{}
                    _ = v10
                }
                switch (("" + "a") + "") {
                case "\n\t":
                    v11 := GenS2{F0: 42, F1: -299, F2: -32768}
                    _ = v11
                case "foo bar":
                    switch -1 {
                    case 32768:
                        v12 := GenS2{F0: 1, F1: 42, F2: -32768}
                        _ = v12
                    case 425:
                        if v1 {
                            v13 := v3[42]
                            var v14 int = -49
                            _ = v13
                            _ = v14
                        } else {
                            if (-1 >= 0) {
                                var v15 uint = 418
                                v16 := GenS2{F0: 2, F1: 0, F2: -32768}
                                v17 := GenS1{F0: 298, F1: 121}
                                _ = v15
                                _ = v16
                                _ = v17
                            }
                            var v18 uint = 261
                            v8 = v8
                            _ = v18
                        }
                    }
                }
            }
        } else {
            if v1 {
                fmt.Println(v1, v1, v1)
            } else {
                var v19 uint64 = 18446744073709551615
                switch "a" {
                case "\n\t":
                    var v20 uint8 = (v8 >> ((123) & 63))
                    _ = v20
                default:
                    for fi21 := 0; fi21 < 11; fi21++ {
                        v22 := []int{fi21}
                        var v23 uint64 = (^(v19))
                        _ = fi21
                        _ = v22
                        _ = v23
                    }
                }
                var v24 uint16 = 1
                _ = v19
                _ = v24
            }
        }
        for fi25 := 0; fi25 < 1; fi25++ {
            fmt.Println(v1, fi25)
            _ = fi25
        }
        switch 32768 {
        case 32768:
            v2 -= (v2 % ((v2) | 1))
        case -58:
            v26 := []int{}
            _ = v26
        default:
            v1 = (v1 && v1)
        }
    } else {
        fmt.Println(v2)
    }
    v2 -= (v2 / (((v2 - (-(v2)))) | 1))
    var v27 int32 = 0
    v28 := GenS3{F0: 0}
    var v29 uint64 = 18446744073709551615
    for fi30 := 0; fi30 < 2; fi30++ {
        v8 += 128
        if ((65535 >= 442) && (false || (32768 >= fi30))) {
            for fi31 := 0; fi31 < 18; fi31++ {
                v27 -= v27
                var v32 uint64 = v29
                switch (fi31 >> ((689) & 63)) {
                case 1:
                    switch 9223372036854775807 {
                    case 604:
                        var v33 uint = 436
                        _ = v33
                    default:
                        var v34 float64 = 3.14159
                        _ = v34
                    }
                case -9223372036854775808:
                    v27 -= 1
                }
                switch ("a" + "foo bar") {
                case "":
                    if v1 {
                        switch "" {
                        case "a":
                            var v35 uint8 = (v8 / ((v8) | 1))
                            _ = v35
                        default:
                            v36 := []int{fi30, 94, fi30}
                            _ = v36
                        }
                        v3[(^(fi30))] = true
                    }
                default:
                    if (!false) {
                        fmt.Println(v29)
                        if true {
                            fmt.Println(v32, v1, v2)
                            delete(v3, fi30)
                        }
                        v8 += (v8 | 127)
                    } else {
                        v29 += v32
                    }
                }
                _ = fi31
                _ = v32
            }
            var v37 uint8 = v8
            _ = v37
        }
        _ = fi30
    }
    var v38 uint = 0
    v8 = v8
    v2 -= (v2 ^ -128)
    for fi39 := 0; fi39 < 17; fi39++ {
        fmt.Println(v1, v1)
        if (!true) {
            var v40 uint8 = 127
            var v41 bool = true
            _ = v40
            _ = v41
        } else {
            v1 = false
            var v42 uint = (v38 | v38)
            _ = v42
        }
        _ = fi39
    }
    var v43 string = ("foo bar" + ("\n\t" + (("foo bar" + "a") + ("a" + "foo bar"))))
    switch 128 {
    case 131:
        switch 985 {
        case -2147483648:
            var v44 int32 = int32(v29)
            _ = v44
        case 42:
            _ = len(v3)
        default:
            v45 := GenS1{F0: 42, F1: 777}
            _ = v45
        }
    case 128:
        for fi46 := 0; fi46 < 20; fi46++ {
            for fi47 := 0; fi47 < 3; fi47++ {
                v48 := v3[(fi47 % (((fi46 + 9223372036854775807)) | 1))]
                switch (("hello" + v43) + "\n\t") {
                case "a":
                    var v49 int64 = -1
                    _ = v49
                }
                delete(v3, 1)
                var v50 uint = 2
                _ = fi47
                _ = v48
                _ = v50
            }
            v1 = v1
            _ = fi46
        }
    case 42:
        if true {
            var v51 int32 = 1
            for fi52 := 0; fi52 < 9; fi52++ {
                for fi53 := 0; fi53 < 10; fi53++ {
                    v54 := []int{fi52, (fi53 & 1), (fi53 | (-(fi52))), (fi52 - (^(fi52)))}
                    v55 := v3[(^(fi52))]
                    fmt.Println(v38)
                    _ = fi53
                    _ = v54
                    _ = v55
                }
                _ = fi52
            }
            _ = v51
        }
    default:
        v56 := GenS4{F0: -1, F1: (v29 & 42), F2: 904}
        _ = v56
    }
    for fi57 := 0; fi57 < 17; fi57++ {
        fmt.Println(v1, v8)
        v2 -= genFn0((v8 ^ v8))
        _ = fi57
    }
    v3[4294967295] = true
    if ((v43 >= ("foo bar" + v43)) && ((("foo bar" + v43) + "\n\t") <= ((v43 + "foo bar") + v43))) {
        if false {
            for fi58 := 0; fi58 < 18; fi58++ {
                var v59 uint8 = (v8 << ((255) & 63))
                if (v1 || (v59 <= 0)) {
                    if true {
                        v29 = v29
                    } else {
                        v60 := GenS5{F0: 0, F1: (v27 >= int32(v59)), F2: v8}
                        v61 := v3[(fi58 + fi58)]
                        switch v43 {
                        case "\n\t":
                            _ = len(v3)
                        default:
                            v62 := GenS2{F0: (v29 % ((1) | 1)), F1: 9223372036854775807, F2: -47}
                            _ = v62
                        }
                        _ = v60
                        _ = v61
                    }
                    v3[(fi58 + fi58)] = v1
                    v43 = (v43 + ((("a" + ("\n\t" + v43)) + (v43 + v43)) + v43))
                }
                v63 := GenS3{F0: 65535}
                _ = fi58
                _ = v59
                _ = v63
            }
            switch (v43 + (v43 + v43)) {
            case "":
                switch -2147483648 {
                case 32767:
                    if ((!v1) && true) {
                        v64 := []int{128, 127}
                        _ = v64
                    } else {
                        if v1 {
                            fmt.Println(v27)
                        } else {
                            fmt.Println(v43)
                        }
                    }
                case 497:
                    switch v43 {
                    case "a":
                        if (899 != 0) {
                            if true {
                                for fi65 := 0; fi65 < 19; fi65++ {
                                    var v66 float32 = (((0.0 - (1.5 * 0.0)) * 0.1) * -2.25)
                                    _ = fi65
                                    _ = v66
                                }
                                v67 := GenS3{F0: 65535}
                                v1 = v1
                                _ = v67
                            }
                            fmt.Println(v1)
                        } else {
                            if ((true && v1) && v1) {
                                v68 := GenS5{F0: int32(v8), F1: (42 >= (v27 ^ v27)), F2: v8}
                                _ = v68
                            }
                        }
                    default:
                        _ = len(v3)
                    }
                case 32768:
                    v69 := []int{-9223372036854775808, 32767, 0}
                    _ = v69
                default:
                    v70 := []int{int(v8), -725, 4294967295}
                    _ = v70
                }
            default:
                if true {
                    v27 = (-(v27))
                    for fi71 := 0; fi71 < 12; fi71++ {
                        if (!(!false)) {
                            if v1 {
                                var v72 uint64 = (v29 & (v29 << (((v38 - uint(v29))) & 63)))
                                switch fi71 {
                                case 4294967295:
                                    fmt.Println(v27, v1, v8)
                                case -966:
                                    var v73 float64 = ((float64(1.5) / 1e10) / 1.0)
                                    _ = v73
                                case 2147483647:
                                    for fi74 := 0; fi74 < 13; fi74++ {
                                        var v75 uint16 = 379
                                        if ((!(v1 && true)) || true) {
                                            v29 = uint64(v75)
                                            v27 += v27
                                        } else {
                                            v43 = v43
                                            fmt.Println(v2)
                                        }
                                        _ = fi74
                                        _ = v75
                                    }
                                }
                                for fi76 := 0; fi76 < 18; fi76++ {
                                    if (int64(v27) != -1) {
                                        var v77 uint32 = 425
                                        if (-1 <= -32768) {
                                            v78 := GenS4{F0: -32768, F1: v72, F2: 1}
                                            v27 -= (v27 * int32(v77))
                                            v79 := []int{}
                                            if v1 {
                                                v29 = (v29 + 1)
                                                if true {
                                                    switch v43 {
                                                    case "hello":
                                                        if ((3.14159 / -2.25) > 3.14159) {
                                                            var v80 int8 = (v2 / ((100) | 1))
                                                            v2 += genFn0(255)
                                                            switch (fi76 >> (((v38 + (v38 | v38))) & 63)) {
                                                            case 0:
                                                                v8 -= (v8 << ((v38) & 63))
                                                            case 32768:
                                                                v81 := GenS5{F0: (v27 + (v27 + v27)), F1: v1, F2: 1}
                                                                _ = v81
                                                            }
                                                            _ = v80
                                                        } else {
                                                            fmt.Println(v29, v8, v29)
                                                        }
                                                    case "":
                                                        switch v43 {
                                                        case "":
                                                            fmt.Println(v77, v77)
                                                        case "hello":
                                                            var v82 bool = (fi71 > 4294967295)
                                                            _ = v82
                                                        case "foo bar":
                                                            if len(v79) >= 2 {
                                                                v79 = v79[1:len(v79)]
                                                            }
                                                        default:
                                                            if v1 {
                                                                v8 -= 100
                                                                if v1 {
                                                                    var v83 uint64 = v29
                                                                    _ = v83
                                                                } else {
                                                                    switch "foo bar" {
                                                                    case "\n\t":
                                                                        if true {
                                                                            v84 := v3[fi76]
                                                                            if v1 {
                                                                                var v85 uint = uint(v72)
                                                                                v27 -= v27
                                                                                _ = v85
                                                                            } else {
                                                                                v43 = ("foo bar" + v43)
                                                                                fmt.Println(fi71)
                                                                            }
                                                                            _ = v84
                                                                        }
                                                                        var v86 int8 = v2
                                                                        _ = v86
                                                                    case "":
                                                                        v87 := GenS6{F0: (-1 == -766), F1: true}
                                                                        _ = v87
                                                                    default:
                                                                        if ((-9223372036854775808 < fi71) && false) {
                                                                            v88 := GenS5{F0: int32(v2), F1: (v2 >= v2), F2: (v8 & uint8(v72))}
                                                                            switch "\n\t" {
                                                                            case "\n\t":
                                                                                var v89 string = ""
                                                                                _ = v89
                                                                            }
                                                                            var v90 uint = 255
                                                                            fmt.Println(v43, v27, v90)
                                                                            _ = v88
                                                                            _ = v90
                                                                        }
                                                                    }
                                                                    fmt.Println(v1)
                                                                }
                                                            }
                                                        }
                                                    }
                                                    var v91 uint32 = (v77 << ((v38) & 63))
                                                    _ = v91
                                                } else {
                                                    var v92 string = (v43 + "a")
                                                    fmt.Println(v38, v29)
                                                    _ = v92
                                                }
                                                if (!false) {
                                                    if v1 {
                                                        switch v43 {
                                                        case "a":
                                                            v93 := GenS5{F0: (v27 | 0), F1: (v1 || ((("a" + v43) <= "a") || v1)), F2: v8}
                                                            _ = v93
                                                        case "\n\t":
                                                            fmt.Println(v2, v1, v8)
                                                        case "hello":
                                                            v94 := v3[fi71]
                                                            _ = v94
                                                        default:
                                                            v8 = v8
                                                        }
                                                        fmt.Println(v72)
                                                        if (!v1) {
                                                            var v95 float64 = float64(((1.5 / 1e10) * (-2.25 * float32(((3.14159 * (0.0 * (3.14159 * 1e10))) * float64(1.5))))))
                                                            v96 := GenS4{F0: 1, F1: (v29 >> (((v38 - (v38 - (v38 >> (((v38 & uint(v72))) & 63))))) & 63)), F2: -9223372036854775808}
                                                            if false {
                                                                v27 += (v27 / (((v27 + v27)) | 1))
                                                                fmt.Println(v2, v95, v43)
                                                            }
                                                            switch (fi71 >> ((2) & 63)) {
                                                            case 32768:
                                                                v3[(fi76 ^ int(v2))] = v1
                                                            case 128:
                                                                if len(v79) > 0 {
                                                                    v79[0] = (^(fi71))
                                                                }
                                                            case 32767:
                                                                var v97 string = v43
                                                                _ = v97
                                                            }
                                                            _ = v95
                                                            _ = v96
                                                        } else {
                                                            switch "foo bar" {
                                                            case "a":
                                                                var v98 uint32 = v77
                                                                _ = v98
                                                            case "foo bar":
                                                                _ = len(v3)
                                                            default:
                                                                v72 = (v29 % ((18446744073709551615) | 1))
                                                            }
                                                        }
                                                        v38 = 0
                                                        v38 = 255
                                                    }
                                                }
                                                var v99 float32 = 3.14159
                                                _ = v99
                                            }
                                            v3[fi76] = (!((v1 || ((v8 / (((v8 & (v8 - v8))) | 1)) <= v8)) && v1))
                                            _ = v78
                                            _ = v79
                                        }
                                        var v100 int64 = -9223372036854775808
                                        v101 := []int{fi76, (fi71 * int(v38))}
                                        _ = v100
                                        _ = v101
                                        _ = v77
                                    }
                                    if ((!(v29 == 1)) && v1) {
                                        fmt.Println(v1, v38)
                                        var v102 uint64 = (v29 ^ (^(v72)))
                                        _ = v102
                                    }
                                    _ = fi76
                                }
                                _ = v72
                            } else {
                                if v1 {
                                    v8 = 255
                                    switch fi71 {
                                    case -8:
                                        var v103 float64 = (-2.25 * float64(0.1))
                                        _ = v103
                                    case 42:
                                        v1 = (v43 < v43)
                                    }
                                    _ = len(v3)
                                } else {
                                    var v104 uint64 = v29
                                    _ = v104
                                }
                            }
                            var v105 int = fi71
                            _ = v105
                        } else {
                            var v106 int = (fi71 >> (((v38 % (((v38 << (((^(v38))) & 63))) | 1))) & 63))
                            v3[fi71] = (-1 < 0)
                            v107 := GenS3{F0: 1}
                            _ = v106
                            _ = v107
                        }
                        _ = fi71
                    }
                }
            }
        }
        if false {
            var v108 uint16 = 42
            v27 = v27
            _ = v108
        } else {
            switch v43 {
            case "\n\t":
                var v109 string = v43
                _ = v109
            case "hello":
                v38 += 1
            case "":
                v3[-9223372036854775808] = true
            }
            v110 := GenS4{F0: -32768, F1: uint64(v38), F2: -9223372036854775808}
            v111 := []int{}
            _ = v110
            _ = v111
        }
        for fi112 := 0; fi112 < 10; fi112++ {
            v8 += 117
            _ = fi112
        }
    }
    v2 += genFn0((v8 << ((v38) & 63)))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("v43:", v43)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v38
    _ = v43
    _ = v8
}
