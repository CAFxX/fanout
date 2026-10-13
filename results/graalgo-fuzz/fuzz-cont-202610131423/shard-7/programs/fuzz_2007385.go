package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8 }
type GenS2 struct { F0 string; F1 uint }
type GenS3 struct { F0 uint8; F1 uint32; F2 uint64 }
type GenS4 struct { F0 uint64; F1 uint64 }
type GenS5 struct { F0 uint32; F1 int16; F2 int64 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 uint32; F1 uint64 }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 int32; F1 uint32 }

func main() {
    var v1 uint32 = 389
    var v2 uint32 = (v1 | 4294967295)
    var v3 float64 = 0.1
    var v4 int32 = -272
    var v5 uint32 = 2
    for fi6 := 0; fi6 < 20; fi6++ {
        v7 := []int{fi6}
        _ = fi6
        _ = v7
    }
    switch 2147483647 {
    case 108:
        v5 = 355
    default:
        if false {
            v1 -= 0
            v8 := GenS1{F0: v1, F1: 2}
            v9 := GenS2{F0: "", F1: 4294967295}
            _ = v8
            _ = v9
        }
    }
    switch 127 {
    case 32767:
        v5 = v2
    case -2147483648:
        var v10 uint32 = v1
        _ = v10
    }
    v5 -= (v5 << ((1) & 63))
    v11 := GenS3{F0: 128, F1: (v1 - (v5 % ((4294967295) | 1))), F2: 2}
    v12 := GenS3{F0: 184, F1: 981, F2: 42}
    for fi13 := 0; fi13 < 6; fi13++ {
        v5 = v5
        var v14 int64 = -9223372036854775808
        _ = fi13
        _ = v14
    }
    if true {
        fmt.Println(v5)
        if true {
            for fi15 := 0; fi15 < 18; fi15++ {
                for fi16 := 0; fi16 < 1; fi16++ {
                    v17 := GenS4{F0: 477, F1: 2}
                    _ = fi16
                    _ = v17
                }
                for fi18 := 0; fi18 < 6; fi18++ {
                    v19 := []int{fi15, (fi15 / ((fi15) | 1)), fi18}
                    if (!(!(!true))) {
                        v2 -= (v2 % (((v1 & 0)) | 1))
                        _ = len(v19)
                        v5 -= v2
                    } else {
                        v5 -= (v2 | v5)
                    }
                    _ = fi18
                    _ = v19
                }
                if (0 <= 592) {
                    v1 += (^(v2))
                    if true {
                        for fi20 := 0; fi20 < 4; fi20++ {
                            v3 -= (v3 + (v3 - (v3 / v3)))
                            v1 = 2
                            _ = fi20
                        }
                    } else {
                        v21 := []int{2147483647, (fi15 * -1)}
                        _ = v21
                    }
                } else {
                    v22 := []int{fi15}
                    _ = v22
                }
                _ = fi15
            }
        } else {
            switch "hello" {
            case "foo bar":
                v23 := GenS5{F0: (v2 ^ v5), F1: 1, F2: 1}
                _ = v23
            default:
                v3 += (0.0 / (((-2.25 - (v3 * (v3 / (-2.25 * (v3 * (0.0 + v3)))))) / 1.5) + v3))
            }
            v24 := GenS2{F0: ("hello" + "foo bar"), F1: 980}
            var v25 int64 = -9223372036854775808
            _ = v24
            _ = v25
        }
        v26 := GenS3{F0: 128, F1: (v5 >> ((365) & 63)), F2: 977}
        _ = v26
    }
    var v27 int16 = 682
    var v28 uint = 495
    if (((0.0 / (v3 + ((v3 - v3) * (v3 + ((v3 - v3) / 1.0))))) * 1.5) == v3) {
        var v29 bool = false
        if ((v2 > 597) || v29) {
            fmt.Println(v1)
            v28 = 42
            switch 0 {
            case 1:
                v30 := []int{-343}
                _ = v30
            default:
                for fi31 := 0; fi31 < 1; fi31++ {
                    var v32 int64 = 0
                    v33 := []int{}
                    _ = fi31
                    _ = v32
                    _ = v33
                }
            }
        }
        _ = v29
    } else {
        var v34 int64 = -125
        v35 := GenS6{F0: 1}
        var v36 float64 = float64(1.5)
        _ = v34
        _ = v35
        _ = v36
    }
    var v37 uint32 = 266
    v37 = (^(v37))
    if (false || false) {
        if true {
            var v38 uint = v28
            for fi39 := 0; fi39 < 8; fi39++ {
                switch (-(fi39)) {
                case 42:
                    v40 := GenS7{F0: v2, F1: 239}
                    _ = v40
                case 4294967295:
                    fmt.Println(fi39, v5, v4)
                case 32767:
                    if true {
                        v41 := []int{fi39}
                        _ = v41
                    }
                default:
                    var v42 int32 = (^(v4))
                    _ = v42
                }
                if true {
                    v37 = (v37 % ((v37) | 1))
                    if false {
                        v38 = 916
                    }
                    switch "\n\t" {
                    case "\n\t":
                        var v43 int16 = 1
                        _ = v43
                    case "a":
                        for fi44 := 0; fi44 < 8; fi44++ {
                            v45 := []int{fi39, 1, -9223372036854775808, fi39}
                            v46 := GenS8{F0: 0}
                            v37 = 0
                            _ = fi44
                            _ = v45
                            _ = v46
                        }
                    default:
                        if true {
                            var v47 int8 = -128
                            _ = v47
                        }
                    }
                }
                _ = fi39
            }
            _ = v38
        } else {
            var v48 float64 = float64(0.1)
            _ = v48
        }
    } else {
        v27 = v27
        for fi49 := 0; fi49 < 7; fi49++ {
            if false {
                if false {
                    if true {
                        if (v3 != ((1.5 / (((((-2.25 - 3.14159) / ((0.1 - v3) * (float64(0.0) * (v3 * v3)))) / 1e10) / v3) - v3)) / ((v3 - (v3 / (v3 - -2.25))) - v3))) {
                            v5 -= 2147483648
                        } else {
                            var v50 int32 = -1
                            _ = v50
                        }
                        v51 := []int{(fi49 + fi49)}
                        _ = v51
                    } else {
                        v2 -= (v37 << ((v28) & 63))
                        if false {
                            var v52 float64 = (v3 * (v3 + v3))
                            var v53 uint16 = 638
                            for fi54 := 0; fi54 < 18; fi54++ {
                                v55 := GenS7{F0: 2147483647, F1: 1}
                                _ = fi54
                                _ = v55
                            }
                            _ = v52
                            _ = v53
                        }
                        var v56 string = "a"
                        _ = v56
                    }
                    v57 := GenS3{F0: 255, F1: v1, F2: 4294967295}
                    for fi58 := 0; fi58 < 9; fi58++ {
                        v5 = (v5 - v5)
                        _ = fi58
                    }
                    _ = v57
                }
                v59 := GenS9{F0: v4, F1: 65535}
                _ = v59
            }
            _ = fi49
        }
        v60 := GenS3{F0: 128, F1: (v5 - 2), F2: 42}
        _ = v60
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v37
    _ = v4
    _ = v5
}
