package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int8; F1 int16 }
type GenS5 struct { F0 int16; F1 uint32 }
type GenS6 struct { F0 int8; F1 bool }
type GenS7 struct { F0 string; F1 bool }
type GenS8 struct { F0 uint32; F1 uint8; F2 uint32 }
func genFn0() uint {
    var v4 int64 = 539
    if false {
        v5 := GenS1{F0: 224}
        if true {
            v4 = v4
        } else {
            var v6 int8 = 127
            _ = v6
        }
        var v7 uint16 = 467
        _ = v5
        _ = v7
    }
    if false {
        var v8 bool = ((827 == 2) || true)
        _ = v8
    } else {
        switch "foo bar" {
        case "\n\t":
            for fi9 := 0; fi9 < 12; fi9++ {
                var v10 bool = (304 == 255)
                _ = fi9
                _ = v10
            }
        default:
            v4 = v4
        }
        for fi11 := 0; fi11 < 9; fi11++ {
            switch 32767 {
            case 127:
                switch "foo bar" {
                case "hello":
                    if (127 < (fi11 & fi11)) {
                        var v12 int = (fi11 + fi11)
                        var v13 int32 = 0
                        _ = v12
                        _ = v13
                    }
                case "":
                    v14 := []int{}
                    _ = v14
                }
            case -861:
                if true {
                    v4 += v4
                    var v15 float32 = (1e10 * (0.0 - (1.5 / 0.1)))
                    _ = v15
                } else {
                    v16 := []int{}
                    if (true && ((760 == 603) && true)) {
                        switch "foo bar" {
                        case "hello":
                            v17 := GenS2{F0: 2147483647}
                            _ = v17
                        default:
                            v18 := GenS3{F0: v4}
                            _ = v18
                        }
                    }
                    if (0.0 == 0.0) {
                        v4 = (^(v4))
                    }
                    _ = v16
                }
            case 1:
                if (-128 == -128) {
                    v19 := []int{(fi11 ^ 127), (fi11 << ((1) & 63)), fi11, fi11}
                    if len(v19) > 0 {
                        v19[0] = fi11
                    }
                    var v20 bool = true
                    _ = v19
                    _ = v20
                }
            default:
                v4 = 1
            }
            _ = fi11
        }
        switch ("foo bar" + (((("" + "") + "hello") + "\n\t") + "hello")) {
        case "\n\t":
            v21 := GenS1{F0: 255}
            _ = v21
        }
    }
    v4 += v4
    v4 = v4
    _ = v4
    return 1
}
func genFn1(p0 int, p1 uint16, p2 uint8) uint32 {
    var v22 float64 = (0.1 - -2.25)
    var v23 uint64 = 121
    _ = p0
    _ = p1
    _ = p2
    _ = v22
    _ = v23
    return 42
}

