package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16 }
type GenS2 struct { F0 uint32; F1 int32 }
type GenS3 struct { F0 int64; F1 bool }
type GenS4 struct { F0 int8 }

func main() {
    v1 := []int{32768, -9223372036854775808, 32768}
    var v2 uint = 255
    var v3 uint64 = 42
    v4 := make(map[string]string)
    v5 := make(map[int]uint32)
    var v6 int = 0
    if (("" > "hello") || true) {
        v1 = append(v1, 789)
    } else {
        for fi7 := 0; fi7 < 14; fi7++ {
            var v8 uint32 = uint32(v2)
            switch (fi7 | -749) {
            case 42:
                var v9 uint16 = uint16(v8)
                _ = v9
            }
            if (false || (false || (!true))) {
                v10 := GenS1{F0: 42, F1: -389}
                for fi11 := 0; fi11 < 14; fi11++ {
                    v3 = v3
                    _ = fi11
                }
                _ = v10
            } else {
                v2 = (v2 << ((uint(v8)) & 63))
                v8 = (v8 << (((v2 + (v2 % ((1) | 1)))) & 63))
            }
            _ = fi7
            _ = v8
        }
        v12 := v4["a"]
        _ = v12
    }
    switch (v6 + -9223372036854775808) {
    case 127:
        switch v6 {
        case -663:
            v3 = (v3 << ((v2) & 63))
        case 219:
            fmt.Println(v6)
        case -508:
            if true {
                if (true && false) {
                    fmt.Println(v2, v3, v2)
                } else {
                    for fi13 := 0; fi13 < 5; fi13++ {
                        switch "hello" {
                        case "\n\t":
                            v2 = v2
                        }
                        _ = fi13
                    }
                    _ = len(v1)
                }
                fmt.Println(v6, v2, v3)
            } else {
                _ = len(v1)
                v6 = 2147483647
            }
        }
    case 42:
        if true {
            for fi14 := 0; fi14 < 14; fi14++ {
                var v15 string = (("\n\t" + (("foo bar" + ("" + "a")) + "\n\t")) + "hello")
                v16 := GenS1{F0: 0, F1: int16(v2)}
                _ = fi14
                _ = v15
                _ = v16
            }
            switch "hello" {
            case "foo bar":
                switch v6 {
                case -2147483648:
                    v3 += (^(v3))
                case 837:
                    var v17 int16 = 1
                    _ = v17
                case 32768:
                    for fi18 := 0; fi18 < 1; fi18++ {
                        for fi19 := 0; fi19 < 13; fi19++ {
                            fmt.Println(fi19, fi19, v2)
                            _ = fi19
                        }
                        v1 = append(v1, fi18)
                        _ = fi18
                    }
                default:
                    for fi20 := 0; fi20 < 15; fi20++ {
                        if (!false) {
                            for fi21 := 0; fi21 < 12; fi21++ {
                                if false {
                                    switch "\n\t" {
                                    case "":
                                        if false {
                                            v22 := GenS2{F0: 2147483647, F1: 2147483647}
                                            _ = v22
                                        }
                                    case "foo bar":
                                        v23 := GenS2{F0: 4294967295, F1: 42}
                                        _ = v23
                                    default:
                                        v5[322] = 65535
                                    }
                                    v5[(v6 % (((fi21 * (v6 & fi21))) | 1))] = 731
                                } else {
                                    var v24 uint = v2
                                    var v25 uint8 = 2
                                    if ((v3 ^ v3) > 2) {
                                        switch v6 {
                                        case -9223372036854775808:
                                            var v26 int64 = -9223372036854775808
                                            v26 = (v26 ^ int64(v3))
                                            _ = v26
                                        default:
                                            v25 = v25
                                        }
                                    } else {
                                        fmt.Println(v24, v6)
                                        var v27 string = "hello"
                                        _ = v27
                                    }
                                    _ = v24
                                    _ = v25
                                }
                                fmt.Println(v6, v6)
                                _ = fi21
                            }
                            fmt.Println(v6)
                        } else {
                            var v28 uint16 = 489
                            _ = v28
                        }
                        _ = fi20
                    }
                }
            }
        }
    default:
        var v29 uint64 = (^(v3))
        _ = v29
    }
    v30 := GenS3{F0: -383, F1: false}
    fmt.Println(v6)
    fmt.Println(v6, v2, v3)
    v3 = v3
    v31 := GenS4{F0: -128}
    v32 := GenS3{F0: 0, F1: (-2.25 <= 3.14159)}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
