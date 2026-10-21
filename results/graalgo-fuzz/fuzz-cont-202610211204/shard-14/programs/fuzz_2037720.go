package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16; F2 uint }
type GenS2 struct { F0 uint; F1 int16 }
type GenS3 struct { F0 int16; F1 string }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint8; F1 int16; F2 string }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 bool; F1 int16; F2 string }
type GenS9 struct { F0 uint16 }
type GenS10 struct { F0 uint32; F1 uint8 }
func genFn0() string {
    v6 := []int{128}
    for fi7 := 0; fi7 < 10; fi7++ {
        for fi8 := 0; fi8 < 7; fi8++ {
            var v9 int32 = 1
            if true {
                switch fi8 {
                case -429:
                    var v10 int8 = -1
                    fmt.Println(fi8, v10)
                    _ = v10
                case -1:
                    fmt.Println(fi7, fi7)
                default:
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                }
                switch (fi7 | -9223372036854775808) {
                case 1:
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                default:
                    var v11 uint = 65535
                    _ = v11
                }
            }
            v9 = -84
            _ = fi8
            _ = v9
        }
        _ = fi7
    }
    v6 = append(v6, 128)
    var v12 int16 = 1
    var v13 uint64 = 2
    _ = v12
    _ = v13
    _ = v6
    return ("hello" + ("a" + "\n\t"))
}
func genFn1() string {
    for fi14 := 0; fi14 < 7; fi14++ {
        v15 := GenS1{F0: ((fi14 ^ (-(fi14))) > (fi14 >> ((18446744073709551615) & 63))), F1: -32768, F2: 793}
        switch "a" {
        case "":
            switch 0 {
            case 9223372036854775807:
                var v16 uint64 = 2
                _ = v16
            case 42:
                v17 := GenS2{F0: 2, F1: -32768}
                _ = v17
            case 32768:
                var v18 uint16 = 25
                _ = v18
            default:
                v19 := GenS2{F0: 0, F1: 32767}
                _ = v19
            }
        case "\n\t":
            v20 := []int{}
            _ = v20
        default:
            fmt.Println(fi14)
        }
        switch fi14 {
        case 128:
            if (false || true) {
                var v21 float64 = 1e10
                _ = v21
            }
        case 32768:
            var v22 int16 = -1
            _ = v22
        case 42:
            if false {
                var v23 int64 = 989
                _ = v23
            }
        }
        _ = fi14
        _ = v15
    }
    var v24 uint16 = 65535
    for fi25 := 0; fi25 < 2; fi25++ {
        switch "" {
        case "hello":
            if false {
                if false {
                    var v26 float64 = 3.14159
                    for fi27 := 0; fi27 < 7; fi27++ {
                        switch (-(fi27)) {
                        case -653:
                            fmt.Println(fi27, fi27)
                        }
                        switch "" {
                        case "foo bar":
                            if (-697 == -1) {
                                if true {
                                    switch "" {
                                    case "hello":
                                        v28 := GenS3{F0: 0, F1: "foo bar"}
                                        _ = v28
                                    case "foo bar":
                                        if false {
                                            v29 := []int{(fi27 + fi27), fi27, fi27}
                                            _ = v29
                                        }
                                    case "":
                                        v30 := GenS2{F0: 721, F1: -328}
                                        _ = v30
                                    }
                                } else {
                                    v31 := GenS4{F0: 255}
                                    v32 := GenS2{F0: 209, F1: 32767}
                                    v33 := GenS3{F0: -32768, F1: "foo bar"}
                                    _ = v31
                                    _ = v32
                                    _ = v33
                                }
                                v34 := []int{-177}
                                var v35 int8 = 42
                                v34 = append(v34, (fi27 >> ((4294967295) & 63)))
                                v35 += 127
                                _ = v34
                                _ = v35
                            }
                        case "":
                            if (!(!(255 > 65535))) {
                                var v36 uint = 4294967295
                                _ = v36
                            }
                            v26 += v26
                        case "\n\t":
                            fmt.Println(v26, fi25)
                        default:
                            if true {
                                switch "a" {
                                case "hello":
                                    switch -229 {
                                    case -661:
                                        fmt.Println(fi27, v26, v26)
                                    case 32767:
                                        if (("a" + "foo bar") < "foo bar") {
                                            v37 := []int{}
                                            if false {
                                                v24 = v24
                                                v37 = append(v37, (fi27 * fi25))
                                                if ((-1 >= -128) || false) {
                                                    var v38 float64 = -2.25
                                                    _ = len(v37)
                                                    _ = v38
                                                }
                                                switch (("a" + ("\n\t" + ("\n\t" + "foo bar"))) + ("hello" + ("\n\t" + "\n\t"))) {
                                                case "":
                                                    v26 = (v26 + (v26 * 0.1))
                                                default:
                                                    fmt.Println(fi25)
                                                }
                                            } else {
                                                if true {
                                                    if len(v37) >= 2 {
                                                        v37 = v37[1:len(v37)]
                                                    }
                                                }
                                                v39 := GenS2{F0: 192, F1: -939}
                                                v24 += (v24 % (((^(v24))) | 1))
                                                _ = v39
                                            }
                                            if (v26 > v26) {
                                                v26 = -2.25
                                            } else {
                                                if (fi27 == (fi25 ^ 686)) {
                                                    v40 := GenS3{F0: 1, F1: ""}
                                                    var v41 int32 = -1
                                                    _ = len(v37)
                                                    fmt.Println(fi25)
                                                    _ = v40
                                                    _ = v41
                                                } else {
                                                    var v42 int8 = 100
                                                    _ = v42
                                                }
                                            }
                                            _ = v37
                                        }
                                    default:
                                        var v43 int = (-(fi27))
                                        _ = v43
                                    }
                                case "foo bar":
                                    fmt.Println(v26, fi27, fi27)
                                }
                            } else {
                                v44 := []int{fi25, fi25}
                                if len(v44) >= 2 {
                                    v44 = v44[1:len(v44)]
                                }
                                if len(v44) > 0 {
                                    v44[0] = fi25
                                }
                                v24 = 65535
                                v24 -= (v24 % ((v24) | 1))
                                _ = v44
                            }
                        }
                        switch -192 {
                        case 1:
                            v26 -= ((float64((0.0 - (0.1 * 1.5))) - v26) * v26)
                        case 9223372036854775807:
                            v26 += v26
                        case 4294967295:
                            switch fi27 {
                            case -9223372036854775808:
                                v45 := GenS5{F0: true}
                                _ = v45
                            }
                        }
                        _ = fi27
                    }
                    _ = v26
                } else {
                    v46 := []int{-507, fi25, (fi25 % (((-(fi25))) | 1)), (fi25 & (-(fi25)))}
                    if (!(false && (1 <= 90))) {
                        var v47 int8 = 0
                        _ = v47
                    }
                    _ = v46
                }
                for fi48 := 0; fi48 < 19; fi48++ {
                    var v49 int64 = -9223372036854775808
                    var v50 int8 = -59
                    if true {
                        if ("" != "a") {
                            v51 := []int{(fi25 + fi25), -9223372036854775808, fi48}
                            _ = v51
                        } else {
                            fmt.Println(fi48, v24)
                            if true {
                                v24 += v24
                            }
                            var v52 int8 = -1
                            _ = v52
                        }
                    } else {
                        switch (fi48 - (fi25 | (fi48 >> ((314) & 63)))) {
                        case 127:
                            var v53 int8 = 0
                            _ = v53
                        }
                        v24 = v24
                    }
                    _ = fi48
                    _ = v49
                    _ = v50
                }
                fmt.Println(fi25, fi25)
            }
        }
        v54 := []int{}
        v55 := GenS4{F0: v24}
        _ = fi25
        _ = v54
        _ = v55
    }
    v56 := []int{-103, int(v24)}
    _ = v24
    _ = v56
    return "foo bar"
}

