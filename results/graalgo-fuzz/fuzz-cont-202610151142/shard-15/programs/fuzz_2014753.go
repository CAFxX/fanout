package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64; F2 uint16 }
type GenS2 struct { F0 uint16; F1 uint }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 bool }

func main() {
    v1 := make(map[uint32]string)
    var v2 string = ((("a" + "hello") + "") + "hello")
    v3 := make(map[uint32]string)
    v4 := make(map[int]int)
    v5 := make(map[int]uint32)
    v6 := make(map[string]uint32)
    v2 = "foo bar"
    v7 := []int{}
    delete(v5, 1)
    v2 = v2
    var v8 float32 = ((1e10 * -2.25) / (0.0 - 1e10))
    _ = len(v7)
    if true {
        for fi9 := 0; fi9 < 11; fi9++ {
            if true {
                for fi10 := 0; fi10 < 20; fi10++ {
                    switch "foo bar" {
                    case "":
                        v8 = v8
                    case "hello":
                        switch "" {
                        case "a":
                            if len(v7) > 0 {
                                v7[0] = (fi10 % (((fi9 + fi9)) | 1))
                            }
                        default:
                            switch fi10 {
                            case 128:
                                v11 := GenS1{F0: 1, F1: 483, F2: 1}
                                _ = v11
                            case 352:
                                if len(v7) >= 2 {
                                    v7 = v7[1:len(v7)]
                                }
                            default:
                                switch (fi10 / (((fi10 % ((fi10) | 1))) | 1)) {
                                case -830:
                                    v8 = v8
                                case -9223372036854775808:
                                    if len(v7) > 0 {
                                        v7[0] = 32768
                                    }
                                case 9223372036854775807:
                                    v2 = (v2 + "foo bar")
                                default:
                                    var v12 int64 = -1
                                    _ = v12
                                }
                            }
                            v5[fi10] = 783
                            v2 = v2
                        }
                    default:
                        v13 := GenS1{F0: 4294967295, F1: 0, F2: 255}
                        _ = v13
                    }
                    v8 = ((float32((((3.14159 / (0.0 - (0.1 / ((0.0 + 1e10) / 1.5)))) * (3.14159 + float64((0.1 - float32(3.14159))))) / (-2.25 / 1e10))) - 1.5) * (float32(1e10) + -2.25))
                    v2 = v2
                    if false {
                        var v14 uint8 = 39
                        if true {
                            var v15 float32 = v8
                            var v16 uint8 = 255
                            _ = v15
                            _ = v16
                        } else {
                            v7 = append(v7, 744)
                        }
                        fmt.Println(v2)
                        fmt.Println(fi10, v14)
                        v14 -= 42
                        _ = v14
                    }
                    v17 := GenS2{F0: 305, F1: 637}
                    _ = fi10
                    _ = v17
                }
                fmt.Println(v8, v2)
            } else {
                v18 := GenS2{F0: 65535, F1: 2}
                v19 := GenS1{F0: 65535, F1: 18446744073709551615, F2: 2}
                v8 = 1e10
                _ = v18
                _ = v19
            }
            if (((!(18446744073709551615 <= 0)) && ((v8 / (float32(-2.25) / 3.14159)) != v8)) || (519 <= 898)) {
                v8 += (v8 * float32(((0.0 * 0.1) * 0.0)))
                var v20 uint64 = 2
                v8 = float32((float64(0.1) - ((float64(0.1) + (0.1 + (-2.25 * 3.14159))) - -2.25)))
                _ = v20
            }
            v21 := GenS1{F0: 18446744073709551615, F1: 42, F2: 979}
            _ = fi9
            _ = v21
        }
        if true {
            var v22 float32 = v8
            _ = v22
        } else {
            v8 = v8
            switch 42 {
            case 1:
                for fi23 := 0; fi23 < 1; fi23++ {
                    v8 -= ((float32(0.1) + v8) / v8)
                    _ = fi23
                }
            case 32768:
                for fi24 := 0; fi24 < 12; fi24++ {
                    if false {
                        if false {
                            v25 := GenS1{F0: 4294967295, F1: 42, F2: 184}
                            var v26 int32 = 1
                            _ = v25
                            _ = v26
                        } else {
                            for fi27 := 0; fi27 < 9; fi27++ {
                                v8 += v8
                                _ = fi27
                            }
                            v8 -= v8
                        }
                        v2 = ""
                        var v28 uint8 = 42
                        _ = v28
                    } else {
                        v29 := GenS3{F0: 42}
                        fmt.Println(fi24, v2, v8)
                        _ = v29
                    }
                    for fi30 := 0; fi30 < 16; fi30++ {
                        fmt.Println(fi30, fi30)
                        v31 := GenS2{F0: 2, F1: 712}
                        v2 = v2
                        _ = fi30
                        _ = v31
                    }
                    for fi32 := 0; fi32 < 9; fi32++ {
                        fmt.Println(fi32, fi24)
                        var v33 bool = false
                        if (1 > 915) {
                            fmt.Println(v33, v33)
                        }
                        fmt.Println(fi32, fi24, v2)
                        _ = fi32
                        _ = v33
                    }
                    _ = fi24
                }
            case -848:
                for fi34 := 0; fi34 < 17; fi34++ {
                    var v35 int = (fi34 | (fi34 + fi34))
                    for fi36 := 0; fi36 < 19; fi36++ {
                        var v37 int16 = 586
                        _ = fi36
                        _ = v37
                    }
                    _ = fi34
                    _ = v35
                }
            }
            v8 -= v8
        }
        _ = len(v7)
    }
    for fi38 := 0; fi38 < 7; fi38++ {
        v2 = ""
        var v39 float64 = 0.0
        var v40 int64 = -102
        _ = fi38
        _ = v39
        _ = v40
    }
    switch 9223372036854775807 {
    case -644:
        v8 = 1e10
    default:
        if true {
            for fi41 := 0; fi41 < 12; fi41++ {
                fmt.Println(fi41, v8)
                v42 := GenS4{F0: 980}
                _ = fi41
                _ = v42
            }
        } else {
            v8 += v8
            if true {
                if false {
                    v2 = ("\n\t" + "a")
                }
                for fi43 := 0; fi43 < 5; fi43++ {
                    for fi44 := 0; fi44 < 2; fi44++ {
                        var v45 int16 = -245
                        v46 := GenS2{F0: 255, F1: 255}
                        _ = fi44
                        _ = v45
                        _ = v46
                    }
                    v8 += (3.14159 - v8)
                    if (2 == 136) {
                        v7 = append(v7, (fi43 % ((fi43) | 1)))
                        switch -881 {
                        case -2147483648:
                            switch v2 {
                            case "":
                                fmt.Println(v2, v8, fi43)
                            case "foo bar":
                                var v47 int8 = -71
                                _ = v47
                            case "a":
                                fmt.Println(v2, v8, v2)
                            }
                        default:
                            var v48 int32 = 1
                            _ = v48
                        }
                    } else {
                        v3[27] = (v2 + v2)
                    }
                    _ = fi43
                }
                if (-1 < -2147483648) {
                    v2 = v2
                    var v49 uint32 = 65535
                    _ = v49
                } else {
                    _ = len(v1)
                    var v50 int32 = -1
                    _ = v50
                }
            } else {
                v51 := GenS5{F0: ((v2 <= ("a" + "foo bar")) && ((0 > 2) || true))}
                _ = v51
            }
        }
    }
    var v52 uint32 = 42
    var v53 uint = 606
    v53 = v53
    var v54 float32 = (v8 * v8)
    fmt.Println(v54)
    if false {
        _ = len(v7)
        v54 -= v54
    }
    var v55 uint64 = 524
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    fmt.Println("v54:", v54)
    fmt.Println("v55:", v55)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v52
    _ = v53
    _ = v54
    _ = v55
    _ = v6
    _ = v7
    _ = v8
}
