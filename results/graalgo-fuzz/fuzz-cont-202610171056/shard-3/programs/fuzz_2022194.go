package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32; F2 uint }
type GenS2 struct { F0 int8; F1 int8; F2 uint64 }
type GenS3 struct { F0 uint }

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[uint32]bool)
    v3 := []int{127, -9223372036854775808, 4294967295, 128}
    var v4 int = 4294967295
    var v5 int8 = 127
    var v6 int16 = 669
    v5 = v5
    switch ("a" + (("hello" + "a") + ("a" + "foo bar"))) {
    case "hello":
        if (true && (int16(v4) < int16(v4))) {
            fmt.Println(v6)
        } else {
            v1[0] = 0
            if true {
                v5 = v5
                v5 = -28
            } else {
                fmt.Println(v5)
                if len(v3) > 0 {
                    v3[0] = (v4 >> ((65535) & 63))
                }
            }
        }
    case "foo bar":
        v7 := GenS1{F0: v6, F1: 130, F2: 1}
        _ = v7
    }
    delete(v2, 2147483647)
    _ = len(v3)
    if true {
        for fi8 := 0; fi8 < 14; fi8++ {
            v4 = (fi8 << ((858) & 63))
            var v9 uint8 = 127
            switch (fi8 * fi8) {
            case 9223372036854775807:
                var v10 bool = true
                _ = v10
            case 32768:
                v11 := GenS1{F0: -32768, F1: 1, F2: 895}
                _ = v11
            }
            _ = fi8
            _ = v9
        }
    } else {
        var v12 int = v4
        _ = v12
    }
    if true {
        var v13 string = "\n\t"
        if true {
            v5 = (v5 | (-(v5)))
            v14 := GenS1{F0: 148, F1: -853, F2: 65535}
            v6 += (v6 ^ 42)
            _ = v14
        }
        _ = v13
    } else {
        var v15 uint16 = 920
        fmt.Println(v5)
        _ = v15
    }
    v5 -= v5
    var v16 uint = 42
    v17 := GenS1{F0: v6, F1: -2147483648, F2: 340}
    if len(v3) > 0 {
        v3[0] = 4294967295
    }
    if (192 <= 18446744073709551615) {
        if (!false) {
            switch "hello" {
            case "\n\t":
                v18 := v2[4294967295]
                _ = v18
            }
            var v19 uint32 = 289
            var v20 uint16 = 105
            _ = v19
            _ = v20
        }
        for fi21 := 0; fi21 < 15; fi21++ {
            for fi22 := 0; fi22 < 12; fi22++ {
                var v23 uint32 = 2147483647
                v5 = (v5 + (v5 + 88))
                v24 := v2[v23]
                _ = fi22
                _ = v23
                _ = v24
            }
            for fi25 := 0; fi25 < 18; fi25++ {
                v16 -= v16
                if true {
                    if (2 != 2147483648) {
                        var v26 int16 = v6
                        var v27 float64 = float64((0.0 / 1.0))
                        fmt.Println(v5, v5)
                        _ = v26
                        _ = v27
                    }
                    v28 := GenS1{F0: v6, F1: 999, F2: (v16 >> ((v16) & 63))}
                    delete(v1, (v4 + (fi25 << ((v16) & 63))))
                    v29 := GenS1{F0: -1, F1: 317, F2: 2}
                    v30 := GenS2{F0: (v5 / ((127) | 1)), F1: (v5 * v5), F2: 42}
                    _ = v28
                    _ = v29
                    _ = v30
                } else {
                    if ((v5 / ((100) | 1)) <= v5) {
                        v4 = (^(fi21))
                        v31 := GenS1{F0: (v6 << (((v16 ^ 363)) & 63)), F1: 0, F2: v16}
                        _ = v31
                    }
                    switch "\n\t" {
                    case "hello":
                        var v32 float32 = (1e10 + 3.14159)
                        _ = v32
                    case "\n\t":
                        fmt.Println(v5, v6, v5)
                    default:
                        delete(v1, (^(v4)))
                    }
                    v33 := v2[924]
                    _ = v33
                }
                var v34 int8 = (-(v5))
                switch fi21 {
                case 762:
                    switch ("\n\t" + ("hello" + "foo bar")) {
                    case "hello":
                        var v35 int16 = -138
                        _ = v35
                    case "":
                        switch fi21 {
                        case 127:
                            v3 = append(v3, (fi25 & (fi25 % ((1) | 1))))
                        case -559:
                            var v36 bool = false
                            if (65535 == uint32(v16)) {
                                switch "\n\t" {
                                case "foo bar":
                                    v37 := GenS3{F0: 1}
                                    _ = v37
                                case "":
                                    var v38 int32 = -164
                                    _ = v38
                                default:
                                    var v39 uint16 = 1
                                    _ = v39
                                }
                                v40 := GenS2{F0: (v34 / ((77) | 1)), F1: (^(v34)), F2: 290}
                                var v41 int8 = v5
                                _ = v40
                                _ = v41
                            }
                            v16 = (v16 | 62)
                            _ = v36
                        }
                    }
                }
                _ = fi25
                _ = v34
            }
            fmt.Println(fi21)
            _ = fi21
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
