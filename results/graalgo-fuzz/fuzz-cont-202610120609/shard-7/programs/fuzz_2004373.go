package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32; F2 int8 }
type GenS2 struct { F0 int32; F1 int32 }
type GenS3 struct { F0 int32; F1 int64; F2 int }
type GenS4 struct { F0 int16; F1 int16 }
type GenS5 struct { F0 uint16; F1 int32; F2 uint32 }
type GenS6 struct { F0 int16; F1 bool; F2 bool }
type GenS7 struct { F0 uint32; F1 uint8 }

func main() {
    v1 := []int{-2147483648, -2147483648, 32768}
    v2 := make(map[int]int)
    var v3 int64 = 42
    var v4 string = "foo bar"
    v5 := []string{}
    if true {
        v4 = v4
        v4 = (v4 + "")
        if len(v1) > 0 {
            v1[0] = -9223372036854775808
        }
    } else {
        var v6 int = -9223372036854775808
        _ = v6
    }
    if false {
        v3 = v3
    } else {
        var v7 int = 0
        for fi8 := 0; fi8 < 8; fi8++ {
            if len(v1) > 0 {
                v1[0] = (fi8 + (v7 | fi8))
            }
            var v9 float64 = 0.0
            _ = fi8
            _ = v9
        }
        fmt.Println(v3)
        _ = v7
    }
    for fi10 := 0; fi10 < 12; fi10++ {
        v11 := GenS1{F0: 1, F1: 0, F2: 1}
        if true {
            for fi12 := 0; fi12 < 16; fi12++ {
                switch fi10 {
                case 42:
                    var v13 uint8 = 127
                    _ = v13
                }
                if (!(696 >= 2147483647)) {
                    fmt.Println(v4)
                } else {
                    v14 := GenS2{F0: -591, F1: -1}
                    if true {
                        v15 := GenS1{F0: 14, F1: -1, F2: -99}
                        _ = v15
                    } else {
                        if (42 <= 42) {
                            v16 := GenS1{F0: 87, F1: -1, F2: 100}
                            _ = v16
                        } else {
                            var v17 int16 = -1
                            var v18 uint8 = 0
                            _ = v17
                            _ = v18
                        }
                        if true {
                            if (int32(v3) > 556) {
                                var v19 int8 = 0
                                v19 += -1
                                v3 = 1
                                _ = v19
                            }
                            v3 -= (v3 / ((v3) | 1))
                        }
                        v3 = (^(v3))
                    }
                    _ = v14
                }
                _ = fi12
            }
            v20 := GenS2{F0: -1, F1: 2147483647}
            _ = v20
        } else {
            if false {
                v21 := v2[(fi10 ^ 592)]
                for fi22 := 0; fi22 < 19; fi22++ {
                    fmt.Println(v4, v4)
                    var v23 string = "\n\t"
                    fmt.Println(v23)
                    _ = fi22
                    _ = v23
                }
                for fi24 := 0; fi24 < 16; fi24++ {
                    v3 = (-(v3))
                    var v25 string = "hello"
                    _ = fi24
                    _ = v25
                }
                _ = v21
            }
            v26 := GenS3{F0: 42, F1: (v3 ^ v3), F2: fi10}
            for fi27 := 0; fi27 < 20; fi27++ {
                fmt.Println(v4, fi27, v4)
                if ((fi27 - fi27) != fi27) {
                    if (42 >= 4294967295) {
                        fmt.Println(v3, v4, fi10)
                    } else {
                        if (42 < -233) {
                            v4 = "foo bar"
                        } else {
                            if (!((!(fi27 != 231)) && (!(!(!false))))) {
                                v28 := v2[(fi27 ^ 133)]
                                var v29 bool = true
                                var v30 int64 = (^(v3))
                                _ = v28
                                _ = v29
                                _ = v30
                            }
                            v3 = (v3 - (v3 | (v3 / ((-9223372036854775808) | 1))))
                            if true {
                                var v31 uint8 = 2
                                if false {
                                    v4 = ((v4 + "") + "foo bar")
                                    delete(v2, (fi10 + (fi10 ^ 787)))
                                }
                                v32 := GenS1{F0: -60, F1: -317, F2: 127}
                                _ = v31
                                _ = v32
                            } else {
                                if true {
                                    v4 = ""
                                    var v33 int16 = -424
                                    var v34 bool = ((1e10 + float64((1e10 + -2.25))) != 3.14159)
                                    var v35 int32 = -1
                                    fmt.Println(v3)
                                    v4 = "foo bar"
                                    _ = v33
                                    _ = v34
                                    _ = v35
                                }
                                v4 = "a"
                                v3 -= (v3 - (-(v3)))
                            }
                        }
                        var v36 uint32 = 740
                        var v37 string = v4
                        _ = v36
                        _ = v37
                    }
                    switch (fi10 * fi27) {
                    case 9223372036854775807:
                        v4 = (v4 + v4)
                    case 719:
                        var v38 int = fi10
                        switch "a" {
                        case "hello":
                            v4 = v4
                        case "a":
                            var v39 int64 = (v3 - -9223372036854775808)
                            _ = v39
                        case "foo bar":
                            var v40 uint8 = 127
                            _ = v40
                        }
                        _ = v38
                    }
                    if true {
                        switch "\n\t" {
                        case "foo bar":
                            if true {
                                var v41 int32 = 33
                                fmt.Println(v3)
                                switch ("hello" + ("\n\t" + ("\n\t" + "a"))) {
                                case "":
                                    if (774 >= 323) {
                                        v41 += (v41 % ((v41) | 1))
                                        fmt.Println(v41, v4)
                                        v41 += (v41 + 625)
                                        if (((1.5 >= 1e10) || (114 >= 127)) && false) {
                                            switch v4 {
                                            case "a":
                                                var v42 float32 = 1e10
                                                v43 := GenS3{F0: v41, F1: 0, F2: fi27}
                                                var v44 string = "hello"
                                                _ = v42
                                                _ = v43
                                                _ = v44
                                            case "\n\t":
                                                switch 4294967295 {
                                                case 42:
                                                    switch fi27 {
                                                    case -426:
                                                        if (42 <= 4294967295) {
                                                            var v45 int = fi27
                                                            var v46 uint16 = 0
                                                            if (18446744073709551615 <= 0) {
                                                                v47 := GenS4{F0: -32768, F1: 507}
                                                                v3 += 1
                                                                if ((3.14159 + -2.25) < 0.1) {
                                                                    var v48 bool = true
                                                                    _ = v48
                                                                } else {
                                                                    v45 = (fi27 >> ((65535) & 63))
                                                                    v49 := GenS4{F0: -32768, F1: 32767}
                                                                    switch v45 {
                                                                    case 4294967295:
                                                                        v45 = (fi27 << ((4294967295) & 63))
                                                                    default:
                                                                        fmt.Println(fi10)
                                                                    }
                                                                    _ = v49
                                                                }
                                                                _ = v47
                                                            }
                                                            var v50 int64 = v3
                                                            _ = v45
                                                            _ = v46
                                                            _ = v50
                                                        } else {
                                                            v41 = v41
                                                        }
                                                    case 419:
                                                        switch v4 {
                                                        case "hello":
                                                            v3 += v3
                                                        case "foo bar":
                                                            v41 -= (v41 + int32(v3))
                                                        default:
                                                            v4 = (v4 + "hello")
                                                        }
                                                    case 42:
                                                        var v51 float64 = 3.14159
                                                        _ = v51
                                                    default:
                                                        v3 -= -555
                                                    }
                                                case 1:
                                                    v52 := GenS1{F0: 1, F1: v41, F2: 120}
                                                    _ = v52
                                                case 32767:
                                                    v1 = append(v1, -1)
                                                default:
                                                    var v53 float64 = 1.5
                                                    _ = v53
                                                }
                                            case "foo bar":
                                                v54 := GenS5{F0: 255, F1: (v41 % ((v41) | 1)), F2: 0}
                                                _ = v54
                                            default:
                                                var v55 uint16 = 0
                                                var v56 int64 = v3
                                                v56 += (v3 * (v3 - v3))
                                                _ = v55
                                                _ = v56
                                            }
                                        } else {
                                            switch fi27 {
                                            case 4294967295:
                                                if (2 < 574) {
                                                    var v57 bool = true
                                                    v58 := GenS3{F0: (-(v41)), F1: v3, F2: -1}
                                                    _ = v57
                                                    _ = v58
                                                } else {
                                                    var v59 uint16 = 42
                                                    v41 = (v41 | v41)
                                                    _ = v59
                                                }
                                            case 995:
                                                if (!(0 > 65535)) {
                                                    delete(v2, fi27)
                                                    var v60 uint8 = 128
                                                    _ = v60
                                                } else {
                                                    var v61 uint16 = 0
                                                    _ = v61
                                                }
                                            case -2147483648:
                                                if true {
                                                    var v62 uint = 65535
                                                    _ = v62
                                                }
                                            default:
                                                var v63 bool = (!(((1.5 - 3.14159) + 1.5) >= ((float64(float32(((float64(1e10) - 3.14159) - -2.25))) / 1.0) / 3.14159)))
                                                switch -2147483648 {
                                                case -934:
                                                    v64 := GenS3{F0: (^(v41)), F1: v3, F2: fi27}
                                                    _ = v64
                                                }
                                                _ = v63
                                            }
                                            v3 = v3
                                            v65 := GenS2{F0: v41, F1: 42}
                                            _ = v65
                                        }
                                    }
                                default:
                                    var v66 int8 = 1
                                    _ = v66
                                }
                                var v67 int16 = -298
                                _ = v41
                                _ = v67
                            } else {
                                v4 = "hello"
                            }
                        }
                        v3 = v3
                    } else {
                        switch (fi10 | fi10) {
                        case 42:
                            v4 = v4
                        default:
                            switch "\n\t" {
                            case "a":
                                v4 = v4
                                fmt.Println(v4, v3)
                            default:
                                v68 := GenS6{F0: 635, F1: (42 < 2), F2: true}
                                _ = v68
                            }
                        }
                        var v69 uint8 = 255
                        delete(v2, fi10)
                        v70 := GenS3{F0: -1, F1: v3, F2: (^(fi10))}
                        _ = v69
                        _ = v70
                    }
                } else {
                    v3 = (v3 ^ (v3 ^ -1))
                    if (246 <= 185) {
                        fmt.Println(fi27, v4)
                    }
                    var v71 string = "\n\t"
                    _ = v71
                }
                v4 = "a"
                v72 := GenS7{F0: 0, F1: 236}
                v4 = v4
                _ = fi27
                _ = v72
            }
            _ = v26
        }
        if true {
            if (false || true) {
                switch fi10 {
                case -115:
                    var v73 int16 = 977
                    _ = v73
                case 297:
                    _ = len(v5)
                case 881:
                    v74 := GenS3{F0: 0, F1: (v3 << ((42) & 63)), F2: (fi10 << ((18446744073709551615) & 63))}
                    _ = v74
                }
                for fi75 := 0; fi75 < 16; fi75++ {
                    v76 := GenS1{F0: 0, F1: 0, F2: 3}
                    var v77 int32 = -2147483648
                    v4 = (v4 + (((v4 + v4) + v4) + (("hello" + "hello") + "hello")))
                    var v78 int64 = v3
                    _ = fi75
                    _ = v76
                    _ = v77
                    _ = v78
                }
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
        }
        _ = fi10
        _ = v11
    }
    var v79 int = -780
    switch v79 {
    case 32768:
        var v80 uint32 = 4294967295
        _ = v80
    default:
        switch -2147483648 {
        case 2147483647:
            var v81 int64 = 42
            _ = v81
        }
    }
    v5 = append(v5, "\n\t")
    v4 = "foo bar"
    var v82 int64 = v3
    if len(v5) > 0 {
        v5[0] = (v4 + ((v4 + v4) + v4))
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v79:", v79)
    fmt.Println("v82:", v82)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v79
    _ = v82
}
