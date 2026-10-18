package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int; F2 uint16 }
type GenS2 struct { F0 int16 }

func main() {
    var v1 uint64 = 0
    v2 := make(map[uint32]int)
    var v3 uint8 = 122
    var v4 uint16 = 561
    v5 := []int{9223372036854775807, 510}
    var v6 int8 = 127
    v7 := GenS1{F0: -436, F1: -9223372036854775808, F2: (^(v4))}
    v8 := GenS1{F0: 2147483647, F1: int(v6), F2: 2}
    fmt.Println(v4, v3)
    var v9 uint16 = (v4 - (v4 >> ((42) & 63)))
    fmt.Println(v6, v9, v1)
    var v10 uint64 = (^(v1))
    for fi11 := 0; fi11 < 19; fi11++ {
        switch (fi11 + (^(fi11))) {
        case 9223372036854775807:
            switch "\n\t" {
            case "":
                switch 815 {
                case -2147483648:
                    fmt.Println(v1, v3)
                default:
                    if ((!true) || true) {
                        fmt.Println(v4)
                    }
                }
            case "foo bar":
                v3 -= v3
            default:
                for fi12 := 0; fi12 < 16; fi12++ {
                    var v13 float32 = 1e10
                    _ = fi12
                    _ = v13
                }
            }
        case 958:
            var v14 int8 = (v6 >> ((2) & 63))
            _ = v14
        case 42:
            v4 += v4
        default:
            var v15 float32 = (3.14159 / 3.14159)
            _ = v15
        }
        v6 += (v6 << ((0) & 63))
        for fi16 := 0; fi16 < 6; fi16++ {
            delete(v2, 2147483648)
            fmt.Println(fi11, v9, v4)
            v17 := GenS2{F0: 51}
            fmt.Println(v4)
            _ = fi16
            _ = v17
        }
        _ = fi11
    }
    var v18 uint16 = uint16(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
