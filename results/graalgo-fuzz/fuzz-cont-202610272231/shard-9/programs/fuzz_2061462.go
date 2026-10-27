package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool }
type GenS2 struct { F0 uint8; F1 uint8; F2 uint8 }
type GenS3 struct { F0 int32; F1 int }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 uint32 }
func genFn0(p0 uint64, p1 bool, p2 bool) uint {
    var v7 int32 = 255
    for fi8 := 0; fi8 < 20; fi8++ {
        for fi9 := 0; fi9 < 4; fi9++ {
            var v10 int8 = -28
            if (!p2) {
                var v11 uint16 = 42
                v12 := GenS1{F0: p2, F1: (!p1)}
                _ = v11
                _ = v12
            } else {
                var v13 int16 = 236
                switch ("a" + "") {
                case "hello":
                    if true {
                        var v14 uint = uint(p0)
                        _ = v14
                    }
                case "\n\t":
                    fmt.Println(v13, p0)
                case "foo bar":
                    fmt.Println(v13, p2, fi9)
                }
                switch (^(fi8)) {
                case 32767:
                    var v15 uint32 = 65535
                    _ = v15
                default:
                    if p2 {
                        v16 := GenS1{F0: (128 < uint8(p0)), F1: (v13 > -32768)}
                        var v17 int64 = 9223372036854775807
                        var v18 int16 = 613
                        var v19 int64 = (v17 % ((-1) | 1))
                        v17 = (v19 * (v17 / ((v19) | 1)))
                        _ = v16
                        _ = v17
                        _ = v18
                        _ = v19
                    } else {
                        fmt.Println(p1, v7)
                    }
                }
                _ = v13
            }
            _ = fi9
            _ = v10
        }
        _ = fi8
    }
    v7 = (v7 % ((int32(p0)) | 1))
    v20 := []int{128}
    _ = p0
    _ = p1
    _ = p2
    _ = v20
    _ = v7
    return 4294967295
}

