package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8 }
type GenS2 struct { F0 uint64; F1 uint8 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint64; F1 uint16 }
type GenS5 struct { F0 int64; F1 int; F2 bool }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 bool; F1 int8 }

func main() {
    var v1 float32 = -2.25
    var v2 uint32 = 0
    var v3 int8 = -2
    v4 := GenS1{F0: 2, F1: 255}
    var v5 float32 = 1e10
    var v6 bool = (float64(-2.25) < 1e10)
    v7 := []int{42, 0, -9223372036854775808, 252}
    var v8 uint8 = 128
    if false {
        if true {
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
        } else {
            for fi9 := 0; fi9 < 17; fi9++ {
                var v10 uint64 = 4294967295
                _ = fi9
                _ = v10
            }
            var v11 uint16 = 42
            _ = v11
        }
    } else {
        v2 -= (v2 | (v2 + v2))
    }
    v2 -= v2
    v7 = append(v7, 9223372036854775807)
    switch "" {
    case "hello":
        if (!v6) {
            v12 := GenS1{F0: 18446744073709551615, F1: (v8 % ((42) | 1))}
            var v13 int64 = 42
            _ = v12
            _ = v13
        } else {
            v8 -= v8
            if v6 {
                for fi14 := 0; fi14 < 5; fi14++ {
                    var v15 int64 = 9223372036854775807
                    var v16 int64 = v15
                    v17 := GenS1{F0: uint64(v8), F1: v8}
                    _ = fi14
                    _ = v15
                    _ = v16
                    _ = v17
                }
                if v6 {
                    if len(v7) >= 2 {
                        v7 = v7[1:len(v7)]
                    }
                    fmt.Println(v2, v5)
                    var v18 uint = 120
                    _ = v18
                } else {
                    switch 1 {
                    case 35:
                        v5 = 3.14159
                    case -600:
                        v6 = true
                    case 32768:
                        var v19 int64 = 1
                        _ = v19
                    default:
                        switch "foo bar" {
                        case "\n\t":
                            switch "" {
                            case "hello":
                                if (-40 == 1000) {
                                    var v20 float64 = 1e10
                                    v20 = v20
                                    switch "hello" {
                                    case "":
                                        v21 := GenS1{F0: 4294967295, F1: 127}
                                        _ = v21
                                    case "\n\t":
                                        if (true || v6) {
                                            if len(v7) > 0 {
                                                v7[0] = 0
                                            }
                                        }
                                    default:
                                        _ = len(v7)
                                    }
                                    _ = v20
                                }
                            case "a":
                                var v22 int8 = v3
                                _ = v22
                            default:
                                v23 := GenS2{F0: 4294967295, F1: (v8 >> ((1) & 63))}
                                _ = v23
                            }
                        case "":
                            var v24 int = -884
                            _ = v24
                        case "a":
                            v25 := GenS1{F0: 490, F1: v8}
                            _ = v25
                        }
                    }
                    switch "\n\t" {
                    case "":
                        v6 = v6
                    case "foo bar":
                        fmt.Println(v2)
                    default:
                        v8 = (v8 | 42)
                    }
                }
                v26 := GenS1{F0: 4294967295, F1: (v8 | v8)}
                _ = v26
            } else {
                _ = len(v7)
                for fi27 := 0; fi27 < 20; fi27++ {
                    var v28 uint32 = (v2 | (v2 >> ((65535) & 63)))
                    if len(v7) >= 2 {
                        v7 = v7[1:len(v7)]
                    }
                    _ = fi27
                    _ = v28
                }
                v29 := GenS3{F0: 30}
                _ = v29
            }
            v2 = (v2 * v2)
        }
    default:
        var v30 float64 = 0.0
        _ = v30
    }
    var v31 int8 = -1
    var v32 int64 = 1
    if v6 {
        var v33 uint = 65535
        _ = v33
    }
    if v6 {
        v2 += v2
        for fi34 := 0; fi34 < 10; fi34++ {
            for fi35 := 0; fi35 < 13; fi35++ {
                if v6 {
                    v36 := GenS1{F0: 698, F1: (v8 << ((65535) & 63))}
                    v32 = -620
                    var v37 uint32 = (^(v2))
                    v38 := GenS1{F0: 18446744073709551615, F1: v8}
                    _ = v36
                    _ = v37
                    _ = v38
                }
                v1 = 0.0
                _ = fi35
            }
            if len(v7) > 0 {
                v7[0] = (^(fi34))
            }
            var v39 uint64 = 1
            _ = fi34
            _ = v39
        }
    } else {
        var v40 float64 = float64(v1)
        v6 = ((false && v6) || v6)
        if (v6 || false) {
            v32 -= (v32 * 0)
            switch -692 {
            case -2147483648:
                switch "hello" {
                case "hello":
                    if ((539 != 0) && v6) {
                        var v41 bool = v6
                        _ = v41
                    } else {
                        v42 := GenS2{F0: 1, F1: (v8 >> ((4294967295) & 63))}
                        v2 = v2
                        switch "foo bar" {
                        case "foo bar":
                            var v43 int = 544
                            _ = v43
                        case "a":
                            v32 += -9223372036854775808
                        case "":
                            if len(v7) > 0 {
                                v7[0] = -1
                            }
                        default:
                            if len(v7) > 0 {
                                v7[0] = -495
                            }
                        }
                        _ = v42
                    }
                case "foo bar":
                    _ = len(v7)
                case "":
                    fmt.Println(v5)
                default:
                    v40 -= v40
                }
            case -9223372036854775808:
                for fi44 := 0; fi44 < 6; fi44++ {
                    var v45 int64 = (v32 - v32)
                    _ = fi44
                    _ = v45
                }
            }
            if (int16(v2) < -1) {
                for fi46 := 0; fi46 < 2; fi46++ {
                    if ((v32 < (v32 / ((v32) | 1))) || (v32 >= -9223372036854775808)) {
                        for fi47 := 0; fi47 < 5; fi47++ {
                            v6 = ((255 <= 42) || v6)
                            v48 := GenS4{F0: 0, F1: 255}
                            var v49 uint8 = v8
                            if len(v7) >= 2 {
                                v7 = v7[1:len(v7)]
                            }
                            _ = fi47
                            _ = v48
                            _ = v49
                        }
                        if v6 {
                            fmt.Println(v32, v40, v2)
                            v50 := GenS5{F0: (v32 >> ((4294967295) & 63)), F1: fi46, F2: true}
                            _ = v50
                        } else {
                            _ = len(v7)
                        }
                        if (894 <= 0) {
                            v3 += v31
                        }
                    } else {
                        v6 = v6
                        v7 = append(v7, fi46)
                        v51 := GenS4{F0: 556, F1: 32767}
                        _ = v51
                    }
                    v8 -= 2
                    _ = fi46
                }
                v52 := GenS5{F0: v32, F1: 4294967295, F2: false}
                for fi53 := 0; fi53 < 12; fi53++ {
                    if (v31 >= v3) {
                        v32 -= (^(v32))
                    } else {
                        _ = len(v7)
                        if true {
                            if true {
                                v32 = v32
                                for fi54 := 0; fi54 < 19; fi54++ {
                                    v1 = v1
                                    if v6 {
                                        if len(v7) >= 2 {
                                            v7 = v7[1:len(v7)]
                                        }
                                        var v55 uint32 = uint32(v8)
                                        fmt.Println(v1)
                                        v56 := GenS6{F0: 42}
                                        v5 += v5
                                        _ = v55
                                        _ = v56
                                    }
                                    _ = fi54
                                }
                                if len(v7) > 0 {
                                    v7[0] = -854
                                }
                            }
                        } else {
                            if len(v7) >= 2 {
                                v7 = v7[1:len(v7)]
                            }
                            if v6 {
                                fmt.Println(v3)
                                if v6 {
                                    v2 = (^(v2))
                                    if len(v7) >= 2 {
                                        v7 = v7[1:len(v7)]
                                    }
                                } else {
                                    v32 = (v32 - v32)
                                    if true {
                                        var v57 uint64 = 149
                                        v57 -= (v57 | (v57 | 0))
                                        v1 += (v5 * v1)
                                        _ = v57
                                    }
                                    v1 += v1
                                }
                                var v58 uint64 = 962
                                _ = v58
                            } else {
                                v59 := GenS6{F0: 1}
                                for fi60 := 0; fi60 < 16; fi60++ {
                                    var v61 bool = v6
                                    if (!true) {
                                        var v62 uint = 965
                                        var v63 int16 = -32768
                                        _ = v62
                                        _ = v63
                                    } else {
                                        if (v6 && v61) {
                                            v2 = 6
                                        }
                                        v3 = v31
                                    }
                                    _ = fi60
                                    _ = v61
                                }
                                var v64 uint8 = v8
                                _ = v59
                                _ = v64
                            }
                        }
                    }
                    v3 += (v31 << ((4294967295) & 63))
                    v65 := GenS1{F0: 846, F1: (^(v8))}
                    _ = fi53
                    _ = v65
                }
                _ = v52
            } else {
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
            }
        } else {
            v1 = 1e10
            var v66 float32 = ((float32(1e10) / 0.1) + (v1 - (v1 + 0.1)))
            v67 := GenS7{F0: 4294967295}
            _ = v66
            _ = v67
        }
        _ = v40
    }
    for fi68 := 0; fi68 < 14; fi68++ {
        var v69 float32 = ((0.0 + v5) * (((float32(1.5) - 1.5) * -2.25) + ((v5 * v1) - v1)))
        _ = fi68
        _ = v69
    }
    fmt.Println(v6, v8)
    v7 = append(v7, int(v8))
    for fi70 := 0; fi70 < 11; fi70++ {
        if v6 {
            v3 -= (v31 >> ((413) & 63))
            if (!(0.0 >= (v1 / 1.0))) {
                if (!v6) {
                    v71 := GenS8{F0: ((^(v2)) <= (v2 & 2)), F1: v31}
                    _ = v71
                } else {
                    if (false || v6) {
                        v2 = v2
                    }
                    v5 -= 1e10
                }
                var v72 int8 = v31
                v3 = (v72 >> ((4294967295) & 63))
                _ = v72
            } else {
                if (("a" + ((((("hello" + "foo bar") + "hello") + "hello") + ("foo bar" + "foo bar")) + "hello")) < "foo bar") {
                    _ = len(v7)
                } else {
                    if (v5 < 0.1) {
                        var v73 uint8 = (v8 | 112)
                        if v6 {
                            for fi74 := 0; fi74 < 8; fi74++ {
                                v31 -= 27
                                var v75 uint32 = v2
                                _ = fi74
                                _ = v75
                            }
                            var v76 int8 = v31
                            _ = v76
                        }
                        if (32767 <= 255) {
                            fmt.Println(v73, fi70, v3)
                        }
                        _ = v73
                    }
                }
            }
            _ = len(v7)
        } else {
            if (v6 && (-1 <= -1)) {
                v7 = append(v7, (fi70 + 2147483647))
            }
        }
        _ = fi70
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
