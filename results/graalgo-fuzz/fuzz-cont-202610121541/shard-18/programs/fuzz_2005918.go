package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int; F2 uint64 }
type GenS2 struct { F0 uint64 }

func main() {
    var v1 float64 = 0.0
    var v2 uint64 = 18446744073709551615
    var v3 float32 = 0.1
    v4 := make(map[uint32]string)
    var v5 uint8 = 51
    switch 0 {
    case 32767:
        v5 += (^(v5))
    case 0:
        switch 927 {
        case 718:
            for fi6 := 0; fi6 < 14; fi6++ {
                var v7 uint = 301
                for fi8 := 0; fi8 < 12; fi8++ {
                    var v9 int8 = -1
                    _ = fi8
                    _ = v9
                }
                _ = fi6
                _ = v7
            }
        case -996:
            v3 += (float32(v1) * v3)
        case 996:
            var v10 uint8 = v5
            _ = v10
        }
    default:
        var v11 float32 = float32((((v1 * ((float64(float32((v1 / 0.1))) * (v1 * v1)) / -2.25)) - 0.0) - ((v1 - 3.14159) + v1)))
        _ = v11
    }
    v4[2] = "\n\t"
    fmt.Println(v1)
    var v12 bool = false
    for fi13 := 0; fi13 < 1; fi13++ {
        var v14 uint64 = 503
        var v15 uint32 = 2147483648
        switch 905 {
        case 9223372036854775807:
            fmt.Println(v5, v15, v5)
        default:
            switch "a" {
            case "":
                if (808 == 42) {
                    var v16 float64 = 0.0
                    fmt.Println(fi13, v2, v15)
                    _ = v16
                }
            }
        }
        _ = fi13
        _ = v14
        _ = v15
    }
    fmt.Println(v1, v2)
    var v17 uint8 = (^(v5))
    v5 += 0
    v12 = (v3 >= v3)
    v18 := GenS1{F0: (v5 % ((76) | 1)), F1: 128, F2: (v2 % ((v2) | 1))}
    var v19 int32 = -170
    v20 := GenS1{F0: 165, F1: 32768, F2: (^(v2))}
    v21 := GenS2{F0: (v2 / (((v2 & (v2 * (v2 + v2)))) | 1))}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
