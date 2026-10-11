package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint }
type GenS2 struct { F0 bool; F1 uint8 }
type GenS3 struct { F0 int8; F1 int }
type GenS4 struct { F0 int16; F1 int; F2 uint }
type GenS5 struct { F0 int64; F1 uint16; F2 int }
type GenS6 struct { F0 uint16; F1 int32 }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 bool; F1 int16 }
type GenS9 struct { F0 uint; F1 int }

func main() {
    var v1 int8 = 42
    var v2 uint = 1
    v3 := make(map[uint32]string)
    var v4 int64 = 588
    v5 := make(map[int]bool)
    v6 := GenS1{F0: v4, F1: 42}
    for fi7 := 0; fi7 < 14; fi7++ {
        v8 := []int{(fi7 >> ((255) & 63)), (-(fi7)), 9223372036854775807, (fi7 | fi7)}
        if false {
            switch (fi7 ^ (fi7 ^ (fi7 ^ fi7))) {
            case -2147483648:
                switch "\n\t" {
                case "a":
                    fmt.Println(v1)
                case "":
                    for fi9 := 0; fi9 < 11; fi9++ {
                        v4 -= 253
                        v10 := GenS1{F0: (v4 % ((v4) | 1)), F1: (v2 | v2)}
                        if (18446744073709551615 == 2) {
                            var v11 int = fi7
                            var v12 bool = true
                            fmt.Println(v12)
                            _ = v11
                            _ = v12
                        } else {
                            v1 -= (v1 & v1)
                            switch (("" + "a") + ("foo bar" + ("" + (("\n\t" + "hello") + "foo bar")))) {
                            case "":
                                var v13 uint16 = 866
                                _ = v13
                            default:
                                v14 := GenS1{F0: v4, F1: v2}
                                _ = v14
                            }
                        }
                        v15 := v5[(fi9 + (fi7 | fi9))]
                        _ = fi9
                        _ = v10
                        _ = v15
                    }
                case "foo bar":
                    var v16 int32 = 42
                    _ = v16
                default:
                    var v17 int32 = -2147483648
                    _ = v17
                }
            case 0:
                v18 := GenS2{F0: (v2 == v2), F1: 2}
                _ = v18
            default:
                v19 := GenS1{F0: (v4 + v4), F1: 42}
                _ = v19
            }
            var v20 int16 = 32767
            _ = v20
        } else {
            if true {
                v21 := GenS1{F0: -9223372036854775808, F1: 540}
                _ = v21
            } else {
                var v22 int64 = 512
                _ = v22
            }
            if false {
                switch "\n\t" {
                case "\n\t":
                    v23 := GenS3{F0: -128, F1: 32767}
                    _ = v23
                case "":
                    for fi24 := 0; fi24 < 10; fi24++ {
                        v25 := GenS4{F0: -722, F1: 0, F2: v2}
                        switch (fi24 >> (((v2 >> ((18446744073709551615) & 63))) & 63)) {
                        case 9223372036854775807:
                            v2 -= 255
                        }
                        fmt.Println(fi24, v1, v1)
                        var v26 int8 = (v1 % ((1) | 1))
                        _ = fi24
                        _ = v25
                        _ = v26
                    }
                default:
                    if true {
                        var v27 uint32 = 288
                        var v28 float32 = 1.5
                        _ = v27
                        _ = v28
                    } else {
                        v29 := GenS5{F0: v4, F1: 42, F2: 9223372036854775807}
                        if (1 > 2147483647) {
                            if len(v8) >= 2 {
                                v8 = v8[1:len(v8)]
                            }
                            v4 = (v4 | (v4 << ((v2) & 63)))
                            v30 := v3[2147483647]
                            _ = v30
                        }
                        _ = v29
                    }
                }
                for fi31 := 0; fi31 < 5; fi31++ {
                    v32 := GenS4{F0: -31, F1: -9223372036854775808, F2: (v2 + 2)}
                    var v33 bool = (1e10 != 1e10)
                    _ = fi31
                    _ = v32
                    _ = v33
                }
                var v34 float32 = ((1e10 + 0.0) / (3.14159 / -2.25))
                _ = v34
            }
            fmt.Println(v1, v1)
        }
        for fi35 := 0; fi35 < 16; fi35++ {
            if true {
                fmt.Println(fi35, fi35)
                var v36 uint8 = 42
                if (!(((1.5 * (0.0 / 1.0)) <= (float64(1e10) - 0.0)) || true)) {
                    var v37 int64 = (v4 ^ v4)
                    v37 = v4
                    v37 -= v37
                    _ = v37
                } else {
                    var v38 float32 = 3.14159
                    _ = v38
                }
                _ = v36
            } else {
                var v39 int32 = 1
                switch (fi35 ^ fi35) {
                case 1:
                    _ = len(v3)
                case 32767:
                    v40 := GenS6{F0: 555, F1: v39}
                    _ = v40
                }
                v41 := GenS3{F0: 13, F1: fi35}
                _ = v39
                _ = v41
            }
            _ = fi35
        }
        _ = fi7
        _ = v8
    }
    for fi42 := 0; fi42 < 6; fi42++ {
        var v43 bool = true
        var v44 float32 = 1.5
        _ = fi42
        _ = v43
        _ = v44
    }
    fmt.Println(v1, v2, v1)
    var v45 float32 = -2.25
    for fi46 := 0; fi46 < 1; fi46++ {
        fmt.Println(v45, v4, v2)
        var v47 uint32 = 467
        v45 += -2.25
        _ = fi46
        _ = v47
    }
    var v48 uint8 = 128
    var v49 int64 = int64(v1)
    if false {
        v1 -= 0
        if true {
            var v50 int = -9223372036854775808
            var v51 string = "foo bar"
            _ = v50
            _ = v51
        } else {
            switch "\n\t" {
            case "":
                v3[608] = "hello"
            case "foo bar":
                for fi52 := 0; fi52 < 18; fi52++ {
                    for fi53 := 0; fi53 < 3; fi53++ {
                        var v54 float64 = 1.5
                        _ = fi53
                        _ = v54
                    }
                    for fi55 := 0; fi55 < 16; fi55++ {
                        v56 := GenS6{F0: 1, F1: 2147483647}
                        var v57 string = "a"
                        v4 = 86
                        v58 := []int{fi52, fi52}
                        v49 += (v4 & v4)
                        _ = fi55
                        _ = v56
                        _ = v57
                        _ = v58
                    }
                    _ = fi52
                }
            default:
                var v59 float32 = 0.1
                _ = v59
            }
            for fi60 := 0; fi60 < 9; fi60++ {
                var v61 int = fi60
                fmt.Println(v48)
                _ = fi60
                _ = v61
            }
            delete(v3, 65535)
        }
    } else {
        var v62 uint8 = 0
        var v63 float32 = float32(0.1)
        _ = v62
        _ = v63
    }
    switch "a" {
    case "hello":
        var v64 string = "foo bar"
        _ = v64
    case "a":
        for fi65 := 0; fi65 < 9; fi65++ {
            fmt.Println(v45, v2)
            var v66 uint64 = 4294967295
            var v67 float64 = float64(0.0)
            _ = fi65
            _ = v66
            _ = v67
        }
    }
    v45 -= v45
    var v68 bool = true
    if (("" != "") || v68) {
        for fi69 := 0; fi69 < 10; fi69++ {
            v48 -= v48
            if (796 <= 1) {
                for fi70 := 0; fi70 < 10; fi70++ {
                    if ((!true) && v68) {
                        v71 := GenS5{F0: (v4 / ((v49) | 1)), F1: 0, F2: -9223372036854775808}
                        _ = v71
                    } else {
                        if v68 {
                            v49 -= v49
                        } else {
                            v2 = 924
                            v72 := GenS5{F0: 0, F1: 0, F2: fi69}
                            _ = v72
                        }
                        if v68 {
                            v73 := []int{(fi70 % (((^(fi69))) | 1)), (fi70 ^ fi69), (-(fi70)), (^(fi69))}
                            v45 = (v45 + v45)
                            _ = v73
                        }
                        if v68 {
                            var v74 float32 = (v45 / (v45 * ((v45 * v45) + (v45 - ((v45 * (float32(0.1) / v45)) / v45)))))
                            _ = v74
                        }
                        if (0.1 == 1e10) {
                            var v75 int64 = (v4 % (((v49 % ((v49) | 1))) | 1))
                            var v76 uint32 = 65535
                            _ = len(v3)
                            _ = v75
                            _ = v76
                        }
                        switch (fi69 | fi70) {
                        case 56:
                            if (0 <= (v1 * v1)) {
                                switch (fi69 + fi69) {
                                case 4294967295:
                                    v68 = true
                                }
                                v45 += (v45 + ((3.14159 - v45) / (1e10 / float32(0.1))))
                                fmt.Println(v49, v68)
                                v77 := v3[2]
                                v78 := []int{(^(fi70)), fi70, fi69, (fi70 + fi69)}
                                _ = v77
                                _ = v78
                            } else {
                                var v79 uint32 = 2147483648
                                switch "a" {
                                case "a":
                                    var v80 string = "hello"
                                    _ = v80
                                case "\n\t":
                                    v48 = (v48 % ((v48) | 1))
                                case "":
                                    switch (fi70 * (-(fi70))) {
                                    case 127:
                                        var v81 float64 = 0.1
                                        _ = v81
                                    case -550:
                                        v82 := GenS7{F0: 511}
                                        _ = v82
                                    case 9223372036854775807:
                                        var v83 uint64 = 749
                                        _ = v83
                                    default:
                                        var v84 int8 = v1
                                        _ = v84
                                    }
                                default:
                                    var v85 float64 = ((3.14159 - 1.5) + 3.14159)
                                    _ = v85
                                }
                                _ = v79
                            }
                        default:
                            v45 = v45
                        }
                    }
                    _ = fi70
                }
                v86 := []int{}
                v48 -= v48
                _ = v86
            }
            v2 += (v2 & (^(v2)))
            _ = fi69
        }
        v87 := GenS2{F0: (("" + (((("hello" + "") + "hello") + "") + "a")) <= "foo bar"), F1: 0}
        if false {
            delete(v5, 4294967295)
        } else {
            v1 -= (v1 << (((v2 % ((v2) | 1))) & 63))
        }
        _ = v87
    } else {
        var v88 string = ("hello" + ("a" + "\n\t"))
        var v89 float32 = float32(1e10)
        switch 179 {
        case 38:
            v90 := []int{0, -2147483648}
            _ = v90
        case -9223372036854775808:
            if (!v68) {
                for fi91 := 0; fi91 < 13; fi91++ {
                    for fi92 := 0; fi92 < 7; fi92++ {
                        v89 -= (0.0 * (0.1 - v45))
                        v93 := GenS8{F0: (v68 && (32767 != 1)), F1: 654}
                        _ = fi92
                        _ = v93
                    }
                    for fi94 := 0; fi94 < 2; fi94++ {
                        v95 := GenS3{F0: 0, F1: (fi91 | 397)}
                        fmt.Println(v49, v45, v4)
                        v49 = -1
                        _ = fi94
                        _ = v95
                    }
                    _ = fi91
                }
                v96 := GenS3{F0: int8(v48), F1: -1}
                var v97 uint16 = 1
                _ = v96
                _ = v97
            } else {
                switch (("\n\t" + (v88 + "")) + "hello") {
                case "\n\t":
                    v5[32768] = v68
                default:
                    var v98 float64 = -2.25
                    _ = v98
                }
            }
        case 32768:
            switch 4294967295 {
            case 2147483647:
                if (v68 || (!true)) {
                    switch v88 {
                    case "foo bar":
                        v99 := GenS9{F0: v2, F1: 1}
                        _ = v99
                    default:
                        for fi100 := 0; fi100 < 14; fi100++ {
                            fmt.Println(v89, v4, v1)
                            v3[0] = ("\n\t" + "hello")
                            _ = fi100
                        }
                    }
                } else {
                    _ = len(v3)
                }
            case 32768:
                v49 += (v4 | v4)
            default:
                var v101 float32 = float32(1e10)
                _ = v101
            }
        }
        _ = v88
        _ = v89
    }
    if (127 >= v1) {
        var v102 int64 = v49
        _ = v102
    }
    delete(v5, 1)
    v49 = v4
    fmt.Println(v45, v49, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v68:", v68)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v48
    _ = v49
    _ = v5
    _ = v6
    _ = v68
}
