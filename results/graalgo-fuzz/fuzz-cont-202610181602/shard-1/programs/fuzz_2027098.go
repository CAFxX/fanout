package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string }
type GenS2 struct { F0 string; F1 uint64; F2 string }
type GenS3 struct { F0 uint; F1 uint16; F2 uint8 }
type GenS4 struct { F0 int64; F1 string }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 uint }
func genFn0(p0 string, p1 uint) int64 {
    fmt.Println(p1, p1)
    var v6 int64 = -1
    for fi7 := 0; fi7 < 6; fi7++ {
        v8 := []int{(fi7 << (((p1 & 255)) & 63))}
        var v9 int16 = -32768
        _ = fi7
        _ = v8
        _ = v9
    }
    var v10 uint16 = 2
    v11 := []int{-1, -1, 1, 2147483647}
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v6
    return (v6 << (((p1 << ((p1) & 63))) & 63))
}

func main() {
    var v1 int32 = 42
    var v2 float32 = 0.1
    v3 := make(map[uint32]uint32)
    v4 := make(map[string]string)
    var v5 bool = (!(2147483648 <= 4294967295))
    fmt.Println(v2, v1)
    switch -1 {
    case 9223372036854775807:
        if false {
            v12 := GenS1{F0: 100, F1: "foo bar"}
            switch "a" {
            case "foo bar":
                v3[65535] = 42
            case "\n\t":
                v13 := GenS1{F0: 1, F1: "a"}
                _ = v13
            case "":
                v2 -= (float32((0.0 - 0.1)) / (v2 + 0.1))
            }
            for fi14 := 0; fi14 < 9; fi14++ {
                if (-1 <= 197) {
                    v4[("" + "\n\t")] = "foo bar"
                    for fi15 := 0; fi15 < 15; fi15++ {
                        switch ("\n\t" + "foo bar") {
                        case "a":
                            fmt.Println(fi15, v1)
                        case "foo bar":
                            delete(v3, 2147483648)
                        case "hello":
                            fmt.Println(v1)
                        default:
                            v2 = (v2 / (v2 - (v2 + v2)))
                        }
                        if (18446744073709551615 == 42) {
                            v16 := GenS1{F0: 42, F1: "hello"}
                            _ = v16
                        }
                        _ = fi15
                    }
                    var v17 string = "hello"
                    _ = v17
                }
                fmt.Println(v2)
                var v18 int64 = 0
                _ = fi14
                _ = v18
            }
            _ = v12
        } else {
            fmt.Println(v5)
        }
    case 127:
        switch 1 {
        case -9223372036854775808:
            for fi19 := 0; fi19 < 6; fi19++ {
                v20 := []int{(-(fi19)), fi19, (fi19 | fi19), fi19}
                _ = fi19
                _ = v20
            }
        case 42:
            var v21 int32 = v1
            _ = v21
        case 2147483647:
            for fi22 := 0; fi22 < 4; fi22++ {
                fmt.Println(v5, fi22, v1)
                v1 -= (^(v1))
                _ = fi22
            }
        }
    }
    fmt.Println(v1)
    v5 = ((v1 == (v1 & (v1 << ((0) & 63)))) && (v5 && v5))
    switch 32768 {
    case 42:
        var v23 float32 = v2
        _ = v23
    case 373:
        if (0 == 255) {
            delete(v4, "")
            _ = len(v4)
        } else {
            var v24 string = (("hello" + "a") + "")
            var v25 int16 = 0
            _ = v24
            _ = v25
        }
    default:
        var v26 uint8 = 255
        _ = v26
    }
    fmt.Println(v5)
    switch "" {
    case "\n\t":
        v1 = (v1 ^ v1)
    }
    fmt.Println(v2, v1)
    if (true && ((!v5) || (1.5 > 0.0))) {
        for fi27 := 0; fi27 < 3; fi27++ {
            var v28 int8 = -10
            var v29 int16 = 32767
            v5 = (817 >= genFn0(("" + ("\n\t" + "")), 18446744073709551615))
            _ = fi27
            _ = v28
            _ = v29
        }
        for fi30 := 0; fi30 < 13; fi30++ {
            v31 := []int{int(v1), (fi30 / ((fi30) | 1))}
            _ = fi30
            _ = v31
        }
    }
    v32 := GenS2{F0: "a", F1: 561, F2: ("foo bar" + "hello")}
    var v33 string = "hello"
    switch "foo bar" {
    case "hello":
        for fi34 := 0; fi34 < 7; fi34++ {
            var v35 uint8 = 239
            switch "" {
            case "hello":
                v33 = (((((v33 + v33) + "foo bar") + v33) + "a") + "a")
            case "\n\t":
                v2 = (1e10 + v2)
            case "foo bar":
                v36 := GenS3{F0: 2, F1: 255, F2: 128}
                _ = v36
            default:
                v37 := GenS3{F0: 2, F1: 32767, F2: v35}
                _ = v37
            }
            v1 = v1
            _ = fi34
            _ = v35
        }
    case "a":
        var v38 int8 = 17
        _ = v38
    case "":
        v39 := []int{}
        _ = v39
    }
    var v40 int32 = v1
    if v5 {
        _ = len(v4)
        for fi41 := 0; fi41 < 12; fi41++ {
            var v42 int32 = 2147483647
            v43 := GenS3{F0: 777, F1: 2, F2: 2}
            switch ((v33 + v33) + "") {
            case "foo bar":
                switch fi41 {
                case 675:
                    for fi44 := 0; fi44 < 7; fi44++ {
                        v45 := GenS2{F0: v33, F1: 0, F2: "hello"}
                        _ = fi44
                        _ = v45
                    }
                default:
                    var v46 uint8 = 1
                    _ = v46
                }
            case "hello":
                var v47 int = 128
                _ = v47
            case "":
                if (v40 != 2147483647) {
                    v48 := GenS1{F0: -128, F1: v33}
                    var v49 uint16 = 818
                    v50 := GenS4{F0: -133, F1: "a"}
                    _ = v48
                    _ = v49
                    _ = v50
                }
            }
            _ = fi41
            _ = v42
            _ = v43
        }
    } else {
        for fi51 := 0; fi51 < 16; fi51++ {
            v5 = v5
            if (-105 >= 100) {
                if v5 {
                    var v52 string = v33
                    _ = v52
                } else {
                    v40 += (v1 + (v1 % (((v40 % ((v40) | 1))) | 1)))
                }
                var v53 uint16 = 888
                _ = v53
            }
            _ = fi51
        }
    }
    for fi54 := 0; fi54 < 7; fi54++ {
        if (-32768 > 0) {
            var v55 int16 = 995
            for fi56 := 0; fi56 < 20; fi56++ {
                v3[1] = 2147483648
                if (((-(v55)) > v55) || false) {
                    fmt.Println(fi56, v1, fi54)
                } else {
                    fmt.Println(v33, v33, v40)
                    v57 := []int{(fi56 ^ (fi54 - fi56))}
                    var v58 int = (-(fi54))
                    if v5 {
                        var v59 int8 = 1
                        var v60 float64 = 0.0
                        var v61 float64 = v60
                        _ = v59
                        _ = v60
                        _ = v61
                    } else {
                        if (v5 || v5) {
                            v5 = v5
                            if true {
                                if (!v5) {
                                    v40 = (^(v1))
                                    if v5 {
                                        fmt.Println(v40)
                                        fmt.Println(v58, v5, fi54)
                                        if (42 == 1) {
                                            v55 += int16(fi54)
                                            if v5 {
                                                v62 := GenS3{F0: 834, F1: 42, F2: 42}
                                                delete(v3, 2147483648)
                                                if v5 {
                                                    if len(v57) > 0 {
                                                        v57[0] = 745
                                                    }
                                                }
                                                _ = v62
                                            } else {
                                                v63 := GenS3{F0: 42, F1: 2, F2: 1}
                                                var v64 float32 = v2
                                                _ = v63
                                                _ = v64
                                            }
                                        }
                                        _ = len(v3)
                                    }
                                    var v65 bool = (-68 <= 0)
                                    _ = v65
                                }
                                v40 = 2147483647
                            }
                        } else {
                            v66 := GenS2{F0: ("\n\t" + v33), F1: 4294967295, F2: v33}
                            _ = v66
                        }
                    }
                    _ = v57
                    _ = v58
                }
                _ = fi56
            }
            _ = v55
        }
        _ = fi54
    }
    switch v33 {
    case "\n\t":
        for fi67 := 0; fi67 < 1; fi67++ {
            v3[470] = 717
            var v68 int16 = -1
            _ = fi67
            _ = v68
        }
    case "":
        _ = len(v3)
    case "foo bar":
        if v5 {
            if v5 {
                if (18446744073709551615 == 42) {
                    fmt.Println(v1, v2, v33)
                } else {
                    for fi69 := 0; fi69 < 20; fi69++ {
                        v70 := GenS5{F0: (v40 - v1)}
                        v71 := []int{fi69, (fi69 % ((fi69) | 1)), (fi69 + fi69)}
                        if v5 {
                            v1 -= v1
                            v40 = v40
                        } else {
                            for fi72 := 0; fi72 < 7; fi72++ {
                                delete(v3, 1)
                                switch 321 {
                                case 9223372036854775807:
                                    var v73 int32 = 0
                                    _ = v73
                                default:
                                    fmt.Println(fi72, fi69, fi72)
                                }
                                _ = fi72
                            }
                        }
                        _ = fi69
                        _ = v70
                        _ = v71
                    }
                    if (false || (20 < 0)) {
                        var v74 int64 = -1
                        switch v33 {
                        case "\n\t":
                            var v75 int32 = (v1 | v1)
                            _ = v75
                        case "hello":
                            v76 := []int{128, 4294967295, -14, 42}
                            _ = v76
                        default:
                            fmt.Println(v74, v1)
                        }
                        v1 += (v40 ^ v40)
                        _ = v74
                    } else {
                        fmt.Println(v2, v33, v40)
                        v77 := []int{0, -2147483648, 42, -960}
                        _ = v77
                    }
                }
                v78 := GenS3{F0: 2, F1: 257, F2: 1}
                _ = v78
            } else {
                if false {
                    if v5 {
                        var v79 int64 = 0
                        _ = v79
                    }
                } else {
                    var v80 uint = 1
                    _ = v80
                }
                v81 := GenS6{F0: 303}
                switch (v33 + v33) {
                case "hello":
                    v82 := []int{2147483647, 356, 366}
                    _ = v82
                }
                _ = v81
            }
            v1 = 2147483647
        } else {
            v83 := GenS6{F0: 65535}
            if v5 {
                for fi84 := 0; fi84 < 14; fi84++ {
                    v85 := GenS2{F0: (v33 + v33), F1: 2, F2: v33}
                    _ = fi84
                    _ = v85
                }
            }
            _ = v83
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v40
    _ = v5
}
