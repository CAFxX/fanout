package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16 }
type GenS2 struct { F0 uint }

func main() {
    var v1 uint = 42
    v2 := make(map[int]bool)
    var v3 uint16 = 1
    var v4 uint16 = (v3 << (((v1 / ((uint(v3)) | 1))) & 63))
    var v5 uint32 = uint32(v3)
    var v6 string = (("foo bar" + "") + "\n\t")
    for fi7 := 0; fi7 < 16; fi7++ {
        if (true && (false && ((0 > -59) || (-2147483648 == -853)))) {
            if ((57 > 210) && false) {
                v8 := GenS1{F0: 226, F1: 344}
                var v9 float64 = (0.1 * 1e10)
                _ = v8
                _ = v9
            }
            var v10 uint8 = 1
            _ = v10
        } else {
            fmt.Println(v6, v1)
            if (-2 >= -1) {
                v4 += v4
                if ((v5 / (((v5 | (v5 >> ((v1) & 63)))) | 1)) > v5) {
                    v11 := []int{128, fi7}
                    _ = v11
                } else {
                    var v12 string = "a"
                    v12 = v12
                    _ = v12
                }
            } else {
                fmt.Println(v3, fi7, fi7)
            }
            v13 := GenS2{F0: (v1 ^ (v1 - 4294967295))}
            _ = v13
        }
        _ = fi7
    }
    var v14 int8 = -9
    switch "" {
    case "foo bar":
        v15 := []int{1}
        _ = v15
    case "\n\t":
        v2[135] = (((float32((1.5 + 0.0)) / 0.1) + (0.0 + (1.5 + (1e10 - (1.5 + 1.5))))) >= 3.14159)
    case "hello":
        v16 := []int{128, 128, 127}
        _ = v16
    }
    fmt.Println(v14, v14, v5)
    var v17 bool = (true || (!true))
    fmt.Println(v17)
    switch v6 {
    case "hello":
        fmt.Println(v5, v14, v3)
    case "foo bar":
        v18 := GenS2{F0: (v1 ^ v1)}
        _ = v18
    default:
        v5 = v5
    }
    for fi19 := 0; fi19 < 15; fi19++ {
        switch v6 {
        case "hello":
            v1 = (v1 << ((v1) & 63))
        case "\n\t":
            v1 = v1
        case "foo bar":
            if v17 {
                switch fi19 {
                case -842:
                    var v20 int = 42
                    _ = v20
                case 1:
                    _ = len(v2)
                }
            }
        }
        fmt.Println(v17)
        _ = fi19
    }
    fmt.Println(v14)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
