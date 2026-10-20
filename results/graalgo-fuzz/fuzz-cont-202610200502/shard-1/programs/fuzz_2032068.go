package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 string; F1 uint64; F2 uint8 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int; F1 int32 }
type GenS6 struct { F0 int32; F1 uint8 }
type GenS7 struct { F0 int16; F1 uint16 }
type GenS8 struct { F0 int16; F1 int8 }
func genFn0() uint8 {
    v5 := GenS1{F0: 127}
    v6 := GenS2{F0: "foo bar", F1: 670, F2: 127}
    var v7 uint8 = 85
    _ = v5
    _ = v6
    _ = v7
    return (v7 - (^(v7)))
}
func genFn1() uint {
    if false {
        var v8 int16 = 286
        _ = v8
    } else {
        for fi9 := 0; fi9 < 14; fi9++ {
            var v10 int8 = 29
            var v11 int = fi9
            _ = fi9
            _ = v10
            _ = v11
        }
    }
    var v12 int64 = 1
    v12 -= 1
    v12 = (v12 * v12)
    v13 := GenS2{F0: ("a" + "\n\t"), F1: 0, F2: 2}
    _ = v12
    _ = v13
    return 0
}

func main() {
    v1 := []uint32{}
    var v2 uint32 = 0
    var v3 float64 = 0.1
    var v4 int32 = 0
    v14 := GenS3{F0: v2}
    switch 561 {
    case 2147483647:
        fmt.Println(v3)
    case 707:
        switch "" {
        case "hello":
            _ = len(v1)
        default:
            fmt.Println(v3, v4)
        }
    case -9223372036854775808:
        for fi15 := 0; fi15 < 6; fi15++ {
            for fi16 := 0; fi16 < 3; fi16++ {
                var v17 uint8 = 255
                var v18 float64 = -2.25
                _ = fi16
                _ = v17
                _ = v18
            }
            v2 = v2
            switch "\n\t" {
            case "":
                v2 = v2
            default:
                v2 = v2
            }
            _ = fi15
        }
    default:
        var v19 float32 = 1.5
        _ = v19
    }
    var v20 uint16 = 42
    var v21 uint = 2
    switch "\n\t" {
    case "foo bar":
        var v22 bool = false
        _ = v22
    }
    var v23 uint32 = v2
    var v24 uint = v21
    var v25 string = "a"
    switch (v25 + (v25 + ("foo bar" + ""))) {
    case "a":
        for fi26 := 0; fi26 < 14; fi26++ {
            v3 += (((v3 * v3) + v3) - (v3 / v3))
            var v27 int8 = -110
            for fi28 := 0; fi28 < 4; fi28++ {
                switch (-(fi28)) {
                case -881:
                    v29 := GenS2{F0: v25, F1: 1, F2: 127}
                    _ = v29
                default:
                    switch v25 {
                    case "":
                        v1 = append(v1, 2147483647)
                    case "\n\t":
                        v30 := GenS4{F0: v27}
                        var v31 float64 = (0.0 / -2.25)
                        var v32 uint32 = 291
                        _ = v30
                        _ = v31
                        _ = v32
                    default:
                        switch fi28 {
                        case 532:
                            var v33 int = (fi28 % ((fi28) | 1))
                            var v34 uint32 = (v23 << ((255) & 63))
                            if len(v1) > 0 {
                                v1[0] = v2
                            }
                            _ = v33
                            _ = v34
                        }
                    }
                }
                v1 = append(v1, (v2 + (^(v23))))
                _ = fi28
            }
            _ = fi26
            _ = v27
        }
    }
    fmt.Println(v2, v21)
    for fi35 := 0; fi35 < 4; fi35++ {
        v2 += 0
        fmt.Println(v24, v3, fi35)
        _ = fi35
    }
    if (v21 < 255) {
        if (v21 < v21) {
            var v36 int32 = v4
            v21 = 4294967295
            _ = v36
        }
    }
    for fi37 := 0; fi37 < 19; fi37++ {
        var v38 uint32 = (^(v2))
        if ((v20 ^ v20) > (v20 - v20)) {
            switch v25 {
            case "a":
                if (!("" <= "foo bar")) {
                    var v39 uint = v24
                    v25 = ""
                    switch v25 {
                    case "hello":
                        if (1e10 == (-2.25 - 1e10)) {
                            fmt.Println(v24, fi37, v21)
                        } else {
                            v40 := GenS4{F0: -128}
                            var v41 string = "\n\t"
                            _ = v40
                            _ = v41
                        }
                    }
                    _ = v39
                } else {
                    v25 = "hello"
                    switch "\n\t" {
                    case "":
                        fmt.Println(v4, v23, v21)
                    case "foo bar":
                        v23 -= v23
                    case "hello":
                        var v42 float32 = 0.0
                        _ = v42
                    default:
                        v4 += (v4 * v4)
                    }
                    if ((^(v4)) > v4) {
                        var v43 int16 = -1
                        v2 = (v38 % ((42) | 1))
                        v43 = int16(v20)
                        _ = v43
                    }
                }
            case "foo bar":
                if len(v1) > 0 {
                    v1[0] = (v2 - 2)
                }
            default:
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
            for fi44 := 0; fi44 < 4; fi44++ {
                v45 := GenS3{F0: 65535}
                _ = len(v1)
                _ = fi44
                _ = v45
            }
            if len(v1) > 0 {
                v1[0] = (v2 % (((v23 << ((18446744073709551615) & 63))) | 1))
            }
        }
        _ = fi37
        _ = v38
    }
    v1 = append(v1, v23)
    var v46 float64 = (-2.25 - ((v3 / v3) - 1.5))
    v23 = 0
    v24 += v24
    if false {
        for fi47 := 0; fi47 < 8; fi47++ {
            v25 = (v25 + v25)
            if true {
                if true {
                    if true {
                        var v48 uint16 = (^(v20))
                        if (false && false) {
                            v4 = v4
                            v49 := GenS4{F0: 0}
                            _ = v49
                        }
                        v50 := GenS4{F0: 42}
                        _ = v48
                        _ = v50
                    } else {
                        var v51 int64 = 0
                        for fi52 := 0; fi52 < 8; fi52++ {
                            var v53 int64 = (v51 * v51)
                            if (false || true) {
                                _ = len(v1)
                                v54 := GenS2{F0: "foo bar", F1: 598, F2: 150}
                                if false {
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    v55 := GenS3{F0: v23}
                                    var v56 uint64 = 18446744073709551615
                                    _ = v55
                                    _ = v56
                                } else {
                                    fmt.Println(v25, v53, fi52)
                                    if (3.14159 > v3) {
                                        v57 := GenS5{F0: fi47, F1: v4}
                                        _ = v57
                                    } else {
                                        v1 = append(v1, (v23 + v2))
                                    }
                                }
                                _ = v54
                            }
                            _ = fi52
                            _ = v53
                        }
                        _ = v51
                    }
                    v58 := GenS6{F0: v4, F1: 42}
                    _ = v58
                } else {
                    if (!((0.0 > 0.0) || true)) {
                        var v59 int64 = 1
                        var v60 int16 = 1000
                        for fi61 := 0; fi61 < 2; fi61++ {
                            v62 := GenS7{F0: 32767, F1: 32767}
                            if (((fi47 / ((fi47) | 1)) < fi47) && true) {
                                v63 := GenS8{F0: (v60 - v60), F1: 1}
                                var v64 uint = (v24 % ((v24) | 1))
                                fmt.Println(v60, v20, v4)
                                _ = v63
                                _ = v64
                            } else {
                                var v65 uint32 = 1
                                _ = v65
                            }
                            _ = fi61
                            _ = v62
                        }
                        _ = v59
                        _ = v60
                    }
                    fmt.Println(v46, v3)
                }
                var v66 int8 = 100
                _ = v66
            }
            var v67 int64 = -156
            _ = fi47
            _ = v67
        }
        for fi68 := 0; fi68 < 19; fi68++ {
            var v69 bool = (true || true)
            _ = fi68
            _ = v69
        }
        if ((1 == -926) || (!(uint64(v20) > 253))) {
            var v70 uint = uint(v2)
            _ = v70
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    _ = v1
    _ = v14
    _ = v2
    _ = v20
    _ = v21
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v46
}
