package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16; F2 string }
type GenS2 struct { F0 int8; F1 uint8; F2 uint64 }
type GenS3 struct { F0 int; F1 uint }
type GenS4 struct { F0 uint32; F1 int16; F2 uint64 }
type GenS5 struct { F0 uint; F1 uint32; F2 bool }
type GenS6 struct { F0 uint8; F1 uint64; F2 uint }
type GenS7 struct { F0 bool; F1 int8; F2 uint32 }
func genFn0(p0 bool) int16 {
    var v5 uint16 = 32767
    v6 := []int{-847}
    switch 872 {
    case -832:
        fmt.Println(v5)
    case 9223372036854775807:
        if len(v6) > 0 {
            v6[0] = -1
        }
    case -1:
        for fi7 := 0; fi7 < 10; fi7++ {
            var v8 float32 = float32(3.14159)
            v8 -= (float32(((float64(v8) * ((1.5 + 1e10) - 3.14159)) / -2.25)) * 0.1)
            _ = fi7
            _ = v8
        }
    }
    _ = p0
    _ = v5
    _ = v6
    return 137
}
func genFn1(p0 int16, p1 bool, p2 int32) int {
    v9 := GenS1{F0: (-(p2)), F1: 1, F2: "foo bar"}
    var v10 string = "foo bar"
    v11 := []int{-1}
    v11 = append(v11, -833)
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v9
    return 128
}

