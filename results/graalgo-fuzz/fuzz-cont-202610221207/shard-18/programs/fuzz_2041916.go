package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint }
type GenS2 struct { F0 uint8; F1 uint64 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint64; F1 uint64 }
type GenS5 struct { F0 int32; F1 string; F2 int }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 int32; F1 uint; F2 uint32 }
func genFn0() int32 {
    var v4 uint8 = 127
    var v5 int8 = -46
    v5 -= (v5 & 1)
    _ = v4
    _ = v5
    return 1
}

func main() {
    var v1 int64 = 42
    v2 := make(map[int]uint32)
    var v3 uint = 499
    _ = len(v2)
    var v6 uint64 = 18446744073709551615
    var v7 uint8 = 42
    for fi8 := 0; fi8 < 5; fi8++ {
        v1 = v1
        _ = fi8
    }
    var v9 int16 = -32768
    v10 := GenS1{F0: 484, F1: (^(v3))}
    v7 += uint8(v6)
    fmt.Println(v7)
    v9 = v9
    if true {
        v11 := GenS1{F0: 2147483648, F1: v3}
        for fi12 := 0; fi12 < 7; fi12++ {
            var v13 int8 = -80
            fmt.Println(v3, v6)
            if true {
                v9 = (-(v9))
                for fi14 := 0; fi14 < 19; fi14++ {
                    v15 := []int{fi14, -145, fi12, (fi12 & (-(fi14)))}
                    var v16 int64 = (v1 << ((0) & 63))
                    _ = fi14
                    _ = v15
                    _ = v16
                }
                v17 := GenS1{F0: 578, F1: v3}
                _ = v17
            } else {
                switch ("foo bar" + "foo bar") {
                case "\n\t":
                    if true {
                        var v18 uint8 = 42
                        var v19 bool = false
                        if ((v3 | (^(v3))) != (v3 / (((v3 & 255)) | 1))) {
                            switch "foo bar" {
                            case "":
                                switch "hello" {
                                case "hello":
                                    for fi20 := 0; fi20 < 12; fi20++ {
                                        fmt.Println(v3, v19, fi20)
                                        if v19 {
                                            var v21 int = fi12
                                            _ = v21
                                        }
                                        if ((!(3.14159 <= -2.25)) || (true || (v7 == v18))) {
                                            switch fi12 {
                                            case -2147483648:
                                                v22 := GenS1{F0: 65535, F1: (v3 * 255)}
                                                _ = v22
                                            }
                                            var v23 int = fi12
                                            v24 := GenS2{F0: (^(v7)), F1: 42}
                                            _ = v23
                                            _ = v24
                                        }
                                        _ = fi20
                                    }
                                case "a":
                                    var v25 int64 = v1
                                    _ = v25
                                case "foo bar":
                                    var v26 uint32 = 791
                                    _ = v26
                                default:
                                    fmt.Println(fi12, v13)
                                }
                            }
                            if v19 {
                                v27 := GenS2{F0: (v18 >> ((2) & 63)), F1: (v6 ^ v6)}
                                var v28 string = "\n\t"
                                var v29 int8 = -1
                                _ = v27
                                _ = v28
                                _ = v29
                            } else {
                                var v30 int16 = v9
                                v13 += -128
                                var v31 string = ("a" + "a")
                                _ = v30
                                _ = v31
                            }
                            if (v13 < (-(v13))) {
                                v1 += v1
                                if true {
                                    v13 = (v13 | (v13 * (-(v13))))
                                }
                            } else {
                                v32 := GenS1{F0: 961, F1: v3}
                                _ = v32
                            }
                        }
                        _ = v18
                        _ = v19
                    }
                case "":
                    switch "foo bar" {
                    case "foo bar":
                        for fi33 := 0; fi33 < 1; fi33++ {
                            switch "a" {
                            case "":
                                var v34 float32 = (0.1 - (3.14159 + 0.0))
                                _ = v34
                            case "\n\t":
                                v35 := v2[fi33]
                                _ = v35
                            case "hello":
                                v9 = (v9 & (v9 | v9))
                            default:
                                var v36 uint32 = 1
                                fmt.Println(v36, v3, v7)
                                _ = v36
                            }
                            v37 := v2[128]
                            _ = fi33
                            _ = v37
                        }
                    case "hello":
                        switch fi12 {
                        case 97:
                            switch (fi12 << ((uint(v6)) & 63)) {
                            case 4294967295:
                                var v38 uint16 = 255
                                _ = v38
                            case 42:
                                v3 -= (v3 >> (((v3 - 65535)) & 63))
                            }
                        case 128:
                            switch (("foo bar" + "foo bar") + "") {
                            case "":
                                v39 := GenS2{F0: v7, F1: v6}
                                _ = v39
                            case "foo bar":
                                var v40 float32 = 1.5
                                _ = v40
                            case "a":
                                for fi41 := 0; fi41 < 1; fi41++ {
                                    switch "foo bar" {
                                    case "foo bar":
                                        v1 = (v1 >> ((31) & 63))
                                    case "":
                                        v1 = (v1 ^ v1)
                                    case "hello":
                                        v42 := GenS2{F0: v7, F1: (v6 | (v6 % (((v6 & 113)) | 1)))}
                                        _ = v42
                                    }
                                    _ = fi41
                                }
                            }
                        case 127:
                            v43 := v2[(fi12 * (fi12 << ((520) & 63)))]
                            _ = v43
                        }
                    default:
                        v1 = 486
                    }
                case "foo bar":
                    v2[fi12] = 65535
                }
            }
            _ = fi12
            _ = v13
        }
        _ = v11
    } else {
        v6 = (v6 | (v6 % (((v6 >> ((v3) & 63))) | 1)))
        for fi44 := 0; fi44 < 19; fi44++ {
            var v45 uint64 = 764
            var v46 uint16 = 42
            if (42 >= 0) {
                v47 := GenS3{F0: ("foo bar" + "hello")}
                if (false && (((0 >= 526) || ((v46 - v46) <= (v46 + 1))) && true)) {
                    for fi48 := 0; fi48 < 8; fi48++ {
                        v3 += 705
                        v6 = v45
                        fmt.Println(v6, fi44, v9)
                        _ = fi48
                    }
                }
                if false {
                    fmt.Println(v1, v9)
                    var v49 uint16 = v46
                    for fi50 := 0; fi50 < 14; fi50++ {
                        if (v9 <= 758) {
                            v51 := []int{}
                            v1 += (v1 + v1)
                            _ = v51
                        } else {
                            var v52 int8 = -111
                            if ((v1 << ((0) & 63)) < int64(v3)) {
                                v53 := GenS4{F0: v6, F1: v6}
                                _ = v53
                            }
                            _ = v52
                        }
                        _ = fi50
                    }
                    _ = v49
                } else {
                    if false {
                        for fi54 := 0; fi54 < 11; fi54++ {
                            if (2 != 78) {
                                if ((v46 | (v46 % ((v46) | 1))) >= v46) {
                                    v7 += 128
                                    v45 -= uint64(v3)
                                    v46 += 570
                                }
                                v55 := GenS5{F0: 2147483647, F1: ("foo bar" + "\n\t"), F2: (fi44 / ((int(v7)) | 1))}
                                v56 := GenS3{F0: ""}
                                _ = v55
                                _ = v56
                            }
                            v57 := GenS6{F0: ((^(v45)) != 4294967295)}
                            v46 += (v46 & v46)
                            _ = fi54
                            _ = v57
                        }
                        var v58 float64 = 1.5
                        _ = v58
                    }
                }
                _ = v47
            } else {
                if (!(true && false)) {
                    var v59 uint8 = (v7 & (^(v7)))
                    var v60 uint32 = 585
                    _ = v59
                    _ = v60
                }
                if true {
                    v61 := GenS5{F0: -1, F1: "\n\t", F2: (fi44 % (((fi44 ^ 32768)) | 1))}
                    var v62 uint16 = v46
                    var v63 int16 = (v9 * v9)
                    _ = v61
                    _ = v62
                    _ = v63
                }
                var v64 uint16 = (v46 / (((v46 + (v46 | 377))) | 1))
                _ = v64
            }
            _ = fi44
            _ = v45
            _ = v46
        }
    }
    var v65 int = 32768
    var v66 int = v65
    v67 := []int{-1, (v65 | v65), (v66 + 4294967295), -662}
    if (!(683 > 65535)) {
        fmt.Println(v7, v9)
        fmt.Println(v6, v1)
        for fi68 := 0; fi68 < 8; fi68++ {
            var v69 uint = 42
            _ = fi68
            _ = v69
        }
    }
    v65 = v66
    if false {
        v3 -= 4294967295
        if (!(!true)) {
            for fi70 := 0; fi70 < 5; fi70++ {
                var v71 uint = (v3 << (((v3 * (v3 & v3))) & 63))
                _ = fi70
                _ = v71
            }
            for fi72 := 0; fi72 < 12; fi72++ {
                var v73 float32 = (1.5 / -2.25)
                delete(v2, fi72)
                _ = fi72
                _ = v73
            }
            v74 := GenS5{F0: genFn0(), F1: ("" + "foo bar"), F2: (v66 / (((v66 << ((v3) & 63))) | 1))}
            _ = v74
        }
    } else {
        switch "hello" {
        case "hello":
            var v75 uint8 = (v7 ^ (v7 & (v7 + (v7 % (((v7 * 2)) | 1)))))
            _ = v75
        case "foo bar":
            var v76 float64 = float64(1.5)
            _ = v76
        }
        if true {
            v1 = (v1 >> ((v3) & 63))
        } else {
            v65 += v65
            for fi77 := 0; fi77 < 4; fi77++ {
                for fi78 := 0; fi78 < 12; fi78++ {
                    v66 -= (fi78 % ((fi78) | 1))
                    if len(v67) > 0 {
                        v67[0] = (v66 >> (((v3 ^ v3)) & 63))
                    }
                    var v79 int32 = int32(v9)
                    _ = fi78
                    _ = v79
                }
                _ = fi77
            }
        }
        for fi80 := 0; fi80 < 4; fi80++ {
            for fi81 := 0; fi81 < 1; fi81++ {
                var v82 int8 = 0
                _ = fi81
                _ = v82
            }
            v1 += 25
            v3 += (v3 - v3)
            _ = fi80
        }
    }
    v83 := GenS6{F0: (1e10 < 0.1)}
    v84 := GenS7{F0: -2147483648, F1: v3, F2: uint32(v6)}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("v65:", v65)
    fmt.Println("v66:", v66)
    fmt.Println("slicelen v67:", len(v67), v67)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v6
    _ = v65
    _ = v66
    _ = v67
    _ = v7
    _ = v83
    _ = v84
    _ = v9
}
