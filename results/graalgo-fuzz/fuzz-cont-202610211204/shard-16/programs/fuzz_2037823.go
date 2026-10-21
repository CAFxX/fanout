package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string; F2 uint8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 string; F1 uint64 }
type GenS4 struct { F0 uint16; F1 int32; F2 uint64 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint64; F1 uint16 }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 uint; F1 uint8; F2 int64 }
type GenS9 struct { F0 int64; F1 int32; F2 int64 }
type GenS10 struct { F0 uint8 }
type GenS11 struct { F0 uint }
func genFn0() uint64 {
    var v5 uint16 = 57
    v5 -= v5
    for fi6 := 0; fi6 < 20; fi6++ {
        var v7 int32 = 1
        if (127 != 86) {
            v8 := []int{fi6, (fi6 + fi6)}
            switch fi6 {
            case 9223372036854775807:
                v9 := GenS1{F0: v5, F1: ("foo bar" + "foo bar"), F2: uint8(v5)}
                _ = v9
            }
            var v10 float64 = float64(1e10)
            _ = v10
            _ = v8
        } else {
            v5 = 1
        }
        _ = fi6
        _ = v7
    }
    switch 128 {
    case -887:
        for fi11 := 0; fi11 < 14; fi11++ {
            var v12 int8 = 26
            v13 := GenS2{F0: fi11}
            for fi14 := 0; fi14 < 18; fi14++ {
                if true {
                    v5 = (v5 - v5)
                    v5 += 255
                    var v15 int8 = v12
                    _ = v15
                }
                v5 -= (^(v5))
                v16 := GenS1{F0: v5, F1: "\n\t", F2: 2}
                _ = fi14
                _ = v16
            }
            _ = fi11
            _ = v12
            _ = v13
        }
    case 663:
        if (((3.14159 - 0.1) / float64(-2.25)) == 1e10) {
            v5 = (^(v5))
            v17 := GenS3{F0: "a", F1: 2}
            var v18 uint64 = 1
            _ = v17
            _ = v18
        }
    case 9223372036854775807:
        for fi19 := 0; fi19 < 17; fi19++ {
            v20 := GenS1{F0: 0, F1: "a", F2: 195}
            var v21 string = (("a" + (("\n\t" + "foo bar") + "")) + "\n\t")
            _ = fi19
            _ = v20
            _ = v21
        }
    }
    var v22 bool = true
    _ = v22
    _ = v5
    return 4294967295
}
func genFn1(p0 int, p1 uint64, p2 uint8) int64 {
    switch "" {
    case "foo bar":
        p0 = (p0 % ((1) | 1))
    default:
        v23 := GenS3{F0: "foo bar", F1: (p1 % ((41) | 1))}
        _ = v23
    }
    v24 := GenS4{F0: 348, F1: 42, F2: (p1 * (p1 ^ (p1 * (p1 >> ((19) & 63)))))}
    p1 = (p1 << ((970) & 63))
    if true {
        if (0 == 1) {
            var v25 float64 = 0.1
            _ = v25
        }
    } else {
        fmt.Println(p2)
        p1 += (^(p1))
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v24
    return -381
}

func main() {
    v1 := make(map[int]string)
    var v2 float32 = 0.1
    v3 := []int{32768}
    var v4 int8 = 0
    v26 := GenS2{F0: 0}
    v27 := GenS5{F0: true}
    for fi28 := 0; fi28 < 7; fi28++ {
        switch "hello" {
        case "foo bar":
            v2 -= (((v2 + v2) + 1.5) * -2.25)
        case "hello":
            if true {
                switch fi28 {
                case 32768:
                    fmt.Println(v2, v2)
                }
                if true {
                    v4 = 71
                    if (!false) {
                        v29 := GenS3{F0: "foo bar", F1: genFn0()}
                        if true {
                            v30 := GenS5{F0: (!(false && true))}
                            _ = v30
                        } else {
                            v4 += v4
                            fmt.Println(fi28)
                            delete(v1, -124)
                        }
                        var v31 int16 = 0
                        _ = v29
                        _ = v31
                    }
                } else {
                    var v32 uint8 = 127
                    _ = v32
                }
            }
        case "":
            var v33 uint8 = 2
            _ = v33
        }
        var v34 uint16 = 2
        _ = fi28
        _ = v34
    }
    var v35 float32 = float32(float64((-2.25 * v2)))
    var v36 uint = 588
    v37 := GenS6{F0: 503, F1: 42}
    fmt.Println(v35, v4, v4)
    var v38 int64 = genFn1(-921, 0, 127)
    var v39 uint16 = 32767
    if (42 <= 0) {
        var v40 uint = (^(v36))
        fmt.Println(v2, v35, v36)
        var v41 int32 = -1
        _ = v40
        _ = v41
    } else {
        _ = len(v3)
        for fi42 := 0; fi42 < 8; fi42++ {
            v43 := GenS7{F0: (v38 | (-(v38)))}
            var v44 int16 = 1000
            switch (fi42 % ((0) | 1)) {
            case -9223372036854775808:
                fmt.Println(v36, fi42, v4)
            case 4294967295:
                if true {
                    for fi45 := 0; fi45 < 10; fi45++ {
                        fmt.Println(v2, fi45, v2)
                        v2 += ((v35 / 0.1) + 0.0)
                        var v46 uint64 = 490
                        fmt.Println(v35, v36, v44)
                        _ = fi45
                        _ = v46
                    }
                    if false {
                        v44 = -1
                        var v47 int = (fi42 + fi42)
                        _ = v47
                    } else {
                        v48 := v1[128]
                        var v49 int16 = v44
                        var v50 float64 = 0.1
                        _ = v48
                        _ = v49
                        _ = v50
                    }
                } else {
                    var v51 int64 = (v38 * int64(fi42))
                    _ = v51
                }
            default:
                for fi52 := 0; fi52 < 5; fi52++ {
                    var v53 uint16 = 775
                    _ = fi52
                    _ = v53
                }
            }
            _ = fi42
            _ = v43
            _ = v44
        }
        for fi54 := 0; fi54 < 11; fi54++ {
            switch fi54 {
            case 32768:
                fmt.Println(v39, v39)
            case -2147483648:
                for fi55 := 0; fi55 < 16; fi55++ {
                    switch "" {
                    case "foo bar":
                        switch "foo bar" {
                        case "foo bar":
                            if true {
                                var v56 float32 = float32((0.0 / 1e10))
                                var v57 float64 = 3.14159
                                var v58 uint = (v36 + v36)
                                v59 := GenS1{F0: (^(v39)), F1: ((("a" + "") + "a") + ""), F2: 127}
                                v56 += (3.14159 + 0.1)
                                _ = v56
                                _ = v57
                                _ = v58
                                _ = v59
                            }
                        default:
                            v35 -= ((0.1 - 1.5) - (float32(3.14159) + float32(0.1)))
                        }
                    case "hello":
                        delete(v1, fi54)
                    case "\n\t":
                        v60 := v1[fi54]
                        _ = v60
                    default:
                        switch (fi55 - (fi55 & fi55)) {
                        case 1:
                            v39 += v39
                        case -1:
                            v2 = v35
                        case 136:
                            v61 := GenS7{F0: (^(v38))}
                            _ = v61
                        }
                    }
                    if ((true || (v35 >= v35)) && ((!false) || false)) {
                        fmt.Println(v39, v39, v4)
                        var v62 float64 = 1e10
                        _ = v62
                    } else {
                        switch fi55 {
                        case -1:
                            switch fi55 {
                            case -701:
                                v36 = v36
                            case -1:
                                v3 = append(v3, fi55)
                            case 128:
                                v63 := GenS8{F0: 255, F1: 2, F2: v38}
                                _ = v63
                            default:
                                v64 := GenS1{F0: 255, F1: "\n\t", F2: uint8(v36)}
                                _ = v64
                            }
                        case 32768:
                            if (false && false) {
                                var v65 uint16 = 44
                                v2 = v35
                                v38 -= (v38 ^ int64(v65))
                                _ = v65
                            }
                        case 42:
                            if false {
                                v66 := GenS2{F0: (fi55 >> ((255) & 63))}
                                if false {
                                    v36 += (v36 + v36)
                                    if false {
                                        if true {
                                            if true {
                                                v39 -= 559
                                                var v67 int = fi55
                                                v3 = append(v3, (v67 ^ fi55))
                                                _ = v67
                                            }
                                            if true {
                                                if (0 > 0) {
                                                    var v68 int64 = (v38 | 0)
                                                    _ = v68
                                                } else {
                                                    fmt.Println(fi54, v38)
                                                    if (!(true && ((v36 & 2) < (v36 >> ((904) & 63))))) {
                                                        fmt.Println(v36)
                                                        var v69 int16 = 349
                                                        _ = v69
                                                    } else {
                                                        var v70 int8 = (v4 | (v4 | 1))
                                                        fmt.Println(v2)
                                                        _ = v70
                                                    }
                                                    v71 := GenS9{F0: (v38 - v38), F1: -1, F2: 9223372036854775807}
                                                    _ = v71
                                                }
                                            } else {
                                                var v72 int8 = 0
                                                if false {
                                                    v73 := GenS10{F0: 2}
                                                    _ = v73
                                                } else {
                                                    _ = len(v3)
                                                    if len(v3) > 0 {
                                                        v3[0] = int(v38)
                                                    }
                                                    fmt.Println(v35, v35, v39)
                                                }
                                                _ = v72
                                            }
                                            fmt.Println(v36, v39, fi55)
                                        }
                                        fmt.Println(fi55)
                                        var v74 uint8 = 42
                                        _ = v74
                                    } else {
                                        var v75 uint = 753
                                        v76 := v1[int(v4)]
                                        _ = v75
                                        _ = v76
                                    }
                                } else {
                                    if len(v3) > 0 {
                                        v3[0] = fi55
                                    }
                                    switch "" {
                                    case "hello":
                                        var v77 string = "a"
                                        _ = v77
                                    default:
                                        v78 := GenS10{F0: 255}
                                        _ = v78
                                    }
                                    var v79 bool = (false && (-1 >= 42))
                                    _ = v79
                                }
                                v39 = v39
                                if (42 < (v4 ^ (v4 >> ((18446744073709551615) & 63)))) {
                                    var v80 float32 = (((v2 * -2.25) + (float32(float64((v2 - v35))) * v35)) - v35)
                                    var v81 uint32 = 729
                                    _ = v80
                                    _ = v81
                                } else {
                                    v82 := GenS4{F0: (v39 ^ v39), F1: 0, F2: 867}
                                    var v83 uint16 = v39
                                    if ((v39 >> (((^(v36))) & 63)) <= (^(v39))) {
                                        v83 -= (v83 + v83)
                                        fmt.Println(v35)
                                        v1[(fi54 * (fi55 * fi55))] = "foo bar"
                                    } else {
                                        var v84 int64 = v38
                                        v85 := GenS11{F0: v36}
                                        switch (-(fi54)) {
                                        case -1:
                                            v86 := v1[(fi55 ^ (-(fi54)))]
                                            _ = v86
                                        }
                                        _ = v84
                                        _ = v85
                                    }
                                    _ = v82
                                    _ = v83
                                }
                                v36 = v36
                                _ = v66
                            }
                        default:
                            switch 32767 {
                            case 32767:
                                v3 = append(v3, fi54)
                            default:
                                var v87 uint16 = v39
                                _ = v87
                            }
                            v88 := GenS9{F0: (v38 / (((v38 >> ((455) & 63))) | 1)), F1: 1, F2: v38}
                            v36 += (v36 % ((61) | 1))
                            var v89 uint8 = 0
                            _ = v88
                            _ = v89
                        }
                    }
                    _ = fi55
                }
            case -764:
                for fi90 := 0; fi90 < 18; fi90++ {
                    var v91 int64 = (v38 ^ 61)
                    _ = fi90
                    _ = v91
                }
            }
            _ = fi54
        }
    }
    var v92 uint32 = 2
    for fi93 := 0; fi93 < 7; fi93++ {
        if (-2.25 < (3.14159 * 1e10)) {
            switch "hello" {
            case "hello":
                fmt.Println(v35)
            case "foo bar":
                if false {
                    if len(v3) > 0 {
                        v3[0] = -299
                    }
                }
            case "a":
                var v94 uint = (v36 | v36)
                _ = v94
            default:
                v92 = v92
            }
            var v95 uint32 = v92
            _ = v95
        }
        _ = fi93
    }
    var v96 uint = (^(v36))
    var v97 uint16 = 32767
    for fi98 := 0; fi98 < 15; fi98++ {
        _ = len(v3)
        v97 -= uint16(v92)
        var v99 uint = v96
        _ = fi98
        _ = v99
    }
    var v100 int16 = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v100:", v100)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v92:", v92)
    fmt.Println("v96:", v96)
    fmt.Println("v97:", v97)
    _ = v1
    _ = v100
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v92
    _ = v96
    _ = v97
}
