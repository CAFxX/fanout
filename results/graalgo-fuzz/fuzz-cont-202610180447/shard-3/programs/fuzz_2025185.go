package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint }
type GenS2 struct { F0 int16; F1 int8; F2 int32 }
type GenS3 struct { F0 int16; F1 uint8 }
type GenS4 struct { F0 int64; F1 int32 }
type GenS5 struct { F0 uint64; F1 uint16 }
func genFn0(p0 int32, p1 int64, p2 uint64) uint8 {
    if false {
        var v7 bool = true
        _ = v7
    }
    var v8 int64 = (p1 << ((138) & 63))
    v9 := GenS1{F0: (p2 ^ (^(p2))), F1: 1}
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return uint8(p2)
}
func genFn1(p0 string, p1 int64) int64 {
    v10 := GenS2{F0: 758, F1: 119, F2: 2147483647}
    switch "foo bar" {
    case "\n\t":
        for fi11 := 0; fi11 < 10; fi11++ {
            var v12 float32 = (-2.25 / 1.5)
            var v13 int64 = p1
            var v14 int64 = v13
            _ = fi11
            _ = v12
            _ = v13
            _ = v14
        }
    case "foo bar":
        var v15 string = p0
        _ = v15
    case "":
        for fi16 := 0; fi16 < 6; fi16++ {
            v17 := GenS1{F0: 4294967295, F1: 65535}
            var v18 uint8 = 31
            for fi19 := 0; fi19 < 4; fi19++ {
                fmt.Println(fi16, v18, p1)
                _ = fi19
            }
            _ = fi16
            _ = v17
            _ = v18
        }
    default:
        p1 += (p1 << ((756) & 63))
    }
    _ = p0
    _ = p1
    _ = v10
    return p1
}

