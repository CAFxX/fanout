package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint8; F1 int }

func main() {
    v1 := make(map[string]int)
    var v2 uint32 = 0
    var v3 bool = true
    var v4 uint16 = 2
    v5 := []int{32768, 9223372036854775807}
    var v6 int16 = 0
    v5 = append(v5, 4294967295)
    for fi7 := 0; fi7 < 3; fi7++ {
        var v8 uint64 = 142
        _ = fi7
        _ = v8
    }
    v4 -= (v4 + v4)
    v2 = (v2 >> ((65535) & 63))
    var v9 uint64 = 888
    v9 = v9
    if (!(-2147483648 < -823)) {
        for fi10 := 0; fi10 < 16; fi10++ {
            for fi11 := 0; fi11 < 16; fi11++ {
                switch (fi11 / ((fi10) | 1)) {
                case 407:
                    v12 := v1["a"]
                    _ = v12
                case 42:
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    var v13 int8 = 100
                    _ = v13
                default:
                    fmt.Println(v9, fi10, v3)
                }
                _ = fi11
            }
            v6 -= (-(v6))
            v6 = -32768
            _ = fi10
        }
        var v14 int8 = 127
        if len(v5) > 0 {
            v5[0] = int(v9)
        }
        _ = v14
    }
    if (v3 || ((((!v3) || v3) && (v3 || false)) && (!((1 > 42) || v3)))) {
        delete(v1, "hello")
        v6 += v6
        fmt.Println(v4, v3, v4)
    } else {
        v15 := GenS1{F0: (v4 % ((v4) | 1))}
        switch "a" {
        case "a":
            fmt.Println(v9, v6, v6)
        case "\n\t":
            var v16 uint64 = v9
            _ = v16
        case "foo bar":
            var v17 uint64 = v9
            _ = v17
        default:
            v18 := GenS1{F0: 428}
            _ = v18
        }
        _ = v15
    }
    v9 += (^(v9))
    v19 := GenS2{F0: 183, F1: 127}
    var v20 uint16 = (v4 + (^(v4)))
    var v21 int64 = 9223372036854775807
    v2 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
