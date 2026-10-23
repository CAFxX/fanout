package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16 }
func genFn0(p0 uint64, p1 uint) uint16 {
    var v7 string = "a"
    var v8 int16 = 608
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 255
}

func main() {
    v1 := []string{("a" + (("foo bar" + "a") + "a")), ((("foo bar" + "foo bar") + (("hello" + "foo bar") + "hello")) + ("" + "hello")), ("\n\t" + ""), "hello"}
    var v2 int = 128
    var v3 int32 = 31
    v4 := []string{((("\n\t" + "hello") + "a") + "a"), "foo bar", "a"}
    v5 := make(map[string]bool)
    var v6 float64 = (0.0 / ((3.14159 / 1e10) - 0.0))
    if (!false) {
        var v9 int32 = (v3 >> ((549) & 63))
        _ = v9
    }
    var v10 bool = false
    var v11 float32 = 1e10
    var v12 int = v2
    v12 -= v12
    switch 32767 {
    case 1:
        var v13 int = 4294967295
        _ = v13
    default:
        for fi14 := 0; fi14 < 4; fi14++ {
            v15 := GenS1{F0: 2147483647, F1: 42}
            _ = fi14
            _ = v15
        }
    }
    v16 := GenS1{F0: 2147483648, F1: 32767}
    switch (("hello" + "") + "") {
    case "\n\t":
        v10 = false
    case "foo bar":
        v2 = (v2 - (v12 / ((42) | 1)))
    case "hello":
        var v17 int16 = 1
        _ = v17
    default:
        var v18 int = (v12 - 127)
        _ = v18
    }
    switch "foo bar" {
    case "foo bar":
        var v19 uint = 4294967295
        _ = v19
    default:
        fmt.Println(v2)
    }
    switch v12 {
    case -299:
        delete(v5, "foo bar")
    case 32767:
        fmt.Println(v3, v10, v3)
    default:
        v2 += v2
    }
    v12 = 42
    v20 := GenS1{F0: 839, F1: 42}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v16
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
