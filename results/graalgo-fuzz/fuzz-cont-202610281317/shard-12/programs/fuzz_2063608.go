package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64; F2 uint }
type GenS2 struct { F0 int16; F1 uint16 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint8 = 128
    v3 := []string{"a", "\n\t", ""}
    switch -9223372036854775808 {
    case 2147483647:
        switch "a" {
        case "foo bar":
            if true {
                var v4 float64 = ((((1.5 - (0.0 / 3.14159)) - 1e10) - -2.25) / float64(((1e10 * (3.14159 + 0.0)) / ((((1.5 / 1.5) - 0.1) / (0.0 - float32(1e10))) / 1.0))))
                _ = v4
            }
        case "a":
            var v5 int8 = -1
            _ = v5
        case "\n\t":
            var v6 int8 = int8(v2)
            _ = v6
        }
    }
    var v7 int16 = 1000
    if true {
        for fi8 := 0; fi8 < 1; fi8++ {
            var v9 int64 = 0
            var v10 int8 = 42
            _ = fi8
            _ = v10
            _ = v9
        }
        if false {
            fmt.Println(v2, v2)
            for fi11 := 0; fi11 < 8; fi11++ {
                fmt.Println(fi11, v2, v2)
                var v12 int8 = -1
                if (4294967295 == 402) {
                    for fi13 := 0; fi13 < 9; fi13++ {
                        v14 := GenS1{F0: 1, F1: 42, F2: 777}
                        var v15 bool = false
                        var v16 uint = 42
                        if ((v12 | (v12 << ((2) & 63))) <= (v12 >> (((^(v16))) & 63))) {
                            var v17 float64 = (float64(0.1) + 0.0)
                            var v18 int32 = 2147483647
                            switch "a" {
                            case "a":
                                v12 -= 127
                            case "\n\t":
                                v7 = (v7 - (v7 & (v7 ^ -794)))
                                var v19 string = ""
                                _ = v19
                            case "":
                                v1[65535] = ((v17 / (float64(((-2.25 + -2.25) - 1.5)) / float64(3.14159))) == ((v17 + ((v17 * v17) * v17)) / (v17 / ((v17 * (v17 * (1e10 - float64((1e10 + float32(0.0)))))) * 1e10))))
                            }
                            v20 := GenS1{F0: 4294967295, F1: 42, F2: 2}
                            _ = v17
                            _ = v18
                            _ = v20
                        }
                        var v21 bool = false
                        _ = fi13
                        _ = v14
                        _ = v15
                        _ = v16
                        _ = v21
                    }
                    fmt.Println(fi11)
                }
                _ = fi11
                _ = v12
            }
        }
        v7 += v7
    } else {
        if (!((^(v2)) >= 255)) {
            switch ("" + ("a" + ("foo bar" + "\n\t"))) {
            case "\n\t":
                v22 := GenS1{F0: 229, F1: 2, F2: 343}
                _ = v22
            }
            switch 587 {
            case 32768:
                for fi23 := 0; fi23 < 2; fi23++ {
                    v1[2147483648] = false
                    fmt.Println(fi23, fi23, v2)
                    v2 -= (^(v2))
                    _ = fi23
                }
            case -427:
                v7 -= (v7 & (-(v7)))
            case 127:
                v3 = append(v3, ("foo bar" + ("a" + "foo bar")))
            default:
                if (720 != 255) {
                    v2 = (v2 ^ 127)
                    var v24 uint = 255
                    _ = v24
                }
            }
            v25 := GenS2{F0: (v7 - v7), F1: 2}
            _ = v25
        }
        var v26 int8 = 0
        _ = v26
    }
    v27 := GenS1{F0: uint64(v2), F1: 1, F2: 0}
    delete(v1, 330)
    v28 := GenS1{F0: 2, F1: 2, F2: 255}
    switch "" {
    case "\n\t":
        switch ("foo bar" + "foo bar") {
        case "a":
            fmt.Println(v7, v2)
        }
    default:
        if (100 <= -49) {
            v3 = append(v3, "foo bar")
        } else {
            v7 -= (v7 | 1)
            v29 := GenS2{F0: v7, F1: 2}
            _ = v29
        }
    }
    if true {
        for fi30 := 0; fi30 < 10; fi30++ {
            var v31 int16 = 1000
            var v32 float32 = 1e10
            v32 = v32
            _ = fi30
            _ = v31
            _ = v32
        }
    }
    for fi33 := 0; fi33 < 12; fi33++ {
        _ = len(v1)
        fmt.Println(v7)
        if true {
            fmt.Println(fi33, fi33, v7)
        } else {
            v7 += 1000
            v7 += v7
            v34 := GenS2{F0: v7, F1: 65535}
            _ = v34
        }
        _ = fi33
    }
    v7 -= (-(v7))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v7
}
