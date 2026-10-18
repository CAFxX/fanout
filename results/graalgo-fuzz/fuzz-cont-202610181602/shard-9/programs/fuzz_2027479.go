package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8 }

func main() {
    var v1 uint = 255
    var v2 uint64 = 1
    var v3 int = 2147483647
    var v4 float64 = 1.5
    v3 = (v3 >> ((v1) & 63))
    var v5 int = (v3 - 32767)
    v3 -= (v3 * v5)
    v6 := []int{(v3 & (v5 & v3)), v3, v3}
    var v7 string = ("" + "hello")
    _ = len(v6)
    if false {
        switch (v7 + (("\n\t" + "\n\t") + "foo bar")) {
        case "\n\t":
            var v8 uint32 = 1
            _ = v8
        case "foo bar":
            v9 := GenS1{F0: 997, F1: 127}
            _ = v9
        default:
            v3 -= -9223372036854775808
        }
        fmt.Println(v4, v2)
    } else {
        if len(v6) > 0 {
            v6[0] = v3
        }
        for fi10 := 0; fi10 < 1; fi10++ {
            for fi11 := 0; fi11 < 14; fi11++ {
                v12 := GenS1{F0: 823, F1: 0}
                var v13 int8 = 42
                _ = fi11
                _ = v12
                _ = v13
            }
            fmt.Println(fi10)
            v1 = v1
            _ = fi10
        }
        v6 = append(v6, (v5 % ((v3) | 1)))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
