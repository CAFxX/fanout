package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int }
type GenS2 struct { F0 int8; F1 string }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 string; F1 int8; F2 int16 }
type GenS6 struct { F0 int64; F1 int32 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 int }
type GenS9 struct { F0 bool; F1 string }
type GenS10 struct { F0 int64; F1 int16; F2 uint8 }
func genFn0(p0 uint16, p1 uint16) uint32 {
    var v4 int64 = 42
    var v5 uint16 = (p0 * p1)
    v4 -= (v4 >> ((42) & 63))
    var v6 int32 = 42
    v7 := []int{-130, -240, 32767, 0}
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return 1
}
func genFn1(p0 uint) uint32 {
    switch -574 {
    case 127:
        var v8 float32 = 3.14159
        _ = v8
    case 42:
        fmt.Println(p0)
    default:
        var v9 string = ("foo bar" + "\n\t")
        _ = v9
    }
    switch "\n\t" {
    case "a":
        for fi10 := 0; fi10 < 10; fi10++ {
            fmt.Println(p0)
            for fi11 := 0; fi11 < 4; fi11++ {
                v12 := []int{(fi11 | 154)}
                _ = fi11
                _ = v12
            }
            _ = fi10
        }
    case "\n\t":
        if true {
            if false {
                v13 := []int{42, 151, -211}
                p0 = (p0 + (p0 >> ((p0) & 63)))
                for fi14 := 0; fi14 < 6; fi14++ {
                    for fi15 := 0; fi15 < 19; fi15++ {
                        p0 = p0
                        if true {
                            if true {
                                v16 := GenS1{F0: 736, F1: fi14}
                                var v17 int8 = 100
                                v13 = append(v13, fi15)
                                var v18 uint16 = 2
                                if len(v13) >= 2 {
                                    v13 = v13[1:len(v13)]
                                }
                                _ = v16
                                _ = v17
                                _ = v18
                            }
                            p0 = (p0 + (p0 + p0))
                        } else {
                            switch "a" {
                            case "foo bar":
                                var v19 uint = (p0 | (^(p0)))
                                v19 = 2
                                _ = v19
                            case "a":
                                if (!false) {
                                    var v20 uint = p0
                                    _ = v20
                                }
                            default:
                                var v21 int64 = 1
                                if (127 > -1) {
                                    fmt.Println(fi14, fi15, fi14)
                                    var v22 string = "hello"
                                    v23 := GenS1{F0: 0, F1: 42}
                                    _ = v22
                                    _ = v23
                                } else {
                                    var v24 uint8 = 184
                                    var v25 uint = (p0 % ((p0) | 1))
                                    if len(v13) > 0 {
                                        v13[0] = (fi14 >> (((^(p0))) & 63))
                                    }
                                    _ = v24
                                    _ = v25
                                }
                                _ = v21
                            }
                        }
                        v13 = append(v13, fi14)
                        _ = fi15
                    }
                    _ = fi14
                }
                _ = v13
            } else {
                p0 = (^(p0))
                v26 := GenS2{F0: -88, F1: ("" + ("hello" + ("\n\t" + "hello")))}
                _ = v26
            }
            v27 := GenS2{F0: 59, F1: ("foo bar" + ("a" + ("a" + "")))}
            _ = v27
        }
    }
    _ = p0
    return 635
}

