package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32; F2 int64 }
type GenS2 struct { F0 uint32; F1 int; F2 int32 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 string; F1 int; F2 uint8 }
type GenS5 struct { F0 uint64; F1 uint }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 uint8; F1 int64; F2 uint }
func genFn0(p0 int) int64 {
    for fi7 := 0; fi7 < 14; fi7++ {
        v8 := []int{p0, (p0 + p0)}
        var v9 float64 = (((float64(((-2.25 - (((0.1 / (1e10 / 1.0)) - (0.1 * (1e10 - 0.0))) * 3.14159)) + -2.25)) / 1.5) - float64(float32(float64(0.1)))) + 0.0)
        for fi10 := 0; fi10 < 6; fi10++ {
            fmt.Println(fi10, p0, fi7)
            fmt.Println(p0)
            var v11 uint16 = 766
            if true {
                if false {
                    p0 = 818
                } else {
                    v9 = 3.14159
                    p0 += 415
                }
                v12 := GenS1{F0: 0, F1: -2147483648, F2: 0}
                v13 := GenS2{F0: 65535, F1: (^(fi7)), F2: 1}
                var v14 uint16 = (v11 + (v11 ^ 2))
                var v15 uint16 = v14
                _ = v12
                _ = v13
                _ = v14
                _ = v15
            } else {
                v16 := GenS2{F0: 1, F1: 32767, F2: 2147483647}
                var v17 uint16 = v11
                _ = v16
                _ = v17
            }
            _ = fi10
            _ = v11
        }
        _ = fi7
        _ = v8
        _ = v9
    }
    var v18 uint = 1
    _ = p0
    _ = v18
    return 1
}
func genFn1() string {
    switch -2147483648 {
    case 720:
        if (-1 != -1) {
            for fi19 := 0; fi19 < 5; fi19++ {
                v20 := []int{(fi19 & 4294967295), 42}
                var v21 int64 = 42
                var v22 string = "foo bar"
                _ = fi19
                _ = v20
                _ = v21
                _ = v22
            }
            var v23 uint8 = 118
            _ = v23
        } else {
            var v24 int8 = 127
            _ = v24
        }
    case 0:
        if (-1 >= 2147483647) {
            v25 := []int{}
            _ = v25
        }
    default:
        var v26 uint16 = 65535
        _ = v26
    }
    v27 := GenS1{F0: 523, F1: 2147483647, F2: genFn0(2147483647)}
    if (42 == 65535) {
        var v28 int = 2147483647
        _ = v28
    } else {
        if (!false) {
            if (1 > 1) {
                v29 := GenS3{F0: 2147483647}
                if ((1 < 12) && ((true || ("" != "foo bar")) || false)) {
                    var v30 uint32 = 2
                    for fi31 := 0; fi31 < 15; fi31++ {
                        v30 += 2147483647
                        v30 += (^(v30))
                        _ = fi31
                    }
                    fmt.Println(v30)
                    _ = v30
                }
                for fi32 := 0; fi32 < 2; fi32++ {
                    var v33 uint8 = 255
                    _ = fi32
                    _ = v33
                }
                _ = v29
            } else {
                v34 := GenS3{F0: 42}
                _ = v34
            }
        }
    }
    v35 := []int{128}
    switch "hello" {
    case "\n\t":
        var v36 uint8 = 127
        _ = v36
    case "a":
        switch "" {
        case "\n\t":
            if (!true) {
                v37 := GenS1{F0: 42, F1: 42, F2: 516}
                var v38 int64 = -1
                switch "foo bar" {
                case "hello":
                    fmt.Println(v38)
                case "a":
                    var v39 bool = (-2147483648 < 0)
                    _ = v39
                case "":
                    v38 -= v38
                }
                _ = v37
                _ = v38
            }
        case "a":
            if true {
                v40 := GenS4{F0: "a", F1: -9223372036854775808, F2: 12}
                _ = v40
            }
        default:
            var v41 int32 = 601
            _ = v41
        }
    default:
        switch -489 {
        case 42:
            var v42 float32 = -2.25
            _ = v42
        case -9223372036854775808:
            if true {
                for fi43 := 0; fi43 < 15; fi43++ {
                    v44 := GenS1{F0: -542, F1: 0, F2: -1}
                    var v45 float64 = ((float64(0.0) / 1.0) * (3.14159 * 0.0))
                    v45 = 0.0
                    _ = fi43
                    _ = v44
                    _ = v45
                }
                for fi46 := 0; fi46 < 7; fi46++ {
                    v47 := GenS5{F0: 183, F1: 1}
                    var v48 int16 = -32768
                    v35 = append(v35, -1)
                    _ = fi46
                    _ = v47
                    _ = v48
                }
                var v49 string = "hello"
                _ = v49
            } else {
                if len(v35) >= 2 {
                    v35 = v35[1:len(v35)]
                }
            }
        default:
            var v50 uint8 = 127
            _ = v50
        }
    }
    _ = v27
    _ = v35
    return "a"
}

