package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint32; F1 int64 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint8; F1 uint8; F2 uint }
type GenS6 struct { F0 uint16 }

func main() {
    var v1 uint = 42
    v2 := make(map[string]bool)
    v3 := make(map[uint32]int)
    var v4 uint64 = 2
    v5 := []int{0, -867, -284}
    var v6 int = 271
    v7 := GenS1{F0: (v6 & (v6 & v6))}
    v1 = (v1 + (v1 + 1))
    for fi8 := 0; fi8 < 7; fi8++ {
        v9 := GenS1{F0: v6}
        _ = fi8
        _ = v9
    }
    if true {
        var v10 uint64 = v4
        for fi11 := 0; fi11 < 16; fi11++ {
            v10 = 934
            v12 := GenS1{F0: (v6 - v6)}
            _ = fi11
            _ = v12
        }
        _ = v10
    } else {
        if (3.14159 > float32(-2.25)) {
            v13 := GenS2{F0: 757, F1: 42}
            for fi14 := 0; fi14 < 13; fi14++ {
                var v15 uint = v1
                if false {
                    v15 = v1
                } else {
                    var v16 int = fi14
                    _ = v16
                }
                for fi17 := 0; fi17 < 7; fi17++ {
                    switch (fi17 << (((v1 | v1)) & 63)) {
                    case 149:
                        v18 := GenS3{F0: 634}
                        _ = v18
                    }
                    if (-1 <= -112) {
                        if true {
                            var v19 string = "foo bar"
                            _ = v19
                        }
                    }
                    v4 = (v4 / (((v4 ^ (v4 % (((v4 << (((^(v1))) & 63))) | 1)))) | 1))
                    _ = fi17
                }
                _ = fi14
                _ = v15
            }
            _ = v13
        } else {
            v4 = v4
        }
        v20 := GenS4{F0: 42}
        switch "foo bar" {
        case "hello":
            if true {
                v1 = (v1 + 165)
                if ((0.1 * ((1.5 * (-2.25 - float64(float32((float64(1e10) / 1e10))))) * 0.0)) == (0.1 / (0.0 + 1e10))) {
                    var v21 float32 = 1e10
                    v1 += 584
                    var v22 float32 = v21
                    _ = v21
                    _ = v22
                }
                v23 := GenS2{F0: 65535, F1: 1}
                _ = v23
            }
        case "":
            fmt.Println(v4)
        case "a":
            for fi24 := 0; fi24 < 10; fi24++ {
                _ = len(v5)
                if false {
                    v25 := GenS5{F0: 42, F1: 214, F2: v1}
                    _ = v25
                } else {
                    v2["\n\t"] = (!((v1 + (v1 + v1)) > uint(v4)))
                    if (-719 <= 0) {
                        delete(v2, "foo bar")
                    } else {
                        switch (v6 + fi24) {
                        case -9223372036854775808:
                            switch v6 {
                            case 128:
                                switch "\n\t" {
                                case "foo bar":
                                    var v26 string = ("" + (("a" + "foo bar") + "\n\t"))
                                    _ = v26
                                case "a":
                                    v27 := GenS2{F0: 4294967295, F1: 941}
                                    _ = v27
                                case "":
                                    if false {
                                        fmt.Println(fi24, v4, v4)
                                        if true {
                                            v1 = v1
                                        }
                                    } else {
                                        for fi28 := 0; fi28 < 11; fi28++ {
                                            var v29 float32 = 0.1
                                            if (42 != -818) {
                                                var v30 int64 = 9223372036854775807
                                                fmt.Println(v6, v6, fi24)
                                                var v31 uint32 = 2
                                                _ = v30
                                                _ = v31
                                            } else {
                                                var v32 uint64 = uint64(v1)
                                                var v33 int = (^(fi28))
                                                v2["hello"] = (!true)
                                                _ = v32
                                                _ = v33
                                            }
                                            fmt.Println(fi28, v4, v1)
                                            v34 := GenS2{F0: 2147483647, F1: -826}
                                            _ = fi28
                                            _ = v29
                                            _ = v34
                                        }
                                        v1 = (v1 << (((v1 - 0)) & 63))
                                    }
                                }
                            default:
                                for fi35 := 0; fi35 < 15; fi35++ {
                                    v6 += 2147483647
                                    _ = fi35
                                }
                            }
                        case 42:
                            for fi36 := 0; fi36 < 2; fi36++ {
                                var v37 int16 = 1000
                                v38 := GenS3{F0: 591}
                                fmt.Println(v4, fi24, v6)
                                _ = fi36
                                _ = v37
                                _ = v38
                            }
                        default:
                            _ = len(v5)
                        }
                        switch "hello" {
                        case "hello":
                            _ = len(v5)
                        }
                        v6 = (v6 ^ v6)
                    }
                }
                v2["foo bar"] = false
                _ = fi24
            }
        default:
            fmt.Println(v6, v6)
        }
        _ = v20
    }
    if true {
        _ = len(v5)
        fmt.Println(v6, v4)
        fmt.Println(v1)
    }
    var v39 int8 = -128
    fmt.Println(v39, v6, v6)
    v40 := GenS6{F0: uint16(v1)}
    var v41 uint64 = v4
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
    _ = v7
}