func main() {
    var v1 uint = 313
    v2 := []int{205, -863, 9223372036854775807}
    var v3 uint16 = 0
    if (((0.0 - 3.14159) > (0.1 / 1e10)) && ((-2.25 + 1e10) > (1.5 + (1.5 + ((0.1 / (0.1 + (-2.25 * 3.14159))) - 1e10))))) {
        fmt.Println(v1, v3)
        for fi24 := 0; fi24 < 16; fi24++ {
            fmt.Println(fi24)
            if true {
                if len(v2) > 0 {
                    v2[0] = fi24
                }
                if len(v2) > 0 {
                    v2[0] = fi24
                }
                if ((true || (0.1 <= (1.5 + float64((float32((1e10 * 0.1)) / 0.1))))) || false) {
                    for fi25 := 0; fi25 < 19; fi25++ {
                        switch "" {
                        case "\n\t":
                            var v26 int = fi24
                            _ = v26
                        default:
                            switch (fi24 % (((fi25 - fi24)) | 1)) {
                            case 2147483647:
                                v2 = append(v2, fi25)
                            default:
                                v1 = v1
                            }
                        }
                        _ = fi25
                    }
                    v2 = append(v2, (fi24 >> ((v1) & 63)))
                    fmt.Println(v1, fi24, v3)
                }
            } else {
                v3 -= 847
                v3 = (v3 | v3)
            }
            v1 = v1
            _ = fi24
        }
        if (v3 != 811) {
            switch 127 {
            case -476:
                if (!false) {
                    for fi27 := 0; fi27 < 11; fi27++ {
                        v1 += v1
                        fmt.Println(v3, v1)
                        var v28 uint32 = 65535
                        _ = fi27
                        _ = v28
                    }
                    var v29 uint16 = 255
                    _ = v29
                } else {
                    if (!((float64((((3.14159 * (0.1 / (float32(1e10) / float32(float64(0.1))))) / -2.25) - 0.0)) / 1.5) >= 1.5)) {
                        if (2147483647 <= 1) {
                            for fi30 := 0; fi30 < 19; fi30++ {
                                v1 = (v1 + 42)
                                var v31 uint8 = 19
                                _ = fi30
                                _ = v31
                            }
                            fmt.Println(v1, v1)
                        } else {
                            if false {
                                for fi32 := 0; fi32 < 8; fi32++ {
                                    v33 := GenS2{F0: 0}
                                    for fi34 := 0; fi34 < 11; fi34++ {
                                        v3 = (^(v3))
                                        var v35 int8 = 127
                                        if len(v2) > 0 {
                                            v2[0] = fi34
                                        }
                                        _ = fi34
                                        _ = v35
                                    }
                                    v1 = v1
                                    _ = fi32
                                    _ = v33
                                }
                                v1 += (v1 >> ((genFn0()) & 63))
                            }
                        }
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        if true {
                            var v36 bool = true
                            if false {
                                v1 -= v1
                            } else {
                                var v37 int8 = -128
                                if v36 {
                                    v1 += 42
                                } else {
                                    if (!(false || false)) {
                                        v3 -= v3
                                        var v38 uint8 = 128
                                        v38 += 255
                                        _ = v38
                                    } else {
                                        v36 = (v37 == (v37 - -128))
                                        var v39 bool = (!v36)
                                        if false {
                                            var v40 int = -1
                                            if v36 {
                                                switch ("hello" + "foo bar") {
                                                case "hello":
                                                    if v36 {
                                                        if true {
                                                            var v41 int32 = -2147483648
                                                            for fi42 := 0; fi42 < 20; fi42++ {
                                                                if ((v3 < 349) || v36) {
                                                                    for fi43 := 0; fi43 < 20; fi43++ {
                                                                        if v36 {
                                                                            v36 = (-2.25 < -2.25)
                                                                            v41 = (v41 + int32(v40))
                                                                            v44 := GenS1{F0: (^(v1))}
                                                                            var v45 float32 = 0.1
                                                                            v41 = int32(v1)
                                                                            var v46 string = "a"
                                                                            fmt.Println(v45, v36, v41)
                                                                            _ = v44
                                                                            _ = v45
                                                                            _ = v46
                                                                        } else {
                                                                            var v47 float64 = (1.5 - float64(float32((0.1 - -2.25))))
                                                                            v39 = (true && (int(v41) > fi43))
                                                                            v48 := GenS4{F0: v37, F1: -1}
                                                                            _ = v47
                                                                            _ = v48
                                                                        }
                                                                        _ = fi43
                                                                    }
                                                                    var v49 uint8 = 128
                                                                    switch (v40 / (((fi42 ^ (fi42 * v40))) | 1)) {
                                                                    case 0:
                                                                        v50 := GenS5{F0: 0, F1: uint32(v3)}
                                                                        _ = v50
                                                                    }
                                                                    _ = v49
                                                                }
                                                                var v51 uint8 = 202
                                                                var v52 float32 = 0.0
                                                                _ = fi42
                                                                _ = v51
                                                                _ = v52
                                                            }
                                                            fmt.Println(v3, v39, v41)
                                                            _ = v41
                                                        } else {
                                                            for fi53 := 0; fi53 < 9; fi53++ {
                                                                v54 := GenS2{F0: 147}
                                                                v55 := GenS4{F0: v37, F1: int16(v40)}
                                                                switch (("a" + ((("\n\t" + "") + ("" + ("a" + "a"))) + "\n\t")) + ("hello" + "\n\t")) {
                                                                case "":
                                                                    v2 = append(v2, 9223372036854775807)
                                                                case "hello":
                                                                    for fi56 := 0; fi56 < 19; fi56++ {
                                                                        v39 = v39
                                                                        v3 = v3
                                                                        var v57 uint8 = 255
                                                                        _ = fi56
                                                                        _ = v57
                                                                    }
                                                                case "a":
                                                                    _ = len(v2)
                                                                }
                                                                _ = fi53
                                                                _ = v54
                                                                _ = v55
                                                            }
                                                            var v58 uint64 = 2
                                                            if v39 {
                                                                _ = len(v2)
                                                                v59 := GenS5{F0: 42, F1: 2}
                                                                fmt.Println(v39, v37)
                                                                _ = v59
                                                            }
                                                            _ = v58
                                                        }
                                                        v37 = 86
                                                        var v60 int16 = 183
                                                        _ = v60
                                                    }
                                                default:
                                                    v1 += v1
                                                }
                                                var v61 uint = (v1 | v1)
                                                var v62 float32 = 1.5
                                                _ = v61
                                                _ = v62
                                            } else {
                                                for fi63 := 0; fi63 < 13; fi63++ {
                                                    v3 = (v3 - v3)
                                                    if ("a" != ("" + (("\n\t" + ("foo bar" + "a")) + ("hello" + ("\n\t" + "foo bar"))))) {
                                                        var v64 int = fi63
                                                        if v39 {
                                                            var v65 uint64 = 18446744073709551615
                                                            v66 := GenS6{F0: (v37 | v37), F1: (2147483647 > 0)}
                                                            _ = v65
                                                            _ = v66
                                                        }
                                                        _ = v64
                                                    } else {
                                                        var v67 bool = (uint8(v1) != 119)
                                                        _ = v67
                                                    }
                                                    v37 = 0
                                                    _ = fi63
                                                }
                                                var v68 bool = true
                                                switch (v40 + v40) {
                                                case 367:
                                                    var v69 uint = v1
                                                    _ = v69
                                                case 4294967295:
                                                    var v70 uint = v1
                                                    _ = v70
                                                }
                                                _ = v68
                                            }
                                            _ = v40
                                        }
                                        _ = v39
                                    }
                                    fmt.Println(v37, v36, v36)
                                    v1 = (v1 ^ 1)
                                }
                                _ = v37
                            }
                            var v71 int64 = 0
                            _ = v36
                            _ = v71
                        } else {
                            for fi72 := 0; fi72 < 12; fi72++ {
                                fmt.Println(v1, v1, fi72)
                                if true {
                                    var v73 int8 = -105
                                    _ = v73
                                }
                                var v74 string = ("a" + "\n\t")
                                _ = fi72
                                _ = v74
                            }
                            var v75 float32 = (0.1 * -2.25)
                            _ = v75
                        }
                    }
                    v1 = v1
                }
            }
            v1 += (v1 / ((genFn0()) | 1))
            v76 := GenS7{F0: ("hello" + "a"), F1: (!(true || true))}
            _ = v76
        }
    } else {
        v77 := GenS8{F0: genFn1(-1, 42, 128), F1: 128, F2: 0}
        switch -512 {
        case 214:
            if true {
                var v78 bool = (!false)
                _ = v78
            } else {
                for fi79 := 0; fi79 < 1; fi79++ {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = fi79
                }
                v1 += genFn0()
                _ = len(v2)
            }
        }
        var v80 bool = (!(v3 <= (^(v3))))
        _ = v77
        _ = v80
    }
    for fi81 := 0; fi81 < 17; fi81++ {
        v82 := GenS6{F0: 42, F1: (!(fi81 != fi81))}
        _ = fi81
        _ = v82
    }
    var v83 uint32 = genFn1(-641, 233, 54)
    switch -853 {
    case 127:
        if (!(false && (1 != 2))) {
            v83 -= 1
            if len(v2) > 0 {
                v2[0] = 127
            }
        } else {
            v84 := GenS6{F0: -100, F1: (false || ("" <= ("foo bar" + "\n\t")))}
            for fi85 := 0; fi85 < 16; fi85++ {
                if len(v2) > 0 {
                    v2[0] = (fi85 % ((fi85) | 1))
                }
                if len(v2) > 0 {
                    v2[0] = 0
                }
                _ = fi85
            }
            for fi86 := 0; fi86 < 2; fi86++ {
                switch -750 {
                case 127:
                    v83 = (v83 << (((^(v1))) & 63))
                case -1:
                    fmt.Println(v3)
                case 0:
                    var v87 float64 = 0.0
                    _ = v87
                default:
                    v3 = (v3 | v3)
                }
                _ = fi86
            }
            _ = v84
        }
    }
    var v88 string = "a"
    v3 = 42
    v88 = "a"
    for fi89 := 0; fi89 < 17; fi89++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = fi89
    }
    v3 = (v3 / ((617) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v83:", v83)
    fmt.Println("v88:", v88)
    _ = v1
    _ = v2
    _ = v3
    _ = v83
    _ = v88
}
