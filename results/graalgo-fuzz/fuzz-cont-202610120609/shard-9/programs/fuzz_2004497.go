package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64; F2 int32 }

func main() {
    v1 := make(map[uint32]string)
    v2 := []string{"a", (("\n\t" + ("foo bar" + "")) + "\n\t"), ""}
    v3 := make(map[uint32]string)
    var v4 bool = true
    v4 = (42 <= 4294967295)
    v5 := GenS1{F0: -32768, F1: 18446744073709551615, F2: 1}
    v4 = (v4 && true)
    if v4 {
        if len(v2) > 0 {
            v2[0] = "foo bar"
        }
        v6 := GenS1{F0: 827, F1: 0, F2: 1}
        v4 = (true || (!v4))
        _ = v6
    }
    var v7 int64 = -559
    v4 = (3.14159 != (1e10 / 1e10))
    var v8 uint32 = 0
    var v9 int16 = -32768
    for fi10 := 0; fi10 < 9; fi10++ {
        v11 := v1[v8]
        v12 := GenS1{F0: (-(v9)), F1: 42, F2: 42}
        var v13 uint32 = (v8 >> ((979) & 63))
        _ = fi10
        _ = v11
        _ = v12
        _ = v13
    }
    for fi14 := 0; fi14 < 7; fi14++ {
        v8 = (^(v8))
        _ = fi14
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
