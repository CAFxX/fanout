package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64 }
type GenS2 struct { F0 uint32; F1 uint64; F2 int }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 string; F1 uint32; F2 string }
type GenS5 struct { F0 int64; F1 uint16; F2 string }
type GenS6 struct { F0 uint16; F1 uint }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 uint8; F1 int }
func genFn0(p0 uint16, p1 int32, p2 int32) uint64 {
    p0 = 2
    v6 := []int{132, 0}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 42
}
func genFn1() int32 {
    v7 := GenS1{F0: 21, F1: 1}
    var v8 int64 = -9223372036854775808
    switch 635 {
    case 9223372036854775807:
        v9 := []int{0}
        _ = v9
    case 0:
        if (!(!false)) {
            v10 := GenS1{F0: 65535, F1: v8}
            _ = v10
        } else {
            if (1 != 0) {
                v11 := []int{1, 0, 793, 42}
                v11 = append(v11, -134)
                var v12 int64 = 0
                _ = v11
                _ = v12
            }
            for fi13 := 0; fi13 < 7; fi13++ {
                v14 := GenS2{F0: 0, F1: 340, F2: (fi13 >> ((0) & 63))}
                v15 := GenS1{F0: 781, F1: (v8 & (v8 >> ((65535) & 63)))}
                _ = fi13
                _ = v14
                _ = v15
            }
        }
    case 954:
        for fi16 := 0; fi16 < 18; fi16++ {
            v17 := GenS2{F0: 65535, F1: 274, F2: (fi16 * fi16)}
            if false {
                var v18 uint16 = 32767
                _ = v18
            }
            switch "foo bar" {
            case "a":
                v19 := GenS1{F0: 1, F1: v8}
                _ = v19
            case "hello":
                v20 := []int{128, fi16, -9223372036854775808}
                _ = v20
            case "foo bar":
                if false {
                    if (1 >= 2) {
                        fmt.Println(fi16, v8)
                    } else {
                        v21 := GenS1{F0: 65535, F1: v8}
                        for fi22 := 0; fi22 < 15; fi22++ {
                            v23 := GenS1{F0: 784, F1: (v8 ^ v8)}
                            if (false && (42 < 487)) {
                                var v24 uint16 = 0
                                v8 += 9223372036854775807
                                var v25 uint8 = 2
                                _ = v24
                                _ = v25
                            } else {
                                var v26 uint = 0
                                fmt.Println(v26, fi22)
                                v27 := []int{fi16}
                                _ = len(v27)
                                _ = v26
                                _ = v27
                            }
                            v28 := GenS2{F0: 2147483648, F1: 2, F2: (fi22 | 2147483647)}
                            v29 := GenS1{F0: 2, F1: -1}
                            var v30 uint = 770
                            switch ("foo bar" + "a") {
                            case "\n\t":
                                fmt.Println(fi22, fi16, fi22)
                            case "a":
                                switch ("\n\t" + ("hello" + ("\n\t" + "foo bar"))) {
                                case "\n\t":
                                    if (false && (342 >= 2147483647)) {
                                        v31 := []int{fi22}
                                        switch "a" {
                                        case "hello":
                                            switch fi16 {
                                            case -239:
                                                var v32 uint64 = genFn0(65535, -1, 29)
                                                _ = v32
                                            case 128:
                                                var v33 float64 = 0.0
                                                _ = v33
                                            case 0:
                                                v30 = v30
                                            }
                                            var v34 int16 = 0
                                            _ = v34
                                        default:
                                            var v35 int8 = 0
                                            _ = v35
                                        }
                                        var v36 float32 = (3.14159 + (0.1 * ((-2.25 - 1e10) / (1e10 + float32(3.14159)))))
                                        _ = v31
                                        _ = v36
                                    } else {
                                        if true {
                                            var v37 float32 = ((((((float32(3.14159) + 1.5) - (0.1 + -2.25)) - 1e10) / 1.5) / float32(1e10)) - -2.25)
                                            _ = v37
                                        } else {
                                            fmt.Println(fi22, v8, v30)
                                            switch "hello" {
                                            case "":
                                                var v38 uint64 = 153
                                                v8 = (v8 - v8)
                                                v30 -= (v30 - v30)
                                                v38 -= 1
                                                v8 = (v8 << (((v30 >> ((4294967295) & 63))) & 63))
                                                _ = v38
                                            }
                                            switch fi16 {
                                            case -785:
                                                v39 := []int{}
                                                _ = v39
                                            case 287:
                                                var v40 float32 = 1e10
                                                _ = v40
                                            }
                                        }
                                        v8 = (v8 | (v8 & 0))
                                        v41 := []int{(fi22 | (fi22 | fi16)), (fi16 << (((v30 / (((v30 >> ((v30) & 63))) | 1))) & 63))}
                                        _ = v41
                                    }
                                case "":
                                    switch ("\n\t" + "hello") {
                                    case "":
                                        v42 := []int{fi16}
                                        _ = v42
                                    case "a":
                                        var v43 float32 = -2.25
                                        _ = v43
                                    }
                                }
                            default:
                                fmt.Println(v8)
                            }
                            _ = fi22
                            _ = v23
                            _ = v28
                            _ = v29
                            _ = v30
                        }
                        var v44 uint32 = 756
                        _ = v21
                        _ = v44
                    }
                    if false {
                        for fi45 := 0; fi45 < 18; fi45++ {
                            v8 = (v8 >> ((4294967295) & 63))
                            var v46 int16 = 32767
                            v47 := GenS3{F0: 42}
                            _ = fi45
                            _ = v46
                            _ = v47
                        }
                    }
                    v48 := GenS4{F0: "\n\t", F1: 4294967295, F2: (("a" + "\n\t") + "hello")}
                    _ = v48
                }
            }
            _ = fi16
            _ = v17
        }
    }
    switch "foo bar" {
    case "hello":
        fmt.Println(v8)
    case "":
        var v49 int32 = 0
        _ = v49
    case "foo bar":
        v50 := []int{-81, 2147483647, 884}
        _ = v50
    }
    _ = v7
    _ = v8
    return -1
}

