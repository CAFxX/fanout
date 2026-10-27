package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16 }
func genFn0(p0 int32) int {
    var v7 float32 = float32(0.1)
    fmt.Println(p0, p0)
    _ = p0
    _ = v7
    return 9223372036854775807
}

func main() {
    var v1 int32 = 674
    v2 := make(map[string]bool)
    var v3 uint64 = 2
    var v4 uint32 = 211
    v5 := make(map[int]uint32)
    v6 := make(map[uint32]int)
    v3 -= (v3 - v3)
    var v8 int = -9223372036854775808
    v8 = genFn0((v1 + 42))
    v8 -= v8
    var v9 uint = 258
    var v10 uint32 = (^(v4))
    v8 += v8
    v11 := GenS1{F0: -91, F1: 42}
    for fi12 := 0; fi12 < 18; fi12++ {
        fmt.Println(v1, v3)
        _ = fi12
    }
    switch "foo bar" {
    case "hello":
        for fi13 := 0; fi13 < 7; fi13++ {
            v1 -= 0
            _ = fi13
        }
    case "\n\t":
        switch (v8 + (^(v8))) {
        case 9223372036854775807:
            for fi14 := 0; fi14 < 16; fi14++ {
                v15 := GenS1{F0: 1, F1: 2}
                v3 += (v3 & (^(v3)))
                var v16 uint32 = v10
                _ = fi14
                _ = v15
                _ = v16
            }
        case -746:
            switch "\n\t" {
            case "\n\t":
                v3 = 0
            case "a":
                for fi17 := 0; fi17 < 9; fi17++ {
                    for fi18 := 0; fi18 < 10; fi18++ {
                        v19 := v5[v8]
                        var v20 int64 = 1
                        _ = fi18
                        _ = v19
                        _ = v20
                    }
                    v21 := []int{(v8 - (v8 * v8)), (v8 ^ genFn0(v1))}
                    _ = fi17
                    _ = v21
                }
            default:
                var v22 int = v8
                _ = v22
            }
        case 1:
            v5[genFn0((v1 & v1))] = (^(v4))
        default:
            v3 -= v3
        }
    case "foo bar":
        _ = len(v5)
    }
    var v23 bool = false
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
