package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int; F2 int8 }
type GenS2 struct { F0 uint64; F1 int16; F2 int8 }

func main() {
    var v1 int16 = 1000
    v2 := []int{77}
    v3 := make(map[string]uint32)
    v4 := make(map[string]uint32)
    var v5 uint = 4294967295
    v2 = append(v2, 77)
    delete(v4, "hello")
    var v6 uint8 = 2
    v7 := GenS1{F0: 127, F1: 32768, F2: 38}
    switch 720 {
    case -9223372036854775808:
        v8 := v3[("a" + "a")]
        _ = v8
    case -2147483648:
        v6 += (v6 ^ (^(v6)))
    }
    v5 = 65535
    for fi9 := 0; fi9 < 15; fi9++ {
        fmt.Println(v5, v5, v1)
        _ = fi9
    }
    v6 += v6
    switch (("\n\t" + "foo bar") + "") {
    case "":
        for fi10 := 0; fi10 < 20; fi10++ {
            switch -202 {
            case 2147483647:
                v11 := GenS1{F0: 106, F1: 32767, F2: 42}
                _ = v11
            case 32767:
                v12 := GenS2{F0: 4294967295, F1: (v1 - 412), F2: 1}
                _ = v12
            default:
                v5 -= (v5 / ((18446744073709551615) | 1))
            }
            _ = fi10
        }
    case "hello":
        v6 = v6
    case "\n\t":
        var v13 int8 = 0
        _ = v13
    }
    v5 = v5
    var v14 int16 = (v1 >> ((1) & 63))
    var v15 int = 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