func main() {
    var v1 float64 = 1.5
    var v2 float32 = 3.14159
    var v3 string = "foo bar"
    var v4 uint64 = 13
    var v5 uint8 = 0
    fmt.Println(v4, v3, v2)
    v3 = (v3 + (("hello" + (v3 + "hello")) + v3))
    fmt.Println(v1, v3)
    v51 := []int{-44, 2147483647, 750, 32768}
    for fi52 := 0; fi52 < 7; fi52++ {
        for fi53 := 0; fi53 < 15; fi53++ {
            v5 = 255
            var v54 uint32 = 65535
            _ = fi53
            _ = v54
        }
        _ = fi52
    }
    fmt.Println(v1, v2)
    fmt.Println(v1, v4, v5)
    if false {
        v51 = append(v51, 971)
        v55 := GenS4{F0: (((((v3 + v3) + v3) + ((v3 + ((v3 + (v3 + "foo bar")) + "foo bar")) + "\n\t")) + v3) + v3), F1: 65535, F2: v3}
        _ = v55
    }
    if (v5 == (v5 & uint8(v4))) {
        v3 = (("hello" + ((v3 + (v3 + "foo bar")) + v3)) + "hello")
        if false {
            var v56 float64 = (float64(v2) / v1)
            for fi57 := 0; fi57 < 11; fi57++ {
                var v58 int32 = 2147483647
                var v59 uint64 = (v4 | v4)
                _ = fi57
                _ = v58
                _ = v59
            }
            for fi60 := 0; fi60 < 6; fi60++ {
                v61 := GenS3{F0: 334}
                if true {
                    v62 := GenS5{F0: 9223372036854775807, F1: 255, F2: v3}
                    if len(v51) > 0 {
                        v51[0] = fi60
                    }
                    fmt.Println(v4, v2)
                    _ = v62
                }
                _ = fi60
                _ = v61
            }
            _ = v56
        } else {
            var v63 float32 = (float32(v1) - v2)
            _ = v63
        }
        v64 := GenS4{F0: ("hello" + (v3 + v3)), F1: 65535, F2: (("a" + v3) + v3)}
        _ = v64
    } else {
        for fi65 := 0; fi65 < 9; fi65++ {
            for fi66 := 0; fi66 < 3; fi66++ {
                var v67 uint32 = uint32(v4)
                var v68 float32 = (v2 / (v2 / (v2 * 3.14159)))
                v3 = "hello"
                v69 := GenS6{F0: uint16(v67), F1: 0}
                v67 += (^(v67))
                _ = fi66
                _ = v67
                _ = v68
                _ = v69
            }
            v1 -= v1
            _ = fi65
        }
        switch 253 {
        case 42:
            var v70 string = "foo bar"
            _ = v70
        default:
            for fi71 := 0; fi71 < 6; fi71++ {
                switch -1 {
                case -907:
                    if (2 < 286) {
                        var v72 uint64 = genFn0(0, 1, 0)
                        var v73 int64 = -9223372036854775808
                        switch (v3 + v3) {
                        case "\n\t":
                            v74 := GenS7{F0: -1}
                            _ = v74
                        case "foo bar":
                            v75 := GenS8{F0: v5, F1: -530}
                            _ = v75
                        case "":
                            var v76 float64 = (3.14159 + 1e10)
                            _ = v76
                        }
                        _ = v72
                        _ = v73
                    } else {
                        v4 -= (v4 | (v4 | v4))
                    }
                }
                _ = fi71
            }
        }
        var v77 float32 = v2
        _ = v77
    }
    if false {
        var v78 int64 = -981
        _ = v78
    }
    var v79 int8 = 1
    for fi80 := 0; fi80 < 9; fi80++ {
        v51 = append(v51, int(v5))
        var v81 int = 42
        v2 -= float32((v1 * (1e10 + float64(((v2 / (0.0 * v2)) * v2)))))
        _ = fi80
        _ = v81
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v51:", len(v51), v51)
    fmt.Println("v79:", v79)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v51
    _ = v79
}
