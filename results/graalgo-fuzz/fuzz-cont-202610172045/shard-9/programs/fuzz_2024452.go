package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint; F2 uint8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int; F1 int8 }
type GenS4 struct { F0 uint; F1 int8; F2 bool }
type GenS5 struct { F0 bool; F1 int64 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 uint8; F1 int16; F2 uint16 }
type GenS9 struct { F0 int; F1 uint32 }
type GenS10 struct { F0 uint8 }
type GenS11 struct { F0 int32 }
type GenS12 struct { F0 uint8; F1 int32; F2 uint16 }
type GenS13 struct { F0 int8; F1 int64; F2 int16 }
func genFn0(p0 uint16, p1 int16, p2 int32) int {
    v5 := []int{-117}
    fmt.Println(p2, p0, p2)
    var v6 string = "\n\t"
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return -182
}

func main() {
    v1 := []uint32{705}
    v2 := []int{32767, -2147483648, 32767}
    v3 := []uint32{}
    v4 := []uint32{2147483647, 1, 1}
    switch 725 {
    case 256:
        if len(v3) > 0 {
            v3[0] = 1
        }
    case -52:
        for fi7 := 0; fi7 < 1; fi7++ {
            v8 := GenS1{F0: 42, F1: 1, F2: 193}
            _ = fi7
            _ = v8
        }
    default:
        v9 := GenS2{F0: 0}
        _ = v9
    }
    var v10 int64 = 301
    if (!(v10 >= (v10 * (v10 & v10)))) {
        var v11 uint64 = 946
        if ((true && (127 >= -1)) && true) {
            var v12 int32 = int32(v11)
            switch 0 {
            case 0:
                var v13 int16 = 32767
                _ = v13
            case 128:
                for fi14 := 0; fi14 < 11; fi14++ {
                    v4 = append(v4, 2)
                    switch ("hello" + "\n\t") {
                    case "foo bar":
                        v12 = 1
                    }
                    _ = fi14
                }
            case 609:
                switch 1 {
                case 518:
                    var v15 int = -303
                    _ = v15
                case -549:
                    v10 += 203
                case 255:
                    v12 = (v12 & -167)
                default:
                    fmt.Println(v11, v10, v12)
                }
            }
            _ = v12
        } else {
            var v16 int64 = 1
            _ = v16
        }
        v10 -= v10
        _ = v11
    } else {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        v17 := GenS3{F0: -9223372036854775808, F1: 127}
        _ = v17
    }
    if (false || (!true)) {
        var v18 uint32 = 0
        if len(v1) > 0 {
            v1[0] = v18
        }
        for fi19 := 0; fi19 < 19; fi19++ {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            if (0 >= 216) {
                if false {
                    v18 -= 595
                    switch "\n\t" {
                    case "foo bar":
                        for fi20 := 0; fi20 < 14; fi20++ {
                            v10 -= -9223372036854775808
                            v18 = 42
                            var v21 int8 = 127
                            _ = fi20
                            _ = v21
                        }
                    case "":
                        v22 := GenS4{F0: 402, F1: 0, F2: true}
                        _ = v22
                    case "a":
                        v23 := GenS3{F0: (fi19 | (fi19 ^ -9223372036854775808)), F1: 50}
                        _ = v23
                    }
                    var v24 uint8 = 162
                    _ = v24
                }
            } else {
                var v25 uint = 255
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                fmt.Println(v10, v10, v18)
                _ = v25
            }
            _ = fi19
        }
        _ = v18
    }
    var v26 uint32 = 2
    var v27 int = -37
    for fi28 := 0; fi28 < 2; fi28++ {
        var v29 string = "a"
        if len(v1) > 0 {
            v1[0] = 0
        }
        _ = fi28
        _ = v29
    }
    v30 := GenS5{F0: false, F1: (v10 & (v10 + v10))}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if ((!((!(true || true)) || false)) && (100 > 1)) {
        v31 := GenS6{F0: false}
        if true {
            v10 -= v10
            var v32 uint32 = (^(v26))
            _ = v32
        } else {
            var v33 int8 = 42
            switch (v27 - (v27 + (^(v27)))) {
            case -9223372036854775808:
                v3 = append(v3, v26)
            case 42:
                switch v27 {
                case 990:
                    fmt.Println(v33, v33)
                case 426:
                    v10 -= v10
                case -2147483648:
                    v34 := GenS1{F0: 28, F1: 4294967295, F2: 255}
                    _ = v34
                }
            case -69:
                v35 := GenS7{F0: 0}
                _ = v35
            default:
                v36 := GenS8{F0: 175, F1: -32768, F2: 42}
                _ = v36
            }
            _ = v33
        }
        if false {
            switch v27 {
            case 42:
                v27 += genFn0(255, -1, 0)
            case 74:
                if false {
                    var v37 float64 = -2.25
                    v38 := GenS2{F0: 100}
                    var v39 float64 = (float64(3.14159) + float64(3.14159))
                    _ = v37
                    _ = v38
                    _ = v39
                } else {
                    if len(v3) > 0 {
                        v3[0] = (v26 * v26)
                    }
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                }
            default:
                if (1 > 198) {
                    v26 += v26
                    if (false && (((3.14159 / 1.5) * 1.5) < float64(3.14159))) {
                        var v40 int16 = 42
                        if false {
                            var v41 float64 = (1.5 - (0.0 - 1e10))
                            _ = v41
                        }
                        if true {
                            v42 := GenS9{F0: v27, F1: 0}
                            v43 := GenS10{F0: 42}
                            if (1.5 < 1.5) {
                                fmt.Println(v26, v27)
                                var v44 float32 = 1e10
                                _ = v44
                            }
                            _ = v42
                            _ = v43
                        }
                        _ = v40
                    } else {
                        v45 := GenS8{F0: 128, F1: 1000, F2: 720}
                        _ = v45
                    }
                    var v46 bool = false
                    _ = v46
                }
            }
            var v47 uint64 = 792
            _ = v47
        }
        _ = v31
    } else {
        v27 = (v27 ^ genFn0(42, -1, -1))
        for fi48 := 0; fi48 < 16; fi48++ {
            if true {
                var v49 uint = 0
                if len(v4) > 0 {
                    v4[0] = 2
                }
                switch 401 {
                case -50:
                    if len(v4) > 0 {
                        v4[0] = (v26 & 4)
                    }
                case 146:
                    if (943 == 42) {
                        fmt.Println(v27, v27)
                    }
                default:
                    var v50 uint32 = 0
                    _ = v50
                }
                _ = v49
            } else {
                for fi51 := 0; fi51 < 14; fi51++ {
                    fmt.Println(fi51, fi48)
                    v27 -= (v27 + 9223372036854775807)
                    v52 := GenS3{F0: int(v26), F1: 42}
                    _ = fi51
                    _ = v52
                }
                var v53 bool = true
                if (740 > 1) {
                    v54 := GenS3{F0: genFn0(1, 32767, 42), F1: 100}
                    _ = v54
                } else {
                    if (!v53) {
                        _ = len(v4)
                        v55 := GenS9{F0: (v27 - (fi48 & (fi48 ^ v27))), F1: v26}
                        _ = v55
                    } else {
                        _ = len(v1)
                        for fi56 := 0; fi56 < 3; fi56++ {
                            v57 := GenS6{F0: true}
                            _ = fi56
                            _ = v57
                        }
                    }
                    var v58 int32 = -2147483648
                    _ = v58
                }
                _ = v53
            }
            if (-719 > 1) {
                v10 -= (v10 / ((int64(v27)) | 1))
            }
            _ = fi48
        }
        if len(v1) > 0 {
            v1[0] = 65535
        }
    }
    switch (v27 * -1) {
    case -891:
        for fi59 := 0; fi59 < 11; fi59++ {
            v26 = v26
            var v60 uint = 18446744073709551615
            _ = fi59
            _ = v60
        }
    default:
        fmt.Println(v27, v27)
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if (false && (-2.25 == -2.25)) {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        if ((v10 * v10) > (v10 >> ((206) & 63))) {
            v61 := GenS3{F0: genFn0(178, -1, int32(v10)), F1: 42}
            var v62 uint32 = (^(v26))
            for fi63 := 0; fi63 < 4; fi63++ {
                var v64 int8 = 42
                v65 := GenS11{F0: 778}
                _ = fi63
                _ = v64
                _ = v65
            }
            _ = v61
            _ = v62
        }
        if len(v1) > 0 {
            v1[0] = (v26 | (^(v26)))
        }
    } else {
        switch (^(v27)) {
        case -1:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        case 2147483647:
            v10 = -963
        case -633:
            for fi66 := 0; fi66 < 9; fi66++ {
                if false {
                    v67 := GenS12{F0: 2, F1: -2147483648, F2: 42}
                    var v68 int = (-(v27))
                    _ = v67
                    _ = v68
                } else {
                    switch (fi66 ^ (-(fi66))) {
                    case 32768:
                        var v69 string = (("a" + ((((("foo bar" + "hello") + "hello") + "") + "hello") + "hello")) + "")
                        _ = v69
                    default:
                        var v70 int8 = 1
                        _ = v70
                    }
                }
                var v71 float32 = (1.5 / 1.5)
                fmt.Println(v10)
                _ = fi66
                _ = v71
            }
        }
        for fi72 := 0; fi72 < 17; fi72++ {
            var v73 uint = 1
            _ = fi72
            _ = v73
        }
        for fi74 := 0; fi74 < 9; fi74++ {
            switch ("a" + "") {
            case "a":
                v26 = 42
            }
            v75 := GenS12{F0: 193, F1: -461, F2: 65535}
            var v76 string = ("\n\t" + ("\n\t" + "a"))
            _ = fi74
            _ = v75
            _ = v76
        }
    }
    switch v27 {
    case -1:
        if false {
            fmt.Println(v27, v10, v10)
        }
    case 436:
        switch ("foo bar" + "foo bar") {
        case "foo bar":
            fmt.Println(v26, v27, v27)
        }
    case 42:
        fmt.Println(v26)
    }
    switch 0 {
    case 32767:
        switch "foo bar" {
        case "a":
            v27 -= 1
        case "\n\t":
            v10 += v10
        default:
            var v77 string = ""
            _ = v77
        }
    case 128:
        v10 -= (^(v10))
    case 42:
        for fi78 := 0; fi78 < 18; fi78++ {
            v79 := GenS13{F0: 116, F1: 42, F2: -1}
            v10 -= (v10 & (v10 >> ((4294967295) & 63)))
            _ = fi78
            _ = v79
        }
    default:
        if true {
            var v80 bool = true
            fmt.Println(v26, v27, v27)
            var v81 bool = (((!v80) && (!v80)) || true)
            _ = v80
            _ = v81
        }
    }
    fmt.Println(v10, v10)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v10
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v30
    _ = v4
}
