package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 string; F1 uint; F2 uint8 }
type GenS3 struct { F0 uint32; F1 int32 }
type GenS4 struct { F0 uint64; F1 uint16; F2 uint64 }
func genFn0() uint64 {
    var v4 bool = true
    if true {
        v5 := GenS1{F0: (-20 > -128)}
        var v6 int8 = -1
        var v7 int64 = 134
        _ = v5
        _ = v6
        _ = v7
    }
    var v8 uint32 = 1
    for fi9 := 0; fi9 < 2; fi9++ {
        var v10 uint16 = 42
        _ = fi9
        _ = v10
    }
    _ = v4
    _ = v8
    return 4294967295
}

func main() {
    v1 := make(map[int]string)
    var v2 bool = false
    var v3 uint64 = 920
    var v11 int8 = 118
    v12 := GenS2{F0: (("hello" + "") + "a"), F1: 255, F2: 128}
    fmt.Println(v2, v2, v3)
    var v13 int8 = (-(v11))
    if true {
        switch "" {
        case "":
            for fi14 := 0; fi14 < 2; fi14++ {
                v1[(fi14 % (((fi14 ^ 127)) | 1))] = ""
                _ = fi14
            }
        case "a":
            v15 := GenS1{F0: (!(32767 > 255))}
            _ = v15
        case "hello":
            v11 = 0
        default:
            var v16 bool = v2
            _ = v16
        }
        v1[32767] = "\n\t"
    } else {
        var v17 uint = 255
        v11 += (-(v13))
        for fi18 := 0; fi18 < 18; fi18++ {
            switch 42 {
            case 4294967295:
                if ((1 > uint32(v17)) || (2 != 0)) {
                    var v19 uint16 = 255
                    _ = v19
                }
            default:
                for fi20 := 0; fi20 < 12; fi20++ {
                    delete(v1, 547)
                    _ = fi20
                }
            }
            if (0.1 >= (1.5 / 3.14159)) {
                var v21 int64 = -1
                _ = v21
            } else {
                v13 = int8(v17)
                if v2 {
                    if v2 {
                        v3 -= genFn0()
                        if (0.1 > (1e10 / 1.5)) {
                            for fi22 := 0; fi22 < 4; fi22++ {
                                v23 := GenS2{F0: ("a" + ("a" + ("a" + ("" + "foo bar")))), F1: (^(v17)), F2: 2}
                                var v24 bool = (42 <= 47)
                                var v25 int = fi18
                                _ = fi22
                                _ = v23
                                _ = v24
                                _ = v25
                            }
                            v13 -= (v11 >> ((v17) & 63))
                            v3 = v3
                        }
                    } else {
                        v13 += (v11 - (v13 - (v13 >> (((v17 & v17)) & 63))))
                        var v26 uint64 = v3
                        _ = v26
                    }
                    for fi27 := 0; fi27 < 6; fi27++ {
                        if v2 {
                            if (!true) {
                                if (((65535 >= 42) || true) && v2) {
                                    v2 = (false && v2)
                                    v3 = v3
                                    fmt.Println(v17, v2, v3)
                                } else {
                                    switch "a" {
                                    case "":
                                        var v28 float32 = 3.14159
                                        _ = v28
                                    default:
                                        var v29 string = "hello"
                                        _ = v29
                                    }
                                    var v30 int64 = -1
                                    _ = v30
                                }
                                switch "\n\t" {
                                case "hello":
                                    v31 := GenS2{F0: "a", F1: 2, F2: 2}
                                    _ = v31
                                case "":
                                    v3 -= (v3 / (((v3 % (((v3 % ((0) | 1))) | 1))) | 1))
                                case "a":
                                    var v32 int16 = 32767
                                    _ = v32
                                }
                            } else {
                                _ = len(v1)
                            }
                            if v2 {
                                if (-32768 != 42) {
                                    v33 := GenS3{F0: 2147483648, F1: int32(v11)}
                                    if v2 {
                                        v34 := GenS3{F0: 353, F1: -2147483648}
                                        _ = v34
                                    }
                                    v35 := GenS2{F0: "\n\t", F1: (v17 % (((v17 * (^(v17)))) | 1)), F2: 42}
                                    _ = v33
                                    _ = v35
                                }
                            } else {
                                var v36 bool = v2
                                if (!true) {
                                    v36 = (864 >= 42)
                                    var v37 int8 = 1
                                    _ = v37
                                }
                                switch fi18 {
                                case 1:
                                    var v38 uint16 = 2
                                    v13 += v11
                                    var v39 int64 = 640
                                    _ = v38
                                    _ = v39
                                case 128:
                                    v36 = (!(!(v36 || v2)))
                                default:
                                    v13 += v13
                                }
                                _ = v36
                            }
                            v3 = (v3 << (((v17 / (((v17 - v17)) | 1))) & 63))
                        } else {
                            v40 := []int{fi18, 2147483647, (fi18 | 32768)}
                            switch ("\n\t" + "foo bar") {
                            case "":
                                var v41 int = (^(fi27))
                                _ = v41
                            default:
                                v40 = append(v40, fi27)
                            }
                            _ = v40
                        }
                        if true {
                            var v42 float64 = 1.5
                            _ = v42
                        }
                        _ = fi27
                    }
                }
            }
            _ = fi18
        }
        _ = v17
    }
    v43 := v1[0]
    if ((-2147483648 != 42) && (v2 && v2)) {
        switch v43 {
        case "a":
            if (1.5 < -2.25) {
                if (65535 != 42) {
                    v44 := GenS4{F0: v3, F1: 65535, F2: (^(v3))}
                    _ = v44
                } else {
                    if false {
                        fmt.Println(v3, v13, v2)
                        v45 := []int{672, 129}
                        _ = v45
                    } else {
                        for fi46 := 0; fi46 < 4; fi46++ {
                            v11 -= (^(v11))
                            v3 -= 355
                            _ = fi46
                        }
                        if (("\n\t" + (v43 + v43)) > ((v43 + v43) + v43)) {
                            if ((((127 != -1) && (2147483648 <= 1)) && (1.5 <= 0.1)) || v2) {
                                switch v43 {
                                case "hello":
                                    delete(v1, 32767)
                                default:
                                    v43 = v43
                                }
                                if (431 > 0) {
                                    v47 := GenS3{F0: 821, F1: -838}
                                    _ = v47
                                }
                            } else {
                                if (0 > 2) {
                                    fmt.Println(v11)
                                    v3 = v3
                                } else {
                                    if true {
                                        switch "a" {
                                        case "\n\t":
                                            switch v43 {
                                            case "":
                                                v1[1] = ((v43 + v43) + v43)
                                            case "hello":
                                                if v2 {
                                                    v48 := []int{42, 127, 42}
                                                    switch v43 {
                                                    case "":
                                                        v2 = true
                                                    default:
                                                        var v49 int64 = -285
                                                        _ = v49
                                                    }
                                                    _ = v48
                                                } else {
                                                    if ((v13 + 0) > 127) {
                                                        v43 = v43
                                                    } else {
                                                        for fi50 := 0; fi50 < 13; fi50++ {
                                                            if false {
                                                                v1[fi50] = ("" + v43)
                                                                var v51 uint = 2
                                                                if (1e10 != -2.25) {
                                                                    v52 := []int{(fi50 * 361), fi50, fi50}
                                                                    _ = v52
                                                                }
                                                                _ = v51
                                                            }
                                                            switch (("" + "") + v43) {
                                                            case "\n\t":
                                                                fmt.Println(v2, v13, v2)
                                                            case "":
                                                                v2 = (true && (-1 <= -2147483648))
                                                            case "foo bar":
                                                                v53 := GenS4{F0: 42, F1: 65535, F2: (v3 % ((genFn0()) | 1))}
                                                                _ = v53
                                                            }
                                                            v43 = "foo bar"
                                                            _ = fi50
                                                        }
                                                    }
                                                    fmt.Println(v43, v43)
                                                }
                                            case "foo bar":
                                                var v54 uint8 = 42
                                                _ = v54
                                            default:
                                                for fi55 := 0; fi55 < 11; fi55++ {
                                                    var v56 uint64 = (v3 / (((v3 % ((v3) | 1))) | 1))
                                                    var v57 int8 = (-(v13))
                                                    _ = fi55
                                                    _ = v56
                                                    _ = v57
                                                }
                                            }
                                        }
                                    }
                                }
                            }
                            for fi58 := 0; fi58 < 11; fi58++ {
                                switch ((v43 + v43) + v43) {
                                case "hello":
                                    var v59 bool = ((v2 || (2147483647 <= 42)) || (1000 <= 0))
                                    _ = v59
                                }
                                _ = fi58
                            }
                            v3 -= v3
                        }
                    }
                    for fi60 := 0; fi60 < 15; fi60++ {
                        for fi61 := 0; fi61 < 1; fi61++ {
                            v11 -= v11
                            if (v2 && (-9223372036854775808 < -1)) {
                                v62 := []int{-552}
                                var v63 uint = 42
                                v2 = v2
                                _ = v62
                                _ = v63
                            }
                            _ = fi61
                        }
                        v64 := GenS4{F0: v3, F1: 65535, F2: 2}
                        _ = fi60
                        _ = v64
                    }
                }
                var v65 float64 = 1.5
                _ = v65
            }
        case "":
            v2 = v2
        case "\n\t":
            var v66 int32 = 762
            _ = v66
        default:
            for fi67 := 0; fi67 < 18; fi67++ {
                v68 := GenS4{F0: (v3 | v3), F1: 875, F2: 42}
                v2 = (-1 < 0)
                _ = fi67
                _ = v68
            }
        }
    }
    if v2 {
        if false {
            switch 1 {
            case 95:
                if false {
                    var v69 int32 = -649
                    var v70 float32 = 1.5
                    _ = v69
                    _ = v70
                }
            case 127:
                v3 += v3
            default:
                v43 = (v43 + v43)
            }
        } else {
            switch -335 {
            case -977:
                v71 := GenS3{F0: 42, F1: 0}
                _ = v71
            case 4294967295:
                if false {
                    var v72 uint = 255
                    _ = v72
                } else {
                    var v73 uint64 = v3
                    _ = v73
                }
            case -9223372036854775808:
                var v74 int64 = 630
                _ = v74
            }
        }
        var v75 uint16 = 2
        _ = v75
    } else {
        var v76 int = -115
        v77 := []int{9223372036854775807, (v76 / ((-562) | 1)), 227}
        v78 := v1[(v76 - v76)]
        _ = v76
        _ = v77
        _ = v78
    }
    v79 := GenS2{F0: "foo bar", F1: 99, F2: 1}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v43
    _ = v79
}
