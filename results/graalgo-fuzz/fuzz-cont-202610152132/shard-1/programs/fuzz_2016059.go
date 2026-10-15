package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint64; F1 string; F2 uint8 }
type GenS3 struct { F0 uint8; F1 bool }
type GenS4 struct { F0 int16; F1 uint8 }
type GenS5 struct { F0 uint8; F1 uint32; F2 int16 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 bool; F1 int16; F2 int64 }
type GenS8 struct { F0 string; F1 int64; F2 bool }
type GenS9 struct { F0 uint16; F1 int; F2 uint8 }
func genFn0(p0 int8) bool {
    v5 := GenS1{F0: 0}
    for fi6 := 0; fi6 < 8; fi6++ {
        for fi7 := 0; fi7 < 2; fi7++ {
            if false {
                v8 := GenS2{F0: 42, F1: "a", F2: 2}
                v9 := GenS2{F0: 18446744073709551615, F1: ("foo bar" + ("hello" + "foo bar")), F2: 143}
                if (42 >= 0) {
                    p0 += 42
                    p0 = p0
                    v10 := []int{-9223372036854775808, fi6, (^(fi7))}
                    if len(v10) > 0 {
                        v10[0] = -2147483648
                    }
                    p0 += (p0 * 100)
                    p0 = p0
                    _ = v10
                }
                p0 = p0
                _ = v8
                _ = v9
            }
            _ = fi7
        }
        _ = fi6
    }
    v11 := []int{2147483647, 42}
    var v12 uint32 = 42
    if (808 == 18446744073709551615) {
        switch "a" {
        case "a":
            var v13 uint16 = 0
            _ = v13
        case "":
            var v14 uint64 = 0
            _ = v14
        case "\n\t":
            var v15 float64 = 1e10
            _ = v15
        default:
            switch ("hello" + "") {
            case "\n\t":
                v12 = (v12 | (v12 - 2))
            default:
                v16 := GenS3{F0: 127, F1: true}
                _ = v16
            }
        }
        p0 += p0
    }
    _ = p0
    _ = v11
    _ = v12
    _ = v5
    return true
}

