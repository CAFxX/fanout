package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint64; F1 bool; F2 int64 }
type GenS3 struct { F0 uint8; F1 int64; F2 int16 }
type GenS4 struct { F0 uint32; F1 int16; F2 uint8 }
type GenS5 struct { F0 uint8; F1 uint64; F2 bool }
type GenS6 struct { F0 int16; F1 uint64 }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 uint16; F1 int64; F2 uint }
type GenS9 struct { F0 bool; F1 uint64 }
type GenS10 struct { F0 uint64; F1 int16; F2 int8 }
func genFn0() int32 {
    var v6 bool = true
    var v7 float32 = 1.5
    switch "\n\t" {
    case "":
        v7 -= v7
    case "a":
        v6 = v6
    case "hello":
        var v8 int16 = -32768
        _ = v8
    default:
        v9 := []int{-282}
        _ = v9
    }
    var v10 int32 = -2147483648
    for fi11 := 0; fi11 < 3; fi11++ {
        for fi12 := 0; fi12 < 10; fi12++ {
            var v13 float64 = float64((-2.25 * 0.1))
            var v14 uint64 = 2
            _ = fi12
            _ = v13
            _ = v14
        }
        var v15 float64 = 3.14159
        v16 := GenS1{F0: 65535}
        _ = fi11
        _ = v15
        _ = v16
    }
    _ = v10
    _ = v6
    _ = v7
    return v10
}
func genFn1(p0 uint8, p1 int8) int8 {
    v17 := GenS1{F0: 865}
    var v18 bool = (false && (-32768 > 0))
    if v18 {
        for fi19 := 0; fi19 < 7; fi19++ {
            v20 := GenS2{F0: 2, F1: v18, F2: 42}
            _ = fi19
            _ = v20
        }
    }
    p0 = (p0 >> ((18446744073709551615) & 63))
    var v21 bool = (1 >= 127)
    _ = p0
    _ = p1
    _ = v17
    _ = v18
    _ = v21
    return p1
}