func main() {
    var v1 string = "hello"
    v2 := make(map[uint32]int)
    var v3 uint = 2
    v28 := GenS2{F0: 42, F1: ("hello" + "hello")}
    v29 := GenS3{F0: 1}
    if true {
        switch 32768 {
        case -904:
            v3 = v3
        case 32768:
            fmt.Println(v1)
        default:
            v30 := GenS4{F0: v3}
            _ = v30
        }
        for fi31 := 0; fi31 < 19; fi31++ {
            switch (v1 + (v1 + v1)) {
            case "hello":
                switch "hello" {
                case "foo bar":
                    v32 := []int{}
                    _ = v32
                }
            case "":
                v33 := GenS5{F0: v1, F1: -84, F2: 32767}
                _ = v33
            default:
                for fi34 := 0; fi34 < 4; fi34++ {
                    v3 = (v3 << (((v3 | (v3 | 42))) & 63))
                    v35 := []int{-2147483648, fi31, fi31, -293}
                    _ = fi34
                    _ = v35
                }
            }
            _ = fi31
        }
        if true {
            v3 -= v3
            v1 = "hello"
            v3 += 1
        }
    }
    v36 := GenS6{F0: -9223372036854775808, F1: -269}
    if (true && (255 > 220)) {
        v37 := []int{32767, 343}
        _ = v37
    }
    if false {
        for fi38 := 0; fi38 < 18; fi38++ {
            v39 := GenS2{F0: 127, F1: (v1 + v1)}
            if false {
                for fi40 := 0; fi40 < 9; fi40++ {
                    v41 := GenS7{F0: (v3 - v3)}
                    v42 := []int{127, fi38, (-(fi40)), (-(fi40))}
                    _ = fi40
                    _ = v41
                    _ = v42
                }
                if ((-128 < -128) && ((true || true) || false)) {
                    var v43 bool = (0.1 == -2.25)
                    _ = v43
                }
                if false {
                    for fi44 := 0; fi44 < 13; fi44++ {
                        fmt.Println(fi38)
                        _ = fi44
                    }
                    switch "a" {
                    case "":
                        var v45 int32 = 763
                        _ = v45
                    case "\n\t":
                        var v46 int = fi38
                        _ = v46
                    case "a":
                        v47 := GenS8{F0: fi38}
                        _ = v47
                    default:
                        var v48 string = "a"
                        _ = v48
                    }
                }
            } else {
                v2[4294967295] = fi38
                for fi49 := 0; fi49 < 12; fi49++ {
                    v3 = (v3 ^ 0)
                    if (!(775 > 2)) {
                        v50 := []int{}
                        _ = v50
                    } else {
                        var v51 string = v1
                        _ = v51
                    }
                    _ = fi49
                }
                var v52 int64 = 9223372036854775807
                _ = v52
            }
            _ = fi38
            _ = v39
        }
    } else {
        v3 += 0
        fmt.Println(v3)
    }
    delete(v2, 2147483648)
    var v53 int64 = 42
    var v54 uint16 = 3
    v55 := GenS2{F0: 100, F1: "a"}
    if false {
        if (0 >= 2) {
            fmt.Println(v54)
            if false {
                var v56 uint64 = 171
                v57 := GenS7{F0: v3}
                switch "a" {
                case "":
                    for fi58 := 0; fi58 < 6; fi58++ {
                        if true {
                            v54 = 2
                            var v59 int = fi58
                            _ = v59
                        } else {
                            if true {
                                if true {
                                    var v60 int16 = 0
                                    if (2147483647 > 0) {
                                        fmt.Println(v54, v60, v56)
                                        v53 += v53
                                    } else {
                                        for fi61 := 0; fi61 < 1; fi61++ {
                                            if (227 > 222) {
                                                var v62 int64 = v53
                                                switch ("hello" + v1) {
                                                case "hello":
                                                    var v63 float64 = -2.25
                                                    _ = v63
                                                case "\n\t":
                                                    v60 = v60
                                                default:
                                                    v64 := GenS9{F0: (766 != (v56 << ((uint(v54)) & 63))), F1: ""}
                                                    _ = v64
                                                }
                                                var v65 int64 = (v62 | -940)
                                                _ = v62
                                                _ = v65
                                            } else {
                                                v66 := GenS8{F0: 32767}
                                                var v67 uint16 = (v54 | (v54 & (v54 + (^(v54)))))
                                                _ = v66
                                                _ = v67
                                            }
                                            _ = fi61
                                        }
                                    }
                                    _ = v60
                                }
                                v68 := GenS10{F0: v53, F1: -319, F2: 37}
                                _ = v68
                            } else {
                                v69 := v2[300]
                                if false {
                                    v70 := GenS1{F0: v54, F1: v69}
                                    _ = v70
                                }
                                switch v1 {
                                case "a":
                                    for fi71 := 0; fi71 < 9; fi71++ {
                                        if ((2147483647 <= genFn0((v54 / (((v54 - (v54 - v54))) | 1)), (v54 >> (((v3 / (((v3 | (v3 - (^(v3))))) | 1))) & 63)))) && true) {
                                            v1 = "a"
                                            v72 := GenS9{F0: (1 != -1), F1: "a"}
                                            if false {
                                                v3 = v3
                                                v1 = v1
                                            } else {
                                                fmt.Println(fi71)
                                                var v73 float64 = 0.1
                                                _ = v73
                                            }
                                            _ = v72
                                        }
                                        switch (v1 + v1) {
                                        case "foo bar":
                                            v53 = (v53 >> ((18446744073709551615) & 63))
                                        default:
                                            v56 += 1
                                        }
                                        var v74 int64 = (v53 & v53)
                                        _ = fi71
                                        _ = v74
                                    }
                                case "hello":
                                    delete(v2, 65535)
                                }
                                _ = v69
                            }
                            v1 = ("\n\t" + ((("hello" + "foo bar") + "a") + ((v1 + v1) + v1)))
                        }
                        _ = fi58
                    }
                case "foo bar":
                    v75 := []int{-722, -9223372036854775808, -9223372036854775808, 0}
                    _ = v75
                case "a":
                    var v76 uint = (v3 / (((v3 * 0)) | 1))
                    _ = v76
                default:
                    var v77 string = v1
                    _ = v77
                }
                _ = v56
                _ = v57
            } else {
                v3 += (v3 - 255)
            }
            v53 -= (^(v53))
        }
        v78 := GenS2{F0: 1, F1: (("" + (v1 + "\n\t")) + ("" + v1))}
        _ = v78
    } else {
        var v79 int = 691
        for fi80 := 0; fi80 < 12; fi80++ {
            var v81 float64 = 1e10
            _ = fi80
            _ = v81
        }
        switch v1 {
        case "hello":
            var v82 uint = v3
            _ = v82
        case "foo bar":
            for fi83 := 0; fi83 < 3; fi83++ {
                v1 = v1
                for fi84 := 0; fi84 < 20; fi84++ {
                    switch ((v1 + v1) + v1) {
                    case "a":
                        var v85 int32 = 42
                        _ = v85
                    }
                    _ = fi84
                }
                _ = fi83
            }
        case "a":
            for fi86 := 0; fi86 < 5; fi86++ {
                if ((false || (!false)) || false) {
                    v54 += (v54 >> (((v3 | v3)) & 63))
                    fmt.Println(v54, v53)
                    v87 := GenS10{F0: (v53 % ((v53) | 1)), F1: 42, F2: 60}
                    _ = v87
                } else {
                    v54 = (v54 - 42)
                    v1 = v1
                }
                v53 = (v53 >> (((v3 / ((680) | 1))) & 63))
                _ = fi86
            }
        default:
            fmt.Println(v79, v53, v79)
        }
        _ = v79
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v53:", v53)
    fmt.Println("v54:", v54)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v36
    _ = v53
    _ = v54
    _ = v55
}
