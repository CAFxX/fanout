package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint; F1 string }

func main() {
    v1 := make(map[uint32]bool)
    var v2 int64 = 0
    var v3 uint16 = 144
    var v4 int = -569
    var v5 int64 = 9223372036854775807
    v6 := make(map[int]string)
    if (v3 >= v3) {
        var v7 uint64 = 2
        v8 := GenS1{F0: (^(v7))}
        v7 += v7
        _ = v7
        _ = v8
    }
    v9 := GenS1{F0: 311}
    v3 -= (v3 << ((0) & 63))
    for fi10 := 0; fi10 < 11; fi10++ {
        fmt.Println(v2, v2)
        fmt.Println(v2, fi10)
        _ = fi10
    }
    v4 -= 0
    v11 := GenS1{F0: 895}
    v12 := GenS1{F0: 1}
    if false {
        for fi13 := 0; fi13 < 3; fi13++ {
            switch "foo bar" {
            case "foo bar":
                v14 := v1[0]
                _ = v14
            case "hello":
                v15 := []int{v4, v4, 9223372036854775807, fi13}
                _ = v15
            default:
                switch fi13 {
                case 42:
                    v16 := GenS2{F0: 255, F1: ("hello" + "hello")}
                    _ = v16
                default:
                    var v17 int = fi13
                    _ = v17
                }
            }
            if (!false) {
                var v18 float64 = 3.14159
                v4 -= (v4 - -844)
                _ = v18
            }
            _ = fi13
        }
    }
    v3 = v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
