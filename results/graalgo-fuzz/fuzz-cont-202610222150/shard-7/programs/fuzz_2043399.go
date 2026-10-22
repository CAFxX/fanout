package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8; F2 int }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int8; F1 int32; F2 uint8 }
type GenS4 struct { F0 uint64; F1 uint; F2 int8 }
type GenS5 struct { F0 string; F1 int; F2 uint }
func genFn0() uint {
    for fi6 := 0; fi6 < 5; fi6++ {
        var v7 float64 = (-2.25 / (3.14159 - -2.25))
        _ = fi6
        _ = v7
    }
    if false {
        v8 := []int{}
        _ = len(v8)
        _ = v8
    }
    var v9 uint64 = 1
    _ = v9
    return 1
}

func main() {
    var v1 int = -9223372036854775808
    v2 := []uint32{}
    v3 := make(map[int]string)
    var v4 uint8 = 127
    v5 := []string{"a", "a"}
    if true {
        var v10 float32 = (1.5 / 0.1)
        _ = v10
    } else {
        var v11 int64 = -852
        _ = v11
    }
    var v12 int = (v1 << ((12) & 63))
    _ = len(v5)
    var v13 int = v1
    if false {
        switch ("foo bar" + "") {
        case "":
            fmt.Println(v12, v13)
        case "foo bar":
            for fi14 := 0; fi14 < 5; fi14++ {
                v13 += -240
                _ = fi14
            }
        case "hello":
            v2 = append(v2, 4294967295)
        default:
            var v15 uint8 = v4
            _ = v15
        }
        v16 := GenS1{F0: 825, F1: 4, F2: (v13 ^ -868)}
        if (42 > 2) {
            if (false || (350 >= 2147483648)) {
                switch (("a" + ("\n\t" + "a")) + "hello") {
                case "":
                    var v17 uint64 = 1
                    _ = v17
                case "foo bar":
                    v18 := v3[v13]
                    _ = v18
                case "hello":
                    v4 = (v4 << ((0) & 63))
                default:
                    for fi19 := 0; fi19 < 6; fi19++ {
                        for fi20 := 0; fi20 < 16; fi20++ {
                            var v21 bool = (!(true && true))
                            var v22 uint32 = 2
                            _ = fi20
                            _ = v21
                            _ = v22
                        }
                        _ = fi19
                    }
                }
            }
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            var v23 int32 = 2147483647
            _ = v23
        }
        _ = v16
    }
    v4 += (^(v4))
    v4 = (v4 & (v4 & v4))
    var v24 string = ("foo bar" + ("" + "\n\t"))
    switch (v13 << ((0) & 63)) {
    case -409:
        if true {
            v24 = "hello"
            for fi25 := 0; fi25 < 19; fi25++ {
                delete(v3, (fi25 ^ v1))
                if true {
                    v2 = append(v2, 65535)
                }
                _ = fi25
            }
        } else {
            var v26 uint32 = 1
            v12 -= v1
            var v27 int16 = 42
            _ = v26
            _ = v27
        }
    case 32768:
        v5 = append(v5, v24)
    case 128:
        var v28 int16 = -32768
        _ = v28
    default:
        for fi29 := 0; fi29 < 1; fi29++ {
            for fi30 := 0; fi30 < 3; fi30++ {
                var v31 int = v1
                _ = fi30
                _ = v31
            }
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            _ = fi29
        }
    }
    if true {
        v3[0] = "foo bar"
    } else {
        if (0 <= 42) {
            v32 := GenS1{F0: 2, F1: -128, F2: v1}
            switch v1 {
            case 127:
                v33 := GenS1{F0: 2, F1: 1, F2: v1}
                _ = v33
            case 9223372036854775807:
                for fi34 := 0; fi34 < 13; fi34++ {
                    var v35 uint = 18446744073709551615
                    var v36 int16 = -1
                    var v37 uint = 240
                    _ = fi34
                    _ = v35
                    _ = v36
                    _ = v37
                }
            case -2147483648:
                var v38 int = v13
                _ = v38
            default:
                switch v1 {
                case 0:
                    v39 := GenS1{F0: 255, F1: 44, F2: 32767}
                    _ = v39
                case 128:
                    for fi40 := 0; fi40 < 16; fi40++ {
                        for fi41 := 0; fi41 < 2; fi41++ {
                            v42 := GenS1{F0: 255, F1: 127, F2: (v12 - fi40)}
                            if len(v5) > 0 {
                                v5[0] = v24
                            }
                            v13 = v12
                            switch v24 {
                            case "a":
                                if true {
                                    v1 += -9223372036854775808
                                } else {
                                    if (!true) {
                                        if true {
                                            v43 := GenS2{F0: 18446744073709551615}
                                            _ = v43
                                        }
                                        v13 = 492
                                    } else {
                                        v3[v1] = ("" + v24)
                                        _ = len(v3)
                                    }
                                }
                            default:
                                if (1000 == 704) {
                                    v1 = v13
                                }
                            }
                            var v44 int16 = -336
                            _ = fi41
                            _ = v42
                            _ = v44
                        }
                        for fi45 := 0; fi45 < 13; fi45++ {
                            v12 = (^(fi40))
                            var v46 int = (fi45 << ((genFn0()) & 63))
                            v24 = v24
                            _ = fi45
                            _ = v46
                        }
                        var v47 int = fi40
                        _ = fi40
                        _ = v47
                    }
                case -593:
                    v1 -= 9223372036854775807
                default:
                    v12 = (v13 - v12)
                }
            }
            _ = v32
        } else {
            _ = len(v3)
            for fi48 := 0; fi48 < 6; fi48++ {
                var v49 uint = genFn0()
                _ = fi48
                _ = v49
            }
        }
    }
    var v50 int32 = 679
    switch v1 {
    case -9223372036854775808:
        var v51 uint16 = 32767
        _ = v51
    case -1:
        v50 += v50
    case 232:
        var v52 uint64 = 2
        _ = v52
    default:
        if true {
            v53 := GenS3{F0: 127, F1: 42, F2: (v4 - v4)}
            if ((203 >= uint16(v4)) || false) {
                v54 := GenS2{F0: 0}
                switch 32767 {
                case 385:
                    for fi55 := 0; fi55 < 9; fi55++ {
                        for fi56 := 0; fi56 < 19; fi56++ {
                            var v57 int = (v12 | (v12 % (((fi55 % ((v1) | 1))) | 1)))
                            var v58 int64 = int64(v50)
                            delete(v3, (-(v13)))
                            _ = fi56
                            _ = v57
                            _ = v58
                        }
                        var v59 uint32 = 956
                        var v60 int = -2147483648
                        _ = fi55
                        _ = v59
                        _ = v60
                    }
                default:
                    fmt.Println(v50, v12, v12)
                }
                _ = v54
            }
            _ = v53
        }
    }
    var v61 int8 = -128
    var v62 float64 = float64(1e10)
    fmt.Println(v24, v62, v1)
    for fi63 := 0; fi63 < 4; fi63++ {
        v64 := GenS4{F0: 746, F1: genFn0(), F2: (-(v61))}
        v65 := GenS5{F0: "\n\t", F1: 983, F2: 42}
        v66 := GenS5{F0: v24, F1: (v1 | v1), F2: 255}
        _ = fi63
        _ = v64
        _ = v65
        _ = v66
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v50:", v50)
    fmt.Println("v61:", v61)
    fmt.Println("v62:", v62)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v61
    _ = v62
}
