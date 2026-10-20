package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint16; F1 int; F2 bool }
type GenS3 struct { F0 string; F1 int; F2 int }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int; F1 uint8; F2 uint64 }
type GenS6 struct { F0 int32; F1 uint8; F2 uint8 }
func genFn0(p0 uint, p1 uint8, p2 uint16) bool {
    p0 -= (p0 + (^(p0)))
    p0 += p0
    for fi4 := 0; fi4 < 7; fi4++ {
        v5 := GenS1{F0: 1000}
        for fi6 := 0; fi6 < 17; fi6++ {
            if true {
                v7 := GenS1{F0: 42}
                var v8 int8 = int8(p1)
                p0 -= (p0 >> ((42) & 63))
                _ = v7
                _ = v8
            }
            p2 += (^(p2))
            _ = fi6
        }
        _ = fi4
        _ = v5
    }
    v9 := GenS1{F0: 655}
    if false {
        v10 := []int{-714}
        fmt.Println(p0, p0, p1)
        _ = v10
    } else {
        switch "\n\t" {
        case "hello":
            var v11 uint32 = 1
            _ = v11
        case "":
            p2 = (^(p2))
        case "\n\t":
            if (!true) {
                for fi12 := 0; fi12 < 1; fi12++ {
                    var v13 int64 = 9223372036854775807
                    _ = fi12
                    _ = v13
                }
                v14 := []int{9223372036854775807, -2147483648}
                v15 := GenS1{F0: 1}
                _ = v14
                _ = v15
            }
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return true
}

func main() {
    var v1 string = "hello"
    v2 := make(map[uint32]string)
    var v3 int32 = -903
    v16 := GenS1{F0: 42}
    var v17 uint32 = 2
    var v18 bool = (0.0 < (0.0 * 1.5))
    var v19 int32 = int32(v17)
    v20 := GenS1{F0: 0}
    var v21 uint = 42
    var v22 int16 = 42
    var v23 bool = v18
    if (!false) {
        for fi24 := 0; fi24 < 20; fi24++ {
            v25 := GenS2{F0: 32767, F1: (fi24 ^ int(v21)), F2: (920 <= fi24)}
            v19 = (v3 ^ (v3 | (^(v3))))
            var v26 int8 = 6
            _ = fi24
            _ = v25
            _ = v26
        }
        var v27 int32 = (v3 & (v3 + -1))
        _ = v27
    }
    var v28 bool = (v18 || (582 <= 805))
    v18 = v28
    for fi29 := 0; fi29 < 3; fi29++ {
        if ((108 > 100) && false) {
            var v30 int16 = 1000
            var v31 uint = v21
            if v28 {
                v21 -= (v21 | (v21 * v21))
                v32 := GenS1{F0: v22}
                _ = v32
            }
            _ = v30
            _ = v31
        } else {
            switch "a" {
            case "\n\t":
                switch "a" {
                case "\n\t":
                    switch fi29 {
                    case 32767:
                        for fi33 := 0; fi33 < 12; fi33++ {
                            switch v1 {
                            case "":
                                switch (fi33 << ((v21) & 63)) {
                                case 64:
                                    var v34 uint32 = v17
                                    _ = v34
                                default:
                                    v35 := v2[(v17 * v17)]
                                    _ = v35
                                }
                            }
                            if v18 {
                                fmt.Println(v18, v18)
                            } else {
                                var v36 bool = v28
                                if true {
                                    var v37 int = fi29
                                    var v38 uint = 2
                                    v39 := GenS1{F0: (-(v22))}
                                    _ = v37
                                    _ = v38
                                    _ = v39
                                } else {
                                    v40 := GenS2{F0: 640, F1: (fi33 % ((fi29) | 1)), F2: v18}
                                    switch "a" {
                                    case "":
                                        switch (("a" + v1) + (("a" + v1) + ("\n\t" + v1))) {
                                        case "\n\t":
                                            if (v18 || v36) {
                                                var v41 int = 9223372036854775807
                                                _ = v41
                                            }
                                        }
                                    }
                                    _ = v40
                                }
                                _ = v36
                            }
                            var v42 int8 = 42
                            v42 += (v42 & v42)
                            var v43 uint64 = 4294967295
                            _ = fi33
                            _ = v42
                            _ = v43
                        }
                    case -617:
                        v2[v17] = v1
                    }
                }
            case "":
                if v18 {
                    for fi44 := 0; fi44 < 20; fi44++ {
                        fmt.Println(v21)
                        v19 = (-(v19))
                        _ = fi44
                    }
                }
            default:
                if false {
                    for fi45 := 0; fi45 < 2; fi45++ {
                        v46 := GenS1{F0: (-(v22))}
                        _ = fi45
                        _ = v46
                    }
                    if (427 < (fi29 >> ((v21) & 63))) {
                        switch v1 {
                        case "a":
                            switch "foo bar" {
                            case "\n\t":
                                var v47 bool = false
                                _ = v47
                            case "":
                                switch fi29 {
                                case 128:
                                    _ = len(v2)
                                case -860:
                                    if (((v1 + v1) + "a") < (v1 + v1)) {
                                        v48 := GenS3{F0: v1, F1: fi29, F2: (fi29 << ((537) & 63))}
                                        v49 := GenS2{F0: 65535, F1: fi29, F2: (uint32(v21) <= v17)}
                                        _ = v48
                                        _ = v49
                                    }
                                case 570:
                                    v50 := GenS3{F0: ("a" + "\n\t"), F1: (fi29 >> ((v21) & 63)), F2: (fi29 + 32767)}
                                    _ = v50
                                default:
                                    v17 += 2
                                }
                            case "a":
                                var v51 uint8 = 2
                                _ = v51
                            default:
                                for fi52 := 0; fi52 < 8; fi52++ {
                                    var v53 int = (fi52 - (fi29 / (((fi29 % ((fi52) | 1))) | 1)))
                                    var v54 int = (fi52 >> ((v21) & 63))
                                    v54 = 1
                                    _ = fi52
                                    _ = v53
                                    _ = v54
                                }
                            }
                        }
                        for fi55 := 0; fi55 < 7; fi55++ {
                            if ((981 != 9223372036854775807) && v18) {
                                switch fi29 {
                                case 128:
                                    v56 := GenS3{F0: v1, F1: fi29, F2: (fi55 * (-(fi29)))}
                                    v57 := []int{42}
                                    _ = v56
                                    _ = v57
                                case 527:
                                    if false {
                                        if (!false) {
                                            var v58 uint = v21
                                            v59 := GenS1{F0: 737}
                                            var v60 float32 = ((float32(0.0) / (-2.25 - 1.5)) - 1.5)
                                            fmt.Println(v21, fi55)
                                            _ = v58
                                            _ = v59
                                            _ = v60
                                        } else {
                                            v21 += 0
                                            var v61 int8 = -46
                                            _ = v61
                                        }
                                    } else {
                                        fmt.Println(v28)
                                    }
                                case 0:
                                    v62 := []int{}
                                    _ = v62
                                default:
                                    switch "a" {
                                    case "foo bar":
                                        v2[(v17 << ((v21) & 63))] = v1
                                    }
                                }
                                var v63 uint = (v21 * v21)
                                _ = v63
                            }
                            v64 := []int{(fi29 | fi55)}
                            _ = fi55
                            _ = v64
                        }
                    }
                    v65 := []int{(fi29 - fi29), (^(fi29)), (-(fi29))}
                    _ = v65
                } else {
                    fmt.Println(v28)
                    switch v1 {
                    case "":
                        var v66 float32 = 0.1
                        _ = v66
                    case "\n\t":
                        if v23 {
                            fmt.Println(v21, v28, v22)
                        }
                    case "a":
                        v67 := GenS4{F0: 130}
                        _ = v67
                    }
                }
            }
        }
        _ = fi29
    }
    for fi68 := 0; fi68 < 10; fi68++ {
        v69 := []int{(^(fi68)), (fi68 % (((fi68 / (((fi68 - (fi68 + 1))) | 1))) | 1)), 127}
        var v70 uint16 = 32767
        _ = fi68
        _ = v69
        _ = v70
    }
    if (1 == 43) {
        v17 += (v17 | v17)
        var v71 string = ("a" + v1)
        _ = v71
    } else {
        if v28 {
            for fi72 := 0; fi72 < 10; fi72++ {
                v73 := GenS4{F0: 1}
                _ = fi72
                _ = v73
            }
        } else {
            var v74 uint32 = (v17 & v17)
            v19 -= v3
            if (!v23) {
                var v75 uint8 = 127
                for fi76 := 0; fi76 < 11; fi76++ {
                    fmt.Println(fi76, v22, v19)
                    var v77 int = fi76
                    _ = fi76
                    _ = v77
                }
                switch "foo bar" {
                case "hello":
                    for fi78 := 0; fi78 < 13; fi78++ {
                        fmt.Println(v22)
                        switch (v1 + "a") {
                        case "hello":
                            var v79 string = v1
                            _ = v79
                        case "":
                            v80 := []int{int(v19)}
                            _ = v80
                        default:
                            var v81 uint = uint(v17)
                            _ = v81
                        }
                        for fi82 := 0; fi82 < 7; fi82++ {
                            _ = len(v2)
                            switch fi78 {
                            case 127:
                                if (!v28) {
                                    _ = len(v2)
                                    switch "\n\t" {
                                    case "hello":
                                        fmt.Println(fi78, v23, v19)
                                    default:
                                        fmt.Println(v21, v74)
                                    }
                                    v28 = true
                                    var v83 bool = (((v1 + v1) == "hello") || (32768 <= (-(fi82))))
                                    v84 := []int{0, fi78}
                                    _ = v83
                                    _ = v84
                                } else {
                                    if (0 <= 0) {
                                        var v85 int8 = -99
                                        switch fi78 {
                                        case 286:
                                            fmt.Println(v17, v3)
                                        default:
                                            var v86 int = fi82
                                            _ = v86
                                        }
                                        _ = v85
                                    }
                                }
                            case -2147483648:
                                v87 := GenS5{F0: (fi78 & -705), F1: 7, F2: 644}
                                _ = v87
                            case 42:
                                _ = len(v2)
                            }
                            _ = fi82
                        }
                        _ = fi78
                    }
                case "\n\t":
                    v21 = (v21 ^ (v21 >> (((v21 & v21)) & 63)))
                }
                _ = v75
            } else {
                var v88 float64 = 3.14159
                _ = v88
            }
            _ = v74
        }
        for fi89 := 0; fi89 < 13; fi89++ {
            v90 := v2[v17]
            v91 := GenS6{F0: v3, F1: 1, F2: 128}
            fmt.Println(v18)
            _ = fi89
            _ = v90
            _ = v91
        }
        switch v1 {
        case "hello":
            fmt.Println(v19, v17, v21)
        case "foo bar":
            var v92 int16 = 377
            _ = v92
        default:
            v93 := []int{32768, -2147483648}
            _ = v93
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v28
    _ = v3
}
