package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 bool; F1 uint }
type GenS4 struct { F0 bool; F1 int32 }
type GenS5 struct { F0 uint64; F1 uint16 }
type GenS6 struct { F0 uint8; F1 bool }
type GenS7 struct { F0 uint; F1 uint; F2 bool }
type GenS8 struct { F0 uint32; F1 bool }
type GenS9 struct { F0 int16; F1 int64; F2 int8 }
type GenS10 struct { F0 int16 }
type GenS11 struct { F0 int32 }
type GenS12 struct { F0 int32; F1 uint32 }
func genFn0() uint32 {
    var v6 int64 = -354
    for fi7 := 0; fi7 < 7; fi7++ {
        if (!(100 != -1)) {
            fmt.Println(fi7, v6)
            v6 -= 730
            var v8 float64 = (float64(0.1) / 1.0)
            _ = v8
        }
        switch "a" {
        case "":
            for fi9 := 0; fi9 < 6; fi9++ {
                v10 := []int{fi9, (fi9 % ((fi7) | 1))}
                var v11 uint = 146
                switch "a" {
                case "a":
                    if false {
                        var v12 int16 = 32767
                        _ = v12
                    }
                case "foo bar":
                    v13 := GenS1{F0: (fi7 >> (((v11 >> ((v11) & 63))) & 63))}
                    _ = v13
                case "":
                    if len(v10) > 0 {
                        v10[0] = fi9
                    }
                default:
                    v11 += (v11 & (v11 + 105))
                }
                _ = fi9
                _ = v10
                _ = v11
            }
        case "hello":
            fmt.Println(v6)
        default:
            if (true || false) {
                v6 = v6
                fmt.Println(fi7)
            } else {
                for fi14 := 0; fi14 < 10; fi14++ {
                    var v15 int = (fi14 + 1)
                    _ = fi14
                    _ = v15
                }
                var v16 uint16 = 394
                _ = v16
            }
        }
        var v17 uint32 = 896
        _ = fi7
        _ = v17
    }
    if (255 >= 255) {
        var v18 float32 = 1.5
        _ = v18
    } else {
        fmt.Println(v6)
        v6 += 892
        v6 = (v6 * v6)
    }
    switch ("a" + "foo bar") {
    case "a":
        for fi19 := 0; fi19 < 6; fi19++ {
            var v20 uint = 65535
            _ = fi19
            _ = v20
        }
    }
    _ = v6
    return 42
}
func genFn1(p0 int32) int64 {
    p0 = p0
    fmt.Println(p0)
    v21 := GenS2{F0: 127}
    fmt.Println(p0)
    _ = p0
    _ = v21
    return 1
}

