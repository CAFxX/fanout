package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string; F2 uint16 }
type GenS2 struct { F0 int16; F1 uint8; F2 uint8 }
type GenS3 struct { F0 int32; F1 uint8; F2 bool }
type GenS4 struct { F0 uint64; F1 uint64; F2 int }
type GenS5 struct { F0 uint8; F1 uint8 }
type GenS6 struct { F0 bool; F1 int64; F2 uint }
type GenS7 struct { F0 uint16; F1 uint64 }
type GenS8 struct { F0 int8; F1 bool }
type GenS9 struct { F0 uint64; F1 int16 }
func genFn0(p0 string, p1 uint8, p2 uint32) uint32 {
    var v7 int = -2147483648
    v8 := []int{343, (v7 | 32768), (v7 & (v7 << ((255) & 63)))}
    p1 = p1
    var v9 float64 = float64(((0.1 + (-2.25 + 0.0)) + 1.5))
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return p2
}

func main() {
    var v1 uint64 = 188
    v2 := make(map[int]string)
    v3 := make(map[string]string)
    v4 := []string{("" + "a")}
    v5 := make(map[uint32]bool)
    v6 := make(map[string]int)
    if (255 != 255) {
        var v10 string = ""
        if (447 != 9223372036854775807) {
            switch 1 {
            case 128:
                v1 += v1
            case -9223372036854775808:
                v10 = (("hello" + ("a" + ("hello" + v10))) + "a")
            default:
                v1 = v1
            }
            var v11 string = "\n\t"
            _ = v11
        } else {
            if false {
                v10 = (v10 + (v10 + v10))
            } else {
                for fi12 := 0; fi12 < 15; fi12++ {
                    v13 := GenS1{F0: 0, F1: ("a" + "hello"), F2: 865}
                    for fi14 := 0; fi14 < 8; fi14++ {
                        if (0 > 2147483647) {
                            var v15 int64 = 9223372036854775807
                            v1 = (v1 % ((18446744073709551615) | 1))
                            _ = v15
                        }
                        fmt.Println(fi12, fi14)
                        if false {
                            _ = len(v2)
                            var v16 string = v10
                            if true {
                                if false {
                                    switch v16 {
                                    case "hello":
                                        var v17 int8 = -48
                                        _ = v17
                                    case "\n\t":
                                        var v18 float64 = 1.5
                                        _ = v18
                                    case "foo bar":
                                        v10 = v16
                                        if len(v4) > 0 {
                                            v4[0] = (v10 + v16)
                                        }
                                        v19 := GenS2{F0: -908, F1: 128, F2: 0}
                                        _ = v19
                                    default:
                                        var v20 uint32 = 311
                                        _ = v20
                                    }
                                }
                                fmt.Println(fi14, fi14, v10)
                            } else {
                                v21 := GenS3{F0: 1, F1: 0, F2: true}
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                _ = v21
                            }
                            _ = v16
                        } else {
                            if true {
                                v22 := GenS3{F0: 42, F1: 255, F2: (!true)}
                                _ = v22
                            }
                        }
                        _ = fi14
                    }
                    _ = fi12
                    _ = v13
                }
                for fi23 := 0; fi23 < 18; fi23++ {
                    var v24 uint16 = 959
                    v24 += v24
                    v24 -= (v24 | v24)
                    _ = fi23
                    _ = v24
                }
                switch 32768 {
                case 627:
                    if true {
                        if (!(!false)) {
                            v10 = ("" + v10)
                        } else {
                            fmt.Println(v10)
                            if (!true) {
                                v10 = v10
                                var v25 string = ("hello" + v10)
                                _ = v25
                            } else {
                                v26 := GenS2{F0: 510, F1: 2, F2: 42}
                                _ = v26
                            }
                            var v27 float32 = 3.14159
                            _ = v27
                        }
                        switch v10 {
                        case "\n\t":
                            v28 := GenS4{F0: 361, F1: (v1 * (v1 >> ((699) & 63))), F2: 920}
                            _ = v28
                        case "hello":
                            v1 += v1
                        case "":
                            v29 := v5[513]
                            _ = v29
                        }
                    } else {
                        for fi30 := 0; fi30 < 4; fi30++ {
                            var v31 uint16 = 2
                            _ = fi30
                            _ = v31
                        }
                    }
                case -531:
                    switch v10 {
                    case "\n\t":
                        for fi32 := 0; fi32 < 2; fi32++ {
                            for fi33 := 0; fi33 < 12; fi33++ {
                                delete(v2, fi32)
                                v4 = append(v4, v10)
                                v1 += (v1 % (((^(v1))) | 1))
                                _ = fi33
                            }
                            fmt.Println(v1)
                            _ = fi32
                        }
                    case "a":
                        var v34 float32 = 0.1
                        _ = v34
                    case "":
                        _ = len(v4)
                    default:
                        v35 := GenS5{F0: 155, F1: 0}
                        _ = v35
                    }
                case 42:
                    var v36 int8 = 42
                    _ = v36
                default:
                    v37 := GenS6{F0: (false && false), F1: 42, F2: 846}
                    _ = v37
                }
            }
        }
        _ = v10
    }
    fmt.Println(v1)
    v38 := GenS7{F0: 1, F1: v1}
    var v39 int = -2147483648
    if false {
        var v40 int32 = 265
        if len(v4) > 0 {
            v4[0] = "foo bar"
        }
        v40 += v40
        _ = v40
    } else {
        var v41 int64 = 9223372036854775807
        _ = v41
    }
    v1 = v1
    var v42 int32 = 0
    var v43 int64 = -9223372036854775808
    v39 -= (v39 << ((972) & 63))
    for fi44 := 0; fi44 < 11; fi44++ {
        if true {
            v39 = (fi44 / ((fi44) | 1))
        } else {
            for fi45 := 0; fi45 < 8; fi45++ {
                fmt.Println(fi44, v1, v42)
                v46 := GenS8{F0: -118, F1: false}
                var v47 float32 = (1.5 + -2.25)
                v39 -= fi45
                _ = fi45
                _ = v46
                _ = v47
            }
        }
        switch (v39 >> ((0) & 63)) {
        case -2147483648:
            var v48 float32 = (0.0 - (1e10 - 3.14159))
            _ = v48
        default:
            var v49 uint32 = 42
            _ = v49
        }
        _ = fi44
    }
    v39 -= (v39 & 127)
    for fi50 := 0; fi50 < 17; fi50++ {
        if false {
            _ = len(v4)
            fmt.Println(fi50, v1)
        }
        if false {
            v42 += v42
        } else {
            switch ("\n\t" + "\n\t") {
            case "\n\t":
                var v51 float32 = 0.0
                _ = v51
            default:
                delete(v2, fi50)
            }
            v52 := GenS9{F0: (v1 - (^(v1))), F1: 0}
            v1 -= (v1 / ((182) | 1))
            _ = v52
        }
        _ = fi50
    }
    v53 := GenS4{F0: v1, F1: v1, F2: v39}
    v54 := GenS6{F0: true, F1: (v43 | v43), F2: 0}
    var v55 int32 = 752
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v55:", v55)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v42
    _ = v43
    _ = v5
    _ = v53
    _ = v54
    _ = v55
    _ = v6
}
