package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64 }
type GenS2 struct { F0 int8; F1 uint16 }
type GenS3 struct { F0 bool; F1 uint64; F2 int64 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 bool; F1 int64 }
type GenS6 struct { F0 int16; F1 bool }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 uint; F1 int16 }
type GenS9 struct { F0 uint16; F1 uint }
type GenS10 struct { F0 int16; F1 bool; F2 int8 }
type GenS11 struct { F0 int32; F1 string; F2 int16 }
type GenS12 struct { F0 string; F1 int8 }
type GenS13 struct { F0 int; F1 bool }
type GenS14 struct { F0 int32; F1 string }
type GenS15 struct { F0 int64 }
func genFn0(p0 uint8, p1 uint16, p2 int8) int16 {
    v7 := GenS1{F0: 1, F1: 884}
    p1 += (^(p1))
    p1 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return -973
}
func genFn1(p0 uint16, p1 int32, p2 bool) uint {
    fmt.Println(p1)
    v8 := GenS1{F0: 789, F1: 367}
    if (0 != 42) {
        p0 += p0
        var v9 bool = false
        var v10 int8 = 1
        _ = v10
        _ = v9
    } else {
        p2 = true
    }
    p1 = (-(p1))
    var v11 float32 = ((float32(0.1) * 0.0) * 1.5)
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v8
    return 42
}