func main() {
    var v1 string = ("" + "hello")
    v2 := make(map[string]int)
    v3 := []uint32{12, 931, 74, 283}
    var v4 float32 = (((((0.0 + float32((float64((3.14159 * -2.25)) / -2.25))) - ((0.0 + (0.1 - -2.25)) * (0.0 * (0.1 / 1.0)))) / 1.0) + (0.0 / float32(0.1))) * 3.14159)
    v4 = v4
    var v17 int8 = 0
    v18 := GenS4{F0: -1, F1: 128}
    switch -121 {
    case -117:
        v19 := GenS5{F0: 103, F1: 4294967295, F2: -682}
        _ = v19
    case 0:
        var v20 uint16 = 32767
        _ = v20
    case 2147483647:
        v21 := GenS6{F0: 95}
        _ = v21
    default:
        if ((((0.1 / 3.14159) * (-2.25 / ((3.14159 - ((((-2.25 + ((float64(0.1) / (1e10 / 0.1)) + -2.25)) * (0.1 / -2.25)) * (3.14159 / 1.0)) - float64(float32(1e10)))) + ((0.0 * 3.14159) - 0.1)))) - (-2.25 / 1.0)) > float64(float32(1e10))) {
            if (!false) {
                for fi22 := 0; fi22 < 19; fi22++ {
                    var v23 int = (fi22 / ((232) | 1))
                    _ = fi22
                    _ = v23
                }
                if (243 <= 1) {
                    if false {
                        if (2 < 148) {
                            if (!false) {
                                switch 32767 {
                                case 664:
                                    delete(v2, ((v1 + ("a" + "a")) + v1))
                                case 503:
                                    var v24 uint32 = 42
                                    _ = v24
                                default:
                                    if false {
                                        if true {
                                            var v25 uint8 = 8
                                            switch ("\n\t" + "hello") {
                                            case "hello":
                                                var v26 int = 0
                                                _ = v26
                                            case "":
                                                v27 := GenS4{F0: -1, F1: v25}
                                                _ = v27
                                            case "\n\t":
                                                v4 = (v4 + float32(((-2.25 / 1.5) / ((-2.25 * 0.0) - 1e10))))
                                            }
                                            _ = v25
                                        }
                                        var v28 uint8 = 128
                                        for fi29 := 0; fi29 < 1; fi29++ {
                                            if false {
                                                v1 = "\n\t"
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                            }
                                            v30 := GenS6{F0: 789}
                                            var v31 string = ((("\n\t" + ((v1 + ("foo bar" + v1)) + v1)) + ((v1 + v1) + v1)) + ("" + v1))
                                            _ = fi29
                                            _ = v30
                                            _ = v31
                                        }
                                        _ = v28
                                    }
                                }
                                fmt.Println(v1)
                            }
                        } else {
                            if true {
                                v2[v1] = 42
                            } else {
                                _ = len(v2)
                                fmt.Println(v4, v1, v17)
                                v3 = append(v3, 536)
                            }
                            var v32 float64 = (((3.14159 - (1.5 + 0.0)) * 1e10) / 1.0)
                            _ = v32
                        }
                    }
                    if (true || false) {
                        v33 := GenS2{F0: 1, F1: (v1 + v1), F2: 225}
                        if (-2147483648 != 776) {
                            if false {
                                var v34 uint8 = 185
                                if (false && (!true)) {
                                    fmt.Println(v4)
                                }
                                switch (v1 + "") {
                                case "a":
                                    v35 := GenS7{F0: true, F1: -365, F2: -965}
                                    _ = v35
                                case "\n\t":
                                    if false {
                                        fmt.Println(v4, v1, v17)
                                        delete(v2, "\n\t")
                                    }
                                }
                                _ = v34
                            } else {
                                var v36 uint64 = 2
                                var v37 int64 = 1
                                _ = v36
                                _ = v37
                            }
                        }
                        _ = v33
                    }
                    v38 := GenS2{F0: 1, F1: "foo bar", F2: 42}
                    _ = v38
                }
                if true {
                    delete(v2, ((v1 + "foo bar") + "a"))
                } else {
                    v1 = "a"
                    if (((!(!(((v1 >= ("hello" + (v1 + "foo bar"))) && (!true)) || (true || true)))) && ((v1 > "foo bar") || false)) || ((((-2.25 + 3.14159) / (0.1 * ((-2.25 - -2.25) - 3.14159))) != (0.1 * 1e10)) || false)) {
                        fmt.Println(v4)
                    } else {
                        fmt.Println(v17)
                    }
                }
            }
            switch 1 {
            case 0:
                var v39 uint64 = 18446744073709551615
                _ = v39
            default:
                switch 1 {
                case 712:
                    if (28 != 49) {
                        var v40 uint64 = 2
                        var v41 int32 = 568
                        _ = v40
                        _ = v41
                    } else {
                        if false {
                            v4 -= v4
                            var v42 bool = false
                            _ = v42
                        }
                        fmt.Println(v4, v1, v1)
                    }
                case -717:
                    v43 := v2[v1]
                    _ = v43
                case -531:
                    for fi44 := 0; fi44 < 16; fi44++ {
                        var v45 int8 = v17
                        var v46 int64 = 42
                        _ = fi44
                        _ = v45
                        _ = v46
                    }
                default:
                    var v47 int16 = -264
                    _ = v47
                }
            }
        }
    }
    v17 -= (-(v17))
    if (!false) {
        var v48 int = 32767
        var v49 int32 = 1
        _ = v48
        _ = v49
    }
    v50 := GenS4{F0: -32768, F1: 105}
    for fi51 := 0; fi51 < 17; fi51++ {
        v52 := GenS8{F0: v1, F1: 1, F2: (128 != 127)}
        _ = fi51
        _ = v52
    }
    var v53 int16 = 0
    v2["hello"] = 699
    _ = len(v3)
    fmt.Println(v17)
    var v54 uint32 = 391
    switch 32768 {
    case 68:
        v2[(v1 + v1)] = 2147483647
    default:
        if true {
            v3 = append(v3, (v54 * 2147483647))
            var v55 uint32 = v54
            _ = v55
        } else {
            for fi56 := 0; fi56 < 13; fi56++ {
                var v57 uint8 = 1
                _ = fi56
                _ = v57
            }
            if true {
                v17 -= (v17 | -1)
                v4 -= v4
                var v58 float64 = 1e10
                _ = v58
            } else {
                v59 := GenS2{F0: 4294967295, F1: ("" + v1), F2: 202}
                for fi60 := 0; fi60 < 9; fi60++ {
                    var v61 int32 = 42
                    _ = len(v3)
                    if true {
                        v62 := GenS9{F0: 457, F1: (fi60 ^ 2147483647), F2: 73}
                        _ = v62
                    }
                    _ = fi60
                    _ = v61
                }
                v63 := GenS2{F0: uint64(v54), F1: "a", F2: 0}
                _ = v59
                _ = v63
            }
            for fi64 := 0; fi64 < 3; fi64++ {
                fmt.Println(v54, v17, v4)
                if (v53 == 1000) {
                    switch fi64 {
                    case -249:
                        v53 = v53
                    case 339:
                        for fi65 := 0; fi65 < 17; fi65++ {
                            v66 := GenS2{F0: 42, F1: ("" + v1), F2: 255}
                            v54 -= v54
                            var v67 uint = 2
                            _ = fi65
                            _ = v66
                            _ = v67
                        }
                    default:
                        v53 = v53
                    }
                    var v68 uint32 = v54
                    _ = v68
                } else {
                    fmt.Println(v17, v53)
                    v4 -= -2.25
                }
                if false {
                    for fi69 := 0; fi69 < 18; fi69++ {
                        var v70 float32 = float32((3.14159 / 1.5))
                        switch (fi64 << ((18446744073709551615) & 63)) {
                        case -9223372036854775808:
                            v71 := v2[v1]
                            _ = v71
                        case 4294967295:
                            v53 = (v53 - (^(v53)))
                        }
                        switch 128 {
                        case 556:
                            v3 = append(v3, (v54 << ((4294967295) & 63)))
                        default:
                            fmt.Println(v70, v4)
                        }
                        _ = fi69
                        _ = v70
                    }
                } else {
                    for fi72 := 0; fi72 < 7; fi72++ {
                        var v73 uint = 255
                        _ = fi72
                        _ = v73
                    }
                    var v74 int = -2147483648
                    _ = v74
                }
                _ = fi64
            }
        }
    }
    var v75 int16 = (-(v53))
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v53:", v53)
    fmt.Println("v54:", v54)
    fmt.Println("v75:", v75)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v50
    _ = v53
    _ = v54
    _ = v75
}
