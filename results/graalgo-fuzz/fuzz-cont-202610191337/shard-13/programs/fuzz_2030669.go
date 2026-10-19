package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint16; F1 uint; F2 bool }

func main() {
    v1 := make(map[string]bool)
    var v2 int16 = 42
    v3 := []string{"a", "hello", (("hello" + ("foo bar" + "\n\t")) + "a"), (("\n\t" + "a") + "foo bar")}
    var v4 int32 = -1
    var v5 uint8 = 180
    var v6 int = -9223372036854775808
    var v7 int = int(v5)
    v5 = v5
    for fi8 := 0; fi8 < 2; fi8++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        switch (v7 & (v6 & v7)) {
        case 32768:
            switch v7 {
            case -112:
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            default:
                v4 = 0
            }
        case 2147483647:
            var v9 bool = false
            _ = v9
        }
        v10 := GenS1{F0: (v5 & (v5 ^ 97))}
        _ = fi8
        _ = v10
    }
    var v11 int64 = 0
    v7 += v7
    delete(v1, "")
    for fi12 := 0; fi12 < 7; fi12++ {
        fmt.Println(v5)
        for fi13 := 0; fi13 < 10; fi13++ {
            fmt.Println(v4)
            v14 := GenS2{F0: 255, F1: 42, F2: false}
            var v15 bool = true
            var v16 int = (-(v6))
            _ = fi13
            _ = v14
            _ = v15
            _ = v16
        }
        _ = fi12
    }
    switch v7 {
    case 32768:
        if len(v3) > 0 {
            v3[0] = "foo bar"
        }
    case 921:
        var v17 uint64 = 2
        _ = v17
    case 32767:
        v5 = (v5 << ((174) & 63))
    default:
        var v18 bool = true
        _ = v18
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
