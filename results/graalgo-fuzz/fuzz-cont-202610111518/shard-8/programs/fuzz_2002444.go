package main

import (
    "fmt"
)

type GenS1 struct { F0 int }

func main() {
    v1 := make(map[string]string)
    var v2 uint64 = 4294967295
    var v3 int = 1
    var v4 uint8 = 127
    var v5 int16 = 32767
    for fi6 := 0; fi6 < 2; fi6++ {
        fmt.Println(v4, v5)
        v7 := GenS1{F0: v3}
        v2 -= v2
        _ = fi6
        _ = v7
    }
    v3 += v3
    v4 += (v4 >> ((73) & 63))
    fmt.Println(v2, v3, v5)
    for fi8 := 0; fi8 < 15; fi8++ {
        var v9 int16 = 1
        v5 += (^(v9))
        v10 := v1[(("a" + "foo bar") + "foo bar")]
        _ = fi8
        _ = v10
        _ = v9
    }
    v1["hello"] = "a"
    _ = len(v1)
    v4 += (v4 + (^(v4)))
    v3 = (v3 ^ v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
