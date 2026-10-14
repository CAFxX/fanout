package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int; F1 uint64; F2 string }

func main() {
    var v1 int64 = -1
    var v2 string = ("a" + "foo bar")
    v3 := make(map[string]uint32)
    var v4 float32 = 1.5
    var v5 int32 = 1
    v6 := make(map[string]uint32)
    v7 := GenS1{F0: 0}
    fmt.Println(v5)
    var v8 int = 128
    v2 = v2
    fmt.Println(v2, v4)
    var v9 int64 = v1
    for fi10 := 0; fi10 < 8; fi10++ {
        var v11 bool = true
        v1 = v1
        var v12 int = -636
        _ = fi10
        _ = v11
        _ = v12
    }
    v3[(v2 + (v2 + v2))] = 0
    v4 = ((1e10 * -2.25) / 0.1)
    for fi13 := 0; fi13 < 19; fi13++ {
        v14 := GenS1{F0: (v1 ^ (v9 | v1))}
        _ = fi13
        _ = v14
    }
    var v15 string = v2
    v6[("foo bar" + "hello")] = 4294967295
    v16 := GenS2{F0: (v8 % (((v8 << ((18446744073709551615) & 63))) | 1)), F1: 42, F2: v2}
    var v17 uint64 = 4294967295
    v15 = (("foo bar" + ("\n\t" + ("" + (v2 + (("" + "") + "a"))))) + "")
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
