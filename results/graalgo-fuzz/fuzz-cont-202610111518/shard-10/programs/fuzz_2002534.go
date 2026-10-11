package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int; F2 uint16 }
type GenS2 struct { F0 uint16; F1 uint8 }
func genFn0(p0 int16) int64 {
    var v4 float64 = 3.14159
    fmt.Println(p0, p0)
    _ = p0
    _ = v4
    return 9223372036854775807
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int8 = 100
    var v3 uint8 = 42
    switch "\n\t" {
    case "":
        v5 := []int{-802, 42}
        _ = v5
    case "a":
        v6 := []int{128, 32767, 255}
        _ = v6
    case "hello":
        v3 -= 104
    default:
        var v7 uint8 = (v3 / ((v3) | 1))
        _ = v7
    }
    fmt.Println(v2, v3)
    if true {
        if true {
            switch "foo bar" {
            case "hello":
                v8 := []int{385}
                _ = v8
            case "foo bar":
                v1[0] = ((v3 / ((v3) | 1)) >= 127)
            }
        }
        switch -9223372036854775808 {
        case 2147483647:
            for fi9 := 0; fi9 < 19; fi9++ {
                v2 = 42
                _ = fi9
            }
        case -9223372036854775808:
            fmt.Println(v2, v3)
        default:
            fmt.Println(v3)
        }
    }
    v3 = (v3 + (v3 ^ (v3 * (v3 - v3))))
    v10 := GenS1{F0: 42, F1: -166, F2: 32767}
    v11 := GenS2{F0: 2, F1: (v3 * v3)}
    fmt.Println(v2, v3)
    var v12 float64 = -2.25
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
}
