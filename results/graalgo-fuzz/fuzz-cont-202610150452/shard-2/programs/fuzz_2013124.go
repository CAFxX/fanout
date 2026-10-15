package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string }
type GenS2 struct { F0 uint64; F1 uint16; F2 uint }
type GenS3 struct { F0 int; F1 uint8 }
type GenS4 struct { F0 uint8; F1 int16; F2 bool }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 string; F1 int32 }
type GenS7 struct { F0 uint; F1 int8 }
type GenS8 struct { F0 uint8; F1 int16 }
type GenS9 struct { F0 int8; F1 uint64; F2 uint8 }
type GenS10 struct { F0 uint32; F1 bool; F2 uint8 }
type GenS11 struct { F0 int16; F1 uint64; F2 uint }
type GenS12 struct { F0 int; F1 uint8; F2 int16 }
func genFn0(p0 int64, p1 uint64) uint32 {
    v4 := []int{}
    fmt.Println(p1, p0)
    var v5 uint64 = p1
    _ = len(v4)
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 1
}
func genFn1(p0 int32, p1 int8) uint16 {
    for fi6 := 0; fi6 < 1; fi6++ {
        var v7 uint64 = 1
        for fi8 := 0; fi8 < 3; fi8++ {
            v9 := GenS1{F0: fi8, F1: ("foo bar" + ("foo bar" + ("a" + "")))}
            _ = fi8
            _ = v9
        }
        v10 := GenS1{F0: fi6, F1: "hello"}
        _ = fi6
        _ = v10
        _ = v7
    }
    var v11 int16 = -318
    _ = p0
    _ = p1
    _ = v11
    return 2
}

