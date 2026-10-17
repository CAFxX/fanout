package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8 }
type GenS2 struct { F0 int32; F1 int }
type GenS3 struct { F0 uint16; F1 uint32; F2 int }
type GenS4 struct { F0 string }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint8; F1 uint; F2 bool }
func genFn0(p0 int16, p1 int, p2 int64) uint16 {
    var v7 uint8 = 168
    p0 -= 1000
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 1
}
func genFn1(p0 int64, p1 int32) uint8 {
    p0 = -692
    p1 -= p1
    if false {
        var v8 float32 = (1.5 / ((1.5 / (0.1 / ((0.0 + 1.5) / 0.1))) + (0.0 / (-2.25 - 1e10))))
        v9 := []int{32767}
        _ = v8
        _ = v9
    } else {
        p1 = p1
        switch "" {
        case "a":
            v10 := GenS1{F0: 2147483647, F1: -128}
            _ = v10
        }
    }
    if true {
        switch ("hello" + "") {
        case "":
            switch "" {
            case "\n\t":
                if (false && true) {
                    var v11 float32 = 0.0
                    _ = v11
                } else {
                    if false {
                        p0 -= p0
                        p1 += p1
                        p0 = (p0 - p0)
                    } else {
                        v12 := []int{127}
                        _ = v12
                    }
                    v13 := []int{-84}
                    if (((3.14159 * 0.1) / 1e10) != 0.1) {
                        var v14 int = 42
                        p1 += (p1 >> ((7) & 63))
                        _ = v14
                    }
                    _ = v13
                }
            case "":
                for fi15 := 0; fi15 < 4; fi15++ {
                    p0 = (p0 / (((p0 / (((p0 / (((p0 & -9223372036854775808)) | 1))) | 1))) | 1))
                    for fi16 := 0; fi16 < 11; fi16++ {
                        v17 := []int{993, (fi15 | fi16)}
                        fmt.Println(p1, p1, p1)
                        p0 = -1
                        _ = fi16
                        _ = v17
                    }
                    for fi18 := 0; fi18 < 17; fi18++ {
                        v19 := GenS1{F0: 65535, F1: -1}
                        _ = fi18
                        _ = v19
                    }
                    _ = fi15
                }
            case "hello":
                var v20 string = "foo bar"
                _ = v20
            default:
                for fi21 := 0; fi21 < 16; fi21++ {
                    var v22 bool = (float32(float64(-2.25)) <= 1.5)
                    _ = fi21
                    _ = v22
                }
            }
        case "a":
            fmt.Println(p0)
        case "hello":
            if false {
                v23 := GenS1{F0: 2, F1: 0}
                for fi24 := 0; fi24 < 1; fi24++ {
                    p0 = p0
                    _ = fi24
                }
                _ = v23
            }
        }
    } else {
        var v25 int8 = 100
        switch -576 {
        case -252:
            fmt.Println(v25, p0, p0)
        case -835:
            v26 := []int{42, 42, 4294967295}
            _ = v26
        case 0:
            switch "\n\t" {
            case "a":
                p0 -= (p0 >> ((32) & 63))
            case "hello":
                var v27 int8 = (v25 / ((49) | 1))
                _ = v27
            case "\n\t":
                var v28 bool = (0 != 0)
                _ = v28
            default:
                var v29 bool = (1 < -108)
                _ = v29
            }
        }
        _ = v25
    }
    _ = p0
    _ = p1
    return 42
}

