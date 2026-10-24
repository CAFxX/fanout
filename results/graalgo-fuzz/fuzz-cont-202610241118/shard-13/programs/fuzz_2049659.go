package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32 }
type GenS2 struct { F0 uint32; F1 int32; F2 int32 }

func main() {
    v1 := []uint32{4294967295, 0}
    v2 := make(map[uint32]bool)
    var v3 int64 = 0
    v4 := make(map[uint32]bool)
    v5 := []uint32{42, 1, 2, 1}
    for fi6 := 0; fi6 < 19; fi6++ {
        if (255 > 717) {
            v3 = (-(v3))
            var v7 int32 = 997
            _ = v7
        }
        v3 -= 42
        _ = fi6
    }
    switch 1 {
    case 16:
        if (-353 > -2147483648) {
            v3 = -9223372036854775808
        } else {
            var v8 uint = 4294967295
            var v9 bool = (uint32(v8) == 2147483647)
            var v10 int64 = v3
            _ = v10
            _ = v8
            _ = v9
        }
    case 4294967295:
        fmt.Println(v3)
    case -710:
        v11 := GenS1{F0: 42, F1: 4294967295}
        _ = v11
    }
    v12 := GenS1{F0: 0, F1: 766}
    var v13 float32 = (0.1 - 0.0)
    fmt.Println(v3, v3)
    for fi14 := 0; fi14 < 12; fi14++ {
        v3 -= -1
        v1 = append(v1, 2147483647)
        _ = fi14
    }
    var v15 int32 = -1
    var v16 int = 127
    v16 -= v16
    fmt.Println(v3, v16)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v17 uint16 = 420
    switch "hello" {
    case "":
        v17 = v17
    case "\n\t":
        var v18 uint8 = 42
        _ = v18
    default:
        if ((264 < v3) || false) {
            if len(v1) > 0 {
                v1[0] = 2
            }
            switch v16 {
            case 42:
                for fi19 := 0; fi19 < 15; fi19++ {
                    v16 = v16
                    _ = len(v1)
                    _ = fi19
                }
            case -1:
                var v20 int64 = v3
                _ = v20
            case 676:
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            default:
                var v21 bool = true
                _ = v21
            }
            fmt.Println(v15, v15)
        } else {
            if true {
                if true {
                    switch ("foo bar" + "hello") {
                    case "":
                        fmt.Println(v15)
                    case "a":
                        var v22 uint8 = 255
                        _ = v22
                    case "foo bar":
                        switch "\n\t" {
                        case "":
                            fmt.Println(v17, v3)
                        case "foo bar":
                            for fi23 := 0; fi23 < 2; fi23++ {
                                v3 -= (^(v3))
                                _ = fi23
                            }
                        case "hello":
                            if true {
                                v5 = append(v5, 2)
                                fmt.Println(v16, v17, v15)
                            } else {
                                switch "\n\t" {
                                case "":
                                    fmt.Println(v17, v16)
                                default:
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                }
                            }
                        }
                    }
                    for fi24 := 0; fi24 < 11; fi24++ {
                        for fi25 := 0; fi25 < 17; fi25++ {
                            switch "foo bar" {
                            case "a":
                                fmt.Println(v17)
                            default:
                                if false {
                                    var v26 bool = true
                                    switch ("foo bar" + "\n\t") {
                                    case "a":
                                        v13 = (1.5 + (v13 - v13))
                                    case "":
                                        v27 := GenS1{F0: v15, F1: 2}
                                        _ = v27
                                    }
                                    _ = v26
                                } else {
                                    v3 -= -1
                                }
                            }
                            v13 += 3.14159
                            fmt.Println(v13, v15, v13)
                            _ = fi25
                        }
                        v28 := GenS2{F0: 0, F1: -1, F2: (v15 ^ v15)}
                        _ = fi24
                        _ = v28
                    }
                    var v29 uint = 47
                    _ = v29
                }
                v30 := GenS2{F0: 65535, F1: (v15 << ((1) & 63)), F2: (v15 + v15)}
                v31 := GenS1{F0: v15, F1: 1}
                _ = v30
                _ = v31
            }
            for fi32 := 0; fi32 < 6; fi32++ {
                if false {
                    if false {
                        delete(v2, 464)
                        v16 = v16
                        if (-32768 >= 42) {
                            v15 = (v15 & -623)
                        } else {
                            v33 := GenS2{F0: 994, F1: (-(v15)), F2: 42}
                            _ = v33
                        }
                    }
                    for fi34 := 0; fi34 < 12; fi34++ {
                        v3 -= (v3 | 42)
                        _ = len(v4)
                        v13 = (v13 / 1e10)
                        _ = fi34
                    }
                }
                _ = fi32
            }
            v3 = v3
        }
    }
    v4[2147483648] = true
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
