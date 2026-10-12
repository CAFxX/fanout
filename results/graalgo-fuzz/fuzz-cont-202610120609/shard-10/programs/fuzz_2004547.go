package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64; F2 uint32 }
type GenS2 struct { F0 int16; F1 uint16; F2 int }
type GenS3 struct { F0 bool; F1 int32 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 int8; F1 bool; F2 int64 }
func genFn0(p0 uint8, p1 int8, p2 uint) int32 {
    p2 -= 18446744073709551615
    v7 := GenS1{F0: 65535, F1: 854, F2: 42}
    var v8 uint16 = 137
    p0 -= (p0 >> ((p2) & 63))
    for fi9 := 0; fi9 < 20; fi9++ {
        for fi10 := 0; fi10 < 2; fi10++ {
            v11 := []int{}
            _ = len(v11)
            if ((p2 | (p2 * (p2 + p2))) == (p2 | p2)) {
                fmt.Println(p0, p0, p1)
                var v12 string = "a"
                _ = v12
            }
            _ = fi10
            _ = v11
        }
        _ = fi9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 42
}

func main() {
    var v1 int = 0
    var v2 int = (v1 | (^(v1)))
    v3 := make(map[uint32]bool)
    v4 := []int{(v1 ^ (v2 & (v1 + v1))), (-(v1))}
    var v5 int64 = 42
    v6 := []uint32{42, 2, 2}
    v13 := GenS2{F0: 42, F1: 1, F2: (v1 - (v1 / ((32768) | 1)))}
    v14 := GenS2{F0: 32767, F1: 344, F2: (-(v1))}
    if (!false) {
        for fi15 := 0; fi15 < 16; fi15++ {
            fmt.Println(v1, v2)
            _ = fi15
        }
        for fi16 := 0; fi16 < 12; fi16++ {
            var v17 float64 = 0.1
            if true {
                if true {
                    fmt.Println(v17)
                } else {
                    var v18 uint = 1
                    v19 := GenS3{F0: false, F1: 482}
                    _ = v18
                    _ = v19
                }
                var v20 int = (-(v1))
                _ = v20
            } else {
                for fi21 := 0; fi21 < 14; fi21++ {
                    v5 += v5
                    switch fi21 {
                    case 32767:
                        if true {
                            v17 -= ((((float64(0.1) / ((v17 + float64(0.0)) - 3.14159)) * (((float64(-2.25) - (v17 / ((v17 * v17) * (float64((3.14159 / ((1.5 - 0.1) * 3.14159))) * v17)))) - v17) - (1.5 + v17))) - 0.0) - v17)
                        }
                    case 2147483647:
                        var v22 int64 = -1
                        _ = v22
                    }
                    if true {
                        v3[2147483647] = ((v5 | 1) <= v5)
                    }
                    _ = fi21
                }
                if (-1 != -344) {
                    var v23 int8 = 42
                    var v24 int = (v1 ^ 42)
                    _ = v23
                    _ = v24
                } else {
                    var v25 uint = 2
                    _ = len(v4)
                    _ = v25
                }
                switch v1 {
                case -312:
                    var v26 int8 = 100
                    _ = v26
                }
            }
            _ = fi16
            _ = v17
        }
    }
    if true {
        _ = len(v4)
        if ((162 != 755) && true) {
            var v27 uint8 = 126
            v28 := GenS2{F0: -570, F1: 0, F2: v1}
            v2 = v1
            _ = v27
            _ = v28
        }
        switch "\n\t" {
        case "\n\t":
            if true {
                fmt.Println(v1, v5, v1)
            } else {
                for fi29 := 0; fi29 < 12; fi29++ {
                    _ = len(v4)
                    for fi30 := 0; fi30 < 11; fi30++ {
                        fmt.Println(v5, fi29, v2)
                        _ = fi30
                    }
                    _ = fi29
                }
                if len(v6) > 0 {
                    v6[0] = 42
                }
            }
        case "foo bar":
            switch -462 {
            case 127:
                v31 := GenS2{F0: 42, F1: 609, F2: v2}
                _ = v31
            case -2147483648:
                v3[42] = false
            case 32768:
                v32 := GenS2{F0: int16(v5), F1: 0, F2: v2}
                _ = v32
            }
        }
    }
    if (46 >= 964) {
        var v33 uint8 = 0
        switch (v2 | v1) {
        case -978:
            var v34 float32 = ((1e10 + float32(-2.25)) + 1e10)
            _ = v34
        case 475:
            for fi35 := 0; fi35 < 11; fi35++ {
                var v36 int8 = 42
                v1 = v1
                var v37 int = v1
                _ = fi35
                _ = v36
                _ = v37
            }
        }
        v38 := v3[2147483648]
        _ = v33
        _ = v38
    } else {
        for fi39 := 0; fi39 < 9; fi39++ {
            for fi40 := 0; fi40 < 11; fi40++ {
                v1 += (v1 - fi39)
                fmt.Println(fi39, fi39)
                v2 -= fi39
                _ = fi40
            }
            switch "a" {
            case "foo bar":
                if (-2147483648 == 42) {
                    var v41 int8 = 90
                    _ = v41
                } else {
                    v42 := GenS4{F0: (v5 >> ((18446744073709551615) & 63))}
                    _ = v42
                }
            case "a":
                if false {
                    v5 = (v5 & (v5 & 1))
                    if (255 >= 529) {
                        var v43 float32 = (3.14159 - 1.5)
                        var v44 uint16 = 65535
                        for fi45 := 0; fi45 < 14; fi45++ {
                            v2 -= v2
                            v3[2] = true
                            switch "" {
                            case "":
                                v1 += (fi45 * 0)
                            default:
                                delete(v3, 809)
                            }
                            v46 := GenS2{F0: 414, F1: 0, F2: 1}
                            _ = fi45
                            _ = v46
                        }
                        _ = v43
                        _ = v44
                    }
                }
            case "hello":
                var v47 int8 = 42
                _ = v47
            default:
                var v48 bool = (v5 > (v5 / ((v5) | 1)))
                _ = v48
            }
            delete(v3, 665)
            _ = fi39
        }
        var v49 int8 = 102
        var v50 int16 = 990
        _ = v49
        _ = v50
    }
    if false {
        v2 = v1
        v1 -= v1
    } else {
        for fi51 := 0; fi51 < 18; fi51++ {
            var v52 int32 = 0
            if true {
                v2 += -9223372036854775808
            } else {
                if false {
                    v1 = (v2 % ((42) | 1))
                }
                var v53 string = "\n\t"
                v52 = (v52 ^ v52)
                _ = v53
            }
            _ = fi51
            _ = v52
        }
    }
    switch -9223372036854775808 {
    case -9223372036854775808:
        if true {
            switch "foo bar" {
            case "hello":
                for fi54 := 0; fi54 < 17; fi54++ {
                    fmt.Println(v5, v1, fi54)
                    v5 += -196
                    _ = fi54
                }
            case "\n\t":
                var v55 float64 = (0.1 + -2.25)
                _ = v55
            case "":
                v2 = (v1 * -491)
            }
            v2 -= (v1 % ((v2) | 1))
            if true {
                v1 += v1
            } else {
                if true {
                    v56 := GenS2{F0: -305, F1: 2, F2: (v2 >> ((255) & 63))}
                    v1 -= v1
                    _ = v56
                }
            }
        }
    case -97:
        switch v1 {
        case 2147483647:
            var v57 uint16 = 1
            _ = v57
        default:
            fmt.Println(v2)
        }
    case 127:
        for fi58 := 0; fi58 < 1; fi58++ {
            v59 := GenS5{F0: 0}
            _ = fi58
            _ = v59
        }
    }
    var v60 uint = 255
    var v61 int64 = (v5 / (((v5 ^ (v5 ^ -9223372036854775808))) | 1))
    for fi62 := 0; fi62 < 12; fi62++ {
        switch (fi62 * 1) {
        case 32768:
            fmt.Println(v5, v5)
        default:
            v60 = (v60 >> ((42) & 63))
        }
        var v63 uint16 = 578
        v60 -= 65535
        _ = fi62
        _ = v63
    }
    switch ("hello" + "\n\t") {
    case "foo bar":
        for fi64 := 0; fi64 < 13; fi64++ {
            for fi65 := 0; fi65 < 20; fi65++ {
                switch (-(v1)) {
                case -9223372036854775808:
                    var v66 uint32 = 498
                    _ = v66
                }
                v61 -= v5
                _ = fi65
            }
            v1 += (^(v2))
            fmt.Println(fi64, v1)
            _ = fi64
        }
    case "\n\t":
        if false {
            var v67 int16 = -805
            _ = v67
        }
    case "":
        if false {
            if (0 < 670) {
                v2 -= (v1 ^ 0)
            }
            if false {
                for fi68 := 0; fi68 < 19; fi68++ {
                    switch v1 {
                    case -9223372036854775808:
                        var v69 string = ""
                        _ = v69
                    }
                    v2 = (fi68 * 673)
                    for fi70 := 0; fi70 < 13; fi70++ {
                        v71 := v3[110]
                        _ = fi70
                        _ = v71
                    }
                    _ = fi68
                }
                v5 = v5
                v61 += (v61 >> ((v60) & 63))
            } else {
                switch -2147483648 {
                case 2147483647:
                    var v72 int32 = -2147483648
                    _ = v72
                case 128:
                    _ = len(v3)
                case 32768:
                    v2 -= v1
                default:
                    v3[641] = (!true)
                }
                v73 := GenS6{F0: -1, F1: false, F2: (v61 >> ((v60) & 63))}
                _ = v73
            }
            v61 = -1
        } else {
            var v74 uint32 = 520
            _ = v74
        }
    }
    if (false || (355 <= 6)) {
        fmt.Println(v2, v1)
        if true {
            fmt.Println(v2, v1, v60)
            fmt.Println(v5)
        }
    } else {
        for fi75 := 0; fi75 < 12; fi75++ {
            v5 = -1
            for fi76 := 0; fi76 < 9; fi76++ {
                if (((2147483648 >= 1) || false) && ("" <= "hello")) {
                    switch 32768 {
                    case 633:
                        var v77 float64 = 1e10
                        _ = v77
                    default:
                        v2 -= (^(v2))
                    }
                    v78 := GenS3{F0: (false || (!true)), F1: 42}
                    _ = v78
                }
                _ = fi76
            }
            for fi79 := 0; fi79 < 7; fi79++ {
                fmt.Println(v1, v61)
                if (!false) {
                    v80 := GenS3{F0: (-68 < -128), F1: -1}
                    fmt.Println(v1, v61)
                    var v81 uint16 = 255
                    _ = v80
                    _ = v81
                }
                _ = fi79
            }
            _ = fi75
        }
    }
    v82 := v3[41]
    var v83 float64 = 1e10
    if len(v6) > 0 {
        v6[0] = 42
    }
    if (!v82) {
        var v84 uint8 = 2
        if (v82 || (!(-768 == 852))) {
            _ = len(v6)
            v84 = v84
            var v85 int32 = 2147483647
            _ = v85
        } else {
            v2 = v1
            var v86 uint8 = (v84 | (v84 ^ (v84 / ((v84) | 1))))
            _ = v86
        }
        _ = v84
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v60:", v60)
    fmt.Println("v61:", v61)
    fmt.Println("v82:", v82)
    fmt.Println("v83:", v83)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v60
    _ = v61
    _ = v82
    _ = v83
}
