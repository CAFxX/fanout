package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 int64 }

func main() {
    v1 := []string{"foo bar", (("hello" + "hello") + "a"), "foo bar"}
    var v2 int = -1
    var v3 string = "hello"
    var v4 uint16 = 42
    if false {
        v4 -= (v4 >> ((65535) & 63))
    } else {
        var v5 int16 = -1
        if len(v1) > 0 {
            v1[0] = v3
        }
        v6 := GenS1{F0: -128, F1: 2, F2: 873}
        _ = v5
        _ = v6
    }
    var v7 uint = 961
    v7 -= (v7 & (v7 >> ((0) & 63)))
    v7 -= 65535
    var v8 int64 = 42
    switch "" {
    case "a":
        if false {
            if (!true) {
                var v9 string = v3
                _ = v9
            } else {
                var v10 float32 = ((1e10 / 1e10) + 1e10)
                _ = v10
            }
        }
    case "":
        v7 = v7
    case "foo bar":
        var v11 int = (v2 - 4294967295)
        _ = v11
    default:
        if len(v1) > 0 {
            v1[0] = v3
        }
    }
    v7 = 42
    var v12 int = v2
    switch v2 {
    case 42:
        var v13 uint = v7
        _ = v13
    case 525:
        v7 -= v7
    default:
        v1 = append(v1, v3)
    }
    for fi14 := 0; fi14 < 9; fi14++ {
        fmt.Println(v12)
        var v15 bool = (-9223372036854775808 <= v8)
        v4 -= 42
        _ = fi14
        _ = v15
    }
    var v16 uint32 = 2147483647
    var v17 string = v3
    v18 := GenS1{F0: 100, F1: (^(v7)), F2: v8}
    if len(v1) > 0 {
        v1[0] = "foo bar"
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
