package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool; F2 uint64 }
type GenS2 struct { F0 int64 }

func main() {
    var v1 uint = 65535
    var v2 bool = false
    var v3 uint16 = 1
    v4 := make(map[int]bool)
    var v5 uint16 = (v3 + (v3 - (v3 >> ((2) & 63))))
    v6 := GenS1{F0: 397, F1: (v2 || (true && (!true))), F2: 0}
    fmt.Println(v3, v1)
    v7 := []int{2147483647, 42, 127, 0}
    var v8 float64 = 1e10
    fmt.Println(v3, v8, v8)
    v5 -= (^(v5))
    for fi9 := 0; fi9 < 2; fi9++ {
        for fi10 := 0; fi10 < 19; fi10++ {
            var v11 int = fi10
            if v2 {
                switch ("hello" + (("\n\t" + "hello") + "a")) {
                case "\n\t":
                    var v12 int16 = 42
                    _ = v12
                case "a":
                    v13 := GenS1{F0: (^(v3)), F1: (0 == 18446744073709551615), F2: 2}
                    _ = v13
                case "hello":
                    switch "a" {
                    case "\n\t":
                        var v14 uint16 = (^(v3))
                        _ = v14
                    case "a":
                        switch (("" + "\n\t") + ("hello" + "foo bar")) {
                        case "foo bar":
                            v8 -= v8
                        case "\n\t":
                            v2 = v2
                        }
                    case "foo bar":
                        var v15 uint32 = uint32(v1)
                        _ = v15
                    }
                }
            } else {
                var v16 uint16 = v3
                v2 = (((((0.1 * v8) / v8) + ((v8 * (float64(-2.25) + (v8 - v8))) / (v8 / v8))) / 0.1) <= (1.5 / ((v8 / v8) * v8)))
                switch 4294967295 {
                case 951:
                    delete(v4, v11)
                    _ = len(v7)
                    var v17 string = ("foo bar" + ("\n\t" + "a"))
                    v5 += v16
                    _ = v17
                }
                v1 = (v1 & v1)
                v2 = (2 < 255)
                _ = v16
            }
            var v18 uint = (v1 / ((255) | 1))
            _ = fi10
            _ = v11
            _ = v18
        }
        if true {
            var v19 float32 = float32(((v8 * (((v8 - v8) - (v8 - (((0.1 / 3.14159) + (v8 - 0.0)) * v8))) - (1e10 + float64(float32(1.5))))) / (float64(((1e10 + 1.5) * float32(v8))) - (v8 + 0.1))))
            v8 = (v8 / 1.0)
            fmt.Println(v2)
            _ = v19
        } else {
            v3 = 822
            v3 += (^(v5))
            if false {
                _ = len(v4)
                v5 = (v5 / (((v5 - v5)) | 1))
            } else {
                switch fi9 {
                case 128:
                    v1 += 255
                case 32767:
                    if ((v5 + 0) == v5) {
                        fmt.Println(v5)
                    }
                case 0:
                    for fi20 := 0; fi20 < 6; fi20++ {
                        switch "" {
                        case "foo bar":
                            v5 += (v3 / ((v3) | 1))
                        case "a":
                            v3 = v5
                        }
                        v21 := GenS1{F0: (v5 << ((v1) & 63)), F1: (true && (v2 && (v1 < v1))), F2: uint64(v3)}
                        var v22 int8 = 100
                        _ = fi20
                        _ = v21
                        _ = v22
                    }
                default:
                    v23 := GenS2{F0: -832}
                    _ = v23
                }
                for fi24 := 0; fi24 < 3; fi24++ {
                    var v25 int16 = int16(v3)
                    var v26 int64 = 9223372036854775807
                    if (100 >= 100) {
                        var v27 uint8 = 2
                        var v28 float32 = 1e10
                        v25 = v25
                        _ = v27
                        _ = v28
                    }
                    _ = fi24
                    _ = v25
                    _ = v26
                }
                var v29 int8 = 100
                _ = v29
            }
        }
        _ = fi9
    }
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    v4[-992] = v2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
