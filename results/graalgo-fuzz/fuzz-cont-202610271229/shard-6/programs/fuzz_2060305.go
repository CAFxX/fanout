package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int }
func genFn0() uint32 {
    var v5 int32 = -829
    v6 := []int{721, 128}
    for fi7 := 0; fi7 < 5; fi7++ {
        v8 := GenS1{F0: -94, F1: 128}
        var v9 uint8 = 255
        fmt.Println(v9, v9, v9)
        _ = fi7
        _ = v8
        _ = v9
    }
    _ = v5
    _ = v6
    return 4294967295
}

func main() {
    var v1 uint = 255
    var v2 uint8 = 255
    var v3 int8 = 1
    var v4 int16 = 1000
    v3 -= 100
    if true {
        var v10 int16 = v4
        _ = v10
    }
    var v11 int = 2147483647
    var v12 int8 = v3
    fmt.Println(v12, v2, v11)
    v13 := GenS1{F0: v12, F1: 32768}
    switch (("\n\t" + (("foo bar" + "a") + "")) + "a") {
    case "hello":
        if (!false) {
            v14 := []int{(v11 << (((v1 & v1)) & 63)), (v11 - (-(v11)))}
            _ = v14
        }
    case "a":
        v15 := []int{(v11 & v11), -2147483648, v11, (v11 % ((v11) | 1))}
        _ = v15
    case "\n\t":
        v2 -= (v2 & (^(v2)))
    default:
        for fi16 := 0; fi16 < 12; fi16++ {
            v17 := []int{-9223372036854775808}
            fmt.Println(v12, fi16)
            _ = fi16
            _ = v17
        }
    }
    for fi18 := 0; fi18 < 8; fi18++ {
        var v19 bool = false
        _ = fi18
        _ = v19
    }
    var v20 uint8 = (v2 ^ (v2 * (^(v2))))
    var v21 int32 = -76
    v2 = 82
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
}
