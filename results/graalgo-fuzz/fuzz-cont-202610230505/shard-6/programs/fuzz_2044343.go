package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64 }
type GenS2 struct { F0 string; F1 string }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint32; F1 int32; F2 uint64 }
type GenS5 struct { F0 uint8; F1 uint32 }
type GenS6 struct { F0 uint32; F1 int64; F2 uint16 }
type GenS7 struct { F0 int64; F1 int8; F2 int8 }

func main() {
    var v1 float32 = 0.1
    var v2 string = ("foo bar" + "hello")
    v3 := make(map[string]uint32)
    var v4 int8 = -1
    var v5 int = -2147483648
    var v6 int16 = 0
    for fi7 := 0; fi7 < 11; fi7++ {
        v8 := GenS1{F0: -2147483648, F1: -418}
        for fi9 := 0; fi9 < 15; fi9++ {
            v10 := GenS1{F0: 489, F1: 1}
            fmt.Println(v6, fi9)
            _ = fi9
            _ = v10
        }
        _ = fi7
        _ = v8
    }
    for fi11 := 0; fi11 < 7; fi11++ {
        v12 := GenS2{F0: (("a" + (v2 + v2)) + v2), F1: "\n\t"}
        _ = fi11
        _ = v12
    }
    v3["\n\t"] = 329
    v6 -= v6
    switch (("hello" + ((((("" + (v2 + v2)) + (v2 + v2)) + (("hello" + "") + v2)) + "hello") + "foo bar")) + "") {
    case "foo bar":
        for fi13 := 0; fi13 < 19; fi13++ {
            if false {
                delete(v3, v2)
                for fi14 := 0; fi14 < 1; fi14++ {
                    v15 := GenS3{F0: (fi14 & (^(fi13)))}
                    if false {
                        switch "a" {
                        case "":
                            v16 := GenS4{F0: 684, F1: 2147483647, F2: 4294967295}
                            _ = v16
                        case "hello":
                            v2 = v2
                        default:
                            v3["foo bar"] = 0
                        }
                        v5 = fi13
                    }
                    if true {
                        var v17 uint = 92
                        v18 := []int{1, fi14, 127, 9223372036854775807}
                        if (42 == 42) {
                            v6 = (^(v6))
                        }
                        v5 += v5
                        var v19 uint = v17
                        _ = v17
                        _ = v18
                        _ = v19
                    }
                    switch v5 {
                    case 4294967295:
                        switch 128 {
                        case -295:
                            var v20 int64 = 42
                            _ = v20
                        case -1:
                            v4 += 127
                            var v21 int8 = 1
                            _ = v21
                        case 2147483647:
                            switch "foo bar" {
                            case "":
                                if true {
                                    v22 := GenS2{F0: "", F1: (v2 + "")}
                                    v1 += ((0.0 + float32((1.5 / 1e10))) + 1e10)
                                    v23 := []int{128, fi13, 9223372036854775807, (fi14 & fi13)}
                                    _ = v22
                                    _ = v23
                                } else {
                                    if true {
                                        var v24 uint32 = 2147483647
                                        var v25 uint8 = 183
                                        _ = v24
                                        _ = v25
                                    } else {
                                        var v26 int = (fi13 / ((v5) | 1))
                                        v27 := GenS5{F0: 185, F1: 0}
                                        if true {
                                            var v28 int16 = int16(fi13)
                                            v5 = v26
                                            fmt.Println(v6, v1, v2)
                                            _ = v28
                                        }
                                        if (!(false && (!true))) {
                                            fmt.Println(fi13)
                                            v29 := GenS6{F0: 2147483648, F1: -1, F2: 295}
                                            _ = v29
                                        } else {
                                            var v30 uint = 395
                                            var v31 int = 32767
                                            _ = v30
                                            _ = v31
                                        }
                                        v32 := GenS1{F0: 737, F1: 0}
                                        if ("" >= v2) {
                                            var v33 int = -745
                                            _ = v33
                                        }
                                        var v34 int16 = (v6 - v6)
                                        _ = v26
                                        _ = v27
                                        _ = v32
                                        _ = v34
                                    }
                                    var v35 int8 = 1
                                    v36 := GenS7{F0: -9223372036854775808, F1: v4, F2: 103}
                                    _ = v35
                                    _ = v36
                                }
                                fmt.Println(fi13)
                            case "hello":
                                v37 := GenS1{F0: -940, F1: -1}
                                _ = v37
                            default:
                                _ = len(v3)
                            }
                        }
                    default:
                        if false {
                            if (("hello" + (v2 + "")) != v2) {
                                if true {
                                    var v38 bool = false
                                    _ = v38
                                }
                                v4 = (v4 % ((int8(v6)) | 1))
                                var v39 string = "hello"
                                switch (v5 - fi14) {
                                case 0:
                                    if false {
                                        var v40 uint64 = 0
                                        _ = v40
                                    }
                                case -1:
                                    var v41 uint16 = 1
                                    _ = v41
                                }
                                v42 := GenS7{F0: 9223372036854775807, F1: v4, F2: 100}
                                _ = v39
                                _ = v42
                            }
                        } else {
                            v43 := GenS7{F0: 1, F1: v4, F2: v4}
                            v4 = (v4 * v4)
                            switch "foo bar" {
                            case "hello":
                                var v44 float64 = 0.0
                                _ = v44
                            default:
                                var v45 int16 = (v6 >> ((446) & 63))
                                _ = v45
                            }
                            _ = v43
                        }
                    }
                    _ = fi14
                    _ = v15
                }
                fmt.Println(v6, v6, v2)
            }
            _ = fi13
        }
    case "a":
        switch v5 {
        case -1:
            for fi46 := 0; fi46 < 16; fi46++ {
                v5 = fi46
                _ = fi46
            }
        default:
            v5 += v5
        }
    case "":
        v6 = 693
    default:
        switch (v5 / ((int(v6)) | 1)) {
        case 9223372036854775807:
            for fi47 := 0; fi47 < 12; fi47++ {
                v48 := GenS1{F0: -394, F1: 0}
                for fi49 := 0; fi49 < 10; fi49++ {
                    v50 := []int{(v5 | fi47), (v5 * fi49), (^(v5)), fi49}
                    v5 -= (fi47 ^ (fi49 & (fi47 | 32767)))
                    v1 = v1
                    _ = fi49
                    _ = v50
                }
                fmt.Println(v4)
                _ = fi47
                _ = v48
            }
        case -9223372036854775808:
            delete(v3, v2)
        case 32767:
            for fi51 := 0; fi51 < 5; fi51++ {
                v52 := GenS5{F0: 0, F1: 1}
                _ = fi51
                _ = v52
            }
        }
    }
    if (82 < 385) {
        if false {
            if true {
                v1 += v1
            }
            var v53 int = 32767
            var v54 uint32 = 430
            _ = v53
            _ = v54
        } else {
            var v55 int64 = 305
            var v56 bool = (!false)
            if (2147483648 != 4294967295) {
                var v57 uint32 = 735
                var v58 int8 = -59
                v4 += v58
                _ = v57
                _ = v58
            }
            _ = v55
            _ = v56
        }
    } else {
        fmt.Println(v1)
        var v59 bool = true
        _ = v59
    }
    switch 32768 {
    case -255:
        v60 := []int{(-(v5)), (v5 * (v5 | 1))}
        _ = v60
    case -1:
        fmt.Println(v6, v2, v4)
    case -289:
        switch (v5 << ((0) & 63)) {
        case 32768:
            for fi61 := 0; fi61 < 3; fi61++ {
                var v62 bool = false
                _ = fi61
                _ = v62
            }
        }
    default:
        v1 = 0.0
    }
    v5 = v5
    v63 := GenS2{F0: v2, F1: "foo bar"}
    v64 := GenS7{F0: -854, F1: 33, F2: v4}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v63
    _ = v64
}
