package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int8; F1 uint; F2 int32 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint; F1 uint16; F2 int32 }
type GenS6 struct { F0 int64; F1 bool }
type GenS7 struct { F0 string; F1 uint; F2 uint8 }
type GenS8 struct { F0 uint; F1 int64; F2 int16 }
type GenS9 struct { F0 uint8; F1 int32 }
type GenS10 struct { F0 int8; F1 int16; F2 uint8 }
func genFn0() uint16 {
    if true {
        if (!true) {
            v5 := []int{127}
            _ = v5
        } else {
            v6 := GenS1{F0: -830}
            var v7 int64 = 9223372036854775807
            _ = v6
            _ = v7
        }
    } else {
        for fi8 := 0; fi8 < 10; fi8++ {
            var v9 int = fi8
            v10 := GenS1{F0: 0}
            for fi11 := 0; fi11 < 3; fi11++ {
                v9 += (v9 / (((fi8 ^ (fi8 << ((42) & 63)))) | 1))
                var v12 int = 0
                _ = fi11
                _ = v12
            }
            _ = fi8
            _ = v10
            _ = v9
        }
    }
    var v13 float64 = ((1e10 + 1.5) - float64(1e10))
    _ = v13
    return 449
}
func genFn1(p0 bool, p1 uint32, p2 int8) int8 {
    for fi14 := 0; fi14 < 4; fi14++ {
        v15 := []int{fi14, -955}
        v16 := GenS1{F0: 425}
        var v17 float32 = (0.1 + 0.1)
        _ = fi14
        _ = v15
        _ = v16
        _ = v17
    }
    p1 -= (p1 << ((65535) & 63))
    for fi18 := 0; fi18 < 11; fi18++ {
        var v19 int = int(p2)
        var v20 int32 = 2147483647
        v19 = (fi18 / ((v19) | 1))
        _ = fi18
        _ = v19
        _ = v20
    }
    _ = p0
    _ = p1
    _ = p2
    return (p2 & (p2 * (-(p2))))
}

