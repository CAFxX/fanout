package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }

func main() {
    var v1 int32 = 717
    v2 := make(map[uint32]bool)
    var v3 string = "a"
    var v4 string = ("foo bar" + "\n\t")
    var v5 int64 = 9223372036854775807
    var v6 uint16 = 175
    fmt.Println(v3, v3, v4)
    v7 := GenS1{F0: 18446744073709551615}
    v6 -= (^(v6))
    if false {
        for fi8 := 0; fi8 < 5; fi8++ {
            v6 = v6
            switch fi8 {
            case -779:
                for fi9 := 0; fi9 < 13; fi9++ {
                    var v10 int8 = -112
                    _ = fi9
                    _ = v10
                }
            default:
                var v11 uint32 = 42
                _ = v11
            }
            fmt.Println(v6, v4)
            _ = fi8
        }
    }
    var v12 uint8 = 220
    for fi13 := 0; fi13 < 18; fi13++ {
        for fi14 := 0; fi14 < 19; fi14++ {
            switch "\n\t" {
            case "":
                fmt.Println(fi13, v1, v3)
            case "\n\t":
                var v15 int = fi14
                _ = v15
            }
            v12 = (v12 / ((163) | 1))
            _ = fi14
        }
        _ = fi13
    }
    var v16 float32 = float32(1e10)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
