package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 bool; F1 int; F2 int32 }
type GenS5 struct { F0 int64; F1 uint }
type GenS6 struct { F0 uint16; F1 int16; F2 uint64 }
type GenS7 struct { F0 int64; F1 uint; F2 int8 }
type GenS8 struct { F0 int8 }
func genFn0(p0 uint8) int8 {
    p0 = (p0 / (((p0 - (p0 >> ((0) & 63)))) | 1))
    if ((94 >= 42) || (600 >= 667)) {
        p0 = (p0 - 248)
    } else {
        p0 -= (p0 % (((p0 << ((81) & 63))) | 1))
        if false {
            p0 += p0
            switch "hello" {
            case "\n\t":
                p0 -= (p0 >> ((1) & 63))
            default:
                var v6 int64 = 42
                _ = v6
            }
            if (32768 < -9223372036854775808) {
                var v7 uint16 = 988
                for fi8 := 0; fi8 < 5; fi8++ {
                    if (!false) {
                        v7 = uint16(p0)
                        var v9 uint64 = 2
                        _ = v9
                    } else {
                        var v10 int8 = 1
                        v7 = 69
                        _ = v10
                    }
                    p0 = 2
                    _ = fi8
                }
                _ = v7
            } else {
                p0 -= (p0 % ((0) | 1))
            }
        }
    }
    var v11 uint32 = 4294967295
    _ = p0
    _ = v11
    return -64
}

