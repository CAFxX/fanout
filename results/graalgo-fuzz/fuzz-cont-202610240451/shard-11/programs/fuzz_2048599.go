package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int }
type GenS4 struct { F0 string; F1 uint; F2 int32 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 uint; F1 bool }
type GenS7 struct { F0 uint8; F1 int16 }

func main() {
    var v1 uint16 = 2
    var v2 uint32 = 2147483648
    var v3 uint64 = 4294967295
    var v4 int = 904
    v4 -= 948
    fmt.Println(v4, v4, v3)
    fmt.Println(v3)
    var v5 uint64 = v3
    v6 := []int{32768, (v4 & 1), v4}
    var v7 uint = 1
    if true {
        v4 = int(v1)
    } else {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
    }
    for fi8 := 0; fi8 < 12; fi8++ {
        if false {
            v7 = 65535
        } else {
            v1 = v1
            v1 += 65535
            v9 := GenS1{F0: ("hello" + ("\n\t" + ("foo bar" + "hello")))}
            _ = v9
        }
        _ = fi8
    }
    v10 := GenS2{F0: "foo bar"}
    v4 += int(v2)
    v11 := GenS3{F0: -2147483648}
    v12 := GenS4{F0: "hello", F1: v7, F2: -2147483648}
    var v13 int32 = -2147483648
    if len(v6) > 0 {
        v6[0] = int(v13)
    }
    v14 := GenS5{F0: (v13 + -1)}
    for fi15 := 0; fi15 < 4; fi15++ {
        for fi16 := 0; fi16 < 18; fi16++ {
            if true {
                v17 := GenS5{F0: v13}
                var v18 int32 = int32(v1)
                v7 += v7
                _ = v17
                _ = v18
            } else {
                v5 = uint64(v1)
                switch -91 {
                case -2147483648:
                    fmt.Println(fi15)
                    switch (("\n\t" + "hello") + ("" + "a")) {
                    case "a":
                        switch "hello" {
                        case "":
                            switch (("a" + "") + "foo bar") {
                            case "":
                                if (-39 < -128) {
                                    if len(v6) >= 2 {
                                        v6 = v6[1:len(v6)]
                                    }
                                    v2 = v2
                                    v4 -= v4
                                } else {
                                    v19 := GenS6{F0: (^(v7)), F1: (false && true)}
                                    switch "a" {
                                    case "\n\t":
                                        var v20 bool = (1 >= v1)
                                        _ = v20
                                    default:
                                        switch (fi16 % ((0) | 1)) {
                                        case 128:
                                            var v21 uint64 = v5
                                            _ = v21
                                        }
                                    }
                                    _ = v19
                                }
                            case "hello":
                                _ = len(v6)
                            default:
                                v22 := GenS7{F0: 70, F1: 0}
                                _ = v22
                            }
                        case "a":
                            var v23 int16 = -32768
                            _ = v23
                        case "foo bar":
                            fmt.Println(v2, v5, v4)
                        }
                    case "foo bar":
                        v24 := GenS5{F0: v13}
                        _ = v24
                    default:
                        v1 += (v1 / ((1) | 1))
                    }
                    fmt.Println(v13)
                case 32768:
                    if (3.14159 <= 1e10) {
                        var v25 string = "\n\t"
                        _ = v25
                    }
                default:
                    switch (fi15 / (((fi15 << ((v7) & 63))) | 1)) {
                    case -9223372036854775808:
                        if (true || (0 >= 127)) {
                            var v26 int = int(v5)
                            _ = len(v6)
                            _ = v26
                        }
                    case 1:
                        fmt.Println(v1, v2, v13)
                        v3 = (v3 | uint64(v1))
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                        if (false && (uint16(v5) < uint16(v2))) {
                            v4 = (fi16 | 751)
                        } else {
                            v6 = append(v6, v4)
                        }
                        v27 := GenS2{F0: "\n\t"}
                        var v28 int = 0
                        _ = v27
                        _ = v28
                    }
                }
            }
            v7 += 4294967295
            if ((1.5 / ((0.0 + -2.25) - float32(float64(float32(0.1))))) < (3.14159 / 1.0)) {
                v3 -= v5
                var v29 float32 = (float32(float64(1e10)) * (((1e10 - 3.14159) * (1.5 + (0.0 + (1e10 + (((1e10 / (0.1 + (1e10 - -2.25))) / float32(-2.25)) - (0.1 + (float32(float64((0.0 - ((0.0 - (1e10 - (0.0 + float32(0.1)))) / float32((0.0 - 0.1)))))) / -2.25))))))) + (3.14159 - 1e10)))
                _ = v29
            }
            _ = fi16
        }
        _ = fi15
    }
    switch v4 {
    case 720:
        var v30 int64 = -992
        _ = v30
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
