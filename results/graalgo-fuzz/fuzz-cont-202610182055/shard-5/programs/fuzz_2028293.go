package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int32; F1 int }
type GenS3 struct { F0 string; F1 int16; F2 string }
type GenS4 struct { F0 int16; F1 string; F2 uint }
type GenS5 struct { F0 int8; F1 string }
type GenS6 struct { F0 int }
type GenS7 struct { F0 int64; F1 int64; F2 string }
type GenS8 struct { F0 int64; F1 int8; F2 int16 }
type GenS9 struct { F0 bool; F1 uint }

func main() {
    var v1 float32 = 0.1
    var v2 int64 = -836
    var v3 uint64 = 726
    var v4 int64 = 9223372036854775807
    v5 := make(map[uint32]string)
    var v6 uint16 = 255
    switch ("foo bar" + "") {
    case "hello":
        v7 := GenS1{F0: 244}
        _ = v7
    case "":
        v8 := []int{}
        _ = v8
    case "a":
        fmt.Println(v6, v1)
    default:
        v4 = (v2 | (v4 / ((v2) | 1)))
    }
    for fi9 := 0; fi9 < 19; fi9++ {
        v10 := GenS1{F0: 2}
        v1 = (0.1 * v1)
        for fi11 := 0; fi11 < 11; fi11++ {
            fmt.Println(v6, v6)
            _ = fi11
        }
        _ = fi9
        _ = v10
    }
    var v12 int8 = -128
    if false {
        v3 = 42
        delete(v5, 170)
    } else {
        for fi13 := 0; fi13 < 18; fi13++ {
            fmt.Println(v12)
            fmt.Println(v12, v2)
            _ = fi13
        }
        if (false || ("foo bar" == "")) {
            if true {
                v4 = v2
                for fi14 := 0; fi14 < 6; fi14++ {
                    var v15 float64 = float64(float32(-2.25))
                    delete(v5, 2147483647)
                    _ = fi14
                    _ = v15
                }
                if (!false) {
                    var v16 uint8 = 1
                    if ((!(v12 < (v12 % ((v12) | 1)))) && ((!false) || (v12 >= (^(v12))))) {
                        v12 = v12
                        v12 -= (v12 | 1)
                    } else {
                        v5[42] = "a"
                        v3 -= 2
                        v17 := GenS1{F0: (v16 + (v16 + v16))}
                        _ = v17
                    }
                    v18 := []int{-2147483648, -292}
                    _ = v16
                    _ = v18
                } else {
                    switch "\n\t" {
                    case "\n\t":
                        v6 = 0
                    default:
                        switch -1 {
                        case 561:
                            if false {
                                fmt.Println(v4, v4, v3)
                                v19 := []int{0, 294}
                                var v20 uint16 = 42
                                _ = v19
                                _ = v20
                            }
                        }
                    }
                }
            } else {
                v21 := GenS1{F0: 0}
                _ = v21
            }
            for fi22 := 0; fi22 < 6; fi22++ {
                if (false && true) {
                    if false {
                        if true {
                            var v23 float32 = v1
                            if false {
                                for fi24 := 0; fi24 < 15; fi24++ {
                                    fmt.Println(v2)
                                    _ = fi24
                                }
                                fmt.Println(v12, v6)
                            }
                            _ = v23
                        }
                        var v25 uint8 = 127
                        _ = v25
                    } else {
                        if ((0 == int16(v3)) && (v4 > v2)) {
                            _ = len(v5)
                        }
                    }
                    for fi26 := 0; fi26 < 3; fi26++ {
                        switch fi26 {
                        case 2147483647:
                            if false {
                                if (-713 <= 32767) {
                                    if false {
                                        v4 = v4
                                        if (false || (-2147483648 < 0)) {
                                            v4 += -973
                                            var v27 int16 = 1000
                                            fmt.Println(v1, fi26, v4)
                                            var v28 int8 = v12
                                            _ = v27
                                            _ = v28
                                        }
                                    } else {
                                        v29 := GenS2{F0: -2147483648, F1: fi22}
                                        _ = v29
                                    }
                                } else {
                                    v12 = (v12 - (v12 - (v12 / (((v12 / (((v12 & v12)) | 1))) | 1))))
                                    v6 = 2
                                    v30 := []int{fi22}
                                    _ = v30
                                }
                            } else {
                                v4 = 42
                                delete(v5, 2)
                                var v31 float64 = (((3.14159 * -2.25) / 3.14159) - 0.0)
                                v32 := GenS3{F0: "foo bar", F1: 1000, F2: ("" + ("" + "\n\t"))}
                                v3 -= v3
                                _ = v31
                                _ = v32
                            }
                        }
                        fmt.Println(fi26)
                        _ = fi26
                    }
                }
                _ = fi22
            }
            v33 := GenS4{F0: 106, F1: ("hello" + ""), F2: 255}
            _ = v33
        } else {
            switch 32768 {
            case 2147483647:
                v6 -= (v6 / ((0) | 1))
            case 32767:
                for fi34 := 0; fi34 < 4; fi34++ {
                    var v35 uint8 = 34
                    v36 := GenS5{F0: v12, F1: ("" + "a")}
                    _ = fi34
                    _ = v35
                    _ = v36
                }
            }
            for fi37 := 0; fi37 < 18; fi37++ {
                v1 -= v1
                v38 := GenS6{F0: (fi37 << ((18446744073709551615) & 63))}
                v2 += v2
                _ = fi37
                _ = v38
            }
        }
    }
    fmt.Println(v1, v4)
    var v39 uint16 = v6
    var v40 int16 = 548
    switch "foo bar" {
    case "hello":
        v41 := []int{230, -1}
        _ = v41
    case "":
        var v42 uint16 = v39
        _ = v42
    case "foo bar":
        v43 := GenS7{F0: (v4 ^ int64(v3)), F1: (^(v2)), F2: ""}
        _ = v43
    default:
        if true {
            v44 := GenS2{F0: 1, F1: 32767}
            if ((!(42 < 0)) || true) {
                v45 := []int{98}
                v6 = (v6 >> ((65535) & 63))
                _ = v45
            }
            _ = v44
        } else {
            switch -2147483648 {
            case -9223372036854775808:
                v3 = v3
            case 551:
                var v46 uint16 = v39
                _ = v46
            }
            for fi47 := 0; fi47 < 17; fi47++ {
                if (!true) {
                    for fi48 := 0; fi48 < 6; fi48++ {
                        var v49 bool = false
                        switch "hello" {
                        case "foo bar":
                            fmt.Println(fi47, fi47)
                        case "\n\t":
                            var v50 float64 = 1.5
                            _ = v50
                        }
                        var v51 uint32 = 2147483647
                        _ = fi48
                        _ = v49
                        _ = v51
                    }
                    for fi52 := 0; fi52 < 6; fi52++ {
                        switch (fi47 * (fi47 + 340)) {
                        case 127:
                            v3 += (v3 % ((42) | 1))
                        case -1:
                            v40 += -32768
                        default:
                            var v53 int64 = (v4 | int64(v40))
                            var v54 uint32 = 466
                            _ = len(v5)
                            var v55 string = ("foo bar" + "foo bar")
                            delete(v5, (v54 % ((1) | 1)))
                            fmt.Println(v40, v54)
                            fmt.Println(v4, v3, v40)
                            if true {
                                v56 := []int{fi47}
                                v55 = v55
                                _ = v56
                            } else {
                                var v57 int = (fi47 | fi52)
                                v58 := GenS5{F0: (v12 / ((v12) | 1)), F1: (v55 + ((v55 + v55) + "hello"))}
                                v59 := []int{32767, (fi52 ^ v57), fi47, 2147483647}
                                _ = v57
                                _ = v58
                                _ = v59
                            }
                            _ = v53
                            _ = v54
                            _ = v55
                        }
                        v5[2147483647] = (("\n\t" + "hello") + (("" + ((("foo bar" + "") + "") + ("" + ("foo bar" + "a")))) + "a"))
                        var v60 int = fi47
                        _ = fi52
                        _ = v60
                    }
                }
                switch "a" {
                case "foo bar":
                    var v61 bool = (-363 < v4)
                    _ = v61
                case "hello":
                    var v62 int16 = (v40 % ((v40) | 1))
                    _ = v62
                case "\n\t":
                    v12 -= (v12 * (-(v12)))
                }
                _ = fi47
            }
        }
    }
    v2 = v2
    v1 += 1e10
    var v63 uint = 0
    var v64 int8 = v12
    v1 = v1
    if false {
        switch -786 {
        case 1:
            switch "" {
            case "hello":
                var v65 uint32 = 190
                _ = v65
            default:
                var v66 float64 = 3.14159
                _ = v66
            }
        case 128:
            for fi67 := 0; fi67 < 20; fi67++ {
                v64 = v64
                v2 = int64(v39)
                if ((v6 * v39) != (^(v6))) {
                    if ((("hello" + ("" + "a")) + "hello") >= ("hello" + ("foo bar" + "\n\t"))) {
                        fmt.Println(v1)
                    } else {
                        var v68 int16 = (v40 ^ (v40 | (v40 ^ -1)))
                        for fi69 := 0; fi69 < 18; fi69++ {
                            v5[0] = ""
                            var v70 int32 = -2147483648
                            delete(v5, 2147483647)
                            _ = fi69
                            _ = v70
                        }
                        v71 := GenS8{F0: v2, F1: -5, F2: 349}
                        _ = v68
                        _ = v71
                    }
                    var v72 bool = true
                    switch "hello" {
                    case "\n\t":
                        if (fi67 < (fi67 % (((fi67 * fi67)) | 1))) {
                            var v73 int8 = v64
                            v5[2] = "hello"
                            if v72 {
                                for fi74 := 0; fi74 < 12; fi74++ {
                                    v39 = (v6 >> (((v63 + (v63 / (((v63 % ((301) | 1))) | 1)))) & 63))
                                    v75 := []int{}
                                    var v76 float64 = 3.14159
                                    v77 := GenS9{F0: v72, F1: 920}
                                    if ((v73 + v12) <= (v64 * (v64 * 2))) {
                                        v78 := GenS1{F0: 42}
                                        _ = v78
                                    } else {
                                        v79 := GenS9{F0: true, F1: (v63 * (v63 + 255))}
                                        v12 = v64
                                        _ = v79
                                    }
                                    _ = fi74
                                    _ = v75
                                    _ = v76
                                    _ = v77
                                }
                            } else {
                                v80 := v5[2147483647]
                                for fi81 := 0; fi81 < 19; fi81++ {
                                    var v82 uint16 = v6
                                    _ = fi81
                                    _ = v82
                                }
                                _ = v80
                            }
                            _ = v73
                        }
                    case "a":
                        if v72 {
                            v12 = (v12 & (v64 | (v12 << ((1) & 63))))
                            for fi83 := 0; fi83 < 14; fi83++ {
                                var v84 uint64 = v3
                                var v85 int16 = v40
                                var v86 int64 = 9223372036854775807
                                _ = fi83
                                _ = v84
                                _ = v85
                                _ = v86
                            }
                        }
                    default:
                        var v87 int64 = int64(v63)
                        _ = v87
                    }
                    _ = v72
                }
                _ = fi67
            }
        }
    } else {
        var v88 uint16 = (v39 & (^(v39)))
        var v89 uint16 = (v6 | v39)
        _ = v88
        _ = v89
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v63:", v63)
    fmt.Println("v64:", v64)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
    _ = v63
    _ = v64
}
