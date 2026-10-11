package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int64; F1 uint64 }

func main() {
    var v1 int64 = 0
    v2 := []string{}
    var v3 uint64 = 18446744073709551615
    var v4 int32 = int32(v3)
    var v5 bool = false
    var v6 int8 = 100
    var v7 float32 = float32((float64(1e10) / 0.1))
    var v8 string = (((("" + ("" + "\n\t")) + (("a" + "\n\t") + "")) + ("hello" + ("a" + "\n\t"))) + (("foo bar" + ((("" + (("hello" + ("a" + ("a" + (("a" + "") + (("\n\t" + "foo bar") + ("foo bar" + "foo bar")))))) + "\n\t")) + ("a" + ("" + "a"))) + "hello")) + "a"))
    _ = len(v2)
    if false {
        var v9 uint = 157
        if true {
            fmt.Println(v8, v7)
        }
        switch -9223372036854775808 {
        case -426:
            fmt.Println(v8)
        case -2147483648:
            v9 -= v9
        }
        _ = v9
    } else {
        v7 = (0.1 / -2.25)
        v10 := GenS1{F0: -428, F1: v5}
        _ = v10
    }
    fmt.Println(v7, v1, v7)
    v11 := GenS2{F0: v5}
    fmt.Println(v6)
    v7 = v7
    v4 -= (v4 << ((255) & 63))
    var v12 uint64 = v3
    fmt.Println(v7, v4, v4)
    for fi13 := 0; fi13 < 16; fi13++ {
        fmt.Println(v12, v4, v6)
        v14 := GenS1{F0: (fi13 << ((996) & 63)), F1: v5}
        v15 := GenS3{F0: v1, F1: (^(v12))}
        _ = fi13
        _ = v14
        _ = v15
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
