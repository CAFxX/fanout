package main

import (
    "fmt"
)

type GenS1 struct { F0 string }

func main() {
    var v1 int64 = -1
    var v2 uint64 = 0
    v3 := make(map[uint32]int)
    var v4 float64 = float64(0.0)
    v5 := []string{"hello"}
    v2 = v2
    v6 := GenS1{F0: ""}
    v7 := GenS1{F0: ("foo bar" + "")}
    fmt.Println(v4, v4)
    for fi8 := 0; fi8 < 11; fi8++ {
        var v9 uint16 = 255
        _ = fi8
        _ = v9
    }
    _ = len(v5)
    var v10 int32 = 2147483647
    for fi11 := 0; fi11 < 12; fi11++ {
        var v12 int = -2147483648
        _ = fi11
        _ = v12
    }
    var v13 uint = 2
    v13 += v13
    fmt.Println(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
