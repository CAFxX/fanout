package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string; F2 int8 }
type GenS2 struct { F0 bool; F1 int8; F2 uint8 }
type GenS3 struct { F0 uint32; F1 uint8; F2 uint8 }
type GenS4 struct { F0 int8; F1 int64 }
type GenS5 struct { F0 bool; F1 int16 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 uint16 }
func genFn0() string {
    for fi6 := 0; fi6 < 17; fi6++ {
        for fi7 := 0; fi7 < 18; fi7++ {
            var v8 int16 = 1
            fmt.Println(fi6, v8)
            switch "hello" {
            case "a":
                if (-128 <= 127) {
                    v9 := []int{fi7}
                    _ = v9
                } else {
                    v10 := GenS1{F0: 2, F1: "foo bar", F2: -1}
                    _ = v10
                }
            case "":
                var v11 int = fi7
                _ = v11
            case "\n\t":
                if true {
                    if false {
                        switch fi7 {
                        case -477:
                            v12 := []int{fi6}
                            _ = v12
                        case 127:
                            var v13 bool = (65535 >= 65535)
                            _ = v13
                        }
                        var v14 uint32 = 1
                        _ = v14
                    } else {
                        if true {
                            v15 := []int{-2147483648, (fi6 & fi6), fi7, fi7}
                            var v16 uint8 = 2
                            _ = v15
                            _ = v16
                        }
                        if (-2.25 < 0.0) {
                            switch "hello" {
                            case "hello":
                                v8 += v8
                            case "foo bar":
                                switch (fi7 * -1) {
                                case -2147483648:
                                    v17 := GenS2{F0: (-2147483648 <= 965), F1: 12, F2: 255}
                                    _ = v17
                                default:
                                    v8 += (-(v8))
                                }
                            case "":
                                v18 := []int{(fi6 * -1), (fi6 / (((fi7 << ((0) & 63))) | 1))}
                                _ = v18
                            default:
                                v19 := GenS1{F0: 0, F1: ("foo bar" + "foo bar"), F2: 42}
                                v8 = (-(v8))
                                v20 := GenS3{F0: 4294967295, F1: 0, F2: 2}
                                _ = v19
                                _ = v20
                            }
                            fmt.Println(fi7, fi6)
                        }
                        v21 := []int{fi6, fi7}
                        if true {
                            v8 += v8
                            _ = len(v21)
                            v8 -= v8
                        } else {
                            v8 = (v8 - 32767)
                            if true {
                                switch "foo bar" {
                                case "foo bar":
                                    v22 := GenS2{F0: (50 <= 0), F1: 30, F2: 128}
                                    _ = v22
                                case "\n\t":
                                    if true {
                                        var v23 int32 = 0
                                        v24 := GenS1{F0: 65535, F1: ("" + "foo bar"), F2: 1}
                                        var v25 int64 = -1
                                        _ = v23
                                        _ = v24
                                        _ = v25
                                    } else {
                                        var v26 int32 = -2147483648
                                        var v27 int32 = (v26 >> ((375) & 63))
                                        _ = v26
                                        _ = v27
                                    }
                                    switch ("" + "") {
                                    case "":
                                        if (255 != 18446744073709551615) {
                                            v8 -= v8
                                        } else {
                                            v28 := GenS4{F0: 127, F1: 1}
                                            _ = v28
                                        }
                                    case "foo bar":
                                        switch "foo bar" {
                                        case "\n\t":
                                            if (!(!true)) {
                                                var v29 string = "\n\t"
                                                _ = v29
                                            }
                                        case "a":
                                            switch (fi7 ^ (-(fi6))) {
                                            case 9223372036854775807:
                                                v30 := GenS1{F0: 42, F1: ("" + ("foo bar" + "hello")), F2: 100}
                                                _ = v30
                                            case 989:
                                                v31 := GenS5{F0: true, F1: 939}
                                                _ = v31
                                            default:
                                                var v32 string = "foo bar"
                                                _ = v32
                                            }
                                        default:
                                            var v33 uint64 = 2
                                            _ = v33
                                        }
                                    default:
                                        v21 = append(v21, -2147483648)
                                    }
                                    if len(v21) > 0 {
                                        v21[0] = (fi7 | (fi7 & fi6))
                                    }
                                }
                                v8 = v8
                            }
                            if false {
                                var v34 uint32 = 4294967295
                                _ = v34
                            } else {
                                if (!(!(float64((float32(3.14159) * -2.25)) < 3.14159))) {
                                    if len(v21) > 0 {
                                        v21[0] = fi6
                                    }
                                    switch fi7 {
                                    case 723:
                                        v8 -= v8
                                    case 0:
                                        fmt.Println(fi6)
                                    default:
                                        if (1 == 100) {
                                            fmt.Println(fi7)
                                        } else {
                                            v8 = (v8 % ((v8) | 1))
                                            v8 = v8
                                            var v35 int64 = 783
                                            _ = v35
                                        }
                                    }
                                } else {
                                    v8 = -1
                                }
                                if (v8 != (v8 & v8)) {
                                    var v36 int16 = v8
                                    _ = v36
                                } else {
                                    switch "foo bar" {
                                    case "\n\t":
                                        _ = len(v21)
                                    case "hello":
                                        v8 += v8
                                    default:
                                        var v37 bool = ("\n\t" == "hello")
                                        v37 = v37
                                        _ = v37
                                    }
                                    v8 -= (v8 + v8)
                                }
                                if (!(42 < -1)) {
                                    v38 := GenS1{F0: 2147483648, F1: "", F2: -1}
                                    _ = v38
                                } else {
                                    var v39 int32 = 683
                                    v39 -= (v39 - v39)
                                    if true {
                                        fmt.Println(fi6, v39, v8)
                                        v8 += (-(v8))
                                        if true {
                                            fmt.Println(v8, fi6, v39)
                                            v8 -= v8
                                        } else {
                                            switch (fi7 & fi7) {
                                            case -1:
                                                if len(v21) > 0 {
                                                    v21[0] = fi7
                                                }
                                            }
                                            v8 = (v8 | 0)
                                            v39 += (v39 & 1)
                                        }
                                    }
                                    _ = v39
                                }
                            }
                        }
                        _ = len(v21)
                        _ = v21
                    }
                    v40 := []int{fi7, (fi6 % (((fi6 % ((fi6) | 1))) | 1)), fi7, (fi7 / ((2147483647) | 1))}
                    var v41 uint8 = 127
                    _ = v40
                    _ = v41
                } else {
                    var v42 uint64 = 394
                    v43 := []int{}
                    v42 -= v42
                    fmt.Println(v42, fi6, fi6)
                    _ = v42
                    _ = v43
                }
            }
            _ = fi7
            _ = v8
        }
        _ = fi6
    }
    var v44 uint32 = 69
    fmt.Println(v44)
    _ = v44
    return (("hello" + ("\n\t" + ((("hello" + "a") + "\n\t") + (("" + "foo bar") + "")))) + "a")
}
func genFn1(p0 uint16, p1 int16) uint64 {
    var v45 string = ""
    fmt.Println(p0, v45)
    v46 := GenS1{F0: 42, F1: (v45 + v45), F2: 1}
    v47 := []int{127, 0, -356, 4294967295}
    _ = p0
    _ = p1
    _ = v45
    _ = v46
    _ = v47
    return 4294967295
}

