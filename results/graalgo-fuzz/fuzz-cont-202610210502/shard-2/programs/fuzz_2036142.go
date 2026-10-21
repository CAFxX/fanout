package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int }
type GenS2 struct { F0 bool; F1 uint32 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int; F1 int32; F2 uint32 }
type GenS5 struct { F0 uint32; F1 string; F2 int8 }
type GenS6 struct { F0 int; F1 uint32; F2 int64 }
type GenS7 struct { F0 uint; F1 uint32 }
func genFn0() uint32 {
    if false {
        var v7 bool = (-1 > 9223372036854775807)
        _ = v7
    } else {
        for fi8 := 0; fi8 < 19; fi8++ {
            v9 := []int{1, (fi8 - fi8)}
            switch fi8 {
            case 2147483647:
                v9 = append(v9, -2)
            case 32767:
                if false {
                    var v10 float32 = -2.25
                    for fi11 := 0; fi11 < 4; fi11++ {
                        v10 += v10
                        if false {
                            switch "a" {
                            case "":
                                if false {
                                    fmt.Println(fi8, v10)
                                    v12 := GenS1{F0: int32(fi11), F1: fi8}
                                    var v13 float32 = 0.1
                                    _ = v12
                                    _ = v13
                                }
                            case "\n\t":
                                v14 := GenS1{F0: 42, F1: -989}
                                _ = v14
                            case "hello":
                                v15 := GenS2{F0: (v10 < v10), F1: 42}
                                _ = v15
                            }
                        }
                        v10 = v10
                        _ = fi11
                    }
                    _ = v10
                }
            }
            _ = fi8
            _ = v9
        }
    }
    v16 := []int{42}
    if (!(483 != 42)) {
        for fi17 := 0; fi17 < 12; fi17++ {
            v18 := GenS1{F0: 980, F1: (fi17 << ((65535) & 63))}
            var v19 float64 = 1.5
            _ = fi17
            _ = v18
            _ = v19
        }
        var v20 int = 127
        if (!((-1 >= 127) && false)) {
            if true {
                for fi21 := 0; fi21 < 4; fi21++ {
                    if len(v16) > 0 {
                        v16[0] = 2147483647
                    }
                    _ = fi21
                }
                var v22 uint = 896
                if (85 >= -1) {
                    if true {
                        fmt.Println(v22, v20)
                        var v23 int16 = 0
                        for fi24 := 0; fi24 < 3; fi24++ {
                            switch -477 {
                            case -464:
                                if len(v16) > 0 {
                                    v16[0] = (fi24 / (((-(v20))) | 1))
                                }
                            default:
                                for fi25 := 0; fi25 < 10; fi25++ {
                                    v16 = append(v16, (^(fi25)))
                                    _ = fi25
                                }
                            }
                            if len(v16) >= 2 {
                                v16 = v16[1:len(v16)]
                            }
                            _ = fi24
                        }
                        _ = v23
                    }
                    for fi26 := 0; fi26 < 5; fi26++ {
                        fmt.Println(fi26, v20, fi26)
                        fmt.Println(v20)
                        _ = fi26
                    }
                } else {
                    var v27 int16 = 32767
                    switch "foo bar" {
                    case "\n\t":
                        var v28 int16 = v27
                        _ = v28
                    case "":
                        if true {
                            v20 = (v20 & (v20 % (((v20 ^ v20)) | 1)))
                            if false {
                                for fi29 := 0; fi29 < 15; fi29++ {
                                    v22 += (v22 - v22)
                                    v30 := GenS2{F0: false, F1: 298}
                                    for fi31 := 0; fi31 < 8; fi31++ {
                                        if len(v16) > 0 {
                                            v16[0] = (fi29 >> (((v22 / (((v22 & 149)) | 1))) & 63))
                                        }
                                        _ = fi31
                                    }
                                    _ = fi29
                                    _ = v30
                                }
                                v22 -= (v22 / ((v22) | 1))
                            } else {
                                if false {
                                    var v32 int8 = 42
                                    var v33 int8 = v32
                                    _ = v32
                                    _ = v33
                                } else {
                                    for fi34 := 0; fi34 < 19; fi34++ {
                                        v22 -= (v22 % ((v22) | 1))
                                        _ = fi34
                                    }
                                    fmt.Println(v22, v27)
                                }
                                fmt.Println(v20, v20, v27)
                                for fi35 := 0; fi35 < 14; fi35++ {
                                    if true {
                                        for fi36 := 0; fi36 < 19; fi36++ {
                                            if (("foo bar" + ("foo bar" + "")) >= "") {
                                                v37 := GenS3{F0: true}
                                                var v38 uint = v22
                                                _ = v37
                                                _ = v38
                                            }
                                            var v39 uint16 = 65535
                                            switch (fi36 & fi35) {
                                            case -930:
                                                if len(v16) > 0 {
                                                    v16[0] = fi36
                                                }
                                            case -1:
                                                v40 := GenS4{F0: (fi36 / ((int(v39)) | 1)), F1: -2147483648, F2: 2}
                                                _ = v40
                                            case 32768:
                                                fmt.Println(v20)
                                            }
                                            _ = fi36
                                            _ = v39
                                        }
                                        fmt.Println(v20, v20, v20)
                                    } else {
                                        if (!(65535 != 905)) {
                                            if true {
                                                fmt.Println(fi35, v22, fi35)
                                            }
                                        } else {
                                            var v41 int32 = 42
                                            v16 = append(v16, (-(fi35)))
                                            v16 = append(v16, (^(v20)))
                                            _ = v41
                                        }
                                    }
                                    _ = fi35
                                }
                            }
                            if true {
                                for fi42 := 0; fi42 < 15; fi42++ {
                                    if len(v16) > 0 {
                                        v16[0] = fi42
                                    }
                                    v43 := GenS4{F0: (^(fi42)), F1: 423, F2: 751}
                                    _ = fi42
                                    _ = v43
                                }
                                if false {
                                    _ = len(v16)
                                    for fi44 := 0; fi44 < 15; fi44++ {
                                        var v45 uint64 = 42
                                        _ = len(v16)
                                        switch (-(fi44)) {
                                        case 42:
                                            v46 := GenS5{F0: 376, F1: "foo bar", F2: 42}
                                            _ = v46
                                        case -638:
                                            _ = len(v16)
                                        case 4294967295:
                                            if ((0.1 - (float32((-2.25 + (float64(float32(3.14159)) * 0.0))) - ((-2.25 / (0.0 - 1e10)) * (0.1 - (-2.25 + 1.5))))) > 3.14159) {
                                                var v47 int = (v20 | 127)
                                                if len(v16) >= 2 {
                                                    v16 = v16[1:len(v16)]
                                                }
                                                _ = v47
                                            }
                                        }
                                        _ = fi44
                                        _ = v45
                                    }
                                    var v48 float64 = 3.14159
                                    _ = v48
                                } else {
                                    for fi49 := 0; fi49 < 14; fi49++ {
                                        for fi50 := 0; fi50 < 14; fi50++ {
                                            var v51 int32 = int32(v27)
                                            v52 := GenS2{F0: false, F1: 42}
                                            fmt.Println(v22, v27, v22)
                                            if (true || (1e10 != 3.14159)) {
                                                switch fi50 {
                                                case -1:
                                                    var v53 float64 = float64((float32(-2.25) / 0.1))
                                                    _ = v53
                                                case 32767:
                                                    var v54 string = ((((("" + "") + "foo bar") + (("hello" + "") + ((("\n\t" + ("" + "foo bar")) + "a") + "foo bar"))) + ("" + "hello")) + "foo bar")
                                                    if len(v16) > 0 {
                                                        v16[0] = 412
                                                    }
                                                    _ = v54
                                                case -9223372036854775808:
                                                    fmt.Println(v20, v20)
                                                }
                                                v27 -= int16(v22)
                                                if ((759 <= 1) || true) {
                                                    v20 = (v20 & fi50)
                                                } else {
                                                    fmt.Println(fi49, v22)
                                                    switch "hello" {
                                                    case "a":
                                                        var v55 uint8 = 103
                                                        _ = v55
                                                    case "foo bar":
                                                        if len(v16) >= 2 {
                                                            v16 = v16[1:len(v16)]
                                                        }
                                                    }
                                                    if false {
                                                        var v56 uint8 = 128
                                                        v51 -= v51
                                                        _ = v56
                                                    }
                                                }
                                            }
                                            v16 = append(v16, (-(fi49)))
                                            _ = fi50
                                            _ = v51
                                            _ = v52
                                        }
                                        _ = fi49
                                    }
                                    switch v20 {
                                    case 42:
                                        var v57 uint8 = 2
                                        _ = v57
                                    case 0:
                                        switch v20 {
                                        case 4294967295:
                                            fmt.Println(v22, v20)
                                        case -1:
                                            v22 += v22
                                        case -2147483648:
                                            v20 = (v20 & v20)
                                        }
                                    case 32768:
                                        fmt.Println(v20, v22, v20)
                                    default:
                                        var v58 int16 = (v27 - v27)
                                        _ = v58
                                    }
                                    for fi59 := 0; fi59 < 4; fi59++ {
                                        v27 -= v27
                                        v22 -= (v22 & (v22 + v22))
                                        for fi60 := 0; fi60 < 12; fi60++ {
                                            v61 := GenS2{F0: true, F1: 156}
                                            if len(v16) >= 2 {
                                                v16 = v16[1:len(v16)]
                                            }
                                            var v62 uint8 = 0
                                            _ = fi60
                                            _ = v61
                                            _ = v62
                                        }
                                        _ = fi59
                                    }
                                }
                            }
                        }
                    case "a":
                        for fi63 := 0; fi63 < 12; fi63++ {
                            var v64 uint64 = 42
                            for fi65 := 0; fi65 < 1; fi65++ {
                                var v66 int16 = (v27 / ((85) | 1))
                                _ = fi65
                                _ = v66
                            }
                            if (570 >= 0) {
                                if len(v16) >= 2 {
                                    v16 = v16[1:len(v16)]
                                }
                            }
                            _ = fi63
                            _ = v64
                        }
                    }
                    _ = v27
                }
                _ = v22
            }
            if len(v16) > 0 {
                v16[0] = -9223372036854775808
            }
        } else {
            v20 = (^(v20))
        }
        _ = v20
    }
    var v67 float32 = ((1e10 * 0.0) * (((-2.25 / (-2.25 * ((1.5 - 0.1) / 1.5))) + ((1.5 - 0.0) * float32(0.1))) / 1.0))
    fmt.Println(v67)
    _ = v16
    _ = v67
    return 2147483648
}
func genFn1() bool {
    switch (("\n\t" + "a") + "hello") {
    case "a":
        if false {
            switch "foo bar" {
            case "\n\t":
                var v68 int16 = 428
                _ = v68
            case "a":
                var v69 int16 = 44
                _ = v69
            }
        }
    case "hello":
        v70 := GenS2{F0: (((255 <= 844) && false) && (32767 >= 0)), F1: 2147483647}
        _ = v70
    case "foo bar":
        v71 := []int{42, 127, 1, 0}
        _ = v71
    }
    v72 := GenS1{F0: -655, F1: 32768}
    v73 := []int{0, -676, -729, -646}
    var v74 uint32 = genFn0()
    _ = v72
    _ = v73
    _ = v74
    return true
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint64 = 1
    var v3 uint8 = 196
    var v4 uint = 576
    var v5 float32 = (0.1 - (0.1 * -2.25))
    v6 := make(map[string]bool)
    var v75 float64 = 1.5
    v76 := GenS6{F0: 42, F1: uint32(v4), F2: 716}
    v77 := []int{-2147483648}
    fmt.Println(v75, v5)
    var v78 int16 = 42
    fmt.Println(v4, v2, v3)
    var v79 float32 = (float32(v75) - 0.1)
    v80 := GenS7{F0: v4, F1: 2147483648}
    v77 = append(v77, 4294967295)
    v2 = v2
    var v81 uint32 = 520
    for fi82 := 0; fi82 < 14; fi82++ {
        for fi83 := 0; fi83 < 16; fi83++ {
            v81 -= v81
            _ = fi83
        }
        _ = fi82
    }
    if len(v77) >= 2 {
        v77 = v77[1:len(v77)]
    }
    v77 = append(v77, 2147483647)
    var v84 uint8 = v3
    var v85 uint = v4
    var v86 float64 = v75
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v75:", v75)
    fmt.Println("slicelen v77:", len(v77), v77)
    fmt.Println("v78:", v78)
    fmt.Println("v79:", v79)
    fmt.Println("v81:", v81)
    fmt.Println("v84:", v84)
    fmt.Println("v85:", v85)
    fmt.Println("v86:", v86)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v75
    _ = v76
    _ = v77
    _ = v78
    _ = v79
    _ = v80
    _ = v81
    _ = v84
    _ = v85
    _ = v86
}
