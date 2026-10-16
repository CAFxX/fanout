package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int }

func main() {
    v1 := []string{}
    v2 := make(map[int]bool)
    var v3 string = ""
    v4 := make(map[int]bool)
    v5 := make(map[int]uint32)
    v6 := GenS1{F0: 917}
    fmt.Println(v3)
    var v7 float32 = -2.25
    v8 := GenS1{F0: 936}
    switch -2147483648 {
    case 9223372036854775807:
        v9 := v5[0]
        _ = v9
    case 900:
        v7 += (v7 * -2.25)
    }
    switch 32767 {
    case 128:
        if false {
            fmt.Println(v7, v3)
            v10 := GenS1{F0: 0}
            _ = v10
        } else {
            for fi11 := 0; fi11 < 20; fi11++ {
                v7 -= (float32(float64((float32(((float64(3.14159) + (-2.25 * 0.1)) + 0.1)) / 1.0))) / 3.14159)
                var v12 uint16 = 1
                v1 = append(v1, v3)
                _ = fi11
                _ = v12
            }
        }
    case -9223372036854775808:
        fmt.Println(v3, v7)
    case -2147483648:
        switch 2147483647 {
        case 128:
            delete(v4, 1)
        case 32767:
            if (((3.14159 * (0.1 + v7)) + 0.1) != (v7 - v7)) {
                v13 := GenS1{F0: 255}
                _ = v13
            }
        default:
            if (18446744073709551615 > 4294967295) {
                fmt.Println(v7, v3)
                v14 := GenS2{F0: 128}
                var v15 int64 = 1
                _ = v14
                _ = v15
            }
        }
    }
    var v16 bool = (-32768 >= 32767)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v17 int16 = 42
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
