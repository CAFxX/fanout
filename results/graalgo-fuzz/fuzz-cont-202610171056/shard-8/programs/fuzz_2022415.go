package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16 }

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[int]uint32)
    var v3 uint32 = 718
    var v4 int8 = 0
    v5 := make(map[uint32]bool)
    if false {
        v3 = 652
    }
    var v6 uint8 = 127
    fmt.Println(v6, v3)
    var v7 float64 = float64((1e10 - float32((1.5 / 3.14159))))
    if ((true && (-773 >= -1)) && (true && ((v6 > (v6 << ((4294967295) & 63))) || ((-(v4)) != -77)))) {
        v3 += (v3 >> ((677) & 63))
        switch "hello" {
        case "foo bar":
            v7 = v7
        case "\n\t":
            v8 := []int{4294967295, 32767}
            _ = v8
        case "hello":
            fmt.Println(v6)
        default:
            v4 += 0
        }
    } else {
        for fi9 := 0; fi9 < 1; fi9++ {
            if (-2.25 < 0.1) {
                v10 := GenS1{F0: 1, F1: 1000}
                v6 = v6
                _ = v10
            } else {
                fmt.Println(v4, v4, v6)
            }
            for fi11 := 0; fi11 < 12; fi11++ {
                delete(v1, 1)
                v4 += v4
                _ = fi11
            }
            _ = fi9
        }
    }
    for fi12 := 0; fi12 < 1; fi12++ {
        var v13 int64 = -758
        if false {
            switch "foo bar" {
            case "hello":
                for fi14 := 0; fi14 < 12; fi14++ {
                    switch fi12 {
                    case -9223372036854775808:
                        if (((!(428 >= -410)) || (1e10 != (3.14159 + (1.5 / 1.0)))) || (false && false)) {
                            v7 = ((0.0 * v7) / v7)
                            var v15 uint32 = (v3 + 2147483648)
                            _ = v15
                        }
                    case 32768:
                        fmt.Println(v3, fi14, v7)
                    case -750:
                        var v16 uint16 = 1
                        _ = v16
                    }
                    _ = fi14
                }
            }
        } else {
            fmt.Println(v13, v4)
            if (v6 == v6) {
                fmt.Println(v6)
                if (4294967295 != fi12) {
                    for fi17 := 0; fi17 < 3; fi17++ {
                        fmt.Println(v13)
                        _ = fi17
                    }
                    v7 += v7
                }
                v7 = v7
            }
        }
        _ = fi12
        _ = v13
    }
    _ = len(v5)
    v18 := GenS1{F0: 810, F1: 1}
    var v19 int64 = 0
    var v20 int32 = -1
    v21 := []int{42, 9223372036854775807, 2147483647, 9223372036854775807}
    var v22 int = 0
    var v23 uint16 = 574
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