func main() {
    var v1 int32 = -1
    v2 := make(map[int]bool)
    v3 := []string{"foo bar", "\n\t", ("" + "foo bar"), "a"}
    var v4 uint16 = 255
    var v5 int32 = v1
    var v57 float32 = float32(float64(-2.25))
    v5 -= 0
    v58 := GenS6{F0: 127, F1: -691, F2: "hello"}
    for fi59 := 0; fi59 < 4; fi59++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi59
    }
    v3 = append(v3, ("\n\t" + "\n\t"))
    _ = len(v3)
    var v60 int32 = (v5 / ((v1) | 1))
    for fi61 := 0; fi61 < 8; fi61++ {
        fmt.Println(v5)
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi61
    }
    fmt.Println(v60, v60)
    if ((-658 == 370) && ((v60 - (v60 >> ((255) & 63))) <= (v60 / (((v1 * (v1 ^ (v60 >> ((433) & 63))))) | 1)))) {
        if len(v3) > 0 {
            v3[0] = "foo bar"
        }
    } else {
        for fi62 := 0; fi62 < 4; fi62++ {
            var v63 uint16 = 255
            v64 := GenS7{F0: 905}
            for fi65 := 0; fi65 < 16; fi65++ {
                v57 = 0.0
                v66 := GenS8{F0: (((float32((float64((3.14159 / float32(0.1))) + float64(3.14159))) + v57) - (1.5 / -2.25)) != float32(0.1)), F1: 1, F2: ((("hello" + ("hello" + "\n\t")) + "\n\t") + "")}
                _ = fi65
                _ = v66
            }
            _ = fi62
            _ = v63
            _ = v64
        }
    }
    var v67 string = "foo bar"
    v67 = v67
    fmt.Println(v60)
    var v68 uint8 = 186
    if (!true) {
        if false {
            for fi69 := 0; fi69 < 15; fi69++ {
                v67 = v67
                _ = fi69
            }
            switch -126 {
            case 42:
                switch -115 {
                case 9223372036854775807:
                    if true {
                        var v70 int16 = 0
                        _ = v70
                    } else {
                        if ((0.1 * (float64((float32(float64(v57)) / (v57 / 1e10))) * 0.1)) < float64(1e10)) {
                            if (1 <= 2147483648) {
                                v71 := v2[128]
                                _ = v71
                            }
                            var v72 bool = (1 < -661)
                            _ = v72
                        }
                        fmt.Println(v68)
                        for fi73 := 0; fi73 < 3; fi73++ {
                            v3 = append(v3, ("foo bar" + (v67 + v67)))
                            for fi74 := 0; fi74 < 20; fi74++ {
                                v75 := GenS9{F0: (v4 - 686)}
                                v76 := GenS10{F0: 0, F1: v68}
                                _ = fi74
                                _ = v75
                                _ = v76
                            }
                            _ = fi73
                        }
                    }
                case -2147483648:
                    v57 = 1e10
                }
            case 127:
                switch 32767 {
                case 128:
                    fmt.Println(v5, v4)
                default:
                    v4 -= (v4 + (^(v4)))
                }
            }
            var v77 int8 = 1
            _ = v77
        }
        if len(v3) > 0 {
            v3[0] = v67
        }
        var v78 int64 = -9223372036854775808
        _ = v78
    }
    fmt.Println(v4, v67)
    if len(v3) > 0 {
        v3[0] = "a"
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v57:", v57)
    fmt.Println("v60:", v60)
    fmt.Println("v67:", v67)
    fmt.Println("v68:", v68)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v57
    _ = v58
    _ = v60
    _ = v67
    _ = v68
}
