package main

import (
    "fmt"
)


func main() {
    v1 := []string{"", "foo bar"}
    var v2 uint16 = 42
    var v3 float32 = 3.14159
    var v4 bool = true
    var v5 float64 = -2.25
    v2 += v2
    var v6 int = 128
    var v7 uint16 = 833
    for fi8 := 0; fi8 < 2; fi8++ {
        v6 = (v6 & (fi8 ^ v6))
        switch fi8 {
        case 172:
            v6 = (v6 << ((42) & 63))
        default:
            var v9 float32 = ((v3 - v3) - v3)
            _ = v9
        }
        _ = fi8
    }
    fmt.Println(v4)
    v6 += -1
    var v10 uint8 = 42
    if len(v1) > 0 {
        v1[0] = "a"
    }
    v4 = (65535 > v7)
    if len(v1) > 0 {
        v1[0] = "foo bar"
    }
    fmt.Println(v10)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
