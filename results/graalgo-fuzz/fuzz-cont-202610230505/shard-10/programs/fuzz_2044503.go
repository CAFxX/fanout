package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 string; F1 uint8 }
type GenS4 struct { F0 int64; F1 uint8 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int64; F1 uint32 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 bool; F1 string }
type GenS9 struct { F0 uint8; F1 int8 }
func genFn0() int32 {
    var v6 uint64 = 4294967295
    var v7 int16 = 32767
    _ = v6
    _ = v7
    return 1
}
func genFn1(p0 string, p1 int, p2 string) int16 {
    if false {
        v8 := GenS1{F0: 300}
        var v9 int16 = 112
        fmt.Println(v9)
        _ = v8
        _ = v9
    } else {
        if true {
            p0 = p2
            fmt.Println(p0, p0)
            if (42 != -45) {
                v10 := []int{(^(p1)), (p1 / (((p1 % ((p1) | 1))) | 1)), 42}
                v11 := GenS1{F0: 65535}
                if len(v10) > 0 {
                    v10[0] = (-(p1))
                }
                _ = v10
                _ = v11
            } else {
                fmt.Println(p2, p0)
            }
        }
    }
    var v12 float32 = (3.14159 * ((-2.25 * 3.14159) + 0.1))
    p0 = ((p2 + p2) + "hello")
    switch (p1 + (p1 >> ((2) & 63))) {
    case 128:
        var v13 int32 = 1
        _ = v13
    case -761:
        v14 := []int{(p1 - 383)}
        _ = v14
    case 42:
        switch (p1 % ((p1) | 1)) {
        case 32768:
            p2 = "\n\t"
        case 42:
            for fi15 := 0; fi15 < 4; fi15++ {
                for fi16 := 0; fi16 < 8; fi16++ {
                    var v17 uint16 = 65535
                    if (!false) {
                        v18 := []int{fi15}
                        _ = v18
                    }
                    v19 := GenS2{F0: -1}
                    _ = fi16
                    _ = v17
                    _ = v19
                }
                v20 := GenS1{F0: 65535}
                switch 4294967295 {
                case 32767:
                    if (580 >= 0) {
                        v12 = v12
                        fmt.Println(v12, p0, p1)
                        v21 := []int{1}
                        _ = v21
                    }
                case -1:
                    fmt.Println(p2, p1, p1)
                case 127:
                    p2 = (p2 + p2)
                }
                _ = fi15
                _ = v20
            }
        case -900:
            p1 -= p1
        default:
            var v22 uint16 = 776
            _ = v22
        }
    default:
        v23 := GenS3{F0: "hello", F1: 255}
        _ = v23
    }
    p2 = ""
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return 42
}

