package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 string; F1 int16; F2 int8 }
type GenS4 struct { F0 int64; F1 uint8 }
type GenS5 struct { F0 bool; F1 int16; F2 uint }
type GenS6 struct { F0 bool; F1 int16; F2 string }
func genFn0() uint16 {
    if false {
        if true {
            v6 := GenS1{F0: 432, F1: 1}
            var v7 uint8 = 2
            _ = v6
            _ = v7
        } else {
            v8 := []int{128, 4294967295, 406, 584}
            var v9 float64 = -2.25
            _ = v8
            _ = v9
        }
    }
    for fi10 := 0; fi10 < 20; fi10++ {
        for fi11 := 0; fi11 < 18; fi11++ {
            v12 := GenS1{F0: 42, F1: 0}
            if (30 < 885) {
                v13 := GenS2{F0: -32768}
                var v14 int = fi11
                var v15 int64 = 814
                v16 := GenS2{F0: 1}
                _ = v13
                _ = v14
                _ = v15
                _ = v16
            }
            fmt.Println(fi11, fi10)
            var v17 uint = 0
            v18 := GenS1{F0: 255, F1: 0}
            if (!true) {
                v17 -= v17
                var v19 uint = 18446744073709551615
                v19 += (v17 / ((v19) | 1))
                _ = v19
            } else {
                v20 := GenS3{F0: "", F1: 32767, F2: int8(v17)}
                fmt.Println(fi10, fi10, fi10)
                _ = v20
            }
            var v21 bool = (!(-2.25 <= (-2.25 / 3.14159)))
            _ = fi11
            _ = v12
            _ = v17
            _ = v18
            _ = v21
        }
        _ = fi10
    }
    if (65535 != 2147483648) {
        var v22 int8 = 127
        _ = v22
    }
    return 0
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[int]int)
    v3 := make(map[int]bool)
    var v4 float32 = 1.5
    var v5 uint8 = 0
    v23 := GenS4{F0: -9223372036854775808, F1: 2}
    switch 32767 {
    case -826:
        var v24 float64 = 3.14159
        _ = v24
    case 405:
        var v25 float32 = 1.5
        _ = v25
    case 32768:
        var v26 uint16 = genFn0()
        _ = v26
    }
    var v27 uint = 227
    for fi28 := 0; fi28 < 2; fi28++ {
        var v29 uint = v27
        for fi30 := 0; fi30 < 8; fi30++ {
            _ = len(v1)
            v4 = 0.0
            _ = fi30
        }
        _ = fi28
        _ = v29
    }
    switch 60 {
    case 9223372036854775807:
        var v31 uint8 = v5
        _ = v31
    }
    for fi32 := 0; fi32 < 14; fi32++ {
        fmt.Println(v27, v4)
        var v33 float32 = 3.14159
        for fi34 := 0; fi34 < 10; fi34++ {
            var v35 float32 = v4
            v36 := []int{(fi32 ^ fi34), fi34}
            v37 := GenS1{F0: 1, F1: 1}
            _ = fi34
            _ = v35
            _ = v36
            _ = v37
        }
        _ = fi32
        _ = v33
    }
    v4 = 0.1
    fmt.Println(v5)
    if (true && false) {
        if false {
            for fi38 := 0; fi38 < 4; fi38++ {
                if false {
                    switch "hello" {
                    case "\n\t":
                        fmt.Println(v5, v4, v5)
                    case "hello":
                        _ = len(v3)
                    default:
                        for fi39 := 0; fi39 < 4; fi39++ {
                            v4 += v4
                            _ = fi39
                        }
                    }
                } else {
                    if (0 >= 2147483648) {
                        v3[fi38] = (!(("foo bar" + "a") > "hello"))
                        v3[(fi38 - (^(fi38)))] = (uint64(v27) <= 18446744073709551615)
                    } else {
                        if (!false) {
                            switch "\n\t" {
                            case "":
                                v40 := GenS3{F0: "hello", F1: 1000, F2: -124}
                                _ = v40
                            case "hello":
                                if false {
                                    if ("" < "") {
                                        for fi41 := 0; fi41 < 19; fi41++ {
                                            v5 = (^(v5))
                                            _ = fi41
                                        }
                                    } else {
                                        v27 -= (v27 << ((0) & 63))
                                        for fi42 := 0; fi42 < 5; fi42++ {
                                            var v43 uint = v27
                                            if false {
                                                var v44 uint32 = uint32(v5)
                                                _ = v44
                                            } else {
                                                fmt.Println(v5, fi38, fi42)
                                                if false {
                                                    if (!(!((2 == 65535) && true))) {
                                                        v45 := []int{}
                                                        fmt.Println(v27, v27, v5)
                                                        _ = v45
                                                    }
                                                    var v46 string = "hello"
                                                    v5 = v5
                                                    _ = len(v3)
                                                    var v47 int64 = 42
                                                    _ = v46
                                                    _ = v47
                                                }
                                            }
                                            _ = fi42
                                            _ = v43
                                        }
                                    }
                                }
                            }
                            v48 := GenS2{F0: -1}
                            v49 := GenS1{F0: 65535, F1: 127}
                            _ = v48
                            _ = v49
                        } else {
                            v5 += (v5 ^ 1)
                            fmt.Println(v27, fi38)
                        }
                        v27 += v27
                        v4 = (0.0 * v4)
                    }
                    if (2 == 1) {
                        v5 += v5
                        var v50 int = (^(fi38))
                        for fi51 := 0; fi51 < 8; fi51++ {
                            v52 := []int{fi38}
                            if (2 >= 2147483647) {
                                v5 += v5
                                v53 := GenS3{F0: "\n\t", F1: 42, F2: 120}
                                fmt.Println(v4, fi51, v5)
                                v4 = float32(1e10)
                                _ = v53
                            } else {
                                var v54 int32 = 0
                                v4 = v4
                                var v55 uint8 = 230
                                v3[-164] = (4294967295 < 2147483648)
                                _ = v54
                                _ = v55
                            }
                            var v56 float64 = (((0.0 * ((0.1 / (float64((0.1 + (1e10 - ((v4 / 1.0) / v4)))) * -2.25)) + -2.25)) / (3.14159 + -2.25)) / (float64(1e10) + ((((float64(v4) * (0.0 / (1.5 - 3.14159))) * float64(float32(1e10))) / 1e10) + 1.5)))
                            _ = fi51
                            _ = v52
                            _ = v56
                        }
                        _ = v50
                    }
                }
                switch fi38 {
                case 9223372036854775807:
                    fmt.Println(v4, v4, fi38)
                case 4294967295:
                    if (-29 >= -128) {
                        delete(v3, (-(fi38)))
                        var v57 uint = (v27 & v27)
                        for fi58 := 0; fi58 < 14; fi58++ {
                            v57 = (v57 / (((v57 >> ((v27) & 63))) | 1))
                            _ = fi58
                        }
                        _ = v57
                    }
                case -1:
                    v4 += (0.0 - 0.1)
                default:
                    for fi59 := 0; fi59 < 14; fi59++ {
                        v4 -= (v4 * v4)
                        _ = fi59
                    }
                }
                _ = fi38
            }
            var v60 uint = v27
            var v61 uint8 = v5
            _ = v60
            _ = v61
        } else {
            for fi62 := 0; fi62 < 5; fi62++ {
                if (!((858 != -2147483648) || (!false))) {
                    v63 := GenS2{F0: 0}
                    _ = v63
                } else {
                    v27 = v27
                    var v64 int16 = 254
                    _ = v64
                }
                v27 = 42
                _ = fi62
            }
        }
        var v65 uint = (v27 / (((^(v27))) | 1))
        _ = v65
    }
    if ((v27 ^ v27) >= v27) {
        var v66 int64 = -9223372036854775808
        var v67 uint = 130
        _ = v66
        _ = v67
    } else {
        switch 127 {
        case -2147483648:
            for fi68 := 0; fi68 < 18; fi68++ {
                if (32767 != 734) {
                    for fi69 := 0; fi69 < 9; fi69++ {
                        switch "\n\t" {
                        case "foo bar":
                            v27 = 51
                            if ((-(fi68)) >= fi68) {
                                var v70 uint16 = uint16(v5)
                                switch 32767 {
                                case 32767:
                                    switch fi68 {
                                    case 0:
                                        v70 -= (v70 << ((294) & 63))
                                        v70 = (^(v70))
                                    default:
                                        v71 := GenS2{F0: -874}
                                        _ = v71
                                    }
                                case 1:
                                    v27 -= v27
                                }
                                _ = v70
                            }
                            v4 += v4
                        }
                        v27 += (v27 * (^(v27)))
                        var v72 uint = (^(v27))
                        _ = fi69
                        _ = v72
                    }
                    for fi73 := 0; fi73 < 8; fi73++ {
                        fmt.Println(v5, fi73)
                        _ = fi73
                    }
                }
                switch (("hello" + "hello") + "") {
                case "\n\t":
                    v74 := GenS2{F0: 32767}
                    _ = v74
                case "foo bar":
                    var v75 uint = 4294967295
                    _ = v75
                default:
                    for fi76 := 0; fi76 < 18; fi76++ {
                        v27 = v27
                        _ = fi76
                    }
                }
                _ = fi68
            }
        case 289:
            switch "" {
            case "a":
                v77 := GenS2{F0: 1000}
                _ = v77
            case "foo bar":
                if (!false) {
                    var v78 float32 = v4
                    v5 = v5
                    _ = v78
                } else {
                    for fi79 := 0; fi79 < 5; fi79++ {
                        v4 += (float32((((((0.1 * (3.14159 + float64(0.1))) - (float64(v4) + 1e10)) + 1.5) - 1e10) * float64(0.1))) - v4)
                        v80 := GenS1{F0: 2, F1: 1}
                        _ = fi79
                        _ = v80
                    }
                }
            }
        case 290:
            var v81 int8 = 0
            _ = v81
        }
        v82 := []int{1, 4294967295}
        _ = v82
    }
    v27 += (v27 * v27)
    switch 631 {
    case 32768:
        v3[32768] = false
    }
    v5 = v5
    for fi83 := 0; fi83 < 19; fi83++ {
        var v84 int64 = 573
        if true {
            v85 := []int{fi83, 721, 9223372036854775807, fi83}
            for fi86 := 0; fi86 < 16; fi86++ {
                fmt.Println(v4)
                var v87 string = "foo bar"
                _ = fi86
                _ = v87
            }
            for fi88 := 0; fi88 < 8; fi88++ {
                if (!true) {
                    v5 += v5
                    if (false && (42 == int32(v27))) {
                        fmt.Println(v4, v84)
                    } else {
                        var v89 float64 = 1e10
                        if len(v85) > 0 {
                            v85[0] = fi88
                        }
                        v5 = (^(v5))
                        if len(v85) >= 2 {
                            v85 = v85[1:len(v85)]
                        }
                        _ = v89
                    }
                    v27 = v27
                }
                _ = fi88
            }
            _ = v85
        }
        for fi90 := 0; fi90 < 19; fi90++ {
            fmt.Println(v27)
            _ = fi90
        }
        _ = fi83
        _ = v84
    }
    if ((false && false) || true) {
        if false {
            if (-32768 != 1000) {
                v4 = v4
                for fi91 := 0; fi91 < 14; fi91++ {
                    v27 = (v27 / ((18446744073709551615) | 1))
                    for fi92 := 0; fi92 < 6; fi92++ {
                        fmt.Println(v27, fi91, v4)
                        v93 := GenS2{F0: -32768}
                        _ = fi92
                        _ = v93
                    }
                    _ = fi91
                }
            } else {
                fmt.Println(v27, v27, v4)
                v94 := GenS5{F0: (!false), F1: 32767, F2: v27}
                v4 -= (v4 / (((3.14159 + v4) * v4) - -2.25))
                _ = v94
            }
            var v95 uint16 = 854
            _ = v95
        } else {
            var v96 uint16 = 42
            _ = v96
        }
        v97 := GenS6{F0: false, F1: 1000, F2: "\n\t"}
        var v98 uint32 = uint32(v27)
        _ = v97
        _ = v98
    } else {
        if true {
            v27 += (v27 & (v27 / ((v27) | 1)))
        }
        v4 = (3.14159 - float32(1e10))
        v4 -= 0.0
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
