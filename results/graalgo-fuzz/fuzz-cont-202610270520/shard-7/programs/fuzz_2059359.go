package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8 }

func main() {
    v1 := []string{}
    v2 := make(map[string]string)
    var v3 int8 = -1
    var v4 int64 = -62
    var v5 uint64 = 42
    fmt.Println(v5, v5, v5)
    var v6 uint64 = (v5 - v5)
    if ((!("" > "hello")) && true) {
        var v7 int32 = -705
        _ = v7
    }
    v3 += -1
    for fi8 := 0; fi8 < 10; fi8++ {
        if ((1e10 - (0.0 + ((0.0 + 0.1) - ((3.14159 / -2.25) / 1.5)))) < ((float64(((((0.0 - -2.25) + 0.0) * 3.14159) * 1e10)) + (3.14159 + 3.14159)) * -2.25)) {
            v9 := GenS1{F0: 1, F1: v3}
            _ = v9
        }
        for fi10 := 0; fi10 < 11; fi10++ {
            v4 = (v4 >> ((622) & 63))
            var v11 string = (("" + "foo bar") + "foo bar")
            _ = fi10
            _ = v11
        }
        _ = fi8
    }
    var v12 uint32 = 38
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
