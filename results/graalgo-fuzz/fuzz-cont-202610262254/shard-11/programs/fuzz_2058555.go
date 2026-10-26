package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint64; F1 string; F2 uint64 }
type GenS3 struct { F0 string; F1 int16; F2 int16 }
type GenS4 struct { F0 uint; F1 uint8; F2 uint32 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 int8; F1 int16 }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 int32; F1 int16; F2 int16 }
type GenS10 struct { F0 int32; F1 int16 }
type GenS11 struct { F0 string; F1 bool; F2 int8 }
type GenS12 struct { F0 int8 }
type GenS13 struct { F0 int32 }
func genFn0(p0 uint8) int32 {
    for fi4 := 0; fi4 < 8; fi4++ {
        switch (fi4 * (fi4 ^ fi4)) {
        case 127:
            var v5 int8 = -128
            _ = v5
        case -2147483648:
            var v6 float32 = 1.5
            _ = v6
        default:
            v7 := GenS1{F0: -906}
            _ = v7
        }
        _ = fi4
    }
    p0 -= (p0 << ((42) & 63))
    var v8 uint8 = 42
    var v9 uint8 = (v8 - p0)
    _ = p0
    _ = v8
    _ = v9
    return 1
}
func genFn1(p0 uint32, p1 int32, p2 bool) int64 {
    switch 128 {
    case 1:
        switch 128 {
        case -2147483648:
            var v10 uint32 = p0
            _ = v10
        default:
            v11 := []int{}
            _ = v11
        }
    case 32767:
        var v12 uint8 = 255
        _ = v12
    default:
        fmt.Println(p2, p0, p2)
    }
    p0 += 4294967295
    if (!(100 >= 100)) {
        if false {
            p1 += (p1 + p1)
            if false {
                v13 := []int{-1}
                if (0 <= p1) {
                    if (!p2) {
                        if len(v13) > 0 {
                            v13[0] = 254
                        }
                        var v14 int32 = p1
                        _ = v14
                    }
                    v15 := GenS1{F0: genFn0(68)}
                    _ = v15
                } else {
                    for fi16 := 0; fi16 < 16; fi16++ {
                        v17 := GenS2{F0: 2, F1: "a", F2: 1}
                        p1 -= (p1 & (p1 * (p1 & p1)))
                        _ = fi16
                        _ = v17
                    }
                    var v18 int32 = p1
                    _ = v18
                }
                p0 = 42
                _ = v13
            } else {
                if p2 {
                    var v19 uint = 0
                    _ = v19
                } else {
                    fmt.Println(p1, p1, p1)
                }
                v20 := []int{}
                switch "a" {
                case "foo bar":
                    for fi21 := 0; fi21 < 5; fi21++ {
                        if p2 {
                            if ((0.1 + 1.5) != (-2.25 * -2.25)) {
                                var v22 uint32 = p0
                                _ = v22
                            } else {
                                var v23 string = "foo bar"
                                var v24 bool = true
                                if false {
                                    v25 := GenS3{F0: ("hello" + "\n\t"), F1: -1, F2: 0}
                                    _ = v25
                                }
                                _ = v23
                                _ = v24
                            }
                        }
                        p1 = (p1 >> ((635) & 63))
                        _ = fi21
                    }
                case "a":
                    var v26 uint32 = 65535
                    _ = v26
                default:
                    var v27 bool = p2
                    _ = v27
                }
                _ = v20
            }
        } else {
            fmt.Println(p1, p2)
            if true {
                p0 += 310
            }
        }
        v28 := GenS2{F0: 232, F1: "foo bar", F2: 4294967295}
        _ = v28
    } else {
        var v29 int64 = -1
        fmt.Println(p0, v29, p2)
        _ = v29
    }
    var v30 string = "\n\t"
    _ = p0
    _ = p1
    _ = p2
    _ = v30
    return -9223372036854775808
}

