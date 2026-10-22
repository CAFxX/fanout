package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int8; F1 int32; F2 int16 }
type GenS4 struct { F0 uint8; F1 bool; F2 int32 }

func main() {
    v1 := make(map[uint32]int)
    v2 := []uint32{42, 1}
    var v3 uint64 = 1
    v4 := []string{(("foo bar" + "\n\t") + ("foo bar" + "hello"))}
    var v5 float32 = (float32(-2.25) / 1.0)
    var v6 int32 = 1
    var v7 uint = 887
    for fi8 := 0; fi8 < 11; fi8++ {
        if (uint16(v7) < 32767) {
            _ = len(v1)
        } else {
            if false {
                for fi9 := 0; fi9 < 14; fi9++ {
                    if false {
                        if true {
                            var v10 bool = (47 >= 255)
                            _ = v10
                        } else {
                            v5 = v5
                            if false {
                                switch "\n\t" {
                                case "hello":
                                    var v11 int64 = 1
                                    _ = v11
                                case "\n\t":
                                    var v12 int16 = 1000
                                    _ = v12
                                case "a":
                                    fmt.Println(v6)
                                default:
                                    switch "hello" {
                                    case "a":
                                        v13 := v1[65535]
                                        _ = v13
                                    default:
                                        switch "\n\t" {
                                        case "foo bar":
                                            _ = len(v1)
                                        case "a":
                                            fmt.Println(v5, v7)
                                        default:
                                            var v14 uint64 = 854
                                            _ = v14
                                        }
                                    }
                                }
                            } else {
                                v2 = append(v2, 2147483647)
                                fmt.Println(v5, v6, fi8)
                                delete(v1, 170)
                            }
                            if true {
                                var v15 uint64 = (^(v3))
                                _ = v15
                            }
                            v5 = 1.5
                            if true {
                                v6 -= (v6 - v6)
                            } else {
                                var v16 int8 = int8(v6)
                                if (false && ((false && false) && (!false))) {
                                    var v17 uint8 = 85
                                    var v18 uint64 = 341
                                    v19 := GenS1{F0: 65535, F1: 42}
                                    _ = v17
                                    _ = v18
                                    _ = v19
                                }
                                var v20 int = fi8
                                _ = v16
                                _ = v20
                            }
                        }
                    } else {
                        _ = len(v1)
                    }
                    var v21 float32 = float32(1.5)
                    _ = fi9
                    _ = v21
                }
                var v22 string = "hello"
                v23 := GenS2{F0: 42}
                _ = v22
                _ = v23
            } else {
                v6 = v6
                for fi24 := 0; fi24 < 12; fi24++ {
                    var v25 float32 = (v5 * (v5 + v5))
                    fmt.Println(fi8, v6, fi8)
                    _ = fi24
                    _ = v25
                }
            }
            v7 -= (v7 - (v7 ^ (^(v7))))
        }
        for fi26 := 0; fi26 < 18; fi26++ {
            switch "hello" {
            case "":
                v5 = v5
            default:
                var v27 int8 = 1
                _ = v27
            }
            fmt.Println(v5)
            var v28 bool = true
            _ = fi26
            _ = v28
        }
        _ = fi8
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    for fi29 := 0; fi29 < 8; fi29++ {
        var v30 uint8 = 1
        switch int(v7) {
        case 32768:
            var v31 uint8 = v30
            _ = v31
        case -1:
            v30 -= (v30 | v30)
        case 32767:
            fmt.Println(v5)
        }
        var v32 int64 = 962
        _ = fi29
        _ = v30
        _ = v32
    }
    var v33 uint16 = 442
    for fi34 := 0; fi34 < 2; fi34++ {
        if len(v4) > 0 {
            v4[0] = "\n\t"
        }
        for fi35 := 0; fi35 < 11; fi35++ {
            _ = len(v4)
            if true {
                if ((true && true) && false) {
                    fmt.Println(fi35, v7)
                    if len(v4) > 0 {
                        v4[0] = "a"
                    }
                    if len(v4) > 0 {
                        v4[0] = ""
                    }
                } else {
                    v5 += v5
                    if len(v4) > 0 {
                        v4[0] = ((("" + "hello") + "\n\t") + "")
                    }
                }
                if true {
                    v33 += v33
                    switch (fi34 & fi35) {
                    case 128:
                        var v36 float32 = v5
                        _ = v36
                    case 42:
                        switch (fi34 * fi34) {
                        case 706:
                            switch "foo bar" {
                            case "\n\t":
                                var v37 uint32 = 928
                                _ = v37
                            case "a":
                                v38 := v1[262]
                                _ = v38
                            case "foo bar":
                                if false {
                                    var v39 uint16 = (v33 ^ v33)
                                    v40 := GenS3{F0: -1, F1: 4, F2: 1}
                                    v6 -= v6
                                    _ = v39
                                    _ = v40
                                } else {
                                    if true {
                                        v5 = 0.1
                                    }
                                }
                            default:
                                v41 := GenS3{F0: 127, F1: v6, F2: 32767}
                                _ = v41
                            }
                        default:
                            if (!("\n\t" > "hello")) {
                                _ = len(v1)
                                fmt.Println(fi34)
                            } else {
                                if false {
                                    v42 := GenS2{F0: v33}
                                    var v43 int = fi34
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    _ = v42
                                    _ = v43
                                }
                            }
                        }
                    default:
                        v44 := GenS3{F0: 1, F1: 0, F2: -472}
                        _ = v44
                    }
                } else {
                    v45 := GenS2{F0: (v33 % ((v33) | 1))}
                    _ = len(v1)
                    v46 := GenS2{F0: 1}
                    _ = v45
                    _ = v46
                }
                _ = len(v1)
            } else {
                fmt.Println(fi34, v33)
                switch ("foo bar" + "foo bar") {
                case "":
                    v3 = (v3 >> (((v7 * (v7 * (v7 - 18446744073709551615)))) & 63))
                case "a":
                    v4 = append(v4, "a")
                default:
                    v3 = v3
                }
            }
            var v47 int = fi34
            _ = fi35
            _ = v47
        }
        v3 = (^(v3))
        _ = fi34
    }
    switch 811 {
    case -2147483648:
        v48 := GenS2{F0: 173}
        _ = v48
    case 477:
        v6 -= (v6 << ((v7) & 63))
    default:
        v49 := GenS4{F0: 127, F1: false, F2: 42}
        _ = v49
    }
    v50 := v1[2]
    for fi51 := 0; fi51 < 16; fi51++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        v33 = v33
        _ = fi51
    }
    for fi52 := 0; fi52 < 17; fi52++ {
        v53 := GenS1{F0: 0, F1: (v33 ^ (^(v33)))}
        _ = fi52
        _ = v53
    }
    fmt.Println(v7, v3)
    v6 += v6
    v5 = v5
    var v54 uint32 = 42
    var v55 uint = v7
    if true {
        v50 = (v50 * (v50 / ((v50) | 1)))
        var v56 uint8 = 1
        _ = v56
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v54:", v54)
    fmt.Println("v55:", v55)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v50
    _ = v54
    _ = v55
    _ = v6
    _ = v7
}
