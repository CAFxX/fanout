package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }

func main() {
    var v1 float64 = (3.14159 + (float64(1.5) + 0.1))
    var v2 uint64 = 0
    v3 := make(map[uint32]uint32)
    v4 := []string{"hello"}
    var v5 int64 = -1
    var v6 uint32 = 204
    if (906 >= -244) {
        v6 -= v6
        v5 += 9223372036854775807
        v2 -= (v2 << ((255) & 63))
    }
    var v7 bool = ((v5 << ((255) & 63)) != (v5 & (v5 + (v5 | v5))))
    var v8 string = ("" + ((("" + "foo bar") + "hello") + ("\n\t" + "foo bar")))
    v9 := GenS1{F0: v6}
    _ = len(v3)
    v4 = append(v4, v8)
    var v10 float32 = 0.0
    if false {
        var v11 uint64 = (v2 + v2)
        for fi12 := 0; fi12 < 13; fi12++ {
            v6 -= 42
            v13 := v3[(^(v6))]
            v2 = v2
            _ = fi12
            _ = v13
        }
        _ = v11
    }
    _ = len(v3)
    v14 := GenS1{F0: v6}
    switch "" {
    case "":
        v7 = (v2 > (^(v2)))
    case "\n\t":
        if false {
            fmt.Println(v10, v1)
        }
    }
    var v15 int64 = -711
    for fi16 := 0; fi16 < 16; fi16++ {
        var v17 int64 = -304
        _ = fi16
        _ = v17
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
