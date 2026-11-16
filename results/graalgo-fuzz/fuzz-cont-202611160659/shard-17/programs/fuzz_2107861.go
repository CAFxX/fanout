package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8 }

func main() {
    var v1 bool = (false || (246 > 468))
    var v2 uint32 = 65535
    v3 := []string{"hello", "a", "hello", "foo bar"}
    var v4 string = "\n\t"
    v5 := make(map[string]bool)
    var v6 float32 = 1e10
    if (!true) {
        var v7 uint8 = 2
        var v8 uint8 = (v7 << ((42) & 63))
        _ = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 19; fi9++ {
        v2 = 918
        v4 = ((v4 + (v4 + (v4 + "\n\t"))) + v4)
        _ = fi9
    }
    v6 = (v6 / ((1e10 / v6) * 1e10))
    v10 := GenS1{F0: 32767, F1: 42}
    var v11 int16 = -1
    v2 = v2
    var v12 uint16 = 2
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