func main() {
    v1 := []uint32{}
    var v2 int32 = 0
    var v3 string = "hello"
    var v4 int8 = 100
    var v5 uint16 = 675
    v6 := make(map[int]uint32)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v5 = v5
    fmt.Println(v2)
    if true {
        if false {
            fmt.Println(v3)
            var v20 bool = ((^(v5)) != v5)
            _ = v20
        } else {
            v21 := GenS2{F0: 1000, F1: int8(v2), F2: (v2 ^ v2)}
            if (328 != 42) {
                if (!(!true)) {
                    v22 := v6[int(v4)]
                    var v23 int64 = 446
                    v23 -= v23
                    _ = v22
                    _ = v23
                } else {
                    fmt.Println(v4)
                    v5 = v5
                }
                fmt.Println(v5, v4, v4)
            }
            _ = v21
        }
    } else {
        fmt.Println(v5, v5, v4)
        for fi24 := 0; fi24 < 17; fi24++ {
            if (0 < fi24) {
                if (!(109 != 2)) {
                    var v25 float64 = (1.5 + (-2.25 / (-2.25 * float64((1.5 / 1.0)))))
                    _ = len(v6)
                    v26 := GenS3{F0: 1000, F1: 191}
                    _ = v25
                    _ = v26
                } else {
                    switch (fi24 + 316) {
                    case 32768:
                        for fi27 := 0; fi27 < 7; fi27++ {
                            if false {
                                switch (fi24 / (((fi24 * -634)) | 1)) {
                                case 127:
                                    var v28 int32 = (v2 | v2)
                                    _ = v28
                                case -1:
                                    fmt.Println(fi24, v2, fi27)
                                default:
                                    if len(v1) > 0 {
                                        v1[0] = 42
                                    }
                                }
                                switch (fi24 ^ fi27) {
                                case -670:
                                    if (genFn1("a", 1) != 9223372036854775807) {
                                        v29 := GenS1{F0: uint64(v5), F1: 0}
                                        v2 -= 42
                                        _ = v29
                                    }
                                case 60:
                                    switch (fi24 - -4) {
                                    case 2147483647:
                                        if true {
                                            v2 -= (v2 << ((0) & 63))
                                            v30 := v6[(fi24 | (fi24 >> ((298) & 63)))]
                                            _ = v30
                                        }
                                    case -9223372036854775808:
                                        var v31 uint32 = 42
                                        _ = v31
                                    }
                                case -307:
                                    fmt.Println(v2, v4)
                                default:
                                    v2 -= v2
                                }
                            }
                            _ = fi27
                        }
                    case -9223372036854775808:
                        for fi32 := 0; fi32 < 7; fi32++ {
                            var v33 int8 = -58
                            switch v3 {
                            case "\n\t":
                                v34 := GenS2{F0: 0, F1: 1, F2: (-(v2))}
                                _ = v34
                            case "a":
                                var v35 uint32 = 4294967295
                                _ = v35
                            default:
                                v36 := GenS4{F0: -262, F1: (v2 + -769)}
                                _ = v36
                            }
                            _ = fi32
                            _ = v33
                        }
                    case 0:
                        var v37 int32 = v2
                        _ = v37
                    default:
                        var v38 int64 = 63
                        _ = v38
                    }
                }
                switch v3 {
                case "foo bar":
                    v4 = (v4 | -27)
                case "":
                    if false {
                        v3 = v3
                        for fi39 := 0; fi39 < 3; fi39++ {
                            fmt.Println(v2)
                            var v40 int16 = 902
                            var v41 bool = false
                            _ = fi39
                            _ = v40
                            _ = v41
                        }
                        fmt.Println(v5)
                    }
                case "hello":
                    for fi42 := 0; fi42 < 15; fi42++ {
                        if (!(!(true && true))) {
                            var v43 int8 = v4
                            _ = v43
                        } else {
                            if (!true) {
                                v3 = ("\n\t" + ((v3 + "hello") + ("a" + v3)))
                                v5 -= v5
                                v4 = v4
                                switch (-(fi42)) {
                                case 128:
                                    v2 += (v2 >> ((4294967295) & 63))
                                default:
                                    if (!false) {
                                        var v44 float32 = (1e10 - (1e10 * 3.14159))
                                        _ = v44
                                    }
                                }
                            }
                            v3 = v3
                        }
                        _ = fi42
                    }
                default:
                    var v45 uint64 = 18446744073709551615
                    _ = v45
                }
            } else {
                if (684 > 65535) {
                    if true {
                        v46 := v6[(fi24 & (^(fi24)))]
                        var v47 float64 = 0.1
                        _ = v46
                        _ = v47
                    }
                    if true {
                        if false {
                            var v48 float32 = 3.14159
                            var v49 int64 = -1
                            _ = v48
                            _ = v49
                        }
                        fmt.Println(v2)
                    }
                }
                fmt.Println(v3)
                var v50 float32 = 3.14159
                _ = v50
            }
            _ = fi24
        }
    }
    var v51 uint8 = 127
    var v52 uint8 = v51
    _ = len(v1)
    var v53 string = ("\n\t" + (v3 + v3))
    v3 = "hello"
    if true {
        fmt.Println(v51)
        v3 = v3
        switch (v53 + (v53 + v53)) {
        case "":
            fmt.Println(v4)
        case "hello":
            switch -9223372036854775808 {
            case 127:
                var v54 uint64 = 685
                _ = v54
            }
        default:
            var v55 float32 = (1e10 * (3.14159 + -2.25))
            _ = v55
        }
    }
    v5 -= 108
    switch int(v51) {
    case 2147483647:
        var v56 int16 = 254
        _ = v56
    case -2147483648:
        v51 += (v52 << ((18446744073709551615) & 63))
    case -1:
        switch 42 {
        case -985:
            v57 := GenS5{F0: 624, F1: 255}
            _ = v57
        case 2147483647:
            v58 := GenS4{F0: 1, F1: (v2 ^ (^(v2)))}
            _ = v58
        default:
            v59 := GenS5{F0: 0, F1: 0}
            _ = v59
        }
    default:
        if false {
            for fi60 := 0; fi60 < 12; fi60++ {
                var v61 uint32 = 118
                _ = fi60
                _ = v61
            }
        } else {
            for fi62 := 0; fi62 < 14; fi62++ {
                v4 = v4
                v52 += 1
                _ = fi62
            }
            var v63 float32 = ((float32(0.0) * (float32(3.14159) + 0.0)) - ((1e10 + 1.5) + 0.0))
            _ = v63
        }
    }
    var v64 uint16 = v5
    v52 -= v52
    var v65 int = -470
    var v66 uint16 = 11
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v64:", v64)
    fmt.Println("v65:", v65)
    fmt.Println("v66:", v66)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v51
    _ = v52
    _ = v53
    _ = v6
    _ = v64
    _ = v65
    _ = v66
}
