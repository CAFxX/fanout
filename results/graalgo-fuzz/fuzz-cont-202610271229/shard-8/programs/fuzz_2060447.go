package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint32; F1 int; F2 int16 }
type GenS3 struct { F0 string; F1 uint64; F2 uint8 }

func main() {
    v1 := make(map[string]string)
    v2 := make(map[int]int)
    v3 := make(map[int]string)
    var v4 uint64 = 4294967295
    var v5 float64 = 0.0
    v4 = v4
    switch "" {
    case "foo bar":
        for fi6 := 0; fi6 < 9; fi6++ {
            v5 = v5
            fmt.Println(v5, fi6, v5)
            switch "\n\t" {
            case "\n\t":
                v5 = 0.0
            case "a":
                var v7 uint = 18446744073709551615
                _ = v7
            }
            _ = fi6
        }
    case "\n\t":
        v4 -= v4
    case "":
        if true {
            fmt.Println(v5, v5)
            v4 -= (^(v4))
            v5 -= float64((1e10 * (0.1 / (float32(v5) + 0.1))))
        } else {
            v8 := []int{42, 128, 32768}
            v5 = ((v5 * 0.0) - (0.1 - 3.14159))
            for fi9 := 0; fi9 < 12; fi9++ {
                var v10 int64 = -11
                _ = fi9
                _ = v10
            }
            _ = v8
        }
    }
    var v11 float64 = (1.5 + v5)
    v11 = v5
    v5 = v11
    var v12 uint32 = 2147483647
    v13 := []int{-1}
    switch "\n\t" {
    case "hello":
        v12 = (v12 << ((1) & 63))
    }
    v11 += (3.14159 - (((0.0 * -2.25) + 3.14159) / v11))
    v13 = append(v13, 9223372036854775807)
    v12 = v12
    v14 := GenS1{F0: (false || (float64(0.0) >= (v5 - v11)))}
    switch "hello" {
    case "\n\t":
        v15 := GenS2{F0: 2147483648, F1: -126, F2: -1}
        _ = v15
    case "hello":
        var v16 int16 = 32767
        _ = v16
    case "":
        v5 = 0.0
    default:
        v17 := GenS3{F0: ("" + "hello"), F1: (^(v4)), F2: 127}
        _ = v17
    }
    v18 := v2[42]
    if (42 > 128) {
        for fi19 := 0; fi19 < 11; fi19++ {
            var v20 bool = false
            v21 := GenS1{F0: true}
            _ = fi19
            _ = v20
            _ = v21
        }
        switch v18 {
        case -169:
            switch (v18 * -769) {
            case -526:
                v4 = 0
            case -476:
                v13 = append(v13, 127)
            }
        case 42:
            v5 -= -2.25
        case -1:
            var v22 uint16 = 255
            _ = v22
        default:
            var v23 string = "hello"
            _ = v23
        }
    }
    var v24 uint8 = 139
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v18
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
