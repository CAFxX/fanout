package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint32; F1 int16 }
type GenS3 struct { F0 int8; F1 int64; F2 uint16 }
type GenS4 struct { F0 uint32; F1 uint16; F2 uint64 }
type GenS5 struct { F0 uint8; F1 int }
func genFn0(p0 int32, p1 bool) uint {
    switch "foo bar" {
    case "a":
        var v7 uint = 1
        _ = v7
    default:
        var v8 int32 = 732
        _ = v8
    }
    p0 += p0
    p1 = p1
    var v9 float64 = 0.1
    _ = p0
    _ = p1
    _ = v9
    return 0
}

func main() {
    v1 := []string{("a" + "hello")}
    var v2 int32 = 2147483647
    var v3 int64 = -9223372036854775808
    var v4 uint = 4294967295
    var v5 uint32 = 1
    v6 := make(map[string]uint32)
    v10 := GenS1{F0: 90}
    if true {
        switch "hello" {
        case "hello":
            for fi11 := 0; fi11 < 4; fi11++ {
                v12 := GenS1{F0: 1}
                if (fi11 > -163) {
                    if (false && true) {
                        var v13 string = "a"
                        v3 = 0
                        if len(v1) > 0 {
                            v1[0] = v13
                        }
                        _ = v13
                    } else {
                        for fi14 := 0; fi14 < 20; fi14++ {
                            switch "foo bar" {
                            case "\n\t":
                                if true {
                                    if len(v1) > 0 {
                                        v1[0] = (("a" + "foo bar") + "\n\t")
                                    }
                                    v15 := GenS1{F0: 255}
                                    _ = v15
                                }
                            case "foo bar":
                                v2 = (v2 ^ v2)
                            default:
                                var v16 uint8 = 1
                                _ = v16
                            }
                            var v17 uint16 = 42
                            fmt.Println(v3, v3, fi11)
                            _ = fi14
                            _ = v17
                        }
                        switch ("foo bar" + "\n\t") {
                        case "a":
                            var v18 uint32 = (v5 * v5)
                            _ = v18
                        }
                    }
                    var v19 int64 = int64(v4)
                    var v20 int8 = 12
                    _ = v19
                    _ = v20
                } else {
                    var v21 uint64 = 546
                    _ = v21
                }
                _ = fi11
                _ = v12
            }
        case "a":
            for fi22 := 0; fi22 < 13; fi22++ {
                v5 += (v5 + v5)
                _ = fi22
            }
        case "":
            var v23 int = 4294967295
            _ = v23
        }
        for fi24 := 0; fi24 < 17; fi24++ {
            v5 = (v5 ^ 42)
            _ = len(v1)
            var v25 float32 = 1.5
            _ = fi24
            _ = v25
        }
    } else {
        if false {
            var v26 uint64 = uint64(v5)
            var v27 int64 = v3
            _ = v26
            _ = v27
        }
        v3 = 0
    }
    var v28 int64 = (^(v3))
    for fi29 := 0; fi29 < 10; fi29++ {
        v4 -= 42
        fmt.Println(v28, v4)
        for fi30 := 0; fi30 < 19; fi30++ {
            if false {
                v31 := v6[("a" + (("hello" + "foo bar") + "\n\t"))]
                switch ("" + ("foo bar" + ("foo bar" + ("" + ("" + ("foo bar" + "hello")))))) {
                case "\n\t":
                    if false {
                        delete(v6, ("\n\t" + ("foo bar" + "hello")))
                        v32 := GenS1{F0: 2}
                        _ = v32
                    }
                case "a":
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    var v33 int = (^(fi29))
                    fmt.Println(v4, v28)
                    _ = v33
                default:
                    var v34 string = "foo bar"
                    _ = v34
                }
                v35 := GenS2{F0: (^(v31)), F1: 1000}
                _ = v31
                _ = v35
            } else {
                var v36 int64 = (v28 + v28)
                switch ((("\n\t" + ("hello" + "")) + "a") + "a") {
                case "":
                    if (!(true && ((42 > 32767) && false))) {
                        var v37 float32 = 1.5
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        _ = v37
                    }
                case "foo bar":
                    if len(v1) > 0 {
                        v1[0] = "a"
                    }
                    switch "" {
                    case "\n\t":
                        var v38 int64 = v28
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        if (1e10 >= float64(0.1)) {
                            v39 := GenS1{F0: 127}
                            if len(v1) > 0 {
                                v1[0] = ((("" + "") + ("a" + (((("" + (("" + "a") + "a")) + ("foo bar" + "hello")) + "hello") + ""))) + ("foo bar" + ""))
                            }
                            v3 = (v38 - v28)
                            v5 = 0
                            _ = v39
                        }
                        var v40 int16 = -480
                        var v41 int64 = v36
                        _ = v38
                        _ = v40
                        _ = v41
                    }
                    v6["foo bar"] = v5
                default:
                    v42 := GenS3{F0: 1, F1: (v36 | v28), F2: 2}
                    _ = v42
                }
                v28 -= (v3 + int64(fi29))
                v43 := GenS4{F0: 0, F1: 42, F2: 42}
                _ = v36
                _ = v43
            }
            v1 = append(v1, "\n\t")
            _ = fi30
        }
        _ = fi29
    }
    switch ("foo bar" + "hello") {
    case "foo bar":
        switch "" {
        case "hello":
            switch "a" {
            case "\n\t":
                for fi44 := 0; fi44 < 15; fi44++ {
                    switch "a" {
                    case "hello":
                        if ((201 < 127) && false) {
                            var v45 int64 = (v3 - (v28 ^ 0))
                            var v46 int32 = v2
                            _ = v45
                            _ = v46
                        } else {
                            fmt.Println(v3)
                            for fi47 := 0; fi47 < 17; fi47++ {
                                var v48 bool = ((v5 * v5) == 2147483648)
                                _ = fi47
                                _ = v48
                            }
                        }
                    default:
                        var v49 uint16 = uint16(v5)
                        _ = v49
                    }
                    v4 -= 1
                    _ = fi44
                }
            case "foo bar":
                if false {
                    for fi50 := 0; fi50 < 11; fi50++ {
                        switch "a" {
                        case "a":
                            if (-607 <= 32767) {
                                for fi51 := 0; fi51 < 16; fi51++ {
                                    var v52 float32 = 1.5
                                    _ = fi51
                                    _ = v52
                                }
                            }
                        case "\n\t":
                            v5 -= (v5 >> ((v4) & 63))
                        default:
                            fmt.Println(v5)
                        }
                        v53 := GenS3{F0: 82, F1: v3, F2: uint16(v5)}
                        v2 = 42
                        _ = fi50
                        _ = v53
                    }
                    fmt.Println(v4, v3)
                    for fi54 := 0; fi54 < 15; fi54++ {
                        v2 -= (v2 * int32(v3))
                        _ = fi54
                    }
                } else {
                    if (false && (0 < uint64(v5))) {
                        v3 += (v28 ^ 42)
                        fmt.Println(v4)
                        delete(v6, "\n\t")
                    }
                }
            }
        }
    case "hello":
        if false {
            var v55 string = "foo bar"
            if (26 == 1) {
                var v56 float64 = -2.25
                var v57 float64 = ((float64(-2.25) / v56) / float64(((-2.25 + 0.0) / (1e10 + -2.25))))
                v58 := GenS4{F0: v5, F1: 195, F2: uint64(v4)}
                _ = v56
                _ = v57
                _ = v58
            } else {
                v59 := GenS1{F0: 255}
                _ = v59
            }
            if false {
                delete(v6, (v55 + (v55 + ("" + "\n\t"))))
                if (472 < 255) {
                    fmt.Println(v2, v55)
                } else {
                    var v60 string = (v55 + "hello")
                    _ = v60
                }
            }
            _ = v55
        } else {
            if false {
                switch -992 {
                case 393:
                    if true {
                        var v61 uint = (v4 / ((v4) | 1))
                        for fi62 := 0; fi62 < 5; fi62++ {
                            v63 := GenS2{F0: v5, F1: 653}
                            _ = fi62
                            _ = v63
                        }
                        _ = v61
                    } else {
                        v4 += (v4 | 368)
                        v5 += (v5 % ((v5) | 1))
                        if false {
                            var v64 bool = true
                            for fi65 := 0; fi65 < 18; fi65++ {
                                fmt.Println(v3)
                                var v66 uint = (v4 | v4)
                                v67 := GenS5{F0: 42, F1: (fi65 ^ (fi65 ^ fi65))}
                                _ = fi65
                                _ = v66
                                _ = v67
                            }
                            _ = v64
                        }
                    }
                case -2147483648:
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                default:
                    fmt.Println(v4)
                }
                fmt.Println(v4, v5, v4)
            } else {
                v5 = 42
                v28 -= (^(v3))
                fmt.Println(v3)
            }
            for fi68 := 0; fi68 < 11; fi68++ {
                v69 := GenS1{F0: 2}
                if true {
                    v5 = (v5 >> ((v4) & 63))
                    for fi70 := 0; fi70 < 13; fi70++ {
                        var v71 bool = false
                        _ = fi70
                        _ = v71
                    }
                } else {
                    fmt.Println(v2, v28, v4)
                }
                _ = fi68
                _ = v69
            }
        }
    }
    switch "" {
    case "foo bar":
        var v72 float32 = 0.0
        _ = v72
    default:
        v28 += -950
    }
    var v73 int8 = -128
    v74 := v6["a"]
    fmt.Println(v73)
    for fi75 := 0; fi75 < 17; fi75++ {
        v28 = v3
        v2 = v2
        _ = fi75
    }
    var v76 float64 = ((1.5 - 1e10) / (1.5 / 1.0))
    var v77 int64 = 1
    v78 := v6["a"]
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v73:", v73)
    fmt.Println("v74:", v74)
    fmt.Println("v76:", v76)
    fmt.Println("v77:", v77)
    fmt.Println("v78:", v78)
    _ = v1
    _ = v10
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v73
    _ = v74
    _ = v76
    _ = v77
    _ = v78
}
