package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16; F2 bool }
type GenS2 struct { F0 int; F1 string }
type GenS3 struct { F0 uint64; F1 int8 }
type GenS4 struct { F0 uint32; F1 uint; F2 int32 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 uint; F1 int }

func main() {
    var v1 uint8 = 146
    v2 := make(map[string]string)
    var v3 string = "\n\t"
    v4 := []int{-445, 4294967295, 2147483647}
    var v5 int8 = 40
    var v6 uint8 = v1
    v2[v3] = v3
    fmt.Println(v5, v1)
    for fi7 := 0; fi7 < 6; fi7++ {
        switch ("a" + (("\n\t" + (("foo bar" + ("hello" + v3)) + (v3 + v3))) + v3)) {
        case "\n\t":
            v5 = v5
        case "foo bar":
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        case "a":
            v8 := GenS1{F0: 1000, F1: 32767, F2: (!false)}
            _ = v8
        default:
            var v9 bool = ((false || (!(true || true))) || false)
            _ = v9
        }
        fmt.Println(v6, v1)
        for fi10 := 0; fi10 < 19; fi10++ {
            switch "\n\t" {
            case "":
                if true {
                    v11 := GenS2{F0: -1, F1: v3}
                    _ = v11
                } else {
                    v12 := GenS2{F0: (fi7 / ((4294967295) | 1)), F1: v3}
                    v2[v3] = (v3 + (v3 + (v3 + v3)))
                    v13 := GenS2{F0: fi10, F1: (v3 + (v3 + v3))}
                    var v14 int8 = v5
                    v15 := v2[v3]
                    _ = v12
                    _ = v13
                    _ = v14
                    _ = v15
                }
            case "a":
                v6 = (v6 ^ (v1 - v1))
            case "\n\t":
                var v16 int32 = 1
                _ = v16
            default:
                if (0.1 <= 1e10) {
                    v3 = "foo bar"
                    if (3.14159 <= 3.14159) {
                        v1 = (v6 << ((42) & 63))
                    }
                    v17 := GenS1{F0: 1000, F1: 42, F2: (!(671 < 1))}
                    if (false || (v3 > (((v3 + v3) + "a") + (v3 + "")))) {
                        var v18 bool = ((-(v5)) >= (-(v5)))
                        v2[v3] = "a"
                        if (v18 && v18) {
                            v18 = (1.5 != (0.1 + -2.25))
                        }
                        _ = v18
                    }
                    _ = v17
                } else {
                    var v19 int16 = -1
                    if (v19 == 42) {
                        if len(v4) > 0 {
                            v4[0] = (^(fi7))
                        }
                        switch v3 {
                        case "\n\t":
                            var v20 int64 = -473
                            v20 = 42
                            v1 -= 1
                            _ = v20
                        }
                        var v21 float64 = -2.25
                        _ = v21
                    } else {
                        switch "a" {
                        case "":
                            _ = len(v4)
                        }
                        v3 = v3
                    }
                    switch (fi7 * (fi10 * (fi10 << ((0) & 63)))) {
                    case 324:
                        v22 := GenS3{F0: uint64(v6), F1: (v5 - (v5 | (v5 + v5)))}
                        _ = v22
                    case 995:
                        v23 := v2[v3]
                        _ = v23
                    default:
                        switch (^(fi7)) {
                        case -1:
                            v24 := GenS3{F0: 968, F1: -42}
                            _ = v24
                        case 2147483647:
                            v19 = (v19 + v19)
                        case 128:
                            fmt.Println(v6, v3)
                        }
                    }
                    _ = v19
                }
            }
            var v25 float64 = float64(-2.25)
            var v26 bool = false
            if false {
                var v27 int64 = -581
                _ = v27
            }
            v28 := GenS4{F0: 0, F1: 519, F2: 1}
            _ = len(v2)
            _ = fi10
            _ = v25
            _ = v26
            _ = v28
        }
        _ = fi7
    }
    for fi29 := 0; fi29 < 17; fi29++ {
        v30 := GenS4{F0: 395, F1: 242, F2: 2147483647}
        var v31 float64 = 1e10
        v6 = v1
        _ = fi29
        _ = v30
        _ = v31
    }
    fmt.Println(v5, v6)
    v32 := GenS3{F0: 557, F1: (v5 / ((-128) | 1))}
    var v33 bool = (false || (true || (749 >= 32767)))
    if len(v4) > 0 {
        v4[0] = 4294967295
    }
    if (-1 < 32768) {
        v33 = (int16(v6) != 42)
        for fi34 := 0; fi34 < 10; fi34++ {
            v5 += (^(v5))
            v35 := GenS5{F0: 42}
            _ = fi34
            _ = v35
        }
    } else {
        v2["a"] = v3
    }
    if (!(42 < 1)) {
        switch 32768 {
        case 6:
            if (2147483648 == 2) {
                var v36 uint16 = 42
                for fi37 := 0; fi37 < 9; fi37++ {
                    if (-9223372036854775808 == (fi37 ^ (fi37 >> ((2) & 63)))) {
                        if true {
                            v6 -= 255
                        }
                        v3 = ""
                    }
                    v33 = (!((v6 - 42) != 255))
                    _ = fi37
                }
                _ = v36
            }
        case 507:
            v38 := GenS6{F0: 10}
            _ = v38
        default:
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
        v39 := GenS7{F0: 255, F1: -826}
        var v40 string = v3
        _ = v39
        _ = v40
    } else {
        switch -656 {
        case 32767:
            if true {
                fmt.Println(v1, v33, v5)
                if true {
                    switch 2147483647 {
                    case 459:
                        if ((1e10 > 1e10) || v33) {
                            if (255 > 255) {
                                switch 32768 {
                                case -9223372036854775808:
                                    for fi41 := 0; fi41 < 2; fi41++ {
                                        var v42 float64 = 0.0
                                        delete(v2, "a")
                                        if v33 {
                                            var v43 uint32 = 735
                                            _ = v43
                                        } else {
                                            v3 = "hello"
                                        }
                                        _ = fi41
                                        _ = v42
                                    }
                                case 0:
                                    if ((((float64((3.14159 + 1.5)) * 1.5) * -2.25) != -2.25) && true) {
                                        var v44 uint = 0
                                        var v45 int32 = -1
                                        _ = v44
                                        _ = v45
                                    } else {
                                        var v46 uint16 = 457
                                        var v47 int64 = 162
                                        v6 += (^(v1))
                                        _ = v46
                                        _ = v47
                                    }
                                }
                            }
                            fmt.Println(v6)
                            v3 = (v3 + (v3 + v3))
                        }
                    case -628:
                        switch v3 {
                        case "\n\t":
                            v1 += 255
                        case "hello":
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        case "":
                            fmt.Println(v3)
                        default:
                            v48 := GenS6{F0: 373}
                            _ = v48
                        }
                    case 12:
                        if v33 {
                            if len(v4) > 0 {
                                v4[0] = 1
                            }
                            v6 -= (v1 << ((42) & 63))
                        } else {
                            switch v3 {
                            case "":
                                if ((^(v5)) == (v5 / (((^(v5))) | 1))) {
                                    _ = len(v4)
                                }
                            case "a":
                                v49 := GenS7{F0: uint(v1), F1: 42}
                                _ = v49
                            }
                            v50 := GenS7{F0: 1, F1: 42}
                            _ = v50
                        }
                    default:
                        fmt.Println(v6)
                    }
                } else {
                    v33 = false
                    v51 := GenS2{F0: 9223372036854775807, F1: ("a" + (v3 + v3))}
                    if ((v3 + (v3 + "hello")) == "hello") {
                        for fi52 := 0; fi52 < 9; fi52++ {
                            var v53 int8 = (v5 ^ v5)
                            delete(v2, v3)
                            _ = fi52
                            _ = v53
                        }
                        if true {
                            v54 := GenS6{F0: 1}
                            if false {
                                v6 = (v1 & v1)
                            } else {
                                fmt.Println(v3, v1, v6)
                                var v55 uint64 = 42
                                _ = v55
                            }
                            switch v3 {
                            case "":
                                v2[(v3 + (v3 + (((v3 + v3) + ("hello" + (v3 + ""))) + v3)))] = v3
                            case "hello":
                                var v56 float32 = 0.1
                                _ = v56
                            }
                            _ = v54
                        } else {
                            v1 = (v1 & (^(v6)))
                            v2[((v3 + v3) + "foo bar")] = "a"
                        }
                    } else {
                        var v57 string = v3
                        var v58 int64 = -9223372036854775808
                        _ = v57
                        _ = v58
                    }
                    _ = v51
                }
            } else {
                for fi59 := 0; fi59 < 7; fi59++ {
                    if v33 {
                        fmt.Println(v5, fi59, fi59)
                        var v60 int64 = -1
                        if true {
                            var v61 int16 = 139
                            v5 += 127
                            var v62 uint32 = 1
                            _ = v61
                            _ = v62
                        }
                        _ = v60
                    } else {
                        fmt.Println(v3, v33)
                        v33 = (1 >= -257)
                    }
                    _ = len(v4)
                    _ = fi59
                }
                for fi63 := 0; fi63 < 5; fi63++ {
                    fmt.Println(v33)
                    _ = fi63
                }
                var v64 uint = 0
                _ = v64
            }
        }
        for fi65 := 0; fi65 < 6; fi65++ {
            if (704 < -624) {
                var v66 float64 = (3.14159 / 1.5)
                var v67 int32 = 0
                var v68 uint64 = 4294967295
                _ = v66
                _ = v67
                _ = v68
            } else {
                var v69 bool = false
                var v70 float32 = 3.14159
                _ = v69
                _ = v70
            }
            _ = fi65
        }
    }
    var v71 uint16 = 2
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v71:", v71)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
    _ = v71
}
