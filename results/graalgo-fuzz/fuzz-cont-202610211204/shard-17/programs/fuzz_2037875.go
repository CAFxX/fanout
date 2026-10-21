package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32; F2 uint16 }

func main() {
    var v1 int64 = 142
    v2 := make(map[uint32]uint32)
    var v3 int8 = -126
    var v4 int64 = v1
    for fi5 := 0; fi5 < 12; fi5++ {
        var v6 bool = (((1.5 / -2.25) * 3.14159) == (1e10 * (1e10 / 1.0)))
        v6 = (32767 != 32767)
        v4 = v1
        _ = fi5
        _ = v6
    }
    v1 = (^(v4))
    v4 += v4
    fmt.Println(v4)
    v1 = v4
    fmt.Println(v1)
    if (1 > 42) {
        v3 -= 42
        if ((("\n\t" + "") + "\n\t") <= "foo bar") {
            var v7 uint8 = 44
            v4 = v1
            v8 := GenS1{F0: 698, F1: 4294967295, F2: 1}
            _ = v7
            _ = v8
        } else {
            v9 := []int{-202, 42, 2147483647}
            _ = v9
        }
    }
    delete(v2, 42)
    if (true || (!false)) {
        fmt.Println(v1)
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}
