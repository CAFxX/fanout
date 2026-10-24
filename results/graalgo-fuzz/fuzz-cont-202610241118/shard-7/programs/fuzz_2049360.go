package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64; F2 uint }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint16; F1 uint8 }
type GenS4 struct { F0 uint16; F1 uint64 }
type GenS5 struct { F0 int32; F1 uint32; F2 int16 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 int8; F1 uint8 }
type GenS8 struct { F0 int16; F1 uint8; F2 int64 }
type GenS9 struct { F0 string; F1 int64; F2 int }
func genFn0(p0 bool, p1 int64) int16 {
    var v4 uint16 = 65535
    p1 -= p1
    for fi5 := 0; fi5 < 18; fi5++ {
        switch ("hello" + (("" + ("foo bar" + "a")) + (("\n\t" + "foo bar") + "a"))) {
        case "":
            switch (("hello" + (((("\n\t" + "a") + ("hello" + ("\n\t" + "hello"))) + "") + "hello")) + ((("\n\t" + "\n\t") + (((("foo bar" + "hello") + ("\n\t" + ("foo bar" + (((("hello" + ("hello" + "hello")) + "a") + (("\n\t" + "") + "")) + ("\n\t" + ("foo bar" + ("a" + (("" + "\n\t") + ("foo bar" + "\n\t"))))))))) + "foo bar") + "a")) + "\n\t")) {
            case "hello":
                v6 := GenS1{F0: 42, F1: 209, F2: uint(v4)}
                _ = v6
            default:
                v7 := []int{fi5, fi5}
                _ = v7
            }
        case "a":
            if p0 {
                if p0 {
                    v8 := GenS1{F0: int8(p1), F1: 706, F2: 42}
                    if p0 {
                        v4 = (v4 & v4)
                        var v9 int64 = p1
                        var v10 uint8 = 2
                        _ = v10
                        _ = v9
                    }
                    if p0 {
                        v11 := GenS2{F0: p1}
                        for fi12 := 0; fi12 < 18; fi12++ {
                            fmt.Println(fi5, fi5, p1)
                            v13 := []int{32767, (fi5 - fi12), (^(fi5)), fi5}
                            p0 = p0
                            _ = fi12
                            _ = v13
                        }
                        _ = v11
                    } else {
                        v14 := GenS3{F0: (v4 ^ (v4 >> ((1) & 63))), F1: 2}
                        fmt.Println(p1)
                        switch ("a" + ("\n\t" + "foo bar")) {
                        case "":
                            switch (fi5 + 106) {
                            case -766:
                                v15 := GenS4{F0: 42, F1: 186}
                                _ = v15
                            default:
                                v4 -= v4
                            }
                        }
                        _ = v14
                    }
                    _ = v8
                } else {
                    v16 := GenS1{F0: -1, F1: 18446744073709551615, F2: 0}
                    v17 := []int{(fi5 % (((fi5 | -349)) | 1)), (fi5 / ((fi5) | 1))}
                    _ = v16
                    _ = v17
                }
                if (!((100 < 42) && p0)) {
                    switch (fi5 * (fi5 >> ((1) & 63))) {
                    case 0:
                        for fi18 := 0; fi18 < 19; fi18++ {
                            v19 := []int{(^(fi18)), (fi5 ^ fi5), (fi5 / (((fi18 | (fi5 / ((fi18) | 1)))) | 1)), fi5}
                            if p0 {
                                var v20 uint8 = 255
                                if (1e10 < 3.14159) {
                                    v21 := GenS5{F0: -1, F1: 42, F2: 777}
                                    switch "" {
                                    case "a":
                                        if len(v19) >= 2 {
                                            v19 = v19[1:len(v19)]
                                        }
                                    case "foo bar":
                                        var v22 int32 = 237
                                        _ = v22
                                    case "":
                                        var v23 int8 = 42
                                        fmt.Println(v23, v4, v4)
                                        _ = v23
                                    }
                                    _ = v21
                                }
                                _ = v20
                            }
                            p0 = (0 <= 2)
                            fmt.Println(fi18)
                            v19 = append(v19, 0)
                            _ = fi18
                            _ = v19
                        }
                    case 937:
                        fmt.Println(p0, v4, fi5)
                    }
                    fmt.Println(p0, p0)
                    switch (fi5 & 128) {
                    case 127:
                        v24 := []int{}
                        _ = v24
                    }
                } else {
                    p1 -= (p1 | (p1 | p1))
                    if false {
                        v25 := GenS4{F0: v4, F1: 817}
                        if (((1.5 + 0.1) - 1.5) <= ((float32((-2.25 + 0.0)) / (0.1 + 1e10)) - (float32(1e10) + 1.5))) {
                            switch fi5 {
                            case 9223372036854775807:
                                v26 := GenS6{F0: p0}
                                _ = v26
                            }
                        }
                        v27 := GenS2{F0: (p1 % (((p1 & int64(fi5))) | 1))}
                        _ = v25
                        _ = v27
                    } else {
                        p0 = (804 == 1)
                    }
                }
                v4 -= v4
            }
        default:
            if (p0 && true) {
                v28 := []int{fi5, fi5, int(p1), (fi5 << ((uint(v4)) & 63))}
                _ = v28
            } else {
                for fi29 := 0; fi29 < 14; fi29++ {
                    var v30 uint = 65535
                    var v31 uint = v30
                    _ = fi29
                    _ = v30
                    _ = v31
                }
                v4 = (v4 + v4)
            }
        }
        v32 := GenS2{F0: 1}
        if (p0 || (!(0.1 == float64(1.5)))) {
            var v33 int32 = -2147483648
            var v34 string = "a"
            _ = v33
            _ = v34
        } else {
            var v35 int32 = -25
            v36 := []int{-303, fi5, (^(fi5)), fi5}
            var v37 uint = 0
            _ = v35
            _ = v36
            _ = v37
        }
        _ = fi5
        _ = v32
    }
    fmt.Println(p1, p0, p1)
    _ = p0
    _ = p1
    _ = v4
    return -244
}