func main() {
    v1 := make(map[int]string)
    var v2 int8 = 100
    v3 := make(map[int]uint32)
    var v4 int8 = (v2 * v2)
    v5 := []string{(("foo bar" + "hello") + "foo bar")}
    var v6 float32 = 1.5
    v30 := GenS2{F0: -2147483648, F1: 32768}
    var v31 string = "\n\t"
    var v32 uint = 65535
    if ((4294967295 == 113) && true) {
        v6 = 0.1
    }
    if len(v5) > 0 {
        v5[0] = v31
    }
    for fi33 := 0; fi33 < 8; fi33++ {
        var v34 int8 = 42
        for fi35 := 0; fi35 < 12; fi35++ {
            fmt.Println(v34, fi33)
            v36 := GenS3{F0: 1, F1: 42, F2: 9223372036854775807}
            v2 = v34
            _ = fi35
            _ = v36
        }
        _ = fi33
        _ = v34
    }
    v37 := GenS1{F0: 418, F1: v2}
    for fi38 := 0; fi38 < 20; fi38++ {
        v6 -= 3.14159
        if false {
            for fi39 := 0; fi39 < 3; fi39++ {
                switch -1 {
                case 0:
                    v2 = 1
                }
                if (!(-2.25 <= (0.0 - (-2.25 * ((0.1 * (0.0 * (0.1 * float64(v6)))) + 3.14159))))) {
                    v6 = 3.14159
                    fmt.Println(v4, fi38)
                }
                _ = fi39
            }
            if ((genFn1(-979, 0) > 128) && (false || false)) {
                var v40 float64 = 0.0
                fmt.Println(v31, v2)
                _ = v40
            }
        } else {
            v41 := GenS1{F0: 640, F1: (v2 << ((0) & 63))}
            _ = v41
        }
        fmt.Println(v32, v31)
        _ = fi38
    }
    if true {
        if ((v32 / ((4294967295) | 1)) != v32) {
            v1[128] = v31
            v1[-2147483648] = v31
        }
    }
    if false {
        var v42 int8 = v2
        v2 = (v2 / ((v42) | 1))
        v43 := GenS3{F0: 0, F1: 65535, F2: 605}
        _ = v42
        _ = v43
    }
    v4 -= (^(v4))
    fmt.Println(v4, v6, v4)
    v44 := GenS3{F0: 0, F1: 65535, F2: 42}
    var v45 int8 = -128
    for fi46 := 0; fi46 < 1; fi46++ {
        for fi47 := 0; fi47 < 18; fi47++ {
            if false {
                var v48 uint8 = 42
                _ = v48
            } else {
                _ = len(v5)
            }
            _ = fi47
        }
        switch (fi46 % (((fi46 ^ (fi46 / ((-364) | 1)))) | 1)) {
        case 1:
            var v49 int = (fi46 >> (((^(v32))) & 63))
            _ = v49
        case -9223372036854775808:
            for fi50 := 0; fi50 < 13; fi50++ {
                v51 := GenS2{F0: 1, F1: int(v4)}
                _ = fi50
                _ = v51
            }
        case 32768:
            switch v31 {
            case "":
                if false {
                    var v52 uint32 = 42
                    for fi53 := 0; fi53 < 15; fi53++ {
                        v31 = v31
                        if ((fi46 - fi46) < fi46) {
                            v32 = (^(v32))
                        }
                        fmt.Println(v4)
                        _ = fi53
                    }
                    _ = v52
                } else {
                    v32 = (v32 >> (((v32 | (v32 % ((2) | 1)))) & 63))
                    v2 -= (v45 & (-(v4)))
                    var v54 int16 = 1
                    _ = v54
                }
            case "foo bar":
                var v55 float64 = (float64(1e10) - (0.0 - 3.14159))
                _ = v55
            case "hello":
                for fi56 := 0; fi56 < 20; fi56++ {
                    if false {
                        v57 := v3[(fi46 | 42)]
                        _ = v57
                    }
                    var v58 uint = v32
                    _ = fi56
                    _ = v58
                }
            default:
                v45 -= v2
            }
        }
        if (!false) {
            for fi59 := 0; fi59 < 5; fi59++ {
                var v60 int = -1
                v4 -= (v4 >> ((v32) & 63))
                _ = fi59
                _ = v60
            }
            var v61 float32 = v6
            if (42 > 42) {
                for fi62 := 0; fi62 < 19; fi62++ {
                    var v63 int8 = 42
                    if true {
                        v61 -= ((v6 * float32(1.5)) + ((v61 + v61) + ((((float32(3.14159) - (float32(3.14159) - ((v61 - (v6 / (float32(3.14159) + v61))) - 0.0))) / (v61 * 0.1)) + ((((3.14159 - v6) / (float32(1e10) - (1e10 / v6))) + 1.5) - (float32((((-2.25 + 3.14159) / (3.14159 - float64((v61 / float32(((float64((float32(3.14159) * v6)) * float64((((v61 + (((v61 - 1.5) + float32(0.1)) / 0.1)) + 1e10) / v61))) + -2.25)))))) - 1e10)) - (v61 - 1.5)))) - 1.5)))
                        var v64 uint64 = uint64(v32)
                        if (2147483647 > 610) {
                            var v65 int32 = -566
                            _ = v65
                        } else {
                            v66 := GenS4{F0: "\n\t"}
                            if ((^(v64)) > (^(v64))) {
                                var v67 float64 = -2.25
                                _ = v67
                            }
                            _ = v66
                        }
                        var v68 bool = true
                        switch (v31 + (v31 + "a")) {
                        case "":
                            v69 := GenS4{F0: "a"}
                            _ = v69
                        default:
                            if (-32768 <= -428) {
                                var v70 uint64 = v64
                                v3[(fi62 & fi46)] = 1
                                if v68 {
                                    _ = len(v3)
                                    if v68 {
                                        if v68 {
                                            v71 := GenS3{F0: 42, F1: 431, F2: (^(fi46))}
                                            _ = v71
                                        }
                                    } else {
                                        v63 = (v4 * v4)
                                    }
                                } else {
                                    var v72 float64 = 0.0
                                    _ = v72
                                }
                                _ = v70
                            } else {
                                if false {
                                    v61 = ((float32(((0.1 / 1.0) * 3.14159)) * -2.25) / (0.1 + v61))
                                    v73 := GenS2{F0: -34, F1: fi46}
                                    _ = v73
                                }
                            }
                        }
                        _ = v64
                        _ = v68
                    }
                    v6 -= -2.25
                    _ = fi62
                    _ = v63
                }
                if true {
                    if true {
                        var v74 uint8 = 163
                        _ = v74
                    } else {
                        v3[-2147483648] = 2
                        v31 = ""
                        switch (v31 + v31) {
                        case "a":
                            _ = len(v5)
                        default:
                            if (4294967295 >= 804) {
                                v75 := GenS2{F0: -170, F1: (fi46 >> ((v32) & 63))}
                                _ = v75
                            }
                        }
                    }
                } else {
                    for fi76 := 0; fi76 < 11; fi76++ {
                        var v77 uint64 = 0
                        if true {
                            var v78 bool = (-32768 >= 1)
                            switch (fi76 & (-(fi46))) {
                            case 753:
                                if ((!true) || v78) {
                                    fmt.Println(v61, v32)
                                    v4 = (v45 % (((v2 - 127)) | 1))
                                }
                            case 128:
                                var v79 uint64 = 42
                                _ = v79
                            default:
                                switch 127 {
                                case 9223372036854775807:
                                    v77 = v77
                                case -361:
                                    fmt.Println(v32, v6)
                                    var v80 string = v31
                                    var v81 uint8 = 127
                                    _ = v80
                                    _ = v81
                                case 821:
                                    v82 := GenS5{F0: v31}
                                    _ = v82
                                default:
                                    var v83 int8 = 52
                                    _ = v83
                                }
                            }
                            v61 = float32(0.1)
                            _ = v78
                        } else {
                            if true {
                                fmt.Println(fi46, v4, v32)
                            }
                            if true {
                                var v84 float64 = 1.5
                                v85 := GenS3{F0: 32767, F1: 42, F2: fi46}
                                _ = v84
                                _ = v85
                            } else {
                                v86 := GenS6{F0: 148, F1: (v32 | (v32 ^ 4294967295)), F2: false}
                                var v87 bool = true
                                v5 = append(v5, "hello")
                                v5 = append(v5, v31)
                                _ = v86
                                _ = v87
                            }
                            v77 = 18446744073709551615
                        }
                        _ = fi76
                        _ = v77
                    }
                }
            }
            _ = v61
        } else {
            for fi88 := 0; fi88 < 13; fi88++ {
                var v89 uint32 = 2147483648
                var v90 uint8 = 42
                _ = fi88
                _ = v89
                _ = v90
            }
        }
        _ = fi46
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v37
    _ = v4
    _ = v44
    _ = v45
    _ = v5
    _ = v6
}
