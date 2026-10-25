package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64 }
type GenS2 struct { F0 int64; F1 uint32 }
type GenS3 struct { F0 int32; F1 int32 }

func main() {
    v1 := []int{}
    var v2 float64 = float64(3.14159)
    var v3 uint8 = 29
    v4 := make(map[uint32]bool)
    var v5 int16 = 346
    v6 := []uint32{}
    for fi7 := 0; fi7 < 13; fi7++ {
        fmt.Println(v2, v2)
        _ = fi7
    }
    v8 := GenS1{F0: "hello", F1: 2}
    if ((!("\n\t" != "")) || (92 >= 127)) {
        for fi9 := 0; fi9 < 19; fi9++ {
            delete(v4, 1)
            for fi10 := 0; fi10 < 8; fi10++ {
                switch fi10 {
                case 127:
                    var v11 bool = (float32(((3.14159 / v2) - v2)) >= -2.25)
                    _ = v11
                case 32767:
                    _ = len(v4)
                default:
                    switch ("a" + ((("\n\t" + "") + "hello") + "\n\t")) {
                    case "foo bar":
                        v5 -= (v5 + v5)
                    case "hello":
                        var v12 uint = 1
                        _ = v12
                    default:
                        v13 := GenS2{F0: 402, F1: 933}
                        _ = v13
                    }
                }
                v5 += (-(v5))
                _ = fi10
            }
            v4[4294967295] = true
            _ = fi9
        }
        var v14 string = ""
        v1 = append(v1, 0)
        _ = v14
    }
    _ = len(v1)
    for fi15 := 0; fi15 < 15; fi15++ {
        var v16 float64 = 0.1
        _ = fi15
        _ = v16
    }
    var v17 int = 127
    if true {
        v2 = v2
        switch "\n\t" {
        case "\n\t":
            v6 = append(v6, 2)
        case "a":
            var v18 bool = true
            _ = v18
        case "":
            for fi19 := 0; fi19 < 12; fi19++ {
                for fi20 := 0; fi20 < 14; fi20++ {
                    if len(v1) > 0 {
                        v1[0] = (v17 / ((128) | 1))
                    }
                    _ = fi20
                }
                _ = fi19
            }
        }
    }
    var v21 float32 = -2.25
    switch ("foo bar" + "hello") {
    case "foo bar":
        var v22 string = ("" + "hello")
        _ = v22
    default:
        var v23 uint16 = 42
        _ = v23
    }
    v24 := GenS3{F0: -1, F1: -2147483648}
    v21 = -2.25
    var v25 uint16 = 42
    var v26 int8 = 0
    fmt.Println(v21)
    switch "" {
    case "\n\t":
        switch (v17 + v17) {
        case -152:
            v3 = v3
        case 4294967295:
            v5 = v5
        case 32768:
            var v27 int64 = int64(v17)
            _ = v27
        }
    }
    if (18446744073709551615 == 1) {
        var v28 uint64 = 18446744073709551615
        v29 := GenS2{F0: 42, F1: 42}
        _ = v28
        _ = v29
    }
    v1 = append(v1, 9223372036854775807)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v21
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
