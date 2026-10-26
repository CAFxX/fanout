package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16; F2 int16 }
type GenS2 struct { F0 int8 }

func main() {
    var v1 int = -2147483648
    var v2 int32 = 2147483647
    var v3 uint16 = 2
    var v4 uint8 = 0
    if (3.14159 >= (1.5 + 3.14159)) {
        v5 := GenS1{F0: ("" + "foo bar"), F1: 0, F2: -32768}
        v6 := []int{v1}
        _ = v5
        _ = v6
    } else {
        var v7 int16 = 832
        for fi8 := 0; fi8 < 9; fi8++ {
            for fi9 := 0; fi9 < 10; fi9++ {
                v10 := GenS1{F0: "\n\t", F1: v3, F2: 475}
                _ = fi9
                _ = v10
            }
            v3 += v3
            _ = fi8
        }
        var v11 uint32 = 199
        _ = v11
        _ = v7
    }
    v4 -= 1
    var v12 int16 = 32767
    for fi13 := 0; fi13 < 14; fi13++ {
        var v14 bool = true
        v4 = (v4 >> ((2) & 63))
        switch "" {
        case "":
            v15 := []int{280, fi13, fi13, v1}
            _ = v15
        case "foo bar":
            if v14 {
                v16 := GenS2{F0: 121}
                v1 += (v1 % (((fi13 % (((fi13 ^ (fi13 - (v1 ^ (fi13 | 0))))) | 1))) | 1))
                _ = v16
            }
        case "\n\t":
            v2 = -470
        default:
            switch "foo bar" {
            case "foo bar":
                v3 += (^(v3))
            }
        }
        _ = fi13
        _ = v14
    }
    var v17 int8 = 42
    var v18 uint32 = 839
    var v19 int64 = -407
    fmt.Println(v12)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