func main() {
    var v1 int32 = 2147483647
    v2 := make(map[int]uint32)
    v3 := make(map[uint32]uint32)
    for fi38 := 0; fi38 < 7; fi38++ {
        var v39 int32 = v1
        var v40 uint8 = 42
        v40 -= (v40 + 128)
        _ = fi38
        _ = v39
        _ = v40
    }
    v3[288] = 65535
    if true {
        var v41 float32 = 3.14159
        for fi42 := 0; fi42 < 12; fi42++ {
            var v43 uint64 = 42
            _ = fi42
            _ = v43
        }
        switch ("" + ("" + "")) {
        case "\n\t":
            if true {
                v44 := []int{-2147483648}
                v41 += 0.0
                _ = v44
            } else {
                if (((!(0 != -390)) || (0.0 != float64((-2.25 + (v41 * v41))))) && ((!(v1 == (v1 - v1))) && (!true))) {
                    for fi45 := 0; fi45 < 3; fi45++ {
                        v46 := []int{(^(fi45)), fi45}
                        for fi47 := 0; fi47 < 7; fi47++ {
                            _ = len(v2)
                            if (false && (-1 == 0)) {
                                var v48 uint64 = 518
                                _ = v48
                            } else {
                                var v49 int32 = (v1 * (v1 >> ((0) & 63)))
                                _ = v49
                            }
                            _ = fi47
                        }
                        _ = fi45
                        _ = v46
                    }
                }
                switch ("\n\t" + "") {
                case "a":
                    v1 = (v1 % ((v1) | 1))
                }
                var v50 uint8 = 255
                _ = v50
            }
        case "a":
            v51 := GenS3{F0: 387, F1: 1}
            _ = v51
        }
        _ = v41
    }
    v1 += (^(v1))
    v52 := GenS4{F0: 407, F1: 42}
    if (42 >= 26) {
        v53 := GenS3{F0: 777, F1: 1}
        switch 2147483647 {
        case 42:
            for fi54 := 0; fi54 < 15; fi54++ {
                switch ("" + "hello") {
                case "hello":
                    v55 := []int{fi54}
                    _ = v55
                }
                for fi56 := 0; fi56 < 5; fi56++ {
                    switch 1 {
                    case -359:
                        v57 := []int{}
                        _ = v57
                    case 55:
                        v1 = (v1 + (v1 / ((v1) | 1)))
                    case 9223372036854775807:
                        v1 -= v1
                    }
                    v1 += 803
                    _ = fi56
                }
                var v58 uint16 = 2
                _ = fi54
                _ = v58
            }
        case -554:
            v1 -= v1
        }
        _ = v53
    }
    fmt.Println(v1)
    v59 := []int{335, -298, -2147483648, 127}
    if false {
        if len(v59) > 0 {
            v59[0] = 32768
        }
        switch "hello" {
        case "":
            switch 839 {
            case 32767:
                v1 = v1
            case 4294967295:
                v1 = (^(v1))
            case 0:
                switch 128 {
                case 537:
                    v1 = v1
                case -9223372036854775808:
                    for fi60 := 0; fi60 < 6; fi60++ {
                        if false {
                            v61 := GenS7{F0: 127, F1: 29}
                            v1 -= v1
                            _ = v61
                        } else {
                            fmt.Println(fi60, v1)
                            if false {
                                v1 = v1
                            } else {
                                v1 = (v1 | (v1 ^ (v1 * v1)))
                                for fi62 := 0; fi62 < 1; fi62++ {
                                    var v63 int16 = 702
                                    v63 -= 122
                                    if true {
                                        var v64 uint8 = 255
                                        v1 -= -2147483648
                                        _ = v64
                                    } else {
                                        fmt.Println(fi60, fi60, fi60)
                                        if (false && (2 > 1)) {
                                            v65 := GenS1{F0: -128, F1: 18446744073709551615, F2: 141}
                                            if len(v59) > 0 {
                                                v59[0] = (fi62 << ((944) & 63))
                                            }
                                            v63 += -339
                                            _ = v65
                                        }
                                        var v66 float32 = ((-2.25 * float32((-2.25 - 1.5))) - 3.14159)
                                        _ = v66
                                    }
                                    _ = fi62
                                    _ = v63
                                }
                                for fi67 := 0; fi67 < 6; fi67++ {
                                    v1 -= v1
                                    v1 -= v1
                                    if false {
                                        if (false || false) {
                                            switch (fi60 / (((fi67 | (fi60 - 2147483647))) | 1)) {
                                            case 139:
                                                fmt.Println(v1, fi67)
                                            case 491:
                                                v2[990] = 2147483647
                                            }
                                            v2[fi67] = 1
                                            v68 := GenS8{F0: 0, F1: 186, F2: 9223372036854775807}
                                            _ = v68
                                        }
                                        var v69 int8 = 100
                                        v69 = (v69 & 1)
                                        _ = v69
                                    } else {
                                        v1 += -2147483648
                                    }
                                    _ = fi67
                                }
                            }
                            var v70 int64 = 42
                            _ = v70
                        }
                        var v71 float32 = (1e10 * (float32(0.1) + 0.1))
                        _ = fi60
                        _ = v71
                    }
                case 42:
                    var v72 uint16 = 42
                    _ = v72
                default:
                    switch ("foo bar" + ("foo bar" + "foo bar")) {
                    case "\n\t":
                        if (826 > 2) {
                            fmt.Println(v1)
                        } else {
                            if len(v59) > 0 {
                                v59[0] = 4294967295
                            }
                            if len(v59) >= 2 {
                                v59 = v59[1:len(v59)]
                            }
                            v1 = (v1 ^ 0)
                        }
                    default:
                        var v73 float32 = 1.5
                        _ = v73
                    }
                }
            }
        case "\n\t":
            if false {
                if len(v59) >= 2 {
                    v59 = v59[1:len(v59)]
                }
                var v74 uint64 = 42
                _ = v74
            } else {
                _ = len(v59)
                v59 = append(v59, 42)
            }
        case "hello":
            if false {
                v1 = (v1 ^ v1)
            }
        default:
            var v75 float32 = (((3.14159 + (0.0 / 0.1)) - float32((float64(((-2.25 - float32((1e10 * 1.5))) - 3.14159)) - 3.14159))) * 1.5)
            _ = v75
        }
        v1 -= 2147483647
    } else {
        fmt.Println(v1)
        v1 = 988
        if (!(2 != 42)) {
            if len(v59) >= 2 {
                v59 = v59[1:len(v59)]
            }
            v76 := GenS1{F0: 0, F1: 42, F2: 255}
            _ = v76
        } else {
            v77 := GenS9{F0: "", F1: -867, F2: 32768}
            _ = v77
        }
    }
    var v78 string = ("" + ("\n\t" + (("" + "a") + "\n\t")))
    switch ((("hello" + (v78 + "")) + ("\n\t" + "foo bar")) + "hello") {
    case "foo bar":
        v1 += 227
    default:
        _ = len(v59)
    }
    if true {
        for fi79 := 0; fi79 < 16; fi79++ {
            if (65535 <= 2) {
                if false {
                    v80 := GenS1{F0: 74, F1: 1, F2: 2}
                    _ = v80
                }
                v81 := GenS2{F0: -974}
                _ = v81
            }
            _ = fi79
        }
        v1 = 369
        v78 = v78
    }
    _ = len(v3)
    if true {
        for fi82 := 0; fi82 < 13; fi82++ {
            v3[689] = 65535
            _ = fi82
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v59:", len(v59), v59)
    fmt.Println("v78:", v78)
    _ = v1
    _ = v2
    _ = v3
    _ = v52
    _ = v59
    _ = v78
}
