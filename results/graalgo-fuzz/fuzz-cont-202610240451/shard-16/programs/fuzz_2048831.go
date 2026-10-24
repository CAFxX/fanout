package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8 }
type GenS2 struct { F0 uint32; F1 int64 }

func main() {
    var v1 uint16 = 32767
    v2 := make(map[int]uint32)
    v3 := make(map[int]uint32)
    var v4 bool = ((-9223372036854775808 < 42) || false)
    var v5 int = 128
    fmt.Println(v4, v1)
    v6 := v2[v5]
    switch "a" {
    case "\n\t":
        v7 := GenS1{F0: 2147483647, F1: -62}
        _ = v7
    case "a":
        v8 := GenS1{F0: 42, F1: -7}
        _ = v8
    case "hello":
        if (v4 && true) {
            switch "a" {
            case "a":
                var v9 int64 = -1
                _ = v9
            case "hello":
                v10 := GenS1{F0: -1, F1: 57}
                _ = v10
            case "\n\t":
                var v11 float32 = 0.0
                _ = v11
            default:
                var v12 int = v5
                _ = v12
            }
            switch "a" {
            case "a":
                for fi13 := 0; fi13 < 10; fi13++ {
                    v14 := []int{v5}
                    _ = fi13
                    _ = v14
                }
            case "foo bar":
                for fi15 := 0; fi15 < 9; fi15++ {
                    v1 = 42
                    _ = fi15
                }
            default:
                if (!v4) {
                    fmt.Println(v5, v1)
                } else {
                    var v16 bool = (79 > 42)
                    v6 = (v6 % (((v6 - 4294967295)) | 1))
                    _ = v16
                }
            }
        } else {
            fmt.Println(v5, v4, v1)
        }
    default:
        v6 = (v6 / (((v6 - v6)) | 1))
    }
    var v17 uint8 = 127
    v18 := GenS1{F0: -2147483648, F1: 100}
    v19 := GenS2{F0: v6, F1: 9223372036854775807}
    for fi20 := 0; fi20 < 3; fi20++ {
        var v21 float64 = ((float64(0.1) + (1e10 * -2.25)) * 1.5)
        for fi22 := 0; fi22 < 16; fi22++ {
            var v23 uint32 = (v6 + 1)
            _ = fi22
            _ = v23
        }
        _ = fi20
        _ = v21
    }
    if v4 {
        v24 := []int{(v5 * v5)}
        v4 = v4
        _ = v24
    } else {
        v17 -= (v17 | (v17 - 2))
    }
    if (v1 < 1) {
        v17 += (v17 + v17)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
