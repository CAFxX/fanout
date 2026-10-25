package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32 }

func main() {
    var v1 uint64 = 1
    v2 := make(map[string]bool)
    var v3 bool = (("hello" + "foo bar") >= "")
    var v4 int64 = -1
    v5 := GenS1{F0: 32767, F1: 2147483647}
    for fi6 := 0; fi6 < 4; fi6++ {
        var v7 int8 = 61
        var v8 float64 = (1.5 * 3.14159)
        v2[(("" + "foo bar") + (("" + "foo bar") + ""))] = (!v3)
        _ = fi6
        _ = v7
        _ = v8
    }
    v3 = false
    var v9 uint = 65535
    switch 32767 {
    case -306:
        v4 += v4
    case 2147483647:
        v9 = v9
    case 850:
        var v10 uint16 = 255
        _ = v10
    }
    v11 := GenS1{F0: 0, F1: 42}
    var v12 bool = (0 != v4)
    switch "foo bar" {
    case "\n\t":
        if (true || v12) {
            var v13 uint = 42
            _ = v13
        }
    case "hello":
        var v14 int = 916
        _ = v14
    }
    v4 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
