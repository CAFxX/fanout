package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint; F1 uint32; F2 int16 }
type GenS3 struct { F0 uint; F1 int8 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int64; F1 int16; F2 bool }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 int32 }
func genFn0(p0 int16, p1 int16) uint16 {
    fmt.Println(p0)
    v7 := []int{42, 9223372036854775807}
    _ = p0
    _ = p1
    _ = v7
    return 32767
}
func genFn1() uint {
    if false {
        var v8 float32 = 1e10
        for fi9 := 0; fi9 < 3; fi9++ {
            if false {
                v10 := []int{(fi9 * -569), fi9, fi9, fi9}
                if len(v10) >= 2 {
                    v10 = v10[1:len(v10)]
                }
                if len(v10) >= 2 {
                    v10 = v10[1:len(v10)]
                }
                _ = v10
            } else {
                v11 := GenS1{F0: ("" + ("hello" + ("" + (("" + "") + "hello"))))}
                v12 := []int{fi9, (fi9 % (((fi9 - 128)) | 1)), fi9}
                _ = v11
                _ = v12
            }
            _ = fi9
        }
        _ = v8
    }
    for fi13 := 0; fi13 < 15; fi13++ {
        var v14 uint8 = 42
        v14 += 190
        var v15 bool = (false || true)
        _ = fi13
        _ = v14
        _ = v15
    }
    if ((!false) || true) {
        var v16 int = 4294967295
        _ = v16
    } else {
        var v17 string = "foo bar"
        _ = v17
    }
    switch -9223372036854775808 {
    case 128:
        for fi18 := 0; fi18 < 9; fi18++ {
            for fi19 := 0; fi19 < 8; fi19++ {
                var v20 int64 = -9223372036854775808
                if (true || ((v20 ^ (v20 - 9223372036854775807)) <= (v20 * v20))) {
                    fmt.Println(fi18, v20)
                    v21 := []int{0, 2147483647, fi18, -1}
                    v20 = (v20 - (v20 & (v20 ^ -1)))
                    _ = v21
                }
                _ = fi19
                _ = v20
            }
            _ = fi18
        }
    case -2147483648:
        var v22 float32 = ((((1e10 - 1e10) * (1e10 + ((0.1 + -2.25) / 1.0))) / 0.1) + 1.5)
        _ = v22
    case 453:
        var v23 float64 = 0.0
        _ = v23
    }
    return 881
}

