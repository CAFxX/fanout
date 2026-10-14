package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint; F2 int32 }
func genFn0(p0 int32) uint8 {
    fmt.Println(p0)
    p0 += -733
    _ = p0
    return 216
}

func main() {
    var v1 bool = (2147483648 != 2147483647)
    v2 := make(map[uint32]bool)
    var v3 int = 1
    for fi4 := 0; fi4 < 6; fi4++ {
        v1 = ((-1 < -77) && v1)
        _ = fi4
    }
    v2[1] = v1
    v1 = v1
    v1 = (128 < v3)
    var v5 int32 = 2147483647
    for fi6 := 0; fi6 < 18; fi6++ {
        v3 += v3
        v7 := GenS1{F0: 42, F1: 581, F2: (v5 - v5)}
        switch "" {
        case "\n\t":
            var v8 int16 = 1
            _ = v8
        case "foo bar":
            var v9 float64 = 1e10
            _ = v9
        }
        _ = fi6
        _ = v7
    }
    var v10 int16 = 0
    v11 := []int{v3}
    v12 := v2[97]
    v3 = v3
    for fi13 := 0; fi13 < 16; fi13++ {
        _ = len(v11)
        _ = fi13
    }
    v14 := v2[4294967295]
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v14
    _ = v2
    _ = v3
    _ = v5
}
