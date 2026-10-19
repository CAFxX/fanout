package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint16; F1 uint64; F2 int64 }

func main() {
    var v1 string = "hello"
    v2 := make(map[string]bool)
    var v3 int = 1
    v4 := make(map[uint32]uint32)
    v5 := make(map[int]int)
    v6 := []string{((v1 + v1) + "a"), "hello"}
    fmt.Println(v3, v1)
    _ = len(v6)
    _ = len(v4)
    fmt.Println(v1, v3)
    v3 = (v3 << ((1) & 63))
    var v7 uint = 18446744073709551615
    fmt.Println(v1)
    if true {
        switch v3 {
        case 808:
            switch (v1 + "foo bar") {
            case "hello":
                _ = len(v4)
            case "":
                v1 = ((v1 + "hello") + v1)
            case "\n\t":
                for fi8 := 0; fi8 < 12; fi8++ {
                    if false {
                        v3 -= (fi8 - 9223372036854775807)
                        for fi9 := 0; fi9 < 13; fi9++ {
                            v3 -= (v3 << ((v7) & 63))
                            fmt.Println(fi9, v3, fi8)
                            _ = fi9
                        }
                        v10 := GenS1{F0: 0, F1: v7}
                        _ = v10
                    } else {
                        switch v1 {
                        case "a":
                            v3 -= (fi8 / (((v3 | (v3 % ((782) | 1)))) | 1))
                        case "":
                            v1 = v1
                        case "\n\t":
                            for fi11 := 0; fi11 < 12; fi11++ {
                                if (true || false) {
                                    switch "" {
                                    case "foo bar":
                                        switch (v3 / (((fi8 | (v3 << ((v7) & 63)))) | 1)) {
                                        case 127:
                                            fmt.Println(v1, fi8)
                                        case 121:
                                            fmt.Println(v7, v3, v1)
                                        default:
                                            if len(v6) >= 2 {
                                                v6 = v6[1:len(v6)]
                                            }
                                        }
                                    default:
                                        if false {
                                            var v12 int16 = 42
                                            _ = v12
                                        } else {
                                            var v13 float32 = 3.14159
                                            v14 := GenS1{F0: 0, F1: (v7 & v7)}
                                            v7 = 625
                                            _ = v13
                                            _ = v14
                                        }
                                    }
                                    var v15 int16 = 42
                                    v7 = 546
                                    _ = v15
                                } else {
                                    v3 = (fi8 * 4294967295)
                                    v1 = ("\n\t" + "a")
                                    v16 := GenS2{F0: 1}
                                    _ = v16
                                }
                                fmt.Println(v3, v3)
                                _ = fi11
                            }
                        }
                    }
                    fmt.Println(v1, v7)
                    _ = fi8
                }
            }
        case 272:
            v17 := GenS2{F0: 32767}
            _ = v17
        }
    } else {
        v3 -= (v3 | 128)
        for fi18 := 0; fi18 < 6; fi18++ {
            v2[("foo bar" + ((v1 + v1) + (v1 + (v1 + v1))))] = true
            v7 -= (v7 % ((v7) | 1))
            if ((1.5 != 0.1) || false) {
                v19 := GenS2{F0: 698}
                if (false || false) {
                    v3 = (^(fi18))
                } else {
                    for fi20 := 0; fi20 < 4; fi20++ {
                        switch (v1 + (("" + v1) + v1)) {
                        case "":
                            var v21 int32 = 614
                            _ = v21
                        case "\n\t":
                            var v22 float32 = 1e10
                            _ = v22
                        default:
                            var v23 int64 = -9223372036854775808
                            _ = v23
                        }
                        _ = fi20
                    }
                    fmt.Println(fi18, v3)
                }
                _ = v19
            } else {
                var v24 bool = false
                v6 = append(v6, v1)
                fmt.Println(v1, fi18)
                _ = v24
            }
            _ = fi18
        }
        if false {
            for fi25 := 0; fi25 < 9; fi25++ {
                var v26 uint32 = 0
                _ = fi25
                _ = v26
            }
            v27 := GenS3{F0: 1}
            _ = len(v5)
            _ = v27
        } else {
            var v28 int32 = 2147483647
            if true {
                v1 = v1
                v3 += v3
            }
            _ = v28
        }
    }
    if true {
        fmt.Println(v7, v3, v1)
        v3 = v3
        v29 := GenS4{F0: 2, F1: 698, F2: 0}
        _ = v29
    }
    for fi30 := 0; fi30 < 7; fi30++ {
        for fi31 := 0; fi31 < 11; fi31++ {
            switch v3 {
            case -915:
                if true {
                    fmt.Println(v7)
                    if ((2 == 891) && false) {
                        var v32 uint16 = 65535
                        var v33 bool = true
                        _ = v32
                        _ = v33
                    }
                    v34 := GenS3{F0: -540}
                    var v35 int8 = 42
                    _ = v34
                    _ = v35
                }
            default:
                var v36 int = fi31
                _ = v36
            }
            fmt.Println(v3, fi30)
            var v37 bool = true
            _ = fi31
            _ = v37
        }
        v1 = ("hello" + v1)
        _ = fi30
    }
    var v38 uint8 = 2
    var v39 int32 = -2147483648
    for fi40 := 0; fi40 < 12; fi40++ {
        v41 := GenS4{F0: 42, F1: 4294967295, F2: -1}
        _ = fi40
        _ = v41
    }
    if len(v6) > 0 {
        v6[0] = "a"
    }
    v1 = ""
    var v42 uint64 = 1
    v39 = int32(v3)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v42
    _ = v5
    _ = v6
    _ = v7
}
