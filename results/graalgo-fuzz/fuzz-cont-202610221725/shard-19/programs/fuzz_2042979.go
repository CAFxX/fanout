package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint; F2 int16 }
type GenS2 struct { F0 string; F1 uint64 }
type GenS3 struct { F0 int32; F1 uint32; F2 string }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int64; F1 bool }
type GenS6 struct { F0 int8; F1 int32 }
func genFn0(p0 int64, p1 uint8) uint {
    v6 := []int{}
    p1 -= p1
    for fi7 := 0; fi7 < 5; fi7++ {
        switch "a" {
        case "a":
            v8 := GenS1{F0: 1, F1: uint(p1), F2: -1}
            _ = v8
        case "\n\t":
            p1 = (^(p1))
        }
        v9 := GenS1{F0: 118, F1: 4294967295, F2: -1}
        var v10 string = ((("hello" + "foo bar") + "a") + "a")
        _ = fi7
        _ = v10
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v6
    return 830
}
func genFn1(p0 uint64, p1 uint16, p2 int) uint {
    v11 := GenS2{F0: (("hello" + "hello") + ("\n\t" + ("a" + "a"))), F1: 2}
    v12 := []int{(p2 * (p2 | (p2 / (((p2 | p2)) | 1))))}
    fmt.Println(p0, p1, p0)
    for fi13 := 0; fi13 < 17; fi13++ {
        if len(v12) >= 2 {
            v12 = v12[1:len(v12)]
        }
        _ = fi13
    }
    v14 := GenS3{F0: 1, F1: 2147483648, F2: ""}
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    _ = v14
    return genFn0(-1, uint8(p1))
}

func main() {
    v1 := []uint32{2147483648, 42, 2147483648}
    v2 := []int{-520, 2147483647, 300, 4294967295}
    v3 := make(map[uint32]bool)
    v4 := make(map[int]int)
    var v5 uint16 = 0
    var v15 int64 = 999
    if false {
        for fi16 := 0; fi16 < 10; fi16++ {
            fmt.Println(v5)
            _ = fi16
        }
    } else {
        v5 = 42
        switch "a" {
        case "foo bar":
            var v17 int = 0
            _ = v17
        case "":
            for fi18 := 0; fi18 < 4; fi18++ {
                if false {
                    fmt.Println(fi18, v5)
                    v19 := GenS2{F0: ("hello" + "hello"), F1: 105}
                    _ = v19
                }
                switch fi18 {
                case 32767:
                    var v20 float32 = ((-2.25 + 0.1) / 1.0)
                    _ = v20
                default:
                    v21 := GenS2{F0: "", F1: 0}
                    _ = v21
                }
                _ = fi18
            }
        case "\n\t":
            var v22 int16 = -374
            _ = v22
        }
    }
    switch 127 {
    case -2147483648:
        if true {
            for fi23 := 0; fi23 < 9; fi23++ {
                for fi24 := 0; fi24 < 9; fi24++ {
                    switch 32768 {
                    case 127:
                        v25 := GenS4{F0: (fi24 * fi24)}
                        _ = v25
                    case 662:
                        v5 = v5
                    default:
                        var v26 int = fi24
                        _ = v26
                    }
                    _ = fi24
                }
                for fi27 := 0; fi27 < 17; fi27++ {
                    v5 -= 42
                    v28 := GenS3{F0: 1, F1: 2147483647, F2: "foo bar"}
                    _ = fi27
                    _ = v28
                }
                fmt.Println(v15)
                _ = fi23
            }
        }
    }
    var v29 int32 = 0
    var v30 float64 = (3.14159 + 0.0)
    switch "hello" {
    case "hello":
        for fi31 := 0; fi31 < 9; fi31++ {
            for fi32 := 0; fi32 < 9; fi32++ {
                var v33 float32 = -2.25
                v34 := GenS3{F0: -505, F1: 2147483647, F2: (("foo bar" + "\n\t") + "hello")}
                v30 = (0.0 - (v30 * v30))
                _ = fi32
                _ = v33
                _ = v34
            }
            for fi35 := 0; fi35 < 7; fi35++ {
                if false {
                    if false {
                        v15 += -9223372036854775808
                        var v36 int = 250
                        _ = v36
                    } else {
                        if (255 == 128) {
                            if true {
                                fmt.Println(fi35, v30, v5)
                            } else {
                                var v37 int32 = v29
                                _ = v37
                            }
                            var v38 int32 = v29
                            v39 := GenS5{F0: 585, F1: (float32((float64(-2.25) / float64(1e10))) >= 0.1)}
                            _ = v38
                            _ = v39
                        } else {
                            _ = len(v4)
                            v5 = v5
                            var v40 uint = 724
                            v40 = genFn0(1, 127)
                            _ = v40
                        }
                        v2 = append(v2, fi35)
                        var v41 uint = 42
                        var v42 uint16 = (v5 << ((v41) & 63))
                        _ = v41
                        _ = v42
                    }
                    if true {
                        fmt.Println(v15, v29)
                        var v43 int16 = -802
                        fmt.Println(v43, v30, fi31)
                        _ = v43
                    } else {
                        switch -798 {
                        case -135:
                            v29 -= (v29 - 42)
                        case 127:
                            var v44 bool = false
                            _ = v44
                        case 1:
                            v45 := GenS3{F0: (v29 >> ((65535) & 63)), F1: 188, F2: ("hello" + "\n\t")}
                            _ = v45
                        }
                        if (-88 == 1) {
                            if (!true) {
                                v2 = append(v2, fi31)
                                var v46 int32 = 42
                                _ = v46
                            }
                            var v47 bool = true
                            fmt.Println(v30)
                            _ = v47
                        }
                        v15 = 0
                        v5 = v5
                    }
                    v48 := GenS1{F0: 4294967295, F1: 42, F2: 0}
                    _ = v48
                } else {
                    if true {
                        if false {
                            var v49 int = (^(fi31))
                            _ = v49
                        } else {
                            if len(v1) > 0 {
                                v1[0] = 65535
                            }
                            var v50 int8 = 42
                            v51 := GenS6{F0: v50, F1: v29}
                            _ = v50
                            _ = v51
                        }
                        v15 = (v15 - v15)
                        if (!true) {
                            fmt.Println(fi35)
                        } else {
                            if (v15 != v15) {
                                v5 = 389
                            } else {
                                var v52 uint8 = 128
                                var v53 int32 = v29
                                _ = v52
                                _ = v53
                            }
                            var v54 uint8 = 255
                            if false {
                                switch 32768 {
                                case 9223372036854775807:
                                    var v55 float64 = 1.5
                                    _ = v55
                                default:
                                    var v56 uint16 = (v5 - (v5 | (^(v5))))
                                    _ = v56
                                }
                                fmt.Println(v30)
                                if true {
                                    v15 += v15
                                    var v57 uint64 = 530
                                    v58 := GenS4{F0: fi35}
                                    _ = v57
                                    _ = v58
                                }
                            } else {
                                var v59 int16 = 677
                                v15 -= 42
                                _ = v59
                            }
                            _ = v54
                        }
                    } else {
                        if true {
                            var v60 bool = (!(!((!true) && false)))
                            if v60 {
                                v29 = v29
                                if len(v2) > 0 {
                                    v2[0] = fi31
                                }
                            }
                            var v61 int64 = v15
                            _ = v60
                            _ = v61
                        } else {
                            v62 := GenS5{F0: (v15 * (v15 << ((genFn0(v15, 234)) & 63))), F1: true}
                            _ = v62
                        }
                        var v63 uint8 = 1
                        _ = v63
                    }
                }
                v64 := GenS5{F0: 0, F1: (true || false)}
                var v65 int32 = (v29 * (-(v29)))
                _ = fi35
                _ = v64
                _ = v65
            }
            _ = fi31
        }
    }
    for fi66 := 0; fi66 < 4; fi66++ {
        if true {
            if (255 == 127) {
                for fi67 := 0; fi67 < 15; fi67++ {
                    var v68 int64 = (v15 / ((1) | 1))
                    v69 := GenS2{F0: "\n\t", F1: 18446744073709551615}
                    _ = fi67
                    _ = v68
                    _ = v69
                }
            } else {
                fmt.Println(fi66, v5)
                v70 := GenS4{F0: (fi66 | -749)}
                _ = v70
            }
            v71 := GenS1{F0: 0, F1: 2, F2: 1}
            v30 = v30
            _ = v71
        }
        for fi72 := 0; fi72 < 13; fi72++ {
            fmt.Println(v5, fi66)
            _ = fi72
        }
        _ = fi66
    }
    var v73 float64 = v30
    var v74 int16 = 32767
    fmt.Println(v5, v29)
    v5 += v5
    v29 = (v29 | v29)
    v73 += ((0.0 / 1e10) * v73)
    for fi75 := 0; fi75 < 15; fi75++ {
        v76 := GenS6{F0: 42, F1: (v29 * 806)}
        _ = fi75
        _ = v76
    }
    var v77 uint64 = 4294967295
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v73:", v73)
    fmt.Println("v74:", v74)
    fmt.Println("v77:", v77)
    _ = v1
    _ = v15
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v73
    _ = v74
    _ = v77
}
