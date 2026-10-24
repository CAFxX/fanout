package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool; F2 uint8 }
type GenS2 struct { F0 int64; F1 uint8; F2 uint }
type GenS3 struct { F0 int8; F1 string }
type GenS4 struct { F0 uint8; F1 uint8 }
type GenS5 struct { F0 int; F1 uint }
type GenS6 struct { F0 int }
type GenS7 struct { F0 int8; F1 int }
type GenS8 struct { F0 uint8; F1 uint32 }
type GenS9 struct { F0 string; F1 uint }
type GenS10 struct { F0 int8; F1 int }

func main() {
    var v1 bool = (531 >= 443)
    var v2 float64 = (0.0 / 3.14159)
    var v3 string = (("hello" + "foo bar") + "a")
    var v4 int32 = 0
    v5 := make(map[int]uint32)
    v6 := []uint32{1, 480}
    fmt.Println(v3, v1)
    if (!(v1 && (true && (0 < 151)))) {
        v6 = append(v6, 1)
        var v7 uint = 42
        _ = v7
    } else {
        v3 = ((v3 + ("\n\t" + v3)) + v3)
        var v8 bool = (884 > 32767)
        _ = v8
    }
    v9 := GenS1{F0: (!v1), F1: v1, F2: 127}
    fmt.Println(v1, v1)
    var v10 int64 = -92
    if (false || v1) {
        fmt.Println(v3)
        fmt.Println(v2)
        if v1 {
            v11 := GenS2{F0: v10, F1: 2, F2: 602}
            fmt.Println(v4)
            delete(v5, 551)
            _ = v11
        }
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    var v12 uint8 = 60
    v1 = v1
    var v13 uint32 = 65535
    if v1 {
        fmt.Println(v12, v1, v3)
        for fi14 := 0; fi14 < 12; fi14++ {
            delete(v5, -2147483648)
            _ = fi14
        }
    }
    v15 := GenS2{F0: -9223372036854775808, F1: (^(v12)), F2: 370}
    v2 -= (v2 + 1.5)
    v16 := GenS3{F0: -31, F1: v3}
    for fi17 := 0; fi17 < 13; fi17++ {
        v18 := GenS3{F0: 41, F1: "hello"}
        if (v13 >= 830) {
            var v19 uint32 = 871
            _ = len(v5)
            switch ((v3 + (v3 + ("\n\t" + v3))) + (("hello" + v3) + (v3 + v3))) {
            case "":
                var v20 float64 = v2
                _ = v20
            case "foo bar":
                _ = len(v6)
            default:
                v4 += -2147483648
            }
            _ = v19
        }
        _ = fi17
        _ = v18
    }
    for fi21 := 0; fi21 < 3; fi21++ {
        fmt.Println(v1, v13)
        if len(v6) > 0 {
            v6[0] = (^(v13))
        }
        var v22 bool = ((797 < 2) || v1)
        _ = fi21
        _ = v22
    }
    if v1 {
        for fi23 := 0; fi23 < 17; fi23++ {
            v5[0] = (v13 & v13)
            _ = fi23
        }
        for fi24 := 0; fi24 < 20; fi24++ {
            v3 = v3
            if true {
                var v25 int64 = int64(v4)
                if v1 {
                    v26 := GenS3{F0: -79, F1: ("hello" + v3)}
                    var v27 uint64 = 0
                    _ = v26
                    _ = v27
                } else {
                    v13 -= v13
                }
                _ = v25
            }
            var v28 float64 = -2.25
            _ = fi24
            _ = v28
        }
        if (!(1 != 1)) {
            fmt.Println(v10, v10, v12)
            if v1 {
                fmt.Println(v3, v10, v10)
            } else {
                for fi29 := 0; fi29 < 14; fi29++ {
                    if ((v10 <= v10) && ((v3 + v3) != v3)) {
                        for fi30 := 0; fi30 < 15; fi30++ {
                            fmt.Println(v12)
                            v31 := v5[fi30]
                            _ = fi30
                            _ = v31
                        }
                    } else {
                        v5[fi29] = (v13 >> ((807) & 63))
                    }
                    _ = fi29
                }
                var v32 int64 = int64(v4)
                for fi33 := 0; fi33 < 11; fi33++ {
                    for fi34 := 0; fi34 < 12; fi34++ {
                        if (1000 < -1) {
                            v35 := GenS4{F0: 0, F1: v12}
                            switch (v3 + (v3 + v3)) {
                            case "hello":
                                var v36 float32 = 3.14159
                                _ = v36
                            case "\n\t":
                                v1 = (570 <= 42)
                                if (false && (!v1)) {
                                    v37 := GenS5{F0: fi33, F1: 0}
                                    var v38 float64 = (v2 + 3.14159)
                                    _ = v37
                                    _ = v38
                                } else {
                                    var v39 uint = 255
                                    _ = v39
                                }
                            default:
                                v40 := GenS1{F0: v1, F1: (-606 >= (v4 * 1)), F2: (v12 * v12)}
                                _ = v40
                            }
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                            if v1 {
                                var v41 uint8 = (^(v12))
                                _ = v41
                            } else {
                                if v1 {
                                    fmt.Println(fi34, v4, v32)
                                    var v42 int = (fi33 * (fi33 ^ fi34))
                                    _ = v42
                                }
                                if len(v6) > 0 {
                                    v6[0] = (v13 ^ (v13 << ((105) & 63)))
                                }
                                v4 = (^(v4))
                                v43 := v5[(-(fi33))]
                                fmt.Println(v10, fi34)
                                v10 -= (^(v32))
                                var v44 uint32 = v13
                                v45 := GenS6{F0: (^(fi33))}
                                _ = v43
                                _ = v44
                                _ = v45
                            }
                            v46 := GenS5{F0: (-(fi34)), F1: 255}
                            _ = v35
                            _ = v46
                        }
                        _ = fi34
                    }
                    for fi47 := 0; fi47 < 13; fi47++ {
                        v10 -= (v10 >> ((65535) & 63))
                        fmt.Println(v32, v3, v1)
                        _ = fi47
                    }
                    _ = fi33
                }
                _ = v32
            }
        } else {
            var v48 float32 = (3.14159 + -2.25)
            _ = v48
        }
    } else {
        if true {
            var v49 uint8 = 42
            v6 = append(v6, (v13 - v13))
            var v50 uint16 = 615
            _ = v49
            _ = v50
        } else {
            v51 := GenS1{F0: false, F1: v1, F2: 42}
            v52 := GenS3{F0: 0, F1: (v3 + v3)}
            for fi53 := 0; fi53 < 3; fi53++ {
                for fi54 := 0; fi54 < 3; fi54++ {
                    if v1 {
                        var v55 int64 = int64(fi54)
                        if (!true) {
                            var v56 uint32 = 42
                            if (128 != 2) {
                                if v1 {
                                    if v1 {
                                        v55 -= (v55 >> ((1) & 63))
                                        v57 := GenS1{F0: v1, F1: (fi53 <= (fi54 * (fi54 - 128))), F2: (v12 >> ((18446744073709551615) & 63))}
                                        v55 += int64(v4)
                                        if v1 {
                                            switch (v3 + "hello") {
                                            case "":
                                                switch (fi54 ^ fi54) {
                                                case 2147483647:
                                                    v12 = 255
                                                case 128:
                                                    if v1 {
                                                        switch v3 {
                                                        case "a":
                                                            var v58 uint16 = 926
                                                            _ = v58
                                                        case "\n\t":
                                                            _ = len(v6)
                                                        case "":
                                                            var v59 float32 = 0.1
                                                            switch fi54 {
                                                            case -2147483648:
                                                                var v60 string = "\n\t"
                                                                _ = v60
                                                            }
                                                            _ = v59
                                                        }
                                                        v4 += v4
                                                        v61 := GenS1{F0: v1, F1: (fi53 <= (fi53 & (fi53 / ((fi54) | 1)))), F2: 128}
                                                        _ = v61
                                                    }
                                                default:
                                                    var v62 float64 = (-2.25 * v2)
                                                    _ = v62
                                                }
                                            }
                                            v12 += v12
                                        }
                                        v63 := GenS7{F0: 1, F1: fi53}
                                        v10 += (v10 | (^(v55)))
                                        v64 := GenS6{F0: (fi54 * -1)}
                                        _ = v57
                                        _ = v63
                                        _ = v64
                                    } else {
                                        v55 = 293
                                    }
                                }
                                v65 := GenS8{F0: v12, F1: 126}
                                _ = v65
                            }
                            v4 += v4
                            _ = v56
                        }
                        if (!v1) {
                            if true {
                                if (0 == -2147483648) {
                                    v66 := v5[fi53]
                                    if (((0.1 / 1.5) == 1.5) && ((^(v66)) >= (v13 * v66))) {
                                        if (!false) {
                                            _ = len(v5)
                                            v55 -= (v55 ^ v55)
                                            fmt.Println(v10, v13)
                                        } else {
                                            if (v1 || (v10 == 1)) {
                                                v12 = (v12 - (v12 << ((uint(v12)) & 63)))
                                            } else {
                                                var v67 string = v3
                                                _ = v67
                                            }
                                            if true {
                                                v68 := GenS3{F0: 28, F1: v3}
                                                var v69 float32 = float32((v2 / v2))
                                                _ = v68
                                                _ = v69
                                            }
                                            var v70 float64 = (v2 - 3.14159)
                                            v71 := v5[fi53]
                                            v71 = (v13 | v66)
                                            _ = v70
                                            _ = v71
                                        }
                                        v3 = v3
                                        var v72 uint16 = 42
                                        _ = v72
                                    }
                                    if false {
                                        var v73 int32 = v4
                                        v74 := GenS7{F0: 1, F1: -1}
                                        _ = v73
                                        _ = v74
                                    } else {
                                        v75 := GenS9{F0: "foo bar", F1: 18446744073709551615}
                                        var v76 bool = v1
                                        _ = v75
                                        _ = v76
                                    }
                                    _ = v66
                                } else {
                                    if v1 {
                                        var v77 int32 = 1
                                        if len(v6) > 0 {
                                            v6[0] = (^(v13))
                                        }
                                        _ = v77
                                    } else {
                                        var v78 uint = 953
                                        var v79 int8 = int8(v78)
                                        v80 := GenS10{F0: (v79 / ((0) | 1)), F1: fi53}
                                        _ = v78
                                        _ = v79
                                        _ = v80
                                    }
                                    fmt.Println(v55, v12)
                                }
                                switch fi54 {
                                case 32767:
                                    var v81 uint = 2
                                    _ = v81
                                case 270:
                                    if v1 {
                                        v12 += (v12 - (v12 >> ((255) & 63)))
                                    }
                                case 32768:
                                    var v82 string = v3
                                    _ = v82
                                default:
                                    if (true || false) {
                                        fmt.Println(v13, v3)
                                        v1 = v1
                                        var v83 float32 = 0.1
                                        _ = v83
                                    } else {
                                        v4 = v4
                                        if len(v6) >= 2 {
                                            v6 = v6[1:len(v6)]
                                        }
                                    }
                                }
                                v4 = v4
                            }
                            switch 9223372036854775807 {
                            case 128:
                                if true {
                                    v84 := GenS1{F0: v1, F1: (v1 && ((-1 >= v10) && true)), F2: v12}
                                    v55 = 1
                                    switch fi54 {
                                    case 950:
                                        switch (-(fi54)) {
                                        case 32768:
                                            if false {
                                                var v85 int32 = v4
                                                if v1 {
                                                    if true {
                                                        v85 += (v4 ^ (v4 << ((65535) & 63)))
                                                    }
                                                    var v86 bool = ((v1 && (v1 || false)) && (!("hello" != v3)))
                                                    v4 = int32(v12)
                                                    _ = v86
                                                }
                                                _ = v85
                                            } else {
                                                fmt.Println(v55, v3)
                                            }
                                        default:
                                            var v87 uint16 = 314
                                            _ = v87
                                        }
                                    }
                                    v88 := GenS2{F0: v10, F1: v12, F2: 28}
                                    switch (v3 + v3) {
                                    case "":
                                        v89 := GenS2{F0: (v10 >> ((4294967295) & 63)), F1: (v12 >> ((65535) & 63)), F2: 65535}
                                        _ = v89
                                    default:
                                        var v90 int8 = -1
                                        _ = v90
                                    }
                                    _ = v84
                                    _ = v88
                                } else {
                                    v12 += (v12 ^ 2)
                                    var v91 uint8 = (v12 | (^(v12)))
                                    _ = v91
                                }
                            case 1:
                                var v92 int16 = 900
                                _ = v92
                            case -9223372036854775808:
                                _ = len(v6)
                            default:
                                var v93 int32 = (v4 & (-(v4)))
                                _ = v93
                            }
                            _ = len(v5)
                            _ = len(v5)
                            delete(v5, (-(fi53)))
                        } else {
                            v6 = append(v6, (v13 | v13))
                            if len(v6) > 0 {
                                v6[0] = v13
                            }
                            fmt.Println(fi54, v1)
                            var v94 uint32 = 497
                            _ = v94
                        }
                        _ = v55
                    }
                    _ = fi54
                }
                _ = fi53
            }
            _ = v51
            _ = v52
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
