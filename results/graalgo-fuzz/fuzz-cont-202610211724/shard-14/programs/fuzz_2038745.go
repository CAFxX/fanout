package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64; F2 int64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint32 }

func main() {
    v1 := make(map[uint32]int)
    var v2 float32 = 1e10
    v3 := make(map[uint32]bool)
    switch 42 {
    case 42:
        var v4 float32 = 0.1
        _ = v4
    default:
        var v5 string = ((((("a" + "foo bar") + "\n\t") + "\n\t") + "") + "foo bar")
        _ = v5
    }
    fmt.Println(v2)
    v6 := []int{797}
    fmt.Println(v2)
    switch 569 {
    case -784:
        v7 := GenS1{F0: 1, F1: -1, F2: -9223372036854775808}
        _ = v7
    case 9223372036854775807:
        var v8 int32 = 83
        _ = v8
    default:
        var v9 uint16 = 255
        _ = v9
    }
    var v10 uint = 42
    v11 := GenS2{F0: 4294967295}
    switch "" {
    case "\n\t":
        v10 -= 65535
    default:
        switch 1 {
        case -623:
            v2 -= (-2.25 / v2)
        case 203:
            var v12 int16 = 1000
            _ = v12
        default:
            if (!false) {
                v2 += (v2 - (v2 - v2))
                for fi13 := 0; fi13 < 15; fi13++ {
                    v6 = append(v6, fi13)
                    _ = fi13
                }
                switch "hello" {
                case "foo bar":
                    var v14 float64 = ((1.5 * (1.5 - float64(0.0))) - 0.1)
                    _ = v14
                }
            }
        }
    }
    v15 := v3[2147483647]
    _ = len(v1)
    v16 := GenS3{F0: 458}
    fmt.Println(v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v6
}
