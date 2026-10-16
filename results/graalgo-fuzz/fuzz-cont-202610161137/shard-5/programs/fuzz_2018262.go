package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64 }

func main() {
    v1 := []string{}
    var v2 int32 = 274
    var v3 uint64 = 42
    v4 := make(map[uint32]int)
    var v5 uint8 = 128
    v6 := v4[uint32(v3)]
    for fi7 := 0; fi7 < 19; fi7++ {
        for fi8 := 0; fi8 < 7; fi8++ {
            var v9 bool = ((v2 >> ((18446744073709551615) & 63)) > v2)
            v1 = append(v1, "a")
            _ = fi8
            _ = v9
        }
        v2 -= (v2 * (^(v2)))
        v10 := GenS1{F0: true, F1: v3}
        _ = fi7
        _ = v10
    }
    if false {
        for fi11 := 0; fi11 < 9; fi11++ {
            switch "\n\t" {
            case "":
                v5 -= (v5 | 33)
            case "\n\t":
                v12 := GenS1{F0: (-1 < (fi11 + 42)), F1: (v3 >> ((863) & 63))}
                _ = v12
            default:
                var v13 uint16 = 0
                _ = v13
            }
            _ = fi11
        }
        fmt.Println(v6, v5)
    } else {
        delete(v4, 42)
        v14 := GenS1{F0: false, F1: v3}
        _ = v14
    }
    v15 := GenS1{F0: true, F1: (v3 + 4294967295)}
    var v16 float64 = ((1e10 - 1.5) / 0.1)
    for fi17 := 0; fi17 < 12; fi17++ {
        v18 := GenS1{F0: false, F1: (v3 + (v3 - (v3 >> ((255) & 63))))}
        v3 = 653
        _ = fi17
        _ = v18
    }
    if false {
        var v19 string = ("a" + "hello")
        switch "foo bar" {
        case "\n\t":
            v16 = (0.1 * (v16 - v16))
        case "a":
            var v20 uint = 4294967295
            _ = v20
        default:
            var v21 uint16 = 1
            _ = v21
        }
        _ = v19
    }
    v2 = (-(v2))
    v5 = v5
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
