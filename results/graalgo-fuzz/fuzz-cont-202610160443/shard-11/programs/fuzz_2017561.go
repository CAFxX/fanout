package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8; F2 int }
type GenS2 struct { F0 uint; F1 int8 }
type GenS3 struct { F0 int32; F1 int64; F2 int }
type GenS4 struct { F0 uint64; F1 int32; F2 int64 }
type GenS5 struct { F0 int; F1 uint8; F2 bool }
type GenS6 struct { F0 bool; F1 uint64 }
type GenS7 struct { F0 int64; F1 uint16 }
type GenS8 struct { F0 string }
func genFn0(p0 string, p1 bool) int32 {
    var v6 uint64 = 705
    v7 := []int{-2147483648, -9223372036854775808, 128, 55}
    p1 = false
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 0
}
func genFn1(p0 uint64, p1 uint64, p2 uint) int64 {
    v8 := []int{-2147483648}
    var v9 uint16 = uint16(p1)
    v10 := GenS1{F0: (("foo bar" + "foo bar") + "hello"), F1: 100, F2: 4294967295}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v8
    _ = v9
    return -1
}

func main() {
    v1 := []int{}
    var v2 uint8 = 42
    var v3 float32 = ((0.1 / (-2.25 * (((3.14159 + 0.0) / (3.14159 - ((((3.14159 / (-2.25 - 1.5)) + (0.0 + (0.0 + 1.5))) - ((1e10 + 1e10) + (((1.5 / -2.25) - 0.0) * (3.14159 + 1e10)))) - 3.14159))) - 0.1))) * ((-2.25 + 0.1) / 1.5))
    v4 := make(map[string]bool)
    v5 := make(map[int]bool)
    v3 = (((v3 * (-2.25 - ((float32((1.5 + 0.0)) * 3.14159) + 0.1))) + 1.5) / (1.5 - (((((v3 - v3) / v3) * ((v3 - v3) + v3)) + v3) * 0.0)))
    for fi11 := 0; fi11 < 1; fi11++ {
        v12 := GenS1{F0: "a", F1: -125, F2: 42}
        v2 -= (v2 + v2)
        _ = fi11
        _ = v12
    }
    v2 -= (v2 & v2)
    switch -496 {
    case 32768:
        v2 = 255
    case -203:
        if true {
            v3 += v3
            v3 = float32((0.0 * 0.0))
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
    default:
        if ((!(884 >= 286)) && (18446744073709551615 != 195)) {
            var v13 uint64 = 1
            _ = v13
        }
    }
    var v14 uint8 = v2
    if false {
        v2 = 42
    }
    if (true && false) {
        v3 = 3.14159
    }
    v15 := v5[-1]
    var v16 float32 = 0.0
    v3 += (v16 + 0.0)
    fmt.Println(v15, v2, v16)
    var v17 uint64 = 1
    v16 -= v3
    delete(v4, "foo bar")
    switch "\n\t" {
    case "a":
        for fi18 := 0; fi18 < 16; fi18++ {
            v15 = (-2.25 != float32(3.14159))
            v16 += ((v16 - float32(((1.5 / (3.14159 + float64(0.0))) + (0.1 * 0.1)))) / ((float32(float64(v16)) / (1e10 - v16)) / float32(1.5)))
            v19 := v5[(fi18 << ((255) & 63))]
            _ = fi18
            _ = v19
        }
    default:
        switch -834 {
        case -1:
            var v20 uint32 = 65535
            _ = v20
        case 0:
            if (!(!((v17 >> ((0) & 63)) != (v17 & 2)))) {
                if len(v1) > 0 {
                    v1[0] = 4294967295
                }
                fmt.Println(v3)
            } else {
                fmt.Println(v14)
                v14 = (v14 << ((1) & 63))
                if false {
                    v17 = v17
                    v2 = (^(v2))
                } else {
                    for fi21 := 0; fi21 < 16; fi21++ {
                        switch fi21 {
                        case 4294967295:
                            if (fi21 >= fi21) {
                                v22 := GenS1{F0: "", F1: 55, F2: (^(fi21))}
                                _ = v22
                            }
                        case 32767:
                            for fi23 := 0; fi23 < 13; fi23++ {
                                switch ("\n\t" + ("a" + "\n\t")) {
                                case "\n\t":
                                    var v24 bool = v15
                                    _ = v24
                                case "foo bar":
                                    fmt.Println(fi21, fi23, v14)
                                case "":
                                    if (!((0.1 + v16) >= 1.5)) {
                                        var v25 int16 = 42
                                        var v26 string = ("foo bar" + ("hello" + "\n\t"))
                                        var v27 uint = 65535
                                        fmt.Println(v3, v27)
                                        _ = v25
                                        _ = v26
                                        _ = v27
                                    }
                                }
                                if v15 {
                                    if v15 {
                                        v28 := GenS2{F0: 255, F1: 0}
                                        switch ("\n\t" + "\n\t") {
                                        case "foo bar":
                                            v14 -= 42
                                        case "hello":
                                            var v29 uint = 160
                                            _ = v29
                                        default:
                                            var v30 uint8 = 2
                                            if (v15 && false) {
                                                var v31 uint32 = 2147483648
                                                var v32 string = ("foo bar" + "foo bar")
                                                v33 := v5[-9223372036854775808]
                                                _ = v31
                                                _ = v32
                                                _ = v33
                                            }
                                            var v34 uint16 = 65535
                                            _ = v30
                                            _ = v34
                                        }
                                        v35 := GenS3{F0: 1, F1: 121, F2: fi23}
                                        _ = v28
                                        _ = v35
                                    }
                                }
                                v1 = append(v1, fi23)
                                _ = fi23
                            }
                        case 128:
                            v1 = append(v1, 1)
                        }
                        v2 = v14
                        var v36 int64 = 9223372036854775807
                        _ = fi21
                        _ = v36
                    }
                    v15 = ((38 <= 32767) || false)
                }
            }
        default:
            switch 4294967295 {
            case 32768:
                _ = len(v4)
            case -2147483648:
                if false {
                    var v37 int = -1
                    v14 -= 42
                    switch (v37 / (((v37 & v37)) | 1)) {
                    case -1:
                        for fi38 := 0; fi38 < 3; fi38++ {
                            v3 = (v16 * float32((1.5 / 1.0)))
                            if ("foo bar" == ("foo bar" + ("\n\t" + ""))) {
                                switch (-(v37)) {
                                case -2147483648:
                                    if v15 {
                                        v3 -= (0.0 / (-2.25 / float32(1.5)))
                                        var v39 uint64 = (v17 % ((219) | 1))
                                        _ = v39
                                    } else {
                                        v2 = v14
                                        var v40 int16 = 0
                                        _ = v40
                                    }
                                case -9223372036854775808:
                                    var v41 int8 = -128
                                    _ = v41
                                case -1:
                                    v42 := GenS3{F0: 934, F1: 1, F2: (v37 | (fi38 * (fi38 - 42)))}
                                    _ = v42
                                default:
                                    if (!false) {
                                        var v43 uint32 = 65535
                                        _ = v43
                                    } else {
                                        var v44 int = (fi38 - (v37 % (((v37 - fi38)) | 1)))
                                        for fi45 := 0; fi45 < 3; fi45++ {
                                            fmt.Println(v17)
                                            v46 := GenS2{F0: 18446744073709551615, F1: 42}
                                            var v47 string = "\n\t"
                                            v3 += 1e10
                                            _ = fi45
                                            _ = v46
                                            _ = v47
                                        }
                                        v15 = ("" >= "foo bar")
                                        _ = v44
                                    }
                                }
                                var v48 float32 = (float32(0.0) * v3)
                                _ = v48
                            }
                            fmt.Println(v14, v17)
                            _ = fi38
                        }
                    }
                    _ = v37
                }
            default:
                if v15 {
                    for fi49 := 0; fi49 < 14; fi49++ {
                        v17 = 18446744073709551615
                        _ = fi49
                    }
                } else {
                    if len(v1) > 0 {
                        v1[0] = -9223372036854775808
                    }
                    v50 := v5[194]
                    _ = v50
                }
            }
        }
    }
    var v51 uint8 = (v14 << ((255) & 63))
    if v15 {
        var v52 float64 = 1.5
        for fi53 := 0; fi53 < 1; fi53++ {
            for fi54 := 0; fi54 < 7; fi54++ {
                if ("" < ("foo bar" + ("" + "hello"))) {
                    var v55 int32 = 0
                    var v56 int = fi54
                    _ = v55
                    _ = v56
                }
                switch (^(fi53)) {
                case 127:
                    v52 = v52
                default:
                    var v57 float64 = v52
                    _ = v57
                }
                var v58 int64 = genFn1((v17 + (v17 - (v17 - 863))), 0, 0)
                _ = fi54
                _ = v58
            }
            var v59 int32 = 2147483647
            var v60 int8 = -1
            _ = fi53
            _ = v59
            _ = v60
        }
        _ = v52
    } else {
        if v15 {
            var v61 uint = 2
            if false {
                switch "\n\t" {
                case "hello":
                    var v62 string = ""
                    _ = v62
                default:
                    v2 = (^(v2))
                }
                for fi63 := 0; fi63 < 6; fi63++ {
                    v64 := GenS4{F0: v17, F1: genFn0("\n\t", false), F2: -1}
                    if ((v17 - (v17 >> (((v61 - v61)) & 63))) < v17) {
                        v17 -= v17
                        if (v16 <= 3.14159) {
                            var v65 float32 = 3.14159
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            switch "\n\t" {
                            case "":
                                for fi66 := 0; fi66 < 11; fi66++ {
                                    v67 := GenS5{F0: 9223372036854775807, F1: v14, F2: ((v17 <= v17) || ((int(v17) < (^(fi63))) || ((v51 / ((v2) | 1)) != v51)))}
                                    if false {
                                        var v68 uint16 = 1
                                        if len(v1) > 0 {
                                            v1[0] = 1
                                        }
                                        _ = v68
                                    } else {
                                        v61 += 18446744073709551615
                                        v69 := GenS6{F0: v15, F1: v17}
                                        _ = v69
                                    }
                                    _ = fi66
                                    _ = v67
                                }
                            default:
                                var v70 uint = (v61 % (((v61 ^ v61)) | 1))
                                _ = v70
                            }
                            _ = v65
                        } else {
                            v71 := GenS7{F0: 9223372036854775807, F1: 0}
                            for fi72 := 0; fi72 < 16; fi72++ {
                                var v73 uint32 = uint32(v61)
                                v74 := GenS7{F0: 1, F1: 503}
                                v15 = ("\n\t" < "a")
                                v75 := GenS7{F0: 0, F1: 1}
                                _ = fi72
                                _ = v73
                                _ = v74
                                _ = v75
                            }
                            _ = v71
                        }
                        v76 := GenS5{F0: 2147483647, F1: 0, F2: ((false && false) || (4294967295 < (v61 + 65535)))}
                        _ = v76
                    }
                    _ = fi63
                    _ = v64
                }
                var v77 int64 = -8
                _ = v77
            }
            var v78 uint = (v61 - v61)
            _ = v61
            _ = v78
        } else {
            for fi79 := 0; fi79 < 12; fi79++ {
                for fi80 := 0; fi80 < 6; fi80++ {
                    if false {
                        var v81 uint = 65535
                        switch (fi80 % (((fi79 >> ((0) & 63))) | 1)) {
                        case 42:
                            v81 -= v81
                        case 1:
                            v3 = v3
                        case -927:
                            var v82 int8 = 42
                            _ = v82
                        }
                        _ = v81
                    }
                    v83 := GenS8{F0: "\n\t"}
                    _ = fi80
                    _ = v83
                }
                fmt.Println(v17)
                _ = fi79
            }
            if (2147483647 > 42) {
                if false {
                    v84 := GenS5{F0: int(v51), F1: v51, F2: v15}
                    _ = v84
                }
                for fi85 := 0; fi85 < 3; fi85++ {
                    switch ("\n\t" + ("hello" + "a")) {
                    case "foo bar":
                        if (!false) {
                            v14 = v14
                            v4[(("a" + ("" + ("a" + ("a" + "hello")))) + "hello")] = (!false)
                        }
                    case "":
                        var v86 int8 = -20
                        _ = v86
                    }
                    v87 := GenS5{F0: 37, F1: 127, F2: false}
                    _ = fi85
                    _ = v87
                }
            }
        }
        v88 := GenS7{F0: -9223372036854775808, F1: 2}
        _ = v88
    }
    v89 := GenS4{F0: 1, F1: 2147483647, F2: -1}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v51:", v51)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v51
    _ = v89
}
