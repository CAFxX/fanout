package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32; F2 string }
type GenS2 struct { F0 uint8; F1 bool; F2 string }
type GenS3 struct { F0 string; F1 int16; F2 uint8 }
func genFn0(p0 uint8, p1 int) int16 {
    p1 = p1
    var v7 string = "foo bar"
    _ = p0
    _ = p1
    _ = v7
    return -1
}
func genFn1() int64 {
    var v8 uint16 = 42
    fmt.Println(v8)
    for fi9 := 0; fi9 < 19; fi9++ {
        if false {
            var v10 string = "\n\t"
            for fi11 := 0; fi11 < 9; fi11++ {
                var v12 int16 = 148
                _ = fi11
                _ = v12
            }
            for fi13 := 0; fi13 < 9; fi13++ {
                switch fi9 {
                case 473:
                    v8 = 889
                case -2147483648:
                    v14 := []int{fi9, -867, (^(fi9))}
                    _ = v14
                case 183:
                    if ((82 > 0) || true) {
                        if (806 > 18446744073709551615) {
                            var v15 float32 = 1e10
                            var v16 float32 = float32((0.0 - 3.14159))
                            _ = v15
                            _ = v16
                        }
                        v17 := []int{}
                        v8 = v8
                        if (3.14159 == 1e10) {
                            fmt.Println(fi9, v10)
                            var v18 string = "a"
                            _ = v18
                        }
                        v19 := GenS1{F0: v8, F1: int32(fi13), F2: "hello"}
                        switch v10 {
                        case "a":
                            var v20 int32 = 0
                            _ = v20
                        case "foo bar":
                            if (42 <= v8) {
                                v21 := GenS1{F0: v8, F1: 2147483647, F2: "hello"}
                                _ = v21
                            } else {
                                var v22 bool = true
                                v8 = 283
                                v22 = (127 < 178)
                                _ = v22
                            }
                        default:
                            fmt.Println(v8, fi13)
                        }
                        _ = v17
                        _ = v19
                    } else {
                        v10 = v10
                        v10 = ""
                    }
                default:
                    v10 = "\n\t"
                }
                fmt.Println(fi9)
                _ = fi13
            }
            _ = v10
        } else {
            var v23 int8 = -58
            v24 := []int{9223372036854775807, fi9}
            _ = v23
            _ = v24
        }
        v8 -= (v8 << ((42) & 63))
        _ = fi9
    }
    for fi25 := 0; fi25 < 14; fi25++ {
        if true {
            var v26 bool = (true && ((false || ((652 > 18446744073709551615) && true)) && (!false)))
            _ = v26
        }
        for fi27 := 0; fi27 < 18; fi27++ {
            v8 -= (v8 / (((v8 * 65535)) | 1))
            v8 += v8
            _ = fi27
        }
        if (!(!true)) {
            v8 -= 2
            for fi28 := 0; fi28 < 20; fi28++ {
                var v29 string = "hello"
                v29 = v29
                if (2 < 184) {
                    var v30 bool = (-9223372036854775808 == 9223372036854775807)
                    v31 := []int{(fi28 - -2147483648), fi25}
                    v30 = v30
                    var v32 int = 127
                    _ = v30
                    _ = v31
                    _ = v32
                } else {
                    v29 = "a"
                    var v33 float32 = 0.0
                    if false {
                        v8 = (v8 * (^(v8)))
                        v33 = float32(0.1)
                        v8 += v8
                    }
                    _ = v33
                }
                _ = fi28
                _ = v29
            }
        } else {
            v8 = (v8 | 0)
        }
        _ = fi25
    }
    fmt.Println(v8)
    _ = v8
    return 9223372036854775807
}

