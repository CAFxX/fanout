package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 string; F1 uint16; F2 int32 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint16; F1 uint32; F2 int32 }
func genFn0(p0 int8, p1 bool) uint32 {
    var v4 uint = 0
    p0 = 127
    for fi5 := 0; fi5 < 17; fi5++ {
        var v6 string = ("\n\t" + "hello")
        v7 := GenS1{F0: 229}
        _ = fi5
        _ = v6
        _ = v7
    }
    for fi8 := 0; fi8 < 12; fi8++ {
        v9 := GenS2{F0: "", F1: 42, F2: -1}
        v10 := []int{(fi8 << (((v4 + 1)) & 63)), 128, 127, (fi8 % ((fi8) | 1))}
        _ = fi8
        _ = v10
        _ = v9
    }
    var v11 int8 = (^(p0))
    _ = p0
    _ = p1
    _ = v11
    _ = v4
    return uint32(v4)
}
func genFn1(p0 uint32, p1 bool, p2 uint64) int8 {
    var v12 int16 = 42
    p0 += (^(p0))
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return -1
}

func main() {
    v1 := []uint32{4294967295}
    var v2 bool = ((307 >= 65535) || (!true))
    var v3 int = 127
    for fi13 := 0; fi13 < 15; fi13++ {
        switch (v3 + (v3 | (fi13 % ((fi13) | 1)))) {
        case 128:
            v2 = (4294967295 != 18446744073709551615)
        case -567:
            v3 += fi13
        case 32768:
            v2 = false
        default:
            v3 -= 32768
        }
        _ = fi13
    }
    if (551 <= 326) {
        for fi14 := 0; fi14 < 2; fi14++ {
            if false {
                v15 := GenS1{F0: 185}
                if len(v1) > 0 {
                    v1[0] = 363
                }
                var v16 float64 = (1e10 + float64(float32((0.0 * 0.0))))
                _ = v15
                _ = v16
            } else {
                var v17 uint = 0
                for fi18 := 0; fi18 < 7; fi18++ {
                    if ((v3 + -140) > (fi18 / ((393) | 1))) {
                        switch "" {
                        case "hello":
                            var v19 float64 = ((float64(1e10) - 1e10) + 1.5)
                            _ = v19
                        case "a":
                            v3 -= (fi18 | (fi14 >> ((v17) & 63)))
                        case "\n\t":
                            v17 = v17
                        }
                        if v2 {
                            fmt.Println(v17, fi14)
                        } else {
                            var v20 uint16 = 391
                            _ = v20
                        }
                        v21 := GenS3{F0: 1}
                        _ = v21
                    } else {
                        v2 = (!(56 > 0))
                        var v22 int16 = 1
                        var v23 int32 = 42
                        var v24 uint64 = 42
                        _ = v22
                        _ = v23
                        _ = v24
                    }
                    _ = fi18
                }
                if v2 {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    var v25 int64 = 42
                    _ = v25
                }
                _ = v17
            }
            _ = fi14
        }
    } else {
        fmt.Println(v3)
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if ((0 == 0) || (385 <= -1)) {
        for fi26 := 0; fi26 < 15; fi26++ {
            v27 := GenS3{F0: int8(fi26)}
            var v28 int64 = 1
            _ = fi26
            _ = v27
            _ = v28
        }
        for fi29 := 0; fi29 < 17; fi29++ {
            v2 = ((247 >= 2147483647) || v2)
            if (862 < 65535) {
                if v2 {
                    fmt.Println(fi29, v3)
                    var v30 int = 32768
                    _ = v30
                }
                v31 := GenS2{F0: "", F1: 1, F2: -792}
                var v32 int64 = -9223372036854775808
                _ = v31
                _ = v32
            } else {
                if len(v1) > 0 {
                    v1[0] = 2
                }
                fmt.Println(v2, v3)
            }
            _ = fi29
        }
        switch "" {
        case "\n\t":
            var v33 float64 = 0.0
            _ = v33
        }
    } else {
        for fi34 := 0; fi34 < 5; fi34++ {
            if v2 {
                var v35 int8 = 0
                switch ("" + (("a" + "a") + "")) {
                case "hello":
                    v36 := GenS2{F0: (((("a" + "hello") + "") + "foo bar") + "hello"), F1: 0, F2: -2147483648}
                    _ = v36
                case "":
                    var v37 float64 = 1e10
                    _ = v37
                default:
                    for fi38 := 0; fi38 < 1; fi38++ {
                        var v39 int = -832
                        fmt.Println(fi38, v3, v35)
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        _ = fi38
                        _ = v39
                    }
                }
                _ = v35
            } else {
                v1 = append(v1, 308)
            }
            _ = fi34
        }
    }
    switch (v3 - 42) {
    case 127:
        var v40 int32 = -2147483648
        _ = v40
    }
    var v41 bool = (("foo bar" + "\n\t") < "\n\t")
    var v42 float32 = ((((3.14159 - ((float32((0.0 / 1e10)) + ((0.1 / -2.25) / 1e10)) / float32(3.14159))) * float32(0.1)) / 1.0) / (1e10 / 0.1))
    var v43 uint32 = 522
    v3 = v3
    switch "\n\t" {
    case "":
        v44 := GenS3{F0: genFn1(0, false, 4294967295)}
        _ = v44
    case "\n\t":
        switch -1 {
        case 42:
            v42 = ((v42 + (((v42 / v42) / 1.0) - v42)) - float32(1e10))
        case 2147483647:
            for fi45 := 0; fi45 < 5; fi45++ {
                if (v2 && v41) {
                    if v2 {
                        v3 -= (v3 % ((v3) | 1))
                    } else {
                        if false {
                            if (((-1 != 707) && (v3 < v3)) || ((v42 * (v42 * 0.1)) == (v42 * 0.1))) {
                                var v46 float32 = 3.14159
                                fmt.Println(fi45)
                                _ = v46
                            } else {
                                var v47 int8 = int8(v43)
                                if (!false) {
                                    if v2 {
                                        v42 = (0.1 / v42)
                                        var v48 float32 = v42
                                        if ((1 != 0) || (true && false)) {
                                            var v49 uint64 = 640
                                            if false {
                                                var v50 uint = 1
                                                fmt.Println(v42, v48, fi45)
                                                _ = v50
                                            } else {
                                                switch 9223372036854775807 {
                                                case 128:
                                                    v42 = v42
                                                case 397:
                                                    if v41 {
                                                        var v51 uint16 = 0
                                                        _ = v51
                                                    } else {
                                                        var v52 int16 = -892
                                                        fmt.Println(v47)
                                                        _ = v52
                                                    }
                                                }
                                            }
                                            v47 -= genFn1(v43, false, 42)
                                            _ = v49
                                        } else {
                                            var v53 float64 = 3.14159
                                            if len(v1) > 0 {
                                                v1[0] = v43
                                            }
                                            _ = v53
                                        }
                                        _ = v48
                                    }
                                    fmt.Println(v43, v41, v42)
                                    fmt.Println(v41)
                                } else {
                                    v41 = ("a" < "")
                                }
                                _ = v47
                            }
                        }
                    }
                    v41 = true
                    switch ("a" + "") {
                    case "hello":
                        for fi54 := 0; fi54 < 13; fi54++ {
                            if false {
                                switch "hello" {
                                case "hello":
                                    switch fi45 {
                                    case 13:
                                        fmt.Println(fi45)
                                    case -374:
                                        v41 = ((v43 >> ((18446744073709551615) & 63)) < v43)
                                    case 32768:
                                        if v41 {
                                            if (false || v41) {
                                                var v55 uint8 = 197
                                                _ = v55
                                            }
                                            fmt.Println(fi45, v3, fi45)
                                            if len(v1) > 0 {
                                                v1[0] = v43
                                            }
                                            _ = len(v1)
                                        } else {
                                            fmt.Println(v42, v41, v42)
                                        }
                                    default:
                                        var v56 int8 = 27
                                        _ = v56
                                    }
                                case "\n\t":
                                    v41 = v41
                                }
                                var v57 int16 = 32767
                                _ = len(v1)
                                if (!(v41 || (!v2))) {
                                    switch "\n\t" {
                                    case "hello":
                                        v57 -= v57
                                    }
                                    if v2 {
                                        var v58 float64 = (3.14159 - (-2.25 - ((1.5 * (-2.25 / (float64(v42) / 0.1))) + float64(v42))))
                                        _ = v58
                                    }
                                    v59 := GenS3{F0: -128}
                                    _ = v59
                                } else {
                                    switch -9223372036854775808 {
                                    case 42:
                                        v1 = append(v1, 719)
                                    case 0:
                                        _ = len(v1)
                                    case -9223372036854775808:
                                        var v60 uint32 = v43
                                        _ = v60
                                    }
                                    var v61 uint = 863
                                    fmt.Println(v57, v2, v42)
                                    _ = v61
                                }
                                _ = v57
                            } else {
                                fmt.Println(v2)
                                var v62 uint32 = v43
                                if (42 == 65535) {
                                    v1 = append(v1, (v43 >> ((65535) & 63)))
                                    v42 = float32(3.14159)
                                    v63 := GenS2{F0: "", F1: 275, F2: 0}
                                    _ = v63
                                }
                                _ = v62
                            }
                            _ = fi54
                        }
                    case "a":
                        if (0 < 1) {
                            _ = len(v1)
                            v64 := GenS4{F0: 4294967295}
                            _ = v64
                        } else {
                            switch "a" {
                            case "a":
                                v3 += fi45
                            default:
                                v43 = (^(v43))
                            }
                            if len(v1) > 0 {
                                v1[0] = (v43 ^ genFn0(42, true))
                            }
                        }
                    case "":
                        v65 := GenS5{F0: 1, F1: 2147483648, F2: -1}
                        _ = v65
                    default:
                        var v66 uint8 = 2
                        _ = v66
                    }
                }
                _ = fi45
            }
        case -9223372036854775808:
            for fi67 := 0; fi67 < 19; fi67++ {
                if len(v1) > 0 {
                    v1[0] = genFn0(56, v2)
                }
                v41 = v2
                fmt.Println(fi67, v42)
                _ = fi67
            }
        }
    }
    for fi68 := 0; fi68 < 11; fi68++ {
        fmt.Println(v41, v3, v2)
        v69 := GenS5{F0: 777, F1: genFn0(-128, (v42 > v42)), F2: -574}
        _ = fi68
        _ = v69
    }
    v3 += -2147483648
    if (-32768 == 42) {
        var v70 float32 = ((1e10 + (v42 * v42)) + 3.14159)
        v1 = append(v1, (^(v43)))
        _ = v70
    }
    fmt.Println(v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v2
    _ = v3
    _ = v41
    _ = v42
    _ = v43
}
