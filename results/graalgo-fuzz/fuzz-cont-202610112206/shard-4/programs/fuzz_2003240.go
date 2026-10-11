package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }

func main() {
    var v1 uint32 = 4
    var v2 string = ("foo bar" + "foo bar")
    var v3 int64 = -202
    if false {
        v4 := []int{-9223372036854775808}
        var v5 int8 = 125
        _ = v4
        _ = v5
    }
    if (-1 > 0) {
        for fi6 := 0; fi6 < 13; fi6++ {
            v2 = v2
            fmt.Println(v1)
            _ = fi6
        }
        v7 := []int{149, 32768}
        _ = v7
    }
    fmt.Println(v2, v1)
    switch 0 {
    case 32768:
        if (585 != 4294967295) {
            for fi8 := 0; fi8 < 11; fi8++ {
                v9 := GenS1{F0: -294}
                _ = fi8
                _ = v9
            }
            var v10 uint64 = 1
            _ = v10
        }
    case -920:
        fmt.Println(v3, v2, v1)
    }
    var v11 int32 = 125
    var v12 float64 = -2.25
    var v13 string = ("a" + (v2 + ("" + v2)))
    for fi14 := 0; fi14 < 20; fi14++ {
        v15 := []int{(fi14 - fi14)}
        v16 := GenS1{F0: 0}
        _ = fi14
        _ = v15
        _ = v16
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
}
