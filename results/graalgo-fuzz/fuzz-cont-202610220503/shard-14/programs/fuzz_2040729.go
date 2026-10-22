package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint; F1 uint8; F2 int32 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint32; F1 uint64 }
type GenS6 struct { F0 uint32; F1 string; F2 uint8 }
type GenS7 struct { F0 bool; F1 int64; F2 int64 }
type GenS8 struct { F0 int8; F1 uint64 }
type GenS9 struct { F0 uint8; F1 uint8 }
type GenS10 struct { F0 bool; F1 uint32; F2 uint16 }

func main() {
    var v1 int = 279
    v2 := []string{("" + "hello"), "foo bar", "", ("foo bar" + "")}
    v3 := make(map[int]uint32)
    v4 := []int{(^(v1)), (v1 & (^(v1))), -1}
    var v5 int = v1
    if false {
        if false {
            var v6 int8 = -87
            var v7 int64 = 964
            _ = v6
            _ = v7
        } else {
            if false {
                var v8 string = "\n\t"
                if (2 > 687) {
                    v9 := GenS1{F0: v8}
                    _ = v9
                } else {
                    switch v5 {
                    case -1:
                        delete(v3, v1)
                    }
                    if false {
                        if (2147483647 <= v1) {
                            if ((32767 >= -1) || (714 > 984)) {
                                var v10 int16 = 405
                                switch (v8 + v8) {
                                case "foo bar":
                                    v11 := GenS1{F0: v8}
                                    _ = v11
                                case "":
                                    fmt.Println(v10, v10, v1)
                                case "\n\t":
                                    var v12 float32 = ((1.5 - -2.25) / 1e10)
                                    _ = v12
                                }
                                if ((-1 != -2147483648) && false) {
                                    switch (v8 + v8) {
                                    case "":
                                        v10 -= (^(v10))
                                    }
                                    v13 := GenS2{F0: 255, F1: 0, F2: -1}
                                    v14 := GenS1{F0: v8}
                                    _ = v13
                                    _ = v14
                                } else {
                                    if ((1 == -304) && false) {
                                        v8 = ""
                                        delete(v3, (v1 ^ v1))
                                    }
                                    if false {
                                        if false {
                                            v15 := GenS2{F0: 458, F1: 2, F2: -1}
                                            var v16 int64 = -296
                                            _ = v15
                                            _ = v16
                                        }
                                        switch v1 {
                                        case -1:
                                            if len(v2) > 0 {
                                                v2[0] = ("hello" + v8)
                                            }
                                        case 128:
                                            var v17 string = "foo bar"
                                            _ = v17
                                        case 32767:
                                            for fi18 := 0; fi18 < 6; fi18++ {
                                                v8 = v8
                                                v5 += v1
                                                _ = fi18
                                            }
                                        default:
                                            switch "hello" {
                                            case "hello":
                                                if false {
                                                    fmt.Println(v1, v1)
                                                } else {
                                                    v3[(v5 + v1)] = 760
                                                    for fi19 := 0; fi19 < 4; fi19++ {
                                                        _ = len(v3)
                                                        v1 = (-(fi19))
                                                        _ = fi19
                                                    }
                                                }
                                            default:
                                                for fi20 := 0; fi20 < 16; fi20++ {
                                                    for fi21 := 0; fi21 < 2; fi21++ {
                                                        fmt.Println(fi20)
                                                        var v22 uint8 = 127
                                                        _ = fi21
                                                        _ = v22
                                                    }
                                                    var v23 float64 = float64((0.1 - -2.25))
                                                    _ = fi20
                                                    _ = v23
                                                }
                                            }
                                        }
                                        v2 = append(v2, (v8 + (v8 + "a")))
                                    } else {
                                        if (!true) {
                                            v5 -= (v1 >> ((4294967295) & 63))
                                        }
                                        var v24 int32 = -657
                                        var v25 uint8 = 42
                                        _ = v24
                                        _ = v25
                                    }
                                }
                                _ = v10
                            }
                            v26 := GenS2{F0: 196, F1: 2, F2: 124}
                            _ = v26
                        }
                    } else {
                        v4 = append(v4, 84)
                    }
                    v27 := GenS2{F0: 809, F1: 1, F2: 1}
                    _ = v27
                }
                switch v5 {
                case 127:
                    var v28 string = (((("a" + (("hello" + v8) + (v8 + v8))) + ("\n\t" + v8)) + "foo bar") + (v8 + "\n\t"))
                    _ = v28
                }
                _ = v8
            } else {
                v1 -= v1
                for fi29 := 0; fi29 < 4; fi29++ {
                    v5 -= (v1 / ((v1) | 1))
                    if true {
                        _ = len(v3)
                        if len(v2) > 0 {
                            v2[0] = "\n\t"
                        }
                        v5 = (v5 << ((2) & 63))
                    } else {
                        v30 := GenS3{F0: (v1 / ((fi29) | 1))}
                        var v31 uint8 = 127
                        v32 := GenS3{F0: v5}
                        _ = v30
                        _ = v31
                        _ = v32
                    }
                    if true {
                        fmt.Println(v1, v5, v1)
                        for fi33 := 0; fi33 < 11; fi33++ {
                            var v34 string = "hello"
                            v3[v1] = 1
                            _ = fi33
                            _ = v34
                        }
                    }
                    _ = fi29
                }
            }
            for fi35 := 0; fi35 < 13; fi35++ {
                for fi36 := 0; fi36 < 10; fi36++ {
                    if false {
                        fmt.Println(fi36, fi35)
                    }
                    var v37 uint8 = 2
                    var v38 uint16 = 1
                    if true {
                        if (!(3.14159 != 3.14159)) {
                            switch "" {
                            case "a":
                                v1 -= int(v38)
                            case "":
                                var v39 uint8 = 128
                                _ = v39
                            }
                            var v40 int32 = 42
                            v41 := GenS4{F0: -128}
                            _ = v40
                            _ = v41
                        }
                        v1 += int(v38)
                        v42 := GenS1{F0: "foo bar"}
                        _ = v42
                    } else {
                        switch "foo bar" {
                        case "foo bar":
                            var v43 uint64 = 18446744073709551615
                            _ = v43
                        default:
                            v44 := GenS5{F0: 2147483647, F1: 2}
                            _ = v44
                        }
                    }
                    v45 := GenS4{F0: -128}
                    if false {
                        var v46 int8 = 127
                        var v47 int16 = 533
                        _ = v46
                        _ = v47
                    } else {
                        if true {
                            var v48 int8 = 51
                            v49 := GenS6{F0: 0, F1: "foo bar", F2: v37}
                            _ = v48
                            _ = v49
                        }
                        v50 := GenS7{F0: false, F1: -391, F2: 726}
                        v3[v1] = 2147483647
                        _ = v50
                    }
                    if ((v37 - v37) != (v37 - v37)) {
                        v3[fi36] = 2147483647
                        v1 = (fi35 | v1)
                        if len(v2) > 0 {
                            v2[0] = ""
                        }
                    }
                    v38 = v38
                    _ = len(v4)
                    _ = fi36
                    _ = v37
                    _ = v38
                    _ = v45
                }
                var v51 int16 = 1
                _ = fi35
                _ = v51
            }
            if true {
                var v52 int16 = 1
                _ = v52
            }
        }
        for fi53 := 0; fi53 < 4; fi53++ {
            if true {
                fmt.Println(v1)
            } else {
                for fi54 := 0; fi54 < 15; fi54++ {
                    fmt.Println(fi54, v5, v5)
                    var v55 int32 = 156
                    _ = fi54
                    _ = v55
                }
                v56 := GenS8{F0: -127, F1: 0}
                v1 += (fi53 / (((v5 >> ((4294967295) & 63))) | 1))
                _ = v56
            }
            v5 = (v1 >> ((65535) & 63))
            v2 = append(v2, "\n\t")
            _ = fi53
        }
    }
    if (!(2 > 128)) {
        switch v5 {
        case 2147483647:
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        default:
            _ = len(v3)
        }
        var v57 int = (v1 * (v1 * v5))
        v5 += (v5 * 32767)
        _ = v57
    }
    var v58 int = v1
    switch (v5 >> ((485) & 63)) {
    case -87:
        for fi59 := 0; fi59 < 11; fi59++ {
            var v60 int64 = 0
            v1 += -62
            v61 := GenS9{F0: 10, F1: 0}
            _ = fi59
            _ = v60
            _ = v61
        }
    }
    var v62 int16 = 32767
    for fi63 := 0; fi63 < 12; fi63++ {
        if false {
            v64 := GenS10{F0: true, F1: 938, F2: 1}
            _ = v64
        } else {
            switch (fi63 % ((-9223372036854775808) | 1)) {
            case 548:
                fmt.Println(fi63, v1, v58)
            default:
                if (!true) {
                    var v65 uint32 = 4294967295
                    v5 = (v58 / (((v1 * 32767)) | 1))
                    switch "" {
                    case "\n\t":
                        v58 -= 128
                    default:
                        delete(v3, v58)
                    }
                    _ = v65
                }
            }
            for fi66 := 0; fi66 < 16; fi66++ {
                _ = len(v3)
                v58 = fi63
                _ = len(v2)
                _ = fi66
            }
            var v67 uint32 = 65535
            _ = v67
        }
        var v68 uint = 42
        _ = fi63
        _ = v68
    }
    var v69 uint32 = 65535
    var v70 int64 = 0
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v58:", v58)
    fmt.Println("v62:", v62)
    fmt.Println("v69:", v69)
    fmt.Println("v70:", v70)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v58
    _ = v62
    _ = v69
    _ = v70
}
