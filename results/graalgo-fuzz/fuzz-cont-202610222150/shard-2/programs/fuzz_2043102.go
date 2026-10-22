package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int64; F1 string; F2 uint8 }
type GenS3 struct { F0 uint; F1 int64; F2 int64 }
type GenS4 struct { F0 uint16; F1 string }
func genFn0() uint {
    v5 := GenS1{F0: 127}
    var v6 float32 = 1.5
    var v7 uint16 = 0
    if (!false) {
        if (-223 != 57) {
            v7 = (v7 - 1)
            v7 -= (v7 % (((v7 & (v7 >> ((255) & 63)))) | 1))
            for fi8 := 0; fi8 < 11; fi8++ {
                v9 := []int{}
                _ = fi8
                _ = v9
            }
        }
        for fi10 := 0; fi10 < 3; fi10++ {
            v11 := GenS2{F0: 1, F1: ("foo bar" + "hello"), F2: 42}
            var v12 int = (-(fi10))
            _ = fi10
            _ = v11
            _ = v12
        }
    } else {
        fmt.Println(v6)
        if true {
            v13 := []int{-734, 9223372036854775807, 375}
            if true {
                for fi14 := 0; fi14 < 6; fi14++ {
                    v7 -= 65535
                    _ = fi14
                }
            } else {
                if ((239 < 326) || true) {
                    v6 = 0.1
                    switch ("" + "hello") {
                    case "\n\t":
                        if false {
                            var v15 int64 = -681
                            var v16 bool = (1 > v15)
                            _ = v15
                            _ = v16
                        }
                    case "a":
                        v17 := GenS2{F0: 42, F1: ("a" + "\n\t"), F2: 1}
                        _ = v17
                    case "hello":
                        var v18 int8 = 100
                        _ = v18
                    default:
                        var v19 uint64 = 4294967295
                        _ = v19
                    }
                    v20 := GenS1{F0: 1}
                    _ = v20
                } else {
                    switch 42 {
                    case 42:
                        if true {
                            if (true || false) {
                                v13 = append(v13, 32768)
                                v7 -= 255
                                var v21 uint16 = (v7 << ((333) & 63))
                                _ = v21
                            } else {
                                var v22 uint64 = 2
                                _ = len(v13)
                                v7 = (v7 & (^(v7)))
                                _ = v22
                            }
                        }
                    case 127:
                        for fi23 := 0; fi23 < 8; fi23++ {
                            v6 = v6
                            _ = fi23
                        }
                    }
                    var v24 int64 = -9223372036854775808
                    if (((v6 * v6) * v6) == 0.0) {
                        if true {
                            v6 += 1e10
                            switch -2147483648 {
                            case 127:
                                v25 := GenS1{F0: 0}
                                _ = v25
                            case 128:
                                for fi26 := 0; fi26 < 1; fi26++ {
                                    if (false || (!false)) {
                                        _ = len(v13)
                                        if false {
                                            if (!((v6 / ((0.1 / (1.5 / (((0.1 - (v6 * (3.14159 / 1.5))) * v6) + ((0.0 / ((1e10 + v6) + float32(float64(float32((0.1 / (-2.25 - 1e10))))))) + 1.5)))) + v6)) < v6)) {
                                                var v27 string = "foo bar"
                                                _ = v27
                                            } else {
                                                v24 += (v24 - (-(v24)))
                                                var v28 int32 = 752
                                                _ = v28
                                            }
                                        }
                                    } else {
                                        v29 := GenS2{F0: (v24 | v24), F1: ("\n\t" + (("a" + ("\n\t" + "a")) + "foo bar")), F2: 247}
                                        v7 = 0
                                        for fi30 := 0; fi30 < 17; fi30++ {
                                            switch (fi30 - (-(fi26))) {
                                            case 4294967295:
                                                v31 := GenS3{F0: 577, F1: v24, F2: v24}
                                                _ = v31
                                            case -1:
                                                switch -695 {
                                                case -493:
                                                    if true {
                                                        switch "foo bar" {
                                                        case "hello":
                                                            var v32 uint16 = (v7 | (v7 + v7))
                                                            _ = v32
                                                        }
                                                    } else {
                                                        var v33 string = "a"
                                                        switch fi30 {
                                                        case -2147483648:
                                                            if false {
                                                                var v34 float64 = 0.0
                                                                _ = v34
                                                            }
                                                        case 42:
                                                            var v35 uint32 = 4294967295
                                                            _ = v35
                                                        default:
                                                            var v36 bool = false
                                                            _ = v36
                                                        }
                                                        v6 = (3.14159 + v6)
                                                        _ = v33
                                                    }
                                                case -118:
                                                    fmt.Println(v6, fi30)
                                                }
                                            default:
                                                v24 -= v24
                                            }
                                            _ = fi30
                                        }
                                        _ = v29
                                    }
                                    _ = fi26
                                }
                            case -244:
                                fmt.Println(v24)
                            }
                        }
                    } else {
                        if len(v13) >= 2 {
                            v13 = v13[1:len(v13)]
                        }
                        v6 = -2.25
                    }
                    _ = v24
                }
                if true {
                    v7 -= (^(v7))
                    v13 = append(v13, 668)
                    if len(v13) >= 2 {
                        v13 = v13[1:len(v13)]
                    }
                }
                var v37 uint = 18446744073709551615
                _ = v37
            }
            _ = v13
        }
    }
    _ = v5
    _ = v6
    _ = v7
    return 65535
}
func genFn1(p0 uint, p1 int32) uint16 {
    for fi38 := 0; fi38 < 1; fi38++ {
        p1 = 42
        var v39 int32 = 2147483647
        for fi40 := 0; fi40 < 9; fi40++ {
            v39 = v39
            _ = fi40
        }
        _ = fi38
        _ = v39
    }
    fmt.Println(p1, p0)
    _ = p0
    _ = p1
    return 42
}

