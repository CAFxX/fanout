package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16; F2 uint }
type GenS2 struct { F0 int64; F1 uint32 }
type GenS3 struct { F0 int; F1 uint8 }
type GenS4 struct { F0 uint64; F1 int16; F2 int16 }
type GenS5 struct { F0 uint; F1 int8; F2 int8 }
type GenS6 struct { F0 uint; F1 int64 }
type GenS7 struct { F0 uint; F1 bool; F2 string }
type GenS8 struct { F0 int32; F1 int; F2 string }
func genFn0(p0 uint32, p1 int, p2 int16) uint32 {
    var v5 float64 = 0.1
    for fi6 := 0; fi6 < 14; fi6++ {
        var v7 uint32 = (p0 << ((278) & 63))
        v7 = 0
        _ = fi6
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 2
}
func genFn1() uint8 {
    for fi8 := 0; fi8 < 7; fi8++ {
        if (!true) {
            v9 := GenS1{F0: 0, F1: 1, F2: 655}
            _ = v9
        } else {
            for fi10 := 0; fi10 < 5; fi10++ {
                if true {
                    var v11 uint8 = 100
                    switch "" {
                    case "":
                        if false {
                            var v12 float64 = (3.14159 / (-2.25 - 0.0))
                            var v13 uint8 = 255
                            fmt.Println(fi8, fi10, fi10)
                            _ = v12
                            _ = v13
                        }
                    case "hello":
                        if (!(1 > -1)) {
                            if true {
                                if true {
                                    if (-1 <= 127) {
                                        v11 = (v11 & (v11 * (v11 << ((18446744073709551615) & 63))))
                                        var v14 uint32 = 147
                                        _ = v14
                                    } else {
                                        var v15 float32 = 3.14159
                                        _ = v15
                                    }
                                    v16 := []int{-122, (-(fi10)), 141}
                                    if true {
                                        v17 := GenS1{F0: 299, F1: -32768, F2: 0}
                                        fmt.Println(fi10)
                                        if len(v16) > 0 {
                                            v16[0] = fi10
                                        }
                                        _ = v17
                                    }
                                    _ = v16
                                } else {
                                    var v18 float64 = 1e10
                                    _ = v18
                                }
                            } else {
                                fmt.Println(fi10, fi8, fi10)
                                switch "foo bar" {
                                case "a":
                                    var v19 uint = 42
                                    _ = v19
                                case "\n\t":
                                    var v20 int32 = 2147483647
                                    _ = v20
                                }
                                var v21 float64 = 1e10
                                _ = v21
                            }
                            v11 = 128
                            var v22 int32 = -107
                            _ = v22
                        } else {
                            if false {
                                v23 := GenS2{F0: -1, F1: 4294967295}
                                var v24 uint16 = 102
                                fmt.Println(fi8, v11)
                                _ = v23
                                _ = v24
                            }
                            v25 := []int{}
                            switch fi8 {
                            case -9223372036854775808:
                                var v26 uint64 = 4294967295
                                _ = v26
                            case -563:
                                var v27 float32 = -2.25
                                _ = v27
                            case 128:
                                if len(v25) > 0 {
                                    v25[0] = (fi8 % ((32768) | 1))
                                }
                            default:
                                var v28 int64 = 1
                                _ = v28
                            }
                            v29 := GenS2{F0: 0, F1: 0}
                            _ = v25
                            _ = v29
                        }
                    }
                    fmt.Println(fi10, fi8, fi10)
                    if (!true) {
                        v11 += 127
                    }
                    _ = v11
                }
                v30 := GenS1{F0: 890, F1: 1000, F2: 131}
                _ = fi10
                _ = v30
            }
            if false {
                v31 := []int{}
                _ = v31
            }
            for fi32 := 0; fi32 < 6; fi32++ {
                var v33 float32 = 1e10
                v33 -= v33
                _ = fi32
                _ = v33
            }
        }
        switch fi8 {
        case 2147483647:
            v34 := GenS2{F0: 42, F1: 0}
            _ = v34
        case 42:
            var v35 int64 = -9223372036854775808
            _ = v35
        }
        var v36 string = "foo bar"
        _ = fi8
        _ = v36
    }
    var v37 uint32 = 65535
    v37 += v37
    v37 -= genFn0(4294967295, 128, 42)
    switch 4294967295 {
    case 2147483647:
        switch "a" {
        case "foo bar":
            var v38 float64 = 1.5
            _ = v38
        case "":
            for fi39 := 0; fi39 < 12; fi39++ {
                for fi40 := 0; fi40 < 14; fi40++ {
                    if true {
                        if (100 < 71) {
                            v41 := []int{}
                            if (1 == 2) {
                                v37 += genFn0(v37, (fi40 + fi39), -871)
                            } else {
                                v42 := GenS2{F0: -1, F1: (v37 ^ v37)}
                                _ = v42
                            }
                            _ = v41
                        } else {
                            v43 := []int{fi40, fi40, (fi39 & (fi40 ^ fi40)), 9223372036854775807}
                            if (927 != 0) {
                                v44 := GenS3{F0: (fi39 ^ (^(fi39))), F1: 127}
                                var v45 int8 = 0
                                var v46 int = fi39
                                _ = v44
                                _ = v45
                                _ = v46
                            }
                            _ = v43
                        }
                        fmt.Println(fi40, v37, fi39)
                        if true {
                            v37 = v37
                            switch fi39 {
                            case 667:
                                v37 += (^(v37))
                            default:
                                if true {
                                    var v47 int16 = -221
                                    if true {
                                        fmt.Println(fi39, fi39, fi39)
                                        v48 := []int{}
                                        _ = v48
                                    } else {
                                        v49 := GenS1{F0: (v47 >> ((775) & 63)), F1: (v47 << ((109) & 63)), F2: 1}
                                        v50 := GenS4{F0: 0, F1: (v47 << ((525) & 63)), F2: (v47 % ((v47) | 1))}
                                        var v51 string = ("a" + ("hello" + ("a" + "foo bar")))
                                        _ = v49
                                        _ = v50
                                        _ = v51
                                    }
                                    _ = v47
                                }
                            }
                        } else {
                            switch "a" {
                            case "a":
                                fmt.Println(fi39, v37, fi39)
                            case "hello":
                                var v52 int64 = -9223372036854775808
                                _ = v52
                            case "":
                                var v53 int8 = -128
                                _ = v53
                            default:
                                var v54 int32 = 331
                                var v55 bool = false
                                var v56 float64 = (0.1 * (3.14159 - -2.25))
                                v57 := []int{fi40, 4294967295}
                                if (0 > -729) {
                                    var v58 uint64 = 42
                                    v55 = v55
                                    if len(v57) >= 2 {
                                        v57 = v57[1:len(v57)]
                                    }
                                    _ = v58
                                }
                                if v55 {
                                    v59 := GenS5{F0: 18446744073709551615, F1: 42, F2: 42}
                                    var v60 bool = ((true || (v55 || v55)) || true)
                                    var v61 float64 = v56
                                    _ = v59
                                    _ = v60
                                    _ = v61
                                }
                                switch int(v37) {
                                case 117:
                                    var v62 uint64 = 1
                                    _ = v62
                                default:
                                    var v63 int8 = 127
                                    _ = v63
                                }
                                _ = v54
                                _ = v55
                                _ = v56
                                _ = v57
                            }
                            var v64 float64 = -2.25
                            v65 := []int{fi39, 592, (fi39 * fi39)}
                            var v66 int64 = 42
                            _ = v64
                            _ = v65
                            _ = v66
                        }
                        fmt.Println(fi39)
                        v67 := []int{(fi40 << ((255) & 63)), 1}
                        _ = v67
                    }
                    if true {
                        var v68 uint16 = 2
                        var v69 uint64 = 326
                        _ = v68
                        _ = v69
                    }
                    _ = fi40
                }
                if false {
                    var v70 int32 = -1
                    var v71 int64 = 0
                    _ = v70
                    _ = v71
                } else {
                    for fi72 := 0; fi72 < 18; fi72++ {
                        var v73 bool = (genFn0(v37, (fi39 / ((fi39) | 1)), 1000) <= (v37 | v37))
                        _ = fi72
                        _ = v73
                    }
                    var v74 int = (fi39 + 122)
                    _ = v74
                }
                _ = fi39
            }
        case "hello":
            var v75 uint32 = v37
            _ = v75
        }
    case 353:
        switch 235 {
        case -1:
            var v76 string = "hello"
            _ = v76
        case 128:
            if (1 > 693) {
                v37 = v37
                v37 = (^(v37))
            }
        case -2147483648:
            var v77 string = ("foo bar" + ("" + "hello"))
            _ = v77
        default:
            for fi78 := 0; fi78 < 11; fi78++ {
                var v79 int64 = -1
                v80 := []int{fi78, fi78, (fi78 * (fi78 * fi78))}
                _ = fi78
                _ = v79
                _ = v80
            }
        }
    case 585:
        v81 := GenS6{F0: 2, F1: 0}
        _ = v81
    default:
        v37 += v37
    }
    _ = v37
    return 255
}

func main() {
    v1 := make(map[int]bool)
    v2 := []uint32{65535}
    var v3 uint64 = 844
    var v4 uint64 = 0
    v82 := GenS7{F0: 42, F1: false, F2: "hello"}
    v1[318] = (651 < 9223372036854775807)
    fmt.Println(v4, v4)
    var v83 bool = (127 == 305)
    var v84 int = 438
    var v85 uint32 = 4294967295
    var v86 uint = 255
    v4 = 18446744073709551615
    delete(v1, (^(v84)))
    v86 = (v86 | 517)
    if len(v2) > 0 {
        v2[0] = 2147483647
    }
    v83 = v83
    v87 := GenS8{F0: 42, F1: (v84 << ((v86) & 63)), F2: "foo bar"}
    v4 = v3
    var v88 bool = (-76 >= 0)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v83:", v83)
    fmt.Println("v84:", v84)
    fmt.Println("v85:", v85)
    fmt.Println("v86:", v86)
    fmt.Println("v88:", v88)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v82
    _ = v83
    _ = v84
    _ = v85
    _ = v86
    _ = v87
    _ = v88
}
