package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8; F2 uint32 }
type GenS2 struct { F0 string }

func main() {
    var v1 uint16 = 0
    v2 := make(map[uint32]string)
    var v3 uint = 1
    v4 := GenS1{F0: 42, F1: int8(v1), F2: 42}
    v1 = (v1 & v1)
    for fi5 := 0; fi5 < 15; fi5++ {
        v6 := GenS1{F0: 981, F1: 1, F2: 765}
        _ = fi5
        _ = v6
    }
    for fi7 := 0; fi7 < 20; fi7++ {
        v8 := GenS1{F0: 4294967295, F1: 127, F2: 0}
        switch 2147483647 {
        case 32767:
            fmt.Println(v3, v3, v1)
        case 4294967295:
            _ = len(v2)
        case 127:
            v2[232] = "hello"
        default:
            var v9 bool = (42 == 100)
            _ = v9
        }
        v10 := GenS1{F0: 622, F1: -128, F2: 42}
        _ = fi7
        _ = v10
        _ = v8
    }
    if false {
        v3 = uint(v1)
        v3 -= v3
    } else {
        v11 := GenS1{F0: 279, F1: -72, F2: 2147483648}
        var v12 int64 = 0
        fmt.Println(v3)
        _ = v11
        _ = v12
    }
    fmt.Println(v3, v3)
    switch "" {
    case "":
        v13 := GenS2{F0: ""}
        _ = v13
    case "foo bar":
        switch (("\n\t" + (("" + "a") + "")) + "") {
        case "hello":
            var v14 uint16 = (v1 ^ (v1 + (v1 + (v1 + 1))))
            _ = v14
        default:
            v15 := []int{728, -2147483648, 128}
            _ = v15
        }
    case "a":
        v3 -= v3
    default:
        fmt.Println(v1, v3)
    }
    v3 = (v3 & v3)
    var v16 int64 = -184
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
