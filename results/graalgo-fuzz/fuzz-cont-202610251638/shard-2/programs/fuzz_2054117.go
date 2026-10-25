package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }

func main() {
    var v1 int = 9223372036854775807
    var v2 uint16 = 0
    var v3 int = (v1 / (((v1 | (v1 - (v1 / (((v1 * (v1 & (v1 * -2147483648)))) | 1))))) | 1))
    var v4 int16 = 32767
    var v5 uint32 = 65535
    var v6 int8 = 113
    var v7 int = v3
    if (-9223372036854775808 <= (-(v1))) {
        v8 := GenS1{F0: -144}
        _ = v8
    }
    switch v1 {
    case 815:
        var v9 int8 = (v6 + (v6 % ((v6) | 1)))
        _ = v9
    case 1:
        for fi10 := 0; fi10 < 17; fi10++ {
            fmt.Println(v1, v4, v4)
            for fi11 := 0; fi11 < 3; fi11++ {
                var v12 uint16 = v2
                var v13 float64 = (1e10 / (1e10 / 1e10))
                _ = fi11
                _ = v12
                _ = v13
            }
            _ = fi10
        }
    }
    v6 = (-(v6))
    switch "a" {
    case "\n\t":
        var v14 int = (^(v1))
        _ = v14
    }
    fmt.Println(v6, v3, v3)
    var v15 int = v3
    var v16 uint = 255
    var v17 int64 = 9223372036854775807
    v18 := GenS1{F0: 2147483647}
    var v19 uint = v16
    v20 := []int{-47, v7}
    var v21 float64 = 1e10
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
