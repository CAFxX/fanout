package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16 }
type GenS2 struct { F0 int32; F1 uint32 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 string; F1 uint32 }
type GenS6 struct { F0 int64; F1 int32; F2 int }
type GenS7 struct { F0 string; F1 int16 }
type GenS8 struct { F0 uint32; F1 int16; F2 int32 }
type GenS9 struct { F0 int64; F1 int32; F2 uint32 }
type GenS10 struct { F0 uint8; F1 uint16 }

func main() {
    var v1 uint = 42
    var v2 string = "\n\t"
    var v3 string = v2
    var v4 int64 = -708
    var v5 uint32 = 2147483648
    v6 := GenS1{F0: 100, F1: 2}
    if (1000 > -125) {
        if (!true) {
            var v7 uint64 = 1
            _ = v7
        }
        v1 -= v1
    } else {
        var v8 int32 = -792
        _ = v8
    }
    v5 = (v5 << ((v1) & 63))
    switch 127 {
    case -962:
        for fi9 := 0; fi9 < 2; fi9++ {
            for fi10 := 0; fi10 < 6; fi10++ {
                if false {
                    v11 := GenS2{F0: 42, F1: v5}
                    if (!false) {
                        v3 = v3
                        if true {
                            var v12 int8 = 0
                            var v13 int32 = -145
                            v5 -= (v5 & (v5 / (((v5 | (v5 >> (((v1 | (^(v1)))) & 63)))) | 1)))
                            _ = v12
                            _ = v13
                        }
                    }
                    _ = v11
                } else {
                    switch 2147483647 {
                    case -127:
                        var v14 uint32 = (^(v5))
                        v15 := GenS3{F0: 631}
                        _ = v14
                        _ = v15
                    case -1:
                        v16 := []int{fi9, (fi10 << ((18446744073709551615) & 63)), fi9}
                        if true {
                            if true {
                                v17 := GenS2{F0: 2147483647, F1: v5}
                                _ = v17
                            } else {
                                fmt.Println(v4)
                                fmt.Println(fi10, fi9, fi9)
                                v16 = append(v16, fi10)
                                var v18 string = ""
                                _ = v18
                            }
                            v3 = ("a" + "\n\t")
                            v1 = (v1 + (^(v1)))
                        }
                        if ((v5 + v5) >= (^(v5))) {
                            fmt.Println(fi9, fi9, v5)
                            v5 -= (v5 % ((v5) | 1))
                            var v19 int16 = 1000
                            if len(v16) > 0 {
                                v16[0] = 32768
                            }
                            _ = v19
                        }
                        _ = v16
                    case 2147483647:
                        v20 := []int{(^(fi9)), int(v4), fi10}
                        switch (fi9 - (fi10 * 32767)) {
                        case -1:
                            v4 += (v4 | v4)
                        case 32767:
                            var v21 int8 = int8(v1)
                            _ = v21
                        case -2147483648:
                            var v22 bool = (!(0.0 <= (1.5 * -2.25)))
                            _ = v22
                        default:
                            v4 = -1
                        }
                        _ = v20
                    }
                    v23 := []int{fi9, fi9, fi10}
                    if true {
                        v1 = v1
                    } else {
                        v2 = ("hello" + (v3 + v2))
                        switch fi10 {
                        case 42:
                            if len(v23) > 0 {
                                v23[0] = fi9
                            }
                        case 453:
                            fmt.Println(v1, fi9, v4)
                        }
                        switch 2147483647 {
                        case 384:
                            v2 = "a"
                        case 4294967295:
                            if len(v23) >= 2 {
                                v23 = v23[1:len(v23)]
                            }
                        case 9223372036854775807:
                            var v24 int32 = 0
                            _ = v24
                        default:
                            switch v2 {
                            case "a":
                                _ = len(v23)
                            case "":
                                fmt.Println(v1, v3)
                            case "\n\t":
                                if (939 >= 4294967295) {
                                    var v25 uint64 = 4294967295
                                    _ = v25
                                } else {
                                    fmt.Println(v3, v3)
                                    v3 = v3
                                }
                            }
                        }
                    }
                    v4 = v4
                    _ = v23
                }
                if false {
                    var v26 float64 = 0.1
                    v1 = v1
                    v26 = v26
                    _ = v26
                } else {
                    if false {
                        v27 := GenS4{F0: (v3 + v2)}
                        var v28 uint = 1
                        _ = v27
                        _ = v28
                    }
                    v2 = "a"
                    v29 := []int{4294967295, fi9, fi9}
                    if len(v29) >= 2 {
                        v29 = v29[1:len(v29)]
                    }
                    _ = v29
                }
                _ = fi10
            }
            var v30 int8 = 100
            for fi31 := 0; fi31 < 8; fi31++ {
                fmt.Println(v4, fi31, fi9)
                v4 = v4
                v2 = v2
                _ = fi31
            }
            _ = fi9
            _ = v30
        }
    case 2147483647:
        v32 := GenS2{F0: 980, F1: 742}
        _ = v32
    }
    var v33 uint = (v1 ^ 1)
    v5 = (v5 | v5)
    v34 := GenS5{F0: v3, F1: v5}
    v33 = (v33 - 1)
    if true {
        switch "foo bar" {
        case "a":
            for fi35 := 0; fi35 < 4; fi35++ {
                v33 = (v1 & v1)
                _ = fi35
            }
        }
        v36 := GenS4{F0: (("hello" + v2) + "a")}
        var v37 int16 = 1
        _ = v36
        _ = v37
    }
    v4 += 480
    if false {
        v38 := GenS1{F0: 1, F1: 1}
        v2 = (v3 + (v3 + "hello"))
        if false {
            v33 -= (v1 * v33)
            if true {
                for fi39 := 0; fi39 < 16; fi39++ {
                    v40 := GenS3{F0: 901}
                    _ = fi39
                    _ = v40
                }
            } else {
                var v41 uint64 = 42
                _ = v41
            }
        }
        _ = v38
    } else {
        for fi42 := 0; fi42 < 10; fi42++ {
            v43 := []int{(^(fi42)), (fi42 | (fi42 / ((fi42) | 1))), (fi42 << (((^(v33))) & 63)), (fi42 % ((fi42) | 1))}
            _ = fi42
            _ = v43
        }
    }
    fmt.Println(v1, v33)
    var v44 uint64 = 18446744073709551615
    switch v2 {
    case "":
        if true {
            switch "a" {
            case "a":
                var v45 uint = v1
                _ = v45
            default:
                if (-645 != (v4 >> (((v33 & (v1 | 23))) & 63))) {
                    v46 := GenS3{F0: 65535}
                    switch 127 {
                    case 344:
                        if (371 <= 1) {
                            fmt.Println(v5)
                            v4 += -1
                            var v47 uint16 = 65535
                            _ = v47
                        } else {
                            v48 := []int{0, 9223372036854775807, 264}
                            fmt.Println(v44, v3, v2)
                            if len(v48) >= 2 {
                                v48 = v48[1:len(v48)]
                            }
                            _ = v48
                        }
                    case 128:
                        v3 = (("a" + v2) + (v2 + (v2 + "foo bar")))
                    default:
                        fmt.Println(v44, v33)
                    }
                    _ = v46
                }
            }
            switch -550 {
            case -400:
                fmt.Println(v4, v1)
            case 32767:
                v49 := GenS6{F0: v4, F1: 903, F2: 32768}
                _ = v49
            default:
                if false {
                    fmt.Println(v33, v1, v4)
                    v4 -= v4
                }
            }
            fmt.Println(v5, v1, v1)
        } else {
            if (0.0 > 3.14159) {
                if true {
                    for fi50 := 0; fi50 < 6; fi50++ {
                        var v51 int64 = (v4 % (((v4 * v4)) | 1))
                        _ = fi50
                        _ = v51
                    }
                    if ((1.5 > 1e10) || true) {
                        fmt.Println(v5, v5, v3)
                        var v52 int32 = -2147483648
                        v53 := []int{32767, 32768}
                        _ = v52
                        _ = v53
                    }
                    fmt.Println(v3, v33)
                } else {
                    var v54 float32 = 1.5
                    switch (v2 + "foo bar") {
                    case "hello":
                        if ((!true) || false) {
                            if true {
                                v55 := []int{740}
                                fmt.Println(v4)
                                switch int(v33) {
                                case 42:
                                    for fi56 := 0; fi56 < 20; fi56++ {
                                        v4 = (v4 - 0)
                                        if false {
                                            var v57 uint32 = 617
                                            v3 = v3
                                            _ = v57
                                        } else {
                                            var v58 int = fi56
                                            v4 += (v4 ^ (^(v4)))
                                            v59 := GenS6{F0: -1, F1: 601, F2: (v58 - 604)}
                                            _ = v58
                                            _ = v59
                                        }
                                        v33 = (v33 ^ (v33 + v33))
                                        _ = fi56
                                    }
                                case -9223372036854775808:
                                    v54 = ((v54 - (v54 * 1.5)) + v54)
                                case 184:
                                    if true {
                                        fmt.Println(v44, v3)
                                        v33 = 0
                                    } else {
                                        var v60 int64 = (v4 / (((v4 ^ v4)) | 1))
                                        fmt.Println(v5)
                                        if len(v55) >= 2 {
                                            v55 = v55[1:len(v55)]
                                        }
                                        _ = v60
                                    }
                                }
                                _ = v55
                            } else {
                                v61 := GenS7{F0: (v3 + "\n\t"), F1: int16(v33)}
                                v62 := []int{823}
                                switch 850 {
                                case -368:
                                    var v63 int32 = 42
                                    _ = v63
                                case -9223372036854775808:
                                    var v64 int = 1
                                    _ = v64
                                case -417:
                                    if (false || true) {
                                        for fi65 := 0; fi65 < 5; fi65++ {
                                            v4 += v4
                                            _ = fi65
                                        }
                                    }
                                default:
                                    if (float64(0.1) == 3.14159) {
                                        v33 -= 2
                                        var v66 int64 = int64(v1)
                                        _ = v66
                                    }
                                }
                                _ = v61
                                _ = v62
                            }
                            var v67 uint32 = v5
                            if ((v54 + (v54 / float32(((((1e10 - ((-2.25 / 1.0) + float64(1.5))) / 1.0) + -2.25) * float64(v54))))) <= 1.5) {
                                v4 = (v4 | (v4 << ((122) & 63)))
                                if (!(false || false)) {
                                    var v68 uint = v1
                                    v69 := GenS4{F0: (v2 + (((((v3 + ((v3 + v3) + "")) + (v2 + v3)) + (("hello" + "foo bar") + "a")) + ((("\n\t" + v3) + (v2 + v2)) + v2)) + v2))}
                                    _ = v68
                                    _ = v69
                                } else {
                                    fmt.Println(v3, v67)
                                    var v70 int64 = 9223372036854775807
                                    _ = v70
                                }
                                v44 -= v44
                            } else {
                                if true {
                                    switch "foo bar" {
                                    case "\n\t":
                                        for fi71 := 0; fi71 < 4; fi71++ {
                                            v72 := GenS4{F0: v3}
                                            v73 := GenS2{F0: int32(v44), F1: 4294967295}
                                            var v74 uint8 = 2
                                            _ = fi71
                                            _ = v72
                                            _ = v73
                                            _ = v74
                                        }
                                    }
                                    fmt.Println(v5)
                                } else {
                                    var v75 string = v2
                                    for fi76 := 0; fi76 < 7; fi76++ {
                                        v77 := GenS1{F0: 42, F1: 878}
                                        v54 += float32(float64((v54 / 1.0)))
                                        _ = fi76
                                        _ = v77
                                    }
                                    v78 := []int{-319, -32}
                                    _ = v75
                                    _ = v78
                                }
                            }
                            _ = v67
                        } else {
                            var v79 int16 = 1
                            for fi80 := 0; fi80 < 2; fi80++ {
                                v44 += 1
                                _ = fi80
                            }
                            fmt.Println(v54, v33, v79)
                            _ = v79
                        }
                    case "foo bar":
                        fmt.Println(v2, v4, v54)
                    case "\n\t":
                        v44 = v44
                    default:
                        if true {
                            v3 = (("a" + ("foo bar" + ("hello" + v2))) + ("a" + (v3 + ((v2 + ("" + ("a" + v3))) + (v3 + v3)))))
                            v81 := GenS8{F0: v5, F1: 0, F2: 1}
                            var v82 float64 = float64(1.5)
                            _ = v81
                            _ = v82
                        } else {
                            v83 := []int{128}
                            var v84 uint8 = uint8(v5)
                            _ = v83
                            _ = v84
                        }
                    }
                    _ = v54
                }
                v85 := []int{2, 9223372036854775807, -9223372036854775808}
                fmt.Println(v33, v33)
                _ = v85
            }
            v5 += 682
            v86 := GenS7{F0: v2, F1: int16(v44)}
            _ = v86
        }
    case "hello":
        v87 := []int{2147483647}
        _ = v87
    default:
        var v88 int32 = 0
        _ = v88
    }
    for fi89 := 0; fi89 < 7; fi89++ {
        if true {
            var v90 bool = (2147483647 >= -2147483648)
            fmt.Println(v33, v3)
            _ = v90
        } else {
            v91 := []int{fi89, (fi89 + fi89), (fi89 * fi89), 127}
            var v92 int = (fi89 & (fi89 - fi89))
            _ = v91
            _ = v92
        }
        v3 = v3
        v1 = (v33 - (v33 | 1))
        _ = fi89
    }
    for fi93 := 0; fi93 < 16; fi93++ {
        v94 := []int{660, fi93, fi93}
        var v95 uint16 = 1
        switch v3 {
        case "\n\t":
            v33 += v1
        }
        _ = fi93
        _ = v94
        _ = v95
    }
    switch 0 {
    case 32768:
        fmt.Println(v44, v3)
    case 128:
        switch -2147483648 {
        case 9223372036854775807:
            v4 += v4
        case -410:
            v96 := []int{517, 32768, -1, 32768}
            _ = v96
        case 193:
            v97 := GenS9{F0: v4, F1: -2147483648, F2: 713}
            _ = v97
        default:
            for fi98 := 0; fi98 < 3; fi98++ {
                fmt.Println(v4, fi98, v44)
                fmt.Println(v2)
                switch v3 {
                case "\n\t":
                    if ((18446744073709551615 < v44) && (1.5 > ((float64((3.14159 / 1.0)) * -2.25) + (float64(1.5) - 0.1)))) {
                        var v99 int64 = (v4 << ((1) & 63))
                        _ = v99
                    } else {
                        var v100 float64 = 0.1
                        for fi101 := 0; fi101 < 11; fi101++ {
                            fmt.Println(v4)
                            v102 := []int{(fi101 / ((fi101) | 1)), (^(fi98))}
                            _ = fi101
                            _ = v102
                        }
                        _ = v100
                    }
                case "hello":
                    v4 -= 0
                }
                _ = fi98
            }
        }
    case 108:
        var v103 bool = false
        _ = v103
    default:
        v104 := GenS10{F0: 139, F1: 255}
        _ = v104
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v44
    _ = v5
    _ = v6
}
