package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32 }
type GenS2 struct { F0 uint32; F1 int32 }
type GenS3 struct { F0 string; F1 uint32 }

func main() {
    var v1 uint16 = 255
    var v2 float32 = -2.25
    v3 := []int{-459, 2147483647, -348, -9223372036854775808}
    var v4 float64 = float64(float32((1e10 - -2.25)))
    v5 := GenS1{F0: (^(v1)), F1: 1}
    var v6 uint16 = 255
    v2 += v2
    fmt.Println(v4)
    v2 -= float32(((float64(v2) + (1.5 + v4)) * 3.14159))
    for fi7 := 0; fi7 < 11; fi7++ {
        if ((((759 <= 42) || ((fi7 << ((717) & 63)) == (^(fi7)))) || (830 < 65535)) && true) {
            var v8 uint32 = 2147483647
            v8 = 42
            for fi9 := 0; fi9 < 14; fi9++ {
                var v10 float64 = v4
                var v11 uint8 = 127
                v1 -= uint16(v11)
                v4 += 0.0
                if false {
                    if true {
                        var v12 int = (fi9 / ((int(v11)) | 1))
                        var v13 int64 = 1
                        fmt.Println(v13, v11, v4)
                        _ = v12
                        _ = v13
                    }
                } else {
                    v14 := GenS1{F0: v6, F1: -128}
                    _ = v14
                }
                _ = fi9
                _ = v10
                _ = v11
            }
            _ = v8
        }
        fmt.Println(fi7, v2, v2)
        _ = fi7
    }
    switch "foo bar" {
    case "hello":
        v15 := GenS2{F0: 65535, F1: 0}
        _ = v15
    default:
        switch -2147483648 {
        case 32767:
            for fi16 := 0; fi16 < 20; fi16++ {
                if true {
                    fmt.Println(v2, v1)
                    v1 = (v1 % (((v1 * (v6 | v6))) | 1))
                    var v17 int8 = 127
                    _ = v17
                } else {
                    var v18 uint64 = 42
                    _ = v18
                }
                var v19 float64 = (v4 - 1.5)
                v6 += (v1 >> ((65535) & 63))
                _ = fi16
                _ = v19
            }
        case 42:
            fmt.Println(v1, v2, v4)
        case -763:
            _ = len(v3)
        }
    }
    if true {
        for fi20 := 0; fi20 < 9; fi20++ {
            v21 := GenS3{F0: "hello", F1: 0}
            if len(v3) > 0 {
                v3[0] = 1
            }
            _ = fi20
            _ = v21
        }
        if false {
            _ = len(v3)
            if (18446744073709551615 < 1) {
                v3 = append(v3, 127)
                switch ("\n\t" + "a") {
                case "a":
                    fmt.Println(v1, v1, v1)
                case "foo bar":
                    v2 += (v2 + -2.25)
                case "hello":
                    switch "a" {
                    case "foo bar":
                        switch "" {
                        case "a":
                            fmt.Println(v6)
                        }
                    default:
                        switch "foo bar" {
                        case "\n\t":
                            v6 -= (v1 * (v6 / (((v6 | v1)) | 1)))
                        }
                    }
                }
                switch -133 {
                case 42:
                    var v22 int32 = -2
                    _ = v22
                }
            } else {
                var v23 uint16 = (v1 - (^(v1)))
                switch 4294967295 {
                case 4294967295:
                    var v24 int16 = 42
                    _ = v24
                default:
                    var v25 float64 = 0.1
                    _ = v25
                }
                v6 = v23
                _ = v23
            }
        } else {
            fmt.Println(v6, v2)
            v2 -= v2
            v2 -= float32(1.5)
        }
    } else {
        v4 = ((1.5 + (v4 - 1.5)) / (0.1 - v4))
        var v26 bool = (0 > 269)
        _ = v26
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
