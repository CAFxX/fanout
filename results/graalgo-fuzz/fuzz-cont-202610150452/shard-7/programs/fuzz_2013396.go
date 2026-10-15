package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32 }

func main() {
    var v1 int16 = -113
    var v2 int8 = 113
    var v3 int16 = (^(v1))
    var v4 uint64 = 42
    var v5 int64 = -1
    v5 += (v5 * (v5 + v5))
    fmt.Println(v1, v4, v3)
    v4 = (v4 >> ((42) & 63))
    var v6 uint = 910
    v1 += v3
    var v7 uint16 = uint16(v4)
    v4 = v4
    for fi8 := 0; fi8 < 18; fi8++ {
        fmt.Println(v6, v7)
        _ = fi8
    }
    var v9 int = -152
    for fi10 := 0; fi10 < 11; fi10++ {
        var v11 int32 = 1
        v12 := GenS1{F0: (!true), F1: 42}
        _ = fi10
        _ = v11
        _ = v12
    }
    var v13 uint8 = 255
    if true {
        if false {
            v14 := []int{(v9 + v9), v9, (v9 * 2147483647)}
            _ = v14
        } else {
            fmt.Println(v9)
        }
        for fi15 := 0; fi15 < 20; fi15++ {
            fmt.Println(v1, v1)
            var v16 string = "foo bar"
            _ = fi15
            _ = v16
        }
    } else {
        v17 := []int{2147483647, (v9 >> (((^(v6))) & 63)), 128}
        if false {
            for fi18 := 0; fi18 < 3; fi18++ {
                if len(v17) > 0 {
                    v17[0] = (^(fi18))
                }
                fmt.Println(v7)
                _ = fi18
            }
            var v19 int = -456
            _ = v19
        }
        _ = v17
    }
    var v20 float32 = (((-2.25 / 3.14159) - 0.1) + (0.0 + 0.0))
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