func main() {
    var v1 float32 = 1e10
    v2 := make(map[int]uint32)
    v3 := make(map[int]int)
    var v4 int16 = 0
    v5 := []uint32{2147483648}
    v1 -= v1
    v1 += (float32(0.0) - v1)
    var v24 int32 = -681
    if false {
        v25 := GenS4{F0: -1, F1: 13}
        v5 = append(v5, 0)
        v24 -= 2147483647
        _ = v25
    }
    for fi26 := 0; fi26 < 19; fi26++ {
        if (!(false && (95 != 42))) {
            if (!true) {
                var v27 bool = false
                v28 := GenS4{F0: 42, F1: 144}
                _ = v27
                _ = v28
            } else {
                for fi29 := 0; fi29 < 15; fi29++ {
                    if len(v5) > 0 {
                        v5[0] = 868
                    }
                    var v30 int16 = 1
                    v24 += (v24 >> ((42) & 63))
                    _ = fi29
                    _ = v30
                }
            }
            if false {
                delete(v2, fi26)
            } else {
                v4 += 1000
            }
        } else {
            switch fi26 {
            case 32767:
                v1 = float32(float64(0.0))
            case -1:
                if (-2.25 > (0.0 * -2.25)) {
                    v31 := GenS5{F0: -631}
                    v1 = v1
                    for fi32 := 0; fi32 < 3; fi32++ {
                        v33 := GenS1{F0: 0}
                        if len(v5) > 0 {
                            v5[0] = 483
                        }
                        if ((2 >= 0) || ((1.5 / v1) < ((0.0 + (-2.25 + 0.1)) / v1))) {
                            v4 = (v4 * 848)
                        }
                        _ = fi32
                        _ = v33
                    }
                    _ = v31
                } else {
                    fmt.Println(v24, v4, v24)
                    switch fi26 {
                    case 128:
                        delete(v3, fi26)
                    case -2147483648:
                        var v34 uint16 = 65535
                        _ = v34
                    default:
                        v24 += 1
                    }
                }
            default:
                if len(v5) > 0 {
                    v5[0] = 502
                }
            }
        }
        for fi35 := 0; fi35 < 13; fi35++ {
            var v36 uint8 = 137
            _ = fi35
            _ = v36
        }
        switch "hello" {
        case "a":
            var v37 int8 = 100
            _ = v37
        case "\n\t":
            var v38 bool = (3.14159 < ((0.1 + 0.1) - 1e10))
            _ = v38
        case "hello":
            switch "\n\t" {
            case "a":
                v24 -= (-(v24))
            }
        }
        _ = fi26
    }
    var v39 int64 = 616
    var v40 uint32 = 98
    v41 := GenS6{F0: 42, F1: 42}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    switch 4294967295 {
    case 9223372036854775807:
        v42 := GenS3{F0: "hello", F1: 0}
        _ = v42
    case 1:
        switch (("" + "a") + "foo bar") {
        case "a":
            switch "\n\t" {
            case "hello":
                v24 = 1
            case "\n\t":
                fmt.Println(v39, v1, v40)
            default:
                if true {
                    var v43 string = (("" + "a") + "foo bar")
                    v4 -= (v4 % (((^(v4))) | 1))
                    _ = v43
                } else {
                    if true {
                        if (true && false) {
                            v40 = (v40 >> ((4294967295) & 63))
                        }
                    }
                }
            }
        case "hello":
            v44 := GenS4{F0: (^(v39)), F1: 1}
            _ = v44
        case "foo bar":
            if false {
                if false {
                    if (0 < 238) {
                        var v45 int32 = 0
                        v46 := GenS7{F0: false}
                        fmt.Println(v4)
                        _ = v45
                        _ = v46
                    } else {
                        if (("hello" + "a") != "hello") {
                            v39 = (v39 % ((-9223372036854775808) | 1))
                            if true {
                                if true {
                                    v47 := GenS7{F0: (0 >= 739)}
                                    fmt.Println(v24, v4, v1)
                                    _ = v47
                                } else {
                                    delete(v3, 106)
                                }
                                var v48 int64 = (v39 + 9223372036854775807)
                                _ = v48
                            } else {
                                fmt.Println(v4, v40)
                            }
                            var v49 int = 32768
                            _ = v49
                        } else {
                            var v50 uint = 65535
                            _ = v50
                        }
                        v1 += 1e10
                        var v51 uint8 = 118
                        _ = v51
                    }
                }
                _ = len(v5)
                v5 = append(v5, v40)
            } else {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                v52 := GenS8{F0: true, F1: "foo bar"}
                _ = v52
            }
        }
    case 4294967295:
        if true {
            switch ("" + "\n\t") {
            case "":
                v53 := v3[2147483647]
                _ = v53
            case "hello":
                for fi54 := 0; fi54 < 19; fi54++ {
                    if (!(!(genFn0() >= v24))) {
                        v24 += (^(v24))
                        v55 := v3[(fi54 + 0)]
                        v56 := GenS9{F0: 31, F1: 100}
                        _ = v55
                        _ = v56
                    }
                    _ = fi54
                }
            default:
                var v57 uint = 9
                _ = v57
            }
            if (v4 > v4) {
                _ = len(v5)
                var v58 uint64 = 18446744073709551615
                var v59 uint8 = 0
                _ = v58
                _ = v59
            }
        }
    default:
        fmt.Println(v39)
    }
    for fi60 := 0; fi60 < 15; fi60++ {
        for fi61 := 0; fi61 < 3; fi61++ {
            v40 += v40
            var v62 int32 = (v24 << ((2) & 63))
            _ = fi61
            _ = v62
        }
        var v63 uint32 = 42
        _ = fi60
        _ = v63
    }
    var v64 uint64 = 2
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v64:", v64)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v64
}
