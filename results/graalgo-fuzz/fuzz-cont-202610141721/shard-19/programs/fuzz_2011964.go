package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint8; F1 uint }

func main() {
    var v1 string = "\n\t"
    var v2 int8 = 0
    var v3 string = "\n\t"
    v4 := make(map[string]string)
    v1 = v3
    fmt.Println(v3, v3, v1)
    var v5 uint32 = 4294967295
    switch 684 {
    case -195:
        if true {
            v6 := []int{9223372036854775807, int(v2), int(v5), 0}
            fmt.Println(v1, v3, v2)
            _ = v6
        } else {
            fmt.Println(v3, v1)
            var v7 uint16 = 2
            _ = v7
        }
    case -9223372036854775808:
        var v8 int = 2147483647
        _ = v8
    default:
        switch 9223372036854775807 {
        case -790:
            for fi9 := 0; fi9 < 11; fi9++ {
                switch 32767 {
                case 1:
                    var v10 float64 = (1.5 * 3.14159)
                    _ = v10
                case 127:
                    var v11 int16 = -1
                    _ = v11
                case 128:
                    switch 1 {
                    case 9223372036854775807:
                        v12 := GenS1{F0: -9223372036854775808}
                        _ = v12
                    case 1:
                        v5 = (v5 / (((v5 * (^(v5)))) | 1))
                    }
                }
                v2 = v2
                v13 := GenS1{F0: -1}
                _ = fi9
                _ = v13
            }
        case -375:
            v1 = ("foo bar" + v3)
        }
    }
    v2 = -83
    v14 := v4["a"]
    v15 := GenS1{F0: 42}
    v16 := GenS2{F0: 1, F1: 0}
    v17 := GenS1{F0: 0}
    var v18 uint32 = 254
    fmt.Println(v3)
    var v19 int16 = 32767
    var v20 uint32 = 0
    v21 := []int{-926, -348, -9223372036854775808, 42}
    _ = len(v21)
    v22 := GenS1{F0: -433}
    fmt.Println(v3, v18, v14)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
