package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8 }
type GenS2 struct { F0 uint16; F1 uint64; F2 uint32 }
type GenS3 struct { F0 uint8; F1 int64 }
type GenS4 struct { F0 uint32; F1 int64; F2 uint16 }
type GenS5 struct { F0 int16; F1 bool }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 int16; F1 int16 }
type GenS8 struct { F0 int }
type GenS9 struct { F0 int64; F1 uint32; F2 uint16 }
func genFn0(p0 bool, p1 int64, p2 uint64) int8 {
    v5 := []int{42, 1, 739}
    fmt.Println(p2, p0, p1)
    var v6 uint = 255
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 42
}

func main() {
    var v1 uint8 = 0
    var v2 int64 = 9223372036854775807
    var v3 string = "hello"
    var v4 uint16 = 65535
    if false {
        v7 := GenS1{F0: 32768, F1: -25}
        var v8 int32 = 0
        for fi9 := 0; fi9 < 5; fi9++ {
            v10 := GenS2{F0: v4, F1: 0, F2: 1}
            v11 := GenS1{F0: fi9, F1: -71}
            _ = fi9
            _ = v10
            _ = v11
        }
        _ = v7
        _ = v8
    }
    var v12 int8 = genFn0(false, v2, 0)
    if false {
        v1 -= v1
    }
    v13 := []int{2147483647, 0, -303, 794}
    v13 = append(v13, 128)
    for fi14 := 0; fi14 < 3; fi14++ {
        if (true || true) {
            if false {
                if false {
                    v2 -= (v2 & 0)
                    for fi15 := 0; fi15 < 6; fi15++ {
                        v16 := GenS1{F0: -2147483648, F1: v12}
                        var v17 string = v3
                        v4 += (^(v4))
                        _ = fi15
                        _ = v16
                        _ = v17
                    }
                    if len(v13) >= 2 {
                        v13 = v13[1:len(v13)]
                    }
                } else {
                    v3 = "foo bar"
                }
            }
            v3 = ((v3 + (v3 + "\n\t")) + "")
            if len(v13) > 0 {
                v13[0] = fi14
            }
        } else {
            fmt.Println(fi14, v3)
            switch v3 {
            case "a":
                for fi18 := 0; fi18 < 17; fi18++ {
                    if (918 < 42) {
                        var v19 uint8 = v1
                        switch 397 {
                        case 4294967295:
                            fmt.Println(v4)
                        case -288:
                            if len(v13) >= 2 {
                                v13 = v13[1:len(v13)]
                            }
                        case -2147483648:
                            if true {
                                if len(v13) >= 2 {
                                    v13 = v13[1:len(v13)]
                                }
                                var v20 int16 = 1
                                _ = v20
                            }
                        default:
                            if (64 < 2) {
                                if len(v13) >= 2 {
                                    v13 = v13[1:len(v13)]
                                }
                                var v21 string = (v3 + v3)
                                switch fi14 {
                                case 535:
                                    if false {
                                        var v22 bool = false
                                        _ = v22
                                    } else {
                                        v23 := GenS2{F0: 897, F1: 4294967295, F2: 4294967295}
                                        if true {
                                            if true {
                                                var v24 uint = 206
                                                v24 += uint(v1)
                                                if false {
                                                    if len(v13) >= 2 {
                                                        v13 = v13[1:len(v13)]
                                                    }
                                                    switch (v21 + v21) {
                                                    case "\n\t":
                                                        if len(v13) > 0 {
                                                            v13[0] = (fi18 | (fi14 + 2147483647))
                                                        }
                                                    }
                                                    if (false && false) {
                                                        if true {
                                                            if (-2.25 == -2.25) {
                                                                v25 := GenS1{F0: fi14, F1: -128}
                                                                _ = v25
                                                            }
                                                        } else {
                                                            if (!(!false)) {
                                                                if len(v13) > 0 {
                                                                    v13[0] = fi18
                                                                }
                                                                v4 -= v4
                                                                var v26 int32 = 1
                                                                _ = v26
                                                            }
                                                        }
                                                        if len(v13) > 0 {
                                                            v13[0] = (fi18 & 42)
                                                        }
                                                        v27 := GenS1{F0: 42, F1: v12}
                                                        _ = v27
                                                    } else {
                                                        v1 -= v1
                                                        v12 = -77
                                                    }
                                                } else {
                                                    switch (fi14 / ((fi14) | 1)) {
                                                    case -381:
                                                        fmt.Println(v21)
                                                    default:
                                                        switch (v3 + v3) {
                                                        case "a":
                                                            switch (("foo bar" + "hello") + ((v3 + "a") + v21)) {
                                                            case "a":
                                                                if ((v24 & (v24 & 4294967295)) <= 2) {
                                                                    var v28 int64 = 417
                                                                    var v29 int8 = (v12 & (v12 ^ (v12 % (((v12 & v12)) | 1))))
                                                                    _ = v28
                                                                    _ = v29
                                                                }
                                                            default:
                                                                if false {
                                                                    v30 := GenS3{F0: 42, F1: (v2 * (v2 % (((v2 + (v2 ^ (-(v2))))) | 1)))}
                                                                    _ = v30
                                                                } else {
                                                                    v31 := GenS4{F0: 693, F1: (v2 >> ((888) & 63)), F2: v4}
                                                                    var v32 int8 = int8(v24)
                                                                    _ = v31
                                                                    _ = v32
                                                                }
                                                            }
                                                        case "\n\t":
                                                            if (v2 >= (v2 << ((18446744073709551615) & 63))) {
                                                                var v33 uint64 = 42
                                                                v34 := GenS5{F0: 98, F1: true}
                                                                v35 := GenS5{F0: 1000, F1: true}
                                                                _ = v33
                                                                _ = v34
                                                                _ = v35
                                                            }
                                                        default:
                                                            fmt.Println(fi18, v4)
                                                        }
                                                    }
                                                }
                                                v36 := GenS5{F0: -354, F1: ((-382 >= -650) || true)}
                                                v37 := GenS2{F0: 1, F1: 735, F2: 4294967295}
                                                var v38 uint32 = 0
                                                _ = v24
                                                _ = v36
                                                _ = v37
                                                _ = v38
                                            } else {
                                                if len(v13) > 0 {
                                                    v13[0] = fi18
                                                }
                                            }
                                            fmt.Println(v12, fi14)
                                            var v39 float64 = 3.14159
                                            v12 += 100
                                            _ = len(v13)
                                            _ = v39
                                        } else {
                                            if (3.14159 <= 1e10) {
                                                v21 = (v21 + ("foo bar" + v21))
                                                switch fi14 {
                                                case -799:
                                                    v4 = (v4 >> ((255) & 63))
                                                case 127:
                                                    var v40 uint32 = 218
                                                    _ = v40
                                                case -305:
                                                    var v41 bool = (!((v12 + v12) == v12))
                                                    _ = v41
                                                default:
                                                    v1 = v19
                                                }
                                                if false {
                                                    v42 := GenS5{F0: -966, F1: true}
                                                    _ = v42
                                                } else {
                                                    if (float64((1e10 - 1e10)) > float64(-2.25)) {
                                                        if (!(true && false)) {
                                                            fmt.Println(v3, v4, v2)
                                                            v19 += 215
                                                            fmt.Println(v2, v21, v1)
                                                        }
                                                        v43 := GenS6{F0: 917}
                                                        _ = v43
                                                    } else {
                                                        v12 += -1
                                                        if ((1.5 / 1.0) == 0.1) {
                                                            if (true && (3.14159 < 0.1)) {
                                                                if true {
                                                                    var v44 float32 = float32(0.1)
                                                                    v2 -= v2
                                                                    var v45 uint16 = v4
                                                                    v44 -= v44
                                                                    if (uint16(v19) < (^(v45))) {
                                                                        var v46 int8 = (v12 >> ((42) & 63))
                                                                        if true {
                                                                            switch (fi14 << ((1) & 63)) {
                                                                            case 476:
                                                                                v12 += v12
                                                                            case 4294967295:
                                                                                var v47 int = fi14
                                                                                _ = v47
                                                                            case -188:
                                                                                v12 += (v12 << ((80) & 63))
                                                                            }
                                                                            if len(v13) > 0 {
                                                                                v13[0] = -1
                                                                            }
                                                                        } else {
                                                                            var v48 bool = (false || false)
                                                                            v19 = v19
                                                                            v49 := GenS2{F0: 390, F1: 18446744073709551615, F2: uint32(v4)}
                                                                            _ = v48
                                                                            _ = v49
                                                                        }
                                                                        _ = v46
                                                                    }
                                                                    _ = v44
                                                                    _ = v45
                                                                }
                                                                v13 = append(v13, 1)
                                                            }
                                                        } else {
                                                            if len(v13) > 0 {
                                                                v13[0] = 128
                                                            }
                                                        }
                                                    }
                                                    if (0 < 1) {
                                                        var v50 uint32 = 628
                                                        if true {
                                                            var v51 uint = 990
                                                            var v52 int32 = -1
                                                            _ = v51
                                                            _ = v52
                                                        } else {
                                                            v12 = genFn0(true, -1, 808)
                                                        }
                                                        fmt.Println(fi14, fi14)
                                                        _ = v50
                                                    } else {
                                                        if (!false) {
                                                            v12 -= v12
                                                        }
                                                        if len(v13) >= 2 {
                                                            v13 = v13[1:len(v13)]
                                                        }
                                                        fmt.Println(v2)
                                                    }
                                                    var v53 int16 = -1
                                                    _ = v53
                                                }
                                                var v54 int8 = 100
                                                v55 := GenS3{F0: v1, F1: v2}
                                                _ = v54
                                                _ = v55
                                            } else {
                                                v13 = append(v13, (fi18 | (-(fi14))))
                                            }
                                            if false {
                                                var v56 float32 = float32(3.14159)
                                                _ = v56
                                            }
                                            switch fi18 {
                                            case -802:
                                                switch v21 {
                                                case "\n\t":
                                                    if false {
                                                        v57 := GenS1{F0: fi14, F1: (-(v12))}
                                                        _ = v57
                                                    }
                                                default:
                                                    var v58 int8 = 95
                                                    _ = v58
                                                }
                                            case -1:
                                                fmt.Println(fi18, v4)
                                            default:
                                                var v59 uint = 28
                                                _ = v59
                                            }
                                        }
                                        var v60 int16 = -406
                                        _ = v23
                                        _ = v60
                                    }
                                }
                                _ = v21
                            } else {
                                if false {
                                    switch ("" + ("a" + "a")) {
                                    case "":
                                        v61 := GenS3{F0: v19, F1: v2}
                                        _ = v61
                                    case "a":
                                        var v62 float64 = (((0.0 * (1.5 - -2.25)) + (0.1 / 1.0)) - (1e10 + float64(float32(1.5))))
                                        _ = v62
                                    }
                                    v63 := GenS1{F0: fi18, F1: v12}
                                    _ = v63
                                }
                                v64 := GenS7{F0: 32767, F1: 1000}
                                v19 = 42
                                if true {
                                    fmt.Println(v19, v2, v4)
                                }
                                _ = v64
                            }
                        }
                        v19 += (v19 | 2)
                        _ = v19
                    }
                    switch ("hello" + ((("hello" + "\n\t") + v3) + v3)) {
                    case "hello":
                        if false {
                            v65 := GenS8{F0: fi14}
                            v13 = append(v13, fi18)
                            _ = v65
                        }
                    }
                    _ = fi18
                }
            }
            if (fi14 != (fi14 - fi14)) {
                switch fi14 {
                case 9223372036854775807:
                    v13 = append(v13, (fi14 & (fi14 / ((fi14) | 1))))
                case -875:
                    var v66 bool = true
                    _ = v66
                default:
                    if (!(((!(-1 <= -1)) && false) && (0.0 == float64(-2.25)))) {
                        var v67 bool = true
                        if ((v4 != (v4 << ((4294967295) & 63))) && v67) {
                            for fi68 := 0; fi68 < 3; fi68++ {
                                var v69 int64 = (v2 & -180)
                                _ = fi68
                                _ = v69
                            }
                            var v70 int32 = int32(v4)
                            _ = v70
                        } else {
                            var v71 bool = (v67 || v67)
                            _ = v71
                        }
                        if (v67 && (!v67)) {
                            if v67 {
                                for fi72 := 0; fi72 < 20; fi72++ {
                                    if false {
                                        fmt.Println(v4)
                                        switch ("hello" + (v3 + "\n\t")) {
                                        case "foo bar":
                                            switch fi72 {
                                            case 32768:
                                                if len(v13) >= 2 {
                                                    v13 = v13[1:len(v13)]
                                                }
                                            default:
                                                v73 := GenS7{F0: 585, F1: int16(v4)}
                                                _ = v73
                                            }
                                            if true {
                                                switch v3 {
                                                case "foo bar":
                                                    if true {
                                                        if len(v13) > 0 {
                                                            v13[0] = 0
                                                        }
                                                        fmt.Println(v4)
                                                    }
                                                case "":
                                                    v12 = 1
                                                case "hello":
                                                    if v67 {
                                                        v3 = v3
                                                        v12 -= 42
                                                        v1 -= v1
                                                    }
                                                default:
                                                    v13 = append(v13, -1)
                                                }
                                            } else {
                                                var v74 float64 = 1.5
                                                if (1 < -580) {
                                                    v2 -= (v2 ^ -9223372036854775808)
                                                    var v75 bool = (!(int(v1) >= (fi14 % ((fi72) | 1))))
                                                    v4 -= (v4 >> ((1) & 63))
                                                    _ = v75
                                                }
                                                _ = v74
                                            }
                                        case "a":
                                            v76 := GenS9{F0: v2, F1: 2147483648, F2: (^(v4))}
                                            _ = v76
                                        default:
                                            v2 = 3
                                        }
                                    }
                                    _ = fi72
                                }
                            } else {
                                var v77 uint64 = uint64(v4)
                                var v78 uint16 = (v4 + (^(v4)))
                                _ = v77
                                _ = v78
                            }
                        }
                        _ = v67
                    } else {
                        v12 = v12
                    }
                }
                v1 -= 29
            }
        }
        if len(v13) >= 2 {
            v13 = v13[1:len(v13)]
        }
        var v79 int64 = v2
        _ = fi14
        _ = v79
    }
    fmt.Println(v4, v1, v1)
    v13 = append(v13, 9223372036854775807)
    v2 -= -331
    v12 -= int8(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
}