func main() {
    v1 := make(map[int]uint32)
    var v2 uint = 542
    var v3 int16 = 32767
    var v4 int16 = (v3 << (((v2 + (^(v2)))) & 63))
    v4 = (v3 | int16(v2))
    v12 := GenS1{F0: -745, F1: 65535, F2: "a"}
    var v13 uint64 = 75
    var v14 uint64 = uint64(v2)
    var v15 uint16 = 255
    if false {
        v16 := v1[214]
        if false {
            if true {
                var v17 int32 = 2147483647
                _ = v17
            } else {
                var v18 uint64 = uint64(v16)
                var v19 uint16 = 255
                _ = v18
                _ = v19
            }
            v2 = 18446744073709551615
            switch 766 {
            case 32768:
                if (v4 >= v4) {
                    v20 := []int{}
                    _ = v20
                }
            default:
                v21 := []int{genFn1(1, ((v15 + 42) >= 42), 42), 1}
                _ = v21
            }
        }
        if false {
            if (!false) {
                for fi22 := 0; fi22 < 15; fi22++ {
                    fmt.Println(v2)
                    _ = len(v1)
                    _ = fi22
                }
            } else {
                switch -1 {
                case -9223372036854775808:
                    if (0.0 < (1.5 + (0.1 / 1e10))) {
                        v23 := GenS2{F0: -1, F1: 0, F2: v14}
                        if false {
                            v24 := []int{2147483647}
                            _ = v24
                        }
                        v2 = v2
                        _ = v23
                    }
                }
            }
            if (-2.25 <= (((((float32(0.1) + (1.5 * (1.5 - ((1.5 - 3.14159) + 0.0)))) * 0.1) * 3.14159) / ((-2.25 / -2.25) - ((((0.1 + (((float32(1.5) / 1.0) * 0.0) - (1.5 - ((3.14159 - ((-2.25 / (((-2.25 / float32(0.1)) / -2.25) / 1.5)) + 3.14159)) - 0.1)))) / 1.5) - 0.0) * 1e10))) - 1.5)) {
                v25 := GenS1{F0: -610, F1: v15, F2: (("" + "foo bar") + "\n\t")}
                if (812 < 1) {
                    if (true || true) {
                        if (!false) {
                            if true {
                                v26 := []int{48, int(v15), -846}
                                _ = v26
                            } else {
                                if (true && (1 <= uint8(v14))) {
                                    var v27 uint32 = v16
                                    var v28 int64 = 570
                                    v29 := []int{-439}
                                    _ = v27
                                    _ = v28
                                    _ = v29
                                } else {
                                    v30 := []int{-2147483648}
                                    _ = v30
                                }
                                var v31 uint64 = v13
                                v2 = (v2 % ((uint(v15)) | 1))
                                _ = v31
                            }
                            var v32 int = 4294967295
                            var v33 int = v32
                            _ = v32
                            _ = v33
                        }
                        v34 := []int{32767, 127}
                        v15 = v15
                        _ = v34
                    } else {
                        for fi35 := 0; fi35 < 15; fi35++ {
                            if (!(!false)) {
                                v16 -= v16
                                fmt.Println(v16, v15)
                                for fi36 := 0; fi36 < 2; fi36++ {
                                    var v37 int8 = 1
                                    v38 := GenS2{F0: (-(v37)), F1: 2, F2: 1}
                                    fmt.Println(v4)
                                    if true {
                                        v14 += v14
                                        if true {
                                            var v39 uint64 = (v13 / ((v14) | 1))
                                            var v40 uint32 = (^(v16))
                                            v41 := []int{-333, (^(fi36)), fi36, (fi35 >> ((v2) & 63))}
                                            _ = v39
                                            _ = v40
                                            _ = v41
                                        }
                                        v42 := GenS1{F0: 551, F1: v15, F2: ("\n\t" + "hello")}
                                        _ = v42
                                    } else {
                                        switch ("a" + "hello") {
                                        case "\n\t":
                                            v1[(fi35 % ((fi36) | 1))] = (v16 - v16)
                                        case "foo bar":
                                            v43 := GenS3{F0: (fi36 << ((uint(v15)) & 63)), F1: (v2 - v2)}
                                            _ = v43
                                        case "hello":
                                            fmt.Println(v15, v4, v4)
                                        default:
                                            var v44 bool = false
                                            var v45 float64 = 0.1
                                            v46 := GenS1{F0: 57, F1: (v15 % ((140) | 1)), F2: "a"}
                                            _ = v44
                                            _ = v45
                                            _ = v46
                                        }
                                        var v47 int32 = 42
                                        v48 := []int{}
                                        _ = v47
                                        _ = v48
                                    }
                                    _ = fi36
                                    _ = v37
                                    _ = v38
                                }
                            }
                            delete(v1, int(v13))
                            switch "hello" {
                            case "a":
                                var v49 int64 = -9223372036854775808
                                _ = v49
                            case "foo bar":
                                v14 = v14
                            default:
                                v3 = v3
                            }
                            _ = fi35
                        }
                    }
                    v50 := GenS4{F0: 65535, F1: (v3 * (v3 + -32768)), F2: 18446744073709551615}
                    _ = v50
                }
                _ = v25
            } else {
                v51 := v1[70]
                for fi52 := 0; fi52 < 12; fi52++ {
                    var v53 string = ("hello" + ("foo bar" + ""))
                    v54 := GenS5{F0: 695, F1: 2147483647, F2: ((^(v3)) != v4)}
                    for fi55 := 0; fi55 < 5; fi55++ {
                        var v56 uint16 = (v15 >> (((v2 | (v2 << (((v2 & v2)) & 63)))) & 63))
                        _ = fi55
                        _ = v56
                    }
                    _ = fi52
                    _ = v53
                    _ = v54
                }
                v57 := GenS6{F0: 127, F1: v14, F2: (^(v2))}
                _ = v51
                _ = v57
            }
        } else {
            for fi58 := 0; fi58 < 18; fi58++ {
                v59 := GenS7{F0: (((false || false) || (!(int16(v15) >= -195))) && false), F1: 127, F2: v16}
                for fi60 := 0; fi60 < 12; fi60++ {
                    v61 := []int{(fi60 >> ((v2) & 63)), (fi60 & 127)}
                    _ = fi60
                    _ = v61
                }
                if true {
                    var v62 uint32 = uint32(v15)
                    _ = v62
                }
                _ = fi58
                _ = v59
            }
            fmt.Println(v16)
        }
        _ = v16
    }
    var v63 int64 = -137
    if (false || true) {
        v63 = (v63 | (v63 % ((40) | 1)))
        v64 := GenS7{F0: (!false), F1: 30, F2: 0}
        _ = v64
    }
    fmt.Println(v4, v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v63:", v63)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v63
}
