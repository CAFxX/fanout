package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8 }

func main() {
    var v1 uint16 = 65535
    var v2 uint8 = 128
    v3 := make(map[string]int)
    var v4 float64 = (0.0 / (0.0 + -2.25))
    v5 := []int{-2147483648, 32768, -9223372036854775808}
    fmt.Println(v4)
    if false {
        _ = len(v3)
        var v6 int32 = -2147483648
        _ = v6
    } else {
        var v7 int32 = 2147483647
        _ = len(v5)
        for fi8 := 0; fi8 < 19; fi8++ {
            switch (fi8 - (fi8 / ((32767) | 1))) {
            case 2147483647:
                if false {
                    var v9 int = (fi8 >> ((18446744073709551615) & 63))
                    _ = v9
                }
            default:
                if ((^(fi8)) <= -706) {
                    v7 = v7
                }
            }
            _ = fi8
        }
        _ = v7
    }
    v2 = (v2 - (v2 * 255))
    var v10 uint16 = (v1 | (v1 | (^(v1))))
    for fi11 := 0; fi11 < 9; fi11++ {
        for fi12 := 0; fi12 < 15; fi12++ {
            if false {
                v13 := GenS1{F0: "foo bar", F1: -1}
                fmt.Println(v10, v1)
                var v14 bool = true
                _ = v13
                _ = v14
            } else {
                if false {
                    v15 := GenS1{F0: "hello", F1: -128}
                    _ = v15
                }
            }
            _ = fi12
        }
        v16 := GenS1{F0: "hello", F1: 48}
        _ = fi11
        _ = v16
    }
    if false {
        var v17 uint = 4294967295
        _ = v17
    } else {
        var v18 int16 = 1
        _ = v18
    }
    fmt.Println(v2)
    for fi19 := 0; fi19 < 19; fi19++ {
        var v20 int64 = 1
        fmt.Println(v2, v10)
        _ = fi19
        _ = v20
    }
    var v21 float32 = (1.5 + 0.1)
    fmt.Println(v1, v2, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
