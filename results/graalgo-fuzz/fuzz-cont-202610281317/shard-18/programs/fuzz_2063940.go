package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int; F1 uint; F2 int }
type GenS3 struct { F0 uint8; F1 int32; F2 uint64 }
type GenS4 struct { F0 string; F1 bool }
type GenS5 struct { F0 int32; F1 uint8; F2 string }
type GenS6 struct { F0 int64 }
func genFn0(p0 uint64) uint8 {
    v7 := []int{-9223372036854775808, 987, 0}
    if (42 < 2) {
        p0 -= 1
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
        _ = len(v7)
    } else {
        p0 = (^(p0))
    }
    var v8 uint64 = p0
    _ = p0
    _ = v7
    _ = v8
    return 221
}

func main() {
    v1 := []uint32{2147483647, 2, 42}
    var v2 int8 = 69
    var v3 bool = (42 < 42)
    var v4 int8 = -1
    var v5 int8 = (v2 + (^(v2)))
    var v6 int32 = -317
    v9 := GenS1{F0: 127}
    if v3 {
        var v10 int = 0
        v11 := GenS2{F0: v10, F1: 783, F2: v10}
        v4 += (v4 & v2)
        _ = v10
        _ = v11
    }
    for fi12 := 0; fi12 < 4; fi12++ {
        fmt.Println(v2, v4)
        fmt.Println(v5, v6, v3)
        switch 451 {
        case -9223372036854775808:
            var v13 int64 = -9223372036854775808
            _ = v13
        case 1:
            if false {
                v1 = append(v1, 1)
                var v14 float32 = float32(0.1)
                _ = v14
            } else {
                fmt.Println(v6)
                for fi15 := 0; fi15 < 14; fi15++ {
                    v5 = (v2 % ((127) | 1))
                    v3 = ((v2 * v4) == (-(v4)))
                    switch (fi12 / (((fi12 * -1)) | 1)) {
                    case 504:
                        fmt.Println(v6)
                    case 128:
                        var v16 int = fi12
                        _ = v16
                    default:
                        _ = len(v1)
                    }
                    _ = fi15
                }
            }
        case 32767:
            if (65535 != 0) {
                v5 -= (v4 << ((4294967295) & 63))
            } else {
                var v17 uint = 4294967295
                v4 = v5
                _ = v17
            }
        }
        _ = fi12
    }
    if (2147483647 == 663) {
        if len(v1) > 0 {
            v1[0] = 0
        }
    }
    v4 += (v4 | v5)
    switch (("hello" + ("" + "")) + (("hello" + ("" + "a")) + "a")) {
    case "\n\t":
        switch "foo bar" {
        case "\n\t":
            var v18 int8 = v2
            _ = v18
        case "":
            if (true && v3) {
                switch -9223372036854775808 {
                case -2147483648:
                    var v19 int = 4294967295
                    _ = v19
                case -455:
                    fmt.Println(v4, v2)
                default:
                    if (!v3) {
                        if (65535 == 290) {
                            switch -213 {
                            case 2147483647:
                                fmt.Println(v3)
                            case -9223372036854775808:
                                v5 = (v5 + 127)
                            default:
                                var v20 string = "a"
                                _ = v20
                            }
                            var v21 uint64 = 323
                            _ = v21
                        } else {
                            for fi22 := 0; fi22 < 1; fi22++ {
                                v23 := GenS3{F0: 0, F1: v6, F2: 1}
                                switch "" {
                                case "\n\t":
                                    v4 += (^(v5))
                                default:
                                    v4 = v4
                                }
                                _ = fi22
                                _ = v23
                            }
                        }
                    }
                }
                v1 = append(v1, 0)
            } else {
                fmt.Println(v2)
            }
        case "foo bar":
            for fi24 := 0; fi24 < 13; fi24++ {
                v6 += v6
                v25 := GenS4{F0: "a", F1: (0.0 <= 0.0)}
                if false {
                    _ = len(v1)
                }
                _ = fi24
                _ = v25
            }
        }
    case "foo bar":
        var v26 float32 = float32(0.0)
        _ = v26
    }
    var v27 float32 = 0.0
    if true {
        var v28 uint64 = 4294967295
        _ = v28
    } else {
        for fi29 := 0; fi29 < 1; fi29++ {
            v1 = append(v1, 914)
            if (1.5 > 0.0) {
                var v30 uint8 = 127
                var v31 float32 = ((((0.0 + v27) + (1.5 + float32((1e10 * float64(v27))))) + float32(1e10)) / ((v27 * (v27 / v27)) * ((1.5 / v27) / v27)))
                _ = v30
                _ = v31
            }
            _ = fi29
        }
        v32 := GenS5{F0: v6, F1: 2, F2: ("foo bar" + (("hello" + ("" + "\n\t")) + ""))}
        v33 := GenS6{F0: 42}
        _ = v32
        _ = v33
    }
    var v34 uint16 = 32767
    fmt.Println(v5, v34)
    var v35 int32 = v6
    for fi36 := 0; fi36 < 6; fi36++ {
        var v37 uint8 = genFn0(1)
        var v38 string = "foo bar"
        v2 -= (v2 * (-(v2)))
        _ = fi36
        _ = v37
        _ = v38
    }
    var v39 uint16 = 2
    v1 = append(v1, 227)
    var v40 int = 32768
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v34
    _ = v35
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
    _ = v9
}
