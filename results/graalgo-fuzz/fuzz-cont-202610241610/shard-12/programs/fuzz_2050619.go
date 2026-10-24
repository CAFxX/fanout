package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16; F2 uint8 }
type GenS2 struct { F0 int64; F1 int; F2 uint16 }
type GenS3 struct { F0 uint32; F1 int16; F2 bool }
type GenS4 struct { F0 uint16; F1 uint32 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int32; F1 int32; F2 uint32 }
func genFn0(p0 uint32, p1 string, p2 bool) uint64 {
    if ((0 > int32(p0)) && p2) {
        switch -1 {
        case 672:
            p1 = p1
        case 127:
            v4 := []int{-54, 516, -1}
            _ = v4
        default:
            var v5 int = 9223372036854775807
            _ = v5
        }
        var v6 float64 = -2.25
        _ = v6
    } else {
        v7 := GenS1{F0: p2, F1: 76, F2: 128}
        v8 := GenS1{F0: p2, F1: -1, F2: 1}
        _ = v7
        _ = v8
    }
    var v9 int16 = 1
    p1 = ""
    p2 = false
    v10 := []int{-9223372036854775808, 128, 128}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return 1
}
func genFn1(p0 uint, p1 int64) uint {
    switch "hello" {
    case "foo bar":
        var v11 int16 = 1000
        _ = v11
    case "a":
        if (-292 == 541) {
            fmt.Println(p1, p1)
            var v12 int8 = 105
            switch 0 {
            case 2147483647:
                p1 = 1
            case -90:
                var v13 int64 = (p1 >> ((p0) & 63))
                _ = v13
            }
            _ = v12
        }
    default:
        var v14 float32 = 1e10
        _ = v14
    }
    var v15 int32 = 182
    _ = p0
    _ = p1
    _ = v15
    return p0
}

