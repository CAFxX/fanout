package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }

func main() {
    var v1 int64 = 9223372036854775807
    v2 := []string{("foo bar" + "\n\t"), "foo bar"}
    var v3 int = 723
    var v4 int = 9223372036854775807
    var v5 int64 = 9223372036854775807
    v1 -= (v1 | v5)
    for fi6 := 0; fi6 < 9; fi6++ {
        fmt.Println(fi6, v5)
        _ = fi6
    }
    switch "foo bar" {
    case "":
        for fi7 := 0; fi7 < 7; fi7++ {
            var v8 uint32 = 42
            var v9 int16 = 1
            v10 := GenS1{F0: 820}
            _ = fi7
            _ = v10
            _ = v8
            _ = v9
        }
    }
    v1 = (v1 >> ((4294967295) & 63))
    fmt.Println(v4, v5)
    fmt.Println(v1)
    var v11 int8 = 42
    var v12 float64 = 1.5
    switch v3 {
    case 127:
        for fi13 := 0; fi13 < 20; fi13++ {
            fmt.Println(fi13, v3, v3)
            var v14 int32 = -1
            var v15 int64 = 685
            _ = fi13
            _ = v14
            _ = v15
        }
    case 1:
        var v16 uint8 = 128
        _ = v16
    case 42:
        if len(v2) > 0 {
            v2[0] = "hello"
        }
    default:
        var v17 uint8 = 255
        _ = v17
    }
    for fi18 := 0; fi18 < 9; fi18++ {
        v4 = (v4 / (((v3 >> ((42) & 63))) | 1))
        _ = fi18
    }
    v11 -= v11
    _ = len(v2)
    if false {
        var v19 uint8 = 83
        v5 = v5
        _ = v19
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
