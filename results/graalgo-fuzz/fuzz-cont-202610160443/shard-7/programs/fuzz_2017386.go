package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16; F2 uint32 }
type GenS2 struct { F0 uint; F1 uint }
type GenS3 struct { F0 int16; F1 int8; F2 string }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 uint32; F1 int64; F2 uint8 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 string }
type GenS9 struct { F0 uint16; F1 int32 }
func genFn0(p0 uint32, p1 int64) int32 {
    var v5 int8 = -1
    for fi6 := 0; fi6 < 4; fi6++ {
        v7 := GenS1{F0: -13, F1: 1, F2: 4294967295}
        v5 = 100
        fmt.Println(p0, p0)
        _ = fi6
        _ = v7
    }
    var v8 uint64 = 2
    var v9 float64 = -2.25
    _ = p0
    _ = p1
    _ = v5
    _ = v8
    _ = v9
    return -2147483648
}

func main() {
    v1 := []int{}
    v2 := []string{""}
    var v3 uint32 = 683
    v4 := make(map[int]string)
    if true {
        for fi10 := 0; fi10 < 19; fi10++ {
            switch fi10 {
            case -1:
                if false {
                    _ = len(v4)
                    if (1.5 >= ((1.5 - ((0.0 / 3.14159) - -2.25)) * (-2.25 + (1e10 / 1.0)))) {
                        v11 := GenS1{F0: 0, F1: 733, F2: 497}
                        switch (((("" + (("" + "a") + "\n\t")) + "\n\t") + "hello") + "hello") {
                        case "a":
                            if false {
                                var v12 float32 = 1e10
                                var v13 string = "\n\t"
                                _ = v12
                                _ = v13
                            } else {
                                if (1.5 >= -2.25) {
                                    var v14 int = fi10
                                    _ = v14
                                }
                                fmt.Println(fi10, fi10)
                            }
                        case "":
                            v15 := GenS1{F0: -1, F1: 32767, F2: v3}
                            _ = v15
                        case "foo bar":
                            v16 := GenS1{F0: 100, F1: 65535, F2: 638}
                            _ = v16
                        }
                        v3 -= (v3 | (v3 & v3))
                        _ = v11
                    } else {
                        var v17 uint32 = 2147483648
                        v18 := GenS1{F0: -91, F1: 604, F2: (v17 * v3)}
                        _ = v17
                        _ = v18
                    }
                }
            default:
                var v19 uint64 = 2
                _ = v19
            }
            if (0 == 9223372036854775807) {
                for fi20 := 0; fi20 < 3; fi20++ {
                    v3 = v3
                    fmt.Println(fi10)
                    v3 = (v3 & v3)
                    if true {
                        v3 -= (v3 + 4294967295)
                        fmt.Println(fi20)
                        switch ("foo bar" + "\n\t") {
                        case "hello":
                            var v21 int32 = -1
                            _ = v21
                        case "":
                            v3 += v3
                            var v22 float32 = 1e10
                            if ((-1 != 0) && true) {
                                v22 += float32((-2.25 + -2.25))
                            }
                            _ = v22
                        default:
                            v3 = v3
                        }
                        var v23 uint = 2
                        v3 = v3
                        v3 -= v3
                        _ = v23
                    } else {
                        _ = len(v1)
                    }
                    _ = fi20
                }
                v3 = (v3 * v3)
                for fi24 := 0; fi24 < 7; fi24++ {
                    v3 = (v3 & (v3 * 2))
                    v25 := GenS1{F0: 1, F1: 2, F2: (v3 * (v3 << ((255) & 63)))}
                    _ = fi24
                    _ = v25
                }
            } else {
                switch 2147483647 {
                case 369:
                    if (42 >= 1) {
                        var v26 int64 = -651
                        if true {
                            v27 := GenS1{F0: 11, F1: 117, F2: v3}
                            fmt.Println(v3, fi10)
                            _ = v27
                        } else {
                            var v28 bool = true
                            _ = v28
                        }
                        for fi29 := 0; fi29 < 19; fi29++ {
                            var v30 int16 = -811
                            v26 += v26
                            if true {
                                if true {
                                    _ = len(v2)
                                    switch "foo bar" {
                                    case "hello":
                                        v26 -= (v26 * -355)
                                    case "":
                                        v26 = int64(v3)
                                    default:
                                        var v31 int8 = 127
                                        _ = v31
                                    }
                                    v3 = (v3 << ((0) & 63))
                                } else {
                                    v32 := GenS2{F0: 255, F1: 830}
                                    var v33 uint32 = v3
                                    if (0 != (v26 + (v26 | 42))) {
                                        v26 = v26
                                    } else {
                                        v33 += (^(v33))
                                        v1 = append(v1, fi10)
                                    }
                                    _ = v32
                                    _ = v33
                                }
                                v3 += (v3 | (v3 << ((18446744073709551615) & 63)))
                                v26 = (-(v26))
                            }
                            _ = fi29
                            _ = v30
                        }
                        _ = v26
                    } else {
                        fmt.Println(v3, v3)
                        var v34 float64 = 3.14159
                        v34 += ((v34 * ((v34 - v34) - 1e10)) / v34)
                        _ = v34
                    }
                case 42:
                    v3 = (v3 - (^(v3)))
                }
                v35 := GenS2{F0: 481, F1: 133}
                _ = v35
            }
            _ = fi10
        }
    } else {
        v3 = (v3 & 0)
        v36 := GenS1{F0: 0, F1: 0, F2: (v3 / (((v3 ^ (v3 - v3))) | 1))}
        _ = v36
    }
    v2 = append(v2, ((("\n\t" + "") + "") + "a"))
    fmt.Println(v3)
    var v37 uint32 = v3
    v38 := GenS3{F0: 1000, F1: 127, F2: "a"}
    fmt.Println(v3, v3)
    v3 = 42
    v37 += v37
    if true {
        switch -1 {
        case 294:
            fmt.Println(v3, v37)
        }
    }
    if (18446744073709551615 >= 2) {
        v39 := GenS4{F0: 4294967295}
        for fi40 := 0; fi40 < 1; fi40++ {
            v41 := GenS5{F0: 42}
            v3 = 2
            for fi42 := 0; fi42 < 4; fi42++ {
                switch (fi42 << ((1) & 63)) {
                case 9223372036854775807:
                    var v43 uint64 = 67
                    _ = v43
                case 32768:
                    if false {
                        v44 := GenS2{F0: 18446744073709551615, F1: 496}
                        v45 := GenS5{F0: 2147483647}
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        if (((0.1 / 1.0) == 0.0) || true) {
                            var v46 string = "\n\t"
                            v37 = v3
                            fmt.Println(v46)
                            _ = v46
                        } else {
                            v3 = v37
                            fmt.Println(v3, v3, fi40)
                            v37 -= (^(v37))
                        }
                        if (v3 != v37) {
                            fmt.Println(fi40, v3)
                        }
                        delete(v4, 127)
                        var v47 uint = 0
                        var v48 uint64 = 18446744073709551615
                        _ = v44
                        _ = v45
                        _ = v47
                        _ = v48
                    }
                case -2147483648:
                    if true {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        var v49 uint = 63
                        if false {
                            if (false || false) {
                                if (961 == 42) {
                                    if ((!((-1 >= 582) && ((1e10 / (1e10 + 0.0)) < (1.5 + float64(1e10))))) && (3.14159 < 3.14159)) {
                                        switch "foo bar" {
                                        case "hello":
                                            if true {
                                                v50 := GenS6{F0: (v37 ^ (v37 & 65535)), F1: -9223372036854775808, F2: uint8(v49)}
                                                if false {
                                                    if (true || true) {
                                                        var v51 string = "a"
                                                        if len(v1) > 0 {
                                                            v1[0] = fi40
                                                        }
                                                        _ = v51
                                                    } else {
                                                        v3 = (^(v37))
                                                        var v52 uint = (v49 >> (((^(v49))) & 63))
                                                        _ = v52
                                                    }
                                                    v53 := GenS7{F0: 725}
                                                    var v54 float32 = 1e10
                                                    _ = v53
                                                    _ = v54
                                                }
                                                switch fi42 {
                                                case 41:
                                                    v49 = v49
                                                }
                                                switch "" {
                                                case "hello":
                                                    fmt.Println(v37)
                                                    var v55 string = ""
                                                    v3 -= (v37 ^ v3)
                                                    _ = v55
                                                case "\n\t":
                                                    var v56 uint16 = 978
                                                    _ = v56
                                                default:
                                                    fmt.Println(v3, v3)
                                                }
                                                _ = v50
                                            } else {
                                                fmt.Println(v37, v37, v3)
                                                v37 = (^(v37))
                                            }
                                        }
                                        var v57 int8 = -1
                                        _ = v57
                                    } else {
                                        v49 = v49
                                    }
                                    switch "a" {
                                    case "\n\t":
                                        fmt.Println(fi40, v49)
                                    case "":
                                        var v58 int32 = genFn0(v37, 9223372036854775807)
                                        _ = v58
                                    }
                                    _ = len(v2)
                                } else {
                                    v59 := v4[(fi42 * (fi40 * (-(fi40))))]
                                    v3 += uint32(v49)
                                    switch (fi42 % ((fi42) | 1)) {
                                    case 1:
                                        if true {
                                            var v60 int16 = 554
                                            _ = v60
                                        }
                                    }
                                    _ = v59
                                }
                                v2 = append(v2, "")
                            }
                        }
                        _ = v49
                    } else {
                        v3 = (v37 - v3)
                        var v61 float32 = 3.14159
                        _ = v61
                    }
                default:
                    var v62 uint8 = 2
                    v37 = 1
                    v63 := GenS8{F0: "foo bar"}
                    var v64 uint64 = 4294967295
                    fmt.Println(v64, v3, v3)
                    _ = v62
                    _ = v63
                    _ = v64
                }
                _ = fi42
            }
            _ = fi40
            _ = v41
        }
        _ = v39
    }
    if true {
        fmt.Println(v3, v3)
        v2 = append(v2, ("foo bar" + "a"))
        v65 := GenS9{F0: 2, F1: genFn0((v3 + 726), 1)}
        _ = v65
    } else {
        v66 := v4[4294967295]
        _ = v66
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v4
}