func main() {
    var v1 float64 = 3.14159
    var v2 int = 0
    v3 := make(map[int]uint32)
    v31 := []int{v2, v2, (v2 & 32768), v2}
    for fi32 := 0; fi32 < 15; fi32++ {
        switch fi32 {
        case 32767:
            fmt.Println(v1)
        }
        if (76 >= 255) {
            v2 = 555
            var v33 uint32 = 42
            v2 -= -1
            _ = v33
        } else {
            fmt.Println(v2, v1, v2)
            v34 := GenS4{F0: 42, F1: 1, F2: 65535}
            _ = v34
        }
        for fi35 := 0; fi35 < 9; fi35++ {
            switch "foo bar" {
            case "a":
                v1 = float64(0.1)
            }
            v36 := GenS5{F0: 46}
            v2 = (fi35 + -63)
            v2 = fi32
            switch (("" + ("hello" + "hello")) + "a") {
            case "":
                v37 := GenS6{F0: 100, F1: -287}
                if false {
                    fmt.Println(v1)
                }
                var v38 uint8 = 16
                var v39 uint8 = 0
                _ = v37
                _ = v38
                _ = v39
            case "hello":
                v2 -= fi32
            case "foo bar":
                if (!false) {
                    var v40 float64 = v1
                    _ = v40
                }
            }
            _ = fi35
            _ = v36
        }
        _ = fi32
    }
    v1 = (v1 / v1)
    var v41 int32 = 1
    var v42 uint8 = 243
    switch "foo bar" {
    case "a":
        for fi43 := 0; fi43 < 12; fi43++ {
            var v44 int8 = 42
            var v45 uint8 = (v42 << ((2) & 63))
            _ = fi43
            _ = v44
            _ = v45
        }
    case "\n\t":
        for fi46 := 0; fi46 < 11; fi46++ {
            v47 := GenS7{F0: int64(v42)}
            _ = fi46
            _ = v47
        }
    default:
        for fi48 := 0; fi48 < 7; fi48++ {
            for fi49 := 0; fi49 < 16; fi49++ {
                fmt.Println(v41)
                v50 := GenS5{F0: -608}
                _ = fi49
                _ = v50
            }
            if (!false) {
                if false {
                    if (-2.25 >= v1) {
                        v31 = append(v31, (v2 / ((v2) | 1)))
                    } else {
                        for fi51 := 0; fi51 < 19; fi51++ {
                            fmt.Println(fi48, fi48)
                            _ = fi51
                        }
                    }
                    var v52 int32 = (v41 % ((v41) | 1))
                    if (32767 <= 551) {
                        var v53 bool = true
                        if len(v31) >= 2 {
                            v31 = v31[1:len(v31)]
                        }
                        if (-127 > 0) {
                            var v54 uint32 = 2147483647
                            _ = v54
                        }
                        _ = v53
                    }
                    _ = v52
                }
            }
            _ = fi48
        }
    }
    if false {
        for fi55 := 0; fi55 < 10; fi55++ {
            for fi56 := 0; fi56 < 13; fi56++ {
                switch "\n\t" {
                case "a":
                    v57 := GenS2{F0: 18446744073709551615, F1: "", F2: 2}
                    _ = v57
                }
                _ = fi56
            }
            _ = fi55
        }
    } else {
        v1 = ((0.1 / (v1 + (v1 + v1))) + v1)
        if true {
            v58 := GenS1{F0: (v41 ^ (v41 - (v41 - v41)))}
            var v59 bool = (true || true)
            _ = v58
            _ = v59
        } else {
            if true {
                fmt.Println(v2, v2, v42)
                for fi60 := 0; fi60 < 14; fi60++ {
                    v1 = v1
                    v61 := GenS8{F0: 0}
                    for fi62 := 0; fi62 < 17; fi62++ {
                        switch "foo bar" {
                        case "":
                            if false {
                                v3[32767] = 611
                            }
                        default:
                            var v63 uint16 = 997
                            _ = v63
                        }
                        switch fi62 {
                        case 9223372036854775807:
                            fmt.Println(v41, fi60)
                        }
                        _ = fi62
                    }
                    _ = fi60
                    _ = v61
                }
                if true {
                    for fi64 := 0; fi64 < 13; fi64++ {
                        v42 = (v42 ^ 128)
                        switch fi64 {
                        case 0:
                            v65 := GenS7{F0: 1}
                            _ = v65
                        case -2147483648:
                            v66 := GenS9{F0: genFn0((v42 + v42)), F1: -741, F2: 32767}
                            _ = v66
                        }
                        _ = fi64
                    }
                } else {
                    var v67 float64 = v1
                    _ = v67
                }
            }
            v68 := GenS10{F0: genFn0(2), F1: -32768}
            _ = v68
        }
    }
    var v69 int8 = -128
    v70 := GenS11{F0: "a", F1: (false || (false || (!false))), F2: v69}
    var v71 float64 = v1
    v3[32767] = 2147483647
    v71 = v1
    if true {
        v2 -= -723
        if true {
            if false {
                var v72 int32 = genFn0(v42)
                _ = v72
            } else {
                if false {
                    for fi73 := 0; fi73 < 19; fi73++ {
                        v41 = 42
                        var v74 string = (((("\n\t" + ("\n\t" + ("foo bar" + "foo bar"))) + "a") + "foo bar") + "")
                        _ = fi73
                        _ = v74
                    }
                    if (65535 >= 558) {
                        if true {
                            fmt.Println(v1, v41, v1)
                        }
                    } else {
                        for fi75 := 0; fi75 < 8; fi75++ {
                            v71 = (v71 / v71)
                            if (true && true) {
                                var v76 bool = true
                                _ = v76
                            }
                            v77 := GenS12{F0: (v69 - (v69 + v69))}
                            _ = fi75
                            _ = v77
                        }
                        for fi78 := 0; fi78 < 11; fi78++ {
                            fmt.Println(v69, v1)
                            var v79 int16 = -32768
                            _ = fi78
                            _ = v79
                        }
                        var v80 uint = 2
                        _ = v80
                    }
                } else {
                    v42 = 2
                    if ((v2 + (v2 >> ((533) & 63))) == v2) {
                        for fi81 := 0; fi81 < 7; fi81++ {
                            if (140 < (v42 | (v42 >> ((4294967295) & 63)))) {
                                if false {
                                    var v82 int64 = 0
                                    v83 := GenS3{F0: ("a" + "hello"), F1: 5, F2: 32767}
                                    fmt.Println(v2, v82, v1)
                                    _ = v82
                                    _ = v83
                                } else {
                                    v84 := GenS13{F0: (v41 << ((2) & 63))}
                                    _ = v84
                                }
                            }
                            switch "" {
                            case "a":
                                if false {
                                    for fi85 := 0; fi85 < 4; fi85++ {
                                        v41 += (v41 - 1)
                                        _ = fi85
                                    }
                                    v69 = (v69 - (v69 >> ((728) & 63)))
                                    var v86 float64 = (v1 + 0.0)
                                    _ = v86
                                }
                            case "hello":
                                if (v42 < (v42 << ((42) & 63))) {
                                    for fi87 := 0; fi87 < 16; fi87++ {
                                        var v88 int = fi87
                                        v89 := v3[(-(v2))]
                                        _ = fi87
                                        _ = v88
                                        _ = v89
                                    }
                                }
                            default:
                                for fi90 := 0; fi90 < 5; fi90++ {
                                    _ = len(v31)
                                    var v91 float64 = v1
                                    _ = fi90
                                    _ = v91
                                }
                            }
                            _ = fi81
                        }
                    }
                }
                fmt.Println(v69, v2, v71)
            }
            v1 = v71
        }
        var v92 uint32 = 65535
        _ = v92
    } else {
        fmt.Println(v41, v42, v1)
        v42 += v42
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v69:", v69)
    fmt.Println("v71:", v71)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v41
    _ = v42
    _ = v69
    _ = v70
    _ = v71
}