func main() {
    var v1 int16 = 42
    v2 := make(map[string]string)
    var v3 float64 = ((0.0 - 0.0) * (1.5 * (-2.25 - (float64((3.14159 + (3.14159 * (-2.25 * (1.5 + ((0.0 * (0.1 / 0.1)) - 1.5)))))) / 3.14159))))
    v4 := make(map[int]string)
    v5 := []string{("" + "foo bar")}
    v3 -= v3
    for fi48 := 0; fi48 < 11; fi48++ {
        if true {
            var v49 float64 = v3
            var v50 float64 = (v3 + (v49 * (v3 * float64((1e10 / (3.14159 - (0.0 - (0.0 / (float32(v3) - (-2.25 / (1.5 * (((float32(v3) * (float32(-2.25) * ((-2.25 * 3.14159) - 0.1))) / 1.5) + 0.1))))))))))))
            _ = v49
            _ = v50
        } else {
            var v51 int8 = -128
            var v52 int8 = 49
            switch (fi48 >> ((18446744073709551615) & 63)) {
            case 127:
                if true {
                    switch "hello" {
                    case "":
                        v53 := GenS5{F0: (false || (0 > 0)), F1: (v1 >> ((0) & 63))}
                        _ = v53
                    case "foo bar":
                        v52 -= v51
                    }
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                }
            case 2147483647:
                v2["hello"] = ""
            case 641:
                fmt.Println(fi48)
            }
            _ = v51
            _ = v52
        }
        v1 -= (v1 + v1)
        for fi54 := 0; fi54 < 16; fi54++ {
            switch (^(fi48)) {
            case -59:
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            case 4294967295:
                v1 = v1
            case 739:
                var v55 string = "\n\t"
                _ = v55
            }
            v1 -= (v1 >> ((25) & 63))
            _ = fi54
        }
        _ = fi48
    }
    for fi56 := 0; fi56 < 17; fi56++ {
        var v57 uint64 = 0
        _ = fi56
        _ = v57
    }
    if false {
        v58 := GenS6{F0: v1}
        var v59 uint8 = 70
        if true {
            v3 = float64((3.14159 - ((0.1 - 0.1) - (1.5 * 1e10))))
            var v60 uint = 18446744073709551615
            switch ("\n\t" + "a") {
            case "":
                if (2 <= 383) {
                    var v61 float32 = -2.25
                    var v62 int16 = 32767
                    _ = v61
                    _ = v62
                }
            }
            _ = v60
        }
        _ = v58
        _ = v59
    } else {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        if len(v5) > 0 {
            v5[0] = "a"
        }
    }
    fmt.Println(v1, v3)
    if (1 == 435) {
        switch "a" {
        case "\n\t":
            v63 := GenS7{F0: 32767}
            _ = v63
        case "hello":
            if true {
                switch "foo bar" {
                case "foo bar":
                    v64 := GenS2{F0: (true && true), F1: 76, F2: 0}
                    _ = v64
                case "hello":
                    for fi65 := 0; fi65 < 13; fi65++ {
                        var v66 uint16 = 65
                        _ = fi65
                        _ = v66
                    }
                }
                delete(v2, "a")
            }
        default:
            var v67 bool = ((42 > 42) || false)
            _ = v67
        }
        var v68 float32 = 0.0
        v69 := GenS4{F0: 1, F1: 1}
        _ = v68
        _ = v69
    }
    v1 = (v1 << ((42) & 63))
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    fmt.Println(v3, v3)
    for fi70 := 0; fi70 < 14; fi70++ {
        fmt.Println(fi70, fi70, v3)
        for fi71 := 0; fi71 < 11; fi71++ {
            if (!(!false)) {
                v72 := GenS3{F0: 379, F1: 2, F2: 102}
                if false {
                    v3 = v3
                    v3 -= 0.1
                } else {
                    var v73 float32 = 3.14159
                    if (true || false) {
                        fmt.Println(v73, v3, v1)
                        v73 = v73
                    } else {
                        fmt.Println(v73)
                        var v74 int = fi71
                        _ = v74
                    }
                    fmt.Println(v1, fi71, v73)
                    _ = v73
                }
                if (true || true) {
                    delete(v2, (("a" + ("\n\t" + ("a" + "hello"))) + "a"))
                }
                _ = v72
            }
            _ = fi71
        }
        _ = fi70
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
