package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64 }

func main() {
    v1 := []string{"foo bar", (("\n\t" + ("" + "a")) + ("a" + "a")), "\n\t"}
    var v2 bool = false
    var v3 float64 = -2.25
    var v4 int8 = -101
    var v5 int32 = -2147483648
    if len(v1) > 0 {
        v1[0] = "a"
    }
    switch 1 {
    case 2147483647:
        v1 = append(v1, "\n\t")
    default:
        var v6 uint = 0
        _ = v6
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi7 := 0; fi7 < 6; fi7++ {
        var v8 int32 = int32(v4)
        _ = fi7
        _ = v8
    }
    v2 = v2
    if v2 {
        var v9 uint32 = 65535
        v3 += (v3 - 0.1)
        _ = v9
    }
    v10 := GenS1{F0: 42, F1: 18446744073709551615}
    v11 := GenS1{F0: -265, F1: 18446744073709551615}
    var v12 float64 = v3
    fmt.Println(v5)
    v13 := GenS1{F0: 9223372036854775807, F1: 18446744073709551615}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
