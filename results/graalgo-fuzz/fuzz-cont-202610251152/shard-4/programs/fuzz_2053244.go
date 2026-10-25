package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64; F2 int16 }

func main() {
    v1 := []string{"\n\t", "hello", ""}
    v2 := []string{"\n\t", ("" + "")}
    v3 := make(map[uint32]int)
    var v4 uint64 = 154
    var v5 int16 = -850
    var v6 int = 0
    v5 -= (v5 - (v5 | (v5 / ((v5) | 1))))
    for fi7 := 0; fi7 < 2; fi7++ {
        v3[42] = v6
        _ = fi7
    }
    var v8 int32 = -2147483648
    if true {
        var v9 uint64 = v4
        v8 -= -40
        v5 += v5
        _ = v9
    }
    var v10 uint64 = (v4 % ((v4) | 1))
    if true {
        v6 += (-(v6))
    }
    var v11 bool = true
    var v12 int16 = (v5 % ((v5) | 1))
    switch "\n\t" {
    case "hello":
        v5 += (^(v12))
    case "":
        if (2 == 127) {
            v12 -= v5
        }
    default:
        switch ("hello" + ("foo bar" + "\n\t")) {
        case "":
            v5 -= v5
        case "\n\t":
            v13 := GenS1{F0: v8, F1: 0, F2: v5}
            _ = v13
        }
    }
    fmt.Println(v12, v8, v12)
    for fi14 := 0; fi14 < 7; fi14++ {
        for fi15 := 0; fi15 < 12; fi15++ {
            fmt.Println(v5, v6, v5)
            _ = fi15
        }
        _ = fi14
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
