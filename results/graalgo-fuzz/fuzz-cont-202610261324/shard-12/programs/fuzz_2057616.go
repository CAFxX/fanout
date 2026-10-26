package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool; F2 bool }
type GenS2 struct { F0 int; F1 uint16 }
type GenS3 struct { F0 int16; F1 int; F2 uint16 }
type GenS4 struct { F0 string; F1 int32 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 string; F1 int64 }
func genFn0(p0 int8) int8 {
    if (1 >= 42) {
        for fi5 := 0; fi5 < 18; fi5++ {
            v6 := []int{fi5}
            switch "\n\t" {
            case "a":
                var v7 int8 = (p0 ^ p0)
                _ = v7
            default:
                for fi8 := 0; fi8 < 7; fi8++ {
                    if (true && (775 < 0)) {
                        var v9 float64 = 3.14159
                        var v10 int16 = 42
                        _ = v10
                        _ = v9
                    }
                    p0 = (p0 * -124)
                    p0 -= (p0 ^ (p0 * 127))
                    _ = fi8
                }
            }
            for fi11 := 0; fi11 < 1; fi11++ {
                v12 := GenS1{F0: 2, F1: false, F2: false}
                if len(v6) > 0 {
                    v6[0] = (fi11 % ((9223372036854775807) | 1))
                }
                _ = fi11
                _ = v12
            }
            _ = fi5
            _ = v6
        }
    }
    for fi13 := 0; fi13 < 16; fi13++ {
        switch "foo bar" {
        case "\n\t":
            if false {
                var v14 int8 = 100
                if false {
                    var v15 uint = 255
                    var v16 int16 = -1
                    _ = v15
                    _ = v16
                } else {
                    if ((0 > 0) || true) {
                        v17 := []int{127, fi13}
                        _ = v17
                    }
                }
                p0 = (v14 + p0)
                _ = v14
            } else {
                var v18 string = ""
                v19 := GenS1{F0: 42, F1: (!("" < (v18 + ("a" + "foo bar")))), F2: (!false)}
                _ = v18
                _ = v19
            }
        case "hello":
            switch (^(fi13)) {
            case 1:
                switch fi13 {
                case 4294967295:
                    if (255 < 255) {
                        if (0 <= 581) {
                            v20 := GenS2{F0: fi13, F1: 2}
                            if true {
                                if (!false) {
                                    var v21 uint8 = 128
                                    var v22 int = fi13
                                    var v23 uint16 = 65535
                                    _ = v21
                                    _ = v22
                                    _ = v23
                                }
                                fmt.Println(fi13, fi13)
                            } else {
                                switch ("foo bar" + ("a" + (("" + ("" + "\n\t")) + "a"))) {
                                case "a":
                                    if (65535 != 528) {
                                        v24 := []int{-888, (fi13 << ((4294967295) & 63)), (fi13 % (((fi13 ^ -459)) | 1)), fi13}
                                        _ = v24
                                    } else {
                                        if (!false) {
                                            p0 -= p0
                                        }
                                        for fi25 := 0; fi25 < 3; fi25++ {
                                            p0 -= p0
                                            _ = fi25
                                        }
                                    }
                                case "":
                                    switch fi13 {
                                    case 128:
                                        for fi26 := 0; fi26 < 12; fi26++ {
                                            v27 := GenS1{F0: 0, F1: (42 <= 42), F2: (-326 == 981)}
                                            _ = fi26
                                            _ = v27
                                        }
                                    default:
                                        if false {
                                            p0 -= (^(p0))
                                            p0 = (p0 | p0)
                                        }
                                    }
                                case "\n\t":
                                    if true {
                                        if true {
                                            switch "" {
                                            case "hello":
                                                var v28 int16 = 0
                                                _ = v28
                                            }
                                            p0 += p0
                                            if (true && false) {
                                                var v29 uint64 = 42
                                                v29 += (v29 % (((v29 / ((v29) | 1))) | 1))
                                                _ = v29
                                            } else {
                                                for fi30 := 0; fi30 < 16; fi30++ {
                                                    p0 = p0
                                                    _ = fi30
                                                }
                                            }
                                        }
                                        var v31 int16 = -1
                                        _ = v31
                                    }
                                default:
                                    v32 := GenS1{F0: 18446744073709551615, F1: false, F2: true}
                                    _ = v32
                                }
                                p0 = 100
                                p0 -= p0
                            }
                            v33 := []int{-323}
                            _ = v20
                            _ = v33
                        } else {
                            if (0 >= 1) {
                                v34 := GenS3{F0: 32767, F1: (fi13 | fi13), F2: 42}
                                var v35 uint = 0
                                _ = v34
                                _ = v35
                            }
                            switch 32768 {
                            case 2147483647:
                                v36 := []int{}
                                _ = v36
                            }
                        }
                        p0 -= p0
                    }
                case -265:
                    for fi37 := 0; fi37 < 16; fi37++ {
                        fmt.Println(fi37, p0)
                        _ = fi37
                    }
                case 46:
                    fmt.Println(fi13)
                }
            case 0:
                v38 := GenS4{F0: "", F1: -804}
                _ = v38
            default:
                var v39 uint8 = 2
                _ = v39
            }
        case "a":
            var v40 int16 = 1
            _ = v40
        }
        var v41 bool = true
        _ = fi13
        _ = v41
    }
    _ = p0
    return p0
}
func genFn1(p0 uint64) int {
    p0 = (p0 * (p0 % ((451) | 1)))
    v42 := []int{-863, -738, 128}
    p0 += 42
    _ = p0
    _ = v42
    return 1
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[string]int)
    var v3 bool = true
    var v4 float64 = 1e10
    for fi43 := 0; fi43 < 5; fi43++ {
        var v44 float64 = (0.0 - (float64(1.5) - (v4 / (((1.5 - v4) / float64((float32(-2.25) - 1e10))) * 0.1))))
        fmt.Println(v44)
        _ = fi43
        _ = v44
    }
    v3 = (1 == 509)
    switch 4294967295 {
    case -2147483648:
        for fi45 := 0; fi45 < 5; fi45++ {
            v4 += 0.1
            _ = fi45
        }
    case 128:
        v46 := GenS5{F0: -1}
        _ = v46
    case 183:
        v47 := []int{32768, 0, 592, 32767}
        _ = v47
    default:
        if ((9223372036854775807 <= -9223372036854775808) || false) {
            if v3 {
                if (-571 >= 32767) {
                    v4 -= (v4 + v4)
                    _ = len(v1)
                }
                switch 32767 {
                case 357:
                    var v48 uint = 255
                    _ = v48
                default:
                    var v49 float64 = v4
                    _ = v49
                }
            } else {
                v4 += 0.1
            }
        } else {
            for fi50 := 0; fi50 < 14; fi50++ {
                fmt.Println(fi50, fi50, v4)
                var v51 int16 = 88
                v52 := []int{2147483647, (-(fi50)), genFn1(42)}
                _ = fi50
                _ = v51
                _ = v52
            }
        }
    }
    var v53 int32 = 42
    var v54 bool = (42 == 0)
    if (-1 >= 42) {
        var v55 int64 = -1
        switch ("hello" + ("\n\t" + "\n\t")) {
        case "\n\t":
            v56 := GenS6{F0: 0}
            _ = v56
        case "foo bar":
            v57 := GenS2{F0: 32767, F1: 2}
            _ = v57
        }
        var v58 uint32 = 42
        _ = v55
        _ = v58
    }
    for fi59 := 0; fi59 < 20; fi59++ {
        for fi60 := 0; fi60 < 12; fi60++ {
            if (1 != 1) {
                if v54 {
                    if false {
                        v54 = (3.14159 == float64(0.1))
                        v3 = v54
                        var v61 int64 = -1000
                        _ = v61
                    }
                    if (false && true) {
                        if v54 {
                            if v54 {
                                var v62 int64 = 543
                                _ = v62
                            } else {
                                if (!false) {
                                    v54 = (2147483647 >= 4294967295)
                                    var v63 int16 = -271
                                    _ = v63
                                } else {
                                    v53 = (v53 % ((v53) | 1))
                                    fmt.Println(v4)
                                }
                                var v64 int16 = 42
                                v65 := GenS7{F0: ("\n\t" + ("a" + "hello")), F1: 1}
                                _ = v64
                                _ = v65
                            }
                            v4 -= v4
                        } else {
                            if (((v3 && (false && ((0 < 128) && v3))) || v54) && ((0.1 / 1.5) <= 1.5)) {
                                v3 = v3
                            }
                        }
                        var v66 int16 = 1
                        _ = v66
                    }
                } else {
                    v67 := v1[(fi59 - genFn1(1))]
                    var v68 float32 = 0.0
                    _ = v67
                    _ = v68
                }
                v2[""] = fi60
            }
            if ((315 >= 4294967295) && (v54 || (!((true && true) && (!("\n\t" < "")))))) {
                switch fi60 {
                case -842:
                    if v3 {
                        var v69 uint = 2
                        _ = v69
                    }
                case 32767:
                    v3 = (699 == 493)
                case 32768:
                    v70 := GenS6{F0: 870}
                    _ = v70
                }
                var v71 float32 = float32(v4)
                switch (((("hello" + "") + "hello") + "a") + "hello") {
                case "foo bar":
                    var v72 int8 = 100
                    _ = v72
                case "":
                    v71 -= v71
                    fmt.Println(v3)
                    v3 = (!v54)
                case "a":
                    var v73 int32 = v53
                    _ = v73
                }
                fmt.Println(fi60, fi60, v4)
                _ = v71
            } else {
                switch genFn1(4294967295) {
                case 308:
                    var v74 string = "hello"
                    _ = v74
                case -299:
                    switch genFn1(0) {
                    case -1:
                        v53 += (v53 >> ((795) & 63))
                    }
                default:
                    v75 := v2[("" + "\n\t")]
                    _ = v75
                }
                var v76 uint8 = 2
                var v77 uint16 = 422
                _ = v76
                _ = v77
            }
            _ = fi60
        }
        for fi78 := 0; fi78 < 5; fi78++ {
            fmt.Println(v4, fi59)
            var v79 int64 = -1
            if ((false || true) || ((v53 - (^(v53))) == v53)) {
                var v80 uint = 4294967295
                var v81 uint32 = 430
                v81 = v81
                _ = v80
                _ = v81
            } else {
                v53 = (v53 + -348)
            }
            var v82 int32 = (v53 % ((2147483647) | 1))
            v3 = (!((((546 == 2) && (42 <= 744)) && (!v3)) && ("foo bar" < "foo bar")))
            var v83 uint16 = 2
            _ = fi78
            _ = v79
            _ = v82
            _ = v83
        }
        _ = fi59
    }
    v84 := []int{-880}
    v53 = (v53 + (v53 | (v53 - v53)))
    fmt.Println(v54)
    var v85 uint16 = 255
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v53:", v53)
    fmt.Println("v54:", v54)
    fmt.Println("slicelen v84:", len(v84), v84)
    fmt.Println("v85:", v85)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v53
    _ = v54
    _ = v84
    _ = v85
}
