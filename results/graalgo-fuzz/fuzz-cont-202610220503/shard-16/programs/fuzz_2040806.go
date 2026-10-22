package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int16; F1 int32; F2 uint }
type GenS3 struct { F0 uint; F1 int32; F2 uint8 }
type GenS4 struct { F0 int64; F1 uint32; F2 int8 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 bool; F1 bool; F2 uint8 }
type GenS7 struct { F0 int32; F1 int16; F2 bool }
type GenS8 struct { F0 int16; F1 uint64; F2 int64 }
func genFn0(p0 int32) uint64 {
    var v7 bool = true
    fmt.Println(v7)
    v8 := []int{127, 524}
    v7 = (3.14159 <= 0.1)
    fmt.Println(v7, v7)
    _ = p0
    _ = v7
    _ = v8
    return 4294967295
}

func main() {
    v1 := []int{0}
    var v2 int8 = 100
    v3 := make(map[uint32]string)
    var v4 uint32 = 65535
    v5 := []string{(("\n\t" + "a") + ((("foo bar" + "a") + "a") + "\n\t")), "\n\t"}
    var v6 uint = 364
    v9 := GenS1{F0: 1}
    for fi10 := 0; fi10 < 14; fi10++ {
        if (false && false) {
            var v11 uint64 = 0
            v4 += 2
            _ = v11
        } else {
            switch (fi10 | (fi10 - (fi10 ^ (-(fi10))))) {
            case 0:
                var v12 uint64 = 594
                _ = v12
            }
            var v13 int64 = 227
            var v14 float32 = float32((1.5 - -2.25))
            _ = v13
            _ = v14
        }
        if (-2.25 > -2.25) {
            if true {
                for fi15 := 0; fi15 < 10; fi15++ {
                    switch (fi10 | fi10) {
                    case 0:
                        if (!true) {
                            switch "a" {
                            case "a":
                                var v16 int32 = -2147483648
                                _ = v16
                            case "hello":
                                if (((1e10 + 1.5) > ((3.14159 - 0.0) + (1e10 - (((1e10 + ((float64(0.0) / -2.25) - 0.0)) / 1.0) * -2.25)))) || (!true)) {
                                    var v17 bool = ("a" <= ("foo bar" + "hello"))
                                    if false {
                                        var v18 float64 = ((((0.1 + 0.1) * 1.5) / 1.0) * 0.1)
                                        var v19 float64 = float64(0.1)
                                        _ = v18
                                        _ = v19
                                    }
                                    v20 := GenS2{F0: 1000, F1: 1, F2: v6}
                                    _ = v17
                                    _ = v20
                                }
                            }
                            fmt.Println(v4)
                        } else {
                            var v21 uint8 = 1
                            _ = v21
                        }
                        if (4294967295 != 230) {
                            var v22 uint = (^(v6))
                            var v23 int32 = 1
                            _ = v22
                            _ = v23
                        }
                        var v24 int16 = -32768
                        _ = v24
                    case 9223372036854775807:
                        v6 = 4294967295
                    case -416:
                        if (-564 != 1000) {
                            v6 -= (v6 % ((v6) | 1))
                            var v25 string = ("" + "a")
                            v2 += v2
                            _ = v25
                        }
                    default:
                        fmt.Println(v2, v6)
                    }
                    if true {
                        v4 -= (v4 ^ v4)
                    } else {
                        v26 := GenS2{F0: 42, F1: 2147483647, F2: 2}
                        v27 := GenS3{F0: v6, F1: 42, F2: 255}
                        _ = v26
                        _ = v27
                    }
                    v4 = v4
                    _ = fi15
                }
                if true {
                    v2 -= -1
                }
                switch "hello" {
                case "a":
                    v2 = (v2 + (v2 - (v2 << (((^(v6))) & 63))))
                case "":
                    _ = len(v3)
                default:
                    var v28 uint32 = (v4 & v4)
                    _ = v28
                }
            } else {
                if (0 >= 0) {
                    v3[(v4 ^ 65535)] = "\n\t"
                    var v29 float64 = (1e10 / 3.14159)
                    var v30 uint = (v6 & 65535)
                    _ = v29
                    _ = v30
                } else {
                    if ((float64((((0.0 / -2.25) + -2.25) - -2.25)) - 0.1) >= float64(3.14159)) {
                        v4 += v4
                        v6 -= v6
                        var v31 uint8 = 30
                        _ = v31
                    }
                }
            }
            for fi32 := 0; fi32 < 20; fi32++ {
                fmt.Println(v6, v4, v2)
                v4 = v4
                v4 -= (^(v4))
                var v33 uint8 = 127
                v4 = 355
                _ = fi32
                _ = v33
            }
        }
        _ = fi10
    }
    switch ("hello" + ("a" + "a")) {
    case "":
        if (0 >= 1) {
            v2 += v2
        }
    case "foo bar":
        switch "foo bar" {
        case "foo bar":
            var v34 int = -9223372036854775808
            _ = v34
        case "":
            for fi35 := 0; fi35 < 4; fi35++ {
                switch "hello" {
                case "\n\t":
                    v36 := GenS1{F0: 128}
                    _ = v36
                case "":
                    fmt.Println(v4, v6, v4)
                }
                var v37 uint8 = 127
                _ = fi35
                _ = v37
            }
        case "\n\t":
            for fi38 := 0; fi38 < 4; fi38++ {
                var v39 uint16 = 32767
                var v40 bool = true
                _ = fi38
                _ = v39
                _ = v40
            }
        default:
            fmt.Println(v6)
        }
    default:
        v41 := GenS4{F0: -9223372036854775808, F1: v4, F2: (v2 % (((-(v2))) | 1))}
        _ = v41
    }
    fmt.Println(v6)
    if true {
        v4 = (v4 << ((v6) & 63))
        delete(v3, (v4 * v4))
    } else {
        for fi42 := 0; fi42 < 15; fi42++ {
            switch 1 {
            case 32767:
                var v43 int8 = v2
                _ = v43
            case 30:
                if false {
                    v4 = (v4 + 2147483647)
                } else {
                    var v44 float64 = 1.5
                    v44 -= v44
                    var v45 string = ("foo bar" + "a")
                    _ = v44
                    _ = v45
                }
            case -21:
                for fi46 := 0; fi46 < 1; fi46++ {
                    if true {
                        if (1 != v2) {
                            v1 = append(v1, 9223372036854775807)
                            fmt.Println(fi42, fi46, v4)
                            switch ("a" + "a") {
                            case "foo bar":
                                if true {
                                    v47 := GenS5{F0: (^(v4))}
                                    _ = len(v3)
                                    _ = v47
                                }
                            case "":
                                var v48 float64 = 1.5
                                _ = v48
                            }
                        } else {
                            fmt.Println(fi46, v4)
                        }
                        v6 = (v6 + (v6 / ((uint(v4)) | 1)))
                    }
                    v6 += v6
                    v2 += (v2 / ((v2) | 1))
                    _ = fi46
                }
            }
            _ = fi42
        }
    }
    v2 = (v2 * -109)
    v2 = v2
    fmt.Println(v4)
    for fi49 := 0; fi49 < 13; fi49++ {
        var v50 int32 = 42
        _ = fi49
        _ = v50
    }
    if false {
        if false {
            v51 := GenS5{F0: v4}
            var v52 float64 = -2.25
            switch "foo bar" {
            case "":
                var v53 uint = v6
                _ = v53
            case "foo bar":
                if true {
                    var v54 uint32 = v4
                    _ = v54
                }
            case "\n\t":
                v55 := GenS6{F0: (0.0 > float32(v52)), F1: ((6 != 1) || (0 >= -1)), F2: 183}
                _ = v55
            default:
                var v56 int16 = 647
                _ = v56
            }
            _ = v51
            _ = v52
        } else {
            v6 -= v6
            if true {
                v6 = v6
                var v57 int64 = 1
                _ = v57
            } else {
                var v58 int = 127
                _ = v58
            }
        }
        fmt.Println(v2, v2)
    } else {
        fmt.Println(v2, v6, v2)
        var v59 int8 = v2
        _ = v59
    }
    if false {
        v2 += (v2 ^ v2)
        v60 := GenS5{F0: (v4 ^ v4)}
        _ = v60
    } else {
        v4 -= v4
        v6 = (v6 | v6)
        switch "hello" {
        case "hello":
            _ = len(v3)
        case "a":
            v61 := GenS6{F0: false, F1: false, F2: 95}
            _ = v61
        case "":
            switch 32768 {
            case 127:
                switch (("" + "\n\t") + "hello") {
                case "a":
                    v62 := GenS7{F0: 0, F1: -32768, F2: false}
                    _ = v62
                case "":
                    if (0 == 128) {
                        if true {
                            v1 = append(v1, 896)
                        }
                    } else {
                        v2 -= (^(v2))
                        var v63 string = ""
                        _ = v63
                    }
                }
            case 4294967295:
                if (-307 < -9223372036854775808) {
                    fmt.Println(v4, v6, v4)
                    if true {
                        if false {
                            for fi64 := 0; fi64 < 8; fi64++ {
                                if true {
                                    for fi65 := 0; fi65 < 17; fi65++ {
                                        var v66 float32 = (0.0 * float32((float64((float32(0.1) - 1e10)) + (1e10 + 1e10))))
                                        _ = len(v3)
                                        fmt.Println(v4)
                                        _ = fi65
                                        _ = v66
                                    }
                                    for fi67 := 0; fi67 < 16; fi67++ {
                                        v2 = (v2 / (((v2 ^ (v2 - (v2 * (v2 & -31))))) | 1))
                                        switch "foo bar" {
                                        case "a":
                                            if true {
                                                if ((!false) || (!(!(v2 < v2)))) {
                                                    if false {
                                                        v6 = 2
                                                        if ((!false) || true) {
                                                            var v68 uint32 = v4
                                                            var v69 int8 = v2
                                                            var v70 int32 = 42
                                                            _ = v68
                                                            _ = v69
                                                            _ = v70
                                                        } else {
                                                            if false {
                                                                v71 := GenS5{F0: 1}
                                                                if (false || (false || ((!false) || false))) {
                                                                    var v72 int16 = -32768
                                                                    switch 42 {
                                                                    case 465:
                                                                        delete(v3, v4)
                                                                    }
                                                                    _ = v72
                                                                }
                                                                v6 = (v6 + v6)
                                                                _ = v71
                                                            }
                                                            if false {
                                                                v4 = 65535
                                                                var v73 int64 = 700
                                                                v3[(v4 * 42)] = ("hello" + "")
                                                                _ = v73
                                                            }
                                                        }
                                                        v4 = v4
                                                        var v74 bool = (1 < 1)
                                                        _ = v74
                                                    }
                                                } else {
                                                    var v75 float32 = 3.14159
                                                    _ = v75
                                                }
                                            }
                                        case "":
                                            v5 = append(v5, (("" + "\n\t") + ("\n\t" + (("foo bar" + ("a" + ((("" + "foo bar") + "hello") + ""))) + ("hello" + "a")))))
                                        case "foo bar":
                                            v4 = v4
                                        default:
                                            switch "foo bar" {
                                            case "a":
                                                v6 = 4294967295
                                            }
                                        }
                                        v76 := v3[(^(v4))]
                                        _ = fi67
                                        _ = v76
                                    }
                                    v77 := GenS3{F0: v6, F1: 0, F2: 127}
                                    _ = v77
                                } else {
                                    _ = len(v3)
                                    if true {
                                        var v78 uint = (v6 >> (((^(v6))) & 63))
                                        _ = v78
                                    } else {
                                        v3[42] = "hello"
                                        var v79 float64 = (-2.25 - -2.25)
                                        var v80 uint8 = 209
                                        _ = v79
                                        _ = v80
                                    }
                                }
                                if false {
                                    var v81 uint16 = uint16(v4)
                                    _ = v81
                                } else {
                                    for fi82 := 0; fi82 < 2; fi82++ {
                                        v2 = (v2 % ((v2) | 1))
                                        v6 = v6
                                        _ = len(v3)
                                        _ = fi82
                                    }
                                    for fi83 := 0; fi83 < 2; fi83++ {
                                        fmt.Println(v2)
                                        var v84 string = "foo bar"
                                        _ = fi83
                                        _ = v84
                                    }
                                    switch 4294967295 {
                                    case 127:
                                        var v85 int = (fi64 + fi64)
                                        _ = v85
                                    }
                                }
                                v4 = v4
                                _ = fi64
                            }
                            v6 -= (^(v6))
                            if len(v1) > 0 {
                                v1[0] = 4294967295
                            }
                        } else {
                            if true {
                                var v86 uint16 = 65535
                                _ = v86
                            }
                            var v87 uint16 = 111
                            fmt.Println(v4, v4, v87)
                            _ = v87
                        }
                    } else {
                        if false {
                            if (-1 != int32(v2)) {
                                switch 1 {
                                case 548:
                                    v2 = v2
                                case -965:
                                    switch "foo bar" {
                                    case "hello":
                                        for fi88 := 0; fi88 < 10; fi88++ {
                                            v2 = (v2 >> ((18446744073709551615) & 63))
                                            _ = fi88
                                        }
                                    case "\n\t":
                                        var v89 int32 = 618
                                        _ = v89
                                    case "a":
                                        if true {
                                            var v90 bool = ((v4 << (((v6 >> (((v6 + 2)) & 63))) & 63)) > v4)
                                            _ = v90
                                        }
                                    }
                                case -743:
                                    v91 := GenS8{F0: 42, F1: 2, F2: -537}
                                    _ = v91
                                default:
                                    var v92 int16 = -678
                                    _ = v92
                                }
                            }
                        }
                        fmt.Println(v6, v6)
                    }
                    fmt.Println(v2)
                }
            default:
                v6 -= (v6 & (v6 >> ((v6) & 63)))
            }
        default:
            var v93 float32 = (1e10 / 1.0)
            _ = v93
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
