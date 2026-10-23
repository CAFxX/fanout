package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16; F2 uint16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint32; F1 uint8 }
type GenS4 struct { F0 uint16; F1 int8 }
type GenS5 struct { F0 bool; F1 int16; F2 bool }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 int8; F1 int32; F2 int32 }
type GenS8 struct { F0 uint8; F1 uint32 }
func genFn0(p0 uint16, p1 bool) uint32 {
    for fi4 := 0; fi4 < 13; fi4++ {
        p1 = p1
        v5 := GenS1{F0: 255, F1: p0, F2: p0}
        _ = fi4
        _ = v5
    }
    for fi6 := 0; fi6 < 19; fi6++ {
        p1 = p1
        _ = fi6
    }
    v7 := []int{32767, 9223372036854775807, 1000, -251}
    fmt.Println(p0, p1)
    _ = p0
    _ = p1
    _ = v7
    return 1
}

func main() {
    var v1 uint8 = 255
    var v2 int64 = 42
    var v3 uint32 = 2147483648
    v8 := GenS2{F0: ("a" + ("hello" + "\n\t"))}
    v9 := []int{-9223372036854775808}
    if false {
        if false {
            if len(v9) > 0 {
                v9[0] = 482
            }
            _ = len(v9)
        } else {
            for fi10 := 0; fi10 < 2; fi10++ {
                var v11 uint16 = 47
                if len(v9) > 0 {
                    v9[0] = fi10
                }
                for fi12 := 0; fi12 < 6; fi12++ {
                    var v13 int32 = 2147483647
                    if true {
                        v14 := GenS2{F0: ("foo bar" + "foo bar")}
                        _ = len(v9)
                        v15 := GenS3{F0: 42, F1: (v1 ^ 31)}
                        var v16 bool = true
                        var v17 uint32 = v3
                        _ = v14
                        _ = v15
                        _ = v16
                        _ = v17
                    } else {
                        var v18 bool = (!true)
                        v1 = v1
                        var v19 uint = 18446744073709551615
                        var v20 uint64 = 241
                        v19 = v19
                        _ = v18
                        _ = v19
                        _ = v20
                    }
                    var v21 uint = 1
                    v22 := GenS4{F0: (v11 | v11), F1: 0}
                    if false {
                        v9 = append(v9, int(v2))
                        v23 := GenS3{F0: (v3 ^ 2147483648), F1: (^(v1))}
                        _ = v23
                    } else {
                        var v24 float32 = (0.0 - (-2.25 + (1.5 + 1.5)))
                        switch fi12 {
                        case 128:
                            _ = len(v9)
                        case -1:
                            var v25 int32 = v13
                            _ = v25
                        default:
                            if ("hello" < "") {
                                if len(v9) >= 2 {
                                    v9 = v9[1:len(v9)]
                                }
                                var v26 uint8 = v1
                                var v27 int8 = 122
                                _ = v26
                                _ = v27
                            } else {
                                v28 := GenS4{F0: v11, F1: 97}
                                var v29 int32 = v13
                                var v30 int64 = (v2 * v2)
                                _ = v28
                                _ = v29
                                _ = v30
                            }
                        }
                        switch (-(fi10)) {
                        case 32767:
                            var v31 float32 = (0.1 * v24)
                            _ = v31
                        case 2147483647:
                            var v32 string = (("hello" + "\n\t") + "")
                            _ = v32
                        default:
                            _ = len(v9)
                            v33 := GenS1{F0: v1, F1: v11, F2: 255}
                            _ = v33
                        }
                        _ = v24
                    }
                    _ = fi12
                    _ = v13
                    _ = v21
                    _ = v22
                }
                _ = fi10
                _ = v11
            }
        }
        if false {
            v3 = 2
            fmt.Println(v1, v1, v1)
        }
    }
    switch 940 {
    case -295:
        if false {
            v9 = append(v9, 803)
            v1 -= (v1 >> ((65535) & 63))
            v1 = v1
        }
    case 127:
        v2 = v2
    default:
        v34 := GenS3{F0: (v3 % ((v3) | 1)), F1: 2}
        _ = v34
    }
    var v35 bool = true
    v3 = (v3 - 2)
    if (0 > 2) {
        var v36 bool = (65535 >= 2)
        v35 = (0 >= 127)
        _ = v36
    }
    switch ("" + "\n\t") {
    case "":
        var v37 float64 = ((1e10 / (float64(1.5) + (1.5 / (1e10 - 0.0)))) + 0.0)
        _ = v37
    default:
        if (2 == 1) {
            if v35 {
                if (4294967295 <= 542) {
                    v9 = append(v9, 127)
                } else {
                    fmt.Println(v1)
                    if v35 {
                        switch "a" {
                        case "\n\t":
                            var v38 uint16 = 32767
                            _ = v38
                        }
                    } else {
                        v1 = (^(v1))
                        switch 42 {
                        case -2147483648:
                            if len(v9) >= 2 {
                                v9 = v9[1:len(v9)]
                            }
                        case 127:
                            if v35 {
                                if len(v9) > 0 {
                                    v9[0] = 128
                                }
                                if true {
                                    if (-1 <= 0) {
                                        for fi39 := 0; fi39 < 9; fi39++ {
                                            if v35 {
                                                switch int(v2) {
                                                case -349:
                                                    fmt.Println(v35, v3)
                                                default:
                                                    for fi40 := 0; fi40 < 17; fi40++ {
                                                        v41 := GenS2{F0: ("hello" + ("a" + "\n\t"))}
                                                        _ = fi40
                                                        _ = v41
                                                    }
                                                }
                                            } else {
                                                fmt.Println(fi39, fi39)
                                            }
                                            _ = fi39
                                        }
                                        if v35 {
                                            fmt.Println(v3, v3)
                                            v42 := GenS3{F0: genFn0(836, true), F1: (v1 + v1)}
                                            _ = v42
                                        }
                                    } else {
                                        for fi43 := 0; fi43 < 5; fi43++ {
                                            for fi44 := 0; fi44 < 11; fi44++ {
                                                fmt.Println(v1)
                                                if false {
                                                    v45 := GenS2{F0: ("\n\t" + "")}
                                                    v2 = (^(v2))
                                                    v2 = (v2 % (((-(v2))) | 1))
                                                    _ = v45
                                                } else {
                                                    if v35 {
                                                        v3 = (v3 ^ v3)
                                                        _ = len(v9)
                                                        var v46 uint8 = 26
                                                        _ = v46
                                                    } else {
                                                        v2 += (v2 / ((v2) | 1))
                                                    }
                                                    fmt.Println(v3, fi43)
                                                    if len(v9) >= 2 {
                                                        v9 = v9[1:len(v9)]
                                                    }
                                                }
                                                _ = fi44
                                            }
                                            v3 -= v3
                                            _ = fi43
                                        }
                                        for fi47 := 0; fi47 < 3; fi47++ {
                                            for fi48 := 0; fi48 < 16; fi48++ {
                                                if len(v9) >= 2 {
                                                    v9 = v9[1:len(v9)]
                                                }
                                                _ = fi48
                                            }
                                            var v49 int = (fi47 << ((uint(v1)) & 63))
                                            _ = fi47
                                            _ = v49
                                        }
                                        fmt.Println(v2)
                                    }
                                    for fi50 := 0; fi50 < 9; fi50++ {
                                        for fi51 := 0; fi51 < 1; fi51++ {
                                            var v52 int8 = -1
                                            _ = fi51
                                            _ = v52
                                        }
                                        _ = fi50
                                    }
                                }
                            }
                        case 538:
                            var v53 uint32 = (v3 ^ (v3 >> ((65535) & 63)))
                            _ = v53
                        default:
                            v3 = (v3 | (v3 * genFn0(255, v35)))
                        }
                    }
                    if len(v9) >= 2 {
                        v9 = v9[1:len(v9)]
                    }
                }
            } else {
                var v54 int8 = 127
                v1 = v1
                _ = v54
            }
            var v55 int32 = 1
            _ = v55
        }
    }
    switch "a" {
    case "\n\t":
        var v56 uint16 = 255
        _ = v56
    case "a":
        fmt.Println(v1, v3)
    default:
        for fi57 := 0; fi57 < 5; fi57++ {
            v58 := GenS3{F0: uint32(v1), F1: v1}
            v9 = append(v9, fi57)
            for fi59 := 0; fi59 < 2; fi59++ {
                if v35 {
                    switch "foo bar" {
                    case "\n\t":
                        switch "a" {
                        case "a":
                            fmt.Println(v1, v3)
                        case "foo bar":
                            var v60 uint8 = v1
                            _ = v60
                        case "\n\t":
                            if (18446744073709551615 <= 42) {
                                var v61 int32 = -423
                                _ = v61
                            }
                        }
                        v62 := GenS1{F0: v1, F1: 421, F2: 42}
                        _ = v62
                    case "foo bar":
                        var v63 int8 = 100
                        _ = v63
                    case "a":
                        var v64 int64 = v2
                        _ = v64
                    }
                }
                _ = fi59
            }
            _ = fi57
            _ = v58
        }
    }
    fmt.Println(v35, v35)
    var v65 uint32 = (v3 % ((v3) | 1))
    fmt.Println(v65, v2, v35)
    if (true && ((-1 == 42) && false)) {
        if v35 {
            for fi66 := 0; fi66 < 11; fi66++ {
                v9 = append(v9, 9223372036854775807)
                if true {
                    var v67 int16 = 0
                    _ = len(v9)
                    _ = v67
                } else {
                    switch (-(fi66)) {
                    case 0:
                        v68 := GenS5{F0: v35, F1: 1, F2: (v35 || v35)}
                        _ = v68
                    default:
                        switch "" {
                        case "foo bar":
                            v1 += v1
                        case "\n\t":
                            v69 := GenS6{F0: 1}
                            _ = v69
                        default:
                            switch ("\n\t" + ("a" + ("foo bar" + ("\n\t" + "a")))) {
                            case "":
                                fmt.Println(v65, v1, v3)
                            }
                        }
                    }
                }
                v2 += v2
                _ = fi66
            }
        } else {
            v70 := GenS7{F0: -1, F1: 2147483647, F2: -2147483648}
            if true {
                switch "foo bar" {
                case "a":
                    _ = len(v9)
                }
                var v71 int16 = 0
                _ = v71
            }
            _ = v70
        }
        var v72 int16 = 32767
        if (0 <= 717) {
            if len(v9) >= 2 {
                v9 = v9[1:len(v9)]
            }
            fmt.Println(v35)
            switch -1 {
            case -2147483648:
                if (v35 && (-684 > 2147483647)) {
                    v3 = v65
                    var v73 bool = false
                    _ = v73
                }
            default:
                for fi74 := 0; fi74 < 18; fi74++ {
                    var v75 float64 = (((1e10 - 0.0) * 0.0) - (1e10 / ((0.1 / 1.0) + -2.25)))
                    if (true && ((!(v35 && (-9223372036854775808 >= int64(fi74)))) && false)) {
                        if v35 {
                            fmt.Println(fi74)
                            switch fi74 {
                            case -681:
                                v76 := GenS5{F0: (!false), F1: 32767, F2: v35}
                                _ = v76
                            default:
                                var v77 uint64 = 4294967295
                                _ = v77
                            }
                            v75 -= (v75 - v75)
                        } else {
                            for fi78 := 0; fi78 < 15; fi78++ {
                                v79 := GenS7{F0: -100, F1: 2147483647, F2: -316}
                                switch "foo bar" {
                                case "\n\t":
                                    if len(v9) > 0 {
                                        v9[0] = fi78
                                    }
                                case "hello":
                                    v80 := GenS8{F0: v1, F1: v65}
                                    _ = len(v9)
                                    var v81 uint64 = 21
                                    _ = v80
                                    _ = v81
                                case "a":
                                    switch "hello" {
                                    case "a":
                                        var v82 int8 = -128
                                        _ = v82
                                    case "":
                                        fmt.Println(fi78, v65, v1)
                                    case "hello":
                                        var v83 int32 = 741
                                        _ = v83
                                    }
                                }
                                if v35 {
                                    v65 += (v3 ^ (^(v3)))
                                    fmt.Println(fi74)
                                    if len(v9) >= 2 {
                                        v9 = v9[1:len(v9)]
                                    }
                                }
                                _ = fi78
                                _ = v79
                            }
                            v9 = append(v9, 42)
                            var v84 bool = true
                            _ = v84
                        }
                        v9 = append(v9, (fi74 ^ fi74))
                        fmt.Println(v72)
                    } else {
                        var v85 int16 = 720
                        v86 := GenS1{F0: v1, F1: uint16(v1), F2: 1}
                        var v87 int8 = 22
                        _ = v85
                        _ = v86
                        _ = v87
                    }
                    var v88 uint = 255
                    _ = fi74
                    _ = v75
                    _ = v88
                }
            }
        }
        _ = v72
    }
    v35 = (v35 && v35)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v65:", v65)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v65
    _ = v8
    _ = v9
}
