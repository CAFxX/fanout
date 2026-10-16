package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32 }
type GenS2 struct { F0 uint; F1 int8 }

func main() {
    v1 := []uint32{2147483648, 0, 2147483648, 271}
    var v2 int32 = 2147483647
    v3 := []int{2147483647, 127, 1, -721}
    v4 := []int{127, 32768, 32768}
    v5 := make(map[string]string)
    fmt.Println(v2)
    v2 = 42
    if false {
        switch "foo bar" {
        case "\n\t":
            v2 = v2
        case "":
            v2 = 470
        default:
            switch 32767 {
            case 128:
                if (-256 > -9223372036854775808) {
                    fmt.Println(v2)
                    fmt.Println(v2)
                    fmt.Println(v2)
                }
            case 9223372036854775807:
                v2 += 1
            default:
                var v6 bool = (!true)
                _ = v6
            }
        }
    } else {
        var v7 uint8 = 0
        fmt.Println(v7)
        _ = v7
    }
    if true {
        if len(v3) > 0 {
            v3[0] = -725
        }
        var v8 uint32 = 551
        v2 -= v2
        _ = v8
    }
    v2 += v2
    if (!true) {
        v9 := GenS1{F0: (true || (false || false)), F1: v2}
        _ = v9
    } else {
        var v10 uint8 = 81
        fmt.Println(v10)
        _ = v10
    }
    var v11 float64 = 3.14159
    v12 := v5[""]
    var v13 int64 = -459
    v13 -= v13
    v13 = 9223372036854775807
    v14 := GenS2{F0: 789, F1: 100}
    v15 := GenS1{F0: false, F1: (v2 ^ v2)}
    var v16 uint64 = 4294967295
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