func main() {
    v1 := make(map[int]int)
    v2 := []string{}
    v3 := []uint32{1, 121}
    var v4 int8 = -1
    v5 := make(map[uint32]int)
    v6 := make(map[int]int)
    if (-357 == 0) {
        var v21 int8 = 100
        _ = v21
    }
    var v22 uint16 = 255
    v23 := GenS1{F0: true, F1: (false || (!(-2147483648 == -2147483648)))}
    switch "hello" {
    case "foo bar":
        if true {
            for fi24 := 0; fi24 < 5; fi24++ {
                v25 := GenS2{F0: 42, F1: 71, F2: 42}
                _ = fi24
                _ = v25
            }
            var v26 uint32 = 938
            _ = v26
        } else {
            if (!false) {
                v27 := GenS3{F0: 42, F1: 32768}
                if len(v3) > 0 {
                    v3[0] = 1
                }
                fmt.Println(v4)
                _ = v27
            } else {
                v28 := GenS1{F0: false, F1: false}
                if (((0.1 / -2.25) != 1.5) && false) {
                    for fi29 := 0; fi29 < 12; fi29++ {
                        if true {
                            v22 = v22
                            if ((0 >= fi29) && (0 == 0)) {
                                if len(v2) > 0 {
                                    v2[0] = "hello"
                                }
                            }
                        } else {
                            v4 = 127
                            if true {
                                fmt.Println(fi29, v22, fi29)
                                v30 := GenS3{F0: 354, F1: 1}
                                _ = v30
                            } else {
                                fmt.Println(v4)
                            }
                        }
                        var v31 int8 = v4
                        _ = fi29
                        _ = v31
                    }
                    _ = len(v3)
                    var v32 int8 = (^(v4))
                    _ = v32
                }
                _ = v28
            }
        }
    }
    for fi33 := 0; fi33 < 19; fi33++ {
        for fi34 := 0; fi34 < 20; fi34++ {
            v35 := GenS1{F0: false, F1: true}
            var v36 int16 = 0
            if false {
                v4 -= (v4 << ((94) & 63))
                if (!true) {
                    var v37 uint16 = v22
                    var v38 int = int(v37)
                    _ = v37
                    _ = v38
                } else {
                    v4 -= 95
                    if (true || true) {
                        fmt.Println(v36)
                        var v39 int8 = 86
                        var v40 int32 = 42
                        _ = v39
                        _ = v40
                    }
                }
            }
            var v41 int8 = 0
            _ = fi34
            _ = v35
            _ = v36
            _ = v41
        }
        _ = fi33
    }
    v4 = v4
    if (true || false) {
        fmt.Println(v4, v22)
        if true {
            var v42 float64 = float64(float32(3.14159))
            for fi43 := 0; fi43 < 5; fi43++ {
                v3 = append(v3, 65535)
                if true {
                    var v44 uint8 = 127
                    for fi45 := 0; fi45 < 7; fi45++ {
                        fmt.Println(v44, v44, v42)
                        switch 32767 {
                        case -29:
                            v42 -= (float64(-2.25) - 0.0)
                        default:
                            v42 = v42
                        }
                        var v46 uint32 = 42
                        fmt.Println(v44)
                        if true {
                            v44 += (v44 % ((2) | 1))
                            v47 := GenS3{F0: 2147483647, F1: fi45}
                            _ = v47
                        }
                        var v48 int8 = 0
                        _ = fi45
                        _ = v46
                        _ = v48
                    }
                    fmt.Println(v22)
                    _ = v44
                }
                v49 := GenS1{F0: false, F1: (533 > 18446744073709551615)}
                _ = fi43
                _ = v49
            }
            _ = v42
        }
    } else {
        if len(v3) > 0 {
            v3[0] = 65535
        }
        var v50 int64 = 0
        var v51 int32 = 42
        _ = v50
        _ = v51
    }
    if true {
        if (false && false) {
            if (0 < 935) {
                v22 = 255
            } else {
                v22 -= (v22 / ((65535) | 1))
                fmt.Println(v22, v4)
            }
        }
        fmt.Println(v22)
        if true {
            v52 := GenS4{F0: 1}
            _ = v52
        }
    } else {
        for fi53 := 0; fi53 < 12; fi53++ {
            var v54 string = ("hello" + "a")
            if true {
                for fi55 := 0; fi55 < 19; fi55++ {
                    if false {
                        v6[9223372036854775807] = fi53
                    }
                    var v56 uint64 = 0
                    _ = fi55
                    _ = v56
                }
                for fi57 := 0; fi57 < 18; fi57++ {
                    var v58 float64 = -2.25
                    var v59 uint8 = 221
                    _ = fi57
                    _ = v58
                    _ = v59
                }
            }
            if ((v22 & v22) < v22) {
                v60 := GenS1{F0: false, F1: ((false || ((796 <= 2) && ((fi53 == 937) && false))) || (true && false))}
                switch 127 {
                case 32768:
                    switch ((v54 + ((v54 + ("foo bar" + "hello")) + v54)) + "") {
                    case "foo bar":
                        if (!(false || (-1 == v4))) {
                            if (!false) {
                                switch (v54 + "hello") {
                                case "a":
                                    if (!(v4 <= (v4 - (v4 % (((-(v4))) | 1))))) {
                                        v61 := GenS5{F0: v22}
                                        _ = v61
                                    }
                                }
                                if len(v2) > 0 {
                                    v2[0] = (v54 + (v54 + (v54 + v54)))
                                }
                            } else {
                                for fi62 := 0; fi62 < 8; fi62++ {
                                    delete(v1, int(v22))
                                    if true {
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        var v63 uint = 944
                                        fmt.Println(v54)
                                        _ = v63
                                    }
                                    _ = fi62
                                }
                            }
                        } else {
                            if (true && (!(42 != 655))) {
                                switch 684 {
                                case 32767:
                                    for fi64 := 0; fi64 < 16; fi64++ {
                                        var v65 bool = false
                                        var v66 float64 = 0.0
                                        _ = fi64
                                        _ = v65
                                        _ = v66
                                    }
                                default:
                                    fmt.Println(fi53)
                                }
                                v22 = v22
                                for fi67 := 0; fi67 < 6; fi67++ {
                                    v22 -= 810
                                    _ = fi67
                                }
                            }
                            var v68 int16 = 457
                            var v69 float32 = 0.1
                            _ = v68
                            _ = v69
                        }
                    }
                case 128:
                    var v70 string = v54
                    _ = v70
                default:
                    v22 -= (v22 / (((v22 & (v22 % (((v22 & v22)) | 1)))) | 1))
                }
                _ = v60
            }
            _ = fi53
            _ = v54
        }
        for fi71 := 0; fi71 < 4; fi71++ {
            var v72 uint64 = 0
            _ = fi71
            _ = v72
        }
        v73 := GenS6{F0: 1}
        _ = v73
    }
    if (65535 > 65535) {
        v74 := GenS4{F0: -969}
        if false {
            switch -984 {
            case 513:
                v22 = v22
            case -190:
                for fi75 := 0; fi75 < 19; fi75++ {
                    v76 := GenS6{F0: 42}
                    var v77 uint = 2
                    _ = fi75
                    _ = v76
                    _ = v77
                }
            default:
                var v78 string = "hello"
                _ = v78
            }
        }
        _ = v74
    }
    var v79 int32 = 42
    fmt.Println(v79)
    if false {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    for fi80 := 0; fi80 < 18; fi80++ {
        var v81 uint16 = 255
        for fi82 := 0; fi82 < 11; fi82++ {
            var v83 bool = (42 != 23)
            _ = fi82
            _ = v83
        }
        _ = fi80
        _ = v81
    }
    var v84 int64 = 0
    fmt.Println(v4, v22, v22)
    var v85 string = "foo bar"
    var v86 int16 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v79:", v79)
    fmt.Println("v84:", v84)
    fmt.Println("v85:", v85)
    fmt.Println("v86:", v86)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v79
    _ = v84
    _ = v85
    _ = v86
}