func main() {
    var v1 int32 = 983
    var v2 int64 = -557
    v3 := []uint32{0}
    v4 := []int{-9223372036854775808}
    v5 := []int{-849, 32767}
    var v6 float64 = float64(3.14159)
    if (42 <= -128) {
        var v51 uint64 = 2
        for fi52 := 0; fi52 < 13; fi52++ {
            v2 = (v2 >> ((672) & 63))
            v6 = (v6 / v6)
            var v53 int16 = 1
            _ = fi52
            _ = v53
        }
        v3 = append(v3, 42)
        _ = v51
    }
    v6 += 1.5
    switch 130 {
    case 32768:
        switch 441 {
        case 32768:
            var v54 float64 = (((v6 + (((v6 - (float64(float32((v6 / v6))) / ((0.1 / ((v6 / ((v6 / ((-2.25 - v6) - 1e10)) * v6)) / v6)) - 0.1))) + v6) + v6)) - v6) / 1.0)
            _ = v54
        default:
            v6 = (-2.25 - ((v6 * (v6 / (v6 / float64(0.1)))) * float64(3.14159)))
        }
    case -213:
        v2 = genFn0(127)
    }
    for fi55 := 0; fi55 < 4; fi55++ {
        for fi56 := 0; fi56 < 7; fi56++ {
            v2 = v2
            _ = fi56
        }
        _ = fi55
    }
    for fi57 := 0; fi57 < 16; fi57++ {
        v4 = append(v4, fi57)
        _ = fi57
    }
    if false {
        v6 = (1e10 + (-2.25 + 1.5))
        if false {
            var v58 uint32 = 2147483647
            for fi59 := 0; fi59 < 4; fi59++ {
                v60 := GenS5{F0: 18446744073709551615, F1: 804}
                if true {
                    if true {
                        fmt.Println(v1)
                        v58 = 2147483648
                    } else {
                        for fi61 := 0; fi61 < 1; fi61++ {
                            if len(v5) > 0 {
                                v5[0] = fi59
                            }
                            fmt.Println(v58, fi59)
                            var v62 bool = ((fi61 < fi61) || false)
                            switch (-(fi59)) {
                            case -1:
                                v3 = append(v3, 42)
                            default:
                                if (!("hello" == "\n\t")) {
                                    v6 = v6
                                    v63 := GenS4{F0: "a", F1: fi61, F2: 42}
                                    v64 := GenS3{F0: 4294967295}
                                    _ = v63
                                    _ = v64
                                } else {
                                    var v65 int8 = 69
                                    _ = v65
                                }
                            }
                            if v62 {
                                if (v62 || true) {
                                    v4 = append(v4, 32768)
                                } else {
                                    var v66 int = fi59
                                    v58 += 2
                                    _ = v66
                                }
                            } else {
                                fmt.Println(v62)
                                if (v62 && ((v6 - (v6 * ((v6 + float64(-2.25)) + float64(((((0.1 / 1.5) / -2.25) + (0.0 * ((((0.0 + 1.5) * 1e10) / 1.0) / 3.14159))) / float32((0.0 - v6))))))) <= v6)) {
                                    v67 := GenS3{F0: 1}
                                    v6 += float64(1e10)
                                    _ = v67
                                } else {
                                    var v68 uint32 = v58
                                    var v69 bool = (false && v62)
                                    _ = v68
                                    _ = v69
                                }
                                var v70 float64 = 0.0
                                if false {
                                    if true {
                                        if (263 != 1) {
                                            var v71 float64 = float64(3.14159)
                                            v72 := GenS1{F0: (v1 << ((2) & 63)), F1: int32(v58), F2: v2}
                                            _ = v71
                                            _ = v72
                                        }
                                    } else {
                                        var v73 int16 = 1
                                        _ = v73
                                    }
                                    v74 := GenS6{F0: 18446744073709551615}
                                    _ = v74
                                }
                                var v75 uint16 = 65535
                                _ = v70
                                _ = v75
                            }
                            _ = fi61
                            _ = v62
                        }
                        v76 := GenS4{F0: "a", F1: fi59, F2: 255}
                        v6 += v6
                        _ = v76
                    }
                    v2 = (v2 % ((v2) | 1))
                    _ = len(v4)
                } else {
                    var v77 uint16 = 369
                    for fi78 := 0; fi78 < 4; fi78++ {
                        v79 := GenS3{F0: 965}
                        v80 := GenS4{F0: "foo bar", F1: fi59, F2: 2}
                        v81 := GenS5{F0: 960, F1: 2}
                        _ = fi78
                        _ = v79
                        _ = v80
                        _ = v81
                    }
                    _ = v77
                }
                var v82 int16 = 1
                _ = fi59
                _ = v60
                _ = v82
            }
            _ = v58
        } else {
            for fi83 := 0; fi83 < 19; fi83++ {
                v6 += v6
                var v84 int = fi83
                if len(v4) > 0 {
                    v4[0] = 4294967295
                }
                _ = fi83
                _ = v84
            }
            var v85 uint64 = 18446744073709551615
            _ = v85
        }
        for fi86 := 0; fi86 < 9; fi86++ {
            var v87 int8 = -83
            _ = fi86
            _ = v87
        }
    } else {
        var v88 uint8 = 255
        v6 = v6
        _ = v88
    }
    v89 := GenS7{F0: 2, F1: 9223372036854775807, F2: 2}
    v90 := GenS1{F0: v1, F1: v1, F2: v2}
    var v91 int64 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v91:", v91)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v89
    _ = v90
    _ = v91
}
