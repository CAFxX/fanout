package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint32 }

func main() {
    v1 := []uint32{0, 42, 0}
    var v2 bool = false
    var v3 uint16 = 2
    var v4 bool = (false && v2)
    for fi5 := 0; fi5 < 19; fi5++ {
        fmt.Println(fi5, v2, v4)
        _ = fi5
    }
    v4 = v2
    for fi6 := 0; fi6 < 10; fi6++ {
        fmt.Println(fi6)
        v2 = v2
        switch 1 {
        case -2147483648:
            var v7 uint8 = 127
            _ = v7
        }
        _ = fi6
    }
    v2 = false
    for fi8 := 0; fi8 < 7; fi8++ {
        for fi9 := 0; fi9 < 9; fi9++ {
            switch "\n\t" {
            case "\n\t":
                switch "foo bar" {
                case "\n\t":
                    if ((23 <= 1) && (741 == 196)) {
                        var v10 uint64 = 91
                        _ = v10
                    }
                }
            case "foo bar":
                if len(v1) > 0 {
                    v1[0] = 719
                }
            case "":
                v11 := GenS1{F0: 255, F1: 128}
                _ = v11
            }
            switch fi9 {
            case 127:
                if ((2 >= 737) && false) {
                    fmt.Println(fi8)
                } else {
                    v4 = (v2 || (1e10 > 1e10))
                    fmt.Println(v4, v4)
                    if len(v1) > 0 {
                        v1[0] = 1
                    }
                    var v12 int8 = -78
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    v1 = append(v1, 0)
                    if (-2147483648 > 1) {
                        switch (fi9 % ((fi9) | 1)) {
                        case 2147483647:
                            var v13 float64 = 0.1
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            v12 += v12
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            if true {
                                v12 += v12
                                v13 += ((v13 + 0.1) + 1.5)
                                fmt.Println(v13)
                            }
                            _ = v13
                        case 1:
                            fmt.Println(v4, v4, v2)
                        default:
                            if len(v1) > 0 {
                                v1[0] = 42
                            }
                        }
                        switch (fi8 + (fi9 / (((-(fi9))) | 1))) {
                        case 32767:
                            if len(v1) > 0 {
                                v1[0] = 65535
                            }
                        case 498:
                            v12 -= (v12 >> ((255) & 63))
                        default:
                            switch "hello" {
                            case "a":
                                if true {
                                    v14 := GenS2{F0: 1}
                                    var v15 uint8 = 1
                                    _ = v14
                                    _ = v15
                                } else {
                                    fmt.Println(v2, v4, v3)
                                    if (("\n\t" + "") != "foo bar") {
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        v16 := GenS1{F0: 128, F1: 42}
                                        fmt.Println(v3, v12, v12)
                                        fmt.Println(v12, v3, fi8)
                                        _ = v16
                                    } else {
                                        if v4 {
                                            switch "a" {
                                            case "a":
                                                if false {
                                                    v2 = (v4 && true)
                                                    v3 -= v3
                                                } else {
                                                    var v17 uint = 268
                                                    if (127 < 2) {
                                                        var v18 int64 = 232
                                                        _ = v18
                                                    }
                                                    switch -990 {
                                                    case 0:
                                                        var v19 int = fi9
                                                        _ = v19
                                                    }
                                                    var v20 uint32 = 1
                                                    var v21 bool = (128 <= 127)
                                                    _ = v17
                                                    _ = v20
                                                    _ = v21
                                                }
                                            }
                                            v4 = (-9223372036854775808 == 1)
                                            var v22 float32 = -2.25
                                            _ = v22
                                        } else {
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                        }
                                        v2 = (68 <= 447)
                                        v1 = append(v1, 886)
                                    }
                                    switch "a" {
                                    case "hello":
                                        v3 += (v3 % ((v3) | 1))
                                    case "a":
                                        fmt.Println(v3, v12)
                                    default:
                                        fmt.Println(v2, v3)
                                    }
                                }
                            default:
                                v1 = append(v1, 42)
                            }
                        }
                        v2 = true
                    } else {
                        if v4 {
                            fmt.Println(v12, v12)
                            var v23 bool = (true || true)
                            _ = v23
                        } else {
                            var v24 uint16 = v3
                            _ = v24
                        }
                        v25 := GenS3{F0: 0}
                        _ = v25
                    }
                    if v2 {
                        v2 = (1 < 2147483647)
                    }
                    _ = v12
                }
            case -9223372036854775808:
                var v26 uint8 = 168
                _ = v26
            case 9223372036854775807:
                v27 := GenS2{F0: -815}
                _ = v27
            default:
                fmt.Println(fi9, fi8, fi9)
            }
            var v28 uint16 = v3
            _ = fi9
            _ = v28
        }
        v3 = 65535
        fmt.Println(v3, v4, v4)
        _ = fi8
    }
    if v4 {
        v3 = (v3 + (v3 << ((255) & 63)))
        v4 = v4
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    } else {
        fmt.Println(v2, v4)
        switch -978 {
        case 42:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        case 2147483647:
            v29 := GenS1{F0: 127, F1: 189}
            _ = v29
        }
    }
    if ((1e10 <= 0.1) || v2) {
        v4 = true
    }
    if (159 == 0) {
        var v30 uint8 = 42
        for fi31 := 0; fi31 < 18; fi31++ {
            v32 := GenS3{F0: 4294967295}
            _ = fi31
            _ = v32
        }
        var v33 uint = 939
        _ = v30
        _ = v33
    } else {
        var v34 uint = 1
        _ = v34
    }
    v35 := GenS4{F0: 76}
    fmt.Println(v2)
    if v4 {
        if len(v1) > 0 {
            v1[0] = 65535
        }
        v1 = append(v1, 993)
    }
    v36 := GenS4{F0: 962}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v4
}
