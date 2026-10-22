package main

import (
    "fmt"
)

type GenS1 struct { F0 string }

func main() {
    var v1 float64 = 1.5
    v2 := make(map[uint32]string)
    var v3 int64 = -1
    var v4 uint64 = 76
    var v5 uint64 = v4
    var v6 int = -871
    _ = len(v2)
    for fi7 := 0; fi7 < 12; fi7++ {
        for fi8 := 0; fi8 < 18; fi8++ {
            var v9 float64 = v1
            v10 := v2[1]
            fmt.Println(fi7, fi8, v5)
            _ = fi8
            _ = v10
            _ = v9
        }
        v11 := []int{(fi7 + 1), fi7, fi7, (fi7 << ((18446744073709551615) & 63))}
        var v12 uint64 = 18446744073709551615
        _ = fi7
        _ = v11
        _ = v12
    }
    var v13 float64 = 0.0
    v5 -= v4
    v14 := GenS1{F0: "foo bar"}
    var v15 uint32 = 4294967295
    fmt.Println(v3, v5, v6)
    switch v6 {
    case -9223372036854775808:
        fmt.Println(v13)
    case 1:
        var v16 uint = 766
        _ = v16
    case -155:
        var v17 bool = false
        _ = v17
    }
    for fi18 := 0; fi18 < 20; fi18++ {
        fmt.Println(v5)
        if (368 < 243) {
            v19 := GenS1{F0: ("foo bar" + ("foo bar" + "foo bar"))}
            _ = v19
        }
        v20 := GenS1{F0: ""}
        _ = fi18
        _ = v20
    }
    v2[(v15 % ((208) | 1))] = ("hello" + "\n\t")
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
