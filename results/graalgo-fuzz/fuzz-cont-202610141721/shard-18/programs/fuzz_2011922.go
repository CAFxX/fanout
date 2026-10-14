package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint; F1 int32 }
type GenS5 struct { F0 int16; F1 bool }
type GenS6 struct { F0 uint8 }
func genFn0(p0 uint32, p1 int16) uint {
    if (false || true) {
        v6 := GenS1{F0: 32767, F1: 295}
        _ = v6
    }
    var v7 int = -9223372036854775808
    v8 := GenS1{F0: v7, F1: -936}
    v9 := GenS2{F0: -2147483648}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return 255
}
func genFn1(p0 int64) int16 {
    p0 -= -1
    p0 = (p0 - (p0 >> ((1) & 63)))
    var v10 int32 = -1
    var v11 int = 2147483647
    _ = p0
    _ = v10
    _ = v11
    return 42
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[int]string)
    var v3 float32 = 1e10
    v4 := make(map[int]uint32)
    v5 := []int{-758, 979, 128, -2147483648}
    if ("foo bar" != "a") {
        v12 := GenS2{F0: 128}
        _ = len(v4)
        if (true || (true || (true && false))) {
            var v13 int8 = 100
            if (42 > 2) {
                switch -9223372036854775808 {
                case 9223372036854775807:
                    v13 = 35
                default:
                    v3 = v3
                }
                v14 := GenS2{F0: 2147483647}
                _ = v14
            } else {
                fmt.Println(v3, v13)
                for fi15 := 0; fi15 < 13; fi15++ {
                    v3 -= (float32(0.1) - (v3 / float32(1.5)))
                    v13 = v13
                    v13 -= (v13 >> ((255) & 63))
                    _ = fi15
                }
            }
            _ = v13
        } else {
            for fi16 := 0; fi16 < 13; fi16++ {
                v3 = v3
                var v17 uint64 = 524
                for fi18 := 0; fi18 < 8; fi18++ {
                    v3 += v3
                    if len(v5) > 0 {
                        v5[0] = fi16
                    }
                    _ = fi18
                }
                _ = fi16
                _ = v17
            }
            switch "\n\t" {
            case "":
                var v19 float64 = float64(3.14159)
                _ = v19
            default:
                if (42 < 568) {
                    v20 := GenS3{F0: 1}
                    _ = len(v4)
                    switch 9223372036854775807 {
                    case -1:
                        _ = len(v5)
                    case -116:
                        var v21 bool = (42 < 578)
                        _ = v21
                    case 32768:
                        v3 += v3
                    default:
                        for fi22 := 0; fi22 < 1; fi22++ {
                            var v23 float32 = v3
                            if ((42 == 1) && true) {
                                _ = len(v2)
                                v24 := GenS4{F0: 42, F1: 1}
                                _ = v24
                            } else {
                                switch (-(fi22)) {
                                case 1:
                                    if false {
                                        v23 = v3
                                    } else {
                                        if true {
                                            var v25 int8 = 0
                                            v26 := GenS4{F0: 65535, F1: -2147483648}
                                            _ = v25
                                            _ = v26
                                        } else {
                                            for fi27 := 0; fi27 < 20; fi27++ {
                                                var v28 float64 = float64((-2.25 + float32((1e10 / 1.5))))
                                                var v29 int64 = -304
                                                if true {
                                                    if (18446744073709551615 == 940) {
                                                        v30 := GenS4{F0: 1, F1: -2147483648}
                                                        var v31 uint64 = 18446744073709551615
                                                        _ = v30
                                                        _ = v31
                                                    }
                                                    var v32 uint32 = 17
                                                    var v33 uint16 = 2
                                                    var v34 uint32 = v32
                                                    _ = v32
                                                    _ = v33
                                                    _ = v34
                                                }
                                                _ = fi27
                                                _ = v28
                                                _ = v29
                                            }
                                        }
                                    }
                                case 32768:
                                    for fi35 := 0; fi35 < 1; fi35++ {
                                        var v36 int16 = 1000
                                        v36 -= v36
                                        v23 = v3
                                        _ = fi35
                                        _ = v36
                                    }
                                case 32767:
                                    var v37 uint32 = 2
                                    _ = v37
                                }
                                for fi38 := 0; fi38 < 11; fi38++ {
                                    v39 := GenS3{F0: 2}
                                    _ = fi38
                                    _ = v39
                                }
                            }
                            var v40 uint8 = 128
                            _ = fi22
                            _ = v23
                            _ = v40
                        }
                    }
                    _ = v20
                }
            }
            var v41 int64 = 208
            _ = v41
        }
        _ = v12
    } else {
        for fi42 := 0; fi42 < 8; fi42++ {
            v43 := GenS5{F0: 42, F1: (2 <= 42)}
            _ = fi42
            _ = v43
        }
        if true {
            if true {
                fmt.Println(v3)
                var v44 uint32 = 1
                _ = v44
            }
            if true {
                v3 += ((((-2.25 - (((-2.25 - (float32(1e10) - float32(3.14159))) - 0.1) * v3)) + v3) / ((-2.25 / v3) / 3.14159)) - -2.25)
                switch -620 {
                case -909:
                    var v45 int64 = -1
                    _ = v45
                case -106:
                    fmt.Println(v3)
                }
                v3 -= v3
            }
            var v46 string = ("" + "")
            _ = v46
        }
    }
    var v47 float64 = -2.25
    switch -9223372036854775808 {
    case 0:
        var v48 int = 364
        _ = v48
    default:
        fmt.Println(v3)
    }
    switch "foo bar" {
    case "foo bar":
        for fi49 := 0; fi49 < 8; fi49++ {
            var v50 int16 = 640
            switch (fi49 / ((fi49) | 1)) {
            case 537:
                var v51 uint8 = 128
                _ = v51
            }
            var v52 uint8 = 122
            _ = fi49
            _ = v50
            _ = v52
        }
    case "hello":
        v5 = append(v5, 9223372036854775807)
    default:
        v53 := GenS1{F0: 9223372036854775807, F1: -589}
        _ = v53
    }
    fmt.Println(v3, v47)
    fmt.Println(v47, v3)
    var v54 int32 = -324
    _ = len(v2)
    v55 := GenS1{F0: -2147483648, F1: 0}
    delete(v4, 0)
    if false {
        if (v3 < 1e10) {
            var v56 uint32 = 42
            for fi57 := 0; fi57 < 2; fi57++ {
                var v58 uint = 65535
                var v59 float64 = 0.0
                if ((9223372036854775807 == -9223372036854775808) || false) {
                    fmt.Println(v3)
                    switch "\n\t" {
                    case "foo bar":
                        if false {
                            v60 := GenS6{F0: 0}
                            _ = v60
                        } else {
                            for fi61 := 0; fi61 < 2; fi61++ {
                                v54 -= v54
                                fmt.Println(v58)
                                v54 -= (^(v54))
                                _ = fi61
                            }
                        }
                    default:
                        v54 = v54
                    }
                }
                _ = fi57
                _ = v58
                _ = v59
            }
            _ = v56
        } else {
            v54 = (v54 | (v54 % ((v54) | 1)))
        }
    } else {
        for fi62 := 0; fi62 < 9; fi62++ {
            var v63 uint64 = 2
            var v64 int64 = -9223372036854775808
            if len(v5) > 0 {
                v5[0] = -577
            }
            _ = fi62
            _ = v63
            _ = v64
        }
        v1[4] = 65535
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v47:", v47)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v54:", v54)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v5
    _ = v54
    _ = v55
}
