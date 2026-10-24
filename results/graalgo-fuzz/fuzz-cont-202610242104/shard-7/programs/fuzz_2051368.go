package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint }

func main() {
    var v1 int32 = 487
    var v2 float64 = 0.1
    var v3 uint16 = 0
    v1 = 0
    var v4 uint32 = 42
    v3 -= (^(v3))
    fmt.Println(v2, v1, v2)
    if ("foo bar" != "") {
        v1 -= v1
        v4 -= v4
        var v5 uint16 = v3
        _ = v5
    } else {
        switch "hello" {
        case "hello":
            v2 += v2
        case "":
            switch 0 {
            case -574:
                fmt.Println(v4, v2, v4)
            case 347:
                v6 := []int{-1, 1}
                _ = v6
            case -9223372036854775808:
                for fi7 := 0; fi7 < 1; fi7++ {
                    v8 := GenS1{F0: int32(v3), F1: 472}
                    _ = fi7
                    _ = v8
                }
            default:
                v9 := GenS1{F0: (-(v1)), F1: 671}
                _ = v9
            }
        case "a":
            v3 += (^(v3))
        }
    }
    v2 -= 3.14159
    switch 2147483647 {
    case 32767:
        switch 128 {
        case 32768:
            v1 += 42
        case 127:
            var v10 bool = (true || false)
            _ = v10
        default:
            var v11 uint = 42
            _ = v11
        }
    }
    v12 := []int{977, 42, 1}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
}
