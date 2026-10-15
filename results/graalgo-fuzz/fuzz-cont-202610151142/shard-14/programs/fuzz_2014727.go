package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }

func main() {
    var v1 int = -76
    var v2 int16 = 828
    v3 := make(map[int]int)
    v4 := GenS1{F0: 100}
    v2 += v2
    v5 := v3[(v1 | v1)]
    if (2147483648 == 42) {
        var v6 uint32 = 2147483648
        var v7 uint32 = 2147483648
        _ = v6
        _ = v7
    }
    for fi8 := 0; fi8 < 16; fi8++ {
        v9 := v3[0]
        _ = fi8
        _ = v9
    }
    var v10 string = ("" + "foo bar")
    v2 = (v2 >> ((2) & 63))
    fmt.Println(v2, v2, v5)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
