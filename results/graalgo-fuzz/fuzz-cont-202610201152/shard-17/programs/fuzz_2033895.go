package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8; F2 int64 }
type GenS2 struct { F0 int16; F1 bool }

func main() {
    v1 := make(map[uint32]bool)
    var v2 int32 = 629
    var v3 uint16 = 32767
    v4 := make(map[string]int)
    var v5 uint32 = 0
    v6 := make(map[int]string)
    v7 := GenS1{F0: -1, F1: 128, F2: 9223372036854775807}
    v3 += 0
    v8 := v6[232]
    _ = len(v1)
    v2 -= v2
    switch -817 {
    case -9223372036854775808:
        var v9 float32 = 1e10
        _ = v9
    case 590:
        v2 -= (v2 + (-(v2)))
    case 2147483647:
        switch -2147483648 {
        case 32768:
            v1[(v5 * 2147483648)] = (!true)
        default:
            v10 := GenS1{F0: 459, F1: 241, F2: 804}
            _ = v10
        }
    }
    var v11 int32 = 1
    switch "foo bar" {
    case "hello":
        var v12 uint64 = 680
        _ = v12
    case "\n\t":
        delete(v1, 4294967295)
    case "foo bar":
        fmt.Println(v2)
    default:
        v13 := []int{806, 4294967295}
        _ = v13
    }
    _ = len(v4)
    var v14 uint = 4294967295
    switch -1 {
    case 42:
        if (!false) {
            v15 := []int{4294967295}
            var v16 int64 = 1
            fmt.Println(v14, v3, v3)
            _ = v15
            _ = v16
        }
    case -590:
        switch (v8 + v8) {
        case "a":
            if (-231 != (v2 | (v2 ^ (^(v2))))) {
                v17 := []int{32767, 106}
                v18 := GenS1{F0: 9223372036854775807, F1: 127, F2: 42}
                _ = v17
                _ = v18
            }
        case "foo bar":
            delete(v4, "")
        case "hello":
            v19 := []int{4294967295, -2147483648, 218, 32768}
            _ = v19
        default:
            v14 += 4294967295
        }
    case -2147483648:
        var v20 int32 = (^(v11))
        _ = v20
    default:
        v21 := GenS2{F0: -32768, F1: false}
        _ = v21
    }
    v5 += 42
    v22 := GenS2{F0: 32767, F1: (true && false)}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v14
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
