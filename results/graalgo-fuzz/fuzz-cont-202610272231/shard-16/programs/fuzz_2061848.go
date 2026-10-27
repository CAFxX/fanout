package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64 }
type GenS2 struct { F0 string; F1 uint }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int8; F1 uint }
type GenS5 struct { F0 int32; F1 int8 }
func genFn0(p0 int, p1 int16) uint32 {
    v5 := []int{p0}
    p0 -= (p0 / ((128) | 1))
    for fi6 := 0; fi6 < 18; fi6++ {
        var v7 int16 = (p1 >> ((18446744073709551615) & 63))
        for fi8 := 0; fi8 < 18; fi8++ {
            var v9 uint32 = 42
            if false {
                v10 := GenS1{F0: -2147483648, F1: 4294967295}
                _ = v10
            }
            if ((p1 | (v7 & p1)) != (v7 * p1)) {
                v11 := GenS1{F0: -1, F1: 883}
                _ = v11
            }
            fmt.Println(p1, v7, fi6)
            _ = fi8
            _ = v9
        }
        _ = fi6
        _ = v7
    }
    fmt.Println(p0)
    for fi12 := 0; fi12 < 6; fi12++ {
        var v13 int8 = 1
        _ = fi12
        _ = v13
    }
    _ = p0
    _ = p1
    _ = v5
    return 432
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint32 = 2
    var v3 uint8 = 0
    v4 := make(map[string]int)
    var v14 float32 = (3.14159 + -2.25)
    if false {
        v15 := GenS1{F0: 42, F1: 2}
        if false {
            switch "hello" {
            case "\n\t":
                switch "\n\t" {
                case "foo bar":
                    switch "" {
                    case "foo bar":
                        if (42 > 2) {
                            if (!("foo bar" != "hello")) {
                                if false {
                                    v16 := GenS2{F0: "\n\t", F1: 414}
                                    var v17 string = ("a" + "")
                                    _ = v16
                                    _ = v17
                                }
                                switch "a" {
                                case "hello":
                                    if false {
                                        for fi18 := 0; fi18 < 16; fi18++ {
                                            v14 -= (v14 * 0.1)
                                            _ = fi18
                                        }
                                        for fi19 := 0; fi19 < 1; fi19++ {
                                            v20 := GenS1{F0: -2147483648, F1: 42}
                                            _ = fi19
                                            _ = v20
                                        }
                                        if false {
                                            for fi21 := 0; fi21 < 12; fi21++ {
                                                var v22 uint = 18446744073709551615
                                                fmt.Println(v14, v3)
                                                _ = fi21
                                                _ = v22
                                            }
                                            var v23 uint64 = 42
                                            _ = v23
                                        }
                                    }
                                default:
                                    for fi24 := 0; fi24 < 16; fi24++ {
                                        var v25 bool = (((((v14 / v14) + 0.0) + ((float32(1.5) * (-2.25 * v14)) / v14)) > float32(0.0)) || true)
                                        switch (-(fi24)) {
                                        case 478:
                                            var v26 uint64 = 1
                                            _ = v26
                                        case -586:
                                            for fi27 := 0; fi27 < 7; fi27++ {
                                                if ((v25 || (v25 || v25)) || (!(!((v25 && v25) || (v25 && ("hello" < "a")))))) {
                                                    _ = len(v1)
                                                    fmt.Println(v2, v2, v2)
                                                    v2 += (v2 >> ((2) & 63))
                                                } else {
                                                    var v28 int8 = 53
                                                    v4[(("hello" + ("" + "a")) + "hello")] = (fi24 | (fi27 & (fi27 >> ((18446744073709551615) & 63))))
                                                    _ = v28
                                                }
                                                v29 := GenS1{F0: 2147483647, F1: uint64(v2)}
                                                v3 += (v3 * v3)
                                                _ = fi27
                                                _ = v29
                                            }
                                        }
                                        _ = fi24
                                        _ = v25
                                    }
                                }
                                switch 4294967295 {
                                case -2147483648:
                                    fmt.Println(v14)
                                case -9223372036854775808:
                                    if true {
                                        v14 -= v14
                                        if (!false) {
                                            v30 := GenS2{F0: ("" + ("hello" + "hello")), F1: 255}
                                            var v31 int32 = 42
                                            v32 := []int{727, 294, 1}
                                            _ = v30
                                            _ = v31
                                            _ = v32
                                        }
                                        fmt.Println(v3, v3, v14)
                                    }
                                case -1:
                                    v1[2147483647] = false
                                }
                            }
                            var v33 uint16 = 65535
                            _ = v33
                        } else {
                            var v34 int64 = -502
                            v14 -= (((v14 / float32((1.5 / 1e10))) - v14) / float32(1e10))
                            v14 = 0.1
                            _ = v34
                        }
                    case "\n\t":
                        v35 := GenS2{F0: "foo bar", F1: 987}
                        _ = v35
                    case "":
                        var v36 uint64 = 985
                        _ = v36
                    }
                case "hello":
                    if (211 > (v3 - v3)) {
                        v2 = genFn0(9223372036854775807, 42)
                        for fi37 := 0; fi37 < 18; fi37++ {
                            fmt.Println(v14, fi37, v14)
                            _ = fi37
                        }
                        v38 := GenS2{F0: ("foo bar" + (("" + "foo bar") + "")), F1: 25}
                        _ = v38
                    }
                case "a":
                    v39 := GenS1{F0: 2147483647, F1: 1}
                    _ = v39
                default:
                    if (false || (65535 <= 388)) {
                        v40 := GenS3{F0: v3}
                        var v41 uint32 = v2
                        _ = v40
                        _ = v41
                    } else {
                        if (32767 <= -766) {
                            for fi42 := 0; fi42 < 4; fi42++ {
                                v1[2147483647] = (1 > 1)
                                var v43 uint16 = 139
                                var v44 int32 = -329
                                _ = fi42
                                _ = v43
                                _ = v44
                            }
                            var v45 uint32 = 1
                            _ = v45
                        } else {
                            if true {
                                var v46 uint64 = 1
                                var v47 uint = 18446744073709551615
                                _ = v46
                                _ = v47
                            }
                        }
                        switch 0 {
                        case 42:
                            v14 = v14
                        case -9223372036854775808:
                            if (!((((v14 + v14) * v14) * v14) <= ((float32((-2.25 + ((0.0 - 1.5) + (-2.25 - ((0.1 / 1.5) / ((0.1 + 3.14159) * 1.5)))))) + v14) * (1.5 - v14)))) {
                                for fi48 := 0; fi48 < 6; fi48++ {
                                    var v49 int16 = 0
                                    for fi50 := 0; fi50 < 7; fi50++ {
                                        var v51 float64 = 0.0
                                        _ = fi50
                                        _ = v51
                                    }
                                    if ((v2 + v2) < v2) {
                                        switch "foo bar" {
                                        case "foo bar":
                                            v52 := []int{fi48, fi48, fi48, fi48}
                                            _ = v52
                                        case "hello":
                                            v49 -= v49
                                        case "\n\t":
                                            v3 += (v3 + (^(v3)))
                                        default:
                                            if (2147483647 != -1) {
                                                var v53 int64 = 0
                                                v53 -= v53
                                                _ = v53
                                            }
                                        }
                                        if (v49 <= (^(v49))) {
                                            if false {
                                                var v54 int = (fi48 >> ((386) & 63))
                                                v55 := GenS3{F0: (^(v3))}
                                                _ = v54
                                                _ = v55
                                            } else {
                                                var v56 uint16 = 225
                                                for fi57 := 0; fi57 < 10; fi57++ {
                                                    v58 := GenS1{F0: -2147483648, F1: 4294967295}
                                                    _ = fi57
                                                    _ = v58
                                                }
                                                v49 -= (-(v49))
                                                _ = v56
                                            }
                                            if true {
                                                _ = len(v1)
                                                v49 += v49
                                                var v59 string = "hello"
                                                _ = v59
                                            }
                                            switch (fi48 | (^(fi48))) {
                                            case -626:
                                                for fi60 := 0; fi60 < 2; fi60++ {
                                                    var v61 uint = 489
                                                    fmt.Println(fi60, fi48)
                                                    _ = fi60
                                                    _ = v61
                                                }
                                            }
                                        } else {
                                            for fi62 := 0; fi62 < 19; fi62++ {
                                                v4[("hello" + ("hello" + "foo bar"))] = fi48
                                                v3 += v3
                                                var v63 float64 = 1.5
                                                _ = fi62
                                                _ = v63
                                            }
                                            for fi64 := 0; fi64 < 11; fi64++ {
                                                if false {
                                                    v14 -= 1.5
                                                    var v65 uint8 = 1
                                                    _ = v65
                                                } else {
                                                    v2 -= uint32(v3)
                                                    var v66 uint64 = 18446744073709551615
                                                    _ = v66
                                                }
                                                switch 2147483647 {
                                                case 42:
                                                    v14 = float32(-2.25)
                                                case -308:
                                                    if false {
                                                        v67 := GenS4{F0: 127, F1: 1}
                                                        _ = v67
                                                    } else {
                                                        v49 -= (v49 >> ((42) & 63))
                                                        v68 := GenS2{F0: "hello", F1: 423}
                                                        v2 -= (v2 / (((v2 ^ (v2 + (v2 & genFn0((fi48 * fi64), v49))))) | 1))
                                                        v49 += v49
                                                        var v69 uint32 = v2
                                                        v2 = (v2 + (v69 - 521))
                                                        _ = v68
                                                        _ = v69
                                                    }
                                                }
                                                fmt.Println(v14)
                                                _ = fi64
                                            }
                                            v3 -= v3
                                        }
                                        v70 := []int{fi48, fi48}
                                        _ = v70
                                    }
                                    _ = fi48
                                    _ = v49
                                }
                                v71 := []int{}
                                v2 = (v2 % ((genFn0(4294967295, 42)) | 1))
                                _ = v71
                            } else {
                                var v72 uint16 = 2
                                _ = v72
                            }
                        case -475:
                            fmt.Println(v14)
                        default:
                            var v73 uint32 = v2
                            _ = v73
                        }
                    }
                }
            case "hello":
                var v74 uint8 = 42
                _ = v74
            case "foo bar":
                var v75 int32 = 827
                _ = v75
            default:
                var v76 bool = false
                _ = v76
            }
        }
        _ = v15
    } else {
        var v77 uint64 = 4294967295
        v14 = v14
        _ = v77
    }
    switch "" {
    case "a":
        v14 -= v14
    }
    var v78 string = ("foo bar" + "")
    v14 -= float32((1.5 * float64(float32(0.1))))
    var v79 uint64 = 1
    fmt.Println(v14, v3, v79)
    if true {
        v80 := GenS5{F0: 1, F1: 0}
        _ = v80
    }
    var v81 int = 128
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v78:", v78)
    fmt.Println("v79:", v79)
    fmt.Println("v81:", v81)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v78
    _ = v79
    _ = v81
}
