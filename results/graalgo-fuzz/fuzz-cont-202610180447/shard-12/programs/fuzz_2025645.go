package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int; F2 string }
type GenS2 struct { F0 bool; F1 int64; F2 int64 }
type GenS3 struct { F0 int; F1 uint64; F2 string }
type GenS4 struct { F0 uint32; F1 uint; F2 uint64 }
type GenS5 struct { F0 int16; F1 int8; F2 uint }
type GenS6 struct { F0 int }
type GenS7 struct { F0 uint8; F1 int32 }
type GenS8 struct { F0 uint32; F1 string }
type GenS9 struct { F0 uint16; F1 int16; F2 int16 }
func genFn0() int16 {
    if true {
        var v4 string = "\n\t"
        var v5 int16 = 42
        _ = v4
        _ = v5
    }
    for fi6 := 0; fi6 < 1; fi6++ {
        var v7 uint8 = 253
        fmt.Println(v7, fi6)
        _ = fi6
        _ = v7
    }
    if false {
        var v8 float32 = 0.0
        v8 += 3.14159
        switch 127 {
        case 682:
            if (true && true) {
                v8 -= v8
                v9 := []int{-9223372036854775808, 170, 4294967295}
                v10 := GenS1{F0: -9223372036854775808, F1: 32768, F2: ""}
                _ = v10
                _ = v9
            }
        case 0:
            switch "foo bar" {
            case "foo bar":
                v8 = (0.1 / v8)
            }
        case 42:
            if false {
                v8 = v8
                var v11 uint64 = 42
                _ = v11
            } else {
                for fi12 := 0; fi12 < 13; fi12++ {
                    v8 -= v8
                    v8 = v8
                    _ = fi12
                }
                for fi13 := 0; fi13 < 8; fi13++ {
                    var v14 int64 = 798
                    if true {
                        v15 := GenS2{F0: true, F1: v14, F2: v14}
                        switch fi13 {
                        case 127:
                            var v16 uint32 = 453
                            _ = v16
                        case 909:
                            if (v14 == 0) {
                                for fi17 := 0; fi17 < 16; fi17++ {
                                    v14 += v14
                                    _ = fi17
                                }
                                v14 = (^(v14))
                                v18 := GenS2{F0: (0 < 701), F1: v14, F2: v14}
                                _ = v18
                            } else {
                                v14 = 0
                                v19 := GenS2{F0: true, F1: v14, F2: v14}
                                _ = v19
                            }
                        }
                        for fi20 := 0; fi20 < 16; fi20++ {
                            v14 -= v14
                            _ = fi20
                        }
                        _ = v15
                    } else {
                        fmt.Println(fi13, fi13)
                        v8 = 0.1
                    }
                    v14 -= (v14 | v14)
                    _ = fi13
                    _ = v14
                }
                var v21 uint8 = 255
                _ = v21
            }
        }
        _ = v8
    }
    var v22 uint16 = 0
    v23 := GenS2{F0: (true || (!true)), F1: -561, F2: -1}
    _ = v22
    _ = v23
    return -794
}

