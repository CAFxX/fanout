package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64 }

func main() {
    v1 := make(map[uint32]int)
    v2 := []string{("foo bar" + "\n\t")}
    v3 := make(map[int]string)
    switch 127 {
    case 340:
        var v4 uint = 779
        _ = v4
    default:
        var v5 float64 = ((((1.5 * float64(-2.25)) - (-2.25 / 1.0)) * 0.1) + 3.14159)
        _ = v5
    }
    var v6 int = 9223372036854775807
    v2 = append(v2, "a")
    var v7 uint64 = 59
    switch "" {
    case "foo bar":
        v7 = v7
    case "hello":
        var v8 uint64 = (v7 & (v7 * v7))
        _ = v8
    case "a":
        v9 := GenS1{F0: 0, F1: (v7 - (^(v7)))}
        _ = v9
    }
    v6 += v6
    var v10 uint8 = 82
    var v11 int32 = 435
    switch "a" {
    case "a":
        for fi12 := 0; fi12 < 2; fi12++ {
            switch -2147483648 {
            case -837:
                v10 -= 251
            case -9223372036854775808:
                fmt.Println(v7)
            case 127:
                delete(v3, (v6 >> ((822) & 63)))
            }
            _ = fi12
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v6
    _ = v7
}
