package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64 }

func main() {
    var v1 float64 = 0.0
    var v2 uint16 = 356
    var v3 uint8 = 0
    v4 := []string{("a" + "")}
    v5 := []string{}
    v1 -= v1
    v6 := GenS1{F0: (true && false), F1: 18446744073709551615}
    fmt.Println(v2)
    _ = len(v5)
    var v7 int8 = -27
    var v8 int16 = 801
    var v9 uint64 = 754
    for fi10 := 0; fi10 < 17; fi10++ {
        var v11 uint16 = v2
        for fi12 := 0; fi12 < 9; fi12++ {
            v11 = (v11 << ((83) & 63))
            if ("foo bar" < "") {
                v13 := GenS1{F0: (false || false), F1: (v9 - (^(v9)))}
                v14 := GenS1{F0: true, F1: v9}
                _ = v13
                _ = v14
            } else {
                var v15 int16 = (-(v8))
                v16 := GenS1{F0: true, F1: (v9 ^ (v9 * (^(v9))))}
                _ = v15
                _ = v16
            }
            _ = fi12
        }
        _ = fi10
        _ = v11
    }
    fmt.Println(v7, v8)
    var v17 float64 = float64(-2.25)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
