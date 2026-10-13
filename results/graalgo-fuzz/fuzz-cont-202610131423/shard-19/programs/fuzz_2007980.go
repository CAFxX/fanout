package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint; F2 int32 }
type GenS2 struct { F0 uint8; F1 int; F2 int16 }
type GenS3 struct { F0 uint16; F1 uint32; F2 string }
type GenS4 struct { F0 int32; F1 int }
type GenS5 struct { F0 uint16; F1 int; F2 uint16 }
type GenS6 struct { F0 string }
func genFn0(p0 int16, p1 uint8) int8 {
    v5 := GenS1{F0: p1, F1: 65535, F2: 1}
    var v6 int16 = (p0 ^ p0)
    v7 := []int{1, -815}
    var v8 int16 = v6
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return 42
}
func genFn1() bool {
    for fi9 := 0; fi9 < 14; fi9++ {
        v10 := []int{fi9, (fi9 & fi9), (fi9 & (fi9 % ((fi9) | 1)))}
        var v11 uint8 = 122
        if (70 < 0) {
            var v12 uint16 = 65535
            if (0 <= -12) {
                for fi13 := 0; fi13 < 20; fi13++ {
                    if true {
                        v10 = append(v10, (fi13 << ((65535) & 63)))
                        v12 = (v12 | uint16(v11))
                        fmt.Println(v12, fi9, v12)
                    } else {
                        switch ("foo bar" + "") {
                        case "\n\t":
                            var v14 int32 = 2147483647
                            _ = v14
                        case "":
                            if len(v10) >= 2 {
                                v10 = v10[1:len(v10)]
                            }
                        }
                        v12 = 993
                    }
                    _ = fi13
                }
            } else {
                var v15 float64 = 1.5
                v15 -= (v15 * v15)
                _ = v15
            }
            var v16 string = (("foo bar" + "\n\t") + "hello")
            _ = v12
            _ = v16
        } else {
            for fi17 := 0; fi17 < 9; fi17++ {
                var v18 uint8 = v11
                if (42 != 100) {
                    var v19 int8 = -73
                    v20 := GenS1{F0: 255, F1: 433, F2: 42}
                    _ = v19
                    _ = v20
                }
                v21 := GenS2{F0: (^(v11)), F1: (fi9 - -702), F2: 970}
                fmt.Println(fi17, fi17, v11)
                v10 = append(v10, fi9)
                if len(v10) > 0 {
                    v10[0] = fi17
                }
                _ = fi17
                _ = v18
                _ = v21
            }
        }
        _ = fi9
        _ = v10
        _ = v11
    }
    var v22 uint = 630
    v22 = 42
    var v23 int16 = 42
    _ = v22
    _ = v23
    return (314 == 802)
}

