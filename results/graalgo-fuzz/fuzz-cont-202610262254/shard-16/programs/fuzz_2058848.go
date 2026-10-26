package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int; F1 int64 }

func main() {
    var v1 uint8 = 127
    v2 := make(map[string]uint32)
    var v3 uint32 = 1
    var v4 int8 = 127
    switch "\n\t" {
    case "":
        var v5 uint8 = v1
        _ = v5
    case "hello":
        if (!true) {
            fmt.Println(v4, v1)
        }
    case "\n\t":
        switch ("\n\t" + "foo bar") {
        case "hello":
            v4 = v4
        case "":
            for fi6 := 0; fi6 < 10; fi6++ {
                v3 = v3
                _ = fi6
            }
        default:
            delete(v2, "")
        }
    }
    v1 += v1
    for fi7 := 0; fi7 < 6; fi7++ {
        var v8 string = ""
        _ = fi7
        _ = v8
    }
    if (-449 > 0) {
        if (true || true) {
            v2[("" + "hello")] = 424
        } else {
            v4 += v4
            v9 := GenS1{F0: 42, F1: -32768}
            _ = v9
        }
    } else {
        var v10 float32 = 1e10
        fmt.Println(v4, v4, v1)
        fmt.Println(v1, v1, v10)
        _ = v10
    }
    for fi11 := 0; fi11 < 2; fi11++ {
        v3 += v3
        v12 := GenS1{F0: 341, F1: 42}
        _ = fi11
        _ = v12
    }
    fmt.Println(v1, v1)
    for fi13 := 0; fi13 < 17; fi13++ {
        for fi14 := 0; fi14 < 18; fi14++ {
            v15 := GenS2{F0: 0}
            v16 := GenS3{F0: -292, F1: -952}
            v4 += (v4 | -64)
            _ = fi14
            _ = v15
            _ = v16
        }
        var v17 int16 = -1
        _ = fi13
        _ = v17
    }
    v1 -= v1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}
