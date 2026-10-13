package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint; F2 string }

func main() {
    var v1 int32 = -1
    var v2 bool = true
    var v3 uint64 = 18446744073709551615
    v4 := make(map[uint32]bool)
    fmt.Println(v2, v3)
    if (!v2) {
        for fi5 := 0; fi5 < 20; fi5++ {
            v6 := GenS1{F0: 603, F1: 487, F2: ((("" + "a") + "foo bar") + ("foo bar" + "\n\t"))}
            _ = fi5
            _ = v6
        }
        var v7 int8 = 100
        _ = v7
    } else {
        fmt.Println(v3, v3, v3)
        delete(v4, 2147483648)
    }
    switch "foo bar" {
    case "":
        v8 := GenS1{F0: -988, F1: 1, F2: "foo bar"}
        _ = v8
    case "foo bar":
        v9 := GenS1{F0: 0, F1: 709, F2: ("foo bar" + ("\n\t" + "hello"))}
        _ = v9
    case "hello":
        var v10 int64 = -9223372036854775808
        _ = v10
    default:
        v2 = v2
    }
    switch "foo bar" {
    case "foo bar":
        var v11 uint64 = (v3 - (^(v3)))
        _ = v11
    }
    v3 -= (v3 * v3)
    var v12 int = 42
    fmt.Println(v2, v1, v12)
    fmt.Println(v3, v1)
    v4[42] = (988 > v3)
    for fi13 := 0; fi13 < 14; fi13++ {
        v14 := GenS1{F0: -1, F1: 4294967295, F2: ""}
        var v15 int16 = -1
        v2 = (v2 && v2)
        _ = fi13
        _ = v14
        _ = v15
    }
    for fi16 := 0; fi16 < 11; fi16++ {
        if false {
            v1 = v1
            _ = len(v4)
            var v17 float32 = ((0.1 + 0.0) / 3.14159)
            _ = v17
        }
        fmt.Println(v2, v2)
        _ = fi16
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
}
