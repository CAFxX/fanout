package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64; F2 int32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int; F1 uint16; F2 bool }
type GenS4 struct { F0 int8; F1 bool }
type GenS5 struct { F0 string; F1 uint32 }
type GenS6 struct { F0 bool; F1 bool }
type GenS7 struct { F0 int32; F1 uint16 }
type GenS8 struct { F0 int16; F1 uint8 }
type GenS9 struct { F0 uint64; F1 int }
type GenS10 struct { F0 int; F1 int8 }
type GenS11 struct { F0 uint64; F1 uint32; F2 bool }
func genFn0(p0 int8, p1 int16) int32 {
    var v5 uint64 = 764
    v5 += v5
    _ = p0
    _ = p1
    _ = v5
    return 1
}
func genFn1() string {
    v6 := GenS1{F0: -57, F1: 0, F2: 2147483647}
    v7 := []int{1, -1, 128}
    _ = v6
    _ = v7
    return "foo bar"
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[int]string)
    var v3 uint8 = 128
    v4 := make(map[int]bool)
    if (!false) {
        var v8 uint = 255
        v9 := GenS2{F0: 989}
        var v10 bool = false
        _ = v10
        _ = v8
        _ = v9
    }
    for fi11 := 0; fi11 < 5; fi11++ {
        var v12 bool = true
        var v13 int32 = -2147483648
        var v14 uint16 = 2
        _ = fi11
        _ = v12
        _ = v13
        _ = v14
    }
    fmt.Println(v3)
    if (false || ((!false) && true)) {
        fmt.Println(v3)
        fmt.Println(v3)
        for fi15 := 0; fi15 < 15; fi15++ {
            delete(v2, 2147483647)
            _ = fi15
        }
    }
    if (float32(0.0) == 3.14159) {
        v3 -= (v3 + (v3 * v3))
        if (32768 >= -4) {
            switch -2147483648 {
            case 1:
                var v16 int8 = -1
                _ = v16
            case 9223372036854775807:
                _ = len(v2)
            case 60:
                var v17 uint8 = 255
                _ = v17
            default:
                switch 0 {
                case -2147483648:
                    var v18 int8 = 14
                    _ = v18
                case -496:
                    v3 = v3
                }
            }
            fmt.Println(v3)
        } else {
            switch -480 {
            case 1:
                if true {
                    v3 -= v3
                } else {
                    _ = len(v4)
                }
            case 9223372036854775807:
                v19 := GenS3{F0: -88, F1: 2, F2: (v3 == (v3 << ((18446744073709551615) & 63)))}
                _ = v19
            default:
                var v20 int32 = 0
                _ = v20
            }
            for fi21 := 0; fi21 < 1; fi21++ {
                v3 = 127
                var v22 string = ""
                _ = fi21
                _ = v22
            }
        }
        v23 := []int{9223372036854775807, -9223372036854775808, 128, 2147483647}
        _ = v23
    } else {
        for fi24 := 0; fi24 < 7; fi24++ {
            var v25 int8 = 100
            switch fi24 {
            case 127:
                switch "" {
                case "":
                    var v26 int = (fi24 - fi24)
                    _ = v26
                case "foo bar":
                    if true {
                        if true {
                            v27 := GenS2{F0: 32767}
                            _ = v27
                        }
                    } else {
                        v3 -= v3
                        v28 := GenS4{F0: v25, F1: false}
                        _ = v28
                    }
                case "a":
                    fmt.Println(fi24, fi24, v25)
                default:
                    fmt.Println(v25, v3)
                }
            case -886:
                v29 := []int{(fi24 - fi24), fi24, 127}
                _ = v29
            default:
                v30 := GenS5{F0: "", F1: 349}
                _ = v30
            }
            _ = fi24
            _ = v25
        }
        v3 -= (^(v3))
        switch -714 {
        case 128:
            if (4294967295 < 0) {
                var v31 bool = (!(927 > 42))
                _ = v31
            }
        case 2147483647:
            v3 = 128
        case -1:
            v3 = 0
        default:
            var v32 uint16 = 2
            _ = v32
        }
    }
    v3 -= v3
    switch ("hello" + "a") {
    case "hello":
        switch ("foo bar" + ((("a" + "hello") + ("" + "foo bar")) + ("hello" + "foo bar"))) {
        case "hello":
            v3 = v3
        case "foo bar":
            if (!(!(!(42 > -1)))) {
                v2[-901] = ("\n\t" + "a")
            }
        case "\n\t":
            fmt.Println(v3)
        }
    default:
        for fi33 := 0; fi33 < 13; fi33++ {
            fmt.Println(v3, fi33)
            var v34 float32 = float32(((0.1 + 0.0) - 1.5))
            _ = fi33
            _ = v34
        }
    }
    if false {
        for fi35 := 0; fi35 < 16; fi35++ {
            for fi36 := 0; fi36 < 5; fi36++ {
                if (!(-2.25 <= (0.0 - (((0.1 / 1.0) / ((1e10 / 1.5) / (1.5 * 0.1))) * float64(3.14159))))) {
                    if false {
                        switch 972 {
                        case -9223372036854775808:
                            var v37 uint32 = 65535
                            _ = v37
                        case 42:
                            if true {
                                var v38 int16 = -529
                                switch (fi36 << ((uint(v3)) & 63)) {
                                case 9223372036854775807:
                                    v39 := GenS6{F0: (127 <= v3), F1: true}
                                    _ = v39
                                case 4294967295:
                                    switch -2147483648 {
                                    case 4294967295:
                                        var v40 uint32 = 222
                                        _ = v40
                                    case -683:
                                        var v41 uint = uint(v3)
                                        _ = v41
                                    case 42:
                                        v42 := GenS4{F0: -128, F1: false}
                                        _ = v42
                                    }
                                case 32767:
                                    fmt.Println(fi35)
                                default:
                                    var v43 uint8 = v3
                                    _ = v43
                                }
                                _ = v38
                            } else {
                                v44 := []int{}
                                _ = v44
                            }
                            var v45 bool = (!false)
                            v46 := GenS7{F0: -364, F1: 1}
                            fmt.Println(fi36)
                            v47 := []int{}
                            _ = v45
                            _ = v46
                            _ = v47
                        case 1:
                            v3 = v3
                        default:
                            switch (fi35 << ((976) & 63)) {
                            case 9223372036854775807:
                                v48 := GenS8{F0: 0, F1: 2}
                                v49 := GenS9{F0: 0, F1: (fi35 + (fi36 & 1))}
                                _ = v48
                                _ = v49
                            case -9223372036854775808:
                                if true {
                                    v50 := GenS10{F0: fi35, F1: 1}
                                    var v51 float64 = 0.1
                                    var v52 float32 = (0.0 - (0.0 + (1e10 / 3.14159)))
                                    _ = v50
                                    _ = v51
                                    _ = v52
                                } else {
                                    var v53 bool = true
                                    var v54 int16 = -32768
                                    if ((32767 <= (-(v54))) || (9223372036854775807 > 42)) {
                                        fmt.Println(fi36, v53, v54)
                                        v53 = v53
                                    } else {
                                        v55 := GenS1{F0: 42, F1: 1, F2: 0}
                                        v56 := GenS4{F0: 42, F1: v53}
                                        v54 += v54
                                        _ = v55
                                        _ = v56
                                    }
                                    v3 -= v3
                                    var v57 float32 = 0.0
                                    _ = v53
                                    _ = v54
                                    _ = v57
                                }
                            default:
                                if true {
                                    switch "" {
                                    case "foo bar":
                                        v3 -= v3
                                    default:
                                        var v58 float32 = ((((0.1 - (float32(((1e10 / (float64(float32(0.0)) + float64(3.14159))) * 0.0)) - (0.1 / 1.0))) - ((float32((0.1 * (((((1.5 * 3.14159) / 1.0) / -2.25) - ((0.1 / 0.1) / 1.5)) * (-2.25 + (-2.25 + float64((float32(-2.25) + -2.25))))))) - 1.5) + (-2.25 / ((1.5 / (1.5 - (float32((0.1 / 1.0)) + 1e10))) + (0.0 * 1.5))))) * 0.1) / float32((3.14159 + float64(1.5))))
                                        _ = v58
                                    }
                                    v59 := GenS1{F0: -35, F1: 0, F2: -409}
                                    _ = v59
                                }
                                var v60 float64 = 1.5
                                _ = v60
                            }
                        }
                    } else {
                        if false {
                            v3 = v3
                        } else {
                            var v61 uint64 = 2
                            var v62 int32 = 42
                            v63 := v4[(fi36 - fi36)]
                            _ = v61
                            _ = v62
                            _ = v63
                        }
                        delete(v4, (^(fi35)))
                    }
                } else {
                    v3 = 2
                    v64 := GenS8{F0: -320, F1: (v3 - v3)}
                    _ = v64
                }
                _ = fi36
            }
            v4[fi35] = (1 < 2)
            var v65 int8 = -128
            _ = fi35
            _ = v65
        }
    } else {
        v66 := v4[132]
        if true {
            v67 := GenS9{F0: 2, F1: -367}
            switch 42 {
            case 0:
                if (307 != 1) {
                    var v68 int32 = -499
                    _ = v68
                }
            case 922:
                v1[2147483648] = 4294967295
            case 987:
                var v69 uint8 = (v3 >> ((4294967295) & 63))
                _ = v69
            default:
                if (1 < 127) {
                    v70 := GenS5{F0: "hello", F1: 4294967295}
                    delete(v4, -2147483648)
                    delete(v4, 127)
                    _ = v70
                } else {
                    v71 := []int{32768, 892, 539}
                    var v72 uint8 = v3
                    var v73 string = ""
                    _ = v71
                    _ = v72
                    _ = v73
                }
            }
            _ = len(v1)
            _ = v67
        }
        v66 = v66
        _ = v66
    }
    var v74 int32 = 0
    for fi75 := 0; fi75 < 10; fi75++ {
        if false {
            v3 = 128
            v76 := v4[(fi75 * fi75)]
            _ = v76
        }
        fmt.Println(v3, v3, v74)
        _ = fi75
    }
    delete(v1, 4294967295)
    if (-109 == -128) {
        for fi77 := 0; fi77 < 18; fi77++ {
            var v78 bool = (1e10 == (((0.1 * (0.0 / 1e10)) - (float64(0.0) / ((((1.5 - -2.25) + 3.14159) + (0.1 / float64((1e10 - 1.5)))) - 0.0))) / 1.0))
            switch "\n\t" {
            case "":
                v79 := []int{(fi77 & (^(fi77))), (fi77 >> ((720) & 63)), (fi77 & -2147483648), fi77}
                _ = v79
            case "foo bar":
                v3 += v3
            case "a":
                v3 = v3
            }
            var v80 int64 = 42
            _ = fi77
            _ = v78
            _ = v80
        }
        v81 := GenS11{F0: 1, F1: 101, F2: (969 > 42)}
        v82 := GenS3{F0: -777, F1: 152, F2: true}
        _ = v81
        _ = v82
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v74:", v74)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v74
}
