package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32 }
type GenS2 struct { F0 uint16; F1 bool }

func main() {
    var v1 uint64 = 728
    v2 := make(map[string]int)
    var v3 int32 = -258
    var v4 bool = ((^(v1)) <= v1)
    var v5 int = -9223372036854775808
    for fi6 := 0; fi6 < 9; fi6++ {
        v5 -= (v5 - (fi6 + fi6))
        _ = fi6
    }
    if (18446744073709551615 < 65535) {
        fmt.Println(v4, v5, v5)
        _ = len(v2)
    } else {
        fmt.Println(v5, v5, v3)
        var v7 int8 = -1
        v2[((("hello" + "\n\t") + ("" + ("hello" + "a"))) + ("foo bar" + ""))] = (v5 / ((v5) | 1))
        _ = v7
    }
    v4 = (v4 && (-24 < 42))
    for fi8 := 0; fi8 < 13; fi8++ {
        v1 = v1
        var v9 uint64 = (v1 - (v1 >> ((1) & 63)))
        _ = fi8
        _ = v9
    }
    var v10 int = (-(v5))
    v5 = (v10 << ((1) & 63))
    v11 := GenS1{F0: v1, F1: -689}
    if (!v4) {
        v12 := v2[((("\n\t" + (("" + ((("a" + "hello") + ("a" + "a")) + ("foo bar" + ("hello" + "foo bar")))) + ("a" + "a"))) + "foo bar") + ("hello" + "a"))]
        _ = v12
    } else {
        var v13 uint8 = 245
        v14 := GenS2{F0: 2, F1: (0 <= v3)}
        var v15 int16 = 32767
        _ = v13
        _ = v14
        _ = v15
    }
    var v16 uint8 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
