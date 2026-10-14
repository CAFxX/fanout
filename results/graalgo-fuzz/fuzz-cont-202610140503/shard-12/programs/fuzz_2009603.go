package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint; F2 bool }

func main() {
    var v1 uint8 = 2
    v2 := make(map[string]int)
    v3 := make(map[string]int)
    var v4 int16 = 42
    v5 := []string{"hello"}
    v1 = (v1 * v1)
    fmt.Println(v4, v4)
    v2["a"] = 1
    var v6 int64 = 0
    v7 := GenS1{F0: 380, F1: 107, F2: (!((!(!true)) && ("a" == ("\n\t" + "\n\t"))))}
    for fi8 := 0; fi8 < 4; fi8++ {
        v4 += 787
        _ = fi8
    }
    var v9 float64 = 1e10
    switch "hello" {
    case "\n\t":
        var v10 uint8 = v1
        _ = v10
    case "a":
        v9 -= ((v9 + ((float64(3.14159) - ((3.14159 - (v9 / (1e10 / float64((1e10 / 1e10))))) / (0.0 / v9))) + -2.25)) / 1.0)
    }
    var v11 uint = uint(v1)
    v4 = (v4 << ((v11) & 63))
    switch -1 {
    case 127:
        fmt.Println(v11, v6)
    case -9223372036854775808:
        for fi12 := 0; fi12 < 13; fi12++ {
            v6 = (v6 & 1)
            _ = fi12
        }
    default:
        var v13 int32 = -532
        _ = v13
    }
    if (629 != 0) {
        v4 = (v4 * -32768)
    }
    if (!(v1 > v1)) {
        v1 = (^(v1))
    }
    var v14 float64 = 1.5
    fmt.Println(v1)
    fmt.Println(v14)
    v15 := GenS1{F0: 682, F1: 0, F2: false}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
