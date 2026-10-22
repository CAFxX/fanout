package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 int64 }
type GenS2 struct { F0 uint16; F1 bool; F2 bool }
type GenS3 struct { F0 int64; F1 uint }
type GenS4 struct { F0 int16; F1 uint64; F2 int }
type GenS5 struct { F0 string; F1 uint64; F2 int32 }
type GenS6 struct { F0 uint64; F1 int16; F2 int32 }
type GenS7 struct { F0 int64; F1 int; F2 int64 }
func genFn0() int {
    var v6 int32 = 42
    var v7 uint = 0
    _ = v6
    _ = v7
    return 128
}

func main() {
    var v1 int64 = -133
    var v2 int8 = -3
    var v3 int64 = v1
    var v4 int = 127
    var v5 uint16 = 2
    v8 := GenS1{F0: v5, F1: 42, F2: v1}
    var v9 bool = (v5 != 220)
    v5 -= v5
    v4 += -1
    switch "foo bar" {
    case "a":
        v10 := []int{}
        _ = v10
    case "foo bar":
        if ((42 == (-(v3))) || false) {
            v3 -= v3
            v11 := []int{(v4 % ((v4) | 1)), -9223372036854775808}
            _ = v11
        }
    case "":
        fmt.Println(v2, v1, v4)
    default:
        var v12 int32 = 1
        _ = v12
    }
    for fi13 := 0; fi13 < 6; fi13++ {
        if (!(!(2147483647 <= 2))) {
            for fi14 := 0; fi14 < 7; fi14++ {
                fmt.Println(v3)
                if (int16(v2) == 32767) {
                    var v15 int64 = -259
                    v4 += (fi14 % (((fi14 | 127)) | 1))
                    _ = v15
                } else {
                    v16 := GenS1{F0: v5, F1: 179, F2: 0}
                    switch 9223372036854775807 {
                    case 4294967295:
                        fmt.Println(v3, fi14, v4)
                    case 32767:
                        if false {
                            var v17 uint = 1
                            var v18 bool = (2 >= 65535)
                            _ = v17
                            _ = v18
                        } else {
                            v19 := []int{}
                            if (4294967295 == 228) {
                                v1 -= -9223372036854775808
                            }
                            _ = v19
                        }
                    default:
                        v5 = 2
                    }
                    v20 := []int{}
                    _ = v16
                    _ = v20
                }
                fmt.Println(v3, v1)
                _ = fi14
            }
            var v21 bool = (false && true)
            _ = v21
        } else {
            fmt.Println(v5, v4, v3)
        }
        _ = fi13
    }
    var v22 int8 = -1
    for fi23 := 0; fi23 < 16; fi23++ {
        v24 := []int{(v4 % (((^(v4))) | 1)), int(v2), v4}
        _ = fi23
        _ = v24
    }
    if v9 {
        switch genFn0() {
        case -1:
            v25 := []int{}
            _ = v25
        case 1:
            var v26 int16 = -1
            _ = v26
        default:
            switch v4 {
            case -1:
                v27 := []int{v4}
                _ = v27
            default:
                switch "\n\t" {
                case "hello":
                    for fi28 := 0; fi28 < 11; fi28++ {
                        v9 = v9
                        if ((v5 - (v5 & (v5 & v5))) < v5) {
                            if v9 {
                                fmt.Println(v5)
                            }
                        }
                        _ = fi28
                    }
                case "a":
                    switch v4 {
                    case 42:
                        v3 = v3
                    case -376:
                        v9 = (-32768 > 42)
                    case -339:
                        v3 -= int64(v22)
                    }
                case "\n\t":
                    v29 := GenS1{F0: (v5 % ((255) | 1)), F1: 2, F2: (-(v1))}
                    _ = v29
                default:
                    fmt.Println(v3)
                }
            }
        }
        v30 := GenS1{F0: (^(v5)), F1: 158, F2: v3}
        for fi31 := 0; fi31 < 8; fi31++ {
            for fi32 := 0; fi32 < 16; fi32++ {
                fmt.Println(v5)
                _ = fi32
            }
            _ = fi31
        }
        _ = v30
    } else {
        for fi33 := 0; fi33 < 20; fi33++ {
            var v34 uint8 = 96
            var v35 float32 = float32(1.5)
            v36 := []int{genFn0(), fi33, (fi33 - fi33)}
            _ = fi33
            _ = v34
            _ = v35
            _ = v36
        }
        for fi37 := 0; fi37 < 11; fi37++ {
            for fi38 := 0; fi38 < 1; fi38++ {
                var v39 float32 = (float32((1.5 / 1.0)) + (0.0 + 3.14159))
                v40 := GenS2{F0: (v5 >> ((42) & 63)), F1: (!((0.0 + ((v39 * (v39 / (v39 * 3.14159))) - 1.5)) == v39)), F2: v9}
                _ = fi38
                _ = v39
                _ = v40
            }
            v22 = v2
            switch -143 {
            case 140:
                switch v4 {
                case 4294967295:
                    v41 := []int{831}
                    _ = v41
                }
            default:
                fmt.Println(v1)
            }
            _ = fi37
        }
    }
    v42 := []int{}
    var v43 uint8 = 255
    switch "a" {
    case "":
        for fi44 := 0; fi44 < 15; fi44++ {
            var v45 bool = (872 == v5)
            if ((!v9) || (3.14159 <= 0.0)) {
                var v46 int64 = 1
                _ = v46
            } else {
                var v47 uint8 = (v43 % (((v43 << ((1) & 63))) | 1))
                fmt.Println(v22, v3, v9)
                _ = v47
            }
            fmt.Println(v43)
            _ = fi44
            _ = v45
        }
    case "hello":
        var v48 uint32 = 42
        _ = v48
    case "foo bar":
        for fi49 := 0; fi49 < 11; fi49++ {
            fmt.Println(v9, v9, v1)
            _ = fi49
        }
    default:
        for fi50 := 0; fi50 < 13; fi50++ {
            _ = len(v42)
            var v51 string = ("" + "foo bar")
            v52 := GenS1{F0: 2, F1: (v43 ^ v43), F2: (v3 / (((v1 / ((v1) | 1))) | 1))}
            _ = fi50
            _ = v51
            _ = v52
        }
    }
    v22 = (v22 >> ((18446744073709551615) & 63))
    if len(v42) > 0 {
        v42[0] = v4
    }
    for fi53 := 0; fi53 < 5; fi53++ {
        var v54 int16 = -981
        if v9 {
            var v55 uint32 = 42
            switch -9223372036854775808 {
            case 4294967295:
                var v56 string = "hello"
                _ = v56
            case -246:
                for fi57 := 0; fi57 < 12; fi57++ {
                    v55 += (v55 % ((42) | 1))
                    if false {
                        if (493 <= 46) {
                            v5 -= v5
                            switch (fi53 ^ (v4 ^ genFn0())) {
                            case -9223372036854775808:
                                if v9 {
                                    v58 := GenS1{F0: v5, F1: (v43 % (((v43 % (((v43 << ((65535) & 63))) | 1))) | 1)), F2: v1}
                                    v22 = 1
                                    v42 = append(v42, (v4 ^ -2147483648))
                                    _ = v58
                                }
                            default:
                                var v59 int64 = v1
                                _ = v59
                            }
                            v55 -= 682
                        } else {
                            var v60 int8 = 0
                            v22 = (v60 | v22)
                            var v61 int16 = (v54 * 32767)
                            _ = v60
                            _ = v61
                        }
                        if (!(v54 > (v54 % ((v54) | 1)))) {
                            v62 := GenS3{F0: 9223372036854775807, F1: 65535}
                            v3 = (^(v3))
                            _ = v62
                        }
                    }
                    _ = fi57
                }
            case -9223372036854775808:
                v42 = append(v42, genFn0())
            default:
                for fi63 := 0; fi63 < 14; fi63++ {
                    var v64 bool = false
                    switch 648 {
                    case 2147483647:
                        switch v4 {
                        case -796:
                            switch (fi53 / ((-712) | 1)) {
                            case -9223372036854775808:
                                if false {
                                    var v65 float32 = 1e10
                                    _ = v65
                                } else {
                                    var v66 int16 = (^(v54))
                                    fmt.Println(v2, v66)
                                    v55 += v55
                                    v43 -= (v43 - 127)
                                    v67 := GenS4{F0: v54, F1: 1, F2: -760}
                                    _ = v66
                                    _ = v67
                                }
                            case -1:
                                v2 += int8(v55)
                                var v68 int = v4
                                _ = v68
                            case 127:
                                if (true || (v4 != (v4 << ((2) & 63)))) {
                                    if (!true) {
                                        var v69 uint32 = v55
                                        if (!v64) {
                                            fmt.Println(v69)
                                            v70 := GenS4{F0: (v54 / ((v54) | 1)), F1: uint64(v69), F2: 723}
                                            _ = v70
                                        }
                                        v64 = (!(1e10 > 0.0))
                                        _ = v69
                                    } else {
                                        switch ("foo bar" + ("a" + "")) {
                                        case "":
                                            switch v4 {
                                            case -595:
                                                fmt.Println(fi63)
                                            case 85:
                                                v22 = v22
                                            default:
                                                switch 32768 {
                                                case 2147483647:
                                                    if (!v9) {
                                                        _ = len(v42)
                                                    }
                                                    v9 = v64
                                                    if (true && ((false || false) && (((v54 % (((v54 << ((2) & 63))) | 1)) <= v54) && v64))) {
                                                        fmt.Println(v3)
                                                        fmt.Println(v3, v64)
                                                        fmt.Println(v2, v3, v9)
                                                        var v71 bool = v9
                                                        v72 := GenS5{F0: "foo bar", F1: 2, F2: 1}
                                                        if (!v71) {
                                                            v3 = v1
                                                        } else {
                                                            v54 -= (^(v54))
                                                        }
                                                        _ = v71
                                                        _ = v72
                                                    }
                                                default:
                                                    v1 = (v3 >> ((1) & 63))
                                                }
                                            }
                                        }
                                        v2 -= (v2 | (v22 + v22))
                                        v55 = (v55 >> ((255) & 63))
                                        var v73 float32 = 1.5
                                        _ = v73
                                    }
                                    fmt.Println(v54, fi53, v1)
                                    v5 -= (v5 % (((v5 * v5)) | 1))
                                } else {
                                    v74 := GenS6{F0: 4294967295, F1: 32767, F2: -1}
                                    switch ((("a" + ((("hello" + "\n\t") + "foo bar") + "")) + "a") + "\n\t") {
                                    case "":
                                        var v75 bool = true
                                        _ = v75
                                    case "foo bar":
                                        _ = len(v42)
                                    default:
                                        var v76 string = ((("foo bar" + "\n\t") + ("\n\t" + "foo bar")) + "foo bar")
                                        _ = v76
                                    }
                                    _ = v74
                                }
                            }
                        case 42:
                            switch (fi53 ^ (v4 / ((fi63) | 1))) {
                            case 4294967295:
                                v77 := GenS3{F0: v3, F1: 255}
                                _ = v77
                            case -981:
                                var v78 uint8 = (v43 - (^(v43)))
                                _ = v78
                            case -9223372036854775808:
                                switch v4 {
                                case 2147483647:
                                    v2 += v22
                                case -743:
                                    fmt.Println(fi63)
                                case 4294967295:
                                    fmt.Println(v2, fi53, v5)
                                default:
                                    v55 -= 2
                                }
                            }
                        case -2147483648:
                            var v79 bool = (255 == 112)
                            _ = v79
                        }
                    }
                    v3 += (v1 / ((v1) | 1))
                    v80 := GenS6{F0: 4294967295, F1: 493, F2: 1}
                    _ = fi63
                    _ = v64
                    _ = v80
                }
            }
            _ = v55
        }
        if v9 {
            fmt.Println(v1, v22)
        }
        _ = fi53
        _ = v54
    }
    if v9 {
        switch v4 {
        case -1:
            for fi81 := 0; fi81 < 8; fi81++ {
                if ((0 >= 1) || (fi81 >= (^(v4)))) {
                    for fi82 := 0; fi82 < 8; fi82++ {
                        if len(v42) > 0 {
                            v42[0] = v4
                        }
                        _ = fi82
                    }
                } else {
                    if false {
                        if true {
                            v5 = 32767
                            var v83 uint32 = 2147483648
                            _ = v83
                        }
                        v84 := GenS7{F0: v3, F1: fi81, F2: (v1 % (((v1 | int64(v43))) | 1))}
                        _ = v84
                    }
                }
                _ = fi81
            }
        case 128:
            var v85 float32 = 0.0
            _ = v85
        default:
            if v9 {
                v22 = v2
                if (true || (0 < v2)) {
                    v86 := GenS3{F0: int64(v43), F1: 65535}
                    var v87 bool = true
                    var v88 uint = 18446744073709551615
                    _ = v86
                    _ = v87
                    _ = v88
                }
            }
        }
        fmt.Println(v3, v3, v1)
        v1 -= v3
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v42:", len(v42), v42)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v5
    _ = v8
    _ = v9
}