func main() {
    var v1 string = "foo bar"
    var v2 int8 = 100
    var v3 uint16 = 42
    var v4 uint16 = v3
    v5 := make(map[uint32]uint32)
    var v6 uint32 = 4294967295
    if ((-(v2)) < v2) {
        fmt.Println(v3)
        if (892 <= 0) {
            var v34 float32 = (3.14159 * (0.0 / 1e10))
            if (v1 <= ("\n\t" + v1)) {
                v34 -= float32(0.1)
                v4 = (^(v3))
            }
            switch ("a" + v1) {
            case "hello":
                var v35 uint16 = (v4 | v4)
                _ = v35
            case "a":
                switch v1 {
                case "":
                    v34 = (0.0 * v34)
                case "\n\t":
                    if false {
                        switch "\n\t" {
                        case "hello":
                            fmt.Println(v2)
                        case "a":
                            if false {
                                var v36 uint32 = 4294967295
                                var v37 int16 = 32767
                                _ = v36
                                _ = v37
                            } else {
                                switch 9223372036854775807 {
                                case 1:
                                    switch v1 {
                                    case "":
                                        v6 = v6
                                    }
                                case -725:
                                    v6 -= (v6 + 1)
                                case 128:
                                    v3 -= (v4 / ((v4) | 1))
                                default:
                                    var v38 int32 = int32(v6)
                                    _ = v38
                                }
                                if true {
                                    v5[(v6 & v6)] = (v6 % ((v6) | 1))
                                } else {
                                    v1 = (v1 + ((v1 + (v1 + (("a" + (v1 + ("a" + ""))) + v1))) + "hello"))
                                }
                            }
                        case "":
                            fmt.Println(v6, v1, v6)
                        default:
                            for fi39 := 0; fi39 < 6; fi39++ {
                                v40 := GenS2{F0: 0, F1: true, F2: v1}
                                _ = fi39
                                _ = v40
                            }
                        }
                    }
                case "foo bar":
                    switch (((("a" + v1) + v1) + ("hello" + v1)) + (v1 + (v1 + "\n\t"))) {
                    case "":
                        var v41 int = 128
                        _ = v41
                    case "a":
                        if false {
                            if false {
                                if true {
                                    if false {
                                        if (!true) {
                                            for fi42 := 0; fi42 < 4; fi42++ {
                                                fmt.Println(v2, v2)
                                                _ = fi42
                                            }
                                            var v43 uint8 = 2
                                            _ = v43
                                        } else {
                                            v6 -= (v6 / ((uint32(v4)) | 1))
                                        }
                                    } else {
                                        var v44 float32 = v34
                                        v45 := GenS2{F0: 0, F1: ((!(true || (genFn0(uint8(v6), -461) < 32767))) || true), F2: v1}
                                        delete(v5, (v6 ^ v6))
                                        _ = v44
                                        _ = v45
                                    }
                                    fmt.Println(v1)
                                    if ((!false) && (42 >= 0)) {
                                        v5[(v6 ^ v6)] = (v6 ^ uint32(v4))
                                    } else {
                                        if true {
                                            if true {
                                                var v46 float32 = v34
                                                _ = v46
                                            }
                                            if (1 <= 0) {
                                                if (true || true) {
                                                    var v47 uint = 0
                                                    _ = len(v5)
                                                    _ = v47
                                                } else {
                                                    var v48 uint64 = 211
                                                    v49 := GenS3{F0: (v1 + v1), F1: -1, F2: uint8(v4)}
                                                    for fi50 := 0; fi50 < 8; fi50++ {
                                                        delete(v5, (v6 + v6))
                                                        _ = fi50
                                                    }
                                                    _ = v48
                                                    _ = v49
                                                }
                                            }
                                        }
                                    }
                                } else {
                                    _ = len(v5)
                                    var v51 int = -9223372036854775808
                                    _ = v51
                                }
                                var v52 uint8 = 3
                                v3 = v3
                                _ = v52
                            }
                        } else {
                            v53 := []int{32768, -1}
                            _ = v53
                        }
                    }
                default:
                    v3 = (v4 / (((^(v3))) | 1))
                }
            case "foo bar":
                if (0 == 1) {
                    fmt.Println(v1, v6)
                    if (!true) {
                        for fi54 := 0; fi54 < 14; fi54++ {
                            if false {
                                for fi55 := 0; fi55 < 14; fi55++ {
                                    v56 := []int{fi55, fi54}
                                    _ = fi55
                                    _ = v56
                                }
                                for fi57 := 0; fi57 < 8; fi57++ {
                                    v58 := GenS3{F0: v1, F1: -1, F2: 0}
                                    var v59 int64 = 1
                                    _ = fi57
                                    _ = v58
                                    _ = v59
                                }
                                var v60 uint8 = 1
                                _ = v60
                            }
                            v6 = (v6 - v6)
                            v61 := GenS2{F0: uint8(v4), F1: true, F2: "foo bar"}
                            _ = fi54
                            _ = v61
                        }
                        v62 := v5[(v6 * v6)]
                        _ = v62
                    }
                    var v63 int64 = int64(v2)
                    _ = v63
                }
            }
            _ = v34
        } else {
            switch v1 {
            case "foo bar":
                var v64 uint16 = (v4 / ((602) | 1))
                _ = v64
            case "a":
                fmt.Println(v4)
            default:
                if true {
                    fmt.Println(v1, v2, v2)
                    v1 = "hello"
                    var v65 int = 127
                    _ = v65
                } else {
                    _ = len(v5)
                }
            }
        }
        if true {
            if (331 != 4294967295) {
                for fi66 := 0; fi66 < 13; fi66++ {
                    v67 := GenS1{F0: (^(v3)), F1: 0, F2: v1}
                    if true {
                        if false {
                            fmt.Println(v4, v1, v1)
                        } else {
                            _ = len(v5)
                            for fi68 := 0; fi68 < 4; fi68++ {
                                switch ("\n\t" + "") {
                                case "":
                                    switch "foo bar" {
                                    case "":
                                        v5[0] = (v6 & 4294967295)
                                    }
                                case "\n\t":
                                    v1 = v1
                                case "foo bar":
                                    fmt.Println(v6, v6, v6)
                                }
                                v69 := []int{(fi68 - fi66), (fi68 & fi66), 2147483647, fi68}
                                _ = fi68
                                _ = v69
                            }
                        }
                        var v70 uint8 = 1
                        _ = v70
                    } else {
                        if (!false) {
                            var v71 int8 = (v2 | 42)
                            v2 = 81
                            v72 := GenS2{F0: uint8(v6), F1: (!false), F2: v1}
                            _ = v71
                            _ = v72
                        } else {
                            if true {
                                for fi73 := 0; fi73 < 3; fi73++ {
                                    var v74 uint16 = (v4 ^ (v3 ^ v3))
                                    _ = fi73
                                    _ = v74
                                }
                            }
                            if false {
                                v3 -= v4
                                fmt.Println(v6, v3)
                                fmt.Println(fi66, v6, v6)
                            }
                        }
                        v75 := []int{}
                        _ = v75
                    }
                    v1 = v1
                    _ = fi66
                    _ = v67
                }
                fmt.Println(v4, v2)
            } else {
                for fi76 := 0; fi76 < 11; fi76++ {
                    var v77 uint = 0
                    _ = fi76
                    _ = v77
                }
                for fi78 := 0; fi78 < 19; fi78++ {
                    v3 += v4
                    v1 = (("a" + "") + v1)
                    var v79 uint = 335
                    _ = fi78
                    _ = v79
                }
                fmt.Println(v6, v2)
            }
            if false {
                switch 955 {
                case -610:
                    fmt.Println(v2)
                }
            }
        } else {
            v2 += (v2 / ((v2) | 1))
            var v80 int = -870
            var v81 uint16 = (v4 >> ((42) & 63))
            _ = v80
            _ = v81
        }
    }
    var v82 string = ("foo bar" + ("a" + "a"))
    var v83 int64 = -750
    v5[(^(v6))] = v6
    for fi84 := 0; fi84 < 1; fi84++ {
        var v85 int8 = 42
        v86 := []int{(fi84 & (fi84 % ((fi84) | 1))), (^(fi84)), (-(fi84))}
        v2 -= (v85 | v85)
        _ = fi84
        _ = v85
        _ = v86
    }
    v87 := []int{2147483647}
    v88 := GenS3{F0: v82, F1: genFn0(1, 840), F2: 2}
    v3 = 1
    var v89 uint16 = 278
    if true {
        var v90 bool = (((!(!false)) || (3.14159 <= 1.5)) || (-172 > 0))
        v91 := GenS3{F0: v82, F1: 1, F2: 45}
        for fi92 := 0; fi92 < 2; fi92++ {
            v5[(v6 & (^(v6)))] = 1
            var v93 int16 = 592
            if (2147483647 == v6) {
                var v94 int64 = v83
                _ = v94
            } else {
                fmt.Println(v93)
                var v95 int64 = -9223372036854775808
                _ = v95
            }
            _ = fi92
            _ = v93
        }
        _ = v90
        _ = v91
    } else {
        v82 = "foo bar"
        _ = len(v5)
        v82 = v1
    }
    v83 += v83
    var v96 int32 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v82:", v82)
    fmt.Println("v83:", v83)
    fmt.Println("slicelen v87:", len(v87), v87)
    fmt.Println("v89:", v89)
    fmt.Println("v96:", v96)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v82
    _ = v83
    _ = v87
    _ = v88
    _ = v89
    _ = v96
}
