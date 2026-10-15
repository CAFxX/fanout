package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16 }
type GenS2 struct { F0 int64; F1 uint8 }

func main() {
    v1 := []string{("foo bar" + ("a" + "foo bar")), ""}
    v2 := []string{"a", ("\n\t" + "hello"), ("hello" + (("hello" + "") + "a"))}
    v3 := make(map[uint32]int)
    var v4 string = ""
    var v5 uint32 = 236
    var v6 uint = 4294967295
    v3[(v5 - 4294967295)] = 2147483647
    if false {
        v4 = v4
    }
    if true {
        fmt.Println(v4, v6, v6)
        v4 = v4
        var v7 bool = ((42 != 0) || (65535 != 1))
        _ = v7
    } else {
        v8 := v3[2147483648]
        _ = v8
    }
    switch v4 {
    case "\n\t":
        var v9 int64 = 82
        _ = v9
    case "hello":
        if len(v1) > 0 {
            v1[0] = "\n\t"
        }
    case "":
        var v10 int = 1
        _ = v10
    }
    var v11 uint = (v6 % ((v6) | 1))
    if len(v1) > 0 {
        v1[0] = v4
    }
    if len(v2) > 0 {
        v2[0] = (((v4 + v4) + v4) + v4)
    }
    for fi12 := 0; fi12 < 14; fi12++ {
        var v13 int64 = 9223372036854775807
        for fi14 := 0; fi14 < 1; fi14++ {
            var v15 int16 = 632
            v16 := GenS1{F0: -1, F1: v15}
            v5 = (v5 | v5)
            _ = fi14
            _ = v15
            _ = v16
        }
        _ = len(v3)
        _ = fi12
        _ = v13
    }
    v17 := GenS1{F0: -174, F1: 0}
    v18 := GenS1{F0: 0, F1: 1000}
    v19 := GenS1{F0: 357, F1: 42}
    var v20 uint32 = (^(v5))
    v21 := GenS2{F0: -955, F1: 1}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
