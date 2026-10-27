package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64 }

func main() {
    var v1 float64 = float64(1.5)
    var v2 float32 = 3.14159
    var v3 int32 = -2147483648
    v4 := []string{("foo bar" + "hello"), ("" + "a"), "foo bar", ""}
    var v5 uint64 = 18446744073709551615
    v1 = 1.5
    var v6 uint8 = 127
    fmt.Println(v2, v6)
    var v7 int16 = -32768
    switch ("" + ("foo bar" + "\n\t")) {
    case "\n\t":
        if (!true) {
            v8 := GenS1{F0: v6, F1: -1}
            _ = v8
        }
    case "":
        for fi9 := 0; fi9 < 12; fi9++ {
            var v10 uint = 1
            _ = fi9
            _ = v10
        }
    default:
        v3 = v3
    }
    var v11 int32 = int32(v7)
    var v12 float32 = v2
    fmt.Println(v12, v7)
    for fi13 := 0; fi13 < 4; fi13++ {
        for fi14 := 0; fi14 < 14; fi14++ {
            var v15 int = fi14
            _ = fi14
            _ = v15
        }
        if false {
            var v16 int16 = 840
            _ = v16
        } else {
            var v17 uint64 = (v5 << ((4294967295) & 63))
            v12 += 1e10
            _ = v17
        }
        var v18 int64 = 42
        _ = fi13
        _ = v18
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
