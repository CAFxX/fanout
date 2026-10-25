package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool }
type GenS2 struct { F0 string; F1 bool }
type GenS3 struct { F0 uint16; F1 uint32 }
type GenS4 struct { F0 bool; F1 uint64; F2 int8 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 uint8 }

func main() {
    var v1 uint64 = 2
    var v2 uint = 831
    var v3 string = (("foo bar" + "a") + ("a" + "a"))
    var v4 int16 = -32768
    var v5 uint8 = 100
    var v6 bool = (true || false)
    fmt.Println(v6, v5)
    fmt.Println(v2, v1)
    if v6 {
        fmt.Println(v5, v5, v1)
        if (-2147483648 <= -1) {
            for fi7 := 0; fi7 < 19; fi7++ {
                for fi8 := 0; fi8 < 19; fi8++ {
                    v9 := []int{fi7}
                    _ = fi8
                    _ = v9
                }
                v3 = (("foo bar" + v3) + "a")
                v1 += 4294967295
                _ = fi7
            }
            for fi10 := 0; fi10 < 7; fi10++ {
                v6 = v6
                if v6 {
                    v11 := []int{fi10}
                    if v6 {
                        if (false && ((v2 < 297) && (((-2.25 - 1e10) == (1.5 + (1.5 * ((1.5 * -2.25) * 0.0)))) && false))) {
                            v12 := GenS1{F0: (v4 - 1), F1: true}
                            var v13 float32 = 0.1
                            _ = v12
                            _ = v13
                        } else {
                            var v14 int8 = 0
                            v15 := GenS1{F0: 42, F1: (24 >= 2147483648)}
                            _ = v14
                            _ = v15
                        }
                        if v6 {
                            switch -1 {
                            case 9223372036854775807:
                                if len(v11) > 0 {
                                    v11[0] = fi10
                                }
                            default:
                                if ((((((-2.25 + 1e10) + 0.0) / 1e10) * 0.0) - 1.5) == -2.25) {
                                    var v16 bool = true
                                    var v17 int = (fi10 | fi10)
                                    fmt.Println(v3)
                                    _ = v16
                                    _ = v17
                                } else {
                                    fmt.Println(v1)
                                    for fi18 := 0; fi18 < 12; fi18++ {
                                        var v19 uint8 = 2
                                        _ = fi18
                                        _ = v19
                                    }
                                    var v20 int = fi10
                                    _ = v20
                                }
                            }
                        } else {
                            var v21 int16 = 42
                            v22 := GenS1{F0: (v4 / ((v21) | 1)), F1: false}
                            v2 = v2
                            _ = v21
                            _ = v22
                        }
                        var v23 uint16 = 2
                        _ = v23
                    } else {
                        fmt.Println(fi10, v3)
                        if (-463 >= 417) {
                            v24 := GenS1{F0: (v4 ^ v4), F1: v6}
                            v25 := GenS2{F0: v3, F1: v6}
                            for fi26 := 0; fi26 < 17; fi26++ {
                                switch ("hello" + (v3 + "\n\t")) {
                                case "a":
                                    if v6 {
                                        v5 -= v5
                                        if (!v6) {
                                            switch (v3 + (v3 + "\n\t")) {
                                            case "hello":
                                                fmt.Println(v1)
                                            case "a":
                                                if (!false) {
                                                    var v27 int8 = 59
                                                    var v28 int32 = -1
                                                    _ = v27
                                                    _ = v28
                                                } else {
                                                    if len(v11) >= 2 {
                                                        v11 = v11[1:len(v11)]
                                                    }
                                                    if v6 {
                                                        v6 = true
                                                    } else {
                                                        if ((!(0.1 != float32(1.5))) || (v6 && (-1 >= -17))) {
                                                            v6 = (v6 && (-35 <= 92))
                                                            switch fi26 {
                                                            case 4294967295:
                                                                switch (fi10 << ((732) & 63)) {
                                                                case -9223372036854775808:
                                                                    switch (v3 + v3) {
                                                                    case "foo bar":
                                                                        v29 := GenS1{F0: (v4 ^ v4), F1: ((v1 - (^(v1))) != (v1 >> (((v2 + (v2 * v2))) & 63)))}
                                                                        _ = v29
                                                                    }
                                                                default:
                                                                    v30 := GenS2{F0: "a", F1: v6}
                                                                    _ = v30
                                                                }
                                                            case 32768:
                                                                var v31 int = -675
                                                                _ = v31
                                                            }
                                                        } else {
                                                            if (0.0 < (float32(-2.25) - (-2.25 - 1e10))) {
                                                                v4 -= 32767
                                                            }
                                                            fmt.Println(fi10)
                                                            _ = len(v11)
                                                        }
                                                    }
                                                    fmt.Println(v6, v4)
                                                }
                                            case "\n\t":
                                                v32 := GenS3{F0: 411, F1: 4294967295}
                                                _ = v32
                                            default:
                                                v33 := GenS4{F0: v6, F1: v1, F2: 125}
                                                _ = v33
                                            }
                                            v34 := GenS1{F0: v4, F1: v6}
                                            _ = v34
                                        } else {
                                            if (297 <= 2) {
                                                v6 = (1.5 >= float64(-2.25))
                                            } else {
                                                var v35 int64 = int64(v5)
                                                if len(v11) >= 2 {
                                                    v11 = v11[1:len(v11)]
                                                }
                                                switch "hello" {
                                                case "":
                                                    fmt.Println(v4, v2)
                                                }
                                                _ = v35
                                            }
                                            var v36 int = fi26
                                            _ = v36
                                        }
                                    } else {
                                        if false {
                                            v5 -= v5
                                            v2 = (v2 | (v2 + (^(v2))))
                                            var v37 float64 = -2.25
                                            _ = v37
                                        }
                                        switch v3 {
                                        case "foo bar":
                                            switch (fi26 + (^(fi26))) {
                                            case 32768:
                                                v11 = append(v11, (fi10 * fi26))
                                            default:
                                                var v38 int32 = -2147483648
                                                _ = v38
                                            }
                                        case "":
                                            v4 = v4
                                        case "\n\t":
                                            v39 := GenS3{F0: 255, F1: 2}
                                            _ = v39
                                        default:
                                            if (!(32767 >= 1)) {
                                                if len(v11) >= 2 {
                                                    v11 = v11[1:len(v11)]
                                                }
                                            } else {
                                                var v40 uint64 = v1
                                                v41 := GenS5{F0: 337}
                                                if (v6 || v6) {
                                                    fmt.Println(fi10, fi26, fi10)
                                                } else {
                                                    fmt.Println(v1, v4)
                                                    v6 = true
                                                    v4 -= 993
                                                }
                                                var v42 int32 = 42
                                                _ = v40
                                                _ = v41
                                                _ = v42
                                            }
                                        }
                                        v5 = (v5 >> (((v2 + (v2 - (v2 >> ((v2) & 63))))) & 63))
                                    }
                                case "foo bar":
                                    var v43 float64 = 0.1
                                    _ = v43
                                default:
                                    v44 := GenS2{F0: "foo bar", F1: false}
                                    _ = v44
                                }
                                _ = fi26
                            }
                            _ = v24
                            _ = v25
                        }
                    }
                    fmt.Println(v1, v5, v2)
                    _ = v11
                }
                switch "\n\t" {
                case "foo bar":
                    for fi45 := 0; fi45 < 9; fi45++ {
                        v46 := GenS3{F0: uint16(v2), F1: 4294967295}
                        _ = fi45
                        _ = v46
                    }
                case "hello":
                    if v6 {
                        if v6 {
                            for fi47 := 0; fi47 < 4; fi47++ {
                                var v48 float64 = ((1.5 + ((3.14159 - ((0.0 / 0.1) * (1e10 * (1e10 / ((float64(3.14159) * -2.25) / 1.0))))) + (1e10 + 0.0))) - (-2.25 + 1e10))
                                var v49 uint32 = 2147483648
                                _ = fi47
                                _ = v48
                                _ = v49
                            }
                            v6 = true
                            for fi50 := 0; fi50 < 19; fi50++ {
                                if ((1e10 + (1e10 / -2.25)) >= -2.25) {
                                    v3 = v3
                                }
                                _ = fi50
                            }
                        }
                        if v6 {
                            fmt.Println(v6, v1)
                            if (v2 != 0) {
                                var v51 bool = (v6 && (3.14159 <= 1.5))
                                var v52 uint16 = 32767
                                v3 = v3
                                _ = v51
                                _ = v52
                            }
                            var v53 int8 = int8(fi10)
                            _ = v53
                        }
                    } else {
                        var v54 uint16 = 240
                        if v6 {
                            v4 += (v4 + (v4 / ((32767) | 1)))
                        } else {
                            v2 += v2
                            v55 := GenS4{F0: v6, F1: (v1 & v1), F2: 0}
                            var v56 int = fi10
                            _ = v55
                            _ = v56
                        }
                        v57 := []int{}
                        _ = v54
                        _ = v57
                    }
                case "a":
                    var v58 int32 = int32(v5)
                    _ = v58
                default:
                    v59 := []int{(fi10 << ((v2) & 63)), (fi10 << ((uint(v1)) & 63)), int(v4)}
                    _ = v59
                }
                _ = fi10
            }
        }
    } else {
        v4 = v4
    }
    v60 := GenS6{F0: (^(v5))}
    v61 := GenS2{F0: v3, F1: (v5 <= v5)}
    switch 590 {
    case 1:
        fmt.Println(v5, v3)
    case 4294967295:
        if (!(42 <= -873)) {
            for fi62 := 0; fi62 < 13; fi62++ {
                switch (fi62 << (((v2 % (((v2 | (^(v2)))) | 1))) & 63)) {
                case 32768:
                    var v63 int64 = int64(v1)
                    _ = v63
                }
                if (42 >= 86) {
                    var v64 uint32 = 2147483648
                    if (int(v1) > 15) {
                        for fi65 := 0; fi65 < 19; fi65++ {
                            v66 := []int{}
                            var v67 uint16 = 2
                            fmt.Println(v6, v4)
                            _ = fi65
                            _ = v66
                            _ = v67
                        }
                        var v68 float32 = ((1e10 / (3.14159 - float32(((((0.0 * 3.14159) + ((1.5 * 0.0) - 1.5)) / -2.25) + float64(0.0))))) / 3.14159)
                        _ = v68
                    }
                    fmt.Println(v6, v5)
                    _ = v64
                }
                _ = fi62
            }
            var v69 float32 = float32(3.14159)
            for fi70 := 0; fi70 < 9; fi70++ {
                if true {
                    v71 := []int{(fi70 - fi70), fi70, 0}
                    v69 = (1.5 + 1e10)
                    _ = v71
                } else {
                    if (330 <= 2147483647) {
                        v72 := []int{}
                        var v73 int16 = -184
                        _ = v72
                        _ = v73
                    }
                    v3 = "hello"
                    v5 -= v5
                }
                var v74 uint16 = 276
                var v75 uint = (v2 - uint(v74))
                _ = fi70
                _ = v74
                _ = v75
            }
            _ = v69
        } else {
            for fi76 := 0; fi76 < 2; fi76++ {
                v5 -= v5
                _ = fi76
            }
        }
    case 9223372036854775807:
        var v77 int16 = v4
        _ = v77
    default:
        v6 = (false && false)
    }
    switch "hello" {
    case "":
        var v78 float32 = (float32(1.5) * 0.1)
        _ = v78
    case "foo bar":
        var v79 int32 = -2147483648
        _ = v79
    }
    fmt.Println(v6, v4)
    for fi80 := 0; fi80 < 6; fi80++ {
        if v6 {
            v81 := []int{}
            _ = v81
        } else {
            var v82 uint = (v2 >> (((v2 | v2)) & 63))
            var v83 uint8 = uint8(v1)
            v84 := []int{}
            _ = v82
            _ = v83
            _ = v84
        }
        _ = fi80
    }
    fmt.Println(v5)
    switch v3 {
    case "a":
        if v6 {
            v85 := GenS3{F0: 1, F1: 2}
            fmt.Println(v4, v5, v6)
            _ = v85
        }
    }
    for fi86 := 0; fi86 < 20; fi86++ {
        if true {
            switch fi86 {
            case 2147483647:
                switch v3 {
                case "a":
                    if ((-128 < 127) || false) {
                        v87 := GenS2{F0: (v3 + ((v3 + ((((("a" + (("" + "hello") + (v3 + v3))) + (v3 + "hello")) + ((v3 + v3) + "foo bar")) + v3) + v3)) + (((v3 + "hello") + "foo bar") + "a"))), F1: v6}
                        _ = v87
                    } else {
                        fmt.Println(v1, v4, fi86)
                        var v88 float32 = 1e10
                        _ = v88
                    }
                }
            case 128:
                var v89 int32 = int32(v5)
                _ = v89
            }
        } else {
            v4 += (^(v4))
        }
        _ = fi86
    }
    var v90 float32 = -2.25
    v91 := []int{}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v90:", v90)
    fmt.Println("slicelen v91:", len(v91), v91)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v60
    _ = v61
    _ = v90
    _ = v91
}
