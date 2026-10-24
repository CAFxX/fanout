package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }

func main() {
    v1 := make(map[uint32]string)
    v2 := []int{-894, 32768, -9, 127}
    var v3 int8 = 0
    var v4 int8 = 0
    var v5 int64 = 1
    switch ("a" + ("foo bar" + "foo bar")) {
    case "a":
        v3 = (v4 & (v3 ^ v4))
    case "\n\t":
        var v6 int16 = -1
        _ = v6
    }
    v4 = v4
    for fi7 := 0; fi7 < 6; fi7++ {
        fmt.Println(v4, fi7, v3)
        v8 := GenS1{F0: 141}
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 14; fi9++ {
        var v10 int = (fi9 ^ (fi9 * (fi9 % ((fi9) | 1))))
        _ = fi9
        _ = v10
    }
    var v11 uint16 = 484
    v5 += v5
    if (0 != 1) {
        v12 := GenS1{F0: (^(v11))}
        var v13 int8 = (v4 & 0)
        _ = v12
        _ = v13
    }
    for fi14 := 0; fi14 < 14; fi14++ {
        v15 := GenS1{F0: (v11 * (v11 * v11))}
        _ = fi14
        _ = v15
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
