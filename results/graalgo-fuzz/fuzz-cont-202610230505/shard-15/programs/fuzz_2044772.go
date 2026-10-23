package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int; F2 uint32 }
type GenS2 struct { F0 uint; F1 string }

func main() {
    var v1 bool = true
    v2 := make(map[string]uint32)
    var v3 uint64 = 42
    var v4 int = 1
    v5 := make(map[string]bool)
    switch (v4 + v4) {
    case 128:
        fmt.Println(v3, v1, v3)
    case -2147483648:
        v4 = (v4 * v4)
    }
    v6 := v5[("a" + "\n\t")]
    v7 := []int{v4, 32768}
    for fi8 := 0; fi8 < 6; fi8++ {
        v3 -= 42
        for fi9 := 0; fi9 < 20; fi9++ {
            fmt.Println(fi9)
            if len(v7) > 0 {
                v7[0] = fi8
            }
            fmt.Println(v3, fi9, v4)
            _ = fi9
        }
        _ = fi8
    }
    if true {
        v10 := GenS1{F0: 42, F1: v4, F2: 2147483648}
        v4 += v4
        fmt.Println(v1)
        _ = v10
    }
    var v11 int16 = -32768
    v3 = 212
    v12 := GenS1{F0: 32767, F1: 128, F2: 2147483647}
    for fi13 := 0; fi13 < 20; fi13++ {
        v14 := GenS2{F0: 0, F1: "hello"}
        var v15 int = (v4 + 854)
        _ = fi13
        _ = v14
        _ = v15
    }
    var v16 int = v4
    var v17 int64 = 1
    delete(v5, ((("foo bar" + "hello") + ("foo bar" + ((("a" + "") + "") + ("hello" + "a")))) + ("a" + (("\n\t" + "a") + ("foo bar" + "hello")))))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
