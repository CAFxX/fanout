package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int64; F1 int16 }
func genFn0(p0 string, p1 int64) int32 {
    p1 -= (p1 | -450)
    if true {
        if true {
            fmt.Println(p1, p0)
            switch 678 {
            case 32768:
                v6 := GenS1{F0: 255, F1: 1000}
                _ = v6
            case 978:
                p0 = p0
            case 9223372036854775807:
                var v7 float64 = 0.1
                _ = v7
            }
        }
    } else {
        p1 -= 9223372036854775807
    }
    _ = p0
    _ = p1
    return 42
}

func main() {
    var v1 int16 = 42
    var v2 float64 = 0.0
    v3 := make(map[string]uint32)
    var v4 uint = 921
    var v5 string = ("foo bar" + ("a" + ("foo bar" + "a")))
    v8 := GenS1{F0: 65535, F1: (-(v1))}
    v9 := GenS1{F0: 160, F1: (v1 % (((v1 >> (((v4 / (((^(v4))) | 1))) & 63))) | 1))}
    var v10 int64 = -604
    v11 := GenS1{F0: 2, F1: v1}
    v2 += v2
    if false {
        v10 -= v10
        switch "foo bar" {
        case "foo bar":
            v4 = (v4 / ((v4) | 1))
        case "hello":
            v1 -= (^(v1))
        case "\n\t":
            if (!true) {
                v1 += v1
                switch v5 {
                case "a":
                    v4 = v4
                }
                v12 := []int{493, 32768}
                _ = v12
            } else {
                if false {
                    fmt.Println(v2, v1)
                } else {
                    fmt.Println(v1, v10, v1)
                }
                v13 := []int{}
                for fi14 := 0; fi14 < 16; fi14++ {
                    if ((3.14159 * ((1e10 + (1.5 / 1.0)) + 1e10)) > 0.1) {
                        switch ("" + v5) {
                        case "foo bar":
                            v15 := GenS1{F0: 390, F1: v1}
                            _ = v15
                        case "":
                            fmt.Println(v1)
                        case "\n\t":
                            for fi16 := 0; fi16 < 11; fi16++ {
                                var v17 float64 = v2
                                v18 := GenS2{F0: int8(v10)}
                                fmt.Println(v10)
                                _ = fi16
                                _ = v17
                                _ = v18
                            }
                        default:
                            _ = len(v13)
                        }
                    } else {
                        fmt.Println(v1, v10, v4)
                    }
                    switch (fi14 & fi14) {
                    case 42:
                        for fi19 := 0; fi19 < 16; fi19++ {
                            var v20 float32 = -2.25
                            _ = fi19
                            _ = v20
                        }
                    case 891:
                        switch (fi14 * (-(fi14))) {
                        case -9223372036854775808:
                            _ = len(v13)
                        }
                    }
                    if ((v1 >> (((^(v4))) & 63)) <= int16(fi14)) {
                        var v21 int8 = 42
                        v5 = (v5 + v5)
                        _ = v21
                    } else {
                        var v22 uint8 = 128
                        v22 -= 1
                        for fi23 := 0; fi23 < 13; fi23++ {
                            v13 = append(v13, 32767)
                            v1 -= v1
                            _ = fi23
                        }
                        _ = v22
                    }
                    _ = fi14
                }
                _ = v13
            }
        default:
            var v24 int = 2147483647
            _ = v24
        }
    } else {
        if (!(553 < 32767)) {
            fmt.Println(v2, v1)
            for fi25 := 0; fi25 < 19; fi25++ {
                v26 := GenS3{F0: -1, F1: 1000}
                v5 = v5
                _ = fi25
                _ = v26
            }
        }
    }
    v27 := GenS3{F0: 386, F1: 42}
    if false {
        var v28 int32 = 0
        var v29 int = int(v1)
        v28 = (v28 - genFn0("", (v10 ^ 9223372036854775807)))
        _ = v28
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
