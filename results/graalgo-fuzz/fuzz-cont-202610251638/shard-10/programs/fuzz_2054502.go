package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int }
type GenS2 struct { F0 int; F1 uint16; F2 int8 }
type GenS3 struct { F0 uint8; F1 uint; F2 uint64 }
type GenS4 struct { F0 uint16; F1 uint8; F2 int8 }
type GenS5 struct { F0 uint8; F1 uint8; F2 bool }
type GenS6 struct { F0 uint8; F1 bool; F2 string }
type GenS7 struct { F0 int16; F1 uint32; F2 uint64 }
func genFn0(p0 uint16, p1 uint32, p2 uint8) uint32 {
    for fi6 := 0; fi6 < 8; fi6++ {
        fmt.Println(p1, p0, p1)
        p1 -= p1
        _ = fi6
    }
    var v7 int32 = 2147483647
    p0 = p0
    v8 := []int{128, -9223372036854775808, 654, 438}
    var v9 uint32 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return v9
}
func genFn1(p0 string, p1 uint) uint {
    for fi10 := 0; fi10 < 14; fi10++ {
        var v11 uint8 = 148
        _ = fi10
        _ = v11
    }
    p0 = "a"
    fmt.Println(p0)
    _ = p0
    _ = p1
    return 0
}