func main() {
    v1 := make(map[int]bool)
    v2 := []uint32{516}
    var v3 uint32 = 42
    v4 := make(map[uint32]uint32)
    v5 := make(map[string]bool)
    var v6 float64 = 0.0
    v12 := GenS2{F0: 0, F1: 42}
    fmt.Println(v6, v6)
    fmt.Println(v6, v6)
    v3 = v3
    for fi13 := 0; fi13 < 18; fi13++ {
        var v14 int = (fi13 - (fi13 + (fi13 ^ (-(fi13)))))
        _ = fi13
        _ = v14
    }
    switch "a" {
    case "":
        var v15 bool = (-32768 != 1)
        _ = v15
    case "foo bar":
        var v16 float32 = (-2.25 * 0.1)
        _ = v16
    default:
        var v17 int = 128
        _ = v17
    }
    v6 -= 0.0
    if true {
        for fi18 := 0; fi18 < 7; fi18++ {
            v3 = 2
            for fi19 := 0; fi19 < 17; fi19++ {
                fmt.Println(fi18, fi18)
                fmt.Println(v3)
                switch "a" {
                case "":
                    switch fi18 {
                    case 2147483647:
                        v20 := GenS3{F0: (1 == 98), F1: 2, F2: 9223372036854775807}
                        _ = v20
                    default:
                        v6 += 0.0
                    }
                case "foo bar":
                    fmt.Println(v3, fi18, fi19)
                case "a":
                    switch fi19 {
                    case 127:
                        v21 := GenS4{F0: 981}
                        _ = v21
                    }
                default:
                    v6 -= float64(3.14159)
                }
                _ = fi19
            }
            _ = fi18
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        for fi22 := 0; fi22 < 1; fi22++ {
            v23 := GenS5{F0: false, F1: 770}
            switch "\n\t" {
            case "foo bar":
                fmt.Println(v6)
            case "":
                if len(v2) > 0 {
                    v2[0] = 42
                }
            default:
                if true {
                    delete(v4, v3)
                    if ((v3 & v3) >= v3) {
                        var v24 float64 = v6
                        if true {
                            var v25 int = fi22
                            _ = v25
                        }
                        _ = v24
                    } else {
                        var v26 float32 = 3.14159
                        _ = v26
                    }
                } else {
                    v6 += (v6 / 1.0)
                }
            }
            _ = fi22
            _ = v23
        }
    } else {
        var v27 float64 = (v6 - ((-2.25 / float64((1e10 - (0.0 / 1.5)))) * v6))
        for fi28 := 0; fi28 < 16; fi28++ {
            if true {
                switch 42 {
                case 32767:
                    for fi29 := 0; fi29 < 6; fi29++ {
                        v3 = v3
                        if false {
                            var v30 int = 0
                            _ = v30
                        } else {
                            v27 += (v27 / v6)
                            switch 1 {
                            case 328:
                                var v31 uint32 = v3
                                _ = v31
                            case -383:
                                var v32 uint32 = 4294967295
                                _ = v32
                            }
                        }
                        _ = fi29
                    }
                default:
                    v6 = float64(1e10)
                }
                v33 := GenS6{F0: genFn0(2, 255, 1), F1: (true || false)}
                v3 = v3
                _ = v33
            }
            fmt.Println(fi28, v6, v6)
            if true {
                _ = len(v2)
            }
            _ = fi28
        }
        if false {
            for fi34 := 0; fi34 < 12; fi34++ {
                if len(v2) > 0 {
                    v2[0] = v3
                }
                var v35 uint = 245
                _ = fi34
                _ = v35
            }
            fmt.Println(v6, v27, v27)
            fmt.Println(v3)
        } else {
            switch "foo bar" {
            case "hello":
                switch -9223372036854775808 {
                case 0:
                    fmt.Println(v3, v3, v27)
                case 32768:
                    v36 := GenS7{F0: -613}
                    _ = v36
                }
            case "foo bar":
                var v37 float64 = ((-2.25 - (1.5 + float64((-2.25 - -2.25)))) * ((v6 / (v6 + v6)) - float64(3.14159)))
                _ = v37
            case "\n\t":
                if false {
                    v27 -= v27
                    if false {
                        if false {
                            switch 32767 {
                            case 905:
                                v38 := GenS3{F0: (1 >= 1), F1: 0, F2: 1}
                                _ = v38
                            case -934:
                                v6 = (-2.25 * v27)
                            case -1:
                                var v39 int = -429
                                _ = v39
                            default:
                                var v40 int32 = 0
                                _ = v40
                            }
                            if (!true) {
                                if true {
                                    v41 := GenS8{F0: 65535, F1: genFn0(1, 65535, 0)}
                                    v5[""] = false
                                    if true {
                                        fmt.Println(v27)
                                    } else {
                                        switch 1 {
                                        case -554:
                                            var v42 int16 = -32768
                                            _ = v42
                                        }
                                        v27 -= ((v6 / (v6 + v27)) / v27)
                                    }
                                    _ = v41
                                } else {
                                    delete(v1, 452)
                                    v43 := GenS9{F0: 2, F1: 255}
                                    _ = v43
                                }
                            } else {
                                v2 = append(v2, (v3 / ((4294967295) | 1)))
                                v6 = (float64(float32((v27 + v6))) - (((((v27 / 0.1) - (-2.25 - v6)) * 1e10) / v27) / v6))
                            }
                            var v44 uint16 = 65535
                            _ = v44
                        } else {
                            switch 225 {
                            case 537:
                                fmt.Println(v6)
                            case -830:
                                fmt.Println(v6, v27, v27)
                            case 32767:
                                if false {
                                    if false {
                                        v45 := GenS2{F0: 97, F1: 255}
                                        if false {
                                            var v46 uint64 = 170
                                            if (v46 > v46) {
                                                var v47 int16 = 1
                                                switch "foo bar" {
                                                case "":
                                                    var v48 int32 = 2147483647
                                                    _ = v48
                                                default:
                                                    switch "\n\t" {
                                                    case "hello":
                                                        var v49 uint64 = v46
                                                        _ = v49
                                                    }
                                                }
                                                if (32767 != 255) {
                                                    var v50 int8 = -123
                                                    _ = v50
                                                } else {
                                                    v46 += v46
                                                    switch (("\n\t" + "foo bar") + "hello") {
                                                    case "":
                                                        var v51 int16 = (v47 - (v47 ^ v47))
                                                        _ = v51
                                                    case "foo bar":
                                                        v6 = 0.1
                                                    }
                                                    if false {
                                                        var v52 int32 = -1
                                                        _ = v52
                                                    } else {
                                                        var v53 float64 = (-2.25 / 0.1)
                                                        _ = v53
                                                    }
                                                }
                                                _ = v47
                                            }
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                            _ = v46
                                        }
                                        _ = v45
                                    } else {
                                        fmt.Println(v27, v6, v3)
                                        v2 = append(v2, 65535)
                                        var v54 uint64 = 1
                                        _ = v54
                                    }
                                }
                            default:
                                for fi55 := 0; fi55 < 3; fi55++ {
                                    var v56 bool = (1.5 <= (v6 * 0.0))
                                    _ = fi55
                                    _ = v56
                                }
                            }
                            for fi57 := 0; fi57 < 11; fi57++ {
                                var v58 int16 = int16(v3)
                                _ = fi57
                                _ = v58
                            }
                        }
                    } else {
                        v59 := GenS5{F0: false, F1: 9223372036854775807}
                        _ = v59
                    }
                }
            default:
                v1[520] = false
            }
        }
        _ = v27
    }
    var v60 int64 = 42
    switch (("a" + "foo bar") + ("foo bar" + ("\n\t" + "a"))) {
    case "a":
        v6 = v6
    case "\n\t":
        if (!((1.5 * 0.1) >= 1e10)) {
            for fi61 := 0; fi61 < 14; fi61++ {
                if false {
                    switch "foo bar" {
                    case "hello":
                        v62 := GenS10{F0: -1, F1: (!true), F2: 100}
                        _ = v62
                    case "":
                        var v63 float32 = ((0.1 * 0.1) - ((float32((v6 / 1.0)) + float32((3.14159 - v6))) + 1e10))
                        _ = v63
                    case "foo bar":
                        v64 := GenS11{F0: 1, F1: "\n\t", F2: 0}
                        _ = v64
                    }
                    for fi65 := 0; fi65 < 3; fi65++ {
                        if (1e10 == -2.25) {
                            v3 = (v3 | v3)
                        }
                        v3 -= 2
                        var v66 int16 = 422
                        v67 := v5[""]
                        _ = fi65
                        _ = v66
                        _ = v67
                    }
                }
                switch ("hello" + "a") {
                case "hello":
                    var v68 int = (fi61 >> ((18446744073709551615) & 63))
                    _ = v68
                case "\n\t":
                    fmt.Println(v6, v3, v3)
                case "a":
                    v6 -= v6
                default:
                    if (1.5 != ((0.1 * (float32(0.0) * 1.5)) * 1.5)) {
                        _ = len(v2)
                        v69 := GenS12{F0: "a", F1: -128}
                        for fi70 := 0; fi70 < 15; fi70++ {
                            if true {
                                v60 += (v60 + (v60 << ((4294967295) & 63)))
                            } else {
                                v2 = append(v2, v3)
                                switch "\n\t" {
                                case "\n\t":
                                    var v71 float32 = 1e10
                                    _ = v71
                                case "":
                                    _ = len(v4)
                                }
                            }
                            delete(v5, "hello")
                            _ = fi70
                        }
                        _ = v69
                    }
                }
                _ = fi61
            }
            if true {
                v3 = 4294967295
                var v72 float64 = float64(((1.5 / (float32(v6) - 1.5)) / -2.25))
                if (0 <= 663) {
                    var v73 bool = (true || true)
                    for fi74 := 0; fi74 < 13; fi74++ {
                        var v75 bool = (42 >= 0)
                        v3 = v3
                        v72 = float64(float32(v72))
                        _ = fi74
                        _ = v75
                    }
                    switch "a" {
                    case "":
                        fmt.Println(v60)
                    }
                    _ = v73
                } else {
                    if true {
                        v76 := GenS12{F0: "hello", F1: 100}
                        v77 := GenS13{F0: 0, F1: (v60 < v60)}
                        _ = v76
                        _ = v77
                    }
                    for fi78 := 0; fi78 < 4; fi78++ {
                        fmt.Println(v3, v3, v60)
                        for fi79 := 0; fi79 < 16; fi79++ {
                            v80 := GenS14{F0: -773, F1: ""}
                            fmt.Println(fi78, v72)
                            _ = len(v4)
                            _ = fi79
                            _ = v80
                        }
                        v72 = float64(0.1)
                        _ = fi78
                    }
                }
                _ = v72
            } else {
                for fi81 := 0; fi81 < 20; fi81++ {
                    if ((0 == 255) && (!false)) {
                        var v82 int64 = v60
                        _ = v82
                    }
                    _ = fi81
                }
                v6 -= v6
            }
            v3 -= v3
        } else {
            v83 := GenS9{F0: 874, F1: 4294967295}
            var v84 uint = 1
            _ = v83
            _ = v84
        }
    case "hello":
        if (v3 != (^(v3))) {
            v85 := GenS1{F0: 0, F1: v60}
            _ = v85
        } else {
            var v86 uint16 = 434
            _ = v86
        }
    default:
        v87 := GenS10{F0: genFn0(42, 466, 127), F1: (29 >= (v3 + (v3 / ((v3) | 1)))), F2: 127}
        _ = v87
    }
    v88 := GenS15{F0: v60}
    var v89 float64 = v6
    var v90 float64 = v6
    var v91 uint = 65535
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v60:", v60)
    fmt.Println("v89:", v89)
    fmt.Println("v90:", v90)
    fmt.Println("v91:", v91)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v60
    _ = v88
    _ = v89
    _ = v90
    _ = v91
}
