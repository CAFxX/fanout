package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 string; F1 uint64 }
type GenS5 struct { F0 uint; F1 uint32 }
type GenS6 struct { F0 int64; F1 int8; F2 int64 }
type GenS7 struct { F0 int64; F1 uint; F2 uint }
func genFn0(p0 uint16, p1 int16) int32 {
    if true {
        v5 := GenS1{F0: -71}
        v6 := GenS1{F0: 127}
        _ = v5
        _ = v6
    } else {
        v7 := GenS1{F0: 42}
        for fi8 := 0; fi8 < 10; fi8++ {
            fmt.Println(fi8, p1)
            var v9 uint32 = 65535
            _ = fi8
            _ = v9
        }
        if (p0 == 0) {
            if (!false) {
                var v10 string = ("a" + "hello")
                _ = v10
            }
            for fi11 := 0; fi11 < 7; fi11++ {
                if (!(1e10 == (-2.25 + 3.14159))) {
                    for fi12 := 0; fi12 < 4; fi12++ {
                        p0 = 672
                        _ = fi12
                    }
                    var v13 uint = 4294967295
                    _ = v13
                } else {
                    for fi14 := 0; fi14 < 3; fi14++ {
                        v15 := GenS2{F0: 2}
                        switch -703 {
                        case -9223372036854775808:
                            switch (fi14 - (fi14 << ((978) & 63))) {
                            case 32768:
                                switch 969 {
                                case -536:
                                    p0 += p0
                                }
                            case 0:
                                p1 += p1
                            }
                        case -272:
                            p1 -= -1
                        }
                        switch 9223372036854775807 {
                        case 32768:
                            switch "" {
                            case "hello":
                                if (!(true && (("foo bar" + ("" + "\n\t")) > ("a" + "hello")))) {
                                    switch "" {
                                    case "foo bar":
                                        v16 := []int{}
                                        var v17 bool = (381 > uint(p0))
                                        _ = v16
                                        _ = v17
                                    case "a":
                                        var v18 int16 = 843
                                        _ = v18
                                    case "\n\t":
                                        var v19 uint64 = 42
                                        _ = v19
                                    default:
                                        var v20 uint64 = 835
                                        _ = v20
                                    }
                                    p0 = p0
                                    if false {
                                        if false {
                                            v21 := GenS3{F0: 1}
                                            fmt.Println(p0, fi14)
                                            _ = v21
                                        } else {
                                            var v22 int = 2147483647
                                            _ = v22
                                        }
                                    } else {
                                        var v23 int = -322
                                        p1 += p1
                                        _ = v23
                                    }
                                } else {
                                    v24 := []int{fi14, fi11, fi14}
                                    _ = v24
                                }
                                v25 := GenS2{F0: 2}
                                _ = v25
                            case "foo bar":
                                fmt.Println(p1)
                            case "":
                                v26 := GenS3{F0: 585}
                                _ = v26
                            }
                        case -1:
                            v27 := []int{}
                            p1 += (p1 % ((p1) | 1))
                            var v28 float32 = 0.1
                            _ = v27
                            _ = v28
                        case 42:
                            fmt.Println(p1)
                        default:
                            v29 := []int{}
                            _ = v29
                        }
                        _ = fi14
                        _ = v15
                    }
                }
                var v30 uint64 = 42
                switch "" {
                case "":
                    if false {
                        v31 := []int{(fi11 + fi11), (fi11 ^ (fi11 >> ((1) & 63)))}
                        _ = v31
                    }
                case "hello":
                    var v32 bool = ((("" + ("a" + "a")) <= ((("foo bar" + "\n\t") + ("hello" + "")) + "\n\t")) || (32767 == (p1 & (p1 % (((p1 << ((374) & 63))) | 1)))))
                    _ = v32
                case "\n\t":
                    var v33 string = "foo bar"
                    _ = v33
                default:
                    v34 := GenS4{F0: "", F1: 353}
                    _ = v34
                }
                _ = fi11
                _ = v30
            }
            if (608 <= -1) {
                switch "a" {
                case "":
                    switch 42 {
                    case 192:
                        if (true && (0.0 == -2.25)) {
                            for fi35 := 0; fi35 < 6; fi35++ {
                                fmt.Println(p1)
                                var v36 int32 = 232
                                _ = fi35
                                _ = v36
                            }
                            var v37 float64 = (float64((((((1.5 - (-2.25 + (0.0 - (1.5 - 3.14159)))) / 0.1) - -2.25) * 3.14159) * float32(3.14159))) * (-2.25 * 1e10))
                            _ = v37
                        }
                    case 4294967295:
                        p0 -= (p0 & 255)
                    }
                case "a":
                    v38 := []int{-314}
                    _ = v38
                case "foo bar":
                    v39 := []int{906, -795}
                    _ = v39
                default:
                    p1 += (p1 * -264)
                }
            } else {
                fmt.Println(p1, p1)
            }
        }
        _ = v7
    }
    v40 := []int{2147483647, 127, -555, -987}
    if len(v40) > 0 {
        v40[0] = 763
    }
    var v41 uint8 = 127
    switch 9223372036854775807 {
    case -632:
        v40 = append(v40, 9223372036854775807)
    default:
        if len(v40) > 0 {
            v40[0] = -871
        }
    }
    _ = p0
    _ = p1
    _ = v40
    _ = v41
    return 2147483647
}
func genFn1(p0 uint32) int64 {
    fmt.Println(p0)
    fmt.Println(p0)
    for fi42 := 0; fi42 < 14; fi42++ {
        p0 += (^(p0))
        _ = fi42
    }
    var v43 uint32 = 65535
    _ = p0
    _ = v43
    return 1
}

