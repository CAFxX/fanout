package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }

func main() {
    v1 := []uint32{1, 762, 4294967295, 2147483648}
    v2 := make(map[int]string)
    var v3 uint = 315
    var v4 int64 = 42
    var v5 string = ("hello" + ("a" + "a"))
    v6 := GenS1{F0: -502}
    v4 += 9223372036854775807
    switch 186 {
    case -126:
        var v7 int16 = 1000
        _ = v7
    default:
        var v8 float32 = 0.0
        _ = v8
    }
    if false {
        for fi9 := 0; fi9 < 16; fi9++ {
            v4 += (-(v4))
            _ = fi9
        }
        switch "a" {
        case "foo bar":
            v3 = 65535
        default:
            var v10 int16 = -1
            _ = v10
        }
        var v11 string = ((("hello" + ("" + ("a" + "a"))) + "foo bar") + v5)
        _ = v11
    }
    var v12 int64 = v4
    var v13 string = v5
    v14 := v2[32767]
    var v15 int16 = 42
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
