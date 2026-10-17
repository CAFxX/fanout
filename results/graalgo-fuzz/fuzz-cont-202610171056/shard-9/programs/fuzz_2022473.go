package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16 }
type GenS2 struct { F0 int32; F1 string }
type GenS3 struct { F0 int16; F1 int8; F2 uint64 }
type GenS4 struct { F0 int32; F1 bool; F2 int }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 int64 }
func genFn0(p0 uint32, p1 uint16, p2 uint8) int8 {
    switch 42 {
    case -1:
        p1 -= p1
    case 1:
        p2 -= (p2 | (p2 % ((255) | 1)))
    case 501:
        var v5 string = ("\n\t" + (("foo bar" + "hello") + ""))
        _ = v5
    default:
        var v6 int64 = 0
        _ = v6
    }
    if (!true) {
        var v7 int = 32767
        switch "hello" {
        case "hello":
            v7 -= (v7 << ((240) & 63))
        default:
            v8 := []int{v7, (^(v7)), (^(v7))}
            _ = v8
        }
        p0 = uint32(p2)
        _ = v7
    } else {
        p1 = p1
        p0 = 65535
        if true {
            switch "hello" {
            case "a":
                p1 -= (p1 & (p1 - (^(p1))))
            case "\n\t":
                var v9 uint32 = p0
                _ = v9
            default:
                v10 := GenS1{F0: -236, F1: -32768}
                _ = v10
            }
            for fi11 := 0; fi11 < 12; fi11++ {
                var v12 uint16 = p1
                v13 := []int{fi11, (-(fi11)), fi11}
                _ = fi11
                _ = v12
                _ = v13
            }
            if false {
                fmt.Println(p0, p0, p2)
                for fi14 := 0; fi14 < 1; fi14++ {
                    for fi15 := 0; fi15 < 10; fi15++ {
                        v16 := GenS1{F0: fi15, F1: -1}
                        v17 := GenS1{F0: fi14, F1: -1}
                        _ = fi15
                        _ = v16
                        _ = v17
                    }
                    switch int(p2) {
                    case -9223372036854775808:
                        for fi18 := 0; fi18 < 19; fi18++ {
                            var v19 int32 = -2147483648
                            if ((-9223372036854775808 >= 42) && false) {
                                switch (("hello" + "foo bar") + ((("\n\t" + (("" + "") + ("a" + "\n\t"))) + "\n\t") + (("hello" + ("" + ("foo bar" + "a"))) + ("foo bar" + ("hello" + "hello"))))) {
                                case "foo bar":
                                    fmt.Println(p2, fi14)
                                }
                                if (((9223372036854775807 == 309) || true) || false) {
                                    if (int64(fi14) == 9223372036854775807) {
                                        var v20 uint64 = 727
                                        v21 := []int{(fi18 / ((9223372036854775807) | 1)), (fi14 ^ 32768), (-(fi18))}
                                        _ = v20
                                        _ = v21
                                    }
                                    switch fi14 {
                                    case 128:
                                        p0 += (p0 >> ((42) & 63))
                                    case 137:
                                        var v22 uint32 = uint32(p2)
                                        _ = v22
                                    }
                                    v23 := []int{}
                                    p1 = (p1 - 396)
                                    p1 += 547
                                    _ = v23
                                }
                            } else {
                                v24 := []int{(-(fi18)), fi14, (fi18 + (fi14 | (fi14 + -1))), fi14}
                                switch "\n\t" {
                                case "":
                                    var v25 uint64 = 42
                                    _ = v25
                                case "a":
                                    var v26 uint = 556
                                    _ = v26
                                case "foo bar":
                                    var v27 int32 = (v19 * v19)
                                    _ = v27
                                default:
                                    v28 := GenS2{F0: 0, F1: ("foo bar" + ("hello" + ((("\n\t" + "a") + ("foo bar" + (("hello" + "") + ("" + "a")))) + "")))}
                                    _ = v28
                                }
                                var v29 int32 = v19
                                _ = v24
                                _ = v29
                            }
                            p2 += (^(p2))
                            _ = fi18
                            _ = v19
                        }
                    }
                    p1 = 1
                    _ = fi14
                }
                v30 := GenS2{F0: 2147483647, F1: ("\n\t" + "hello")}
                _ = v30
            }
        } else {
            p1 = p1
            for fi31 := 0; fi31 < 20; fi31++ {
                p2 += p2
                _ = fi31
            }
            switch 100 {
            case 32767:
                p0 = 0
            case -203:
                if (p2 > 128) {
                    if true {
                        var v32 string = ""
                        _ = v32
                    } else {
                        v33 := []int{32768}
                        _ = v33
                    }
                    var v34 uint = 42
                    _ = v34
                } else {
                    var v35 bool = (false && true)
                    v36 := []int{2147483647, -577, 0, 725}
                    _ = v35
                    _ = v36
                }
            }
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return 0
}
func genFn1() uint64 {
    v37 := []int{}
    for fi38 := 0; fi38 < 6; fi38++ {
        var v39 int8 = 127
        _ = fi38
        _ = v39
    }
    switch "" {
    case "hello":
        var v40 uint64 = 1
        _ = v40
    default:
        var v41 float32 = 1.5
        _ = v41
    }
    var v42 uint8 = 128
    _ = v37
    _ = v42
    return 257
}

func main() {
    v1 := []uint32{}
    v2 := make(map[uint32]string)
    var v3 uint64 = 209
    var v4 bool = false
    for fi43 := 0; fi43 < 17; fi43++ {
        v44 := GenS3{F0: 0, F1: 1, F2: 567}
        var v45 string = (("foo bar" + (("foo bar" + "\n\t") + "foo bar")) + "foo bar")
        var v46 uint = 18446744073709551615
        _ = fi43
        _ = v44
        _ = v45
        _ = v46
    }
    for fi47 := 0; fi47 < 16; fi47++ {
        for fi48 := 0; fi48 < 14; fi48++ {
            v3 -= v3
            v4 = v4
            v49 := GenS4{F0: 42, F1: v4, F2: (-(fi48))}
            _ = fi48
            _ = v49
        }
        _ = fi47
    }
    var v50 uint32 = 1
    var v51 int64 = 42
    v3 += (v3 << ((484) & 63))
    v50 -= v50
    v52 := GenS3{F0: -32768, F1: 42, F2: (v3 & (^(v3)))}
    for fi53 := 0; fi53 < 15; fi53++ {
        v3 = (^(v3))
        switch 1 {
        case -7:
            if v4 {
                var v54 float64 = 1e10
                var v55 uint8 = 85
                fmt.Println(fi53)
                _ = v54
                _ = v55
            } else {
                var v56 int16 = -32768
                switch "\n\t" {
                case "foo bar":
                    v51 = v51
                case "a":
                    for fi57 := 0; fi57 < 10; fi57++ {
                        v50 = 42
                        v51 += v51
                        if true {
                            _ = len(v1)
                        } else {
                            if ((v50 - 645) != (v50 | v50)) {
                                if ((true || (v4 || (0 <= 289))) || false) {
                                    v58 := GenS5{F0: 255}
                                    var v59 uint32 = v50
                                    _ = v58
                                    _ = v59
                                } else {
                                    v60 := GenS2{F0: 1, F1: ("foo bar" + "hello")}
                                    _ = v60
                                }
                            } else {
                                v61 := GenS6{F0: v56}
                                delete(v2, v50)
                                if ((v4 && false) || (0 == 42)) {
                                    fmt.Println(v51, v56, v56)
                                    v62 := GenS6{F0: (v56 + (v56 * 1))}
                                    _ = v62
                                } else {
                                    if false {
                                        v63 := GenS7{F0: v51}
                                        v4 = (true && false)
                                        var v64 uint16 = 42
                                        _ = v63
                                        _ = v64
                                    }
                                }
                                _ = v61
                            }
                            if v4 {
                                var v65 float32 = 1.5
                                v66 := v2[v50]
                                switch fi57 {
                                case 680:
                                    var v67 float64 = float64(v65)
                                    _ = v67
                                case 127:
                                    var v68 bool = false
                                    _ = v68
                                default:
                                    switch (v66 + "") {
                                    case "hello":
                                        v56 -= 42
                                    case "a":
                                        fmt.Println(v51, v65, v4)
                                    case "foo bar":
                                        if (((v66 + v66) + "a") < ((v66 + "") + v66)) {
                                            v3 -= (v3 << ((4294967295) & 63))
                                            var v69 float64 = 3.14159
                                            v65 = float32(v69)
                                            _ = v69
                                        } else {
                                            switch fi57 {
                                            case 127:
                                                v66 = v66
                                                v4 = (true && v4)
                                                v50 = 2147483647
                                                var v70 int64 = v51
                                                _ = v70
                                            case -345:
                                                switch v66 {
                                                case "foo bar":
                                                    var v71 int64 = (v51 % (((^(v51))) | 1))
                                                    _ = v71
                                                default:
                                                    switch v66 {
                                                    case "a":
                                                        v72 := GenS4{F0: 42, F1: v4, F2: 128}
                                                        _ = v72
                                                    case "foo bar":
                                                        var v73 int64 = v51
                                                        _ = v73
                                                    case "hello":
                                                        fmt.Println(fi57)
                                                    default:
                                                        if true {
                                                            fmt.Println(v51, v51, v56)
                                                            switch 463 {
                                                            case 9223372036854775807:
                                                                fmt.Println(v3, fi53)
                                                            }
                                                            v74 := GenS2{F0: 0, F1: "hello"}
                                                            _ = v74
                                                        } else {
                                                            if (65535 <= 529) {
                                                                v75 := GenS5{F0: 1}
                                                                switch ("hello" + (v66 + v66)) {
                                                                case "":
                                                                    v76 := v2[2147483647]
                                                                    _ = v76
                                                                default:
                                                                    switch "a" {
                                                                    case "a":
                                                                        fmt.Println(v51, fi53)
                                                                    default:
                                                                        var v77 uint8 = 244
                                                                        _ = v77
                                                                    }
                                                                }
                                                                v1 = append(v1, 2)
                                                                _ = v75
                                                            } else {
                                                                v51 = -1
                                                            }
                                                        }
                                                    }
                                                }
                                            }
                                        }
                                    default:
                                        if len(v1) > 0 {
                                            v1[0] = 369
                                        }
                                        v66 = "\n\t"
                                    }
                                }
                                _ = v65
                                _ = v66
                            } else {
                                var v78 int32 = 0
                                if v4 {
                                    var v79 float32 = ((0.1 / 1.0) * float32(3.14159))
                                    v80 := GenS7{F0: v51}
                                    var v81 string = ("hello" + "hello")
                                    fmt.Println(fi57, fi53)
                                    v82 := GenS1{F0: 32768, F1: v56}
                                    _ = v79
                                    _ = v80
                                    _ = v81
                                    _ = v82
                                }
                                _ = v78
                            }
                            if v4 {
                                v3 -= v3
                            } else {
                                v2[150] = "hello"
                                v83 := GenS1{F0: 439, F1: (v56 / ((42) | 1))}
                                v51 -= (^(v51))
                                _ = v83
                            }
                        }
                        _ = fi57
                    }
                default:
                    v3 += (v3 / ((v3) | 1))
                }
                v3 = v3
                _ = v56
            }
        case 2147483647:
            v51 += 9223372036854775807
        }
        _ = fi53
    }
    _ = len(v2)
    var v84 int64 = v51
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    fmt.Println("v84:", v84)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v50
    _ = v51
    _ = v52
    _ = v84
}