func main() {
    v1 := []int{391, -444, -1}
    var v2 float64 = 0.1
    var v3 int = -670
    var v4 int8 = 42
    fmt.Println(v3, v3, v3)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if len(v1) > 0 {
        v1[0] = v3
    }
    v2 -= v2
    if false {
        var v44 int = v3
        for fi45 := 0; fi45 < 6; fi45++ {
            v1 = append(v1, -982)
            v1 = append(v1, fi45)
            for fi46 := 0; fi46 < 10; fi46++ {
                v47 := GenS3{F0: 4294967295}
                var v48 string = "hello"
                _ = fi46
                _ = v47
                _ = v48
            }
            _ = fi45
        }
        switch (("\n\t" + ("a" + "foo bar")) + "") {
        case "foo bar":
            var v49 int64 = -214
            _ = v49
        }
        _ = v44
    } else {
        var v50 int = v3
        for fi51 := 0; fi51 < 10; fi51++ {
            var v52 int32 = 0
            _ = fi51
            _ = v52
        }
        if true {
            v53 := GenS5{F0: 42, F1: 2147483647}
            _ = v53
        }
        _ = v50
    }
    fmt.Println(v3, v3)
    var v54 uint8 = 255
    v55 := GenS5{F0: 344, F1: 885}
    fmt.Println(v4, v54, v4)
    v56 := GenS2{F0: v54}
    for fi57 := 0; fi57 < 5; fi57++ {
        if false {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            v58 := GenS5{F0: 42, F1: 2}
            _ = v58
        } else {
            switch "foo bar" {
            case "hello":
                v2 = v2
            }
            fmt.Println(v4, v3)
            v59 := GenS6{F0: 513, F1: -128, F2: -1}
            _ = v59
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        if ((v54 | 120) == (^(v54))) {
            v4 -= (v4 | 100)
        }
        _ = fi57
    }
    for fi60 := 0; fi60 < 9; fi60++ {
        var v61 int8 = int8(fi60)
        _ = fi60
        _ = v61
    }
    var v62 int16 = 1000
    v63 := GenS2{F0: (v54 & v54)}
    v64 := GenS3{F0: 4294967295}
    if (!true) {
        var v65 int8 = v4
        _ = v65
    }
    for fi66 := 0; fi66 < 4; fi66++ {
        var v67 int8 = 1
        if true {
            v2 -= (float64(float32(v2)) + (((v2 / v2) / 1.0) * v2))
            v67 += (-(v4))
            for fi68 := 0; fi68 < 9; fi68++ {
                if false {
                    var v69 string = "a"
                    _ = len(v1)
                    var v70 uint16 = 32767
                    _ = v69
                    _ = v70
                } else {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                }
                v54 = 255
                fmt.Println(v54, v3)
                _ = fi68
            }
        }
        v67 -= 100
        _ = fi66
        _ = v67
    }
    for fi71 := 0; fi71 < 13; fi71++ {
        if true {
            _ = len(v1)
            for fi72 := 0; fi72 < 5; fi72++ {
                v73 := GenS3{F0: 1}
                _ = fi72
                _ = v73
            }
            for fi74 := 0; fi74 < 18; fi74++ {
                if (!(1.5 == float32(1.5))) {
                    if (-9223372036854775808 < 42) {
                        var v75 int8 = 25
                        if true {
                            v3 -= 4294967295
                            v3 -= fi71
                        }
                        _ = v75
                    } else {
                        if false {
                            if true {
                                var v76 float64 = 1.5
                                _ = len(v1)
                                if true {
                                    var v77 float64 = 1e10
                                    _ = v77
                                }
                                v78 := GenS7{F0: 394, F1: 4294967295, F2: 25}
                                if false {
                                    if (false && false) {
                                        if false {
                                            v4 = -128
                                        } else {
                                            fmt.Println(v54)
                                        }
                                        var v79 uint16 = 417
                                        if (65535 >= 2) {
                                            if (!(-2.25 > (float32(1e10) * (-2.25 + ((0.0 / ((((1e10 - -2.25) + (float32(v76) * (0.1 * 3.14159))) / 1.0) + 0.0)) - 1.5))))) {
                                                fmt.Println(fi74, v79, fi71)
                                                v54 = v54
                                            }
                                            v62 -= (v62 % (((v62 & (v62 - v62))) | 1))
                                            fmt.Println(v3)
                                        }
                                        _ = v79
                                    }
                                    var v80 string = "\n\t"
                                    _ = v80
                                }
                                v4 += 100
                                v81 := GenS2{F0: (v54 | (^(v54)))}
                                switch "" {
                                case "a":
                                    if ((!false) && (!true)) {
                                        v3 = 815
                                    } else {
                                        v2 -= -2.25
                                        v1 = append(v1, (v3 * 1))
                                        var v82 int8 = (v4 | (v4 ^ (v4 ^ (v4 ^ v4))))
                                        _ = v82
                                    }
                                default:
                                    if true {
                                        v3 = int(v54)
                                        v76 -= (1e10 + float64((0.1 / 3.14159)))
                                        if ((!false) && ((("\n\t" + "") + "") > "foo bar")) {
                                            v54 = 50
                                            v83 := GenS7{F0: genFn1(0), F1: 1, F2: 2}
                                            fmt.Println(v62, fi71)
                                            var v84 int = 462
                                            _ = v83
                                            _ = v84
                                        } else {
                                            fmt.Println(fi74, v2, v76)
                                        }
                                    }
                                }
                                _ = v76
                                _ = v78
                                _ = v81
                            }
                            var v85 string = "hello"
                            fmt.Println(v3, v85)
                            _ = v85
                        }
                    }
                    var v86 bool = ((float64(3.14159) * v2) == 0.0)
                    var v87 uint = 0
                    _ = v86
                    _ = v87
                }
                var v88 uint32 = 2
                _ = fi74
                _ = v88
            }
        }
        var v89 float64 = float64(0.1)
        v54 = (^(v54))
        _ = fi71
        _ = v89
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v54:", v54)
    fmt.Println("v62:", v62)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v54
    _ = v55
    _ = v56
    _ = v62
    _ = v63
    _ = v64
}
