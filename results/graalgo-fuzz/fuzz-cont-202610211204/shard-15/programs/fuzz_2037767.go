package main

import (
    "fmt"
)

type GenS1 struct { F0 string }

func main() {
    var v1 int64 = 126
    v2 := make(map[uint32]bool)
    var v3 int = -531
    v4 := []string{(("" + ("foo bar" + "hello")) + ("foo bar" + "hello"))}
    var v5 string = "foo bar"
    for fi6 := 0; fi6 < 14; fi6++ {
        v1 += (-(v1))
        _ = fi6
    }
    for fi7 := 0; fi7 < 6; fi7++ {
        var v8 float32 = (3.14159 * 0.1)
        var v9 bool = true
        v10 := GenS1{F0: v5}
        _ = fi7
        _ = v10
        _ = v8
        _ = v9
    }
    switch "a" {
    case "hello":
        if (114 >= 1) {
            v3 += (v3 * 32767)
            var v11 float64 = 1.5
            v3 = 32768
            _ = v11
        }
    case "":
        fmt.Println(v3, v5)
    default:
        for fi12 := 0; fi12 < 1; fi12++ {
            var v13 bool = (0 == 614)
            _ = len(v4)
            var v14 uint8 = 0
            _ = fi12
            _ = v13
            _ = v14
        }
    }
    if (v3 != (v3 / ((-575) | 1))) {
        v5 = "foo bar"
        var v15 int8 = 100
        _ = v15
    }
    var v16 int64 = 9223372036854775807
    v3 -= (v3 << ((42) & 63))
    switch "" {
    case "a":
        switch (v5 + v5) {
        case "":
            v1 = (v1 / ((v1) | 1))
        }
    case "foo bar":
        _ = len(v4)
    default:
        fmt.Println(v16, v5, v16)
    }
    fmt.Println(v1, v5, v1)
    v17 := v2[1]
    fmt.Println(v1, v16, v1)
    fmt.Println(v1, v16, v1)
    for fi18 := 0; fi18 < 9; fi18++ {
        for fi19 := 0; fi19 < 15; fi19++ {
            var v20 uint8 = 2
            fmt.Println(fi19, v16, v16)
            fmt.Println(v17, v5, v1)
            _ = fi19
            _ = v20
        }
        var v21 string = "a"
        _ = fi18
        _ = v21
    }
    var v22 uint64 = 18446744073709551615
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
