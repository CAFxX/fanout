package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16; F2 bool }

func main() {
    var v1 uint32 = 2147483648
    var v2 int64 = 42
    v3 := []string{"hello", "a", "a", ("" + "\n\t")}
    if ((!true) || true) {
        fmt.Println(v1)
    } else {
        switch "hello" {
        case "foo bar":
            var v4 int32 = -596
            _ = v4
        }
    }
    var v5 uint8 = 219
    fmt.Println(v2, v1)
    fmt.Println(v2, v5)
    for fi6 := 0; fi6 < 9; fi6++ {
        if (65535 > 2) {
            var v7 int64 = v2
            v5 = (v5 << ((2) & 63))
            _ = v7
        }
        for fi8 := 0; fi8 < 12; fi8++ {
            v2 = (v2 | (v2 ^ v2))
            _ = fi8
        }
        _ = fi6
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println(v5, v1, v1)
    var v9 int32 = -657
    for fi10 := 0; fi10 < 6; fi10++ {
        v5 -= v5
        _ = fi10
    }
    for fi11 := 0; fi11 < 4; fi11++ {
        if len(v3) > 0 {
            v3[0] = ("" + "hello")
        }
        _ = fi11
    }
    v12 := GenS1{F0: -853, F1: 586, F2: (42 < 748)}
    if false {
        _ = len(v3)
    } else {
        v2 += v2
    }
    var v13 uint = 18446744073709551615
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v5
    _ = v9
}
