package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string }

func main() {
    var v1 int8 = 100
    var v2 int16 = -32768
    var v3 int = -9223372036854775808
    v3 -= (v3 + (v3 << ((4294967295) & 63)))
    v1 = v1
    if false {
        var v4 int8 = v1
        fmt.Println(v1)
        _ = v4
    } else {
        var v5 int64 = 42
        v1 += -1
        _ = v5
    }
    var v6 bool = (4294967295 == 1)
    var v7 int32 = -2147483648
    switch (v3 << ((859) & 63)) {
    case 1:
        var v8 string = ""
        _ = v8
    case 987:
        v1 = 34
    }
    v9 := GenS1{F0: (v2 * (-(v2))), F1: "\n\t"}
    switch 478 {
    case 32767:
        v10 := []int{}
        _ = v10
    case 42:
        fmt.Println(v3, v7, v6)
    default:
        for fi11 := 0; fi11 < 7; fi11++ {
            fmt.Println(v6)
            _ = fi11
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v9
}
