package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[uint32]uint32)
    v3 := make(map[string]bool)
    v4 := []string{"\n\t", "hello", "\n\t"}
    v5 := make(map[string]bool)
    delete(v3, ("\n\t" + "hello"))
    var v6 uint32 = 697
    v6 -= 4294967295
    if true {
        v4 = append(v4, (((("" + "a") + ("\n\t" + "")) + "foo bar") + ("a" + "hello")))
        var v7 int8 = 1
        v7 -= v7
        _ = v7
    } else {
        v8 := GenS1{F0: v6}
        v6 = (v6 + 820)
        v6 = 42
        _ = v8
    }
    var v9 uint32 = 0
    var v10 int64 = -1
    for fi11 := 0; fi11 < 3; fi11++ {
        var v12 int = fi11
        _ = fi11
        _ = v12
    }
    var v13 uint16 = 65535
    for fi14 := 0; fi14 < 3; fi14++ {
        for fi15 := 0; fi15 < 2; fi15++ {
            var v16 float32 = (-2.25 * -2.25)
            v13 = v13
            v10 = (v10 / (((v10 + (-(v10)))) | 1))
            _ = fi15
            _ = v16
        }
        v13 += v13
        _ = fi14
    }
    fmt.Println(v9, v10)
    v9 += (v6 & (^(v6)))
    switch 32767 {
    case 4294967295:
        var v17 int64 = (v10 ^ (v10 << ((65535) & 63)))
        _ = v17
    default:
        var v18 bool = false
        _ = v18
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
