package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }

func main() {
    var v1 uint16 = 65535
    var v2 uint32 = 42
    var v3 int64 = -9223372036854775808
    for fi4 := 0; fi4 < 3; fi4++ {
        fmt.Println(v2, v3)
        var v5 float32 = 0.0
        _ = fi4
        _ = v5
    }
    v6 := GenS1{F0: 255}
    fmt.Println(v2, v2)
    v1 = v1
    var v7 int8 = -56
    v1 = 1
    v1 = 65535
    v2 = v2
    var v8 int8 = (v7 & (-(v7)))
    var v9 int16 = 32767
    var v10 uint64 = 733
    fmt.Println(v9, v3)
    if false {
        for fi11 := 0; fi11 < 16; fi11++ {
            fmt.Println(v7, v1)
            _ = fi11
        }
        var v12 uint64 = v10
        _ = v12
    } else {
        fmt.Println(v8, v8, v2)
        var v13 uint32 = (v2 * 0)
        _ = v13
    }
    if false {
        v10 = v10
        v10 += 350
        var v14 int16 = (v9 | v9)
        _ = v14
    } else {
        for fi15 := 0; fi15 < 1; fi15++ {
            v10 = v10
            var v16 uint32 = 2
            var v17 uint64 = v10
            _ = fi15
            _ = v16
            _ = v17
        }
        v8 += v8
        switch 42 {
        case 127:
            if (false && true) {
                fmt.Println(v7, v10)
                var v18 uint8 = 255
                fmt.Println(v9, v7, v3)
                _ = v18
            } else {
                v19 := []int{}
                fmt.Println(v10, v7)
                _ = v19
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