func main() {
    v1 := []uint32{174, 198, 2147483648}
    var v2 int8 = 100
    var v3 uint64 = 42
    v4 := make(map[uint32]uint32)
    if (1 == 42) {
        v21 := GenS2{F0: v2, F1: 65535, F2: -500}
        for fi22 := 0; fi22 < 13; fi22++ {
            fmt.Println(fi22)
            _ = fi22
        }
        if true {
            v2 = v2
        }
        _ = v21
    } else {
        var v23 float64 = -2.25
        for fi24 := 0; fi24 < 16; fi24++ {
            for fi25 := 0; fi25 < 5; fi25++ {
                v2 += (^(v2))
                v26 := v4[42]
                _ = fi25
                _ = v26
            }
            v27 := GenS3{F0: 2}
            for fi28 := 0; fi28 < 20; fi28++ {
                v29 := GenS4{F0: 42}
                fmt.Println(v23, fi28)
                v23 = v23
                _ = fi28
                _ = v29
            }
            _ = fi24
            _ = v27
        }
        fmt.Println(v23)
        _ = v23
    }
    switch "foo bar" {
    case "":
        if (-1 < 42) {
            if false {
                v30 := GenS3{F0: 42}
                _ = v30
            } else {
                switch 1 {
                case -337:
                    switch -2147483648 {
                    case -2147483648:
                        _ = len(v4)
                    case -9223372036854775808:
                        for fi31 := 0; fi31 < 15; fi31++ {
                            switch ((("foo bar" + ("a" + "a")) + "hello") + "a") {
                            case "a":
                                for fi32 := 0; fi32 < 9; fi32++ {
                                    if (false && (false && (false && (int32(v2) != 42)))) {
                                        delete(v4, 958)
                                        if false {
                                            v2 = -67
                                            var v33 uint64 = (v3 * v3)
                                            _ = v33
                                        }
                                        var v34 uint8 = 220
                                        _ = v34
                                    }
                                    _ = fi32
                                }
                            case "":
                                v35 := GenS5{F0: 50, F1: 42, F2: 42}
                                _ = v35
                            default:
                                var v36 int16 = -32768
                                _ = v36
                            }
                            if true {
                                v37 := GenS3{F0: 2147483648}
                                for fi38 := 0; fi38 < 20; fi38++ {
                                    if (443 > 383) {
                                        if false {
                                            delete(v4, 42)
                                            if false {
                                                switch fi38 {
                                                case 4294967295:
                                                    if true {
                                                        fmt.Println(fi38, fi38, v2)
                                                        if len(v1) >= 2 {
                                                            v1 = v1[1:len(v1)]
                                                        }
                                                    } else {
                                                        var v39 string = "\n\t"
                                                        _ = v39
                                                    }
                                                }
                                                if (false && false) {
                                                    var v40 uint16 = 607
                                                    v41 := GenS4{F0: 491}
                                                    _ = len(v1)
                                                    fmt.Println(v40, v2, v2)
                                                    _ = len(v1)
                                                    _ = v40
                                                    _ = v41
                                                } else {
                                                    switch fi38 {
                                                    case 4294967295:
                                                        var v42 uint8 = 0
                                                        _ = v42
                                                    default:
                                                        switch 42 {
                                                        case 32768:
                                                            fmt.Println(v2, v2)
                                                        case 1:
                                                            var v43 int = (fi38 << ((487) & 63))
                                                            _ = v43
                                                        case 127:
                                                            v2 += v2
                                                        default:
                                                            v3 = (^(v3))
                                                        }
                                                    }
                                                    var v44 uint8 = 1
                                                    var v45 bool = (51 != 0)
                                                    _ = v44
                                                    _ = v45
                                                }
                                                var v46 uint16 = genFn0()
                                                _ = v46
                                            }
                                        }
                                        v47 := GenS4{F0: 1}
                                        _ = v47
                                    } else {
                                        if (true || false) {
                                            v3 = (v3 & (v3 + (v3 >> ((18446744073709551615) & 63))))
                                            if (-703 < 1) {
                                                v48 := v4[275]
                                                fmt.Println(v3, fi38, fi31)
                                                _ = v48
                                            } else {
                                                switch "" {
                                                case "":
                                                    var v49 float64 = (((-2.25 - (1e10 / (-2.25 * (1e10 / 3.14159)))) / 0.1) * 0.1)
                                                    _ = v49
                                                case "a":
                                                    var v50 int64 = 1
                                                    _ = v50
                                                }
                                                _ = len(v1)
                                                v51 := GenS6{F0: 1, F1: (("foo bar" < "foo bar") && (1.5 >= (1e10 / 1e10)))}
                                                fmt.Println(fi38, fi38)
                                                _ = v51
                                            }
                                        } else {
                                            v52 := GenS3{F0: 2147483648}
                                            _ = v52
                                        }
                                        v2 = (-(v2))
                                        v3 -= v3
                                    }
                                    fmt.Println(fi38)
                                    var v53 int = (-(fi38))
                                    _ = fi38
                                    _ = v53
                                }
                                var v54 int8 = genFn1(false, 2147483648, (v2 >> ((4294967295) & 63)))
                                _ = v37
                                _ = v54
                            } else {
                                fmt.Println(v2, v3)
                                var v55 uint8 = 128
                                var v56 bool = true
                                _ = v55
                                _ = v56
                            }
                            _ = fi31
                        }
                    default:
                        fmt.Println(v2, v3)
                    }
                case -158:
                    if (!(("hello" + ("" + ("\n\t" + (("foo bar" + (("a" + (("foo bar" + "hello") + (("a" + (("hello" + ("foo bar" + "")) + "hello")) + "hello"))) + "hello")) + "foo bar")))) == "a")) {
                        v57 := GenS3{F0: 290}
                        v58 := GenS7{F0: "a", F1: 2, F2: 128}
                        _ = v57
                        _ = v58
                    }
                }
            }
            switch 2147483647 {
            case -257:
                var v59 int16 = 42
                _ = v59
            case 42:
                if (false || ((4294967295 >= 2147483647) || (1 == 9223372036854775807))) {
                    if ((0 <= 143) && (v3 > (v3 % (((v3 >> ((287) & 63))) | 1)))) {
                        switch "hello" {
                        case "a":
                            v3 = v3
                        case "hello":
                            var v60 uint16 = 251
                            _ = v60
                        }
                        if (false || (false || false)) {
                            v4[0] = 65535
                        } else {
                            v3 -= (^(v3))
                        }
                        var v61 float64 = -2.25
                        _ = v61
                    }
                    v3 = 18446744073709551615
                    v3 += (v3 << ((183) & 63))
                } else {
                    var v62 int8 = genFn1((!(!true)), 723, genFn1((277 != -32768), 915, 1))
                    var v63 int64 = -1
                    v62 += (v2 / ((v62) | 1))
                    _ = v62
                    _ = v63
                }
            }
            v64 := GenS5{F0: 18446744073709551615, F1: 530, F2: 702}
            _ = v64
        }
    }
    if (-933 >= 0) {
        fmt.Println(v3, v2)
        for fi65 := 0; fi65 < 3; fi65++ {
            if (65535 != 0) {
                v66 := GenS5{F0: 0, F1: 402, F2: 1}
                v67 := GenS5{F0: 367, F1: 0, F2: 2147483647}
                _ = v66
                _ = v67
            } else {
                var v68 string = ("foo bar" + "")
                v68 = v68
                switch v68 {
                case "\n\t":
                    var v69 float32 = float32(float64(-2.25))
                    _ = v69
                }
                _ = v68
            }
            _ = len(v4)
            _ = fi65
        }
    }
    v2 += (v2 % ((-128) | 1))
    if (!((true && true) && (42 > 2))) {
        if true {
            var v70 int32 = 324
            v71 := GenS8{F0: 42, F1: 42, F2: 238}
            if false {
                var v72 int64 = 9223372036854775807
                _ = v72
            } else {
                v73 := GenS3{F0: 1}
                var v74 int = 4294967295
                _ = v73
                _ = v74
            }
            _ = v70
            _ = v71
        }
        v3 = 42
        var v75 int = -1
        _ = v75
    } else {
        var v76 int = 0
        v77 := GenS9{F0: 0, F1: 454}
        _ = v76
        _ = v77
    }
    v3 = 4294967295
    fmt.Println(v2)
    var v78 float64 = 0.1
    if true {
        var v79 uint16 = 432
        var v80 uint32 = 84
        _ = v79
        _ = v80
    }
    for fi81 := 0; fi81 < 3; fi81++ {
        var v82 int32 = -2147483648
        v83 := GenS10{F0: (v2 >> ((453) & 63)), F1: 32767, F2: 127}
        _ = fi81
        _ = v82
        _ = v83
    }
    v84 := GenS4{F0: 317}
    v78 -= ((((float64(-2.25) - v78) + v78) - (v78 - float64(float32(3.14159)))) / 1.0)
    v85 := GenS9{F0: 0, F1: -74}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v78:", v78)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v78
    _ = v84
    _ = v85
}
