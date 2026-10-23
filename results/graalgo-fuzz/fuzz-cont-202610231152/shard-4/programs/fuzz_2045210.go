package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32; F2 uint32 }

func main() {
    var v1 float32 = -2.25
    var v2 uint = 4294967295
    var v3 uint8 = 1
    var v4 uint32 = 2
    if (163 == 2147483647) {
        fmt.Println(v2)
        if true {
            v2 -= v2
            var v5 uint16 = 2
            _ = v5
        }
    }
    v6 := []int{}
    var v7 string = ((("foo bar" + "") + "") + ("\n\t" + ("\n\t" + "foo bar")))
    v3 += 2
    var v8 int64 = 9223372036854775807
    switch v7 {
    case "":
        v9 := GenS1{F0: (v3 << (((v2 / (((^(v2))) | 1))) & 63)), F1: uint32(v2), F2: v4}
        _ = v9
    default:
        var v10 bool = false
        _ = v10
    }
    v11 := GenS1{F0: v3, F1: 2, F2: 42}
    fmt.Println(v2)
    v3 -= 141
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
}
