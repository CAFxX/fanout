package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }

func main() {
    var v1 int16 = 42
    var v2 int16 = 1000
    v3 := make(map[uint32]string)
    v2 += 42
    if false {
        var v4 uint64 = 4294967295
        var v5 float64 = 0.1
        _ = v4
        _ = v5
    }
    v6 := GenS1{F0: ((127 > 1) && ((0.1 - 1e10) < (1.5 - (float64(0.1) / (1.5 + 1e10)))))}
    switch 32767 {
    case -2147483648:
        var v7 uint8 = 127
        _ = v7
    default:
        var v8 bool = true
        _ = v8
    }
    v9 := GenS1{F0: false}
    v10 := []int{1, 168, 0, -1}
    v2 += (^(v1))
    for fi11 := 0; fi11 < 16; fi11++ {
        if false {
            var v12 int32 = 42
            switch fi11 {
            case 533:
                switch "foo bar" {
                case "a":
                    if len(v10) > 0 {
                        v10[0] = fi11
                    }
                default:
                    var v13 uint = 65535
                    _ = v13
                }
            case -9223372036854775808:
                if len(v10) > 0 {
                    v10[0] = (fi11 - fi11)
                }
            }
            var v14 int32 = (v12 - 1)
            _ = v12
            _ = v14
        }
        var v15 int = 938
        v16 := v3[1]
        _ = fi11
        _ = v15
        _ = v16
    }
    v1 = (v2 + (v1 + 32767))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v6
    _ = v9
}
