package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8 }
type GenS2 struct { F0 int16; F1 int16 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint8; F1 uint16; F2 uint64 }
type GenS5 struct { F0 uint64; F1 int8; F2 string }
type GenS6 struct { F0 uint; F1 int }

func main() {
    v1 := []uint32{0, 2, 2147483647}
    var v2 int32 = 2147483647
    v3 := make(map[uint32]int)
    var v4 int8 = int8(v2)
    _ = len(v3)
    v4 -= -128
    v5 := GenS1{F0: "foo bar", F1: (v4 + (-(v4)))}
    v2 = -502
    v6 := GenS1{F0: "hello", F1: -2}
    var v7 bool = true
    fmt.Println(v2)
    v8 := GenS1{F0: "hello", F1: v4}
    fmt.Println(v4, v7)
    if v7 {
        fmt.Println(v2, v2, v2)
        v7 = (true && (0.1 != 1.5))
    } else {
        fmt.Println(v4, v7, v2)
    }
    switch 32768 {
    case 1:
        var v9 float64 = ((0.0 / float64(1.5)) * 0.1)
        _ = v9
    default:
        v7 = (v7 || v7)
    }
    var v10 float32 = (((1.5 / 1.5) - (0.0 + 1.5)) * 1e10)
    if v7 {
        v7 = (!(!(0.0 == 3.14159)))
        if v7 {
            if true {
                v7 = ((65535 > 999) && (2 <= 414))
                for fi11 := 0; fi11 < 1; fi11++ {
                    var v12 float64 = (0.0 - 1e10)
                    switch (fi11 + (fi11 % ((fi11) | 1))) {
                    case 32768:
                        if len(v1) > 0 {
                            v1[0] = 0
                        }
                    case 128:
                        if false {
                            v4 = -22
                            v3[42] = fi11
                        }
                    }
                    v10 += v10
                    _ = fi11
                    _ = v12
                }
                for fi13 := 0; fi13 < 14; fi13++ {
                    v4 += -25
                    _ = fi13
                }
            }
            var v14 string = "\n\t"
            var v15 uint32 = 65535
            _ = v14
            _ = v15
        }
        var v16 uint32 = 897
        _ = v16
    } else {
        if ((42 >= 18446744073709551615) || false) {
            var v17 int16 = 0
            v18 := GenS1{F0: "\n\t", F1: (^(v4))}
            if v7 {
                switch "a" {
                case "foo bar":
                    for fi19 := 0; fi19 < 17; fi19++ {
                        switch ("\n\t" + "foo bar") {
                        case "hello":
                            var v20 uint16 = 65535
                            _ = v20
                        default:
                            if v7 {
                                var v21 uint64 = 857
                                _ = v21
                            } else {
                                if v7 {
                                    _ = len(v3)
                                } else {
                                    v1 = append(v1, 279)
                                    if true {
                                        if v7 {
                                            _ = len(v3)
                                            if (v7 || v7) {
                                                v1 = append(v1, 1)
                                                var v22 float64 = -2.25
                                                _ = v22
                                            } else {
                                                v4 += (^(v4))
                                                var v23 int16 = v17
                                                if v7 {
                                                    var v24 int = (fi19 & (-(fi19)))
                                                    _ = v24
                                                } else {
                                                    fmt.Println(v2, v10)
                                                    v25 := GenS1{F0: "foo bar", F1: v4}
                                                    v4 = -118
                                                    _ = v25
                                                }
                                                _ = v23
                                            }
                                            v7 = ((v2 >> ((2) & 63)) < v2)
                                        }
                                        v10 += v10
                                    } else {
                                        v1 = append(v1, 784)
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        switch "\n\t" {
                                        case "a":
                                            for fi26 := 0; fi26 < 18; fi26++ {
                                                v27 := GenS2{F0: 78, F1: v17}
                                                var v28 string = "a"
                                                var v29 int64 = -1
                                                _ = fi26
                                                _ = v27
                                                _ = v28
                                                _ = v29
                                            }
                                        case "":
                                            var v30 int32 = (v2 % (((v2 - (^(v2)))) | 1))
                                            _ = v30
                                        case "\n\t":
                                            for fi31 := 0; fi31 < 2; fi31++ {
                                                var v32 int8 = -3
                                                if (65535 < 1) {
                                                    if (-752 <= 42) {
                                                        switch (fi19 | fi19) {
                                                        case 630:
                                                            v3[150] = (-(fi19))
                                                        case 111:
                                                            v32 = -118
                                                        default:
                                                            v33 := v3[2]
                                                            _ = v33
                                                        }
                                                        v34 := GenS1{F0: "a", F1: (v32 & (^(v4)))}
                                                        _ = v34
                                                    }
                                                    v35 := GenS3{F0: 42}
                                                    _ = v35
                                                } else {
                                                    var v36 int16 = v17
                                                    v36 -= v17
                                                    var v37 string = ""
                                                    _ = v36
                                                    _ = v37
                                                }
                                                _ = len(v1)
                                                _ = fi31
                                                _ = v32
                                            }
                                        }
                                    }
                                    _ = len(v1)
                                }
                                v7 = (1 > 2)
                                v1 = append(v1, 42)
                            }
                        }
                        fmt.Println(v7, fi19)
                        _ = fi19
                    }
                }
            } else {
                if v7 {
                    v4 -= (^(v4))
                    v4 = v4
                    switch "a" {
                    case "hello":
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    case "":
                        var v38 int8 = (v4 ^ -128)
                        _ = v38
                    case "\n\t":
                        v39 := GenS1{F0: (("" + "a") + ("hello" + "")), F1: (v4 * (^(v4)))}
                        _ = v39
                    default:
                        var v40 bool = (v7 && v7)
                        _ = v40
                    }
                } else {
                    if v7 {
                        for fi41 := 0; fi41 < 5; fi41++ {
                            for fi42 := 0; fi42 < 9; fi42++ {
                                delete(v3, 249)
                                _ = fi42
                            }
                            _ = fi41
                        }
                    } else {
                        var v43 uint = 956
                        _ = v43
                    }
                    delete(v3, 2147483647)
                    if v7 {
                        var v44 float32 = v10
                        _ = v44
                    } else {
                        fmt.Println(v7, v2, v10)
                        switch "hello" {
                        case "\n\t":
                            for fi45 := 0; fi45 < 10; fi45++ {
                                if v7 {
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                } else {
                                    v46 := GenS1{F0: ("hello" + ""), F1: (v4 << ((881) & 63))}
                                    _ = v46
                                }
                                _ = fi45
                            }
                        default:
                            if (!v7) {
                                if (!v7) {
                                    for fi47 := 0; fi47 < 10; fi47++ {
                                        v48 := GenS2{F0: 1, F1: (v17 / ((int16(v4)) | 1))}
                                        if (v7 && (((((0.0 / 1.5) + -2.25) + (1e10 / float64((((v10 - ((-2.25 / -2.25) * v10)) / v10) / (v10 - (float32(-2.25) / v10)))))) + 0.1) <= 0.0)) {
                                            v17 -= (v17 * 1000)
                                        }
                                        _ = fi47
                                        _ = v48
                                    }
                                } else {
                                    var v49 float64 = 3.14159
                                    var v50 int32 = v2
                                    _ = v49
                                    _ = v50
                                }
                                v51 := GenS4{F0: 42, F1: 42, F2: 4294967295}
                                var v52 uint = 4294967295
                                _ = v51
                                _ = v52
                            }
                        }
                        delete(v3, 1)
                    }
                }
            }
            _ = v17
            _ = v18
        } else {
            switch "\n\t" {
            case "\n\t":
                v53 := GenS3{F0: 4294967295}
                _ = v53
            case "a":
                v7 = v7
            }
        }
    }
    var v54 uint32 = 781
    var v55 bool = true
    for fi56 := 0; fi56 < 8; fi56++ {
        _ = len(v1)
        switch "a" {
        case "a":
            v4 = (v4 / ((2) | 1))
        case "hello":
            v1 = append(v1, v54)
        case "\n\t":
            for fi57 := 0; fi57 < 17; fi57++ {
                v3[v54] = fi56
                if v7 {
                    if (fi56 == (^(fi56))) {
                        if ((1 == 2) && true) {
                            if ((-(fi56)) != fi56) {
                                switch "foo bar" {
                                case "":
                                    var v58 uint64 = 4294967295
                                    _ = v58
                                }
                                _ = len(v1)
                            }
                            fmt.Println(v10, fi56, v10)
                            v59 := GenS2{F0: -32768, F1: 225}
                            _ = v59
                        } else {
                            var v60 int8 = (v4 >> ((2) & 63))
                            _ = v60
                        }
                    } else {
                        v61 := GenS3{F0: 42}
                        fmt.Println(v7, v54, fi56)
                        var v62 int16 = -396
                        v63 := GenS5{F0: 1, F1: v4, F2: "\n\t"}
                        _ = v61
                        _ = v62
                        _ = v63
                    }
                    switch (fi57 ^ (fi57 + fi56)) {
                    case 1:
                        v55 = v7
                    case 128:
                        var v64 int64 = 0
                        _ = v64
                    default:
                        v65 := GenS2{F0: -1, F1: 0}
                        delete(v3, (v54 + (^(v54))))
                        _ = v65
                    }
                } else {
                    var v66 uint = 187
                    v67 := GenS6{F0: 2, F1: fi56}
                    delete(v3, 913)
                    v7 = ("hello" > "hello")
                    _ = v66
                    _ = v67
                }
                _ = fi57
            }
        }
        _ = fi56
    }
    var v68 uint8 = 1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v54:", v54)
    fmt.Println("v55:", v55)
    fmt.Println("v68:", v68)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v54
    _ = v55
    _ = v6
    _ = v68
    _ = v7
    _ = v8
}