func main() {
    var v1 int8 = 99
    v2 := make(map[string]bool)
    var v3 int32 = -1
    var v4 float32 = 0.1
    var v5 string = ("a" + "foo bar")
    v6 := []string{"a", (v5 + ("\n\t" + (v5 + "a")))}
    v5 = "hello"
    v24 := GenS2{F0: 903, F1: 2147483647, F2: 1000}
    if (960 == 42) {
        switch -230 {
        case -426:
            for fi25 := 0; fi25 < 20; fi25++ {
                if (42 >= 2) {
                    var v26 float64 = 0.0
                    _ = v26
                }
                _ = fi25
            }
        case 32768:
            _ = len(v2)
        default:
            if true {
                var v27 uint16 = 255
                var v28 int = -408
                for fi29 := 0; fi29 < 15; fi29++ {
                    for fi30 := 0; fi30 < 15; fi30++ {
                        v27 = v27
                        if (((fi29 >> ((1) & 63)) != (^(v28))) || true) {
                            v27 += v27
                            var v31 uint64 = 4294967295
                            switch v28 {
                            case 127:
                                fmt.Println(v1)
                                var v32 uint32 = 585
                                if false {
                                    var v33 uint16 = genFn0(1000, -1)
                                    v34 := GenS2{F0: genFn1(), F1: (v32 / ((v32) | 1)), F2: 724}
                                    var v35 int8 = 38
                                    _ = v33
                                    _ = v34
                                    _ = v35
                                }
                                _ = v32
                            case -269:
                                fmt.Println(v1, v4, fi30)
                            default:
                                v1 += v1
                            }
                            fmt.Println(v28, v1, fi29)
                            v31 -= (v31 & v31)
                            if false {
                                v36 := GenS3{F0: 65535, F1: v1}
                                fmt.Println(v4, fi30)
                                _ = v36
                            }
                            fmt.Println(v27)
                            _ = v31
                        } else {
                            v4 += v4
                        }
                        if true {
                            var v37 float32 = -2.25
                            if true {
                                if false {
                                    v28 = fi29
                                    var v38 uint16 = (v27 << ((569) & 63))
                                    _ = v38
                                } else {
                                    switch v5 {
                                    case "a":
                                        if (-371 > v3) {
                                            _ = len(v2)
                                            if true {
                                                v39 := GenS3{F0: 65535, F1: v1}
                                                v5 = v5
                                                _ = v39
                                            } else {
                                                v3 = (v3 << ((2) & 63))
                                                var v40 string = "hello"
                                                var v41 int32 = (^(v3))
                                                var v42 int8 = v1
                                                _ = v40
                                                _ = v41
                                                _ = v42
                                            }
                                            switch v5 {
                                            case "a":
                                                _ = len(v2)
                                            default:
                                                v3 = v3
                                            }
                                        } else {
                                            fmt.Println(fi29, v5)
                                            if true {
                                                v37 += v37
                                                delete(v2, v5)
                                                fmt.Println(v27)
                                                fmt.Println(v27, v5, v37)
                                                var v43 float64 = -2.25
                                                _ = v43
                                            }
                                            _ = len(v2)
                                            v37 = 0.0
                                            var v44 uint8 = 1
                                            var v45 int32 = (v3 ^ (v3 % ((v3) | 1)))
                                            _ = v44
                                            _ = v45
                                        }
                                    case "\n\t":
                                        switch fi29 {
                                        case -2147483648:
                                            v46 := GenS3{F0: 342, F1: (v1 ^ v1)}
                                            _ = v46
                                        }
                                    case "hello":
                                        v4 -= 1.5
                                    }
                                }
                                switch v5 {
                                case "\n\t":
                                    var v47 int32 = 0
                                    _ = v47
                                case "hello":
                                    var v48 int64 = -9223372036854775808
                                    _ = v48
                                case "foo bar":
                                    v4 = 1e10
                                }
                                fmt.Println(v4)
                                v37 -= ((v4 * (v37 * v4)) * v37)
                            } else {
                                switch v28 {
                                case -2147483648:
                                    fmt.Println(v1)
                                case 42:
                                    fmt.Println(fi30, fi29, v4)
                                }
                                v49 := GenS4{F0: fi30}
                                _ = v49
                            }
                            var v50 bool = (!((!true) && false))
                            _ = v37
                            _ = v50
                        }
                        _ = fi30
                    }
                    if true {
                        for fi51 := 0; fi51 < 10; fi51++ {
                            v3 = 0
                            if false {
                                _ = len(v2)
                            } else {
                                fmt.Println(v27)
                                switch (v28 * fi29) {
                                case 9223372036854775807:
                                    if len(v6) > 0 {
                                        v6[0] = "foo bar"
                                    }
                                case 470:
                                    var v52 uint8 = 1
                                    _ = v52
                                }
                            }
                            _ = fi51
                        }
                        if (2 == 0) {
                            v2[(v5 + (v5 + ("hello" + v5)))] = false
                            _ = len(v2)
                        } else {
                            if ((fi29 ^ 127) >= -907) {
                                fmt.Println(v3)
                                v53 := GenS5{F0: -9223372036854775808, F1: 1, F2: (0 < 0)}
                                _ = v53
                            } else {
                                for fi54 := 0; fi54 < 11; fi54++ {
                                    v55 := GenS1{F0: ("" + v5)}
                                    var v56 uint = 2
                                    if len(v6) >= 2 {
                                        v6 = v6[1:len(v6)]
                                    }
                                    _ = fi54
                                    _ = v55
                                    _ = v56
                                }
                                if true {
                                    v57 := v2[("hello" + v5)]
                                    _ = v57
                                } else {
                                    var v58 string = "hello"
                                    _ = v58
                                }
                                var v59 int64 = 9223372036854775807
                                _ = v59
                            }
                            var v60 int8 = v1
                            _ = v60
                        }
                        v61 := GenS4{F0: (-(fi29))}
                        _ = v61
                    } else {
                        v5 = ""
                        if (42 > 1) {
                            if false {
                                fmt.Println(v3, v28)
                                var v62 uint32 = 65535
                                _ = v62
                            }
                        } else {
                            _ = len(v6)
                            fmt.Println(v28)
                        }
                        v5 = "foo bar"
                    }
                    v63 := v2["\n\t"]
                    _ = fi29
                    _ = v63
                }
                _ = v27
                _ = v28
            }
        }
    }
    if (-2.25 >= 1e10) {
        if true {
            fmt.Println(v4, v4, v1)
            switch ("hello" + "") {
            case "\n\t":
                v64 := GenS2{F0: 1, F1: 934, F2: 64}
                _ = v64
            case "hello":
                v5 = v5
            case "foo bar":
                if true {
                    if true {
                        if false {
                            v65 := GenS3{F0: 18446744073709551615, F1: v1}
                            _ = v65
                        } else {
                            switch v5 {
                            case "a":
                                var v66 uint16 = 0
                                _ = v66
                            case "hello":
                                for fi67 := 0; fi67 < 16; fi67++ {
                                    v5 = "foo bar"
                                    _ = fi67
                                }
                            }
                            var v68 bool = true
                            _ = len(v2)
                            _ = v68
                        }
                    } else {
                        for fi69 := 0; fi69 < 11; fi69++ {
                            var v70 float32 = v4
                            v6 = append(v6, "\n\t")
                            v5 = "\n\t"
                            _ = fi69
                            _ = v70
                        }
                    }
                    for fi71 := 0; fi71 < 4; fi71++ {
                        for fi72 := 0; fi72 < 3; fi72++ {
                            fmt.Println(v5, v1)
                            switch (^(fi71)) {
                            case -79:
                                v73 := GenS6{F0: 472}
                                _ = v73
                            }
                            _ = fi72
                        }
                        v74 := GenS7{F0: -1}
                        v75 := GenS5{F0: 42, F1: -32768, F2: (v4 > -2.25)}
                        _ = fi71
                        _ = v74
                        _ = v75
                    }
                } else {
                    var v76 uint8 = 255
                    fmt.Println(v3, v5, v5)
                    _ = v76
                }
            default:
                var v77 uint8 = 42
                _ = v77
            }
            var v78 bool = (true || false)
            _ = v78
        }
    }
    delete(v2, ((((v5 + ((v5 + v5) + "\n\t")) + "") + "a") + ((v5 + "a") + v5)))
    var v79 uint8 = 0
    switch v5 {
    case "\n\t":
        var v80 int16 = 1000
        _ = v80
    case "hello":
        for fi81 := 0; fi81 < 14; fi81++ {
            var v82 string = ("a" + "a")
            _ = fi81
            _ = v82
        }
    }
    v83 := v2["\n\t"]
    v84 := GenS2{F0: 511, F1: 940, F2: -299}
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    for fi85 := 0; fi85 < 16; fi85++ {
        v2[((v5 + v5) + ("hello" + ((v5 + "hello") + (v5 + "foo bar"))))] = false
        _ = fi85
    }
    if v83 {
        v5 = v5
    }
    fmt.Println(v3, v79)
    fmt.Println(v83, v83)
    var v86 uint = 4294967295
    for fi87 := 0; fi87 < 3; fi87++ {
        var v88 uint16 = genFn0(805, 707)
        _ = fi87
        _ = v88
    }
    var v89 bool = v83
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v79:", v79)
    fmt.Println("v83:", v83)
    fmt.Println("v86:", v86)
    fmt.Println("v89:", v89)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v79
    _ = v83
    _ = v84
    _ = v86
    _ = v89
}
