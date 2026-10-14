package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32 }
type GenS2 struct { F0 int16; F1 uint16; F2 uint64 }
type GenS3 struct { F0 int8; F1 uint16 }
type GenS4 struct { F0 uint32; F1 uint8; F2 uint }
func genFn0(p0 uint16, p1 int) int {
    var v5 float64 = (-2.25 / float64((1.5 - -2.25)))
    fmt.Println(v5, v5, p0)
    v6 := GenS1{F0: 2147483647, F1: 268}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 128
}

func main() {
    var v1 string = "a"
    v2 := make(map[string]int)
    var v3 int32 = 2147483647
    var v4 float64 = 1.5
    fmt.Println(v3, v4, v1)
    fmt.Println(v1)
    v7 := []int{genFn0(65535, -9223372036854775808)}
    v7 = append(v7, 0)
    fmt.Println(v1)
    for fi8 := 0; fi8 < 3; fi8++ {
        var v9 float64 = v4
        if (1 > 0) {
            var v10 int = (fi8 ^ -47)
            switch genFn0(907, v10) {
            case 127:
                fmt.Println(fi8, v1)
            case 2147483647:
                v11 := GenS2{F0: -1, F1: 65535, F2: 18446744073709551615}
                _ = v11
            default:
                v12 := GenS3{F0: -128, F1: 1}
                _ = v12
            }
            switch "a" {
            case "\n\t":
                for fi13 := 0; fi13 < 13; fi13++ {
                    switch v1 {
                    case "\n\t":
                        v1 = "foo bar"
                    }
                    fmt.Println(fi8, v9, v9)
                    var v14 float64 = v9
                    _ = fi13
                    _ = v14
                }
            case "hello":
                var v15 float32 = (0.1 / (float32(v9) / (0.1 * (1e10 + 3.14159))))
                _ = v15
            }
            _ = v10
        } else {
            fmt.Println(fi8, v3, v1)
            fmt.Println(v4, fi8)
        }
        v4 -= 3.14159
        _ = fi8
        _ = v9
    }
    switch ((v1 + ((v1 + "foo bar") + ("" + ("" + v1)))) + (v1 + v1)) {
    case "foo bar":
        v16 := GenS4{F0: 2, F1: 128, F2: 255}
        _ = v16
    case "":
        v17 := GenS4{F0: 2147483648, F1: 128, F2: 0}
        _ = v17
    }
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
