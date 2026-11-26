package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint; F1 int16; F2 int8 }
type GenS3 struct { F0 uint16; F1 int32; F2 int8 }
type GenS4 struct { F0 int64; F1 bool; F2 uint8 }
type GenS5 struct { F0 int64; F1 uint32; F2 int8 }
type GenS6 struct { F0 uint16; F1 uint32; F2 string }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 string; F1 uint16; F2 int }
type GenS9 struct { F0 int16; F1 uint }

func main() {
    v1 := make(map[int]string)
    v2 := make(map[int]string)
    var v3 int16 = -32768
    var v4 int32 = -81
    v5 := make(map[int]bool)
    var v6 uint32 = 834
    var v7 uint64 = 4294967295
    fmt.Println(v3)
    for fi8 := 0; fi8 < 20; fi8++ {
        var v9 uint64 = 0
        switch fi8 {
        case -379:
            v10 := GenS1{F0: -1}
            _ = v10
        default:
            var v11 int16 = (-(v3))
            _ = v11
        }
        _ = fi8
        _ = v9
    }
    v12 := []int{536, 127, 9223372036854775807, 85}
    v7 = v7
    if false {
        v7 = 821
        var v13 string = "foo bar"
        _ = v13
    } else {
        switch ("a" + "foo bar") {
        case "\n\t":
            if (!true) {
                for fi14 := 0; fi14 < 8; fi14++ {
                    for fi15 := 0; fi15 < 10; fi15++ {
                        var v16 int8 = 42
                        var v17 bool = (!false)
                        _ = fi15
                        _ = v16
                        _ = v17
                    }
                    for fi18 := 0; fi18 < 6; fi18++ {
                        var v19 float32 = -2.25
                        _ = fi18
                        _ = v19
                    }
                    var v20 int64 = 748
                    _ = fi14
                    _ = v20
                }
                _ = len(v12)
                var v21 uint8 = 0
                _ = v21
            } else {
                var v22 int64 = -1
                v6 = (v6 + 592)
                switch "\n\t" {
                case "hello":
                    if false {
                        if (63 > 0) {
                            for fi23 := 0; fi23 < 8; fi23++ {
                                var v24 uint = 647
                                v24 = (v24 << ((v24) & 63))
                                _ = fi23
                                _ = v24
                            }
                            for fi25 := 0; fi25 < 19; fi25++ {
                                v26 := GenS2{F0: 255, F1: int16(v6), F2: -57}
                                _ = fi25
                                _ = v26
                            }
                        } else {
                            if (18446744073709551615 > 4294967295) {
                                v27 := GenS3{F0: 0, F1: v4, F2: 115}
                                _ = v27
                            }
                        }
                        if true {
                            if true {
                                for fi28 := 0; fi28 < 7; fi28++ {
                                    v29 := GenS3{F0: 1, F1: (v4 - (v4 % ((v4) | 1))), F2: 0}
                                    _ = fi28
                                    _ = v29
                                }
                            } else {
                                v30 := GenS2{F0: 271, F1: v3, F2: 127}
                                v22 += 900
                                _ = v30
                            }
                            var v31 int8 = -97
                            v7 = v7
                            _ = v31
                        }
                        v32 := GenS4{F0: v22, F1: false, F2: 1}
                        _ = v32
                    } else {
                        if true {
                            var v33 float32 = 0.1
                            _ = v33
                        } else {
                            v4 = (v4 + (v4 << ((65535) & 63)))
                        }
                        for fi34 := 0; fi34 < 4; fi34++ {
                            var v35 string = "a"
                            var v36 uint8 = 128
                            var v37 string = v35
                            _ = fi34
                            _ = v35
                            _ = v36
                            _ = v37
                        }
                        for fi38 := 0; fi38 < 3; fi38++ {
                            v39 := GenS5{F0: -326, F1: (v6 % ((2147483647) | 1)), F2: -56}
                            _ = fi38
                            _ = v39
                        }
                    }
                case "":
                    var v40 float32 = ((-2.25 + 1e10) + 1e10)
                    _ = v40
                case "\n\t":
                    for fi41 := 0; fi41 < 13; fi41++ {
                        fmt.Println(v4)
                        v22 -= v22
                        v22 = v22
                        _ = fi41
                    }
                default:
                    var v42 bool = false
                    _ = v42
                }
                _ = v22
            }
        }
    }
    v43 := GenS6{F0: 0, F1: v6, F2: "a"}
    var v44 int8 = 0
    _ = len(v12)
    if true {
        if (!(240 <= 176)) {
            v45 := GenS5{F0: 556, F1: (v6 / ((v6) | 1)), F2: (v44 ^ (v44 | v44))}
            if (!true) {
                v6 += 4294967295
                v46 := GenS7{F0: uint16(v7)}
                _ = v46
            } else {
                var v47 uint16 = 1
                var v48 uint16 = (^(v47))
                if false {
                    for fi49 := 0; fi49 < 17; fi49++ {
                        v44 = v44
                        switch 32768 {
                        case 127:
                            if true {
                                var v50 uint16 = v47
                                if ((false && false) || (true || true)) {
                                    fmt.Println(fi49)
                                    fmt.Println(v3, v48)
                                    for fi51 := 0; fi51 < 16; fi51++ {
                                        v47 -= (^(v47))
                                        _ = fi51
                                    }
                                }
                                _ = v50
                            } else {
                                switch "foo bar" {
                                case "a":
                                    if (182 > uint8(v6)) {
                                        if true {
                                            fmt.Println(v48, v6, fi49)
                                            if (false && (42 == 0)) {
                                                fmt.Println(v4)
                                            }
                                            if false {
                                                if (v3 > 32767) {
                                                    for fi52 := 0; fi52 < 12; fi52++ {
                                                        _ = len(v12)
                                                        _ = fi52
                                                    }
                                                    if len(v12) >= 2 {
                                                        v12 = v12[1:len(v12)]
                                                    }
                                                }
                                                var v53 bool = ((v6 <= 42) && true)
                                                _ = v53
                                            } else {
                                                switch (fi49 + (fi49 & 127)) {
                                                case 32767:
                                                    for fi54 := 0; fi54 < 4; fi54++ {
                                                        v55 := GenS8{F0: "foo bar", F1: 692, F2: fi54}
                                                        _ = fi54
                                                        _ = v55
                                                    }
                                                }
                                                if len(v12) >= 2 {
                                                    v12 = v12[1:len(v12)]
                                                }
                                                v56 := GenS5{F0: 266, F1: (v6 % ((v6) | 1)), F2: int8(v6)}
                                                _ = v56
                                            }
                                        }
                                        v3 += (v3 ^ v3)
                                        v6 += v6
                                    }
                                default:
                                    for fi57 := 0; fi57 < 2; fi57++ {
                                        var v58 string = ""
                                        _ = fi57
                                        _ = v58
                                    }
                                }
                            }
                        }
                        var v59 int32 = v4
                        _ = fi49
                        _ = v59
                    }
                    v6 -= v6
                }
                _ = v47
                _ = v48
            }
            var v60 uint32 = v6
            _ = v45
            _ = v60
        }
    } else {
        v6 = (v6 << ((0) & 63))
        _ = len(v1)
        v3 += -587
    }
    switch ("" + ("a" + "hello")) {
    case "\n\t":
        switch ("\n\t" + "") {
        case "":
            v3 -= 32767
        case "\n\t":
            for fi61 := 0; fi61 < 9; fi61++ {
                for fi62 := 0; fi62 < 2; fi62++ {
                    var v63 uint8 = 1
                    if false {
                        var v64 int = (fi62 / (((fi61 + fi61)) | 1))
                        _ = len(v12)
                        _ = v64
                    }
                    fmt.Println(v44, v44, v7)
                    _ = fi62
                    _ = v63
                }
                v3 = v3
                var v65 int64 = -1
                _ = fi61
                _ = v65
            }
        default:
            v12 = append(v12, -1)
        }
    case "hello":
        v66 := v5[-343]
        _ = v66
    case "a":
        fmt.Println(v6, v6)
    default:
        v3 = v3
    }
    if false {
        fmt.Println(v6)
    }
    switch "" {
    case "foo bar":
        v67 := v2[-1]
        _ = v67
    default:
        v44 -= -7
    }
    switch 1 {
    case -473:
        if (!false) {
            v68 := GenS9{F0: (^(v3)), F1: 257}
            for fi69 := 0; fi69 < 1; fi69++ {
                v12 = append(v12, fi69)
                if len(v12) >= 2 {
                    v12 = v12[1:len(v12)]
                }
                _ = fi69
            }
            if true {
                var v70 bool = (!true)
                v44 = int8(v6)
                var v71 int8 = (v44 * v44)
                _ = v70
                _ = v71
            } else {
                v7 = (^(v7))
            }
            _ = v68
        }
    case 127:
        v72 := v5[-9223372036854775808]
        _ = v72
    default:
        v6 -= (v6 >> ((42) & 63))
    }
    v7 -= v7
    if ((false || true) || ((v7 + v7) >= (^(v7)))) {
        v7 = (v7 + (v7 / ((v7) | 1)))
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v5
    _ = v6
    _ = v7
}