func main() {
    v1 := make(map[int]int)
    var v2 float64 = ((-2.25 * 1.5) * 1.5)
    var v3 uint16 = 1
    var v4 uint64 = 42
    var v5 int8 = -1
    if (false && false) {
        var v12 int = -279
        switch "" {
        case "a":
            var v13 int64 = 42
            _ = v13
        case "":
            for fi14 := 0; fi14 < 12; fi14++ {
                var v15 uint32 = 970
                _ = fi14
                _ = v15
            }
        default:
            for fi16 := 0; fi16 < 7; fi16++ {
                switch -309 {
                case 128:
                    fmt.Println(v12)
                case 1:
                    switch (-(fi16)) {
                    case 32768:
                        v17 := GenS1{F0: (v4 ^ v4), F1: fi16}
                        _ = v17
                    case 4294967295:
                        if true {
                            if true {
                                var v18 uint32 = 4294967295
                                v2 = (v2 - (v2 + v2))
                                v12 -= v12
                                _ = v18
                            } else {
                                v4 = v4
                                for fi19 := 0; fi19 < 8; fi19++ {
                                    if true {
                                        _ = len(v1)
                                        if (int64(v3) > 872) {
                                            var v20 uint16 = 32767
                                            var v21 int16 = -1
                                            _ = v20
                                            _ = v21
                                        }
                                        switch "\n\t" {
                                        case "":
                                            fmt.Println(v4, fi16)
                                            var v22 uint64 = 541
                                            switch "hello" {
                                            case "hello":
                                                v3 = 831
                                            case "\n\t":
                                                v4 = v4
                                            case "foo bar":
                                                v5 += 100
                                            default:
                                                var v23 int = (fi16 - (-(fi16)))
                                                _ = v23
                                            }
                                            _ = v22
                                        case "foo bar":
                                            var v24 uint32 = 2
                                            _ = v24
                                        case "a":
                                            if (true && ((false && true) && false)) {
                                                var v25 int64 = 1
                                                var v26 uint = genFn1("foo bar", 42)
                                                switch "hello" {
                                                case "a":
                                                    var v27 float32 = 3.14159
                                                    _ = v27
                                                case "\n\t":
                                                    var v28 int16 = -27
                                                    _ = v28
                                                }
                                                _ = v25
                                                _ = v26
                                            }
                                        }
                                    } else {
                                        var v29 float64 = (v2 + v2)
                                        var v30 int = 2147483647
                                        fmt.Println(fi16, v4, v5)
                                        _ = v29
                                        _ = v30
                                    }
                                    v2 = (v2 * -2.25)
                                    _ = fi19
                                }
                            }
                            fmt.Println(fi16, v3, v5)
                        } else {
                            delete(v1, (v12 >> ((65535) & 63)))
                            if true {
                                switch ((("hello" + "") + "a") + "a") {
                                case "a":
                                    for fi31 := 0; fi31 < 9; fi31++ {
                                        v12 = (fi16 / ((fi16) | 1))
                                        v32 := []int{}
                                        if false {
                                            v5 = (v5 % ((v5) | 1))
                                        }
                                        if len(v32) >= 2 {
                                            v32 = v32[1:len(v32)]
                                        }
                                        _ = fi31
                                        _ = v32
                                    }
                                }
                            } else {
                                var v33 uint64 = (^(v4))
                                _ = v33
                            }
                            v34 := GenS2{F0: fi16, F1: 63, F2: 0}
                            _ = v34
                        }
                    default:
                        var v35 int32 = 0
                        _ = v35
                    }
                case 701:
                    v36 := []int{(fi16 + v12)}
                    _ = v36
                }
                if false {
                    v1[fi16] = 428
                    for fi37 := 0; fi37 < 19; fi37++ {
                        fmt.Println(v4, fi16)
                        var v38 uint16 = 2
                        if false {
                            switch v12 {
                            case 128:
                                v39 := GenS3{F0: 1, F1: 238, F2: (^(v4))}
                                _ = v39
                            case 399:
                                v4 = (v4 / ((v4) | 1))
                            }
                            switch ("a" + "hello") {
                            case "":
                                switch 984 {
                                case 795:
                                    fmt.Println(v38, v5, fi37)
                                case 32767:
                                    var v40 int32 = -1
                                    _ = v40
                                case 2147483647:
                                    if ("" <= "") {
                                        var v41 uint32 = 2
                                        var v42 uint16 = v3
                                        var v43 float64 = v2
                                        _ = v41
                                        _ = v42
                                        _ = v43
                                    }
                                default:
                                    if ((1e10 * v2) > v2) {
                                        v5 = v5
                                        fmt.Println(v2, v12)
                                    } else {
                                        var v44 uint = genFn1("\n\t", 2)
                                        _ = v44
                                    }
                                    var v45 uint8 = 127
                                    _ = v45
                                }
                                v2 = float64(1.5)
                                var v46 uint32 = 761
                                _ = v46
                            case "\n\t":
                                var v47 string = "hello"
                                _ = v47
                            case "hello":
                                var v48 int8 = 127
                                _ = v48
                            }
                        } else {
                            fmt.Println(v3)
                        }
                        _ = fi37
                        _ = v38
                    }
                    fmt.Println(v4, fi16, v4)
                } else {
                    for fi49 := 0; fi49 < 8; fi49++ {
                        if (1.5 == -2.25) {
                            switch 146 {
                            case -2147483648:
                                fmt.Println(v5)
                            case 9223372036854775807:
                                _ = len(v1)
                            case 32768:
                                v50 := GenS4{F0: v3, F1: 7, F2: (^(v5))}
                                _ = v50
                            default:
                                switch v12 {
                                case -1:
                                    if false {
                                        v51 := GenS2{F0: (-(fi49)), F1: (v3 + 2), F2: v5}
                                        _ = v51
                                    } else {
                                        v52 := []int{32767, (fi16 >> ((42) & 63)), -2147483648}
                                        var v53 uint16 = 5
                                        v54 := GenS5{F0: 157, F1: 2, F2: (true && true)}
                                        _ = v52
                                        _ = v53
                                        _ = v54
                                    }
                                    fmt.Println(v2, v4, fi49)
                                }
                            }
                        } else {
                            v55 := GenS5{F0: 2, F1: 0, F2: false}
                            fmt.Println(v4)
                            switch "hello" {
                            case "":
                                switch "hello" {
                                case "foo bar":
                                    v56 := GenS3{F0: 127, F1: 2, F2: v4}
                                    _ = v56
                                case "\n\t":
                                    v4 += 42
                                case "a":
                                    if false {
                                        fmt.Println(v12, v3, v3)
                                        v57 := GenS3{F0: 255, F1: 42, F2: (v4 << ((1) & 63))}
                                        if true {
                                            v58 := []int{v12}
                                            var v59 int32 = -2147483648
                                            _ = v58
                                            _ = v59
                                        }
                                        _ = v57
                                    } else {
                                        v2 = 0.0
                                        fmt.Println(v5, v3, v2)
                                    }
                                default:
                                    switch "hello" {
                                    case "\n\t":
                                        var v60 string = "foo bar"
                                        _ = v60
                                    }
                                }
                            default:
                                if false {
                                    if true {
                                        var v61 int32 = 2147483647
                                        v4 = (v4 + v4)
                                        _ = v61
                                    } else {
                                        var v62 uint16 = v3
                                        v4 -= 2
                                        var v63 float64 = (((v2 + (v2 / v2)) - v2) * v2)
                                        _ = v62
                                        _ = v63
                                    }
                                }
                            }
                            v64 := []int{fi16, (fi16 + int(v3)), (-(fi49))}
                            v64 = append(v64, (fi49 + -505))
                            var v65 int64 = -484
                            _ = v55
                            _ = v64
                            _ = v65
                        }
                        v5 += v5
                        v5 = 100
                        _ = fi49
                    }
                    v66 := []int{}
                    _ = v66
                }
                if false {
                    var v67 uint32 = genFn0((v3 ^ (v3 >> ((18) & 63))), genFn0(535, 42, 0), 1)
                    _ = v67
                }
                _ = fi16
            }
        }
        fmt.Println(v5, v5)
        _ = v12
    } else {
        for fi68 := 0; fi68 < 16; fi68++ {
            v5 = (v5 - (v5 >> ((genFn1("\n\t", 42)) & 63)))
            v5 += (-(v5))
            switch "a" {
            case "":
                v69 := GenS6{F0: 127, F1: true, F2: "hello"}
                _ = v69
            default:
                v70 := GenS6{F0: 1, F1: true, F2: "a"}
                _ = v70
            }
            _ = fi68
        }
        for fi71 := 0; fi71 < 6; fi71++ {
            var v72 uint16 = v3
            v4 -= v4
            v4 = (v4 - 4294967295)
            _ = fi71
            _ = v72
        }
    }
    for fi73 := 0; fi73 < 16; fi73++ {
        v74 := v1[(fi73 >> ((643) & 63))]
        var v75 int8 = 110
        _ = fi73
        _ = v74
        _ = v75
    }
    if true {
        v76 := GenS4{F0: v3, F1: 1, F2: (v5 % ((-128) | 1))}
        _ = v76
    }
    fmt.Println(v4, v3, v2)
    v77 := []int{}
    v2 += 0.0
    fmt.Println(v2, v5, v2)
    for fi78 := 0; fi78 < 13; fi78++ {
        fmt.Println(v2, v2)
        var v79 int = (-(fi78))
        _ = fi78
        _ = v79
    }
    v4 -= v4
    var v80 float64 = float64(0.0)
    _ = len(v1)
    if false {
        v77 = append(v77, 0)
        if (!true) {
            v81 := v1[-415]
            var v82 int32 = 2147483647
            v81 = (v81 ^ v81)
            _ = v81
            _ = v82
        }
        v4 += (v4 * v4)
    }
    v3 += (v3 ^ (v3 >> ((279) & 63)))
    var v83 uint8 = 74
    for fi84 := 0; fi84 < 5; fi84++ {
        var v85 uint8 = 1
        switch "a" {
        case "\n\t":
            v86 := GenS2{F0: fi84, F1: (v3 * v3), F2: 1}
            _ = v86
        case "a":
            fmt.Println(v80, v85)
        case "":
            var v87 uint32 = 4294967295
            _ = v87
        default:
            if len(v77) > 0 {
                v77[0] = fi84
            }
        }
        _ = fi84
        _ = v85
    }
    switch 788 {
    case 113:
        v80 += 0.0
    default:
        switch ("" + "a") {
        case "":
            for fi88 := 0; fi88 < 9; fi88++ {
                v89 := GenS3{F0: 0, F1: 42, F2: (v4 << ((255) & 63))}
                _ = fi88
                _ = v89
            }
        case "\n\t":
            var v90 uint32 = 555
            _ = v90
        case "a":
            _ = len(v1)
        default:
            fmt.Println(v80)
        }
    }
    v91 := GenS4{F0: v3, F1: v83, F2: (v5 % ((100) | 1))}
    v92 := GenS7{F0: 42, F1: 2, F2: 4294967295}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v77:", len(v77), v77)
    fmt.Println("v80:", v80)
    fmt.Println("v83:", v83)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v77
    _ = v80
    _ = v83
    _ = v91
    _ = v92
}