func main() {
    var v1 int32 = 2147483647
    var v2 int32 = (^(v1))
    v3 := make(map[int]int)
    for fi16 := 0; fi16 < 15; fi16++ {
        fmt.Println(v1, v1, v1)
        _ = fi16
    }
    for fi17 := 0; fi17 < 4; fi17++ {
        if true {
            switch (fi17 >> ((1) & 63)) {
            case 700:
                fmt.Println(v2)
            case 0:
                var v18 uint64 = genFn0(1, "\n\t", (!true))
                _ = v18
            case -22:
                v19 := GenS1{F0: (0 > 0), F1: 0, F2: 1}
                _ = v19
            default:
                for fi20 := 0; fi20 < 3; fi20++ {
                    if false {
                        v1 = (v1 - -822)
                        var v21 float32 = 3.14159
                        switch ("" + "\n\t") {
                        case "":
                            v2 = (v2 - v2)
                        case "a":
                            var v22 uint = 2
                            _ = v22
                        case "\n\t":
                            fmt.Println(v2, fi17, v1)
                        }
                        _ = v21
                    } else {
                        var v23 float32 = (3.14159 * 3.14159)
                        _ = v23
                    }
                    switch (fi17 | fi20) {
                    case 32767:
                        v24 := GenS2{F0: 9223372036854775807, F1: (^(fi17)), F2: 0}
                        _ = v24
                    default:
                        v1 -= v1
                    }
                    var v25 uint16 = 42
                    _ = fi20
                    _ = v25
                }
            }
            if (!false) {
                v26 := []int{(fi17 / (((^(fi17))) | 1)), fi17, (fi17 - fi17), (fi17 & fi17)}
                switch "\n\t" {
                case "hello":
                    var v27 bool = (127 == 157)
                    _ = v27
                case "":
                    var v28 float64 = (1e10 * 1.5)
                    _ = v28
                case "foo bar":
                    v29 := GenS2{F0: 42, F1: (fi17 << ((2) & 63)), F2: 0}
                    _ = v29
                default:
                    v30 := GenS3{F0: 464, F1: 1000, F2: true}
                    _ = v30
                }
                _ = v26
            }
        } else {
            for fi31 := 0; fi31 < 15; fi31++ {
                var v32 int8 = 100
                _ = fi31
                _ = v32
            }
            for fi33 := 0; fi33 < 8; fi33++ {
                var v34 int8 = -1
                _ = fi33
                _ = v34
            }
        }
        for fi35 := 0; fi35 < 4; fi35++ {
            v3[(fi35 * 32767)] = fi35
            if false {
                v1 -= (v1 % ((v2) | 1))
                var v36 int = -1
                var v37 uint = 255
                v2 = (v1 << (((v37 ^ v37)) & 63))
                _ = v36
                _ = v37
            }
            var v38 int8 = -1
            if (!((v2 & v2) != int32(fi35))) {
                fmt.Println(v38, v1)
                var v39 int = fi17
                _ = v39
            } else {
                if false {
                    if (!(1 > 1)) {
                        v38 += v38
                        if true {
                            v40 := GenS4{F0: 557, F1: 42}
                            switch "foo bar" {
                            case "\n\t":
                                v1 -= (-(v2))
                            case "":
                                v38 = v38
                            case "hello":
                                v41 := GenS5{F0: 0}
                                _ = v41
                            default:
                                fmt.Println(fi35)
                            }
                            fmt.Println(v2)
                            _ = v40
                        }
                    }
                } else {
                    fmt.Println(v1, v2)
                    if true {
                        if true {
                            var v42 int = (fi35 << ((1) & 63))
                            _ = len(v3)
                            v38 = -128
                            _ = v42
                        } else {
                            v43 := GenS6{F0: v2, F1: v2, F2: 65535}
                            _ = v43
                        }
                        var v44 string = ("hello" + ("hello" + ""))
                        _ = v44
                    }
                }
                switch "a" {
                case "":
                    var v45 bool = false
                    _ = v45
                case "foo bar":
                    if false {
                        v2 -= -1
                        var v46 uint32 = 65535
                        v47 := GenS2{F0: -1, F1: fi17, F2: 65535}
                        _ = v46
                        _ = v47
                    } else {
                        v48 := []int{}
                        v38 -= 42
                        if true {
                            var v49 uint64 = genFn0(2, (("" + "a") + ("a" + "a")), false)
                            _ = v49
                        }
                        _ = v48
                    }
                default:
                    v2 += (v2 | 42)
                }
            }
            v2 -= -1
            _ = fi35
            _ = v38
        }
        switch fi17 {
        case 128:
            v2 = (v2 / ((v2) | 1))
        case 42:
            v3[(fi17 | 128)] = fi17
        }
        _ = fi17
    }
    var v50 uint16 = 0
    if ((1.5 * 1e10) < 0.0) {
        fmt.Println(v50, v2, v1)
        v50 -= (v50 ^ v50)
    }
    var v51 int = 870
    v52 := GenS2{F0: 1, F1: (v51 & int(v1)), F2: 28}
    switch ("\n\t" + "foo bar") {
    case "":
        fmt.Println(v51, v2, v2)
    case "foo bar":
        v53 := GenS3{F0: 789, F1: 32767, F2: ((0.0 * -2.25) != 1.5)}
        _ = v53
    default:
        v50 = v50
    }
    v54 := GenS1{F0: (!true), F1: 0, F2: 127}
    delete(v3, (v51 + v51))
    v51 = (v51 | v51)
    if false {
        var v55 bool = (915 != 18446744073709551615)
        v50 = v50
        _ = v55
    }
    v56 := []int{v51, (^(v51))}
    switch (-(v51)) {
    case 42:
        if (!(0 >= int8(v1))) {
            for fi57 := 0; fi57 < 2; fi57++ {
                fmt.Println(v51, v51)
                for fi58 := 0; fi58 < 15; fi58++ {
                    v59 := v3[(fi57 ^ -908)]
                    var v60 uint32 = 2
                    delete(v3, v59)
                    _ = fi58
                    _ = v59
                    _ = v60
                }
                switch (fi57 & fi57) {
                case 4294967295:
                    v50 = v50
                }
                _ = fi57
            }
            if true {
                if false {
                    var v61 bool = true
                    v62 := GenS4{F0: 2, F1: 42}
                    v61 = v61
                    _ = v61
                    _ = v62
                }
                v2 = v1
                var v63 uint64 = 0
                _ = v63
            }
        } else {
            var v64 uint8 = 127
            _ = v64
        }
    case -924:
        if (0.0 > 1.5) {
            v65 := GenS6{F0: (v2 ^ -956), F1: v2, F2: 188}
            _ = v65
        } else {
            var v66 int64 = -9223372036854775808
            if len(v56) >= 2 {
                v56 = v56[1:len(v56)]
            }
            var v67 uint8 = 1
            _ = v66
            _ = v67
        }
    default:
        if (!(!false)) {
            fmt.Println(v2)
        } else {
            fmt.Println(v50, v1, v2)
            v1 = int32(v50)
            v2 = (v1 >> ((793) & 63))
        }
    }
    if (true && ((42 != 100) && (2 != 856))) {
        var v68 uint8 = 43
        _ = v68
    }
    v51 -= 127
    if false {
        var v69 uint8 = 127
        fmt.Println(v1, v51, v69)
        v2 -= v1
        _ = v69
    }
    _ = len(v56)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    fmt.Println("slicelen v56:", len(v56), v56)
    _ = v1
    _ = v2
    _ = v3
    _ = v50
    _ = v51
    _ = v52
    _ = v54
    _ = v56
}