func main() {
    var v1 float32 = 0.0
    var v2 uint8 = 127
    v3 := make(map[string]bool)
    v2 += (v2 + (v2 - v2))
    var v12 float64 = 1e10
    for fi13 := 0; fi13 < 1; fi13++ {
        fmt.Println(v1, fi13)
        _ = fi13
    }
    v12 -= (v12 / float64(-2.25))
    if (!(1 >= 716)) {
        v2 = v2
        switch ("foo bar" + "foo bar") {
        case "":
            switch 127 {
            case 0:
                if (!(false && true)) {
                    var v14 uint8 = (v2 / ((v2) | 1))
                    var v15 bool = (true || true)
                    switch 127 {
                    case 0:
                        var v16 int8 = -123
                        _ = v16
                    case -891:
                        v17 := GenS1{F0: -260, F1: "a"}
                        _ = v17
                    case 127:
                        var v18 int32 = 1
                        _ = v18
                    }
                    _ = v14
                    _ = v15
                } else {
                    v19 := GenS1{F0: 127, F1: "hello"}
                    var v20 int8 = 1
                    v21 := GenS1{F0: 0, F1: "\n\t"}
                    _ = v19
                    _ = v20
                    _ = v21
                }
            case 32767:
                v1 = ((v1 + -2.25) + (v1 * v1))
            case 4294967295:
                for fi22 := 0; fi22 < 13; fi22++ {
                    v23 := GenS2{F0: 42, F1: 32767, F2: 2}
                    if true {
                        switch "\n\t" {
                        case "foo bar":
                            if (!true) {
                                v24 := GenS3{F0: fi22, F1: (v2 + (v2 & 20))}
                                _ = v24
                            }
                        }
                    } else {
                        v1 = v1
                        var v25 uint64 = 42
                        _ = v25
                    }
                    v26 := GenS2{F0: 1, F1: 917, F2: 498}
                    _ = fi22
                    _ = v23
                    _ = v26
                }
            }
        }
        var v27 uint64 = 826
        _ = v27
    } else {
        for fi28 := 0; fi28 < 14; fi28++ {
            if true {
                fmt.Println(v1)
                v1 -= v1
                switch ((("foo bar" + "") + "") + ((("a" + ("hello" + "hello")) + "hello") + "foo bar")) {
                case "hello":
                    var v29 uint64 = 18446744073709551615
                    _ = v29
                }
            } else {
                for fi30 := 0; fi30 < 12; fi30++ {
                    v3["a"] = (true || (!false))
                    _ = fi30
                }
                if true {
                    v31 := []int{(fi28 - 0), -358, fi28}
                    _ = v31
                }
                var v32 uint16 = genFn1(-1, -1)
                _ = v32
            }
            _ = fi28
        }
    }
    v2 -= v2
    if ((v1 > 1.5) || (false && (!false))) {
        v12 += 3.14159
    } else {
        var v33 int64 = 42
        if (128 != 32768) {
            v12 = float64(1.5)
            if (v2 <= 43) {
                v1 = 0.1
                v34 := GenS2{F0: 4294967295, F1: 0, F2: 2}
                _ = v34
            }
            v12 = (v12 - (v12 / v12))
        }
        _ = v33
    }
    switch -351 {
    case 0:
        if (!true) {
            v2 = 177
        } else {
            var v35 uint32 = 1
            _ = v35
        }
    case -1:
        for fi36 := 0; fi36 < 15; fi36++ {
            switch (fi36 / ((fi36) | 1)) {
            case 2147483647:
                v37 := GenS3{F0: (fi36 ^ (fi36 * fi36)), F1: (v2 & v2)}
                _ = v37
            case 575:
                switch "hello" {
                case "\n\t":
                    for fi38 := 0; fi38 < 4; fi38++ {
                        v39 := GenS1{F0: (fi38 % ((-348) | 1)), F1: ("\n\t" + "hello")}
                        v2 = (v2 / ((v2) | 1))
                        _ = fi38
                        _ = v39
                    }
                case "foo bar":
                    v40 := GenS4{F0: (v2 + 0), F1: 1, F2: ((false || (!false)) && (!true))}
                    _ = v40
                }
            }
            var v41 string = ""
            _ = fi36
            _ = v41
        }
    case 177:
        var v42 int8 = 32
        _ = v42
    default:
        v3["\n\t"] = false
    }
    v2 += (^(v2))
    switch "foo bar" {
    case "hello":
        if (927 <= 4294967295) {
            switch "" {
            case "a":
                for fi43 := 0; fi43 < 6; fi43++ {
                    v44 := GenS1{F0: (fi43 % (((fi43 ^ 125)) | 1)), F1: "a"}
                    v1 -= float32((0.0 * 1e10))
                    _ = fi43
                    _ = v44
                }
            case "foo bar":
                v45 := GenS5{F0: 0}
                _ = v45
            case "\n\t":
                var v46 uint8 = v2
                _ = v46
            default:
                fmt.Println(v1, v2, v2)
            }
            v47 := GenS6{F0: ("a" + "hello"), F1: 0}
            _ = v47
        }
    }
    for fi48 := 0; fi48 < 11; fi48++ {
        var v49 uint8 = 42
        v12 += (v12 + v12)
        for fi50 := 0; fi50 < 18; fi50++ {
            if (65535 != 362) {
                if false {
                    v51 := GenS2{F0: 18446744073709551615, F1: 0, F2: 2}
                    if ((!((v1 <= (v1 * ((-2.25 + ((float32(((((float64(v1) * 1.5) - 0.1) * v12) - v12)) * (-2.25 - (((1.5 - (v1 + v1)) / v1) * v1))) / ((v1 - v1) - 0.1))) * 0.1))) || (!false))) || false) {
                        v52 := GenS2{F0: 0, F1: 482, F2: 0}
                        _ = v52
                    }
                    switch "foo bar" {
                    case "\n\t":
                        if (2147483648 == genFn0(-116, 4294967295)) {
                            var v53 int16 = 844
                            v54 := GenS7{F0: 163, F1: int8(v2)}
                            if false {
                                v49 += (^(v2))
                                if false {
                                    var v55 int = (fi50 << ((2) & 63))
                                    v12 += v12
                                    var v56 int = -9223372036854775808
                                    switch fi48 {
                                    case -371:
                                        v49 -= (v2 % (((v2 | v49)) | 1))
                                    case -2147483648:
                                        if ((0.0 * 1.5) == (1e10 + (((v12 * v12) - v12) - (0.0 * float64((v1 + 0.0)))))) {
                                            v2 = (v2 - 128)
                                            v57 := []int{v55, -290}
                                            _ = v57
                                        } else {
                                            if (128 < (v49 + 228)) {
                                                v58 := []int{42, (v55 & (-(v55)))}
                                                _ = v58
                                            } else {
                                                var v59 int32 = -2147483648
                                                _ = v59
                                            }
                                        }
                                        var v60 int64 = 916
                                        if false {
                                            var v61 float64 = v12
                                            switch "foo bar" {
                                            case "\n\t":
                                                v62 := GenS8{F0: (v49 | v2), F1: 0}
                                                _ = v62
                                            }
                                            var v63 int16 = v53
                                            _ = v61
                                            _ = v63
                                        }
                                        _ = v60
                                    default:
                                        v64 := v3["\n\t"]
                                        _ = v64
                                    }
                                    _ = v55
                                    _ = v56
                                } else {
                                    switch "\n\t" {
                                    case "":
                                        v65 := GenS9{F0: 1, F1: 754, F2: 56}
                                        _ = v65
                                    case "hello":
                                        v3[("\n\t" + ("foo bar" + (("foo bar" + "") + "\n\t")))] = (4294967295 <= 2)
                                    default:
                                        v49 -= 42
                                    }
                                }
                                var v66 int32 = -672
                                if false {
                                    var v67 uint16 = genFn1(v66, -128)
                                    if (v49 > v49) {
                                        var v68 int = (fi50 & -9223372036854775808)
                                        delete(v3, "foo bar")
                                        v49 = (v2 % ((v49) | 1))
                                        _ = v68
                                    }
                                    fmt.Println(v12, v53)
                                    _ = v67
                                } else {
                                    switch ("a" + "hello") {
                                    case "\n\t":
                                        v69 := GenS10{F0: 190, F1: (true || (42 <= v49)), F2: (v49 - (v2 & v49))}
                                        _ = v69
                                    default:
                                        var v70 uint8 = v49
                                        v12 -= (v12 * ((v12 * v12) - v12))
                                        switch 42 {
                                        case -195:
                                            if true {
                                                v70 = 2
                                                v2 = (v70 - v70)
                                            } else {
                                                fmt.Println(v66)
                                                v2 = (v2 ^ v2)
                                                var v71 uint = 2
                                                _ = v71
                                            }
                                        }
                                        _ = v70
                                    }
                                    var v72 uint64 = 47
                                    _ = v72
                                }
                                _ = v66
                            }
                            _ = v53
                            _ = v54
                        } else {
                            if false {
                                v73 := GenS2{F0: 1, F1: 1, F2: 18446744073709551615}
                                var v74 int = 975
                                _ = v73
                                _ = v74
                            }
                            v1 = 0.1
                            switch (fi48 % ((-9223372036854775808) | 1)) {
                            case 615:
                                var v75 int8 = 127
                                _ = v75
                            }
                        }
                    case "a":
                        v2 = v2
                    case "foo bar":
                        var v76 uint8 = (^(v2))
                        _ = v76
                    }
                    _ = v51
                }
            } else {
                var v77 int = (^(fi50))
                _ = v77
            }
            var v78 int = (^(fi48))
            _ = fi50
            _ = v78
        }
        _ = fi48
        _ = v49
    }
    var v79 uint8 = (v2 % ((v2) | 1))
    v80 := GenS10{F0: 807, F1: (false || (!(float32((0.1 - v12)) == 1.5))), F2: (v79 % ((v79) | 1))}
    _ = len(v3)
    v81 := GenS3{F0: -9223372036854775808, F1: v79}
    v82 := GenS4{F0: (v2 % (((^(v79))) | 1)), F1: -1, F2: true}
    if false {
        if (!(1 >= 529)) {
            v83 := GenS7{F0: uint(v2), F1: 1}
            _ = v83
        }
        v84 := GenS3{F0: 490, F1: (v79 % ((72) | 1))}
        if true {
            v2 = (v79 + (v79 + (v2 + (v79 / ((v2) | 1)))))
            if true {
                v79 += v2
            }
            var v85 int = 42
            _ = v85
        } else {
            switch -1 {
            case -174:
                v86 := GenS11{F0: 1, F1: 1, F2: 42}
                _ = v86
            case 9223372036854775807:
                v87 := GenS12{F0: 454, F1: v2, F2: -32768}
                _ = v87
            default:
                for fi88 := 0; fi88 < 11; fi88++ {
                    for fi89 := 0; fi89 < 19; fi89++ {
                        _ = len(v3)
                        switch "foo bar" {
                        case "foo bar":
                            v2 = 127
                        default:
                            var v90 int8 = 121
                            v12 = (v12 * ((v12 / 1.0) - v12))
                            _ = v90
                        }
                        fmt.Println(fi88, v79, v79)
                        _ = fi89
                    }
                    _ = fi88
                }
            }
            v91 := GenS5{F0: 0}
            var v92 int16 = -655
            _ = v91
            _ = v92
        }
        _ = v84
    }
    var v93 float64 = (float64(v1) / (v12 - 3.14159))
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v79:", v79)
    fmt.Println("v93:", v93)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v79
    _ = v80
    _ = v81
    _ = v82
    _ = v93
}
