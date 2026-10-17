package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64; F2 uint16 }
type GenS2 struct { F0 int; F1 bool; F2 uint32 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 bool; F1 uint8 }
type GenS5 struct { F0 uint64; F1 int16; F2 uint32 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 uint; F1 uint32 }
type GenS9 struct { F0 int8; F1 uint8 }
type GenS10 struct { F0 int32; F1 string; F2 int32 }
type GenS11 struct { F0 int32; F1 int64; F2 int16 }

func main() {
    var v1 uint8 = 0
    var v2 string = "\n\t"
    var v3 uint32 = 1
    var v4 int8 = 59
    var v5 int16 = 32767
    var v6 int8 = v4
    v7 := []int{194, -799, 32768, -61}
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    fmt.Println(v5, v1, v2)
    switch 51 {
    case -9223372036854775808:
        for fi8 := 0; fi8 < 8; fi8++ {
            v9 := GenS1{F0: 0, F1: 21, F2: 42}
            switch "a" {
            case "":
                v10 := GenS2{F0: fi8, F1: (v1 == v1), F2: (^(v3))}
                _ = v10
            }
            var v11 uint64 = 0
            _ = fi8
            _ = v11
            _ = v9
        }
    case 2147483647:
        for fi12 := 0; fi12 < 13; fi12++ {
            fmt.Println(fi12, v3, v4)
            switch "a" {
            case "":
                if true {
                    if (65535 <= 0) {
                        if (18446744073709551615 > 18446744073709551615) {
                            for fi13 := 0; fi13 < 9; fi13++ {
                                var v14 string = ((v2 + "hello") + "")
                                if (false || false) {
                                    v15 := GenS2{F0: fi13, F1: ((fi12 << ((58) & 63)) >= (fi12 * (fi13 / ((0) | 1)))), F2: (v3 % ((v3) | 1))}
                                    _ = v15
                                } else {
                                    switch 1 {
                                    case 4294967295:
                                        switch (fi12 & 175) {
                                        case 1:
                                            var v16 int64 = -201
                                            _ = v16
                                        default:
                                            var v17 bool = ((!(false || true)) && true)
                                            _ = v17
                                        }
                                        fmt.Println(fi12, v6, v1)
                                        switch v14 {
                                        case "\n\t":
                                            if (!((true && true) || (447 >= -596))) {
                                                _ = len(v7)
                                            } else {
                                                v18 := GenS2{F0: fi12, F1: false, F2: v3}
                                                _ = v18
                                            }
                                        }
                                        switch fi12 {
                                        case 812:
                                            v5 -= 1
                                            v19 := GenS1{F0: v1, F1: 910, F2: 2}
                                            _ = v19
                                        default:
                                            switch "foo bar" {
                                            case "foo bar":
                                                switch 0 {
                                                case 128:
                                                    if len(v7) >= 2 {
                                                        v7 = v7[1:len(v7)]
                                                    }
                                                default:
                                                    var v20 int32 = -728
                                                    _ = v20
                                                }
                                            case "\n\t":
                                                if true {
                                                    fmt.Println(v2, v6, v14)
                                                    v1 += v1
                                                    switch 32768 {
                                                    case -524:
                                                        var v21 uint = 18446744073709551615
                                                        _ = v21
                                                    case 0:
                                                        var v22 int32 = 2147483647
                                                        _ = v22
                                                    case -969:
                                                        if len(v7) >= 2 {
                                                            v7 = v7[1:len(v7)]
                                                        }
                                                        var v23 float32 = (0.0 + (-2.25 + 0.1))
                                                        if false {
                                                            if true {
                                                                if len(v7) >= 2 {
                                                                    v7 = v7[1:len(v7)]
                                                                }
                                                            }
                                                            if len(v7) >= 2 {
                                                                v7 = v7[1:len(v7)]
                                                            }
                                                            v24 := GenS1{F0: 128, F1: 1, F2: 12}
                                                            _ = v24
                                                        } else {
                                                            v23 += (v23 + v23)
                                                            v5 = (-(v5))
                                                            fmt.Println(v2, v3)
                                                        }
                                                        _ = v23
                                                    }
                                                }
                                            case "":
                                                var v25 bool = (!false)
                                                _ = v25
                                            default:
                                                if len(v7) > 0 {
                                                    v7[0] = -9223372036854775808
                                                }
                                            }
                                        }
                                    default:
                                        switch v14 {
                                        case "\n\t":
                                            v26 := GenS2{F0: (fi12 * 42), F1: true, F2: (v3 + v3)}
                                            _ = v26
                                        case "foo bar":
                                            v27 := GenS3{F0: true}
                                            _ = v27
                                        default:
                                            if true {
                                                v6 = (v4 << ((476) & 63))
                                                v3 = 65535
                                                if (true || (!true)) {
                                                    fmt.Println(v6, v1, fi13)
                                                    var v28 float32 = 1.5
                                                    v4 += v4
                                                    _ = v28
                                                } else {
                                                    switch 98 {
                                                    case -1:
                                                        v29 := GenS2{F0: 4294967295, F1: (false && ((^(v3)) <= 2)), F2: (v3 | (v3 ^ (v3 & 1)))}
                                                        _ = v29
                                                    }
                                                }
                                            }
                                        }
                                    }
                                    v30 := GenS4{F0: ((!false) && (!false)), F1: 127}
                                    v7 = append(v7, (fi12 << ((528) & 63)))
                                    _ = v30
                                }
                                var v31 int8 = v6
                                _ = fi13
                                _ = v14
                                _ = v31
                            }
                        }
                        if len(v7) > 0 {
                            v7[0] = (fi12 - -151)
                        }
                        var v32 string = v2
                        _ = v32
                    } else {
                        var v33 string = (v2 + v2)
                        for fi34 := 0; fi34 < 8; fi34++ {
                            if false {
                                v35 := GenS5{F0: 1, F1: v5, F2: v3}
                                v7 = append(v7, 32768)
                                _ = v35
                            }
                            if (-1 <= 3) {
                                var v36 int = (fi12 - int(v6))
                                if ((((3.14159 / -2.25) / 1.0) + 0.1) <= (1e10 + (((float64(-2.25) + 0.0) * 0.1) / (-2.25 * (3.14159 * (0.1 / 1.0)))))) {
                                    if len(v7) >= 2 {
                                        v7 = v7[1:len(v7)]
                                    }
                                    v37 := GenS3{F0: false}
                                    if (4294967295 < 2) {
                                        v5 = v5
                                    }
                                    _ = v37
                                } else {
                                    fmt.Println(fi34, v33, v2)
                                    switch (fi12 - v36) {
                                    case 9223372036854775807:
                                        var v38 uint16 = 0
                                        _ = v38
                                    case 32767:
                                        v36 = (fi34 << ((288) & 63))
                                    case -360:
                                        v39 := GenS6{F0: 1}
                                        _ = v39
                                    }
                                }
                                _ = v36
                            } else {
                                if false {
                                    v6 -= -1
                                    v5 += (v5 | v5)
                                    fmt.Println(v33)
                                } else {
                                    var v40 string = "foo bar"
                                    v40 = "a"
                                    _ = v40
                                }
                                v41 := GenS7{F0: 2}
                                var v42 uint = 18446744073709551615
                                _ = v41
                                _ = v42
                            }
                            v43 := GenS3{F0: true}
                            _ = fi34
                            _ = v43
                        }
                        _ = v33
                    }
                } else {
                    v44 := GenS8{F0: 20, F1: 132}
                    switch ("foo bar" + (v2 + "")) {
                    case "":
                        var v45 int8 = v6
                        _ = v45
                    case "a":
                        for fi46 := 0; fi46 < 12; fi46++ {
                            switch v2 {
                            case "foo bar":
                                if (!false) {
                                    var v47 float32 = -2.25
                                    v3 = (v3 ^ 2147483648)
                                    _ = v47
                                } else {
                                    var v48 float64 = 1.5
                                    _ = v48
                                }
                            }
                            if true {
                                v7 = append(v7, (fi12 * fi12))
                                switch (v2 + v2) {
                                case "\n\t":
                                    v4 = -7
                                default:
                                    if true {
                                        v49 := GenS9{F0: v4, F1: v1}
                                        switch (v2 + v2) {
                                        case "\n\t":
                                            v50 := GenS10{F0: int32(v3), F1: "\n\t", F2: 2147483647}
                                            _ = v50
                                        }
                                        var v51 uint32 = 533
                                        _ = v49
                                        _ = v51
                                    }
                                }
                            }
                            switch fi12 {
                            case -2147483648:
                                var v52 bool = true
                                _ = v52
                            default:
                                var v53 bool = (0 <= v5)
                                _ = v53
                            }
                            _ = fi46
                        }
                    default:
                        var v54 uint32 = v3
                        _ = v54
                    }
                    _ = v44
                }
            }
            var v55 int8 = 1
            _ = fi12
            _ = v55
        }
    default:
        if true {
            var v56 int32 = -1
            _ = v56
        }
    }
    if false {
        v6 = v4
        v6 = v4
    }
    v1 = (^(v1))
    v4 -= -60
    for fi57 := 0; fi57 < 1; fi57++ {
        v58 := GenS6{F0: 42}
        v2 = v2
        var v59 int64 = 1
        _ = fi57
        _ = v58
        _ = v59
    }
    for fi60 := 0; fi60 < 11; fi60++ {
        v6 += 35
        _ = fi60
    }
    v1 = (^(v1))
    v61 := GenS11{F0: 42, F1: 9223372036854775807, F2: (^(v5))}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v61
    _ = v7
}
