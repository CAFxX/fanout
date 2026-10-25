package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint }
type GenS2 struct { F0 uint16; F1 uint64; F2 int16 }

func main() {
    v1 := make(map[int]int)
    var v2 uint16 = 42
    var v3 float32 = (1e10 * 0.1)
    v4 := []int{}
    _ = len(v4)
    switch "" {
    case "foo bar":
        var v5 int8 = 1
        _ = v5
    case "\n\t":
        v3 -= (0.0 - 1e10)
    case "a":
        if false {
            if false {
                fmt.Println(v3)
            } else {
                v3 -= ((v3 - (v3 * v3)) / v3)
                v4 = append(v4, 32768)
                _ = len(v1)
            }
        }
    default:
        v3 = 1e10
    }
    var v6 bool = (-2147483648 != -1)
    if v6 {
        v7 := GenS1{F0: -2147483648, F1: 634}
        v2 = v2
        _ = v7
    } else {
        v3 += 0.1
        v8 := GenS1{F0: 0, F1: 288}
        if (("hello" > "hello") || v6) {
            v6 = true
            if (189 == 42) {
                var v9 uint16 = 0
                _ = v9
            } else {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                v2 = (v2 >> ((42) & 63))
            }
            var v10 uint64 = uint64(v2)
            _ = v10
        }
        _ = v8
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if v6 {
        for fi11 := 0; fi11 < 17; fi11++ {
            if v6 {
                var v12 uint = 2
                _ = v12
            } else {
                _ = len(v1)
                v1[(fi11 + (fi11 % (((fi11 >> ((255) & 63))) | 1)))] = fi11
                if ((!(128 >= 133)) && v6) {
                    v4 = append(v4, -1)
                    v13 := GenS1{F0: -2147483648, F1: 18446744073709551615}
                    _ = v13
                } else {
                    v2 += v2
                    _ = len(v1)
                }
            }
            _ = fi11
        }
        v14 := GenS2{F0: v2, F1: 408, F2: 849}
        _ = v14
    }
    var v15 uint16 = v2
    var v16 int64 = 9223372036854775807
    v15 = 2
    var v17 float64 = -2.25
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v6
}
