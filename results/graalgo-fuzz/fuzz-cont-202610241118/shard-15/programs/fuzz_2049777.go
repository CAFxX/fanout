package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int8; F1 uint32; F2 int16 }
type GenS3 struct { F0 int16; F1 bool; F2 uint32 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int32; F1 int32 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 uint16; F1 bool }
type GenS8 struct { F0 int16; F1 uint16; F2 int16 }
type GenS9 struct { F0 uint16; F1 int32; F2 int64 }
func genFn0(p0 uint, p1 uint64) int16 {
    if (62 > 54) {
        switch "a" {
        case "\n\t":
            v5 := []int{-1}
            _ = v5
        }
        p1 -= (p1 | (p1 >> ((65535) & 63)))
    }
    v6 := GenS1{F0: 0}
    var v7 uint16 = 65535
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return -1
}
func genFn1() string {
    for fi8 := 0; fi8 < 2; fi8++ {
        v9 := GenS1{F0: (^(fi8))}
        _ = fi8
        _ = v9
    }
    v10 := GenS2{F0: 0, F1: 2, F2: genFn0(1, 4294967295)}
    var v11 uint8 = 1
    v11 += v11
    if (42 <= -32768) {
        v11 = (v11 >> ((47) & 63))
    } else {
        switch "hello" {
        case "foo bar":
            switch 867 {
            case 484:
                v12 := []int{0, -2147483648}
                _ = v12
            case -1:
                fmt.Println(v11)
            default:
                for fi13 := 0; fi13 < 3; fi13++ {
                    for fi14 := 0; fi14 < 9; fi14++ {
                        var v15 int8 = 127
                        _ = fi14
                        _ = v15
                    }
                    if false {
                        v16 := []int{fi13, fi13, fi13}
                        for fi17 := 0; fi17 < 8; fi17++ {
                            var v18 bool = false
                            _ = fi17
                            _ = v18
                        }
                        _ = v16
                    }
                    _ = fi13
                }
            }
        }
    }
    _ = v10
    _ = v11
    return ""
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint16 = 0
    var v3 int64 = 0
    var v4 uint32 = uint32(v2)
    v3 -= v3
    v3 -= v3
    v2 += v2
    for fi19 := 0; fi19 < 1; fi19++ {
        v20 := GenS3{F0: 1000, F1: true, F2: v4}
        _ = fi19
        _ = v20
    }
    v21 := GenS2{F0: -103, F1: 367, F2: 543}
    if true {
        if false {
            for fi22 := 0; fi22 < 13; fi22++ {
                for fi23 := 0; fi23 < 8; fi23++ {
                    var v24 int = fi22
                    _ = fi23
                    _ = v24
                }
                v25 := GenS4{F0: 18446744073709551615}
                v26 := []int{(^(fi22)), -510}
                _ = fi22
                _ = v25
                _ = v26
            }
            switch "a" {
            case "a":
                for fi27 := 0; fi27 < 16; fi27++ {
                    fmt.Println(v4, v2, fi27)
                    v28 := []int{(fi27 * 42), 32768}
                    _ = fi27
                    _ = v28
                }
            case "hello":
                v4 = (v4 | v4)
            default:
                if false {
                    v3 = 0
                }
            }
            v29 := GenS5{F0: -2147483648, F1: 0}
            _ = v29
        } else {
            switch "a" {
            case "":
                v30 := GenS2{F0: -66, F1: v4, F2: 1000}
                _ = v30
            default:
                v31 := []int{4294967295, 9223372036854775807}
                _ = v31
            }
            v32 := GenS5{F0: 1, F1: 42}
            delete(v1, v4)
            _ = v32
        }
        v33 := GenS1{F0: 32767}
        _ = v33
    }
    v4 = (v4 | uint32(v2))
    v34 := GenS6{F0: -2147483648}
    if (!false) {
        v4 = (v4 & 65535)
        if false {
            fmt.Println(v4)
            if false {
                var v35 int8 = -1
                _ = v35
            }
            v36 := GenS3{F0: 0, F1: true, F2: (^(v4))}
            _ = v36
        } else {
            v37 := GenS1{F0: 4294967295}
            var v38 int8 = 23
            _ = v37
            _ = v38
        }
    } else {
        v3 -= (v3 << ((0) & 63))
    }
    if true {
        var v39 int64 = v3
        _ = v39
    } else {
        for fi40 := 0; fi40 < 5; fi40++ {
            v41 := GenS4{F0: 1}
            for fi42 := 0; fi42 < 11; fi42++ {
                v43 := GenS3{F0: genFn0(42, 85), F1: ((v2 % ((v2) | 1)) == 255), F2: (v4 % (((v4 >> ((4294967295) & 63))) | 1))}
                _ = fi42
                _ = v43
            }
            fmt.Println(v2, v4)
            _ = fi40
            _ = v41
        }
        switch 42 {
        case -2147483648:
            if (32767 != genFn0(255, 311)) {
                v2 = v2
                var v44 string = ("" + ("a" + "hello"))
                _ = v44
            }
        default:
            if (553 == -2147483648) {
                switch ((("" + "foo bar") + ("hello" + ("foo bar" + "foo bar"))) + "a") {
                case "hello":
                    v45 := v1[(v4 >> ((101) & 63))]
                    _ = v45
                case "":
                    for fi46 := 0; fi46 < 8; fi46++ {
                        v47 := GenS6{F0: 0}
                        switch (fi46 | (fi46 * 0)) {
                        case -9223372036854775808:
                            v3 = (v3 * (v3 + (v3 & 256)))
                        case 628:
                            var v48 uint32 = v4
                            _ = v48
                        case 32768:
                            delete(v1, (v4 / ((v4) | 1)))
                        default:
                            for fi49 := 0; fi49 < 7; fi49++ {
                                v50 := GenS7{F0: v2, F1: false}
                                v4 += v4
                                v2 += (v2 >> ((42) & 63))
                                _ = fi49
                                _ = v50
                            }
                        }
                        for fi51 := 0; fi51 < 18; fi51++ {
                            _ = len(v1)
                            _ = fi51
                        }
                        _ = fi46
                        _ = v47
                    }
                }
                var v52 int64 = -117
                _ = v52
            }
        }
        if false {
            v4 = (v4 % (((^(v4))) | 1))
            v3 -= (v3 & (v3 - v3))
            for fi53 := 0; fi53 < 18; fi53++ {
                for fi54 := 0; fi54 < 4; fi54++ {
                    var v55 float64 = 0.0
                    v4 -= v4
                    _ = fi54
                    _ = v55
                }
                if (!(4294967295 >= 481)) {
                    var v56 int8 = 0
                    switch (-(fi53)) {
                    case 4294967295:
                        for fi57 := 0; fi57 < 4; fi57++ {
                            delete(v1, v4)
                            v4 += 2
                            _ = fi57
                        }
                    default:
                        if (true && false) {
                            for fi58 := 0; fi58 < 6; fi58++ {
                                v59 := []int{fi58, 679, (fi53 / ((128) | 1))}
                                var v60 float32 = (3.14159 + 1e10)
                                fmt.Println(fi53)
                                _ = fi58
                                _ = v59
                                _ = v60
                            }
                            v3 -= v3
                            var v61 float64 = 1e10
                            _ = v61
                        } else {
                            for fi62 := 0; fi62 < 7; fi62++ {
                                switch (fi53 * -993) {
                                case 481:
                                    var v63 uint64 = 554
                                    _ = v63
                                case -1:
                                    switch fi62 {
                                    case -434:
                                        var v64 int32 = 2147483647
                                        _ = v64
                                    default:
                                        switch "foo bar" {
                                        case "hello":
                                            if (32767 == 32767) {
                                                var v65 int32 = 42
                                                _ = v65
                                            } else {
                                                var v66 float32 = (-2.25 * 3.14159)
                                                if false {
                                                    v2 = (^(v2))
                                                    if true {
                                                        var v67 int32 = -2147483648
                                                        _ = v67
                                                    } else {
                                                        v3 += v3
                                                    }
                                                    fmt.Println(v56)
                                                }
                                                fmt.Println(v56)
                                                _ = v66
                                            }
                                        }
                                    }
                                }
                                switch "" {
                                case "foo bar":
                                    v4 = 65535
                                }
                                if (!(621 < 72)) {
                                    var v68 string = "a"
                                    _ = v68
                                } else {
                                    switch "a" {
                                    case "foo bar":
                                        v69 := []int{fi62, int(v56)}
                                        _ = v69
                                    }
                                }
                                _ = fi62
                            }
                            v56 = v56
                            var v70 string = "hello"
                            _ = v70
                        }
                    }
                    _ = v56
                } else {
                    v4 += (v4 + (v4 + 2))
                    v71 := GenS8{F0: -32768, F1: v2, F2: 920}
                    for fi72 := 0; fi72 < 7; fi72++ {
                        v73 := GenS4{F0: 448}
                        if (!true) {
                            var v74 float32 = 0.1
                            _ = v74
                        } else {
                            var v75 int = fi72
                            _ = v75
                        }
                        if false {
                            if false {
                                if true {
                                    var v76 int32 = -2147483648
                                    v4 = (v4 >> ((42) & 63))
                                    var v77 int32 = v76
                                    var v78 int32 = 2147483647
                                    var v79 int16 = 1000
                                    _ = v76
                                    _ = v77
                                    _ = v78
                                    _ = v79
                                } else {
                                    var v80 int64 = v3
                                    var v81 float64 = float64(0.1)
                                    _ = v80
                                    _ = v81
                                }
                                v2 += (v2 - v2)
                            } else {
                                v1[65535] = (-128 == 1)
                                fmt.Println(v2, v4, v4)
                            }
                        } else {
                            if true {
                                fmt.Println(v3)
                            }
                            switch 32768 {
                            case -440:
                                fmt.Println(fi72, v3, fi72)
                                v82 := []int{}
                                _ = v82
                            default:
                                var v83 int32 = 1
                                _ = v83
                            }
                            v1[(v4 | (v4 + v4))] = true
                            fmt.Println(v2)
                        }
                        fmt.Println(v2)
                        _ = fi72
                        _ = v73
                    }
                    _ = v71
                }
                var v84 uint8 = 2
                _ = fi53
                _ = v84
            }
        }
    }
    fmt.Println(v4, v2, v3)
    v3 = v3
    fmt.Println(v2, v4)
    for fi85 := 0; fi85 < 8; fi85++ {
        if true {
            var v86 int8 = -1
            v87 := GenS7{F0: (v2 << ((195) & 63)), F1: true}
            if true {
                switch ("\n\t" + "hello") {
                case "hello":
                    v88 := GenS2{F0: (v86 % ((v86) | 1)), F1: v4, F2: genFn0(0, 42)}
                    _ = v88
                case "\n\t":
                    var v89 bool = (0.1 < 0.1)
                    _ = v89
                case "foo bar":
                    v90 := []int{}
                    _ = v90
                }
                fmt.Println(v2, fi85, v4)
            } else {
                switch "" {
                case "a":
                    if ((193 != 965) && false) {
                        delete(v1, 2147483647)
                    }
                case "hello":
                    _ = len(v1)
                }
                for fi91 := 0; fi91 < 2; fi91++ {
                    v4 = 493
                    fmt.Println(v3, fi85, v2)
                    switch (("\n\t" + "a") + ("foo bar" + ("foo bar" + "a"))) {
                    case "foo bar":
                        if (2147483647 < 0) {
                            var v92 uint8 = 167
                            _ = v92
                        }
                    case "":
                        v86 = v86
                    case "hello":
                        var v93 uint8 = 185
                        _ = v93
                    }
                    if false {
                        var v94 float32 = 0.1
                        _ = v94
                    } else {
                        switch (fi91 % (((^(fi91))) | 1)) {
                        case 42:
                            var v95 float64 = 0.0
                            _ = v95
                        default:
                            v86 = v86
                        }
                        var v96 uint64 = 42
                        _ = v96
                    }
                    v1[(v4 / ((4294967295) | 1))] = false
                    _ = fi91
                }
                fmt.Println(fi85, v2)
            }
            _ = v86
            _ = v87
        } else {
            fmt.Println(v3, v3, v2)
            v1[(v4 & v4)] = (3.14159 < ((((1.5 - ((3.14159 * 1.5) - 0.0)) / 1.5) - 0.0) - 1.5))
        }
        v97 := []int{fi85, (fi85 & fi85), fi85, fi85}
        _ = len(v1)
        _ = fi85
        _ = v97
    }
    v3 -= (^(v3))
    for fi98 := 0; fi98 < 6; fi98++ {
        v99 := GenS9{F0: v2, F1: -51, F2: (v3 | (v3 * v3))}
        _ = fi98
        _ = v99
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v34
    _ = v4
}
