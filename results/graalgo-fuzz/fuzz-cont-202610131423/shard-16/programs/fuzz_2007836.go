package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64; F2 string }
type GenS2 struct { F0 int64; F1 uint16; F2 uint64 }
type GenS3 struct { F0 int32; F1 uint32 }
type GenS4 struct { F0 int32; F1 string; F2 int8 }

func main() {
    var v1 uint = 203
    var v2 uint16 = 42
    var v3 int = 128
    var v4 uint64 = 4294967295
    v5 := make(map[uint32]int)
    var v6 uint = 65535
    for fi7 := 0; fi7 < 11; fi7++ {
        if (!true) {
            for fi8 := 0; fi8 < 10; fi8++ {
                v9 := []int{(-(fi7)), fi8, (fi7 >> (((v6 * 42)) & 63)), 1}
                var v10 float32 = 0.0
                var v11 int16 = -1
                _ = fi8
                _ = v10
                _ = v11
                _ = v9
            }
            for fi12 := 0; fi12 < 20; fi12++ {
                if true {
                    if (!false) {
                        _ = len(v5)
                    }
                    switch ("a" + ("" + "foo bar")) {
                    case "hello":
                        var v13 uint8 = 9
                        _ = v13
                    case "\n\t":
                        if (!false) {
                            fmt.Println(v4)
                            v14 := GenS1{F0: 62, F1: 1, F2: "hello"}
                            v1 = 2
                            _ = v14
                        }
                    case "foo bar":
                        switch ("a" + (((("" + "foo bar") + ((("a" + ("foo bar" + "")) + "\n\t") + "foo bar")) + "hello") + "")) {
                        case "a":
                            var v15 int = fi7
                            var v16 uint16 = 0
                            if false {
                                var v17 bool = (true || true)
                                var v18 int64 = -377
                                _ = v17
                                _ = v18
                            }
                            _ = v15
                            _ = v16
                        case "foo bar":
                            var v19 int8 = 66
                            _ = v19
                        case "hello":
                            v6 += uint(v4)
                        }
                    }
                    v4 = 187
                }
                var v20 uint64 = (^(v4))
                var v21 bool = (!(127 < int8(v2)))
                _ = fi12
                _ = v20
                _ = v21
            }
            v22 := []int{(v3 - v3), (v3 * v3), fi7}
            _ = v22
        }
        _ = fi7
    }
    v23 := []int{v3, v3, 1, v3}
    fmt.Println(v4, v6, v4)
    if (!(37 != 0)) {
        v3 = v3
        if (!(!(((float32(float64(1.5)) / float32(float64((float32(1.5) / (3.14159 * -2.25))))) * (((0.0 + 3.14159) - -2.25) - (-2.25 * float32(0.0)))) >= 0.0))) {
            v6 = v1
        } else {
            v4 = (v4 % (((^(v4))) | 1))
            if (true && false) {
                for fi24 := 0; fi24 < 14; fi24++ {
                    var v25 uint8 = 1
                    var v26 float64 = 1.5
                    _ = fi24
                    _ = v25
                    _ = v26
                }
                v2 -= (v2 - 681)
                for fi27 := 0; fi27 < 6; fi27++ {
                    if (v4 > (v4 << ((v6) & 63))) {
                        var v28 int = (fi27 * v3)
                        for fi29 := 0; fi29 < 6; fi29++ {
                            v30 := GenS1{F0: -32768, F1: -9223372036854775808, F2: ("a" + "\n\t")}
                            v31 := GenS2{F0: int64(v2), F1: 673, F2: 259}
                            _ = fi29
                            _ = v30
                            _ = v31
                        }
                        _ = v28
                    } else {
                        v4 -= 18446744073709551615
                    }
                    var v32 float64 = -2.25
                    v33 := GenS2{F0: 42, F1: v2, F2: v4}
                    _ = fi27
                    _ = v32
                    _ = v33
                }
            }
            var v34 int = 4294967295
            _ = v34
        }
    }
    if true {
        if false {
            var v35 int32 = -263
            var v36 float32 = ((0.0 - 1.5) * 1.5)
            var v37 float32 = v36
            _ = v35
            _ = v36
            _ = v37
        }
    }
    if true {
        switch "\n\t" {
        case "foo bar":
            v2 = v2
        case "a":
            v2 = (v2 % (((v2 / ((v2) | 1))) | 1))
        default:
            if (1 > 42) {
                var v38 int = (v3 - v3)
                delete(v5, 2)
                if (1.5 <= 0.0) {
                    for fi39 := 0; fi39 < 11; fi39++ {
                        var v40 uint = (v6 % ((v6) | 1))
                        v2 -= v2
                        fmt.Println(v3, v40)
                        _ = fi39
                        _ = v40
                    }
                    var v41 float64 = 1.5
                    _ = v41
                } else {
                    for fi42 := 0; fi42 < 11; fi42++ {
                        if true {
                            fmt.Println(v38, v4)
                            switch v38 {
                            case -9223372036854775808:
                                var v43 uint16 = v2
                                _ = v43
                            case 51:
                                var v44 uint32 = 9
                                _ = v44
                            case 32768:
                                fmt.Println(fi42, v6)
                            default:
                                for fi45 := 0; fi45 < 11; fi45++ {
                                    fmt.Println(v3)
                                    var v46 uint16 = v2
                                    _ = fi45
                                    _ = v46
                                }
                            }
                            v3 = (-(fi42))
                        } else {
                            var v47 float64 = 1e10
                            v48 := GenS1{F0: 42, F1: -9223372036854775808, F2: ("foo bar" + (((("\n\t" + "\n\t") + "a") + "a") + ""))}
                            var v49 float32 = (3.14159 / -2.25)
                            _ = v47
                            _ = v48
                            _ = v49
                        }
                        var v50 uint16 = 338
                        for fi51 := 0; fi51 < 18; fi51++ {
                            v38 = (v38 >> ((0) & 63))
                            if (v2 < v2) {
                                v5[2] = v38
                                if false {
                                    fmt.Println(fi51, v1, v6)
                                    v2 -= v50
                                }
                            } else {
                                var v52 uint32 = 295
                                if (false || true) {
                                    var v53 bool = false
                                    _ = v53
                                } else {
                                    delete(v5, (v52 / ((65535) | 1)))
                                    var v54 string = "foo bar"
                                    _ = v54
                                }
                                _ = v52
                            }
                            _ = fi51
                        }
                        _ = fi42
                        _ = v50
                    }
                }
                _ = v38
            }
        }
        var v55 uint8 = 128
        _ = v55
    }
    v2 -= v2
    for fi56 := 0; fi56 < 17; fi56++ {
        v57 := GenS3{F0: int32(v4), F1: 2147483647}
        _ = fi56
        _ = v57
    }
    v4 -= 1
    var v58 uint = 4294967295
    var v59 float32 = (1e10 + (float32(float64((1e10 / 1e10))) / 1e10))
    var v60 int32 = -2147483648
    for fi61 := 0; fi61 < 7; fi61++ {
        var v62 int8 = 127
        v5[65535] = v3
        _ = fi61
        _ = v62
    }
    v63 := GenS3{F0: v60, F1: uint32(v1)}
    v64 := GenS4{F0: (-(v60)), F1: "", F2: 100}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v58:", v58)
    fmt.Println("v59:", v59)
    fmt.Println("v6:", v6)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v58
    _ = v59
    _ = v6
    _ = v60
    _ = v63
    _ = v64
}
