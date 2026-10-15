package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint; F2 int32 }
type GenS2 struct { F0 uint16; F1 int32; F2 uint8 }

func main() {
    var v1 string = "hello"
    var v2 int64 = 1
    var v3 bool = true
    var v4 int8 = 0
    var v5 string = (v1 + "a")
    v6 := make(map[uint32]bool)
    v5 = (v1 + (v1 + (("\n\t" + v1) + (v1 + v1))))
    switch 0 {
    case 2147483647:
        var v7 int = 233
        _ = v7
    }
    v8 := GenS1{F0: 1, F1: 1, F2: -2147483648}
    v1 = (v1 + "")
    switch "" {
    case "hello":
        v2 += -9223372036854775808
    case "\n\t":
        v3 = true
    case "":
        if (1 >= 42) {
            var v9 bool = v3
            fmt.Println(v3, v2)
            _ = v9
        }
    }
    var v10 int32 = 42
    for fi11 := 0; fi11 < 10; fi11++ {
        v12 := []int{}
        _ = fi11
        _ = v12
    }
    var v13 int16 = -32768
    if (!(((283 < 283) || true) || false)) {
        var v14 int32 = -1
        var v15 int16 = (v13 & (^(v13)))
        v1 = v5
        _ = v14
        _ = v15
    }
    v16 := GenS2{F0: 531, F1: -47, F2: 0}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v13
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
