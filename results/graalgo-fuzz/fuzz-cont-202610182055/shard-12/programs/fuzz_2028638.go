package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint; F2 uint }
type GenS2 struct { F0 uint32; F1 int32 }

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[string]uint32)
    var v3 bool = true
    v4 := make(map[int]uint32)
    v5 := GenS1{F0: (770 >= 981), F1: 702, F2: 65535}
    var v6 int8 = -7
    for fi7 := 0; fi7 < 1; fi7++ {
        v8 := GenS1{F0: v3, F1: 2, F2: 0}
        _ = fi7
        _ = v8
    }
    v3 = (!((v6 ^ -128) != v6))
    switch "hello" {
    case "hello":
        if false {
            for fi9 := 0; fi9 < 15; fi9++ {
                v10 := GenS1{F0: (-2.25 > 1e10), F1: 255, F2: 1}
                _ = fi9
                _ = v10
            }
            var v11 uint64 = 18446744073709551615
            v12 := GenS2{F0: 2, F1: 2147483647}
            _ = v11
            _ = v12
        }
    case "a":
        var v13 uint = 1
        _ = v13
    case "foo bar":
        switch 32767 {
        case -153:
            v14 := []int{}
            _ = v14
        case 927:
            if v3 {
                var v15 int = -9223372036854775808
                _ = v15
            } else {
                if (42 <= -881) {
                    v16 := []int{9223372036854775807}
                    v17 := GenS2{F0: 42, F1: 681}
                    _ = v16
                    _ = v17
                }
            }
        case -9223372036854775808:
            v18 := GenS1{F0: (v3 || false), F1: 255, F2: 41}
            _ = v18
        }
    default:
        v2["hello"] = 0
    }
    var v19 uint8 = 2
    switch 30 {
    case -589:
        var v20 float64 = 3.14159
        _ = v20
    case 700:
        var v21 uint8 = v19
        _ = v21
    }
    v22 := GenS2{F0: 1, F1: 42}
    v23 := GenS1{F0: (-2147483648 < 2147483647), F1: 49, F2: 42}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
