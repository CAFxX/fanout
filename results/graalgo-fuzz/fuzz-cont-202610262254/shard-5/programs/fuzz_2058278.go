package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8 }
func genFn0(p0 int64, p1 uint16) bool {
    v7 := GenS1{F0: 0, F1: 166}
    var v8 uint = 646
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return (!(805 < 698))
}

func main() {
    v1 := make(map[string]string)
    var v2 int64 = -1
    var v3 uint64 = 18446744073709551615
    var v4 uint8 = 122
    var v5 int64 = (v2 & (v2 & v2))
    var v6 float64 = 1e10
    switch "\n\t" {
    case "":
        v9 := GenS1{F0: 1, F1: (v4 * v4)}
        _ = v9
    case "hello":
        fmt.Println(v2)
    case "\n\t":
        var v10 int = 884
        _ = v10
    }
    var v11 bool = (65535 == 65535)
    var v12 uint = 65535
    v13 := GenS1{F0: (v12 >> (((v12 & v12)) & 63)), F1: 1}
    v6 -= -2.25
    for fi14 := 0; fi14 < 20; fi14++ {
        if v11 {
            v15 := GenS1{F0: v12, F1: (v4 % (((v4 | (v4 & 127))) | 1))}
            fmt.Println(v6, v5)
            _ = v15
        } else {
            v16 := []int{fi14}
            var v17 uint8 = 1
            v5 += (^(v5))
            _ = v16
            _ = v17
        }
        _ = fi14
    }
    v18 := []int{}
    v1["a"] = "\n\t"
    for fi19 := 0; fi19 < 9; fi19++ {
        if (!(!true)) {
            var v20 int = fi19
            _ = v20
        }
        _ = fi19
    }
    v21 := GenS1{F0: v12, F1: v4}
    v4 += (v4 ^ v4)
    v12 = (v12 / ((v12) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v18
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
