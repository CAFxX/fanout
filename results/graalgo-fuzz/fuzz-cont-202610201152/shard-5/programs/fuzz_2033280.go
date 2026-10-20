package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8; F2 uint64 }

func main() {
    v1 := make(map[uint32]int)
    var v2 uint8 = 127
    var v3 float32 = (-2.25 * 0.1)
    var v4 int16 = 1
    v2 = (v2 - (v2 | v2))
    v2 = (v2 + 1)
    v5 := []int{0}
    var v6 int16 = (^(v4))
    _ = len(v5)
    for fi7 := 0; fi7 < 17; fi7++ {
        switch int(v6) {
        case 32767:
            fmt.Println(v2, v6, v6)
        case 4294967295:
            v8 := GenS1{F0: ("foo bar" == ""), F1: (v2 ^ v2), F2: 18446744073709551615}
            _ = v8
        default:
            fmt.Println(v6)
        }
        v9 := GenS1{F0: (true || (!true)), F1: (v2 % (((v2 % ((1) | 1))) | 1)), F2: 1}
        _ = fi7
        _ = v9
    }
    if true {
        _ = len(v5)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
