package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16 }
type GenS2 struct { F0 bool; F1 bool; F2 uint16 }
type GenS3 struct { F0 int8; F1 uint8; F2 int32 }
type GenS4 struct { F0 int8; F1 uint8; F2 uint64 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 string; F1 uint32; F2 int8 }
type GenS8 struct { F0 uint32; F1 int8 }
type GenS9 struct { F0 int64 }
type GenS10 struct { F0 uint32; F1 int }
type GenS11 struct { F0 uint32; F1 int32; F2 int16 }

func main() {
    var v1 int = -379
    var v2 uint32 = 65535
    v3 := []uint32{(v2 / ((v2) | 1)), (v2 * (v2 - 65535)), (v2 % ((v2) | 1)), (v2 / (((^(v2))) | 1))}
    v4 := make(map[uint32]uint32)
    v5 := make(map[uint32]int)
    var v6 uint16 = 599
    var v7 string = "foo bar"
    fmt.Println(v6, v1)
    switch v1 {
    case 913:
        v2 = v2
    default:
        v3 = append(v3, (v2 - v2))
    }
    var v8 int16 = 216
    for fi9 := 0; fi9 < 14; fi9++ {
        fmt.Println(v6, fi9, v6)
        var v10 int32 = -1
        for fi11 := 0; fi11 < 20; fi11++ {
            if false {
                var v12 uint = 42
                v8 += -728
                fmt.Println(fi9)
                _ = v12
            }
            switch (v1 >> ((18446744073709551615) & 63)) {
            case -9223372036854775808:
                fmt.Println(v1)
                v1 -= (-(v1))
            }
            _ = fi11
        }
        _ = fi9
        _ = v10
    }
    if false {
        if false {
            var v13 bool = (42 >= v6)
            var v14 int32 = 42
            _ = v13
            _ = v14
        }
        for fi15 := 0; fi15 < 3; fi15++ {
            v16 := GenS1{F0: -591, F1: 42}
            v8 = -1
            var v17 int16 = (v8 % ((-1) | 1))
            _ = fi15
            _ = v16
            _ = v17
        }
        var v18 string = ("\n\t" + v7)
        _ = v18
    }
    fmt.Println(v6, v6, v6)
    switch v7 {
    case "":
        v19 := GenS2{F0: true, F1: false, F2: v6}
        _ = v19
    case "hello":
        v5[837] = (v1 << ((42) & 63))
    case "foo bar":
        fmt.Println(v6)
    default:
        v1 -= (^(v1))
    }
    for fi20 := 0; fi20 < 9; fi20++ {
        for fi21 := 0; fi21 < 14; fi21++ {
            v22 := GenS1{F0: int16(v2), F1: (v8 * 217)}
            v23 := GenS3{F0: 100, F1: 127, F2: 0}
            if (false && false) {
                v24 := GenS2{F0: true, F1: true, F2: 744}
                if (true && (1 < -128)) {
                    var v25 uint64 = 42
                    if (v2 != 455) {
                        fmt.Println(fi20, v1, fi21)
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        fmt.Println(v25)
                    } else {
                        var v26 int64 = 1
                        v27 := GenS4{F0: -128, F1: uint8(v25), F2: v25}
                        fmt.Println(v8)
                        _ = v26
                        _ = v27
                    }
                    _ = len(v5)
                    if true {
                        delete(v4, (v2 * (^(v2))))
                        var v28 int64 = 420
                        v7 = v7
                        _ = v28
                    } else {
                        v7 = v7
                        var v29 float32 = 0.0
                        fmt.Println(fi20)
                        fmt.Println(v1, fi20, v1)
                        _ = v29
                    }
                    _ = len(v3)
                    v30 := GenS2{F0: false, F1: false, F2: (v6 << ((0) & 63))}
                    _ = v25
                    _ = v30
                } else {
                    switch v7 {
                    case "\n\t":
                        v2 = v2
                    case "a":
                        v5[(v2 + v2)] = fi20
                    case "":
                        v4[v2] = 4294967295
                    default:
                        if ((v6 == v6) && (42 <= 100)) {
                            v2 = (v2 - v2)
                        }
                    }
                    v1 = fi21
                    v4[(v2 >> ((255) & 63))] = v2
                }
                _ = len(v4)
                _ = len(v4)
                _ = v24
            } else {
                switch (((v7 + v7) + "hello") + v7) {
                case "":
                    var v31 uint = 721
                    _ = v31
                default:
                    switch 0 {
                    case 128:
                        v5[v2] = 510
                    }
                }
                v8 += (-(v8))
                switch (^(fi20)) {
                case 4294967295:
                    fmt.Println(fi21, v7)
                default:
                    if (true && false) {
                        switch "" {
                        case "\n\t":
                            var v32 uint16 = v6
                            _ = v32
                        case "hello":
                            switch ("hello" + "hello") {
                            case "\n\t":
                                v33 := v4[v2]
                                _ = v33
                            default:
                                fmt.Println(fi21, v2)
                            }
                        case "":
                            fmt.Println(fi20, v2)
                        default:
                            if (!(42 != 1)) {
                                v34 := GenS2{F0: (4294967295 == 486), F1: (v7 >= v7), F2: v6}
                                v35 := GenS5{F0: v8}
                                if false {
                                    if (1 < 2) {
                                        fmt.Println(v6, v8)
                                        v36 := GenS6{F0: 944}
                                        _ = v36
                                    }
                                    _ = len(v3)
                                    if ((42 != 316) && false) {
                                        var v37 int = (v1 << ((728) & 63))
                                        _ = v37
                                    } else {
                                        v38 := GenS7{F0: v7, F1: v2, F2: 127}
                                        v1 = (fi21 * 1)
                                        switch ("" + v7) {
                                        case "foo bar":
                                            v6 += 15
                                        default:
                                            switch v7 {
                                            case "":
                                                v39 := GenS6{F0: 385}
                                                _ = v39
                                            case "hello":
                                                if (0 == 0) {
                                                    if (!false) {
                                                        v5[v2] = -1
                                                        fmt.Println(v1, fi20)
                                                    }
                                                } else {
                                                    var v40 uint16 = v6
                                                    v8 -= (v8 & (v8 & v8))
                                                    _ = v40
                                                }
                                            case "a":
                                                v7 = (v7 + "\n\t")
                                            default:
                                                if true {
                                                    var v41 int32 = 0
                                                    if (978 == 2) {
                                                        if (320 > 4294967295) {
                                                            v2 += (v2 - 721)
                                                            var v42 int8 = 0
                                                            var v43 uint32 = (v2 % (((^(v2))) | 1))
                                                            fmt.Println(v2, v43)
                                                            _ = v42
                                                            _ = v43
                                                        }
                                                        delete(v4, (v2 << ((255) & 63)))
                                                    } else {
                                                        v44 := GenS8{F0: v2, F1: 62}
                                                        fmt.Println(fi20, v41)
                                                        _ = v44
                                                    }
                                                    v6 = v6
                                                    _ = v41
                                                }
                                            }
                                        }
                                        switch (v7 + v7) {
                                        case "a":
                                            if ((-1 <= 42) && ("hello" < (v7 + "foo bar"))) {
                                                if len(v3) > 0 {
                                                    v3[0] = v2
                                                }
                                                v45 := GenS9{F0: 1}
                                                _ = v45
                                            }
                                        case "foo bar":
                                            v46 := GenS6{F0: 42}
                                            _ = v46
                                        case "\n\t":
                                            var v47 uint64 = 2
                                            _ = v47
                                        default:
                                            v48 := v4[v2]
                                            _ = v48
                                        }
                                        v49 := GenS10{F0: 42, F1: -2147483648}
                                        _ = v38
                                        _ = v49
                                    }
                                } else {
                                    v50 := GenS4{F0: 1, F1: 0, F2: 2}
                                    v3 = append(v3, v2)
                                    _ = v50
                                }
                                _ = v34
                                _ = v35
                            } else {
                                var v51 uint64 = 2
                                _ = len(v4)
                                var v52 string = (v7 + v7)
                                _ = v51
                                _ = v52
                            }
                        }
                    } else {
                        v8 += v8
                        fmt.Println(v1)
                        var v53 uint32 = (v2 / (((v2 / (((v2 | (v2 % ((v2) | 1)))) | 1))) | 1))
                        _ = v53
                    }
                }
            }
            _ = fi21
            _ = v22
            _ = v23
        }
        _ = fi20
    }
    v54 := GenS5{F0: -32768}
    if len(v3) > 0 {
        v3[0] = (^(v2))
    }
    v55 := GenS11{F0: (v2 >> ((2) & 63)), F1: -891, F2: 0}
    var v56 uint32 = (v2 ^ (v2 + v2))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v56:", v56)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v54
    _ = v55
    _ = v56
    _ = v6
    _ = v7
    _ = v8
}
