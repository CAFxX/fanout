package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int8; F1 int32; F2 int8 }
type GenS3 struct { F0 int; F1 bool; F2 int64 }
type GenS4 struct { F0 uint16; F1 uint32; F2 int }
type GenS5 struct { F0 int16; F1 int64 }
func genFn0() int64 {
    switch 32767 {
    case -2147483648:
        var v7 uint = 4294967295
        _ = v7
    case 95:
        v8 := GenS1{F0: 253}
        _ = v8
    case 32767:
        var v9 int32 = 42
        _ = v9
    default:
        switch 42 {
        case 110:
            if false {
                if (3.14159 > 0.1) {
                    if (true || (!false)) {
                        var v10 string = "foo bar"
                        if false {
                            v10 = v10
                        }
                        _ = v10
                    } else {
                        var v11 int16 = -32768
                        var v12 uint16 = 32767
                        var v13 float32 = float32(0.0)
                        _ = v11
                        _ = v12
                        _ = v13
                    }
                    var v14 float32 = (float32(0.1) * 1e10)
                    switch "hello" {
                    case "foo bar":
                        if true {
                            v14 = ((v14 * v14) + 3.14159)
                            if true {
                                fmt.Println(v14)
                            } else {
                                v15 := []int{-386, 1, 0, 4294967295}
                                _ = v15
                            }
                            v16 := []int{32767}
                            _ = v16
                        } else {
                            var v17 uint = 1
                            v18 := []int{}
                            var v19 string = "a"
                            _ = v17
                            _ = v18
                            _ = v19
                        }
                    case "hello":
                        v20 := []int{}
                        _ = v20
                    case "":
                        v21 := GenS2{F0: 0, F1: -1, F2: -58}
                        _ = v21
                    }
                    _ = v14
                } else {
                    switch 32768 {
                    case 2147483647:
                        for fi22 := 0; fi22 < 5; fi22++ {
                            var v23 int16 = 32767
                            v23 = v23
                            _ = fi22
                            _ = v23
                        }
                    }
                }
                v24 := GenS3{F0: -2147483648, F1: (!((32767 <= 170) && (!true))), F2: -926}
                for fi25 := 0; fi25 < 16; fi25++ {
                    var v26 int16 = 0
                    switch 4294967295 {
                    case 456:
                        v27 := []int{(fi25 / ((0) | 1)), 258, fi25, -292}
                        _ = v27
                    case 0:
                        v28 := GenS2{F0: -1, F1: 0, F2: -1}
                        _ = v28
                    }
                    if ((("\n\t" + "hello") < "hello") && (255 != 255)) {
                        fmt.Println(v26, v26)
                        var v29 int64 = -1
                        _ = v29
                    } else {
                        for fi30 := 0; fi30 < 13; fi30++ {
                            fmt.Println(v26, fi30)
                            _ = fi30
                        }
                        var v31 float64 = (1.5 / ((0.1 + (1.5 + -2.25)) - float64(1.5)))
                        _ = v31
                    }
                    _ = fi25
                    _ = v26
                }
                _ = v24
            } else {
                var v32 int64 = 1
                _ = v32
            }
        case -348:
            for fi33 := 0; fi33 < 15; fi33++ {
                var v34 uint64 = 449
                v34 = (v34 * v34)
                for fi35 := 0; fi35 < 2; fi35++ {
                    v36 := []int{fi33, (fi35 ^ fi33)}
                    var v37 int = (fi33 << ((1) & 63))
                    if (0.0 > 0.1) {
                        if (127 < fi33) {
                            fmt.Println(fi35, v34, fi33)
                            var v38 uint = 750
                            switch "foo bar" {
                            case "":
                                switch "foo bar" {
                                case "foo bar":
                                    v38 = (v38 ^ (^(v38)))
                                case "\n\t":
                                    if len(v36) > 0 {
                                        v36[0] = v37
                                    }
                                default:
                                    var v39 float32 = 1.5
                                    _ = v39
                                }
                            case "foo bar":
                                var v40 int64 = int64(v37)
                                _ = v40
                            case "\n\t":
                                var v41 uint8 = 127
                                _ = v41
                            default:
                                v42 := GenS2{F0: 0, F1: 0, F2: -128}
                                _ = v42
                            }
                            _ = v38
                        } else {
                            var v43 int8 = 127
                            _ = v43
                        }
                        v44 := GenS2{F0: 6, F1: 555, F2: 62}
                        _ = v44
                    } else {
                        var v45 int8 = -1
                        switch fi33 {
                        case -1:
                            v36 = append(v36, -1)
                        default:
                            var v46 string = "\n\t"
                            _ = v46
                        }
                        _ = v45
                    }
                    _ = fi35
                    _ = v36
                    _ = v37
                }
                _ = fi33
                _ = v34
            }
        case -9223372036854775808:
            v47 := GenS3{F0: -9223372036854775808, F1: true, F2: 349}
            _ = v47
        default:
            if (false && (42 != -1)) {
                var v48 int = 128
                _ = v48
            } else {
                v49 := []int{128, 42}
                var v50 float32 = ((1e10 + 1.5) + (1.5 / 1.5))
                v51 := GenS1{F0: -32768}
                _ = v49
                _ = v50
                _ = v51
            }
        }
    }
    var v52 float32 = (((3.14159 + (-2.25 - (((3.14159 * 1.5) + -2.25) + 1.5))) - 3.14159) * 0.1)
    for fi53 := 0; fi53 < 7; fi53++ {
        v54 := GenS4{F0: 42, F1: 1, F2: -9223372036854775808}
        _ = fi53
        _ = v54
    }
    var v55 int16 = 409
    _ = v52
    _ = v55
    return -9223372036854775808
}
func genFn1(p0 int64) int16 {
    p0 -= (p0 % ((p0) | 1))
    p0 = genFn0()
    if false {
        fmt.Println(p0)
    } else {
        p0 += (-(p0))
        var v56 uint64 = 554
        _ = v56
    }
    _ = p0
    return 32767
}

func main() {
    var v1 int64 = 1
    var v2 int64 = -424
    v3 := make(map[string]int)
    v4 := make(map[int]uint32)
    v5 := make(map[int]int)
    v6 := []string{(("hello" + "foo bar") + "foo bar"), ((("hello" + "") + "") + "\n\t"), ("" + "hello")}
    v1 -= (v2 >> ((2) & 63))
    fmt.Println(v2, v1)
    v2 = (-(v2))
    for fi57 := 0; fi57 < 8; fi57++ {
        v1 += -1
        _ = fi57
    }
    v58 := GenS5{F0: genFn1(956), F1: 0}
    var v59 bool = (42 <= 798)
    v59 = false
    fmt.Println(v2)
    var v60 float64 = ((-2.25 + 1.5) * float64((3.14159 * 1.5)))
    for fi61 := 0; fi61 < 12; fi61++ {
        var v62 float32 = (-2.25 + 3.14159)
        fmt.Println(fi61, v59, v59)
        _ = fi61
        _ = v62
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v59:", v59)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v58
    _ = v59
    _ = v6
    _ = v60
}