func main() {
    v1 := []uint32{2147483647, 383}
    v2 := []uint32{}
    v3 := make(map[uint32]string)
    v24 := GenS3{F0: 128, F1: 2, F2: ("\n\t" + ("a" + "hello"))}
    switch 9223372036854775807 {
    case 32768:
        v25 := GenS1{F0: -9223372036854775808, F1: 127, F2: "foo bar"}
        _ = v25
    case 127:
        switch 9223372036854775807 {
        case 32767:
            var v26 int = 590
            _ = v26
        case 9223372036854775807:
            var v27 uint = 42
            _ = v27
        default:
            var v28 uint16 = 0
            _ = v28
        }
    }
    var v29 uint16 = 1
    _ = len(v3)
    var v30 uint32 = 27
    v2 = append(v2, 0)
    var v31 int32 = -123
    v31 = (v31 >> ((18446744073709551615) & 63))
    for fi32 := 0; fi32 < 10; fi32++ {
        if (65535 >= 18446744073709551615) {
            fmt.Println(v31)
            var v33 float64 = (0.1 - (1e10 + 0.1))
            switch "a" {
            case "foo bar":
                switch "foo bar" {
                case "\n\t":
                    v34 := GenS1{F0: 0, F1: (fi32 % (((-(fi32))) | 1)), F2: "foo bar"}
                    _ = v34
                case "":
                    v31 = v31
                default:
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                }
            case "a":
                for fi35 := 0; fi35 < 10; fi35++ {
                    v30 = (^(v30))
                    _ = len(v1)
                    var v36 int64 = 0
                    v36 -= v36
                    _ = fi35
                    _ = v36
                }
            default:
                fmt.Println(fi32, v30)
            }
            _ = v33
        } else {
            v29 += 561
        }
        var v37 uint32 = 816
        v38 := GenS4{F0: (v37 - (v37 | (v30 << ((4294967295) & 63)))), F1: 4294967295, F2: 664}
        _ = fi32
        _ = v37
        _ = v38
    }
    v39 := GenS3{F0: -1, F1: 958, F2: ""}
    var v40 uint32 = 942
    var v41 bool = ((v31 % ((507) | 1)) > v31)
    v40 += (^(v40))
    var v42 float32 = (0.1 / (((1e10 + ((-2.25 - float32((((1e10 * 0.0) + 0.0) * (1e10 + 0.0)))) + 1e10)) + (((1e10 + 0.0) + ((0.1 / (1.5 / ((float32((((float64(0.0) - 3.14159) + 3.14159) + 1.5)) / 1.0) + float32((1.5 - 3.14159))))) * -2.25)) + 0.1)) / 1.0))
    if true {
        fmt.Println(v30, v29)
        for fi43 := 0; fi43 < 6; fi43++ {
            for fi44 := 0; fi44 < 4; fi44++ {
                v31 += (v31 >> ((uint(v29)) & 63))
                v45 := GenS5{F0: -1, F1: -22, F2: 674}
                v46 := GenS6{F0: -1}
                _ = fi44
                _ = v45
                _ = v46
            }
            if (!true) {
                for fi47 := 0; fi47 < 19; fi47++ {
                    v48 := GenS7{F0: 2, F1: v31}
                    _ = fi47
                    _ = v48
                }
                if ("\n\t" != "\n\t") {
                    v49 := GenS7{F0: 38, F1: v31}
                    var v50 int = fi43
                    v51 := GenS2{F0: v41, F1: 1, F2: -9223372036854775808}
                    _ = v49
                    _ = v50
                    _ = v51
                } else {
                    switch "\n\t" {
                    case "hello":
                        v52 := GenS7{F0: 2, F1: (^(v31))}
                        _ = v52
                    case "":
                        var v53 uint = 912
                        _ = v53
                    case "a":
                        var v54 string = "\n\t"
                        _ = v54
                    }
                    fmt.Println(v31)
                    v41 = (-231 >= 0)
                }
                var v55 uint = 18446744073709551615
                _ = v55
            } else {
                var v56 uint8 = 42
                var v57 uint8 = (v56 % ((v56) | 1))
                if (979 >= 760) {
                    _ = len(v3)
                    switch (((("a" + "a") + "\n\t") + "foo bar") + ("hello" + "foo bar")) {
                    case "hello":
                        switch "foo bar" {
                        case "":
                            if (genFn0() == -174) {
                                v29 -= 42
                            } else {
                                v57 -= 179
                                var v58 int8 = -1
                                if v41 {
                                    var v59 int64 = 42
                                    _ = v59
                                } else {
                                    v60 := GenS8{F0: v30, F1: (("foo bar" + ("hello" + ("hello" + "a"))) + (("hello" + "") + ("hello" + ("" + "\n\t"))))}
                                    v61 := GenS9{F0: 130, F1: 0, F2: -1}
                                    v40 = 42
                                    _ = v60
                                    _ = v61
                                }
                                _ = v58
                            }
                        case "hello":
                            if v41 {
                                v30 = v40
                                switch (fi43 | (fi43 << ((448) & 63))) {
                                case 1:
                                    v29 -= v29
                                case 811:
                                    v57 = v56
                                default:
                                    var v62 bool = ((v31 & v31) >= 1)
                                    _ = v62
                                }
                                _ = len(v2)
                            }
                        case "\n\t":
                            var v63 int32 = v31
                            _ = v63
                        default:
                            v56 = v57
                        }
                    default:
                        var v64 uint8 = v56
                        _ = v64
                    }
                    var v65 int = 735
                    _ = v65
                }
                _ = v56
                _ = v57
            }
            var v66 int32 = (v31 | v31)
            _ = fi43
            _ = v66
        }
    } else {
        var v67 float64 = 3.14159
        var v68 string = (((("a" + "") + "") + ("" + "")) + "")
        var v69 uint16 = (v29 % ((v29) | 1))
        _ = v67
        _ = v68
        _ = v69
    }
    delete(v3, 4294967295)
    v42 = ((v42 / float32(1.5)) + (v42 + v42))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v2
    _ = v24
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v39
    _ = v40
    _ = v41
    _ = v42
}