func main() {
    var v1 uint16 = 32767
    v2 := make(map[string]string)
    v3 := make(map[int]bool)
    var v4 float64 = 3.14159
    for fi41 := 0; fi41 < 19; fi41++ {
        v4 = v4
        v1 += v1
        _ = fi41
    }
    v1 = (v1 + (v1 - v1))
    if true {
        var v42 uint64 = 238
        v43 := GenS4{F0: 42, F1: "\n\t"}
        fmt.Println(v4, v42)
        _ = v42
        _ = v43
    }
    var v44 bool = false
    for fi45 := 0; fi45 < 4; fi45++ {
        v46 := GenS4{F0: v1, F1: "a"}
        var v47 uint = 8
        _ = fi45
        _ = v46
        _ = v47
    }
    switch -1 {
    case 32768:
        v4 += (v4 + 1e10)
    case 1:
        var v48 int32 = -604
        _ = v48
    }
    v4 = (-2.25 + float64(-2.25))
    v4 = (float64(1e10) / 1.0)
    v3[42] = v44
    v44 = (18446744073709551615 == 42)
    var v49 uint = 0
    v49 += v49
    if true {
        if true {
            var v50 uint = 2
            v51 := v2[("\n\t" + "a")]
            _ = v50
            _ = v51
        } else {
            var v52 float64 = 0.0
            for fi53 := 0; fi53 < 7; fi53++ {
                if (!((-2.25 - 3.14159) < (0.1 + 0.1))) {
                    v44 = v44
                    for fi54 := 0; fi54 < 2; fi54++ {
                        var v55 int32 = 1
                        _ = fi54
                        _ = v55
                    }
                    v56 := []int{1, fi53, (fi53 - (fi53 & -533)), (fi53 & (^(fi53)))}
                    _ = v56
                }
                for fi57 := 0; fi57 < 4; fi57++ {
                    v44 = ((0 <= -128) && (!((1 < 455) && false)))
                    if (v44 && v44) {
                        delete(v2, "foo bar")
                        switch "foo bar" {
                        case "":
                            v52 += v52
                        case "foo bar":
                            v3[fi53] = (-330 > fi53)
                        default:
                            v44 = (!false)
                        }
                        v58 := []int{}
                        v59 := GenS2{F0: -18, F1: "hello", F2: 176}
                        fmt.Println(v52, v1)
                        _ = v58
                        _ = v59
                    } else {
                        _ = len(v3)
                        switch fi57 {
                        case 42:
                            if (0.1 != ((-2.25 + (0.0 + -2.25)) - 0.1)) {
                                v2[""] = "\n\t"
                            }
                        default:
                            v52 = (v52 - v4)
                        }
                    }
                    _ = fi57
                }
                _ = fi53
            }
            v49 += (v49 & (^(v49)))
            _ = v52
        }
    } else {
        v60 := GenS3{F0: (v49 ^ (v49 - v49)), F1: -1, F2: 42}
        for fi61 := 0; fi61 < 12; fi61++ {
            if (127 == 127) {
                for fi62 := 0; fi62 < 18; fi62++ {
                    var v63 uint32 = 1
                    _ = fi62
                    _ = v63
                }
                var v64 uint32 = 4294967295
                _ = v64
            } else {
                var v65 int64 = 0
                _ = v65
            }
            _ = fi61
        }
        _ = v60
    }
    var v66 string = "\n\t"
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v49:", v49)
    fmt.Println("v66:", v66)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v49
    _ = v66
}