func main() {
    var v1 uint = 822
    var v2 float32 = float32((0.1 / -2.25))
    var v3 uint64 = 957
    var v4 uint8 = 2
    v24 := GenS1{F0: v4, F1: (v1 * (v1 - v1)), F2: 1}
    v2 += 3.14159
    for fi25 := 0; fi25 < 12; fi25++ {
        if false {
            fmt.Println(v4, v4)
            for fi26 := 0; fi26 < 19; fi26++ {
                if (v4 <= (v4 * v4)) {
                    if (128 > 2) {
                        if false {
                            v27 := []int{fi26, fi26, fi25}
                            _ = v27
                        } else {
                            if (884 >= 0) {
                                v1 = v1
                            } else {
                                v28 := []int{}
                                v4 = (^(v4))
                                fmt.Println(fi25, fi25, v1)
                                if len(v28) > 0 {
                                    v28[0] = (fi25 - 522)
                                }
                                _ = v28
                            }
                            if false {
                                fmt.Println(fi25, fi25, fi26)
                                switch fi25 {
                                case 32767:
                                    var v29 float64 = ((-2.25 + 3.14159) / float64(-2.25))
                                    _ = v29
                                default:
                                    var v30 uint = 255
                                    _ = v30
                                }
                            }
                            if false {
                                fmt.Println(v4)
                            }
                        }
                        var v31 int16 = 545
                        var v32 uint32 = 670
                        _ = v31
                        _ = v32
                    }
                    fmt.Println(v3)
                    v1 += 504
                    v1 += (v1 + (v1 >> ((v1) & 63)))
                } else {
                    if true {
                        v3 = v3
                        v4 -= (^(v4))
                    }
                    if (!(0 < 42)) {
                        if true {
                            v33 := []int{(fi25 % (((fi25 & fi26)) | 1)), fi26, fi25}
                            switch (fi26 >> ((v1) & 63)) {
                            case 2147483647:
                                if false {
                                    if false {
                                        v2 = (v2 / v2)
                                    } else {
                                        fmt.Println(fi26)
                                        if true {
                                            v34 := GenS1{F0: v4, F1: uint(v4), F2: int32(v4)}
                                            _ = len(v33)
                                            _ = v34
                                        }
                                    }
                                    var v35 string = "hello"
                                    _ = v35
                                } else {
                                    var v36 int64 = 1
                                    var v37 float32 = ((v2 / 3.14159) * v2)
                                    var v38 int16 = 1000
                                    _ = v36
                                    _ = v37
                                    _ = v38
                                }
                            case 721:
                                v4 = (v4 & 128)
                            case -874:
                                fmt.Println(fi25)
                            }
                            var v39 int32 = -833
                            var v40 int16 = -815
                            _ = v33
                            _ = v39
                            _ = v40
                        }
                        v41 := GenS1{F0: (^(v4)), F1: v1, F2: 2147483647}
                        switch "a" {
                        case "":
                            switch ((("foo bar" + "\n\t") + "\n\t") + "foo bar") {
                            case "foo bar":
                                fmt.Println(fi26, v4, v4)
                            default:
                                fmt.Println(v1, fi25, v1)
                            }
                        case "\n\t":
                            if (-821 >= (fi25 - fi25)) {
                                v2 -= v2
                                v42 := []int{int(v1)}
                                v3 = v3
                                _ = v42
                            } else {
                                switch (("a" + "\n\t") + ("hello" + (((("foo bar" + "hello") + "\n\t") + ("hello" + "a")) + "foo bar"))) {
                                case "foo bar":
                                    var v43 float32 = 3.14159
                                    _ = v43
                                default:
                                    if false {
                                        switch (fi26 % (((fi26 * 32768)) | 1)) {
                                        case 127:
                                            var v44 string = "\n\t"
                                            _ = v44
                                        default:
                                            v45 := GenS3{F0: 65535, F1: 867, F2: ("" + "\n\t")}
                                            _ = v45
                                        }
                                    }
                                }
                            }
                        }
                        fmt.Println(v3, fi25)
                        v46 := []int{}
                        _ = v41
                        _ = v46
                    }
                }
                if (v4 >= (^(v4))) {
                    v4 = (v4 % ((128) | 1))
                    v47 := GenS4{F0: -2147483648, F1: (fi26 | (fi25 - (fi26 - fi26)))}
                    _ = v47
                } else {
                    v48 := GenS2{F0: (v4 % (((^(v4))) | 1)), F1: (fi25 ^ fi26), F2: 42}
                    if true {
                        var v49 uint64 = v3
                        _ = v49
                    } else {
                        fmt.Println(v1, v4)
                        if false {
                            v2 = v2
                            v50 := GenS5{F0: 534, F1: -2147483648, F2: uint16(v4)}
                            v3 -= 4294967295
                            _ = v50
                        } else {
                            v4 = v4
                            var v51 int64 = 1
                            _ = v51
                        }
                        v52 := []int{}
                        v53 := GenS6{F0: "\n\t"}
                        v4 = 255
                        _ = v52
                        _ = v53
                    }
                    v1 += 255
                    _ = v48
                }
                _ = fi26
            }
            if (42 > -663) {
                var v54 uint64 = v3
                var v55 uint32 = uint32(v4)
                _ = v54
                _ = v55
            }
        }
        _ = fi25
    }
    var v56 int8 = 42
    fmt.Println(v4)
    v57 := []int{-1, 2147483647, 42, int(v1)}
    fmt.Println(v2, v4, v2)
    v1 -= (v1 - v1)
    var v58 uint32 = uint32(v4)
    switch "hello" {
    case "":
        v1 += v1
    case "foo bar":
        switch -2147483648 {
        case 4294967295:
            for fi59 := 0; fi59 < 5; fi59++ {
                if ((55 < (v56 + v56)) || false) {
                    v60 := GenS4{F0: -661, F1: (fi59 | 32768)}
                    _ = v60
                } else {
                    v2 = ((v2 + (v2 + (0.0 + v2))) / v2)
                    v1 -= v1
                }
                var v61 int = fi59
                if len(v57) >= 2 {
                    v57 = v57[1:len(v57)]
                }
                _ = fi59
                _ = v61
            }
        case 0:
            for fi62 := 0; fi62 < 15; fi62++ {
                v1 = uint(v58)
                switch int(v4) {
                case -1:
                    var v63 int = (fi62 - (-(fi62)))
                    _ = v63
                case 32767:
                    if len(v57) >= 2 {
                        v57 = v57[1:len(v57)]
                    }
                }
                v57 = append(v57, fi62)
                _ = fi62
            }
        }
    case "a":
        var v64 float64 = float64(v2)
        _ = v64
    }
    var v65 int = 217
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v56:", v56)
    fmt.Println("slicelen v57:", len(v57), v57)
    fmt.Println("v58:", v58)
    fmt.Println("v65:", v65)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v56
    _ = v57
    _ = v58
    _ = v65
}
