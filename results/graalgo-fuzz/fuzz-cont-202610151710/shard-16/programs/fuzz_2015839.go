package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32 }
type GenS2 struct { F0 int32; F1 int32 }
type GenS3 struct { F0 uint32; F1 uint; F2 int8 }
type GenS4 struct { F0 uint8; F1 uint32 }
type GenS5 struct { F0 uint16; F1 uint32 }
type GenS6 struct { F0 int64; F1 uint32; F2 bool }
type GenS7 struct { F0 uint16; F1 int16 }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 bool; F1 string }
type GenS10 struct { F0 int64 }
type GenS11 struct { F0 uint32; F1 int16; F2 int16 }
type GenS12 struct { F0 int32; F1 uint32 }
type GenS13 struct { F0 int8; F1 int8 }
type GenS14 struct { F0 uint }
func genFn0(p0 uint, p1 string) uint32 {
    p0 = 4294967295
    fmt.Println(p1)
    _ = p0
    _ = p1
    return 1
}

func main() {
    var v1 uint16 = 255
    v2 := []uint32{}
    var v3 int16 = 1000
    v4 := []int{}
    var v5 int = 1
    var v6 int64 = 1
    v2 = append(v2, 1)
    var v7 string = ("foo bar" + "a")
    if (false && true) {
        for fi8 := 0; fi8 < 7; fi8++ {
            v9 := GenS1{F0: v6, F1: 65535}
            _ = len(v2)
            switch "" {
            case "hello":
                fmt.Println(v3)
            default:
                if true {
                    v3 += v3
                    _ = len(v4)
                    if (42 != 2) {
                        var v10 uint8 = 128
                        _ = v10
                    } else {
                        _ = len(v4)
                        v1 = (v1 - (v1 - (^(v1))))
                        for fi11 := 0; fi11 < 1; fi11++ {
                            v1 += 1
                            if false {
                                v12 := GenS2{F0: 1, F1: -1}
                                _ = v12
                            } else {
                                if (!true) {
                                    if (float32((((0.0 + 0.1) + 1e10) - 0.0)) >= 1e10) {
                                        if false {
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                        } else {
                                            var v13 string = v7
                                            _ = v13
                                        }
                                        var v14 int32 = -931
                                        v15 := GenS1{F0: v6, F1: 1}
                                        _ = v14
                                        _ = v15
                                    } else {
                                        v16 := GenS2{F0: -1, F1: 42}
                                        v5 += (^(v5))
                                        _ = v16
                                    }
                                } else {
                                    switch int(v3) {
                                    case 128:
                                        var v17 int16 = v3
                                        _ = v17
                                    case 32767:
                                        if false {
                                            var v18 int64 = (v6 & v6)
                                            v4 = append(v4, 4294967295)
                                            if false {
                                                switch (fi8 * (-(fi8))) {
                                                case 2147483647:
                                                    v19 := GenS3{F0: 2, F1: 731, F2: 1}
                                                    _ = v19
                                                default:
                                                    switch v7 {
                                                    case "a":
                                                        if (!false) {
                                                            fmt.Println(fi8, v7)
                                                        }
                                                    case "\n\t":
                                                        var v20 uint = 987
                                                        _ = v20
                                                    case "foo bar":
                                                        fmt.Println(fi11, v18)
                                                        v21 := GenS1{F0: v6, F1: uint32(v1)}
                                                        v6 = int64(v1)
                                                        if false {
                                                            var v22 uint32 = genFn0(73, (((v7 + v7) + v7) + "\n\t"))
                                                            fmt.Println(v18)
                                                            _ = v22
                                                        }
                                                        if len(v4) > 0 {
                                                            v4[0] = 32768
                                                        }
                                                        fmt.Println(v6, fi8, v3)
                                                        _ = len(v2)
                                                        var v23 int64 = (v18 + (v18 / ((9223372036854775807) | 1)))
                                                        _ = v21
                                                        _ = v23
                                                    default:
                                                        var v24 int16 = -1
                                                        _ = v24
                                                    }
                                                    v1 += (v1 >> ((2) & 63))
                                                    v3 = int16(v18)
                                                }
                                            }
                                            _ = v18
                                        } else {
                                            var v25 int = (-(v5))
                                            var v26 int16 = v3
                                            v1 += v1
                                            if ((0.0 < (3.14159 * -2.25)) || true) {
                                                switch v7 {
                                                case "hello":
                                                    var v27 int8 = 42
                                                    _ = v27
                                                case "":
                                                    v7 = (v7 + "\n\t")
                                                case "foo bar":
                                                    fmt.Println(fi11, v25, v25)
                                                default:
                                                    if (false || true) {
                                                        if len(v4) >= 2 {
                                                            v4 = v4[1:len(v4)]
                                                        }
                                                        v25 += 32768
                                                        switch ("a" + v7) {
                                                        case "foo bar":
                                                            v4 = append(v4, (fi8 | (v25 >> ((42) & 63))))
                                                        default:
                                                            var v28 int32 = 2147483647
                                                            _ = v28
                                                        }
                                                    } else {
                                                        v6 += (v6 + v6)
                                                        fmt.Println(fi11, v3)
                                                        v26 = v3
                                                    }
                                                }
                                                v26 += v26
                                                var v29 float64 = 1e10
                                                var v30 float64 = (v29 + 1e10)
                                                if (false || (42 == 0)) {
                                                    v1 -= (v1 >> ((42) & 63))
                                                    switch v7 {
                                                    case "hello":
                                                        v6 -= (v6 ^ -1)
                                                    case "a":
                                                        var v31 uint64 = 1
                                                        _ = v31
                                                    case "\n\t":
                                                        v5 -= 127
                                                    }
                                                    v32 := GenS1{F0: -605, F1: 400}
                                                    _ = v32
                                                } else {
                                                    if true {
                                                        var v33 int32 = 1
                                                        _ = v33
                                                    }
                                                    if len(v4) > 0 {
                                                        v4[0] = fi11
                                                    }
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                }
                                                var v34 uint64 = 2
                                                v35 := GenS2{F0: 42, F1: -1}
                                                _ = v29
                                                _ = v30
                                                _ = v34
                                                _ = v35
                                            }
                                            var v36 float64 = 1e10
                                            _ = v25
                                            _ = v26
                                            _ = v36
                                        }
                                        if len(v4) > 0 {
                                            v4[0] = 42
                                        }
                                        var v37 int16 = (v3 * (v3 >> ((18446744073709551615) & 63)))
                                        _ = v37
                                    }
                                    v6 += -1
                                    if (-128 <= 87) {
                                        var v38 int16 = int16(fi11)
                                        var v39 float64 = (1.5 + 3.14159)
                                        _ = v38
                                        _ = v39
                                    } else {
                                        var v40 uint64 = 42
                                        v40 = (v40 << ((4294967295) & 63))
                                        _ = v40
                                    }
                                }
                                v41 := GenS4{F0: 2, F1: genFn0(595, v7)}
                                v42 := GenS5{F0: (v1 * 402), F1: genFn0(255, v7)}
                                _ = v41
                                _ = v42
                            }
                            _ = fi11
                        }
                    }
                }
            }
            _ = fi8
            _ = v9
        }
        switch (((v7 + ("\n\t" + v7)) + "\n\t") + v7) {
        case "a":
            v43 := GenS6{F0: (v6 | (v6 >> ((2) & 63))), F1: 592, F2: true}
            _ = v43
        default:
            fmt.Println(v7, v5)
        }
        v5 = (v5 & v5)
    }
    v44 := GenS3{F0: 482, F1: 42, F2: 19}
    if (!(true || true)) {
        var v45 float64 = (-2.25 - 3.14159)
        v5 = v5
        _ = v45
    } else {
        fmt.Println(v3, v3, v5)
    }
    switch v7 {
    case "a":
        var v46 bool = true
        _ = v46
    case "hello":
        if (!false) {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        } else {
            fmt.Println(v5)
            if true {
                v1 -= (v1 << ((2) & 63))
            } else {
                switch v7 {
                case "foo bar":
                    if true {
                        v1 += (v1 + v1)
                    }
                case "a":
                    v47 := GenS7{F0: (v1 / (((v1 / ((v1) | 1))) | 1)), F1: v3}
                    _ = v47
                case "hello":
                    v5 += (v5 * 42)
                default:
                    switch ((v7 + ((v7 + v7) + v7)) + v7) {
                    case "hello":
                        var v48 bool = (true || (!true))
                        _ = v48
                    case "":
                        if ((-2147483648 < 279) && false) {
                            var v49 uint32 = genFn0(887, v7)
                            for fi50 := 0; fi50 < 16; fi50++ {
                                var v51 uint64 = 18446744073709551615
                                for fi52 := 0; fi52 < 14; fi52++ {
                                    switch -2147483648 {
                                    case 128:
                                        var v53 int32 = 1
                                        _ = v53
                                    }
                                    v7 = ("foo bar" + "")
                                    if (28 >= (v51 << ((2) & 63))) {
                                        v54 := GenS1{F0: 42, F1: (v49 + v49)}
                                        _ = v54
                                    }
                                    _ = fi52
                                }
                                _ = fi50
                                _ = v51
                            }
                            _ = v49
                        } else {
                            var v55 int16 = 32767
                            if true {
                                fmt.Println(v6, v55, v6)
                                v56 := GenS5{F0: (^(v1)), F1: 0}
                                v7 = "foo bar"
                                _ = v56
                            } else {
                                v57 := GenS8{F0: 42}
                                _ = v57
                            }
                            _ = v55
                        }
                    case "foo bar":
                        if true {
                            switch v7 {
                            case "":
                                for fi58 := 0; fi58 < 13; fi58++ {
                                    v3 = (v3 / (((v3 % ((140) | 1))) | 1))
                                    v5 += 0
                                    _ = fi58
                                }
                            case "foo bar":
                                if (!((v6 == v6) || false)) {
                                    var v59 uint8 = 94
                                    _ = v59
                                } else {
                                    var v60 uint64 = 553
                                    if (!(int8(v5) != -1)) {
                                        if (-1 == int32(v60)) {
                                            v6 = int64(v5)
                                            fmt.Println(v60, v5, v1)
                                            v5 = v5
                                        }
                                        v1 = v1
                                        if true {
                                            var v61 uint16 = 25
                                            var v62 float64 = 1.5
                                            switch v5 {
                                            case 0:
                                                switch (-(v5)) {
                                                case -748:
                                                    v63 := GenS9{F0: false, F1: (("foo bar" + "hello") + v7)}
                                                    _ = v63
                                                case -2147483648:
                                                    var v64 int = (v5 + v5)
                                                    _ = v64
                                                default:
                                                    var v65 float64 = 1.5
                                                    _ = v65
                                                }
                                            default:
                                                v3 += -264
                                            }
                                            _ = v61
                                            _ = v62
                                        } else {
                                            for fi66 := 0; fi66 < 2; fi66++ {
                                                v4 = append(v4, (v5 * (v5 % (((v5 % ((1) | 1))) | 1))))
                                                for fi67 := 0; fi67 < 17; fi67++ {
                                                    _ = len(v2)
                                                    v68 := GenS10{F0: v6}
                                                    _ = fi67
                                                    _ = v68
                                                }
                                                _ = fi66
                                            }
                                            var v69 float64 = (0.0 + 1e10)
                                            var v70 string = (v7 + v7)
                                            _ = v69
                                            _ = v70
                                        }
                                    } else {
                                        switch (v7 + "\n\t") {
                                        case "a":
                                            v71 := GenS2{F0: 2147483647, F1: 2147483647}
                                            _ = v71
                                        case "foo bar":
                                            v7 = "\n\t"
                                        default:
                                            for fi72 := 0; fi72 < 17; fi72++ {
                                                v73 := GenS1{F0: (v6 & 772), F1: 42}
                                                var v74 uint = 255
                                                v75 := GenS11{F0: genFn0(v74, v7), F1: (-(v3)), F2: v3}
                                                _ = fi72
                                                _ = v73
                                                _ = v74
                                                _ = v75
                                            }
                                        }
                                        switch (v5 + 4294967295) {
                                        case 127:
                                            for fi76 := 0; fi76 < 9; fi76++ {
                                                for fi77 := 0; fi77 < 7; fi77++ {
                                                    var v78 bool = (2 < 592)
                                                    fmt.Println(v3, fi76, v1)
                                                    var v79 int32 = 0
                                                    _ = fi77
                                                    _ = v78
                                                    _ = v79
                                                }
                                                v80 := GenS7{F0: v1, F1: (v3 + (v3 ^ v3))}
                                                _ = fi76
                                                _ = v80
                                            }
                                        default:
                                            v60 += (v60 | (v60 - 201))
                                        }
                                        _ = len(v4)
                                    }
                                    v7 = v7
                                    _ = v60
                                }
                            case "a":
                                for fi81 := 0; fi81 < 2; fi81++ {
                                    _ = len(v2)
                                    _ = fi81
                                }
                            default:
                                for fi82 := 0; fi82 < 4; fi82++ {
                                    v83 := GenS12{F0: 42, F1: genFn0(18446744073709551615, v7)}
                                    _ = fi82
                                    _ = v83
                                }
                            }
                            if (false || (true || ((((float32(0.1) - 3.14159) - 0.1) / 3.14159) >= 0.0))) {
                                fmt.Println(v6, v3)
                                v7 = v7
                                var v84 uint32 = genFn0(65535, v7)
                                _ = v84
                            } else {
                                fmt.Println(v6)
                                var v85 float64 = ((0.0 / (float64(1e10) - 3.14159)) + 3.14159)
                                _ = v85
                            }
                            switch v5 {
                            case 4294967295:
                                v86 := GenS5{F0: 42, F1: 2}
                                _ = v86
                            case 42:
                                if (1 > 42) {
                                    _ = len(v2)
                                    switch v7 {
                                    case "foo bar":
                                        v5 = v5
                                    case "hello":
                                        var v87 uint = 787
                                        _ = v87
                                    case "a":
                                        for fi88 := 0; fi88 < 6; fi88++ {
                                            var v89 int64 = v6
                                            var v90 uint64 = 888
                                            _ = fi88
                                            _ = v89
                                            _ = v90
                                        }
                                    }
                                    v2 = append(v2, 294)
                                }
                            case 0:
                                var v91 uint16 = (^(v1))
                                _ = v91
                            }
                        }
                    }
                }
            }
            for fi92 := 0; fi92 < 8; fi92++ {
                if false {
                    fmt.Println(v3, fi92, fi92)
                }
                switch v5 {
                case 32767:
                    var v93 int = (v5 ^ v5)
                    _ = v93
                }
                fmt.Println(v7, fi92, fi92)
                _ = fi92
            }
        }
    case "":
        var v94 string = "hello"
        _ = v94
    }
    v6 = -9223372036854775808
    if false {
        if false {
            v1 -= 2
        } else {
            v95 := GenS13{F0: -35, F1: -128}
            var v96 uint = 42
            for fi97 := 0; fi97 < 4; fi97++ {
                var v98 int16 = (v3 - v3)
                _ = fi97
                _ = v98
            }
            _ = v95
            _ = v96
        }
        v99 := GenS14{F0: 2}
        v6 += v6
        _ = v99
    }
    v100 := GenS3{F0: 1, F1: 255, F2: 0}
    switch v7 {
    case "a":
        for fi101 := 0; fi101 < 6; fi101++ {
            v102 := GenS4{F0: 42, F1: uint32(v1)}
            var v103 string = (v7 + ("hello" + v7))
            _ = fi101
            _ = v102
            _ = v103
        }
    case "":
        v7 = ((v7 + ("\n\t" + (v7 + v7))) + v7)
    }
    fmt.Println(v6)
    if len(v2) > 0 {
        v2[0] = genFn0(623, ("\n\t" + "foo bar"))
    }
    for fi104 := 0; fi104 < 11; fi104++ {
        v2 = append(v2, 301)
        if ((-153 > v3) || (!(int32(fi104) >= -865))) {
            v1 = (v1 | v1)
            v5 -= fi104
        } else {
            switch fi104 {
            case 918:
                var v105 int8 = -92
                _ = v105
            case -438:
                var v106 int8 = 0
                _ = v106
            case -2147483648:
                if true {
                    var v107 uint64 = 1
                    v5 = fi104
                    _ = v107
                }
            }
            fmt.Println(fi104)
            v6 = v6
        }
        _ = fi104
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v100
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v5
    _ = v6
    _ = v7
}