func main() {
    v1 := []string{}
    var v2 int64 = 42
    var v3 uint16 = 1
    v4 := make(map[string]string)
    var v5 float32 = float32((-2.25 - (float64(1e10) - 1.5)))
    _ = len(v1)
    if false {
        fmt.Println(v2, v5)
        var v22 float32 = (0.1 / v5)
        var v23 uint16 = 976
        _ = v22
        _ = v23
    }
    v2 += v2
    if (-1 != -148) {
        for fi24 := 0; fi24 < 15; fi24++ {
            var v25 int16 = -215
            for fi26 := 0; fi26 < 1; fi26++ {
                var v27 uint32 = 227
                switch 127 {
                case -460:
                    v28 := v4["foo bar"]
                    _ = v28
                default:
                    v29 := GenS2{F0: 1, F1: false, F2: v2}
                    _ = v29
                }
                v27 -= v27
                v25 = 32767
                var v30 uint = 1
                var v31 int16 = (v25 - v25)
                fmt.Println(fi24, v31)
                _ = fi26
                _ = v27
                _ = v30
                _ = v31
            }
            _ = fi24
            _ = v25
        }
        if (true || (0 <= -452)) {
            switch 26 {
            case 1:
                v3 = (v3 + 42)
            case 362:
                if false {
                    v32 := GenS1{F0: 18446744073709551615}
                    switch 59 {
                    case 32767:
                        _ = len(v1)
                    case -2147483648:
                        fmt.Println(v3, v3, v2)
                    case 2147483647:
                        switch "\n\t" {
                        case "foo bar":
                            v33 := GenS3{F0: 128, F1: (v2 ^ (-(v2))), F2: -388}
                            _ = v33
                        default:
                            if false {
                                switch ("\n\t" + ("a" + "\n\t")) {
                                case "hello":
                                    switch "a" {
                                    case "hello":
                                        if (((int(v3) > -9223372036854775808) && false) || true) {
                                            if (!(32767 != -1)) {
                                                if false {
                                                    if true {
                                                        if true {
                                                            v34 := GenS4{F0: 931, F1: 42, F2: 2}
                                                            _ = v34
                                                        }
                                                        if len(v1) > 0 {
                                                            v1[0] = (("a" + ("a" + "\n\t")) + (((("\n\t" + "hello") + "\n\t") + "hello") + "\n\t"))
                                                        }
                                                        for fi35 := 0; fi35 < 15; fi35++ {
                                                            var v36 uint8 = 2
                                                            v3 = (v3 << ((2) & 63))
                                                            _ = fi35
                                                            _ = v36
                                                        }
                                                    }
                                                }
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                            } else {
                                                var v37 bool = (359 <= 317)
                                                _ = v37
                                            }
                                            if true {
                                                v38 := GenS4{F0: 2, F1: 1000, F2: 0}
                                                if len(v1) > 0 {
                                                    v1[0] = ("hello" + "foo bar")
                                                }
                                                _ = v38
                                            }
                                            fmt.Println(v5)
                                        }
                                    case "":
                                        var v39 int16 = -1
                                        _ = v39
                                    }
                                case "a":
                                    var v40 int64 = (^(v2))
                                    _ = v40
                                case "\n\t":
                                    switch 1 {
                                    case 127:
                                        switch 2147483647 {
                                        case 32767:
                                            v5 -= (v5 * ((v5 / float32(float64(3.14159))) * 1e10))
                                        case 4294967295:
                                            var v41 int16 = 1000
                                            _ = v41
                                        case 2147483647:
                                            v42 := GenS3{F0: 128, F1: v2, F2: 1000}
                                            _ = v42
                                        }
                                    case 1:
                                        var v43 float64 = float64(-2.25)
                                        _ = v43
                                    }
                                }
                                v44 := GenS5{F0: 78, F1: 4294967295, F2: false}
                                _ = v44
                            }
                        }
                    default:
                        for fi45 := 0; fi45 < 1; fi45++ {
                            fmt.Println(v5, fi45)
                            _ = len(v1)
                            _ = fi45
                        }
                    }
                    for fi46 := 0; fi46 < 19; fi46++ {
                        _ = len(v1)
                        if true {
                            v47 := GenS5{F0: 127, F1: 42, F2: true}
                            var v48 int8 = -128
                            if false {
                                fmt.Println(v3, v5)
                            }
                            _ = v47
                            _ = v48
                        }
                        _ = fi46
                    }
                    _ = v32
                }
            default:
                v2 -= v2
            }
            if true {
                var v49 uint = 18446744073709551615
                switch 32768 {
                case -802:
                    _ = len(v4)
                case 4294967295:
                    v5 -= ((1.5 - (3.14159 + 3.14159)) - 0.0)
                case 128:
                    v50 := GenS6{F0: 1, F1: 724}
                    _ = v50
                default:
                    v51 := GenS6{F0: 42, F1: 0}
                    _ = v51
                }
                _ = v49
            } else {
                if true {
                    v52 := GenS7{F0: 362}
                    _ = v52
                } else {
                    var v53 int32 = genFn0()
                    if (true || (1 != 42)) {
                        v54 := GenS7{F0: 0}
                        if (!((0 == 134) && true)) {
                            fmt.Println(v2, v53, v5)
                        }
                        _ = v54
                    } else {
                        if ((((v5 / v5) + v5) > v5) || (true && false)) {
                            v2 -= (v2 - v2)
                        } else {
                            var v55 float32 = v5
                            _ = v55
                        }
                        for fi56 := 0; fi56 < 7; fi56++ {
                            for fi57 := 0; fi57 < 5; fi57++ {
                                v58 := GenS8{F0: (^(v3)), F1: 1, F2: 255}
                                var v59 float32 = float32(1.5)
                                _ = fi57
                                _ = v58
                                _ = v59
                            }
                            fmt.Println(v5, v2, v2)
                            _ = fi56
                        }
                    }
                    v53 = -625
                    _ = v53
                }
                _ = len(v1)
            }
            v3 = 65535
        } else {
            var v60 uint = 2
            if true {
                if ((0.0 == -2.25) || false) {
                    v61 := GenS9{F0: ((false && true) && true), F1: 99}
                    if false {
                        if (!true) {
                            fmt.Println(v3, v2, v3)
                        } else {
                            if true {
                                v62 := GenS5{F0: 0, F1: 0, F2: false}
                                v5 -= ((v5 - float32(1e10)) / v5)
                                _ = v62
                            } else {
                                if (2147483647 == genFn0()) {
                                    for fi63 := 0; fi63 < 4; fi63++ {
                                        switch fi63 {
                                        case -2147483648:
                                            fmt.Println(v3, v3)
                                        case 4294967295:
                                            if len(v1) > 0 {
                                                v1[0] = "\n\t"
                                            }
                                        case -450:
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                        default:
                                            if len(v1) > 0 {
                                                v1[0] = ("" + "foo bar")
                                            }
                                        }
                                        _ = fi63
                                    }
                                    v5 = v5
                                    v5 = v5
                                }
                                var v64 float64 = float64(0.1)
                                if true {
                                    var v65 uint32 = 4294967295
                                    _ = v65
                                }
                                _ = v64
                            }
                        }
                    }
                    if false {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        v66 := GenS5{F0: 127, F1: uint64(v60), F2: false}
                        fmt.Println(v5)
                        _ = v66
                    }
                    _ = v61
                }
                var v67 uint32 = uint32(v3)
                _ = v67
            }
            _ = v60
        }
        var v68 uint16 = (v3 / ((v3) | 1))
        _ = v68
    }
    v69 := GenS10{F0: 1, F1: 0, F2: -1}
    fmt.Println(v5, v3, v5)
    v2 -= v2
    var v70 int32 = 2147483647
    v5 += v5
    for fi71 := 0; fi71 < 5; fi71++ {
        v2 += (v2 & (v2 + v2))
        _ = fi71
    }
    if true {
        var v72 uint8 = 2
        _ = v72
    }
    var v73 float64 = float64(v5)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v70:", v70)
    fmt.Println("v73:", v73)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v69
    _ = v70
    _ = v73
}
