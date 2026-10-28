package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint16; F1 string }

func main() {
    var v1 int32 = -1
    var v2 int32 = v1
    var v3 uint8 = 0
    v4 := make(map[int]int)
    var v5 int16 = 243
    v6 := GenS1{F0: false}
    v5 -= v5
    fmt.Println(v3)
    fmt.Println(v5, v3)
    var v7 bool = (987 > 255)
    fmt.Println(v3, v5)
    var v8 int = 9223372036854775807
    for fi9 := 0; fi9 < 6; fi9++ {
        var v10 uint64 = 149
        _ = fi9
        _ = v10
    }
    for fi11 := 0; fi11 < 4; fi11++ {
        v12 := GenS2{F0: 1, F1: "hello"}
        v13 := []int{795, fi11}
        _ = fi11
        _ = v12
        _ = v13
    }
    v5 = (v5 % (((v5 | v5)) | 1))
    switch v8 {
    case 56:
        v3 -= (v3 ^ v3)
    default:
        switch "" {
        case "":
            for fi14 := 0; fi14 < 14; fi14++ {
                v15 := []int{v8}
                v2 = v2
                switch "hello" {
                case "a":
                    v16 := v4[v8]
                    _ = v16
                case "":
                    _ = len(v15)
                case "\n\t":
                    var v17 uint = 1
                    _ = v17
                }
                _ = fi14
                _ = v15
            }
        default:
            for fi18 := 0; fi18 < 13; fi18++ {
                var v19 uint16 = 1
                _ = fi18
                _ = v19
            }
        }
    }
    v8 = v8
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
