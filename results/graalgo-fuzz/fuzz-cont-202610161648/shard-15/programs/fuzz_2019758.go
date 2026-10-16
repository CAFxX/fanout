package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8; F2 int16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint; F1 int; F2 bool }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint32; F1 bool; F2 bool }
type GenS6 struct { F0 int16; F1 int; F2 bool }

func main() {
    var v1 int8 = 27
    var v2 int16 = -32768
    var v3 int16 = 42
    v4 := GenS1{F0: (-(v1)), F1: 128, F2: -1}
    if true {
        v5 := GenS1{F0: (-(v1)), F1: 231, F2: (^(v2))}
        if false {
            v6 := []int{9223372036854775807, 128, -9223372036854775808}
            v1 += (-(v1))
            for fi7 := 0; fi7 < 13; fi7++ {
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                for fi8 := 0; fi8 < 12; fi8++ {
                    if true {
                        v9 := GenS2{F0: -1}
                        fmt.Println(v2, fi8, v1)
                        _ = v9
                    } else {
                        var v10 int16 = v2
                        v1 = v1
                        _ = v10
                    }
                    fmt.Println(fi7, fi8)
                    _ = fi8
                }
                _ = fi7
            }
            _ = v6
        } else {
            var v11 int32 = 410
            for fi12 := 0; fi12 < 2; fi12++ {
                switch ("" + "a") {
                case "":
                    fmt.Println(v1)
                }
                _ = fi12
            }
            var v13 uint32 = 1
            _ = v11
            _ = v13
        }
        _ = v5
    }
    for fi14 := 0; fi14 < 11; fi14++ {
        v15 := GenS3{F0: 18446744073709551615, F1: (fi14 >> ((221) & 63)), F2: false}
        _ = fi14
        _ = v15
    }
    v3 = 42
    if true {
        var v16 uint = 255
        var v17 float32 = 0.0
        for fi18 := 0; fi18 < 16; fi18++ {
            if true {
                v19 := []int{fi18}
                v20 := GenS2{F0: -201}
                _ = v19
                _ = v20
            }
            _ = fi18
        }
        _ = v16
        _ = v17
    } else {
        v3 = -25
        fmt.Println(v1, v3, v3)
        v21 := GenS4{F0: -55}
        _ = v21
    }
    if (26 > 0) {
        v22 := GenS3{F0: 160, F1: 128, F2: ((^(v1)) >= v1)}
        if true {
            v2 -= v3
            fmt.Println(v2, v1, v1)
        } else {
            switch ((("a" + "") + "\n\t") + "hello") {
            case "foo bar":
                var v23 uint16 = 2
                _ = v23
            case "":
                v24 := []int{127, -775}
                _ = v24
            case "a":
                if ((1e10 / (3.14159 / -2.25)) <= (float32(1.5) / ((0.0 - 1e10) * (1e10 - float32(((0.1 + (-2.25 / 3.14159)) * 0.1)))))) {
                    var v25 uint64 = 42
                    _ = v25
                }
            }
            for fi26 := 0; fi26 < 20; fi26++ {
                v3 -= -406
                for fi27 := 0; fi27 < 15; fi27++ {
                    if false {
                        var v28 int16 = 0
                        _ = v28
                    } else {
                        fmt.Println(v1)
                        if true {
                            var v29 int16 = 798
                            _ = v29
                        } else {
                            v3 = v3
                            v30 := []int{}
                            if ((0.1 * 0.1) != 3.14159) {
                                if (-118 == 0) {
                                    _ = len(v30)
                                } else {
                                    if false {
                                        var v31 bool = false
                                        v32 := GenS1{F0: (v1 + v1), F1: 127, F2: (v3 + v2)}
                                        switch fi26 {
                                        case 128:
                                            if true {
                                                var v33 uint16 = 1
                                                _ = v33
                                            }
                                        case 0:
                                            if ((("\n\t" + "\n\t") > ("hello" + "foo bar")) && false) {
                                                if v31 {
                                                    var v34 bool = (v31 || (v3 != 1000))
                                                    _ = v34
                                                } else {
                                                    if ((42 >= 2147483647) || v31) {
                                                        var v35 bool = v31
                                                        if (-2.25 <= 3.14159) {
                                                            var v36 uint32 = 2147483648
                                                            _ = v36
                                                        }
                                                        _ = v35
                                                    }
                                                }
                                            } else {
                                                if len(v30) > 0 {
                                                    v30[0] = (fi27 / ((fi27) | 1))
                                                }
                                                v2 += v2
                                                var v37 uint8 = 127
                                                _ = v37
                                            }
                                        case 922:
                                            v31 = (4294967295 > 2)
                                        default:
                                            switch "a" {
                                            case "foo bar":
                                                var v38 float64 = 3.14159
                                                _ = v38
                                            case "\n\t":
                                                var v39 int16 = int16(v1)
                                                _ = v39
                                            case "":
                                                if (true && (-9223372036854775808 <= -9223372036854775808)) {
                                                    v40 := GenS5{F0: 42, F1: (!(255 == 32767)), F2: (!(32767 != 2))}
                                                    _ = v40
                                                }
                                            }
                                        }
                                        _ = v31
                                        _ = v32
                                    } else {
                                        v2 = v3
                                        if true {
                                            var v41 float64 = ((0.1 / 1.0) * (((3.14159 - 1e10) - 0.1) - 3.14159))
                                            if false {
                                                var v42 int8 = (v1 - v1)
                                                v1 = (^(v1))
                                                _ = v42
                                            }
                                            v30 = append(v30, fi27)
                                            _ = v41
                                        }
                                        fmt.Println(v2, fi26, fi27)
                                    }
                                }
                                if ((0.0 / float64(float32((1e10 / float64(1.5))))) <= 3.14159) {
                                    fmt.Println(v1, v3)
                                    v3 += (v3 % ((1000) | 1))
                                }
                            }
                            _ = v30
                        }
                    }
                    v2 = (v3 | (v3 >> ((576) & 63)))
                    var v43 int8 = -1
                    _ = fi27
                    _ = v43
                }
                fmt.Println(fi26, v2)
                _ = fi26
            }
        }
        _ = v22
    }
    if true {
        v1 = (v1 ^ (^(v1)))
    } else {
        for fi44 := 0; fi44 < 12; fi44++ {
            v45 := GenS6{F0: v2, F1: fi44, F2: (432 < 4294967295)}
            _ = fi44
            _ = v45
        }
        v46 := GenS5{F0: 65535, F1: true, F2: (2 > 349)}
        _ = v46
    }
    if false {
        v1 = v1
        v1 -= 42
    } else {
        for fi47 := 0; fi47 < 18; fi47++ {
            if (!true) {
                fmt.Println(v2, v3)
            }
            var v48 bool = true
            _ = fi47
            _ = v48
        }
    }
    if true {
        v1 -= 0
        for fi49 := 0; fi49 < 16; fi49++ {
            var v50 int8 = (v1 ^ (v1 ^ v1))
            _ = fi49
            _ = v50
        }
        var v51 uint64 = 42
        _ = v51
    } else {
        v52 := GenS5{F0: 20, F1: (663 != 42), F2: false}
        _ = v52
    }
    var v53 float32 = 0.1
    var v54 int16 = v3
    if true {
        var v55 uint64 = 0
        switch -2147483648 {
        case 127:
            var v56 uint = 255
            _ = v56
        case 42:
            for fi57 := 0; fi57 < 16; fi57++ {
                var v58 uint8 = 128
                fmt.Println(v1)
                v59 := []int{(fi57 >> ((2) & 63)), (fi57 | 127), int(v1)}
                _ = fi57
                _ = v58
                _ = v59
            }
        }
        _ = v55
    }
    fmt.Println(v3, v3)
    switch -269 {
    case 537:
        v54 = v2
    default:
        if (!(42 < 42)) {
            v53 -= v53
            if true {
                if (!true) {
                    fmt.Println(v2, v54, v54)
                    fmt.Println(v1, v54, v3)
                    var v60 uint8 = 128
                    _ = v60
                } else {
                    switch ("a" + "foo bar") {
                    case "a":
                        for fi61 := 0; fi61 < 3; fi61++ {
                            v62 := []int{128, (fi61 & 710), 0, (^(fi61))}
                            var v63 uint = 4294967295
                            _ = fi61
                            _ = v62
                            _ = v63
                        }
                    case "":
                        v64 := GenS4{F0: -787}
                        _ = v64
                    case "foo bar":
                        var v65 float32 = (((0.1 / 1.0) * float32((0.1 - float64((float32(1e10) * v53))))) / (v53 * v53))
                        _ = v65
                    default:
                        v66 := GenS4{F0: 764}
                        _ = v66
                    }
                }
                for fi67 := 0; fi67 < 18; fi67++ {
                    v68 := GenS6{F0: v3, F1: (fi67 * (fi67 * 797)), F2: false}
                    v69 := GenS6{F0: v3, F1: (fi67 / (((fi67 * 306)) | 1)), F2: true}
                    for fi70 := 0; fi70 < 12; fi70++ {
                        var v71 uint16 = 285
                        var v72 uint8 = 211
                        v3 = v2
                        _ = fi70
                        _ = v71
                        _ = v72
                    }
                    _ = fi67
                    _ = v68
                    _ = v69
                }
            }
            for fi73 := 0; fi73 < 18; fi73++ {
                for fi74 := 0; fi74 < 18; fi74++ {
                    fmt.Println(fi74)
                    if ((-(v1)) == 88) {
                        if (4294967295 >= 65535) {
                            var v75 string = "hello"
                            _ = v75
                        }
                        var v76 int16 = v3
                        _ = v76
                    }
                    switch fi74 {
                    case -2147483648:
                        var v77 bool = (!(v53 <= (v53 / 3.14159)))
                        _ = v77
                    case 32767:
                        var v78 int32 = 42
                        _ = v78
                    case -9223372036854775808:
                        var v79 int64 = 405
                        _ = v79
                    }
                    _ = fi74
                }
                v53 -= (v53 * (v53 / v53))
                switch "hello" {
                case "a":
                    v80 := GenS6{F0: (v54 * (v2 * v2)), F1: fi73, F2: (false && (!true))}
                    _ = v80
                case "\n\t":
                    v81 := GenS5{F0: 931, F1: ((1.5 >= 0.0) || true), F2: (2 < 4294967295)}
                    _ = v81
                case "foo bar":
                    var v82 float32 = ((-2.25 / ((float32((1.5 / 1.5)) * 0.0) - 0.1)) - v53)
                    _ = v82
                }
                _ = fi73
            }
        }
    }
    v2 = 1000
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v53:", v53)
    fmt.Println("v54:", v54)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v53
    _ = v54
}