func main() {
    var v1 float64 = (0.0 + (0.0 + 0.0))
    v2 := make(map[int]string)
    var v3 int16 = 1
    var v4 uint = 255
    var v5 uint32 = 0
    if (!((false && (v1 != (3.14159 / float64((3.14159 * (((3.14159 - 1.5) - float32(3.14159)) / 1e10)))))) || false)) {
        fmt.Println(v5, v4, v5)
        switch 127 {
        case 574:
            var v12 int32 = 2147483647
            _ = v12
        default:
            var v13 uint = (v4 - (v4 * v4))
            _ = v13
        }
        for fi14 := 0; fi14 < 20; fi14++ {
            v4 = v4
            _ = fi14
        }
    } else {
        var v15 int8 = -5
        var v16 int = 32768
        _ = v15
        _ = v16
    }
    v17 := GenS1{F0: genFn0(128)}
    for fi18 := 0; fi18 < 7; fi18++ {
        var v19 string = "a"
        for fi20 := 0; fi20 < 6; fi20++ {
            v4 -= v4
            var v21 int8 = 127
            _ = fi20
            _ = v21
        }
        if true {
            v19 = (((("" + "a") + "foo bar") + v19) + v19)
        }
        _ = fi18
        _ = v19
    }
    if true {
        var v22 uint = v4
        var v23 uint64 = 1
        switch -47 {
        case -9223372036854775808:
            v24 := GenS1{F0: 127}
            _ = v24
        case 0:
            var v25 string = ((("foo bar" + "foo bar") + "") + (((("\n\t" + "hello") + ("hello" + ((("foo bar" + "\n\t") + "") + ""))) + ("" + "hello")) + (("" + "hello") + "foo bar")))
            _ = v25
        case 4294967295:
            switch 4294967295 {
            case -632:
                for fi26 := 0; fi26 < 15; fi26++ {
                    for fi27 := 0; fi27 < 1; fi27++ {
                        if (v23 >= (v23 >> (((^(v22))) & 63))) {
                            fmt.Println(fi26)
                            fmt.Println(v1, v22, fi26)
                            v1 = v1
                        }
                        var v28 int16 = 0
                        var v29 string = "hello"
                        _ = fi27
                        _ = v28
                        _ = v29
                    }
                    switch ("" + "\n\t") {
                    case "hello":
                        switch fi26 {
                        case 83:
                            v23 = (v23 + 660)
                        }
                    case "a":
                        for fi30 := 0; fi30 < 15; fi30++ {
                            v2[fi30] = ""
                            v31 := []int{fi26, (^(fi30)), (fi26 + (fi26 & (fi30 | -660)))}
                            var v32 float64 = (float64((0.1 - -2.25)) + (v1 / (v1 * (v1 * (((-2.25 * v1) - float64(((1.5 - 1.5) + 1e10))) * v1)))))
                            fmt.Println(fi30)
                            v22 += v4
                            if true {
                                var v33 int8 = int8(v22)
                                _ = v33
                            }
                            v2[128] = "a"
                            _ = fi30
                            _ = v31
                            _ = v32
                        }
                    case "foo bar":
                        v34 := GenS1{F0: 0}
                        _ = v34
                    default:
                        v35 := GenS2{F0: uint16(v5)}
                        _ = v35
                    }
                    _ = fi26
                }
            case 127:
                for fi36 := 0; fi36 < 16; fi36++ {
                    if ((v1 >= 0.1) && true) {
                        v37 := []int{(fi36 + int(v3))}
                        for fi38 := 0; fi38 < 18; fi38++ {
                            v1 = (-2.25 + (((((v1 * v1) * v1) + v1) - 1.5) * (((v1 + v1) / 3.14159) + 1e10)))
                            if len(v37) >= 2 {
                                v37 = v37[1:len(v37)]
                            }
                            _ = fi38
                        }
                        v3 = (-(v3))
                        _ = v37
                    } else {
                        v39 := GenS1{F0: 100}
                        _ = v39
                    }
                    _ = fi36
                }
            case 0:
                fmt.Println(v1)
            }
        }
        _ = v22
        _ = v23
    }
    v40 := GenS3{F0: v3}
    var v41 int = -2147483648
    var v42 uint64 = 220
    for fi43 := 0; fi43 < 16; fi43++ {
        var v44 float64 = float64((1e10 * (0.0 - 1.5)))
        var v45 int8 = 42
        if true {
            var v46 int = v41
            _ = v46
        } else {
            switch "" {
            case "\n\t":
                if (!true) {
                    v47 := []int{-411, (v41 | fi43)}
                    for fi48 := 0; fi48 < 1; fi48++ {
                        var v49 int8 = genFn0(0)
                        v44 += 1e10
                        var v50 int32 = 0
                        _ = fi48
                        _ = v49
                        _ = v50
                    }
                    switch (v41 - (^(v41))) {
                    case 9223372036854775807:
                        fmt.Println(fi43, v5)
                    case 0:
                        var v51 float32 = -2.25
                        _ = v51
                    default:
                        v52 := GenS4{F0: (!true), F1: fi43, F2: 193}
                        _ = v52
                    }
                    _ = v47
                }
            case "foo bar":
                var v53 uint = (v4 ^ 4294967295)
                _ = v53
            case "":
                v1 -= (-2.25 / 1e10)
            default:
                v4 = (v4 % (((v4 - v4)) | 1))
            }
            var v54 int8 = v45
            if (316 == 938) {
                v5 += 2147483647
                if (-9223372036854775808 == 9223372036854775807) {
                    v55 := []int{4294967295, (fi43 | 55), fi43}
                    if (1 == 65535) {
                        if false {
                            for fi56 := 0; fi56 < 13; fi56++ {
                                _ = len(v2)
                                _ = fi56
                            }
                            for fi57 := 0; fi57 < 3; fi57++ {
                                v44 = ((v1 / (v44 / v1)) - (3.14159 / 3.14159))
                                _ = len(v2)
                                if (-2.25 <= (3.14159 + 1e10)) {
                                    if false {
                                        v1 = float64(3.14159)
                                        v58 := GenS2{F0: 42}
                                        v59 := GenS2{F0: 2}
                                        _ = v58
                                        _ = v59
                                    }
                                    delete(v2, 127)
                                    var v60 int64 = -9223372036854775808
                                    _ = v60
                                } else {
                                    v61 := GenS5{F0: int64(v41), F1: v4}
                                    var v62 int8 = (v54 << (((v4 % ((154) | 1))) & 63))
                                    _ = v61
                                    _ = v62
                                }
                                _ = fi57
                            }
                        } else {
                            if len(v55) > 0 {
                                v55[0] = v41
                            }
                            v42 = v42
                            var v63 uint64 = 1
                            _ = v63
                        }
                        v64 := v2[(-(v41))]
                        if ((v42 >> (((^(v4))) & 63)) < v42) {
                            fmt.Println(v54)
                        } else {
                            var v65 int8 = (v45 ^ v45)
                            _ = len(v55)
                            for fi66 := 0; fi66 < 19; fi66++ {
                                if (((v44 / (((float64(((-2.25 / -2.25) / ((float32(float64(((3.14159 / 3.14159) - 1e10))) / 3.14159) + float32((v44 + float64(-2.25)))))) / 0.1) / -2.25) - v44)) - ((((v44 * v1) + 0.0) * 1e10) / v1)) != v44) {
                                    fmt.Println(v44, v4)
                                }
                                v1 = 0.1
                                _ = fi66
                            }
                            _ = v65
                        }
                        _ = v64
                    } else {
                        var v67 bool = (true && false)
                        _ = v67
                    }
                    _ = v55
                }
                v44 = (v1 - float64(1.5))
            } else {
                switch (v41 & v41) {
                case -724:
                    var v68 int64 = -1
                    _ = v68
                case -262:
                    if (128 > 1) {
                        fmt.Println(v45, v1, v42)
                        var v69 uint = v4
                        if false {
                            v4 += 368
                        }
                        _ = v69
                    }
                case -9223372036854775808:
                    _ = len(v2)
                }
            }
            _ = v54
        }
        _ = fi43
        _ = v44
        _ = v45
    }
    if ("" <= "a") {
        var v70 uint16 = 32767
        _ = v70
    } else {
        for fi71 := 0; fi71 < 6; fi71++ {
            for fi72 := 0; fi72 < 16; fi72++ {
                var v73 float64 = v1
                v74 := GenS1{F0: -126}
                v75 := []int{v41, 42}
                v41 -= 127
                _ = fi72
                _ = v73
                _ = v74
                _ = v75
            }
            if (!false) {
                var v76 uint8 = 128
                _ = v76
            } else {
                switch "hello" {
                case "a":
                    var v77 uint64 = v42
                    _ = v77
                case "":
                    v78 := GenS3{F0: (^(v3))}
                    _ = v78
                default:
                    v79 := GenS6{F0: 1, F1: v3, F2: v42}
                    _ = v79
                }
                for fi80 := 0; fi80 < 6; fi80++ {
                    v81 := []int{int(v5)}
                    _ = fi80
                    _ = v81
                }
                var v82 int32 = 0
                _ = v82
            }
            v42 -= 2
            _ = fi71
        }
    }
    v83 := GenS2{F0: 255}
    v42 -= v42
    v84 := GenS7{F0: 0, F1: v4, F2: -1}
    if (!true) {
        if true {
            v85 := GenS3{F0: v3}
            var v86 uint16 = 32767
            fmt.Println(v5, v41)
            _ = v85
            _ = v86
        } else {
            var v87 float64 = float64(float32(v1))
            _ = v87
        }
        switch (v41 + (v41 - (^(v41)))) {
        case 32768:
            var v88 int64 = -155
            _ = v88
        case 14:
            if (v42 >= v42) {
                var v89 float64 = float64(1.5)
                _ = v89
            } else {
                switch (v41 % ((int(v5)) | 1)) {
                case 4294967295:
                    v90 := GenS6{F0: 899, F1: v3, F2: v42}
                    _ = v90
                case 32767:
                    _ = len(v2)
                }
                v42 += 887
            }
        }
        v91 := GenS8{F0: 85}
        _ = v91
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v5
    _ = v83
    _ = v84
}
