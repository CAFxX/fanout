package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64 }
type GenS2 struct { F0 string; F1 uint8 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int16; F1 int; F2 string }
type GenS5 struct { F0 int8; F1 uint32 }

func main() {
    var v1 uint = 110
    v2 := make(map[int]string)
    v3 := make(map[string]uint32)
    var v4 string = "foo bar"
    var v5 uint8 = 99
    v6 := make(map[string]uint32)
    if (!false) {
        fmt.Println(v5, v1, v5)
        if (1 == 2147483647) {
            fmt.Println(v4, v4, v4)
            fmt.Println(v1, v1)
            var v7 float32 = float32(1.5)
            _ = v7
        } else {
            var v8 int = 424
            _ = v8
        }
        switch 42 {
        case 0:
            var v9 int8 = -1
            _ = v9
        }
    } else {
        var v10 uint8 = 146
        var v11 float32 = (0.0 + ((float32(3.14159) + float32(1.5)) + float32(((3.14159 - (float64(-2.25) - 0.0)) - (0.1 * (0.1 + (((float64(((float32((3.14159 * 1.5)) * (((-2.25 + 0.0) + 3.14159) + 1.5)) + 0.1)) - ((-2.25 + (1e10 / ((0.1 * 1e10) / 1e10))) * 3.14159)) / (float64(3.14159) + 3.14159)) / 1.0)))))))
        v4 = ((v4 + v4) + ((v4 + (v4 + v4)) + v4))
        _ = v10
        _ = v11
    }
    switch "\n\t" {
    case "hello":
        v12 := GenS1{F0: 32767, F1: 42}
        _ = v12
    case "foo bar":
        for fi13 := 0; fi13 < 8; fi13++ {
            for fi14 := 0; fi14 < 20; fi14++ {
                v1 -= (v1 & 73)
                _ = fi14
            }
            var v15 uint16 = 2
            _ = fi13
            _ = v15
        }
    default:
        if false {
            var v16 int8 = 10
            v17 := []int{32768}
            _ = v16
            _ = v17
        }
    }
    fmt.Println(v1)
    v4 = v4
    var v18 float32 = (0.0 - 3.14159)
    fmt.Println(v18)
    for fi19 := 0; fi19 < 4; fi19++ {
        fmt.Println(v5, v1, v1)
        _ = fi19
    }
    v18 += 0.0
    if (!true) {
        if true {
            if (65535 < 653) {
                switch "a" {
                case "\n\t":
                    for fi20 := 0; fi20 < 12; fi20++ {
                        v5 += v5
                        _ = fi20
                    }
                case "hello":
                    var v21 float64 = (0.0 / 1.0)
                    _ = v21
                }
                var v22 float32 = -2.25
                _ = v22
            } else {
                var v23 uint64 = 42
                _ = v23
            }
        } else {
            if false {
                v1 -= 42
            }
            fmt.Println(v5, v18, v18)
            v5 = 128
        }
        v4 = v4
    }
    for fi24 := 0; fi24 < 1; fi24++ {
        v5 -= 64
        v25 := GenS2{F0: "hello", F1: 234}
        for fi26 := 0; fi26 < 1; fi26++ {
            var v27 int64 = 804
            if false {
                v5 -= 0
            }
            v28 := GenS3{F0: 427}
            v29 := GenS1{F0: 1000, F1: 0}
            _ = fi26
            _ = v27
            _ = v28
            _ = v29
        }
        _ = fi24
        _ = v25
    }
    v30 := GenS4{F0: 32767, F1: 4294967295, F2: v4}
    v1 = 255
    v1 += v1
    v31 := GenS5{F0: 1, F1: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
