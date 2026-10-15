package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string; F2 string }

func main() {
    var v1 int16 = 882
    var v2 uint64 = 4294967295
    v3 := []uint32{4294967295, 2147483648, 65535, 363}
    _ = len(v3)
    if false {
        v4 := GenS1{F0: 877, F1: "hello", F2: "foo bar"}
        fmt.Println(v1, v2)
        var v5 uint = 18446744073709551615
        _ = v4
        _ = v5
    }
    switch "" {
    case "":
        fmt.Println(v1, v1)
    case "\n\t":
        if len(v3) > 0 {
            v3[0] = 208
        }
    default:
        var v6 uint32 = 112
        _ = v6
    }
    fmt.Println(v1, v2)
    var v7 int32 = 1
    v3 = append(v3, 65535)
    var v8 uint = 1
    switch (("hello" + ((("\n\t" + "foo bar") + "hello") + "a")) + "\n\t") {
    case "a":
        _ = len(v3)
    case "foo bar":
        v1 = (v1 | (v1 >> (((v8 << ((65535) & 63))) & 63)))
    }
    v9 := GenS1{F0: 578, F1: "hello", F2: "\n\t"}
    for fi10 := 0; fi10 < 19; fi10++ {
        var v11 int8 = -12
        _ = fi10
        _ = v11
    }
    var v12 int16 = (v1 >> (((v8 / ((v8) | 1))) & 63))
    if true {
        v13 := GenS1{F0: 412, F1: ("hello" + "\n\t"), F2: "a"}
        v2 += v2
        _ = len(v3)
        _ = v13
    }
    switch 42 {
    case 2147483647:
        _ = len(v3)
    case 127:
        var v14 int = 9223372036854775807
        _ = v14
    case 32767:
        fmt.Println(v12, v8)
    default:
        for fi15 := 0; fi15 < 1; fi15++ {
            var v16 int32 = (v7 + v7)
            var v17 uint = (v8 % (((v8 + (v8 + (v8 - 4294967295)))) | 1))
            _ = fi15
            _ = v16
            _ = v17
        }
    }
    v2 -= (v2 | v2)
    v2 += v2
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
