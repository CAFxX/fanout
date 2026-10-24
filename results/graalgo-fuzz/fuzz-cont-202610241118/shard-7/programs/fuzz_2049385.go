package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16 }
type GenS2 struct { F0 int16; F1 int8 }
type GenS3 struct { F0 int }
func genFn0(p0 bool, p1 int, p2 int32) uint16 {
    fmt.Println(p0, p1, p2)
    for fi4 := 0; fi4 < 8; fi4++ {
        p0 = (p0 || p0)
        v5 := GenS1{F0: 6, F1: 2}
        switch "hello" {
        case "\n\t":
            v6 := GenS2{F0: 1000, F1: -55}
            _ = v6
        case "foo bar":
            switch -2147483648 {
            case -9223372036854775808:
                var v7 uint8 = 79
                _ = v7
            case -1:
                p2 -= p2
            default:
                fmt.Println(p0)
            }
        case "hello":
            p1 = 128
        }
        _ = fi4
        _ = v5
    }
    switch p1 {
    case 230:
        v8 := []int{(p1 >> ((255) & 63)), p1, p1}
        _ = v8
    case 4294967295:
        p0 = (127 >= 154)
    }
    _ = p0
    _ = p1
    _ = p2
    return 42
}

func main() {
    var v1 string = "\n\t"
    var v2 uint16 = 65535
    var v3 int8 = 42
    for fi9 := 0; fi9 < 13; fi9++ {
        v10 := GenS2{F0: -559, F1: v3}
        v11 := []int{fi9, (fi9 - -99), (-(fi9))}
        _ = fi9
        _ = v10
        _ = v11
    }
    var v12 uint64 = 4294967295
    v12 = v12
    v13 := GenS2{F0: 1000, F1: (v3 * 16)}
    v14 := []int{-102, 32768}
    for fi15 := 0; fi15 < 3; fi15++ {
        for fi16 := 0; fi16 < 8; fi16++ {
            if len(v14) > 0 {
                v14[0] = fi15
            }
            var v17 int = (fi15 * fi16)
            _ = fi16
            _ = v17
        }
        v2 -= (v2 | 880)
        if len(v14) >= 2 {
            v14 = v14[1:len(v14)]
        }
        _ = fi15
    }
    v18 := GenS2{F0: -898, F1: (v3 + -21)}
    var v19 float64 = (0.1 - (1.5 - 1.5))
    if (3.14159 <= ((float32((1.5 * v19)) / -2.25) / 0.1)) {
        if (0 != 314) {
            switch v1 {
            case "a":
                var v20 uint8 = 1
                _ = v20
            case "":
                v12 = uint64(v2)
            case "hello":
                switch v1 {
                case "":
                    v1 = "a"
                case "hello":
                    v3 -= v3
                default:
                    switch v1 {
                    case "foo bar":
                        v21 := GenS2{F0: 193, F1: 100}
                        _ = v21
                    case "a":
                        v22 := GenS1{F0: (v12 + 18446744073709551615), F1: (v2 / (((v2 / ((v2) | 1))) | 1))}
                        _ = v22
                    }
                }
            default:
                v3 += 42
            }
            v23 := GenS1{F0: 0, F1: v2}
            if true {
                for fi24 := 0; fi24 < 4; fi24++ {
                    v25 := GenS1{F0: 1, F1: genFn0(true, fi24, 0)}
                    v3 = v3
                    _ = fi24
                    _ = v25
                }
                var v26 float64 = v19
                v27 := GenS2{F0: -829, F1: v3}
                _ = v26
                _ = v27
            } else {
                var v28 int = 4294967295
                _ = v28
            }
            _ = v23
        } else {
            v14 = append(v14, 32767)
        }
        v29 := GenS3{F0: 2147483647}
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v18
    _ = v19
    _ = v2
    _ = v3
}
