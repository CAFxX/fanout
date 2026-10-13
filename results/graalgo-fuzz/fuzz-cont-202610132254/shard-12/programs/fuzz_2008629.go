package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16 }
type GenS2 struct { F0 bool; F1 int; F2 string }

func main() {
    var v1 uint8 = 132
    v2 := make(map[int]string)
    v3 := make(map[uint32]bool)
    var v4 uint32 = 0
    var v5 int16 = 42
    v6 := make(map[uint32]int)
    v4 = v4
    switch "\n\t" {
    case "hello":
        v7 := []int{-1, 4294967295, -286}
        _ = v7
    case "\n\t":
        v1 -= (v1 & 2)
    case "foo bar":
        var v8 int32 = -1
        _ = v8
    default:
        v5 += (-(v5))
    }
    if true {
        v5 -= -207
    } else {
        _ = len(v6)
        switch ("foo bar" + ("hello" + ("foo bar" + "foo bar"))) {
        case "hello":
            var v9 int = 9223372036854775807
            _ = v9
        case "foo bar":
            for fi10 := 0; fi10 < 17; fi10++ {
                switch "foo bar" {
                case "a":
                    switch "a" {
                    case "foo bar":
                        var v11 int16 = v5
                        _ = v11
                    case "":
                        for fi12 := 0; fi12 < 11; fi12++ {
                            var v13 uint32 = 444
                            _ = fi12
                            _ = v13
                        }
                    case "\n\t":
                        fmt.Println(v5, v1, v5)
                    }
                case "\n\t":
                    for fi14 := 0; fi14 < 16; fi14++ {
                        fmt.Println(v5, v5, v1)
                        _ = fi14
                    }
                case "foo bar":
                    if true {
                        for fi15 := 0; fi15 < 7; fi15++ {
                            v16 := GenS1{F0: -32768, F1: (v5 + v5)}
                            fmt.Println(v1, fi10)
                            var v17 float64 = (1e10 + 0.1)
                            _ = fi15
                            _ = v16
                            _ = v17
                        }
                    }
                default:
                    fmt.Println(v5)
                }
                if true {
                    fmt.Println(v4, v4, fi10)
                    v1 = (v1 % ((128) | 1))
                } else {
                    var v18 int16 = v5
                    switch fi10 {
                    case 127:
                        delete(v2, fi10)
                    }
                    var v19 int8 = -128
                    _ = v18
                    _ = v19
                }
                switch "hello" {
                case "a":
                    v20 := []int{127, (fi10 * (fi10 | -244)), fi10}
                    _ = v20
                case "foo bar":
                    v5 = 0
                case "":
                    v1 = v1
                }
                _ = fi10
            }
        default:
            v21 := []int{-499, -2147483648, -603, 49}
            _ = v21
        }
        fmt.Println(v1, v5)
    }
    for fi22 := 0; fi22 < 18; fi22++ {
        var v23 int = fi22
        _ = fi22
        _ = v23
    }
    if false {
        v4 += uint32(v1)
        v2[127] = "hello"
        v24 := GenS1{F0: -1, F1: (v5 >> ((18446744073709551615) & 63))}
        _ = v24
    } else {
        var v25 bool = (!true)
        _ = v25
    }
    v4 += v4
    v1 = (v1 - (v1 / (((v1 & 128)) | 1)))
    var v26 uint = 2
    var v27 uint8 = 1
    v5 = v5
    v28 := []int{206, -856}
    v5 += v5
    if (743 < -928) {
        _ = len(v3)
    } else {
        delete(v3, v4)
        if false {
            v29 := GenS1{F0: 32767, F1: v5}
            var v30 float64 = (-2.25 - float64(0.0))
            _ = v29
            _ = v30
        }
        if (!false) {
            v31 := GenS1{F0: (^(v5)), F1: (v5 >> ((797) & 63))}
            _ = v31
        }
    }
    switch "" {
    case "":
        v32 := GenS2{F0: (true && (((v27 >> ((0) & 63)) >= (v27 | v1)) && (0.1 > 1e10))), F1: 882, F2: (("" + "hello") + "\n\t")}
        _ = v32
    }
    var v33 int32 = 1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
