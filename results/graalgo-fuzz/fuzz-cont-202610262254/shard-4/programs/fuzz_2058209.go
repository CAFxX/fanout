package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 bool; F1 uint; F2 int }
type GenS3 struct { F0 uint32; F1 bool; F2 int8 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 uint64; F1 uint16; F2 int64 }
type GenS7 struct { F0 int64; F1 int32 }
type GenS8 struct { F0 uint; F1 uint32; F2 uint32 }
func genFn0() string {
    var v7 uint = 358
    var v8 string = ""
    v8 = (("a" + (v8 + "hello")) + v8)
    var v9 int32 = 270
    if false {
        var v10 uint8 = 0
        if (!false) {
            v11 := []int{478, 108}
            fmt.Println(v8)
            v12 := GenS1{F0: ((((v8 + v8) + v8) + "\n\t") + v8)}
            _ = v11
            _ = v12
        } else {
            var v13 uint8 = (^(v10))
            _ = v13
        }
        v14 := GenS2{F0: (false || (!false)), F1: v7, F2: -2147483648}
        _ = v10
        _ = v14
    }
    _ = v7
    _ = v8
    _ = v9
    return ("\n\t" + ((("a" + v8) + v8) + "\n\t"))
}
func genFn1(p0 int32, p1 uint16) uint16 {
    for fi15 := 0; fi15 < 10; fi15++ {
        p0 = (p0 | (p0 >> ((65535) & 63)))
        _ = fi15
    }
    v16 := GenS1{F0: "foo bar"}
    p1 += 826
    _ = p0
    _ = p1
    _ = v16
    return (p1 * p1)
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[string]uint32)
    var v3 int = 32768
    v4 := []int{127}
    v5 := make(map[uint32]int)
    var v6 uint = 42
    v17 := GenS3{F0: 4294967295, F1: true, F2: 127}
    if true {
        fmt.Println(v3)
        for fi18 := 0; fi18 < 1; fi18++ {
            fmt.Println(fi18, fi18, fi18)
            if (-32768 < 32767) {
                v3 = v3
                if false {
                    var v19 bool = (1000 > 0)
                    switch 128 {
                    case -2147483648:
                        var v20 int32 = -382
                        _ = v20
                    case 128:
                        fmt.Println(v19, v19, v6)
                    case -9223372036854775808:
                        fmt.Println(fi18)
                    default:
                        v5[4294967295] = fi18
                    }
                    _ = v19
                } else {
                    v3 = (-(v3))
                    if (float64(3.14159) <= ((1.5 - 0.0) + (0.0 - 1e10))) {
                        fmt.Println(v3, v3, v3)
                        var v21 uint16 = 854
                        _ = v21
                    } else {
                        switch "" {
                        case "foo bar":
                            if true {
                                fmt.Println(fi18, v3)
                            } else {
                                v6 += v6
                                fmt.Println(fi18)
                            }
                        case "hello":
                            v6 -= 65535
                        case "\n\t":
                            fmt.Println(fi18, fi18, v3)
                        default:
                            v6 += 993
                        }
                    }
                    v22 := GenS4{F0: -1}
                    _ = v22
                }
            }
            v6 = v6
            _ = fi18
        }
    }
    for fi23 := 0; fi23 < 6; fi23++ {
        v24 := GenS2{F0: (2 == 614), F1: 109, F2: 813}
        _ = fi23
        _ = v24
    }
    v6 = 4294967295
    var v25 int = v3
    fmt.Println(v3, v25)
    var v26 int = v3
    switch v3 {
    case 0:
        fmt.Println(v26, v3)
    case 421:
        for fi27 := 0; fi27 < 9; fi27++ {
            var v28 float64 = 3.14159
            fmt.Println(v6)
            _ = fi27
            _ = v28
        }
    case -2147483648:
        v29 := GenS5{F0: -32768}
        _ = v29
    default:
        v3 = -554
    }
    v3 = v3
    for fi30 := 0; fi30 < 10; fi30++ {
        switch (("\n\t" + ("\n\t" + "hello")) + (((("\n\t" + "\n\t") + "a") + "\n\t") + "foo bar")) {
        case "hello":
            v6 = v6
        case "a":
            if len(v4) > 0 {
                v4[0] = (v3 & (fi30 % ((v26) | 1)))
            }
        case "":
            fmt.Println(v6)
        }
        v6 += v6
        _ = fi30
    }
    var v31 uint16 = 32767
    v3 = v25
    switch "hello" {
    case "\n\t":
        if len(v4) > 0 {
            v4[0] = (v25 * v3)
        }
    case "":
        var v32 float32 = ((-2.25 - -2.25) + (0.1 / (1e10 * ((1.5 * ((1.5 + 1e10) / 0.1)) * 1e10))))
        _ = v32
    case "foo bar":
        v25 -= v26
    default:
        if ((724 > v31) || (1 >= uint8(v6))) {
            v26 += v25
            if ((!(0 < 1)) && (!false)) {
                if false {
                    if (!((v31 / ((738) | 1)) < (v31 << ((uint(v31)) & 63)))) {
                        v26 = v26
                        for fi33 := 0; fi33 < 1; fi33++ {
                            var v34 float64 = 0.0
                            _ = fi33
                            _ = v34
                        }
                    }
                    v35 := GenS6{F0: 42, F1: (v31 * v31), F2: 1}
                    switch (^(v3)) {
                    case 1:
                        if true {
                            var v36 uint = (v6 + 42)
                            for fi37 := 0; fi37 < 6; fi37++ {
                                var v38 uint8 = 1
                                v6 += 267
                                fmt.Println(v38, v26)
                                _ = fi37
                                _ = v38
                            }
                            for fi39 := 0; fi39 < 13; fi39++ {
                                switch 42 {
                                case 32768:
                                    v40 := GenS1{F0: ("\n\t" + "foo bar")}
                                    _ = v40
                                case 127:
                                    if true {
                                        for fi41 := 0; fi41 < 3; fi41++ {
                                            if false {
                                                var v42 uint64 = 789
                                                _ = v42
                                            }
                                            var v43 int = (v26 / ((v25) | 1))
                                            v36 += v6
                                            _ = fi41
                                            _ = v43
                                        }
                                        switch "" {
                                        case "":
                                            if false {
                                                v44 := GenS7{F0: -289, F1: -2147483648}
                                                fmt.Println(v26, v25, v26)
                                                fmt.Println(v25)
                                                _ = v44
                                            }
                                        case "hello":
                                            var v45 int = fi39
                                            _ = v45
                                        case "a":
                                            var v46 string = ""
                                            _ = v46
                                        default:
                                            v47 := GenS8{F0: (v36 / ((v36) | 1)), F1: 635, F2: 178}
                                            _ = v47
                                        }
                                    }
                                case -2147483648:
                                    var v48 int = (-(v26))
                                    _ = v48
                                default:
                                    var v49 bool = false
                                    _ = v49
                                }
                                for fi50 := 0; fi50 < 4; fi50++ {
                                    var v51 int16 = 1
                                    fmt.Println(fi39)
                                    var v52 uint8 = 144
                                    if (int8(v52) >= 1) {
                                        v3 = (v3 >> ((365) & 63))
                                    } else {
                                        switch int(v6) {
                                        case 42:
                                            var v53 int8 = -1
                                            _ = v53
                                        default:
                                            var v54 float64 = 3.14159
                                            _ = v54
                                        }
                                        if len(v4) > 0 {
                                            v4[0] = fi50
                                        }
                                        v55 := GenS2{F0: false, F1: (v36 * (v36 + 402)), F2: fi50}
                                        _ = v55
                                    }
                                    var v56 uint64 = 4294967295
                                    _ = fi50
                                    _ = v51
                                    _ = v52
                                    _ = v56
                                }
                                switch "foo bar" {
                                case "foo bar":
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                case "\n\t":
                                    for fi57 := 0; fi57 < 19; fi57++ {
                                        v31 = 255
                                        _ = fi57
                                    }
                                case "hello":
                                    for fi58 := 0; fi58 < 4; fi58++ {
                                        _ = len(v1)
                                        _ = len(v5)
                                        v59 := GenS2{F0: false, F1: 924, F2: (fi58 | (v3 & 32767))}
                                        _ = fi58
                                        _ = v59
                                    }
                                }
                                _ = fi39
                            }
                            _ = v36
                        }
                    default:
                        switch v26 {
                        case 1:
                            v6 = (v6 >> ((v6) & 63))
                        case 128:
                            fmt.Println(v3, v3, v25)
                        }
                    }
                    _ = v35
                }
            }
            var v60 int64 = 71
            _ = v60
        }
    }
    v31 = 255
    fmt.Println(v3, v25)
    fmt.Println(v3, v31, v6)
    for fi61 := 0; fi61 < 18; fi61++ {
        for fi62 := 0; fi62 < 16; fi62++ {
            fmt.Println(v26, v31)
            var v63 int64 = -9223372036854775808
            if false {
                v63 -= (v63 << ((uint(v31)) & 63))
            } else {
                v6 = 42
                var v64 int16 = -526
                var v65 uint32 = uint32(v31)
                var v66 uint16 = v31
                _ = v64
                _ = v65
                _ = v66
            }
            var v67 int64 = (v63 ^ v63)
            if ((2 <= 0) || false) {
                v26 -= (v26 ^ (v3 | (v25 % ((fi61) | 1))))
                var v68 uint = (v6 >> (((v6 - 1)) & 63))
                _ = v68
            } else {
                fmt.Println(v3, fi61)
                if (-896 < 42) {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                }
                fmt.Println(fi62, v25)
            }
            _ = fi62
            _ = v63
            _ = v67
        }
        var v69 uint8 = 255
        var v70 int64 = -1
        _ = fi61
        _ = v69
        _ = v70
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
