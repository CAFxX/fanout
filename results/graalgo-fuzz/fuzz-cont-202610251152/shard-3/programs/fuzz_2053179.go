package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8 }
type GenS2 struct { F0 bool }

func main() {
    var v1 int64 = -1
    v2 := make(map[int]bool)
    var v3 uint32 = 4294967295
    v4 := make(map[string]uint32)
    var v5 uint64 = 0
    v6 := []string{(("\n\t" + (("foo bar" + ("foo bar" + "a")) + "foo bar")) + "hello")}
    switch "\n\t" {
    case "\n\t":
        var v7 int32 = 0
        _ = v7
    case "foo bar":
        _ = len(v6)
    case "a":
        switch 431 {
        case 9223372036854775807:
            for fi8 := 0; fi8 < 5; fi8++ {
                if true {
                    var v9 float64 = 1.5
                    if true {
                        var v10 float64 = ((3.14159 / 1.5) + (float64(-2.25) / v9))
                        var v11 int = 0
                        _ = v10
                        _ = v11
                    }
                    _ = v9
                } else {
                    v12 := GenS1{F0: 550, F1: 127}
                    _ = v12
                }
                _ = fi8
            }
        default:
            v1 = v1
        }
    default:
        v13 := GenS1{F0: 0, F1: 11}
        _ = v13
    }
    v1 += (v1 % (((^(v1))) | 1))
    switch "" {
    case "":
        switch -779 {
        case 127:
            var v14 uint = 65535
            _ = v14
        case 32768:
            v5 = v5
        }
    case "hello":
        fmt.Println(v5, v1, v3)
    case "\n\t":
        fmt.Println(v1, v1, v5)
    }
    v5 = (v5 << ((42) & 63))
    for fi15 := 0; fi15 < 6; fi15++ {
        switch (("a" + "\n\t") + "foo bar") {
        case "a":
            var v16 float32 = (1.5 * ((float32((((0.1 / 3.14159) * -2.25) / 1.5)) - float32(3.14159)) + (0.1 + 0.1)))
            _ = v16
        case "hello":
            if (true && (-259 < 0)) {
                fmt.Println(v5)
                var v17 float64 = ((float64(1.5) * 0.0) - (0.0 * (0.1 / 0.1)))
                switch (fi15 ^ fi15) {
                case -1:
                    for fi18 := 0; fi18 < 2; fi18++ {
                        var v19 int8 = 0
                        v17 += (float64(0.1) - v17)
                        _ = fi18
                        _ = v19
                    }
                case 42:
                    if (3.14159 != 0.0) {
                        v20 := GenS1{F0: 2147483647, F1: 127}
                        for fi21 := 0; fi21 < 3; fi21++ {
                            v5 += v5
                            v5 -= (v5 | v5)
                            _ = fi21
                        }
                        _ = v20
                    } else {
                        fmt.Println(v17, v17)
                        if false {
                            v6 = append(v6, ("\n\t" + "a"))
                            v5 += (^(v5))
                        } else {
                            var v22 float32 = 0.0
                            v17 = 0.1
                            _ = v22
                        }
                        if true {
                            var v23 uint32 = (^(v3))
                            _ = v23
                        } else {
                            for fi24 := 0; fi24 < 1; fi24++ {
                                v3 = (v3 | v3)
                                _ = fi24
                            }
                        }
                    }
                default:
                    var v25 string = "\n\t"
                    _ = v25
                }
                _ = v17
            }
        case "":
            v5 = (^(v5))
        }
        var v26 uint = 660
        v5 = uint64(v26)
        _ = fi15
        _ = v26
    }
    var v27 string = "\n\t"
    v28 := GenS2{F0: ((1.5 / -2.25) != 3.14159)}
    if (false || (!true)) {
        v3 += 0
    } else {
        var v29 int64 = (^(v1))
        _ = v29
    }
    v6 = append(v6, v27)
    fmt.Println(v1, v1)
    fmt.Println(v5, v1, v3)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
