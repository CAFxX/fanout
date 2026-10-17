package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint }

func main() {
    v1 := make(map[int]string)
    v2 := make(map[string]string)
    var v3 string = "a"
    var v4 uint32 = 42
    v3 = ("a" + ((("foo bar" + ("a" + v3)) + (v3 + v3)) + (("" + (v3 + (("" + v3) + v3))) + v3)))
    switch -2147483648 {
    case 79:
        var v5 bool = (0.0 <= -2.25)
        _ = v5
    default:
        var v6 int32 = 535
        _ = v6
    }
    for fi7 := 0; fi7 < 17; fi7++ {
        for fi8 := 0; fi8 < 6; fi8++ {
            v9 := GenS1{F0: 402, F1: 42}
            _ = fi8
            _ = v9
        }
        v4 = v4
        var v10 string = "hello"
        _ = fi7
        _ = v10
    }
    v11 := GenS1{F0: 65535, F1: 42}
    var v12 int16 = -32768
    var v13 uint16 = 0
    v4 = v4
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
}
