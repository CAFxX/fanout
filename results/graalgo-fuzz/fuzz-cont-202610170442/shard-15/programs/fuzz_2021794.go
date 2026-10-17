package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8; F2 uint64 }

func main() {
    v1 := []string{("" + "a"), "foo bar"}
    var v2 uint32 = 42
    v3 := make(map[uint32]bool)
    var v4 float32 = (float32(float64(-2.25)) + 1.5)
    fmt.Println(v2, v4)
    v5 := GenS1{F0: true, F1: 255, F2: 915}
    var v6 string = (("" + "\n\t") + ("a" + ("\n\t" + "a")))
    if true {
        v4 = v4
        v2 += (v2 / ((225) | 1))
    }
    for fi7 := 0; fi7 < 12; fi7++ {
        v8 := GenS1{F0: (false && true), F1: 255, F2: 42}
        _ = fi7
        _ = v8
    }
    v9 := GenS1{F0: false, F1: 41, F2: 110}
    v4 = (1e10 / (v4 - v4))
    for fi10 := 0; fi10 < 7; fi10++ {
        fmt.Println(v2, v4)
        var v11 int32 = -2147483648
        _ = fi10
        _ = v11
    }
    v12 := GenS1{F0: false, F1: 255, F2: 4294967295}
    var v13 int16 = 207
    switch v6 {
    case "foo bar":
        v14 := v3[v2]
        _ = v14
    case "\n\t":
        v13 = (v13 << ((789) & 63))
    }
    fmt.Println(v2, v2, v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
