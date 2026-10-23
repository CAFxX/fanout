package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string; F2 string }
type GenS2 struct { F0 int64 }

func main() {
    var v1 uint64 = 2
    var v2 int8 = 42
    var v3 float32 = ((0.0 - 3.14159) - (0.1 - (3.14159 - 3.14159)))
    for fi4 := 0; fi4 < 15; fi4++ {
        v5 := []int{fi4, (fi4 & fi4)}
        _ = fi4
        _ = v5
    }
    v6 := GenS1{F0: 128, F1: ("hello" + "hello"), F2: ""}
    v1 -= (v1 << ((202) & 63))
    v2 = v2
    var v7 float64 = float64(v3)
    var v8 int16 = 0
    v9 := GenS1{F0: 146, F1: "a", F2: "foo bar"}
    for fi10 := 0; fi10 < 14; fi10++ {
        fmt.Println(v2, v7)
        if (80 <= (v2 * (v2 - v2))) {
            v1 += 1
        }
        if (true || ((v2 / ((93) | 1)) < (v2 | (v2 | v2)))) {
            if ((v2 & v2) > (v2 | 127)) {
                v11 := GenS2{F0: 0}
                v7 = 0.0
                v1 -= 2
                _ = v11
            }
            var v12 float64 = ((1.5 * float64(v3)) * 3.14159)
            _ = v12
        }
        _ = fi10
    }
    v3 = ((v3 * (v3 / 1.0)) + v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
