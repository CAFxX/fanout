package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16 }
type GenS2 struct { F0 int64; F1 int; F2 uint }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint64; F1 int64 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 string; F1 int; F2 int }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 uint8; F1 int8; F2 int8 }
type GenS9 struct { F0 string }

func main() {
    var v1 float32 = (1e10 * (1.5 / 1.0))
    v2 := make(map[int]uint32)
    var v3 uint64 = 18446744073709551615
    v4 := GenS1{F0: 1000, F1: 1}
    var v5 int8 = 100
    if false {
        v3 -= (v3 | v3)
        _ = len(v2)
    } else {
        if false {
            var v6 bool = false
            if (0 == 2) {
                v7 := GenS2{F0: -360, F1: -9223372036854775808, F2: 809}
                _ = v7
            } else {
                for fi8 := 0; fi8 < 3; fi8++ {
                    if v6 {
                        var v9 int = fi8
                        switch 42 {
                        case 9223372036854775807:
                            var v10 int64 = 1
                            _ = v10
                        default:
                            switch (^(v9)) {
                            case -469:
                                v9 = (fi8 - (v9 >> ((42) & 63)))
                            case 42:
                                v11 := GenS3{F0: 4294967295}
                                _ = v11
                            case -9223372036854775808:
                                v6 = v6
                            }
                        }
                        for fi12 := 0; fi12 < 13; fi12++ {
                            fmt.Println(v1)
                            v5 = 15
                            _ = fi12
                        }
                        _ = v9
                    } else {
                        for fi13 := 0; fi13 < 7; fi13++ {
                            if (1 >= 42) {
                                fmt.Println(v5)
                                if (2147483647 == 65535) {
                                    if (1 < 42) {
                                        var v14 float32 = ((v1 + -2.25) + -2.25)
                                        _ = v14
                                    } else {
                                        v15 := v2[fi8]
                                        v16 := []int{(fi13 ^ fi13), (fi8 | fi13)}
                                        v1 = ((-2.25 + (v1 - ((-2.25 / (0.1 - v1)) - 0.0))) + 0.0)
                                        var v17 float64 = 0.1
                                        var v18 bool = (2147483647 >= int32(v3))
                                        _ = v15
                                        _ = v16
                                        _ = v17
                                        _ = v18
                                    }
                                    var v19 int64 = 9223372036854775807
                                    fmt.Println(fi13, v6)
                                    _ = v19
                                }
                                v20 := []int{}
                                _ = v20
                            } else {
                                var v21 uint = 4294967295
                                if ("\n\t" <= ("foo bar" + "")) {
                                    var v22 int32 = 2147483647
                                    var v23 float64 = 1.5
                                    var v24 uint64 = v3
                                    if (-32768 > 0) {
                                        if (v6 || v6) {
                                            var v25 float32 = ((v1 + v1) - float32(((((v23 * 3.14159) - float64(0.0)) / v23) + v23)))
                                            _ = v25
                                        }
                                        var v26 string = "hello"
                                        v23 = v23
                                        _ = v26
                                    } else {
                                        v3 -= 42
                                        _ = len(v2)
                                        v23 = 0.1
                                        v27 := v2[910]
                                        _ = v27
                                    }
                                    _ = v22
                                    _ = v23
                                    _ = v24
                                } else {
                                    v28 := []int{fi8}
                                    var v29 uint8 = 223
                                    _ = v28
                                    _ = v29
                                }
                                _ = v21
                            }
                            v30 := GenS1{F0: 32767, F1: 1000}
                            _ = fi13
                            _ = v30
                        }
                    }
                    _ = fi8
                }
            }
            var v31 int64 = 115
            _ = v31
            _ = v6
        } else {
            var v32 int32 = 1
            for fi33 := 0; fi33 < 15; fi33++ {
                if false {
                    v1 = v1
                    v3 -= (v3 << ((uint(v3)) & 63))
                    v34 := GenS4{F0: (v3 >> ((0) & 63)), F1: 0}
                    _ = v34
                }
                _ = fi33
            }
            v3 = 18446744073709551615
            _ = v32
        }
    }
    var v35 uint64 = 0
    if false {
        fmt.Println(v5)
        v2[9223372036854775807] = 4294967295
    } else {
        var v36 float32 = v1
        _ = v36
    }
    if ((!((false || (586 >= 26)) && false)) || (true || true)) {
        var v37 int16 = 632
        var v38 int64 = 42
        _ = v37
        _ = v38
    } else {
        var v39 int = -2147483648
        v40 := GenS4{F0: (v35 + (v3 & 994)), F1: 0}
        fmt.Println(v35, v1, v5)
        _ = v39
        _ = v40
    }
    v1 = (v1 / 0.1)
    v35 += 49
    if false {
        v3 = v35
    } else {
        var v41 float32 = (((v1 * 0.0) + (v1 + v1)) - -2.25)
        for fi42 := 0; fi42 < 15; fi42++ {
            v3 = 18446744073709551615
            v43 := []int{(fi42 / ((0) | 1))}
            _ = fi42
            _ = v43
        }
        _ = v41
    }
    for fi44 := 0; fi44 < 8; fi44++ {
        for fi45 := 0; fi45 < 4; fi45++ {
            v46 := v2[(-(fi44))]
            delete(v2, (fi44 << ((65535) & 63)))
            var v47 int16 = 605
            _ = fi45
            _ = v46
            _ = v47
        }
        var v48 string = "hello"
        _ = fi44
        _ = v48
    }
    if (0 > -457) {
        switch -9223372036854775808 {
        case 0:
            v35 -= (v35 & v3)
        case 128:
            for fi49 := 0; fi49 < 18; fi49++ {
                v50 := GenS1{F0: 0, F1: -32768}
                fmt.Println(v5)
                _ = fi49
                _ = v50
            }
        case 32767:
            v51 := GenS3{F0: 4294967295}
            _ = v51
        }
    } else {
        for fi52 := 0; fi52 < 7; fi52++ {
            v53 := GenS5{F0: (((127 >= 2) || (189 == 128)) || false)}
            _ = fi52
            _ = v53
        }
        var v54 string = ("foo bar" + "a")
        _ = v54
    }
    if (42 < -2147483648) {
        fmt.Println(v1)
        v55 := GenS6{F0: "", F1: 32767, F2: 4294967295}
        v3 -= v3
        _ = v55
    }
    if true {
        if (1 < -1) {
            v56 := GenS7{F0: 2147483648}
            v57 := []int{9223372036854775807, 761, 32768}
            _ = v56
            _ = v57
        } else {
            for fi58 := 0; fi58 < 12; fi58++ {
                var v59 bool = (false || ("\n\t" == "foo bar"))
                _ = fi58
                _ = v59
            }
            for fi60 := 0; fi60 < 14; fi60++ {
                v35 = v3
                _ = len(v2)
                _ = fi60
            }
            var v61 uint64 = v3
            _ = v61
        }
        var v62 string = "foo bar"
        _ = v62
    } else {
        for fi63 := 0; fi63 < 9; fi63++ {
            v5 -= 0
            var v64 int8 = 71
            _ = fi63
            _ = v64
        }
        if true {
            v65 := GenS1{F0: -813, F1: 95}
            fmt.Println(v3, v3)
            _ = v65
        } else {
            var v66 int64 = 0
            switch 0 {
            case -370:
                for fi67 := 0; fi67 < 8; fi67++ {
                    v2[-657] = 0
                    var v68 int16 = -32768
                    var v69 float64 = -2.25
                    _ = fi67
                    _ = v68
                    _ = v69
                }
            case 1:
                if (v66 >= v66) {
                    if (175 < 2) {
                        var v70 string = "\n\t"
                        for fi71 := 0; fi71 < 13; fi71++ {
                            fmt.Println(v35, v70)
                            v72 := GenS8{F0: 255, F1: (v5 + 100), F2: 0}
                            var v73 uint8 = 128
                            _ = fi71
                            _ = v72
                            _ = v73
                        }
                        _ = v70
                    } else {
                        v3 = v3
                        v74 := GenS5{F0: true}
                        _ = v74
                    }
                    if (37 != 0) {
                        v75 := []int{2147483647}
                        _ = v75
                    }
                    if (!((-1 >= -1) || false)) {
                        var v76 int8 = 1
                        v35 -= v3
                        var v77 uint32 = 47
                        _ = v76
                        _ = v77
                    } else {
                        v66 += v66
                        v3 = 766
                        for fi78 := 0; fi78 < 20; fi78++ {
                            var v79 float64 = float64(v1)
                            v3 -= (^(v3))
                            _ = fi78
                            _ = v79
                        }
                    }
                }
            default:
                v80 := []int{32768}
                _ = v80
            }
            _ = v66
        }
        v35 += v3
    }
    v3 += 4294967295
    v81 := v2[32768]
    for fi82 := 0; fi82 < 4; fi82++ {
        v83 := []int{(fi82 * (fi82 << ((422) & 63))), 128, -519, fi82}
        _ = fi82
        _ = v83
    }
    if ("\n\t" == "hello") {
        v3 = (v35 & v3)
        delete(v2, 32768)
        var v84 float64 = 3.14159
        _ = v84
    } else {
        for fi85 := 0; fi85 < 7; fi85++ {
            v1 += float32(float64((1e10 * 1.5)))
            v86 := []int{fi85, (-(fi85)), (fi85 * fi85), fi85}
            v87 := GenS9{F0: "foo bar"}
            _ = fi85
            _ = v86
            _ = v87
        }
    }
    for fi88 := 0; fi88 < 18; fi88++ {
        var v89 uint32 = 65535
        var v90 bool = false
        v1 += ((1e10 - v1) * 3.14159)
        _ = fi88
        _ = v89
        _ = v90
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v5:", v5)
    fmt.Println("v81:", v81)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v5
    _ = v81
}