func main() {
    v1 := make(map[int]int)
    var v2 string = "hello"
    var v3 int16 = -32768
    v4 := make(map[int]int)
    var v5 string = v2
    var v22 uint8 = 127
    v5 = v2
    if false {
        if true {
            v22 += v22
            v23 := GenS3{F0: (42 >= 1), F1: 171}
            _ = v23
        } else {
            v24 := GenS1{F0: -908}
            for fi25 := 0; fi25 < 19; fi25++ {
                var v26 uint = 317
                _ = fi25
                _ = v26
            }
            _ = v24
        }
    } else {
        for fi27 := 0; fi27 < 11; fi27++ {
            var v28 int32 = 42
            _ = fi27
            _ = v28
        }
    }
    v29 := GenS4{F0: true, F1: -635}
    v2 = v2
    if (-2.25 < ((((-2.25 / 3.14159) * (0.0 / 1.5)) + -2.25) * (((float32((0.1 / 3.14159)) / -2.25) * 0.0) + (3.14159 * 1e10)))) {
        v22 = 42
        if true {
            switch ((v5 + v5) + v5) {
            case "foo bar":
                var v30 uint = 42
                _ = v30
            default:
                v31 := GenS1{F0: 32767}
                _ = v31
            }
            v5 = (v2 + "")
        } else {
            switch 42 {
            case 1:
                var v32 uint16 = 255
                _ = v32
            default:
                var v33 string = (v2 + "hello")
                _ = v33
            }
            var v34 uint32 = 2147483647
            for fi35 := 0; fi35 < 10; fi35++ {
                if (((0.0 / -2.25) + -2.25) != float32((0.1 / float64(1.5)))) {
                    v3 += v3
                    var v36 int32 = -1
                    for fi37 := 0; fi37 < 8; fi37++ {
                        switch v2 {
                        case "foo bar":
                            if false {
                                var v38 float64 = (0.1 * (0.1 * 3.14159))
                                var v39 uint16 = 65535
                                if (!(0.0 > ((((0.0 + 0.0) * 3.14159) * float32(v38)) + float32((0.1 / (v38 * (0.1 / v38))))))) {
                                    v39 -= 2
                                } else {
                                    v34 += 65535
                                    var v40 uint16 = uint16(v34)
                                    _ = v40
                                }
                                _ = v38
                                _ = v39
                            }
                        case "a":
                            v41 := []int{}
                            _ = v41
                        }
                        fmt.Println(v36, v2, v22)
                        _ = fi37
                    }
                    _ = v36
                }
                switch ((v2 + v2) + ("" + v2)) {
                case "a":
                    for fi42 := 0; fi42 < 12; fi42++ {
                        var v43 float64 = 3.14159
                        v44 := GenS3{F0: (0 != 9223372036854775807), F1: 4294967295}
                        _ = fi42
                        _ = v43
                        _ = v44
                    }
                case "":
                    for fi45 := 0; fi45 < 5; fi45++ {
                        v46 := GenS5{F0: 453, F1: 812}
                        if (!(fi45 != -1)) {
                            v22 -= 2
                            var v47 int64 = 0
                            fmt.Println(v5)
                            _ = v47
                        }
                        _ = fi45
                        _ = v46
                    }
                case "\n\t":
                    fmt.Println(v3, v34, v22)
                }
                v5 = (("a" + v2) + ((v5 + (v2 + (v2 + v5))) + v5))
                _ = fi35
            }
            _ = v34
        }
        var v48 int8 = -17
        _ = v48
    } else {
        var v49 int64 = -1
        _ = v49
    }
    var v50 float32 = float32((1.5 * 3.14159))
    v51 := GenS6{F0: (v22 << ((0) & 63)), F1: true}
    v1[-754] = 127
    if (true && false) {
        v52 := GenS7{F0: 4294967295, F1: 2, F2: false}
        _ = v52
    } else {
        v5 = v2
        for fi53 := 0; fi53 < 10; fi53++ {
            for fi54 := 0; fi54 < 9; fi54++ {
                switch fi54 {
                case 2147483647:
                    v55 := GenS5{F0: 512, F1: 2}
                    _ = v55
                case -501:
                    switch v5 {
                    case "foo bar":
                        v5 = v5
                    case "":
                        if true {
                            if (true && ((843 > 638) || (((float64(3.14159) / (((3.14159 / 1.5) - 0.1) * (float64(float32((3.14159 + -2.25))) / 1e10))) + 1e10) > float64(0.1)))) {
                                switch (fi54 + (fi53 * (fi54 ^ (^(fi53))))) {
                                case 32768:
                                    var v56 uint = 42
                                    _ = v56
                                case 84:
                                    fmt.Println(v3, v3, v3)
                                }
                                var v57 float32 = float32((3.14159 - 0.0))
                                _ = v57
                            } else {
                                var v58 int = (fi53 - (^(fi54)))
                                _ = v58
                            }
                            v59 := []int{fi54, (fi54 % (((fi54 & (fi54 / ((128) | 1)))) | 1))}
                            v60 := GenS1{F0: -1}
                            _ = v59
                            _ = v60
                        } else {
                            v61 := GenS8{F0: 829, F1: false}
                            v62 := []int{}
                            v63 := GenS5{F0: 646, F1: 1}
                            delete(v4, fi53)
                            _ = v61
                            _ = v62
                            _ = v63
                        }
                    case "hello":
                        v64 := GenS3{F0: true, F1: 65535}
                        _ = v64
                    default:
                        var v65 float64 = (1e10 + 1.5)
                        _ = v65
                    }
                case 127:
                    v22 -= (v22 << ((739) & 63))
                }
                var v66 int = (fi54 << ((0) & 63))
                switch (v5 + v5) {
                case "foo bar":
                    v67 := []int{v66, 4294967295, (fi54 & fi54)}
                    _ = v67
                case "":
                    var v68 uint16 = 255
                    _ = v68
                }
                _ = fi54
                _ = v66
            }
            _ = fi53
        }
        v69 := GenS9{F0: -723, F1: 42, F2: 1}
        _ = v69
    }
    v2 = (("hello" + (v5 + v5)) + "")
    if (-979 != -2147483648) {
        v70 := GenS10{F0: (-(v3))}
        fmt.Println(v22, v5)
        v1[9223372036854775807] = 32768
        _ = v70
    } else {
        delete(v1, 4294967295)
        for fi71 := 0; fi71 < 13; fi71++ {
            var v72 int32 = -410
            if true {
                var v73 int64 = 1
                fmt.Println(v2, v2, v2)
                _ = v73
            }
            _ = fi71
            _ = v72
        }
        for fi74 := 0; fi74 < 13; fi74++ {
            v1[(fi74 / ((0) | 1))] = fi74
            _ = fi74
        }
    }
    v5 = (v2 + ((v5 + (v5 + "foo bar")) + v2))
    delete(v1, int(v3))
    var v75 uint8 = (v22 - (v22 << ((411) & 63)))
    fmt.Println(v2, v2)
    if (2 >= 511) {
        var v76 uint32 = genFn0()
        v22 += 2
        _ = v76
    }
    if (v5 >= (v5 + v2)) {
        for fi77 := 0; fi77 < 12; fi77++ {
            delete(v1, 1)
            v50 += 1.5
            v50 = v50
            _ = fi77
        }
        _ = len(v4)
        fmt.Println(v22, v22)
    } else {
        var v78 uint16 = 41
        for fi79 := 0; fi79 < 9; fi79++ {
            v3 += v3
            for fi80 := 0; fi80 < 19; fi80++ {
                if true {
                    v81 := GenS11{F0: 0}
                    var v82 int = 505
                    v83 := GenS12{F0: -1, F1: 694}
                    _ = v81
                    _ = v82
                    _ = v83
                }
                _ = fi80
            }
            var v84 int8 = -1
            _ = fi79
            _ = v84
        }
        _ = v78
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v75:", v75)
    _ = v1
    _ = v2
    _ = v22
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v51
    _ = v75
}
