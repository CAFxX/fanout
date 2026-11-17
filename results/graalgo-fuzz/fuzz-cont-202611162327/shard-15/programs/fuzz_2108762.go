package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string }

func main() {
    var v1 int32 = 1
    v2 := make(map[uint32]int)
    var v3 uint = 42
    var v4 int = -701
    var v5 uint16 = 255
    v6 := make(map[uint32]string)
    for fi7 := 0; fi7 < 8; fi7++ {
        fmt.Println(v1, v3)
        for fi8 := 0; fi8 < 10; fi8++ {
            var v9 uint = (v3 | 0)
            _ = fi8
            _ = v9
        }
        _ = fi7
    }
    v4 = 4294967295
    v10 := GenS1{F0: -54, F1: ("foo bar" + ("hello" + ((("a" + "foo bar") + "hello") + "")))}
    v3 = v3
    v4 -= (v4 & 32767)
    v5 = (^(v5))
    for fi11 := 0; fi11 < 8; fi11++ {
        v2[2147483647] = fi11
        for fi12 := 0; fi12 < 10; fi12++ {
            var v13 int8 = 1
            _ = fi12
            _ = v13
        }
        _ = fi11
    }
    v14 := GenS1{F0: 1, F1: ""}
    var v15 uint16 = v5
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
