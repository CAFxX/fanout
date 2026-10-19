package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16; F2 int32 }
func genFn0(p0 int8, p1 uint64) uint16 {
    var v6 int32 = -723
    var v7 string = "hello"
    v6 -= (^(v6))
    fmt.Println(v6)
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 1
}

func main() {
    var v1 uint64 = 910
    v2 := make(map[int]bool)
    var v3 uint16 = 2
    var v4 int8 = -128
    v5 := []string{((("hello" + ("" + "hello")) + ("hello" + ("a" + "\n\t"))) + "a")}
    switch -2147483648 {
    case 4294967295:
        var v8 int16 = 32767
        _ = v8
    default:
        var v9 uint32 = 752
        _ = v9
    }
    var v10 uint16 = v3
    switch 42 {
    case 128:
        v2[0] = (!(2 <= 65535))
    case -9223372036854775808:
        fmt.Println(v3, v3, v1)
    case -788:
        var v11 uint64 = v1
        _ = v11
    default:
        v2[394] = false
    }
    v12 := GenS1{F0: 255, F1: -32768, F2: 636}
    if (false || ((v1 < 273) && true)) {
        if true {
            for fi13 := 0; fi13 < 14; fi13++ {
                _ = len(v5)
                var v14 bool = false
                var v15 string = "foo bar"
                _ = fi13
                _ = v14
                _ = v15
            }
        }
    }
    v10 -= (v3 / (((v10 % ((1) | 1))) | 1))
    v10 = 0
    v1 -= 892
    v5 = append(v5, "hello")
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
