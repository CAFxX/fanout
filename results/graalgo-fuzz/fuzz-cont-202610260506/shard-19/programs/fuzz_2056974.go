package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8 }

func main() {
    v1 := make(map[string]bool)
    var v2 uint8 = 1
    var v3 uint = 255
    v2 = (v2 << ((v3) & 63))
    for fi4 := 0; fi4 < 14; fi4++ {
        switch (fi4 + (fi4 ^ fi4)) {
        case 9223372036854775807:
            var v5 int64 = 1
            _ = v5
        case -1:
            v2 = (v2 | (v2 - v2))
        case 1:
            v3 = v3
        }
        var v6 float64 = 1.5
        _ = fi4
        _ = v6
    }
    var v7 int16 = -54
    for fi8 := 0; fi8 < 9; fi8++ {
        var v9 uint32 = 153
        v10 := GenS1{F0: 37, F1: uint8(v9)}
        _ = fi8
        _ = v10
        _ = v9
    }
    fmt.Println(v2, v3)
    v2 = 167
    for fi11 := 0; fi11 < 4; fi11++ {
        if (true || (false && true)) {
            for fi12 := 0; fi12 < 18; fi12++ {
                if false {
                    var v13 float64 = 3.14159
                    v14 := GenS1{F0: 127, F1: v2}
                    _ = v13
                    _ = v14
                } else {
                    var v15 float32 = 1e10
                    _ = v15
                }
                var v16 float64 = 0.1
                _ = fi12
                _ = v16
            }
            fmt.Println(v2, v7)
        }
        _ = fi11
    }
    v3 = v3
    var v17 int = -2147483648
    var v18 int32 = 2147483647
    v3 = v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v7
}
